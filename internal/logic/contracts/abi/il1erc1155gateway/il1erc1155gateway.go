// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package il1erc1155gateway

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

// Il1erc1155gatewayMetaData contains all meta data concerning the Il1erc1155gateway contract.
var Il1erc1155gatewayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"BatchDepositERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"BatchRefundERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DepositERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"FinalizeBatchWithdrawERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"FinalizeWithdrawERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundERC1155\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"batchDepositERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"batchDepositERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"finalizeBatchWithdrawERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"finalizeWithdrawERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Il1erc1155gatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use Il1erc1155gatewayMetaData.ABI instead.
var Il1erc1155gatewayABI = Il1erc1155gatewayMetaData.ABI

// Il1erc1155gateway is an auto generated Go binding around an Ethereum contract.
type Il1erc1155gateway struct {
	Il1erc1155gatewayCaller     // Read-only binding to the contract
	Il1erc1155gatewayTransactor // Write-only binding to the contract
	Il1erc1155gatewayFilterer   // Log filterer for contract events
}

// Il1erc1155gatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type Il1erc1155gatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il1erc1155gatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Il1erc1155gatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il1erc1155gatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Il1erc1155gatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Il1erc1155gatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Il1erc1155gatewaySession struct {
	Contract     *Il1erc1155gateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Il1erc1155gatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Il1erc1155gatewayCallerSession struct {
	Contract *Il1erc1155gatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// Il1erc1155gatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Il1erc1155gatewayTransactorSession struct {
	Contract     *Il1erc1155gatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// Il1erc1155gatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type Il1erc1155gatewayRaw struct {
	Contract *Il1erc1155gateway // Generic contract binding to access the raw methods on
}

// Il1erc1155gatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Il1erc1155gatewayCallerRaw struct {
	Contract *Il1erc1155gatewayCaller // Generic read-only contract binding to access the raw methods on
}

// Il1erc1155gatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Il1erc1155gatewayTransactorRaw struct {
	Contract *Il1erc1155gatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIl1erc1155gateway creates a new instance of Il1erc1155gateway, bound to a specific deployed contract.
func NewIl1erc1155gateway(address common.Address, backend bind.ContractBackend) (*Il1erc1155gateway, error) {
	contract, err := bindIl1erc1155gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Il1erc1155gateway{Il1erc1155gatewayCaller: Il1erc1155gatewayCaller{contract: contract}, Il1erc1155gatewayTransactor: Il1erc1155gatewayTransactor{contract: contract}, Il1erc1155gatewayFilterer: Il1erc1155gatewayFilterer{contract: contract}}, nil
}

// NewIl1erc1155gatewayCaller creates a new read-only instance of Il1erc1155gateway, bound to a specific deployed contract.
func NewIl1erc1155gatewayCaller(address common.Address, caller bind.ContractCaller) (*Il1erc1155gatewayCaller, error) {
	contract, err := bindIl1erc1155gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Il1erc1155gatewayCaller{contract: contract}, nil
}

// NewIl1erc1155gatewayTransactor creates a new write-only instance of Il1erc1155gateway, bound to a specific deployed contract.
func NewIl1erc1155gatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*Il1erc1155gatewayTransactor, error) {
	contract, err := bindIl1erc1155gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Il1erc1155gatewayTransactor{contract: contract}, nil
}

// NewIl1erc1155gatewayFilterer creates a new log filterer instance of Il1erc1155gateway, bound to a specific deployed contract.
func NewIl1erc1155gatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*Il1erc1155gatewayFilterer, error) {
	contract, err := bindIl1erc1155gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Il1erc1155gatewayFilterer{contract: contract}, nil
}

