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

// IScrollERC1155MetaData contains all meta data concerning the IScrollERC1155 contract.
var (
	IScrollERC1155MetaData = &bind.MetaData{
		ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false}],\"type\":\"event\",\"name\":\"ApprovalForAll\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\",\"indexed\":false},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\",\"indexed\":false}],\"type\":\"event\",\"name\":\"TransferBatch\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"TransferSingle\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true}],\"type\":\"event\",\"name\":\"URI\",\"anonymous\":false},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}]},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"batchBurn\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"batchMint\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"burn\"},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}]},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"mint\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"safeBatchTransferFrom\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"safeTransferFrom\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setApprovalForAll\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}]}]",
	}
	// IScrollERC1155ABI is the input ABI used to generate the binding from.
	IScrollERC1155ABI, _ = IScrollERC1155MetaData.GetAbi()
)

// IScrollERC1155 is an auto generated Go binding around an Ethereum contract.
type IScrollERC1155 struct {
	Name    string
	Address common.Address // contract address
	ABI     *abi.ABI       // IScrollERC1155ABI is the input ABI used to generate the binding from.

	topics  map[common.Hash]string
	parsers map[common.Hash]func(log *types.Log) error

	IScrollERC1155Caller     // Read-only binding to the contract
	IScrollERC1155Transactor // Write-only binding to the contract
}

// GetName return IScrollERC1155's Name.
func (o *IScrollERC1155) GetName() string {
	return o.Name
}

// GetAddress return IScrollERC1155's contract address.
func (o *IScrollERC1155) GetAddress() common.Address {
	return o.Address
}

// GetEventName get event name by event hash.
func (o *IScrollERC1155) GetEventName(sigHash common.Hash) string {
	return o.topics[sigHash]
}

