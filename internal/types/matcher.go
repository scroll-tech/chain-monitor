package types

type MismatchType int

const (
	MismatchTypeUnknown MismatchType = iota
	MismatchTypeAmount
	MismatchTypeL2EventNotExist
	MismatchTypeL1EventNotExist
	MismatchTypeCrossChainTypeNotMatch
	MismatchTypeCrossChainValueNotMatch
)
