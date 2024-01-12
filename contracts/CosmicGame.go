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

// CosmicGameMetaData contains all meta data concerning the CosmicGame contract.
var CosmicGameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newActivationTime\",\"type\":\"uint256\"}],\"name\":\"ActivationTimeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lastBidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"bidPrice\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"randomWalkNFTId\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"numCSTTokens\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prizeTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"BidEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCharity\",\"type\":\"address\"}],\"name\":\"CharityAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCharityPercentage\",\"type\":\"uint256\"}],\"name\":\"CharityPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCosmicSignature\",\"type\":\"address\"}],\"name\":\"CosmicSignatureAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCosmicToken\",\"type\":\"address\"}],\"name\":\"CosmicTokenAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddressdonatedNFTs\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"DonatedNFTClaimedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DonationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newETHToCSTBidRatio\",\"type\":\"uint256\"}],\"name\":\"ETHToCSTBidRatioChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newInitialBidAmountFraction\",\"type\":\"uint256\"}],\"name\":\"InitialBidAmountFractionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newInitialSecondsUntilPrize\",\"type\":\"uint256\"}],\"name\":\"InitialSecondsUntilPrizeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newMarketingWallet\",\"type\":\"address\"}],\"name\":\"MarketingWalletAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"NFTDonationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNanoSecondsExtra\",\"type\":\"uint256\"}],\"name\":\"NanoSecondsExtraChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNumHolderNFTWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"NumHolderNFTWinnersPerRoundChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNumRaffleNFTWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"NumRaffleNFTWinnersPerRoundChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNumRaffleWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"NumRaffleWinnersPerRoundChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPriceIncrease\",\"type\":\"uint256\"}],\"name\":\"PriceIncreaseChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PrizeClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrizePercentage\",\"type\":\"uint256\"}],\"name\":\"PrizePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"}],\"name\":\"RaffleNFTWinnerEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRafflePercentage\",\"type\":\"uint256\"}],\"name\":\"RafflePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRaffleWallet\",\"type\":\"address\"}],\"name\":\"RaffleWalletAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRandomWalk\",\"type\":\"address\"}],\"name\":\"RandomWalkAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRwalkPrizePercentage\",\"type\":\"uint256\"}],\"name\":\"RwalkPrizePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStakingPercentage\",\"type\":\"uint256\"}],\"name\":\"StakingPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newStakingWallet\",\"type\":\"address\"}],\"name\":\"StakingWalletAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimeIncrease\",\"type\":\"uint256\"}],\"name\":\"TimeIncreaseChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimeout\",\"type\":\"uint256\"}],\"name\":\"TimeoutClaimPrizeChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CSTAuctionLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHToCSTBidRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MARKETING_REWARD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_MESSAGE_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MILLION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_REWARD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"randomWalkNFTId\",\"type\":\"int256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"randomWalkNFTId\",\"type\":\"int256\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"bidAndDonateNFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"bidWithCST\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charity\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charityAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charityPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"claimDonatedNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokens\",\"type\":\"uint256[]\"}],\"name\":\"claimManyDonatedNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentCSTPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"donatedNFTs\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialBidAmountFraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialSecondsUntilPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBidType\",\"outputs\":[{\"internalType\":\"enumCosmicGame.BidType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCSTBidTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketingWallet\",\"outputs\":[{\"internalType\":\"contractMarketingWallet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nanoSecondsExtra\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractCosmicSignature\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numCSTBids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numDonatedNFTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numETHBids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numHolderNFTWinnersPerRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleNFTWinnersPerRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleParticipants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleWinnersPerRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizeTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raffleAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raffleEntropy\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"raffleParticipants\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rafflePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raffleWallet\",\"outputs\":[{\"internalType\":\"contractRaffleWallet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomWalk\",\"outputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rwalkAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rwalkPrizePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newActivationTime\",\"type\":\"uint256\"}],\"name\":\"setActivationTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setCharity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCharityPercentage\",\"type\":\"uint256\"}],\"name\":\"setCharityPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newETHToCSTBidRatio\",\"type\":\"uint256\"}],\"name\":\"setETHToCSTBidRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newInitialSecondsUntilPrize\",\"type\":\"uint256\"}],\"name\":\"setInitialSecondsUntilPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setMarketingWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNanoSecondsExtra\",\"type\":\"uint256\"}],\"name\":\"setNanoSecondsExtra\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setNftContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNumHolderNFTWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"setNumHolderNFTWinnersPerRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNumRaffleNFTWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleNFTWinnersPerRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newNumRaffleWinnersPerRound\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleWinnersPerRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPriceIncrease\",\"type\":\"uint256\"}],\"name\":\"setPriceIncrease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPrizePercentage\",\"type\":\"uint256\"}],\"name\":\"setPrizePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRafflePercentage\",\"type\":\"uint256\"}],\"name\":\"setRafflePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setRaffleWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setRandomWalk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRwalkPrizePercentage\",\"type\":\"uint256\"}],\"name\":\"setRwalkPrizePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newStakingPercentage\",\"type\":\"uint256\"}],\"name\":\"setStakingPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setStakingWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimeIncrease\",\"type\":\"uint256\"}],\"name\":\"setTimeIncrease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimeout\",\"type\":\"uint256\"}],\"name\":\"setTimeoutClaimPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingWallet\",\"outputs\":[{\"internalType\":\"contractStakingWallet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingBidPriceCST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeUntilActivation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeUntilPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeoutClaimPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractCosmicToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newInitialBidAmountFraction\",\"type\":\"uint256\"}],\"name\":\"updateInitialBidAmountFraction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usedRandomWalkNFTs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"winners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052620f695060015565034630b8a000600255620f42a46003556201518060045566038d7ea4c680006005556000600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506201518060085560c86009556019600a55601e600b55600a600c556005600d55600a600e556003600f5560056010556002601155618ca060125560006013556000601455600a60155568056bc75e2d63100000601655600060185563657a4580601b55601b54601c55348015620000e657600080fd5b5062000107620000fb6200019d60201b60201c565b620001a560201b60201c565b42600143620001179190620002a2565b406040516020016200012b9291906200036a565b60405160208183030381529060405280519060200120601d81905550620001576200019d60201b60201c565b600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620003ac565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000620002af8262000269565b9150620002bc8362000269565b9250828203905081811115620002d757620002d662000273565b5b92915050565b600082825260208201905092915050565b7f436f736d6963205369676e617475726520323032330000000000000000000000600082015250565b600062000326601583620002dd565b91506200033382620002ee565b602082019050919050565b620003498162000269565b82525050565b6000819050919050565b62000364816200034f565b82525050565b60006060820190508181036000830152620003858162000317565b90506200039660208301856200033e565b620003a5604083018462000359565b9392505050565b61700a80620003bc6000396000f3fe6080604052600436106104c45760003560e01c806381e1ccba11610276578063cb819dc01161014f578063e1381d7e116100c1578063ed88c68e11610085578063ed88c68e14611289578063f2fde38b14611293578063f7178822146112bc578063f8c34050146112e7578063fb6f71a314611312578063fc0c546a1461133b57610504565b8063e1381d7e146111b6578063e40c53d2146111df578063e5240baf14611208578063ec34866d14611233578063ed0833981461125e57610504565b8063da4493f611610113578063da4493f6146110c5578063da63990b146110f0578063dbc945c01461111b578063dbf8a28a14611146578063de5af6191461116f578063dec08d8e1461118b57610504565b8063cb819dc014610fdb578063d4ba890214611006578063d59d747814611031578063d6e174171461105a578063d94d03161461109a57610504565b8063a672f6e1116101e8578063c01c5de2116101ac578063c01c5de214610edd578063c2b6d2ea14610f06578063c709bf2814610f31578063c7c8378d14610f5c578063c94028c214610f85578063cb722a3c14610fb057610504565b8063a672f6e114610e0e578063a6ceac2c14610e37578063a6f9cc1514610e62578063ba35b1b014610e8b578063bbcd5bbe14610eb457610504565b80638da5cb5b1161023a5780638da5cb5b14610cfc5780639136d6d914610d275780639250c33c14610d52578063934aa02314610d7d57806397441f2e14610da8578063a2fb117514610dd157610504565b806381e1ccba14610c155780638547af3014610c4057806386e378c914610c6b5780638b12227414610ca85780638b1329e014610cd157610504565b80633f7909d4116103a85780635e6e47aa1161031a578063739a3e02116102de578063739a3e0214610b1557806375f0a87414610b40578063785fa62714610b6b578063799d431d14610b965780637c5486a214610bc157806380de163d14610bea57610504565b80635e6e47aa14610a6a578063647b3e7f14610a9357806367f3e4a514610abc57806370740ac914610ae7578063715018a614610afe57610504565b80635111a2d61161036c5780635111a2d61461097d57806351964588146109a857806352f5ad77146109d157806356c96eb7146109fa5780635d098b3814610a255780635d45c96214610a4e57610504565b80633f7909d4146108aa57806340e02322146108d357806347ccca02146108fe5780634a773f33146109295780634ac3a3951461095457610504565b8063150b7a021161044157806332bc934c1161040557806332bc934c146107a857806332d382cd146107d3578063355f01e2146107fe578063376f3448146108295780633b5fa8e3146108545780633bec7b691461087f57610504565b8063150b7a02146106c157806319afe473146106fe5780631a860c3e146107295780632a027211146107525780632aab32231461077d57610504565b80630e38c32b116104885780630e38c32b146105ec578063119b22b31461061557806311dc733514610640578063120791ca1461066b57806313cc72dd1461069657610504565b8063019d354a1461050957806304a57c091461053257806305ba9b671461056f578063062fb1001461059857806306ee6ad8146105c157610504565b3661050457610502604051806020016040528060008152507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff611366565b005b600080fd5b34801561051557600080fd5b50610530600480360381019061052b9190615549565b611793565b005b34801561053e57600080fd5b5061055960048036038101906105549190615549565b611852565b60405161056691906155b7565b60405180910390f35b34801561057b57600080fd5b5061059660048036038101906105919190615549565b611885565b005b3480156105a457600080fd5b506105bf60048036038101906105ba91906155fe565b611944565b005b3480156105cd57600080fd5b506105d6611aaa565b6040516105e3919061568a565b60405180910390f35b3480156105f857600080fd5b50610613600480360381019061060e9190615549565b611ad0565b005b34801561062157600080fd5b5061062a611bfb565b60405161063791906156b4565b60405180910390f35b34801561064c57600080fd5b50610655611c01565b60405161066291906156b4565b60405180910390f35b34801561067757600080fd5b50610680611c07565b60405161068d91906156b4565b60405180910390f35b3480156106a257600080fd5b506106ab611c28565b6040516106b891906156b4565b60405180910390f35b3480156106cd57600080fd5b506106e860048036038101906106e39190615734565b611c2e565b6040516106f591906157f7565b60405180910390f35b34801561070a57600080fd5b50610713611c43565b60405161072091906156b4565b60405180910390f35b34801561073557600080fd5b50610750600480360381019061074b91906155fe565b611c49565b005b34801561075e57600080fd5b50610767611daf565b60405161077491906156b4565b60405180910390f35b34801561078957600080fd5b50610792611db5565b60405161079f91906156b4565b60405180910390f35b3480156107b457600080fd5b506107bd611dbb565b6040516107ca91906156b4565b60405180910390f35b3480156107df57600080fd5b506107e8611dc2565b6040516107f59190615833565b60405180910390f35b34801561080a57600080fd5b50610813611de8565b60405161082091906156b4565b60405180910390f35b34801561083557600080fd5b5061083e611dee565b60405161084b91906156b4565b60405180910390f35b34801561086057600080fd5b50610869611df4565b60405161087691906156b4565b60405180910390f35b34801561088b57600080fd5b50610894611dfa565b6040516108a191906156b4565b60405180910390f35b3480156108b657600080fd5b506108d160048036038101906108cc919061599d565b611e00565b005b3480156108df57600080fd5b506108e8611e46565b6040516108f591906156b4565b60405180910390f35b34801561090a57600080fd5b50610913611e4c565b6040516109209190615a07565b60405180910390f35b34801561093557600080fd5b5061093e611e72565b60405161094b91906156b4565b60405180910390f35b34801561096057600080fd5b5061097b60048036038101906109769190615549565b611e78565b005b34801561098957600080fd5b50610992611f37565b60405161099f9190615a43565b60405180910390f35b3480156109b457600080fd5b506109cf60048036038101906109ca9190615549565b611f5d565b005b3480156109dd57600080fd5b506109f860048036038101906109f391906155fe565b61201c565b005b348015610a0657600080fd5b50610a0f612182565b604051610a1c91906156b4565b60405180910390f35b348015610a3157600080fd5b50610a4c6004803603810190610a4791906155fe565b61221b565b005b610a686004803603810190610a639190615b49565b611366565b005b348015610a7657600080fd5b50610a916004803603810190610a8c9190615549565b612381565b005b348015610a9f57600080fd5b50610aba6004803603810190610ab59190615549565b6124ac565b005b348015610ac857600080fd5b50610ad161256b565b604051610ade91906156b4565b60405180910390f35b348015610af357600080fd5b50610afc612571565b005b348015610b0a57600080fd5b50610b1361363d565b005b348015610b2157600080fd5b50610b2a6136c5565b604051610b3791906156b4565b60405180910390f35b348015610b4c57600080fd5b50610b556136e6565b604051610b629190615bc6565b60405180910390f35b348015610b7757600080fd5b50610b8061370c565b604051610b8d91906156b4565b60405180910390f35b348015610ba257600080fd5b50610bab61372d565b604051610bb891906156b4565b60405180910390f35b348015610bcd57600080fd5b50610be86004803603810190610be39190615549565b613733565b005b348015610bf657600080fd5b50610bff6137f2565b604051610c0c9190615bfa565b60405180910390f35b348015610c2157600080fd5b50610c2a6137f8565b604051610c3791906156b4565b60405180910390f35b348015610c4c57600080fd5b50610c556137fe565b604051610c6291906155b7565b60405180910390f35b348015610c7757600080fd5b50610c926004803603810190610c8d9190615549565b613824565b604051610c9f9190615c30565b60405180910390f35b348015610cb457600080fd5b50610ccf6004803603810190610cca9190615549565b613844565b005b348015610cdd57600080fd5b50610ce6613903565b604051610cf391906156b4565b60405180910390f35b348015610d0857600080fd5b50610d1161392c565b604051610d1e91906155b7565b60405180910390f35b348015610d3357600080fd5b50610d3c613955565b604051610d4991906156b4565b60405180910390f35b348015610d5e57600080fd5b50610d6761395b565b604051610d7491906156b4565b60405180910390f35b348015610d8957600080fd5b50610d92613961565b604051610d9f91906155b7565b60405180910390f35b348015610db457600080fd5b50610dcf6004803603810190610dca9190615549565b613987565b005b348015610ddd57600080fd5b50610df86004803603810190610df39190615549565b613a46565b604051610e0591906155b7565b60405180910390f35b348015610e1a57600080fd5b50610e356004803603810190610e309190615549565b613a79565b005b348015610e4357600080fd5b50610e4c613b38565b604051610e5991906156b4565b60405180910390f35b348015610e6e57600080fd5b50610e896004803603810190610e8491906155fe565b613b3e565b005b348015610e9757600080fd5b50610eb26004803603810190610ead9190615c4b565b613ca4565b005b348015610ec057600080fd5b50610edb6004803603810190610ed691906155fe565b613e8c565b005b348015610ee957600080fd5b50610f046004803603810190610eff9190615549565b613ff2565b005b348015610f1257600080fd5b50610f1b6140b1565b604051610f2891906156b4565b60405180910390f35b348015610f3d57600080fd5b50610f466140bd565b604051610f5391906156b4565b60405180910390f35b348015610f6857600080fd5b50610f836004803603810190610f7e9190615549565b6140c3565b005b348015610f9157600080fd5b50610f9a6141ee565b604051610fa791906156b4565b60405180910390f35b348015610fbc57600080fd5b50610fc561420f565b604051610fd291906156b4565b60405180910390f35b348015610fe757600080fd5b50610ff0614215565b604051610ffd91906156b4565b60405180910390f35b34801561101257600080fd5b5061101b61421b565b60405161102891906156b4565b60405180910390f35b34801561103d57600080fd5b5061105860048036038101906110539190615549565b614228565b005b34801561106657600080fd5b50611081600480360381019061107c9190615549565b61451b565b6040516110919493929190615cb5565b60405180910390f35b3480156110a657600080fd5b506110af614578565b6040516110bc91906156b4565b60405180910390f35b3480156110d157600080fd5b506110da61457e565b6040516110e791906156b4565b60405180910390f35b3480156110fc57600080fd5b50611105614584565b6040516111129190615d71565b60405180910390f35b34801561112757600080fd5b50611130614597565b60405161113d91906156b4565b60405180910390f35b34801561115257600080fd5b5061116d60048036038101906111689190615549565b6145b8565b005b61118960048036038101906111849190615dca565b6146e3565b005b34801561119757600080fd5b506111a06146fd565b6040516111ad91906156b4565b60405180910390f35b3480156111c257600080fd5b506111dd60048036038101906111d89190615549565b614703565b005b3480156111eb57600080fd5b5061120660048036038101906112019190615549565b6147c2565b005b34801561121457600080fd5b5061121d6148ed565b60405161122a91906156b4565b60405180910390f35b34801561123f57600080fd5b506112486148f3565b60405161125591906156b4565b60405180910390f35b34801561126a57600080fd5b50611273614918565b60405161128091906156b4565b60405180910390f35b61129161491e565b005b34801561129f57600080fd5b506112ba60048036038101906112b591906155fe565b6149cb565b005b3480156112c857600080fd5b506112d1614ac2565b6040516112de91906156b4565b60405180910390f35b3480156112f357600080fd5b506112fc614aeb565b60405161130991906156b4565b60405180910390f35b34801561131e57600080fd5b50611339600480360381019061133491906155fe565b614af1565b005b34801561134757600080fd5b50611350614c79565b60405161135d9190615e6e565b60405180910390f35b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461152b576019600082815260200190815260200160002060009054906101000a900460ff16156113ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113e590615f0c565b60405180910390fd5b6113f6614c9f565b73ffffffffffffffffffffffffffffffffffffffff16602760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e836040518263ffffffff1660e01b815260040161146791906156b4565b602060405180830381865afa158015611484573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114a89190615f41565b73ffffffffffffffffffffffffffffffffffffffff16146114fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114f590615fe0565b60405180910390fd5b60016019600083815260200190815260200160002060006101000a81548160ff0219169083151502179055505b60006115356148f3565b90508034101561157a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161157190616072565b60405180910390fd5b8060058190555060007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83146115b15760016115b4565b60005b9050600060028111156115ca576115c9615cfa565b5b8160028111156115dd576115dc615cfa565b5b036115fd576001601460008282546115f591906160c1565b925050819055505b6116078482614ca7565b6005543411156116ee576000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166005543461165a91906160f5565b6040516116669061615a565b60006040518083038185875af1925050503d80600081146116a3576040519150601f19603f3d011682016040523d82523d6000602084013e6116a8565b606091505b50509050806116ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116e3906161bb565b60405180910390fd5b505b601854600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f3ebe28e9be13fedb71392c114461386e763acb563218b28db3690553055cd5a56005547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff806017548a604051611785959493929190616293565b60405180910390a350505050565b61179b614c9f565b73ffffffffffffffffffffffffffffffffffffffff166117b961392c565b73ffffffffffffffffffffffffffffffffffffffff161461180f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161180690616339565b60405180910390fd5b806004819055507fcaf8e5c6bc1bb2c19935f84ddcdaefb232ad06f9f2abd2ad588bea4bbe631d2660045460405161184791906156b4565b60405180910390a150565b601e6020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61188d614c9f565b73ffffffffffffffffffffffffffffffffffffffff166118ab61392c565b73ffffffffffffffffffffffffffffffffffffffff1614611901576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118f890616339565b60405180910390fd5b806002819055507f678d086a09e1be49401b1e3a6e14db1878e8d8b88e71d0cfe24a32726d0e38b960025460405161193991906156b4565b60405180910390a150565b61194c614c9f565b73ffffffffffffffffffffffffffffffffffffffff1661196a61392c565b73ffffffffffffffffffffffffffffffffffffffff16146119c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119b790616339565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611a2f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a26906163a5565b60405180910390fd5b80602060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f508d510537704d37ed64691a706392abc33e59ce8a327b4952e112820ceb88a681604051611a9f91906155b7565b60405180910390a150565b602160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b611ad8614c9f565b73ffffffffffffffffffffffffffffffffffffffff16611af661392c565b73ffffffffffffffffffffffffffffffffffffffff1614611b4c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b4390616339565b60405180910390fd5b80600e819055506064600e54600d54600c54600a54611b6b91906160c1565b611b7591906160c1565b611b7f91906160c1565b10611bbf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bb690616437565b60405180910390fd5b7f9be9203c5d81ee2019998f7020cf412c92a35ba870318f660d4972660210fb4b600e54604051611bf091906156b4565b60405180910390a150565b60185481565b600a5481565b60006064600b5447611c199190616457565b611c2391906164c8565b905090565b60155481565b600063150b7a0260e01b905095945050505050565b60055481565b611c51614c9f565b73ffffffffffffffffffffffffffffffffffffffff16611c6f61392c565b73ffffffffffffffffffffffffffffffffffffffff1614611cc5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611cbc90616339565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611d34576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d2b906163a5565b60405180910390fd5b80602160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f3d112e567ad7f87ef5e5219a98118d33b03b247b007cfbadf4f133e7010f2c3481604051611da491906155b7565b60405180910390a150565b61011881565b60245481565b620f424081565b602060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60045481565b600b5481565b60125481565b601f5481565b60005b8151811015611e4257611e2f828281518110611e2257611e216164f9565b5b6020026020010151614228565b8080611e3a90616528565b915050611e03565b5050565b60095481565b602660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600f5481565b611e80614c9f565b73ffffffffffffffffffffffffffffffffffffffff16611e9e61392c565b73ffffffffffffffffffffffffffffffffffffffff1614611ef4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611eeb90616339565b60405180910390fd5b806003819055507fed46e73b174ced51fb529cdc6c1b4d8abf49387e6d849b71648afb63c81d12cd600354604051611f2c91906156b4565b60405180910390a150565b602760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b611f65614c9f565b73ffffffffffffffffffffffffffffffffffffffff16611f8361392c565b73ffffffffffffffffffffffffffffffffffffffff1614611fd9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fd090616339565b60405180910390fd5b806008819055507f6da281754ba85ee0c5983a8e8f05a92910c2a0c5b80e68c126216d65f162a30560085460405161201191906156b4565b60405180910390a150565b612024614c9f565b73ffffffffffffffffffffffffffffffffffffffff1661204261392c565b73ffffffffffffffffffffffffffffffffffffffff1614612098576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161208f90616339565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612107576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120fe906163a5565b60405180910390fd5b80602660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f7142a592d5404b5fc4a294c66f70e32b2a7776bd807f722d59268def87c765d18160405161217791906155b7565b60405180910390a150565b600080601c544261219391906160f5565b90506000633b9aca006012546002546121ac9190616457565b6121b691906164c8565b90508082106121ca57600092505050612218565b60008183620f42406121dc9190616457565b6121e691906164c8565b620f42406121f491906160f5565b9050620f4240601654826122089190616457565b61221291906164c8565b93505050505b90565b612223614c9f565b73ffffffffffffffffffffffffffffffffffffffff1661224161392c565b73ffffffffffffffffffffffffffffffffffffffff1614612297576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161228e90616339565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612306576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122fd906163a5565b60405180910390fd5b80602260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f4d03942c29c20d1bccfe551e9d148c917c5a44fb558a4fc60270d8f76fb75f548160405161237691906155b7565b60405180910390a150565b612389614c9f565b73ffffffffffffffffffffffffffffffffffffffff166123a761392c565b73ffffffffffffffffffffffffffffffffffffffff16146123fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123f490616339565b60405180910390fd5b80600c819055506064600e54600d54600c54600a5461241c91906160c1565b61242691906160c1565b61243091906160c1565b10612470576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161246790616437565b60405180910390fd5b7f0918b499c15fcb0e087d411d53664cde23577e1aa4a9cbfbdf735ddd1871e7d5600c546040516124a191906156b4565b60405180910390a150565b6124b4614c9f565b73ffffffffffffffffffffffffffffffffffffffff166124d261392c565b73ffffffffffffffffffffffffffffffffffffffff1614612528576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161251f90616339565b60405180910390fd5b80600f819055507f5e3567ae359e61c47ad2f11d2828dea7e062b2ff63dea385cdb7900a50504c7c600f5460405161256091906156b4565b60405180910390a150565b60165481565b4260175411156125b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125ad906165bc565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603612647576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161263e90616628565b60405180910390fd5b6004546017544261265891906160f5565b10156126f657600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661269f614c9f565b73ffffffffffffffffffffffffffffffffffffffff16146126f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126ec906166e0565b60405180910390fd5b5b6000600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000612742614c9f565b905080601a6000601854815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000602760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612807573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061282b9190616715565b90506000602660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561289c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128c09190616715565b905060006128cc61370c565b9050600160028111156128e2576128e1615cfa565b5b600660149054906101000a900460ff16600281111561290457612903615cfa565b5b0361291457612911611c07565b90505b600061291e614597565b9050600061292a6141ee565b905060006129366136c5565b9050600080861115612a8a57602160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168263b6b55f2560e01b60185460405160240161299791906156b4565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051612a01919061677e565b60006040518083038185875af1925050503d8060008114612a3e576040519150601f19603f3d011682016040523d82523d6000602084013e612a43565b606091505b50508091505080612a89576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a80906167e1565b60405180910390fd5b5b6000602660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1960e01b8a601854604051602401612ae2929190616801565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051612b4c919061677e565b6000604051808303816000865af19150503d8060008114612b89576040519150601f19603f3d011682016040523d82523d6000602084013e612b8e565b606091505b5050905080612bd2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612bc99061689c565b60405180910390fd5b6000805b601054811015612dc957612be861518d565b6000601e6000601f54601d5460001c612c0191906168bc565b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506000602660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1960e01b83601854604051602401612c8a929190616801565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051612cf4919061677e565b6000604051808303816000865af19150503d8060008114612d31576040519150601f19603f3d011682016040523d82523d6000602084013e612d36565b606091505b50915050600081806020019051810190612d509190616715565b9050806018548473ffffffffffffffffffffffffffffffffffffffff167f0afb73a48ffcc4cf27ecb1889154b910544b89c39c5e5822c431062629134e5688604051612d9c91906156b4565b60405180910390a4600185612db191906160c1565b94505050508080612dc190616528565b915050612bd6565b5060005b6011548110156132825760008a111561302357612de861518d565b60008a601d5460001c612dfb91906168bc565b90506000602760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e836040518263ffffffff1660e01b8152600401612e5a91906156b4565b602060405180830381865afa158015612e77573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e9b9190615f41565b90506000602660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1960e01b83601854604051602401612ef5929190616801565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051612f5f919061677e565b6000604051808303816000865af19150503d8060008114612f9c576040519150601f19603f3d011682016040523d82523d6000602084013e612fa1565b606091505b50915050600081806020019051810190612fbb9190616715565b9050806018548473ffffffffffffffffffffffffffffffffffffffff167f0afb73a48ffcc4cf27ecb1889154b910544b89c39c5e5822c431062629134e568960405161300791906156b4565b60405180910390a460018661301c91906160c1565b9550505050505b600089111561326f5761303461518d565b600089601d5460001c61304791906168bc565b90506000602660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e836040518263ffffffff1660e01b81526004016130a691906156b4565b602060405180830381865afa1580156130c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130e79190615f41565b90506000602660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1960e01b83601854604051602401613141929190616801565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516131ab919061677e565b6000604051808303816000865af19150503d80600081146131e8576040519150601f19603f3d011682016040523d82523d6000602084013e6131ed565b606091505b509150506000818060200190518101906132079190616715565b9050806018548473ffffffffffffffffffffffffffffffffffffffff167f0afb73a48ffcc4cf27ecb1889154b910544b89c39c5e5822c431062629134e568960405161325391906156b4565b60405180910390a460018661326891906160c1565b9550505050505b808061327a90616528565b915050612dcd565b508973ffffffffffffffffffffffffffffffffffffffff16876040516132a79061615a565b60006040518083038185875af1925050503d80600081146132e4576040519150601f19603f3d011682016040523d82523d6000602084013e6132e9565b606091505b5050809350508261332f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161332690616939565b60405180910390fd5b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16866040516133759061615a565b60006040518083038185875af1925050503d80600081146133b2576040519150601f19603f3d011682016040523d82523d6000602084013e6133b7565b606091505b505080935050826133fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133f4906169cb565b60405180910390fd5b60005b600f548110156135b55761341261518d565b6000601e6000601f54601d5460001c61342b91906168bc565b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050602060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168763f340fa0160e01b836040516024016134af91906155b7565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051613519919061677e565b60006040518083038185875af1925050503d8060008114613556576040519150601f19603f3d011682016040523d82523d6000602084013e61355b565b606091505b505080955050846135a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161359890616a37565b60405180910390fd5b5080806135ad90616528565b915050613400565b506000601f819055506135c66151cf565b8973ffffffffffffffffffffffffffffffffffffffff166018547f27bc828c399c2947fea27bca8a75ced2e94ff2651d607271f051e39db52286ce8960405161360f91906156b4565b60405180910390a360016018600082825461362a91906160c1565b9250508190555050505050505050505050565b613645614c9f565b73ffffffffffffffffffffffffffffffffffffffff1661366361392c565b73ffffffffffffffffffffffffffffffffffffffff16146136b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136b090616339565b60405180910390fd5b6136c360006151e5565b565b60006064600e54476136d79190616457565b6136e191906164c8565b905090565b602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006064600a544761371e9190616457565b61372891906164c8565b905090565b600c5481565b61373b614c9f565b73ffffffffffffffffffffffffffffffffffffffff1661375961392c565b73ffffffffffffffffffffffffffffffffffffffff16146137af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137a690616339565b60405180910390fd5b80601b819055507f584ff9a7b3a86db5397052f2e440da0ed60a95f646a3a884863cd92262e682b6601b546040516137e791906156b4565b60405180910390a150565b601d5481565b600e5481565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60196020528060005260406000206000915054906101000a900460ff1681565b61384c614c9f565b73ffffffffffffffffffffffffffffffffffffffff1661386a61392c565b73ffffffffffffffffffffffffffffffffffffffff16146138c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016138b790616339565b60405180910390fd5b806001819055507fcbc1f49adfa29e2f2f0f5c9e057722496a4bc95a6a5446deaa423a02b30c64ac6001546040516138f891906156b4565b60405180910390a150565b60004260175410156139185760009050613929565b4260175461392691906160f5565b90505b90565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60025481565b600d5481565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61398f614c9f565b73ffffffffffffffffffffffffffffffffffffffff166139ad61392c565b73ffffffffffffffffffffffffffffffffffffffff1614613a03576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016139fa90616339565b60405180910390fd5b806015819055507f22d4bc2aefb73ef04b923582d7ff2f2b74e699aaf7ff92ea24b348ed59161292601554604051613a3b91906156b4565b60405180910390a150565b601a6020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b613a81614c9f565b73ffffffffffffffffffffffffffffffffffffffff16613a9f61392c565b73ffffffffffffffffffffffffffffffffffffffff1614613af5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613aec90616339565b60405180910390fd5b806009819055507f3b311f029da1b90c3de2e0b3168436c5ed7d8b0ae81b7d4894c12da03835c628600954604051613b2d91906156b4565b60405180910390a150565b60105481565b613b46614c9f565b73ffffffffffffffffffffffffffffffffffffffff16613b6461392c565b73ffffffffffffffffffffffffffffffffffffffff1614613bba576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613bb190616339565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603613c29576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613c20906163a5565b60405180910390fd5b80602760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f9dc3707e4b88be65295b0083b44ffa94059c80372add9b0d24d6a3b371a03b9281604051613c9991906155b7565b60405180910390a150565b6000613cae612182565b90506002613cc568056bc75e2d63100000836152a9565b613ccf9190616457565b60168190555042601c81905550600160136000828254613cef91906160c1565b92505081905550601354601554613d069190616457565b6014541015613d3657610e106002601254613d219190616457565b613d2b91906160c1565b601281905550613d51565b600260126000828254613d4991906164c8565b925050819055505b602560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639dc29fac33836040518363ffffffff1660e01b8152600401613dae929190616801565b600060405180830381600087803b158015613dc857600080fd5b505af1158015613ddc573d6000803e3d6000fd5b50505050613deb826002614ca7565b601854600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f3ebe28e9be13fedb71392c114461386e763acb563218b28db3690553055cd5a57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808560175488604051613e80959493929190616a57565b60405180910390a35050565b613e94614c9f565b73ffffffffffffffffffffffffffffffffffffffff16613eb261392c565b73ffffffffffffffffffffffffffffffffffffffff1614613f08576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613eff90616339565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603613f77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613f6e906163a5565b60405180910390fd5b80602560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f3ab2dcf539aa3c30314265602fc86dee8e343e3c3905730f870835f36c400deb81604051613fe791906155b7565b60405180910390a150565b613ffa614c9f565b73ffffffffffffffffffffffffffffffffffffffff1661401861392c565b73ffffffffffffffffffffffffffffffffffffffff161461406e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161406590616339565b60405180910390fd5b806011819055507f0cc7e37c68566f67d0fe13bf38246d7447cf99a0c481c2ef9963969bb4f5ebc86011546040516140a691906156b4565b60405180910390a150565b67d02ab486cedc000081565b601c5481565b6140cb614c9f565b73ffffffffffffffffffffffffffffffffffffffff166140e961392c565b73ffffffffffffffffffffffffffffffffffffffff161461413f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161413690616339565b60405180910390fd5b80600d819055506064600e54600d54600c54600a5461415e91906160c1565b61416891906160c1565b61417291906160c1565b106141b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016141a990616437565b60405180910390fd5b7fd2002a256ab2f8a5c1b865685754b20437c0a26e7249e40ae5df2993966f99f2600d546040516141e391906156b4565b60405180910390a150565b60006064600d54476142009190616457565b61420a91906164c8565b905090565b60145481565b60175481565b68056bc75e2d6310000081565b602454811061426c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161426390616afd565b60405180910390fd5b6000601a60006023600085815260200190815260200160002060020154815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603614329576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161432090616b8f565b60405180910390fd5b6023600083815260200190815260200160002060030160009054906101000a900460ff161561438d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161438490616c21565b60405180910390fd5b60016023600084815260200190815260200160002060030160006101000a81548160ff0219169083151502179055506023600083815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166342842e0e308360236000878152602001908152602001600020600101546040518463ffffffff1660e01b815260040161444593929190616c41565b600060405180830381600087803b15801561445f57600080fd5b505af1158015614473573d6000803e3d6000fd5b5050505060236000838152602001908152602001600020600201547f0d0e4b90a96d931964d5b0421a6b5b48dd73abf538cdb9ae015069d49e9a467983836023600087815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16602360008881526020019081526020016000206001015460405161450f9493929190616c78565b60405180910390a25050565b60236020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030160009054906101000a900460ff16905084565b60035481565b601b5481565b600660149054906101000a900460ff1681565b60006064600c54476145a99190616457565b6145b391906164c8565b905090565b6145c0614c9f565b73ffffffffffffffffffffffffffffffffffffffff166145de61392c565b73ffffffffffffffffffffffffffffffffffffffff1614614634576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161462b90616339565b60405180910390fd5b80600a819055506064600e54600d54600c54600a5461465391906160c1565b61465d91906160c1565b61466791906160c1565b106146a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161469e90616437565b60405180910390fd5b7f595fa5ba64cd6f66df19b53b59bb4a275bae1ec5b362da95e1fa4c5feb62be1e600a546040516146d891906156b4565b60405180910390a150565b6146ed8484611366565b6146f782826152c3565b50505050565b60115481565b61470b614c9f565b73ffffffffffffffffffffffffffffffffffffffff1661472961392c565b73ffffffffffffffffffffffffffffffffffffffff161461477f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161477690616339565b60405180910390fd5b806010819055507f72e4278828b8a868e0ba5b4887f954797ea786f8bac991128636171f5eed471d6010546040516147b791906156b4565b60405180910390a150565b6147ca614c9f565b73ffffffffffffffffffffffffffffffffffffffff166147e861392c565b73ffffffffffffffffffffffffffffffffffffffff161461483e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161483590616339565b60405180910390fd5b80600b819055506064600e54600d54600c54600b5461485d91906160c1565b61486791906160c1565b61487191906160c1565b106148b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016148a890616437565b60405180910390fd5b7ff6a623223b73e9ec7d5b192624a4999b2fa4110d359f6e192ef38e1a1541ee87600b546040516148e291906156b4565b60405180910390a150565b60135481565b6000620f42406001546005546149099190616457565b61491391906164c8565b905090565b60085481565b60003411614961576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161495890616d2f565b60405180910390fd5b601b54421015614974576149736151cf565b5b61497c614c9f565b73ffffffffffffffffffffffffffffffffffffffff167f8b7fe5be5699654fd637d2250cb0d47e88205730710745e78e9d8bcaf8aad8f1346040516149c191906156b4565b60405180910390a2565b6149d3614c9f565b73ffffffffffffffffffffffffffffffffffffffff166149f161392c565b73ffffffffffffffffffffffffffffffffffffffff1614614a47576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614a3e90616339565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603614ab6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614aad90616dc1565b60405180910390fd5b614abf816151e5565b50565b600042601b541015614ad75760009050614ae8565b42601b54614ae591906160f5565b90505b90565b60015481565b614af9614c9f565b73ffffffffffffffffffffffffffffffffffffffff16614b1761392c565b73ffffffffffffffffffffffffffffffffffffffff1614614b6d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614b6490616339565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603614bdc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614bd3906163a5565b60405180910390fd5b80600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051614c6e91906155b7565b60405180910390a150565b602560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600033905090565b601b54421015614cec576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614ce390616e2d565b60405180910390fd5b61011882511115614d32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614d2990616e99565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603614d9d5760085442614d9691906160c1565b6017819055505b614da5614c9f565b600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600660146101000a81548160ff02191690836002811115614e0a57614e09615cfa565b5b0217905550600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16601e6000601f54815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001601f6000828254614e9891906160c1565b925050819055506000602560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1960e01b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1668056bc75e2d63100000604051602401614f20929190616801565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051614f8a919061677e565b6000604051808303816000865af19150503d8060008114614fc7576040519150601f19603f3d011682016040523d82523d6000602084013e614fcc565b606091505b5050905080615010576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161500790616f2b565b60405180910390fd5b602560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1960e01b602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1667d02ab486cedc000060405160240161508e929190616f4b565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516150f8919061677e565b6000604051808303816000865af19150503d8060008114615135576040519150601f19603f3d011682016040523d82523d6000602084013e61513a565b606091505b50508091505080615180576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161517790616f2b565b60405180910390fd5b6151886154a5565b505050565b601d544260014361519e91906160f5565b406040516020016151b193929190616f74565b60405160208183030381529060405280519060200120601d81905550565b600954476151dd91906164c8565b600581905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000818310156152b957816152bb565b825b905092915050565b8173ffffffffffffffffffffffffffffffffffffffff166342842e0e6152e7614c9f565b30846040518463ffffffff1660e01b815260040161530793929190616c41565b600060405180830381600087803b15801561532157600080fd5b505af1158015615335573d6000803e3d6000fd5b5050505060405180608001604052808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200160185481526020016000151581525060236000602454815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002015560608201518160030160006101000a81548160ff02191690831515021790555090505060016024600082825461541b91906160c1565b925050819055506018548273ffffffffffffffffffffffffffffffffffffffff16615444614c9f565b73ffffffffffffffffffffffffffffffffffffffff167fc85be1734ed6a0f4d6adf924d4d41406e2729878c652110a96e2fdec64e118d184600160245461548b91906160f5565b604051615499929190616fab565b60405180910390a45050565b6000633b9aca006002546154b991906164c8565b9050806154c8601754426152a9565b6154d291906160c1565b601781905550620f42406003546002546154ec9190616457565b6154f691906164c8565b60028190555050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61552681615513565b811461553157600080fd5b50565b6000813590506155438161551d565b92915050565b60006020828403121561555f5761555e615509565b5b600061556d84828501615534565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006155a182615576565b9050919050565b6155b181615596565b82525050565b60006020820190506155cc60008301846155a8565b92915050565b6155db81615596565b81146155e657600080fd5b50565b6000813590506155f8816155d2565b92915050565b60006020828403121561561457615613615509565b5b6000615622848285016155e9565b91505092915050565b6000819050919050565b600061565061564b61564684615576565b61562b565b615576565b9050919050565b600061566282615635565b9050919050565b600061567482615657565b9050919050565b61568481615669565b82525050565b600060208201905061569f600083018461567b565b92915050565b6156ae81615513565b82525050565b60006020820190506156c960008301846156a5565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126156f4576156f36156cf565b5b8235905067ffffffffffffffff811115615711576157106156d4565b5b60208301915083600182028301111561572d5761572c6156d9565b5b9250929050565b6000806000806000608086880312156157505761574f615509565b5b600061575e888289016155e9565b955050602061576f888289016155e9565b945050604061578088828901615534565b935050606086013567ffffffffffffffff8111156157a1576157a061550e565b5b6157ad888289016156de565b92509250509295509295909350565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6157f1816157bc565b82525050565b600060208201905061580c60008301846157e8565b92915050565b600061581d82615657565b9050919050565b61582d81615812565b82525050565b60006020820190506158486000830184615824565b92915050565b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6158978261584e565b810181811067ffffffffffffffff821117156158b6576158b561585f565b5b80604052505050565b60006158c96154ff565b90506158d5828261588e565b919050565b600067ffffffffffffffff8211156158f5576158f461585f565b5b602082029050602081019050919050565b6000615919615914846158da565b6158bf565b9050808382526020820190506020840283018581111561593c5761593b6156d9565b5b835b8181101561596557806159518882615534565b84526020840193505060208101905061593e565b5050509392505050565b600082601f830112615984576159836156cf565b5b8135615994848260208601615906565b91505092915050565b6000602082840312156159b3576159b2615509565b5b600082013567ffffffffffffffff8111156159d1576159d061550e565b5b6159dd8482850161596f565b91505092915050565b60006159f182615657565b9050919050565b615a01816159e6565b82525050565b6000602082019050615a1c60008301846159f8565b92915050565b6000615a2d82615657565b9050919050565b615a3d81615a22565b82525050565b6000602082019050615a586000830184615a34565b92915050565b600080fd5b600067ffffffffffffffff821115615a7e57615a7d61585f565b5b615a878261584e565b9050602081019050919050565b82818337600083830152505050565b6000615ab6615ab184615a63565b6158bf565b905082815260208101848484011115615ad257615ad1615a5e565b5b615add848285615a94565b509392505050565b600082601f830112615afa57615af96156cf565b5b8135615b0a848260208601615aa3565b91505092915050565b6000819050919050565b615b2681615b13565b8114615b3157600080fd5b50565b600081359050615b4381615b1d565b92915050565b60008060408385031215615b6057615b5f615509565b5b600083013567ffffffffffffffff811115615b7e57615b7d61550e565b5b615b8a85828601615ae5565b9250506020615b9b85828601615b34565b9150509250929050565b6000615bb082615657565b9050919050565b615bc081615ba5565b82525050565b6000602082019050615bdb6000830184615bb7565b92915050565b6000819050919050565b615bf481615be1565b82525050565b6000602082019050615c0f6000830184615beb565b92915050565b60008115159050919050565b615c2a81615c15565b82525050565b6000602082019050615c456000830184615c21565b92915050565b600060208284031215615c6157615c60615509565b5b600082013567ffffffffffffffff811115615c7f57615c7e61550e565b5b615c8b84828501615ae5565b91505092915050565b6000615c9f82615657565b9050919050565b615caf81615c94565b82525050565b6000608082019050615cca6000830187615ca6565b615cd760208301866156a5565b615ce460408301856156a5565b615cf16060830184615c21565b95945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038110615d3a57615d39615cfa565b5b50565b6000819050615d4b82615d29565b919050565b6000615d5b82615d3d565b9050919050565b615d6b81615d50565b82525050565b6000602082019050615d866000830184615d62565b92915050565b6000615d9782615596565b9050919050565b615da781615d8c565b8114615db257600080fd5b50565b600081359050615dc481615d9e565b92915050565b60008060008060808587031215615de457615de3615509565b5b600085013567ffffffffffffffff811115615e0257615e0161550e565b5b615e0e87828801615ae5565b9450506020615e1f87828801615b34565b9350506040615e3087828801615db5565b9250506060615e4187828801615534565b91505092959194509250565b6000615e5882615657565b9050919050565b615e6881615e4d565b82525050565b6000602082019050615e836000830184615e5f565b92915050565b600082825260208201905092915050565b7f546869732052616e646f6d57616c6b4e46542068617320616c7265616479206260008201527f65656e207573656420666f722062696464696e672e0000000000000000000000602082015250565b6000615ef6603583615e89565b9150615f0182615e9a565b604082019050919050565b60006020820190508181036000830152615f2581615ee9565b9050919050565b600081519050615f3b816155d2565b92915050565b600060208284031215615f5757615f56615509565b5b6000615f6584828501615f2c565b91505092915050565b7f596f75206d75737420626520746865206f776e6572206f66207468652052616e60008201527f646f6d57616c6b4e46542e000000000000000000000000000000000000000000602082015250565b6000615fca602b83615e89565b9150615fd582615f6e565b604082019050919050565b60006020820190508181036000830152615ff981615fbd565b9050919050565b7f5468652076616c7565207375626d69747465642077697468207468697320747260008201527f616e73616374696f6e20697320746f6f206c6f772e0000000000000000000000602082015250565b600061605c603583615e89565b915061606782616000565b604082019050919050565b6000602082019050818103600083015261608b8161604f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006160cc82615513565b91506160d783615513565b92508282019050808211156160ef576160ee616092565b5b92915050565b600061610082615513565b915061610b83615513565b925082820390508181111561612357616122616092565b5b92915050565b600081905092915050565b50565b6000616144600083616129565b915061614f82616134565b600082019050919050565b600061616582616137565b9150819050919050565b7f526566756e64207472616e73666572206661696c65642e000000000000000000600082015250565b60006161a5601783615e89565b91506161b08261616f565b602082019050919050565b600060208201905081810360008301526161d481616198565b9050919050565b6161e481615b13565b82525050565b6000819050919050565b600061620f61620a616205846161ea565b61562b565b615b13565b9050919050565b61621f816161f4565b82525050565b600081519050919050565b60005b8381101561624e578082015181840152602081019050616233565b60008484015250505050565b600061626582616225565b61626f8185615e89565b935061627f818560208601616230565b6162888161584e565b840191505092915050565b600060a0820190506162a860008301886161db565b6162b56020830187616216565b6162c26040830186616216565b6162cf60608301856156a5565b81810360808301526162e1818461625a565b90509695505050505050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000616323602083615e89565b915061632e826162ed565b602082019050919050565b6000602082019050818103600083015261635281616316565b9050919050565b7f5a65726f2d616464726573732077617320676976656e2e000000000000000000600082015250565b600061638f601783615e89565b915061639a82616359565b602082019050919050565b600060208201905081810360008301526163be81616382565b9050919050565b7f50657263656e746167652076616c7565206f766572666c6f772c206d7573742060008201527f6265206c6f776572207468616e203130302e0000000000000000000000000000602082015250565b6000616421603283615e89565b915061642c826163c5565b604082019050919050565b6000602082019050818103600083015261645081616414565b9050919050565b600061646282615513565b915061646d83615513565b925082820261647b81615513565b9150828204841483151761649257616491616092565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006164d382615513565b91506164de83615513565b9250826164ee576164ed616499565b5b828204905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061653382615513565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361656557616564616092565b5b600182019050919050565b7f4e6f7420656e6f7567682074696d652068617320656c61707365642e00000000600082015250565b60006165a6601c83615e89565b91506165b182616570565b602082019050919050565b600060208201905081810360008301526165d581616599565b9050919050565b7f5468657265206973206e6f206c617374206269646465722e0000000000000000600082015250565b6000616612601883615e89565b915061661d826165dc565b602082019050919050565b6000602082019050818103600083015261664181616605565b9050919050565b7f4f6e6c7920746865206c617374206269646465722063616e20636c61696d207460008201527f6865207072697a6520647572696e672074686520666972737420323420686f7560208201527f72732e0000000000000000000000000000000000000000000000000000000000604082015250565b60006166ca604383615e89565b91506166d582616648565b606082019050919050565b600060208201905081810360008301526166f9816166bd565b9050919050565b60008151905061670f8161551d565b92915050565b60006020828403121561672b5761672a615509565b5b600061673984828501616700565b91505092915050565b600081519050919050565b600061675882616742565b6167628185616129565b9350616772818560208601616230565b80840191505092915050565b600061678a828461674d565b915081905092915050565b7f5374616b696e67206465706f736974206661696c65642e000000000000000000600082015250565b60006167cb601783615e89565b91506167d682616795565b602082019050919050565b600060208201905081810360008301526167fa816167be565b9050919050565b600060408201905061681660008301856155a8565b61682360208301846156a5565b9392505050565b7f436f736d69635369676e6174757265206d696e742829206661696c656420746f60008201527f206d696e74204e46542e00000000000000000000000000000000000000000000602082015250565b6000616886602a83615e89565b91506168918261682a565b604082019050919050565b600060208201905081810360008301526168b581616879565b9050919050565b60006168c782615513565b91506168d283615513565b9250826168e2576168e1616499565b5b828206905092915050565b7f5472616e7366657220746f207468652077696e6e6572206661696c65642e0000600082015250565b6000616923601e83615e89565b915061692e826168ed565b602082019050919050565b6000602082019050818103600083015261695281616916565b9050919050565b7f5472616e7366657220746f206368617269747920636f6e74726163742066616960008201527f6c65642e00000000000000000000000000000000000000000000000000000000602082015250565b60006169b5602483615e89565b91506169c082616959565b604082019050919050565b600060208201905081810360008301526169e4816169a8565b9050919050565b7f526166666c65206465706f736974206661696c65642e00000000000000000000600082015250565b6000616a21601683615e89565b9150616a2c826169eb565b602082019050919050565b60006020820190508181036000830152616a5081616a14565b9050919050565b600060a082019050616a6c6000830188616216565b616a796020830187616216565b616a8660408301866161db565b616a9360608301856156a5565b8181036080830152616aa5818461625a565b90509695505050505050565b7f54686520646f6e61746564204e465420646f6573206e6f742065786973742e00600082015250565b6000616ae7601f83615e89565b9150616af282616ab1565b602082019050919050565b60006020820190508181036000830152616b1681616ada565b9050919050565b7f4e6f6e2d6578697374656e742077696e6e657220666f722074686520726f756e60008201527f642e000000000000000000000000000000000000000000000000000000000000602082015250565b6000616b79602283615e89565b9150616b8482616b1d565b604082019050919050565b60006020820190508181036000830152616ba881616b6c565b9050919050565b7f546865204e46542068617320616c7265616479206265656e20636c61696d656460008201527f2e00000000000000000000000000000000000000000000000000000000000000602082015250565b6000616c0b602183615e89565b9150616c1682616baf565b604082019050919050565b60006020820190508181036000830152616c3a81616bfe565b9050919050565b6000606082019050616c5660008301866155a8565b616c6360208301856155a8565b616c7060408301846156a5565b949350505050565b6000608082019050616c8d60008301876156a5565b616c9a60208301866155a8565b616ca760408301856155a8565b616cb460608301846156a5565b95945050505050565b7f446f6e6174696f6e20616d6f756e74206d75737420626520677265617465722060008201527f7468616e20302e00000000000000000000000000000000000000000000000000602082015250565b6000616d19602783615e89565b9150616d2482616cbd565b604082019050919050565b60006020820190508181036000830152616d4881616d0c565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000616dab602683615e89565b9150616db682616d4f565b604082019050919050565b60006020820190508181036000830152616dda81616d9e565b9050919050565b7f4e6f7420616374697665207965742e0000000000000000000000000000000000600082015250565b6000616e17600f83615e89565b9150616e2282616de1565b602082019050919050565b60006020820190508181036000830152616e4681616e0a565b9050919050565b7f4d65737361676520697320746f6f206c6f6e672e000000000000000000000000600082015250565b6000616e83601483615e89565b9150616e8e82616e4d565b602082019050919050565b60006020820190508181036000830152616eb281616e76565b9050919050565b7f436f736d6963546f6b656e206d696e742829206661696c656420746f206d696e60008201527f742072657761726420746f6b656e732e00000000000000000000000000000000602082015250565b6000616f15603083615e89565b9150616f2082616eb9565b604082019050919050565b60006020820190508181036000830152616f4481616f08565b9050919050565b6000604082019050616f606000830185615bb7565b616f6d60208301846156a5565b9392505050565b6000606082019050616f896000830186615beb565b616f9660208301856156a5565b616fa36040830184615beb565b949350505050565b6000604082019050616fc060008301856156a5565b616fcd60208301846156a5565b939250505056fea2646970667358221220c6467cd8c2256d687cad804360d95f64a591e980c9a6f4aa5c7c1a3aa9834e3064736f6c63430008130033",
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
	parsed, err := CosmicGameMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// CSTAuctionLength is a free data retrieval call binding the contract method 0x3b5fa8e3.
