// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ERC20WrapperABI is the input ABI used to generate the binding from.
const ERC20WrapperABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_augurFoundry\",\"type\":\"address\"},{\"internalType\":\"contractIShareToken\",\"name\":\"_shareToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_cash\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"augurFoundry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cash\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shareToken\",\"outputs\":[{\"internalType\":\"contractIShareToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unWrapTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"wrapTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
//"
// ERC20Wrapper is an auto generated Go binding around an Ethereum contract.
type ERC20Wrapper struct {
	ERC20WrapperCaller     // Read-only binding to the contract
	ERC20WrapperTransactor // Write-only binding to the contract
	ERC20WrapperFilterer   // Log filterer for contract events
}

// ERC20WrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20WrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20WrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20WrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20WrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20WrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20WrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20WrapperSession struct {
	Contract     *ERC20Wrapper     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20WrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20WrapperCallerSession struct {
	Contract *ERC20WrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC20WrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20WrapperTransactorSession struct {
	Contract     *ERC20WrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC20WrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20WrapperRaw struct {
	Contract *ERC20Wrapper // Generic contract binding to access the raw methods on
}

// ERC20WrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20WrapperCallerRaw struct {
	Contract *ERC20WrapperCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20WrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20WrapperTransactorRaw struct {
	Contract *ERC20WrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Wrapper creates a new instance of ERC20Wrapper, bound to a specific deployed contract.
func NewERC20Wrapper(address common.Address, backend bind.ContractBackend) (*ERC20Wrapper, error) {
	contract, err := bindERC20Wrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Wrapper{ERC20WrapperCaller: ERC20WrapperCaller{contract: contract}, ERC20WrapperTransactor: ERC20WrapperTransactor{contract: contract}, ERC20WrapperFilterer: ERC20WrapperFilterer{contract: contract}}, nil
}

// NewERC20WrapperCaller creates a new read-only instance of ERC20Wrapper, bound to a specific deployed contract.
func NewERC20WrapperCaller(address common.Address, caller bind.ContractCaller) (*ERC20WrapperCaller, error) {
	contract, err := bindERC20Wrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20WrapperCaller{contract: contract}, nil
}

// NewERC20WrapperTransactor creates a new write-only instance of ERC20Wrapper, bound to a specific deployed contract.
func NewERC20WrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20WrapperTransactor, error) {
	contract, err := bindERC20Wrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20WrapperTransactor{contract: contract}, nil
}

// NewERC20WrapperFilterer creates a new log filterer instance of ERC20Wrapper, bound to a specific deployed contract.
func NewERC20WrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20WrapperFilterer, error) {
	contract, err := bindERC20Wrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20WrapperFilterer{contract: contract}, nil
}

// bindERC20Wrapper binds a generic wrapper to an already deployed contract.
func bindERC20Wrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20WrapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Wrapper *ERC20WrapperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Wrapper.Contract.ERC20WrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Wrapper *ERC20WrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.ERC20WrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Wrapper *ERC20WrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.ERC20WrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Wrapper *ERC20WrapperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Wrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Wrapper *ERC20WrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Wrapper *ERC20WrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Wrapper *ERC20WrapperCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Wrapper.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Wrapper *ERC20WrapperSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Wrapper.Contract.Allowance(&_ERC20Wrapper.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Wrapper *ERC20WrapperCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Wrapper.Contract.Allowance(&_ERC20Wrapper.CallOpts, owner, spender)
}

// AugurFoundry is a free data retrieval call binding the contract method 0x81ebb30e.
//
// Solidity: function augurFoundry() view returns(address)
func (_ERC20Wrapper *ERC20WrapperCaller) AugurFoundry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Wrapper.contract.Call(opts, out, "augurFoundry")
	return *ret0, err
}

// AugurFoundry is a free data retrieval call binding the contract method 0x81ebb30e.
//
// Solidity: function augurFoundry() view returns(address)
func (_ERC20Wrapper *ERC20WrapperSession) AugurFoundry() (common.Address, error) {
	return _ERC20Wrapper.Contract.AugurFoundry(&_ERC20Wrapper.CallOpts)
}

