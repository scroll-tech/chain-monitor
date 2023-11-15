package controller

import (
	"context"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"gorm.io/gorm"

	crosschain "github.com/scroll-tech/chain-monitor/internal/logic/cross_chain"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// CrossChainController is a struct that contains a reference to the Logic object.
type CrossChainController struct {
	crossChainLogic *crosschain.Logic
}

// NewCrossChainController is a constructor function that creates a new CrossChainController object.
func NewCrossChainController(db *gorm.DB, client *ethclient.Client, l1MessengerAddr, l2MessengerAddr common.Address) *CrossChainController {
	return &CrossChainController{
		crossChainLogic: crosschain.NewLogic(db, client, l1MessengerAddr, l2MessengerAddr),
	}
}

// Proposer is a method that triggers the proposer methods for Layer 1 and Layer 2, as well as the
// eth balance checker methods for both layers.
func (c *CrossChainController) Proposer(ctx context.Context) {
	c.l1Proposer(ctx)
	c.l2Proposer(ctx)
	c.l1ETHBalanceChecker(ctx)
	c.l2ETHBalanceChecker(ctx)
}

func (c *CrossChainController) l1Proposer(ctx context.Context) {
	c.crossChainLogic.CheckCrossChainMessage(ctx, types.Layer1)
}

func (c *CrossChainController) l2Proposer(ctx context.Context) {
	c.crossChainLogic.CheckCrossChainMessage(ctx, types.Layer2)
}

func (c *CrossChainController) l1ETHBalanceChecker(ctx context.Context) {
	c.crossChainLogic.CheckETHBalance(ctx, types.Layer1)
}

func (c *CrossChainController) l2ETHBalanceChecker(ctx context.Context) {
	c.crossChainLogic.CheckETHBalance(ctx, types.Layer2)
}
