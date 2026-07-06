// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cosmicgame

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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CharityAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonationReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"charityAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsTransferredToCharity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"charityAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setCharityAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523461002757610011610036565b61001961002c565b61095261022a823961095290f35b610032565b60405190565b5f80fd5b610046610041610110565b610048565b565b610051906100ba565b565b90565b60018060a01b031690565b90565b61007861007361007d92610053565b610061565b610056565b90565b61008990610064565b90565b61009590610056565b90565b6100a19061008c565b9052565b91906100b8905f60208501940190610098565b565b806100d56100cf6100ca5f610080565b61008c565b9161008c565b146100e5576100e3906101ca565b565b6101086100f15f610080565b5f918291631e4fbdf760e01b8352600483016100a5565b0390fd5b5f90565b61011861010c565b503390565b5f1c90565b60018060a01b031690565b61013961013e9161011d565b610122565b90565b61014b905461012d565b90565b5f1b90565b9061016460018060a01b039161014e565b9181191691161790565b61018261017d61018792610056565b610061565b610056565b90565b6101939061016e565b90565b61019f9061018a565b90565b90565b906101ba6101b56101c192610196565b6101a2565b8254610153565b9055565b5f0190565b6101d35f610141565b6101dd825f6101a5565b9061021161020b7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093610196565b91610196565b9161021a61002c565b80610224816101c5565b0390a356fe6080604052600436101561001d575b366103295761001b610814565b005b6100275f35610096565b80630c9be46d14610091578063715018a61461008c5780638da5cb5b14610087578063a52c101e14610082578063afcf2fc41461007d578063b46300ec146100785763f2fde38b0361000e576102f6565b6102c3565b61028e565b610217565b61019e565b610149565b610107565b60e01c90565b60405190565b5f80fd5b5f80fd5b60018060a01b031690565b6100be906100aa565b90565b6100ca816100b5565b036100d157565b5f80fd5b905035906100e2826100c1565b565b906020828203126100fd576100fa915f016100d5565b90565b6100a6565b5f0190565b346101355761011f61011a3660046100e4565b610407565b61012761009c565b8061013181610102565b0390f35b6100a2565b5f91031261014457565b6100a6565b346101775761015936600461013a565b610161610462565b61016961009c565b8061017381610102565b0390f35b6100a2565b610185906100b5565b9052565b919061019c905f6020850194019061017c565b565b346101ce576101ae36600461013a565b6101ca6101b9610496565b6101c161009c565b91829182610189565b0390f35b6100a2565b90565b6101df816101d3565b036101e657565b5f80fd5b905035906101f7826101d6565b565b906020828203126102125761020f915f016101ea565b90565b6100a6565b346102455761022f61022a3660046101f9565b6106b6565b61023761009c565b8061024181610102565b0390f35b6100a2565b1c90565b60018060a01b031690565b61026990600861026e930261024a565b61024e565b90565b9061027c9154610259565b90565b61028b60015f90610271565b90565b346102be5761029e36600461013a565b6102ba6102a961027f565b6102b161009c565b91829182610189565b0390f35b6100a2565b346102f1576102d336600461013a565b6102db610790565b6102e361009c565b806102ed81610102565b0390f35b6100a2565b346103245761030e6103093660046100e4565b610809565b61031661009c565b8061032081610102565b0390f35b6100a2565b5f80fd5b61033e90610339610862565b6103ba565b565b5f1b90565b9061035660018060a01b0391610340565b9181191691161790565b90565b61037761037261037c926100aa565b610360565b6100aa565b90565b61038890610363565b90565b6103949061037f565b90565b90565b906103af6103aa6103b69261038b565b610397565b8254610345565b9055565b6103c581600161039a565b6103ef7f1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c9161038b565b906103f861009c565b8061040281610102565b0390a2565b6104109061032d565b565b61041a610862565b61042261044f565b565b90565b61043b61043661044092610424565b610360565b6100aa565b90565b61044c90610427565b90565b61046061045b5f610443565b6108b0565b565b61046a610412565b565b5f90565b5f1c90565b61048161048691610470565b61024e565b90565b6104939054610475565b90565b61049e61046c565b506104a85f610489565b90565b60209181520190565b5f7f436861726974792061646472657373206e6f74207365742e0000000000000000910152565b6104e860186020926104ab565b6104f1816104b4565b0190565b61050a9060208101905f8183039101526104db565b90565b1561051457565b61051c61009c565b63eac0d38960e01b815280610533600482016104f5565b0390fd5b610540906101d3565b9052565b9190610557905f60208501940190610537565b565b905090565b6105695f8092610559565b0190565b6105769061055e565b90565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906105a190610579565b810190811067ffffffffffffffff8211176105bb57604052565b610583565b906105d36105cc61009c565b9283610597565b565b67ffffffffffffffff81116105f3576105ef602091610579565b0190565b610583565b9061060a610605836105d5565b6105c0565b918252565b606090565b3d5f1461062f576106243d6105f8565b903d5f602084013e5b565b61063761060f565b9061062d565b151590565b5f7f455448207472616e7366657220746f2063686172697479206661696c65642e00910152565b610676601f6020926104ab565b61067f81610642565b0190565b9160406106b49294936106ad6106a2606083018381035f850152610669565b96602083019061017c565b0190610537565b565b6106c06001610489565b6106e5816106de6106d86106d35f610443565b6100b5565b916100b5565b141561050d565b8082906107276107157f1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d9261038b565b9261071e61009c565b91829182610544565b0390a26107575f80838561073961009c565b90816107448161056d565b03925af1610750610614565b501561063d565b61075f575050565b61078061076a61009c565b928392630aa7db6360e11b845260048401610683565b0390fd5b61078d9061037f565b90565b6107a261079c30610784565b316106b6565b565b6107b5906107b0610862565b6107b7565b565b806107d26107cc6107c75f610443565b6100b5565b916100b5565b146107e2576107e0906108b0565b565b6108056107ee5f610443565b5f918291631e4fbdf760e01b835260048301610189565b0390fd5b610812906107a4565b565b61081c61090f565b349061085d61084b7f264f630d9efa0d07053a31163641d9fcc0adafc9d9e76f1c37c2ce3a558d2c529261038b565b9261085461009c565b91829182610544565b0390a2565b61086a610496565b61088361087d61087861090f565b6100b5565b916100b5565b0361088a57565b6108ac61089561090f565b5f91829163118cdaa760e01b835260048301610189565b0390fd5b6108b95f610489565b6108c3825f61039a565b906108f76108f17f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09361038b565b9161038b565b9161090061009c565b8061090a81610102565b0390a3565b61091761046c565b50339056fea2646970667358221220edc8b0d700cdec00afa6ecc27fc01c5fd9de8a775abb24c29901fbd6cf13c20264736f6c634300081c0033",
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

