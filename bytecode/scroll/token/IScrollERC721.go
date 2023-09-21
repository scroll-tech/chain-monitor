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

// IScrollERC721MetaData contains all meta data concerning the IScrollERC721 contract.
var (
	IScrollERC721MetaData = &bind.MetaData{
		ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true}],\"type\":\"event\",\"name\":\"Approval\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false}],\"type\":\"event\",\"name\":\"ApprovalForAll\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true}],\"type\":\"event\",\"name\":\"Transfer\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}]},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"burn\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"mint\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"safeTransferFrom\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"safeTransferFrom\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setApprovalForAll\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\"}]",
	}
	// IScrollERC721ABI is the input ABI used to generate the binding from.
	IScrollERC721ABI, _ = IScrollERC721MetaData.GetAbi()
)

// IScrollERC721 is an auto generated Go binding around an Ethereum contract.
type IScrollERC721 struct {
	Name    string
	Address common.Address // contract address
	ABI     *abi.ABI       // IScrollERC721ABI is the input ABI used to generate the binding from.

	topics  map[common.Hash]string
	parsers map[common.Hash]func(log *types.Log) error

	IScrollERC721Caller     // Read-only binding to the contract
	IScrollERC721Transactor // Write-only binding to the contract
}

// GetName return IScrollERC721's Name.
func (o *IScrollERC721) GetName() string {
	return o.Name
}

// GetAddress return IScrollERC721's contract address.
func (o *IScrollERC721) GetAddress() common.Address {
	return o.Address
}

// GetEventName get event name by event hash.
func (o *IScrollERC721) GetEventName(sigHash common.Hash) string {
	return o.topics[sigHash]
}

// GetSigHashes get sig hash list.
func (o *IScrollERC721) GetSigHashes() []common.Hash {
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
func (o *IScrollERC721) ParseLog(vLog *types.Log) (bool, error) {
	_id := vLog.Topics[0]
	if parse, exist := o.parsers[_id]; exist {
		return true, parse(vLog)
	}
	return false, nil
}

// RegisterApproval, the Approval event ID is 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
func (o *IScrollERC721) RegisterApproval(handler func(vLog *types.Log, data *IScrollERC721ApprovalEvent) error) {
	_id := o.ABI.Events["Approval"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(IScrollERC721ApprovalEvent)
		if err := o.IScrollERC721Caller.contract.UnpackLog(event, "Approval", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
	o.topics[_id] = "Approval"
}

// RegisterApprovalForAll, the ApprovalForAll event ID is 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
func (o *IScrollERC721) RegisterApprovalForAll(handler func(vLog *types.Log, data *IScrollERC721ApprovalForAllEvent) error) {
	_id := o.ABI.Events["ApprovalForAll"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(IScrollERC721ApprovalForAllEvent)
		if err := o.IScrollERC721Caller.contract.UnpackLog(event, "ApprovalForAll", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
	o.topics[_id] = "ApprovalForAll"
}

// RegisterTransfer, the Transfer event ID is 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
func (o *IScrollERC721) RegisterTransfer(handler func(vLog *types.Log, data *IScrollERC721TransferEvent) error) {
	_id := o.ABI.Events["Transfer"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(IScrollERC721TransferEvent)
		if err := o.IScrollERC721Caller.contract.UnpackLog(event, "Transfer", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
	o.topics[_id] = "Transfer"
}

// IScrollERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type IScrollERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScrollERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IScrollERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewIScrollERC721 creates a new instance of IScrollERC721, bound to a specific deployed contract.
func NewIScrollERC721(address common.Address, backend bind.ContractBackend) (*IScrollERC721, error) {
	contract, err := bindIScrollERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	sigAbi, err := IScrollERC721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	topics := make(map[common.Hash]string, len(sigAbi.Events))
	for name, _abi := range sigAbi.Events {
		topics[_abi.ID] = name
	}

	return &IScrollERC721{
		Name:                    "IScrollERC721",
		ABI:                     sigAbi,
		Address:                 address,
		topics:                  topics,
		parsers:                 map[common.Hash]func(log *types.Log) error{},
		IScrollERC721Caller:     IScrollERC721Caller{contract: contract},
		IScrollERC721Transactor: IScrollERC721Transactor{contract: contract},
	}, nil
}

// NewIScrollERC721Caller creates a new read-only instance of IScrollERC721, bound to a specific deployed contract.
func NewIScrollERC721Caller(address common.Address, caller bind.ContractCaller) (*IScrollERC721Caller, error) {
	contract, err := bindIScrollERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IScrollERC721Caller{contract: contract}, nil
}

// NewIScrollERC721Transactor creates a new write-only instance of IScrollERC721, bound to a specific deployed contract.
func NewIScrollERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*IScrollERC721Transactor, error) {
	contract, err := bindIScrollERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IScrollERC721Transactor{contract: contract}, nil
}

// bindIScrollERC721 binds a generic wrapper to an already deployed contract.
func bindIScrollERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IScrollERC721MetaData.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IScrollERC721 *IScrollERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IScrollERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_IScrollERC721 *IScrollERC721Caller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IScrollERC721.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_IScrollERC721 *IScrollERC721Caller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IScrollERC721.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IScrollERC721 *IScrollERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IScrollERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IScrollERC721 *IScrollERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IScrollERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IScrollERC721 *IScrollERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IScrollERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IScrollERC721 *IScrollERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IScrollERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IScrollERC721 *IScrollERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IScrollERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _tokenId) returns()
func (_IScrollERC721 *IScrollERC721Transactor) Burn(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _IScrollERC721.contract.Transact(opts, "burn", _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _tokenId) returns()
func (_IScrollERC721 *IScrollERC721Transactor) Mint(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _IScrollERC721.contract.Transact(opts, "mint", _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IScrollERC721 *IScrollERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IScrollERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IScrollERC721 *IScrollERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IScrollERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IScrollERC721 *IScrollERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _IScrollERC721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IScrollERC721 *IScrollERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IScrollERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// IScrollERC721Approval represents a Approval event raised by the IScrollERC721 contract.
type IScrollERC721ApprovalEvent struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
}

// IScrollERC721ApprovalForAll represents a ApprovalForAll event raised by the IScrollERC721 contract.
type IScrollERC721ApprovalForAllEvent struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
}

// IScrollERC721Transfer represents a Transfer event raised by the IScrollERC721 contract.
type IScrollERC721TransferEvent struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
}
