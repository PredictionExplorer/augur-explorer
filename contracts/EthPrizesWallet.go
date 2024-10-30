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

// EthPrizesWalletMetaData contains all meta data concerning the EthPrizesWallet contract.
var EthPrizesWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"game_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"DepositFromUnauthorizedSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prizeAmount\",\"type\":\"uint256\"}],\"name\":\"PrizeReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prizeAmount\",\"type\":\"uint256\"}],\"name\":\"PrizeWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner_\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner_\",\"type\":\"address\"}],\"name\":\"getWinnerBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523461002f576100196100146100f4565b610246565b610021610034565b6107d261027882396107d290f35b61003a565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100669061003e565b810190811060018060401b0382111761007e57604052565b610048565b9061009661008f610034565b928361005c565b565b5f80fd5b60018060a01b031690565b6100b09061009c565b90565b6100bc816100a7565b036100c357565b5f80fd5b905051906100d4826100b3565b565b906020828203126100ef576100ec915f016100c7565b90565b610098565b610112610a4a8038038061010781610083565b9283398101906100d6565b90565b90565b90565b61012f61012a61013492610115565b610118565b61009c565b90565b6101409061011b565b90565b60209181520190565b5f7f5a65726f2d616464726573732077617320676976656e2e000000000000000000910152565b6101806017602092610143565b6101898161014c565b0190565b6101a29060208101905f818303910152610173565b90565b156101ac57565b6101b4610034565b63eac0d38960e01b8152806101cb6004820161018d565b0390fd5b5f1b90565b906101e560018060a01b03916101cf565b9181191691161790565b6102036101fe6102089261009c565b610118565b61009c565b90565b610214906101ef565b90565b6102209061020b565b90565b90565b9061023b61023661024292610217565b610223565b82546101d4565b9055565b6102759061026f8161026861026261025d5f610137565b6100a7565b916100a7565b14156101a5565b5f610226565b56fe60806040526004361015610013575b61022c565b61001d5f3561005c565b80633ccfd60b14610057578063c3fe3e2814610052578063dc8daa691461004d5763f340fa010361000e57610203565b6101ce565b610133565b610084565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f91031261007a57565b61006c565b5f0190565b346100b257610094366004610070565b61009c6104fc565b6100a4610062565b806100ae8161007f565b0390f35b610068565b1c90565b60018060a01b031690565b6100d69060086100db93026100b7565b6100bb565b90565b906100e991546100c6565b90565b6100f75f5f906100de565b90565b60018060a01b031690565b61010e906100fa565b90565b61011a90610105565b9052565b9190610131905f60208501940190610111565b565b3461016357610143366004610070565b61015f61014e6100ec565b610156610062565b9182918261011e565b0390f35b610068565b61017181610105565b0361017857565b5f80fd5b9050359061018982610168565b565b906020828203126101a4576101a1915f0161017c565b90565b61006c565b90565b6101b5906101a9565b9052565b91906101cc905f602085019401906101ac565b565b346101fe576101fa6101e96101e436600461018b565b6105ba565b6101f1610062565b918291826101b9565b0390f35b610068565b61021661021136600461018b565b6106fd565b61021e610062565b806102288161007f565b0390f35b5f80fd5b90565b61024761024261024c926100fa565b610230565b6100fa565b90565b61025890610233565b90565b634e487b7160e01b5f52603260045260245ffd5b50600160a01b90565b90565b6102848161026f565b82101561029e57610296600191610278565b910201905f90565b61025b565b90565b6102b69060086102bb93026100b7565b6102a3565b90565b906102c991546102a6565b90565b90565b6102e36102de6102e8926102cc565b610230565b6101a9565b90565b1b90565b9190600861030a9102916103045f19846102eb565b926102eb565b9181191691161790565b61032861032361032d926101a9565b610230565b6101a9565b90565b90565b919061034961034461035193610314565b610330565b9083546102ef565b9055565b61035e90610233565b90565b61036a90610355565b90565b905090565b61037d5f809261036d565b0190565b61038a90610372565b90565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906103b59061038d565b810190811067ffffffffffffffff8211176103cf57604052565b610397565b906103e76103e0610062565b92836103ab565b565b67ffffffffffffffff81116104075761040360209161038d565b0190565b610397565b9061041e610419836103e9565b6103d4565b918252565b606090565b3d5f14610443576104383d61040c565b903d5f602084013e5b565b61044b610423565b90610441565b60209181520190565b5f7f5072697a65207769746864726177616c206661696c65642e0000000000000000910152565b61048e6018602092610451565b6104978161045a565b0190565b9160406104cc9294936104c56104ba606083018381035f850152610481565b966020830190610111565b01906101ac565b565b156104d7575050565b6104f86104e2610062565b928392630aa7db6360e11b84526004840161049b565b0390fd5b6105b461051c61051660016105103361024f565b9061027b565b906102be565b6105426105285f6102cf565b61053c60016105363361024f565b9061027b565b90610333565b3381906105846105727f0835f27f462de8bce2ca086b3439451a8842d337fe4bf1fcc4aced1f952a2e2f92610361565b9261057b610062565b918291826101b9565b0390a25f803383610593610062565b908161059e81610381565b03925af16105aa610428565b50903390916104ce565b565b5f90565b6105db6105e1916105c96105b6565b506105d560019161024f565b9061027b565b906102be565b90565b5f1c90565b6105f56105fa916105e4565b6100bb565b90565b61060790546105e9565b90565b60207f7065726d697474656420746f206d616b652061206465706f7369742e00000000917f4f6e6c792074686520436f736d696347616d6520636f6e7472616374206973205f8201520152565b610664603c604092610451565b61066d8161060a565b0190565b919061069490602061068c604086018681035f880152610657565b940190610111565b565b1561069e5750565b6106c0906106aa610062565b918291637ed5977760e11b835260048301610671565b0390fd5b634e487b7160e01b5f52601160045260245ffd5b6106e76106ed919392936101a9565b926101a9565b82018092116106f857565b6106c4565b6107233361071b6107156107105f6105fd565b610105565b91610105565b143390610696565b6107563461075061073e60016107388661024f565b9061027b565b91909261074b83856102be565b6106d8565b91610333565b34906107976107857fee97096f32eed49908b904623ae7ba7af58c121e15fe6e7f31ac379fb7ca1a9892610361565b9261078e610062565b918291826101b9565b0390a256fea26469706673582212207a6288041a8765544b71cb004cc75d3e983d6fdedc816b9b8647a34f3dc6c45464736f6c634300081b0033",
}

// EthPrizesWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use EthPrizesWalletMetaData.ABI instead.
var EthPrizesWalletABI = EthPrizesWalletMetaData.ABI

// EthPrizesWalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthPrizesWalletMetaData.Bin instead.
var EthPrizesWalletBin = EthPrizesWalletMetaData.Bin

// DeployEthPrizesWallet deploys a new Ethereum contract, binding an instance of EthPrizesWallet to it.
func DeployEthPrizesWallet(auth *bind.TransactOpts, backend bind.ContractBackend, game_ common.Address) (common.Address, *types.Transaction, *EthPrizesWallet, error) {
	parsed, err := EthPrizesWalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthPrizesWalletBin), backend, game_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthPrizesWallet{EthPrizesWalletCaller: EthPrizesWalletCaller{contract: contract}, EthPrizesWalletTransactor: EthPrizesWalletTransactor{contract: contract}, EthPrizesWalletFilterer: EthPrizesWalletFilterer{contract: contract}}, nil
}

// EthPrizesWallet is an auto generated Go binding around an Ethereum contract.
type EthPrizesWallet struct {
	EthPrizesWalletCaller     // Read-only binding to the contract
	EthPrizesWalletTransactor // Write-only binding to the contract
	EthPrizesWalletFilterer   // Log filterer for contract events
}

// EthPrizesWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthPrizesWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthPrizesWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthPrizesWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthPrizesWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthPrizesWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthPrizesWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthPrizesWalletSession struct {
	Contract     *EthPrizesWallet  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthPrizesWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthPrizesWalletCallerSession struct {
	Contract *EthPrizesWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// EthPrizesWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthPrizesWalletTransactorSession struct {
	Contract     *EthPrizesWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EthPrizesWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthPrizesWalletRaw struct {
	Contract *EthPrizesWallet // Generic contract binding to access the raw methods on
}

// EthPrizesWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthPrizesWalletCallerRaw struct {
	Contract *EthPrizesWalletCaller // Generic read-only contract binding to access the raw methods on
}

// EthPrizesWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthPrizesWalletTransactorRaw struct {
	Contract *EthPrizesWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthPrizesWallet creates a new instance of EthPrizesWallet, bound to a specific deployed contract.
func NewEthPrizesWallet(address common.Address, backend bind.ContractBackend) (*EthPrizesWallet, error) {
	contract, err := bindEthPrizesWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthPrizesWallet{EthPrizesWalletCaller: EthPrizesWalletCaller{contract: contract}, EthPrizesWalletTransactor: EthPrizesWalletTransactor{contract: contract}, EthPrizesWalletFilterer: EthPrizesWalletFilterer{contract: contract}}, nil
}

// NewEthPrizesWalletCaller creates a new read-only instance of EthPrizesWallet, bound to a specific deployed contract.
func NewEthPrizesWalletCaller(address common.Address, caller bind.ContractCaller) (*EthPrizesWalletCaller, error) {
	contract, err := bindEthPrizesWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthPrizesWalletCaller{contract: contract}, nil
}

// NewEthPrizesWalletTransactor creates a new write-only instance of EthPrizesWallet, bound to a specific deployed contract.
func NewEthPrizesWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*EthPrizesWalletTransactor, error) {
	contract, err := bindEthPrizesWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthPrizesWalletTransactor{contract: contract}, nil
}

// NewEthPrizesWalletFilterer creates a new log filterer instance of EthPrizesWallet, bound to a specific deployed contract.
func NewEthPrizesWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*EthPrizesWalletFilterer, error) {
	contract, err := bindEthPrizesWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthPrizesWalletFilterer{contract: contract}, nil
}

// bindEthPrizesWallet binds a generic wrapper to an already deployed contract.
func bindEthPrizesWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthPrizesWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthPrizesWallet *EthPrizesWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthPrizesWallet.Contract.EthPrizesWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthPrizesWallet *EthPrizesWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthPrizesWallet.Contract.EthPrizesWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthPrizesWallet *EthPrizesWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthPrizesWallet.Contract.EthPrizesWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthPrizesWallet *EthPrizesWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthPrizesWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthPrizesWallet *EthPrizesWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthPrizesWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthPrizesWallet *EthPrizesWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthPrizesWallet.Contract.contract.Transact(opts, method, params...)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_EthPrizesWallet *EthPrizesWalletCaller) Game(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthPrizesWallet.contract.Call(opts, &out, "game")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_EthPrizesWallet *EthPrizesWalletSession) Game() (common.Address, error) {
	return _EthPrizesWallet.Contract.Game(&_EthPrizesWallet.CallOpts)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_EthPrizesWallet *EthPrizesWalletCallerSession) Game() (common.Address, error) {
	return _EthPrizesWallet.Contract.Game(&_EthPrizesWallet.CallOpts)
}

