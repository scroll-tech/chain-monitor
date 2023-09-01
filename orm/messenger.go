package orm

import (
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"gorm.io/gorm"
)

var (
	L1SentMessage          EventType = 1
	L1RelayedMessage       EventType = 2
	L1FailedRelayedMessage EventType = 3

	L2SentMessage          EventType = 1
	L2RelayedMessage       EventType = 2
	L2FailedRelayedMessage EventType = 3
)

type L1MessengerEvent struct {
	*TxHead
	// Make sure cross deposit or withdraw tx successfully confirmed.
	Confirm bool `gorm:"index"`
}

func SaveL1Messenger(db *gorm.DB, eventType EventType, vLog *types.Log, msgHash common.Hash) error {
	return db.Save(&L1MessengerEvent{
		TxHead: NewTxHead(vLog, eventType),
	}).Error
}

type L2MessengerEvent struct {
	Number   uint64    `gorm:"index; comment: block number"`
	Type     EventType `gorm:"index; comment: tx type"`
	MsgNonce uint64    `gorm:"primaryKey"`
	MsgHash  string    `gorm:"index"`
	MsgProof string
	Confirm  bool `gorm:"index"`
}

func SaveL2Messenger(db *gorm.DB, eventType EventType, vLog *types.Log, msgHash common.Hash) error {
	return db.Save(&L2MessengerEvent{
		Number:  vLog.BlockNumber,
		Type:    eventType,
		MsgHash: msgHash.String(),
	}).Error
}

func GetMessengerCount(db *gorm.DB, start, end uint64) (uint64, uint64, error) {
	type Result struct {
		L1Count uint64
		L2Count uint64
	}
	var result Result
	res := db.Table("l2_messenger_events as l2me").
		Select("COUNT(l1me.*) as l1_count, COUNT(l2me.*) as l2_count").
		Joins("LEFT JOIN l1_messenger_events as l1me ON l2me.message_hash = l1me.message_hash").
		Where("l2me.type = ? AND l1me.type = ? AND l2me.number BETWEEN ? AND ?", L2RelayedMessage, L1SentMessage, start, end).
		Scan(&result)
	if res.Error != nil {
		return 0, 0, res.Error
	}
	return result.L1Count, result.L2Count, nil
}
