package utils

import (
	"context"
	"fmt"
	"math/big"

	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/rpc"
	"golang.org/x/sync/errgroup"
	"modernc.org/mathutil"

	"github.com/scroll-tech/chain-monitor/internal/logic/contracts/abi/il2scrollmessenger"
)

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

// GetL2WithdrawRootsInRange gets batch withdraw roots within a block range (inclusive) from the geth node.
func GetL2WithdrawRootsInRange(ctx context.Context, cli *rpc.Client, queueAddr common.Address, startBlockNumber, endBlockNumber uint64) (map[uint64]common.Hash, error) {
	numbers := endBlockNumber - startBlockNumber + 1
	withdrawRoots := make([][]byte, numbers)
	reqs := make([]rpc.BatchElem, numbers)
	for i := startBlockNumber; i <= endBlockNumber; i++ {
		n := big.NewInt(0).SetUint64(i + startBlockNumber)
		reqs[i-startBlockNumber] = rpc.BatchElem{
			Method: "eth_getStorageAt",
			Args:   []interface{}{queueAddr, common.Hash{}, n},
			Result: &withdrawRoots[i-startBlockNumber],
		}
	}
	parallels := 8
	eg := errgroup.Group{}
	eg.SetLimit(parallels)
	for i := 0; i < int(numbers); i += parallels {
		start := i
		eg.Go(func() error {
			return cli.BatchCallContext(ctx, reqs[start:mathutil.Min(start+parallels, len(reqs))])
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	withdrawRootsMap := make(map[uint64]common.Hash)
	for i, withdrawRoot := range withdrawRoots {
		withdrawRootsMap[startBlockNumber+uint64(i)] = common.BytesToHash(withdrawRoot)
	}
	return withdrawRootsMap, nil
}

// UnpackLog unpacks a retrieved log into the provided output structure.
func UnpackLog(c *abi.ABI, out interface{}, event string, log types.Log) error {
	if log.Topics[0] != c.Events[event].ID {
		return fmt.Errorf("event signature mismatch")
	}
	if len(log.Data) > 0 {
		if err := c.UnpackIntoInterface(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range c.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopics(out, indexed, log.Topics[1:])
}