// AugurFoundry is a free data retrieval call binding the contract method 0x81ebb30e.
//
// Solidity: function augurFoundry() view returns(address)
func (_ERC20Wrapper *ERC20WrapperCallerSession) AugurFoundry() (common.Address, error) {
	return _ERC20Wrapper.Contract.AugurFoundry(&_ERC20Wrapper.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Wrapper *ERC20WrapperCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Wrapper.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Wrapper *ERC20WrapperSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Wrapper.Contract.BalanceOf(&_ERC20Wrapper.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Wrapper *ERC20WrapperCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Wrapper.Contract.BalanceOf(&_ERC20Wrapper.CallOpts, account)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_ERC20Wrapper *ERC20WrapperCaller) Cash(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Wrapper.contract.Call(opts, out, "cash")
	return *ret0, err
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_ERC20Wrapper *ERC20WrapperSession) Cash() (common.Address, error) {
	return _ERC20Wrapper.Contract.Cash(&_ERC20Wrapper.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(address)
func (_ERC20Wrapper *ERC20WrapperCallerSession) Cash() (common.Address, error) {
	return _ERC20Wrapper.Contract.Cash(&_ERC20Wrapper.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Wrapper *ERC20WrapperCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ERC20Wrapper.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Wrapper *ERC20WrapperSession) Decimals() (uint8, error) {
	return _ERC20Wrapper.Contract.Decimals(&_ERC20Wrapper.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Wrapper *ERC20WrapperCallerSession) Decimals() (uint8, error) {
	return _ERC20Wrapper.Contract.Decimals(&_ERC20Wrapper.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Wrapper *ERC20WrapperCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20Wrapper.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Wrapper *ERC20WrapperSession) Name() (string, error) {
	return _ERC20Wrapper.Contract.Name(&_ERC20Wrapper.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Wrapper *ERC20WrapperCallerSession) Name() (string, error) {
	return _ERC20Wrapper.Contract.Name(&_ERC20Wrapper.CallOpts)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_ERC20Wrapper *ERC20WrapperCaller) ShareToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Wrapper.contract.Call(opts, out, "shareToken")
	return *ret0, err
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_ERC20Wrapper *ERC20WrapperSession) ShareToken() (common.Address, error) {
	return _ERC20Wrapper.Contract.ShareToken(&_ERC20Wrapper.CallOpts)
}

// ShareToken is a free data retrieval call binding the contract method 0x6c9fa59e.
//
// Solidity: function shareToken() view returns(address)
func (_ERC20Wrapper *ERC20WrapperCallerSession) ShareToken() (common.Address, error) {
	return _ERC20Wrapper.Contract.ShareToken(&_ERC20Wrapper.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC20Wrapper *ERC20WrapperCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC20Wrapper.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC20Wrapper *ERC20WrapperSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC20Wrapper.Contract.SupportsInterface(&_ERC20Wrapper.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC20Wrapper *ERC20WrapperCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC20Wrapper.Contract.SupportsInterface(&_ERC20Wrapper.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Wrapper *ERC20WrapperCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20Wrapper.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Wrapper *ERC20WrapperSession) Symbol() (string, error) {
	return _ERC20Wrapper.Contract.Symbol(&_ERC20Wrapper.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Wrapper *ERC20WrapperCallerSession) Symbol() (string, error) {
	return _ERC20Wrapper.Contract.Symbol(&_ERC20Wrapper.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_ERC20Wrapper *ERC20WrapperCaller) TokenId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Wrapper.contract.Call(opts, out, "tokenId")
	return *ret0, err
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_ERC20Wrapper *ERC20WrapperSession) TokenId() (*big.Int, error) {
	return _ERC20Wrapper.Contract.TokenId(&_ERC20Wrapper.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_ERC20Wrapper *ERC20WrapperCallerSession) TokenId() (*big.Int, error) {
	return _ERC20Wrapper.Contract.TokenId(&_ERC20Wrapper.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Wrapper *ERC20WrapperCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Wrapper.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Wrapper *ERC20WrapperSession) TotalSupply() (*big.Int, error) {
	return _ERC20Wrapper.Contract.TotalSupply(&_ERC20Wrapper.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Wrapper *ERC20WrapperCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Wrapper.Contract.TotalSupply(&_ERC20Wrapper.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Wrapper *ERC20WrapperTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Wrapper *ERC20WrapperSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.Approve(&_ERC20Wrapper.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Wrapper *ERC20WrapperTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.Approve(&_ERC20Wrapper.TransactOpts, spender, amount)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _account) returns()
func (_ERC20Wrapper *ERC20WrapperTransactor) Claim(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _ERC20Wrapper.contract.Transact(opts, "claim", _account)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _account) returns()
func (_ERC20Wrapper *ERC20WrapperSession) Claim(_account common.Address) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.Claim(&_ERC20Wrapper.TransactOpts, _account)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _account) returns()
func (_ERC20Wrapper *ERC20WrapperTransactorSession) Claim(_account common.Address) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.Claim(&_ERC20Wrapper.TransactOpts, _account)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Wrapper *ERC20WrapperTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Wrapper *ERC20WrapperSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.DecreaseAllowance(&_ERC20Wrapper.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Wrapper *ERC20WrapperTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.DecreaseAllowance(&_ERC20Wrapper.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Wrapper *ERC20WrapperTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Wrapper *ERC20WrapperSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.IncreaseAllowance(&_ERC20Wrapper.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Wrapper *ERC20WrapperTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.IncreaseAllowance(&_ERC20Wrapper.TransactOpts, spender, addedValue)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_ERC20Wrapper *ERC20WrapperTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20Wrapper.contract.Transact(opts, "onERC1155BatchReceived", operator, from, ids, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_ERC20Wrapper *ERC20WrapperSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.OnERC1155BatchReceived(&_ERC20Wrapper.TransactOpts, operator, from, ids, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_ERC20Wrapper *ERC20WrapperTransactorSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.OnERC1155BatchReceived(&_ERC20Wrapper.TransactOpts, operator, from, ids, values, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_ERC20Wrapper *ERC20WrapperTransactor) OnERC1155Received(opts *bind.TransactOpts, operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20Wrapper.contract.Transact(opts, "onERC1155Received", operator, from, id, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_ERC20Wrapper *ERC20WrapperSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.OnERC1155Received(&_ERC20Wrapper.TransactOpts, operator, from, id, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_ERC20Wrapper *ERC20WrapperTransactorSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.OnERC1155Received(&_ERC20Wrapper.TransactOpts, operator, from, id, value, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Wrapper *ERC20WrapperTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Wrapper *ERC20WrapperSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.Transfer(&_ERC20Wrapper.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Wrapper *ERC20WrapperTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.Transfer(&_ERC20Wrapper.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Wrapper *ERC20WrapperTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Wrapper *ERC20WrapperSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.TransferFrom(&_ERC20Wrapper.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Wrapper *ERC20WrapperTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.TransferFrom(&_ERC20Wrapper.TransactOpts, sender, recipient, amount)
}

// UnWrapTokens is a paid mutator transaction binding the contract method 0xfbfdc41c.
//
// Solidity: function unWrapTokens(address _account, uint256 _amount) returns()
func (_ERC20Wrapper *ERC20WrapperTransactor) UnWrapTokens(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.contract.Transact(opts, "unWrapTokens", _account, _amount)
}

// UnWrapTokens is a paid mutator transaction binding the contract method 0xfbfdc41c.
//
// Solidity: function unWrapTokens(address _account, uint256 _amount) returns()
func (_ERC20Wrapper *ERC20WrapperSession) UnWrapTokens(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.UnWrapTokens(&_ERC20Wrapper.TransactOpts, _account, _amount)
}

// UnWrapTokens is a paid mutator transaction binding the contract method 0xfbfdc41c.
//
// Solidity: function unWrapTokens(address _account, uint256 _amount) returns()
func (_ERC20Wrapper *ERC20WrapperTransactorSession) UnWrapTokens(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.UnWrapTokens(&_ERC20Wrapper.TransactOpts, _account, _amount)
}

// WrapTokens is a paid mutator transaction binding the contract method 0xa084bccf.
//
// Solidity: function wrapTokens(address _account, uint256 _amount) returns()
func (_ERC20Wrapper *ERC20WrapperTransactor) WrapTokens(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.contract.Transact(opts, "wrapTokens", _account, _amount)
}

// WrapTokens is a paid mutator transaction binding the contract method 0xa084bccf.
//
// Solidity: function wrapTokens(address _account, uint256 _amount) returns()
func (_ERC20Wrapper *ERC20WrapperSession) WrapTokens(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.WrapTokens(&_ERC20Wrapper.TransactOpts, _account, _amount)
}

// WrapTokens is a paid mutator transaction binding the contract method 0xa084bccf.
//
// Solidity: function wrapTokens(address _account, uint256 _amount) returns()
func (_ERC20Wrapper *ERC20WrapperTransactorSession) WrapTokens(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC20Wrapper.Contract.WrapTokens(&_ERC20Wrapper.TransactOpts, _account, _amount)
}

// ERC20WrapperApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Wrapper contract.
type ERC20WrapperApprovalIterator struct {
	Event *ERC20WrapperApproval // Event containing the contract specifics and raw log

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
func (it *ERC20WrapperApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20WrapperApproval)
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
		it.Event = new(ERC20WrapperApproval)
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
func (it *ERC20WrapperApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20WrapperApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20WrapperApproval represents a Approval event raised by the ERC20Wrapper contract.
type ERC20WrapperApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Wrapper *ERC20WrapperFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20WrapperApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Wrapper.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20WrapperApprovalIterator{contract: _ERC20Wrapper.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Wrapper *ERC20WrapperFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20WrapperApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Wrapper.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20WrapperApproval)
				if err := _ERC20Wrapper.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Wrapper *ERC20WrapperFilterer) ParseApproval(log types.Log) (*ERC20WrapperApproval, error) {
	event := new(ERC20WrapperApproval)
	if err := _ERC20Wrapper.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20WrapperTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Wrapper contract.
type ERC20WrapperTransferIterator struct {
	Event *ERC20WrapperTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20WrapperTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20WrapperTransfer)
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
		it.Event = new(ERC20WrapperTransfer)
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
func (it *ERC20WrapperTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20WrapperTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20WrapperTransfer represents a Transfer event raised by the ERC20Wrapper contract.
type ERC20WrapperTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Wrapper *ERC20WrapperFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20WrapperTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Wrapper.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20WrapperTransferIterator{contract: _ERC20Wrapper.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Wrapper *ERC20WrapperFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20WrapperTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Wrapper.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20WrapperTransfer)
				if err := _ERC20Wrapper.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Wrapper *ERC20WrapperFilterer) ParseTransfer(log types.Log) (*ERC20WrapperTransfer, error) {
	event := new(ERC20WrapperTransfer)
	if err := _ERC20Wrapper.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
