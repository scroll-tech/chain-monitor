package orm

import (
	"github.com/scroll-tech/go-ethereum/common"
	"gorm.io/gorm"
)

type ChainConfirm struct {
	Number         uint64      `gorm:"primaryKey"`
	WithdrawRoot   common.Hash `gorm:"-"`
	WithdrawStatus bool
	DepositStatus  bool
	Confirm        bool
}

func GetLatestConfirmedNumber(db *gorm.DB) (uint64, error) {
	var monitor ChainConfirm
	err := db.Where("confirm = true").Last(&monitor).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return 0, err
	}
	return monitor.Number, nil
}

func GetConfirmMsgByNumber(db *gorm.DB, number uint64) (*ChainConfirm, error) {
	var confirmBatch ChainConfirm
	res := db.Where("number = ?", number).First(&confirmBatch)
	return &confirmBatch, res.Error
}
