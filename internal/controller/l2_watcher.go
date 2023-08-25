package controller

import (
	"chain-monitor/internal/config"
	"chain-monitor/internal/logic"
	"chain-monitor/orm"
	"context"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"
	"math/big"
	"modernc.org/mathutil"
	"sync/atomic"
	"time"
)

// L2Watcher return a new instance of L2Watcher.
type L2Watcher struct {
	cfg    *config.L2Config
	client *ethclient.Client

	contracts *logic.L2Contracts

	latestNumber uint64
	safeNumber   uint64
	startNumber  uint64

	db *gorm.DB
}

func NewL2Watcher(cfg *config.L2Config, db *gorm.DB) (*L2Watcher, error) {
	client, err := ethclient.Dial(cfg.L2ChainURL)
	if err != nil {
		return nil, err
	}
	l1Contracts, err := logic.NewL2Contracts(client, db, cfg.L2gateways)
	if err != nil {
		return nil, err
	}

	l2Block, err := orm.GetL2LatestBlock(db)
	if err != nil {
		return nil, err
	}

	watcher := &L2Watcher{
		cfg:         cfg,
		db:          db,
		client:      client,
		contracts:   l1Contracts,
		startNumber: l2Block.Number,
	}

	return watcher, nil
}

func (l2 *L2Watcher) StartNumber() uint64 {
	return atomic.LoadUint64(&l2.startNumber)
}

func (l2 *L2Watcher) WithdrawRoot(ctx context.Context, number uint64) (common.Hash, error) {
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

func (l2 *L2Watcher) ScanL2Chain(ctx context.Context) {
	start, end, err := l2.getStartAndEndNumber(ctx)
	if err != nil {
		log.Error("l2Watcher failed to get start and end number", "err", err)
		return
	}

	cts := l2.contracts
	err = cts.ParseL2Events(ctx, start, end)
	if err != nil {
		log.Error("failed to parse l2chain events", "start", start, "end", end, "err", err)
		return
	}

	atomic.StoreUint64(&l2.startNumber, end)
	log.Info("scan l2chain successful", "start", start, "end", end)
	return
}

func (l2 *L2Watcher) getStartAndEndNumber(ctx context.Context) (uint64, uint64, error) {
	var (
		start = atomic.LoadUint64(&l2.startNumber) + 1
		end   = start + BatchSize - 1
	)
	if end > l2.safeNumber {
		// Get safeNumber block latestHeight.
		number, err := l2.client.BlockNumber(ctx)
		if err != nil {
			log.Error("failed to get safeNumber block latestHeight", "err", err)
			return 0, 0, err
		}
		l2.latestNumber = number
		l2.safeNumber = mathutil.MaxUint64(l2.safeNumber, number-l2.cfg.Confirm)
		time.Sleep(time.Second * 2)
	}
	return start, mathutil.MinUint64(end, l2.safeNumber), nil
}
