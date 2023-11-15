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

// L1WETHGatewayMetaData contains all meta data concerning the L1WETHGateway contract.
var (
	L1WETHGatewayMetaData = &bind.MetaData{
		ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2WETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false}],\"type\":\"event\",\"name\":\"DepositERC20\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false}],\"type\":\"event\",\"name\":\"FinalizeWithdrawERC20\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Initialized\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"RefundERC20\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldRateLimiter\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_newRateLimiter\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"UpdateRateLimiter\",\"anonymous\":false},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"depositERC20\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"depositERC20\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"depositERC20AndCall\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"finalizeWithdrawERC20\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"l2WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"onDropMessage\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"rateLimiter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"renounceOwnership\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferOwnership\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRateLimiter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateRateLimiter\"},{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	}
	// L1WETHGatewayABI is the input ABI used to generate the binding from.
	L1WETHGatewayABI, _ = L1WETHGatewayMetaData.GetAbi()
)

// L1WETHGateway is an auto generated Go binding around an Ethereum contract.
type L1WETHGateway struct {
	Name    string
	Address common.Address // contract address
	ABI     *abi.ABI       // L1WETHGatewayABI is the input ABI used to generate the binding from.

	topics  map[common.Hash]string
	parsers map[common.Hash]func(log *types.Log) error

	L1WETHGatewayCaller     // Read-only binding to the contract
	L1WETHGatewayTransactor // Write-only binding to the contract
}

// GetName return L1WETHGateway's Name.
func (o *L1WETHGateway) GetName() string {
	return o.Name
}

// GetAddress return L1WETHGateway's contract address.
func (o *L1WETHGateway) GetAddress() common.Address {
	return o.Address
}

// GetEventName get event name by event hash.
func (o *L1WETHGateway) GetEventName(sigHash common.Hash) string {
	return o.topics[sigHash]
}

