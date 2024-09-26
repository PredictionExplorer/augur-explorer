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

// CharityWalletMetaData contains all meta data concerning the CharityWallet contract.
var CharityWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroBalance\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCharityAddress\",\"type\":\"address\"}],\"name\":\"CharityUpdatedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonationReceivedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"charity\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonationSentEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"charityAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCharityAddress\",\"type\":\"address\"}],\"name\":\"setCharity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523461002757610011610036565b61001961002c565b610a496102128239610a4990f35b610032565b60405190565b5f80fd5b61003f33610041565b565b61004a906100b3565b565b90565b60018060a01b031690565b90565b61007161006c6100769261004c565b61005a565b61004f565b90565b6100829061005d565b90565b61008e9061004f565b90565b61009a90610085565b9052565b91906100b1905f60208501940190610091565b565b806100ce6100c86100c35f610079565b610085565b91610085565b146100de576100dc906101b2565b565b6101016100ea5f610079565b5f918291631e4fbdf760e01b83526004830161009e565b0390fd5b5f1c90565b60018060a01b031690565b61012161012691610105565b61010a565b90565b6101339054610115565b90565b5f1b90565b9061014c60018060a01b0391610136565b9181191691161790565b61016a61016561016f9261004f565b61005a565b61004f565b90565b61017b90610156565b90565b61018790610172565b90565b90565b906101a261019d6101a99261017e565b61018a565b825461013b565b9055565b5f0190565b6101bb5f610129565b6101c5825f61018d565b906101f96101f37f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09361017e565b9161017e565b9161020261002c565b8061020c816101ad565b0390a356fe6080604052600436101561001d575b366102a25761001b61090b565b005b6100275f35610086565b8063715018a6146100815780638da5cb5b1461007c578063afcf2fc414610077578063b46300ec14610072578063f2fde38b1461006d5763fb6f71a30361000e5761026f565b61023c565b6101c8565b610193565b61011a565b6100ae565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f9103126100a457565b610096565b5f0190565b346100dc576100be36600461009a565b6100c66102f9565b6100ce61008c565b806100d8816100a9565b0390f35b610092565b60018060a01b031690565b6100f5906100e1565b90565b610101906100ec565b9052565b9190610118905f602085019401906100f8565b565b3461014a5761012a36600461009a565b61014661013561032d565b61013d61008c565b91829182610105565b0390f35b610092565b1c90565b60018060a01b031690565b61016e906008610173930261014f565b610153565b90565b90610181915461015e565b90565b61019060015f90610176565b90565b346101c3576101a336600461009a565b6101bf6101ae610184565b6101b661008c565b91829182610105565b0390f35b610092565b346101f6576101d836600461009a565b6101e0610658565b6101e861008c565b806101f2816100a9565b0390f35b610092565b610204816100ec565b0361020b57565b5f80fd5b9050359061021c826101fb565b565b9060208282031261023757610234915f0161020f565b90565b610096565b3461026a5761025461024f36600461021e565b6107a0565b61025c61008c565b80610266816100a9565b0390f35b610092565b3461029d5761028761028236600461021e565b610900565b61028f61008c565b80610299816100a9565b0390f35b610092565b5f80fd5b6102ae610959565b6102b66102e6565b565b90565b90565b6102d26102cd6102d7926102b8565b6102bb565b6100e1565b90565b6102e3906102be565b90565b6102f76102f25f6102da565b6109a7565b565b6103016102a6565b565b5f90565b5f1c90565b61031861031d91610307565b610153565b90565b61032a905461030c565b90565b610335610303565b5061033f5f610320565b90565b60209181520190565b5f7f436861726974792061646472657373206e6f74207365742e0000000000000000910152565b61037f6018602092610342565b6103888161034b565b0190565b6103a19060208101905f818303910152610372565b90565b156103ab57565b6103b361008c565b63eac0d38960e01b8152806103ca6004820161038c565b0390fd5b6103e26103dd6103e7926100e1565b6102bb565b6100e1565b90565b6103f3906103ce565b90565b6103ff906103ea565b90565b90565b61041961041461041e926102b8565b6102bb565b610402565b90565b5f7f4e6f2066756e647320746f2073656e642e000000000000000000000000000000910152565b6104556011602092610342565b61045e81610421565b0190565b6104779060208101905f818303910152610448565b90565b1561048157565b61048961008c565b63cb7450f760e01b8152806104a060048201610462565b0390fd5b905090565b6104b45f80926104a4565b0190565b6104c1906104a9565b90565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906104ec906104c4565b810190811067ffffffffffffffff82111761050657604052565b6104ce565b9061051e61051761008c565b92836104e2565b565b67ffffffffffffffff811161053e5761053a6020916104c4565b0190565b6104ce565b9061055561055083610520565b61050b565b918252565b606090565b3d5f1461057a5761056f3d610543565b903d5f602084013e5b565b61058261055a565b90610578565b5f7f5472616e73666572206661696c65642e00000000000000000000000000000000910152565b6105bc6010602092610342565b6105c581610588565b0190565b6105d290610402565b9052565b9160406106079294936106006105f5606083018381035f8501526105af565b9660208301906105c9565b01906100f8565b565b15610612575050565b61063361061d61008c565b92839263310a0fbb60e21b8452600484016105d6565b0390fd5b610640906103ea565b90565b9190610656905f602085019401906105c9565b565b6106866106656001610320565b61067f6106796106745f6102da565b6100ec565b916100ec565b14156103a4565b61068f306103f6565b316106ac816106a66106a05f610405565b91610402565b1161047a565b6106ed5f806106bb6001610320565b846106c461008c565b90816106cf816104b8565b03925af16106db61055f565b50826106e76001610320565b91610609565b6106f76001610320565b6107366107247f44d398d152fa0735a428b13ebc78f79fe4cb1b4722292cd233e278552fa36d3292610637565b9261072d61008c565b91829182610643565b0390a2565b61074c90610747610959565b61074e565b565b8061076961076361075e5f6102da565b6100ec565b916100ec565b1461077957610777906109a7565b565b61079c6107855f6102da565b5f918291631e4fbdf760e01b835260048301610105565b0390fd5b6107a99061073b565b565b6107bc906107b7610959565b610884565b565b5f7f5a65726f2d616464726573732077617320676976656e2e000000000000000000910152565b6107f26017602092610342565b6107fb816107be565b0190565b6108149060208101905f8183039101526107e5565b90565b1561081e57565b61082661008c565b63eac0d38960e01b81528061083d600482016107ff565b0390fd5b5f1b90565b9061085760018060a01b0391610841565b9181191691161790565b90565b9061087961087461088092610637565b610861565b8254610846565b9055565b6108b4906108ad816108a66108a061089b5f6102da565b6100ec565b916100ec565b1415610817565b6001610864565b6108be6001610320565b6108e87fa0bd6b2fdbf082ae2356710c23fc8d76d56d418cecb4514d119c77a8617b4ffe91610637565b906108f161008c565b806108fb816100a9565b0390a2565b610909906107ab565b565b610913610a06565b34906109546109427f46ff3d75d4645bdbbae4cd6109ba42e6e1b80ea25e69d10472b357b73330036892610637565b9261094b61008c565b91829182610643565b0390a2565b61096161032d565b61097a61097461096f610a06565b6100ec565b916100ec565b0361098157565b6109a361098c610a06565b5f91829163118cdaa760e01b835260048301610105565b0390fd5b6109b05f610320565b6109ba825f610864565b906109ee6109e87f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093610637565b91610637565b916109f761008c565b80610a01816100a9565b0390a3565b610a0e610303565b50339056fea2646970667358221220d8d307725afd6bd5d4ea479187c59c8f3799c037692eaa8553d595ce22cbeb3364736f6c634300081a0033",
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
	parsed, err := CharityWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// CharityAddress is a free data retrieval call binding the contract method 0xafcf2fc4.
