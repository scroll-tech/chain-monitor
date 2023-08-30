package logic

import (
	"context"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"gorm.io/gorm"

	"chain-monitor/bytecode"
	"chain-monitor/bytecode/scroll/L2"
	"chain-monitor/bytecode/scroll/L2/gateway"
	"chain-monitor/bytecode/scroll/L2/predeploys"
	"chain-monitor/internal/config"
	"chain-monitor/orm"
)

type L2Contracts struct {
	db     *gorm.DB
	tx     *gorm.DB
	client *ethclient.Client

	messengerEvents map[string]common.Hash
	ethEvents       []*orm.L2ETHEvent
	erc20Events     []*orm.L2ERC20Event
	erc721Events    []*orm.L2ERC721Event
	erc1155Events   []*orm.L2ERC1155Event

	ETHGateway           *gateway.L2ETHGateway
	WETHGateway          *gateway.L2WETHGateway
	DAIGateway           *gateway.L2DAIGateway
	StandardERC20Gateway *gateway.L2StandardERC20Gateway
	CustomERC20Gateway   *gateway.L2CustomERC20Gateway
	ERC721Gateway        *gateway.L2ERC721Gateway
	ERC1155Gateway       *gateway.L2ERC1155Gateway

	ScrollMessenger *L2.L2ScrollMessenger
	MessageQueue    *predeploys.L2MessageQueue

	filter *bytecode.ContractsFilter
}

func NewL2Contracts(client *ethclient.Client, db *gorm.DB, cfg *config.Gateway) (*L2Contracts, error) {
	var (
		cts = &L2Contracts{
			client: client,
			db:     db,
		}
		err error
	)
	cts.ETHGateway, err = gateway.NewL2ETHGateway(cfg.ETHGateway, client)
	if err != nil {
		return nil, err
	}
	cts.WETHGateway, err = gateway.NewL2WETHGateway(cfg.WETHGateway, client)
	if err != nil {
		return nil, err
	}
	cts.DAIGateway, err = gateway.NewL2DAIGateway(cfg.DAIGateway, client)
	if err != nil {
		return nil, err
	}
	cts.StandardERC20Gateway, err = gateway.NewL2StandardERC20Gateway(cfg.StandardERC20Gateway, client)
	if err != nil {
		return nil, err
	}
	cts.CustomERC20Gateway, err = gateway.NewL2CustomERC20Gateway(cfg.CustomERC20Gateway, client)
	if err != nil {
		return nil, err
	}
	cts.ERC721Gateway, err = gateway.NewL2ERC721Gateway(cfg.ERC721Gateway, client)
	if err != nil {
		return nil, err
	}
	cts.ERC1155Gateway, err = gateway.NewL2ERC1155Gateway(cfg.ERC1155Gateway, client)
	if err != nil {
		return nil, err
	}

	cts.ScrollMessenger, err = L2.NewL2ScrollMessenger(cfg.ScrollMessenger, client)
	if err != nil {
		return nil, err
	}
	cts.MessageQueue, err = predeploys.NewL2MessageQueue(cfg.MessageQueue, client)
	if err != nil {
		return nil, err
	}

	cts.filter = bytecode.NewContractsFilter([]bytecode.ContractAPI{
		cts.ScrollMessenger,
		cts.MessageQueue,

		cts.ETHGateway,
		cts.WETHGateway,
		cts.DAIGateway,
		cts.StandardERC20Gateway,
		cts.CustomERC20Gateway,
		cts.ERC721Gateway,
		cts.ERC1155Gateway,
	}...)

	cts.registerDepositHandlers()

	return cts, nil
}

func (l2 *L2Contracts) clean() {
	l2.messengerEvents = map[string]common.Hash{}
	l2.ethEvents = l2.ethEvents[:0]
	l2.erc20Events = l2.erc20Events[:0]
	l2.erc721Events = l2.erc721Events[:0]
	l2.erc1155Events = l2.erc1155Events[:0]
}

