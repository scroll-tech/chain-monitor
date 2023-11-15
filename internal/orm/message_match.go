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
	L1BlockStatus int    `json:"l1_block_status" gorm:"l1_block_status"`
	L1EventType   int    `json:"l1_event_type" gorm:"l1_event_type"`
	L1BlockNumber uint64 `json:"l1_block_number" gorm:"l1_block_number"`
	L1TxHash      string `json:"l1_tx_hash" gorm:"l1_tx_hash"`
	L1TokenIds    string `json:"l1_token_ids" gorm:"l1_token_ids"`
	L1Amounts     string `json:"l1_amounts" gorm:"l1_amounts"`

	// l2 event info
	L2BlockStatus int    `json:"l2_block_status" gorm:"l2_block_status"`
	L2EventType   int    `json:"l2_event_type" gorm:"l2_event_type"`
	L2BlockNumber uint64 `json:"l2_block_number" gorm:"l2_block_number"`
	L2TxHash      string `json:"l2_tx_hash" gorm:"l2_tx_hash"`
	L2TokenIds    string `json:"l2_token_ids" gorm:"l2_token_ids"`
	L2Amounts     string `json:"l2_amounts" gorm:"l2_amounts"`

	// eth info
	L1MessengerETHBalance decimal.Decimal `json:"l1_messenger_eth_balance" gorm:"l1_messenger_eth_balance"`
	L2MessengerETHBalance decimal.Decimal `json:"l2_messenger_eth_balance" gorm:"l2_messenger_eth_balance"`
	L1ETHBalanceStatus    int             `json:"l1_eth_balance_status" gorm:"l1_eth_balance_status"`
	L2ETHBalanceStatus    int             `json:"l2_eth_balance_status" gorm:"l2_eth_balance_status"`

	// status
	CheckStatus        int    `json:"check_status" gorm:"check_status"`
	WithdrawRootStatus int    `json:"withdraw_root_status" gorm:"withdraw_root_status"`
	L1ChainStatus      int    `json:"l1_chain_status" gorm:"l1_chain_status"`
	L2ChainStatus      int    `json:"l2_chain_status" gorm:"l2_chain_status"`
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

// GetLatestBlockValidMessageMatch fetches the latest valid message match record for the specified layer.
func (m *MessageMatch) GetLatestBlockValidMessageMatch(ctx context.Context, layer types.LayerType) (*MessageMatch, error) {
	var message MessageMatch
	db := m.db.WithContext(ctx)
	switch layer {
	case types.Layer1:
		db = db.Where("l1_block_status = ?", types.BlockStatusTypeValid)
	case types.Layer2:
		db = db.Where("l2_block_status = ?", types.BlockStatusTypeValid)
	}
	if err := db.Last(&message).Error; err != nil {
		log.Warn("MessageMatch.GetLatestBlockValidMessageMatch failed", "error", err)
		return nil, fmt.Errorf("MessageMatch.GetLatestBlockValidMessageMatch failed err:%w", err)
	}
	return &message, nil
}

// GetLatestDoubleValidMessageMatch fetches the latest valid message match record where both layers are valid.
func (m *MessageMatch) GetLatestDoubleValidMessageMatch(ctx context.Context) (*MessageMatch, error) {
	var message MessageMatch
	db := m.db.WithContext(ctx)

	// Look for records where both layers are valid
	db = db.Where("l1_block_status = ?", types.BlockStatusTypeValid)
	db = db.Where("l2_block_status = ?", types.BlockStatusTypeValid)

	if err := db.Last(&message).Error; err != nil {
		log.Warn("MessageMatch.GetLatestDoubleValidMessageMatch failed", "error", err)
		return nil, fmt.Errorf("MessageMatch.GetLatestDoubleValidMessageMatch failed err:%w", err)
	}
	return &message, nil
}

// GetLatestValidCrossChainMessageMatch fetches the latest valid cross chain message match for the specified layer.
func (m *MessageMatch) GetLatestValidCrossChainMessageMatch(ctx context.Context, layerType types.LayerType) (*MessageMatch, error) {
	var message MessageMatch
	db := m.db.WithContext(ctx)
	switch layerType {
	case types.Layer1:
		db = db.Where("l1_cross_chain_status = ?", types.CrossChainStatusTypeValid)
	case types.Layer2:
		db = db.Where("l2_cross_chain_status = ?", types.CrossChainStatusTypeValid)
	}
	if err := db.Last(&message).Error; err != nil {
		log.Warn("MessageMatch.GetLatestValidCrossChainMessageMatch failed", "error", err)
		return nil, fmt.Errorf("MessageMatch.GetLatestValidCrossChainMessageMatch failed err:%w", err)
	}
	return &message, nil
}

