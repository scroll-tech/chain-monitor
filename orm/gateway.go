package orm

import (
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"gorm.io/gorm"
	"math/big"
)

type EventType uint8

var (
	L1DepositETH          EventType = 1
	L1FinalizeWithdrawETH EventType = 2
	L2FinalizeDepositETH  EventType = 101
	L2WithdrawETH         EventType = 102

	L1DepositERC20                  EventType = 1
	L1FinalizeWithdrawERC20         EventType = 2
	L1DepositWETH                   EventType = 3
	L1FinalizeWithdrawWETH          EventType = 4
	L1DepositDAI                    EventType = 5
	L1FinalizeWithdrawDAI           EventType = 6
	L1DepositStandardERC20          EventType = 7
	L1FinalizeWithdrawStandardERC20 EventType = 8

	L2FinalizeDepositERC20         EventType = 101
	L2WithdrawERC20                EventType = 102
	L2FinalizeDepositWETH          EventType = 103
	L2WithdrawWETH                 EventType = 104
	L2FinalizeDepositDAI           EventType = 105
	L2WithdrawDAI                  EventType = 106
	L2FinalizeDepositStandardERC20 EventType = 107
	L2WithdrawStandardERC20        EventType = 108

	L1DepositCustomERC20          EventType = 9
	L1FinalizeWithdrawCustomERC20 EventType = 10
	L2FinalizeDepositCustomERC20  EventType = 109
	L2WithdrawCustomERC20         EventType = 110

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
	Number uint64    `gorm:"index; comment: block number"`
	TxHash string    `gorm:"primaryKey; type: varchar(66); comment: tx hash"`
	Type   EventType `gorm:"index; comment: tx type"`
}

func newTxHead(vLog *types.Log, tp EventType) *TxHead {
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

func SaveL1ETHEvent(db *gorm.DB, eventType EventType, vLog *types.Log, from, to common.Address, amount *big.Int) error {
	return db.Save(&L1ETHEvent{
		TxHead: newTxHead(vLog, eventType),
		From:   from.String(),
		To:     to.String(),
		Amount: amount,
	}).Error
}

func SaveL2ETHEvent(db *gorm.DB, eventType EventType, vLog *types.Log, from, to common.Address, amount *big.Int) error {
	return db.Save(&L2ETHEvent{
		TxHead: newTxHead(vLog, eventType),
		From:   from.String(),
		To:     to.String(),
		Amount: amount,
	}).Error
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

func SaveL1ETH20Event(db *gorm.DB, eventType EventType, vLog *types.Log, l1Token, l2Token, from, to common.Address, amount *big.Int) error {
	return db.Save(&L1ERC20Event{
		TxHead:  newTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		From:    from.String(),
		To:      to.String(),
		Amount:  amount,
	}).Error
}

func SaveL2ETH20Event(db *gorm.DB, eventType EventType, vLog *types.Log, l1Token, l2Token, from, to common.Address, amount *big.Int) error {
	return db.Save(&L2ERC20Event{
		TxHead:  newTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		From:    from.String(),
		To:      to.String(),
		Amount:  amount,
	}).Error
}

type L1ERC721Event struct {
	*TxHead
	L1Token string
	L2Token string
	From    string
	To      string
	TokenId *big.Int `gorm:"type:serial"`
}

func SaveL1ERC721Event(db *gorm.DB, eventType EventType, vLog *types.Log, l1Token, l2Token, from, to common.Address, tokenId *big.Int) error {
	return db.Save(&L1ERC721Event{
		TxHead:  newTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		From:    from.String(),
		To:      to.String(),
		TokenId: tokenId,
	}).Error
}

type L2ERC721Event struct {
	*TxHead
	L1Token string
	L2Token string
	From    string
	To      string
	TokenId *big.Int `gorm:"type:serial"`
}

func SaveL2ERC721Event(db *gorm.DB, eventType EventType, vLog *types.Log, l1Token, l2Token, from, to common.Address, tokenId *big.Int) error {
	return db.Save(&L1ERC721Event{
		TxHead:  newTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		From:    from.String(),
		To:      to.String(),
		TokenId: tokenId,
	}).Error
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

func SaveL1ERC1155Event(db *gorm.DB, eventType EventType, vLog *types.Log, l1Token, l2Token, from, to common.Address, tokenId, amount *big.Int) error {
	return db.Save(&L1ERC1155Event{
		TxHead:  newTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		From:    from.String(),
		To:      to.String(),
		TokenId: tokenId,
		Amount:  amount,
	}).Error
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

func SaveL2ERC1155Event(db *gorm.DB, eventType EventType, vLog *types.Log, l1Token, l2Token, from, to common.Address, tokenId, amount *big.Int) error {
	return db.Save(&L2ERC1155Event{
		TxHead:  newTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		From:    from.String(),
		To:      to.String(),
		TokenId: tokenId,
		Amount:  amount,
	}).Error
}
