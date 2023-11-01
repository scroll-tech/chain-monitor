// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2erc1155gateway

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

// L2erc1155gatewayMetaData contains all meta data concerning the L2erc1155gateway contract.
var L2erc1155gatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"BatchWithdrawERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"FinalizeBatchDepositERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FinalizeDepositERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldL1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newL1Token\",\"type\":\"address\"}],\"name\":\"UpdateTokenMapping\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC1155\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"batchWithdrawERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"batchWithdrawERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"finalizeBatchDepositERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"finalizeDepositERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"updateTokenMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// L2erc1155gatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L2erc1155gatewayMetaData.ABI instead.
var L2erc1155gatewayABI = L2erc1155gatewayMetaData.ABI

// L2erc1155gateway is an auto generated Go binding around an Ethereum contract.
type L2erc1155gateway struct {
	L2erc1155gatewayCaller     // Read-only binding to the contract
	L2erc1155gatewayTransactor // Write-only binding to the contract
	L2erc1155gatewayFilterer   // Log filterer for contract events
}

// L2erc1155gatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2erc1155gatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2erc1155gatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2erc1155gatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2erc1155gatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2erc1155gatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2erc1155gatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2erc1155gatewaySession struct {
	Contract     *L2erc1155gateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2erc1155gatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2erc1155gatewayCallerSession struct {
	Contract *L2erc1155gatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// L2erc1155gatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2erc1155gatewayTransactorSession struct {
	Contract     *L2erc1155gatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// L2erc1155gatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2erc1155gatewayRaw struct {
	Contract *L2erc1155gateway // Generic contract binding to access the raw methods on
}

// L2erc1155gatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2erc1155gatewayCallerRaw struct {
	Contract *L2erc1155gatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L2erc1155gatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2erc1155gatewayTransactorRaw struct {
	Contract *L2erc1155gatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2erc1155gateway creates a new instance of L2erc1155gateway, bound to a specific deployed contract.
func NewL2erc1155gateway(address common.Address, backend bind.ContractBackend) (*L2erc1155gateway, error) {
	contract, err := bindL2erc1155gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2erc1155gateway{L2erc1155gatewayCaller: L2erc1155gatewayCaller{contract: contract}, L2erc1155gatewayTransactor: L2erc1155gatewayTransactor{contract: contract}, L2erc1155gatewayFilterer: L2erc1155gatewayFilterer{contract: contract}}, nil
}

// NewL2erc1155gatewayCaller creates a new read-only instance of L2erc1155gateway, bound to a specific deployed contract.
func NewL2erc1155gatewayCaller(address common.Address, caller bind.ContractCaller) (*L2erc1155gatewayCaller, error) {
	contract, err := bindL2erc1155gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2erc1155gatewayCaller{contract: contract}, nil
}

// NewL2erc1155gatewayTransactor creates a new write-only instance of L2erc1155gateway, bound to a specific deployed contract.
func NewL2erc1155gatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L2erc1155gatewayTransactor, error) {
	contract, err := bindL2erc1155gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2erc1155gatewayTransactor{contract: contract}, nil
}

// NewL2erc1155gatewayFilterer creates a new log filterer instance of L2erc1155gateway, bound to a specific deployed contract.
func NewL2erc1155gatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L2erc1155gatewayFilterer, error) {
	contract, err := bindL2erc1155gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2erc1155gatewayFilterer{contract: contract}, nil
}

