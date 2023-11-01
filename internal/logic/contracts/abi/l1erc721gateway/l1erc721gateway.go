// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l1erc721gateway

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

// L1erc721gatewayMetaData contains all meta data concerning the L1erc721gateway contract.
var L1erc721gatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"BatchDepositERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"BatchRefundERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"DepositERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"FinalizeBatchWithdrawERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"FinalizeWithdrawERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"RefundERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldL2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newL2Token\",\"type\":\"address\"}],\"name\":\"UpdateTokenMapping\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"batchDepositERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"batchDepositERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"depositERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"finalizeBatchWithdrawERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"finalizeWithdrawERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"onDropMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"}],\"name\":\"updateTokenMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// L1erc721gatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L1erc721gatewayMetaData.ABI instead.
var L1erc721gatewayABI = L1erc721gatewayMetaData.ABI

// L1erc721gateway is an auto generated Go binding around an Ethereum contract.
type L1erc721gateway struct {
	L1erc721gatewayCaller     // Read-only binding to the contract
	L1erc721gatewayTransactor // Write-only binding to the contract
	L1erc721gatewayFilterer   // Log filterer for contract events
}

// L1erc721gatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1erc721gatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1erc721gatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1erc721gatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1erc721gatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1erc721gatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1erc721gatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1erc721gatewaySession struct {
	Contract     *L1erc721gateway  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1erc721gatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1erc721gatewayCallerSession struct {
	Contract *L1erc721gatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// L1erc721gatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1erc721gatewayTransactorSession struct {
	Contract     *L1erc721gatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// L1erc721gatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1erc721gatewayRaw struct {
	Contract *L1erc721gateway // Generic contract binding to access the raw methods on
}

// L1erc721gatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1erc721gatewayCallerRaw struct {
	Contract *L1erc721gatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L1erc721gatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1erc721gatewayTransactorRaw struct {
	Contract *L1erc721gatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1erc721gateway creates a new instance of L1erc721gateway, bound to a specific deployed contract.
func NewL1erc721gateway(address common.Address, backend bind.ContractBackend) (*L1erc721gateway, error) {
	contract, err := bindL1erc721gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1erc721gateway{L1erc721gatewayCaller: L1erc721gatewayCaller{contract: contract}, L1erc721gatewayTransactor: L1erc721gatewayTransactor{contract: contract}, L1erc721gatewayFilterer: L1erc721gatewayFilterer{contract: contract}}, nil
}

// NewL1erc721gatewayCaller creates a new read-only instance of L1erc721gateway, bound to a specific deployed contract.
func NewL1erc721gatewayCaller(address common.Address, caller bind.ContractCaller) (*L1erc721gatewayCaller, error) {
	contract, err := bindL1erc721gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1erc721gatewayCaller{contract: contract}, nil
}

// NewL1erc721gatewayTransactor creates a new write-only instance of L1erc721gateway, bound to a specific deployed contract.
func NewL1erc721gatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L1erc721gatewayTransactor, error) {
	contract, err := bindL1erc721gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1erc721gatewayTransactor{contract: contract}, nil
}

// NewL1erc721gatewayFilterer creates a new log filterer instance of L1erc721gateway, bound to a specific deployed contract.
func NewL1erc721gatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L1erc721gatewayFilterer, error) {
	contract, err := bindL1erc721gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1erc721gatewayFilterer{contract: contract}, nil
}

