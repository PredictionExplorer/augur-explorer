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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroBalance\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CharityAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonationReceivedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"charityAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonationSentEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"charityAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setCharityAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523461002757610011610036565b61001961002c565b610a3f61022a8239610a3f90f35b610032565b60405190565b5f80fd5b610046610041610110565b610048565b565b610051906100ba565b565b90565b60018060a01b031690565b90565b61007861007361007d92610053565b610061565b610056565b90565b61008990610064565b90565b61009590610056565b90565b6100a19061008c565b9052565b91906100b8905f60208501940190610098565b565b806100d56100cf6100ca5f610080565b61008c565b9161008c565b146100e5576100e3906101ca565b565b6101086100f15f610080565b5f918291631e4fbdf760e01b8352600483016100a5565b0390fd5b5f90565b61011861010c565b503390565b5f1c90565b60018060a01b031690565b61013961013e9161011d565b610122565b90565b61014b905461012d565b90565b5f1b90565b9061016460018060a01b039161014e565b9181191691161790565b61018261017d61018792610056565b610061565b610056565b90565b6101939061016e565b90565b61019f9061018a565b90565b90565b906101ba6101b56101c192610196565b6101a2565b8254610153565b9055565b5f0190565b6101d35f610141565b6101dd825f6101a5565b9061021161020b7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093610196565b91610196565b9161021a61002c565b80610224816101c5565b0390a356fe6080604052600436101561001d575b366102a25761001b610901565b005b6100275f35610086565b80630c9be46d14610081578063715018a61461007c5780638da5cb5b14610077578063afcf2fc414610072578063b46300ec1461006d5763f2fde38b0361000e5761026f565b61023c565b610207565b61018e565b610139565b6100f7565b60e01c90565b60405190565b5f80fd5b5f80fd5b60018060a01b031690565b6100ae9061009a565b90565b6100ba816100a5565b036100c157565b5f80fd5b905035906100d2826100b1565b565b906020828203126100ed576100ea915f016100c5565b90565b610096565b5f0190565b346101255761010f61010a3660046100d4565b61045c565b61011761008c565b80610121816100f2565b0390f35b610092565b5f91031261013457565b610096565b346101675761014936600461012a565b61015161048c565b61015961008c565b80610163816100f2565b0390f35b610092565b610175906100a5565b9052565b919061018c905f6020850194019061016c565b565b346101be5761019e36600461012a565b6101ba6101a96104c0565b6101b161008c565b91829182610179565b0390f35b610092565b1c90565b60018060a01b031690565b6101e29060086101e793026101c3565b6101c7565b90565b906101f591546101d2565b90565b61020460015f906101ea565b90565b346102375761021736600461012a565b6102336102226101f8565b61022a61008c565b91829182610179565b0390f35b610092565b3461026a5761024c36600461012a565b6102546107ae565b61025c61008c565b80610266816100f2565b0390f35b610092565b3461029d576102876102823660046100d4565b6108f6565b61028f61008c565b80610299816100f2565b0390f35b610092565b5f80fd5b6102b7906102b261094f565b6103ea565b565b90565b90565b6102d36102ce6102d8926102b9565b6102bc565b61009a565b90565b6102e4906102bf565b90565b60209181520190565b5f7f5a65726f2d616464726573732077617320676976656e2e000000000000000000910152565b61032460176020926102e7565b61032d816102f0565b0190565b6103469060208101905f818303910152610317565b90565b1561035057565b61035861008c565b63eac0d38960e01b81528061036f60048201610331565b0390fd5b5f1b90565b9061038960018060a01b0391610373565b9181191691161790565b6103a76103a26103ac9261009a565b6102bc565b61009a565b90565b6103b890610393565b90565b6103c4906103af565b90565b90565b906103df6103da6103e6926103bb565b6103c7565b8254610378565b9055565b61040f816104086104026103fd5f6102db565b6100a5565b916100a5565b1415610349565b61041a8160016103ca565b6104447f1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c916103bb565b9061044d61008c565b80610457816100f2565b0390a2565b610465906102a6565b565b61046f61094f565b610477610479565b565b61048a6104855f6102db565b61099d565b565b610494610467565b565b5f90565b5f1c90565b6104ab6104b09161049a565b6101c7565b90565b6104bd905461049f565b90565b6104c8610496565b506104d25f6104b3565b90565b5f7f436861726974792061646472657373206e6f74207365742e0000000000000000910152565b61050960186020926102e7565b610512816104d5565b0190565b61052b9060208101905f8183039101526104fc565b90565b1561053557565b61053d61008c565b63eac0d38960e01b81528061055460048201610516565b0390fd5b610561906103af565b90565b90565b61057b610576610580926102b9565b6102bc565b610564565b90565b5f7f4e6f2066756e647320746f2073656e642e000000000000000000000000000000910152565b6105b760116020926102e7565b6105c081610583565b0190565b6105d99060208101905f8183039101526105aa565b90565b156105e357565b6105eb61008c565b63cb7450f760e01b815280610602600482016105c4565b0390fd5b905090565b6106165f8092610606565b0190565b6106239061060b565b90565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b9061064e90610626565b810190811067ffffffffffffffff82111761066857604052565b610630565b9061068061067961008c565b9283610644565b565b67ffffffffffffffff81116106a05761069c602091610626565b0190565b610630565b906106b76106b283610682565b61066d565b918252565b606090565b3d5f146106dc576106d13d6106a5565b903d5f602084013e5b565b6106e46106bc565b906106da565b5f7f5472616e7366657220746f2063686172697479206661696c65642e0000000000910152565b61071e601b6020926102e7565b610727816106ea565b0190565b61073490610564565b9052565b916040610769929493610762610757606083018381035f850152610711565b96602083019061016c565b019061072b565b565b15610774575050565b61079561077f61008c565b928392630aa7db6360e11b845260048401610738565b0390fd5b91906107ac905f6020850194019061072b565b565b6107dc6107bb60016104b3565b6107d56107cf6107ca5f6102db565b6100a5565b916100a5565b141561052e565b6107e530610558565b31610802816107fc6107f65f610567565b91610564565b116105dc565b6108435f8061081160016104b3565b8461081a61008c565b90816108258161061a565b03925af16108316106c1565b5061083c60016104b3565b839161076b565b61084d60016104b3565b61088c61087a7f44d398d152fa0735a428b13ebc78f79fe4cb1b4722292cd233e278552fa36d32926103bb565b9261088361008c565b91829182610799565b0390a2565b6108a29061089d61094f565b6108a4565b565b806108bf6108b96108b45f6102db565b6100a5565b916100a5565b146108cf576108cd9061099d565b565b6108f26108db5f6102db565b5f918291631e4fbdf760e01b835260048301610179565b0390fd5b6108ff90610891565b565b6109096109fc565b349061094a6109387f46ff3d75d4645bdbbae4cd6109ba42e6e1b80ea25e69d10472b357b733300368926103bb565b9261094161008c565b91829182610799565b0390a2565b6109576104c0565b61097061096a6109656109fc565b6100a5565b916100a5565b0361097757565b6109996109826109fc565b5f91829163118cdaa760e01b835260048301610179565b0390fd5b6109a65f6104b3565b6109b0825f6103ca565b906109e46109de7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936103bb565b916103bb565b916109ed61008c565b806109f7816100f2565b0390a3565b610a04610496565b50339056fea2646970667358221220d06a643d3e550d924881305c28a16370d61d9b62aebb59183832c182e889cb7764736f6c634300081b0033",
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

