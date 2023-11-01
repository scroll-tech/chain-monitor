// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package il1ethgateway

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

// Il1ethgatewayMetaData contains all meta data concerning the Il1ethgateway contract.
var Il1ethgatewayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DepositETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeWithdrawETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundETH\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositETHAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"finalizeWithdrawETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// Il1ethgatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use Il1ethgatewayMetaData.ABI instead.
var Il1ethgatewayABI = Il1ethgatewayMetaData.ABI

// Il1ethgateway is an auto generated Go binding around an Ethereum contract.
type Il1ethgateway struct {
	Il1ethgatewayCaller     // Read-only binding to the contract
	Il1ethgatewayTransactor // Write-only binding to the contract
	Il1ethgatewayFilterer   // Log filterer for contract events
}

// Il1ethgatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type Il1ethgatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il1ethgatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Il1ethgatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il1ethgatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Il1ethgatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il1ethgatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Il1ethgatewaySession struct {
	Contract     *Il1ethgateway    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Il1ethgatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Il1ethgatewayCallerSession struct {
	Contract *Il1ethgatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// Il1ethgatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Il1ethgatewayTransactorSession struct {
	Contract     *Il1ethgatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Il1ethgatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type Il1ethgatewayRaw struct {
	Contract *Il1ethgateway // Generic contract binding to access the raw methods on
}

// Il1ethgatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Il1ethgatewayCallerRaw struct {
	Contract *Il1ethgatewayCaller // Generic read-only contract binding to access the raw methods on
}

// Il1ethgatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Il1ethgatewayTransactorRaw struct {
	Contract *Il1ethgatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIl1ethgateway creates a new instance of Il1ethgateway, bound to a specific deployed contract.
func NewIl1ethgateway(address common.Address, backend bind.ContractBackend) (*Il1ethgateway, error) {
	contract, err := bindIl1ethgateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Il1ethgateway{Il1ethgatewayCaller: Il1ethgatewayCaller{contract: contract}, Il1ethgatewayTransactor: Il1ethgatewayTransactor{contract: contract}, Il1ethgatewayFilterer: Il1ethgatewayFilterer{contract: contract}}, nil
}

// NewIl1ethgatewayCaller creates a new read-only instance of Il1ethgateway, bound to a specific deployed contract.
func NewIl1ethgatewayCaller(address common.Address, caller bind.ContractCaller) (*Il1ethgatewayCaller, error) {
	contract, err := bindIl1ethgateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Il1ethgatewayCaller{contract: contract}, nil
}

// NewIl1ethgatewayTransactor creates a new write-only instance of Il1ethgateway, bound to a specific deployed contract.
func NewIl1ethgatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*Il1ethgatewayTransactor, error) {
	contract, err := bindIl1ethgateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Il1ethgatewayTransactor{contract: contract}, nil
}

// NewIl1ethgatewayFilterer creates a new log filterer instance of Il1ethgateway, bound to a specific deployed contract.
func NewIl1ethgatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*Il1ethgatewayFilterer, error) {
	contract, err := bindIl1ethgateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Il1ethgatewayFilterer{contract: contract}, nil
}

// bindIl1ethgateway binds a generic wrapper to an already deployed contract.
func bindIl1ethgateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Il1ethgatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Il1ethgateway *Il1ethgatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Il1ethgateway.Contract.Il1ethgatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Il1ethgateway *Il1ethgatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Il1ethgateway.Contract.Il1ethgatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Il1ethgateway *Il1ethgatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Il1ethgateway.Contract.Il1ethgatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Il1ethgateway *Il1ethgatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Il1ethgateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Il1ethgateway *Il1ethgatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Il1ethgateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Il1ethgateway *Il1ethgatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Il1ethgateway.Contract.contract.Transact(opts, method, params...)
}

// DepositETH is a paid mutator transaction binding the contract method 0x9f8420b3.
//
// Solidity: function depositETH(uint256 amount, uint256 gasLimit) payable returns()
func (_Il1ethgateway *Il1ethgatewayTransactor) DepositETH(opts *bind.TransactOpts, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1ethgateway.contract.Transact(opts, "depositETH", amount, gasLimit)
}

