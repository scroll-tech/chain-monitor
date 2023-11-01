// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package il1scrollmessenger

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

// IL1ScrollMessengerL2MessageProof is an auto generated low-level Go binding around an user-defined struct.
type IL1ScrollMessengerL2MessageProof struct {
	BatchIndex  *big.Int
	MerkleProof []byte
}

// Il1scrollmessengerMetaData contains all meta data concerning the Il1scrollmessenger contract.
var Il1scrollmessengerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"FailedRelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"RelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"SentMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMaxReplayTimes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxReplayTimes\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxReplayTimes\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"dropMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"merkleProof\",\"type\":\"bytes\"}],\"internalType\":\"structIL1ScrollMessenger.L2MessageProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"relayMessageWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"newGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"refundAddress\",\"type\":\"address\"}],\"name\":\"replayMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"refundAddress\",\"type\":\"address\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xDomainMessageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Il1scrollmessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use Il1scrollmessengerMetaData.ABI instead.
var Il1scrollmessengerABI = Il1scrollmessengerMetaData.ABI

// Il1scrollmessenger is an auto generated Go binding around an Ethereum contract.
type Il1scrollmessenger struct {
	Il1scrollmessengerCaller     // Read-only binding to the contract
	Il1scrollmessengerTransactor // Write-only binding to the contract
	Il1scrollmessengerFilterer   // Log filterer for contract events
}

// Il1scrollmessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type Il1scrollmessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il1scrollmessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Il1scrollmessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il1scrollmessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Il1scrollmessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il1scrollmessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Il1scrollmessengerSession struct {
	Contract     *Il1scrollmessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Il1scrollmessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Il1scrollmessengerCallerSession struct {
	Contract *Il1scrollmessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// Il1scrollmessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Il1scrollmessengerTransactorSession struct {
	Contract     *Il1scrollmessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// Il1scrollmessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type Il1scrollmessengerRaw struct {
	Contract *Il1scrollmessenger // Generic contract binding to access the raw methods on
}

// Il1scrollmessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Il1scrollmessengerCallerRaw struct {
	Contract *Il1scrollmessengerCaller // Generic read-only contract binding to access the raw methods on
}

// Il1scrollmessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Il1scrollmessengerTransactorRaw struct {
	Contract *Il1scrollmessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIl1scrollmessenger creates a new instance of Il1scrollmessenger, bound to a specific deployed contract.
func NewIl1scrollmessenger(address common.Address, backend bind.ContractBackend) (*Il1scrollmessenger, error) {
	contract, err := bindIl1scrollmessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Il1scrollmessenger{Il1scrollmessengerCaller: Il1scrollmessengerCaller{contract: contract}, Il1scrollmessengerTransactor: Il1scrollmessengerTransactor{contract: contract}, Il1scrollmessengerFilterer: Il1scrollmessengerFilterer{contract: contract}}, nil
}

// NewIl1scrollmessengerCaller creates a new read-only instance of Il1scrollmessenger, bound to a specific deployed contract.
func NewIl1scrollmessengerCaller(address common.Address, caller bind.ContractCaller) (*Il1scrollmessengerCaller, error) {
	contract, err := bindIl1scrollmessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Il1scrollmessengerCaller{contract: contract}, nil
}

// NewIl1scrollmessengerTransactor creates a new write-only instance of Il1scrollmessenger, bound to a specific deployed contract.
func NewIl1scrollmessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*Il1scrollmessengerTransactor, error) {
	contract, err := bindIl1scrollmessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Il1scrollmessengerTransactor{contract: contract}, nil
}

// NewIl1scrollmessengerFilterer creates a new log filterer instance of Il1scrollmessenger, bound to a specific deployed contract.
func NewIl1scrollmessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*Il1scrollmessengerFilterer, error) {
	contract, err := bindIl1scrollmessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Il1scrollmessengerFilterer{contract: contract}, nil
}

