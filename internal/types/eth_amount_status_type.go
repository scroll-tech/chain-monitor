package types

//go:generate stringer -type ETHAmountStatus

// ETHAmountStatus represents if the eth amount of this message is set.
type ETHAmountStatus int

const (
	// ETHAmountStatusTypeUnset represents the eth amount is unset.
	ETHAmountStatusTypeUnset ETHAmountStatus = iota
	// ETHAmountStatusTypeSet represents the eth amount is set.
	ETHAmountStatusTypeSet
)
