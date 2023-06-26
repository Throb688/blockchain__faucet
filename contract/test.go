// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TestABI is the input ABI used to generate the binding from.
const TestABI = "[{\"inputs\":[],\"name\":\"amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bb\",\"type\":\"address\"}],\"name\":\"getBalanceof\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"num\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payme\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quantity_everyday\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_num1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_num2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_num3\",\"type\":\"uint256\"}],\"name\":\"reset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_num1\",\"type\":\"uint256\"}],\"name\":\"setAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_num2\",\"type\":\"uint256\"}],\"name\":\"setNum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_num\",\"type\":\"uint256\"}],\"name\":\"setQuantity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aa\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_ip\",\"type\":\"string\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Test is an auto generated Go binding around an Ethereum contract.
type Test struct {
	TestCaller     // Read-only binding to the contract
	TestTransactor // Write-only binding to the contract
	TestFilterer   // Log filterer for contract events
}

// TestCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestSession struct {
	Contract     *Test             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestCallerSession struct {
	Contract *TestCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestTransactorSession struct {
	Contract     *TestTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestRaw struct {
	Contract *Test // Generic contract binding to access the raw methods on
}

// TestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestCallerRaw struct {
	Contract *TestCaller // Generic read-only contract binding to access the raw methods on
}

// TestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestTransactorRaw struct {
	Contract *TestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTest creates a new instance of Test, bound to a specific deployed contract.
func NewTest(address common.Address, backend bind.ContractBackend) (*Test, error) {
	contract, err := bindTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}, TestFilterer: TestFilterer{contract: contract}}, nil
}

// NewTestCaller creates a new read-only instance of Test, bound to a specific deployed contract.
func NewTestCaller(address common.Address, caller bind.ContractCaller) (*TestCaller, error) {
	contract, err := bindTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestCaller{contract: contract}, nil
}

// NewTestTransactor creates a new write-only instance of Test, bound to a specific deployed contract.
func NewTestTransactor(address common.Address, transactor bind.ContractTransactor) (*TestTransactor, error) {
	contract, err := bindTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestTransactor{contract: contract}, nil
}

// NewTestFilterer creates a new log filterer instance of Test, bound to a specific deployed contract.
func NewTestFilterer(address common.Address, filterer bind.ContractFilterer) (*TestFilterer, error) {
	contract, err := bindTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestFilterer{contract: contract}, nil
}

// bindTest binds a generic wrapper to an already deployed contract.
func bindTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test.Contract.TestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.contract.Transact(opts, method, params...)
}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_Test *TestCaller) Amount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "amount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_Test *TestSession) Amount() (*big.Int, error) {
	return _Test.Contract.Amount(&_Test.CallOpts)
}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_Test *TestCallerSession) Amount() (*big.Int, error) {
	return _Test.Contract.Amount(&_Test.CallOpts)
}

// GetBalanceof is a free data retrieval call binding the contract method 0xaacaa074.
//
// Solidity: function getBalanceof(address bb) view returns(uint256)
func (_Test *TestCaller) GetBalanceof(opts *bind.CallOpts, bb common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "getBalanceof", bb)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceof is a free data retrieval call binding the contract method 0xaacaa074.
//
// Solidity: function getBalanceof(address bb) view returns(uint256)
func (_Test *TestSession) GetBalanceof(bb common.Address) (*big.Int, error) {
	return _Test.Contract.GetBalanceof(&_Test.CallOpts, bb)
}

// GetBalanceof is a free data retrieval call binding the contract method 0xaacaa074.
//
// Solidity: function getBalanceof(address bb) view returns(uint256)
func (_Test *TestCallerSession) GetBalanceof(bb common.Address) (*big.Int, error) {
	return _Test.Contract.GetBalanceof(&_Test.CallOpts, bb)
}

// Num is a free data retrieval call binding the contract method 0x4e70b1dc.
//
// Solidity: function num() view returns(uint256)
func (_Test *TestCaller) Num(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "num")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Num is a free data retrieval call binding the contract method 0x4e70b1dc.
//
// Solidity: function num() view returns(uint256)
func (_Test *TestSession) Num() (*big.Int, error) {
	return _Test.Contract.Num(&_Test.CallOpts)
}

// Num is a free data retrieval call binding the contract method 0x4e70b1dc.
//
// Solidity: function num() view returns(uint256)
func (_Test *TestCallerSession) Num() (*big.Int, error) {
	return _Test.Contract.Num(&_Test.CallOpts)
}

// QuantityEveryday is a free data retrieval call binding the contract method 0xf6d73d46.
//
// Solidity: function quantity_everyday() view returns(uint256)
func (_Test *TestCaller) QuantityEveryday(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "quantity_everyday")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuantityEveryday is a free data retrieval call binding the contract method 0xf6d73d46.
//
// Solidity: function quantity_everyday() view returns(uint256)
func (_Test *TestSession) QuantityEveryday() (*big.Int, error) {
	return _Test.Contract.QuantityEveryday(&_Test.CallOpts)
}

// QuantityEveryday is a free data retrieval call binding the contract method 0xf6d73d46.
//
// Solidity: function quantity_everyday() view returns(uint256)
func (_Test *TestCallerSession) QuantityEveryday() (*big.Int, error) {
	return _Test.Contract.QuantityEveryday(&_Test.CallOpts)
}

// Payme is a paid mutator transaction binding the contract method 0xba6361fb.
//
// Solidity: function payme() payable returns()
func (_Test *TestTransactor) Payme(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "payme")
}

