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

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220adf975bab74bae8a1b2ee567b0874a9ed87daca30bd4fdf4940cf168032ee1d864736f6c63430008130033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// CosmicGameMetaData contains all meta data concerning the CosmicGame contract.
var CosmicGameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newActivationTime\",\"type\":\"uint256\"}],\"name\":\"ActivationTimeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lastBidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"randomWalkNFTId\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prizeTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"BidEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCharity\",\"type\":\"address\"}],\"name\":\"CharityAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCharityPercentage\",\"type\":\"uint256\"}],\"name\":\"CharityPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCosmicSignature\",\"type\":\"address\"}],\"name\":\"CosmicSignatureAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCosmicToken\",\"type\":\"address\"}],\"name\":\"CosmicTokenAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddressdonatedNFTs\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"DonatedNFTClaimedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newInitialBidAmountFraction\",\"type\":\"uint256\"}],\"name\":\"InitialBidAmountFractionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newInitialSecondsUntilPrize\",\"type\":\"uint256\"}],\"name\":\"InitialSecondsUntilPrizeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"NFTDonationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNanoSecondsExtra\",\"type\":\"uint256\"}],\"name\":\"NanoSecondsExtraChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNumRaffleNFTWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"NumRaffleNFTWinnersPerRoundChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNumRaffleWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"NumRaffleWinnersPerRoundChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPriceIncrease\",\"type\":\"uint256\"}],\"name\":\"PriceIncreaseChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PrizeClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrizePercentage\",\"type\":\"uint256\"}],\"name\":\"PrizePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"RaffleNFTClaimedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winner_index\",\"type\":\"uint256\"}],\"name\":\"RaffleNFTWinnerEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRafflePercentage\",\"type\":\"uint256\"}],\"name\":\"RafflePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRaffleWallet\",\"type\":\"address\"}],\"name\":\"RaffleWalletAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRandomWalk\",\"type\":\"address\"}],\"name\":\"RandomWalkAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimeIncrease\",\"type\":\"uint256\"}],\"name\":\"TimeIncreaseChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_MESSAGE_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MILLION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"bidAndDonateNFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"randomWalkNFTId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"bidWithRWLK\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"randomWalkNFTId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"bidWithRWLKAndDonateNFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charity\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charityAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charityPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"claimDonatedNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRaffleNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"donatedNFTs\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialBidAmountFraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialSecondsUntilPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nanoSecondsExtra\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractCosmicSignature\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numDonatedNFTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleNFTWinnersPerRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleParticipants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleWinnersPerRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizeTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raffleAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raffleEntropy\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"raffleNFTWinners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"raffleParticipants\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rafflePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raffleWallet\",\"outputs\":[{\"internalType\":\"contractRaffleWallet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomWalk\",\"outputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newActivationTime\",\"type\":\"uint256\"}],\"name\":\"setActivationTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setCharity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCharityPercentage\",\"type\":\"uint256\"}],\"name\":\"setCharityPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newInitialSecondsUntilPrize\",\"type\":\"uint256\"}],\"name\":\"setInitialSecondsUntilPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNanoSecondsExtra\",\"type\":\"uint256\"}],\"name\":\"setNanoSecondsExtra\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setNftContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNumRaffleNFTWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleNFTWinnersPerRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNumRaffleWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleWinnersPerRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPriceIncrease\",\"type\":\"uint256\"}],\"name\":\"setPriceIncrease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRafflePercentage\",\"type\":\"uint256\"}],\"name\":\"setRafflePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setRaffleWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setRandomWalk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimeIncrease\",\"type\":\"uint256\"}],\"name\":\"setTimeIncrease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeUntilActivation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeUntilPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractCosmicToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newInitialBidAmountFraction\",\"type\":\"uint256\"}],\"name\":\"updateInitialBidAmountFraction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPrizePercentage\",\"type\":\"uint256\"}],\"name\":\"updatePrizePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usedRandomWalkNFTs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"winners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052620f695060015565034630b8a000600255620f42a460035566038d7ea4c680006004556000600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600a600755620151806008556103e86009556019600a556005600b556003600c556005600d556000600f5563644709f0601255348015620000a957600080fd5b50620000ca620000be6200015260201b60201c565b6200015a60201b60201c565b424340604051602001620000e0929190620002b5565b604051602081830303815290604052805190602001206013819055506200010c6200015260201b60201c565b600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620002f7565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600082825260208201905092915050565b7f436f736d6963205369676e617475726520323032330000000000000000000000600082015250565b6000620002676015836200021e565b915062000274826200022f565b602082019050919050565b6000819050919050565b62000294816200027f565b82525050565b6000819050919050565b620002af816200029a565b82525050565b60006060820190508181036000830152620002d08162000258565b9050620002e1602083018562000289565b620002f06040830184620002a4565b9392505050565b61562080620003076000396000f3fe6080604052600436106103d25760003560e01c80638b122274116101fd578063cb819dc011610118578063ec34866d116100ab578063f2fde38b1161007a578063f2fde38b14610dcd578063f717882214610df6578063f8c3405014610e21578063fb6f71a314610e4c578063fc0c546a14610e75576103f1565b8063ec34866d14610d56578063ed08339814610d81578063ed88c68e14610dac578063efeacc5614610db6576103f1565b8063da4493f6116100e7578063da4493f614610c9a578063dbc945c014610cc5578063e1381d7e14610cf0578063ebc3f1d114610d19576103f1565b8063cb819dc014610bdb578063d59d747814610c06578063d6e1741714610c2f578063d94d031614610c6f576103f1565b8063a672f6e111610190578063bb0bc58e1161015f578063bb0bc58e14610b42578063bbcd5bbe14610b5e578063c7c8378d14610b87578063c94028c214610bb0576103f1565b8063a672f6e114610a9c578063a6ceac2c14610ac5578063a6f9cc1514610af0578063ae6247f214610b19576103f1565b80639250c33c116101cc5780639250c33c146109ed578063934aa02314610a18578063a2fb117514610a43578063a325b7d214610a80576103f1565b80638b122274146109435780638b1329e01461096c5780638da5cb5b146109975780639136d6d9146109c2576103f1565b80634ac3a395116102ed578063715018a6116102805780637c5486a21161024f5780637c5486a21461088757806380de163d146108b05780638547af30146108db57806386e378c914610906576103f1565b8063715018a6146107fe578063785fa62714610815578063799d431d146108405780637aef951c1461086b576103f1565b80635e6e47aa116102bc5780635e6e47aa1461076a578063647b3e7f146107935780636e66f6e9146107bc57806370740ac9146107e7576103f1565b80634ac3a395146106c45780635111a2d6146106ed578063519645881461071857806352f5ad7714610741576103f1565b80632aab3223116103655780633c83adc4116103345780633c83adc41461061a57806340e023221461064357806347ccca021461066e5780634a773f3314610699576103f1565b80632aab32231461056e57806332bc934c1461059957806332d382cd146105c45780633bec7b69146105ef576103f1565b806311dc7335116103a157806311dc7335146104b0578063150b7a02146104db57806319afe473146105185780632a02721114610543576103f1565b806304a57c09146103f657806305ba9b6714610433578063062fb1001461045c578063119b22b314610485576103f1565b366103f1576103ef60405180602001604052806000815250610ea0565b005b600080fd5b34801561040257600080fd5b5061041d60048036038101906104189190613dc2565b611087565b60405161042a9190613e30565b60405180910390f35b34801561043f57600080fd5b5061045a60048036038101906104559190613dc2565b6110ba565b005b34801561046857600080fd5b50610483600480360381019061047e9190613e77565b611179565b005b34801561049157600080fd5b5061049a6112df565b6040516104a79190613eb3565b60405180910390f35b3480156104bc57600080fd5b506104c56112e5565b6040516104d29190613eb3565b60405180910390f35b3480156104e757600080fd5b5061050260048036038101906104fd9190613f33565b6112eb565b60405161050f9190613ff6565b60405180910390f35b34801561052457600080fd5b5061052d611300565b60405161053a9190613eb3565b60405180910390f35b34801561054f57600080fd5b50610558611306565b6040516105659190613eb3565b60405180910390f35b34801561057a57600080fd5b5061058361130c565b6040516105909190613eb3565b60405180910390f35b3480156105a557600080fd5b506105ae611312565b6040516105bb9190613eb3565b60405180910390f35b3480156105d057600080fd5b506105d9611319565b6040516105e69190614070565b60405180910390f35b3480156105fb57600080fd5b5061060461133f565b6040516106119190613eb3565b60405180910390f35b34801561062657600080fd5b50610641600480360381019061063c9190613dc2565b611345565b005b34801561064f57600080fd5b50610658611447565b6040516106659190613eb3565b60405180910390f35b34801561067a57600080fd5b5061068361144d565b60405161069091906140ac565b60405180910390f35b3480156106a557600080fd5b506106ae611473565b6040516106bb9190613eb3565b60405180910390f35b3480156106d057600080fd5b506106eb60048036038101906106e69190613dc2565b611479565b005b3480156106f957600080fd5b50610702611538565b60405161070f91906140e8565b60405180910390f35b34801561072457600080fd5b5061073f600480360381019061073a9190613dc2565b61155e565b005b34801561074d57600080fd5b5061076860048036038101906107639190613e77565b61161d565b005b34801561077657600080fd5b50610791600480360381019061078c9190613dc2565b611783565b005b34801561079f57600080fd5b506107ba60048036038101906107b59190613dc2565b611885565b005b3480156107c857600080fd5b506107d1611944565b6040516107de9190613eb3565b60405180910390f35b3480156107f357600080fd5b506107fc611951565b005b34801561080a57600080fd5b5061081361214f565b005b34801561082157600080fd5b5061082a6121d7565b6040516108379190613eb3565b60405180910390f35b34801561084c57600080fd5b506108556121f8565b6040516108629190613eb3565b60405180910390f35b61088560048036038101906108809190614244565b610ea0565b005b34801561089357600080fd5b506108ae60048036038101906108a99190613dc2565b6121fe565b005b3480156108bc57600080fd5b506108c56122bd565b6040516108d291906142a6565b60405180910390f35b3480156108e757600080fd5b506108f06122c3565b6040516108fd9190613e30565b60405180910390f35b34801561091257600080fd5b5061092d60048036038101906109289190613dc2565b6122e9565b60405161093a91906142dc565b60405180910390f35b34801561094f57600080fd5b5061096a60048036038101906109659190613dc2565b612309565b005b34801561097857600080fd5b506109816123c8565b60405161098e9190613eb3565b60405180910390f35b3480156109a357600080fd5b506109ac6123f1565b6040516109b99190613e30565b60405180910390f35b3480156109ce57600080fd5b506109d761241a565b6040516109e49190613eb3565b60405180910390f35b3480156109f957600080fd5b50610a02612420565b604051610a0f9190613eb3565b60405180910390f35b348015610a2457600080fd5b50610a2d612426565b604051610a3a9190613e30565b60405180910390f35b348015610a4f57600080fd5b50610a6a6004803603810190610a659190613dc2565b61244c565b604051610a779190613e30565b60405180910390f35b610a9a6004803603810190610a959190614335565b61247f565b005b348015610aa857600080fd5b50610ac36004803603810190610abe9190613dc2565b612499565b005b348015610ad157600080fd5b50610ada612558565b604051610ae79190613eb3565b60405180910390f35b348015610afc57600080fd5b50610b176004803603810190610b129190613e77565b61255e565b005b348015610b2557600080fd5b50610b406004803603810190610b3b91906143b8565b6126c4565b005b610b5c6004803603810190610b579190614414565b6128eb565b005b348015610b6a57600080fd5b50610b856004803603810190610b809190613e77565b612903565b005b348015610b9357600080fd5b50610bae6004803603810190610ba99190613dc2565b612a69565b005b348015610bbc57600080fd5b50610bc5612b6b565b604051610bd29190613eb3565b60405180910390f35b348015610be757600080fd5b50610bf0612b8c565b604051610bfd9190613eb3565b60405180910390f35b348015610c1257600080fd5b50610c2d6004803603810190610c289190613dc2565b612b92565b005b348015610c3b57600080fd5b50610c566004803603810190610c519190613dc2565b612e8b565b604051610c6694939291906144a4565b60405180910390f35b348015610c7b57600080fd5b50610c84612ee8565b604051610c919190613eb3565b60405180910390f35b348015610ca657600080fd5b50610caf612eee565b604051610cbc9190613eb3565b60405180910390f35b348015610cd157600080fd5b50610cda612ef4565b604051610ce79190613eb3565b60405180910390f35b348015610cfc57600080fd5b50610d176004803603810190610d129190613dc2565b612f15565b005b348015610d2557600080fd5b50610d406004803603810190610d3b9190613e77565b612fd4565b604051610d4d9190613eb3565b60405180910390f35b348015610d6257600080fd5b50610d6b612fec565b604051610d789190613eb3565b60405180910390f35b348015610d8d57600080fd5b50610d96613011565b604051610da39190613eb3565b60405180910390f35b610db4613017565b005b348015610dc257600080fd5b50610dcb613110565b005b348015610dd957600080fd5b50610df46004803603810190610def9190613e77565b61338f565b005b348015610e0257600080fd5b50610e0b613486565b604051610e189190613eb3565b60405180910390f35b348015610e2d57600080fd5b50610e366134af565b604051610e439190613eb3565b60405180910390f35b348015610e5857600080fd5b50610e736004803603810190610e6e9190613e77565b6134b5565b005b348015610e8157600080fd5b50610e8a61363d565b604051610e97919061450a565b60405180910390f35b6000610eaa612fec565b905080341015610eef576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ee6906145a8565b60405180910390fd5b80600481905550610eff82613663565b600454341115610fe6576000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660045434610f5291906145f7565b604051610f5e9061465c565b60006040518083038185875af1925050503d8060008114610f9b576040519150601f19603f3d011682016040523d82523d6000602084013e610fa0565b606091505b5050905080610fe4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fdb906146bd565b60405180910390fd5b505b600f54600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fc7beef8f8d25068377c10e7b7f30ef5622f443eb674e07835df1daf5fe84acbf6004547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600e548760405161107b9493929190614790565b60405180910390a35050565b60156020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6110c26139ae565b73ffffffffffffffffffffffffffffffffffffffff166110e06123f1565b73ffffffffffffffffffffffffffffffffffffffff1614611136576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161112d90614828565b60405180910390fd5b806002819055507f678d086a09e1be49401b1e3a6e14db1878e8d8b88e71d0cfe24a32726d0e38b960025460405161116e9190613eb3565b60405180910390a150565b6111816139ae565b73ffffffffffffffffffffffffffffffffffffffff1661119f6123f1565b73ffffffffffffffffffffffffffffffffffffffff16146111f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111ec90614828565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611264576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161125b90614894565b60405180910390fd5b80601760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f508d510537704d37ed64691a706392abc33e59ce8a327b4952e112820ceb88a6816040516112d49190613e30565b60405180910390a150565b600f5481565b600a5481565b600063150b7a0260e01b905095945050505050565b60045481565b61011881565b60195481565b620f424081565b601760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60165481565b61134d6139ae565b73ffffffffffffffffffffffffffffffffffffffff1661136b6123f1565b73ffffffffffffffffffffffffffffffffffffffff16146113c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113b890614828565b60405180910390fd5b60648110611404576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113fb90614926565b60405180910390fd5b80600a819055507f595fa5ba64cd6f66df19b53b59bb4a275bae1ec5b362da95e1fa4c5feb62be1e600a5460405161143c9190613eb3565b60405180910390a150565b60095481565b601b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600c5481565b6114816139ae565b73ffffffffffffffffffffffffffffffffffffffff1661149f6123f1565b73ffffffffffffffffffffffffffffffffffffffff16146114f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114ec90614828565b60405180910390fd5b806003819055507fed46e73b174ced51fb529cdc6c1b4d8abf49387e6d849b71648afb63c81d12cd60035460405161152d9190613eb3565b60405180910390a150565b601c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6115666139ae565b73ffffffffffffffffffffffffffffffffffffffff166115846123f1565b73ffffffffffffffffffffffffffffffffffffffff16146115da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115d190614828565b60405180910390fd5b806008819055507f6da281754ba85ee0c5983a8e8f05a92910c2a0c5b80e68c126216d65f162a3056008546040516116129190613eb3565b60405180910390a150565b6116256139ae565b73ffffffffffffffffffffffffffffffffffffffff166116436123f1565b73ffffffffffffffffffffffffffffffffffffffff1614611699576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161169090614828565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611708576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116ff90614894565b60405180910390fd5b80601b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f7142a592d5404b5fc4a294c66f70e32b2a7776bd807f722d59268def87c765d1816040516117789190613e30565b60405180910390a150565b61178b6139ae565b73ffffffffffffffffffffffffffffffffffffffff166117a96123f1565b73ffffffffffffffffffffffffffffffffffffffff16146117ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117f690614828565b60405180910390fd5b60648110611842576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161183990614926565b60405180910390fd5b806007819055507f0918b499c15fcb0e087d411d53664cde23577e1aa4a9cbfbdf735ddd1871e7d560075460405161187a9190613eb3565b60405180910390a150565b61188d6139ae565b73ffffffffffffffffffffffffffffffffffffffff166118ab6123f1565b73ffffffffffffffffffffffffffffffffffffffff1614611901576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118f890614828565b60405180910390fd5b80600c819055507f5e3567ae359e61c47ad2f11d2828dea7e062b2ff63dea385cdb7900a50504c7c600c546040516119399190613eb3565b60405180910390a150565b68056bc75e2d6310000081565b42600e541115611996576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161198d90614992565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611a27576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a1e906149fe565b60405180910390fd5b62015180600e5442611a3991906145f7565b1015611ad757600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16611a806139ae565b73ffffffffffffffffffffffffffffffffffffffff1614611ad6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611acd90614ab6565b60405180910390fd5b5b6000611ae16139ae565b90506000600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060116000600f54815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600f6000828254611b8c9190614ad6565b925050819055506000601b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636a62784260e01b83604051602401611be79190613e30565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051611c519190614b46565b6000604051808303816000865af19150503d8060008114611c8e576040519150601f19603f3d011682016040523d82523d6000602084013e611c93565b606091505b5050905080611cd7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611cce90614bcf565b60405180910390fd5b6000611ce16121d7565b90506000611ced612ef4565b90506000611cf9612b6b565b905060005b600d54811015611ddc576000611d126139b6565b90506001601460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611d649190614ad6565b925050819055506001600f54611d7a91906145f7565b8173ffffffffffffffffffffffffffffffffffffffff167f80348bf864c08069d1368c42ed36b7a60560f73267f63d58e9be69f4b021bacc84604051611dc09190613eb3565b60405180910390a3508080611dd490614bef565b915050611cfe565b5060008573ffffffffffffffffffffffffffffffffffffffff1684604051611e039061465c565b60006040518083038185875af1925050503d8060008114611e40576040519150601f19603f3d011682016040523d82523d6000602084013e611e45565b606091505b5050905080611e89576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e8090614c83565b60405180910390fd5b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1683604051611ecf9061465c565b60006040518083038185875af1925050503d8060008114611f0c576040519150601f19603f3d011682016040523d82523d6000602084013e611f11565b606091505b50508091505080611f57576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f4e90614d15565b60405180910390fd5b60005b600c548110156120d9576000611f6e6139b6565b9050601760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16846347e7ef2460e01b836001600f54611fc291906145f7565b604051602401611fd3929190614d35565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161203d9190614b46565b60006040518083038185875af1925050503d806000811461207a576040519150601f19603f3d011682016040523d82523d6000602084013e61207f565b606091505b505080935050826120c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120bc90614daa565b60405180910390fd5b5080806120d190614bef565b915050611f5a565b5060006016819055506120ea613a0b565b8573ffffffffffffffffffffffffffffffffffffffff166001600f5461211091906145f7565b7f27bc828c399c2947fea27bca8a75ced2e94ff2651d607271f051e39db52286ce8660405161213f9190613eb3565b60405180910390a3505050505050565b6121576139ae565b73ffffffffffffffffffffffffffffffffffffffff166121756123f1565b73ffffffffffffffffffffffffffffffffffffffff16146121cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121c290614828565b60405180910390fd5b6121d56000613a28565b565b60006064600a54476121e99190614dca565b6121f39190614e3b565b905090565b60075481565b6122066139ae565b73ffffffffffffffffffffffffffffffffffffffff166122246123f1565b73ffffffffffffffffffffffffffffffffffffffff161461227a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161227190614828565b60405180910390fd5b806012819055507f584ff9a7b3a86db5397052f2e440da0ed60a95f646a3a884863cd92262e682b66012546040516122b29190613eb3565b60405180910390a150565b60135481565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60106020528060005260406000206000915054906101000a900460ff1681565b6123116139ae565b73ffffffffffffffffffffffffffffffffffffffff1661232f6123f1565b73ffffffffffffffffffffffffffffffffffffffff1614612385576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161237c90614828565b60405180910390fd5b806001819055507fcbc1f49adfa29e2f2f0f5c9e057722496a4bc95a6a5446deaa423a02b30c64ac6001546040516123bd9190613eb3565b60405180910390a150565b600042600e5410156123dd57600090506123ee565b42600e546123eb91906145f7565b90505b90565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60025481565b600b5481565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60116020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61248984846126c4565b6124938282613aec565b50505050565b6124a16139ae565b73ffffffffffffffffffffffffffffffffffffffff166124bf6123f1565b73ffffffffffffffffffffffffffffffffffffffff1614612515576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161250c90614828565b60405180910390fd5b806009819055507f3b311f029da1b90c3de2e0b3168436c5ed7d8b0ae81b7d4894c12da03835c62860095460405161254d9190613eb3565b60405180910390a150565b600d5481565b6125666139ae565b73ffffffffffffffffffffffffffffffffffffffff166125846123f1565b73ffffffffffffffffffffffffffffffffffffffff16146125da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125d190614828565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612649576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161264090614894565b60405180910390fd5b80601c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f9dc3707e4b88be65295b0083b44ffa94059c80372add9b0d24d6a3b371a03b92816040516126b99190613e30565b60405180910390a150565b6010600083815260200190815260200160002060009054906101000a900460ff1615612725576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161271c90614ede565b60405180910390fd5b61272d6139ae565b73ffffffffffffffffffffffffffffffffffffffff16601c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e846040518263ffffffff1660e01b815260040161279e9190613eb3565b602060405180830381865afa1580156127bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127df9190614f13565b73ffffffffffffffffffffffffffffffffffffffff1614612835576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161282c90614fb2565b60405180910390fd5b60016010600084815260200190815260200160002060006101000a81548160ff02191690831515021790555061286a81613663565b600f54600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fc7beef8f8d25068377c10e7b7f30ef5622f443eb674e07835df1daf5fe84acbf60045485600e54866040516128df9493929190614fe1565b60405180910390a35050565b6128f483610ea0565b6128fe8282613aec565b505050565b61290b6139ae565b73ffffffffffffffffffffffffffffffffffffffff166129296123f1565b73ffffffffffffffffffffffffffffffffffffffff161461297f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161297690614828565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036129ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129e590614894565b60405180910390fd5b80601a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f3ab2dcf539aa3c30314265602fc86dee8e343e3c3905730f870835f36c400deb81604051612a5e9190613e30565b60405180910390a150565b612a716139ae565b73ffffffffffffffffffffffffffffffffffffffff16612a8f6123f1565b73ffffffffffffffffffffffffffffffffffffffff1614612ae5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612adc90614828565b60405180910390fd5b60648110612b28576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b1f90614926565b60405180910390fd5b80600b819055507fd2002a256ab2f8a5c1b865685754b20437c0a26e7249e40ae5df2993966f99f2600b54604051612b609190613eb3565b60405180910390a150565b60006064600b5447612b7d9190614dca565b612b879190614e3b565b905090565b600e5481565b6019548110612bd6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612bcd90615079565b60405180910390fd5b6000601160006018600085815260200190815260200160002060020154815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff16612c436139ae565b73ffffffffffffffffffffffffffffffffffffffff1614612c99576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c909061510b565b60405180910390fd5b6018600083815260200190815260200160002060030160009054906101000a900460ff1615612cfd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cf49061519d565b60405180910390fd5b60016018600084815260200190815260200160002060030160006101000a81548160ff0219169083151502179055506018600083815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166342842e0e308360186000878152602001908152602001600020600101546040518463ffffffff1660e01b8152600401612db5939291906151bd565b600060405180830381600087803b158015612dcf57600080fd5b505af1158015612de3573d6000803e3d6000fd5b5050505060186000838152602001908152602001600020600201547f0d0e4b90a96d931964d5b0421a6b5b48dd73abf538cdb9ae015069d49e9a467983836018600087815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166018600088815260200190815260200160002060010154604051612e7f94939291906151f4565b60405180910390a25050565b60186020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030160009054906101000a900460ff16905084565b60035481565b60125481565b6000606460075447612f069190614dca565b612f109190614e3b565b905090565b612f1d6139ae565b73ffffffffffffffffffffffffffffffffffffffff16612f3b6123f1565b73ffffffffffffffffffffffffffffffffffffffff1614612f91576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f8890614828565b60405180910390fd5b80600d819055507f72e4278828b8a868e0ba5b4887f954797ea786f8bac991128636171f5eed471d600d54604051612fc99190613eb3565b60405180910390a150565b60146020528060005260406000206000915090505481565b6000620f42406001546004546130029190614dca565b61300c9190614e3b565b905090565b60085481565b6000341161305a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613051906152ab565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036130b9576130b8613a0b565b5b6130c16139ae565b73ffffffffffffffffffffffffffffffffffffffff167f8b7fe5be5699654fd637d2250cb0d47e88205730710745e78e9d8bcaf8aad8f1346040516131069190613eb3565b60405180910390a2565b60006014600061311e6139ae565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205411613199576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131909061533d565b60405180910390fd5b6001601460006131a76139ae565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546131f091906145f7565b925050819055506000601b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636a62784260e01b6132426139ae565b6040516024016132529190613e30565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516132bc9190614b46565b6000604051808303816000865af19150503d80600081146132f9576040519150601f19603f3d011682016040523d82523d6000602084013e6132fe565b606091505b5050905080613342576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161333990614bcf565b60405180910390fd5b61334a6139ae565b73ffffffffffffffffffffffffffffffffffffffff167fe05ba2c5fcd9a60f30b179cb0e775070cc8ce9667b0e663e984ee6a02f694cee60405160405180910390a250565b6133976139ae565b73ffffffffffffffffffffffffffffffffffffffff166133b56123f1565b73ffffffffffffffffffffffffffffffffffffffff161461340b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161340290614828565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361347a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613471906153cf565b60405180910390fd5b61348381613a28565b50565b600042601254101561349b57600090506134ac565b426012546134a991906145f7565b90505b90565b60015481565b6134bd6139ae565b73ffffffffffffffffffffffffffffffffffffffff166134db6123f1565b73ffffffffffffffffffffffffffffffffffffffff1614613531576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161352890614828565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036135a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161359790614894565b60405180910390fd5b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516136329190613e30565b60405180910390a150565b601a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6012544210156136a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161369f9061543b565b60405180910390fd5b610118815111156136ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136e5906154a7565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361375957600854426137529190614ad6565b600e819055505b6137616139ae565b600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660156000601654815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060016016600082825461382a9190614ad6565b925050819055506000601a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1960e01b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1668056bc75e2d631000006040516024016138b2929190614d35565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161391c9190614b46565b6000604051808303816000865af19150503d8060008114613959576040519150601f19603f3d011682016040523d82523d6000602084013e61395e565b606091505b50509050806139a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161399990615539565b60405180910390fd5b6139aa613cce565b5050565b600033905090565b60006139c0613d28565b6015600060165460135460001c6139d79190615559565b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600954613a166121d7565b613a209190614e3b565b600481905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b8173ffffffffffffffffffffffffffffffffffffffff166342842e0e613b106139ae565b30846040518463ffffffff1660e01b8152600401613b30939291906151bd565b600060405180830381600087803b158015613b4a57600080fd5b505af1158015613b5e573d6000803e3d6000fd5b5050505060405180608001604052808373ffffffffffffffffffffffffffffffffffffffff168152602001828152602001600f5481526020016000151581525060186000601954815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002015560608201518160030160006101000a81548160ff021916908315150217905550905050600160196000828254613c449190614ad6565b92505081905550600f548273ffffffffffffffffffffffffffffffffffffffff16613c6d6139ae565b73ffffffffffffffffffffffffffffffffffffffff167fc85be1734ed6a0f4d6adf924d4d41406e2729878c652110a96e2fdec64e118d1846001601954613cb491906145f7565b604051613cc292919061558a565b60405180910390a45050565b6000633b9aca00600254613ce29190614e3b565b905080613cf1600e5442613d5e565b613cfb9190614ad6565b600e81905550620f4240600354600254613d159190614dca565b613d1f9190614e3b565b60028190555050565b601354424340604051602001613d40939291906155b3565b60405160208183030381529060405280519060200120601381905550565b600081831015613d6e5781613d70565b825b905092915050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b613d9f81613d8c565b8114613daa57600080fd5b50565b600081359050613dbc81613d96565b92915050565b600060208284031215613dd857613dd7613d82565b5b6000613de684828501613dad565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000613e1a82613def565b9050919050565b613e2a81613e0f565b82525050565b6000602082019050613e456000830184613e21565b92915050565b613e5481613e0f565b8114613e5f57600080fd5b50565b600081359050613e7181613e4b565b92915050565b600060208284031215613e8d57613e8c613d82565b5b6000613e9b84828501613e62565b91505092915050565b613ead81613d8c565b82525050565b6000602082019050613ec86000830184613ea4565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f840112613ef357613ef2613ece565b5b8235905067ffffffffffffffff811115613f1057613f0f613ed3565b5b602083019150836001820283011115613f2c57613f2b613ed8565b5b9250929050565b600080600080600060808688031215613f4f57613f4e613d82565b5b6000613f5d88828901613e62565b9550506020613f6e88828901613e62565b9450506040613f7f88828901613dad565b935050606086013567ffffffffffffffff811115613fa057613f9f613d87565b5b613fac88828901613edd565b92509250509295509295909350565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b613ff081613fbb565b82525050565b600060208201905061400b6000830184613fe7565b92915050565b6000819050919050565b600061403661403161402c84613def565b614011565b613def565b9050919050565b60006140488261401b565b9050919050565b600061405a8261403d565b9050919050565b61406a8161404f565b82525050565b60006020820190506140856000830184614061565b92915050565b60006140968261403d565b9050919050565b6140a68161408b565b82525050565b60006020820190506140c1600083018461409d565b92915050565b60006140d28261403d565b9050919050565b6140e2816140c7565b82525050565b60006020820190506140fd60008301846140d9565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61415182614108565b810181811067ffffffffffffffff821117156141705761416f614119565b5b80604052505050565b6000614183613d78565b905061418f8282614148565b919050565b600067ffffffffffffffff8211156141af576141ae614119565b5b6141b882614108565b9050602081019050919050565b82818337600083830152505050565b60006141e76141e284614194565b614179565b90508281526020810184848401111561420357614202614103565b5b61420e8482856141c5565b509392505050565b600082601f83011261422b5761422a613ece565b5b813561423b8482602086016141d4565b91505092915050565b60006020828403121561425a57614259613d82565b5b600082013567ffffffffffffffff81111561427857614277613d87565b5b61428484828501614216565b91505092915050565b6000819050919050565b6142a08161428d565b82525050565b60006020820190506142bb6000830184614297565b92915050565b60008115159050919050565b6142d6816142c1565b82525050565b60006020820190506142f160008301846142cd565b92915050565b600061430282613e0f565b9050919050565b614312816142f7565b811461431d57600080fd5b50565b60008135905061432f81614309565b92915050565b6000806000806080858703121561434f5761434e613d82565b5b600061435d87828801613dad565b945050602085013567ffffffffffffffff81111561437e5761437d613d87565b5b61438a87828801614216565b935050604061439b87828801614320565b92505060606143ac87828801613dad565b91505092959194509250565b600080604083850312156143cf576143ce613d82565b5b60006143dd85828601613dad565b925050602083013567ffffffffffffffff8111156143fe576143fd613d87565b5b61440a85828601614216565b9150509250929050565b60008060006060848603121561442d5761442c613d82565b5b600084013567ffffffffffffffff81111561444b5761444a613d87565b5b61445786828701614216565b935050602061446886828701614320565b925050604061447986828701613dad565b9150509250925092565b600061448e8261403d565b9050919050565b61449e81614483565b82525050565b60006080820190506144b96000830187614495565b6144c66020830186613ea4565b6144d36040830185613ea4565b6144e060608301846142cd565b95945050505050565b60006144f48261403d565b9050919050565b614504816144e9565b82525050565b600060208201905061451f60008301846144fb565b92915050565b600082825260208201905092915050565b7f5468652076616c7565207375626d69747465642077697468207468697320747260008201527f616e73616374696f6e20697320746f6f206c6f772e0000000000000000000000602082015250565b6000614592603583614525565b915061459d82614536565b604082019050919050565b600060208201905081810360008301526145c181614585565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061460282613d8c565b915061460d83613d8c565b9250828203905081811115614625576146246145c8565b5b92915050565b600081905092915050565b50565b600061464660008361462b565b915061465182614636565b600082019050919050565b600061466782614639565b9150819050919050565b7f526566756e64207472616e73666572206661696c65642e000000000000000000600082015250565b60006146a7601783614525565b91506146b282614671565b602082019050919050565b600060208201905081810360008301526146d68161469a565b9050919050565b6000819050919050565b6000819050919050565b600061470c614707614702846146dd565b614011565b6146e7565b9050919050565b61471c816146f1565b82525050565b600081519050919050565b60005b8381101561474b578082015181840152602081019050614730565b60008484015250505050565b600061476282614722565b61476c8185614525565b935061477c81856020860161472d565b61478581614108565b840191505092915050565b60006080820190506147a56000830187613ea4565b6147b26020830186614713565b6147bf6040830185613ea4565b81810360608301526147d18184614757565b905095945050505050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000614812602083614525565b915061481d826147dc565b602082019050919050565b6000602082019050818103600083015261484181614805565b9050919050565b7f5a65726f2d616464726573732077617320676976656e2e000000000000000000600082015250565b600061487e601783614525565b915061488982614848565b602082019050919050565b600060208201905081810360008301526148ad81614871565b9050919050565b7f50657263656e746167652076616c7565206f766572666c6f772c206d7573742060008201527f6265206c6f776572207468616e203130302e0000000000000000000000000000602082015250565b6000614910603283614525565b915061491b826148b4565b604082019050919050565b6000602082019050818103600083015261493f81614903565b9050919050565b7f4e6f7420656e6f7567682074696d652068617320656c61707365642e00000000600082015250565b600061497c601c83614525565b915061498782614946565b602082019050919050565b600060208201905081810360008301526149ab8161496f565b9050919050565b7f5468657265206973206e6f206c617374206269646465722e0000000000000000600082015250565b60006149e8601883614525565b91506149f3826149b2565b602082019050919050565b60006020820190508181036000830152614a17816149db565b9050919050565b7f4f6e6c7920746865206c617374206269646465722063616e20636c61696d207460008201527f6865207072697a6520647572696e672074686520666972737420323420686f7560208201527f72732e0000000000000000000000000000000000000000000000000000000000604082015250565b6000614aa0604383614525565b9150614aab82614a1e565b606082019050919050565b60006020820190508181036000830152614acf81614a93565b9050919050565b6000614ae182613d8c565b9150614aec83613d8c565b9250828201905080821115614b0457614b036145c8565b5b92915050565b600081519050919050565b6000614b2082614b0a565b614b2a818561462b565b9350614b3a81856020860161472d565b80840191505092915050565b6000614b528284614b15565b915081905092915050565b7f436f736d69635369676e6174757265206d696e742829206661696c656420746f60008201527f206d696e74204e46542e00000000000000000000000000000000000000000000602082015250565b6000614bb9602a83614525565b9150614bc482614b5d565b604082019050919050565b60006020820190508181036000830152614be881614bac565b9050919050565b6000614bfa82613d8c565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614c2c57614c2b6145c8565b5b600182019050919050565b7f5472616e7366657220746f207468652077696e6e6572206661696c65642e0000600082015250565b6000614c6d601e83614525565b9150614c7882614c37565b602082019050919050565b60006020820190508181036000830152614c9c81614c60565b9050919050565b7f5472616e7366657220746f206368617269747920636f6e74726163742066616960008201527f6c65642e00000000000000000000000000000000000000000000000000000000602082015250565b6000614cff602483614525565b9150614d0a82614ca3565b604082019050919050565b60006020820190508181036000830152614d2e81614cf2565b9050919050565b6000604082019050614d4a6000830185613e21565b614d576020830184613ea4565b9392505050565b7f526166666c65206465706f736974206661696c65642e00000000000000000000600082015250565b6000614d94601683614525565b9150614d9f82614d5e565b602082019050919050565b60006020820190508181036000830152614dc381614d87565b9050919050565b6000614dd582613d8c565b9150614de083613d8c565b9250828202614dee81613d8c565b91508282048414831517614e0557614e046145c8565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000614e4682613d8c565b9150614e5183613d8c565b925082614e6157614e60614e0c565b5b828204905092915050565b7f546869732052616e646f6d57616c6b4e46542068617320616c7265616479206260008201527f65656e207573656420666f722062696464696e672e0000000000000000000000602082015250565b6000614ec8603583614525565b9150614ed382614e6c565b604082019050919050565b60006020820190508181036000830152614ef781614ebb565b9050919050565b600081519050614f0d81613e4b565b92915050565b600060208284031215614f2957614f28613d82565b5b6000614f3784828501614efe565b91505092915050565b7f596f75206d75737420626520746865206f776e6572206f66207468652052616e60008201527f646f6d57616c6b4e46542e000000000000000000000000000000000000000000602082015250565b6000614f9c602b83614525565b9150614fa782614f40565b604082019050919050565b60006020820190508181036000830152614fcb81614f8f565b9050919050565b614fdb816146e7565b82525050565b6000608082019050614ff66000830187613ea4565b6150036020830186614fd2565b6150106040830185613ea4565b81810360608301526150228184614757565b905095945050505050565b7f54686520646f6e61746564204e465420646f6573206e6f742065786973742e00600082015250565b6000615063601f83614525565b915061506e8261502d565b602082019050919050565b6000602082019050818103600083015261509281615056565b9050919050565b7f596f7520617265206e6f74207468652077696e6e6572206f662074686520726f60008201527f756e642e00000000000000000000000000000000000000000000000000000000602082015250565b60006150f5602483614525565b915061510082615099565b604082019050919050565b60006020820190508181036000830152615124816150e8565b9050919050565b7f546865204e46542068617320616c7265616479206265656e20636c61696d656460008201527f2e00000000000000000000000000000000000000000000000000000000000000602082015250565b6000615187602183614525565b91506151928261512b565b604082019050919050565b600060208201905081810360008301526151b68161517a565b9050919050565b60006060820190506151d26000830186613e21565b6151df6020830185613e21565b6151ec6040830184613ea4565b949350505050565b60006080820190506152096000830187613ea4565b6152166020830186613e21565b6152236040830185613e21565b6152306060830184613ea4565b95945050505050565b7f446f6e6174696f6e20616d6f756e74206d75737420626520677265617465722060008201527f7468616e20302e00000000000000000000000000000000000000000000000000602082015250565b6000615295602783614525565b91506152a082615239565b604082019050919050565b600060208201905081810360008301526152c481615288565b9050919050565b7f596f752068617665206e6f20756e636c61696d656420726166666c65204e465460008201527f732e000000000000000000000000000000000000000000000000000000000000602082015250565b6000615327602283614525565b9150615332826152cb565b604082019050919050565b600060208201905081810360008301526153568161531a565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006153b9602683614525565b91506153c48261535d565b604082019050919050565b600060208201905081810360008301526153e8816153ac565b9050919050565b7f4e6f7420616374697665207965742e0000000000000000000000000000000000600082015250565b6000615425600f83614525565b9150615430826153ef565b602082019050919050565b6000602082019050818103600083015261545481615418565b9050919050565b7f4d65737361676520697320746f6f206c6f6e672e000000000000000000000000600082015250565b6000615491601483614525565b915061549c8261545b565b602082019050919050565b600060208201905081810360008301526154c081615484565b9050919050565b7f436f736d6963546f6b656e206d696e742829206661696c656420746f206d696e60008201527f742072657761726420746f6b656e732e00000000000000000000000000000000602082015250565b6000615523603083614525565b915061552e826154c7565b604082019050919050565b6000602082019050818103600083015261555281615516565b9050919050565b600061556482613d8c565b915061556f83613d8c565b92508261557f5761557e614e0c565b5b828206905092915050565b600060408201905061559f6000830185613ea4565b6155ac6020830184613ea4565b9392505050565b60006060820190506155c86000830186614297565b6155d56020830185613ea4565b6155e26040830184614297565b94935050505056fea2646970667358221220208c18d28605972646f40635951f7a251e39f300a183fe8f82959be9c6c7f76764736f6c63430008130033",
}

