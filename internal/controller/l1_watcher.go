package controller

import (
	"context"
	"time"

	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"
	"modernc.org/mathutil"

	"chain-monitor/internal/config"
	"chain-monitor/internal/logic"
	"chain-monitor/orm"
)

var (
	BatchSize uint64 = 100
)

// L1Watcher return a new instance of L1Watcher.
type L1Watcher struct {
	cfg    *config.L1Config
	client *ethclient.Client

	contracts *logic.L1Contracts

	latestNumber uint64
	safeNumber   uint64
	startNumber  uint64

	db *gorm.DB
}

func NewL1Watcher(cfg *config.L1Config, db *gorm.DB) (*L1Watcher, error) {
	client, err := ethclient.Dial(cfg.L1ChainURL)
	contracts, err := logic.NewL1Contracts(client, db, cfg.L1Gateways)
	if err != nil {
		return nil, err
	}

	l1Block, err := orm.GetL1LatestBlock(db)
	if err != nil {
		return nil, err
	}

	watcherClient := &L1Watcher{
		cfg:         cfg,
		db:          db,
		client:      client,
		contracts:   contracts,
		startNumber: mathutil.MaxUint64(l1Block.Number, cfg.StartNumber),
	}

	return watcherClient, nil
}

func (l1 *L1Watcher) ScanL1Chain(ctx context.Context) {
	start, end, err := l1.getStartAndEndNumber(ctx)
	if err != nil {
		log.Error("failed to get l1Chain start and end number", "err", err)
		return
	}
	if start > end {
		return
	}

	cts := l1.contracts
	err = cts.ParseL1Events(ctx, start, end)
	if err != nil {
		log.Error("failed to parse l1chain events", "start", start, "end", end, "err", err)
		return
	}

	l1.startNumber = end
	log.Info("scan l1chain successful", "start", start, "end", end)
	return
}

func (l1 *L1Watcher) getStartAndEndNumber(ctx context.Context) (uint64, uint64, error) {
	var (
		start = l1.startNumber + 1
		end   = start + BatchSize - 1
	)
	if end > l1.safeNumber {
		// Get safeNumber block latestHeight.
		number, err := l1.client.BlockNumber(ctx)
		if err != nil {
			log.Error("failed to get safeNumber block latestHeight", "err", err)
			return 0, 0, err
		}
		l1.latestNumber = number
		l1.safeNumber = mathutil.MaxUint64(l1.safeNumber, number-l1.cfg.Confirm)
		time.Sleep(time.Second * 2)
	}
	return start, mathutil.MinUint64(end, l1.safeNumber), nil
}
