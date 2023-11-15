// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package il2erc1155gateway

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

// Il2erc1155gatewayMetaData contains all meta data concerning the Il2erc1155gateway contract.
var Il2erc1155gatewayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"BatchWithdrawERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"FinalizeBatchDepositERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FinalizeDepositERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC1155\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"batchWithdrawERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"batchWithdrawERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"finalizeBatchDepositERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"finalizeDepositERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// Il2erc1155gatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use Il2erc1155gatewayMetaData.ABI instead.
var Il2erc1155gatewayABI = Il2erc1155gatewayMetaData.ABI

// Il2erc1155gateway is an auto generated Go binding around an Ethereum contract.
type Il2erc1155gateway struct {
	Il2erc1155gatewayCaller     // Read-only binding to the contract
	Il2erc1155gatewayTransactor // Write-only binding to the contract
	Il2erc1155gatewayFilterer   // Log filterer for contract events
}

// Il2erc1155gatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type Il2erc1155gatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il2erc1155gatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Il2erc1155gatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il2erc1155gatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Il2erc1155gatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il2erc1155gatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Il2erc1155gatewaySession struct {
	Contract     *Il2erc1155gateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Il2erc1155gatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Il2erc1155gatewayCallerSession struct {
	Contract *Il2erc1155gatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// Il2erc1155gatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Il2erc1155gatewayTransactorSession struct {
	Contract     *Il2erc1155gatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// Il2erc1155gatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type Il2erc1155gatewayRaw struct {
	Contract *Il2erc1155gateway // Generic contract binding to access the raw methods on
}

// Il2erc1155gatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Il2erc1155gatewayCallerRaw struct {
	Contract *Il2erc1155gatewayCaller // Generic read-only contract binding to access the raw methods on
}

// Il2erc1155gatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Il2erc1155gatewayTransactorRaw struct {
	Contract *Il2erc1155gatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIl2erc1155gateway creates a new instance of Il2erc1155gateway, bound to a specific deployed contract.
func NewIl2erc1155gateway(address common.Address, backend bind.ContractBackend) (*Il2erc1155gateway, error) {
	contract, err := bindIl2erc1155gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Il2erc1155gateway{Il2erc1155gatewayCaller: Il2erc1155gatewayCaller{contract: contract}, Il2erc1155gatewayTransactor: Il2erc1155gatewayTransactor{contract: contract}, Il2erc1155gatewayFilterer: Il2erc1155gatewayFilterer{contract: contract}}, nil
}

// NewIl2erc1155gatewayCaller creates a new read-only instance of Il2erc1155gateway, bound to a specific deployed contract.
func NewIl2erc1155gatewayCaller(address common.Address, caller bind.ContractCaller) (*Il2erc1155gatewayCaller, error) {
	contract, err := bindIl2erc1155gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Il2erc1155gatewayCaller{contract: contract}, nil
}

// NewIl2erc1155gatewayTransactor creates a new write-only instance of Il2erc1155gateway, bound to a specific deployed contract.
func NewIl2erc1155gatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*Il2erc1155gatewayTransactor, error) {
	contract, err := bindIl2erc1155gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Il2erc1155gatewayTransactor{contract: contract}, nil
}

// NewIl2erc1155gatewayFilterer creates a new log filterer instance of Il2erc1155gateway, bound to a specific deployed contract.
func NewIl2erc1155gatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*Il2erc1155gatewayFilterer, error) {
	contract, err := bindIl2erc1155gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Il2erc1155gatewayFilterer{contract: contract}, nil
}

// bindIl2erc1155gateway binds a generic wrapper to an already deployed contract.
func bindIl2erc1155gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Il2erc1155gatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Il2erc1155gateway *Il2erc1155gatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Il2erc1155gateway.Contract.Il2erc1155gatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Il2erc1155gateway *Il2erc1155gatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Il2erc1155gateway.Contract.Il2erc1155gatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Il2erc1155gateway *Il2erc1155gatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Il2erc1155gateway.Contract.Il2erc1155gatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Il2erc1155gateway *Il2erc1155gatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Il2erc1155gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Il2erc1155gateway *Il2erc1155gatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Il2erc1155gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Il2erc1155gateway *Il2erc1155gatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Il2erc1155gateway.Contract.contract.Transact(opts, method, params...)
}

