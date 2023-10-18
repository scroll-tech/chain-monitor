package l1watcher

import (
	"context"
	"encoding/json"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/log"

	"chain-monitor/bytecode"
	"chain-monitor/bytecode/scroll/L1"
	"chain-monitor/bytecode/scroll/L1/rollup"
	l2gateway "chain-monitor/bytecode/scroll/L2/gateway"
	"chain-monitor/internal/orm"
	"chain-monitor/internal/utils"
)

func (l1 *l1Contracts) registerMessengerHandlers() {
	l1.ScrollMessenger.RegisterSentMessage(func(vLog *types.Log, data *L1.L1ScrollMessengerSentMessageEvent) error {
		msgHash := utils.ComputeMessageHash(data.Sender, data.Target, data.Value, data.MessageNonce, data.Message)
		l1.msgSentEvents[vLog.TxHash.String()] = &orm.L1MessengerEvent{
			TxHead: &orm.TxHead{
				Number:  vLog.BlockNumber,
				TxHash:  vLog.TxHash.String(),
				MsgHash: msgHash.String(),
				Type:    orm.L1SentMessage,
			},
			Value:   data.Value,
			Target:  data.Target,
			Message: data.Message,
			Log:     vLog,
		}
		return nil
	})
	l1.ScrollMessenger.RegisterRelayedMessage(func(vLog *types.Log, data *L1.L1ScrollMessengerRelayedMessageEvent) error {
		msgHash := common.BytesToHash(data.MessageHash[:])
		l1.msgSentEvents[vLog.TxHash.String()] = &orm.L1MessengerEvent{
			TxHead: &orm.TxHead{
				Number:  vLog.BlockNumber,
				TxHash:  vLog.TxHash.String(),
				MsgHash: msgHash.String(),
				Type:    orm.L1RelayedMessage,
			},
			Log: vLog,
		}
		return nil
	})
	l1.ScrollMessenger.RegisterFailedRelayedMessage(func(vLog *types.Log, data *L1.L1ScrollMessengerFailedRelayedMessageEvent) error {
		msgHash := common.BytesToHash(data.MessageHash[:])
		return orm.SaveL1Messenger(l1.tx, orm.L1FailedRelayedMessage, vLog, msgHash)
	})
}

func (l1 *l1Contracts) registerScrollHandlers() {
	l1.ScrollChain.RegisterCommitBatch(func(vLog *types.Log, data *rollup.ScrollChainCommitBatchEvent) error {
		l1Tx, _, err := l1.client.TransactionByHash(context.Background(), vLog.TxHash)
		if err != nil {
			return err
		}
		scrollBatch, err := bytecode.GetBatchRangeFromCalldataV2(l1.ScrollChain.ABI, l1Tx.Data())
		if err != nil {
			return err
		}
		return l1.tx.Create(&orm.L1ScrollChainEvent{
			Number:        vLog.BlockNumber,
			BatchIndex:    data.BatchIndex.Uint64(),
			BatchHash:     common.BytesToHash(data.BatchHash[:]).String(),
			CommitNumber:  vLog.BlockNumber,
			CommitHash:    vLog.TxHash.String(),
			L2StartNumber: scrollBatch.L2StartNumber,
			L2EndNumber:   scrollBatch.L2EndNumber,
			Status:        orm.BatchCommit,
		}).Error
	})
	l1.ScrollChain.RegisterFinalizeBatch(func(vLog *types.Log, data *rollup.ScrollChainFinalizeBatchEvent) error {
		eventMsg := orm.L1ScrollChainEvent{
			BatchHash:      common.BytesToHash(data.BatchHash[:]).String(),
			FinalizeNumber: vLog.BlockNumber,
			FinalizeHash:   vLog.TxHash.String(),
			Status:         orm.BatchFinalize,
		}
		return l1.tx.Select("finalize_number", "finalize_hash", "status").Updates(&eventMsg).Error
	})
	l1.ScrollChain.RegisterRevertBatch(func(vLog *types.Log, data *rollup.ScrollChainRevertBatchEvent) error {
		eventMsg := orm.L1ScrollChainEvent{
			BatchHash:    common.BytesToHash(data.BatchHash[:]).String(),
			RevertNumber: vLog.BlockNumber,
			RevertHash:   vLog.TxHash.String(),
			Status:       orm.BatchRevert,
		}
		return l1.tx.Select("revert_number", "revert_hash", "status").Updates(&eventMsg).Error
	})
}

func (l1 *l1Contracts) parseGatewayDeposit(l1msg *orm.L1MessengerEvent) error {
	if len(l1msg.Message) < 4 {
		log.Warn("l1chain sendMessage content less than 4 bytes", "tx_hash", l1msg.TxHash)
		return errMessenger
	}
	id := common.Bytes2Hex(l1msg.Message[:4])
	switch id {
	case "232e8748": // FinalizeDepositETH
		return l1.parseGatewayDepositETH(l1msg.TxHead, l1msg.Message)
	case "8431f5c1": // FinalizeDepositERC20
		return l1.parseGatewayDepositERC20(l1msg.TxHead, l1msg.Target, l1msg.Message)
	case "982b151f": // FinalizeBatchDepositERC721
		return l1.parseGatewayDepositERC721(l1msg.TxHead, l1msg.Message)
	case "4764cc62": // FinalizeDepositERC1155
		return l1.parseGatewayDepositERC1155(l1msg.TxHead, l1msg.Message)
	}
	log.Warn("l1chain sendMessage unexpect method_id", "tx_hash", l1msg.TxHash, "method_id", id)
	return errMessenger
}

