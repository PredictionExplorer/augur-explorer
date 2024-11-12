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

// IStakingWalletRandomWalkNftMetaData contains all meta data concerning the IStakingWalletRandomWalkNft contract.
var IStakingWalletRandomWalkNftMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCosmicGameConstants.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftUnstaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"entropy_\",\"type\":\"bytes32\"}],\"name\":\"pickRandomStakerIfPossible\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IStakingWalletRandomWalkNftABI is the input ABI used to generate the binding from.
// Deprecated: Use IStakingWalletRandomWalkNftMetaData.ABI instead.
var IStakingWalletRandomWalkNftABI = IStakingWalletRandomWalkNftMetaData.ABI

// IStakingWalletRandomWalkNft is an auto generated Go binding around an Ethereum contract.
type IStakingWalletRandomWalkNft struct {
	IStakingWalletRandomWalkNftCaller     // Read-only binding to the contract
	IStakingWalletRandomWalkNftTransactor // Write-only binding to the contract
	IStakingWalletRandomWalkNftFilterer   // Log filterer for contract events
}

// IStakingWalletRandomWalkNftCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakingWalletRandomWalkNftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletRandomWalkNftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakingWalletRandomWalkNftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletRandomWalkNftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakingWalletRandomWalkNftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletRandomWalkNftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakingWalletRandomWalkNftSession struct {
	Contract     *IStakingWalletRandomWalkNft // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IStakingWalletRandomWalkNftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakingWalletRandomWalkNftCallerSession struct {
	Contract *IStakingWalletRandomWalkNftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// IStakingWalletRandomWalkNftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakingWalletRandomWalkNftTransactorSession struct {
	Contract     *IStakingWalletRandomWalkNftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// IStakingWalletRandomWalkNftRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakingWalletRandomWalkNftRaw struct {
	Contract *IStakingWalletRandomWalkNft // Generic contract binding to access the raw methods on
}

// IStakingWalletRandomWalkNftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakingWalletRandomWalkNftCallerRaw struct {
	Contract *IStakingWalletRandomWalkNftCaller // Generic read-only contract binding to access the raw methods on
}

// IStakingWalletRandomWalkNftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakingWalletRandomWalkNftTransactorRaw struct {
	Contract *IStakingWalletRandomWalkNftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStakingWalletRandomWalkNft creates a new instance of IStakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewIStakingWalletRandomWalkNft(address common.Address, backend bind.ContractBackend) (*IStakingWalletRandomWalkNft, error) {
	contract, err := bindIStakingWalletRandomWalkNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRandomWalkNft{IStakingWalletRandomWalkNftCaller: IStakingWalletRandomWalkNftCaller{contract: contract}, IStakingWalletRandomWalkNftTransactor: IStakingWalletRandomWalkNftTransactor{contract: contract}, IStakingWalletRandomWalkNftFilterer: IStakingWalletRandomWalkNftFilterer{contract: contract}}, nil
}

// NewIStakingWalletRandomWalkNftCaller creates a new read-only instance of IStakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewIStakingWalletRandomWalkNftCaller(address common.Address, caller bind.ContractCaller) (*IStakingWalletRandomWalkNftCaller, error) {
	contract, err := bindIStakingWalletRandomWalkNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRandomWalkNftCaller{contract: contract}, nil
}

// NewIStakingWalletRandomWalkNftTransactor creates a new write-only instance of IStakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewIStakingWalletRandomWalkNftTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakingWalletRandomWalkNftTransactor, error) {
	contract, err := bindIStakingWalletRandomWalkNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRandomWalkNftTransactor{contract: contract}, nil
}

// NewIStakingWalletRandomWalkNftFilterer creates a new log filterer instance of IStakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewIStakingWalletRandomWalkNftFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakingWalletRandomWalkNftFilterer, error) {
	contract, err := bindIStakingWalletRandomWalkNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRandomWalkNftFilterer{contract: contract}, nil
}

// bindIStakingWalletRandomWalkNft binds a generic wrapper to an already deployed contract.
func bindIStakingWalletRandomWalkNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStakingWalletRandomWalkNftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingWalletRandomWalkNft.Contract.IStakingWalletRandomWalkNftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.IStakingWalletRandomWalkNftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.IStakingWalletRandomWalkNftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingWalletRandomWalkNft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.contract.Transact(opts, method, params...)
}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftCaller) NumStakedNfts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakingWalletRandomWalkNft.contract.Call(opts, &out, "numStakedNfts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) NumStakedNfts() (*big.Int, error) {
	return _IStakingWalletRandomWalkNft.Contract.NumStakedNfts(&_IStakingWalletRandomWalkNft.CallOpts)
}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftCallerSession) NumStakedNfts() (*big.Int, error) {
	return _IStakingWalletRandomWalkNft.Contract.NumStakedNfts(&_IStakingWalletRandomWalkNft.CallOpts)
}

// PickRandomStakerIfPossible is a free data retrieval call binding the contract method 0x00caccb0.
//
// Solidity: function pickRandomStakerIfPossible(bytes32 entropy_) view returns(address)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftCaller) PickRandomStakerIfPossible(opts *bind.CallOpts, entropy_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IStakingWalletRandomWalkNft.contract.Call(opts, &out, "pickRandomStakerIfPossible", entropy_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PickRandomStakerIfPossible is a free data retrieval call binding the contract method 0x00caccb0.
//
// Solidity: function pickRandomStakerIfPossible(bytes32 entropy_) view returns(address)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) PickRandomStakerIfPossible(entropy_ [32]byte) (common.Address, error) {
	return _IStakingWalletRandomWalkNft.Contract.PickRandomStakerIfPossible(&_IStakingWalletRandomWalkNft.CallOpts, entropy_)
}

