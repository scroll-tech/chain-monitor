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

// IERC721MetaData contains all meta data concerning the IERC721 contract.
var (
	IERC721MetaData = &bind.MetaData{
		ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true}],\"type\":\"event\",\"name\":\"Approval\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false}],\"type\":\"event\",\"name\":\"ApprovalForAll\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true}],\"type\":\"event\",\"name\":\"Transfer\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}]},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}]},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"safeTransferFrom\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"safeTransferFrom\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setApprovalForAll\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\"}]",
	}
	// IERC721ABI is the input ABI used to generate the binding from.
	IERC721ABI, _ = IERC721MetaData.GetAbi()
)

// IERC721 is an auto generated Go binding around an Ethereum contract.
type IERC721 struct {
	Name    string
	Address common.Address // contract address
	ABI     *abi.ABI       // IERC721ABI is the input ABI used to generate the binding from.

	topics  map[common.Hash]string
	parsers map[common.Hash]func(log *types.Log) error

	IERC721Caller     // Read-only binding to the contract
	IERC721Transactor // Write-only binding to the contract
}

// GetName return IERC721's Name.
func (o *IERC721) GetName() string {
	return o.Name
}

// GetAddress return IERC721's contract address.
func (o *IERC721) GetAddress() common.Address {
	return o.Address
}

// GetEventName get event name by event hash.
func (o *IERC721) GetEventName(sigHash common.Hash) string {
	return o.topics[sigHash]
}

// GetSigHashes get sig hash list.
func (o *IERC721) GetSigHashes() []common.Hash {
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
func (o *IERC721) ParseLog(vLog *types.Log) (bool, error) {
	_id := vLog.Topics[0]
	if parse, exist := o.parsers[_id]; exist {
		return true, parse(vLog)
	}
	return false, nil
}

// RegisterApproval, the Approval event ID is 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
func (o *IERC721) RegisterApproval(handler func(vLog *types.Log, data *IERC721ApprovalEvent) error) {
	_id := o.ABI.Events["Approval"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(IERC721ApprovalEvent)
		if err := o.IERC721Caller.contract.UnpackLog(event, "Approval", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "Approval"
}

// RegisterApprovalForAll, the ApprovalForAll event ID is 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
func (o *IERC721) RegisterApprovalForAll(handler func(vLog *types.Log, data *IERC721ApprovalForAllEvent) error) {
	_id := o.ABI.Events["ApprovalForAll"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(IERC721ApprovalForAllEvent)
		if err := o.IERC721Caller.contract.UnpackLog(event, "ApprovalForAll", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "ApprovalForAll"
}

// RegisterTransfer, the Transfer event ID is 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
func (o *IERC721) RegisterTransfer(handler func(vLog *types.Log, data *IERC721TransferEvent) error) {
	_id := o.ABI.Events["Transfer"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(IERC721TransferEvent)
		if err := o.IERC721Caller.contract.UnpackLog(event, "Transfer", *log); err != nil {
			return err
		}
		event.Log = log
		return handler(log, event)
	}
	o.topics[_id] = "Transfer"
}

// IERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewIERC721 creates a new instance of IERC721, bound to a specific deployed contract.
func NewIERC721(address common.Address, backend bind.ContractBackend) (*IERC721, error) {
	contract, err := bindIERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	sigAbi, err := IERC721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	topics := make(map[common.Hash]string, len(sigAbi.Events))
	for name, _abi := range sigAbi.Events {
		topics[_abi.ID] = name
	}

	return &IERC721{
		Name:              "IERC721",
		ABI:               sigAbi,
		Address:           address,
		topics:            topics,
		parsers:           map[common.Hash]func(log *types.Log) error{},
		IERC721Caller:     IERC721Caller{contract: contract},
		IERC721Transactor: IERC721Transactor{contract: contract},
	}, nil
}

// NewIERC721Caller creates a new read-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Caller(address common.Address, caller bind.ContractCaller) (*IERC721Caller, error) {
	contract, err := bindIERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Caller{contract: contract}, nil
}

// NewIERC721Transactor creates a new write-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC721Transactor, error) {
	contract, err := bindIERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Transactor{contract: contract}, nil
}

// bindIERC721 binds a generic wrapper to an already deployed contract.
func bindIERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721MetaData.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721 *IERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721 *IERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721 *IERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721 *IERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721 *IERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "approve", to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IERC721 *IERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// IERC721Approval represents a Approval event raised by the IERC721 contract.
type IERC721ApprovalEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
}

// IERC721ApprovalForAll represents a ApprovalForAll event raised by the IERC721 contract.
type IERC721ApprovalForAllEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	Owner    common.Address
	Operator common.Address
	Approved bool
}

// IERC721Transfer represents a Transfer event raised by the IERC721 contract.
type IERC721TransferEvent struct {
	Log *types.Log `json:"-" gorm:"-"`

	From    common.Address
	To      common.Address
	TokenId *big.Int
}
