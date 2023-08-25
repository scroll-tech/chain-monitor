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

// L2ERC20GatewayMetaData contains all meta data concerning the L2ERC20Gateway contract.
var (
	L2ERC20GatewayMetaData = &bind.MetaData{
		ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false}],\"type\":\"event\",\"name\":\"FinalizeDepositERC20\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Initialized\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false}],\"type\":\"event\",\"name\":\"WithdrawERC20\",\"anonymous\":false},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"finalizeDepositERC20\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"getL1ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"renounceOwnership\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferOwnership\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"withdrawERC20\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"withdrawERC20\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"withdrawERC20AndCall\"}]",
	}
)

// L2ERC20Gateway is an auto generated Go binding around an Ethereum contract.
type L2ERC20Gateway struct {
	Name    string
	Address common.Address // contract address
	ABI     *abi.ABI       // L2ERC20GatewayABI is the input ABI used to generate the binding from.

	topics  map[common.Hash]string
	parsers map[common.Hash]func(log *types.Log) error

	L2ERC20GatewayCaller     // Read-only binding to the contract
	L2ERC20GatewayTransactor // Write-only binding to the contract
}

// GetName return L2ERC20Gateway's Name.
func (o *L2ERC20Gateway) GetName() string {
	return o.Name
}

// GetAddress return L2ERC20Gateway's contract address.
func (o *L2ERC20Gateway) GetAddress() common.Address {
	return o.Address
}

// GetEventName get event name by event hash.
func (o *L2ERC20Gateway) GetEventName(sigHash common.Hash) string {
	return o.topics[sigHash]
}

