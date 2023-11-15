package orm

import "gorm.io/gorm"

// BatchType batch event types.
type BatchType uint8

var (
	// BatchCommit batch commit event.
	BatchCommit BatchType = 0x01
	// BatchFinalize batch finalize event.
	BatchFinalize BatchType = 0x02
	// BatchRevert batch revert event.
	BatchRevert BatchType = 0x03
)

// L1ScrollChainEvent records l1 chain batches.
type L1ScrollChainEvent struct {
	Number     uint64 `gorm:"index"`
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

// GetBatchByIndex gets L1ScrollChainEvent by batch_index.
func GetBatchByIndex(db *gorm.DB, index uint64) (*L1ScrollChainEvent, error) {
	var scrollChain L1ScrollChainEvent
	tx := db.Model(&L1ScrollChainEvent{}).Where("batch_index = ?", index)
	err := tx.Last(&scrollChain).Error
	if err != nil {
		return nil, err
	}
	return &scrollChain, nil
}