// bindIl1scrollmessenger binds a generic wrapper to an already deployed contract.
func bindIl1scrollmessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Il1scrollmessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Il1scrollmessenger *Il1scrollmessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Il1scrollmessenger.Contract.Il1scrollmessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Il1scrollmessenger *Il1scrollmessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Il1scrollmessenger.Contract.Il1scrollmessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Il1scrollmessenger *Il1scrollmessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Il1scrollmessenger.Contract.Il1scrollmessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Il1scrollmessenger *Il1scrollmessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Il1scrollmessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Il1scrollmessenger *Il1scrollmessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Il1scrollmessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Il1scrollmessenger *Il1scrollmessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Il1scrollmessenger.Contract.contract.Transact(opts, method, params...)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_Il1scrollmessenger *Il1scrollmessengerCaller) XDomainMessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Il1scrollmessenger.contract.Call(opts, &out, "xDomainMessageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_Il1scrollmessenger *Il1scrollmessengerSession) XDomainMessageSender() (common.Address, error) {
	return _Il1scrollmessenger.Contract.XDomainMessageSender(&_Il1scrollmessenger.CallOpts)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_Il1scrollmessenger *Il1scrollmessengerCallerSession) XDomainMessageSender() (common.Address, error) {
	return _Il1scrollmessenger.Contract.XDomainMessageSender(&_Il1scrollmessenger.CallOpts)
}

// DropMessage is a paid mutator transaction binding the contract method 0x29907acd.
//
// Solidity: function dropMessage(address from, address to, uint256 value, uint256 messageNonce, bytes message) returns()
func (_Il1scrollmessenger *Il1scrollmessengerTransactor) DropMessage(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, messageNonce *big.Int, message []byte) (*types.Transaction, error) {
	return _Il1scrollmessenger.contract.Transact(opts, "dropMessage", from, to, value, messageNonce, message)
}

// DropMessage is a paid mutator transaction binding the contract method 0x29907acd.
//
// Solidity: function dropMessage(address from, address to, uint256 value, uint256 messageNonce, bytes message) returns()
func (_Il1scrollmessenger *Il1scrollmessengerSession) DropMessage(from common.Address, to common.Address, value *big.Int, messageNonce *big.Int, message []byte) (*types.Transaction, error) {
	return _Il1scrollmessenger.Contract.DropMessage(&_Il1scrollmessenger.TransactOpts, from, to, value, messageNonce, message)
}

// DropMessage is a paid mutator transaction binding the contract method 0x29907acd.
//
// Solidity: function dropMessage(address from, address to, uint256 value, uint256 messageNonce, bytes message) returns()
func (_Il1scrollmessenger *Il1scrollmessengerTransactorSession) DropMessage(from common.Address, to common.Address, value *big.Int, messageNonce *big.Int, message []byte) (*types.Transaction, error) {
	return _Il1scrollmessenger.Contract.DropMessage(&_Il1scrollmessenger.TransactOpts, from, to, value, messageNonce, message)
}

// RelayMessageWithProof is a paid mutator transaction binding the contract method 0xc311b6fc.
//
// Solidity: function relayMessageWithProof(address from, address to, uint256 value, uint256 nonce, bytes message, (uint256,bytes) proof) returns()
func (_Il1scrollmessenger *Il1scrollmessengerTransactor) RelayMessageWithProof(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, nonce *big.Int, message []byte, proof IL1ScrollMessengerL2MessageProof) (*types.Transaction, error) {
	return _Il1scrollmessenger.contract.Transact(opts, "relayMessageWithProof", from, to, value, nonce, message, proof)
}

// RelayMessageWithProof is a paid mutator transaction binding the contract method 0xc311b6fc.
//
// Solidity: function relayMessageWithProof(address from, address to, uint256 value, uint256 nonce, bytes message, (uint256,bytes) proof) returns()
func (_Il1scrollmessenger *Il1scrollmessengerSession) RelayMessageWithProof(from common.Address, to common.Address, value *big.Int, nonce *big.Int, message []byte, proof IL1ScrollMessengerL2MessageProof) (*types.Transaction, error) {
	return _Il1scrollmessenger.Contract.RelayMessageWithProof(&_Il1scrollmessenger.TransactOpts, from, to, value, nonce, message, proof)
}

// RelayMessageWithProof is a paid mutator transaction binding the contract method 0xc311b6fc.
//
// Solidity: function relayMessageWithProof(address from, address to, uint256 value, uint256 nonce, bytes message, (uint256,bytes) proof) returns()
func (_Il1scrollmessenger *Il1scrollmessengerTransactorSession) RelayMessageWithProof(from common.Address, to common.Address, value *big.Int, nonce *big.Int, message []byte, proof IL1ScrollMessengerL2MessageProof) (*types.Transaction, error) {
	return _Il1scrollmessenger.Contract.RelayMessageWithProof(&_Il1scrollmessenger.TransactOpts, from, to, value, nonce, message, proof)
}

