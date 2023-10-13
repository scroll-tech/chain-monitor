package l1watcher

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/log"

	"chain-monitor/bytecode/scroll/L1/gateway"
	l2gateway "chain-monitor/bytecode/scroll/L2/gateway"
	"chain-monitor/internal/controller"
	"chain-monitor/internal/orm"
)

var (
	errMessenger = errors.New("l1chain sendMessage content is not relate to gateway contract")
)

func (l1 *l1Contracts) registerGatewayHandlers() {
	l1.ETHGateway.RegisterDepositETH(func(vLog *types.Log, data *gateway.L1ETHGatewayDepositETHEvent) error {
		controller.ETHEventTotal.WithLabelValues(l1.chainName, orm.L1DepositETH.String()).Inc()
		l1.ethEvents = append(l1.ethEvents, newL1ETHEvent(orm.L1DepositETH, vLog, data.Amount))
		return nil
	})
	l1.ETHGateway.RegisterFinalizeWithdrawETH(func(vLog *types.Log, data *gateway.L1ETHGatewayFinalizeWithdrawETHEvent) error {
		controller.ETHEventTotal.WithLabelValues(l1.chainName, orm.L1FinalizeWithdrawETH.String()).Inc()
		l1.ethEvents = append(l1.ethEvents, newL1ETHEvent(orm.L1FinalizeWithdrawETH, vLog, data.Amount))
		return nil
	})

	l1.WETHGateway.RegisterDepositERC20(func(vLog *types.Log, data *gateway.L1WETHGatewayDepositERC20Event) error {
		controller.WETHEventTotal.WithLabelValues(l1.chainName, orm.L1DepositWETH.String()).Inc()
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1DepositWETH, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l1.WETHGateway.RegisterFinalizeWithdrawERC20(func(vLog *types.Log, data *gateway.L1WETHGatewayFinalizeWithdrawERC20Event) error {
		controller.WETHEventTotal.WithLabelValues(l1.chainName, orm.L1FinalizeWithdrawWETH.String()).Inc()
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1FinalizeWithdrawWETH, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l1.DAIGateway.RegisterDepositERC20(func(vLog *types.Log, data *gateway.L1DAIGatewayDepositERC20Event) error {
		controller.DAIEventTotal.WithLabelValues(l1.chainName, orm.L1DepositDAI.String()).Inc()
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1DepositDAI, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l1.DAIGateway.RegisterFinalizeWithdrawERC20(func(vLog *types.Log, data *gateway.L1DAIGatewayFinalizeWithdrawERC20Event) error {
		controller.DAIEventTotal.WithLabelValues(l1.chainName, orm.L1FinalizeWithdrawDAI.String()).Inc()
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1FinalizeWithdrawDAI, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l1.StandardERC20Gateway.RegisterDepositERC20(func(vLog *types.Log, data *gateway.L1StandardERC20GatewayDepositERC20Event) error {
		controller.StandardERC20EventsTotal.WithLabelValues(l1.chainName, orm.L1DepositStandardERC20.String()).Inc()
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1DepositStandardERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l1.StandardERC20Gateway.RegisterFinalizeWithdrawERC20(func(vLog *types.Log, data *gateway.L1StandardERC20GatewayFinalizeWithdrawERC20Event) error {
		controller.StandardERC20EventsTotal.WithLabelValues(l1.chainName, orm.L1FinalizeWithdrawStandardERC20.String()).Inc()
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1FinalizeWithdrawStandardERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l1.CustomERC20Gateway.RegisterDepositERC20(func(vLog *types.Log, data *gateway.L1CustomERC20GatewayDepositERC20Event) error {
		controller.CustomERC20EventsTotal.WithLabelValues(l1.chainName, orm.L1DepositCustomERC20.String()).Inc()
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1DepositCustomERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l1.CustomERC20Gateway.RegisterFinalizeWithdrawERC20(func(vLog *types.Log, data *gateway.L1CustomERC20GatewayFinalizeWithdrawERC20Event) error {
		controller.CustomERC20EventsTotal.WithLabelValues(l1.chainName, orm.L1FinalizeWithdrawCustomERC20.String()).Inc()
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1FinalizeWithdrawCustomERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l1.USDCERC20Gateway.RegisterDepositERC20(func(vLog *types.Log, data *gateway.L1CustomERC20GatewayDepositERC20Event) error {
		controller.USDCERC20EventsTotal.WithLabelValues(l1.chainName, orm.L1USDCDepositERC20.String()).Inc()
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1USDCDepositERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l1.USDCERC20Gateway.RegisterFinalizeWithdrawERC20(func(vLog *types.Log, data *gateway.L1CustomERC20GatewayFinalizeWithdrawERC20Event) error {
		controller.USDCERC20EventsTotal.WithLabelValues(l1.chainName, orm.L1USDCFinalizeWithdrawERC20.String()).Inc()
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1USDCFinalizeWithdrawERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l1.LIDOERC20Gateway.RegisterDepositERC20(func(vLog *types.Log, data *gateway.L1CustomERC20GatewayDepositERC20Event) error {
		controller.LIDOERC20EventsTotal.WithLabelValues(l1.chainName, orm.L1LIDODepositERC20.String()).Inc()
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1LIDODepositERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l1.LIDOERC20Gateway.RegisterFinalizeWithdrawERC20(func(vLog *types.Log, data *gateway.L1CustomERC20GatewayFinalizeWithdrawERC20Event) error {
		controller.LIDOERC20EventsTotal.WithLabelValues(l1.chainName, orm.L1LIDOFinalizeWithdrawERC20.String()).Inc()
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1LIDOFinalizeWithdrawERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})

	l1.ERC721Gateway.RegisterDepositERC721(func(vLog *types.Log, data *gateway.L1ERC721GatewayDepositERC721Event) error {
		controller.ERC721EventsTotal.WithLabelValues(l1.chainName, orm.L1DepositERC721.String()).Inc()
		l1.erc721Events = append(l1.erc721Events, newL1ERC721Event(orm.L1DepositERC721, vLog, data.L1Token, data.L2Token, data.TokenId))
		return nil
	})
	l1.ERC721Gateway.RegisterFinalizeWithdrawERC721(func(vLog *types.Log, data *gateway.L1ERC721GatewayFinalizeWithdrawERC721Event) error {
		controller.ERC721EventsTotal.WithLabelValues(l1.chainName, orm.L1FinalizeWithdrawERC721.String()).Inc()
		l1.erc721Events = append(l1.erc721Events, newL1ERC721Event(orm.L1FinalizeWithdrawERC721, vLog, data.L1Token, data.L2Token, data.TokenId))
		return nil
	})
	l1.ERC1155Gateway.RegisterDepositERC1155(func(vLog *types.Log, data *gateway.L1ERC1155GatewayDepositERC1155Event) error {
		controller.ERC1155EventsTotal.WithLabelValues(l1.chainName, orm.L1DepositERC1155.String()).Inc()
		l1.erc1155Events = append(l1.erc1155Events, newL1ERC1155Event(orm.L1DepositERC1155, vLog, data.L1Token, data.L2Token, data.TokenId, data.Amount))
		return nil
	})
	l1.ERC1155Gateway.RegisterFinalizeWithdrawERC1155(func(vLog *types.Log, data *gateway.L1ERC1155GatewayFinalizeWithdrawERC1155Event) error {
		controller.ERC1155EventsTotal.WithLabelValues(l1.chainName, orm.L1FinalizeWithdrawERC1155.String()).Inc()
		l1.erc1155Events = append(l1.erc1155Events, newL1ERC1155Event(orm.L1FinalizeWithdrawERC1155, vLog, data.L1Token, data.L2Token, data.TokenId, data.Amount))
		return nil
	})
}

func (l1 *l1Contracts) integrateGatewayEvents() error {
	for i := 0; i < len(l1.ethEvents); i++ {
		event := l1.ethEvents[i]
		if msg, exist := l1.msgSentEvents[event.TxHash]; exist {
			event.MsgHash = msg.MsgHash
			msg.FromGateway = true
		}
	}

	for i := 0; i < len(l1.erc20Events); i++ {
		event := l1.erc20Events[i]
		if msg, exist := l1.msgSentEvents[event.TxHash]; exist {
			event.MsgHash = msg.MsgHash
			msg.FromGateway = true
		}
	}

	for i := 0; i < len(l1.erc721Events); i++ {
		event := l1.erc721Events[i]
		if msg, exist := l1.msgSentEvents[event.TxHash]; exist {
			event.MsgHash = msg.MsgHash
			msg.FromGateway = true
		}
	}

	for i := 0; i < len(l1.erc1155Events); i++ {
		event := l1.erc1155Events[i]
		if msg, exist := l1.msgSentEvents[event.TxHash]; exist {
			event.MsgHash = msg.MsgHash
			msg.FromGateway = true
		}
	}

	for _, msg := range l1.msgSentEvents {
		if !(msg.Type == orm.L1SentMessage) || msg.FromGateway {
			continue
		}
		err := l1.parseGatewayDeposit(msg)
		if errors.Is(err, errMessenger) {
			continue
		}
		if err != nil {
			log.Error("l1chain failed to parse gateway message", "tx_hash", msg.TxHash, "err", err)
			return err
		}
		msg.FromGateway = true
	}

	return nil
}

func (l1 *l1Contracts) parseGatewayDeposit(l1msg *orm.L1MessengerEvent) error {
	if len(l1msg.Message) < 4 {
		log.Warn("l1chain sendMessage content less than 4 bytes", "tx_hash", l1msg.TxHash)
		return errMessenger
	}
	_id := common.Bytes2Hex(l1msg.Message[:4])
	switch _id {
	case "232e8748": // FinalizeDepositETH
		return l1.parseGatewayDepositETH(l1msg.TxHead, l1msg.Message)
	case "8431f5c1": // FinalizeDepositERC20
		return l1.parseGatewayDepositERC20(l1msg.TxHead, l1msg.Target, l1msg.Message)
	case "982b151f": // FinalizeBatchDepositERC721
		return l1.parseGatewayDepositERC721(l1msg.TxHead, l1msg.Message)
	case "4764cc62": // FinalizeDepositERC1155
		return l1.parseGatewayDepositERC1155(l1msg.TxHead, l1msg.Message)
	}
	log.Warn("l1chain sendMessage unexpect method_id", "tx_hash", l1msg.TxHash, "method_id", _id)
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
		log.Warn("l1chain unexpect erc20 deposit transfer", "content", string(buf))
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
		for i := 0; i < length; i++ {
			for k, v := range l1.msgSentEvents {
				// Only store sentMessage events.
				if v.Type == orm.L1SentMessage {
					msgs = append(msgs, l1.msgSentEvents[k])
				}
			}
		}
		if err := l1.tx.Save(msgs).Error; err != nil {
			return err
		}
	}

	return nil
}

func newL1ETHEvent(eventType orm.EventType, vLog *types.Log, amount *big.Int) *orm.L1ETHEvent {
	return &orm.L1ETHEvent{
		TxHead: orm.NewTxHead(vLog, eventType),
		Amount: amount,
	}
}

func newL1ETH20Event(eventType orm.EventType, vLog *types.Log, l1Token, l2Token common.Address, amount *big.Int) *orm.L1ERC20Event {
	return &orm.L1ERC20Event{
		TxHead:  orm.NewTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		Amount:  amount,
	}
}

func newL1ERC721Event(eventType orm.EventType, vLog *types.Log, l1Token, l2Token common.Address, tokenID *big.Int) *orm.L1ERC721Event {
	return &orm.L1ERC721Event{
		TxHead:  orm.NewTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		TokenID: tokenID,
	}
}

func newL1ERC1155Event(eventType orm.EventType, vLog *types.Log, l1Token, l2Token common.Address, tokenID, amount *big.Int) *orm.L1ERC1155Event {
	return &orm.L1ERC1155Event{
		TxHead:  orm.NewTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		TokenID: tokenID,
		Amount:  amount,
	}
}
