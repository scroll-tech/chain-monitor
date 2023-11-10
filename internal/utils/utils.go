package utils

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/modern-go/reflect2"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/rpc"
	"golang.org/x/sync/errgroup"
	"modernc.org/mathutil"

	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il2scrollmessenger"
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
	for {
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

// ComputeMessageHash compute message event fields to get message hash.
func ComputeMessageHash(
	sender common.Address,
	target common.Address,
	value *big.Int,
	messageNonce *big.Int,
	message []byte,
) common.Hash {
	parsed, _ := il2scrollmessenger.Il2scrollmessengerMetaData.GetAbi()
	data, _ := parsed.Pack("relayMessage", sender, target, value, messageNonce, message)
	return common.BytesToHash(crypto.Keccak256(data))
}

// GetLatestConfirmedBlockNumber get confirmed block number by rpc.BlockNumber type.
func GetLatestConfirmedBlockNumber(ctx context.Context, client *ethclient.Client, confirm rpc.BlockNumber) (uint64, error) {
	switch true {
	case confirm == rpc.SafeBlockNumber || confirm == rpc.FinalizedBlockNumber:
		var tag *big.Int
		if confirm == rpc.FinalizedBlockNumber {
			tag = big.NewInt(int64(rpc.FinalizedBlockNumber))
		} else {
			tag = big.NewInt(int64(rpc.SafeBlockNumber))
		}

		header, err := client.HeaderByNumber(ctx, tag)
		if err != nil {
			return 0, err
		}
		if !header.Number.IsInt64() {
			return 0, fmt.Errorf("received invalid block confirm: %v", header.Number)
		}
		return header.Number.Uint64(), nil
	case confirm == rpc.LatestBlockNumber:
		number, err := client.BlockNumber(ctx)
		if err != nil {
			return 0, err
		}
		return number, nil
	case confirm.Int64() >= 0: // If it's positive integer, consider it as a certain confirm value.
		number, err := client.BlockNumber(ctx)
		if err != nil {
			return 0, err
		}
		cfmNum := uint64(confirm.Int64())

		if number >= cfmNum {
			return number - cfmNum, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("unknown confirmation type: %v", confirm)
	}
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

var emptyHash = common.BigToHash(big.NewInt(0))

// GetL2WithdrawRootsInRange gets batch withdraw roots within a block range (inclusive) from the geth node.
func GetL2WithdrawRootsInRange(ctx context.Context, cli *rpc.Client, queueAddr common.Address, startBlockNumber, endBlockNumber uint64) (map[uint64]common.Hash, error) {
	blockNumbers := make([]uint64, endBlockNumber-startBlockNumber+1)
	for i := startBlockNumber; i <= endBlockNumber; i++ {
		blockNumbers[i-startBlockNumber] = i
	}
	reqs := make([]rpc.BatchElem, len(blockNumbers))
	withdrawRootsMap := make(map[uint64]common.Hash)
	for i, number := range blockNumbers {
		nb := big.NewInt(0).SetUint64(number)
		reqs[i] = rpc.BatchElem{
			Method: "eth_getStorageAt",
			Args:   []interface{}{queueAddr, emptyHash, toBlockNumArg(nb)},
			Result: withdrawRootsMap[number],
		}
	}
	parallels := 8
	eg := errgroup.Group{}
	eg.SetLimit(parallels)
	for i := 0; i < len(blockNumbers); i += parallels {
		start := i
		eg.Go(func() error {
			return cli.BatchCallContext(ctx, reqs[start:mathutil.Min(start+parallels, len(reqs))])
		})
	}
	return withdrawRootsMap, eg.Wait()
}

// GetBatchBalances get batch account balances at given block numbers.
func GetBatchBalances(ctx context.Context, cli *rpc.Client, addr common.Address, numbers []uint64) ([]*big.Int, error) {
	stringResults := make([]string, len(numbers))
	reqs := make([]rpc.BatchElem, len(numbers))
	for i, number := range numbers {
		nb := big.NewInt(0).SetUint64(number)
		reqs[i] = rpc.BatchElem{
			Method: "eth_getBalance",
			Args:   []interface{}{addr, toBlockNumArg(nb)},
			Result: &stringResults[i],
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

	err := eg.Wait()
	if err != nil {
		return nil, err
	}

	balances := make([]*big.Int, len(numbers))
	for i, str := range stringResults {
		data, err := hexutil.Decode(str)
		if err != nil {
			return nil, err
		}
		balances[i] = big.NewInt(0).SetBytes(data)
	}
	return balances, nil
}
