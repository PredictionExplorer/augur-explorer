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

// CosmicSignatureConstantsBalanceInfo is an auto generated low-level Go binding around an user-defined struct.
type CosmicSignatureConstantsBalanceInfo struct {
	RoundNum *big.Int
	Amount   *big.Int
}

// CosmicSignatureConstantsDonatedTokenToClaim is an auto generated low-level Go binding around an user-defined struct.
type CosmicSignatureConstantsDonatedTokenToClaim struct {
	RoundNum     *big.Int
	TokenAddress common.Address
}

// IPrizesWalletMetaData contains all meta data concerning the IPrizesWallet contract.
var IPrizesWalletMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DonatedNftClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonatedTokenClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"NftDonated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"TimeoutDurationToWithdrawPrizesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenDonated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"claimDonatedNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"}],\"name\":\"claimDonatedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"indices_\",\"type\":\"uint256[]\"}],\"name\":\"claimManyDonatedNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structCosmicSignatureConstants.DonatedTokenToClaim[]\",\"name\":\"donatedTokensToClaim_\",\"type\":\"tuple[]\"}],\"name\":\"claimManyDonatedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress_\",\"type\":\"address\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"donateNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress_\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"donateToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"}],\"name\":\"getDonatedTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEthBalanceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCosmicSignatureConstants.BalanceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"getEthBalanceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCosmicSignatureConstants.BalanceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"roundMainPrizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"registerRoundEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setTimeoutDurationToWithdrawPrizes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"withdrawEth_\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structCosmicSignatureConstants.DonatedTokenToClaim[]\",\"name\":\"donatedTokensToClaim_\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"donatedNftIndices_\",\"type\":\"uint256[]\"}],\"name\":\"withdrawEverything\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetDonatedTokenAmount is a free data retrieval call binding the contract method 0x854fdf1b.
//
// Solidity: function getDonatedTokenAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_IPrizesWallet *IPrizesWalletCaller) GetDonatedTokenAmount(opts *bind.CallOpts, roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IPrizesWallet.contract.Call(opts, &out, "getDonatedTokenAmount", roundNum_, tokenAddress_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDonatedTokenAmount is a free data retrieval call binding the contract method 0x854fdf1b.
//
// Solidity: function getDonatedTokenAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_IPrizesWallet *IPrizesWalletSession) GetDonatedTokenAmount(roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	return _IPrizesWallet.Contract.GetDonatedTokenAmount(&_IPrizesWallet.CallOpts, roundNum_, tokenAddress_)
}

