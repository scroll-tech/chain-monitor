package l1watcher

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
	"chain-monitor/internal/orm"
)

var l1BatchSize uint64 = 100

// L1Watcher return a new instance of L1Watcher.
type L1Watcher struct {
	cfg    *config.L1Config
	client *ethclient.Client

	filter *l1Contracts

	cacheLen    int
	headerCache []*types.Header

	curTime     time.Time
	startNumber uint64
	safeNumber  uint64

	db *gorm.DB
}

func NewL1Watcher(cfg *config.L1Config, db *gorm.DB) (*L1Watcher, error) {
	client, err := ethclient.Dial(cfg.L1ChainURL)
	if err != nil {
		return nil, err
	}

	contracts, err := newL1Contracts(client, cfg.L1Gateways)
	if err != nil {
		return nil, err
	}

	l1Block, err := orm.GetL1LatestBlock(db)
	if err != nil {
		return nil, err
	}
	latestNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	watcherClient := &L1Watcher{
		cfg:         cfg,
		db:          db,
		client:      client,
		filter:      contracts,
		cacheLen:    32,
		headerCache: make([]*types.Header, 0, 32),
		curTime:     time.Now(),
		startNumber: mathutil.MaxUint64(l1Block.Number, cfg.StartNumber),
		safeNumber:  latestNumber - cfg.Confirm,
	}

	return watcherClient, nil
}

func (l1 *L1Watcher) ScanL1Chain(ctx context.Context) {
	start, end, err := l1.getStartAndEndNumber(ctx)
	if err != nil {
		log.Error("failed to get l1Chain start and end number", "err", err)
		return
	}
	if end > l1.SafeNumber() {
		return
	}

	var count int
	// If we sync events one by one.
	if start == end {
		var header *types.Header
		header, err = l1.checkReorg(ctx)
		if err != nil {
			log.Error("appear error when do l1chain reorg process", "number", start, "err", err)
			return
		}
		// get events by number
		start = header.Number.Uint64()
		end = start
		count, err = l1.filter.ParseL1Events(ctx, l1.db, start, end)
		if err != nil {
			log.Error("failed to parse l1chain events", "start", start, "end", end, "err", err)
			return
		}
		// append block header
		if len(l1.headerCache) >= l1.cacheLen {
			l1.headerCache = l1.headerCache[:l1.cacheLen-1]
		}
		l1.headerCache = append(l1.headerCache, header)
	} else {
		count, err = l1.filter.ParseL1Events(ctx, l1.db, start, end)
		if err != nil {
			log.Error("failed to parse l1chain events", "start", start, "end", end, "err", err)
			return
		}
	}
	l1.setStartNumber(end)

	log.Info("scan l1chain successful", "start", start, "end", end, "event_count", count)
}

func (l1 *L1Watcher) getStartAndEndNumber(ctx context.Context) (uint64, uint64, error) {
	var (
		start = l1.StartNumber() + 1
		end   = start + l1BatchSize - 1
	)
	safeNumber := l1.SafeNumber() - uint64(l1.cacheLen/2)
	if end <= safeNumber {
		return start, end, nil
	}
	if start < safeNumber {
		return start, safeNumber - 1, nil
	}

	// update latest number
	curTime := time.Now()
	if int(curTime.Sub(l1.curTime).Seconds()) >= 5 {
		latestNumber, err := l1.client.BlockNumber(ctx)
		if err != nil {
			return 0, 0, err
		}
		l1.setSafeNumber(latestNumber - l1.cfg.Confirm)
		l1.curTime = curTime
	}

	// Scan l2chain number one by one.
	return start, start, nil
}

func (l1 *L1Watcher) checkReorg(ctx context.Context) (*types.Header, error) {
	var number uint64
	if len(l1.headerCache) == 0 {
		number = l1.StartNumber()
	} else {
		number = l1.headerCache[len(l1.headerCache)-1].Number.Uint64()
	}
	number++

	// get the next block header.
	header, err := l1.client.HeaderByNumber(ctx, big.NewInt(0).SetUint64(number))
	if err != nil {
		return nil, err
	}

	if len(l1.headerCache) == 0 {
		return header, nil
	}

	var reorgNumbers []uint64
	for len(l1.headerCache) > 0 {
		latestHeader := l1.headerCache[len(l1.headerCache)-1]
		if header.ParentHash == latestHeader.Hash() {
			break
		}
		if header.ParentHash != latestHeader.Hash() {
			reorgNumbers = append(reorgNumbers, latestHeader.Number.Uint64())
			l1.headerCache = l1.headerCache[:len(l1.headerCache)-1]
			var (
				parentHash = header.ParentHash
			)
			header, err = l1.client.HeaderByHash(ctx, parentHash)
			if err != nil {
				return nil, err
			}
		}
	}

	// TODO: A deeper rollback is required
	if len(l1.headerCache) == 0 {
		panic(fmt.Errorf("l1chain reorged too deep"))
	}

	// Reorg stored events if the reorg headers is not empty.
	if err = deleteReorgEvents(ctx, l1.db, reorgNumbers); err != nil {
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
		tables = orm.L1Tables
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
