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
)

// CharityWalletMetaData contains all meta data concerning the CharityWallet contract.
var CharityWalletMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCharityAddress\",\"type\":\"address\"}],\"name\":\"CharityUpdatedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonationReceivedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"charity\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonationSentEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCharityAddress\",\"type\":\"address\"}],\"name\":\"setCharity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061002d61002261003260201b60201c565b61003a60201b60201c565b6100fe565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61091f8061010d6000396000f3fe60806040526004361061004e5760003560e01c8063715018a6146100af5780638da5cb5b146100c6578063b46300ec146100f1578063f2fde38b14610108578063fb6f71a314610131576100aa565b366100aa5761005b61015a565b73ffffffffffffffffffffffffffffffffffffffff167f46ff3d75d4645bdbbae4cd6109ba42e6e1b80ea25e69d10472b357b733300368346040516100a09190610653565b60405180910390a2005b600080fd5b3480156100bb57600080fd5b506100c4610162565b005b3480156100d257600080fd5b506100db6101ea565b6040516100e891906106af565b60405180910390f35b3480156100fd57600080fd5b50610106610213565b005b34801561011457600080fd5b5061012f600480360381019061012a91906106fb565b61035a565b005b34801561013d57600080fd5b50610158600480360381019061015391906106fb565b610451565b005b600033905090565b61016a61015a565b73ffffffffffffffffffffffffffffffffffffffff166101886101ea565b73ffffffffffffffffffffffffffffffffffffffff16146101de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101d590610785565b60405180910390fd5b6101e86000610576565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60004790506000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1682604051610260906107d6565b60006040518083038185875af1925050503d806000811461029d576040519150601f19603f3d011682016040523d82523d6000602084013e6102a2565b606091505b50509050806102e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102dd90610837565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f44d398d152fa0735a428b13ebc78f79fe4cb1b4722292cd233e278552fa36d328360405161034e9190610653565b60405180910390a25050565b61036261015a565b73ffffffffffffffffffffffffffffffffffffffff166103806101ea565b73ffffffffffffffffffffffffffffffffffffffff16146103d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103cd90610785565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610445576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161043c906108c9565b60405180910390fd5b61044e81610576565b50565b61045961015a565b73ffffffffffffffffffffffffffffffffffffffff166104776101ea565b73ffffffffffffffffffffffffffffffffffffffff16146104cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104c490610785565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fa0bd6b2fdbf082ae2356710c23fc8d76d56d418cecb4514d119c77a8617b4ffe60405160405180910390a250565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000819050919050565b61064d8161063a565b82525050565b60006020820190506106686000830184610644565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006106998261066e565b9050919050565b6106a98161068e565b82525050565b60006020820190506106c460008301846106a0565b92915050565b600080fd5b6106d88161068e565b81146106e357600080fd5b50565b6000813590506106f5816106cf565b92915050565b600060208284031215610711576107106106ca565b5b600061071f848285016106e6565b91505092915050565b600082825260208201905092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b600061076f602083610728565b915061077a82610739565b602082019050919050565b6000602082019050818103600083015261079e81610762565b9050919050565b600081905092915050565b50565b60006107c06000836107a5565b91506107cb826107b0565b600082019050919050565b60006107e1826107b3565b9150819050919050565b7f5472616e73666572206661696c65642e00000000000000000000000000000000600082015250565b6000610821601083610728565b915061082c826107eb565b602082019050919050565b6000602082019050818103600083015261085081610814565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006108b3602683610728565b91506108be82610857565b604082019050919050565b600060208201905081810360008301526108e2816108a6565b905091905056fea26469706673582212204f8c043f41023d636126211e53f6457d0232440f964d8abd31d9fc5653e1895c64736f6c63430008120033",
}

// CharityWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use CharityWalletMetaData.ABI instead.
var CharityWalletABI = CharityWalletMetaData.ABI

// CharityWalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CharityWalletMetaData.Bin instead.
var CharityWalletBin = CharityWalletMetaData.Bin

