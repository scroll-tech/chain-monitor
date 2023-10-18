package l2watcher

import (
	"errors"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/log"

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

func (l2 *l2Contracts) gatewayEvents() error {
	var msgSentEvents = map[string]*orm.L2MessengerEvent{}
	for _, msgs := range l2.msgSentEvents {
		for _, msg := range msgs {
			msgSentEvents[msg.TxHash] = msg
		}
	}

	for _, event := range l2.ethEvents {
		if msg, exist := msgSentEvents[event.TxHash]; exist {
			event.MsgHash, msg.FromGateway = msg.MsgHash, true
		}
	}

	for _, event := range l2.erc20Events {
		if msg, exist := msgSentEvents[event.TxHash]; exist {
			event.MsgHash, msg.FromGateway = msg.MsgHash, true
		}
	}

	for _, event := range l2.erc721Events {
		if msg, exist := msgSentEvents[event.TxHash]; exist {
			event.MsgHash, msg.FromGateway = msg.MsgHash, true
		}
	}

	for _, event := range l2.erc1155Events {
		if msg, exist := msgSentEvents[event.TxHash]; exist {
			event.MsgHash, msg.FromGateway = msg.MsgHash, true
		}
	}

	for _, msg := range msgSentEvents {
		if msg.IsNotGatewaySentMessage() {
			err := l2.parseGatewayWithdraw(msg)
			if errors.Is(err, errMessenger) {
				continue
			}
			if err != nil {
				log.Error("l2chain failed to parse gateway message", "tx_hash", msg.Log.TxHash.String(), "err", err)
				return err
			}
			msg.FromGateway = true
		}
	}

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