// CosmicGameABI is the input ABI used to generate the binding from.
// Deprecated: Use CosmicGameMetaData.ABI instead.
var CosmicGameABI = CosmicGameMetaData.ABI

// CosmicGameBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CosmicGameMetaData.Bin instead.
var CosmicGameBin = CosmicGameMetaData.Bin

// DeployCosmicGame deploys a new Ethereum contract, binding an instance of CosmicGame to it.
func DeployCosmicGame(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CosmicGame, error) {
	parsed, err := CosmicGameMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CosmicGameBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CosmicGame{CosmicGameCaller: CosmicGameCaller{contract: contract}, CosmicGameTransactor: CosmicGameTransactor{contract: contract}, CosmicGameFilterer: CosmicGameFilterer{contract: contract}}, nil
}

// CosmicGame is an auto generated Go binding around an Ethereum contract.
type CosmicGame struct {
	CosmicGameCaller     // Read-only binding to the contract
	CosmicGameTransactor // Write-only binding to the contract
	CosmicGameFilterer   // Log filterer for contract events
}

// CosmicGameCaller is an auto generated read-only Go binding around an Ethereum contract.
type CosmicGameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmicGameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CosmicGameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmicGameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CosmicGameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmicGameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CosmicGameSession struct {
	Contract     *CosmicGame       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CosmicGameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CosmicGameCallerSession struct {
	Contract *CosmicGameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CosmicGameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CosmicGameTransactorSession struct {
	Contract     *CosmicGameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CosmicGameRaw is an auto generated low-level Go binding around an Ethereum contract.
type CosmicGameRaw struct {
	Contract *CosmicGame // Generic contract binding to access the raw methods on
}

// CosmicGameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CosmicGameCallerRaw struct {
	Contract *CosmicGameCaller // Generic read-only contract binding to access the raw methods on
}

// CosmicGameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CosmicGameTransactorRaw struct {
	Contract *CosmicGameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCosmicGame creates a new instance of CosmicGame, bound to a specific deployed contract.
func NewCosmicGame(address common.Address, backend bind.ContractBackend) (*CosmicGame, error) {
	contract, err := bindCosmicGame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CosmicGame{CosmicGameCaller: CosmicGameCaller{contract: contract}, CosmicGameTransactor: CosmicGameTransactor{contract: contract}, CosmicGameFilterer: CosmicGameFilterer{contract: contract}}, nil
}

// NewCosmicGameCaller creates a new read-only instance of CosmicGame, bound to a specific deployed contract.
func NewCosmicGameCaller(address common.Address, caller bind.ContractCaller) (*CosmicGameCaller, error) {
	contract, err := bindCosmicGame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CosmicGameCaller{contract: contract}, nil
}

// NewCosmicGameTransactor creates a new write-only instance of CosmicGame, bound to a specific deployed contract.
func NewCosmicGameTransactor(address common.Address, transactor bind.ContractTransactor) (*CosmicGameTransactor, error) {
	contract, err := bindCosmicGame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CosmicGameTransactor{contract: contract}, nil
}

// NewCosmicGameFilterer creates a new log filterer instance of CosmicGame, bound to a specific deployed contract.
func NewCosmicGameFilterer(address common.Address, filterer bind.ContractFilterer) (*CosmicGameFilterer, error) {
	contract, err := bindCosmicGame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CosmicGameFilterer{contract: contract}, nil
}

// bindCosmicGame binds a generic wrapper to an already deployed contract.
func bindCosmicGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CosmicGameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmicGame *CosmicGameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmicGame.Contract.CosmicGameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmicGame *CosmicGameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicGame.Contract.CosmicGameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmicGame *CosmicGameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmicGame.Contract.CosmicGameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmicGame *CosmicGameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmicGame.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmicGame *CosmicGameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicGame.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmicGame *CosmicGameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmicGame.Contract.contract.Transact(opts, method, params...)
}

// MAXMESSAGELENGTH is a free data retrieval call binding the contract method 0x2a027211.
//
// Solidity: function MAX_MESSAGE_LENGTH() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) MAXMESSAGELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "MAX_MESSAGE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMESSAGELENGTH is a free data retrieval call binding the contract method 0x2a027211.
//
// Solidity: function MAX_MESSAGE_LENGTH() view returns(uint256)
func (_CosmicGame *CosmicGameSession) MAXMESSAGELENGTH() (*big.Int, error) {
	return _CosmicGame.Contract.MAXMESSAGELENGTH(&_CosmicGame.CallOpts)
}

// MAXMESSAGELENGTH is a free data retrieval call binding the contract method 0x2a027211.
//
// Solidity: function MAX_MESSAGE_LENGTH() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) MAXMESSAGELENGTH() (*big.Int, error) {
	return _CosmicGame.Contract.MAXMESSAGELENGTH(&_CosmicGame.CallOpts)
}

// MILLION is a free data retrieval call binding the contract method 0x32bc934c.
//
// Solidity: function MILLION() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) MILLION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "MILLION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MILLION is a free data retrieval call binding the contract method 0x32bc934c.
//
// Solidity: function MILLION() view returns(uint256)
func (_CosmicGame *CosmicGameSession) MILLION() (*big.Int, error) {
	return _CosmicGame.Contract.MILLION(&_CosmicGame.CallOpts)
}

// MILLION is a free data retrieval call binding the contract method 0x32bc934c.
//
// Solidity: function MILLION() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) MILLION() (*big.Int, error) {
	return _CosmicGame.Contract.MILLION(&_CosmicGame.CallOpts)
}

// ActivationTime is a free data retrieval call binding the contract method 0xda4493f6.
//
// Solidity: function activationTime() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) ActivationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "activationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActivationTime is a free data retrieval call binding the contract method 0xda4493f6.
//
// Solidity: function activationTime() view returns(uint256)
func (_CosmicGame *CosmicGameSession) ActivationTime() (*big.Int, error) {
	return _CosmicGame.Contract.ActivationTime(&_CosmicGame.CallOpts)
}

// ActivationTime is a free data retrieval call binding the contract method 0xda4493f6.
//
// Solidity: function activationTime() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) ActivationTime() (*big.Int, error) {
	return _CosmicGame.Contract.ActivationTime(&_CosmicGame.CallOpts)
}

// BidPrice is a free data retrieval call binding the contract method 0x19afe473.
//
// Solidity: function bidPrice() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) BidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "bidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BidPrice is a free data retrieval call binding the contract method 0x19afe473.
//
// Solidity: function bidPrice() view returns(uint256)
func (_CosmicGame *CosmicGameSession) BidPrice() (*big.Int, error) {
	return _CosmicGame.Contract.BidPrice(&_CosmicGame.CallOpts)
}

// BidPrice is a free data retrieval call binding the contract method 0x19afe473.
//
// Solidity: function bidPrice() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) BidPrice() (*big.Int, error) {
	return _CosmicGame.Contract.BidPrice(&_CosmicGame.CallOpts)
}

// Charity is a free data retrieval call binding the contract method 0x934aa023.
//
// Solidity: function charity() view returns(address)
func (_CosmicGame *CosmicGameCaller) Charity(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "charity")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Charity is a free data retrieval call binding the contract method 0x934aa023.
//
// Solidity: function charity() view returns(address)
func (_CosmicGame *CosmicGameSession) Charity() (common.Address, error) {
	return _CosmicGame.Contract.Charity(&_CosmicGame.CallOpts)
}

// Charity is a free data retrieval call binding the contract method 0x934aa023.
//
// Solidity: function charity() view returns(address)
func (_CosmicGame *CosmicGameCallerSession) Charity() (common.Address, error) {
	return _CosmicGame.Contract.Charity(&_CosmicGame.CallOpts)
}

// CharityAmount is a free data retrieval call binding the contract method 0xdbc945c0.
//
// Solidity: function charityAmount() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) CharityAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "charityAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CharityAmount is a free data retrieval call binding the contract method 0xdbc945c0.
//
// Solidity: function charityAmount() view returns(uint256)
func (_CosmicGame *CosmicGameSession) CharityAmount() (*big.Int, error) {
	return _CosmicGame.Contract.CharityAmount(&_CosmicGame.CallOpts)
}

// CharityAmount is a free data retrieval call binding the contract method 0xdbc945c0.
//
// Solidity: function charityAmount() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) CharityAmount() (*big.Int, error) {
	return _CosmicGame.Contract.CharityAmount(&_CosmicGame.CallOpts)
}

// CharityPercentage is a free data retrieval call binding the contract method 0x799d431d.
//
// Solidity: function charityPercentage() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) CharityPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "charityPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CharityPercentage is a free data retrieval call binding the contract method 0x799d431d.
//
// Solidity: function charityPercentage() view returns(uint256)
func (_CosmicGame *CosmicGameSession) CharityPercentage() (*big.Int, error) {
	return _CosmicGame.Contract.CharityPercentage(&_CosmicGame.CallOpts)
}

// CharityPercentage is a free data retrieval call binding the contract method 0x799d431d.
//
// Solidity: function charityPercentage() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) CharityPercentage() (*big.Int, error) {
	return _CosmicGame.Contract.CharityPercentage(&_CosmicGame.CallOpts)
}

// DonatedNFTs is a free data retrieval call binding the contract method 0xd6e17417.
//
// Solidity: function donatedNFTs(uint256 ) view returns(address nftAddress, uint256 tokenId, uint256 round, bool claimed)
func (_CosmicGame *CosmicGameCaller) DonatedNFTs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NftAddress common.Address
	TokenId    *big.Int
	Round      *big.Int
	Claimed    bool
}, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "donatedNFTs", arg0)

	outstruct := new(struct {
		NftAddress common.Address
		TokenId    *big.Int
		Round      *big.Int
		Claimed    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NftAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Round = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Claimed = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// DonatedNFTs is a free data retrieval call binding the contract method 0xd6e17417.
//
// Solidity: function donatedNFTs(uint256 ) view returns(address nftAddress, uint256 tokenId, uint256 round, bool claimed)
func (_CosmicGame *CosmicGameSession) DonatedNFTs(arg0 *big.Int) (struct {
	NftAddress common.Address
	TokenId    *big.Int
	Round      *big.Int
	Claimed    bool
}, error) {
	return _CosmicGame.Contract.DonatedNFTs(&_CosmicGame.CallOpts, arg0)
}

// DonatedNFTs is a free data retrieval call binding the contract method 0xd6e17417.
//
// Solidity: function donatedNFTs(uint256 ) view returns(address nftAddress, uint256 tokenId, uint256 round, bool claimed)
func (_CosmicGame *CosmicGameCallerSession) DonatedNFTs(arg0 *big.Int) (struct {
	NftAddress common.Address
	TokenId    *big.Int
	Round      *big.Int
	Claimed    bool
}, error) {
	return _CosmicGame.Contract.DonatedNFTs(&_CosmicGame.CallOpts, arg0)
}

// GetBidPrice is a free data retrieval call binding the contract method 0xec34866d.
//
// Solidity: function getBidPrice() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) GetBidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "getBidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBidPrice is a free data retrieval call binding the contract method 0xec34866d.
//
// Solidity: function getBidPrice() view returns(uint256)
func (_CosmicGame *CosmicGameSession) GetBidPrice() (*big.Int, error) {
	return _CosmicGame.Contract.GetBidPrice(&_CosmicGame.CallOpts)
}

// GetBidPrice is a free data retrieval call binding the contract method 0xec34866d.
//
// Solidity: function getBidPrice() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) GetBidPrice() (*big.Int, error) {
	return _CosmicGame.Contract.GetBidPrice(&_CosmicGame.CallOpts)
}

// InitialBidAmountFraction is a free data retrieval call binding the contract method 0x40e02322.
//
// Solidity: function initialBidAmountFraction() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) InitialBidAmountFraction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "initialBidAmountFraction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialBidAmountFraction is a free data retrieval call binding the contract method 0x40e02322.
//
// Solidity: function initialBidAmountFraction() view returns(uint256)
func (_CosmicGame *CosmicGameSession) InitialBidAmountFraction() (*big.Int, error) {
	return _CosmicGame.Contract.InitialBidAmountFraction(&_CosmicGame.CallOpts)
}

// InitialBidAmountFraction is a free data retrieval call binding the contract method 0x40e02322.
//
// Solidity: function initialBidAmountFraction() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) InitialBidAmountFraction() (*big.Int, error) {
	return _CosmicGame.Contract.InitialBidAmountFraction(&_CosmicGame.CallOpts)
}

// InitialSecondsUntilPrize is a free data retrieval call binding the contract method 0xed083398.
//
// Solidity: function initialSecondsUntilPrize() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) InitialSecondsUntilPrize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "initialSecondsUntilPrize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialSecondsUntilPrize is a free data retrieval call binding the contract method 0xed083398.
//
// Solidity: function initialSecondsUntilPrize() view returns(uint256)
func (_CosmicGame *CosmicGameSession) InitialSecondsUntilPrize() (*big.Int, error) {
	return _CosmicGame.Contract.InitialSecondsUntilPrize(&_CosmicGame.CallOpts)
}

// InitialSecondsUntilPrize is a free data retrieval call binding the contract method 0xed083398.
//
// Solidity: function initialSecondsUntilPrize() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) InitialSecondsUntilPrize() (*big.Int, error) {
	return _CosmicGame.Contract.InitialSecondsUntilPrize(&_CosmicGame.CallOpts)
}

// LastBidder is a free data retrieval call binding the contract method 0x8547af30.
//
// Solidity: function lastBidder() view returns(address)
func (_CosmicGame *CosmicGameCaller) LastBidder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "lastBidder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastBidder is a free data retrieval call binding the contract method 0x8547af30.
//
// Solidity: function lastBidder() view returns(address)
func (_CosmicGame *CosmicGameSession) LastBidder() (common.Address, error) {
	return _CosmicGame.Contract.LastBidder(&_CosmicGame.CallOpts)
}

// LastBidder is a free data retrieval call binding the contract method 0x8547af30.
//
// Solidity: function lastBidder() view returns(address)
func (_CosmicGame *CosmicGameCallerSession) LastBidder() (common.Address, error) {
	return _CosmicGame.Contract.LastBidder(&_CosmicGame.CallOpts)
}

// NanoSecondsExtra is a free data retrieval call binding the contract method 0x9136d6d9.
//
// Solidity: function nanoSecondsExtra() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) NanoSecondsExtra(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "nanoSecondsExtra")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NanoSecondsExtra is a free data retrieval call binding the contract method 0x9136d6d9.
//
// Solidity: function nanoSecondsExtra() view returns(uint256)
func (_CosmicGame *CosmicGameSession) NanoSecondsExtra() (*big.Int, error) {
	return _CosmicGame.Contract.NanoSecondsExtra(&_CosmicGame.CallOpts)
}

// NanoSecondsExtra is a free data retrieval call binding the contract method 0x9136d6d9.
//
// Solidity: function nanoSecondsExtra() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) NanoSecondsExtra() (*big.Int, error) {
	return _CosmicGame.Contract.NanoSecondsExtra(&_CosmicGame.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_CosmicGame *CosmicGameCaller) Nft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "nft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_CosmicGame *CosmicGameSession) Nft() (common.Address, error) {
	return _CosmicGame.Contract.Nft(&_CosmicGame.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_CosmicGame *CosmicGameCallerSession) Nft() (common.Address, error) {
	return _CosmicGame.Contract.Nft(&_CosmicGame.CallOpts)
}

// NumDonatedNFTs is a free data retrieval call binding the contract method 0x2aab3223.
//
// Solidity: function numDonatedNFTs() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) NumDonatedNFTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "numDonatedNFTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumDonatedNFTs is a free data retrieval call binding the contract method 0x2aab3223.
//
// Solidity: function numDonatedNFTs() view returns(uint256)
func (_CosmicGame *CosmicGameSession) NumDonatedNFTs() (*big.Int, error) {
	return _CosmicGame.Contract.NumDonatedNFTs(&_CosmicGame.CallOpts)
}

// NumDonatedNFTs is a free data retrieval call binding the contract method 0x2aab3223.
//
// Solidity: function numDonatedNFTs() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) NumDonatedNFTs() (*big.Int, error) {
	return _CosmicGame.Contract.NumDonatedNFTs(&_CosmicGame.CallOpts)
}

// NumRaffleNFTWinnersPerRound is a free data retrieval call binding the contract method 0xa6ceac2c.
//
// Solidity: function numRaffleNFTWinnersPerRound() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) NumRaffleNFTWinnersPerRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "numRaffleNFTWinnersPerRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRaffleNFTWinnersPerRound is a free data retrieval call binding the contract method 0xa6ceac2c.
//
// Solidity: function numRaffleNFTWinnersPerRound() view returns(uint256)
func (_CosmicGame *CosmicGameSession) NumRaffleNFTWinnersPerRound() (*big.Int, error) {
	return _CosmicGame.Contract.NumRaffleNFTWinnersPerRound(&_CosmicGame.CallOpts)
}

// NumRaffleNFTWinnersPerRound is a free data retrieval call binding the contract method 0xa6ceac2c.
//
// Solidity: function numRaffleNFTWinnersPerRound() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) NumRaffleNFTWinnersPerRound() (*big.Int, error) {
	return _CosmicGame.Contract.NumRaffleNFTWinnersPerRound(&_CosmicGame.CallOpts)
}

// NumRaffleParticipants is a free data retrieval call binding the contract method 0x3bec7b69.
//
// Solidity: function numRaffleParticipants() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) NumRaffleParticipants(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "numRaffleParticipants")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRaffleParticipants is a free data retrieval call binding the contract method 0x3bec7b69.
//
// Solidity: function numRaffleParticipants() view returns(uint256)
func (_CosmicGame *CosmicGameSession) NumRaffleParticipants() (*big.Int, error) {
	return _CosmicGame.Contract.NumRaffleParticipants(&_CosmicGame.CallOpts)
}

