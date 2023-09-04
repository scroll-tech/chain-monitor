package types

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scroll-tech/go-ethereum/common"
)

const (
	// Success shows OK.
	Success = 0
	// ErrParameterInvalidNo is invalid params
	ErrParameterInvalidNo = 40001
	// ErrConfirmWithdrawRootByNumber failed to get confirm batch
	ErrConfirmWithdrawRootByNumber = 40005
)

// QueryByBatchIndexOrHashRequest the request parameter of batch index api
type QueryByBatchIndexOrHashRequest struct {
	// BatchIndex can not be 0, because we dont decode the genesis block
	BatchIndex uint64      `form:"batch_index" binding:"required"`
	BatchHash  common.Hash `json:"batch_hash" binding:"required"`
}

// QueryByNumber request withdraw root by block number.
type QueryByNumber struct {
	// block number
	Number uint64 `form:"block_number" binding:"required"`
}

// Response the response schema
type Response struct {
	ErrCode int         `json:"errcode"`
	ErrMsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

// RenderJSON renders response with json
func RenderJSON(ctx *gin.Context, errCode int, err error, data interface{}) {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	renderData := Response{
		ErrCode: errCode,
		ErrMsg:  errMsg,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, renderData)
}
