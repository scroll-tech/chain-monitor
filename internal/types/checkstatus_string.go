// Code generated by "stringer -type CheckStatus"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CheckStatusUnchecked-0]
	_ = x[CheckStatusChecked-1]
}

const _CheckStatus_name = "CheckStatusUncheckedCheckStatusChecked"

var _CheckStatus_index = [...]uint8{0, 20, 38}

func (i CheckStatus) String() string {
	if i < 0 || i >= CheckStatus(len(_CheckStatus_index)-1) {
		return "CheckStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CheckStatus_name[_CheckStatus_index[i]:_CheckStatus_index[i+1]]
}
