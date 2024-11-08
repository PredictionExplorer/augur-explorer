// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.ConvertType
)

// IPrizesWalletMetaData contains all meta data concerning the IPrizesWallet contract.
var IPrizesWalletMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner_\",\"type\":\"address\"}],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner_\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPrizesWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use IPrizesWalletMetaData.ABI instead.
var IPrizesWalletABI = IPrizesWalletMetaData.ABI

// IPrizesWallet is an auto generated Go binding around an Ethereum contract.
type IPrizesWallet struct {
	IPrizesWalletCaller     // Read-only binding to the contract
	IPrizesWalletTransactor // Write-only binding to the contract
	IPrizesWalletFilterer   // Log filterer for contract events
}

// IPrizesWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPrizesWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPrizesWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPrizesWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPrizesWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPrizesWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPrizesWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPrizesWalletSession struct {
	Contract     *IPrizesWallet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPrizesWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPrizesWalletCallerSession struct {
	Contract *IPrizesWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IPrizesWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPrizesWalletTransactorSession struct {
	Contract     *IPrizesWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IPrizesWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPrizesWalletRaw struct {
	Contract *IPrizesWallet // Generic contract binding to access the raw methods on
}

// IPrizesWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPrizesWalletCallerRaw struct {
	Contract *IPrizesWalletCaller // Generic read-only contract binding to access the raw methods on
}

// IPrizesWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPrizesWalletTransactorRaw struct {
	Contract *IPrizesWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPrizesWallet creates a new instance of IPrizesWallet, bound to a specific deployed contract.
func NewIPrizesWallet(address common.Address, backend bind.ContractBackend) (*IPrizesWallet, error) {
	contract, err := bindIPrizesWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPrizesWallet{IPrizesWalletCaller: IPrizesWalletCaller{contract: contract}, IPrizesWalletTransactor: IPrizesWalletTransactor{contract: contract}, IPrizesWalletFilterer: IPrizesWalletFilterer{contract: contract}}, nil
}

// NewIPrizesWalletCaller creates a new read-only instance of IPrizesWallet, bound to a specific deployed contract.
func NewIPrizesWalletCaller(address common.Address, caller bind.ContractCaller) (*IPrizesWalletCaller, error) {
	contract, err := bindIPrizesWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletCaller{contract: contract}, nil
}

// NewIPrizesWalletTransactor creates a new write-only instance of IPrizesWallet, bound to a specific deployed contract.
func NewIPrizesWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*IPrizesWalletTransactor, error) {
	contract, err := bindIPrizesWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletTransactor{contract: contract}, nil
}

// NewIPrizesWalletFilterer creates a new log filterer instance of IPrizesWallet, bound to a specific deployed contract.
func NewIPrizesWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*IPrizesWalletFilterer, error) {
	contract, err := bindIPrizesWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletFilterer{contract: contract}, nil
}

