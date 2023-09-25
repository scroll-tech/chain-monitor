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

// L1ETHGatewayMetaData contains all meta data concerning the L1ETHGateway contract.
var (
	L1ETHGatewayMetaData = &bind.MetaData{
		ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false}],\"type\":\"event\",\"name\":\"DepositETH\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false}],\"type\":\"event\",\"name\":\"FinalizeWithdrawETH\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Initialized\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"RefundETH\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldRateLimiter\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_newRateLimiter\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"UpdateRateLimiter\",\"anonymous\":false},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"depositETH\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"depositETH\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"depositETHAndCall\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"finalizeWithdrawETH\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"onDropMessage\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"rateLimiter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"renounceOwnership\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferOwnership\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRateLimiter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateRateLimiter\"}]",
	}
	// L1ETHGatewayABI is the input ABI used to generate the binding from.
	L1ETHGatewayABI, _ = L1ETHGatewayMetaData.GetAbi()
)

// L1ETHGateway is an auto generated Go binding around an Ethereum contract.
type L1ETHGateway struct {
	Name    string
	Address common.Address // contract address
	ABI     *abi.ABI       // L1ETHGatewayABI is the input ABI used to generate the binding from.

	topics  map[common.Hash]string
	parsers map[common.Hash]func(log *types.Log) error

	L1ETHGatewayCaller     // Read-only binding to the contract
	L1ETHGatewayTransactor // Write-only binding to the contract
}

// GetName return L1ETHGateway's Name.
func (o *L1ETHGateway) GetName() string {
	return o.Name
}

// GetAddress return L1ETHGateway's contract address.
func (o *L1ETHGateway) GetAddress() common.Address {
	return o.Address
}

// GetEventName get event name by event hash.
func (o *L1ETHGateway) GetEventName(sigHash common.Hash) string {
	return o.topics[sigHash]
}

