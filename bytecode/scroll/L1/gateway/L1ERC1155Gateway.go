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

// L1ERC1155GatewayMetaData contains all meta data concerning the L1ERC1155Gateway contract.
var (
	L1ERC1155GatewayMetaData = &bind.MetaData{
		ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\",\"indexed\":false},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\",\"indexed\":false}],\"type\":\"event\",\"name\":\"BatchDepositERC1155\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"indexed\":false},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\",\"indexed\":false}],\"type\":\"event\",\"name\":\"BatchRefundERC1155\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"DepositERC1155\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\",\"indexed\":false},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\",\"indexed\":false}],\"type\":\"event\",\"name\":\"FinalizeBatchWithdrawERC1155\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"FinalizeWithdrawERC1155\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Initialized\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"RefundERC1155\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"oldL2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"newL2Token\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"UpdateTokenMapping\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"batchDepositERC1155\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"batchDepositERC1155\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"depositERC1155\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"depositERC1155\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"finalizeBatchWithdrawERC1155\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"finalizeWithdrawERC1155\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"onDropMessage\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"renounceOwnership\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"tokenMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferOwnership\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateTokenMapping\"}]",
	}
)

// L1ERC1155Gateway is an auto generated Go binding around an Ethereum contract.
type L1ERC1155Gateway struct {
	Name    string
	Address common.Address // contract address
	ABI     *abi.ABI       // L1ERC1155GatewayABI is the input ABI used to generate the binding from.

	topics  map[common.Hash]string
	parsers map[common.Hash]func(log *types.Log) error

	L1ERC1155GatewayCaller     // Read-only binding to the contract
	L1ERC1155GatewayTransactor // Write-only binding to the contract
}

// GetName return L1ERC1155Gateway's Name.
func (o *L1ERC1155Gateway) GetName() string {
	return o.Name
}

// GetAddress return L1ERC1155Gateway's contract address.
func (o *L1ERC1155Gateway) GetAddress() common.Address {
	return o.Address
}

// GetEventName get event name by event hash.
func (o *L1ERC1155Gateway) GetEventName(sigHash common.Hash) string {
	return o.topics[sigHash]
}