// DepositETH is a paid mutator transaction binding the contract method 0x9f8420b3.
//
// Solidity: function depositETH(uint256 amount, uint256 gasLimit) payable returns()
func (_Il1ethgateway *Il1ethgatewaySession) DepositETH(amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1ethgateway.Contract.DepositETH(&_Il1ethgateway.TransactOpts, amount, gasLimit)
}

// DepositETH is a paid mutator transaction binding the contract method 0x9f8420b3.
//
// Solidity: function depositETH(uint256 amount, uint256 gasLimit) payable returns()
func (_Il1ethgateway *Il1ethgatewayTransactorSession) DepositETH(amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1ethgateway.Contract.DepositETH(&_Il1ethgateway.TransactOpts, amount, gasLimit)
}

// DepositETH0 is a paid mutator transaction binding the contract method 0xce0b63ce.
//
// Solidity: function depositETH(address to, uint256 amount, uint256 gasLimit) payable returns()
func (_Il1ethgateway *Il1ethgatewayTransactor) DepositETH0(opts *bind.TransactOpts, to common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1ethgateway.contract.Transact(opts, "depositETH0", to, amount, gasLimit)
}

// DepositETH0 is a paid mutator transaction binding the contract method 0xce0b63ce.
//
// Solidity: function depositETH(address to, uint256 amount, uint256 gasLimit) payable returns()
func (_Il1ethgateway *Il1ethgatewaySession) DepositETH0(to common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1ethgateway.Contract.DepositETH0(&_Il1ethgateway.TransactOpts, to, amount, gasLimit)
}

// DepositETH0 is a paid mutator transaction binding the contract method 0xce0b63ce.
//
// Solidity: function depositETH(address to, uint256 amount, uint256 gasLimit) payable returns()
func (_Il1ethgateway *Il1ethgatewayTransactorSession) DepositETH0(to common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1ethgateway.Contract.DepositETH0(&_Il1ethgateway.TransactOpts, to, amount, gasLimit)
}

// DepositETHAndCall is a paid mutator transaction binding the contract method 0xaac476f8.
//
// Solidity: function depositETHAndCall(address to, uint256 amount, bytes data, uint256 gasLimit) payable returns()
func (_Il1ethgateway *Il1ethgatewayTransactor) DepositETHAndCall(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1ethgateway.contract.Transact(opts, "depositETHAndCall", to, amount, data, gasLimit)
}

// DepositETHAndCall is a paid mutator transaction binding the contract method 0xaac476f8.
//
// Solidity: function depositETHAndCall(address to, uint256 amount, bytes data, uint256 gasLimit) payable returns()
func (_Il1ethgateway *Il1ethgatewaySession) DepositETHAndCall(to common.Address, amount *big.Int, data []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1ethgateway.Contract.DepositETHAndCall(&_Il1ethgateway.TransactOpts, to, amount, data, gasLimit)
}

// DepositETHAndCall is a paid mutator transaction binding the contract method 0xaac476f8.
//
// Solidity: function depositETHAndCall(address to, uint256 amount, bytes data, uint256 gasLimit) payable returns()
func (_Il1ethgateway *Il1ethgatewayTransactorSession) DepositETHAndCall(to common.Address, amount *big.Int, data []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1ethgateway.Contract.DepositETHAndCall(&_Il1ethgateway.TransactOpts, to, amount, data, gasLimit)
}

// FinalizeWithdrawETH is a paid mutator transaction binding the contract method 0x8eaac8a3.
//
// Solidity: function finalizeWithdrawETH(address from, address to, uint256 amount, bytes data) payable returns()
func (_Il1ethgateway *Il1ethgatewayTransactor) FinalizeWithdrawETH(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Il1ethgateway.contract.Transact(opts, "finalizeWithdrawETH", from, to, amount, data)
}

// FinalizeWithdrawETH is a paid mutator transaction binding the contract method 0x8eaac8a3.
//
// Solidity: function finalizeWithdrawETH(address from, address to, uint256 amount, bytes data) payable returns()
func (_Il1ethgateway *Il1ethgatewaySession) FinalizeWithdrawETH(from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Il1ethgateway.Contract.FinalizeWithdrawETH(&_Il1ethgateway.TransactOpts, from, to, amount, data)
}

