// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package il2erc20gateway

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
	_ = abi.ConvertType
)

// Il2erc20gatewayMetaData contains all meta data concerning the Il2erc20gateway contract.
var Il2erc20gatewayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeDepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawERC20\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"finalizeDepositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"}],\"name\":\"getL1ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"}],\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20AndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// Il2erc20gatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use Il2erc20gatewayMetaData.ABI instead.
var Il2erc20gatewayABI = Il2erc20gatewayMetaData.ABI

// Il2erc20gateway is an auto generated Go binding around an Ethereum contract.
type Il2erc20gateway struct {
	Il2erc20gatewayCaller     // Read-only binding to the contract
	Il2erc20gatewayTransactor // Write-only binding to the contract
	Il2erc20gatewayFilterer   // Log filterer for contract events
}

// Il2erc20gatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type Il2erc20gatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il2erc20gatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Il2erc20gatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il2erc20gatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Il2erc20gatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il2erc20gatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Il2erc20gatewaySession struct {
	Contract     *Il2erc20gateway  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Il2erc20gatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Il2erc20gatewayCallerSession struct {
	Contract *Il2erc20gatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// Il2erc20gatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Il2erc20gatewayTransactorSession struct {
	Contract     *Il2erc20gatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Il2erc20gatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type Il2erc20gatewayRaw struct {
	Contract *Il2erc20gateway // Generic contract binding to access the raw methods on
}

// Il2erc20gatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Il2erc20gatewayCallerRaw struct {
	Contract *Il2erc20gatewayCaller // Generic read-only contract binding to access the raw methods on
}

// Il2erc20gatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Il2erc20gatewayTransactorRaw struct {
	Contract *Il2erc20gatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIl2erc20gateway creates a new instance of Il2erc20gateway, bound to a specific deployed contract.
func NewIl2erc20gateway(address common.Address, backend bind.ContractBackend) (*Il2erc20gateway, error) {
	contract, err := bindIl2erc20gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Il2erc20gateway{Il2erc20gatewayCaller: Il2erc20gatewayCaller{contract: contract}, Il2erc20gatewayTransactor: Il2erc20gatewayTransactor{contract: contract}, Il2erc20gatewayFilterer: Il2erc20gatewayFilterer{contract: contract}}, nil
}

// NewIl2erc20gatewayCaller creates a new read-only instance of Il2erc20gateway, bound to a specific deployed contract.
func NewIl2erc20gatewayCaller(address common.Address, caller bind.ContractCaller) (*Il2erc20gatewayCaller, error) {
	contract, err := bindIl2erc20gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Il2erc20gatewayCaller{contract: contract}, nil
}

// NewIl2erc20gatewayTransactor creates a new write-only instance of Il2erc20gateway, bound to a specific deployed contract.
func NewIl2erc20gatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*Il2erc20gatewayTransactor, error) {
	contract, err := bindIl2erc20gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Il2erc20gatewayTransactor{contract: contract}, nil
}

// NewIl2erc20gatewayFilterer creates a new log filterer instance of Il2erc20gateway, bound to a specific deployed contract.
func NewIl2erc20gatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*Il2erc20gatewayFilterer, error) {
	contract, err := bindIl2erc20gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Il2erc20gatewayFilterer{contract: contract}, nil
}

