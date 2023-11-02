package controller

type CrossChainController struct {
}

func NewCrossChainController() *ContractController {
	return &ContractController{}
}

func (c *CrossChainController) Proposer() {
}

func (c *CrossChainController) l1Proposer() {
	// 找到合适的 l1 number
}

func (c *CrossChainController) l2Proposer() {

}
