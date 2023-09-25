// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package L1

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

// IL1ScrollMessengerL2MessageProof is an auto generated low-level Go binding around an user-defined struct.
type IL1ScrollMessengerL2MessageProof struct {
	BatchIndex  *big.Int
	MerkleProof []byte
}

// L1ScrollMessengerMetaData contains all meta data concerning the L1ScrollMessenger contract.
var (
	L1ScrollMessengerMetaData = &bind.MetaData{
		ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\",\"indexed\":true}],\"type\":\"event\",\"name\":\"FailedRelayedMessage\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Initialized\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Paused\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\",\"indexed\":true}],\"type\":\"event\",\"name\":\"RelayedMessage\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false}],\"type\":\"event\",\"name\":\"SentMessage\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Unpaused\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldFeeVault\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"address\",\"name\":\"_newFeeVault\",\"type\":\"address\",\"indexed\":false}],\"type\":\"event\",\"name\":\"UpdateFeeVault\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"oldMaxReplayTimes\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"newMaxReplayTimes\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"UpdateMaxReplayTimes\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldRateLimiter\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_newRateLimiter\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"UpdateRateLimiter\",\"anonymous\":false},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"dropMessage\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"feeVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messageQueue\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"isL1MessageDropped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}]},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"isL1MessageSent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}]},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"isL2MessageExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"maxReplayTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"messageQueue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}]},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"prevReplayIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"rateLimiter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"structIL1ScrollMessenger.L2MessageProof\",\"name\":\"_proof\",\"type\":\"tuple\",\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"merkleProof\",\"type\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"relayMessageWithProof\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"renounceOwnership\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"_newGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"replayMessage\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"replayStates\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"times\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"lastIndex\",\"type\":\"uint128\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"sendMessage\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"sendMessage\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setPause\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferOwnership\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFeeVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateFeeVault\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newMaxReplayTimes\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateMaxReplayTimes\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRateLimiter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateRateLimiter\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"xDomainMessageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	}
	// L1ScrollMessengerABI is the input ABI used to generate the binding from.
	L1ScrollMessengerABI, _ = L1ScrollMessengerMetaData.GetAbi()
)

// L1ScrollMessenger is an auto generated Go binding around an Ethereum contract.
type L1ScrollMessenger struct {
	Name    string
	Address common.Address // contract address
	ABI     *abi.ABI       // L1ScrollMessengerABI is the input ABI used to generate the binding from.

	topics  map[common.Hash]string
	parsers map[common.Hash]func(log *types.Log) error

	L1ScrollMessengerCaller     // Read-only binding to the contract
	L1ScrollMessengerTransactor // Write-only binding to the contract
}

// GetName return L1ScrollMessenger's Name.
func (o *L1ScrollMessenger) GetName() string {
	return o.Name
}

// GetAddress return L1ScrollMessenger's contract address.
func (o *L1ScrollMessenger) GetAddress() common.Address {
	return o.Address
}

// GetEventName get event name by event hash.
func (o *L1ScrollMessenger) GetEventName(sigHash common.Hash) string {
	return o.topics[sigHash]
}