// DeployCharityWallet deploys a new Ethereum contract, binding an instance of CharityWallet to it.
func DeployCharityWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CharityWallet, error) {
	parsed, err := CharityWalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CharityWalletBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CharityWallet{CharityWalletCaller: CharityWalletCaller{contract: contract}, CharityWalletTransactor: CharityWalletTransactor{contract: contract}, CharityWalletFilterer: CharityWalletFilterer{contract: contract}}, nil
}

// CharityWallet is an auto generated Go binding around an Ethereum contract.
type CharityWallet struct {
	CharityWalletCaller     // Read-only binding to the contract
	CharityWalletTransactor // Write-only binding to the contract
	CharityWalletFilterer   // Log filterer for contract events
}

// CharityWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type CharityWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CharityWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CharityWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CharityWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CharityWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CharityWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CharityWalletSession struct {
	Contract     *CharityWallet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CharityWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CharityWalletCallerSession struct {
	Contract *CharityWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CharityWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CharityWalletTransactorSession struct {
	Contract     *CharityWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CharityWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type CharityWalletRaw struct {
	Contract *CharityWallet // Generic contract binding to access the raw methods on
}

// CharityWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CharityWalletCallerRaw struct {
	Contract *CharityWalletCaller // Generic read-only contract binding to access the raw methods on
}

// CharityWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CharityWalletTransactorRaw struct {
	Contract *CharityWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCharityWallet creates a new instance of CharityWallet, bound to a specific deployed contract.
func NewCharityWallet(address common.Address, backend bind.ContractBackend) (*CharityWallet, error) {
	contract, err := bindCharityWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CharityWallet{CharityWalletCaller: CharityWalletCaller{contract: contract}, CharityWalletTransactor: CharityWalletTransactor{contract: contract}, CharityWalletFilterer: CharityWalletFilterer{contract: contract}}, nil
}

// NewCharityWalletCaller creates a new read-only instance of CharityWallet, bound to a specific deployed contract.
func NewCharityWalletCaller(address common.Address, caller bind.ContractCaller) (*CharityWalletCaller, error) {
	contract, err := bindCharityWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CharityWalletCaller{contract: contract}, nil
}

// NewCharityWalletTransactor creates a new write-only instance of CharityWallet, bound to a specific deployed contract.
func NewCharityWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*CharityWalletTransactor, error) {
	contract, err := bindCharityWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CharityWalletTransactor{contract: contract}, nil
}

// NewCharityWalletFilterer creates a new log filterer instance of CharityWallet, bound to a specific deployed contract.
func NewCharityWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*CharityWalletFilterer, error) {
	contract, err := bindCharityWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CharityWalletFilterer{contract: contract}, nil
}