// bindL2erc1155gateway binds a generic wrapper to an already deployed contract.
func bindL2erc1155gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2erc1155gatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2erc1155gateway *L2erc1155gatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2erc1155gateway.Contract.L2erc1155gatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2erc1155gateway *L2erc1155gatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.L2erc1155gatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2erc1155gateway *L2erc1155gatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.L2erc1155gatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2erc1155gateway *L2erc1155gatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2erc1155gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2erc1155gateway *L2erc1155gatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2erc1155gateway *L2erc1155gatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.contract.Transact(opts, method, params...)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2erc1155gateway *L2erc1155gatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2erc1155gateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2erc1155gateway *L2erc1155gatewaySession) Counterpart() (common.Address, error) {
	return _L2erc1155gateway.Contract.Counterpart(&_L2erc1155gateway.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2erc1155gateway *L2erc1155gatewayCallerSession) Counterpart() (common.Address, error) {
	return _L2erc1155gateway.Contract.Counterpart(&_L2erc1155gateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2erc1155gateway *L2erc1155gatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2erc1155gateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2erc1155gateway *L2erc1155gatewaySession) Messenger() (common.Address, error) {
	return _L2erc1155gateway.Contract.Messenger(&_L2erc1155gateway.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2erc1155gateway *L2erc1155gatewayCallerSession) Messenger() (common.Address, error) {
	return _L2erc1155gateway.Contract.Messenger(&_L2erc1155gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2erc1155gateway *L2erc1155gatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2erc1155gateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2erc1155gateway *L2erc1155gatewaySession) Owner() (common.Address, error) {
	return _L2erc1155gateway.Contract.Owner(&_L2erc1155gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2erc1155gateway *L2erc1155gatewayCallerSession) Owner() (common.Address, error) {
	return _L2erc1155gateway.Contract.Owner(&_L2erc1155gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2erc1155gateway *L2erc1155gatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2erc1155gateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2erc1155gateway *L2erc1155gatewaySession) Router() (common.Address, error) {
	return _L2erc1155gateway.Contract.Router(&_L2erc1155gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2erc1155gateway *L2erc1155gatewayCallerSession) Router() (common.Address, error) {
	return _L2erc1155gateway.Contract.Router(&_L2erc1155gateway.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2erc1155gateway *L2erc1155gatewayCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _L2erc1155gateway.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2erc1155gateway *L2erc1155gatewaySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L2erc1155gateway.Contract.SupportsInterface(&_L2erc1155gateway.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2erc1155gateway *L2erc1155gatewayCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L2erc1155gateway.Contract.SupportsInterface(&_L2erc1155gateway.CallOpts, interfaceId)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2erc1155gateway *L2erc1155gatewayCaller) TokenMapping(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2erc1155gateway.contract.Call(opts, &out, "tokenMapping", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2erc1155gateway *L2erc1155gatewaySession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _L2erc1155gateway.Contract.TokenMapping(&_L2erc1155gateway.CallOpts, arg0)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2erc1155gateway *L2erc1155gatewayCallerSession) TokenMapping(arg0 common.Address) (common.Address, error) {
	return _L2erc1155gateway.Contract.TokenMapping(&_L2erc1155gateway.CallOpts, arg0)
}

// BatchWithdrawERC1155 is a paid mutator transaction binding the contract method 0x48de03de.
//
// Solidity: function batchWithdrawERC1155(address _token, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactor) BatchWithdrawERC1155(opts *bind.TransactOpts, _token common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc1155gateway.contract.Transact(opts, "batchWithdrawERC1155", _token, _tokenIds, _amounts, _gasLimit)
}

// BatchWithdrawERC1155 is a paid mutator transaction binding the contract method 0x48de03de.
//
// Solidity: function batchWithdrawERC1155(address _token, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_L2erc1155gateway *L2erc1155gatewaySession) BatchWithdrawERC1155(_token common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.BatchWithdrawERC1155(&_L2erc1155gateway.TransactOpts, _token, _tokenIds, _amounts, _gasLimit)
}

// BatchWithdrawERC1155 is a paid mutator transaction binding the contract method 0x48de03de.
//
// Solidity: function batchWithdrawERC1155(address _token, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactorSession) BatchWithdrawERC1155(_token common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.BatchWithdrawERC1155(&_L2erc1155gateway.TransactOpts, _token, _tokenIds, _amounts, _gasLimit)
}

// BatchWithdrawERC11550 is a paid mutator transaction binding the contract method 0x8c23d5b2.
//
// Solidity: function batchWithdrawERC1155(address _token, address _to, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactor) BatchWithdrawERC11550(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc1155gateway.contract.Transact(opts, "batchWithdrawERC11550", _token, _to, _tokenIds, _amounts, _gasLimit)
}

// BatchWithdrawERC11550 is a paid mutator transaction binding the contract method 0x8c23d5b2.
//
// Solidity: function batchWithdrawERC1155(address _token, address _to, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_L2erc1155gateway *L2erc1155gatewaySession) BatchWithdrawERC11550(_token common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.BatchWithdrawERC11550(&_L2erc1155gateway.TransactOpts, _token, _to, _tokenIds, _amounts, _gasLimit)
}

// BatchWithdrawERC11550 is a paid mutator transaction binding the contract method 0x8c23d5b2.
//
// Solidity: function batchWithdrawERC1155(address _token, address _to, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactorSession) BatchWithdrawERC11550(_token common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.BatchWithdrawERC11550(&_L2erc1155gateway.TransactOpts, _token, _to, _tokenIds, _amounts, _gasLimit)
}

// FinalizeBatchDepositERC1155 is a paid mutator transaction binding the contract method 0xeaa72ad9.
//
// Solidity: function finalizeBatchDepositERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds, uint256[] _amounts) returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactor) FinalizeBatchDepositERC1155(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _L2erc1155gateway.contract.Transact(opts, "finalizeBatchDepositERC1155", _l1Token, _l2Token, _from, _to, _tokenIds, _amounts)
}

// FinalizeBatchDepositERC1155 is a paid mutator transaction binding the contract method 0xeaa72ad9.
//
// Solidity: function finalizeBatchDepositERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds, uint256[] _amounts) returns()
func (_L2erc1155gateway *L2erc1155gatewaySession) FinalizeBatchDepositERC1155(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.FinalizeBatchDepositERC1155(&_L2erc1155gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenIds, _amounts)
}

// FinalizeBatchDepositERC1155 is a paid mutator transaction binding the contract method 0xeaa72ad9.
//
// Solidity: function finalizeBatchDepositERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds, uint256[] _amounts) returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactorSession) FinalizeBatchDepositERC1155(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.FinalizeBatchDepositERC1155(&_L2erc1155gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenIds, _amounts)
}

// FinalizeDepositERC1155 is a paid mutator transaction binding the contract method 0x4764cc62.
//
// Solidity: function finalizeDepositERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId, uint256 _amount) returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactor) FinalizeDepositERC1155(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _L2erc1155gateway.contract.Transact(opts, "finalizeDepositERC1155", _l1Token, _l2Token, _from, _to, _tokenId, _amount)
}

// FinalizeDepositERC1155 is a paid mutator transaction binding the contract method 0x4764cc62.
//
// Solidity: function finalizeDepositERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId, uint256 _amount) returns()
func (_L2erc1155gateway *L2erc1155gatewaySession) FinalizeDepositERC1155(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.FinalizeDepositERC1155(&_L2erc1155gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenId, _amount)
}

// FinalizeDepositERC1155 is a paid mutator transaction binding the contract method 0x4764cc62.
//
// Solidity: function finalizeDepositERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId, uint256 _amount) returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactorSession) FinalizeDepositERC1155(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.FinalizeDepositERC1155(&_L2erc1155gateway.TransactOpts, _l1Token, _l2Token, _from, _to, _tokenId, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2erc1155gateway.contract.Transact(opts, "initialize", _counterpart, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L2erc1155gateway *L2erc1155gatewaySession) Initialize(_counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.Initialize(&_L2erc1155gateway.TransactOpts, _counterpart, _messenger)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactorSession) Initialize(_counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.Initialize(&_L2erc1155gateway.TransactOpts, _counterpart, _messenger)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_L2erc1155gateway *L2erc1155gatewayTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L2erc1155gateway.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_L2erc1155gateway *L2erc1155gatewaySession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.OnERC1155BatchReceived(&_L2erc1155gateway.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_L2erc1155gateway *L2erc1155gatewayTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.OnERC1155BatchReceived(&_L2erc1155gateway.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_L2erc1155gateway *L2erc1155gatewayTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L2erc1155gateway.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_L2erc1155gateway *L2erc1155gatewaySession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.OnERC1155Received(&_L2erc1155gateway.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_L2erc1155gateway *L2erc1155gatewayTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.OnERC1155Received(&_L2erc1155gateway.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2erc1155gateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2erc1155gateway *L2erc1155gatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.RenounceOwnership(&_L2erc1155gateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.RenounceOwnership(&_L2erc1155gateway.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2erc1155gateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2erc1155gateway *L2erc1155gatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.TransferOwnership(&_L2erc1155gateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.TransferOwnership(&_L2erc1155gateway.TransactOpts, newOwner)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactor) UpdateTokenMapping(opts *bind.TransactOpts, _l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2erc1155gateway.contract.Transact(opts, "updateTokenMapping", _l2Token, _l1Token)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2erc1155gateway *L2erc1155gatewaySession) UpdateTokenMapping(_l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.UpdateTokenMapping(&_L2erc1155gateway.TransactOpts, _l2Token, _l1Token)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactorSession) UpdateTokenMapping(_l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.UpdateTokenMapping(&_L2erc1155gateway.TransactOpts, _l2Token, _l1Token)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x0f2da080.
//
// Solidity: function withdrawERC1155(address _token, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactor) WithdrawERC1155(opts *bind.TransactOpts, _token common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc1155gateway.contract.Transact(opts, "withdrawERC1155", _token, _tokenId, _amount, _gasLimit)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x0f2da080.
//
// Solidity: function withdrawERC1155(address _token, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2erc1155gateway *L2erc1155gatewaySession) WithdrawERC1155(_token common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.WithdrawERC1155(&_L2erc1155gateway.TransactOpts, _token, _tokenId, _amount, _gasLimit)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x0f2da080.
//
// Solidity: function withdrawERC1155(address _token, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactorSession) WithdrawERC1155(_token common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.WithdrawERC1155(&_L2erc1155gateway.TransactOpts, _token, _tokenId, _amount, _gasLimit)
}

