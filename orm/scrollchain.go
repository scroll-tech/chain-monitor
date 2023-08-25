package orm

import (
	"gorm.io/gorm"
)

type BatchType uint8

var (
	BatchCommit   BatchType = 0x01
	BatchFinalize BatchType = 0x02
	BatchRevert   BatchType = 0x03
)

// ScrollChainEvent records l1 chain batches.
type ScrollChainEvent struct {
	BatchIndex uint64 `gorm:"index; comment: batch index, increase one by one"`
	BatchHash  string `gorm:"primaryKey; comment: batch content hash, it is unique"`

	CommitNumber uint64 `gorm:"index; comment: commit l1chain number"`
	CommitHash   string

	FinalizeNumber uint64 `gorm:"index; comment: finalize l1chain number"`
	FinalizeHash   string

	RevertNumber uint64 `gorm:"index; comment: revert batch l1chain number"`
	RevertHash   string

	L2StartNumber uint64 `gorm:"comment: l2chain block start number contained in this batch"`
	L2EndNumber   uint64 `gorm:"comment: l2chain block end number contained in this batch"`

	Status BatchType `gorm:"index"`
}

func GetScrollChainByIndex(db *gorm.DB, index uint64) (*ScrollChainEvent, error) {
	tx := db.Where("batch_index = ?", index)
	var batch ScrollChainEvent
	err := tx.First(&batch).Error
	if err != nil {
		return nil, err
	}
	return &batch, nil
}
