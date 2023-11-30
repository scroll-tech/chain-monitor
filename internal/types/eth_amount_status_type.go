package types

//go:generate stringer -type ETHAmountStatus

// ETHAmountStatus represents if the eth amount of this message is set.
type ETHAmountStatus int

const (
	// BlockStatusTypeInvalid represents the eth amount is unset.
	ETHAmountStatusTypeUnset ETHAmountStatus = iota
	// BlockStatusTypeValid represents the eth amount is set.
	ETHAmountStatusTypeSet
)
