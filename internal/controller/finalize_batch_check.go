package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/config"
	messagematch "github.com/scroll-tech/chain-monitor/internal/logic/message_match"
	"github.com/scroll-tech/chain-monitor/internal/types"
)

// FinalizeBatchCheckController Check if the upcoming batch to be submitted is valid
type FinalizeBatchCheckController struct {
	db *gorm.DB

	messageMatchLogic *messagematch.LogicMessageMatch
}

// NewFinalizeBatchCheckController create finalize batch controller instance
func NewFinalizeBatchCheckController(conf *config.Config, db *gorm.DB) *FinalizeBatchCheckController {
	return &FinalizeBatchCheckController{
		db:                db,
		messageMatchLogic: messagematch.NewMessageMatchLogic(conf, db),
	}
}

// BatchStatus get the upcoming finalized batch status
func (f *FinalizeBatchCheckController) BatchStatus(ctx *gin.Context) {
	var finalizeBatchParam types.FinalizeBatchCheckParam
	err := ctx.ShouldBind(&finalizeBatchParam)
	if err != nil {
		types.RenderJSON(ctx, types.ErrParameterInvalidNo, err, nil)
		return
	}

	ret := f.messageMatchLogic.GetBlocksStatus(ctx, finalizeBatchParam.StartBlockNumber, finalizeBatchParam.EndBlockNumber)
	types.RenderJSON(ctx, types.Success, nil, ret)
}
