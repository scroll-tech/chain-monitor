// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iscrollerc20

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

// Iscrollerc20MetaData contains all meta data concerning the Iscrollerc20 contract.
var Iscrollerc20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Iscrollerc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Iscrollerc20MetaData.ABI instead.
var Iscrollerc20ABI = Iscrollerc20MetaData.ABI

// Iscrollerc20 is an auto generated Go binding around an Ethereum contract.
type Iscrollerc20 struct {
	Iscrollerc20Caller     // Read-only binding to the contract
	Iscrollerc20Transactor // Write-only binding to the contract
	Iscrollerc20Filterer   // Log filterer for contract events
}

// Iscrollerc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Iscrollerc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iscrollerc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Iscrollerc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iscrollerc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Iscrollerc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iscrollerc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Iscrollerc20Session struct {
	Contract     *Iscrollerc20     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Iscrollerc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Iscrollerc20CallerSession struct {
	Contract *Iscrollerc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Iscrollerc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Iscrollerc20TransactorSession struct {
	Contract     *Iscrollerc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Iscrollerc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Iscrollerc20Raw struct {
	Contract *Iscrollerc20 // Generic contract binding to access the raw methods on
}

// Iscrollerc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Iscrollerc20CallerRaw struct {
	Contract *Iscrollerc20Caller // Generic read-only contract binding to access the raw methods on
}

// Iscrollerc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Iscrollerc20TransactorRaw struct {
	Contract *Iscrollerc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIscrollerc20 creates a new instance of Iscrollerc20, bound to a specific deployed contract.
func NewIscrollerc20(address common.Address, backend bind.ContractBackend) (*Iscrollerc20, error) {
	contract, err := bindIscrollerc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iscrollerc20{Iscrollerc20Caller: Iscrollerc20Caller{contract: contract}, Iscrollerc20Transactor: Iscrollerc20Transactor{contract: contract}, Iscrollerc20Filterer: Iscrollerc20Filterer{contract: contract}}, nil
}

// NewIscrollerc20Caller creates a new read-only instance of Iscrollerc20, bound to a specific deployed contract.
func NewIscrollerc20Caller(address common.Address, caller bind.ContractCaller) (*Iscrollerc20Caller, error) {
	contract, err := bindIscrollerc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Iscrollerc20Caller{contract: contract}, nil
}

// NewIscrollerc20Transactor creates a new write-only instance of Iscrollerc20, bound to a specific deployed contract.
func NewIscrollerc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Iscrollerc20Transactor, error) {
	contract, err := bindIscrollerc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Iscrollerc20Transactor{contract: contract}, nil
}

// NewIscrollerc20Filterer creates a new log filterer instance of Iscrollerc20, bound to a specific deployed contract.
func NewIscrollerc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Iscrollerc20Filterer, error) {
	contract, err := bindIscrollerc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Iscrollerc20Filterer{contract: contract}, nil
}