// PickRandomStakerIfPossible is a free data retrieval call binding the contract method 0x00caccb0.
//
// Solidity: function pickRandomStakerIfPossible(bytes32 entropy_) view returns(address)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftCallerSession) PickRandomStakerIfPossible(entropy_ [32]byte) (common.Address, error) {
	return _IStakingWalletRandomWalkNft.Contract.PickRandomStakerIfPossible(&_IStakingWalletRandomWalkNft.CallOpts, entropy_)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(bool)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftCaller) WasNftUsed(opts *bind.CallOpts, nftId_ *big.Int) (bool, error) {
	var out []interface{}
	err := _IStakingWalletRandomWalkNft.contract.Call(opts, &out, "wasNftUsed", nftId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(bool)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) WasNftUsed(nftId_ *big.Int) (bool, error) {
	return _IStakingWalletRandomWalkNft.Contract.WasNftUsed(&_IStakingWalletRandomWalkNft.CallOpts, nftId_)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(bool)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftCallerSession) WasNftUsed(nftId_ *big.Int) (bool, error) {
	return _IStakingWalletRandomWalkNft.Contract.WasNftUsed(&_IStakingWalletRandomWalkNft.CallOpts, nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactor) Stake(opts *bind.TransactOpts, nftId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.contract.Transact(opts, "stake", nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.Stake(&_IStakingWalletRandomWalkNft.TransactOpts, nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactorSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.Stake(&_IStakingWalletRandomWalkNft.TransactOpts, nftId_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactor) StakeMany(opts *bind.TransactOpts, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.contract.Transact(opts, "stakeMany", nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.StakeMany(&_IStakingWalletRandomWalkNft.TransactOpts, nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactorSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.StakeMany(&_IStakingWalletRandomWalkNft.TransactOpts, nftIds_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactor) Unstake(opts *bind.TransactOpts, stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.contract.Transact(opts, "unstake", stakeActionId_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) Unstake(stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.Unstake(&_IStakingWalletRandomWalkNft.TransactOpts, stakeActionId_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactorSession) Unstake(stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.Unstake(&_IStakingWalletRandomWalkNft.TransactOpts, stakeActionId_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactor) UnstakeMany(opts *bind.TransactOpts, stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.contract.Transact(opts, "unstakeMany", stakeActionIds_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) UnstakeMany(stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.UnstakeMany(&_IStakingWalletRandomWalkNft.TransactOpts, stakeActionIds_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactorSession) UnstakeMany(stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.UnstakeMany(&_IStakingWalletRandomWalkNft.TransactOpts, stakeActionIds_)
}

// IStakingWalletRandomWalkNftNftStakedIterator is returned from FilterNftStaked and is used to iterate over the raw logs and unpacked data for NftStaked events raised by the IStakingWalletRandomWalkNft contract.
type IStakingWalletRandomWalkNftNftStakedIterator struct {
	Event *IStakingWalletRandomWalkNftNftStaked // Event containing the contract specifics and raw log

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
func (it *IStakingWalletRandomWalkNftNftStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingWalletRandomWalkNftNftStaked)
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
		it.Event = new(IStakingWalletRandomWalkNftNftStaked)
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
func (it *IStakingWalletRandomWalkNftNftStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingWalletRandomWalkNftNftStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingWalletRandomWalkNftNftStaked represents a NftStaked event raised by the IStakingWalletRandomWalkNft contract.
type IStakingWalletRandomWalkNftNftStaked struct {
	StakeActionId *big.Int
	NftTypeCode   uint8
	NftId         *big.Int
	StakerAddress common.Address
	NumStakedNfts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNftStaked is a free log retrieval operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftFilterer) FilterNftStaked(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*IStakingWalletRandomWalkNftNftStakedIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletRandomWalkNft.contract.FilterLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRandomWalkNftNftStakedIterator{contract: _IStakingWalletRandomWalkNft.contract, event: "NftStaked", logs: logs, sub: sub}, nil
}

// WatchNftStaked is a free log subscription operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftFilterer) WatchNftStaked(opts *bind.WatchOpts, sink chan<- *IStakingWalletRandomWalkNftNftStaked, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletRandomWalkNft.contract.WatchLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingWalletRandomWalkNftNftStaked)
				if err := _IStakingWalletRandomWalkNft.contract.UnpackLog(event, "NftStaked", log); err != nil {
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

// ParseNftStaked is a log parse operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftFilterer) ParseNftStaked(log types.Log) (*IStakingWalletRandomWalkNftNftStaked, error) {
	event := new(IStakingWalletRandomWalkNftNftStaked)
	if err := _IStakingWalletRandomWalkNft.contract.UnpackLog(event, "NftStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingWalletRandomWalkNftNftUnstakedIterator is returned from FilterNftUnstaked and is used to iterate over the raw logs and unpacked data for NftUnstaked events raised by the IStakingWalletRandomWalkNft contract.
type IStakingWalletRandomWalkNftNftUnstakedIterator struct {
	Event *IStakingWalletRandomWalkNftNftUnstaked // Event containing the contract specifics and raw log

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
func (it *IStakingWalletRandomWalkNftNftUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingWalletRandomWalkNftNftUnstaked)
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
		it.Event = new(IStakingWalletRandomWalkNftNftUnstaked)
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
func (it *IStakingWalletRandomWalkNftNftUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingWalletRandomWalkNftNftUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingWalletRandomWalkNftNftUnstaked represents a NftUnstaked event raised by the IStakingWalletRandomWalkNft contract.
type IStakingWalletRandomWalkNftNftUnstaked struct {
	StakeActionId *big.Int
	NftId         *big.Int
	StakerAddress common.Address
	NumStakedNfts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNftUnstaked is a free log retrieval operation binding the contract event 0x1792a7a9d5e2f53a8d06f2ae40b4446d91b503e84dd7f6307f40cdeb11541668.
//
// Solidity: event NftUnstaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftFilterer) FilterNftUnstaked(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*IStakingWalletRandomWalkNftNftUnstakedIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletRandomWalkNft.contract.FilterLogs(opts, "NftUnstaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRandomWalkNftNftUnstakedIterator{contract: _IStakingWalletRandomWalkNft.contract, event: "NftUnstaked", logs: logs, sub: sub}, nil
}

// WatchNftUnstaked is a free log subscription operation binding the contract event 0x1792a7a9d5e2f53a8d06f2ae40b4446d91b503e84dd7f6307f40cdeb11541668.
//
// Solidity: event NftUnstaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftFilterer) WatchNftUnstaked(opts *bind.WatchOpts, sink chan<- *IStakingWalletRandomWalkNftNftUnstaked, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletRandomWalkNft.contract.WatchLogs(opts, "NftUnstaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingWalletRandomWalkNftNftUnstaked)
				if err := _IStakingWalletRandomWalkNft.contract.UnpackLog(event, "NftUnstaked", log); err != nil {
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

// ParseNftUnstaked is a log parse operation binding the contract event 0x1792a7a9d5e2f53a8d06f2ae40b4446d91b503e84dd7f6307f40cdeb11541668.
//
// Solidity: event NftUnstaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftFilterer) ParseNftUnstaked(log types.Log) (*IStakingWalletRandomWalkNftNftUnstaked, error) {
	event := new(IStakingWalletRandomWalkNftNftUnstaked)
	if err := _IStakingWalletRandomWalkNft.contract.UnpackLog(event, "NftUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletRandomWalkNftMetaData contains all meta data concerning the StakingWalletRandomWalkNft contract.
var StakingWalletRandomWalkNftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"randomWalkNft_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"NftOneTimeStaking\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"}],\"name\":\"NftStakeActionAccessDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"NftStakeActionInvalidId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCosmicGameConstants.NftTypeCode\",\"name\":\"nftTypeCode\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftUnstaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"actionCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"entropy_\",\"type\":\"bytes32\"}],\"name\":\"pickRandomStakerIfPossible\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomWalkNft\",\"outputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeActionIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftOwnerAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"wasNftUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523461002f57610019610014610100565b610290565b610021610034565b6116206102d3823961162090f35b61003a565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100669061003e565b810190811060018060401b0382111761007e57604052565b610048565b9061009661008f610034565b928361005c565b565b5f80fd5b60018060a01b031690565b6100b09061009c565b90565b6100bc906100a7565b90565b6100c8816100b3565b036100cf57565b5f80fd5b905051906100e0826100bf565b565b906020828203126100fb576100f8915f016100d3565b90565b610098565b61011e6118f38038038061011381610083565b9283398101906100e2565b90565b90565b61013861013361013d9261009c565b610121565b61009c565b90565b61014990610124565b90565b61015590610140565b90565b90565b61016f61016a61017492610158565b610121565b61009c565b90565b6101809061015b565b90565b60209181520190565b60207f616e646f6d57616c6b4e66745f2e000000000000000000000000000000000000917f5a65726f2d616464726573732077617320676976656e20666f722074686520725f8201520152565b6101e6602e604092610183565b6101ef8161018c565b0190565b6102089060208101905f8183039101526101d9565b90565b1561021257565b61021a610034565b63eac0d38960e01b815280610231600482016101f3565b0390fd5b5f1b90565b9061024b60018060a01b0391610235565b9181191691161790565b61025e90610124565b90565b61026a90610255565b90565b90565b9061028561028061028c92610261565b61026d565b825461023a565b9055565b6102d0906102c16102a08261014c565b6102ba6102b46102af5f610177565b6100a7565b916100a7565b141561020b565b68010000000000000002610270565b56fe60806040526004361015610013575b610739565b61001d5f356100cb565b8062caccb0146100c65780630d50c189146100c15780632e17de78146100bc57806357951c74146100b757806360294405146100b2578063755b4ef7146100ad578063a2b136fb146100a8578063a694fc3a146100a3578063ca7c1f921461009e578063fdbd98b0146100995763fe939afc0361000e57610705565b6106d0565b610684565b610651565b610619565b6104d7565b6103f2565b6102ed565b610293565b61021b565b610160565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b90565b6100ef816100e3565b036100f657565b5f80fd5b90503590610107826100e6565b565b906020828203126101225761011f915f016100fa565b90565b6100db565b60018060a01b031690565b61013b90610127565b90565b61014790610132565b9052565b919061015e905f6020850194019061013e565b565b346101905761018c61017b610176366004610109565b6107ee565b6101836100d1565b9182918261014b565b0390f35b6100d7565b5f80fd5b5f80fd5b5f80fd5b909182601f830112156101db5781359167ffffffffffffffff83116101d65760200192602083028401116101d157565b61019d565b610199565b610195565b90602082820312610211575f82013567ffffffffffffffff811161020c5761020892016101a1565b9091565b6100df565b6100db565b5f0190565b3461024a5761023461022e3660046101e0565b906108aa565b61023c6100d1565b8061024681610216565b0390f35b6100d7565b90565b61025b8161024f565b0361026257565b5f80fd5b9050359061027382610252565b565b9060208282031261028e5761028b915f01610266565b90565b6100db565b346102c1576102ab6102a6366004610275565b610d02565b6102b36100d1565b806102bd81610216565b0390f35b6100d7565b151590565b6102d4906102c6565b9052565b91906102eb905f602085019401906102cb565b565b3461031d57610319610308610303366004610275565b611019565b6103106100d1565b918291826102d8565b0390f35b6100d7565b634e487b7160e01b5f52603260045260245ffd5b506801000000000000000090565b90565b61035081610336565b82101561036a57610362600191610344565b910201905f90565b610322565b1c90565b90565b61038690600861038b930261036f565b610373565b90565b906103999154610376565b90565b680400000000000000036103af81610336565b8210156103cc576103c9916103c391610347565b9061038e565b90565b5f80fd5b6103d99061024f565b9052565b91906103f0905f602085019401906103d0565b565b346104225761041e61040d610408366004610275565b61039c565b6104156100d1565b918291826103dd565b0390f35b6100d7565b5f91031261043157565b6100db565b60018060a01b031690565b610451906008610456930261036f565b610436565b90565b906104649154610441565b90565b61047b680100000000000000025f90610459565b90565b90565b61049561049061049a92610127565b61047e565b610127565b90565b6104a690610481565b90565b6104b29061049d565b90565b6104be906104a9565b9052565b91906104d5905f602085019401906104b5565b565b34610507576104e7366004610427565b6105036104f2610467565b6104fa6100d1565b918291826104c2565b0390f35b6100d7565b506801000000000000000090565b90565b6105268161050c565b8210156105405761053860039161051a565b910201905f90565b610322565b5f1c90565b61055661055b91610545565b610373565b90565b610568905461054a565b90565b60018060a01b031690565b61058261058791610545565b61056b565b90565b6105949054610576565b90565b68010000000000000003906105ab8261050c565b8110156105e3576105bb9161051d565b506105c75f820161055e565b916105e060026105d96001850161055e565b930161058a565b90565b5f80fd5b604090610610610617949695939661060660608401985f8501906103d0565b60208301906103d0565b019061013e565b565b3461064c5761064861063461062f366004610275565b610597565b61063f9391936100d1565b938493846105e7565b0390f35b6100d7565b3461067f57610669610664366004610275565b61134c565b6106716100d1565b8061067b81610216565b0390f35b6100d7565b346106b457610694366004610427565b6106b061069f611575565b6106a76100d1565b918291826103dd565b0390f35b6100d7565b6106cd680100000000000000015f9061038e565b90565b34610700576106e0366004610427565b6106fc6106eb6106b9565b6106f36100d1565b918291826103dd565b0390f35b6100d7565b346107345761071e6107183660046101e0565b9061158a565b6107266100d1565b8061073081610216565b0390f35b6100d7565b5f80fd5b5f90565b90565b61075861075361075d92610741565b61047e565b61024f565b90565b61077461076f61077992610741565b61047e565b610127565b90565b61078590610760565b90565b61079c6107976107a19261024f565b61047e565b61024f565b90565b6107b06107b591610545565b610788565b90565b634e487b7160e01b5f52601260045260245ffd5b6107d86107de9161024f565b9161024f565b9081156107e9570690565b6107b8565b6107f661073d565b506108005f61055e565b908161081461080e5f610744565b9161024f565b146108675761085d61084e610848610839610864956108346002966107a4565b6107cc565b68040000000000000003610347565b9061038e565b6801000000000000000361051d565b500161058a565b90565b50506108725f61077c565b90565b6001610881910161024f565b90565b5090565b9190811015610898576020020190565b610322565b356108a781610252565b90565b9190916108b65f610744565b5b806108d46108ce6108c9858890610884565b61024f565b9161024f565b1015610904576108ff906108fa6108f56108f085888591610888565b61089d565b610d02565b610875565b6108b7565b50509050565b90565b906109179061024f565b9052565b9061092590610132565b9052565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b9061095190610929565b810190811067ffffffffffffffff82111761096b57604052565b610933565b9061098361097c6100d1565b9283610947565b565b61098f6060610970565b90565b906109e16109d860026109a3610985565b946109ba6109b25f830161055e565b5f880161090d565b6109d26109c96001830161055e565b6020880161090d565b0161058a565b6040840161091b565b565b6109ec90610992565b90565b6109f99051610132565b90565b60209181520190565b5f7f496e76616c6964204e4654207374616b6520616374696f6e2049442e00000000910152565b610a39601c6020926109fc565b610a4281610a05565b0190565b9190610a69906020610a61604086018681035f880152610a2c565b9401906103d0565b565b60207f6e7374616b652069742e00000000000000000000000000000000000000000000917f4f6e6c79204e4654206f776e6572206973207065726d697474656420746f20755f8201520152565b610ac5602a6040926109fc565b610ace81610a6b565b0190565b916040610b03929493610afc610af1606083018381035f850152610ab8565b9660208301906103d0565b019061013e565b565b90565b610b1c610b17610b2192610b05565b61047e565b61024f565b90565b634e487b7160e01b5f52601160045260245ffd5b610b47610b4d9193929361024f565b9261024f565b8203918211610b5857565b610b24565b5f1b90565b90610b6e5f1991610b5d565b9181191691161790565b90565b90610b90610b8b610b9792610788565b610b78565b8254610b62565b9055565b610ba5905161024f565b90565b1b90565b91906008610bc7910291610bc15f1984610ba8565b92610ba8565b9181191691161790565b9190610be7610be2610bef93610788565b610b78565b908354610bac565b9055565b5f90565b610c0991610c03610bf3565b91610bd1565b565b91906008610c2b910291610c2560018060a01b0384610ba8565b92610ba8565b9181191691161790565b610c3e9061049d565b90565b90565b9190610c5a610c55610c6293610c35565b610c41565b908354610c0b565b9055565b610c7891610c7261073d565b91610c44565b565b610c86610c8b91610545565b610436565b90565b610c989054610c7a565b90565b610ca49061049d565b90565b5f80fd5b60e01b90565b5f910312610cbb57565b6100db565b604090610ce9610cf09496959396610cdf60608401985f85019061013e565b602083019061013e565b01906103d0565b565b610cfa6100d1565b3d5f823e3d90fd5b610d1f610d1968010000000000000003839061051d565b5061090a565b610d28816109e3565b9133610d47610d41610d3c604087016109ef565b610132565b91610132565b03610f3a57610e07610d6b610d5b5f61055e565b610d656001610b08565b90610b38565b92610d76845f610b7b565b610de35f6002610d99610d93680400000000000000038990610347565b9061038e565b93610dc4610da8848b01610b9b565b84610dbd68010000000000000003899061051d565b5001610b7b565b610dd083808301610bf7565b610ddd8360018301610bf7565b01610c66565b610e0168040000000000000003610dfb5f8801610b9b565b90610347565b90610bd1565b610e24610e1e680400000000000000038490610347565b90610bf7565b610e3e610e3968010000000000000002610c8e565b6104a9565b906323b872dd610e4d30610c9b565b3393610e5b60208801610b9b565b92813b15610f35575f610e8191610e8c8296610e756100d1565b998a9788968795610cab565b855260048501610cc0565b03925af1908115610f3057610eab92602092610f04575b509301610b9b565b339192610eff610eed610ee7610ee17f1792a7a9d5e2f53a8d06f2ae40b4446d91b503e84dd7f6307f40cdeb1154166894610788565b94610788565b94610c35565b94610ef66100d1565b918291826103dd565b0390a4565b610f23905f3d8111610f29575b610f1b8183610947565b810190610cb1565b5f610ea3565b503d610f11565b610cf2565b610ca7565b610f46604084016109ef565b610f60610f5a610f555f61077c565b610132565b91610132565b14155f14610f8f573390610f8b610f756100d1565b9283926348aca7ef60e11b845260048401610ad2565b0390fd5b610fb190610f9b6100d1565b91829163023df6b160e21b835260048301610a46565b0390fd5b5f90565b506801000000000000000090565b90565b610fd381610fb9565b821015610fed57610fe5600191610fc7565b910201905f90565b610322565b60ff1690565b61100461100991610545565b610ff2565b90565b6110169054610ff8565b90565b5f61103161103892611029610fb5565b506001610fca565b500161100c565b90565b60407f206f6e6c79206f6e63652e000000000000000000000000000000000000000000917f54686973204e46542068617320616c7265616479206265656e207374616b65645f8201527f2e20416e204e465420697320616c6c6f77656420746f206265207374616b656460208201520152565b6110bb604b6060926109fc565b6110c48161103b565b0190565b91906110eb9060206110e3604086018681035f8801526110ae565b9401906103d0565b565b156110f55750565b611117906111016100d1565b918291633b471d1f60e21b8352600483016110c8565b0390fd5b61112a6111309193929361024f565b9261024f565b820180921161113b57565b610b24565b9061115160018060a01b0391610b5d565b9181191691161790565b9061117061116b61117792610c35565b610c41565b8254611140565b9055565b6111856040610970565b90565b90611192906102c6565b9052565b60018060f81b031690565b6111b56111b06111ba92610741565b61047e565b611196565b90565b906111c790611196565b9052565b634e487b7160e01b5f525f60045260245ffd5b6111e890516102c6565b90565b906111f760ff91610b5d565b9181191691161790565b61120a906102c6565b90565b90565b9061122561122061122c92611201565b61120d565b82546111eb565b9055565b61123a9051611196565b90565b60081b90565b9061125060ff199161123d565b9181191691161790565b61126e61126961127392611196565b61047e565b611196565b90565b90565b9061128e6112896112959261125a565b611276565b8254611243565b9055565b906112c360205f6112c9946112bb8282016112b58488016111de565b90611210565b019201611230565b90611279565b565b91906112dc576112da91611299565b565b6111cb565b634e487b7160e01b5f52602160045260245ffd5b600311156112ff57565b6112e1565b9061130e826112f5565b565b61131990611304565b90565b61132590611310565b9052565b91602061134a92949361134360408201965f83019061131c565b01906103d0565b565b61137661136f6113695f61136260018690610fca565b500161100c565b156102c6565b82906110ed565b61139b61138b6801000000000000000161055e565b6113956001610b08565b9061111b565b6113ae8168010000000000000001610b7b565b61142c6113ce6113c868010000000000000003849061051d565b5061090a565b6113db8460018301610b7b565b6113e8336002830161115b565b6113fe6113f45f61055e565b915f839101610b7b565b61141c83611416680400000000000000038490610347565b90610bd1565b6114266001610b08565b9061111b565b90611437825f610b7b565b611477600161146561145c5f61145761144e61117b565b945f8601611188565b6111a1565b602083016111bd565b61147160018690610fca565b906112cb565b61149161148c68010000000000000002610c8e565b6104a9565b6323b872dd336114a030610c9b565b928692813b15611570575f6114c8916114d382966114bc6100d1565b98899788968795610cab565b855260048501610cc0565b03925af1801561156b5761153f575b5060029290339261152561151f6115197fcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f0082994610788565b94610788565b94610c35565b9461153a6115316100d1565b92839283611329565b0390a4565b61155e905f3d8111611564575b6115568183610947565b810190610cb1565b5f6114e2565b503d61154c565b610cf2565b610ca7565b61157d610bf3565b506115875f61055e565b90565b9190916115965f610744565b5b806115b46115ae6115a9858890610884565b61024f565b9161024f565b10156115e4576115df906115da6115d56115d085888591610888565b61089d565b61134c565b610875565b611597565b5050905056fea264697066735822122016b66b335932186a9417d18063e557700276e70243049f72505d4e283635645164736f6c634300081b0033",
}

// StakingWalletRandomWalkNftABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingWalletRandomWalkNftMetaData.ABI instead.
var StakingWalletRandomWalkNftABI = StakingWalletRandomWalkNftMetaData.ABI

// StakingWalletRandomWalkNftBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingWalletRandomWalkNftMetaData.Bin instead.
var StakingWalletRandomWalkNftBin = StakingWalletRandomWalkNftMetaData.Bin

// DeployStakingWalletRandomWalkNft deploys a new Ethereum contract, binding an instance of StakingWalletRandomWalkNft to it.
func DeployStakingWalletRandomWalkNft(auth *bind.TransactOpts, backend bind.ContractBackend, randomWalkNft_ common.Address) (common.Address, *types.Transaction, *StakingWalletRandomWalkNft, error) {
	parsed, err := StakingWalletRandomWalkNftMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingWalletRandomWalkNftBin), backend, randomWalkNft_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingWalletRandomWalkNft{StakingWalletRandomWalkNftCaller: StakingWalletRandomWalkNftCaller{contract: contract}, StakingWalletRandomWalkNftTransactor: StakingWalletRandomWalkNftTransactor{contract: contract}, StakingWalletRandomWalkNftFilterer: StakingWalletRandomWalkNftFilterer{contract: contract}}, nil
}

// StakingWalletRandomWalkNft is an auto generated Go binding around an Ethereum contract.
type StakingWalletRandomWalkNft struct {
	StakingWalletRandomWalkNftCaller     // Read-only binding to the contract
	StakingWalletRandomWalkNftTransactor // Write-only binding to the contract
	StakingWalletRandomWalkNftFilterer   // Log filterer for contract events
}

// StakingWalletRandomWalkNftCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingWalletRandomWalkNftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletRandomWalkNftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingWalletRandomWalkNftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletRandomWalkNftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingWalletRandomWalkNftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletRandomWalkNftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingWalletRandomWalkNftSession struct {
	Contract     *StakingWalletRandomWalkNft // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// StakingWalletRandomWalkNftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingWalletRandomWalkNftCallerSession struct {
	Contract *StakingWalletRandomWalkNftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// StakingWalletRandomWalkNftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingWalletRandomWalkNftTransactorSession struct {
	Contract     *StakingWalletRandomWalkNftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// StakingWalletRandomWalkNftRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingWalletRandomWalkNftRaw struct {
	Contract *StakingWalletRandomWalkNft // Generic contract binding to access the raw methods on
}

// StakingWalletRandomWalkNftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingWalletRandomWalkNftCallerRaw struct {
	Contract *StakingWalletRandomWalkNftCaller // Generic read-only contract binding to access the raw methods on
}

// StakingWalletRandomWalkNftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingWalletRandomWalkNftTransactorRaw struct {
	Contract *StakingWalletRandomWalkNftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingWalletRandomWalkNft creates a new instance of StakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewStakingWalletRandomWalkNft(address common.Address, backend bind.ContractBackend) (*StakingWalletRandomWalkNft, error) {
	contract, err := bindStakingWalletRandomWalkNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRandomWalkNft{StakingWalletRandomWalkNftCaller: StakingWalletRandomWalkNftCaller{contract: contract}, StakingWalletRandomWalkNftTransactor: StakingWalletRandomWalkNftTransactor{contract: contract}, StakingWalletRandomWalkNftFilterer: StakingWalletRandomWalkNftFilterer{contract: contract}}, nil
}

// NewStakingWalletRandomWalkNftCaller creates a new read-only instance of StakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewStakingWalletRandomWalkNftCaller(address common.Address, caller bind.ContractCaller) (*StakingWalletRandomWalkNftCaller, error) {
	contract, err := bindStakingWalletRandomWalkNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRandomWalkNftCaller{contract: contract}, nil
}

// NewStakingWalletRandomWalkNftTransactor creates a new write-only instance of StakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewStakingWalletRandomWalkNftTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingWalletRandomWalkNftTransactor, error) {
	contract, err := bindStakingWalletRandomWalkNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRandomWalkNftTransactor{contract: contract}, nil
}

// NewStakingWalletRandomWalkNftFilterer creates a new log filterer instance of StakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewStakingWalletRandomWalkNftFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingWalletRandomWalkNftFilterer, error) {
	contract, err := bindStakingWalletRandomWalkNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRandomWalkNftFilterer{contract: contract}, nil
}

// bindStakingWalletRandomWalkNft binds a generic wrapper to an already deployed contract.
func bindStakingWalletRandomWalkNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingWalletRandomWalkNftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWalletRandomWalkNft.Contract.StakingWalletRandomWalkNftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.StakingWalletRandomWalkNftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.StakingWalletRandomWalkNftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWalletRandomWalkNft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.contract.Transact(opts, method, params...)
}

// ActionCounter is a free data retrieval call binding the contract method 0xfdbd98b0.
//
// Solidity: function actionCounter() view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) ActionCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "actionCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActionCounter is a free data retrieval call binding the contract method 0xfdbd98b0.
//
// Solidity: function actionCounter() view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) ActionCounter() (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.ActionCounter(&_StakingWalletRandomWalkNft.CallOpts)
}

// ActionCounter is a free data retrieval call binding the contract method 0xfdbd98b0.
//
// Solidity: function actionCounter() view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) ActionCounter() (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.ActionCounter(&_StakingWalletRandomWalkNft.CallOpts)
}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) NumStakedNfts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "numStakedNfts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) NumStakedNfts() (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.NumStakedNfts(&_StakingWalletRandomWalkNft.CallOpts)
}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) NumStakedNfts() (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.NumStakedNfts(&_StakingWalletRandomWalkNft.CallOpts)
}

// PickRandomStakerIfPossible is a free data retrieval call binding the contract method 0x00caccb0.
//
// Solidity: function pickRandomStakerIfPossible(bytes32 entropy_) view returns(address)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) PickRandomStakerIfPossible(opts *bind.CallOpts, entropy_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "pickRandomStakerIfPossible", entropy_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PickRandomStakerIfPossible is a free data retrieval call binding the contract method 0x00caccb0.
//
// Solidity: function pickRandomStakerIfPossible(bytes32 entropy_) view returns(address)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) PickRandomStakerIfPossible(entropy_ [32]byte) (common.Address, error) {
	return _StakingWalletRandomWalkNft.Contract.PickRandomStakerIfPossible(&_StakingWalletRandomWalkNft.CallOpts, entropy_)
}

// PickRandomStakerIfPossible is a free data retrieval call binding the contract method 0x00caccb0.
//
// Solidity: function pickRandomStakerIfPossible(bytes32 entropy_) view returns(address)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) PickRandomStakerIfPossible(entropy_ [32]byte) (common.Address, error) {
	return _StakingWalletRandomWalkNft.Contract.PickRandomStakerIfPossible(&_StakingWalletRandomWalkNft.CallOpts, entropy_)
}

// RandomWalkNft is a free data retrieval call binding the contract method 0x755b4ef7.
//
// Solidity: function randomWalkNft() view returns(address)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) RandomWalkNft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "randomWalkNft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RandomWalkNft is a free data retrieval call binding the contract method 0x755b4ef7.
//
// Solidity: function randomWalkNft() view returns(address)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) RandomWalkNft() (common.Address, error) {
	return _StakingWalletRandomWalkNft.Contract.RandomWalkNft(&_StakingWalletRandomWalkNft.CallOpts)
}

// RandomWalkNft is a free data retrieval call binding the contract method 0x755b4ef7.
//
// Solidity: function randomWalkNft() view returns(address)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) RandomWalkNft() (common.Address, error) {
	return _StakingWalletRandomWalkNft.Contract.RandomWalkNft(&_StakingWalletRandomWalkNft.CallOpts)
}

// StakeActionIds is a free data retrieval call binding the contract method 0x60294405.
//
// Solidity: function stakeActionIds(uint256 ) view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) StakeActionIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "stakeActionIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeActionIds is a free data retrieval call binding the contract method 0x60294405.
//
// Solidity: function stakeActionIds(uint256 ) view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) StakeActionIds(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.StakeActionIds(&_StakingWalletRandomWalkNft.CallOpts, arg0)
}

// StakeActionIds is a free data retrieval call binding the contract method 0x60294405.
//
// Solidity: function stakeActionIds(uint256 ) view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) StakeActionIds(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.StakeActionIds(&_StakingWalletRandomWalkNft.CallOpts, arg0)
}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 index, uint256 nftId, address nftOwnerAddress)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) StakeActions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Index           *big.Int
	NftId           *big.Int
	NftOwnerAddress common.Address
}, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "stakeActions", arg0)

	outstruct := new(struct {
		Index           *big.Int
		NftId           *big.Int
		NftOwnerAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NftId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NftOwnerAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 index, uint256 nftId, address nftOwnerAddress)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) StakeActions(arg0 *big.Int) (struct {
	Index           *big.Int
	NftId           *big.Int
	NftOwnerAddress common.Address
}, error) {
	return _StakingWalletRandomWalkNft.Contract.StakeActions(&_StakingWalletRandomWalkNft.CallOpts, arg0)
}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 index, uint256 nftId, address nftOwnerAddress)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) StakeActions(arg0 *big.Int) (struct {
	Index           *big.Int
	NftId           *big.Int
	NftOwnerAddress common.Address
}, error) {
	return _StakingWalletRandomWalkNft.Contract.StakeActions(&_StakingWalletRandomWalkNft.CallOpts, arg0)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(bool)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) WasNftUsed(opts *bind.CallOpts, nftId_ *big.Int) (bool, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "wasNftUsed", nftId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(bool)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) WasNftUsed(nftId_ *big.Int) (bool, error) {
	return _StakingWalletRandomWalkNft.Contract.WasNftUsed(&_StakingWalletRandomWalkNft.CallOpts, nftId_)
}

