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
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea264697066735822122053e1a83c441c07ad4262e278ac750d5e876ea21b9ae1aefd5d5e80acc24d86f464736f6c63430008150033",
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newActivationTime\",\"type\":\"uint256\"}],\"name\":\"ActivationTimeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lastBidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"randomWalkNFTId\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prizeTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"BidEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCharity\",\"type\":\"address\"}],\"name\":\"CharityAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCharityPercentage\",\"type\":\"uint256\"}],\"name\":\"CharityPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCosmicSignature\",\"type\":\"address\"}],\"name\":\"CosmicSignatureAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCosmicToken\",\"type\":\"address\"}],\"name\":\"CosmicTokenAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddressdonatedNFTs\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"DonatedNFTClaimedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newInitialBidAmountFraction\",\"type\":\"uint256\"}],\"name\":\"InitialBidAmountFractionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newInitialSecondsUntilPrize\",\"type\":\"uint256\"}],\"name\":\"InitialSecondsUntilPrizeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"NFTDonationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNanoSecondsExtra\",\"type\":\"uint256\"}],\"name\":\"NanoSecondsExtraChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNumHolderNFTWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"NumHolderNFTWinnersPerRoundChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNumRaffleNFTWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"NumRaffleNFTWinnersPerRoundChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNumRaffleWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"NumRaffleWinnersPerRoundChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPriceIncrease\",\"type\":\"uint256\"}],\"name\":\"PriceIncreaseChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PrizeClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrizePercentage\",\"type\":\"uint256\"}],\"name\":\"PrizePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"}],\"name\":\"RaffleNFTWinnerEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRafflePercentage\",\"type\":\"uint256\"}],\"name\":\"RafflePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRaffleWallet\",\"type\":\"address\"}],\"name\":\"RaffleWalletAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRandomWalk\",\"type\":\"address\"}],\"name\":\"RandomWalkAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimeIncrease\",\"type\":\"uint256\"}],\"name\":\"TimeIncreaseChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimeout\",\"type\":\"uint256\"}],\"name\":\"TimeoutClaimPrizeChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_MESSAGE_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MILLION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_REWARD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"bidAndDonateNFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"randomWalkNFTId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"bidWithRWLK\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"randomWalkNFTId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"bidWithRWLKAndDonateNFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charity\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charityAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charityPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"claimDonatedNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokens\",\"type\":\"uint256[]\"}],\"name\":\"claimManyDonatedNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"donatedNFTs\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialBidAmountFraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialSecondsUntilPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nanoSecondsExtra\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractCosmicSignature\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numDonatedNFTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numHolderNFTWinnersPerRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleNFTWinnersPerRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleParticipants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleWinnersPerRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizeTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raffleAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raffleEntropy\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"raffleParticipants\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rafflePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raffleWallet\",\"outputs\":[{\"internalType\":\"contractRaffleWallet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomWalk\",\"outputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newActivationTime\",\"type\":\"uint256\"}],\"name\":\"setActivationTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setCharity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCharityPercentage\",\"type\":\"uint256\"}],\"name\":\"setCharityPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newInitialSecondsUntilPrize\",\"type\":\"uint256\"}],\"name\":\"setInitialSecondsUntilPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNanoSecondsExtra\",\"type\":\"uint256\"}],\"name\":\"setNanoSecondsExtra\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setNftContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNumHolderNFTWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"setNumHolderNFTWinnersPerRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNumRaffleNFTWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleNFTWinnersPerRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNumRaffleWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleWinnersPerRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPriceIncrease\",\"type\":\"uint256\"}],\"name\":\"setPriceIncrease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRafflePercentage\",\"type\":\"uint256\"}],\"name\":\"setRafflePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setRaffleWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setRandomWalk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimeIncrease\",\"type\":\"uint256\"}],\"name\":\"setTimeIncrease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimeout\",\"type\":\"uint256\"}],\"name\":\"setTimeoutClaimPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeUntilActivation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeUntilPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeoutClaimPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractCosmicToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newInitialBidAmountFraction\",\"type\":\"uint256\"}],\"name\":\"updateInitialBidAmountFraction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPrizePercentage\",\"type\":\"uint256\"}],\"name\":\"updatePrizePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usedRandomWalkNFTs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"winners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052620f695060015565034630b8a000600255620f42a46003556201518060045566038d7ea4c680006005555f60065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506201518060085560c86009556019600a55600a600b556005600c556003600d556005600e556002600f555f60115563644709f0601455348015620000b0575f80fd5b50620000d1620000c56200015860201b60201c565b6200015f60201b60201c565b424340604051602001620000e7929190620002b2565b60405160208183030381529060405280519060200120601581905550620001136200015860201b60201c565b60075f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620002f2565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f82825260208201905092915050565b7f436f736d6963205369676e6174757265203230323300000000000000000000005f82015250565b5f6200026660158362000220565b9150620002738262000230565b602082019050919050565b5f819050919050565b62000292816200027e565b82525050565b5f819050919050565b620002ac8162000298565b82525050565b5f6060820190508181035f830152620002cb8162000258565b9050620002dc602083018562000287565b620002eb6040830184620002a1565b9392505050565b615d2780620003005f395ff3fe6080604052600436106103f2575f3560e01c80638b12227411610207578063cb819dc011610117578063e1381d7e116100aa578063f2fde38b11610079578063f2fde38b14610e2d578063f717882214610e55578063f8c3405014610e7f578063fb6f71a314610ea9578063fc0c546a14610ed157610410565b8063e1381d7e14610da7578063ec34866d14610dcf578063ed08339814610df9578063ed88c68e14610e2357610410565b8063d94d0316116100e6578063d94d031614610cff578063da4493f614610d29578063dbc945c014610d53578063dec08d8e14610d7d57610410565b8063cb819dc014610c44578063d4ba890214610c6e578063d59d747814610c98578063d6e1741714610cc057610410565b8063a672f6e11161019a578063bb0bc58e11610169578063bb0bc58e14610b86578063bbcd5bbe14610ba2578063c01c5de214610bca578063c7c8378d14610bf2578063c94028c214610c1a57610410565b8063a672f6e114610ae4578063a6ceac2c14610b0c578063a6f9cc1514610b36578063ae6247f214610b5e57610410565b80639250c33c116101d65780639250c33c14610a38578063934aa02314610a62578063a2fb117514610a8c578063a325b7d214610ac857610410565b80638b122274146109925780638b1329e0146109ba5780638da5cb5b146109e45780639136d6d914610a0e57610410565b806347ccca021161030257806370740ac9116102955780637aef951c116102645780637aef951c146108be5780637c5486a2146108da57806380de163d146109025780638547af301461092c57806386e378c91461095657610410565b806370740ac91461083e578063715018a614610854578063785fa6271461086a578063799d431d1461089457610410565b806351964588116102d1578063519645881461079e57806352f5ad77146107c65780635e6e47aa146107ee578063647b3e7f1461081657610410565b806347ccca02146106f85780634a773f33146107225780634ac3a3951461074c5780635111a2d61461077457610410565b80632a02721111610385578063355f01e211610354578063355f01e21461062a5780633bec7b69146106545780633c83adc41461067e5780633f7909d4146106a657806340e02322146106ce57610410565b80632a027211146105825780632aab3223146105ac57806332bc934c146105d657806332d382cd1461060057610410565b8063119b22b3116103c1578063119b22b3146104c857806311dc7335146104f2578063150b7a021461051c57806319afe4731461055857610410565b8063019d354a1461041457806304a57c091461043c57806305ba9b6714610478578063062fb100146104a057610410565b366104105761040e60405180602001604052805f815250610efb565b005b5f80fd5b34801561041f575f80fd5b5061043a60048036038101906104359190614477565b6110db565b005b348015610447575f80fd5b50610462600480360381019061045d9190614477565b61119a565b60405161046f91906144e1565b60405180910390f35b348015610483575f80fd5b5061049e60048036038101906104999190614477565b6111ca565b005b3480156104ab575f80fd5b506104c660048036038101906104c19190614524565b611289565b005b3480156104d3575f80fd5b506104dc6113ed565b6040516104e9919061455e565b60405180910390f35b3480156104fd575f80fd5b506105066113f3565b604051610513919061455e565b60405180910390f35b348015610527575f80fd5b50610542600480360381019061053d91906145d8565b6113f9565b60405161054f9190614696565b60405180910390f35b348015610563575f80fd5b5061056c61140d565b604051610579919061455e565b60405180910390f35b34801561058d575f80fd5b50610596611413565b6040516105a3919061455e565b60405180910390f35b3480156105b7575f80fd5b506105c0611419565b6040516105cd919061455e565b60405180910390f35b3480156105e1575f80fd5b506105ea61141f565b6040516105f7919061455e565b60405180910390f35b34801561060b575f80fd5b50610614611426565b604051610621919061470a565b60405180910390f35b348015610635575f80fd5b5061063e61144b565b60405161064b919061455e565b60405180910390f35b34801561065f575f80fd5b50610668611451565b604051610675919061455e565b60405180910390f35b348015610689575f80fd5b506106a4600480360381019061069f9190614477565b611457565b005b3480156106b1575f80fd5b506106cc60048036038101906106c7919061486b565b611575565b005b3480156106d9575f80fd5b506106e26115ba565b6040516106ef919061455e565b60405180910390f35b348015610703575f80fd5b5061070c6115c0565b60405161071991906148d2565b60405180910390f35b34801561072d575f80fd5b506107366115e5565b604051610743919061455e565b60405180910390f35b348015610757575f80fd5b50610772600480360381019061076d9190614477565b6115eb565b005b34801561077f575f80fd5b506107886116aa565b604051610795919061490b565b60405180910390f35b3480156107a9575f80fd5b506107c460048036038101906107bf9190614477565b6116cf565b005b3480156107d1575f80fd5b506107ec60048036038101906107e79190614524565b61178e565b005b3480156107f9575f80fd5b50610814600480360381019061080f9190614477565b6118f2565b005b348015610821575f80fd5b5061083c60048036038101906108379190614477565b611a10565b005b348015610849575f80fd5b50610852611acf565b005b34801561085f575f80fd5b50610868612a81565b005b348015610875575f80fd5b5061087e612b08565b60405161088b919061455e565b60405180910390f35b34801561089f575f80fd5b506108a8612b28565b6040516108b5919061455e565b60405180910390f35b6108d860048036038101906108d391906149d4565b610efb565b005b3480156108e5575f80fd5b5061090060048036038101906108fb9190614477565b612b2e565b005b34801561090d575f80fd5b50610916612bed565b6040516109239190614a33565b60405180910390f35b348015610937575f80fd5b50610940612bf3565b60405161094d91906144e1565b60405180910390f35b348015610961575f80fd5b5061097c60048036038101906109779190614477565b612c18565b6040516109899190614a66565b60405180910390f35b34801561099d575f80fd5b506109b860048036038101906109b39190614477565b612c35565b005b3480156109c5575f80fd5b506109ce612cf4565b6040516109db919061455e565b60405180910390f35b3480156109ef575f80fd5b506109f8612d1b565b604051610a0591906144e1565b60405180910390f35b348015610a19575f80fd5b50610a22612d42565b604051610a2f919061455e565b60405180910390f35b348015610a43575f80fd5b50610a4c612d48565b604051610a59919061455e565b60405180910390f35b348015610a6d575f80fd5b50610a76612d4e565b604051610a8391906144e1565b60405180910390f35b348015610a97575f80fd5b50610ab26004803603810190610aad9190614477565b612d73565b604051610abf91906144e1565b60405180910390f35b610ae26004803603810190610add9190614aba565b612da3565b005b348015610aef575f80fd5b50610b0a6004803603810190610b059190614477565b612dbd565b005b348015610b17575f80fd5b50610b20612e7c565b604051610b2d919061455e565b60405180910390f35b348015610b41575f80fd5b50610b5c6004803603810190610b579190614524565b612e82565b005b348015610b69575f80fd5b50610b846004803603810190610b7f9190614b3a565b612fe6565b005b610ba06004803603810190610b9b9190614b94565b613201565b005b348015610bad575f80fd5b50610bc86004803603810190610bc39190614524565b613219565b005b348015610bd5575f80fd5b50610bf06004803603810190610beb9190614477565b61337d565b005b348015610bfd575f80fd5b50610c186004803603810190610c139190614477565b61343c565b005b348015610c25575f80fd5b50610c2e61355a565b604051610c3b919061455e565b60405180910390f35b348015610c4f575f80fd5b50610c5861357a565b604051610c65919061455e565b60405180910390f35b348015610c79575f80fd5b50610c82613580565b604051610c8f919061455e565b60405180910390f35b348015610ca3575f80fd5b50610cbe6004803603810190610cb99190614477565b61358d565b005b348015610ccb575f80fd5b50610ce66004803603810190610ce19190614477565b613867565b604051610cf69493929190614c20565b60405180910390f35b348015610d0a575f80fd5b50610d136138be565b604051610d20919061455e565b60405180910390f35b348015610d34575f80fd5b50610d3d6138c4565b604051610d4a919061455e565b60405180910390f35b348015610d5e575f80fd5b50610d676138ca565b604051610d74919061455e565b60405180910390f35b348015610d88575f80fd5b50610d916138ea565b604051610d9e919061455e565b60405180910390f35b348015610db2575f80fd5b50610dcd6004803603810190610dc89190614477565b6138f0565b005b348015610dda575f80fd5b50610de36139af565b604051610df0919061455e565b60405180910390f35b348015610e04575f80fd5b50610e0d6139d3565b604051610e1a919061455e565b60405180910390f35b610e2b6139d9565b005b348015610e38575f80fd5b50610e536004803603810190610e4e9190614524565b613acf565b005b348015610e60575f80fd5b50610e69613bc5565b604051610e76919061455e565b60405180910390f35b348015610e8a575f80fd5b50610e93613bec565b604051610ea0919061455e565b60405180910390f35b348015610eb4575f80fd5b50610ecf6004803603810190610eca9190614524565b613bf2565b005b348015610edc575f80fd5b50610ee5613d77565b604051610ef29190614c83565b60405180910390f35b5f610f046139af565b905080341015610f49576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f4090614d1c565b60405180910390fd5b80600581905550610f5982613d9c565b60055434111561103b575f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660055434610faa9190614d67565b604051610fb690614dc7565b5f6040518083038185875af1925050503d805f8114610ff0576040519150601f19603f3d011682016040523d82523d5f602084013e610ff5565b606091505b5050905080611039576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103090614e25565b60405180910390fd5b505b60115460065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fc7beef8f8d25068377c10e7b7f30ef5622f443eb674e07835df1daf5fe84acbf6005547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff601054876040516110cf9493929190614eef565b60405180910390a35050565b6110e36140d8565b73ffffffffffffffffffffffffffffffffffffffff16611101612d1b565b73ffffffffffffffffffffffffffffffffffffffff1614611157576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161114e90614f83565b60405180910390fd5b806004819055507fcaf8e5c6bc1bb2c19935f84ddcdaefb232ad06f9f2abd2ad588bea4bbe631d2660045460405161118f919061455e565b60405180910390a150565b6016602052805f5260405f205f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6111d26140d8565b73ffffffffffffffffffffffffffffffffffffffff166111f0612d1b565b73ffffffffffffffffffffffffffffffffffffffff1614611246576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161123d90614f83565b60405180910390fd5b806002819055507f678d086a09e1be49401b1e3a6e14db1878e8d8b88e71d0cfe24a32726d0e38b960025460405161127e919061455e565b60405180910390a150565b6112916140d8565b73ffffffffffffffffffffffffffffffffffffffff166112af612d1b565b73ffffffffffffffffffffffffffffffffffffffff1614611305576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112fc90614f83565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611373576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161136a90614feb565b60405180910390fd5b8060185f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f508d510537704d37ed64691a706392abc33e59ce8a327b4952e112820ceb88a6816040516113e291906144e1565b60405180910390a150565b60115481565b600a5481565b5f63150b7a0260e01b905095945050505050565b60055481565b61011881565b601a5481565b620f424081565b60185f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60045481565b60175481565b61145f6140d8565b73ffffffffffffffffffffffffffffffffffffffff1661147d612d1b565b73ffffffffffffffffffffffffffffffffffffffff16146114d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114ca90614f83565b60405180910390fd5b80600a819055506064600c54600b54600a546114ef9190615009565b6114f99190615009565b10611539576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611530906150ac565b60405180910390fd5b7f595fa5ba64cd6f66df19b53b59bb4a275bae1ec5b362da95e1fa4c5feb62be1e600a5460405161156a919061455e565b60405180910390a150565b5f5b81518110156115b6576115a3828281518110611596576115956150ca565b5b602002602001015161358d565b80806115ae906150f7565b915050611577565b5050565b60095481565b601c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600d5481565b6115f36140d8565b73ffffffffffffffffffffffffffffffffffffffff16611611612d1b565b73ffffffffffffffffffffffffffffffffffffffff1614611667576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161165e90614f83565b60405180910390fd5b806003819055507fed46e73b174ced51fb529cdc6c1b4d8abf49387e6d849b71648afb63c81d12cd60035460405161169f919061455e565b60405180910390a150565b601d5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6116d76140d8565b73ffffffffffffffffffffffffffffffffffffffff166116f5612d1b565b73ffffffffffffffffffffffffffffffffffffffff161461174b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161174290614f83565b60405180910390fd5b806008819055507f6da281754ba85ee0c5983a8e8f05a92910c2a0c5b80e68c126216d65f162a305600854604051611783919061455e565b60405180910390a150565b6117966140d8565b73ffffffffffffffffffffffffffffffffffffffff166117b4612d1b565b73ffffffffffffffffffffffffffffffffffffffff161461180a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161180190614f83565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611878576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161186f90614feb565b60405180910390fd5b80601c5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f7142a592d5404b5fc4a294c66f70e32b2a7776bd807f722d59268def87c765d1816040516118e791906144e1565b60405180910390a150565b6118fa6140d8565b73ffffffffffffffffffffffffffffffffffffffff16611918612d1b565b73ffffffffffffffffffffffffffffffffffffffff161461196e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161196590614f83565b60405180910390fd5b80600b819055506064600c54600b54600a5461198a9190615009565b6119949190615009565b106119d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119cb906150ac565b60405180910390fd5b7f0918b499c15fcb0e087d411d53664cde23577e1aa4a9cbfbdf735ddd1871e7d5600b54604051611a05919061455e565b60405180910390a150565b611a186140d8565b73ffffffffffffffffffffffffffffffffffffffff16611a36612d1b565b73ffffffffffffffffffffffffffffffffffffffff1614611a8c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a8390614f83565b60405180910390fd5b80600d819055507f5e3567ae359e61c47ad2f11d2828dea7e062b2ff63dea385cdb7900a50504c7c600d54604051611ac4919061455e565b60405180910390a150565b426010541115611b14576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b0b90615188565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff1660065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611ba3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b9a906151f0565b60405180910390fd5b60045460105442611bb49190614d67565b1015611c515760065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16611bfa6140d8565b73ffffffffffffffffffffffffffffffffffffffff1614611c50576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c47906152a4565b60405180910390fd5b5b5f60065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f611c9a6140d8565b90508060135f60115481526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160115f828254611cff9190615009565b925050819055505f601c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1960e01b836001601154611d579190614d67565b604051602401611d689291906152c2565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051611dd29190615323565b5f604051808303815f865af19150503d805f8114611e0b576040519150601f19603f3d011682016040523d82523d5f602084013e611e10565b606091505b5050905080611e54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e4b906153a9565b60405180910390fd5b5f805b600e5481101561207c57611e696140df565b5f60165f6017546015545f1c611e7f91906153f4565b81526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f8173ffffffffffffffffffffffffffffffffffffffff163b1115611ed35750612069565b5f601c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1960e01b836001601154611f249190614d67565b604051602401611f359291906152c2565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051611f9f9190615323565b5f604051808303815f865af19150503d805f8114611fd8576040519150601f19603f3d011682016040523d82523d5f602084013e611fdd565b606091505b509150505f81806020019051810190611ff69190615438565b90508060016011546120089190614d67565b8473ffffffffffffffffffffffffffffffffffffffff167f0afb73a48ffcc4cf27ecb1889154b910544b89c39c5e5822c431062629134e568860405161204e919061455e565b60405180910390a46001856120639190615009565b94505050505b8080612074906150f7565b915050611e57565b505f601d5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156120e8573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061210c9190615438565b90505f601c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612179573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061219d9190615438565b90505f5b600f548110156126b7575f831115612427576121bb6140df565b5f836015545f1c6121cc91906153f4565b90505f601d5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e836040518263ffffffff1660e01b8152600401612229919061455e565b602060405180830381865afa158015612244573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906122689190615477565b90505f8173ffffffffffffffffffffffffffffffffffffffff163b11156122905750506126a4565b5f601c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1960e01b8360016011546122e19190614d67565b6040516024016122f29291906152c2565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161235c9190615323565b5f604051808303815f865af19150503d805f8114612395576040519150601f19603f3d011682016040523d82523d5f602084013e61239a565b606091505b509150505f818060200190518101906123b39190615438565b90508060016011546123c59190614d67565b8473ffffffffffffffffffffffffffffffffffffffff167f0afb73a48ffcc4cf27ecb1889154b910544b89c39c5e5822c431062629134e568b60405161240b919061455e565b60405180910390a46001886124209190615009565b9750505050505b5f8211156126a3576124376140df565b5f826015545f1c61244891906153f4565b90505f601c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e836040518263ffffffff1660e01b81526004016124a5919061455e565b602060405180830381865afa1580156124c0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906124e49190615477565b90505f8173ffffffffffffffffffffffffffffffffffffffff163b111561250c5750506126a4565b5f601c5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1960e01b83600160115461255d9190614d67565b60405160240161256e9291906152c2565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516125d89190615323565b5f604051808303815f865af19150503d805f8114612611576040519150601f19603f3d011682016040523d82523d5f602084013e612616565b606091505b509150505f8180602001905181019061262f9190615438565b90508060016011546126419190614d67565b8473ffffffffffffffffffffffffffffffffffffffff167f0afb73a48ffcc4cf27ecb1889154b910544b89c39c5e5822c431062629134e568b604051612687919061455e565b60405180910390a460018861269c9190615009565b9750505050505b5b80806126af906150f7565b9150506121a1565b505f6126c1612b08565b90505f6126cc6138ca565b90505f6126d761355a565b90505f8873ffffffffffffffffffffffffffffffffffffffff16846040516126fe90614dc7565b5f6040518083038185875af1925050503d805f8114612738576040519150601f19603f3d011682016040523d82523d5f602084013e61273d565b606091505b5050905080612781576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612778906154ec565b60405180910390fd5b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16836040516127c690614dc7565b5f6040518083038185875af1925050503d805f8114612800576040519150601f19603f3d011682016040523d82523d5f602084013e612805565b606091505b5050809150508061284b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128429061557a565b60405180910390fd5b5f5b600d54811015612a095761285f6140df565b5f60165f6017546015545f1c61287591906153f4565b81526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060185f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168463f340fa0160e01b8360016011546128f59190614d67565b6040516024016129069291906152c2565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516129709190615323565b5f6040518083038185875af1925050503d805f81146129aa576040519150601f19603f3d011682016040523d82523d5f602084013e6129af565b606091505b505080935050826129f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129ec906155e2565b60405180910390fd5b508080612a01906150f7565b91505061284d565b505f601781905550612a19614115565b8873ffffffffffffffffffffffffffffffffffffffff166001601154612a3f9190614d67565b7f27bc828c399c2947fea27bca8a75ced2e94ff2651d607271f051e39db52286ce86604051612a6e919061455e565b60405180910390a3505050505050505050565b612a896140d8565b73ffffffffffffffffffffffffffffffffffffffff16612aa7612d1b565b73ffffffffffffffffffffffffffffffffffffffff1614612afd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612af490614f83565b60405180910390fd5b612b065f61412b565b565b5f6064600a5447612b199190615600565b612b239190615641565b905090565b600b5481565b612b366140d8565b73ffffffffffffffffffffffffffffffffffffffff16612b54612d1b565b73ffffffffffffffffffffffffffffffffffffffff1614612baa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ba190614f83565b60405180910390fd5b806014819055507f584ff9a7b3a86db5397052f2e440da0ed60a95f646a3a884863cd92262e682b6601454604051612be2919061455e565b60405180910390a150565b60155481565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6012602052805f5260405f205f915054906101000a900460ff1681565b612c3d6140d8565b73ffffffffffffffffffffffffffffffffffffffff16612c5b612d1b565b73ffffffffffffffffffffffffffffffffffffffff1614612cb1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ca890614f83565b60405180910390fd5b806001819055507fcbc1f49adfa29e2f2f0f5c9e057722496a4bc95a6a5446deaa423a02b30c64ac600154604051612ce9919061455e565b60405180910390a150565b5f426010541015612d07575f9050612d18565b42601054612d159190614d67565b90505b90565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60025481565b600c5481565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6013602052805f5260405f205f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b612dad8484612fe6565b612db782826141ec565b50505050565b612dc56140d8565b73ffffffffffffffffffffffffffffffffffffffff16612de3612d1b565b73ffffffffffffffffffffffffffffffffffffffff1614612e39576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e3090614f83565b60405180910390fd5b806009819055507f3b311f029da1b90c3de2e0b3168436c5ed7d8b0ae81b7d4894c12da03835c628600954604051612e71919061455e565b60405180910390a150565b600e5481565b612e8a6140d8565b73ffffffffffffffffffffffffffffffffffffffff16612ea8612d1b565b73ffffffffffffffffffffffffffffffffffffffff1614612efe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ef590614f83565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612f6c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f6390614feb565b60405180910390fd5b80601d5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f9dc3707e4b88be65295b0083b44ffa94059c80372add9b0d24d6a3b371a03b9281604051612fdb91906144e1565b60405180910390a150565b60125f8381526020019081526020015f205f9054906101000a900460ff1615613044576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161303b906156e1565b60405180910390fd5b61304c6140d8565b73ffffffffffffffffffffffffffffffffffffffff16601d5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e846040518263ffffffff1660e01b81526004016130bc919061455e565b602060405180830381865afa1580156130d7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130fb9190615477565b73ffffffffffffffffffffffffffffffffffffffff1614613151576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131489061576f565b60405180910390fd5b600160125f8481526020019081526020015f205f6101000a81548160ff02191690831515021790555061318381613d9c565b60115460065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fc7beef8f8d25068377c10e7b7f30ef5622f443eb674e07835df1daf5fe84acbf5f85601054866040516131f594939291906157d5565b60405180910390a35050565b61320a83610efb565b61321482826141ec565b505050565b6132216140d8565b73ffffffffffffffffffffffffffffffffffffffff1661323f612d1b565b73ffffffffffffffffffffffffffffffffffffffff1614613295576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161328c90614f83565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603613303576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016132fa90614feb565b60405180910390fd5b80601b5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f3ab2dcf539aa3c30314265602fc86dee8e343e3c3905730f870835f36c400deb8160405161337291906144e1565b60405180910390a150565b6133856140d8565b73ffffffffffffffffffffffffffffffffffffffff166133a3612d1b565b73ffffffffffffffffffffffffffffffffffffffff16146133f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133f090614f83565b60405180910390fd5b80600f819055507f0cc7e37c68566f67d0fe13bf38246d7447cf99a0c481c2ef9963969bb4f5ebc8600f54604051613431919061455e565b60405180910390a150565b6134446140d8565b73ffffffffffffffffffffffffffffffffffffffff16613462612d1b565b73ffffffffffffffffffffffffffffffffffffffff16146134b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016134af90614f83565b60405180910390fd5b80600c819055506064600c54600b54600a546134d49190615009565b6134de9190615009565b1061351e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613515906150ac565b60405180910390fd5b7fd2002a256ab2f8a5c1b865685754b20437c0a26e7249e40ae5df2993966f99f2600c5460405161354f919061455e565b60405180910390a150565b5f6064600c544761356b9190615600565b6135759190615641565b905090565b60105481565b68056bc75e2d6310000081565b601a5481106135d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016135c890615869565b60405180910390fd5b5f60135f60195f8581526020019081526020015f206002015481526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff166136386140d8565b73ffffffffffffffffffffffffffffffffffffffff161461368e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613685906158f7565b60405180910390fd5b60195f8381526020019081526020015f206003015f9054906101000a900460ff16156136ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136e690615985565b60405180910390fd5b600160195f8481526020019081526020015f206003015f6101000a81548160ff02191690831515021790555060195f8381526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166342842e0e308360195f8781526020019081526020015f20600101546040518463ffffffff1660e01b815260040161379e939291906159a3565b5f604051808303815f87803b1580156137b5575f80fd5b505af11580156137c7573d5f803e3d5ffd5b5050505060195f8381526020019081526020015f20600201547f0d0e4b90a96d931964d5b0421a6b5b48dd73abf538cdb9ae015069d49e9a4679838360195f8781526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660195f8881526020019081526020015f206001015460405161385b94939291906159d8565b60405180910390a25050565b6019602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015490806002015490806003015f9054906101000a900460ff16905084565b60035481565b60145481565b5f6064600b54476138db9190615600565b6138e59190615641565b905090565b600f5481565b6138f86140d8565b73ffffffffffffffffffffffffffffffffffffffff16613916612d1b565b73ffffffffffffffffffffffffffffffffffffffff161461396c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161396390614f83565b60405180910390fd5b80600e819055507f72e4278828b8a868e0ba5b4887f954797ea786f8bac991128636171f5eed471d600e546040516139a4919061455e565b60405180910390a150565b5f620f42406001546005546139c49190615600565b6139ce9190615641565b905090565b60085481565b5f3411613a1b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613a1290615a8b565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff1660065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603613a7857613a77614115565b5b613a806140d8565b73ffffffffffffffffffffffffffffffffffffffff167f8b7fe5be5699654fd637d2250cb0d47e88205730710745e78e9d8bcaf8aad8f134604051613ac5919061455e565b60405180910390a2565b613ad76140d8565b73ffffffffffffffffffffffffffffffffffffffff16613af5612d1b565b73ffffffffffffffffffffffffffffffffffffffff1614613b4b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613b4290614f83565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603613bb9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613bb090615b19565b60405180910390fd5b613bc28161412b565b50565b5f426014541015613bd8575f9050613be9565b42601454613be69190614d67565b90505b90565b60015481565b613bfa6140d8565b73ffffffffffffffffffffffffffffffffffffffff16613c18612d1b565b73ffffffffffffffffffffffffffffffffffffffff1614613c6e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613c6590614f83565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603613cdc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613cd390614feb565b60405180910390fd5b8060075f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051613d6c91906144e1565b60405180910390a150565b601b5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601454421015613de1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613dd890615b81565b60405180910390fd5b61011881511115613e27576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613e1e90615be9565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff1660065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603613e905760085442613e899190615009565b6010819055505b613e986140d8565b60065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660165f60175481526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160175f828254613f5b9190615009565b925050819055505f601b5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1960e01b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1668056bc75e2d63100000604051602401613fe09291906152c2565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161404a9190615323565b5f604051808303815f865af19150503d805f8114614083576040519150601f19603f3d011682016040523d82523d5f602084013e614088565b606091505b50509050806140cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016140c390615c77565b60405180910390fd5b6140d46143c1565b5050565b5f33905090565b6015544243406040516020016140f793929190615c95565b60405160208183030381529060405280519060200120601581905550565b600954476141239190615641565b600581905550565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b8173ffffffffffffffffffffffffffffffffffffffff166342842e0e6142106140d8565b30846040518463ffffffff1660e01b8152600401614230939291906159a3565b5f604051808303815f87803b158015614247575f80fd5b505af1158015614259573d5f803e3d5ffd5b5050505060405180608001604052808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200160115481526020015f151581525060195f601a5481526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff0219169083151502179055509050506001601a5f8282546143379190615009565b925050819055506011548273ffffffffffffffffffffffffffffffffffffffff166143606140d8565b73ffffffffffffffffffffffffffffffffffffffff167fc85be1734ed6a0f4d6adf924d4d41406e2729878c652110a96e2fdec64e118d1846001601a546143a79190614d67565b6040516143b5929190615cca565b60405180910390a45050565b5f633b9aca006002546143d49190615641565b9050806143e36010544261441a565b6143ed9190615009565b601081905550620f42406003546002546144079190615600565b6144119190615641565b60028190555050565b5f81831015614429578161442b565b825b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b61445681614444565b8114614460575f80fd5b50565b5f813590506144718161444d565b92915050565b5f6020828403121561448c5761448b61443c565b5b5f61449984828501614463565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6144cb826144a2565b9050919050565b6144db816144c1565b82525050565b5f6020820190506144f45f8301846144d2565b92915050565b614503816144c1565b811461450d575f80fd5b50565b5f8135905061451e816144fa565b92915050565b5f602082840312156145395761453861443c565b5b5f61454684828501614510565b91505092915050565b61455881614444565b82525050565b5f6020820190506145715f83018461454f565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f84011261459857614597614577565b5b8235905067ffffffffffffffff8111156145b5576145b461457b565b5b6020830191508360018202830111156145d1576145d061457f565b5b9250929050565b5f805f805f608086880312156145f1576145f061443c565b5b5f6145fe88828901614510565b955050602061460f88828901614510565b945050604061462088828901614463565b935050606086013567ffffffffffffffff81111561464157614640614440565b5b61464d88828901614583565b92509250509295509295909350565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6146908161465c565b82525050565b5f6020820190506146a95f830184614687565b92915050565b5f819050919050565b5f6146d26146cd6146c8846144a2565b6146af565b6144a2565b9050919050565b5f6146e3826146b8565b9050919050565b5f6146f4826146d9565b9050919050565b614704816146ea565b82525050565b5f60208201905061471d5f8301846146fb565b92915050565b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61476982614723565b810181811067ffffffffffffffff8211171561478857614787614733565b5b80604052505050565b5f61479a614433565b90506147a68282614760565b919050565b5f67ffffffffffffffff8211156147c5576147c4614733565b5b602082029050602081019050919050565b5f6147e86147e3846147ab565b614791565b9050808382526020820190506020840283018581111561480b5761480a61457f565b5b835b8181101561483457806148208882614463565b84526020840193505060208101905061480d565b5050509392505050565b5f82601f83011261485257614851614577565b5b81356148628482602086016147d6565b91505092915050565b5f602082840312156148805761487f61443c565b5b5f82013567ffffffffffffffff81111561489d5761489c614440565b5b6148a98482850161483e565b91505092915050565b5f6148bc826146d9565b9050919050565b6148cc816148b2565b82525050565b5f6020820190506148e55f8301846148c3565b92915050565b5f6148f5826146d9565b9050919050565b614905816148eb565b82525050565b5f60208201905061491e5f8301846148fc565b92915050565b5f80fd5b5f67ffffffffffffffff82111561494257614941614733565b5b61494b82614723565b9050602081019050919050565b828183375f83830152505050565b5f61497861497384614928565b614791565b90508281526020810184848401111561499457614993614924565b5b61499f848285614958565b509392505050565b5f82601f8301126149bb576149ba614577565b5b81356149cb848260208601614966565b91505092915050565b5f602082840312156149e9576149e861443c565b5b5f82013567ffffffffffffffff811115614a0657614a05614440565b5b614a12848285016149a7565b91505092915050565b5f819050919050565b614a2d81614a1b565b82525050565b5f602082019050614a465f830184614a24565b92915050565b5f8115159050919050565b614a6081614a4c565b82525050565b5f602082019050614a795f830184614a57565b92915050565b5f614a89826144c1565b9050919050565b614a9981614a7f565b8114614aa3575f80fd5b50565b5f81359050614ab481614a90565b92915050565b5f805f8060808587031215614ad257614ad161443c565b5b5f614adf87828801614463565b945050602085013567ffffffffffffffff811115614b0057614aff614440565b5b614b0c878288016149a7565b9350506040614b1d87828801614aa6565b9250506060614b2e87828801614463565b91505092959194509250565b5f8060408385031215614b5057614b4f61443c565b5b5f614b5d85828601614463565b925050602083013567ffffffffffffffff811115614b7e57614b7d614440565b5b614b8a858286016149a7565b9150509250929050565b5f805f60608486031215614bab57614baa61443c565b5b5f84013567ffffffffffffffff811115614bc857614bc7614440565b5b614bd4868287016149a7565b9350506020614be586828701614aa6565b9250506040614bf686828701614463565b9150509250925092565b5f614c0a826146d9565b9050919050565b614c1a81614c00565b82525050565b5f608082019050614c335f830187614c11565b614c40602083018661454f565b614c4d604083018561454f565b614c5a6060830184614a57565b95945050505050565b5f614c6d826146d9565b9050919050565b614c7d81614c63565b82525050565b5f602082019050614c965f830184614c74565b92915050565b5f82825260208201905092915050565b7f5468652076616c7565207375626d6974746564207769746820746869732074725f8201527f616e73616374696f6e20697320746f6f206c6f772e0000000000000000000000602082015250565b5f614d06603583614c9c565b9150614d1182614cac565b604082019050919050565b5f6020820190508181035f830152614d3381614cfa565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f614d7182614444565b9150614d7c83614444565b9250828203905081811115614d9457614d93614d3a565b5b92915050565b5f81905092915050565b50565b5f614db25f83614d9a565b9150614dbd82614da4565b5f82019050919050565b5f614dd182614da7565b9150819050919050565b7f526566756e64207472616e73666572206661696c65642e0000000000000000005f82015250565b5f614e0f601783614c9c565b9150614e1a82614ddb565b602082019050919050565b5f6020820190508181035f830152614e3c81614e03565b9050919050565b5f819050919050565b5f819050919050565b5f614e6f614e6a614e6584614e43565b6146af565b614e4c565b9050919050565b614e7f81614e55565b82525050565b5f81519050919050565b5f5b83811015614eac578082015181840152602081019050614e91565b5f8484015250505050565b5f614ec182614e85565b614ecb8185614c9c565b9350614edb818560208601614e8f565b614ee481614723565b840191505092915050565b5f608082019050614f025f83018761454f565b614f0f6020830186614e76565b614f1c604083018561454f565b8181036060830152614f2e8184614eb7565b905095945050505050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f614f6d602083614c9c565b9150614f7882614f39565b602082019050919050565b5f6020820190508181035f830152614f9a81614f61565b9050919050565b7f5a65726f2d616464726573732077617320676976656e2e0000000000000000005f82015250565b5f614fd5601783614c9c565b9150614fe082614fa1565b602082019050919050565b5f6020820190508181035f83015261500281614fc9565b9050919050565b5f61501382614444565b915061501e83614444565b925082820190508082111561503657615035614d3a565b5b92915050565b7f50657263656e746167652076616c7565206f766572666c6f772c206d757374205f8201527f6265206c6f776572207468616e203130302e0000000000000000000000000000602082015250565b5f615096603283614c9c565b91506150a18261503c565b604082019050919050565b5f6020820190508181035f8301526150c38161508a565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f61510182614444565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361513357615132614d3a565b5b600182019050919050565b7f4e6f7420656e6f7567682074696d652068617320656c61707365642e000000005f82015250565b5f615172601c83614c9c565b915061517d8261513e565b602082019050919050565b5f6020820190508181035f83015261519f81615166565b9050919050565b7f5468657265206973206e6f206c617374206269646465722e00000000000000005f82015250565b5f6151da601883614c9c565b91506151e5826151a6565b602082019050919050565b5f6020820190508181035f830152615207816151ce565b9050919050565b7f4f6e6c7920746865206c617374206269646465722063616e20636c61696d20745f8201527f6865207072697a6520647572696e672074686520666972737420323420686f7560208201527f72732e0000000000000000000000000000000000000000000000000000000000604082015250565b5f61528e604383614c9c565b91506152998261520e565b606082019050919050565b5f6020820190508181035f8301526152bb81615282565b9050919050565b5f6040820190506152d55f8301856144d2565b6152e2602083018461454f565b9392505050565b5f81519050919050565b5f6152fd826152e9565b6153078185614d9a565b9350615317818560208601614e8f565b80840191505092915050565b5f61532e82846152f3565b915081905092915050565b7f436f736d69635369676e6174757265206d696e742829206661696c656420746f5f8201527f206d696e74204e46542e00000000000000000000000000000000000000000000602082015250565b5f615393602a83614c9c565b915061539e82615339565b604082019050919050565b5f6020820190508181035f8301526153c081615387565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6153fe82614444565b915061540983614444565b925082615419576154186153c7565b5b828206905092915050565b5f815190506154328161444d565b92915050565b5f6020828403121561544d5761544c61443c565b5b5f61545a84828501615424565b91505092915050565b5f81519050615471816144fa565b92915050565b5f6020828403121561548c5761548b61443c565b5b5f61549984828501615463565b91505092915050565b7f5472616e7366657220746f207468652077696e6e6572206661696c65642e00005f82015250565b5f6154d6601e83614c9c565b91506154e1826154a2565b602082019050919050565b5f6020820190508181035f830152615503816154ca565b9050919050565b7f5472616e7366657220746f206368617269747920636f6e7472616374206661695f8201527f6c65642e00000000000000000000000000000000000000000000000000000000602082015250565b5f615564602483614c9c565b915061556f8261550a565b604082019050919050565b5f6020820190508181035f83015261559181615558565b9050919050565b7f526166666c65206465706f736974206661696c65642e000000000000000000005f82015250565b5f6155cc601683614c9c565b91506155d782615598565b602082019050919050565b5f6020820190508181035f8301526155f9816155c0565b9050919050565b5f61560a82614444565b915061561583614444565b925082820261562381614444565b9150828204841483151761563a57615639614d3a565b5b5092915050565b5f61564b82614444565b915061565683614444565b925082615666576156656153c7565b5b828204905092915050565b7f546869732052616e646f6d57616c6b4e46542068617320616c726561647920625f8201527f65656e207573656420666f722062696464696e672e0000000000000000000000602082015250565b5f6156cb603583614c9c565b91506156d682615671565b604082019050919050565b5f6020820190508181035f8301526156f8816156bf565b9050919050565b7f596f75206d75737420626520746865206f776e6572206f66207468652052616e5f8201527f646f6d57616c6b4e46542e000000000000000000000000000000000000000000602082015250565b5f615759602b83614c9c565b9150615764826156ff565b604082019050919050565b5f6020820190508181035f8301526157868161574d565b9050919050565b5f819050919050565b5f6157b06157ab6157a68461578d565b6146af565b614444565b9050919050565b6157c081615796565b82525050565b6157cf81614e4c565b82525050565b5f6080820190506157e85f8301876157b7565b6157f560208301866157c6565b615802604083018561454f565b81810360608301526158148184614eb7565b905095945050505050565b7f54686520646f6e61746564204e465420646f6573206e6f742065786973742e005f82015250565b5f615853601f83614c9c565b915061585e8261581f565b602082019050919050565b5f6020820190508181035f83015261588081615847565b9050919050565b7f596f7520617265206e6f74207468652077696e6e6572206f662074686520726f5f8201527f756e642e00000000000000000000000000000000000000000000000000000000602082015250565b5f6158e1602483614c9c565b91506158ec82615887565b604082019050919050565b5f6020820190508181035f83015261590e816158d5565b9050919050565b7f546865204e46542068617320616c7265616479206265656e20636c61696d65645f8201527f2e00000000000000000000000000000000000000000000000000000000000000602082015250565b5f61596f602183614c9c565b915061597a82615915565b604082019050919050565b5f6020820190508181035f83015261599c81615963565b9050919050565b5f6060820190506159b65f8301866144d2565b6159c360208301856144d2565b6159d0604083018461454f565b949350505050565b5f6080820190506159eb5f83018761454f565b6159f860208301866144d2565b615a0560408301856144d2565b615a12606083018461454f565b95945050505050565b7f446f6e6174696f6e20616d6f756e74206d7573742062652067726561746572205f8201527f7468616e20302e00000000000000000000000000000000000000000000000000602082015250565b5f615a75602783614c9c565b9150615a8082615a1b565b604082019050919050565b5f6020820190508181035f830152615aa281615a69565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f615b03602683614c9c565b9150615b0e82615aa9565b604082019050919050565b5f6020820190508181035f830152615b3081615af7565b9050919050565b7f4e6f7420616374697665207965742e00000000000000000000000000000000005f82015250565b5f615b6b600f83614c9c565b9150615b7682615b37565b602082019050919050565b5f6020820190508181035f830152615b9881615b5f565b9050919050565b7f4d65737361676520697320746f6f206c6f6e672e0000000000000000000000005f82015250565b5f615bd3601483614c9c565b9150615bde82615b9f565b602082019050919050565b5f6020820190508181035f830152615c0081615bc7565b9050919050565b7f436f736d6963546f6b656e206d696e742829206661696c656420746f206d696e5f8201527f742072657761726420746f6b656e732e00000000000000000000000000000000602082015250565b5f615c61603083614c9c565b9150615c6c82615c07565b604082019050919050565b5f6020820190508181035f830152615c8e81615c55565b9050919050565b5f606082019050615ca85f830186614a24565b615cb5602083018561454f565b615cc26040830184614a24565b949350505050565b5f604082019050615cdd5f83018561454f565b615cea602083018461454f565b939250505056fea26469706673582212201d708a3b80bad43190e2d80ed747856a45922aa58c40f1e439c65ec1712c010664736f6c63430008150033",
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

