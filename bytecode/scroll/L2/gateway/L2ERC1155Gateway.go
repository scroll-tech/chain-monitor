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

// L2ERC1155GatewayMetaData contains all meta data concerning the L2ERC1155Gateway contract.
var (
	L2ERC1155GatewayMetaData = &bind.MetaData{
		ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"indexed\":false},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\",\"indexed\":false}],\"type\":\"event\",\"name\":\"BatchWithdrawERC1155\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"indexed\":false},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\",\"indexed\":false}],\"type\":\"event\",\"name\":\"FinalizeBatchDepositERC1155\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"FinalizeDepositERC1155\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Initialized\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"oldL1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"newL1Token\",\"type\":\"address\",\"indexed\":true}],\"type\":\"event\",\"name\":\"UpdateTokenMapping\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"WithdrawERC1155\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"batchWithdrawERC1155\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"batchWithdrawERC1155\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"finalizeBatchDepositERC1155\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"finalizeDepositERC1155\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"renounceOwnership\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"tokenMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferOwnership\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"updateTokenMapping\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"withdrawERC1155\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"withdrawERC1155\"}]",
	}
)

// L2ERC1155Gateway is an auto generated Go binding around an Ethereum contract.
type L2ERC1155Gateway struct {
	Name    string
	Address common.Address // contract address
	ABI     *abi.ABI       // L2ERC1155GatewayABI is the input ABI used to generate the binding from.

	topics  map[common.Hash]string
	parsers map[common.Hash]func(log *types.Log) error

	L2ERC1155GatewayCaller     // Read-only binding to the contract
	L2ERC1155GatewayTransactor // Write-only binding to the contract
}

// GetName return L2ERC1155Gateway's Name.
func (o *L2ERC1155Gateway) GetName() string {
	return o.Name
}

// GetAddress return L2ERC1155Gateway's contract address.
func (o *L2ERC1155Gateway) GetAddress() common.Address {
	return o.Address
}

// GetEventName get event name by event hash.
func (o *L2ERC1155Gateway) GetEventName(sigHash common.Hash) string {
	return o.topics[sigHash]
}

