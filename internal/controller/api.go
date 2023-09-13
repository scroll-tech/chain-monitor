package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"chain-monitor/internal/orm"
	"chain-monitor/internal/types"
)

// ChainConfirm controller is responsible for handling confirmation operations on the chain.
type ChainConfirm struct {
	db *gorm.DB
}

// NewMetricsController creates and returns a new instance of the ChainConfirm controller.
func NewMetricsController(db *gorm.DB) *ChainConfirm {
	return &ChainConfirm{db: db}
}

func (m *ChainConfirm) confirmBlocksStatus(start, end uint64) (bool, error) {
	l2FailedConfirms, err := orm.GetL2ConfirmMsgByNumber(m.db, start, end)
	if err != nil {
		return false, err
	}
	return len(l2FailedConfirms) == 0, nil
}

// ConfirmBlocksStatus returns the blocks status based on the requested start_number and end_number.
func (m *ChainConfirm) ConfirmBlocksStatus(ctx *gin.Context) {
	var req types.QueryByBatchNumber
	err := ctx.ShouldBind(&req)
	if err != nil {
		types.RenderJSON(ctx, types.ErrParameterInvalidNo, err, nil)
		return
	}

	isOK, err := m.confirmBlocksStatus(req.StartNumber, req.EndNumber)
	if err != nil {
		types.RenderJSON(ctx, types.ErrBlocksStatus, err, nil)
		return
	}
	types.RenderJSON(ctx, types.Success, nil, isOK)
}

// ConfirmBatchStatus returns the batch status based on the requested batch index.
func (m *ChainConfirm) ConfirmBatchStatus(ctx *gin.Context) {
	var req types.QueryByBatchIndex
	err := ctx.ShouldBind(&req)
	if err != nil {
		types.RenderJSON(ctx, types.ErrParameterInvalidNo, err, nil)
		return
	}
	scrollBatch, err := orm.GetBatchByIndex(m.db, req.BatchIndex)
	if err != nil {
		types.RenderJSON(ctx, types.ErrBatchStatus, err, nil)
		return
	}
	isOK, err := m.confirmBlocksStatus(scrollBatch.L2StartNumber, scrollBatch.L2EndNumber)
	if err != nil {
		types.RenderJSON(ctx, types.ErrBatchStatus, err, nil)
		return
	}
	types.RenderJSON(ctx, types.Success, nil, isOK)
}
