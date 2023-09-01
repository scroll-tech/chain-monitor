package orm

import (
	"math/big"

	"github.com/scroll-tech/go-ethereum/core/types"
)

type EventType uint8

var (
	L1DepositETH          EventType = 1
	L1FinalizeWithdrawETH EventType = 2
	L2FinalizeDepositETH  EventType = 101
	L2WithdrawETH         EventType = 102

	L1DepositWETH                   EventType = 1
	L1FinalizeWithdrawWETH          EventType = 2
	L1DepositDAI                    EventType = 3
	L1FinalizeWithdrawDAI           EventType = 4
	L1DepositStandardERC20          EventType = 5
	L1FinalizeWithdrawStandardERC20 EventType = 6
	L1DepositCustomERC20            EventType = 7
	L1FinalizeWithdrawCustomERC20   EventType = 8
	L2FinalizeDepositWETH           EventType = 101
	L2WithdrawWETH                  EventType = 102
	L2FinalizeDepositDAI            EventType = 103
	L2WithdrawDAI                   EventType = 104
	L2FinalizeDepositStandardERC20  EventType = 105
	L2WithdrawStandardERC20         EventType = 106
	L2FinalizeDepositCustomERC20    EventType = 107
	L2WithdrawCustomERC20           EventType = 108

	L1DepositERC721          EventType = 1
	L1FinalizeWithdrawERC721 EventType = 2
	L2FinalizeDepositERC721  EventType = 101
	L2WithdrawERC721         EventType = 102

	L1DepositERC1155          EventType = 1
	L1FinalizeWithdrawERC1155 EventType = 2
	L2FinalizeDepositERC1155  EventType = 101
	L2WithdrawERC1155         EventType = 102
)

type TxHead struct {
	Number  uint64    `gorm:"index; comment: block number"`
	TxHash  string    `gorm:"primaryKey; type: varchar(66); comment: tx hash"`
	MsgHash string    `gorm:"index"`
	Type    EventType `gorm:"index; comment: tx type"`
}

func NewTxHead(vLog *types.Log, tp EventType) *TxHead {
	return &TxHead{
		Number: vLog.BlockNumber,
		TxHash: vLog.TxHash.String(),
		Type:   tp,
	}
}

type L1ETHEvent struct {
	*TxHead
	// event message
	From   string
	To     string
	Amount *big.Int `gorm:"type:numeric(32,0)"`
}

type L2ETHEvent struct {
	*TxHead
	// event message
	From   string
	To     string
	Amount *big.Int `gorm:"type:numeric(32,0)"`
}

type L1ERC20Event struct {
	*TxHead
	// event message
	L1Token string `gorm:"comment: input token contract address"`
	L2Token string `gorm:"comment: output token contract address"`
	From    string
	To      string
	Amount  *big.Int `gorm:"type:numeric(32,0)"`
}

type L2ERC20Event struct {
	*TxHead
	// event message
	L1Token string `gorm:"comment: input token contract address"`
	L2Token string `gorm:"comment: output token contract address"`
	From    string
	To      string
	Amount  *big.Int `gorm:"type:numeric(32,0)"`
}

type L1ERC721Event struct {
	*TxHead
	L1Token string
	L2Token string
	From    string
	To      string
	TokenId *big.Int `gorm:"type:serial"`
}

type L2ERC721Event struct {
	*TxHead
	L1Token string
	L2Token string
	From    string
	To      string
	TokenId *big.Int `gorm:"type:serial"`
}

type L1ERC1155Event struct {
	*TxHead
	L1Token string
	L2Token string
	From    string
	To      string
	TokenId *big.Int `gorm:"type:serial"`
	Amount  *big.Int `gorm:"type:numeric(32,0)"`
}

type L2ERC1155Event struct {
	*TxHead
	L1Token string
	L2Token string
	From    string
	To      string
	TokenId *big.Int `gorm:"type:serial"`
	Amount  *big.Int `gorm:"type:numeric(32,0)"`
}
