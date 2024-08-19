// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package istandarderc20

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

// Istandarderc20MetaData contains all meta data concerning the Istandarderc20 contract.
var Istandarderc20MetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]",
}

// Istandarderc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Istandarderc20MetaData.ABI instead.
var Istandarderc20ABI = Istandarderc20MetaData.ABI

// Istandarderc20 is an auto generated Go binding around an Ethereum contract.
type Istandarderc20 struct {
	Istandarderc20Caller     // Read-only binding to the contract
	Istandarderc20Transactor // Write-only binding to the contract
	Istandarderc20Filterer   // Log filterer for contract events
}

// Istandarderc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Istandarderc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Istandarderc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Istandarderc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Istandarderc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Istandarderc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Istandarderc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Istandarderc20Session struct {
	Contract     *Istandarderc20   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Istandarderc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Istandarderc20CallerSession struct {
	Contract *Istandarderc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Istandarderc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Istandarderc20TransactorSession struct {
	Contract     *Istandarderc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Istandarderc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Istandarderc20Raw struct {
	Contract *Istandarderc20 // Generic contract binding to access the raw methods on
}

// Istandarderc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Istandarderc20CallerRaw struct {
	Contract *Istandarderc20Caller // Generic read-only contract binding to access the raw methods on
}

// Istandarderc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Istandarderc20TransactorRaw struct {
	Contract *Istandarderc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIstandarderc20 creates a new instance of Istandarderc20, bound to a specific deployed contract.
func NewIstandarderc20(address common.Address, backend bind.ContractBackend) (*Istandarderc20, error) {
	contract, err := bindIstandarderc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Istandarderc20{Istandarderc20Caller: Istandarderc20Caller{contract: contract}, Istandarderc20Transactor: Istandarderc20Transactor{contract: contract}, Istandarderc20Filterer: Istandarderc20Filterer{contract: contract}}, nil
}

// NewIstandarderc20Caller creates a new read-only instance of Istandarderc20, bound to a specific deployed contract.
func NewIstandarderc20Caller(address common.Address, caller bind.ContractCaller) (*Istandarderc20Caller, error) {
	contract, err := bindIstandarderc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Istandarderc20Caller{contract: contract}, nil
}

// NewIstandarderc20Transactor creates a new write-only instance of Istandarderc20, bound to a specific deployed contract.
func NewIstandarderc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Istandarderc20Transactor, error) {
	contract, err := bindIstandarderc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Istandarderc20Transactor{contract: contract}, nil
}

// NewIstandarderc20Filterer creates a new log filterer instance of Istandarderc20, bound to a specific deployed contract.
func NewIstandarderc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Istandarderc20Filterer, error) {
	contract, err := bindIstandarderc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Istandarderc20Filterer{contract: contract}, nil
}

