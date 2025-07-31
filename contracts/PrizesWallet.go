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

// IPrizesWalletDonatedTokenToClaim is an auto generated low-level Go binding around an user-defined struct.
type IPrizesWalletDonatedTokenToClaim struct {
	RoundNum     *big.Int
	TokenAddress common.Address
	Amount       *big.Int
}

// IPrizesWalletEthBalanceInfo is an auto generated low-level Go binding around an user-defined struct.
type IPrizesWalletEthBalanceInfo struct {
	RoundNum *big.Int
	Amount   *big.Int
}

// IPrizesWalletEthDeposit is an auto generated low-level Go binding around an user-defined struct.
type IPrizesWalletEthDeposit struct {
	PrizeWinnerAddress common.Address
	Amount             *big.Int
}

// IPrizesWalletMetaData contains all meta data concerning the IPrizesWallet contract.
var IPrizesWalletMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DonatedNftClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonatedTokenClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prizeWinnerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prizeWinnerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"NftDonated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"TimeoutDurationToWithdrawPrizesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenDonated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"claimDonatedNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"claimDonatedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"indexes_\",\"type\":\"uint256[]\"}],\"name\":\"claimManyDonatedNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.DonatedTokenToClaim[]\",\"name\":\"donatedTokensToClaim_\",\"type\":\"tuple[]\"}],\"name\":\"claimManyDonatedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"prizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress_\",\"type\":\"address\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"donateNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress_\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"donateToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"}],\"name\":\"getDonatedTokenBalanceAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEthBalanceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.EthBalanceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"getEthBalanceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.EthBalanceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mainPrizeBeneficiaryAddress_\",\"type\":\"address\"}],\"name\":\"registerRoundEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mainPrizeBeneficiaryAddress_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"prizeWinnerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.EthDeposit[]\",\"name\":\"ethDeposits_\",\"type\":\"tuple[]\"}],\"name\":\"registerRoundEndAndDepositEthMany\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setTimeoutDurationToWithdrawPrizes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"withdrawEth_\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.DonatedTokenToClaim[]\",\"name\":\"donatedTokensToClaim_\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"donatedNftIndexes_\",\"type\":\"uint256[]\"}],\"name\":\"withdrawEverything\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPrizesWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use IPrizesWalletMetaData.ABI instead.
var IPrizesWalletABI = IPrizesWalletMetaData.ABI

// IPrizesWallet is an auto generated Go binding around an Ethereum contract.
type IPrizesWallet struct {
	IPrizesWalletCaller     // Read-only binding to the contract
	IPrizesWalletTransactor // Write-only binding to the contract
	IPrizesWalletFilterer   // Log filterer for contract events
}

// IPrizesWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPrizesWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPrizesWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPrizesWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPrizesWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPrizesWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPrizesWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPrizesWalletSession struct {
	Contract     *IPrizesWallet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPrizesWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPrizesWalletCallerSession struct {
	Contract *IPrizesWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IPrizesWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPrizesWalletTransactorSession struct {
	Contract     *IPrizesWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IPrizesWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPrizesWalletRaw struct {
	Contract *IPrizesWallet // Generic contract binding to access the raw methods on
}

// IPrizesWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPrizesWalletCallerRaw struct {
	Contract *IPrizesWalletCaller // Generic read-only contract binding to access the raw methods on
}

// IPrizesWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPrizesWalletTransactorRaw struct {
	Contract *IPrizesWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPrizesWallet creates a new instance of IPrizesWallet, bound to a specific deployed contract.
func NewIPrizesWallet(address common.Address, backend bind.ContractBackend) (*IPrizesWallet, error) {
	contract, err := bindIPrizesWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPrizesWallet{IPrizesWalletCaller: IPrizesWalletCaller{contract: contract}, IPrizesWalletTransactor: IPrizesWalletTransactor{contract: contract}, IPrizesWalletFilterer: IPrizesWalletFilterer{contract: contract}}, nil
}

// NewIPrizesWalletCaller creates a new read-only instance of IPrizesWallet, bound to a specific deployed contract.
func NewIPrizesWalletCaller(address common.Address, caller bind.ContractCaller) (*IPrizesWalletCaller, error) {
	contract, err := bindIPrizesWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletCaller{contract: contract}, nil
}

// NewIPrizesWalletTransactor creates a new write-only instance of IPrizesWallet, bound to a specific deployed contract.
func NewIPrizesWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*IPrizesWalletTransactor, error) {
	contract, err := bindIPrizesWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletTransactor{contract: contract}, nil
}

// NewIPrizesWalletFilterer creates a new log filterer instance of IPrizesWallet, bound to a specific deployed contract.
func NewIPrizesWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*IPrizesWalletFilterer, error) {
	contract, err := bindIPrizesWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletFilterer{contract: contract}, nil
}

// bindIPrizesWallet binds a generic wrapper to an already deployed contract.
func bindIPrizesWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPrizesWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPrizesWallet *IPrizesWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPrizesWallet.Contract.IPrizesWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPrizesWallet *IPrizesWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.IPrizesWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPrizesWallet *IPrizesWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.IPrizesWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPrizesWallet *IPrizesWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPrizesWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPrizesWallet *IPrizesWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPrizesWallet *IPrizesWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.contract.Transact(opts, method, params...)
}

// GetDonatedTokenBalanceAmount is a free data retrieval call binding the contract method 0xe86a49d7.
//
// Solidity: function getDonatedTokenBalanceAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_IPrizesWallet *IPrizesWalletCaller) GetDonatedTokenBalanceAmount(opts *bind.CallOpts, roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPrizesWallet.contract.Call(opts, &out, "getDonatedTokenBalanceAmount", roundNum_, tokenAddress_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDonatedTokenBalanceAmount is a free data retrieval call binding the contract method 0xe86a49d7.
//
// Solidity: function getDonatedTokenBalanceAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_IPrizesWallet *IPrizesWalletSession) GetDonatedTokenBalanceAmount(roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	return _IPrizesWallet.Contract.GetDonatedTokenBalanceAmount(&_IPrizesWallet.CallOpts, roundNum_, tokenAddress_)
}

// GetDonatedTokenBalanceAmount is a free data retrieval call binding the contract method 0xe86a49d7.
//
// Solidity: function getDonatedTokenBalanceAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_IPrizesWallet *IPrizesWalletCallerSession) GetDonatedTokenBalanceAmount(roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	return _IPrizesWallet.Contract.GetDonatedTokenBalanceAmount(&_IPrizesWallet.CallOpts, roundNum_, tokenAddress_)
}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletCaller) GetEthBalanceInfo(opts *bind.CallOpts) (IPrizesWalletEthBalanceInfo, error) {
	var out []interface{}
	err := _IPrizesWallet.contract.Call(opts, &out, "getEthBalanceInfo")

	if err != nil {
		return *new(IPrizesWalletEthBalanceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPrizesWalletEthBalanceInfo)).(*IPrizesWalletEthBalanceInfo)

	return out0, err

}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletSession) GetEthBalanceInfo() (IPrizesWalletEthBalanceInfo, error) {
	return _IPrizesWallet.Contract.GetEthBalanceInfo(&_IPrizesWallet.CallOpts)
}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletCallerSession) GetEthBalanceInfo() (IPrizesWalletEthBalanceInfo, error) {
	return _IPrizesWallet.Contract.GetEthBalanceInfo(&_IPrizesWallet.CallOpts)
}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address prizeWinnerAddress_) view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletCaller) GetEthBalanceInfo0(opts *bind.CallOpts, prizeWinnerAddress_ common.Address) (IPrizesWalletEthBalanceInfo, error) {
	var out []interface{}
	err := _IPrizesWallet.contract.Call(opts, &out, "getEthBalanceInfo0", prizeWinnerAddress_)

	if err != nil {
		return *new(IPrizesWalletEthBalanceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPrizesWalletEthBalanceInfo)).(*IPrizesWalletEthBalanceInfo)

	return out0, err

}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address prizeWinnerAddress_) view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletSession) GetEthBalanceInfo0(prizeWinnerAddress_ common.Address) (IPrizesWalletEthBalanceInfo, error) {
	return _IPrizesWallet.Contract.GetEthBalanceInfo0(&_IPrizesWallet.CallOpts, prizeWinnerAddress_)
}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address prizeWinnerAddress_) view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletCallerSession) GetEthBalanceInfo0(prizeWinnerAddress_ common.Address) (IPrizesWalletEthBalanceInfo, error) {
	return _IPrizesWallet.Contract.GetEthBalanceInfo0(&_IPrizesWallet.CallOpts, prizeWinnerAddress_)
}

// ClaimDonatedNft is a paid mutator transaction binding the contract method 0x94d907fc.
//
// Solidity: function claimDonatedNft(uint256 index_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) ClaimDonatedNft(opts *bind.TransactOpts, index_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "claimDonatedNft", index_)
}

// ClaimDonatedNft is a paid mutator transaction binding the contract method 0x94d907fc.
//
// Solidity: function claimDonatedNft(uint256 index_) returns()
func (_IPrizesWallet *IPrizesWalletSession) ClaimDonatedNft(index_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimDonatedNft(&_IPrizesWallet.TransactOpts, index_)
}

// ClaimDonatedNft is a paid mutator transaction binding the contract method 0x94d907fc.
//
// Solidity: function claimDonatedNft(uint256 index_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) ClaimDonatedNft(index_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimDonatedNft(&_IPrizesWallet.TransactOpts, index_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x9523b237.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_, uint256 amount_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) ClaimDonatedToken(opts *bind.TransactOpts, roundNum_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "claimDonatedToken", roundNum_, tokenAddress_, amount_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x9523b237.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_, uint256 amount_) returns()
func (_IPrizesWallet *IPrizesWalletSession) ClaimDonatedToken(roundNum_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimDonatedToken(&_IPrizesWallet.TransactOpts, roundNum_, tokenAddress_, amount_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x9523b237.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_, uint256 amount_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) ClaimDonatedToken(roundNum_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimDonatedToken(&_IPrizesWallet.TransactOpts, roundNum_, tokenAddress_, amount_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indexes_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) ClaimManyDonatedNfts(opts *bind.TransactOpts, indexes_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "claimManyDonatedNfts", indexes_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indexes_) returns()
func (_IPrizesWallet *IPrizesWalletSession) ClaimManyDonatedNfts(indexes_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimManyDonatedNfts(&_IPrizesWallet.TransactOpts, indexes_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indexes_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) ClaimManyDonatedNfts(indexes_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimManyDonatedNfts(&_IPrizesWallet.TransactOpts, indexes_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x5e00aaba.
//
// Solidity: function claimManyDonatedTokens((uint256,address,uint256)[] donatedTokensToClaim_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) ClaimManyDonatedTokens(opts *bind.TransactOpts, donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "claimManyDonatedTokens", donatedTokensToClaim_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x5e00aaba.
//
// Solidity: function claimManyDonatedTokens((uint256,address,uint256)[] donatedTokensToClaim_) returns()
func (_IPrizesWallet *IPrizesWalletSession) ClaimManyDonatedTokens(donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimManyDonatedTokens(&_IPrizesWallet.TransactOpts, donatedTokensToClaim_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x5e00aaba.
//
// Solidity: function claimManyDonatedTokens((uint256,address,uint256)[] donatedTokensToClaim_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) ClaimManyDonatedTokens(donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimManyDonatedTokens(&_IPrizesWallet.TransactOpts, donatedTokensToClaim_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x46f6b4e1.
//
// Solidity: function depositEth(uint256 roundNum_, address prizeWinnerAddress_) payable returns()
func (_IPrizesWallet *IPrizesWalletTransactor) DepositEth(opts *bind.TransactOpts, roundNum_ *big.Int, prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "depositEth", roundNum_, prizeWinnerAddress_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x46f6b4e1.
//
// Solidity: function depositEth(uint256 roundNum_, address prizeWinnerAddress_) payable returns()
func (_IPrizesWallet *IPrizesWalletSession) DepositEth(roundNum_ *big.Int, prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.DepositEth(&_IPrizesWallet.TransactOpts, roundNum_, prizeWinnerAddress_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x46f6b4e1.
//
// Solidity: function depositEth(uint256 roundNum_, address prizeWinnerAddress_) payable returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) DepositEth(roundNum_ *big.Int, prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.DepositEth(&_IPrizesWallet.TransactOpts, roundNum_, prizeWinnerAddress_)
}

// DonateNft is a paid mutator transaction binding the contract method 0xfe673fd3.
//
// Solidity: function donateNft(uint256 roundNum_, address donorAddress_, address nftAddress_, uint256 nftId_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) DonateNft(opts *bind.TransactOpts, roundNum_ *big.Int, donorAddress_ common.Address, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "donateNft", roundNum_, donorAddress_, nftAddress_, nftId_)
}

// DonateNft is a paid mutator transaction binding the contract method 0xfe673fd3.
//
// Solidity: function donateNft(uint256 roundNum_, address donorAddress_, address nftAddress_, uint256 nftId_) returns()
func (_IPrizesWallet *IPrizesWalletSession) DonateNft(roundNum_ *big.Int, donorAddress_ common.Address, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.DonateNft(&_IPrizesWallet.TransactOpts, roundNum_, donorAddress_, nftAddress_, nftId_)
}

// DonateNft is a paid mutator transaction binding the contract method 0xfe673fd3.
//
// Solidity: function donateNft(uint256 roundNum_, address donorAddress_, address nftAddress_, uint256 nftId_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) DonateNft(roundNum_ *big.Int, donorAddress_ common.Address, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.DonateNft(&_IPrizesWallet.TransactOpts, roundNum_, donorAddress_, nftAddress_, nftId_)
}

// DonateToken is a paid mutator transaction binding the contract method 0xe2051c7e.
//
// Solidity: function donateToken(uint256 roundNum_, address donorAddress_, address tokenAddress_, uint256 amount_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) DonateToken(opts *bind.TransactOpts, roundNum_ *big.Int, donorAddress_ common.Address, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "donateToken", roundNum_, donorAddress_, tokenAddress_, amount_)
}

// DonateToken is a paid mutator transaction binding the contract method 0xe2051c7e.
//
// Solidity: function donateToken(uint256 roundNum_, address donorAddress_, address tokenAddress_, uint256 amount_) returns()
func (_IPrizesWallet *IPrizesWalletSession) DonateToken(roundNum_ *big.Int, donorAddress_ common.Address, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.DonateToken(&_IPrizesWallet.TransactOpts, roundNum_, donorAddress_, tokenAddress_, amount_)
}

// DonateToken is a paid mutator transaction binding the contract method 0xe2051c7e.
//
// Solidity: function donateToken(uint256 roundNum_, address donorAddress_, address tokenAddress_, uint256 amount_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) DonateToken(roundNum_ *big.Int, donorAddress_ common.Address, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.DonateToken(&_IPrizesWallet.TransactOpts, roundNum_, donorAddress_, tokenAddress_, amount_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address mainPrizeBeneficiaryAddress_) returns(uint256)
func (_IPrizesWallet *IPrizesWalletTransactor) RegisterRoundEnd(opts *bind.TransactOpts, roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "registerRoundEnd", roundNum_, mainPrizeBeneficiaryAddress_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address mainPrizeBeneficiaryAddress_) returns(uint256)
func (_IPrizesWallet *IPrizesWalletSession) RegisterRoundEnd(roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.RegisterRoundEnd(&_IPrizesWallet.TransactOpts, roundNum_, mainPrizeBeneficiaryAddress_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address mainPrizeBeneficiaryAddress_) returns(uint256)
func (_IPrizesWallet *IPrizesWalletTransactorSession) RegisterRoundEnd(roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.RegisterRoundEnd(&_IPrizesWallet.TransactOpts, roundNum_, mainPrizeBeneficiaryAddress_)
}

// RegisterRoundEndAndDepositEthMany is a paid mutator transaction binding the contract method 0x87565d14.
//
// Solidity: function registerRoundEndAndDepositEthMany(uint256 roundNum_, address mainPrizeBeneficiaryAddress_, (address,uint256)[] ethDeposits_) payable returns(uint256)
func (_IPrizesWallet *IPrizesWalletTransactor) RegisterRoundEndAndDepositEthMany(opts *bind.TransactOpts, roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address, ethDeposits_ []IPrizesWalletEthDeposit) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "registerRoundEndAndDepositEthMany", roundNum_, mainPrizeBeneficiaryAddress_, ethDeposits_)
}

// RegisterRoundEndAndDepositEthMany is a paid mutator transaction binding the contract method 0x87565d14.
//
// Solidity: function registerRoundEndAndDepositEthMany(uint256 roundNum_, address mainPrizeBeneficiaryAddress_, (address,uint256)[] ethDeposits_) payable returns(uint256)
func (_IPrizesWallet *IPrizesWalletSession) RegisterRoundEndAndDepositEthMany(roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address, ethDeposits_ []IPrizesWalletEthDeposit) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.RegisterRoundEndAndDepositEthMany(&_IPrizesWallet.TransactOpts, roundNum_, mainPrizeBeneficiaryAddress_, ethDeposits_)
}

// RegisterRoundEndAndDepositEthMany is a paid mutator transaction binding the contract method 0x87565d14.
//
// Solidity: function registerRoundEndAndDepositEthMany(uint256 roundNum_, address mainPrizeBeneficiaryAddress_, (address,uint256)[] ethDeposits_) payable returns(uint256)
func (_IPrizesWallet *IPrizesWalletTransactorSession) RegisterRoundEndAndDepositEthMany(roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address, ethDeposits_ []IPrizesWalletEthDeposit) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.RegisterRoundEndAndDepositEthMany(&_IPrizesWallet.TransactOpts, roundNum_, mainPrizeBeneficiaryAddress_, ethDeposits_)
}

// SetTimeoutDurationToWithdrawPrizes is a paid mutator transaction binding the contract method 0x9e2842a8.
//
// Solidity: function setTimeoutDurationToWithdrawPrizes(uint256 newValue_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) SetTimeoutDurationToWithdrawPrizes(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "setTimeoutDurationToWithdrawPrizes", newValue_)
}

// SetTimeoutDurationToWithdrawPrizes is a paid mutator transaction binding the contract method 0x9e2842a8.
//
// Solidity: function setTimeoutDurationToWithdrawPrizes(uint256 newValue_) returns()
func (_IPrizesWallet *IPrizesWalletSession) SetTimeoutDurationToWithdrawPrizes(newValue_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.SetTimeoutDurationToWithdrawPrizes(&_IPrizesWallet.TransactOpts, newValue_)
}

// SetTimeoutDurationToWithdrawPrizes is a paid mutator transaction binding the contract method 0x9e2842a8.
//
// Solidity: function setTimeoutDurationToWithdrawPrizes(uint256 newValue_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) SetTimeoutDurationToWithdrawPrizes(newValue_ *big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.SetTimeoutDurationToWithdrawPrizes(&_IPrizesWallet.TransactOpts, newValue_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address prizeWinnerAddress_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) WithdrawEth(opts *bind.TransactOpts, prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "withdrawEth", prizeWinnerAddress_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address prizeWinnerAddress_) returns()
func (_IPrizesWallet *IPrizesWalletSession) WithdrawEth(prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEth(&_IPrizesWallet.TransactOpts, prizeWinnerAddress_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address prizeWinnerAddress_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) WithdrawEth(prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEth(&_IPrizesWallet.TransactOpts, prizeWinnerAddress_)
}

// WithdrawEth0 is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_IPrizesWallet *IPrizesWalletTransactor) WithdrawEth0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "withdrawEth0")
}

// WithdrawEth0 is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_IPrizesWallet *IPrizesWalletSession) WithdrawEth0() (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEth0(&_IPrizesWallet.TransactOpts)
}

// WithdrawEth0 is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) WithdrawEth0() (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEth0(&_IPrizesWallet.TransactOpts)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xb4e6f803.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address,uint256)[] donatedTokensToClaim_, uint256[] donatedNftIndexes_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) WithdrawEverything(opts *bind.TransactOpts, withdrawEth_ bool, donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim, donatedNftIndexes_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "withdrawEverything", withdrawEth_, donatedTokensToClaim_, donatedNftIndexes_)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xb4e6f803.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address,uint256)[] donatedTokensToClaim_, uint256[] donatedNftIndexes_) returns()
func (_IPrizesWallet *IPrizesWalletSession) WithdrawEverything(withdrawEth_ bool, donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim, donatedNftIndexes_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEverything(&_IPrizesWallet.TransactOpts, withdrawEth_, donatedTokensToClaim_, donatedNftIndexes_)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xb4e6f803.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address,uint256)[] donatedTokensToClaim_, uint256[] donatedNftIndexes_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) WithdrawEverything(withdrawEth_ bool, donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim, donatedNftIndexes_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEverything(&_IPrizesWallet.TransactOpts, withdrawEth_, donatedTokensToClaim_, donatedNftIndexes_)
}

