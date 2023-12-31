// Code generated by "stringer -type MismatchType"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MismatchTypeUnknown-0]
	_ = x[MismatchTypeValid-1]
	_ = x[MismatchTypeL1EventNotMatch-2]
	_ = x[MismatchTypeL2EventNotMatch-3]
	_ = x[MismatchTypeL1AmountNotMatch-4]
	_ = x[MismatchTypeL2AmountNotMatch-5]
}

const _MismatchType_name = "MismatchTypeUnknownMismatchTypeValidMismatchTypeL1EventNotMatchMismatchTypeL2EventNotMatchMismatchTypeL1AmountNotMatchMismatchTypeL2AmountNotMatch"

var _MismatchType_index = [...]uint8{0, 19, 36, 63, 90, 118, 146}

func (i MismatchType) String() string {
	if i < 0 || i >= MismatchType(len(_MismatchType_index)-1) {
		return "MismatchType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MismatchType_name[_MismatchType_index[i]:_MismatchType_index[i+1]]
}
