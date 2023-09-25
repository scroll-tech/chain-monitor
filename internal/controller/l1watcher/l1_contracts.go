package l1watcher

import (
	"context"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"gorm.io/gorm"

	"chain-monitor/bytecode"
	"chain-monitor/bytecode/scroll/L1"
	"chain-monitor/bytecode/scroll/L1/gateway"
	"chain-monitor/bytecode/scroll/L1/rollup"
	"chain-monitor/bytecode/scroll/token"
	"chain-monitor/internal/config"
	"chain-monitor/internal/controller"
	"chain-monitor/internal/orm"
)

type l1Contracts struct {
	cfg       *config.L1Contracts
	tx        *gorm.DB
	client    *ethclient.Client
	chainName string

	txHashMsgHash map[string]common.Hash
	ethEvents     []*orm.L1ETHEvent
	erc20Events   []*orm.L1ERC20Event
	erc721Events  []*orm.L1ERC721Event
	erc1155Events []*orm.L1ERC1155Event

	ETHGateway           *gateway.L1ETHGateway
	WETHGateway          *gateway.L1WETHGateway
	DAIGateway           *gateway.L1DAIGateway
	StandardERC20Gateway *gateway.L1StandardERC20Gateway
	CustomERC20Gateway   *gateway.L1CustomERC20Gateway
	ERC721Gateway        *gateway.L1ERC721Gateway
	ERC1155Gateway       *gateway.L1ERC1155Gateway

	ScrollChain     *rollup.ScrollChain
	ScrollMessenger *L1.L1ScrollMessenger
	// MessageQueue    *rollup.L1MessageQueue

	// this fields are used check balance.
	checkBalance     bool
	latestETHBalance *big.Int
	transferEvents   map[string]*token.IERC20TransferEvent
	iERC20           *token.IERC20

	gatewayFilter   *bytecode.ContractsFilter
	depositFilter   *bytecode.ContractsFilter
	fWithdrawFilter *bytecode.ContractsFilter
}

func newL1Contracts(client *ethclient.Client, cfg *config.L1Contracts) (*l1Contracts, error) {
	var (
		cts = &l1Contracts{
			cfg:           cfg,
			client:        client,
			chainName:     "l1_chain",
			txHashMsgHash: map[string]common.Hash{},
			ethEvents:     []*orm.L1ETHEvent{},
			erc20Events:   []*orm.L1ERC20Event{},
			erc721Events:  []*orm.L1ERC721Event{},
			erc1155Events: []*orm.L1ERC1155Event{},
		}
		err error
	)
	cts.ScrollMessenger, err = L1.NewL1ScrollMessenger(cfg.ScrollMessenger, client)
	if err != nil {
		return nil, err
	}
	//cts.MessageQueue, err = rollup.NewL1MessageQueue(cfg.MessageQueue, client)
	//if err != nil {
	//	return nil, err
	//}
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
	cts.iERC20, err = token.NewIERC20(common.Address{}, client)
	if err != nil {
		return nil, err
	}

	cts.gatewayFilter = bytecode.NewContractsFilter(nil, []bytecode.ContractAPI{
		cts.ScrollMessenger,
		// cts.MessageQueue,
		cts.ETHGateway,
		cts.DAIGateway,
		cts.WETHGateway,
		cts.StandardERC20Gateway,
		cts.CustomERC20Gateway,
		cts.ERC721Gateway,
		cts.ERC1155Gateway,
		cts.ScrollChain,
	}...)
	cts.depositFilter = bytecode.NewContractsFilter([][]common.Hash{
		{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")},
		{},
		{
			common.BytesToHash(cfg.DAIGateway[:]),
			common.BytesToHash(cfg.WETHGateway[:]),
			common.BytesToHash(cfg.StandardERC20Gateway[:]),
			common.BytesToHash(cfg.CustomERC20Gateway[:]),
			common.BytesToHash(cfg.ERC721Gateway[:]),
		},
	})
	cts.fWithdrawFilter = bytecode.NewContractsFilter([][]common.Hash{
		{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")},
		{
			common.BytesToHash(cfg.DAIGateway[:]),
			common.BytesToHash(cfg.WETHGateway[:]),
			common.BytesToHash(cfg.StandardERC20Gateway[:]),
			common.BytesToHash(cfg.CustomERC20Gateway[:]),
			common.BytesToHash(cfg.ERC721Gateway[:]),
		},
	})

	cts.registerGatewayHandlers()
	cts.registerMessengerHandlers()
	cts.registerScrollHandlers()
	cts.registerTransfer()

	return cts, nil
}

func (l1 *l1Contracts) clean() {
	l1.txHashMsgHash = map[string]common.Hash{}
	l1.transferEvents = map[string]*token.IERC20TransferEvent{}
	l1.ethEvents = l1.ethEvents[:0]
	l1.erc20Events = l1.erc20Events[:0]
	l1.erc721Events = l1.erc721Events[:0]
	l1.erc1155Events = l1.erc1155Events[:0]
}

func (l1 *l1Contracts) ParseL1Events(ctx context.Context, db *gorm.DB, start, end uint64) (int, error) {
	l1.clean()
	l1.tx = db.Begin().WithContext(ctx)

	// init eth balance.
	if l1.latestETHBalance == nil {
		balance, err := l1.client.BalanceAt(ctx, l1.cfg.ScrollMessenger, big.NewInt(0).SetUint64(start-1))
		if err != nil {
			return 0, err
		}
		l1.latestETHBalance = balance
	}

	// Parse gateway logs.
	count, err := l1.gatewayFilter.GetLogs(ctx, l1.client, start, end, l1.gatewayFilter.ParseLogs)
	if err != nil {
		controller.ParseLogsFailureTotal.WithLabelValues(l1.chainName).Inc()
		l1.tx.Rollback()
		return 0, err
	}

	// Parse finalizeWithdraw transfer event logs.
	_, err = l1.fWithdrawFilter.GetLogs(ctx, l1.client, start, end, l1.parseTransferLogs)
	if err != nil {
		controller.ParseLogsFailureTotal.WithLabelValues(l1.chainName).Inc()
		l1.tx.Rollback()
		return 0, err
	}

	// Parse deposit transfer event logs.
	_, err = l1.depositFilter.GetLogs(ctx, l1.client, start, end, l1.parseTransferLogs)
	if err != nil {
		controller.ParseLogsFailureTotal.WithLabelValues(l1.chainName).Inc()
		l1.tx.Rollback()
		return 0, err
	}

	if l1.checkBalance {
		// Check erc20 balance.
		l1.checkERC20Balance()
		// Check eth balance.
		if err = l1.checkETHBalance(ctx, end); err != nil {
			l1.tx.Rollback()
			return 0, err
		}
	}

	// store l1chain gateway events.
	if err = l1.storeGatewayEvents(); err != nil {
		l1.tx.Rollback()
		return 0, err
	}

	// Store l1 confirm.
	var l1Confirms = make([]orm.L1ChainConfirm, 0, end-start+1)
	for number := start; number <= end; number++ {
		l1Confirms = append(l1Confirms, orm.L1ChainConfirm{
			Number: number,
		})
	}
	if err = l1.tx.Save(l1Confirms).Error; err != nil {
		l1.tx.Rollback()
		return 0, err
	}

	// store the latest l1 block numbers
	err = l1.tx.Save(&orm.L1Block{
		Number:     end,
		EventCount: count,
	}).Error
	if err != nil {
		l1.tx.Rollback()
		return 0, err
	}

	if err = l1.tx.Commit().Error; err != nil {
		l1.tx.Rollback()
		return 0, err
	}
	return count, nil
}