// GetLatesValidETHBalanceMessageMatch fetches the latest valid Ethereum balance match record for the specified layer.
func (m *MessageMatch) GetLatesValidETHBalanceMessageMatch(ctx context.Context, layer types.LayerType) (*MessageMatch, error) {
	var message MessageMatch
	db := m.db.WithContext(ctx)
	switch layer {
	case types.Layer1:
		db = db.Where("l1_eth_balance_status = ?", types.ETHBalanceStatusTypeValid)
	case types.Layer2:
		db = db.Where("l2_eth_balance_status = ?", types.ETHBalanceStatusTypeValid)
	}
	if err := db.Last(&message).Error; err != nil {
		log.Warn("MessageMatch.GetLatestBlockValidMessageMatch failed", "error", err)
		return nil, fmt.Errorf("MessageMatch.GetLatestBlockValidMessageMatch failed err:%w", err)
	}
	return &message, nil
}

// GetMessageMatchesByBlockNumberRange fetches all MessageMatch records in the block number range (inclusive).
func (m *MessageMatch) GetMessageMatchesByBlockNumberRange(ctx context.Context, layer types.LayerType, startHeight, endHeight uint64) ([]MessageMatch, error) {
	var messages []MessageMatch
	db := m.db.WithContext(ctx)
	switch layer {
	case types.Layer1:
		db = db.Where("l1_block_number >= ?", startHeight).Where("l1_block_number <= ?", endHeight)
	case types.Layer2:
		db = db.Where("l2_block_number >= ?", startHeight).Where("l2_block_number <= ?", endHeight)
	}
	if err := db.Find(&messages).Error; err != nil {
		log.Warn("MessageMatch.GetMessageMatchesByBlockNumberRange failed", "start height", startHeight, "end height", endHeight, "error", err)
		return nil, fmt.Errorf("MessageMatch.GetMessageMatchesByBlockNumberRange failed, start height: %v, end height: %v, err: %w", startHeight, endHeight, err)
	}
	return messages, nil
}

// GetMessageMatchByL2BlockNumber fetches the message match record by L2 block number with the maximum id.
func (m *MessageMatch) GetMessageMatchByL2BlockNumber(ctx context.Context, blockNumber uint64) (*MessageMatch, error) {
	var message MessageMatch
	db := m.db.WithContext(ctx)
	db = db.Where("l2_block_number = ?", blockNumber)
	db = db.Order("id desc")
	db = db.Limit(1)
	if err := db.Find(&message).Error; err != nil {
		log.Warn("GetMessageMatchByL2BlockNumber failed", "block number", blockNumber, "error", err)
		return nil, fmt.Errorf("GetMessageMatchByL2BlockNumber failed, block number:%v, err:%w", blockNumber, err)
	}
	return &message, nil
}

// InsertOrUpdate is a placeholder function for inserting or updating message matches.
// @todo: message insert everywhere, ensure l1_block_status & l2_block_status are updated correctly.
func (m *MessageMatch) InsertOrUpdate(ctx context.Context, messages []MessageMatch) (int, error) {
	// insert or update
	return 0, nil
}

// UpdateGatewayStatus updates the gateway status for the message matches with the provided ids.
func (m *MessageMatch) UpdateGatewayStatus(ctx context.Context, id []int64, layerType types.LayerType, status types.GatewayStatusType) error {
	db := m.db.WithContext(ctx)
	db = db.Model(&MessageMatch{})
	db = db.Where("id = (?)", id)

	var err error
	switch layerType {
	case types.Layer1:
		err = db.Update("l1_chain_status", status).Error
	case types.Layer2:
		err = db.Update("l2_chain_status", status).Error
	}

	if err != nil {
		log.Warn("MessageMatch.UpdateGatewayStatus failed", "error", err)
		return fmt.Errorf("MessageMatch.UpdateGatewayStatus failed err:%w", err)
	}
	return nil
}

// UpdateCrossChainStatus updates the cross chain status for the message matches with the provided ids.
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
