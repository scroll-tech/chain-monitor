package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"
	"modernc.org/mathutil"

	"chain-monitor/internal/controller"
	"chain-monitor/internal/orm"
)

var (
	l2ethSQL = `select 
    l1ee.tx_hash as l1_tx_hash, l1ee.amount as l1_amount, 
    l2ee.tx_hash as l2_tx_hash, l2ee.number as l2_number, l2ee.amount as l2_amount 
from l2_eth_events as l2ee full join l1_eth_events as l1ee 
    on l1ee.msg_hash = l2ee.msg_hash  
where l2ee.number BETWEEN ? AND ? and l2ee.type = ?;`

	l2erc20SQL = `select 
    l1ee.tx_hash as l1_tx_hash, l1ee.amount as l1_amount, 
    l2ee.tx_hash as l2_tx_hash, l2ee.number as l2_number, l2ee.type as l2_type, l2ee.amount as l2_amount 
from l2_erc20_events as l2ee full join l1_erc20_events as l1ee 
    on l1ee.msg_hash = l2ee.msg_hash  
where l2ee.number BETWEEN ? AND ? and l2ee.type in (?, ?, ?, ?);`

	l2erc721SQL = `select 
    l1ee.tx_hash as l1_tx_hash, l1ee.token_id as l1_token_id, 
    l2ee.tx_hash as l2_tx_hash, l2ee.number as l2_number, l2ee.token_id as l2_token_id
from l2_erc721_events as l2ee full join l1_erc721_events as l1ee 
    on l1ee.msg_hash = l2ee.msg_hash 
where l2ee.number BETWEEN ? AND ? and l2ee.type = ?;`

	l2erc1155SQL = `select 
    l1ee.tx_hash as l1_tx_hash, l1ee.amount as l1_amount, l1ee.token_id as l1_token_id, 
    l2ee.tx_hash as l2_tx_hash, l2ee.number as l2_number, l2ee.amount as l2_amount, l2ee.token_id as l2_token_id
from l2_erc1155_events as l2ee full join l1_erc1155_events as l1ee 
    on l1ee.msg_hash = l2ee.msg_hash 
where l2ee.number BETWEEN ? AND ? and l2ee.type = ?;`

	l2NoGatewaySQL = `select
    l1me.tx_hash as l1_tx_hash, l1me.number as l1_number,
    l2me.tx_hash as l2_tx_hash, l2me.number as l2_number 
from l2_messenger_events as l2me full join l1_messenger_events as l1me 
    on l2me.msg_hash = l1me.msg_hash
where l2me.number BETWEEN ? AND ? and l2me.type = ? and l2me.from_gateway = false;`
)

// DepositConfirm monitors the blockchain and confirms the deposit events.
func (ch *ChainMonitor) DepositConfirm(ctx context.Context) {
	// Make sure the l1Watcher is ready to use.
	if !ch.l1watcher.IsReady() {
		log.Debug("l1watcher is not ready, sleep 3 seconds")
		time.Sleep(time.Second * 3)
		return
	}
	start, end := ch.getDepositStartAndEndNumber()
	if end > ch.depositSafeNumber {
		log.Debug("l2watcher is not ready", "l2_start_number", ch.depositSafeNumber)
		time.Sleep(time.Second * 3)
		return
	}
	// Get unmatched deposit
	failedNumbers, err := ch.confirmDepositEvents(ctx, start, end)
	if err != nil {
		log.Error("failed to confirm deposit events", "start", start, "end", end, "err", err)
		return
	}
	err = ch.db.Transaction(func(tx *gorm.DB) error {
		// Update deposit records.
		sTx := tx.Model(&orm.L2ChainConfirm{}).Select("deposit_status", "confirm").
			Where("number BETWEEN ? AND ?", start, end)
		sTx = sTx.Update("deposit_status", true).Update("confirm", true)
		if sTx.Error != nil {
			return sTx.Error
		}

		if len(failedNumbers) > 0 {
			fTx := tx.Model(&orm.L2ChainConfirm{}).Select("deposit_status").
				Where("number in ?", failedNumbers)
			fTx = fTx.Update("deposit_status", false)
			if fTx.Error != nil {
				return fTx.Error
			}
		}
		return nil
	})
	if err != nil {
		log.Error("failed to check deposit events", "start", start, "end", end, "err", err)
		time.Sleep(time.Second * 5)
		return
	}
	ch.depositStartNumber = end

	// Metrics records current goroutine.
	controller.WorkerStartedTotal.WithLabelValues("deposit_confirm").Inc()

	log.Info("confirm layer2 deposit transactions", "start", start, "end", end)
}

