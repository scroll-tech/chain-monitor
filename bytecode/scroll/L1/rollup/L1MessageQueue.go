// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rollup

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// L1MessageQueueMetaData contains all meta data concerning the L1MessageQueue contract.
var (
	L1MessageQueueMetaData = &bind.MetaData{
		ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"skippedBitmap\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"DequeueTransaction\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"DropTransaction\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Initialized\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint64\",\"name\":\"queueIndex\",\"type\":\"uint64\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false}],\"type\":\"event\",\"name\":\"QueueTransaction\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldGateway\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_newGateway\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"UpdateEnforcedTxGateway\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldGasOracle\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_newGasOracle\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"UpdateGasOracle\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_oldMaxGasLimit\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"_newMaxGasLimit\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"UpdateMaxGasLimit\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"appendCrossDomainMessage\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"appendEnforcedTransaction\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calculateIntrinsicGasFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_queueIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"computeTransactionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}]},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"dropCrossDomainMessage\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"enforcedTxGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"estimateCrossDomainMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_queueIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"getCrossDomainMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_scrollChain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_enforcedTxGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gasOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxGasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"messageQueue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"nextCrossDomainMessageIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pendingQueueIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_skippedBitmap\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"popCrossDomainMessage\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"renounceOwnership\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"scrollChain\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferOwnership\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateEnforcedTxGateway\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGasOracle\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateGasOracle\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newMaxGasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateMaxGasLimit\"}]",
	}
	// L1MessageQueueABI is the input ABI used to generate the binding from.
	L1MessageQueueABI, _ = L1MessageQueueMetaData.GetAbi()
)

// L1MessageQueue is an auto generated Go binding around an Ethereum contract.
type L1MessageQueue struct {
	Name    string
	Address common.Address // contract address
	ABI     *abi.ABI       // L1MessageQueueABI is the input ABI used to generate the binding from.

	topics  map[common.Hash]string
	parsers map[common.Hash]func(log *types.Log) error

	L1MessageQueueCaller     // Read-only binding to the contract
	L1MessageQueueTransactor // Write-only binding to the contract
}

// GetName return L1MessageQueue's Name.
func (o *L1MessageQueue) GetName() string {
	return o.Name
}

// GetAddress return L1MessageQueue's contract address.
func (o *L1MessageQueue) GetAddress() common.Address {
	return o.Address
}

// GetEventName get event name by event hash.
func (o *L1MessageQueue) GetEventName(sigHash common.Hash) string {
	return o.topics[sigHash]
}

// GetSigHashes get sig hash list.
func (o *L1MessageQueue) GetSigHashes() []common.Hash {
	if len(o.parsers) == 0 {
		return nil
	}
	var sigHashes = make([]common.Hash, 0, len(o.parsers))
	for id := range o.parsers {
		sigHashes = append(sigHashes, id)
	}
	return sigHashes
}

// ParseLog parse the log if parse func is exist.
func (o *L1MessageQueue) ParseLog(vLog *types.Log) (bool, error) {
	_id := vLog.Topics[0]
	if parse, exist := o.parsers[_id]; exist {
		return true, parse(vLog)
	}
	return false, nil
}