// BatchWithdrawERC1155 is a paid mutator transaction binding the contract method 0x48de03de.
//
// Solidity: function batchWithdrawERC1155(address token, uint256[] tokenIds, uint256[] amounts, uint256 gasLimit) payable returns()
func (_Il2erc1155gateway *Il2erc1155gatewayTransactor) BatchWithdrawERC1155(opts *bind.TransactOpts, token common.Address, tokenIds []*big.Int, amounts []*big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc1155gateway.contract.Transact(opts, "batchWithdrawERC1155", token, tokenIds, amounts, gasLimit)
}

// BatchWithdrawERC1155 is a paid mutator transaction binding the contract method 0x48de03de.
//
// Solidity: function batchWithdrawERC1155(address token, uint256[] tokenIds, uint256[] amounts, uint256 gasLimit) payable returns()
func (_Il2erc1155gateway *Il2erc1155gatewaySession) BatchWithdrawERC1155(token common.Address, tokenIds []*big.Int, amounts []*big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc1155gateway.Contract.BatchWithdrawERC1155(&_Il2erc1155gateway.TransactOpts, token, tokenIds, amounts, gasLimit)
}

// BatchWithdrawERC1155 is a paid mutator transaction binding the contract method 0x48de03de.
//
// Solidity: function batchWithdrawERC1155(address token, uint256[] tokenIds, uint256[] amounts, uint256 gasLimit) payable returns()
func (_Il2erc1155gateway *Il2erc1155gatewayTransactorSession) BatchWithdrawERC1155(token common.Address, tokenIds []*big.Int, amounts []*big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc1155gateway.Contract.BatchWithdrawERC1155(&_Il2erc1155gateway.TransactOpts, token, tokenIds, amounts, gasLimit)
}

// BatchWithdrawERC11550 is a paid mutator transaction binding the contract method 0x8c23d5b2.
//
// Solidity: function batchWithdrawERC1155(address token, address to, uint256[] tokenIds, uint256[] amounts, uint256 gasLimit) payable returns()
func (_Il2erc1155gateway *Il2erc1155gatewayTransactor) BatchWithdrawERC11550(opts *bind.TransactOpts, token common.Address, to common.Address, tokenIds []*big.Int, amounts []*big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc1155gateway.contract.Transact(opts, "batchWithdrawERC11550", token, to, tokenIds, amounts, gasLimit)
}

// BatchWithdrawERC11550 is a paid mutator transaction binding the contract method 0x8c23d5b2.
//
// Solidity: function batchWithdrawERC1155(address token, address to, uint256[] tokenIds, uint256[] amounts, uint256 gasLimit) payable returns()
func (_Il2erc1155gateway *Il2erc1155gatewaySession) BatchWithdrawERC11550(token common.Address, to common.Address, tokenIds []*big.Int, amounts []*big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc1155gateway.Contract.BatchWithdrawERC11550(&_Il2erc1155gateway.TransactOpts, token, to, tokenIds, amounts, gasLimit)
}

// BatchWithdrawERC11550 is a paid mutator transaction binding the contract method 0x8c23d5b2.
//
// Solidity: function batchWithdrawERC1155(address token, address to, uint256[] tokenIds, uint256[] amounts, uint256 gasLimit) payable returns()
func (_Il2erc1155gateway *Il2erc1155gatewayTransactorSession) BatchWithdrawERC11550(token common.Address, to common.Address, tokenIds []*big.Int, amounts []*big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc1155gateway.Contract.BatchWithdrawERC11550(&_Il2erc1155gateway.TransactOpts, token, to, tokenIds, amounts, gasLimit)
}

// FinalizeBatchDepositERC1155 is a paid mutator transaction binding the contract method 0xeaa72ad9.
//
// Solidity: function finalizeBatchDepositERC1155(address l1Token, address l2Token, address from, address to, uint256[] tokenIds, uint256[] amounts) returns()
func (_Il2erc1155gateway *Il2erc1155gatewayTransactor) FinalizeBatchDepositERC1155(opts *bind.TransactOpts, l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Il2erc1155gateway.contract.Transact(opts, "finalizeBatchDepositERC1155", l1Token, l2Token, from, to, tokenIds, amounts)
}

// FinalizeBatchDepositERC1155 is a paid mutator transaction binding the contract method 0xeaa72ad9.
//
// Solidity: function finalizeBatchDepositERC1155(address l1Token, address l2Token, address from, address to, uint256[] tokenIds, uint256[] amounts) returns()
func (_Il2erc1155gateway *Il2erc1155gatewaySession) FinalizeBatchDepositERC1155(l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Il2erc1155gateway.Contract.FinalizeBatchDepositERC1155(&_Il2erc1155gateway.TransactOpts, l1Token, l2Token, from, to, tokenIds, amounts)
}