// WasNftUsed is a free data retrieval call binding the contract method 0x57951c74.
//
// Solidity: function wasNftUsed(uint256 nftId_) view returns(bool)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) WasNftUsed(nftId_ *big.Int) (bool, error) {
	return _StakingWalletRandomWalkNft.Contract.WasNftUsed(&_StakingWalletRandomWalkNft.CallOpts, nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactor) Stake(opts *bind.TransactOpts, nftId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.contract.Transact(opts, "stake", nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.Stake(&_StakingWalletRandomWalkNft.TransactOpts, nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactorSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.Stake(&_StakingWalletRandomWalkNft.TransactOpts, nftId_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactor) StakeMany(opts *bind.TransactOpts, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.contract.Transact(opts, "stakeMany", nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.StakeMany(&_StakingWalletRandomWalkNft.TransactOpts, nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactorSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.StakeMany(&_StakingWalletRandomWalkNft.TransactOpts, nftIds_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactor) Unstake(opts *bind.TransactOpts, stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.contract.Transact(opts, "unstake", stakeActionId_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) Unstake(stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.Unstake(&_StakingWalletRandomWalkNft.TransactOpts, stakeActionId_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactorSession) Unstake(stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.Unstake(&_StakingWalletRandomWalkNft.TransactOpts, stakeActionId_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactor) UnstakeMany(opts *bind.TransactOpts, stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.contract.Transact(opts, "unstakeMany", stakeActionIds_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) UnstakeMany(stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.UnstakeMany(&_StakingWalletRandomWalkNft.TransactOpts, stakeActionIds_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactorSession) UnstakeMany(stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.UnstakeMany(&_StakingWalletRandomWalkNft.TransactOpts, stakeActionIds_)
}

// StakingWalletRandomWalkNftNftStakedIterator is returned from FilterNftStaked and is used to iterate over the raw logs and unpacked data for NftStaked events raised by the StakingWalletRandomWalkNft contract.
type StakingWalletRandomWalkNftNftStakedIterator struct {
	Event *StakingWalletRandomWalkNftNftStaked // Event containing the contract specifics and raw log

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
func (it *StakingWalletRandomWalkNftNftStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletRandomWalkNftNftStaked)
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
		it.Event = new(StakingWalletRandomWalkNftNftStaked)
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
func (it *StakingWalletRandomWalkNftNftStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletRandomWalkNftNftStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletRandomWalkNftNftStaked represents a NftStaked event raised by the StakingWalletRandomWalkNft contract.
type StakingWalletRandomWalkNftNftStaked struct {
	StakeActionId *big.Int
	NftTypeCode   uint8
	NftId         *big.Int
	StakerAddress common.Address
	NumStakedNfts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNftStaked is a free log retrieval operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftFilterer) FilterNftStaked(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*StakingWalletRandomWalkNftNftStakedIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletRandomWalkNft.contract.FilterLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRandomWalkNftNftStakedIterator{contract: _StakingWalletRandomWalkNft.contract, event: "NftStaked", logs: logs, sub: sub}, nil
}

// WatchNftStaked is a free log subscription operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftFilterer) WatchNftStaked(opts *bind.WatchOpts, sink chan<- *StakingWalletRandomWalkNftNftStaked, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletRandomWalkNft.contract.WatchLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletRandomWalkNftNftStaked)
				if err := _StakingWalletRandomWalkNft.contract.UnpackLog(event, "NftStaked", log); err != nil {
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

// ParseNftStaked is a log parse operation binding the contract event 0xcbd8e5368101e7829397d044213d371dac0f6727ea4bdc42d2f5a1f2e2f00829.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint8 nftTypeCode, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftFilterer) ParseNftStaked(log types.Log) (*StakingWalletRandomWalkNftNftStaked, error) {
	event := new(StakingWalletRandomWalkNftNftStaked)
	if err := _StakingWalletRandomWalkNft.contract.UnpackLog(event, "NftStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletRandomWalkNftNftUnstakedIterator is returned from FilterNftUnstaked and is used to iterate over the raw logs and unpacked data for NftUnstaked events raised by the StakingWalletRandomWalkNft contract.
type StakingWalletRandomWalkNftNftUnstakedIterator struct {
	Event *StakingWalletRandomWalkNftNftUnstaked // Event containing the contract specifics and raw log

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
func (it *StakingWalletRandomWalkNftNftUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletRandomWalkNftNftUnstaked)
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
		it.Event = new(StakingWalletRandomWalkNftNftUnstaked)
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
func (it *StakingWalletRandomWalkNftNftUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletRandomWalkNftNftUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletRandomWalkNftNftUnstaked represents a NftUnstaked event raised by the StakingWalletRandomWalkNft contract.
type StakingWalletRandomWalkNftNftUnstaked struct {
	StakeActionId *big.Int
	NftId         *big.Int
	StakerAddress common.Address
	NumStakedNfts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNftUnstaked is a free log retrieval operation binding the contract event 0x1792a7a9d5e2f53a8d06f2ae40b4446d91b503e84dd7f6307f40cdeb11541668.
//
// Solidity: event NftUnstaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftFilterer) FilterNftUnstaked(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*StakingWalletRandomWalkNftNftUnstakedIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletRandomWalkNft.contract.FilterLogs(opts, "NftUnstaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRandomWalkNftNftUnstakedIterator{contract: _StakingWalletRandomWalkNft.contract, event: "NftUnstaked", logs: logs, sub: sub}, nil
}

// WatchNftUnstaked is a free log subscription operation binding the contract event 0x1792a7a9d5e2f53a8d06f2ae40b4446d91b503e84dd7f6307f40cdeb11541668.
//
// Solidity: event NftUnstaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftFilterer) WatchNftUnstaked(opts *bind.WatchOpts, sink chan<- *StakingWalletRandomWalkNftNftUnstaked, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletRandomWalkNft.contract.WatchLogs(opts, "NftUnstaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletRandomWalkNftNftUnstaked)
				if err := _StakingWalletRandomWalkNft.contract.UnpackLog(event, "NftUnstaked", log); err != nil {
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

// ParseNftUnstaked is a log parse operation binding the contract event 0x1792a7a9d5e2f53a8d06f2ae40b4446d91b503e84dd7f6307f40cdeb11541668.
//
// Solidity: event NftUnstaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftFilterer) ParseNftUnstaked(log types.Log) (*StakingWalletRandomWalkNftNftUnstaked, error) {
	event := new(StakingWalletRandomWalkNftNftUnstaked)
	if err := _StakingWalletRandomWalkNft.contract.UnpackLog(event, "NftUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