// TOKENREWARD is a free data retrieval call binding the contract method 0xd4ba8902.
//
// Solidity: function TOKEN_REWARD() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) TOKENREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "TOKEN_REWARD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOKENREWARD is a free data retrieval call binding the contract method 0xd4ba8902.
//
// Solidity: function TOKEN_REWARD() view returns(uint256)
func (_CosmicGame *CosmicGameSession) TOKENREWARD() (*big.Int, error) {
	return _CosmicGame.Contract.TOKENREWARD(&_CosmicGame.CallOpts)
}

// TOKENREWARD is a free data retrieval call binding the contract method 0xd4ba8902.
//
// Solidity: function TOKEN_REWARD() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) TOKENREWARD() (*big.Int, error) {
	return _CosmicGame.Contract.TOKENREWARD(&_CosmicGame.CallOpts)
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

// NumHolderNFTWinnersPerRound is a free data retrieval call binding the contract method 0xdec08d8e.
//
// Solidity: function numHolderNFTWinnersPerRound() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) NumHolderNFTWinnersPerRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "numHolderNFTWinnersPerRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumHolderNFTWinnersPerRound is a free data retrieval call binding the contract method 0xdec08d8e.
//
// Solidity: function numHolderNFTWinnersPerRound() view returns(uint256)
func (_CosmicGame *CosmicGameSession) NumHolderNFTWinnersPerRound() (*big.Int, error) {
	return _CosmicGame.Contract.NumHolderNFTWinnersPerRound(&_CosmicGame.CallOpts)
}

