package monitor

import (
	"context"
	"math/big"

	"github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"

	"chain-monitor/bytecode/scroll/L1"
	"chain-monitor/bytecode/scroll/L2"
	"chain-monitor/internal/orm"
	"chain-monitor/internal/utils"
)

const batch = 500

// repairEvent this struct is temporarily used for repairing messenger.
type repairEvent struct {
	*orm.TxHead
	Amount string
}

// RepairMessenger repair l1 and l2 messenger db due to 000019 .
func (ch *ChainMonitor) RepairMessenger() error {
	if err := ch.repairL1Messenger(); err != nil {
		return err
	}
	if err := ch.repairL2Messenger(); err != nil {
		return err
	}
	return nil
}

func (ch *ChainMonitor) repairL1Messenger() error {
	for {
		var l1Msgs []orm.L1MessengerEvent
		rTx := ch.db.Model(&orm.L1MessengerEvent{}).
			Select("number", "tx_hash", "msg_hash", "amount", "type").
			Where("msg_hash = ''").
			Order("number DESC").Limit(batch)
		err := rTx.Scan(&l1Msgs).Error
		if err != nil || len(l1Msgs) == 0 {
			return err
		}

		log.Info("repairing l1_messenger_events", "start number", l1Msgs[len(l1Msgs)-1].Number, "end number", l1Msgs[0].Number)

		_, err = ch.fillL1Messenger(l1Msgs)
		if err != nil {
			return err
		}

		db := ch.db.Begin()
		for i := range l1Msgs {
			msg := &l1Msgs[i]
			wTx := db.Model(&orm.L1MessengerEvent{}).Select("msg_hash", "amount").Where("tx_hash = ?", msg.TxHash)
			err = wTx.Updates(msg).Error
			if err != nil {
				return err
			}
		}
		if err = db.Commit().Error; err != nil {
			return err
		}
	}
}

func (ch *ChainMonitor) fillL1Messenger(l1Msgs []orm.L1MessengerEvent) ([]string, error) {
	var (
		db        = ch.db
		client    = ethclient.NewClient(ch.l1watcher.RPCClient())
		contracts = ch.l1watcher.L1Contracts()
		txHashes  = make([]string, len(l1Msgs))
		msgs      = make(map[string]*orm.L1MessengerEvent, len(l1Msgs))
	)
	for i := range l1Msgs {
		msg := &l1Msgs[i]
		txHashes[i] = msg.TxHash
		msgs[msg.TxHash] = msg
	}

	var ethEvents []repairEvent
	rTx := db.Model(&orm.L1ETHEvent{}).Select("tx_hash", "msg_hash", "amount").Where("tx_hash IN ?", txHashes)
	err := rTx.Scan(&ethEvents).Error
	if err != nil {
		return nil, err
	}
	for _, event := range ethEvents {
		msg := msgs[event.TxHash]
		msg.MsgHash = event.MsgHash
		if msg.Type == orm.L1SentMessage {
			msg.Amount = event.Amount
		}
		delete(msgs, event.TxHash)
	}

	var erc20Event []repairEvent
	rTx = db.Model(&orm.L1ERC20Event{}).Select("tx_hash", "msg_hash", "amount").Where("tx_hash IN ?", txHashes)
	err = rTx.Scan(&erc20Event).Error
	if err != nil {
		return nil, err
	}
	for _, event := range erc20Event {
		msgs[event.TxHash].MsgHash = event.MsgHash
		delete(msgs, event.TxHash)
	}

	for _, msg := range msgs {
		msgBind := bind.NewBoundContract(contracts.ScrollMessenger, *L1.L1ScrollMessengerABI, nil, nil, nil)
		// get non gateway messenger events.
		filter := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(msg.Number)),
			ToBlock:   big.NewInt(int64(msg.Number)),
			Addresses: []common.Address{contracts.ScrollMessenger},
		}
		logs, err := client.FilterLogs(context.Background(), filter)
		if err != nil {
			return nil, err
		}
		for _, log := range logs {
			if msg.MsgHash != "" {
				break
			}
			if log.TxHash.String() != msg.TxHash {
				continue
			}
			if msg.Type == orm.L1SentMessage {
				event := new(L1.L1ScrollMessengerSentMessageEvent)
				if err = msgBind.UnpackLog(event, "SentMessage", log); err != nil {
					return nil, err
				}
				msgHash := utils.ComputeMessageHash(event.Sender, event.Target, event.Value, event.MessageNonce, event.Message)
				msg.MsgHash = msgHash.String()
				msg.Amount = event.Value.String()
			}
			if msg.Type == orm.L1RelayedMessage {
				event := new(L1.L1ScrollMessengerRelayedMessageEvent)
				if err = msgBind.UnpackLog(event, "RelayedMessage", log); err != nil {
					return nil, err
				}
				msg.MsgHash = common.BytesToHash(event.MessageHash[:]).String()
			}
			if msg.Type == orm.L1FailedRelayedMessage {
				event := new(L1.L1ScrollMessengerFailedRelayedMessageEvent)
				if err = msgBind.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
					return nil, err
				}
				msg.MsgHash = common.BytesToHash(event.MessageHash[:]).String()
			}
		}
	}

	return txHashes, nil
}

