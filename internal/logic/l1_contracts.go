package logic

import (
	"context"

	"github.com/scroll-tech/go-ethereum/common"
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
	MessageQueue    *rollup.L1MessageQueue

	filter *bytecode.ContractsFilter
}

func NewL1Contracts(client *ethclient.Client, db *gorm.DB, cfg *config.L1Contracts) (*L1Contracts, error) {
	var (
		cts = &L1Contracts{
			client:        client,
			db:            db,
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
	cts.MessageQueue, err = rollup.NewL1MessageQueue(cfg.MessageQueue, client)
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
		cts.MessageQueue,
		cts.ETHGateway,
		cts.DAIGateway,
		cts.WETHGateway,
		cts.StandardERC20Gateway,
		cts.CustomERC20Gateway,
		cts.ERC721Gateway,
		cts.ERC1155Gateway,
		cts.ScrollChain,
	}...)

	cts.registerGatewayHandlers()
	cts.registerMessengerHandlers()
	cts.registerScrollHandlers()

	return cts, nil
}

func (l1 *L1Contracts) clean() {
	l1.txHashMsgHash = map[string]common.Hash{}
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

	// store l1chain gateway events.
	if err = l1.storeGatewayEvents(); err != nil {
		l1.tx.Rollback()
		return err
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