// bindIl2erc20gateway binds a generic wrapper to an already deployed contract.
func bindIl2erc20gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Il2erc20gatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Il2erc20gateway *Il2erc20gatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Il2erc20gateway.Contract.Il2erc20gatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Il2erc20gateway *Il2erc20gatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Il2erc20gateway.Contract.Il2erc20gatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Il2erc20gateway *Il2erc20gatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Il2erc20gateway.Contract.Il2erc20gatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Il2erc20gateway *Il2erc20gatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Il2erc20gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Il2erc20gateway *Il2erc20gatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Il2erc20gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Il2erc20gateway *Il2erc20gatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Il2erc20gateway.Contract.contract.Transact(opts, method, params...)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address l2Token) view returns(address)
func (_Il2erc20gateway *Il2erc20gatewayCaller) GetL1ERC20Address(opts *bind.CallOpts, l2Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _Il2erc20gateway.contract.Call(opts, &out, "getL1ERC20Address", l2Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address l2Token) view returns(address)
func (_Il2erc20gateway *Il2erc20gatewaySession) GetL1ERC20Address(l2Token common.Address) (common.Address, error) {
	return _Il2erc20gateway.Contract.GetL1ERC20Address(&_Il2erc20gateway.CallOpts, l2Token)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address l2Token) view returns(address)
func (_Il2erc20gateway *Il2erc20gatewayCallerSession) GetL1ERC20Address(l2Token common.Address) (common.Address, error) {
	return _Il2erc20gateway.Contract.GetL1ERC20Address(&_Il2erc20gateway.CallOpts, l2Token)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address l1Token) view returns(address)
func (_Il2erc20gateway *Il2erc20gatewayCaller) GetL2ERC20Address(opts *bind.CallOpts, l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _Il2erc20gateway.contract.Call(opts, &out, "getL2ERC20Address", l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address l1Token) view returns(address)
func (_Il2erc20gateway *Il2erc20gatewaySession) GetL2ERC20Address(l1Token common.Address) (common.Address, error) {
	return _Il2erc20gateway.Contract.GetL2ERC20Address(&_Il2erc20gateway.CallOpts, l1Token)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address l1Token) view returns(address)
func (_Il2erc20gateway *Il2erc20gatewayCallerSession) GetL2ERC20Address(l1Token common.Address) (common.Address, error) {
	return _Il2erc20gateway.Contract.GetL2ERC20Address(&_Il2erc20gateway.CallOpts, l1Token)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address l1Token, address l2Token, address from, address to, uint256 amount, bytes data) payable returns()
func (_Il2erc20gateway *Il2erc20gatewayTransactor) FinalizeDepositERC20(opts *bind.TransactOpts, l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Il2erc20gateway.contract.Transact(opts, "finalizeDepositERC20", l1Token, l2Token, from, to, amount, data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address l1Token, address l2Token, address from, address to, uint256 amount, bytes data) payable returns()
func (_Il2erc20gateway *Il2erc20gatewaySession) FinalizeDepositERC20(l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Il2erc20gateway.Contract.FinalizeDepositERC20(&_Il2erc20gateway.TransactOpts, l1Token, l2Token, from, to, amount, data)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address l1Token, address l2Token, address from, address to, uint256 amount, bytes data) payable returns()
func (_Il2erc20gateway *Il2erc20gatewayTransactorSession) FinalizeDepositERC20(l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Il2erc20gateway.Contract.FinalizeDepositERC20(&_Il2erc20gateway.TransactOpts, l1Token, l2Token, from, to, amount, data)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address token, uint256 amount, uint256 gasLimit) payable returns()
func (_Il2erc20gateway *Il2erc20gatewayTransactor) WithdrawERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc20gateway.contract.Transact(opts, "withdrawERC20", token, amount, gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address token, uint256 amount, uint256 gasLimit) payable returns()
func (_Il2erc20gateway *Il2erc20gatewaySession) WithdrawERC20(token common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc20gateway.Contract.WithdrawERC20(&_Il2erc20gateway.TransactOpts, token, amount, gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address token, uint256 amount, uint256 gasLimit) payable returns()
func (_Il2erc20gateway *Il2erc20gatewayTransactorSession) WithdrawERC20(token common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc20gateway.Contract.WithdrawERC20(&_Il2erc20gateway.TransactOpts, token, amount, gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address token, address to, uint256 amount, uint256 gasLimit) payable returns()
func (_Il2erc20gateway *Il2erc20gatewayTransactor) WithdrawERC200(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc20gateway.contract.Transact(opts, "withdrawERC200", token, to, amount, gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address token, address to, uint256 amount, uint256 gasLimit) payable returns()
func (_Il2erc20gateway *Il2erc20gatewaySession) WithdrawERC200(token common.Address, to common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc20gateway.Contract.WithdrawERC200(&_Il2erc20gateway.TransactOpts, token, to, amount, gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address token, address to, uint256 amount, uint256 gasLimit) payable returns()
func (_Il2erc20gateway *Il2erc20gatewayTransactorSession) WithdrawERC200(token common.Address, to common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc20gateway.Contract.WithdrawERC200(&_Il2erc20gateway.TransactOpts, token, to, amount, gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address token, address to, uint256 amount, bytes data, uint256 gasLimit) payable returns()
func (_Il2erc20gateway *Il2erc20gatewayTransactor) WithdrawERC20AndCall(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, data []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc20gateway.contract.Transact(opts, "withdrawERC20AndCall", token, to, amount, data, gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address token, address to, uint256 amount, bytes data, uint256 gasLimit) payable returns()
func (_Il2erc20gateway *Il2erc20gatewaySession) WithdrawERC20AndCall(token common.Address, to common.Address, amount *big.Int, data []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc20gateway.Contract.WithdrawERC20AndCall(&_Il2erc20gateway.TransactOpts, token, to, amount, data, gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address token, address to, uint256 amount, bytes data, uint256 gasLimit) payable returns()
func (_Il2erc20gateway *Il2erc20gatewayTransactorSession) WithdrawERC20AndCall(token common.Address, to common.Address, amount *big.Int, data []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc20gateway.Contract.WithdrawERC20AndCall(&_Il2erc20gateway.TransactOpts, token, to, amount, data, gasLimit)
}

// Il2erc20gatewayFinalizeDepositERC20Iterator is returned from FilterFinalizeDepositERC20 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC20 events raised by the Il2erc20gateway contract.
type Il2erc20gatewayFinalizeDepositERC20Iterator struct {
	Event *Il2erc20gatewayFinalizeDepositERC20 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Il2erc20gatewayFinalizeDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il2erc20gatewayFinalizeDepositERC20)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Il2erc20gatewayFinalizeDepositERC20)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Il2erc20gatewayFinalizeDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il2erc20gatewayFinalizeDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il2erc20gatewayFinalizeDepositERC20 represents a FinalizeDepositERC20 event raised by the Il2erc20gateway contract.
type Il2erc20gatewayFinalizeDepositERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDepositERC20 is a free log retrieval operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_Il2erc20gateway *Il2erc20gatewayFilterer) FilterFinalizeDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*Il2erc20gatewayFinalizeDepositERC20Iterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Il2erc20gateway.contract.FilterLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &Il2erc20gatewayFinalizeDepositERC20Iterator{contract: _Il2erc20gateway.contract, event: "FinalizeDepositERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC20 is a free log subscription operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_Il2erc20gateway *Il2erc20gatewayFilterer) WatchFinalizeDepositERC20(opts *bind.WatchOpts, sink chan<- *Il2erc20gatewayFinalizeDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Il2erc20gateway.contract.WatchLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il2erc20gatewayFinalizeDepositERC20)
				if err := _Il2erc20gateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFinalizeDepositERC20 is a log parse operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_Il2erc20gateway *Il2erc20gatewayFilterer) ParseFinalizeDepositERC20(log types.Log) (*Il2erc20gatewayFinalizeDepositERC20, error) {
	event := new(Il2erc20gatewayFinalizeDepositERC20)
	if err := _Il2erc20gateway.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il2erc20gatewayWithdrawERC20Iterator is returned from FilterWithdrawERC20 and is used to iterate over the raw logs and unpacked data for WithdrawERC20 events raised by the Il2erc20gateway contract.
type Il2erc20gatewayWithdrawERC20Iterator struct {
	Event *Il2erc20gatewayWithdrawERC20 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Il2erc20gatewayWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il2erc20gatewayWithdrawERC20)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Il2erc20gatewayWithdrawERC20)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Il2erc20gatewayWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il2erc20gatewayWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il2erc20gatewayWithdrawERC20 represents a WithdrawERC20 event raised by the Il2erc20gateway contract.
type Il2erc20gatewayWithdrawERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawERC20 is a free log retrieval operation binding the contract event 0xd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e8.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_Il2erc20gateway *Il2erc20gatewayFilterer) FilterWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*Il2erc20gatewayWithdrawERC20Iterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Il2erc20gateway.contract.FilterLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &Il2erc20gatewayWithdrawERC20Iterator{contract: _Il2erc20gateway.contract, event: "WithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20 is a free log subscription operation binding the contract event 0xd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e8.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_Il2erc20gateway *Il2erc20gatewayFilterer) WatchWithdrawERC20(opts *bind.WatchOpts, sink chan<- *Il2erc20gatewayWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Il2erc20gateway.contract.WatchLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il2erc20gatewayWithdrawERC20)
				if err := _Il2erc20gateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawERC20 is a log parse operation binding the contract event 0xd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e8.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_Il2erc20gateway *Il2erc20gatewayFilterer) ParseWithdrawERC20(log types.Log) (*Il2erc20gatewayWithdrawERC20, error) {
	event := new(Il2erc20gatewayWithdrawERC20)
	if err := _Il2erc20gateway.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
