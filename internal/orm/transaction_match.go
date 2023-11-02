package orm

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// TransactionMatch contains the tx of l1 & l2
type TransactionMatch struct {
	db *gorm.DB `gorm:"column:-"`

	ID          int64  `json:"id" gorm:"column:id"`
	MessageHash string `json:"messageHash" gorm:"message_hash"`
	TokenType   int    `json:"token_type" gorm:"token_type"`

	// tx info of l1
	L1EventType   int             `json:"l1_event_type" gorm:"l1_event_type"`
	L1BlockNumber uint64          `json:"l1_block_number" gorm:"l1_block_number"`
	L1TxHash      string          `json:"l1_tx_hash" gorm:"l1_tx_hash"`
	L1TokenId     string          `json:"l1_token_id" gorm:"l1_token_id"`
	L1Value       decimal.Decimal `json:"l1_value" gorm:"l1_value"`

	// tx info of l2
	L2EventType   int             `json:"l2_event_type" gorm:"l2_event_type"`
	L2BlockNumber uint64          `json:"l2_block_number" gorm:"l2_block_number"`
	L2TxHash      string          `json:"l2_tx_hash" gorm:"l2_tx_hash"`
	L2TokenId     string          `json:"l2_token_id" gorm:"l2_token_id"`
	L2Value       decimal.Decimal `json:"l2_value" gorm:"l2_value"`

	// status
	WithdrawRootStatus int    `json:"withdraw_root_status" gorm:"withdraw_root_status"`
	DepositStatus      int    `json:"deposit_status" gorm:"deposit_status"`
	BalanceStatus      int    `json:"balance_status" gorm:"balance_status"`
	MessageProof       string `json:"msg_proof" gorm:"message_proof"`

	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

// NewTransactionMatch creates a new TransactionMatch database instance.
func NewTransactionMatch(db *gorm.DB) *TransactionMatch {
	return &TransactionMatch{db: db}
}

// TableName returns the table name for the Batch model.
func (*TransactionMatch) TableName() string {
	return "transaction_match"
}

func (t *TransactionMatch) InsertOrUpdate(ctx context.Context, transactions []TransactionMatch) (int, error) {
	// insert or update
}