// GetSigHashes get sig hash list.
func (o *L1ScrollMessenger) GetSigHashes() []common.Hash {
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
func (o *L1ScrollMessenger) ParseLog(vLog *types.Log) (bool, error) {
	_id := vLog.Topics[0]
	if parse, exist := o.parsers[_id]; exist {
		return true, parse(vLog)
	}
	return false, nil
}

// RegisterFailedRelayedMessage, the FailedRelayedMessage event ID is 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
func (o *L1ScrollMessenger) RegisterFailedRelayedMessage(handler func(vLog *types.Log, data *L1ScrollMessengerFailedRelayedMessageEvent) error) {
	_id := o.ABI.Events["FailedRelayedMessage"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ScrollMessengerFailedRelayedMessageEvent)
		if err := o.L1ScrollMessengerCaller.contract.UnpackLog(event, "FailedRelayedMessage", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "FailedRelayedMessage"
}

// RegisterInitialized, the Initialized event ID is 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
func (o *L1ScrollMessenger) RegisterInitialized(handler func(vLog *types.Log, data *L1ScrollMessengerInitializedEvent) error) {
	_id := o.ABI.Events["Initialized"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ScrollMessengerInitializedEvent)
		if err := o.L1ScrollMessengerCaller.contract.UnpackLog(event, "Initialized", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "Initialized"
}

// RegisterOwnershipTransferred, the OwnershipTransferred event ID is 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
func (o *L1ScrollMessenger) RegisterOwnershipTransferred(handler func(vLog *types.Log, data *L1ScrollMessengerOwnershipTransferredEvent) error) {
	_id := o.ABI.Events["OwnershipTransferred"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ScrollMessengerOwnershipTransferredEvent)
		if err := o.L1ScrollMessengerCaller.contract.UnpackLog(event, "OwnershipTransferred", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "OwnershipTransferred"
}

// RegisterPaused, the Paused event ID is 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
func (o *L1ScrollMessenger) RegisterPaused(handler func(vLog *types.Log, data *L1ScrollMessengerPausedEvent) error) {
	_id := o.ABI.Events["Paused"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ScrollMessengerPausedEvent)
		if err := o.L1ScrollMessengerCaller.contract.UnpackLog(event, "Paused", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "Paused"
}

// RegisterRelayedMessage, the RelayedMessage event ID is 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
func (o *L1ScrollMessenger) RegisterRelayedMessage(handler func(vLog *types.Log, data *L1ScrollMessengerRelayedMessageEvent) error) {
	_id := o.ABI.Events["RelayedMessage"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ScrollMessengerRelayedMessageEvent)
		if err := o.L1ScrollMessengerCaller.contract.UnpackLog(event, "RelayedMessage", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "RelayedMessage"
}

// RegisterSentMessage, the SentMessage event ID is 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
func (o *L1ScrollMessenger) RegisterSentMessage(handler func(vLog *types.Log, data *L1ScrollMessengerSentMessageEvent) error) {
	_id := o.ABI.Events["SentMessage"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ScrollMessengerSentMessageEvent)
		if err := o.L1ScrollMessengerCaller.contract.UnpackLog(event, "SentMessage", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "SentMessage"
}

// RegisterUnpaused, the Unpaused event ID is 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
func (o *L1ScrollMessenger) RegisterUnpaused(handler func(vLog *types.Log, data *L1ScrollMessengerUnpausedEvent) error) {
	_id := o.ABI.Events["Unpaused"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ScrollMessengerUnpausedEvent)
		if err := o.L1ScrollMessengerCaller.contract.UnpackLog(event, "Unpaused", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "Unpaused"
}

// RegisterUpdateFeeVault, the UpdateFeeVault event ID is 0x4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f5.
func (o *L1ScrollMessenger) RegisterUpdateFeeVault(handler func(vLog *types.Log, data *L1ScrollMessengerUpdateFeeVaultEvent) error) {
	_id := o.ABI.Events["UpdateFeeVault"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ScrollMessengerUpdateFeeVaultEvent)
		if err := o.L1ScrollMessengerCaller.contract.UnpackLog(event, "UpdateFeeVault", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "UpdateFeeVault"
}

// RegisterUpdateMaxReplayTimes, the UpdateMaxReplayTimes event ID is 0xd700562df02eb66951f6f5275df7ebd7c0ec58b3422915789b3b1877aab2e52b.
func (o *L1ScrollMessenger) RegisterUpdateMaxReplayTimes(handler func(vLog *types.Log, data *L1ScrollMessengerUpdateMaxReplayTimesEvent) error) {
	_id := o.ABI.Events["UpdateMaxReplayTimes"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ScrollMessengerUpdateMaxReplayTimesEvent)
		if err := o.L1ScrollMessengerCaller.contract.UnpackLog(event, "UpdateMaxReplayTimes", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "UpdateMaxReplayTimes"
}

// RegisterUpdateRateLimiter, the UpdateRateLimiter event ID is 0x53d49a5b37a3ed13692a8b217b8c79f53b0144c069de23e0cc5704d89bc23006.
func (o *L1ScrollMessenger) RegisterUpdateRateLimiter(handler func(vLog *types.Log, data *L1ScrollMessengerUpdateRateLimiterEvent) error) {
	_id := o.ABI.Events["UpdateRateLimiter"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ScrollMessengerUpdateRateLimiterEvent)
		if err := o.L1ScrollMessengerCaller.contract.UnpackLog(event, "UpdateRateLimiter", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "UpdateRateLimiter"
}

// L1ScrollMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1ScrollMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ScrollMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1ScrollMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewL1ScrollMessenger creates a new instance of L1ScrollMessenger, bound to a specific deployed contract.
func NewL1ScrollMessenger(address common.Address, backend bind.ContractBackend) (*L1ScrollMessenger, error) {
	contract, err := bindL1ScrollMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	sigAbi, err := L1ScrollMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	topics := make(map[common.Hash]string, len(sigAbi.Events))
	for name, _abi := range sigAbi.Events {
		topics[_abi.ID] = name
	}

	return &L1ScrollMessenger{
		Name:                        "L1ScrollMessenger",
		ABI:                         sigAbi,
		Address:                     address,
		topics:                      topics,
		parsers:                     map[common.Hash]func(log *types.Log) error{},
		L1ScrollMessengerCaller:     L1ScrollMessengerCaller{contract: contract},
		L1ScrollMessengerTransactor: L1ScrollMessengerTransactor{contract: contract},
	}, nil
}

// NewL1ScrollMessengerCaller creates a new read-only instance of L1ScrollMessenger, bound to a specific deployed contract.
func NewL1ScrollMessengerCaller(address common.Address, caller bind.ContractCaller) (*L1ScrollMessengerCaller, error) {
	contract, err := bindL1ScrollMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1ScrollMessengerCaller{contract: contract}, nil
}

// NewL1ScrollMessengerTransactor creates a new write-only instance of L1ScrollMessenger, bound to a specific deployed contract.
func NewL1ScrollMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*L1ScrollMessengerTransactor, error) {
	contract, err := bindL1ScrollMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1ScrollMessengerTransactor{contract: contract}, nil
}

// bindL1ScrollMessenger binds a generic wrapper to an already deployed contract.
func bindL1ScrollMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1ScrollMessengerMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) FeeVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "feeVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IsL1MessageDropped is a free data retrieval call binding the contract method 0xb604bf4c.
//
// Solidity: function isL1MessageDropped(bytes32 ) view returns(bool)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) IsL1MessageDropped(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "isL1MessageDropped", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsL1MessageSent is a free data retrieval call binding the contract method 0x69058083.
//
// Solidity: function isL1MessageSent(bytes32 ) view returns(bool)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) IsL1MessageSent(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "isL1MessageSent", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsL2MessageExecuted is a free data retrieval call binding the contract method 0x088681a7.
//
// Solidity: function isL2MessageExecuted(bytes32 ) view returns(bool)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) IsL2MessageExecuted(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "isL2MessageExecuted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MaxReplayTimes is a free data retrieval call binding the contract method 0x946130d8.
//
// Solidity: function maxReplayTimes() view returns(uint256)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) MaxReplayTimes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "maxReplayTimes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageQueue is a free data retrieval call binding the contract method 0x3b70c18a.
//
// Solidity: function messageQueue() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) MessageQueue(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "messageQueue")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PrevReplayIndex is a free data retrieval call binding the contract method 0xea7ec514.
//
// Solidity: function prevReplayIndex(uint256 ) view returns(uint256)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) PrevReplayIndex(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "prevReplayIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RateLimiter is a free data retrieval call binding the contract method 0x53d4fe33.
//
// Solidity: function rateLimiter() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) RateLimiter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "rateLimiter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReplayStates is a free data retrieval call binding the contract method 0x846d4d7a.
//
// Solidity: function replayStates(bytes32 ) view returns(uint128 times, uint128 lastIndex)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) ReplayStates(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Times     *big.Int
	LastIndex *big.Int
}, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "replayStates", arg0)

	outstruct := new(struct {
		Times     *big.Int
		LastIndex *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Times = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) XDomainMessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "xDomainMessageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DropMessage is a paid mutator transaction binding the contract method 0x29907acd.
//
// Solidity: function dropMessage(address _from, address _to, uint256 _value, uint256 _messageNonce, bytes _message) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) DropMessage(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int, _messageNonce *big.Int, _message []byte) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "dropMessage", _from, _to, _value, _messageNonce, _message)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _counterpart, address _feeVault, address _rollup, address _messageQueue) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _feeVault common.Address, _rollup common.Address, _messageQueue common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "initialize", _counterpart, _feeVault, _rollup, _messageQueue)
}

// RelayMessageWithProof is a paid mutator transaction binding the contract method 0xc311b6fc.
//
// Solidity: function relayMessageWithProof(address _from, address _to, uint256 _value, uint256 _nonce, bytes _message, (uint256,bytes) _proof) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) RelayMessageWithProof(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int, _nonce *big.Int, _message []byte, _proof IL1ScrollMessengerL2MessageProof) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "relayMessageWithProof", _from, _to, _value, _nonce, _message, _proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "renounceOwnership")
}

// ReplayMessage is a paid mutator transaction binding the contract method 0x55004105.
//
// Solidity: function replayMessage(address _from, address _to, uint256 _value, uint256 _messageNonce, bytes _message, uint32 _newGasLimit, address _refundAddress) payable returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) ReplayMessage(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int, _messageNonce *big.Int, _message []byte, _newGasLimit uint32, _refundAddress common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "replayMessage", _from, _to, _value, _messageNonce, _message, _newGasLimit, _refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address _to, uint256 _value, bytes _message, uint256 _gasLimit, address _refundAddress) payable returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) SendMessage(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _message []byte, _gasLimit *big.Int, _refundAddress common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "sendMessage", _to, _value, _message, _gasLimit, _refundAddress)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address _to, uint256 _value, bytes _message, uint256 _gasLimit) payable returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) SendMessage0(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _message []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "sendMessage0", _to, _value, _message, _gasLimit)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) SetPause(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "setPause", _status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "transferOwnership", newOwner)
}

// UpdateFeeVault is a paid mutator transaction binding the contract method 0x2a6cccb2.
//
// Solidity: function updateFeeVault(address _newFeeVault) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) UpdateFeeVault(opts *bind.TransactOpts, _newFeeVault common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "updateFeeVault", _newFeeVault)
}

// UpdateMaxReplayTimes is a paid mutator transaction binding the contract method 0x407c1955.
//
// Solidity: function updateMaxReplayTimes(uint256 _newMaxReplayTimes) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) UpdateMaxReplayTimes(opts *bind.TransactOpts, _newMaxReplayTimes *big.Int) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "updateMaxReplayTimes", _newMaxReplayTimes)
}

// UpdateRateLimiter is a paid mutator transaction binding the contract method 0xe157820a.
//
// Solidity: function updateRateLimiter(address _newRateLimiter) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) UpdateRateLimiter(opts *bind.TransactOpts, _newRateLimiter common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "updateRateLimiter", _newRateLimiter)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// L1ScrollMessengerFailedRelayedMessage represents a FailedRelayedMessage event raised by the L1ScrollMessenger contract.
type L1ScrollMessengerFailedRelayedMessageEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	MessageHash [32]byte
}

// L1ScrollMessengerInitialized represents a Initialized event raised by the L1ScrollMessenger contract.
type L1ScrollMessengerInitializedEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	Version uint8
}

// L1ScrollMessengerOwnershipTransferred represents a OwnershipTransferred event raised by the L1ScrollMessenger contract.
type L1ScrollMessengerOwnershipTransferredEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	PreviousOwner common.Address
	NewOwner      common.Address
}