// Send is a paid mutator transaction binding the contract method 0xa52c101e.
//
// Solidity: function send(uint256 amount_) returns()
func (_CharityWallet *CharityWalletTransactor) Send(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _CharityWallet.contract.Transact(opts, "send", amount_)
}

// Send is a paid mutator transaction binding the contract method 0xa52c101e.
//
// Solidity: function send(uint256 amount_) returns()
func (_CharityWallet *CharityWalletSession) Send(amount_ *big.Int) (*types.Transaction, error) {
	return _CharityWallet.Contract.Send(&_CharityWallet.TransactOpts, amount_)
}

// Send is a paid mutator transaction binding the contract method 0xa52c101e.
//
// Solidity: function send(uint256 amount_) returns()
func (_CharityWallet *CharityWalletTransactorSession) Send(amount_ *big.Int) (*types.Transaction, error) {
	return _CharityWallet.Contract.Send(&_CharityWallet.TransactOpts, amount_)
}

// Send0 is a paid mutator transaction binding the contract method 0xb46300ec.
//
// Solidity: function send() returns()
func (_CharityWallet *CharityWalletTransactor) Send0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CharityWallet.contract.Transact(opts, "send0")
}

// Send0 is a paid mutator transaction binding the contract method 0xb46300ec.
//
// Solidity: function send() returns()
func (_CharityWallet *CharityWalletSession) Send0() (*types.Transaction, error) {
	return _CharityWallet.Contract.Send0(&_CharityWallet.TransactOpts)
}

