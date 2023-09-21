package l2watcher

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"

	"chain-monitor/bytecode/scroll/L2/gateway"
	"chain-monitor/internal/controller"
	"chain-monitor/internal/orm"
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

	l2.ERC721Gateway.RegisterWithdrawERC721(func(vLog *types.Log, data *gateway.L2ERC721GatewayWithdrawERC721Event) error {
		controller.ERC721EventsTotal.WithLabelValues(l2.chainName, orm.L2WithdrawERC721.String()).Inc()
		l2.erc721Events = append(l2.erc721Events, newL2ERC721Event(orm.L2WithdrawERC721, vLog, data.L1Token, data.L2Token, data.TokenID))
		return nil
	})
	l2.ERC721Gateway.RegisterFinalizeDepositERC721(func(vLog *types.Log, data *gateway.L2ERC721GatewayFinalizeDepositERC721Event) error {
		controller.ERC721EventsTotal.WithLabelValues(l2.chainName, orm.L2FinalizeDepositERC721.String()).Inc()
		l2.erc721Events = append(l2.erc721Events, newL2ERC721Event(orm.L2FinalizeDepositERC721, vLog, data.L1Token, data.L2Token, data.TokenID))
		return nil
	})

	l2.ERC1155Gateway.RegisterWithdrawERC1155(func(vLog *types.Log, data *gateway.L2ERC1155GatewayWithdrawERC1155Event) error {
		controller.ERC1155EventsTotal.WithLabelValues(l2.chainName, orm.L2WithdrawERC1155.String()).Inc()
		l2.erc1155Events = append(l2.erc1155Events, newL2ERC1155Event(orm.L2WithdrawERC1155, vLog, data.L1Token, data.L2Token, data.TokenID, data.Amount))
		return nil
	})
	l2.ERC1155Gateway.RegisterFinalizeDepositERC1155(func(vLog *types.Log, data *gateway.L2ERC1155GatewayFinalizeDepositERC1155Event) error {
		controller.ERC1155EventsTotal.WithLabelValues(l2.chainName, orm.L2FinalizeDepositERC1155.String()).Inc()
		l2.erc1155Events = append(l2.erc1155Events, newL2ERC1155Event(orm.L2FinalizeDepositERC1155, vLog, data.L1Token, data.L2Token, data.TokenID, data.Amount))
		return nil
	})
}

func (l2 *l2Contracts) normalTransferCheck(tp orm.EventType, txHash string, amount *big.Int) {
	event, exist := l2.transferEvents[txHash]
	if !exist {
		data, _ := json.Marshal(event)
		go controller.SlackNotify(
			fmt.Sprintf("can't find the %s matched transfer tx, tx_hash: %s, content: %s",
				tp.String(), txHash, string(data)),
		)
	} else if event.Value.Cmp(amount) != 0 {
		data, _ := json.Marshal(event)
		go controller.SlackNotify(
			fmt.Sprintf(
				"te %s transfer value cann't match tx, tx_hash: %s, expect_value: %s, actual_value: %s, content: %s",
				tp.String(), txHash, amount.String(), event.Value.String(), string(data)),
		)
		delete(l2.transferEvents, txHash)
	}
}

func (l2 *l2Contracts) abnormalTransferCheck() {
	// unexpect mint or burn operation.
	for txHash, event := range l2.transferEvents {
		switch event.To {
		case l2.cfg.WETHGateway:
			fallthrough
		case l2.cfg.DAIGateway:
			fallthrough
		case l2.cfg.StandardERC20Gateway:
			fallthrough
		case l2.cfg.CustomERC20Gateway:
			data, _ := json.Marshal(event)
			go controller.SlackNotify(
				fmt.Sprintf("unexpect mint tx from 0x000...000 address, tx_hash: %x, content: %s",
					txHash, string(data)),
			)
		}
		switch event.From {
		case l2.cfg.WETHGateway:
			fallthrough
		case l2.cfg.DAIGateway:
			fallthrough
		case l2.cfg.StandardERC20Gateway:
			fallthrough
		case l2.cfg.CustomERC20Gateway:
			data, _ := json.Marshal(event)
			go controller.SlackNotify(
				fmt.Sprintf("unexpect burn tx from 0x000...000 address, tx_hash: %x, content: %s",
					txHash, string(data)),
			)
		}
	}
}

func (l2 *l2Contracts) storeGatewayEvents(ctx context.Context, start, end uint64) error {
	// store l2 eth events.
	if len(l2.ethEvents) > 0 {
		startBls, err := l2.client.BalanceAt(ctx, l2.cfg.ScrollMessenger, big.NewInt(0).SetUint64(start))
		if err != nil {
			return err
		}

		var bls = big.NewInt(0)
		for _, event := range l2.ethEvents {
			if msgHash, exist := l2.txHashMsgHash[event.TxHash]; exist {
				event.MsgHash = msgHash.String()
			}
			bls.Add(bls, event.Amount)
		}

		endBls, err := l2.client.BalanceAt(ctx, l2.cfg.ScrollMessenger, big.NewInt(0).SetUint64(end))
		if err != nil {
			return err
		}
		endBls.Add(endBls, bls)
		if endBls.Cmp(startBls) != 0 {
			go controller.SlackNotify(
				fmt.Sprintf("the eth balance doesn't match, start: %d, end: %d, expect_balance: %s, actual_balance: %s",
					start, end, endBls.String(), startBls.String()),
			)
		}

		if err := l2.tx.Save(l2.ethEvents).Error; err != nil {
			return err
		}
	}

	// store l2 erc20 events.
	if len(l2.erc20Events) > 0 {
		for _, event := range l2.erc20Events {
			if msgHash, exist := l2.txHashMsgHash[event.TxHash]; exist {
				event.MsgHash = msgHash.String()
			}
			l2.normalTransferCheck(event.Type, event.TxHash, event.Amount)
		}
		if err := l2.tx.Save(l2.erc20Events).Error; err != nil {
			return err
		}
	}

	// store l2 err721 events.
	if len(l2.erc721Events) > 0 {
		for _, event := range l2.erc721Events {
			if msgHash, exist := l2.txHashMsgHash[event.TxHash]; exist {
				event.MsgHash = msgHash.String()
			}
			l2.normalTransferCheck(event.Type, event.TxHash, event.TokenID)
		}
		if err := l2.tx.Save(l2.erc721Events).Error; err != nil {
			return err
		}
	}

	// store l2 erc1155 events.
	if len(l2.erc1155Events) > 0 {
		for _, event := range l2.erc1155Events {
			if msgHash, exist := l2.txHashMsgHash[event.TxHash]; exist {
				event.MsgHash = msgHash.String()
			}
		}
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
