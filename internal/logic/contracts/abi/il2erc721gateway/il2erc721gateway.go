// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package il2erc721gateway

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

// Il2erc721gatewayMetaData contains all meta data concerning the Il2erc721gateway contract.
var Il2erc721gatewayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"BatchWithdrawERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"FinalizeBatchDepositERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"FinalizeDepositERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC721\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"batchWithdrawERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"batchWithdrawERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"finalizeBatchDepositERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"finalizeDepositERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// Il2erc721gatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use Il2erc721gatewayMetaData.ABI instead.
var Il2erc721gatewayABI = Il2erc721gatewayMetaData.ABI

// Il2erc721gateway is an auto generated Go binding around an Ethereum contract.
type Il2erc721gateway struct {
	Il2erc721gatewayCaller     // Read-only binding to the contract
	Il2erc721gatewayTransactor // Write-only binding to the contract
	Il2erc721gatewayFilterer   // Log filterer for contract events
}

// Il2erc721gatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type Il2erc721gatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il2erc721gatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Il2erc721gatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il2erc721gatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Il2erc721gatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il2erc721gatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Il2erc721gatewaySession struct {
	Contract     *Il2erc721gateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Il2erc721gatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Il2erc721gatewayCallerSession struct {
	Contract *Il2erc721gatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Il2erc721gatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Il2erc721gatewayTransactorSession struct {
	Contract     *Il2erc721gatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Il2erc721gatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type Il2erc721gatewayRaw struct {
	Contract *Il2erc721gateway // Generic contract binding to access the raw methods on
}

// Il2erc721gatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Il2erc721gatewayCallerRaw struct {
	Contract *Il2erc721gatewayCaller // Generic read-only contract binding to access the raw methods on
}

// Il2erc721gatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Il2erc721gatewayTransactorRaw struct {
	Contract *Il2erc721gatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIl2erc721gateway creates a new instance of Il2erc721gateway, bound to a specific deployed contract.
func NewIl2erc721gateway(address common.Address, backend bind.ContractBackend) (*Il2erc721gateway, error) {
	contract, err := bindIl2erc721gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Il2erc721gateway{Il2erc721gatewayCaller: Il2erc721gatewayCaller{contract: contract}, Il2erc721gatewayTransactor: Il2erc721gatewayTransactor{contract: contract}, Il2erc721gatewayFilterer: Il2erc721gatewayFilterer{contract: contract}}, nil
}

// NewIl2erc721gatewayCaller creates a new read-only instance of Il2erc721gateway, bound to a specific deployed contract.
func NewIl2erc721gatewayCaller(address common.Address, caller bind.ContractCaller) (*Il2erc721gatewayCaller, error) {
	contract, err := bindIl2erc721gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Il2erc721gatewayCaller{contract: contract}, nil
}

// NewIl2erc721gatewayTransactor creates a new write-only instance of Il2erc721gateway, bound to a specific deployed contract.
func NewIl2erc721gatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*Il2erc721gatewayTransactor, error) {
	contract, err := bindIl2erc721gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Il2erc721gatewayTransactor{contract: contract}, nil
}

// NewIl2erc721gatewayFilterer creates a new log filterer instance of Il2erc721gateway, bound to a specific deployed contract.
func NewIl2erc721gatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*Il2erc721gatewayFilterer, error) {
	contract, err := bindIl2erc721gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Il2erc721gatewayFilterer{contract: contract}, nil
}

// bindIl2erc721gateway binds a generic wrapper to an already deployed contract.
func bindIl2erc721gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Il2erc721gatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Il2erc721gateway *Il2erc721gatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Il2erc721gateway.Contract.Il2erc721gatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Il2erc721gateway *Il2erc721gatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Il2erc721gateway.Contract.Il2erc721gatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Il2erc721gateway *Il2erc721gatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Il2erc721gateway.Contract.Il2erc721gatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Il2erc721gateway *Il2erc721gatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Il2erc721gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Il2erc721gateway *Il2erc721gatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Il2erc721gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Il2erc721gateway *Il2erc721gatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Il2erc721gateway.Contract.contract.Transact(opts, method, params...)
}

