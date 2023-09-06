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

// ConfirmWithdrawRoot returns the batch status based on the requested block number.
func (m *ChainConfirm) ConfirmWithdrawRoot(ctx *gin.Context) {
	var req types.QueryByNumber
	err := ctx.ShouldBind(&req)
	if err != nil {
		types.RenderJSON(ctx, types.ErrParameterInvalidNo, err, nil)
		return
	}

	cfm, err := orm.GetConfirmMsgByNumber(m.db, req.Number)
	if err != nil {
		types.RenderJSON(ctx, types.ErrConfirmWithdrawRootByNumber, err, nil)
		return
	}

	types.RenderJSON(
		ctx,
		types.Success,
		nil,
		cfm.WithdrawRootStatus && cfm.DepositStatus && cfm.WithdrawStatus,
	)
}
