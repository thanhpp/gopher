// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kyberswapbsc

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

// KyberswapbscMetaData contains all meta data concerning the Kyberswapbsc contract.
var KyberswapbscMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"ClientData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"Error\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"}],\"name\":\"Exchange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spentAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"name\":\"Swapped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouter.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"executorData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor1Inch\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"srcReceiver1Inch\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structSwapDescriptionExecutor1Inch\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"executor1InchData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"swapExecutor1Inch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router1Inch\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"router1InchData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouter.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"swapRouter1Inch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouter.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"executorData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"swapSimpleMode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"updateWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// KyberswapbscABI is the input ABI used to generate the binding from.
// Deprecated: Use KyberswapbscMetaData.ABI instead.
var KyberswapbscABI = KyberswapbscMetaData.ABI

// Kyberswapbsc is an auto generated Go binding around an Ethereum contract.
type Kyberswapbsc struct {
	KyberswapbscCaller     // Read-only binding to the contract
	KyberswapbscTransactor // Write-only binding to the contract
	KyberswapbscFilterer   // Log filterer for contract events
}

// KyberswapbscCaller is an auto generated read-only Go binding around an Ethereum contract.
type KyberswapbscCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberswapbscTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KyberswapbscTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberswapbscFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KyberswapbscFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KyberswapbscSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KyberswapbscSession struct {
	Contract     *Kyberswapbsc     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KyberswapbscCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KyberswapbscCallerSession struct {
	Contract *KyberswapbscCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// KyberswapbscTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KyberswapbscTransactorSession struct {
	Contract     *KyberswapbscTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// KyberswapbscRaw is an auto generated low-level Go binding around an Ethereum contract.
type KyberswapbscRaw struct {
	Contract *Kyberswapbsc // Generic contract binding to access the raw methods on
}

// KyberswapbscCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KyberswapbscCallerRaw struct {
	Contract *KyberswapbscCaller // Generic read-only contract binding to access the raw methods on
}

// KyberswapbscTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KyberswapbscTransactorRaw struct {
	Contract *KyberswapbscTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKyberswapbsc creates a new instance of Kyberswapbsc, bound to a specific deployed contract.
func NewKyberswapbsc(address common.Address, backend bind.ContractBackend) (*Kyberswapbsc, error) {
	contract, err := bindKyberswapbsc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kyberswapbsc{KyberswapbscCaller: KyberswapbscCaller{contract: contract}, KyberswapbscTransactor: KyberswapbscTransactor{contract: contract}, KyberswapbscFilterer: KyberswapbscFilterer{contract: contract}}, nil
}

// NewKyberswapbscCaller creates a new read-only instance of Kyberswapbsc, bound to a specific deployed contract.
func NewKyberswapbscCaller(address common.Address, caller bind.ContractCaller) (*KyberswapbscCaller, error) {
	contract, err := bindKyberswapbsc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KyberswapbscCaller{contract: contract}, nil
}

// NewKyberswapbscTransactor creates a new write-only instance of Kyberswapbsc, bound to a specific deployed contract.
func NewKyberswapbscTransactor(address common.Address, transactor bind.ContractTransactor) (*KyberswapbscTransactor, error) {
	contract, err := bindKyberswapbsc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KyberswapbscTransactor{contract: contract}, nil
}

// NewKyberswapbscFilterer creates a new log filterer instance of Kyberswapbsc, bound to a specific deployed contract.
func NewKyberswapbscFilterer(address common.Address, filterer bind.ContractFilterer) (*KyberswapbscFilterer, error) {
	contract, err := bindKyberswapbsc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KyberswapbscFilterer{contract: contract}, nil
}

// bindKyberswapbsc binds a generic wrapper to an already deployed contract.
func bindKyberswapbsc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KyberswapbscABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kyberswapbsc *KyberswapbscRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kyberswapbsc.Contract.KyberswapbscCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kyberswapbsc *KyberswapbscRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.KyberswapbscTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kyberswapbsc *KyberswapbscRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.KyberswapbscTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kyberswapbsc *KyberswapbscCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kyberswapbsc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kyberswapbsc *KyberswapbscTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kyberswapbsc *KyberswapbscTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Kyberswapbsc *KyberswapbscCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kyberswapbsc.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Kyberswapbsc *KyberswapbscSession) WETH() (common.Address, error) {
	return _Kyberswapbsc.Contract.WETH(&_Kyberswapbsc.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Kyberswapbsc *KyberswapbscCallerSession) WETH() (common.Address, error) {
	return _Kyberswapbsc.Contract.WETH(&_Kyberswapbsc.CallOpts)
}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(bool)
func (_Kyberswapbsc *KyberswapbscCaller) IsWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Kyberswapbsc.contract.Call(opts, &out, "isWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(bool)
func (_Kyberswapbsc *KyberswapbscSession) IsWhitelist(arg0 common.Address) (bool, error) {
	return _Kyberswapbsc.Contract.IsWhitelist(&_Kyberswapbsc.CallOpts, arg0)
}

// IsWhitelist is a free data retrieval call binding the contract method 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(bool)
func (_Kyberswapbsc *KyberswapbscCallerSession) IsWhitelist(arg0 common.Address) (bool, error) {
	return _Kyberswapbsc.Contract.IsWhitelist(&_Kyberswapbsc.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Kyberswapbsc *KyberswapbscCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kyberswapbsc.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Kyberswapbsc *KyberswapbscSession) Owner() (common.Address, error) {
	return _Kyberswapbsc.Contract.Owner(&_Kyberswapbsc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Kyberswapbsc *KyberswapbscCallerSession) Owner() (common.Address, error) {
	return _Kyberswapbsc.Contract.Owner(&_Kyberswapbsc.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Kyberswapbsc *KyberswapbscTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kyberswapbsc.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Kyberswapbsc *KyberswapbscSession) RenounceOwnership() (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.RenounceOwnership(&_Kyberswapbsc.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Kyberswapbsc *KyberswapbscTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.RenounceOwnership(&_Kyberswapbsc.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Kyberswapbsc *KyberswapbscTransactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kyberswapbsc.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Kyberswapbsc *KyberswapbscSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.RescueFunds(&_Kyberswapbsc.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_Kyberswapbsc *KyberswapbscTransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.RescueFunds(&_Kyberswapbsc.TransactOpts, token, amount)
}

// Swap is a paid mutator transaction binding the contract method 0xabcffc26.
//
// Solidity: function swap(address caller, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapbsc *KyberswapbscTransactor) Swap(opts *bind.TransactOpts, caller common.Address, desc MetaAggregationRouterSwapDescription, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapbsc.contract.Transact(opts, "swap", caller, desc, executorData, clientData)
}

// Swap is a paid mutator transaction binding the contract method 0xabcffc26.
//
// Solidity: function swap(address caller, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapbsc *KyberswapbscSession) Swap(caller common.Address, desc MetaAggregationRouterSwapDescription, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.Swap(&_Kyberswapbsc.TransactOpts, caller, desc, executorData, clientData)
}

// Swap is a paid mutator transaction binding the contract method 0xabcffc26.
//
// Solidity: function swap(address caller, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapbsc *KyberswapbscTransactorSession) Swap(caller common.Address, desc MetaAggregationRouterSwapDescription, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.Swap(&_Kyberswapbsc.TransactOpts, caller, desc, executorData, clientData)
}

// SwapExecutor1Inch is a paid mutator transaction binding the contract method 0x1fb3e235.
//
// Solidity: function swapExecutor1Inch(address caller, (address,address,address,address,address[],uint256[],uint256,uint256,uint256,bytes) desc, bytes executor1InchData, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapbsc *KyberswapbscTransactor) SwapExecutor1Inch(opts *bind.TransactOpts, caller common.Address, desc SwapDescriptionExecutor1Inch, executor1InchData []byte, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapbsc.contract.Transact(opts, "swapExecutor1Inch", caller, desc, executor1InchData, clientData)
}

// SwapExecutor1Inch is a paid mutator transaction binding the contract method 0x1fb3e235.
//
// Solidity: function swapExecutor1Inch(address caller, (address,address,address,address,address[],uint256[],uint256,uint256,uint256,bytes) desc, bytes executor1InchData, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapbsc *KyberswapbscSession) SwapExecutor1Inch(caller common.Address, desc SwapDescriptionExecutor1Inch, executor1InchData []byte, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.SwapExecutor1Inch(&_Kyberswapbsc.TransactOpts, caller, desc, executor1InchData, clientData)
}

// SwapExecutor1Inch is a paid mutator transaction binding the contract method 0x1fb3e235.
//
// Solidity: function swapExecutor1Inch(address caller, (address,address,address,address,address[],uint256[],uint256,uint256,uint256,bytes) desc, bytes executor1InchData, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapbsc *KyberswapbscTransactorSession) SwapExecutor1Inch(caller common.Address, desc SwapDescriptionExecutor1Inch, executor1InchData []byte, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.SwapExecutor1Inch(&_Kyberswapbsc.TransactOpts, caller, desc, executor1InchData, clientData)
}

// SwapRouter1Inch is a paid mutator transaction binding the contract method 0xa72bd2f5.
//
// Solidity: function swapRouter1Inch(address router1Inch, bytes router1InchData, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapbsc *KyberswapbscTransactor) SwapRouter1Inch(opts *bind.TransactOpts, router1Inch common.Address, router1InchData []byte, desc MetaAggregationRouterSwapDescription, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapbsc.contract.Transact(opts, "swapRouter1Inch", router1Inch, router1InchData, desc, clientData)
}

// SwapRouter1Inch is a paid mutator transaction binding the contract method 0xa72bd2f5.
//
// Solidity: function swapRouter1Inch(address router1Inch, bytes router1InchData, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapbsc *KyberswapbscSession) SwapRouter1Inch(router1Inch common.Address, router1InchData []byte, desc MetaAggregationRouterSwapDescription, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.SwapRouter1Inch(&_Kyberswapbsc.TransactOpts, router1Inch, router1InchData, desc, clientData)
}

