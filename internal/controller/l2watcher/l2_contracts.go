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
	"chain-monitor/internal/config"
	"chain-monitor/internal/controller"
	"chain-monitor/internal/orm"
	"chain-monitor/internal/utils/msgproof"
)

type l2Contracts struct {
	tx        *gorm.DB
	cfg       *config.Gateway
	chainName string

	rpcCli *rpc.Client
	client *ethclient.Client

	monitorAPI controller.MonitorAPI

	withdraw *msgproof.WithdrawTrie

	txHashMsgHash map[string]common.Hash
	msgSentEvents map[uint64][]*orm.L2MessengerEvent
	ethEvents     []*orm.L2ETHEvent
	erc20Events   []*orm.L2ERC20Event
	erc721Events  []*orm.L2ERC721Event
	erc1155Events []*orm.L2ERC1155Event

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

func newL2Contracts(l2chainURL string, db *gorm.DB, cfg *config.Gateway) (*l2Contracts, error) {
	rpcCli, err := rpc.Dial(l2chainURL)
	if err != nil {
		return nil, err
	}

	var (
		client = ethclient.NewClient(rpcCli)
		cts    = &l2Contracts{
			rpcCli:        rpcCli,
			client:        client,
			cfg:           cfg,
			chainName:     "l2_chain",
			withdraw:      msgproof.NewWithdrawTrie(),
			txHashMsgHash: map[string]common.Hash{},
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

	cts.filter = bytecode.NewContractsFilter("l2Watcher", []bytecode.ContractAPI{
		cts.ScrollMessenger,
		// cts.MessageQueue,

		cts.ETHGateway,
		cts.WETHGateway,
		cts.DAIGateway,
		cts.StandardERC20Gateway,
		cts.CustomERC20Gateway,
		cts.ERC721Gateway,
		cts.ERC1155Gateway,
	}...)

	// Init withdraw root.
	if err = cts.initWithdraw(db); err != nil {
		return nil, err
	}

	cts.registerGatewayHandlers()
	cts.registerMessengerHandlers()

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
	l2.txHashMsgHash = map[string]common.Hash{}
	l2.msgSentEvents = map[uint64][]*orm.L2MessengerEvent{}
	l2.ethEvents = l2.ethEvents[:0]
	l2.erc20Events = l2.erc20Events[:0]
	l2.erc721Events = l2.erc721Events[:0]
	l2.erc1155Events = l2.erc1155Events[:0]
}

func (l2 *l2Contracts) ParseL2Events(ctx context.Context, db *gorm.DB, start, end uint64) (int, error) {
	l2.clean()
	l2.tx = db.Begin().WithContext(ctx)
	count, err := l2.filter.ParseLogs(ctx, l2.client, start, end)
	if err != nil {
		controller.ParseLogsFailureTotal.WithLabelValues(l2.chainName).Inc()
		l2.tx.Rollback()
		return 0, err
	}

	// store l2Messenger sentMessenger events.
	if err = l2.storeMessengerEvents(ctx, start, end); err != nil {
		l2.tx.Rollback()
		return 0, err
	}

	// store l2chain gateway events.
	if err = l2.storeGatewayEvents(); err != nil {
		l2.tx.Rollback()
		return 0, err
	}

	err = l2.tx.Save(&orm.L2Block{
		Number:     end,
		EventCount: count,
	}).Error
	if err != nil {
		l2.tx.Rollback()
		return 0, err
	}

	if err = l2.tx.Commit().Error; err != nil {
		l2.tx.Rollback()
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

func (l2 *l2Contracts) setMonitorAPI(monitor controller.MonitorAPI) {
	l2.monitorAPI = monitor
}
