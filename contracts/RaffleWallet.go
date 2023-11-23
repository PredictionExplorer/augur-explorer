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

// RaffleWalletMetaData contains all meta data concerning the RaffleWallet contract.
var RaffleWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCosmicGame\",\"name\":\"game_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RaffleDepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RaffleWithdrawalEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"contractCosmicGame\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000fa338038062000fa38339818101604052810190620000379190620001e9565b620000576200004b6200009f60201b60201c565b620000a760201b60201c565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200021b565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200019d8262000170565b9050919050565b6000620001b18262000190565b9050919050565b620001c381620001a4565b8114620001cf57600080fd5b50565b600081519050620001e381620001b8565b92915050565b6000602082840312156200020257620002016200016b565b5b60006200021284828501620001d2565b91505092915050565b610d78806200022b6000396000f3fe6080604052600436106100705760003560e01c80638da5cb5b1161004e5780638da5cb5b146100e0578063c3fe3e281461010b578063f2fde38b14610136578063f340fa011461015f57610070565b806327e235e3146100755780633ccfd60b146100b2578063715018a6146100c9575b600080fd5b34801561008157600080fd5b5061009c60048036038101906100979190610843565b61017b565b6040516100a99190610889565b60405180910390f35b3480156100be57600080fd5b506100c7610193565b005b3480156100d557600080fd5b506100de61035d565b005b3480156100ec57600080fd5b506100f56103e5565b60405161010291906108b3565b60405180910390f35b34801561011757600080fd5b5061012061040e565b60405161012d919061092d565b60405180910390f35b34801561014257600080fd5b5061015d60048036038101906101589190610843565b610434565b005b61017960048036038101906101749190610843565b61052b565b005b60026020528060005260406000206000915090505481565b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000811161021a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610211906109a5565b60405180910390fd5b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060003373ffffffffffffffffffffffffffffffffffffffff1682604051610285906109f6565b60006040518083038185875af1925050503d80600081146102c2576040519150601f19603f3d011682016040523d82523d6000602084013e6102c7565b606091505b505090508061030b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030290610a57565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167f49775450da95afb818c4651b894f124d05629f06572a18900bf29db74b04a0d8836040516103519190610889565b60405180910390a25050565b610365610714565b73ffffffffffffffffffffffffffffffffffffffff166103836103e5565b73ffffffffffffffffffffffffffffffffffffffff16146103d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d090610ac3565b60405180910390fd5b6103e3600061071c565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61043c610714565b73ffffffffffffffffffffffffffffffffffffffff1661045a6103e5565b73ffffffffffffffffffffffffffffffffffffffff16146104b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a790610ac3565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361051f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161051690610b55565b60405180910390fd5b6105288161071c565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361059a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059190610bc1565b60405180910390fd5b600034116105dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105d490610c2d565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461066d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066490610cbf565b60405180910390fd5b34600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546106bc9190610d0e565b925050819055508073ffffffffffffffffffffffffffffffffffffffff167fcf6f6dcf9306290f700abd5d57c300a4ad7bf0d4086d5a5e88040fcecb0fef90346040516107099190610889565b60405180910390a250565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610810826107e5565b9050919050565b61082081610805565b811461082b57600080fd5b50565b60008135905061083d81610817565b92915050565b600060208284031215610859576108586107e0565b5b60006108678482850161082e565b91505092915050565b6000819050919050565b61088381610870565b82525050565b600060208201905061089e600083018461087a565b92915050565b6108ad81610805565b82525050565b60006020820190506108c860008301846108a4565b92915050565b6000819050919050565b60006108f36108ee6108e9846107e5565b6108ce565b6107e5565b9050919050565b6000610905826108d8565b9050919050565b6000610917826108fa565b9050919050565b6109278161090c565b82525050565b6000602082019050610942600083018461091e565b92915050565b600082825260208201905092915050565b7f596f75722062616c616e636520697320302e0000000000000000000000000000600082015250565b600061098f601283610948565b915061099a82610959565b602082019050919050565b600060208201905081810360008301526109be81610982565b9050919050565b600081905092915050565b50565b60006109e06000836109c5565b91506109eb826109d0565b600082019050919050565b6000610a01826109d3565b9150819050919050565b7f5472616e73666572206661696c65642e00000000000000000000000000000000600082015250565b6000610a41601083610948565b9150610a4c82610a0b565b602082019050919050565b60006020820190508181036000830152610a7081610a34565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000610aad602083610948565b9150610ab882610a77565b602082019050919050565b60006020820190508181036000830152610adc81610aa0565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000610b3f602683610948565b9150610b4a82610ae3565b604082019050919050565b60006020820190508181036000830152610b6e81610b32565b9050919050565b7f5a65726f2d616464726573732077617320676976656e2e000000000000000000600082015250565b6000610bab601783610948565b9150610bb682610b75565b602082019050919050565b60006020820190508181036000830152610bda81610b9e565b9050919050565b7f4e6f2045544820686173206265656e2073656e742e0000000000000000000000600082015250565b6000610c17601583610948565b9150610c2282610be1565b602082019050919050565b60006020820190508181036000830152610c4681610c0a565b9050919050565b7f4f6e6c7920436f736d696347616d6520697320616c6c6f77656420746f20646560008201527f706f7369742e0000000000000000000000000000000000000000000000000000602082015250565b6000610ca9602683610948565b9150610cb482610c4d565b604082019050919050565b60006020820190508181036000830152610cd881610c9c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610d1982610870565b9150610d2483610870565b9250828201905080821115610d3c57610d3b610cdf565b5b9291505056fea2646970667358221220c69d459fe58278ac929259a9f243fc1ddb2c5674ab0913125cb7b4c599a07bdc64736f6c63430008130033",
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
	parsed, err := abi.JSON(strings.NewReader(RaffleWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// RandomWalkNFTMetaData contains all meta data concerning the RandomWalkNFT contract.
var RandomWalkNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"seed\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"MintEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"TokenNameEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entropy\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMintTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMinter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numWithdrawals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"seeds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"seedsOfOwner\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setTokenName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeUntilSale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeUntilWithdrawal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenGenerationScript\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenNames\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"walletOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawalAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawalNums\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalWaitSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405263618dae80600b5566038d7ea4c68000600c5562278d00600d5560006010556000601460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600b5460155560006016556040518060600160405280603581526020016200560660359139601890816200009a9190620004db565b50348015620000a857600080fd5b506040518060400160405280600d81526020017f52616e646f6d57616c6b4e4654000000000000000000000000000000000000008152506040518060400160405280600481526020017f52574c4b000000000000000000000000000000000000000000000000000000008152508160009081620001269190620004db565b508060019081620001389190620004db565b5050506200015b6200014f6200019360201b60201c565b6200019b60201b60201c565b42434060405160200162000171929190620006c1565b6040516020818303038152906040528051906020012060138190555062000703565b600033905090565b6000600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620002e357607f821691505b602082108103620002f957620002f86200029b565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620003637fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000324565b6200036f868362000324565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620003bc620003b6620003b08462000387565b62000391565b62000387565b9050919050565b6000819050919050565b620003d8836200039b565b620003f0620003e782620003c3565b84845462000331565b825550505050565b600090565b62000407620003f8565b62000414818484620003cd565b505050565b5b818110156200043c5762000430600082620003fd565b6001810190506200041a565b5050565b601f8211156200048b576200045581620002ff565b620004608462000314565b8101602085101562000470578190505b620004886200047f8562000314565b83018262000419565b50505b505050565b600082821c905092915050565b6000620004b06000198460080262000490565b1980831691505092915050565b6000620004cb83836200049d565b9150826002028217905092915050565b620004e68262000261565b67ffffffffffffffff8111156200050257620005016200026c565b5b6200050e8254620002ca565b6200051b82828562000440565b600060209050601f8311600181146200055357600084156200053e578287015190505b6200054a8582620004bd565b865550620005ba565b601f1984166200056386620002ff565b60005b828110156200058d5784890151825560018201915060208501945060208101905062000566565b86831015620005ad5784890151620005a9601f8916826200049d565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b7f412074776f2d64696d656e73696f6e616c2072616e646f6d2077616c6b20776960008201527f6c6c2072657475726e20746f2074686520706f696e742077686572652069742060208201527f737461727465642c2062757420612074687265652d64696d656e73696f6e616c60408201527f206f6e65206d6179206e6f742e00000000000000000000000000000000000000606082015250565b60006200067d606d83620005c2565b91506200068a82620005d3565b608082019050919050565b620006a08162000387565b82525050565b6000819050919050565b620006bb81620006a6565b82525050565b60006060820190508181036000830152620006dc816200066e565b9050620006ed602083018562000695565b620006fc6040830184620006b0565b9392505050565b614ef380620007136000396000f3fe6080604052600436106102515760003560e01c8063661a1cfe11610139578063a7f93ebd116100b6578063e6e268f41161007a578063e6e268f4146108d1578063e985e9c5146108fc578063ee67d6a014610939578063f0503e8014610964578063f2fde38b146109a1578063f454aae1146109ca57610251565b8063a7f93ebd146107da578063b88d4fde14610805578063c87b56dd1461082e578063cb8efe951461086b578063cdb0e89e146108a857610251565b80638da5cb5b116100fd5780638da5cb5b1461070557806395d89b41146107305780639d4635201461075b578063a035b1fe14610786578063a22cb465146107b157610251565b8063661a1cfe1461061e5780636e56f6f91461064957806370a0823114610686578063715018a6146106c357806375794a3c146106da57610251565b80632f745c59116101d2578063438b630011610196578063438b6300146104e857806347ce07cc146105255780634cd609bc146105505780634f6ccce71461057b57806355f804b3146105b85780636352211e146105e157610251565b80632f745c5914610403578063310495ab146104405780633ccfd60b1461047d5780633e8af4da1461049457806342842e0e146104bf57610251565b8063157e394511610219578063157e39451461032e5780631596facb1461035957806317d209ab1461038457806318160ddd146103af57806323b872dd146103da57610251565b806301ffc9a71461025657806306fdde0314610293578063081812fc146102be578063095ea7b3146102fb5780631249c58b14610324575b600080fd5b34801561026257600080fd5b5061027d600480360381019061027891906131fd565b610a07565b60405161028a9190613245565b60405180910390f35b34801561029f57600080fd5b506102a8610a81565b6040516102b591906132f0565b60405180910390f35b3480156102ca57600080fd5b506102e560048036038101906102e09190613348565b610b13565b6040516102f291906133b6565b60405180910390f35b34801561030757600080fd5b50610322600480360381019061031d91906133fd565b610b98565b005b61032c610caf565b005b34801561033a57600080fd5b50610343610fbc565b604051610350919061344c565b60405180910390f35b34801561036557600080fd5b5061036e610fc2565b60405161037b919061344c565b60405180910390f35b34801561039057600080fd5b50610399610fc8565b6040516103a6919061344c565b60405180910390f35b3480156103bb57600080fd5b506103c4610fce565b6040516103d1919061344c565b60405180910390f35b3480156103e657600080fd5b5061040160048036038101906103fc9190613467565b610fdb565b005b34801561040f57600080fd5b5061042a600480360381019061042591906133fd565b61103b565b604051610437919061344c565b60405180910390f35b34801561044c57600080fd5b5061046760048036038101906104629190613348565b6110e0565b60405161047491906132f0565b60405180910390f35b34801561048957600080fd5b50610492611180565b005b3480156104a057600080fd5b506104a9611421565b6040516104b6919061344c565b60405180910390f35b3480156104cb57600080fd5b506104e660048036038101906104e19190613467565b61145b565b005b3480156104f457600080fd5b5061050f600480360381019061050a91906134ba565b61147b565b60405161051c91906135a5565b60405180910390f35b34801561053157600080fd5b5061053a611584565b60405161054791906135e0565b60405180910390f35b34801561055c57600080fd5b5061056561158a565b60405161057291906133b6565b60405180910390f35b34801561058757600080fd5b506105a2600480360381019061059d9190613348565b6115b0565b6040516105af919061344c565b60405180910390f35b3480156105c457600080fd5b506105df60048036038101906105da9190613730565b611621565b005b3480156105ed57600080fd5b5061060860048036038101906106039190613348565b6116b0565b60405161061591906133b6565b60405180910390f35b34801561062a57600080fd5b50610633611761565b604051610640919061344c565b60405180910390f35b34801561065557600080fd5b50610670600480360381019061066b9190613348565b61178a565b60405161067d919061344c565b60405180910390f35b34801561069257600080fd5b506106ad60048036038101906106a891906134ba565b6117a2565b6040516106ba919061344c565b60405180910390f35b3480156106cf57600080fd5b506106d8611859565b005b3480156106e657600080fd5b506106ef6118e1565b6040516106fc919061344c565b60405180910390f35b34801561071157600080fd5b5061071a6118e7565b60405161072791906133b6565b60405180910390f35b34801561073c57600080fd5b50610745611911565b60405161075291906132f0565b60405180910390f35b34801561076757600080fd5b506107706119a3565b60405161077d919061344c565b60405180910390f35b34801561079257600080fd5b5061079b6119a9565b6040516107a8919061344c565b60405180910390f35b3480156107bd57600080fd5b506107d860048036038101906107d391906137a5565b6119af565b005b3480156107e657600080fd5b506107ef611b2f565b6040516107fc919061344c565b60405180910390f35b34801561081157600080fd5b5061082c60048036038101906108279190613886565b611b53565b005b34801561083a57600080fd5b5061085560048036038101906108509190613348565b611bb5565b60405161086291906132f0565b60405180910390f35b34801561087757600080fd5b50610892600480360381019061088d91906134ba565b611c5c565b60405161089f91906139c7565b60405180910390f35b3480156108b457600080fd5b506108cf60048036038101906108ca91906139e9565b611d7e565b005b3480156108dd57600080fd5b506108e6611e71565b6040516108f3919061344c565b60405180910390f35b34801561090857600080fd5b50610923600480360381019061091e9190613a45565b611e85565b6040516109309190613245565b60405180910390f35b34801561094557600080fd5b5061094e611f19565b60405161095b91906132f0565b60405180910390f35b34801561097057600080fd5b5061098b60048036038101906109869190613348565b611fa7565b60405161099891906135e0565b60405180910390f35b3480156109ad57600080fd5b506109c860048036038101906109c391906134ba565b611fbf565b005b3480156109d657600080fd5b506109f160048036038101906109ec9190613348565b6120b6565b6040516109fe919061344c565b60405180910390f35b60007f780e9d63000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610a7a5750610a79826120ce565b5b9050919050565b606060008054610a9090613ab4565b80601f0160208091040260200160405190810160405280929190818152602001828054610abc90613ab4565b8015610b095780601f10610ade57610100808354040283529160200191610b09565b820191906000526020600020905b815481529060010190602001808311610aec57829003601f168201915b5050505050905090565b6000610b1e826121b0565b610b5d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b5490613b57565b60405180910390fd5b6004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6000610ba3826116b0565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610c13576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c0a90613be9565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16610c3261221c565b73ffffffffffffffffffffffffffffffffffffffff161480610c615750610c6081610c5b61221c565b611e85565b5b610ca0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c9790613c7b565b60405180910390fd5b610caa8383612224565b505050565b6000610cb9611b2f565b905080341015610cfe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf590613d0d565b60405180910390fd5b600b54421015610d43576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3a90613d79565b60405180910390fd5b610d4b61221c565b601460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055504260158190555080600c8190555060006016549050600160166000828254610db39190613dc8565b9250508190555060135442434083601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051602001610df8959493929190613dfc565b60405160208183030381529060405280519060200120601381905550601354600e600083815260200190815260200160002081905550610e5a601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16826122dd565b600c54341115610f41576000601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600c5434610ead9190613e4f565b604051610eb990613eb4565b60006040518083038185875af1925050503d8060008114610ef6576040519150601f19603f3d011682016040523d82523d6000602084013e610efb565b606091505b5050905080610f3f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f3690613f15565b60405180910390fd5b505b601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16817fad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec601354600c54604051610fb0929190613f35565b60405180910390a35050565b60105481565b600b5481565b600d5481565b6000600880549050905090565b610fec610fe661221c565b826122fb565b61102b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102290613fd0565b60405180910390fd5b6110368383836123d9565b505050565b6000611046836117a2565b8210611087576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161107e90614062565b60405180910390fd5b600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002054905092915050565b600f60205280600052604060002060009150905080546110ff90613ab4565b80601f016020809104026020016040519081016040528092919081815260200182805461112b90613ab4565b80156111785780601f1061114d57610100808354040283529160200191611178565b820191906000526020600020905b81548152906001019060200180831161115b57829003601f168201915b505050505081565b601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166111c161221c565b73ffffffffffffffffffffffffffffffffffffffff1614611217576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161120e906140ce565b60405180910390fd5b6000611221611421565b14611261576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112589061413a565b60405180910390fd5b6000601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000601460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600060016016546112db9190613e4f565b905060006112e7611e71565b90506001601060008282546112fc9190613dc8565b92505081905550601054601160008481526020019081526020016000208190555080601260008481526020019081526020016000208190555060008373ffffffffffffffffffffffffffffffffffffffff168260405161135b90613eb4565b60006040518083038185875af1925050503d8060008114611398576040519150601f19603f3d011682016040523d82523d6000602084013e61139d565b606091505b50509050806113e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113d890613f15565b60405180910390fd5b827fa11b556ace4b11a5cae8675a293b51e8cde3a06387d34010861789dfd9e9abc7858460405161141392919061415a565b60405180910390a250505050565b600080600d546015546114349190613dc8565b905042811015611448576000915050611458565b42816114549190613e4f565b9150505b90565b61147683838360405180602001604052806000815250611b53565b505050565b60606000611488836117a2565b9050600081036114e457600067ffffffffffffffff8111156114ad576114ac613605565b5b6040519080825280602002602001820160405280156114db5781602001602082028036833780820191505090505b5091505061157f565b60008167ffffffffffffffff811115611500576114ff613605565b5b60405190808252806020026020018201604052801561152e5781602001602082028036833780820191505090505b50905060005b8281101561157857611546858261103b565b82828151811061155957611558614183565b5b6020026020010181815250508080611570906141b2565b915050611534565b5080925050505b919050565b60135481565b601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006115ba610fce565b82106115fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115f29061426c565b60405180910390fd5b6008828154811061160f5761160e614183565b5b90600052602060002001549050919050565b61162961221c565b73ffffffffffffffffffffffffffffffffffffffff166116476118e7565b73ffffffffffffffffffffffffffffffffffffffff161461169d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611694906142d8565b60405180910390fd5b80601790816116ac91906144a4565b5050565b6000806002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611758576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161174f906145e8565b60405180910390fd5b80915050919050565b600042600b5410156117765760009050611787565b42600b546117849190613e4f565b90505b90565b60116020528060005260406000206000915090505481565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611812576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118099061467a565b60405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b61186161221c565b73ffffffffffffffffffffffffffffffffffffffff1661187f6118e7565b73ffffffffffffffffffffffffffffffffffffffff16146118d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118cc906142d8565b60405180910390fd5b6118df6000612634565b565b60165481565b6000600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606001805461192090613ab4565b80601f016020809104026020016040519081016040528092919081815260200182805461194c90613ab4565b80156119995780601f1061196e57610100808354040283529160200191611999565b820191906000526020600020905b81548152906001019060200180831161197c57829003601f168201915b5050505050905090565b60155481565b600c5481565b6119b761221c565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611a24576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a1b906146e6565b60405180910390fd5b8060056000611a3161221c565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff16611ade61221c565b73ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051611b239190613245565b60405180910390a35050565b600061271061271b600c54611b449190614706565b611b4e9190614777565b905090565b611b64611b5e61221c565b836122fb565b611ba3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b9a90613fd0565b60405180910390fd5b611baf848484846126fa565b50505050565b6060611bc0826121b0565b611bff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bf69061481a565b60405180910390fd5b6000611c09612756565b90506000815111611c295760405180602001604052806000815250611c54565b80611c33846127e8565b604051602001611c44929190614876565b6040516020818303038152906040525b915050919050565b60606000611c69836117a2565b905060008103611cc557600067ffffffffffffffff811115611c8e57611c8d613605565b5b604051908082528060200260200182016040528015611cbc5781602001602082028036833780820191505090505b50915050611d79565b60008167ffffffffffffffff811115611ce157611ce0613605565b5b604051908082528060200260200182016040528015611d0f5781602001602082028036833780820191505090505b50905060005b82811015611d72576000611d29868361103b565b9050600e600082815260200190815260200160002054838381518110611d5257611d51614183565b5b602002602001018181525050508080611d6a906141b2565b915050611d15565b5080925050505b919050565b611d8f611d8961221c565b836122fb565b611dce576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611dc59061490c565b60405180910390fd5b602081511115611e13576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e0a90614978565b60405180910390fd5b80600f60008481526020019081526020016000209081611e3391906144a4565b507f8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f128282604051611e65929190614998565b60405180910390a15050565b6000600247611e809190614777565b905090565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60188054611f2690613ab4565b80601f0160208091040260200160405190810160405280929190818152602001828054611f5290613ab4565b8015611f9f5780601f10611f7457610100808354040283529160200191611f9f565b820191906000526020600020905b815481529060010190602001808311611f8257829003601f168201915b505050505081565b600e6020528060005260406000206000915090505481565b611fc761221c565b73ffffffffffffffffffffffffffffffffffffffff16611fe56118e7565b73ffffffffffffffffffffffffffffffffffffffff161461203b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612032906142d8565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036120aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120a190614a3a565b60405180910390fd5b6120b381612634565b50565b60126020528060005260406000206000915090505481565b60007f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061219957507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806121a957506121a882612948565b5b9050919050565b60008073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b600033905090565b816004600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff16612297836116b0565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6122f78282604051806020016040528060008152506129b2565b5050565b6000612306826121b0565b612345576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161233c90614acc565b60405180910390fd5b6000612350836116b0565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614806123bf57508373ffffffffffffffffffffffffffffffffffffffff166123a784610b13565b73ffffffffffffffffffffffffffffffffffffffff16145b806123d057506123cf8185611e85565b5b91505092915050565b8273ffffffffffffffffffffffffffffffffffffffff166123f9826116b0565b73ffffffffffffffffffffffffffffffffffffffff161461244f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161244690614b5e565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036124be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124b590614bf0565b60405180910390fd5b6124c9838383612a0d565b6124d4600082612224565b6001600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546125249190613e4f565b925050819055506001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461257b9190613dc8565b92505081905550816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b6000600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6127058484846123d9565b61271184848484612b1f565b612750576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161274790614c82565b60405180910390fd5b50505050565b60606017805461276590613ab4565b80601f016020809104026020016040519081016040528092919081815260200182805461279190613ab4565b80156127de5780601f106127b3576101008083540402835291602001916127de565b820191906000526020600020905b8154815290600101906020018083116127c157829003601f168201915b5050505050905090565b60606000820361282f576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050612943565b600082905060005b6000821461286157808061284a906141b2565b915050600a8261285a9190614777565b9150612837565b60008167ffffffffffffffff81111561287d5761287c613605565b5b6040519080825280601f01601f1916602001820160405280156128af5781602001600182028036833780820191505090505b5090505b6000851461293c576001826128c89190613e4f565b9150600a856128d79190614ca2565b60306128e39190613dc8565b60f81b8183815181106128f9576128f8614183565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a856129359190614777565b94506128b3565b8093505050505b919050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b6129bc8383612ca6565b6129c96000848484612b1f565b612a08576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129ff90614c82565b60405180910390fd5b505050565b612a18838383612e73565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603612a5a57612a5581612e78565b612a99565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614612a9857612a978382612ec1565b5b5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603612adb57612ad68161302e565b612b1a565b8273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614612b1957612b1882826130ff565b5b5b505050565b6000612b408473ffffffffffffffffffffffffffffffffffffffff1661317e565b15612c99578373ffffffffffffffffffffffffffffffffffffffff1663150b7a02612b6961221c565b8786866040518563ffffffff1660e01b8152600401612b8b9493929190614d28565b6020604051808303816000875af1925050508015612bc757506040513d601f19601f82011682018060405250810190612bc49190614d89565b60015b612c49573d8060008114612bf7576040519150601f19603f3d011682016040523d82523d6000602084013e612bfc565b606091505b506000815103612c41576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c3890614c82565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614915050612c9e565b600190505b949350505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603612d15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d0c90614e02565b60405180910390fd5b612d1e816121b0565b15612d5e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d5590614e6e565b60405180910390fd5b612d6a60008383612a0d565b6001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254612dba9190613dc8565b92505081905550816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a45050565b505050565b6008805490506009600083815260200190815260200160002081905550600881908060018154018082558091505060019003906000526020600020016000909190919091505550565b60006001612ece846117a2565b612ed89190613e4f565b9050600060076000848152602001908152602001600020549050818114612fbd576000600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002054905080600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002081905550816007600083815260200190815260200160002081905550505b6007600084815260200190815260200160002060009055600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008381526020019081526020016000206000905550505050565b600060016008805490506130429190613e4f565b905060006009600084815260200190815260200160002054905060006008838154811061307257613071614183565b5b90600052602060002001549050806008838154811061309457613093614183565b5b9060005260206000200181905550816009600083815260200190815260200160002081905550600960008581526020019081526020016000206000905560088054806130e3576130e2614e8e565b5b6001900381819060005260206000200160009055905550505050565b600061310a836117a2565b905081600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002081905550806007600084815260200190815260200160002081905550505050565b600080823b905060008111915050919050565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6131da816131a5565b81146131e557600080fd5b50565b6000813590506131f7816131d1565b92915050565b6000602082840312156132135761321261319b565b5b6000613221848285016131e8565b91505092915050565b60008115159050919050565b61323f8161322a565b82525050565b600060208201905061325a6000830184613236565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561329a57808201518184015260208101905061327f565b60008484015250505050565b6000601f19601f8301169050919050565b60006132c282613260565b6132cc818561326b565b93506132dc81856020860161327c565b6132e5816132a6565b840191505092915050565b6000602082019050818103600083015261330a81846132b7565b905092915050565b6000819050919050565b61332581613312565b811461333057600080fd5b50565b6000813590506133428161331c565b92915050565b60006020828403121561335e5761335d61319b565b5b600061336c84828501613333565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006133a082613375565b9050919050565b6133b081613395565b82525050565b60006020820190506133cb60008301846133a7565b92915050565b6133da81613395565b81146133e557600080fd5b50565b6000813590506133f7816133d1565b92915050565b600080604083850312156134145761341361319b565b5b6000613422858286016133e8565b925050602061343385828601613333565b9150509250929050565b61344681613312565b82525050565b6000602082019050613461600083018461343d565b92915050565b6000806000606084860312156134805761347f61319b565b5b600061348e868287016133e8565b935050602061349f868287016133e8565b92505060406134b086828701613333565b9150509250925092565b6000602082840312156134d0576134cf61319b565b5b60006134de848285016133e8565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61351c81613312565b82525050565b600061352e8383613513565b60208301905092915050565b6000602082019050919050565b6000613552826134e7565b61355c81856134f2565b935061356783613503565b8060005b8381101561359857815161357f8882613522565b975061358a8361353a565b92505060018101905061356b565b5085935050505092915050565b600060208201905081810360008301526135bf8184613547565b905092915050565b6000819050919050565b6135da816135c7565b82525050565b60006020820190506135f560008301846135d1565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61363d826132a6565b810181811067ffffffffffffffff8211171561365c5761365b613605565b5b80604052505050565b600061366f613191565b905061367b8282613634565b919050565b600067ffffffffffffffff82111561369b5761369a613605565b5b6136a4826132a6565b9050602081019050919050565b82818337600083830152505050565b60006136d36136ce84613680565b613665565b9050828152602081018484840111156136ef576136ee613600565b5b6136fa8482856136b1565b509392505050565b600082601f830112613717576137166135fb565b5b81356137278482602086016136c0565b91505092915050565b6000602082840312156137465761374561319b565b5b600082013567ffffffffffffffff811115613764576137636131a0565b5b61377084828501613702565b91505092915050565b6137828161322a565b811461378d57600080fd5b50565b60008135905061379f81613779565b92915050565b600080604083850312156137bc576137bb61319b565b5b60006137ca858286016133e8565b92505060206137db85828601613790565b9150509250929050565b600067ffffffffffffffff821115613800576137ff613605565b5b613809826132a6565b9050602081019050919050565b6000613829613824846137e5565b613665565b90508281526020810184848401111561384557613844613600565b5b6138508482856136b1565b509392505050565b600082601f83011261386d5761386c6135fb565b5b813561387d848260208601613816565b91505092915050565b600080600080608085870312156138a05761389f61319b565b5b60006138ae878288016133e8565b94505060206138bf878288016133e8565b93505060406138d087828801613333565b925050606085013567ffffffffffffffff8111156138f1576138f06131a0565b5b6138fd87828801613858565b91505092959194509250565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61393e816135c7565b82525050565b60006139508383613935565b60208301905092915050565b6000602082019050919050565b600061397482613909565b61397e8185613914565b935061398983613925565b8060005b838110156139ba5781516139a18882613944565b97506139ac8361395c565b92505060018101905061398d565b5085935050505092915050565b600060208201905081810360008301526139e18184613969565b905092915050565b60008060408385031215613a00576139ff61319b565b5b6000613a0e85828601613333565b925050602083013567ffffffffffffffff811115613a2f57613a2e6131a0565b5b613a3b85828601613702565b9150509250929050565b60008060408385031215613a5c57613a5b61319b565b5b6000613a6a858286016133e8565b9250506020613a7b858286016133e8565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680613acc57607f821691505b602082108103613adf57613ade613a85565b5b50919050565b7f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860008201527f697374656e7420746f6b656e0000000000000000000000000000000000000000602082015250565b6000613b41602c8361326b565b9150613b4c82613ae5565b604082019050919050565b60006020820190508181036000830152613b7081613b34565b9050919050565b7f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560008201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b6000613bd360218361326b565b9150613bde82613b77565b604082019050919050565b60006020820190508181036000830152613c0281613bc6565b9050919050565b7f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f7760008201527f6e6572206e6f7220617070726f76656420666f7220616c6c0000000000000000602082015250565b6000613c6560388361326b565b9150613c7082613c09565b604082019050919050565b60006020820190508181036000830152613c9481613c58565b9050919050565b7f5468652076616c7565207375626d69747465642077697468207468697320747260008201527f616e73616374696f6e20697320746f6f206c6f772e0000000000000000000000602082015250565b6000613cf760358361326b565b9150613d0282613c9b565b604082019050919050565b60006020820190508181036000830152613d2681613cea565b9050919050565b7f5468652073616c65206973206e6f74206f70656e207965742e00000000000000600082015250565b6000613d6360198361326b565b9150613d6e82613d2d565b602082019050919050565b60006020820190508181036000830152613d9281613d56565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000613dd382613312565b9150613dde83613312565b9250828201905080821115613df657613df5613d99565b5b92915050565b600060a082019050613e1160008301886135d1565b613e1e602083018761343d565b613e2b60408301866135d1565b613e38606083018561343d565b613e4560808301846133a7565b9695505050505050565b6000613e5a82613312565b9150613e6583613312565b9250828203905081811115613e7d57613e7c613d99565b5b92915050565b600081905092915050565b50565b6000613e9e600083613e83565b9150613ea982613e8e565b600082019050919050565b6000613ebf82613e91565b9150819050919050565b7f5472616e73666572206661696c65642e00000000000000000000000000000000600082015250565b6000613eff60108361326b565b9150613f0a82613ec9565b602082019050919050565b60006020820190508181036000830152613f2e81613ef2565b9050919050565b6000604082019050613f4a60008301856135d1565b613f57602083018461343d565b9392505050565b7f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f60008201527f776e6572206e6f7220617070726f766564000000000000000000000000000000602082015250565b6000613fba60318361326b565b9150613fc582613f5e565b604082019050919050565b60006020820190508181036000830152613fe981613fad565b9050919050565b7f455243373231456e756d657261626c653a206f776e657220696e646578206f7560008201527f74206f6620626f756e6473000000000000000000000000000000000000000000602082015250565b600061404c602b8361326b565b915061405782613ff0565b604082019050919050565b6000602082019050818103600083015261407b8161403f565b9050919050565b7f4f6e6c79206c617374206d696e7465722063616e2077697468647261772e0000600082015250565b60006140b8601e8361326b565b91506140c382614082565b602082019050919050565b600060208201905081810360008301526140e7816140ab565b9050919050565b7f4e6f7420656e6f7567682074696d652068617320656c61707365642e00000000600082015250565b6000614124601c8361326b565b915061412f826140ee565b602082019050919050565b6000602082019050818103600083015261415381614117565b9050919050565b600060408201905061416f60008301856133a7565b61417c602083018461343d565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006141bd82613312565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036141ef576141ee613d99565b5b600182019050919050565b7f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60008201527f7574206f6620626f756e64730000000000000000000000000000000000000000602082015250565b6000614256602c8361326b565b9150614261826141fa565b604082019050919050565b6000602082019050818103600083015261428581614249565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006142c260208361326b565b91506142cd8261428c565b602082019050919050565b600060208201905081810360008301526142f1816142b5565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261435a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261431d565b614364868361431d565b95508019841693508086168417925050509392505050565b6000819050919050565b60006143a161439c61439784613312565b61437c565b613312565b9050919050565b6000819050919050565b6143bb83614386565b6143cf6143c7826143a8565b84845461432a565b825550505050565b600090565b6143e46143d7565b6143ef8184846143b2565b505050565b5b81811015614413576144086000826143dc565b6001810190506143f5565b5050565b601f82111561445857614429816142f8565b6144328461430d565b81016020851015614441578190505b61445561444d8561430d565b8301826143f4565b50505b505050565b600082821c905092915050565b600061447b6000198460080261445d565b1980831691505092915050565b6000614494838361446a565b9150826002028217905092915050565b6144ad82613260565b67ffffffffffffffff8111156144c6576144c5613605565b5b6144d08254613ab4565b6144db828285614417565b600060209050601f83116001811461450e57600084156144fc578287015190505b6145068582614488565b86555061456e565b601f19841661451c866142f8565b60005b828110156145445784890151825560018201915060208501945060208101905061451f565b86831015614561578489015161455d601f89168261446a565b8355505b6001600288020188555050505b505050505050565b7f4552433732313a206f776e657220717565727920666f72206e6f6e657869737460008201527f656e7420746f6b656e0000000000000000000000000000000000000000000000602082015250565b60006145d260298361326b565b91506145dd82614576565b604082019050919050565b60006020820190508181036000830152614601816145c5565b9050919050565b7f4552433732313a2062616c616e636520717565727920666f7220746865207a6560008201527f726f206164647265737300000000000000000000000000000000000000000000602082015250565b6000614664602a8361326b565b915061466f82614608565b604082019050919050565b6000602082019050818103600083015261469381614657565b9050919050565b7f4552433732313a20617070726f766520746f2063616c6c657200000000000000600082015250565b60006146d060198361326b565b91506146db8261469a565b602082019050919050565b600060208201905081810360008301526146ff816146c3565b9050919050565b600061471182613312565b915061471c83613312565b925082820261472a81613312565b9150828204841483151761474157614740613d99565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061478282613312565b915061478d83613312565b92508261479d5761479c614748565b5b828204905092915050565b7f4552433732314d657461646174613a2055524920717565727920666f72206e6f60008201527f6e6578697374656e7420746f6b656e0000000000000000000000000000000000602082015250565b6000614804602f8361326b565b915061480f826147a8565b604082019050919050565b60006020820190508181036000830152614833816147f7565b9050919050565b600081905092915050565b600061485082613260565b61485a818561483a565b935061486a81856020860161327c565b80840191505092915050565b60006148828285614845565b915061488e8284614845565b91508190509392505050565b7f736574546f6b656e4e616d652063616c6c6572206973206e6f74206f776e657260008201527f206e6f7220617070726f76656400000000000000000000000000000000000000602082015250565b60006148f6602d8361326b565b91506149018261489a565b604082019050919050565b60006020820190508181036000830152614925816148e9565b9050919050565b7f546f6b656e206e616d6520697320746f6f206c6f6e672e000000000000000000600082015250565b600061496260178361326b565b915061496d8261492c565b602082019050919050565b6000602082019050818103600083015261499181614955565b9050919050565b60006040820190506149ad600083018561343d565b81810360208301526149bf81846132b7565b90509392505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000614a2460268361326b565b9150614a2f826149c8565b604082019050919050565b60006020820190508181036000830152614a5381614a17565b9050919050565b7f4552433732313a206f70657261746f7220717565727920666f72206e6f6e657860008201527f697374656e7420746f6b656e0000000000000000000000000000000000000000602082015250565b6000614ab6602c8361326b565b9150614ac182614a5a565b604082019050919050565b60006020820190508181036000830152614ae581614aa9565b9050919050565b7f4552433732313a207472616e73666572206f6620746f6b656e2074686174206960008201527f73206e6f74206f776e0000000000000000000000000000000000000000000000602082015250565b6000614b4860298361326b565b9150614b5382614aec565b604082019050919050565b60006020820190508181036000830152614b7781614b3b565b9050919050565b7f4552433732313a207472616e7366657220746f20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000614bda60248361326b565b9150614be582614b7e565b604082019050919050565b60006020820190508181036000830152614c0981614bcd565b9050919050565b7f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560008201527f63656976657220696d706c656d656e7465720000000000000000000000000000602082015250565b6000614c6c60328361326b565b9150614c7782614c10565b604082019050919050565b60006020820190508181036000830152614c9b81614c5f565b9050919050565b6000614cad82613312565b9150614cb883613312565b925082614cc857614cc7614748565b5b828206905092915050565b600081519050919050565b600082825260208201905092915050565b6000614cfa82614cd3565b614d048185614cde565b9350614d1481856020860161327c565b614d1d816132a6565b840191505092915050565b6000608082019050614d3d60008301876133a7565b614d4a60208301866133a7565b614d57604083018561343d565b8181036060830152614d698184614cef565b905095945050505050565b600081519050614d83816131d1565b92915050565b600060208284031215614d9f57614d9e61319b565b5b6000614dad84828501614d74565b91505092915050565b7f4552433732313a206d696e7420746f20746865207a65726f2061646472657373600082015250565b6000614dec60208361326b565b9150614df782614db6565b602082019050919050565b60006020820190508181036000830152614e1b81614ddf565b9050919050565b7f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000600082015250565b6000614e58601c8361326b565b9150614e6382614e22565b602082019050919050565b60006020820190508181036000830152614e8781614e4b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea264697066735822122047852f6ed623c78af33027a5946e09618e4c4c3862daee4d6a23ca6c550b736264736f6c63430008130033697066733a2f2f516d50375a385662514c7079747a586e6365654141633444357458333958567a6f4565555a77454b3861506b3857",
}

// RandomWalkNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use RandomWalkNFTMetaData.ABI instead.
var RandomWalkNFTABI = RandomWalkNFTMetaData.ABI

// RandomWalkNFTBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RandomWalkNFTMetaData.Bin instead.
var RandomWalkNFTBin = RandomWalkNFTMetaData.Bin

// DeployRandomWalkNFT deploys a new Ethereum contract, binding an instance of RandomWalkNFT to it.
func DeployRandomWalkNFT(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RandomWalkNFT, error) {
	parsed, err := RandomWalkNFTMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RandomWalkNFTBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RandomWalkNFT{RandomWalkNFTCaller: RandomWalkNFTCaller{contract: contract}, RandomWalkNFTTransactor: RandomWalkNFTTransactor{contract: contract}, RandomWalkNFTFilterer: RandomWalkNFTFilterer{contract: contract}}, nil
}

// RandomWalkNFT is an auto generated Go binding around an Ethereum contract.
type RandomWalkNFT struct {
	RandomWalkNFTCaller     // Read-only binding to the contract
	RandomWalkNFTTransactor // Write-only binding to the contract
	RandomWalkNFTFilterer   // Log filterer for contract events
}

// RandomWalkNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type RandomWalkNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomWalkNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RandomWalkNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomWalkNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RandomWalkNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomWalkNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RandomWalkNFTSession struct {
	Contract     *RandomWalkNFT    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RandomWalkNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RandomWalkNFTCallerSession struct {
	Contract *RandomWalkNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RandomWalkNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RandomWalkNFTTransactorSession struct {
	Contract     *RandomWalkNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RandomWalkNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type RandomWalkNFTRaw struct {
	Contract *RandomWalkNFT // Generic contract binding to access the raw methods on
}

// RandomWalkNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RandomWalkNFTCallerRaw struct {
	Contract *RandomWalkNFTCaller // Generic read-only contract binding to access the raw methods on
}

// RandomWalkNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RandomWalkNFTTransactorRaw struct {
	Contract *RandomWalkNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRandomWalkNFT creates a new instance of RandomWalkNFT, bound to a specific deployed contract.
func NewRandomWalkNFT(address common.Address, backend bind.ContractBackend) (*RandomWalkNFT, error) {
	contract, err := bindRandomWalkNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFT{RandomWalkNFTCaller: RandomWalkNFTCaller{contract: contract}, RandomWalkNFTTransactor: RandomWalkNFTTransactor{contract: contract}, RandomWalkNFTFilterer: RandomWalkNFTFilterer{contract: contract}}, nil
}

// NewRandomWalkNFTCaller creates a new read-only instance of RandomWalkNFT, bound to a specific deployed contract.
func NewRandomWalkNFTCaller(address common.Address, caller bind.ContractCaller) (*RandomWalkNFTCaller, error) {
	contract, err := bindRandomWalkNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTCaller{contract: contract}, nil
}

// NewRandomWalkNFTTransactor creates a new write-only instance of RandomWalkNFT, bound to a specific deployed contract.
func NewRandomWalkNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*RandomWalkNFTTransactor, error) {
	contract, err := bindRandomWalkNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTTransactor{contract: contract}, nil
}

// NewRandomWalkNFTFilterer creates a new log filterer instance of RandomWalkNFT, bound to a specific deployed contract.
func NewRandomWalkNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*RandomWalkNFTFilterer, error) {
	contract, err := bindRandomWalkNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTFilterer{contract: contract}, nil
}