// Payme is a paid mutator transaction binding the contract method 0xba6361fb.
//
// Solidity: function payme() payable returns()
func (_Test *TestSession) Payme() (*types.Transaction, error) {
	return _Test.Contract.Payme(&_Test.TransactOpts)
}

// Payme is a paid mutator transaction binding the contract method 0xba6361fb.
//
// Solidity: function payme() payable returns()
func (_Test *TestTransactorSession) Payme() (*types.Transaction, error) {
	return _Test.Contract.Payme(&_Test.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xa6801cbd.
//
// Solidity: function reset(uint256 _num1, uint256 _num2, uint256 _num3) returns()
func (_Test *TestTransactor) Reset(opts *bind.TransactOpts, _num1 *big.Int, _num2 *big.Int, _num3 *big.Int) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "reset", _num1, _num2, _num3)
}

// Reset is a paid mutator transaction binding the contract method 0xa6801cbd.
//
// Solidity: function reset(uint256 _num1, uint256 _num2, uint256 _num3) returns()
func (_Test *TestSession) Reset(_num1 *big.Int, _num2 *big.Int, _num3 *big.Int) (*types.Transaction, error) {
	return _Test.Contract.Reset(&_Test.TransactOpts, _num1, _num2, _num3)
}

// Reset is a paid mutator transaction binding the contract method 0xa6801cbd.
//
// Solidity: function reset(uint256 _num1, uint256 _num2, uint256 _num3) returns()
func (_Test *TestTransactorSession) Reset(_num1 *big.Int, _num2 *big.Int, _num3 *big.Int) (*types.Transaction, error) {
	return _Test.Contract.Reset(&_Test.TransactOpts, _num1, _num2, _num3)
}

// SetAmount is a paid mutator transaction binding the contract method 0x271f88b4.
//
// Solidity: function setAmount(uint256 _num1) returns()
func (_Test *TestTransactor) SetAmount(opts *bind.TransactOpts, _num1 *big.Int) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "setAmount", _num1)
}

// SetAmount is a paid mutator transaction binding the contract method 0x271f88b4.
//
// Solidity: function setAmount(uint256 _num1) returns()
func (_Test *TestSession) SetAmount(_num1 *big.Int) (*types.Transaction, error) {
	return _Test.Contract.SetAmount(&_Test.TransactOpts, _num1)
}

// SetAmount is a paid mutator transaction binding the contract method 0x271f88b4.
//
// Solidity: function setAmount(uint256 _num1) returns()
func (_Test *TestTransactorSession) SetAmount(_num1 *big.Int) (*types.Transaction, error) {
	return _Test.Contract.SetAmount(&_Test.TransactOpts, _num1)
}

// SetNum is a paid mutator transaction binding the contract method 0xcd16ecbf.
//
// Solidity: function setNum(uint256 _num2) returns()
func (_Test *TestTransactor) SetNum(opts *bind.TransactOpts, _num2 *big.Int) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "setNum", _num2)
}

// SetNum is a paid mutator transaction binding the contract method 0xcd16ecbf.
//
// Solidity: function setNum(uint256 _num2) returns()
func (_Test *TestSession) SetNum(_num2 *big.Int) (*types.Transaction, error) {
	return _Test.Contract.SetNum(&_Test.TransactOpts, _num2)
}

// SetNum is a paid mutator transaction binding the contract method 0xcd16ecbf.
//
// Solidity: function setNum(uint256 _num2) returns()
func (_Test *TestTransactorSession) SetNum(_num2 *big.Int) (*types.Transaction, error) {
	return _Test.Contract.SetNum(&_Test.TransactOpts, _num2)
}

// SetQuantity is a paid mutator transaction binding the contract method 0x0bf15326.
//
// Solidity: function setQuantity(uint256 _num) returns()
func (_Test *TestTransactor) SetQuantity(opts *bind.TransactOpts, _num *big.Int) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "setQuantity", _num)
}

// SetQuantity is a paid mutator transaction binding the contract method 0x0bf15326.
//
// Solidity: function setQuantity(uint256 _num) returns()
func (_Test *TestSession) SetQuantity(_num *big.Int) (*types.Transaction, error) {
	return _Test.Contract.SetQuantity(&_Test.TransactOpts, _num)
}

// SetQuantity is a paid mutator transaction binding the contract method 0x0bf15326.
//
// Solidity: function setQuantity(uint256 _num) returns()
func (_Test *TestTransactorSession) SetQuantity(_num *big.Int) (*types.Transaction, error) {
	return _Test.Contract.SetQuantity(&_Test.TransactOpts, _num)
}

// Withdraw is a paid mutator transaction binding the contract method 0x6e6b1101.
//
// Solidity: function withdraw(address aa, string _ip) returns()
func (_Test *TestTransactor) Withdraw(opts *bind.TransactOpts, aa common.Address, _ip string) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "withdraw", aa, _ip)
}

// Withdraw is a paid mutator transaction binding the contract method 0x6e6b1101.
//
// Solidity: function withdraw(address aa, string _ip) returns()
func (_Test *TestSession) Withdraw(aa common.Address, _ip string) (*types.Transaction, error) {
	return _Test.Contract.Withdraw(&_Test.TransactOpts, aa, _ip)
}

// Withdraw is a paid mutator transaction binding the contract method 0x6e6b1101.
//
// Solidity: function withdraw(address aa, string _ip) returns()
func (_Test *TestTransactorSession) Withdraw(aa common.Address, _ip string) (*types.Transaction, error) {
	return _Test.Contract.Withdraw(&_Test.TransactOpts, aa, _ip)
}
