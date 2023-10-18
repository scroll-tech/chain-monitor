package l2watcher

import (
	l1gateway "chain-monitor/bytecode/scroll/L1/gateway"
	"context"
	"encoding/json"
	"fmt"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/log"

	"chain-monitor/bytecode/scroll/L2"
	"chain-monitor/internal/controller"
	"chain-monitor/internal/orm"
	"chain-monitor/internal/utils"
)

func (l2 *l2Contracts) registerMessengerHandlers() {
	l2.ScrollMessenger.RegisterSentMessage(func(vLog *types.Log, data *L2.L2ScrollMessengerSentMessageEvent) error {
		msgHash := utils.ComputeMessageHash(data.Sender, data.Target, data.Value, data.MessageNonce, data.Message)
		number := vLog.BlockNumber
		l2.msgSentEvents[number] = append(l2.msgSentEvents[number], &orm.L2MessengerEvent{
			Number:   vLog.BlockNumber,
			TxHash:   vLog.TxHash.String(),
			MsgHash:  msgHash.String(),
			Type:     orm.L2SentMessage,
			Target:   data.Target,
			Message:  data.Message,
			Log:      vLog,
			Value:    data.Value,
			MsgNonce: data.MessageNonce.Uint64(),
		})
		return nil
	})
	l2.ScrollMessenger.RegisterRelayedMessage(func(vLog *types.Log, data *L2.L2ScrollMessengerRelayedMessageEvent) error {
		l2.msgSentEvents[vLog.BlockNumber] = append(l2.msgSentEvents[vLog.BlockNumber], &orm.L2MessengerEvent{
			Number:  vLog.BlockNumber,
			TxHash:  vLog.TxHash.String(),
			MsgHash: common.BytesToHash(data.MessageHash[:]).String(),
			Type:    orm.L2RelayedMessage,
		})
		return nil
	})
	l2.ScrollMessenger.RegisterFailedRelayedMessage(func(vLog *types.Log, data *L2.L2ScrollMessengerFailedRelayedMessageEvent) error {
		l2.msgSentEvents[vLog.BlockNumber] = append(l2.msgSentEvents[vLog.BlockNumber], &orm.L2MessengerEvent{
			Number:  vLog.BlockNumber,
			TxHash:  vLog.TxHash.String(),
			MsgHash: common.BytesToHash(data.MessageHash[:]).String(),
			Type:    orm.L2FailedRelayedMessage,
		})
		return nil
	})
}

func (l2 *l2Contracts) storeMessengerEvents(ctx context.Context, start, end uint64) error {
	// Calculate withdraw root.
	var msgSentEvents []*orm.L2MessengerEvent
	for number := start; number <= end; number++ {
		var (
			msgHashes       []common.Hash
			latestSentEvent *orm.L2MessengerEvent
		)
		for _, msg := range l2.msgSentEvents[number] {
			msgSentEvents = append(msgSentEvents, msg)
			if msg.Type == orm.L2SentMessage {
				msgHashes = append(msgHashes, common.HexToHash(msg.MsgHash))
				latestSentEvent = msg
			}
		}
		if len(msgHashes) == 0 {
			l2.l2Confirms[number].WithdrawRootStatus = true
			continue
		}
		proofs := l2.withdraw.AppendMessages(msgHashes)
		latestSentEvent.MsgProof = common.Bytes2Hex(proofs[len(proofs)-1])
		l2.l2Confirms[number].WithdrawRoot = l2.withdraw.MessageRoot()
	}

	// Store messenger events.
	if len(msgSentEvents) > 0 {
		tx := l2.tx.WithContext(ctx)
		if err := tx.Model(&orm.L2MessengerEvent{}).Save(msgSentEvents).Error; err != nil {
			return err
		}
	}

	return nil
}

