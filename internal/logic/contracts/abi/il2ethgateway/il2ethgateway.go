// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package il2ethgateway

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

// Il2ethgatewayMetaData contains all meta data concerning the Il2ethgateway contract.
var Il2ethgatewayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeDepositETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawETH\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeDepositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawETHAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// Il2ethgatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use Il2ethgatewayMetaData.ABI instead.
var Il2ethgatewayABI = Il2ethgatewayMetaData.ABI

// Il2ethgateway is an auto generated Go binding around an Ethereum contract.
type Il2ethgateway struct {
	Il2ethgatewayCaller     // Read-only binding to the contract
	Il2ethgatewayTransactor // Write-only binding to the contract
	Il2ethgatewayFilterer   // Log filterer for contract events
}

// Il2ethgatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type Il2ethgatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il2ethgatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Il2ethgatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il2ethgatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Il2ethgatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il2ethgatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Il2ethgatewaySession struct {
	Contract     *Il2ethgateway    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Il2ethgatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Il2ethgatewayCallerSession struct {
	Contract *Il2ethgatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// Il2ethgatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Il2ethgatewayTransactorSession struct {
	Contract     *Il2ethgatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Il2ethgatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type Il2ethgatewayRaw struct {
	Contract *Il2ethgateway // Generic contract binding to access the raw methods on
}

// Il2ethgatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Il2ethgatewayCallerRaw struct {
	Contract *Il2ethgatewayCaller // Generic read-only contract binding to access the raw methods on
}

// Il2ethgatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Il2ethgatewayTransactorRaw struct {
	Contract *Il2ethgatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIl2ethgateway creates a new instance of Il2ethgateway, bound to a specific deployed contract.
func NewIl2ethgateway(address common.Address, backend bind.ContractBackend) (*Il2ethgateway, error) {
	contract, err := bindIl2ethgateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Il2ethgateway{Il2ethgatewayCaller: Il2ethgatewayCaller{contract: contract}, Il2ethgatewayTransactor: Il2ethgatewayTransactor{contract: contract}, Il2ethgatewayFilterer: Il2ethgatewayFilterer{contract: contract}}, nil
}

// NewIl2ethgatewayCaller creates a new read-only instance of Il2ethgateway, bound to a specific deployed contract.
func NewIl2ethgatewayCaller(address common.Address, caller bind.ContractCaller) (*Il2ethgatewayCaller, error) {
	contract, err := bindIl2ethgateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Il2ethgatewayCaller{contract: contract}, nil
}

// NewIl2ethgatewayTransactor creates a new write-only instance of Il2ethgateway, bound to a specific deployed contract.
func NewIl2ethgatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*Il2ethgatewayTransactor, error) {
	contract, err := bindIl2ethgateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Il2ethgatewayTransactor{contract: contract}, nil
}

// NewIl2ethgatewayFilterer creates a new log filterer instance of Il2ethgateway, bound to a specific deployed contract.
func NewIl2ethgatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*Il2ethgatewayFilterer, error) {
	contract, err := bindIl2ethgateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Il2ethgatewayFilterer{contract: contract}, nil
}