// SwapRouter1Inch is a paid mutator transaction binding the contract method 0xa72bd2f5.
//
// Solidity: function swapRouter1Inch(address router1Inch, bytes router1InchData, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes clientData) payable returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapbsc *KyberswapbscTransactorSession) SwapRouter1Inch(router1Inch common.Address, router1InchData []byte, desc MetaAggregationRouterSwapDescription, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.SwapRouter1Inch(&_Kyberswapbsc.TransactOpts, router1Inch, router1InchData, desc, clientData)
}

// SwapSimpleMode is a paid mutator transaction binding the contract method 0xa7f5c104.
//
// Solidity: function swapSimpleMode(address caller, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapbsc *KyberswapbscTransactor) SwapSimpleMode(opts *bind.TransactOpts, caller common.Address, desc MetaAggregationRouterSwapDescription, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapbsc.contract.Transact(opts, "swapSimpleMode", caller, desc, executorData, clientData)
}

// SwapSimpleMode is a paid mutator transaction binding the contract method 0xa7f5c104.
//
// Solidity: function swapSimpleMode(address caller, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapbsc *KyberswapbscSession) SwapSimpleMode(caller common.Address, desc MetaAggregationRouterSwapDescription, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.SwapSimpleMode(&_Kyberswapbsc.TransactOpts, caller, desc, executorData, clientData)
}

