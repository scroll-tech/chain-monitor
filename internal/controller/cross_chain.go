package controller

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/config"
	crosschain "github.com/scroll-tech/chain-monitor/internal/logic/cross_chain"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// CrossChainController is a struct that contains a reference to the Logic object.
type CrossChainController struct {
	gatewayCrossChainLogic   *crosschain.LogicGatewayCrossChain
	messengerCrossChainLogic *crosschain.LogicMessengerCrossChain

	stopL1CrossChainChan chan struct{}
	stopL2CrossChainChan chan struct{}

	crossChainControllerRunningTotal *prometheus.CounterVec
}

// NewCrossChainController is a constructor function that creates a new CrossChainController object.
func NewCrossChainController(cfg *config.Config, db *gorm.DB, l1Client, l2Client *ethclient.Client) *CrossChainController {
	l1MessengerAddr := cfg.L1Config.L1Contracts.ScrollMessenger
	l2MessengerAddr := cfg.L2Config.L2Contracts.ScrollMessenger
	return &CrossChainController{
		stopL1CrossChainChan:     make(chan struct{}),
		stopL2CrossChainChan:     make(chan struct{}),
		gatewayCrossChainLogic:   crosschain.NewLogicGatewayCrossChain(db),
		messengerCrossChainLogic: crosschain.NewLogicMessengerCrossChain(db, l1Client, l2Client, l1MessengerAddr, l2MessengerAddr, cfg.L1Config.StartMessengerBalance),
		crossChainControllerRunningTotal: promauto.With(prometheus.DefaultRegisterer).NewCounterVec(prometheus.CounterOpts{
			Name: "cross_chain_check_controller_running_total",
			Help: "The total number of cross chain controllers running.",
		}, []string{"layer"}),
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
	c.stopL1CrossChainChan <- struct{}{}
	c.stopL2CrossChainChan <- struct{}{}
}

func (c *CrossChainController) watcherStart(ctx context.Context, layer types.LayerType) {
	log.Info("cross chain controller start successful", "layer", layer.String())

	tick := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ctx.Done():
			tick.Stop()
			if ctx.Err() != nil {
				log.Error("CrossChainController watch canceled with error", "layer", layer.String(), "error", ctx.Err())
			}
			return
		case <-c.stopL1CrossChainChan:
			tick.Stop()
			log.Info("CrossChainController l1 the run loop exit", "layer", layer.String())
			return
		case <-c.stopL2CrossChainChan:
			tick.Stop()
			log.Info("CrossChainController l2 the run loop exit", "layer", layer.String())
			return
		case <-tick.C:
			c.crossChainControllerRunningTotal.WithLabelValues(layer.String()).Inc()
			c.gatewayCrossChainLogic.CheckCrossChainGatewayMessage(ctx, layer)
			c.messengerCrossChainLogic.CheckETHBalance(ctx, layer)
		}
	}
}
