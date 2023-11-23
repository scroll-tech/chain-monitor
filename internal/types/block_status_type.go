package types

//go:generate stringer -type BlockStatus

// BlockStatus represents the status of a block in the blockchain.
type BlockStatus int

const (
	// BlockStatusTypeInvalid represents an invalid block.
	BlockStatusTypeInvalid BlockStatus = iota
	// BlockStatusTypeValid represents a valid block.
	BlockStatusTypeValid
)