// NumRaffleParticipants is a free data retrieval call binding the contract method 0x3bec7b69.
//
// Solidity: function numRaffleParticipants() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) NumRaffleParticipants() (*big.Int, error) {
	return _CosmicGame.Contract.NumRaffleParticipants(&_CosmicGame.CallOpts)
}

// NumRaffleWinnersPerRound is a free data retrieval call binding the contract method 0x4a773f33.
//
// Solidity: function numRaffleWinnersPerRound() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) NumRaffleWinnersPerRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "numRaffleWinnersPerRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRaffleWinnersPerRound is a free data retrieval call binding the contract method 0x4a773f33.
//
// Solidity: function numRaffleWinnersPerRound() view returns(uint256)
func (_CosmicGame *CosmicGameSession) NumRaffleWinnersPerRound() (*big.Int, error) {
	return _CosmicGame.Contract.NumRaffleWinnersPerRound(&_CosmicGame.CallOpts)
}

// NumRaffleWinnersPerRound is a free data retrieval call binding the contract method 0x4a773f33.
//
// Solidity: function numRaffleWinnersPerRound() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) NumRaffleWinnersPerRound() (*big.Int, error) {
	return _CosmicGame.Contract.NumRaffleWinnersPerRound(&_CosmicGame.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_CosmicGame *CosmicGameCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_CosmicGame *CosmicGameSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _CosmicGame.Contract.OnERC721Received(&_CosmicGame.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_CosmicGame *CosmicGameCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _CosmicGame.Contract.OnERC721Received(&_CosmicGame.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CosmicGame *CosmicGameCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CosmicGame *CosmicGameSession) Owner() (common.Address, error) {
	return _CosmicGame.Contract.Owner(&_CosmicGame.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CosmicGame *CosmicGameCallerSession) Owner() (common.Address, error) {
	return _CosmicGame.Contract.Owner(&_CosmicGame.CallOpts)
}

// PriceIncrease is a free data retrieval call binding the contract method 0xf8c34050.
//
// Solidity: function priceIncrease() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) PriceIncrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "priceIncrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceIncrease is a free data retrieval call binding the contract method 0xf8c34050.
//
// Solidity: function priceIncrease() view returns(uint256)
func (_CosmicGame *CosmicGameSession) PriceIncrease() (*big.Int, error) {
	return _CosmicGame.Contract.PriceIncrease(&_CosmicGame.CallOpts)
}

// PriceIncrease is a free data retrieval call binding the contract method 0xf8c34050.
//
// Solidity: function priceIncrease() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) PriceIncrease() (*big.Int, error) {
	return _CosmicGame.Contract.PriceIncrease(&_CosmicGame.CallOpts)
}

// PrizeAmount is a free data retrieval call binding the contract method 0x785fa627.
//
// Solidity: function prizeAmount() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) PrizeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "prizeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrizeAmount is a free data retrieval call binding the contract method 0x785fa627.
//
// Solidity: function prizeAmount() view returns(uint256)
func (_CosmicGame *CosmicGameSession) PrizeAmount() (*big.Int, error) {
	return _CosmicGame.Contract.PrizeAmount(&_CosmicGame.CallOpts)
}

// PrizeAmount is a free data retrieval call binding the contract method 0x785fa627.
//
// Solidity: function prizeAmount() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) PrizeAmount() (*big.Int, error) {
	return _CosmicGame.Contract.PrizeAmount(&_CosmicGame.CallOpts)
}

// PrizePercentage is a free data retrieval call binding the contract method 0x11dc7335.
//
// Solidity: function prizePercentage() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) PrizePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "prizePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrizePercentage is a free data retrieval call binding the contract method 0x11dc7335.
//
// Solidity: function prizePercentage() view returns(uint256)
func (_CosmicGame *CosmicGameSession) PrizePercentage() (*big.Int, error) {
	return _CosmicGame.Contract.PrizePercentage(&_CosmicGame.CallOpts)
}

// PrizePercentage is a free data retrieval call binding the contract method 0x11dc7335.
//
// Solidity: function prizePercentage() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) PrizePercentage() (*big.Int, error) {
	return _CosmicGame.Contract.PrizePercentage(&_CosmicGame.CallOpts)
}

// PrizeTime is a free data retrieval call binding the contract method 0xcb819dc0.
//
// Solidity: function prizeTime() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) PrizeTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "prizeTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrizeTime is a free data retrieval call binding the contract method 0xcb819dc0.
//
// Solidity: function prizeTime() view returns(uint256)
func (_CosmicGame *CosmicGameSession) PrizeTime() (*big.Int, error) {
	return _CosmicGame.Contract.PrizeTime(&_CosmicGame.CallOpts)
}

// PrizeTime is a free data retrieval call binding the contract method 0xcb819dc0.
//
// Solidity: function prizeTime() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) PrizeTime() (*big.Int, error) {
	return _CosmicGame.Contract.PrizeTime(&_CosmicGame.CallOpts)
}

// RaffleAmount is a free data retrieval call binding the contract method 0xc94028c2.
//
// Solidity: function raffleAmount() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) RaffleAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "raffleAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaffleAmount is a free data retrieval call binding the contract method 0xc94028c2.
//
// Solidity: function raffleAmount() view returns(uint256)
func (_CosmicGame *CosmicGameSession) RaffleAmount() (*big.Int, error) {
	return _CosmicGame.Contract.RaffleAmount(&_CosmicGame.CallOpts)
}

// RaffleAmount is a free data retrieval call binding the contract method 0xc94028c2.
//
// Solidity: function raffleAmount() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) RaffleAmount() (*big.Int, error) {
	return _CosmicGame.Contract.RaffleAmount(&_CosmicGame.CallOpts)
}

// RaffleEntropy is a free data retrieval call binding the contract method 0x80de163d.
//
// Solidity: function raffleEntropy() view returns(bytes32)
func (_CosmicGame *CosmicGameCaller) RaffleEntropy(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "raffleEntropy")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RaffleEntropy is a free data retrieval call binding the contract method 0x80de163d.
//
// Solidity: function raffleEntropy() view returns(bytes32)
func (_CosmicGame *CosmicGameSession) RaffleEntropy() ([32]byte, error) {
	return _CosmicGame.Contract.RaffleEntropy(&_CosmicGame.CallOpts)
}

// RaffleEntropy is a free data retrieval call binding the contract method 0x80de163d.
//
// Solidity: function raffleEntropy() view returns(bytes32)
func (_CosmicGame *CosmicGameCallerSession) RaffleEntropy() ([32]byte, error) {
	return _CosmicGame.Contract.RaffleEntropy(&_CosmicGame.CallOpts)
}

// RaffleNFTWinners is a free data retrieval call binding the contract method 0xebc3f1d1.
//
// Solidity: function raffleNFTWinners(address ) view returns(uint256)
func (_CosmicGame *CosmicGameCaller) RaffleNFTWinners(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "raffleNFTWinners", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaffleNFTWinners is a free data retrieval call binding the contract method 0xebc3f1d1.
//
// Solidity: function raffleNFTWinners(address ) view returns(uint256)
func (_CosmicGame *CosmicGameSession) RaffleNFTWinners(arg0 common.Address) (*big.Int, error) {
	return _CosmicGame.Contract.RaffleNFTWinners(&_CosmicGame.CallOpts, arg0)
}

// RaffleNFTWinners is a free data retrieval call binding the contract method 0xebc3f1d1.
//
// Solidity: function raffleNFTWinners(address ) view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) RaffleNFTWinners(arg0 common.Address) (*big.Int, error) {
	return _CosmicGame.Contract.RaffleNFTWinners(&_CosmicGame.CallOpts, arg0)
}

// RaffleParticipants is a free data retrieval call binding the contract method 0x04a57c09.
//
// Solidity: function raffleParticipants(uint256 ) view returns(address)
func (_CosmicGame *CosmicGameCaller) RaffleParticipants(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "raffleParticipants", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RaffleParticipants is a free data retrieval call binding the contract method 0x04a57c09.
//
// Solidity: function raffleParticipants(uint256 ) view returns(address)
func (_CosmicGame *CosmicGameSession) RaffleParticipants(arg0 *big.Int) (common.Address, error) {
	return _CosmicGame.Contract.RaffleParticipants(&_CosmicGame.CallOpts, arg0)
}

// RaffleParticipants is a free data retrieval call binding the contract method 0x04a57c09.
//
// Solidity: function raffleParticipants(uint256 ) view returns(address)
func (_CosmicGame *CosmicGameCallerSession) RaffleParticipants(arg0 *big.Int) (common.Address, error) {
	return _CosmicGame.Contract.RaffleParticipants(&_CosmicGame.CallOpts, arg0)
}

// RafflePercentage is a free data retrieval call binding the contract method 0x9250c33c.
//
// Solidity: function rafflePercentage() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) RafflePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "rafflePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RafflePercentage is a free data retrieval call binding the contract method 0x9250c33c.
//
// Solidity: function rafflePercentage() view returns(uint256)
func (_CosmicGame *CosmicGameSession) RafflePercentage() (*big.Int, error) {
	return _CosmicGame.Contract.RafflePercentage(&_CosmicGame.CallOpts)
}

// RafflePercentage is a free data retrieval call binding the contract method 0x9250c33c.
//
// Solidity: function rafflePercentage() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) RafflePercentage() (*big.Int, error) {
	return _CosmicGame.Contract.RafflePercentage(&_CosmicGame.CallOpts)
}

// RaffleWallet is a free data retrieval call binding the contract method 0x32d382cd.
//
// Solidity: function raffleWallet() view returns(address)
func (_CosmicGame *CosmicGameCaller) RaffleWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "raffleWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RaffleWallet is a free data retrieval call binding the contract method 0x32d382cd.
//
// Solidity: function raffleWallet() view returns(address)
func (_CosmicGame *CosmicGameSession) RaffleWallet() (common.Address, error) {
	return _CosmicGame.Contract.RaffleWallet(&_CosmicGame.CallOpts)
}

// RaffleWallet is a free data retrieval call binding the contract method 0x32d382cd.
//
// Solidity: function raffleWallet() view returns(address)
func (_CosmicGame *CosmicGameCallerSession) RaffleWallet() (common.Address, error) {
	return _CosmicGame.Contract.RaffleWallet(&_CosmicGame.CallOpts)
}

// RandomWalk is a free data retrieval call binding the contract method 0x5111a2d6.
//
// Solidity: function randomWalk() view returns(address)
func (_CosmicGame *CosmicGameCaller) RandomWalk(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "randomWalk")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RandomWalk is a free data retrieval call binding the contract method 0x5111a2d6.
//
// Solidity: function randomWalk() view returns(address)
func (_CosmicGame *CosmicGameSession) RandomWalk() (common.Address, error) {
	return _CosmicGame.Contract.RandomWalk(&_CosmicGame.CallOpts)
}

// RandomWalk is a free data retrieval call binding the contract method 0x5111a2d6.
//
// Solidity: function randomWalk() view returns(address)
func (_CosmicGame *CosmicGameCallerSession) RandomWalk() (common.Address, error) {
	return _CosmicGame.Contract.RandomWalk(&_CosmicGame.CallOpts)
}

// RoundNum is a free data retrieval call binding the contract method 0x119b22b3.
//
// Solidity: function roundNum() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) RoundNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "roundNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundNum is a free data retrieval call binding the contract method 0x119b22b3.
//
// Solidity: function roundNum() view returns(uint256)
func (_CosmicGame *CosmicGameSession) RoundNum() (*big.Int, error) {
	return _CosmicGame.Contract.RoundNum(&_CosmicGame.CallOpts)
}

// RoundNum is a free data retrieval call binding the contract method 0x119b22b3.
//
// Solidity: function roundNum() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) RoundNum() (*big.Int, error) {
	return _CosmicGame.Contract.RoundNum(&_CosmicGame.CallOpts)
}

// TimeIncrease is a free data retrieval call binding the contract method 0xd94d0316.
//
// Solidity: function timeIncrease() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) TimeIncrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "timeIncrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeIncrease is a free data retrieval call binding the contract method 0xd94d0316.
//
// Solidity: function timeIncrease() view returns(uint256)
func (_CosmicGame *CosmicGameSession) TimeIncrease() (*big.Int, error) {
	return _CosmicGame.Contract.TimeIncrease(&_CosmicGame.CallOpts)
}

// TimeIncrease is a free data retrieval call binding the contract method 0xd94d0316.
//
// Solidity: function timeIncrease() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) TimeIncrease() (*big.Int, error) {
	return _CosmicGame.Contract.TimeIncrease(&_CosmicGame.CallOpts)
}

// TimeUntilActivation is a free data retrieval call binding the contract method 0xf7178822.
//
// Solidity: function timeUntilActivation() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) TimeUntilActivation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "timeUntilActivation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeUntilActivation is a free data retrieval call binding the contract method 0xf7178822.
//
// Solidity: function timeUntilActivation() view returns(uint256)
func (_CosmicGame *CosmicGameSession) TimeUntilActivation() (*big.Int, error) {
	return _CosmicGame.Contract.TimeUntilActivation(&_CosmicGame.CallOpts)
}

// TimeUntilActivation is a free data retrieval call binding the contract method 0xf7178822.
//
// Solidity: function timeUntilActivation() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) TimeUntilActivation() (*big.Int, error) {
	return _CosmicGame.Contract.TimeUntilActivation(&_CosmicGame.CallOpts)
}

// TimeUntilPrize is a free data retrieval call binding the contract method 0x8b1329e0.
//
// Solidity: function timeUntilPrize() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) TimeUntilPrize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "timeUntilPrize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeUntilPrize is a free data retrieval call binding the contract method 0x8b1329e0.
//
// Solidity: function timeUntilPrize() view returns(uint256)
func (_CosmicGame *CosmicGameSession) TimeUntilPrize() (*big.Int, error) {
	return _CosmicGame.Contract.TimeUntilPrize(&_CosmicGame.CallOpts)
}

// TimeUntilPrize is a free data retrieval call binding the contract method 0x8b1329e0.
//
// Solidity: function timeUntilPrize() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) TimeUntilPrize() (*big.Int, error) {
	return _CosmicGame.Contract.TimeUntilPrize(&_CosmicGame.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CosmicGame *CosmicGameCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CosmicGame *CosmicGameSession) Token() (common.Address, error) {
	return _CosmicGame.Contract.Token(&_CosmicGame.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CosmicGame *CosmicGameCallerSession) Token() (common.Address, error) {
	return _CosmicGame.Contract.Token(&_CosmicGame.CallOpts)
}

// TokenReward is a free data retrieval call binding the contract method 0x6e66f6e9.
//
// Solidity: function tokenReward() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) TokenReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "tokenReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenReward is a free data retrieval call binding the contract method 0x6e66f6e9.
//
// Solidity: function tokenReward() view returns(uint256)
func (_CosmicGame *CosmicGameSession) TokenReward() (*big.Int, error) {
	return _CosmicGame.Contract.TokenReward(&_CosmicGame.CallOpts)
}

// TokenReward is a free data retrieval call binding the contract method 0x6e66f6e9.
//
// Solidity: function tokenReward() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) TokenReward() (*big.Int, error) {
	return _CosmicGame.Contract.TokenReward(&_CosmicGame.CallOpts)
}

// UsedRandomWalkNFTs is a free data retrieval call binding the contract method 0x86e378c9.
//
// Solidity: function usedRandomWalkNFTs(uint256 ) view returns(bool)
func (_CosmicGame *CosmicGameCaller) UsedRandomWalkNFTs(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "usedRandomWalkNFTs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedRandomWalkNFTs is a free data retrieval call binding the contract method 0x86e378c9.
//
// Solidity: function usedRandomWalkNFTs(uint256 ) view returns(bool)
func (_CosmicGame *CosmicGameSession) UsedRandomWalkNFTs(arg0 *big.Int) (bool, error) {
	return _CosmicGame.Contract.UsedRandomWalkNFTs(&_CosmicGame.CallOpts, arg0)
}

// UsedRandomWalkNFTs is a free data retrieval call binding the contract method 0x86e378c9.
//
// Solidity: function usedRandomWalkNFTs(uint256 ) view returns(bool)
func (_CosmicGame *CosmicGameCallerSession) UsedRandomWalkNFTs(arg0 *big.Int) (bool, error) {
	return _CosmicGame.Contract.UsedRandomWalkNFTs(&_CosmicGame.CallOpts, arg0)
}

// Winners is a free data retrieval call binding the contract method 0xa2fb1175.
//
// Solidity: function winners(uint256 ) view returns(address)
func (_CosmicGame *CosmicGameCaller) Winners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "winners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Winners is a free data retrieval call binding the contract method 0xa2fb1175.
//
// Solidity: function winners(uint256 ) view returns(address)
func (_CosmicGame *CosmicGameSession) Winners(arg0 *big.Int) (common.Address, error) {
	return _CosmicGame.Contract.Winners(&_CosmicGame.CallOpts, arg0)
}

// Winners is a free data retrieval call binding the contract method 0xa2fb1175.
//
// Solidity: function winners(uint256 ) view returns(address)
func (_CosmicGame *CosmicGameCallerSession) Winners(arg0 *big.Int) (common.Address, error) {
	return _CosmicGame.Contract.Winners(&_CosmicGame.CallOpts, arg0)
}

// Bid is a paid mutator transaction binding the contract method 0x7aef951c.
//
// Solidity: function bid(string message) payable returns()
func (_CosmicGame *CosmicGameTransactor) Bid(opts *bind.TransactOpts, message string) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "bid", message)
}

// Bid is a paid mutator transaction binding the contract method 0x7aef951c.
//
// Solidity: function bid(string message) payable returns()
func (_CosmicGame *CosmicGameSession) Bid(message string) (*types.Transaction, error) {
	return _CosmicGame.Contract.Bid(&_CosmicGame.TransactOpts, message)
}

// Bid is a paid mutator transaction binding the contract method 0x7aef951c.
//
// Solidity: function bid(string message) payable returns()
func (_CosmicGame *CosmicGameTransactorSession) Bid(message string) (*types.Transaction, error) {
	return _CosmicGame.Contract.Bid(&_CosmicGame.TransactOpts, message)
}

// BidAndDonateNFT is a paid mutator transaction binding the contract method 0xbb0bc58e.
//
// Solidity: function bidAndDonateNFT(string message, address nftAddress, uint256 tokenId) payable returns()
func (_CosmicGame *CosmicGameTransactor) BidAndDonateNFT(opts *bind.TransactOpts, message string, nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "bidAndDonateNFT", message, nftAddress, tokenId)
}

// BidAndDonateNFT is a paid mutator transaction binding the contract method 0xbb0bc58e.
//
// Solidity: function bidAndDonateNFT(string message, address nftAddress, uint256 tokenId) payable returns()
func (_CosmicGame *CosmicGameSession) BidAndDonateNFT(message string, nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.BidAndDonateNFT(&_CosmicGame.TransactOpts, message, nftAddress, tokenId)
}

// BidAndDonateNFT is a paid mutator transaction binding the contract method 0xbb0bc58e.
//
// Solidity: function bidAndDonateNFT(string message, address nftAddress, uint256 tokenId) payable returns()
func (_CosmicGame *CosmicGameTransactorSession) BidAndDonateNFT(message string, nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.BidAndDonateNFT(&_CosmicGame.TransactOpts, message, nftAddress, tokenId)
}

// BidWithRWLK is a paid mutator transaction binding the contract method 0xae6247f2.
//
// Solidity: function bidWithRWLK(uint256 randomWalkNFTId, string message) returns()
func (_CosmicGame *CosmicGameTransactor) BidWithRWLK(opts *bind.TransactOpts, randomWalkNFTId *big.Int, message string) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "bidWithRWLK", randomWalkNFTId, message)
}

// BidWithRWLK is a paid mutator transaction binding the contract method 0xae6247f2.
//
// Solidity: function bidWithRWLK(uint256 randomWalkNFTId, string message) returns()
func (_CosmicGame *CosmicGameSession) BidWithRWLK(randomWalkNFTId *big.Int, message string) (*types.Transaction, error) {
	return _CosmicGame.Contract.BidWithRWLK(&_CosmicGame.TransactOpts, randomWalkNFTId, message)
}

// BidWithRWLK is a paid mutator transaction binding the contract method 0xae6247f2.
//
// Solidity: function bidWithRWLK(uint256 randomWalkNFTId, string message) returns()
func (_CosmicGame *CosmicGameTransactorSession) BidWithRWLK(randomWalkNFTId *big.Int, message string) (*types.Transaction, error) {
	return _CosmicGame.Contract.BidWithRWLK(&_CosmicGame.TransactOpts, randomWalkNFTId, message)
}

// BidWithRWLKAndDonateNFT is a paid mutator transaction binding the contract method 0xa325b7d2.
//
// Solidity: function bidWithRWLKAndDonateNFT(uint256 randomWalkNFTId, string message, address nftAddress, uint256 tokenId) payable returns()
func (_CosmicGame *CosmicGameTransactor) BidWithRWLKAndDonateNFT(opts *bind.TransactOpts, randomWalkNFTId *big.Int, message string, nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "bidWithRWLKAndDonateNFT", randomWalkNFTId, message, nftAddress, tokenId)
}

// BidWithRWLKAndDonateNFT is a paid mutator transaction binding the contract method 0xa325b7d2.
//
// Solidity: function bidWithRWLKAndDonateNFT(uint256 randomWalkNFTId, string message, address nftAddress, uint256 tokenId) payable returns()
func (_CosmicGame *CosmicGameSession) BidWithRWLKAndDonateNFT(randomWalkNFTId *big.Int, message string, nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.BidWithRWLKAndDonateNFT(&_CosmicGame.TransactOpts, randomWalkNFTId, message, nftAddress, tokenId)
}

// BidWithRWLKAndDonateNFT is a paid mutator transaction binding the contract method 0xa325b7d2.
//
// Solidity: function bidWithRWLKAndDonateNFT(uint256 randomWalkNFTId, string message, address nftAddress, uint256 tokenId) payable returns()
func (_CosmicGame *CosmicGameTransactorSession) BidWithRWLKAndDonateNFT(randomWalkNFTId *big.Int, message string, nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.BidWithRWLKAndDonateNFT(&_CosmicGame.TransactOpts, randomWalkNFTId, message, nftAddress, tokenId)
}

// ClaimDonatedNFT is a paid mutator transaction binding the contract method 0xd59d7478.
//
// Solidity: function claimDonatedNFT(uint256 num) returns()
func (_CosmicGame *CosmicGameTransactor) ClaimDonatedNFT(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "claimDonatedNFT", num)
}

// ClaimDonatedNFT is a paid mutator transaction binding the contract method 0xd59d7478.
//
// Solidity: function claimDonatedNFT(uint256 num) returns()
func (_CosmicGame *CosmicGameSession) ClaimDonatedNFT(num *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.ClaimDonatedNFT(&_CosmicGame.TransactOpts, num)
}

// ClaimDonatedNFT is a paid mutator transaction binding the contract method 0xd59d7478.
//
// Solidity: function claimDonatedNFT(uint256 num) returns()
func (_CosmicGame *CosmicGameTransactorSession) ClaimDonatedNFT(num *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.ClaimDonatedNFT(&_CosmicGame.TransactOpts, num)
}

// ClaimPrize is a paid mutator transaction binding the contract method 0x70740ac9.
//
// Solidity: function claimPrize() returns()
func (_CosmicGame *CosmicGameTransactor) ClaimPrize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "claimPrize")
}

// ClaimPrize is a paid mutator transaction binding the contract method 0x70740ac9.
//
// Solidity: function claimPrize() returns()
func (_CosmicGame *CosmicGameSession) ClaimPrize() (*types.Transaction, error) {
	return _CosmicGame.Contract.ClaimPrize(&_CosmicGame.TransactOpts)
}

// ClaimPrize is a paid mutator transaction binding the contract method 0x70740ac9.
//
// Solidity: function claimPrize() returns()
func (_CosmicGame *CosmicGameTransactorSession) ClaimPrize() (*types.Transaction, error) {
	return _CosmicGame.Contract.ClaimPrize(&_CosmicGame.TransactOpts)
}

// ClaimRaffleNFT is a paid mutator transaction binding the contract method 0xefeacc56.
//
// Solidity: function claimRaffleNFT() returns()
func (_CosmicGame *CosmicGameTransactor) ClaimRaffleNFT(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "claimRaffleNFT")
}

// ClaimRaffleNFT is a paid mutator transaction binding the contract method 0xefeacc56.
//
// Solidity: function claimRaffleNFT() returns()
func (_CosmicGame *CosmicGameSession) ClaimRaffleNFT() (*types.Transaction, error) {
	return _CosmicGame.Contract.ClaimRaffleNFT(&_CosmicGame.TransactOpts)
}

// ClaimRaffleNFT is a paid mutator transaction binding the contract method 0xefeacc56.
//
// Solidity: function claimRaffleNFT() returns()
func (_CosmicGame *CosmicGameTransactorSession) ClaimRaffleNFT() (*types.Transaction, error) {
	return _CosmicGame.Contract.ClaimRaffleNFT(&_CosmicGame.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_CosmicGame *CosmicGameTransactor) Donate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "donate")
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_CosmicGame *CosmicGameSession) Donate() (*types.Transaction, error) {
	return _CosmicGame.Contract.Donate(&_CosmicGame.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_CosmicGame *CosmicGameTransactorSession) Donate() (*types.Transaction, error) {
	return _CosmicGame.Contract.Donate(&_CosmicGame.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CosmicGame *CosmicGameTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CosmicGame *CosmicGameSession) RenounceOwnership() (*types.Transaction, error) {
	return _CosmicGame.Contract.RenounceOwnership(&_CosmicGame.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CosmicGame *CosmicGameTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CosmicGame.Contract.RenounceOwnership(&_CosmicGame.TransactOpts)
}

// SetActivationTime is a paid mutator transaction binding the contract method 0x7c5486a2.
//
// Solidity: function setActivationTime(uint256 newActivationTime) returns()
func (_CosmicGame *CosmicGameTransactor) SetActivationTime(opts *bind.TransactOpts, newActivationTime *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setActivationTime", newActivationTime)
}

// SetActivationTime is a paid mutator transaction binding the contract method 0x7c5486a2.
//
// Solidity: function setActivationTime(uint256 newActivationTime) returns()
func (_CosmicGame *CosmicGameSession) SetActivationTime(newActivationTime *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetActivationTime(&_CosmicGame.TransactOpts, newActivationTime)
}

// SetActivationTime is a paid mutator transaction binding the contract method 0x7c5486a2.
//
// Solidity: function setActivationTime(uint256 newActivationTime) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetActivationTime(newActivationTime *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetActivationTime(&_CosmicGame.TransactOpts, newActivationTime)
}

// SetCharity is a paid mutator transaction binding the contract method 0xfb6f71a3.
//
// Solidity: function setCharity(address addr) returns()
func (_CosmicGame *CosmicGameTransactor) SetCharity(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setCharity", addr)
}

// SetCharity is a paid mutator transaction binding the contract method 0xfb6f71a3.
//
// Solidity: function setCharity(address addr) returns()
func (_CosmicGame *CosmicGameSession) SetCharity(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetCharity(&_CosmicGame.TransactOpts, addr)
}

// SetCharity is a paid mutator transaction binding the contract method 0xfb6f71a3.
//
// Solidity: function setCharity(address addr) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetCharity(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetCharity(&_CosmicGame.TransactOpts, addr)
}

// SetCharityPercentage is a paid mutator transaction binding the contract method 0x5e6e47aa.
//
// Solidity: function setCharityPercentage(uint256 newCharityPercentage) returns()
func (_CosmicGame *CosmicGameTransactor) SetCharityPercentage(opts *bind.TransactOpts, newCharityPercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setCharityPercentage", newCharityPercentage)
}

// SetCharityPercentage is a paid mutator transaction binding the contract method 0x5e6e47aa.
//
// Solidity: function setCharityPercentage(uint256 newCharityPercentage) returns()
func (_CosmicGame *CosmicGameSession) SetCharityPercentage(newCharityPercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetCharityPercentage(&_CosmicGame.TransactOpts, newCharityPercentage)
}

// SetCharityPercentage is a paid mutator transaction binding the contract method 0x5e6e47aa.
//
// Solidity: function setCharityPercentage(uint256 newCharityPercentage) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetCharityPercentage(newCharityPercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetCharityPercentage(&_CosmicGame.TransactOpts, newCharityPercentage)
}

// SetInitialSecondsUntilPrize is a paid mutator transaction binding the contract method 0x51964588.
//
// Solidity: function setInitialSecondsUntilPrize(uint256 newInitialSecondsUntilPrize) returns()
func (_CosmicGame *CosmicGameTransactor) SetInitialSecondsUntilPrize(opts *bind.TransactOpts, newInitialSecondsUntilPrize *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setInitialSecondsUntilPrize", newInitialSecondsUntilPrize)
}

// SetInitialSecondsUntilPrize is a paid mutator transaction binding the contract method 0x51964588.
//
// Solidity: function setInitialSecondsUntilPrize(uint256 newInitialSecondsUntilPrize) returns()
func (_CosmicGame *CosmicGameSession) SetInitialSecondsUntilPrize(newInitialSecondsUntilPrize *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetInitialSecondsUntilPrize(&_CosmicGame.TransactOpts, newInitialSecondsUntilPrize)
}

