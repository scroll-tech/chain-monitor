package controller

import (
	"context"
	"math/big"
	"sync/atomic"
	"time"

	"chain-monitor/internal/config"
	"chain-monitor/internal/logic"
	"chain-monitor/orm"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"
)

// L2Watcher return a new instance of L2Watcher.
type L2Watcher struct {
	cfg    *config.L2Config
	client *ethclient.Client

	contracts *logic.L2Contracts

	curTime     time.Time
	startNumber uint64
	safeNumber  uint64

	db *gorm.DB
}

func NewL2Watcher(cfg *config.L2Config, db *gorm.DB) (*L2Watcher, error) {
	client, err := ethclient.Dial(cfg.L2ChainURL)
	if err != nil {
		return nil, err
	}

	l2Block, err := orm.GetL2LatestBlock(db)
	if err != nil {
		return nil, err
	}
	latestNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	// Create a event filter instance.
	l1Contracts, err := logic.NewL2Contracts(client, cfg.L2gateways)
	if err != nil {
		return nil, err
	}

	watcher := &L2Watcher{
		cfg:         cfg,
		db:          db,
		client:      client,
		contracts:   l1Contracts,
		curTime:     time.Now(),
		startNumber: l2Block.Number,
		safeNumber:  latestNumber - cfg.Confirm,
	}

	return watcher, nil
}

func (l2 *L2Watcher) StartNumber() uint64 {
	return atomic.LoadUint64(&l2.startNumber)
}

func (l2 *L2Watcher) WithdrawRoot(ctx context.Context, number uint64) (common.Hash, error) {
	data, err := l2.client.StorageAt(
		ctx,
		l2.cfg.L2gateways.MessageQueue,
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
	if end > l2.safeNumber {
		return
	}

	cts := l2.contracts
	err = cts.ParseL2Events(ctx, l2.db, start, end)
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

	if end <= l2.safeNumber {
		return start, end, nil
	}
	if start < l2.safeNumber {
		return start, l2.safeNumber - 1, nil
	}

	curTime := time.Now()
	if int(curTime.Sub(l2.curTime).Seconds()) >= 5 {
		latestNumber, err := l2.client.BlockNumber(ctx)
		if err != nil {
			return 0, 0, err
		}
		l2.safeNumber = latestNumber - l2.cfg.Confirm
		l2.curTime = curTime
	}
	return start, start, nil
}
