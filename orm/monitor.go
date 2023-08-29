package orm

import (
	"github.com/scroll-tech/go-ethereum/common"
	"gorm.io/gorm"
)

type BatchConfirm struct {
	BatchIndex         uint64 `gorm:"primaryKey"`
	DepositStatus      bool
	WithdrawRootStatus bool
}

func GetConfirmedBatchIndex(db *gorm.DB) (uint64, error) {
	var monitor BatchConfirm
	err := db.Last(&monitor).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return 0, err
	}
	return monitor.BatchIndex, nil
}

func GetConfirmBatchByIndex(db *gorm.DB, batchIndex uint64) (*BatchConfirm, error) {
	var confirmBatch BatchConfirm
	res := db.Where("batch_index = ?", batchIndex).First(&confirmBatch)
	return &confirmBatch, res.Error
}

func GetConfirmBatchByHash(db *gorm.DB, batchHash common.Hash) (*BatchConfirm, error) {
	var confirmBatch BatchConfirm
	res := db.Where("batch_hash = ?", batchHash.String()).First(&confirmBatch)
	return &confirmBatch, res.Error
}