// SetInitialSecondsUntilPrize is a paid mutator transaction binding the contract method 0x51964588.
//
// Solidity: function setInitialSecondsUntilPrize(uint256 newInitialSecondsUntilPrize) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetInitialSecondsUntilPrize(newInitialSecondsUntilPrize *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetInitialSecondsUntilPrize(&_CosmicGame.TransactOpts, newInitialSecondsUntilPrize)
}

// SetNanoSecondsExtra is a paid mutator transaction binding the contract method 0x05ba9b67.
//
// Solidity: function setNanoSecondsExtra(uint256 newNanoSecondsExtra) returns()
func (_CosmicGame *CosmicGameTransactor) SetNanoSecondsExtra(opts *bind.TransactOpts, newNanoSecondsExtra *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setNanoSecondsExtra", newNanoSecondsExtra)
}

// SetNanoSecondsExtra is a paid mutator transaction binding the contract method 0x05ba9b67.
//
// Solidity: function setNanoSecondsExtra(uint256 newNanoSecondsExtra) returns()
func (_CosmicGame *CosmicGameSession) SetNanoSecondsExtra(newNanoSecondsExtra *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNanoSecondsExtra(&_CosmicGame.TransactOpts, newNanoSecondsExtra)
}

// SetNanoSecondsExtra is a paid mutator transaction binding the contract method 0x05ba9b67.
//
// Solidity: function setNanoSecondsExtra(uint256 newNanoSecondsExtra) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetNanoSecondsExtra(newNanoSecondsExtra *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNanoSecondsExtra(&_CosmicGame.TransactOpts, newNanoSecondsExtra)
}

// SetNftContract is a paid mutator transaction binding the contract method 0x52f5ad77.
//
// Solidity: function setNftContract(address addr) returns()
func (_CosmicGame *CosmicGameTransactor) SetNftContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setNftContract", addr)
}

// SetNftContract is a paid mutator transaction binding the contract method 0x52f5ad77.
//
// Solidity: function setNftContract(address addr) returns()
func (_CosmicGame *CosmicGameSession) SetNftContract(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNftContract(&_CosmicGame.TransactOpts, addr)
}

// SetNftContract is a paid mutator transaction binding the contract method 0x52f5ad77.
//
// Solidity: function setNftContract(address addr) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetNftContract(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNftContract(&_CosmicGame.TransactOpts, addr)
}

// SetNumRaffleNFTWinnersPerRound is a paid mutator transaction binding the contract method 0xe1381d7e.
//
// Solidity: function setNumRaffleNFTWinnersPerRound(uint256 newNumRaffleNFTWinnersPerRound) returns()
func (_CosmicGame *CosmicGameTransactor) SetNumRaffleNFTWinnersPerRound(opts *bind.TransactOpts, newNumRaffleNFTWinnersPerRound *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setNumRaffleNFTWinnersPerRound", newNumRaffleNFTWinnersPerRound)
}

// SetNumRaffleNFTWinnersPerRound is a paid mutator transaction binding the contract method 0xe1381d7e.
//
// Solidity: function setNumRaffleNFTWinnersPerRound(uint256 newNumRaffleNFTWinnersPerRound) returns()
func (_CosmicGame *CosmicGameSession) SetNumRaffleNFTWinnersPerRound(newNumRaffleNFTWinnersPerRound *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNumRaffleNFTWinnersPerRound(&_CosmicGame.TransactOpts, newNumRaffleNFTWinnersPerRound)
}

// SetNumRaffleNFTWinnersPerRound is a paid mutator transaction binding the contract method 0xe1381d7e.
//
// Solidity: function setNumRaffleNFTWinnersPerRound(uint256 newNumRaffleNFTWinnersPerRound) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetNumRaffleNFTWinnersPerRound(newNumRaffleNFTWinnersPerRound *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNumRaffleNFTWinnersPerRound(&_CosmicGame.TransactOpts, newNumRaffleNFTWinnersPerRound)
}

// SetNumRaffleWinnersPerRound is a paid mutator transaction binding the contract method 0x647b3e7f.
//
// Solidity: function setNumRaffleWinnersPerRound(uint256 newNumRaffleWinnersPerRound) returns()
func (_CosmicGame *CosmicGameTransactor) SetNumRaffleWinnersPerRound(opts *bind.TransactOpts, newNumRaffleWinnersPerRound *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setNumRaffleWinnersPerRound", newNumRaffleWinnersPerRound)
}

// SetNumRaffleWinnersPerRound is a paid mutator transaction binding the contract method 0x647b3e7f.
//
// Solidity: function setNumRaffleWinnersPerRound(uint256 newNumRaffleWinnersPerRound) returns()
func (_CosmicGame *CosmicGameSession) SetNumRaffleWinnersPerRound(newNumRaffleWinnersPerRound *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNumRaffleWinnersPerRound(&_CosmicGame.TransactOpts, newNumRaffleWinnersPerRound)
}

// SetNumRaffleWinnersPerRound is a paid mutator transaction binding the contract method 0x647b3e7f.
//
// Solidity: function setNumRaffleWinnersPerRound(uint256 newNumRaffleWinnersPerRound) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetNumRaffleWinnersPerRound(newNumRaffleWinnersPerRound *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNumRaffleWinnersPerRound(&_CosmicGame.TransactOpts, newNumRaffleWinnersPerRound)
}

// SetPriceIncrease is a paid mutator transaction binding the contract method 0x8b122274.
//
// Solidity: function setPriceIncrease(uint256 newPriceIncrease) returns()
func (_CosmicGame *CosmicGameTransactor) SetPriceIncrease(opts *bind.TransactOpts, newPriceIncrease *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setPriceIncrease", newPriceIncrease)
}

// SetPriceIncrease is a paid mutator transaction binding the contract method 0x8b122274.
//
// Solidity: function setPriceIncrease(uint256 newPriceIncrease) returns()
func (_CosmicGame *CosmicGameSession) SetPriceIncrease(newPriceIncrease *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetPriceIncrease(&_CosmicGame.TransactOpts, newPriceIncrease)
}

// SetPriceIncrease is a paid mutator transaction binding the contract method 0x8b122274.
//
// Solidity: function setPriceIncrease(uint256 newPriceIncrease) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetPriceIncrease(newPriceIncrease *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetPriceIncrease(&_CosmicGame.TransactOpts, newPriceIncrease)
}

// SetRafflePercentage is a paid mutator transaction binding the contract method 0xc7c8378d.
//
// Solidity: function setRafflePercentage(uint256 newRafflePercentage) returns()
func (_CosmicGame *CosmicGameTransactor) SetRafflePercentage(opts *bind.TransactOpts, newRafflePercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setRafflePercentage", newRafflePercentage)
}

// SetRafflePercentage is a paid mutator transaction binding the contract method 0xc7c8378d.
//
// Solidity: function setRafflePercentage(uint256 newRafflePercentage) returns()
func (_CosmicGame *CosmicGameSession) SetRafflePercentage(newRafflePercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetRafflePercentage(&_CosmicGame.TransactOpts, newRafflePercentage)
}

// SetRafflePercentage is a paid mutator transaction binding the contract method 0xc7c8378d.
//
// Solidity: function setRafflePercentage(uint256 newRafflePercentage) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetRafflePercentage(newRafflePercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetRafflePercentage(&_CosmicGame.TransactOpts, newRafflePercentage)
}

// SetRaffleWallet is a paid mutator transaction binding the contract method 0x062fb100.
//
// Solidity: function setRaffleWallet(address addr) returns()
func (_CosmicGame *CosmicGameTransactor) SetRaffleWallet(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setRaffleWallet", addr)
}

// SetRaffleWallet is a paid mutator transaction binding the contract method 0x062fb100.
//
// Solidity: function setRaffleWallet(address addr) returns()
func (_CosmicGame *CosmicGameSession) SetRaffleWallet(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetRaffleWallet(&_CosmicGame.TransactOpts, addr)
}

// SetRaffleWallet is a paid mutator transaction binding the contract method 0x062fb100.
//
// Solidity: function setRaffleWallet(address addr) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetRaffleWallet(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetRaffleWallet(&_CosmicGame.TransactOpts, addr)
}

// SetRandomWalk is a paid mutator transaction binding the contract method 0xa6f9cc15.
//
// Solidity: function setRandomWalk(address addr) returns()
func (_CosmicGame *CosmicGameTransactor) SetRandomWalk(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setRandomWalk", addr)
}

// SetRandomWalk is a paid mutator transaction binding the contract method 0xa6f9cc15.
//
// Solidity: function setRandomWalk(address addr) returns()
func (_CosmicGame *CosmicGameSession) SetRandomWalk(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetRandomWalk(&_CosmicGame.TransactOpts, addr)
}

// SetRandomWalk is a paid mutator transaction binding the contract method 0xa6f9cc15.
//
// Solidity: function setRandomWalk(address addr) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetRandomWalk(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetRandomWalk(&_CosmicGame.TransactOpts, addr)
}

// SetTimeIncrease is a paid mutator transaction binding the contract method 0x4ac3a395.
//
// Solidity: function setTimeIncrease(uint256 newTimeIncrease) returns()
func (_CosmicGame *CosmicGameTransactor) SetTimeIncrease(opts *bind.TransactOpts, newTimeIncrease *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setTimeIncrease", newTimeIncrease)
}

// SetTimeIncrease is a paid mutator transaction binding the contract method 0x4ac3a395.
//
// Solidity: function setTimeIncrease(uint256 newTimeIncrease) returns()
func (_CosmicGame *CosmicGameSession) SetTimeIncrease(newTimeIncrease *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetTimeIncrease(&_CosmicGame.TransactOpts, newTimeIncrease)
}

// SetTimeIncrease is a paid mutator transaction binding the contract method 0x4ac3a395.
//
// Solidity: function setTimeIncrease(uint256 newTimeIncrease) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetTimeIncrease(newTimeIncrease *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetTimeIncrease(&_CosmicGame.TransactOpts, newTimeIncrease)
}

// SetTokenContract is a paid mutator transaction binding the contract method 0xbbcd5bbe.
//
// Solidity: function setTokenContract(address addr) returns()
func (_CosmicGame *CosmicGameTransactor) SetTokenContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setTokenContract", addr)
}

// SetTokenContract is a paid mutator transaction binding the contract method 0xbbcd5bbe.
//
// Solidity: function setTokenContract(address addr) returns()
func (_CosmicGame *CosmicGameSession) SetTokenContract(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetTokenContract(&_CosmicGame.TransactOpts, addr)
}

// SetTokenContract is a paid mutator transaction binding the contract method 0xbbcd5bbe.
//
// Solidity: function setTokenContract(address addr) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetTokenContract(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetTokenContract(&_CosmicGame.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CosmicGame *CosmicGameTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CosmicGame *CosmicGameSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.TransferOwnership(&_CosmicGame.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CosmicGame *CosmicGameTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.TransferOwnership(&_CosmicGame.TransactOpts, newOwner)
}

// UpdateInitialBidAmountFraction is a paid mutator transaction binding the contract method 0xa672f6e1.
//
// Solidity: function updateInitialBidAmountFraction(uint256 newInitialBidAmountFraction) returns()
func (_CosmicGame *CosmicGameTransactor) UpdateInitialBidAmountFraction(opts *bind.TransactOpts, newInitialBidAmountFraction *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "updateInitialBidAmountFraction", newInitialBidAmountFraction)
}

// UpdateInitialBidAmountFraction is a paid mutator transaction binding the contract method 0xa672f6e1.
//
// Solidity: function updateInitialBidAmountFraction(uint256 newInitialBidAmountFraction) returns()
func (_CosmicGame *CosmicGameSession) UpdateInitialBidAmountFraction(newInitialBidAmountFraction *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.UpdateInitialBidAmountFraction(&_CosmicGame.TransactOpts, newInitialBidAmountFraction)
}

// UpdateInitialBidAmountFraction is a paid mutator transaction binding the contract method 0xa672f6e1.
//
// Solidity: function updateInitialBidAmountFraction(uint256 newInitialBidAmountFraction) returns()
func (_CosmicGame *CosmicGameTransactorSession) UpdateInitialBidAmountFraction(newInitialBidAmountFraction *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.UpdateInitialBidAmountFraction(&_CosmicGame.TransactOpts, newInitialBidAmountFraction)
}

// UpdatePrizePercentage is a paid mutator transaction binding the contract method 0x3c83adc4.
//
// Solidity: function updatePrizePercentage(uint256 newPrizePercentage) returns()
func (_CosmicGame *CosmicGameTransactor) UpdatePrizePercentage(opts *bind.TransactOpts, newPrizePercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "updatePrizePercentage", newPrizePercentage)
}

// UpdatePrizePercentage is a paid mutator transaction binding the contract method 0x3c83adc4.
//
// Solidity: function updatePrizePercentage(uint256 newPrizePercentage) returns()
func (_CosmicGame *CosmicGameSession) UpdatePrizePercentage(newPrizePercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.UpdatePrizePercentage(&_CosmicGame.TransactOpts, newPrizePercentage)
}

