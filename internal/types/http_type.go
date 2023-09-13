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
	// ErrBlocksStatus failed to get blocks status
	ErrBlocksStatus = 40005
	// ErrBatchStatus failed to get batch status
	ErrBatchStatus = 40006
)

// QueryByBatchIndexOrHashRequest the request parameter of batch index api
type QueryByBatchIndexOrHashRequest struct {
	// BatchIndex can not be 0, because we dont decode the genesis block
	BatchIndex uint64      `form:"batch_index" binding:"required"`
	BatchHash  common.Hash `json:"batch_hash" binding:"required"`
}

// QueryByBatchNumber request withdraw root by block number.
type QueryByBatchNumber struct {
	// block number
	StartNumber uint64 `form:"start_number" binding:"required"`
	EndNumber   uint64 `form:"end_number" binding:"required"`
}

// QueryByBatchIndex request withdraw root by block number.
type QueryByBatchIndex struct {
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
