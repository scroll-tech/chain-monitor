package orm

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/scroll-tech/go-ethereum/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/scroll-tech/chain-monitor/internal/types"
	"github.com/scroll-tech/chain-monitor/internal/utils"
)

// MessageMatch contains the tx of l1 & l2
type MessageMatch struct {
	db *gorm.DB `gorm:"column:-"`

	ID          int64  `json:"id" gorm:"column:id"`
	MessageHash string `json:"message_hash" gorm:"message_hash"`
	TokenType   int    `json:"token_type" gorm:"token_type"`

	// l1 event info
	L1EventType   int    `json:"l1_event_type" gorm:"l1_event_type"`
	L1BlockNumber uint64 `json:"l1_block_number" gorm:"l1_block_number"`
	L1TxHash      string `json:"l1_tx_hash" gorm:"l1_tx_hash"`
	L1TokenIds    string `json:"l1_token_ids" gorm:"l1_token_ids"`
	L1Amounts     string `json:"l1_amounts" gorm:"l1_amounts"`

	// l2 event info
	L2EventType   int    `json:"l2_event_type" gorm:"l2_event_type"`
	L2BlockNumber uint64 `json:"l2_block_number" gorm:"l2_block_number"`
	L2TxHash      string `json:"l2_tx_hash" gorm:"l2_tx_hash"`
	L2TokenIds    string `json:"l2_token_ids" gorm:"l2_token_ids"`
	L2Amounts     string `json:"l2_amounts" gorm:"l2_amounts"`

	// eth event info
	L1MessengerETHBalance decimal.Decimal `json:"l1_messenger_eth_balance" gorm:"l1_messenger_eth_balance"`
	L1ETHBalanceStatus    int             `json:"l1_eth_balance_status" gorm:"l1_eth_balance_status"`
	L2MessengerETHBalance decimal.Decimal `json:"l2_messenger_eth_balance" gorm:"l2_messenger_eth_balance"`
	L2ETHBalanceStatus    int             `json:"l2_eth_balance_status" gorm:"l2_eth_balance_status"`
	ETHAmount             string          `json:"eth_amount" gorm:"eth_amount"`
	ETHAmountStatus       int             `json:"eth_amount_status" gorm:"eth_amount_status"`

	// status
	L1BlockStatus      int `json:"l1_block_status" gorm:"l1_block_status"`
	L2BlockStatus      int `json:"l2_block_status" gorm:"l2_block_status"`
	L1CrossChainStatus int `json:"l1_cross_chain_status" gorm:"l1_cross_chain_status"`
	L2CrossChainStatus int `json:"l2_cross_chain_status" gorm:"l2_cross_chain_status"`
	WithdrawRootStatus int `json:"withdraw_root_status" gorm:"withdraw_root_status"`

	// only not null in the last message of each block.
	MessageProof []byte `json:"message_proof" gorm:"message_proof"`
	// only not null in l2 sent messages, and use next message nonce (+1) to distinguish from the zero values.
	NextMessageNonce uint64 `json:"next_message_nonce" gorm:"next_message_nonce"`

	L1BlockStatusUpdatedAt      time.Time      `json:"l1_block_status_updated_at" gorm:"l1_block_status_updated_at"`
	L2BlockStatusUpdatedAt      time.Time      `json:"l2_block_status_updated_at" gorm:"l2_block_status_updated_at"`
	L1CrossChainStatusUpdatedAt time.Time      `json:"l1_cross_chain_status_updated_at" gorm:"l1_cross_chain_status_updated_at"`
	L2CrossChainStatusUpdatedAt time.Time      `json:"l2_cross_chain_status_updated_at" gorm:"l2_cross_chain_status_updated_at"`
	L1EthBalanceStatusUpdatedAt time.Time      `json:"l1_eth_balance_status_updated_at" gorm:"l1_eth_balance_status_updated_at"`
	L2EthBalanceStatusUpdatedAt time.Time      `json:"l2_eth_balance_status_updated_at" gorm:"l2_eth_balance_status_updated_at"`
	MessageProofUpdatedAt       time.Time      `json:"message_proof_updated_at" gorm:"message_proof_updated_at"`
	CreatedAt                   time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt                   time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt                   gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

// NewMessageMatch creates a new MessageMatch database instance.
func NewMessageMatch(db *gorm.DB) *MessageMatch {
	return &MessageMatch{db: db}
}

// TableName returns the table name for the Batch model.
func (*MessageMatch) TableName() string {
	return "message_match"
}

// GetUncheckedAndDoubleLayerValidGatewayMessageMatches retrieves the earliest unchecked gateway message match records
// that are valid in both Layer1 and Layer2.
func (m *MessageMatch) GetUncheckedAndDoubleLayerValidGatewayMessageMatches(ctx context.Context, layer types.LayerType, limit int) ([]MessageMatch, error) {
	var messages []MessageMatch
	db := m.db.WithContext(ctx)
	db = db.Where("l1_block_status = ?", types.BlockStatusTypeValid)
	db = db.Where("l2_block_status = ?", types.BlockStatusTypeValid)
	switch layer {
	case types.Layer1:
		db = db.Where("l1_cross_chain_status = ?", types.CrossChainStatusTypeInvalid)
	case types.Layer2:
		db = db.Where("l2_cross_chain_status = ?", types.CrossChainStatusTypeInvalid)
	}
	db = db.Limit(limit)
	if err := db.Find(&messages).Error; err != nil {
		log.Warn("MessageMatch.GetUncheckedAndDoubleLayerValidGatewayMessageMatches failed", "error", err)
		return nil, fmt.Errorf("MessageMatch.GetUncheckedAndDoubleLayerValidGatewayMessageMatches failed err:%w", err)
	}
	return messages, nil
}

// GetUncheckedLatestETHMessageMatch get the latest uncheck eth message match record
func (m *MessageMatch) GetUncheckedLatestETHMessageMatch(ctx context.Context, layer types.LayerType, limit int) ([]MessageMatch, error) {
	var messages []MessageMatch
	db := m.db.WithContext(ctx)
	switch layer {
	case types.Layer1:
		db = db.Where("l1_eth_balance_status = ?", types.ETHBalanceStatusTypeInvalid)
		db = db.Where("l1_block_status = ?", types.BlockStatusTypeValid)
		db = db.Order("l1_block_number asc")
	case types.Layer2:
		db = db.Where("l2_eth_balance_status = ?", types.ETHBalanceStatusTypeInvalid)
		db = db.Where("l2_block_status = ?", types.BlockStatusTypeValid)
		db = db.Order("l2_block_number asc")
	}
	db = db.Where("token_type = ? or token_type = ?", types.TokenTypeETH, types.TokenTypeERC20)
	db = db.Limit(limit)
	if err := db.Find(&messages).Error; err != nil {
		log.Warn("MessageMatch.GetUncheckedLatestETHMessageMatch failed", "error", err)
		return nil, fmt.Errorf("MessageMatch.GetUncheckedLatestETHMessageMatch failed err:%w", err)
	}
	return messages, nil
}

// GetLatestBlockValidMessageMatch fetches the latest valid message match record for the specified layer.
func (m *MessageMatch) GetLatestBlockValidMessageMatch(ctx context.Context, layer types.LayerType) (*MessageMatch, error) {
	var message MessageMatch
	db := m.db.WithContext(ctx)
	switch layer {
	case types.Layer1:
		db = db.Where("l1_block_status = ?", types.BlockStatusTypeValid)
		db = db.Order("l1_block_number desc")
	case types.Layer2:
		db = db.Where("l2_block_status = ?", types.BlockStatusTypeValid)
		db = db.Order("l2_block_number desc")
	}
	err := db.First(&message).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Warn("MessageMatch.GetLatestBlockValidMessageMatch failed", "error", err)
		return nil, fmt.Errorf("MessageMatch.GetLatestBlockValidMessageMatch failed err:%w", err)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &message, nil
}

// GetETHCheckStartBlockNumberAndBalance fetches the latest valid Ethereum balance match record for the specified layer
// and returns the block number and messenger balance for the specified layer.
func (m *MessageMatch) GetETHCheckStartBlockNumberAndBalance(ctx context.Context, layer types.LayerType) (*big.Int, error) {
	var message MessageMatch
	db := m.db.WithContext(ctx)
	switch layer {
	case types.Layer1:
		db = db.Where("l1_eth_balance_status = ?", types.ETHBalanceStatusTypeValid)
		db = db.Order("l1_block_number desc")
	case types.Layer2:
		db = db.Where("l2_eth_balance_status = ?", types.ETHBalanceStatusTypeValid)
		db = db.Order("l2_block_number desc")
	}
	err := db.First(&message).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return big.NewInt(0), nil
		}
		log.Warn("MessageMatch.GetETHCheckStartBlockNumberAndBalance failed", "error", err)
		return big.NewInt(0), fmt.Errorf("MessageMatch.GetETHCheckStartBlockNumberAndBalance failed err:%w", err)
	}

	// Return the block number and messenger balance for the specified layer
	switch layer {
	case types.Layer1:
		return message.L1MessengerETHBalance.BigInt(), nil
	case types.Layer2:
		return message.L2MessengerETHBalance.BigInt(), nil
	default:
		return big.NewInt(0), fmt.Errorf("invalid layer: %v", layer)
	}
}

