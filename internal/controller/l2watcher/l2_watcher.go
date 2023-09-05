package l2watcher

import (
	"context"
	"math/big"
	"time"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"

	"chain-monitor/internal/config"
	"chain-monitor/internal/orm"
)

var l2BatchSize uint64 = 500

// L2Watcher is a watcher for the Layer 2 chain. It observes and tracks events on the L2 chain.
type L2Watcher struct {
	cfg    *config.L2Config
	client *ethclient.Client

	filter *l2Contracts

	curTime     time.Time
	startNumber uint64
	safeNumber  uint64

	db *gorm.DB
}

// NewL2Watcher initializes a new L2Watcher with the given L2 configuration and database connection.
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
	l1Contracts, err := newL2Contracts(cfg.L2ChainURL, db, cfg.L2gateways)
	if err != nil {
		return nil, err
	}

	watcher := &L2Watcher{
		cfg:         cfg,
		db:          db,
		client:      client,
		filter:      l1Contracts,
		curTime:     time.Now(),
		startNumber: l2Block.Number,
		safeNumber:  latestNumber - cfg.Confirm,
	}

	return watcher, nil
}

// WithdrawRoot fetches the root hash of withdrawal data from the storage of the MessageQueue contract.
func (l2 *L2Watcher) WithdrawRoot(ctx context.Context, number uint64) (common.Hash, error) {
	data, err := l2.client.StorageAt(
		ctx,
		l2.cfg.L2gateways.MessageQueue,
		common.BigToHash(big.NewInt(0)), big.NewInt(0).SetUint64(number),
	)
	if err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(data), nil
}

// ScanL2Chain scans a range of blocks on the L2 chain for events.
func (l2 *L2Watcher) ScanL2Chain(ctx context.Context) {
	start, end, err := l2.getStartAndEndNumber(ctx)
	if err != nil {
		log.Error("l2Watcher failed to get start and end number", "err", err)
		return
	}
	if end > l2.SafeNumber() {
		return
	}

	cts := l2.filter
	count, err := cts.ParseL2Events(ctx, l2.db, start, end)
	if err != nil {
		log.Error("failed to parse l2chain events", "start", start, "end", end, "err", err)
		return
	}

	l2.setStartNumber(end)
	log.Info("scan l2chain successful", "start", start, "end", end, "event_count", count)
}

func (l2 *L2Watcher) getStartAndEndNumber(ctx context.Context) (uint64, uint64, error) {
	var (
		start = l2.StartNumber() + 1
		end   = start + l2BatchSize - 1
	)

	if end <= l2.SafeNumber() {
		return start, end, nil
	}
	if start < l2.SafeNumber() {
		return start, l2.SafeNumber() - 1, nil
	}

	curTime := time.Now()
	if int(curTime.Sub(l2.curTime).Seconds()) >= 5 {
		latestNumber, err := l2.client.BlockNumber(ctx)
		if err != nil {
			return 0, 0, err
		}
		l2.setSafeNumber(latestNumber - l2.cfg.Confirm)
		l2.curTime = curTime
	}
	return start, start, nil
}
