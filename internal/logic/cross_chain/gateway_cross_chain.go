package crosschain

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/logic/slack"
	"github.com/scroll-tech/chain-monitor/internal/orm"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// LogicGatewayCrossChain is a struct for checking the l1/l2 event match from the database.
// FinalizeWithdraw value corresponds to the withdrawal value.
// FinalizeDeposit value corresponds to the deposit value.
// This is because not every deposit/withdrawal event in the system will have a finalize event,
// as users have the ability to refund deposits independently.
type LogicGatewayCrossChain struct {
	db                          *gorm.DB
	gatewayMessageOrm           *orm.GatewayMessageMatch
	checker                     *GatewayCrossEventMatcher
	crossChainGatewayCheckTotal *prometheus.CounterVec
}

// NewLogicGatewayCrossChain is a constructor for Logic.
func NewLogicGatewayCrossChain(db *gorm.DB) *LogicGatewayCrossChain {
	return &LogicGatewayCrossChain{
		db:                db,
		checker:           NewGatewayCrossEventMatcher(),
		gatewayMessageOrm: orm.NewGatewayMessageMatch(db),

		crossChainGatewayCheckTotal: promauto.With(prometheus.DefaultRegisterer).NewCounterVec(prometheus.CounterOpts{
			Name: "cross_chain_checked_gateway_event_check_total",
			Help: "the total number of cross chain gateway checked",
		}, []string{"layer"}),
	}
}

// CheckCrossChainGatewayMessage is a method for checking cross-chain messages.
func (c *LogicGatewayCrossChain) CheckCrossChainGatewayMessage(ctx context.Context, layerType types.LayerType) {
	messages, err := c.gatewayMessageOrm.GetUncheckedAndDoubleLayerValidGatewayMessageMatches(ctx, layerType, 1000)
	if err != nil {
		log.Error("CheckCrossChainGatewayMessage.GetUncheckedAndDoubleLayerValidGatewayMessageMatches failed", "error", err)
		return
	}

	if len(messages) == 0 {
		time.Sleep(time.Second)
		return
	}

	log.Info("checking cross chain gateway messages", "layer", layerType.String(), "number of messages", len(messages))

	var messageMatchIds []int64
	for _, message := range messages {
		c.crossChainGatewayCheckTotal.WithLabelValues(layerType.String()).Inc()
		checkResult := c.checker.GatewayCrossChainCheck(layerType, message)
		if checkResult == types.MismatchTypeValid {
			messageMatchIds = append(messageMatchIds, message.ID)
			continue
		}
		slack.Notify(slack.MrkDwnGatewayCrossChainMessage(message, checkResult))
	}

	if err = c.gatewayMessageOrm.UpdateCrossChainStatus(ctx, messageMatchIds, layerType, types.CrossChainStatusTypeValid); err != nil {
		log.Error("Logic.CheckCrossChainMessage UpdateCrossChainStatus failed", "error", err)
		return
	}
}
