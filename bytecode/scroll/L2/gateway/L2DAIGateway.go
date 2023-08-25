// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gateway

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
)

// L2DAIGatewayMetaData contains all meta data concerning the L2DAIGateway contract.
var (
	L2DAIGatewayMetaData = &bind.MetaData{
		ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false}],\"type\":\"event\",\"name\":\"FinalizeDepositERC20\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Initialized\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"oldL1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"newL1Token\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"UpdateTokenMapping\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false}],\"type\":\"event\",\"name\":\"WithdrawERC20\",\"anonymous\":false},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"finalizeDepositERC20\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"getL1ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"renounceOwnership\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"tokenMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferOwnership\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateTokenMapping\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"withdrawERC20\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"withdrawERC20\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"withdrawERC20AndCall\"}]",
	}
)

// L2DAIGateway is an auto generated Go binding around an Ethereum contract.
type L2DAIGateway struct {
	Name    string
	Address common.Address // contract address
	ABI     *abi.ABI       // L2DAIGatewayABI is the input ABI used to generate the binding from.

	topics  map[common.Hash]string
	parsers map[common.Hash]func(log *types.Log) error

	L2DAIGatewayCaller     // Read-only binding to the contract
	L2DAIGatewayTransactor // Write-only binding to the contract
}

// GetName return L2DAIGateway's Name.
func (o *L2DAIGateway) GetName() string {
	return o.Name
}

// GetAddress return L2DAIGateway's contract address.
func (o *L2DAIGateway) GetAddress() common.Address {
	return o.Address
}

// GetEventName get event name by event hash.
func (o *L2DAIGateway) GetEventName(sigHash common.Hash) string {
	return o.topics[sigHash]
}

// GetSigHashes get sig hash list.
func (o *L2DAIGateway) GetSigHashes() []common.Hash {
	if len(o.parsers) == 0 {
		return nil
	}
	var sigHashes = make([]common.Hash, 0, len(o.parsers))
	for id := range o.parsers {
		sigHashes = append(sigHashes, id)
	}
	return sigHashes
}

// ParseLog parse the log if parse func is exist.
func (o *L2DAIGateway) ParseLog(vLog *types.Log) error {
	_id := vLog.Topics[0]
	if parse, exist := o.parsers[_id]; exist {
		return parse(vLog)
	}
	return nil
}