// GetSigHashes get sig hash list.
func (o *L2ERC1155Gateway) GetSigHashes() []common.Hash {
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
func (o *L2ERC1155Gateway) ParseLog(vLog *types.Log) error {
	_id := vLog.Topics[0]
	if parse, exist := o.parsers[_id]; exist {
		return parse(vLog)
	}
	return nil
}

// RegisterBatchWithdrawERC1155, the BatchWithdrawERC1155 event ID is 0x5d2d5d4cdbf7b115e43f0b9986644dd8b9514b10be6a019ab6a4a87f12290970.
func (o *L2ERC1155Gateway) RegisterBatchWithdrawERC1155(handler func(vLog *types.Log, data *L2ERC1155GatewayBatchWithdrawERC1155Event) error) {
	o.parsers[o.ABI.Events["BatchWithdrawERC1155"].ID] = func(log *types.Log) error {
		event := new(L2ERC1155GatewayBatchWithdrawERC1155Event)
		if err := o.L2ERC1155GatewayCaller.contract.UnpackLog(event, "BatchWithdrawERC1155", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// RegisterFinalizeBatchDepositERC1155, the FinalizeBatchDepositERC1155 event ID is 0xf07745bfeb45fb1184165136e9148689adf57ba578a5b90dde949f26066b7756.
func (o *L2ERC1155Gateway) RegisterFinalizeBatchDepositERC1155(handler func(vLog *types.Log, data *L2ERC1155GatewayFinalizeBatchDepositERC1155Event) error) {
	o.parsers[o.ABI.Events["FinalizeBatchDepositERC1155"].ID] = func(log *types.Log) error {
		event := new(L2ERC1155GatewayFinalizeBatchDepositERC1155Event)
		if err := o.L2ERC1155GatewayCaller.contract.UnpackLog(event, "FinalizeBatchDepositERC1155", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// RegisterFinalizeDepositERC1155, the FinalizeDepositERC1155 event ID is 0x5399dc7b86d085e50a28946dbc213966bb7a7ac78d312aedd6018c791ad6cef9.
func (o *L2ERC1155Gateway) RegisterFinalizeDepositERC1155(handler func(vLog *types.Log, data *L2ERC1155GatewayFinalizeDepositERC1155Event) error) {
	o.parsers[o.ABI.Events["FinalizeDepositERC1155"].ID] = func(log *types.Log) error {
		event := new(L2ERC1155GatewayFinalizeDepositERC1155Event)
		if err := o.L2ERC1155GatewayCaller.contract.UnpackLog(event, "FinalizeDepositERC1155", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// RegisterInitialized, the Initialized event ID is 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
func (o *L2ERC1155Gateway) RegisterInitialized(handler func(vLog *types.Log, data *L2ERC1155GatewayInitializedEvent) error) {
	o.parsers[o.ABI.Events["Initialized"].ID] = func(log *types.Log) error {
		event := new(L2ERC1155GatewayInitializedEvent)
		if err := o.L2ERC1155GatewayCaller.contract.UnpackLog(event, "Initialized", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// RegisterOwnershipTransferred, the OwnershipTransferred event ID is 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
func (o *L2ERC1155Gateway) RegisterOwnershipTransferred(handler func(vLog *types.Log, data *L2ERC1155GatewayOwnershipTransferredEvent) error) {
	o.parsers[o.ABI.Events["OwnershipTransferred"].ID] = func(log *types.Log) error {
		event := new(L2ERC1155GatewayOwnershipTransferredEvent)
		if err := o.L2ERC1155GatewayCaller.contract.UnpackLog(event, "OwnershipTransferred", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// RegisterUpdateTokenMapping, the UpdateTokenMapping event ID is 0x2069a26c43c36ffaabe0c2d19bf65e55dd03abecdc449f5cc9663491e97f709d.
func (o *L2ERC1155Gateway) RegisterUpdateTokenMapping(handler func(vLog *types.Log, data *L2ERC1155GatewayUpdateTokenMappingEvent) error) {
	o.parsers[o.ABI.Events["UpdateTokenMapping"].ID] = func(log *types.Log) error {
		event := new(L2ERC1155GatewayUpdateTokenMappingEvent)
		if err := o.L2ERC1155GatewayCaller.contract.UnpackLog(event, "UpdateTokenMapping", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// RegisterWithdrawERC1155, the WithdrawERC1155 event ID is 0x1f9dcda7fce6f73a13055f044ffecaed2032a7a844e0a37a3eb8bbb17488d01a.
func (o *L2ERC1155Gateway) RegisterWithdrawERC1155(handler func(vLog *types.Log, data *L2ERC1155GatewayWithdrawERC1155Event) error) {
	o.parsers[o.ABI.Events["WithdrawERC1155"].ID] = func(log *types.Log) error {
		event := new(L2ERC1155GatewayWithdrawERC1155Event)
		if err := o.L2ERC1155GatewayCaller.contract.UnpackLog(event, "WithdrawERC1155", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// L2ERC1155GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ERC1155GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ERC1155GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ERC1155GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewL2ERC1155Gateway creates a new instance of L2ERC1155Gateway, bound to a specific deployed contract.
func NewL2ERC1155Gateway(address common.Address, backend bind.ContractBackend) (*L2ERC1155Gateway, error) {
	contract, err := bindL2ERC1155Gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	sigAbi, err := L2ERC1155GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	topics := make(map[common.Hash]string, len(sigAbi.Events))
	for name, _abi := range sigAbi.Events {
		topics[_abi.ID] = name
	}

	return &L2ERC1155Gateway{
		Name:                       "L2ERC1155Gateway",
		ABI:                        sigAbi,
		Address:                    address,
		topics:                     topics,
		parsers:                    map[common.Hash]func(log *types.Log) error{},
		L2ERC1155GatewayCaller:     L2ERC1155GatewayCaller{contract: contract},
		L2ERC1155GatewayTransactor: L2ERC1155GatewayTransactor{contract: contract},
	}, nil
}

// NewL2ERC1155GatewayCaller creates a new read-only instance of L2ERC1155Gateway, bound to a specific deployed contract.
func NewL2ERC1155GatewayCaller(address common.Address, caller bind.ContractCaller) (*L2ERC1155GatewayCaller, error) {
	contract, err := bindL2ERC1155Gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ERC1155GatewayCaller{contract: contract}, nil
}

// NewL2ERC1155GatewayTransactor creates a new write-only instance of L2ERC1155Gateway, bound to a specific deployed contract.
func NewL2ERC1155GatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L2ERC1155GatewayTransactor, error) {
	contract, err := bindL2ERC1155Gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ERC1155GatewayTransactor{contract: contract}, nil
}

// bindL2ERC1155Gateway binds a generic wrapper to an already deployed contract.
func bindL2ERC1155Gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2ERC1155GatewayMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewayCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC1155Gateway.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewayCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC1155Gateway.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC1155Gateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ERC1155Gateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2ERC1155Gateway *L2ERC1155GatewayCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _L2ERC1155Gateway.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address)
func (_L2ERC1155Gateway *L2ERC1155GatewayCaller) TokenMapping(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2ERC1155Gateway.contract.Call(opts, &out, "tokenMapping", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BatchWithdrawERC1155 is a paid mutator transaction binding the contract method 0x48de03de.
//
// Solidity: function batchWithdrawERC1155(address _token, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) BatchWithdrawERC1155(opts *bind.TransactOpts, _token common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "batchWithdrawERC1155", _token, _tokenIds, _amounts, _gasLimit)
}

// BatchWithdrawERC11550 is a paid mutator transaction binding the contract method 0x8c23d5b2.
//
// Solidity: function batchWithdrawERC1155(address _token, address _to, uint256[] _tokenIds, uint256[] _amounts, uint256 _gasLimit) payable returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) BatchWithdrawERC11550(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "batchWithdrawERC11550", _token, _to, _tokenIds, _amounts, _gasLimit)
}

// FinalizeBatchDepositERC1155 is a paid mutator transaction binding the contract method 0xeaa72ad9.
//
// Solidity: function finalizeBatchDepositERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256[] _tokenIds, uint256[] _amounts) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) FinalizeBatchDepositERC1155(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "finalizeBatchDepositERC1155", _l1Token, _l2Token, _from, _to, _tokenIds, _amounts)
}

// FinalizeDepositERC1155 is a paid mutator transaction binding the contract method 0x4764cc62.
//
// Solidity: function finalizeDepositERC1155(address _l1Token, address _l2Token, address _from, address _to, uint256 _tokenId, uint256 _amount) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) FinalizeDepositERC1155(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "finalizeDepositERC1155", _l1Token, _l2Token, _from, _to, _tokenId, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpart, address _messenger) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _messenger common.Address) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "initialize", _counterpart, _messenger)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "renounceOwnership")
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// UpdateTokenMapping is a paid mutator transaction binding the contract method 0xfac752eb.
//
// Solidity: function updateTokenMapping(address _l2Token, address _l1Token) returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) UpdateTokenMapping(opts *bind.TransactOpts, _l2Token common.Address, _l1Token common.Address) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "updateTokenMapping", _l2Token, _l1Token)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0x0f2da080.
//
// Solidity: function withdrawERC1155(address _token, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) WithdrawERC1155(opts *bind.TransactOpts, _token common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "withdrawERC1155", _token, _tokenId, _amount, _gasLimit)
}

// WithdrawERC11550 is a paid mutator transaction binding the contract method 0x21fedfc9.
//
// Solidity: function withdrawERC1155(address _token, address _to, uint256 _tokenId, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2ERC1155Gateway *L2ERC1155GatewayTransactor) WithdrawERC11550(opts *bind.TransactOpts, _token common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2ERC1155Gateway.contract.Transact(opts, "withdrawERC11550", _token, _to, _tokenId, _amount, _gasLimit)
}

// L2ERC1155GatewayBatchWithdrawERC1155 represents a BatchWithdrawERC1155 event raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayBatchWithdrawERC1155Event struct {
	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Amounts  []*big.Int
}

// L2ERC1155GatewayFinalizeBatchDepositERC1155 represents a FinalizeBatchDepositERC1155 event raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayFinalizeBatchDepositERC1155Event struct {
	L1Token  common.Address
	L2Token  common.Address
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Amounts  []*big.Int
}

// L2ERC1155GatewayFinalizeDepositERC1155 represents a FinalizeDepositERC1155 event raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayFinalizeDepositERC1155Event struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Amount  *big.Int
}

// L2ERC1155GatewayInitialized represents a Initialized event raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayInitializedEvent struct {
	Version uint8
}

// L2ERC1155GatewayOwnershipTransferred represents a OwnershipTransferred event raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayOwnershipTransferredEvent struct {
	PreviousOwner common.Address
	NewOwner      common.Address
}

// L2ERC1155GatewayUpdateTokenMapping represents a UpdateTokenMapping event raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayUpdateTokenMappingEvent struct {
	L2Token    common.Address
	OldL1Token common.Address
	NewL1Token common.Address
}

// L2ERC1155GatewayWithdrawERC1155 represents a WithdrawERC1155 event raised by the L2ERC1155Gateway contract.
type L2ERC1155GatewayWithdrawERC1155Event struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Amount  *big.Int
}