// bindRandomWalkNFT binds a generic wrapper to an already deployed contract.
func bindRandomWalkNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RandomWalkNFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomWalkNFT *RandomWalkNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomWalkNFT.Contract.RandomWalkNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomWalkNFT *RandomWalkNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.RandomWalkNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomWalkNFT *RandomWalkNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.RandomWalkNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomWalkNFT *RandomWalkNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomWalkNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomWalkNFT *RandomWalkNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomWalkNFT *RandomWalkNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _RandomWalkNFT.Contract.BalanceOf(&_RandomWalkNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _RandomWalkNFT.Contract.BalanceOf(&_RandomWalkNFT.CallOpts, owner)
}

// Entropy is a free data retrieval call binding the contract method 0x47ce07cc.
//
// Solidity: function entropy() view returns(bytes32)
func (_RandomWalkNFT *RandomWalkNFTCaller) Entropy(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "entropy")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Entropy is a free data retrieval call binding the contract method 0x47ce07cc.
//
// Solidity: function entropy() view returns(bytes32)
func (_RandomWalkNFT *RandomWalkNFTSession) Entropy() ([32]byte, error) {
	return _RandomWalkNFT.Contract.Entropy(&_RandomWalkNFT.CallOpts)
}

// Entropy is a free data retrieval call binding the contract method 0x47ce07cc.
//
// Solidity: function entropy() view returns(bytes32)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) Entropy() ([32]byte, error) {
	return _RandomWalkNFT.Contract.Entropy(&_RandomWalkNFT.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RandomWalkNFT *RandomWalkNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RandomWalkNFT *RandomWalkNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _RandomWalkNFT.Contract.GetApproved(&_RandomWalkNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _RandomWalkNFT.Contract.GetApproved(&_RandomWalkNFT.CallOpts, tokenId)
}

// GetMintPrice is a free data retrieval call binding the contract method 0xa7f93ebd.
//
// Solidity: function getMintPrice() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) GetMintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "getMintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMintPrice is a free data retrieval call binding the contract method 0xa7f93ebd.
//
// Solidity: function getMintPrice() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) GetMintPrice() (*big.Int, error) {
	return _RandomWalkNFT.Contract.GetMintPrice(&_RandomWalkNFT.CallOpts)
}