func (ch *ChainMonitor) confirmDepositEvents(ctx context.Context, start, end uint64) ([]uint64, error) {
	var (
		db            = ch.db.WithContext(ctx)
		failedNumbers []uint64
		flagNumbers   = map[uint64]bool{}
	)

	// check eth events.
	var ethEvents []msgEvents
	db = db.Raw(l2ethSQL, start, end, orm.L2FinalizeDepositETH)
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
			controller.DepositFailedTotal.WithLabelValues(orm.L2FinalizeDepositETH.String()).Inc()
			// If eth msg don't match, alert it.
			data, _ := json.Marshal(msg)
			go controller.SlackNotify(fmt.Sprintf("deposit eth don't match, message: %s", string(data)))
			log.Error("the eth deposit count or amount don't match", "start", start, "end", end, "event_type", orm.L2FinalizeDepositETH, "l1_tx_hash", msg.L1TxHash, "l2_tx_hash", msg.L2TxHash)
		}
	}

	var erc20Events []msgEvents
	// check erc20 events.
	db = db.Raw(l2erc20SQL,
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
			controller.DepositFailedTotal.WithLabelValues(msg.L2Type.String()).Inc()
			// If erc20 msg don't match, alert it.
			data, _ := json.Marshal(msg)
			go controller.SlackNotify(fmt.Sprintf("erc20 deposit don't match, message: %s", string(data)))
			log.Error(
				"the erc20 deposit count or amount doesn't match",
				"start", start,
				"end", end,
				"event_type", []orm.EventType{orm.L2FinalizeDepositDAI, orm.L2FinalizeDepositWETH, orm.L2FinalizeDepositStandardERC20, orm.L2FinalizeDepositCustomERC20},
				"l1_tx_hash", msg.L1TxHash,
				"l2_tx_hash", msg.L2TxHash,
			)
		}
	}

	// check erc721 events.
	var erc721Events []msgEvents
	db = db.Raw(l2erc721SQL, start, end, orm.L2FinalizeDepositERC721)
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
			controller.DepositFailedTotal.WithLabelValues(orm.L2FinalizeDepositERC721.String()).Inc()
			// If erc721 event don't match, alert it.
			data, _ := json.Marshal(msg)
			go controller.SlackNotify(fmt.Sprintf("erc721 event don't match, message: %s", string(data)))
			log.Error("the erc721 deposit count or amount doesn't match", "start", start, "end", end, "event_type", orm.L2FinalizeDepositERC721, "l1_tx_hash", msg.L1TxHash, "l2_tx_hash", msg.L2TxHash)
		}
	}

	// check erc1155 events.
	var erc1155Events []msgEvents
	db = db.Raw(l2erc1155SQL, start, end, orm.L2FinalizeDepositERC1155)
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
			data, _ := json.Marshal(msg)
			go controller.SlackNotify(fmt.Sprintf("erc1155 event don't match, message: %s", string(data)))
			log.Error("the erc1155 deposit count or amount doesn't match", "start", start, "end", end, "event_type", orm.L2FinalizeDepositERC1155, "l1_tx_hash", msg.L1TxHash, "l2_tx_hash", msg.L2TxHash)
		}
	}

	// check no gateway events.
	var noGateways []msgEvents
	db = db.Raw(l2NoGatewaySQL, start, end, orm.L2SentMessage)
	if err := db.Scan(noGateways).Error; err != nil {
		return nil, err
	}
	for i := 0; i < len(noGateways); i++ {
		msg := noGateways[i]
		if msg.L1Number == 0 || msg.L2Number == 0 || msg.L1Type == orm.L1FailedRelayedMessage {
			if msg.L1Type != orm.L1FailedRelayedMessage && !flagNumbers[msg.L2Number] {
				flagNumbers[msg.L2Number] = true
				failedNumbers = append(failedNumbers, msg.L2Number)
			}
			// no gateway event don't match, alert it.
			data, _ := json.Marshal(msg)
			go controller.SlackNotify(fmt.Sprintf("l2chain's sentMessage event can't match l1chain relayMessage event, content: %s", string(data)))
			log.Error("l2chain's sentMessage event can't match l1chain relayMessage event", "start", start, "end", end, "l1_tx_hash", msg.L1TxHash, "l2_tx_hash", msg.L2TxHash)
		}
	}

	return failedNumbers, nil
}

func (ch *ChainMonitor) getDepositStartAndEndNumber() (uint64, uint64) {
	ch.depositSafeNumber = ch.l2watcher.CurrentNumber()
	var (
		start = ch.depositStartNumber + 1
		end   = mathutil.MinUint64(start+batchSize-1, ch.depositSafeNumber)
	)
	if start <= end {
		return start, end
	}
	return start, start
}
