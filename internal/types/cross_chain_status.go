package types

//go:generate stringer -type CrossChainStatusType

// CrossChainStatusType represents the status of a cross-chain transaction.
type CrossChainStatusType int

const (
	// CrossChainStatusTypeUnknown represents an unknown or undefined cross-chain status.
	CrossChainStatusTypeUnknown CrossChainStatusType = iota
	// CrossChainStatusTypeInvalid represents an invalid cross-chain transaction.
	CrossChainStatusTypeInvalid
	// CrossChainStatusTypeValid represents a valid cross-chain transaction.
	CrossChainStatusTypeValid
)
