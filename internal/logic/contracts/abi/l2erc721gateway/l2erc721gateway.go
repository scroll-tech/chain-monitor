// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2erc721gateway

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

// L2erc721gatewayMetaData contains all meta data concerning the L2erc721gateway contract.
var L2erc721gatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"BatchWithdrawERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"FinalizeBatchDepositERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"FinalizeDepositERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldL1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newL1Token\",\"type\":\"address\"}],\"name\":\"UpdateTokenMapping\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC721\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"batchWithdrawERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"batchWithdrawERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"finalizeBatchDepositERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"finalizeDepositERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"updateTokenMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// L2erc721gatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L2erc721gatewayMetaData.ABI instead.
var L2erc721gatewayABI = L2erc721gatewayMetaData.ABI

// L2erc721gateway is an auto generated Go binding around an Ethereum contract.
type L2erc721gateway struct {
	L2erc721gatewayCaller     // Read-only binding to the contract
	L2erc721gatewayTransactor // Write-only binding to the contract
	L2erc721gatewayFilterer   // Log filterer for contract events
}

// L2erc721gatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2erc721gatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2erc721gatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2erc721gatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2erc721gatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2erc721gatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2erc721gatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2erc721gatewaySession struct {
	Contract     *L2erc721gateway  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2erc721gatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2erc721gatewayCallerSession struct {
	Contract *L2erc721gatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// L2erc721gatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2erc721gatewayTransactorSession struct {
	Contract     *L2erc721gatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// L2erc721gatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2erc721gatewayRaw struct {
	Contract *L2erc721gateway // Generic contract binding to access the raw methods on
}

// L2erc721gatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2erc721gatewayCallerRaw struct {
	Contract *L2erc721gatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L2erc721gatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2erc721gatewayTransactorRaw struct {
	Contract *L2erc721gatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2erc721gateway creates a new instance of L2erc721gateway, bound to a specific deployed contract.
func NewL2erc721gateway(address common.Address, backend bind.ContractBackend) (*L2erc721gateway, error) {
	contract, err := bindL2erc721gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2erc721gateway{L2erc721gatewayCaller: L2erc721gatewayCaller{contract: contract}, L2erc721gatewayTransactor: L2erc721gatewayTransactor{contract: contract}, L2erc721gatewayFilterer: L2erc721gatewayFilterer{contract: contract}}, nil
}

// NewL2erc721gatewayCaller creates a new read-only instance of L2erc721gateway, bound to a specific deployed contract.
func NewL2erc721gatewayCaller(address common.Address, caller bind.ContractCaller) (*L2erc721gatewayCaller, error) {
	contract, err := bindL2erc721gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2erc721gatewayCaller{contract: contract}, nil
}

// NewL2erc721gatewayTransactor creates a new write-only instance of L2erc721gateway, bound to a specific deployed contract.
func NewL2erc721gatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L2erc721gatewayTransactor, error) {
	contract, err := bindL2erc721gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2erc721gatewayTransactor{contract: contract}, nil
}

// NewL2erc721gatewayFilterer creates a new log filterer instance of L2erc721gateway, bound to a specific deployed contract.
func NewL2erc721gatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L2erc721gatewayFilterer, error) {
	contract, err := bindL2erc721gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2erc721gatewayFilterer{contract: contract}, nil
}

