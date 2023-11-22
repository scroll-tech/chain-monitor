package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/rpc"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/config"
	crosschain "github.com/scroll-tech/chain-monitor/internal/logic/cross_chain"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// CrossChainController is a struct that contains a reference to the Logic object.
type CrossChainController struct {
	crossChainLogic *crosschain.Logic
	stopTimeoutChan chan struct{}
}

// NewCrossChainController is a constructor function that creates a new CrossChainController object.
func NewCrossChainController(cfg *config.Config, db *gorm.DB, l1Client, l2Client *rpc.Client) *CrossChainController {
	l1MessengerAddr := cfg.L1Config.L1Contracts.ScrollMessenger
	l2MessengerAddr := cfg.L2Config.L2Contracts.ScrollMessenger
	return &CrossChainController{
		stopTimeoutChan: make(chan struct{}),
		crossChainLogic: crosschain.NewLogic(db, l1Client, l2Client, l1MessengerAddr, l2MessengerAddr),
	}
}

// Watch is a method that triggers the proposer methods for Layer 1 and Layer 2, as well as the
// eth balance checker methods for both layers.
func (c *CrossChainController) Watch(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			nerr := fmt.Errorf("CrossChainController watch panic error: %v", err)
			log.Warn(nerr.Error())
		}
	}()

	log.Info("cross chain controller start successful")

	ticker := time.NewTicker(time.Millisecond * 500)
	for {
		select {
		case <-ticker.C:
			go c.l1Watcher(ctx)
			go c.l2Watcher(ctx)
		case <-ctx.Done():
			if ctx.Err() != nil {
				log.Error("CrossChainController proposer canceled with error", "error", ctx.Err())
			}
			return
		case <-c.stopTimeoutChan:
			log.Info("CrossChainController proposer the run loop exit")
			return
		}
	}
}

// Stop all the cross chain controller
func (c *CrossChainController) Stop() {
	c.stopTimeoutChan <- struct{}{}
}

func (c *CrossChainController) l1Watcher(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			nerr := fmt.Errorf("l1Proposer panic error: %v", err)
			log.Warn(nerr.Error())
		}
	}()
	c.crossChainLogic.CheckCrossChainGatewayMessage(ctx, types.Layer1)
	c.crossChainLogic.CheckETHBalance(ctx, types.Layer1)
}

func (c *CrossChainController) l2Watcher(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			nerr := fmt.Errorf("l2Proposer panic error: %v", err)
			log.Warn(nerr.Error())
		}
	}()
	c.crossChainLogic.CheckCrossChainGatewayMessage(ctx, types.Layer2)
	c.crossChainLogic.CheckETHBalance(ctx, types.Layer2)
}