// bindIscrollerc20 binds a generic wrapper to an already deployed contract.
func bindIscrollerc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Iscrollerc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iscrollerc20 *Iscrollerc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iscrollerc20.Contract.Iscrollerc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iscrollerc20 *Iscrollerc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iscrollerc20.Contract.Iscrollerc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iscrollerc20 *Iscrollerc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iscrollerc20.Contract.Iscrollerc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iscrollerc20 *Iscrollerc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iscrollerc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iscrollerc20 *Iscrollerc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iscrollerc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iscrollerc20 *Iscrollerc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iscrollerc20.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Iscrollerc20 *Iscrollerc20Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Iscrollerc20.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Iscrollerc20 *Iscrollerc20Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _Iscrollerc20.Contract.DOMAINSEPARATOR(&_Iscrollerc20.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Iscrollerc20 *Iscrollerc20CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Iscrollerc20.Contract.DOMAINSEPARATOR(&_Iscrollerc20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Iscrollerc20 *Iscrollerc20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Iscrollerc20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Iscrollerc20 *Iscrollerc20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Iscrollerc20.Contract.Allowance(&_Iscrollerc20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Iscrollerc20 *Iscrollerc20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Iscrollerc20.Contract.Allowance(&_Iscrollerc20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Iscrollerc20 *Iscrollerc20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Iscrollerc20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Iscrollerc20 *Iscrollerc20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _Iscrollerc20.Contract.BalanceOf(&_Iscrollerc20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Iscrollerc20 *Iscrollerc20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Iscrollerc20.Contract.BalanceOf(&_Iscrollerc20.CallOpts, account)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_Iscrollerc20 *Iscrollerc20Caller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iscrollerc20.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_Iscrollerc20 *Iscrollerc20Session) Counterpart() (common.Address, error) {
	return _Iscrollerc20.Contract.Counterpart(&_Iscrollerc20.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_Iscrollerc20 *Iscrollerc20CallerSession) Counterpart() (common.Address, error) {
	return _Iscrollerc20.Contract.Counterpart(&_Iscrollerc20.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Iscrollerc20 *Iscrollerc20Caller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iscrollerc20.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Iscrollerc20 *Iscrollerc20Session) Gateway() (common.Address, error) {
	return _Iscrollerc20.Contract.Gateway(&_Iscrollerc20.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Iscrollerc20 *Iscrollerc20CallerSession) Gateway() (common.Address, error) {
	return _Iscrollerc20.Contract.Gateway(&_Iscrollerc20.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Iscrollerc20 *Iscrollerc20Caller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Iscrollerc20.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Iscrollerc20 *Iscrollerc20Session) Nonces(owner common.Address) (*big.Int, error) {
	return _Iscrollerc20.Contract.Nonces(&_Iscrollerc20.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Iscrollerc20 *Iscrollerc20CallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Iscrollerc20.Contract.Nonces(&_Iscrollerc20.CallOpts, owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Iscrollerc20 *Iscrollerc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iscrollerc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Iscrollerc20 *Iscrollerc20Session) TotalSupply() (*big.Int, error) {
	return _Iscrollerc20.Contract.TotalSupply(&_Iscrollerc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Iscrollerc20 *Iscrollerc20CallerSession) TotalSupply() (*big.Int, error) {
	return _Iscrollerc20.Contract.TotalSupply(&_Iscrollerc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Iscrollerc20 *Iscrollerc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Iscrollerc20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Iscrollerc20 *Iscrollerc20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Iscrollerc20.Contract.Approve(&_Iscrollerc20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Iscrollerc20 *Iscrollerc20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Iscrollerc20.Contract.Approve(&_Iscrollerc20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_Iscrollerc20 *Iscrollerc20Transactor) Burn(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Iscrollerc20.contract.Transact(opts, "burn", _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_Iscrollerc20 *Iscrollerc20Session) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Iscrollerc20.Contract.Burn(&_Iscrollerc20.TransactOpts, _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_Iscrollerc20 *Iscrollerc20TransactorSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Iscrollerc20.Contract.Burn(&_Iscrollerc20.TransactOpts, _from, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Iscrollerc20 *Iscrollerc20Transactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Iscrollerc20.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Iscrollerc20 *Iscrollerc20Session) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Iscrollerc20.Contract.Mint(&_Iscrollerc20.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Iscrollerc20 *Iscrollerc20TransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Iscrollerc20.Contract.Mint(&_Iscrollerc20.TransactOpts, _to, _amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Iscrollerc20 *Iscrollerc20Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Iscrollerc20.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Iscrollerc20 *Iscrollerc20Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Iscrollerc20.Contract.Permit(&_Iscrollerc20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Iscrollerc20 *Iscrollerc20TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Iscrollerc20.Contract.Permit(&_Iscrollerc20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Iscrollerc20 *Iscrollerc20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Iscrollerc20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Iscrollerc20 *Iscrollerc20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Iscrollerc20.Contract.Transfer(&_Iscrollerc20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Iscrollerc20 *Iscrollerc20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Iscrollerc20.Contract.Transfer(&_Iscrollerc20.TransactOpts, to, amount)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address receiver, uint256 amount, bytes data) returns(bool success)
func (_Iscrollerc20 *Iscrollerc20Transactor) TransferAndCall(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Iscrollerc20.contract.Transact(opts, "transferAndCall", receiver, amount, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address receiver, uint256 amount, bytes data) returns(bool success)
func (_Iscrollerc20 *Iscrollerc20Session) TransferAndCall(receiver common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Iscrollerc20.Contract.TransferAndCall(&_Iscrollerc20.TransactOpts, receiver, amount, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address receiver, uint256 amount, bytes data) returns(bool success)
func (_Iscrollerc20 *Iscrollerc20TransactorSession) TransferAndCall(receiver common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Iscrollerc20.Contract.TransferAndCall(&_Iscrollerc20.TransactOpts, receiver, amount, data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Iscrollerc20 *Iscrollerc20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Iscrollerc20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Iscrollerc20 *Iscrollerc20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Iscrollerc20.Contract.TransferFrom(&_Iscrollerc20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Iscrollerc20 *Iscrollerc20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Iscrollerc20.Contract.TransferFrom(&_Iscrollerc20.TransactOpts, from, to, amount)
}

// Iscrollerc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Iscrollerc20 contract.
type Iscrollerc20ApprovalIterator struct {
	Event *Iscrollerc20Approval // Event containing the contract specifics and raw log

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
func (it *Iscrollerc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Iscrollerc20Approval)
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
		it.Event = new(Iscrollerc20Approval)
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
func (it *Iscrollerc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Iscrollerc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Iscrollerc20Approval represents a Approval event raised by the Iscrollerc20 contract.
type Iscrollerc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Iscrollerc20 *Iscrollerc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Iscrollerc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Iscrollerc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Iscrollerc20ApprovalIterator{contract: _Iscrollerc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Iscrollerc20 *Iscrollerc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Iscrollerc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Iscrollerc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Iscrollerc20Approval)
				if err := _Iscrollerc20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Iscrollerc20 *Iscrollerc20Filterer) ParseApproval(log types.Log) (*Iscrollerc20Approval, error) {
	event := new(Iscrollerc20Approval)
	if err := _Iscrollerc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Iscrollerc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Iscrollerc20 contract.
type Iscrollerc20TransferIterator struct {
	Event *Iscrollerc20Transfer // Event containing the contract specifics and raw log

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
func (it *Iscrollerc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Iscrollerc20Transfer)
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
		it.Event = new(Iscrollerc20Transfer)
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
func (it *Iscrollerc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Iscrollerc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Iscrollerc20Transfer represents a Transfer event raised by the Iscrollerc20 contract.
type Iscrollerc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Iscrollerc20 *Iscrollerc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Iscrollerc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Iscrollerc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Iscrollerc20TransferIterator{contract: _Iscrollerc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Iscrollerc20 *Iscrollerc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Iscrollerc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Iscrollerc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Iscrollerc20Transfer)
				if err := _Iscrollerc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Iscrollerc20 *Iscrollerc20Filterer) ParseTransfer(log types.Log) (*Iscrollerc20Transfer, error) {
	event := new(Iscrollerc20Transfer)
	if err := _Iscrollerc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