// bindIPrizesWallet binds a generic wrapper to an already deployed contract.
func bindIPrizesWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPrizesWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPrizesWallet *IPrizesWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPrizesWallet.Contract.IPrizesWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPrizesWallet *IPrizesWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.IPrizesWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPrizesWallet *IPrizesWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.IPrizesWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPrizesWallet *IPrizesWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPrizesWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPrizesWallet *IPrizesWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPrizesWallet *IPrizesWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.contract.Transact(opts, method, params...)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address winner_) view returns(uint256)
func (_IPrizesWallet *IPrizesWalletCaller) GetEthBalance(opts *bind.CallOpts, winner_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPrizesWallet.contract.Call(opts, &out, "getEthBalance", winner_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address winner_) view returns(uint256)
func (_IPrizesWallet *IPrizesWalletSession) GetEthBalance(winner_ common.Address) (*big.Int, error) {
	return _IPrizesWallet.Contract.GetEthBalance(&_IPrizesWallet.CallOpts, winner_)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address winner_) view returns(uint256)
func (_IPrizesWallet *IPrizesWalletCallerSession) GetEthBalance(winner_ common.Address) (*big.Int, error) {
	return _IPrizesWallet.Contract.GetEthBalance(&_IPrizesWallet.CallOpts, winner_)
}

// GetEthBalance0 is a free data retrieval call binding the contract method 0x70ed0ada.
//
// Solidity: function getEthBalance() view returns(uint256)
func (_IPrizesWallet *IPrizesWalletCaller) GetEthBalance0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IPrizesWallet.contract.Call(opts, &out, "getEthBalance0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance0 is a free data retrieval call binding the contract method 0x70ed0ada.
//
// Solidity: function getEthBalance() view returns(uint256)
func (_IPrizesWallet *IPrizesWalletSession) GetEthBalance0() (*big.Int, error) {
	return _IPrizesWallet.Contract.GetEthBalance0(&_IPrizesWallet.CallOpts)
}

// GetEthBalance0 is a free data retrieval call binding the contract method 0x70ed0ada.
//
// Solidity: function getEthBalance() view returns(uint256)
func (_IPrizesWallet *IPrizesWalletCallerSession) GetEthBalance0() (*big.Int, error) {
	return _IPrizesWallet.Contract.GetEthBalance0(&_IPrizesWallet.CallOpts)
}

// DepositEth is a paid mutator transaction binding the contract method 0xad9d4ba3.
//
// Solidity: function depositEth(address winner_) payable returns()
func (_IPrizesWallet *IPrizesWalletTransactor) DepositEth(opts *bind.TransactOpts, winner_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "depositEth", winner_)
}

// DepositEth is a paid mutator transaction binding the contract method 0xad9d4ba3.
//
// Solidity: function depositEth(address winner_) payable returns()
func (_IPrizesWallet *IPrizesWalletSession) DepositEth(winner_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.DepositEth(&_IPrizesWallet.TransactOpts, winner_)
}

// DepositEth is a paid mutator transaction binding the contract method 0xad9d4ba3.
//
// Solidity: function depositEth(address winner_) payable returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) DepositEth(winner_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.DepositEth(&_IPrizesWallet.TransactOpts, winner_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_IPrizesWallet *IPrizesWalletTransactor) WithdrawEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "withdrawEth")
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_IPrizesWallet *IPrizesWalletSession) WithdrawEth() (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEth(&_IPrizesWallet.TransactOpts)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) WithdrawEth() (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEth(&_IPrizesWallet.TransactOpts)
}