// SetCharityAddress is a paid mutator transaction binding the contract method 0x0c9be46d.
//
// Solidity: function setCharityAddress(address newValue_) returns()
func (_CharityWallet *CharityWalletTransactor) SetCharityAddress(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _CharityWallet.contract.Transact(opts, "setCharityAddress", newValue_)
}

// SetCharityAddress is a paid mutator transaction binding the contract method 0x0c9be46d.
//
// Solidity: function setCharityAddress(address newValue_) returns()
func (_CharityWallet *CharityWalletSession) SetCharityAddress(newValue_ common.Address) (*types.Transaction, error) {
	return _CharityWallet.Contract.SetCharityAddress(&_CharityWallet.TransactOpts, newValue_)
}

// SetCharityAddress is a paid mutator transaction binding the contract method 0x0c9be46d.
//
// Solidity: function setCharityAddress(address newValue_) returns()
func (_CharityWallet *CharityWalletTransactorSession) SetCharityAddress(newValue_ common.Address) (*types.Transaction, error) {
	return _CharityWallet.Contract.SetCharityAddress(&_CharityWallet.TransactOpts, newValue_)
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

// CharityWalletCharityAddressChangedIterator is returned from FilterCharityAddressChanged and is used to iterate over the raw logs and unpacked data for CharityAddressChanged events raised by the CharityWallet contract.
type CharityWalletCharityAddressChangedIterator struct {
	Event *CharityWalletCharityAddressChanged // Event containing the contract specifics and raw log

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
func (it *CharityWalletCharityAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharityWalletCharityAddressChanged)
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
		it.Event = new(CharityWalletCharityAddressChanged)
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
func (it *CharityWalletCharityAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharityWalletCharityAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharityWalletCharityAddressChanged represents a CharityAddressChanged event raised by the CharityWallet contract.
type CharityWalletCharityAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCharityAddressChanged is a free log retrieval operation binding the contract event 0x1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c.
//
// Solidity: event CharityAddressChanged(address indexed newValue)
func (_CharityWallet *CharityWalletFilterer) FilterCharityAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*CharityWalletCharityAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CharityWallet.contract.FilterLogs(opts, "CharityAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &CharityWalletCharityAddressChangedIterator{contract: _CharityWallet.contract, event: "CharityAddressChanged", logs: logs, sub: sub}, nil
}

// WatchCharityAddressChanged is a free log subscription operation binding the contract event 0x1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c.
//
// Solidity: event CharityAddressChanged(address indexed newValue)
func (_CharityWallet *CharityWalletFilterer) WatchCharityAddressChanged(opts *bind.WatchOpts, sink chan<- *CharityWalletCharityAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CharityWallet.contract.WatchLogs(opts, "CharityAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharityWalletCharityAddressChanged)
				if err := _CharityWallet.contract.UnpackLog(event, "CharityAddressChanged", log); err != nil {
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

// ParseCharityAddressChanged is a log parse operation binding the contract event 0x1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c.
//
// Solidity: event CharityAddressChanged(address indexed newValue)
func (_CharityWallet *CharityWalletFilterer) ParseCharityAddressChanged(log types.Log) (*CharityWalletCharityAddressChanged, error) {
	event := new(CharityWalletCharityAddressChanged)
	if err := _CharityWallet.contract.UnpackLog(event, "CharityAddressChanged", log); err != nil {
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
	DonorAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDonationReceivedEvent is a free log retrieval operation binding the contract event 0x46ff3d75d4645bdbbae4cd6109ba42e6e1b80ea25e69d10472b357b733300368.
//
// Solidity: event DonationReceivedEvent(address indexed donorAddress, uint256 amount)
func (_CharityWallet *CharityWalletFilterer) FilterDonationReceivedEvent(opts *bind.FilterOpts, donorAddress []common.Address) (*CharityWalletDonationReceivedEventIterator, error) {

	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}

	logs, sub, err := _CharityWallet.contract.FilterLogs(opts, "DonationReceivedEvent", donorAddressRule)
	if err != nil {
		return nil, err
	}
	return &CharityWalletDonationReceivedEventIterator{contract: _CharityWallet.contract, event: "DonationReceivedEvent", logs: logs, sub: sub}, nil
}

// WatchDonationReceivedEvent is a free log subscription operation binding the contract event 0x46ff3d75d4645bdbbae4cd6109ba42e6e1b80ea25e69d10472b357b733300368.
//
// Solidity: event DonationReceivedEvent(address indexed donorAddress, uint256 amount)
func (_CharityWallet *CharityWalletFilterer) WatchDonationReceivedEvent(opts *bind.WatchOpts, sink chan<- *CharityWalletDonationReceivedEvent, donorAddress []common.Address) (event.Subscription, error) {

	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}

	logs, sub, err := _CharityWallet.contract.WatchLogs(opts, "DonationReceivedEvent", donorAddressRule)
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
// Solidity: event DonationReceivedEvent(address indexed donorAddress, uint256 amount)
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
	CharityAddress common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDonationSentEvent is a free log retrieval operation binding the contract event 0x44d398d152fa0735a428b13ebc78f79fe4cb1b4722292cd233e278552fa36d32.
//
// Solidity: event DonationSentEvent(address indexed charityAddress, uint256 amount)
func (_CharityWallet *CharityWalletFilterer) FilterDonationSentEvent(opts *bind.FilterOpts, charityAddress []common.Address) (*CharityWalletDonationSentEventIterator, error) {

	var charityAddressRule []interface{}
	for _, charityAddressItem := range charityAddress {
		charityAddressRule = append(charityAddressRule, charityAddressItem)
	}

	logs, sub, err := _CharityWallet.contract.FilterLogs(opts, "DonationSentEvent", charityAddressRule)
	if err != nil {
		return nil, err
	}
	return &CharityWalletDonationSentEventIterator{contract: _CharityWallet.contract, event: "DonationSentEvent", logs: logs, sub: sub}, nil
}

// WatchDonationSentEvent is a free log subscription operation binding the contract event 0x44d398d152fa0735a428b13ebc78f79fe4cb1b4722292cd233e278552fa36d32.
//
// Solidity: event DonationSentEvent(address indexed charityAddress, uint256 amount)
func (_CharityWallet *CharityWalletFilterer) WatchDonationSentEvent(opts *bind.WatchOpts, sink chan<- *CharityWalletDonationSentEvent, charityAddress []common.Address) (event.Subscription, error) {

	var charityAddressRule []interface{}
	for _, charityAddressItem := range charityAddress {
		charityAddressRule = append(charityAddressRule, charityAddressItem)
	}

	logs, sub, err := _CharityWallet.contract.WatchLogs(opts, "DonationSentEvent", charityAddressRule)
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
// Solidity: event DonationSentEvent(address indexed charityAddress, uint256 amount)
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CharityAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonationReceivedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"charityAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonationSentEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setCharityAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// SetCharityAddress is a paid mutator transaction binding the contract method 0x0c9be46d.
//
// Solidity: function setCharityAddress(address newValue_) returns()
func (_ICharityWallet *ICharityWalletTransactor) SetCharityAddress(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _ICharityWallet.contract.Transact(opts, "setCharityAddress", newValue_)
}

// SetCharityAddress is a paid mutator transaction binding the contract method 0x0c9be46d.
//
// Solidity: function setCharityAddress(address newValue_) returns()
func (_ICharityWallet *ICharityWalletSession) SetCharityAddress(newValue_ common.Address) (*types.Transaction, error) {
	return _ICharityWallet.Contract.SetCharityAddress(&_ICharityWallet.TransactOpts, newValue_)
}

// SetCharityAddress is a paid mutator transaction binding the contract method 0x0c9be46d.
//
// Solidity: function setCharityAddress(address newValue_) returns()
func (_ICharityWallet *ICharityWalletTransactorSession) SetCharityAddress(newValue_ common.Address) (*types.Transaction, error) {
	return _ICharityWallet.Contract.SetCharityAddress(&_ICharityWallet.TransactOpts, newValue_)
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

// ICharityWalletCharityAddressChangedIterator is returned from FilterCharityAddressChanged and is used to iterate over the raw logs and unpacked data for CharityAddressChanged events raised by the ICharityWallet contract.
type ICharityWalletCharityAddressChangedIterator struct {
	Event *ICharityWalletCharityAddressChanged // Event containing the contract specifics and raw log

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
func (it *ICharityWalletCharityAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICharityWalletCharityAddressChanged)
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
		it.Event = new(ICharityWalletCharityAddressChanged)
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
func (it *ICharityWalletCharityAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICharityWalletCharityAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICharityWalletCharityAddressChanged represents a CharityAddressChanged event raised by the ICharityWallet contract.
type ICharityWalletCharityAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCharityAddressChanged is a free log retrieval operation binding the contract event 0x1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c.
//
// Solidity: event CharityAddressChanged(address indexed newValue)
func (_ICharityWallet *ICharityWalletFilterer) FilterCharityAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*ICharityWalletCharityAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _ICharityWallet.contract.FilterLogs(opts, "CharityAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &ICharityWalletCharityAddressChangedIterator{contract: _ICharityWallet.contract, event: "CharityAddressChanged", logs: logs, sub: sub}, nil
}

// WatchCharityAddressChanged is a free log subscription operation binding the contract event 0x1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c.
//
// Solidity: event CharityAddressChanged(address indexed newValue)
func (_ICharityWallet *ICharityWalletFilterer) WatchCharityAddressChanged(opts *bind.WatchOpts, sink chan<- *ICharityWalletCharityAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _ICharityWallet.contract.WatchLogs(opts, "CharityAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICharityWalletCharityAddressChanged)
				if err := _ICharityWallet.contract.UnpackLog(event, "CharityAddressChanged", log); err != nil {
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

// ParseCharityAddressChanged is a log parse operation binding the contract event 0x1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c.
//
// Solidity: event CharityAddressChanged(address indexed newValue)
func (_ICharityWallet *ICharityWalletFilterer) ParseCharityAddressChanged(log types.Log) (*ICharityWalletCharityAddressChanged, error) {
	event := new(ICharityWalletCharityAddressChanged)
	if err := _ICharityWallet.contract.UnpackLog(event, "CharityAddressChanged", log); err != nil {
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
	DonorAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDonationReceivedEvent is a free log retrieval operation binding the contract event 0x46ff3d75d4645bdbbae4cd6109ba42e6e1b80ea25e69d10472b357b733300368.
//
// Solidity: event DonationReceivedEvent(address indexed donorAddress, uint256 amount)
func (_ICharityWallet *ICharityWalletFilterer) FilterDonationReceivedEvent(opts *bind.FilterOpts, donorAddress []common.Address) (*ICharityWalletDonationReceivedEventIterator, error) {

	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}

	logs, sub, err := _ICharityWallet.contract.FilterLogs(opts, "DonationReceivedEvent", donorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ICharityWalletDonationReceivedEventIterator{contract: _ICharityWallet.contract, event: "DonationReceivedEvent", logs: logs, sub: sub}, nil
}

// WatchDonationReceivedEvent is a free log subscription operation binding the contract event 0x46ff3d75d4645bdbbae4cd6109ba42e6e1b80ea25e69d10472b357b733300368.
//
// Solidity: event DonationReceivedEvent(address indexed donorAddress, uint256 amount)
func (_ICharityWallet *ICharityWalletFilterer) WatchDonationReceivedEvent(opts *bind.WatchOpts, sink chan<- *ICharityWalletDonationReceivedEvent, donorAddress []common.Address) (event.Subscription, error) {

	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}

	logs, sub, err := _ICharityWallet.contract.WatchLogs(opts, "DonationReceivedEvent", donorAddressRule)
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
// Solidity: event DonationReceivedEvent(address indexed donorAddress, uint256 amount)
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
	CharityAddress common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDonationSentEvent is a free log retrieval operation binding the contract event 0x44d398d152fa0735a428b13ebc78f79fe4cb1b4722292cd233e278552fa36d32.
//
// Solidity: event DonationSentEvent(address indexed charityAddress, uint256 amount)
func (_ICharityWallet *ICharityWalletFilterer) FilterDonationSentEvent(opts *bind.FilterOpts, charityAddress []common.Address) (*ICharityWalletDonationSentEventIterator, error) {

	var charityAddressRule []interface{}
	for _, charityAddressItem := range charityAddress {
		charityAddressRule = append(charityAddressRule, charityAddressItem)
	}

	logs, sub, err := _ICharityWallet.contract.FilterLogs(opts, "DonationSentEvent", charityAddressRule)
	if err != nil {
		return nil, err
	}
	return &ICharityWalletDonationSentEventIterator{contract: _ICharityWallet.contract, event: "DonationSentEvent", logs: logs, sub: sub}, nil
}

// WatchDonationSentEvent is a free log subscription operation binding the contract event 0x44d398d152fa0735a428b13ebc78f79fe4cb1b4722292cd233e278552fa36d32.
//
// Solidity: event DonationSentEvent(address indexed charityAddress, uint256 amount)
func (_ICharityWallet *ICharityWalletFilterer) WatchDonationSentEvent(opts *bind.WatchOpts, sink chan<- *ICharityWalletDonationSentEvent, charityAddress []common.Address) (event.Subscription, error) {

	var charityAddressRule []interface{}
	for _, charityAddressItem := range charityAddress {
		charityAddressRule = append(charityAddressRule, charityAddressItem)
	}

	logs, sub, err := _ICharityWallet.contract.WatchLogs(opts, "DonationSentEvent", charityAddressRule)
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
// Solidity: event DonationSentEvent(address indexed charityAddress, uint256 amount)
func (_ICharityWallet *ICharityWalletFilterer) ParseDonationSentEvent(log types.Log) (*ICharityWalletDonationSentEvent, error) {
	event := new(ICharityWalletDonationSentEvent)
	if err := _ICharityWallet.contract.UnpackLog(event, "DonationSentEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