// bindIl1erc1155gateway binds a generic wrapper to an already deployed contract.
func bindIl1erc1155gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Il1erc1155gatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Il1erc1155gateway *Il1erc1155gatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Il1erc1155gateway.Contract.Il1erc1155gatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Il1erc1155gateway *Il1erc1155gatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Il1erc1155gateway.Contract.Il1erc1155gatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Il1erc1155gateway *Il1erc1155gatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Il1erc1155gateway.Contract.Il1erc1155gatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Il1erc1155gateway *Il1erc1155gatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Il1erc1155gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Il1erc1155gateway *Il1erc1155gatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Il1erc1155gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Il1erc1155gateway *Il1erc1155gatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Il1erc1155gateway.Contract.contract.Transact(opts, method, params...)
}

// BatchDepositERC1155 is a paid mutator transaction binding the contract method 0x5ee8e74c.
//
// Solidity: function batchDepositERC1155(address _token, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_Il1erc1155gateway *Il1erc1155gatewayTransactor) BatchDepositERC1155(opts *bind.TransactOpts, _token common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc1155gateway.contract.Transact(opts, "batchDepositERC1155", _token, _tokenIds, _amounts, _gasLimit)
}

// BatchDepositERC1155 is a paid mutator transaction binding the contract method 0x5ee8e74c.
//
// Solidity: function batchDepositERC1155(address _token, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_Il1erc1155gateway *Il1erc1155gatewaySession) BatchDepositERC1155(_token common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc1155gateway.Contract.BatchDepositERC1155(&_Il1erc1155gateway.TransactOpts, _token, _tokenIds, _amounts, _gasLimit)
}

// BatchDepositERC1155 is a paid mutator transaction binding the contract method 0x5ee8e74c.
//
// Solidity: function batchDepositERC1155(address _token, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_Il1erc1155gateway *Il1erc1155gatewayTransactorSession) BatchDepositERC1155(_token common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc1155gateway.Contract.BatchDepositERC1155(&_Il1erc1155gateway.TransactOpts, _token, _tokenIds, _amounts, _gasLimit)
}

// BatchDepositERC11550 is a paid mutator transaction binding the contract method 0xc99dac9b.
//
// Solidity: function batchDepositERC1155(address _token, address _to, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_Il1erc1155gateway *Il1erc1155gatewayTransactor) BatchDepositERC11550(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc1155gateway.contract.Transact(opts, "batchDepositERC11550", _token, _to, _tokenIds, _amounts, _gasLimit)
}

// BatchDepositERC11550 is a paid mutator transaction binding the contract method 0xc99dac9b.
//
// Solidity: function batchDepositERC1155(address _token, address _to, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_Il1erc1155gateway *Il1erc1155gatewaySession) BatchDepositERC11550(_token common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc1155gateway.Contract.BatchDepositERC11550(&_Il1erc1155gateway.TransactOpts, _token, _to, _tokenIds, _amounts, _gasLimit)
}

// BatchDepositERC11550 is a paid mutator transaction binding the contract method 0xc99dac9b.
//
// Solidity: function batchDepositERC1155(address _token, address _to, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_Il1erc1155gateway *Il1erc1155gatewayTransactorSession) BatchDepositERC11550(_token common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc1155gateway.Contract.BatchDepositERC11550(&_Il1erc1155gateway.TransactOpts, _token, _to, _tokenIds, _amounts, _gasLimit)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0xa901cf8a.
//
// Solidity: function depositERC1155(address _token, address _to, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_Il1erc1155gateway *Il1erc1155gatewayTransactor) DepositERC1155(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc1155gateway.contract.Transact(opts, "depositERC1155", _token, _to, _tokenId, _amount, _gasLimit)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0xa901cf8a.
//
// Solidity: function depositERC1155(address _token, address _to, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_Il1erc1155gateway *Il1erc1155gatewaySession) DepositERC1155(_token common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc1155gateway.Contract.DepositERC1155(&_Il1erc1155gateway.TransactOpts, _token, _to, _tokenId, _amount, _gasLimit)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0xa901cf8a.
//
// Solidity: function depositERC1155(address _token, address _to, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_Il1erc1155gateway *Il1erc1155gatewayTransactorSession) DepositERC1155(_token common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc1155gateway.Contract.DepositERC1155(&_Il1erc1155gateway.TransactOpts, _token, _to, _tokenId, _amount, _gasLimit)
}

