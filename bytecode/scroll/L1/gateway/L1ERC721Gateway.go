// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gateway

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

// L1ERC721GatewayMetaData contains all meta data concerning the L1ERC721Gateway contract.
var (
	L1ERC721GatewayMetaData = &bind.MetaData{
		ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\",\"indexed\":false}],\"type\":\"event\",\"name\":\"BatchDepositERC721\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"indexed\":false}],\"type\":\"event\",\"name\":\"BatchRefundERC721\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"DepositERC721\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\",\"indexed\":false}],\"type\":\"event\",\"name\":\"FinalizeBatchWithdrawERC721\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"FinalizeWithdrawERC721\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Initialized\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"RefundERC721\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldRateLimiter\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_newRateLimiter\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"UpdateRateLimiter\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"oldL2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"newL2Token\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"UpdateTokenMapping\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"batchDepositERC721\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"batchDepositERC721\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"depositERC721\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"depositERC721\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"finalizeBatchWithdrawERC721\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"finalizeWithdrawERC721\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"onDropMessage\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"rateLimiter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"renounceOwnership\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"tokenMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferOwnership\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRateLimiter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateRateLimiter\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateTokenMapping\"}]",
	}
	// L1ERC721GatewayABI is the input ABI used to generate the binding from.
	L1ERC721GatewayABI, _ = L1ERC721GatewayMetaData.GetAbi()
)

// L1ERC721Gateway is an auto generated Go binding around an Ethereum contract.
type L1ERC721Gateway struct {
	Name    string
	Address common.Address // contract address
	ABI     *abi.ABI       // L1ERC721GatewayABI is the input ABI used to generate the binding from.

	topics  map[common.Hash]string
	parsers map[common.Hash]func(log *types.Log) error

	L1ERC721GatewayCaller     // Read-only binding to the contract
	L1ERC721GatewayTransactor // Write-only binding to the contract
}

// GetName return L1ERC721Gateway's Name.
func (o *L1ERC721Gateway) GetName() string {
	return o.Name
}

// GetAddress return L1ERC721Gateway's contract address.
func (o *L1ERC721Gateway) GetAddress() common.Address {
	return o.Address
}

// GetEventName get event name by event hash.
func (o *L1ERC721Gateway) GetEventName(sigHash common.Hash) string {
	return o.topics[sigHash]
}

