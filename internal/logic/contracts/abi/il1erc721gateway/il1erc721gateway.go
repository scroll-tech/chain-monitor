// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package il1erc721gateway

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

// Il1erc721gatewayMetaData contains all meta data concerning the Il1erc721gateway contract.
var Il1erc721gatewayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"BatchDepositERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"BatchRefundERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"DepositERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"FinalizeBatchWithdrawERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"FinalizeWithdrawERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"RefundERC721\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"batchDepositERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"batchDepositERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"finalizeBatchWithdrawERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"finalizeWithdrawERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Il1erc721gatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use Il1erc721gatewayMetaData.ABI instead.
var Il1erc721gatewayABI = Il1erc721gatewayMetaData.ABI

// Il1erc721gateway is an auto generated Go binding around an Ethereum contract.
type Il1erc721gateway struct {
	Il1erc721gatewayCaller     // Read-only binding to the contract
	Il1erc721gatewayTransactor // Write-only binding to the contract
	Il1erc721gatewayFilterer   // Log filterer for contract events
}

// Il1erc721gatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type Il1erc721gatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il1erc721gatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Il1erc721gatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il1erc721gatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Il1erc721gatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il1erc721gatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Il1erc721gatewaySession struct {
	Contract     *Il1erc721gateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Il1erc721gatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Il1erc721gatewayCallerSession struct {
	Contract *Il1erc721gatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Il1erc721gatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Il1erc721gatewayTransactorSession struct {
	Contract     *Il1erc721gatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Il1erc721gatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type Il1erc721gatewayRaw struct {
	Contract *Il1erc721gateway // Generic contract binding to access the raw methods on
}

// Il1erc721gatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Il1erc721gatewayCallerRaw struct {
	Contract *Il1erc721gatewayCaller // Generic read-only contract binding to access the raw methods on
}

// Il1erc721gatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Il1erc721gatewayTransactorRaw struct {
	Contract *Il1erc721gatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIl1erc721gateway creates a new instance of Il1erc721gateway, bound to a specific deployed contract.
func NewIl1erc721gateway(address common.Address, backend bind.ContractBackend) (*Il1erc721gateway, error) {
	contract, err := bindIl1erc721gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Il1erc721gateway{Il1erc721gatewayCaller: Il1erc721gatewayCaller{contract: contract}, Il1erc721gatewayTransactor: Il1erc721gatewayTransactor{contract: contract}, Il1erc721gatewayFilterer: Il1erc721gatewayFilterer{contract: contract}}, nil
}

// NewIl1erc721gatewayCaller creates a new read-only instance of Il1erc721gateway, bound to a specific deployed contract.
func NewIl1erc721gatewayCaller(address common.Address, caller bind.ContractCaller) (*Il1erc721gatewayCaller, error) {
	contract, err := bindIl1erc721gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Il1erc721gatewayCaller{contract: contract}, nil
}

// NewIl1erc721gatewayTransactor creates a new write-only instance of Il1erc721gateway, bound to a specific deployed contract.
func NewIl1erc721gatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*Il1erc721gatewayTransactor, error) {
	contract, err := bindIl1erc721gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Il1erc721gatewayTransactor{contract: contract}, nil
}

// NewIl1erc721gatewayFilterer creates a new log filterer instance of Il1erc721gateway, bound to a specific deployed contract.
func NewIl1erc721gatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*Il1erc721gatewayFilterer, error) {
	contract, err := bindIl1erc721gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Il1erc721gatewayFilterer{contract: contract}, nil
}

// bindIl1erc721gateway binds a generic wrapper to an already deployed contract.
func bindIl1erc721gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Il1erc721gatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Il1erc721gateway *Il1erc721gatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Il1erc721gateway.Contract.Il1erc721gatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Il1erc721gateway *Il1erc721gatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Il1erc721gateway.Contract.Il1erc721gatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Il1erc721gateway *Il1erc721gatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Il1erc721gateway.Contract.Il1erc721gatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Il1erc721gateway *Il1erc721gatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Il1erc721gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Il1erc721gateway *Il1erc721gatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Il1erc721gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Il1erc721gateway *Il1erc721gatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Il1erc721gateway.Contract.contract.Transact(opts, method, params...)
}

