package types

// LayerType represents the type of the chain, which can be Layer 1 or Layer 2.
type LayerType int

const (
	// LayerUnknown represents an unknown or undefined chain type.
	LayerUnknown LayerType = iota
	// Layer1 represents a Layer 1 chain type.
	Layer1
	// Layer2 represents a Layer 2 chain type.
	Layer2
)
