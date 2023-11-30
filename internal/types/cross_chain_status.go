package types

//go:generate stringer -type CrossChainStatusType

// CrossChainStatusType represents the status of a cross-chain transaction.
type CrossChainStatusType int

const (
	// CrossChainStatusTypeInvalid represents an invalid cross-chain transaction.
	CrossChainStatusTypeInvalid CrossChainStatusType = iota
	// CrossChainStatusTypeValid represents a valid cross-chain transaction.
	CrossChainStatusTypeValid
)