//
// Solidity: function charityAddress() view returns(address)
func (_CharityWallet *CharityWalletCaller) CharityAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CharityWallet.contract.Call(opts, &out, "charityAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CharityAddress is a free data retrieval call binding the contract method 0xafcf2fc4.
//
// Solidity: function charityAddress() view returns(address)
func (_CharityWallet *CharityWalletSession) CharityAddress() (common.Address, error) {
	return _CharityWallet.Contract.CharityAddress(&_CharityWallet.CallOpts)
}

// CharityAddress is a free data retrieval call binding the contract method 0xafcf2fc4.
//
// Solidity: function charityAddress() view returns(address)
func (_CharityWallet *CharityWalletCallerSession) CharityAddress() (common.Address, error) {
	return _CharityWallet.Contract.CharityAddress(&_CharityWallet.CallOpts)
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

// ICharityWalletMetaData contains all meta data concerning the ICharityWallet contract.
var ICharityWalletMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCharityAddress\",\"type\":\"address\"}],\"name\":\"CharityUpdatedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonationReceivedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"charity\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonationSentEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCharityAddress\",\"type\":\"address\"}],\"name\":\"setCharity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ICharityWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use ICharityWalletMetaData.ABI instead.
var ICharityWalletABI = ICharityWalletMetaData.ABI

// ICharityWallet is an auto generated Go binding around an Ethereum contract.
type ICharityWallet struct {
	ICharityWalletCaller     // Read-only binding to the contract
	ICharityWalletTransactor // Write-only binding to the contract
	ICharityWalletFilterer   // Log filterer for contract events
}

// ICharityWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICharityWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICharityWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICharityWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICharityWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICharityWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICharityWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICharityWalletSession struct {
	Contract     *ICharityWallet   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICharityWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICharityWalletCallerSession struct {
	Contract *ICharityWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ICharityWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICharityWalletTransactorSession struct {
	Contract     *ICharityWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ICharityWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICharityWalletRaw struct {
	Contract *ICharityWallet // Generic contract binding to access the raw methods on
}

// ICharityWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICharityWalletCallerRaw struct {
	Contract *ICharityWalletCaller // Generic read-only contract binding to access the raw methods on
}

// ICharityWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICharityWalletTransactorRaw struct {
	Contract *ICharityWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICharityWallet creates a new instance of ICharityWallet, bound to a specific deployed contract.
func NewICharityWallet(address common.Address, backend bind.ContractBackend) (*ICharityWallet, error) {
	contract, err := bindICharityWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICharityWallet{ICharityWalletCaller: ICharityWalletCaller{contract: contract}, ICharityWalletTransactor: ICharityWalletTransactor{contract: contract}, ICharityWalletFilterer: ICharityWalletFilterer{contract: contract}}, nil
}

// NewICharityWalletCaller creates a new read-only instance of ICharityWallet, bound to a specific deployed contract.
func NewICharityWalletCaller(address common.Address, caller bind.ContractCaller) (*ICharityWalletCaller, error) {
	contract, err := bindICharityWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICharityWalletCaller{contract: contract}, nil
}

// NewICharityWalletTransactor creates a new write-only instance of ICharityWallet, bound to a specific deployed contract.
func NewICharityWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*ICharityWalletTransactor, error) {
	contract, err := bindICharityWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICharityWalletTransactor{contract: contract}, nil
}

// NewICharityWalletFilterer creates a new log filterer instance of ICharityWallet, bound to a specific deployed contract.
func NewICharityWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*ICharityWalletFilterer, error) {
	contract, err := bindICharityWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICharityWalletFilterer{contract: contract}, nil
}

// bindICharityWallet binds a generic wrapper to an already deployed contract.
func bindICharityWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICharityWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICharityWallet *ICharityWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICharityWallet.Contract.ICharityWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICharityWallet *ICharityWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICharityWallet.Contract.ICharityWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICharityWallet *ICharityWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICharityWallet.Contract.ICharityWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICharityWallet *ICharityWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICharityWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICharityWallet *ICharityWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICharityWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICharityWallet *ICharityWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICharityWallet.Contract.contract.Transact(opts, method, params...)
}