// IPrizesWalletEthReceivedIterator is returned from FilterEthReceived and is used to iterate over the raw logs and unpacked data for EthReceived events raised by the IPrizesWallet contract.
type IPrizesWalletEthReceivedIterator struct {
	Event *IPrizesWalletEthReceived // Event containing the contract specifics and raw log

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
func (it *IPrizesWalletEthReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPrizesWalletEthReceived)
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
		it.Event = new(IPrizesWalletEthReceived)
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
func (it *IPrizesWalletEthReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPrizesWalletEthReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPrizesWalletEthReceived represents a EthReceived event raised by the IPrizesWallet contract.
type IPrizesWalletEthReceived struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEthReceived is a free log retrieval operation binding the contract event 0x85177f287940f2f05425a4029951af0e047a7f9c4eaa9a6e6917bcd869f86695.
//
// Solidity: event EthReceived(address indexed winner, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterEthReceived(opts *bind.FilterOpts, winner []common.Address) (*IPrizesWalletEthReceivedIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "EthReceived", winnerRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletEthReceivedIterator{contract: _IPrizesWallet.contract, event: "EthReceived", logs: logs, sub: sub}, nil
}

// WatchEthReceived is a free log subscription operation binding the contract event 0x85177f287940f2f05425a4029951af0e047a7f9c4eaa9a6e6917bcd869f86695.
//
// Solidity: event EthReceived(address indexed winner, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchEthReceived(opts *bind.WatchOpts, sink chan<- *IPrizesWalletEthReceived, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "EthReceived", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPrizesWalletEthReceived)
				if err := _IPrizesWallet.contract.UnpackLog(event, "EthReceived", log); err != nil {
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

// ParseEthReceived is a log parse operation binding the contract event 0x85177f287940f2f05425a4029951af0e047a7f9c4eaa9a6e6917bcd869f86695.
//
// Solidity: event EthReceived(address indexed winner, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) ParseEthReceived(log types.Log) (*IPrizesWalletEthReceived, error) {
	event := new(IPrizesWalletEthReceived)
	if err := _IPrizesWallet.contract.UnpackLog(event, "EthReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPrizesWalletEthWithdrawnIterator is returned from FilterEthWithdrawn and is used to iterate over the raw logs and unpacked data for EthWithdrawn events raised by the IPrizesWallet contract.
type IPrizesWalletEthWithdrawnIterator struct {
	Event *IPrizesWalletEthWithdrawn // Event containing the contract specifics and raw log

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
func (it *IPrizesWalletEthWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPrizesWalletEthWithdrawn)
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
		it.Event = new(IPrizesWalletEthWithdrawn)
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
func (it *IPrizesWalletEthWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPrizesWalletEthWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPrizesWalletEthWithdrawn represents a EthWithdrawn event raised by the IPrizesWallet contract.
type IPrizesWalletEthWithdrawn struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawn is a free log retrieval operation binding the contract event 0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b.
//
// Solidity: event EthWithdrawn(address indexed winner, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterEthWithdrawn(opts *bind.FilterOpts, winner []common.Address) (*IPrizesWalletEthWithdrawnIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "EthWithdrawn", winnerRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletEthWithdrawnIterator{contract: _IPrizesWallet.contract, event: "EthWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawn is a free log subscription operation binding the contract event 0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b.
//
// Solidity: event EthWithdrawn(address indexed winner, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchEthWithdrawn(opts *bind.WatchOpts, sink chan<- *IPrizesWalletEthWithdrawn, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "EthWithdrawn", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPrizesWalletEthWithdrawn)
				if err := _IPrizesWallet.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
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

// ParseEthWithdrawn is a log parse operation binding the contract event 0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b.
//
// Solidity: event EthWithdrawn(address indexed winner, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) ParseEthWithdrawn(log types.Log) (*IPrizesWalletEthWithdrawn, error) {
	event := new(IPrizesWalletEthWithdrawn)
	if err := _IPrizesWallet.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizesWalletMetaData contains all meta data concerning the PrizesWallet contract.
var PrizesWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"game_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner_\",\"type\":\"address\"}],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner_\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523461002f576100196100146100f4565b610246565b610021610034565b610856610278823961085690f35b61003a565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100669061003e565b810190811060018060401b0382111761007e57604052565b610048565b9061009661008f610034565b928361005c565b565b5f80fd5b60018060a01b031690565b6100b09061009c565b90565b6100bc816100a7565b036100c357565b5f80fd5b905051906100d4826100b3565b565b906020828203126100ef576100ec915f016100c7565b90565b610098565b610112610ace8038038061010781610083565b9283398101906100d6565b90565b90565b90565b61012f61012a61013492610115565b610118565b61009c565b90565b6101409061011b565b90565b60209181520190565b5f7f5a65726f2d616464726573732077617320676976656e2e000000000000000000910152565b6101806017602092610143565b6101898161014c565b0190565b6101a29060208101905f818303910152610173565b90565b156101ac57565b6101b4610034565b63eac0d38960e01b8152806101cb6004820161018d565b0390fd5b5f1b90565b906101e560018060a01b03916101cf565b9181191691161790565b6102036101fe6102089261009c565b610118565b61009c565b90565b610214906101ef565b90565b6102209061020b565b90565b90565b9061023b61023661024292610217565b610223565b82546101d4565b9055565b6102759061026f8161026861026261025d5f610137565b6100a7565b916100a7565b14156101a5565b5f610226565b56fe60806040526004361015610013575b610271565b61001d5f3561006c565b80634d2301cc1461006757806370ed0ada14610062578063a0ef91df1461005d578063ad9d4ba3146100585763c3fe3e280361000e5761023c565b6101ae565b61017b565b610141565b6100fd565b60e01c90565b60405190565b5f80fd5b5f80fd5b60018060a01b031690565b61009490610080565b90565b6100a08161008b565b036100a757565b5f80fd5b905035906100b882610097565b565b906020828203126100d3576100d0915f016100ab565b90565b61007c565b90565b6100e4906100d8565b9052565b91906100fb905f602085019401906100db565b565b3461012d576101296101186101133660046100ba565b610315565b610120610072565b918291826100e8565b0390f35b610078565b5f91031261013c57565b61007c565b3461017157610151366004610132565b61016d61015c61033f565b610164610072565b918291826100e8565b0390f35b610078565b5f0190565b346101a95761018b366004610132565b610193610598565b61019b610072565b806101a581610176565b0390f35b610078565b6101c16101bc3660046100ba565b610815565b6101c9610072565b806101d381610176565b0390f35b1c90565b60018060a01b031690565b6101f69060086101fb93026101d7565b6101db565b90565b9061020991546101e6565b90565b6102175f5f906101fe565b90565b6102239061008b565b9052565b919061023a905f6020850194019061021a565b565b3461026c5761024c366004610132565b61026861025761020c565b61025f610072565b91829182610227565b0390f35b610078565b5f80fd5b5f90565b90565b61029061028b61029592610080565b610279565b610080565b90565b6102a19061027c565b90565b634e487b7160e01b5f52603260045260245ffd5b50600160a01b90565b90565b6102cd816102b8565b8210156102e7576102df6001916102c1565b910201905f90565b6102a4565b90565b6102ff90600861030493026101d7565b6102ec565b90565b9061031291546102ef565b90565b61033661033c91610324610275565b50610330600191610298565b906102c4565b90610307565b90565b610347610275565b5061036561035f600161035933610298565b906102c4565b90610307565b90565b90565b61037f61037a61038492610368565b610279565b6100d8565b90565b1b90565b919060086103a69102916103a05f1984610387565b92610387565b9181191691161790565b6103c46103bf6103c9926100d8565b610279565b6100d8565b90565b90565b91906103e56103e06103ed936103b0565b6103cc565b90835461038b565b9055565b6103fa9061027c565b90565b610406906103f1565b90565b905090565b6104195f8092610409565b0190565b6104269061040e565b90565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b9061045190610429565b810190811067ffffffffffffffff82111761046b57604052565b610433565b9061048361047c610072565b9283610447565b565b67ffffffffffffffff81116104a35761049f602091610429565b0190565b610433565b906104ba6104b583610485565b610470565b918252565b606090565b3d5f146104df576104d43d6104a8565b903d5f602084013e5b565b6104e76104bf565b906104dd565b60209181520190565b5f7f455448207769746864726177616c206661696c65642e00000000000000000000910152565b61052a60166020926104ed565b610533816104f6565b0190565b916040610568929493610561610556606083018381035f85015261051d565b96602083019061021a565b01906100db565b565b15610573575050565b61059461057e610072565b928392630aa7db6360e11b845260048401610537565b0390fd5b6106506105b86105b260016105ac33610298565b906102c4565b90610307565b6105de6105c45f61036b565b6105d860016105d233610298565b906102c4565b906103cf565b33819061062061060e7f8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b926103fd565b92610617610072565b918291826100e8565b0390a25f80338361062f610072565b908161063a8161041d565b03925af16106466104c4565b509033909161056a565b565b5f1c90565b61066361066891610652565b6101db565b90565b6106759054610657565b90565b60207f7065726d697474656420746f2063616c6c2074686973206d6574686f642e0000917f4f6e6c792074686520436f736d696347616d6520636f6e7472616374206973205f8201520152565b6106d2603e6040926104ed565b6106db81610678565b0190565b91906107029060206106fa604086018681035f8801526106c5565b94019061021a565b565b1561070c5750565b61072e90610718610072565b91829163336e7de360e11b8352600483016106df565b0390fd5b6107619061075c3361075461074e6107495f61066b565b61008b565b9161008b565b143390610704565b61079c565b565b634e487b7160e01b5f52601160045260245ffd5b61078661078c919392936100d8565b926100d8565b820180921161079757565b610763565b6107cf346107c96107b760016107b186610298565b906102c4565b9190926107c48385610307565b610777565b916103cf565b34906108106107fe7f85177f287940f2f05425a4029951af0e047a7f9c4eaa9a6e6917bcd869f86695926103fd565b92610807610072565b918291826100e8565b0390a2565b61081e90610732565b56fea26469706673582212207a46aa207cda90d0a4e24471277855864bf9f5684aa4a1e9611196fa5913ba7164736f6c634300081b0033",
}

// PrizesWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use PrizesWalletMetaData.ABI instead.
var PrizesWalletABI = PrizesWalletMetaData.ABI

// PrizesWalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PrizesWalletMetaData.Bin instead.
var PrizesWalletBin = PrizesWalletMetaData.Bin

// DeployPrizesWallet deploys a new Ethereum contract, binding an instance of PrizesWallet to it.
func DeployPrizesWallet(auth *bind.TransactOpts, backend bind.ContractBackend, game_ common.Address) (common.Address, *types.Transaction, *PrizesWallet, error) {
	parsed, err := PrizesWalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PrizesWalletBin), backend, game_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PrizesWallet{PrizesWalletCaller: PrizesWalletCaller{contract: contract}, PrizesWalletTransactor: PrizesWalletTransactor{contract: contract}, PrizesWalletFilterer: PrizesWalletFilterer{contract: contract}}, nil
}

// PrizesWallet is an auto generated Go binding around an Ethereum contract.
type PrizesWallet struct {
	PrizesWalletCaller     // Read-only binding to the contract
	PrizesWalletTransactor // Write-only binding to the contract
	PrizesWalletFilterer   // Log filterer for contract events
}

// PrizesWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrizesWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrizesWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrizesWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrizesWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrizesWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrizesWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrizesWalletSession struct {
	Contract     *PrizesWallet     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrizesWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrizesWalletCallerSession struct {
	Contract *PrizesWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PrizesWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrizesWalletTransactorSession struct {
	Contract     *PrizesWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PrizesWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrizesWalletRaw struct {
	Contract *PrizesWallet // Generic contract binding to access the raw methods on
}

// PrizesWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrizesWalletCallerRaw struct {
	Contract *PrizesWalletCaller // Generic read-only contract binding to access the raw methods on
}

// PrizesWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrizesWalletTransactorRaw struct {
	Contract *PrizesWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrizesWallet creates a new instance of PrizesWallet, bound to a specific deployed contract.
func NewPrizesWallet(address common.Address, backend bind.ContractBackend) (*PrizesWallet, error) {
	contract, err := bindPrizesWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrizesWallet{PrizesWalletCaller: PrizesWalletCaller{contract: contract}, PrizesWalletTransactor: PrizesWalletTransactor{contract: contract}, PrizesWalletFilterer: PrizesWalletFilterer{contract: contract}}, nil
}

// NewPrizesWalletCaller creates a new read-only instance of PrizesWallet, bound to a specific deployed contract.
func NewPrizesWalletCaller(address common.Address, caller bind.ContractCaller) (*PrizesWalletCaller, error) {
	contract, err := bindPrizesWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletCaller{contract: contract}, nil
}

// NewPrizesWalletTransactor creates a new write-only instance of PrizesWallet, bound to a specific deployed contract.
func NewPrizesWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*PrizesWalletTransactor, error) {
	contract, err := bindPrizesWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletTransactor{contract: contract}, nil
}

// NewPrizesWalletFilterer creates a new log filterer instance of PrizesWallet, bound to a specific deployed contract.
func NewPrizesWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*PrizesWalletFilterer, error) {
	contract, err := bindPrizesWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletFilterer{contract: contract}, nil
}

// bindPrizesWallet binds a generic wrapper to an already deployed contract.
func bindPrizesWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PrizesWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrizesWallet *PrizesWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrizesWallet.Contract.PrizesWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrizesWallet *PrizesWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrizesWallet.Contract.PrizesWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrizesWallet *PrizesWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrizesWallet.Contract.PrizesWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrizesWallet *PrizesWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrizesWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrizesWallet *PrizesWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrizesWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrizesWallet *PrizesWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrizesWallet.Contract.contract.Transact(opts, method, params...)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_PrizesWallet *PrizesWalletCaller) Game(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "game")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_PrizesWallet *PrizesWalletSession) Game() (common.Address, error) {
	return _PrizesWallet.Contract.Game(&_PrizesWallet.CallOpts)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_PrizesWallet *PrizesWalletCallerSession) Game() (common.Address, error) {
	return _PrizesWallet.Contract.Game(&_PrizesWallet.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address winner_) view returns(uint256)
func (_PrizesWallet *PrizesWalletCaller) GetEthBalance(opts *bind.CallOpts, winner_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "getEthBalance", winner_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address winner_) view returns(uint256)
func (_PrizesWallet *PrizesWalletSession) GetEthBalance(winner_ common.Address) (*big.Int, error) {
	return _PrizesWallet.Contract.GetEthBalance(&_PrizesWallet.CallOpts, winner_)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address winner_) view returns(uint256)
func (_PrizesWallet *PrizesWalletCallerSession) GetEthBalance(winner_ common.Address) (*big.Int, error) {
	return _PrizesWallet.Contract.GetEthBalance(&_PrizesWallet.CallOpts, winner_)
}

// GetEthBalance0 is a free data retrieval call binding the contract method 0x70ed0ada.
//
// Solidity: function getEthBalance() view returns(uint256)
func (_PrizesWallet *PrizesWalletCaller) GetEthBalance0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "getEthBalance0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance0 is a free data retrieval call binding the contract method 0x70ed0ada.
//
// Solidity: function getEthBalance() view returns(uint256)
func (_PrizesWallet *PrizesWalletSession) GetEthBalance0() (*big.Int, error) {
	return _PrizesWallet.Contract.GetEthBalance0(&_PrizesWallet.CallOpts)
}

// GetEthBalance0 is a free data retrieval call binding the contract method 0x70ed0ada.
//
// Solidity: function getEthBalance() view returns(uint256)
func (_PrizesWallet *PrizesWalletCallerSession) GetEthBalance0() (*big.Int, error) {
	return _PrizesWallet.Contract.GetEthBalance0(&_PrizesWallet.CallOpts)
}

// DepositEth is a paid mutator transaction binding the contract method 0xad9d4ba3.
//
// Solidity: function depositEth(address winner_) payable returns()
func (_PrizesWallet *PrizesWalletTransactor) DepositEth(opts *bind.TransactOpts, winner_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "depositEth", winner_)
}