// WithdrawERC11550 is a paid mutator transaction binding the contract method 0x21fedfc9.
//
// Solidity: function withdrawERC1155(address _token, address _to, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactor) WithdrawERC11550(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc1155gateway.contract.Transact(opts, "withdrawERC11550", _token, _to, _tokenId, _amount, _gasLimit)
}

// WithdrawERC11550 is a paid mutator transaction binding the contract method 0x21fedfc9.
//
// Solidity: function withdrawERC1155(address _token, address _to, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2erc1155gateway *L2erc1155gatewaySession) WithdrawERC11550(_token common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.WithdrawERC11550(&_L2erc1155gateway.TransactOpts, _token, _to, _tokenId, _amount, _gasLimit)
}

// WithdrawERC11550 is a paid mutator transaction binding the contract method 0x21fedfc9.
//
// Solidity: function withdrawERC1155(address _token, address _to, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2erc1155gateway *L2erc1155gatewayTransactorSession) WithdrawERC11550(_token common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2erc1155gateway.Contract.WithdrawERC11550(&_L2erc1155gateway.TransactOpts, _token, _to, _tokenId, _amount, _gasLimit)
}

// L2erc1155gatewayBatchWithdrawERC1155Iterator is returned from FilterBatchWithdrawERC1155 and is used to iterate over the raw logs and unpacked data for BatchWithdrawERC1155 events raised by the L2erc1155gateway contract.
type L2erc1155gatewayBatchWithdrawERC1155Iterator struct {
	Event *L2erc1155gatewayBatchWithdrawERC1155 // Event containing the contract specifics and raw log

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
func (it *L2erc1155gatewayBatchWithdrawERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2erc1155gatewayBatchWithdrawERC1155)
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
		it.Event = new(L2erc1155gatewayBatchWithdrawERC1155)
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
func (it *L2erc1155gatewayBatchWithdrawERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2erc1155gatewayBatchWithdrawERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2erc1155gatewayBatchWithdrawERC1155 represents a BatchWithdrawERC1155 event raised by the L2erc1155gateway contract.
type L2erc1155gatewayBatchWithdrawERC1155 struct {
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
func (_L2erc1155gateway *L2erc1155gatewayFilterer) FilterBatchWithdrawERC1155(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2erc1155gatewayBatchWithdrawERC1155Iterator, error) {

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

	logs, sub, err := _L2erc1155gateway.contract.FilterLogs(opts, "BatchWithdrawERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2erc1155gatewayBatchWithdrawERC1155Iterator{contract: _L2erc1155gateway.contract, event: "BatchWithdrawERC1155", logs: logs, sub: sub}, nil
}

// WatchBatchWithdrawERC1155 is a free log subscription operation binding the contract event 0x5d2d5d4cdbf7b115e43f0b9986644dd8b9514b10be6a019ab6a4a87f12290970.
//
// Solidity: event BatchWithdrawERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds, uint256[] amounts)
func (_L2erc1155gateway *L2erc1155gatewayFilterer) WatchBatchWithdrawERC1155(opts *bind.WatchOpts, sink chan<- *L2erc1155gatewayBatchWithdrawERC1155, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2erc1155gateway.contract.WatchLogs(opts, "BatchWithdrawERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2erc1155gatewayBatchWithdrawERC1155)
				if err := _L2erc1155gateway.contract.UnpackLog(event, "BatchWithdrawERC1155", log); err != nil {
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
func (_L2erc1155gateway *L2erc1155gatewayFilterer) ParseBatchWithdrawERC1155(log types.Log) (*L2erc1155gatewayBatchWithdrawERC1155, error) {
	event := new(L2erc1155gatewayBatchWithdrawERC1155)
	if err := _L2erc1155gateway.contract.UnpackLog(event, "BatchWithdrawERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2erc1155gatewayFinalizeBatchDepositERC1155Iterator is returned from FilterFinalizeBatchDepositERC1155 and is used to iterate over the raw logs and unpacked data for FinalizeBatchDepositERC1155 events raised by the L2erc1155gateway contract.
type L2erc1155gatewayFinalizeBatchDepositERC1155Iterator struct {
	Event *L2erc1155gatewayFinalizeBatchDepositERC1155 // Event containing the contract specifics and raw log

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
func (it *L2erc1155gatewayFinalizeBatchDepositERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2erc1155gatewayFinalizeBatchDepositERC1155)
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
		it.Event = new(L2erc1155gatewayFinalizeBatchDepositERC1155)
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
func (it *L2erc1155gatewayFinalizeBatchDepositERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2erc1155gatewayFinalizeBatchDepositERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2erc1155gatewayFinalizeBatchDepositERC1155 represents a FinalizeBatchDepositERC1155 event raised by the L2erc1155gateway contract.
type L2erc1155gatewayFinalizeBatchDepositERC1155 struct {
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
func (_L2erc1155gateway *L2erc1155gatewayFilterer) FilterFinalizeBatchDepositERC1155(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2erc1155gatewayFinalizeBatchDepositERC1155Iterator, error) {

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

	logs, sub, err := _L2erc1155gateway.contract.FilterLogs(opts, "FinalizeBatchDepositERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2erc1155gatewayFinalizeBatchDepositERC1155Iterator{contract: _L2erc1155gateway.contract, event: "FinalizeBatchDepositERC1155", logs: logs, sub: sub}, nil
}

// WatchFinalizeBatchDepositERC1155 is a free log subscription operation binding the contract event 0xf07745bfeb45fb1184165136e9148689adf57ba578a5b90dde949f26066b7756.
//
// Solidity: event FinalizeBatchDepositERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256[] tokenIds, uint256[] amounts)
func (_L2erc1155gateway *L2erc1155gatewayFilterer) WatchFinalizeBatchDepositERC1155(opts *bind.WatchOpts, sink chan<- *L2erc1155gatewayFinalizeBatchDepositERC1155, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2erc1155gateway.contract.WatchLogs(opts, "FinalizeBatchDepositERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2erc1155gatewayFinalizeBatchDepositERC1155)
				if err := _L2erc1155gateway.contract.UnpackLog(event, "FinalizeBatchDepositERC1155", log); err != nil {
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
func (_L2erc1155gateway *L2erc1155gatewayFilterer) ParseFinalizeBatchDepositERC1155(log types.Log) (*L2erc1155gatewayFinalizeBatchDepositERC1155, error) {
	event := new(L2erc1155gatewayFinalizeBatchDepositERC1155)
	if err := _L2erc1155gateway.contract.UnpackLog(event, "FinalizeBatchDepositERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2erc1155gatewayFinalizeDepositERC1155Iterator is returned from FilterFinalizeDepositERC1155 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC1155 events raised by the L2erc1155gateway contract.
type L2erc1155gatewayFinalizeDepositERC1155Iterator struct {
	Event *L2erc1155gatewayFinalizeDepositERC1155 // Event containing the contract specifics and raw log

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
func (it *L2erc1155gatewayFinalizeDepositERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2erc1155gatewayFinalizeDepositERC1155)
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
		it.Event = new(L2erc1155gatewayFinalizeDepositERC1155)
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
func (it *L2erc1155gatewayFinalizeDepositERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2erc1155gatewayFinalizeDepositERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2erc1155gatewayFinalizeDepositERC1155 represents a FinalizeDepositERC1155 event raised by the L2erc1155gateway contract.
type L2erc1155gatewayFinalizeDepositERC1155 struct {
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
func (_L2erc1155gateway *L2erc1155gatewayFilterer) FilterFinalizeDepositERC1155(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2erc1155gatewayFinalizeDepositERC1155Iterator, error) {

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

	logs, sub, err := _L2erc1155gateway.contract.FilterLogs(opts, "FinalizeDepositERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2erc1155gatewayFinalizeDepositERC1155Iterator{contract: _L2erc1155gateway.contract, event: "FinalizeDepositERC1155", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC1155 is a free log subscription operation binding the contract event 0x5399dc7b86d085e50a28946dbc213966bb7a7ac78d312aedd6018c791ad6cef9.
//
// Solidity: event FinalizeDepositERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId, uint256 amount)
func (_L2erc1155gateway *L2erc1155gatewayFilterer) WatchFinalizeDepositERC1155(opts *bind.WatchOpts, sink chan<- *L2erc1155gatewayFinalizeDepositERC1155, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2erc1155gateway.contract.WatchLogs(opts, "FinalizeDepositERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2erc1155gatewayFinalizeDepositERC1155)
				if err := _L2erc1155gateway.contract.UnpackLog(event, "FinalizeDepositERC1155", log); err != nil {
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
func (_L2erc1155gateway *L2erc1155gatewayFilterer) ParseFinalizeDepositERC1155(log types.Log) (*L2erc1155gatewayFinalizeDepositERC1155, error) {
	event := new(L2erc1155gatewayFinalizeDepositERC1155)
	if err := _L2erc1155gateway.contract.UnpackLog(event, "FinalizeDepositERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2erc1155gatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2erc1155gateway contract.
type L2erc1155gatewayInitializedIterator struct {
	Event *L2erc1155gatewayInitialized // Event containing the contract specifics and raw log

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
func (it *L2erc1155gatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2erc1155gatewayInitialized)
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
		it.Event = new(L2erc1155gatewayInitialized)
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
func (it *L2erc1155gatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2erc1155gatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2erc1155gatewayInitialized represents a Initialized event raised by the L2erc1155gateway contract.
type L2erc1155gatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2erc1155gateway *L2erc1155gatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2erc1155gatewayInitializedIterator, error) {

	logs, sub, err := _L2erc1155gateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2erc1155gatewayInitializedIterator{contract: _L2erc1155gateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2erc1155gateway *L2erc1155gatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2erc1155gatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _L2erc1155gateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2erc1155gatewayInitialized)
				if err := _L2erc1155gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2erc1155gateway *L2erc1155gatewayFilterer) ParseInitialized(log types.Log) (*L2erc1155gatewayInitialized, error) {
	event := new(L2erc1155gatewayInitialized)
	if err := _L2erc1155gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2erc1155gatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2erc1155gateway contract.
type L2erc1155gatewayOwnershipTransferredIterator struct {
	Event *L2erc1155gatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2erc1155gatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2erc1155gatewayOwnershipTransferred)
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
		it.Event = new(L2erc1155gatewayOwnershipTransferred)
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
func (it *L2erc1155gatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2erc1155gatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2erc1155gatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L2erc1155gateway contract.
type L2erc1155gatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2erc1155gateway *L2erc1155gatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2erc1155gatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2erc1155gateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2erc1155gatewayOwnershipTransferredIterator{contract: _L2erc1155gateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2erc1155gateway *L2erc1155gatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2erc1155gatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2erc1155gateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2erc1155gatewayOwnershipTransferred)
				if err := _L2erc1155gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2erc1155gateway *L2erc1155gatewayFilterer) ParseOwnershipTransferred(log types.Log) (*L2erc1155gatewayOwnershipTransferred, error) {
	event := new(L2erc1155gatewayOwnershipTransferred)
	if err := _L2erc1155gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2erc1155gatewayUpdateTokenMappingIterator is returned from FilterUpdateTokenMapping and is used to iterate over the raw logs and unpacked data for UpdateTokenMapping events raised by the L2erc1155gateway contract.
type L2erc1155gatewayUpdateTokenMappingIterator struct {
	Event *L2erc1155gatewayUpdateTokenMapping // Event containing the contract specifics and raw log

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
func (it *L2erc1155gatewayUpdateTokenMappingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2erc1155gatewayUpdateTokenMapping)
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
		it.Event = new(L2erc1155gatewayUpdateTokenMapping)
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
func (it *L2erc1155gatewayUpdateTokenMappingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2erc1155gatewayUpdateTokenMappingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2erc1155gatewayUpdateTokenMapping represents a UpdateTokenMapping event raised by the L2erc1155gateway contract.
type L2erc1155gatewayUpdateTokenMapping struct {
	L2Token    common.Address
	OldL1Token common.Address
	NewL1Token common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenMapping is a free log retrieval operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l2Token, address indexed oldL1Token, address indexed newL1Token)
func (_L2erc1155gateway *L2erc1155gatewayFilterer) FilterUpdateTokenMapping(opts *bind.FilterOpts, l2Token []common.Address, oldL1Token []common.Address, newL1Token []common.Address) (*L2erc1155gatewayUpdateTokenMappingIterator, error) {

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

	logs, sub, err := _L2erc1155gateway.contract.FilterLogs(opts, "UpdateTokenMapping", l2TokenRule, oldL1TokenRule, newL1TokenRule)
	if err != nil {
		return nil, err
	}
	return &L2erc1155gatewayUpdateTokenMappingIterator{contract: _L2erc1155gateway.contract, event: "UpdateTokenMapping", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenMapping is a free log subscription operation binding the contract event 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
//
// Solidity: event UpdateTokenMapping(address indexed l2Token, address indexed oldL1Token, address indexed newL1Token)
func (_L2erc1155gateway *L2erc1155gatewayFilterer) WatchUpdateTokenMapping(opts *bind.WatchOpts, sink chan<- *L2erc1155gatewayUpdateTokenMapping, l2Token []common.Address, oldL1Token []common.Address, newL1Token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2erc1155gateway.contract.WatchLogs(opts, "UpdateTokenMapping", l2TokenRule, oldL1TokenRule, newL1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2erc1155gatewayUpdateTokenMapping)
				if err := _L2erc1155gateway.contract.UnpackLog(event, "UpdateTokenMapping", log); err != nil {
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
func (_L2erc1155gateway *L2erc1155gatewayFilterer) ParseUpdateTokenMapping(log types.Log) (*L2erc1155gatewayUpdateTokenMapping, error) {
	event := new(L2erc1155gatewayUpdateTokenMapping)
	if err := _L2erc1155gateway.contract.UnpackLog(event, "UpdateTokenMapping", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2erc1155gatewayWithdrawERC1155Iterator is returned from FilterWithdrawERC1155 and is used to iterate over the raw logs and unpacked data for WithdrawERC1155 events raised by the L2erc1155gateway contract.
type L2erc1155gatewayWithdrawERC1155Iterator struct {
	Event *L2erc1155gatewayWithdrawERC1155 // Event containing the contract specifics and raw log

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
func (it *L2erc1155gatewayWithdrawERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2erc1155gatewayWithdrawERC1155)
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
		it.Event = new(L2erc1155gatewayWithdrawERC1155)
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
func (it *L2erc1155gatewayWithdrawERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2erc1155gatewayWithdrawERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2erc1155gatewayWithdrawERC1155 represents a WithdrawERC1155 event raised by the L2erc1155gateway contract.
type L2erc1155gatewayWithdrawERC1155 struct {
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
func (_L2erc1155gateway *L2erc1155gatewayFilterer) FilterWithdrawERC1155(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2erc1155gatewayWithdrawERC1155Iterator, error) {

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

	logs, sub, err := _L2erc1155gateway.contract.FilterLogs(opts, "WithdrawERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2erc1155gatewayWithdrawERC1155Iterator{contract: _L2erc1155gateway.contract, event: "WithdrawERC1155", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC1155 is a free log subscription operation binding the contract event 0x1f9dcda7fce6f73a13055f044ffecaed2032a7a844e0a37a3eb8bbb17488d01a.
//
// Solidity: event WithdrawERC1155(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 tokenId, uint256 amount)
func (_L2erc1155gateway *L2erc1155gatewayFilterer) WatchWithdrawERC1155(opts *bind.WatchOpts, sink chan<- *L2erc1155gatewayWithdrawERC1155, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2erc1155gateway.contract.WatchLogs(opts, "WithdrawERC1155", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2erc1155gatewayWithdrawERC1155)
				if err := _L2erc1155gateway.contract.UnpackLog(event, "WithdrawERC1155", log); err != nil {
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
func (_L2erc1155gateway *L2erc1155gatewayFilterer) ParseWithdrawERC1155(log types.Log) (*L2erc1155gatewayWithdrawERC1155, error) {
	event := new(L2erc1155gatewayWithdrawERC1155)
	if err := _L2erc1155gateway.contract.UnpackLog(event, "WithdrawERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