// FinalizeBatchDepositERC1155 is a paid mutator transaction binding the contract method 0xeaa72ad9.
//
// Solidity: function finalizeBatchDepositERC1155(address l1Token, address l2Token, address from, address to, uint256[] tokenIds, uint256[] amounts) returns()
func (_Il2erc1155gateway *Il2erc1155gatewayTransactorSession) FinalizeBatchDepositERC1155(l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Il2erc1155gateway.Contract.FinalizeBatchDepositERC1155(&_Il2erc1155gateway.TransactOpts, l1Token, l2Token, from, to, tokenIds, amounts)
}

// FinalizeDepositERC1155 is a paid mutator transaction binding the contract method 0x4764cc62.
//
// Solidity: function finalizeDepositERC1155(address l1Token, address l2Token, address from, address to, uint256 tokenId, uint256 amount) returns()
func (_Il2erc1155gateway *Il2erc1155gatewayTransactor) FinalizeDepositERC1155(opts *bind.TransactOpts, l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Il2erc1155gateway.contract.Transact(opts, "finalizeDepositERC1155", l1Token, l2Token, from, to, tokenId, amount)
}

// FinalizeDepositERC1155 is a paid mutator transaction binding the contract method 0x4764cc62.
//
// Solidity: function finalizeDepositERC1155(address l1Token, address l2Token, address from, address to, uint256 tokenId, uint256 amount) returns()
func (_Il2erc1155gateway *Il2erc1155gatewaySession) FinalizeDepositERC1155(l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Il2erc1155gateway.Contract.FinalizeDepositERC1155(&_Il2erc1155gateway.TransactOpts, l1Token, l2Token, from, to, tokenId, amount)
}