// bindL1erc721gateway binds a generic wrapper to an already deployed contract.
func bindL1erc721gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1erc721gatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1erc721gateway *L1erc721gatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1erc721gateway.Contract.L1erc721gatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1erc721gateway *L1erc721gatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.L1erc721gatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1erc721gateway *L1erc721gatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.L1erc721gatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1erc721gateway *L1erc721gatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1erc721gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1erc721gateway *L1erc721gatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1erc721gateway *L1erc721gatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.contract.Transact(opts, method, params...)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1erc721gateway *L1erc721gatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1erc721gateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1erc721gateway *L1erc721gatewaySession) Counterpart() (common.Address, error) {
	return _L1erc721gateway.Contract.Counterpart(&_L1erc721gateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1erc721gateway *L1erc721gatewayCallerSession) Counterpart() (common.Address, error) {
	return _L1erc721gateway.Contract.Counterpart(&_L1erc721gateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1erc721gateway *L1erc721gatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1erc721gateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1erc721gateway *L1erc721gatewaySession) Messenger() (common.Address, error) {
	return _L1erc721gateway.Contract.Messenger(&_L1erc721gateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1erc721gateway *L1erc721gatewayCallerSession) Messenger() (common.Address, error) {
	return _L1erc721gateway.Contract.Messenger(&_L1erc721gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1erc721gateway *L1erc721gatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1erc721gateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1erc721gateway *L1erc721gatewaySession) Owner() (common.Address, error) {
	return _L1erc721gateway.Contract.Owner(&_L1erc721gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1erc721gateway *L1erc721gatewayCallerSession) Owner() (common.Address, error) {
	return _L1erc721gateway.Contract.Owner(&_L1erc721gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1erc721gateway *L1erc721gatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1erc721gateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1erc721gateway *L1erc721gatewaySession) Router() (common.Address, error) {
	return _L1erc721gateway.Contract.Router(&_L1erc721gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1erc721gateway *L1erc721gatewayCallerSession) Router() (common.Address, error) {
	return _L1erc721gateway.Contract.Router(&_L1erc721gateway.CallOpts)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L1erc721gateway *L1erc721gatewayCaller) TokenMapping(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1erc721gateway.contract.Call(opts, &out, "tokenMapping", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L1erc721gateway *L1erc721gatewaySession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _L1erc721gateway.Contract.TokenMapping(&_L1erc721gateway.CallOpts, arg0)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L1erc721gateway *L1erc721gatewayCallerSession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _L1erc721gateway.Contract.TokenMapping(&_L1erc721gateway.CallOpts, arg0)
}

// BatchDepositERC721 is a paid mutator transaction binding the contract method 0x1b997a93.
//
// Solidity: function batchDepositERC721(address _token, address _to, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L1erc721gateway *L1erc721gatewayTransactor) BatchDepositERC721(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1erc721gateway.contract.Transact(opts, "batchDepositERC721", _token, _to, _tokenIds, _gasLimit)
}

// BatchDepositERC721 is a paid mutator transaction binding the contract method 0x1b997a93.
//
// Solidity: function batchDepositERC721(address _token, address _to, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L1erc721gateway *L1erc721gatewaySession) BatchDepositERC721(_token common.Address, _to common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.BatchDepositERC721(&_L1erc721gateway.TransactOpts, _token, _to, _tokenIds, _gasLimit)
}

// BatchDepositERC721 is a paid mutator transaction binding the contract method 0x1b997a93.
//
// Solidity: function batchDepositERC721(address _token, address _to, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L1erc721gateway *L1erc721gatewayTransactorSession) BatchDepositERC721(_token common.Address, _to common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.BatchDepositERC721(&_L1erc721gateway.TransactOpts, _token, _to, _tokenIds, _gasLimit)
}

// BatchDepositERC7210 is a paid mutator transaction binding the contract method 0xd96c8ecf.
//
// Solidity: function batchDepositERC721(address _token, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L1erc721gateway *L1erc721gatewayTransactor) BatchDepositERC7210(opts *bind.TransactOpts, _token common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1erc721gateway.contract.Transact(opts, "batchDepositERC7210", _token, _tokenIds, _gasLimit)
}

// BatchDepositERC7210 is a paid mutator transaction binding the contract method 0xd96c8ecf.
//
// Solidity: function batchDepositERC721(address _token, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L1erc721gateway *L1erc721gatewaySession) BatchDepositERC7210(_token common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.BatchDepositERC7210(&_L1erc721gateway.TransactOpts, _token, _tokenIds, _gasLimit)
}

// BatchDepositERC7210 is a paid mutator transaction binding the contract method 0xd96c8ecf.
//
// Solidity: function batchDepositERC721(address _token, uint256[] _tokenIds, uint256 _gasLimit) payable returns()
func (_L1erc721gateway *L1erc721gatewayTransactorSession) BatchDepositERC7210(_token common.Address, _tokenIds []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.BatchDepositERC7210(&_L1erc721gateway.TransactOpts, _token, _tokenIds, _gasLimit)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x0a7aa196.
//
// Solidity: function depositERC721(address _token, address _to, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L1erc721gateway *L1erc721gatewayTransactor) DepositERC721(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1erc721gateway.contract.Transact(opts, "depositERC721", _token, _to, _tokenId, _gasLimit)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x0a7aa196.
//
// Solidity: function depositERC721(address _token, address _to, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L1erc721gateway *L1erc721gatewaySession) DepositERC721(_token common.Address, _to common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.DepositERC721(&_L1erc721gateway.TransactOpts, _token, _to, _tokenId, _gasLimit)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x0a7aa196.
//
// Solidity: function depositERC721(address _token, address _to, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L1erc721gateway *L1erc721gatewayTransactorSession) DepositERC721(_token common.Address, _to common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.DepositERC721(&_L1erc721gateway.TransactOpts, _token, _to, _tokenId, _gasLimit)
}

// DepositERC7210 is a paid mutator transaction binding the contract method 0x45a4276b.
//
// Solidity: function depositERC721(address _token, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L1erc721gateway *L1erc721gatewayTransactor) DepositERC7210(opts *bind.TransactOpts, _token common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1erc721gateway.contract.Transact(opts, "depositERC7210", _token, _tokenId, _gasLimit)
}

// DepositERC7210 is a paid mutator transaction binding the contract method 0x45a4276b.
//
// Solidity: function depositERC721(address _token, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L1erc721gateway *L1erc721gatewaySession) DepositERC7210(_token common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.DepositERC7210(&_L1erc721gateway.TransactOpts, _token, _tokenId, _gasLimit)
}

// DepositERC7210 is a paid mutator transaction binding the contract method 0x45a4276b.
//
// Solidity: function depositERC721(address _token, uint256 _tokenId, uint256 _gasLimit) payable returns()
func (_L1erc721gateway *L1erc721gatewayTransactorSession) DepositERC7210(_token common.Address, _tokenId *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.DepositERC7210(&_L1erc721gateway.TransactOpts, _token, _tokenId, _gasLimit)
}

// FinalizeBatchWithdrawERC721 is a paid mutator transaction binding the contract method 0x9f0a68b3.
//
// Solidity: function finalizeBatchWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds) returns()
func (_L1erc721gateway *L1erc721gatewayTransactor) FinalizeBatchWithdrawERC721(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _L1erc721gateway.contract.Transact(opts, "finalizeBatchWithdrawERC721", _l1Token, _l2Token, _from, _to, _tokenIds)
}

// FinalizeBatchWithdrawERC721 is a paid mutator transaction binding the contract method 0x9f0a68b3.
//
// Solidity: function finalizeBatchWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds) returns()
func (_L1erc721gateway *L1erc721gatewaySession) FinalizeBatchWithdrawERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.FinalizeBatchWithdrawERC721(&_L1erc721gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenIds)
}

