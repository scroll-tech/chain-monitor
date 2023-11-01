package types

type ContractType int

const (
	ContractUnknown ContractType = iota
	ContractETH
	ContractWETH
)

type LayerType int

const (
	LayerUnknown LayerType = iota
	Layer1
	Layer2
)