// GetSigHashes get sig hash list.
func (o *L2ERC20Gateway) GetSigHashes() []common.Hash {
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
func (o *L2ERC20Gateway) ParseLog(vLog *types.Log) error {
	_id := vLog.Topics[0]
	if parse, exist := o.parsers[_id]; exist {
		return parse(vLog)
	}
	return nil
}

// RegisterFinalizeDepositERC20, the FinalizeDepositERC20 event ID is 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
func (o *L2ERC20Gateway) RegisterFinalizeDepositERC20(handler func(vLog *types.Log, data *L2ERC20GatewayFinalizeDepositERC20Event) error) {
	o.parsers[o.ABI.Events["FinalizeDepositERC20"].ID] = func(log *types.Log) error {
		event := new(L2ERC20GatewayFinalizeDepositERC20Event)
		if err := o.L2ERC20GatewayCaller.contract.UnpackLog(event, "FinalizeDepositERC20", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// RegisterInitialized, the Initialized event ID is 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
func (o *L2ERC20Gateway) RegisterInitialized(handler func(vLog *types.Log, data *L2ERC20GatewayInitializedEvent) error) {
	o.parsers[o.ABI.Events["Initialized"].ID] = func(log *types.Log) error {
		event := new(L2ERC20GatewayInitializedEvent)
		if err := o.L2ERC20GatewayCaller.contract.UnpackLog(event, "Initialized", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// RegisterOwnershipTransferred, the OwnershipTransferred event ID is 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
func (o *L2ERC20Gateway) RegisterOwnershipTransferred(handler func(vLog *types.Log, data *L2ERC20GatewayOwnershipTransferredEvent) error) {
	o.parsers[o.ABI.Events["OwnershipTransferred"].ID] = func(log *types.Log) error {
		event := new(L2ERC20GatewayOwnershipTransferredEvent)
		if err := o.L2ERC20GatewayCaller.contract.UnpackLog(event, "OwnershipTransferred", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// RegisterWithdrawERC20, the WithdrawERC20 event ID is 0xd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e8.
func (o *L2ERC20Gateway) RegisterWithdrawERC20(handler func(vLog *types.Log, data *L2ERC20GatewayWithdrawERC20Event) error) {
	o.parsers[o.ABI.Events["WithdrawERC20"].ID] = func(log *types.Log) error {
		event := new(L2ERC20GatewayWithdrawERC20Event)
		if err := o.L2ERC20GatewayCaller.contract.UnpackLog(event, "WithdrawERC20", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// L2ERC20GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ERC20GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ERC20GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ERC20GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewL2ERC20Gateway creates a new instance of L2ERC20Gateway, bound to a specific deployed contract.
func NewL2ERC20Gateway(address common.Address, backend bind.ContractBackend) (*L2ERC20Gateway, error) {
	contract, err := bindL2ERC20Gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	sigAbi, err := L2ERC20GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	topics := make(map[common.Hash]string, len(sigAbi.Events))
	for name, _abi := range sigAbi.Events {
		topics[_abi.ID] = name
	}

	return &L2ERC20Gateway{
		Name:                     "L2ERC20Gateway",
		ABI:                      sigAbi,
		Address:                  address,
		topics:                   topics,
		parsers:                  map[common.Hash]func(log *types.Log) error{},
		L2ERC20GatewayCaller:     L2ERC20GatewayCaller{contract: contract},
		L2ERC20GatewayTransactor: L2ERC20GatewayTransactor{contract: contract},
	}, nil
}

// NewL2ERC20GatewayCaller creates a new read-only instance of L2ERC20Gateway, bound to a specific deployed contract.
func NewL2ERC20GatewayCaller(address common.Address, caller bind.ContractCaller) (*L2ERC20GatewayCaller, error) {
	contract, err := bindL2ERC20Gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ERC20GatewayCaller{contract: contract}, nil
}

// NewL2ERC20GatewayTransactor creates a new write-only instance of L2ERC20Gateway, bound to a specific deployed contract.
func NewL2ERC20GatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L2ERC20GatewayTransactor, error) {
	contract, err := bindL2ERC20Gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ERC20GatewayTransactor{contract: contract}, nil
}

// bindL2ERC20Gateway binds a generic wrapper to an already deployed contract.
func bindL2ERC20Gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2ERC20GatewayMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2ERC20Gateway *L2ERC20GatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC20Gateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address l2Token) view returns(address)
func (_L2ERC20Gateway *L2ERC20GatewayCaller) GetL1ERC20Address(opts *bind.CallOpts, l2Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2ERC20Gateway.contract.Call(opts, &out, "getL1ERC20Address", l2Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address l1Token) view returns(address)
func (_L2ERC20Gateway *L2ERC20GatewayCaller) GetL2ERC20Address(opts *bind.CallOpts, l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2ERC20Gateway.contract.Call(opts, &out, "getL2ERC20Address", l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ERC20Gateway *L2ERC20GatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC20Gateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2ERC20Gateway *L2ERC20GatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC20Gateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2ERC20Gateway *L2ERC20GatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC20Gateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address l1Token, address l2Token, address from, address to, uint256 amount, bytes data) payable returns()
func (_L2ERC20Gateway *L2ERC20GatewayTransactor) FinalizeDepositERC20(opts *bind.TransactOpts, l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _L2ERC20Gateway.contract.Transact(opts, "finalizeDepositERC20", l1Token, l2Token, from, to, amount, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2ERC20Gateway *L2ERC20GatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ERC20Gateway.contract.Transact(opts, "renounceOwnership")
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2ERC20Gateway *L2ERC20GatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2ERC20Gateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ERC20Gateway *L2ERC20GatewayTransactor) WithdrawERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC20Gateway.contract.Transact(opts, "withdrawERC20", _token, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ERC20Gateway *L2ERC20GatewayTransactor) WithdrawERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC20Gateway.contract.Transact(opts, "withdrawERC200", _token, _to, _amount, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2ERC20Gateway *L2ERC20GatewayTransactor) WithdrawERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC20Gateway.contract.Transact(opts, "withdrawERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// L2ERC20GatewayFinalizeDepositERC20 represents a FinalizeDepositERC20 event raised by the L2ERC20Gateway contract.
type L2ERC20GatewayFinalizeDepositERC20Event struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
}

// L2ERC20GatewayInitialized represents a Initialized event raised by the L2ERC20Gateway contract.
type L2ERC20GatewayInitializedEvent struct {
	Version uint8
}

// L2ERC20GatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L2ERC20Gateway contract.
type L2ERC20GatewayOwnershipTransferredEvent struct {
	PreviousOwner common.Address
	NewOwner      common.Address
}

// L2ERC20GatewayWithdrawERC20 represents a WithdrawERC20 event raised by the L2ERC20Gateway contract.
type L2ERC20GatewayWithdrawERC20Event struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
}
