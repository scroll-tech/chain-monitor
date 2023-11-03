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

type CheckStatusType int

const (
	CheckStatusUnknown CheckStatusType = iota
	CheckStatusUnchecked
	CheckStatusChecked
)

type CrossChainStatusType int

const (
	CrossChainStatusTypeUnknown CrossChainStatusType = iota
	CrossChainStatusTypeInvalid
	CrossChainStatusTypeValid
)

type GatewayStatusType int

const (
	GatewayStatusTypeUnknown GatewayStatusType = iota
	GatewayStatusTypeInvalid
	GatewayStatusTypeValid
)