// GetMintPrice is a free data retrieval call binding the contract method 0xa7f93ebd.
//
// Solidity: function getMintPrice() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) GetMintPrice() (*big.Int, error) {
	return _RandomWalkNFT.Contract.GetMintPrice(&_RandomWalkNFT.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RandomWalkNFT *RandomWalkNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RandomWalkNFT *RandomWalkNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _RandomWalkNFT.Contract.IsApprovedForAll(&_RandomWalkNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _RandomWalkNFT.Contract.IsApprovedForAll(&_RandomWalkNFT.CallOpts, owner, operator)
}

// LastMintTime is a free data retrieval call binding the contract method 0x9d463520.
//
// Solidity: function lastMintTime() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) LastMintTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "lastMintTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastMintTime is a free data retrieval call binding the contract method 0x9d463520.
//
// Solidity: function lastMintTime() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) LastMintTime() (*big.Int, error) {
	return _RandomWalkNFT.Contract.LastMintTime(&_RandomWalkNFT.CallOpts)
}

// LastMintTime is a free data retrieval call binding the contract method 0x9d463520.
//
// Solidity: function lastMintTime() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) LastMintTime() (*big.Int, error) {
	return _RandomWalkNFT.Contract.LastMintTime(&_RandomWalkNFT.CallOpts)
}