// BatchDepositERC721 is a paid mutator transaction binding the contract method 0x1b997a93.
//
// Solidity: function batchDepositERC721(address _token, address _to, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_Il1erc721gateway *Il1erc721gatewayTransactor) BatchDepositERC721(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc721gateway.contract.Transact(opts, "batchDepositERC721", _token, _to, _tokenIds, _gasLimit)
}

// BatchDepositERC721 is a paid mutator transaction binding the contract method 0x1b997a93.
//
// Solidity: function batchDepositERC721(address _token, address _to, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_Il1erc721gateway *Il1erc721gatewaySession) BatchDepositERC721(_token common.Address, _to common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc721gateway.Contract.BatchDepositERC721(&_Il1erc721gateway.TransactOpts, _token, _to, _tokenIds, _gasLimit)
}

// BatchDepositERC721 is a paid mutator transaction binding the contract method 0x1b997a93.
//
// Solidity: function batchDepositERC721(address _token, address _to, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_Il1erc721gateway *Il1erc721gatewayTransactorSession) BatchDepositERC721(_token common.Address, _to common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc721gateway.Contract.BatchDepositERC721(&_Il1erc721gateway.TransactOpts, _token, _to, _tokenIds, _gasLimit)
}

// BatchDepositERC7210 is a paid mutator transaction binding the contract method 0xd96c8ecf.
//
// Solidity: function batchDepositERC721(address _token, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_Il1erc721gateway *Il1erc721gatewayTransactor) BatchDepositERC7210(opts *bind.TransactOpts, _token common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc721gateway.contract.Transact(opts, "batchDepositERC7210", _token, _tokenIds, _gasLimit)
}

// BatchDepositERC7210 is a paid mutator transaction binding the contract method 0xd96c8ecf.
//
// Solidity: function batchDepositERC721(address _token, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_Il1erc721gateway *Il1erc721gatewaySession) BatchDepositERC7210(_token common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc721gateway.Contract.BatchDepositERC7210(&_Il1erc721gateway.TransactOpts, _token, _tokenIds, _gasLimit)
}

// BatchDepositERC7210 is a paid mutator transaction binding the contract method 0xd96c8ecf.
//
// Solidity: function batchDepositERC721(address _token, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_Il1erc721gateway *Il1erc721gatewayTransactorSession) BatchDepositERC7210(_token common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc721gateway.Contract.BatchDepositERC7210(&_Il1erc721gateway.TransactOpts, _token, _tokenIds, _gasLimit)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x0a7aa196.
//
// Solidity: function depositERC721(address _token, address _to, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_Il1erc721gateway *Il1erc721gatewayTransactor) DepositERC721(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc721gateway.contract.Transact(opts, "depositERC721", _token, _to, _tokenId, _gasLimit)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x0a7aa196.
//
// Solidity: function depositERC721(address _token, address _to, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_Il1erc721gateway *Il1erc721gatewaySession) DepositERC721(_token common.Address, _to common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc721gateway.Contract.DepositERC721(&_Il1erc721gateway.TransactOpts, _token, _to, _tokenId, _gasLimit)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x0a7aa196.
//
// Solidity: function depositERC721(address _token, address _to, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_Il1erc721gateway *Il1erc721gatewayTransactorSession) DepositERC721(_token common.Address, _to common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc721gateway.Contract.DepositERC721(&_Il1erc721gateway.TransactOpts, _token, _to, _tokenId, _gasLimit)
}

// DepositERC7210 is a paid mutator transaction binding the contract method 0x45a4276b.
//
// Solidity: function depositERC721(address _token, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_Il1erc721gateway *Il1erc721gatewayTransactor) DepositERC7210(opts *bind.TransactOpts, _token common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc721gateway.contract.Transact(opts, "depositERC7210", _token, _tokenId, _gasLimit)
}

// DepositERC7210 is a paid mutator transaction binding the contract method 0x45a4276b.
//
// Solidity: function depositERC721(address _token, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_Il1erc721gateway *Il1erc721gatewaySession) DepositERC7210(_token common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc721gateway.Contract.DepositERC7210(&_Il1erc721gateway.TransactOpts, _token, _tokenId, _gasLimit)
}

