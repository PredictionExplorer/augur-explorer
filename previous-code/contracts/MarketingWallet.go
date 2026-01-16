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

// ICosmicSignatureTokenMintSpec is an auto generated low-level Go binding around an user-defined struct.
type ICosmicSignatureTokenMintSpec struct {
	Account common.Address
	Value   *big.Int
}

// IMarketingWalletMetaData contains all meta data concerning the IMarketingWallet contract.
var IMarketingWalletMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"marketerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"TreasurerAddressChanged\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structICosmicSignatureToken.MintSpec[]\",\"name\":\"specs_\",\"type\":\"tuple[]\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"marketerAddresses_\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"marketerAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"payReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setTreasurerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IMarketingWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use IMarketingWalletMetaData.ABI instead.
var IMarketingWalletABI = IMarketingWalletMetaData.ABI

// IMarketingWallet is an auto generated Go binding around an Ethereum contract.
type IMarketingWallet struct {
	IMarketingWalletCaller     // Read-only binding to the contract
	IMarketingWalletTransactor // Write-only binding to the contract
	IMarketingWalletFilterer   // Log filterer for contract events
}

// IMarketingWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMarketingWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMarketingWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMarketingWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMarketingWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMarketingWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMarketingWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMarketingWalletSession struct {
	Contract     *IMarketingWallet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMarketingWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMarketingWalletCallerSession struct {
	Contract *IMarketingWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IMarketingWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMarketingWalletTransactorSession struct {
	Contract     *IMarketingWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IMarketingWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMarketingWalletRaw struct {
	Contract *IMarketingWallet // Generic contract binding to access the raw methods on
}

// IMarketingWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMarketingWalletCallerRaw struct {
	Contract *IMarketingWalletCaller // Generic read-only contract binding to access the raw methods on
}

// IMarketingWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMarketingWalletTransactorRaw struct {
	Contract *IMarketingWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMarketingWallet creates a new instance of IMarketingWallet, bound to a specific deployed contract.
func NewIMarketingWallet(address common.Address, backend bind.ContractBackend) (*IMarketingWallet, error) {
	contract, err := bindIMarketingWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMarketingWallet{IMarketingWalletCaller: IMarketingWalletCaller{contract: contract}, IMarketingWalletTransactor: IMarketingWalletTransactor{contract: contract}, IMarketingWalletFilterer: IMarketingWalletFilterer{contract: contract}}, nil
}

// NewIMarketingWalletCaller creates a new read-only instance of IMarketingWallet, bound to a specific deployed contract.
func NewIMarketingWalletCaller(address common.Address, caller bind.ContractCaller) (*IMarketingWalletCaller, error) {
	contract, err := bindIMarketingWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMarketingWalletCaller{contract: contract}, nil
}

// NewIMarketingWalletTransactor creates a new write-only instance of IMarketingWallet, bound to a specific deployed contract.
func NewIMarketingWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*IMarketingWalletTransactor, error) {
	contract, err := bindIMarketingWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMarketingWalletTransactor{contract: contract}, nil
}

// NewIMarketingWalletFilterer creates a new log filterer instance of IMarketingWallet, bound to a specific deployed contract.
func NewIMarketingWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*IMarketingWalletFilterer, error) {
	contract, err := bindIMarketingWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMarketingWalletFilterer{contract: contract}, nil
}

// bindIMarketingWallet binds a generic wrapper to an already deployed contract.
func bindIMarketingWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMarketingWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMarketingWallet *IMarketingWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMarketingWallet.Contract.IMarketingWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMarketingWallet *IMarketingWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.IMarketingWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMarketingWallet *IMarketingWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.IMarketingWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMarketingWallet *IMarketingWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMarketingWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMarketingWallet *IMarketingWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMarketingWallet *IMarketingWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.contract.Transact(opts, method, params...)
}

// PayManyRewards is a paid mutator transaction binding the contract method 0xc603317f.
//
// Solidity: function payManyRewards((address,uint256)[] specs_) returns()
func (_IMarketingWallet *IMarketingWalletTransactor) PayManyRewards(opts *bind.TransactOpts, specs_ []ICosmicSignatureTokenMintSpec) (*types.Transaction, error) {
	return _IMarketingWallet.contract.Transact(opts, "payManyRewards", specs_)
}

// PayManyRewards is a paid mutator transaction binding the contract method 0xc603317f.
//
// Solidity: function payManyRewards((address,uint256)[] specs_) returns()
func (_IMarketingWallet *IMarketingWalletSession) PayManyRewards(specs_ []ICosmicSignatureTokenMintSpec) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.PayManyRewards(&_IMarketingWallet.TransactOpts, specs_)
}

// PayManyRewards is a paid mutator transaction binding the contract method 0xc603317f.
//
// Solidity: function payManyRewards((address,uint256)[] specs_) returns()
func (_IMarketingWallet *IMarketingWalletTransactorSession) PayManyRewards(specs_ []ICosmicSignatureTokenMintSpec) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.PayManyRewards(&_IMarketingWallet.TransactOpts, specs_)
}

// PayManyRewards0 is a paid mutator transaction binding the contract method 0xffcbf0b4.
//
// Solidity: function payManyRewards(address[] marketerAddresses_, uint256 amount_) returns()
func (_IMarketingWallet *IMarketingWalletTransactor) PayManyRewards0(opts *bind.TransactOpts, marketerAddresses_ []common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IMarketingWallet.contract.Transact(opts, "payManyRewards0", marketerAddresses_, amount_)
}

// PayManyRewards0 is a paid mutator transaction binding the contract method 0xffcbf0b4.
//
// Solidity: function payManyRewards(address[] marketerAddresses_, uint256 amount_) returns()
func (_IMarketingWallet *IMarketingWalletSession) PayManyRewards0(marketerAddresses_ []common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.PayManyRewards0(&_IMarketingWallet.TransactOpts, marketerAddresses_, amount_)
}

// PayManyRewards0 is a paid mutator transaction binding the contract method 0xffcbf0b4.
//
// Solidity: function payManyRewards(address[] marketerAddresses_, uint256 amount_) returns()
func (_IMarketingWallet *IMarketingWalletTransactorSession) PayManyRewards0(marketerAddresses_ []common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.PayManyRewards0(&_IMarketingWallet.TransactOpts, marketerAddresses_, amount_)
}

