// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package il2scrollmessenger

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

// Il2scrollmessengerMetaData contains all meta data concerning the Il2scrollmessenger contract.
var Il2scrollmessengerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"FailedRelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"RelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"SentMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMaxFailedExecutionTimes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxFailedExecutionTimes\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxFailedExecutionTimes\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"relayMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"refundAddress\",\"type\":\"address\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xDomainMessageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Il2scrollmessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use Il2scrollmessengerMetaData.ABI instead.
var Il2scrollmessengerABI = Il2scrollmessengerMetaData.ABI

// Il2scrollmessenger is an auto generated Go binding around an Ethereum contract.
type Il2scrollmessenger struct {
	Il2scrollmessengerCaller     // Read-only binding to the contract
	Il2scrollmessengerTransactor // Write-only binding to the contract
	Il2scrollmessengerFilterer   // Log filterer for contract events
}

// Il2scrollmessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type Il2scrollmessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il2scrollmessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Il2scrollmessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il2scrollmessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Il2scrollmessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il2scrollmessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Il2scrollmessengerSession struct {
	Contract     *Il2scrollmessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Il2scrollmessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Il2scrollmessengerCallerSession struct {
	Contract *Il2scrollmessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// Il2scrollmessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Il2scrollmessengerTransactorSession struct {
	Contract     *Il2scrollmessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// Il2scrollmessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type Il2scrollmessengerRaw struct {
	Contract *Il2scrollmessenger // Generic contract binding to access the raw methods on
}

// Il2scrollmessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Il2scrollmessengerCallerRaw struct {
	Contract *Il2scrollmessengerCaller // Generic read-only contract binding to access the raw methods on
}

// Il2scrollmessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Il2scrollmessengerTransactorRaw struct {
	Contract *Il2scrollmessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIl2scrollmessenger creates a new instance of Il2scrollmessenger, bound to a specific deployed contract.
func NewIl2scrollmessenger(address common.Address, backend bind.ContractBackend) (*Il2scrollmessenger, error) {
	contract, err := bindIl2scrollmessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Il2scrollmessenger{Il2scrollmessengerCaller: Il2scrollmessengerCaller{contract: contract}, Il2scrollmessengerTransactor: Il2scrollmessengerTransactor{contract: contract}, Il2scrollmessengerFilterer: Il2scrollmessengerFilterer{contract: contract}}, nil
}

// NewIl2scrollmessengerCaller creates a new read-only instance of Il2scrollmessenger, bound to a specific deployed contract.
func NewIl2scrollmessengerCaller(address common.Address, caller bind.ContractCaller) (*Il2scrollmessengerCaller, error) {
	contract, err := bindIl2scrollmessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Il2scrollmessengerCaller{contract: contract}, nil
}

// NewIl2scrollmessengerTransactor creates a new write-only instance of Il2scrollmessenger, bound to a specific deployed contract.
func NewIl2scrollmessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*Il2scrollmessengerTransactor, error) {
	contract, err := bindIl2scrollmessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Il2scrollmessengerTransactor{contract: contract}, nil
}

// NewIl2scrollmessengerFilterer creates a new log filterer instance of Il2scrollmessenger, bound to a specific deployed contract.
func NewIl2scrollmessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*Il2scrollmessengerFilterer, error) {
	contract, err := bindIl2scrollmessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Il2scrollmessengerFilterer{contract: contract}, nil
}

// bindIl2scrollmessenger binds a generic wrapper to an already deployed contract.
func bindIl2scrollmessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Il2scrollmessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Il2scrollmessenger *Il2scrollmessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Il2scrollmessenger.Contract.Il2scrollmessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Il2scrollmessenger *Il2scrollmessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Il2scrollmessenger.Contract.Il2scrollmessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Il2scrollmessenger *Il2scrollmessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Il2scrollmessenger.Contract.Il2scrollmessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Il2scrollmessenger *Il2scrollmessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Il2scrollmessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Il2scrollmessenger *Il2scrollmessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Il2scrollmessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Il2scrollmessenger *Il2scrollmessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Il2scrollmessenger.Contract.contract.Transact(opts, method, params...)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_Il2scrollmessenger *Il2scrollmessengerCaller) XDomainMessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Il2scrollmessenger.contract.Call(opts, &out, "xDomainMessageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_Il2scrollmessenger *Il2scrollmessengerSession) XDomainMessageSender() (common.Address, error) {
	return _Il2scrollmessenger.Contract.XDomainMessageSender(&_Il2scrollmessenger.CallOpts)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_Il2scrollmessenger *Il2scrollmessengerCallerSession) XDomainMessageSender() (common.Address, error) {
	return _Il2scrollmessenger.Contract.XDomainMessageSender(&_Il2scrollmessenger.CallOpts)
}

