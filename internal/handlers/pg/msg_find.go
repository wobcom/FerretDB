// Copyright 2021 FerretDB Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pg

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v4"

	"github.com/FerretDB/FerretDB/internal/handlers/common"
	"github.com/FerretDB/FerretDB/internal/handlers/pg/pgdb"
	"github.com/FerretDB/FerretDB/internal/types"
	"github.com/FerretDB/FerretDB/internal/util/iterator"
	"github.com/FerretDB/FerretDB/internal/util/lazyerrors"
	"github.com/FerretDB/FerretDB/internal/util/must"
	"github.com/FerretDB/FerretDB/internal/wire"
)

// MsgFind implements HandlerInterface.
func (h *Handler) MsgFind(ctx context.Context, msg *wire.OpMsg) (*wire.OpMsg, error) {
	dbPool, err := h.DBPool(ctx)
	if err != nil {
		return nil, lazyerrors.Error(err)
	}

	document, err := msg.Document()
	if err != nil {
		return nil, lazyerrors.Error(err)
	}

	params, err := common.GetFindParams(document, h.L)
	if err != nil {
		return nil, err
	}

	if params.MaxTimeMS != 0 {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Duration(params.MaxTimeMS)*time.Millisecond)
		defer cancel()

		ctx = ctxWithTimeout
	}

	qp := pgdb.QueryParams{
		DB:         params.DB,
		Collection: params.Collection,
		Comment:    params.Comment,
		Filter:     params.Filter,
	}

	// get comment from query, e.g. db.collection.find({$comment: "test"})
	if qp.Filter != nil {
		if qp.Comment, err = common.GetOptionalParam(qp.Filter, "$comment", qp.Comment); err != nil {
			return nil, err
		}
	}

	var resDocs []*types.Document
	err = dbPool.InTransaction(ctx, func(tx pgx.Tx) error {
		resDocs, err = fetchAndFilterDocs(ctx, &fetchParams{tx, &qp, h.DisablePushdown})
		return err
	})

	if err != nil {
		return nil, err
	}

	if err = common.SortDocuments(resDocs, params.Sort); err != nil {
		return nil, err
	}

	if resDocs, err = common.LimitDocuments(resDocs, params.Limit); err != nil {
		return nil, err
	}

	if err = common.ProjectDocuments(resDocs, params.Projection); err != nil {
		return nil, err
	}

	firstBatch := types.MakeArray(len(resDocs))
	for _, doc := range resDocs {
		firstBatch.Append(doc)
	}

	var reply wire.OpMsg
	must.NoError(reply.SetSections(wire.OpMsgSection{
		Documents: []*types.Document{must.NotFail(types.NewDocument(
			"cursor", must.NotFail(types.NewDocument(
				"firstBatch", firstBatch,
				"id", int64(0), // TODO
				"ns", qp.DB+"."+qp.Collection,
			)),
			"ok", float64(1),
		))},
	}))

	return &reply, nil
}

// fetchParams is used to pass parameters to fetchAndFilterDocs.
type fetchParams struct {
	tx              pgx.Tx
	qp              *pgdb.QueryParams
	disablePushdown bool
}

// fetchAndFilterDocs fetches documents from the database and filters them using the provided sqlParam.Filter.
func fetchAndFilterDocs(ctx context.Context, fp *fetchParams) ([]*types.Document, error) {
	// filter is used to filter documents on the FerretDB side,
	// qp.Filter is used to filter documents on the PostgreSQL side (query pushdown).
	filter := fp.qp.Filter

	if fp.disablePushdown {
		fp.qp.Filter = nil
	}

	iter, err := pgdb.QueryDocuments(ctx, fp.tx, fp.qp)
	if err != nil {
		return nil, err
	}

	defer iter.Close()

	resDocs := make([]*types.Document, 0, 16)

	for {
		_, doc, err := iter.Next()
		if err != nil {
			if errors.Is(err, iterator.ErrIteratorDone) {
				return resDocs, nil
			}

			return nil, err
		}

		matches, err := common.FilterDocument(doc, filter)
		if err != nil {
			return nil, err
		}

		if !matches {
			continue
		}

		resDocs = append(resDocs, doc)
	}
}
