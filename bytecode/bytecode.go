package bytecode

import (
	"chain-monitor/internal/utils"
	"context"
	"encoding/binary"
	"errors"
	"github.com/scroll-tech/go-ethereum/log"
	"math/big"

	geth "github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/ethclient"
)

type EventHandler func(vLog *types.Log, value interface{}) error

// ContractAPI it's common for contract.
type ContractAPI interface {
	GetName() string
	GetAddress() common.Address
	GetEventName(sigHash common.Hash) string
	GetSigHashes() []common.Hash
	ParseLog(vLog *types.Log) (bool, error)
}

type Contract struct {
	Name     string
	Address  common.Address
	ABI      *abi.ABI
	parsers  map[common.Hash]func(log *types.Log) (interface{}, error)
	handlers map[common.Hash]EventHandler
}

// ContractsFilter contracts filter struct.
type ContractsFilter struct {
	name         string
	query        *geth.FilterQuery
	contractAPIs map[common.Address]ContractAPI
}

// NewContractsFilter return a contracts filter instance.
func NewContractsFilter(name string, cAPIs ...ContractAPI) *ContractsFilter {
	contractAPIs := make(map[common.Address]ContractAPI)
	for _, cABI := range cAPIs {
		addr := cABI.GetAddress()
		if !utils.IsNil(contractAPIs[addr]) {
			log.Warn("conflict contract name", "name", cABI.GetName())
			continue
		}
		contractAPIs[addr] = cABI
	}
	addrList := make([]common.Address, 0, len(contractAPIs))
	for addr := range contractAPIs {
		addrList = append(addrList, addr)
	}
	return &ContractsFilter{
		name: name,
		query: &geth.FilterQuery{
			FromBlock: big.NewInt(0),
			ToBlock:   big.NewInt(0),
			Addresses: addrList,
		},
		contractAPIs: contractAPIs,
	}
}

// ParseLogs parse logs.
func (c *ContractsFilter) ParseLogs(ctx context.Context, client *ethclient.Client, start, end uint64) (int, error) {
	c.query.FromBlock.SetUint64(start)
	c.query.ToBlock.SetUint64(end)
	logs, err := client.FilterLogs(ctx, *c.query)
	if err != nil {
		return 0, err
	}
	for _, vLog := range logs {
		cAPI := c.contractAPIs[vLog.Address]
		exist, err := cAPI.ParseLog(&vLog)
		if err != nil {
			return 0, err
		}
		if !exist {
			log.Debug("unsupported event", "tx_hash", vLog.TxHash.String(), "event_name", cAPI.GetEventName(vLog.Topics[0]))
		}
	}
	return len(logs), nil
}

type commitBatchArgs struct {
	Version                uint8
	ParentBatchHeader      []byte
	Chunks                 [][]byte
	SkippedL1MessageBitmap []byte
}

type ScrollBatch struct {
	BatchIndex    uint64
	L2StartNumber uint64
	L2EndNumber   uint64
}

// GetBatchRangeFromCalldataV2 find the block range from calldata, both inclusive.
func GetBatchRangeFromCalldataV2(ABI *abi.ABI, data []byte) (*ScrollBatch, error) {
	method := ABI.Methods["commitBatch"]
	values, err := method.Inputs.Unpack(data[4:])
	if err != nil {
		// special case: import genesis batch
		method = ABI.Methods["importGenesisBatch"]
		_, err2 := method.Inputs.Unpack(data[4:])
		if err2 == nil {
			// genesis batch
			return &ScrollBatch{}, nil
		}
		// none of "commitBatch" and "importGenesisBatch" match, give up
		return nil, err
	}
	args := commitBatchArgs{}
	err = method.Inputs.Copy(&args, values)
	if err != nil {
		return nil, err
	}

	var startBlock uint64
	var finishBlock uint64

	// decode batchIndex from ParentBatchHeader
	if len(args.ParentBatchHeader) < 9 {
		return nil, errors.New("invalid parent batch header")
	}
	batchIndex := binary.BigEndian.Uint64(args.ParentBatchHeader[1:9]) + 1

	// decode blocks from chunk and assume that there's no empty chunk
	// |   1 byte   | 60 bytes | ... | 60 bytes |
	// | num blocks |  block 1 | ... |  block n |
	if len(args.Chunks) == 0 {
		return nil, errors.New("invalid chunks")
	}
	chunk := args.Chunks[0]
	block := chunk[1:61] // first block in chunk
	startBlock = binary.BigEndian.Uint64(block[0:8])

	chunk = args.Chunks[len(args.Chunks)-1]
	lastBlockIndex := int(chunk[0]) - 1
	block = chunk[1+lastBlockIndex*60 : 1+lastBlockIndex*60+60] // last block in chunk
	finishBlock = binary.BigEndian.Uint64(block[0:8])

	return &ScrollBatch{
		BatchIndex:    batchIndex,
		L2StartNumber: startBlock,
		L2EndNumber:   finishBlock,
	}, err
}