// DepositERC7210 is a paid mutator transaction binding the contract method 0x45a4276b.
//
// Solidity: function depositERC721(address _token, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_Il1erc721gateway *Il1erc721gatewayTransactorSession) DepositERC7210(_token common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc721gateway.Contract.DepositERC7210(&_Il1erc721gateway.TransactOpts, _token, _tokenId, _gasLimit)
}

// FinalizeBatchWithdrawERC721 is a paid mutator transaction binding the contract method 0x9f0a68b3.
//
// Solidity: function finalizeBatchWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds) returns()
func (_Il1erc721gateway *Il1erc721gatewayTransactor) FinalizeBatchWithdrawERC721(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Il1erc721gateway.contract.Transact(opts, "finalizeBatchWithdrawERC721", _l1Token, _l2Token, _from, _to, _tokenIds)
}

// FinalizeBatchWithdrawERC721 is a paid mutator transaction binding the contract method 0x9f0a68b3.
//
// Solidity: function finalizeBatchWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds) returns()
func (_Il1erc721gateway *Il1erc721gatewaySession) FinalizeBatchWithdrawERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Il1erc721gateway.Contract.FinalizeBatchWithdrawERC721(&_Il1erc721gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenIds)
}

// FinalizeBatchWithdrawERC721 is a paid mutator transaction binding the contract method 0x9f0a68b3.
//
// Solidity: function finalizeBatchWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds) returns()
func (_Il1erc721gateway *Il1erc721gatewayTransactorSession) FinalizeBatchWithdrawERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Il1erc721gateway.Contract.FinalizeBatchWithdrawERC721(&_Il1erc721gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenIds)
}

// FinalizeWithdrawERC721 is a paid mutator transaction binding the contract method 0xd606b4dc.
//
// Solidity: function finalizeWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId) returns()
func (_Il1erc721gateway *Il1erc721gatewayTransactor) FinalizeWithdrawERC721(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Il1erc721gateway.contract.Transact(opts, "finalizeWithdrawERC721", _l1Token, _l2Token, _from, _to, _tokenId)
}

// FinalizeWithdrawERC721 is a paid mutator transaction binding the contract method 0xd606b4dc.
//
// Solidity: function finalizeWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId) returns()
func (_Il1erc721gateway *Il1erc721gatewaySession) FinalizeWithdrawERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Il1erc721gateway.Contract.FinalizeWithdrawERC721(&_Il1erc721gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenId)
}

// FinalizeWithdrawERC721 is a paid mutator transaction binding the contract method 0xd606b4dc.
//
// Solidity: function finalizeWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId) returns()
func (_Il1erc721gateway *Il1erc721gatewayTransactorSession) FinalizeWithdrawERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Il1erc721gateway.Contract.FinalizeWithdrawERC721(&_Il1erc721gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenId)
}