// FinalizeBatchWithdrawERC721 is a paid mutator transaction binding the contract method 0x9f0a68b3.
//
// Solidity: function finalizeBatchWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds) returns()
func (_L1erc721gateway *L1erc721gatewayTransactorSession) FinalizeBatchWithdrawERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.FinalizeBatchWithdrawERC721(&_L1erc721gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenIds)
}

// FinalizeWithdrawERC721 is a paid mutator transaction binding the contract method 0xd606b4dc.
//
// Solidity: function finalizeWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId) returns()
func (_L1erc721gateway *L1erc721gatewayTransactor) FinalizeWithdrawERC721(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _L1erc721gateway.contract.Transact(opts, "finalizeWithdrawERC721", _l1Token, _l2Token, _from, _to, _tokenId)
}

// FinalizeWithdrawERC721 is a paid mutator transaction binding the contract method 0xd606b4dc.
//
// Solidity: function finalizeWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId) returns()
func (_L1erc721gateway *L1erc721gatewaySession) FinalizeWithdrawERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.FinalizeWithdrawERC721(&_L1erc721gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenId)
}

// FinalizeWithdrawERC721 is a paid mutator transaction binding the contract method 0xd606b4dc.
//
// Solidity: function finalizeWithdrawERC721(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId) returns()
func (_L1erc721gateway *L1erc721gatewayTransactorSession) FinalizeWithdrawERC721(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.FinalizeWithdrawERC721(&_L1erc721gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L1erc721gateway *L1erc721gatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1erc721gateway.contract.Transact(opts, "initialize", _counterpart, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L1erc721gateway *L1erc721gatewaySession) Initialize(_counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.Initialize(&_L1erc721gateway.TransactOpts, _counterpart, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L1erc721gateway *L1erc721gatewayTransactorSession) Initialize(_counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.Initialize(&_L1erc721gateway.TransactOpts, _counterpart, _messenger)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1erc721gateway *L1erc721gatewayTransactor) OnDropMessage(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _L1erc721gateway.contract.Transact(opts, "onDropMessage", _message)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1erc721gateway *L1erc721gatewaySession) OnDropMessage(_message []byte) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.OnDropMessage(&_L1erc721gateway.TransactOpts, _message)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1erc721gateway *L1erc721gatewayTransactorSession) OnDropMessage(_message []byte) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.OnDropMessage(&_L1erc721gateway.TransactOpts, _message)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_L1erc721gateway *L1erc721gatewayTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L1erc721gateway.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_L1erc721gateway *L1erc721gatewaySession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.OnERC721Received(&_L1erc721gateway.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_L1erc721gateway *L1erc721gatewayTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.OnERC721Received(&_L1erc721gateway.TransactOpts, arg0, arg1, arg2, arg3)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1erc721gateway *L1erc721gatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1erc721gateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1erc721gateway *L1erc721gatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _L1erc721gateway.Contract.RenounceOwnership(&_L1erc721gateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1erc721gateway *L1erc721gatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1erc721gateway.Contract.RenounceOwnership(&_L1erc721gateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1erc721gateway *L1erc721gatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1erc721gateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1erc721gateway *L1erc721gatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.TransferOwnership(&_L1erc721gateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1erc721gateway *L1erc721gatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.TransferOwnership(&_L1erc721gateway.TransactOpts, newOwner)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l1Token, address _l2Token) returns()
func (_L1erc721gateway *L1erc721gatewayTransactor) UpdateTokenMapping(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address) (*types.Transaction, error) {
	return _L1erc721gateway.contract.Transact(opts, "updateTokenMapping", _l1Token, _l2Token)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l1Token, address _l2Token) returns()
func (_L1erc721gateway *L1erc721gatewaySession) UpdateTokenMapping(_l1Token common.Address, _l2Token common.Address) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.UpdateTokenMapping(&_L1erc721gateway.TransactOpts, _l1Token, _l2Token)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l1Token, address _l2Token) returns()
func (_L1erc721gateway *L1erc721gatewayTransactorSession) UpdateTokenMapping(_l1Token common.Address, _l2Token common.Address) (*types.Transaction, error) {
	return _L1erc721gateway.Contract.UpdateTokenMapping(&_L1erc721gateway.TransactOpts, _l1Token, _l2Token)
}

// L1erc721gatewayBatchDepositERC721Iterator is returned from FilterBatchDepositERC721 and is used to iterate over the raw logs and unpacked data for BatchDepositERC721 events raised by the L1erc721gateway contract.
type L1erc721gatewayBatchDepositERC721Iterator struct {
	Event *L1erc721gatewayBatchDepositERC721 // Event containing the contract specifics and raw log

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
func (it *L1erc721gatewayBatchDepositERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1erc721gatewayBatchDepositERC721)
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
		it.Event = new(L1erc721gatewayBatchDepositERC721)
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
func (it *L1erc721gatewayBatchDepositERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1erc721gatewayBatchDepositERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1erc721gatewayBatchDepositERC721 represents a BatchDepositERC721 event raised by the L1erc721gateway contract.
type L1erc721gatewayBatchDepositERC721 struct {
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
func (_L1erc721gateway *L1erc721gatewayFilterer) FilterBatchDepositERC721(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*L1erc721gatewayBatchDepositERC721Iterator, error) {

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

	logs, sub, err := _L1erc721gateway.contract.FilterLogs(opts, "BatchDepositERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &L1erc721gatewayBatchDepositERC721Iterator{contract: _L1erc721gateway.contract, event: "BatchDepositERC721", logs: logs, sub: sub}, nil
}

// WatchBatchDepositERC721 is a free log subscription operation binding the contract event 0xf05915e3b4fbd6f61b8b6f80b07f10e1cad039ccc7abe7c7fec115d038fe3dd6.
//
// Solidity: event BatchDepositERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds)
func (_L1erc721gateway *L1erc721gatewayFilterer) WatchBatchDepositERC721(opts *bind.WatchOpts, sink chan<- *L1erc721gatewayBatchDepositERC721, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1erc721gateway.contract.WatchLogs(opts, "BatchDepositERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1erc721gatewayBatchDepositERC721)
				if err := _L1erc721gateway.contract.UnpackLog(event, "BatchDepositERC721", log); err != nil {
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
func (_L1erc721gateway *L1erc721gatewayFilterer) ParseBatchDepositERC721(log types.Log) (*L1erc721gatewayBatchDepositERC721, error) {
	event := new(L1erc721gatewayBatchDepositERC721)
	if err := _L1erc721gateway.contract.UnpackLog(event, "BatchDepositERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1erc721gatewayBatchRefundERC721Iterator is returned from FilterBatchRefundERC721 and is used to iterate over the raw logs and unpacked data for BatchRefundERC721 events raised by the L1erc721gateway contract.
type L1erc721gatewayBatchRefundERC721Iterator struct {
	Event *L1erc721gatewayBatchRefundERC721 // Event containing the contract specifics and raw log

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
func (it *L1erc721gatewayBatchRefundERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1erc721gatewayBatchRefundERC721)
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
		it.Event = new(L1erc721gatewayBatchRefundERC721)
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
func (it *L1erc721gatewayBatchRefundERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1erc721gatewayBatchRefundERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1erc721gatewayBatchRefundERC721 represents a BatchRefundERC721 event raised by the L1erc721gateway contract.
type L1erc721gatewayBatchRefundERC721 struct {
	Token     common.Address
	Recipient common.Address
	TokenIds  []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBatchRefundERC721 is a free log retrieval operation binding the contract event 0x998a3ef0a23771412ff48d871a2288502a89da39c5db04a2a66e5eb85586cc22.
//
// Solidity: event BatchRefundERC721(address indexed token, address indexed recipient, uint256[] tokenIds)
func (_L1erc721gateway *L1erc721gatewayFilterer) FilterBatchRefundERC721(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*L1erc721gatewayBatchRefundERC721Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1erc721gateway.contract.FilterLogs(opts, "BatchRefundERC721", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &L1erc721gatewayBatchRefundERC721Iterator{contract: _L1erc721gateway.contract, event: "BatchRefundERC721", logs: logs, sub: sub}, nil
}

// WatchBatchRefundERC721 is a free log subscription operation binding the contract event 0x998a3ef0a23771412ff48d871a2288502a89da39c5db04a2a66e5eb85586cc22.
//
// Solidity: event BatchRefundERC721(address indexed token, address indexed recipient, uint256[] tokenIds)
func (_L1erc721gateway *L1erc721gatewayFilterer) WatchBatchRefundERC721(opts *bind.WatchOpts, sink chan<- *L1erc721gatewayBatchRefundERC721, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1erc721gateway.contract.WatchLogs(opts, "BatchRefundERC721", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1erc721gatewayBatchRefundERC721)
				if err := _L1erc721gateway.contract.UnpackLog(event, "BatchRefundERC721", log); err != nil {
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
func (_L1erc721gateway *L1erc721gatewayFilterer) ParseBatchRefundERC721(log types.Log) (*L1erc721gatewayBatchRefundERC721, error) {
	event := new(L1erc721gatewayBatchRefundERC721)
	if err := _L1erc721gateway.contract.UnpackLog(event, "BatchRefundERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1erc721gatewayDepositERC721Iterator is returned from FilterDepositERC721 and is used to iterate over the raw logs and unpacked data for DepositERC721 events raised by the L1erc721gateway contract.
type L1erc721gatewayDepositERC721Iterator struct {
	Event *L1erc721gatewayDepositERC721 // Event containing the contract specifics and raw log

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
func (it *L1erc721gatewayDepositERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1erc721gatewayDepositERC721)
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
		it.Event = new(L1erc721gatewayDepositERC721)
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
func (it *L1erc721gatewayDepositERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1erc721gatewayDepositERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1erc721gatewayDepositERC721 represents a DepositERC721 event raised by the L1erc721gateway contract.
type L1erc721gatewayDepositERC721 struct {
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
func (_L1erc721gateway *L1erc721gatewayFilterer) FilterDepositERC721(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*L1erc721gatewayDepositERC721Iterator, error) {

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

	logs, sub, err := _L1erc721gateway.contract.FilterLogs(opts, "DepositERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &L1erc721gatewayDepositERC721Iterator{contract: _L1erc721gateway.contract, event: "DepositERC721", logs: logs, sub: sub}, nil
}

// WatchDepositERC721 is a free log subscription operation binding the contract event 0xfc1d17c06ff1e4678321cc30660a73f3f1436df8195108a288d3159a961febec.
//
// Solidity: event DepositERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId)
func (_L1erc721gateway *L1erc721gatewayFilterer) WatchDepositERC721(opts *bind.WatchOpts, sink chan<- *L1erc721gatewayDepositERC721, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1erc721gateway.contract.WatchLogs(opts, "DepositERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1erc721gatewayDepositERC721)
				if err := _L1erc721gateway.contract.UnpackLog(event, "DepositERC721", log); err != nil {
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
func (_L1erc721gateway *L1erc721gatewayFilterer) ParseDepositERC721(log types.Log) (*L1erc721gatewayDepositERC721, error) {
	event := new(L1erc721gatewayDepositERC721)
	if err := _L1erc721gateway.contract.UnpackLog(event, "DepositERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1erc721gatewayFinalizeBatchWithdrawERC721Iterator is returned from FilterFinalizeBatchWithdrawERC721 and is used to iterate over the raw logs and unpacked data for FinalizeBatchWithdrawERC721 events raised by the L1erc721gateway contract.
type L1erc721gatewayFinalizeBatchWithdrawERC721Iterator struct {
	Event *L1erc721gatewayFinalizeBatchWithdrawERC721 // Event containing the contract specifics and raw log

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
func (it *L1erc721gatewayFinalizeBatchWithdrawERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1erc721gatewayFinalizeBatchWithdrawERC721)
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
		it.Event = new(L1erc721gatewayFinalizeBatchWithdrawERC721)
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
func (it *L1erc721gatewayFinalizeBatchWithdrawERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1erc721gatewayFinalizeBatchWithdrawERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1erc721gatewayFinalizeBatchWithdrawERC721 represents a FinalizeBatchWithdrawERC721 event raised by the L1erc721gateway contract.
type L1erc721gatewayFinalizeBatchWithdrawERC721 struct {
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
func (_L1erc721gateway *L1erc721gatewayFilterer) FilterFinalizeBatchWithdrawERC721(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*L1erc721gatewayFinalizeBatchWithdrawERC721Iterator, error) {

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

	logs, sub, err := _L1erc721gateway.contract.FilterLogs(opts, "FinalizeBatchWithdrawERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &L1erc721gatewayFinalizeBatchWithdrawERC721Iterator{contract: _L1erc721gateway.contract, event: "FinalizeBatchWithdrawERC721", logs: logs, sub: sub}, nil
}

// WatchFinalizeBatchWithdrawERC721 is a free log subscription operation binding the contract event 0x9b8e51c8f180115b421b26c9042287d6bf95e0ce9c0c5434784e2af3d0b9de7d.
//
// Solidity: event FinalizeBatchWithdrawERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256[] _tokenIds)
func (_L1erc721gateway *L1erc721gatewayFilterer) WatchFinalizeBatchWithdrawERC721(opts *bind.WatchOpts, sink chan<- *L1erc721gatewayFinalizeBatchWithdrawERC721, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1erc721gateway.contract.WatchLogs(opts, "FinalizeBatchWithdrawERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1erc721gatewayFinalizeBatchWithdrawERC721)
				if err := _L1erc721gateway.contract.UnpackLog(event, "FinalizeBatchWithdrawERC721", log); err != nil {
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
func (_L1erc721gateway *L1erc721gatewayFilterer) ParseFinalizeBatchWithdrawERC721(log types.Log) (*L1erc721gatewayFinalizeBatchWithdrawERC721, error) {
	event := new(L1erc721gatewayFinalizeBatchWithdrawERC721)
	if err := _L1erc721gateway.contract.UnpackLog(event, "FinalizeBatchWithdrawERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1erc721gatewayFinalizeWithdrawERC721Iterator is returned from FilterFinalizeWithdrawERC721 and is used to iterate over the raw logs and unpacked data for FinalizeWithdrawERC721 events raised by the L1erc721gateway contract.
type L1erc721gatewayFinalizeWithdrawERC721Iterator struct {
	Event *L1erc721gatewayFinalizeWithdrawERC721 // Event containing the contract specifics and raw log

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
func (it *L1erc721gatewayFinalizeWithdrawERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1erc721gatewayFinalizeWithdrawERC721)
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
		it.Event = new(L1erc721gatewayFinalizeWithdrawERC721)
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
func (it *L1erc721gatewayFinalizeWithdrawERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1erc721gatewayFinalizeWithdrawERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1erc721gatewayFinalizeWithdrawERC721 represents a FinalizeWithdrawERC721 event raised by the L1erc721gateway contract.
type L1erc721gatewayFinalizeWithdrawERC721 struct {
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
func (_L1erc721gateway *L1erc721gatewayFilterer) FilterFinalizeWithdrawERC721(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*L1erc721gatewayFinalizeWithdrawERC721Iterator, error) {

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

	logs, sub, err := _L1erc721gateway.contract.FilterLogs(opts, "FinalizeWithdrawERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &L1erc721gatewayFinalizeWithdrawERC721Iterator{contract: _L1erc721gateway.contract, event: "FinalizeWithdrawERC721", logs: logs, sub: sub}, nil
}

// WatchFinalizeWithdrawERC721 is a free log subscription operation binding the contract event 0xacdbfefc030b5ccccd5f60ca6d9ca371c6d6d6956fe16ebe10f81920198206e9.
//
// Solidity: event FinalizeWithdrawERC721(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _tokenId)
func (_L1erc721gateway *L1erc721gatewayFilterer) WatchFinalizeWithdrawERC721(opts *bind.WatchOpts, sink chan<- *L1erc721gatewayFinalizeWithdrawERC721, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1erc721gateway.contract.WatchLogs(opts, "FinalizeWithdrawERC721", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1erc721gatewayFinalizeWithdrawERC721)
				if err := _L1erc721gateway.contract.UnpackLog(event, "FinalizeWithdrawERC721", log); err != nil {
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
func (_L1erc721gateway *L1erc721gatewayFilterer) ParseFinalizeWithdrawERC721(log types.Log) (*L1erc721gatewayFinalizeWithdrawERC721, error) {
	event := new(L1erc721gatewayFinalizeWithdrawERC721)
	if err := _L1erc721gateway.contract.UnpackLog(event, "FinalizeWithdrawERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1erc721gatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1erc721gateway contract.
type L1erc721gatewayInitializedIterator struct {
	Event *L1erc721gatewayInitialized // Event containing the contract specifics and raw log

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
func (it *L1erc721gatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1erc721gatewayInitialized)
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
		it.Event = new(L1erc721gatewayInitialized)
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
func (it *L1erc721gatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1erc721gatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1erc721gatewayInitialized represents a Initialized event raised by the L1erc721gateway contract.
type L1erc721gatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1erc721gateway *L1erc721gatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1erc721gatewayInitializedIterator, error) {

	logs, sub, err := _L1erc721gateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1erc721gatewayInitializedIterator{contract: _L1erc721gateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1erc721gateway *L1erc721gatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1erc721gatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _L1erc721gateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1erc721gatewayInitialized)
				if err := _L1erc721gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L1erc721gateway *L1erc721gatewayFilterer) ParseInitialized(log types.Log) (*L1erc721gatewayInitialized, error) {
	event := new(L1erc721gatewayInitialized)
	if err := _L1erc721gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1erc721gatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1erc721gateway contract.
type L1erc721gatewayOwnershipTransferredIterator struct {
	Event *L1erc721gatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L1erc721gatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1erc721gatewayOwnershipTransferred)
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
		it.Event = new(L1erc721gatewayOwnershipTransferred)
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
func (it *L1erc721gatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1erc721gatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1erc721gatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L1erc721gateway contract.
type L1erc721gatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1erc721gateway *L1erc721gatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1erc721gatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1erc721gateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1erc721gatewayOwnershipTransferredIterator{contract: _L1erc721gateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1erc721gateway *L1erc721gatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1erc721gatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1erc721gateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1erc721gatewayOwnershipTransferred)
				if err := _L1erc721gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L1erc721gateway *L1erc721gatewayFilterer) ParseOwnershipTransferred(log types.Log) (*L1erc721gatewayOwnershipTransferred, error) {
	event := new(L1erc721gatewayOwnershipTransferred)
	if err := _L1erc721gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1erc721gatewayRefundERC721Iterator is returned from FilterRefundERC721 and is used to iterate over the raw logs and unpacked data for RefundERC721 events raised by the L1erc721gateway contract.
type L1erc721gatewayRefundERC721Iterator struct {
	Event *L1erc721gatewayRefundERC721 // Event containing the contract specifics and raw log

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
func (it *L1erc721gatewayRefundERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1erc721gatewayRefundERC721)
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
		it.Event = new(L1erc721gatewayRefundERC721)
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
func (it *L1erc721gatewayRefundERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1erc721gatewayRefundERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1erc721gatewayRefundERC721 represents a RefundERC721 event raised by the L1erc721gateway contract.
type L1erc721gatewayRefundERC721 struct {
	Token     common.Address
	Recipient common.Address
	TokenId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundERC721 is a free log retrieval operation binding the contract event 0xb9a838365634e4fb87a9333edf0ea86f82836e361b311a125aefd14135581208.
//
// Solidity: event RefundERC721(address indexed token, address indexed recipient, uint256 tokenId)
func (_L1erc721gateway *L1erc721gatewayFilterer) FilterRefundERC721(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*L1erc721gatewayRefundERC721Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1erc721gateway.contract.FilterLogs(opts, "RefundERC721", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &L1erc721gatewayRefundERC721Iterator{contract: _L1erc721gateway.contract, event: "RefundERC721", logs: logs, sub: sub}, nil
}

// WatchRefundERC721 is a free log subscription operation binding the contract event 0xb9a838365634e4fb87a9333edf0ea86f82836e361b311a125aefd14135581208.
//
// Solidity: event RefundERC721(address indexed token, address indexed recipient, uint256 tokenId)
func (_L1erc721gateway *L1erc721gatewayFilterer) WatchRefundERC721(opts *bind.WatchOpts, sink chan<- *L1erc721gatewayRefundERC721, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _L1erc721gateway.contract.WatchLogs(opts, "RefundERC721", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1erc721gatewayRefundERC721)
				if err := _L1erc721gateway.contract.UnpackLog(event, "RefundERC721", log); err != nil {
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
func (_L1erc721gateway *L1erc721gatewayFilterer) ParseRefundERC721(log types.Log) (*L1erc721gatewayRefundERC721, error) {
	event := new(L1erc721gatewayRefundERC721)
	if err := _L1erc721gateway.contract.UnpackLog(event, "RefundERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1erc721gatewayUpdateTokenMappingIterator is returned from FilterUpdateTokenMapping and is used to iterate over the raw logs and unpacked data for UpdateTokenMapping events raised by the L1erc721gateway contract.
type L1erc721gatewayUpdateTokenMappingIterator struct {
	Event *L1erc721gatewayUpdateTokenMapping // Event containing the contract specifics and raw log

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
func (it *L1erc721gatewayUpdateTokenMappingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1erc721gatewayUpdateTokenMapping)
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
		it.Event = new(L1erc721gatewayUpdateTokenMapping)
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
func (it *L1erc721gatewayUpdateTokenMappingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1erc721gatewayUpdateTokenMappingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1erc721gatewayUpdateTokenMapping represents a UpdateTokenMapping event raised by the L1erc721gateway contract.
type L1erc721gatewayUpdateTokenMapping struct {
	L1Token    common.Address
	OldL2Token common.Address
	NewL2Token common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenMapping is a free log retrieval operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l1Token, address indexed oldL2Token, address indexed newL2Token)
func (_L1erc721gateway *L1erc721gatewayFilterer) FilterUpdateTokenMapping(opts *bind.FilterOpts, l1Token []common.Address, oldL2Token []common.Address, newL2Token []common.Address) (*L1erc721gatewayUpdateTokenMappingIterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var oldL2TokenRule []interface{}
	for _, oldL2TokenItem := range oldL2Token {
		oldL2TokenRule = append(oldL2TokenRule, oldL2TokenItem)
	}
	var newL2TokenRule []interface{}
	for _, newL2TokenItem := range newL2Token {
		newL2TokenRule = append(newL2TokenRule, newL2TokenItem)
	}

	logs, sub, err := _L1erc721gateway.contract.FilterLogs(opts, "UpdateTokenMapping", l1TokenRule, oldL2TokenRule, newL2TokenRule)
	if err != nil {
		return nil, err
	}
	return &L1erc721gatewayUpdateTokenMappingIterator{contract: _L1erc721gateway.contract, event: "UpdateTokenMapping", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenMapping is a free log subscription operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l1Token, address indexed oldL2Token, address indexed newL2Token)
func (_L1erc721gateway *L1erc721gatewayFilterer) WatchUpdateTokenMapping(opts *bind.WatchOpts, sink chan<- *L1erc721gatewayUpdateTokenMapping, l1Token []common.Address, oldL2Token []common.Address, newL2Token []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var oldL2TokenRule []interface{}
	for _, oldL2TokenItem := range oldL2Token {
		oldL2TokenRule = append(oldL2TokenRule, oldL2TokenItem)
	}
	var newL2TokenRule []interface{}
	for _, newL2TokenItem := range newL2Token {
		newL2TokenRule = append(newL2TokenRule, newL2TokenItem)
	}

	logs, sub, err := _L1erc721gateway.contract.WatchLogs(opts, "UpdateTokenMapping", l1TokenRule, oldL2TokenRule, newL2TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1erc721gatewayUpdateTokenMapping)
				if err := _L1erc721gateway.contract.UnpackLog(event, "UpdateTokenMapping", log); err != nil {
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
// Solidity: event UpdateTokenMapping(address indexed l1Token, address indexed oldL2Token, address indexed newL2Token)
func (_L1erc721gateway *L1erc721gatewayFilterer) ParseUpdateTokenMapping(log types.Log) (*L1erc721gatewayUpdateTokenMapping, error) {
	event := new(L1erc721gatewayUpdateTokenMapping)
	if err := _L1erc721gateway.contract.UnpackLog(event, "UpdateTokenMapping", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
