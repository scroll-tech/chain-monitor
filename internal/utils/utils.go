package utils

import (
	"context"
	"math/big"
	"time"

	"github.com/modern-go/reflect2"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/rpc"
	"golang.org/x/sync/errgroup"
	"modernc.org/mathutil"
)

// TryTimes try run several times until the function return true.
func TryTimes(times int, run func() bool) bool {
	for i := 0; i < times; i++ {
		if run() {
			return true
		}
		time.Sleep(time.Millisecond * 500)
	}
	return false
}

// LoopWithContext Run the f func with context periodically.
func LoopWithContext(ctx context.Context, period time.Duration, f func(ctx context.Context)) {
	tick := time.NewTicker(period)
	defer tick.Stop()
	for true {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			f(ctx)
		}
	}
}

// IsNil Check if the interface is empty.
func IsNil(i interface{}) bool {
	return i == nil || reflect2.IsNil(i)
}

func ComputeMessageHash(ABI *abi.ABI,
	sender common.Address,
	target common.Address,
	value *big.Int,
	messageNonce *big.Int,
	message []byte,
) common.Hash {
	data, _ := ABI.Pack("relayMessage", sender, target, value, messageNonce, message)
	return common.BytesToHash(crypto.Keccak256(data))
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	latest := big.NewInt(int64(rpc.LatestBlockNumber))
	if number.Cmp(latest) == 0 {
		return "latest"
	}
	pending := big.NewInt(int64(rpc.PendingBlockNumber))
	if number.Cmp(pending) == 0 {
		return "pending"
	}
	finalized := big.NewInt(int64(rpc.FinalizedBlockNumber))
	if number.Cmp(finalized) == 0 {
		return "finalized"
	}
	safe := big.NewInt(int64(rpc.SafeBlockNumber))
	if number.Cmp(safe) == 0 {
		return "safe"
	}
	return hexutil.EncodeBig(number)
}

var (
	emptyHash = common.BigToHash(big.NewInt(0))
)

// GetBatchWithdrawRoots get batch withdraw roots.
func GetBatchWithdrawRoots(ctx context.Context, cli *rpc.Client, queueAddr common.Address, numbers []uint64) ([]common.Hash, error) {
	// If the numbers count is too less, just get them.
	if len(numbers) == 1 {
		var client = ethclient.NewClient(cli)
		root, err := client.StorageAt(ctx, queueAddr, emptyHash, big.NewInt(0).SetUint64(numbers[0]))
		if err != nil {
			return nil, err
		}
		return []common.Hash{common.BytesToHash(root)}, nil
	}

	withdrawRoots := make([]common.Hash, len(numbers))
	reqs := make([]rpc.BatchElem, len(numbers))
	for i, number := range numbers {
		nb := big.NewInt(0).SetUint64(number)
		reqs[i] = rpc.BatchElem{
			Method: "eth_getStorageAt",
			Args:   []interface{}{queueAddr, emptyHash, toBlockNumArg(nb)},
			Result: &withdrawRoots[i],
		}
	}
	parallels := 8
	eg := errgroup.Group{}
	eg.SetLimit(parallels)
	for i := 0; i < len(numbers); i += parallels {
		start := i
		eg.Go(func() error {
			return cli.BatchCallContext(ctx, reqs[start:mathutil.Min(start+parallels, len(reqs))])
		})
	}

	return withdrawRoots, eg.Wait()
}