//
// Solidity: function CSTAuctionLength() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) CSTAuctionLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "CSTAuctionLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CSTAuctionLength is a free data retrieval call binding the contract method 0x3b5fa8e3.
//
// Solidity: function CSTAuctionLength() view returns(uint256)
func (_CosmicGame *CosmicGameSession) CSTAuctionLength() (*big.Int, error) {
	return _CosmicGame.Contract.CSTAuctionLength(&_CosmicGame.CallOpts)
}

// CSTAuctionLength is a free data retrieval call binding the contract method 0x3b5fa8e3.
//
// Solidity: function CSTAuctionLength() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) CSTAuctionLength() (*big.Int, error) {
	return _CosmicGame.Contract.CSTAuctionLength(&_CosmicGame.CallOpts)
}

// ETHToCSTBidRatio is a free data retrieval call binding the contract method 0x13cc72dd.
//
// Solidity: function ETHToCSTBidRatio() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) ETHToCSTBidRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "ETHToCSTBidRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ETHToCSTBidRatio is a free data retrieval call binding the contract method 0x13cc72dd.
//
// Solidity: function ETHToCSTBidRatio() view returns(uint256)
func (_CosmicGame *CosmicGameSession) ETHToCSTBidRatio() (*big.Int, error) {
	return _CosmicGame.Contract.ETHToCSTBidRatio(&_CosmicGame.CallOpts)
}