// IPrizesWalletDonatedNftClaimedIterator is returned from FilterDonatedNftClaimed and is used to iterate over the raw logs and unpacked data for DonatedNftClaimed events raised by the IPrizesWallet contract.
type IPrizesWalletDonatedNftClaimedIterator struct {
	Event *IPrizesWalletDonatedNftClaimed // Event containing the contract specifics and raw log

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
func (it *IPrizesWalletDonatedNftClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPrizesWalletDonatedNftClaimed)
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
		it.Event = new(IPrizesWalletDonatedNftClaimed)
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
func (it *IPrizesWalletDonatedNftClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPrizesWalletDonatedNftClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPrizesWalletDonatedNftClaimed represents a DonatedNftClaimed event raised by the IPrizesWallet contract.
type IPrizesWalletDonatedNftClaimed struct {
	RoundNum           *big.Int
	BeneficiaryAddress common.Address
	NftAddress         common.Address
	NftId              *big.Int
	Index              *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDonatedNftClaimed is a free log retrieval operation binding the contract event 0x03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3.
//
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterDonatedNftClaimed(opts *bind.FilterOpts, roundNum []*big.Int, beneficiaryAddress []common.Address, nftAddress []common.Address) (*IPrizesWalletDonatedNftClaimedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "DonatedNftClaimed", roundNumRule, beneficiaryAddressRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletDonatedNftClaimedIterator{contract: _IPrizesWallet.contract, event: "DonatedNftClaimed", logs: logs, sub: sub}, nil
}

// WatchDonatedNftClaimed is a free log subscription operation binding the contract event 0x03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3.
//
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchDonatedNftClaimed(opts *bind.WatchOpts, sink chan<- *IPrizesWalletDonatedNftClaimed, roundNum []*big.Int, beneficiaryAddress []common.Address, nftAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "DonatedNftClaimed", roundNumRule, beneficiaryAddressRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPrizesWalletDonatedNftClaimed)
				if err := _IPrizesWallet.contract.UnpackLog(event, "DonatedNftClaimed", log); err != nil {
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

// ParseDonatedNftClaimed is a log parse operation binding the contract event 0x03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3.
//
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_IPrizesWallet *IPrizesWalletFilterer) ParseDonatedNftClaimed(log types.Log) (*IPrizesWalletDonatedNftClaimed, error) {
	event := new(IPrizesWalletDonatedNftClaimed)
	if err := _IPrizesWallet.contract.UnpackLog(event, "DonatedNftClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPrizesWalletDonatedTokenClaimedIterator is returned from FilterDonatedTokenClaimed and is used to iterate over the raw logs and unpacked data for DonatedTokenClaimed events raised by the IPrizesWallet contract.
type IPrizesWalletDonatedTokenClaimedIterator struct {
	Event *IPrizesWalletDonatedTokenClaimed // Event containing the contract specifics and raw log

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
func (it *IPrizesWalletDonatedTokenClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPrizesWalletDonatedTokenClaimed)
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
		it.Event = new(IPrizesWalletDonatedTokenClaimed)
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
func (it *IPrizesWalletDonatedTokenClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPrizesWalletDonatedTokenClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPrizesWalletDonatedTokenClaimed represents a DonatedTokenClaimed event raised by the IPrizesWallet contract.
type IPrizesWalletDonatedTokenClaimed struct {
	RoundNum           *big.Int
	BeneficiaryAddress common.Address
	TokenAddress       common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDonatedTokenClaimed is a free log retrieval operation binding the contract event 0xaf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0.
//
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed tokenAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterDonatedTokenClaimed(opts *bind.FilterOpts, roundNum []*big.Int, beneficiaryAddress []common.Address, tokenAddress []common.Address) (*IPrizesWalletDonatedTokenClaimedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "DonatedTokenClaimed", roundNumRule, beneficiaryAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletDonatedTokenClaimedIterator{contract: _IPrizesWallet.contract, event: "DonatedTokenClaimed", logs: logs, sub: sub}, nil
}

// WatchDonatedTokenClaimed is a free log subscription operation binding the contract event 0xaf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0.
//
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed tokenAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchDonatedTokenClaimed(opts *bind.WatchOpts, sink chan<- *IPrizesWalletDonatedTokenClaimed, roundNum []*big.Int, beneficiaryAddress []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "DonatedTokenClaimed", roundNumRule, beneficiaryAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPrizesWalletDonatedTokenClaimed)
				if err := _IPrizesWallet.contract.UnpackLog(event, "DonatedTokenClaimed", log); err != nil {
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

// ParseDonatedTokenClaimed is a log parse operation binding the contract event 0xaf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0.
//
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed tokenAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) ParseDonatedTokenClaimed(log types.Log) (*IPrizesWalletDonatedTokenClaimed, error) {
	event := new(IPrizesWalletDonatedTokenClaimed)
	if err := _IPrizesWallet.contract.UnpackLog(event, "DonatedTokenClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPrizesWalletEthReceivedIterator is returned from FilterEthReceived and is used to iterate over the raw logs and unpacked data for EthReceived events raised by the IPrizesWallet contract.
type IPrizesWalletEthReceivedIterator struct {
	Event *IPrizesWalletEthReceived // Event containing the contract specifics and raw log

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
func (it *IPrizesWalletEthReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPrizesWalletEthReceived)
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
		it.Event = new(IPrizesWalletEthReceived)
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
func (it *IPrizesWalletEthReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPrizesWalletEthReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPrizesWalletEthReceived represents a EthReceived event raised by the IPrizesWallet contract.
type IPrizesWalletEthReceived struct {
	RoundNum           *big.Int
	PrizeWinnerAddress common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterEthReceived is a free log retrieval operation binding the contract event 0x999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326.
//
// Solidity: event EthReceived(uint256 indexed roundNum, address indexed prizeWinnerAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterEthReceived(opts *bind.FilterOpts, roundNum []*big.Int, prizeWinnerAddress []common.Address) (*IPrizesWalletEthReceivedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var prizeWinnerAddressRule []interface{}
	for _, prizeWinnerAddressItem := range prizeWinnerAddress {
		prizeWinnerAddressRule = append(prizeWinnerAddressRule, prizeWinnerAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "EthReceived", roundNumRule, prizeWinnerAddressRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletEthReceivedIterator{contract: _IPrizesWallet.contract, event: "EthReceived", logs: logs, sub: sub}, nil
}

// WatchEthReceived is a free log subscription operation binding the contract event 0x999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326.
//
// Solidity: event EthReceived(uint256 indexed roundNum, address indexed prizeWinnerAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchEthReceived(opts *bind.WatchOpts, sink chan<- *IPrizesWalletEthReceived, roundNum []*big.Int, prizeWinnerAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var prizeWinnerAddressRule []interface{}
	for _, prizeWinnerAddressItem := range prizeWinnerAddress {
		prizeWinnerAddressRule = append(prizeWinnerAddressRule, prizeWinnerAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "EthReceived", roundNumRule, prizeWinnerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPrizesWalletEthReceived)
				if err := _IPrizesWallet.contract.UnpackLog(event, "EthReceived", log); err != nil {
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

// ParseEthReceived is a log parse operation binding the contract event 0x999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326.
//
// Solidity: event EthReceived(uint256 indexed roundNum, address indexed prizeWinnerAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) ParseEthReceived(log types.Log) (*IPrizesWalletEthReceived, error) {
	event := new(IPrizesWalletEthReceived)
	if err := _IPrizesWallet.contract.UnpackLog(event, "EthReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPrizesWalletEthWithdrawnIterator is returned from FilterEthWithdrawn and is used to iterate over the raw logs and unpacked data for EthWithdrawn events raised by the IPrizesWallet contract.
type IPrizesWalletEthWithdrawnIterator struct {
	Event *IPrizesWalletEthWithdrawn // Event containing the contract specifics and raw log

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
func (it *IPrizesWalletEthWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPrizesWalletEthWithdrawn)
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
		it.Event = new(IPrizesWalletEthWithdrawn)
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
func (it *IPrizesWalletEthWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPrizesWalletEthWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPrizesWalletEthWithdrawn represents a EthWithdrawn event raised by the IPrizesWallet contract.
type IPrizesWalletEthWithdrawn struct {
	PrizeWinnerAddress common.Address
	BeneficiaryAddress common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawn is a free log retrieval operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed prizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterEthWithdrawn(opts *bind.FilterOpts, prizeWinnerAddress []common.Address, beneficiaryAddress []common.Address) (*IPrizesWalletEthWithdrawnIterator, error) {

	var prizeWinnerAddressRule []interface{}
	for _, prizeWinnerAddressItem := range prizeWinnerAddress {
		prizeWinnerAddressRule = append(prizeWinnerAddressRule, prizeWinnerAddressItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "EthWithdrawn", prizeWinnerAddressRule, beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletEthWithdrawnIterator{contract: _IPrizesWallet.contract, event: "EthWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawn is a free log subscription operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed prizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchEthWithdrawn(opts *bind.WatchOpts, sink chan<- *IPrizesWalletEthWithdrawn, prizeWinnerAddress []common.Address, beneficiaryAddress []common.Address) (event.Subscription, error) {

	var prizeWinnerAddressRule []interface{}
	for _, prizeWinnerAddressItem := range prizeWinnerAddress {
		prizeWinnerAddressRule = append(prizeWinnerAddressRule, prizeWinnerAddressItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "EthWithdrawn", prizeWinnerAddressRule, beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPrizesWalletEthWithdrawn)
				if err := _IPrizesWallet.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
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

// ParseEthWithdrawn is a log parse operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed prizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) ParseEthWithdrawn(log types.Log) (*IPrizesWalletEthWithdrawn, error) {
	event := new(IPrizesWalletEthWithdrawn)
	if err := _IPrizesWallet.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPrizesWalletNftDonatedIterator is returned from FilterNftDonated and is used to iterate over the raw logs and unpacked data for NftDonated events raised by the IPrizesWallet contract.
type IPrizesWalletNftDonatedIterator struct {
	Event *IPrizesWalletNftDonated // Event containing the contract specifics and raw log

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
func (it *IPrizesWalletNftDonatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPrizesWalletNftDonated)
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
		it.Event = new(IPrizesWalletNftDonated)
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
func (it *IPrizesWalletNftDonatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPrizesWalletNftDonatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPrizesWalletNftDonated represents a NftDonated event raised by the IPrizesWallet contract.
type IPrizesWalletNftDonated struct {
	RoundNum     *big.Int
	DonorAddress common.Address
	NftAddress   common.Address
	NftId        *big.Int
	Index        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNftDonated is a free log retrieval operation binding the contract event 0xb12e72bab0c2a8fe98175a3c02792645ebdf65d37cf3983517e6d1c5ab842e23.
//
// Solidity: event NftDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterNftDonated(opts *bind.FilterOpts, roundNum []*big.Int, donorAddress []common.Address, nftAddress []common.Address) (*IPrizesWalletNftDonatedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "NftDonated", roundNumRule, donorAddressRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletNftDonatedIterator{contract: _IPrizesWallet.contract, event: "NftDonated", logs: logs, sub: sub}, nil
}

// WatchNftDonated is a free log subscription operation binding the contract event 0xb12e72bab0c2a8fe98175a3c02792645ebdf65d37cf3983517e6d1c5ab842e23.
//
// Solidity: event NftDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchNftDonated(opts *bind.WatchOpts, sink chan<- *IPrizesWalletNftDonated, roundNum []*big.Int, donorAddress []common.Address, nftAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "NftDonated", roundNumRule, donorAddressRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPrizesWalletNftDonated)
				if err := _IPrizesWallet.contract.UnpackLog(event, "NftDonated", log); err != nil {
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

// ParseNftDonated is a log parse operation binding the contract event 0xb12e72bab0c2a8fe98175a3c02792645ebdf65d37cf3983517e6d1c5ab842e23.
//
// Solidity: event NftDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_IPrizesWallet *IPrizesWalletFilterer) ParseNftDonated(log types.Log) (*IPrizesWalletNftDonated, error) {
	event := new(IPrizesWalletNftDonated)
	if err := _IPrizesWallet.contract.UnpackLog(event, "NftDonated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator is returned from FilterTimeoutDurationToWithdrawPrizesChanged and is used to iterate over the raw logs and unpacked data for TimeoutDurationToWithdrawPrizesChanged events raised by the IPrizesWallet contract.
type IPrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator struct {
	Event *IPrizesWalletTimeoutDurationToWithdrawPrizesChanged // Event containing the contract specifics and raw log

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
func (it *IPrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPrizesWalletTimeoutDurationToWithdrawPrizesChanged)
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
		it.Event = new(IPrizesWalletTimeoutDurationToWithdrawPrizesChanged)
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
func (it *IPrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPrizesWalletTimeoutDurationToWithdrawPrizesChanged represents a TimeoutDurationToWithdrawPrizesChanged event raised by the IPrizesWallet contract.
type IPrizesWalletTimeoutDurationToWithdrawPrizesChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTimeoutDurationToWithdrawPrizesChanged is a free log retrieval operation binding the contract event 0x8717bb199c6bc4a5dadb21547205f9ef8ec037dda246a5526d6a6471306ea52e.
//
// Solidity: event TimeoutDurationToWithdrawPrizesChanged(uint256 newValue)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterTimeoutDurationToWithdrawPrizesChanged(opts *bind.FilterOpts) (*IPrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator, error) {

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "TimeoutDurationToWithdrawPrizesChanged")
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator{contract: _IPrizesWallet.contract, event: "TimeoutDurationToWithdrawPrizesChanged", logs: logs, sub: sub}, nil
}

// WatchTimeoutDurationToWithdrawPrizesChanged is a free log subscription operation binding the contract event 0x8717bb199c6bc4a5dadb21547205f9ef8ec037dda246a5526d6a6471306ea52e.
//
// Solidity: event TimeoutDurationToWithdrawPrizesChanged(uint256 newValue)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchTimeoutDurationToWithdrawPrizesChanged(opts *bind.WatchOpts, sink chan<- *IPrizesWalletTimeoutDurationToWithdrawPrizesChanged) (event.Subscription, error) {

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "TimeoutDurationToWithdrawPrizesChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPrizesWalletTimeoutDurationToWithdrawPrizesChanged)
				if err := _IPrizesWallet.contract.UnpackLog(event, "TimeoutDurationToWithdrawPrizesChanged", log); err != nil {
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

// ParseTimeoutDurationToWithdrawPrizesChanged is a log parse operation binding the contract event 0x8717bb199c6bc4a5dadb21547205f9ef8ec037dda246a5526d6a6471306ea52e.
//
// Solidity: event TimeoutDurationToWithdrawPrizesChanged(uint256 newValue)
func (_IPrizesWallet *IPrizesWalletFilterer) ParseTimeoutDurationToWithdrawPrizesChanged(log types.Log) (*IPrizesWalletTimeoutDurationToWithdrawPrizesChanged, error) {
	event := new(IPrizesWalletTimeoutDurationToWithdrawPrizesChanged)
	if err := _IPrizesWallet.contract.UnpackLog(event, "TimeoutDurationToWithdrawPrizesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPrizesWalletTokenDonatedIterator is returned from FilterTokenDonated and is used to iterate over the raw logs and unpacked data for TokenDonated events raised by the IPrizesWallet contract.
type IPrizesWalletTokenDonatedIterator struct {
	Event *IPrizesWalletTokenDonated // Event containing the contract specifics and raw log

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
func (it *IPrizesWalletTokenDonatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPrizesWalletTokenDonated)
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
		it.Event = new(IPrizesWalletTokenDonated)
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
func (it *IPrizesWalletTokenDonatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPrizesWalletTokenDonatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPrizesWalletTokenDonated represents a TokenDonated event raised by the IPrizesWallet contract.
type IPrizesWalletTokenDonated struct {
	RoundNum     *big.Int
	DonorAddress common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenDonated is a free log retrieval operation binding the contract event 0x3f94f6171c96ab491d8c268f80da5e460b148878065711a0bb099bbc297615af.
//
// Solidity: event TokenDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed tokenAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterTokenDonated(opts *bind.FilterOpts, roundNum []*big.Int, donorAddress []common.Address, tokenAddress []common.Address) (*IPrizesWalletTokenDonatedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "TokenDonated", roundNumRule, donorAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletTokenDonatedIterator{contract: _IPrizesWallet.contract, event: "TokenDonated", logs: logs, sub: sub}, nil
}

// WatchTokenDonated is a free log subscription operation binding the contract event 0x3f94f6171c96ab491d8c268f80da5e460b148878065711a0bb099bbc297615af.
//
// Solidity: event TokenDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed tokenAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchTokenDonated(opts *bind.WatchOpts, sink chan<- *IPrizesWalletTokenDonated, roundNum []*big.Int, donorAddress []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "TokenDonated", roundNumRule, donorAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPrizesWalletTokenDonated)
				if err := _IPrizesWallet.contract.UnpackLog(event, "TokenDonated", log); err != nil {
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

// ParseTokenDonated is a log parse operation binding the contract event 0x3f94f6171c96ab491d8c268f80da5e460b148878065711a0bb099bbc297615af.
//
// Solidity: event TokenDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed tokenAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) ParseTokenDonated(log types.Log) (*IPrizesWalletTokenDonated, error) {
	event := new(IPrizesWalletTokenDonated)
	if err := _IPrizesWallet.contract.UnpackLog(event, "TokenDonated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizesWalletMetaData contains all meta data concerning the PrizesWallet contract.
var PrizesWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"game_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DonatedNftAlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operationPermittedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"DonatedNftClaimDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operationPermittedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"DonatedTokenClaimDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"prizeWinnerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operationPermittedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawalDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidDonatedNftIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"}],\"name\":\"UnauthorizedCaller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DonatedNftClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonatedTokenClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prizeWinnerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prizeWinnerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"NftDonated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"TimeoutDurationToWithdrawPrizesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenDonated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"claimDonatedNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"claimDonatedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"indexes_\",\"type\":\"uint256[]\"}],\"name\":\"claimManyDonatedNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.DonatedTokenToClaim[]\",\"name\":\"donatedTokensToClaim_\",\"type\":\"tuple[]\"}],\"name\":\"claimManyDonatedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"prizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress_\",\"type\":\"address\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"donateNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress_\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"donateToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"donatedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"donatedTokens\",\"outputs\":[{\"internalType\":\"contractDonatedTokenHolder\",\"name\":\"holder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"}],\"name\":\"getDonatedTokenBalanceAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEthBalanceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.EthBalanceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"getEthBalanceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.EthBalanceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mainPrizeBeneficiaryAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextDonatedNftIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mainPrizeBeneficiaryAddress_\",\"type\":\"address\"}],\"name\":\"registerRoundEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mainPrizeBeneficiaryAddress_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"prizeWinnerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.EthDeposit[]\",\"name\":\"ethDeposits_\",\"type\":\"tuple[]\"}],\"name\":\"registerRoundEndAndDepositEthMany\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roundTimeoutTimesToWithdrawPrizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setTimeoutDurationToWithdrawPrizes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeoutDurationToWithdrawPrizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"withdrawEth_\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPrizesWallet.DonatedTokenToClaim[]\",\"name\":\"donatedTokensToClaim_\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"donatedNftIndexes_\",\"type\":\"uint256[]\"}],\"name\":\"withdrawEverything\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523461003f57610019610014610104565b61020b565b610021610044565b613ddc6104db8239608051818181610c0701526123420152613ddc90f35b61004a565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100769061004e565b810190811060018060401b0382111761008e57604052565b610058565b906100a661009f610044565b928361006c565b565b5f80fd5b60018060a01b031690565b6100c0906100ac565b90565b6100cc816100b7565b036100d357565b5f80fd5b905051906100e4826100c3565b565b906020828203126100ff576100fc915f016100d7565b90565b6100a8565b6101226142b78038038061011781610093565b9283398101906100e6565b90565b90565b90565b90565b61014261013d61014792610125565b61012b565b610128565b90565b610156622e248061012e565b90565b5f1b90565b9061016a5f1991610159565b9181191691161790565b61018861018361018d92610128565b61012b565b610128565b90565b90565b906101a86101a36101af92610174565b610190565b825461015e565b9055565b90565b6101ca6101c56101cf926101b3565b61012b565b610128565b90565b906101e76101e26101ee926101b6565b610190565b825461015e565b9055565b610204906101ff8161038f565b610206565b565b608052565b61025d9061021f61021a610320565b61025f565b61023961022a61014a565b68010000000000000001610193565b6102585f740200000000000000000000000300000000000000026101d2565b6101f2565b565b6102689061026a565b565b61027390610275565b565b61027e906102ca565b565b61029461028f610299926101b3565b61012b565b6100ac565b90565b6102a590610280565b90565b6102b1906100b7565b9052565b91906102c8905f602085019401906102a8565b565b806102e56102df6102da5f61029c565b6100b7565b916100b7565b146102f5576102f39061047b565b565b6103186103015f61029c565b5f918291631e4fbdf760e01b8352600483016102b5565b0390fd5b5f90565b61032861031c565b503390565b60209181520190565b5f7f5468652070726f76696465642061646472657373206973207a65726f2e000000910152565b61036a601d60209261032d565b61037381610336565b0190565b61038c9060208101905f81830391015261035d565b90565b6103a96103a361039e5f61029c565b6100b7565b916100b7565b146103b057565b6103b8610044565b63eac0d38960e01b8152806103cf60048201610377565b0390fd5b5f1c90565b60018060a01b031690565b6103ef6103f4916103d3565b6103d8565b90565b61040190546103e3565b90565b9061041560018060a01b0391610159565b9181191691161790565b61043361042e610438926100ac565b61012b565b6100ac565b90565b6104449061041f565b90565b6104509061043b565b90565b90565b9061046b61046661047292610447565b610453565b8254610404565b9055565b5f0190565b6104845f6103f7565b61048e825f610456565b906104c26104bc7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093610447565b91610447565b916104cb610044565b806104d581610476565b0390a356fe60806040526004361015610013575b610fc8565b61001d5f356101bc565b80631a18c889146101b757806325e16063146101b257806346f6b4e1146101ad5780634b5e1b19146101a85780635e00aaba146101a35780636224dd3f1461019e5780636ff1facd14610199578063715018a61461019457806376e42c6a1461018f57806387565d141461018a5780638da5cb5b1461018557806394d907fc146101805780639523b2371461017b5780639cf10d32146101765780639e2842a814610171578063a089e0be1461016c578063a0ef91df14610167578063b4e6f80314610162578063c3fe3e281461015d578063cc5810d814610158578063d7f4f8be14610153578063e2051c7e1461014e578063e4a6c2a414610149578063e86a49d714610144578063f2fde38b1461013f5763fe673fd30361000e57610f91565b610ee9565b610eb3565b610e52565b610de5565b610d6a565b610c5e565b610c29565b610bce565b610ac9565b610a94565b610a61565b610a29565b610898565b6107fc565b6107c7565b610798565b6106d2565b6105e7565b6105b2565b61055a565b61050f565b610459565b610341565b6102be565b610228565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f9103126101da57565b6101cc565b90565b6101eb906101df565b9052565b90602080610211936102075f8201515f8601906101e2565b01519101906101e2565b565b9190610226905f604085019401906101ef565b565b34610258576102383660046101d0565b6102546102436110fd565b61024b6101c2565b91829182610213565b0390f35b6101c8565b5f80fd5b60018060a01b031690565b61027590610261565b90565b6102818161026c565b0361028857565b5f80fd5b9050359061029982610278565b565b906020828203126102b4576102b1915f0161028c565b90565b6101cc565b5f0190565b346102ec576102d66102d136600461029b565b611595565b6102de6101c2565b806102e8816102b9565b0390f35b6101c8565b6102fa816101df565b0361030157565b5f80fd5b90503590610312826102f1565b565b919060408382031261033c5780610330610339925f8601610305565b9360200161028c565b90565b6101cc565b61035561034f366004610314565b906115df565b61035d6101c2565b80610367816102b9565b0390f35b9060208282031261038457610381915f01610305565b90565b6101cc565b634e487b7160e01b5f52603260045260245ffd5b506801000000000000000090565b90565b6103b78161039d565b8210156103d1576103c96001916103ab565b910201905f90565b610389565b1c90565b90565b6103ed9060086103f293026103d6565b6103da565b90565b9061040091546103dd565b90565b680100000000000000026104168161039d565b821015610433576104309161042a916103ae565b906103f5565b90565b5f80fd5b610440906101df565b9052565b9190610457905f60208501940190610437565b565b346104895761048561047461046f36600461036b565b610403565b61047c6101c2565b91829182610444565b0390f35b6101c8565b5f80fd5b5f80fd5b5f80fd5b909182601f830112156104d45781359167ffffffffffffffff83116104cf5760200192606083028401116104ca57565b610496565b610492565b61048e565b9060208282031261050a575f82013567ffffffffffffffff811161050557610501920161049a565b9091565b61025d565b6101cc565b3461053e576105286105223660046104d9565b90611613565b6105306101c2565b8061053a816102b9565b0390f35b6101c8565b610557680100000000000000015f906103f5565b90565b3461058a5761056a3660046101d0565b610586610575610543565b61057d6101c2565b91829182610444565b0390f35b6101c8565b6105af740200000000000000000000000300000000000000025f906103f5565b90565b346105e2576105c23660046101d0565b6105de6105cd61058f565b6105d56101c2565b91829182610444565b0390f35b6101c8565b34610615576105f73660046101d0565b6105ff61166c565b6106076101c2565b80610611816102b9565b0390f35b6101c8565b506801000000000000000090565b90565b6106348161061a565b82101561064e57610646600191610628565b910201905f90565b610389565b60018060a01b031690565b61066e90600861067393026103d6565b610653565b90565b90610681915461065e565b90565b600161068f8161061a565b8210156106ac576106a9916106a39161062b565b90610676565b90565b5f80fd5b6106b99061026c565b9052565b91906106d0905f602085019401906106b0565b565b34610702576106fe6106ed6106e836600461036b565b610684565b6106f56101c2565b918291826106bd565b0390f35b6101c8565b909182601f830112156107415781359167ffffffffffffffff831161073c57602001926040830284011161073757565b610496565b610492565b61048e565b916060838303126107935761075d825f8501610305565b9261076b836020830161028c565b92604082013567ffffffffffffffff811161078e5761078a9201610707565b9091565b61025d565b6101cc565b6107c36107b26107a9366004610746565b9291909161178d565b6107ba6101c2565b91829182610444565b0390f35b346107f7576107d73660046101d0565b6107f36107e26117c9565b6107ea6101c2565b918291826106bd565b0390f35b6101c8565b3461082a5761081461080f36600461036b565b611804565b61081c6101c2565b80610826816102b9565b0390f35b6101c8565b6108389061026c565b90565b6108448161082f565b0361084b57565b5f80fd5b9050359061085c8261083b565b565b909160608284031261089357610890610879845f8501610305565b93610887816020860161084f565b93604001610305565b90565b6101cc565b346108c7576108b16108ab36600461085e565b9161183b565b6108b96101c2565b806108c3816102b9565b0390f35b6101c8565b506801000000000000000090565b90565b6108e6816108cc565b821015610900576108f86003916108da565b910201905f90565b610389565b5f1c90565b61091661091b91610905565b6103da565b90565b610928905461090a565b90565b60018060a01b031690565b61094261094791610905565b61092b565b90565b6109549054610936565b90565b7402000000000000000000000003000000000000000390610977826108cc565b8110156109af57610987916108dd565b506109935f820161091e565b916109ac60026109a56001850161094a565b930161091e565b90565b5f80fd5b90565b6109ca6109c56109cf92610261565b6109b3565b610261565b90565b6109db906109b6565b90565b6109e7906109d2565b90565b6109f3906109de565b9052565b604090610a20610a279496959396610a1660608401985f850190610437565b60208301906109ea565b0190610437565b565b34610a5c57610a58610a44610a3f36600461036b565b610957565b610a4f9391936101c2565b938493846109f7565b0390f35b6101c8565b34610a8f57610a79610a7436600461036b565b6118e4565b610a816101c2565b80610a8b816102b9565b0390f35b6101c8565b34610ac457610ac0610aaf610aaa36600461029b565b6118ef565b610ab76101c2565b91829182610213565b0390f35b6101c8565b34610af757610ad93660046101d0565b610ae1611945565b610ae96101c2565b80610af3816102b9565b0390f35b6101c8565b151590565b610b0a81610afc565b03610b1157565b5f80fd5b90503590610b2282610b01565b565b909182601f83011215610b5e5781359167ffffffffffffffff8311610b59576020019260208302840111610b5457565b610496565b610492565b61048e565b606081830312610bc957610b79825f8301610b15565b92602082013567ffffffffffffffff8111610bc45783610b9a91840161049a565b929093604082013567ffffffffffffffff8111610bbf57610bbb9201610b24565b9091565b61025d565b61025d565b6101cc565b34610c0057610bea610be1366004610b63565b93929092611997565b610bf26101c2565b80610bfc816102b9565b0390f35b6101c8565b7f000000000000000000000000000000000000000000000000000000000000000090565b34610c5957610c393660046101d0565b610c55610c44610c05565b610c4c6101c2565b918291826106bd565b0390f35b6101c8565b34610c8f57610c8b610c7a610c74366004610314565b906119e9565b610c826101c2565b91829182610444565b0390f35b6101c8565b506801000000000000000090565b90565b610cae81610c94565b821015610cc857610cc0600191610ca2565b910201905f90565b610389565b60018060a01b031690565b610ce4610ce991610905565b610ccd565b90565b610cf69054610cd8565b90565b7402000000000000000000000002000000000000000290610d1982610c94565b811015610d3857610d2e5f91610d3593610ca5565b5001610cec565b90565b5f80fd5b610d45906109d2565b90565b610d5190610d3c565b9052565b9190610d68905f60208501940190610d48565b565b34610d9a57610d96610d85610d8036600461036b565b610cf9565b610d8d6101c2565b91829182610d55565b0390f35b6101c8565b608081830312610de057610db5825f8301610305565b92610ddd610dc6846020850161028c565b93610dd4816040860161084f565b93606001610305565b90565b6101cc565b34610e1757610e01610df8366004610d9f565b92919091611d53565b610e096101c2565b80610e13816102b9565b0390f35b6101c8565b90602082820312610e4d575f82013567ffffffffffffffff8111610e4857610e449201610b24565b9091565b61025d565b6101cc565b34610e8157610e6b610e65366004610e1c565b90611d89565b610e736101c2565b80610e7d816102b9565b0390f35b6101c8565b9190604083820312610eae5780610ea2610eab925f8601610305565b9360200161084f565b90565b6101cc565b34610ee457610ee0610ecf610ec9366004610e86565b90611dc2565b610ed76101c2565b91829182610444565b0390f35b6101c8565b34610f1757610f01610efc36600461029b565b611f39565b610f096101c2565b80610f13816102b9565b0390f35b6101c8565b610f259061026c565b90565b610f3181610f1c565b03610f3857565b5f80fd5b90503590610f4982610f28565b565b608081830312610f8c57610f61825f8301610305565b92610f89610f72846020850161028c565b93610f808160408601610f3c565b93606001610305565b90565b6101cc565b34610fc357610fad610fa4366004610f4b565b929190916121b6565b610fb56101c2565b80610fbf816102b9565b0390f35b6101c8565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b90610ff490610fcc565b810190811067ffffffffffffffff82111761100e57604052565b610fd6565b9061102661101f6101c2565b9283610fea565b565b6110326040611013565b90565b5f90565b611041611028565b906020808361104e611035565b815201611059611035565b81525050565b611067611039565b90565b611073906109b6565b90565b50600160a01b90565b90565b61108b81611076565b8210156110a55761109d60029161107f565b910201905f90565b610389565b906110b4906101df565b9052565b906110ef6110e660016110c9611028565b946110e06110d85f830161091e565b5f88016110aa565b0161091e565b602084016110aa565b565b6110fa906110b8565b90565b61110561105f565b5061113261112c680200000000000000026111266111216121c4565b61106a565b90611082565b506110f1565b90565b6111469061114161221f565b611418565b61114e612264565b565b90565b90565b61116a61116561116f92611153565b6109b3565b6101df565b90565b60209181520190565b60407f206265666f726520612074696d656f757420657870697265732e000000000000917f4f6e6c792074686520455448207072697a652077696e6e6572206973207065725f8201527f6d697474656420746f2077697468647261772074686569722062616c616e636560208201520152565b6111fb605a606092611172565b6112048161117b565b0190565b61124d6112549461124360809498979561123961122e60a087018781035f8901526111ee565b9a60208701906106b0565b60408501906106b0565b6060830190610437565b0190610437565b565b939193929092156112675750505050565b9061128a916112746101c2565b94859463380b9dd360e21b865260048601611208565b0390fd5b1b90565b919060086112ad9102916112a75f198461128e565b9261128e565b9181191691161790565b6112cb6112c66112d0926101df565b6109b3565b6101df565b90565b90565b91906112ec6112e76112f4936112b7565b6112d3565b908354611292565b9055565b5f90565b61130e916113086112f8565b916112d6565b565b611319906109d2565b90565b905090565b61132c5f809261131c565b0190565b61133990611321565b90565b67ffffffffffffffff811161135a57611356602091610fcc565b0190565b610fd6565b9061137161136c8361133c565b611013565b918252565b606090565b3d5f146113965761138b3d61135f565b903d5f602084013e5b565b61139e611376565b90611394565b5f7f455448207769746864726177616c206661696c65642e00000000000000000000910152565b6113d86016602092611172565b6113e1816113a4565b0190565b91604061141692949361140f611404606083018381035f8501526113cb565b9660208301906106b0565b0190610437565b565b61143d611437680200000000000000026114318461106a565b90611082565b50611150565b906114976114686114626801000000000000000261145c5f870161091e565b906103ae565b906103f5565b4261147b611475836101df565b916101df565b10159081611579575b839061148e6121c4565b90914293611256565b6114bc5f806114a86001860161091e565b946114b682600183016112fc565b016112fc565b6114c46121c4565b829161150e6114fc6114f67f4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f5929238393611310565b93611310565b936115056101c2565b91829182610444565b0390a36115455f8061151e6121c4565b846115276101c2565b908161153281611330565b03925af161153e61137b565b5015610afc565b61154c5750565b6115546121c4565b61157561155f6101c2565b928392630aa7db6360e11b8452600484016113e5565b0390fd5b80915061158e6115885f611156565b916101df565b1190611484565b61159e90611135565b565b906115b2916115ad61221f565b6115bc565b6115ba612264565b565b906115ce916115c9612332565b6115d0565b565b906115dd919034916123c5565b565b906115e9916115a0565b565b906115fd916115f861221f565b611607565b611605612264565b565b906116119161248f565b565b9061161d916115eb565b565b611627612505565b61162f611659565b565b61164561164061164a92611153565b6109b3565b610261565b90565b61165690611631565b90565b61166a6116655f61164d565b612576565b565b61167461161f565b565b9061168b9493929161168661221f565b611696565b90611694612264565b565b906116ab949392916116a6612332565b611711565b90565b5090565b634e487b7160e01b5f52601160045260245ffd5b6116cf906101df565b5f81146116dd576001900390565b6116b2565b91908110156116f2576040020190565b610389565b3561170181610278565b90565b3561170e816102f1565b90565b91611720919495925084612621565b9061172c8386906116ae565b5b8061174061173a5f611156565b916101df565b11156117845761174f906116c6565b9161177e61175f858886916116e2565b869061177860206117715f84016116f7565b9201611704565b916123c5565b9161172d565b50935091505090565b906117a193929161179c6112f8565b611676565b90565b5f90565b6117b46117b991610905565b610653565b90565b6117c690546117a8565b90565b6117d16117a4565b506117db5f6117bc565b90565b6117ef906117ea61221f565b6117f9565b6117f7612264565b565b6118029061296b565b565b61180d906117de565b565b90611822929161181d61221f565b61182c565b61182a612264565b565b9161183992919091612dbd565b565b90611846929161180f565b565b61185990611854612505565b611896565b565b5f1b90565b9061186c5f199161185b565b9181191691161790565b9061188b611886611892926112b7565b6112d3565b8254611860565b9055565b6118a98168010000000000000001611876565b6118df7f8717bb199c6bc4a5dadb21547205f9ef8ec037dda246a5526d6a6471306ea52e916118d66101c2565b91829182610444565b0390a1565b6118ed90611848565b565b61191861191e916118fe61105f565b50611912680200000000000000029161106a565b90611082565b506110f1565b90565b61192961221f565b61193161193b565b611939612264565b565b611943612fe3565b565b61194d611921565b565b906119649493929161195f61221f565b61196e565b61196c612264565b565b916119839161198895949361198a575b61248f565b613112565b565b611992612fe3565b61197e565b906119a49493929161194f565b565b906119b992916119b461221f565b6119c4565b906119c2612264565b565b906119d792916119d2612332565b6119da565b90565b906119e6929150612621565b90565b906119fb916119f66112f8565b6119a6565b90565b90611a12939291611a0d61221f565b611a1c565b611a1a612264565b565b90611a30939291611a2b612332565b611b5b565b565b90565b611a3e9061026c565b90565b90611a4b90611a35565b9052565b611a596020611013565b90565b90611a7b611a735f611a6c611a4f565b9401610cec565b5f8401611a41565b565b611a8690611a5c565b90565b611a939051611a35565b90565b5f80fd5b60e01b90565b5f910312611aaa57565b6101cc565b611ab8906109d2565b90565b611ac490611aaf565b9052565b9190611adb905f60208501940190611abb565b565b611ae56101c2565b3d5f823e3d90fd5b90611afe60018060a01b039161185b565b9181191691161790565b611b11906109b6565b90565b611b1d90611b08565b90565b90565b90611b38611b33611b3f92611b14565b611b20565b8254611aed565b9055565b611b4c906109b6565b90565b611b5890611b43565b90565b92909192611b88611b82740200000000000000000000000200000000000000028390610ca5565b50611a32565b90611b9282611a7d565b91611ba6611ba15f8501611a89565b610d3c565b611bc0611bba611bb55f61164d565b61026c565b9161026c565b145f14611cb05785611bd06101c2565b90610a8a820182811067ffffffffffffffff821117611cab578291611bfc91610a8a61331d8539611ac8565b03905ff0908115611ca657611ca496611c9d93611c365f93611c23611c9896868a01611a41565b84611c2f818a01611a89565b9101611b23565b5b8782908892611c8d611c7b611c75611c6f7f3f94f6171c96ab491d8c268f80da5e460b148878065711a0bb099bbc297615af946112b7565b94611310565b94611b4f565b94611c846101c2565b91829182610444565b0390a4959301611a89565b610d3c565b919261319c565b565b611add565b610fd6565b50611cc4611cbf5f8401611a89565b610d3c565b630b820b5486823b15611d4e57611cfa92611cef5f8094611ce36101c2565b96879586948593611a9a565b835260048301611ac8565b03925af18015611d4957611ca496611c9d93611c98935f93611d1d575b50611c37565b611d3c90843d8111611d42575b611d348183610fea565b810190611aa0565b5f611d17565b503d611d2a565b611add565b611a96565b90611d5f9392916119fe565b565b90611d7391611d6e61221f565b611d7d565b611d7b612264565b565b90611d8791613112565b565b90611d9391611d61565b565b90505190611da2826102f1565b565b90602082820312611dbd57611dba915f01611d95565b90565b6101cc565b611df6611df0611dfb92611dd46112f8565b5074020000000000000000000000020000000000000002610ca5565b50611a32565b611a7d565b90611e0f611e0a5f8401611a89565b610d3c565b611e29611e23611e1e5f61164d565b61026c565b9161026c565b145f14611e3f575050611e3b5f611156565b5b90565b611e8991611e4e602092611aaf565b611e7e611e69611e645f6370a082319501611a89565b610d3c565b92611e726101c2565b95869485938493611a9a565b8352600483016106bd565b03915afa908115611ecf575f91611ea1575b50611e3c565b611ec2915060203d8111611ec8575b611eba8183610fea565b810190611da4565b5f611e9b565b503d611eb0565b611add565b611ee590611ee0612505565b611ee7565b565b80611f02611efc611ef75f61164d565b61026c565b9161026c565b14611f1257611f1090612576565b565b611f35611f1e5f61164d565b5f918291631e4fbdf760e01b8352600483016106bd565b0390fd5b611f4290611ed4565b565b90611f58939291611f5361221f565b611f62565b611f60612264565b565b90611f76939291611f71612332565b612033565b565b90565b611f84906109b6565b90565b611f9090611f7b565b90565b90565b90611fab611fa6611fb292611f87565b611f93565b8254611aed565b9055565b916020611fd7929493611fd060408201965f830190610437565b0190610437565b565b611fe2906101df565b5f198114611ff05760010190565b6116b2565b611ffe906109d2565b90565b60409061202a612031949695939661202060608401985f8501906106b0565b60208301906106b0565b0190610437565b565b9161212761210c61212c939461205d7402000000000000000000000003000000000000000261091e565b906120ad61208a6120847402000000000000000000000003000000000000000385906108dd565b50611f78565b612096835f8301611876565b6120a38760018301611f96565b60028a9101611876565b8685908992846120ef6120e96120e37fb12e72bab0c2a8fe98175a3c02792645ebdf65d37cf3983517e6d1c5ab842e23946112b7565b94611310565b94611f87565b946121046120fb6101c2565b92839283611fb6565b0390a4611fd9565b74020000000000000000000000030000000000000002611876565b6109de565b6323b872dd919061213c30611ff5565b9392813b156121b1575f6121639161216e82966121576101c2565b98899788968795611a9a565b855260048501612001565b03925af180156121ac57612180575b50565b61219f905f3d81116121a5575b6121978183610fea565b810190611aa0565b5f61217d565b503d61218d565b611add565b611a96565b906121c2939291611f44565b565b6121cc6117a4565b503390565b90565b90565b6121eb6121e66121f0926121d1565b61185b565b6121d4565b90565b61221c7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f006121d7565b90565b6122276131f0565b6122485761224661223e6122396121f3565b613224565b600190613239565b565b5f633ee5aeb560e01b815280612260600482016102b9565b0390fd5b61227e6122776122726121f3565b613224565b5f90613239565b565b60407f6d6574686f642e00000000000000000000000000000000000000000000000000917f4f6e6c792074686520436f736d69635369676e617475726547616d6520636f6e5f8201527f7472616374206973207065726d697474656420746f2063616c6c20746869732060208201520152565b6123006047606092611172565b61230981612280565b0190565b9190612330906020612328604086018681035f8801526122f3565b9401906106b0565b565b61233a6121c4565b61236c6123667f000000000000000000000000000000000000000000000000000000000000000061026c565b9161026c565b0361237357565b61239c61237e6121c4565b6123866101c2565b91829163ced50f6760e01b83526004830161230d565b0390fd5b6123af6123b5919392936101df565b926101df565b82018092116123c057565b6116b2565b9190916124196123f06123ea680200000000000000026123e48761106a565b90611082565b50611150565b6123fc835f8301611876565b61241360018592019161240e8361091e565b6123a0565b90611876565b91909161246461245261244c7f999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326936112b7565b93611310565b9361245b6101c2565b91829182610444565b0390a3565b5090565b919081101561247d576060020190565b610389565b3561248c8161083b565b90565b909161249c828490612469565b5b806124b06124aa5f611156565b916101df565b11156124ff576124bf906116c6565b906124f96124cf8486859161246d565b6124da5f8201611704565b906124f360406124ec60208401612482565b9201611704565b91612dbd565b9061249d565b50915050565b61250d6117c9565b61252661252061251b6121c4565b61026c565b9161026c565b0361252d57565b61254f6125386121c4565b5f91829163118cdaa760e01b8352600483016106bd565b0390fd5b90565b9061256b61256661257292611310565b612553565b8254611aed565b9055565b61257f5f6117bc565b612589825f612556565b906125bd6125b77f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093611310565b91611310565b916125c66101c2565b806125d0816102b9565b0390a3565b919060086125f59102916125ef60018060a01b038461128e565b9261128e565b9181191691161790565b919061261561261061261d93611310565b612553565b9083546125d5565b9055565b906126419061262e6112f8565b5061263b6001849061062b565b906125ff565b61267b6126604261265a6801000000000000000161091e565b906123a0565b916126758391680100000000000000026103ae565b906112d6565b90565b9061268890610f1c565b9052565b6126966060611013565b90565b906126e86126df60026126aa61268c565b946126c16126b95f830161091e565b5f88016110aa565b6126d96126d06001830161094a565b6020880161267e565b0161091e565b604084016110aa565b565b6126f390612699565b90565b6127009051610f1c565b90565b5f7f496e76616c696420646f6e61746564204e465420696e6465782e000000000000910152565b612737601a602092611172565b61274081612703565b0190565b91604061277592949361276e612763606083018381035f85015261272a565b9660208301906106b0565b0190610437565b565b5f7f446f6e61746564204e465420616c726561647920636c61696d65642e00000000910152565b6127ab601c602092611172565b6127b481612777565b0190565b9160406127e99294936127e26127d7606083018381035f85015261279e565b9660208301906106b0565b0190610437565b565b6127f590516101df565b90565b60607f70697265732e0000000000000000000000000000000000000000000000000000917f4f6e6c79207468652062696464696e6720726f756e64206d61696e207072697a5f8201527f652062656e6566696369617279206973207065726d697474656420746f20636c60208201527f61696d2074686973204e4654206265666f726520612074696d656f757420657860408201520152565b61289e6066608092611172565b6128a7816127f8565b0190565b6128f06128f7946128e66080949897956128dc6128d160a087018781035f890152612891565b9a60208701906106b0565b6040850190610437565b6060830190610437565b0190610437565b565b9391939290921561290a5750505050565b9061292d916129176101c2565b94859463b97dcd1760e01b8652600486016128ab565b0390fd5b919061294761294261294f93611f87565b611f93565b9083546125d5565b9055565b5f90565b61296991612963612953565b91612931565b565b61299461298e7402000000000000000000000003000000000000000383906108dd565b50611f78565b9061299e826126ea565b916129b36129ae602085016126f6565b6109de565b6129cd6129c76129c25f61164d565b61026c565b9161026c565b14612be2575f6002612a36926129e16121c4565b612a11612a0b612a06612a0060016129fa898d016127eb565b9061062b565b90610676565b61026c565b9161026c565b03612b68575b612a23838083016112fc565b612a308360018301612957565b016112fc565b612a415f83016127eb565b612a496121c4565b612a55602085016126f6565b91612a62604086016127eb565b93612a9f612a99612a937f03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3946112b7565b94611310565b94611f87565b94612ab4612aab6101c2565b92839283611fb6565b0390a4612acb612ac6602083016126f6565b6109de565b6323b872dd90612ada30611ff5565b90612aef6040612ae86121c4565b95016127eb565b92813b15612b63575f612b1591612b208296612b096101c2565b98899788968795611a9a565b855260048501612001565b03925af18015612b5e57612b32575b50565b612b51905f3d8111612b57575b612b498183610fea565b810190611aa0565b5f612b2f565b503d612b3f565b611add565b611a96565b612bc1612b92612b8c68010000000000000002612b86878b016127eb565b906103ae565b906103f5565b42612ba5612b9f836101df565b916101df565b10159081612bc6575b612bb66121c4565b9088909142936128f9565b612a17565b809150612bdb612bd587611156565b916101df565b1190612bae565b5080612c13612c0d612c087402000000000000000000000003000000000000000261091e565b6101df565b916101df565b1015612c4657612c216121c4565b612c42612c2c6101c2565b92839263581c778d60e01b8452600484016127b8565b0390fd5b612c4e6121c4565b612c6f612c596101c2565b9283926373b047ef60e11b845260048401612744565b0390fd5b60607f65666f726520612074696d656f757420657870697265732e0000000000000000917f4f6e6c79207468652062696464696e6720726f756e64206d61696e207072697a5f8201527f652062656e6566696369617279206973207065726d697474656420746f20636c60208201527f61696d2074686973204552432d323020746f6b656e20646f6e6174696f6e206260408201520152565b612d196078608092611172565b612d2281612c73565b0190565b90959492612d8194612d70612d7a92612d6660a096612d5c612d5160c089018981035f8b0152612d0c565b9c6020890190610437565b60408701906106b0565b6060850190611abb565b6080830190610437565b0190610437565b565b94929093919415612d95575050505050565b90612db99291612da36101c2565b95869563c8568d6360e01b875260048701612d26565b0390fd5b9091612dc76121c4565b612ded612de7612de2612ddc6001879061062b565b90610676565b61026c565b9161026c565b03612f72575b612e24612e1f612e19740200000000000000000000000200000000000000028590610ca5565b50611a32565b611a7d565b9181612e38612e325f611156565b916101df565b14612ec7575b612eb15f612ec595612eb693612e526121c4565b82908792612ea7612e95612e8f612e897faf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0946112b7565b94611310565b94611b4f565b94612e9e6101c2565b91829182610444565b0390a49401611a89565b610d3c565b612ebe6121c4565b919261319c565b565b9050826020612ed8612f1495611aaf565b6370a0823190612f09612ef4612eef5f8901611a89565b610d3c565b92612efd6101c2565b98899485938493611a9a565b8352600483016106bd565b03915afa8015612f6d575f612eb692612eb192612ec5978391612f3f575b5094935095505050612e3e565b612f60915060203d8111612f66575b612f588183610fea565b810190611da4565b5f612f32565b503d612f4e565b611add565b612fc2612f92612f8c6801000000000000000285906103ae565b906103f5565b42612fa5612f9f836101df565b916101df565b10159081612fc7575b8490612fb86121c4565b8791924294612d83565b612df3565b809150612fdc612fd65f611156565b916101df565b1190612fae565b61300f61300968020000000000000002613003612ffe6121c4565b61106a565b90611082565b50611150565b6130345f806130206001850161091e565b9361302e82600183016112fc565b016112fc565b61303c6121c4565b6130446121c4565b829161308e61307c6130767f4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f5929238393611310565b93611310565b936130856101c2565b91829182610444565b0390a36130c55f8061309e6121c4565b846130a76101c2565b90816130b281611330565b03925af16130be61137b565b5015610afc565b6130cc5750565b6130d46121c4565b6130f56130df6101c2565b928392630aa7db6360e11b8452600484016113e5565b0390fd5b5090565b919081101561310d576020020190565b610389565b909161311f8284906130f9565b5b8061313361312d5f611156565b916101df565b111561316557613142906116c6565b9061315f61315a613155858786916130fd565b611704565b61296b565b90613120565b50915050565b63ffffffff1690565b63ffffffff60e01b1690565b61319461318f6131999261316b565b611a9a565b613174565b90565b6004926131d66131ea95936131e593946131bd6323b872dd92949192613180565b936131c66101c2565b9788956020870190815201612001565b60208201810382520383610fea565b61325b565b565b5f90565b6131f86131ec565b5061321161320c6132076121f3565b613224565b61330f565b90565b5f90565b613221906121d4565b90565b61323690613230613214565b50613218565b90565b5d565b90565b61325361324e6132589261323c565b6109b3565b6101df565b90565b905f6020916132686112f8565b506132716112f8565b50828151910182855af115613304573d5f51906132966132905f611156565b916101df565b145f146132ea57506132a781611aaf565b3b6132ba6132b45f611156565b916101df565b145b6132c35750565b6132cf6132e691611aaf565b5f918291635274afe760e01b8352600483016106bd565b0390fd5b6132fd6132f7600161323f565b916101df565b14156132bc565b6040513d5f823e3d90fd5b6133176131ec565b505c9056fe60a06040523461003f57610019610014610110565b610131565b610021610044565b6105d46104b6823960805181818161018e01526101eb01526105d490f35b61004a565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100769061004e565b810190811060018060401b0382111761008e57604052565b610058565b906100a661009f610044565b928361006c565b565b5f80fd5b60018060a01b031690565b6100c0906100ac565b90565b6100cc906100b7565b90565b6100d8816100c3565b036100df57565b5f80fd5b905051906100f0826100cf565b565b9060208282031261010b57610108915f016100e3565b90565b6100a8565b61012e610a8a8038038061012381610093565b9283398101906100f2565b90565b61013e903360805261014d565b565b61014a90516100b7565b90565b6101639061015b6080610140565b5f1991610267565b565b90565b61017c610177610181926100ac565b610165565b6100ac565b90565b61018d90610168565b90565b61019990610184565b90565b63ffffffff1690565b63ffffffff60e01b1690565b60e01b90565b6101cb6101c66101d09261019c565b6101b1565b6101a5565b90565b6101dc906100b7565b9052565b90565b6101ec906101e0565b9052565b91602061021192949361020a60408201965f8301906101d3565b01906101e3565b565b151590565b90565b61022f61022a61023492610218565b610165565b6101e0565b90565b6102409061021b565b9052565b91602061026592949361025e60408201965f8301906101d3565b0190610237565b565b90916102aa60049161029b63095ea7b36102828793916101b7565b9261028b610044565b95869460208601908152016101f0565b6020820181038252038261006c565b906102bf6102b982849061034a565b15610213565b6102c9575b505050565b610316600461031b94610311849161030263095ea7b36102e95f916101b7565b926102f2610044565b9687946020860190815201610244565b6020820181038252038361006c565b610401565b610401565b5f80806102c4565b5f90565b5f90565b90565b61034261033d6103479261032b565b610165565b6101e0565b90565b905f602091610357610323565b50610360610323565b50610369610327565b50610372610327565b50828151910182855af13d915f5191928361038e575b50505090565b909192506103a461039e5f61021b565b916101e0565b145f146103d2576103b59150610190565b3b6103c86103c25f61021b565b916101e0565b115b5f8080610388565b506103e66103e0600161032e565b916101e0565b146103ca565b91906103ff905f602085019401906101d3565b565b905f60209161040e610327565b50610417610327565b50828151910182855af1156104aa573d5f519061043c6104365f61021b565b916101e0565b145f14610490575061044d81610190565b3b61046061045a5f61021b565b916101e0565b145b6104695750565b61047561048c91610190565b5f918291635274afe760e01b8352600483016103ec565b0390fd5b6104a361049d600161032e565b916101e0565b1415610462565b6040513d5f823e3d90fdfe60806040526004361015610013575b6100dc565b61001d5f3561002c565b630b820b540361000e576100a9565b60e01c90565b60405190565b5f80fd5b5f80fd5b60018060a01b031690565b61005490610040565b90565b6100609061004b565b90565b61006c81610057565b0361007357565b5f80fd5b9050359061008482610063565b565b9060208282031261009f5761009c915f01610077565b90565b61003c565b5f0190565b346100d7576100c16100bc366004610086565b6100fe565b6100c9610032565b806100d3816100a4565b0390f35b610038565b5f80fd5b6100f1906100ec610185565b6100f3565b565b6100fc906101e5565b565b610107906100e0565b565b60209181520190565b5f7f4465706c6f796572206f6e6c792e000000000000000000000000000000000000910152565b610146600e602092610109565b61014f81610112565b0190565b61015c9061004b565b9052565b919061018390602061017b604086018681035f880152610139565b940190610153565b565b336101b86101b27f000000000000000000000000000000000000000000000000000000000000000061004b565b9161004b565b036101bf57565b6101e1336101cb610032565b91829163ced50f6760e01b835260048301610160565b0390fd5b610212907f00000000000000000000000000000000000000000000000000000000000000005f1991610350565b565b90565b61022b61022661023092610040565b610214565b610040565b90565b61023c90610217565b90565b61024890610233565b90565b63ffffffff1690565b63ffffffff60e01b1690565b60e01b90565b61027a61027561027f9261024b565b610260565b610254565b90565b90565b61028e90610282565b9052565b9160206102b39294936102ac60408201965f830190610153565b0190610285565b565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906102dd906102b5565b810190811067ffffffffffffffff8211176102f757604052565b6102bf565b151590565b90565b61031861031361031d92610301565b610214565b610282565b90565b61032990610304565b9052565b91602061034e92949361034760408201965f830190610153565b0190610320565b565b909161039360049161038463095ea7b361036b879391610266565b92610374610032565b9586946020860190815201610292565b602082018103825203826102d3565b906103a86103a2828490610433565b156102fc565b6103b2575b505050565b6103ff6004610404946103fa84916103eb63095ea7b36103d25f91610266565b926103db610032565b968794602086019081520161032d565b602082018103825203836102d3565b6104ea565b6104ea565b5f80806103ad565b5f90565b5f90565b90565b61042b61042661043092610414565b610214565b610282565b90565b905f60209161044061040c565b5061044961040c565b50610452610410565b5061045b610410565b50828151910182855af13d915f51919283610477575b50505090565b9091925061048d6104875f610304565b91610282565b145f146104bb5761049e915061023f565b3b6104b16104ab5f610304565b91610282565b115b5f8080610471565b506104cf6104c96001610417565b91610282565b146104b3565b91906104e8905f60208501940190610153565b565b905f6020916104f7610410565b50610500610410565b50828151910182855af115610593573d5f519061052561051f5f610304565b91610282565b145f1461057957506105368161023f565b3b6105496105435f610304565b91610282565b145b6105525750565b61055e6105759161023f565b5f918291635274afe760e01b8352600483016104d5565b0390fd5b61058c6105866001610417565b91610282565b141561054b565b6040513d5f823e3d90fdfea26469706673582212200d1c69ec99903662f13c2ff672c7a937c04a2ca046231ec23ffc3ec22d8614de64736f6c634300081d0033a2646970667358221220f888b2e3da93525ca8433b6f8dfba22847742e70b9c7d90fd9046ccc8ba6584a64736f6c634300081d0033",
}

// PrizesWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use PrizesWalletMetaData.ABI instead.
var PrizesWalletABI = PrizesWalletMetaData.ABI

// PrizesWalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PrizesWalletMetaData.Bin instead.
var PrizesWalletBin = PrizesWalletMetaData.Bin

// DeployPrizesWallet deploys a new Ethereum contract, binding an instance of PrizesWallet to it.
func DeployPrizesWallet(auth *bind.TransactOpts, backend bind.ContractBackend, game_ common.Address) (common.Address, *types.Transaction, *PrizesWallet, error) {
	parsed, err := PrizesWalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PrizesWalletBin), backend, game_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PrizesWallet{PrizesWalletCaller: PrizesWalletCaller{contract: contract}, PrizesWalletTransactor: PrizesWalletTransactor{contract: contract}, PrizesWalletFilterer: PrizesWalletFilterer{contract: contract}}, nil
}

// PrizesWallet is an auto generated Go binding around an Ethereum contract.
type PrizesWallet struct {
	PrizesWalletCaller     // Read-only binding to the contract
	PrizesWalletTransactor // Write-only binding to the contract
	PrizesWalletFilterer   // Log filterer for contract events
}

// PrizesWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrizesWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrizesWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrizesWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrizesWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrizesWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrizesWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrizesWalletSession struct {
	Contract     *PrizesWallet     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrizesWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrizesWalletCallerSession struct {
	Contract *PrizesWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PrizesWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrizesWalletTransactorSession struct {
	Contract     *PrizesWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PrizesWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrizesWalletRaw struct {
	Contract *PrizesWallet // Generic contract binding to access the raw methods on
}

// PrizesWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrizesWalletCallerRaw struct {
	Contract *PrizesWalletCaller // Generic read-only contract binding to access the raw methods on
}

// PrizesWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrizesWalletTransactorRaw struct {
	Contract *PrizesWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrizesWallet creates a new instance of PrizesWallet, bound to a specific deployed contract.
func NewPrizesWallet(address common.Address, backend bind.ContractBackend) (*PrizesWallet, error) {
	contract, err := bindPrizesWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrizesWallet{PrizesWalletCaller: PrizesWalletCaller{contract: contract}, PrizesWalletTransactor: PrizesWalletTransactor{contract: contract}, PrizesWalletFilterer: PrizesWalletFilterer{contract: contract}}, nil
}

// NewPrizesWalletCaller creates a new read-only instance of PrizesWallet, bound to a specific deployed contract.
func NewPrizesWalletCaller(address common.Address, caller bind.ContractCaller) (*PrizesWalletCaller, error) {
	contract, err := bindPrizesWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletCaller{contract: contract}, nil
}

// NewPrizesWalletTransactor creates a new write-only instance of PrizesWallet, bound to a specific deployed contract.
func NewPrizesWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*PrizesWalletTransactor, error) {
	contract, err := bindPrizesWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletTransactor{contract: contract}, nil
}

// NewPrizesWalletFilterer creates a new log filterer instance of PrizesWallet, bound to a specific deployed contract.
func NewPrizesWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*PrizesWalletFilterer, error) {
	contract, err := bindPrizesWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletFilterer{contract: contract}, nil
}

// bindPrizesWallet binds a generic wrapper to an already deployed contract.
func bindPrizesWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PrizesWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrizesWallet *PrizesWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrizesWallet.Contract.PrizesWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrizesWallet *PrizesWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrizesWallet.Contract.PrizesWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrizesWallet *PrizesWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrizesWallet.Contract.PrizesWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrizesWallet *PrizesWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrizesWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrizesWallet *PrizesWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrizesWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrizesWallet *PrizesWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrizesWallet.Contract.contract.Transact(opts, method, params...)
}

// DonatedNfts is a free data retrieval call binding the contract method 0x9cf10d32.
//
// Solidity: function donatedNfts(uint256 ) view returns(uint256 roundNum, address nftAddress, uint256 nftId)
func (_PrizesWallet *PrizesWalletCaller) DonatedNfts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RoundNum   *big.Int
	NftAddress common.Address
	NftId      *big.Int
}, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "donatedNfts", arg0)

	outstruct := new(struct {
		RoundNum   *big.Int
		NftAddress common.Address
		NftId      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundNum = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NftAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.NftId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DonatedNfts is a free data retrieval call binding the contract method 0x9cf10d32.
//
// Solidity: function donatedNfts(uint256 ) view returns(uint256 roundNum, address nftAddress, uint256 nftId)
func (_PrizesWallet *PrizesWalletSession) DonatedNfts(arg0 *big.Int) (struct {
	RoundNum   *big.Int
	NftAddress common.Address
	NftId      *big.Int
}, error) {
	return _PrizesWallet.Contract.DonatedNfts(&_PrizesWallet.CallOpts, arg0)
}

// DonatedNfts is a free data retrieval call binding the contract method 0x9cf10d32.
//
// Solidity: function donatedNfts(uint256 ) view returns(uint256 roundNum, address nftAddress, uint256 nftId)
func (_PrizesWallet *PrizesWalletCallerSession) DonatedNfts(arg0 *big.Int) (struct {
	RoundNum   *big.Int
	NftAddress common.Address
	NftId      *big.Int
}, error) {
	return _PrizesWallet.Contract.DonatedNfts(&_PrizesWallet.CallOpts, arg0)
}

// DonatedTokens is a free data retrieval call binding the contract method 0xd7f4f8be.
//
// Solidity: function donatedTokens(uint256 ) view returns(address holder)
func (_PrizesWallet *PrizesWalletCaller) DonatedTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "donatedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DonatedTokens is a free data retrieval call binding the contract method 0xd7f4f8be.
//
// Solidity: function donatedTokens(uint256 ) view returns(address holder)
func (_PrizesWallet *PrizesWalletSession) DonatedTokens(arg0 *big.Int) (common.Address, error) {
	return _PrizesWallet.Contract.DonatedTokens(&_PrizesWallet.CallOpts, arg0)
}

// DonatedTokens is a free data retrieval call binding the contract method 0xd7f4f8be.
//
// Solidity: function donatedTokens(uint256 ) view returns(address holder)
func (_PrizesWallet *PrizesWalletCallerSession) DonatedTokens(arg0 *big.Int) (common.Address, error) {
	return _PrizesWallet.Contract.DonatedTokens(&_PrizesWallet.CallOpts, arg0)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_PrizesWallet *PrizesWalletCaller) Game(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "game")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_PrizesWallet *PrizesWalletSession) Game() (common.Address, error) {
	return _PrizesWallet.Contract.Game(&_PrizesWallet.CallOpts)
}

// Game is a free data retrieval call binding the contract method 0xc3fe3e28.
//
// Solidity: function game() view returns(address)
func (_PrizesWallet *PrizesWalletCallerSession) Game() (common.Address, error) {
	return _PrizesWallet.Contract.Game(&_PrizesWallet.CallOpts)
}

// GetDonatedTokenBalanceAmount is a free data retrieval call binding the contract method 0xe86a49d7.
//
// Solidity: function getDonatedTokenBalanceAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_PrizesWallet *PrizesWalletCaller) GetDonatedTokenBalanceAmount(opts *bind.CallOpts, roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "getDonatedTokenBalanceAmount", roundNum_, tokenAddress_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDonatedTokenBalanceAmount is a free data retrieval call binding the contract method 0xe86a49d7.
//
// Solidity: function getDonatedTokenBalanceAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_PrizesWallet *PrizesWalletSession) GetDonatedTokenBalanceAmount(roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	return _PrizesWallet.Contract.GetDonatedTokenBalanceAmount(&_PrizesWallet.CallOpts, roundNum_, tokenAddress_)
}

// GetDonatedTokenBalanceAmount is a free data retrieval call binding the contract method 0xe86a49d7.
//
// Solidity: function getDonatedTokenBalanceAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_PrizesWallet *PrizesWalletCallerSession) GetDonatedTokenBalanceAmount(roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	return _PrizesWallet.Contract.GetDonatedTokenBalanceAmount(&_PrizesWallet.CallOpts, roundNum_, tokenAddress_)
}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletCaller) GetEthBalanceInfo(opts *bind.CallOpts) (IPrizesWalletEthBalanceInfo, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "getEthBalanceInfo")

	if err != nil {
		return *new(IPrizesWalletEthBalanceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPrizesWalletEthBalanceInfo)).(*IPrizesWalletEthBalanceInfo)

	return out0, err

}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletSession) GetEthBalanceInfo() (IPrizesWalletEthBalanceInfo, error) {
	return _PrizesWallet.Contract.GetEthBalanceInfo(&_PrizesWallet.CallOpts)
}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletCallerSession) GetEthBalanceInfo() (IPrizesWalletEthBalanceInfo, error) {
	return _PrizesWallet.Contract.GetEthBalanceInfo(&_PrizesWallet.CallOpts)
}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address prizeWinnerAddress_) view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletCaller) GetEthBalanceInfo0(opts *bind.CallOpts, prizeWinnerAddress_ common.Address) (IPrizesWalletEthBalanceInfo, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "getEthBalanceInfo0", prizeWinnerAddress_)

	if err != nil {
		return *new(IPrizesWalletEthBalanceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPrizesWalletEthBalanceInfo)).(*IPrizesWalletEthBalanceInfo)

	return out0, err

}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address prizeWinnerAddress_) view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletSession) GetEthBalanceInfo0(prizeWinnerAddress_ common.Address) (IPrizesWalletEthBalanceInfo, error) {
	return _PrizesWallet.Contract.GetEthBalanceInfo0(&_PrizesWallet.CallOpts, prizeWinnerAddress_)
}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address prizeWinnerAddress_) view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletCallerSession) GetEthBalanceInfo0(prizeWinnerAddress_ common.Address) (IPrizesWalletEthBalanceInfo, error) {
	return _PrizesWallet.Contract.GetEthBalanceInfo0(&_PrizesWallet.CallOpts, prizeWinnerAddress_)
}

// MainPrizeBeneficiaryAddresses is a free data retrieval call binding the contract method 0x76e42c6a.
//
// Solidity: function mainPrizeBeneficiaryAddresses(uint256 ) view returns(address)
func (_PrizesWallet *PrizesWalletCaller) MainPrizeBeneficiaryAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "mainPrizeBeneficiaryAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MainPrizeBeneficiaryAddresses is a free data retrieval call binding the contract method 0x76e42c6a.
//
// Solidity: function mainPrizeBeneficiaryAddresses(uint256 ) view returns(address)
func (_PrizesWallet *PrizesWalletSession) MainPrizeBeneficiaryAddresses(arg0 *big.Int) (common.Address, error) {
	return _PrizesWallet.Contract.MainPrizeBeneficiaryAddresses(&_PrizesWallet.CallOpts, arg0)
}

// MainPrizeBeneficiaryAddresses is a free data retrieval call binding the contract method 0x76e42c6a.
//
// Solidity: function mainPrizeBeneficiaryAddresses(uint256 ) view returns(address)
func (_PrizesWallet *PrizesWalletCallerSession) MainPrizeBeneficiaryAddresses(arg0 *big.Int) (common.Address, error) {
	return _PrizesWallet.Contract.MainPrizeBeneficiaryAddresses(&_PrizesWallet.CallOpts, arg0)
}

// NextDonatedNftIndex is a free data retrieval call binding the contract method 0x6ff1facd.
//
// Solidity: function nextDonatedNftIndex() view returns(uint256)
func (_PrizesWallet *PrizesWalletCaller) NextDonatedNftIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "nextDonatedNftIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextDonatedNftIndex is a free data retrieval call binding the contract method 0x6ff1facd.
//
// Solidity: function nextDonatedNftIndex() view returns(uint256)
func (_PrizesWallet *PrizesWalletSession) NextDonatedNftIndex() (*big.Int, error) {
	return _PrizesWallet.Contract.NextDonatedNftIndex(&_PrizesWallet.CallOpts)
}

// NextDonatedNftIndex is a free data retrieval call binding the contract method 0x6ff1facd.
//
// Solidity: function nextDonatedNftIndex() view returns(uint256)
func (_PrizesWallet *PrizesWalletCallerSession) NextDonatedNftIndex() (*big.Int, error) {
	return _PrizesWallet.Contract.NextDonatedNftIndex(&_PrizesWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PrizesWallet *PrizesWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PrizesWallet *PrizesWalletSession) Owner() (common.Address, error) {
	return _PrizesWallet.Contract.Owner(&_PrizesWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PrizesWallet *PrizesWalletCallerSession) Owner() (common.Address, error) {
	return _PrizesWallet.Contract.Owner(&_PrizesWallet.CallOpts)
}

// RoundTimeoutTimesToWithdrawPrizes is a free data retrieval call binding the contract method 0x4b5e1b19.
//
// Solidity: function roundTimeoutTimesToWithdrawPrizes(uint256 ) view returns(uint256)
func (_PrizesWallet *PrizesWalletCaller) RoundTimeoutTimesToWithdrawPrizes(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "roundTimeoutTimesToWithdrawPrizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundTimeoutTimesToWithdrawPrizes is a free data retrieval call binding the contract method 0x4b5e1b19.
//
// Solidity: function roundTimeoutTimesToWithdrawPrizes(uint256 ) view returns(uint256)
func (_PrizesWallet *PrizesWalletSession) RoundTimeoutTimesToWithdrawPrizes(arg0 *big.Int) (*big.Int, error) {
	return _PrizesWallet.Contract.RoundTimeoutTimesToWithdrawPrizes(&_PrizesWallet.CallOpts, arg0)
}

// RoundTimeoutTimesToWithdrawPrizes is a free data retrieval call binding the contract method 0x4b5e1b19.
//
// Solidity: function roundTimeoutTimesToWithdrawPrizes(uint256 ) view returns(uint256)
func (_PrizesWallet *PrizesWalletCallerSession) RoundTimeoutTimesToWithdrawPrizes(arg0 *big.Int) (*big.Int, error) {
	return _PrizesWallet.Contract.RoundTimeoutTimesToWithdrawPrizes(&_PrizesWallet.CallOpts, arg0)
}

// TimeoutDurationToWithdrawPrizes is a free data retrieval call binding the contract method 0x6224dd3f.
//
// Solidity: function timeoutDurationToWithdrawPrizes() view returns(uint256)
func (_PrizesWallet *PrizesWalletCaller) TimeoutDurationToWithdrawPrizes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "timeoutDurationToWithdrawPrizes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeoutDurationToWithdrawPrizes is a free data retrieval call binding the contract method 0x6224dd3f.
//
// Solidity: function timeoutDurationToWithdrawPrizes() view returns(uint256)
func (_PrizesWallet *PrizesWalletSession) TimeoutDurationToWithdrawPrizes() (*big.Int, error) {
	return _PrizesWallet.Contract.TimeoutDurationToWithdrawPrizes(&_PrizesWallet.CallOpts)
}

// TimeoutDurationToWithdrawPrizes is a free data retrieval call binding the contract method 0x6224dd3f.
//
// Solidity: function timeoutDurationToWithdrawPrizes() view returns(uint256)
func (_PrizesWallet *PrizesWalletCallerSession) TimeoutDurationToWithdrawPrizes() (*big.Int, error) {
	return _PrizesWallet.Contract.TimeoutDurationToWithdrawPrizes(&_PrizesWallet.CallOpts)
}

// ClaimDonatedNft is a paid mutator transaction binding the contract method 0x94d907fc.
//
// Solidity: function claimDonatedNft(uint256 index_) returns()
func (_PrizesWallet *PrizesWalletTransactor) ClaimDonatedNft(opts *bind.TransactOpts, index_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "claimDonatedNft", index_)
}

// ClaimDonatedNft is a paid mutator transaction binding the contract method 0x94d907fc.
//
// Solidity: function claimDonatedNft(uint256 index_) returns()
func (_PrizesWallet *PrizesWalletSession) ClaimDonatedNft(index_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimDonatedNft(&_PrizesWallet.TransactOpts, index_)
}

// ClaimDonatedNft is a paid mutator transaction binding the contract method 0x94d907fc.
//
// Solidity: function claimDonatedNft(uint256 index_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) ClaimDonatedNft(index_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimDonatedNft(&_PrizesWallet.TransactOpts, index_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x9523b237.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_, uint256 amount_) returns()
func (_PrizesWallet *PrizesWalletTransactor) ClaimDonatedToken(opts *bind.TransactOpts, roundNum_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "claimDonatedToken", roundNum_, tokenAddress_, amount_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x9523b237.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_, uint256 amount_) returns()
func (_PrizesWallet *PrizesWalletSession) ClaimDonatedToken(roundNum_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimDonatedToken(&_PrizesWallet.TransactOpts, roundNum_, tokenAddress_, amount_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x9523b237.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_, uint256 amount_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) ClaimDonatedToken(roundNum_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimDonatedToken(&_PrizesWallet.TransactOpts, roundNum_, tokenAddress_, amount_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indexes_) returns()
func (_PrizesWallet *PrizesWalletTransactor) ClaimManyDonatedNfts(opts *bind.TransactOpts, indexes_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "claimManyDonatedNfts", indexes_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indexes_) returns()
func (_PrizesWallet *PrizesWalletSession) ClaimManyDonatedNfts(indexes_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimManyDonatedNfts(&_PrizesWallet.TransactOpts, indexes_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indexes_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) ClaimManyDonatedNfts(indexes_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimManyDonatedNfts(&_PrizesWallet.TransactOpts, indexes_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x5e00aaba.
//
// Solidity: function claimManyDonatedTokens((uint256,address,uint256)[] donatedTokensToClaim_) returns()
func (_PrizesWallet *PrizesWalletTransactor) ClaimManyDonatedTokens(opts *bind.TransactOpts, donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "claimManyDonatedTokens", donatedTokensToClaim_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x5e00aaba.
//
// Solidity: function claimManyDonatedTokens((uint256,address,uint256)[] donatedTokensToClaim_) returns()
func (_PrizesWallet *PrizesWalletSession) ClaimManyDonatedTokens(donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimManyDonatedTokens(&_PrizesWallet.TransactOpts, donatedTokensToClaim_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x5e00aaba.
//
// Solidity: function claimManyDonatedTokens((uint256,address,uint256)[] donatedTokensToClaim_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) ClaimManyDonatedTokens(donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimManyDonatedTokens(&_PrizesWallet.TransactOpts, donatedTokensToClaim_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x46f6b4e1.
//
// Solidity: function depositEth(uint256 roundNum_, address prizeWinnerAddress_) payable returns()
func (_PrizesWallet *PrizesWalletTransactor) DepositEth(opts *bind.TransactOpts, roundNum_ *big.Int, prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "depositEth", roundNum_, prizeWinnerAddress_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x46f6b4e1.
//
// Solidity: function depositEth(uint256 roundNum_, address prizeWinnerAddress_) payable returns()
func (_PrizesWallet *PrizesWalletSession) DepositEth(roundNum_ *big.Int, prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.DepositEth(&_PrizesWallet.TransactOpts, roundNum_, prizeWinnerAddress_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x46f6b4e1.
//
// Solidity: function depositEth(uint256 roundNum_, address prizeWinnerAddress_) payable returns()
func (_PrizesWallet *PrizesWalletTransactorSession) DepositEth(roundNum_ *big.Int, prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.DepositEth(&_PrizesWallet.TransactOpts, roundNum_, prizeWinnerAddress_)
}

// DonateNft is a paid mutator transaction binding the contract method 0xfe673fd3.
//
// Solidity: function donateNft(uint256 roundNum_, address donorAddress_, address nftAddress_, uint256 nftId_) returns()
func (_PrizesWallet *PrizesWalletTransactor) DonateNft(opts *bind.TransactOpts, roundNum_ *big.Int, donorAddress_ common.Address, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "donateNft", roundNum_, donorAddress_, nftAddress_, nftId_)
}

// DonateNft is a paid mutator transaction binding the contract method 0xfe673fd3.
//
// Solidity: function donateNft(uint256 roundNum_, address donorAddress_, address nftAddress_, uint256 nftId_) returns()
func (_PrizesWallet *PrizesWalletSession) DonateNft(roundNum_ *big.Int, donorAddress_ common.Address, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.DonateNft(&_PrizesWallet.TransactOpts, roundNum_, donorAddress_, nftAddress_, nftId_)
}

// DonateNft is a paid mutator transaction binding the contract method 0xfe673fd3.
//
// Solidity: function donateNft(uint256 roundNum_, address donorAddress_, address nftAddress_, uint256 nftId_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) DonateNft(roundNum_ *big.Int, donorAddress_ common.Address, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.DonateNft(&_PrizesWallet.TransactOpts, roundNum_, donorAddress_, nftAddress_, nftId_)
}

// DonateToken is a paid mutator transaction binding the contract method 0xe2051c7e.
//
// Solidity: function donateToken(uint256 roundNum_, address donorAddress_, address tokenAddress_, uint256 amount_) returns()
func (_PrizesWallet *PrizesWalletTransactor) DonateToken(opts *bind.TransactOpts, roundNum_ *big.Int, donorAddress_ common.Address, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "donateToken", roundNum_, donorAddress_, tokenAddress_, amount_)
}

// DonateToken is a paid mutator transaction binding the contract method 0xe2051c7e.
//
// Solidity: function donateToken(uint256 roundNum_, address donorAddress_, address tokenAddress_, uint256 amount_) returns()
func (_PrizesWallet *PrizesWalletSession) DonateToken(roundNum_ *big.Int, donorAddress_ common.Address, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.DonateToken(&_PrizesWallet.TransactOpts, roundNum_, donorAddress_, tokenAddress_, amount_)
}

// DonateToken is a paid mutator transaction binding the contract method 0xe2051c7e.
//
// Solidity: function donateToken(uint256 roundNum_, address donorAddress_, address tokenAddress_, uint256 amount_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) DonateToken(roundNum_ *big.Int, donorAddress_ common.Address, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.DonateToken(&_PrizesWallet.TransactOpts, roundNum_, donorAddress_, tokenAddress_, amount_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address mainPrizeBeneficiaryAddress_) returns(uint256)
func (_PrizesWallet *PrizesWalletTransactor) RegisterRoundEnd(opts *bind.TransactOpts, roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "registerRoundEnd", roundNum_, mainPrizeBeneficiaryAddress_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address mainPrizeBeneficiaryAddress_) returns(uint256)
func (_PrizesWallet *PrizesWalletSession) RegisterRoundEnd(roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.RegisterRoundEnd(&_PrizesWallet.TransactOpts, roundNum_, mainPrizeBeneficiaryAddress_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address mainPrizeBeneficiaryAddress_) returns(uint256)
func (_PrizesWallet *PrizesWalletTransactorSession) RegisterRoundEnd(roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.RegisterRoundEnd(&_PrizesWallet.TransactOpts, roundNum_, mainPrizeBeneficiaryAddress_)
}

// RegisterRoundEndAndDepositEthMany is a paid mutator transaction binding the contract method 0x87565d14.
//
// Solidity: function registerRoundEndAndDepositEthMany(uint256 roundNum_, address mainPrizeBeneficiaryAddress_, (address,uint256)[] ethDeposits_) payable returns(uint256)
func (_PrizesWallet *PrizesWalletTransactor) RegisterRoundEndAndDepositEthMany(opts *bind.TransactOpts, roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address, ethDeposits_ []IPrizesWalletEthDeposit) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "registerRoundEndAndDepositEthMany", roundNum_, mainPrizeBeneficiaryAddress_, ethDeposits_)
}

// RegisterRoundEndAndDepositEthMany is a paid mutator transaction binding the contract method 0x87565d14.
//
// Solidity: function registerRoundEndAndDepositEthMany(uint256 roundNum_, address mainPrizeBeneficiaryAddress_, (address,uint256)[] ethDeposits_) payable returns(uint256)
func (_PrizesWallet *PrizesWalletSession) RegisterRoundEndAndDepositEthMany(roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address, ethDeposits_ []IPrizesWalletEthDeposit) (*types.Transaction, error) {
	return _PrizesWallet.Contract.RegisterRoundEndAndDepositEthMany(&_PrizesWallet.TransactOpts, roundNum_, mainPrizeBeneficiaryAddress_, ethDeposits_)
}

// RegisterRoundEndAndDepositEthMany is a paid mutator transaction binding the contract method 0x87565d14.
//
// Solidity: function registerRoundEndAndDepositEthMany(uint256 roundNum_, address mainPrizeBeneficiaryAddress_, (address,uint256)[] ethDeposits_) payable returns(uint256)
func (_PrizesWallet *PrizesWalletTransactorSession) RegisterRoundEndAndDepositEthMany(roundNum_ *big.Int, mainPrizeBeneficiaryAddress_ common.Address, ethDeposits_ []IPrizesWalletEthDeposit) (*types.Transaction, error) {
	return _PrizesWallet.Contract.RegisterRoundEndAndDepositEthMany(&_PrizesWallet.TransactOpts, roundNum_, mainPrizeBeneficiaryAddress_, ethDeposits_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PrizesWallet *PrizesWalletTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PrizesWallet *PrizesWalletSession) RenounceOwnership() (*types.Transaction, error) {
	return _PrizesWallet.Contract.RenounceOwnership(&_PrizesWallet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PrizesWallet *PrizesWalletTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PrizesWallet.Contract.RenounceOwnership(&_PrizesWallet.TransactOpts)
}

// SetTimeoutDurationToWithdrawPrizes is a paid mutator transaction binding the contract method 0x9e2842a8.
//
// Solidity: function setTimeoutDurationToWithdrawPrizes(uint256 newValue_) returns()
func (_PrizesWallet *PrizesWalletTransactor) SetTimeoutDurationToWithdrawPrizes(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "setTimeoutDurationToWithdrawPrizes", newValue_)
}

// SetTimeoutDurationToWithdrawPrizes is a paid mutator transaction binding the contract method 0x9e2842a8.
//
// Solidity: function setTimeoutDurationToWithdrawPrizes(uint256 newValue_) returns()
func (_PrizesWallet *PrizesWalletSession) SetTimeoutDurationToWithdrawPrizes(newValue_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.SetTimeoutDurationToWithdrawPrizes(&_PrizesWallet.TransactOpts, newValue_)
}

// SetTimeoutDurationToWithdrawPrizes is a paid mutator transaction binding the contract method 0x9e2842a8.
//
// Solidity: function setTimeoutDurationToWithdrawPrizes(uint256 newValue_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) SetTimeoutDurationToWithdrawPrizes(newValue_ *big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.SetTimeoutDurationToWithdrawPrizes(&_PrizesWallet.TransactOpts, newValue_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PrizesWallet *PrizesWalletTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PrizesWallet *PrizesWalletSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.TransferOwnership(&_PrizesWallet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.TransferOwnership(&_PrizesWallet.TransactOpts, newOwner)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address prizeWinnerAddress_) returns()
func (_PrizesWallet *PrizesWalletTransactor) WithdrawEth(opts *bind.TransactOpts, prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "withdrawEth", prizeWinnerAddress_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address prizeWinnerAddress_) returns()
func (_PrizesWallet *PrizesWalletSession) WithdrawEth(prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEth(&_PrizesWallet.TransactOpts, prizeWinnerAddress_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address prizeWinnerAddress_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) WithdrawEth(prizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEth(&_PrizesWallet.TransactOpts, prizeWinnerAddress_)
}

// WithdrawEth0 is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_PrizesWallet *PrizesWalletTransactor) WithdrawEth0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "withdrawEth0")
}

// WithdrawEth0 is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_PrizesWallet *PrizesWalletSession) WithdrawEth0() (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEth0(&_PrizesWallet.TransactOpts)
}

// WithdrawEth0 is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_PrizesWallet *PrizesWalletTransactorSession) WithdrawEth0() (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEth0(&_PrizesWallet.TransactOpts)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xb4e6f803.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address,uint256)[] donatedTokensToClaim_, uint256[] donatedNftIndexes_) returns()
func (_PrizesWallet *PrizesWalletTransactor) WithdrawEverything(opts *bind.TransactOpts, withdrawEth_ bool, donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim, donatedNftIndexes_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "withdrawEverything", withdrawEth_, donatedTokensToClaim_, donatedNftIndexes_)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xb4e6f803.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address,uint256)[] donatedTokensToClaim_, uint256[] donatedNftIndexes_) returns()
func (_PrizesWallet *PrizesWalletSession) WithdrawEverything(withdrawEth_ bool, donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim, donatedNftIndexes_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEverything(&_PrizesWallet.TransactOpts, withdrawEth_, donatedTokensToClaim_, donatedNftIndexes_)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xb4e6f803.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address,uint256)[] donatedTokensToClaim_, uint256[] donatedNftIndexes_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) WithdrawEverything(withdrawEth_ bool, donatedTokensToClaim_ []IPrizesWalletDonatedTokenToClaim, donatedNftIndexes_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEverything(&_PrizesWallet.TransactOpts, withdrawEth_, donatedTokensToClaim_, donatedNftIndexes_)
}

// PrizesWalletDonatedNftClaimedIterator is returned from FilterDonatedNftClaimed and is used to iterate over the raw logs and unpacked data for DonatedNftClaimed events raised by the PrizesWallet contract.
type PrizesWalletDonatedNftClaimedIterator struct {
	Event *PrizesWalletDonatedNftClaimed // Event containing the contract specifics and raw log

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
func (it *PrizesWalletDonatedNftClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizesWalletDonatedNftClaimed)
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
		it.Event = new(PrizesWalletDonatedNftClaimed)
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
func (it *PrizesWalletDonatedNftClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizesWalletDonatedNftClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizesWalletDonatedNftClaimed represents a DonatedNftClaimed event raised by the PrizesWallet contract.
type PrizesWalletDonatedNftClaimed struct {
	RoundNum           *big.Int
	BeneficiaryAddress common.Address
	NftAddress         common.Address
	NftId              *big.Int
	Index              *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDonatedNftClaimed is a free log retrieval operation binding the contract event 0x03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3.
//
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_PrizesWallet *PrizesWalletFilterer) FilterDonatedNftClaimed(opts *bind.FilterOpts, roundNum []*big.Int, beneficiaryAddress []common.Address, nftAddress []common.Address) (*PrizesWalletDonatedNftClaimedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "DonatedNftClaimed", roundNumRule, beneficiaryAddressRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletDonatedNftClaimedIterator{contract: _PrizesWallet.contract, event: "DonatedNftClaimed", logs: logs, sub: sub}, nil
}

// WatchDonatedNftClaimed is a free log subscription operation binding the contract event 0x03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3.
//
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_PrizesWallet *PrizesWalletFilterer) WatchDonatedNftClaimed(opts *bind.WatchOpts, sink chan<- *PrizesWalletDonatedNftClaimed, roundNum []*big.Int, beneficiaryAddress []common.Address, nftAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "DonatedNftClaimed", roundNumRule, beneficiaryAddressRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizesWalletDonatedNftClaimed)
				if err := _PrizesWallet.contract.UnpackLog(event, "DonatedNftClaimed", log); err != nil {
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

// ParseDonatedNftClaimed is a log parse operation binding the contract event 0x03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3.
//
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_PrizesWallet *PrizesWalletFilterer) ParseDonatedNftClaimed(log types.Log) (*PrizesWalletDonatedNftClaimed, error) {
	event := new(PrizesWalletDonatedNftClaimed)
	if err := _PrizesWallet.contract.UnpackLog(event, "DonatedNftClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizesWalletDonatedTokenClaimedIterator is returned from FilterDonatedTokenClaimed and is used to iterate over the raw logs and unpacked data for DonatedTokenClaimed events raised by the PrizesWallet contract.
type PrizesWalletDonatedTokenClaimedIterator struct {
	Event *PrizesWalletDonatedTokenClaimed // Event containing the contract specifics and raw log

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
func (it *PrizesWalletDonatedTokenClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizesWalletDonatedTokenClaimed)
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
		it.Event = new(PrizesWalletDonatedTokenClaimed)
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
func (it *PrizesWalletDonatedTokenClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizesWalletDonatedTokenClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizesWalletDonatedTokenClaimed represents a DonatedTokenClaimed event raised by the PrizesWallet contract.
type PrizesWalletDonatedTokenClaimed struct {
	RoundNum           *big.Int
	BeneficiaryAddress common.Address
	TokenAddress       common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDonatedTokenClaimed is a free log retrieval operation binding the contract event 0xaf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0.
//
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed tokenAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) FilterDonatedTokenClaimed(opts *bind.FilterOpts, roundNum []*big.Int, beneficiaryAddress []common.Address, tokenAddress []common.Address) (*PrizesWalletDonatedTokenClaimedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "DonatedTokenClaimed", roundNumRule, beneficiaryAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletDonatedTokenClaimedIterator{contract: _PrizesWallet.contract, event: "DonatedTokenClaimed", logs: logs, sub: sub}, nil
}

// WatchDonatedTokenClaimed is a free log subscription operation binding the contract event 0xaf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0.
//
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed tokenAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) WatchDonatedTokenClaimed(opts *bind.WatchOpts, sink chan<- *PrizesWalletDonatedTokenClaimed, roundNum []*big.Int, beneficiaryAddress []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "DonatedTokenClaimed", roundNumRule, beneficiaryAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizesWalletDonatedTokenClaimed)
				if err := _PrizesWallet.contract.UnpackLog(event, "DonatedTokenClaimed", log); err != nil {
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

// ParseDonatedTokenClaimed is a log parse operation binding the contract event 0xaf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0.
//
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, address indexed tokenAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) ParseDonatedTokenClaimed(log types.Log) (*PrizesWalletDonatedTokenClaimed, error) {
	event := new(PrizesWalletDonatedTokenClaimed)
	if err := _PrizesWallet.contract.UnpackLog(event, "DonatedTokenClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizesWalletEthReceivedIterator is returned from FilterEthReceived and is used to iterate over the raw logs and unpacked data for EthReceived events raised by the PrizesWallet contract.
type PrizesWalletEthReceivedIterator struct {
	Event *PrizesWalletEthReceived // Event containing the contract specifics and raw log

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
func (it *PrizesWalletEthReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizesWalletEthReceived)
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
		it.Event = new(PrizesWalletEthReceived)
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
func (it *PrizesWalletEthReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizesWalletEthReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizesWalletEthReceived represents a EthReceived event raised by the PrizesWallet contract.
type PrizesWalletEthReceived struct {
	RoundNum           *big.Int
	PrizeWinnerAddress common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterEthReceived is a free log retrieval operation binding the contract event 0x999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326.
//
// Solidity: event EthReceived(uint256 indexed roundNum, address indexed prizeWinnerAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) FilterEthReceived(opts *bind.FilterOpts, roundNum []*big.Int, prizeWinnerAddress []common.Address) (*PrizesWalletEthReceivedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var prizeWinnerAddressRule []interface{}
	for _, prizeWinnerAddressItem := range prizeWinnerAddress {
		prizeWinnerAddressRule = append(prizeWinnerAddressRule, prizeWinnerAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "EthReceived", roundNumRule, prizeWinnerAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletEthReceivedIterator{contract: _PrizesWallet.contract, event: "EthReceived", logs: logs, sub: sub}, nil
}

// WatchEthReceived is a free log subscription operation binding the contract event 0x999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326.
//
// Solidity: event EthReceived(uint256 indexed roundNum, address indexed prizeWinnerAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) WatchEthReceived(opts *bind.WatchOpts, sink chan<- *PrizesWalletEthReceived, roundNum []*big.Int, prizeWinnerAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var prizeWinnerAddressRule []interface{}
	for _, prizeWinnerAddressItem := range prizeWinnerAddress {
		prizeWinnerAddressRule = append(prizeWinnerAddressRule, prizeWinnerAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "EthReceived", roundNumRule, prizeWinnerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizesWalletEthReceived)
				if err := _PrizesWallet.contract.UnpackLog(event, "EthReceived", log); err != nil {
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

// ParseEthReceived is a log parse operation binding the contract event 0x999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326.
//
// Solidity: event EthReceived(uint256 indexed roundNum, address indexed prizeWinnerAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) ParseEthReceived(log types.Log) (*PrizesWalletEthReceived, error) {
	event := new(PrizesWalletEthReceived)
	if err := _PrizesWallet.contract.UnpackLog(event, "EthReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizesWalletEthWithdrawnIterator is returned from FilterEthWithdrawn and is used to iterate over the raw logs and unpacked data for EthWithdrawn events raised by the PrizesWallet contract.
type PrizesWalletEthWithdrawnIterator struct {
	Event *PrizesWalletEthWithdrawn // Event containing the contract specifics and raw log

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
func (it *PrizesWalletEthWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizesWalletEthWithdrawn)
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
		it.Event = new(PrizesWalletEthWithdrawn)
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
func (it *PrizesWalletEthWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizesWalletEthWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizesWalletEthWithdrawn represents a EthWithdrawn event raised by the PrizesWallet contract.
type PrizesWalletEthWithdrawn struct {
	PrizeWinnerAddress common.Address
	BeneficiaryAddress common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawn is a free log retrieval operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed prizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) FilterEthWithdrawn(opts *bind.FilterOpts, prizeWinnerAddress []common.Address, beneficiaryAddress []common.Address) (*PrizesWalletEthWithdrawnIterator, error) {

	var prizeWinnerAddressRule []interface{}
	for _, prizeWinnerAddressItem := range prizeWinnerAddress {
		prizeWinnerAddressRule = append(prizeWinnerAddressRule, prizeWinnerAddressItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "EthWithdrawn", prizeWinnerAddressRule, beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletEthWithdrawnIterator{contract: _PrizesWallet.contract, event: "EthWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawn is a free log subscription operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed prizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) WatchEthWithdrawn(opts *bind.WatchOpts, sink chan<- *PrizesWalletEthWithdrawn, prizeWinnerAddress []common.Address, beneficiaryAddress []common.Address) (event.Subscription, error) {

	var prizeWinnerAddressRule []interface{}
	for _, prizeWinnerAddressItem := range prizeWinnerAddress {
		prizeWinnerAddressRule = append(prizeWinnerAddressRule, prizeWinnerAddressItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "EthWithdrawn", prizeWinnerAddressRule, beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizesWalletEthWithdrawn)
				if err := _PrizesWallet.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
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

// ParseEthWithdrawn is a log parse operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed prizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) ParseEthWithdrawn(log types.Log) (*PrizesWalletEthWithdrawn, error) {
	event := new(PrizesWalletEthWithdrawn)
	if err := _PrizesWallet.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizesWalletNftDonatedIterator is returned from FilterNftDonated and is used to iterate over the raw logs and unpacked data for NftDonated events raised by the PrizesWallet contract.
type PrizesWalletNftDonatedIterator struct {
	Event *PrizesWalletNftDonated // Event containing the contract specifics and raw log

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
func (it *PrizesWalletNftDonatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizesWalletNftDonated)
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
		it.Event = new(PrizesWalletNftDonated)
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
func (it *PrizesWalletNftDonatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizesWalletNftDonatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizesWalletNftDonated represents a NftDonated event raised by the PrizesWallet contract.
type PrizesWalletNftDonated struct {
	RoundNum     *big.Int
	DonorAddress common.Address
	NftAddress   common.Address
	NftId        *big.Int
	Index        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNftDonated is a free log retrieval operation binding the contract event 0xb12e72bab0c2a8fe98175a3c02792645ebdf65d37cf3983517e6d1c5ab842e23.
//
// Solidity: event NftDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_PrizesWallet *PrizesWalletFilterer) FilterNftDonated(opts *bind.FilterOpts, roundNum []*big.Int, donorAddress []common.Address, nftAddress []common.Address) (*PrizesWalletNftDonatedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "NftDonated", roundNumRule, donorAddressRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletNftDonatedIterator{contract: _PrizesWallet.contract, event: "NftDonated", logs: logs, sub: sub}, nil
}

// WatchNftDonated is a free log subscription operation binding the contract event 0xb12e72bab0c2a8fe98175a3c02792645ebdf65d37cf3983517e6d1c5ab842e23.
//
// Solidity: event NftDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_PrizesWallet *PrizesWalletFilterer) WatchNftDonated(opts *bind.WatchOpts, sink chan<- *PrizesWalletNftDonated, roundNum []*big.Int, donorAddress []common.Address, nftAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "NftDonated", roundNumRule, donorAddressRule, nftAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizesWalletNftDonated)
				if err := _PrizesWallet.contract.UnpackLog(event, "NftDonated", log); err != nil {
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

// ParseNftDonated is a log parse operation binding the contract event 0xb12e72bab0c2a8fe98175a3c02792645ebdf65d37cf3983517e6d1c5ab842e23.
//
// Solidity: event NftDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed nftAddress, uint256 nftId, uint256 index)
func (_PrizesWallet *PrizesWalletFilterer) ParseNftDonated(log types.Log) (*PrizesWalletNftDonated, error) {
	event := new(PrizesWalletNftDonated)
	if err := _PrizesWallet.contract.UnpackLog(event, "NftDonated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizesWalletOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PrizesWallet contract.
type PrizesWalletOwnershipTransferredIterator struct {
	Event *PrizesWalletOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PrizesWalletOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizesWalletOwnershipTransferred)
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
		it.Event = new(PrizesWalletOwnershipTransferred)
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
func (it *PrizesWalletOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizesWalletOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizesWalletOwnershipTransferred represents a OwnershipTransferred event raised by the PrizesWallet contract.
type PrizesWalletOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PrizesWallet *PrizesWalletFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PrizesWalletOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletOwnershipTransferredIterator{contract: _PrizesWallet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PrizesWallet *PrizesWalletFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PrizesWalletOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizesWalletOwnershipTransferred)
				if err := _PrizesWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PrizesWallet *PrizesWalletFilterer) ParseOwnershipTransferred(log types.Log) (*PrizesWalletOwnershipTransferred, error) {
	event := new(PrizesWalletOwnershipTransferred)
	if err := _PrizesWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator is returned from FilterTimeoutDurationToWithdrawPrizesChanged and is used to iterate over the raw logs and unpacked data for TimeoutDurationToWithdrawPrizesChanged events raised by the PrizesWallet contract.
type PrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator struct {
	Event *PrizesWalletTimeoutDurationToWithdrawPrizesChanged // Event containing the contract specifics and raw log

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
func (it *PrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizesWalletTimeoutDurationToWithdrawPrizesChanged)
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
		it.Event = new(PrizesWalletTimeoutDurationToWithdrawPrizesChanged)
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
func (it *PrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizesWalletTimeoutDurationToWithdrawPrizesChanged represents a TimeoutDurationToWithdrawPrizesChanged event raised by the PrizesWallet contract.
type PrizesWalletTimeoutDurationToWithdrawPrizesChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTimeoutDurationToWithdrawPrizesChanged is a free log retrieval operation binding the contract event 0x8717bb199c6bc4a5dadb21547205f9ef8ec037dda246a5526d6a6471306ea52e.
//
// Solidity: event TimeoutDurationToWithdrawPrizesChanged(uint256 newValue)
func (_PrizesWallet *PrizesWalletFilterer) FilterTimeoutDurationToWithdrawPrizesChanged(opts *bind.FilterOpts) (*PrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator, error) {

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "TimeoutDurationToWithdrawPrizesChanged")
	if err != nil {
		return nil, err
	}
	return &PrizesWalletTimeoutDurationToWithdrawPrizesChangedIterator{contract: _PrizesWallet.contract, event: "TimeoutDurationToWithdrawPrizesChanged", logs: logs, sub: sub}, nil
}

// WatchTimeoutDurationToWithdrawPrizesChanged is a free log subscription operation binding the contract event 0x8717bb199c6bc4a5dadb21547205f9ef8ec037dda246a5526d6a6471306ea52e.
//
// Solidity: event TimeoutDurationToWithdrawPrizesChanged(uint256 newValue)
func (_PrizesWallet *PrizesWalletFilterer) WatchTimeoutDurationToWithdrawPrizesChanged(opts *bind.WatchOpts, sink chan<- *PrizesWalletTimeoutDurationToWithdrawPrizesChanged) (event.Subscription, error) {

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "TimeoutDurationToWithdrawPrizesChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizesWalletTimeoutDurationToWithdrawPrizesChanged)
				if err := _PrizesWallet.contract.UnpackLog(event, "TimeoutDurationToWithdrawPrizesChanged", log); err != nil {
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

// ParseTimeoutDurationToWithdrawPrizesChanged is a log parse operation binding the contract event 0x8717bb199c6bc4a5dadb21547205f9ef8ec037dda246a5526d6a6471306ea52e.
//
// Solidity: event TimeoutDurationToWithdrawPrizesChanged(uint256 newValue)
func (_PrizesWallet *PrizesWalletFilterer) ParseTimeoutDurationToWithdrawPrizesChanged(log types.Log) (*PrizesWalletTimeoutDurationToWithdrawPrizesChanged, error) {
	event := new(PrizesWalletTimeoutDurationToWithdrawPrizesChanged)
	if err := _PrizesWallet.contract.UnpackLog(event, "TimeoutDurationToWithdrawPrizesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrizesWalletTokenDonatedIterator is returned from FilterTokenDonated and is used to iterate over the raw logs and unpacked data for TokenDonated events raised by the PrizesWallet contract.
type PrizesWalletTokenDonatedIterator struct {
	Event *PrizesWalletTokenDonated // Event containing the contract specifics and raw log

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
func (it *PrizesWalletTokenDonatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrizesWalletTokenDonated)
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
		it.Event = new(PrizesWalletTokenDonated)
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
func (it *PrizesWalletTokenDonatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrizesWalletTokenDonatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrizesWalletTokenDonated represents a TokenDonated event raised by the PrizesWallet contract.
type PrizesWalletTokenDonated struct {
	RoundNum     *big.Int
	DonorAddress common.Address
	TokenAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenDonated is a free log retrieval operation binding the contract event 0x3f94f6171c96ab491d8c268f80da5e460b148878065711a0bb099bbc297615af.
//
// Solidity: event TokenDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed tokenAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) FilterTokenDonated(opts *bind.FilterOpts, roundNum []*big.Int, donorAddress []common.Address, tokenAddress []common.Address) (*PrizesWalletTokenDonatedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "TokenDonated", roundNumRule, donorAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletTokenDonatedIterator{contract: _PrizesWallet.contract, event: "TokenDonated", logs: logs, sub: sub}, nil
}

// WatchTokenDonated is a free log subscription operation binding the contract event 0x3f94f6171c96ab491d8c268f80da5e460b148878065711a0bb099bbc297615af.
//
// Solidity: event TokenDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed tokenAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) WatchTokenDonated(opts *bind.WatchOpts, sink chan<- *PrizesWalletTokenDonated, roundNum []*big.Int, donorAddress []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "TokenDonated", roundNumRule, donorAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrizesWalletTokenDonated)
				if err := _PrizesWallet.contract.UnpackLog(event, "TokenDonated", log); err != nil {
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

// ParseTokenDonated is a log parse operation binding the contract event 0x3f94f6171c96ab491d8c268f80da5e460b148878065711a0bb099bbc297615af.
//
// Solidity: event TokenDonated(uint256 indexed roundNum, address indexed donorAddress, address indexed tokenAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) ParseTokenDonated(log types.Log) (*PrizesWalletTokenDonated, error) {
	event := new(PrizesWalletTokenDonated)
	if err := _PrizesWallet.contract.UnpackLog(event, "TokenDonated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
