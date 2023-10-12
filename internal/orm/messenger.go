package orm

import (
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"gorm.io/gorm"
)

const (
	// L1SentMessage l1 sent message.
	L1SentMessage EventType = 201
	// L1RelayedMessage l1 relayed message.
	L1RelayedMessage EventType = 202
	// L1FailedRelayedMessage l1 failed relayed message.
	L1FailedRelayedMessage EventType = 203

	// L2SentMessage l2 sent message.
	L2SentMessage EventType = 201
	// L2RelayedMessage l2 relayed message.
	L2RelayedMessage EventType = 202
	// L2FailedRelayedMessage l2 failed relayed message.
	L2FailedRelayedMessage EventType = 203
)

// L1MessengerEvent represents an event related to L1 messenger activities.
type L1MessengerEvent struct {
	*TxHead
	Target  common.Address `gorm:"-"`
	Message []byte         `gorm:"-"`
	Log     *types.Log     `gorm:"-"`

	// Make sure cross deposit or withdraw tx successfully confirmed.
	Confirm bool `gorm:"index"`
}

// SaveL1Messenger saves an L1 messenger event into the database.
func SaveL1Messenger(db *gorm.DB, eventType EventType, vLog *types.Log, msgHash common.Hash) error {
	return db.Save(&L1MessengerEvent{
		TxHead: NewTxHead(vLog, eventType),
	}).Error
}

// L2MessengerEvent represents an event related to L2 messenger activities.
type L2MessengerEvent struct {
	Target  common.Address `gorm:"-"`
	Message []byte         `gorm:"-"`
	Log     *types.Log     `gorm:"-"`

	Number   uint64    `gorm:"index; comment: block number"`
	Type     EventType `gorm:"index; comment: tx type"`
	MsgNonce uint64    `gorm:"type: msg_nonce"`
	MsgHash  string    `gorm:"primaryKey"`
	MsgProof string
	Confirm  bool `gorm:"index"`
}

// SaveL2Messenger saves an L2 messenger event into the database.
func SaveL2Messenger(db *gorm.DB, eventType EventType, vLog *types.Log, msgHash common.Hash) error {
	return db.Save(&L2MessengerEvent{
		Number:  vLog.BlockNumber,
		Type:    eventType,
		MsgHash: msgHash.String(),
	}).Error
}

// GetMessengerCount retrieves the count of L1 and L2 messenger events between given block numbers.
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