// ETHToCSTBidRatio is a free data retrieval call binding the contract method 0x13cc72dd.
//
// Solidity: function ETHToCSTBidRatio() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) ETHToCSTBidRatio() (*big.Int, error) {
	return _CosmicGame.Contract.ETHToCSTBidRatio(&_CosmicGame.CallOpts)
}

// MARKETINGREWARD is a free data retrieval call binding the contract method 0xc2b6d2ea.
//
// Solidity: function MARKETING_REWARD() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) MARKETINGREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "MARKETING_REWARD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MARKETINGREWARD is a free data retrieval call binding the contract method 0xc2b6d2ea.
//
// Solidity: function MARKETING_REWARD() view returns(uint256)
func (_CosmicGame *CosmicGameSession) MARKETINGREWARD() (*big.Int, error) {
	return _CosmicGame.Contract.MARKETINGREWARD(&_CosmicGame.CallOpts)
}

// MARKETINGREWARD is a free data retrieval call binding the contract method 0xc2b6d2ea.
//
// Solidity: function MARKETING_REWARD() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) MARKETINGREWARD() (*big.Int, error) {
	return _CosmicGame.Contract.MARKETINGREWARD(&_CosmicGame.CallOpts)
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

// CurrentCSTPrice is a free data retrieval call binding the contract method 0x56c96eb7.
//
// Solidity: function currentCSTPrice() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) CurrentCSTPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "currentCSTPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentCSTPrice is a free data retrieval call binding the contract method 0x56c96eb7.
//
// Solidity: function currentCSTPrice() view returns(uint256)
func (_CosmicGame *CosmicGameSession) CurrentCSTPrice() (*big.Int, error) {
	return _CosmicGame.Contract.CurrentCSTPrice(&_CosmicGame.CallOpts)
}