func (l2 *l2Contracts) storeL1ChainConfirms(ctx context.Context) error {
	var (
		numbers       []uint64
		withdrawRoots = map[uint64]common.Hash{}
		l2Confirms    = make([]*orm.L2ChainConfirm, 0, len(l2.l2Confirms))
		err           error
	)
	for i, monitor := range l2.l2Confirms {
		if !monitor.WithdrawRootStatus {
			numbers = append(numbers, monitor.Number)
		}
		l2Confirms = append(l2Confirms, l2.l2Confirms[i])
	}

	utils.TryTimes(3, func() bool {
		if len(numbers) == 0 {
			return true
		}
		// get withdraw root by batch.
		var roots []common.Hash
		roots, err = utils.GetBatchWithdrawRoots(ctx, l2.rpcCli, l2.MessageQueue.Address, numbers)
		if err != nil {
			return false
		}
		for i, number := range numbers {
			withdrawRoots[number] = roots[i]
		}
		return true
	})
	if err != nil {
		return err
	}

	for _, monitor := range l2Confirms {
		if len(withdrawRoots) == 0 {
			break
		}
		if monitor.WithdrawRootStatus {
			continue
		}
		expectRoot := withdrawRoots[monitor.Number]
		monitor.WithdrawRootStatus = monitor.WithdrawRoot == expectRoot
		// If the withdraw root doesn't match, alert it.
		if !monitor.WithdrawRootStatus {
			controller.WithdrawRootMisMatchTotal.Inc()
			msg := fmt.Sprintf(
				"withdraw root doestn't match, l2_number: %d, expect_withdraw_root: %s, actual_withdraw_root: %s",
				monitor.Number,
				expectRoot.String(),
				monitor.WithdrawRoot.String(),
			)
			log.Error("withdraw root doesn't match", "Number", monitor.Number, "expect_root", expectRoot.String(), "actual_root", monitor.WithdrawRoot.String())
			go controller.SlackNotify(msg)
		}
	}
	if err = l2.tx.Model(&orm.L2ChainConfirm{}).Save(l2Confirms).Error; err != nil {
		return err
	}
	return nil
}

func (l2 *l2Contracts) parseGatewayWithdraw(l2msg *orm.L2MessengerEvent) error {
	if len(l2msg.Message) < 4 {
		log.Warn("l2chain sendMessage content less than 4 bytes", "tx_hash", l2msg.Log.TxHash.String())
		return errMessenger
	}
	id := common.Bytes2Hex(l2msg.Message[:4])
	switch id {
	case "8eaac8a3": // FinalizeWithdrawETH
		return l2.parseGatewayWithdrawETH(l2msg)
	case "84bd13b0": // FinalizeWithdrawERC20
		return l2.parseGatewayWithdrawERC20(l2msg)
	case "d606b4dc": // FinalizeWithdrawERC721
		return l2.parseGatewayWithdrawERC721(l2msg)
	case "730608b3": // FinalizeWithdrawERC1155
		return l2.parseGatewayWithdrawERC155(l2msg)
	}
	log.Warn("l2chain sendMessage unexpect method_id", "tx_hash", l2msg.Log.TxHash.String(), "method_id", id)
	return errMessenger
}

func (l2 *l2Contracts) parseGatewayWithdrawETH(l2msg *orm.L2MessengerEvent) error {
	method, err := l1gateway.L1ETHGatewayABI.MethodById(l2msg.Message)
	if err != nil {
		return err
	}
	params, err := method.Inputs.Unpack(l2msg.Message[4:])
	if err != nil {
		return err
	}
	event := new(l1gateway.L1ETHGatewayFinalizeWithdrawETHEvent)
	err = method.Inputs.Copy(event, params)
	if err != nil {
		return err
	}
	l2.ethEvents = append(l2.ethEvents, &orm.L2ETHEvent{
		TxHead: &orm.TxHead{
			Number:  l2msg.Number,
			TxHash:  l2msg.Log.TxHash.String(),
			MsgHash: l2msg.MsgHash,
			Type:    orm.L2WithdrawETH,
		},
		Amount: event.Amount,
	})
	return nil
}

