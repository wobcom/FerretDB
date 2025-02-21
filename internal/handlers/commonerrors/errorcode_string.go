// Code generated by "stringer -linecomment -type ErrorCode"; DO NOT EDIT.

package commonerrors

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[errUnset-0]
	_ = x[errInternalError-1]
	_ = x[ErrBadValue-2]
	_ = x[ErrFailedToParse-9]
	_ = x[ErrTypeMismatch-14]
	_ = x[ErrNamespaceNotFound-26]
	_ = x[ErrUnsuitableValueType-28]
	_ = x[ErrConflictingUpdateOperators-40]
	_ = x[ErrCursorNotFound-43]
	_ = x[ErrNamespaceExists-48]
	_ = x[ErrInvalidID-53]
	_ = x[ErrEmptyName-56]
	_ = x[ErrCommandNotFound-59]
	_ = x[ErrInvalidNamespace-73]
	_ = x[ErrOperationFailed-96]
	_ = x[ErrDocumentValidationFailure-121]
	_ = x[ErrNotImplemented-238]
	_ = x[ErrMechanismUnavailable-334]
	_ = x[ErrDuplicateKey-11000]
	_ = x[ErrStageGroupID-15948]
	_ = x[ErrMatchBadExpression-15959]
	_ = x[ErrSortBadExpression-15973]
	_ = x[ErrSortBadValue-15974]
	_ = x[ErrSortBadOrder-15975]
	_ = x[ErrSortMissingKey-15976]
	_ = x[ErrInvalidArg-28667]
	_ = x[ErrSliceFirstArg-28724]
	_ = x[ErrProjectionInEx-31253]
	_ = x[ErrProjectionExIn-31254]
	_ = x[ErrStageCountNonString-40156]
	_ = x[ErrStageCountNonEmptyString-40157]
	_ = x[ErrStageCountBadPrefix-40158]
	_ = x[ErrStageCountBadValue-40160]
	_ = x[ErrStageInvalid-40323]
	_ = x[ErrEmptyFieldPath-40352]
	_ = x[ErrMissingField-40414]
	_ = x[ErrFailedToParseInput-40415]
	_ = x[ErrFreeMonitoringDisabled-50840]
	_ = x[ErrBatchSizeNegative-51024]
	_ = x[ErrRegexOptions-51075]
	_ = x[ErrRegexMissingParen-51091]
	_ = x[ErrBadRegexOption-51108]
}

const _ErrorCode_name = "UnsetInternalErrorBadValueFailedToParseTypeMismatchNamespaceNotFoundUnsuitableValueTypeConflictingUpdateOperatorsCursorNotFoundNamespaceExistsInvalidIDEmptyNameCommandNotFoundInvalidNamespaceOperationFailedDocumentValidationFailureNotImplementedMechanismUnavailableLocation11000Location15948Location15959Location15973Location15974Location15975Location15976Location28667Location28724Location31253Location31254Location40156Location40157Location40158Location40160Location40323Location40352Location40414Location40415Location50840Location51024Location51075Location51091Location51108"

var _ErrorCode_map = map[ErrorCode]string{
	0:     _ErrorCode_name[0:5],
	1:     _ErrorCode_name[5:18],
	2:     _ErrorCode_name[18:26],
	9:     _ErrorCode_name[26:39],
	14:    _ErrorCode_name[39:51],
	26:    _ErrorCode_name[51:68],
	28:    _ErrorCode_name[68:87],
	40:    _ErrorCode_name[87:113],
	43:    _ErrorCode_name[113:127],
	48:    _ErrorCode_name[127:142],
	53:    _ErrorCode_name[142:151],
	56:    _ErrorCode_name[151:160],
	59:    _ErrorCode_name[160:175],
	73:    _ErrorCode_name[175:191],
	96:    _ErrorCode_name[191:206],
	121:   _ErrorCode_name[206:231],
	238:   _ErrorCode_name[231:245],
	334:   _ErrorCode_name[245:265],
	11000: _ErrorCode_name[265:278],
	15948: _ErrorCode_name[278:291],
	15959: _ErrorCode_name[291:304],
	15973: _ErrorCode_name[304:317],
	15974: _ErrorCode_name[317:330],
	15975: _ErrorCode_name[330:343],
	15976: _ErrorCode_name[343:356],
	28667: _ErrorCode_name[356:369],
	28724: _ErrorCode_name[369:382],
	31253: _ErrorCode_name[382:395],
	31254: _ErrorCode_name[395:408],
	40156: _ErrorCode_name[408:421],
	40157: _ErrorCode_name[421:434],
	40158: _ErrorCode_name[434:447],
	40160: _ErrorCode_name[447:460],
	40323: _ErrorCode_name[460:473],
	40352: _ErrorCode_name[473:486],
	40414: _ErrorCode_name[486:499],
	40415: _ErrorCode_name[499:512],
	50840: _ErrorCode_name[512:525],
	51024: _ErrorCode_name[525:538],
	51075: _ErrorCode_name[538:551],
	51091: _ErrorCode_name[551:564],
	51108: _ErrorCode_name[564:577],
}

func (i ErrorCode) String() string {
	if str, ok := _ErrorCode_map[i]; ok {
		return str
	}
	return "ErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
}
