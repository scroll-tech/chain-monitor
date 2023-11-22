package types

//go:generate stringer -type MismatchType

// MismatchType represents the type of mismatch in cross-chain transactions.
type MismatchType int

const (
	// MismatchTypeUnknown represents an unknown or undefined mismatch.
	MismatchTypeUnknown MismatchType = iota
	// MismatchTypeValid represents a no mismatch condition, meaning the transaction is valid.
	MismatchTypeValid
	// MismatchTypeL1EventNotMatch represents a mismatch where the Layer 1 event does not match the Layer 2 event.
	MismatchTypeL1EventNotMatch
	// MismatchTypeL2EventNotMatch represents a mismatch where the Layer 2 event does not match the Layer 1 event.
	MismatchTypeL2EventNotMatch
	// MismatchTypeL1MountNotMatch represents a mismatch where the layer1 amount does not match the Layer 2
	MismatchTypeL1MountNotMatch
	// MismatchTypeL2MountNotMatch represents a mismatch where the layer2 amount does not match the Layer 1
	MismatchTypeL2MountNotMatch
)