// FinalizeWithdrawETH is a paid mutator transaction binding the contract method 0x8eaac8a3.
//
// Solidity: function finalizeWithdrawETH(address from, address to, uint256 amount, bytes data) payable returns()
func (_Il1ethgateway *Il1ethgatewayTransactorSession) FinalizeWithdrawETH(from common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Il1ethgateway.Contract.FinalizeWithdrawETH(&_Il1ethgateway.TransactOpts, from, to, amount, data)
}

// Il1ethgatewayDepositETHIterator is returned from FilterDepositETH and is used to iterate over the raw logs and unpacked data for DepositETH events raised by the Il1ethgateway contract.
type Il1ethgatewayDepositETHIterator struct {
	Event *Il1ethgatewayDepositETH // Event containing the contract specifics and raw log

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
func (it *Il1ethgatewayDepositETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il1ethgatewayDepositETH)
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
		it.Event = new(Il1ethgatewayDepositETH)
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
func (it *Il1ethgatewayDepositETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il1ethgatewayDepositETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il1ethgatewayDepositETH represents a DepositETH event raised by the Il1ethgateway contract.
type Il1ethgatewayDepositETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositETH is a free log retrieval operation binding the contract event 0x6670de856ec8bf5cb2b7e957c5dc24759716056f79d97ea5e7c939ca0ba5a675.
//
// Solidity: event DepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_Il1ethgateway *Il1ethgatewayFilterer) FilterDepositETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Il1ethgatewayDepositETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Il1ethgateway.contract.FilterLogs(opts, "DepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Il1ethgatewayDepositETHIterator{contract: _Il1ethgateway.contract, event: "DepositETH", logs: logs, sub: sub}, nil
}

// WatchDepositETH is a free log subscription operation binding the contract event 0x6670de856ec8bf5cb2b7e957c5dc24759716056f79d97ea5e7c939ca0ba5a675.
//
// Solidity: event DepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_Il1ethgateway *Il1ethgatewayFilterer) WatchDepositETH(opts *bind.WatchOpts, sink chan<- *Il1ethgatewayDepositETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Il1ethgateway.contract.WatchLogs(opts, "DepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il1ethgatewayDepositETH)
				if err := _Il1ethgateway.contract.UnpackLog(event, "DepositETH", log); err != nil {
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

// ParseDepositETH is a log parse operation binding the contract event 0x6670de856ec8bf5cb2b7e957c5dc24759716056f79d97ea5e7c939ca0ba5a675.
//
// Solidity: event DepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_Il1ethgateway *Il1ethgatewayFilterer) ParseDepositETH(log types.Log) (*Il1ethgatewayDepositETH, error) {
	event := new(Il1ethgatewayDepositETH)
	if err := _Il1ethgateway.contract.UnpackLog(event, "DepositETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il1ethgatewayFinalizeWithdrawETHIterator is returned from FilterFinalizeWithdrawETH and is used to iterate over the raw logs and unpacked data for FinalizeWithdrawETH events raised by the Il1ethgateway contract.
type Il1ethgatewayFinalizeWithdrawETHIterator struct {
	Event *Il1ethgatewayFinalizeWithdrawETH // Event containing the contract specifics and raw log

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
func (it *Il1ethgatewayFinalizeWithdrawETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il1ethgatewayFinalizeWithdrawETH)
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
		it.Event = new(Il1ethgatewayFinalizeWithdrawETH)
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
func (it *Il1ethgatewayFinalizeWithdrawETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il1ethgatewayFinalizeWithdrawETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il1ethgatewayFinalizeWithdrawETH represents a FinalizeWithdrawETH event raised by the Il1ethgateway contract.
type Il1ethgatewayFinalizeWithdrawETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFinalizeWithdrawETH is a free log retrieval operation binding the contract event 0x96db5d1cee1dd2760826bb56fabd9c9f6e978083e0a8b88559c741a29e9746e7.
//
// Solidity: event FinalizeWithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_Il1ethgateway *Il1ethgatewayFilterer) FilterFinalizeWithdrawETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Il1ethgatewayFinalizeWithdrawETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Il1ethgateway.contract.FilterLogs(opts, "FinalizeWithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Il1ethgatewayFinalizeWithdrawETHIterator{contract: _Il1ethgateway.contract, event: "FinalizeWithdrawETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeWithdrawETH is a free log subscription operation binding the contract event 0x96db5d1cee1dd2760826bb56fabd9c9f6e978083e0a8b88559c741a29e9746e7.
//
// Solidity: event FinalizeWithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_Il1ethgateway *Il1ethgatewayFilterer) WatchFinalizeWithdrawETH(opts *bind.WatchOpts, sink chan<- *Il1ethgatewayFinalizeWithdrawETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Il1ethgateway.contract.WatchLogs(opts, "FinalizeWithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il1ethgatewayFinalizeWithdrawETH)
				if err := _Il1ethgateway.contract.UnpackLog(event, "FinalizeWithdrawETH", log); err != nil {
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

// ParseFinalizeWithdrawETH is a log parse operation binding the contract event 0x96db5d1cee1dd2760826bb56fabd9c9f6e978083e0a8b88559c741a29e9746e7.
//
// Solidity: event FinalizeWithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_Il1ethgateway *Il1ethgatewayFilterer) ParseFinalizeWithdrawETH(log types.Log) (*Il1ethgatewayFinalizeWithdrawETH, error) {
	event := new(Il1ethgatewayFinalizeWithdrawETH)
	if err := _Il1ethgateway.contract.UnpackLog(event, "FinalizeWithdrawETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il1ethgatewayRefundETHIterator is returned from FilterRefundETH and is used to iterate over the raw logs and unpacked data for RefundETH events raised by the Il1ethgateway contract.
type Il1ethgatewayRefundETHIterator struct {
	Event *Il1ethgatewayRefundETH // Event containing the contract specifics and raw log

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
func (it *Il1ethgatewayRefundETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il1ethgatewayRefundETH)
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
		it.Event = new(Il1ethgatewayRefundETH)
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
func (it *Il1ethgatewayRefundETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il1ethgatewayRefundETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il1ethgatewayRefundETH represents a RefundETH event raised by the Il1ethgateway contract.
type Il1ethgatewayRefundETH struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundETH is a free log retrieval operation binding the contract event 0x289360176646a5f99cb4b6300628426dca46b723f40db3c04449d6ed1745a0e7.
//
// Solidity: event RefundETH(address indexed recipient, uint256 amount)
func (_Il1ethgateway *Il1ethgatewayFilterer) FilterRefundETH(opts *bind.FilterOpts, recipient []common.Address) (*Il1ethgatewayRefundETHIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Il1ethgateway.contract.FilterLogs(opts, "RefundETH", recipientRule)
	if err != nil {
		return nil, err
	}
	return &Il1ethgatewayRefundETHIterator{contract: _Il1ethgateway.contract, event: "RefundETH", logs: logs, sub: sub}, nil
}

// WatchRefundETH is a free log subscription operation binding the contract event 0x289360176646a5f99cb4b6300628426dca46b723f40db3c04449d6ed1745a0e7.
//
// Solidity: event RefundETH(address indexed recipient, uint256 amount)
func (_Il1ethgateway *Il1ethgatewayFilterer) WatchRefundETH(opts *bind.WatchOpts, sink chan<- *Il1ethgatewayRefundETH, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Il1ethgateway.contract.WatchLogs(opts, "RefundETH", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il1ethgatewayRefundETH)
				if err := _Il1ethgateway.contract.UnpackLog(event, "RefundETH", log); err != nil {
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

// ParseRefundETH is a log parse operation binding the contract event 0x289360176646a5f99cb4b6300628426dca46b723f40db3c04449d6ed1745a0e7.
//
// Solidity: event RefundETH(address indexed recipient, uint256 amount)
func (_Il1ethgateway *Il1ethgatewayFilterer) ParseRefundETH(log types.Log) (*Il1ethgatewayRefundETH, error) {
	event := new(Il1ethgatewayRefundETH)
	if err := _Il1ethgateway.contract.UnpackLog(event, "RefundETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
