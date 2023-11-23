// Code generated by "stringer -type EventCategory"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[EventCategoryUnknown-0]
	_ = x[ERC20EventCategory-1]
	_ = x[ERC721EventCategory-2]
	_ = x[ERC1155EventCategory-3]
	_ = x[MessengerEventCategory-4]
}

const _EventCategory_name = "EventCategoryUnknownERC20EventCategoryERC721EventCategoryERC1155EventCategoryMessengerEventCategory"

var _EventCategory_index = [...]uint8{0, 20, 38, 57, 77, 99}

func (i EventCategory) String() string {
	if i < 0 || i >= EventCategory(len(_EventCategory_index)-1) {
		return "EventCategory(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EventCategory_name[_EventCategory_index[i]:_EventCategory_index[i+1]]
}
