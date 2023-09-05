package orm

import (
	"gorm.io/gorm"
)

// L1Block represents a block on Layer 1. It has a unique block number and the count of events.
type L1Block struct {
	Number     uint64 `gorm:"primaryKey; comment: the last block number"`
	EventCount int
}

// L2Block represents a block on Layer 2. It has a unique block number and the count of events.
type L2Block struct {
	Number     uint64 `gorm:"primaryKey; comment: block number"`
	EventCount int
}

// GetL1LatestBlock retrieves the latest L1Block from the database.
func GetL1LatestBlock(db *gorm.DB) (*L1Block, error) {
	var block L1Block
	err := db.Last(&block).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &block, nil
}

// GetL2LatestBlock retrieves the latest L2Block from the database.
func GetL2LatestBlock(db *gorm.DB) (*L2Block, error) {
	var block L2Block
	err := db.Last(&block).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &block, nil
}
