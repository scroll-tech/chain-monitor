package orm

import (
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
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

	Amount string `gorm:"type:numeric(32,0)"`
}

// L2MessengerEvent represents an event related to L2 messenger activities.
type L2MessengerEvent struct {
	Number  uint64    `gorm:"index; comment: block number"`
	TxHash  string    `gorm:"index; type: varchar(66); comment: tx hash"`
	MsgHash string    `gorm:"primaryKey"`
	Type    EventType `gorm:"index; comment: tx type"`

	Target  common.Address `gorm:"-"`
	Message []byte         `gorm:"-"`
	Log     *types.Log     `gorm:"-"`

	Amount   string `gorm:"type:numeric(32,0)"`
	MsgNonce uint64 `gorm:"type: msg_nonce"`
	MsgProof string `gorm:"msg_proof"`
}
