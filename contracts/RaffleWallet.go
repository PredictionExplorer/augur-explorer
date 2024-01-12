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

// RaffleWalletMetaData contains all meta data concerning the RaffleWallet contract.
var RaffleWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCosmicGame\",\"name\":\"game_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RaffleDepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RaffleWithdrawalEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"contractCosmicGame\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000fa338038062000fa38339818101604052810190620000379190620001e9565b620000576200004b6200009f60201b60201c565b620000a760201b60201c565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200021b565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200019d8262000170565b9050919050565b6000620001b18262000190565b9050919050565b620001c381620001a4565b8114620001cf57600080fd5b50565b600081519050620001e381620001b8565b92915050565b6000602082840312156200020257620002016200016b565b5b60006200021284828501620001d2565b91505092915050565b610d78806200022b6000396000f3fe6080604052600436106100705760003560e01c80638da5cb5b1161004e5780638da5cb5b146100e0578063c3fe3e281461010b578063f2fde38b14610136578063f340fa011461015f57610070565b806327e235e3146100755780633ccfd60b146100b2578063715018a6146100c9575b600080fd5b34801561008157600080fd5b5061009c60048036038101906100979190610843565b61017b565b6040516100a99190610889565b60405180910390f35b3480156100be57600080fd5b506100c7610193565b005b3480156100d557600080fd5b506100de61035d565b005b3480156100ec57600080fd5b506100f56103e5565b60405161010291906108b3565b60405180910390f35b34801561011757600080fd5b5061012061040e565b60405161012d919061092d565b60405180910390f35b34801561014257600080fd5b5061015d60048036038101906101589190610843565b610434565b005b61017960048036038101906101749190610843565b61052b565b005b60026020528060005260406000206000915090505481565b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000811161021a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610211906109a5565b60405180910390fd5b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060003373ffffffffffffffffffffffffffffffffffffffff1682604051610285906109f6565b60006040518083038185875af1925050503d80600081146102c2576040519150601f19603f3d011682016040523d82523d6000602084013e6102c7565b606091505b505090508061030b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030290610a57565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167f49775450da95afb818c4651b894f124d05629f06572a18900bf29db74b04a0d8836040516103519190610889565b60405180910390a25050565b610365610714565b73ffffffffffffffffffffffffffffffffffffffff166103836103e5565b73ffffffffffffffffffffffffffffffffffffffff16146103d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d090610ac3565b60405180910390fd5b6103e3600061071c565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61043c610714565b73ffffffffffffffffffffffffffffffffffffffff1661045a6103e5565b73ffffffffffffffffffffffffffffffffffffffff16146104b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a790610ac3565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361051f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161051690610b55565b60405180910390fd5b6105288161071c565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361059a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059190610bc1565b60405180910390fd5b600034116105dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105d490610c2d565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461066d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066490610cbf565b60405180910390fd5b34600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546106bc9190610d0e565b925050819055508073ffffffffffffffffffffffffffffffffffffffff167fcf6f6dcf9306290f700abd5d57c300a4ad7bf0d4086d5a5e88040fcecb0fef90346040516107099190610889565b60405180910390a250565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610810826107e5565b9050919050565b61082081610805565b811461082b57600080fd5b50565b60008135905061083d81610817565b92915050565b600060208284031215610859576108586107e0565b5b60006108678482850161082e565b91505092915050565b6000819050919050565b61088381610870565b82525050565b600060208201905061089e600083018461087a565b92915050565b6108ad81610805565b82525050565b60006020820190506108c860008301846108a4565b92915050565b6000819050919050565b60006108f36108ee6108e9846107e5565b6108ce565b6107e5565b9050919050565b6000610905826108d8565b9050919050565b6000610917826108fa565b9050919050565b6109278161090c565b82525050565b6000602082019050610942600083018461091e565b92915050565b600082825260208201905092915050565b7f596f75722062616c616e636520697320302e0000000000000000000000000000600082015250565b600061098f601283610948565b915061099a82610959565b602082019050919050565b600060208201905081810360008301526109be81610982565b9050919050565b600081905092915050565b50565b60006109e06000836109c5565b91506109eb826109d0565b600082019050919050565b6000610a01826109d3565b9150819050919050565b7f5472616e73666572206661696c65642e00000000000000000000000000000000600082015250565b6000610a41601083610948565b9150610a4c82610a0b565b602082019050919050565b60006020820190508181036000830152610a7081610a34565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000610aad602083610948565b9150610ab882610a77565b602082019050919050565b60006020820190508181036000830152610adc81610aa0565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000610b3f602683610948565b9150610b4a82610ae3565b604082019050919050565b60006020820190508181036000830152610b6e81610b32565b9050919050565b7f5a65726f2d616464726573732077617320676976656e2e000000000000000000600082015250565b6000610bab601783610948565b9150610bb682610b75565b602082019050919050565b60006020820190508181036000830152610bda81610b9e565b9050919050565b7f4e6f2045544820686173206265656e2073656e742e0000000000000000000000600082015250565b6000610c17601583610948565b9150610c2282610be1565b602082019050919050565b60006020820190508181036000830152610c4681610c0a565b9050919050565b7f4f6e6c7920436f736d696347616d6520697320616c6c6f77656420746f20646560008201527f706f7369742e0000000000000000000000000000000000000000000000000000602082015250565b6000610ca9602683610948565b9150610cb482610c4d565b604082019050919050565b60006020820190508181036000830152610cd881610c9c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610d1982610870565b9150610d2483610870565b9250828201905080821115610d3c57610d3b610cdf565b5b9291505056fea264697066735822122073fbfa1bd7ee4bc9446ec61ce08d554968744a5387f529c6e4800a9b57e414c264736f6c63430008130033",
}

// RaffleWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use RaffleWalletMetaData.ABI instead.
var RaffleWalletABI = RaffleWalletMetaData.ABI

// RaffleWalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RaffleWalletMetaData.Bin instead.
var RaffleWalletBin = RaffleWalletMetaData.Bin

// DeployRaffleWallet deploys a new Ethereum contract, binding an instance of RaffleWallet to it.
func DeployRaffleWallet(auth *bind.TransactOpts, backend bind.ContractBackend, game_ common.Address) (common.Address, *types.Transaction, *RaffleWallet, error) {
	parsed, err := RaffleWalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RaffleWalletBin), backend, game_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RaffleWallet{RaffleWalletCaller: RaffleWalletCaller{contract: contract}, RaffleWalletTransactor: RaffleWalletTransactor{contract: contract}, RaffleWalletFilterer: RaffleWalletFilterer{contract: contract}}, nil
}

// RaffleWallet is an auto generated Go binding around an Ethereum contract.
type RaffleWallet struct {
	RaffleWalletCaller     // Read-only binding to the contract
	RaffleWalletTransactor // Write-only binding to the contract
	RaffleWalletFilterer   // Log filterer for contract events
}

// RaffleWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type RaffleWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaffleWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RaffleWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaffleWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RaffleWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaffleWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RaffleWalletSession struct {
	Contract     *RaffleWallet     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RaffleWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RaffleWalletCallerSession struct {
	Contract *RaffleWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RaffleWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RaffleWalletTransactorSession struct {
	Contract     *RaffleWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RaffleWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type RaffleWalletRaw struct {
	Contract *RaffleWallet // Generic contract binding to access the raw methods on
}

// RaffleWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RaffleWalletCallerRaw struct {
	Contract *RaffleWalletCaller // Generic read-only contract binding to access the raw methods on
}

// RaffleWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RaffleWalletTransactorRaw struct {
	Contract *RaffleWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRaffleWallet creates a new instance of RaffleWallet, bound to a specific deployed contract.
func NewRaffleWallet(address common.Address, backend bind.ContractBackend) (*RaffleWallet, error) {
	contract, err := bindRaffleWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RaffleWallet{RaffleWalletCaller: RaffleWalletCaller{contract: contract}, RaffleWalletTransactor: RaffleWalletTransactor{contract: contract}, RaffleWalletFilterer: RaffleWalletFilterer{contract: contract}}, nil
}

// NewRaffleWalletCaller creates a new read-only instance of RaffleWallet, bound to a specific deployed contract.
func NewRaffleWalletCaller(address common.Address, caller bind.ContractCaller) (*RaffleWalletCaller, error) {
	contract, err := bindRaffleWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RaffleWalletCaller{contract: contract}, nil
}

// NewRaffleWalletTransactor creates a new write-only instance of RaffleWallet, bound to a specific deployed contract.
func NewRaffleWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*RaffleWalletTransactor, error) {
	contract, err := bindRaffleWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RaffleWalletTransactor{contract: contract}, nil
}