// GetLatestValidL2SentMessageMatch fetches the valid l2 sent message with the largest message nonce.
func (m *MessageMatch) GetLatestValidL2SentMessageMatch(ctx context.Context) (*MessageMatch, error) {
	var message MessageMatch
	db := m.db.WithContext(ctx)
	db = db.Where("withdraw_root_status = ?", types.WithdrawRootStatusTypeValid)
	db = db.Where("next_message_nonce > 0")
	db = db.Order("next_message_nonce DESC")
	err := db.First(&message).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		log.Warn("GetLatestValidL2SentMessageMatch failed", "error", err)
		return nil, fmt.Errorf("GetLatestValidL2SentMessageMatch failed, err:%w", err)
	}
	return &message, nil
}

// GetL2SentMessagesInBlockRange fetches the message match records of l2 sent message within the block range.
func (m *MessageMatch) GetL2SentMessagesInBlockRange(ctx context.Context, startBlockNumber, endBlockNumber uint64) ([]*MessageMatch, error) {
	var messages []*MessageMatch
	db := m.db.WithContext(ctx)
	db = db.Where("withdraw_root_status = ?", types.WithdrawRootStatusTypeUnknown)
	db = db.Where("l2_block_number >= ?", startBlockNumber)
	db = db.Where("l2_block_number <= ?", endBlockNumber)
	db = db.Where("next_message_nonce > 0")
	db = db.Order("next_message_nonce ASC")
	err := db.Find(&messages).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		log.Warn("GetL2SentMessagesInBlockRange failed", "error", err)
		return nil, fmt.Errorf("GetL2SentMessagesInBlockRange failed, err:%w", err)
	}
	return messages, nil
}

