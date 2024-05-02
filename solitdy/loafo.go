// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package solitdy

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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

// LostItemsFoundApplication is an auto generated low-level Go binding around an user-defined struct.
type LostItemsFoundApplication struct {
	ApplicationId *big.Int
	RelateItemID  *big.Int
	Mysender      common.Address
	Receiver      common.Address
	SubmitDate    *big.Int
	Result        bool
	IsConfirmed   bool
	RewardPaid    bool
}

// LostItemsFoundOneChange is an auto generated low-level Go binding around an user-defined struct.
type LostItemsFoundOneChange struct {
	FoundOnes []common.Address
	Total     *big.Int
	TraceHash [32]byte
}

// LostItemsHistory is an auto generated low-level Go binding around an user-defined struct.
type LostItemsHistory struct {
	HistoryID         *big.Int
	Operator          common.Address
	Operation         string
	OperationTime     *big.Int
	RelatedLostItemID *big.Int
}

// LoAfoMetaData contains all meta data concerning the LoAfo contract.
var LoAfoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_LostItemID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_Operation\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_OperationTime\",\"type\":\"uint256\"}],\"name\":\"addHistoryRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_lostItemName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_lostDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_itemCategory\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"}],\"name\":\"addLostItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lostItemID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_applicationId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"}],\"name\":\"confirmFoundApplication\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"traceHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_points\",\"type\":\"uint256\"}],\"name\":\"decereaseUserPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lostItemID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_applicationId\",\"type\":\"uint256\"}],\"name\":\"deleteApplication\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lostItemID\",\"type\":\"uint256\"}],\"name\":\"deleteLostItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_points\",\"type\":\"uint256\"}],\"name\":\"increaseUserPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lostItemID\",\"type\":\"uint256\"}],\"name\":\"submitFoundApplication\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lostItemID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"lostItemName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"LostItemAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendToCaller\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdminAddress\",\"type\":\"address\"}],\"name\":\"setAdminAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lostItemID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_lostItemName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_lostDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_itemCategory\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"}],\"name\":\"updateLostItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lostItemID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_applicationId\",\"type\":\"uint256\"}],\"name\":\"applicationExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lostItemID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_applicationId\",\"type\":\"uint256\"}],\"name\":\"ApplicationInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"applicationId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"relateItemID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mysender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"submitDate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isConfirmed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rewardPaid\",\"type\":\"bool\"}],\"internalType\":\"structLostItems.FoundApplication\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"foundOneChanges\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"traceHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdminAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lostItemID\",\"type\":\"uint256\"}],\"name\":\"getFoundOneChange\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"foundOnes\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"traceHash\",\"type\":\"bytes32\"}],\"internalType\":\"structLostItems.FoundOneChange\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lostItemID\",\"type\":\"uint256\"}],\"name\":\"getFoundOneChangeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getHistoryRecordsByOperator\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"HistoryID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"Operator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"Operation\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"OperationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"RelatedLostItemID\",\"type\":\"uint256\"}],\"internalType\":\"structLostItems.History[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lostItemID\",\"type\":\"uint256\"}],\"name\":\"getLostItemOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_traceHash\",\"type\":\"bytes32\"}],\"name\":\"getTraceInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lostItemID\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"foundOnes\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"traceHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"histories\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"HistoryID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"Operator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"Operation\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"OperationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"RelatedLostItemID\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lostItemID\",\"type\":\"uint256\"}],\"name\":\"lostItemExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lostItems\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"LostItemID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"User\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"FoundOne\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"LostItemName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"LostDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"Location\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ItemCategory\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"IsFound\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"FoundDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"FoundApplicationsCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextApplicationId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextHistoryID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextLostItemID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TotalHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TotalLostItems\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LoAfoABI is the input ABI used to generate the binding from.
// Deprecated: Use LoAfoMetaData.ABI instead.
var LoAfoABI = LoAfoMetaData.ABI

// LoAfo is an auto generated Go binding around an Ethereum contract.
type LoAfo struct {
	LoAfoCaller     // Read-only binding to the contract
	LoAfoTransactor // Write-only binding to the contract
	LoAfoFilterer   // Log filterer for contract events
}

// LoAfoCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoAfoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoAfoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoAfoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoAfoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoAfoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoAfoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoAfoSession struct {
	Contract     *LoAfo            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoAfoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoAfoCallerSession struct {
	Contract *LoAfoCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LoAfoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoAfoTransactorSession struct {
	Contract     *LoAfoTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoAfoRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoAfoRaw struct {
	Contract *LoAfo // Generic contract binding to access the raw methods on
}

// LoAfoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoAfoCallerRaw struct {
	Contract *LoAfoCaller // Generic read-only contract binding to access the raw methods on
}

// LoAfoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoAfoTransactorRaw struct {
	Contract *LoAfoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLoAfo creates a new instance of LoAfo, bound to a specific deployed contract.
func NewLoAfo(address common.Address, backend bind.ContractBackend) (*LoAfo, error) {
	contract, err := bindLoAfo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LoAfo{LoAfoCaller: LoAfoCaller{contract: contract}, LoAfoTransactor: LoAfoTransactor{contract: contract}, LoAfoFilterer: LoAfoFilterer{contract: contract}}, nil
}

// NewLoAfoCaller creates a new read-only instance of LoAfo, bound to a specific deployed contract.
func NewLoAfoCaller(address common.Address, caller bind.ContractCaller) (*LoAfoCaller, error) {
	contract, err := bindLoAfo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoAfoCaller{contract: contract}, nil
}

// NewLoAfoTransactor creates a new write-only instance of LoAfo, bound to a specific deployed contract.
func NewLoAfoTransactor(address common.Address, transactor bind.ContractTransactor) (*LoAfoTransactor, error) {
	contract, err := bindLoAfo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoAfoTransactor{contract: contract}, nil
}

// NewLoAfoFilterer creates a new log filterer instance of LoAfo, bound to a specific deployed contract.
func NewLoAfoFilterer(address common.Address, filterer bind.ContractFilterer) (*LoAfoFilterer, error) {
	contract, err := bindLoAfo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoAfoFilterer{contract: contract}, nil
}

// bindLoAfo binds a generic wrapper to an already deployed contract.
func bindLoAfo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LoAfoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoAfo *LoAfoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LoAfo.Contract.LoAfoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoAfo *LoAfoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoAfo.Contract.LoAfoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoAfo *LoAfoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoAfo.Contract.LoAfoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoAfo *LoAfoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LoAfo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoAfo *LoAfoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoAfo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoAfo *LoAfoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoAfo.Contract.contract.Transact(opts, method, params...)
}

