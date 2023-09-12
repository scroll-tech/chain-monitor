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
	var req types.QueryByBatchNumber
	err := ctx.ShouldBind(&req)
	if err != nil {
		types.RenderJSON(ctx, types.ErrParameterInvalidNo, err, nil)
		return
	}

	l2FailedConfirms, err := orm.GetL2ConfirmMsgByNumber(m.db, req.StartNumber, req.EndNumber)
	if err != nil {
		types.RenderJSON(ctx, types.ErrConfirmWithdrawRootByNumber, err, nil)
		return
	}
	types.RenderJSON(ctx, types.Success, nil, len(l2FailedConfirms) == 0)
}
