package l1watcher

import (
	"errors"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/log"

	"chain-monitor/bytecode/scroll/L1/gateway"
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

func (l1 *l1Contracts) gatewayEvents() error {
	for _, event := range l1.ethEvents {
		if msg, exist := l1.msgSentEvents[event.TxHash]; exist {
			event.MsgHash, msg.FromGateway = msg.MsgHash, true
		}
	}

	for _, event := range l1.erc20Events {
		if msg, exist := l1.msgSentEvents[event.TxHash]; exist {
			event.MsgHash, msg.FromGateway = msg.MsgHash, true
		}
	}

	for _, event := range l1.erc721Events {
		if msg, exist := l1.msgSentEvents[event.TxHash]; exist {
			event.MsgHash, msg.FromGateway = msg.MsgHash, true
		}
	}

	for _, event := range l1.erc1155Events {
		if msg, exist := l1.msgSentEvents[event.TxHash]; exist {
			event.MsgHash, msg.FromGateway = msg.MsgHash, true
		}
	}

	for _, msg := range l1.msgSentEvents {
		if msg.IsNotGatewaySentMessage() {
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