// NewRaffleWalletFilterer creates a new log filterer instance of RaffleWallet, bound to a specific deployed contract.
func NewRaffleWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*RaffleWalletFilterer, error) {
	contract, err := bindRaffleWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RaffleWalletFilterer{contract: contract}, nil
}

// bindRaffleWallet binds a generic wrapper to an already deployed contract.
func bindRaffleWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RaffleWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RaffleWallet *RaffleWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RaffleWallet.Contract.RaffleWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RaffleWallet *RaffleWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RaffleWallet.Contract.RaffleWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RaffleWallet *RaffleWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RaffleWallet.Contract.RaffleWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RaffleWallet *RaffleWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RaffleWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RaffleWallet *RaffleWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RaffleWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RaffleWallet *RaffleWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RaffleWallet.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_RaffleWallet *RaffleWalletCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RaffleWallet.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_RaffleWallet *RaffleWalletSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _RaffleWallet.Contract.Balances(&_RaffleWallet.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_RaffleWallet *RaffleWalletCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _RaffleWallet.Contract.Balances(&_RaffleWallet.CallOpts, arg0)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_RaffleWallet *RaffleWalletCaller) Game(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RaffleWallet.contract.Call(opts, &out, "game")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_RaffleWallet *RaffleWalletSession) Game() (common.Address, error) {
	return _RaffleWallet.Contract.Game(&_RaffleWallet.CallOpts)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_RaffleWallet *RaffleWalletCallerSession) Game() (common.Address, error) {
	return _RaffleWallet.Contract.Game(&_RaffleWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RaffleWallet *RaffleWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RaffleWallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RaffleWallet *RaffleWalletSession) Owner() (common.Address, error) {
	return _RaffleWallet.Contract.Owner(&_RaffleWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RaffleWallet *RaffleWalletCallerSession) Owner() (common.Address, error) {
	return _RaffleWallet.Contract.Owner(&_RaffleWallet.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address winner) payable returns()
func (_RaffleWallet *RaffleWalletTransactor) Deposit(opts *bind.TransactOpts, winner common.Address) (*types.Transaction, error) {
	return _RaffleWallet.contract.Transact(opts, "deposit", winner)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address winner) payable returns()
func (_RaffleWallet *RaffleWalletSession) Deposit(winner common.Address) (*types.Transaction, error) {
	return _RaffleWallet.Contract.Deposit(&_RaffleWallet.TransactOpts, winner)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address winner) payable returns()
func (_RaffleWallet *RaffleWalletTransactorSession) Deposit(winner common.Address) (*types.Transaction, error) {
	return _RaffleWallet.Contract.Deposit(&_RaffleWallet.TransactOpts, winner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RaffleWallet *RaffleWalletTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RaffleWallet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RaffleWallet *RaffleWalletSession) RenounceOwnership() (*types.Transaction, error) {
	return _RaffleWallet.Contract.RenounceOwnership(&_RaffleWallet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RaffleWallet *RaffleWalletTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RaffleWallet.Contract.RenounceOwnership(&_RaffleWallet.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RaffleWallet *RaffleWalletTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RaffleWallet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RaffleWallet *RaffleWalletSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RaffleWallet.Contract.TransferOwnership(&_RaffleWallet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RaffleWallet *RaffleWalletTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RaffleWallet.Contract.TransferOwnership(&_RaffleWallet.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RaffleWallet *RaffleWalletTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RaffleWallet.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RaffleWallet *RaffleWalletSession) Withdraw() (*types.Transaction, error) {
	return _RaffleWallet.Contract.Withdraw(&_RaffleWallet.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RaffleWallet *RaffleWalletTransactorSession) Withdraw() (*types.Transaction, error) {
	return _RaffleWallet.Contract.Withdraw(&_RaffleWallet.TransactOpts)
}

// RaffleWalletOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RaffleWallet contract.
type RaffleWalletOwnershipTransferredIterator struct {
	Event *RaffleWalletOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RaffleWalletOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleWalletOwnershipTransferred)
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
		it.Event = new(RaffleWalletOwnershipTransferred)
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
func (it *RaffleWalletOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleWalletOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleWalletOwnershipTransferred represents a OwnershipTransferred event raised by the RaffleWallet contract.
type RaffleWalletOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RaffleWallet *RaffleWalletFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RaffleWalletOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RaffleWallet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RaffleWalletOwnershipTransferredIterator{contract: _RaffleWallet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RaffleWallet *RaffleWalletFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RaffleWalletOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RaffleWallet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleWalletOwnershipTransferred)
				if err := _RaffleWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RaffleWallet *RaffleWalletFilterer) ParseOwnershipTransferred(log types.Log) (*RaffleWalletOwnershipTransferred, error) {
	event := new(RaffleWalletOwnershipTransferred)
	if err := _RaffleWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleWalletRaffleDepositEventIterator is returned from FilterRaffleDepositEvent and is used to iterate over the raw logs and unpacked data for RaffleDepositEvent events raised by the RaffleWallet contract.
type RaffleWalletRaffleDepositEventIterator struct {
	Event *RaffleWalletRaffleDepositEvent // Event containing the contract specifics and raw log

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
func (it *RaffleWalletRaffleDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleWalletRaffleDepositEvent)
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
		it.Event = new(RaffleWalletRaffleDepositEvent)
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
func (it *RaffleWalletRaffleDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleWalletRaffleDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleWalletRaffleDepositEvent represents a RaffleDepositEvent event raised by the RaffleWallet contract.
type RaffleWalletRaffleDepositEvent struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRaffleDepositEvent is a free log retrieval operation binding the contract event 0xcf6f6dcf9306290f700abd5d57c300a4ad7bf0d4086d5a5e88040fcecb0fef90.
//
// Solidity: event RaffleDepositEvent(address indexed winner, uint256 amount)
func (_RaffleWallet *RaffleWalletFilterer) FilterRaffleDepositEvent(opts *bind.FilterOpts, winner []common.Address) (*RaffleWalletRaffleDepositEventIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _RaffleWallet.contract.FilterLogs(opts, "RaffleDepositEvent", winnerRule)
	if err != nil {
		return nil, err
	}
	return &RaffleWalletRaffleDepositEventIterator{contract: _RaffleWallet.contract, event: "RaffleDepositEvent", logs: logs, sub: sub}, nil
}

// WatchRaffleDepositEvent is a free log subscription operation binding the contract event 0xcf6f6dcf9306290f700abd5d57c300a4ad7bf0d4086d5a5e88040fcecb0fef90.
//
// Solidity: event RaffleDepositEvent(address indexed winner, uint256 amount)
func (_RaffleWallet *RaffleWalletFilterer) WatchRaffleDepositEvent(opts *bind.WatchOpts, sink chan<- *RaffleWalletRaffleDepositEvent, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _RaffleWallet.contract.WatchLogs(opts, "RaffleDepositEvent", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleWalletRaffleDepositEvent)
				if err := _RaffleWallet.contract.UnpackLog(event, "RaffleDepositEvent", log); err != nil {
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

// ParseRaffleDepositEvent is a log parse operation binding the contract event 0xcf6f6dcf9306290f700abd5d57c300a4ad7bf0d4086d5a5e88040fcecb0fef90.
//
// Solidity: event RaffleDepositEvent(address indexed winner, uint256 amount)
func (_RaffleWallet *RaffleWalletFilterer) ParseRaffleDepositEvent(log types.Log) (*RaffleWalletRaffleDepositEvent, error) {
	event := new(RaffleWalletRaffleDepositEvent)
	if err := _RaffleWallet.contract.UnpackLog(event, "RaffleDepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaffleWalletRaffleWithdrawalEventIterator is returned from FilterRaffleWithdrawalEvent and is used to iterate over the raw logs and unpacked data for RaffleWithdrawalEvent events raised by the RaffleWallet contract.
type RaffleWalletRaffleWithdrawalEventIterator struct {
	Event *RaffleWalletRaffleWithdrawalEvent // Event containing the contract specifics and raw log

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
func (it *RaffleWalletRaffleWithdrawalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleWalletRaffleWithdrawalEvent)
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
		it.Event = new(RaffleWalletRaffleWithdrawalEvent)
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
func (it *RaffleWalletRaffleWithdrawalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleWalletRaffleWithdrawalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleWalletRaffleWithdrawalEvent represents a RaffleWithdrawalEvent event raised by the RaffleWallet contract.
type RaffleWalletRaffleWithdrawalEvent struct {
	Destination common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRaffleWithdrawalEvent is a free log retrieval operation binding the contract event 0x49775450da95afb818c4651b894f124d05629f06572a18900bf29db74b04a0d8.
//
// Solidity: event RaffleWithdrawalEvent(address indexed destination, uint256 amount)
func (_RaffleWallet *RaffleWalletFilterer) FilterRaffleWithdrawalEvent(opts *bind.FilterOpts, destination []common.Address) (*RaffleWalletRaffleWithdrawalEventIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _RaffleWallet.contract.FilterLogs(opts, "RaffleWithdrawalEvent", destinationRule)
	if err != nil {
		return nil, err
	}
	return &RaffleWalletRaffleWithdrawalEventIterator{contract: _RaffleWallet.contract, event: "RaffleWithdrawalEvent", logs: logs, sub: sub}, nil
}

// WatchRaffleWithdrawalEvent is a free log subscription operation binding the contract event 0x49775450da95afb818c4651b894f124d05629f06572a18900bf29db74b04a0d8.
//
// Solidity: event RaffleWithdrawalEvent(address indexed destination, uint256 amount)
func (_RaffleWallet *RaffleWalletFilterer) WatchRaffleWithdrawalEvent(opts *bind.WatchOpts, sink chan<- *RaffleWalletRaffleWithdrawalEvent, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _RaffleWallet.contract.WatchLogs(opts, "RaffleWithdrawalEvent", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleWalletRaffleWithdrawalEvent)
				if err := _RaffleWallet.contract.UnpackLog(event, "RaffleWithdrawalEvent", log); err != nil {
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

// ParseRaffleWithdrawalEvent is a log parse operation binding the contract event 0x49775450da95afb818c4651b894f124d05629f06572a18900bf29db74b04a0d8.
//
// Solidity: event RaffleWithdrawalEvent(address indexed destination, uint256 amount)
func (_RaffleWallet *RaffleWalletFilterer) ParseRaffleWithdrawalEvent(log types.Log) (*RaffleWalletRaffleWithdrawalEvent, error) {
	event := new(RaffleWalletRaffleWithdrawalEvent)
	if err := _RaffleWallet.contract.UnpackLog(event, "RaffleWithdrawalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
