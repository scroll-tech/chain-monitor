package types

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
)

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

// ETHBalanceStatus represents the status of an Ethereum balance check.
type ETHBalanceStatus int

const (
	// ETHBalanceStatusTypeUnknown represents an unknown or undefined balance status.
	ETHBalanceStatusTypeUnknown ETHBalanceStatus = iota
	// ETHBalanceStatusTypeInvalid represents an invalid balance.
	ETHBalanceStatusTypeInvalid
	// ETHBalanceStatusTypeValid represents a valid balance.
	ETHBalanceStatusTypeValid
)
