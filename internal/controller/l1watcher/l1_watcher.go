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
	"chain-monitor/internal/controller"
	"chain-monitor/internal/orm"
	"chain-monitor/internal/utils"
)

var l1BatchSize uint64 = 50

// L1Watcher return a new instance of L1Watcher.
type L1Watcher struct {
	cfg    *config.L1Config
	client *ethclient.Client

	filter *l1Contracts

	isOneByOne  bool
	cacheLen    int
	headerCache []*types.Header

	curTime    time.Time
	currNumber uint64
	safeNumber uint64

	db *gorm.DB
}

// NewL1Watcher create a l1watcher instance.
func NewL1Watcher(cfg *config.L1Config, db *gorm.DB) (*L1Watcher, error) {
	client, err := ethclient.Dial(cfg.L1URL)
	if err != nil {
		return nil, err
	}

	l2Filter, err := newL1Contracts(cfg.L1URL, cfg.L1Contracts)
	if err != nil {
		return nil, err
	}

	l1Block, err := orm.GetL1LatestBlock(db)
	if err != nil {
		return nil, err
	}

	// Get confirm number.
	number, err := utils.GetLatestConfirmedBlockNumber(context.Background(), client, cfg.Confirm)
	if err != nil {
		return nil, err
	}

	watcher := &L1Watcher{
		cfg:         cfg,
		db:          db,
		client:      client,
		filter:      l2Filter,
		cacheLen:    32,
		headerCache: make([]*types.Header, 0, 32),
		curTime:     time.Now(),
		currNumber:  mathutil.MaxUint64(l1Block.Number, cfg.StartNumber),
		safeNumber:  number,
	}

	return watcher, nil
}

// ScanL1Chain scan l1chain entrypoint function.
func (l1 *L1Watcher) ScanL1Chain(ctx context.Context) {
	start, end, err := l1.getStartAndEndNumber(ctx)
	if err != nil {
		log.Error("failed to get l1Chain start and end number", "err", err)
		return
	}
	safeNumber := l1.SafeNumber()
	if end > safeNumber {
		return
	}
	l1.filter.checkBalance = (safeNumber - start) <= 10

	var count int
	// If we sync events one by one.
	if l1.isOneByOne || start == end {
		l1.isOneByOne = true
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
			l1.headerCache = l1.headerCache[1:]
		}
		l1.headerCache = append(l1.headerCache, header)
	} else {
		count, err = l1.filter.ParseL1Events(ctx, l1.db, start, end)
		if err != nil {
			log.Error("failed to parse l1chain events", "start", start, "end", end, "err", err)
			return
		}
	}
	l1.setCurrentNumber(end)

	// Metrics records current goroutine.
	controller.WorkerStartedTotal.WithLabelValues("l1_watcher").Inc()

	log.Info("scan l1chain successful", "start", start, "end", end, "check_balance", l1.filter.checkBalance, "event_count", count)
}

func (l1 *L1Watcher) getStartAndEndNumber(ctx context.Context) (uint64, uint64, error) {
	var (
		start = l1.CurrentNumber() + 1
		end   = mathutil.MinUint64(start+l1BatchSize-1, l1.SafeNumber())
	)

	// update latest number
	curTime := time.Now()
	if int(curTime.Sub(l1.curTime).Seconds()) >= 5 {
		number, err := utils.GetLatestConfirmedBlockNumber(ctx, l1.client, l1.cfg.Confirm)
		if err != nil {
			return 0, 0, err
		}
		number = mathutil.MaxUint64(number, l1.SafeNumber())
		l1.setSafeNumber(number)
		l1.curTime = curTime
		controller.BlockNumber.WithLabelValues(l1.filter.chainName).Set(float64(number))
	}

	if start <= end {
		return start, end, nil
	}

	// Scan l2chain number one by one.
	return start, start, nil
}

func (l1 *L1Watcher) checkReorg(ctx context.Context) (*types.Header, error) {
	var number uint64
	if len(l1.headerCache) == 0 {
		number = l1.CurrentNumber()
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
		// reorg appeared.
		reorgNumbers = append(reorgNumbers, latestHeader.Number.Uint64())
		l1.headerCache = l1.headerCache[:len(l1.headerCache)-1]
		header, err = l1.client.HeaderByNumber(ctx, latestHeader.Number)
		if err != nil {
			return nil, err
		}
	}

	// TODO: A deeper rollback is required
	if len(l1.headerCache) == 0 {
		panic(fmt.Errorf("l1chain reorged too deep"))
	}

	// Reorg stored events if the reorg headers is not empty.
	if err = deleteReorgEvents(ctx, l1.db, reorgNumbers, l1.filter.chainName); err != nil {
		return nil, err
	}

	return header, nil
}

func deleteReorgEvents(ctx context.Context, db *gorm.DB, numbers []uint64, chainName string) error {
	if len(numbers) == 0 {
		return nil
	}

	log.Warn("L1 chain reorg", "reorg numbers", numbers)
	controller.ReorgTotal.WithLabelValues(chainName).Inc()

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