// ApplicationInfo is a free data retrieval call binding the contract method 0xf7e7171d.
//
// Solidity: function ApplicationInfo(uint256 _lostItemID, uint256 _applicationId) view returns((uint256,uint256,address,address,uint256,bool,bool,bool))
func (_LoAfo *LoAfoCaller) ApplicationInfo(opts *bind.CallOpts, _lostItemID *big.Int, _applicationId *big.Int) (LostItemsFoundApplication, error) {
	var out []interface{}
	err := _LoAfo.contract.Call(opts, &out, "ApplicationInfo", _lostItemID, _applicationId)

	if err != nil {
		return *new(LostItemsFoundApplication), err
	}

	out0 := *abi.ConvertType(out[0], new(LostItemsFoundApplication)).(*LostItemsFoundApplication)

	return out0, err

}

// ApplicationInfo is a free data retrieval call binding the contract method 0xf7e7171d.
//
// Solidity: function ApplicationInfo(uint256 _lostItemID, uint256 _applicationId) view returns((uint256,uint256,address,address,uint256,bool,bool,bool))
func (_LoAfo *LoAfoSession) ApplicationInfo(_lostItemID *big.Int, _applicationId *big.Int) (LostItemsFoundApplication, error) {
	return _LoAfo.Contract.ApplicationInfo(&_LoAfo.CallOpts, _lostItemID, _applicationId)
}

// ApplicationInfo is a free data retrieval call binding the contract method 0xf7e7171d.
//
// Solidity: function ApplicationInfo(uint256 _lostItemID, uint256 _applicationId) view returns((uint256,uint256,address,address,uint256,bool,bool,bool))
func (_LoAfo *LoAfoCallerSession) ApplicationInfo(_lostItemID *big.Int, _applicationId *big.Int) (LostItemsFoundApplication, error) {
	return _LoAfo.Contract.ApplicationInfo(&_LoAfo.CallOpts, _lostItemID, _applicationId)
}

