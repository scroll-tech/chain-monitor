package types

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// Success shows OK.
	Success = 0
	// ErrParameterInvalidNo is invalid params
	ErrParameterInvalidNo  = 40001
	ErrConfirmBatchBalance = 40004
	// ErrConfirmWithdrawRootByBatchIndex failed to get confirm batch
	ErrConfirmWithdrawRootByBatchIndex = 40005
)

// QueryByBatchIndexRequest the request parameter of batch index api
type QueryByBatchIndexRequest struct {
	// BatchIndex can not be 0, because we dont decode the genesis block
	BatchIndex uint64 `form:"batch_index" binding:"required"`
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