// PayReward is a paid mutator transaction binding the contract method 0x8b8afcd7.
//
// Solidity: function payReward(address marketerAddress_, uint256 amount_) returns()
func (_IMarketingWallet *IMarketingWalletTransactor) PayReward(opts *bind.TransactOpts, marketerAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IMarketingWallet.contract.Transact(opts, "payReward", marketerAddress_, amount_)
}

// PayReward is a paid mutator transaction binding the contract method 0x8b8afcd7.
//
// Solidity: function payReward(address marketerAddress_, uint256 amount_) returns()
func (_IMarketingWallet *IMarketingWalletSession) PayReward(marketerAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.PayReward(&_IMarketingWallet.TransactOpts, marketerAddress_, amount_)
}

// PayReward is a paid mutator transaction binding the contract method 0x8b8afcd7.
//
// Solidity: function payReward(address marketerAddress_, uint256 amount_) returns()
func (_IMarketingWallet *IMarketingWalletTransactorSession) PayReward(marketerAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.PayReward(&_IMarketingWallet.TransactOpts, marketerAddress_, amount_)
}

// SetTreasurerAddress is a paid mutator transaction binding the contract method 0xa351f75a.
//
// Solidity: function setTreasurerAddress(address newValue_) returns()
func (_IMarketingWallet *IMarketingWalletTransactor) SetTreasurerAddress(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _IMarketingWallet.contract.Transact(opts, "setTreasurerAddress", newValue_)
}

// SetTreasurerAddress is a paid mutator transaction binding the contract method 0xa351f75a.
//
// Solidity: function setTreasurerAddress(address newValue_) returns()
func (_IMarketingWallet *IMarketingWalletSession) SetTreasurerAddress(newValue_ common.Address) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.SetTreasurerAddress(&_IMarketingWallet.TransactOpts, newValue_)
}

// SetTreasurerAddress is a paid mutator transaction binding the contract method 0xa351f75a.
//
// Solidity: function setTreasurerAddress(address newValue_) returns()
func (_IMarketingWallet *IMarketingWalletTransactorSession) SetTreasurerAddress(newValue_ common.Address) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.SetTreasurerAddress(&_IMarketingWallet.TransactOpts, newValue_)
}

// IMarketingWalletRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the IMarketingWallet contract.
type IMarketingWalletRewardPaidIterator struct {
	Event *IMarketingWalletRewardPaid // Event containing the contract specifics and raw log

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
func (it *IMarketingWalletRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketingWalletRewardPaid)
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
		it.Event = new(IMarketingWalletRewardPaid)
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
func (it *IMarketingWalletRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketingWalletRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketingWalletRewardPaid represents a RewardPaid event raised by the IMarketingWallet contract.
type IMarketingWalletRewardPaid struct {
	MarketerAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed marketerAddress, uint256 amount)
func (_IMarketingWallet *IMarketingWalletFilterer) FilterRewardPaid(opts *bind.FilterOpts, marketerAddress []common.Address) (*IMarketingWalletRewardPaidIterator, error) {

	var marketerAddressRule []interface{}
	for _, marketerAddressItem := range marketerAddress {
		marketerAddressRule = append(marketerAddressRule, marketerAddressItem)
	}

	logs, sub, err := _IMarketingWallet.contract.FilterLogs(opts, "RewardPaid", marketerAddressRule)
	if err != nil {
		return nil, err
	}
	return &IMarketingWalletRewardPaidIterator{contract: _IMarketingWallet.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed marketerAddress, uint256 amount)
func (_IMarketingWallet *IMarketingWalletFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *IMarketingWalletRewardPaid, marketerAddress []common.Address) (event.Subscription, error) {

	var marketerAddressRule []interface{}
	for _, marketerAddressItem := range marketerAddress {
		marketerAddressRule = append(marketerAddressRule, marketerAddressItem)
	}

	logs, sub, err := _IMarketingWallet.contract.WatchLogs(opts, "RewardPaid", marketerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketingWalletRewardPaid)
				if err := _IMarketingWallet.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed marketerAddress, uint256 amount)
func (_IMarketingWallet *IMarketingWalletFilterer) ParseRewardPaid(log types.Log) (*IMarketingWalletRewardPaid, error) {
	event := new(IMarketingWalletRewardPaid)
	if err := _IMarketingWallet.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketingWalletTreasurerAddressChangedIterator is returned from FilterTreasurerAddressChanged and is used to iterate over the raw logs and unpacked data for TreasurerAddressChanged events raised by the IMarketingWallet contract.
type IMarketingWalletTreasurerAddressChangedIterator struct {
	Event *IMarketingWalletTreasurerAddressChanged // Event containing the contract specifics and raw log

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
func (it *IMarketingWalletTreasurerAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketingWalletTreasurerAddressChanged)
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
		it.Event = new(IMarketingWalletTreasurerAddressChanged)
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
func (it *IMarketingWalletTreasurerAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketingWalletTreasurerAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketingWalletTreasurerAddressChanged represents a TreasurerAddressChanged event raised by the IMarketingWallet contract.
type IMarketingWalletTreasurerAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTreasurerAddressChanged is a free log retrieval operation binding the contract event 0xdf73fc12cc071a4834f7ba0e7c6cfe7d23e98866e191ec9e86e6e61614d9e50e.
//
// Solidity: event TreasurerAddressChanged(address indexed newValue)
func (_IMarketingWallet *IMarketingWalletFilterer) FilterTreasurerAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*IMarketingWalletTreasurerAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _IMarketingWallet.contract.FilterLogs(opts, "TreasurerAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &IMarketingWalletTreasurerAddressChangedIterator{contract: _IMarketingWallet.contract, event: "TreasurerAddressChanged", logs: logs, sub: sub}, nil
}

// WatchTreasurerAddressChanged is a free log subscription operation binding the contract event 0xdf73fc12cc071a4834f7ba0e7c6cfe7d23e98866e191ec9e86e6e61614d9e50e.
//
// Solidity: event TreasurerAddressChanged(address indexed newValue)
func (_IMarketingWallet *IMarketingWalletFilterer) WatchTreasurerAddressChanged(opts *bind.WatchOpts, sink chan<- *IMarketingWalletTreasurerAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _IMarketingWallet.contract.WatchLogs(opts, "TreasurerAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketingWalletTreasurerAddressChanged)
				if err := _IMarketingWallet.contract.UnpackLog(event, "TreasurerAddressChanged", log); err != nil {
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

// ParseTreasurerAddressChanged is a log parse operation binding the contract event 0xdf73fc12cc071a4834f7ba0e7c6cfe7d23e98866e191ec9e86e6e61614d9e50e.
//
// Solidity: event TreasurerAddressChanged(address indexed newValue)
func (_IMarketingWallet *IMarketingWalletFilterer) ParseTreasurerAddressChanged(log types.Log) (*IMarketingWalletTreasurerAddressChanged, error) {
	event := new(IMarketingWalletTreasurerAddressChanged)
	if err := _IMarketingWallet.contract.UnpackLog(event, "TreasurerAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketingWalletMetaData contains all meta data concerning the MarketingWallet contract.
var MarketingWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCosmicSignatureToken\",\"name\":\"token_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"}],\"name\":\"UnauthorizedCaller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"marketerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"TreasurerAddressChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structICosmicSignatureToken.MintSpec[]\",\"name\":\"specs_\",\"type\":\"tuple[]\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"marketerAddresses_\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"payManyRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"marketerAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"payReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setTreasurerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractCosmicSignatureToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasurerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523461004d5761001961001461011e565b6101f8565b610021610052565b6110a861042082396080518181816103d7015281816106f701528181610af80152610d7d01526110a890f35b610058565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100849061005c565b810190811060018060401b0382111761009c57604052565b610066565b906100b46100ad610052565b928361007a565b565b5f80fd5b60018060a01b031690565b6100ce906100ba565b90565b6100da906100c5565b90565b6100e6816100d1565b036100ed57565b5f80fd5b905051906100fe826100dd565b565b9060208282031261011957610116915f016100f1565b90565b6100b6565b61013c6114c880380380610131816100a1565b928339810190610100565b90565b90565b61015661015161015b926100ba565b61013f565b6100ba565b90565b61016790610142565b90565b6101739061015e565b90565b6101909061018b6101868261016a565b610346565b6101e1565b565b5f1b90565b906101a860018060a01b0391610192565b9181191691161790565b6101bb9061015e565b90565b90565b906101d66101d16101dd926101b2565b6101be565b8254610197565b9055565b6101f36101ec6102d7565b60016101c1565b608052565b6102119061020c6102076102d7565b610213565b610176565b565b61021c9061021e565b565b61022790610229565b565b61023290610281565b565b90565b61024b61024661025092610234565b61013f565b6100ba565b90565b61025c90610237565b90565b610268906100c5565b9052565b919061027f905f6020850194019061025f565b565b8061029c6102966102915f610253565b6100c5565b916100c5565b146102ac576102aa906103c0565b565b6102cf6102b85f610253565b5f918291631e4fbdf760e01b83526004830161026c565b0390fd5b5f90565b6102df6102d3565b503390565b60209181520190565b5f7f5468652070726f76696465642061646472657373206973207a65726f2e000000910152565b610321601d6020926102e4565b61032a816102ed565b0190565b6103439060208101905f818303910152610314565b90565b61036061035a6103555f610253565b6100c5565b916100c5565b1461036757565b61036f610052565b63eac0d38960e01b8152806103866004820161032e565b0390fd5b5f1c90565b60018060a01b031690565b6103a66103ab9161038a565b61038f565b90565b6103b8905461039a565b90565b5f0190565b6103c95f6103ae565b6103d3825f6101c1565b906104076104017f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936101b2565b916101b2565b91610410610052565b8061041a816103bb565b0390a356fe60806040526004361015610013575b61053f565b61001d5f356100ac565b80636b34a45a146100a7578063715018a6146100a25780638b8afcd71461009d5780638da5cb5b14610098578063a351f75a14610093578063c603317f1461008e578063f2fde38b14610089578063fc0c546a146100845763ffcbf0b40361000e5761050b565b610452565b6103a2565b61036e565b6102ba565b610267565b610233565b610186565b61014c565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f9103126100ca57565b6100bc565b1c90565b60018060a01b031690565b6100ee9060086100f393026100cf565b6100d3565b90565b9061010191546100de565b90565b61011060015f906100f6565b90565b60018060a01b031690565b61012790610113565b90565b6101339061011e565b9052565b919061014a905f6020850194019061012a565b565b3461017c5761015c3660046100c0565b610178610167610104565b61016f6100b2565b91829182610137565b0390f35b6100b8565b5f0190565b346101b4576101963660046100c0565b61019e610593565b6101a66100b2565b806101b081610181565b0390f35b6100b8565b5f80fd5b6101c68161011e565b036101cd57565b5f80fd5b905035906101de826101bd565b565b90565b6101ec816101e0565b036101f357565b5f80fd5b90503590610204826101e3565b565b919060408382031261022e578061022261022b925f86016101d1565b936020016101f7565b90565b6100bc565b346102625761024c610246366004610206565b9061078c565b6102546100b2565b8061025e81610181565b0390f35b6100b8565b34610297576102773660046100c0565b6102936102826107c2565b61028a6100b2565b91829182610137565b0390f35b6100b8565b906020828203126102b5576102b2915f016101d1565b90565b6100bc565b346102e8576102d26102cd36600461029c565b61088e565b6102da6100b2565b806102e481610181565b0390f35b6100b8565b5f80fd5b5f80fd5b5f80fd5b909182601f830112156103335781359167ffffffffffffffff831161032e57602001926040830284011161032957565b6102f5565b6102f1565b6102ed565b90602082820312610369575f82013567ffffffffffffffff81116103645761036092016102f9565b9091565b6101b9565b6100bc565b3461039d57610387610381366004610338565b90610b9c565b61038f6100b2565b8061039981610181565b0390f35b6100b8565b346103d0576103ba6103b536600461029c565b610c0d565b6103c26100b2565b806103cc81610181565b0390f35b6100b8565b7f000000000000000000000000000000000000000000000000000000000000000090565b90565b61041061040b61041592610113565b6103f9565b610113565b90565b610421906103fc565b90565b61042d90610418565b90565b61043990610424565b9052565b9190610450905f60208501940190610430565b565b34610482576104623660046100c0565b61047e61046d6103d5565b6104756100b2565b9182918261043d565b0390f35b6100b8565b909182601f830112156104c15781359167ffffffffffffffff83116104bc5760200192602083028401116104b757565b6102f5565b6102f1565b6102ed565b91604083830312610506575f83013567ffffffffffffffff8111610501576104f3836104fe928601610487565b9390946020016101f7565b90565b6101b9565b6100bc565b3461053a5761052461051e3660046104c6565b91610e22565b61052c6100b2565b8061053681610181565b0390f35b6100b8565b5f80fd5b61054b610e2f565b610553610580565b565b90565b61056c61056761057192610555565b6103f9565b610113565b90565b61057d90610558565b90565b61059161058c5f610574565b610e7d565b565b61059b610543565b565b906105af916105aa610f71565b6106a9565b565b6105ba90610418565b90565b6105c6906101e0565b9052565b91906105dd905f602085019401906105bd565b565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b9061060b906105e3565b810190811067ffffffffffffffff82111761062557604052565b6105ed565b60e01b90565b151590565b61063e81610630565b0361064557565b5f80fd5b9050519061065682610635565b565b906020828203126106715761066e915f01610649565b90565b6100bc565b91602061069792949361069060408201965f83019061012a565b01906105bd565b565b6106a16100b2565b3d5f823e3d90fd5b906020908281906106ef6106dd7fe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486926105b1565b926106e66100b2565b918291826105ca565b0390a261071b7f0000000000000000000000000000000000000000000000000000000000000000610424565b61073e5f63a9059cbb9593956107496107326100b2565b9788968795869461062a565b845260048401610676565b03925af180156107875761075b575b50565b61077b9060203d8111610780575b6107738183610601565b810190610658565b610758565b503d610769565b610699565b906107969161059d565b565b5f90565b5f1c90565b6107ad6107b29161079c565b6100d3565b90565b6107bf90546107a1565b90565b6107ca610798565b506107d45f6107b5565b90565b6107e8906107e3610e2f565b6107ea565b565b6107fc906107f781611021565b610841565b565b5f1b90565b9061081460018060a01b03916107fe565b9181191691161790565b90565b9061083661083161083d926105b1565b61081e565b8254610803565b9055565b61084c816001610821565b6108767fdf73fc12cc071a4834f7ba0e7c6cfe7d23e98866e191ec9e86e6e61614d9e50e916105b1565b9061087f6100b2565b8061088981610181565b0390a2565b610897906107d7565b565b906108ab916108a6610f71565b610a57565b565b5090565b6108c56108c06108ca92610555565b6103f9565b6101e0565b90565b634e487b7160e01b5f52601160045260245ffd5b6108ea906101e0565b5f81146108f8576001900390565b6108cd565b634e487b7160e01b5f52603260045260245ffd5b9190811015610921576040020190565b6108fd565b35610930816101bd565b90565b3561093d816101e3565b90565b5f91031261094a57565b6100bc565b60209181520190565b90565b5061096a9060208101906101d1565b90565b6109769061011e565b9052565b506109899060208101906101f7565b90565b610995906101e0565b9052565b9060206109c46109cc936109bb6109b25f83018361095b565b5f86019061096d565b8281019061097a565b91019061098c565b565b906109db81604093610999565b0190565b5090565b60400190565b916109f7826109fd9261094f565b92610958565b90815f905b828210610a10575050505090565b90919293610a32610a2c600192610a2788866109df565b6109ce565b956109e3565b920190929192610a02565b9091610a549260208301925f8185039101526109e9565b90565b610a628183906108ad565b5b80610a76610a705f6108b1565b916101e0565b1115610af257610a85906108e1565b92610a9282848691610911565b610aa96020610aa25f8401610926565b9201610933565b90610ae9610ad77fe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486926105b1565b92610ae06100b2565b918291826105ca565b0390a292610a63565b50610b1c7f0000000000000000000000000000000000000000000000000000000000000000610424565b91638d6bd91c919092803b15610b9757610b495f8094610b54610b3d6100b2565b9788968795869461062a565b845260048401610a3d565b03925af18015610b9257610b66575b50565b610b85905f3d8111610b8b575b610b7d8183610601565b810190610940565b5f610b63565b503d610b73565b610699565b6105df565b90610ba691610899565b565b610bb990610bb4610e2f565b610bbb565b565b80610bd6610bd0610bcb5f610574565b61011e565b9161011e565b14610be657610be490610e7d565b565b610c09610bf25f610574565b5f918291631e4fbdf760e01b835260048301610137565b0390fd5b610c1690610ba8565b565b90610c2b9291610c26610f71565b610ce5565b565b5090565b9190811015610c41576020020190565b6108fd565b60209181520190565b90565b90610c5f8160209361096d565b0190565b60200190565b91610c7782610c7d92610c46565b92610c4f565b90815f905b828210610c90575050505090565b90919293610cb2610cac600192610ca7888661095b565b610c52565b95610c63565b920190929192610c82565b939290610cdb602091610ce39460408801918883035f8a0152610c69565b9401906105bd565b565b909192610cf3828490610c2d565b5b80610d07610d015f6108b1565b916101e0565b1115610d7557610d16906108e1565b90610d2b610d2684868591610c31565b610926565b8590610d6c610d5a7fe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486926105b1565b92610d636100b2565b918291826105ca565b0390a290610cf4565b509291610da17f0000000000000000000000000000000000000000000000000000000000000000610424565b630a8e30d992919392813b15610e1d575f610dcf91610dda8296610dc36100b2565b9889978896879561062a565b855260048501610cbd565b03925af18015610e1857610dec575b50565b610e0b905f3d8111610e11575b610e038183610601565b810190610940565b5f610de9565b503d610df9565b610699565b6105df565b90610e2d9291610c18565b565b610e376107c2565b610e50610e4a610e45611065565b61011e565b9161011e565b03610e5757565b610e79610e62611065565b5f91829163118cdaa760e01b835260048301610137565b0390fd5b610e865f6107b5565b610e90825f610821565b90610ec4610ebe7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936105b1565b916105b1565b91610ecd6100b2565b80610ed781610181565b0390a3565b60209181520190565b60207f6f2063616c6c2074686973206d6574686f642e00000000000000000000000000917f4f6e6c7920746865207472657375726572206973207065726d697474656420745f8201520152565b610f3f6033604092610edc565b610f4881610ee5565b0190565b9190610f6f906020610f67604086018681035f880152610f32565b94019061012a565b565b610f79611065565b610f94610f8e610f8960016107b5565b61011e565b9161011e565b03610f9b57565b610fc4610fa6611065565b610fae6100b2565b91829163ced50f6760e01b835260048301610f4c565b0390fd5b5f7f5468652070726f76696465642061646472657373206973207a65726f2e000000910152565b610ffc601d602092610edc565b61100581610fc8565b0190565b61101e9060208101905f818303910152610fef565b90565b61103b6110356110305f610574565b61011e565b9161011e565b1461104257565b61104a6100b2565b63eac0d38960e01b81528061106160048201611009565b0390fd5b61106d610798565b50339056fea26469706673582212203eaf6efa06b3df4b2c1eca1186cf2df82e0bb564cebbe403fb3c85d42c05d27e64736f6c634300081d0033",
}

// MarketingWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketingWalletMetaData.ABI instead.
var MarketingWalletABI = MarketingWalletMetaData.ABI

// MarketingWalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MarketingWalletMetaData.Bin instead.
var MarketingWalletBin = MarketingWalletMetaData.Bin

// DeployMarketingWallet deploys a new Ethereum contract, binding an instance of MarketingWallet to it.
func DeployMarketingWallet(auth *bind.TransactOpts, backend bind.ContractBackend, token_ common.Address) (common.Address, *types.Transaction, *MarketingWallet, error) {
	parsed, err := MarketingWalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MarketingWalletBin), backend, token_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MarketingWallet{MarketingWalletCaller: MarketingWalletCaller{contract: contract}, MarketingWalletTransactor: MarketingWalletTransactor{contract: contract}, MarketingWalletFilterer: MarketingWalletFilterer{contract: contract}}, nil
}

// MarketingWallet is an auto generated Go binding around an Ethereum contract.
type MarketingWallet struct {
	MarketingWalletCaller     // Read-only binding to the contract
	MarketingWalletTransactor // Write-only binding to the contract
	MarketingWalletFilterer   // Log filterer for contract events
}

// MarketingWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketingWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketingWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketingWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketingWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketingWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketingWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketingWalletSession struct {
	Contract     *MarketingWallet  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketingWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketingWalletCallerSession struct {
	Contract *MarketingWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MarketingWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketingWalletTransactorSession struct {
	Contract     *MarketingWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MarketingWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketingWalletRaw struct {
	Contract *MarketingWallet // Generic contract binding to access the raw methods on
}

// MarketingWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketingWalletCallerRaw struct {
	Contract *MarketingWalletCaller // Generic read-only contract binding to access the raw methods on
}

// MarketingWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketingWalletTransactorRaw struct {
	Contract *MarketingWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketingWallet creates a new instance of MarketingWallet, bound to a specific deployed contract.
func NewMarketingWallet(address common.Address, backend bind.ContractBackend) (*MarketingWallet, error) {
	contract, err := bindMarketingWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketingWallet{MarketingWalletCaller: MarketingWalletCaller{contract: contract}, MarketingWalletTransactor: MarketingWalletTransactor{contract: contract}, MarketingWalletFilterer: MarketingWalletFilterer{contract: contract}}, nil
}

// NewMarketingWalletCaller creates a new read-only instance of MarketingWallet, bound to a specific deployed contract.
func NewMarketingWalletCaller(address common.Address, caller bind.ContractCaller) (*MarketingWalletCaller, error) {
	contract, err := bindMarketingWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketingWalletCaller{contract: contract}, nil
}

// NewMarketingWalletTransactor creates a new write-only instance of MarketingWallet, bound to a specific deployed contract.
func NewMarketingWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketingWalletTransactor, error) {
	contract, err := bindMarketingWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketingWalletTransactor{contract: contract}, nil
}

// NewMarketingWalletFilterer creates a new log filterer instance of MarketingWallet, bound to a specific deployed contract.
func NewMarketingWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketingWalletFilterer, error) {
	contract, err := bindMarketingWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketingWalletFilterer{contract: contract}, nil
}

// bindMarketingWallet binds a generic wrapper to an already deployed contract.
func bindMarketingWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketingWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketingWallet *MarketingWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketingWallet.Contract.MarketingWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketingWallet *MarketingWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketingWallet.Contract.MarketingWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketingWallet *MarketingWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketingWallet.Contract.MarketingWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketingWallet *MarketingWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketingWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketingWallet *MarketingWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketingWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketingWallet *MarketingWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketingWallet.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketingWallet *MarketingWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketingWallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketingWallet *MarketingWalletSession) Owner() (common.Address, error) {
	return _MarketingWallet.Contract.Owner(&_MarketingWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketingWallet *MarketingWalletCallerSession) Owner() (common.Address, error) {
	return _MarketingWallet.Contract.Owner(&_MarketingWallet.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MarketingWallet *MarketingWalletCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketingWallet.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MarketingWallet *MarketingWalletSession) Token() (common.Address, error) {
	return _MarketingWallet.Contract.Token(&_MarketingWallet.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MarketingWallet *MarketingWalletCallerSession) Token() (common.Address, error) {
	return _MarketingWallet.Contract.Token(&_MarketingWallet.CallOpts)
}

// TreasurerAddress is a free data retrieval call binding the contract method 0x6b34a45a.
//
// Solidity: function treasurerAddress() view returns(address)
func (_MarketingWallet *MarketingWalletCaller) TreasurerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketingWallet.contract.Call(opts, &out, "treasurerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasurerAddress is a free data retrieval call binding the contract method 0x6b34a45a.
//
// Solidity: function treasurerAddress() view returns(address)
func (_MarketingWallet *MarketingWalletSession) TreasurerAddress() (common.Address, error) {
	return _MarketingWallet.Contract.TreasurerAddress(&_MarketingWallet.CallOpts)
}

// TreasurerAddress is a free data retrieval call binding the contract method 0x6b34a45a.
//
// Solidity: function treasurerAddress() view returns(address)
func (_MarketingWallet *MarketingWalletCallerSession) TreasurerAddress() (common.Address, error) {
	return _MarketingWallet.Contract.TreasurerAddress(&_MarketingWallet.CallOpts)
}

// PayManyRewards is a paid mutator transaction binding the contract method 0xc603317f.
//
// Solidity: function payManyRewards((address,uint256)[] specs_) returns()
func (_MarketingWallet *MarketingWalletTransactor) PayManyRewards(opts *bind.TransactOpts, specs_ []ICosmicSignatureTokenMintSpec) (*types.Transaction, error) {
	return _MarketingWallet.contract.Transact(opts, "payManyRewards", specs_)
}

// PayManyRewards is a paid mutator transaction binding the contract method 0xc603317f.
//
// Solidity: function payManyRewards((address,uint256)[] specs_) returns()
func (_MarketingWallet *MarketingWalletSession) PayManyRewards(specs_ []ICosmicSignatureTokenMintSpec) (*types.Transaction, error) {
	return _MarketingWallet.Contract.PayManyRewards(&_MarketingWallet.TransactOpts, specs_)
}

// PayManyRewards is a paid mutator transaction binding the contract method 0xc603317f.
//
// Solidity: function payManyRewards((address,uint256)[] specs_) returns()
func (_MarketingWallet *MarketingWalletTransactorSession) PayManyRewards(specs_ []ICosmicSignatureTokenMintSpec) (*types.Transaction, error) {
	return _MarketingWallet.Contract.PayManyRewards(&_MarketingWallet.TransactOpts, specs_)
}

// PayManyRewards0 is a paid mutator transaction binding the contract method 0xffcbf0b4.
//
// Solidity: function payManyRewards(address[] marketerAddresses_, uint256 amount_) returns()
func (_MarketingWallet *MarketingWalletTransactor) PayManyRewards0(opts *bind.TransactOpts, marketerAddresses_ []common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MarketingWallet.contract.Transact(opts, "payManyRewards0", marketerAddresses_, amount_)
}

// PayManyRewards0 is a paid mutator transaction binding the contract method 0xffcbf0b4.
//
// Solidity: function payManyRewards(address[] marketerAddresses_, uint256 amount_) returns()
func (_MarketingWallet *MarketingWalletSession) PayManyRewards0(marketerAddresses_ []common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MarketingWallet.Contract.PayManyRewards0(&_MarketingWallet.TransactOpts, marketerAddresses_, amount_)
}

// PayManyRewards0 is a paid mutator transaction binding the contract method 0xffcbf0b4.
//
// Solidity: function payManyRewards(address[] marketerAddresses_, uint256 amount_) returns()
func (_MarketingWallet *MarketingWalletTransactorSession) PayManyRewards0(marketerAddresses_ []common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MarketingWallet.Contract.PayManyRewards0(&_MarketingWallet.TransactOpts, marketerAddresses_, amount_)
}

// PayReward is a paid mutator transaction binding the contract method 0x8b8afcd7.
//
// Solidity: function payReward(address marketerAddress_, uint256 amount_) returns()
func (_MarketingWallet *MarketingWalletTransactor) PayReward(opts *bind.TransactOpts, marketerAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MarketingWallet.contract.Transact(opts, "payReward", marketerAddress_, amount_)
}

// PayReward is a paid mutator transaction binding the contract method 0x8b8afcd7.
//
// Solidity: function payReward(address marketerAddress_, uint256 amount_) returns()
func (_MarketingWallet *MarketingWalletSession) PayReward(marketerAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MarketingWallet.Contract.PayReward(&_MarketingWallet.TransactOpts, marketerAddress_, amount_)
}

// PayReward is a paid mutator transaction binding the contract method 0x8b8afcd7.
//
// Solidity: function payReward(address marketerAddress_, uint256 amount_) returns()
func (_MarketingWallet *MarketingWalletTransactorSession) PayReward(marketerAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _MarketingWallet.Contract.PayReward(&_MarketingWallet.TransactOpts, marketerAddress_, amount_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketingWallet *MarketingWalletTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketingWallet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketingWallet *MarketingWalletSession) RenounceOwnership() (*types.Transaction, error) {
	return _MarketingWallet.Contract.RenounceOwnership(&_MarketingWallet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketingWallet *MarketingWalletTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MarketingWallet.Contract.RenounceOwnership(&_MarketingWallet.TransactOpts)
}

// SetTreasurerAddress is a paid mutator transaction binding the contract method 0xa351f75a.
//
// Solidity: function setTreasurerAddress(address newValue_) returns()
func (_MarketingWallet *MarketingWalletTransactor) SetTreasurerAddress(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _MarketingWallet.contract.Transact(opts, "setTreasurerAddress", newValue_)
}

// SetTreasurerAddress is a paid mutator transaction binding the contract method 0xa351f75a.
//
// Solidity: function setTreasurerAddress(address newValue_) returns()
func (_MarketingWallet *MarketingWalletSession) SetTreasurerAddress(newValue_ common.Address) (*types.Transaction, error) {
	return _MarketingWallet.Contract.SetTreasurerAddress(&_MarketingWallet.TransactOpts, newValue_)
}

// SetTreasurerAddress is a paid mutator transaction binding the contract method 0xa351f75a.
//
// Solidity: function setTreasurerAddress(address newValue_) returns()
func (_MarketingWallet *MarketingWalletTransactorSession) SetTreasurerAddress(newValue_ common.Address) (*types.Transaction, error) {
	return _MarketingWallet.Contract.SetTreasurerAddress(&_MarketingWallet.TransactOpts, newValue_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketingWallet *MarketingWalletTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MarketingWallet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketingWallet *MarketingWalletSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketingWallet.Contract.TransferOwnership(&_MarketingWallet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketingWallet *MarketingWalletTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketingWallet.Contract.TransferOwnership(&_MarketingWallet.TransactOpts, newOwner)
}

// MarketingWalletOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MarketingWallet contract.
type MarketingWalletOwnershipTransferredIterator struct {
	Event *MarketingWalletOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarketingWalletOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketingWalletOwnershipTransferred)
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
		it.Event = new(MarketingWalletOwnershipTransferred)
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
func (it *MarketingWalletOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketingWalletOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketingWalletOwnershipTransferred represents a OwnershipTransferred event raised by the MarketingWallet contract.
type MarketingWalletOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MarketingWallet *MarketingWalletFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarketingWalletOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketingWallet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketingWalletOwnershipTransferredIterator{contract: _MarketingWallet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MarketingWallet *MarketingWalletFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketingWalletOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketingWallet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketingWalletOwnershipTransferred)
				if err := _MarketingWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MarketingWallet *MarketingWalletFilterer) ParseOwnershipTransferred(log types.Log) (*MarketingWalletOwnershipTransferred, error) {
	event := new(MarketingWalletOwnershipTransferred)
	if err := _MarketingWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketingWalletRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the MarketingWallet contract.
type MarketingWalletRewardPaidIterator struct {
	Event *MarketingWalletRewardPaid // Event containing the contract specifics and raw log

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
func (it *MarketingWalletRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketingWalletRewardPaid)
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
		it.Event = new(MarketingWalletRewardPaid)
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
func (it *MarketingWalletRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketingWalletRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketingWalletRewardPaid represents a RewardPaid event raised by the MarketingWallet contract.
type MarketingWalletRewardPaid struct {
	MarketerAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed marketerAddress, uint256 amount)
func (_MarketingWallet *MarketingWalletFilterer) FilterRewardPaid(opts *bind.FilterOpts, marketerAddress []common.Address) (*MarketingWalletRewardPaidIterator, error) {

	var marketerAddressRule []interface{}
	for _, marketerAddressItem := range marketerAddress {
		marketerAddressRule = append(marketerAddressRule, marketerAddressItem)
	}

	logs, sub, err := _MarketingWallet.contract.FilterLogs(opts, "RewardPaid", marketerAddressRule)
	if err != nil {
		return nil, err
	}
	return &MarketingWalletRewardPaidIterator{contract: _MarketingWallet.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed marketerAddress, uint256 amount)
func (_MarketingWallet *MarketingWalletFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *MarketingWalletRewardPaid, marketerAddress []common.Address) (event.Subscription, error) {

	var marketerAddressRule []interface{}
	for _, marketerAddressItem := range marketerAddress {
		marketerAddressRule = append(marketerAddressRule, marketerAddressItem)
	}

	logs, sub, err := _MarketingWallet.contract.WatchLogs(opts, "RewardPaid", marketerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketingWalletRewardPaid)
				if err := _MarketingWallet.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed marketerAddress, uint256 amount)
func (_MarketingWallet *MarketingWalletFilterer) ParseRewardPaid(log types.Log) (*MarketingWalletRewardPaid, error) {
	event := new(MarketingWalletRewardPaid)
	if err := _MarketingWallet.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketingWalletTreasurerAddressChangedIterator is returned from FilterTreasurerAddressChanged and is used to iterate over the raw logs and unpacked data for TreasurerAddressChanged events raised by the MarketingWallet contract.
type MarketingWalletTreasurerAddressChangedIterator struct {
	Event *MarketingWalletTreasurerAddressChanged // Event containing the contract specifics and raw log

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
func (it *MarketingWalletTreasurerAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketingWalletTreasurerAddressChanged)
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
		it.Event = new(MarketingWalletTreasurerAddressChanged)
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
func (it *MarketingWalletTreasurerAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketingWalletTreasurerAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketingWalletTreasurerAddressChanged represents a TreasurerAddressChanged event raised by the MarketingWallet contract.
type MarketingWalletTreasurerAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTreasurerAddressChanged is a free log retrieval operation binding the contract event 0xdf73fc12cc071a4834f7ba0e7c6cfe7d23e98866e191ec9e86e6e61614d9e50e.
//
// Solidity: event TreasurerAddressChanged(address indexed newValue)
func (_MarketingWallet *MarketingWalletFilterer) FilterTreasurerAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*MarketingWalletTreasurerAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _MarketingWallet.contract.FilterLogs(opts, "TreasurerAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &MarketingWalletTreasurerAddressChangedIterator{contract: _MarketingWallet.contract, event: "TreasurerAddressChanged", logs: logs, sub: sub}, nil
}

// WatchTreasurerAddressChanged is a free log subscription operation binding the contract event 0xdf73fc12cc071a4834f7ba0e7c6cfe7d23e98866e191ec9e86e6e61614d9e50e.
//
// Solidity: event TreasurerAddressChanged(address indexed newValue)
func (_MarketingWallet *MarketingWalletFilterer) WatchTreasurerAddressChanged(opts *bind.WatchOpts, sink chan<- *MarketingWalletTreasurerAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _MarketingWallet.contract.WatchLogs(opts, "TreasurerAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketingWalletTreasurerAddressChanged)
				if err := _MarketingWallet.contract.UnpackLog(event, "TreasurerAddressChanged", log); err != nil {
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

// ParseTreasurerAddressChanged is a log parse operation binding the contract event 0xdf73fc12cc071a4834f7ba0e7c6cfe7d23e98866e191ec9e86e6e61614d9e50e.
//
// Solidity: event TreasurerAddressChanged(address indexed newValue)
func (_MarketingWallet *MarketingWalletFilterer) ParseTreasurerAddressChanged(log types.Log) (*MarketingWalletTreasurerAddressChanged, error) {
	event := new(MarketingWalletTreasurerAddressChanged)
	if err := _MarketingWallet.contract.UnpackLog(event, "TreasurerAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageSlotMetaData contains all meta data concerning the StorageSlot contract.
var StorageSlotMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x608060405234601d57600e6021565b603e602c823930815050603e90f35b6027565b60405190565b5f80fdfe60806040525f80fdfea2646970667358221220028f8807dca3f4b7ce909f06ba1286f50697d85d2839d3d84bfc33396bd78ef964736f6c634300081d0033",
}

// StorageSlotABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageSlotMetaData.ABI instead.
var StorageSlotABI = StorageSlotMetaData.ABI

// StorageSlotBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageSlotMetaData.Bin instead.
var StorageSlotBin = StorageSlotMetaData.Bin

// DeployStorageSlot deploys a new Ethereum contract, binding an instance of StorageSlot to it.
func DeployStorageSlot(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StorageSlot, error) {
	parsed, err := StorageSlotMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageSlotBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StorageSlot{StorageSlotCaller: StorageSlotCaller{contract: contract}, StorageSlotTransactor: StorageSlotTransactor{contract: contract}, StorageSlotFilterer: StorageSlotFilterer{contract: contract}}, nil
}

// StorageSlot is an auto generated Go binding around an Ethereum contract.
type StorageSlot struct {
	StorageSlotCaller     // Read-only binding to the contract
	StorageSlotTransactor // Write-only binding to the contract
	StorageSlotFilterer   // Log filterer for contract events
}

// StorageSlotCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageSlotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageSlotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageSlotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSlotSession struct {
	Contract     *StorageSlot      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageSlotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageSlotCallerSession struct {
	Contract *StorageSlotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// StorageSlotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageSlotTransactorSession struct {
	Contract     *StorageSlotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StorageSlotRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageSlotRaw struct {
	Contract *StorageSlot // Generic contract binding to access the raw methods on
}

// StorageSlotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageSlotCallerRaw struct {
	Contract *StorageSlotCaller // Generic read-only contract binding to access the raw methods on
}

// StorageSlotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageSlotTransactorRaw struct {
	Contract *StorageSlotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageSlot creates a new instance of StorageSlot, bound to a specific deployed contract.
func NewStorageSlot(address common.Address, backend bind.ContractBackend) (*StorageSlot, error) {
	contract, err := bindStorageSlot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageSlot{StorageSlotCaller: StorageSlotCaller{contract: contract}, StorageSlotTransactor: StorageSlotTransactor{contract: contract}, StorageSlotFilterer: StorageSlotFilterer{contract: contract}}, nil
}

// NewStorageSlotCaller creates a new read-only instance of StorageSlot, bound to a specific deployed contract.
func NewStorageSlotCaller(address common.Address, caller bind.ContractCaller) (*StorageSlotCaller, error) {
	contract, err := bindStorageSlot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSlotCaller{contract: contract}, nil
}

// NewStorageSlotTransactor creates a new write-only instance of StorageSlot, bound to a specific deployed contract.
func NewStorageSlotTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageSlotTransactor, error) {
	contract, err := bindStorageSlot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSlotTransactor{contract: contract}, nil
}

// NewStorageSlotFilterer creates a new log filterer instance of StorageSlot, bound to a specific deployed contract.
func NewStorageSlotFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageSlotFilterer, error) {
	contract, err := bindStorageSlot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageSlotFilterer{contract: contract}, nil
}

// bindStorageSlot binds a generic wrapper to an already deployed contract.
func bindStorageSlot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageSlotMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageSlot *StorageSlotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageSlot.Contract.StorageSlotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageSlot *StorageSlotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageSlot.Contract.StorageSlotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageSlot *StorageSlotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageSlot.Contract.StorageSlotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageSlot *StorageSlotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageSlot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageSlot *StorageSlotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageSlot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageSlot *StorageSlotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageSlot.Contract.contract.Transact(opts, method, params...)
}

// TimeMetaData contains all meta data concerning the Time contract.
var TimeMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x608060405234601d57600e6021565b603e602c823930815050603e90f35b6027565b60405190565b5f80fdfe60806040525f80fdfea2646970667358221220a85ca47f50ba7a67b6f93406ca59044e215242f2a5a20f7717e54e35e5f3b05864736f6c634300081d0033",
}

// TimeABI is the input ABI used to generate the binding from.
// Deprecated: Use TimeMetaData.ABI instead.
var TimeABI = TimeMetaData.ABI

// TimeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TimeMetaData.Bin instead.
var TimeBin = TimeMetaData.Bin

// DeployTime deploys a new Ethereum contract, binding an instance of Time to it.
func DeployTime(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Time, error) {
	parsed, err := TimeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TimeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Time{TimeCaller: TimeCaller{contract: contract}, TimeTransactor: TimeTransactor{contract: contract}, TimeFilterer: TimeFilterer{contract: contract}}, nil
}

// Time is an auto generated Go binding around an Ethereum contract.
type Time struct {
	TimeCaller     // Read-only binding to the contract
	TimeTransactor // Write-only binding to the contract
	TimeFilterer   // Log filterer for contract events
}

// TimeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimeSession struct {
	Contract     *Time             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimeCallerSession struct {
	Contract *TimeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TimeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimeTransactorSession struct {
	Contract     *TimeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimeRaw struct {
	Contract *Time // Generic contract binding to access the raw methods on
}

// TimeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimeCallerRaw struct {
	Contract *TimeCaller // Generic read-only contract binding to access the raw methods on
}

// TimeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimeTransactorRaw struct {
	Contract *TimeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTime creates a new instance of Time, bound to a specific deployed contract.
func NewTime(address common.Address, backend bind.ContractBackend) (*Time, error) {
	contract, err := bindTime(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Time{TimeCaller: TimeCaller{contract: contract}, TimeTransactor: TimeTransactor{contract: contract}, TimeFilterer: TimeFilterer{contract: contract}}, nil
}

// NewTimeCaller creates a new read-only instance of Time, bound to a specific deployed contract.
func NewTimeCaller(address common.Address, caller bind.ContractCaller) (*TimeCaller, error) {
	contract, err := bindTime(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimeCaller{contract: contract}, nil
}

// NewTimeTransactor creates a new write-only instance of Time, bound to a specific deployed contract.
func NewTimeTransactor(address common.Address, transactor bind.ContractTransactor) (*TimeTransactor, error) {
	contract, err := bindTime(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimeTransactor{contract: contract}, nil
}

// NewTimeFilterer creates a new log filterer instance of Time, bound to a specific deployed contract.
func NewTimeFilterer(address common.Address, filterer bind.ContractFilterer) (*TimeFilterer, error) {
	contract, err := bindTime(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimeFilterer{contract: contract}, nil
}

// bindTime binds a generic wrapper to an already deployed contract.
func bindTime(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TimeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Time *TimeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Time.Contract.TimeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Time *TimeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Time.Contract.TimeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Time *TimeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Time.Contract.TimeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Time *TimeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Time.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Time *TimeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Time.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Time *TimeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Time.Contract.contract.Transact(opts, method, params...)
}