// GetSigHashes get sig hash list.
func (o *L1ERC1155Gateway) GetSigHashes() []common.Hash {
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
func (o *L1ERC1155Gateway) ParseLog(vLog *types.Log) (bool, error) {
	_id := vLog.Topics[0]
	if parse, exist := o.parsers[_id]; exist {
		return true, parse(vLog)
	} else {
		return false, nil
	}
	return true, nil
}

// RegisterBatchDepositERC1155, the BatchDepositERC1155 event ID is 0x743f65db61a23bc629915d35e22af5cf13478a8b3dbd154d3e5db0149509756d.
func (o *L1ERC1155Gateway) RegisterBatchDepositERC1155(handler func(vLog *types.Log, data *L1ERC1155GatewayBatchDepositERC1155Event) error) {
	_id := o.ABI.Events["BatchDepositERC1155"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ERC1155GatewayBatchDepositERC1155Event)
		if err := o.L1ERC1155GatewayCaller.contract.UnpackLog(event, "BatchDepositERC1155", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
	o.topics[_id] = "BatchDepositERC1155"
}

// RegisterBatchRefundERC1155, the BatchRefundERC1155 event ID is 0xe198c04cbd4522ed7825c7e6ab1ae33fdaf6ab3565c4a3fb4c0cf24338f306e6.
func (o *L1ERC1155Gateway) RegisterBatchRefundERC1155(handler func(vLog *types.Log, data *L1ERC1155GatewayBatchRefundERC1155Event) error) {
	_id := o.ABI.Events["BatchRefundERC1155"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ERC1155GatewayBatchRefundERC1155Event)
		if err := o.L1ERC1155GatewayCaller.contract.UnpackLog(event, "BatchRefundERC1155", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
	o.topics[_id] = "BatchRefundERC1155"
}

// RegisterDepositERC1155, the DepositERC1155 event ID is 0x7f6552b688fa94306ca59e44dd4454ff550542445a3f1cb39b8c768be6f5c08a.
func (o *L1ERC1155Gateway) RegisterDepositERC1155(handler func(vLog *types.Log, data *L1ERC1155GatewayDepositERC1155Event) error) {
	_id := o.ABI.Events["DepositERC1155"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ERC1155GatewayDepositERC1155Event)
		if err := o.L1ERC1155GatewayCaller.contract.UnpackLog(event, "DepositERC1155", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
	o.topics[_id] = "DepositERC1155"
}

// RegisterFinalizeBatchWithdrawERC1155, the FinalizeBatchWithdrawERC1155 event ID is 0x45294b6ad6ad2408cc3ee9a37203aa1b0480616667a97b157c52ac9294cbc258.
func (o *L1ERC1155Gateway) RegisterFinalizeBatchWithdrawERC1155(handler func(vLog *types.Log, data *L1ERC1155GatewayFinalizeBatchWithdrawERC1155Event) error) {
	_id := o.ABI.Events["FinalizeBatchWithdrawERC1155"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ERC1155GatewayFinalizeBatchWithdrawERC1155Event)
		if err := o.L1ERC1155GatewayCaller.contract.UnpackLog(event, "FinalizeBatchWithdrawERC1155", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
	o.topics[_id] = "FinalizeBatchWithdrawERC1155"
}

// RegisterFinalizeWithdrawERC1155, the FinalizeWithdrawERC1155 event ID is 0xfcc2841e9e72e6d610944e1b668912e92d5df94003055dbe06d615ba8d9efad4.
func (o *L1ERC1155Gateway) RegisterFinalizeWithdrawERC1155(handler func(vLog *types.Log, data *L1ERC1155GatewayFinalizeWithdrawERC1155Event) error) {
	_id := o.ABI.Events["FinalizeWithdrawERC1155"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ERC1155GatewayFinalizeWithdrawERC1155Event)
		if err := o.L1ERC1155GatewayCaller.contract.UnpackLog(event, "FinalizeWithdrawERC1155", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
	o.topics[_id] = "FinalizeWithdrawERC1155"
}

// RegisterInitialized, the Initialized event ID is 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
func (o *L1ERC1155Gateway) RegisterInitialized(handler func(vLog *types.Log, data *L1ERC1155GatewayInitializedEvent) error) {
	_id := o.ABI.Events["Initialized"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ERC1155GatewayInitializedEvent)
		if err := o.L1ERC1155GatewayCaller.contract.UnpackLog(event, "Initialized", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
	o.topics[_id] = "Initialized"
}

// RegisterOwnershipTransferred, the OwnershipTransferred event ID is 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
func (o *L1ERC1155Gateway) RegisterOwnershipTransferred(handler func(vLog *types.Log, data *L1ERC1155GatewayOwnershipTransferredEvent) error) {
	_id := o.ABI.Events["OwnershipTransferred"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ERC1155GatewayOwnershipTransferredEvent)
		if err := o.L1ERC1155GatewayCaller.contract.UnpackLog(event, "OwnershipTransferred", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
	o.topics[_id] = "OwnershipTransferred"
}

// RegisterRefundERC1155, the RefundERC1155 event ID is 0xee285671d9ac3b0e0ed40037cb6db081095aa6cd68363f3e56989dde39e0df09.
func (o *L1ERC1155Gateway) RegisterRefundERC1155(handler func(vLog *types.Log, data *L1ERC1155GatewayRefundERC1155Event) error) {
	_id := o.ABI.Events["RefundERC1155"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ERC1155GatewayRefundERC1155Event)
		if err := o.L1ERC1155GatewayCaller.contract.UnpackLog(event, "RefundERC1155", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
	o.topics[_id] = "RefundERC1155"
}

// RegisterUpdateTokenMapping, the UpdateTokenMapping event ID is 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
func (o *L1ERC1155Gateway) RegisterUpdateTokenMapping(handler func(vLog *types.Log, data *L1ERC1155GatewayUpdateTokenMappingEvent) error) {
	_id := o.ABI.Events["UpdateTokenMapping"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(L1ERC1155GatewayUpdateTokenMappingEvent)
		if err := o.L1ERC1155GatewayCaller.contract.UnpackLog(event, "UpdateTokenMapping", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
	o.topics[_id] = "UpdateTokenMapping"
}

// L1ERC1155GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1ERC1155GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC1155GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1ERC1155GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewL1ERC1155Gateway creates a new instance of L1ERC1155Gateway, bound to a specific deployed contract.
func NewL1ERC1155Gateway(address common.Address, backend bind.ContractBackend) (*L1ERC1155Gateway, error) {
	contract, err := bindL1ERC1155Gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	sigAbi, err := L1ERC1155GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	topics := make(map[common.Hash]string, len(sigAbi.Events))
	for name, _abi := range sigAbi.Events {
		topics[_abi.ID] = name
	}

	return &L1ERC1155Gateway{
		Name:                       "L1ERC1155Gateway",
		ABI:                        sigAbi,
		Address:                    address,
		topics:                     topics,
		parsers:                    map[common.Hash]func(log *types.Log) error{},
		L1ERC1155GatewayCaller:     L1ERC1155GatewayCaller{contract: contract},
		L1ERC1155GatewayTransactor: L1ERC1155GatewayTransactor{contract: contract},
	}, nil
}

// NewL1ERC1155GatewayCaller creates a new read-only instance of L1ERC1155Gateway, bound to a specific deployed contract.
func NewL1ERC1155GatewayCaller(address common.Address, caller bind.ContractCaller) (*L1ERC1155GatewayCaller, error) {
	contract, err := bindL1ERC1155Gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1ERC1155GatewayCaller{contract: contract}, nil
}

// NewL1ERC1155GatewayTransactor creates a new write-only instance of L1ERC1155Gateway, bound to a specific deployed contract.
func NewL1ERC1155GatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L1ERC1155GatewayTransactor, error) {
	contract, err := bindL1ERC1155Gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1ERC1155GatewayTransactor{contract: contract}, nil
}

// bindL1ERC1155Gateway binds a generic wrapper to an already deployed contract.
func bindL1ERC1155Gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1ERC1155GatewayMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1ERC1155Gateway *L1ERC1155GatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC1155Gateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L1ERC1155Gateway *L1ERC1155GatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC1155Gateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1ERC1155Gateway *L1ERC1155GatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC1155Gateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1ERC1155Gateway *L1ERC1155GatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC1155Gateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L1ERC1155Gateway *L1ERC1155GatewayCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _L1ERC1155Gateway.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L1ERC1155Gateway *L1ERC1155GatewayCaller) TokenMapping(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1ERC1155Gateway.contract.Call(opts, &out, "tokenMapping", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BatchDepositERC1155 is a paid mutator transaction binding the contract method 0x5ee8e74c.
//
// Solidity: function batchDepositERC1155(address _token, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_L1ERC1155Gateway *L1ERC1155GatewayTransactor) BatchDepositERC1155(opts *bind.TransactOpts, _token common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC1155Gateway.contract.Transact(opts, "batchDepositERC1155", _token, _tokenIds, _amounts, _gasLimit)
}

// BatchDepositERC11550 is a paid mutator transaction binding the contract method 0xc99dac9b.
//
// Solidity: function batchDepositERC1155(address _token, address _to, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_L1ERC1155Gateway *L1ERC1155GatewayTransactor) BatchDepositERC11550(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC1155Gateway.contract.Transact(opts, "batchDepositERC11550", _token, _to, _tokenIds, _amounts, _gasLimit)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0xa901cf8a.
//
// Solidity: function depositERC1155(address _token, address _to, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1ERC1155Gateway *L1ERC1155GatewayTransactor) DepositERC1155(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC1155Gateway.contract.Transact(opts, "depositERC1155", _token, _to, _tokenId, _amount, _gasLimit)
}

// DepositERC11550 is a paid mutator transaction binding the contract method 0xf998fe9d.
//
// Solidity: function depositERC1155(address _token, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L1ERC1155Gateway *L1ERC1155GatewayTransactor) DepositERC11550(opts *bind.TransactOpts, _token common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ERC1155Gateway.contract.Transact(opts, "depositERC11550", _token, _tokenId, _amount, _gasLimit)
}

// FinalizeBatchWithdrawERC1155 is a paid mutator transaction binding the contract method 0xf92748d3.
//
// Solidity: function finalizeBatchWithdrawERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds, uint256[] _amounts) returns()
func (_L1ERC1155Gateway *L1ERC1155GatewayTransactor) FinalizeBatchWithdrawERC1155(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _L1ERC1155Gateway.contract.Transact(opts, "finalizeBatchWithdrawERC1155", _l1Token, _l2Token, _from, _to, _tokenIds, _amounts)
}

// FinalizeWithdrawERC1155 is a paid mutator transaction binding the contract method 0x730608b3.
//
// Solidity: function finalizeWithdrawERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId, uint256 _amount) returns()
func (_L1ERC1155Gateway *L1ERC1155GatewayTransactor) FinalizeWithdrawERC1155(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _L1ERC1155Gateway.contract.Transact(opts, "finalizeWithdrawERC1155", _l1Token, _l2Token, _from, _to, _tokenId, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L1ERC1155Gateway *L1ERC1155GatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L1ERC1155Gateway.contract.Transact(opts, "initialize", _counterpart, _messenger)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes _message) payable returns()
func (_L1ERC1155Gateway *L1ERC1155GatewayTransactor) OnDropMessage(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _L1ERC1155Gateway.contract.Transact(opts, "onDropMessage", _message)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_L1ERC1155Gateway *L1ERC1155GatewayTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L1ERC1155Gateway.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_L1ERC1155Gateway *L1ERC1155GatewayTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L1ERC1155Gateway.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1ERC1155Gateway *L1ERC1155GatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC1155Gateway.contract.Transact(opts, "renounceOwnership")
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1ERC1155Gateway *L1ERC1155GatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1ERC1155Gateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l1Token, address _l2Token) returns()
func (_L1ERC1155Gateway *L1ERC1155GatewayTransactor) UpdateTokenMapping(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address) (*types.Transaction, error) {
	return _L1ERC1155Gateway.contract.Transact(opts, "updateTokenMapping", _l1Token, _l2Token)
}

// L1ERC1155GatewayBatchDepositERC1155 represents a BatchDepositERC1155 event raised by the L1ERC1155Gateway contract.
type L1ERC1155GatewayBatchDepositERC1155Event struct {
	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Amounts  []*big.Int
}

// L1ERC1155GatewayBatchRefundERC1155 represents a BatchRefundERC1155 event raised by the L1ERC1155Gateway contract.
type L1ERC1155GatewayBatchRefundERC1155Event struct {
	Token     common.Address
	Recipient common.Address
	TokenIds  []*big.Int
	Amounts   []*big.Int
}

// L1ERC1155GatewayDepositERC1155 represents a DepositERC1155 event raised by the L1ERC1155Gateway contract.
type L1ERC1155GatewayDepositERC1155Event struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Amount  *big.Int
}

// L1ERC1155GatewayFinalizeBatchWithdrawERC1155 represents a FinalizeBatchWithdrawERC1155 event raised by the L1ERC1155Gateway contract.
type L1ERC1155GatewayFinalizeBatchWithdrawERC1155Event struct {
	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Amounts  []*big.Int
}

// L1ERC1155GatewayFinalizeWithdrawERC1155 represents a FinalizeWithdrawERC1155 event raised by the L1ERC1155Gateway contract.
type L1ERC1155GatewayFinalizeWithdrawERC1155Event struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Amount  *big.Int
}

// L1ERC1155GatewayInitialized represents a Initialized event raised by the L1ERC1155Gateway contract.
type L1ERC1155GatewayInitializedEvent struct {
	Version uint8
}

// L1ERC1155GatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L1ERC1155Gateway contract.
type L1ERC1155GatewayOwnershipTransferredEvent struct {
	PreviousOwner common.Address
	NewOwner      common.Address
}

// L1ERC1155GatewayRefundERC1155 represents a RefundERC1155 event raised by the L1ERC1155Gateway contract.
type L1ERC1155GatewayRefundERC1155Event struct {
	Token     common.Address
	Recipient common.Address
	TokenId   *big.Int
	Amount    *big.Int
}

// L1ERC1155GatewayUpdateTokenMapping represents a UpdateTokenMapping event raised by the L1ERC1155Gateway contract.
type L1ERC1155GatewayUpdateTokenMappingEvent struct {
	L1Token    common.Address
	OldL2Token common.Address
	NewL2Token common.Address
}