// RegisterFinalizeDepositERC20, the FinalizeDepositERC20 event ID is 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
func (o *L2DAIGateway) RegisterFinalizeDepositERC20(handler func(vLog *types.Log, data *L2DAIGatewayFinalizeDepositERC20Event) error) {
	o.parsers[o.ABI.Events["FinalizeDepositERC20"].ID] = func(log *types.Log) error {
		event := new(L2DAIGatewayFinalizeDepositERC20Event)
		if err := o.L2DAIGatewayCaller.contract.UnpackLog(event, "FinalizeDepositERC20", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// RegisterInitialized, the Initialized event ID is 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
func (o *L2DAIGateway) RegisterInitialized(handler func(vLog *types.Log, data *L2DAIGatewayInitializedEvent) error) {
	o.parsers[o.ABI.Events["Initialized"].ID] = func(log *types.Log) error {
		event := new(L2DAIGatewayInitializedEvent)
		if err := o.L2DAIGatewayCaller.contract.UnpackLog(event, "Initialized", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// RegisterOwnershipTransferred, the OwnershipTransferred event ID is 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
func (o *L2DAIGateway) RegisterOwnershipTransferred(handler func(vLog *types.Log, data *L2DAIGatewayOwnershipTransferredEvent) error) {
	o.parsers[o.ABI.Events["OwnershipTransferred"].ID] = func(log *types.Log) error {
		event := new(L2DAIGatewayOwnershipTransferredEvent)
		if err := o.L2DAIGatewayCaller.contract.UnpackLog(event, "OwnershipTransferred", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// RegisterUpdateTokenMapping, the UpdateTokenMapping event ID is 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
func (o *L2DAIGateway) RegisterUpdateTokenMapping(handler func(vLog *types.Log, data *L2DAIGatewayUpdateTokenMappingEvent) error) {
	o.parsers[o.ABI.Events["UpdateTokenMapping"].ID] = func(log *types.Log) error {
		event := new(L2DAIGatewayUpdateTokenMappingEvent)
		if err := o.L2DAIGatewayCaller.contract.UnpackLog(event, "UpdateTokenMapping", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// RegisterWithdrawERC20, the WithdrawERC20 event ID is 0xd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e8.
func (o *L2DAIGateway) RegisterWithdrawERC20(handler func(vLog *types.Log, data *L2DAIGatewayWithdrawERC20Event) error) {
	o.parsers[o.ABI.Events["WithdrawERC20"].ID] = func(log *types.Log) error {
		event := new(L2DAIGatewayWithdrawERC20Event)
		if err := o.L2DAIGatewayCaller.contract.UnpackLog(event, "WithdrawERC20", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// L2DAIGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2DAIGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2DAIGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2DAIGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewL2DAIGateway creates a new instance of L2DAIGateway, bound to a specific deployed contract.
func NewL2DAIGateway(address common.Address, backend bind.ContractBackend) (*L2DAIGateway, error) {
	contract, err := bindL2DAIGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	sigAbi, err := L2DAIGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	topics := make(map[common.Hash]string, len(sigAbi.Events))
	for name, _abi := range sigAbi.Events {
		topics[_abi.ID] = name
	}

	return &L2DAIGateway{
		Name:                   "L2DAIGateway",
		ABI:                    sigAbi,
		Address:                address,
		topics:                 topics,
		parsers:                map[common.Hash]func(log *types.Log) error{},
		L2DAIGatewayCaller:     L2DAIGatewayCaller{contract: contract},
		L2DAIGatewayTransactor: L2DAIGatewayTransactor{contract: contract},
	}, nil
}

// NewL2DAIGatewayCaller creates a new read-only instance of L2DAIGateway, bound to a specific deployed contract.
func NewL2DAIGatewayCaller(address common.Address, caller bind.ContractCaller) (*L2DAIGatewayCaller, error) {
	contract, err := bindL2DAIGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2DAIGatewayCaller{contract: contract}, nil
}

// NewL2DAIGatewayTransactor creates a new write-only instance of L2DAIGateway, bound to a specific deployed contract.
func NewL2DAIGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L2DAIGatewayTransactor, error) {
	contract, err := bindL2DAIGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2DAIGatewayTransactor{contract: contract}, nil
}

// bindL2DAIGateway binds a generic wrapper to an already deployed contract.
func bindL2DAIGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2DAIGatewayMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2DAIGateway *L2DAIGatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2DAIGateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Token) view returns(address)
func (_L2DAIGateway *L2DAIGatewayCaller) GetL1ERC20Address(opts *bind.CallOpts, _l2Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2DAIGateway.contract.Call(opts, &out, "getL1ERC20Address", _l2Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) pure returns(address)
func (_L2DAIGateway *L2DAIGatewayCaller) GetL2ERC20Address(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2DAIGateway.contract.Call(opts, &out, "getL2ERC20Address", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2DAIGateway *L2DAIGatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2DAIGateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2DAIGateway *L2DAIGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2DAIGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2DAIGateway *L2DAIGatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2DAIGateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2DAIGateway *L2DAIGatewayCaller) TokenMapping(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2DAIGateway.contract.Call(opts, &out, "tokenMapping", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2DAIGateway *L2DAIGatewayTransactor) FinalizeDepositERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2DAIGateway.contract.Transact(opts, "finalizeDepositERC20", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L2DAIGateway *L2DAIGatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2DAIGateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2DAIGateway *L2DAIGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2DAIGateway.contract.Transact(opts, "renounceOwnership")
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2DAIGateway *L2DAIGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2DAIGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2DAIGateway *L2DAIGatewayTransactor) UpdateTokenMapping(opts *bind.TransactOpts, _l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2DAIGateway.contract.Transact(opts, "updateTokenMapping", _l2Token, _l1Token)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2DAIGateway *L2DAIGatewayTransactor) WithdrawERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2DAIGateway.contract.Transact(opts, "withdrawERC20", _token, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2DAIGateway *L2DAIGatewayTransactor) WithdrawERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2DAIGateway.contract.Transact(opts, "withdrawERC200", _token, _to, _amount, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2DAIGateway *L2DAIGatewayTransactor) WithdrawERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2DAIGateway.contract.Transact(opts, "withdrawERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// L2DAIGatewayFinalizeDepositERC20 represents a FinalizeDepositERC20 event raised by the L2DAIGateway contract.
type L2DAIGatewayFinalizeDepositERC20Event struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
}

// L2DAIGatewayInitialized represents a Initialized event raised by the L2DAIGateway contract.
type L2DAIGatewayInitializedEvent struct {
	Version uint8
}

// L2DAIGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L2DAIGateway contract.
type L2DAIGatewayOwnershipTransferredEvent struct {
	PreviousOwner common.Address
	NewOwner      common.Address
}

// L2DAIGatewayUpdateTokenMapping represents a UpdateTokenMapping event raised by the L2DAIGateway contract.
type L2DAIGatewayUpdateTokenMappingEvent struct {
	L2Token    common.Address
	OldL1Token common.Address
	NewL1Token common.Address
}

// L2DAIGatewayWithdrawERC20 represents a WithdrawERC20 event raised by the L2DAIGateway contract.
type L2DAIGatewayWithdrawERC20Event struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
}