// ReplayMessage is a paid mutator transaction binding the contract method 0x55004105.
//
// Solidity: function replayMessage(address from, address to, uint256 value, uint256 messageNonce, bytes message, uint32 newGasLimit, address refundAddress) payable returns()
func (_Il1scrollmessenger *Il1scrollmessengerTransactor) ReplayMessage(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, messageNonce *big.Int, message []byte, newGasLimit uint32, refundAddress common.Address) (*types.Transaction, error) {
	return _Il1scrollmessenger.contract.Transact(opts, "replayMessage", from, to, value, messageNonce, message, newGasLimit, refundAddress)
}

// ReplayMessage is a paid mutator transaction binding the contract method 0x55004105.
//
// Solidity: function replayMessage(address from, address to, uint256 value, uint256 messageNonce, bytes message, uint32 newGasLimit, address refundAddress) payable returns()
func (_Il1scrollmessenger *Il1scrollmessengerSession) ReplayMessage(from common.Address, to common.Address, value *big.Int, messageNonce *big.Int, message []byte, newGasLimit uint32, refundAddress common.Address) (*types.Transaction, error) {
	return _Il1scrollmessenger.Contract.ReplayMessage(&_Il1scrollmessenger.TransactOpts, from, to, value, messageNonce, message, newGasLimit, refundAddress)
}

