// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iscrollerc1155

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

// Iscrollerc1155MetaData contains all meta data concerning the Iscrollerc1155 contract.
var Iscrollerc1155MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"batchMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Iscrollerc1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use Iscrollerc1155MetaData.ABI instead.
var Iscrollerc1155ABI = Iscrollerc1155MetaData.ABI

// Iscrollerc1155 is an auto generated Go binding around an Ethereum contract.
type Iscrollerc1155 struct {
	Iscrollerc1155Caller     // Read-only binding to the contract
	Iscrollerc1155Transactor // Write-only binding to the contract
	Iscrollerc1155Filterer   // Log filterer for contract events
}

// Iscrollerc1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type Iscrollerc1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iscrollerc1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Iscrollerc1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iscrollerc1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Iscrollerc1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iscrollerc1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Iscrollerc1155Session struct {
	Contract     *Iscrollerc1155   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Iscrollerc1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Iscrollerc1155CallerSession struct {
	Contract *Iscrollerc1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Iscrollerc1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Iscrollerc1155TransactorSession struct {
	Contract     *Iscrollerc1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Iscrollerc1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type Iscrollerc1155Raw struct {
	Contract *Iscrollerc1155 // Generic contract binding to access the raw methods on
}

// Iscrollerc1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Iscrollerc1155CallerRaw struct {
	Contract *Iscrollerc1155Caller // Generic read-only contract binding to access the raw methods on
}

// Iscrollerc1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Iscrollerc1155TransactorRaw struct {
	Contract *Iscrollerc1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIscrollerc1155 creates a new instance of Iscrollerc1155, bound to a specific deployed contract.
func NewIscrollerc1155(address common.Address, backend bind.ContractBackend) (*Iscrollerc1155, error) {
	contract, err := bindIscrollerc1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iscrollerc1155{Iscrollerc1155Caller: Iscrollerc1155Caller{contract: contract}, Iscrollerc1155Transactor: Iscrollerc1155Transactor{contract: contract}, Iscrollerc1155Filterer: Iscrollerc1155Filterer{contract: contract}}, nil
}

// NewIscrollerc1155Caller creates a new read-only instance of Iscrollerc1155, bound to a specific deployed contract.
func NewIscrollerc1155Caller(address common.Address, caller bind.ContractCaller) (*Iscrollerc1155Caller, error) {
	contract, err := bindIscrollerc1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Iscrollerc1155Caller{contract: contract}, nil
}

// NewIscrollerc1155Transactor creates a new write-only instance of Iscrollerc1155, bound to a specific deployed contract.
func NewIscrollerc1155Transactor(address common.Address, transactor bind.ContractTransactor) (*Iscrollerc1155Transactor, error) {
	contract, err := bindIscrollerc1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Iscrollerc1155Transactor{contract: contract}, nil
}

// NewIscrollerc1155Filterer creates a new log filterer instance of Iscrollerc1155, bound to a specific deployed contract.
func NewIscrollerc1155Filterer(address common.Address, filterer bind.ContractFilterer) (*Iscrollerc1155Filterer, error) {
	contract, err := bindIscrollerc1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Iscrollerc1155Filterer{contract: contract}, nil
}

// bindIscrollerc1155 binds a generic wrapper to an already deployed contract.
func bindIscrollerc1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Iscrollerc1155MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iscrollerc1155 *Iscrollerc1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iscrollerc1155.Contract.Iscrollerc1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iscrollerc1155 *Iscrollerc1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iscrollerc1155.Contract.Iscrollerc1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iscrollerc1155 *Iscrollerc1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iscrollerc1155.Contract.Iscrollerc1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iscrollerc1155 *Iscrollerc1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iscrollerc1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iscrollerc1155 *Iscrollerc1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iscrollerc1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iscrollerc1155 *Iscrollerc1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iscrollerc1155.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Iscrollerc1155 *Iscrollerc1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Iscrollerc1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Iscrollerc1155 *Iscrollerc1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Iscrollerc1155.Contract.BalanceOf(&_Iscrollerc1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Iscrollerc1155 *Iscrollerc1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Iscrollerc1155.Contract.BalanceOf(&_Iscrollerc1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Iscrollerc1155 *Iscrollerc1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Iscrollerc1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Iscrollerc1155 *Iscrollerc1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Iscrollerc1155.Contract.BalanceOfBatch(&_Iscrollerc1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Iscrollerc1155 *Iscrollerc1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Iscrollerc1155.Contract.BalanceOfBatch(&_Iscrollerc1155.CallOpts, accounts, ids)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_Iscrollerc1155 *Iscrollerc1155Caller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iscrollerc1155.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_Iscrollerc1155 *Iscrollerc1155Session) Counterpart() (common.Address, error) {
	return _Iscrollerc1155.Contract.Counterpart(&_Iscrollerc1155.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_Iscrollerc1155 *Iscrollerc1155CallerSession) Counterpart() (common.Address, error) {
	return _Iscrollerc1155.Contract.Counterpart(&_Iscrollerc1155.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Iscrollerc1155 *Iscrollerc1155Caller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iscrollerc1155.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Iscrollerc1155 *Iscrollerc1155Session) Gateway() (common.Address, error) {
	return _Iscrollerc1155.Contract.Gateway(&_Iscrollerc1155.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Iscrollerc1155 *Iscrollerc1155CallerSession) Gateway() (common.Address, error) {
	return _Iscrollerc1155.Contract.Gateway(&_Iscrollerc1155.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Iscrollerc1155 *Iscrollerc1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Iscrollerc1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Iscrollerc1155 *Iscrollerc1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Iscrollerc1155.Contract.IsApprovedForAll(&_Iscrollerc1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Iscrollerc1155 *Iscrollerc1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Iscrollerc1155.Contract.IsApprovedForAll(&_Iscrollerc1155.CallOpts, account, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Iscrollerc1155 *Iscrollerc1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Iscrollerc1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Iscrollerc1155 *Iscrollerc1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Iscrollerc1155.Contract.SupportsInterface(&_Iscrollerc1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Iscrollerc1155 *Iscrollerc1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Iscrollerc1155.Contract.SupportsInterface(&_Iscrollerc1155.CallOpts, interfaceId)
}

// BatchBurn is a paid mutator transaction binding the contract method 0xf6eb127a.
//
// Solidity: function batchBurn(address _from, uint256[] _tokenIds, uint256[] _amounts) returns()
func (_Iscrollerc1155 *Iscrollerc1155Transactor) BatchBurn(opts *bind.TransactOpts, _from common.Address, _tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Iscrollerc1155.contract.Transact(opts, "batchBurn", _from, _tokenIds, _amounts)
}

// BatchBurn is a paid mutator transaction binding the contract method 0xf6eb127a.
//
// Solidity: function batchBurn(address _from, uint256[] _tokenIds, uint256[] _amounts) returns()
func (_Iscrollerc1155 *Iscrollerc1155Session) BatchBurn(_from common.Address, _tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Iscrollerc1155.Contract.BatchBurn(&_Iscrollerc1155.TransactOpts, _from, _tokenIds, _amounts)
}

// BatchBurn is a paid mutator transaction binding the contract method 0xf6eb127a.
//
// Solidity: function batchBurn(address _from, uint256[] _tokenIds, uint256[] _amounts) returns()
func (_Iscrollerc1155 *Iscrollerc1155TransactorSession) BatchBurn(_from common.Address, _tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _Iscrollerc1155.Contract.BatchBurn(&_Iscrollerc1155.TransactOpts, _from, _tokenIds, _amounts)
}

// BatchMint is a paid mutator transaction binding the contract method 0xb48ab8b6.
//
// Solidity: function batchMint(address _to, uint256[] _tokenIds, uint256[] _amounts, bytes _data) returns()
func (_Iscrollerc1155 *Iscrollerc1155Transactor) BatchMint(opts *bind.TransactOpts, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _Iscrollerc1155.contract.Transact(opts, "batchMint", _to, _tokenIds, _amounts, _data)
}

// BatchMint is a paid mutator transaction binding the contract method 0xb48ab8b6.
//
// Solidity: function batchMint(address _to, uint256[] _tokenIds, uint256[] _amounts, bytes _data) returns()
func (_Iscrollerc1155 *Iscrollerc1155Session) BatchMint(_to common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _Iscrollerc1155.Contract.BatchMint(&_Iscrollerc1155.TransactOpts, _to, _tokenIds, _amounts, _data)
}

// BatchMint is a paid mutator transaction binding the contract method 0xb48ab8b6.
//
// Solidity: function batchMint(address _to, uint256[] _tokenIds, uint256[] _amounts, bytes _data) returns()
func (_Iscrollerc1155 *Iscrollerc1155TransactorSession) BatchMint(_to common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _Iscrollerc1155.Contract.BatchMint(&_Iscrollerc1155.TransactOpts, _to, _tokenIds, _amounts, _data)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address _from, uint256 _tokenId, uint256 _amount) returns()
func (_Iscrollerc1155 *Iscrollerc1155Transactor) Burn(opts *bind.TransactOpts, _from common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Iscrollerc1155.contract.Transact(opts, "burn", _from, _tokenId, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address _from, uint256 _tokenId, uint256 _amount) returns()
func (_Iscrollerc1155 *Iscrollerc1155Session) Burn(_from common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Iscrollerc1155.Contract.Burn(&_Iscrollerc1155.TransactOpts, _from, _tokenId, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address _from, uint256 _tokenId, uint256 _amount) returns()
func (_Iscrollerc1155 *Iscrollerc1155TransactorSession) Burn(_from common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Iscrollerc1155.Contract.Burn(&_Iscrollerc1155.TransactOpts, _from, _tokenId, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address _to, uint256 _tokenId, uint256 _amount, bytes _data) returns()
func (_Iscrollerc1155 *Iscrollerc1155Transactor) Mint(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _Iscrollerc1155.contract.Transact(opts, "mint", _to, _tokenId, _amount, _data)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address _to, uint256 _tokenId, uint256 _amount, bytes _data) returns()
func (_Iscrollerc1155 *Iscrollerc1155Session) Mint(_to common.Address, _tokenId *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _Iscrollerc1155.Contract.Mint(&_Iscrollerc1155.TransactOpts, _to, _tokenId, _amount, _data)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address _to, uint256 _tokenId, uint256 _amount, bytes _data) returns()
func (_Iscrollerc1155 *Iscrollerc1155TransactorSession) Mint(_to common.Address, _tokenId *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _Iscrollerc1155.Contract.Mint(&_Iscrollerc1155.TransactOpts, _to, _tokenId, _amount, _data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Iscrollerc1155 *Iscrollerc1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Iscrollerc1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Iscrollerc1155 *Iscrollerc1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Iscrollerc1155.Contract.SafeBatchTransferFrom(&_Iscrollerc1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Iscrollerc1155 *Iscrollerc1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Iscrollerc1155.Contract.SafeBatchTransferFrom(&_Iscrollerc1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Iscrollerc1155 *Iscrollerc1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Iscrollerc1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Iscrollerc1155 *Iscrollerc1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Iscrollerc1155.Contract.SafeTransferFrom(&_Iscrollerc1155.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Iscrollerc1155 *Iscrollerc1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Iscrollerc1155.Contract.SafeTransferFrom(&_Iscrollerc1155.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Iscrollerc1155 *Iscrollerc1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Iscrollerc1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Iscrollerc1155 *Iscrollerc1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Iscrollerc1155.Contract.SetApprovalForAll(&_Iscrollerc1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Iscrollerc1155 *Iscrollerc1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Iscrollerc1155.Contract.SetApprovalForAll(&_Iscrollerc1155.TransactOpts, operator, approved)
}

// Iscrollerc1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Iscrollerc1155 contract.
type Iscrollerc1155ApprovalForAllIterator struct {
	Event *Iscrollerc1155ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Iscrollerc1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Iscrollerc1155ApprovalForAll)
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
		it.Event = new(Iscrollerc1155ApprovalForAll)
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
func (it *Iscrollerc1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Iscrollerc1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Iscrollerc1155ApprovalForAll represents a ApprovalForAll event raised by the Iscrollerc1155 contract.
type Iscrollerc1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Iscrollerc1155 *Iscrollerc1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*Iscrollerc1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Iscrollerc1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Iscrollerc1155ApprovalForAllIterator{contract: _Iscrollerc1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Iscrollerc1155 *Iscrollerc1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Iscrollerc1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Iscrollerc1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Iscrollerc1155ApprovalForAll)
				if err := _Iscrollerc1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Iscrollerc1155 *Iscrollerc1155Filterer) ParseApprovalForAll(log types.Log) (*Iscrollerc1155ApprovalForAll, error) {
	event := new(Iscrollerc1155ApprovalForAll)
	if err := _Iscrollerc1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Iscrollerc1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Iscrollerc1155 contract.
type Iscrollerc1155TransferBatchIterator struct {
	Event *Iscrollerc1155TransferBatch // Event containing the contract specifics and raw log

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
func (it *Iscrollerc1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Iscrollerc1155TransferBatch)
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
		it.Event = new(Iscrollerc1155TransferBatch)
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
func (it *Iscrollerc1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Iscrollerc1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Iscrollerc1155TransferBatch represents a TransferBatch event raised by the Iscrollerc1155 contract.
type Iscrollerc1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Iscrollerc1155 *Iscrollerc1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Iscrollerc1155TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Iscrollerc1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Iscrollerc1155TransferBatchIterator{contract: _Iscrollerc1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Iscrollerc1155 *Iscrollerc1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *Iscrollerc1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Iscrollerc1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Iscrollerc1155TransferBatch)
				if err := _Iscrollerc1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Iscrollerc1155 *Iscrollerc1155Filterer) ParseTransferBatch(log types.Log) (*Iscrollerc1155TransferBatch, error) {
	event := new(Iscrollerc1155TransferBatch)
	if err := _Iscrollerc1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Iscrollerc1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Iscrollerc1155 contract.
type Iscrollerc1155TransferSingleIterator struct {
	Event *Iscrollerc1155TransferSingle // Event containing the contract specifics and raw log

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
func (it *Iscrollerc1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Iscrollerc1155TransferSingle)
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
		it.Event = new(Iscrollerc1155TransferSingle)
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
func (it *Iscrollerc1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Iscrollerc1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Iscrollerc1155TransferSingle represents a TransferSingle event raised by the Iscrollerc1155 contract.
type Iscrollerc1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Iscrollerc1155 *Iscrollerc1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*Iscrollerc1155TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Iscrollerc1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Iscrollerc1155TransferSingleIterator{contract: _Iscrollerc1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Iscrollerc1155 *Iscrollerc1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *Iscrollerc1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Iscrollerc1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Iscrollerc1155TransferSingle)
				if err := _Iscrollerc1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Iscrollerc1155 *Iscrollerc1155Filterer) ParseTransferSingle(log types.Log) (*Iscrollerc1155TransferSingle, error) {
	event := new(Iscrollerc1155TransferSingle)
	if err := _Iscrollerc1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Iscrollerc1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Iscrollerc1155 contract.
type Iscrollerc1155URIIterator struct {
	Event *Iscrollerc1155URI // Event containing the contract specifics and raw log

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
func (it *Iscrollerc1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Iscrollerc1155URI)
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
		it.Event = new(Iscrollerc1155URI)
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
func (it *Iscrollerc1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Iscrollerc1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Iscrollerc1155URI represents a URI event raised by the Iscrollerc1155 contract.
type Iscrollerc1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Iscrollerc1155 *Iscrollerc1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*Iscrollerc1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Iscrollerc1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &Iscrollerc1155URIIterator{contract: _Iscrollerc1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Iscrollerc1155 *Iscrollerc1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *Iscrollerc1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Iscrollerc1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Iscrollerc1155URI)
				if err := _Iscrollerc1155.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Iscrollerc1155 *Iscrollerc1155Filterer) ParseURI(log types.Log) (*Iscrollerc1155URI, error) {
	event := new(Iscrollerc1155URI)
	if err := _Iscrollerc1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
