// Code generated by "stringer -type BlockStatus"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BlockStatusTypeInvalid-0]
	_ = x[BlockStatusTypeValid-1]
}

const _BlockStatus_name = "BlockStatusTypeInvalidBlockStatusTypeValid"

var _BlockStatus_index = [...]uint8{0, 22, 42}

func (i BlockStatus) String() string {
	if i < 0 || i >= BlockStatus(len(_BlockStatus_index)-1) {
		return "BlockStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BlockStatus_name[_BlockStatus_index[i]:_BlockStatus_index[i+1]]
}
