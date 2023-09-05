package l2watcher

import (
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"

	"chain-monitor/bytecode/scroll/L2/gateway"
	"chain-monitor/internal/orm"
)

func (l2 *l2Contracts) registerGatewayHandlers() {
	l2.ETHGateway.RegisterWithdrawETH(func(vLog *types.Log, data *gateway.L2ETHGatewayWithdrawETHEvent) error {
		l2.ethEvents = append(l2.ethEvents, newL2ETHEvent(orm.L2WithdrawETH, vLog, data.Amount))
		return nil
	})
	l2.ETHGateway.RegisterFinalizeDepositETH(func(vLog *types.Log, data *gateway.L2ETHGatewayFinalizeDepositETHEvent) error {
		l2.ethEvents = append(l2.ethEvents, newL2ETHEvent(orm.L2FinalizeDepositETH, vLog, data.Amount))
		return nil
	})

	l2.WETHGateway.RegisterWithdrawERC20(func(vLog *types.Log, data *gateway.L2WETHGatewayWithdrawERC20Event) error {
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2WithdrawWETH, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l2.WETHGateway.RegisterFinalizeDepositERC20(func(vLog *types.Log, data *gateway.L2WETHGatewayFinalizeDepositERC20Event) error {
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2FinalizeDepositWETH, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l2.DAIGateway.RegisterWithdrawERC20(func(vLog *types.Log, data *gateway.L2DAIGatewayWithdrawERC20Event) error {
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2WithdrawDAI, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l2.DAIGateway.RegisterFinalizeDepositERC20(func(vLog *types.Log, data *gateway.L2DAIGatewayFinalizeDepositERC20Event) error {
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2FinalizeDepositDAI, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l2.StandardERC20Gateway.RegisterWithdrawERC20(func(vLog *types.Log, data *gateway.L2StandardERC20GatewayWithdrawERC20Event) error {
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2WithdrawStandardERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l2.StandardERC20Gateway.RegisterFinalizeDepositERC20(func(vLog *types.Log, data *gateway.L2StandardERC20GatewayFinalizeDepositERC20Event) error {
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2FinalizeDepositStandardERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l2.CustomERC20Gateway.RegisterWithdrawERC20(func(vLog *types.Log, data *gateway.L2CustomERC20GatewayWithdrawERC20Event) error {
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2WithdrawCustomERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})
	l2.CustomERC20Gateway.RegisterFinalizeDepositERC20(func(vLog *types.Log, data *gateway.L2CustomERC20GatewayFinalizeDepositERC20Event) error {
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2FinalizeDepositCustomERC20, vLog, data.L1Token, data.L2Token, data.Amount))
		return nil
	})

	l2.ERC721Gateway.RegisterWithdrawERC721(func(vLog *types.Log, data *gateway.L2ERC721GatewayWithdrawERC721Event) error {
		l2.erc721Events = append(l2.erc721Events, newL2ERC721Event(orm.L2WithdrawERC721, vLog, data.L1Token, data.L2Token, data.TokenID))
		return nil
	})
	l2.ERC721Gateway.RegisterFinalizeDepositERC721(func(vLog *types.Log, data *gateway.L2ERC721GatewayFinalizeDepositERC721Event) error {
		l2.erc721Events = append(l2.erc721Events, newL2ERC721Event(orm.L2FinalizeDepositERC721, vLog, data.L1Token, data.L2Token, data.TokenID))
		return nil
	})

	l2.ERC1155Gateway.RegisterWithdrawERC1155(func(vLog *types.Log, data *gateway.L2ERC1155GatewayWithdrawERC1155Event) error {
		l2.erc1155Events = append(l2.erc1155Events, newL2ERC1155Event(orm.L2WithdrawERC1155, vLog, data.L1Token, data.L2Token, data.TokenID, data.Amount))
		return nil
	})
	l2.ERC1155Gateway.RegisterFinalizeDepositERC1155(func(vLog *types.Log, data *gateway.L2ERC1155GatewayFinalizeDepositERC1155Event) error {
		l2.erc1155Events = append(l2.erc1155Events, newL2ERC1155Event(orm.L2FinalizeDepositERC1155, vLog, data.L1Token, data.L2Token, data.TokenID, data.Amount))
		return nil
	})
}

func (l2 *l2Contracts) storeGatewayEvents() error {
	// store l2 eth events.
	for i := 0; i < len(l2.ethEvents); i++ {
		event := l2.ethEvents[i]
		if msgHash, exist := l2.txHashMsgHash[event.TxHash]; exist {
			event.MsgHash = msgHash.String()
		}
	}
	if len(l2.ethEvents) > 0 {
		if err := l2.tx.Save(l2.ethEvents).Error; err != nil {
			return err
		}
	}

	// store l2 erc20 events.
	for i := 0; i < len(l2.erc20Events); i++ {
		event := l2.erc20Events[i]
		if msgHash, exist := l2.txHashMsgHash[event.TxHash]; exist {
			event.MsgHash = msgHash.String()
		}
	}
	if len(l2.erc20Events) > 0 {
		if err := l2.tx.Save(l2.erc20Events).Error; err != nil {
			return err
		}
	}

	// store l2 err721 events.
	for i := 0; i < len(l2.erc721Events); i++ {
		event := l2.erc721Events[i]
		if msgHash, exist := l2.txHashMsgHash[event.TxHash]; exist {
			event.MsgHash = msgHash.String()
		}
	}
	if len(l2.erc721Events) > 0 {
		if err := l2.tx.Save(l2.erc721Events).Error; err != nil {
			return err
		}
	}

	// store l2 erc1155 events.
	for i := 0; i < len(l2.erc1155Events); i++ {
		event := l2.erc1155Events[i]
		if msgHash, exist := l2.txHashMsgHash[event.TxHash]; exist {
			event.MsgHash = msgHash.String()
		}
	}
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
