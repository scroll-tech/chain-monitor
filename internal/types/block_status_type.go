package types

//go:generate stringer -type BlockStatus

// BlockStatus represents the status of a block in the blockchain.
type BlockStatus int

const (
	// BlockStatusTypeUnknown represents an unknown or undefined block status.
	BlockStatusTypeUnknown BlockStatus = iota
	// BlockStatusTypeInvalid represents an invalid block.
	BlockStatusTypeInvalid
	// BlockStatusTypeValid represents a valid block.
	BlockStatusTypeValid
)