// CurrentCSTPrice is a free data retrieval call binding the contract method 0x56c96eb7.
//
// Solidity: function currentCSTPrice() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) CurrentCSTPrice() (*big.Int, error) {
	return _CosmicGame.Contract.CurrentCSTPrice(&_CosmicGame.CallOpts)
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

// LastBidType is a free data retrieval call binding the contract method 0xda63990b.
//
// Solidity: function lastBidType() view returns(uint8)
func (_CosmicGame *CosmicGameCaller) LastBidType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "lastBidType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// LastBidType is a free data retrieval call binding the contract method 0xda63990b.
//
// Solidity: function lastBidType() view returns(uint8)
func (_CosmicGame *CosmicGameSession) LastBidType() (uint8, error) {
	return _CosmicGame.Contract.LastBidType(&_CosmicGame.CallOpts)
}

// LastBidType is a free data retrieval call binding the contract method 0xda63990b.
//
// Solidity: function lastBidType() view returns(uint8)
func (_CosmicGame *CosmicGameCallerSession) LastBidType() (uint8, error) {
	return _CosmicGame.Contract.LastBidType(&_CosmicGame.CallOpts)
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

// LastCSTBidTime is a free data retrieval call binding the contract method 0xc709bf28.
//
// Solidity: function lastCSTBidTime() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) LastCSTBidTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "lastCSTBidTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastCSTBidTime is a free data retrieval call binding the contract method 0xc709bf28.
//
// Solidity: function lastCSTBidTime() view returns(uint256)
func (_CosmicGame *CosmicGameSession) LastCSTBidTime() (*big.Int, error) {
	return _CosmicGame.Contract.LastCSTBidTime(&_CosmicGame.CallOpts)
}

// LastCSTBidTime is a free data retrieval call binding the contract method 0xc709bf28.
//
// Solidity: function lastCSTBidTime() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) LastCSTBidTime() (*big.Int, error) {
	return _CosmicGame.Contract.LastCSTBidTime(&_CosmicGame.CallOpts)
}

// MarketingWallet is a free data retrieval call binding the contract method 0x75f0a874.
//
// Solidity: function marketingWallet() view returns(address)
func (_CosmicGame *CosmicGameCaller) MarketingWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "marketingWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketingWallet is a free data retrieval call binding the contract method 0x75f0a874.
//
// Solidity: function marketingWallet() view returns(address)
func (_CosmicGame *CosmicGameSession) MarketingWallet() (common.Address, error) {
	return _CosmicGame.Contract.MarketingWallet(&_CosmicGame.CallOpts)
}

