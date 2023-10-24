package orm

import (
	"fmt"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"gorm.io/gorm"
)

const batch = 1000

// repairEvent this struct is temporarily used for repairing messenger.
type repairEvent struct {
	*TxHead
	Amount string
}

// RepairMessenger repair l1 and l2 messenger db due to 000019 .
func RepairMessenger(db *gorm.DB) error {
	if err := repairL1Messenger(db); err != nil {
		return err
	}
	if err := repairL2Messenger(db); err != nil {
		return err
	}
	return nil
}

func repairL1Messenger(client *ethclient.Client, db *gorm.DB) error {
	for {
		var l1Msgs []L1MessengerEvent
		rTx := db.Model(&L1MessengerEvent{}).
			Select("number", "tx_hash", "msg_hash", "amount").
			Where("msg_hash = ''").
			Order("number DESC").Limit(batch)
		err := rTx.Scan(&l1Msgs).Error
		if err != nil || len(l1Msgs) == 0 {
			return err
		}

		txHashes, err := fillL1Messenger(client, db, l1Msgs)
		if err != nil {
			return err
		}
		err = db.Transaction(func(tx *gorm.DB) error {
			wTx := tx.Model(&L1MessengerEvent{}).Select("msg_hash", "amount").Where("tx_hash IN ?", txHashes)
			return wTx.Updates(&l1Msgs).Error
		})
		if err != nil {
			return err
		}
	}
}

func fillL1Messenger(client *ethclient.Client, db *gorm.DB, l1Msgs []L1MessengerEvent) ([]string, error) {
	var (
		txHashes = make([]string, len(l1Msgs))
		msgs     = make(map[string]*L1MessengerEvent, len(l1Msgs))
	)
	for i := range l1Msgs {
		msg := &l1Msgs[i]
		txHashes[i] = msg.TxHash
		msgs[msg.TxHash] = msg
	}

	var ethEvents []repairEvent
	rTx := db.Model(&L1ETHEvent{}).Select("tx_hash", "msg_hash", "amount").Where("tx_hash IN ?", txHashes)
	err := rTx.Scan(&ethEvents).Error
	if err != nil {
		return nil, err
	}
	for _, event := range ethEvents {
		msg := msgs[event.TxHash]
		msg.MsgHash, msg.Amount = event.MsgHash, event.Amount
		delete(msgs, event.TxHash)
	}

	var erc20Event []repairEvent
	rTx = db.Model(&L1ERC20Event{}).Select("tx_hash", "msg_hash", "amount").Where("tx_hash IN ?", txHashes)
	err = rTx.Scan(&erc20Event).Error
	if err != nil {
		return nil, err
	}
	for _, event := range erc20Event {
		msg := msgs[event.TxHash]
		msg.MsgHash, msg.Amount = event.MsgHash, event.Amount
		delete(msgs, event.TxHash)
	}

	for _, event := range msgs {
		if event.MsgHash == "" {
			fmt.Println(event.TxHash)
		}
	}

	return txHashes, nil
}

func repairL2Messenger(db *gorm.DB) error {
	for {
		var l2Msgs []L2MessengerEvent
		rTx := db.Model(&L2MessengerEvent{}).
			Select("number", "msg_hash", "amount").
			Where("tx_hash = ''").
			Order("number DESC").Limit(batch)
		err := rTx.Scan(&l2Msgs).Error
		if err != nil || len(l2Msgs) == 0 {
			return err
		}

		msgHashes, err := fillL2Messenger(db, l2Msgs)
		if err != nil {
			return err
		}
		err = db.Transaction(func(tx *gorm.DB) error {
			wTx := tx.Model(&L2MessengerEvent{}).Select("tx_hash", "amount").Where("msg_hash IN ?", msgHashes)
			return wTx.Updates(&l2Msgs).Error
		})
		if err != nil {
			return err
		}
	}
}

func fillL2Messenger(db *gorm.DB, l2Msgs []L2MessengerEvent) ([]string, error) {
	var (
		msgHashes = make([]string, len(l2Msgs))
		msgs      = make(map[string]*L2MessengerEvent, len(l2Msgs))
	)
	for i := range l2Msgs {
		msg := &l2Msgs[i]
		msgHashes[i] = msg.MsgHash
		msgs[msg.MsgHash] = msg
	}

	var ethEvents []repairEvent
	rTx := db.Model(&L2ETHEvent{}).Select("tx_hash", "msg_hash", "amount").Where("tx_hash IN ?", msgHashes)
	err := rTx.Scan(&ethEvents).Error
	if err != nil {
		return nil, err
	}
	for _, event := range ethEvents {
		msgs[event.MsgHash].TxHash, msgs[event.MsgHash].Amount = event.TxHash, event.Amount
	}

	var erc20Event []repairEvent
	rTx = db.Model(&L2ERC20Event{}).Select("tx_hash", "msg_hash", "amount").Where("tx_hash IN ?", msgHashes)
	err = rTx.Scan(&erc20Event).Error
	if err != nil {
		return nil, err
	}
	for _, event := range erc20Event {
		msgs[event.MsgHash].TxHash, msgs[event.MsgHash].Amount = event.TxHash, event.Amount
	}
	return msgHashes, nil
}
