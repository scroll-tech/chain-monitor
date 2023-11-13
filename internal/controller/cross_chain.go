package controller

import (
	"context"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/logic/cross_chain"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

type CrossChainController struct {
	crossChainLogic *cross_chain.CrossChainLogic
}

func NewCrossChainController(db *gorm.DB, client *ethclient.Client, l1MessengerAddr, l2MessengerAddr common.Address) *CrossChainController {
	return &CrossChainController{
		crossChainLogic: cross_chain.NewCrossChainLogic(db, client, l1MessengerAddr, l2MessengerAddr),
	}
}

func (c *CrossChainController) Proposer() {
}

func (c *CrossChainController) l1Proposer(ctx context.Context, start, end uint64) {
	c.crossChainLogic.Fetcher(ctx, types.Layer1, start, end)
}

func (c *CrossChainController) l2Proposer(ctx context.Context, start, end uint64) {
	c.crossChainLogic.Fetcher(ctx, types.Layer2, start, end)
}

func (c *CrossChainController) l1ETHBalanceChecker(ctx context.Context) {
	c.crossChainLogic.CheckETHBalance(ctx, types.Layer1)
}

func (c *CrossChainController) l2ETHBalanceChecker(ctx context.Context) {
	c.crossChainLogic.CheckETHBalance(ctx, types.Layer2)
}
