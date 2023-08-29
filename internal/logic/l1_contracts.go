package logic

import (
	"context"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"gorm.io/gorm"

	"chain-monitor/bytecode"
	"chain-monitor/bytecode/scroll/L1"
	"chain-monitor/bytecode/scroll/L1/gateway"
	"chain-monitor/bytecode/scroll/L1/rollup"
	"chain-monitor/internal/config"
	"chain-monitor/orm"
)

type L1Contracts struct {
	db     *gorm.DB
	tx     *gorm.DB
	client *ethclient.Client

	messengerEvents map[string]common.Hash
	ethEvents       []*orm.L1ETHEvent
	erc20Events     []*orm.L1ERC20Event
	erc721Events    []*orm.L1ERC721Event
	erc1155Events   []*orm.L1ERC1155Event

	ETHGateway           *gateway.L1ETHGateway
	WETHGateway          *gateway.L1WETHGateway
	DAIGateway           *gateway.L1DAIGateway
	StandardERC20Gateway *gateway.L1StandardERC20Gateway
	CustomERC20Gateway   *gateway.L1CustomERC20Gateway
	ERC721Gateway        *gateway.L1ERC721Gateway
	ERC1155Gateway       *gateway.L1ERC1155Gateway

	ScrollChain     *rollup.ScrollChain
	ScrollMessenger *L1.L1ScrollMessenger

	filter *bytecode.ContractsFilter
}

func NewL1Contracts(client *ethclient.Client, db *gorm.DB, cfg *config.L1Contracts) (*L1Contracts, error) {
	var (
		cts = &L1Contracts{
			client:          client,
			db:              db,
			messengerEvents: map[string]common.Hash{},
			ethEvents:       []*orm.L1ETHEvent{},
			erc20Events:     []*orm.L1ERC20Event{},
			erc721Events:    []*orm.L1ERC721Event{},
			erc1155Events:   []*orm.L1ERC1155Event{},
		}
		err error
	)
	cts.ScrollMessenger, err = L1.NewL1ScrollMessenger(cfg.ScrollMessenger, client)
	if err != nil {
		return nil, err
	}
	cts.ETHGateway, err = gateway.NewL1ETHGateway(cfg.ETHGateway, client)
	if err != nil {
		return nil, err
	}
	cts.WETHGateway, err = gateway.NewL1WETHGateway(cfg.WETHGateway, client)
	if err != nil {
		return nil, err
	}
	cts.StandardERC20Gateway, err = gateway.NewL1StandardERC20Gateway(cfg.StandardERC20Gateway, client)
	if err != nil {
		return nil, err
	}
	cts.CustomERC20Gateway, err = gateway.NewL1CustomERC20Gateway(cfg.CustomERC20Gateway, client)
	if err != nil {
		return nil, err
	}
	cts.ERC721Gateway, err = gateway.NewL1ERC721Gateway(cfg.ERC721Gateway, client)
	if err != nil {
		return nil, err
	}
	cts.ERC1155Gateway, err = gateway.NewL1ERC1155Gateway(cfg.ERC1155Gateway, client)
	if err != nil {
		return nil, err
	}
	cts.DAIGateway, err = gateway.NewL1DAIGateway(cfg.DAIGateway, client)
	if err != nil {
		return nil, err
	}
	cts.ScrollChain, err = rollup.NewScrollChain(cfg.L1ScrollChain, client)
	if err != nil {
		return nil, err
	}

	cts.filter = bytecode.NewContractsFilter([]bytecode.ContractAPI{
		cts.ScrollMessenger,
		cts.ETHGateway,
		cts.DAIGateway,
		cts.WETHGateway,
		cts.StandardERC20Gateway,
		cts.CustomERC20Gateway,
		cts.ERC721Gateway,
		cts.ERC1155Gateway,
		cts.ScrollChain,
	}...)

	cts.registerEventHandlers()

	return cts, nil
}

func (l1 *L1Contracts) clean() {
	l1.messengerEvents = map[string]common.Hash{}
	l1.ethEvents = l1.ethEvents[:0]
	l1.erc20Events = l1.erc20Events[:0]
	l1.erc721Events = l1.erc721Events[:0]
	l1.erc1155Events = l1.erc1155Events[:0]
}

func (l1 *L1Contracts) ParseL1Events(ctx context.Context, start, end uint64) error {
	l1.clean()
	l1.tx = l1.db.Begin()
	count, err := l1.filter.ParseLogs(ctx, l1.client, start, end)
	if err != nil {
		l1.tx.Rollback()
		return err
	}

	// store l1 eth events.
	for i := 0; i < len(l1.ethEvents); i++ {
		event := l1.ethEvents[i]
		if msgHash, exist := l1.messengerEvents[event.TxHash]; exist {
			event.MessageHash = msgHash.String()
		}
	}
	if len(l1.ethEvents) > 0 {
		if err := l1.tx.Save(l1.ethEvents).Error; err != nil {
			l1.tx.Rollback()
			return err
		}
	}

	// store l1 erc20 events.
	for i := 0; i < len(l1.erc20Events); i++ {
		event := l1.erc20Events[i]
		if msgHash, exist := l1.messengerEvents[event.TxHash]; exist {
			event.MessageHash = msgHash.String()
		}
	}
	if len(l1.erc20Events) > 0 {
		if err := l1.tx.Save(l1.erc20Events).Error; err != nil {
			l1.tx.Rollback()
			return err
		}
	}

	// store l1 err721 events.
	for i := 0; i < len(l1.erc721Events); i++ {
		event := l1.erc721Events[i]
		if msgHash, exist := l1.messengerEvents[event.TxHash]; exist {
			event.MessageHash = msgHash.String()

		}
	}
	if len(l1.erc721Events) > 0 {
		if err := l1.tx.Save(l1.erc721Events).Error; err != nil {
			l1.tx.Rollback()
			return err
		}
	}

	// store l1 erc1155 events.
	for i := 0; i < len(l1.erc1155Events); i++ {
		event := l1.erc1155Events[i]
		if msgHash, exist := l1.messengerEvents[event.TxHash]; exist {
			event.MessageHash = msgHash.String()
		}
	}
	if len(l1.erc1155Events) > 0 {
		if err := l1.tx.Save(l1.erc1155Events).Error; err != nil {
			l1.tx.Rollback()
			return err
		}
	}

	// store the latest l1 block number
	err = l1.tx.Save(&orm.L1Block{
		Number:     end,
		EventCount: count,
	}).Error
	if err != nil {
		l1.tx.Rollback()
		return err
	}

	if err = l1.tx.Commit().Error; err != nil {
		l1.tx.Rollback()
		return err
	}
	return nil
}