// DepositEth is a paid mutator transaction binding the contract method 0xad9d4ba3.
//
// Solidity: function depositEth(address winner_) payable returns()
func (_PrizesWallet *PrizesWalletSession) DepositEth(winner_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.DepositEth(&_PrizesWallet.TransactOpts, winner_)
}

// DepositEth is a paid mutator transaction binding the contract method 0xad9d4ba3.
//
// Solidity: function depositEth(address winner_) payable returns()
func (_PrizesWallet *PrizesWalletTransactorSession) DepositEth(winner_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.DepositEth(&_PrizesWallet.TransactOpts, winner_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_PrizesWallet *PrizesWalletTransactor) WithdrawEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "withdrawEth")
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_PrizesWallet *PrizesWalletSession) WithdrawEth() (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEth(&_PrizesWallet.TransactOpts)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_PrizesWallet *PrizesWalletTransactorSession) WithdrawEth() (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEth(&_PrizesWallet.TransactOpts)
}

// PrizesWalletEthReceivedIterator is returned from FilterEthReceived and is used to iterate over the raw logs and unpacked data for EthReceived events raised by the PrizesWallet contract.
type PrizesWalletEthReceivedIterator struct {
	Event *PrizesWalletEthReceived // Event containing the contract specifics and raw log

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
func (it *PrizesWalletEthReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizesWalletEthReceived)
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
		it.Event = new(PrizesWalletEthReceived)
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
func (it *PrizesWalletEthReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizesWalletEthReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizesWalletEthReceived represents a EthReceived event raised by the PrizesWallet contract.
type PrizesWalletEthReceived struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEthReceived is a free log retrieval operation binding the contract event 0x85177f287940f2f05425a4029951af0e047a7f9c4eaa9a6e6917bcd869f86695.
//
// Solidity: event EthReceived(address indexed winner, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) FilterEthReceived(opts *bind.FilterOpts, winner []common.Address) (*PrizesWalletEthReceivedIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "EthReceived", winnerRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletEthReceivedIterator{contract: _PrizesWallet.contract, event: "EthReceived", logs: logs, sub: sub}, nil
}

// WatchEthReceived is a free log subscription operation binding the contract event 0x85177f287940f2f05425a4029951af0e047a7f9c4eaa9a6e6917bcd869f86695.
//
// Solidity: event EthReceived(address indexed winner, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) WatchEthReceived(opts *bind.WatchOpts, sink chan<- *PrizesWalletEthReceived, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "EthReceived", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizesWalletEthReceived)
				if err := _PrizesWallet.contract.UnpackLog(event, "EthReceived", log); err != nil {
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

// ParseEthReceived is a log parse operation binding the contract event 0x85177f287940f2f05425a4029951af0e047a7f9c4eaa9a6e6917bcd869f86695.
//
// Solidity: event EthReceived(address indexed winner, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) ParseEthReceived(log types.Log) (*PrizesWalletEthReceived, error) {
	event := new(PrizesWalletEthReceived)
	if err := _PrizesWallet.contract.UnpackLog(event, "EthReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizesWalletEthWithdrawnIterator is returned from FilterEthWithdrawn and is used to iterate over the raw logs and unpacked data for EthWithdrawn events raised by the PrizesWallet contract.
type PrizesWalletEthWithdrawnIterator struct {
	Event *PrizesWalletEthWithdrawn // Event containing the contract specifics and raw log

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
func (it *PrizesWalletEthWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizesWalletEthWithdrawn)
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
		it.Event = new(PrizesWalletEthWithdrawn)
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
func (it *PrizesWalletEthWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizesWalletEthWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizesWalletEthWithdrawn represents a EthWithdrawn event raised by the PrizesWallet contract.
type PrizesWalletEthWithdrawn struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawn is a free log retrieval operation binding the contract event 0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b.
//
// Solidity: event EthWithdrawn(address indexed winner, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) FilterEthWithdrawn(opts *bind.FilterOpts, winner []common.Address) (*PrizesWalletEthWithdrawnIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "EthWithdrawn", winnerRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletEthWithdrawnIterator{contract: _PrizesWallet.contract, event: "EthWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawn is a free log subscription operation binding the contract event 0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b.
//
// Solidity: event EthWithdrawn(address indexed winner, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) WatchEthWithdrawn(opts *bind.WatchOpts, sink chan<- *PrizesWalletEthWithdrawn, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "EthWithdrawn", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizesWalletEthWithdrawn)
				if err := _PrizesWallet.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
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

// ParseEthWithdrawn is a log parse operation binding the contract event 0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b.
//
// Solidity: event EthWithdrawn(address indexed winner, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) ParseEthWithdrawn(log types.Log) (*PrizesWalletEthWithdrawn, error) {
	event := new(PrizesWalletEthWithdrawn)
	if err := _PrizesWallet.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
