package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"chain-monitor/internal/types"
	"chain-monitor/orm"
)

type ChainConfirm struct {
	db *gorm.DB
}

func NewMetricsController(db *gorm.DB) *ChainConfirm {
	return &ChainConfirm{db: db}
}

// ConfirmBatch return batch status
func (m *ChainConfirm) ConfirmBatch(ctx *gin.Context) {
	var req types.QueryByBatchIndexRequest
	if err := ctx.ShouldBind(&req); err != nil {
		types.RenderJSON(ctx, types.ErrParameterInvalidNo, err, nil)
		return
	}
	confirmBatch, err := orm.GetConfirmBatchByIndex(m.db, req.BatchIndex)
	if err != nil {
		types.RenderJSON(ctx, types.ErrConfirmWithdrawRootByBatchIndex, err, nil)
		return
	}
	types.RenderJSON(ctx, types.Success, nil, confirmBatch.WithdrawRootStatus && confirmBatch.EventStatus)
}