// bindL2erc721gateway binds a generic wrapper to an already deployed contract.
func bindL2erc721gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2erc721gatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2erc721gateway *L2erc721gatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2erc721gateway.Contract.L2erc721gatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2erc721gateway *L2erc721gatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.L2erc721gatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2erc721gateway *L2erc721gatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.L2erc721gatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2erc721gateway *L2erc721gatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2erc721gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2erc721gateway *L2erc721gatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2erc721gateway *L2erc721gatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.contract.Transact(opts, method, params...)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2erc721gateway *L2erc721gatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2erc721gateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2erc721gateway *L2erc721gatewaySession) Counterpart() (common.Address, error) {
	return _L2erc721gateway.Contract.Counterpart(&_L2erc721gateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2erc721gateway *L2erc721gatewayCallerSession) Counterpart() (common.Address, error) {
	return _L2erc721gateway.Contract.Counterpart(&_L2erc721gateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2erc721gateway *L2erc721gatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2erc721gateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2erc721gateway *L2erc721gatewaySession) Messenger() (common.Address, error) {
	return _L2erc721gateway.Contract.Messenger(&_L2erc721gateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2erc721gateway *L2erc721gatewayCallerSession) Messenger() (common.Address, error) {
	return _L2erc721gateway.Contract.Messenger(&_L2erc721gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2erc721gateway *L2erc721gatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2erc721gateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2erc721gateway *L2erc721gatewaySession) Owner() (common.Address, error) {
	return _L2erc721gateway.Contract.Owner(&_L2erc721gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2erc721gateway *L2erc721gatewayCallerSession) Owner() (common.Address, error) {
	return _L2erc721gateway.Contract.Owner(&_L2erc721gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2erc721gateway *L2erc721gatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2erc721gateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2erc721gateway *L2erc721gatewaySession) Router() (common.Address, error) {
	return _L2erc721gateway.Contract.Router(&_L2erc721gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2erc721gateway *L2erc721gatewayCallerSession) Router() (common.Address, error) {
	return _L2erc721gateway.Contract.Router(&_L2erc721gateway.CallOpts)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2erc721gateway *L2erc721gatewayCaller) TokenMapping(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2erc721gateway.contract.Call(opts, &out, "tokenMapping", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2erc721gateway *L2erc721gatewaySession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _L2erc721gateway.Contract.TokenMapping(&_L2erc721gateway.CallOpts, arg0)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2erc721gateway *L2erc721gatewayCallerSession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _L2erc721gateway.Contract.TokenMapping(&_L2erc721gateway.CallOpts, arg0)
}

// BatchWithdrawERC721 is a paid mutator transaction binding the contract method 0x46aa3411.
//
// Solidity: function batchWithdrawERC721(address _token, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L2erc721gateway *L2erc721gatewayTransactor) BatchWithdrawERC721(opts *bind.TransactOpts, _token common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc721gateway.contract.Transact(opts, "batchWithdrawERC721", _token, _tokenIds, _gasLimit)
}

// BatchWithdrawERC721 is a paid mutator transaction binding the contract method 0x46aa3411.
//
// Solidity: function batchWithdrawERC721(address _token, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L2erc721gateway *L2erc721gatewaySession) BatchWithdrawERC721(_token common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.BatchWithdrawERC721(&_L2erc721gateway.TransactOpts, _token, _tokenIds, _gasLimit)
}

// BatchWithdrawERC721 is a paid mutator transaction binding the contract method 0x46aa3411.
//
// Solidity: function batchWithdrawERC721(address _token, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L2erc721gateway *L2erc721gatewayTransactorSession) BatchWithdrawERC721(_token common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.BatchWithdrawERC721(&_L2erc721gateway.TransactOpts, _token, _tokenIds, _gasLimit)
}

// BatchWithdrawERC7210 is a paid mutator transaction binding the contract method 0xaa4c1158.
//
// Solidity: function batchWithdrawERC721(address _token, address _to, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L2erc721gateway *L2erc721gatewayTransactor) BatchWithdrawERC7210(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc721gateway.contract.Transact(opts, "batchWithdrawERC7210", _token, _to, _tokenIds, _gasLimit)
}

// BatchWithdrawERC7210 is a paid mutator transaction binding the contract method 0xaa4c1158.
//
// Solidity: function batchWithdrawERC721(address _token, address _to, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L2erc721gateway *L2erc721gatewaySession) BatchWithdrawERC7210(_token common.Address, _to common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.BatchWithdrawERC7210(&_L2erc721gateway.TransactOpts, _token, _to, _tokenIds, _gasLimit)
}

// BatchWithdrawERC7210 is a paid mutator transaction binding the contract method 0xaa4c1158.
//
// Solidity: function batchWithdrawERC721(address _token, address _to, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L2erc721gateway *L2erc721gatewayTransactorSession) BatchWithdrawERC7210(_token common.Address, _to common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.BatchWithdrawERC7210(&_L2erc721gateway.TransactOpts, _token, _to, _tokenIds, _gasLimit)
}

// FinalizeBatchDepositERC721 is a paid mutator transaction binding the contract method 0x982b151f.
//
// Solidity: function finalizeBatchDepositERC721(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds) returns()
func (_L2erc721gateway *L2erc721gatewayTransactor) FinalizeBatchDepositERC721(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _L2erc721gateway.contract.Transact(opts, "finalizeBatchDepositERC721", _l1Token, _l2Token, _from, _to, _tokenIds)
}

// FinalizeBatchDepositERC721 is a paid mutator transaction binding the contract method 0x982b151f.
//
// Solidity: function finalizeBatchDepositERC721(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds) returns()
func (_L2erc721gateway *L2erc721gatewaySession) FinalizeBatchDepositERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.FinalizeBatchDepositERC721(&_L2erc721gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenIds)
}

// FinalizeBatchDepositERC721 is a paid mutator transaction binding the contract method 0x982b151f.
//
// Solidity: function finalizeBatchDepositERC721(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds) returns()
func (_L2erc721gateway *L2erc721gatewayTransactorSession) FinalizeBatchDepositERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.FinalizeBatchDepositERC721(&_L2erc721gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenIds)
}