// Send0 is a paid mutator transaction binding the contract method 0xb46300ec.
//
// Solidity: function send() returns()
func (_CharityWallet *CharityWalletTransactorSession) Send0() (*types.Transaction, error) {
	return _CharityWallet.Contract.Send0(&_CharityWallet.TransactOpts)
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

// CharityWalletDonationReceivedIterator is returned from FilterDonationReceived and is used to iterate over the raw logs and unpacked data for DonationReceived events raised by the CharityWallet contract.
type CharityWalletDonationReceivedIterator struct {
	Event *CharityWalletDonationReceived // Event containing the contract specifics and raw log

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
func (it *CharityWalletDonationReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharityWalletDonationReceived)
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
		it.Event = new(CharityWalletDonationReceived)
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
func (it *CharityWalletDonationReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharityWalletDonationReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharityWalletDonationReceived represents a DonationReceived event raised by the CharityWallet contract.
type CharityWalletDonationReceived struct {
	DonorAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDonationReceived is a free log retrieval operation binding the contract event 0x264f630d9efa0d07053a31163641d9fcc0adafc9d9e76f1c37c2ce3a558d2c52.
//
// Solidity: event DonationReceived(address indexed donorAddress, uint256 amount)
func (_CharityWallet *CharityWalletFilterer) FilterDonationReceived(opts *bind.FilterOpts, donorAddress []common.Address) (*CharityWalletDonationReceivedIterator, error) {

	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}

	logs, sub, err := _CharityWallet.contract.FilterLogs(opts, "DonationReceived", donorAddressRule)
	if err != nil {
		return nil, err
	}
	return &CharityWalletDonationReceivedIterator{contract: _CharityWallet.contract, event: "DonationReceived", logs: logs, sub: sub}, nil
}

// WatchDonationReceived is a free log subscription operation binding the contract event 0x264f630d9efa0d07053a31163641d9fcc0adafc9d9e76f1c37c2ce3a558d2c52.
//
// Solidity: event DonationReceived(address indexed donorAddress, uint256 amount)
func (_CharityWallet *CharityWalletFilterer) WatchDonationReceived(opts *bind.WatchOpts, sink chan<- *CharityWalletDonationReceived, donorAddress []common.Address) (event.Subscription, error) {

	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}

	logs, sub, err := _CharityWallet.contract.WatchLogs(opts, "DonationReceived", donorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharityWalletDonationReceived)
				if err := _CharityWallet.contract.UnpackLog(event, "DonationReceived", log); err != nil {
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

// ParseDonationReceived is a log parse operation binding the contract event 0x264f630d9efa0d07053a31163641d9fcc0adafc9d9e76f1c37c2ce3a558d2c52.
//
// Solidity: event DonationReceived(address indexed donorAddress, uint256 amount)
func (_CharityWallet *CharityWalletFilterer) ParseDonationReceived(log types.Log) (*CharityWalletDonationReceived, error) {
	event := new(CharityWalletDonationReceived)
	if err := _CharityWallet.contract.UnpackLog(event, "DonationReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharityWalletFundsTransferredToCharityIterator is returned from FilterFundsTransferredToCharity and is used to iterate over the raw logs and unpacked data for FundsTransferredToCharity events raised by the CharityWallet contract.
type CharityWalletFundsTransferredToCharityIterator struct {
	Event *CharityWalletFundsTransferredToCharity // Event containing the contract specifics and raw log

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
func (it *CharityWalletFundsTransferredToCharityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharityWalletFundsTransferredToCharity)
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
		it.Event = new(CharityWalletFundsTransferredToCharity)
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
func (it *CharityWalletFundsTransferredToCharityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharityWalletFundsTransferredToCharityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharityWalletFundsTransferredToCharity represents a FundsTransferredToCharity event raised by the CharityWallet contract.
type CharityWalletFundsTransferredToCharity struct {
	CharityAddress common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFundsTransferredToCharity is a free log retrieval operation binding the contract event 0x1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d.
//
// Solidity: event FundsTransferredToCharity(address indexed charityAddress, uint256 amount)
func (_CharityWallet *CharityWalletFilterer) FilterFundsTransferredToCharity(opts *bind.FilterOpts, charityAddress []common.Address) (*CharityWalletFundsTransferredToCharityIterator, error) {

	var charityAddressRule []interface{}
	for _, charityAddressItem := range charityAddress {
		charityAddressRule = append(charityAddressRule, charityAddressItem)
	}

	logs, sub, err := _CharityWallet.contract.FilterLogs(opts, "FundsTransferredToCharity", charityAddressRule)
	if err != nil {
		return nil, err
	}
	return &CharityWalletFundsTransferredToCharityIterator{contract: _CharityWallet.contract, event: "FundsTransferredToCharity", logs: logs, sub: sub}, nil
}

// WatchFundsTransferredToCharity is a free log subscription operation binding the contract event 0x1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d.
//
// Solidity: event FundsTransferredToCharity(address indexed charityAddress, uint256 amount)
func (_CharityWallet *CharityWalletFilterer) WatchFundsTransferredToCharity(opts *bind.WatchOpts, sink chan<- *CharityWalletFundsTransferredToCharity, charityAddress []common.Address) (event.Subscription, error) {

	var charityAddressRule []interface{}
	for _, charityAddressItem := range charityAddress {
		charityAddressRule = append(charityAddressRule, charityAddressItem)
	}

	logs, sub, err := _CharityWallet.contract.WatchLogs(opts, "FundsTransferredToCharity", charityAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharityWalletFundsTransferredToCharity)
				if err := _CharityWallet.contract.UnpackLog(event, "FundsTransferredToCharity", log); err != nil {
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

// ParseFundsTransferredToCharity is a log parse operation binding the contract event 0x1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d.
//
// Solidity: event FundsTransferredToCharity(address indexed charityAddress, uint256 amount)
func (_CharityWallet *CharityWalletFilterer) ParseFundsTransferredToCharity(log types.Log) (*CharityWalletFundsTransferredToCharity, error) {
	event := new(CharityWalletFundsTransferredToCharity)
	if err := _CharityWallet.contract.UnpackLog(event, "FundsTransferredToCharity", log); err != nil {
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CharityAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonationReceived\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setCharityAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// Send is a paid mutator transaction binding the contract method 0xa52c101e.
//
// Solidity: function send(uint256 amount_) returns()
func (_ICharityWallet *ICharityWalletTransactor) Send(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _ICharityWallet.contract.Transact(opts, "send", amount_)
}

// Send is a paid mutator transaction binding the contract method 0xa52c101e.
//
// Solidity: function send(uint256 amount_) returns()
func (_ICharityWallet *ICharityWalletSession) Send(amount_ *big.Int) (*types.Transaction, error) {
	return _ICharityWallet.Contract.Send(&_ICharityWallet.TransactOpts, amount_)
}

// Send is a paid mutator transaction binding the contract method 0xa52c101e.
//
// Solidity: function send(uint256 amount_) returns()
func (_ICharityWallet *ICharityWalletTransactorSession) Send(amount_ *big.Int) (*types.Transaction, error) {
	return _ICharityWallet.Contract.Send(&_ICharityWallet.TransactOpts, amount_)
}

// Send0 is a paid mutator transaction binding the contract method 0xb46300ec.
//
// Solidity: function send() returns()
func (_ICharityWallet *ICharityWalletTransactor) Send0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICharityWallet.contract.Transact(opts, "send0")
}

// Send0 is a paid mutator transaction binding the contract method 0xb46300ec.
//
// Solidity: function send() returns()
func (_ICharityWallet *ICharityWalletSession) Send0() (*types.Transaction, error) {
	return _ICharityWallet.Contract.Send0(&_ICharityWallet.TransactOpts)
}

// Send0 is a paid mutator transaction binding the contract method 0xb46300ec.
//
// Solidity: function send() returns()
func (_ICharityWallet *ICharityWalletTransactorSession) Send0() (*types.Transaction, error) {
	return _ICharityWallet.Contract.Send0(&_ICharityWallet.TransactOpts)
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

// ICharityWalletDonationReceivedIterator is returned from FilterDonationReceived and is used to iterate over the raw logs and unpacked data for DonationReceived events raised by the ICharityWallet contract.
type ICharityWalletDonationReceivedIterator struct {
	Event *ICharityWalletDonationReceived // Event containing the contract specifics and raw log

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
func (it *ICharityWalletDonationReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICharityWalletDonationReceived)
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
		it.Event = new(ICharityWalletDonationReceived)
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
func (it *ICharityWalletDonationReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICharityWalletDonationReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICharityWalletDonationReceived represents a DonationReceived event raised by the ICharityWallet contract.
type ICharityWalletDonationReceived struct {
	DonorAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDonationReceived is a free log retrieval operation binding the contract event 0x264f630d9efa0d07053a31163641d9fcc0adafc9d9e76f1c37c2ce3a558d2c52.
//
// Solidity: event DonationReceived(address indexed donorAddress, uint256 amount)
func (_ICharityWallet *ICharityWalletFilterer) FilterDonationReceived(opts *bind.FilterOpts, donorAddress []common.Address) (*ICharityWalletDonationReceivedIterator, error) {

	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}

	logs, sub, err := _ICharityWallet.contract.FilterLogs(opts, "DonationReceived", donorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ICharityWalletDonationReceivedIterator{contract: _ICharityWallet.contract, event: "DonationReceived", logs: logs, sub: sub}, nil
}

// WatchDonationReceived is a free log subscription operation binding the contract event 0x264f630d9efa0d07053a31163641d9fcc0adafc9d9e76f1c37c2ce3a558d2c52.
//
// Solidity: event DonationReceived(address indexed donorAddress, uint256 amount)
func (_ICharityWallet *ICharityWalletFilterer) WatchDonationReceived(opts *bind.WatchOpts, sink chan<- *ICharityWalletDonationReceived, donorAddress []common.Address) (event.Subscription, error) {

	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}

	logs, sub, err := _ICharityWallet.contract.WatchLogs(opts, "DonationReceived", donorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICharityWalletDonationReceived)
				if err := _ICharityWallet.contract.UnpackLog(event, "DonationReceived", log); err != nil {
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

// ParseDonationReceived is a log parse operation binding the contract event 0x264f630d9efa0d07053a31163641d9fcc0adafc9d9e76f1c37c2ce3a558d2c52.
//
// Solidity: event DonationReceived(address indexed donorAddress, uint256 amount)
func (_ICharityWallet *ICharityWalletFilterer) ParseDonationReceived(log types.Log) (*ICharityWalletDonationReceived, error) {
	event := new(ICharityWalletDonationReceived)
	if err := _ICharityWallet.contract.UnpackLog(event, "DonationReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
