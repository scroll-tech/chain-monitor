package controller

import (
	"context"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/config"
	crosschain "github.com/scroll-tech/chain-monitor/internal/logic/cross_chain"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// CrossChainController is a struct that contains a reference to the Logic object.
type CrossChainController struct {
	crossChainLogic *crosschain.LogicCrossChain
	stopTimeoutChan chan struct{}
}

// NewCrossChainController is a constructor function that creates a new CrossChainController object.
func NewCrossChainController(cfg *config.Config, db *gorm.DB, l1Client, l2Client *ethclient.Client) *CrossChainController {
	l1MessengerAddr := cfg.L1Config.L1Contracts.ScrollMessenger
	l2MessengerAddr := cfg.L2Config.L2Contracts.ScrollMessenger
	return &CrossChainController{
		stopTimeoutChan: make(chan struct{}),
		crossChainLogic: crosschain.NewCrossChainLogic(db, l1Client, l2Client, l1MessengerAddr, l2MessengerAddr),
	}
}

// Watch is a method that triggers the proposer methods for Layer 1 and Layer 2, as well as the
// eth balance checker methods for both layers.
func (c *CrossChainController) Watch(ctx context.Context) {
	go c.watcherStart(ctx, types.Layer1)
	go c.watcherStart(ctx, types.Layer2)
}

// Stop all the cross chain controller
func (c *CrossChainController) Stop() {
	c.stopTimeoutChan <- struct{}{}
}

func (c *CrossChainController) watcherStart(ctx context.Context, layer types.LayerType) {
	log.Info("cross chain controller start successful", "layer", layer.String())
	//defer func() {
	//	if err := recover(); err != nil {
	//		nerr := fmt.Errorf("layer:%s watcher start panic error: %v", layer.String(), err)
	//		log.Warn(nerr.Error())
	//	}
	//}()

	for {
		select {
		case <-ctx.Done():
			if ctx.Err() != nil {
				log.Error("CrossChainController proposer canceled with error", "layer", layer.String(), "error", ctx.Err())
			}
			return
		case <-c.stopTimeoutChan:
			log.Info("CrossChainController proposer the run loop exit", "layer", layer.String())
			return
		default:
		}

		c.crossChainLogic.CheckCrossChainGatewayMessage(ctx, layer)
		c.crossChainLogic.CheckETHBalance(ctx, layer)
	}
}
