// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mycontract

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

// MycontractABI is the input ABI used to generate the binding from.
const MycontractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"greet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"_MyEvent\",\"type\":\"event\"}]"

// Mycontract is an auto generated Go binding around an Ethereum contract.
type Mycontract struct {
	MycontractCaller     // Read-only binding to the contract
	MycontractTransactor // Write-only binding to the contract
	MycontractFilterer   // Log filterer for contract events
}

// MycontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MycontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MycontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MycontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MycontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MycontractSession struct {
	Contract     *Mycontract       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MycontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MycontractCallerSession struct {
	Contract *MycontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MycontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MycontractTransactorSession struct {
	Contract     *MycontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MycontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MycontractRaw struct {
	Contract *Mycontract // Generic contract binding to access the raw methods on
}

// MycontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MycontractCallerRaw struct {
	Contract *MycontractCaller // Generic read-only contract binding to access the raw methods on
}

// MycontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MycontractTransactorRaw struct {
	Contract *MycontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMycontract creates a new instance of Mycontract, bound to a specific deployed contract.
func NewMycontract(address common.Address, backend bind.ContractBackend) (*Mycontract, error) {
	contract, err := bindMycontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mycontract{MycontractCaller: MycontractCaller{contract: contract}, MycontractTransactor: MycontractTransactor{contract: contract}, MycontractFilterer: MycontractFilterer{contract: contract}}, nil
}

// NewMycontractCaller creates a new read-only instance of Mycontract, bound to a specific deployed contract.
func NewMycontractCaller(address common.Address, caller bind.ContractCaller) (*MycontractCaller, error) {
	contract, err := bindMycontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MycontractCaller{contract: contract}, nil
}

// NewMycontractTransactor creates a new write-only instance of Mycontract, bound to a specific deployed contract.
func NewMycontractTransactor(address common.Address, transactor bind.ContractTransactor) (*MycontractTransactor, error) {
	contract, err := bindMycontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MycontractTransactor{contract: contract}, nil
}

// NewMycontractFilterer creates a new log filterer instance of Mycontract, bound to a specific deployed contract.
func NewMycontractFilterer(address common.Address, filterer bind.ContractFilterer) (*MycontractFilterer, error) {
	contract, err := bindMycontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MycontractFilterer{contract: contract}, nil
}

// bindMycontract binds a generic wrapper to an already deployed contract.
func bindMycontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MycontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mycontract *MycontractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mycontract.Contract.MycontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mycontract *MycontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.Contract.MycontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mycontract *MycontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mycontract.Contract.MycontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mycontract *MycontractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mycontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mycontract *MycontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mycontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mycontract *MycontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mycontract.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256)
func (_Mycontract *MycontractCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Mycontract.contract.Call(opts, out, "count")
	return *ret0, err
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256)
func (_Mycontract *MycontractSession) Count() (*big.Int, error) {
	return _Mycontract.Contract.Count(&_Mycontract.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() constant returns(uint256)
func (_Mycontract *MycontractCallerSession) Count() (*big.Int, error) {
	return _Mycontract.Contract.Count(&_Mycontract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Mycontract *MycontractCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Mycontract.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Mycontract *MycontractSession) Name() (string, error) {
	return _Mycontract.Contract.Name(&_Mycontract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Mycontract *MycontractCallerSession) Name() (string, error) {
	return _Mycontract.Contract.Name(&_Mycontract.CallOpts)
}

// Greet is a paid mutator transaction binding the contract method 0xead710c4.
//
// Solidity: function greet(_name string) returns()
func (_Mycontract *MycontractTransactor) Greet(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Mycontract.contract.Transact(opts, "greet", _name)
}

// Greet is a paid mutator transaction binding the contract method 0xead710c4.
//
// Solidity: function greet(_name string) returns()
func (_Mycontract *MycontractSession) Greet(_name string) (*types.Transaction, error) {
	return _Mycontract.Contract.Greet(&_Mycontract.TransactOpts, _name)
}

// Greet is a paid mutator transaction binding the contract method 0xead710c4.
//
// Solidity: function greet(_name string) returns()
func (_Mycontract *MycontractTransactorSession) Greet(_name string) (*types.Transaction, error) {
	return _Mycontract.Contract.Greet(&_Mycontract.TransactOpts, _name)
}

// MycontractMyEventIterator is returned from FilterMyEvent and is used to iterate over the raw logs and unpacked data for MyEvent events raised by the Mycontract contract.
type MycontractMyEventIterator struct {
	Event *MycontractMyEvent // Event containing the contract specifics and raw log

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
func (it *MycontractMyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MycontractMyEvent)
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
		it.Event = new(MycontractMyEvent)
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
func (it *MycontractMyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MycontractMyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MycontractMyEvent represents a MyEvent event raised by the Mycontract contract.
type MycontractMyEvent struct {
	Name  string
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMyEvent is a free log retrieval operation binding the contract event 0x58a511ff45ed07d91d943b897d293a3bf9504ac55ca968e356fe4f1d3b7f23b5.
//
// Solidity: event _MyEvent(name string, count uint256)
func (_Mycontract *MycontractFilterer) FilterMyEvent(opts *bind.FilterOpts) (*MycontractMyEventIterator, error) {

	logs, sub, err := _Mycontract.contract.FilterLogs(opts, "_MyEvent")
	if err != nil {
		return nil, err
	}
	return &MycontractMyEventIterator{contract: _Mycontract.contract, event: "_MyEvent", logs: logs, sub: sub}, nil
}

// WatchMyEvent is a free log subscription operation binding the contract event 0x58a511ff45ed07d91d943b897d293a3bf9504ac55ca968e356fe4f1d3b7f23b5.
//
// Solidity: event _MyEvent(name string, count uint256)
func (_Mycontract *MycontractFilterer) WatchMyEvent(opts *bind.WatchOpts, sink chan<- *MycontractMyEvent) (event.Subscription, error) {

	logs, sub, err := _Mycontract.contract.WatchLogs(opts, "_MyEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MycontractMyEvent)
				if err := _Mycontract.contract.UnpackLog(event, "_MyEvent", log); err != nil {
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