// DepositERC11550 is a paid mutator transaction binding the contract method 0xf998fe9d.
//
// Solidity: function depositERC1155(address _token, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_Il1erc1155gateway *Il1erc1155gatewayTransactor) DepositERC11550(opts *bind.TransactOpts, _token common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc1155gateway.contract.Transact(opts, "depositERC11550", _token, _tokenId, _amount, _gasLimit)
}

// DepositERC11550 is a paid mutator transaction binding the contract method 0xf998fe9d.
//
// Solidity: function depositERC1155(address _token, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_Il1erc1155gateway *Il1erc1155gatewaySession) DepositERC11550(_token common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc1155gateway.Contract.DepositERC11550(&_Il1erc1155gateway.TransactOpts, _token, _tokenId, _amount, _gasLimit)
}

// DepositERC11550 is a paid mutator transaction binding the contract method 0xf998fe9d.
//
// Solidity: function depositERC1155(address _token, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_Il1erc1155gateway *Il1erc1155gatewayTransactorSession) DepositERC11550(_token common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Il1erc1155gateway.Contract.DepositERC11550(&_Il1erc1155gateway.TransactOpts, _token, _tokenId, _amount, _gasLimit)
}

// FinalizeBatchWithdrawERC1155 is a paid mutator transaction binding the contract method 0xf92748d3.
//
// Solidity: function finalizeBatchWithdrawERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds, uint256[] _amounts) returns()
func (_Il1erc1155gateway *Il1erc1155gatewayTransactor) FinalizeBatchWithdrawERC1155(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Il1erc1155gateway.contract.Transact(opts, "finalizeBatchWithdrawERC1155", _l1Token, _l2Token, _from, _to, _tokenIds, _amounts)
}

// FinalizeBatchWithdrawERC1155 is a paid mutator transaction binding the contract method 0xf92748d3.
//
// Solidity: function finalizeBatchWithdrawERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds, uint256[] _amounts) returns()
func (_Il1erc1155gateway *Il1erc1155gatewaySession) FinalizeBatchWithdrawERC1155(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Il1erc1155gateway.Contract.FinalizeBatchWithdrawERC1155(&_Il1erc1155gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenIds, _amounts)
}

// FinalizeBatchWithdrawERC1155 is a paid mutator transaction binding the contract method 0xf92748d3.
//
// Solidity: function finalizeBatchWithdrawERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds, uint256[] _amounts) returns()
func (_Il1erc1155gateway *Il1erc1155gatewayTransactorSession) FinalizeBatchWithdrawERC1155(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Il1erc1155gateway.Contract.FinalizeBatchWithdrawERC1155(&_Il1erc1155gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenIds, _amounts)
}

// FinalizeWithdrawERC1155 is a paid mutator transaction binding the contract method 0x730608b3.
//
// Solidity: function finalizeWithdrawERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId, uint256 _amount) returns()
func (_Il1erc1155gateway *Il1erc1155gatewayTransactor) FinalizeWithdrawERC1155(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Il1erc1155gateway.contract.Transact(opts, "finalizeWithdrawERC1155", _l1Token, _l2Token, _from, _to, _tokenId, _amount)
}

// FinalizeWithdrawERC1155 is a paid mutator transaction binding the contract method 0x730608b3.
//
// Solidity: function finalizeWithdrawERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId, uint256 _amount) returns()
func (_Il1erc1155gateway *Il1erc1155gatewaySession) FinalizeWithdrawERC1155(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Il1erc1155gateway.Contract.FinalizeWithdrawERC1155(&_Il1erc1155gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenId, _amount)
}