// LastMinter is a free data retrieval call binding the contract method 0x4cd609bc.
//
// Solidity: function lastMinter() view returns(address)
func (_RandomWalkNFT *RandomWalkNFTCaller) LastMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "lastMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastMinter is a free data retrieval call binding the contract method 0x4cd609bc.
//
// Solidity: function lastMinter() view returns(address)
func (_RandomWalkNFT *RandomWalkNFTSession) LastMinter() (common.Address, error) {
	return _RandomWalkNFT.Contract.LastMinter(&_RandomWalkNFT.CallOpts)
}

// LastMinter is a free data retrieval call binding the contract method 0x4cd609bc.
//
// Solidity: function lastMinter() view returns(address)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) LastMinter() (common.Address, error) {
	return _RandomWalkNFT.Contract.LastMinter(&_RandomWalkNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RandomWalkNFT *RandomWalkNFTSession) Name() (string, error) {
	return _RandomWalkNFT.Contract.Name(&_RandomWalkNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) Name() (string, error) {
	return _RandomWalkNFT.Contract.Name(&_RandomWalkNFT.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) NextTokenId() (*big.Int, error) {
	return _RandomWalkNFT.Contract.NextTokenId(&_RandomWalkNFT.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) NextTokenId() (*big.Int, error) {
	return _RandomWalkNFT.Contract.NextTokenId(&_RandomWalkNFT.CallOpts)
}

// NumWithdrawals is a free data retrieval call binding the contract method 0x157e3945.
//
// Solidity: function numWithdrawals() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) NumWithdrawals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "numWithdrawals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumWithdrawals is a free data retrieval call binding the contract method 0x157e3945.
//
// Solidity: function numWithdrawals() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) NumWithdrawals() (*big.Int, error) {
	return _RandomWalkNFT.Contract.NumWithdrawals(&_RandomWalkNFT.CallOpts)
}

// NumWithdrawals is a free data retrieval call binding the contract method 0x157e3945.
//
// Solidity: function numWithdrawals() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) NumWithdrawals() (*big.Int, error) {
	return _RandomWalkNFT.Contract.NumWithdrawals(&_RandomWalkNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RandomWalkNFT *RandomWalkNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RandomWalkNFT *RandomWalkNFTSession) Owner() (common.Address, error) {
	return _RandomWalkNFT.Contract.Owner(&_RandomWalkNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) Owner() (common.Address, error) {
	return _RandomWalkNFT.Contract.Owner(&_RandomWalkNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RandomWalkNFT *RandomWalkNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RandomWalkNFT *RandomWalkNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _RandomWalkNFT.Contract.OwnerOf(&_RandomWalkNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _RandomWalkNFT.Contract.OwnerOf(&_RandomWalkNFT.CallOpts, tokenId)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) Price() (*big.Int, error) {
	return _RandomWalkNFT.Contract.Price(&_RandomWalkNFT.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) Price() (*big.Int, error) {
	return _RandomWalkNFT.Contract.Price(&_RandomWalkNFT.CallOpts)
}

// SaleTime is a free data retrieval call binding the contract method 0x1596facb.
//
// Solidity: function saleTime() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) SaleTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "saleTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SaleTime is a free data retrieval call binding the contract method 0x1596facb.
//
// Solidity: function saleTime() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) SaleTime() (*big.Int, error) {
	return _RandomWalkNFT.Contract.SaleTime(&_RandomWalkNFT.CallOpts)
}

// SaleTime is a free data retrieval call binding the contract method 0x1596facb.
//
// Solidity: function saleTime() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) SaleTime() (*big.Int, error) {
	return _RandomWalkNFT.Contract.SaleTime(&_RandomWalkNFT.CallOpts)
}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(bytes32)
func (_RandomWalkNFT *RandomWalkNFTCaller) Seeds(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "seeds", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(bytes32)
func (_RandomWalkNFT *RandomWalkNFTSession) Seeds(arg0 *big.Int) ([32]byte, error) {
	return _RandomWalkNFT.Contract.Seeds(&_RandomWalkNFT.CallOpts, arg0)
}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(bytes32)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) Seeds(arg0 *big.Int) ([32]byte, error) {
	return _RandomWalkNFT.Contract.Seeds(&_RandomWalkNFT.CallOpts, arg0)
}

// SeedsOfOwner is a free data retrieval call binding the contract method 0xcb8efe95.
//
// Solidity: function seedsOfOwner(address _owner) view returns(bytes32[])
func (_RandomWalkNFT *RandomWalkNFTCaller) SeedsOfOwner(opts *bind.CallOpts, _owner common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "seedsOfOwner", _owner)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// SeedsOfOwner is a free data retrieval call binding the contract method 0xcb8efe95.
//
// Solidity: function seedsOfOwner(address _owner) view returns(bytes32[])
func (_RandomWalkNFT *RandomWalkNFTSession) SeedsOfOwner(_owner common.Address) ([][32]byte, error) {
	return _RandomWalkNFT.Contract.SeedsOfOwner(&_RandomWalkNFT.CallOpts, _owner)
}

// SeedsOfOwner is a free data retrieval call binding the contract method 0xcb8efe95.
//
// Solidity: function seedsOfOwner(address _owner) view returns(bytes32[])
func (_RandomWalkNFT *RandomWalkNFTCallerSession) SeedsOfOwner(_owner common.Address) ([][32]byte, error) {
	return _RandomWalkNFT.Contract.SeedsOfOwner(&_RandomWalkNFT.CallOpts, _owner)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RandomWalkNFT *RandomWalkNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RandomWalkNFT *RandomWalkNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RandomWalkNFT.Contract.SupportsInterface(&_RandomWalkNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RandomWalkNFT.Contract.SupportsInterface(&_RandomWalkNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RandomWalkNFT *RandomWalkNFTSession) Symbol() (string, error) {
	return _RandomWalkNFT.Contract.Symbol(&_RandomWalkNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) Symbol() (string, error) {
	return _RandomWalkNFT.Contract.Symbol(&_RandomWalkNFT.CallOpts)
}

// TimeUntilSale is a free data retrieval call binding the contract method 0x661a1cfe.
//
// Solidity: function timeUntilSale() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) TimeUntilSale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "timeUntilSale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeUntilSale is a free data retrieval call binding the contract method 0x661a1cfe.
//
// Solidity: function timeUntilSale() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) TimeUntilSale() (*big.Int, error) {
	return _RandomWalkNFT.Contract.TimeUntilSale(&_RandomWalkNFT.CallOpts)
}

// TimeUntilSale is a free data retrieval call binding the contract method 0x661a1cfe.
//
// Solidity: function timeUntilSale() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) TimeUntilSale() (*big.Int, error) {
	return _RandomWalkNFT.Contract.TimeUntilSale(&_RandomWalkNFT.CallOpts)
}

// TimeUntilWithdrawal is a free data retrieval call binding the contract method 0x3e8af4da.
//
// Solidity: function timeUntilWithdrawal() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) TimeUntilWithdrawal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "timeUntilWithdrawal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeUntilWithdrawal is a free data retrieval call binding the contract method 0x3e8af4da.
//
// Solidity: function timeUntilWithdrawal() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) TimeUntilWithdrawal() (*big.Int, error) {
	return _RandomWalkNFT.Contract.TimeUntilWithdrawal(&_RandomWalkNFT.CallOpts)
}

// TimeUntilWithdrawal is a free data retrieval call binding the contract method 0x3e8af4da.
//
// Solidity: function timeUntilWithdrawal() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) TimeUntilWithdrawal() (*big.Int, error) {
	return _RandomWalkNFT.Contract.TimeUntilWithdrawal(&_RandomWalkNFT.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _RandomWalkNFT.Contract.TokenByIndex(&_RandomWalkNFT.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _RandomWalkNFT.Contract.TokenByIndex(&_RandomWalkNFT.CallOpts, index)
}

// TokenGenerationScript is a free data retrieval call binding the contract method 0xee67d6a0.
//
// Solidity: function tokenGenerationScript() view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCaller) TokenGenerationScript(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "tokenGenerationScript")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenGenerationScript is a free data retrieval call binding the contract method 0xee67d6a0.
//
// Solidity: function tokenGenerationScript() view returns(string)
func (_RandomWalkNFT *RandomWalkNFTSession) TokenGenerationScript() (string, error) {
	return _RandomWalkNFT.Contract.TokenGenerationScript(&_RandomWalkNFT.CallOpts)
}

// TokenGenerationScript is a free data retrieval call binding the contract method 0xee67d6a0.
//
// Solidity: function tokenGenerationScript() view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) TokenGenerationScript() (string, error) {
	return _RandomWalkNFT.Contract.TokenGenerationScript(&_RandomWalkNFT.CallOpts)
}

// TokenNames is a free data retrieval call binding the contract method 0x310495ab.
//
// Solidity: function tokenNames(uint256 ) view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCaller) TokenNames(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "tokenNames", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenNames is a free data retrieval call binding the contract method 0x310495ab.
//
// Solidity: function tokenNames(uint256 ) view returns(string)
func (_RandomWalkNFT *RandomWalkNFTSession) TokenNames(arg0 *big.Int) (string, error) {
	return _RandomWalkNFT.Contract.TokenNames(&_RandomWalkNFT.CallOpts, arg0)
}

// TokenNames is a free data retrieval call binding the contract method 0x310495ab.
//
// Solidity: function tokenNames(uint256 ) view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) TokenNames(arg0 *big.Int) (string, error) {
	return _RandomWalkNFT.Contract.TokenNames(&_RandomWalkNFT.CallOpts, arg0)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _RandomWalkNFT.Contract.TokenOfOwnerByIndex(&_RandomWalkNFT.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _RandomWalkNFT.Contract.TokenOfOwnerByIndex(&_RandomWalkNFT.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_RandomWalkNFT *RandomWalkNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _RandomWalkNFT.Contract.TokenURI(&_RandomWalkNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _RandomWalkNFT.Contract.TokenURI(&_RandomWalkNFT.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) TotalSupply() (*big.Int, error) {
	return _RandomWalkNFT.Contract.TotalSupply(&_RandomWalkNFT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) TotalSupply() (*big.Int, error) {
	return _RandomWalkNFT.Contract.TotalSupply(&_RandomWalkNFT.CallOpts)
}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256[])
func (_RandomWalkNFT *RandomWalkNFTCaller) WalletOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "walletOfOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256[])
func (_RandomWalkNFT *RandomWalkNFTSession) WalletOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _RandomWalkNFT.Contract.WalletOfOwner(&_RandomWalkNFT.CallOpts, _owner)
}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256[])
func (_RandomWalkNFT *RandomWalkNFTCallerSession) WalletOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _RandomWalkNFT.Contract.WalletOfOwner(&_RandomWalkNFT.CallOpts, _owner)
}

// WithdrawalAmount is a free data retrieval call binding the contract method 0xe6e268f4.
//
// Solidity: function withdrawalAmount() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) WithdrawalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "withdrawalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalAmount is a free data retrieval call binding the contract method 0xe6e268f4.
//
// Solidity: function withdrawalAmount() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) WithdrawalAmount() (*big.Int, error) {
	return _RandomWalkNFT.Contract.WithdrawalAmount(&_RandomWalkNFT.CallOpts)
}

// WithdrawalAmount is a free data retrieval call binding the contract method 0xe6e268f4.
//
// Solidity: function withdrawalAmount() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) WithdrawalAmount() (*big.Int, error) {
	return _RandomWalkNFT.Contract.WithdrawalAmount(&_RandomWalkNFT.CallOpts)
}

// WithdrawalAmounts is a free data retrieval call binding the contract method 0xf454aae1.
//
// Solidity: function withdrawalAmounts(uint256 ) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) WithdrawalAmounts(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "withdrawalAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalAmounts is a free data retrieval call binding the contract method 0xf454aae1.
//
// Solidity: function withdrawalAmounts(uint256 ) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) WithdrawalAmounts(arg0 *big.Int) (*big.Int, error) {
	return _RandomWalkNFT.Contract.WithdrawalAmounts(&_RandomWalkNFT.CallOpts, arg0)
}

// WithdrawalAmounts is a free data retrieval call binding the contract method 0xf454aae1.
//
// Solidity: function withdrawalAmounts(uint256 ) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) WithdrawalAmounts(arg0 *big.Int) (*big.Int, error) {
	return _RandomWalkNFT.Contract.WithdrawalAmounts(&_RandomWalkNFT.CallOpts, arg0)
}