// GetDonatedTokenAmount is a free data retrieval call binding the contract method 0x854fdf1b.
//
// Solidity: function getDonatedTokenAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_IPrizesWallet *IPrizesWalletCallerSession) GetDonatedTokenAmount(roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	return _IPrizesWallet.Contract.GetDonatedTokenAmount(&_IPrizesWallet.CallOpts, roundNum_, tokenAddress_)
}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletCaller) GetEthBalanceInfo(opts *bind.CallOpts) (CosmicSignatureConstantsBalanceInfo, error) {
	var out []interface{}
	err := _IPrizesWallet.contract.Call(opts, &out, "getEthBalanceInfo")

	if err != nil {
		return *new(CosmicSignatureConstantsBalanceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CosmicSignatureConstantsBalanceInfo)).(*CosmicSignatureConstantsBalanceInfo)

	return out0, err

}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletSession) GetEthBalanceInfo() (CosmicSignatureConstantsBalanceInfo, error) {
	return _IPrizesWallet.Contract.GetEthBalanceInfo(&_IPrizesWallet.CallOpts)
}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletCallerSession) GetEthBalanceInfo() (CosmicSignatureConstantsBalanceInfo, error) {
	return _IPrizesWallet.Contract.GetEthBalanceInfo(&_IPrizesWallet.CallOpts)
}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address roundPrizeWinnerAddress_) view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletCaller) GetEthBalanceInfo0(opts *bind.CallOpts, roundPrizeWinnerAddress_ common.Address) (CosmicSignatureConstantsBalanceInfo, error) {
	var out []interface{}
	err := _IPrizesWallet.contract.Call(opts, &out, "getEthBalanceInfo0", roundPrizeWinnerAddress_)

	if err != nil {
		return *new(CosmicSignatureConstantsBalanceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CosmicSignatureConstantsBalanceInfo)).(*CosmicSignatureConstantsBalanceInfo)

	return out0, err

}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address roundPrizeWinnerAddress_) view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletSession) GetEthBalanceInfo0(roundPrizeWinnerAddress_ common.Address) (CosmicSignatureConstantsBalanceInfo, error) {
	return _IPrizesWallet.Contract.GetEthBalanceInfo0(&_IPrizesWallet.CallOpts, roundPrizeWinnerAddress_)
}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address roundPrizeWinnerAddress_) view returns((uint256,uint256))
func (_IPrizesWallet *IPrizesWalletCallerSession) GetEthBalanceInfo0(roundPrizeWinnerAddress_ common.Address) (CosmicSignatureConstantsBalanceInfo, error) {
	return _IPrizesWallet.Contract.GetEthBalanceInfo0(&_IPrizesWallet.CallOpts, roundPrizeWinnerAddress_)
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

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x95de6c2c.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) ClaimDonatedToken(opts *bind.TransactOpts, roundNum_ *big.Int, tokenAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "claimDonatedToken", roundNum_, tokenAddress_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x95de6c2c.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_) returns()
func (_IPrizesWallet *IPrizesWalletSession) ClaimDonatedToken(roundNum_ *big.Int, tokenAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimDonatedToken(&_IPrizesWallet.TransactOpts, roundNum_, tokenAddress_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x95de6c2c.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) ClaimDonatedToken(roundNum_ *big.Int, tokenAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimDonatedToken(&_IPrizesWallet.TransactOpts, roundNum_, tokenAddress_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indices_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) ClaimManyDonatedNfts(opts *bind.TransactOpts, indices_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "claimManyDonatedNfts", indices_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indices_) returns()
func (_IPrizesWallet *IPrizesWalletSession) ClaimManyDonatedNfts(indices_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimManyDonatedNfts(&_IPrizesWallet.TransactOpts, indices_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indices_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) ClaimManyDonatedNfts(indices_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimManyDonatedNfts(&_IPrizesWallet.TransactOpts, indices_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x50a15deb.
//
// Solidity: function claimManyDonatedTokens((uint256,address)[] donatedTokensToClaim_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) ClaimManyDonatedTokens(opts *bind.TransactOpts, donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "claimManyDonatedTokens", donatedTokensToClaim_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x50a15deb.
//
// Solidity: function claimManyDonatedTokens((uint256,address)[] donatedTokensToClaim_) returns()
func (_IPrizesWallet *IPrizesWalletSession) ClaimManyDonatedTokens(donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimManyDonatedTokens(&_IPrizesWallet.TransactOpts, donatedTokensToClaim_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x50a15deb.
//
// Solidity: function claimManyDonatedTokens((uint256,address)[] donatedTokensToClaim_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) ClaimManyDonatedTokens(donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.ClaimManyDonatedTokens(&_IPrizesWallet.TransactOpts, donatedTokensToClaim_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x46f6b4e1.
//
// Solidity: function depositEth(uint256 roundNum_, address roundPrizeWinnerAddress_) payable returns()
func (_IPrizesWallet *IPrizesWalletTransactor) DepositEth(opts *bind.TransactOpts, roundNum_ *big.Int, roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "depositEth", roundNum_, roundPrizeWinnerAddress_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x46f6b4e1.
//
// Solidity: function depositEth(uint256 roundNum_, address roundPrizeWinnerAddress_) payable returns()
func (_IPrizesWallet *IPrizesWalletSession) DepositEth(roundNum_ *big.Int, roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.DepositEth(&_IPrizesWallet.TransactOpts, roundNum_, roundPrizeWinnerAddress_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x46f6b4e1.
//
// Solidity: function depositEth(uint256 roundNum_, address roundPrizeWinnerAddress_) payable returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) DepositEth(roundNum_ *big.Int, roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.DepositEth(&_IPrizesWallet.TransactOpts, roundNum_, roundPrizeWinnerAddress_)
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
// Solidity: function registerRoundEnd(uint256 roundNum_, address roundMainPrizeWinnerAddress_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) RegisterRoundEnd(opts *bind.TransactOpts, roundNum_ *big.Int, roundMainPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "registerRoundEnd", roundNum_, roundMainPrizeWinnerAddress_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address roundMainPrizeWinnerAddress_) returns()
func (_IPrizesWallet *IPrizesWalletSession) RegisterRoundEnd(roundNum_ *big.Int, roundMainPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.RegisterRoundEnd(&_IPrizesWallet.TransactOpts, roundNum_, roundMainPrizeWinnerAddress_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address roundMainPrizeWinnerAddress_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) RegisterRoundEnd(roundNum_ *big.Int, roundMainPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.RegisterRoundEnd(&_IPrizesWallet.TransactOpts, roundNum_, roundMainPrizeWinnerAddress_)
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
// Solidity: function withdrawEth(address roundPrizeWinnerAddress_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) WithdrawEth(opts *bind.TransactOpts, roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "withdrawEth", roundPrizeWinnerAddress_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address roundPrizeWinnerAddress_) returns()
func (_IPrizesWallet *IPrizesWalletSession) WithdrawEth(roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEth(&_IPrizesWallet.TransactOpts, roundPrizeWinnerAddress_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address roundPrizeWinnerAddress_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) WithdrawEth(roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEth(&_IPrizesWallet.TransactOpts, roundPrizeWinnerAddress_)
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

// WithdrawEverything is a paid mutator transaction binding the contract method 0xa72be1b2.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address)[] donatedTokensToClaim_, uint256[] donatedNftIndices_) returns()
func (_IPrizesWallet *IPrizesWalletTransactor) WithdrawEverything(opts *bind.TransactOpts, withdrawEth_ bool, donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim, donatedNftIndices_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.contract.Transact(opts, "withdrawEverything", withdrawEth_, donatedTokensToClaim_, donatedNftIndices_)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xa72be1b2.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address)[] donatedTokensToClaim_, uint256[] donatedNftIndices_) returns()
func (_IPrizesWallet *IPrizesWalletSession) WithdrawEverything(withdrawEth_ bool, donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim, donatedNftIndices_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEverything(&_IPrizesWallet.TransactOpts, withdrawEth_, donatedTokensToClaim_, donatedNftIndices_)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xa72be1b2.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address)[] donatedTokensToClaim_, uint256[] donatedNftIndices_) returns()
func (_IPrizesWallet *IPrizesWalletTransactorSession) WithdrawEverything(withdrawEth_ bool, donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim, donatedNftIndices_ []*big.Int) (*types.Transaction, error) {
	return _IPrizesWallet.Contract.WithdrawEverything(&_IPrizesWallet.TransactOpts, withdrawEth_, donatedTokensToClaim_, donatedNftIndices_)
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
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address beneficiaryAddress, address nftAddress, uint256 nftId, uint256 index)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterDonatedNftClaimed(opts *bind.FilterOpts, roundNum []*big.Int) (*IPrizesWalletDonatedNftClaimedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "DonatedNftClaimed", roundNumRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletDonatedNftClaimedIterator{contract: _IPrizesWallet.contract, event: "DonatedNftClaimed", logs: logs, sub: sub}, nil
}

// WatchDonatedNftClaimed is a free log subscription operation binding the contract event 0x03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3.
//
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address beneficiaryAddress, address nftAddress, uint256 nftId, uint256 index)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchDonatedNftClaimed(opts *bind.WatchOpts, sink chan<- *IPrizesWalletDonatedNftClaimed, roundNum []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "DonatedNftClaimed", roundNumRule)
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
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address beneficiaryAddress, address nftAddress, uint256 nftId, uint256 index)
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
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address beneficiaryAddress, address tokenAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterDonatedTokenClaimed(opts *bind.FilterOpts, roundNum []*big.Int) (*IPrizesWalletDonatedTokenClaimedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "DonatedTokenClaimed", roundNumRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletDonatedTokenClaimedIterator{contract: _IPrizesWallet.contract, event: "DonatedTokenClaimed", logs: logs, sub: sub}, nil
}

// WatchDonatedTokenClaimed is a free log subscription operation binding the contract event 0xaf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0.
//
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address beneficiaryAddress, address tokenAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchDonatedTokenClaimed(opts *bind.WatchOpts, sink chan<- *IPrizesWalletDonatedTokenClaimed, roundNum []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "DonatedTokenClaimed", roundNumRule)
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
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address beneficiaryAddress, address tokenAddress, uint256 amount)
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
	RoundNum                *big.Int
	RoundPrizeWinnerAddress common.Address
	Amount                  *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterEthReceived is a free log retrieval operation binding the contract event 0x999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326.
//
// Solidity: event EthReceived(uint256 indexed roundNum, address indexed roundPrizeWinnerAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterEthReceived(opts *bind.FilterOpts, roundNum []*big.Int, roundPrizeWinnerAddress []common.Address) (*IPrizesWalletEthReceivedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var roundPrizeWinnerAddressRule []interface{}
	for _, roundPrizeWinnerAddressItem := range roundPrizeWinnerAddress {
		roundPrizeWinnerAddressRule = append(roundPrizeWinnerAddressRule, roundPrizeWinnerAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "EthReceived", roundNumRule, roundPrizeWinnerAddressRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletEthReceivedIterator{contract: _IPrizesWallet.contract, event: "EthReceived", logs: logs, sub: sub}, nil
}

// WatchEthReceived is a free log subscription operation binding the contract event 0x999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326.
//
// Solidity: event EthReceived(uint256 indexed roundNum, address indexed roundPrizeWinnerAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchEthReceived(opts *bind.WatchOpts, sink chan<- *IPrizesWalletEthReceived, roundNum []*big.Int, roundPrizeWinnerAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var roundPrizeWinnerAddressRule []interface{}
	for _, roundPrizeWinnerAddressItem := range roundPrizeWinnerAddress {
		roundPrizeWinnerAddressRule = append(roundPrizeWinnerAddressRule, roundPrizeWinnerAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "EthReceived", roundNumRule, roundPrizeWinnerAddressRule)
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
// Solidity: event EthReceived(uint256 indexed roundNum, address indexed roundPrizeWinnerAddress, uint256 amount)
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
	RoundPrizeWinnerAddress common.Address
	BeneficiaryAddress      common.Address
	Amount                  *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawn is a free log retrieval operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed roundPrizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) FilterEthWithdrawn(opts *bind.FilterOpts, roundPrizeWinnerAddress []common.Address, beneficiaryAddress []common.Address) (*IPrizesWalletEthWithdrawnIterator, error) {

	var roundPrizeWinnerAddressRule []interface{}
	for _, roundPrizeWinnerAddressItem := range roundPrizeWinnerAddress {
		roundPrizeWinnerAddressRule = append(roundPrizeWinnerAddressRule, roundPrizeWinnerAddressItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.FilterLogs(opts, "EthWithdrawn", roundPrizeWinnerAddressRule, beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return &IPrizesWalletEthWithdrawnIterator{contract: _IPrizesWallet.contract, event: "EthWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawn is a free log subscription operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed roundPrizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_IPrizesWallet *IPrizesWalletFilterer) WatchEthWithdrawn(opts *bind.WatchOpts, sink chan<- *IPrizesWalletEthWithdrawn, roundPrizeWinnerAddress []common.Address, beneficiaryAddress []common.Address) (event.Subscription, error) {

	var roundPrizeWinnerAddressRule []interface{}
	for _, roundPrizeWinnerAddressItem := range roundPrizeWinnerAddress {
		roundPrizeWinnerAddressRule = append(roundPrizeWinnerAddressRule, roundPrizeWinnerAddressItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _IPrizesWallet.contract.WatchLogs(opts, "EthWithdrawn", roundPrizeWinnerAddressRule, beneficiaryAddressRule)
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
// Solidity: event EthWithdrawn(address indexed roundPrizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"game_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"}],\"name\":\"CallDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DonatedNftAlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DonatedNftClaimDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"DonatedTokenClaimDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20TransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"operationAllowedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"EarlyWithdrawal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidDonatedNftIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DonatedNftClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonatedTokenClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"NftDonated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"TimeoutDurationToWithdrawPrizesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenDonated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"claimDonatedNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"}],\"name\":\"claimDonatedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"indices_\",\"type\":\"uint256[]\"}],\"name\":\"claimManyDonatedNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structCosmicSignatureConstants.DonatedTokenToClaim[]\",\"name\":\"donatedTokensToClaim_\",\"type\":\"tuple[]\"}],\"name\":\"claimManyDonatedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress_\",\"type\":\"address\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"donateNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress_\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"donateToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"donatedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"donatedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"}],\"name\":\"getDonatedTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEthBalanceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCosmicSignatureConstants.BalanceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"getEthBalanceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCosmicSignatureConstants.BalanceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numDonatedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"roundMainPrizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"registerRoundEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roundMainPrizeWinnerAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roundTimeoutTimesToWithdrawPrizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setTimeoutDurationToWithdrawPrizes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeoutDurationToWithdrawPrizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"roundPrizeWinnerAddress_\",\"type\":\"address\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"withdrawEth_\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structCosmicSignatureConstants.DonatedTokenToClaim[]\",\"name\":\"donatedTokensToClaim_\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"donatedNftIndices_\",\"type\":\"uint256[]\"}],\"name\":\"withdrawEverything\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523461002f576100196100146100f4565b6103fd565b610021610034565b612cc86105998239612cc890f35b61003a565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100669061003e565b810190811060018060401b0382111761007e57604052565b610048565b9061009661008f610034565b928361005c565b565b5f80fd5b60018060a01b031690565b6100b09061009c565b90565b6100bc816100a7565b036100c357565b5f80fd5b905051906100d4826100b3565b565b906020828203126100ef576100ec915f016100c7565b90565b610098565b6101126132618038038061010781610083565b9283398101906100d6565b90565b90565b90565b90565b61013261012d61013792610115565b61011b565b610118565b90565b610144603c61011e565b90565b610151603c61011e565b90565b634e487b7160e01b5f52601160045260245ffd5b61017761017d91939293610118565b92610118565b91610189838202610118565b92818404149015171561019857565b610154565b6101b66101a861013a565b6101b0610147565b90610168565b90565b90565b6101d06101cb6101d5926101b9565b61011b565b610118565b90565b6101e260186101bc565b90565b6101fe6101f061019d565b6101f86101d8565b90610168565b90565b90565b61021861021361021d92610201565b61011b565b610118565b90565b61023b600a6102366102306101e5565b91610204565b610168565b90565b5f1b90565b9061024f5f199161023e565b9181191691161790565b61026d61026861027292610118565b61011b565b610118565b90565b90565b9061028d61028861029492610259565b610275565b8254610243565b9055565b90565b6102af6102aa6102b492610298565b61011b565b610118565b90565b906102cc6102c76102d39261029b565b610275565b8254610243565b9055565b6102eb6102e66102f092610298565b61011b565b61009c565b90565b6102fc906102d7565b90565b60209181520190565b5f7f5a65726f2d616464726573732077617320676976656e2e000000000000000000910152565b61033c60176020926102ff565b61034581610308565b0190565b61035e9060208101905f81830391015261032f565b90565b1561036857565b610370610034565b63eac0d38960e01b81528061038760048201610349565b0390fd5b9061039c60018060a01b039161023e565b9181191691161790565b6103ba6103b56103bf9261009c565b61011b565b61009c565b90565b6103cb906103a6565b90565b6103d7906103c2565b90565b90565b906103f26103ed6103f9926103ce565b6103da565b825461038b565b9055565b6104779061040a33610479565b610424610415610220565b68010000000000000002610278565b61044b5f7c01000000000000000200000000000000000000000200000000000000036102b7565b6104708161046961046361045e5f6102f3565b6100a7565b916100a7565b1415610361565b60016103dd565b565b61048290610484565b565b61048d906104b1565b565b610498906100a7565b9052565b91906104af905f6020850194019061048f565b565b806104cc6104c66104c15f6102f3565b6100a7565b916100a7565b146104dc576104da90610539565b565b6104ff6104e85f6102f3565b5f918291631e4fbdf760e01b83526004830161049c565b0390fd5b5f1c90565b60018060a01b031690565b61051f61052491610503565b610508565b90565b6105319054610513565b90565b5f0190565b6105425f610527565b61054c825f6103dd565b9061058061057a7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936103ce565b916103ce565b91610589610034565b8061059381610534565b0390a356fe60806040526004361015610013575b610f6e565b61001d5f356101bc565b8063150b7a02146101b75780631a18c889146101b25780632299072b146101ad57806325e16063146101a857806346f6b4e1146101a35780634b5e1b191461019e57806350a15deb146101995780636224dd3f14610194578063715018a61461018f578063854fdf1b1461018a5780638da5cb5b14610185578063909466221461018057806394d907fc1461017b57806395de6c2c146101765780639cf10d32146101715780639e2842a81461016c578063a089e0be14610167578063a0ef91df14610162578063a72be1b21461015d578063c3fe3e2814610158578063cc5810d814610153578063d7f4f8be1461014e578063e2051c7e14610149578063e4a6c2a414610144578063f2fde38b1461013f5763fe673fd30361000e57610f37565b610e8f565b610e5b565b610dee565b610d73565b610cc8565b610c93565b610c4d565b610b48565b610b13565b610ae0565b610aa8565b61090f565b6108dc565b6108a7565b610847565b610811565b610782565b61074d565b610702565b610658565b610576565b610516565b6104be565b61039b565b61030d565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b60018060a01b031690565b6101e8906101d4565b90565b6101f4816101df565b036101fb57565b5f80fd5b9050359061020c826101eb565b565b90565b61021a8161020e565b0361022157565b5f80fd5b9050359061023282610211565b565b5f80fd5b5f80fd5b5f80fd5b909182601f8301121561027a5781359167ffffffffffffffff831161027557602001926001830284011161027057565b61023c565b610238565b610234565b906080828203126102da57610296815f84016101ff565b926102a482602085016101ff565b926102b28360408301610225565b92606082013567ffffffffffffffff81116102d5576102d19201610240565b9091565b6101d0565b6101cc565b63ffffffff60e01b1690565b6102f4906102df565b9052565b919061030b905f602085019401906102eb565b565b346103415761033d61032c61032336600461027f565b93929092610f76565b6103346101c2565b918291826102f8565b0390f35b6101c8565b5f91031261035057565b6101cc565b61035e9061020e565b9052565b906020806103849361037a5f8201515f860190610355565b0151910190610355565b565b9190610399905f60408501940190610362565b565b346103cb576103ab366004610346565b6103c76103b66110c0565b6103be6101c2565b91829182610386565b0390f35b6101c8565b906020828203126103e9576103e6915f01610225565b90565b6101cc565b634e487b7160e01b5f52603260045260245ffd5b506801000000000000000090565b90565b61041c81610402565b8210156104365761042e600191610410565b910201905f90565b6103ee565b1c90565b60018060a01b031690565b61045a90600861045f930261043b565b61043f565b90565b9061046d915461044a565b90565b600261047b81610402565b821015610498576104959161048f91610413565b90610462565b90565b5f80fd5b6104a5906101df565b9052565b91906104bc905f6020850194019061049c565b565b346104ee576104ea6104d96104d43660046103d0565b610470565b6104e16101c2565b918291826104a9565b0390f35b6101c8565b9060208282031261050c57610509915f016101ff565b90565b6101cc565b5f0190565b346105445761052e6105293660046104f3565b611376565b6105366101c2565b8061054081610511565b0390f35b6101c8565b9190604083820312610571578061056561056e925f8601610225565b936020016101ff565b90565b6101cc565b61058a610584366004610549565b906116f4565b6105926101c2565b8061059c81610511565b0390f35b506801000000000000000090565b90565b6105ba816105a0565b8210156105d4576105cc6001916105ae565b910201905f90565b6103ee565b90565b6105ec9060086105f1930261043b565b6105d9565b90565b906105ff91546105dc565b90565b68010000000000000003610615816105a0565b8210156106325761062f91610629916105b1565b906105f4565b90565b5f80fd5b61063f9061020e565b9052565b9190610656905f60208501940190610636565b565b346106885761068461067361066e3660046103d0565b610602565b61067b6101c2565b91829182610643565b0390f35b6101c8565b909182601f830112156106c75781359167ffffffffffffffff83116106c25760200192604083028401116106bd57565b61023c565b610238565b610234565b906020828203126106fd575f82013567ffffffffffffffff81116106f8576106f4920161068d565b9091565b6101d0565b6101cc565b346107315761071b6107153660046106cc565b90611742565b6107236101c2565b8061072d81610511565b0390f35b6101c8565b61074a680100000000000000025f906105f4565b90565b3461077d5761075d366004610346565b610779610768610736565b6107706101c2565b91829182610643565b0390f35b6101c8565b346107b057610792366004610346565b61079a611808565b6107a26101c2565b806107ac81610511565b0390f35b6101c8565b6107be906101df565b90565b6107ca816107b5565b036107d157565b5f80fd5b905035906107e2826107c1565b565b919060408382031261080c5780610800610809925f8601610225565b936020016107d5565b90565b6101cc565b346108425761083e61082d6108273660046107e4565b90611812565b6108356101c2565b91829182610643565b0390f35b6101c8565b3461087757610857366004610346565b610873610862611855565b61086a6101c2565b918291826104a9565b0390f35b6101c8565b6108a47c01000000000000000200000000000000000000000200000000000000035f906105f4565b90565b346108d7576108b7366004610346565b6108d36108c261087c565b6108ca6101c2565b91829182610643565b0390f35b6101c8565b3461090a576108f46108ef3660046103d0565b611bdd565b6108fc6101c2565b8061090681610511565b0390f35b6101c8565b3461093e576109286109223660046107e4565b90612139565b6109306101c2565b8061093a81610511565b0390f35b6101c8565b506801000000000000000090565b90565b61095d81610943565b8210156109775761096f600391610951565b910201905f90565b6103ee565b5f1c90565b61098d6109929161097c565b6105d9565b90565b61099f9054610981565b90565b60018060a01b031690565b6109b96109be9161097c565b6109a2565b90565b6109cb90546109ad565b90565b7c0100000000000000020000000000000000000000020000000000000004906109f682610943565b811015610a2e57610a0691610954565b50610a125f8201610995565b91610a2b6002610a24600185016109c1565b9301610995565b90565b5f80fd5b90565b610a49610a44610a4e926101d4565b610a32565b6101d4565b90565b610a5a90610a35565b90565b610a6690610a51565b90565b610a7290610a5d565b9052565b604090610a9f610aa69496959396610a9560608401985f850190610636565b6020830190610a69565b0190610636565b565b34610adb57610ad7610ac3610abe3660046103d0565b6109ce565b610ace9391936101c2565b93849384610a76565b0390f35b6101c8565b34610b0e57610af8610af33660046103d0565b61236c565b610b006101c2565b80610b0a81610511565b0390f35b6101c8565b34610b4357610b3f610b2e610b293660046104f3565b612377565b610b366101c2565b91829182610386565b0390f35b6101c8565b34610b7657610b58366004610346565b610b606123a9565b610b686101c2565b80610b7281610511565b0390f35b6101c8565b151590565b610b8981610b7b565b03610b9057565b5f80fd5b90503590610ba182610b80565b565b909182601f83011215610bdd5781359167ffffffffffffffff8311610bd8576020019260208302840111610bd357565b61023c565b610238565b610234565b606081830312610c4857610bf8825f8301610b94565b92602082013567ffffffffffffffff8111610c435783610c1991840161068d565b929093604082013567ffffffffffffffff8111610c3e57610c3a9201610ba3565b9091565b6101d0565b6101d0565b6101cc565b34610c7f57610c69610c60366004610be2565b93929092612474565b610c716101c2565b80610c7b81610511565b0390f35b6101c8565b610c9060015f90610462565b90565b34610cc357610ca3366004610346565b610cbf610cae610c84565b610cb66101c2565b918291826104a9565b0390f35b6101c8565b34610cf757610ce1610cdb366004610549565b90612545565b610ce96101c2565b80610cf381610511565b0390f35b6101c8565b50600160e01b90565b90565b610d1181610cfc565b821015610d2b57610d23600191610d05565b910201905f90565b6103ee565b7402000000000000000000000002000000000000000390610d5082610cfc565b811015610d6f57610d655f91610d6c93610d08565b5001610995565b90565b5f80fd5b34610da357610d9f610d8e610d893660046103d0565b610d30565b610d966101c2565b91829182610643565b0390f35b6101c8565b608081830312610de957610dbe825f8301610225565b92610de6610dcf84602085016101ff565b93610ddd81604086016107d5565b93606001610225565b90565b6101cc565b34610e2057610e0a610e01366004610da8565b9291909161279a565b610e126101c2565b80610e1c81610511565b0390f35b6101c8565b90602082820312610e56575f82013567ffffffffffffffff8111610e5157610e4d9201610ba3565b9091565b6101d0565b6101cc565b34610e8a57610e74610e6e366004610e25565b906127c1565b610e7c6101c2565b80610e8681610511565b0390f35b6101c8565b34610ebd57610ea7610ea23660046104f3565b612886565b610eaf6101c2565b80610eb981610511565b0390f35b6101c8565b610ecb906101df565b90565b610ed781610ec2565b03610ede57565b5f80fd5b90503590610eef82610ece565b565b608081830312610f3257610f07825f8301610225565b92610f2f610f1884602085016101ff565b93610f268160408601610ee2565b93606001610225565b90565b6101cc565b34610f6957610f53610f4a366004610ef1565b92919091612b08565b610f5b6101c2565b80610f6581610511565b0390f35b6101c8565b5f80fd5b5f90565b5050505050610f83610f72565b50630a85bd0160e11b90565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b90610fb790610f8f565b810190811067ffffffffffffffff821117610fd157604052565b610f99565b90610fe9610fe26101c2565b9283610fad565b565b610ff56040610fd6565b90565b5f90565b611004610feb565b9060208083611011610ff8565b81520161101c610ff8565b81525050565b61102a610ffc565b90565b61103690610a35565b90565b50600160a01b90565b90565b61104e81611039565b82101561106857611060600291611042565b910201905f90565b6103ee565b906110779061020e565b9052565b906110b26110a9600161108c610feb565b946110a361109b5f8301610995565b5f880161106d565b01610995565b6020840161106d565b565b6110bd9061107b565b90565b6110c8611022565b506110ee6110e8680200000000000000036110e23361102d565b90611045565b506110b4565b90565b90565b90565b61110b611106611110926110f4565b610a32565b61020e565b90565b60209181520190565b5f7f4e6f7420656e6f7567682074696d652068617320656c61707365642e00000000910152565b611150601c602092611113565b6111598161111c565b0190565b91604061118e92949361118761117c606083018381035f850152611143565b966020830190610636565b0190610636565b565b15611199575050565b6111ba6111a46101c2565b92839263496fca8f60e11b84526004840161115d565b0390fd5b1b90565b919060086111dd9102916111d75f19846111be565b926111be565b9181191691161790565b6111fb6111f66112009261020e565b610a32565b61020e565b90565b90565b919061121c611217611224936111e7565b611203565b9083546111c2565b9055565b5f90565b61123e91611238611228565b91611206565b565b61124990610a51565b90565b905090565b61125c5f809261124c565b0190565b61126990611251565b90565b67ffffffffffffffff811161128a57611286602091610f8f565b0190565b610f99565b906112a161129c8361126c565b610fd6565b918252565b606090565b3d5f146112c6576112bb3d61128f565b903d5f602084013e5b565b6112ce6112a6565b906112c4565b5f7f455448207769746864726177616c206661696c65642e00000000000000000000910152565b6113086016602092611113565b611311816112d4565b0190565b91604061134692949361133f611334606083018381035f8501526112fb565b96602083019061049c565b0190610636565b565b15611351575050565b61137261135c6101c2565b928392630aa7db6360e11b845260048401611315565b0390fd5b61148d9061139f611399680200000000000000036113938461102d565b90611045565b506110f1565b906113ed6113ca6113c4680100000000000000036113be5f8701610995565b906105b1565b906105f4565b426113dd6113d78361020e565b9161020e565b1015908161148f575b4291611190565b6114125f806113fe60018601610995565b9461140c826001830161122c565b0161122c565b33829161145d61144b6114457f4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f5929238393611240565b93611240565b936114546101c2565b91829182610643565b0390a35f80338361146c6101c2565b908161147781611260565b03925af16114836112ab565b5090339091611348565b565b8091506114a461149e5f6110f7565b9161020e565b11906113e6565b6114b76114bc9161097c565b61043f565b90565b6114c990546114ab565b90565b60407f6d6574686f642e00000000000000000000000000000000000000000000000000917f4f6e6c792074686520436f736d69635369676e617475726547616d6520636f6e5f8201527f7472616374206973207065726d697474656420746f2063616c6c20746869732060208201520152565b61154c6047606092611113565b611555816114cc565b0190565b919061157c906020611574604086018681035f88015261153f565b94019061049c565b565b156115865750565b6115a8906115926101c2565b91829163336e7de360e11b835260048301611559565b0390fd5b906115dd916115d8336115d06115ca6115c560016114bf565b6101df565b916101df565b14339061157e565b611653565b565b5f1b90565b906115f05f19916115df565b9181191691161790565b9061160f61160a611616926111e7565b611203565b82546115e4565b9055565b634e487b7160e01b5f52601160045260245ffd5b61163d6116439193929361020e565b9261020e565b820180921161164e57565b61161a565b6116a461167b6116756802000000000000000361166f8661102d565b90611045565b506110f1565b611687835f83016115fa565b61169e60013492019161169983610995565b61162e565b906115fa565b9034916116ef6116dd6116d77f999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326936111e7565b93611240565b936116e66101c2565b91829182610643565b0390a3565b906116fe916115ac565b565b600161170c910161020e565b90565b5090565b9190811015611723576040020190565b6103ee565b3561173281610211565b90565b3561173f816107c1565b90565b919061174d5f6110f7565b5b8061176b61176561176087869061170f565b61020e565b9161020e565b10156117b5576117b0906117ab61178e5f61178888878691611713565b01611728565b6117a5602061179f89888791611713565b01611735565b90612139565b611700565b61174e565b50509050565b6117c3612b16565b6117cb6117f5565b565b6117e16117dc6117e6926110f4565b610a32565b6101d4565b90565b6117f2906117cd565b90565b6118066118015f6117e9565b612b84565b565b6118106117bb565b565b61184761182c61184e935f93611826611228565b50612c4b565b74020000000000000000000000020000000000000003610d08565b5001610995565b90565b5f90565b61185d611851565b506118675f6114bf565b90565b90565b9061187790610ec2565b9052565b6118856060610fd6565b90565b906118d76118ce600261189961187b565b946118b06118a85f8301610995565b5f880161106d565b6118c86118bf600183016109c1565b6020880161186d565b01610995565b6040840161106d565b565b6118e290611888565b90565b6118ef9051610ec2565b90565b5f7f446f6e61746564204e465420616c726561647920636c61696d65642e00000000910152565b611926601c602092611113565b61192f816118f2565b0190565b919061195690602061194e604086018681035f880152611919565b940190610636565b565b5f7f496e76616c696420646f6e61746564204e465420696e6465782e000000000000910152565b61198c601a602092611113565b61199581611958565b0190565b91906119bc9060206119b4604086018681035f88015261197f565b940190610636565b565b6119c8905161020e565b90565b60407f686973204e46542e000000000000000000000000000000000000000000000000917f4f6e6c79207468652062696464696e6720726f756e64206d61696e207072697a5f8201527f652077696e6e6572206973207065726d697474656420746f20636c61696d207460208201520152565b611a4b6048606092611113565b611a54816119cb565b0190565b916040611a89929493611a82611a77606083018381035f850152611a3e565b96602083019061049c565b0190610636565b565b15611a94575050565b611ab5611a9f6101c2565b92839263d1b9b13b60e01b845260048401611a58565b0390fd5b91906008611ad9910291611ad360018060a01b03846111be565b926111be565b9181191691161790565b611aec90610a35565b90565b611af890611ae3565b90565b90565b9190611b14611b0f611b1c93611aef565b611afb565b908354611ab9565b9055565b5f90565b611b3691611b30611b20565b91611afe565b565b611b6d611b7494611b63606094989795611b59608086019a5f87019061049c565b6020850190610a69565b6040830190610636565b0190610636565b565b611b7f90610a51565b90565b5f80fd5b60e01b90565b5f910312611b9657565b6101cc565b604090611bc4611bcb9496959396611bba60608401985f85019061049c565b602083019061049c565b0190610636565b565b611bd56101c2565b3d5f823e3d90fd5b90611c0f611c097c01000000000000000200000000000000000000000200000000000000048490610954565b5061186a565b611c18816118d9565b90611c2d611c28602084016118e5565b610a5d565b611c47611c41611c3c5f6117e9565b6101df565b916101df565b145f14611cdb578380611c87611c81611c7c7c0100000000000000020000000000000000000000020000000000000003610995565b61020e565b9161020e565b10155f14611cb557611cb190611c9b6101c2565b9182916306420ad160e41b835260048301611999565b0390fd5b611cd790611cc16101c2565b91829163aa576f4960e01b835260048301611933565b0390fd5b5f6002611d3b929593949533611d16611d10611d0b611d0586611cff898d016119be565b90610413565b90610462565b6101df565b916101df565b03611e4d575b611d288380830161122c565b611d358360018301611b24565b0161122c565b611d465f83016119be565b33611d53602085016118e5565b90611da0611d63604087016119be565b94611d8e7f03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3956111e7565b95611d976101c2565b94859485611b38565b0390a2611db7611db2602083016118e5565b610a5d565b6342842e0e90611dc630611b76565b90611dd460403395016119be565b92813b15611e48575f611dfa91611e058296611dee6101c2565b98899788968795611b86565b855260048501611b9b565b03925af18015611e4357611e17575b50565b611e36905f3d8111611e3c575b611e2e8183610fad565b810190611b8c565b5f611e14565b503d611e24565b611bcd565b611b82565b611e9c611e77611e7168010000000000000003611e6b878b016119be565b906105b1565b906105f4565b42611e8a611e848361020e565b9161020e565b10159081611ea1575b50338791611a8b565b611d1c565b9050611eb5611eaf866110f7565b9161020e565b115f611e93565b60407f686973204552432d323020746f6b656e20646f6e6174696f6e2e000000000000917f4f6e6c79207468652062696464696e6720726f756e64206d61696e207072697a5f8201527f652077696e6e6572206973207065726d697474656420746f20636c61696d207460208201520152565b611f3c605a606092611113565b611f4581611ebc565b0190565b611f5290610a51565b90565b611f5e90611f49565b9052565b606090611f9b611fa29496959396611f91611f86608085018581035f870152611f2f565b986020850190610636565b604083019061049c565b0190611f55565b565b9290919215611fb257505050565b611fd490611fbe6101c2565b93849363b039722960e01b855260048501611f62565b0390fd5b90565b611fe56020610fd6565b90565b90612007611fff5f611ff8611fdb565b9401610995565b5f840161106d565b565b61201290611fe8565b90565b60409061203e612045949695939661203460608401985f85019061049c565b6020830190611f55565b0190610636565b565b9050519061205482610b80565b565b9060208282031261206f5761206c915f01612047565b90565b6101cc565b91602061209592949361208e60408201965f83019061049c565b0190610636565b565b5f7f5472616e73666572206661696c65642e00000000000000000000000000000000910152565b6120cb6010602092611113565b6120d481612097565b0190565b9160406121099294936121026120f7606083018381035f8501526120be565b96602083019061049c565b0190610636565b565b15612114575050565b61213561211f6101c2565b92839263f7fce64560e01b8452600484016120d8565b0390fd5b90612209903361216561215f61215a61215460028890610413565b90610462565b6101df565b916101df565b036122a4575b61219d61219761217c858490612c4b565b74020000000000000000000000020000000000000003610d08565b50611fd8565b926121b35f806121ac87612009565b960161122c565b3390826121c15f87016119be565b926122016121ef7faf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0946111e7565b946121f86101c2565b93849384612015565b0390a2611f49565b90602063a9059cbb92339061223b5f6122238187016119be565b9661224661222f6101c2565b98899687958694611b86565b845260048401612074565b03925af1801561229f5761226f925f91612271575b50906122695f3392016119be565b9161210b565b565b612292915060203d8111612298575b61228a8183610fad565b810190612056565b5f61225b565b503d612280565b611bcd565b6122eb6122c46122be6801000000000000000386906105b1565b906105f4565b426122d76122d18361020e565b9161020e565b101590816122f0575b508433908492611fa4565b61216b565b90506123046122fe5f6110f7565b9161020e565b115f6122e0565b61231c90612317612b16565b61231e565b565b61233181680100000000000000026115fa565b6123677f8717bb199c6bc4a5dadb21547205f9ef8ec037dda246a5526d6a6471306ea52e9161235e6101c2565b91829182610643565b0390a1565b6123759061230b565b565b6123a06123a691612386611022565b5061239a680200000000000000039161102d565b90611045565b506110b4565b90565b6124726123d16123cb680200000000000000036123c53361102d565b90611045565b506110f1565b6123f65f806123e260018501610995565b936123f0826001830161122c565b0161122c565b3333829161244261243061242a7f4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f5929238393611240565b93611240565b936124396101c2565b91829182610643565b0390a35f8033836124516101c2565b908161245c81611260565b03925af16124686112ab565b5090339091611348565b565b916124899161248e959493612490575b611742565b6127c1565b565b6124986123a9565b612484565b906124ce916124c9336124c16124bb6124b660016114bf565b6101df565b916101df565b14339061157e565b6124f5565b565b90565b91906124e96124e46124f193611240565b6124d0565b908354611ab9565b9055565b61250e6125439261250860028490610413565b906124d3565b61253d61252d4261252768010000000000000002610995565b9061162e565b91680100000000000000036105b1565b90611206565b565b9061254f9161249d565b565b9061258493929161257f3361257761257161256c60016114bf565b6101df565b916101df565b14339061157e565b612621565b565b5f7f5a65726f2d616464726573732077617320676976656e2e000000000000000000910152565b6125ba6017602092611113565b6125c381612586565b0190565b6125dc9060208101905f8183039101526125ad565b90565b156125e657565b6125ee6101c2565b63eac0d38960e01b815280612605600482016125c7565b0390fd5b61261290610a35565b90565b61261e90612609565b90565b60209193926127069161265761263683611f49565b61265061264a6126455f6117e9565b6101df565b916101df565b14156125df565b6126a861268c61268661266b848690612c4b565b74020000000000000000000000020000000000000003610d08565b50611fd8565b6126a25f8892019161269d83610995565b61162e565b906115fa565b85829086926126fe6126ec6126e66126e07f3f94f6171c96ab491d8c268f80da5e460b148878065711a0bb099bbc297615af946111e7565b94611240565b94612615565b946126f56101c2565b91829182610643565b0390a4611f49565b6323b872dd93906127335f61271a30611b76565b9661273e876127276101c2565b998a9788968795611b86565b855260048501611b9b565b03925af1801561279557612765925f91612767575b509061275e30611b76565b909161210b565b565b612788915060203d811161278e575b6127808183610fad565b810190612056565b5f612753565b503d612776565b611bcd565b906127a6939291612551565b565b5090565b91908110156127bc576020020190565b6103ee565b9190916127cd5f6110f7565b5b806127eb6127e56127e08588906127a8565b61020e565b9161020e565b101561281b576128169061281161280c612807858885916127ac565b611728565b611bdd565b611700565b6127ce565b50509050565b6128329061282d612b16565b612834565b565b8061284f6128496128445f6117e9565b6101df565b916101df565b1461285f5761285d90612b84565b565b61288261286b5f6117e9565b5f918291631e4fbdf760e01b8352600483016104a9565b0390fd5b61288f90612821565b565b906128c49392916128bf336128b76128b16128ac60016114bf565b6101df565b916101df565b14339061157e565b612940565b565b906128d760018060a01b03916115df565b9181191691161790565b906128f66128f16128fd92611aef565b611afb565b82546128c6565b9055565b91602061292292949361291b60408201965f830190610636565b0190610636565b565b61292d9061020e565b5f19811461293b5760010190565b61161a565b91612a79612a56612a7e939461297961295885610a5d565b61297261296c6129675f6117e9565b6101df565b916101df565b14156125df565b61299f7c0100000000000000020000000000000000000000020000000000000003610995565b906129f76129d46129ce7c01000000000000000200000000000000000000000200000000000000048590610954565b5061186a565b6129e0835f83016115fa565b6129ed87600183016128e1565b60028a91016115fa565b868590899284612a39612a33612a2d7fb12e72bab0c2a8fe98175a3c02792645ebdf65d37cf3983517e6d1c5ab842e23946111e7565b94611240565b94611aef565b94612a4e612a456101c2565b92839283612901565b0390a4612924565b7c01000000000000000200000000000000000000000200000000000000036115fa565b610a5d565b6342842e0e9190612a8e30611b76565b9392813b15612b03575f612ab591612ac08296612aa96101c2565b98899788968795611b86565b855260048501611b9b565b03925af18015612afe57612ad2575b50565b612af1905f3d8111612af7575b612ae98183610fad565b810190611b8c565b5f612acf565b503d612adf565b611bcd565b611b82565b90612b14939291612891565b565b612b1e611855565b612b37612b31612b2c612c85565b6101df565b916101df565b03612b3e57565b612b60612b49612c85565b5f91829163118cdaa760e01b8352600483016104a9565b0390fd5b90612b79612b74612b8092611240565b6124d0565b82546128c6565b9055565b612b8d5f6114bf565b612b97825f612b64565b90612bcb612bc57f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093611240565b91611240565b91612bd46101c2565b80612bde81610511565b0390a3565b612bf7612bf2612bfc926101d4565b610a32565b61020e565b90565b90565b60ff1690565b612c1c612c17612c2192612bff565b610a32565b612c02565b90565b612c4390612c3d612c37612c4894612c02565b9161020e565b906111be565b61020e565b90565b612c71612c6c612c67612c8193612c60611228565b5094611f49565b61102d565b612be3565b612c7b6040612c08565b90612c24565b1790565b612c8d611851565b50339056fea26469706673582212202b671a0571a883476a42da499921690086ac4b8ddf6dd6c6bb5653419aac1db164736f6c634300081b0033",
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
// Solidity: function donatedTokens(uint256 ) view returns(uint256 amount)
func (_PrizesWallet *PrizesWalletCaller) DonatedTokens(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "donatedTokens", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DonatedTokens is a free data retrieval call binding the contract method 0xd7f4f8be.
//
// Solidity: function donatedTokens(uint256 ) view returns(uint256 amount)
func (_PrizesWallet *PrizesWalletSession) DonatedTokens(arg0 *big.Int) (*big.Int, error) {
	return _PrizesWallet.Contract.DonatedTokens(&_PrizesWallet.CallOpts, arg0)
}

// DonatedTokens is a free data retrieval call binding the contract method 0xd7f4f8be.
//
// Solidity: function donatedTokens(uint256 ) view returns(uint256 amount)
func (_PrizesWallet *PrizesWalletCallerSession) DonatedTokens(arg0 *big.Int) (*big.Int, error) {
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

// GetDonatedTokenAmount is a free data retrieval call binding the contract method 0x854fdf1b.
//
// Solidity: function getDonatedTokenAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_PrizesWallet *PrizesWalletCaller) GetDonatedTokenAmount(opts *bind.CallOpts, roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "getDonatedTokenAmount", roundNum_, tokenAddress_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDonatedTokenAmount is a free data retrieval call binding the contract method 0x854fdf1b.
//
// Solidity: function getDonatedTokenAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_PrizesWallet *PrizesWalletSession) GetDonatedTokenAmount(roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	return _PrizesWallet.Contract.GetDonatedTokenAmount(&_PrizesWallet.CallOpts, roundNum_, tokenAddress_)
}

// GetDonatedTokenAmount is a free data retrieval call binding the contract method 0x854fdf1b.
//
// Solidity: function getDonatedTokenAmount(uint256 roundNum_, address tokenAddress_) view returns(uint256)
func (_PrizesWallet *PrizesWalletCallerSession) GetDonatedTokenAmount(roundNum_ *big.Int, tokenAddress_ common.Address) (*big.Int, error) {
	return _PrizesWallet.Contract.GetDonatedTokenAmount(&_PrizesWallet.CallOpts, roundNum_, tokenAddress_)
}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletCaller) GetEthBalanceInfo(opts *bind.CallOpts) (CosmicSignatureConstantsBalanceInfo, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "getEthBalanceInfo")

	if err != nil {
		return *new(CosmicSignatureConstantsBalanceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CosmicSignatureConstantsBalanceInfo)).(*CosmicSignatureConstantsBalanceInfo)

	return out0, err

}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletSession) GetEthBalanceInfo() (CosmicSignatureConstantsBalanceInfo, error) {
	return _PrizesWallet.Contract.GetEthBalanceInfo(&_PrizesWallet.CallOpts)
}

// GetEthBalanceInfo is a free data retrieval call binding the contract method 0x1a18c889.
//
// Solidity: function getEthBalanceInfo() view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletCallerSession) GetEthBalanceInfo() (CosmicSignatureConstantsBalanceInfo, error) {
	return _PrizesWallet.Contract.GetEthBalanceInfo(&_PrizesWallet.CallOpts)
}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address roundPrizeWinnerAddress_) view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletCaller) GetEthBalanceInfo0(opts *bind.CallOpts, roundPrizeWinnerAddress_ common.Address) (CosmicSignatureConstantsBalanceInfo, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "getEthBalanceInfo0", roundPrizeWinnerAddress_)

	if err != nil {
		return *new(CosmicSignatureConstantsBalanceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CosmicSignatureConstantsBalanceInfo)).(*CosmicSignatureConstantsBalanceInfo)

	return out0, err

}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address roundPrizeWinnerAddress_) view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletSession) GetEthBalanceInfo0(roundPrizeWinnerAddress_ common.Address) (CosmicSignatureConstantsBalanceInfo, error) {
	return _PrizesWallet.Contract.GetEthBalanceInfo0(&_PrizesWallet.CallOpts, roundPrizeWinnerAddress_)
}

// GetEthBalanceInfo0 is a free data retrieval call binding the contract method 0xa089e0be.
//
// Solidity: function getEthBalanceInfo(address roundPrizeWinnerAddress_) view returns((uint256,uint256))
func (_PrizesWallet *PrizesWalletCallerSession) GetEthBalanceInfo0(roundPrizeWinnerAddress_ common.Address) (CosmicSignatureConstantsBalanceInfo, error) {
	return _PrizesWallet.Contract.GetEthBalanceInfo0(&_PrizesWallet.CallOpts, roundPrizeWinnerAddress_)
}

// NumDonatedNfts is a free data retrieval call binding the contract method 0x90946622.
//
// Solidity: function numDonatedNfts() view returns(uint256)
func (_PrizesWallet *PrizesWalletCaller) NumDonatedNfts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "numDonatedNfts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumDonatedNfts is a free data retrieval call binding the contract method 0x90946622.
//
// Solidity: function numDonatedNfts() view returns(uint256)
func (_PrizesWallet *PrizesWalletSession) NumDonatedNfts() (*big.Int, error) {
	return _PrizesWallet.Contract.NumDonatedNfts(&_PrizesWallet.CallOpts)
}

// NumDonatedNfts is a free data retrieval call binding the contract method 0x90946622.
//
// Solidity: function numDonatedNfts() view returns(uint256)
func (_PrizesWallet *PrizesWalletCallerSession) NumDonatedNfts() (*big.Int, error) {
	return _PrizesWallet.Contract.NumDonatedNfts(&_PrizesWallet.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_PrizesWallet *PrizesWalletCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_PrizesWallet *PrizesWalletSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _PrizesWallet.Contract.OnERC721Received(&_PrizesWallet.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_PrizesWallet *PrizesWalletCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _PrizesWallet.Contract.OnERC721Received(&_PrizesWallet.CallOpts, arg0, arg1, arg2, arg3)
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

// RoundMainPrizeWinnerAddresses is a free data retrieval call binding the contract method 0x2299072b.
//
// Solidity: function roundMainPrizeWinnerAddresses(uint256 ) view returns(address)
func (_PrizesWallet *PrizesWalletCaller) RoundMainPrizeWinnerAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PrizesWallet.contract.Call(opts, &out, "roundMainPrizeWinnerAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoundMainPrizeWinnerAddresses is a free data retrieval call binding the contract method 0x2299072b.
//
// Solidity: function roundMainPrizeWinnerAddresses(uint256 ) view returns(address)
func (_PrizesWallet *PrizesWalletSession) RoundMainPrizeWinnerAddresses(arg0 *big.Int) (common.Address, error) {
	return _PrizesWallet.Contract.RoundMainPrizeWinnerAddresses(&_PrizesWallet.CallOpts, arg0)
}

// RoundMainPrizeWinnerAddresses is a free data retrieval call binding the contract method 0x2299072b.
//
// Solidity: function roundMainPrizeWinnerAddresses(uint256 ) view returns(address)
func (_PrizesWallet *PrizesWalletCallerSession) RoundMainPrizeWinnerAddresses(arg0 *big.Int) (common.Address, error) {
	return _PrizesWallet.Contract.RoundMainPrizeWinnerAddresses(&_PrizesWallet.CallOpts, arg0)
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

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x95de6c2c.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_) returns()
func (_PrizesWallet *PrizesWalletTransactor) ClaimDonatedToken(opts *bind.TransactOpts, roundNum_ *big.Int, tokenAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "claimDonatedToken", roundNum_, tokenAddress_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x95de6c2c.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_) returns()
func (_PrizesWallet *PrizesWalletSession) ClaimDonatedToken(roundNum_ *big.Int, tokenAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimDonatedToken(&_PrizesWallet.TransactOpts, roundNum_, tokenAddress_)
}

// ClaimDonatedToken is a paid mutator transaction binding the contract method 0x95de6c2c.
//
// Solidity: function claimDonatedToken(uint256 roundNum_, address tokenAddress_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) ClaimDonatedToken(roundNum_ *big.Int, tokenAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimDonatedToken(&_PrizesWallet.TransactOpts, roundNum_, tokenAddress_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indices_) returns()
func (_PrizesWallet *PrizesWalletTransactor) ClaimManyDonatedNfts(opts *bind.TransactOpts, indices_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "claimManyDonatedNfts", indices_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indices_) returns()
func (_PrizesWallet *PrizesWalletSession) ClaimManyDonatedNfts(indices_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimManyDonatedNfts(&_PrizesWallet.TransactOpts, indices_)
}

// ClaimManyDonatedNfts is a paid mutator transaction binding the contract method 0xe4a6c2a4.
//
// Solidity: function claimManyDonatedNfts(uint256[] indices_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) ClaimManyDonatedNfts(indices_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimManyDonatedNfts(&_PrizesWallet.TransactOpts, indices_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x50a15deb.
//
// Solidity: function claimManyDonatedTokens((uint256,address)[] donatedTokensToClaim_) returns()
func (_PrizesWallet *PrizesWalletTransactor) ClaimManyDonatedTokens(opts *bind.TransactOpts, donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "claimManyDonatedTokens", donatedTokensToClaim_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x50a15deb.
//
// Solidity: function claimManyDonatedTokens((uint256,address)[] donatedTokensToClaim_) returns()
func (_PrizesWallet *PrizesWalletSession) ClaimManyDonatedTokens(donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimManyDonatedTokens(&_PrizesWallet.TransactOpts, donatedTokensToClaim_)
}

// ClaimManyDonatedTokens is a paid mutator transaction binding the contract method 0x50a15deb.
//
// Solidity: function claimManyDonatedTokens((uint256,address)[] donatedTokensToClaim_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) ClaimManyDonatedTokens(donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim) (*types.Transaction, error) {
	return _PrizesWallet.Contract.ClaimManyDonatedTokens(&_PrizesWallet.TransactOpts, donatedTokensToClaim_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x46f6b4e1.
//
// Solidity: function depositEth(uint256 roundNum_, address roundPrizeWinnerAddress_) payable returns()
func (_PrizesWallet *PrizesWalletTransactor) DepositEth(opts *bind.TransactOpts, roundNum_ *big.Int, roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "depositEth", roundNum_, roundPrizeWinnerAddress_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x46f6b4e1.
//
// Solidity: function depositEth(uint256 roundNum_, address roundPrizeWinnerAddress_) payable returns()
func (_PrizesWallet *PrizesWalletSession) DepositEth(roundNum_ *big.Int, roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.DepositEth(&_PrizesWallet.TransactOpts, roundNum_, roundPrizeWinnerAddress_)
}

// DepositEth is a paid mutator transaction binding the contract method 0x46f6b4e1.
//
// Solidity: function depositEth(uint256 roundNum_, address roundPrizeWinnerAddress_) payable returns()
func (_PrizesWallet *PrizesWalletTransactorSession) DepositEth(roundNum_ *big.Int, roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.DepositEth(&_PrizesWallet.TransactOpts, roundNum_, roundPrizeWinnerAddress_)
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
// Solidity: function registerRoundEnd(uint256 roundNum_, address roundMainPrizeWinnerAddress_) returns()
func (_PrizesWallet *PrizesWalletTransactor) RegisterRoundEnd(opts *bind.TransactOpts, roundNum_ *big.Int, roundMainPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "registerRoundEnd", roundNum_, roundMainPrizeWinnerAddress_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address roundMainPrizeWinnerAddress_) returns()
func (_PrizesWallet *PrizesWalletSession) RegisterRoundEnd(roundNum_ *big.Int, roundMainPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.RegisterRoundEnd(&_PrizesWallet.TransactOpts, roundNum_, roundMainPrizeWinnerAddress_)
}

// RegisterRoundEnd is a paid mutator transaction binding the contract method 0xcc5810d8.
//
// Solidity: function registerRoundEnd(uint256 roundNum_, address roundMainPrizeWinnerAddress_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) RegisterRoundEnd(roundNum_ *big.Int, roundMainPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.RegisterRoundEnd(&_PrizesWallet.TransactOpts, roundNum_, roundMainPrizeWinnerAddress_)
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
// Solidity: function withdrawEth(address roundPrizeWinnerAddress_) returns()
func (_PrizesWallet *PrizesWalletTransactor) WithdrawEth(opts *bind.TransactOpts, roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "withdrawEth", roundPrizeWinnerAddress_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address roundPrizeWinnerAddress_) returns()
func (_PrizesWallet *PrizesWalletSession) WithdrawEth(roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEth(&_PrizesWallet.TransactOpts, roundPrizeWinnerAddress_)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x25e16063.
//
// Solidity: function withdrawEth(address roundPrizeWinnerAddress_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) WithdrawEth(roundPrizeWinnerAddress_ common.Address) (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEth(&_PrizesWallet.TransactOpts, roundPrizeWinnerAddress_)
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

// WithdrawEverything is a paid mutator transaction binding the contract method 0xa72be1b2.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address)[] donatedTokensToClaim_, uint256[] donatedNftIndices_) returns()
func (_PrizesWallet *PrizesWalletTransactor) WithdrawEverything(opts *bind.TransactOpts, withdrawEth_ bool, donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim, donatedNftIndices_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.contract.Transact(opts, "withdrawEverything", withdrawEth_, donatedTokensToClaim_, donatedNftIndices_)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xa72be1b2.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address)[] donatedTokensToClaim_, uint256[] donatedNftIndices_) returns()
func (_PrizesWallet *PrizesWalletSession) WithdrawEverything(withdrawEth_ bool, donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim, donatedNftIndices_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEverything(&_PrizesWallet.TransactOpts, withdrawEth_, donatedTokensToClaim_, donatedNftIndices_)
}

// WithdrawEverything is a paid mutator transaction binding the contract method 0xa72be1b2.
//
// Solidity: function withdrawEverything(bool withdrawEth_, (uint256,address)[] donatedTokensToClaim_, uint256[] donatedNftIndices_) returns()
func (_PrizesWallet *PrizesWalletTransactorSession) WithdrawEverything(withdrawEth_ bool, donatedTokensToClaim_ []CosmicSignatureConstantsDonatedTokenToClaim, donatedNftIndices_ []*big.Int) (*types.Transaction, error) {
	return _PrizesWallet.Contract.WithdrawEverything(&_PrizesWallet.TransactOpts, withdrawEth_, donatedTokensToClaim_, donatedNftIndices_)
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
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address beneficiaryAddress, address nftAddress, uint256 nftId, uint256 index)
func (_PrizesWallet *PrizesWalletFilterer) FilterDonatedNftClaimed(opts *bind.FilterOpts, roundNum []*big.Int) (*PrizesWalletDonatedNftClaimedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "DonatedNftClaimed", roundNumRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletDonatedNftClaimedIterator{contract: _PrizesWallet.contract, event: "DonatedNftClaimed", logs: logs, sub: sub}, nil
}

// WatchDonatedNftClaimed is a free log subscription operation binding the contract event 0x03c2b6e01c9ca39e4073132f279b61b987d41a82b08cea3dd4a2fabc37067ac3.
//
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address beneficiaryAddress, address nftAddress, uint256 nftId, uint256 index)
func (_PrizesWallet *PrizesWalletFilterer) WatchDonatedNftClaimed(opts *bind.WatchOpts, sink chan<- *PrizesWalletDonatedNftClaimed, roundNum []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "DonatedNftClaimed", roundNumRule)
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
// Solidity: event DonatedNftClaimed(uint256 indexed roundNum, address beneficiaryAddress, address nftAddress, uint256 nftId, uint256 index)
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
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address beneficiaryAddress, address tokenAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) FilterDonatedTokenClaimed(opts *bind.FilterOpts, roundNum []*big.Int) (*PrizesWalletDonatedTokenClaimedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "DonatedTokenClaimed", roundNumRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletDonatedTokenClaimedIterator{contract: _PrizesWallet.contract, event: "DonatedTokenClaimed", logs: logs, sub: sub}, nil
}

// WatchDonatedTokenClaimed is a free log subscription operation binding the contract event 0xaf1adae2e1e983ca738335dc2e37194114142793f394ac934a45ea632a8a5bd0.
//
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address beneficiaryAddress, address tokenAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) WatchDonatedTokenClaimed(opts *bind.WatchOpts, sink chan<- *PrizesWalletDonatedTokenClaimed, roundNum []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "DonatedTokenClaimed", roundNumRule)
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
// Solidity: event DonatedTokenClaimed(uint256 indexed roundNum, address beneficiaryAddress, address tokenAddress, uint256 amount)
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
	RoundNum                *big.Int
	RoundPrizeWinnerAddress common.Address
	Amount                  *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterEthReceived is a free log retrieval operation binding the contract event 0x999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326.
//
// Solidity: event EthReceived(uint256 indexed roundNum, address indexed roundPrizeWinnerAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) FilterEthReceived(opts *bind.FilterOpts, roundNum []*big.Int, roundPrizeWinnerAddress []common.Address) (*PrizesWalletEthReceivedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var roundPrizeWinnerAddressRule []interface{}
	for _, roundPrizeWinnerAddressItem := range roundPrizeWinnerAddress {
		roundPrizeWinnerAddressRule = append(roundPrizeWinnerAddressRule, roundPrizeWinnerAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "EthReceived", roundNumRule, roundPrizeWinnerAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletEthReceivedIterator{contract: _PrizesWallet.contract, event: "EthReceived", logs: logs, sub: sub}, nil
}

// WatchEthReceived is a free log subscription operation binding the contract event 0x999946acc98c7b7dacc26921697d55abbcb1637484b0a73040f0b06287e43326.
//
// Solidity: event EthReceived(uint256 indexed roundNum, address indexed roundPrizeWinnerAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) WatchEthReceived(opts *bind.WatchOpts, sink chan<- *PrizesWalletEthReceived, roundNum []*big.Int, roundPrizeWinnerAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var roundPrizeWinnerAddressRule []interface{}
	for _, roundPrizeWinnerAddressItem := range roundPrizeWinnerAddress {
		roundPrizeWinnerAddressRule = append(roundPrizeWinnerAddressRule, roundPrizeWinnerAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "EthReceived", roundNumRule, roundPrizeWinnerAddressRule)
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
// Solidity: event EthReceived(uint256 indexed roundNum, address indexed roundPrizeWinnerAddress, uint256 amount)
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
	RoundPrizeWinnerAddress common.Address
	BeneficiaryAddress      common.Address
	Amount                  *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawn is a free log retrieval operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed roundPrizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) FilterEthWithdrawn(opts *bind.FilterOpts, roundPrizeWinnerAddress []common.Address, beneficiaryAddress []common.Address) (*PrizesWalletEthWithdrawnIterator, error) {

	var roundPrizeWinnerAddressRule []interface{}
	for _, roundPrizeWinnerAddressItem := range roundPrizeWinnerAddress {
		roundPrizeWinnerAddressRule = append(roundPrizeWinnerAddressRule, roundPrizeWinnerAddressItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.FilterLogs(opts, "EthWithdrawn", roundPrizeWinnerAddressRule, beneficiaryAddressRule)
	if err != nil {
		return nil, err
	}
	return &PrizesWalletEthWithdrawnIterator{contract: _PrizesWallet.contract, event: "EthWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawn is a free log subscription operation binding the contract event 0x4f43b861ba36494acfe938f3815fba7fac6981bdc611b6ccdc14c08f59292383.
//
// Solidity: event EthWithdrawn(address indexed roundPrizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
func (_PrizesWallet *PrizesWalletFilterer) WatchEthWithdrawn(opts *bind.WatchOpts, sink chan<- *PrizesWalletEthWithdrawn, roundPrizeWinnerAddress []common.Address, beneficiaryAddress []common.Address) (event.Subscription, error) {

	var roundPrizeWinnerAddressRule []interface{}
	for _, roundPrizeWinnerAddressItem := range roundPrizeWinnerAddress {
		roundPrizeWinnerAddressRule = append(roundPrizeWinnerAddressRule, roundPrizeWinnerAddressItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	logs, sub, err := _PrizesWallet.contract.WatchLogs(opts, "EthWithdrawn", roundPrizeWinnerAddressRule, beneficiaryAddressRule)
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
// Solidity: event EthWithdrawn(address indexed roundPrizeWinnerAddress, address indexed beneficiaryAddress, uint256 amount)
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