// NumHolderNFTWinnersPerRound is a free data retrieval call binding the contract method 0xdec08d8e.
//
// Solidity: function numHolderNFTWinnersPerRound() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) NumHolderNFTWinnersPerRound() (*big.Int, error) {
	return _CosmicGame.Contract.NumHolderNFTWinnersPerRound(&_CosmicGame.CallOpts)
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

// TimeoutClaimPrize is a free data retrieval call binding the contract method 0x355f01e2.
//
// Solidity: function timeoutClaimPrize() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) TimeoutClaimPrize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "timeoutClaimPrize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeoutClaimPrize is a free data retrieval call binding the contract method 0x355f01e2.
//
// Solidity: function timeoutClaimPrize() view returns(uint256)
func (_CosmicGame *CosmicGameSession) TimeoutClaimPrize() (*big.Int, error) {
	return _CosmicGame.Contract.TimeoutClaimPrize(&_CosmicGame.CallOpts)
}

// TimeoutClaimPrize is a free data retrieval call binding the contract method 0x355f01e2.
//
// Solidity: function timeoutClaimPrize() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) TimeoutClaimPrize() (*big.Int, error) {
	return _CosmicGame.Contract.TimeoutClaimPrize(&_CosmicGame.CallOpts)
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

// ClaimManyDonatedNFTs is a paid mutator transaction binding the contract method 0x3f7909d4.
//
// Solidity: function claimManyDonatedNFTs(uint256[] tokens) returns()
func (_CosmicGame *CosmicGameTransactor) ClaimManyDonatedNFTs(opts *bind.TransactOpts, tokens []*big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "claimManyDonatedNFTs", tokens)
}