// GetSigHashes get sig hash list.
func (o *IScrollERC1155) GetSigHashes() []common.Hash {
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
func (o *IScrollERC1155) ParseLog(vLog *types.Log) (bool, error) {
	_id := vLog.Topics[0]
	if parse, exist := o.parsers[_id]; exist {
		return true, parse(vLog)
	}
	return false, nil
}

// RegisterApprovalForAll, the ApprovalForAll event ID is 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
func (o *IScrollERC1155) RegisterApprovalForAll(handler func(vLog *types.Log, data *IScrollERC1155ApprovalForAllEvent) error) {
	_id := o.ABI.Events["ApprovalForAll"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(IScrollERC1155ApprovalForAllEvent)
		if err := o.IScrollERC1155Caller.contract.UnpackLog(event, "ApprovalForAll", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
	o.topics[_id] = "ApprovalForAll"
}

// RegisterTransferBatch, the TransferBatch event ID is 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
func (o *IScrollERC1155) RegisterTransferBatch(handler func(vLog *types.Log, data *IScrollERC1155TransferBatchEvent) error) {
	_id := o.ABI.Events["TransferBatch"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(IScrollERC1155TransferBatchEvent)
		if err := o.IScrollERC1155Caller.contract.UnpackLog(event, "TransferBatch", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
	o.topics[_id] = "TransferBatch"
}

// RegisterTransferSingle, the TransferSingle event ID is 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
func (o *IScrollERC1155) RegisterTransferSingle(handler func(vLog *types.Log, data *IScrollERC1155TransferSingleEvent) error) {
	_id := o.ABI.Events["TransferSingle"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(IScrollERC1155TransferSingleEvent)
		if err := o.IScrollERC1155Caller.contract.UnpackLog(event, "TransferSingle", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
	o.topics[_id] = "TransferSingle"
}

// RegisterURI, the URI event ID is 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
func (o *IScrollERC1155) RegisterURI(handler func(vLog *types.Log, data *IScrollERC1155URIEvent) error) {
	_id := o.ABI.Events["URI"].ID
	o.parsers[_id] = func(log *types.Log) error {
		event := new(IScrollERC1155URIEvent)
		if err := o.IScrollERC1155Caller.contract.UnpackLog(event, "URI", *log); err != nil {
			return err
		}
		return handler(log, event)
	}
	o.topics[_id] = "URI"
}

// IScrollERC1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type IScrollERC1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScrollERC1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IScrollERC1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewIScrollERC1155 creates a new instance of IScrollERC1155, bound to a specific deployed contract.
func NewIScrollERC1155(address common.Address, backend bind.ContractBackend) (*IScrollERC1155, error) {
	contract, err := bindIScrollERC1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	sigAbi, err := IScrollERC1155MetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	topics := make(map[common.Hash]string, len(sigAbi.Events))
	for name, _abi := range sigAbi.Events {
		topics[_abi.ID] = name
	}

	return &IScrollERC1155{
		Name:                     "IScrollERC1155",
		ABI:                      sigAbi,
		Address:                  address,
		topics:                   topics,
		parsers:                  map[common.Hash]func(log *types.Log) error{},
		IScrollERC1155Caller:     IScrollERC1155Caller{contract: contract},
		IScrollERC1155Transactor: IScrollERC1155Transactor{contract: contract},
	}, nil
}

// NewIScrollERC1155Caller creates a new read-only instance of IScrollERC1155, bound to a specific deployed contract.
func NewIScrollERC1155Caller(address common.Address, caller bind.ContractCaller) (*IScrollERC1155Caller, error) {
	contract, err := bindIScrollERC1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IScrollERC1155Caller{contract: contract}, nil
}

// NewIScrollERC1155Transactor creates a new write-only instance of IScrollERC1155, bound to a specific deployed contract.
func NewIScrollERC1155Transactor(address common.Address, transactor bind.ContractTransactor) (*IScrollERC1155Transactor, error) {
	contract, err := bindIScrollERC1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IScrollERC1155Transactor{contract: contract}, nil
}

// bindIScrollERC1155 binds a generic wrapper to an already deployed contract.
func bindIScrollERC1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IScrollERC1155MetaData.ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_IScrollERC1155 *IScrollERC1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IScrollERC1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_IScrollERC1155 *IScrollERC1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IScrollERC1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_IScrollERC1155 *IScrollERC1155Caller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IScrollERC1155.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_IScrollERC1155 *IScrollERC1155Caller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IScrollERC1155.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_IScrollERC1155 *IScrollERC1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IScrollERC1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IScrollERC1155 *IScrollERC1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IScrollERC1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BatchBurn is a paid mutator transaction binding the contract method 0xf6eb127a.
//
// Solidity: function batchBurn(address _from, uint256[] _tokenIds, uint256[] _amounts) returns()
func (_IScrollERC1155 *IScrollERC1155Transactor) BatchBurn(opts *bind.TransactOpts, _from common.Address, _tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _IScrollERC1155.contract.Transact(opts, "batchBurn", _from, _tokenIds, _amounts)
}

// BatchMint is a paid mutator transaction binding the contract method 0xb48ab8b6.
//
// Solidity: function batchMint(address _to, uint256[] _tokenIds, uint256[] _amounts, bytes _data) returns()
func (_IScrollERC1155 *IScrollERC1155Transactor) BatchMint(opts *bind.TransactOpts, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _IScrollERC1155.contract.Transact(opts, "batchMint", _to, _tokenIds, _amounts, _data)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address _from, uint256 _tokenId, uint256 _amount) returns()
func (_IScrollERC1155 *IScrollERC1155Transactor) Burn(opts *bind.TransactOpts, _from common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _IScrollERC1155.contract.Transact(opts, "burn", _from, _tokenId, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address _to, uint256 _tokenId, uint256 _amount, bytes _data) returns()
func (_IScrollERC1155 *IScrollERC1155Transactor) Mint(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IScrollERC1155.contract.Transact(opts, "mint", _to, _tokenId, _amount, _data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_IScrollERC1155 *IScrollERC1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _IScrollERC1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_IScrollERC1155 *IScrollERC1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IScrollERC1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IScrollERC1155 *IScrollERC1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _IScrollERC1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// IScrollERC1155ApprovalForAll represents a ApprovalForAll event raised by the IScrollERC1155 contract.
type IScrollERC1155ApprovalForAllEvent struct {
	Account  common.Address
	Operator common.Address
	Approved bool
}

// IScrollERC1155TransferBatch represents a TransferBatch event raised by the IScrollERC1155 contract.
type IScrollERC1155TransferBatchEvent struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
}

// IScrollERC1155TransferSingle represents a TransferSingle event raised by the IScrollERC1155 contract.
type IScrollERC1155TransferSingleEvent struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
}

// IScrollERC1155URI represents a URI event raised by the IScrollERC1155 contract.
type IScrollERC1155URIEvent struct {
	Value string
	Id    *big.Int
}