func (l1 *L1Contracts) registerEventHandlers() {
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

	l1.ScrollMessenger.RegisterSentMessage(func(vLog *types.Log, data *L1.L1ScrollMessengerSentMessageEvent) error {
		msgHash := crypto.Keccak256Hash(data.Message)
		l1.messengerEvents[vLog.TxHash.String()] = msgHash
		return nil //orm.SaveL1Messenger(l1.tx, orm.L1SentMessage, vLog, msgHash)
	})
	l1.ScrollMessenger.RegisterRelayedMessage(func(vLog *types.Log, data *L1.L1ScrollMessengerRelayedMessageEvent) error {
		msgHash := common.BytesToHash(data.MessageHash[:])
		l1.messengerEvents[vLog.TxHash.String()] = msgHash
		return nil //orm.SaveL1Messenger(l1.tx, orm.L1RelayedMessage, vLog, msgHash)
	})
	l1.ScrollMessenger.RegisterFailedRelayedMessage(func(vLog *types.Log, data *L1.L1ScrollMessengerFailedRelayedMessageEvent) error {
		msgHash := common.BytesToHash(data.MessageHash[:])
		l1.messengerEvents[vLog.TxHash.String()] = msgHash
		return nil //orm.SaveL1Messenger(l1.tx, orm.L1FailedRelayedMessage, vLog, msgHash)
	})

	l1.ScrollChain.RegisterCommitBatch(func(vLog *types.Log, data *rollup.ScrollChainCommitBatchEvent) error {
		l1Tx, _, err := l1.client.TransactionByHash(context.Background(), vLog.TxHash)
		if err != nil {
			return err
		}
		scrollBatch, err := bytecode.GetBatchRangeFromCalldataV2(l1.ScrollChain.ABI, l1Tx.Data())
		if err != nil {
			return err
		}
		return l1.tx.Create(&orm.ScrollChainEvent{
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
		eventMsg := orm.ScrollChainEvent{
			BatchHash:      common.BytesToHash(data.BatchHash[:]).String(),
			FinalizeNumber: vLog.BlockNumber,
			FinalizeHash:   vLog.TxHash.String(),
			Status:         orm.BatchFinalize,
		}
		return l1.tx.Select("finalize_number", "finalize_hash", "status").Updates(&eventMsg).Error
	})
	l1.ScrollChain.RegisterRevertBatch(func(vLog *types.Log, data *rollup.ScrollChainRevertBatchEvent) error {
		eventMsg := orm.ScrollChainEvent{
			BatchHash:    common.BytesToHash(data.BatchHash[:]).String(),
			RevertNumber: vLog.BlockNumber,
			RevertHash:   vLog.TxHash.String(),
			Status:       orm.BatchRevert,
		}
		return l1.tx.Select("revert_number", "revert_hash", "status").Updates(&eventMsg).Error
	})
}

func newL1ETHEvent(eventType orm.EventType, vLog *types.Log, from, to common.Address, amount *big.Int) *orm.L1ETHEvent {
	return &orm.L1ETHEvent{
		TxHead: orm.NewTxHead(vLog, eventType),
		From:   from.String(),
		To:     to.String(),
		Amount: amount,
	}
}

func newL1ETH20Event(eventType orm.EventType, vLog *types.Log, l1Token, l2Token, from, to common.Address, amount *big.Int) *orm.L1ERC20Event {
	return &orm.L1ERC20Event{
		TxHead:  orm.NewTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		From:    from.String(),
		To:      to.String(),
		Amount:  amount,
	}
}

func newL1ERC721Event(eventType orm.EventType, vLog *types.Log, l1Token, l2Token, from, to common.Address, tokenId *big.Int) *orm.L1ERC721Event {
	return &orm.L1ERC721Event{
		TxHead:  orm.NewTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		From:    from.String(),
		To:      to.String(),
		TokenId: tokenId,
	}
}

func newL1ERC1155Event(eventType orm.EventType, vLog *types.Log, l1Token, l2Token, from, to common.Address, tokenId, amount *big.Int) *orm.L1ERC1155Event {
	return &orm.L1ERC1155Event{
		TxHead:  orm.NewTxHead(vLog, eventType),
		L1Token: l1Token.String(),
		L2Token: l2Token.String(),
		From:    from.String(),
		To:      to.String(),
		TokenId: tokenId,
		Amount:  amount,
	}
}