// Il1erc721gatewayBatchDepositERC721Iterator is returned from FilterBatchDepositERC721 and is used to iterate over the raw logs and unpacked data for BatchDepositERC721 events raised by the Il1erc721gateway contract.
type Il1erc721gatewayBatchDepositERC721Iterator struct {
	Event *Il1erc721gatewayBatchDepositERC721 // Event containing the contract specifics and raw log

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
func (it *Il1erc721gatewayBatchDepositERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il1erc721gatewayBatchDepositERC721)
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
		it.Event = new(Il1erc721gatewayBatchDepositERC721)
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
func (it *Il1erc721gatewayBatchDepositERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il1erc721gatewayBatchDepositERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il1erc721gatewayBatchDepositERC721 represents a BatchDepositERC721 event raised by the Il1erc721gateway contract.
type Il1erc721gatewayBatchDepositERC721 struct {
	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBatchDepositERC721 is a free log retrieval operation binding the contract event 0xf05915e3b4fbd6f61b8b6f80b07f10e1cad039ccc7abe7c7fec115d038fe3dd6.
//
// Solidity: event BatchDepositERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds)
func (_Il1erc721gateway *Il1erc721gatewayFilterer) FilterBatchDepositERC721(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*Il1erc721gatewayBatchDepositERC721Iterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Il1erc721gateway.contract.FilterLogs(opts, "BatchDepositERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &Il1erc721gatewayBatchDepositERC721Iterator{contract: _Il1erc721gateway.contract, event: "BatchDepositERC721", logs: logs, sub: sub}, nil
}

// WatchBatchDepositERC721 is a free log subscription operation binding the contract event 0xf05915e3b4fbd6f61b8b6f80b07f10e1cad039ccc7abe7c7fec115d038fe3dd6.
//
// Solidity: event BatchDepositERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds)
func (_Il1erc721gateway *Il1erc721gatewayFilterer) WatchBatchDepositERC721(opts *bind.WatchOpts, sink chan<- *Il1erc721gatewayBatchDepositERC721, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Il1erc721gateway.contract.WatchLogs(opts, "BatchDepositERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il1erc721gatewayBatchDepositERC721)
				if err := _Il1erc721gateway.contract.UnpackLog(event, "BatchDepositERC721", log); err != nil {
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

// ParseBatchDepositERC721 is a log parse operation binding the contract event 0xf05915e3b4fbd6f61b8b6f80b07f10e1cad039ccc7abe7c7fec115d038fe3dd6.
//
// Solidity: event BatchDepositERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds)
func (_Il1erc721gateway *Il1erc721gatewayFilterer) ParseBatchDepositERC721(log types.Log) (*Il1erc721gatewayBatchDepositERC721, error) {
	event := new(Il1erc721gatewayBatchDepositERC721)
	if err := _Il1erc721gateway.contract.UnpackLog(event, "BatchDepositERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il1erc721gatewayBatchRefundERC721Iterator is returned from FilterBatchRefundERC721 and is used to iterate over the raw logs and unpacked data for BatchRefundERC721 events raised by the Il1erc721gateway contract.
type Il1erc721gatewayBatchRefundERC721Iterator struct {
	Event *Il1erc721gatewayBatchRefundERC721 // Event containing the contract specifics and raw log

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
func (it *Il1erc721gatewayBatchRefundERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il1erc721gatewayBatchRefundERC721)
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
		it.Event = new(Il1erc721gatewayBatchRefundERC721)
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
func (it *Il1erc721gatewayBatchRefundERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il1erc721gatewayBatchRefundERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il1erc721gatewayBatchRefundERC721 represents a BatchRefundERC721 event raised by the Il1erc721gateway contract.
type Il1erc721gatewayBatchRefundERC721 struct {
	Token     common.Address
	Recipient common.Address
	TokenIds  []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBatchRefundERC721 is a free log retrieval operation binding the contract event 0x998a3ef0a23771412ff48d871a2288502a89da39c5db04a2a66e5eb85586cc22.
//
// Solidity: event BatchRefundERC721(address indexed token, address indexed recipient, uint256[] tokenIds)
func (_Il1erc721gateway *Il1erc721gatewayFilterer) FilterBatchRefundERC721(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*Il1erc721gatewayBatchRefundERC721Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Il1erc721gateway.contract.FilterLogs(opts, "BatchRefundERC721", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &Il1erc721gatewayBatchRefundERC721Iterator{contract: _Il1erc721gateway.contract, event: "BatchRefundERC721", logs: logs, sub: sub}, nil
}

// WatchBatchRefundERC721 is a free log subscription operation binding the contract event 0x998a3ef0a23771412ff48d871a2288502a89da39c5db04a2a66e5eb85586cc22.
//
// Solidity: event BatchRefundERC721(address indexed token, address indexed recipient, uint256[] tokenIds)
func (_Il1erc721gateway *Il1erc721gatewayFilterer) WatchBatchRefundERC721(opts *bind.WatchOpts, sink chan<- *Il1erc721gatewayBatchRefundERC721, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Il1erc721gateway.contract.WatchLogs(opts, "BatchRefundERC721", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il1erc721gatewayBatchRefundERC721)
				if err := _Il1erc721gateway.contract.UnpackLog(event, "BatchRefundERC721", log); err != nil {
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

// ParseBatchRefundERC721 is a log parse operation binding the contract event 0x998a3ef0a23771412ff48d871a2288502a89da39c5db04a2a66e5eb85586cc22.
//
// Solidity: event BatchRefundERC721(address indexed token, address indexed recipient, uint256[] tokenIds)
func (_Il1erc721gateway *Il1erc721gatewayFilterer) ParseBatchRefundERC721(log types.Log) (*Il1erc721gatewayBatchRefundERC721, error) {
	event := new(Il1erc721gatewayBatchRefundERC721)
	if err := _Il1erc721gateway.contract.UnpackLog(event, "BatchRefundERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il1erc721gatewayDepositERC721Iterator is returned from FilterDepositERC721 and is used to iterate over the raw logs and unpacked data for DepositERC721 events raised by the Il1erc721gateway contract.
type Il1erc721gatewayDepositERC721Iterator struct {
	Event *Il1erc721gatewayDepositERC721 // Event containing the contract specifics and raw log

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
func (it *Il1erc721gatewayDepositERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il1erc721gatewayDepositERC721)
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
		it.Event = new(Il1erc721gatewayDepositERC721)
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
func (it *Il1erc721gatewayDepositERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il1erc721gatewayDepositERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il1erc721gatewayDepositERC721 represents a DepositERC721 event raised by the Il1erc721gateway contract.
type Il1erc721gatewayDepositERC721 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositERC721 is a free log retrieval operation binding the contract event 0xfc1d17c06ff1e4678321cc30660a73f3f1436df8195108a288d3159a961febec.
//
// Solidity: event DepositERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId)
func (_Il1erc721gateway *Il1erc721gatewayFilterer) FilterDepositERC721(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*Il1erc721gatewayDepositERC721Iterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Il1erc721gateway.contract.FilterLogs(opts, "DepositERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &Il1erc721gatewayDepositERC721Iterator{contract: _Il1erc721gateway.contract, event: "DepositERC721", logs: logs, sub: sub}, nil
}

// WatchDepositERC721 is a free log subscription operation binding the contract event 0xfc1d17c06ff1e4678321cc30660a73f3f1436df8195108a288d3159a961febec.
//
// Solidity: event DepositERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId)
func (_Il1erc721gateway *Il1erc721gatewayFilterer) WatchDepositERC721(opts *bind.WatchOpts, sink chan<- *Il1erc721gatewayDepositERC721, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Il1erc721gateway.contract.WatchLogs(opts, "DepositERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il1erc721gatewayDepositERC721)
				if err := _Il1erc721gateway.contract.UnpackLog(event, "DepositERC721", log); err != nil {
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

// ParseDepositERC721 is a log parse operation binding the contract event 0xfc1d17c06ff1e4678321cc30660a73f3f1436df8195108a288d3159a961febec.
//
// Solidity: event DepositERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId)
func (_Il1erc721gateway *Il1erc721gatewayFilterer) ParseDepositERC721(log types.Log) (*Il1erc721gatewayDepositERC721, error) {
	event := new(Il1erc721gatewayDepositERC721)
	if err := _Il1erc721gateway.contract.UnpackLog(event, "DepositERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il1erc721gatewayFinalizeBatchWithdrawERC721Iterator is returned from FilterFinalizeBatchWithdrawERC721 and is used to iterate over the raw logs and unpacked data for FinalizeBatchWithdrawERC721 events raised by the Il1erc721gateway contract.
type Il1erc721gatewayFinalizeBatchWithdrawERC721Iterator struct {
	Event *Il1erc721gatewayFinalizeBatchWithdrawERC721 // Event containing the contract specifics and raw log

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
func (it *Il1erc721gatewayFinalizeBatchWithdrawERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il1erc721gatewayFinalizeBatchWithdrawERC721)
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
		it.Event = new(Il1erc721gatewayFinalizeBatchWithdrawERC721)
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
func (it *Il1erc721gatewayFinalizeBatchWithdrawERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il1erc721gatewayFinalizeBatchWithdrawERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il1erc721gatewayFinalizeBatchWithdrawERC721 represents a FinalizeBatchWithdrawERC721 event raised by the Il1erc721gateway contract.
type Il1erc721gatewayFinalizeBatchWithdrawERC721 struct {
	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFinalizeBatchWithdrawERC721 is a free log retrieval operation binding the contract event 0x9b8e51c8f180115b421b26c9042287d6bf95e0ce9c0c5434784e2af3d0b9de7d.
//
// Solidity: event FinalizeBatchWithdrawERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds)
func (_Il1erc721gateway *Il1erc721gatewayFilterer) FilterFinalizeBatchWithdrawERC721(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*Il1erc721gatewayFinalizeBatchWithdrawERC721Iterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Il1erc721gateway.contract.FilterLogs(opts, "FinalizeBatchWithdrawERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &Il1erc721gatewayFinalizeBatchWithdrawERC721Iterator{contract: _Il1erc721gateway.contract, event: "FinalizeBatchWithdrawERC721", logs: logs, sub: sub}, nil
}

// WatchFinalizeBatchWithdrawERC721 is a free log subscription operation binding the contract event 0x9b8e51c8f180115b421b26c9042287d6bf95e0ce9c0c5434784e2af3d0b9de7d.
//
// Solidity: event FinalizeBatchWithdrawERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds)
func (_Il1erc721gateway *Il1erc721gatewayFilterer) WatchFinalizeBatchWithdrawERC721(opts *bind.WatchOpts, sink chan<- *Il1erc721gatewayFinalizeBatchWithdrawERC721, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Il1erc721gateway.contract.WatchLogs(opts, "FinalizeBatchWithdrawERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il1erc721gatewayFinalizeBatchWithdrawERC721)
				if err := _Il1erc721gateway.contract.UnpackLog(event, "FinalizeBatchWithdrawERC721", log); err != nil {
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

// ParseFinalizeBatchWithdrawERC721 is a log parse operation binding the contract event 0x9b8e51c8f180115b421b26c9042287d6bf95e0ce9c0c5434784e2af3d0b9de7d.
//
// Solidity: event FinalizeBatchWithdrawERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds)
func (_Il1erc721gateway *Il1erc721gatewayFilterer) ParseFinalizeBatchWithdrawERC721(log types.Log) (*Il1erc721gatewayFinalizeBatchWithdrawERC721, error) {
	event := new(Il1erc721gatewayFinalizeBatchWithdrawERC721)
	if err := _Il1erc721gateway.contract.UnpackLog(event, "FinalizeBatchWithdrawERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il1erc721gatewayFinalizeWithdrawERC721Iterator is returned from FilterFinalizeWithdrawERC721 and is used to iterate over the raw logs and unpacked data for FinalizeWithdrawERC721 events raised by the Il1erc721gateway contract.
type Il1erc721gatewayFinalizeWithdrawERC721Iterator struct {
	Event *Il1erc721gatewayFinalizeWithdrawERC721 // Event containing the contract specifics and raw log

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
func (it *Il1erc721gatewayFinalizeWithdrawERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il1erc721gatewayFinalizeWithdrawERC721)
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
		it.Event = new(Il1erc721gatewayFinalizeWithdrawERC721)
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
func (it *Il1erc721gatewayFinalizeWithdrawERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il1erc721gatewayFinalizeWithdrawERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il1erc721gatewayFinalizeWithdrawERC721 represents a FinalizeWithdrawERC721 event raised by the Il1erc721gateway contract.
type Il1erc721gatewayFinalizeWithdrawERC721 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinalizeWithdrawERC721 is a free log retrieval operation binding the contract event 0xacdbfefc030b5ccccd5f60ca6d9ca371c6d6d6956fe16ebe10f81920198206e9.
//
// Solidity: event FinalizeWithdrawERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId)
func (_Il1erc721gateway *Il1erc721gatewayFilterer) FilterFinalizeWithdrawERC721(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*Il1erc721gatewayFinalizeWithdrawERC721Iterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Il1erc721gateway.contract.FilterLogs(opts, "FinalizeWithdrawERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &Il1erc721gatewayFinalizeWithdrawERC721Iterator{contract: _Il1erc721gateway.contract, event: "FinalizeWithdrawERC721", logs: logs, sub: sub}, nil
}

// WatchFinalizeWithdrawERC721 is a free log subscription operation binding the contract event 0xacdbfefc030b5ccccd5f60ca6d9ca371c6d6d6956fe16ebe10f81920198206e9.
//
// Solidity: event FinalizeWithdrawERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId)
func (_Il1erc721gateway *Il1erc721gatewayFilterer) WatchFinalizeWithdrawERC721(opts *bind.WatchOpts, sink chan<- *Il1erc721gatewayFinalizeWithdrawERC721, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Il1erc721gateway.contract.WatchLogs(opts, "FinalizeWithdrawERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il1erc721gatewayFinalizeWithdrawERC721)
				if err := _Il1erc721gateway.contract.UnpackLog(event, "FinalizeWithdrawERC721", log); err != nil {
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

// ParseFinalizeWithdrawERC721 is a log parse operation binding the contract event 0xacdbfefc030b5ccccd5f60ca6d9ca371c6d6d6956fe16ebe10f81920198206e9.
//
// Solidity: event FinalizeWithdrawERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId)
func (_Il1erc721gateway *Il1erc721gatewayFilterer) ParseFinalizeWithdrawERC721(log types.Log) (*Il1erc721gatewayFinalizeWithdrawERC721, error) {
	event := new(Il1erc721gatewayFinalizeWithdrawERC721)
	if err := _Il1erc721gateway.contract.UnpackLog(event, "FinalizeWithdrawERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il1erc721gatewayRefundERC721Iterator is returned from FilterRefundERC721 and is used to iterate over the raw logs and unpacked data for RefundERC721 events raised by the Il1erc721gateway contract.
type Il1erc721gatewayRefundERC721Iterator struct {
	Event *Il1erc721gatewayRefundERC721 // Event containing the contract specifics and raw log

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
func (it *Il1erc721gatewayRefundERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il1erc721gatewayRefundERC721)
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
		it.Event = new(Il1erc721gatewayRefundERC721)
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
func (it *Il1erc721gatewayRefundERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il1erc721gatewayRefundERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il1erc721gatewayRefundERC721 represents a RefundERC721 event raised by the Il1erc721gateway contract.
type Il1erc721gatewayRefundERC721 struct {
	Token     common.Address
	Recipient common.Address
	TokenId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundERC721 is a free log retrieval operation binding the contract event 0xb9a838365634e4fb87a9333edf0ea86f82836e361b311a125aefd14135581208.
//
// Solidity: event RefundERC721(address indexed token, address indexed recipient, uint256 tokenId)
func (_Il1erc721gateway *Il1erc721gatewayFilterer) FilterRefundERC721(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*Il1erc721gatewayRefundERC721Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Il1erc721gateway.contract.FilterLogs(opts, "RefundERC721", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &Il1erc721gatewayRefundERC721Iterator{contract: _Il1erc721gateway.contract, event: "RefundERC721", logs: logs, sub: sub}, nil
}

// WatchRefundERC721 is a free log subscription operation binding the contract event 0xb9a838365634e4fb87a9333edf0ea86f82836e361b311a125aefd14135581208.
//
// Solidity: event RefundERC721(address indexed token, address indexed recipient, uint256 tokenId)
func (_Il1erc721gateway *Il1erc721gatewayFilterer) WatchRefundERC721(opts *bind.WatchOpts, sink chan<- *Il1erc721gatewayRefundERC721, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Il1erc721gateway.contract.WatchLogs(opts, "RefundERC721", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il1erc721gatewayRefundERC721)
				if err := _Il1erc721gateway.contract.UnpackLog(event, "RefundERC721", log); err != nil {
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

// ParseRefundERC721 is a log parse operation binding the contract event 0xb9a838365634e4fb87a9333edf0ea86f82836e361b311a125aefd14135581208.
//
// Solidity: event RefundERC721(address indexed token, address indexed recipient, uint256 tokenId)
func (_Il1erc721gateway *Il1erc721gatewayFilterer) ParseRefundERC721(log types.Log) (*Il1erc721gatewayRefundERC721, error) {
	event := new(Il1erc721gatewayRefundERC721)
	if err := _Il1erc721gateway.contract.UnpackLog(event, "RefundERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
