package orm

import (
	"context"
	"fmt"
	"time"

	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/scroll-tech/chain-monitor/internal/types"
	"github.com/scroll-tech/chain-monitor/internal/utils"
)

// GatewayMessageMatch contains the gateway tx of l1 & l2
type GatewayMessageMatch struct {
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

	// status
	L1BlockStatus      int `json:"l1_block_status" gorm:"l1_block_status"`
	L2BlockStatus      int `json:"l2_block_status" gorm:"l2_block_status"`
	L1CrossChainStatus int `json:"l1_cross_chain_status" gorm:"l1_cross_chain_status"`
	L2CrossChainStatus int `json:"l2_cross_chain_status" gorm:"l2_cross_chain_status"`

	L1BlockStatusUpdatedAt      time.Time      `json:"l1_block_status_updated_at" gorm:"l1_block_status_updated_at"`
	L2BlockStatusUpdatedAt      time.Time      `json:"l2_block_status_updated_at" gorm:"l2_block_status_updated_at"`
	L1CrossChainStatusUpdatedAt time.Time      `json:"l1_cross_chain_status_updated_at" gorm:"l1_cross_chain_status_updated_at"`
	L2CrossChainStatusUpdatedAt time.Time      `json:"l2_cross_chain_status_updated_at" gorm:"l2_cross_chain_status_updated_at"`
	CreatedAt                   time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt                   time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt                   gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

// NewGatewayMessageMatch creates a new GatewayMessageMatch database instance.
func NewGatewayMessageMatch(db *gorm.DB) *GatewayMessageMatch {
	return &GatewayMessageMatch{db: db}
}

// TableName returns the table name for the Batch model.
func (*GatewayMessageMatch) TableName() string {
	return "gateway_message_match"
}

// GetUncheckedAndDoubleLayerValidGatewayMessageMatches retrieves the earliest unchecked gateway message match records
// that are valid in both Layer1 and Layer2.
func (m *GatewayMessageMatch) GetUncheckedAndDoubleLayerValidGatewayMessageMatches(ctx context.Context, layer types.LayerType, limit int) ([]GatewayMessageMatch, error) {
	var messages []GatewayMessageMatch
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
		log.Warn("GatewayMessageMatch.GetUncheckedAndDoubleLayerValidGatewayMessageMatches failed", "error", err)
		return nil, fmt.Errorf("GatewayMessageMatch.GetUncheckedAndDoubleLayerValidGatewayMessageMatches failed err:%w", err)
	}
	return messages, nil
}

// InsertOrUpdateEventInfo insert or update event info
func (m *GatewayMessageMatch) InsertOrUpdateEventInfo(ctx context.Context, layer types.LayerType, message GatewayMessageMatch, dbTX ...*gorm.DB) (int64, error) {
	db := m.db
	if len(dbTX) > 0 && dbTX[0] != nil {
		db = dbTX[0]
	}

	db = db.WithContext(ctx)
	db = db.Model(&GatewayMessageMatch{})
	var assignmentColumn clause.Set
	var where clause.Where
	if layer == types.Layer1 {
		assignmentColumn = clause.AssignmentColumns([]string{"token_type", "l1_block_number", "l1_tx_hash", "l1_event_type", "l1_token_ids", "l1_amounts"})
		where = clause.Where{Exprs: []clause.Expression{clause.Eq{Column: "gateway_message_match.l1_block_number", Value: 0}}}
	} else {
		assignmentColumn = clause.AssignmentColumns([]string{"token_type", "l2_block_number", "l2_tx_hash", "l2_event_type", "l2_token_ids", "l2_amounts"})
		where = clause.Where{Exprs: []clause.Expression{clause.Eq{Column: "gateway_message_match.l2_block_number", Value: 0}}}
	}

	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "message_hash"}},
		Where:     where,
		DoUpdates: assignmentColumn,
	})

	result := db.Create(&message)
	if result.Error != nil {
		return 0, fmt.Errorf("GatewayMessageMatch.InsertOrUpdateGatewayEventInfo error: %w, messages: %v", result.Error, message)
	}
	return result.RowsAffected, nil
}

// UpdateCrossChainStatus updates the cross chain status for the message matches with the provided ids.
func (m *GatewayMessageMatch) UpdateCrossChainStatus(ctx context.Context, id []int64, layer types.LayerType, status types.CrossChainStatusType) error {
	db := m.db.WithContext(ctx)
	db = db.Model(&GatewayMessageMatch{})
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
		log.Warn("GatewayMessageMatch.UpdateCrossChainStatus failed", "error", err)
		return fmt.Errorf("GatewayMessageMatch.UpdateCrossChainStatus failed err:%w", err)
	}
	return nil
}

// UpdateBlockStatus updates the block status for the given layer and block number range.
func (m *GatewayMessageMatch) UpdateBlockStatus(ctx context.Context, layer types.LayerType, startBlockNumber, endBlockNumber uint64, dbTX ...*gorm.DB) error {
	db := m.db
	if len(dbTX) > 0 && dbTX[0] != nil {
		db = dbTX[0]
	}

	db = db.WithContext(ctx)
	db = db.Model(&GatewayMessageMatch{})

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
		return fmt.Errorf("GatewayMessageMatch.UpdateBlockStatus failed, start block number: %v, end block number: %v, err: %w", startBlockNumber, endBlockNumber, db.Error)
	}
	return nil
}
