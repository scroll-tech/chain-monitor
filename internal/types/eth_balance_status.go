package types

//go:generate stringer -type ETHBalanceStatus

// ETHBalanceStatus represents the status of an Ethereum balance check.
type ETHBalanceStatus int

const (
	// ETHBalanceStatusTypeInvalid represents an invalid balance.
	ETHBalanceStatusTypeInvalid ETHBalanceStatus = iota
	// ETHBalanceStatusTypeValid represents a valid balance.
	ETHBalanceStatusTypeValid
)