// BatchWithdrawERC721 is a paid mutator transaction binding the contract method 0x46aa3411.
//
// Solidity: function batchWithdrawERC721(address token, uint256[] tokenIds, uint256 gasLimit) payable returns()
func (_Il2erc721gateway *Il2erc721gatewayTransactor) BatchWithdrawERC721(opts *bind.TransactOpts, token common.Address, tokenIds []*big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc721gateway.contract.Transact(opts, "batchWithdrawERC721", token, tokenIds, gasLimit)
}

// BatchWithdrawERC721 is a paid mutator transaction binding the contract method 0x46aa3411.
//
// Solidity: function batchWithdrawERC721(address token, uint256[] tokenIds, uint256 gasLimit) payable returns()
func (_Il2erc721gateway *Il2erc721gatewaySession) BatchWithdrawERC721(token common.Address, tokenIds []*big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc721gateway.Contract.BatchWithdrawERC721(&_Il2erc721gateway.TransactOpts, token, tokenIds, gasLimit)
}

// BatchWithdrawERC721 is a paid mutator transaction binding the contract method 0x46aa3411.
//
// Solidity: function batchWithdrawERC721(address token, uint256[] tokenIds, uint256 gasLimit) payable returns()
func (_Il2erc721gateway *Il2erc721gatewayTransactorSession) BatchWithdrawERC721(token common.Address, tokenIds []*big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc721gateway.Contract.BatchWithdrawERC721(&_Il2erc721gateway.TransactOpts, token, tokenIds, gasLimit)
}

// BatchWithdrawERC7210 is a paid mutator transaction binding the contract method 0xaa4c1158.
//
// Solidity: function batchWithdrawERC721(address token, address to, uint256[] tokenIds, uint256 gasLimit) payable returns()
func (_Il2erc721gateway *Il2erc721gatewayTransactor) BatchWithdrawERC7210(opts *bind.TransactOpts, token common.Address, to common.Address, tokenIds []*big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc721gateway.contract.Transact(opts, "batchWithdrawERC7210", token, to, tokenIds, gasLimit)
}

// BatchWithdrawERC7210 is a paid mutator transaction binding the contract method 0xaa4c1158.
//
// Solidity: function batchWithdrawERC721(address token, address to, uint256[] tokenIds, uint256 gasLimit) payable returns()
func (_Il2erc721gateway *Il2erc721gatewaySession) BatchWithdrawERC7210(token common.Address, to common.Address, tokenIds []*big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc721gateway.Contract.BatchWithdrawERC7210(&_Il2erc721gateway.TransactOpts, token, to, tokenIds, gasLimit)
}

// BatchWithdrawERC7210 is a paid mutator transaction binding the contract method 0xaa4c1158.
//
// Solidity: function batchWithdrawERC721(address token, address to, uint256[] tokenIds, uint256 gasLimit) payable returns()
func (_Il2erc721gateway *Il2erc721gatewayTransactorSession) BatchWithdrawERC7210(token common.Address, to common.Address, tokenIds []*big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc721gateway.Contract.BatchWithdrawERC7210(&_Il2erc721gateway.TransactOpts, token, to, tokenIds, gasLimit)
}

// FinalizeBatchDepositERC721 is a paid mutator transaction binding the contract method 0x982b151f.
//
// Solidity: function finalizeBatchDepositERC721(address l1Token, address l2Token, address from, address to, uint256[] tokenIds) returns()
func (_Il2erc721gateway *Il2erc721gatewayTransactor) FinalizeBatchDepositERC721(opts *bind.TransactOpts, l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Il2erc721gateway.contract.Transact(opts, "finalizeBatchDepositERC721", l1Token, l2Token, from, to, tokenIds)
}

// FinalizeBatchDepositERC721 is a paid mutator transaction binding the contract method 0x982b151f.
//
// Solidity: function finalizeBatchDepositERC721(address l1Token, address l2Token, address from, address to, uint256[] tokenIds) returns()
func (_Il2erc721gateway *Il2erc721gatewaySession) FinalizeBatchDepositERC721(l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Il2erc721gateway.Contract.FinalizeBatchDepositERC721(&_Il2erc721gateway.TransactOpts, l1Token, l2Token, from, to, tokenIds)
}

