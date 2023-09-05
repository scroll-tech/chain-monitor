package l1watcher

import (
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"

	"chain-monitor/bytecode/scroll/L1/gateway"
	"chain-monitor/internal/orm"
)

func (l1 *l1Contracts) registerGatewayHandlers() {
	l1.ETHGateway.RegisterDepositETH(func(vLog *types.Log, data *gateway.L1ETHGatewayDepositETHEvent) error {
		l1.ethEvents = append(l1.ethEvents, newL1ETHEvent(orm.L1DepositETH, vLog, data.From, data.To, data.Amount))
		return nil
	})
	l1.ETHGateway.RegisterFinalizeWithdrawETH(func(vLog *types.Log, data *gateway.L1ETHGatewayFinalizeWithdrawETHEvent) error {
		l1.ethEvents = append(l1.ethEvents, newL1ETHEvent(orm.L1FinalizeWithdrawETH, vLog, data.From, data.To, data.Amount))
		return nil
	})

	l1.WETHGateway.RegisterDepositERC20(func(vLog *types.Log, data *gateway.L1WETHGatewayDepositERC20Event) error {
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1DepositWETH, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount))
		return nil
	})
	l1.WETHGateway.RegisterFinalizeWithdrawERC20(func(vLog *types.Log, data *gateway.L1WETHGatewayFinalizeWithdrawERC20Event) error {
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1FinalizeWithdrawWETH, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount))
		return nil
	})
	l1.DAIGateway.RegisterDepositERC20(func(vLog *types.Log, data *gateway.L1DAIGatewayDepositERC20Event) error {
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1DepositDAI, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount))
		return nil
	})
	l1.DAIGateway.RegisterFinalizeWithdrawERC20(func(vLog *types.Log, data *gateway.L1DAIGatewayFinalizeWithdrawERC20Event) error {
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1FinalizeWithdrawDAI, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount))
		return nil
	})
	l1.StandardERC20Gateway.RegisterDepositERC20(func(vLog *types.Log, data *gateway.L1StandardERC20GatewayDepositERC20Event) error {
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1DepositStandardERC20, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount))
		return nil
	})
	l1.StandardERC20Gateway.RegisterFinalizeWithdrawERC20(func(vLog *types.Log, data *gateway.L1StandardERC20GatewayFinalizeWithdrawERC20Event) error {
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1FinalizeWithdrawStandardERC20, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount))
		return nil
	})
	l1.CustomERC20Gateway.RegisterDepositERC20(func(vLog *types.Log, data *gateway.L1CustomERC20GatewayDepositERC20Event) error {
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1DepositCustomERC20, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount))
		return nil
	})
	l1.CustomERC20Gateway.RegisterFinalizeWithdrawERC20(func(vLog *types.Log, data *gateway.L1CustomERC20GatewayFinalizeWithdrawERC20Event) error {
		l1.erc20Events = append(l1.erc20Events, newL1ETH20Event(orm.L1FinalizeWithdrawCustomERC20, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount))
		return nil
	})

	l1.ERC721Gateway.RegisterDepositERC721(func(vLog *types.Log, data *gateway.L1ERC721GatewayDepositERC721Event) error {
		l1.erc721Events = append(l1.erc721Events, newL1ERC721Event(orm.L1DepositERC721, vLog, data.L1Token, data.L2Token, data.From, data.To, data.TokenId))
		return nil
	})
	l1.ERC721Gateway.RegisterFinalizeWithdrawERC721(func(vLog *types.Log, data *gateway.L1ERC721GatewayFinalizeWithdrawERC721Event) error {
		l1.erc721Events = append(l1.erc721Events, newL1ERC721Event(orm.L1FinalizeWithdrawERC721, vLog, data.L1Token, data.L2Token, data.From, data.To, data.TokenId))
		return nil
	})
	l1.ERC1155Gateway.RegisterDepositERC1155(func(vLog *types.Log, data *gateway.L1ERC1155GatewayDepositERC1155Event) error {
		l1.erc1155Events = append(l1.erc1155Events, newL1ERC1155Event(orm.L1DepositERC1155, vLog, data.L1Token, data.L2Token, data.From, data.To, data.TokenId, data.Amount))
		return nil
	})
	l1.ERC1155Gateway.RegisterFinalizeWithdrawERC1155(func(vLog *types.Log, data *gateway.L1ERC1155GatewayFinalizeWithdrawERC1155Event) error {
		l1.erc1155Events = append(l1.erc1155Events, newL1ERC1155Event(orm.L1FinalizeWithdrawERC1155, vLog, data.L1Token, data.L2Token, data.From, data.To, data.TokenId, data.Amount))
		return nil
	})
}

func (l1 *l1Contracts) storeGatewayEvents() error {
	// store l1 eth events.
	for i := 0; i < len(l1.ethEvents); i++ {
		event := l1.ethEvents[i]
		if msgHash, exist := l1.txHashMsgHash[event.TxHash]; exist {
			event.MsgHash = msgHash.String()
		}
	}
	if len(l1.ethEvents) > 0 {
		if err := l1.tx.Save(l1.ethEvents).Error; err != nil {
			return err
		}
	}

	// store l1 erc20 events.
	for i := 0; i < len(l1.erc20Events); i++ {
		event := l1.erc20Events[i]
		if msgHash, exist := l1.txHashMsgHash[event.TxHash]; exist {
			event.MsgHash = msgHash.String()
		}
	}
	if len(l1.erc20Events) > 0 {
		if err := l1.tx.Save(l1.erc20Events).Error; err != nil {
			return err
		}
	}

	// store l1 err721 events.
	for i := 0; i < len(l1.erc721Events); i++ {
		event := l1.erc721Events[i]
		if msgHash, exist := l1.txHashMsgHash[event.TxHash]; exist {
			event.MsgHash = msgHash.String()
		}
	}
	if len(l1.erc721Events) > 0 {
		if err := l1.tx.Save(l1.erc721Events).Error; err != nil {
			return err
		}
	}

	// store l1 erc1155 events.
	for i := 0; i < len(l1.erc1155Events); i++ {
		event := l1.erc1155Events[i]
		if msgHash, exist := l1.txHashMsgHash[event.TxHash]; exist {
			event.MsgHash = msgHash.String()
		}
	}
	if len(l1.erc1155Events) > 0 {
		if err := l1.tx.Save(l1.erc1155Events).Error; err != nil {
			return err
		}
	}
	return nil
}

func newL1ETHEvent(eventType orm.EventType, vLog *types.Log, from, to common.Address, amount *big.Int) *orm.L1ETHEvent {
	return &orm.L1ETHEvent{
		TxHead: orm.NewTxHead(vLog, eventType),
		Amount: amount,
	}
}

func newL1ETH20Event(eventType orm.EventType, vLog *types.Log, l1Token, l2Token, from, to common.Address, amount *big.Int) *orm.L1ERC20Event {
	return &orm.L1ERC20Event{
		TxHead:  orm.NewTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		Amount:  amount,
	}
}

func newL1ERC721Event(eventType orm.EventType, vLog *types.Log, l1Token, l2Token, from, to common.Address, tokenId *big.Int) *orm.L1ERC721Event {
	return &orm.L1ERC721Event{
		TxHead:  orm.NewTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		TokenId: tokenId,
	}
}

func newL1ERC1155Event(eventType orm.EventType, vLog *types.Log, l1Token, l2Token, from, to common.Address, tokenId, amount *big.Int) *orm.L1ERC1155Event {
	return &orm.L1ERC1155Event{
		TxHead:  orm.NewTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		TokenId: tokenId,
		Amount:  amount,
	}
}
