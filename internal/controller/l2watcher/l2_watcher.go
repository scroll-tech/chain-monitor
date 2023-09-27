package l2watcher

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/log"
	"gorm.io/gorm"
	"modernc.org/mathutil"

	"chain-monitor/internal/config"
	"chain-monitor/internal/controller"
	"chain-monitor/internal/orm"
	"chain-monitor/internal/utils"
)

var l2BatchSize uint64 = 200

// L2Watcher is a watcher for the Layer 2 chain. It observes and tracks events on the L2 chain.
type L2Watcher struct {
	cfg    *config.L2Config
	client *ethclient.Client

	filter *l2Contracts

	isOneByOne  bool
	cacheLen    int
	headerCache []*types.Header

	curTime    time.Time
	currNumber uint64
	safeNumber uint64

	db *gorm.DB
}

// NewL2Watcher initializes a new L2Watcher with the given L2 configuration and database connection.
func NewL2Watcher(cfg *config.L2Config, db *gorm.DB) (*L2Watcher, error) {
	client, err := ethclient.Dial(cfg.L2URL)
	if err != nil {
		return nil, err
	}

	l2Block, err := orm.GetL2LatestBlock(db)
	if err != nil {
		return nil, err
	}

	// Get confirm number.
	number, err := utils.GetLatestConfirmedBlockNumber(context.Background(), client, cfg.Confirm)
	if err != nil {
		return nil, err
	}

	// Create a event gatewayFilter instance.
	l2Filter, err := newL2Contracts(cfg.L2URL, db, cfg.L2Contracts)
	if err != nil {
		return nil, err
	}

	watcher := &L2Watcher{
		cfg:         cfg,
		db:          db,
		client:      client,
		filter:      l2Filter,
		cacheLen:    32,
		headerCache: make([]*types.Header, 0, 32),
		curTime:     time.Now(),
		currNumber:  l2Block.Number,
		safeNumber:  number,
	}

	return watcher, nil
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

	var count int
	if l2.isOneByOne || start == end {
		l2.isOneByOne = true
		var header *types.Header
		header, err = l2.checkReorg(ctx)
		if err != nil {
			log.Error("appear error when do l2chain reorg process", "number", start, "err", err)
			return
		}
		start = header.Number.Uint64()
		end = start
		count, err = l2.filter.ParseL2Events(ctx, l2.db, start, end)
		if err != nil {
			log.Error("failed to parse l2chain events", "start", start, "end", end, "err", err)
			return
		}
		// append block header
		if len(l2.headerCache) >= l2.cacheLen {
			l2.headerCache = l2.headerCache[1:]
		}
		l2.headerCache = append(l2.headerCache, header)
	} else {
		count, err = l2.filter.ParseL2Events(ctx, l2.db, start, end)
		if err != nil {
			log.Error("failed to parse l2chain events", "start", start, "end", end, "err", err)
			return
		}
	}
	l2.setCurrentNumber(end)

	// Metrics records current goroutine.
	controller.WorkerStartedTotal.WithLabelValues("l2_watcher").Inc()

	log.Info("scan l2chain successful", "start", start, "end", end, "event_count", count)
}

func (l2 *L2Watcher) getStartAndEndNumber(ctx context.Context) (uint64, uint64, error) {
	var (
		start = l2.CurrentNumber() + 1
		end   = mathutil.MinUint64(start+l2BatchSize-1, l2.SafeNumber())
	)
	if start <= end {
		return start, end, nil
	}

	curTime := time.Now()
	if int(curTime.Sub(l2.curTime).Seconds()) >= 5 {
		number, err := utils.GetLatestConfirmedBlockNumber(ctx, l2.client, l2.cfg.Confirm)
		if err != nil {
			return 0, 0, err
		}
		number = mathutil.MaxUint64(number, l2.SafeNumber())
		l2.setSafeNumber(number)
		l2.curTime = curTime
		controller.BlockNumber.WithLabelValues(l2.filter.chainName).Set(float64(number))
	}
	return start, start, nil
}

func (l2 *L2Watcher) checkReorg(ctx context.Context) (*types.Header, error) {
	var number uint64
	if len(l2.headerCache) == 0 {
		number = l2.CurrentNumber()
	} else {
		number = l2.headerCache[len(l2.headerCache)-1].Number.Uint64()
	}
	number++

	header, err := l2.client.HeaderByNumber(ctx, big.NewInt(0).SetUint64(number))
	if err != nil {
		return nil, err
	}
	if len(l2.headerCache) == 0 {
		return header, nil
	}

	var reorgNumbers []uint64
	for len(l2.headerCache) > 0 {
		latestHeader := l2.headerCache[len(l2.headerCache)-1]
		if header.ParentHash == latestHeader.Hash() {
			break
		}
		// reorg appeared.
		reorgNumbers = append(reorgNumbers, latestHeader.Number.Uint64())
		l2.headerCache = l2.headerCache[:len(l2.headerCache)-1]
		header, err = l2.client.HeaderByNumber(ctx, latestHeader.Number)
		if err != nil {
			return nil, err
		}
	}

	// TODO: A deeper rollback is required
	if len(l2.headerCache) == 0 {
		panic(fmt.Errorf("l2chain reorged too deep"))
	}
	// Record reorg times.
	if len(reorgNumbers) > 0 {
		log.Warn("L2 chain reorg", "reorg numbers", reorgNumbers)
		controller.ReorgTotal.WithLabelValues(l2.filter.chainName).Inc()
	}

	// Reorg stored events if the reorg headers is not empty.
	if err = deleteReorgEvents(ctx, l2.db, reorgNumbers); err != nil {
		return nil, err
	}

	return header, nil
}

func deleteReorgEvents(ctx context.Context, db *gorm.DB, numbers []uint64) error {
	if len(numbers) == 0 {
		return nil
	}

	var (
		start  = numbers[0]
		end    = numbers[len(numbers)-1]
		tables = orm.L2Tables
		result *gorm.DB
	)
	tx := db.Begin().WithContext(ctx)
	for _, tb := range tables {
		// delete eth events.
		result = tx.Where("number BETWEEN ? AND ?", start, end).Delete(tb)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