// FinalizeBatchDepositERC721 is a paid mutator transaction binding the contract method 0x982b151f.
//
// Solidity: function finalizeBatchDepositERC721(address l1Token, address l2Token, address from, address to, uint256[] tokenIds) returns()
func (_Il2erc721gateway *Il2erc721gatewayTransactorSession) FinalizeBatchDepositERC721(l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Il2erc721gateway.Contract.FinalizeBatchDepositERC721(&_Il2erc721gateway.TransactOpts, l1Token, l2Token, from, to, tokenIds)
}

// FinalizeDepositERC721 is a paid mutator transaction binding the contract method 0xf8c3cf25.
//
// Solidity: function finalizeDepositERC721(address l1Token, address l2Token, address from, address to, uint256 tokenId) returns()
func (_Il2erc721gateway *Il2erc721gatewayTransactor) FinalizeDepositERC721(opts *bind.TransactOpts, l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Il2erc721gateway.contract.Transact(opts, "finalizeDepositERC721", l1Token, l2Token, from, to, tokenId)
}

// FinalizeDepositERC721 is a paid mutator transaction binding the contract method 0xf8c3cf25.
//
// Solidity: function finalizeDepositERC721(address l1Token, address l2Token, address from, address to, uint256 tokenId) returns()
func (_Il2erc721gateway *Il2erc721gatewaySession) FinalizeDepositERC721(l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Il2erc721gateway.Contract.FinalizeDepositERC721(&_Il2erc721gateway.TransactOpts, l1Token, l2Token, from, to, tokenId)
}