// MarketingWallet is a free data retrieval call binding the contract method 0x75f0a874.
//
// Solidity: function marketingWallet() view returns(address)
func (_CosmicGame *CosmicGameCallerSession) MarketingWallet() (common.Address, error) {
	return _CosmicGame.Contract.MarketingWallet(&_CosmicGame.CallOpts)
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

// NumCSTBids is a free data retrieval call binding the contract method 0xe5240baf.
//
// Solidity: function numCSTBids() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) NumCSTBids(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "numCSTBids")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumCSTBids is a free data retrieval call binding the contract method 0xe5240baf.
//
// Solidity: function numCSTBids() view returns(uint256)
func (_CosmicGame *CosmicGameSession) NumCSTBids() (*big.Int, error) {
	return _CosmicGame.Contract.NumCSTBids(&_CosmicGame.CallOpts)
}

// NumCSTBids is a free data retrieval call binding the contract method 0xe5240baf.
//
// Solidity: function numCSTBids() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) NumCSTBids() (*big.Int, error) {
	return _CosmicGame.Contract.NumCSTBids(&_CosmicGame.CallOpts)
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

// NumETHBids is a free data retrieval call binding the contract method 0xcb722a3c.
//
// Solidity: function numETHBids() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) NumETHBids(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "numETHBids")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumETHBids is a free data retrieval call binding the contract method 0xcb722a3c.
//
// Solidity: function numETHBids() view returns(uint256)
func (_CosmicGame *CosmicGameSession) NumETHBids() (*big.Int, error) {
	return _CosmicGame.Contract.NumETHBids(&_CosmicGame.CallOpts)
}

// NumETHBids is a free data retrieval call binding the contract method 0xcb722a3c.
//
// Solidity: function numETHBids() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) NumETHBids() (*big.Int, error) {
	return _CosmicGame.Contract.NumETHBids(&_CosmicGame.CallOpts)
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

// RwalkAmount is a free data retrieval call binding the contract method 0x120791ca.
//
// Solidity: function rwalkAmount() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) RwalkAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "rwalkAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RwalkAmount is a free data retrieval call binding the contract method 0x120791ca.
//
// Solidity: function rwalkAmount() view returns(uint256)
func (_CosmicGame *CosmicGameSession) RwalkAmount() (*big.Int, error) {
	return _CosmicGame.Contract.RwalkAmount(&_CosmicGame.CallOpts)
}

// RwalkAmount is a free data retrieval call binding the contract method 0x120791ca.
//
// Solidity: function rwalkAmount() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) RwalkAmount() (*big.Int, error) {
	return _CosmicGame.Contract.RwalkAmount(&_CosmicGame.CallOpts)
}

// RwalkPrizePercentage is a free data retrieval call binding the contract method 0x376f3448.
//
// Solidity: function rwalkPrizePercentage() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) RwalkPrizePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "rwalkPrizePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RwalkPrizePercentage is a free data retrieval call binding the contract method 0x376f3448.
//
// Solidity: function rwalkPrizePercentage() view returns(uint256)
func (_CosmicGame *CosmicGameSession) RwalkPrizePercentage() (*big.Int, error) {
	return _CosmicGame.Contract.RwalkPrizePercentage(&_CosmicGame.CallOpts)
}

// RwalkPrizePercentage is a free data retrieval call binding the contract method 0x376f3448.
//
// Solidity: function rwalkPrizePercentage() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) RwalkPrizePercentage() (*big.Int, error) {
	return _CosmicGame.Contract.RwalkPrizePercentage(&_CosmicGame.CallOpts)
}

// StakingAmount is a free data retrieval call binding the contract method 0x739a3e02.
//
// Solidity: function stakingAmount() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) StakingAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "stakingAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingAmount is a free data retrieval call binding the contract method 0x739a3e02.
//
// Solidity: function stakingAmount() view returns(uint256)
func (_CosmicGame *CosmicGameSession) StakingAmount() (*big.Int, error) {
	return _CosmicGame.Contract.StakingAmount(&_CosmicGame.CallOpts)
}

// StakingAmount is a free data retrieval call binding the contract method 0x739a3e02.
//
// Solidity: function stakingAmount() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) StakingAmount() (*big.Int, error) {
	return _CosmicGame.Contract.StakingAmount(&_CosmicGame.CallOpts)
}

// StakingPercentage is a free data retrieval call binding the contract method 0x81e1ccba.
//
// Solidity: function stakingPercentage() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) StakingPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "stakingPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingPercentage is a free data retrieval call binding the contract method 0x81e1ccba.
//
// Solidity: function stakingPercentage() view returns(uint256)
func (_CosmicGame *CosmicGameSession) StakingPercentage() (*big.Int, error) {
	return _CosmicGame.Contract.StakingPercentage(&_CosmicGame.CallOpts)
}

// StakingPercentage is a free data retrieval call binding the contract method 0x81e1ccba.
//
// Solidity: function stakingPercentage() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) StakingPercentage() (*big.Int, error) {
	return _CosmicGame.Contract.StakingPercentage(&_CosmicGame.CallOpts)
}

// StakingWallet is a free data retrieval call binding the contract method 0x06ee6ad8.
//
// Solidity: function stakingWallet() view returns(address)
func (_CosmicGame *CosmicGameCaller) StakingWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "stakingWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingWallet is a free data retrieval call binding the contract method 0x06ee6ad8.
//
// Solidity: function stakingWallet() view returns(address)
func (_CosmicGame *CosmicGameSession) StakingWallet() (common.Address, error) {
	return _CosmicGame.Contract.StakingWallet(&_CosmicGame.CallOpts)
}

// StakingWallet is a free data retrieval call binding the contract method 0x06ee6ad8.
//
// Solidity: function stakingWallet() view returns(address)
func (_CosmicGame *CosmicGameCallerSession) StakingWallet() (common.Address, error) {
	return _CosmicGame.Contract.StakingWallet(&_CosmicGame.CallOpts)
}

// StartingBidPriceCST is a free data retrieval call binding the contract method 0x67f3e4a5.
//
// Solidity: function startingBidPriceCST() view returns(uint256)
func (_CosmicGame *CosmicGameCaller) StartingBidPriceCST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicGame.contract.Call(opts, &out, "startingBidPriceCST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingBidPriceCST is a free data retrieval call binding the contract method 0x67f3e4a5.
//
// Solidity: function startingBidPriceCST() view returns(uint256)
func (_CosmicGame *CosmicGameSession) StartingBidPriceCST() (*big.Int, error) {
	return _CosmicGame.Contract.StartingBidPriceCST(&_CosmicGame.CallOpts)
}

// StartingBidPriceCST is a free data retrieval call binding the contract method 0x67f3e4a5.
//
// Solidity: function startingBidPriceCST() view returns(uint256)
func (_CosmicGame *CosmicGameCallerSession) StartingBidPriceCST() (*big.Int, error) {
	return _CosmicGame.Contract.StartingBidPriceCST(&_CosmicGame.CallOpts)
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

// Bid is a paid mutator transaction binding the contract method 0x5d45c962.
//
// Solidity: function bid(string message, int256 randomWalkNFTId) payable returns()
func (_CosmicGame *CosmicGameTransactor) Bid(opts *bind.TransactOpts, message string, randomWalkNFTId *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "bid", message, randomWalkNFTId)
}

// Bid is a paid mutator transaction binding the contract method 0x5d45c962.
//
// Solidity: function bid(string message, int256 randomWalkNFTId) payable returns()
func (_CosmicGame *CosmicGameSession) Bid(message string, randomWalkNFTId *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.Bid(&_CosmicGame.TransactOpts, message, randomWalkNFTId)
}

// Bid is a paid mutator transaction binding the contract method 0x5d45c962.
//
// Solidity: function bid(string message, int256 randomWalkNFTId) payable returns()
func (_CosmicGame *CosmicGameTransactorSession) Bid(message string, randomWalkNFTId *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.Bid(&_CosmicGame.TransactOpts, message, randomWalkNFTId)
}

// BidAndDonateNFT is a paid mutator transaction binding the contract method 0xde5af619.
//
// Solidity: function bidAndDonateNFT(string message, int256 randomWalkNFTId, address nftAddress, uint256 tokenId) payable returns()
func (_CosmicGame *CosmicGameTransactor) BidAndDonateNFT(opts *bind.TransactOpts, message string, randomWalkNFTId *big.Int, nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "bidAndDonateNFT", message, randomWalkNFTId, nftAddress, tokenId)
}

// BidAndDonateNFT is a paid mutator transaction binding the contract method 0xde5af619.
//
// Solidity: function bidAndDonateNFT(string message, int256 randomWalkNFTId, address nftAddress, uint256 tokenId) payable returns()
func (_CosmicGame *CosmicGameSession) BidAndDonateNFT(message string, randomWalkNFTId *big.Int, nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.BidAndDonateNFT(&_CosmicGame.TransactOpts, message, randomWalkNFTId, nftAddress, tokenId)
}

// BidAndDonateNFT is a paid mutator transaction binding the contract method 0xde5af619.
//
// Solidity: function bidAndDonateNFT(string message, int256 randomWalkNFTId, address nftAddress, uint256 tokenId) payable returns()
func (_CosmicGame *CosmicGameTransactorSession) BidAndDonateNFT(message string, randomWalkNFTId *big.Int, nftAddress common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.BidAndDonateNFT(&_CosmicGame.TransactOpts, message, randomWalkNFTId, nftAddress, tokenId)
}

// BidWithCST is a paid mutator transaction binding the contract method 0xba35b1b0.
//
// Solidity: function bidWithCST(string message) returns()
func (_CosmicGame *CosmicGameTransactor) BidWithCST(opts *bind.TransactOpts, message string) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "bidWithCST", message)
}

// BidWithCST is a paid mutator transaction binding the contract method 0xba35b1b0.
//
// Solidity: function bidWithCST(string message) returns()
func (_CosmicGame *CosmicGameSession) BidWithCST(message string) (*types.Transaction, error) {
	return _CosmicGame.Contract.BidWithCST(&_CosmicGame.TransactOpts, message)
}

// BidWithCST is a paid mutator transaction binding the contract method 0xba35b1b0.
//
// Solidity: function bidWithCST(string message) returns()
func (_CosmicGame *CosmicGameTransactorSession) BidWithCST(message string) (*types.Transaction, error) {
	return _CosmicGame.Contract.BidWithCST(&_CosmicGame.TransactOpts, message)
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

// SetETHToCSTBidRatio is a paid mutator transaction binding the contract method 0x97441f2e.
//
// Solidity: function setETHToCSTBidRatio(uint256 newETHToCSTBidRatio) returns()
func (_CosmicGame *CosmicGameTransactor) SetETHToCSTBidRatio(opts *bind.TransactOpts, newETHToCSTBidRatio *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setETHToCSTBidRatio", newETHToCSTBidRatio)
}

// SetETHToCSTBidRatio is a paid mutator transaction binding the contract method 0x97441f2e.
//
// Solidity: function setETHToCSTBidRatio(uint256 newETHToCSTBidRatio) returns()
func (_CosmicGame *CosmicGameSession) SetETHToCSTBidRatio(newETHToCSTBidRatio *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetETHToCSTBidRatio(&_CosmicGame.TransactOpts, newETHToCSTBidRatio)
}

// SetETHToCSTBidRatio is a paid mutator transaction binding the contract method 0x97441f2e.
//
// Solidity: function setETHToCSTBidRatio(uint256 newETHToCSTBidRatio) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetETHToCSTBidRatio(newETHToCSTBidRatio *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetETHToCSTBidRatio(&_CosmicGame.TransactOpts, newETHToCSTBidRatio)
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

// SetMarketingWallet is a paid mutator transaction binding the contract method 0x5d098b38.
//
// Solidity: function setMarketingWallet(address addr) returns()
func (_CosmicGame *CosmicGameTransactor) SetMarketingWallet(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setMarketingWallet", addr)
}

// SetMarketingWallet is a paid mutator transaction binding the contract method 0x5d098b38.
//
// Solidity: function setMarketingWallet(address addr) returns()
func (_CosmicGame *CosmicGameSession) SetMarketingWallet(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetMarketingWallet(&_CosmicGame.TransactOpts, addr)
}

// SetMarketingWallet is a paid mutator transaction binding the contract method 0x5d098b38.
//
// Solidity: function setMarketingWallet(address addr) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetMarketingWallet(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetMarketingWallet(&_CosmicGame.TransactOpts, addr)
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

// SetPrizePercentage is a paid mutator transaction binding the contract method 0xdbf8a28a.
//
// Solidity: function setPrizePercentage(uint256 newPrizePercentage) returns()
func (_CosmicGame *CosmicGameTransactor) SetPrizePercentage(opts *bind.TransactOpts, newPrizePercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setPrizePercentage", newPrizePercentage)
}

// SetPrizePercentage is a paid mutator transaction binding the contract method 0xdbf8a28a.
//
// Solidity: function setPrizePercentage(uint256 newPrizePercentage) returns()
func (_CosmicGame *CosmicGameSession) SetPrizePercentage(newPrizePercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetPrizePercentage(&_CosmicGame.TransactOpts, newPrizePercentage)
}

// SetPrizePercentage is a paid mutator transaction binding the contract method 0xdbf8a28a.
//
// Solidity: function setPrizePercentage(uint256 newPrizePercentage) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetPrizePercentage(newPrizePercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetPrizePercentage(&_CosmicGame.TransactOpts, newPrizePercentage)
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

// SetRwalkPrizePercentage is a paid mutator transaction binding the contract method 0xe40c53d2.
//
// Solidity: function setRwalkPrizePercentage(uint256 newRwalkPrizePercentage) returns()
func (_CosmicGame *CosmicGameTransactor) SetRwalkPrizePercentage(opts *bind.TransactOpts, newRwalkPrizePercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setRwalkPrizePercentage", newRwalkPrizePercentage)
}

// SetRwalkPrizePercentage is a paid mutator transaction binding the contract method 0xe40c53d2.
//
// Solidity: function setRwalkPrizePercentage(uint256 newRwalkPrizePercentage) returns()
func (_CosmicGame *CosmicGameSession) SetRwalkPrizePercentage(newRwalkPrizePercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetRwalkPrizePercentage(&_CosmicGame.TransactOpts, newRwalkPrizePercentage)
}

// SetRwalkPrizePercentage is a paid mutator transaction binding the contract method 0xe40c53d2.
//
// Solidity: function setRwalkPrizePercentage(uint256 newRwalkPrizePercentage) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetRwalkPrizePercentage(newRwalkPrizePercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetRwalkPrizePercentage(&_CosmicGame.TransactOpts, newRwalkPrizePercentage)
}

// SetStakingPercentage is a paid mutator transaction binding the contract method 0x0e38c32b.
//
// Solidity: function setStakingPercentage(uint256 newStakingPercentage) returns()
func (_CosmicGame *CosmicGameTransactor) SetStakingPercentage(opts *bind.TransactOpts, newStakingPercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setStakingPercentage", newStakingPercentage)
}