// RelayMessage is a paid mutator transaction binding the contract method 0x8ef1332e.
//
// Solidity: function relayMessage(address from, address to, uint256 value, uint256 nonce, bytes message) returns()
func (_Il2scrollmessenger *Il2scrollmessengerTransactor) RelayMessage(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, nonce *big.Int, message []byte) (*types.Transaction, error) {
	return _Il2scrollmessenger.contract.Transact(opts, "relayMessage", from, to, value, nonce, message)
}

// RelayMessage is a paid mutator transaction binding the contract method 0x8ef1332e.
//
// Solidity: function relayMessage(address from, address to, uint256 value, uint256 nonce, bytes message) returns()
func (_Il2scrollmessenger *Il2scrollmessengerSession) RelayMessage(from common.Address, to common.Address, value *big.Int, nonce *big.Int, message []byte) (*types.Transaction, error) {
	return _Il2scrollmessenger.Contract.RelayMessage(&_Il2scrollmessenger.TransactOpts, from, to, value, nonce, message)
}

// RelayMessage is a paid mutator transaction binding the contract method 0x8ef1332e.
//
// Solidity: function relayMessage(address from, address to, uint256 value, uint256 nonce, bytes message) returns()
func (_Il2scrollmessenger *Il2scrollmessengerTransactorSession) RelayMessage(from common.Address, to common.Address, value *big.Int, nonce *big.Int, message []byte) (*types.Transaction, error) {
	return _Il2scrollmessenger.Contract.RelayMessage(&_Il2scrollmessenger.TransactOpts, from, to, value, nonce, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_Il2scrollmessenger *Il2scrollmessengerTransactor) SendMessage(opts *bind.TransactOpts, target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _Il2scrollmessenger.contract.Transact(opts, "sendMessage", target, value, message, gasLimit, refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_Il2scrollmessenger *Il2scrollmessengerSession) SendMessage(target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _Il2scrollmessenger.Contract.SendMessage(&_Il2scrollmessenger.TransactOpts, target, value, message, gasLimit, refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_Il2scrollmessenger *Il2scrollmessengerTransactorSession) SendMessage(target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _Il2scrollmessenger.Contract.SendMessage(&_Il2scrollmessenger.TransactOpts, target, value, message, gasLimit, refundAddress)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_Il2scrollmessenger *Il2scrollmessengerTransactor) SendMessage0(opts *bind.TransactOpts, target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2scrollmessenger.contract.Transact(opts, "sendMessage0", target, value, message, gasLimit)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_Il2scrollmessenger *Il2scrollmessengerSession) SendMessage0(target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2scrollmessenger.Contract.SendMessage0(&_Il2scrollmessenger.TransactOpts, target, value, message, gasLimit)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_Il2scrollmessenger *Il2scrollmessengerTransactorSession) SendMessage0(target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2scrollmessenger.Contract.SendMessage0(&_Il2scrollmessenger.TransactOpts, target, value, message, gasLimit)
}

// Il2scrollmessengerFailedRelayedMessageIterator is returned from FilterFailedRelayedMessage and is used to iterate over the raw logs and unpacked data for FailedRelayedMessage events raised by the Il2scrollmessenger contract.
type Il2scrollmessengerFailedRelayedMessageIterator struct {
	Event *Il2scrollmessengerFailedRelayedMessage // Event containing the contract specifics and raw log

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
func (it *Il2scrollmessengerFailedRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il2scrollmessengerFailedRelayedMessage)
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
		it.Event = new(Il2scrollmessengerFailedRelayedMessage)
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
func (it *Il2scrollmessengerFailedRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il2scrollmessengerFailedRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il2scrollmessengerFailedRelayedMessage represents a FailedRelayedMessage event raised by the Il2scrollmessenger contract.
type Il2scrollmessengerFailedRelayedMessage struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFailedRelayedMessage is a free log retrieval operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_Il2scrollmessenger *Il2scrollmessengerFilterer) FilterFailedRelayedMessage(opts *bind.FilterOpts, messageHash [][32]byte) (*Il2scrollmessengerFailedRelayedMessageIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _Il2scrollmessenger.contract.FilterLogs(opts, "FailedRelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &Il2scrollmessengerFailedRelayedMessageIterator{contract: _Il2scrollmessenger.contract, event: "FailedRelayedMessage", logs: logs, sub: sub}, nil
}

// WatchFailedRelayedMessage is a free log subscription operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_Il2scrollmessenger *Il2scrollmessengerFilterer) WatchFailedRelayedMessage(opts *bind.WatchOpts, sink chan<- *Il2scrollmessengerFailedRelayedMessage, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _Il2scrollmessenger.contract.WatchLogs(opts, "FailedRelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il2scrollmessengerFailedRelayedMessage)
				if err := _Il2scrollmessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
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

// ParseFailedRelayedMessage is a log parse operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_Il2scrollmessenger *Il2scrollmessengerFilterer) ParseFailedRelayedMessage(log types.Log) (*Il2scrollmessengerFailedRelayedMessage, error) {
	event := new(Il2scrollmessengerFailedRelayedMessage)
	if err := _Il2scrollmessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il2scrollmessengerRelayedMessageIterator is returned from FilterRelayedMessage and is used to iterate over the raw logs and unpacked data for RelayedMessage events raised by the Il2scrollmessenger contract.
type Il2scrollmessengerRelayedMessageIterator struct {
	Event *Il2scrollmessengerRelayedMessage // Event containing the contract specifics and raw log

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
func (it *Il2scrollmessengerRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il2scrollmessengerRelayedMessage)
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
		it.Event = new(Il2scrollmessengerRelayedMessage)
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
func (it *Il2scrollmessengerRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il2scrollmessengerRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il2scrollmessengerRelayedMessage represents a RelayedMessage event raised by the Il2scrollmessenger contract.
type Il2scrollmessengerRelayedMessage struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRelayedMessage is a free log retrieval operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_Il2scrollmessenger *Il2scrollmessengerFilterer) FilterRelayedMessage(opts *bind.FilterOpts, messageHash [][32]byte) (*Il2scrollmessengerRelayedMessageIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _Il2scrollmessenger.contract.FilterLogs(opts, "RelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &Il2scrollmessengerRelayedMessageIterator{contract: _Il2scrollmessenger.contract, event: "RelayedMessage", logs: logs, sub: sub}, nil
}

// WatchRelayedMessage is a free log subscription operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_Il2scrollmessenger *Il2scrollmessengerFilterer) WatchRelayedMessage(opts *bind.WatchOpts, sink chan<- *Il2scrollmessengerRelayedMessage, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _Il2scrollmessenger.contract.WatchLogs(opts, "RelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il2scrollmessengerRelayedMessage)
				if err := _Il2scrollmessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
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

// ParseRelayedMessage is a log parse operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_Il2scrollmessenger *Il2scrollmessengerFilterer) ParseRelayedMessage(log types.Log) (*Il2scrollmessengerRelayedMessage, error) {
	event := new(Il2scrollmessengerRelayedMessage)
	if err := _Il2scrollmessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il2scrollmessengerSentMessageIterator is returned from FilterSentMessage and is used to iterate over the raw logs and unpacked data for SentMessage events raised by the Il2scrollmessenger contract.
type Il2scrollmessengerSentMessageIterator struct {
	Event *Il2scrollmessengerSentMessage // Event containing the contract specifics and raw log

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
func (it *Il2scrollmessengerSentMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il2scrollmessengerSentMessage)
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
		it.Event = new(Il2scrollmessengerSentMessage)
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
func (it *Il2scrollmessengerSentMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il2scrollmessengerSentMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il2scrollmessengerSentMessage represents a SentMessage event raised by the Il2scrollmessenger contract.
type Il2scrollmessengerSentMessage struct {
	Sender       common.Address
	Target       common.Address
	Value        *big.Int
	MessageNonce *big.Int
	GasLimit     *big.Int
	Message      []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSentMessage is a free log retrieval operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_Il2scrollmessenger *Il2scrollmessengerFilterer) FilterSentMessage(opts *bind.FilterOpts, sender []common.Address, target []common.Address) (*Il2scrollmessengerSentMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Il2scrollmessenger.contract.FilterLogs(opts, "SentMessage", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &Il2scrollmessengerSentMessageIterator{contract: _Il2scrollmessenger.contract, event: "SentMessage", logs: logs, sub: sub}, nil
}

// WatchSentMessage is a free log subscription operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_Il2scrollmessenger *Il2scrollmessengerFilterer) WatchSentMessage(opts *bind.WatchOpts, sink chan<- *Il2scrollmessengerSentMessage, sender []common.Address, target []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Il2scrollmessenger.contract.WatchLogs(opts, "SentMessage", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il2scrollmessengerSentMessage)
				if err := _Il2scrollmessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
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

// ParseSentMessage is a log parse operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_Il2scrollmessenger *Il2scrollmessengerFilterer) ParseSentMessage(log types.Log) (*Il2scrollmessengerSentMessage, error) {
	event := new(Il2scrollmessengerSentMessage)
	if err := _Il2scrollmessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il2scrollmessengerUpdateMaxFailedExecutionTimesIterator is returned from FilterUpdateMaxFailedExecutionTimes and is used to iterate over the raw logs and unpacked data for UpdateMaxFailedExecutionTimes events raised by the Il2scrollmessenger contract.
type Il2scrollmessengerUpdateMaxFailedExecutionTimesIterator struct {
	Event *Il2scrollmessengerUpdateMaxFailedExecutionTimes // Event containing the contract specifics and raw log

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
func (it *Il2scrollmessengerUpdateMaxFailedExecutionTimesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il2scrollmessengerUpdateMaxFailedExecutionTimes)
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
		it.Event = new(Il2scrollmessengerUpdateMaxFailedExecutionTimes)
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
func (it *Il2scrollmessengerUpdateMaxFailedExecutionTimesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il2scrollmessengerUpdateMaxFailedExecutionTimesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il2scrollmessengerUpdateMaxFailedExecutionTimes represents a UpdateMaxFailedExecutionTimes event raised by the Il2scrollmessenger contract.
type Il2scrollmessengerUpdateMaxFailedExecutionTimes struct {
	OldMaxFailedExecutionTimes *big.Int
	NewMaxFailedExecutionTimes *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxFailedExecutionTimes is a free log retrieval operation binding the contract event 0x8a4c22c9b46f23dedd49b843839940ce0c36fa1612073a9bc7dbaeef9ee547ba.
//
// Solidity: event UpdateMaxFailedExecutionTimes(uint256 oldMaxFailedExecutionTimes, uint256 newMaxFailedExecutionTimes)
func (_Il2scrollmessenger *Il2scrollmessengerFilterer) FilterUpdateMaxFailedExecutionTimes(opts *bind.FilterOpts) (*Il2scrollmessengerUpdateMaxFailedExecutionTimesIterator, error) {

	logs, sub, err := _Il2scrollmessenger.contract.FilterLogs(opts, "UpdateMaxFailedExecutionTimes")
	if err != nil {
		return nil, err
	}
	return &Il2scrollmessengerUpdateMaxFailedExecutionTimesIterator{contract: _Il2scrollmessenger.contract, event: "UpdateMaxFailedExecutionTimes", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxFailedExecutionTimes is a free log subscription operation binding the contract event 0x8a4c22c9b46f23dedd49b843839940ce0c36fa1612073a9bc7dbaeef9ee547ba.
//
// Solidity: event UpdateMaxFailedExecutionTimes(uint256 oldMaxFailedExecutionTimes, uint256 newMaxFailedExecutionTimes)
func (_Il2scrollmessenger *Il2scrollmessengerFilterer) WatchUpdateMaxFailedExecutionTimes(opts *bind.WatchOpts, sink chan<- *Il2scrollmessengerUpdateMaxFailedExecutionTimes) (event.Subscription, error) {

	logs, sub, err := _Il2scrollmessenger.contract.WatchLogs(opts, "UpdateMaxFailedExecutionTimes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il2scrollmessengerUpdateMaxFailedExecutionTimes)
				if err := _Il2scrollmessenger.contract.UnpackLog(event, "UpdateMaxFailedExecutionTimes", log); err != nil {
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

// ParseUpdateMaxFailedExecutionTimes is a log parse operation binding the contract event 0x8a4c22c9b46f23dedd49b843839940ce0c36fa1612073a9bc7dbaeef9ee547ba.
//
// Solidity: event UpdateMaxFailedExecutionTimes(uint256 oldMaxFailedExecutionTimes, uint256 newMaxFailedExecutionTimes)
func (_Il2scrollmessenger *Il2scrollmessengerFilterer) ParseUpdateMaxFailedExecutionTimes(log types.Log) (*Il2scrollmessengerUpdateMaxFailedExecutionTimes, error) {
	event := new(Il2scrollmessengerUpdateMaxFailedExecutionTimes)
	if err := _Il2scrollmessenger.contract.UnpackLog(event, "UpdateMaxFailedExecutionTimes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