// ClaimManyDonatedNFTs is a paid mutator transaction binding the contract method 0x3f7909d4.
//
// Solidity: function claimManyDonatedNFTs(uint256[] tokens) returns()
func (_CosmicGame *CosmicGameSession) ClaimManyDonatedNFTs(tokens []*big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.ClaimManyDonatedNFTs(&_CosmicGame.TransactOpts, tokens)
}

// ClaimManyDonatedNFTs is a paid mutator transaction binding the contract method 0x3f7909d4.
//
// Solidity: function claimManyDonatedNFTs(uint256[] tokens) returns()
func (_CosmicGame *CosmicGameTransactorSession) ClaimManyDonatedNFTs(tokens []*big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.ClaimManyDonatedNFTs(&_CosmicGame.TransactOpts, tokens)
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

// SetNumHolderNFTWinnersPerRound is a paid mutator transaction binding the contract method 0xc01c5de2.
//
// Solidity: function setNumHolderNFTWinnersPerRound(uint256 newNumHolderNFTWinnersPerRound) returns()
func (_CosmicGame *CosmicGameTransactor) SetNumHolderNFTWinnersPerRound(opts *bind.TransactOpts, newNumHolderNFTWinnersPerRound *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setNumHolderNFTWinnersPerRound", newNumHolderNFTWinnersPerRound)
}

// SetNumHolderNFTWinnersPerRound is a paid mutator transaction binding the contract method 0xc01c5de2.
//
// Solidity: function setNumHolderNFTWinnersPerRound(uint256 newNumHolderNFTWinnersPerRound) returns()
func (_CosmicGame *CosmicGameSession) SetNumHolderNFTWinnersPerRound(newNumHolderNFTWinnersPerRound *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNumHolderNFTWinnersPerRound(&_CosmicGame.TransactOpts, newNumHolderNFTWinnersPerRound)
}

// SetNumHolderNFTWinnersPerRound is a paid mutator transaction binding the contract method 0xc01c5de2.
//
// Solidity: function setNumHolderNFTWinnersPerRound(uint256 newNumHolderNFTWinnersPerRound) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetNumHolderNFTWinnersPerRound(newNumHolderNFTWinnersPerRound *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetNumHolderNFTWinnersPerRound(&_CosmicGame.TransactOpts, newNumHolderNFTWinnersPerRound)
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

// SetTimeoutClaimPrize is a paid mutator transaction binding the contract method 0x019d354a.
//
// Solidity: function setTimeoutClaimPrize(uint256 newTimeout) returns()
func (_CosmicGame *CosmicGameTransactor) SetTimeoutClaimPrize(opts *bind.TransactOpts, newTimeout *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setTimeoutClaimPrize", newTimeout)
}

// SetTimeoutClaimPrize is a paid mutator transaction binding the contract method 0x019d354a.
//
// Solidity: function setTimeoutClaimPrize(uint256 newTimeout) returns()
func (_CosmicGame *CosmicGameSession) SetTimeoutClaimPrize(newTimeout *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetTimeoutClaimPrize(&_CosmicGame.TransactOpts, newTimeout)
}

// SetTimeoutClaimPrize is a paid mutator transaction binding the contract method 0x019d354a.
//
// Solidity: function setTimeoutClaimPrize(uint256 newTimeout) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetTimeoutClaimPrize(newTimeout *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetTimeoutClaimPrize(&_CosmicGame.TransactOpts, newTimeout)
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

// CosmicGameNumHolderNFTWinnersPerRoundChangedIterator is returned from FilterNumHolderNFTWinnersPerRoundChanged and is used to iterate over the raw logs and unpacked data for NumHolderNFTWinnersPerRoundChanged events raised by the CosmicGame contract.
type CosmicGameNumHolderNFTWinnersPerRoundChangedIterator struct {
	Event *CosmicGameNumHolderNFTWinnersPerRoundChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameNumHolderNFTWinnersPerRoundChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameNumHolderNFTWinnersPerRoundChanged)
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
		it.Event = new(CosmicGameNumHolderNFTWinnersPerRoundChanged)
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
func (it *CosmicGameNumHolderNFTWinnersPerRoundChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameNumHolderNFTWinnersPerRoundChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameNumHolderNFTWinnersPerRoundChanged represents a NumHolderNFTWinnersPerRoundChanged event raised by the CosmicGame contract.
type CosmicGameNumHolderNFTWinnersPerRoundChanged struct {
	NewNumHolderNFTWinnersPerRound *big.Int
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterNumHolderNFTWinnersPerRoundChanged is a free log retrieval operation binding the contract event 0x0cc7e37c68566f67d0fe13bf38246d7447cf99a0c481c2ef9963969bb4f5ebc8.
//
// Solidity: event NumHolderNFTWinnersPerRoundChanged(uint256 newNumHolderNFTWinnersPerRound)
func (_CosmicGame *CosmicGameFilterer) FilterNumHolderNFTWinnersPerRoundChanged(opts *bind.FilterOpts) (*CosmicGameNumHolderNFTWinnersPerRoundChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "NumHolderNFTWinnersPerRoundChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameNumHolderNFTWinnersPerRoundChangedIterator{contract: _CosmicGame.contract, event: "NumHolderNFTWinnersPerRoundChanged", logs: logs, sub: sub}, nil
}

// WatchNumHolderNFTWinnersPerRoundChanged is a free log subscription operation binding the contract event 0x0cc7e37c68566f67d0fe13bf38246d7447cf99a0c481c2ef9963969bb4f5ebc8.
//
// Solidity: event NumHolderNFTWinnersPerRoundChanged(uint256 newNumHolderNFTWinnersPerRound)
func (_CosmicGame *CosmicGameFilterer) WatchNumHolderNFTWinnersPerRoundChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameNumHolderNFTWinnersPerRoundChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "NumHolderNFTWinnersPerRoundChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameNumHolderNFTWinnersPerRoundChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "NumHolderNFTWinnersPerRoundChanged", log); err != nil {
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

// ParseNumHolderNFTWinnersPerRoundChanged is a log parse operation binding the contract event 0x0cc7e37c68566f67d0fe13bf38246d7447cf99a0c481c2ef9963969bb4f5ebc8.
//
// Solidity: event NumHolderNFTWinnersPerRoundChanged(uint256 newNumHolderNFTWinnersPerRound)
func (_CosmicGame *CosmicGameFilterer) ParseNumHolderNFTWinnersPerRoundChanged(log types.Log) (*CosmicGameNumHolderNFTWinnersPerRoundChanged, error) {
	event := new(CosmicGameNumHolderNFTWinnersPerRoundChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "NumHolderNFTWinnersPerRoundChanged", log); err != nil {
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
	TokenId     *big.Int
	WinnerIndex *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRaffleNFTWinnerEvent is a free log retrieval operation binding the contract event 0x0afb73a48ffcc4cf27ecb1889154b910544b89c39c5e5822c431062629134e56.
//
// Solidity: event RaffleNFTWinnerEvent(address indexed winner, uint256 indexed round, uint256 indexed tokenId, uint256 winnerIndex)
func (_CosmicGame *CosmicGameFilterer) FilterRaffleNFTWinnerEvent(opts *bind.FilterOpts, winner []common.Address, round []*big.Int, tokenId []*big.Int) (*CosmicGameRaffleNFTWinnerEventIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "RaffleNFTWinnerEvent", winnerRule, roundRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CosmicGameRaffleNFTWinnerEventIterator{contract: _CosmicGame.contract, event: "RaffleNFTWinnerEvent", logs: logs, sub: sub}, nil
}

// WatchRaffleNFTWinnerEvent is a free log subscription operation binding the contract event 0x0afb73a48ffcc4cf27ecb1889154b910544b89c39c5e5822c431062629134e56.
//
// Solidity: event RaffleNFTWinnerEvent(address indexed winner, uint256 indexed round, uint256 indexed tokenId, uint256 winnerIndex)
func (_CosmicGame *CosmicGameFilterer) WatchRaffleNFTWinnerEvent(opts *bind.WatchOpts, sink chan<- *CosmicGameRaffleNFTWinnerEvent, winner []common.Address, round []*big.Int, tokenId []*big.Int) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "RaffleNFTWinnerEvent", winnerRule, roundRule, tokenIdRule)
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

// ParseRaffleNFTWinnerEvent is a log parse operation binding the contract event 0x0afb73a48ffcc4cf27ecb1889154b910544b89c39c5e5822c431062629134e56.
//
// Solidity: event RaffleNFTWinnerEvent(address indexed winner, uint256 indexed round, uint256 indexed tokenId, uint256 winnerIndex)
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

// CosmicGameTimeoutClaimPrizeChangedIterator is returned from FilterTimeoutClaimPrizeChanged and is used to iterate over the raw logs and unpacked data for TimeoutClaimPrizeChanged events raised by the CosmicGame contract.
type CosmicGameTimeoutClaimPrizeChangedIterator struct {
	Event *CosmicGameTimeoutClaimPrizeChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameTimeoutClaimPrizeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameTimeoutClaimPrizeChanged)
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
		it.Event = new(CosmicGameTimeoutClaimPrizeChanged)
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
func (it *CosmicGameTimeoutClaimPrizeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameTimeoutClaimPrizeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameTimeoutClaimPrizeChanged represents a TimeoutClaimPrizeChanged event raised by the CosmicGame contract.
type CosmicGameTimeoutClaimPrizeChanged struct {
	NewTimeout *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTimeoutClaimPrizeChanged is a free log retrieval operation binding the contract event 0xcaf8e5c6bc1bb2c19935f84ddcdaefb232ad06f9f2abd2ad588bea4bbe631d26.
//
// Solidity: event TimeoutClaimPrizeChanged(uint256 newTimeout)
func (_CosmicGame *CosmicGameFilterer) FilterTimeoutClaimPrizeChanged(opts *bind.FilterOpts) (*CosmicGameTimeoutClaimPrizeChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "TimeoutClaimPrizeChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameTimeoutClaimPrizeChangedIterator{contract: _CosmicGame.contract, event: "TimeoutClaimPrizeChanged", logs: logs, sub: sub}, nil
}

// WatchTimeoutClaimPrizeChanged is a free log subscription operation binding the contract event 0xcaf8e5c6bc1bb2c19935f84ddcdaefb232ad06f9f2abd2ad588bea4bbe631d26.
//
// Solidity: event TimeoutClaimPrizeChanged(uint256 newTimeout)
func (_CosmicGame *CosmicGameFilterer) WatchTimeoutClaimPrizeChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameTimeoutClaimPrizeChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "TimeoutClaimPrizeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameTimeoutClaimPrizeChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "TimeoutClaimPrizeChanged", log); err != nil {
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

// ParseTimeoutClaimPrizeChanged is a log parse operation binding the contract event 0xcaf8e5c6bc1bb2c19935f84ddcdaefb232ad06f9f2abd2ad588bea4bbe631d26.
//
// Solidity: event TimeoutClaimPrizeChanged(uint256 newTimeout)
func (_CosmicGame *CosmicGameFilterer) ParseTimeoutClaimPrizeChanged(log types.Log) (*CosmicGameTimeoutClaimPrizeChanged, error) {
	event := new(CosmicGameTimeoutClaimPrizeChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "TimeoutClaimPrizeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CountersMetaData contains all meta data concerning the Counters contract.
var CountersMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea2646970667358221220fd162b5463951badd050a822a25f5943261cf66d524750dca2748307d93db50c64736f6c63430008150033",
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
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea26469706673582212209709edcf2338d17fe757e98ff438364592856ea5dfe7cd2ae4ba3766b4671c1864736f6c63430008150033",
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

// ERC721MetaData contains all meta data concerning the ERC721 contract.
var ERC721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b506040516200284f3803806200284f8339818101604052810190620000369190620001e6565b815f9081620000469190620004a0565b508060019081620000589190620004a0565b50505062000584565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b620000c2826200007a565b810181811067ffffffffffffffff82111715620000e457620000e36200008a565b5b80604052505050565b5f620000f862000061565b9050620001068282620000b7565b919050565b5f67ffffffffffffffff8211156200012857620001276200008a565b5b62000133826200007a565b9050602081019050919050565b5f5b838110156200015f57808201518184015260208101905062000142565b5f8484015250505050565b5f620001806200017a846200010b565b620000ed565b9050828152602081018484840111156200019f576200019e62000076565b5b620001ac84828562000140565b509392505050565b5f82601f830112620001cb57620001ca62000072565b5b8151620001dd8482602086016200016a565b91505092915050565b5f8060408385031215620001ff57620001fe6200006a565b5b5f83015167ffffffffffffffff8111156200021f576200021e6200006e565b5b6200022d85828601620001b4565b925050602083015167ffffffffffffffff8111156200025157620002506200006e565b5b6200025f85828601620001b4565b9150509250929050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680620002b857607f821691505b602082108103620002ce57620002cd62000273565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620003327fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620002f5565b6200033e8683620002f5565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f62000388620003826200037c8462000356565b6200035f565b62000356565b9050919050565b5f819050919050565b620003a38362000368565b620003bb620003b2826200038f565b84845462000301565b825550505050565b5f90565b620003d1620003c3565b620003de81848462000398565b505050565b5b818110156200040557620003f95f82620003c7565b600181019050620003e4565b5050565b601f82111562000454576200041e81620002d4565b6200042984620002e6565b8101602085101562000439578190505b620004516200044885620002e6565b830182620003e3565b50505b505050565b5f82821c905092915050565b5f620004765f198460080262000459565b1980831691505092915050565b5f62000490838362000465565b9150826002028217905092915050565b620004ab8262000269565b67ffffffffffffffff811115620004c757620004c66200008a565b5b620004d38254620002a0565b620004e082828562000409565b5f60209050601f83116001811462000516575f841562000501578287015190505b6200050d858262000483565b8655506200057c565b601f1984166200052686620002d4565b5f5b828110156200054f5784890151825560018201915060208501945060208101905062000528565b868310156200056f57848901516200056b601f89168262000465565b8355505b6001600288020188555050505b505050505050565b6122bd80620005925f395ff3fe608060405234801561000f575f80fd5b50600436106100cd575f3560e01c80636352211e1161008a578063a22cb46511610064578063a22cb46514610221578063b88d4fde1461023d578063c87b56dd14610259578063e985e9c514610289576100cd565b80636352211e146101a357806370a08231146101d357806395d89b4114610203576100cd565b806301ffc9a7146100d157806306fdde0314610101578063081812fc1461011f578063095ea7b31461014f57806323b872dd1461016b57806342842e0e14610187575b5f80fd5b6100eb60048036038101906100e691906113c1565b6102b9565b6040516100f89190611406565b60405180910390f35b61010961039a565b60405161011691906114a9565b60405180910390f35b610139600480360381019061013491906114fc565b610429565b6040516101469190611566565b60405180910390f35b610169600480360381019061016491906115a9565b6104aa565b005b610185600480360381019061018091906115e7565b6105c0565b005b6101a1600480360381019061019c91906115e7565b610620565b005b6101bd60048036038101906101b891906114fc565b61063f565b6040516101ca9190611566565b60405180910390f35b6101ed60048036038101906101e89190611637565b6106eb565b6040516101fa9190611671565b60405180910390f35b61020b61079f565b60405161021891906114a9565b60405180910390f35b61023b600480360381019061023691906116b4565b61082f565b005b6102576004803603810190610252919061181e565b6109aa565b005b610273600480360381019061026e91906114fc565b610a0c565b60405161028091906114a9565b60405180910390f35b6102a3600480360381019061029e919061189e565b610ab0565b6040516102b09190611406565b60405180910390f35b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061038357507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610393575061039282610b3e565b5b9050919050565b60605f80546103a890611909565b80601f01602080910402602001604051908101604052809291908181526020018280546103d490611909565b801561041f5780601f106103f65761010080835404028352916020019161041f565b820191905f5260205f20905b81548152906001019060200180831161040257829003601f168201915b5050505050905090565b5f61043382610ba7565b610472576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610469906119a9565b60405180910390fd5b60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f6104b48261063f565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610524576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161051b90611a37565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16610543610c0f565b73ffffffffffffffffffffffffffffffffffffffff16148061057257506105718161056c610c0f565b610ab0565b5b6105b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a890611ac5565b60405180910390fd5b6105bb8383610c16565b505050565b6105d16105cb610c0f565b82610ccc565b610610576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060790611b53565b60405180910390fd5b61061b838383610da8565b505050565b61063a83838360405180602001604052805f8152506109aa565b505050565b5f8060025f8481526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036106e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d990611be1565b60405180910390fd5b80915050919050565b5f8073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361075a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075190611c6f565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6060600180546107ae90611909565b80601f01602080910402602001604051908101604052809291908181526020018280546107da90611909565b80156108255780601f106107fc57610100808354040283529160200191610825565b820191905f5260205f20905b81548152906001019060200180831161080857829003601f168201915b5050505050905090565b610837610c0f565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089b90611cd7565b60405180910390fd5b8060055f6108b0610c0f565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff16610959610c0f565b73ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c318360405161099e9190611406565b60405180910390a35050565b6109bb6109b5610c0f565b83610ccc565b6109fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109f190611b53565b60405180910390fd5b610a0684848484610ff8565b50505050565b6060610a1782610ba7565b610a56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4d90611d65565b60405180910390fd5b5f610a5f611054565b90505f815111610a7d5760405180602001604052805f815250610aa8565b80610a878461106a565b604051602001610a98929190611dbd565b6040516020818303038152906040525b915050919050565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f8073ffffffffffffffffffffffffffffffffffffffff1660025f8481526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b5f33905090565b8160045f8381526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff16610c868361063f565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b5f610cd682610ba7565b610d15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d0c90611e50565b60405180910390fd5b5f610d1f8361063f565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480610d8e57508373ffffffffffffffffffffffffffffffffffffffff16610d7684610429565b73ffffffffffffffffffffffffffffffffffffffff16145b80610d9f5750610d9e8185610ab0565b5b91505092915050565b8273ffffffffffffffffffffffffffffffffffffffff16610dc88261063f565b73ffffffffffffffffffffffffffffffffffffffff1614610e1e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e1590611ede565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610e8c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e8390611f6c565b60405180910390fd5b610e978383836111c3565b610ea15f82610c16565b600160035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610eee9190611fb7565b92505081905550600160035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610f429190611fea565b925050819055508160025f8381526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b611003848484610da8565b61100f848484846111c8565b61104e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110459061208d565b60405180910390fd5b50505050565b606060405180602001604052805f815250905090565b60605f82036110b0576040518060400160405280600181526020017f300000000000000000000000000000000000000000000000000000000000000081525090506111be565b5f8290505f5b5f82146110df5780806110c8906120ab565b915050600a826110d8919061211f565b91506110b6565b5f8167ffffffffffffffff8111156110fa576110f96116fa565b5b6040519080825280601f01601f19166020018201604052801561112c5781602001600182028036833780820191505090505b5090505b5f85146111b7576001826111449190611fb7565b9150600a85611153919061214f565b603061115f9190611fea565b60f81b8183815181106111755761117461217f565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600a856111b0919061211f565b9450611130565b8093505050505b919050565b505050565b5f6111e88473ffffffffffffffffffffffffffffffffffffffff1661134a565b1561133d578373ffffffffffffffffffffffffffffffffffffffff1663150b7a02611211610c0f565b8786866040518563ffffffff1660e01b815260040161123394939291906121fe565b6020604051808303815f875af192505050801561126e57506040513d601f19601f8201168201806040525081019061126b919061225c565b60015b6112ed573d805f811461129c576040519150601f19603f3d011682016040523d82523d5f602084013e6112a1565b606091505b505f8151036112e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112dc9061208d565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614915050611342565b600190505b949350505050565b5f80823b90505f8111915050919050565b5f604051905090565b5f80fd5b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6113a08161136c565b81146113aa575f80fd5b50565b5f813590506113bb81611397565b92915050565b5f602082840312156113d6576113d5611364565b5b5f6113e3848285016113ad565b91505092915050565b5f8115159050919050565b611400816113ec565b82525050565b5f6020820190506114195f8301846113f7565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561145657808201518184015260208101905061143b565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61147b8261141f565b6114858185611429565b9350611495818560208601611439565b61149e81611461565b840191505092915050565b5f6020820190508181035f8301526114c18184611471565b905092915050565b5f819050919050565b6114db816114c9565b81146114e5575f80fd5b50565b5f813590506114f6816114d2565b92915050565b5f6020828403121561151157611510611364565b5b5f61151e848285016114e8565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61155082611527565b9050919050565b61156081611546565b82525050565b5f6020820190506115795f830184611557565b92915050565b61158881611546565b8114611592575f80fd5b50565b5f813590506115a38161157f565b92915050565b5f80604083850312156115bf576115be611364565b5b5f6115cc85828601611595565b92505060206115dd858286016114e8565b9150509250929050565b5f805f606084860312156115fe576115fd611364565b5b5f61160b86828701611595565b935050602061161c86828701611595565b925050604061162d868287016114e8565b9150509250925092565b5f6020828403121561164c5761164b611364565b5b5f61165984828501611595565b91505092915050565b61166b816114c9565b82525050565b5f6020820190506116845f830184611662565b92915050565b611693816113ec565b811461169d575f80fd5b50565b5f813590506116ae8161168a565b92915050565b5f80604083850312156116ca576116c9611364565b5b5f6116d785828601611595565b92505060206116e8858286016116a0565b9150509250929050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61173082611461565b810181811067ffffffffffffffff8211171561174f5761174e6116fa565b5b80604052505050565b5f61176161135b565b905061176d8282611727565b919050565b5f67ffffffffffffffff82111561178c5761178b6116fa565b5b61179582611461565b9050602081019050919050565b828183375f83830152505050565b5f6117c26117bd84611772565b611758565b9050828152602081018484840111156117de576117dd6116f6565b5b6117e98482856117a2565b509392505050565b5f82601f830112611805576118046116f2565b5b81356118158482602086016117b0565b91505092915050565b5f805f806080858703121561183657611835611364565b5b5f61184387828801611595565b945050602061185487828801611595565b9350506040611865878288016114e8565b925050606085013567ffffffffffffffff81111561188657611885611368565b5b611892878288016117f1565b91505092959194509250565b5f80604083850312156118b4576118b3611364565b5b5f6118c185828601611595565b92505060206118d285828601611595565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061192057607f821691505b602082108103611933576119326118dc565b5b50919050565b7f4552433732313a20617070726f76656420717565727920666f72206e6f6e65785f8201527f697374656e7420746f6b656e0000000000000000000000000000000000000000602082015250565b5f611993602c83611429565b915061199e82611939565b604082019050919050565b5f6020820190508181035f8301526119c081611987565b9050919050565b7f4552433732313a20617070726f76616c20746f2063757272656e74206f776e655f8201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b5f611a21602183611429565b9150611a2c826119c7565b604082019050919050565b5f6020820190508181035f830152611a4e81611a15565b9050919050565b7f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f775f8201527f6e6572206e6f7220617070726f76656420666f7220616c6c0000000000000000602082015250565b5f611aaf603883611429565b9150611aba82611a55565b604082019050919050565b5f6020820190508181035f830152611adc81611aa3565b9050919050565b7f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f5f8201527f776e6572206e6f7220617070726f766564000000000000000000000000000000602082015250565b5f611b3d603183611429565b9150611b4882611ae3565b604082019050919050565b5f6020820190508181035f830152611b6a81611b31565b9050919050565b7f4552433732313a206f776e657220717565727920666f72206e6f6e65786973745f8201527f656e7420746f6b656e0000000000000000000000000000000000000000000000602082015250565b5f611bcb602983611429565b9150611bd682611b71565b604082019050919050565b5f6020820190508181035f830152611bf881611bbf565b9050919050565b7f4552433732313a2062616c616e636520717565727920666f7220746865207a655f8201527f726f206164647265737300000000000000000000000000000000000000000000602082015250565b5f611c59602a83611429565b9150611c6482611bff565b604082019050919050565b5f6020820190508181035f830152611c8681611c4d565b9050919050565b7f4552433732313a20617070726f766520746f2063616c6c6572000000000000005f82015250565b5f611cc1601983611429565b9150611ccc82611c8d565b602082019050919050565b5f6020820190508181035f830152611cee81611cb5565b9050919050565b7f4552433732314d657461646174613a2055524920717565727920666f72206e6f5f8201527f6e6578697374656e7420746f6b656e0000000000000000000000000000000000602082015250565b5f611d4f602f83611429565b9150611d5a82611cf5565b604082019050919050565b5f6020820190508181035f830152611d7c81611d43565b9050919050565b5f81905092915050565b5f611d978261141f565b611da18185611d83565b9350611db1818560208601611439565b80840191505092915050565b5f611dc88285611d8d565b9150611dd48284611d8d565b91508190509392505050565b7f4552433732313a206f70657261746f7220717565727920666f72206e6f6e65785f8201527f697374656e7420746f6b656e0000000000000000000000000000000000000000602082015250565b5f611e3a602c83611429565b9150611e4582611de0565b604082019050919050565b5f6020820190508181035f830152611e6781611e2e565b9050919050565b7f4552433732313a207472616e73666572206f6620746f6b656e207468617420695f8201527f73206e6f74206f776e0000000000000000000000000000000000000000000000602082015250565b5f611ec8602983611429565b9150611ed382611e6e565b604082019050919050565b5f6020820190508181035f830152611ef581611ebc565b9050919050565b7f4552433732313a207472616e7366657220746f20746865207a65726f206164645f8201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b5f611f56602483611429565b9150611f6182611efc565b604082019050919050565b5f6020820190508181035f830152611f8381611f4a565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611fc1826114c9565b9150611fcc836114c9565b9250828203905081811115611fe457611fe3611f8a565b5b92915050565b5f611ff4826114c9565b9150611fff836114c9565b925082820190508082111561201757612016611f8a565b5b92915050565b7f4552433732313a207472616e7366657220746f206e6f6e2045524337323152655f8201527f63656976657220696d706c656d656e7465720000000000000000000000000000602082015250565b5f612077603283611429565b91506120828261201d565b604082019050919050565b5f6020820190508181035f8301526120a48161206b565b9050919050565b5f6120b5826114c9565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036120e7576120e6611f8a565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f612129826114c9565b9150612134836114c9565b925082612144576121436120f2565b5b828204905092915050565b5f612159826114c9565b9150612164836114c9565b925082612174576121736120f2565b5b828206905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81519050919050565b5f82825260208201905092915050565b5f6121d0826121ac565b6121da81856121b6565b93506121ea818560208601611439565b6121f381611461565b840191505092915050565b5f6080820190506122115f830187611557565b61221e6020830186611557565b61222b6040830185611662565b818103606083015261223d81846121c6565b905095945050505050565b5f8151905061225681611397565b92915050565b5f6020828403121561227157612270611364565b5b5f61227e84828501612248565b9150509291505056fea2646970667358221220c61ac346fd9b6e70d1d74d02080f18e9789a5ab84830259ef6a50b2a94f84aa664736f6c63430008150033",
}

// ERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721MetaData.ABI instead.
var ERC721ABI = ERC721MetaData.ABI

// ERC721Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC721MetaData.Bin instead.
var ERC721Bin = ERC721MetaData.Bin

// DeployERC721 deploys a new Ethereum contract, binding an instance of ERC721 to it.
func DeployERC721(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *ERC721, error) {
	parsed, err := ERC721MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC721Bin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// ERC721 is an auto generated Go binding around an Ethereum contract.
type ERC721 struct {
	ERC721Caller     // Read-only binding to the contract
	ERC721Transactor // Write-only binding to the contract
	ERC721Filterer   // Log filterer for contract events
}

// ERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721Session struct {
	Contract     *ERC721           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721CallerSession struct {
	Contract *ERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721TransactorSession struct {
	Contract     *ERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721Raw struct {
	Contract *ERC721 // Generic contract binding to access the raw methods on
}

// ERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721CallerRaw struct {
	Contract *ERC721Caller // Generic read-only contract binding to access the raw methods on
}

// ERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721TransactorRaw struct {
	Contract *ERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721 creates a new instance of ERC721, bound to a specific deployed contract.
func NewERC721(address common.Address, backend bind.ContractBackend) (*ERC721, error) {
	contract, err := bindERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// NewERC721Caller creates a new read-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Caller(address common.Address, caller bind.ContractCaller) (*ERC721Caller, error) {
	contract, err := bindERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Caller{contract: contract}, nil
}

// NewERC721Transactor creates a new write-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC721Transactor, error) {
	contract, err := bindERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Transactor{contract: contract}, nil
}

// NewERC721Filterer creates a new log filterer instance of ERC721, bound to a specific deployed contract.
func NewERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC721Filterer, error) {
	contract, err := bindERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721Filterer{contract: contract}, nil
}

// bindERC721 binds a generic wrapper to an already deployed contract.
func bindERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.ERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721 *ERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721 *ERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721 *ERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721 *ERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721 *ERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721 *ERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721 *ERC721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721 *ERC721Session) Name() (string, error) {
	return _ERC721.Contract.Name(&_ERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721 *ERC721CallerSession) Name() (string, error) {
	return _ERC721.Contract.Name(&_ERC721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721 *ERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721 *ERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721 *ERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721 *ERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721 *ERC721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721 *ERC721Session) Symbol() (string, error) {
	return _ERC721.Contract.Symbol(&_ERC721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721 *ERC721CallerSession) Symbol() (string, error) {
	return _ERC721.Contract.Symbol(&_ERC721.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721 *ERC721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721 *ERC721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721.Contract.TokenURI(&_ERC721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721 *ERC721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721.Contract.TokenURI(&_ERC721.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom0(&_ERC721.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom0(&_ERC721.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721 *ERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721 *ERC721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721 *ERC721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

// ERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721 contract.
type ERC721ApprovalIterator struct {
	Event *ERC721Approval // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Approval)
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
		it.Event = new(ERC721Approval)
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
func (it *ERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Approval represents a Approval event raised by the ERC721 contract.
type ERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721ApprovalIterator, error) {

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

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalIterator{contract: _ERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Approval)
				if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC721 *ERC721Filterer) ParseApproval(log types.Log) (*ERC721Approval, error) {
	event := new(ERC721Approval)
	if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721 contract.
type ERC721ApprovalForAllIterator struct {
	Event *ERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721ApprovalForAll)
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
		it.Event = new(ERC721ApprovalForAll)
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
func (it *ERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721ApprovalForAll represents a ApprovalForAll event raised by the ERC721 contract.
type ERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalForAllIterator{contract: _ERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721ApprovalForAll)
				if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_ERC721 *ERC721Filterer) ParseApprovalForAll(log types.Log) (*ERC721ApprovalForAll, error) {
	event := new(ERC721ApprovalForAll)
	if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721 contract.
type ERC721TransferIterator struct {
	Event *ERC721Transfer // Event containing the contract specifics and raw log

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
func (it *ERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Transfer)
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
		it.Event = new(ERC721Transfer)
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
func (it *ERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TransferIterator) Close() error {
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

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721TransferIterator, error) {

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

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TransferIterator{contract: _ERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Transfer)
				if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC721 *ERC721Filterer) ParseTransfer(log types.Log) (*ERC721Transfer, error) {
	event := new(ERC721Transfer)
	if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721EnumerableMetaData contains all meta data concerning the ERC721Enumerable contract.
var ERC721EnumerableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC721EnumerableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721EnumerableMetaData.ABI instead.
var ERC721EnumerableABI = ERC721EnumerableMetaData.ABI

// ERC721Enumerable is an auto generated Go binding around an Ethereum contract.
type ERC721Enumerable struct {
	ERC721EnumerableCaller     // Read-only binding to the contract
	ERC721EnumerableTransactor // Write-only binding to the contract
	ERC721EnumerableFilterer   // Log filterer for contract events
}

// ERC721EnumerableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721EnumerableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721EnumerableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721EnumerableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721EnumerableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721EnumerableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721EnumerableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721EnumerableSession struct {
	Contract     *ERC721Enumerable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721EnumerableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721EnumerableCallerSession struct {
	Contract *ERC721EnumerableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC721EnumerableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721EnumerableTransactorSession struct {
	Contract     *ERC721EnumerableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC721EnumerableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721EnumerableRaw struct {
	Contract *ERC721Enumerable // Generic contract binding to access the raw methods on
}

// ERC721EnumerableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721EnumerableCallerRaw struct {
	Contract *ERC721EnumerableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721EnumerableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721EnumerableTransactorRaw struct {
	Contract *ERC721EnumerableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Enumerable creates a new instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721Enumerable(address common.Address, backend bind.ContractBackend) (*ERC721Enumerable, error) {
	contract, err := bindERC721Enumerable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Enumerable{ERC721EnumerableCaller: ERC721EnumerableCaller{contract: contract}, ERC721EnumerableTransactor: ERC721EnumerableTransactor{contract: contract}, ERC721EnumerableFilterer: ERC721EnumerableFilterer{contract: contract}}, nil
}

// NewERC721EnumerableCaller creates a new read-only instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721EnumerableCaller(address common.Address, caller bind.ContractCaller) (*ERC721EnumerableCaller, error) {
	contract, err := bindERC721Enumerable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableCaller{contract: contract}, nil
}

// NewERC721EnumerableTransactor creates a new write-only instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721EnumerableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721EnumerableTransactor, error) {
	contract, err := bindERC721Enumerable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableTransactor{contract: contract}, nil
}

// NewERC721EnumerableFilterer creates a new log filterer instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721EnumerableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721EnumerableFilterer, error) {
	contract, err := bindERC721Enumerable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableFilterer{contract: contract}, nil
}

// bindERC721Enumerable binds a generic wrapper to an already deployed contract.
func bindERC721Enumerable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721EnumerableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Enumerable *ERC721EnumerableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Enumerable.Contract.ERC721EnumerableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Enumerable *ERC721EnumerableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.ERC721EnumerableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Enumerable *ERC721EnumerableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.ERC721EnumerableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Enumerable *ERC721EnumerableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Enumerable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Enumerable *ERC721EnumerableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Enumerable *ERC721EnumerableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721Enumerable.Contract.BalanceOf(&_ERC721Enumerable.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721Enumerable.Contract.BalanceOf(&_ERC721Enumerable.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721Enumerable *ERC721EnumerableCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721Enumerable *ERC721EnumerableSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.GetApproved(&_ERC721Enumerable.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.GetApproved(&_ERC721Enumerable.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721Enumerable *ERC721EnumerableSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721Enumerable.Contract.IsApprovedForAll(&_ERC721Enumerable.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721Enumerable.Contract.IsApprovedForAll(&_ERC721Enumerable.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721Enumerable *ERC721EnumerableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721Enumerable *ERC721EnumerableSession) Name() (string, error) {
	return _ERC721Enumerable.Contract.Name(&_ERC721Enumerable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) Name() (string, error) {
	return _ERC721Enumerable.Contract.Name(&_ERC721Enumerable.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721Enumerable *ERC721EnumerableCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721Enumerable *ERC721EnumerableSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.OwnerOf(&_ERC721Enumerable.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.OwnerOf(&_ERC721Enumerable.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721Enumerable *ERC721EnumerableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721Enumerable.Contract.SupportsInterface(&_ERC721Enumerable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721Enumerable.Contract.SupportsInterface(&_ERC721Enumerable.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721Enumerable *ERC721EnumerableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721Enumerable *ERC721EnumerableSession) Symbol() (string, error) {
	return _ERC721Enumerable.Contract.Symbol(&_ERC721Enumerable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) Symbol() (string, error) {
	return _ERC721Enumerable.Contract.Symbol(&_ERC721Enumerable.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenByIndex(&_ERC721Enumerable.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenByIndex(&_ERC721Enumerable.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenOfOwnerByIndex(&_ERC721Enumerable.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenOfOwnerByIndex(&_ERC721Enumerable.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721Enumerable *ERC721EnumerableCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721Enumerable *ERC721EnumerableSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721Enumerable.Contract.TokenURI(&_ERC721Enumerable.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721Enumerable.Contract.TokenURI(&_ERC721Enumerable.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Enumerable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) TotalSupply() (*big.Int, error) {
	return _ERC721Enumerable.Contract.TotalSupply(&_ERC721Enumerable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721Enumerable.Contract.TotalSupply(&_ERC721Enumerable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.Approve(&_ERC721Enumerable.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.Approve(&_ERC721Enumerable.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SafeTransferFrom(&_ERC721Enumerable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SafeTransferFrom(&_ERC721Enumerable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SafeTransferFrom0(&_ERC721Enumerable.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SafeTransferFrom0(&_ERC721Enumerable.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SetApprovalForAll(&_ERC721Enumerable.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SetApprovalForAll(&_ERC721Enumerable.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.TransferFrom(&_ERC721Enumerable.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.TransferFrom(&_ERC721Enumerable.TransactOpts, from, to, tokenId)
}

// ERC721EnumerableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Enumerable contract.
type ERC721EnumerableApprovalIterator struct {
	Event *ERC721EnumerableApproval // Event containing the contract specifics and raw log

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
func (it *ERC721EnumerableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721EnumerableApproval)
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
		it.Event = new(ERC721EnumerableApproval)
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
func (it *ERC721EnumerableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721EnumerableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721EnumerableApproval represents a Approval event raised by the ERC721Enumerable contract.
type ERC721EnumerableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721EnumerableApprovalIterator, error) {

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

	logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableApprovalIterator{contract: _ERC721Enumerable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721EnumerableApproval)
				if err := _ERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC721Enumerable *ERC721EnumerableFilterer) ParseApproval(log types.Log) (*ERC721EnumerableApproval, error) {
	event := new(ERC721EnumerableApproval)
	if err := _ERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721EnumerableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Enumerable contract.
type ERC721EnumerableApprovalForAllIterator struct {
	Event *ERC721EnumerableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721EnumerableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721EnumerableApprovalForAll)
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
		it.Event = new(ERC721EnumerableApprovalForAll)
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
func (it *ERC721EnumerableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721EnumerableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721EnumerableApprovalForAll represents a ApprovalForAll event raised by the ERC721Enumerable contract.
type ERC721EnumerableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721EnumerableApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableApprovalForAllIterator{contract: _ERC721Enumerable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721EnumerableApprovalForAll)
				if err := _ERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_ERC721Enumerable *ERC721EnumerableFilterer) ParseApprovalForAll(log types.Log) (*ERC721EnumerableApprovalForAll, error) {
	event := new(ERC721EnumerableApprovalForAll)
	if err := _ERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721EnumerableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Enumerable contract.
type ERC721EnumerableTransferIterator struct {
	Event *ERC721EnumerableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721EnumerableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721EnumerableTransfer)
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
		it.Event = new(ERC721EnumerableTransfer)
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
func (it *ERC721EnumerableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721EnumerableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721EnumerableTransfer represents a Transfer event raised by the ERC721Enumerable contract.
type ERC721EnumerableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721EnumerableTransferIterator, error) {

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

	logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableTransferIterator{contract: _ERC721Enumerable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721EnumerableTransfer)
				if err := _ERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC721Enumerable *ERC721EnumerableFilterer) ParseTransfer(log types.Log) (*ERC721EnumerableTransfer, error) {
	event := new(ERC721EnumerableTransfer)
	if err := _ERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// MathMetaData contains all meta data concerning the Math contract.
var MathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea264697066735822122021178070beb97b0380ed012bc26206099b474e0c3c403c243ce842e6075d20fd64736f6c63430008150033",
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
	Bin: "0x608060405263618dae80600b5566038d7ea4c68000600c5562278d00600d555f6010555f60145f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600b546015555f601655604051806060016040528060358152602001620053dd6035913960189081620000969190620004bb565b50348015620000a3575f80fd5b506040518060400160405280600d81526020017f52616e646f6d57616c6b4e4654000000000000000000000000000000000000008152506040518060400160405280600481526020017f52574c4b00000000000000000000000000000000000000000000000000000000815250815f9081620001209190620004bb565b508060019081620001329190620004bb565b50505062000155620001496200018d60201b60201c565b6200019460201b60201c565b4243406040516020016200016b9291906200069a565b60405160208183030381529060405280519060200120601381905550620006da565b5f33905090565b5f600a5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600a5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680620002d357607f821691505b602082108103620002e957620002e86200028e565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026200034d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000310565b62000359868362000310565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620003a36200039d620003978462000371565b6200037a565b62000371565b9050919050565b5f819050919050565b620003be8362000383565b620003d6620003cd82620003aa565b8484546200031c565b825550505050565b5f90565b620003ec620003de565b620003f9818484620003b3565b505050565b5b818110156200042057620004145f82620003e2565b600181019050620003ff565b5050565b601f8211156200046f576200043981620002ef565b620004448462000301565b8101602085101562000454578190505b6200046c620004638562000301565b830182620003fe565b50505b505050565b5f82821c905092915050565b5f620004915f198460080262000474565b1980831691505092915050565b5f620004ab838362000480565b9150826002028217905092915050565b620004c68262000257565b67ffffffffffffffff811115620004e257620004e162000261565b5b620004ee8254620002bb565b620004fb82828562000424565b5f60209050601f83116001811462000531575f84156200051c578287015190505b6200052885826200049e565b86555062000597565b601f1984166200054186620002ef565b5f5b828110156200056a5784890151825560018201915060208501945060208101905062000543565b868310156200058a578489015162000586601f89168262000480565b8355505b6001600288020188555050505b505050505050565b5f82825260208201905092915050565b7f412074776f2d64696d656e73696f6e616c2072616e646f6d2077616c6b2077695f8201527f6c6c2072657475726e20746f2074686520706f696e742077686572652069742060208201527f737461727465642c2062757420612074687265652d64696d656e73696f6e616c60408201527f206f6e65206d6179206e6f742e00000000000000000000000000000000000000606082015250565b5f62000657606d836200059f565b91506200066482620005af565b608082019050919050565b6200067a8162000371565b82525050565b5f819050919050565b620006948162000680565b82525050565b5f6060820190508181035f830152620006b38162000649565b9050620006c460208301856200066f565b620006d3604083018462000689565b9392505050565b614cf580620006e85f395ff3fe608060405260043610610250575f3560e01c8063661a1cfe11610138578063a7f93ebd116100b5578063e6e268f411610079578063e6e268f4146108ac578063e985e9c5146108d6578063ee67d6a014610912578063f0503e801461093c578063f2fde38b14610978578063f454aae1146109a057610250565b8063a7f93ebd146107ba578063b88d4fde146107e4578063c87b56dd1461080c578063cb8efe9514610848578063cdb0e89e1461088457610250565b80638da5cb5b116100fc5780638da5cb5b146106ea57806395d89b41146107145780639d4635201461073e578063a035b1fe14610768578063a22cb4651461079257610250565b8063661a1cfe146106085780636e56f6f91461063257806370a082311461066e578063715018a6146106aa57806375794a3c146106c057610250565b80632f745c59116101d1578063438b630011610195578063438b6300146104d857806347ce07cc146105145780634cd609bc1461053e5780634f6ccce71461056857806355f804b3146105a45780636352211e146105cc57610250565b80632f745c59146103f8578063310495ab146104345780633ccfd60b146104705780633e8af4da1461048657806342842e0e146104b057610250565b8063157e394511610218578063157e3945146103285780631596facb1461035257806317d209ab1461037c57806318160ddd146103a657806323b872dd146103d057610250565b806301ffc9a71461025457806306fdde0314610290578063081812fc146102ba578063095ea7b3146102f65780631249c58b1461031e575b5f80fd5b34801561025f575f80fd5b5061027a600480360381019061027591906130ed565b6109dc565b6040516102879190613132565b60405180910390f35b34801561029b575f80fd5b506102a4610a55565b6040516102b191906131d5565b60405180910390f35b3480156102c5575f80fd5b506102e060048036038101906102db9190613228565b610ae4565b6040516102ed9190613292565b60405180910390f35b348015610301575f80fd5b5061031c600480360381019061031791906132d5565b610b65565b005b610326610c7b565b005b348015610333575f80fd5b5061033c610f7a565b6040516103499190613322565b60405180910390f35b34801561035d575f80fd5b50610366610f80565b6040516103739190613322565b60405180910390f35b348015610387575f80fd5b50610390610f86565b60405161039d9190613322565b60405180910390f35b3480156103b1575f80fd5b506103ba610f8c565b6040516103c79190613322565b60405180910390f35b3480156103db575f80fd5b506103f660048036038101906103f1919061333b565b610f98565b005b348015610403575f80fd5b5061041e600480360381019061041991906132d5565b610ff8565b60405161042b9190613322565b60405180910390f35b34801561043f575f80fd5b5061045a60048036038101906104559190613228565b611098565b60405161046791906131d5565b60405180910390f35b34801561047b575f80fd5b50610484611133565b005b348015610491575f80fd5b5061049a6113c3565b6040516104a79190613322565b60405180910390f35b3480156104bb575f80fd5b506104d660048036038101906104d1919061333b565b6113fb565b005b3480156104e3575f80fd5b506104fe60048036038101906104f9919061338b565b61141a565b60405161050b919061346d565b60405180910390f35b34801561051f575f80fd5b5061052861151e565b60405161053591906134a5565b60405180910390f35b348015610549575f80fd5b50610552611524565b60405161055f9190613292565b60405180910390f35b348015610573575f80fd5b5061058e60048036038101906105899190613228565b611549565b60405161059b9190613322565b60405180910390f35b3480156105af575f80fd5b506105ca60048036038101906105c591906135ea565b6115b7565b005b3480156105d7575f80fd5b506105f260048036038101906105ed9190613228565b611646565b6040516105ff9190613292565b60405180910390f35b348015610613575f80fd5b5061061c6116f2565b6040516106299190613322565b60405180910390f35b34801561063d575f80fd5b5061065860048036038101906106539190613228565b611719565b6040516106659190613322565b60405180910390f35b348015610679575f80fd5b50610694600480360381019061068f919061338b565b61172e565b6040516106a19190613322565b60405180910390f35b3480156106b5575f80fd5b506106be6117e2565b005b3480156106cb575f80fd5b506106d4611869565b6040516106e19190613322565b60405180910390f35b3480156106f5575f80fd5b506106fe61186f565b60405161070b9190613292565b60405180910390f35b34801561071f575f80fd5b50610728611897565b60405161073591906131d5565b60405180910390f35b348015610749575f80fd5b50610752611927565b60405161075f9190613322565b60405180910390f35b348015610773575f80fd5b5061077c61192d565b6040516107899190613322565b60405180910390f35b34801561079d575f80fd5b506107b860048036038101906107b3919061365b565b611933565b005b3480156107c5575f80fd5b506107ce611aae565b6040516107db9190613322565b60405180910390f35b3480156107ef575f80fd5b5061080a60048036038101906108059190613737565b611ad1565b005b348015610817575f80fd5b50610832600480360381019061082d9190613228565b611b33565b60405161083f91906131d5565b60405180910390f35b348015610853575f80fd5b5061086e6004803603810190610869919061338b565b611bd7565b60405161087b919061386e565b60405180910390f35b34801561088f575f80fd5b506108aa60048036038101906108a5919061388e565b611cf1565b005b3480156108b7575f80fd5b506108c0611de2565b6040516108cd9190613322565b60405180910390f35b3480156108e1575f80fd5b506108fc60048036038101906108f791906138e8565b611df5565b6040516109099190613132565b60405180910390f35b34801561091d575f80fd5b50610926611e83565b60405161093391906131d5565b60405180910390f35b348015610947575f80fd5b50610962600480360381019061095d9190613228565b611f0f565b60405161096f91906134a5565b60405180910390f35b348015610983575f80fd5b5061099e6004803603810190610999919061338b565b611f24565b005b3480156109ab575f80fd5b506109c660048036038101906109c19190613228565b61201a565b6040516109d39190613322565b60405180910390f35b5f7f780e9d63000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610a4e5750610a4d8261202f565b5b9050919050565b60605f8054610a6390613953565b80601f0160208091040260200160405190810160405280929190818152602001828054610a8f90613953565b8015610ada5780601f10610ab157610100808354040283529160200191610ada565b820191905f5260205f20905b815481529060010190602001808311610abd57829003601f168201915b5050505050905090565b5f610aee82612110565b610b2d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b24906139f3565b60405180910390fd5b60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f610b6f82611646565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610bdf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bd690613a81565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16610bfe612178565b73ffffffffffffffffffffffffffffffffffffffff161480610c2d5750610c2c81610c27612178565b611df5565b5b610c6c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6390613b0f565b60405180910390fd5b610c76838361217f565b505050565b5f610c84611aae565b905080341015610cc9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cc090613b9d565b60405180910390fd5b600b54421015610d0e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d0590613c05565b60405180910390fd5b610d16612178565b60145f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055504260158190555080600c819055505f6016549050600160165f828254610d7b9190613c50565b925050819055506013544243408360145f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051602001610dbf959493929190613c83565b60405160208183030381529060405280519060200120601381905550601354600e5f8381526020019081526020015f2081905550610e1e60145f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682612235565b600c54341115610f00575f60145f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600c5434610e6f9190613cd4565b604051610e7b90613d34565b5f6040518083038185875af1925050503d805f8114610eb5576040519150601f19603f3d011682016040523d82523d5f602084013e610eba565b606091505b5050905080610efe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ef590613d92565b60405180910390fd5b505b60145f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16817fad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec601354600c54604051610f6e929190613db0565b60405180910390a35050565b60105481565b600b5481565b600d5481565b5f600880549050905090565b610fa9610fa3612178565b82612252565b610fe8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fdf90613e47565b60405180910390fd5b610ff383838361232e565b505050565b5f6110028361172e565b8210611043576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103a90613ed5565b60405180910390fd5b60065f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8381526020019081526020015f2054905092915050565b600f602052805f5260405f205f9150905080546110b490613953565b80601f01602080910402602001604051908101604052809291908181526020018280546110e090613953565b801561112b5780601f106111025761010080835404028352916020019161112b565b820191905f5260205f20905b81548152906001019060200180831161110e57829003601f168201915b505050505081565b60145f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16611173612178565b73ffffffffffffffffffffffffffffffffffffffff16146111c9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111c090613f3d565b60405180910390fd5b5f6111d26113c3565b14611212576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161120990613fa5565b60405180910390fd5b5f60145f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f60145f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f60016016546112879190613cd4565b90505f611292611de2565b9050600160105f8282546112a69190613c50565b9250508190555060105460115f8481526020019081526020015f20819055508060125f8481526020019081526020015f20819055505f8373ffffffffffffffffffffffffffffffffffffffff168260405161130090613d34565b5f6040518083038185875af1925050503d805f811461133a576040519150601f19603f3d011682016040523d82523d5f602084013e61133f565b606091505b5050905080611383576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161137a90613d92565b60405180910390fd5b827fa11b556ace4b11a5cae8675a293b51e8cde3a06387d34010861789dfd9e9abc785846040516113b5929190613fc3565b60405180910390a250505050565b5f80600d546015546113d59190613c50565b9050428110156113e8575f9150506113f8565b42816113f49190613cd4565b9150505b90565b61141583838360405180602001604052805f815250611ad1565b505050565b60605f6114268361172e565b90505f8103611480575f67ffffffffffffffff811115611449576114486134c6565b5b6040519080825280602002602001820160405280156114775781602001602082028036833780820191505090505b50915050611519565b5f8167ffffffffffffffff81111561149b5761149a6134c6565b5b6040519080825280602002602001820160405280156114c95781602001602082028036833780820191505090505b5090505f5b82811015611512576114e08582610ff8565b8282815181106114f3576114f2613fea565b5b602002602001018181525050808061150a90614017565b9150506114ce565b5080925050505b919050565b60135481565b60145f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f611552610f8c565b8210611593576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161158a906140ce565b60405180910390fd5b600882815481106115a7576115a6613fea565b5b905f5260205f2001549050919050565b6115bf612178565b73ffffffffffffffffffffffffffffffffffffffff166115dd61186f565b73ffffffffffffffffffffffffffffffffffffffff1614611633576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161162a90614136565b60405180910390fd5b806017908161164291906142f1565b5050565b5f8060025f8481526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036116e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116e090614430565b60405180910390fd5b80915050919050565b5f42600b541015611705575f9050611716565b42600b546117139190613cd4565b90505b90565b6011602052805f5260405f205f915090505481565b5f8073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361179d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611794906144be565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6117ea612178565b73ffffffffffffffffffffffffffffffffffffffff1661180861186f565b73ffffffffffffffffffffffffffffffffffffffff161461185e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161185590614136565b60405180910390fd5b6118675f61257e565b565b60165481565b5f600a5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600180546118a690613953565b80601f01602080910402602001604051908101604052809291908181526020018280546118d290613953565b801561191d5780601f106118f45761010080835404028352916020019161191d565b820191905f5260205f20905b81548152906001019060200180831161190057829003601f168201915b5050505050905090565b60155481565b600c5481565b61193b612178565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036119a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161199f90614526565b60405180910390fd5b8060055f6119b4612178565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff16611a5d612178565b73ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051611aa29190613132565b60405180910390a35050565b5f61271061271b600c54611ac29190614544565b611acc91906145b2565b905090565b611ae2611adc612178565b83612252565b611b21576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b1890613e47565b60405180910390fd5b611b2d84848484612641565b50505050565b6060611b3e82612110565b611b7d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b7490614652565b60405180910390fd5b5f611b8661269d565b90505f815111611ba45760405180602001604052805f815250611bcf565b80611bae8461272d565b604051602001611bbf9291906146aa565b6040516020818303038152906040525b915050919050565b60605f611be38361172e565b90505f8103611c3d575f67ffffffffffffffff811115611c0657611c056134c6565b5b604051908082528060200260200182016040528015611c345781602001602082028036833780820191505090505b50915050611cec565b5f8167ffffffffffffffff811115611c5857611c576134c6565b5b604051908082528060200260200182016040528015611c865781602001602082028036833780820191505090505b5090505f5b82811015611ce5575f611c9e8683610ff8565b9050600e5f8281526020019081526020015f2054838381518110611cc557611cc4613fea565b5b602002602001018181525050508080611cdd90614017565b915050611c8b565b5080925050505b919050565b611d02611cfc612178565b83612252565b611d41576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d389061473d565b60405180910390fd5b602081511115611d86576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d7d906147a5565b60405180910390fd5b80600f5f8481526020019081526020015f209081611da491906142f1565b507f8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f128282604051611dd69291906147c3565b60405180910390a15050565b5f600247611df091906145b2565b905090565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b60188054611e9090613953565b80601f0160208091040260200160405190810160405280929190818152602001828054611ebc90613953565b8015611f075780601f10611ede57610100808354040283529160200191611f07565b820191905f5260205f20905b815481529060010190602001808311611eea57829003601f168201915b505050505081565b600e602052805f5260405f205f915090505481565b611f2c612178565b73ffffffffffffffffffffffffffffffffffffffff16611f4a61186f565b73ffffffffffffffffffffffffffffffffffffffff1614611fa0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f9790614136565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361200e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161200590614861565b60405180910390fd5b6120178161257e565b50565b6012602052805f5260405f205f915090505481565b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806120f957507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80612109575061210882612886565b5b9050919050565b5f8073ffffffffffffffffffffffffffffffffffffffff1660025f8481526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b5f33905090565b8160045f8381526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff166121ef83611646565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b61224e828260405180602001604052805f8152506128ef565b5050565b5f61225c82612110565b61229b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612292906148ef565b60405180910390fd5b5f6122a583611646565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16148061231457508373ffffffffffffffffffffffffffffffffffffffff166122fc84610ae4565b73ffffffffffffffffffffffffffffffffffffffff16145b8061232557506123248185611df5565b5b91505092915050565b8273ffffffffffffffffffffffffffffffffffffffff1661234e82611646565b73ffffffffffffffffffffffffffffffffffffffff16146123a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161239b9061497d565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603612412576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161240990614a0b565b60405180910390fd5b61241d838383612949565b6124275f8261217f565b600160035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546124749190613cd4565b92505081905550600160035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546124c89190613c50565b925050819055508160025f8381526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b5f600a5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600a5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61264c84848461232e565b61265884848484612a59565b612697576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161268e90614a99565b60405180910390fd5b50505050565b6060601780546126ac90613953565b80601f01602080910402602001604051908101604052809291908181526020018280546126d890613953565b80156127235780601f106126fa57610100808354040283529160200191612723565b820191905f5260205f20905b81548152906001019060200180831161270657829003601f168201915b5050505050905090565b60605f8203612773576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050612881565b5f8290505f5b5f82146127a257808061278b90614017565b915050600a8261279b91906145b2565b9150612779565b5f8167ffffffffffffffff8111156127bd576127bc6134c6565b5b6040519080825280601f01601f1916602001820160405280156127ef5781602001600182028036833780820191505090505b5090505b5f851461287a576001826128079190613cd4565b9150600a856128169190614ab7565b60306128229190613c50565b60f81b81838151811061283857612837613fea565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600a8561287391906145b2565b94506127f3565b8093505050505b919050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b6128f98383612bdb565b6129055f848484612a59565b612944576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161293b90614a99565b60405180910390fd5b505050565b612954838383612d9f565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036129955761299081612da4565b6129d4565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146129d3576129d28382612de8565b5b5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603612a1557612a1081612f3e565b612a54565b8273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614612a5357612a528282612ffe565b5b5b505050565b5f612a798473ffffffffffffffffffffffffffffffffffffffff16613076565b15612bce578373ffffffffffffffffffffffffffffffffffffffff1663150b7a02612aa2612178565b8786866040518563ffffffff1660e01b8152600401612ac49493929190614b39565b6020604051808303815f875af1925050508015612aff57506040513d601f19601f82011682018060405250810190612afc9190614b97565b60015b612b7e573d805f8114612b2d576040519150601f19603f3d011682016040523d82523d5f602084013e612b32565b606091505b505f815103612b76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b6d90614a99565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614915050612bd3565b600190505b949350505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603612c49576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c4090614c0c565b60405180910390fd5b612c5281612110565b15612c92576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c8990614c74565b60405180910390fd5b612c9d5f8383612949565b600160035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254612cea9190613c50565b925050819055508160025f8381526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a45050565b505050565b60088054905060095f8381526020019081526020015f2081905550600881908060018154018082558091505060019003905f5260205f20015f909190919091505550565b5f6001612df48461172e565b612dfe9190613cd4565b90505f60075f8481526020019081526020015f20549050818114612ed5575f60065f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8481526020019081526020015f205490508060065f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8481526020019081526020015f20819055508160075f8381526020019081526020015f2081905550505b60075f8481526020019081526020015f205f905560065f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8381526020019081526020015f205f905550505050565b5f6001600880549050612f519190613cd4565b90505f60095f8481526020019081526020015f205490505f60088381548110612f7d57612f7c613fea565b5b905f5260205f20015490508060088381548110612f9d57612f9c613fea565b5b905f5260205f2001819055508160095f8381526020019081526020015f208190555060095f8581526020019081526020015f205f90556008805480612fe557612fe4614c92565b5b600190038181905f5260205f20015f9055905550505050565b5f6130088361172e565b90508160065f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8381526020019081526020015f20819055508060075f8481526020019081526020015f2081905550505050565b5f80823b90505f8111915050919050565b5f604051905090565b5f80fd5b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6130cc81613098565b81146130d6575f80fd5b50565b5f813590506130e7816130c3565b92915050565b5f6020828403121561310257613101613090565b5b5f61310f848285016130d9565b91505092915050565b5f8115159050919050565b61312c81613118565b82525050565b5f6020820190506131455f830184613123565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015613182578082015181840152602081019050613167565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6131a78261314b565b6131b18185613155565b93506131c1818560208601613165565b6131ca8161318d565b840191505092915050565b5f6020820190508181035f8301526131ed818461319d565b905092915050565b5f819050919050565b613207816131f5565b8114613211575f80fd5b50565b5f81359050613222816131fe565b92915050565b5f6020828403121561323d5761323c613090565b5b5f61324a84828501613214565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61327c82613253565b9050919050565b61328c81613272565b82525050565b5f6020820190506132a55f830184613283565b92915050565b6132b481613272565b81146132be575f80fd5b50565b5f813590506132cf816132ab565b92915050565b5f80604083850312156132eb576132ea613090565b5b5f6132f8858286016132c1565b925050602061330985828601613214565b9150509250929050565b61331c816131f5565b82525050565b5f6020820190506133355f830184613313565b92915050565b5f805f6060848603121561335257613351613090565b5b5f61335f868287016132c1565b9350506020613370868287016132c1565b925050604061338186828701613214565b9150509250925092565b5f602082840312156133a05761339f613090565b5b5f6133ad848285016132c1565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6133e8816131f5565b82525050565b5f6133f983836133df565b60208301905092915050565b5f602082019050919050565b5f61341b826133b6565b61342581856133c0565b9350613430836133d0565b805f5b8381101561346057815161344788826133ee565b975061345283613405565b925050600181019050613433565b5085935050505092915050565b5f6020820190508181035f8301526134858184613411565b905092915050565b5f819050919050565b61349f8161348d565b82525050565b5f6020820190506134b85f830184613496565b92915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6134fc8261318d565b810181811067ffffffffffffffff8211171561351b5761351a6134c6565b5b80604052505050565b5f61352d613087565b905061353982826134f3565b919050565b5f67ffffffffffffffff821115613558576135576134c6565b5b6135618261318d565b9050602081019050919050565b828183375f83830152505050565b5f61358e6135898461353e565b613524565b9050828152602081018484840111156135aa576135a96134c2565b5b6135b584828561356e565b509392505050565b5f82601f8301126135d1576135d06134be565b5b81356135e184826020860161357c565b91505092915050565b5f602082840312156135ff576135fe613090565b5b5f82013567ffffffffffffffff81111561361c5761361b613094565b5b613628848285016135bd565b91505092915050565b61363a81613118565b8114613644575f80fd5b50565b5f8135905061365581613631565b92915050565b5f806040838503121561367157613670613090565b5b5f61367e858286016132c1565b925050602061368f85828601613647565b9150509250929050565b5f67ffffffffffffffff8211156136b3576136b26134c6565b5b6136bc8261318d565b9050602081019050919050565b5f6136db6136d684613699565b613524565b9050828152602081018484840111156136f7576136f66134c2565b5b61370284828561356e565b509392505050565b5f82601f83011261371e5761371d6134be565b5b813561372e8482602086016136c9565b91505092915050565b5f805f806080858703121561374f5761374e613090565b5b5f61375c878288016132c1565b945050602061376d878288016132c1565b935050604061377e87828801613214565b925050606085013567ffffffffffffffff81111561379f5761379e613094565b5b6137ab8782880161370a565b91505092959194509250565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6137e98161348d565b82525050565b5f6137fa83836137e0565b60208301905092915050565b5f602082019050919050565b5f61381c826137b7565b61382681856137c1565b9350613831836137d1565b805f5b8381101561386157815161384888826137ef565b975061385383613806565b925050600181019050613834565b5085935050505092915050565b5f6020820190508181035f8301526138868184613812565b905092915050565b5f80604083850312156138a4576138a3613090565b5b5f6138b185828601613214565b925050602083013567ffffffffffffffff8111156138d2576138d1613094565b5b6138de858286016135bd565b9150509250929050565b5f80604083850312156138fe576138fd613090565b5b5f61390b858286016132c1565b925050602061391c858286016132c1565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061396a57607f821691505b60208210810361397d5761397c613926565b5b50919050565b7f4552433732313a20617070726f76656420717565727920666f72206e6f6e65785f8201527f697374656e7420746f6b656e0000000000000000000000000000000000000000602082015250565b5f6139dd602c83613155565b91506139e882613983565b604082019050919050565b5f6020820190508181035f830152613a0a816139d1565b9050919050565b7f4552433732313a20617070726f76616c20746f2063757272656e74206f776e655f8201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b5f613a6b602183613155565b9150613a7682613a11565b604082019050919050565b5f6020820190508181035f830152613a9881613a5f565b9050919050565b7f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f775f8201527f6e6572206e6f7220617070726f76656420666f7220616c6c0000000000000000602082015250565b5f613af9603883613155565b9150613b0482613a9f565b604082019050919050565b5f6020820190508181035f830152613b2681613aed565b9050919050565b7f5468652076616c7565207375626d6974746564207769746820746869732074725f8201527f616e73616374696f6e20697320746f6f206c6f772e0000000000000000000000602082015250565b5f613b87603583613155565b9150613b9282613b2d565b604082019050919050565b5f6020820190508181035f830152613bb481613b7b565b9050919050565b7f5468652073616c65206973206e6f74206f70656e207965742e000000000000005f82015250565b5f613bef601983613155565b9150613bfa82613bbb565b602082019050919050565b5f6020820190508181035f830152613c1c81613be3565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f613c5a826131f5565b9150613c65836131f5565b9250828201905080821115613c7d57613c7c613c23565b5b92915050565b5f60a082019050613c965f830188613496565b613ca36020830187613313565b613cb06040830186613496565b613cbd6060830185613313565b613cca6080830184613283565b9695505050505050565b5f613cde826131f5565b9150613ce9836131f5565b9250828203905081811115613d0157613d00613c23565b5b92915050565b5f81905092915050565b50565b5f613d1f5f83613d07565b9150613d2a82613d11565b5f82019050919050565b5f613d3e82613d14565b9150819050919050565b7f5472616e73666572206661696c65642e000000000000000000000000000000005f82015250565b5f613d7c601083613155565b9150613d8782613d48565b602082019050919050565b5f6020820190508181035f830152613da981613d70565b9050919050565b5f604082019050613dc35f830185613496565b613dd06020830184613313565b9392505050565b7f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f5f8201527f776e6572206e6f7220617070726f766564000000000000000000000000000000602082015250565b5f613e31603183613155565b9150613e3c82613dd7565b604082019050919050565b5f6020820190508181035f830152613e5e81613e25565b9050919050565b7f455243373231456e756d657261626c653a206f776e657220696e646578206f755f8201527f74206f6620626f756e6473000000000000000000000000000000000000000000602082015250565b5f613ebf602b83613155565b9150613eca82613e65565b604082019050919050565b5f6020820190508181035f830152613eec81613eb3565b9050919050565b7f4f6e6c79206c617374206d696e7465722063616e2077697468647261772e00005f82015250565b5f613f27601e83613155565b9150613f3282613ef3565b602082019050919050565b5f6020820190508181035f830152613f5481613f1b565b9050919050565b7f4e6f7420656e6f7567682074696d652068617320656c61707365642e000000005f82015250565b5f613f8f601c83613155565b9150613f9a82613f5b565b602082019050919050565b5f6020820190508181035f830152613fbc81613f83565b9050919050565b5f604082019050613fd65f830185613283565b613fe36020830184613313565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f614021826131f5565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361405357614052613c23565b5b600182019050919050565b7f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f5f8201527f7574206f6620626f756e64730000000000000000000000000000000000000000602082015250565b5f6140b8602c83613155565b91506140c38261405e565b604082019050919050565b5f6020820190508181035f8301526140e5816140ac565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f614120602083613155565b915061412b826140ec565b602082019050919050565b5f6020820190508181035f83015261414d81614114565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026141b07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82614175565b6141ba8683614175565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6141f56141f06141eb846131f5565b6141d2565b6131f5565b9050919050565b5f819050919050565b61420e836141db565b61422261421a826141fc565b848454614181565b825550505050565b5f90565b61423661422a565b614241818484614205565b505050565b5b81811015614264576142595f8261422e565b600181019050614247565b5050565b601f8211156142a95761427a81614154565b61428384614166565b81016020851015614292578190505b6142a661429e85614166565b830182614246565b50505b505050565b5f82821c905092915050565b5f6142c95f19846008026142ae565b1980831691505092915050565b5f6142e183836142ba565b9150826002028217905092915050565b6142fa8261314b565b67ffffffffffffffff811115614313576143126134c6565b5b61431d8254613953565b614328828285614268565b5f60209050601f831160018114614359575f8415614347578287015190505b61435185826142d6565b8655506143b8565b601f19841661436786614154565b5f5b8281101561438e57848901518255600182019150602085019450602081019050614369565b868310156143ab57848901516143a7601f8916826142ba565b8355505b6001600288020188555050505b505050505050565b7f4552433732313a206f776e657220717565727920666f72206e6f6e65786973745f8201527f656e7420746f6b656e0000000000000000000000000000000000000000000000602082015250565b5f61441a602983613155565b9150614425826143c0565b604082019050919050565b5f6020820190508181035f8301526144478161440e565b9050919050565b7f4552433732313a2062616c616e636520717565727920666f7220746865207a655f8201527f726f206164647265737300000000000000000000000000000000000000000000602082015250565b5f6144a8602a83613155565b91506144b38261444e565b604082019050919050565b5f6020820190508181035f8301526144d58161449c565b9050919050565b7f4552433732313a20617070726f766520746f2063616c6c6572000000000000005f82015250565b5f614510601983613155565b915061451b826144dc565b602082019050919050565b5f6020820190508181035f83015261453d81614504565b9050919050565b5f61454e826131f5565b9150614559836131f5565b9250828202614567816131f5565b9150828204841483151761457e5761457d613c23565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6145bc826131f5565b91506145c7836131f5565b9250826145d7576145d6614585565b5b828204905092915050565b7f4552433732314d657461646174613a2055524920717565727920666f72206e6f5f8201527f6e6578697374656e7420746f6b656e0000000000000000000000000000000000602082015250565b5f61463c602f83613155565b9150614647826145e2565b604082019050919050565b5f6020820190508181035f83015261466981614630565b9050919050565b5f81905092915050565b5f6146848261314b565b61468e8185614670565b935061469e818560208601613165565b80840191505092915050565b5f6146b5828561467a565b91506146c1828461467a565b91508190509392505050565b7f736574546f6b656e4e616d652063616c6c6572206973206e6f74206f776e65725f8201527f206e6f7220617070726f76656400000000000000000000000000000000000000602082015250565b5f614727602d83613155565b9150614732826146cd565b604082019050919050565b5f6020820190508181035f8301526147548161471b565b9050919050565b7f546f6b656e206e616d6520697320746f6f206c6f6e672e0000000000000000005f82015250565b5f61478f601783613155565b915061479a8261475b565b602082019050919050565b5f6020820190508181035f8301526147bc81614783565b9050919050565b5f6040820190506147d65f830185613313565b81810360208301526147e8818461319d565b90509392505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f61484b602683613155565b9150614856826147f1565b604082019050919050565b5f6020820190508181035f8301526148788161483f565b9050919050565b7f4552433732313a206f70657261746f7220717565727920666f72206e6f6e65785f8201527f697374656e7420746f6b656e0000000000000000000000000000000000000000602082015250565b5f6148d9602c83613155565b91506148e48261487f565b604082019050919050565b5f6020820190508181035f830152614906816148cd565b9050919050565b7f4552433732313a207472616e73666572206f6620746f6b656e207468617420695f8201527f73206e6f74206f776e0000000000000000000000000000000000000000000000602082015250565b5f614967602983613155565b91506149728261490d565b604082019050919050565b5f6020820190508181035f8301526149948161495b565b9050919050565b7f4552433732313a207472616e7366657220746f20746865207a65726f206164645f8201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b5f6149f5602483613155565b9150614a008261499b565b604082019050919050565b5f6020820190508181035f830152614a22816149e9565b9050919050565b7f4552433732313a207472616e7366657220746f206e6f6e2045524337323152655f8201527f63656976657220696d706c656d656e7465720000000000000000000000000000602082015250565b5f614a83603283613155565b9150614a8e82614a29565b604082019050919050565b5f6020820190508181035f830152614ab081614a77565b9050919050565b5f614ac1826131f5565b9150614acc836131f5565b925082614adc57614adb614585565b5b828206905092915050565b5f81519050919050565b5f82825260208201905092915050565b5f614b0b82614ae7565b614b158185614af1565b9350614b25818560208601613165565b614b2e8161318d565b840191505092915050565b5f608082019050614b4c5f830187613283565b614b596020830186613283565b614b666040830185613313565b8181036060830152614b788184614b01565b905095945050505050565b5f81519050614b91816130c3565b92915050565b5f60208284031215614bac57614bab613090565b5b5f614bb984828501614b83565b91505092915050565b7f4552433732313a206d696e7420746f20746865207a65726f20616464726573735f82015250565b5f614bf6602083613155565b9150614c0182614bc2565b602082019050919050565b5f6020820190508181035f830152614c2381614bea565b9050919050565b7f4552433732313a20746f6b656e20616c7265616479206d696e746564000000005f82015250565b5f614c5e601c83613155565b9150614c6982614c2a565b602082019050919050565b5f6020820190508181035f830152614c8b81614c52565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea26469706673582212200c758b2b472581b10586f3cd487179092796933717e6171b675c73ef438f2c2564736f6c63430008150033697066733a2f2f516d50375a385662514c7079747a586e6365654141633444357458333958567a6f4565555a77454b3861506b3857",
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
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea2646970667358221220c7872ee4c0bd4adc1cd6bf8a34480392bf3a3a5c5390a96e44f916496871e38a64736f6c63430008150033",
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
	Bin: "0x6055604b600b8282823980515f1a607314603f577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea2646970667358221220c7de152dc192bd43ce9b573ce15810e2213407eea7297f1f30aacc39da3971b264736f6c63430008150033",
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
