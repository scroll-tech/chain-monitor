package l2watcher

import (
	"context"
	"fmt"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/rpc"
	"gorm.io/gorm"

	"chain-monitor/bytecode"
	"chain-monitor/bytecode/scroll/L2"
	"chain-monitor/bytecode/scroll/L2/gateway"
	"chain-monitor/bytecode/scroll/L2/predeploys"
	"chain-monitor/bytecode/scroll/token"
	"chain-monitor/internal/config"
	"chain-monitor/internal/controller"
	"chain-monitor/internal/orm"
	"chain-monitor/internal/utils/msgproof"
)

type l2Contracts struct {
	tx           *gorm.DB
	cfg          *config.L2Contracts
	l1gatewayCfg *config.Gateway
	chainName    string

	rpcCli *rpc.Client
	client *ethclient.Client

	withdraw *msgproof.WithdrawTrie

	msgSentEvents map[uint64][]*orm.L2MessengerEvent
	ethEvents     []*orm.L2ETHEvent
	erc20Events   []*orm.L2ERC20Event
	erc721Events  []*orm.L2ERC721Event
	erc1155Events []*orm.L2ERC1155Event

	gatewayAPIs          []bytecode.ContractAPI
	ETHGateway           *gateway.L2ETHGateway
	WETHGateway          *gateway.L2WETHGateway
	DAIGateway           *gateway.L2DAIGateway
	StandardERC20Gateway *gateway.L2StandardERC20Gateway
	CustomERC20Gateway   *gateway.L2CustomERC20Gateway
	USDCERC20Gateway     *gateway.L2CustomERC20Gateway
	LIDOERC20Gateway     *gateway.L2CustomERC20Gateway
	ERC721Gateway        *gateway.L2ERC721Gateway
	ERC1155Gateway       *gateway.L2ERC1155Gateway

	ScrollMessenger *L2.L2ScrollMessenger
	MessageQueue    *predeploys.L2MessageQueue

	// this fields are used check balance.
	transferEvents map[string]*token.IERC20TransferEvent
	iERC20         *token.IERC20

	l2Confirms map[uint64]*orm.L2ChainConfirm

	gatewayFilter  *bytecode.ContractsFilter
	fDepositFilter *bytecode.ContractsFilter
	withdrawFilter *bytecode.ContractsFilter
}

