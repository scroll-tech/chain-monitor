package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/config"
	messagematch "github.com/scroll-tech/chain-monitor/internal/logic/message_match"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// FinalizeBatchCheckController Check if the upcoming batch to be submitted is valid
type FinalizeBatchCheckController struct {
	db *gorm.DB

	messageMatchLogic *messagematch.LogicMessageMatch

	gatewayBatchFinalizeCheckFailed   prometheus.Counter
	messengerBatchFinalizeCheckFailed prometheus.Counter
}

// NewFinalizeBatchCheckController create finalize batch controller instance
func NewFinalizeBatchCheckController(conf *config.Config, db *gorm.DB) *FinalizeBatchCheckController {
	return &FinalizeBatchCheckController{
		db:                db,
		messageMatchLogic: messagematch.NewMessageMatchLogic(conf, db),

		gatewayBatchFinalizeCheckFailed: promauto.With(prometheus.DefaultRegisterer).NewCounter(prometheus.CounterOpts{
			Name: "gateway_batch_finalized_failed_total",
			Help: "The total number of gateway batch finalized failed.",
		}),
		messengerBatchFinalizeCheckFailed: promauto.With(prometheus.DefaultRegisterer).NewCounter(prometheus.CounterOpts{
			Name: "messenger_batch_finalized_failed_total",
			Help: "The total number of messenger batch finalized failed.",
		}),
	}
}

// BatchStatus get the upcoming finalized batch status
func (f *FinalizeBatchCheckController) BatchStatus(ctx *gin.Context) {
	var finalizeBatchParam types.FinalizeBatchCheckParam
	err := ctx.ShouldBind(&finalizeBatchParam)
	if err != nil {
		log.Error("batch status failed", "error", err)
		types.RenderJSON(ctx, types.ErrParameterInvalidNo, err, nil)
		return
	}

	// because the l2CurrentMaxBlockNumber updated after the contract watcher loop, the block number less than l2CurrentMaxBlockNumber must
	// have insert to db (If block's event aligns with the event being observed). So, if rollup's query block number less than l2CurrentMaxBlockNumber,
	// chain_monitor have checked the status of these blocks.
	if finalizeBatchParam.StartBlockNumber > l2CurrentMaxBlockNumber.Load() || finalizeBatchParam.EndBlockNumber > l2CurrentMaxBlockNumber.Load() {
		log.Error("batch status failed for query number large than current l2 block number",
			"current l2 block number", l2CurrentMaxBlockNumber.Load(),
			"start number", finalizeBatchParam.StartBlockNumber,
			"end number", finalizeBatchParam.EndBlockNumber,
		)
		types.RenderJSON(ctx, types.ErrParameterInvalidNo, err, nil)
		return
	}

	gatewayCheck, messengerCheck := f.messageMatchLogic.GetBlocksStatus(ctx, finalizeBatchParam.StartBlockNumber, finalizeBatchParam.EndBlockNumber)
	if !gatewayCheck {
		f.gatewayBatchFinalizeCheckFailed.Inc()
	}

	if !messengerCheck {
		f.messengerBatchFinalizeCheckFailed.Inc()
	}

	types.RenderJSON(ctx, types.Success, nil, gatewayCheck && messengerCheck)
}