// InsertOrUpdateEventInfo insert or update event info
func (m *MessageMatch) InsertOrUpdateEventInfo(ctx context.Context, layer types.LayerType, message MessageMatch, dbTX ...*gorm.DB) (int64, error) {
	db := m.db
	if len(dbTX) > 0 && dbTX[0] != nil {
		db = dbTX[0]
	}

	db = db.WithContext(ctx)
	db = db.Model(&MessageMatch{})
	var assignmentColumn clause.Set
	if layer == types.Layer1 {
		if message.L1EventType == int(types.L1SentMessage) { // sent
			assignmentColumn = clause.AssignmentColumns([]string{"token_type", "l1_event_type", "l1_block_number", "l1_tx_hash", "l1_amounts", "l2_amounts", "eth_amount", "eth_amount_status"})
		} else if message.L1EventType == int(types.L1RelayedMessage) { // relayed
			assignmentColumn = clause.AssignmentColumns([]string{"token_type", "l1_event_type", "l1_block_number", "l1_tx_hash", "l1_token_ids"})
		} else if message.L1EventType == int(types.L1DepositERC20) || message.L1EventType == int(types.L1DepositERC721) || message.L1EventType == int(types.L1DepositERC1155) { // sent
			assignmentColumn = clause.AssignmentColumns([]string{"token_type", "l1_block_number", "l1_tx_hash", "l1_event_type", "l1_token_ids", "l1_amounts", "eth_amount", "eth_amount_status"})
		} else { // relayed
			assignmentColumn = clause.AssignmentColumns([]string{"token_type", "l1_block_number", "l1_tx_hash", "l1_event_type", "l1_token_ids", "l1_amounts"})
		}
	}

	if layer == types.Layer2 {
		if message.L2EventType == int(types.L2SentMessage) { // sent
			assignmentColumn = clause.AssignmentColumns([]string{"token_type", "l2_event_type", "l2_block_number", "l2_tx_hash", "l1_amounts", "l2_amounts", "eth_amount", "eth_amount_status"})
		} else if message.L2EventType == int(types.L2RelayedMessage) { // relayed
			assignmentColumn = clause.AssignmentColumns([]string{"token_type", "l2_event_type", "l2_block_number", "l2_tx_hash"})
		} else if message.L2EventType == int(types.L2WithdrawERC20) || message.L2EventType == int(types.L2WithdrawERC721) || message.L2EventType == int(types.L2WithdrawERC1155) { // sent
			assignmentColumn = clause.AssignmentColumns([]string{"token_type", "l2_block_number", "l2_tx_hash", "l2_event_type", "l2_token_ids", "l2_amounts", "eth_amount", "eth_amount_status"})
		} else { // relayed
			assignmentColumn = clause.AssignmentColumns([]string{"token_type", "l2_block_number", "l2_tx_hash", "l2_event_type", "l2_token_ids", "l2_amounts"})
		}
	}

	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "message_hash"}},
		DoUpdates: assignmentColumn,
	})

	result := db.Create(&message)
	if result.Error != nil {
		return 0, fmt.Errorf("MessageMatch.InsertOrUpdateGatewayEventInfo error: %w, messages: %v", result.Error, message)
	}
	return result.RowsAffected, nil
}