// WithdrawalNums is a free data retrieval call binding the contract method 0x6e56f6f9.
//
// Solidity: function withdrawalNums(uint256 ) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) WithdrawalNums(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "withdrawalNums", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalNums is a free data retrieval call binding the contract method 0x6e56f6f9.
//
// Solidity: function withdrawalNums(uint256 ) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) WithdrawalNums(arg0 *big.Int) (*big.Int, error) {
	return _RandomWalkNFT.Contract.WithdrawalNums(&_RandomWalkNFT.CallOpts, arg0)
}

// WithdrawalNums is a free data retrieval call binding the contract method 0x6e56f6f9.
//
// Solidity: function withdrawalNums(uint256 ) view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) WithdrawalNums(arg0 *big.Int) (*big.Int, error) {
	return _RandomWalkNFT.Contract.WithdrawalNums(&_RandomWalkNFT.CallOpts, arg0)
}

// WithdrawalWaitSeconds is a free data retrieval call binding the contract method 0x17d209ab.
//
// Solidity: function withdrawalWaitSeconds() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCaller) WithdrawalWaitSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomWalkNFT.contract.Call(opts, &out, "withdrawalWaitSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalWaitSeconds is a free data retrieval call binding the contract method 0x17d209ab.
//
// Solidity: function withdrawalWaitSeconds() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTSession) WithdrawalWaitSeconds() (*big.Int, error) {
	return _RandomWalkNFT.Contract.WithdrawalWaitSeconds(&_RandomWalkNFT.CallOpts)
}

