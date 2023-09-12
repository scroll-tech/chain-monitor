package orm

import (
	"github.com/scroll-tech/go-ethereum/common"
	"gorm.io/gorm"
)

// L1ChainConfirm represents the confirmation status of various events in the blockchain.
// It keeps track of the confirmation status for deposits, withdrawals, and overall confirmations for a given block number.
type L1ChainConfirm struct {
	Number         uint64 `gorm:"primaryKey"`
	WithdrawStatus bool   `gorm:"type: withdraw_status"`
	Confirm        bool   `gorm:"type: confirm"`
}

// L2ChainConfirm represents the confirmation status of various events in the blockchain.
// It keeps track of the confirmation status for deposits, withdrawals, and overall confirmations for a given block number.
type L2ChainConfirm struct {
	Number             uint64      `gorm:"primaryKey"`
	WithdrawRoot       common.Hash `gorm:"-"`
	WithdrawRootStatus bool        `gorm:"type: withdraw_status"`

	DepositStatus bool `gorm:"type: deposit_status"`
	Confirm       bool `gorm:"type: confirm"`
}

// GetL2DepositNumber retrieves the latest block number with confirmation status set to true.
func GetL2DepositNumber(db *gorm.DB) (uint64, error) {
	var monitor L2ChainConfirm
	err := db.Model(&L2ChainConfirm{}).Where("confirm = true").Last(&monitor).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return 0, err
	}
	return monitor.Number, nil
}

// GetL1WithdrawNumber retrieves the latest block number with confirmation status set to true.
func GetL1WithdrawNumber(db *gorm.DB) (uint64, error) {
	var monitor L1ChainConfirm
	err := db.Model(&L1ChainConfirm{}).Where("confirm = true").Last(&monitor).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return 0, err
	}
	return monitor.Number, nil
}

// GetL2ConfirmMsgByNumber fetches the L2ChainConfirm message for a given block number.
func GetL2ConfirmMsgByNumber(db *gorm.DB, start, end uint64) ([]L2ChainConfirm, error) {
	var failedConfirms []L2ChainConfirm
	tx := db.Model(&L2ChainConfirm{}).Where("number BETWEEN ? AND ? AND (deposit_status = false OR withdraw_root_status = false)", start, end)
	err := tx.Scan(&failedConfirms).Error
	if err != nil {
		return nil, err
	}
	return failedConfirms, nil
}