// bindIl2ethgateway binds a generic wrapper to an already deployed contract.
func bindIl2ethgateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Il2ethgatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Il2ethgateway *Il2ethgatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Il2ethgateway.Contract.Il2ethgatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Il2ethgateway *Il2ethgatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Il2ethgateway.Contract.Il2ethgatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Il2ethgateway *Il2ethgatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Il2ethgateway.Contract.Il2ethgatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Il2ethgateway *Il2ethgatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Il2ethgateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Il2ethgateway *Il2ethgatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Il2ethgateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Il2ethgateway *Il2ethgatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Il2ethgateway.Contract.contract.Transact(opts, method, params...)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_Il2ethgateway *Il2ethgatewayTransactor) FinalizeDepositETH(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _Il2ethgateway.contract.Transact(opts, "finalizeDepositETH", _from, _to, _amount, _data)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_Il2ethgateway *Il2ethgatewaySession) FinalizeDepositETH(_from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _Il2ethgateway.Contract.FinalizeDepositETH(&_Il2ethgateway.TransactOpts, _from, _to, _amount, _data)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_Il2ethgateway *Il2ethgatewayTransactorSession) FinalizeDepositETH(_from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _Il2ethgateway.Contract.FinalizeDepositETH(&_Il2ethgateway.TransactOpts, _from, _to, _amount, _data)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address to, uint256 amount, uint256 gasLimit) payable returns()
func (_Il2ethgateway *Il2ethgatewayTransactor) WithdrawETH(opts *bind.TransactOpts, to common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2ethgateway.contract.Transact(opts, "withdrawETH", to, amount, gasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address to, uint256 amount, uint256 gasLimit) payable returns()
func (_Il2ethgateway *Il2ethgatewaySession) WithdrawETH(to common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2ethgateway.Contract.WithdrawETH(&_Il2ethgateway.TransactOpts, to, amount, gasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address to, uint256 amount, uint256 gasLimit) payable returns()
func (_Il2ethgateway *Il2ethgatewayTransactorSession) WithdrawETH(to common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2ethgateway.Contract.WithdrawETH(&_Il2ethgateway.TransactOpts, to, amount, gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 amount, uint256 gasLimit) payable returns()
func (_Il2ethgateway *Il2ethgatewayTransactor) WithdrawETH0(opts *bind.TransactOpts, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2ethgateway.contract.Transact(opts, "withdrawETH0", amount, gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 amount, uint256 gasLimit) payable returns()
func (_Il2ethgateway *Il2ethgatewaySession) WithdrawETH0(amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2ethgateway.Contract.WithdrawETH0(&_Il2ethgateway.TransactOpts, amount, gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 amount, uint256 gasLimit) payable returns()
func (_Il2ethgateway *Il2ethgatewayTransactorSession) WithdrawETH0(amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2ethgateway.Contract.WithdrawETH0(&_Il2ethgateway.TransactOpts, amount, gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address to, uint256 amount, bytes data, uint256 gasLimit) payable returns()
func (_Il2ethgateway *Il2ethgatewayTransactor) WithdrawETHAndCall(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2ethgateway.contract.Transact(opts, "withdrawETHAndCall", to, amount, data, gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address to, uint256 amount, bytes data, uint256 gasLimit) payable returns()
func (_Il2ethgateway *Il2ethgatewaySession) WithdrawETHAndCall(to common.Address, amount *big.Int, data []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2ethgateway.Contract.WithdrawETHAndCall(&_Il2ethgateway.TransactOpts, to, amount, data, gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address to, uint256 amount, bytes data, uint256 gasLimit) payable returns()
func (_Il2ethgateway *Il2ethgatewayTransactorSession) WithdrawETHAndCall(to common.Address, amount *big.Int, data []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2ethgateway.Contract.WithdrawETHAndCall(&_Il2ethgateway.TransactOpts, to, amount, data, gasLimit)
}

// Il2ethgatewayFinalizeDepositETHIterator is returned from FilterFinalizeDepositETH and is used to iterate over the raw logs and unpacked data for FinalizeDepositETH events raised by the Il2ethgateway contract.
type Il2ethgatewayFinalizeDepositETHIterator struct {
	Event *Il2ethgatewayFinalizeDepositETH // Event containing the contract specifics and raw log

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
func (it *Il2ethgatewayFinalizeDepositETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il2ethgatewayFinalizeDepositETH)
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
		it.Event = new(Il2ethgatewayFinalizeDepositETH)
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
func (it *Il2ethgatewayFinalizeDepositETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il2ethgatewayFinalizeDepositETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il2ethgatewayFinalizeDepositETH represents a FinalizeDepositETH event raised by the Il2ethgateway contract.
type Il2ethgatewayFinalizeDepositETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDepositETH is a free log retrieval operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_Il2ethgateway *Il2ethgatewayFilterer) FilterFinalizeDepositETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Il2ethgatewayFinalizeDepositETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Il2ethgateway.contract.FilterLogs(opts, "FinalizeDepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Il2ethgatewayFinalizeDepositETHIterator{contract: _Il2ethgateway.contract, event: "FinalizeDepositETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositETH is a free log subscription operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_Il2ethgateway *Il2ethgatewayFilterer) WatchFinalizeDepositETH(opts *bind.WatchOpts, sink chan<- *Il2ethgatewayFinalizeDepositETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Il2ethgateway.contract.WatchLogs(opts, "FinalizeDepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il2ethgatewayFinalizeDepositETH)
				if err := _Il2ethgateway.contract.UnpackLog(event, "FinalizeDepositETH", log); err != nil {
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

// ParseFinalizeDepositETH is a log parse operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_Il2ethgateway *Il2ethgatewayFilterer) ParseFinalizeDepositETH(log types.Log) (*Il2ethgatewayFinalizeDepositETH, error) {
	event := new(Il2ethgatewayFinalizeDepositETH)
	if err := _Il2ethgateway.contract.UnpackLog(event, "FinalizeDepositETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il2ethgatewayWithdrawETHIterator is returned from FilterWithdrawETH and is used to iterate over the raw logs and unpacked data for WithdrawETH events raised by the Il2ethgateway contract.
type Il2ethgatewayWithdrawETHIterator struct {
	Event *Il2ethgatewayWithdrawETH // Event containing the contract specifics and raw log

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
func (it *Il2ethgatewayWithdrawETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il2ethgatewayWithdrawETH)
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
		it.Event = new(Il2ethgatewayWithdrawETH)
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
func (it *Il2ethgatewayWithdrawETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il2ethgatewayWithdrawETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il2ethgatewayWithdrawETH represents a WithdrawETH event raised by the Il2ethgateway contract.
type Il2ethgatewayWithdrawETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawETH is a free log retrieval operation binding the contract event 0xd8ed6eaa9a7a8980d7901e911fde6686810b989d3082182d1d3a3df6306ce20e.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_Il2ethgateway *Il2ethgatewayFilterer) FilterWithdrawETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Il2ethgatewayWithdrawETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Il2ethgateway.contract.FilterLogs(opts, "WithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Il2ethgatewayWithdrawETHIterator{contract: _Il2ethgateway.contract, event: "WithdrawETH", logs: logs, sub: sub}, nil
}

// WatchWithdrawETH is a free log subscription operation binding the contract event 0xd8ed6eaa9a7a8980d7901e911fde6686810b989d3082182d1d3a3df6306ce20e.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_Il2ethgateway *Il2ethgatewayFilterer) WatchWithdrawETH(opts *bind.WatchOpts, sink chan<- *Il2ethgatewayWithdrawETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Il2ethgateway.contract.WatchLogs(opts, "WithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il2ethgatewayWithdrawETH)
				if err := _Il2ethgateway.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
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

// ParseWithdrawETH is a log parse operation binding the contract event 0xd8ed6eaa9a7a8980d7901e911fde6686810b989d3082182d1d3a3df6306ce20e.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_Il2ethgateway *Il2ethgatewayFilterer) ParseWithdrawETH(log types.Log) (*Il2ethgatewayWithdrawETH, error) {
	event := new(Il2ethgatewayWithdrawETH)
	if err := _Il2ethgateway.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
