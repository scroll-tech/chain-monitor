package controller

import (
	"context"

	"github.com/scroll-tech/chain-monitor/internal/logic/cross_chain"
	"github.com/scroll-tech/chain-monitor/internal/types"
	"gorm.io/gorm"
)

type CrossChainController struct {
	crossChainLogic *cross_chain.CrossChainLogic
}

func NewCrossChainController(db *gorm.DB) *CrossChainController {
	return &CrossChainController{
		crossChainLogic: cross_chain.NewCrossChainLogic(db),
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