// FinalizeWithdrawERC1155 is a paid mutator transaction binding the contract method 0x730608b3.
//
// Solidity: function finalizeWithdrawERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId, uint256 _amount) returns()
func (_Il1erc1155gateway *Il1erc1155gatewayTransactorSession) FinalizeWithdrawERC1155(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Il1erc1155gateway.Contract.FinalizeWithdrawERC1155(&_Il1erc1155gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenId, _amount)
}

// Il1erc1155gatewayBatchDepositERC1155Iterator is returned from FilterBatchDepositERC1155 and is used to iterate over the raw logs and unpacked data for BatchDepositERC1155 events raised by the Il1erc1155gateway contract.
type Il1erc1155gatewayBatchDepositERC1155Iterator struct {
	Event *Il1erc1155gatewayBatchDepositERC1155 // Event containing the contract specifics and raw log

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
func (it *Il1erc1155gatewayBatchDepositERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il1erc1155gatewayBatchDepositERC1155)
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
		it.Event = new(Il1erc1155gatewayBatchDepositERC1155)
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
func (it *Il1erc1155gatewayBatchDepositERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il1erc1155gatewayBatchDepositERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il1erc1155gatewayBatchDepositERC1155 represents a BatchDepositERC1155 event raised by the Il1erc1155gateway contract.
type Il1erc1155gatewayBatchDepositERC1155 struct {
	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Amounts  []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBatchDepositERC1155 is a free log retrieval operation binding the contract event 0x743f65db61a23bc629915d35e22af5cf13478a8b3dbd154d3e5db0149509756d.
//
// Solidity: event BatchDepositERC1155(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds, uint256[] _amounts)
func (_Il1erc1155gateway *Il1erc1155gatewayFilterer) FilterBatchDepositERC1155(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*Il1erc1155gatewayBatchDepositERC1155Iterator, error) {

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

	logs, sub, err := _Il1erc1155gateway.contract.FilterLogs(opts, "BatchDepositERC1155", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &Il1erc1155gatewayBatchDepositERC1155Iterator{contract: _Il1erc1155gateway.contract, event: "BatchDepositERC1155", logs: logs, sub: sub}, nil
}

// WatchBatchDepositERC1155 is a free log subscription operation binding the contract event 0x743f65db61a23bc629915d35e22af5cf13478a8b3dbd154d3e5db0149509756d.
//
// Solidity: event BatchDepositERC1155(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds, uint256[] _amounts)
func (_Il1erc1155gateway *Il1erc1155gatewayFilterer) WatchBatchDepositERC1155(opts *bind.WatchOpts, sink chan<- *Il1erc1155gatewayBatchDepositERC1155, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Il1erc1155gateway.contract.WatchLogs(opts, "BatchDepositERC1155", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il1erc1155gatewayBatchDepositERC1155)
				if err := _Il1erc1155gateway.contract.UnpackLog(event, "BatchDepositERC1155", log); err != nil {
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

// ParseBatchDepositERC1155 is a log parse operation binding the contract event 0x743f65db61a23bc629915d35e22af5cf13478a8b3dbd154d3e5db0149509756d.
//
// Solidity: event BatchDepositERC1155(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds, uint256[] _amounts)
func (_Il1erc1155gateway *Il1erc1155gatewayFilterer) ParseBatchDepositERC1155(log types.Log) (*Il1erc1155gatewayBatchDepositERC1155, error) {
	event := new(Il1erc1155gatewayBatchDepositERC1155)
	if err := _Il1erc1155gateway.contract.UnpackLog(event, "BatchDepositERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il1erc1155gatewayBatchRefundERC1155Iterator is returned from FilterBatchRefundERC1155 and is used to iterate over the raw logs and unpacked data for BatchRefundERC1155 events raised by the Il1erc1155gateway contract.
type Il1erc1155gatewayBatchRefundERC1155Iterator struct {
	Event *Il1erc1155gatewayBatchRefundERC1155 // Event containing the contract specifics and raw log

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
func (it *Il1erc1155gatewayBatchRefundERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il1erc1155gatewayBatchRefundERC1155)
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
		it.Event = new(Il1erc1155gatewayBatchRefundERC1155)
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
func (it *Il1erc1155gatewayBatchRefundERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il1erc1155gatewayBatchRefundERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il1erc1155gatewayBatchRefundERC1155 represents a BatchRefundERC1155 event raised by the Il1erc1155gateway contract.
type Il1erc1155gatewayBatchRefundERC1155 struct {
	Token     common.Address
	Recipient common.Address
	TokenIds  []*big.Int
	Amounts   []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBatchRefundERC1155 is a free log retrieval operation binding the contract event 0xe198c04cbd4522ed7825c7e6ab1ae33fdaf6ab3565c4a3fb4c0cf24338f306e6.
//
// Solidity: event BatchRefundERC1155(address indexed token, address indexed recipient, uint256[] tokenIds, uint256[] amounts)
func (_Il1erc1155gateway *Il1erc1155gatewayFilterer) FilterBatchRefundERC1155(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*Il1erc1155gatewayBatchRefundERC1155Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Il1erc1155gateway.contract.FilterLogs(opts, "BatchRefundERC1155", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &Il1erc1155gatewayBatchRefundERC1155Iterator{contract: _Il1erc1155gateway.contract, event: "BatchRefundERC1155", logs: logs, sub: sub}, nil
}

// WatchBatchRefundERC1155 is a free log subscription operation binding the contract event 0xe198c04cbd4522ed7825c7e6ab1ae33fdaf6ab3565c4a3fb4c0cf24338f306e6.
//
// Solidity: event BatchRefundERC1155(address indexed token, address indexed recipient, uint256[] tokenIds, uint256[] amounts)
func (_Il1erc1155gateway *Il1erc1155gatewayFilterer) WatchBatchRefundERC1155(opts *bind.WatchOpts, sink chan<- *Il1erc1155gatewayBatchRefundERC1155, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Il1erc1155gateway.contract.WatchLogs(opts, "BatchRefundERC1155", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il1erc1155gatewayBatchRefundERC1155)
				if err := _Il1erc1155gateway.contract.UnpackLog(event, "BatchRefundERC1155", log); err != nil {
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

// ParseBatchRefundERC1155 is a log parse operation binding the contract event 0xe198c04cbd4522ed7825c7e6ab1ae33fdaf6ab3565c4a3fb4c0cf24338f306e6.
//
// Solidity: event BatchRefundERC1155(address indexed token, address indexed recipient, uint256[] tokenIds, uint256[] amounts)
func (_Il1erc1155gateway *Il1erc1155gatewayFilterer) ParseBatchRefundERC1155(log types.Log) (*Il1erc1155gatewayBatchRefundERC1155, error) {
	event := new(Il1erc1155gatewayBatchRefundERC1155)
	if err := _Il1erc1155gateway.contract.UnpackLog(event, "BatchRefundERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il1erc1155gatewayDepositERC1155Iterator is returned from FilterDepositERC1155 and is used to iterate over the raw logs and unpacked data for DepositERC1155 events raised by the Il1erc1155gateway contract.
type Il1erc1155gatewayDepositERC1155Iterator struct {
	Event *Il1erc1155gatewayDepositERC1155 // Event containing the contract specifics and raw log

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
func (it *Il1erc1155gatewayDepositERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il1erc1155gatewayDepositERC1155)
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
		it.Event = new(Il1erc1155gatewayDepositERC1155)
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
func (it *Il1erc1155gatewayDepositERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il1erc1155gatewayDepositERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il1erc1155gatewayDepositERC1155 represents a DepositERC1155 event raised by the Il1erc1155gateway contract.
type Il1erc1155gatewayDepositERC1155 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositERC1155 is a free log retrieval operation binding the contract event 0x7f6552b688fa94306ca59e44dd4454ff550542445a3f1cb39b8c768be6f5c08a.
//
// Solidity: event DepositERC1155(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId, uint256 _amount)
func (_Il1erc1155gateway *Il1erc1155gatewayFilterer) FilterDepositERC1155(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*Il1erc1155gatewayDepositERC1155Iterator, error) {

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

	logs, sub, err := _Il1erc1155gateway.contract.FilterLogs(opts, "DepositERC1155", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &Il1erc1155gatewayDepositERC1155Iterator{contract: _Il1erc1155gateway.contract, event: "DepositERC1155", logs: logs, sub: sub}, nil
}

// WatchDepositERC1155 is a free log subscription operation binding the contract event 0x7f6552b688fa94306ca59e44dd4454ff550542445a3f1cb39b8c768be6f5c08a.
//
// Solidity: event DepositERC1155(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId, uint256 _amount)
func (_Il1erc1155gateway *Il1erc1155gatewayFilterer) WatchDepositERC1155(opts *bind.WatchOpts, sink chan<- *Il1erc1155gatewayDepositERC1155, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Il1erc1155gateway.contract.WatchLogs(opts, "DepositERC1155", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il1erc1155gatewayDepositERC1155)
				if err := _Il1erc1155gateway.contract.UnpackLog(event, "DepositERC1155", log); err != nil {
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

// ParseDepositERC1155 is a log parse operation binding the contract event 0x7f6552b688fa94306ca59e44dd4454ff550542445a3f1cb39b8c768be6f5c08a.
//
// Solidity: event DepositERC1155(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId, uint256 _amount)
func (_Il1erc1155gateway *Il1erc1155gatewayFilterer) ParseDepositERC1155(log types.Log) (*Il1erc1155gatewayDepositERC1155, error) {
	event := new(Il1erc1155gatewayDepositERC1155)
	if err := _Il1erc1155gateway.contract.UnpackLog(event, "DepositERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il1erc1155gatewayFinalizeBatchWithdrawERC1155Iterator is returned from FilterFinalizeBatchWithdrawERC1155 and is used to iterate over the raw logs and unpacked data for FinalizeBatchWithdrawERC1155 events raised by the Il1erc1155gateway contract.
type Il1erc1155gatewayFinalizeBatchWithdrawERC1155Iterator struct {
	Event *Il1erc1155gatewayFinalizeBatchWithdrawERC1155 // Event containing the contract specifics and raw log

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
func (it *Il1erc1155gatewayFinalizeBatchWithdrawERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il1erc1155gatewayFinalizeBatchWithdrawERC1155)
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
		it.Event = new(Il1erc1155gatewayFinalizeBatchWithdrawERC1155)
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
func (it *Il1erc1155gatewayFinalizeBatchWithdrawERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il1erc1155gatewayFinalizeBatchWithdrawERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il1erc1155gatewayFinalizeBatchWithdrawERC1155 represents a FinalizeBatchWithdrawERC1155 event raised by the Il1erc1155gateway contract.
type Il1erc1155gatewayFinalizeBatchWithdrawERC1155 struct {
	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Amounts  []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFinalizeBatchWithdrawERC1155 is a free log retrieval operation binding the contract event 0x45294b6ad6ad2408cc3ee9a37203aa1b0480616667a97b157c52ac9294cbc258.
//
// Solidity: event FinalizeBatchWithdrawERC1155(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds, uint256[] _amounts)
func (_Il1erc1155gateway *Il1erc1155gatewayFilterer) FilterFinalizeBatchWithdrawERC1155(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*Il1erc1155gatewayFinalizeBatchWithdrawERC1155Iterator, error) {

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

	logs, sub, err := _Il1erc1155gateway.contract.FilterLogs(opts, "FinalizeBatchWithdrawERC1155", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &Il1erc1155gatewayFinalizeBatchWithdrawERC1155Iterator{contract: _Il1erc1155gateway.contract, event: "FinalizeBatchWithdrawERC1155", logs: logs, sub: sub}, nil
}

// WatchFinalizeBatchWithdrawERC1155 is a free log subscription operation binding the contract event 0x45294b6ad6ad2408cc3ee9a37203aa1b0480616667a97b157c52ac9294cbc258.
//
// Solidity: event FinalizeBatchWithdrawERC1155(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds, uint256[] _amounts)
func (_Il1erc1155gateway *Il1erc1155gatewayFilterer) WatchFinalizeBatchWithdrawERC1155(opts *bind.WatchOpts, sink chan<- *Il1erc1155gatewayFinalizeBatchWithdrawERC1155, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Il1erc1155gateway.contract.WatchLogs(opts, "FinalizeBatchWithdrawERC1155", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il1erc1155gatewayFinalizeBatchWithdrawERC1155)
				if err := _Il1erc1155gateway.contract.UnpackLog(event, "FinalizeBatchWithdrawERC1155", log); err != nil {
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

// ParseFinalizeBatchWithdrawERC1155 is a log parse operation binding the contract event 0x45294b6ad6ad2408cc3ee9a37203aa1b0480616667a97b157c52ac9294cbc258.
//
// Solidity: event FinalizeBatchWithdrawERC1155(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds, uint256[] _amounts)
func (_Il1erc1155gateway *Il1erc1155gatewayFilterer) ParseFinalizeBatchWithdrawERC1155(log types.Log) (*Il1erc1155gatewayFinalizeBatchWithdrawERC1155, error) {
	event := new(Il1erc1155gatewayFinalizeBatchWithdrawERC1155)
	if err := _Il1erc1155gateway.contract.UnpackLog(event, "FinalizeBatchWithdrawERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il1erc1155gatewayFinalizeWithdrawERC1155Iterator is returned from FilterFinalizeWithdrawERC1155 and is used to iterate over the raw logs and unpacked data for FinalizeWithdrawERC1155 events raised by the Il1erc1155gateway contract.
type Il1erc1155gatewayFinalizeWithdrawERC1155Iterator struct {
	Event *Il1erc1155gatewayFinalizeWithdrawERC1155 // Event containing the contract specifics and raw log

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
func (it *Il1erc1155gatewayFinalizeWithdrawERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il1erc1155gatewayFinalizeWithdrawERC1155)
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
		it.Event = new(Il1erc1155gatewayFinalizeWithdrawERC1155)
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
func (it *Il1erc1155gatewayFinalizeWithdrawERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il1erc1155gatewayFinalizeWithdrawERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il1erc1155gatewayFinalizeWithdrawERC1155 represents a FinalizeWithdrawERC1155 event raised by the Il1erc1155gateway contract.
type Il1erc1155gatewayFinalizeWithdrawERC1155 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinalizeWithdrawERC1155 is a free log retrieval operation binding the contract event 0xfcc2841e9e72e6d610944e1b668912e92d5df94003055dbe06d615ba8d9efad4.
//
// Solidity: event FinalizeWithdrawERC1155(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId, uint256 _amount)
func (_Il1erc1155gateway *Il1erc1155gatewayFilterer) FilterFinalizeWithdrawERC1155(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*Il1erc1155gatewayFinalizeWithdrawERC1155Iterator, error) {

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

	logs, sub, err := _Il1erc1155gateway.contract.FilterLogs(opts, "FinalizeWithdrawERC1155", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &Il1erc1155gatewayFinalizeWithdrawERC1155Iterator{contract: _Il1erc1155gateway.contract, event: "FinalizeWithdrawERC1155", logs: logs, sub: sub}, nil
}

// WatchFinalizeWithdrawERC1155 is a free log subscription operation binding the contract event 0xfcc2841e9e72e6d610944e1b668912e92d5df94003055dbe06d615ba8d9efad4.
//
// Solidity: event FinalizeWithdrawERC1155(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId, uint256 _amount)
func (_Il1erc1155gateway *Il1erc1155gatewayFilterer) WatchFinalizeWithdrawERC1155(opts *bind.WatchOpts, sink chan<- *Il1erc1155gatewayFinalizeWithdrawERC1155, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Il1erc1155gateway.contract.WatchLogs(opts, "FinalizeWithdrawERC1155", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il1erc1155gatewayFinalizeWithdrawERC1155)
				if err := _Il1erc1155gateway.contract.UnpackLog(event, "FinalizeWithdrawERC1155", log); err != nil {
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

// ParseFinalizeWithdrawERC1155 is a log parse operation binding the contract event 0xfcc2841e9e72e6d610944e1b668912e92d5df94003055dbe06d615ba8d9efad4.
//
// Solidity: event FinalizeWithdrawERC1155(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId, uint256 _amount)
func (_Il1erc1155gateway *Il1erc1155gatewayFilterer) ParseFinalizeWithdrawERC1155(log types.Log) (*Il1erc1155gatewayFinalizeWithdrawERC1155, error) {
	event := new(Il1erc1155gatewayFinalizeWithdrawERC1155)
	if err := _Il1erc1155gateway.contract.UnpackLog(event, "FinalizeWithdrawERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Il1erc1155gatewayRefundERC1155Iterator is returned from FilterRefundERC1155 and is used to iterate over the raw logs and unpacked data for RefundERC1155 events raised by the Il1erc1155gateway contract.
type Il1erc1155gatewayRefundERC1155Iterator struct {
	Event *Il1erc1155gatewayRefundERC1155 // Event containing the contract specifics and raw log

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
func (it *Il1erc1155gatewayRefundERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Il1erc1155gatewayRefundERC1155)
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
		it.Event = new(Il1erc1155gatewayRefundERC1155)
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
func (it *Il1erc1155gatewayRefundERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Il1erc1155gatewayRefundERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Il1erc1155gatewayRefundERC1155 represents a RefundERC1155 event raised by the Il1erc1155gateway contract.
type Il1erc1155gatewayRefundERC1155 struct {
	Token     common.Address
	Recipient common.Address
	TokenId   *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundERC1155 is a free log retrieval operation binding the contract event 0xee285671d9ac3b0e0ed40037cb6db081095aa6cd68363f3e56989dde39e0df09.
//
// Solidity: event RefundERC1155(address indexed token, address indexed recipient, uint256 tokenId, uint256 amount)
func (_Il1erc1155gateway *Il1erc1155gatewayFilterer) FilterRefundERC1155(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*Il1erc1155gatewayRefundERC1155Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Il1erc1155gateway.contract.FilterLogs(opts, "RefundERC1155", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &Il1erc1155gatewayRefundERC1155Iterator{contract: _Il1erc1155gateway.contract, event: "RefundERC1155", logs: logs, sub: sub}, nil
}

// WatchRefundERC1155 is a free log subscription operation binding the contract event 0xee285671d9ac3b0e0ed40037cb6db081095aa6cd68363f3e56989dde39e0df09.
//
// Solidity: event RefundERC1155(address indexed token, address indexed recipient, uint256 tokenId, uint256 amount)
func (_Il1erc1155gateway *Il1erc1155gatewayFilterer) WatchRefundERC1155(opts *bind.WatchOpts, sink chan<- *Il1erc1155gatewayRefundERC1155, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Il1erc1155gateway.contract.WatchLogs(opts, "RefundERC1155", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Il1erc1155gatewayRefundERC1155)
				if err := _Il1erc1155gateway.contract.UnpackLog(event, "RefundERC1155", log); err != nil {
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

// ParseRefundERC1155 is a log parse operation binding the contract event 0xee285671d9ac3b0e0ed40037cb6db081095aa6cd68363f3e56989dde39e0df09.
//
// Solidity: event RefundERC1155(address indexed token, address indexed recipient, uint256 tokenId, uint256 amount)
func (_Il1erc1155gateway *Il1erc1155gatewayFilterer) ParseRefundERC1155(log types.Log) (*Il1erc1155gatewayRefundERC1155, error) {
	event := new(Il1erc1155gatewayRefundERC1155)
	if err := _Il1erc1155gateway.contract.UnpackLog(event, "RefundERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
