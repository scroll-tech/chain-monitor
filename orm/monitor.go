package orm

import (
	"github.com/scroll-tech/go-ethereum/common"
	"gorm.io/gorm"
)

type ChainConfirm struct {
	Number             uint64 `gorm:"primaryKey"`
	ActualWithdrawRoot string
	Confirm            bool
	DepositStatus      bool
	WithdrawStatus     bool
}

func GetLatestConfirmedNumber(db *gorm.DB) (uint64, error) {
	var monitor ChainConfirm
	err := db.Where("confirm = true").Last(&monitor).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return 0, err
	}
	return monitor.Number, nil
}

func GetConfirmBatchByIndex(db *gorm.DB, batchIndex uint64) (*ChainConfirm, error) {
	var confirmBatch ChainConfirm
	res := db.Where("batch_index = ?", batchIndex).First(&confirmBatch)
	return &confirmBatch, res.Error
}

func GetConfirmBatchByHash(db *gorm.DB, batchHash common.Hash) (*ChainConfirm, error) {
	var confirmBatch ChainConfirm
	res := db.Where("batch_hash = ?", batchHash.String()).First(&confirmBatch)
	return &confirmBatch, res.Error
}