// Send is a paid mutator transaction binding the contract method 0xb46300ec.
//
// Solidity: function send() returns()
func (_ICharityWallet *ICharityWalletTransactor) Send(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICharityWallet.contract.Transact(opts, "send")
}

// Send is a paid mutator transaction binding the contract method 0xb46300ec.
//
// Solidity: function send() returns()
func (_ICharityWallet *ICharityWalletSession) Send() (*types.Transaction, error) {
	return _ICharityWallet.Contract.Send(&_ICharityWallet.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xb46300ec.
//
// Solidity: function send() returns()
func (_ICharityWallet *ICharityWalletTransactorSession) Send() (*types.Transaction, error) {
	return _ICharityWallet.Contract.Send(&_ICharityWallet.TransactOpts)
}

// SetCharity is a paid mutator transaction binding the contract method 0xfb6f71a3.
//
// Solidity: function setCharity(address newCharityAddress) returns()
func (_ICharityWallet *ICharityWalletTransactor) SetCharity(opts *bind.TransactOpts, newCharityAddress common.Address) (*types.Transaction, error) {
	return _ICharityWallet.contract.Transact(opts, "setCharity", newCharityAddress)
}

// SetCharity is a paid mutator transaction binding the contract method 0xfb6f71a3.
//
// Solidity: function setCharity(address newCharityAddress) returns()
func (_ICharityWallet *ICharityWalletSession) SetCharity(newCharityAddress common.Address) (*types.Transaction, error) {
	return _ICharityWallet.Contract.SetCharity(&_ICharityWallet.TransactOpts, newCharityAddress)
}

// SetCharity is a paid mutator transaction binding the contract method 0xfb6f71a3.
//
// Solidity: function setCharity(address newCharityAddress) returns()
func (_ICharityWallet *ICharityWalletTransactorSession) SetCharity(newCharityAddress common.Address) (*types.Transaction, error) {
	return _ICharityWallet.Contract.SetCharity(&_ICharityWallet.TransactOpts, newCharityAddress)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ICharityWallet *ICharityWalletTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICharityWallet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ICharityWallet *ICharityWalletSession) Receive() (*types.Transaction, error) {
	return _ICharityWallet.Contract.Receive(&_ICharityWallet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ICharityWallet *ICharityWalletTransactorSession) Receive() (*types.Transaction, error) {
	return _ICharityWallet.Contract.Receive(&_ICharityWallet.TransactOpts)
}

// ICharityWalletCharityUpdatedEventIterator is returned from FilterCharityUpdatedEvent and is used to iterate over the raw logs and unpacked data for CharityUpdatedEvent events raised by the ICharityWallet contract.
type ICharityWalletCharityUpdatedEventIterator struct {
	Event *ICharityWalletCharityUpdatedEvent // Event containing the contract specifics and raw log

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
func (it *ICharityWalletCharityUpdatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICharityWalletCharityUpdatedEvent)
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
		it.Event = new(ICharityWalletCharityUpdatedEvent)
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
func (it *ICharityWalletCharityUpdatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICharityWalletCharityUpdatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICharityWalletCharityUpdatedEvent represents a CharityUpdatedEvent event raised by the ICharityWallet contract.
type ICharityWalletCharityUpdatedEvent struct {
	NewCharityAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCharityUpdatedEvent is a free log retrieval operation binding the contract event 0xa0bd6b2fdbf082ae2356710c23fc8d76d56d418cecb4514d119c77a8617b4ffe.
//
// Solidity: event CharityUpdatedEvent(address indexed newCharityAddress)
func (_ICharityWallet *ICharityWalletFilterer) FilterCharityUpdatedEvent(opts *bind.FilterOpts, newCharityAddress []common.Address) (*ICharityWalletCharityUpdatedEventIterator, error) {

	var newCharityAddressRule []interface{}
	for _, newCharityAddressItem := range newCharityAddress {
		newCharityAddressRule = append(newCharityAddressRule, newCharityAddressItem)
	}

	logs, sub, err := _ICharityWallet.contract.FilterLogs(opts, "CharityUpdatedEvent", newCharityAddressRule)
	if err != nil {
		return nil, err
	}
	return &ICharityWalletCharityUpdatedEventIterator{contract: _ICharityWallet.contract, event: "CharityUpdatedEvent", logs: logs, sub: sub}, nil
}

// WatchCharityUpdatedEvent is a free log subscription operation binding the contract event 0xa0bd6b2fdbf082ae2356710c23fc8d76d56d418cecb4514d119c77a8617b4ffe.
//
// Solidity: event CharityUpdatedEvent(address indexed newCharityAddress)
func (_ICharityWallet *ICharityWalletFilterer) WatchCharityUpdatedEvent(opts *bind.WatchOpts, sink chan<- *ICharityWalletCharityUpdatedEvent, newCharityAddress []common.Address) (event.Subscription, error) {

	var newCharityAddressRule []interface{}
	for _, newCharityAddressItem := range newCharityAddress {
		newCharityAddressRule = append(newCharityAddressRule, newCharityAddressItem)
	}

	logs, sub, err := _ICharityWallet.contract.WatchLogs(opts, "CharityUpdatedEvent", newCharityAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICharityWalletCharityUpdatedEvent)
				if err := _ICharityWallet.contract.UnpackLog(event, "CharityUpdatedEvent", log); err != nil {
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
func (_ICharityWallet *ICharityWalletFilterer) ParseCharityUpdatedEvent(log types.Log) (*ICharityWalletCharityUpdatedEvent, error) {
	event := new(ICharityWalletCharityUpdatedEvent)
	if err := _ICharityWallet.contract.UnpackLog(event, "CharityUpdatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICharityWalletDonationReceivedEventIterator is returned from FilterDonationReceivedEvent and is used to iterate over the raw logs and unpacked data for DonationReceivedEvent events raised by the ICharityWallet contract.
type ICharityWalletDonationReceivedEventIterator struct {
	Event *ICharityWalletDonationReceivedEvent // Event containing the contract specifics and raw log

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
func (it *ICharityWalletDonationReceivedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICharityWalletDonationReceivedEvent)
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
		it.Event = new(ICharityWalletDonationReceivedEvent)
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
func (it *ICharityWalletDonationReceivedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICharityWalletDonationReceivedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICharityWalletDonationReceivedEvent represents a DonationReceivedEvent event raised by the ICharityWallet contract.
type ICharityWalletDonationReceivedEvent struct {
	Donor  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDonationReceivedEvent is a free log retrieval operation binding the contract event 0x46ff3d75d4645bdbbae4cd6109ba42e6e1b80ea25e69d10472b357b733300368.
//
// Solidity: event DonationReceivedEvent(address indexed donor, uint256 amount)
func (_ICharityWallet *ICharityWalletFilterer) FilterDonationReceivedEvent(opts *bind.FilterOpts, donor []common.Address) (*ICharityWalletDonationReceivedEventIterator, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _ICharityWallet.contract.FilterLogs(opts, "DonationReceivedEvent", donorRule)
	if err != nil {
		return nil, err
	}
	return &ICharityWalletDonationReceivedEventIterator{contract: _ICharityWallet.contract, event: "DonationReceivedEvent", logs: logs, sub: sub}, nil
}

// WatchDonationReceivedEvent is a free log subscription operation binding the contract event 0x46ff3d75d4645bdbbae4cd6109ba42e6e1b80ea25e69d10472b357b733300368.
//
// Solidity: event DonationReceivedEvent(address indexed donor, uint256 amount)
func (_ICharityWallet *ICharityWalletFilterer) WatchDonationReceivedEvent(opts *bind.WatchOpts, sink chan<- *ICharityWalletDonationReceivedEvent, donor []common.Address) (event.Subscription, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _ICharityWallet.contract.WatchLogs(opts, "DonationReceivedEvent", donorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICharityWalletDonationReceivedEvent)
				if err := _ICharityWallet.contract.UnpackLog(event, "DonationReceivedEvent", log); err != nil {
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
func (_ICharityWallet *ICharityWalletFilterer) ParseDonationReceivedEvent(log types.Log) (*ICharityWalletDonationReceivedEvent, error) {
	event := new(ICharityWalletDonationReceivedEvent)
	if err := _ICharityWallet.contract.UnpackLog(event, "DonationReceivedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICharityWalletDonationSentEventIterator is returned from FilterDonationSentEvent and is used to iterate over the raw logs and unpacked data for DonationSentEvent events raised by the ICharityWallet contract.
type ICharityWalletDonationSentEventIterator struct {
	Event *ICharityWalletDonationSentEvent // Event containing the contract specifics and raw log

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
func (it *ICharityWalletDonationSentEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICharityWalletDonationSentEvent)
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
		it.Event = new(ICharityWalletDonationSentEvent)
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
func (it *ICharityWalletDonationSentEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICharityWalletDonationSentEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICharityWalletDonationSentEvent represents a DonationSentEvent event raised by the ICharityWallet contract.
type ICharityWalletDonationSentEvent struct {
	Charity common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDonationSentEvent is a free log retrieval operation binding the contract event 0x44d398d152fa0735a428b13ebc78f79fe4cb1b4722292cd233e278552fa36d32.
//
// Solidity: event DonationSentEvent(address indexed charity, uint256 amount)
func (_ICharityWallet *ICharityWalletFilterer) FilterDonationSentEvent(opts *bind.FilterOpts, charity []common.Address) (*ICharityWalletDonationSentEventIterator, error) {

	var charityRule []interface{}
	for _, charityItem := range charity {
		charityRule = append(charityRule, charityItem)
	}

	logs, sub, err := _ICharityWallet.contract.FilterLogs(opts, "DonationSentEvent", charityRule)
	if err != nil {
		return nil, err
	}
	return &ICharityWalletDonationSentEventIterator{contract: _ICharityWallet.contract, event: "DonationSentEvent", logs: logs, sub: sub}, nil
}

// WatchDonationSentEvent is a free log subscription operation binding the contract event 0x44d398d152fa0735a428b13ebc78f79fe4cb1b4722292cd233e278552fa36d32.
//
// Solidity: event DonationSentEvent(address indexed charity, uint256 amount)
func (_ICharityWallet *ICharityWalletFilterer) WatchDonationSentEvent(opts *bind.WatchOpts, sink chan<- *ICharityWalletDonationSentEvent, charity []common.Address) (event.Subscription, error) {

	var charityRule []interface{}
	for _, charityItem := range charity {
		charityRule = append(charityRule, charityItem)
	}

	logs, sub, err := _ICharityWallet.contract.WatchLogs(opts, "DonationSentEvent", charityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICharityWalletDonationSentEvent)
				if err := _ICharityWallet.contract.UnpackLog(event, "DonationSentEvent", log); err != nil {
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
func (_ICharityWallet *ICharityWalletFilterer) ParseDonationSentEvent(log types.Log) (*ICharityWalletDonationSentEvent, error) {
	event := new(ICharityWalletDonationSentEvent)
	if err := _ICharityWallet.contract.UnpackLog(event, "DonationSentEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
