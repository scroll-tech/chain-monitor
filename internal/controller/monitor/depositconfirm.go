package monitor

import (
	"context"
	"fmt"

	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"

	"chain-monitor/internal/orm"
)

type msgEvents struct {
	L2Number uint64 `gorm:"l2_number"`

	L1TxHash string `gorm:"l1_tx_hash"`
	L2TxHash string `gorm:"l2_tx_hash"`

	// asset fields
	L1Amount  string `gorm:"l1_amount"`
	L2Amount  string `gorm:"l2_amount"`
	L1TokenID string `gorm:"l1_token_id"`
	L2TokenID string `gorm:"l2_token_id"`
}

func (ch *ChainMonitor) confirmDepositEvents(ctx context.Context, db *gorm.DB, start, end uint64) ([]uint64, error) {
	db = db.WithContext(ctx)

	var (
		failedNumbers []uint64
		flagNumbers   = map[uint64]bool{}
	)

	// check eth events.
	var ethEvents []msgEvents
	sql := `select 
    l1ee.tx_hash as l1_tx_hash, l1ee.amount as l1_amount, 
    l2ee.tx_hash as l2_tx_hash, l2ee.number as l2_number, l2ee.amount as l2_amount 
from l2_eth_events as l2ee full join l1_eth_events as l1ee 
    on l1ee.msg_hash = l2ee.msg_hash  
where l2ee.number BETWEEN ? AND ? and l2ee.type = ?;`
	db = db.Raw(sql, start, end, orm.L2FinalizeDepositETH)
	if err := db.Scan(&ethEvents).Error; err != nil {
		return nil, err
	}
	for i := 0; i < len(ethEvents); i++ {
		msg := ethEvents[i]
		if msg.L1Amount != msg.L2Amount {
			if !flagNumbers[msg.L2Number] {
				flagNumbers[msg.L2Number] = true
				failedNumbers = append(failedNumbers, msg.L2Number)
			}
			// If eth msg don't match, alert it.
			go ch.SlackNotify(fmt.Sprintf("deposit eth don't match, message: %v", msg))
			log.Error("the eth deposit hash or amount don't match", "l1_tx_hash", msg.L1TxHash, "l2_tx_hash", msg.L2TxHash)
		}
	}

	var erc20Events []msgEvents
	// check erc20 events.
	sql = `select 
    l1ee.tx_hash as l1_tx_hash, l1ee.amount as l1_amount, 
    l2ee.tx_hash as l2_tx_hash, l2ee.number as l2_number, l2ee.amount as l2_amount 
from l2_erc20_events as l2ee full join l1_erc20_events as l1ee 
    on l1ee.msg_hash = l2ee.msg_hash  
where l2ee.number BETWEEN ? AND ? and l2ee.type in (?, ?, ?, ?);`
	db = db.Raw(sql,
		start, end,
		orm.L2FinalizeDepositDAI,
		orm.L2FinalizeDepositWETH,
		orm.L2FinalizeDepositStandardERC20,
		orm.L2FinalizeDepositCustomERC20)
	if err := db.Scan(&erc20Events).Error; err != nil {
		return nil, err
	}
	for i := 0; i < len(erc20Events); i++ {
		msg := erc20Events[i]
		if msg.L1Amount != msg.L2Amount {
			if !flagNumbers[msg.L2Number] {
				flagNumbers[msg.L2Number] = true
				failedNumbers = append(failedNumbers, msg.L2Number)
			}
			// If erc20 msg don't match, alert it.
			go ch.SlackNotify(fmt.Sprintf("erc20 deposit don't match, message: %v", msg))
			log.Error("the erc20 deposit hash or amount doesn't match", "l1_tx_hash", msg.L1TxHash, "l2_tx_hash", msg.L2TxHash)
		}
	}

	// check erc721 events.
	var erc721Events []msgEvents
	sql = `select 
    l1ee.tx_hash as l1_tx_hash, l1ee.token_id as l1_token_id, 
    l2ee.tx_hash as l2_tx_hash, l2ee.number as l2_number, l2ee.token_id as l2_token_id
from l2_erc721_events as l2ee full join l1_erc721_events as l1ee 
    on l1ee.msg_hash = l2ee.msg_hash 
where l2ee.number BETWEEN ? AND ? and l2ee.type = ?;`
	db = db.Raw(sql, start, end, orm.L2FinalizeDepositERC721)
	if err := db.Scan(&erc721Events).Error; err != nil {
		return nil, err
	}
	for i := 0; i < len(erc721Events); i++ {
		msg := erc721Events[i]
		if msg.L1TokenID != msg.L2TokenID {
			if !flagNumbers[msg.L2Number] {
				flagNumbers[msg.L2Number] = true
				failedNumbers = append(failedNumbers, msg.L2Number)
			}
			// If erc721 event don't match, alert it.
			go ch.SlackNotify(fmt.Sprintf("erc721 event don't match, message: %v", msg))
			log.Error("the erc721 deposit hash or amount doesn't match", "l1_tx_hash", msg.L1TxHash, "l2_tx_hash", msg.L2TxHash)
		}
	}

	// check erc1155 events.
	var erc1155Events []msgEvents
	sql = `select 
    l1ee.tx_hash as l1_tx_hash, l1ee.amount as l1_amount, l1ee.token_id as l1_token_id, 
    l2ee.tx_hash as l2_tx_hash, l2ee.number as l2_number, l2ee.amount as l2_amount, l2ee.token_id as l2_token_id
from l2_erc1155_events as l2ee full join l1_erc1155_events as l1ee 
    on l1ee.msg_hash = l2ee.msg_hash 
where l2ee.number BETWEEN ? AND ? and l2ee.type = ?;`
	db = db.Raw(sql, start, end, orm.L2FinalizeDepositERC1155)
	if err := db.Scan(&erc1155Events).Error; err != nil {
		return nil, err
	}
	for i := 0; i < len(erc1155Events); i++ {
		msg := erc1155Events[i]
		if msg.L1TokenID != msg.L2TokenID || msg.L1Amount != msg.L2Amount {
			if !flagNumbers[msg.L2Number] {
				flagNumbers[msg.L2Number] = true
				failedNumbers = append(failedNumbers, msg.L2Number)
			}
			// If erc1155 event don't match, alert it.
			go ch.SlackNotify(fmt.Sprintf("erc1155 event don't match, message: %v", msg))
			log.Error("the erc1155 deposit hash or amount doesn't match", "l1_tx_hash", msg.L1TxHash, "l2_tx_hash", msg.L2TxHash)
		}
	}

	return failedNumbers, nil
}