// GetSigHashes get sig hash list.
func (o *L1WETHGateway) GetSigHashes() []common.Hash {
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
func (o *L1WETHGateway) ParseLog(vLog *types.Log) (bool, error) {
	_id := vLog.Topics[0]
	if parse, exist := o.parsers[_id]; exist {
		return true, parse(vLog)
	}
	return false, nil
}

// RegisterDepositERC20, the DepositERC20 event ID is 0x31cd3b976e4d654022bf95c68a2ce53f1d5d94afabe0454d2832208eeb40af25.
func (o *L1WETHGateway) RegisterDepositERC20(handler func(vLog *types.Log, data *L1WETHGatewayDepositERC20Event) error) {
	_id := o.ABI.Events["DepositERC20"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1WETHGatewayDepositERC20Event)
		if err := o.L1WETHGatewayCaller.contract.UnpackLog(event, "DepositERC20", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "DepositERC20"
}

// RegisterFinalizeWithdrawERC20, the FinalizeWithdrawERC20 event ID is 0xc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a7.
func (o *L1WETHGateway) RegisterFinalizeWithdrawERC20(handler func(vLog *types.Log, data *L1WETHGatewayFinalizeWithdrawERC20Event) error) {
	_id := o.ABI.Events["FinalizeWithdrawERC20"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1WETHGatewayFinalizeWithdrawERC20Event)
		if err := o.L1WETHGatewayCaller.contract.UnpackLog(event, "FinalizeWithdrawERC20", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "FinalizeWithdrawERC20"
}

// RegisterInitialized, the Initialized event ID is 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
func (o *L1WETHGateway) RegisterInitialized(handler func(vLog *types.Log, data *L1WETHGatewayInitializedEvent) error) {
	_id := o.ABI.Events["Initialized"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1WETHGatewayInitializedEvent)
		if err := o.L1WETHGatewayCaller.contract.UnpackLog(event, "Initialized", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "Initialized"
}

// RegisterOwnershipTransferred, the OwnershipTransferred event ID is 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
func (o *L1WETHGateway) RegisterOwnershipTransferred(handler func(vLog *types.Log, data *L1WETHGatewayOwnershipTransferredEvent) error) {
	_id := o.ABI.Events["OwnershipTransferred"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1WETHGatewayOwnershipTransferredEvent)
		if err := o.L1WETHGatewayCaller.contract.UnpackLog(event, "OwnershipTransferred", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "OwnershipTransferred"
}

// RegisterRefundERC20, the RefundERC20 event ID is 0xdbdf8eb487847e4c0f22847f5dac07f2d3690f96f581a6ae4b102769917645a8.
func (o *L1WETHGateway) RegisterRefundERC20(handler func(vLog *types.Log, data *L1WETHGatewayRefundERC20Event) error) {
	_id := o.ABI.Events["RefundERC20"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1WETHGatewayRefundERC20Event)
		if err := o.L1WETHGatewayCaller.contract.UnpackLog(event, "RefundERC20", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "RefundERC20"
}

// RegisterUpdateRateLimiter, the UpdateRateLimiter event ID is 0x53d49a5b37a3ed13692a8b217b8c79f53b0144c069de23e0cc5704d89bc23006.
func (o *L1WETHGateway) RegisterUpdateRateLimiter(handler func(vLog *types.Log, data *L1WETHGatewayUpdateRateLimiterEvent) error) {
	_id := o.ABI.Events["UpdateRateLimiter"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1WETHGatewayUpdateRateLimiterEvent)
		if err := o.L1WETHGatewayCaller.contract.UnpackLog(event, "UpdateRateLimiter", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "UpdateRateLimiter"
}

// L1WETHGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1WETHGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1WETHGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1WETHGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewL1WETHGateway creates a new instance of L1WETHGateway, bound to a specific deployed contract.
func NewL1WETHGateway(address common.Address, backend bind.ContractBackend) (*L1WETHGateway, error) {
	contract, err := bindL1WETHGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	sigAbi, err := L1WETHGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	topics := make(map[common.Hash]string, len(sigAbi.Events))
	for name, _abi := range sigAbi.Events {
		topics[_abi.ID] = name
	}

	return &L1WETHGateway{
		Name:                    "L1WETHGateway",
		ABI:                     sigAbi,
		Address:                 address,
		topics:                  topics,
		parsers:                 map[common.Hash]func(log *types.Log) error{},
		L1WETHGatewayCaller:     L1WETHGatewayCaller{contract: contract},
		L1WETHGatewayTransactor: L1WETHGatewayTransactor{contract: contract},
	}, nil
}

// NewL1WETHGatewayCaller creates a new read-only instance of L1WETHGateway, bound to a specific deployed contract.
func NewL1WETHGatewayCaller(address common.Address, caller bind.ContractCaller) (*L1WETHGatewayCaller, error) {
	contract, err := bindL1WETHGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1WETHGatewayCaller{contract: contract}, nil
}

// NewL1WETHGatewayTransactor creates a new write-only instance of L1WETHGateway, bound to a specific deployed contract.
func NewL1WETHGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L1WETHGatewayTransactor, error) {
	contract, err := bindL1WETHGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1WETHGatewayTransactor{contract: contract}, nil
}

// bindL1WETHGateway binds a generic wrapper to an already deployed contract.
func bindL1WETHGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1WETHGatewayMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_L1WETHGateway *L1WETHGatewayCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1WETHGateway.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1WETHGateway *L1WETHGatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1WETHGateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) view returns(address)
func (_L1WETHGateway *L1WETHGatewayCaller) GetL2ERC20Address(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1WETHGateway.contract.Call(opts, &out, "getL2ERC20Address", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2WETH is a free data retrieval call binding the contract method 0x88558687.
//
// Solidity: function l2WETH() view returns(address)
func (_L1WETHGateway *L1WETHGatewayCaller) L2WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1WETHGateway.contract.Call(opts, &out, "l2WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1WETHGateway *L1WETHGatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1WETHGateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1WETHGateway *L1WETHGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1WETHGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RateLimiter is a free data retrieval call binding the contract method 0x53d4fe33.
//
// Solidity: function rateLimiter() view returns(address)
func (_L1WETHGateway *L1WETHGatewayCaller) RateLimiter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1WETHGateway.contract.Call(opts, &out, "rateLimiter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1WETHGateway *L1WETHGatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1WETHGateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositERC20 is a paid mutator transaction binding the contract method 0x21425ee0.
//
// Solidity: function depositERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1WETHGateway *L1WETHGatewayTransactor) DepositERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1WETHGateway.contract.Transact(opts, "depositERC20", _token, _amount, _gasLimit)
}

// DepositERC200 is a paid mutator transaction binding the contract method 0xf219fa66.
//
// Solidity: function depositERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1WETHGateway *L1WETHGatewayTransactor) DepositERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1WETHGateway.contract.Transact(opts, "depositERC200", _token, _to, _amount, _gasLimit)
}

// DepositERC20AndCall is a paid mutator transaction binding the contract method 0x0aea8c26.
//
// Solidity: function depositERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L1WETHGateway *L1WETHGatewayTransactor) DepositERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1WETHGateway.contract.Transact(opts, "depositERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// FinalizeWithdrawERC20 is a paid mutator transaction binding the contract method 0x84bd13b0.
//
// Solidity: function finalizeWithdrawERC20(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1WETHGateway *L1WETHGatewayTransactor) FinalizeWithdrawERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1WETHGateway.contract.Transact(opts, "finalizeWithdrawERC20", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _counterpart, address _router, address _messenger) returns()
func (_L1WETHGateway *L1WETHGatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _router common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1WETHGateway.contract.Transact(opts, "initialize", _counterpart, _router, _messenger)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1WETHGateway *L1WETHGatewayTransactor) OnDropMessage(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _L1WETHGateway.contract.Transact(opts, "onDropMessage", _message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1WETHGateway *L1WETHGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1WETHGateway.contract.Transact(opts, "renounceOwnership")
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1WETHGateway *L1WETHGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1WETHGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// UpdateRateLimiter is a paid mutator transaction binding the contract method 0xe157820a.
//
// Solidity: function updateRateLimiter(address _newRateLimiter) returns()
func (_L1WETHGateway *L1WETHGatewayTransactor) UpdateRateLimiter(opts *bind.TransactOpts, _newRateLimiter common.Address) (*types.Transaction, error) {
	return _L1WETHGateway.contract.Transact(opts, "updateRateLimiter", _newRateLimiter)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1WETHGateway *L1WETHGatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1WETHGateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// L1WETHGatewayDepositERC20 represents a DepositERC20 event raised by the L1WETHGateway contract.
type L1WETHGatewayDepositERC20Event struct {
	Log *types.Log `json:"-" gorm:"-"`

	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
}

// L1WETHGatewayFinalizeWithdrawERC20 represents a FinalizeWithdrawERC20 event raised by the L1WETHGateway contract.
type L1WETHGatewayFinalizeWithdrawERC20Event struct {
	Log *types.Log `json:"-" gorm:"-"`

	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
}

// L1WETHGatewayInitialized represents a Initialized event raised by the L1WETHGateway contract.
type L1WETHGatewayInitializedEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	Version uint8
}

// L1WETHGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L1WETHGateway contract.
type L1WETHGatewayOwnershipTransferredEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	PreviousOwner common.Address
	NewOwner      common.Address
}

// L1WETHGatewayRefundERC20 represents a RefundERC20 event raised by the L1WETHGateway contract.
type L1WETHGatewayRefundERC20Event struct {
	Log *types.Log `json:"-" gorm:"-"`

	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
}

// L1WETHGatewayUpdateRateLimiter represents a UpdateRateLimiter event raised by the L1WETHGateway contract.
type L1WETHGatewayUpdateRateLimiterEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	OldRateLimiter common.Address
	NewRateLimiter common.Address
}