// TotalHistory is a free data retrieval call binding the contract method 0x1102109c.
//
// Solidity: function TotalHistory() view returns(uint256)
func (_LoAfo *LoAfoCaller) TotalHistory(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LoAfo.contract.Call(opts, &out, "TotalHistory")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalHistory is a free data retrieval call binding the contract method 0x1102109c.
//
// Solidity: function TotalHistory() view returns(uint256)
func (_LoAfo *LoAfoSession) TotalHistory() (*big.Int, error) {
	return _LoAfo.Contract.TotalHistory(&_LoAfo.CallOpts)
}

// TotalHistory is a free data retrieval call binding the contract method 0x1102109c.
//
// Solidity: function TotalHistory() view returns(uint256)
func (_LoAfo *LoAfoCallerSession) TotalHistory() (*big.Int, error) {
	return _LoAfo.Contract.TotalHistory(&_LoAfo.CallOpts)
}

// TotalLostItems is a free data retrieval call binding the contract method 0xfd4ca838.
//
// Solidity: function TotalLostItems() view returns(uint256)
func (_LoAfo *LoAfoCaller) TotalLostItems(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LoAfo.contract.Call(opts, &out, "TotalLostItems")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLostItems is a free data retrieval call binding the contract method 0xfd4ca838.
//
// Solidity: function TotalLostItems() view returns(uint256)
func (_LoAfo *LoAfoSession) TotalLostItems() (*big.Int, error) {
	return _LoAfo.Contract.TotalLostItems(&_LoAfo.CallOpts)
}

// TotalLostItems is a free data retrieval call binding the contract method 0xfd4ca838.
//
// Solidity: function TotalLostItems() view returns(uint256)
func (_LoAfo *LoAfoCallerSession) TotalLostItems() (*big.Int, error) {
	return _LoAfo.Contract.TotalLostItems(&_LoAfo.CallOpts)
}

// ApplicationExists is a free data retrieval call binding the contract method 0x7b60eb55.
//
// Solidity: function applicationExists(uint256 _lostItemID, uint256 _applicationId) view returns(bool)
func (_LoAfo *LoAfoCaller) ApplicationExists(opts *bind.CallOpts, _lostItemID *big.Int, _applicationId *big.Int) (bool, error) {
	var out []interface{}
	err := _LoAfo.contract.Call(opts, &out, "applicationExists", _lostItemID, _applicationId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApplicationExists is a free data retrieval call binding the contract method 0x7b60eb55.
//
// Solidity: function applicationExists(uint256 _lostItemID, uint256 _applicationId) view returns(bool)
func (_LoAfo *LoAfoSession) ApplicationExists(_lostItemID *big.Int, _applicationId *big.Int) (bool, error) {
	return _LoAfo.Contract.ApplicationExists(&_LoAfo.CallOpts, _lostItemID, _applicationId)
}

// ApplicationExists is a free data retrieval call binding the contract method 0x7b60eb55.
//
// Solidity: function applicationExists(uint256 _lostItemID, uint256 _applicationId) view returns(bool)
func (_LoAfo *LoAfoCallerSession) ApplicationExists(_lostItemID *big.Int, _applicationId *big.Int) (bool, error) {
	return _LoAfo.Contract.ApplicationExists(&_LoAfo.CallOpts, _lostItemID, _applicationId)
}

// FoundOneChanges is a free data retrieval call binding the contract method 0x3d89b730.
//
// Solidity: function foundOneChanges(uint256 ) view returns(uint256 total, bytes32 traceHash)
func (_LoAfo *LoAfoCaller) FoundOneChanges(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Total     *big.Int
	TraceHash [32]byte
}, error) {
	var out []interface{}
	err := _LoAfo.contract.Call(opts, &out, "foundOneChanges", arg0)

	outstruct := new(struct {
		Total     *big.Int
		TraceHash [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Total = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TraceHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// FoundOneChanges is a free data retrieval call binding the contract method 0x3d89b730.
//
// Solidity: function foundOneChanges(uint256 ) view returns(uint256 total, bytes32 traceHash)
func (_LoAfo *LoAfoSession) FoundOneChanges(arg0 *big.Int) (struct {
	Total     *big.Int
	TraceHash [32]byte
}, error) {
	return _LoAfo.Contract.FoundOneChanges(&_LoAfo.CallOpts, arg0)
}

// FoundOneChanges is a free data retrieval call binding the contract method 0x3d89b730.
//
// Solidity: function foundOneChanges(uint256 ) view returns(uint256 total, bytes32 traceHash)
func (_LoAfo *LoAfoCallerSession) FoundOneChanges(arg0 *big.Int) (struct {
	Total     *big.Int
	TraceHash [32]byte
}, error) {
	return _LoAfo.Contract.FoundOneChanges(&_LoAfo.CallOpts, arg0)
}

// GetAdminAddress is a free data retrieval call binding the contract method 0xb2e6b912.
//
// Solidity: function getAdminAddress() view returns(address)
func (_LoAfo *LoAfoCaller) GetAdminAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LoAfo.contract.Call(opts, &out, "getAdminAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAdminAddress is a free data retrieval call binding the contract method 0xb2e6b912.
//
// Solidity: function getAdminAddress() view returns(address)
func (_LoAfo *LoAfoSession) GetAdminAddress() (common.Address, error) {
	return _LoAfo.Contract.GetAdminAddress(&_LoAfo.CallOpts)
}

// GetAdminAddress is a free data retrieval call binding the contract method 0xb2e6b912.
//
// Solidity: function getAdminAddress() view returns(address)
func (_LoAfo *LoAfoCallerSession) GetAdminAddress() (common.Address, error) {
	return _LoAfo.Contract.GetAdminAddress(&_LoAfo.CallOpts)
}

// GetFoundOneChange is a free data retrieval call binding the contract method 0x5a3090a1.
//
// Solidity: function getFoundOneChange(uint256 _lostItemID) view returns((address[],uint256,bytes32))
func (_LoAfo *LoAfoCaller) GetFoundOneChange(opts *bind.CallOpts, _lostItemID *big.Int) (LostItemsFoundOneChange, error) {
	var out []interface{}
	err := _LoAfo.contract.Call(opts, &out, "getFoundOneChange", _lostItemID)

	if err != nil {
		return *new(LostItemsFoundOneChange), err
	}

	out0 := *abi.ConvertType(out[0], new(LostItemsFoundOneChange)).(*LostItemsFoundOneChange)

	return out0, err

}

// GetFoundOneChange is a free data retrieval call binding the contract method 0x5a3090a1.
//
// Solidity: function getFoundOneChange(uint256 _lostItemID) view returns((address[],uint256,bytes32))
func (_LoAfo *LoAfoSession) GetFoundOneChange(_lostItemID *big.Int) (LostItemsFoundOneChange, error) {
	return _LoAfo.Contract.GetFoundOneChange(&_LoAfo.CallOpts, _lostItemID)
}

// GetFoundOneChange is a free data retrieval call binding the contract method 0x5a3090a1.
//
// Solidity: function getFoundOneChange(uint256 _lostItemID) view returns((address[],uint256,bytes32))
func (_LoAfo *LoAfoCallerSession) GetFoundOneChange(_lostItemID *big.Int) (LostItemsFoundOneChange, error) {
	return _LoAfo.Contract.GetFoundOneChange(&_LoAfo.CallOpts, _lostItemID)
}

// GetFoundOneChangeCount is a free data retrieval call binding the contract method 0x4748e999.
//
// Solidity: function getFoundOneChangeCount(uint256 _lostItemID) view returns(uint256)
func (_LoAfo *LoAfoCaller) GetFoundOneChangeCount(opts *bind.CallOpts, _lostItemID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LoAfo.contract.Call(opts, &out, "getFoundOneChangeCount", _lostItemID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFoundOneChangeCount is a free data retrieval call binding the contract method 0x4748e999.
//
// Solidity: function getFoundOneChangeCount(uint256 _lostItemID) view returns(uint256)
func (_LoAfo *LoAfoSession) GetFoundOneChangeCount(_lostItemID *big.Int) (*big.Int, error) {
	return _LoAfo.Contract.GetFoundOneChangeCount(&_LoAfo.CallOpts, _lostItemID)
}

// GetFoundOneChangeCount is a free data retrieval call binding the contract method 0x4748e999.
//
// Solidity: function getFoundOneChangeCount(uint256 _lostItemID) view returns(uint256)
func (_LoAfo *LoAfoCallerSession) GetFoundOneChangeCount(_lostItemID *big.Int) (*big.Int, error) {
	return _LoAfo.Contract.GetFoundOneChangeCount(&_LoAfo.CallOpts, _lostItemID)
}

// GetHistoryRecordsByOperator is a free data retrieval call binding the contract method 0x0aa6cdad.
//
// Solidity: function getHistoryRecordsByOperator(address _operator) view returns((uint256,address,string,uint256,uint256)[])
func (_LoAfo *LoAfoCaller) GetHistoryRecordsByOperator(opts *bind.CallOpts, _operator common.Address) ([]LostItemsHistory, error) {
	var out []interface{}
	err := _LoAfo.contract.Call(opts, &out, "getHistoryRecordsByOperator", _operator)

	if err != nil {
		return *new([]LostItemsHistory), err
	}

	out0 := *abi.ConvertType(out[0], new([]LostItemsHistory)).(*[]LostItemsHistory)

	return out0, err

}

// GetHistoryRecordsByOperator is a free data retrieval call binding the contract method 0x0aa6cdad.
//
// Solidity: function getHistoryRecordsByOperator(address _operator) view returns((uint256,address,string,uint256,uint256)[])
func (_LoAfo *LoAfoSession) GetHistoryRecordsByOperator(_operator common.Address) ([]LostItemsHistory, error) {
	return _LoAfo.Contract.GetHistoryRecordsByOperator(&_LoAfo.CallOpts, _operator)
}

// GetHistoryRecordsByOperator is a free data retrieval call binding the contract method 0x0aa6cdad.
//
// Solidity: function getHistoryRecordsByOperator(address _operator) view returns((uint256,address,string,uint256,uint256)[])
func (_LoAfo *LoAfoCallerSession) GetHistoryRecordsByOperator(_operator common.Address) ([]LostItemsHistory, error) {
	return _LoAfo.Contract.GetHistoryRecordsByOperator(&_LoAfo.CallOpts, _operator)
}

// GetLostItemOwner is a free data retrieval call binding the contract method 0x039705ee.
//
// Solidity: function getLostItemOwner(uint256 _lostItemID) view returns(address)
func (_LoAfo *LoAfoCaller) GetLostItemOwner(opts *bind.CallOpts, _lostItemID *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LoAfo.contract.Call(opts, &out, "getLostItemOwner", _lostItemID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLostItemOwner is a free data retrieval call binding the contract method 0x039705ee.
//
// Solidity: function getLostItemOwner(uint256 _lostItemID) view returns(address)
func (_LoAfo *LoAfoSession) GetLostItemOwner(_lostItemID *big.Int) (common.Address, error) {
	return _LoAfo.Contract.GetLostItemOwner(&_LoAfo.CallOpts, _lostItemID)
}

// GetLostItemOwner is a free data retrieval call binding the contract method 0x039705ee.
//
// Solidity: function getLostItemOwner(uint256 _lostItemID) view returns(address)
func (_LoAfo *LoAfoCallerSession) GetLostItemOwner(_lostItemID *big.Int) (common.Address, error) {
	return _LoAfo.Contract.GetLostItemOwner(&_LoAfo.CallOpts, _lostItemID)
}

// GetTraceInfo is a free data retrieval call binding the contract method 0xeebb8d52.
//
// Solidity: function getTraceInfo(bytes32 _traceHash) view returns(uint256 lostItemID, address[] foundOnes, uint256 total, bytes32 traceHash)
func (_LoAfo *LoAfoCaller) GetTraceInfo(opts *bind.CallOpts, _traceHash [32]byte) (struct {
	LostItemID *big.Int
	FoundOnes  []common.Address
	Total      *big.Int
	TraceHash  [32]byte
}, error) {
	var out []interface{}
	err := _LoAfo.contract.Call(opts, &out, "getTraceInfo", _traceHash)

	outstruct := new(struct {
		LostItemID *big.Int
		FoundOnes  []common.Address
		Total      *big.Int
		TraceHash  [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LostItemID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FoundOnes = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Total = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TraceHash = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetTraceInfo is a free data retrieval call binding the contract method 0xeebb8d52.
//
// Solidity: function getTraceInfo(bytes32 _traceHash) view returns(uint256 lostItemID, address[] foundOnes, uint256 total, bytes32 traceHash)
func (_LoAfo *LoAfoSession) GetTraceInfo(_traceHash [32]byte) (struct {
	LostItemID *big.Int
	FoundOnes  []common.Address
	Total      *big.Int
	TraceHash  [32]byte
}, error) {
	return _LoAfo.Contract.GetTraceInfo(&_LoAfo.CallOpts, _traceHash)
}

// GetTraceInfo is a free data retrieval call binding the contract method 0xeebb8d52.
//
// Solidity: function getTraceInfo(bytes32 _traceHash) view returns(uint256 lostItemID, address[] foundOnes, uint256 total, bytes32 traceHash)
func (_LoAfo *LoAfoCallerSession) GetTraceInfo(_traceHash [32]byte) (struct {
	LostItemID *big.Int
	FoundOnes  []common.Address
	Total      *big.Int
	TraceHash  [32]byte
}, error) {
	return _LoAfo.Contract.GetTraceInfo(&_LoAfo.CallOpts, _traceHash)
}

// Histories is a free data retrieval call binding the contract method 0x56564b7c.
//
// Solidity: function histories(uint256 ) view returns(uint256 HistoryID, address Operator, string Operation, uint256 OperationTime, uint256 RelatedLostItemID)
func (_LoAfo *LoAfoCaller) Histories(opts *bind.CallOpts, arg0 *big.Int) (struct {
	HistoryID         *big.Int
	Operator          common.Address
	Operation         string
	OperationTime     *big.Int
	RelatedLostItemID *big.Int
}, error) {
	var out []interface{}
	err := _LoAfo.contract.Call(opts, &out, "histories", arg0)

	outstruct := new(struct {
		HistoryID         *big.Int
		Operator          common.Address
		Operation         string
		OperationTime     *big.Int
		RelatedLostItemID *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.HistoryID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Operator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Operation = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.OperationTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RelatedLostItemID = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Histories is a free data retrieval call binding the contract method 0x56564b7c.
//
// Solidity: function histories(uint256 ) view returns(uint256 HistoryID, address Operator, string Operation, uint256 OperationTime, uint256 RelatedLostItemID)
func (_LoAfo *LoAfoSession) Histories(arg0 *big.Int) (struct {
	HistoryID         *big.Int
	Operator          common.Address
	Operation         string
	OperationTime     *big.Int
	RelatedLostItemID *big.Int
}, error) {
	return _LoAfo.Contract.Histories(&_LoAfo.CallOpts, arg0)
}

// Histories is a free data retrieval call binding the contract method 0x56564b7c.
//
// Solidity: function histories(uint256 ) view returns(uint256 HistoryID, address Operator, string Operation, uint256 OperationTime, uint256 RelatedLostItemID)
func (_LoAfo *LoAfoCallerSession) Histories(arg0 *big.Int) (struct {
	HistoryID         *big.Int
	Operator          common.Address
	Operation         string
	OperationTime     *big.Int
	RelatedLostItemID *big.Int
}, error) {
	return _LoAfo.Contract.Histories(&_LoAfo.CallOpts, arg0)
}

// LostItemExists is a free data retrieval call binding the contract method 0xbf21c30d.
//
// Solidity: function lostItemExists(uint256 _lostItemID) view returns(bool)
func (_LoAfo *LoAfoCaller) LostItemExists(opts *bind.CallOpts, _lostItemID *big.Int) (bool, error) {
	var out []interface{}
	err := _LoAfo.contract.Call(opts, &out, "lostItemExists", _lostItemID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LostItemExists is a free data retrieval call binding the contract method 0xbf21c30d.
//
// Solidity: function lostItemExists(uint256 _lostItemID) view returns(bool)
func (_LoAfo *LoAfoSession) LostItemExists(_lostItemID *big.Int) (bool, error) {
	return _LoAfo.Contract.LostItemExists(&_LoAfo.CallOpts, _lostItemID)
}

// LostItemExists is a free data retrieval call binding the contract method 0xbf21c30d.
//
// Solidity: function lostItemExists(uint256 _lostItemID) view returns(bool)
func (_LoAfo *LoAfoCallerSession) LostItemExists(_lostItemID *big.Int) (bool, error) {
	return _LoAfo.Contract.LostItemExists(&_LoAfo.CallOpts, _lostItemID)
}

// LostItems is a free data retrieval call binding the contract method 0x9e8d43a2.
//
// Solidity: function lostItems(uint256 ) view returns(uint256 LostItemID, address User, address FoundOne, string LostItemName, string Description, uint256 LostDate, string Location, string ItemCategory, bool IsFound, uint256 FoundDate, uint256 reward, uint256 FoundApplicationsCount)
func (_LoAfo *LoAfoCaller) LostItems(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LostItemID             *big.Int
	User                   common.Address
	FoundOne               common.Address
	LostItemName           string
	Description            string
	LostDate               *big.Int
	Location               string
	ItemCategory           string
	IsFound                bool
	FoundDate              *big.Int
	Reward                 *big.Int
	FoundApplicationsCount *big.Int
}, error) {
	var out []interface{}
	err := _LoAfo.contract.Call(opts, &out, "lostItems", arg0)

	outstruct := new(struct {
		LostItemID             *big.Int
		User                   common.Address
		FoundOne               common.Address
		LostItemName           string
		Description            string
		LostDate               *big.Int
		Location               string
		ItemCategory           string
		IsFound                bool
		FoundDate              *big.Int
		Reward                 *big.Int
		FoundApplicationsCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LostItemID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.User = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.FoundOne = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.LostItemName = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.LostDate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Location = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.ItemCategory = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.IsFound = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.FoundDate = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.Reward = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.FoundApplicationsCount = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LostItems is a free data retrieval call binding the contract method 0x9e8d43a2.
//
// Solidity: function lostItems(uint256 ) view returns(uint256 LostItemID, address User, address FoundOne, string LostItemName, string Description, uint256 LostDate, string Location, string ItemCategory, bool IsFound, uint256 FoundDate, uint256 reward, uint256 FoundApplicationsCount)
func (_LoAfo *LoAfoSession) LostItems(arg0 *big.Int) (struct {
	LostItemID             *big.Int
	User                   common.Address
	FoundOne               common.Address
	LostItemName           string
	Description            string
	LostDate               *big.Int
	Location               string
	ItemCategory           string
	IsFound                bool
	FoundDate              *big.Int
	Reward                 *big.Int
	FoundApplicationsCount *big.Int
}, error) {
	return _LoAfo.Contract.LostItems(&_LoAfo.CallOpts, arg0)
}

// LostItems is a free data retrieval call binding the contract method 0x9e8d43a2.
//
// Solidity: function lostItems(uint256 ) view returns(uint256 LostItemID, address User, address FoundOne, string LostItemName, string Description, uint256 LostDate, string Location, string ItemCategory, bool IsFound, uint256 FoundDate, uint256 reward, uint256 FoundApplicationsCount)
func (_LoAfo *LoAfoCallerSession) LostItems(arg0 *big.Int) (struct {
	LostItemID             *big.Int
	User                   common.Address
	FoundOne               common.Address
	LostItemName           string
	Description            string
	LostDate               *big.Int
	Location               string
	ItemCategory           string
	IsFound                bool
	FoundDate              *big.Int
	Reward                 *big.Int
	FoundApplicationsCount *big.Int
}, error) {
	return _LoAfo.Contract.LostItems(&_LoAfo.CallOpts, arg0)
}

// NextApplicationId is a free data retrieval call binding the contract method 0x625338cb.
//
// Solidity: function nextApplicationId() view returns(uint256)
func (_LoAfo *LoAfoCaller) NextApplicationId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LoAfo.contract.Call(opts, &out, "nextApplicationId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextApplicationId is a free data retrieval call binding the contract method 0x625338cb.
//
// Solidity: function nextApplicationId() view returns(uint256)
func (_LoAfo *LoAfoSession) NextApplicationId() (*big.Int, error) {
	return _LoAfo.Contract.NextApplicationId(&_LoAfo.CallOpts)
}

// NextApplicationId is a free data retrieval call binding the contract method 0x625338cb.
//
// Solidity: function nextApplicationId() view returns(uint256)
func (_LoAfo *LoAfoCallerSession) NextApplicationId() (*big.Int, error) {
	return _LoAfo.Contract.NextApplicationId(&_LoAfo.CallOpts)
}

// NextHistoryID is a free data retrieval call binding the contract method 0x43613992.
//
// Solidity: function nextHistoryID() view returns(uint256)
func (_LoAfo *LoAfoCaller) NextHistoryID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LoAfo.contract.Call(opts, &out, "nextHistoryID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextHistoryID is a free data retrieval call binding the contract method 0x43613992.
//
// Solidity: function nextHistoryID() view returns(uint256)
func (_LoAfo *LoAfoSession) NextHistoryID() (*big.Int, error) {
	return _LoAfo.Contract.NextHistoryID(&_LoAfo.CallOpts)
}

// NextHistoryID is a free data retrieval call binding the contract method 0x43613992.
//
// Solidity: function nextHistoryID() view returns(uint256)
func (_LoAfo *LoAfoCallerSession) NextHistoryID() (*big.Int, error) {
	return _LoAfo.Contract.NextHistoryID(&_LoAfo.CallOpts)
}

// NextLostItemID is a free data retrieval call binding the contract method 0x85d761ee.
//
// Solidity: function nextLostItemID() view returns(uint256)
func (_LoAfo *LoAfoCaller) NextLostItemID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LoAfo.contract.Call(opts, &out, "nextLostItemID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextLostItemID is a free data retrieval call binding the contract method 0x85d761ee.
//
// Solidity: function nextLostItemID() view returns(uint256)
func (_LoAfo *LoAfoSession) NextLostItemID() (*big.Int, error) {
	return _LoAfo.Contract.NextLostItemID(&_LoAfo.CallOpts)
}

// NextLostItemID is a free data retrieval call binding the contract method 0x85d761ee.
//
// Solidity: function nextLostItemID() view returns(uint256)
func (_LoAfo *LoAfoCallerSession) NextLostItemID() (*big.Int, error) {
	return _LoAfo.Contract.NextLostItemID(&_LoAfo.CallOpts)
}

// UserPoints is a free data retrieval call binding the contract method 0xf50ddbc7.
//
// Solidity: function userPoints(address ) view returns(uint256)
func (_LoAfo *LoAfoCaller) UserPoints(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LoAfo.contract.Call(opts, &out, "userPoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPoints is a free data retrieval call binding the contract method 0xf50ddbc7.
//
// Solidity: function userPoints(address ) view returns(uint256)
func (_LoAfo *LoAfoSession) UserPoints(arg0 common.Address) (*big.Int, error) {
	return _LoAfo.Contract.UserPoints(&_LoAfo.CallOpts, arg0)
}

// UserPoints is a free data retrieval call binding the contract method 0xf50ddbc7.
//
// Solidity: function userPoints(address ) view returns(uint256)
func (_LoAfo *LoAfoCallerSession) UserPoints(arg0 common.Address) (*big.Int, error) {
	return _LoAfo.Contract.UserPoints(&_LoAfo.CallOpts, arg0)
}

// AddHistoryRecord is a paid mutator transaction binding the contract method 0xcb4a5c2c.
//
// Solidity: function addHistoryRecord(uint256 _LostItemID, string _Operation, uint256 _OperationTime) returns()
func (_LoAfo *LoAfoTransactor) AddHistoryRecord(opts *bind.TransactOpts, _LostItemID *big.Int, _Operation string, _OperationTime *big.Int) (*types.Transaction, error) {
	return _LoAfo.contract.Transact(opts, "addHistoryRecord", _LostItemID, _Operation, _OperationTime)
}

// AddHistoryRecord is a paid mutator transaction binding the contract method 0xcb4a5c2c.
//
// Solidity: function addHistoryRecord(uint256 _LostItemID, string _Operation, uint256 _OperationTime) returns()
func (_LoAfo *LoAfoSession) AddHistoryRecord(_LostItemID *big.Int, _Operation string, _OperationTime *big.Int) (*types.Transaction, error) {
	return _LoAfo.Contract.AddHistoryRecord(&_LoAfo.TransactOpts, _LostItemID, _Operation, _OperationTime)
}

// AddHistoryRecord is a paid mutator transaction binding the contract method 0xcb4a5c2c.
//
// Solidity: function addHistoryRecord(uint256 _LostItemID, string _Operation, uint256 _OperationTime) returns()
func (_LoAfo *LoAfoTransactorSession) AddHistoryRecord(_LostItemID *big.Int, _Operation string, _OperationTime *big.Int) (*types.Transaction, error) {
	return _LoAfo.Contract.AddHistoryRecord(&_LoAfo.TransactOpts, _LostItemID, _Operation, _OperationTime)
}

// AddLostItem is a paid mutator transaction binding the contract method 0xe45c925c.
//
// Solidity: function addLostItem(string _lostItemName, string _description, uint256 _lostDate, string _location, string _itemCategory, uint256 _reward) returns()
func (_LoAfo *LoAfoTransactor) AddLostItem(opts *bind.TransactOpts, _lostItemName string, _description string, _lostDate *big.Int, _location string, _itemCategory string, _reward *big.Int) (*types.Transaction, error) {
	return _LoAfo.contract.Transact(opts, "addLostItem", _lostItemName, _description, _lostDate, _location, _itemCategory, _reward)
}

// AddLostItem is a paid mutator transaction binding the contract method 0xe45c925c.
//
// Solidity: function addLostItem(string _lostItemName, string _description, uint256 _lostDate, string _location, string _itemCategory, uint256 _reward) returns()
func (_LoAfo *LoAfoSession) AddLostItem(_lostItemName string, _description string, _lostDate *big.Int, _location string, _itemCategory string, _reward *big.Int) (*types.Transaction, error) {
	return _LoAfo.Contract.AddLostItem(&_LoAfo.TransactOpts, _lostItemName, _description, _lostDate, _location, _itemCategory, _reward)
}

// AddLostItem is a paid mutator transaction binding the contract method 0xe45c925c.
//
// Solidity: function addLostItem(string _lostItemName, string _description, uint256 _lostDate, string _location, string _itemCategory, uint256 _reward) returns()
func (_LoAfo *LoAfoTransactorSession) AddLostItem(_lostItemName string, _description string, _lostDate *big.Int, _location string, _itemCategory string, _reward *big.Int) (*types.Transaction, error) {
	return _LoAfo.Contract.AddLostItem(&_LoAfo.TransactOpts, _lostItemName, _description, _lostDate, _location, _itemCategory, _reward)
}

// ConfirmFoundApplication is a paid mutator transaction binding the contract method 0x40913c25.
//
// Solidity: function confirmFoundApplication(uint256 _lostItemID, uint256 _applicationId, bool _result) payable returns(bytes32 traceHash)
func (_LoAfo *LoAfoTransactor) ConfirmFoundApplication(opts *bind.TransactOpts, _lostItemID *big.Int, _applicationId *big.Int, _result bool) (*types.Transaction, error) {
	return _LoAfo.contract.Transact(opts, "confirmFoundApplication", _lostItemID, _applicationId, _result)
}

// ConfirmFoundApplication is a paid mutator transaction binding the contract method 0x40913c25.
//
// Solidity: function confirmFoundApplication(uint256 _lostItemID, uint256 _applicationId, bool _result) payable returns(bytes32 traceHash)
func (_LoAfo *LoAfoSession) ConfirmFoundApplication(_lostItemID *big.Int, _applicationId *big.Int, _result bool) (*types.Transaction, error) {
	return _LoAfo.Contract.ConfirmFoundApplication(&_LoAfo.TransactOpts, _lostItemID, _applicationId, _result)
}

// ConfirmFoundApplication is a paid mutator transaction binding the contract method 0x40913c25.
//
// Solidity: function confirmFoundApplication(uint256 _lostItemID, uint256 _applicationId, bool _result) payable returns(bytes32 traceHash)
func (_LoAfo *LoAfoTransactorSession) ConfirmFoundApplication(_lostItemID *big.Int, _applicationId *big.Int, _result bool) (*types.Transaction, error) {
	return _LoAfo.Contract.ConfirmFoundApplication(&_LoAfo.TransactOpts, _lostItemID, _applicationId, _result)
}

// DecereaseUserPoints is a paid mutator transaction binding the contract method 0xf22ce773.
//
// Solidity: function decereaseUserPoints(address _user, uint256 _points) returns()
func (_LoAfo *LoAfoTransactor) DecereaseUserPoints(opts *bind.TransactOpts, _user common.Address, _points *big.Int) (*types.Transaction, error) {
	return _LoAfo.contract.Transact(opts, "decereaseUserPoints", _user, _points)
}

// DecereaseUserPoints is a paid mutator transaction binding the contract method 0xf22ce773.
//
// Solidity: function decereaseUserPoints(address _user, uint256 _points) returns()
func (_LoAfo *LoAfoSession) DecereaseUserPoints(_user common.Address, _points *big.Int) (*types.Transaction, error) {
	return _LoAfo.Contract.DecereaseUserPoints(&_LoAfo.TransactOpts, _user, _points)
}

// DecereaseUserPoints is a paid mutator transaction binding the contract method 0xf22ce773.
//
// Solidity: function decereaseUserPoints(address _user, uint256 _points) returns()
func (_LoAfo *LoAfoTransactorSession) DecereaseUserPoints(_user common.Address, _points *big.Int) (*types.Transaction, error) {
	return _LoAfo.Contract.DecereaseUserPoints(&_LoAfo.TransactOpts, _user, _points)
}

// DeleteApplication is a paid mutator transaction binding the contract method 0xf7f41850.
//
// Solidity: function deleteApplication(uint256 _lostItemID, uint256 _applicationId) returns()
func (_LoAfo *LoAfoTransactor) DeleteApplication(opts *bind.TransactOpts, _lostItemID *big.Int, _applicationId *big.Int) (*types.Transaction, error) {
	return _LoAfo.contract.Transact(opts, "deleteApplication", _lostItemID, _applicationId)
}

// DeleteApplication is a paid mutator transaction binding the contract method 0xf7f41850.
//
// Solidity: function deleteApplication(uint256 _lostItemID, uint256 _applicationId) returns()
func (_LoAfo *LoAfoSession) DeleteApplication(_lostItemID *big.Int, _applicationId *big.Int) (*types.Transaction, error) {
	return _LoAfo.Contract.DeleteApplication(&_LoAfo.TransactOpts, _lostItemID, _applicationId)
}

// DeleteApplication is a paid mutator transaction binding the contract method 0xf7f41850.
//
// Solidity: function deleteApplication(uint256 _lostItemID, uint256 _applicationId) returns()
func (_LoAfo *LoAfoTransactorSession) DeleteApplication(_lostItemID *big.Int, _applicationId *big.Int) (*types.Transaction, error) {
	return _LoAfo.Contract.DeleteApplication(&_LoAfo.TransactOpts, _lostItemID, _applicationId)
}

// DeleteLostItem is a paid mutator transaction binding the contract method 0x88b560c3.
//
// Solidity: function deleteLostItem(uint256 _lostItemID) returns()
func (_LoAfo *LoAfoTransactor) DeleteLostItem(opts *bind.TransactOpts, _lostItemID *big.Int) (*types.Transaction, error) {
	return _LoAfo.contract.Transact(opts, "deleteLostItem", _lostItemID)
}

// DeleteLostItem is a paid mutator transaction binding the contract method 0x88b560c3.
//
// Solidity: function deleteLostItem(uint256 _lostItemID) returns()
func (_LoAfo *LoAfoSession) DeleteLostItem(_lostItemID *big.Int) (*types.Transaction, error) {
	return _LoAfo.Contract.DeleteLostItem(&_LoAfo.TransactOpts, _lostItemID)
}

// DeleteLostItem is a paid mutator transaction binding the contract method 0x88b560c3.
//
// Solidity: function deleteLostItem(uint256 _lostItemID) returns()
func (_LoAfo *LoAfoTransactorSession) DeleteLostItem(_lostItemID *big.Int) (*types.Transaction, error) {
	return _LoAfo.Contract.DeleteLostItem(&_LoAfo.TransactOpts, _lostItemID)
}

// IncreaseUserPoints is a paid mutator transaction binding the contract method 0x1d8d83b0.
//
// Solidity: function increaseUserPoints(address _user, uint256 _points) returns()
func (_LoAfo *LoAfoTransactor) IncreaseUserPoints(opts *bind.TransactOpts, _user common.Address, _points *big.Int) (*types.Transaction, error) {
	return _LoAfo.contract.Transact(opts, "increaseUserPoints", _user, _points)
}

// IncreaseUserPoints is a paid mutator transaction binding the contract method 0x1d8d83b0.
//
// Solidity: function increaseUserPoints(address _user, uint256 _points) returns()
func (_LoAfo *LoAfoSession) IncreaseUserPoints(_user common.Address, _points *big.Int) (*types.Transaction, error) {
	return _LoAfo.Contract.IncreaseUserPoints(&_LoAfo.TransactOpts, _user, _points)
}

// IncreaseUserPoints is a paid mutator transaction binding the contract method 0x1d8d83b0.
//
// Solidity: function increaseUserPoints(address _user, uint256 _points) returns()
func (_LoAfo *LoAfoTransactorSession) IncreaseUserPoints(_user common.Address, _points *big.Int) (*types.Transaction, error) {
	return _LoAfo.Contract.IncreaseUserPoints(&_LoAfo.TransactOpts, _user, _points)
}

// SendToCaller is a paid mutator transaction binding the contract method 0x45dcb6ad.
//
// Solidity: function sendToCaller(uint256 amount) payable returns()
func (_LoAfo *LoAfoTransactor) SendToCaller(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LoAfo.contract.Transact(opts, "sendToCaller", amount)
}

// SendToCaller is a paid mutator transaction binding the contract method 0x45dcb6ad.
//
// Solidity: function sendToCaller(uint256 amount) payable returns()
func (_LoAfo *LoAfoSession) SendToCaller(amount *big.Int) (*types.Transaction, error) {
	return _LoAfo.Contract.SendToCaller(&_LoAfo.TransactOpts, amount)
}

// SendToCaller is a paid mutator transaction binding the contract method 0x45dcb6ad.
//
// Solidity: function sendToCaller(uint256 amount) payable returns()
func (_LoAfo *LoAfoTransactorSession) SendToCaller(amount *big.Int) (*types.Transaction, error) {
	return _LoAfo.Contract.SendToCaller(&_LoAfo.TransactOpts, amount)
}

// SetAdminAddress is a paid mutator transaction binding the contract method 0x2c1e816d.
//
// Solidity: function setAdminAddress(address _newAdminAddress) returns()
func (_LoAfo *LoAfoTransactor) SetAdminAddress(opts *bind.TransactOpts, _newAdminAddress common.Address) (*types.Transaction, error) {
	return _LoAfo.contract.Transact(opts, "setAdminAddress", _newAdminAddress)
}

// SetAdminAddress is a paid mutator transaction binding the contract method 0x2c1e816d.
//
// Solidity: function setAdminAddress(address _newAdminAddress) returns()
func (_LoAfo *LoAfoSession) SetAdminAddress(_newAdminAddress common.Address) (*types.Transaction, error) {
	return _LoAfo.Contract.SetAdminAddress(&_LoAfo.TransactOpts, _newAdminAddress)
}

// SetAdminAddress is a paid mutator transaction binding the contract method 0x2c1e816d.
//
// Solidity: function setAdminAddress(address _newAdminAddress) returns()
func (_LoAfo *LoAfoTransactorSession) SetAdminAddress(_newAdminAddress common.Address) (*types.Transaction, error) {
	return _LoAfo.Contract.SetAdminAddress(&_LoAfo.TransactOpts, _newAdminAddress)
}

// SubmitFoundApplication is a paid mutator transaction binding the contract method 0xef7df1af.
//
// Solidity: function submitFoundApplication(uint256 _lostItemID) returns()
func (_LoAfo *LoAfoTransactor) SubmitFoundApplication(opts *bind.TransactOpts, _lostItemID *big.Int) (*types.Transaction, error) {
	return _LoAfo.contract.Transact(opts, "submitFoundApplication", _lostItemID)
}

// SubmitFoundApplication is a paid mutator transaction binding the contract method 0xef7df1af.
//
// Solidity: function submitFoundApplication(uint256 _lostItemID) returns()
func (_LoAfo *LoAfoSession) SubmitFoundApplication(_lostItemID *big.Int) (*types.Transaction, error) {
	return _LoAfo.Contract.SubmitFoundApplication(&_LoAfo.TransactOpts, _lostItemID)
}

// SubmitFoundApplication is a paid mutator transaction binding the contract method 0xef7df1af.
//
// Solidity: function submitFoundApplication(uint256 _lostItemID) returns()
func (_LoAfo *LoAfoTransactorSession) SubmitFoundApplication(_lostItemID *big.Int) (*types.Transaction, error) {
	return _LoAfo.Contract.SubmitFoundApplication(&_LoAfo.TransactOpts, _lostItemID)
}

// UpdateLostItem is a paid mutator transaction binding the contract method 0xa357963b.
//
// Solidity: function updateLostItem(uint256 _lostItemID, string _lostItemName, string _description, uint256 _lostDate, string _location, string _itemCategory, uint256 _reward) returns()
func (_LoAfo *LoAfoTransactor) UpdateLostItem(opts *bind.TransactOpts, _lostItemID *big.Int, _lostItemName string, _description string, _lostDate *big.Int, _location string, _itemCategory string, _reward *big.Int) (*types.Transaction, error) {
	return _LoAfo.contract.Transact(opts, "updateLostItem", _lostItemID, _lostItemName, _description, _lostDate, _location, _itemCategory, _reward)
}

// UpdateLostItem is a paid mutator transaction binding the contract method 0xa357963b.
//
// Solidity: function updateLostItem(uint256 _lostItemID, string _lostItemName, string _description, uint256 _lostDate, string _location, string _itemCategory, uint256 _reward) returns()
func (_LoAfo *LoAfoSession) UpdateLostItem(_lostItemID *big.Int, _lostItemName string, _description string, _lostDate *big.Int, _location string, _itemCategory string, _reward *big.Int) (*types.Transaction, error) {
	return _LoAfo.Contract.UpdateLostItem(&_LoAfo.TransactOpts, _lostItemID, _lostItemName, _description, _lostDate, _location, _itemCategory, _reward)
}

// UpdateLostItem is a paid mutator transaction binding the contract method 0xa357963b.
//
// Solidity: function updateLostItem(uint256 _lostItemID, string _lostItemName, string _description, uint256 _lostDate, string _location, string _itemCategory, uint256 _reward) returns()
func (_LoAfo *LoAfoTransactorSession) UpdateLostItem(_lostItemID *big.Int, _lostItemName string, _description string, _lostDate *big.Int, _location string, _itemCategory string, _reward *big.Int) (*types.Transaction, error) {
	return _LoAfo.Contract.UpdateLostItem(&_LoAfo.TransactOpts, _lostItemID, _lostItemName, _description, _lostDate, _location, _itemCategory, _reward)
}

// LoAfoLostItemAddedIterator is returned from FilterLostItemAdded and is used to iterate over the raw logs and unpacked data for LostItemAdded events raised by the LoAfo contract.
type LoAfoLostItemAddedIterator struct {
	Event *LoAfoLostItemAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LoAfoLostItemAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoAfoLostItemAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LoAfoLostItemAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LoAfoLostItemAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoAfoLostItemAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoAfoLostItemAdded represents a LostItemAdded event raised by the LoAfo contract.
type LoAfoLostItemAdded struct {
	LostItemID   *big.Int
	User         common.Address
	LostItemName string
	Location     string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLostItemAdded is a free log retrieval operation binding the contract event 0xe39daa90fdc4de7343e85d7b94da86f5d7c7da29e2f83c1105be099dcf44fc9f.
//
// Solidity: event LostItemAdded(uint256 lostItemID, address user, string lostItemName, string location)
func (_LoAfo *LoAfoFilterer) FilterLostItemAdded(opts *bind.FilterOpts) (*LoAfoLostItemAddedIterator, error) {

	logs, sub, err := _LoAfo.contract.FilterLogs(opts, "LostItemAdded")
	if err != nil {
		return nil, err
	}
	return &LoAfoLostItemAddedIterator{contract: _LoAfo.contract, event: "LostItemAdded", logs: logs, sub: sub}, nil
}

// WatchLostItemAdded is a free log subscription operation binding the contract event 0xe39daa90fdc4de7343e85d7b94da86f5d7c7da29e2f83c1105be099dcf44fc9f.
//
// Solidity: event LostItemAdded(uint256 lostItemID, address user, string lostItemName, string location)
func (_LoAfo *LoAfoFilterer) WatchLostItemAdded(opts *bind.WatchOpts, sink chan<- *LoAfoLostItemAdded) (event.Subscription, error) {

	logs, sub, err := _LoAfo.contract.WatchLogs(opts, "LostItemAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoAfoLostItemAdded)
				if err := _LoAfo.contract.UnpackLog(event, "LostItemAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLostItemAdded is a log parse operation binding the contract event 0xe39daa90fdc4de7343e85d7b94da86f5d7c7da29e2f83c1105be099dcf44fc9f.
//
// Solidity: event LostItemAdded(uint256 lostItemID, address user, string lostItemName, string location)
func (_LoAfo *LoAfoFilterer) ParseLostItemAdded(log types.Log) (*LoAfoLostItemAdded, error) {
	event := new(LoAfoLostItemAdded)
	if err := _LoAfo.contract.UnpackLog(event, "LostItemAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
