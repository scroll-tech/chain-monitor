package types

//go:generate stringer -type ETHBalanceStatus

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