// FinalizeDepositERC721 is a paid mutator transaction binding the contract method 0xf8c3cf25.
//
// Solidity: function finalizeDepositERC721(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId) returns()
func (_L2erc721gateway *L2erc721gatewayTransactor) FinalizeDepositERC721(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _L2erc721gateway.contract.Transact(opts, "finalizeDepositERC721", _l1Token, _l2Token, _from, _to, _tokenId)
}

// FinalizeDepositERC721 is a paid mutator transaction binding the contract method 0xf8c3cf25.
//
// Solidity: function finalizeDepositERC721(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId) returns()
func (_L2erc721gateway *L2erc721gatewaySession) FinalizeDepositERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.FinalizeDepositERC721(&_L2erc721gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenId)
}

// FinalizeDepositERC721 is a paid mutator transaction binding the contract method 0xf8c3cf25.
//
// Solidity: function finalizeDepositERC721(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId) returns()
func (_L2erc721gateway *L2erc721gatewayTransactorSession) FinalizeDepositERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.FinalizeDepositERC721(&_L2erc721gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L2erc721gateway *L2erc721gatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2erc721gateway.contract.Transact(opts, "initialize", _counterpart, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L2erc721gateway *L2erc721gatewaySession) Initialize(_counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.Initialize(&_L2erc721gateway.TransactOpts, _counterpart, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L2erc721gateway *L2erc721gatewayTransactorSession) Initialize(_counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.Initialize(&_L2erc721gateway.TransactOpts, _counterpart, _messenger)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_L2erc721gateway *L2erc721gatewayTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2erc721gateway.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_L2erc721gateway *L2erc721gatewaySession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.OnERC721Received(&_L2erc721gateway.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_L2erc721gateway *L2erc721gatewayTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.OnERC721Received(&_L2erc721gateway.TransactOpts, arg0, arg1, arg2, arg3)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2erc721gateway *L2erc721gatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2erc721gateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2erc721gateway *L2erc721gatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _L2erc721gateway.Contract.RenounceOwnership(&_L2erc721gateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2erc721gateway *L2erc721gatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2erc721gateway.Contract.RenounceOwnership(&_L2erc721gateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2erc721gateway *L2erc721gatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2erc721gateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2erc721gateway *L2erc721gatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.TransferOwnership(&_L2erc721gateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2erc721gateway *L2erc721gatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.TransferOwnership(&_L2erc721gateway.TransactOpts, newOwner)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2erc721gateway *L2erc721gatewayTransactor) UpdateTokenMapping(opts *bind.TransactOpts, _l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2erc721gateway.contract.Transact(opts, "updateTokenMapping", _l2Token, _l1Token)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2erc721gateway *L2erc721gatewaySession) UpdateTokenMapping(_l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.UpdateTokenMapping(&_L2erc721gateway.TransactOpts, _l2Token, _l1Token)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2erc721gateway *L2erc721gatewayTransactorSession) UpdateTokenMapping(_l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.UpdateTokenMapping(&_L2erc721gateway.TransactOpts, _l2Token, _l1Token)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x2a491247.
//
// Solidity: function withdrawERC721(address _token, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L2erc721gateway *L2erc721gatewayTransactor) WithdrawERC721(opts *bind.TransactOpts, _token common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc721gateway.contract.Transact(opts, "withdrawERC721", _token, _tokenId, _gasLimit)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x2a491247.
//
// Solidity: function withdrawERC721(address _token, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L2erc721gateway *L2erc721gatewaySession) WithdrawERC721(_token common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.WithdrawERC721(&_L2erc721gateway.TransactOpts, _token, _tokenId, _gasLimit)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x2a491247.
//
// Solidity: function withdrawERC721(address _token, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L2erc721gateway *L2erc721gatewayTransactorSession) WithdrawERC721(_token common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.WithdrawERC721(&_L2erc721gateway.TransactOpts, _token, _tokenId, _gasLimit)
}

// WithdrawERC7210 is a paid mutator transaction binding the contract method 0xee5a8db2.
//
// Solidity: function withdrawERC721(address _token, address _to, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L2erc721gateway *L2erc721gatewayTransactor) WithdrawERC7210(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc721gateway.contract.Transact(opts, "withdrawERC7210", _token, _to, _tokenId, _gasLimit)
}

// WithdrawERC7210 is a paid mutator transaction binding the contract method 0xee5a8db2.
//
// Solidity: function withdrawERC721(address _token, address _to, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L2erc721gateway *L2erc721gatewaySession) WithdrawERC7210(_token common.Address, _to common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.WithdrawERC7210(&_L2erc721gateway.TransactOpts, _token, _to, _tokenId, _gasLimit)
}

// WithdrawERC7210 is a paid mutator transaction binding the contract method 0xee5a8db2.
//
// Solidity: function withdrawERC721(address _token, address _to, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L2erc721gateway *L2erc721gatewayTransactorSession) WithdrawERC7210(_token common.Address, _to common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc721gateway.Contract.WithdrawERC7210(&_L2erc721gateway.TransactOpts, _token, _to, _tokenId, _gasLimit)
}

// L2erc721gatewayBatchWithdrawERC721Iterator is returned from FilterBatchWithdrawERC721 and is used to iterate over the raw logs and unpacked data for BatchWithdrawERC721 events raised by the L2erc721gateway contract.
type L2erc721gatewayBatchWithdrawERC721Iterator struct {
	Event *L2erc721gatewayBatchWithdrawERC721 // Event containing the contract specifics and raw log

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
func (it *L2erc721gatewayBatchWithdrawERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2erc721gatewayBatchWithdrawERC721)
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
		it.Event = new(L2erc721gatewayBatchWithdrawERC721)
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
func (it *L2erc721gatewayBatchWithdrawERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2erc721gatewayBatchWithdrawERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2erc721gatewayBatchWithdrawERC721 represents a BatchWithdrawERC721 event raised by the L2erc721gateway contract.
type L2erc721gatewayBatchWithdrawERC721 struct {
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
func (_L2erc721gateway *L2erc721gatewayFilterer) FilterBatchWithdrawERC721(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2erc721gatewayBatchWithdrawERC721Iterator, error) {

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

	logs, sub, err := _L2erc721gateway.contract.FilterLogs(opts, "BatchWithdrawERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2erc721gatewayBatchWithdrawERC721Iterator{contract: _L2erc721gateway.contract, event: "BatchWithdrawERC721", logs: logs, sub: sub}, nil
}

// WatchBatchWithdrawERC721 is a free log subscription operation binding the contract event 0xbdb7b5cec70093e3ce49b258071951d245c0871c006fd9327778c69d0e9f244d.
//
// Solidity: event BatchWithdrawERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds)
func (_L2erc721gateway *L2erc721gatewayFilterer) WatchBatchWithdrawERC721(opts *bind.WatchOpts, sink chan<- *L2erc721gatewayBatchWithdrawERC721, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2erc721gateway.contract.WatchLogs(opts, "BatchWithdrawERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2erc721gatewayBatchWithdrawERC721)
				if err := _L2erc721gateway.contract.UnpackLog(event, "BatchWithdrawERC721", log); err != nil {
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
func (_L2erc721gateway *L2erc721gatewayFilterer) ParseBatchWithdrawERC721(log types.Log) (*L2erc721gatewayBatchWithdrawERC721, error) {
	event := new(L2erc721gatewayBatchWithdrawERC721)
	if err := _L2erc721gateway.contract.UnpackLog(event, "BatchWithdrawERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2erc721gatewayFinalizeBatchDepositERC721Iterator is returned from FilterFinalizeBatchDepositERC721 and is used to iterate over the raw logs and unpacked data for FinalizeBatchDepositERC721 events raised by the L2erc721gateway contract.
type L2erc721gatewayFinalizeBatchDepositERC721Iterator struct {
	Event *L2erc721gatewayFinalizeBatchDepositERC721 // Event containing the contract specifics and raw log

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
func (it *L2erc721gatewayFinalizeBatchDepositERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2erc721gatewayFinalizeBatchDepositERC721)
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
		it.Event = new(L2erc721gatewayFinalizeBatchDepositERC721)
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
func (it *L2erc721gatewayFinalizeBatchDepositERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2erc721gatewayFinalizeBatchDepositERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2erc721gatewayFinalizeBatchDepositERC721 represents a FinalizeBatchDepositERC721 event raised by the L2erc721gateway contract.
type L2erc721gatewayFinalizeBatchDepositERC721 struct {
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
func (_L2erc721gateway *L2erc721gatewayFilterer) FilterFinalizeBatchDepositERC721(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2erc721gatewayFinalizeBatchDepositERC721Iterator, error) {

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

	logs, sub, err := _L2erc721gateway.contract.FilterLogs(opts, "FinalizeBatchDepositERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2erc721gatewayFinalizeBatchDepositERC721Iterator{contract: _L2erc721gateway.contract, event: "FinalizeBatchDepositERC721", logs: logs, sub: sub}, nil
}

// WatchFinalizeBatchDepositERC721 is a free log subscription operation binding the contract event 0xafa88b850da44ca05b319e813873eac8d08e7c041d2d9b3072db0f087e3cd29e.
//
// Solidity: event FinalizeBatchDepositERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds)
func (_L2erc721gateway *L2erc721gatewayFilterer) WatchFinalizeBatchDepositERC721(opts *bind.WatchOpts, sink chan<- *L2erc721gatewayFinalizeBatchDepositERC721, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2erc721gateway.contract.WatchLogs(opts, "FinalizeBatchDepositERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2erc721gatewayFinalizeBatchDepositERC721)
				if err := _L2erc721gateway.contract.UnpackLog(event, "FinalizeBatchDepositERC721", log); err != nil {
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
func (_L2erc721gateway *L2erc721gatewayFilterer) ParseFinalizeBatchDepositERC721(log types.Log) (*L2erc721gatewayFinalizeBatchDepositERC721, error) {
	event := new(L2erc721gatewayFinalizeBatchDepositERC721)
	if err := _L2erc721gateway.contract.UnpackLog(event, "FinalizeBatchDepositERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2erc721gatewayFinalizeDepositERC721Iterator is returned from FilterFinalizeDepositERC721 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC721 events raised by the L2erc721gateway contract.
type L2erc721gatewayFinalizeDepositERC721Iterator struct {
	Event *L2erc721gatewayFinalizeDepositERC721 // Event containing the contract specifics and raw log

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
func (it *L2erc721gatewayFinalizeDepositERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2erc721gatewayFinalizeDepositERC721)
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
		it.Event = new(L2erc721gatewayFinalizeDepositERC721)
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
func (it *L2erc721gatewayFinalizeDepositERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2erc721gatewayFinalizeDepositERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2erc721gatewayFinalizeDepositERC721 represents a FinalizeDepositERC721 event raised by the L2erc721gateway contract.
type L2erc721gatewayFinalizeDepositERC721 struct {
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
func (_L2erc721gateway *L2erc721gatewayFilterer) FilterFinalizeDepositERC721(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2erc721gatewayFinalizeDepositERC721Iterator, error) {

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

	logs, sub, err := _L2erc721gateway.contract.FilterLogs(opts, "FinalizeDepositERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2erc721gatewayFinalizeDepositERC721Iterator{contract: _L2erc721gateway.contract, event: "FinalizeDepositERC721", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC721 is a free log subscription operation binding the contract event 0xc655ec1de34d98630aa4572239414f926d6b3d07653dde093a6df97377e31b41.
//
// Solidity: event FinalizeDepositERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId)
func (_L2erc721gateway *L2erc721gatewayFilterer) WatchFinalizeDepositERC721(opts *bind.WatchOpts, sink chan<- *L2erc721gatewayFinalizeDepositERC721, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2erc721gateway.contract.WatchLogs(opts, "FinalizeDepositERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2erc721gatewayFinalizeDepositERC721)
				if err := _L2erc721gateway.contract.UnpackLog(event, "FinalizeDepositERC721", log); err != nil {
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
func (_L2erc721gateway *L2erc721gatewayFilterer) ParseFinalizeDepositERC721(log types.Log) (*L2erc721gatewayFinalizeDepositERC721, error) {
	event := new(L2erc721gatewayFinalizeDepositERC721)
	if err := _L2erc721gateway.contract.UnpackLog(event, "FinalizeDepositERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2erc721gatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2erc721gateway contract.
type L2erc721gatewayInitializedIterator struct {
	Event *L2erc721gatewayInitialized // Event containing the contract specifics and raw log

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
func (it *L2erc721gatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2erc721gatewayInitialized)
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
		it.Event = new(L2erc721gatewayInitialized)
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
func (it *L2erc721gatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2erc721gatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2erc721gatewayInitialized represents a Initialized event raised by the L2erc721gateway contract.
type L2erc721gatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2erc721gateway *L2erc721gatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2erc721gatewayInitializedIterator, error) {

	logs, sub, err := _L2erc721gateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2erc721gatewayInitializedIterator{contract: _L2erc721gateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2erc721gateway *L2erc721gatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2erc721gatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _L2erc721gateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2erc721gatewayInitialized)
				if err := _L2erc721gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2erc721gateway *L2erc721gatewayFilterer) ParseInitialized(log types.Log) (*L2erc721gatewayInitialized, error) {
	event := new(L2erc721gatewayInitialized)
	if err := _L2erc721gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2erc721gatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2erc721gateway contract.
type L2erc721gatewayOwnershipTransferredIterator struct {
	Event *L2erc721gatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2erc721gatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2erc721gatewayOwnershipTransferred)
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
		it.Event = new(L2erc721gatewayOwnershipTransferred)
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
func (it *L2erc721gatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2erc721gatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2erc721gatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L2erc721gateway contract.
type L2erc721gatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2erc721gateway *L2erc721gatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2erc721gatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2erc721gateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2erc721gatewayOwnershipTransferredIterator{contract: _L2erc721gateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2erc721gateway *L2erc721gatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2erc721gatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2erc721gateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2erc721gatewayOwnershipTransferred)
				if err := _L2erc721gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2erc721gateway *L2erc721gatewayFilterer) ParseOwnershipTransferred(log types.Log) (*L2erc721gatewayOwnershipTransferred, error) {
	event := new(L2erc721gatewayOwnershipTransferred)
	if err := _L2erc721gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2erc721gatewayUpdateTokenMappingIterator is returned from FilterUpdateTokenMapping and is used to iterate over the raw logs and unpacked data for UpdateTokenMapping events raised by the L2erc721gateway contract.
type L2erc721gatewayUpdateTokenMappingIterator struct {
	Event *L2erc721gatewayUpdateTokenMapping // Event containing the contract specifics and raw log

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
func (it *L2erc721gatewayUpdateTokenMappingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2erc721gatewayUpdateTokenMapping)
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
		it.Event = new(L2erc721gatewayUpdateTokenMapping)
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
func (it *L2erc721gatewayUpdateTokenMappingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2erc721gatewayUpdateTokenMappingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2erc721gatewayUpdateTokenMapping represents a UpdateTokenMapping event raised by the L2erc721gateway contract.
type L2erc721gatewayUpdateTokenMapping struct {
	L2Token    common.Address
	OldL1Token common.Address
	NewL1Token common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenMapping is a free log retrieval operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l2Token, address indexed oldL1Token, address indexed newL1Token)
func (_L2erc721gateway *L2erc721gatewayFilterer) FilterUpdateTokenMapping(opts *bind.FilterOpts, l2Token []common.Address, oldL1Token []common.Address, newL1Token []common.Address) (*L2erc721gatewayUpdateTokenMappingIterator, error) {

	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var oldL1TokenRule []interface{}
	for _, oldL1TokenItem := range oldL1Token {
		oldL1TokenRule = append(oldL1TokenRule, oldL1TokenItem)
	}
	var newL1TokenRule []interface{}
	for _, newL1TokenItem := range newL1Token {
		newL1TokenRule = append(newL1TokenRule, newL1TokenItem)
	}

	logs, sub, err := _L2erc721gateway.contract.FilterLogs(opts, "UpdateTokenMapping", l2TokenRule, oldL1TokenRule, newL1TokenRule)
	if err != nil {
		return nil, err
	}
	return &L2erc721gatewayUpdateTokenMappingIterator{contract: _L2erc721gateway.contract, event: "UpdateTokenMapping", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenMapping is a free log subscription operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l2Token, address indexed oldL1Token, address indexed newL1Token)
func (_L2erc721gateway *L2erc721gatewayFilterer) WatchUpdateTokenMapping(opts *bind.WatchOpts, sink chan<- *L2erc721gatewayUpdateTokenMapping, l2Token []common.Address, oldL1Token []common.Address, newL1Token []common.Address) (event.Subscription, error) {

	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var oldL1TokenRule []interface{}
	for _, oldL1TokenItem := range oldL1Token {
		oldL1TokenRule = append(oldL1TokenRule, oldL1TokenItem)
	}
	var newL1TokenRule []interface{}
	for _, newL1TokenItem := range newL1Token {
		newL1TokenRule = append(newL1TokenRule, newL1TokenItem)
	}

	logs, sub, err := _L2erc721gateway.contract.WatchLogs(opts, "UpdateTokenMapping", l2TokenRule, oldL1TokenRule, newL1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2erc721gatewayUpdateTokenMapping)
				if err := _L2erc721gateway.contract.UnpackLog(event, "UpdateTokenMapping", log); err != nil {
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

// ParseUpdateTokenMapping is a log parse operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l2Token, address indexed oldL1Token, address indexed newL1Token)
func (_L2erc721gateway *L2erc721gatewayFilterer) ParseUpdateTokenMapping(log types.Log) (*L2erc721gatewayUpdateTokenMapping, error) {
	event := new(L2erc721gatewayUpdateTokenMapping)
	if err := _L2erc721gateway.contract.UnpackLog(event, "UpdateTokenMapping", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2erc721gatewayWithdrawERC721Iterator is returned from FilterWithdrawERC721 and is used to iterate over the raw logs and unpacked data for WithdrawERC721 events raised by the L2erc721gateway contract.
type L2erc721gatewayWithdrawERC721Iterator struct {
	Event *L2erc721gatewayWithdrawERC721 // Event containing the contract specifics and raw log

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
func (it *L2erc721gatewayWithdrawERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2erc721gatewayWithdrawERC721)
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
		it.Event = new(L2erc721gatewayWithdrawERC721)
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
func (it *L2erc721gatewayWithdrawERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2erc721gatewayWithdrawERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2erc721gatewayWithdrawERC721 represents a WithdrawERC721 event raised by the L2erc721gateway contract.
type L2erc721gatewayWithdrawERC721 struct {
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
func (_L2erc721gateway *L2erc721gatewayFilterer) FilterWithdrawERC721(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2erc721gatewayWithdrawERC721Iterator, error) {

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

	logs, sub, err := _L2erc721gateway.contract.FilterLogs(opts, "WithdrawERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2erc721gatewayWithdrawERC721Iterator{contract: _L2erc721gateway.contract, event: "WithdrawERC721", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC721 is a free log subscription operation binding the contract event 0xe9e85cf0c862dd491ecda3c9a230e12ada8956472028ebde4fdc4f8e2d77bcda.
//
// Solidity: event WithdrawERC721(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId)
func (_L2erc721gateway *L2erc721gatewayFilterer) WatchWithdrawERC721(opts *bind.WatchOpts, sink chan<- *L2erc721gatewayWithdrawERC721, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2erc721gateway.contract.WatchLogs(opts, "WithdrawERC721", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2erc721gatewayWithdrawERC721)
				if err := _L2erc721gateway.contract.UnpackLog(event, "WithdrawERC721", log); err != nil {
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
func (_L2erc721gateway *L2erc721gatewayFilterer) ParseWithdrawERC721(log types.Log) (*L2erc721gatewayWithdrawERC721, error) {
	event := new(L2erc721gatewayWithdrawERC721)
	if err := _L2erc721gateway.contract.UnpackLog(event, "WithdrawERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