// GetSigHashes get sig hash list.
func (o *L1ERC721Gateway) GetSigHashes() []common.Hash {
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
func (o *L1ERC721Gateway) ParseLog(vLog *types.Log) (bool, error) {
	_id := vLog.Topics[0]
	if parse, exist := o.parsers[_id]; exist {
		return true, parse(vLog)
	}
	return false, nil
}

// RegisterBatchDepositERC721, the BatchDepositERC721 event ID is 0xf05915e3b4fbd6f61b8b6f80b07f10e1cad039ccc7abe7c7fec115d038fe3dd6.
func (o *L1ERC721Gateway) RegisterBatchDepositERC721(handler func(vLog *types.Log, data *L1ERC721GatewayBatchDepositERC721Event) error) {
	_id := o.ABI.Events["BatchDepositERC721"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ERC721GatewayBatchDepositERC721Event)
		if err := o.L1ERC721GatewayCaller.contract.UnpackLog(event, "BatchDepositERC721", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "BatchDepositERC721"
}

// RegisterBatchRefundERC721, the BatchRefundERC721 event ID is 0x998a3ef0a23771412ff48d871a2288502a89da39c5db04a2a66e5eb85586cc22.
func (o *L1ERC721Gateway) RegisterBatchRefundERC721(handler func(vLog *types.Log, data *L1ERC721GatewayBatchRefundERC721Event) error) {
	_id := o.ABI.Events["BatchRefundERC721"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ERC721GatewayBatchRefundERC721Event)
		if err := o.L1ERC721GatewayCaller.contract.UnpackLog(event, "BatchRefundERC721", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "BatchRefundERC721"
}

// RegisterDepositERC721, the DepositERC721 event ID is 0xfc1d17c06ff1e4678321cc30660a73f3f1436df8195108a288d3159a961febec.
func (o *L1ERC721Gateway) RegisterDepositERC721(handler func(vLog *types.Log, data *L1ERC721GatewayDepositERC721Event) error) {
	_id := o.ABI.Events["DepositERC721"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ERC721GatewayDepositERC721Event)
		if err := o.L1ERC721GatewayCaller.contract.UnpackLog(event, "DepositERC721", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "DepositERC721"
}

// RegisterFinalizeBatchWithdrawERC721, the FinalizeBatchWithdrawERC721 event ID is 0x9b8e51c8f180115b421b26c9042287d6bf95e0ce9c0c5434784e2af3d0b9de7d.
func (o *L1ERC721Gateway) RegisterFinalizeBatchWithdrawERC721(handler func(vLog *types.Log, data *L1ERC721GatewayFinalizeBatchWithdrawERC721Event) error) {
	_id := o.ABI.Events["FinalizeBatchWithdrawERC721"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ERC721GatewayFinalizeBatchWithdrawERC721Event)
		if err := o.L1ERC721GatewayCaller.contract.UnpackLog(event, "FinalizeBatchWithdrawERC721", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "FinalizeBatchWithdrawERC721"
}

// RegisterFinalizeWithdrawERC721, the FinalizeWithdrawERC721 event ID is 0xacdbfefc030b5ccccd5f60ca6d9ca371c6d6d6956fe16ebe10f81920198206e9.
func (o *L1ERC721Gateway) RegisterFinalizeWithdrawERC721(handler func(vLog *types.Log, data *L1ERC721GatewayFinalizeWithdrawERC721Event) error) {
	_id := o.ABI.Events["FinalizeWithdrawERC721"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ERC721GatewayFinalizeWithdrawERC721Event)
		if err := o.L1ERC721GatewayCaller.contract.UnpackLog(event, "FinalizeWithdrawERC721", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "FinalizeWithdrawERC721"
}

// RegisterInitialized, the Initialized event ID is 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
func (o *L1ERC721Gateway) RegisterInitialized(handler func(vLog *types.Log, data *L1ERC721GatewayInitializedEvent) error) {
	_id := o.ABI.Events["Initialized"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ERC721GatewayInitializedEvent)
		if err := o.L1ERC721GatewayCaller.contract.UnpackLog(event, "Initialized", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "Initialized"
}

// RegisterOwnershipTransferred, the OwnershipTransferred event ID is 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
func (o *L1ERC721Gateway) RegisterOwnershipTransferred(handler func(vLog *types.Log, data *L1ERC721GatewayOwnershipTransferredEvent) error) {
	_id := o.ABI.Events["OwnershipTransferred"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ERC721GatewayOwnershipTransferredEvent)
		if err := o.L1ERC721GatewayCaller.contract.UnpackLog(event, "OwnershipTransferred", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "OwnershipTransferred"
}

// RegisterRefundERC721, the RefundERC721 event ID is 0xb9a838365634e4fb87a9333edf0ea86f82836e361b311a125aefd14135581208.
func (o *L1ERC721Gateway) RegisterRefundERC721(handler func(vLog *types.Log, data *L1ERC721GatewayRefundERC721Event) error) {
	_id := o.ABI.Events["RefundERC721"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ERC721GatewayRefundERC721Event)
		if err := o.L1ERC721GatewayCaller.contract.UnpackLog(event, "RefundERC721", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "RefundERC721"
}

// RegisterUpdateRateLimiter, the UpdateRateLimiter event ID is 0x53d49a5b37a3ed13692a8b217b8c79f53b0144c069de23e0cc5704d89bc23006.
func (o *L1ERC721Gateway) RegisterUpdateRateLimiter(handler func(vLog *types.Log, data *L1ERC721GatewayUpdateRateLimiterEvent) error) {
	_id := o.ABI.Events["UpdateRateLimiter"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ERC721GatewayUpdateRateLimiterEvent)
		if err := o.L1ERC721GatewayCaller.contract.UnpackLog(event, "UpdateRateLimiter", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "UpdateRateLimiter"
}

// RegisterUpdateTokenMapping, the UpdateTokenMapping event ID is 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
func (o *L1ERC721Gateway) RegisterUpdateTokenMapping(handler func(vLog *types.Log, data *L1ERC721GatewayUpdateTokenMappingEvent) error) {
	_id := o.ABI.Events["UpdateTokenMapping"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ERC721GatewayUpdateTokenMappingEvent)
		if err := o.L1ERC721GatewayCaller.contract.UnpackLog(event, "UpdateTokenMapping", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "UpdateTokenMapping"
}

// L1ERC721GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1ERC721GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC721GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1ERC721GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewL1ERC721Gateway creates a new instance of L1ERC721Gateway, bound to a specific deployed contract.
func NewL1ERC721Gateway(address common.Address, backend bind.ContractBackend) (*L1ERC721Gateway, error) {
	contract, err := bindL1ERC721Gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	sigAbi, err := L1ERC721GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	topics := make(map[common.Hash]string, len(sigAbi.Events))
	for name, _abi := range sigAbi.Events {
		topics[_abi.ID] = name
	}

	return &L1ERC721Gateway{
		Name:                      "L1ERC721Gateway",
		ABI:                       sigAbi,
		Address:                   address,
		topics:                    topics,
		parsers:                   map[common.Hash]func(log *types.Log) error{},
		L1ERC721GatewayCaller:     L1ERC721GatewayCaller{contract: contract},
		L1ERC721GatewayTransactor: L1ERC721GatewayTransactor{contract: contract},
	}, nil
}

// NewL1ERC721GatewayCaller creates a new read-only instance of L1ERC721Gateway, bound to a specific deployed contract.
func NewL1ERC721GatewayCaller(address common.Address, caller bind.ContractCaller) (*L1ERC721GatewayCaller, error) {
	contract, err := bindL1ERC721Gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1ERC721GatewayCaller{contract: contract}, nil
}

// NewL1ERC721GatewayTransactor creates a new write-only instance of L1ERC721Gateway, bound to a specific deployed contract.
func NewL1ERC721GatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L1ERC721GatewayTransactor, error) {
	contract, err := bindL1ERC721Gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1ERC721GatewayTransactor{contract: contract}, nil
}

// bindL1ERC721Gateway binds a generic wrapper to an already deployed contract.
func bindL1ERC721Gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1ERC721GatewayMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC721Gateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC721Gateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC721Gateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RateLimiter is a free data retrieval call binding the contract method 0x53d4fe33.
//
// Solidity: function rateLimiter() view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewayCaller) RateLimiter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC721Gateway.contract.Call(opts, &out, "rateLimiter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC721Gateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L1ERC721Gateway *L1ERC721GatewayCaller) TokenMapping(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1ERC721Gateway.contract.Call(opts, &out, "tokenMapping", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BatchDepositERC721 is a paid mutator transaction binding the contract method 0x1b997a93.
//
// Solidity: function batchDepositERC721(address _token, address _to, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) BatchDepositERC721(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "batchDepositERC721", _token, _to, _tokenIds, _gasLimit)
}

// BatchDepositERC7210 is a paid mutator transaction binding the contract method 0xd96c8ecf.
//
// Solidity: function batchDepositERC721(address _token, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) BatchDepositERC7210(opts *bind.TransactOpts, _token common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "batchDepositERC7210", _token, _tokenIds, _gasLimit)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x0a7aa196.
//
// Solidity: function depositERC721(address _token, address _to, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) DepositERC721(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "depositERC721", _token, _to, _tokenId, _gasLimit)
}

// DepositERC7210 is a paid mutator transaction binding the contract method 0x45a4276b.
//
// Solidity: function depositERC721(address _token, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) DepositERC7210(opts *bind.TransactOpts, _token common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "depositERC7210", _token, _tokenId, _gasLimit)
}

// FinalizeBatchWithdrawERC721 is a paid mutator transaction binding the contract method 0x9f0a68b3.
//
// Solidity: function finalizeBatchWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds) returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) FinalizeBatchWithdrawERC721(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "finalizeBatchWithdrawERC721", _l1Token, _l2Token, _from, _to, _tokenIds)
}

// FinalizeWithdrawERC721 is a paid mutator transaction binding the contract method 0xd606b4dc.
//
// Solidity: function finalizeWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId) returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) FinalizeWithdrawERC721(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "finalizeWithdrawERC721", _l1Token, _l2Token, _from, _to, _tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "initialize", _counterpart, _messenger)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) OnDropMessage(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "onDropMessage", _message)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "renounceOwnership")
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// UpdateRateLimiter is a paid mutator transaction binding the contract method 0xe157820a.
//
// Solidity: function updateRateLimiter(address _newRateLimiter) returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) UpdateRateLimiter(opts *bind.TransactOpts, _newRateLimiter common.Address) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "updateRateLimiter", _newRateLimiter)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l1Token, address _l2Token) returns()
func (_L1ERC721Gateway *L1ERC721GatewayTransactor) UpdateTokenMapping(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address) (*types.Transaction, error) {
	return _L1ERC721Gateway.contract.Transact(opts, "updateTokenMapping", _l1Token, _l2Token)
}

// L1ERC721GatewayBatchDepositERC721 represents a BatchDepositERC721 event raised by the L1ERC721Gateway contract.
type L1ERC721GatewayBatchDepositERC721Event struct {
	Log *types.Log `json:"-" gorm:"-"`

	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
}

// L1ERC721GatewayBatchRefundERC721 represents a BatchRefundERC721 event raised by the L1ERC721Gateway contract.
type L1ERC721GatewayBatchRefundERC721Event struct {
	Log *types.Log `json:"-" gorm:"-"`

	Token     common.Address
	Recipient common.Address
	TokenIds  []*big.Int
}

// L1ERC721GatewayDepositERC721 represents a DepositERC721 event raised by the L1ERC721Gateway contract.
type L1ERC721GatewayDepositERC721Event struct {
	Log *types.Log `json:"-" gorm:"-"`

	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
}

// L1ERC721GatewayFinalizeBatchWithdrawERC721 represents a FinalizeBatchWithdrawERC721 event raised by the L1ERC721Gateway contract.
type L1ERC721GatewayFinalizeBatchWithdrawERC721Event struct {
	Log *types.Log `json:"-" gorm:"-"`

	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
}

// L1ERC721GatewayFinalizeWithdrawERC721 represents a FinalizeWithdrawERC721 event raised by the L1ERC721Gateway contract.
type L1ERC721GatewayFinalizeWithdrawERC721Event struct {
	Log *types.Log `json:"-" gorm:"-"`

	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
}

// L1ERC721GatewayInitialized represents a Initialized event raised by the L1ERC721Gateway contract.
type L1ERC721GatewayInitializedEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	Version uint8
}

// L1ERC721GatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L1ERC721Gateway contract.
type L1ERC721GatewayOwnershipTransferredEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	PreviousOwner common.Address
	NewOwner      common.Address
}

// L1ERC721GatewayRefundERC721 represents a RefundERC721 event raised by the L1ERC721Gateway contract.
type L1ERC721GatewayRefundERC721Event struct {
	Log *types.Log `json:"-" gorm:"-"`

	Token     common.Address
	Recipient common.Address
	TokenId   *big.Int
}

// L1ERC721GatewayUpdateRateLimiter represents a UpdateRateLimiter event raised by the L1ERC721Gateway contract.
type L1ERC721GatewayUpdateRateLimiterEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	OldRateLimiter common.Address
	NewRateLimiter common.Address
}

// L1ERC721GatewayUpdateTokenMapping represents a UpdateTokenMapping event raised by the L1ERC721Gateway contract.
type L1ERC721GatewayUpdateTokenMappingEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	L1Token    common.Address
	OldL2Token common.Address
	NewL2Token common.Address
}
