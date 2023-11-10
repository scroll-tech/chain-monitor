package orm

import (
	"context"
	"fmt"
	"time"

	"github.com/scroll-tech/go-ethereum/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/types"
)

// MessageMatch contains the tx of l1 & l2
type MessageMatch struct {
	db *gorm.DB `gorm:"column:-"`

	ID          int64  `json:"id" gorm:"column:id"`
	MessageHash string `json:"message_hash" gorm:"message_hash"`
	TokenType   int    `json:"token_type" gorm:"token_type"`

	// l1 event info
	L1BlockStatus int             `json:"l1_block_status" gorm:"l1_block_status"`
	L1EventType   int             `json:"l1_event_type" gorm:"l1_event_type"`
	L1BlockNumber uint64          `json:"l1_block_number" gorm:"l1_block_number"`
	L1TxHash      string          `json:"l1_tx_hash" gorm:"l1_tx_hash"`
	L1TokenId     string          `json:"l1_token_id" gorm:"l1_token_id"`
	L1Amount      decimal.Decimal `json:"l1_amount" gorm:"l1_amount"`

	// l2 event info
	L2BlockStatus int             `json:"l2_block_status" gorm:"l2_block_status"`
	L2EventType   int             `json:"l2_event_type" gorm:"l2_event_type"`
	L2BlockNumber uint64          `json:"l2_block_number" gorm:"l2_block_number"`
	L2TxHash      string          `json:"l2_tx_hash" gorm:"l2_tx_hash"`
	L2TokenId     string          `json:"l2_token_id" gorm:"l2_token_id"`
	L2Amount      decimal.Decimal `json:"l2_amount" gorm:"l2_amount"`

	// status
	CheckStatus        int    `json:"check_status" gorm:"check_status"`
	WithdrawRootStatus int    `json:"withdraw_root_status" gorm:"withdraw_root_status"`
	L1GatewayStatus    int    `json:"l1_gateway_status" gorm:"l1_gateway_status"`
	L2GatewayStatus    int    `json:"l2_gateway_status" gorm:"l2_gateway_status"`
	L1CrossChainStatus int    `json:"l1_cross_chain_status" gorm:"l1_cross_chain_status"`
	L2CrossChainStatus int    `json:"l2_cross_chain_status" gorm:"l2_cross_chain_status"`
	MessageProof       []byte `json:"message_proof" gorm:"message_proof"` // only not null in the last message of each block.
	MessageNonce       uint64 `json:"message_nonce" gorm:"message_nonce"` // only not null in the last message of each block.

	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

// NewMessageMatch creates a new MessageMatch database instance.
func NewMessageMatch(db *gorm.DB) *MessageMatch {
	return &MessageMatch{db: db}
}

// TableName returns the table name for the Batch model.
func (*MessageMatch) TableName() string {
	return "message_match"
}

// GetLatestBlockValidMessageMatch get the latest layer message match record
func (m *MessageMatch) GetLatestBlockValidMessageMatch(ctx context.Context, layer1 types.LayerType) (*MessageMatch, error) {
	var transaction MessageMatch
	db := m.db.WithContext(ctx)
	switch layer1 {
	case types.Layer1:
		db = db.Where("l1_block_status = ?", types.BlockStatusTypeValid)
	case types.Layer2:
		db = db.Where("l2_block_status = ?", types.BlockStatusTypeValid)
	}
	if err := db.Last(&transaction).Error; err != nil {
		log.Warn("MessageMatch.GetLatestBlockValidMessageMatch failed", "error", err)
		return nil, fmt.Errorf("MessageMatch.GetLatestBlockValidMessageMatch failed err:%w", err)
	}
	return &transaction, nil
}

// GetMessageMatchByL2BlockNumber gets message match record by L2 block number.
func (m *MessageMatch) GetMessageMatchByL2BlockNumber(ctx context.Context, blockNumber uint64) (*MessageMatch, error) {
	var message MessageMatch
	db := m.db.WithContext(ctx)
	db = db.Where("l2_block_number = ?", blockNumber)
	db = db.Order("id asc")
	db = db.Limit(1)
	if err := db.Find(&message).Error; err != nil {
		log.Warn("GetMessageMatchByL2BlockNumber failed", "block number", blockNumber, "error", err)
		return nil, fmt.Errorf("GetMessageMatchByL2BlockNumber failed, block number:%v, err:%w", blockNumber, err)
	}
	return &message, nil
}

// GetUncheckedLatestMessageMatch get the latest uncheck message match record
func (m *MessageMatch) GetUncheckedLatestMessageMatch(ctx context.Context, limit int) ([]MessageMatch, error) {
	var messages []MessageMatch
	db := m.db.WithContext(ctx)
	db = db.Where("check_status = ?", types.CheckStatusUnchecked)
	db = db.Order("id asc")
	db = db.Limit(limit)
	if err := db.Find(&messages).Error; err != nil {
		log.Warn("MessageMatch.GetUncheckedLatestMessageMatch failed", "error", err)
		return nil, fmt.Errorf("MessageMatch.GetUncheckedLatestMessageMatch failed err:%w", err)
	}
	return messages, nil
}

func (m *MessageMatch) InsertOrUpdate(ctx context.Context, messages []MessageMatch) (int, error) {
	// insert or update
	return 0, nil
}

func (m *MessageMatch) UpdateGatewayStatus(ctx context.Context, id []int64, layerType types.LayerType, status types.GatewayStatusType) error {
	db := m.db.WithContext(ctx)
	db = db.Model(&MessageMatch{})
	db = db.Where("id = (?)", id)

	var err error
	switch layerType {
	case types.Layer1:
		err = db.Update("l1_gateway_status", status).Error
	case types.Layer2:
		err = db.Update("l2_gateway_status", status).Error
	}

	if err != nil {
		log.Warn("MessageMatch.UpdateGatewayStatus failed", "error", err)
		return fmt.Errorf("MessageMatch.UpdateGatewayStatus failed err:%w", err)
	}
	return nil
}

func (m *MessageMatch) UpdateCrossChainStatus(ctx context.Context, id []int64, layerType types.LayerType, status types.CrossChainStatusType) error {
	db := m.db.WithContext(ctx)
	db = db.Model(&MessageMatch{})
	db = db.Where("id in (?)", id)

	var err error
	switch layerType {
	case types.Layer1:
		err = db.Update("l1_cross_chain_status", status).Error
	case types.Layer2:
		err = db.Update("l2_cross_chain_status", status).Error
	}

	if err != nil {
		log.Warn("MessageMatch.UpdateCrossChainStatus failed", "error", err)
		return fmt.Errorf("MessageMatch.UpdateCrossChainStatus failed err:%w", err)
	}
	return nil
}