// SetStakingPercentage is a paid mutator transaction binding the contract method 0x0e38c32b.
//
// Solidity: function setStakingPercentage(uint256 newStakingPercentage) returns()
func (_CosmicGame *CosmicGameSession) SetStakingPercentage(newStakingPercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetStakingPercentage(&_CosmicGame.TransactOpts, newStakingPercentage)
}

// SetStakingPercentage is a paid mutator transaction binding the contract method 0x0e38c32b.
//
// Solidity: function setStakingPercentage(uint256 newStakingPercentage) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetStakingPercentage(newStakingPercentage *big.Int) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetStakingPercentage(&_CosmicGame.TransactOpts, newStakingPercentage)
}

// SetStakingWallet is a paid mutator transaction binding the contract method 0x1a860c3e.
//
// Solidity: function setStakingWallet(address addr) returns()
func (_CosmicGame *CosmicGameTransactor) SetStakingWallet(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.contract.Transact(opts, "setStakingWallet", addr)
}

// SetStakingWallet is a paid mutator transaction binding the contract method 0x1a860c3e.
//
// Solidity: function setStakingWallet(address addr) returns()
func (_CosmicGame *CosmicGameSession) SetStakingWallet(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetStakingWallet(&_CosmicGame.TransactOpts, addr)
}