// FinalizeDepositERC721 is a paid mutator transaction binding the contract method 0xf8c3cf25.
//
// Solidity: function finalizeDepositERC721(address l1Token, address l2Token, address from, address to, uint256 tokenId) returns()
func (_Il2erc721gateway *Il2erc721gatewayTransactorSession) FinalizeDepositERC721(l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Il2erc721gateway.Contract.FinalizeDepositERC721(&_Il2erc721gateway.TransactOpts, l1Token, l2Token, from, to, tokenId)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x2a491247.
//
// Solidity: function withdrawERC721(address token, uint256 tokenId, uint256 gasLimit) payable returns()
func (_Il2erc721gateway *Il2erc721gatewayTransactor) WithdrawERC721(opts *bind.TransactOpts, token common.Address, tokenId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc721gateway.contract.Transact(opts, "withdrawERC721", token, tokenId, gasLimit)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x2a491247.
//
// Solidity: function withdrawERC721(address token, uint256 tokenId, uint256 gasLimit) payable returns()
func (_Il2erc721gateway *Il2erc721gatewaySession) WithdrawERC721(token common.Address, tokenId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc721gateway.Contract.WithdrawERC721(&_Il2erc721gateway.TransactOpts, token, tokenId, gasLimit)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x2a491247.
//
// Solidity: function withdrawERC721(address token, uint256 tokenId, uint256 gasLimit) payable returns()
func (_Il2erc721gateway *Il2erc721gatewayTransactorSession) WithdrawERC721(token common.Address, tokenId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc721gateway.Contract.WithdrawERC721(&_Il2erc721gateway.TransactOpts, token, tokenId, gasLimit)
}

// WithdrawERC7210 is a paid mutator transaction binding the contract method 0xee5a8db2.
//
// Solidity: function withdrawERC721(address token, address to, uint256 tokenId, uint256 gasLimit) payable returns()
func (_Il2erc721gateway *Il2erc721gatewayTransactor) WithdrawERC7210(opts *bind.TransactOpts, token common.Address, to common.Address, tokenId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc721gateway.contract.Transact(opts, "withdrawERC7210", token, to, tokenId, gasLimit)
}

// WithdrawERC7210 is a paid mutator transaction binding the contract method 0xee5a8db2.
//
// Solidity: function withdrawERC721(address token, address to, uint256 tokenId, uint256 gasLimit) payable returns()
func (_Il2erc721gateway *Il2erc721gatewaySession) WithdrawERC7210(token common.Address, to common.Address, tokenId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc721gateway.Contract.WithdrawERC7210(&_Il2erc721gateway.TransactOpts, token, to, tokenId, gasLimit)
}

// WithdrawERC7210 is a paid mutator transaction binding the contract method 0xee5a8db2.
//
// Solidity: function withdrawERC721(address token, address to, uint256 tokenId, uint256 gasLimit) payable returns()
func (_Il2erc721gateway *Il2erc721gatewayTransactorSession) WithdrawERC7210(token common.Address, to common.Address, tokenId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc721gateway.Contract.WithdrawERC7210(&_Il2erc721gateway.TransactOpts, token, to, tokenId, gasLimit)
}

// Il2erc721gatewayBatchWithdrawERC721Iterator is returned from FilterBatchWithdrawERC721 and is used to iterate over the raw logs and unpacked data for BatchWithdrawERC721 events raised by the Il2erc721gateway contract.
type Il2erc721gatewayBatchWithdrawERC721Iterator struct {
	Event *Il2erc721gatewayBatchWithdrawERC721 // Event containing the contract specifics and raw log

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
func (it *Il2erc721gatewayBatchWithdrawERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il2erc721gatewayBatchWithdrawERC721)
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
		it.Event = new(Il2erc721gatewayBatchWithdrawERC721)
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
func (it *Il2erc721gatewayBatchWithdrawERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il2erc721gatewayBatchWithdrawERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il2erc721gatewayBatchWithdrawERC721 represents a BatchWithdrawERC721 event raised by the Il2erc721gateway contract.
type Il2erc721gatewayBatchWithdrawERC721 struct {
	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBatchWithdrawERC721 is a free log retrieval operation binding the contract event 0xbdb7b5cec70093e3ce49b258071951d245c0871c006fd9327778c69d0e9f244d.
//
// Solidity: event BatchWithdrawERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds)
func (_Il2erc721gateway *Il2erc721gatewayFilterer) FilterBatchWithdrawERC721(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*Il2erc721gatewayBatchWithdrawERC721Iterator, error) {

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

	logs, sub, err := _Il2erc721gateway.contract.FilterLogs(opts, "BatchWithdrawERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &Il2erc721gatewayBatchWithdrawERC721Iterator{contract: _Il2erc721gateway.contract, event: "BatchWithdrawERC721", logs: logs, sub: sub}, nil
}

// WatchBatchWithdrawERC721 is a free log subscription operation binding the contract event 0xbdb7b5cec70093e3ce49b258071951d245c0871c006fd9327778c69d0e9f244d.
//
// Solidity: event BatchWithdrawERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds)
func (_Il2erc721gateway *Il2erc721gatewayFilterer) WatchBatchWithdrawERC721(opts *bind.WatchOpts, sink chan<- *Il2erc721gatewayBatchWithdrawERC721, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Il2erc721gateway.contract.WatchLogs(opts, "BatchWithdrawERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il2erc721gatewayBatchWithdrawERC721)
				if err := _Il2erc721gateway.contract.UnpackLog(event, "BatchWithdrawERC721", log); err != nil {
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

// ParseBatchWithdrawERC721 is a log parse operation binding the contract event 0xbdb7b5cec70093e3ce49b258071951d245c0871c006fd9327778c69d0e9f244d.
//
// Solidity: event BatchWithdrawERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds)
func (_Il2erc721gateway *Il2erc721gatewayFilterer) ParseBatchWithdrawERC721(log types.Log) (*Il2erc721gatewayBatchWithdrawERC721, error) {
	event := new(Il2erc721gatewayBatchWithdrawERC721)
	if err := _Il2erc721gateway.contract.UnpackLog(event, "BatchWithdrawERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il2erc721gatewayFinalizeBatchDepositERC721Iterator is returned from FilterFinalizeBatchDepositERC721 and is used to iterate over the raw logs and unpacked data for FinalizeBatchDepositERC721 events raised by the Il2erc721gateway contract.
type Il2erc721gatewayFinalizeBatchDepositERC721Iterator struct {
	Event *Il2erc721gatewayFinalizeBatchDepositERC721 // Event containing the contract specifics and raw log

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
func (it *Il2erc721gatewayFinalizeBatchDepositERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il2erc721gatewayFinalizeBatchDepositERC721)
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
		it.Event = new(Il2erc721gatewayFinalizeBatchDepositERC721)
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
func (it *Il2erc721gatewayFinalizeBatchDepositERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il2erc721gatewayFinalizeBatchDepositERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il2erc721gatewayFinalizeBatchDepositERC721 represents a FinalizeBatchDepositERC721 event raised by the Il2erc721gateway contract.
type Il2erc721gatewayFinalizeBatchDepositERC721 struct {
	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFinalizeBatchDepositERC721 is a free log retrieval operation binding the contract event 0xafa88b850da44ca05b319e813873eac8d08e7c041d2d9b3072db0f087e3cd29e.
//
// Solidity: event FinalizeBatchDepositERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds)
func (_Il2erc721gateway *Il2erc721gatewayFilterer) FilterFinalizeBatchDepositERC721(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*Il2erc721gatewayFinalizeBatchDepositERC721Iterator, error) {

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

	logs, sub, err := _Il2erc721gateway.contract.FilterLogs(opts, "FinalizeBatchDepositERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &Il2erc721gatewayFinalizeBatchDepositERC721Iterator{contract: _Il2erc721gateway.contract, event: "FinalizeBatchDepositERC721", logs: logs, sub: sub}, nil
}

// WatchFinalizeBatchDepositERC721 is a free log subscription operation binding the contract event 0xafa88b850da44ca05b319e813873eac8d08e7c041d2d9b3072db0f087e3cd29e.
//
// Solidity: event FinalizeBatchDepositERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds)
func (_Il2erc721gateway *Il2erc721gatewayFilterer) WatchFinalizeBatchDepositERC721(opts *bind.WatchOpts, sink chan<- *Il2erc721gatewayFinalizeBatchDepositERC721, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Il2erc721gateway.contract.WatchLogs(opts, "FinalizeBatchDepositERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il2erc721gatewayFinalizeBatchDepositERC721)
				if err := _Il2erc721gateway.contract.UnpackLog(event, "FinalizeBatchDepositERC721", log); err != nil {
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

// ParseFinalizeBatchDepositERC721 is a log parse operation binding the contract event 0xafa88b850da44ca05b319e813873eac8d08e7c041d2d9b3072db0f087e3cd29e.
//
// Solidity: event FinalizeBatchDepositERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds)
func (_Il2erc721gateway *Il2erc721gatewayFilterer) ParseFinalizeBatchDepositERC721(log types.Log) (*Il2erc721gatewayFinalizeBatchDepositERC721, error) {
	event := new(Il2erc721gatewayFinalizeBatchDepositERC721)
	if err := _Il2erc721gateway.contract.UnpackLog(event, "FinalizeBatchDepositERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il2erc721gatewayFinalizeDepositERC721Iterator is returned from FilterFinalizeDepositERC721 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC721 events raised by the Il2erc721gateway contract.
type Il2erc721gatewayFinalizeDepositERC721Iterator struct {
	Event *Il2erc721gatewayFinalizeDepositERC721 // Event containing the contract specifics and raw log

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
func (it *Il2erc721gatewayFinalizeDepositERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il2erc721gatewayFinalizeDepositERC721)
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
		it.Event = new(Il2erc721gatewayFinalizeDepositERC721)
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
func (it *Il2erc721gatewayFinalizeDepositERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il2erc721gatewayFinalizeDepositERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il2erc721gatewayFinalizeDepositERC721 represents a FinalizeDepositERC721 event raised by the Il2erc721gateway contract.
type Il2erc721gatewayFinalizeDepositERC721 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDepositERC721 is a free log retrieval operation binding the contract event 0xc655ec1de34d98630aa4572239414f926d6b3d07653dde093a6df97377e31b41.
//
// Solidity: event FinalizeDepositERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId)
func (_Il2erc721gateway *Il2erc721gatewayFilterer) FilterFinalizeDepositERC721(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*Il2erc721gatewayFinalizeDepositERC721Iterator, error) {

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

	logs, sub, err := _Il2erc721gateway.contract.FilterLogs(opts, "FinalizeDepositERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &Il2erc721gatewayFinalizeDepositERC721Iterator{contract: _Il2erc721gateway.contract, event: "FinalizeDepositERC721", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC721 is a free log subscription operation binding the contract event 0xc655ec1de34d98630aa4572239414f926d6b3d07653dde093a6df97377e31b41.
//
// Solidity: event FinalizeDepositERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId)
func (_Il2erc721gateway *Il2erc721gatewayFilterer) WatchFinalizeDepositERC721(opts *bind.WatchOpts, sink chan<- *Il2erc721gatewayFinalizeDepositERC721, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Il2erc721gateway.contract.WatchLogs(opts, "FinalizeDepositERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il2erc721gatewayFinalizeDepositERC721)
				if err := _Il2erc721gateway.contract.UnpackLog(event, "FinalizeDepositERC721", log); err != nil {
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

// ParseFinalizeDepositERC721 is a log parse operation binding the contract event 0xc655ec1de34d98630aa4572239414f926d6b3d07653dde093a6df97377e31b41.
//
// Solidity: event FinalizeDepositERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId)
func (_Il2erc721gateway *Il2erc721gatewayFilterer) ParseFinalizeDepositERC721(log types.Log) (*Il2erc721gatewayFinalizeDepositERC721, error) {
	event := new(Il2erc721gatewayFinalizeDepositERC721)
	if err := _Il2erc721gateway.contract.UnpackLog(event, "FinalizeDepositERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il2erc721gatewayWithdrawERC721Iterator is returned from FilterWithdrawERC721 and is used to iterate over the raw logs and unpacked data for WithdrawERC721 events raised by the Il2erc721gateway contract.
type Il2erc721gatewayWithdrawERC721Iterator struct {
	Event *Il2erc721gatewayWithdrawERC721 // Event containing the contract specifics and raw log

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
func (it *Il2erc721gatewayWithdrawERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il2erc721gatewayWithdrawERC721)
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
		it.Event = new(Il2erc721gatewayWithdrawERC721)
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
func (it *Il2erc721gatewayWithdrawERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il2erc721gatewayWithdrawERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il2erc721gatewayWithdrawERC721 represents a WithdrawERC721 event raised by the Il2erc721gateway contract.
type Il2erc721gatewayWithdrawERC721 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawERC721 is a free log retrieval operation binding the contract event 0xe9e85cf0c862dd491ecda3c9a230e12ada8956472028ebde4fdc4f8e2d77bcda.
//
// Solidity: event WithdrawERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId)
func (_Il2erc721gateway *Il2erc721gatewayFilterer) FilterWithdrawERC721(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*Il2erc721gatewayWithdrawERC721Iterator, error) {

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

	logs, sub, err := _Il2erc721gateway.contract.FilterLogs(opts, "WithdrawERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &Il2erc721gatewayWithdrawERC721Iterator{contract: _Il2erc721gateway.contract, event: "WithdrawERC721", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC721 is a free log subscription operation binding the contract event 0xe9e85cf0c862dd491ecda3c9a230e12ada8956472028ebde4fdc4f8e2d77bcda.
//
// Solidity: event WithdrawERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId)
func (_Il2erc721gateway *Il2erc721gatewayFilterer) WatchWithdrawERC721(opts *bind.WatchOpts, sink chan<- *Il2erc721gatewayWithdrawERC721, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Il2erc721gateway.contract.WatchLogs(opts, "WithdrawERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il2erc721gatewayWithdrawERC721)
				if err := _Il2erc721gateway.contract.UnpackLog(event, "WithdrawERC721", log); err != nil {
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

// ParseWithdrawERC721 is a log parse operation binding the contract event 0xe9e85cf0c862dd491ecda3c9a230e12ada8956472028ebde4fdc4f8e2d77bcda.
//
// Solidity: event WithdrawERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId)
func (_Il2erc721gateway *Il2erc721gatewayFilterer) ParseWithdrawERC721(log types.Log) (*Il2erc721gatewayWithdrawERC721, error) {
	event := new(Il2erc721gatewayWithdrawERC721)
	if err := _Il2erc721gateway.contract.UnpackLog(event, "WithdrawERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