func (ch *ChainMonitor) repairL2Messenger() error {
	for {
		var l2Msgs []orm.L2MessengerEvent
		rTx := ch.db.Model(&orm.L2MessengerEvent{}).
			Select("number", "msg_hash", "amount", "type").
			Where("tx_hash = ''").
			Order("number DESC").Limit(batch)
		err := rTx.Scan(&l2Msgs).Error
		if err != nil || len(l2Msgs) == 0 {
			return err
		}

		log.Info("repairing l2_messenger_events", "start number", l2Msgs[len(l2Msgs)-1].Number, "end number", l2Msgs[0].Number)

		_, err = ch.fillL2Messenger(l2Msgs)
		if err != nil {
			return err
		}

		db := ch.db.Begin()
		for i := range l2Msgs {
			msg := &l2Msgs[i]
			wTx := db.Model(&orm.L2MessengerEvent{}).Select("tx_hash", "amount").Where("msg_hash = ?", msg.MsgHash)
			err = wTx.Updates(msg).Error
			if err != nil {
				return err
			}
		}
		if err = db.Commit().Error; err != nil {
			return err
		}
	}
}

func (ch *ChainMonitor) fillL2Messenger(l2Msgs []orm.L2MessengerEvent) ([]string, error) {
	var (
		db        = ch.db
		contracts = ch.l2watcher.L2Contracts()
		client    = ethclient.NewClient(ch.l2watcher.RPCClient())
		msgHashes = make([]string, len(l2Msgs))
		msgs      = make(map[string]*orm.L2MessengerEvent, len(l2Msgs))
	)
	for i := range l2Msgs {
		msg := &l2Msgs[i]
		msgHashes[i] = msg.MsgHash
		msgs[msg.MsgHash] = msg
	}

	var ethEvents []repairEvent
	rTx := db.Model(&orm.L2ETHEvent{}).Select("tx_hash", "msg_hash", "amount").Where("msg_hash IN ?", msgHashes)
	err := rTx.Scan(&ethEvents).Error
	if err != nil {
		return nil, err
	}
	for _, event := range ethEvents {
		msg := msgs[event.MsgHash]
		msg.TxHash = event.TxHash
		if msg.Type == orm.L2SentMessage {
			msg.Amount = event.Amount
		}
		delete(msgs, event.MsgHash)
	}

	var erc20Event []repairEvent
	rTx = db.Model(&orm.L2ERC20Event{}).Select("tx_hash", "msg_hash", "amount").Where("msg_hash IN ?", msgHashes)
	err = rTx.Scan(&erc20Event).Error
	if err != nil {
		return nil, err
	}
	for _, event := range erc20Event {
		msg := msgs[event.MsgHash]
		msg.TxHash = event.TxHash
		delete(msgs, event.MsgHash)
	}

	for _, msg := range msgs {
		abi := L2.L2ScrollMessengerABI
		msgBind := bind.NewBoundContract(contracts.ScrollMessenger, *abi, nil, nil, nil)
		// get non gateway messenger events.
		filter := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(msg.Number)),
			ToBlock:   big.NewInt(int64(msg.Number)),
			Addresses: []common.Address{contracts.ScrollMessenger},
		}
		logs, err := client.FilterLogs(context.Background(), filter)
		if err != nil {
			return nil, err
		}
		for _, vLog := range logs {
			if msg.TxHash != "" {
				break
			}
			switch vLog.Topics[0] {
			case abi.Events["SentMessage"].ID:
				if msg.Type != orm.L2SentMessage {
					continue
				}
				event := new(L2.L2ScrollMessengerSentMessageEvent)
				if err = msgBind.UnpackLog(event, "SentMessage", vLog); err != nil {
					return nil, err
				}
				msgHash := utils.ComputeMessageHash(event.Sender, event.Target, event.Value, event.MessageNonce, event.Message)
				if msgHash.String() == msg.MsgHash {
					msg.TxHash = vLog.TxHash.String()
					msg.Amount = event.Value.String()
				}
			case abi.Events["RelayedMessage"].ID:
				if msg.Type != orm.L2RelayedMessage {
					continue
				}
				event := new(L2.L2ScrollMessengerRelayedMessageEvent)
				if err = msgBind.UnpackLog(event, "RelayedMessage", vLog); err != nil {
					return nil, err
				}
				msgHash := common.BytesToHash(event.MessageHash[:]).String()
				if msgHash == msg.MsgHash {
					msg.TxHash = vLog.TxHash.String()
				}
			case abi.Events["FailedRelayedMessage"].ID:
				if msg.Type != orm.L2FailedRelayedMessage {
					continue
				}
				event := new(L1.L1ScrollMessengerFailedRelayedMessageEvent)
				if err = msgBind.UnpackLog(event, "FailedRelayedMessage", vLog); err != nil {
					return nil, err
				}
				msgHash := common.BytesToHash(event.MessageHash[:]).String()
				if msgHash == msg.MsgHash {
					msg.TxHash = vLog.TxHash.String()
				}
			}
		}
	}

	return msgHashes, nil
}