// bindIstandarderc20 binds a generic wrapper to an already deployed contract.
func bindIstandarderc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Istandarderc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Istandarderc20 *Istandarderc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Istandarderc20.Contract.Istandarderc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Istandarderc20 *Istandarderc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Istandarderc20.Contract.Istandarderc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Istandarderc20 *Istandarderc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Istandarderc20.Contract.Istandarderc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Istandarderc20 *Istandarderc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Istandarderc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Istandarderc20 *Istandarderc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Istandarderc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Istandarderc20 *Istandarderc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Istandarderc20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Istandarderc20 *Istandarderc20Caller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Istandarderc20.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Istandarderc20 *Istandarderc20Session) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Istandarderc20.Contract.Allowance(&_Istandarderc20.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Istandarderc20 *Istandarderc20CallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Istandarderc20.Contract.Allowance(&_Istandarderc20.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Istandarderc20 *Istandarderc20Caller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Istandarderc20.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Istandarderc20 *Istandarderc20Session) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Istandarderc20.Contract.BalanceOf(&_Istandarderc20.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Istandarderc20 *Istandarderc20CallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Istandarderc20.Contract.BalanceOf(&_Istandarderc20.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Istandarderc20 *Istandarderc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Istandarderc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Istandarderc20 *Istandarderc20Session) Decimals() (uint8, error) {
	return _Istandarderc20.Contract.Decimals(&_Istandarderc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Istandarderc20 *Istandarderc20CallerSession) Decimals() (uint8, error) {
	return _Istandarderc20.Contract.Decimals(&_Istandarderc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Istandarderc20 *Istandarderc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Istandarderc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Istandarderc20 *Istandarderc20Session) Name() (string, error) {
	return _Istandarderc20.Contract.Name(&_Istandarderc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Istandarderc20 *Istandarderc20CallerSession) Name() (string, error) {
	return _Istandarderc20.Contract.Name(&_Istandarderc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Istandarderc20 *Istandarderc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Istandarderc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Istandarderc20 *Istandarderc20Session) Symbol() (string, error) {
	return _Istandarderc20.Contract.Symbol(&_Istandarderc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Istandarderc20 *Istandarderc20CallerSession) Symbol() (string, error) {
	return _Istandarderc20.Contract.Symbol(&_Istandarderc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Istandarderc20 *Istandarderc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Istandarderc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Istandarderc20 *Istandarderc20Session) TotalSupply() (*big.Int, error) {
	return _Istandarderc20.Contract.TotalSupply(&_Istandarderc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Istandarderc20 *Istandarderc20CallerSession) TotalSupply() (*big.Int, error) {
	return _Istandarderc20.Contract.TotalSupply(&_Istandarderc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Istandarderc20 *Istandarderc20Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Istandarderc20.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Istandarderc20 *Istandarderc20Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Istandarderc20.Contract.Approve(&_Istandarderc20.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Istandarderc20 *Istandarderc20TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Istandarderc20.Contract.Approve(&_Istandarderc20.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Istandarderc20 *Istandarderc20Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Istandarderc20.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Istandarderc20 *Istandarderc20Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Istandarderc20.Contract.Transfer(&_Istandarderc20.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Istandarderc20 *Istandarderc20TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Istandarderc20.Contract.Transfer(&_Istandarderc20.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Istandarderc20 *Istandarderc20Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Istandarderc20.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Istandarderc20 *Istandarderc20Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Istandarderc20.Contract.TransferFrom(&_Istandarderc20.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Istandarderc20 *Istandarderc20TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Istandarderc20.Contract.TransferFrom(&_Istandarderc20.TransactOpts, _from, _to, _value)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Istandarderc20 *Istandarderc20Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Istandarderc20.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Istandarderc20 *Istandarderc20Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Istandarderc20.Contract.Fallback(&_Istandarderc20.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Istandarderc20 *Istandarderc20TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Istandarderc20.Contract.Fallback(&_Istandarderc20.TransactOpts, calldata)
}

// Istandarderc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Istandarderc20 contract.
type Istandarderc20ApprovalIterator struct {
	Event *Istandarderc20Approval // Event containing the contract specifics and raw log

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
func (it *Istandarderc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Istandarderc20Approval)
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
		it.Event = new(Istandarderc20Approval)
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
func (it *Istandarderc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Istandarderc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Istandarderc20Approval represents a Approval event raised by the Istandarderc20 contract.
type Istandarderc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Istandarderc20 *Istandarderc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Istandarderc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Istandarderc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Istandarderc20ApprovalIterator{contract: _Istandarderc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Istandarderc20 *Istandarderc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Istandarderc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Istandarderc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Istandarderc20Approval)
				if err := _Istandarderc20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Istandarderc20 *Istandarderc20Filterer) ParseApproval(log types.Log) (*Istandarderc20Approval, error) {
	event := new(Istandarderc20Approval)
	if err := _Istandarderc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Istandarderc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Istandarderc20 contract.
type Istandarderc20TransferIterator struct {
	Event *Istandarderc20Transfer // Event containing the contract specifics and raw log

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
func (it *Istandarderc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Istandarderc20Transfer)
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
		it.Event = new(Istandarderc20Transfer)
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
func (it *Istandarderc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Istandarderc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Istandarderc20Transfer represents a Transfer event raised by the Istandarderc20 contract.
type Istandarderc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Istandarderc20 *Istandarderc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Istandarderc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Istandarderc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Istandarderc20TransferIterator{contract: _Istandarderc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Istandarderc20 *Istandarderc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Istandarderc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Istandarderc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Istandarderc20Transfer)
				if err := _Istandarderc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Istandarderc20 *Istandarderc20Filterer) ParseTransfer(log types.Log) (*Istandarderc20Transfer, error) {
	event := new(Istandarderc20Transfer)
	if err := _Istandarderc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