// L1ScrollMessengerPaused represents a Paused event raised by the L1ScrollMessenger contract.
type L1ScrollMessengerPausedEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	Account common.Address
}

// L1ScrollMessengerRelayedMessage represents a RelayedMessage event raised by the L1ScrollMessenger contract.
type L1ScrollMessengerRelayedMessageEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	MessageHash [32]byte
}

// L1ScrollMessengerSentMessage represents a SentMessage event raised by the L1ScrollMessenger contract.
type L1ScrollMessengerSentMessageEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	Sender       common.Address
	Target       common.Address
	Value        *big.Int
	MessageNonce *big.Int
	GasLimit     *big.Int
	Message      []byte
}

// L1ScrollMessengerUnpaused represents a Unpaused event raised by the L1ScrollMessenger contract.
type L1ScrollMessengerUnpausedEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	Account common.Address
}

// L1ScrollMessengerUpdateFeeVault represents a UpdateFeeVault event raised by the L1ScrollMessenger contract.
type L1ScrollMessengerUpdateFeeVaultEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	OldFeeVault common.Address
	NewFeeVault common.Address
}

// L1ScrollMessengerUpdateMaxReplayTimes represents a UpdateMaxReplayTimes event raised by the L1ScrollMessenger contract.
type L1ScrollMessengerUpdateMaxReplayTimesEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	OldMaxReplayTimes *big.Int
	NewMaxReplayTimes *big.Int
}

// L1ScrollMessengerUpdateRateLimiter represents a UpdateRateLimiter event raised by the L1ScrollMessenger contract.
type L1ScrollMessengerUpdateRateLimiterEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	OldRateLimiter common.Address
	NewRateLimiter common.Address
}
