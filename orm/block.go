package orm

import (
	"gorm.io/gorm"
)

type L1Block struct {
	Number     uint64 `gorm:"primaryKey; comment: the last block number"`
	EventCount int
}

type L2Block struct {
	Number     uint64 `gorm:"primaryKey; comment: block number"`
	EventCount int
}

func GetL1LatestBlock(db *gorm.DB) (*L1Block, error) {
	var block L1Block
	err := db.Last(&block).Error
	if err != nil && (err.Error() != gorm.ErrRecordNotFound.Error()) {
		return nil, err
	}
	return &block, nil
}

func GetL2LatestBlock(db *gorm.DB) (*L2Block, error) {
	var block L2Block
	err := db.Last(&block).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &block, nil
}
