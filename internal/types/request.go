package types

// FinalizeBatchCheckParam the param of batch status check
type FinalizeBatchCheckParam struct {
	BatchIndex       uint64 `form:"batch_index" json:"batch_index" binding:"required"`
	StartBlockNumber uint64 `form:"start_block_number" json:"start_block_number" binding:"required"`
	EndBlockNumber   uint64 `form:"end_block_number" json:"end_block_number" binding:"required"`
}