// SwapSimpleMode is a paid mutator transaction binding the contract method 0xa7f5c104.
//
// Solidity: function swapSimpleMode(address caller, (address,address,address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) returns(uint256 returnAmount, uint256 gasUsed)
func (_Kyberswapbsc *KyberswapbscTransactorSession) SwapSimpleMode(caller common.Address, desc MetaAggregationRouterSwapDescription, executorData []byte, clientData []byte) (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.SwapSimpleMode(&_Kyberswapbsc.TransactOpts, caller, desc, executorData, clientData)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Kyberswapbsc *KyberswapbscTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Kyberswapbsc.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Kyberswapbsc *KyberswapbscSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.TransferOwnership(&_Kyberswapbsc.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Kyberswapbsc *KyberswapbscTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.TransferOwnership(&_Kyberswapbsc.TransactOpts, newOwner)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x0d392cd9.
//
// Solidity: function updateWhitelist(address addr, bool value) returns()
func (_Kyberswapbsc *KyberswapbscTransactor) UpdateWhitelist(opts *bind.TransactOpts, addr common.Address, value bool) (*types.Transaction, error) {
	return _Kyberswapbsc.contract.Transact(opts, "updateWhitelist", addr, value)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x0d392cd9.
//
// Solidity: function updateWhitelist(address addr, bool value) returns()
func (_Kyberswapbsc *KyberswapbscSession) UpdateWhitelist(addr common.Address, value bool) (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.UpdateWhitelist(&_Kyberswapbsc.TransactOpts, addr, value)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x0d392cd9.
//
// Solidity: function updateWhitelist(address addr, bool value) returns()
func (_Kyberswapbsc *KyberswapbscTransactorSession) UpdateWhitelist(addr common.Address, value bool) (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.UpdateWhitelist(&_Kyberswapbsc.TransactOpts, addr, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Kyberswapbsc *KyberswapbscTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kyberswapbsc.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Kyberswapbsc *KyberswapbscSession) Receive() (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.Receive(&_Kyberswapbsc.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Kyberswapbsc *KyberswapbscTransactorSession) Receive() (*types.Transaction, error) {
	return _Kyberswapbsc.Contract.Receive(&_Kyberswapbsc.TransactOpts)
}

// KyberswapbscClientDataIterator is returned from FilterClientData and is used to iterate over the raw logs and unpacked data for ClientData events raised by the Kyberswapbsc contract.
type KyberswapbscClientDataIterator struct {
	Event *KyberswapbscClientData // Event containing the contract specifics and raw log

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
func (it *KyberswapbscClientDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberswapbscClientData)
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
		it.Event = new(KyberswapbscClientData)
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
func (it *KyberswapbscClientDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberswapbscClientDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberswapbscClientData represents a ClientData event raised by the Kyberswapbsc contract.
type KyberswapbscClientData struct {
	ClientData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClientData is a free log retrieval operation binding the contract event 0x095e66fa4dd6a6f7b43fb8444a7bd0edb870508c7abf639bc216efb0bcff9779.
//
// Solidity: event ClientData(bytes clientData)
func (_Kyberswapbsc *KyberswapbscFilterer) FilterClientData(opts *bind.FilterOpts) (*KyberswapbscClientDataIterator, error) {

	logs, sub, err := _Kyberswapbsc.contract.FilterLogs(opts, "ClientData")
	if err != nil {
		return nil, err
	}
	return &KyberswapbscClientDataIterator{contract: _Kyberswapbsc.contract, event: "ClientData", logs: logs, sub: sub}, nil
}

// WatchClientData is a free log subscription operation binding the contract event 0x095e66fa4dd6a6f7b43fb8444a7bd0edb870508c7abf639bc216efb0bcff9779.
//
// Solidity: event ClientData(bytes clientData)
func (_Kyberswapbsc *KyberswapbscFilterer) WatchClientData(opts *bind.WatchOpts, sink chan<- *KyberswapbscClientData) (event.Subscription, error) {

	logs, sub, err := _Kyberswapbsc.contract.WatchLogs(opts, "ClientData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberswapbscClientData)
				if err := _Kyberswapbsc.contract.UnpackLog(event, "ClientData", log); err != nil {
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
func (_Kyberswapbsc *KyberswapbscFilterer) ParseClientData(log types.Log) (*KyberswapbscClientData, error) {
	event := new(KyberswapbscClientData)
	if err := _Kyberswapbsc.contract.UnpackLog(event, "ClientData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberswapbscErrorIterator is returned from FilterError and is used to iterate over the raw logs and unpacked data for Error events raised by the Kyberswapbsc contract.
type KyberswapbscErrorIterator struct {
	Event *KyberswapbscError // Event containing the contract specifics and raw log

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
func (it *KyberswapbscErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberswapbscError)
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
		it.Event = new(KyberswapbscError)
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
func (it *KyberswapbscErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberswapbscErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberswapbscError represents a Error event raised by the Kyberswapbsc contract.
type KyberswapbscError struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterError is a free log retrieval operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_Kyberswapbsc *KyberswapbscFilterer) FilterError(opts *bind.FilterOpts) (*KyberswapbscErrorIterator, error) {

	logs, sub, err := _Kyberswapbsc.contract.FilterLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return &KyberswapbscErrorIterator{contract: _Kyberswapbsc.contract, event: "Error", logs: logs, sub: sub}, nil
}

// WatchError is a free log subscription operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_Kyberswapbsc *KyberswapbscFilterer) WatchError(opts *bind.WatchOpts, sink chan<- *KyberswapbscError) (event.Subscription, error) {

	logs, sub, err := _Kyberswapbsc.contract.WatchLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberswapbscError)
				if err := _Kyberswapbsc.contract.UnpackLog(event, "Error", log); err != nil {
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
func (_Kyberswapbsc *KyberswapbscFilterer) ParseError(log types.Log) (*KyberswapbscError, error) {
	event := new(KyberswapbscError)
	if err := _Kyberswapbsc.contract.UnpackLog(event, "Error", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberswapbscExchangeIterator is returned from FilterExchange and is used to iterate over the raw logs and unpacked data for Exchange events raised by the Kyberswapbsc contract.
type KyberswapbscExchangeIterator struct {
	Event *KyberswapbscExchange // Event containing the contract specifics and raw log

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
func (it *KyberswapbscExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberswapbscExchange)
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
		it.Event = new(KyberswapbscExchange)
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
func (it *KyberswapbscExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberswapbscExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberswapbscExchange represents a Exchange event raised by the Kyberswapbsc contract.
type KyberswapbscExchange struct {
	Pair      common.Address
	AmountOut *big.Int
	Output    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExchange is a free log retrieval operation binding the contract event 0xddac40937f35385a34f721af292e5a83fc5b840f722bff57c2fc71adba708c48.
//
// Solidity: event Exchange(address pair, uint256 amountOut, address output)
func (_Kyberswapbsc *KyberswapbscFilterer) FilterExchange(opts *bind.FilterOpts) (*KyberswapbscExchangeIterator, error) {

	logs, sub, err := _Kyberswapbsc.contract.FilterLogs(opts, "Exchange")
	if err != nil {
		return nil, err
	}
	return &KyberswapbscExchangeIterator{contract: _Kyberswapbsc.contract, event: "Exchange", logs: logs, sub: sub}, nil
}

// WatchExchange is a free log subscription operation binding the contract event 0xddac40937f35385a34f721af292e5a83fc5b840f722bff57c2fc71adba708c48.
//
// Solidity: event Exchange(address pair, uint256 amountOut, address output)
func (_Kyberswapbsc *KyberswapbscFilterer) WatchExchange(opts *bind.WatchOpts, sink chan<- *KyberswapbscExchange) (event.Subscription, error) {

	logs, sub, err := _Kyberswapbsc.contract.WatchLogs(opts, "Exchange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberswapbscExchange)
				if err := _Kyberswapbsc.contract.UnpackLog(event, "Exchange", log); err != nil {
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
func (_Kyberswapbsc *KyberswapbscFilterer) ParseExchange(log types.Log) (*KyberswapbscExchange, error) {
	event := new(KyberswapbscExchange)
	if err := _Kyberswapbsc.contract.UnpackLog(event, "Exchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberswapbscOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Kyberswapbsc contract.
type KyberswapbscOwnershipTransferredIterator struct {
	Event *KyberswapbscOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KyberswapbscOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberswapbscOwnershipTransferred)
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
		it.Event = new(KyberswapbscOwnershipTransferred)
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
func (it *KyberswapbscOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberswapbscOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberswapbscOwnershipTransferred represents a OwnershipTransferred event raised by the Kyberswapbsc contract.
type KyberswapbscOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kyberswapbsc *KyberswapbscFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KyberswapbscOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kyberswapbsc.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KyberswapbscOwnershipTransferredIterator{contract: _Kyberswapbsc.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kyberswapbsc *KyberswapbscFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KyberswapbscOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kyberswapbsc.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberswapbscOwnershipTransferred)
				if err := _Kyberswapbsc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Kyberswapbsc *KyberswapbscFilterer) ParseOwnershipTransferred(log types.Log) (*KyberswapbscOwnershipTransferred, error) {
	event := new(KyberswapbscOwnershipTransferred)
	if err := _Kyberswapbsc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KyberswapbscSwappedIterator is returned from FilterSwapped and is used to iterate over the raw logs and unpacked data for Swapped events raised by the Kyberswapbsc contract.
type KyberswapbscSwappedIterator struct {
	Event *KyberswapbscSwapped // Event containing the contract specifics and raw log

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
func (it *KyberswapbscSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KyberswapbscSwapped)
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
		it.Event = new(KyberswapbscSwapped)
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
func (it *KyberswapbscSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KyberswapbscSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KyberswapbscSwapped represents a Swapped event raised by the Kyberswapbsc contract.
type KyberswapbscSwapped struct {
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
func (_Kyberswapbsc *KyberswapbscFilterer) FilterSwapped(opts *bind.FilterOpts) (*KyberswapbscSwappedIterator, error) {

	logs, sub, err := _Kyberswapbsc.contract.FilterLogs(opts, "Swapped")
	if err != nil {
		return nil, err
	}
	return &KyberswapbscSwappedIterator{contract: _Kyberswapbsc.contract, event: "Swapped", logs: logs, sub: sub}, nil
}

// WatchSwapped is a free log subscription operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_Kyberswapbsc *KyberswapbscFilterer) WatchSwapped(opts *bind.WatchOpts, sink chan<- *KyberswapbscSwapped) (event.Subscription, error) {

	logs, sub, err := _Kyberswapbsc.contract.WatchLogs(opts, "Swapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KyberswapbscSwapped)
				if err := _Kyberswapbsc.contract.UnpackLog(event, "Swapped", log); err != nil {
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
func (_Kyberswapbsc *KyberswapbscFilterer) ParseSwapped(log types.Log) (*KyberswapbscSwapped, error) {
	event := new(KyberswapbscSwapped)
	if err := _Kyberswapbsc.contract.UnpackLog(event, "Swapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