func (l2 *l2Contracts) parseGatewayWithdrawERC20(l2msg *orm.L2MessengerEvent) error {
	method, err := l1gateway.L1ERC20GatewayABI.MethodById(l2msg.Message)
	if err != nil {
		return err
	}
	params, err := method.Inputs.Unpack(l2msg.Message[4:])
	if err != nil {
		return err
	}
	event := new(l1gateway.L1ERC20GatewayFinalizeWithdrawERC20Event)
	err = method.Inputs.Copy(event, params)
	if err != nil {
		return err
	}

	var (
		_tp        orm.EventType
		gatewayCfg = l2.l1gatewayCfg
	)
	switch l2msg.Target {
	case gatewayCfg.DAIGateway:
		_tp = orm.L2WithdrawDAI
	case gatewayCfg.WETHGateway:
		_tp = orm.L2WithdrawWETH
	case gatewayCfg.StandardERC20Gateway:
		_tp = orm.L2WithdrawStandardERC20
	case gatewayCfg.CustomERC20Gateway:
		_tp = orm.L2WithdrawCustomERC20
	case gatewayCfg.USDCGateway:
		_tp = orm.L2USDCWithdrawERC20
	case gatewayCfg.LIDOGateway:
		_tp = orm.L2LIDOWithdrawERC20
	default:
		buf, _ := json.Marshal(event)
		log.Warn("l1chain unexpect erc20 withdraw transfer", "target address", l2msg.Target.String(), "content", string(buf))
	}

	l2.erc20Events = append(l2.erc20Events, &orm.L2ERC20Event{
		TxHead: &orm.TxHead{
			Number:  l2msg.Number,
			TxHash:  l2msg.Log.TxHash.String(),
			MsgHash: l2msg.MsgHash,
			Type:    _tp,
		},
		L1Token: event.L1Token.String(),
		L2Token: event.L2Token.String(),
		Amount:  event.Amount,
	})

	return nil
}

func (l2 *l2Contracts) parseGatewayWithdrawERC721(l2msg *orm.L2MessengerEvent) error {
	method, err := l1gateway.L1ERC721GatewayABI.MethodById(l2msg.Message)
	if err != nil {
		return err
	}
	params, err := method.Inputs.Unpack(l2msg.Message[4:])
	if err != nil {
		return err
	}
	event := new(l1gateway.L1ERC721GatewayFinalizeWithdrawERC721Event)
	err = method.Inputs.Copy(event, params)
	if err != nil {
		return err
	}
	l2.erc721Events = append(l2.erc721Events, &orm.L2ERC721Event{
		TxHead: &orm.TxHead{
			Number:  l2msg.Number,
			TxHash:  l2msg.Log.TxHash.String(),
			MsgHash: l2msg.MsgHash,
			Type:    orm.L2WithdrawERC721,
		},
		L1Token: event.L1Token.String(),
		L2Token: event.L2Token.String(),
		TokenID: event.TokenId,
	})

	return nil
}

func (l2 *l2Contracts) parseGatewayWithdrawERC155(l2msg *orm.L2MessengerEvent) error {
	method, err := l1gateway.L1ERC1155GatewayABI.MethodById(l2msg.Message)
	if err != nil {
		return err
	}
	params, err := method.Inputs.Unpack(l2msg.Message[4:])
	if err != nil {
		return err
	}
	event := new(l1gateway.L1ERC1155GatewayFinalizeWithdrawERC1155Event)
	err = method.Inputs.Copy(event, params)
	if err != nil {
		return err
	}
	l2.erc1155Events = append(l2.erc1155Events, &orm.L2ERC1155Event{
		TxHead: &orm.TxHead{
			Number:  l2msg.Number,
			TxHash:  l2msg.Log.TxHash.String(),
			MsgHash: l2msg.MsgHash,
			Type:    orm.L2WithdrawERC1155,
		},
		L1Token: event.L1Token.String(),
		L2Token: event.L2Token.String(),
		TokenID: event.TokenId,
		Amount:  event.Amount,
	})

	return nil
}