// UpdateCrossChainStatus updates the cross chain status for the message matches with the provided ids.
func (m *MessageMatch) UpdateCrossChainStatus(ctx context.Context, id []int64, layer types.LayerType, status types.CrossChainStatusType) error {
	db := m.db.WithContext(ctx)
	db = db.Model(&MessageMatch{})
	db = db.Where("id in (?)", id)

	var updateFields map[string]interface{}
	switch layer {
	case types.Layer1:
		updateFields = map[string]interface{}{
			"l1_cross_chain_status":            status,
			"l1_cross_chain_status_updated_at": utils.NowUTC(),
		}
	case types.Layer2:
		updateFields = map[string]interface{}{
			"l2_cross_chain_status":            status,
			"l2_cross_chain_status_updated_at": utils.NowUTC(),
		}
	}

	if err := db.Updates(updateFields).Error; err != nil {
		log.Warn("MessageMatch.UpdateCrossChainStatus failed", "error", err)
		return fmt.Errorf("MessageMatch.UpdateCrossChainStatus failed err:%w", err)
	}
	return nil
}

// UpdateMsgProofAndStatus insert or update the withdrawal tree root's message proof and withdraw root status
func (m *MessageMatch) UpdateMsgProofAndStatus(ctx context.Context, message *MessageMatch, dbTX ...*gorm.DB) error {
	if message == nil {
		return nil
	}
	db := m.db
	if len(dbTX) > 0 && dbTX[0] != nil {
		db = dbTX[0]
	}
	db = db.WithContext(ctx)
	db = db.Model(&MessageMatch{})
	db = db.Where("message_hash = ?", message.MessageHash)

	updateFields := map[string]interface{}{
		"message_proof":        message.MessageProof,
		"withdraw_root_status": message.WithdrawRootStatus,
	}

	if err := db.Updates(updateFields).Error; err != nil {
		return fmt.Errorf("MessageMatch.UpdateMsgProofAndStatus failed err:%w", err)
	}
	return nil
}

// UpdateBlockStatus updates the block status for the given layer and block number range.
func (m *MessageMatch) UpdateBlockStatus(ctx context.Context, layer types.LayerType, startBlockNumber, endBlockNumber uint64, dbTX ...*gorm.DB) error {
	db := m.db
	if len(dbTX) > 0 && dbTX[0] != nil {
		db = dbTX[0]
	}

	db = db.WithContext(ctx)
	db = db.Model(&MessageMatch{})

	var updateFields map[string]interface{}
	switch layer {
	case types.Layer1:
		db = db.Where("l1_block_number >= ? AND l1_block_number <= ?", startBlockNumber, endBlockNumber)
		updateFields = map[string]interface{}{
			"l1_block_status":            types.BlockStatusTypeValid,
			"l1_block_status_updated_at": utils.NowUTC(),
		}
	case types.Layer2:
		db = db.Where("l2_block_number >= ? AND l2_block_number <= ?", startBlockNumber, endBlockNumber)
		updateFields = map[string]interface{}{
			"l2_block_status":            types.BlockStatusTypeValid,
			"l2_block_status_updated_at": utils.NowUTC(),
		}
	}

	if err := db.Updates(updateFields).Error; err != nil {
		return fmt.Errorf("MessageMatch.UpdateBlockStatus failed, start block number: %v, end block number: %v, err: %w", startBlockNumber, endBlockNumber, db.Error)
	}
	return nil
}

// UpdateETHBalance update the eth balance and eth status
func (m *MessageMatch) UpdateETHBalance(ctx context.Context, layer types.LayerType, messageMatch MessageMatch, dbTX ...*gorm.DB) error {
	db := m.db
	if len(dbTX) > 0 && dbTX[0] != nil {
		db = dbTX[0]
	}

	db = db.WithContext(ctx)
	db = db.Model(&MessageMatch{})
	db = db.Where("id = ?", messageMatch.ID)

	var updateFields map[string]interface{}
	switch layer {
	case types.Layer1:
		updateFields = map[string]interface{}{
			"l1_messenger_eth_balance":         messageMatch.L1MessengerETHBalance,
			"l1_eth_balance_status":            types.ETHBalanceStatusTypeValid,
			"l1_eth_balance_status_updated_at": utils.NowUTC(),
		}
	case types.Layer2:
		updateFields = map[string]interface{}{
			"l2_messenger_eth_balance":         messageMatch.L2MessengerETHBalance,
			"l2_eth_balance_status":            types.ETHBalanceStatusTypeValid,
			"l2_eth_balance_status_updated_at": utils.NowUTC(),
		}
	}
	if err := db.Updates(updateFields).Error; err != nil {
		log.Warn("MessageMatch.UpdateETHBalance failed", "error", err)
		return fmt.Errorf("MessageMatch.UpdateETHBalance failed err:%w", err)
	}
	return nil
}