// WithdrawalWaitSeconds is a free data retrieval call binding the contract method 0x17d209ab.
//
// Solidity: function withdrawalWaitSeconds() view returns(uint256)
func (_RandomWalkNFT *RandomWalkNFTCallerSession) WithdrawalWaitSeconds() (*big.Int, error) {
	return _RandomWalkNFT.Contract.WithdrawalWaitSeconds(&_RandomWalkNFT.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_RandomWalkNFT *RandomWalkNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.Approve(&_RandomWalkNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.Approve(&_RandomWalkNFT.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_RandomWalkNFT *RandomWalkNFTSession) Mint() (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.Mint(&_RandomWalkNFT.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) Mint() (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.Mint(&_RandomWalkNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RandomWalkNFT *RandomWalkNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.RenounceOwnership(&_RandomWalkNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.RenounceOwnership(&_RandomWalkNFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RandomWalkNFT *RandomWalkNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SafeTransferFrom(&_RandomWalkNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SafeTransferFrom(&_RandomWalkNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_RandomWalkNFT *RandomWalkNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SafeTransferFrom0(&_RandomWalkNFT.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SafeTransferFrom0(&_RandomWalkNFT.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RandomWalkNFT *RandomWalkNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SetApprovalForAll(&_RandomWalkNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SetApprovalForAll(&_RandomWalkNFT.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI string) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "setBaseURI", baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_RandomWalkNFT *RandomWalkNFTSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SetBaseURI(&_RandomWalkNFT.TransactOpts, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SetBaseURI(&_RandomWalkNFT.TransactOpts, baseURI)
}

// SetTokenName is a paid mutator transaction binding the contract method 0xcdb0e89e.
//
// Solidity: function setTokenName(uint256 tokenId, string name) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) SetTokenName(opts *bind.TransactOpts, tokenId *big.Int, name string) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "setTokenName", tokenId, name)
}

// SetTokenName is a paid mutator transaction binding the contract method 0xcdb0e89e.
//
// Solidity: function setTokenName(uint256 tokenId, string name) returns()
func (_RandomWalkNFT *RandomWalkNFTSession) SetTokenName(tokenId *big.Int, name string) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SetTokenName(&_RandomWalkNFT.TransactOpts, tokenId, name)
}

// SetTokenName is a paid mutator transaction binding the contract method 0xcdb0e89e.
//
// Solidity: function setTokenName(uint256 tokenId, string name) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) SetTokenName(tokenId *big.Int, name string) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.SetTokenName(&_RandomWalkNFT.TransactOpts, tokenId, name)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_RandomWalkNFT *RandomWalkNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.TransferFrom(&_RandomWalkNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.TransferFrom(&_RandomWalkNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RandomWalkNFT *RandomWalkNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.TransferOwnership(&_RandomWalkNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.TransferOwnership(&_RandomWalkNFT.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RandomWalkNFT *RandomWalkNFTTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomWalkNFT.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RandomWalkNFT *RandomWalkNFTSession) Withdraw() (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.Withdraw(&_RandomWalkNFT.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_RandomWalkNFT *RandomWalkNFTTransactorSession) Withdraw() (*types.Transaction, error) {
	return _RandomWalkNFT.Contract.Withdraw(&_RandomWalkNFT.TransactOpts)
}

// RandomWalkNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RandomWalkNFT contract.
type RandomWalkNFTApprovalIterator struct {
	Event *RandomWalkNFTApproval // Event containing the contract specifics and raw log

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
func (it *RandomWalkNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomWalkNFTApproval)
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
		it.Event = new(RandomWalkNFTApproval)
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
func (it *RandomWalkNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomWalkNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomWalkNFTApproval represents a Approval event raised by the RandomWalkNFT contract.
type RandomWalkNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_RandomWalkNFT *RandomWalkNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*RandomWalkNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTApprovalIterator{contract: _RandomWalkNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_RandomWalkNFT *RandomWalkNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RandomWalkNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomWalkNFTApproval)
				if err := _RandomWalkNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_RandomWalkNFT *RandomWalkNFTFilterer) ParseApproval(log types.Log) (*RandomWalkNFTApproval, error) {
	event := new(RandomWalkNFTApproval)
	if err := _RandomWalkNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomWalkNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the RandomWalkNFT contract.
type RandomWalkNFTApprovalForAllIterator struct {
	Event *RandomWalkNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *RandomWalkNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomWalkNFTApprovalForAll)
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
		it.Event = new(RandomWalkNFTApprovalForAll)
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
func (it *RandomWalkNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomWalkNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomWalkNFTApprovalForAll represents a ApprovalForAll event raised by the RandomWalkNFT contract.
type RandomWalkNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_RandomWalkNFT *RandomWalkNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*RandomWalkNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTApprovalForAllIterator{contract: _RandomWalkNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_RandomWalkNFT *RandomWalkNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *RandomWalkNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomWalkNFTApprovalForAll)
				if err := _RandomWalkNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_RandomWalkNFT *RandomWalkNFTFilterer) ParseApprovalForAll(log types.Log) (*RandomWalkNFTApprovalForAll, error) {
	event := new(RandomWalkNFTApprovalForAll)
	if err := _RandomWalkNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomWalkNFTMintEventIterator is returned from FilterMintEvent and is used to iterate over the raw logs and unpacked data for MintEvent events raised by the RandomWalkNFT contract.
type RandomWalkNFTMintEventIterator struct {
	Event *RandomWalkNFTMintEvent // Event containing the contract specifics and raw log

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
func (it *RandomWalkNFTMintEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomWalkNFTMintEvent)
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
		it.Event = new(RandomWalkNFTMintEvent)
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
func (it *RandomWalkNFTMintEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomWalkNFTMintEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomWalkNFTMintEvent represents a MintEvent event raised by the RandomWalkNFT contract.
type RandomWalkNFTMintEvent struct {
	TokenId *big.Int
	Owner   common.Address
	Seed    [32]byte
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMintEvent is a free log retrieval operation binding the contract event 0xad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec.
//
// Solidity: event MintEvent(uint256 indexed tokenId, address indexed owner, bytes32 seed, uint256 price)
func (_RandomWalkNFT *RandomWalkNFTFilterer) FilterMintEvent(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address) (*RandomWalkNFTMintEventIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.FilterLogs(opts, "MintEvent", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTMintEventIterator{contract: _RandomWalkNFT.contract, event: "MintEvent", logs: logs, sub: sub}, nil
}

// WatchMintEvent is a free log subscription operation binding the contract event 0xad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec.
//
// Solidity: event MintEvent(uint256 indexed tokenId, address indexed owner, bytes32 seed, uint256 price)
func (_RandomWalkNFT *RandomWalkNFTFilterer) WatchMintEvent(opts *bind.WatchOpts, sink chan<- *RandomWalkNFTMintEvent, tokenId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.WatchLogs(opts, "MintEvent", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomWalkNFTMintEvent)
				if err := _RandomWalkNFT.contract.UnpackLog(event, "MintEvent", log); err != nil {
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

// ParseMintEvent is a log parse operation binding the contract event 0xad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec.
//
// Solidity: event MintEvent(uint256 indexed tokenId, address indexed owner, bytes32 seed, uint256 price)
func (_RandomWalkNFT *RandomWalkNFTFilterer) ParseMintEvent(log types.Log) (*RandomWalkNFTMintEvent, error) {
	event := new(RandomWalkNFTMintEvent)
	if err := _RandomWalkNFT.contract.UnpackLog(event, "MintEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomWalkNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RandomWalkNFT contract.
type RandomWalkNFTOwnershipTransferredIterator struct {
	Event *RandomWalkNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RandomWalkNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomWalkNFTOwnershipTransferred)
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
		it.Event = new(RandomWalkNFTOwnershipTransferred)
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
func (it *RandomWalkNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomWalkNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomWalkNFTOwnershipTransferred represents a OwnershipTransferred event raised by the RandomWalkNFT contract.
type RandomWalkNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RandomWalkNFT *RandomWalkNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RandomWalkNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTOwnershipTransferredIterator{contract: _RandomWalkNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RandomWalkNFT *RandomWalkNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RandomWalkNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomWalkNFTOwnershipTransferred)
				if err := _RandomWalkNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RandomWalkNFT *RandomWalkNFTFilterer) ParseOwnershipTransferred(log types.Log) (*RandomWalkNFTOwnershipTransferred, error) {
	event := new(RandomWalkNFTOwnershipTransferred)
	if err := _RandomWalkNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomWalkNFTTokenNameEventIterator is returned from FilterTokenNameEvent and is used to iterate over the raw logs and unpacked data for TokenNameEvent events raised by the RandomWalkNFT contract.
type RandomWalkNFTTokenNameEventIterator struct {
	Event *RandomWalkNFTTokenNameEvent // Event containing the contract specifics and raw log

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
func (it *RandomWalkNFTTokenNameEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomWalkNFTTokenNameEvent)
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
		it.Event = new(RandomWalkNFTTokenNameEvent)
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
func (it *RandomWalkNFTTokenNameEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomWalkNFTTokenNameEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomWalkNFTTokenNameEvent represents a TokenNameEvent event raised by the RandomWalkNFT contract.
type RandomWalkNFTTokenNameEvent struct {
	TokenId *big.Int
	NewName string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenNameEvent is a free log retrieval operation binding the contract event 0x8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12.
//
// Solidity: event TokenNameEvent(uint256 tokenId, string newName)
func (_RandomWalkNFT *RandomWalkNFTFilterer) FilterTokenNameEvent(opts *bind.FilterOpts) (*RandomWalkNFTTokenNameEventIterator, error) {

	logs, sub, err := _RandomWalkNFT.contract.FilterLogs(opts, "TokenNameEvent")
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTTokenNameEventIterator{contract: _RandomWalkNFT.contract, event: "TokenNameEvent", logs: logs, sub: sub}, nil
}

// WatchTokenNameEvent is a free log subscription operation binding the contract event 0x8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12.
//
// Solidity: event TokenNameEvent(uint256 tokenId, string newName)
func (_RandomWalkNFT *RandomWalkNFTFilterer) WatchTokenNameEvent(opts *bind.WatchOpts, sink chan<- *RandomWalkNFTTokenNameEvent) (event.Subscription, error) {

	logs, sub, err := _RandomWalkNFT.contract.WatchLogs(opts, "TokenNameEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomWalkNFTTokenNameEvent)
				if err := _RandomWalkNFT.contract.UnpackLog(event, "TokenNameEvent", log); err != nil {
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

// ParseTokenNameEvent is a log parse operation binding the contract event 0x8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12.
//
// Solidity: event TokenNameEvent(uint256 tokenId, string newName)
func (_RandomWalkNFT *RandomWalkNFTFilterer) ParseTokenNameEvent(log types.Log) (*RandomWalkNFTTokenNameEvent, error) {
	event := new(RandomWalkNFTTokenNameEvent)
	if err := _RandomWalkNFT.contract.UnpackLog(event, "TokenNameEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomWalkNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RandomWalkNFT contract.
type RandomWalkNFTTransferIterator struct {
	Event *RandomWalkNFTTransfer // Event containing the contract specifics and raw log

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
func (it *RandomWalkNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomWalkNFTTransfer)
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
		it.Event = new(RandomWalkNFTTransfer)
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
func (it *RandomWalkNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomWalkNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomWalkNFTTransfer represents a Transfer event raised by the RandomWalkNFT contract.
type RandomWalkNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_RandomWalkNFT *RandomWalkNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*RandomWalkNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTTransferIterator{contract: _RandomWalkNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_RandomWalkNFT *RandomWalkNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RandomWalkNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomWalkNFTTransfer)
				if err := _RandomWalkNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_RandomWalkNFT *RandomWalkNFTFilterer) ParseTransfer(log types.Log) (*RandomWalkNFTTransfer, error) {
	event := new(RandomWalkNFTTransfer)
	if err := _RandomWalkNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomWalkNFTWithdrawalEventIterator is returned from FilterWithdrawalEvent and is used to iterate over the raw logs and unpacked data for WithdrawalEvent events raised by the RandomWalkNFT contract.
type RandomWalkNFTWithdrawalEventIterator struct {
	Event *RandomWalkNFTWithdrawalEvent // Event containing the contract specifics and raw log

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
func (it *RandomWalkNFTWithdrawalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomWalkNFTWithdrawalEvent)
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
		it.Event = new(RandomWalkNFTWithdrawalEvent)
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
func (it *RandomWalkNFTWithdrawalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomWalkNFTWithdrawalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomWalkNFTWithdrawalEvent represents a WithdrawalEvent event raised by the RandomWalkNFT contract.
type RandomWalkNFTWithdrawalEvent struct {
	TokenId     *big.Int
	Destination common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalEvent is a free log retrieval operation binding the contract event 0xa11b556ace4b11a5cae8675a293b51e8cde3a06387d34010861789dfd9e9abc7.
//
// Solidity: event WithdrawalEvent(uint256 indexed tokenId, address destination, uint256 amount)
func (_RandomWalkNFT *RandomWalkNFTFilterer) FilterWithdrawalEvent(opts *bind.FilterOpts, tokenId []*big.Int) (*RandomWalkNFTWithdrawalEventIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.FilterLogs(opts, "WithdrawalEvent", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RandomWalkNFTWithdrawalEventIterator{contract: _RandomWalkNFT.contract, event: "WithdrawalEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawalEvent is a free log subscription operation binding the contract event 0xa11b556ace4b11a5cae8675a293b51e8cde3a06387d34010861789dfd9e9abc7.
//
// Solidity: event WithdrawalEvent(uint256 indexed tokenId, address destination, uint256 amount)
func (_RandomWalkNFT *RandomWalkNFTFilterer) WatchWithdrawalEvent(opts *bind.WatchOpts, sink chan<- *RandomWalkNFTWithdrawalEvent, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RandomWalkNFT.contract.WatchLogs(opts, "WithdrawalEvent", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomWalkNFTWithdrawalEvent)
				if err := _RandomWalkNFT.contract.UnpackLog(event, "WithdrawalEvent", log); err != nil {
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

// ParseWithdrawalEvent is a log parse operation binding the contract event 0xa11b556ace4b11a5cae8675a293b51e8cde3a06387d34010861789dfd9e9abc7.
//
// Solidity: event WithdrawalEvent(uint256 indexed tokenId, address destination, uint256 amount)
func (_RandomWalkNFT *RandomWalkNFTFilterer) ParseWithdrawalEvent(log types.Log) (*RandomWalkNFTWithdrawalEvent, error) {
	event := new(RandomWalkNFTWithdrawalEvent)
	if err := _RandomWalkNFT.contract.UnpackLog(event, "WithdrawalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

