// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package L2

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

// L2ScrollMessengerMetaData contains all meta data concerning the L2ScrollMessenger contract.
var (
	L2ScrollMessengerMetaData = &bind.MetaData{
		ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageQueue\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\",\"indexed\":true}],\"type\":\"event\",\"name\":\"FailedRelayedMessage\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Initialized\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Paused\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\",\"indexed\":true}],\"type\":\"event\",\"name\":\"RelayedMessage\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false}],\"type\":\"event\",\"name\":\"SentMessage\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Unpaused\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldFeeVault\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"address\",\"name\":\"_newFeeVault\",\"type\":\"address\",\"indexed\":false}],\"type\":\"event\",\"name\":\"UpdateFeeVault\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"oldMaxFailedExecutionTimes\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"newMaxFailedExecutionTimes\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"UpdateMaxFailedExecutionTimes\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldRateLimiter\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_newRateLimiter\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"UpdateRateLimiter\",\"anonymous\":false},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"feeVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"isL1MessageExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}]},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"isL2MessageSent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"messageQueue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"rateLimiter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"relayMessage\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"renounceOwnership\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"sendMessage\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"sendMessage\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setPause\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferOwnership\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFeeVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateFeeVault\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRateLimiter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateRateLimiter\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"xDomainMessageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	}
	// L2ScrollMessengerABI is the input ABI used to generate the binding from.
	L2ScrollMessengerABI, _ = L2ScrollMessengerMetaData.GetAbi()
)

// L2ScrollMessenger is an auto generated Go binding around an Ethereum contract.
type L2ScrollMessenger struct {
	Name    string
	Address common.Address // contract address
	ABI     *abi.ABI       // L2ScrollMessengerABI is the input ABI used to generate the binding from.

	topics  map[common.Hash]string
	parsers map[common.Hash]func(log *types.Log) error

	L2ScrollMessengerCaller     // Read-only binding to the contract
	L2ScrollMessengerTransactor // Write-only binding to the contract
}

// GetName return L2ScrollMessenger's Name.
func (o *L2ScrollMessenger) GetName() string {
	return o.Name
}

// GetAddress return L2ScrollMessenger's contract address.
func (o *L2ScrollMessenger) GetAddress() common.Address {
	return o.Address
}

// GetEventName get event name by event hash.
func (o *L2ScrollMessenger) GetEventName(sigHash common.Hash) string {
	return o.topics[sigHash]
}

// GetSigHashes get sig hash list.
func (o *L2ScrollMessenger) GetSigHashes() []common.Hash {
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
func (o *L2ScrollMessenger) ParseLog(vLog *types.Log) (bool, error) {
	_id := vLog.Topics[0]
	if parse, exist := o.parsers[_id]; exist {
		return true, parse(vLog)
	}
	return false, nil
}

// RegisterFailedRelayedMessage, the FailedRelayedMessage event ID is 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
func (o *L2ScrollMessenger) RegisterFailedRelayedMessage(handler func(vLog *types.Log, data *L2ScrollMessengerFailedRelayedMessageEvent) error) {
	_id := o.ABI.Events["FailedRelayedMessage"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L2ScrollMessengerFailedRelayedMessageEvent)
		if err := o.L2ScrollMessengerCaller.contract.UnpackLog(event, "FailedRelayedMessage", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "FailedRelayedMessage"
}

// RegisterInitialized, the Initialized event ID is 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
func (o *L2ScrollMessenger) RegisterInitialized(handler func(vLog *types.Log, data *L2ScrollMessengerInitializedEvent) error) {
	_id := o.ABI.Events["Initialized"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L2ScrollMessengerInitializedEvent)
		if err := o.L2ScrollMessengerCaller.contract.UnpackLog(event, "Initialized", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "Initialized"
}

// RegisterOwnershipTransferred, the OwnershipTransferred event ID is 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
func (o *L2ScrollMessenger) RegisterOwnershipTransferred(handler func(vLog *types.Log, data *L2ScrollMessengerOwnershipTransferredEvent) error) {
	_id := o.ABI.Events["OwnershipTransferred"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L2ScrollMessengerOwnershipTransferredEvent)
		if err := o.L2ScrollMessengerCaller.contract.UnpackLog(event, "OwnershipTransferred", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "OwnershipTransferred"
}

// RegisterPaused, the Paused event ID is 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
func (o *L2ScrollMessenger) RegisterPaused(handler func(vLog *types.Log, data *L2ScrollMessengerPausedEvent) error) {
	_id := o.ABI.Events["Paused"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L2ScrollMessengerPausedEvent)
		if err := o.L2ScrollMessengerCaller.contract.UnpackLog(event, "Paused", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "Paused"
}

// RegisterRelayedMessage, the RelayedMessage event ID is 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
func (o *L2ScrollMessenger) RegisterRelayedMessage(handler func(vLog *types.Log, data *L2ScrollMessengerRelayedMessageEvent) error) {
	_id := o.ABI.Events["RelayedMessage"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L2ScrollMessengerRelayedMessageEvent)
		if err := o.L2ScrollMessengerCaller.contract.UnpackLog(event, "RelayedMessage", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "RelayedMessage"
}

// RegisterSentMessage, the SentMessage event ID is 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
func (o *L2ScrollMessenger) RegisterSentMessage(handler func(vLog *types.Log, data *L2ScrollMessengerSentMessageEvent) error) {
	_id := o.ABI.Events["SentMessage"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L2ScrollMessengerSentMessageEvent)
		if err := o.L2ScrollMessengerCaller.contract.UnpackLog(event, "SentMessage", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "SentMessage"
}

// RegisterUnpaused, the Unpaused event ID is 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
func (o *L2ScrollMessenger) RegisterUnpaused(handler func(vLog *types.Log, data *L2ScrollMessengerUnpausedEvent) error) {
	_id := o.ABI.Events["Unpaused"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L2ScrollMessengerUnpausedEvent)
		if err := o.L2ScrollMessengerCaller.contract.UnpackLog(event, "Unpaused", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "Unpaused"
}

// RegisterUpdateFeeVault, the UpdateFeeVault event ID is 0x4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f5.
func (o *L2ScrollMessenger) RegisterUpdateFeeVault(handler func(vLog *types.Log, data *L2ScrollMessengerUpdateFeeVaultEvent) error) {
	_id := o.ABI.Events["UpdateFeeVault"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L2ScrollMessengerUpdateFeeVaultEvent)
		if err := o.L2ScrollMessengerCaller.contract.UnpackLog(event, "UpdateFeeVault", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "UpdateFeeVault"
}

// RegisterUpdateMaxFailedExecutionTimes, the UpdateMaxFailedExecutionTimes event ID is 0x8a4c22c9b46f23dedd49b843839940ce0c36fa1612073a9bc7dbaeef9ee547ba.
func (o *L2ScrollMessenger) RegisterUpdateMaxFailedExecutionTimes(handler func(vLog *types.Log, data *L2ScrollMessengerUpdateMaxFailedExecutionTimesEvent) error) {
	_id := o.ABI.Events["UpdateMaxFailedExecutionTimes"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L2ScrollMessengerUpdateMaxFailedExecutionTimesEvent)
		if err := o.L2ScrollMessengerCaller.contract.UnpackLog(event, "UpdateMaxFailedExecutionTimes", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "UpdateMaxFailedExecutionTimes"
}

// RegisterUpdateRateLimiter, the UpdateRateLimiter event ID is 0x53d49a5b37a3ed13692a8b217b8c79f53b0144c069de23e0cc5704d89bc23006.
func (o *L2ScrollMessenger) RegisterUpdateRateLimiter(handler func(vLog *types.Log, data *L2ScrollMessengerUpdateRateLimiterEvent) error) {
	_id := o.ABI.Events["UpdateRateLimiter"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L2ScrollMessengerUpdateRateLimiterEvent)
		if err := o.L2ScrollMessengerCaller.contract.UnpackLog(event, "UpdateRateLimiter", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "UpdateRateLimiter"
}

// L2ScrollMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ScrollMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ScrollMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ScrollMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewL2ScrollMessenger creates a new instance of L2ScrollMessenger, bound to a specific deployed contract.
func NewL2ScrollMessenger(address common.Address, backend bind.ContractBackend) (*L2ScrollMessenger, error) {
	contract, err := bindL2ScrollMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	sigAbi, err := L2ScrollMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	topics := make(map[common.Hash]string, len(sigAbi.Events))
	for name, _abi := range sigAbi.Events {
		topics[_abi.ID] = name
	}

	return &L2ScrollMessenger{
		Name:                        "L2ScrollMessenger",
		ABI:                         sigAbi,
		Address:                     address,
		topics:                      topics,
		parsers:                     map[common.Hash]func(log *types.Log) error{},
		L2ScrollMessengerCaller:     L2ScrollMessengerCaller{contract: contract},
		L2ScrollMessengerTransactor: L2ScrollMessengerTransactor{contract: contract},
	}, nil
}

// NewL2ScrollMessengerCaller creates a new read-only instance of L2ScrollMessenger, bound to a specific deployed contract.
func NewL2ScrollMessengerCaller(address common.Address, caller bind.ContractCaller) (*L2ScrollMessengerCaller, error) {
	contract, err := bindL2ScrollMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ScrollMessengerCaller{contract: contract}, nil
}

// NewL2ScrollMessengerTransactor creates a new write-only instance of L2ScrollMessenger, bound to a specific deployed contract.
func NewL2ScrollMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*L2ScrollMessengerTransactor, error) {
	contract, err := bindL2ScrollMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ScrollMessengerTransactor{contract: contract}, nil
}

// bindL2ScrollMessenger binds a generic wrapper to an already deployed contract.
func bindL2ScrollMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2ScrollMessengerMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2ScrollMessenger *L2ScrollMessengerCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ScrollMessenger.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_L2ScrollMessenger *L2ScrollMessengerCaller) FeeVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ScrollMessenger.contract.Call(opts, &out, "feeVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IsL1MessageExecuted is a free data retrieval call binding the contract method 0x02345b50.
//
// Solidity: function isL1MessageExecuted(bytes32 ) view returns(bool)
func (_L2ScrollMessenger *L2ScrollMessengerCaller) IsL1MessageExecuted(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L2ScrollMessenger.contract.Call(opts, &out, "isL1MessageExecuted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsL2MessageSent is a free data retrieval call binding the contract method 0x84a7d81f.
//
// Solidity: function isL2MessageSent(bytes32 ) view returns(bool)
func (_L2ScrollMessenger *L2ScrollMessengerCaller) IsL2MessageSent(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L2ScrollMessenger.contract.Call(opts, &out, "isL2MessageSent", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MessageQueue is a free data retrieval call binding the contract method 0x3b70c18a.
//
// Solidity: function messageQueue() view returns(address)
func (_L2ScrollMessenger *L2ScrollMessengerCaller) MessageQueue(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ScrollMessenger.contract.Call(opts, &out, "messageQueue")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2ScrollMessenger *L2ScrollMessengerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ScrollMessenger.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2ScrollMessenger *L2ScrollMessengerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L2ScrollMessenger.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RateLimiter is a free data retrieval call binding the contract method 0x53d4fe33.
//
// Solidity: function rateLimiter() view returns(address)
func (_L2ScrollMessenger *L2ScrollMessengerCaller) RateLimiter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ScrollMessenger.contract.Call(opts, &out, "rateLimiter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_L2ScrollMessenger *L2ScrollMessengerCaller) XDomainMessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ScrollMessenger.contract.Call(opts, &out, "xDomainMessageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _counterpart) returns()
func (_L2ScrollMessenger *L2ScrollMessengerTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address) (*types.Transaction, error) {
	return _L2ScrollMessenger.contract.Transact(opts, "initialize", _counterpart)
}

// RelayMessage is a paid mutator transaction binding the contract method 0x8ef1332e.
//
// Solidity: function relayMessage(address _from, address _to, uint256 _value, uint256 _nonce, bytes _message) returns()
func (_L2ScrollMessenger *L2ScrollMessengerTransactor) RelayMessage(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int, _nonce *big.Int, _message []byte) (*types.Transaction, error) {
	return _L2ScrollMessenger.contract.Transact(opts, "relayMessage", _from, _to, _value, _nonce, _message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2ScrollMessenger *L2ScrollMessengerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ScrollMessenger.contract.Transact(opts, "renounceOwnership")
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address _to, uint256 _value, bytes _message, uint256 _gasLimit, address ) payable returns()
func (_L2ScrollMessenger *L2ScrollMessengerTransactor) SendMessage(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _message []byte, _gasLimit *big.Int, arg4 common.Address) (*types.Transaction, error) {
	return _L2ScrollMessenger.contract.Transact(opts, "sendMessage", _to, _value, _message, _gasLimit, arg4)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address _to, uint256 _value, bytes _message, uint256 _gasLimit) payable returns()
func (_L2ScrollMessenger *L2ScrollMessengerTransactor) SendMessage0(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _message []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ScrollMessenger.contract.Transact(opts, "sendMessage0", _to, _value, _message, _gasLimit)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_L2ScrollMessenger *L2ScrollMessengerTransactor) SetPause(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _L2ScrollMessenger.contract.Transact(opts, "setPause", _status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2ScrollMessenger *L2ScrollMessengerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2ScrollMessenger.contract.Transact(opts, "transferOwnership", newOwner)
}

// UpdateFeeVault is a paid mutator transaction binding the contract method 0x2a6cccb2.
//
// Solidity: function updateFeeVault(address _newFeeVault) returns()
func (_L2ScrollMessenger *L2ScrollMessengerTransactor) UpdateFeeVault(opts *bind.TransactOpts, _newFeeVault common.Address) (*types.Transaction, error) {
	return _L2ScrollMessenger.contract.Transact(opts, "updateFeeVault", _newFeeVault)
}

// UpdateRateLimiter is a paid mutator transaction binding the contract method 0xe157820a.
//
// Solidity: function updateRateLimiter(address _newRateLimiter) returns()
func (_L2ScrollMessenger *L2ScrollMessengerTransactor) UpdateRateLimiter(opts *bind.TransactOpts, _newRateLimiter common.Address) (*types.Transaction, error) {
	return _L2ScrollMessenger.contract.Transact(opts, "updateRateLimiter", _newRateLimiter)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2ScrollMessenger *L2ScrollMessengerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ScrollMessenger.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// L2ScrollMessengerFailedRelayedMessage represents a FailedRelayedMessage event raised by the L2ScrollMessenger contract.
type L2ScrollMessengerFailedRelayedMessageEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	MessageHash [32]byte
}

// L2ScrollMessengerInitialized represents a Initialized event raised by the L2ScrollMessenger contract.
type L2ScrollMessengerInitializedEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	Version uint8
}

// L2ScrollMessengerOwnershipTransferred represents a OwnershipTransferred event raised by the L2ScrollMessenger contract.
type L2ScrollMessengerOwnershipTransferredEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	PreviousOwner common.Address
	NewOwner      common.Address
}

// L2ScrollMessengerPaused represents a Paused event raised by the L2ScrollMessenger contract.
type L2ScrollMessengerPausedEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	Account common.Address
}

// L2ScrollMessengerRelayedMessage represents a RelayedMessage event raised by the L2ScrollMessenger contract.
type L2ScrollMessengerRelayedMessageEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	MessageHash [32]byte
}

// L2ScrollMessengerSentMessage represents a SentMessage event raised by the L2ScrollMessenger contract.
type L2ScrollMessengerSentMessageEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	Sender       common.Address
	Target       common.Address
	Value        *big.Int
	MessageNonce *big.Int
	GasLimit     *big.Int
	Message      []byte
}

// L2ScrollMessengerUnpaused represents a Unpaused event raised by the L2ScrollMessenger contract.
type L2ScrollMessengerUnpausedEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	Account common.Address
}

// L2ScrollMessengerUpdateFeeVault represents a UpdateFeeVault event raised by the L2ScrollMessenger contract.
type L2ScrollMessengerUpdateFeeVaultEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	OldFeeVault common.Address
	NewFeeVault common.Address
}

// L2ScrollMessengerUpdateMaxFailedExecutionTimes represents a UpdateMaxFailedExecutionTimes event raised by the L2ScrollMessenger contract.
type L2ScrollMessengerUpdateMaxFailedExecutionTimesEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	OldMaxFailedExecutionTimes *big.Int
	NewMaxFailedExecutionTimes *big.Int
}

// L2ScrollMessengerUpdateRateLimiter represents a UpdateRateLimiter event raised by the L2ScrollMessenger contract.
type L2ScrollMessengerUpdateRateLimiterEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	OldRateLimiter common.Address
	NewRateLimiter common.Address
}
