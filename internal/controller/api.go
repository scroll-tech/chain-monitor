package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"chain-monitor/internal/orm"
	"chain-monitor/internal/types"
)

type ChainConfirm struct {
	db *gorm.DB
}

func NewMetricsController(db *gorm.DB) *ChainConfirm {
	return &ChainConfirm{db: db}
}

// ConfirmWithdrawRoot return batch status
func (m *ChainConfirm) ConfirmWithdrawRoot(ctx *gin.Context) {
	var req types.QueryByNumber
	err := ctx.ShouldBind(&req)
	if err != nil {
		types.RenderJSON(ctx, types.ErrParameterInvalidNo, err, nil)
		return
	}

	confirmBlock, err := orm.GetConfirmMsgByNumber(m.db, req.Number)
	if err != nil {
		types.RenderJSON(ctx, types.ErrConfirmWithdrawRootByNumber, err, nil)
		return
	}

	types.RenderJSON(ctx, types.Success, nil, confirmBlock.WithdrawStatus && confirmBlock.DepositStatus)
}
