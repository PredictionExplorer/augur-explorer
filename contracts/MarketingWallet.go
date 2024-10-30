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

// IMarketingWalletMetaData contains all meta data concerning the IMarketingWallet contract.
var IMarketingWalletMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractICosmicToken\",\"name\":\"newCosmicToken\",\"type\":\"address\"}],\"name\":\"CosmicTokenAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"marketer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardSentEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICosmicToken\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// Send is a paid mutator transaction binding the contract method 0x785d04f5.
//
// Solidity: function send(uint256 amount, address to) returns()
func (_IMarketingWallet *IMarketingWalletTransactor) Send(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _IMarketingWallet.contract.Transact(opts, "send", amount, to)
}

// Send is a paid mutator transaction binding the contract method 0x785d04f5.
//
// Solidity: function send(uint256 amount, address to) returns()
func (_IMarketingWallet *IMarketingWalletSession) Send(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.Send(&_IMarketingWallet.TransactOpts, amount, to)
}

// Send is a paid mutator transaction binding the contract method 0x785d04f5.
//
// Solidity: function send(uint256 amount, address to) returns()
func (_IMarketingWallet *IMarketingWalletTransactorSession) Send(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.Send(&_IMarketingWallet.TransactOpts, amount, to)
}

// SetTokenContract is a paid mutator transaction binding the contract method 0xbbcd5bbe.
//
// Solidity: function setTokenContract(address addr) returns()
func (_IMarketingWallet *IMarketingWalletTransactor) SetTokenContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _IMarketingWallet.contract.Transact(opts, "setTokenContract", addr)
}

// SetTokenContract is a paid mutator transaction binding the contract method 0xbbcd5bbe.
//
// Solidity: function setTokenContract(address addr) returns()
func (_IMarketingWallet *IMarketingWalletSession) SetTokenContract(addr common.Address) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.SetTokenContract(&_IMarketingWallet.TransactOpts, addr)
}

// SetTokenContract is a paid mutator transaction binding the contract method 0xbbcd5bbe.
//
// Solidity: function setTokenContract(address addr) returns()
func (_IMarketingWallet *IMarketingWalletTransactorSession) SetTokenContract(addr common.Address) (*types.Transaction, error) {
	return _IMarketingWallet.Contract.SetTokenContract(&_IMarketingWallet.TransactOpts, addr)
}