func (l2 *L2Contracts) ParseL2Events(ctx context.Context, start, end uint64) error {
	l2.clean()
	l2.tx = l2.db.Begin()
	count, err := l2.filter.ParseLogs(ctx, l2.client, start, end)
	if err != nil {
		l2.tx.Rollback()
		return err
	}

	// store l2 eth events.
	for i := 0; i < len(l2.ethEvents); i++ {
		event := l2.ethEvents[i]
		if msgHash, exist := l2.messengerEvents[event.TxHash]; exist {
			event.MessageHash = msgHash.String()
		}
	}
	if len(l2.ethEvents) > 0 {
		if err := l2.tx.Save(l2.ethEvents).Error; err != nil {
			l2.tx.Rollback()
			return err
		}
	}

	// store l2 erc20 events.
	for i := 0; i < len(l2.erc20Events); i++ {
		event := l2.erc20Events[i]
		if msgHash, exist := l2.messengerEvents[event.TxHash]; exist {
			event.MessageHash = msgHash.String()
		}
	}
	if len(l2.erc20Events) > 0 {
		if err := l2.tx.Save(l2.erc20Events).Error; err != nil {
			l2.tx.Rollback()
			return err
		}
	}

	// store l2 err721 events.
	for i := 0; i < len(l2.erc721Events); i++ {
		event := l2.erc721Events[i]
		if msgHash, exist := l2.messengerEvents[event.TxHash]; exist {
			event.MessageHash = msgHash.String()

		}
	}
	if len(l2.erc721Events) > 0 {
		if err := l2.tx.Save(l2.erc721Events).Error; err != nil {
			l2.tx.Rollback()
			return err
		}
	}

	// store l2 erc1155 events.
	for i := 0; i < len(l2.erc1155Events); i++ {
		event := l2.erc1155Events[i]
		if msgHash, exist := l2.messengerEvents[event.TxHash]; exist {
			event.MessageHash = msgHash.String()
		}
	}
	if len(l2.erc1155Events) > 0 {
		if err := l2.tx.Save(l2.erc1155Events).Error; err != nil {
			l2.tx.Rollback()
			return err
		}
	}

	err = l2.tx.Save(&orm.L2Block{
		Number:     end,
		EventCount: count,
	}).Error
	if err != nil {
		l2.tx.Rollback()
		return err
	}

	if err = l2.tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func (l2 *L2Contracts) registerDepositHandlers() {
	l2.ETHGateway.RegisterWithdrawETH(func(vLog *types.Log, data *gateway.L2ETHGatewayWithdrawETHEvent) error {
		l2.ethEvents = append(l2.ethEvents, newL2ETHEvent(orm.L2WithdrawETH, vLog, data.From, data.To, data.Amount))
		return nil
	})
	l2.ETHGateway.RegisterFinalizeDepositETH(func(vLog *types.Log, data *gateway.L2ETHGatewayFinalizeDepositETHEvent) error {
		l2.ethEvents = append(l2.ethEvents, newL2ETHEvent(orm.L2FinalizeDepositETH, vLog, data.From, data.To, data.Amount))
		return nil
	})
	l2.WETHGateway.RegisterWithdrawERC20(func(vLog *types.Log, data *gateway.L2WETHGatewayWithdrawERC20Event) error {
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2WithdrawWETH, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount))
		return nil
	})
	l2.WETHGateway.RegisterFinalizeDepositERC20(func(vLog *types.Log, data *gateway.L2WETHGatewayFinalizeDepositERC20Event) error {
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2FinalizeDepositWETH, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount))
		return nil
	})
	l2.DAIGateway.RegisterWithdrawERC20(func(vLog *types.Log, data *gateway.L2DAIGatewayWithdrawERC20Event) error {
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2WithdrawDAI, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount))
		return nil
	})
	l2.DAIGateway.RegisterFinalizeDepositERC20(func(vLog *types.Log, data *gateway.L2DAIGatewayFinalizeDepositERC20Event) error {
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2FinalizeDepositDAI, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount))
		return nil
	})

	l2.StandardERC20Gateway.RegisterWithdrawERC20(func(vLog *types.Log, data *gateway.L2StandardERC20GatewayWithdrawERC20Event) error {
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2WithdrawStandardERC20, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount))
		return nil
	})
	l2.StandardERC20Gateway.RegisterFinalizeDepositERC20(func(vLog *types.Log, data *gateway.L2StandardERC20GatewayFinalizeDepositERC20Event) error {
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2FinalizeDepositStandardERC20, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount))
		return nil
	})

	l2.CustomERC20Gateway.RegisterWithdrawERC20(func(vLog *types.Log, data *gateway.L2CustomERC20GatewayWithdrawERC20Event) error {
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2WithdrawCustomERC20, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount))
		return nil
	})
	l2.CustomERC20Gateway.RegisterFinalizeDepositERC20(func(vLog *types.Log, data *gateway.L2CustomERC20GatewayFinalizeDepositERC20Event) error {
		l2.erc20Events = append(l2.erc20Events, newL2ETH20Event(orm.L2FinalizeDepositCustomERC20, vLog, data.L1Token, data.L2Token, data.From, data.To, data.Amount))
		return nil
	})

	l2.ERC721Gateway.RegisterWithdrawERC721(func(vLog *types.Log, data *gateway.L2ERC721GatewayWithdrawERC721Event) error {
		l2.erc721Events = append(l2.erc721Events, newL2ERC721Event(orm.L2WithdrawERC721, vLog, data.L1Token, data.L2Token, data.From, data.To, data.TokenId))
		return nil
	})
	l2.ERC721Gateway.RegisterFinalizeDepositERC721(func(vLog *types.Log, data *gateway.L2ERC721GatewayFinalizeDepositERC721Event) error {
		l2.erc721Events = append(l2.erc721Events, newL2ERC721Event(orm.L2FinalizeDepositERC721, vLog, data.L1Token, data.L2Token, data.From, data.To, data.TokenId))
		return nil
	})
	l2.ERC1155Gateway.RegisterWithdrawERC1155(func(vLog *types.Log, data *gateway.L2ERC1155GatewayWithdrawERC1155Event) error {
		l2.erc1155Events = append(l2.erc1155Events, newL2ERC1155Event(orm.L2WithdrawERC1155, vLog, data.L1Token, data.L2Token, data.From, data.To, data.TokenId, data.Amount))
		return nil
	})
	l2.ERC1155Gateway.RegisterFinalizeDepositERC1155(func(vLog *types.Log, data *gateway.L2ERC1155GatewayFinalizeDepositERC1155Event) error {
		l2.erc1155Events = append(l2.erc1155Events, newL2ERC1155Event(orm.L2FinalizeDepositERC1155, vLog, data.L1Token, data.L2Token, data.From, data.To, data.TokenId, data.Amount))
		return nil
	})

	l2.MessageQueue.RegisterAppendMessage(func(vLog *types.Log, data *predeploys.L2MessageQueueAppendMessageEvent) error {
		l2.messengerEvents[vLog.TxHash.String()] = data.MessageHash
		return orm.SaveL2Messenger(l2.tx, orm.L2SentMessage, vLog, data.MessageHash)
	})
	l2.ScrollMessenger.RegisterRelayedMessage(func(vLog *types.Log, data *L2.L2ScrollMessengerRelayedMessageEvent) error {
		l2.messengerEvents[vLog.TxHash.String()] = data.MessageHash
		return orm.SaveL2Messenger(l2.tx, orm.L2RelayedMessage, vLog, data.MessageHash)
	})

	l2.ScrollMessenger.RegisterFailedRelayedMessage(func(vLog *types.Log, data *L2.L2ScrollMessengerFailedRelayedMessageEvent) error {
		l2.messengerEvents[vLog.TxHash.String()] = data.MessageHash
		return orm.SaveL2Messenger(l2.tx, orm.L2FailedRelayedMessage, vLog, data.MessageHash)
	})
}

