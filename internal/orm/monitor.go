package orm

import (
	"github.com/scroll-tech/go-ethereum/common"
	"gorm.io/gorm"
)

// ChainConfirm represents the confirmation status of various events in the blockchain.
// It keeps track of the confirmation status for deposits, withdrawals, and overall confirmations for a given block number.
type ChainConfirm struct {
	Number         uint64      `gorm:"primaryKey"`
	WithdrawRoot   common.Hash `gorm:"-"`
	WithdrawStatus bool
	DepositStatus  bool
	Confirm        bool
}

// GetLatestConfirmedNumber retrieves the latest block number with confirmation status set to true.
func GetLatestConfirmedNumber(db *gorm.DB) (uint64, error) {
	var monitor ChainConfirm
	err := db.Where("confirm = true").Last(&monitor).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return 0, err
	}
	return monitor.Number, nil
}

// GetConfirmMsgByNumber fetches the ChainConfirm message for a given block number.
func GetConfirmMsgByNumber(db *gorm.DB, number uint64) (*ChainConfirm, error) {
	var confirmBatch ChainConfirm
	res := db.Where("number = ?", number).First(&confirmBatch)
	return &confirmBatch, res.Error
}