func newL2Contracts(l2chainURL string, db *gorm.DB, cfg *config.L2Contracts) (*l2Contracts, error) {
	rpcCli, err := rpc.Dial(l2chainURL)
	if err != nil {
		return nil, err
	}

	var (
		client = ethclient.NewClient(rpcCli)
		cts    = &l2Contracts{
			rpcCli:    rpcCli,
			client:    client,
			cfg:       cfg,
			chainName: "l2_chain",
			withdraw:  msgproof.NewWithdrawTrie(),
		}
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
	cts.USDCERC20Gateway, err = gateway.NewL2CustomERC20Gateway(cfg.USDCGateway, client)
	if err != nil {
		return nil, err
	}
	cts.LIDOERC20Gateway, err = gateway.NewL2CustomERC20Gateway(cfg.LIDOGateway, client)
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
	cts.iERC20, err = token.NewIERC20(common.Address{}, client)
	if err != nil {
		return nil, err
	}

	cts.gatewayAPIs = []bytecode.ContractAPI{
		cts.ETHGateway,
		cts.WETHGateway,
		cts.StandardERC20Gateway,
		cts.CustomERC20Gateway,
		cts.ERC721Gateway,
		cts.ERC1155Gateway,
	}
	if cfg.USDCGateway != (common.Address{}) {
		cts.gatewayAPIs = append(cts.gatewayAPIs, cts.USDCERC20Gateway)
	}
	if cfg.LIDOGateway != (common.Address{}) {
		cts.gatewayAPIs = append(cts.gatewayAPIs, cts.LIDOERC20Gateway)
	}
	if cfg.DAIGateway != (common.Address{}) {
		cts.gatewayAPIs = append(cts.gatewayAPIs, cts.DAIGateway)
	}
	cts.gatewayFilter = bytecode.NewContractsFilter(nil, append(cts.gatewayAPIs, []bytecode.ContractAPI{
		cts.ScrollMessenger,
		// cts.MessageQueue,
	}...)...)

	// Filter the Transfer event ID is 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
	// The Topic[1] value should be 0x000000.
	cts.fDepositFilter = bytecode.NewContractsFilter([][]common.Hash{
		{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")},
		{common.BigToHash(big.NewInt(0))},
	})
	// Filter the Transfer event ID is 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
	// The Topic[2] value should be 0x000000.
	cts.withdrawFilter = bytecode.NewContractsFilter([][]common.Hash{
		{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")},
		{},
		{common.BigToHash(big.NewInt(0))},
	})

	// Init withdraw root.
	if err = cts.initWithdraw(db); err != nil {
		return nil, err
	}

	cts.registerGatewayHandlers()
	cts.registerMessengerHandlers()
	cts.registerTransfer()

	return cts, nil
}

func (l2 *l2Contracts) initWithdraw(db *gorm.DB) error {
	tx := db.Where("type = ? AND msg_proof != ''", orm.L2SentMessage).Order("number DESC")
	var msg orm.L2MessengerEvent
	err := tx.Last(&msg).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return err
	}
	if err != nil {
		return nil
	}
	l2.withdraw.Initialize(msg.MsgNonce, common.HexToHash(msg.MsgHash), common.Hex2Bytes(msg.MsgProof))

	// Check the current withdraw root is right or not.
	expectRoot, err := l2.withdrawRoot(context.Background(), msg.Number)
	if err != nil {
		return err
	}
	actualRoot := l2.withdraw.MessageRoot()
	if expectRoot != actualRoot {
		return fmt.Errorf("withdraw root is not right, withdraw init failed, numbers: %d, expect_root: %s, actual_root: %s", msg.Number, expectRoot.String(), actualRoot.String())
	}
	return nil
}

func (l2 *l2Contracts) clean() {
	l2.msgSentEvents = map[uint64][]*orm.L2MessengerEvent{}
	l2.transferEvents = map[string]*token.IERC20TransferEvent{}
	l2.l2Confirms = map[uint64]*orm.L2ChainConfirm{}
	l2.ethEvents = l2.ethEvents[:0]
	l2.erc20Events = l2.erc20Events[:0]
	l2.erc721Events = l2.erc721Events[:0]
	l2.erc1155Events = l2.erc1155Events[:0]
}

func (l2 *l2Contracts) ParseL2Events(ctx context.Context, db *gorm.DB, start, end uint64) (int, error) {
	l2.tx = db.Begin().WithContext(ctx)

	// parse l2 event logs.
	count, err := l2.parseL2Events(ctx, start, end)
	if err != nil {
		l2.tx.Rollback()
		return 0, err
	}

	// commit db.
	if err = l2.tx.Commit().Error; err != nil {
		l2.tx.Rollback()
		return 0, err
	}

	return count, nil
}

func (l2 *l2Contracts) parseL2Events(ctx context.Context, start, end uint64) (int, error) {
	l2.clean()

	// Parse gateway logs
	count, err := l2.gatewayFilter.GetLogs(ctx, l2.client, start, end, l2.gatewayFilter.ParseLogs)
	if err != nil {
		controller.ParseLogsFailureTotal.WithLabelValues(l2.chainName).Inc()
		return 0, err
	}

	// Parse finalizeDeposit transfer event logs.
	_, err = l2.fDepositFilter.GetLogs(ctx, l2.client, start, end, l2.parseTransferLogs)
	if err != nil {
		controller.ParseLogsFailureTotal.WithLabelValues(l2.chainName).Inc()
		return 0, err
	}

	// Parse withdraw transfer event logs.
	_, err = l2.withdrawFilter.GetLogs(ctx, l2.client, start, end, l2.parseTransferLogs)
	if err != nil {
		controller.ParseLogsFailureTotal.WithLabelValues(l2.chainName).Inc()
		return 0, err
	}

	// Create l2Confirms
	for number := start; number <= end; number++ {
		l2.l2Confirms[number] = &orm.L2ChainConfirm{
			Number:        number,
			BalanceStatus: true,
		}
	}

	// If cross tx not used gateway contract, then the function can mock a gateway event.
	if err = l2.gatewayEvents(); err != nil {
		return 0, err
	}

	// Check balance.
	if err = l2.checkL2Balance(ctx, start, end); err != nil {
		return 0, err
	}

	// Check eth balance.
	if err = l2.storeGatewayEvents(); err != nil {
		return 0, err
	}

	// store l2Messenger sentMessenger events.
	if err = l2.storeMessengerEvents(ctx, start, end); err != nil {
		return 0, err
	}

	// Check l2chain confirms.
	if err = l2.storeL1ChainConfirms(ctx); err != nil {
		return 0, err
	}

	err = l2.tx.Save(&orm.L2Block{
		Number:     end,
		EventCount: count,
	}).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (l2 *l2Contracts) withdrawRoot(ctx context.Context, number uint64) (common.Hash, error) {
	data, err := l2.client.StorageAt(
		ctx,
		l2.cfg.MessageQueue,
		common.BigToHash(big.NewInt(0)), big.NewInt(0).SetUint64(number),
	)
	if err != nil {
		return [32]byte{}, err
	}
	return common.BytesToHash(data), nil
}
