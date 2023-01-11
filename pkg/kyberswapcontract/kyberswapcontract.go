// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kyberswapcontract

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

// MetaAggregationRouterSwapDescription is an auto generated low-level Go binding around an user-defined struct.
type MetaAggregationRouterSwapDescription struct {
	SrcToken        common.Address
	DstToken        common.Address
	SrcReceivers    []common.Address
	SrcAmounts      []*big.Int
	DstReceiver     common.Address
	Amount          *big.Int
	MinReturnAmount *big.Int
	Flags           *big.Int
	Permit          []byte
}

// SwapDescriptionExecutor1Inch is an auto generated low-level Go binding around an user-defined struct.
type SwapDescriptionExecutor1Inch struct {
	SrcToken         common.Address
	DstToken         common.Address
	SrcReceiver1Inch common.Address
	DstReceiver      common.Address
	SrcReceivers     []common.Address
	SrcAmounts       []*big.Int
	Amount           *big.Int
	MinReturnAmount  *big.Int
	Flags            *big.Int
	Permit           []byte
}

// KyberswapcontractMetaData contains all meta data concerning the Kyberswapcontract contract.
var KyberswapcontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"ClientData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"Error\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"}],\"name\":\"Exchange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spentAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"name\":\"Swapped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouter.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"executorData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor1Inch\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"srcReceiver1Inch\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structSwapDescriptionExecutor1Inch\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"executor1InchData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"swapExecutor1Inch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router1Inch\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"router1InchData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouter.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"swapRouter1Inch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouter.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"executorData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"swapSimpleMode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"updateWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// KyberswapcontractABI is the input ABI used to generate the binding from.
// Deprecated: Use KyberswapcontractMetaData.ABI instead.
var KyberswapcontractABI = KyberswapcontractMetaData.ABI

// Kyberswapcontract is an auto generated Go binding around an Ethereum contract.
type Kyberswapcontract struct {
	KyberswapcontractCaller     // Read-only binding to the contract
	KyberswapcontractTransactor // Write-only binding to the contract
	KyberswapcontractFilterer   // Log filterer for contract events
}

// KyberswapcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type KyberswapcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberswapcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KyberswapcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberswapcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KyberswapcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberswapcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KyberswapcontractSession struct {
	Contract     *Kyberswapcontract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// KyberswapcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KyberswapcontractCallerSession struct {
	Contract *KyberswapcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// KyberswapcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KyberswapcontractTransactorSession struct {
	Contract     *KyberswapcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// KyberswapcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type KyberswapcontractRaw struct {
	Contract *Kyberswapcontract // Generic contract binding to access the raw methods on
}

// KyberswapcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KyberswapcontractCallerRaw struct {
	Contract *KyberswapcontractCaller // Generic read-only contract binding to access the raw methods on
}

// KyberswapcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KyberswapcontractTransactorRaw struct {
	Contract *KyberswapcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKyberswapcontract creates a new instance of Kyberswapcontract, bound to a specific deployed contract.
func NewKyberswapcontract(address common.Address, backend bind.ContractBackend) (*Kyberswapcontract, error) {
	contract, err := bindKyberswapcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kyberswapcontract{KyberswapcontractCaller: KyberswapcontractCaller{contract: contract}, KyberswapcontractTransactor: KyberswapcontractTransactor{contract: contract}, KyberswapcontractFilterer: KyberswapcontractFilterer{contract: contract}}, nil
}

// NewKyberswapcontractCaller creates a new read-only instance of Kyberswapcontract, bound to a specific deployed contract.
func NewKyberswapcontractCaller(address common.Address, caller bind.ContractCaller) (*KyberswapcontractCaller, error) {
	contract, err := bindKyberswapcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KyberswapcontractCaller{contract: contract}, nil
}

// NewKyberswapcontractTransactor creates a new write-only instance of Kyberswapcontract, bound to a specific deployed contract.
func NewKyberswapcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*KyberswapcontractTransactor, error) {
	contract, err := bindKyberswapcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KyberswapcontractTransactor{contract: contract}, nil
}