func newL2ETHEvent(eventType orm.EventType, vLog *types.Log, from, to common.Address, amount *big.Int) *orm.L2ETHEvent {
	return &orm.L2ETHEvent{
		TxHead: orm.NewTxHead(vLog, eventType),
		From:   from.String(),
		To:     to.String(),
		Amount: amount,
	}
}

func newL2ETH20Event(eventType orm.EventType, vLog *types.Log, l1Token, l2Token, from, to common.Address, amount *big.Int) *orm.L2ERC20Event {
	return &orm.L2ERC20Event{
		TxHead:  orm.NewTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		From:    from.String(),
		To:      to.String(),
		Amount:  amount,
	}
}

func newL2ERC721Event(eventType orm.EventType, vLog *types.Log, l1Token, l2Token, from, to common.Address, tokenId *big.Int) *orm.L2ERC721Event {
	return &orm.L2ERC721Event{
		TxHead:  orm.NewTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		From:    from.String(),
		To:      to.String(),
		TokenId: tokenId,
	}
}

func newL2ERC1155Event(eventType orm.EventType, vLog *types.Log, l1Token, l2Token, from, to common.Address, tokenId, amount *big.Int) *orm.L2ERC1155Event {
	return &orm.L2ERC1155Event{
		TxHead:  orm.NewTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		From:    from.String(),
		To:      to.String(),
		TokenId: tokenId,
		Amount:  amount,
	}
}