// bindCharityWallet binds a generic wrapper to an already deployed contract.
func bindCharityWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CharityWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CharityWallet *CharityWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CharityWallet.Contract.CharityWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CharityWallet *CharityWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CharityWallet.Contract.CharityWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CharityWallet *CharityWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CharityWallet.Contract.CharityWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CharityWallet *CharityWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CharityWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CharityWallet *CharityWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CharityWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CharityWallet *CharityWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CharityWallet.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CharityWallet *CharityWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CharityWallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CharityWallet *CharityWalletSession) Owner() (common.Address, error) {
	return _CharityWallet.Contract.Owner(&_CharityWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CharityWallet *CharityWalletCallerSession) Owner() (common.Address, error) {
	return _CharityWallet.Contract.Owner(&_CharityWallet.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CharityWallet *CharityWalletTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CharityWallet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CharityWallet *CharityWalletSession) RenounceOwnership() (*types.Transaction, error) {
	return _CharityWallet.Contract.RenounceOwnership(&_CharityWallet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CharityWallet *CharityWalletTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CharityWallet.Contract.RenounceOwnership(&_CharityWallet.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xb46300ec.
//
// Solidity: function send() returns()
func (_CharityWallet *CharityWalletTransactor) Send(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CharityWallet.contract.Transact(opts, "send")
}

// Send is a paid mutator transaction binding the contract method 0xb46300ec.
//
// Solidity: function send() returns()
func (_CharityWallet *CharityWalletSession) Send() (*types.Transaction, error) {
	return _CharityWallet.Contract.Send(&_CharityWallet.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xb46300ec.
//
// Solidity: function send() returns()
func (_CharityWallet *CharityWalletTransactorSession) Send() (*types.Transaction, error) {
	return _CharityWallet.Contract.Send(&_CharityWallet.TransactOpts)
}

// SetCharity is a paid mutator transaction binding the contract method 0xfb6f71a3.
//
// Solidity: function setCharity(address newCharityAddress) returns()
func (_CharityWallet *CharityWalletTransactor) SetCharity(opts *bind.TransactOpts, newCharityAddress common.Address) (*types.Transaction, error) {
	return _CharityWallet.contract.Transact(opts, "setCharity", newCharityAddress)
}

// SetCharity is a paid mutator transaction binding the contract method 0xfb6f71a3.
//
// Solidity: function setCharity(address newCharityAddress) returns()
func (_CharityWallet *CharityWalletSession) SetCharity(newCharityAddress common.Address) (*types.Transaction, error) {
	return _CharityWallet.Contract.SetCharity(&_CharityWallet.TransactOpts, newCharityAddress)
}

// SetCharity is a paid mutator transaction binding the contract method 0xfb6f71a3.
//
// Solidity: function setCharity(address newCharityAddress) returns()
func (_CharityWallet *CharityWalletTransactorSession) SetCharity(newCharityAddress common.Address) (*types.Transaction, error) {
	return _CharityWallet.Contract.SetCharity(&_CharityWallet.TransactOpts, newCharityAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CharityWallet *CharityWalletTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CharityWallet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CharityWallet *CharityWalletSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CharityWallet.Contract.TransferOwnership(&_CharityWallet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CharityWallet *CharityWalletTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CharityWallet.Contract.TransferOwnership(&_CharityWallet.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CharityWallet *CharityWalletTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CharityWallet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CharityWallet *CharityWalletSession) Receive() (*types.Transaction, error) {
	return _CharityWallet.Contract.Receive(&_CharityWallet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CharityWallet *CharityWalletTransactorSession) Receive() (*types.Transaction, error) {
	return _CharityWallet.Contract.Receive(&_CharityWallet.TransactOpts)
}

// CharityWalletCharityUpdatedEventIterator is returned from FilterCharityUpdatedEvent and is used to iterate over the raw logs and unpacked data for CharityUpdatedEvent events raised by the CharityWallet contract.
type CharityWalletCharityUpdatedEventIterator struct {
	Event *CharityWalletCharityUpdatedEvent // Event containing the contract specifics and raw log

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
func (it *CharityWalletCharityUpdatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharityWalletCharityUpdatedEvent)
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
		it.Event = new(CharityWalletCharityUpdatedEvent)
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
func (it *CharityWalletCharityUpdatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharityWalletCharityUpdatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharityWalletCharityUpdatedEvent represents a CharityUpdatedEvent event raised by the CharityWallet contract.
type CharityWalletCharityUpdatedEvent struct {
	NewCharityAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCharityUpdatedEvent is a free log retrieval operation binding the contract event 0xa0bd6b2fdbf082ae2356710c23fc8d76d56d418cecb4514d119c77a8617b4ffe.
//
// Solidity: event CharityUpdatedEvent(address indexed newCharityAddress)
func (_CharityWallet *CharityWalletFilterer) FilterCharityUpdatedEvent(opts *bind.FilterOpts, newCharityAddress []common.Address) (*CharityWalletCharityUpdatedEventIterator, error) {

	var newCharityAddressRule []interface{}
	for _, newCharityAddressItem := range newCharityAddress {
		newCharityAddressRule = append(newCharityAddressRule, newCharityAddressItem)
	}

	logs, sub, err := _CharityWallet.contract.FilterLogs(opts, "CharityUpdatedEvent", newCharityAddressRule)
	if err != nil {
		return nil, err
	}
	return &CharityWalletCharityUpdatedEventIterator{contract: _CharityWallet.contract, event: "CharityUpdatedEvent", logs: logs, sub: sub}, nil
}

// WatchCharityUpdatedEvent is a free log subscription operation binding the contract event 0xa0bd6b2fdbf082ae2356710c23fc8d76d56d418cecb4514d119c77a8617b4ffe.
//
// Solidity: event CharityUpdatedEvent(address indexed newCharityAddress)
func (_CharityWallet *CharityWalletFilterer) WatchCharityUpdatedEvent(opts *bind.WatchOpts, sink chan<- *CharityWalletCharityUpdatedEvent, newCharityAddress []common.Address) (event.Subscription, error) {

	var newCharityAddressRule []interface{}
	for _, newCharityAddressItem := range newCharityAddress {
		newCharityAddressRule = append(newCharityAddressRule, newCharityAddressItem)
	}

	logs, sub, err := _CharityWallet.contract.WatchLogs(opts, "CharityUpdatedEvent", newCharityAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharityWalletCharityUpdatedEvent)
				if err := _CharityWallet.contract.UnpackLog(event, "CharityUpdatedEvent", log); err != nil {
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

// ParseCharityUpdatedEvent is a log parse operation binding the contract event 0xa0bd6b2fdbf082ae2356710c23fc8d76d56d418cecb4514d119c77a8617b4ffe.
//
// Solidity: event CharityUpdatedEvent(address indexed newCharityAddress)
func (_CharityWallet *CharityWalletFilterer) ParseCharityUpdatedEvent(log types.Log) (*CharityWalletCharityUpdatedEvent, error) {
	event := new(CharityWalletCharityUpdatedEvent)
	if err := _CharityWallet.contract.UnpackLog(event, "CharityUpdatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharityWalletDonationReceivedEventIterator is returned from FilterDonationReceivedEvent and is used to iterate over the raw logs and unpacked data for DonationReceivedEvent events raised by the CharityWallet contract.
type CharityWalletDonationReceivedEventIterator struct {
	Event *CharityWalletDonationReceivedEvent // Event containing the contract specifics and raw log

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
func (it *CharityWalletDonationReceivedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharityWalletDonationReceivedEvent)
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
		it.Event = new(CharityWalletDonationReceivedEvent)
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
func (it *CharityWalletDonationReceivedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharityWalletDonationReceivedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharityWalletDonationReceivedEvent represents a DonationReceivedEvent event raised by the CharityWallet contract.
type CharityWalletDonationReceivedEvent struct {
	Donor  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDonationReceivedEvent is a free log retrieval operation binding the contract event 0x46ff3d75d4645bdbbae4cd6109ba42e6e1b80ea25e69d10472b357b733300368.
//
// Solidity: event DonationReceivedEvent(address indexed donor, uint256 amount)
func (_CharityWallet *CharityWalletFilterer) FilterDonationReceivedEvent(opts *bind.FilterOpts, donor []common.Address) (*CharityWalletDonationReceivedEventIterator, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _CharityWallet.contract.FilterLogs(opts, "DonationReceivedEvent", donorRule)
	if err != nil {
		return nil, err
	}
	return &CharityWalletDonationReceivedEventIterator{contract: _CharityWallet.contract, event: "DonationReceivedEvent", logs: logs, sub: sub}, nil
}

// WatchDonationReceivedEvent is a free log subscription operation binding the contract event 0x46ff3d75d4645bdbbae4cd6109ba42e6e1b80ea25e69d10472b357b733300368.
//
// Solidity: event DonationReceivedEvent(address indexed donor, uint256 amount)
func (_CharityWallet *CharityWalletFilterer) WatchDonationReceivedEvent(opts *bind.WatchOpts, sink chan<- *CharityWalletDonationReceivedEvent, donor []common.Address) (event.Subscription, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _CharityWallet.contract.WatchLogs(opts, "DonationReceivedEvent", donorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharityWalletDonationReceivedEvent)
				if err := _CharityWallet.contract.UnpackLog(event, "DonationReceivedEvent", log); err != nil {
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

// ParseDonationReceivedEvent is a log parse operation binding the contract event 0x46ff3d75d4645bdbbae4cd6109ba42e6e1b80ea25e69d10472b357b733300368.
//
// Solidity: event DonationReceivedEvent(address indexed donor, uint256 amount)
func (_CharityWallet *CharityWalletFilterer) ParseDonationReceivedEvent(log types.Log) (*CharityWalletDonationReceivedEvent, error) {
	event := new(CharityWalletDonationReceivedEvent)
	if err := _CharityWallet.contract.UnpackLog(event, "DonationReceivedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharityWalletDonationSentEventIterator is returned from FilterDonationSentEvent and is used to iterate over the raw logs and unpacked data for DonationSentEvent events raised by the CharityWallet contract.
type CharityWalletDonationSentEventIterator struct {
	Event *CharityWalletDonationSentEvent // Event containing the contract specifics and raw log

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
func (it *CharityWalletDonationSentEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharityWalletDonationSentEvent)
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
		it.Event = new(CharityWalletDonationSentEvent)
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
func (it *CharityWalletDonationSentEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharityWalletDonationSentEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharityWalletDonationSentEvent represents a DonationSentEvent event raised by the CharityWallet contract.
type CharityWalletDonationSentEvent struct {
	Charity common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDonationSentEvent is a free log retrieval operation binding the contract event 0x44d398d152fa0735a428b13ebc78f79fe4cb1b4722292cd233e278552fa36d32.
//
// Solidity: event DonationSentEvent(address indexed charity, uint256 amount)
func (_CharityWallet *CharityWalletFilterer) FilterDonationSentEvent(opts *bind.FilterOpts, charity []common.Address) (*CharityWalletDonationSentEventIterator, error) {

	var charityRule []interface{}
	for _, charityItem := range charity {
		charityRule = append(charityRule, charityItem)
	}

	logs, sub, err := _CharityWallet.contract.FilterLogs(opts, "DonationSentEvent", charityRule)
	if err != nil {
		return nil, err
	}
	return &CharityWalletDonationSentEventIterator{contract: _CharityWallet.contract, event: "DonationSentEvent", logs: logs, sub: sub}, nil
}

// WatchDonationSentEvent is a free log subscription operation binding the contract event 0x44d398d152fa0735a428b13ebc78f79fe4cb1b4722292cd233e278552fa36d32.
//
// Solidity: event DonationSentEvent(address indexed charity, uint256 amount)
func (_CharityWallet *CharityWalletFilterer) WatchDonationSentEvent(opts *bind.WatchOpts, sink chan<- *CharityWalletDonationSentEvent, charity []common.Address) (event.Subscription, error) {

	var charityRule []interface{}
	for _, charityItem := range charity {
		charityRule = append(charityRule, charityItem)
	}

	logs, sub, err := _CharityWallet.contract.WatchLogs(opts, "DonationSentEvent", charityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharityWalletDonationSentEvent)
				if err := _CharityWallet.contract.UnpackLog(event, "DonationSentEvent", log); err != nil {
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

// ParseDonationSentEvent is a log parse operation binding the contract event 0x44d398d152fa0735a428b13ebc78f79fe4cb1b4722292cd233e278552fa36d32.
//
// Solidity: event DonationSentEvent(address indexed charity, uint256 amount)
func (_CharityWallet *CharityWalletFilterer) ParseDonationSentEvent(log types.Log) (*CharityWalletDonationSentEvent, error) {
	event := new(CharityWalletDonationSentEvent)
	if err := _CharityWallet.contract.UnpackLog(event, "DonationSentEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharityWalletOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CharityWallet contract.
type CharityWalletOwnershipTransferredIterator struct {
	Event *CharityWalletOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CharityWalletOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharityWalletOwnershipTransferred)
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
		it.Event = new(CharityWalletOwnershipTransferred)
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
func (it *CharityWalletOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharityWalletOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharityWalletOwnershipTransferred represents a OwnershipTransferred event raised by the CharityWallet contract.
type CharityWalletOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CharityWallet *CharityWalletFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CharityWalletOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CharityWallet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CharityWalletOwnershipTransferredIterator{contract: _CharityWallet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CharityWallet *CharityWalletFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CharityWalletOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CharityWallet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharityWalletOwnershipTransferred)
				if err := _CharityWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CharityWallet *CharityWalletFilterer) ParseOwnershipTransferred(log types.Log) (*CharityWalletOwnershipTransferred, error) {
	event := new(CharityWalletOwnershipTransferred)
	if err := _CharityWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