func (l1 *l1Contracts) parseGatewayDepositETH(txHead *orm.TxHead, data []byte) error {
	method, err := l2gateway.L2ETHGatewayABI.MethodById(data)
	if err != nil {
		return err
	}
	params, err := method.Inputs.Unpack(data[4:])
	if err != nil {
		return err
	}
	event := new(l2gateway.L2ETHGatewayFinalizeDepositETHEvent)
	err = method.Inputs.Copy(event, params)
	if err != nil {
		return err
	}
	l1.ethEvents = append(l1.ethEvents, &orm.L1ETHEvent{
		TxHead: &orm.TxHead{
			Number:  txHead.Number,
			TxHash:  txHead.TxHash,
			MsgHash: txHead.MsgHash,
			Type:    orm.L1DepositETH,
		},
		Amount: event.Amount,
	})
	return nil
}

func (l1 *l1Contracts) parseGatewayDepositERC20(txHead *orm.TxHead, target common.Address, data []byte) error {
	method, err := l2gateway.L2ERC20GatewayABI.MethodById(data)
	if err != nil {
		return err
	}
	params, err := method.Inputs.Unpack(data[4:])
	if err != nil {
		return err
	}
	event := new(l2gateway.L2ERC20GatewayFinalizeDepositERC20Event)
	err = method.Inputs.Copy(event, params)
	if err != nil {
		return err
	}

	var (
		_tp        orm.EventType
		gatewayCfg = l1.l2gatewayCfg
	)
	switch target {
	case gatewayCfg.DAIGateway:
		_tp = orm.L1DepositDAI
	case gatewayCfg.WETHGateway:
		_tp = orm.L1DepositWETH
	case gatewayCfg.StandardERC20Gateway:
		_tp = orm.L1DepositStandardERC20
	case gatewayCfg.CustomERC20Gateway:
		_tp = orm.L1DepositCustomERC20
	case gatewayCfg.USDCGateway:
		_tp = orm.L1USDCDepositERC20
	case gatewayCfg.LIDOGateway:
		_tp = orm.L1LIDODepositERC20
	default:
		buf, _ := json.Marshal(event)
		log.Warn("l1chain unexpect erc20 deposit transfer", "target address", target.String(), "content", string(buf))
	}

	l1.erc20Events = append(l1.erc20Events, &orm.L1ERC20Event{
		TxHead: &orm.TxHead{
			Number:  txHead.Number,
			TxHash:  txHead.TxHash,
			MsgHash: txHead.MsgHash,
			Type:    _tp,
		},
		L1Token: event.L1Token.String(),
		L2Token: event.L2Token.String(),
		Amount:  event.Amount,
	})

	return nil
}

func (l1 *l1Contracts) parseGatewayDepositERC721(txHead *orm.TxHead, data []byte) error {
	method, err := l2gateway.L2ERC721GatewayABI.MethodById(data)
	if err != nil {
		return err
	}
	params, err := method.Inputs.Unpack(data[4:])
	if err != nil {
		return err
	}
	event := new(l2gateway.L2ERC721GatewayFinalizeDepositERC721Event)
	err = method.Inputs.Copy(event, params)
	if err != nil {
		return err
	}
	l1.erc721Events = append(l1.erc721Events, &orm.L1ERC721Event{
		TxHead: &orm.TxHead{
			Number:  txHead.Number,
			TxHash:  txHead.TxHash,
			MsgHash: txHead.MsgHash,
			Type:    orm.L1DepositERC721,
		},
		L1Token: event.L1Token.String(),
		L2Token: event.L2Token.String(),
		TokenID: event.TokenId,
	})

	return nil
}

func (l1 *l1Contracts) parseGatewayDepositERC1155(txHead *orm.TxHead, data []byte) error {
	method, err := l2gateway.L2ERC1155GatewayABI.MethodById(data)
	if err != nil {
		return err
	}
	params, err := method.Inputs.Unpack(data[4:])
	if err != nil {
		return err
	}
	event := new(l2gateway.L2ERC1155GatewayFinalizeDepositERC1155Event)
	err = method.Inputs.Copy(event, params)
	if err != nil {
		return err
	}
	l1.erc1155Events = append(l1.erc1155Events, &orm.L1ERC1155Event{
		TxHead: &orm.TxHead{
			Number:  txHead.Number,
			TxHash:  txHead.TxHash,
			MsgHash: txHead.MsgHash,
			Type:    orm.L1DepositERC1155,
		},
		L1Token: event.L1Token.String(),
		L2Token: event.L2Token.String(),
		TokenID: event.TokenId,
		Amount:  event.Amount,
	})
	return nil
}

func (l1 *l1Contracts) storeL1WatcherEvents() error {
	// store l1 eth events.
	if len(l1.ethEvents) > 0 {
		if err := l1.tx.Save(l1.ethEvents).Error; err != nil {
			return err
		}
	}

	// store l1 erc20 events.
	if len(l1.erc20Events) > 0 {
		if err := l1.tx.Save(l1.erc20Events).Error; err != nil {
			return err
		}
	}

	// store l1 err721 events.
	if len(l1.erc721Events) > 0 {
		if err := l1.tx.Save(l1.erc721Events).Error; err != nil {
			return err
		}
	}

	// store l1 erc1155 events.
	if len(l1.erc1155Events) > 0 {
		if err := l1.tx.Save(l1.erc1155Events).Error; err != nil {
			return err
		}
	}

	// store l1 scroll_messenger sentMessage events.
	if length := len(l1.msgSentEvents); length > 0 {
		var msgs = make([]*orm.L1MessengerEvent, 0, length)
		for k := range l1.msgSentEvents {
			msgs = append(msgs, l1.msgSentEvents[k])
		}
		if err := l1.tx.Save(msgs).Error; err != nil {
			return err
		}
	}

	return nil
}