// NewKyberswapcontractFilterer creates a new log filterer instance of Kyberswapcontract, bound to a specific deployed contract.
func NewKyberswapcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*KyberswapcontractFilterer, error) {
	contract, err := bindKyberswapcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KyberswapcontractFilterer{contract: contract}, nil
}

// bindKyberswapcontract binds a generic wrapper to an already deployed contract.
func bindKyberswapcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KyberswapcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kyberswapcontract *KyberswapcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kyberswapcontract.Contract.KyberswapcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kyberswapcontract *KyberswapcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.KyberswapcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kyberswapcontract *KyberswapcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.KyberswapcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kyberswapcontract *KyberswapcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kyberswapcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kyberswapcontract *KyberswapcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kyberswapcontract *KyberswapcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Kyberswapcontract *KyberswapcontractCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kyberswapcontract.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Kyberswapcontract *KyberswapcontractSession) WETH() (common.Address, error) {
	return _Kyberswapcontract.Contract.WETH(&_Kyberswapcontract.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Kyberswapcontract *KyberswapcontractCallerSession) WETH() (common.Address, error) {
	return _Kyberswapcontract.Contract.WETH(&_Kyberswapcontract.CallOpts)
}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(bool)
func (_Kyberswapcontract *KyberswapcontractCaller) IsWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Kyberswapcontract.contract.Call(opts, &out, "isWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(bool)
func (_Kyberswapcontract *KyberswapcontractSession) IsWhitelist(arg0 common.Address) (bool, error) {
	return _Kyberswapcontract.Contract.IsWhitelist(&_Kyberswapcontract.CallOpts, arg0)
}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(bool)
func (_Kyberswapcontract *KyberswapcontractCallerSession) IsWhitelist(arg0 common.Address) (bool, error) {
	return _Kyberswapcontract.Contract.IsWhitelist(&_Kyberswapcontract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Kyberswapcontract *KyberswapcontractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kyberswapcontract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Kyberswapcontract *KyberswapcontractSession) Owner() (common.Address, error) {
	return _Kyberswapcontract.Contract.Owner(&_Kyberswapcontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Kyberswapcontract *KyberswapcontractCallerSession) Owner() (common.Address, error) {
	return _Kyberswapcontract.Contract.Owner(&_Kyberswapcontract.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Kyberswapcontract *KyberswapcontractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kyberswapcontract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Kyberswapcontract *KyberswapcontractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.RenounceOwnership(&_Kyberswapcontract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Kyberswapcontract *KyberswapcontractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.RenounceOwnership(&_Kyberswapcontract.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Kyberswapcontract *KyberswapcontractTransactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kyberswapcontract.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Kyberswapcontract *KyberswapcontractSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.RescueFunds(&_Kyberswapcontract.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Kyberswapcontract *KyberswapcontractTransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.RescueFunds(&_Kyberswapcontract.TransactOpts, token, amount)
}

// Swap is a paid mutator transaction binding the contract method 0xabcffc26.
//
// Solidity: function swap(address caller, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapcontract *KyberswapcontractTransactor) Swap(opts *bind.TransactOpts, caller common.Address, desc MetaAggregationRouterSwapDescription, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapcontract.contract.Transact(opts, "swap", caller, desc, executorData, clientData)
}

// Swap is a paid mutator transaction binding the contract method 0xabcffc26.
//
// Solidity: function swap(address caller, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapcontract *KyberswapcontractSession) Swap(caller common.Address, desc MetaAggregationRouterSwapDescription, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.Swap(&_Kyberswapcontract.TransactOpts, caller, desc, executorData, clientData)
}

// Swap is a paid mutator transaction binding the contract method 0xabcffc26.
//
// Solidity: function swap(address caller, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapcontract *KyberswapcontractTransactorSession) Swap(caller common.Address, desc MetaAggregationRouterSwapDescription, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.Swap(&_Kyberswapcontract.TransactOpts, caller, desc, executorData, clientData)
}

// SwapExecutor1Inch is a paid mutator transaction binding the contract method 0x1fb3e235.
//
// Solidity: function swapExecutor1Inch(address caller, (address,address,address,address,address[],uint256[],uint256,uint256,uint256,bytes) desc, bytes executor1InchData, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapcontract *KyberswapcontractTransactor) SwapExecutor1Inch(opts *bind.TransactOpts, caller common.Address, desc SwapDescriptionExecutor1Inch, executor1InchData []byte, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapcontract.contract.Transact(opts, "swapExecutor1Inch", caller, desc, executor1InchData, clientData)
}

// SwapExecutor1Inch is a paid mutator transaction binding the contract method 0x1fb3e235.
//
// Solidity: function swapExecutor1Inch(address caller, (address,address,address,address,address[],uint256[],uint256,uint256,uint256,bytes) desc, bytes executor1InchData, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapcontract *KyberswapcontractSession) SwapExecutor1Inch(caller common.Address, desc SwapDescriptionExecutor1Inch, executor1InchData []byte, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.SwapExecutor1Inch(&_Kyberswapcontract.TransactOpts, caller, desc, executor1InchData, clientData)
}

// SwapExecutor1Inch is a paid mutator transaction binding the contract method 0x1fb3e235.
//
// Solidity: function swapExecutor1Inch(address caller, (address,address,address,address,address[],uint256[],uint256,uint256,uint256,bytes) desc, bytes executor1InchData, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapcontract *KyberswapcontractTransactorSession) SwapExecutor1Inch(caller common.Address, desc SwapDescriptionExecutor1Inch, executor1InchData []byte, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.SwapExecutor1Inch(&_Kyberswapcontract.TransactOpts, caller, desc, executor1InchData, clientData)
}

// SwapRouter1Inch is a paid mutator transaction binding the contract method 0xa72bd2f5.
//
// Solidity: function swapRouter1Inch(address router1Inch, bytes router1InchData, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapcontract *KyberswapcontractTransactor) SwapRouter1Inch(opts *bind.TransactOpts, router1Inch common.Address, router1InchData []byte, desc MetaAggregationRouterSwapDescription, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapcontract.contract.Transact(opts, "swapRouter1Inch", router1Inch, router1InchData, desc, clientData)
}

// SwapRouter1Inch is a paid mutator transaction binding the contract method 0xa72bd2f5.
//
// Solidity: function swapRouter1Inch(address router1Inch, bytes router1InchData, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapcontract *KyberswapcontractSession) SwapRouter1Inch(router1Inch common.Address, router1InchData []byte, desc MetaAggregationRouterSwapDescription, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.SwapRouter1Inch(&_Kyberswapcontract.TransactOpts, router1Inch, router1InchData, desc, clientData)
}