// SetStakingWallet is a paid mutator transaction binding the contract method 0x1a860c3e.
//
// Solidity: function setStakingWallet(address addr) returns()
func (_CosmicGame *CosmicGameTransactorSession) SetStakingWallet(addr common.Address) (*types.Transaction, error) {
	return _CosmicGame.Contract.SetStakingWallet(&_CosmicGame.TransactOpts, addr)
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
	NumCSTTokens    *big.Int
	PrizeTime       *big.Int
	Message         string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBidEvent is a free log retrieval operation binding the contract event 0x3ebe28e9be13fedb71392c114461386e763acb563218b28db3690553055cd5a5.
//
// Solidity: event BidEvent(address indexed lastBidder, uint256 indexed round, int256 bidPrice, int256 randomWalkNFTId, int256 numCSTTokens, uint256 prizeTime, string message)
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

// WatchBidEvent is a free log subscription operation binding the contract event 0x3ebe28e9be13fedb71392c114461386e763acb563218b28db3690553055cd5a5.
//
// Solidity: event BidEvent(address indexed lastBidder, uint256 indexed round, int256 bidPrice, int256 randomWalkNFTId, int256 numCSTTokens, uint256 prizeTime, string message)
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

// ParseBidEvent is a log parse operation binding the contract event 0x3ebe28e9be13fedb71392c114461386e763acb563218b28db3690553055cd5a5.
//
// Solidity: event BidEvent(address indexed lastBidder, uint256 indexed round, int256 bidPrice, int256 randomWalkNFTId, int256 numCSTTokens, uint256 prizeTime, string message)
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

// CosmicGameETHToCSTBidRatioChangedIterator is returned from FilterETHToCSTBidRatioChanged and is used to iterate over the raw logs and unpacked data for ETHToCSTBidRatioChanged events raised by the CosmicGame contract.
type CosmicGameETHToCSTBidRatioChangedIterator struct {
	Event *CosmicGameETHToCSTBidRatioChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameETHToCSTBidRatioChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameETHToCSTBidRatioChanged)
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
		it.Event = new(CosmicGameETHToCSTBidRatioChanged)
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
func (it *CosmicGameETHToCSTBidRatioChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameETHToCSTBidRatioChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameETHToCSTBidRatioChanged represents a ETHToCSTBidRatioChanged event raised by the CosmicGame contract.
type CosmicGameETHToCSTBidRatioChanged struct {
	NewETHToCSTBidRatio *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterETHToCSTBidRatioChanged is a free log retrieval operation binding the contract event 0x22d4bc2aefb73ef04b923582d7ff2f2b74e699aaf7ff92ea24b348ed59161292.
//
// Solidity: event ETHToCSTBidRatioChanged(uint256 newETHToCSTBidRatio)
func (_CosmicGame *CosmicGameFilterer) FilterETHToCSTBidRatioChanged(opts *bind.FilterOpts) (*CosmicGameETHToCSTBidRatioChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "ETHToCSTBidRatioChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameETHToCSTBidRatioChangedIterator{contract: _CosmicGame.contract, event: "ETHToCSTBidRatioChanged", logs: logs, sub: sub}, nil
}

// WatchETHToCSTBidRatioChanged is a free log subscription operation binding the contract event 0x22d4bc2aefb73ef04b923582d7ff2f2b74e699aaf7ff92ea24b348ed59161292.
//
// Solidity: event ETHToCSTBidRatioChanged(uint256 newETHToCSTBidRatio)
func (_CosmicGame *CosmicGameFilterer) WatchETHToCSTBidRatioChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameETHToCSTBidRatioChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "ETHToCSTBidRatioChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameETHToCSTBidRatioChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "ETHToCSTBidRatioChanged", log); err != nil {
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

// ParseETHToCSTBidRatioChanged is a log parse operation binding the contract event 0x22d4bc2aefb73ef04b923582d7ff2f2b74e699aaf7ff92ea24b348ed59161292.
//
// Solidity: event ETHToCSTBidRatioChanged(uint256 newETHToCSTBidRatio)
func (_CosmicGame *CosmicGameFilterer) ParseETHToCSTBidRatioChanged(log types.Log) (*CosmicGameETHToCSTBidRatioChanged, error) {
	event := new(CosmicGameETHToCSTBidRatioChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "ETHToCSTBidRatioChanged", log); err != nil {
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

// CosmicGameMarketingWalletAddressChangedIterator is returned from FilterMarketingWalletAddressChanged and is used to iterate over the raw logs and unpacked data for MarketingWalletAddressChanged events raised by the CosmicGame contract.
type CosmicGameMarketingWalletAddressChangedIterator struct {
	Event *CosmicGameMarketingWalletAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameMarketingWalletAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameMarketingWalletAddressChanged)
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
		it.Event = new(CosmicGameMarketingWalletAddressChanged)
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
func (it *CosmicGameMarketingWalletAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameMarketingWalletAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameMarketingWalletAddressChanged represents a MarketingWalletAddressChanged event raised by the CosmicGame contract.
type CosmicGameMarketingWalletAddressChanged struct {
	NewMarketingWallet common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMarketingWalletAddressChanged is a free log retrieval operation binding the contract event 0x4d03942c29c20d1bccfe551e9d148c917c5a44fb558a4fc60270d8f76fb75f54.
//
// Solidity: event MarketingWalletAddressChanged(address newMarketingWallet)
func (_CosmicGame *CosmicGameFilterer) FilterMarketingWalletAddressChanged(opts *bind.FilterOpts) (*CosmicGameMarketingWalletAddressChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "MarketingWalletAddressChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameMarketingWalletAddressChangedIterator{contract: _CosmicGame.contract, event: "MarketingWalletAddressChanged", logs: logs, sub: sub}, nil
}

// WatchMarketingWalletAddressChanged is a free log subscription operation binding the contract event 0x4d03942c29c20d1bccfe551e9d148c917c5a44fb558a4fc60270d8f76fb75f54.
//
// Solidity: event MarketingWalletAddressChanged(address newMarketingWallet)
func (_CosmicGame *CosmicGameFilterer) WatchMarketingWalletAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameMarketingWalletAddressChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "MarketingWalletAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameMarketingWalletAddressChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "MarketingWalletAddressChanged", log); err != nil {
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

// ParseMarketingWalletAddressChanged is a log parse operation binding the contract event 0x4d03942c29c20d1bccfe551e9d148c917c5a44fb558a4fc60270d8f76fb75f54.
//
// Solidity: event MarketingWalletAddressChanged(address newMarketingWallet)
func (_CosmicGame *CosmicGameFilterer) ParseMarketingWalletAddressChanged(log types.Log) (*CosmicGameMarketingWalletAddressChanged, error) {
	event := new(CosmicGameMarketingWalletAddressChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "MarketingWalletAddressChanged", log); err != nil {
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

// CosmicGameRwalkPrizePercentageChangedIterator is returned from FilterRwalkPrizePercentageChanged and is used to iterate over the raw logs and unpacked data for RwalkPrizePercentageChanged events raised by the CosmicGame contract.
type CosmicGameRwalkPrizePercentageChangedIterator struct {
	Event *CosmicGameRwalkPrizePercentageChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameRwalkPrizePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameRwalkPrizePercentageChanged)
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
		it.Event = new(CosmicGameRwalkPrizePercentageChanged)
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
func (it *CosmicGameRwalkPrizePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameRwalkPrizePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameRwalkPrizePercentageChanged represents a RwalkPrizePercentageChanged event raised by the CosmicGame contract.
type CosmicGameRwalkPrizePercentageChanged struct {
	NewRwalkPrizePercentage *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterRwalkPrizePercentageChanged is a free log retrieval operation binding the contract event 0xf6a623223b73e9ec7d5b192624a4999b2fa4110d359f6e192ef38e1a1541ee87.
//
// Solidity: event RwalkPrizePercentageChanged(uint256 newRwalkPrizePercentage)
func (_CosmicGame *CosmicGameFilterer) FilterRwalkPrizePercentageChanged(opts *bind.FilterOpts) (*CosmicGameRwalkPrizePercentageChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "RwalkPrizePercentageChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameRwalkPrizePercentageChangedIterator{contract: _CosmicGame.contract, event: "RwalkPrizePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchRwalkPrizePercentageChanged is a free log subscription operation binding the contract event 0xf6a623223b73e9ec7d5b192624a4999b2fa4110d359f6e192ef38e1a1541ee87.
//
// Solidity: event RwalkPrizePercentageChanged(uint256 newRwalkPrizePercentage)
func (_CosmicGame *CosmicGameFilterer) WatchRwalkPrizePercentageChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameRwalkPrizePercentageChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "RwalkPrizePercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameRwalkPrizePercentageChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "RwalkPrizePercentageChanged", log); err != nil {
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

// ParseRwalkPrizePercentageChanged is a log parse operation binding the contract event 0xf6a623223b73e9ec7d5b192624a4999b2fa4110d359f6e192ef38e1a1541ee87.
//
// Solidity: event RwalkPrizePercentageChanged(uint256 newRwalkPrizePercentage)
func (_CosmicGame *CosmicGameFilterer) ParseRwalkPrizePercentageChanged(log types.Log) (*CosmicGameRwalkPrizePercentageChanged, error) {
	event := new(CosmicGameRwalkPrizePercentageChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "RwalkPrizePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameStakingPercentageChangedIterator is returned from FilterStakingPercentageChanged and is used to iterate over the raw logs and unpacked data for StakingPercentageChanged events raised by the CosmicGame contract.
type CosmicGameStakingPercentageChangedIterator struct {
	Event *CosmicGameStakingPercentageChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameStakingPercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameStakingPercentageChanged)
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
		it.Event = new(CosmicGameStakingPercentageChanged)
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
func (it *CosmicGameStakingPercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameStakingPercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameStakingPercentageChanged represents a StakingPercentageChanged event raised by the CosmicGame contract.
type CosmicGameStakingPercentageChanged struct {
	NewStakingPercentage *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterStakingPercentageChanged is a free log retrieval operation binding the contract event 0x9be9203c5d81ee2019998f7020cf412c92a35ba870318f660d4972660210fb4b.
//
// Solidity: event StakingPercentageChanged(uint256 newStakingPercentage)
func (_CosmicGame *CosmicGameFilterer) FilterStakingPercentageChanged(opts *bind.FilterOpts) (*CosmicGameStakingPercentageChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "StakingPercentageChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameStakingPercentageChangedIterator{contract: _CosmicGame.contract, event: "StakingPercentageChanged", logs: logs, sub: sub}, nil
}

// WatchStakingPercentageChanged is a free log subscription operation binding the contract event 0x9be9203c5d81ee2019998f7020cf412c92a35ba870318f660d4972660210fb4b.
//
// Solidity: event StakingPercentageChanged(uint256 newStakingPercentage)
func (_CosmicGame *CosmicGameFilterer) WatchStakingPercentageChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameStakingPercentageChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "StakingPercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameStakingPercentageChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "StakingPercentageChanged", log); err != nil {
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

// ParseStakingPercentageChanged is a log parse operation binding the contract event 0x9be9203c5d81ee2019998f7020cf412c92a35ba870318f660d4972660210fb4b.
//
// Solidity: event StakingPercentageChanged(uint256 newStakingPercentage)
func (_CosmicGame *CosmicGameFilterer) ParseStakingPercentageChanged(log types.Log) (*CosmicGameStakingPercentageChanged, error) {
	event := new(CosmicGameStakingPercentageChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "StakingPercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicGameStakingWalletAddressChangedIterator is returned from FilterStakingWalletAddressChanged and is used to iterate over the raw logs and unpacked data for StakingWalletAddressChanged events raised by the CosmicGame contract.
type CosmicGameStakingWalletAddressChangedIterator struct {
	Event *CosmicGameStakingWalletAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicGameStakingWalletAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicGameStakingWalletAddressChanged)
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
		it.Event = new(CosmicGameStakingWalletAddressChanged)
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
func (it *CosmicGameStakingWalletAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicGameStakingWalletAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicGameStakingWalletAddressChanged represents a StakingWalletAddressChanged event raised by the CosmicGame contract.
type CosmicGameStakingWalletAddressChanged struct {
	NewStakingWallet common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterStakingWalletAddressChanged is a free log retrieval operation binding the contract event 0x3d112e567ad7f87ef5e5219a98118d33b03b247b007cfbadf4f133e7010f2c34.
//
// Solidity: event StakingWalletAddressChanged(address newStakingWallet)
func (_CosmicGame *CosmicGameFilterer) FilterStakingWalletAddressChanged(opts *bind.FilterOpts) (*CosmicGameStakingWalletAddressChangedIterator, error) {

	logs, sub, err := _CosmicGame.contract.FilterLogs(opts, "StakingWalletAddressChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicGameStakingWalletAddressChangedIterator{contract: _CosmicGame.contract, event: "StakingWalletAddressChanged", logs: logs, sub: sub}, nil
}

// WatchStakingWalletAddressChanged is a free log subscription operation binding the contract event 0x3d112e567ad7f87ef5e5219a98118d33b03b247b007cfbadf4f133e7010f2c34.
//
// Solidity: event StakingWalletAddressChanged(address newStakingWallet)
func (_CosmicGame *CosmicGameFilterer) WatchStakingWalletAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicGameStakingWalletAddressChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicGame.contract.WatchLogs(opts, "StakingWalletAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicGameStakingWalletAddressChanged)
				if err := _CosmicGame.contract.UnpackLog(event, "StakingWalletAddressChanged", log); err != nil {
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

// ParseStakingWalletAddressChanged is a log parse operation binding the contract event 0x3d112e567ad7f87ef5e5219a98118d33b03b247b007cfbadf4f133e7010f2c34.
//
// Solidity: event StakingWalletAddressChanged(address newStakingWallet)
func (_CosmicGame *CosmicGameFilterer) ParseStakingWalletAddressChanged(log types.Log) (*CosmicGameStakingWalletAddressChanged, error) {
	event := new(CosmicGameStakingWalletAddressChanged)
	if err := _CosmicGame.contract.UnpackLog(event, "StakingWalletAddressChanged", log); err != nil {
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
	parsed, err := CountersMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := ECDSAMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := EIP712MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := ERC165MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// ERC20BurnableMetaData contains all meta data concerning the ERC20Burnable contract.
var ERC20BurnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC20BurnableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20BurnableMetaData.ABI instead.
var ERC20BurnableABI = ERC20BurnableMetaData.ABI

// ERC20Burnable is an auto generated Go binding around an Ethereum contract.
type ERC20Burnable struct {
	ERC20BurnableCaller     // Read-only binding to the contract
	ERC20BurnableTransactor // Write-only binding to the contract
	ERC20BurnableFilterer   // Log filterer for contract events
}

// ERC20BurnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BurnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BurnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BurnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BurnableSession struct {
	Contract     *ERC20Burnable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BurnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BurnableCallerSession struct {
	Contract *ERC20BurnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC20BurnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BurnableTransactorSession struct {
	Contract     *ERC20BurnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20BurnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BurnableRaw struct {
	Contract *ERC20Burnable // Generic contract binding to access the raw methods on
}

// ERC20BurnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BurnableCallerRaw struct {
	Contract *ERC20BurnableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BurnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BurnableTransactorRaw struct {
	Contract *ERC20BurnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Burnable creates a new instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20Burnable(address common.Address, backend bind.ContractBackend) (*ERC20Burnable, error) {
	contract, err := bindERC20Burnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Burnable{ERC20BurnableCaller: ERC20BurnableCaller{contract: contract}, ERC20BurnableTransactor: ERC20BurnableTransactor{contract: contract}, ERC20BurnableFilterer: ERC20BurnableFilterer{contract: contract}}, nil
}

// NewERC20BurnableCaller creates a new read-only instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableCaller(address common.Address, caller bind.ContractCaller) (*ERC20BurnableCaller, error) {
	contract, err := bindERC20Burnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableCaller{contract: contract}, nil
}

// NewERC20BurnableTransactor creates a new write-only instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BurnableTransactor, error) {
	contract, err := bindERC20Burnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableTransactor{contract: contract}, nil
}

// NewERC20BurnableFilterer creates a new log filterer instance of ERC20Burnable, bound to a specific deployed contract.
func NewERC20BurnableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BurnableFilterer, error) {
	contract, err := bindERC20Burnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableFilterer{contract: contract}, nil
}

// bindERC20Burnable binds a generic wrapper to an already deployed contract.
func bindERC20Burnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20BurnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Burnable *ERC20BurnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Burnable.Contract.ERC20BurnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Burnable *ERC20BurnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.ERC20BurnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Burnable *ERC20BurnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.ERC20BurnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Burnable *ERC20BurnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Burnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Burnable *ERC20BurnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Burnable *ERC20BurnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.Allowance(&_ERC20Burnable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.Allowance(&_ERC20Burnable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.BalanceOf(&_ERC20Burnable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Burnable.Contract.BalanceOf(&_ERC20Burnable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableSession) Decimals() (uint8, error) {
	return _ERC20Burnable.Contract.Decimals(&_ERC20Burnable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Burnable *ERC20BurnableCallerSession) Decimals() (uint8, error) {
	return _ERC20Burnable.Contract.Decimals(&_ERC20Burnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableSession) Name() (string, error) {
	return _ERC20Burnable.Contract.Name(&_ERC20Burnable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Burnable *ERC20BurnableCallerSession) Name() (string, error) {
	return _ERC20Burnable.Contract.Name(&_ERC20Burnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableSession) Symbol() (string, error) {
	return _ERC20Burnable.Contract.Symbol(&_ERC20Burnable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Burnable *ERC20BurnableCallerSession) Symbol() (string, error) {
	return _ERC20Burnable.Contract.Symbol(&_ERC20Burnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Burnable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableSession) TotalSupply() (*big.Int, error) {
	return _ERC20Burnable.Contract.TotalSupply(&_ERC20Burnable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Burnable *ERC20BurnableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Burnable.Contract.TotalSupply(&_ERC20Burnable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Approve(&_ERC20Burnable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Approve(&_ERC20Burnable.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Burn(&_ERC20Burnable.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Burn(&_ERC20Burnable.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.BurnFrom(&_ERC20Burnable.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20Burnable *ERC20BurnableTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.BurnFrom(&_ERC20Burnable.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.DecreaseAllowance(&_ERC20Burnable.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.DecreaseAllowance(&_ERC20Burnable.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.IncreaseAllowance(&_ERC20Burnable.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.IncreaseAllowance(&_ERC20Burnable.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Transfer(&_ERC20Burnable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.Transfer(&_ERC20Burnable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.TransferFrom(&_ERC20Burnable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Burnable *ERC20BurnableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Burnable.Contract.TransferFrom(&_ERC20Burnable.TransactOpts, sender, recipient, amount)
}

// ERC20BurnableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Burnable contract.
type ERC20BurnableApprovalIterator struct {
	Event *ERC20BurnableApproval // Event containing the contract specifics and raw log

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
func (it *ERC20BurnableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableApproval)
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
		it.Event = new(ERC20BurnableApproval)
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
func (it *ERC20BurnableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableApproval represents a Approval event raised by the ERC20Burnable contract.
type ERC20BurnableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20BurnableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableApprovalIterator{contract: _ERC20Burnable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20BurnableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableApproval)
				if err := _ERC20Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20Burnable *ERC20BurnableFilterer) ParseApproval(log types.Log) (*ERC20BurnableApproval, error) {
	event := new(ERC20BurnableApproval)
	if err := _ERC20Burnable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BurnableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Burnable contract.
type ERC20BurnableTransferIterator struct {
	Event *ERC20BurnableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BurnableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableTransfer)
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
		it.Event = new(ERC20BurnableTransfer)
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
func (it *ERC20BurnableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableTransfer represents a Transfer event raised by the ERC20Burnable contract.
type ERC20BurnableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BurnableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Burnable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableTransferIterator{contract: _ERC20Burnable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Burnable *ERC20BurnableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BurnableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Burnable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableTransfer)
				if err := _ERC20Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20Burnable *ERC20BurnableFilterer) ParseTransfer(log types.Log) (*ERC20BurnableTransfer, error) {
	event := new(ERC20BurnableTransfer)
	if err := _ERC20Burnable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PermitMetaData contains all meta data concerning the ERC20Permit contract.
var ERC20PermitMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC20PermitABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20PermitMetaData.ABI instead.
var ERC20PermitABI = ERC20PermitMetaData.ABI

// ERC20Permit is an auto generated Go binding around an Ethereum contract.
type ERC20Permit struct {
	ERC20PermitCaller     // Read-only binding to the contract
	ERC20PermitTransactor // Write-only binding to the contract
	ERC20PermitFilterer   // Log filterer for contract events
}

// ERC20PermitCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20PermitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PermitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20PermitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PermitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20PermitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20PermitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20PermitSession struct {
	Contract     *ERC20Permit      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20PermitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20PermitCallerSession struct {
	Contract *ERC20PermitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC20PermitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20PermitTransactorSession struct {
	Contract     *ERC20PermitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC20PermitRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20PermitRaw struct {
	Contract *ERC20Permit // Generic contract binding to access the raw methods on
}

// ERC20PermitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20PermitCallerRaw struct {
	Contract *ERC20PermitCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20PermitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20PermitTransactorRaw struct {
	Contract *ERC20PermitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Permit creates a new instance of ERC20Permit, bound to a specific deployed contract.
func NewERC20Permit(address common.Address, backend bind.ContractBackend) (*ERC20Permit, error) {
	contract, err := bindERC20Permit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Permit{ERC20PermitCaller: ERC20PermitCaller{contract: contract}, ERC20PermitTransactor: ERC20PermitTransactor{contract: contract}, ERC20PermitFilterer: ERC20PermitFilterer{contract: contract}}, nil
}

// NewERC20PermitCaller creates a new read-only instance of ERC20Permit, bound to a specific deployed contract.
func NewERC20PermitCaller(address common.Address, caller bind.ContractCaller) (*ERC20PermitCaller, error) {
	contract, err := bindERC20Permit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PermitCaller{contract: contract}, nil
}

// NewERC20PermitTransactor creates a new write-only instance of ERC20Permit, bound to a specific deployed contract.
func NewERC20PermitTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20PermitTransactor, error) {
	contract, err := bindERC20Permit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20PermitTransactor{contract: contract}, nil
}

// NewERC20PermitFilterer creates a new log filterer instance of ERC20Permit, bound to a specific deployed contract.
func NewERC20PermitFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20PermitFilterer, error) {
	contract, err := bindERC20Permit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20PermitFilterer{contract: contract}, nil
}

// bindERC20Permit binds a generic wrapper to an already deployed contract.
func bindERC20Permit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20PermitMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Permit *ERC20PermitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Permit.Contract.ERC20PermitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Permit *ERC20PermitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Permit.Contract.ERC20PermitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Permit *ERC20PermitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Permit.Contract.ERC20PermitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Permit *ERC20PermitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Permit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Permit *ERC20PermitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Permit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Permit *ERC20PermitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Permit.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ERC20Permit *ERC20PermitCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Permit.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ERC20Permit *ERC20PermitSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ERC20Permit.Contract.DOMAINSEPARATOR(&_ERC20Permit.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ERC20Permit *ERC20PermitCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ERC20Permit.Contract.DOMAINSEPARATOR(&_ERC20Permit.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Permit *ERC20PermitCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Permit.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Permit *ERC20PermitSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Permit.Contract.Allowance(&_ERC20Permit.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Permit *ERC20PermitCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Permit.Contract.Allowance(&_ERC20Permit.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Permit *ERC20PermitCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Permit.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Permit *ERC20PermitSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Permit.Contract.BalanceOf(&_ERC20Permit.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Permit *ERC20PermitCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Permit.Contract.BalanceOf(&_ERC20Permit.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Permit *ERC20PermitCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20Permit.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Permit *ERC20PermitSession) Decimals() (uint8, error) {
	return _ERC20Permit.Contract.Decimals(&_ERC20Permit.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Permit *ERC20PermitCallerSession) Decimals() (uint8, error) {
	return _ERC20Permit.Contract.Decimals(&_ERC20Permit.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Permit *ERC20PermitCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Permit.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Permit *ERC20PermitSession) Name() (string, error) {
	return _ERC20Permit.Contract.Name(&_ERC20Permit.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Permit *ERC20PermitCallerSession) Name() (string, error) {
	return _ERC20Permit.Contract.Name(&_ERC20Permit.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ERC20Permit *ERC20PermitCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Permit.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ERC20Permit *ERC20PermitSession) Nonces(owner common.Address) (*big.Int, error) {
	return _ERC20Permit.Contract.Nonces(&_ERC20Permit.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_ERC20Permit *ERC20PermitCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _ERC20Permit.Contract.Nonces(&_ERC20Permit.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Permit *ERC20PermitCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Permit.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Permit *ERC20PermitSession) Symbol() (string, error) {
	return _ERC20Permit.Contract.Symbol(&_ERC20Permit.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Permit *ERC20PermitCallerSession) Symbol() (string, error) {
	return _ERC20Permit.Contract.Symbol(&_ERC20Permit.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Permit *ERC20PermitCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Permit.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Permit *ERC20PermitSession) TotalSupply() (*big.Int, error) {
	return _ERC20Permit.Contract.TotalSupply(&_ERC20Permit.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Permit *ERC20PermitCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Permit.Contract.TotalSupply(&_ERC20Permit.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Permit *ERC20PermitTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Permit *ERC20PermitSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.Approve(&_ERC20Permit.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Permit *ERC20PermitTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.Approve(&_ERC20Permit.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Permit *ERC20PermitTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Permit *ERC20PermitSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.DecreaseAllowance(&_ERC20Permit.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Permit *ERC20PermitTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.DecreaseAllowance(&_ERC20Permit.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Permit *ERC20PermitTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Permit *ERC20PermitSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.IncreaseAllowance(&_ERC20Permit.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Permit *ERC20PermitTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.IncreaseAllowance(&_ERC20Permit.TransactOpts, spender, addedValue)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ERC20Permit *ERC20PermitTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ERC20Permit.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ERC20Permit *ERC20PermitSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ERC20Permit.Contract.Permit(&_ERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ERC20Permit *ERC20PermitTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ERC20Permit.Contract.Permit(&_ERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Permit *ERC20PermitTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Permit *ERC20PermitSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.Transfer(&_ERC20Permit.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Permit *ERC20PermitTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.Transfer(&_ERC20Permit.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Permit *ERC20PermitTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Permit *ERC20PermitSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.TransferFrom(&_ERC20Permit.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Permit *ERC20PermitTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Permit.Contract.TransferFrom(&_ERC20Permit.TransactOpts, sender, recipient, amount)
}

// ERC20PermitApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Permit contract.
type ERC20PermitApprovalIterator struct {
	Event *ERC20PermitApproval // Event containing the contract specifics and raw log

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
func (it *ERC20PermitApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PermitApproval)
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
		it.Event = new(ERC20PermitApproval)
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
func (it *ERC20PermitApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PermitApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PermitApproval represents a Approval event raised by the ERC20Permit contract.
type ERC20PermitApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Permit *ERC20PermitFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20PermitApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Permit.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PermitApprovalIterator{contract: _ERC20Permit.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Permit *ERC20PermitFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20PermitApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Permit.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PermitApproval)
				if err := _ERC20Permit.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20Permit *ERC20PermitFilterer) ParseApproval(log types.Log) (*ERC20PermitApproval, error) {
	event := new(ERC20PermitApproval)
	if err := _ERC20Permit.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PermitTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Permit contract.
type ERC20PermitTransferIterator struct {
	Event *ERC20PermitTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20PermitTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PermitTransfer)
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
		it.Event = new(ERC20PermitTransfer)
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
func (it *ERC20PermitTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PermitTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PermitTransfer represents a Transfer event raised by the ERC20Permit contract.
type ERC20PermitTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Permit *ERC20PermitFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20PermitTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Permit.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PermitTransferIterator{contract: _ERC20Permit.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Permit *ERC20PermitFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20PermitTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Permit.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PermitTransfer)
				if err := _ERC20Permit.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20Permit *ERC20PermitFilterer) ParseTransfer(log types.Log) (*ERC20PermitTransfer, error) {
	event := new(ERC20PermitTransfer)
	if err := _ERC20Permit.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721MetaData contains all meta data concerning the ERC721 contract.
var ERC721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002946380380620029468339818101604052810190620000379190620001f6565b8160009081620000489190620004c6565b5080600190816200005a9190620004c6565b505050620005ad565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620000cc8262000081565b810181811067ffffffffffffffff82111715620000ee57620000ed62000092565b5b80604052505050565b60006200010362000063565b9050620001118282620000c1565b919050565b600067ffffffffffffffff82111562000134576200013362000092565b5b6200013f8262000081565b9050602081019050919050565b60005b838110156200016c5780820151818401526020810190506200014f565b60008484015250505050565b60006200018f620001898462000116565b620000f7565b905082815260208101848484011115620001ae57620001ad6200007c565b5b620001bb8482856200014c565b509392505050565b600082601f830112620001db57620001da62000077565b5b8151620001ed84826020860162000178565b91505092915050565b6000806040838503121562000210576200020f6200006d565b5b600083015167ffffffffffffffff81111562000231576200023062000072565b5b6200023f85828601620001c3565b925050602083015167ffffffffffffffff81111562000263576200026262000072565b5b6200027185828601620001c3565b9150509250929050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620002ce57607f821691505b602082108103620002e457620002e362000286565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026200034e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200030f565b6200035a86836200030f565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620003a7620003a16200039b8462000372565b6200037c565b62000372565b9050919050565b6000819050919050565b620003c38362000386565b620003db620003d282620003ae565b8484546200031c565b825550505050565b600090565b620003f2620003e3565b620003ff818484620003b8565b505050565b5b8181101562000427576200041b600082620003e8565b60018101905062000405565b5050565b601f82111562000476576200044081620002ea565b6200044b84620002ff565b810160208510156200045b578190505b620004736200046a85620002ff565b83018262000404565b50505b505050565b600082821c905092915050565b60006200049b600019846008026200047b565b1980831691505092915050565b6000620004b6838362000488565b9150826002028217905092915050565b620004d1826200027b565b67ffffffffffffffff811115620004ed57620004ec62000092565b5b620004f98254620002b5565b620005068282856200042b565b600060209050601f8311600181146200053e576000841562000529578287015190505b620005358582620004a8565b865550620005a5565b601f1984166200054e86620002ea565b60005b82811015620005785784890151825560018201915060208501945060208101905062000551565b8683101562000598578489015162000594601f89168262000488565b8355505b6001600288020188555050505b505050505050565b61238980620005bd6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80636352211e1161008c578063a22cb46511610066578063a22cb46514610224578063b88d4fde14610240578063c87b56dd1461025c578063e985e9c51461028c576100cf565b80636352211e146101a657806370a08231146101d657806395d89b4114610206576100cf565b806301ffc9a7146100d457806306fdde0314610104578063081812fc14610122578063095ea7b31461015257806323b872dd1461016e57806342842e0e1461018a575b600080fd5b6100ee60048036038101906100e99190611411565b6102bc565b6040516100fb9190611459565b60405180910390f35b61010c61039e565b6040516101199190611504565b60405180910390f35b61013c6004803603810190610137919061155c565b610430565b60405161014991906115ca565b60405180910390f35b61016c60048036038101906101679190611611565b6104b5565b005b61018860048036038101906101839190611651565b6105cc565b005b6101a4600480360381019061019f9190611651565b61062c565b005b6101c060048036038101906101bb919061155c565b61064c565b6040516101cd91906115ca565b60405180910390f35b6101f060048036038101906101eb91906116a4565b6106fd565b6040516101fd91906116e0565b60405180910390f35b61020e6107b4565b60405161021b9190611504565b60405180910390f35b61023e60048036038101906102399190611727565b610846565b005b61025a6004803603810190610255919061189c565b6109c6565b005b6102766004803603810190610271919061155c565b610a28565b6040516102839190611504565b60405180910390f35b6102a660048036038101906102a1919061191f565b610acf565b6040516102b39190611459565b60405180910390f35b60007f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061038757507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610397575061039682610b63565b5b9050919050565b6060600080546103ad9061198e565b80601f01602080910402602001604051908101604052809291908181526020018280546103d99061198e565b80156104265780601f106103fb57610100808354040283529160200191610426565b820191906000526020600020905b81548152906001019060200180831161040957829003601f168201915b5050505050905090565b600061043b82610bcd565b61047a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047190611a31565b60405180910390fd5b6004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60006104c08261064c565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610530576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052790611ac3565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1661054f610c39565b73ffffffffffffffffffffffffffffffffffffffff16148061057e575061057d81610578610c39565b610acf565b5b6105bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b490611b55565b60405180910390fd5b6105c78383610c41565b505050565b6105dd6105d7610c39565b82610cfa565b61061c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061390611be7565b60405180910390fd5b610627838383610dd8565b505050565b610647838383604051806020016040528060008152506109c6565b505050565b6000806002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036106f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106eb90611c79565b60405180910390fd5b80915050919050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361076d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076490611d0b565b60405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600180546107c39061198e565b80601f01602080910402602001604051908101604052809291908181526020018280546107ef9061198e565b801561083c5780601f106108115761010080835404028352916020019161083c565b820191906000526020600020905b81548152906001019060200180831161081f57829003601f168201915b5050505050905090565b61084e610c39565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b290611d77565b60405180910390fd5b80600560006108c8610c39565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff16610975610c39565b73ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516109ba9190611459565b60405180910390a35050565b6109d76109d1610c39565b83610cfa565b610a16576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0d90611be7565b60405180910390fd5b610a2284848484611033565b50505050565b6060610a3382610bcd565b610a72576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6990611e09565b60405180910390fd5b6000610a7c61108f565b90506000815111610a9c5760405180602001604052806000815250610ac7565b80610aa6846110a6565b604051602001610ab7929190611e65565b6040516020818303038152906040525b915050919050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b60008073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b600033905090565b816004600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff16610cb48361064c565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000610d0582610bcd565b610d44576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3b90611efb565b60405180910390fd5b6000610d4f8361064c565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480610dbe57508373ffffffffffffffffffffffffffffffffffffffff16610da684610430565b73ffffffffffffffffffffffffffffffffffffffff16145b80610dcf5750610dce8185610acf565b5b91505092915050565b8273ffffffffffffffffffffffffffffffffffffffff16610df88261064c565b73ffffffffffffffffffffffffffffffffffffffff1614610e4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4590611f8d565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610ebd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eb49061201f565b60405180910390fd5b610ec8838383611206565b610ed3600082610c41565b6001600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610f23919061206e565b925050819055506001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610f7a91906120a2565b92505081905550816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b61103e848484610dd8565b61104a8484848461120b565b611089576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161108090612148565b60405180910390fd5b50505050565b606060405180602001604052806000815250905090565b6060600082036110ed576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050611201565b600082905060005b6000821461111f57808061110890612168565b915050600a8261111891906121df565b91506110f5565b60008167ffffffffffffffff81111561113b5761113a611771565b5b6040519080825280601f01601f19166020018201604052801561116d5781602001600182028036833780820191505090505b5090505b600085146111fa57600182611186919061206e565b9150600a856111959190612210565b60306111a191906120a2565b60f81b8183815181106111b7576111b6612241565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a856111f391906121df565b9450611171565b8093505050505b919050565b505050565b600061122c8473ffffffffffffffffffffffffffffffffffffffff16611392565b15611385578373ffffffffffffffffffffffffffffffffffffffff1663150b7a02611255610c39565b8786866040518563ffffffff1660e01b815260040161127794939291906122c5565b6020604051808303816000875af19250505080156112b357506040513d601f19601f820116820180604052508101906112b09190612326565b60015b611335573d80600081146112e3576040519150601f19603f3d011682016040523d82523d6000602084013e6112e8565b606091505b50600081510361132d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161132490612148565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161491505061138a565b600190505b949350505050565b600080823b905060008111915050919050565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6113ee816113b9565b81146113f957600080fd5b50565b60008135905061140b816113e5565b92915050565b600060208284031215611427576114266113af565b5b6000611435848285016113fc565b91505092915050565b60008115159050919050565b6114538161143e565b82525050565b600060208201905061146e600083018461144a565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156114ae578082015181840152602081019050611493565b60008484015250505050565b6000601f19601f8301169050919050565b60006114d682611474565b6114e0818561147f565b93506114f0818560208601611490565b6114f9816114ba565b840191505092915050565b6000602082019050818103600083015261151e81846114cb565b905092915050565b6000819050919050565b61153981611526565b811461154457600080fd5b50565b60008135905061155681611530565b92915050565b600060208284031215611572576115716113af565b5b600061158084828501611547565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006115b482611589565b9050919050565b6115c4816115a9565b82525050565b60006020820190506115df60008301846115bb565b92915050565b6115ee816115a9565b81146115f957600080fd5b50565b60008135905061160b816115e5565b92915050565b60008060408385031215611628576116276113af565b5b6000611636858286016115fc565b925050602061164785828601611547565b9150509250929050565b60008060006060848603121561166a576116696113af565b5b6000611678868287016115fc565b9350506020611689868287016115fc565b925050604061169a86828701611547565b9150509250925092565b6000602082840312156116ba576116b96113af565b5b60006116c8848285016115fc565b91505092915050565b6116da81611526565b82525050565b60006020820190506116f560008301846116d1565b92915050565b6117048161143e565b811461170f57600080fd5b50565b600081359050611721816116fb565b92915050565b6000806040838503121561173e5761173d6113af565b5b600061174c858286016115fc565b925050602061175d85828601611712565b9150509250929050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6117a9826114ba565b810181811067ffffffffffffffff821117156117c8576117c7611771565b5b80604052505050565b60006117db6113a5565b90506117e782826117a0565b919050565b600067ffffffffffffffff82111561180757611806611771565b5b611810826114ba565b9050602081019050919050565b82818337600083830152505050565b600061183f61183a846117ec565b6117d1565b90508281526020810184848401111561185b5761185a61176c565b5b61186684828561181d565b509392505050565b600082601f83011261188357611882611767565b5b813561189384826020860161182c565b91505092915050565b600080600080608085870312156118b6576118b56113af565b5b60006118c4878288016115fc565b94505060206118d5878288016115fc565b93505060406118e687828801611547565b925050606085013567ffffffffffffffff811115611907576119066113b4565b5b6119138782880161186e565b91505092959194509250565b60008060408385031215611936576119356113af565b5b6000611944858286016115fc565b9250506020611955858286016115fc565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806119a657607f821691505b6020821081036119b9576119b861195f565b5b50919050565b7f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860008201527f697374656e7420746f6b656e0000000000000000000000000000000000000000602082015250565b6000611a1b602c8361147f565b9150611a26826119bf565b604082019050919050565b60006020820190508181036000830152611a4a81611a0e565b9050919050565b7f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560008201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b6000611aad60218361147f565b9150611ab882611a51565b604082019050919050565b60006020820190508181036000830152611adc81611aa0565b9050919050565b7f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f7760008201527f6e6572206e6f7220617070726f76656420666f7220616c6c0000000000000000602082015250565b6000611b3f60388361147f565b9150611b4a82611ae3565b604082019050919050565b60006020820190508181036000830152611b6e81611b32565b9050919050565b7f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f60008201527f776e6572206e6f7220617070726f766564000000000000000000000000000000602082015250565b6000611bd160318361147f565b9150611bdc82611b75565b604082019050919050565b60006020820190508181036000830152611c0081611bc4565b9050919050565b7f4552433732313a206f776e657220717565727920666f72206e6f6e657869737460008201527f656e7420746f6b656e0000000000000000000000000000000000000000000000602082015250565b6000611c6360298361147f565b9150611c6e82611c07565b604082019050919050565b60006020820190508181036000830152611c9281611c56565b9050919050565b7f4552433732313a2062616c616e636520717565727920666f7220746865207a6560008201527f726f206164647265737300000000000000000000000000000000000000000000602082015250565b6000611cf5602a8361147f565b9150611d0082611c99565b604082019050919050565b60006020820190508181036000830152611d2481611ce8565b9050919050565b7f4552433732313a20617070726f766520746f2063616c6c657200000000000000600082015250565b6000611d6160198361147f565b9150611d6c82611d2b565b602082019050919050565b60006020820190508181036000830152611d9081611d54565b9050919050565b7f4552433732314d657461646174613a2055524920717565727920666f72206e6f60008201527f6e6578697374656e7420746f6b656e0000000000000000000000000000000000602082015250565b6000611df3602f8361147f565b9150611dfe82611d97565b604082019050919050565b60006020820190508181036000830152611e2281611de6565b9050919050565b600081905092915050565b6000611e3f82611474565b611e498185611e29565b9350611e59818560208601611490565b80840191505092915050565b6000611e718285611e34565b9150611e7d8284611e34565b91508190509392505050565b7f4552433732313a206f70657261746f7220717565727920666f72206e6f6e657860008201527f697374656e7420746f6b656e0000000000000000000000000000000000000000602082015250565b6000611ee5602c8361147f565b9150611ef082611e89565b604082019050919050565b60006020820190508181036000830152611f1481611ed8565b9050919050565b7f4552433732313a207472616e73666572206f6620746f6b656e2074686174206960008201527f73206e6f74206f776e0000000000000000000000000000000000000000000000602082015250565b6000611f7760298361147f565b9150611f8282611f1b565b604082019050919050565b60006020820190508181036000830152611fa681611f6a565b9050919050565b7f4552433732313a207472616e7366657220746f20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b600061200960248361147f565b915061201482611fad565b604082019050919050565b6000602082019050818103600083015261203881611ffc565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061207982611526565b915061208483611526565b925082820390508181111561209c5761209b61203f565b5b92915050565b60006120ad82611526565b91506120b883611526565b92508282019050808211156120d0576120cf61203f565b5b92915050565b7f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560008201527f63656976657220696d706c656d656e7465720000000000000000000000000000602082015250565b600061213260328361147f565b915061213d826120d6565b604082019050919050565b6000602082019050818103600083015261216181612125565b9050919050565b600061217382611526565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036121a5576121a461203f565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006121ea82611526565b91506121f583611526565b925082612205576122046121b0565b5b828204905092915050565b600061221b82611526565b915061222683611526565b925082612236576122356121b0565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081519050919050565b600082825260208201905092915050565b600061229782612270565b6122a1818561227b565b93506122b1818560208601611490565b6122ba816114ba565b840191505092915050565b60006080820190506122da60008301876115bb565b6122e760208301866115bb565b6122f460408301856116d1565b8181036060830152612306818461228c565b905095945050505050565b600081519050612320816113e5565b92915050565b60006020828403121561233c5761233b6113af565b5b600061234a84828501612311565b9150509291505056fea2646970667358221220fb1be3510c119434556b065ac9555a3e4ddcae4e3b6ee61680e47c45ff5b744e64736f6c63430008130033",
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
	parsed, err := ERC721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := ERC721EnumerableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := IERC165MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := IERC721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := IERC721EnumerableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := IERC721MetadataMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := IERC721ReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := MathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := OwnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := SafeCastMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := StringsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