// ReplayMessage is a paid mutator transaction binding the contract method 0x55004105.
//
// Solidity: function replayMessage(address from, address to, uint256 value, uint256 messageNonce, bytes message, uint32 newGasLimit, address refundAddress) payable returns()
func (_Il1scrollmessenger *Il1scrollmessengerTransactorSession) ReplayMessage(from common.Address, to common.Address, value *big.Int, messageNonce *big.Int, message []byte, newGasLimit uint32, refundAddress common.Address) (*types.Transaction, error) {
	return _Il1scrollmessenger.Contract.ReplayMessage(&_Il1scrollmessenger.TransactOpts, from, to, value, messageNonce, message, newGasLimit, refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_Il1scrollmessenger *Il1scrollmessengerTransactor) SendMessage(opts *bind.TransactOpts, target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _Il1scrollmessenger.contract.Transact(opts, "sendMessage", target, value, message, gasLimit, refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_Il1scrollmessenger *Il1scrollmessengerSession) SendMessage(target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _Il1scrollmessenger.Contract.SendMessage(&_Il1scrollmessenger.TransactOpts, target, value, message, gasLimit, refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_Il1scrollmessenger *Il1scrollmessengerTransactorSession) SendMessage(target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _Il1scrollmessenger.Contract.SendMessage(&_Il1scrollmessenger.TransactOpts, target, value, message, gasLimit, refundAddress)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_Il1scrollmessenger *Il1scrollmessengerTransactor) SendMessage0(opts *bind.TransactOpts, target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1scrollmessenger.contract.Transact(opts, "sendMessage0", target, value, message, gasLimit)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_Il1scrollmessenger *Il1scrollmessengerSession) SendMessage0(target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1scrollmessenger.Contract.SendMessage0(&_Il1scrollmessenger.TransactOpts, target, value, message, gasLimit)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_Il1scrollmessenger *Il1scrollmessengerTransactorSession) SendMessage0(target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1scrollmessenger.Contract.SendMessage0(&_Il1scrollmessenger.TransactOpts, target, value, message, gasLimit)
}

// Il1scrollmessengerFailedRelayedMessageIterator is returned from FilterFailedRelayedMessage and is used to iterate over the raw logs and unpacked data for FailedRelayedMessage events raised by the Il1scrollmessenger contract.
type Il1scrollmessengerFailedRelayedMessageIterator struct {
	Event *Il1scrollmessengerFailedRelayedMessage // Event containing the contract specifics and raw log

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
func (it *Il1scrollmessengerFailedRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il1scrollmessengerFailedRelayedMessage)
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
		it.Event = new(Il1scrollmessengerFailedRelayedMessage)
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
func (it *Il1scrollmessengerFailedRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il1scrollmessengerFailedRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il1scrollmessengerFailedRelayedMessage represents a FailedRelayedMessage event raised by the Il1scrollmessenger contract.
type Il1scrollmessengerFailedRelayedMessage struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFailedRelayedMessage is a free log retrieval operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_Il1scrollmessenger *Il1scrollmessengerFilterer) FilterFailedRelayedMessage(opts *bind.FilterOpts, messageHash [][32]byte) (*Il1scrollmessengerFailedRelayedMessageIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _Il1scrollmessenger.contract.FilterLogs(opts, "FailedRelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &Il1scrollmessengerFailedRelayedMessageIterator{contract: _Il1scrollmessenger.contract, event: "FailedRelayedMessage", logs: logs, sub: sub}, nil
}

// WatchFailedRelayedMessage is a free log subscription operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_Il1scrollmessenger *Il1scrollmessengerFilterer) WatchFailedRelayedMessage(opts *bind.WatchOpts, sink chan<- *Il1scrollmessengerFailedRelayedMessage, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _Il1scrollmessenger.contract.WatchLogs(opts, "FailedRelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il1scrollmessengerFailedRelayedMessage)
				if err := _Il1scrollmessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
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
func (_Il1scrollmessenger *Il1scrollmessengerFilterer) ParseFailedRelayedMessage(log types.Log) (*Il1scrollmessengerFailedRelayedMessage, error) {
	event := new(Il1scrollmessengerFailedRelayedMessage)
	if err := _Il1scrollmessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il1scrollmessengerRelayedMessageIterator is returned from FilterRelayedMessage and is used to iterate over the raw logs and unpacked data for RelayedMessage events raised by the Il1scrollmessenger contract.
type Il1scrollmessengerRelayedMessageIterator struct {
	Event *Il1scrollmessengerRelayedMessage // Event containing the contract specifics and raw log

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
func (it *Il1scrollmessengerRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il1scrollmessengerRelayedMessage)
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
		it.Event = new(Il1scrollmessengerRelayedMessage)
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
func (it *Il1scrollmessengerRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il1scrollmessengerRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il1scrollmessengerRelayedMessage represents a RelayedMessage event raised by the Il1scrollmessenger contract.
type Il1scrollmessengerRelayedMessage struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRelayedMessage is a free log retrieval operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_Il1scrollmessenger *Il1scrollmessengerFilterer) FilterRelayedMessage(opts *bind.FilterOpts, messageHash [][32]byte) (*Il1scrollmessengerRelayedMessageIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _Il1scrollmessenger.contract.FilterLogs(opts, "RelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &Il1scrollmessengerRelayedMessageIterator{contract: _Il1scrollmessenger.contract, event: "RelayedMessage", logs: logs, sub: sub}, nil
}

// WatchRelayedMessage is a free log subscription operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_Il1scrollmessenger *Il1scrollmessengerFilterer) WatchRelayedMessage(opts *bind.WatchOpts, sink chan<- *Il1scrollmessengerRelayedMessage, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _Il1scrollmessenger.contract.WatchLogs(opts, "RelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il1scrollmessengerRelayedMessage)
				if err := _Il1scrollmessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
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
func (_Il1scrollmessenger *Il1scrollmessengerFilterer) ParseRelayedMessage(log types.Log) (*Il1scrollmessengerRelayedMessage, error) {
	event := new(Il1scrollmessengerRelayedMessage)
	if err := _Il1scrollmessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il1scrollmessengerSentMessageIterator is returned from FilterSentMessage and is used to iterate over the raw logs and unpacked data for SentMessage events raised by the Il1scrollmessenger contract.
type Il1scrollmessengerSentMessageIterator struct {
	Event *Il1scrollmessengerSentMessage // Event containing the contract specifics and raw log

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
func (it *Il1scrollmessengerSentMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il1scrollmessengerSentMessage)
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
		it.Event = new(Il1scrollmessengerSentMessage)
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
func (it *Il1scrollmessengerSentMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il1scrollmessengerSentMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il1scrollmessengerSentMessage represents a SentMessage event raised by the Il1scrollmessenger contract.
type Il1scrollmessengerSentMessage struct {
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
func (_Il1scrollmessenger *Il1scrollmessengerFilterer) FilterSentMessage(opts *bind.FilterOpts, sender []common.Address, target []common.Address) (*Il1scrollmessengerSentMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Il1scrollmessenger.contract.FilterLogs(opts, "SentMessage", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &Il1scrollmessengerSentMessageIterator{contract: _Il1scrollmessenger.contract, event: "SentMessage", logs: logs, sub: sub}, nil
}

// WatchSentMessage is a free log subscription operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_Il1scrollmessenger *Il1scrollmessengerFilterer) WatchSentMessage(opts *bind.WatchOpts, sink chan<- *Il1scrollmessengerSentMessage, sender []common.Address, target []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Il1scrollmessenger.contract.WatchLogs(opts, "SentMessage", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il1scrollmessengerSentMessage)
				if err := _Il1scrollmessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
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
func (_Il1scrollmessenger *Il1scrollmessengerFilterer) ParseSentMessage(log types.Log) (*Il1scrollmessengerSentMessage, error) {
	event := new(Il1scrollmessengerSentMessage)
	if err := _Il1scrollmessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il1scrollmessengerUpdateMaxReplayTimesIterator is returned from FilterUpdateMaxReplayTimes and is used to iterate over the raw logs and unpacked data for UpdateMaxReplayTimes events raised by the Il1scrollmessenger contract.
type Il1scrollmessengerUpdateMaxReplayTimesIterator struct {
	Event *Il1scrollmessengerUpdateMaxReplayTimes // Event containing the contract specifics and raw log

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
func (it *Il1scrollmessengerUpdateMaxReplayTimesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il1scrollmessengerUpdateMaxReplayTimes)
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
		it.Event = new(Il1scrollmessengerUpdateMaxReplayTimes)
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
func (it *Il1scrollmessengerUpdateMaxReplayTimesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il1scrollmessengerUpdateMaxReplayTimesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il1scrollmessengerUpdateMaxReplayTimes represents a UpdateMaxReplayTimes event raised by the Il1scrollmessenger contract.
type Il1scrollmessengerUpdateMaxReplayTimes struct {
	OldMaxReplayTimes *big.Int
	NewMaxReplayTimes *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxReplayTimes is a free log retrieval operation binding the contract event 0xd700562df02eb66951f6f5275df7ebd7c0ec58b3422915789b3b1877aab2e52b.
//
// Solidity: event UpdateMaxReplayTimes(uint256 oldMaxReplayTimes, uint256 newMaxReplayTimes)
func (_Il1scrollmessenger *Il1scrollmessengerFilterer) FilterUpdateMaxReplayTimes(opts *bind.FilterOpts) (*Il1scrollmessengerUpdateMaxReplayTimesIterator, error) {

	logs, sub, err := _Il1scrollmessenger.contract.FilterLogs(opts, "UpdateMaxReplayTimes")
	if err != nil {
		return nil, err
	}
	return &Il1scrollmessengerUpdateMaxReplayTimesIterator{contract: _Il1scrollmessenger.contract, event: "UpdateMaxReplayTimes", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxReplayTimes is a free log subscription operation binding the contract event 0xd700562df02eb66951f6f5275df7ebd7c0ec58b3422915789b3b1877aab2e52b.
//
// Solidity: event UpdateMaxReplayTimes(uint256 oldMaxReplayTimes, uint256 newMaxReplayTimes)
func (_Il1scrollmessenger *Il1scrollmessengerFilterer) WatchUpdateMaxReplayTimes(opts *bind.WatchOpts, sink chan<- *Il1scrollmessengerUpdateMaxReplayTimes) (event.Subscription, error) {

	logs, sub, err := _Il1scrollmessenger.contract.WatchLogs(opts, "UpdateMaxReplayTimes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il1scrollmessengerUpdateMaxReplayTimes)
				if err := _Il1scrollmessenger.contract.UnpackLog(event, "UpdateMaxReplayTimes", log); err != nil {
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

// ParseUpdateMaxReplayTimes is a log parse operation binding the contract event 0xd700562df02eb66951f6f5275df7ebd7c0ec58b3422915789b3b1877aab2e52b.
//
// Solidity: event UpdateMaxReplayTimes(uint256 oldMaxReplayTimes, uint256 newMaxReplayTimes)
func (_Il1scrollmessenger *Il1scrollmessengerFilterer) ParseUpdateMaxReplayTimes(log types.Log) (*Il1scrollmessengerUpdateMaxReplayTimes, error) {
	event := new(Il1scrollmessengerUpdateMaxReplayTimes)
	if err := _Il1scrollmessenger.contract.UnpackLog(event, "UpdateMaxReplayTimes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
