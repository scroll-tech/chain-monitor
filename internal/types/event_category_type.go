package types

//go:generate stringer -type EventCategory

// EventCategory represents the category of events.
type EventCategory int

const (
	// EventCategoryUnknown represents an unknown or undefined event category.
	EventCategoryUnknown EventCategory = iota
	// ERC20EventCategory represents the ERC20 gateway events.
	ERC20EventCategory
	// ERC721EventCategory represents the ERC721 gateway events.
	ERC721EventCategory
	// ERC1155EventCategory represents the ERC1155 gateway events.
	ERC1155EventCategory
	// MessengerEventCategory represents the messenger events.
	MessengerEventCategory
)
