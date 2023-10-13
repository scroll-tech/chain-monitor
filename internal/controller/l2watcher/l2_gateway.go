package l2watcher

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/log"

	l1gateway "chain-monitor/bytecode/scroll/L1/gateway"
	"chain-monitor/bytecode/scroll/L2/gateway"
	"chain-monitor/internal/controller"
	"chain-monitor/internal/orm"
)

var (
	errMessenger = errors.New("l2chain sendMessage content is not relate to gateway contract")
)

func (l2 *l2Contracts) registerGatewayHandlers() {
	l2.ETHGateway.RegisterWithdrawETH(func(vLog *types.Log, data *gateway.L2ETHGatewayWithdrawETHEvent) error {
		controller.ETHEventTotal.WithLabelValues(l2.chainName, orm.L2WithdrawETH.String()).Inc()
		l2.ethEvents = append(l2.ethEvents, newL2ETHEvent(orm.L2WithdrawETH, vLog, data.Amount))
		return nil
	})
	l2.ETHGateway.RegisterFinalizeDepositETH(func(vLog *types.Log, data *gateway.L2ETHGatewayFinalizeDepositETHEvent) error {
		controller.ETHEventTotal.WithLabelValues(l2.chainName, orm.L2FinalizeDepositETH.String()).Inc()
		l2.ethEvents = append(l2.ethEvents, newL2ETHEvent(orm.L2FinalizeDepositETH, vLog, data.Amount))
		return nil
	})

	l2.WETHGateway.RegisterWithdrawERC20(func(vLog *types.Log, data *gateway.L2WETHGatewayWithdrawERC20Event) error {
		controller.WETHEventTotal.WithLabelValues(l2.chainName, orm.L2WithdrawWETH.String()).Inc()
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2WithdrawWETH, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l2.WETHGateway.RegisterFinalizeDepositERC20(func(vLog *types.Log, data *gateway.L2WETHGatewayFinalizeDepositERC20Event) error {
		controller.WETHEventTotal.WithLabelValues(l2.chainName, orm.L2FinalizeDepositWETH.String()).Inc()
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2FinalizeDepositWETH, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l2.DAIGateway.RegisterWithdrawERC20(func(vLog *types.Log, data *gateway.L2DAIGatewayWithdrawERC20Event) error {
		controller.DAIEventTotal.WithLabelValues(l2.chainName, orm.L2WithdrawDAI.String()).Inc()
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2WithdrawDAI, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l2.DAIGateway.RegisterFinalizeDepositERC20(func(vLog *types.Log, data *gateway.L2DAIGatewayFinalizeDepositERC20Event) error {
		controller.DAIEventTotal.WithLabelValues(l2.chainName, orm.L2FinalizeDepositDAI.String()).Inc()
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2FinalizeDepositDAI, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l2.StandardERC20Gateway.RegisterWithdrawERC20(func(vLog *types.Log, data *gateway.L2StandardERC20GatewayWithdrawERC20Event) error {
		controller.StandardERC20EventsTotal.WithLabelValues(l2.chainName, orm.L2WithdrawStandardERC20.String()).Inc()
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2WithdrawStandardERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l2.StandardERC20Gateway.RegisterFinalizeDepositERC20(func(vLog *types.Log, data *gateway.L2StandardERC20GatewayFinalizeDepositERC20Event) error {
		controller.StandardERC20EventsTotal.WithLabelValues(l2.chainName, orm.L2FinalizeDepositStandardERC20.String()).Inc()
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2FinalizeDepositStandardERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l2.CustomERC20Gateway.RegisterWithdrawERC20(func(vLog *types.Log, data *gateway.L2CustomERC20GatewayWithdrawERC20Event) error {
		controller.CustomERC20EventsTotal.WithLabelValues(l2.chainName, orm.L2WithdrawCustomERC20.String()).Inc()
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2WithdrawCustomERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l2.CustomERC20Gateway.RegisterFinalizeDepositERC20(func(vLog *types.Log, data *gateway.L2CustomERC20GatewayFinalizeDepositERC20Event) error {
		controller.CustomERC20EventsTotal.WithLabelValues(l2.chainName, orm.L2FinalizeDepositCustomERC20.String()).Inc()
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2FinalizeDepositCustomERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l2.USDCERC20Gateway.RegisterWithdrawERC20(func(vLog *types.Log, data *gateway.L2CustomERC20GatewayWithdrawERC20Event) error {
		controller.USDCERC20EventsTotal.WithLabelValues(l2.chainName, orm.L2USDCWithdrawERC20.String()).Inc()
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2USDCWithdrawERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l2.USDCERC20Gateway.RegisterFinalizeDepositERC20(func(vLog *types.Log, data *gateway.L2CustomERC20GatewayFinalizeDepositERC20Event) error {
		controller.USDCERC20EventsTotal.WithLabelValues(l2.chainName, orm.L2USDCFinalizeDepositERC20.String()).Inc()
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2USDCFinalizeDepositERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l2.LIDOERC20Gateway.RegisterWithdrawERC20(func(vLog *types.Log, data *gateway.L2CustomERC20GatewayWithdrawERC20Event) error {
		controller.LIDOERC20EventsTotal.WithLabelValues(l2.chainName, orm.L2LIDOWithdrawERC20.String()).Inc()
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2LIDOWithdrawERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l2.LIDOERC20Gateway.RegisterFinalizeDepositERC20(func(vLog *types.Log, data *gateway.L2CustomERC20GatewayFinalizeDepositERC20Event) error {
		controller.LIDOERC20EventsTotal.WithLabelValues(l2.chainName, orm.L2LIDOFinalizeDepositERC20.String()).Inc()
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2LIDOFinalizeDepositERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})

	l2.ERC721Gateway.RegisterWithdrawERC721(func(vLog *types.Log, data *gateway.L2ERC721GatewayWithdrawERC721Event) error {
		controller.ERC721EventsTotal.WithLabelValues(l2.chainName, orm.L2WithdrawERC721.String()).Inc()
		l2.erc721Events = append(l2.erc721Events, newL2ERC721Event(orm.L2WithdrawERC721, vLog, data.L1Token, data.L2Token, data.TokenId))
		return nil
	})
	l2.ERC721Gateway.RegisterFinalizeDepositERC721(func(vLog *types.Log, data *gateway.L2ERC721GatewayFinalizeDepositERC721Event) error {
		controller.ERC721EventsTotal.WithLabelValues(l2.chainName, orm.L2FinalizeDepositERC721.String()).Inc()
		l2.erc721Events = append(l2.erc721Events, newL2ERC721Event(orm.L2FinalizeDepositERC721, vLog, data.L1Token, data.L2Token, data.TokenId))
		return nil
	})

	l2.ERC1155Gateway.RegisterWithdrawERC1155(func(vLog *types.Log, data *gateway.L2ERC1155GatewayWithdrawERC1155Event) error {
		controller.ERC1155EventsTotal.WithLabelValues(l2.chainName, orm.L2WithdrawERC1155.String()).Inc()
		l2.erc1155Events = append(l2.erc1155Events, newL2ERC1155Event(orm.L2WithdrawERC1155, vLog, data.L1Token, data.L2Token, data.TokenId, data.Amount))
		return nil
	})
	l2.ERC1155Gateway.RegisterFinalizeDepositERC1155(func(vLog *types.Log, data *gateway.L2ERC1155GatewayFinalizeDepositERC1155Event) error {
		controller.ERC1155EventsTotal.WithLabelValues(l2.chainName, orm.L2FinalizeDepositERC1155.String()).Inc()
		l2.erc1155Events = append(l2.erc1155Events, newL2ERC1155Event(orm.L2FinalizeDepositERC1155, vLog, data.L1Token, data.L2Token, data.TokenId, data.Amount))
		return nil
	})
}

func (l2 *l2Contracts) integrateGatewayEvents() error {
	for _, event := range l2.ethEvents {
		if msgHash, exist := l2.txHashMsgHash[event.TxHash]; exist {
			event.MsgHash = msgHash.String()
			delete(l2.txHashMsgHash, event.TxHash)
		}
	}

	for _, event := range l2.erc20Events {
		if msgHash, exist := l2.txHashMsgHash[event.TxHash]; exist {
			event.MsgHash = msgHash.String()
			delete(l2.txHashMsgHash, event.TxHash)
		}
	}

	for _, event := range l2.erc721Events {
		if msgHash, exist := l2.txHashMsgHash[event.TxHash]; exist {
			event.MsgHash = msgHash.String()
			delete(l2.txHashMsgHash, event.TxHash)
		}
	}

	for _, event := range l2.erc1155Events {
		if msgHash, exist := l2.txHashMsgHash[event.TxHash]; exist {
			event.MsgHash = msgHash.String()
			delete(l2.txHashMsgHash, event.TxHash)
		}
	}

	for _, msgs := range l2.msgSentEvents {
		for _, msg := range msgs {
			if _, exist := l2.txHashMsgHash[msg.Log.TxHash.String()]; !exist {
				continue
			}
			err := l2.parseGatewayWithdraw(msg)
			if errors.Is(err, errMessenger) {
				msg.FromGateway = false
				continue
			}
			if err != nil {
				log.Error("l2chain failed to parse gateway message", "tx_hash", msg.Log.TxHash.String(), "err", err)
				return err
			}
		}
	}

	return nil
}

func (l2 *l2Contracts) parseGatewayWithdraw(l2msg *orm.L2MessengerEvent) error {
	if len(l2msg.Message) < 4 {
		log.Warn("l2chain sendMessage content less than 4 bytes", "tx_hash", l2msg.Log.TxHash.String())
		return errMessenger
	}
	_id := common.Bytes2Hex(l2msg.Message[:4])
	switch _id {
	case "8eaac8a3": // FinalizeWithdrawETH
		return l2.parseGatewayWithdrawETH(l2msg)
	case "84bd13b0": // FinalizeWithdrawERC20
		return l2.parseGatewayWithdrawERC20(l2msg)
	case "d606b4dc": // FinalizeWithdrawERC721
		return l2.parseGatewayWithdrawERC721(l2msg)
	case "730608b3": // FinalizeWithdrawERC1155
		return l2.parseGatewayWithdrawERC155(l2msg)
	}
	log.Warn("l2chain sendMessage unexpect method_id", "tx_hash", l2msg.Log.TxHash.String(), "method_id", _id)
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
		log.Warn("l1chain unexpect erc20 deposit transfer", "content", string(buf))
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

func (l2 *l2Contracts) storeGatewayEvents() error {
	// store l2 eth events.
	if len(l2.ethEvents) > 0 {
		if err := l2.tx.Save(l2.ethEvents).Error; err != nil {
			return err
		}
	}

	// store l2 erc20 events.
	if len(l2.erc20Events) > 0 {
		if err := l2.tx.Save(l2.erc20Events).Error; err != nil {
			return err
		}
	}

	// store l2 err721 events.
	if len(l2.erc721Events) > 0 {
		if err := l2.tx.Save(l2.erc721Events).Error; err != nil {
			return err
		}
	}

	// store l2 erc1155 events.
	if len(l2.erc1155Events) > 0 {
		if err := l2.tx.Save(l2.erc1155Events).Error; err != nil {
			return err
		}
	}

	return nil
}

func newL2ETHEvent(eventType orm.EventType, vLog *types.Log, amount *big.Int) *orm.L2ETHEvent {
	return &orm.L2ETHEvent{
		TxHead: orm.NewTxHead(vLog, eventType),
		Amount: amount,
	}
}

func newL2ETH20Event(eventType orm.EventType, vLog *types.Log, l1Token, l2Token common.Address, amount *big.Int) *orm.L2ERC20Event {
	return &orm.L2ERC20Event{
		TxHead:  orm.NewTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		Amount:  amount,
	}
}

func newL2ERC721Event(eventType orm.EventType, vLog *types.Log, l1Token, l2Token common.Address, tokenID *big.Int) *orm.L2ERC721Event {
	return &orm.L2ERC721Event{
		TxHead:  orm.NewTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		TokenID: tokenID,
	}
}

func newL2ERC1155Event(eventType orm.EventType, vLog *types.Log, l1Token, l2Token common.Address, tokenID, amount *big.Int) *orm.L2ERC1155Event {
	return &orm.L2ERC1155Event{
		TxHead:  orm.NewTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		TokenID: tokenID,
		Amount:  amount,
	}
}
