package orm

import (
	"math/big"

	"github.com/scroll-tech/go-ethereum/core/types"
)

// EventType represents the type of blockchain event.
type EventType uint8

var (
	// L1DepositETH represents the event for depositing ETH on Layer 1.
	L1DepositETH EventType = 1
	// L1FinalizeWithdrawETH represents the event for finalizing ETH withdrawal on Layer 1.
	L1FinalizeWithdrawETH EventType = 2
	// L2FinalizeDepositETH represents the event for finalizing ETH deposit on Layer 2.
	L2FinalizeDepositETH EventType = 101
	// L2WithdrawETH represents the event for withdrawing ETH on Layer 2.
	L2WithdrawETH EventType = 102

	// L1DepositWETH represents the event for depositing Wrapped ETH on Layer 1.
	L1DepositWETH EventType = 1
	// L1FinalizeWithdrawWETH represents the event for finalizing Wrapped ETH withdrawal on Layer 1.
	L1FinalizeWithdrawWETH EventType = 2
	// L1DepositDAI represents the event for depositing DAI on Layer 1.
	L1DepositDAI EventType = 3
	// L1FinalizeWithdrawDAI represents the event for finalizing DAI withdrawal on Layer 1.
	L1FinalizeWithdrawDAI EventType = 4
	// L1DepositStandardERC20 represents the event for depositing standard ERC20 tokens on Layer 1.
	L1DepositStandardERC20 EventType = 5
	// L1FinalizeWithdrawStandardERC20 represents the event for finalizing standard ERC20 token withdrawal on Layer 1.
	L1FinalizeWithdrawStandardERC20 EventType = 6
	// L1DepositCustomERC20 represents the event for depositing custom ERC20 tokens on Layer 1.
	L1DepositCustomERC20 EventType = 7
	// L1FinalizeWithdrawCustomERC20 represents the event for finalizing custom ERC20 token withdrawal on Layer 1.
	L1FinalizeWithdrawCustomERC20 EventType = 8
	// L2FinalizeDepositWETH represents the event for finalizing Wrapped ETH deposit on Layer 2.
	L2FinalizeDepositWETH EventType = 101
	// L2WithdrawWETH represents the event for withdrawing Wrapped ETH on Layer 2.
	L2WithdrawWETH EventType = 102
	// L2FinalizeDepositDAI represents the event for finalizing DAI deposit on Layer 2.
	L2FinalizeDepositDAI EventType = 103
	// L2WithdrawDAI represents the event for withdrawing DAI on Layer 2.
	L2WithdrawDAI EventType = 104
	// L2FinalizeDepositStandardERC20 represents the event for finalizing standard ERC20 token deposit on Layer 2.
	L2FinalizeDepositStandardERC20 EventType = 105
	// L2WithdrawStandardERC20 represents the event for withdrawing standard ERC20 tokens on Layer 2.
	L2WithdrawStandardERC20 EventType = 106
	// L2FinalizeDepositCustomERC20 represents the event for finalizing custom ERC20 token deposit on Layer 2.
	L2FinalizeDepositCustomERC20 EventType = 107
	// L2WithdrawCustomERC20 represents the event for withdrawing custom ERC20 tokens on Layer 2.
	L2WithdrawCustomERC20 EventType = 108

	// L1DepositERC721 represents the event for depositing ERC721 tokens on Layer 1.
	L1DepositERC721 EventType = 1
	// L1FinalizeWithdrawERC721 represents the event for finalizing ERC721 token withdrawal on Layer 1.
	L1FinalizeWithdrawERC721 EventType = 2
	// L2FinalizeDepositERC721 represents the event for finalizing ERC721 token deposit on Layer 2.
	L2FinalizeDepositERC721 EventType = 101
	// L2WithdrawERC721 represents the event for withdrawing ERC721 tokens on Layer 2.
	L2WithdrawERC721 EventType = 102

	// L1DepositERC1155 represents the event for depositing ERC1155 tokens on Layer 1.
	L1DepositERC1155 EventType = 1
	// L1FinalizeWithdrawERC1155 represents the event for finalizing ERC1155 token withdrawal on Layer 1.
	L1FinalizeWithdrawERC1155 EventType = 2
	// L2FinalizeDepositERC1155 represents the event for finalizing ERC1155 token deposit on Layer 2.
	L2FinalizeDepositERC1155 EventType = 101
	// L2WithdrawERC1155 represents the event for withdrawing ERC1155 tokens on Layer 2.
	L2WithdrawERC1155 EventType = 102
)

// TxHead represents the essential attributes of a transaction.
type TxHead struct {
	Number  uint64    `gorm:"index; comment: block number"`
	TxHash  string    `gorm:"primaryKey; type: varchar(66); comment: tx hash"`
	MsgHash string    `gorm:"index"`
	Type    EventType `gorm:"index; comment: tx type"`
}

// NewTxHead creates a new transaction header from a log and event type.
func NewTxHead(vLog *types.Log, tp EventType) *TxHead {
	return &TxHead{
		Number: vLog.BlockNumber,
		TxHash: vLog.TxHash.String(),
		Type:   tp,
	}
}

// L1ETHEvent represents an ETH event on Layer 1.
type L1ETHEvent struct {
	*TxHead
	// event message
	Amount *big.Int `gorm:"type:numeric(32,0)"`
}

// L2ETHEvent represents an ETH event on Layer 2.
type L2ETHEvent struct {
	*TxHead
	Amount *big.Int `gorm:"type:numeric(32,0)"`
}

// L1ERC20Event represents an ERC20 event on Layer 1.
type L1ERC20Event struct {
	*TxHead
	L1Token string   `gorm:"comment: input token contract address"`
	L2Token string   `gorm:"comment: output token contract address"`
	Amount  *big.Int `gorm:"type:numeric(32,0)"`
}

// L2ERC20Event represents an ERC20 event on Layer 2.
type L2ERC20Event struct {
	*TxHead
	L1Token string   `gorm:"comment: input token contract address"`
	L2Token string   `gorm:"comment: output token contract address"`
	Amount  *big.Int `gorm:"type:numeric(32,0)"`
}

// L1ERC721Event represents an ERC721 event on Layer 1.
type L1ERC721Event struct {
	*TxHead
	L1Token string
	L2Token string
	TokenID *big.Int `gorm:"type:serial"`
}

// L2ERC721Event represents an ERC721 event on Layer 2.
type L2ERC721Event struct {
	*TxHead
	L1Token string
	L2Token string
	TokenID *big.Int `gorm:"type:serial"`
}

// L1ERC1155Event represents an ERC1155 event on Layer 1.
type L1ERC1155Event struct {
	*TxHead
	L1Token string
	L2Token string
	TokenID *big.Int `gorm:"type:serial"`
	Amount  *big.Int `gorm:"type:numeric(32,0)"`
}

// L2ERC1155Event represents an ERC1155 event on Layer 2.
type L2ERC1155Event struct {
	*TxHead
	L1Token string
	L2Token string
	TokenID *big.Int `gorm:"type:serial"`
	Amount  *big.Int `gorm:"type:numeric(32,0)"`
}