// GetWinnerBalance is a free data retrieval call binding the contract method 0xdc8daa69.
//
// Solidity: function getWinnerBalance(address winner_) view returns(uint256)
func (_EthPrizesWallet *EthPrizesWalletCaller) GetWinnerBalance(opts *bind.CallOpts, winner_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EthPrizesWallet.contract.Call(opts, &out, "getWinnerBalance", winner_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWinnerBalance is a free data retrieval call binding the contract method 0xdc8daa69.
//
// Solidity: function getWinnerBalance(address winner_) view returns(uint256)
func (_EthPrizesWallet *EthPrizesWalletSession) GetWinnerBalance(winner_ common.Address) (*big.Int, error) {
	return _EthPrizesWallet.Contract.GetWinnerBalance(&_EthPrizesWallet.CallOpts, winner_)
}

// GetWinnerBalance is a free data retrieval call binding the contract method 0xdc8daa69.
//
// Solidity: function getWinnerBalance(address winner_) view returns(uint256)
func (_EthPrizesWallet *EthPrizesWalletCallerSession) GetWinnerBalance(winner_ common.Address) (*big.Int, error) {
	return _EthPrizesWallet.Contract.GetWinnerBalance(&_EthPrizesWallet.CallOpts, winner_)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address winner_) payable returns()
func (_EthPrizesWallet *EthPrizesWalletTransactor) Deposit(opts *bind.TransactOpts, winner_ common.Address) (*types.Transaction, error) {
	return _EthPrizesWallet.contract.Transact(opts, "deposit", winner_)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address winner_) payable returns()
func (_EthPrizesWallet *EthPrizesWalletSession) Deposit(winner_ common.Address) (*types.Transaction, error) {
	return _EthPrizesWallet.Contract.Deposit(&_EthPrizesWallet.TransactOpts, winner_)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address winner_) payable returns()
func (_EthPrizesWallet *EthPrizesWalletTransactorSession) Deposit(winner_ common.Address) (*types.Transaction, error) {
	return _EthPrizesWallet.Contract.Deposit(&_EthPrizesWallet.TransactOpts, winner_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_EthPrizesWallet *EthPrizesWalletTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthPrizesWallet.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_EthPrizesWallet *EthPrizesWalletSession) Withdraw() (*types.Transaction, error) {
	return _EthPrizesWallet.Contract.Withdraw(&_EthPrizesWallet.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_EthPrizesWallet *EthPrizesWalletTransactorSession) Withdraw() (*types.Transaction, error) {
	return _EthPrizesWallet.Contract.Withdraw(&_EthPrizesWallet.TransactOpts)
}

// EthPrizesWalletPrizeReceivedIterator is returned from FilterPrizeReceived and is used to iterate over the raw logs and unpacked data for PrizeReceived events raised by the EthPrizesWallet contract.
type EthPrizesWalletPrizeReceivedIterator struct {
	Event *EthPrizesWalletPrizeReceived // Event containing the contract specifics and raw log

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
func (it *EthPrizesWalletPrizeReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthPrizesWalletPrizeReceived)
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
		it.Event = new(EthPrizesWalletPrizeReceived)
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
func (it *EthPrizesWalletPrizeReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthPrizesWalletPrizeReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthPrizesWalletPrizeReceived represents a PrizeReceived event raised by the EthPrizesWallet contract.
type EthPrizesWalletPrizeReceived struct {
	Winner      common.Address
	PrizeAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPrizeReceived is a free log retrieval operation binding the contract event 0xee97096f32eed49908b904623ae7ba7af58c121e15fe6e7f31ac379fb7ca1a98.
//
// Solidity: event PrizeReceived(address indexed winner, uint256 prizeAmount)
func (_EthPrizesWallet *EthPrizesWalletFilterer) FilterPrizeReceived(opts *bind.FilterOpts, winner []common.Address) (*EthPrizesWalletPrizeReceivedIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _EthPrizesWallet.contract.FilterLogs(opts, "PrizeReceived", winnerRule)
	if err != nil {
		return nil, err
	}
	return &EthPrizesWalletPrizeReceivedIterator{contract: _EthPrizesWallet.contract, event: "PrizeReceived", logs: logs, sub: sub}, nil
}

// WatchPrizeReceived is a free log subscription operation binding the contract event 0xee97096f32eed49908b904623ae7ba7af58c121e15fe6e7f31ac379fb7ca1a98.
//
// Solidity: event PrizeReceived(address indexed winner, uint256 prizeAmount)
func (_EthPrizesWallet *EthPrizesWalletFilterer) WatchPrizeReceived(opts *bind.WatchOpts, sink chan<- *EthPrizesWalletPrizeReceived, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _EthPrizesWallet.contract.WatchLogs(opts, "PrizeReceived", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthPrizesWalletPrizeReceived)
				if err := _EthPrizesWallet.contract.UnpackLog(event, "PrizeReceived", log); err != nil {
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

// ParsePrizeReceived is a log parse operation binding the contract event 0xee97096f32eed49908b904623ae7ba7af58c121e15fe6e7f31ac379fb7ca1a98.
//
// Solidity: event PrizeReceived(address indexed winner, uint256 prizeAmount)
func (_EthPrizesWallet *EthPrizesWalletFilterer) ParsePrizeReceived(log types.Log) (*EthPrizesWalletPrizeReceived, error) {
	event := new(EthPrizesWalletPrizeReceived)
	if err := _EthPrizesWallet.contract.UnpackLog(event, "PrizeReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthPrizesWalletPrizeWithdrawnIterator is returned from FilterPrizeWithdrawn and is used to iterate over the raw logs and unpacked data for PrizeWithdrawn events raised by the EthPrizesWallet contract.
type EthPrizesWalletPrizeWithdrawnIterator struct {
	Event *EthPrizesWalletPrizeWithdrawn // Event containing the contract specifics and raw log

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
func (it *EthPrizesWalletPrizeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthPrizesWalletPrizeWithdrawn)
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
		it.Event = new(EthPrizesWalletPrizeWithdrawn)
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
func (it *EthPrizesWalletPrizeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthPrizesWalletPrizeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthPrizesWalletPrizeWithdrawn represents a PrizeWithdrawn event raised by the EthPrizesWallet contract.
type EthPrizesWalletPrizeWithdrawn struct {
	Winner      common.Address
	PrizeAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPrizeWithdrawn is a free log retrieval operation binding the contract event 0x0835f27f462de8bce2ca086b3439451a8842d337fe4bf1fcc4aced1f952a2e2f.
//
// Solidity: event PrizeWithdrawn(address indexed winner, uint256 prizeAmount)
func (_EthPrizesWallet *EthPrizesWalletFilterer) FilterPrizeWithdrawn(opts *bind.FilterOpts, winner []common.Address) (*EthPrizesWalletPrizeWithdrawnIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _EthPrizesWallet.contract.FilterLogs(opts, "PrizeWithdrawn", winnerRule)
	if err != nil {
		return nil, err
	}
	return &EthPrizesWalletPrizeWithdrawnIterator{contract: _EthPrizesWallet.contract, event: "PrizeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchPrizeWithdrawn is a free log subscription operation binding the contract event 0x0835f27f462de8bce2ca086b3439451a8842d337fe4bf1fcc4aced1f952a2e2f.
//
// Solidity: event PrizeWithdrawn(address indexed winner, uint256 prizeAmount)
func (_EthPrizesWallet *EthPrizesWalletFilterer) WatchPrizeWithdrawn(opts *bind.WatchOpts, sink chan<- *EthPrizesWalletPrizeWithdrawn, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _EthPrizesWallet.contract.WatchLogs(opts, "PrizeWithdrawn", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthPrizesWalletPrizeWithdrawn)
				if err := _EthPrizesWallet.contract.UnpackLog(event, "PrizeWithdrawn", log); err != nil {
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

// ParsePrizeWithdrawn is a log parse operation binding the contract event 0x0835f27f462de8bce2ca086b3439451a8842d337fe4bf1fcc4aced1f952a2e2f.
//
// Solidity: event PrizeWithdrawn(address indexed winner, uint256 prizeAmount)
func (_EthPrizesWallet *EthPrizesWalletFilterer) ParsePrizeWithdrawn(log types.Log) (*EthPrizesWalletPrizeWithdrawn, error) {
	event := new(EthPrizesWalletPrizeWithdrawn)
	if err := _EthPrizesWallet.contract.UnpackLog(event, "PrizeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEthPrizesWalletMetaData contains all meta data concerning the IEthPrizesWallet contract.
var IEthPrizesWalletMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prizeAmount\",\"type\":\"uint256\"}],\"name\":\"PrizeReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prizeAmount\",\"type\":\"uint256\"}],\"name\":\"PrizeWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner_\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner_\",\"type\":\"address\"}],\"name\":\"getWinnerBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IEthPrizesWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use IEthPrizesWalletMetaData.ABI instead.
var IEthPrizesWalletABI = IEthPrizesWalletMetaData.ABI

// IEthPrizesWallet is an auto generated Go binding around an Ethereum contract.
type IEthPrizesWallet struct {
	IEthPrizesWalletCaller     // Read-only binding to the contract
	IEthPrizesWalletTransactor // Write-only binding to the contract
	IEthPrizesWalletFilterer   // Log filterer for contract events
}

// IEthPrizesWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthPrizesWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthPrizesWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthPrizesWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthPrizesWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthPrizesWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthPrizesWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthPrizesWalletSession struct {
	Contract     *IEthPrizesWallet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEthPrizesWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthPrizesWalletCallerSession struct {
	Contract *IEthPrizesWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IEthPrizesWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthPrizesWalletTransactorSession struct {
	Contract     *IEthPrizesWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IEthPrizesWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthPrizesWalletRaw struct {
	Contract *IEthPrizesWallet // Generic contract binding to access the raw methods on
}

// IEthPrizesWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthPrizesWalletCallerRaw struct {
	Contract *IEthPrizesWalletCaller // Generic read-only contract binding to access the raw methods on
}

// IEthPrizesWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthPrizesWalletTransactorRaw struct {
	Contract *IEthPrizesWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthPrizesWallet creates a new instance of IEthPrizesWallet, bound to a specific deployed contract.
func NewIEthPrizesWallet(address common.Address, backend bind.ContractBackend) (*IEthPrizesWallet, error) {
	contract, err := bindIEthPrizesWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthPrizesWallet{IEthPrizesWalletCaller: IEthPrizesWalletCaller{contract: contract}, IEthPrizesWalletTransactor: IEthPrizesWalletTransactor{contract: contract}, IEthPrizesWalletFilterer: IEthPrizesWalletFilterer{contract: contract}}, nil
}

// NewIEthPrizesWalletCaller creates a new read-only instance of IEthPrizesWallet, bound to a specific deployed contract.
func NewIEthPrizesWalletCaller(address common.Address, caller bind.ContractCaller) (*IEthPrizesWalletCaller, error) {
	contract, err := bindIEthPrizesWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthPrizesWalletCaller{contract: contract}, nil
}

// NewIEthPrizesWalletTransactor creates a new write-only instance of IEthPrizesWallet, bound to a specific deployed contract.
func NewIEthPrizesWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthPrizesWalletTransactor, error) {
	contract, err := bindIEthPrizesWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthPrizesWalletTransactor{contract: contract}, nil
}

// NewIEthPrizesWalletFilterer creates a new log filterer instance of IEthPrizesWallet, bound to a specific deployed contract.
func NewIEthPrizesWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthPrizesWalletFilterer, error) {
	contract, err := bindIEthPrizesWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthPrizesWalletFilterer{contract: contract}, nil
}

// bindIEthPrizesWallet binds a generic wrapper to an already deployed contract.
func bindIEthPrizesWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEthPrizesWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthPrizesWallet *IEthPrizesWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthPrizesWallet.Contract.IEthPrizesWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthPrizesWallet *IEthPrizesWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthPrizesWallet.Contract.IEthPrizesWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthPrizesWallet *IEthPrizesWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthPrizesWallet.Contract.IEthPrizesWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthPrizesWallet *IEthPrizesWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthPrizesWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthPrizesWallet *IEthPrizesWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthPrizesWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthPrizesWallet *IEthPrizesWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthPrizesWallet.Contract.contract.Transact(opts, method, params...)
}

// GetWinnerBalance is a free data retrieval call binding the contract method 0xdc8daa69.
//
// Solidity: function getWinnerBalance(address winner_) view returns(uint256)
func (_IEthPrizesWallet *IEthPrizesWalletCaller) GetWinnerBalance(opts *bind.CallOpts, winner_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IEthPrizesWallet.contract.Call(opts, &out, "getWinnerBalance", winner_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWinnerBalance is a free data retrieval call binding the contract method 0xdc8daa69.
//
// Solidity: function getWinnerBalance(address winner_) view returns(uint256)
func (_IEthPrizesWallet *IEthPrizesWalletSession) GetWinnerBalance(winner_ common.Address) (*big.Int, error) {
	return _IEthPrizesWallet.Contract.GetWinnerBalance(&_IEthPrizesWallet.CallOpts, winner_)
}

// GetWinnerBalance is a free data retrieval call binding the contract method 0xdc8daa69.
//
// Solidity: function getWinnerBalance(address winner_) view returns(uint256)
func (_IEthPrizesWallet *IEthPrizesWalletCallerSession) GetWinnerBalance(winner_ common.Address) (*big.Int, error) {
	return _IEthPrizesWallet.Contract.GetWinnerBalance(&_IEthPrizesWallet.CallOpts, winner_)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address winner_) payable returns()
func (_IEthPrizesWallet *IEthPrizesWalletTransactor) Deposit(opts *bind.TransactOpts, winner_ common.Address) (*types.Transaction, error) {
	return _IEthPrizesWallet.contract.Transact(opts, "deposit", winner_)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address winner_) payable returns()
func (_IEthPrizesWallet *IEthPrizesWalletSession) Deposit(winner_ common.Address) (*types.Transaction, error) {
	return _IEthPrizesWallet.Contract.Deposit(&_IEthPrizesWallet.TransactOpts, winner_)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address winner_) payable returns()
func (_IEthPrizesWallet *IEthPrizesWalletTransactorSession) Deposit(winner_ common.Address) (*types.Transaction, error) {
	return _IEthPrizesWallet.Contract.Deposit(&_IEthPrizesWallet.TransactOpts, winner_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IEthPrizesWallet *IEthPrizesWalletTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthPrizesWallet.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IEthPrizesWallet *IEthPrizesWalletSession) Withdraw() (*types.Transaction, error) {
	return _IEthPrizesWallet.Contract.Withdraw(&_IEthPrizesWallet.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IEthPrizesWallet *IEthPrizesWalletTransactorSession) Withdraw() (*types.Transaction, error) {
	return _IEthPrizesWallet.Contract.Withdraw(&_IEthPrizesWallet.TransactOpts)
}

// IEthPrizesWalletPrizeReceivedIterator is returned from FilterPrizeReceived and is used to iterate over the raw logs and unpacked data for PrizeReceived events raised by the IEthPrizesWallet contract.
type IEthPrizesWalletPrizeReceivedIterator struct {
	Event *IEthPrizesWalletPrizeReceived // Event containing the contract specifics and raw log

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
func (it *IEthPrizesWalletPrizeReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEthPrizesWalletPrizeReceived)
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
		it.Event = new(IEthPrizesWalletPrizeReceived)
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
func (it *IEthPrizesWalletPrizeReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEthPrizesWalletPrizeReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEthPrizesWalletPrizeReceived represents a PrizeReceived event raised by the IEthPrizesWallet contract.
type IEthPrizesWalletPrizeReceived struct {
	Winner      common.Address
	PrizeAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPrizeReceived is a free log retrieval operation binding the contract event 0xee97096f32eed49908b904623ae7ba7af58c121e15fe6e7f31ac379fb7ca1a98.
//
// Solidity: event PrizeReceived(address indexed winner, uint256 prizeAmount)
func (_IEthPrizesWallet *IEthPrizesWalletFilterer) FilterPrizeReceived(opts *bind.FilterOpts, winner []common.Address) (*IEthPrizesWalletPrizeReceivedIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _IEthPrizesWallet.contract.FilterLogs(opts, "PrizeReceived", winnerRule)
	if err != nil {
		return nil, err
	}
	return &IEthPrizesWalletPrizeReceivedIterator{contract: _IEthPrizesWallet.contract, event: "PrizeReceived", logs: logs, sub: sub}, nil
}

// WatchPrizeReceived is a free log subscription operation binding the contract event 0xee97096f32eed49908b904623ae7ba7af58c121e15fe6e7f31ac379fb7ca1a98.
//
// Solidity: event PrizeReceived(address indexed winner, uint256 prizeAmount)
func (_IEthPrizesWallet *IEthPrizesWalletFilterer) WatchPrizeReceived(opts *bind.WatchOpts, sink chan<- *IEthPrizesWalletPrizeReceived, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _IEthPrizesWallet.contract.WatchLogs(opts, "PrizeReceived", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEthPrizesWalletPrizeReceived)
				if err := _IEthPrizesWallet.contract.UnpackLog(event, "PrizeReceived", log); err != nil {
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

// ParsePrizeReceived is a log parse operation binding the contract event 0xee97096f32eed49908b904623ae7ba7af58c121e15fe6e7f31ac379fb7ca1a98.
//
// Solidity: event PrizeReceived(address indexed winner, uint256 prizeAmount)
func (_IEthPrizesWallet *IEthPrizesWalletFilterer) ParsePrizeReceived(log types.Log) (*IEthPrizesWalletPrizeReceived, error) {
	event := new(IEthPrizesWalletPrizeReceived)
	if err := _IEthPrizesWallet.contract.UnpackLog(event, "PrizeReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEthPrizesWalletPrizeWithdrawnIterator is returned from FilterPrizeWithdrawn and is used to iterate over the raw logs and unpacked data for PrizeWithdrawn events raised by the IEthPrizesWallet contract.
type IEthPrizesWalletPrizeWithdrawnIterator struct {
	Event *IEthPrizesWalletPrizeWithdrawn // Event containing the contract specifics and raw log

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
func (it *IEthPrizesWalletPrizeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEthPrizesWalletPrizeWithdrawn)
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
		it.Event = new(IEthPrizesWalletPrizeWithdrawn)
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
func (it *IEthPrizesWalletPrizeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEthPrizesWalletPrizeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEthPrizesWalletPrizeWithdrawn represents a PrizeWithdrawn event raised by the IEthPrizesWallet contract.
type IEthPrizesWalletPrizeWithdrawn struct {
	Winner      common.Address
	PrizeAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPrizeWithdrawn is a free log retrieval operation binding the contract event 0x0835f27f462de8bce2ca086b3439451a8842d337fe4bf1fcc4aced1f952a2e2f.
//
// Solidity: event PrizeWithdrawn(address indexed winner, uint256 prizeAmount)
func (_IEthPrizesWallet *IEthPrizesWalletFilterer) FilterPrizeWithdrawn(opts *bind.FilterOpts, winner []common.Address) (*IEthPrizesWalletPrizeWithdrawnIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _IEthPrizesWallet.contract.FilterLogs(opts, "PrizeWithdrawn", winnerRule)
	if err != nil {
		return nil, err
	}
	return &IEthPrizesWalletPrizeWithdrawnIterator{contract: _IEthPrizesWallet.contract, event: "PrizeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchPrizeWithdrawn is a free log subscription operation binding the contract event 0x0835f27f462de8bce2ca086b3439451a8842d337fe4bf1fcc4aced1f952a2e2f.
//
// Solidity: event PrizeWithdrawn(address indexed winner, uint256 prizeAmount)
func (_IEthPrizesWallet *IEthPrizesWalletFilterer) WatchPrizeWithdrawn(opts *bind.WatchOpts, sink chan<- *IEthPrizesWalletPrizeWithdrawn, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _IEthPrizesWallet.contract.WatchLogs(opts, "PrizeWithdrawn", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEthPrizesWalletPrizeWithdrawn)
				if err := _IEthPrizesWallet.contract.UnpackLog(event, "PrizeWithdrawn", log); err != nil {
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

// ParsePrizeWithdrawn is a log parse operation binding the contract event 0x0835f27f462de8bce2ca086b3439451a8842d337fe4bf1fcc4aced1f952a2e2f.
//
// Solidity: event PrizeWithdrawn(address indexed winner, uint256 prizeAmount)
func (_IEthPrizesWallet *IEthPrizesWalletFilterer) ParsePrizeWithdrawn(log types.Log) (*IEthPrizesWalletPrizeWithdrawn, error) {
	event := new(IEthPrizesWalletPrizeWithdrawn)
	if err := _IEthPrizesWallet.contract.UnpackLog(event, "PrizeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
