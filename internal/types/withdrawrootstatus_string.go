// Code generated by "stringer -type WithdrawRootStatus"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[WithdrawRootStatusTypeUnknown-0]
	_ = x[WithdrawRootStatusTypeValid-1]
}

const _WithdrawRootStatus_name = "WithdrawRootStatusTypeUnknownWithdrawRootStatusTypeValid"

var _WithdrawRootStatus_index = [...]uint8{0, 29, 56}

func (i WithdrawRootStatus) String() string {
	if i < 0 || i >= WithdrawRootStatus(len(_WithdrawRootStatus_index)-1) {
		return "WithdrawRootStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _WithdrawRootStatus_name[_WithdrawRootStatus_index[i]:_WithdrawRootStatus_index[i+1]]
}