// GetSigHashes get sig hash list.
func (o *L1ETHGateway) GetSigHashes() []common.Hash {
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
func (o *L1ETHGateway) ParseLog(vLog *types.Log) (bool, error) {
	_id := vLog.Topics[0]
	if parse, exist := o.parsers[_id]; exist {
		return true, parse(vLog)
	}
	return false, nil
}

// RegisterDepositETH, the DepositETH event ID is 0x6670de856ec8bf5cb2b7e957c5dc24759716056f79d97ea5e7c939ca0ba5a675.
func (o *L1ETHGateway) RegisterDepositETH(handler func(vLog *types.Log, data *L1ETHGatewayDepositETHEvent) error) {
	_id := o.ABI.Events["DepositETH"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ETHGatewayDepositETHEvent)
		if err := o.L1ETHGatewayCaller.contract.UnpackLog(event, "DepositETH", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "DepositETH"
}

// RegisterFinalizeWithdrawETH, the FinalizeWithdrawETH event ID is 0x96db5d1cee1dd2760826bb56fabd9c9f6e978083e0a8b88559c741a29e9746e7.
func (o *L1ETHGateway) RegisterFinalizeWithdrawETH(handler func(vLog *types.Log, data *L1ETHGatewayFinalizeWithdrawETHEvent) error) {
	_id := o.ABI.Events["FinalizeWithdrawETH"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ETHGatewayFinalizeWithdrawETHEvent)
		if err := o.L1ETHGatewayCaller.contract.UnpackLog(event, "FinalizeWithdrawETH", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "FinalizeWithdrawETH"
}

// RegisterInitialized, the Initialized event ID is 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
func (o *L1ETHGateway) RegisterInitialized(handler func(vLog *types.Log, data *L1ETHGatewayInitializedEvent) error) {
	_id := o.ABI.Events["Initialized"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ETHGatewayInitializedEvent)
		if err := o.L1ETHGatewayCaller.contract.UnpackLog(event, "Initialized", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "Initialized"
}

// RegisterOwnershipTransferred, the OwnershipTransferred event ID is 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
func (o *L1ETHGateway) RegisterOwnershipTransferred(handler func(vLog *types.Log, data *L1ETHGatewayOwnershipTransferredEvent) error) {
	_id := o.ABI.Events["OwnershipTransferred"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ETHGatewayOwnershipTransferredEvent)
		if err := o.L1ETHGatewayCaller.contract.UnpackLog(event, "OwnershipTransferred", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "OwnershipTransferred"
}

// RegisterRefundETH, the RefundETH event ID is 0x289360176646a5f99cb4b6300628426dca46b723f40db3c04449d6ed1745a0e7.
func (o *L1ETHGateway) RegisterRefundETH(handler func(vLog *types.Log, data *L1ETHGatewayRefundETHEvent) error) {
	_id := o.ABI.Events["RefundETH"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ETHGatewayRefundETHEvent)
		if err := o.L1ETHGatewayCaller.contract.UnpackLog(event, "RefundETH", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "RefundETH"
}

// RegisterUpdateRateLimiter, the UpdateRateLimiter event ID is 0x53d49a5b37a3ed13692a8b217b8c79f53b0144c069de23e0cc5704d89bc23006.
func (o *L1ETHGateway) RegisterUpdateRateLimiter(handler func(vLog *types.Log, data *L1ETHGatewayUpdateRateLimiterEvent) error) {
	_id := o.ABI.Events["UpdateRateLimiter"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ETHGatewayUpdateRateLimiterEvent)
		if err := o.L1ETHGatewayCaller.contract.UnpackLog(event, "UpdateRateLimiter", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "UpdateRateLimiter"
}

// L1ETHGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1ETHGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ETHGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1ETHGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewL1ETHGateway creates a new instance of L1ETHGateway, bound to a specific deployed contract.
func NewL1ETHGateway(address common.Address, backend bind.ContractBackend) (*L1ETHGateway, error) {
	contract, err := bindL1ETHGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	sigAbi, err := L1ETHGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	topics := make(map[common.Hash]string, len(sigAbi.Events))
	for name, _abi := range sigAbi.Events {
		topics[_abi.ID] = name
	}

	return &L1ETHGateway{
		Name:                   "L1ETHGateway",
		ABI:                    sigAbi,
		Address:                address,
		topics:                 topics,
		parsers:                map[common.Hash]func(log *types.Log) error{},
		L1ETHGatewayCaller:     L1ETHGatewayCaller{contract: contract},
		L1ETHGatewayTransactor: L1ETHGatewayTransactor{contract: contract},
	}, nil
}

// NewL1ETHGatewayCaller creates a new read-only instance of L1ETHGateway, bound to a specific deployed contract.
func NewL1ETHGatewayCaller(address common.Address, caller bind.ContractCaller) (*L1ETHGatewayCaller, error) {
	contract, err := bindL1ETHGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1ETHGatewayCaller{contract: contract}, nil
}

// NewL1ETHGatewayTransactor creates a new write-only instance of L1ETHGateway, bound to a specific deployed contract.
func NewL1ETHGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L1ETHGatewayTransactor, error) {
	contract, err := bindL1ETHGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1ETHGatewayTransactor{contract: contract}, nil
}

// bindL1ETHGateway binds a generic wrapper to an already deployed contract.
func bindL1ETHGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1ETHGatewayMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1ETHGateway *L1ETHGatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ETHGateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1ETHGateway *L1ETHGatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ETHGateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1ETHGateway *L1ETHGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ETHGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RateLimiter is a free data retrieval call binding the contract method 0x53d4fe33.
//
// Solidity: function rateLimiter() view returns(address)
func (_L1ETHGateway *L1ETHGatewayCaller) RateLimiter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ETHGateway.contract.Call(opts, &out, "rateLimiter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1ETHGateway *L1ETHGatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ETHGateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositETH is a paid mutator transaction binding the contract method 0x9f8420b3.
//
// Solidity: function depositETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1ETHGateway *L1ETHGatewayTransactor) DepositETH(opts *bind.TransactOpts, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ETHGateway.contract.Transact(opts, "depositETH", _amount, _gasLimit)
}

// DepositETH0 is a paid mutator transaction binding the contract method 0xce0b63ce.
//
// Solidity: function depositETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1ETHGateway *L1ETHGatewayTransactor) DepositETH0(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ETHGateway.contract.Transact(opts, "depositETH0", _to, _amount, _gasLimit)
}

// DepositETHAndCall is a paid mutator transaction binding the contract method 0xaac476f8.
//
// Solidity: function depositETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1ETHGateway *L1ETHGatewayTransactor) DepositETHAndCall(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ETHGateway.contract.Transact(opts, "depositETHAndCall", _to, _amount, _data, _gasLimit)
}

// FinalizeWithdrawETH is a paid mutator transaction binding the contract method 0x8eaac8a3.
//
// Solidity: function finalizeWithdrawETH(address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1ETHGateway *L1ETHGatewayTransactor) FinalizeWithdrawETH(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1ETHGateway.contract.Transact(opts, "finalizeWithdrawETH", _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L1ETHGateway *L1ETHGatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1ETHGateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1ETHGateway *L1ETHGatewayTransactor) OnDropMessage(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _L1ETHGateway.contract.Transact(opts, "onDropMessage", _message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1ETHGateway *L1ETHGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ETHGateway.contract.Transact(opts, "renounceOwnership")
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1ETHGateway *L1ETHGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1ETHGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// UpdateRateLimiter is a paid mutator transaction binding the contract method 0xe157820a.
//
// Solidity: function updateRateLimiter(address _newRateLimiter) returns()
func (_L1ETHGateway *L1ETHGatewayTransactor) UpdateRateLimiter(opts *bind.TransactOpts, _newRateLimiter common.Address) (*types.Transaction, error) {
	return _L1ETHGateway.contract.Transact(opts, "updateRateLimiter", _newRateLimiter)
}

// L1ETHGatewayDepositETH represents a DepositETH event raised by the L1ETHGateway contract.
type L1ETHGatewayDepositETHEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
}

// L1ETHGatewayFinalizeWithdrawETH represents a FinalizeWithdrawETH event raised by the L1ETHGateway contract.
type L1ETHGatewayFinalizeWithdrawETHEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
}

// L1ETHGatewayInitialized represents a Initialized event raised by the L1ETHGateway contract.
type L1ETHGatewayInitializedEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	Version uint8
}

// L1ETHGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L1ETHGateway contract.
type L1ETHGatewayOwnershipTransferredEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	PreviousOwner common.Address
	NewOwner      common.Address
}

// L1ETHGatewayRefundETH represents a RefundETH event raised by the L1ETHGateway contract.
type L1ETHGatewayRefundETHEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	Recipient common.Address
	Amount    *big.Int
}

// L1ETHGatewayUpdateRateLimiter represents a UpdateRateLimiter event raised by the L1ETHGateway contract.
type L1ETHGatewayUpdateRateLimiterEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	OldRateLimiter common.Address
	NewRateLimiter common.Address
}
