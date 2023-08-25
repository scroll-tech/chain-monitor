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

// L2ETHGatewayMetaData contains all meta data concerning the L2ETHGateway contract.
var (
	L2ETHGatewayMetaData = &bind.MetaData{
		ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false}],\"type\":\"event\",\"name\":\"FinalizeDepositETH\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Initialized\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false}],\"type\":\"event\",\"name\":\"WithdrawETH\",\"anonymous\":false},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"finalizeDepositETH\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"renounceOwnership\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferOwnership\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"withdrawETH\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"withdrawETH\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"withdrawETHAndCall\"}]",
	}
)

// L2ETHGateway is an auto generated Go binding around an Ethereum contract.
type L2ETHGateway struct {
	Name    string
	Address common.Address // contract address
	ABI     *abi.ABI       // L2ETHGatewayABI is the input ABI used to generate the binding from.

	topics  map[common.Hash]string
	parsers map[common.Hash]func(log *types.Log) error

	L2ETHGatewayCaller     // Read-only binding to the contract
	L2ETHGatewayTransactor // Write-only binding to the contract
}

// GetName return L2ETHGateway's Name.
func (o *L2ETHGateway) GetName() string {
	return o.Name
}

// GetAddress return L2ETHGateway's contract address.
func (o *L2ETHGateway) GetAddress() common.Address {
	return o.Address
}

// GetEventName get event name by event hash.
func (o *L2ETHGateway) GetEventName(sigHash common.Hash) string {
	return o.topics[sigHash]
}

// GetSigHashes get sig hash list.
func (o *L2ETHGateway) GetSigHashes() []common.Hash {
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
func (o *L2ETHGateway) ParseLog(vLog *types.Log) error {
	_id := vLog.Topics[0]
	if parse, exist := o.parsers[_id]; exist {
		return parse(vLog)
	}
	return nil
}

// RegisterFinalizeDepositETH, the FinalizeDepositETH event ID is 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
func (o *L2ETHGateway) RegisterFinalizeDepositETH(handler func(vLog *types.Log, data *L2ETHGatewayFinalizeDepositETHEvent) error) {
	o.parsers[o.ABI.Events["FinalizeDepositETH"].ID] = func(log *types.Log) error {
		event := new(L2ETHGatewayFinalizeDepositETHEvent)
		if err := o.L2ETHGatewayCaller.contract.UnpackLog(event, "FinalizeDepositETH", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// RegisterInitialized, the Initialized event ID is 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
func (o *L2ETHGateway) RegisterInitialized(handler func(vLog *types.Log, data *L2ETHGatewayInitializedEvent) error) {
	o.parsers[o.ABI.Events["Initialized"].ID] = func(log *types.Log) error {
		event := new(L2ETHGatewayInitializedEvent)
		if err := o.L2ETHGatewayCaller.contract.UnpackLog(event, "Initialized", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// RegisterOwnershipTransferred, the OwnershipTransferred event ID is 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
func (o *L2ETHGateway) RegisterOwnershipTransferred(handler func(vLog *types.Log, data *L2ETHGatewayOwnershipTransferredEvent) error) {
	o.parsers[o.ABI.Events["OwnershipTransferred"].ID] = func(log *types.Log) error {
		event := new(L2ETHGatewayOwnershipTransferredEvent)
		if err := o.L2ETHGatewayCaller.contract.UnpackLog(event, "OwnershipTransferred", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// RegisterWithdrawETH, the WithdrawETH event ID is 0xd8ed6eaa9a7a8980d7901e911fde6686810b989d3082182d1d3a3df6306ce20e.
func (o *L2ETHGateway) RegisterWithdrawETH(handler func(vLog *types.Log, data *L2ETHGatewayWithdrawETHEvent) error) {
	o.parsers[o.ABI.Events["WithdrawETH"].ID] = func(log *types.Log) error {
		event := new(L2ETHGatewayWithdrawETHEvent)
		if err := o.L2ETHGatewayCaller.contract.UnpackLog(event, "WithdrawETH", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// L2ETHGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ETHGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ETHGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ETHGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewL2ETHGateway creates a new instance of L2ETHGateway, bound to a specific deployed contract.
func NewL2ETHGateway(address common.Address, backend bind.ContractBackend) (*L2ETHGateway, error) {
	contract, err := bindL2ETHGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	sigAbi, err := L2ETHGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	topics := make(map[common.Hash]string, len(sigAbi.Events))
	for name, _abi := range sigAbi.Events {
		topics[_abi.ID] = name
	}

	return &L2ETHGateway{
		Name:                   "L2ETHGateway",
		ABI:                    sigAbi,
		Address:                address,
		topics:                 topics,
		parsers:                map[common.Hash]func(log *types.Log) error{},
		L2ETHGatewayCaller:     L2ETHGatewayCaller{contract: contract},
		L2ETHGatewayTransactor: L2ETHGatewayTransactor{contract: contract},
	}, nil
}

// NewL2ETHGatewayCaller creates a new read-only instance of L2ETHGateway, bound to a specific deployed contract.
func NewL2ETHGatewayCaller(address common.Address, caller bind.ContractCaller) (*L2ETHGatewayCaller, error) {
	contract, err := bindL2ETHGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ETHGatewayCaller{contract: contract}, nil
}

// NewL2ETHGatewayTransactor creates a new write-only instance of L2ETHGateway, bound to a specific deployed contract.
func NewL2ETHGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L2ETHGatewayTransactor, error) {
	contract, err := bindL2ETHGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ETHGatewayTransactor{contract: contract}, nil
}

// bindL2ETHGateway binds a generic wrapper to an already deployed contract.
func bindL2ETHGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2ETHGatewayMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2ETHGateway *L2ETHGatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ETHGateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ETHGateway *L2ETHGatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ETHGateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2ETHGateway *L2ETHGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ETHGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2ETHGateway *L2ETHGatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ETHGateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2ETHGateway *L2ETHGatewayTransactor) FinalizeDepositETH(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ETHGateway.contract.Transact(opts, "finalizeDepositETH", _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2ETHGateway *L2ETHGatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2ETHGateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2ETHGateway *L2ETHGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ETHGateway.contract.Transact(opts, "renounceOwnership")
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2ETHGateway *L2ETHGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2ETHGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ETHGateway *L2ETHGatewayTransactor) WithdrawETH(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ETHGateway.contract.Transact(opts, "withdrawETH", _to, _amount, _gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ETHGateway *L2ETHGatewayTransactor) WithdrawETH0(opts *bind.TransactOpts, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ETHGateway.contract.Transact(opts, "withdrawETH0", _amount, _gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2ETHGateway *L2ETHGatewayTransactor) WithdrawETHAndCall(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ETHGateway.contract.Transact(opts, "withdrawETHAndCall", _to, _amount, _data, _gasLimit)
}

// L2ETHGatewayFinalizeDepositETH represents a FinalizeDepositETH event raised by the L2ETHGateway contract.
type L2ETHGatewayFinalizeDepositETHEvent struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
}

// L2ETHGatewayInitialized represents a Initialized event raised by the L2ETHGateway contract.
type L2ETHGatewayInitializedEvent struct {
	Version uint8
}

// L2ETHGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L2ETHGateway contract.
type L2ETHGatewayOwnershipTransferredEvent struct {
	PreviousOwner common.Address
	NewOwner      common.Address
}

// L2ETHGatewayWithdrawETH represents a WithdrawETH event raised by the L2ETHGateway contract.
type L2ETHGatewayWithdrawETHEvent struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
}
