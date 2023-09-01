package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/scroll-tech/go-ethereum/common"
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
	var req types.QueryByBatchIndexOrHashRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		types.RenderJSON(ctx, types.ErrParameterInvalidNo, err, nil)
		return
	}

	var confirmBatch *orm.ChainConfirm
	if req.BatchHash != (common.Hash{}) {
		confirmBatch, err = orm.GetConfirmBatchByHash(m.db, req.BatchHash)
		if err != nil {
			types.RenderJSON(ctx, types.ErrConfirmWithdrawRootByBatchIndex, err, nil)
			return
		}
	} else if req.BatchIndex != 0 {
		confirmBatch, err = orm.GetConfirmBatchByIndex(m.db, req.BatchIndex)
		if err != nil {
			types.RenderJSON(ctx, types.ErrConfirmWithdrawRootByBatchHash, err, nil)
			return
		}
	} else {
		types.RenderJSON(ctx, types.ErrParameterInvalidNo, errors.New("invalid parameters"), nil)
		return
	}

	types.RenderJSON(ctx, types.Success, nil, confirmBatch.WithdrawStatus && confirmBatch.DepositStatus)
}
