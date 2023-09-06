package orm

import (
	"github.com/scroll-tech/go-ethereum/common"
	"gorm.io/gorm"
)

// ChainConfirm represents the confirmation status of various events in the blockchain.
// It keeps track of the confirmation status for deposits, withdrawals, and overall confirmations for a given block number.
type ChainConfirm struct {
	Number             uint64      `gorm:"primaryKey"`
	WithdrawRoot       common.Hash `gorm:"-"`
	WithdrawRootStatus bool        `gorm:"type: withdraw_root_status"`

	DepositStatus   bool `gorm:"type: deposit_status"`
	DepositConfirm  bool `gorm:"type: deposit_confirm"`
	WithdrawStatus  bool `gorm:"type: withdraw_status"`
	WithdrawConfirm bool `gorm:"type: withdraw_confirm"`
}

// GetLatestDepositConfirmedNumber retrieves the latest block number with confirmation status set to true.
func GetLatestDepositConfirmedNumber(db *gorm.DB) (uint64, error) {
	var monitor ChainConfirm
	err := db.Where("deposit_confirm = true").Last(&monitor).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return 0, err
	}
	return monitor.Number, nil
}

// GetLatestWithdrawConfirmNumber retrieves the latest block number with confirmation status set to true.
func GetLatestWithdrawConfirmNumber(db *gorm.DB) (uint64, error) {
	var monitor ChainConfirm
	err := db.Where("withdraw_confirm = true").Last(&monitor).Error
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