// RegisterDequeueTransaction, the DequeueTransaction event ID is 0xc77f792f838ae38399ac31acc3348389aeb110ce7bedf3cfdbdd5e6679267970.
func (o *L1MessageQueue) RegisterDequeueTransaction(handler func(vLog *types.Log, data *L1MessageQueueDequeueTransactionEvent) error) {
	_id := o.ABI.Events["DequeueTransaction"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1MessageQueueDequeueTransactionEvent)
		if err := o.L1MessageQueueCaller.contract.UnpackLog(event, "DequeueTransaction", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "DequeueTransaction"
}

// RegisterDropTransaction, the DropTransaction event ID is 0x43a375005206d20a83abc71722cba68c24434a8dc1f583775be7c3fde0396cbf.
func (o *L1MessageQueue) RegisterDropTransaction(handler func(vLog *types.Log, data *L1MessageQueueDropTransactionEvent) error) {
	_id := o.ABI.Events["DropTransaction"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1MessageQueueDropTransactionEvent)
		if err := o.L1MessageQueueCaller.contract.UnpackLog(event, "DropTransaction", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "DropTransaction"
}

// RegisterInitialized, the Initialized event ID is 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
func (o *L1MessageQueue) RegisterInitialized(handler func(vLog *types.Log, data *L1MessageQueueInitializedEvent) error) {
	_id := o.ABI.Events["Initialized"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1MessageQueueInitializedEvent)
		if err := o.L1MessageQueueCaller.contract.UnpackLog(event, "Initialized", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "Initialized"
}

// RegisterOwnershipTransferred, the OwnershipTransferred event ID is 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
func (o *L1MessageQueue) RegisterOwnershipTransferred(handler func(vLog *types.Log, data *L1MessageQueueOwnershipTransferredEvent) error) {
	_id := o.ABI.Events["OwnershipTransferred"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1MessageQueueOwnershipTransferredEvent)
		if err := o.L1MessageQueueCaller.contract.UnpackLog(event, "OwnershipTransferred", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "OwnershipTransferred"
}

// RegisterQueueTransaction, the QueueTransaction event ID is 0x69cfcb8e6d4192b8aba9902243912587f37e550d75c1fa801491fce26717f37e.
func (o *L1MessageQueue) RegisterQueueTransaction(handler func(vLog *types.Log, data *L1MessageQueueQueueTransactionEvent) error) {
	_id := o.ABI.Events["QueueTransaction"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1MessageQueueQueueTransactionEvent)
		if err := o.L1MessageQueueCaller.contract.UnpackLog(event, "QueueTransaction", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "QueueTransaction"
}

// RegisterUpdateEnforcedTxGateway, the UpdateEnforcedTxGateway event ID is 0x5fd1d27c789fb50eafa108fba89345986a66d9d0aba85d48adee09f5e3307bbd.
func (o *L1MessageQueue) RegisterUpdateEnforcedTxGateway(handler func(vLog *types.Log, data *L1MessageQueueUpdateEnforcedTxGatewayEvent) error) {
	_id := o.ABI.Events["UpdateEnforcedTxGateway"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1MessageQueueUpdateEnforcedTxGatewayEvent)
		if err := o.L1MessageQueueCaller.contract.UnpackLog(event, "UpdateEnforcedTxGateway", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "UpdateEnforcedTxGateway"
}

// RegisterUpdateGasOracle, the UpdateGasOracle event ID is 0x9ed5ec28f252b3e7f62f1ace8e54c5ebabf4c61cc2a7c33a806365b2ff7ecc5e.
func (o *L1MessageQueue) RegisterUpdateGasOracle(handler func(vLog *types.Log, data *L1MessageQueueUpdateGasOracleEvent) error) {
	_id := o.ABI.Events["UpdateGasOracle"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1MessageQueueUpdateGasOracleEvent)
		if err := o.L1MessageQueueCaller.contract.UnpackLog(event, "UpdateGasOracle", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "UpdateGasOracle"
}

// RegisterUpdateMaxGasLimit, the UpdateMaxGasLimit event ID is 0xa030881e03ff723954dd0d35500564afab9603555d09d4456a32436f2b2373c5.
func (o *L1MessageQueue) RegisterUpdateMaxGasLimit(handler func(vLog *types.Log, data *L1MessageQueueUpdateMaxGasLimitEvent) error) {
	_id := o.ABI.Events["UpdateMaxGasLimit"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1MessageQueueUpdateMaxGasLimitEvent)
		if err := o.L1MessageQueueCaller.contract.UnpackLog(event, "UpdateMaxGasLimit", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "UpdateMaxGasLimit"
}

// L1MessageQueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1MessageQueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1MessageQueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1MessageQueueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewL1MessageQueue creates a new instance of L1MessageQueue, bound to a specific deployed contract.
func NewL1MessageQueue(address common.Address, backend bind.ContractBackend) (*L1MessageQueue, error) {
	contract, err := bindL1MessageQueue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	sigAbi, err := L1MessageQueueMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	topics := make(map[common.Hash]string, len(sigAbi.Events))
	for name, _abi := range sigAbi.Events {
		topics[_abi.ID] = name
	}

	return &L1MessageQueue{
		Name:                     "L1MessageQueue",
		ABI:                      sigAbi,
		Address:                  address,
		topics:                   topics,
		parsers:                  map[common.Hash]func(log *types.Log) error{},
		L1MessageQueueCaller:     L1MessageQueueCaller{contract: contract},
		L1MessageQueueTransactor: L1MessageQueueTransactor{contract: contract},
	}, nil
}

// NewL1MessageQueueCaller creates a new read-only instance of L1MessageQueue, bound to a specific deployed contract.
func NewL1MessageQueueCaller(address common.Address, caller bind.ContractCaller) (*L1MessageQueueCaller, error) {
	contract, err := bindL1MessageQueue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueCaller{contract: contract}, nil
}

// NewL1MessageQueueTransactor creates a new write-only instance of L1MessageQueue, bound to a specific deployed contract.
func NewL1MessageQueueTransactor(address common.Address, transactor bind.ContractTransactor) (*L1MessageQueueTransactor, error) {
	contract, err := bindL1MessageQueue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1MessageQueueTransactor{contract: contract}, nil
}

// bindL1MessageQueue binds a generic wrapper to an already deployed contract.
func bindL1MessageQueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1MessageQueueMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// CalculateIntrinsicGasFee is a free data retrieval call binding the contract method 0xe172d3a1.
//
// Solidity: function calculateIntrinsicGasFee(bytes _calldata) view returns(uint256)
func (_L1MessageQueue *L1MessageQueueCaller) CalculateIntrinsicGasFee(opts *bind.CallOpts, _calldata []byte) (*big.Int, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "calculateIntrinsicGasFee", _calldata)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeTransactionHash is a free data retrieval call binding the contract method 0x5ad9945a.
//
// Solidity: function computeTransactionHash(address _sender, uint256 _queueIndex, uint256 _value, address _target, uint256 _gasLimit, bytes _data) pure returns(bytes32)
func (_L1MessageQueue *L1MessageQueueCaller) ComputeTransactionHash(opts *bind.CallOpts, _sender common.Address, _queueIndex *big.Int, _value *big.Int, _target common.Address, _gasLimit *big.Int, _data []byte) ([32]byte, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "computeTransactionHash", _sender, _queueIndex, _value, _target, _gasLimit, _data)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EnforcedTxGateway is a free data retrieval call binding the contract method 0x3e83496c.
//
// Solidity: function enforcedTxGateway() view returns(address)
func (_L1MessageQueue *L1MessageQueueCaller) EnforcedTxGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "enforcedTxGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EstimateCrossDomainMessageFee is a free data retrieval call binding the contract method 0xd7704bae.
//
// Solidity: function estimateCrossDomainMessageFee(uint256 _gasLimit) view returns(uint256)
func (_L1MessageQueue *L1MessageQueueCaller) EstimateCrossDomainMessageFee(opts *bind.CallOpts, _gasLimit *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "estimateCrossDomainMessageFee", _gasLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_L1MessageQueue *L1MessageQueueCaller) GasOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "gasOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCrossDomainMessage is a free data retrieval call binding the contract method 0xae453cd5.
//
// Solidity: function getCrossDomainMessage(uint256 _queueIndex) view returns(bytes32)
func (_L1MessageQueue *L1MessageQueueCaller) GetCrossDomainMessage(opts *bind.CallOpts, _queueIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "getCrossDomainMessage", _queueIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MaxGasLimit is a free data retrieval call binding the contract method 0x5e45da23.
//
// Solidity: function maxGasLimit() view returns(uint256)
func (_L1MessageQueue *L1MessageQueueCaller) MaxGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "maxGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageQueue is a free data retrieval call binding the contract method 0x29aa604b.
//
// Solidity: function messageQueue(uint256 ) view returns(bytes32)
func (_L1MessageQueue *L1MessageQueueCaller) MessageQueue(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "messageQueue", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1MessageQueue *L1MessageQueueCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NextCrossDomainMessageIndex is a free data retrieval call binding the contract method 0xfd0ad31e.
//
// Solidity: function nextCrossDomainMessageIndex() view returns(uint256)
func (_L1MessageQueue *L1MessageQueueCaller) NextCrossDomainMessageIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "nextCrossDomainMessageIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1MessageQueue *L1MessageQueueCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingQueueIndex is a free data retrieval call binding the contract method 0xa85006ca.
//
// Solidity: function pendingQueueIndex() view returns(uint256)
func (_L1MessageQueue *L1MessageQueueCaller) PendingQueueIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "pendingQueueIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScrollChain is a free data retrieval call binding the contract method 0x897630dd.
//
// Solidity: function scrollChain() view returns(address)
func (_L1MessageQueue *L1MessageQueueCaller) ScrollChain(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1MessageQueue.contract.Call(opts, &out, "scrollChain")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AppendCrossDomainMessage is a paid mutator transaction binding the contract method 0x9b159782.
//
// Solidity: function appendCrossDomainMessage(address _target, uint256 _gasLimit, bytes _data) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) AppendCrossDomainMessage(opts *bind.TransactOpts, _target common.Address, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "appendCrossDomainMessage", _target, _gasLimit, _data)
}

// AppendEnforcedTransaction is a paid mutator transaction binding the contract method 0xbdc6f0a0.
//
// Solidity: function appendEnforcedTransaction(address _sender, address _target, uint256 _value, uint256 _gasLimit, bytes _data) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) AppendEnforcedTransaction(opts *bind.TransactOpts, _sender common.Address, _target common.Address, _value *big.Int, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "appendEnforcedTransaction", _sender, _target, _value, _gasLimit, _data)
}

// DropCrossDomainMessage is a paid mutator transaction binding the contract method 0x91652461.
//
// Solidity: function dropCrossDomainMessage(uint256 _index) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) DropCrossDomainMessage(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "dropCrossDomainMessage", _index)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address _messenger, address _scrollChain, address _enforcedTxGateway, address _gasOracle, uint256 _maxGasLimit) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) Initialize(opts *bind.TransactOpts, _messenger common.Address, _scrollChain common.Address, _enforcedTxGateway common.Address, _gasOracle common.Address, _maxGasLimit *big.Int) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "initialize", _messenger, _scrollChain, _enforcedTxGateway, _gasOracle, _maxGasLimit)
}

// PopCrossDomainMessage is a paid mutator transaction binding the contract method 0x55f613ce.
//
// Solidity: function popCrossDomainMessage(uint256 _startIndex, uint256 _count, uint256 _skippedBitmap) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) PopCrossDomainMessage(opts *bind.TransactOpts, _startIndex *big.Int, _count *big.Int, _skippedBitmap *big.Int) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "popCrossDomainMessage", _startIndex, _count, _skippedBitmap)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1MessageQueue *L1MessageQueueTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "renounceOwnership")
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "transferOwnership", newOwner)
}

// UpdateEnforcedTxGateway is a paid mutator transaction binding the contract method 0xc94864e1.
//
// Solidity: function updateEnforcedTxGateway(address _newGateway) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) UpdateEnforcedTxGateway(opts *bind.TransactOpts, _newGateway common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "updateEnforcedTxGateway", _newGateway)
}

// UpdateGasOracle is a paid mutator transaction binding the contract method 0x70cee67f.
//
// Solidity: function updateGasOracle(address _newGasOracle) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) UpdateGasOracle(opts *bind.TransactOpts, _newGasOracle common.Address) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "updateGasOracle", _newGasOracle)
}

// UpdateMaxGasLimit is a paid mutator transaction binding the contract method 0xd5ad4a97.
//
// Solidity: function updateMaxGasLimit(uint256 _newMaxGasLimit) returns()
func (_L1MessageQueue *L1MessageQueueTransactor) UpdateMaxGasLimit(opts *bind.TransactOpts, _newMaxGasLimit *big.Int) (*types.Transaction, error) {
	return _L1MessageQueue.contract.Transact(opts, "updateMaxGasLimit", _newMaxGasLimit)
}

// L1MessageQueueDequeueTransaction represents a DequeueTransaction event raised by the L1MessageQueue contract.
type L1MessageQueueDequeueTransactionEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	StartIndex    *big.Int
	Count         *big.Int
	SkippedBitmap *big.Int
}

// L1MessageQueueDropTransaction represents a DropTransaction event raised by the L1MessageQueue contract.
type L1MessageQueueDropTransactionEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	Index *big.Int
}

// L1MessageQueueInitialized represents a Initialized event raised by the L1MessageQueue contract.
type L1MessageQueueInitializedEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	Version uint8
}

// L1MessageQueueOwnershipTransferred represents a OwnershipTransferred event raised by the L1MessageQueue contract.
type L1MessageQueueOwnershipTransferredEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	PreviousOwner common.Address
	NewOwner      common.Address
}

// L1MessageQueueQueueTransaction represents a QueueTransaction event raised by the L1MessageQueue contract.
type L1MessageQueueQueueTransactionEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	Sender     common.Address
	Target     common.Address
	Value      *big.Int
	QueueIndex uint64
	GasLimit   *big.Int
	Data       []byte
}

// L1MessageQueueUpdateEnforcedTxGateway represents a UpdateEnforcedTxGateway event raised by the L1MessageQueue contract.
type L1MessageQueueUpdateEnforcedTxGatewayEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	OldGateway common.Address
	NewGateway common.Address
}

// L1MessageQueueUpdateGasOracle represents a UpdateGasOracle event raised by the L1MessageQueue contract.
type L1MessageQueueUpdateGasOracleEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	OldGasOracle common.Address
	NewGasOracle common.Address
}

// L1MessageQueueUpdateMaxGasLimit represents a UpdateMaxGasLimit event raised by the L1MessageQueue contract.
type L1MessageQueueUpdateMaxGasLimitEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	OldMaxGasLimit *big.Int
	NewMaxGasLimit *big.Int
}