// SwapRouter1Inch is a paid mutator transaction binding the contract method 0xa72bd2f5.
//
// Solidity: function swapRouter1Inch(address router1Inch, bytes router1InchData, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapcontract *KyberswapcontractTransactorSession) SwapRouter1Inch(router1Inch common.Address, router1InchData []byte, desc MetaAggregationRouterSwapDescription, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.SwapRouter1Inch(&_Kyberswapcontract.TransactOpts, router1Inch, router1InchData, desc, clientData)
}

// SwapSimpleMode is a paid mutator transaction binding the contract method 0xa7f5c104.
//
// Solidity: function swapSimpleMode(address caller, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapcontract *KyberswapcontractTransactor) SwapSimpleMode(opts *bind.TransactOpts, caller common.Address, desc MetaAggregationRouterSwapDescription, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapcontract.contract.Transact(opts, "swapSimpleMode", caller, desc, executorData, clientData)
}

// SwapSimpleMode is a paid mutator transaction binding the contract method 0xa7f5c104.
//
// Solidity: function swapSimpleMode(address caller, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapcontract *KyberswapcontractSession) SwapSimpleMode(caller common.Address, desc MetaAggregationRouterSwapDescription, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.SwapSimpleMode(&_Kyberswapcontract.TransactOpts, caller, desc, executorData, clientData)
}