// FinalizeDepositERC1155 is a paid mutator transaction binding the contract method 0x4764cc62.
//
// Solidity: function finalizeDepositERC1155(address l1Token, address l2Token, address from, address to, uint256 tokenId, uint256 amount) returns()
func (_Il2erc1155gateway *Il2erc1155gatewayTransactorSession) FinalizeDepositERC1155(l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Il2erc1155gateway.Contract.FinalizeDepositERC1155(&_Il2erc1155gateway.TransactOpts, l1Token, l2Token, from, to, tokenId, amount)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x0f2da080.
//
// Solidity: function withdrawERC1155(address token, uint256 tokenId, uint256 amount, uint256 gasLimit) payable returns()
func (_Il2erc1155gateway *Il2erc1155gatewayTransactor) WithdrawERC1155(opts *bind.TransactOpts, token common.Address, tokenId *big.Int, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc1155gateway.contract.Transact(opts, "withdrawERC1155", token, tokenId, amount, gasLimit)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x0f2da080.
//
// Solidity: function withdrawERC1155(address token, uint256 tokenId, uint256 amount, uint256 gasLimit) payable returns()
func (_Il2erc1155gateway *Il2erc1155gatewaySession) WithdrawERC1155(token common.Address, tokenId *big.Int, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc1155gateway.Contract.WithdrawERC1155(&_Il2erc1155gateway.TransactOpts, token, tokenId, amount, gasLimit)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x0f2da080.
//
// Solidity: function withdrawERC1155(address token, uint256 tokenId, uint256 amount, uint256 gasLimit) payable returns()
func (_Il2erc1155gateway *Il2erc1155gatewayTransactorSession) WithdrawERC1155(token common.Address, tokenId *big.Int, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc1155gateway.Contract.WithdrawERC1155(&_Il2erc1155gateway.TransactOpts, token, tokenId, amount, gasLimit)
}

// WithdrawERC11550 is a paid mutator transaction binding the contract method 0x21fedfc9.
//
// Solidity: function withdrawERC1155(address token, address to, uint256 tokenId, uint256 amount, uint256 gasLimit) payable returns()
func (_Il2erc1155gateway *Il2erc1155gatewayTransactor) WithdrawERC11550(opts *bind.TransactOpts, token common.Address, to common.Address, tokenId *big.Int, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc1155gateway.contract.Transact(opts, "withdrawERC11550", token, to, tokenId, amount, gasLimit)
}

// WithdrawERC11550 is a paid mutator transaction binding the contract method 0x21fedfc9.
//
// Solidity: function withdrawERC1155(address token, address to, uint256 tokenId, uint256 amount, uint256 gasLimit) payable returns()
func (_Il2erc1155gateway *Il2erc1155gatewaySession) WithdrawERC11550(token common.Address, to common.Address, tokenId *big.Int, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc1155gateway.Contract.WithdrawERC11550(&_Il2erc1155gateway.TransactOpts, token, to, tokenId, amount, gasLimit)
}

// WithdrawERC11550 is a paid mutator transaction binding the contract method 0x21fedfc9.
//
// Solidity: function withdrawERC1155(address token, address to, uint256 tokenId, uint256 amount, uint256 gasLimit) payable returns()
func (_Il2erc1155gateway *Il2erc1155gatewayTransactorSession) WithdrawERC11550(token common.Address, to common.Address, tokenId *big.Int, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _Il2erc1155gateway.Contract.WithdrawERC11550(&_Il2erc1155gateway.TransactOpts, token, to, tokenId, amount, gasLimit)
}

// Il2erc1155gatewayBatchWithdrawERC1155Iterator is returned from FilterBatchWithdrawERC1155 and is used to iterate over the raw logs and unpacked data for BatchWithdrawERC1155 events raised by the Il2erc1155gateway contract.
type Il2erc1155gatewayBatchWithdrawERC1155Iterator struct {
	Event *Il2erc1155gatewayBatchWithdrawERC1155 // Event containing the contract specifics and raw log

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
func (it *Il2erc1155gatewayBatchWithdrawERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il2erc1155gatewayBatchWithdrawERC1155)
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
		it.Event = new(Il2erc1155gatewayBatchWithdrawERC1155)
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
func (it *Il2erc1155gatewayBatchWithdrawERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il2erc1155gatewayBatchWithdrawERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il2erc1155gatewayBatchWithdrawERC1155 represents a BatchWithdrawERC1155 event raised by the Il2erc1155gateway contract.
type Il2erc1155gatewayBatchWithdrawERC1155 struct {
	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Amounts  []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBatchWithdrawERC1155 is a free log retrieval operation binding the contract event 0x5d2d5d4cdbf7b115e43f0b9986644dd8b9514b10be6a019ab6a4a87f12290970.
//
// Solidity: event BatchWithdrawERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds, uint256[] amounts)
func (_Il2erc1155gateway *Il2erc1155gatewayFilterer) FilterBatchWithdrawERC1155(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*Il2erc1155gatewayBatchWithdrawERC1155Iterator, error) {

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

	logs, sub, err := _Il2erc1155gateway.contract.FilterLogs(opts, "BatchWithdrawERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &Il2erc1155gatewayBatchWithdrawERC1155Iterator{contract: _Il2erc1155gateway.contract, event: "BatchWithdrawERC1155", logs: logs, sub: sub}, nil
}

// WatchBatchWithdrawERC1155 is a free log subscription operation binding the contract event 0x5d2d5d4cdbf7b115e43f0b9986644dd8b9514b10be6a019ab6a4a87f12290970.
//
// Solidity: event BatchWithdrawERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds, uint256[] amounts)
func (_Il2erc1155gateway *Il2erc1155gatewayFilterer) WatchBatchWithdrawERC1155(opts *bind.WatchOpts, sink chan<- *Il2erc1155gatewayBatchWithdrawERC1155, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Il2erc1155gateway.contract.WatchLogs(opts, "BatchWithdrawERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il2erc1155gatewayBatchWithdrawERC1155)
				if err := _Il2erc1155gateway.contract.UnpackLog(event, "BatchWithdrawERC1155", log); err != nil {
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

// ParseBatchWithdrawERC1155 is a log parse operation binding the contract event 0x5d2d5d4cdbf7b115e43f0b9986644dd8b9514b10be6a019ab6a4a87f12290970.
//
// Solidity: event BatchWithdrawERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds, uint256[] amounts)
func (_Il2erc1155gateway *Il2erc1155gatewayFilterer) ParseBatchWithdrawERC1155(log types.Log) (*Il2erc1155gatewayBatchWithdrawERC1155, error) {
	event := new(Il2erc1155gatewayBatchWithdrawERC1155)
	if err := _Il2erc1155gateway.contract.UnpackLog(event, "BatchWithdrawERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il2erc1155gatewayFinalizeBatchDepositERC1155Iterator is returned from FilterFinalizeBatchDepositERC1155 and is used to iterate over the raw logs and unpacked data for FinalizeBatchDepositERC1155 events raised by the Il2erc1155gateway contract.
type Il2erc1155gatewayFinalizeBatchDepositERC1155Iterator struct {
	Event *Il2erc1155gatewayFinalizeBatchDepositERC1155 // Event containing the contract specifics and raw log

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
func (it *Il2erc1155gatewayFinalizeBatchDepositERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il2erc1155gatewayFinalizeBatchDepositERC1155)
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
		it.Event = new(Il2erc1155gatewayFinalizeBatchDepositERC1155)
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
func (it *Il2erc1155gatewayFinalizeBatchDepositERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il2erc1155gatewayFinalizeBatchDepositERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il2erc1155gatewayFinalizeBatchDepositERC1155 represents a FinalizeBatchDepositERC1155 event raised by the Il2erc1155gateway contract.
type Il2erc1155gatewayFinalizeBatchDepositERC1155 struct {
	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Amounts  []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFinalizeBatchDepositERC1155 is a free log retrieval operation binding the contract event 0xf07745bfeb45fb1184165136e9148689adf57ba578a5b90dde949f26066b7756.
//
// Solidity: event FinalizeBatchDepositERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds, uint256[] amounts)
func (_Il2erc1155gateway *Il2erc1155gatewayFilterer) FilterFinalizeBatchDepositERC1155(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*Il2erc1155gatewayFinalizeBatchDepositERC1155Iterator, error) {

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

	logs, sub, err := _Il2erc1155gateway.contract.FilterLogs(opts, "FinalizeBatchDepositERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &Il2erc1155gatewayFinalizeBatchDepositERC1155Iterator{contract: _Il2erc1155gateway.contract, event: "FinalizeBatchDepositERC1155", logs: logs, sub: sub}, nil
}

// WatchFinalizeBatchDepositERC1155 is a free log subscription operation binding the contract event 0xf07745bfeb45fb1184165136e9148689adf57ba578a5b90dde949f26066b7756.
//
// Solidity: event FinalizeBatchDepositERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds, uint256[] amounts)
func (_Il2erc1155gateway *Il2erc1155gatewayFilterer) WatchFinalizeBatchDepositERC1155(opts *bind.WatchOpts, sink chan<- *Il2erc1155gatewayFinalizeBatchDepositERC1155, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Il2erc1155gateway.contract.WatchLogs(opts, "FinalizeBatchDepositERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il2erc1155gatewayFinalizeBatchDepositERC1155)
				if err := _Il2erc1155gateway.contract.UnpackLog(event, "FinalizeBatchDepositERC1155", log); err != nil {
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

// ParseFinalizeBatchDepositERC1155 is a log parse operation binding the contract event 0xf07745bfeb45fb1184165136e9148689adf57ba578a5b90dde949f26066b7756.
//
// Solidity: event FinalizeBatchDepositERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds, uint256[] amounts)
func (_Il2erc1155gateway *Il2erc1155gatewayFilterer) ParseFinalizeBatchDepositERC1155(log types.Log) (*Il2erc1155gatewayFinalizeBatchDepositERC1155, error) {
	event := new(Il2erc1155gatewayFinalizeBatchDepositERC1155)
	if err := _Il2erc1155gateway.contract.UnpackLog(event, "FinalizeBatchDepositERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il2erc1155gatewayFinalizeDepositERC1155Iterator is returned from FilterFinalizeDepositERC1155 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC1155 events raised by the Il2erc1155gateway contract.
type Il2erc1155gatewayFinalizeDepositERC1155Iterator struct {
	Event *Il2erc1155gatewayFinalizeDepositERC1155 // Event containing the contract specifics and raw log

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
func (it *Il2erc1155gatewayFinalizeDepositERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il2erc1155gatewayFinalizeDepositERC1155)
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
		it.Event = new(Il2erc1155gatewayFinalizeDepositERC1155)
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
func (it *Il2erc1155gatewayFinalizeDepositERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il2erc1155gatewayFinalizeDepositERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il2erc1155gatewayFinalizeDepositERC1155 represents a FinalizeDepositERC1155 event raised by the Il2erc1155gateway contract.
type Il2erc1155gatewayFinalizeDepositERC1155 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDepositERC1155 is a free log retrieval operation binding the contract event 0x5399dc7b86d085e50a28946dbc213966bb7a7ac78d312aedd6018c791ad6cef9.
//
// Solidity: event FinalizeDepositERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId, uint256 amount)
func (_Il2erc1155gateway *Il2erc1155gatewayFilterer) FilterFinalizeDepositERC1155(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*Il2erc1155gatewayFinalizeDepositERC1155Iterator, error) {

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

	logs, sub, err := _Il2erc1155gateway.contract.FilterLogs(opts, "FinalizeDepositERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &Il2erc1155gatewayFinalizeDepositERC1155Iterator{contract: _Il2erc1155gateway.contract, event: "FinalizeDepositERC1155", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC1155 is a free log subscription operation binding the contract event 0x5399dc7b86d085e50a28946dbc213966bb7a7ac78d312aedd6018c791ad6cef9.
//
// Solidity: event FinalizeDepositERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId, uint256 amount)
func (_Il2erc1155gateway *Il2erc1155gatewayFilterer) WatchFinalizeDepositERC1155(opts *bind.WatchOpts, sink chan<- *Il2erc1155gatewayFinalizeDepositERC1155, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Il2erc1155gateway.contract.WatchLogs(opts, "FinalizeDepositERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il2erc1155gatewayFinalizeDepositERC1155)
				if err := _Il2erc1155gateway.contract.UnpackLog(event, "FinalizeDepositERC1155", log); err != nil {
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

// ParseFinalizeDepositERC1155 is a log parse operation binding the contract event 0x5399dc7b86d085e50a28946dbc213966bb7a7ac78d312aedd6018c791ad6cef9.
//
// Solidity: event FinalizeDepositERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId, uint256 amount)
func (_Il2erc1155gateway *Il2erc1155gatewayFilterer) ParseFinalizeDepositERC1155(log types.Log) (*Il2erc1155gatewayFinalizeDepositERC1155, error) {
	event := new(Il2erc1155gatewayFinalizeDepositERC1155)
	if err := _Il2erc1155gateway.contract.UnpackLog(event, "FinalizeDepositERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il2erc1155gatewayWithdrawERC1155Iterator is returned from FilterWithdrawERC1155 and is used to iterate over the raw logs and unpacked data for WithdrawERC1155 events raised by the Il2erc1155gateway contract.
type Il2erc1155gatewayWithdrawERC1155Iterator struct {
	Event *Il2erc1155gatewayWithdrawERC1155 // Event containing the contract specifics and raw log

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
func (it *Il2erc1155gatewayWithdrawERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il2erc1155gatewayWithdrawERC1155)
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
		it.Event = new(Il2erc1155gatewayWithdrawERC1155)
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
func (it *Il2erc1155gatewayWithdrawERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il2erc1155gatewayWithdrawERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il2erc1155gatewayWithdrawERC1155 represents a WithdrawERC1155 event raised by the Il2erc1155gateway contract.
type Il2erc1155gatewayWithdrawERC1155 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawERC1155 is a free log retrieval operation binding the contract event 0x1f9dcda7fce6f73a13055f044ffecaed2032a7a844e0a37a3eb8bbb17488d01a.
//
// Solidity: event WithdrawERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId, uint256 amount)
func (_Il2erc1155gateway *Il2erc1155gatewayFilterer) FilterWithdrawERC1155(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*Il2erc1155gatewayWithdrawERC1155Iterator, error) {

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

	logs, sub, err := _Il2erc1155gateway.contract.FilterLogs(opts, "WithdrawERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &Il2erc1155gatewayWithdrawERC1155Iterator{contract: _Il2erc1155gateway.contract, event: "WithdrawERC1155", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC1155 is a free log subscription operation binding the contract event 0x1f9dcda7fce6f73a13055f044ffecaed2032a7a844e0a37a3eb8bbb17488d01a.
//
// Solidity: event WithdrawERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId, uint256 amount)
func (_Il2erc1155gateway *Il2erc1155gatewayFilterer) WatchWithdrawERC1155(opts *bind.WatchOpts, sink chan<- *Il2erc1155gatewayWithdrawERC1155, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Il2erc1155gateway.contract.WatchLogs(opts, "WithdrawERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il2erc1155gatewayWithdrawERC1155)
				if err := _Il2erc1155gateway.contract.UnpackLog(event, "WithdrawERC1155", log); err != nil {
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

// ParseWithdrawERC1155 is a log parse operation binding the contract event 0x1f9dcda7fce6f73a13055f044ffecaed2032a7a844e0a37a3eb8bbb17488d01a.
//
// Solidity: event WithdrawERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId, uint256 amount)
func (_Il2erc1155gateway *Il2erc1155gatewayFilterer) ParseWithdrawERC1155(log types.Log) (*Il2erc1155gatewayWithdrawERC1155, error) {
	event := new(Il2erc1155gatewayWithdrawERC1155)
	if err := _Il2erc1155gateway.contract.UnpackLog(event, "WithdrawERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