// UpdatePrizePercentage is a paid mutator transaction binding the contract method 0x3c83adc4.
//
// Solidity: function updatePrizePercentage(uint256 newPrizePercentage) returns()
func (_CosmicGame *CosmicGameTransactorSession) UpdatePrizePercentage(newPrizePercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.UpdatePrizePercentage(&_CosmicGame.TransactOpts, newPrizePercentage)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CosmicGame *CosmicGameTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicGame.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CosmicGame *CosmicGameSession) Receive() (*types.Transaction, error) {
	return _CosmicGame.Contract.Receive(&_CosmicGame.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CosmicGame *CosmicGameTransactorSession) Receive() (*types.Transaction, error) {
	return _CosmicGame.Contract.Receive(&_CosmicGame.TransactOpts)
}

// CosmicGameActivationTimeChangedIterator is returned from FilterActivationTimeChanged and is used to iterate over the raw logs and unpacked data for ActivationTimeChanged events raised by the CosmicGame contract.
type CosmicGameActivationTimeChangedIterator struct {
	Event *CosmicGameActivationTimeChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameActivationTimeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameActivationTimeChanged)
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
		it.Event = new(CosmicGameActivationTimeChanged)
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
func (it *CosmicGameActivationTimeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameActivationTimeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameActivationTimeChanged represents a ActivationTimeChanged event raised by the CosmicGame contract.
type CosmicGameActivationTimeChanged struct {
	NewActivationTime *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterActivationTimeChanged is a free log retrieval operation binding the contract event 0x584ff9a7b3a86db5397052f2e440da0ed60a95f646a3a884863cd92262e682b6.
//
// Solidity: event ActivationTimeChanged(uint256 newActivationTime)
func (_CosmicGame *CosmicGameFilterer) FilterActivationTimeChanged(opts *bind.FilterOpts) (*CosmicGameActivationTimeChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "ActivationTimeChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameActivationTimeChangedIterator{contract: _CosmicGame.contract, event: "ActivationTimeChanged", logs: logs, sub: sub}, nil
}

// WatchActivationTimeChanged is a free log subscription operation binding the contract event 0x584ff9a7b3a86db5397052f2e440da0ed60a95f646a3a884863cd92262e682b6.
//
// Solidity: event ActivationTimeChanged(uint256 newActivationTime)
func (_CosmicGame *CosmicGameFilterer) WatchActivationTimeChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameActivationTimeChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "ActivationTimeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameActivationTimeChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "ActivationTimeChanged", log); err != nil {
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

// ParseActivationTimeChanged is a log parse operation binding the contract event 0x584ff9a7b3a86db5397052f2e440da0ed60a95f646a3a884863cd92262e682b6.
//
// Solidity: event ActivationTimeChanged(uint256 newActivationTime)
func (_CosmicGame *CosmicGameFilterer) ParseActivationTimeChanged(log types.Log) (*CosmicGameActivationTimeChanged, error) {
	event := new(CosmicGameActivationTimeChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "ActivationTimeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameBidEventIterator is returned from FilterBidEvent and is used to iterate over the raw logs and unpacked data for BidEvent events raised by the CosmicGame contract.
type CosmicGameBidEventIterator struct {
	Event *CosmicGameBidEvent // Event containing the contract specifics and raw log

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
func (it *CosmicGameBidEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameBidEvent)
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
		it.Event = new(CosmicGameBidEvent)
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
func (it *CosmicGameBidEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameBidEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameBidEvent represents a BidEvent event raised by the CosmicGame contract.
type CosmicGameBidEvent struct {
	LastBidder      common.Address
	Round           *big.Int
	BidPrice        *big.Int
	RandomWalkNFTId *big.Int
	PrizeTime       *big.Int
	Message         string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBidEvent is a free log retrieval operation binding the contract event 0xc7beef8f8d25068377c10e7b7f30ef5622f443eb674e07835df1daf5fe84acbf.
//
// Solidity: event BidEvent(address indexed lastBidder, uint256 indexed round, uint256 bidPrice, int256 randomWalkNFTId, uint256 prizeTime, string message)
func (_CosmicGame *CosmicGameFilterer) FilterBidEvent(opts *bind.FilterOpts, lastBidder []common.Address, round []*big.Int) (*CosmicGameBidEventIterator, error) {

	var lastBidderRule []interface{}
	for _, lastBidderItem := range lastBidder {
		lastBidderRule = append(lastBidderRule, lastBidderItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "BidEvent", lastBidderRule, roundRule)
	if err != nil {
		return nil, err
	}
	return &CosmicGameBidEventIterator{contract: _CosmicGame.contract, event: "BidEvent", logs: logs, sub: sub}, nil
}

// WatchBidEvent is a free log subscription operation binding the contract event 0xc7beef8f8d25068377c10e7b7f30ef5622f443eb674e07835df1daf5fe84acbf.
//
// Solidity: event BidEvent(address indexed lastBidder, uint256 indexed round, uint256 bidPrice, int256 randomWalkNFTId, uint256 prizeTime, string message)
func (_CosmicGame *CosmicGameFilterer) WatchBidEvent(opts *bind.WatchOpts, sink chan<- *CosmicGameBidEvent, lastBidder []common.Address, round []*big.Int) (event.Subscription, error) {

	var lastBidderRule []interface{}
	for _, lastBidderItem := range lastBidder {
		lastBidderRule = append(lastBidderRule, lastBidderItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "BidEvent", lastBidderRule, roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameBidEvent)
				if err := _CosmicGame.contract.UnpackLog(event, "BidEvent", log); err != nil {
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

// ParseBidEvent is a log parse operation binding the contract event 0xc7beef8f8d25068377c10e7b7f30ef5622f443eb674e07835df1daf5fe84acbf.
//
// Solidity: event BidEvent(address indexed lastBidder, uint256 indexed round, uint256 bidPrice, int256 randomWalkNFTId, uint256 prizeTime, string message)
func (_CosmicGame *CosmicGameFilterer) ParseBidEvent(log types.Log) (*CosmicGameBidEvent, error) {
	event := new(CosmicGameBidEvent)
	if err := _CosmicGame.contract.UnpackLog(event, "BidEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameCharityAddressChangedIterator is returned from FilterCharityAddressChanged and is used to iterate over the raw logs and unpacked data for CharityAddressChanged events raised by the CosmicGame contract.
type CosmicGameCharityAddressChangedIterator struct {
	Event *CosmicGameCharityAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameCharityAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameCharityAddressChanged)
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
		it.Event = new(CosmicGameCharityAddressChanged)
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
func (it *CosmicGameCharityAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameCharityAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameCharityAddressChanged represents a CharityAddressChanged event raised by the CosmicGame contract.
type CosmicGameCharityAddressChanged struct {
	NewCharity common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCharityAddressChanged is a free log retrieval operation binding the contract event 0x1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c.
//
// Solidity: event CharityAddressChanged(address newCharity)
func (_CosmicGame *CosmicGameFilterer) FilterCharityAddressChanged(opts *bind.FilterOpts) (*CosmicGameCharityAddressChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "CharityAddressChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameCharityAddressChangedIterator{contract: _CosmicGame.contract, event: "CharityAddressChanged", logs: logs, sub: sub}, nil
}

// WatchCharityAddressChanged is a free log subscription operation binding the contract event 0x1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c.
//
// Solidity: event CharityAddressChanged(address newCharity)
func (_CosmicGame *CosmicGameFilterer) WatchCharityAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameCharityAddressChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "CharityAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameCharityAddressChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "CharityAddressChanged", log); err != nil {
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
// Solidity: event CharityAddressChanged(address newCharity)
func (_CosmicGame *CosmicGameFilterer) ParseCharityAddressChanged(log types.Log) (*CosmicGameCharityAddressChanged, error) {
	event := new(CosmicGameCharityAddressChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "CharityAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameCharityPercentageChangedIterator is returned from FilterCharityPercentageChanged and is used to iterate over the raw logs and unpacked data for CharityPercentageChanged events raised by the CosmicGame contract.
type CosmicGameCharityPercentageChangedIterator struct {
	Event *CosmicGameCharityPercentageChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameCharityPercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameCharityPercentageChanged)
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
		it.Event = new(CosmicGameCharityPercentageChanged)
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
func (it *CosmicGameCharityPercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameCharityPercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameCharityPercentageChanged represents a CharityPercentageChanged event raised by the CosmicGame contract.
type CosmicGameCharityPercentageChanged struct {
	NewCharityPercentage *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterCharityPercentageChanged is a free log retrieval operation binding the contract event 0x0918b499c15fcb0e087d411d53664cde23577e1aa4a9cbfbdf735ddd1871e7d5.
//
// Solidity: event CharityPercentageChanged(uint256 newCharityPercentage)
func (_CosmicGame *CosmicGameFilterer) FilterCharityPercentageChanged(opts *bind.FilterOpts) (*CosmicGameCharityPercentageChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "CharityPercentageChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameCharityPercentageChangedIterator{contract: _CosmicGame.contract, event: "CharityPercentageChanged", logs: logs, sub: sub}, nil
}

// WatchCharityPercentageChanged is a free log subscription operation binding the contract event 0x0918b499c15fcb0e087d411d53664cde23577e1aa4a9cbfbdf735ddd1871e7d5.
//
// Solidity: event CharityPercentageChanged(uint256 newCharityPercentage)
func (_CosmicGame *CosmicGameFilterer) WatchCharityPercentageChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameCharityPercentageChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "CharityPercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameCharityPercentageChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "CharityPercentageChanged", log); err != nil {
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

// ParseCharityPercentageChanged is a log parse operation binding the contract event 0x0918b499c15fcb0e087d411d53664cde23577e1aa4a9cbfbdf735ddd1871e7d5.
//
// Solidity: event CharityPercentageChanged(uint256 newCharityPercentage)
func (_CosmicGame *CosmicGameFilterer) ParseCharityPercentageChanged(log types.Log) (*CosmicGameCharityPercentageChanged, error) {
	event := new(CosmicGameCharityPercentageChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "CharityPercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameCosmicSignatureAddressChangedIterator is returned from FilterCosmicSignatureAddressChanged and is used to iterate over the raw logs and unpacked data for CosmicSignatureAddressChanged events raised by the CosmicGame contract.
type CosmicGameCosmicSignatureAddressChangedIterator struct {
	Event *CosmicGameCosmicSignatureAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameCosmicSignatureAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameCosmicSignatureAddressChanged)
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
		it.Event = new(CosmicGameCosmicSignatureAddressChanged)
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
func (it *CosmicGameCosmicSignatureAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameCosmicSignatureAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameCosmicSignatureAddressChanged represents a CosmicSignatureAddressChanged event raised by the CosmicGame contract.
type CosmicGameCosmicSignatureAddressChanged struct {
	NewCosmicSignature common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCosmicSignatureAddressChanged is a free log retrieval operation binding the contract event 0x7142a592d5404b5fc4a294c66f70e32b2a7776bd807f722d59268def87c765d1.
//
// Solidity: event CosmicSignatureAddressChanged(address newCosmicSignature)
func (_CosmicGame *CosmicGameFilterer) FilterCosmicSignatureAddressChanged(opts *bind.FilterOpts) (*CosmicGameCosmicSignatureAddressChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "CosmicSignatureAddressChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameCosmicSignatureAddressChangedIterator{contract: _CosmicGame.contract, event: "CosmicSignatureAddressChanged", logs: logs, sub: sub}, nil
}

// WatchCosmicSignatureAddressChanged is a free log subscription operation binding the contract event 0x7142a592d5404b5fc4a294c66f70e32b2a7776bd807f722d59268def87c765d1.
//
// Solidity: event CosmicSignatureAddressChanged(address newCosmicSignature)
func (_CosmicGame *CosmicGameFilterer) WatchCosmicSignatureAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameCosmicSignatureAddressChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "CosmicSignatureAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameCosmicSignatureAddressChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "CosmicSignatureAddressChanged", log); err != nil {
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

// ParseCosmicSignatureAddressChanged is a log parse operation binding the contract event 0x7142a592d5404b5fc4a294c66f70e32b2a7776bd807f722d59268def87c765d1.
//
// Solidity: event CosmicSignatureAddressChanged(address newCosmicSignature)
func (_CosmicGame *CosmicGameFilterer) ParseCosmicSignatureAddressChanged(log types.Log) (*CosmicGameCosmicSignatureAddressChanged, error) {
	event := new(CosmicGameCosmicSignatureAddressChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "CosmicSignatureAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameCosmicTokenAddressChangedIterator is returned from FilterCosmicTokenAddressChanged and is used to iterate over the raw logs and unpacked data for CosmicTokenAddressChanged events raised by the CosmicGame contract.
type CosmicGameCosmicTokenAddressChangedIterator struct {
	Event *CosmicGameCosmicTokenAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameCosmicTokenAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameCosmicTokenAddressChanged)
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
		it.Event = new(CosmicGameCosmicTokenAddressChanged)
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
func (it *CosmicGameCosmicTokenAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameCosmicTokenAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameCosmicTokenAddressChanged represents a CosmicTokenAddressChanged event raised by the CosmicGame contract.
type CosmicGameCosmicTokenAddressChanged struct {
	NewCosmicToken common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCosmicTokenAddressChanged is a free log retrieval operation binding the contract event 0x3ab2dcf539aa3c30314265602fc86dee8e343e3c3905730f870835f36c400deb.
//
// Solidity: event CosmicTokenAddressChanged(address newCosmicToken)
func (_CosmicGame *CosmicGameFilterer) FilterCosmicTokenAddressChanged(opts *bind.FilterOpts) (*CosmicGameCosmicTokenAddressChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "CosmicTokenAddressChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameCosmicTokenAddressChangedIterator{contract: _CosmicGame.contract, event: "CosmicTokenAddressChanged", logs: logs, sub: sub}, nil
}

// WatchCosmicTokenAddressChanged is a free log subscription operation binding the contract event 0x3ab2dcf539aa3c30314265602fc86dee8e343e3c3905730f870835f36c400deb.
//
// Solidity: event CosmicTokenAddressChanged(address newCosmicToken)
func (_CosmicGame *CosmicGameFilterer) WatchCosmicTokenAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameCosmicTokenAddressChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "CosmicTokenAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameCosmicTokenAddressChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "CosmicTokenAddressChanged", log); err != nil {
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

// ParseCosmicTokenAddressChanged is a log parse operation binding the contract event 0x3ab2dcf539aa3c30314265602fc86dee8e343e3c3905730f870835f36c400deb.
//
// Solidity: event CosmicTokenAddressChanged(address newCosmicToken)
func (_CosmicGame *CosmicGameFilterer) ParseCosmicTokenAddressChanged(log types.Log) (*CosmicGameCosmicTokenAddressChanged, error) {
	event := new(CosmicGameCosmicTokenAddressChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "CosmicTokenAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameDonatedNFTClaimedEventIterator is returned from FilterDonatedNFTClaimedEvent and is used to iterate over the raw logs and unpacked data for DonatedNFTClaimedEvent events raised by the CosmicGame contract.
type CosmicGameDonatedNFTClaimedEventIterator struct {
	Event *CosmicGameDonatedNFTClaimedEvent // Event containing the contract specifics and raw log

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
func (it *CosmicGameDonatedNFTClaimedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameDonatedNFTClaimedEvent)
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
		it.Event = new(CosmicGameDonatedNFTClaimedEvent)
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
func (it *CosmicGameDonatedNFTClaimedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameDonatedNFTClaimedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameDonatedNFTClaimedEvent represents a DonatedNFTClaimedEvent event raised by the CosmicGame contract.
type CosmicGameDonatedNFTClaimedEvent struct {
	Round                 *big.Int
	Index                 *big.Int
	Winner                common.Address
	NftAddressdonatedNFTs common.Address
	TokenId               *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterDonatedNFTClaimedEvent is a free log retrieval operation binding the contract event 0x0d0e4b90a96d931964d5b0421a6b5b48dd73abf538cdb9ae015069d49e9a4679.
//
// Solidity: event DonatedNFTClaimedEvent(uint256 indexed round, uint256 index, address winner, address nftAddressdonatedNFTs, uint256 tokenId)
func (_CosmicGame *CosmicGameFilterer) FilterDonatedNFTClaimedEvent(opts *bind.FilterOpts, round []*big.Int) (*CosmicGameDonatedNFTClaimedEventIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "DonatedNFTClaimedEvent", roundRule)
	if err != nil {
		return nil, err
	}
	return &CosmicGameDonatedNFTClaimedEventIterator{contract: _CosmicGame.contract, event: "DonatedNFTClaimedEvent", logs: logs, sub: sub}, nil
}

// WatchDonatedNFTClaimedEvent is a free log subscription operation binding the contract event 0x0d0e4b90a96d931964d5b0421a6b5b48dd73abf538cdb9ae015069d49e9a4679.
//
// Solidity: event DonatedNFTClaimedEvent(uint256 indexed round, uint256 index, address winner, address nftAddressdonatedNFTs, uint256 tokenId)
func (_CosmicGame *CosmicGameFilterer) WatchDonatedNFTClaimedEvent(opts *bind.WatchOpts, sink chan<- *CosmicGameDonatedNFTClaimedEvent, round []*big.Int) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "DonatedNFTClaimedEvent", roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameDonatedNFTClaimedEvent)
				if err := _CosmicGame.contract.UnpackLog(event, "DonatedNFTClaimedEvent", log); err != nil {
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

// ParseDonatedNFTClaimedEvent is a log parse operation binding the contract event 0x0d0e4b90a96d931964d5b0421a6b5b48dd73abf538cdb9ae015069d49e9a4679.
//
// Solidity: event DonatedNFTClaimedEvent(uint256 indexed round, uint256 index, address winner, address nftAddressdonatedNFTs, uint256 tokenId)
func (_CosmicGame *CosmicGameFilterer) ParseDonatedNFTClaimedEvent(log types.Log) (*CosmicGameDonatedNFTClaimedEvent, error) {
	event := new(CosmicGameDonatedNFTClaimedEvent)
	if err := _CosmicGame.contract.UnpackLog(event, "DonatedNFTClaimedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameDonationEventIterator is returned from FilterDonationEvent and is used to iterate over the raw logs and unpacked data for DonationEvent events raised by the CosmicGame contract.
type CosmicGameDonationEventIterator struct {
	Event *CosmicGameDonationEvent // Event containing the contract specifics and raw log

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
func (it *CosmicGameDonationEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameDonationEvent)
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
		it.Event = new(CosmicGameDonationEvent)
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
func (it *CosmicGameDonationEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameDonationEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameDonationEvent represents a DonationEvent event raised by the CosmicGame contract.
type CosmicGameDonationEvent struct {
	Donor  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDonationEvent is a free log retrieval operation binding the contract event 0x8b7fe5be5699654fd637d2250cb0d47e88205730710745e78e9d8bcaf8aad8f1.
//
// Solidity: event DonationEvent(address indexed donor, uint256 amount)
func (_CosmicGame *CosmicGameFilterer) FilterDonationEvent(opts *bind.FilterOpts, donor []common.Address) (*CosmicGameDonationEventIterator, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "DonationEvent", donorRule)
	if err != nil {
		return nil, err
	}
	return &CosmicGameDonationEventIterator{contract: _CosmicGame.contract, event: "DonationEvent", logs: logs, sub: sub}, nil
}

// WatchDonationEvent is a free log subscription operation binding the contract event 0x8b7fe5be5699654fd637d2250cb0d47e88205730710745e78e9d8bcaf8aad8f1.
//
// Solidity: event DonationEvent(address indexed donor, uint256 amount)
func (_CosmicGame *CosmicGameFilterer) WatchDonationEvent(opts *bind.WatchOpts, sink chan<- *CosmicGameDonationEvent, donor []common.Address) (event.Subscription, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "DonationEvent", donorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameDonationEvent)
				if err := _CosmicGame.contract.UnpackLog(event, "DonationEvent", log); err != nil {
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

// ParseDonationEvent is a log parse operation binding the contract event 0x8b7fe5be5699654fd637d2250cb0d47e88205730710745e78e9d8bcaf8aad8f1.
//
// Solidity: event DonationEvent(address indexed donor, uint256 amount)
func (_CosmicGame *CosmicGameFilterer) ParseDonationEvent(log types.Log) (*CosmicGameDonationEvent, error) {
	event := new(CosmicGameDonationEvent)
	if err := _CosmicGame.contract.UnpackLog(event, "DonationEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameInitialBidAmountFractionChangedIterator is returned from FilterInitialBidAmountFractionChanged and is used to iterate over the raw logs and unpacked data for InitialBidAmountFractionChanged events raised by the CosmicGame contract.
type CosmicGameInitialBidAmountFractionChangedIterator struct {
	Event *CosmicGameInitialBidAmountFractionChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameInitialBidAmountFractionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameInitialBidAmountFractionChanged)
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
		it.Event = new(CosmicGameInitialBidAmountFractionChanged)
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
func (it *CosmicGameInitialBidAmountFractionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameInitialBidAmountFractionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameInitialBidAmountFractionChanged represents a InitialBidAmountFractionChanged event raised by the CosmicGame contract.
type CosmicGameInitialBidAmountFractionChanged struct {
	NewInitialBidAmountFraction *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterInitialBidAmountFractionChanged is a free log retrieval operation binding the contract event 0x3b311f029da1b90c3de2e0b3168436c5ed7d8b0ae81b7d4894c12da03835c628.
//
// Solidity: event InitialBidAmountFractionChanged(uint256 newInitialBidAmountFraction)
func (_CosmicGame *CosmicGameFilterer) FilterInitialBidAmountFractionChanged(opts *bind.FilterOpts) (*CosmicGameInitialBidAmountFractionChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "InitialBidAmountFractionChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameInitialBidAmountFractionChangedIterator{contract: _CosmicGame.contract, event: "InitialBidAmountFractionChanged", logs: logs, sub: sub}, nil
}

// WatchInitialBidAmountFractionChanged is a free log subscription operation binding the contract event 0x3b311f029da1b90c3de2e0b3168436c5ed7d8b0ae81b7d4894c12da03835c628.
//
// Solidity: event InitialBidAmountFractionChanged(uint256 newInitialBidAmountFraction)
func (_CosmicGame *CosmicGameFilterer) WatchInitialBidAmountFractionChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameInitialBidAmountFractionChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "InitialBidAmountFractionChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameInitialBidAmountFractionChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "InitialBidAmountFractionChanged", log); err != nil {
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

// ParseInitialBidAmountFractionChanged is a log parse operation binding the contract event 0x3b311f029da1b90c3de2e0b3168436c5ed7d8b0ae81b7d4894c12da03835c628.
//
// Solidity: event InitialBidAmountFractionChanged(uint256 newInitialBidAmountFraction)
func (_CosmicGame *CosmicGameFilterer) ParseInitialBidAmountFractionChanged(log types.Log) (*CosmicGameInitialBidAmountFractionChanged, error) {
	event := new(CosmicGameInitialBidAmountFractionChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "InitialBidAmountFractionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameInitialSecondsUntilPrizeChangedIterator is returned from FilterInitialSecondsUntilPrizeChanged and is used to iterate over the raw logs and unpacked data for InitialSecondsUntilPrizeChanged events raised by the CosmicGame contract.
type CosmicGameInitialSecondsUntilPrizeChangedIterator struct {
	Event *CosmicGameInitialSecondsUntilPrizeChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameInitialSecondsUntilPrizeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameInitialSecondsUntilPrizeChanged)
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
		it.Event = new(CosmicGameInitialSecondsUntilPrizeChanged)
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
func (it *CosmicGameInitialSecondsUntilPrizeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameInitialSecondsUntilPrizeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameInitialSecondsUntilPrizeChanged represents a InitialSecondsUntilPrizeChanged event raised by the CosmicGame contract.
type CosmicGameInitialSecondsUntilPrizeChanged struct {
	NewInitialSecondsUntilPrize *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterInitialSecondsUntilPrizeChanged is a free log retrieval operation binding the contract event 0x6da281754ba85ee0c5983a8e8f05a92910c2a0c5b80e68c126216d65f162a305.
//
// Solidity: event InitialSecondsUntilPrizeChanged(uint256 newInitialSecondsUntilPrize)
func (_CosmicGame *CosmicGameFilterer) FilterInitialSecondsUntilPrizeChanged(opts *bind.FilterOpts) (*CosmicGameInitialSecondsUntilPrizeChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "InitialSecondsUntilPrizeChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameInitialSecondsUntilPrizeChangedIterator{contract: _CosmicGame.contract, event: "InitialSecondsUntilPrizeChanged", logs: logs, sub: sub}, nil
}

// WatchInitialSecondsUntilPrizeChanged is a free log subscription operation binding the contract event 0x6da281754ba85ee0c5983a8e8f05a92910c2a0c5b80e68c126216d65f162a305.
//
// Solidity: event InitialSecondsUntilPrizeChanged(uint256 newInitialSecondsUntilPrize)
func (_CosmicGame *CosmicGameFilterer) WatchInitialSecondsUntilPrizeChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameInitialSecondsUntilPrizeChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "InitialSecondsUntilPrizeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameInitialSecondsUntilPrizeChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "InitialSecondsUntilPrizeChanged", log); err != nil {
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

// ParseInitialSecondsUntilPrizeChanged is a log parse operation binding the contract event 0x6da281754ba85ee0c5983a8e8f05a92910c2a0c5b80e68c126216d65f162a305.
//
// Solidity: event InitialSecondsUntilPrizeChanged(uint256 newInitialSecondsUntilPrize)
func (_CosmicGame *CosmicGameFilterer) ParseInitialSecondsUntilPrizeChanged(log types.Log) (*CosmicGameInitialSecondsUntilPrizeChanged, error) {
	event := new(CosmicGameInitialSecondsUntilPrizeChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "InitialSecondsUntilPrizeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameNFTDonationEventIterator is returned from FilterNFTDonationEvent and is used to iterate over the raw logs and unpacked data for NFTDonationEvent events raised by the CosmicGame contract.
type CosmicGameNFTDonationEventIterator struct {
	Event *CosmicGameNFTDonationEvent // Event containing the contract specifics and raw log

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
func (it *CosmicGameNFTDonationEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameNFTDonationEvent)
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
		it.Event = new(CosmicGameNFTDonationEvent)
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
func (it *CosmicGameNFTDonationEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameNFTDonationEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameNFTDonationEvent represents a NFTDonationEvent event raised by the CosmicGame contract.
type CosmicGameNFTDonationEvent struct {
	Donor      common.Address
	NftAddress common.Address
	Round      *big.Int
	TokenId    *big.Int
	Index      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNFTDonationEvent is a free log retrieval operation binding the contract event 0xc85be1734ed6a0f4d6adf924d4d41406e2729878c652110a96e2fdec64e118d1.
//
// Solidity: event NFTDonationEvent(address indexed donor, address indexed nftAddress, uint256 indexed round, uint256 tokenId, uint256 index)
func (_CosmicGame *CosmicGameFilterer) FilterNFTDonationEvent(opts *bind.FilterOpts, donor []common.Address, nftAddress []common.Address, round []*big.Int) (*CosmicGameNFTDonationEventIterator, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "NFTDonationEvent", donorRule, nftAddressRule, roundRule)
	if err != nil {
		return nil, err
	}
	return &CosmicGameNFTDonationEventIterator{contract: _CosmicGame.contract, event: "NFTDonationEvent", logs: logs, sub: sub}, nil
}

// WatchNFTDonationEvent is a free log subscription operation binding the contract event 0xc85be1734ed6a0f4d6adf924d4d41406e2729878c652110a96e2fdec64e118d1.
//
// Solidity: event NFTDonationEvent(address indexed donor, address indexed nftAddress, uint256 indexed round, uint256 tokenId, uint256 index)
func (_CosmicGame *CosmicGameFilterer) WatchNFTDonationEvent(opts *bind.WatchOpts, sink chan<- *CosmicGameNFTDonationEvent, donor []common.Address, nftAddress []common.Address, round []*big.Int) (event.Subscription, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}
	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "NFTDonationEvent", donorRule, nftAddressRule, roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameNFTDonationEvent)
				if err := _CosmicGame.contract.UnpackLog(event, "NFTDonationEvent", log); err != nil {
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

// ParseNFTDonationEvent is a log parse operation binding the contract event 0xc85be1734ed6a0f4d6adf924d4d41406e2729878c652110a96e2fdec64e118d1.
//
// Solidity: event NFTDonationEvent(address indexed donor, address indexed nftAddress, uint256 indexed round, uint256 tokenId, uint256 index)
func (_CosmicGame *CosmicGameFilterer) ParseNFTDonationEvent(log types.Log) (*CosmicGameNFTDonationEvent, error) {
	event := new(CosmicGameNFTDonationEvent)
	if err := _CosmicGame.contract.UnpackLog(event, "NFTDonationEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameNanoSecondsExtraChangedIterator is returned from FilterNanoSecondsExtraChanged and is used to iterate over the raw logs and unpacked data for NanoSecondsExtraChanged events raised by the CosmicGame contract.
type CosmicGameNanoSecondsExtraChangedIterator struct {
	Event *CosmicGameNanoSecondsExtraChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameNanoSecondsExtraChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameNanoSecondsExtraChanged)
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
		it.Event = new(CosmicGameNanoSecondsExtraChanged)
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
func (it *CosmicGameNanoSecondsExtraChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameNanoSecondsExtraChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameNanoSecondsExtraChanged represents a NanoSecondsExtraChanged event raised by the CosmicGame contract.
type CosmicGameNanoSecondsExtraChanged struct {
	NewNanoSecondsExtra *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNanoSecondsExtraChanged is a free log retrieval operation binding the contract event 0x678d086a09e1be49401b1e3a6e14db1878e8d8b88e71d0cfe24a32726d0e38b9.
//
// Solidity: event NanoSecondsExtraChanged(uint256 newNanoSecondsExtra)
func (_CosmicGame *CosmicGameFilterer) FilterNanoSecondsExtraChanged(opts *bind.FilterOpts) (*CosmicGameNanoSecondsExtraChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "NanoSecondsExtraChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameNanoSecondsExtraChangedIterator{contract: _CosmicGame.contract, event: "NanoSecondsExtraChanged", logs: logs, sub: sub}, nil
}

// WatchNanoSecondsExtraChanged is a free log subscription operation binding the contract event 0x678d086a09e1be49401b1e3a6e14db1878e8d8b88e71d0cfe24a32726d0e38b9.
//
// Solidity: event NanoSecondsExtraChanged(uint256 newNanoSecondsExtra)
func (_CosmicGame *CosmicGameFilterer) WatchNanoSecondsExtraChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameNanoSecondsExtraChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "NanoSecondsExtraChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameNanoSecondsExtraChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "NanoSecondsExtraChanged", log); err != nil {
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

// ParseNanoSecondsExtraChanged is a log parse operation binding the contract event 0x678d086a09e1be49401b1e3a6e14db1878e8d8b88e71d0cfe24a32726d0e38b9.
//
// Solidity: event NanoSecondsExtraChanged(uint256 newNanoSecondsExtra)
func (_CosmicGame *CosmicGameFilterer) ParseNanoSecondsExtraChanged(log types.Log) (*CosmicGameNanoSecondsExtraChanged, error) {
	event := new(CosmicGameNanoSecondsExtraChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "NanoSecondsExtraChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameNumRaffleNFTWinnersPerRoundChangedIterator is returned from FilterNumRaffleNFTWinnersPerRoundChanged and is used to iterate over the raw logs and unpacked data for NumRaffleNFTWinnersPerRoundChanged events raised by the CosmicGame contract.
type CosmicGameNumRaffleNFTWinnersPerRoundChangedIterator struct {
	Event *CosmicGameNumRaffleNFTWinnersPerRoundChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameNumRaffleNFTWinnersPerRoundChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameNumRaffleNFTWinnersPerRoundChanged)
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
		it.Event = new(CosmicGameNumRaffleNFTWinnersPerRoundChanged)
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
func (it *CosmicGameNumRaffleNFTWinnersPerRoundChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameNumRaffleNFTWinnersPerRoundChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameNumRaffleNFTWinnersPerRoundChanged represents a NumRaffleNFTWinnersPerRoundChanged event raised by the CosmicGame contract.
type CosmicGameNumRaffleNFTWinnersPerRoundChanged struct {
	NewNumRaffleNFTWinnersPerRound *big.Int
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterNumRaffleNFTWinnersPerRoundChanged is a free log retrieval operation binding the contract event 0x72e4278828b8a868e0ba5b4887f954797ea786f8bac991128636171f5eed471d.
//
// Solidity: event NumRaffleNFTWinnersPerRoundChanged(uint256 newNumRaffleNFTWinnersPerRound)
func (_CosmicGame *CosmicGameFilterer) FilterNumRaffleNFTWinnersPerRoundChanged(opts *bind.FilterOpts) (*CosmicGameNumRaffleNFTWinnersPerRoundChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "NumRaffleNFTWinnersPerRoundChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameNumRaffleNFTWinnersPerRoundChangedIterator{contract: _CosmicGame.contract, event: "NumRaffleNFTWinnersPerRoundChanged", logs: logs, sub: sub}, nil
}

// WatchNumRaffleNFTWinnersPerRoundChanged is a free log subscription operation binding the contract event 0x72e4278828b8a868e0ba5b4887f954797ea786f8bac991128636171f5eed471d.
//
// Solidity: event NumRaffleNFTWinnersPerRoundChanged(uint256 newNumRaffleNFTWinnersPerRound)
func (_CosmicGame *CosmicGameFilterer) WatchNumRaffleNFTWinnersPerRoundChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameNumRaffleNFTWinnersPerRoundChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "NumRaffleNFTWinnersPerRoundChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameNumRaffleNFTWinnersPerRoundChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "NumRaffleNFTWinnersPerRoundChanged", log); err != nil {
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

// ParseNumRaffleNFTWinnersPerRoundChanged is a log parse operation binding the contract event 0x72e4278828b8a868e0ba5b4887f954797ea786f8bac991128636171f5eed471d.
//
// Solidity: event NumRaffleNFTWinnersPerRoundChanged(uint256 newNumRaffleNFTWinnersPerRound)
func (_CosmicGame *CosmicGameFilterer) ParseNumRaffleNFTWinnersPerRoundChanged(log types.Log) (*CosmicGameNumRaffleNFTWinnersPerRoundChanged, error) {
	event := new(CosmicGameNumRaffleNFTWinnersPerRoundChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "NumRaffleNFTWinnersPerRoundChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameNumRaffleWinnersPerRoundChangedIterator is returned from FilterNumRaffleWinnersPerRoundChanged and is used to iterate over the raw logs and unpacked data for NumRaffleWinnersPerRoundChanged events raised by the CosmicGame contract.
type CosmicGameNumRaffleWinnersPerRoundChangedIterator struct {
	Event *CosmicGameNumRaffleWinnersPerRoundChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameNumRaffleWinnersPerRoundChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameNumRaffleWinnersPerRoundChanged)
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
		it.Event = new(CosmicGameNumRaffleWinnersPerRoundChanged)
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
func (it *CosmicGameNumRaffleWinnersPerRoundChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameNumRaffleWinnersPerRoundChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameNumRaffleWinnersPerRoundChanged represents a NumRaffleWinnersPerRoundChanged event raised by the CosmicGame contract.
type CosmicGameNumRaffleWinnersPerRoundChanged struct {
	NewNumRaffleWinnersPerRound *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterNumRaffleWinnersPerRoundChanged is a free log retrieval operation binding the contract event 0x5e3567ae359e61c47ad2f11d2828dea7e062b2ff63dea385cdb7900a50504c7c.
//
// Solidity: event NumRaffleWinnersPerRoundChanged(uint256 newNumRaffleWinnersPerRound)
func (_CosmicGame *CosmicGameFilterer) FilterNumRaffleWinnersPerRoundChanged(opts *bind.FilterOpts) (*CosmicGameNumRaffleWinnersPerRoundChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "NumRaffleWinnersPerRoundChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameNumRaffleWinnersPerRoundChangedIterator{contract: _CosmicGame.contract, event: "NumRaffleWinnersPerRoundChanged", logs: logs, sub: sub}, nil
}

// WatchNumRaffleWinnersPerRoundChanged is a free log subscription operation binding the contract event 0x5e3567ae359e61c47ad2f11d2828dea7e062b2ff63dea385cdb7900a50504c7c.
//
// Solidity: event NumRaffleWinnersPerRoundChanged(uint256 newNumRaffleWinnersPerRound)
func (_CosmicGame *CosmicGameFilterer) WatchNumRaffleWinnersPerRoundChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameNumRaffleWinnersPerRoundChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "NumRaffleWinnersPerRoundChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameNumRaffleWinnersPerRoundChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "NumRaffleWinnersPerRoundChanged", log); err != nil {
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

// ParseNumRaffleWinnersPerRoundChanged is a log parse operation binding the contract event 0x5e3567ae359e61c47ad2f11d2828dea7e062b2ff63dea385cdb7900a50504c7c.
//
// Solidity: event NumRaffleWinnersPerRoundChanged(uint256 newNumRaffleWinnersPerRound)
func (_CosmicGame *CosmicGameFilterer) ParseNumRaffleWinnersPerRoundChanged(log types.Log) (*CosmicGameNumRaffleWinnersPerRoundChanged, error) {
	event := new(CosmicGameNumRaffleWinnersPerRoundChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "NumRaffleWinnersPerRoundChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CosmicGame contract.
type CosmicGameOwnershipTransferredIterator struct {
	Event *CosmicGameOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CosmicGameOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameOwnershipTransferred)
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
		it.Event = new(CosmicGameOwnershipTransferred)
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
func (it *CosmicGameOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameOwnershipTransferred represents a OwnershipTransferred event raised by the CosmicGame contract.
type CosmicGameOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CosmicGame *CosmicGameFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CosmicGameOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CosmicGameOwnershipTransferredIterator{contract: _CosmicGame.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CosmicGame *CosmicGameFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CosmicGameOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameOwnershipTransferred)
				if err := _CosmicGame.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CosmicGame *CosmicGameFilterer) ParseOwnershipTransferred(log types.Log) (*CosmicGameOwnershipTransferred, error) {
	event := new(CosmicGameOwnershipTransferred)
	if err := _CosmicGame.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGamePriceIncreaseChangedIterator is returned from FilterPriceIncreaseChanged and is used to iterate over the raw logs and unpacked data for PriceIncreaseChanged events raised by the CosmicGame contract.
type CosmicGamePriceIncreaseChangedIterator struct {
	Event *CosmicGamePriceIncreaseChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGamePriceIncreaseChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGamePriceIncreaseChanged)
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
		it.Event = new(CosmicGamePriceIncreaseChanged)
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
func (it *CosmicGamePriceIncreaseChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGamePriceIncreaseChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGamePriceIncreaseChanged represents a PriceIncreaseChanged event raised by the CosmicGame contract.
type CosmicGamePriceIncreaseChanged struct {
	NewPriceIncrease *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPriceIncreaseChanged is a free log retrieval operation binding the contract event 0xcbc1f49adfa29e2f2f0f5c9e057722496a4bc95a6a5446deaa423a02b30c64ac.
//
// Solidity: event PriceIncreaseChanged(uint256 newPriceIncrease)
func (_CosmicGame *CosmicGameFilterer) FilterPriceIncreaseChanged(opts *bind.FilterOpts) (*CosmicGamePriceIncreaseChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "PriceIncreaseChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGamePriceIncreaseChangedIterator{contract: _CosmicGame.contract, event: "PriceIncreaseChanged", logs: logs, sub: sub}, nil
}

// WatchPriceIncreaseChanged is a free log subscription operation binding the contract event 0xcbc1f49adfa29e2f2f0f5c9e057722496a4bc95a6a5446deaa423a02b30c64ac.
//
// Solidity: event PriceIncreaseChanged(uint256 newPriceIncrease)
func (_CosmicGame *CosmicGameFilterer) WatchPriceIncreaseChanged(opts *bind.WatchOpts, sink chan<- *CosmicGamePriceIncreaseChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "PriceIncreaseChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGamePriceIncreaseChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "PriceIncreaseChanged", log); err != nil {
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

// ParsePriceIncreaseChanged is a log parse operation binding the contract event 0xcbc1f49adfa29e2f2f0f5c9e057722496a4bc95a6a5446deaa423a02b30c64ac.
//
// Solidity: event PriceIncreaseChanged(uint256 newPriceIncrease)
func (_CosmicGame *CosmicGameFilterer) ParsePriceIncreaseChanged(log types.Log) (*CosmicGamePriceIncreaseChanged, error) {
	event := new(CosmicGamePriceIncreaseChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "PriceIncreaseChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGamePrizeClaimEventIterator is returned from FilterPrizeClaimEvent and is used to iterate over the raw logs and unpacked data for PrizeClaimEvent events raised by the CosmicGame contract.
type CosmicGamePrizeClaimEventIterator struct {
	Event *CosmicGamePrizeClaimEvent // Event containing the contract specifics and raw log

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
func (it *CosmicGamePrizeClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGamePrizeClaimEvent)
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
		it.Event = new(CosmicGamePrizeClaimEvent)
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
func (it *CosmicGamePrizeClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGamePrizeClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGamePrizeClaimEvent represents a PrizeClaimEvent event raised by the CosmicGame contract.
type CosmicGamePrizeClaimEvent struct {
	PrizeNum    *big.Int
	Destination common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPrizeClaimEvent is a free log retrieval operation binding the contract event 0x27bc828c399c2947fea27bca8a75ced2e94ff2651d607271f051e39db52286ce.
//
// Solidity: event PrizeClaimEvent(uint256 indexed prizeNum, address indexed destination, uint256 amount)
func (_CosmicGame *CosmicGameFilterer) FilterPrizeClaimEvent(opts *bind.FilterOpts, prizeNum []*big.Int, destination []common.Address) (*CosmicGamePrizeClaimEventIterator, error) {

	var prizeNumRule []interface{}
	for _, prizeNumItem := range prizeNum {
		prizeNumRule = append(prizeNumRule, prizeNumItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "PrizeClaimEvent", prizeNumRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return &CosmicGamePrizeClaimEventIterator{contract: _CosmicGame.contract, event: "PrizeClaimEvent", logs: logs, sub: sub}, nil
}

// WatchPrizeClaimEvent is a free log subscription operation binding the contract event 0x27bc828c399c2947fea27bca8a75ced2e94ff2651d607271f051e39db52286ce.
//
// Solidity: event PrizeClaimEvent(uint256 indexed prizeNum, address indexed destination, uint256 amount)
func (_CosmicGame *CosmicGameFilterer) WatchPrizeClaimEvent(opts *bind.WatchOpts, sink chan<- *CosmicGamePrizeClaimEvent, prizeNum []*big.Int, destination []common.Address) (event.Subscription, error) {

	var prizeNumRule []interface{}
	for _, prizeNumItem := range prizeNum {
		prizeNumRule = append(prizeNumRule, prizeNumItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "PrizeClaimEvent", prizeNumRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGamePrizeClaimEvent)
				if err := _CosmicGame.contract.UnpackLog(event, "PrizeClaimEvent", log); err != nil {
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

// ParsePrizeClaimEvent is a log parse operation binding the contract event 0x27bc828c399c2947fea27bca8a75ced2e94ff2651d607271f051e39db52286ce.
//
// Solidity: event PrizeClaimEvent(uint256 indexed prizeNum, address indexed destination, uint256 amount)
func (_CosmicGame *CosmicGameFilterer) ParsePrizeClaimEvent(log types.Log) (*CosmicGamePrizeClaimEvent, error) {
	event := new(CosmicGamePrizeClaimEvent)
	if err := _CosmicGame.contract.UnpackLog(event, "PrizeClaimEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGamePrizePercentageChangedIterator is returned from FilterPrizePercentageChanged and is used to iterate over the raw logs and unpacked data for PrizePercentageChanged events raised by the CosmicGame contract.
type CosmicGamePrizePercentageChangedIterator struct {
	Event *CosmicGamePrizePercentageChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGamePrizePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGamePrizePercentageChanged)
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
		it.Event = new(CosmicGamePrizePercentageChanged)
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
func (it *CosmicGamePrizePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGamePrizePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGamePrizePercentageChanged represents a PrizePercentageChanged event raised by the CosmicGame contract.
type CosmicGamePrizePercentageChanged struct {
	NewPrizePercentage *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPrizePercentageChanged is a free log retrieval operation binding the contract event 0x595fa5ba64cd6f66df19b53b59bb4a275bae1ec5b362da95e1fa4c5feb62be1e.
//
// Solidity: event PrizePercentageChanged(uint256 newPrizePercentage)
func (_CosmicGame *CosmicGameFilterer) FilterPrizePercentageChanged(opts *bind.FilterOpts) (*CosmicGamePrizePercentageChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "PrizePercentageChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGamePrizePercentageChangedIterator{contract: _CosmicGame.contract, event: "PrizePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchPrizePercentageChanged is a free log subscription operation binding the contract event 0x595fa5ba64cd6f66df19b53b59bb4a275bae1ec5b362da95e1fa4c5feb62be1e.
//
// Solidity: event PrizePercentageChanged(uint256 newPrizePercentage)
func (_CosmicGame *CosmicGameFilterer) WatchPrizePercentageChanged(opts *bind.WatchOpts, sink chan<- *CosmicGamePrizePercentageChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "PrizePercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGamePrizePercentageChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "PrizePercentageChanged", log); err != nil {
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

// ParsePrizePercentageChanged is a log parse operation binding the contract event 0x595fa5ba64cd6f66df19b53b59bb4a275bae1ec5b362da95e1fa4c5feb62be1e.
//
// Solidity: event PrizePercentageChanged(uint256 newPrizePercentage)
func (_CosmicGame *CosmicGameFilterer) ParsePrizePercentageChanged(log types.Log) (*CosmicGamePrizePercentageChanged, error) {
	event := new(CosmicGamePrizePercentageChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "PrizePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameRaffleNFTClaimedEventIterator is returned from FilterRaffleNFTClaimedEvent and is used to iterate over the raw logs and unpacked data for RaffleNFTClaimedEvent events raised by the CosmicGame contract.
type CosmicGameRaffleNFTClaimedEventIterator struct {
	Event *CosmicGameRaffleNFTClaimedEvent // Event containing the contract specifics and raw log

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
func (it *CosmicGameRaffleNFTClaimedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameRaffleNFTClaimedEvent)
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
		it.Event = new(CosmicGameRaffleNFTClaimedEvent)
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
func (it *CosmicGameRaffleNFTClaimedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameRaffleNFTClaimedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameRaffleNFTClaimedEvent represents a RaffleNFTClaimedEvent event raised by the CosmicGame contract.
type CosmicGameRaffleNFTClaimedEvent struct {
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRaffleNFTClaimedEvent is a free log retrieval operation binding the contract event 0xe05ba2c5fcd9a60f30b179cb0e775070cc8ce9667b0e663e984ee6a02f694cee.
//
// Solidity: event RaffleNFTClaimedEvent(address indexed winner)
func (_CosmicGame *CosmicGameFilterer) FilterRaffleNFTClaimedEvent(opts *bind.FilterOpts, winner []common.Address) (*CosmicGameRaffleNFTClaimedEventIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "RaffleNFTClaimedEvent", winnerRule)
	if err != nil {
		return nil, err
	}
	return &CosmicGameRaffleNFTClaimedEventIterator{contract: _CosmicGame.contract, event: "RaffleNFTClaimedEvent", logs: logs, sub: sub}, nil
}

// WatchRaffleNFTClaimedEvent is a free log subscription operation binding the contract event 0xe05ba2c5fcd9a60f30b179cb0e775070cc8ce9667b0e663e984ee6a02f694cee.
//
// Solidity: event RaffleNFTClaimedEvent(address indexed winner)
func (_CosmicGame *CosmicGameFilterer) WatchRaffleNFTClaimedEvent(opts *bind.WatchOpts, sink chan<- *CosmicGameRaffleNFTClaimedEvent, winner []common.Address) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "RaffleNFTClaimedEvent", winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameRaffleNFTClaimedEvent)
				if err := _CosmicGame.contract.UnpackLog(event, "RaffleNFTClaimedEvent", log); err != nil {
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

// ParseRaffleNFTClaimedEvent is a log parse operation binding the contract event 0xe05ba2c5fcd9a60f30b179cb0e775070cc8ce9667b0e663e984ee6a02f694cee.
//
// Solidity: event RaffleNFTClaimedEvent(address indexed winner)
func (_CosmicGame *CosmicGameFilterer) ParseRaffleNFTClaimedEvent(log types.Log) (*CosmicGameRaffleNFTClaimedEvent, error) {
	event := new(CosmicGameRaffleNFTClaimedEvent)
	if err := _CosmicGame.contract.UnpackLog(event, "RaffleNFTClaimedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameRaffleNFTWinnerEventIterator is returned from FilterRaffleNFTWinnerEvent and is used to iterate over the raw logs and unpacked data for RaffleNFTWinnerEvent events raised by the CosmicGame contract.
type CosmicGameRaffleNFTWinnerEventIterator struct {
	Event *CosmicGameRaffleNFTWinnerEvent // Event containing the contract specifics and raw log

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
func (it *CosmicGameRaffleNFTWinnerEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameRaffleNFTWinnerEvent)
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
		it.Event = new(CosmicGameRaffleNFTWinnerEvent)
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
func (it *CosmicGameRaffleNFTWinnerEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameRaffleNFTWinnerEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameRaffleNFTWinnerEvent represents a RaffleNFTWinnerEvent event raised by the CosmicGame contract.
type CosmicGameRaffleNFTWinnerEvent struct {
	Winner      common.Address
	Round       *big.Int
	WinnerIndex *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRaffleNFTWinnerEvent is a free log retrieval operation binding the contract event 0x80348bf864c08069d1368c42ed36b7a60560f73267f63d58e9be69f4b021bacc.
//
// Solidity: event RaffleNFTWinnerEvent(address indexed winner, uint256 indexed round, uint256 winner_index)
func (_CosmicGame *CosmicGameFilterer) FilterRaffleNFTWinnerEvent(opts *bind.FilterOpts, winner []common.Address, round []*big.Int) (*CosmicGameRaffleNFTWinnerEventIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "RaffleNFTWinnerEvent", winnerRule, roundRule)
	if err != nil {
		return nil, err
	}
	return &CosmicGameRaffleNFTWinnerEventIterator{contract: _CosmicGame.contract, event: "RaffleNFTWinnerEvent", logs: logs, sub: sub}, nil
}

// WatchRaffleNFTWinnerEvent is a free log subscription operation binding the contract event 0x80348bf864c08069d1368c42ed36b7a60560f73267f63d58e9be69f4b021bacc.
//
// Solidity: event RaffleNFTWinnerEvent(address indexed winner, uint256 indexed round, uint256 winner_index)
func (_CosmicGame *CosmicGameFilterer) WatchRaffleNFTWinnerEvent(opts *bind.WatchOpts, sink chan<- *CosmicGameRaffleNFTWinnerEvent, winner []common.Address, round []*big.Int) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "RaffleNFTWinnerEvent", winnerRule, roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameRaffleNFTWinnerEvent)
				if err := _CosmicGame.contract.UnpackLog(event, "RaffleNFTWinnerEvent", log); err != nil {
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

// ParseRaffleNFTWinnerEvent is a log parse operation binding the contract event 0x80348bf864c08069d1368c42ed36b7a60560f73267f63d58e9be69f4b021bacc.
//
// Solidity: event RaffleNFTWinnerEvent(address indexed winner, uint256 indexed round, uint256 winner_index)
func (_CosmicGame *CosmicGameFilterer) ParseRaffleNFTWinnerEvent(log types.Log) (*CosmicGameRaffleNFTWinnerEvent, error) {
	event := new(CosmicGameRaffleNFTWinnerEvent)
	if err := _CosmicGame.contract.UnpackLog(event, "RaffleNFTWinnerEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameRafflePercentageChangedIterator is returned from FilterRafflePercentageChanged and is used to iterate over the raw logs and unpacked data for RafflePercentageChanged events raised by the CosmicGame contract.
type CosmicGameRafflePercentageChangedIterator struct {
	Event *CosmicGameRafflePercentageChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameRafflePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameRafflePercentageChanged)
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
		it.Event = new(CosmicGameRafflePercentageChanged)
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
func (it *CosmicGameRafflePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameRafflePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameRafflePercentageChanged represents a RafflePercentageChanged event raised by the CosmicGame contract.
type CosmicGameRafflePercentageChanged struct {
	NewRafflePercentage *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRafflePercentageChanged is a free log retrieval operation binding the contract event 0xd2002a256ab2f8a5c1b865685754b20437c0a26e7249e40ae5df2993966f99f2.
//
// Solidity: event RafflePercentageChanged(uint256 newRafflePercentage)
func (_CosmicGame *CosmicGameFilterer) FilterRafflePercentageChanged(opts *bind.FilterOpts) (*CosmicGameRafflePercentageChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "RafflePercentageChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameRafflePercentageChangedIterator{contract: _CosmicGame.contract, event: "RafflePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchRafflePercentageChanged is a free log subscription operation binding the contract event 0xd2002a256ab2f8a5c1b865685754b20437c0a26e7249e40ae5df2993966f99f2.
//
// Solidity: event RafflePercentageChanged(uint256 newRafflePercentage)
func (_CosmicGame *CosmicGameFilterer) WatchRafflePercentageChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameRafflePercentageChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "RafflePercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameRafflePercentageChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "RafflePercentageChanged", log); err != nil {
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

// ParseRafflePercentageChanged is a log parse operation binding the contract event 0xd2002a256ab2f8a5c1b865685754b20437c0a26e7249e40ae5df2993966f99f2.
//
// Solidity: event RafflePercentageChanged(uint256 newRafflePercentage)
func (_CosmicGame *CosmicGameFilterer) ParseRafflePercentageChanged(log types.Log) (*CosmicGameRafflePercentageChanged, error) {
	event := new(CosmicGameRafflePercentageChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "RafflePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameRaffleWalletAddressChangedIterator is returned from FilterRaffleWalletAddressChanged and is used to iterate over the raw logs and unpacked data for RaffleWalletAddressChanged events raised by the CosmicGame contract.
type CosmicGameRaffleWalletAddressChangedIterator struct {
	Event *CosmicGameRaffleWalletAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameRaffleWalletAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameRaffleWalletAddressChanged)
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
		it.Event = new(CosmicGameRaffleWalletAddressChanged)
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
func (it *CosmicGameRaffleWalletAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameRaffleWalletAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameRaffleWalletAddressChanged represents a RaffleWalletAddressChanged event raised by the CosmicGame contract.
type CosmicGameRaffleWalletAddressChanged struct {
	NewRaffleWallet common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRaffleWalletAddressChanged is a free log retrieval operation binding the contract event 0x508d510537704d37ed64691a706392abc33e59ce8a327b4952e112820ceb88a6.
//
// Solidity: event RaffleWalletAddressChanged(address newRaffleWallet)
func (_CosmicGame *CosmicGameFilterer) FilterRaffleWalletAddressChanged(opts *bind.FilterOpts) (*CosmicGameRaffleWalletAddressChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "RaffleWalletAddressChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameRaffleWalletAddressChangedIterator{contract: _CosmicGame.contract, event: "RaffleWalletAddressChanged", logs: logs, sub: sub}, nil
}

// WatchRaffleWalletAddressChanged is a free log subscription operation binding the contract event 0x508d510537704d37ed64691a706392abc33e59ce8a327b4952e112820ceb88a6.
//
// Solidity: event RaffleWalletAddressChanged(address newRaffleWallet)
func (_CosmicGame *CosmicGameFilterer) WatchRaffleWalletAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameRaffleWalletAddressChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "RaffleWalletAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameRaffleWalletAddressChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "RaffleWalletAddressChanged", log); err != nil {
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

// ParseRaffleWalletAddressChanged is a log parse operation binding the contract event 0x508d510537704d37ed64691a706392abc33e59ce8a327b4952e112820ceb88a6.
//
// Solidity: event RaffleWalletAddressChanged(address newRaffleWallet)
func (_CosmicGame *CosmicGameFilterer) ParseRaffleWalletAddressChanged(log types.Log) (*CosmicGameRaffleWalletAddressChanged, error) {
	event := new(CosmicGameRaffleWalletAddressChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "RaffleWalletAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameRandomWalkAddressChangedIterator is returned from FilterRandomWalkAddressChanged and is used to iterate over the raw logs and unpacked data for RandomWalkAddressChanged events raised by the CosmicGame contract.
type CosmicGameRandomWalkAddressChangedIterator struct {
	Event *CosmicGameRandomWalkAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameRandomWalkAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameRandomWalkAddressChanged)
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
		it.Event = new(CosmicGameRandomWalkAddressChanged)
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
func (it *CosmicGameRandomWalkAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameRandomWalkAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameRandomWalkAddressChanged represents a RandomWalkAddressChanged event raised by the CosmicGame contract.
type CosmicGameRandomWalkAddressChanged struct {
	NewRandomWalk common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRandomWalkAddressChanged is a free log retrieval operation binding the contract event 0x9dc3707e4b88be65295b0083b44ffa94059c80372add9b0d24d6a3b371a03b92.
//
// Solidity: event RandomWalkAddressChanged(address newRandomWalk)
func (_CosmicGame *CosmicGameFilterer) FilterRandomWalkAddressChanged(opts *bind.FilterOpts) (*CosmicGameRandomWalkAddressChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "RandomWalkAddressChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameRandomWalkAddressChangedIterator{contract: _CosmicGame.contract, event: "RandomWalkAddressChanged", logs: logs, sub: sub}, nil
}

// WatchRandomWalkAddressChanged is a free log subscription operation binding the contract event 0x9dc3707e4b88be65295b0083b44ffa94059c80372add9b0d24d6a3b371a03b92.
//
// Solidity: event RandomWalkAddressChanged(address newRandomWalk)
func (_CosmicGame *CosmicGameFilterer) WatchRandomWalkAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameRandomWalkAddressChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "RandomWalkAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameRandomWalkAddressChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "RandomWalkAddressChanged", log); err != nil {
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

// ParseRandomWalkAddressChanged is a log parse operation binding the contract event 0x9dc3707e4b88be65295b0083b44ffa94059c80372add9b0d24d6a3b371a03b92.
//
// Solidity: event RandomWalkAddressChanged(address newRandomWalk)
func (_CosmicGame *CosmicGameFilterer) ParseRandomWalkAddressChanged(log types.Log) (*CosmicGameRandomWalkAddressChanged, error) {
	event := new(CosmicGameRandomWalkAddressChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "RandomWalkAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameTimeIncreaseChangedIterator is returned from FilterTimeIncreaseChanged and is used to iterate over the raw logs and unpacked data for TimeIncreaseChanged events raised by the CosmicGame contract.
type CosmicGameTimeIncreaseChangedIterator struct {
	Event *CosmicGameTimeIncreaseChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameTimeIncreaseChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameTimeIncreaseChanged)
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
		it.Event = new(CosmicGameTimeIncreaseChanged)
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
func (it *CosmicGameTimeIncreaseChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameTimeIncreaseChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameTimeIncreaseChanged represents a TimeIncreaseChanged event raised by the CosmicGame contract.
type CosmicGameTimeIncreaseChanged struct {
	NewTimeIncrease *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimeIncreaseChanged is a free log retrieval operation binding the contract event 0xed46e73b174ced51fb529cdc6c1b4d8abf49387e6d849b71648afb63c81d12cd.
//
// Solidity: event TimeIncreaseChanged(uint256 newTimeIncrease)
func (_CosmicGame *CosmicGameFilterer) FilterTimeIncreaseChanged(opts *bind.FilterOpts) (*CosmicGameTimeIncreaseChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "TimeIncreaseChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameTimeIncreaseChangedIterator{contract: _CosmicGame.contract, event: "TimeIncreaseChanged", logs: logs, sub: sub}, nil
}

// WatchTimeIncreaseChanged is a free log subscription operation binding the contract event 0xed46e73b174ced51fb529cdc6c1b4d8abf49387e6d849b71648afb63c81d12cd.
//
// Solidity: event TimeIncreaseChanged(uint256 newTimeIncrease)
func (_CosmicGame *CosmicGameFilterer) WatchTimeIncreaseChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameTimeIncreaseChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "TimeIncreaseChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameTimeIncreaseChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "TimeIncreaseChanged", log); err != nil {
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

// ParseTimeIncreaseChanged is a log parse operation binding the contract event 0xed46e73b174ced51fb529cdc6c1b4d8abf49387e6d849b71648afb63c81d12cd.
//
// Solidity: event TimeIncreaseChanged(uint256 newTimeIncrease)
func (_CosmicGame *CosmicGameFilterer) ParseTimeIncreaseChanged(log types.Log) (*CosmicGameTimeIncreaseChanged, error) {
	event := new(CosmicGameTimeIncreaseChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "TimeIncreaseChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CountersMetaData contains all meta data concerning the Counters contract.
var CountersMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202b30207ad4c731f97f8f5c63a174f7a19eb2f35182448202f31f4ff8d52b1a2f64736f6c63430008130033",
}

// CountersABI is the input ABI used to generate the binding from.
// Deprecated: Use CountersMetaData.ABI instead.
var CountersABI = CountersMetaData.ABI

// CountersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CountersMetaData.Bin instead.
var CountersBin = CountersMetaData.Bin

// DeployCounters deploys a new Ethereum contract, binding an instance of Counters to it.
func DeployCounters(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counters, error) {
	parsed, err := CountersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CountersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counters{CountersCaller: CountersCaller{contract: contract}, CountersTransactor: CountersTransactor{contract: contract}, CountersFilterer: CountersFilterer{contract: contract}}, nil
}

// Counters is an auto generated Go binding around an Ethereum contract.
type Counters struct {
	CountersCaller     // Read-only binding to the contract
	CountersTransactor // Write-only binding to the contract
	CountersFilterer   // Log filterer for contract events
}

// CountersCaller is an auto generated read-only Go binding around an Ethereum contract.
type CountersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CountersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CountersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CountersSession struct {
	Contract     *Counters         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CountersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CountersCallerSession struct {
	Contract *CountersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CountersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CountersTransactorSession struct {
	Contract     *CountersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CountersRaw is an auto generated low-level Go binding around an Ethereum contract.
type CountersRaw struct {
	Contract *Counters // Generic contract binding to access the raw methods on
}

// CountersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CountersCallerRaw struct {
	Contract *CountersCaller // Generic read-only contract binding to access the raw methods on
}

// CountersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CountersTransactorRaw struct {
	Contract *CountersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounters creates a new instance of Counters, bound to a specific deployed contract.
func NewCounters(address common.Address, backend bind.ContractBackend) (*Counters, error) {
	contract, err := bindCounters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counters{CountersCaller: CountersCaller{contract: contract}, CountersTransactor: CountersTransactor{contract: contract}, CountersFilterer: CountersFilterer{contract: contract}}, nil
}

// NewCountersCaller creates a new read-only instance of Counters, bound to a specific deployed contract.
func NewCountersCaller(address common.Address, caller bind.ContractCaller) (*CountersCaller, error) {
	contract, err := bindCounters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CountersCaller{contract: contract}, nil
}

// NewCountersTransactor creates a new write-only instance of Counters, bound to a specific deployed contract.
func NewCountersTransactor(address common.Address, transactor bind.ContractTransactor) (*CountersTransactor, error) {
	contract, err := bindCounters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CountersTransactor{contract: contract}, nil
}

// NewCountersFilterer creates a new log filterer instance of Counters, bound to a specific deployed contract.
func NewCountersFilterer(address common.Address, filterer bind.ContractFilterer) (*CountersFilterer, error) {
	contract, err := bindCounters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CountersFilterer{contract: contract}, nil
}

// bindCounters binds a generic wrapper to an already deployed contract.
func bindCounters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CountersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counters *CountersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counters.Contract.CountersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counters *CountersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counters.Contract.CountersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counters *CountersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counters.Contract.CountersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counters *CountersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counters *CountersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counters *CountersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counters.Contract.contract.Transact(opts, method, params...)
}

// ECDSAMetaData contains all meta data concerning the ECDSA contract.
var ECDSAMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220fddca1388b4cfc79615214a0c4649257a2326f1b0fbd4f2a3cfbd72f6a47c3ca64736f6c63430008130033",
}

// ECDSAABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSAMetaData.ABI instead.
var ECDSAABI = ECDSAMetaData.ABI

// ECDSABin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ECDSAMetaData.Bin instead.
var ECDSABin = ECDSAMetaData.Bin

// DeployECDSA deploys a new Ethereum contract, binding an instance of ECDSA to it.
func DeployECDSA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSA, error) {
	parsed, err := ECDSAMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ECDSABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// ECDSA is an auto generated Go binding around an Ethereum contract.
type ECDSA struct {
	ECDSACaller     // Read-only binding to the contract
	ECDSATransactor // Write-only binding to the contract
	ECDSAFilterer   // Log filterer for contract events
}

// ECDSACaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSASession struct {
	Contract     *ECDSA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSACallerSession struct {
	Contract *ECDSACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ECDSATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSATransactorSession struct {
	Contract     *ECDSATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSARaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSARaw struct {
	Contract *ECDSA // Generic contract binding to access the raw methods on
}

// ECDSACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSACallerRaw struct {
	Contract *ECDSACaller // Generic read-only contract binding to access the raw methods on
}

// ECDSATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSATransactorRaw struct {
	Contract *ECDSATransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSA creates a new instance of ECDSA, bound to a specific deployed contract.
func NewECDSA(address common.Address, backend bind.ContractBackend) (*ECDSA, error) {
	contract, err := bindECDSA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// NewECDSACaller creates a new read-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSACaller(address common.Address, caller bind.ContractCaller) (*ECDSACaller, error) {
	contract, err := bindECDSA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACaller{contract: contract}, nil
}

// NewECDSATransactor creates a new write-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSATransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSATransactor, error) {
	contract, err := bindECDSA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATransactor{contract: contract}, nil
}

// NewECDSAFilterer creates a new log filterer instance of ECDSA, bound to a specific deployed contract.
func NewECDSAFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAFilterer, error) {
	contract, err := bindECDSA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAFilterer{contract: contract}, nil
}

// bindECDSA binds a generic wrapper to an already deployed contract.
func bindECDSA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.ECDSACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transact(opts, method, params...)
}

// EIP712MetaData contains all meta data concerning the EIP712 contract.
var EIP712MetaData = &bind.MetaData{
	ABI: "[]",
}

// EIP712ABI is the input ABI used to generate the binding from.
// Deprecated: Use EIP712MetaData.ABI instead.
var EIP712ABI = EIP712MetaData.ABI

// EIP712 is an auto generated Go binding around an Ethereum contract.
type EIP712 struct {
	EIP712Caller     // Read-only binding to the contract
	EIP712Transactor // Write-only binding to the contract
	EIP712Filterer   // Log filterer for contract events
}

// EIP712Caller is an auto generated read-only Go binding around an Ethereum contract.
type EIP712Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP712Transactor is an auto generated write-only Go binding around an Ethereum contract.
type EIP712Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP712Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EIP712Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP712Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EIP712Session struct {
	Contract     *EIP712           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EIP712CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EIP712CallerSession struct {
	Contract *EIP712Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EIP712TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EIP712TransactorSession struct {
	Contract     *EIP712Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EIP712Raw is an auto generated low-level Go binding around an Ethereum contract.
type EIP712Raw struct {
	Contract *EIP712 // Generic contract binding to access the raw methods on
}

// EIP712CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EIP712CallerRaw struct {
	Contract *EIP712Caller // Generic read-only contract binding to access the raw methods on
}

// EIP712TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EIP712TransactorRaw struct {
	Contract *EIP712Transactor // Generic write-only contract binding to access the raw methods on
}

// NewEIP712 creates a new instance of EIP712, bound to a specific deployed contract.
func NewEIP712(address common.Address, backend bind.ContractBackend) (*EIP712, error) {
	contract, err := bindEIP712(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EIP712{EIP712Caller: EIP712Caller{contract: contract}, EIP712Transactor: EIP712Transactor{contract: contract}, EIP712Filterer: EIP712Filterer{contract: contract}}, nil
}

// NewEIP712Caller creates a new read-only instance of EIP712, bound to a specific deployed contract.
func NewEIP712Caller(address common.Address, caller bind.ContractCaller) (*EIP712Caller, error) {
	contract, err := bindEIP712(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EIP712Caller{contract: contract}, nil
}

// NewEIP712Transactor creates a new write-only instance of EIP712, bound to a specific deployed contract.
func NewEIP712Transactor(address common.Address, transactor bind.ContractTransactor) (*EIP712Transactor, error) {
	contract, err := bindEIP712(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EIP712Transactor{contract: contract}, nil
}

// NewEIP712Filterer creates a new log filterer instance of EIP712, bound to a specific deployed contract.
func NewEIP712Filterer(address common.Address, filterer bind.ContractFilterer) (*EIP712Filterer, error) {
	contract, err := bindEIP712(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EIP712Filterer{contract: contract}, nil
}

// bindEIP712 binds a generic wrapper to an already deployed contract.
func bindEIP712(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EIP712ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP712 *EIP712Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP712.Contract.EIP712Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP712 *EIP712Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP712.Contract.EIP712Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP712 *EIP712Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP712.Contract.EIP712Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP712 *EIP712CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP712.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP712 *EIP712TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP712.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP712 *EIP712TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP712.Contract.contract.Transact(opts, method, params...)
}

// ERC165MetaData contains all meta data concerning the ERC165 contract.
var ERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC165MetaData.ABI instead.
var ERC165ABI = ERC165MetaData.ABI

// ERC165 is an auto generated Go binding around an Ethereum contract.
type ERC165 struct {
	ERC165Caller     // Read-only binding to the contract
	ERC165Transactor // Write-only binding to the contract
	ERC165Filterer   // Log filterer for contract events
}

// ERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165Session struct {
	Contract     *ERC165           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165CallerSession struct {
	Contract *ERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165TransactorSession struct {
	Contract     *ERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165Raw struct {
	Contract *ERC165 // Generic contract binding to access the raw methods on
}

// ERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165CallerRaw struct {
	Contract *ERC165Caller // Generic read-only contract binding to access the raw methods on
}

// ERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165TransactorRaw struct {
	Contract *ERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165 creates a new instance of ERC165, bound to a specific deployed contract.
func NewERC165(address common.Address, backend bind.ContractBackend) (*ERC165, error) {
	contract, err := bindERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// NewERC165Caller creates a new read-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Caller(address common.Address, caller bind.ContractCaller) (*ERC165Caller, error) {
	contract, err := bindERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Caller{contract: contract}, nil
}

// NewERC165Transactor creates a new write-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC165Transactor, error) {
	contract, err := bindERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Transactor{contract: contract}, nil
}

// NewERC165Filterer creates a new log filterer instance of ERC165, bound to a specific deployed contract.
func NewERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC165Filterer, error) {
	contract, err := bindERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165Filterer{contract: contract}, nil
}

// bindERC165 binds a generic wrapper to an already deployed contract.
func bindERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.ERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// IERC165MetaData contains all meta data concerning the IERC165 contract.
var IERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC165MetaData.ABI instead.
var IERC165ABI = IERC165MetaData.ABI

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetadataMetaData contains all meta data concerning the IERC20Metadata contract.
var IERC20MetadataMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20MetadataABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetadataMetaData.ABI instead.
var IERC20MetadataABI = IERC20MetadataMetaData.ABI

// IERC20Metadata is an auto generated Go binding around an Ethereum contract.
type IERC20Metadata struct {
	IERC20MetadataCaller     // Read-only binding to the contract
	IERC20MetadataTransactor // Write-only binding to the contract
	IERC20MetadataFilterer   // Log filterer for contract events
}

// IERC20MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20MetadataSession struct {
	Contract     *IERC20Metadata   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20MetadataCallerSession struct {
	Contract *IERC20MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IERC20MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20MetadataTransactorSession struct {
	Contract     *IERC20MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IERC20MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20MetadataRaw struct {
	Contract *IERC20Metadata // Generic contract binding to access the raw methods on
}

// IERC20MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20MetadataCallerRaw struct {
	Contract *IERC20MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20MetadataTransactorRaw struct {
	Contract *IERC20MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Metadata creates a new instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20Metadata(address common.Address, backend bind.ContractBackend) (*IERC20Metadata, error) {
	contract, err := bindIERC20Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Metadata{IERC20MetadataCaller: IERC20MetadataCaller{contract: contract}, IERC20MetadataTransactor: IERC20MetadataTransactor{contract: contract}, IERC20MetadataFilterer: IERC20MetadataFilterer{contract: contract}}, nil
}

// NewIERC20MetadataCaller creates a new read-only instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataCaller(address common.Address, caller bind.ContractCaller) (*IERC20MetadataCaller, error) {
	contract, err := bindIERC20Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataCaller{contract: contract}, nil
}

// NewIERC20MetadataTransactor creates a new write-only instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20MetadataTransactor, error) {
	contract, err := bindIERC20Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataTransactor{contract: contract}, nil
}

// NewIERC20MetadataFilterer creates a new log filterer instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20MetadataFilterer, error) {
	contract, err := bindIERC20Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataFilterer{contract: contract}, nil
}

// bindIERC20Metadata binds a generic wrapper to an already deployed contract.
func bindIERC20Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20MetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Metadata *IERC20MetadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Metadata.Contract.IERC20MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Metadata *IERC20MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.IERC20MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Metadata *IERC20MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.IERC20MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Metadata *IERC20MetadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Metadata *IERC20MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Metadata *IERC20MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.Allowance(&_IERC20Metadata.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.Allowance(&_IERC20Metadata.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.BalanceOf(&_IERC20Metadata.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.BalanceOf(&_IERC20Metadata.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataSession) Decimals() (uint8, error) {
	return _IERC20Metadata.Contract.Decimals(&_IERC20Metadata.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataCallerSession) Decimals() (uint8, error) {
	return _IERC20Metadata.Contract.Decimals(&_IERC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataSession) Name() (string, error) {
	return _IERC20Metadata.Contract.Name(&_IERC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataCallerSession) Name() (string, error) {
	return _IERC20Metadata.Contract.Name(&_IERC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataSession) Symbol() (string, error) {
	return _IERC20Metadata.Contract.Symbol(&_IERC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataCallerSession) Symbol() (string, error) {
	return _IERC20Metadata.Contract.Symbol(&_IERC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) TotalSupply() (*big.Int, error) {
	return _IERC20Metadata.Contract.TotalSupply(&_IERC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20Metadata.Contract.TotalSupply(&_IERC20Metadata.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Approve(&_IERC20Metadata.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Approve(&_IERC20Metadata.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Transfer(&_IERC20Metadata.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Transfer(&_IERC20Metadata.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.TransferFrom(&_IERC20Metadata.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.TransferFrom(&_IERC20Metadata.TransactOpts, sender, recipient, amount)
}

// IERC20MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20Metadata contract.
type IERC20MetadataApprovalIterator struct {
	Event *IERC20MetadataApproval // Event containing the contract specifics and raw log

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
func (it *IERC20MetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MetadataApproval)
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
		it.Event = new(IERC20MetadataApproval)
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
func (it *IERC20MetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MetadataApproval represents a Approval event raised by the IERC20Metadata contract.
type IERC20MetadataApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20MetadataApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Metadata.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataApprovalIterator{contract: _IERC20Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20MetadataApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Metadata.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MetadataApproval)
				if err := _IERC20Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20Metadata *IERC20MetadataFilterer) ParseApproval(log types.Log) (*IERC20MetadataApproval, error) {
	event := new(IERC20MetadataApproval)
	if err := _IERC20Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20Metadata contract.
type IERC20MetadataTransferIterator struct {
	Event *IERC20MetadataTransfer // Event containing the contract specifics and raw log

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
func (it *IERC20MetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MetadataTransfer)
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
		it.Event = new(IERC20MetadataTransfer)
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
func (it *IERC20MetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MetadataTransfer represents a Transfer event raised by the IERC20Metadata contract.
type IERC20MetadataTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20MetadataTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Metadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataTransferIterator{contract: _IERC20Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20MetadataTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Metadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MetadataTransfer)
				if err := _IERC20Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20Metadata *IERC20MetadataFilterer) ParseTransfer(log types.Log) (*IERC20MetadataTransfer, error) {
	event := new(IERC20MetadataTransfer)
	if err := _IERC20Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20PermitMetaData contains all meta data concerning the IERC20Permit contract.
var IERC20PermitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20PermitABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20PermitMetaData.ABI instead.
var IERC20PermitABI = IERC20PermitMetaData.ABI

// IERC20Permit is an auto generated Go binding around an Ethereum contract.
type IERC20Permit struct {
	IERC20PermitCaller     // Read-only binding to the contract
	IERC20PermitTransactor // Write-only binding to the contract
	IERC20PermitFilterer   // Log filterer for contract events
}

// IERC20PermitCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20PermitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20PermitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20PermitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20PermitSession struct {
	Contract     *IERC20Permit     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20PermitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20PermitCallerSession struct {
	Contract *IERC20PermitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IERC20PermitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20PermitTransactorSession struct {
	Contract     *IERC20PermitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IERC20PermitRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20PermitRaw struct {
	Contract *IERC20Permit // Generic contract binding to access the raw methods on
}

// IERC20PermitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20PermitCallerRaw struct {
	Contract *IERC20PermitCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20PermitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20PermitTransactorRaw struct {
	Contract *IERC20PermitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Permit creates a new instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20Permit(address common.Address, backend bind.ContractBackend) (*IERC20Permit, error) {
	contract, err := bindIERC20Permit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Permit{IERC20PermitCaller: IERC20PermitCaller{contract: contract}, IERC20PermitTransactor: IERC20PermitTransactor{contract: contract}, IERC20PermitFilterer: IERC20PermitFilterer{contract: contract}}, nil
}

// NewIERC20PermitCaller creates a new read-only instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitCaller(address common.Address, caller bind.ContractCaller) (*IERC20PermitCaller, error) {
	contract, err := bindIERC20Permit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitCaller{contract: contract}, nil
}

// NewIERC20PermitTransactor creates a new write-only instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20PermitTransactor, error) {
	contract, err := bindIERC20Permit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitTransactor{contract: contract}, nil
}

// NewIERC20PermitFilterer creates a new log filterer instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20PermitFilterer, error) {
	contract, err := bindIERC20Permit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitFilterer{contract: contract}, nil
}

// bindIERC20Permit binds a generic wrapper to an already deployed contract.
func bindIERC20Permit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20PermitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Permit *IERC20PermitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Permit.Contract.IERC20PermitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Permit *IERC20PermitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Permit.Contract.IERC20PermitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Permit *IERC20PermitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Permit.Contract.IERC20PermitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Permit *IERC20PermitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Permit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Permit *IERC20PermitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Permit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Permit *IERC20PermitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Permit.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20Permit *IERC20PermitCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IERC20Permit.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20Permit *IERC20PermitSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC20Permit.Contract.DOMAINSEPARATOR(&_IERC20Permit.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20Permit *IERC20PermitCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC20Permit.Contract.DOMAINSEPARATOR(&_IERC20Permit.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20Permit *IERC20PermitCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Permit.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20Permit *IERC20PermitSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IERC20Permit.Contract.Nonces(&_IERC20Permit.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20Permit *IERC20PermitCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IERC20Permit.Contract.Nonces(&_IERC20Permit.CallOpts, owner)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.Contract.Permit(&_IERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.Contract.Permit(&_IERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// IERC3156FlashBorrowerMetaData contains all meta data concerning the IERC3156FlashBorrower contract.
var IERC3156FlashBorrowerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onFlashLoan\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC3156FlashBorrowerABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC3156FlashBorrowerMetaData.ABI instead.
var IERC3156FlashBorrowerABI = IERC3156FlashBorrowerMetaData.ABI

// IERC3156FlashBorrower is an auto generated Go binding around an Ethereum contract.
type IERC3156FlashBorrower struct {
	IERC3156FlashBorrowerCaller     // Read-only binding to the contract
	IERC3156FlashBorrowerTransactor // Write-only binding to the contract
	IERC3156FlashBorrowerFilterer   // Log filterer for contract events
}

// IERC3156FlashBorrowerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC3156FlashBorrowerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC3156FlashBorrowerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC3156FlashBorrowerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC3156FlashBorrowerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC3156FlashBorrowerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC3156FlashBorrowerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC3156FlashBorrowerSession struct {
	Contract     *IERC3156FlashBorrower // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IERC3156FlashBorrowerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC3156FlashBorrowerCallerSession struct {
	Contract *IERC3156FlashBorrowerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IERC3156FlashBorrowerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC3156FlashBorrowerTransactorSession struct {
	Contract     *IERC3156FlashBorrowerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IERC3156FlashBorrowerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC3156FlashBorrowerRaw struct {
	Contract *IERC3156FlashBorrower // Generic contract binding to access the raw methods on
}

// IERC3156FlashBorrowerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC3156FlashBorrowerCallerRaw struct {
	Contract *IERC3156FlashBorrowerCaller // Generic read-only contract binding to access the raw methods on
}

// IERC3156FlashBorrowerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC3156FlashBorrowerTransactorRaw struct {
	Contract *IERC3156FlashBorrowerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC3156FlashBorrower creates a new instance of IERC3156FlashBorrower, bound to a specific deployed contract.
func NewIERC3156FlashBorrower(address common.Address, backend bind.ContractBackend) (*IERC3156FlashBorrower, error) {
	contract, err := bindIERC3156FlashBorrower(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashBorrower{IERC3156FlashBorrowerCaller: IERC3156FlashBorrowerCaller{contract: contract}, IERC3156FlashBorrowerTransactor: IERC3156FlashBorrowerTransactor{contract: contract}, IERC3156FlashBorrowerFilterer: IERC3156FlashBorrowerFilterer{contract: contract}}, nil
}

// NewIERC3156FlashBorrowerCaller creates a new read-only instance of IERC3156FlashBorrower, bound to a specific deployed contract.
func NewIERC3156FlashBorrowerCaller(address common.Address, caller bind.ContractCaller) (*IERC3156FlashBorrowerCaller, error) {
	contract, err := bindIERC3156FlashBorrower(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashBorrowerCaller{contract: contract}, nil
}

// NewIERC3156FlashBorrowerTransactor creates a new write-only instance of IERC3156FlashBorrower, bound to a specific deployed contract.
func NewIERC3156FlashBorrowerTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC3156FlashBorrowerTransactor, error) {
	contract, err := bindIERC3156FlashBorrower(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashBorrowerTransactor{contract: contract}, nil
}

// NewIERC3156FlashBorrowerFilterer creates a new log filterer instance of IERC3156FlashBorrower, bound to a specific deployed contract.
func NewIERC3156FlashBorrowerFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC3156FlashBorrowerFilterer, error) {
	contract, err := bindIERC3156FlashBorrower(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashBorrowerFilterer{contract: contract}, nil
}

// bindIERC3156FlashBorrower binds a generic wrapper to an already deployed contract.
func bindIERC3156FlashBorrower(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC3156FlashBorrowerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC3156FlashBorrower.Contract.IERC3156FlashBorrowerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.Contract.IERC3156FlashBorrowerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.Contract.IERC3156FlashBorrowerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC3156FlashBorrower.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.Contract.contract.Transact(opts, method, params...)
}

// OnFlashLoan is a paid mutator transaction binding the contract method 0x23e30c8b.
//
// Solidity: function onFlashLoan(address initiator, address token, uint256 amount, uint256 fee, bytes data) returns(bytes32)
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerTransactor) OnFlashLoan(opts *bind.TransactOpts, initiator common.Address, token common.Address, amount *big.Int, fee *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.contract.Transact(opts, "onFlashLoan", initiator, token, amount, fee, data)
}

// OnFlashLoan is a paid mutator transaction binding the contract method 0x23e30c8b.
//
// Solidity: function onFlashLoan(address initiator, address token, uint256 amount, uint256 fee, bytes data) returns(bytes32)
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerSession) OnFlashLoan(initiator common.Address, token common.Address, amount *big.Int, fee *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.Contract.OnFlashLoan(&_IERC3156FlashBorrower.TransactOpts, initiator, token, amount, fee, data)
}

// OnFlashLoan is a paid mutator transaction binding the contract method 0x23e30c8b.
//
// Solidity: function onFlashLoan(address initiator, address token, uint256 amount, uint256 fee, bytes data) returns(bytes32)
func (_IERC3156FlashBorrower *IERC3156FlashBorrowerTransactorSession) OnFlashLoan(initiator common.Address, token common.Address, amount *big.Int, fee *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC3156FlashBorrower.Contract.OnFlashLoan(&_IERC3156FlashBorrower.TransactOpts, initiator, token, amount, fee, data)
}

// IERC3156FlashLenderMetaData contains all meta data concerning the IERC3156FlashLender contract.
var IERC3156FlashLenderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"flashFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC3156FlashBorrower\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"maxFlashLoan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IERC3156FlashLenderABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC3156FlashLenderMetaData.ABI instead.
var IERC3156FlashLenderABI = IERC3156FlashLenderMetaData.ABI

// IERC3156FlashLender is an auto generated Go binding around an Ethereum contract.
type IERC3156FlashLender struct {
	IERC3156FlashLenderCaller     // Read-only binding to the contract
	IERC3156FlashLenderTransactor // Write-only binding to the contract
	IERC3156FlashLenderFilterer   // Log filterer for contract events
}

// IERC3156FlashLenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC3156FlashLenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC3156FlashLenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC3156FlashLenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC3156FlashLenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC3156FlashLenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC3156FlashLenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC3156FlashLenderSession struct {
	Contract     *IERC3156FlashLender // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IERC3156FlashLenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC3156FlashLenderCallerSession struct {
	Contract *IERC3156FlashLenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IERC3156FlashLenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC3156FlashLenderTransactorSession struct {
	Contract     *IERC3156FlashLenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IERC3156FlashLenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC3156FlashLenderRaw struct {
	Contract *IERC3156FlashLender // Generic contract binding to access the raw methods on
}

// IERC3156FlashLenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC3156FlashLenderCallerRaw struct {
	Contract *IERC3156FlashLenderCaller // Generic read-only contract binding to access the raw methods on
}

// IERC3156FlashLenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC3156FlashLenderTransactorRaw struct {
	Contract *IERC3156FlashLenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC3156FlashLender creates a new instance of IERC3156FlashLender, bound to a specific deployed contract.
func NewIERC3156FlashLender(address common.Address, backend bind.ContractBackend) (*IERC3156FlashLender, error) {
	contract, err := bindIERC3156FlashLender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashLender{IERC3156FlashLenderCaller: IERC3156FlashLenderCaller{contract: contract}, IERC3156FlashLenderTransactor: IERC3156FlashLenderTransactor{contract: contract}, IERC3156FlashLenderFilterer: IERC3156FlashLenderFilterer{contract: contract}}, nil
}

// NewIERC3156FlashLenderCaller creates a new read-only instance of IERC3156FlashLender, bound to a specific deployed contract.
func NewIERC3156FlashLenderCaller(address common.Address, caller bind.ContractCaller) (*IERC3156FlashLenderCaller, error) {
	contract, err := bindIERC3156FlashLender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashLenderCaller{contract: contract}, nil
}

// NewIERC3156FlashLenderTransactor creates a new write-only instance of IERC3156FlashLender, bound to a specific deployed contract.
func NewIERC3156FlashLenderTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC3156FlashLenderTransactor, error) {
	contract, err := bindIERC3156FlashLender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashLenderTransactor{contract: contract}, nil
}

// NewIERC3156FlashLenderFilterer creates a new log filterer instance of IERC3156FlashLender, bound to a specific deployed contract.
func NewIERC3156FlashLenderFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC3156FlashLenderFilterer, error) {
	contract, err := bindIERC3156FlashLender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC3156FlashLenderFilterer{contract: contract}, nil
}

// bindIERC3156FlashLender binds a generic wrapper to an already deployed contract.
func bindIERC3156FlashLender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC3156FlashLenderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC3156FlashLender *IERC3156FlashLenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC3156FlashLender.Contract.IERC3156FlashLenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC3156FlashLender *IERC3156FlashLenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC3156FlashLender.Contract.IERC3156FlashLenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC3156FlashLender *IERC3156FlashLenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC3156FlashLender.Contract.IERC3156FlashLenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC3156FlashLender *IERC3156FlashLenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC3156FlashLender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC3156FlashLender *IERC3156FlashLenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC3156FlashLender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC3156FlashLender *IERC3156FlashLenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC3156FlashLender.Contract.contract.Transact(opts, method, params...)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_IERC3156FlashLender *IERC3156FlashLenderCaller) FlashFee(opts *bind.CallOpts, token common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC3156FlashLender.contract.Call(opts, &out, "flashFee", token, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_IERC3156FlashLender *IERC3156FlashLenderSession) FlashFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _IERC3156FlashLender.Contract.FlashFee(&_IERC3156FlashLender.CallOpts, token, amount)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_IERC3156FlashLender *IERC3156FlashLenderCallerSession) FlashFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _IERC3156FlashLender.Contract.FlashFee(&_IERC3156FlashLender.CallOpts, token, amount)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_IERC3156FlashLender *IERC3156FlashLenderCaller) MaxFlashLoan(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC3156FlashLender.contract.Call(opts, &out, "maxFlashLoan", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_IERC3156FlashLender *IERC3156FlashLenderSession) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _IERC3156FlashLender.Contract.MaxFlashLoan(&_IERC3156FlashLender.CallOpts, token)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_IERC3156FlashLender *IERC3156FlashLenderCallerSession) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _IERC3156FlashLender.Contract.MaxFlashLoan(&_IERC3156FlashLender.CallOpts, token)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_IERC3156FlashLender *IERC3156FlashLenderTransactor) FlashLoan(opts *bind.TransactOpts, receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC3156FlashLender.contract.Transact(opts, "flashLoan", receiver, token, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_IERC3156FlashLender *IERC3156FlashLenderSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC3156FlashLender.Contract.FlashLoan(&_IERC3156FlashLender.TransactOpts, receiver, token, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_IERC3156FlashLender *IERC3156FlashLenderTransactorSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC3156FlashLender.Contract.FlashLoan(&_IERC3156FlashLender.TransactOpts, receiver, token, amount, data)
}

// IERC721MetaData contains all meta data concerning the IERC721 contract.
var IERC721MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC721MetaData.ABI instead.
var IERC721ABI = IERC721MetaData.ABI

// IERC721 is an auto generated Go binding around an Ethereum contract.
type IERC721 struct {
	IERC721Caller     // Read-only binding to the contract
	IERC721Transactor // Write-only binding to the contract
	IERC721Filterer   // Log filterer for contract events
}

// IERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721Session struct {
	Contract     *IERC721          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721CallerSession struct {
	Contract *IERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721TransactorSession struct {
	Contract     *IERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721Raw struct {
	Contract *IERC721 // Generic contract binding to access the raw methods on
}

// IERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721CallerRaw struct {
	Contract *IERC721Caller // Generic read-only contract binding to access the raw methods on
}

// IERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721TransactorRaw struct {
	Contract *IERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721 creates a new instance of IERC721, bound to a specific deployed contract.
func NewIERC721(address common.Address, backend bind.ContractBackend) (*IERC721, error) {
	contract, err := bindIERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721{IERC721Caller: IERC721Caller{contract: contract}, IERC721Transactor: IERC721Transactor{contract: contract}, IERC721Filterer: IERC721Filterer{contract: contract}}, nil
}

// NewIERC721Caller creates a new read-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Caller(address common.Address, caller bind.ContractCaller) (*IERC721Caller, error) {
	contract, err := bindIERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Caller{contract: contract}, nil
}

// NewIERC721Transactor creates a new write-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC721Transactor, error) {
	contract, err := bindIERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Transactor{contract: contract}, nil
}

// NewIERC721Filterer creates a new log filterer instance of IERC721, bound to a specific deployed contract.
func NewIERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC721Filterer, error) {
	contract, err := bindIERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721Filterer{contract: contract}, nil
}

// bindIERC721 binds a generic wrapper to an already deployed contract.
func bindIERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.IERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721 *IERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721 *IERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721 *IERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721 *IERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721 *IERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721 *IERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721 *IERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721 *IERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721 *IERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721 *IERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721 *IERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721 *IERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721 *IERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721 *IERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721 *IERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Session) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721TransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// IERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721 contract.
type IERC721ApprovalIterator struct {
	Event *IERC721Approval // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Approval)
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
		it.Event = new(IERC721Approval)
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
func (it *IERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Approval represents a Approval event raised by the IERC721 contract.
type IERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721ApprovalIterator, error) {

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

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalIterator{contract: _IERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Approval)
				if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC721 *IERC721Filterer) ParseApproval(log types.Log) (*IERC721Approval, error) {
	event := new(IERC721Approval)
	if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721 contract.
type IERC721ApprovalForAllIterator struct {
	Event *IERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721ApprovalForAll)
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
		it.Event = new(IERC721ApprovalForAll)
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
func (it *IERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721ApprovalForAll represents a ApprovalForAll event raised by the IERC721 contract.
type IERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalForAllIterator{contract: _IERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721ApprovalForAll)
				if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_IERC721 *IERC721Filterer) ParseApprovalForAll(log types.Log) (*IERC721ApprovalForAll, error) {
	event := new(IERC721ApprovalForAll)
	if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721 contract.
type IERC721TransferIterator struct {
	Event *IERC721Transfer // Event containing the contract specifics and raw log

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
func (it *IERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Transfer)
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
		it.Event = new(IERC721Transfer)
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
func (it *IERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Transfer represents a Transfer event raised by the IERC721 contract.
type IERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721TransferIterator, error) {

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

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721TransferIterator{contract: _IERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Transfer)
				if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC721 *IERC721Filterer) ParseTransfer(log types.Log) (*IERC721Transfer, error) {
	event := new(IERC721Transfer)
	if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721EnumerableMetaData contains all meta data concerning the IERC721Enumerable contract.
var IERC721EnumerableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC721EnumerableABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC721EnumerableMetaData.ABI instead.
var IERC721EnumerableABI = IERC721EnumerableMetaData.ABI

// IERC721Enumerable is an auto generated Go binding around an Ethereum contract.
type IERC721Enumerable struct {
	IERC721EnumerableCaller     // Read-only binding to the contract
	IERC721EnumerableTransactor // Write-only binding to the contract
	IERC721EnumerableFilterer   // Log filterer for contract events
}

// IERC721EnumerableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721EnumerableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721EnumerableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721EnumerableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721EnumerableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721EnumerableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721EnumerableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721EnumerableSession struct {
	Contract     *IERC721Enumerable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC721EnumerableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721EnumerableCallerSession struct {
	Contract *IERC721EnumerableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IERC721EnumerableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721EnumerableTransactorSession struct {
	Contract     *IERC721EnumerableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IERC721EnumerableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721EnumerableRaw struct {
	Contract *IERC721Enumerable // Generic contract binding to access the raw methods on
}

// IERC721EnumerableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721EnumerableCallerRaw struct {
	Contract *IERC721EnumerableCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721EnumerableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721EnumerableTransactorRaw struct {
	Contract *IERC721EnumerableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Enumerable creates a new instance of IERC721Enumerable, bound to a specific deployed contract.
func NewIERC721Enumerable(address common.Address, backend bind.ContractBackend) (*IERC721Enumerable, error) {
	contract, err := bindIERC721Enumerable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Enumerable{IERC721EnumerableCaller: IERC721EnumerableCaller{contract: contract}, IERC721EnumerableTransactor: IERC721EnumerableTransactor{contract: contract}, IERC721EnumerableFilterer: IERC721EnumerableFilterer{contract: contract}}, nil
}

// NewIERC721EnumerableCaller creates a new read-only instance of IERC721Enumerable, bound to a specific deployed contract.
func NewIERC721EnumerableCaller(address common.Address, caller bind.ContractCaller) (*IERC721EnumerableCaller, error) {
	contract, err := bindIERC721Enumerable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableCaller{contract: contract}, nil
}

// NewIERC721EnumerableTransactor creates a new write-only instance of IERC721Enumerable, bound to a specific deployed contract.
func NewIERC721EnumerableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721EnumerableTransactor, error) {
	contract, err := bindIERC721Enumerable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableTransactor{contract: contract}, nil
}

// NewIERC721EnumerableFilterer creates a new log filterer instance of IERC721Enumerable, bound to a specific deployed contract.
func NewIERC721EnumerableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721EnumerableFilterer, error) {
	contract, err := bindIERC721Enumerable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableFilterer{contract: contract}, nil
}

// bindIERC721Enumerable binds a generic wrapper to an already deployed contract.
func bindIERC721Enumerable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721EnumerableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Enumerable *IERC721EnumerableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Enumerable.Contract.IERC721EnumerableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Enumerable *IERC721EnumerableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.IERC721EnumerableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Enumerable *IERC721EnumerableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.IERC721EnumerableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Enumerable *IERC721EnumerableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Enumerable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Enumerable *IERC721EnumerableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Enumerable *IERC721EnumerableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Enumerable *IERC721EnumerableCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Enumerable *IERC721EnumerableSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Enumerable.Contract.BalanceOf(&_IERC721Enumerable.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Enumerable.Contract.BalanceOf(&_IERC721Enumerable.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Enumerable *IERC721EnumerableCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Enumerable *IERC721EnumerableSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Enumerable.Contract.GetApproved(&_IERC721Enumerable.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Enumerable.Contract.GetApproved(&_IERC721Enumerable.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Enumerable *IERC721EnumerableCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Enumerable *IERC721EnumerableSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Enumerable.Contract.IsApprovedForAll(&_IERC721Enumerable.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Enumerable.Contract.IsApprovedForAll(&_IERC721Enumerable.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Enumerable *IERC721EnumerableCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Enumerable *IERC721EnumerableSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Enumerable.Contract.OwnerOf(&_IERC721Enumerable.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Enumerable.Contract.OwnerOf(&_IERC721Enumerable.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Enumerable *IERC721EnumerableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Enumerable *IERC721EnumerableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Enumerable.Contract.SupportsInterface(&_IERC721Enumerable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Enumerable.Contract.SupportsInterface(&_IERC721Enumerable.CallOpts, interfaceId)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _IERC721Enumerable.Contract.TokenByIndex(&_IERC721Enumerable.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _IERC721Enumerable.Contract.TokenByIndex(&_IERC721Enumerable.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_IERC721Enumerable *IERC721EnumerableCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_IERC721Enumerable *IERC721EnumerableSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _IERC721Enumerable.Contract.TokenOfOwnerByIndex(&_IERC721Enumerable.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _IERC721Enumerable.Contract.TokenOfOwnerByIndex(&_IERC721Enumerable.CallOpts, owner, index)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC721Enumerable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableSession) TotalSupply() (*big.Int, error) {
	return _IERC721Enumerable.Contract.TotalSupply(&_IERC721Enumerable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC721Enumerable.Contract.TotalSupply(&_IERC721Enumerable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.Approve(&_IERC721Enumerable.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.Approve(&_IERC721Enumerable.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.SafeTransferFrom(&_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.SafeTransferFrom(&_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.SafeTransferFrom0(&_IERC721Enumerable.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.SafeTransferFrom0(&_IERC721Enumerable.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.SetApprovalForAll(&_IERC721Enumerable.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.SetApprovalForAll(&_IERC721Enumerable.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.TransferFrom(&_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.TransferFrom(&_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

// IERC721EnumerableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721Enumerable contract.
type IERC721EnumerableApprovalIterator struct {
	Event *IERC721EnumerableApproval // Event containing the contract specifics and raw log

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
func (it *IERC721EnumerableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721EnumerableApproval)
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
		it.Event = new(IERC721EnumerableApproval)
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
func (it *IERC721EnumerableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721EnumerableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721EnumerableApproval represents a Approval event raised by the IERC721Enumerable contract.
type IERC721EnumerableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721EnumerableApprovalIterator, error) {

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

	logs, sub, err := _IERC721Enumerable.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableApprovalIterator{contract: _IERC721Enumerable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721EnumerableApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _IERC721Enumerable.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721EnumerableApproval)
				if err := _IERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC721Enumerable *IERC721EnumerableFilterer) ParseApproval(log types.Log) (*IERC721EnumerableApproval, error) {
	event := new(IERC721EnumerableApproval)
	if err := _IERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721EnumerableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721Enumerable contract.
type IERC721EnumerableApprovalForAllIterator struct {
	Event *IERC721EnumerableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721EnumerableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721EnumerableApprovalForAll)
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
		it.Event = new(IERC721EnumerableApprovalForAll)
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
func (it *IERC721EnumerableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721EnumerableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721EnumerableApprovalForAll represents a ApprovalForAll event raised by the IERC721Enumerable contract.
type IERC721EnumerableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Enumerable *IERC721EnumerableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721EnumerableApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableApprovalForAllIterator{contract: _IERC721Enumerable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Enumerable *IERC721EnumerableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721EnumerableApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721EnumerableApprovalForAll)
				if err := _IERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_IERC721Enumerable *IERC721EnumerableFilterer) ParseApprovalForAll(log types.Log) (*IERC721EnumerableApprovalForAll, error) {
	event := new(IERC721EnumerableApprovalForAll)
	if err := _IERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721EnumerableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721Enumerable contract.
type IERC721EnumerableTransferIterator struct {
	Event *IERC721EnumerableTransfer // Event containing the contract specifics and raw log

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
func (it *IERC721EnumerableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721EnumerableTransfer)
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
		it.Event = new(IERC721EnumerableTransfer)
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
func (it *IERC721EnumerableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721EnumerableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721EnumerableTransfer represents a Transfer event raised by the IERC721Enumerable contract.
type IERC721EnumerableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721EnumerableTransferIterator, error) {

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

	logs, sub, err := _IERC721Enumerable.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableTransferIterator{contract: _IERC721Enumerable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721EnumerableTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _IERC721Enumerable.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721EnumerableTransfer)
				if err := _IERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC721Enumerable *IERC721EnumerableFilterer) ParseTransfer(log types.Log) (*IERC721EnumerableTransfer, error) {
	event := new(IERC721EnumerableTransfer)
	if err := _IERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721MetadataMetaData contains all meta data concerning the IERC721Metadata contract.
var IERC721MetadataMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC721MetadataABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC721MetadataMetaData.ABI instead.
var IERC721MetadataABI = IERC721MetadataMetaData.ABI

// IERC721Metadata is an auto generated Go binding around an Ethereum contract.
type IERC721Metadata struct {
	IERC721MetadataCaller     // Read-only binding to the contract
	IERC721MetadataTransactor // Write-only binding to the contract
	IERC721MetadataFilterer   // Log filterer for contract events
}

// IERC721MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721MetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721MetadataSession struct {
	Contract     *IERC721Metadata  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721MetadataCallerSession struct {
	Contract *IERC721MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IERC721MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721MetadataTransactorSession struct {
	Contract     *IERC721MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IERC721MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721MetadataRaw struct {
	Contract *IERC721Metadata // Generic contract binding to access the raw methods on
}

// IERC721MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721MetadataCallerRaw struct {
	Contract *IERC721MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721MetadataTransactorRaw struct {
	Contract *IERC721MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Metadata creates a new instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721Metadata(address common.Address, backend bind.ContractBackend) (*IERC721Metadata, error) {
	contract, err := bindIERC721Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Metadata{IERC721MetadataCaller: IERC721MetadataCaller{contract: contract}, IERC721MetadataTransactor: IERC721MetadataTransactor{contract: contract}, IERC721MetadataFilterer: IERC721MetadataFilterer{contract: contract}}, nil
}

// NewIERC721MetadataCaller creates a new read-only instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721MetadataCaller(address common.Address, caller bind.ContractCaller) (*IERC721MetadataCaller, error) {
	contract, err := bindIERC721Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataCaller{contract: contract}, nil
}

// NewIERC721MetadataTransactor creates a new write-only instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721MetadataTransactor, error) {
	contract, err := bindIERC721Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataTransactor{contract: contract}, nil
}

// NewIERC721MetadataFilterer creates a new log filterer instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721MetadataFilterer, error) {
	contract, err := bindIERC721Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataFilterer{contract: contract}, nil
}

// bindIERC721Metadata binds a generic wrapper to an already deployed contract.
func bindIERC721Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721MetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Metadata *IERC721MetadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Metadata.Contract.IERC721MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Metadata *IERC721MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.IERC721MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Metadata *IERC721MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.IERC721MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Metadata *IERC721MetadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Metadata *IERC721MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Metadata *IERC721MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Metadata *IERC721MetadataCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Metadata *IERC721MetadataSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Metadata.Contract.BalanceOf(&_IERC721Metadata.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721Metadata *IERC721MetadataCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Metadata.Contract.BalanceOf(&_IERC721Metadata.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Metadata *IERC721MetadataCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Metadata *IERC721MetadataSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.GetApproved(&_IERC721Metadata.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721Metadata *IERC721MetadataCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.GetApproved(&_IERC721Metadata.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Metadata *IERC721MetadataCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Metadata *IERC721MetadataSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Metadata.Contract.IsApprovedForAll(&_IERC721Metadata.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721Metadata *IERC721MetadataCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Metadata.Contract.IsApprovedForAll(&_IERC721Metadata.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC721Metadata *IERC721MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC721Metadata *IERC721MetadataSession) Name() (string, error) {
	return _IERC721Metadata.Contract.Name(&_IERC721Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC721Metadata *IERC721MetadataCallerSession) Name() (string, error) {
	return _IERC721Metadata.Contract.Name(&_IERC721Metadata.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Metadata *IERC721MetadataCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Metadata *IERC721MetadataSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.OwnerOf(&_IERC721Metadata.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721Metadata *IERC721MetadataCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.OwnerOf(&_IERC721Metadata.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Metadata *IERC721MetadataCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Metadata *IERC721MetadataSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Metadata.Contract.SupportsInterface(&_IERC721Metadata.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721Metadata *IERC721MetadataCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Metadata.Contract.SupportsInterface(&_IERC721Metadata.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC721Metadata *IERC721MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC721Metadata *IERC721MetadataSession) Symbol() (string, error) {
	return _IERC721Metadata.Contract.Symbol(&_IERC721Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC721Metadata *IERC721MetadataCallerSession) Symbol() (string, error) {
	return _IERC721Metadata.Contract.Symbol(&_IERC721Metadata.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IERC721Metadata *IERC721MetadataCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _IERC721Metadata.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IERC721Metadata *IERC721MetadataSession) TokenURI(tokenId *big.Int) (string, error) {
	return _IERC721Metadata.Contract.TokenURI(&_IERC721Metadata.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IERC721Metadata *IERC721MetadataCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _IERC721Metadata.Contract.TokenURI(&_IERC721Metadata.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.Approve(&_IERC721Metadata.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.Approve(&_IERC721Metadata.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Metadata *IERC721MetadataSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom0(&_IERC721Metadata.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom0(&_IERC721Metadata.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Metadata *IERC721MetadataSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SetApprovalForAll(&_IERC721Metadata.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.SetApprovalForAll(&_IERC721Metadata.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.TransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.TransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// IERC721MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721Metadata contract.
type IERC721MetadataApprovalIterator struct {
	Event *IERC721MetadataApproval // Event containing the contract specifics and raw log

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
func (it *IERC721MetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721MetadataApproval)
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
		it.Event = new(IERC721MetadataApproval)
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
func (it *IERC721MetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721MetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}
// ERC721Transfer represents a Transfer event raised by the ERC721 contract.
type ERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}
// ERC721MetaData contains all meta data concerning the ERC721 contract.
var ERC721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002946380380620029468339818101604052810190620000379190620001f6565b8160009081620000489190620004c6565b5080600190816200005a9190620004c6565b505050620005ad565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620000cc8262000081565b810181811067ffffffffffffffff82111715620000ee57620000ed62000092565b5b80604052505050565b60006200010362000063565b9050620001118282620000c1565b919050565b600067ffffffffffffffff82111562000134576200013362000092565b5b6200013f8262000081565b9050602081019050919050565b60005b838110156200016c5780820151818401526020810190506200014f565b60008484015250505050565b60006200018f620001898462000116565b620000f7565b905082815260208101848484011115620001ae57620001ad6200007c565b5b620001bb8482856200014c565b509392505050565b600082601f830112620001db57620001da62000077565b5b8151620001ed84826020860162000178565b91505092915050565b6000806040838503121562000210576200020f6200006d565b5b600083015167ffffffffffffffff81111562000231576200023062000072565b5b6200023f85828601620001c3565b925050602083015167ffffffffffffffff81111562000263576200026262000072565b5b6200027185828601620001c3565b9150509250929050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620002ce57607f821691505b602082108103620002e457620002e362000286565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026200034e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200030f565b6200035a86836200030f565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620003a7620003a16200039b8462000372565b6200037c565b62000372565b9050919050565b6000819050919050565b620003c38362000386565b620003db620003d282620003ae565b8484546200031c565b825550505050565b600090565b620003f2620003e3565b620003ff818484620003b8565b505050565b5b8181101562000427576200041b600082620003e8565b60018101905062000405565b5050565b601f82111562000476576200044081620002ea565b6200044b84620002ff565b810160208510156200045b578190505b620004736200046a85620002ff565b83018262000404565b50505b505050565b600082821c905092915050565b60006200049b600019846008026200047b565b1980831691505092915050565b6000620004b6838362000488565b9150826002028217905092915050565b620004d1826200027b565b67ffffffffffffffff811115620004ed57620004ec62000092565b5b620004f98254620002b5565b620005068282856200042b565b600060209050601f8311600181146200053e576000841562000529578287015190505b620005358582620004a8565b865550620005a5565b601f1984166200054e86620002ea565b60005b82811015620005785784890151825560018201915060208501945060208101905062000551565b8683101562000598578489015162000594601f89168262000488565b8355505b6001600288020188555050505b505050505050565b61238980620005bd6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80636352211e1161008c578063a22cb46511610066578063a22cb46514610224578063b88d4fde14610240578063c87b56dd1461025c578063e985e9c51461028c576100cf565b80636352211e146101a657806370a08231146101d657806395d89b4114610206576100cf565b806301ffc9a7146100d457806306fdde0314610104578063081812fc14610122578063095ea7b31461015257806323b872dd1461016e57806342842e0e1461018a575b600080fd5b6100ee60048036038101906100e99190611411565b6102bc565b6040516100fb9190611459565b60405180910390f35b61010c61039e565b6040516101199190611504565b60405180910390f35b61013c6004803603810190610137919061155c565b610430565b60405161014991906115ca565b60405180910390f35b61016c60048036038101906101679190611611565b6104b5565b005b61018860048036038101906101839190611651565b6105cc565b005b6101a4600480360381019061019f9190611651565b61062c565b005b6101c060048036038101906101bb919061155c565b61064c565b6040516101cd91906115ca565b60405180910390f35b6101f060048036038101906101eb91906116a4565b6106fd565b6040516101fd91906116e0565b60405180910390f35b61020e6107b4565b60405161021b9190611504565b60405180910390f35b61023e60048036038101906102399190611727565b610846565b005b61025a6004803603810190610255919061189c565b6109c6565b005b6102766004803603810190610271919061155c565b610a28565b6040516102839190611504565b60405180910390f35b6102a660048036038101906102a1919061191f565b610acf565b6040516102b39190611459565b60405180910390f35b60007f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061038757507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610397575061039682610b63565b5b9050919050565b6060600080546103ad9061198e565b80601f01602080910402602001604051908101604052809291908181526020018280546103d99061198e565b80156104265780601f106103fb57610100808354040283529160200191610426565b820191906000526020600020905b81548152906001019060200180831161040957829003601f168201915b5050505050905090565b600061043b82610bcd565b61047a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047190611a31565b60405180910390fd5b6004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60006104c08261064c565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610530576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052790611ac3565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1661054f610c39565b73ffffffffffffffffffffffffffffffffffffffff16148061057e575061057d81610578610c39565b610acf565b5b6105bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b490611b55565b60405180910390fd5b6105c78383610c41565b505050565b6105dd6105d7610c39565b82610cfa565b61061c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061390611be7565b60405180910390fd5b610627838383610dd8565b505050565b610647838383604051806020016040528060008152506109c6565b505050565b6000806002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036106f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106eb90611c79565b60405180910390fd5b80915050919050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361076d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076490611d0b565b60405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600180546107c39061198e565b80601f01602080910402602001604051908101604052809291908181526020018280546107ef9061198e565b801561083c5780601f106108115761010080835404028352916020019161083c565b820191906000526020600020905b81548152906001019060200180831161081f57829003601f168201915b5050505050905090565b61084e610c39565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b290611d77565b60405180910390fd5b80600560006108c8610c39565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff16610975610c39565b73ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516109ba9190611459565b60405180910390a35050565b6109d76109d1610c39565b83610cfa565b610a16576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0d90611be7565b60405180910390fd5b610a2284848484611033565b50505050565b6060610a3382610bcd565b610a72576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6990611e09565b60405180910390fd5b6000610a7c61108f565b90506000815111610a9c5760405180602001604052806000815250610ac7565b80610aa6846110a6565b604051602001610ab7929190611e65565b6040516020818303038152906040525b915050919050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b60008073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b600033905090565b816004600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff16610cb48361064c565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000610d0582610bcd565b610d44576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3b90611efb565b60405180910390fd5b6000610d4f8361064c565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480610dbe57508373ffffffffffffffffffffffffffffffffffffffff16610da684610430565b73ffffffffffffffffffffffffffffffffffffffff16145b80610dcf5750610dce8185610acf565b5b91505092915050565b8273ffffffffffffffffffffffffffffffffffffffff16610df88261064c565b73ffffffffffffffffffffffffffffffffffffffff1614610e4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4590611f8d565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610ebd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eb49061201f565b60405180910390fd5b610ec8838383611206565b610ed3600082610c41565b6001600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610f23919061206e565b925050819055506001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610f7a91906120a2565b92505081905550816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b61103e848484610dd8565b61104a8484848461120b565b611089576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161108090612148565b60405180910390fd5b50505050565b606060405180602001604052806000815250905090565b6060600082036110ed576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050611201565b600082905060005b6000821461111f57808061110890612168565b915050600a8261111891906121df565b91506110f5565b60008167ffffffffffffffff81111561113b5761113a611771565b5b6040519080825280601f01601f19166020018201604052801561116d5781602001600182028036833780820191505090505b5090505b600085146111fa57600182611186919061206e565b9150600a856111959190612210565b60306111a191906120a2565b60f81b8183815181106111b7576111b6612241565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a856111f391906121df565b9450611171565b8093505050505b919050565b505050565b600061122c8473ffffffffffffffffffffffffffffffffffffffff16611392565b15611385578373ffffffffffffffffffffffffffffffffffffffff1663150b7a02611255610c39565b8786866040518563ffffffff1660e01b815260040161127794939291906122c5565b6020604051808303816000875af19250505080156112b357506040513d601f19601f820116820180604052508101906112b09190612326565b60015b611335573d80600081146112e3576040519150601f19603f3d011682016040523d82523d6000602084013e6112e8565b606091505b50600081510361132d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161132490612148565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161491505061138a565b600190505b949350505050565b600080823b905060008111915050919050565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6113ee816113b9565b81146113f957600080fd5b50565b60008135905061140b816113e5565b92915050565b600060208284031215611427576114266113af565b5b6000611435848285016113fc565b91505092915050565b60008115159050919050565b6114538161143e565b82525050565b600060208201905061146e600083018461144a565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156114ae578082015181840152602081019050611493565b60008484015250505050565b6000601f19601f8301169050919050565b60006114d682611474565b6114e0818561147f565b93506114f0818560208601611490565b6114f9816114ba565b840191505092915050565b6000602082019050818103600083015261151e81846114cb565b905092915050565b6000819050919050565b61153981611526565b811461154457600080fd5b50565b60008135905061155681611530565b92915050565b600060208284031215611572576115716113af565b5b600061158084828501611547565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006115b482611589565b9050919050565b6115c4816115a9565b82525050565b60006020820190506115df60008301846115bb565b92915050565b6115ee816115a9565b81146115f957600080fd5b50565b60008135905061160b816115e5565b92915050565b60008060408385031215611628576116276113af565b5b6000611636858286016115fc565b925050602061164785828601611547565b9150509250929050565b60008060006060848603121561166a576116696113af565b5b6000611678868287016115fc565b9350506020611689868287016115fc565b925050604061169a86828701611547565b9150509250925092565b6000602082840312156116ba576116b96113af565b5b60006116c8848285016115fc565b91505092915050565b6116da81611526565b82525050565b60006020820190506116f560008301846116d1565b92915050565b6117048161143e565b811461170f57600080fd5b50565b600081359050611721816116fb565b92915050565b6000806040838503121561173e5761173d6113af565b5b600061174c858286016115fc565b925050602061175d85828601611712565b9150509250929050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6117a9826114ba565b810181811067ffffffffffffffff821117156117c8576117c7611771565b5b80604052505050565b60006117db6113a5565b90506117e782826117a0565b919050565b600067ffffffffffffffff82111561180757611806611771565b5b611810826114ba565b9050602081019050919050565b82818337600083830152505050565b600061183f61183a846117ec565b6117d1565b90508281526020810184848401111561185b5761185a61176c565b5b61186684828561181d565b509392505050565b600082601f83011261188357611882611767565b5b813561189384826020860161182c565b91505092915050565b600080600080608085870312156118b6576118b56113af565b5b60006118c4878288016115fc565b94505060206118d5878288016115fc565b93505060406118e687828801611547565b925050606085013567ffffffffffffffff811115611907576119066113b4565b5b6119138782880161186e565b91505092959194509250565b60008060408385031215611936576119356113af565b5b6000611944858286016115fc565b9250506020611955858286016115fc565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806119a657607f821691505b6020821081036119b9576119b861195f565b5b50919050565b7f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860008201527f697374656e7420746f6b656e0000000000000000000000000000000000000000602082015250565b6000611a1b602c8361147f565b9150611a26826119bf565b604082019050919050565b60006020820190508181036000830152611a4a81611a0e565b9050919050565b7f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560008201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b6000611aad60218361147f565b9150611ab882611a51565b604082019050919050565b60006020820190508181036000830152611adc81611aa0565b9050919050565b7f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f7760008201527f6e6572206e6f7220617070726f76656420666f7220616c6c0000000000000000602082015250565b6000611b3f60388361147f565b9150611b4a82611ae3565b604082019050919050565b60006020820190508181036000830152611b6e81611b32565b9050919050565b7f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f60008201527f776e6572206e6f7220617070726f766564000000000000000000000000000000602082015250565b6000611bd160318361147f565b9150611bdc82611b75565b604082019050919050565b60006020820190508181036000830152611c0081611bc4565b9050919050565b7f4552433732313a206f776e657220717565727920666f72206e6f6e657869737460008201527f656e7420746f6b656e0000000000000000000000000000000000000000000000602082015250565b6000611c6360298361147f565b9150611c6e82611c07565b604082019050919050565b60006020820190508181036000830152611c9281611c56565b9050919050565b7f4552433732313a2062616c616e636520717565727920666f7220746865207a6560008201527f726f206164647265737300000000000000000000000000000000000000000000602082015250565b6000611cf5602a8361147f565b9150611d0082611c99565b604082019050919050565b60006020820190508181036000830152611d2481611ce8565b9050919050565b7f4552433732313a20617070726f766520746f2063616c6c657200000000000000600082015250565b6000611d6160198361147f565b9150611d6c82611d2b565b602082019050919050565b60006020820190508181036000830152611d9081611d54565b9050919050565b7f4552433732314d657461646174613a2055524920717565727920666f72206e6f60008201527f6e6578697374656e7420746f6b656e0000000000000000000000000000000000602082015250565b6000611df3602f8361147f565b9150611dfe82611d97565b604082019050919050565b60006020820190508181036000830152611e2281611de6565b9050919050565b600081905092915050565b6000611e3f82611474565b611e498185611e29565b9350611e59818560208601611490565b80840191505092915050565b6000611e718285611e34565b9150611e7d8284611e34565b91508190509392505050565b7f4552433732313a206f70657261746f7220717565727920666f72206e6f6e657860008201527f697374656e7420746f6b656e0000000000000000000000000000000000000000602082015250565b6000611ee5602c8361147f565b9150611ef082611e89565b604082019050919050565b60006020820190508181036000830152611f1481611ed8565b9050919050565b7f4552433732313a207472616e73666572206f6620746f6b656e2074686174206960008201527f73206e6f74206f776e0000000000000000000000000000000000000000000000602082015250565b6000611f7760298361147f565b9150611f8282611f1b565b604082019050919050565b60006020820190508181036000830152611fa681611f6a565b9050919050565b7f4552433732313a207472616e7366657220746f20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b600061200960248361147f565b915061201482611fad565b604082019050919050565b6000602082019050818103600083015261203881611ffc565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061207982611526565b915061208483611526565b925082820390508181111561209c5761209b61203f565b5b92915050565b60006120ad82611526565b91506120b883611526565b92508282019050808211156120d0576120cf61203f565b5b92915050565b7f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560008201527f63656976657220696d706c656d656e7465720000000000000000000000000000602082015250565b600061213260328361147f565b915061213d826120d6565b604082019050919050565b6000602082019050818103600083015261216181612125565b9050919050565b600061217382611526565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036121a5576121a461203f565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006121ea82611526565b91506121f583611526565b925082612205576122046121b0565b5b828204905092915050565b600061221b82611526565b915061222683611526565b925082612236576122356121b0565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081519050919050565b600082825260208201905092915050565b600061229782612270565b6122a1818561227b565b93506122b1818560208601611490565b6122ba816114ba565b840191505092915050565b60006080820190506122da60008301876115bb565b6122e760208301866115bb565b6122f460408301856116d1565b8181036060830152612306818461228c565b905095945050505050565b600081519050612320816113e5565b92915050565b60006020828403121561233c5761233b6113af565b5b600061234a84828501612311565b9150509291505056fea2646970667358221220fb1be3510c119434556b065ac9555a3e4ddcae4e3b6ee61680e47c45ff5b744e64736f6c63430008130033",
}

// ERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721MetaData.ABI instead.
var ERC721ABI = ERC721MetaData.ABI

// IERC721MetadataApproval represents a Approval event raised by the IERC721Metadata contract.
type IERC721MetadataApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721MetadataApprovalIterator, error) {

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

	logs, sub, err := _IERC721Metadata.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataApprovalIterator{contract: _IERC721Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721MetadataApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _IERC721Metadata.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721MetadataApproval)
				if err := _IERC721Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC721Metadata *IERC721MetadataFilterer) ParseApproval(log types.Log) (*IERC721MetadataApproval, error) {
	event := new(IERC721MetadataApproval)
	if err := _IERC721Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721MetadataApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721Metadata contract.
type IERC721MetadataApprovalForAllIterator struct {
	Event *IERC721MetadataApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721MetadataApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721MetadataApprovalForAll)
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
		it.Event = new(IERC721MetadataApprovalForAll)
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
func (it *IERC721MetadataApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721MetadataApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721MetadataApprovalForAll represents a ApprovalForAll event raised by the IERC721Metadata contract.
type IERC721MetadataApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Metadata *IERC721MetadataFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721MetadataApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Metadata.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataApprovalForAllIterator{contract: _IERC721Metadata.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Metadata *IERC721MetadataFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721MetadataApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Metadata.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721MetadataApprovalForAll)
				if err := _IERC721Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_IERC721Metadata *IERC721MetadataFilterer) ParseApprovalForAll(log types.Log) (*IERC721MetadataApprovalForAll, error) {
	event := new(IERC721MetadataApprovalForAll)
	if err := _IERC721Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721Metadata contract.
type IERC721MetadataTransferIterator struct {
	Event *IERC721MetadataTransfer // Event containing the contract specifics and raw log

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
func (it *IERC721MetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721MetadataTransfer)
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
		it.Event = new(IERC721MetadataTransfer)
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
func (it *IERC721MetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721MetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721MetadataTransfer represents a Transfer event raised by the IERC721Metadata contract.
type IERC721MetadataTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721MetadataTransferIterator, error) {

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

	logs, sub, err := _IERC721Metadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataTransferIterator{contract: _IERC721Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721MetadataTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _IERC721Metadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721MetadataTransfer)
				if err := _IERC721Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC721Metadata *IERC721MetadataFilterer) ParseTransfer(log types.Log) (*IERC721MetadataTransfer, error) {
	event := new(IERC721MetadataTransfer)
	if err := _IERC721Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721ReceiverMetaData contains all meta data concerning the IERC721Receiver contract.
var IERC721ReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC721ReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC721ReceiverMetaData.ABI instead.
var IERC721ReceiverABI = IERC721ReceiverMetaData.ABI

// IERC721Receiver is an auto generated Go binding around an Ethereum contract.
type IERC721Receiver struct {
	IERC721ReceiverCaller     // Read-only binding to the contract
	IERC721ReceiverTransactor // Write-only binding to the contract
	IERC721ReceiverFilterer   // Log filterer for contract events
}

// IERC721ReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721ReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721ReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721ReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721ReceiverSession struct {
	Contract     *IERC721Receiver  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721ReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721ReceiverCallerSession struct {
	Contract *IERC721ReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IERC721ReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721ReceiverTransactorSession struct {
	Contract     *IERC721ReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IERC721ReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721ReceiverRaw struct {
	Contract *IERC721Receiver // Generic contract binding to access the raw methods on
}

// IERC721ReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721ReceiverCallerRaw struct {
	Contract *IERC721ReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721ReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721ReceiverTransactorRaw struct {
	Contract *IERC721ReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Receiver creates a new instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721Receiver(address common.Address, backend bind.ContractBackend) (*IERC721Receiver, error) {
	contract, err := bindIERC721Receiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Receiver{IERC721ReceiverCaller: IERC721ReceiverCaller{contract: contract}, IERC721ReceiverTransactor: IERC721ReceiverTransactor{contract: contract}, IERC721ReceiverFilterer: IERC721ReceiverFilterer{contract: contract}}, nil
}

// NewIERC721ReceiverCaller creates a new read-only instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverCaller(address common.Address, caller bind.ContractCaller) (*IERC721ReceiverCaller, error) {
	contract, err := bindIERC721Receiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverCaller{contract: contract}, nil
}

// NewIERC721ReceiverTransactor creates a new write-only instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721ReceiverTransactor, error) {
	contract, err := bindIERC721Receiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverTransactor{contract: contract}, nil
}

// NewIERC721ReceiverFilterer creates a new log filterer instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721ReceiverFilterer, error) {
	contract, err := bindIERC721Receiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverFilterer{contract: contract}, nil
}

// bindIERC721Receiver binds a generic wrapper to an already deployed contract.
func bindIERC721Receiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721ReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Receiver *IERC721ReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Receiver.Contract.IERC721ReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Receiver *IERC721ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.IERC721ReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Receiver *IERC721ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.IERC721ReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Receiver *IERC721ReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Receiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Receiver *IERC721ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Receiver *IERC721ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.contract.Transact(opts, method, params...)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.OnERC721Received(&_IERC721Receiver.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.OnERC721Received(&_IERC721Receiver.TransactOpts, operator, from, tokenId, data)
}

// MathMetaData contains all meta data concerning the Math contract.
var MathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ce99fef1524ff99fb01e1f85f06090ad10a2f76c97737945f99ae93c007563ad64736f6c63430008130033",
}

// MathABI is the input ABI used to generate the binding from.
// Deprecated: Use MathMetaData.ABI instead.
var MathABI = MathMetaData.ABI

// MathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MathMetaData.Bin instead.
var MathBin = MathMetaData.Bin

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := MathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// SafeCastMetaData contains all meta data concerning the SafeCast contract.
var SafeCastMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208845b917a644c445e474788dafb2895e18348825c919f69032b30088a13f102b64736f6c63430008130033",
}

// SafeCastABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeCastMetaData.ABI instead.
var SafeCastABI = SafeCastMetaData.ABI

// SafeCastBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeCastMetaData.Bin instead.
var SafeCastBin = SafeCastMetaData.Bin

// DeploySafeCast deploys a new Ethereum contract, binding an instance of SafeCast to it.
func DeploySafeCast(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeCast, error) {
	parsed, err := SafeCastMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeCastBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// SafeCast is an auto generated Go binding around an Ethereum contract.
type SafeCast struct {
	SafeCastCaller     // Read-only binding to the contract
	SafeCastTransactor // Write-only binding to the contract
	SafeCastFilterer   // Log filterer for contract events
}

// SafeCastCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeCastCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeCastTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeCastFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeCastSession struct {
	Contract     *SafeCast         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeCastCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeCastCallerSession struct {
	Contract *SafeCastCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeCastTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeCastTransactorSession struct {
	Contract     *SafeCastTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeCastRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeCastRaw struct {
	Contract *SafeCast // Generic contract binding to access the raw methods on
}

// SafeCastCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeCastCallerRaw struct {
	Contract *SafeCastCaller // Generic read-only contract binding to access the raw methods on
}

// SafeCastTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeCastTransactorRaw struct {
	Contract *SafeCastTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeCast creates a new instance of SafeCast, bound to a specific deployed contract.
func NewSafeCast(address common.Address, backend bind.ContractBackend) (*SafeCast, error) {
	contract, err := bindSafeCast(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// NewSafeCastCaller creates a new read-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastCaller(address common.Address, caller bind.ContractCaller) (*SafeCastCaller, error) {
	contract, err := bindSafeCast(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastCaller{contract: contract}, nil
}

// NewSafeCastTransactor creates a new write-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeCastTransactor, error) {
	contract, err := bindSafeCast(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastTransactor{contract: contract}, nil
}

// NewSafeCastFilterer creates a new log filterer instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeCastFilterer, error) {
	contract, err := bindSafeCast(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeCastFilterer{contract: contract}, nil
}

// bindSafeCast binds a generic wrapper to an already deployed contract.
func bindSafeCast(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeCastABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.SafeCastCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transact(opts, method, params...)
}

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206b5ef8b0c3688c5123425dff80caa6213642336c7d6195a5211aa3f0ec616ac164736f6c63430008130033",
}

// StringsABI is the input ABI used to generate the binding from.
// Deprecated: Use StringsMetaData.ABI instead.
var StringsABI = StringsMetaData.ABI

// StringsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StringsMetaData.Bin instead.
var StringsBin = StringsMetaData.Bin

// DeployStrings deploys a new Ethereum contract, binding an instance of Strings to it.
func DeployStrings(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Strings, error) {
	parsed, err := StringsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StringsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// Strings is an auto generated Go binding around an Ethereum contract.
type Strings struct {
	StringsCaller     // Read-only binding to the contract
	StringsTransactor // Write-only binding to the contract
	StringsFilterer   // Log filterer for contract events
}

// StringsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StringsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StringsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StringsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StringsSession struct {
	Contract     *Strings          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StringsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StringsCallerSession struct {
	Contract *StringsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StringsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StringsTransactorSession struct {
	Contract     *StringsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StringsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StringsRaw struct {
	Contract *Strings // Generic contract binding to access the raw methods on
}

// StringsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StringsCallerRaw struct {
	Contract *StringsCaller // Generic read-only contract binding to access the raw methods on
}

// StringsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StringsTransactorRaw struct {
	Contract *StringsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrings creates a new instance of Strings, bound to a specific deployed contract.
func NewStrings(address common.Address, backend bind.ContractBackend) (*Strings, error) {
	contract, err := bindStrings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// NewStringsCaller creates a new read-only instance of Strings, bound to a specific deployed contract.
func NewStringsCaller(address common.Address, caller bind.ContractCaller) (*StringsCaller, error) {
	contract, err := bindStrings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StringsCaller{contract: contract}, nil
}

// NewStringsTransactor creates a new write-only instance of Strings, bound to a specific deployed contract.
func NewStringsTransactor(address common.Address, transactor bind.ContractTransactor) (*StringsTransactor, error) {
	contract, err := bindStrings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StringsTransactor{contract: contract}, nil
}

// NewStringsFilterer creates a new log filterer instance of Strings, bound to a specific deployed contract.
func NewStringsFilterer(address common.Address, filterer bind.ContractFilterer) (*StringsFilterer, error) {
	contract, err := bindStrings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StringsFilterer{contract: contract}, nil
}

// bindStrings binds a generic wrapper to an already deployed contract.
func bindStrings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StringsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.StringsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transact(opts, method, params...)
}