// SwapSimpleMode is a paid mutator transaction binding the contract method 0xa7f5c104.
//
// Solidity: function swapSimpleMode(address caller, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapcontract *KyberswapcontractTransactorSession) SwapSimpleMode(caller common.Address, desc MetaAggregationRouterSwapDescription, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.SwapSimpleMode(&_Kyberswapcontract.TransactOpts, caller, desc, executorData, clientData)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Kyberswapcontract *KyberswapcontractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Kyberswapcontract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Kyberswapcontract *KyberswapcontractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.TransferOwnership(&_Kyberswapcontract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Kyberswapcontract *KyberswapcontractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.TransferOwnership(&_Kyberswapcontract.TransactOpts, newOwner)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x0d392cd9.
//
// Solidity: function updateWhitelist(address addr, bool value) returns()
func (_Kyberswapcontract *KyberswapcontractTransactor) UpdateWhitelist(opts *bind.TransactOpts, addr common.Address, value bool) (*types.Transaction, error) {
	return _Kyberswapcontract.contract.Transact(opts, "updateWhitelist", addr, value)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x0d392cd9.
//
// Solidity: function updateWhitelist(address addr, bool value) returns()
func (_Kyberswapcontract *KyberswapcontractSession) UpdateWhitelist(addr common.Address, value bool) (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.UpdateWhitelist(&_Kyberswapcontract.TransactOpts, addr, value)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x0d392cd9.
//
// Solidity: function updateWhitelist(address addr, bool value) returns()
func (_Kyberswapcontract *KyberswapcontractTransactorSession) UpdateWhitelist(addr common.Address, value bool) (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.UpdateWhitelist(&_Kyberswapcontract.TransactOpts, addr, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Kyberswapcontract *KyberswapcontractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kyberswapcontract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Kyberswapcontract *KyberswapcontractSession) Receive() (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.Receive(&_Kyberswapcontract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Kyberswapcontract *KyberswapcontractTransactorSession) Receive() (*types.Transaction, error) {
	return _Kyberswapcontract.Contract.Receive(&_Kyberswapcontract.TransactOpts)
}

// KyberswapcontractClientDataIterator is returned from FilterClientData and is used to iterate over the raw logs and unpacked data for ClientData events raised by the Kyberswapcontract contract.
type KyberswapcontractClientDataIterator struct {
	Event *KyberswapcontractClientData // Event containing the contract specifics and raw log

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
func (it *KyberswapcontractClientDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberswapcontractClientData)
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
		it.Event = new(KyberswapcontractClientData)
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
func (it *KyberswapcontractClientDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberswapcontractClientDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberswapcontractClientData represents a ClientData event raised by the Kyberswapcontract contract.
type KyberswapcontractClientData struct {
	ClientData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClientData is a free log retrieval operation binding the contract event 0x095e66fa4dd6a6f7b43fb8444a7bd0edb870508c7abf639bc216efb0bcff9779.
//
// Solidity: event ClientData(bytes clientData)
func (_Kyberswapcontract *KyberswapcontractFilterer) FilterClientData(opts *bind.FilterOpts) (*KyberswapcontractClientDataIterator, error) {

	logs, sub, err := _Kyberswapcontract.contract.FilterLogs(opts, "ClientData")
	if err != nil {
		return nil, err
	}
	return &KyberswapcontractClientDataIterator{contract: _Kyberswapcontract.contract, event: "ClientData", logs: logs, sub: sub}, nil
}

// WatchClientData is a free log subscription operation binding the contract event 0x095e66fa4dd6a6f7b43fb8444a7bd0edb870508c7abf639bc216efb0bcff9779.
//
// Solidity: event ClientData(bytes clientData)
func (_Kyberswapcontract *KyberswapcontractFilterer) WatchClientData(opts *bind.WatchOpts, sink chan<- *KyberswapcontractClientData) (event.Subscription, error) {

	logs, sub, err := _Kyberswapcontract.contract.WatchLogs(opts, "ClientData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberswapcontractClientData)
				if err := _Kyberswapcontract.contract.UnpackLog(event, "ClientData", log); err != nil {
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

// ParseClientData is a log parse operation binding the contract event 0x095e66fa4dd6a6f7b43fb8444a7bd0edb870508c7abf639bc216efb0bcff9779.
//
// Solidity: event ClientData(bytes clientData)
func (_Kyberswapcontract *KyberswapcontractFilterer) ParseClientData(log types.Log) (*KyberswapcontractClientData, error) {
	event := new(KyberswapcontractClientData)
	if err := _Kyberswapcontract.contract.UnpackLog(event, "ClientData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberswapcontractErrorIterator is returned from FilterError and is used to iterate over the raw logs and unpacked data for Error events raised by the Kyberswapcontract contract.
type KyberswapcontractErrorIterator struct {
	Event *KyberswapcontractError // Event containing the contract specifics and raw log

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
func (it *KyberswapcontractErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberswapcontractError)
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
		it.Event = new(KyberswapcontractError)
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
func (it *KyberswapcontractErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberswapcontractErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberswapcontractError represents a Error event raised by the Kyberswapcontract contract.
type KyberswapcontractError struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterError is a free log retrieval operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_Kyberswapcontract *KyberswapcontractFilterer) FilterError(opts *bind.FilterOpts) (*KyberswapcontractErrorIterator, error) {

	logs, sub, err := _Kyberswapcontract.contract.FilterLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return &KyberswapcontractErrorIterator{contract: _Kyberswapcontract.contract, event: "Error", logs: logs, sub: sub}, nil
}

// WatchError is a free log subscription operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_Kyberswapcontract *KyberswapcontractFilterer) WatchError(opts *bind.WatchOpts, sink chan<- *KyberswapcontractError) (event.Subscription, error) {

	logs, sub, err := _Kyberswapcontract.contract.WatchLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberswapcontractError)
				if err := _Kyberswapcontract.contract.UnpackLog(event, "Error", log); err != nil {
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

// ParseError is a log parse operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_Kyberswapcontract *KyberswapcontractFilterer) ParseError(log types.Log) (*KyberswapcontractError, error) {
	event := new(KyberswapcontractError)
	if err := _Kyberswapcontract.contract.UnpackLog(event, "Error", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberswapcontractExchangeIterator is returned from FilterExchange and is used to iterate over the raw logs and unpacked data for Exchange events raised by the Kyberswapcontract contract.
type KyberswapcontractExchangeIterator struct {
	Event *KyberswapcontractExchange // Event containing the contract specifics and raw log

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
func (it *KyberswapcontractExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberswapcontractExchange)
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
		it.Event = new(KyberswapcontractExchange)
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
func (it *KyberswapcontractExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberswapcontractExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberswapcontractExchange represents a Exchange event raised by the Kyberswapcontract contract.
type KyberswapcontractExchange struct {
	Pair      common.Address
	AmountOut *big.Int
	Output    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExchange is a free log retrieval operation binding the contract event 0xddac40937f35385a34f721af292e5a83fc5b840f722bff57c2fc71adba708c48.
//
// Solidity: event Exchange(address pair, uint256 amountOut, address output)
func (_Kyberswapcontract *KyberswapcontractFilterer) FilterExchange(opts *bind.FilterOpts) (*KyberswapcontractExchangeIterator, error) {

	logs, sub, err := _Kyberswapcontract.contract.FilterLogs(opts, "Exchange")
	if err != nil {
		return nil, err
	}
	return &KyberswapcontractExchangeIterator{contract: _Kyberswapcontract.contract, event: "Exchange", logs: logs, sub: sub}, nil
}

// WatchExchange is a free log subscription operation binding the contract event 0xddac40937f35385a34f721af292e5a83fc5b840f722bff57c2fc71adba708c48.
//
// Solidity: event Exchange(address pair, uint256 amountOut, address output)
func (_Kyberswapcontract *KyberswapcontractFilterer) WatchExchange(opts *bind.WatchOpts, sink chan<- *KyberswapcontractExchange) (event.Subscription, error) {

	logs, sub, err := _Kyberswapcontract.contract.WatchLogs(opts, "Exchange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberswapcontractExchange)
				if err := _Kyberswapcontract.contract.UnpackLog(event, "Exchange", log); err != nil {
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

// ParseExchange is a log parse operation binding the contract event 0xddac40937f35385a34f721af292e5a83fc5b840f722bff57c2fc71adba708c48.
//
// Solidity: event Exchange(address pair, uint256 amountOut, address output)
func (_Kyberswapcontract *KyberswapcontractFilterer) ParseExchange(log types.Log) (*KyberswapcontractExchange, error) {
	event := new(KyberswapcontractExchange)
	if err := _Kyberswapcontract.contract.UnpackLog(event, "Exchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberswapcontractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Kyberswapcontract contract.
type KyberswapcontractOwnershipTransferredIterator struct {
	Event *KyberswapcontractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KyberswapcontractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberswapcontractOwnershipTransferred)
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
		it.Event = new(KyberswapcontractOwnershipTransferred)
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
func (it *KyberswapcontractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberswapcontractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberswapcontractOwnershipTransferred represents a OwnershipTransferred event raised by the Kyberswapcontract contract.
type KyberswapcontractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kyberswapcontract *KyberswapcontractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KyberswapcontractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kyberswapcontract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KyberswapcontractOwnershipTransferredIterator{contract: _Kyberswapcontract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kyberswapcontract *KyberswapcontractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KyberswapcontractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kyberswapcontract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberswapcontractOwnershipTransferred)
				if err := _Kyberswapcontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kyberswapcontract *KyberswapcontractFilterer) ParseOwnershipTransferred(log types.Log) (*KyberswapcontractOwnershipTransferred, error) {
	event := new(KyberswapcontractOwnershipTransferred)
	if err := _Kyberswapcontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberswapcontractSwappedIterator is returned from FilterSwapped and is used to iterate over the raw logs and unpacked data for Swapped events raised by the Kyberswapcontract contract.
type KyberswapcontractSwappedIterator struct {
	Event *KyberswapcontractSwapped // Event containing the contract specifics and raw log

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
func (it *KyberswapcontractSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberswapcontractSwapped)
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
		it.Event = new(KyberswapcontractSwapped)
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
func (it *KyberswapcontractSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberswapcontractSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberswapcontractSwapped represents a Swapped event raised by the Kyberswapcontract contract.
type KyberswapcontractSwapped struct {
	Sender       common.Address
	SrcToken     common.Address
	DstToken     common.Address
	DstReceiver  common.Address
	SpentAmount  *big.Int
	ReturnAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwapped is a free log retrieval operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_Kyberswapcontract *KyberswapcontractFilterer) FilterSwapped(opts *bind.FilterOpts) (*KyberswapcontractSwappedIterator, error) {

	logs, sub, err := _Kyberswapcontract.contract.FilterLogs(opts, "Swapped")
	if err != nil {
		return nil, err
	}
	return &KyberswapcontractSwappedIterator{contract: _Kyberswapcontract.contract, event: "Swapped", logs: logs, sub: sub}, nil
}

// WatchSwapped is a free log subscription operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_Kyberswapcontract *KyberswapcontractFilterer) WatchSwapped(opts *bind.WatchOpts, sink chan<- *KyberswapcontractSwapped) (event.Subscription, error) {

	logs, sub, err := _Kyberswapcontract.contract.WatchLogs(opts, "Swapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberswapcontractSwapped)
				if err := _Kyberswapcontract.contract.UnpackLog(event, "Swapped", log); err != nil {
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

// ParseSwapped is a log parse operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_Kyberswapcontract *KyberswapcontractFilterer) ParseSwapped(log types.Log) (*KyberswapcontractSwapped, error) {
	event := new(KyberswapcontractSwapped)
	if err := _Kyberswapcontract.contract.UnpackLog(event, "Swapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