// IMarketingWalletCosmicTokenAddressChangedIterator is returned from FilterCosmicTokenAddressChanged and is used to iterate over the raw logs and unpacked data for CosmicTokenAddressChanged events raised by the IMarketingWallet contract.
type IMarketingWalletCosmicTokenAddressChangedIterator struct {
	Event *IMarketingWalletCosmicTokenAddressChanged // Event containing the contract specifics and raw log

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
func (it *IMarketingWalletCosmicTokenAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketingWalletCosmicTokenAddressChanged)
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
		it.Event = new(IMarketingWalletCosmicTokenAddressChanged)
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
func (it *IMarketingWalletCosmicTokenAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketingWalletCosmicTokenAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketingWalletCosmicTokenAddressChanged represents a CosmicTokenAddressChanged event raised by the IMarketingWallet contract.
type IMarketingWalletCosmicTokenAddressChanged struct {
	NewCosmicToken common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCosmicTokenAddressChanged is a free log retrieval operation binding the contract event 0x3ab2dcf539aa3c30314265602fc86dee8e343e3c3905730f870835f36c400deb.
//
// Solidity: event CosmicTokenAddressChanged(address newCosmicToken)
func (_IMarketingWallet *IMarketingWalletFilterer) FilterCosmicTokenAddressChanged(opts *bind.FilterOpts) (*IMarketingWalletCosmicTokenAddressChangedIterator, error) {

	logs, sub, err := _IMarketingWallet.contract.FilterLogs(opts, "CosmicTokenAddressChanged")
	if err != nil {
		return nil, err
	}
	return &IMarketingWalletCosmicTokenAddressChangedIterator{contract: _IMarketingWallet.contract, event: "CosmicTokenAddressChanged", logs: logs, sub: sub}, nil
}

// WatchCosmicTokenAddressChanged is a free log subscription operation binding the contract event 0x3ab2dcf539aa3c30314265602fc86dee8e343e3c3905730f870835f36c400deb.
//
// Solidity: event CosmicTokenAddressChanged(address newCosmicToken)
func (_IMarketingWallet *IMarketingWalletFilterer) WatchCosmicTokenAddressChanged(opts *bind.WatchOpts, sink chan<- *IMarketingWalletCosmicTokenAddressChanged) (event.Subscription, error) {

	logs, sub, err := _IMarketingWallet.contract.WatchLogs(opts, "CosmicTokenAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketingWalletCosmicTokenAddressChanged)
				if err := _IMarketingWallet.contract.UnpackLog(event, "CosmicTokenAddressChanged", log); err != nil {
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

// ParseCosmicTokenAddressChanged is a log parse operation binding the contract event 0x3ab2dcf539aa3c30314265602fc86dee8e343e3c3905730f870835f36c400deb.
//
// Solidity: event CosmicTokenAddressChanged(address newCosmicToken)
func (_IMarketingWallet *IMarketingWalletFilterer) ParseCosmicTokenAddressChanged(log types.Log) (*IMarketingWalletCosmicTokenAddressChanged, error) {
	event := new(IMarketingWalletCosmicTokenAddressChanged)
	if err := _IMarketingWallet.contract.UnpackLog(event, "CosmicTokenAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketingWalletRewardSentEventIterator is returned from FilterRewardSentEvent and is used to iterate over the raw logs and unpacked data for RewardSentEvent events raised by the IMarketingWallet contract.
type IMarketingWalletRewardSentEventIterator struct {
	Event *IMarketingWalletRewardSentEvent // Event containing the contract specifics and raw log

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
func (it *IMarketingWalletRewardSentEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketingWalletRewardSentEvent)
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
		it.Event = new(IMarketingWalletRewardSentEvent)
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
func (it *IMarketingWalletRewardSentEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketingWalletRewardSentEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketingWalletRewardSentEvent represents a RewardSentEvent event raised by the IMarketingWallet contract.
type IMarketingWalletRewardSentEvent struct {
	Marketer common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRewardSentEvent is a free log retrieval operation binding the contract event 0xdceb832c3abeeb39542a70bd553842dc873f15f6a5411c645092a93fef4ef9ba.
//
// Solidity: event RewardSentEvent(address indexed marketer, uint256 amount)
func (_IMarketingWallet *IMarketingWalletFilterer) FilterRewardSentEvent(opts *bind.FilterOpts, marketer []common.Address) (*IMarketingWalletRewardSentEventIterator, error) {

	var marketerRule []interface{}
	for _, marketerItem := range marketer {
		marketerRule = append(marketerRule, marketerItem)
	}

	logs, sub, err := _IMarketingWallet.contract.FilterLogs(opts, "RewardSentEvent", marketerRule)
	if err != nil {
		return nil, err
	}
	return &IMarketingWalletRewardSentEventIterator{contract: _IMarketingWallet.contract, event: "RewardSentEvent", logs: logs, sub: sub}, nil
}

// WatchRewardSentEvent is a free log subscription operation binding the contract event 0xdceb832c3abeeb39542a70bd553842dc873f15f6a5411c645092a93fef4ef9ba.
//
// Solidity: event RewardSentEvent(address indexed marketer, uint256 amount)
func (_IMarketingWallet *IMarketingWalletFilterer) WatchRewardSentEvent(opts *bind.WatchOpts, sink chan<- *IMarketingWalletRewardSentEvent, marketer []common.Address) (event.Subscription, error) {

	var marketerRule []interface{}
	for _, marketerItem := range marketer {
		marketerRule = append(marketerRule, marketerItem)
	}

	logs, sub, err := _IMarketingWallet.contract.WatchLogs(opts, "RewardSentEvent", marketerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketingWalletRewardSentEvent)
				if err := _IMarketingWallet.contract.UnpackLog(event, "RewardSentEvent", log); err != nil {
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

// ParseRewardSentEvent is a log parse operation binding the contract event 0xdceb832c3abeeb39542a70bd553842dc873f15f6a5411c645092a93fef4ef9ba.
//
// Solidity: event RewardSentEvent(address indexed marketer, uint256 amount)
func (_IMarketingWallet *IMarketingWalletFilterer) ParseRewardSentEvent(log types.Log) (*IMarketingWalletRewardSentEvent, error) {
	event := new(IMarketingWalletRewardSentEvent)
	if err := _IMarketingWallet.contract.UnpackLog(event, "RewardSentEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketingWalletMetaData contains all meta data concerning the MarketingWallet contract.
var MarketingWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCosmicToken\",\"name\":\"token_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20TransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"NonZeroValueRequired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractICosmicToken\",\"name\":\"newCosmicToken\",\"type\":\"address\"}],\"name\":\"CosmicTokenAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"marketer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardSentEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICosmicToken\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractCosmicToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523461002f57610019610014610100565b61026a565b610021610034565b610bc06103f18239610bc090f35b61003a565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100669061003e565b810190811060018060401b0382111761007e57604052565b610048565b9061009661008f610034565b928361005c565b565b5f80fd5b60018060a01b031690565b6100b09061009c565b90565b6100bc906100a7565b90565b6100c8816100b3565b036100cf57565b5f80fd5b905051906100e0826100bf565b565b906020828203126100fb576100f8915f016100d3565b90565b610098565b61011e610fb18038038061011381610083565b9283398101906100e2565b90565b90565b61013861013361013d9261009c565b610121565b61009c565b90565b61014990610124565b90565b61015590610140565b90565b90565b61016f61016a61017492610158565b610121565b61009c565b90565b6101809061015b565b90565b60209181520190565b5f7f5a65726f2d616464726573732077617320676976656e2e000000000000000000910152565b6101c06017602092610183565b6101c98161018c565b0190565b6101e29060208101905f8183039101526101b3565b90565b156101ec57565b6101f4610034565b63eac0d38960e01b81528061020b600482016101cd565b0390fd5b5f1b90565b9061022560018060a01b039161020f565b9181191691161790565b61023890610124565b90565b6102449061022f565b90565b90565b9061025f61025a6102669261023b565b610247565b8254610214565b9055565b6102ab90610277336102ad565b6102a46102838261014c565b61029d6102976102925f610177565b6100a7565b916100a7565b14156101e5565b600161024a565b565b6102b6906102da565b565b6102c1906100a7565b9052565b91906102d8905f602085019401906102b8565b565b806102f56102ef6102ea5f610177565b6100a7565b916100a7565b146103055761030390610391565b565b6103286103115f610177565b5f918291631e4fbdf760e01b8352600483016102c5565b0390fd5b5f1c90565b60018060a01b031690565b61034861034d9161032c565b610331565b90565b61035a905461033c565b90565b61036690610140565b90565b90565b9061038161037c6103889261035d565b610369565b8254610214565b9055565b5f0190565b61039a5f610350565b6103a4825f61036c565b906103d86103d27f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09361035d565b9161035d565b916103e1610034565b806103eb8161038c565b0390a356fe60806040526004361015610013575b610392565b61001d5f3561007c565b8063715018a614610077578063785d04f5146100725780638da5cb5b1461006d578063bbcd5bbe14610068578063f2fde38b146100635763fc0c546a0361000e5761035d565b61028d565b61023c565b6101ba565b610164565b6100a4565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f91031261009a57565b61008c565b5f0190565b346100d2576100b4366004610090565b6100bc6103e6565b6100c4610082565b806100ce8161009f565b0390f35b610088565b90565b6100e3816100d7565b036100ea57565b5f80fd5b905035906100fb826100da565b565b60018060a01b031690565b610111906100fd565b90565b61011d81610108565b0361012457565b5f80fd5b9050359061013582610114565b565b919060408382031261015f578061015361015c925f86016100ee565b93602001610128565b90565b61008c565b346101935761017d610177366004610137565b90610833565b610185610082565b8061018f8161009f565b0390f35b610088565b6101a190610108565b9052565b91906101b8905f60208501940190610198565b565b346101ea576101ca366004610090565b6101e66101d561086f565b6101dd610082565b918291826101a5565b0390f35b610088565b6101f890610108565b90565b610204816101ef565b0361020b57565b5f80fd5b9050359061021c826101fb565b565b9060208282031261023757610234915f0161020f565b90565b61008c565b3461026a5761025461024f36600461021e565b610a32565b61025c610082565b806102668161009f565b0390f35b610088565b9060208282031261028857610285915f01610128565b90565b61008c565b346102bb576102a56102a036600461026f565b610aa2565b6102ad610082565b806102b78161009f565b0390f35b610088565b1c90565b60018060a01b031690565b6102df9060086102e493026102c0565b6102c4565b90565b906102f291546102cf565b90565b61030160015f906102e7565b90565b90565b61031b610316610320926100fd565b610304565b6100fd565b90565b61032c90610307565b90565b61033890610323565b90565b6103449061032f565b9052565b919061035b905f6020850194019061033b565b565b3461038d5761036d366004610090565b6103896103786102f5565b610380610082565b91829182610348565b0390f35b610088565b5f80fd5b61039e610aad565b6103a66103d3565b565b90565b6103bf6103ba6103c4926103a8565b610304565b6100fd565b90565b6103d0906103ab565b90565b6103e46103df5f6103c7565b610b1e565b565b6103ee610396565b565b90610402916103fd610aad565b6106f9565b565b60209181520190565b60207f2e00000000000000000000000000000000000000000000000000000000000000917f526563697069656e7420616464726573732063616e6e6f74206265207a65726f5f8201520152565b6104676021604092610404565b6104708161040d565b0190565b6104899060208101905f81830391015261045a565b90565b1561049357565b61049b610082565b63eac0d38960e01b8152806104b260048201610474565b0390fd5b6104ca6104c56104cf926103a8565b610304565b6100d7565b90565b60207f2e00000000000000000000000000000000000000000000000000000000000000917f416d6f756e74206d7573742062652067726561746572207468616e207a65726f5f8201520152565b61052c6021604092610404565b610535816104d2565b0190565b61054e9060208101905f81830391015261051f565b90565b1561055857565b610560610082565b63af33979960e01b81528061057760048201610539565b0390fd5b5f1c90565b61058c6105919161057b565b6102c4565b90565b61059e9054610580565b90565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906105c9906105a1565b810190811067ffffffffffffffff8211176105e357604052565b6105ab565b60e01b90565b151590565b6105fc816105ee565b0361060357565b5f80fd5b90505190610614826105f3565b565b9060208282031261062f5761062c915f01610607565b90565b61008c565b61063d906100d7565b9052565b91602061066292949361065b60408201965f830190610198565b0190610634565b565b5f7f5472616e73666572206661696c65642e00000000000000000000000000000000910152565b6106986010602092610404565b6106a181610664565b0190565b9160406106d69294936106cf6106c4606083018381035f85015261068b565b966020830190610198565b0190610634565b565b6106e190610323565b90565b91906106f7905f60208501940190610634565b565b9061071f8161071861071261070d5f6103c7565b610108565b91610108565b141561048c565b61073b8261073561072f5f6104b6565b916100d7565b11610551565b61074d6107486001610594565b61032f565b602063a9059cbb9183906107745f879561077f610768610082565b978896879586946105e8565b845260048401610641565b03925af19081610807575b50155f146108025760016107dd575b6107d86107c67fdceb832c3abeeb39542a70bd553842dc873f15f6a5411c645092a93fef4ef9ba926106d8565b926107cf610082565b918291826106e4565b0390a2565b6107fe6107e8610082565b92839263f7fce64560e01b8452600484016106a5565b0390fd5b610799565b6108279060203d811161082c575b61081f81836105bf565b810190610616565b61078a565b503d610815565b9061083d916103f0565b565b5f90565b60018060a01b031690565b61085a61085f9161057b565b610843565b90565b61086c905461084e565b90565b61087761083f565b506108815f610862565b90565b61089590610890610aad565b6109af565b565b6108a090610323565b90565b5f7f5a65726f2d616464726573732077617320676976656e2e000000000000000000910152565b6108d76017602092610404565b6108e0816108a3565b0190565b6108f99060208101905f8183039101526108ca565b90565b1561090357565b61090b610082565b63eac0d38960e01b815280610922600482016108e4565b0390fd5b61092f90610307565b90565b61093b90610926565b90565b5f1b90565b9061095460018060a01b039161093e565b9181191691161790565b61096790610926565b90565b90565b9061098261097d6109899261095e565b61096a565b8254610943565b9055565b61099690610897565b9052565b91906109ad905f6020850194019061098d565b565b6109dc6109bb82610897565b6109d56109cf6109ca5f6103c7565b610108565b91610108565b14156108fc565b6109f76109f06109eb83610897565b610932565b600161096d565b610a2d7f3ab2dcf539aa3c30314265602fc86dee8e343e3c3905730f870835f36c400deb91610a24610082565b9182918261099a565b0390a1565b610a3b90610884565b565b610a4e90610a49610aad565b610a50565b565b80610a6b610a65610a605f6103c7565b610108565b91610108565b14610a7b57610a7990610b1e565b565b610a9e610a875f6103c7565b5f918291631e4fbdf760e01b8352600483016101a5565b0390fd5b610aab90610a3d565b565b610ab561086f565b610ace610ac8610ac3610b7d565b610108565b91610108565b03610ad557565b610af7610ae0610b7d565b5f91829163118cdaa760e01b8352600483016101a5565b0390fd5b90565b90610b13610b0e610b1a926106d8565b610afb565b8254610943565b9055565b610b275f610862565b610b31825f610afe565b90610b65610b5f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936106d8565b916106d8565b91610b6e610082565b80610b788161009f565b0390a3565b610b8561083f565b50339056fea26469706673582212209b7a67ecc55173b090cf484501ba15e0b40b3b3352f4b45682575213d1189b1364736f6c634300081b0033",
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

// Send is a paid mutator transaction binding the contract method 0x785d04f5.
//
// Solidity: function send(uint256 amount, address to) returns()
func (_MarketingWallet *MarketingWalletTransactor) Send(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _MarketingWallet.contract.Transact(opts, "send", amount, to)
}

// Send is a paid mutator transaction binding the contract method 0x785d04f5.
//
// Solidity: function send(uint256 amount, address to) returns()
func (_MarketingWallet *MarketingWalletSession) Send(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _MarketingWallet.Contract.Send(&_MarketingWallet.TransactOpts, amount, to)
}

// Send is a paid mutator transaction binding the contract method 0x785d04f5.
//
// Solidity: function send(uint256 amount, address to) returns()
func (_MarketingWallet *MarketingWalletTransactorSession) Send(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _MarketingWallet.Contract.Send(&_MarketingWallet.TransactOpts, amount, to)
}

// SetTokenContract is a paid mutator transaction binding the contract method 0xbbcd5bbe.
//
// Solidity: function setTokenContract(address addr) returns()
func (_MarketingWallet *MarketingWalletTransactor) SetTokenContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _MarketingWallet.contract.Transact(opts, "setTokenContract", addr)
}

// SetTokenContract is a paid mutator transaction binding the contract method 0xbbcd5bbe.
//
// Solidity: function setTokenContract(address addr) returns()
func (_MarketingWallet *MarketingWalletSession) SetTokenContract(addr common.Address) (*types.Transaction, error) {
	return _MarketingWallet.Contract.SetTokenContract(&_MarketingWallet.TransactOpts, addr)
}

// SetTokenContract is a paid mutator transaction binding the contract method 0xbbcd5bbe.
//
// Solidity: function setTokenContract(address addr) returns()
func (_MarketingWallet *MarketingWalletTransactorSession) SetTokenContract(addr common.Address) (*types.Transaction, error) {
	return _MarketingWallet.Contract.SetTokenContract(&_MarketingWallet.TransactOpts, addr)
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

// MarketingWalletCosmicTokenAddressChangedIterator is returned from FilterCosmicTokenAddressChanged and is used to iterate over the raw logs and unpacked data for CosmicTokenAddressChanged events raised by the MarketingWallet contract.
type MarketingWalletCosmicTokenAddressChangedIterator struct {
	Event *MarketingWalletCosmicTokenAddressChanged // Event containing the contract specifics and raw log

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
func (it *MarketingWalletCosmicTokenAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketingWalletCosmicTokenAddressChanged)
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
		it.Event = new(MarketingWalletCosmicTokenAddressChanged)
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
func (it *MarketingWalletCosmicTokenAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketingWalletCosmicTokenAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketingWalletCosmicTokenAddressChanged represents a CosmicTokenAddressChanged event raised by the MarketingWallet contract.
type MarketingWalletCosmicTokenAddressChanged struct {
	NewCosmicToken common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCosmicTokenAddressChanged is a free log retrieval operation binding the contract event 0x3ab2dcf539aa3c30314265602fc86dee8e343e3c3905730f870835f36c400deb.
//
// Solidity: event CosmicTokenAddressChanged(address newCosmicToken)
func (_MarketingWallet *MarketingWalletFilterer) FilterCosmicTokenAddressChanged(opts *bind.FilterOpts) (*MarketingWalletCosmicTokenAddressChangedIterator, error) {

	logs, sub, err := _MarketingWallet.contract.FilterLogs(opts, "CosmicTokenAddressChanged")
	if err != nil {
		return nil, err
	}
	return &MarketingWalletCosmicTokenAddressChangedIterator{contract: _MarketingWallet.contract, event: "CosmicTokenAddressChanged", logs: logs, sub: sub}, nil
}

// WatchCosmicTokenAddressChanged is a free log subscription operation binding the contract event 0x3ab2dcf539aa3c30314265602fc86dee8e343e3c3905730f870835f36c400deb.
//
// Solidity: event CosmicTokenAddressChanged(address newCosmicToken)
func (_MarketingWallet *MarketingWalletFilterer) WatchCosmicTokenAddressChanged(opts *bind.WatchOpts, sink chan<- *MarketingWalletCosmicTokenAddressChanged) (event.Subscription, error) {

	logs, sub, err := _MarketingWallet.contract.WatchLogs(opts, "CosmicTokenAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketingWalletCosmicTokenAddressChanged)
				if err := _MarketingWallet.contract.UnpackLog(event, "CosmicTokenAddressChanged", log); err != nil {
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

// ParseCosmicTokenAddressChanged is a log parse operation binding the contract event 0x3ab2dcf539aa3c30314265602fc86dee8e343e3c3905730f870835f36c400deb.
//
// Solidity: event CosmicTokenAddressChanged(address newCosmicToken)
func (_MarketingWallet *MarketingWalletFilterer) ParseCosmicTokenAddressChanged(log types.Log) (*MarketingWalletCosmicTokenAddressChanged, error) {
	event := new(MarketingWalletCosmicTokenAddressChanged)
	if err := _MarketingWallet.contract.UnpackLog(event, "CosmicTokenAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// MarketingWalletRewardSentEventIterator is returned from FilterRewardSentEvent and is used to iterate over the raw logs and unpacked data for RewardSentEvent events raised by the MarketingWallet contract.
type MarketingWalletRewardSentEventIterator struct {
	Event *MarketingWalletRewardSentEvent // Event containing the contract specifics and raw log

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
func (it *MarketingWalletRewardSentEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketingWalletRewardSentEvent)
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
		it.Event = new(MarketingWalletRewardSentEvent)
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
func (it *MarketingWalletRewardSentEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketingWalletRewardSentEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketingWalletRewardSentEvent represents a RewardSentEvent event raised by the MarketingWallet contract.
type MarketingWalletRewardSentEvent struct {
	Marketer common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRewardSentEvent is a free log retrieval operation binding the contract event 0xdceb832c3abeeb39542a70bd553842dc873f15f6a5411c645092a93fef4ef9ba.
//
// Solidity: event RewardSentEvent(address indexed marketer, uint256 amount)
func (_MarketingWallet *MarketingWalletFilterer) FilterRewardSentEvent(opts *bind.FilterOpts, marketer []common.Address) (*MarketingWalletRewardSentEventIterator, error) {

	var marketerRule []interface{}
	for _, marketerItem := range marketer {
		marketerRule = append(marketerRule, marketerItem)
	}

	logs, sub, err := _MarketingWallet.contract.FilterLogs(opts, "RewardSentEvent", marketerRule)
	if err != nil {
		return nil, err
	}
	return &MarketingWalletRewardSentEventIterator{contract: _MarketingWallet.contract, event: "RewardSentEvent", logs: logs, sub: sub}, nil
}

// WatchRewardSentEvent is a free log subscription operation binding the contract event 0xdceb832c3abeeb39542a70bd553842dc873f15f6a5411c645092a93fef4ef9ba.
//
// Solidity: event RewardSentEvent(address indexed marketer, uint256 amount)
func (_MarketingWallet *MarketingWalletFilterer) WatchRewardSentEvent(opts *bind.WatchOpts, sink chan<- *MarketingWalletRewardSentEvent, marketer []common.Address) (event.Subscription, error) {

	var marketerRule []interface{}
	for _, marketerItem := range marketer {
		marketerRule = append(marketerRule, marketerItem)
	}

	logs, sub, err := _MarketingWallet.contract.WatchLogs(opts, "RewardSentEvent", marketerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketingWalletRewardSentEvent)
				if err := _MarketingWallet.contract.UnpackLog(event, "RewardSentEvent", log); err != nil {
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

// ParseRewardSentEvent is a log parse operation binding the contract event 0xdceb832c3abeeb39542a70bd553842dc873f15f6a5411c645092a93fef4ef9ba.
//
// Solidity: event RewardSentEvent(address indexed marketer, uint256 amount)
func (_MarketingWallet *MarketingWalletFilterer) ParseRewardSentEvent(log types.Log) (*MarketingWalletRewardSentEvent, error) {
	event := new(MarketingWalletRewardSentEvent)
	if err := _MarketingWallet.contract.UnpackLog(event, "RewardSentEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageSlotMetaData contains all meta data concerning the StorageSlot contract.
var StorageSlotMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x608060405234601d57600e6021565b603e602c823930815050603e90f35b6027565b60405190565b5f80fdfe60806040525f80fdfea26469706673582212206549347550004ac396a4cdf8ad8cfa3f544ecf22617d52f6bd992f05b08e3dde64736f6c634300081b0033",
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
	Bin: "0x608060405234601d57600e6021565b603e602c823930815050603e90f35b6027565b60405190565b5f80fdfe60806040525f80fdfea2646970667358221220ea46be6f062b7c268a1fd8205f2804eab40b9bbfa0d220e7e8ce3b26ce138ac964736f6c634300081b0033",
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
