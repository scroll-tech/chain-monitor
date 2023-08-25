// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// IERC1155MetaData contains all meta data concerning the IERC1155 contract.
var (
	IERC1155MetaData = &bind.MetaData{
		ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false}],\"type\":\"event\",\"name\":\"ApprovalForAll\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\",\"indexed\":false},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\",\"indexed\":false}],\"type\":\"event\",\"name\":\"TransferBatch\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"TransferSingle\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true}],\"type\":\"event\",\"name\":\"URI\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"safeBatchTransferFrom\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"safeTransferFrom\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setApprovalForAll\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}]}]",
	}
)

// IERC1155 is an auto generated Go binding around an Ethereum contract.
type IERC1155 struct {
	Name    string
	Address common.Address // contract address
	ABI     *abi.ABI       // IERC1155ABI is the input ABI used to generate the binding from.

	topics  map[common.Hash]string
	parsers map[common.Hash]func(log *types.Log) error

	IERC1155Caller     // Read-only binding to the contract
	IERC1155Transactor // Write-only binding to the contract
}

// GetName return IERC1155's Name.
func (o *IERC1155) GetName() string {
	return o.Name
}

// GetAddress return IERC1155's contract address.
func (o *IERC1155) GetAddress() common.Address {
	return o.Address
}

// GetEventName get event name by event hash.
func (o *IERC1155) GetEventName(sigHash common.Hash) string {
	return o.topics[sigHash]
}

// GetSigHashes get sig hash list.
func (o *IERC1155) GetSigHashes() []common.Hash {
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
func (o *IERC1155) ParseLog(vLog *types.Log) error {
	_id := vLog.Topics[0]
	if parse, exist := o.parsers[_id]; exist {
		return parse(vLog)
	}
	return nil
}

// RegisterApprovalForAll, the ApprovalForAll event ID is 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
func (o *IERC1155) RegisterApprovalForAll(handler func(vLog *types.Log, data *IERC1155ApprovalForAllEvent) error) {
	o.parsers[o.ABI.Events["ApprovalForAll"].ID] = func(log *types.Log) error {
		event := new(IERC1155ApprovalForAllEvent)
		if err := o.IERC1155Caller.contract.UnpackLog(event, "ApprovalForAll", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// RegisterTransferBatch, the TransferBatch event ID is 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
func (o *IERC1155) RegisterTransferBatch(handler func(vLog *types.Log, data *IERC1155TransferBatchEvent) error) {
	o.parsers[o.ABI.Events["TransferBatch"].ID] = func(log *types.Log) error {
		event := new(IERC1155TransferBatchEvent)
		if err := o.IERC1155Caller.contract.UnpackLog(event, "TransferBatch", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// RegisterTransferSingle, the TransferSingle event ID is 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
func (o *IERC1155) RegisterTransferSingle(handler func(vLog *types.Log, data *IERC1155TransferSingleEvent) error) {
	o.parsers[o.ABI.Events["TransferSingle"].ID] = func(log *types.Log) error {
		event := new(IERC1155TransferSingleEvent)
		if err := o.IERC1155Caller.contract.UnpackLog(event, "TransferSingle", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// RegisterURI, the URI event ID is 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
func (o *IERC1155) RegisterURI(handler func(vLog *types.Log, data *IERC1155URIEvent) error) {
	o.parsers[o.ABI.Events["URI"].ID] = func(log *types.Log) error {
		event := new(IERC1155URIEvent)
		if err := o.IERC1155Caller.contract.UnpackLog(event, "URI", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
}

// IERC1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewIERC1155 creates a new instance of IERC1155, bound to a specific deployed contract.
func NewIERC1155(address common.Address, backend bind.ContractBackend) (*IERC1155, error) {
	contract, err := bindIERC1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	sigAbi, err := IERC1155MetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	topics := make(map[common.Hash]string, len(sigAbi.Events))
	for name, _abi := range sigAbi.Events {
		topics[_abi.ID] = name
	}

	return &IERC1155{
		Name:               "IERC1155",
		ABI:                sigAbi,
		Address:            address,
		topics:             topics,
		parsers:            map[common.Hash]func(log *types.Log) error{},
		IERC1155Caller:     IERC1155Caller{contract: contract},
		IERC1155Transactor: IERC1155Transactor{contract: contract},
	}, nil
}

// NewIERC1155Caller creates a new read-only instance of IERC1155, bound to a specific deployed contract.
func NewIERC1155Caller(address common.Address, caller bind.ContractCaller) (*IERC1155Caller, error) {
	contract, err := bindIERC1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155Caller{contract: contract}, nil
}

// NewIERC1155Transactor creates a new write-only instance of IERC1155, bound to a specific deployed contract.
func NewIERC1155Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC1155Transactor, error) {
	contract, err := bindIERC1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155Transactor{contract: contract}, nil
}

// bindIERC1155 binds a generic wrapper to an already deployed contract.
func bindIERC1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC1155MetaData.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_IERC1155 *IERC1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_IERC1155 *IERC1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IERC1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_IERC1155 *IERC1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1155 *IERC1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_IERC1155 *IERC1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_IERC1155 *IERC1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IERC1155 *IERC1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _IERC1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// IERC1155ApprovalForAll represents a ApprovalForAll event raised by the IERC1155 contract.
type IERC1155ApprovalForAllEvent struct {
	Account  common.Address
	Operator common.Address
	Approved bool
}

// IERC1155TransferBatch represents a TransferBatch event raised by the IERC1155 contract.
type IERC1155TransferBatchEvent struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
}

// IERC1155TransferSingle represents a TransferSingle event raised by the IERC1155 contract.
type IERC1155TransferSingleEvent struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
}

// IERC1155URI represents a URI event raised by the IERC1155 contract.
type IERC1155URIEvent struct {
	Value string
	Id    *big.Int
}
