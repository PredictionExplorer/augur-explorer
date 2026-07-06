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

// CosmicSignatureGameV2MetaData contains all meta data concerning the CosmicSignatureGameV2 contract.
var CosmicSignatureGameV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit\",\"type\":\"uint256\"}],\"name\":\"BidCstRewardAmountMinLimitNotReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"BidHasBeenPlacedInCurrentRound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"}],\"name\":\"CallerIsNotNftOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientReceivedBidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"InvalidOperationInCurrentState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"lastBidderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"durationUntilOperationIsPermitted\",\"type\":\"uint256\"}],\"name\":\"MainPrizeClaimDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"mainPrizeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"MainPrizeEarlyClaim\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"NoBidsPlacedInCurrentRound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"roundActivationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"RoundIsActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"roundActivationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"RoundIsInactive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"messageLength\",\"type\":\"uint256\"}],\"name\":\"TooLongBidMessage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"randomWalkNftId\",\"type\":\"uint256\"}],\"name\":\"UsedRandomWalkNft\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"WrongBidType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ArbitrumError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"BidCstRewardAmountMultiplierChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"BidMessageLengthMaxLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lastBidderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"paidEthPrice\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"paidCstPrice\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"randomWalkNftId\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstDutchAuctionDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mainPrizeTime\",\"type\":\"uint256\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CharityAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CharityEthDonationAmountPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"ChronoWarriorEthPrizeAmountPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chronoWarriorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeCosmicSignatureNftId\",\"type\":\"uint256\"}],\"name\":\"ChronoWarriorPrizePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICosmicSignatureNft\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CosmicSignatureNftAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICosmicSignatureToken\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CosmicSignatureTokenAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CstDutchAuctionBeginningBidPriceMinLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CstDutchAuctionDurationChangeDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CstDutchAuctionDurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CstPrizeAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"DelayDurationBeforeRoundActivationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"enduranceChampionAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeCosmicSignatureNftId\",\"type\":\"uint256\"}],\"name\":\"EnduranceChampionPrizePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"EthBidPriceIncreaseDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"EthBidRefundAmountInGasToSwallowMaxLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthDonated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ethDonationWithInfoRecordIndex\",\"type\":\"uint256\"}],\"name\":\"EthDonatedWithInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"EthDutchAuctionDurationDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"EthDutchAuctionEndingBidPriceDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"FirstBidPlacedInRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"charityAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsTransferredToCharity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"InitialDurationUntilMainPrizeDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lastCstBidderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeCosmicSignatureNftId\",\"type\":\"uint256\"}],\"name\":\"LastCstBidderPrizePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MainEthPrizeAmountPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeCosmicSignatureNftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeoutTimeToWithdrawSecondaryPrizes\",\"type\":\"uint256\"}],\"name\":\"MainPrizeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MainPrizeTimeIncrementInMicroSecondsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MainPrizeTimeIncrementIncreaseDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"MarketingWalletAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MarketingWalletCstContributionAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"NumRaffleCosmicSignatureNftsForBiddersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"NumRaffleEthPrizesForBiddersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIPrizesWallet\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"PrizesWalletAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RaffleTotalEthPrizeAmountForBiddersPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winnerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethPrizeAmount\",\"type\":\"uint256\"}],\"name\":\"RaffleWinnerBidderEthPrizeAllocated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"winnerIsRandomWalkNftStaker\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winnerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeCosmicSignatureNftId\",\"type\":\"uint256\"}],\"name\":\"RaffleWinnerPrizePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIRandomWalkNFT\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"RandomWalkNftAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RoundActivationTimeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIStakingWalletCosmicSignatureNft\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"StakingWalletCosmicSignatureNftAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIStakingWalletRandomWalkNft\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"StakingWalletRandomWalkNftAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"TimeoutDurationToClaimMainPrizeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bidCstRewardAmountMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bidMessageLengthMaxLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceMaxLimit_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"}],\"name\":\"bidWithCst\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceMaxLimit_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"bidWithCstAndDonateNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceMaxLimit_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"bidWithCstAndDonateToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"randomWalkNftId_\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"}],\"name\":\"bidWithEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"randomWalkNftId_\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"bidWithEthAndDonateNft\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"randomWalkNftId_\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"bidWithEthAndDonateToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"}],\"name\":\"bidderAddresses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"numItems\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidderAddress\",\"type\":\"address\"}],\"name\":\"biddersInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalSpentEthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSpentCstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBidTimeStamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charityAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charityEthDonationAmountPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chronoWarriorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chronoWarriorDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chronoWarriorEthPrizeAmountPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimMainPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cosmicSignatureNftStakingTotalEthRewardAmountPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionBeginningBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionBeginningBidPriceMinLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionBeginningTimeStamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionDurationChangeDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstPrizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayDurationBeforeRoundActivation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donateEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"data_\",\"type\":\"string\"}],\"name\":\"donateEthWithInfo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enduranceChampionAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enduranceChampionDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enduranceChampionStartTimeStamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethBidPriceIncreaseDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethBidRefundAmountInGasToSwallowMaxLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ethDonationWithInfoRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethDutchAuctionBeginningBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethDutchAuctionDurationDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethDutchAuctionEndingBidPriceDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBidCstRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"currentTimeOffset_\",\"type\":\"int256\"}],\"name\":\"getBidCstRewardAmountAdvanced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidIndex_\",\"type\":\"uint256\"}],\"name\":\"getBidderAddressAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidderAddress_\",\"type\":\"address\"}],\"name\":\"getBidderTotalSpentAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCharityEthDonationAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChronoWarriorEthPrizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCosmicSignatureNftStakingTotalEthRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCstDutchAuctionDurations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDurationElapsedSinceRoundActivation\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDurationUntilMainPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDurationUntilMainPrizeRaw\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDurationUntilRoundActivation\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEthDutchAuctionDurations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethBidPrice_\",\"type\":\"uint256\"}],\"name\":\"getEthPlusRandomWalkNftBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitialDurationUntilMainPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMainEthPrizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMainPrizeTimeIncrement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextCstBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"currentTimeOffset_\",\"type\":\"int256\"}],\"name\":\"getNextCstBidPriceAdvanced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextEthBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"currentTimeOffset_\",\"type\":\"int256\"}],\"name\":\"getNextEthBidPriceAdvanced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRaffleTotalEthPrizeAmountForBidders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"}],\"name\":\"getTotalNumBids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"halveEthDutchAuctionEndingBidPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialDurationUntilMainPrizeDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBidderAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCstBidderAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainEthPrizeAmountPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainPrizeTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainPrizeTimeIncrementInMicroSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainPrizeTimeIncrementIncreaseDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketingWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketingWalletCstContributionAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextEthBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRoundFirstCstDutchAuctionBeginningBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractCosmicSignatureNft\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numEthDonationWithInfoRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleCosmicSignatureNftsForBidders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleCosmicSignatureNftsForRandomWalkNftStakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleEthPrizesForBidders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prevEnduranceChampionDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizesWallet\",\"outputs\":[{\"internalType\":\"contractPrizesWallet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raffleTotalEthPrizeAmountForBiddersPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomWalkNft\",\"outputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundActivationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setBidCstRewardAmountMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setBidMessageLengthMaxLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setCharityAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCharityEthDonationAmountPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setChronoWarriorEthPrizeAmountPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICosmicSignatureNft\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setCosmicSignatureNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCosmicSignatureNftStakingTotalEthRewardAmountPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICosmicSignatureToken\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setCosmicSignatureToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCstDutchAuctionBeginningBidPriceMinLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCstDutchAuctionDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCstDutchAuctionDurationChangeDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCstPrizeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setDelayDurationBeforeRoundActivation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setEthBidPriceIncreaseDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setEthBidRefundAmountInGasToSwallowMaxLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setEthDutchAuctionDurationDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setEthDutchAuctionEndingBidPriceDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setInitialDurationUntilMainPrizeDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setMainEthPrizeAmountPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setMainPrizeTimeIncrementInMicroSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setMainPrizeTimeIncrementIncreaseDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setMarketingWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setMarketingWalletCstContributionAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleCosmicSignatureNftsForBidders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleCosmicSignatureNftsForRandomWalkNftStakers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleEthPrizesForBidders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPrizesWallet\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setPrizesWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setRaffleTotalEthPrizeAmountForBiddersPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRandomWalkNFT\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setRandomWalkNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setRoundActivationTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakingWalletCosmicSignatureNft\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setStakingWalletCosmicSignatureNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakingWalletRandomWalkNft\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setStakingWalletRandomWalkNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setTimeoutDurationToClaimMainPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingWalletCosmicSignatureNft\",\"outputs\":[{\"internalType\":\"contractStakingWalletCosmicSignatureNft\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingWalletRandomWalkNft\",\"outputs\":[{\"internalType\":\"contractStakingWalletRandomWalkNft\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeoutDurationToClaimMainPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractCosmicSignatureToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tryGetCurrentChampions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"enduranceChampionAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"enduranceChampionDuration_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"chronoWarriorAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chronoWarriorDuration_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"usedRandomWalkNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nftWasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523461003e5761001161004d565b610019610043565b61b2e761036b82396080518181816194060152818161946b0152619639015261b2e790f35b610049565b60405190565b5f80fd5b61005561005f565b61005d61028e565b565b610067610069565b565b610071610073565b565b61007b61007d565b565b610085610087565b565b61008f610091565b565b61009961009b565b565b6100a36100a5565b565b6100ad6100af565b565b6100b76100b9565b565b6100c16100c3565b565b6100cb6100cd565b565b6100d56100d7565b565b6100df6100e1565b565b6100e96100eb565b565b6100f36100f5565b565b6100fd6100ff565b565b610107610109565b565b610111610113565b565b61011b61011d565b565b610125610127565b565b61012f610131565b565b61013961013b565b565b610143610145565b565b61014d610191565b565b60018060a01b031690565b90565b61017161016c6101769261014f565b61015a565b61014f565b90565b6101829061015d565b90565b61018e90610179565b90565b61019a30610185565b608052565b60401c90565b60ff1690565b6101b76101bc9161019f565b6101a5565b90565b6101c990546101ab565b90565b5f0190565b5f1c90565b60018060401b031690565b6101ed6101f2916101d1565b6101d6565b90565b6101ff90546101e1565b90565b60018060401b031690565b5f1b90565b9061022360018060401b039161020d565b9181191691161790565b61024161023c61024692610202565b61015a565b610202565b90565b90565b9061026161025c6102689261022d565b610249565b8254610212565b9055565b61027590610202565b9052565b919061028c905f6020850194019061026c565b565b610296610346565b6102a15f82016101bf565b61032a576102b05f82016101f5565b6102c86102c260018060401b03610202565b91610202565b036102d1575b50565b6102e4905f60018060401b03910161024c565b60018060401b036103217fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d291610318610043565b91829182610279565b0390a15f6102ce565b5f63f92ee8a960e01b815280610342600482016101cc565b0390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009056fe6080604052600436101561001d575b366135055761001b6167b7565b005b6100275f35610815565b80620ac9f114610810578063040d4d311461080b5780630433847914610806578063096323661461080157806309794bee146107fc5780630a120648146107f75780630b5f95ae146107f25780630c9be46d146107ed5780630eb16be6146107e8578063119b22b3146107e357806311b0d1fe146107de578063135f3d28146107d957806317887731146107d45780631824d5e7146107cf57806318305de2146107ca5780631b410319146107c55780631e9cbb7e146107c05780631f1b4aa4146107bb57806323b31cfc146107b6578063250fadb6146107b15780632665c882146107ac57806327700481146107a757806327995f07146107a25780632afa25801461079d5780632b8dcbba146107985780632b91c7bb146107935780632d809e881461078e5780632d829a2d146107895780632f894cd7146107845780632fb3c48f1461077f578063320c435c1461077a578063329b95a51461077557806336750d2c1461077057806337b99cc71461076b5780633b9d292e146107665780634164b95b14610761578063441b32891461075c578063448c6eb11461075757806344a4b9171461075257806344acc12a1461074d578063477adf2a1461074857806347ccca02146107435780634c2a4a331461073e5780634e452010146107395780634f1ef286146107345780634f7346121461072f5780634feb78b71461072a57806352d1902d14610725578063543f416f1461072057806354ada1d61461071b57806356732241146107165780635863a705146107115780635a1e5bde1461070c5780635b0a45d9146107075780635cd8a76b146107025780635d098b38146106fd5780635f0112fe146106f857806362ed9b53146106f35780636b59acb8146106ee5780636b7cbe85146106e95780636c0613c0146106e45780636c17e3cc146106df5780636e95d286146106da5780636e970834146106d5578063715018a6146106d0578063755b4ef7146106cb57806375f0a874146106c657806377fa1027146106c1578063876d5c36146106bc57806388ce802c146106b75780638da5cb5b146106b2578063928880fa146106ad5780639302020f146106a85780639646d758146106a35780639aa1b38d1461069e5780639e50acc9146106995780639edeaf8e14610694578063a35286d11461068f578063a4be0d401461068a578063a922ab5d14610685578063a974201614610680578063aadd1b031461067b578063ad3cb1cc14610676578063ad4b0e8a14610671578063afcf2fc41461066c578063b30f5bb114610667578063b4f1b13414610662578063b5d1f06f1461065d578063b6a94f4214610658578063b700db5f14610653578063b78d1e2a1461064e578063b9cf9ba514610649578063baab443014610644578063bb4b3e6f1461063f578063be720ad51461063a578063c7e7a60114610635578063c87baab514610630578063cb720d4d1461062b578063cfb4e59914610626578063d1f8fcf214610621578063d7559b9c1461061c578063d9ab9eaa14610617578063da9931dd14610612578063ddd6df071461060d578063de704b4114610608578063dfcd00d114610603578063e2f9185f146105fe578063e5b3cd14146105f9578063eaace302146105f4578063eb13430e146105ef578063ebaa1ea8146105ea578063ebb9bc5c146105e5578063ecb5776e146105e0578063ef22d15b146105db578063efeb248a146105d6578063f0bdab7c146105d1578063f11400f0146105cc578063f2fde38b146105c7578063f34d411c146105c2578063f444b298146105bd578063f49efe9d146105b8578063f7bea078146105b3578063fbaf5084146105ae578063fc0c546a146105a9578063fd77310f146105a4578063fd9b37471461059f5763fdfb9ba40361000e576134d2565b61349b565b613443565b61340e565b613369565b613334565b6132f1565b6132bc565b613279565b613246565b613211565b6131ce565b613199565b613154565b613121565b6130ec565b613087565b613042565b612ffd565b612fb8565b612f73565b612f30565b612efd565b612eca565b612e95565b612e52565b612e1e565b612d96565b612d53565b612d1a565b612ca7565b612c62565b612c1d565b612bd8565b612b93565b612b5e565b612b21565b612a7d565b612a25565b6129ec565b6127aa565b612728565b6126e3565b61269e565b612659565b612557565b612522565b6124dd565b61243b565b612406565b6123d3565b612351565b61230c565b6122c9565b612247565b61220d565b612185565b612152565b612125565b612084565b61204f565b61200a565b611f68565b611f33565b611eee565b611ebb565b611e39565b611df4565b611d80565b611d4b565b611d16565b611ce3565b611cb0565b611c7b565b611c48565b611bc6565b611b81565b611b3c565b611af9565b611ac4565b611a6c565b611a37565b611a0d565b611917565b6118e2565b61189d565b6117f9565b6117b6565b611734565b6116f1565b6116be565b611689565b611644565b6115ff565b6115ca565b611593565b6114c1565b61147c565b61141a565b6113e7565b6113b2565b611310565b6112db565b6112a6565b611202565b6111cd565b611188565b611153565b611119565b611069565b611036565b611001565b610fbc565b610f77565b610f2e565b610dbb565b610d84565b610bc7565b610b55565b610b22565b610aac565b610a77565b6109e1565b6109ae565b61097b565b610923565b6108b9565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b90565b6108398161082d565b0361084057565b5f80fd5b9050359061085182610830565b565b919060408382031261087b578061086f610878925f8601610844565b93602001610844565b90565b610825565b60018060a01b031690565b61089490610880565b90565b6108a09061088b565b9052565b91906108b7905f60208501940190610897565b565b346108ea576108e66108d56108cf366004610853565b90613523565b6108dd61081b565b918291826108a4565b0390f35b610821565b5f9103126108f957565b610825565b90565b61090a906108fe565b9052565b9190610921905f60208501940190610901565b565b34610953576109333660046108ef565b61094f61093e61357d565b61094661081b565b9182918261090e565b0390f35b610821565b906020828203126109715761096e915f01610844565b90565b610825565b5f0190565b346109a95761099361098e366004610958565b613659565b61099b61081b565b806109a581610976565b0390f35b610821565b346109dc576109c66109c1366004610958565b6136d1565b6109ce61081b565b806109d881610976565b0390f35b610821565b34610a0f576109f96109f4366004610958565b613736565b610a0161081b565b80610a0b81610976565b0390f35b610821565b610a1d816108fe565b03610a2457565b5f80fd5b90503590610a3582610a14565b565b90602082820312610a5057610a4d915f01610a28565b90565b610825565b610a5e9061082d565b9052565b9190610a75905f60208501940190610a55565b565b34610aa757610aa3610a92610a8d366004610a37565b613816565b610a9a61081b565b91829182610a62565b0390f35b610821565b34610adc57610abc3660046108ef565b610ad8610ac761392d565b610acf61081b565b91829182610a62565b0390f35b610821565b610aea8161088b565b03610af157565b5f80fd5b90503590610b0282610ae1565b565b90602082820312610b1d57610b1a915f01610af5565b90565b610825565b34610b5057610b3a610b35366004610b04565b613a0a565b610b4261081b565b80610b4c81610976565b0390f35b610821565b34610b8557610b653660046108ef565b610b81610b70613a40565b610b7861081b565b91829182610a62565b0390f35b610821565b1c90565b90565b610ba1906008610ba69302610b8a565b610b8e565b90565b90610bb49154610b91565b90565b610bc461010b5f90610ba9565b90565b34610bf757610bd73660046108ef565b610bf3610be2610bb7565b610bea61081b565b91829182610a62565b0390f35b610821565b5f80fd5b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b90610c2c90610c04565b810190811067ffffffffffffffff821117610c4657604052565b610c0e565b90610c5e610c5761081b565b9283610c22565b565b67ffffffffffffffff8111610c7e57610c7a602091610c04565b0190565b610c0e565b90825f939282370152565b90929192610ca3610c9e82610c60565b610c4b565b93818552602085019082840111610cbf57610cbd92610c83565b565b610c00565b9080601f83011215610ce257816020610cdf93359101610c8e565b90565b610bfc565b610cf09061088b565b90565b610cfc81610ce7565b03610d0357565b5f80fd5b90503590610d1482610cf3565b565b919060a083820312610d7f57610d2e815f8501610844565b92602081013567ffffffffffffffff8111610d7a5782610d4f918301610cc4565b92610d77610d608460408501610844565b93610d6e8160608601610d07565b93608001610844565b90565b610829565b610825565b34610db657610da0610d97366004610d16565b93929092613bf6565b610da861081b565b80610db281610976565b0390f35b610821565b34610de957610dd3610dce366004610958565b613c72565b610ddb61081b565b80610de581610976565b0390f35b610821565b9190604083820312610e165780610e0a610e13925f8601610844565b93602001610af5565b90565b610825565b90565b610e32610e2d610e379261082d565b610e1b565b61082d565b90565b90610e4490610e1e565b5f5260205260405f2090565b610e64610e5f610e6992610880565b610e1b565b610880565b90565b610e7590610e50565b90565b610e8190610e6c565b90565b90610e8e90610e78565b5f5260205260405f2090565b5f1c90565b610eab610eb091610e9a565b610b8e565b90565b610ebd9054610e9f565b90565b90610ed0610ed592610104610e3a565b610e84565b610ee05f8201610eb3565b91610ef96002610ef260018501610eb3565b9301610eb3565b90565b604090610f25610f2c9496959396610f1b60608401985f850190610a55565b6020830190610a55565b0190610a55565b565b34610f6257610f5e610f4a610f44366004610dee565b90610ec0565b610f5593919361081b565b93849384610efc565b0390f35b610821565b610f7461010a5f90610ba9565b90565b34610fa757610f873660046108ef565b610fa3610f92610f67565b610f9a61081b565b91829182610a62565b0390f35b610821565b610fb96101245f90610ba9565b90565b34610fec57610fcc3660046108ef565b610fe8610fd7610fac565b610fdf61081b565b91829182610a62565b0390f35b610821565b610ffe6101165f90610ba9565b90565b34611031576110113660046108ef565b61102d61101c610ff1565b61102461081b565b91829182610a62565b0390f35b610821565b346110645761104e611049366004610958565b613cea565b61105661081b565b8061106081610976565b0390f35b610821565b346110975761108161107c366004610958565b613d62565b61108961081b565b8061109381610976565b0390f35b610821565b5f80fd5b5f80fd5b909182601f830112156110de5781359167ffffffffffffffff83116110d95760200192600183028401116110d457565b6110a0565b61109c565b610bfc565b90602082820312611114575f82013567ffffffffffffffff811161110f5761110b92016110a4565b9091565b610829565b610825565b61112d6111273660046110e3565b906140c2565b61113561081b565b8061113f81610976565b0390f35b6111506101065f90610ba9565b90565b34611183576111633660046108ef565b61117f61116e611143565b61117661081b565b91829182610a62565b0390f35b610821565b346111b8576111983660046108ef565b6111b46111a36140da565b6111ab61081b565b91829182610a62565b0390f35b610821565b6111ca6101145f90610ba9565b90565b346111fd576111dd3660046108ef565b6111f96111e86111bd565b6111f061081b565b91829182610a62565b0390f35b610821565b346112325761122e61121d611218366004610958565b61417e565b61122561081b565b91829182610a62565b0390f35b610821565b60018060a01b031690565b6112529060086112579302610b8a565b611237565b90565b906112659154611242565b90565b61127561012c5f9061125a565b90565b61128190610e6c565b90565b61128d90611278565b9052565b91906112a4905f60208501940190611284565b565b346112d6576112b63660046108ef565b6112d26112c1611268565b6112c961081b565b91829182611291565b0390f35b610821565b3461130b576112eb3660046108ef565b6113076112f66141c0565b6112fe61081b565b91829182610a62565b0390f35b610821565b3461133e57611328611323366004610958565b61421c565b61133061081b565b8061133a81610976565b0390f35b610821565b60018060a01b031690565b61135e9060086113639302610b8a565b611343565b90565b90611371915461134e565b90565b61138161012d5f90611366565b90565b61138d90610e6c565b90565b61139990611384565b9052565b91906113b0905f60208501940190611390565b565b346113e2576113c23660046108ef565b6113de6113cd611374565b6113d561081b565b9182918261139d565b0390f35b610821565b34611415576113ff6113fa366004610958565b614294565b61140761081b565b8061141181610976565b0390f35b610821565b346114485761143261142d366004610958565b61430c565b61143a61081b565b8061144481610976565b0390f35b610821565b9061145790610e1e565b5f5260205260405f2090565b5f6114736114799261010361144d565b01610eb3565b90565b346114ac576114a8611497611492366004610958565b611463565b61149f61081b565b91829182610a62565b0390f35b610821565b6114be61011c5f90610ba9565b90565b346114f1576114d13660046108ef565b6114ed6114dc6114b1565b6114e461081b565b91829182610a62565b0390f35b610821565b6114ff9061088b565b90565b61150b816114f6565b0361151257565b5f80fd5b9050359061152382611502565b565b919060a08382031261158e5761153d815f8501610844565b92602081013567ffffffffffffffff8111611589578261155e918301610cc4565b9261158661156f8460408501610844565b9361157d8160608601611516565b93608001610844565b90565b610829565b610825565b346115c5576115af6115a6366004611525565b93929092614446565b6115b761081b565b806115c181610976565b0390f35b610821565b346115fa576115da3660046108ef565b6115f66115e5614455565b6115ed61081b565b91829182610a62565b0390f35b610821565b3461162f5761160f3660046108ef565b61162b61161a6144a4565b61162261081b565b9182918261090e565b0390f35b610821565b6116416101275f90610ba9565b90565b34611674576116543660046108ef565b61167061165f611634565b61166761081b565b91829182610a62565b0390f35b610821565b6116866101305f90610ba9565b90565b346116b9576116993660046108ef565b6116b56116a4611679565b6116ac61081b565b91829182610a62565b0390f35b610821565b346116ec576116d66116d1366004610958565b614542565b6116de61081b565b806116e881610976565b0390f35b610821565b3461171f576117013660046108ef565b611709614936565b61171161081b565b8061171b81610976565b0390f35b610821565b6117316101235f90610ba9565b90565b34611764576117443660046108ef565b61176061174f611724565b61175761081b565b91829182610a62565b0390f35b610821565b6117729061088b565b90565b61177e81611769565b0361178557565b5f80fd5b9050359061179682611775565b565b906020828203126117b1576117ae915f01611789565b90565b610825565b346117e4576117ce6117c9366004611798565b614a4b565b6117d661081b565b806117e081610976565b0390f35b610821565b6117f661011e5f90610ba9565b90565b34611829576118093660046108ef565b6118256118146117e9565b61181c61081b565b91829182610a62565b0390f35b610821565b60018060a01b031690565b61184990600861184e9302610b8a565b61182e565b90565b9061185c9154611839565b90565b61186c61012b5f90611851565b90565b61187890610e6c565b90565b6118849061186f565b9052565b919061189b905f6020850194019061187b565b565b346118cd576118ad3660046108ef565b6118c96118b861185f565b6118c061081b565b91829182611888565b0390f35b610821565b6118df6101205f90610ba9565b90565b34611912576118f23660046108ef565b61190e6118fd6118d2565b61190561081b565b91829182610a62565b0390f35b610821565b346119475761194361193261192d366004610a37565b614a56565b61193a61081b565b91829182610a62565b0390f35b610821565b67ffffffffffffffff811161196a57611966602091610c04565b0190565b610c0e565b9092919261198461197f8261194c565b610c4b565b938185526020850190828401116119a05761199e92610c83565b565b610c00565b9080601f830112156119c3578160206119c09335910161196f565b90565b610bfc565b919091604081840312611a08576119e1835f8301610af5565b92602082013567ffffffffffffffff8111611a0357611a0092016119a5565b90565b610829565b610825565b611a21611a1b3660046119c8565b90614ba5565b611a2961081b565b80611a3381610976565b0390f35b34611a6757611a473660046108ef565b611a63611a52614c05565b611a5a61081b565b91829182610a62565b0390f35b610821565b34611a9a57611a84611a7f366004610958565b614c9a565b611a8c61081b565b80611a9681610976565b0390f35b610821565b90565b611aab90611a9f565b9052565b9190611ac2905f60208501940190611aa2565b565b34611af457611ad43660046108ef565b611af0611adf614d14565b611ae761081b565b91829182611aaf565b0390f35b610821565b34611b2757611b11611b0c366004610958565b614d94565b611b1961081b565b80611b2381610976565b0390f35b610821565b611b3961011d5f90610ba9565b90565b34611b6c57611b4c3660046108ef565b611b68611b57611b2c565b611b5f61081b565b91829182610a62565b0390f35b610821565b611b7e6101265f90610ba9565b90565b34611bb157611b913660046108ef565b611bad611b9c611b71565b611ba461081b565b91829182610a62565b0390f35b610821565b611bc36101075f90610ba9565b90565b34611bf657611bd63660046108ef565b611bf2611be1611bb6565b611be961081b565b91829182610a62565b0390f35b610821565b611c049061088b565b90565b611c1081611bfb565b03611c1757565b5f80fd5b90503590611c2882611c07565b565b90602082820312611c4357611c40915f01611c1b565b90565b610825565b34611c7657611c60611c5b366004611c2a565b614eaa565b611c6861081b565b80611c7281610976565b0390f35b610821565b34611cab57611c8b3660046108ef565b611ca7611c96614eb5565b611c9e61081b565b91829182610a62565b0390f35b610821565b34611cde57611cc03660046108ef565b611cc8615288565b611cd061081b565b80611cda81610976565b0390f35b610821565b34611d1157611cfb611cf6366004610b04565b61531a565b611d0361081b565b80611d0d81610976565b0390f35b610821565b34611d4657611d263660046108ef565b611d42611d31615325565b611d3961081b565b91829182610a62565b0390f35b610821565b34611d7b57611d5b3660046108ef565b611d77611d66615362565b611d6e61081b565b91829182610a62565b0390f35b610821565b34611dae57611d98611d93366004610958565b6153ec565b611da061081b565b80611daa81610976565b0390f35b610821565b60018060a01b031690565b611dce906008611dd39302610b8a565b611db3565b90565b90611de19154611dbe565b90565b611df16101095f90611dd6565b90565b34611e2457611e043660046108ef565b611e20611e0f611de4565b611e1761081b565b918291826108a4565b0390f35b610821565b611e3661011a5f90610ba9565b90565b34611e6957611e493660046108ef565b611e65611e54611e29565b611e5c61081b565b91829182610a62565b0390f35b610821565b611e779061088b565b90565b611e8381611e6e565b03611e8a57565b5f80fd5b90503590611e9b82611e7a565b565b90602082820312611eb657611eb3915f01611e8e565b90565b610825565b34611ee957611ed3611ece366004611e9d565b615502565b611edb61081b565b80611ee581610976565b0390f35b610821565b34611f1e57611efe3660046108ef565b611f1a611f0961550d565b611f1161081b565b91829182610a62565b0390f35b610821565b611f3061010d5f90610ba9565b90565b34611f6357611f433660046108ef565b611f5f611f4e611f23565b611f5661081b565b91829182610a62565b0390f35b610821565b34611f9657611f783660046108ef565b611f8061554f565b611f8861081b565b80611f9281610976565b0390f35b610821565b60018060a01b031690565b611fb6906008611fbb9302610b8a565b611f9b565b90565b90611fc99154611fa6565b90565b611fd961012a5f90611fbe565b90565b611fe590610e6c565b90565b611ff190611fdc565b9052565b9190612008905f60208501940190611fe8565b565b3461203a5761201a3660046108ef565b612036612025611fcc565b61202d61081b565b91829182611ff5565b0390f35b610821565b61204c61012f5f90611dd6565b90565b3461207f5761205f3660046108ef565b61207b61206a61203f565b61207261081b565b918291826108a4565b0390f35b610821565b346120b25761209c612097366004610958565b6155c6565b6120a461081b565b806120ae81610976565b0390f35b610821565b919060a083820312612120576120cf815f8501610a28565b92602081013567ffffffffffffffff811161211b57826120f0918301610cc4565b926121186121018460408501610844565b9361210f8160608601610d07565b93608001610844565b90565b610829565b610825565b61213c6121333660046120b7565b939290926156a9565b61214461081b565b8061214e81610976565b0390f35b346121805761216a612165366004610958565b615725565b61217261081b565b8061217c81610976565b0390f35b610821565b346121b5576121953660046108ef565b6121b16121a0615730565b6121a861081b565b918291826108a4565b0390f35b610821565b9091606082840312612208576121d2835f8401610a28565b9260208301359067ffffffffffffffff8211612203576121f781612200938601610cc4565b93604001610844565b90565b610829565b610825565b61222161221b3660046121ba565b9161577a565b61222961081b565b8061223381610976565b0390f35b6122446101155f90610ba9565b90565b34612277576122573660046108ef565b612273612262612237565b61226a61081b565b91829182610a62565b0390f35b610821565b6122859061088b565b90565b6122918161227c565b0361229857565b5f80fd5b905035906122a982612288565b565b906020828203126122c4576122c1915f0161229c565b90565b610825565b346122f7576122e16122dc3660046122ab565b615892565b6122e961081b565b806122f381610976565b0390f35b610821565b6123096101135f90610ba9565b90565b3461233c5761231c3660046108ef565b6123386123276122fc565b61232f61081b565b91829182610a62565b0390f35b610821565b61234e6101055f90611dd6565b90565b34612381576123613660046108ef565b61237d61236c612341565b61237461081b565b918291826108a4565b0390f35b610821565b61238f9061088b565b90565b61239b81612386565b036123a257565b5f80fd5b905035906123b382612392565b565b906020828203126123ce576123cb915f016123a6565b90565b610825565b34612401576123eb6123e63660046123b5565b6159a8565b6123f361081b565b806123fd81610976565b0390f35b610821565b34612436576124163660046108ef565b6124326124216159b3565b61242961081b565b91829182610a62565b0390f35b610821565b346124695761245361244e366004610958565b615a21565b61245b61081b565b8061246581610976565b0390f35b610821565b60018060a01b031690565b61248990600861248e9302610b8a565b61246e565b90565b9061249c9154612479565b90565b6124ac61012e5f90612491565b90565b6124b890610e6c565b90565b6124c4906124af565b9052565b91906124db905f602085019401906124bb565b565b3461250d576124ed3660046108ef565b6125096124f861249f565b61250061081b565b918291826124c8565b0390f35b610821565b61251f6101125f90610ba9565b90565b34612552576125323660046108ef565b61254e61253d612512565b61254561081b565b91829182610a62565b0390f35b610821565b6125623660046108ef565b61256a615aa8565b61257261081b565b8061257c81610976565b0390f35b9061259261258d83610c60565b610c4b565b918252565b5f7f352e302e30000000000000000000000000000000000000000000000000000000910152565b6125c86005612580565b906125d560208301612597565b565b6125df6125be565b90565b6125ea6125d7565b90565b6125f56125e2565b90565b5190565b60209181520190565b90825f9392825e0152565b61262f61263860209361263d93612626816125f8565b938480936125fc565b95869101612605565b610c04565b0190565b6126569160208201915f818403910152612610565b90565b34612689576126693660046108ef565b6126856126746125ed565b61267c61081b565b91829182612641565b0390f35b610821565b61269b6101025f90611dd6565b90565b346126ce576126ae3660046108ef565b6126ca6126b961268e565b6126c161081b565b918291826108a4565b0390f35b610821565b6126e06101315f90611dd6565b90565b34612713576126f33660046108ef565b61270f6126fe6126d3565b61270661081b565b918291826108a4565b0390f35b610821565b61272561011b5f90610ba9565b90565b34612758576127383660046108ef565b612754612743612718565b61274b61081b565b91829182610a62565b0390f35b610821565b6127669061088b565b90565b6127728161275d565b0361277957565b5f80fd5b9050359061278a82612769565b565b906020828203126127a5576127a2915f0161277d565b90565b610825565b346127d8576127c26127bd36600461278c565b615bbd565b6127ca61081b565b806127d481610976565b0390f35b610821565b634e487b7160e01b5f52603260045260245ffd5b5490565b5f5260205f2090565b612807816127f1565b821015612821576128196004916127f5565b910201905f90565b6127dd565b61283261283791610e9a565b611db3565b90565b6128449054612826565b90565b634e487b7160e01b5f52602260045260245ffd5b906001600283049216801561287b575b602083101461287657565b612847565b91607f169161286b565b60209181520190565b5f5260205f2090565b905f92918054906128b16128aa8361285b565b8094612885565b916001811690815f1461290857506001146128cc575b505050565b6128d9919293945061288e565b915f925b8184106128f057505001905f80806128c7565b600181602092959395548486015201910192906128dd565b92949550505060ff19168252151560200201905f80806128c7565b9061292d91612897565b90565b906129506129499261294061081b565b93848092612923565b0383610c22565b565b6101009061295f826127f1565b8110156129a55761296f916127fe565b509061297c5f8301610eb3565b916129896001820161283a565b916129a2600361299b60028501610eb3565b9301612930565b90565b5f80fd5b90926129dc906129d26129e996946129c860808601975f870190610a55565b6020850190610897565b6040830190610a55565b6060818403910152612610565b90565b34612a2057612a1c612a07612a02366004610958565b612952565b90612a1394929461081b565b948594856129a9565b0390f35b610821565b34612a5557612a51612a40612a3b366004610a37565b615bc8565b612a4861081b565b91829182610a62565b0390f35b610821565b916020612a7b929493612a7460408201965f830190610a55565b0190610901565b565b34612aae57612a8d3660046108ef565b612a95615c85565b90612aaa612aa161081b565b92839283612a5a565b0390f35b610821565b919060a083820312612b1c57612acb815f8501610a28565b92602081013567ffffffffffffffff8111612b175782612aec918301610cc4565b92612b14612afd8460408501610844565b93612b0b8160608601611516565b93608001610844565b90565b610829565b610825565b612b38612b2f366004612ab3565b93929092615d87565b612b4061081b565b80612b4a81610976565b0390f35b612b5b61010c5f90610ba9565b90565b34612b8e57612b6e3660046108ef565b612b8a612b79612b4e565b612b8161081b565b91829182610a62565b0390f35b610821565b34612bc357612ba33660046108ef565b612bbf612bae615d96565b612bb661081b565b91829182610a62565b0390f35b610821565b612bd56101185f90610ba9565b90565b34612c0857612be83660046108ef565b612c04612bf3612bc8565b612bfb61081b565b91829182610a62565b0390f35b610821565b612c1a6101325f90610ba9565b90565b34612c4d57612c2d3660046108ef565b612c49612c38612c0d565b612c4061081b565b91829182610a62565b0390f35b610821565b612c5f6101175f90610ba9565b90565b34612c9257612c723660046108ef565b612c8e612c7d612c52565b612c8561081b565b91829182610a62565b0390f35b610821565b612ca461010f5f90610ba9565b90565b34612cd757612cb73660046108ef565b612cd3612cc2612c97565b612cca61081b565b91829182610a62565b0390f35b610821565b612d11612d1894612d07606094989795612cfd608086019a5f870190610897565b6020850190610a55565b6040830190610897565b0190610a55565b565b34612d4e57612d2a3660046108ef565b612d4a612d35615dfd565b90612d4194929461081b565b94859485612cdc565b0390f35b610821565b34612d8157612d6b612d66366004610958565b616079565b612d7361081b565b80612d7d81610976565b0390f35b610821565b612d9361010e5f90610ba9565b90565b34612dc657612da63660046108ef565b612dc2612db1612d86565b612db961081b565b91829182610a62565b0390f35b610821565b9091606082840312612e1957612de3835f8401610844565b9260208301359067ffffffffffffffff8211612e1457612e0881612e11938601610cc4565b93604001610844565b90565b610829565b610825565b34612e4d57612e37612e31366004612dcb565b916160b0565b612e3f61081b565b80612e4981610976565b0390f35b610821565b34612e8057612e6a612e65366004610958565b61612a565b612e7261081b565b80612e7c81610976565b0390f35b610821565b612e926101335f90610ba9565b90565b34612ec557612ea53660046108ef565b612ec1612eb0612e85565b612eb861081b565b91829182610a62565b0390f35b610821565b34612ef857612ee2612edd366004610958565b616166565b612eea61081b565b80612ef481610976565b0390f35b610821565b34612f2b57612f15612f10366004610958565b6161de565b612f1d61081b565b80612f2781610976565b0390f35b610821565b34612f5e57612f403660046108ef565b612f486163de565b612f5061081b565b80612f5a81610976565b0390f35b610821565b612f706101215f90610ba9565b90565b34612fa357612f833660046108ef565b612f9f612f8e612f63565b612f9661081b565b91829182610a62565b0390f35b610821565b612fb56101015f90611dd6565b90565b34612fe857612fc83660046108ef565b612fe4612fd3612fa8565b612fdb61081b565b918291826108a4565b0390f35b610821565b612ffa6101085f90610ba9565b90565b3461302d5761300d3660046108ef565b613029613018612fed565b61302061081b565b91829182610a62565b0390f35b610821565b61303f6101255f90610ba9565b90565b34613072576130523660046108ef565b61306e61305d613032565b61306561081b565b91829182610a62565b0390f35b610821565b6130846101105f90610ba9565b90565b346130b7576130973660046108ef565b6130b36130a2613077565b6130aa61081b565b91829182610a62565b0390f35b610821565b906130c690610e1e565b5f5260205260405f2090565b6130e9906130e4610119915f926130bc565b610ba9565b90565b3461311c57613118613107613102366004610958565b6130d2565b61310f61081b565b91829182610a62565b0390f35b610821565b3461314f57613139613134366004610958565b616455565b61314161081b565b8061314b81610976565b0390f35b610821565b34613184576131643660046108ef565b61318061316f61646e565b61317761081b565b9182918261090e565b0390f35b610821565b6131966101115f90610ba9565b90565b346131c9576131a93660046108ef565b6131c56131b4613189565b6131bc61081b565b91829182610a62565b0390f35b610821565b346131fc576131e66131e1366004610958565b6164f7565b6131ee61081b565b806131f881610976565b0390f35b610821565b61320e61011f5f90610ba9565b90565b34613241576132213660046108ef565b61323d61322c613201565b61323461081b565b91829182610a62565b0390f35b610821565b346132745761325e613259366004610b04565b616567565b61326661081b565b8061327081610976565b0390f35b610821565b346132a75761329161328c366004610958565b6165a3565b61329961081b565b806132a381610976565b0390f35b610821565b6132b96101285f90610ba9565b90565b346132ec576132cc3660046108ef565b6132e86132d76132ac565b6132df61081b565b91829182610a62565b0390f35b610821565b3461331f57613309613304366004610958565b61661b565b61331161081b565b8061331b81610976565b0390f35b610821565b6133316101225f90610ba9565b90565b34613364576133443660046108ef565b61336061334f613324565b61335761081b565b91829182610a62565b0390f35b610821565b3461339a576133793660046108ef565b613381616626565b9061339661338d61081b565b92839283612a5a565b0390f35b610821565b60018060a01b031690565b6133ba9060086133bf9302610b8a565b61339f565b90565b906133cd91546133aa565b90565b6133dd6101295f906133c2565b90565b6133e990610e6c565b90565b6133f5906133e0565b9052565b919061340c905f602085019401906133ec565b565b3461343e5761341e3660046108ef565b61343a6134296133d0565b61343161081b565b918291826133f9565b0390f35b610821565b346134735761346f61345e613459366004610958565b61664f565b61346661081b565b91829182610a62565b0390f35b610821565b91602061349992949361349260408201965f830190610a55565b0190610a55565b565b346134cd576134b46134ae366004610dee565b9061667c565b906134c96134c061081b565b92839283613478565b0390f35b610821565b34613500576134ea6134e5366004610958565b616733565b6134f261081b565b806134fc81610976565b0390f35b610821565b5f80fd5b5f90565b9061351790610e1e565b5f5260205260405f2090565b61354c91600161354161354793613538613509565b5061010361144d565b0161350d565b61283a565b90565b5f90565b61356761356261356c9261082d565b610e1b565b6108fe565b90565b9061357a91036108fe565b90565b61358561354f565b506135ab61359242613553565b6135a56135a061010d610eb3565b613553565b9061356f565b90565b6135bf906135ba6167c1565b6135c1565b565b6135d2906135cd6168a9565b613612565b565b5f1b90565b906135e55f19916135d4565b9181191691161790565b90565b9061360761360261360e92610e1e565b6135ef565b82546135d9565b9055565b61361e816101336135f2565b6136547facbc6b6929088e4b2d043625fa7248e00fb6658a425eb9d9dc8c37b18c7f3e6f9161364b61081b565b91829182610a62565b0390a1565b613662906135ae565b565b613675906136706167c1565b613677565b565b613688906136836168a9565b61368a565b565b6136968161011b6135f2565b6136cc7f40b9c59af8c486ccf8c7cc73df5a51e7cc29747ea7d39f99632ecaf9caa2ed1f916136c361081b565b91829182610a62565b0390a1565b6136da90613664565b565b6136ed906136e86167c1565b6136ef565b565b6136fb8161010c6135f2565b6137317fb0868a729f47ae3829aaafe3ca2975d3db2148553c854112f598be6d91ef0d289161372861081b565b91829182610a62565b0390a1565b61373f906136dc565b565b5f90565b90565b61375c61375761376192613745565b610e1b565b610880565b90565b61376d90613748565b90565b9061377b91016108fe565b90565b61379261378d61379792613745565b610e1b565b61082d565b90565b6137ae6137a96137b392613745565b610e1b565b6108fe565b90565b6137ca6137c56137cf926108fe565b610e1b565b61082d565b90565b906137dd910261082d565b90565b634e487b7160e01b5f52601260045260245ffd5b6138006138069161082d565b9161082d565b908115613811570490565b6137e0565b61387c90613822613741565b5061382e61010161283a565b61384861384261383d5f613764565b61088b565b9161088b565b145f146138ec5761387661387061386061010d610eb3565b5b9261386b42613553565b613770565b91613553565b9061356f565b6138855f61377e565b90806138a161389b6138965f61379a565b6108fe565b916108fe565b136138ab575b5090565b6138e691506138d06138bf6138e1926137b6565b6138ca61011b610eb3565b906137d2565b6138db610125610eb3565b906137f4565b616b37565b5f6138a7565b613876613870613928600261392261391161010461390b61010b610eb3565b90610e3a565b61391c61010161283a565b90610e84565b01610eb3565b613861565b613935613741565b506139416101006127f1565b90565b613955906139506167c1565b613957565b565b613968906139636168a9565b61396a565b565b61397c9061397781616efb565b6139bc565b565b9061398f60018060a01b03916135d4565b9181191691161790565b90565b906139b16139ac6139b892610e78565b613999565b825461397e565b9055565b6139c88161013161399c565b6139f27f1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c91610e78565b906139fb61081b565b80613a0581610976565b0390a2565b613a1390613944565b565b613a1e90610e6c565b90565b90565b613a38613a33613a3d92613a21565b610e1b565b61082d565b90565b613a48613741565b50613a7a613a6a613a5830613a15565b31613a64610132610eb3565b906137d2565b613a746064613a24565b906137f4565b90565b90613a9294939291613a8d616f8a565b613b3d565b613a9a616fcf565b565b613aa8613aad91610e9a565b611237565b90565b613aba9054613a9c565b90565b5f80fd5b60e01b90565b5f910312613ad157565b610825565b613adf90610e6c565b90565b613aeb90613ad6565b9052565b613b24613b2b94613b1a606094989795613b10608086019a5f870190610a55565b6020850190610897565b6040830190613ae2565b0190610a55565b565b613b3561081b565b3d5f823e3d90fd5b91613b4d92949394919091617357565b613b60613b5b61012c613ab0565b611278565b9063e2051c7e91613b7261010b610eb3565b91613b7b6177a1565b9490823b15613bf1575f94613bae8692613ba394613b9761081b565b998a9889978896613ac1565b865260048601613aef565b03925af18015613bec57613bc0575b50565b613bdf905f3d8111613be5575b613bd78183610c22565b810190613ac7565b5f613bbd565b503d613bcd565b613b2d565b613abd565b90613c0394939291613a7d565b565b613c1690613c116167c1565b613c18565b565b613c2990613c246168a9565b613c2b565b565b613c37816101216135f2565b613c6d7f3312247fcf207243294680d9103851c8709e19be3d46ee4b1bff5e12d1c5ca7b91613c6461081b565b91829182610a62565b0390a1565b613c7b90613c05565b565b613c8e90613c896167c1565b613c90565b565b613ca190613c9c6168a9565b613ca3565b565b613caf816101206135f2565b613ce57f85d8bf21006916836edd67a5afeed2e891cf876a6c9cd9babf7f42f4b007c24f91613cdc61081b565b91829182610a62565b0390a1565b613cf390613c7d565b565b613d0690613d016167c1565b613d08565b565b613d1990613d146168a9565b613d1b565b565b613d278161011d6135f2565b613d5d7f5581e31f5b8d4b3c45b8ab8bf67e3602ce588b423905eb6ad34bd6bc3c84869991613d5461081b565b91829182610a62565b0390a1565b613d6b90613cf5565b565b90613d7f91613d7a616f8a565b613fe8565b613d87616fcf565b565b90565b5490565b5f5260205f2090565b613da281613d8c565b821015613dbc57613db4600491613d90565b910201905f90565b6127dd565b613dca81613d8c565b68010000000000000000811015613dee57613dea91600182018155613d99565b9091565b610c0e565b90565b5090565b601f602091010490565b1b90565b91906008613e23910291613e1d5f1984613e04565b92613e04565b9181191691161790565b9190613e43613e3e613e4b93610e1e565b6135ef565b908354613e08565b9055565b613e6191613e5b613741565b91613e2d565b565b5f5b828110613e7157505050565b80613e805f6001938501613e4f565b01613e65565b9190601f8111613e96575b505050565b818111613ea3575b613e91565b613eb8613eb2613ed69461288e565b91613dfa565b916020613ec482613dfa565b9110613ede575b809101910390613e63565b5f8080613e9e565b505f613ecb565b90613ef5905f1990600802610b8a565b191690565b81613f0491613ee5565b906002021790565b91613f179082613df6565b9067ffffffffffffffff8211613fd657613f3b82613f35855461285b565b85613e86565b5f90601f8311600114613f6e57918091613f5d935f92613f62575b5050613efa565b90555b565b90915001355f80613f56565b601f19831691613f7d8561288e565b925f5b818110613fbe57509160029391856001969410613fa4575b50505002019055613f60565b613fb4910135601f841690613ee5565b90555f8080613f98565b91936020600181928787013581550195019201613f80565b610c0e565b90613fe69291613f0c565b565b61405690613ff76101006127f1565b92600361401661401061400b610100613d89565b613dc1565b50613df3565b9261402d61402561010b610eb3565b5f86016135f2565b6140416140386177a1565b6001860161399c565b61404e34600286016135f2565b919201613fdb565b61406161010b610eb3565b6140696177a1565b3492916140bd6140ab6140a561409f7fa08049565b10d44a06dca9bf05685b39bc370352043c5a003e8d35d45ebdc53f94610e1e565b94610e78565b94610e1e565b946140b461081b565b91829182610a62565b0390a4565b906140cc91613d6d565b565b6140d790610e6c565b90565b6140e2613741565b506141146141046140f2306140ce565b316140fe61011d610eb3565b906137d2565b61410e6064613a24565b906137f4565b90565b90565b61412e61412961413392614117565b610e1b565b61082d565b90565b614140600261411a565b90565b90565b61415a61415561415f92614143565b610e1b565b61082d565b90565b9061416d910361082d565b90565b9061417b910161082d565b90565b6141af6141bd9161418d613741565b506141a9614199614136565b6141a36001614146565b90614162565b90614170565b6141b7614136565b906137f4565b90565b6141c8613741565b506141e86141d7610125610eb3565b6141e2610123610eb3565b906137f4565b90565b6141fc906141f76167c1565b6141fe565b565b61420f9061420a61782d565b614211565b565b61421a90617885565b565b614225906141eb565b565b614238906142336167c1565b61423a565b565b61424b906142466168a9565b61424d565b565b614259816101326135f2565b61428f7ffe65b6d5007c66dadebc5095104ccd672c070f396dfdcfe1ee7f54201b9efa6d9161428661081b565b91829182610a62565b0390a1565b61429d90614227565b565b6142b0906142ab6167c1565b6142b2565b565b6142c3906142be6168a9565b6142c5565b565b6142d18161011f6135f2565b6143077f4787028773c8f14bc3b4bc41f43a02329ae41105823287201a34782c530d35fd916142fe61081b565b91829182610a62565b0390a1565b6143159061429f565b565b9061432c94939291614327616f8a565b61438d565b614334616fcf565b565b61433f90610e6c565b90565b61434b90614336565b9052565b61438461438b9461437a606094989795614370608086019a5f870190610a55565b6020850190610897565b6040830190614342565b0190610a55565b565b9161439d92949394919091617357565b6143b06143ab61012c613ab0565b611278565b9063fe673fd3916143c261010b610eb3565b916143cb6177a1565b9490823b15614441575f946143fe86926143f3946143e761081b565b998a9889978896613ac1565b86526004860161434f565b03925af1801561443c57614410575b50565b61442f905f3d8111614435575b6144278183610c22565b810190613ac7565b5f61440d565b503d61441d565b613b2d565b613abd565b9061445394939291614317565b565b61445d613741565b506144666144a4565b8061448161447b6144765f61379a565b6108fe565b916108fe565b135f1461449557614491906137b6565b5b90565b5061449f5f61377e565b614492565b6144ac61354f565b506144d26144c36144be610124610eb3565b613553565b6144cc42613553565b9061356f565b90565b6144e6906144e16167c1565b6144e8565b565b6144f9906144f46168a9565b6144fb565b565b614507816101136135f2565b61453d7fa787f26546d7eeea63d70fc31736f27ad28329e95982f3bc5a7e0280f497bbf59161453461081b565b91829182610a62565b0390a1565b61454b906144d5565b565b614555616f8a565b61455d61479d565b614565616fcf565b565b151590565b60207f757272656e742062696464696e6720726f756e64207965742e00000000000000917f54686572652068617665206265656e206e6f206269647320696e2074686520635f8201520152565b6145c660396040926125fc565b6145cf8161456c565b0190565b6145e89060208101905f8183039101526145b9565b90565b634e487b7160e01b5f52601160045260245ffd5b61460e614614919392936108fe565b926108fe565b808301925f8285121581831216928512911215161761462f57565b6145eb565b60607f2e00000000000000000000000000000000000000000000000000000000000000917f4f6e6c7920746865206c61737420626964646572206973207065726d697474655f8201527f6420746f20636c61696d207468652062696464696e6720726f756e64206d616960208201527f6e207072697a65206265666f726520612074696d656f7574206578706972657360408201520152565b6146da60616080926125fc565b6146e381614634565b0190565b606090614720614727949695939661471661470b608085018581035f8701526146cd565b986020850190610897565b6040830190610897565b0190610a55565b565b5f7f4e6f7420656e6f7567682074696d652068617320656c61707365642e00000000910152565b61475d601c6020926125fc565b61476681614729565b0190565b91604061479b929493614794614789606083018381035f850152614750565b966020830190610a55565b0190610a55565b565b6147a56177a1565b6147c16147bb6147b661010161283a565b61088b565b9161088b565b145f1461484a576147f0426147e86147e26147dd610124610eb3565b61082d565b9161082d565b101515614567565b614818575b6147fd6178cc565b614806426179fa565b61480e618373565b6148166192c5565b565b614823610124610eb3565b429061484661483061081b565b928392638d31bb1560e01b84526004840161476a565b0390fd5b61487a61485861010161283a565b61487261486c6148675f613764565b61088b565b9161088b565b141515614567565b614913576148a26148896144a4565b61489c614897610127610eb3565b613553565b906145ff565b6148c8816148c06148ba6148b55f61379a565b6108fe565b916108fe565b131515614567565b6148d257506147f5565b6148dd61010161283a565b61490f6148f16148eb6177a1565b936137b6565b6148f961081b565b93849363336598a360e21b8552600485016146e7565b0390fd5b61491b61081b565b6318844a7d60e31b815280614932600482016145d3565b0390fd5b61493e61454d565b565b6149519061494c6167c1565b614953565b565b6149649061495f6168a9565b614972565b565b61496f90610e6c565b90565b61498c9061498761498282614966565b616efb565b6149ed565b565b61499790610e50565b90565b6149a39061498e565b90565b6149af9061498e565b90565b90565b906149ca6149c56149d1926149a6565b6149b2565b825461397e565b9055565b6149de90610e50565b90565b6149ea906149d5565b90565b614a09614a016149fc83614966565b61499a565b61012b6149b5565b614a337f5bde6238168795ba4e77972a2bdaa5a465f7c9a5d05817f5e8d3fed2e5a4fa60916149e1565b90614a3c61081b565b80614a4681610976565b0390a2565b614a5490614940565b565b614a5e613741565b50614a67613741565b50614a7361010161283a565b614a8d614a87614a825f613764565b61088b565b9161088b565b145f14614b6b57614ab0614aa261010f610eb3565b91614aab61357d565b613770565b80614acb614ac5614ac05f61379a565b6108fe565b916108fe565b13155f14614ad957505b5b90565b90614b02614af282614aec610110610eb3565b906137f4565b614afc6001614146565b90614170565b91614b0b6193be565b90614b15816137b6565b614b27614b218461082d565b9161082d565b105f14614b635790614b52614b5792614b4c614b46614b5d9787614162565b916137b6565b906137d2565b6137f4565b90614162565b5b614ad5565b505050614b5e565b50614b77610111610eb3565b614ad6565b90614b8e91614b896193f5565b614b90565b565b90614ba391614b9e816194ba565b61952a565b565b90614baf91614b7c565b565b90565b614bc8614bc3614bcd92614bb1565b610e1b565b61082d565b90565b614bdb6103e8614bb4565b90565b614bf7614be9614bd0565b614bf1614bd0565b906137d2565b90565b614c02614bde565b90565b614c0d613741565b50614c2a614c1c610125610eb3565b614c24614bfa565b906137f4565b90565b614c3e90614c396167c1565b614c40565b565b614c5190614c4c6168a9565b614c53565b565b614c5f816101226135f2565b614c957f9e44c04f534af356419a731b967f3d56bc748b1f3fdbda7a89f4e1a321ade93491614c8c61081b565b91829182610a62565b0390a1565b614ca390614c2d565b565b5f90565b614cba90614cb5619628565b614d08565b90565b90565b614cd4614ccf614cd992614cbd565b6135d4565b611a9f565b90565b614d057f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc614cc0565b90565b50614d11614cdc565b90565b614d24614d1f614ca5565b614ca9565b90565b614d3890614d336167c1565b614d3a565b565b614d4b90614d466168a9565b614d4d565b565b614d598161011a6135f2565b614d8f7f157c413b0549fd4f45aab72b7828304fb2c45dad53de0f1128c5eabf3aaabaf891614d8661081b565b91829182610a62565b0390a1565b614d9d90614d27565b565b614db090614dab6167c1565b614db2565b565b614dc390614dbe6168a9565b614dd1565b565b614dce90610e6c565b90565b614deb90614de6614de182614dc5565b616efb565b614e4c565b565b614df690610e50565b90565b614e0290614ded565b90565b614e0e90614ded565b90565b90565b90614e29614e24614e3092614e05565b614e11565b825461397e565b9055565b614e3d90610e50565b90565b614e4990614e34565b90565b614e68614e60614e5b83614dc5565b614df9565b61012e614e14565b614e927f4da1815cd654922275d14d2335fd9a0dd0aa6a0d0ff87fb4cc872ebe9704596f91614e40565b90614e9b61081b565b80614ea581610976565b0390a2565b614eb390614d9f565b565b614ebd613741565b50614eef614edf614ecd30613a15565b31614ed9610128610eb3565b906137d2565b614ee96064613a24565b906137f4565b90565b614efa6196a2565b614f02615059565b565b67ffffffffffffffff1690565b614f25614f20614f2a92614117565b610e1b565b614f04565b90565b60401c90565b60ff1690565b614f45614f4a91614f2d565b614f33565b90565b614f579054614f39565b90565b67ffffffffffffffff1690565b614f73614f7891610e9a565b614f5a565b90565b614f859054614f67565b90565b90614f9b67ffffffffffffffff916135d4565b9181191691161790565b614fb9614fb4614fbe92614f04565b610e1b565b614f04565b90565b90565b90614fd9614fd4614fe092614fa5565b614fc1565b8254614f88565b9055565b60401b90565b90614ffe68ff000000000000000091614fe4565b9181191691161790565b61501190614567565b90565b90565b9061502c61502761503392615008565b615014565b8254614fea565b9055565b61504090614f04565b9052565b9190615057905f60208501940190615037565b565b6150636002614f11565b61506b6196f1565b6150765f8201614f4d565b8015615106575b6150ea576150af90615091835f8301614fc4565b61509e60015f8301615017565b6150a661512b565b5f809101615017565b6150e57fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2916150dc61081b565b91829182615044565b0390a1565b5f63f92ee8a960e01b81528061510260048201610976565b0390fd5b506151125f8201614f7b565b61512461511e84614f04565b91614f04565b101561507d565b61513361524d565b565b90565b61514c61514761515192615135565b610e1b565b61082d565b90565b61515f61a8c0615138565b90565b90565b61517961517461517e92615162565b610e1b565b61082d565b90565b61518b60fa615165565b90565b90565b6151a56151a06151aa9261518e565b610e1b565b61082d565b90565b6151b8610e10615191565b90565b90565b6151d26151cd6151d7926151bb565b610e1b565b61082d565b90565b6151e96151ef9193929361082d565b9261082d565b916151fb83820261082d565b92818404149015171561520a57565b6145eb565b61524a61523c6f0241c76b735b154119e2dd30000000006152376152316151ad565b916151be565b6151da565b615244614bfa565b906151da565b90565b615260615258615154565b6101156135f2565b61527361526b615181565b6101336135f2565b61528661527e61520f565b61011b6135f2565b565b615290614ef2565b565b6152a39061529e6167c1565b6152a5565b565b6152b6906152b16168a9565b6152b8565b565b6152ca906152c581616efb565b6152cc565b565b6152d88161012f61399c565b6153027f4d03942c29c20d1bccfe551e9d148c917c5a44fb558a4fc60270d8f76fb75f5491610e78565b9061530b61081b565b8061531581610976565b0390a2565b61532390615292565b565b61532d613741565b5061535f61534f61533d306140ce565b31615349610122610eb3565b906137d2565b6153596064613a24565b906137f4565b90565b61536a613741565b5061537c6153775f61379a565b614a56565b90565b6153909061538b6167c1565b615392565b565b6153a39061539e6168a9565b6153a5565b565b6153b1816101286135f2565b6153e77fb5a05ec7911dd5450a7fa4ae54d50d9d83af6e256db8fc76c82edd7b659cf8bc916153de61081b565b91829182610a62565b0390a1565b6153f59061537f565b565b615408906154036167c1565b61540a565b565b61541b906154166168a9565b615429565b565b61542690610e6c565b90565b6154439061543e6154398261541d565b616efb565b6154a4565b565b61544e90610e50565b90565b61545a90615445565b90565b61546690615445565b90565b90565b9061548161547c6154889261545d565b615469565b825461397e565b9055565b61549590610e50565b90565b6154a19061548c565b90565b6154c06154b86154b38361541d565b615451565b61012c61546c565b6154ea7fb4cecfe1346c94da27291cf5a02969d5fe0b5c36eca49b04fcd60841d28c5e1391615498565b906154f361081b565b806154fd81610976565b0390a2565b61550b906153f7565b565b615515613741565b506155276155225f61379a565b615bc8565b90565b6155326167c1565b61553a61553c565b565b61554d6155485f613764565b619715565b565b61555761552a565b565b61556a906155656167c1565b61556c565b565b61557d906155786168a9565b61557f565b565b61558b816101236135f2565b6155c17fb5edd1f338b34c8f5dd3b1c5cc12f05653c495713c282bf588d34cf14fad0f89916155b861081b565b91829182610a62565b0390a1565b6155cf90615559565b565b906155e6949392916155e1616f8a565b6155f0565b6155ee616fcf565b565b9161560092949394919091619b49565b61561361560e61012c613ab0565b611278565b9063e2051c7e9161562561010b610eb3565b9161562e6177a1565b9490823b156156a4575f9461566186926156569461564a61081b565b998a9889978896613ac1565b865260048601613aef565b03925af1801561569f57615673575b50565b615692905f3d8111615698575b61568a8183610c22565b810190613ac7565b5f615670565b503d615680565b613b2d565b613abd565b906156b6949392916155d1565b565b6156c9906156c46167c1565b6156cb565b565b6156dc906156d76168a9565b6156de565b565b6156ea816101186135f2565b6157207f4e8c80fe79d13b8663de9f6981925ae24c498cc07b0ebd4f4789fa78ca17caff9161571761081b565b91829182610a62565b0390a1565b61572e906156b8565b565b615738613509565b5061574b5f61574561a141565b0161283a565b90565b90615761929161575c616f8a565b61576b565b615769616fcf565b565b9161577892919091619b49565b565b90615785929161574e565b565b615798906157936167c1565b61579a565b565b6157ab906157a66168a9565b6157b9565b565b6157b690610e6c565b90565b6157d3906157ce6157c9826157ad565b616efb565b615834565b565b6157de90610e50565b90565b6157ea906157d5565b90565b6157f6906157d5565b90565b90565b9061581161580c615818926157ed565b6157f9565b825461397e565b9055565b61582590610e50565b90565b6158319061581c565b90565b615850615848615843836157ad565b6157e1565b6101296157fc565b61587a7f9b3eda10f1724a2cf9f7dae4ac263c77908df4d00e92f1377b66fc8be37cd8c191615828565b9061588361081b565b8061588d81610976565b0390a2565b61589b90615787565b565b6158ae906158a96167c1565b6158b0565b565b6158c1906158bc6168a9565b6158cf565b565b6158cc90610e6c565b90565b6158e9906158e46158df826158c3565b616efb565b61594a565b565b6158f490610e50565b90565b615900906158eb565b90565b61590c906158eb565b90565b90565b9061592761592261592e92615903565b61590f565b825461397e565b9055565b61593b90610e50565b90565b61594790615932565b90565b61596661595e615959836158c3565b6158f7565b61012a615912565b6159907fdab38e33e6e11cbb8b085bba9d7426d5e5af01bcc46d2c5957645e4d09e8c49c9161593e565b9061599961081b565b806159a381610976565b0390a2565b6159b19061589d565b565b6159bb613741565b506159ed6159dd6159cb306140ce565b316159d761011e610eb3565b906137d2565b6159e76064613a24565b906137f4565b90565b615a01906159fc6167c1565b615a03565b565b615a1490615a0f6168a9565b615a16565b565b615a1f9061a165565b565b615a2a906159f0565b565b615a34616f8a565b615a3c615a46565b615a44616fcf565b565b615a5161010b610eb3565b615a596177a1565b3491615aa3615a91615a8b7fe32cacf203d00685e2b4d8b0a90e7cd8f3f8a208fdf116f4bb36abe08b7d548e93610e1e565b93610e78565b93615a9a61081b565b91829182610a62565b0390a3565b615ab0615a2c565b565b615ac390615abe6167c1565b615ac5565b565b615ad690615ad16168a9565b615ae4565b565b615ae190610e6c565b90565b615afe90615af9615af482615ad8565b616efb565b615b5f565b565b615b0990610e50565b90565b615b1590615b00565b90565b615b2190615b00565b90565b90565b90615b3c615b37615b4392615b18565b615b24565b825461397e565b9055565b615b5090610e50565b90565b615b5c90615b47565b90565b615b7b615b73615b6e83615ad8565b615b0c565b61012d615b27565b615ba57fbf6e296f85d08cc1ab124aed644bf4b19e4a726a7aea53e3784ab1341738a04091615b53565b90615bae61081b565b80615bb881610976565b0390a2565b615bc690615ab2565b565b615be290615bd4613741565b50615bdd61a1ac565b61356f565b80615bfd615bf7615bf25f61379a565b6108fe565b916108fe565b1315615c7857615c51615c6291615c1561010261283a565b615c2f615c29615c245f613764565b61088b565b9161088b565b145f14615c6557615c4b615c44610117610eb3565b5b916137b6565b906137d2565b615c5c610115610eb3565b906137f4565b90565b615c4b615c73610116610eb3565b615c45565b50615c825f61377e565b90565b615c8d613741565b50615c9661354f565b50615c9f61a1db565b90615cab610115610eb3565b9190565b90615cc494939291615cbf616f8a565b615cce565b615ccc616fcf565b565b91615cde92949394919091619b49565b615cf1615cec61012c613ab0565b611278565b9063fe673fd391615d0361010b610eb3565b91615d0c6177a1565b9490823b15615d82575f94615d3f8692615d3494615d2861081b565b998a9889978896613ac1565b86526004860161434f565b03925af18015615d7d57615d51575b50565b615d70905f3d8111615d76575b615d688183610c22565b810190613ac7565b5f615d4e565b503d615d5e565b613b2d565b613abd565b90615d9494939291615caf565b565b615d9e613741565b50615db0615dab5f61379a565b613816565b90565b615dc2615dc89193929361082d565b9261082d565b8203918211615dd357565b6145eb565b615de7615ded9193929361082d565b9261082d565b8201809211615df857565b6145eb565b615e05613509565b90615e0e613741565b90615e17613509565b90615e20613741565b90615e2c61010161283a565b615e46615e40615e3b5f613764565b61088b565b9161088b565b03615e4e575b565b9350505050615e5e61010561283a565b90615e6a610106610eb3565b615e75610107610eb3565b92615e81610108610eb3565b92615e8d61010961283a565b92615e9961010a610eb3565b92615ed06002615eca615eb9610104615eb361010b610eb3565b90610e3a565b615ec461010161283a565b90610e84565b01610eb3565b96615edc428990615db3565b9282615ef8615ef2615eed5f613764565b61088b565b9161088b565b145f14615f6157505050615f2590615f1f615f1461010161283a565b9791965b4292615dd8565b90615db3565b615f2e81613553565b615f48615f42615f3d86613553565b6108fe565b916108fe565b13615f54575b50615e4c565b925090508391905f615f4e565b92809891979298615f7a615f748a61082d565b9161082d565b11615f8f575b505090615f1f615f2592615f18565b9692615fab829993615fa5615fb1948790615dd8565b92615dd8565b90615db3565b90615fbb82613553565b615fd5615fcf615fca88613553565b6108fe565b916108fe565b13615ffa575b5050615f259094615f1f615ff061010161283a565b9791969192615f80565b909450615f259193509392905f615fdb565b61601d906160186167c1565b61601f565b565b6160309061602b6168a9565b616032565b565b61603e816101266135f2565b6160747f4636d3e567b27988879babd22e50f49104ca65647933fc1623ff3d4d807438d29161606b61081b565b91829182610a62565b0390a1565b6160829061600c565b565b906160979291616092616f8a565b6160a1565b61609f616fcf565b565b916160ae92919091617357565b565b906160bb9291616084565b565b6160ce906160c96167c1565b6160d0565b565b6160e1906160dc6168a9565b6160e3565b565b6160ef816101306135f2565b6161257f2652e6657dd1ed89d7bdcb70f8827cc8449ac4536ebf604dbb2465cdad2640209161611c61081b565b91829182610a62565b0390a1565b616133906160bd565b565b616146906161416167c1565b616148565b565b616159906161546168a9565b61615b565b565b6161649061a20c565b565b61616f90616135565b565b6161829061617d6167c1565b616184565b565b616195906161906168a9565b616197565b565b6161a3816101156135f2565b6161d97f4abea08c196329c357e3175d011af39a8625be99ef0ba5a0f3547a95534fedb7916161d061081b565b91829182610a62565b0390a1565b6161e790616171565b565b6161f16167c1565b6161f96161fb565b565b616203616205565b565b61620d61782d565b616215616292565b565b5f7f546f6f206561726c792e00000000000000000000000000000000000000000000910152565b61624b600a6020926125fc565b61625481616217565b0190565b61626d9060208101905f81830391015261623e565b90565b61627c6162829161082d565b9161082d565b90811561628d570490565b6137e0565b6162c161629d616626565b91906162ba6162b46162af8593613553565b6108fe565b916108fe565b1315614567565b6163bb576163b9906163b46163af61639f6162dd610110610eb3565b61639961631c61630b6162fb6162f461010f610eb3565b8590616270565b6163056001614146565b90615dd8565b92616316600261411a565b906151da565b9561639361638d61637c61636b61635161634161633a61010f610eb3565b8d90616270565b61634b6001614146565b90615dd8565b9661635a613741565b5061636661010f610eb3565b615db3565b616376610125610eb3565b906151da565b9461638861010f610eb3565b615db3565b916137b6565b906151da565b90616270565b6163a96001614146565b90615dd8565b61a253565b61a20c565b565b6163c361081b565b63a29f5c4d60e01b8152806163da60048201616258565b0390fd5b6163e66161e9565b565b6163f9906163f46167c1565b6163fb565b565b61640c906164076168a9565b61640e565b565b61641a8161011c6135f2565b6164507fd95e7f967f9370c11deb15ffbb191b9f2e9795ab0738db5bc72bd2794978f32d9161644761081b565b91829182610a62565b0390a1565b61645e906163e8565b565b61646b905f036108fe565b90565b61647661354f565b5061648761648261357d565b616460565b90565b61649b906164966167c1565b61649d565b565b6164ae906164a96168a9565b6164b0565b565b6164bc816101276135f2565b6164f27f37a332914fac995349420c0419b4423a19dcb762017f691442a0782ce4bf417a916164e961081b565b91829182610a62565b0390a1565b6165009061648a565b565b6165139061650e6167c1565b616515565b565b8061653061652a6165255f613764565b61088b565b9161088b565b146165405761653e90619715565b565b61656361654c5f613764565b5f918291631e4fbdf760e01b8352600483016108a4565b0390fd5b61657090616502565b565b6165839061657e6167c1565b616585565b565b616596906165916168a9565b616598565b565b6165a19061a253565b565b6165ac90616572565b565b6165bf906165ba6167c1565b6165c1565b565b6165d2906165cd6168a9565b6165d4565b565b6165e0816101126135f2565b6166167fdeb71e1d117914859ddde484a9810084d9ee399077d5cd8dcbdf8770d33d8ae49161660d61081b565b91829182610a62565b0390a1565b616624906165ae565b565b61662e613741565b5061663761354f565b506166406193be565b9061664961357d565b90565b90565b5f61667061666b61667693616662613741565b5061010361144d565b61664c565b01610eb3565b90565b90565b6166ab916166a16166a69261668f613741565b50616698613741565b50610104610e3a565b610e84565b616679565b906166c360016166bc5f8501610eb3565b9301610eb3565b90565b6166d7906166d26167c1565b6166d9565b565b6166ea906166e56168a9565b6166ec565b565b6166f88161011e6135f2565b61672e7fbfcd8fb930a57c1598c9760db19c84ec766546c3c9a8565611df8302482bfb179161672561081b565b91829182610a62565b0390a1565b61673c906166c6565b565b616746616f8a565b61674e61678e565b616756616fcf565b565b90565b61676f61676a61677492616758565b610e1b565b6108fe565b90565b6167805f612580565b90565b61678b616777565b90565b6167b55f1961679d5f9161675b565b906167af6167a9616783565b9161377e565b91619b49565b565b6167bf61673e565b565b6167c9615730565b6167e26167dc6167d76177a1565b61088b565b9161088b565b036167e957565b61680b6167f46177a1565b5f91829163118cdaa760e01b8352600483016108a4565b0390fd5b60207f65616479206163746976652e0000000000000000000000000000000000000000917f5468652063757272656e742062696464696e6720726f756e6420697320616c725f8201520152565b616869602c6040926125fc565b6168728161680f565b0190565b9160406168a79294936168a0616895606083018381035f85015261685c565b966020830190610a55565b0190610a55565b565b6168b461010d610eb3565b6168d1426168ca6168c48461082d565b9161082d565b1015614567565b6168d85750565b42906168fb6168e561081b565b92839263d0fd11df60e01b845260048401616876565b0390fd5b90565b61691661691161691b926168ff565b610e1b565b61082d565b90565b90565b60ff1690565b61693b6169366169409261691e565b610e1b565b616921565b90565b6169629061695c61695661696794616921565b9161082d565b90610b8a565b61082d565b90565b90565b61698161697c6169869261696a565b610e1b565b616921565b90565b6169a8906169a261699c6169ad94616921565b9161082d565b90613e04565b61082d565b90565b90565b6169c76169c26169cc926169b0565b610e1b565b61082d565b90565b90565b6169e66169e16169eb926169cf565b610e1b565b616921565b90565b90565b616a05616a00616a0a926169ee565b610e1b565b61082d565b90565b90565b616a24616a1f616a2992616a0d565b610e1b565b616921565b90565b90565b616a43616a3e616a4892616a2c565b610e1b565b61082d565b90565b90565b616a62616a5d616a6792616a4b565b610e1b565b616921565b90565b90565b616a81616a7c616a8692616a6a565b610e1b565b61082d565b90565b90565b616aa0616a9b616aa592616a89565b610e1b565b616921565b90565b616abc616ab7616ac192616a0d565b610e1b565b61082d565b90565b616ad8616ad3616add92614117565b610e1b565b616921565b90565b616af4616aef616af992616a89565b610e1b565b61082d565b90565b616b10616b0b616b1592614143565b610e1b565b616921565b90565b90565b616b2f616b2a616b3492616b18565b610e1b565b61082d565b90565b616b3f613741565b5080616b54616b4e6001614146565b9161082d565b1115616e9f5780616d69616d46616d36616d26616d16616d06616cf6616ce6616cd6616cc6616cb68b616cb0616ca9616d6f9f616c89616c79616c9992616b9b6001614146565b9080616bb3616bad600160801b616902565b9161082d565b1015616e71575b80616bd6616bd0680100000000000000006169b3565b9161082d565b1015616e43575b80616bf5616bef6401000000006169f1565b9161082d565b1015616e15575b80616c12616c0c62010000616a2f565b9161082d565b1015616de7575b80616c2e616c28610100616a6d565b9161082d565b1015616db9575b80616c49616c436010616aa8565b9161082d565b1015616d8b575b616c63616c5d6004616ae0565b9161082d565b1015616d72575b616c746003616b1b565b6137d2565b616c836001616afc565b90616943565b616c9381866137f4565b90614170565b616ca36001616afc565b90616943565b80926137f4565b90614170565b616cc06001616afc565b90616943565b616cd0818c6137f4565b90614170565b616ce06001616afc565b90616943565b616cf0818a6137f4565b90614170565b616d006001616afc565b90616943565b616d1081886137f4565b90614170565b616d206001616afc565b90616943565b616d3081866137f4565b90614170565b616d406001616afc565b90616943565b91616d63616d5d616d588580946137f4565b61082d565b9161082d565b1161a29a565b90614162565b90565b616d8690616d806001616afc565b90616989565b616c6a565b616da2616db391616d9c6004616a8c565b90616943565b91616dad6002616ac4565b90616989565b90616c50565b616dd0616de191616dca6008616a4e565b90616943565b91616ddb6004616a8c565b90616989565b90616c35565b616dfe616e0f91616df86010616a10565b90616943565b91616e096008616a4e565b90616989565b90616c19565b616e2c616e3d91616e2660206169d2565b90616943565b91616e376010616a10565b90616989565b90616bfc565b616e5a616e6b91616e54604061696d565b90616943565b91616e6560206169d2565b90616989565b90616bdd565b616e88616e9991616e826080616927565b90616943565b91616e93604061696d565b90616989565b90616bba565b90565b5f7f5468652070726f76696465642061646472657373206973207a65726f2e000000910152565b616ed6601d6020926125fc565b616edf81616ea2565b0190565b616ef89060208101905f818303910152616ec9565b90565b616f15616f0f616f0a5f613764565b61088b565b9161088b565b14616f1c57565b616f2461081b565b63eac0d38960e01b815280616f3b60048201616ee3565b0390fd5b90565b616f56616f51616f5b92616f3f565b6135d4565b611a9f565b90565b616f877f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00616f42565b90565b616f9261a2ac565b616fb357616fb1616fa9616fa4616f5e565b61a2e0565b60019061a2f5565b565b5f633ee5aeb560e01b815280616fcb60048201610976565b0390fd5b616fe9616fe2616fdd616f5e565b61a2e0565b5f9061a2f5565b565b60407f642e000000000000000000000000000000000000000000000000000000000000917f5468652063757272656e742043535420626964207072696365206973206772655f8201527f61746572207468616e20746865206d6178696d756d20796f7520616c6c6f776560208201520152565b61706b60426060926125fc565b61707481616feb565b0190565b9160406170a99294936170a2617097606083018381035f85015261705e565b966020830190610a55565b0190610a55565b565b6170b76170bc91610e9a565b61339f565b90565b6170c990546170ab565b90565b9160206170ed9294936170e660408201965f830190610897565b0190610a55565b565b67ffffffffffffffff81116171075760208091020190565b610c0e565b9061711e617119836170ef565b610c4b565b918252565b61712d6040610c4b565b90565b5f90565b5f90565b617140617123565b906020808361714d617130565b815201617158617134565b81525050565b617166617138565b90565b5f5b82811061717757505050565b60209061718261715e565b818401520161716b565b906171b16171998361710c565b926020806171a786936170ef565b9201910390617169565b565b5190565b906171c1826171b3565b8110156171d2576020809102010190565b6127dd565b906171e19061088b565b9052565b6171ee906108fe565b600160ff1b81146171fe575f0390565b6145eb565b9061720d906108fe565b9052565b60209181520190565b60200190565b6172299061088b565b9052565b617236906108fe565b9052565b9060208061725c936172525f8201515f860190617220565b015191019061722d565b565b9061726b8160409361723a565b0190565b60200190565b9061729261728c617285846171b3565b8093617211565b9261721a565b905f5b8181106172a25750505090565b9091926172bb6172b5600192865161725e565b9461726f565b9101919091617295565b6172da9160208201915f818403910152617275565b90565b6172e7600261411a565b90565b6172f39061675b565b9052565b919461734461733961734e9360a09661732c6173559a9c9b999c61732260c08a01945f8b01906172ea565b6020890190610901565b8682036040880152612610565b986060850190610a55565b6080830190610a55565b0190610a55565b565b6173686173635f61379a565b613816565b926173878461737f6173798461082d565b9161082d565b101515614567565b617782575061739d6173985f61379a565b615bc8565b906173bc826173b46173ae8461082d565b9161082d565b111515614567565b61775c5750826173d46173ce5f61377e565b9161082d565b115f146176c0576173ed6173e8600261411a565b61718c565b6174136173f86177a1565b5f61740c846174068361377e565b906171b7565b51016171d7565b61744361742761742284613553565b6171e5565b602061743c846174365f61377e565b906171b7565b5101617203565b61746a61744e6177a1565b5f6174638461745d6001614146565b906171b7565b51016171d7565b61749361747685613553565b602061748c846174866001614146565b906171b7565b5101617203565b6174a66174a16101296170bf565b6133e0565b9063b355121490823b156176bb576174dd926174d25f80946174c661081b565b96879586948593613ac1565b8352600483016172c5565b03925af180156176b65761768a575b505b6175358161752f600161751f61751161010461750b61010b610eb3565b90610e3a565b6175196177a1565b90610e84565b019161752a83610eb3565b615dd8565b906135f2565b617541426101146135f2565b617567617556826175506172dd565b906151da565b617561610118610eb3565b9061a2f8565b617573816101166135f2565b61757e61010261283a565b61759861759261758d5f613764565b61088b565b9161088b565b14617678575b506175b26175aa6177a1565b61010261399c565b6175db6175c0610115610eb3565b6175d5816175cf610133610eb3565b90616270565b90615dd8565b6175e7816101156135f2565b6175f08361a468565b6175fb61010b610eb3565b906176046177a1565b926176736176135f1992613553565b925f19969790617624610124610eb3565b9161766161765b6176557f1d1f406c89e99504a7222aaba50bf96b9f91f522ebc6ea825255145919e102ec99610e1e565b99610e78565b9961675b565b9961766a61081b565b968796876172f7565b0390a4565b617684906101176135f2565b5f61759e565b6176a9905f3d81116176af575b6176a18183610c22565b810190613ac7565b5f6174ec565b503d617697565b613b2d565b613abd565b6176d36176ce6101296170bf565b6133e0565b639dc29fac6176e06177a1565b8392803b15617757576177065f80946177116176fa61081b565b97889687958694613ac1565b8452600484016170cc565b03925af1801561775257617726575b506174ee565b617745905f3d811161774b575b61773d8183610c22565b810190613ac7565b5f617720565b503d617733565b613b2d565b613abd565b9061777e61776861081b565b92839263814ac7ff60e01b845260048401617078565b0390fd5b8361779d5f9283926335465b3160e01b845260048401613478565b0390fd5b6177a9613509565b503390565b60207f207468652063757272656e742062696464696e6720726f756e642e0000000000917f41206269642068617320616c7265616479206265656e20706c6163656420696e5f8201520152565b617808603b6040926125fc565b617811816177ae565b0190565b61782a9060208101905f8183039101526177fb565b90565b61785c61783b61010161283a565b61785561784f61784a5f613764565b61088b565b9161088b565b1415614567565b61786257565b61786a61081b565b634283f4b960e01b81528061788160048201617815565b0390fd5b6178918161010d6135f2565b6178c77f9a2159c1f277ddd727551baedc6a6c4cba77cc5219c8563ee3b15fb67548d89b916178be61081b565b91829182610a62565b0390a1565b61790260026178fc6178eb6101046178e561010b610eb3565b90610e3a565b6178f661010161283a565b90610e84565b01610eb3565b61790d428290615db3565b61791861010561283a565b61793261792c6179275f613764565b61088b565b9161088b565b145f146179695761795e6179669261795661794e61010161283a565b61010561399c565b6101066135f2565b6101076135f2565b5b565b8061798661798061797b610107610eb3565b61082d565b9161082d565b11617993575b5050617967565b6179eb826179b76179b26179f3956179ac610107610eb3565b90615dd8565b6179fa565b6179cd6179c5610107610eb3565b6101086135f2565b6179e36179db61010161283a565b61010561399c565b6101066135f2565b6101076135f2565b5f8061798c565b617a2390617a1d617a0c610106610eb3565b617a17610108610eb3565b90615dd8565b90615db3565b617a2c81613553565b617a50617a4a617a45617a4061010a610eb3565b613553565b6108fe565b916108fe565b13617a59575b50565b617a7b90617a73617a6b61010561283a565b61010961399c565b61010a6135f2565b5f617a56565b617a8b6020610c4b565b90565b5f90565b617a9a617a81565b90602082617aa6617a8e565b81525050565b617ab4617a92565b90565b90617ac19061082d565b9052565b67ffffffffffffffff8111617add5760208091020190565b610c0e565b90617af4617aef83617ac5565b610c4b565b918252565b617b036040610c4b565b90565b617b0e617af9565b9060208083617b1b617130565b815201617b26617a8e565b81525050565b617b34617b06565b90565b5f5b828110617b4557505050565b602090617b50617b2c565b8184015201617b39565b90617b7f617b6783617ae2565b92602080617b758693617ac5565b9201910390617b37565b565b5190565b90617b8f82617b81565b811015617ba0576020809102010190565b6127dd565b617bae9061082d565b5f8114617bbc576001900390565b6145eb565b617bcd617bd39161082d565b9161082d565b908115617bde570690565b6137e0565b90505190617bf082610830565b565b90602082820312617c0b57617c08915f01617be3565b90565b610825565b60209181520190565b60200190565b617c289061082d565b9052565b90602080617c4e93617c445f8201515f860190617220565b0151910190617c1f565b565b90617c5d81604093617c2c565b0190565b60200190565b90617c84617c7e617c7784617b81565b8093617c10565b92617c19565b905f5b818110617c945750505090565b909192617cad617ca76001928651617c50565b94617c61565b9101919091617c87565b617cdb617ce8949293617cd160608401955f850190610a55565b6020830190610897565b6040818403910152617c67565b90565b617cf7617cfc91610e9a565b61246e565b90565b617d099054617ceb565b90565b5f9060033d11617d19575b565b905060045f803e617d2a5f51610815565b90617d17565b5f5f9160233d11617d3e575b565b915050602060045f3e6001905f5190617d3c565b90565b617d69617d64617d6e92617d52565b610e1b565b61082d565b90565b617d7b6012617d55565b90565b905090565b617d8e5f8092617d7e565b0190565b617d9b90617d83565b90565b90617db0617dab8361194c565b610c4b565b918252565b606090565b3d5f14617dd557617dca3d617d9e565b903d5f602084013e5b565b617ddd617db5565b90617dd3565b5f7f455448207472616e7366657220746f2063686172697479206661696c65642e00910152565b617e17601f6020926125fc565b617e2081617de3565b0190565b9190617e47906020617e3f604086018681035f880152617e0a565b940190610a55565b565b60207f696e207072697a652062656e6566696369617279206661696c65642e00000000917f455448207472616e7366657220746f2062696464696e6720726f756e64206d615f8201520152565b617ea3603c6040926125fc565b617eac81617e49565b0190565b916040617ee1929493617eda617ecf606083018381035f850152617e96565b966020830190610897565b0190610a55565b565b617ef7617ef2617efc92616b18565b610e1b565b616921565b90565b617f13617f0e617f1892616921565b610e1b565b61082d565b90565b617f27617f2c91610e9a565b611343565b90565b617f399054617f1b565b90565b617f46905161082d565b90565b90565b617f60617f5b617f6592617f49565b610e1b565b61082d565b90565b67ffffffffffffffff8111617f805760208091020190565b610c0e565b90505190617f9282610ae1565b565b90929192617fa9617fa482617f68565b610c4b565b9381855260208086019202830192818411617fe657915b838310617fcd5750505050565b60208091617fdb8486617f85565b815201920191617fc0565b6110a0565b9080601f830112156180095781602061800693519101617f94565b90565b610bfc565b9060208282031261803e575f82015167ffffffffffffffff8111618039576180369201617feb565b90565b610829565b610825565b5190565b9061805961805483617f68565b610c4b565b918252565b369037565b9061808861807083618047565b9260208061807e8693617f68565b920191039061805e565b565b67ffffffffffffffff81116180a25760208091020190565b610c0e565b906180b96180b48361808a565b610c4b565b918252565b6180c86040610c4b565b90565b6180d36180be565b90602080836180e0617130565b8152016180eb617a8e565b81525050565b6180f96180cb565b90565b5f5b82811061810a57505050565b6020906181156180f1565b81840152016180fe565b9061814461812c836180a7565b9260208061813a869361808a565b92019103906180fc565b565b5190565b9061815482618146565b811015618165576020809102010190565b6127dd565b9061817482618043565b811015618185576020809102010190565b6127dd565b618194905161088b565b90565b60209181520190565b60200190565b906020806181c8936181be5f8201515f860190617220565b0151910190617c1f565b565b906181d7816040936181a6565b0190565b60200190565b906181fe6181f86181f184618146565b8093618197565b926181a0565b905f5b81811061820e5750505090565b90919261822761822160019286516181ca565b946181db565b9101919091618201565b6182469160208201915f8184039101526181e1565b90565b61825561825a91610e9a565b61182e565b90565b6182679054618249565b90565b90565b61828161827c6182869261826a565b610e1b565b61082d565b90565b60209181520190565b60200190565b906182a581602093617220565b0190565b60200190565b906182cc6182c66182bf84618043565b8093618289565b92618292565b905f5b8181106182dc5750505090565b9091926182f56182ef6001928651618298565b946182a9565b91019190916182cf565b93929061832a6040916183329461831d60608901925f8a0190610a55565b87820360208901526182af565b940190610a55565b565b61833d90614567565b9052565b60409061836a618371949695939661836060608401985f850190618334565b6020830190610a55565b0190610a55565b565b61837b617aac565b9061838f61838761a667565b5f8401617ab7565b6183ae6183a96101036183a361010b610eb3565b9061144d565b61664c565b906183b7613741565b506183c0614eb5565b916183c96140da565b6183d1613a40565b926183da615325565b6183e26159b3565b956183ee61011f610eb3565b9361840b618406866184006001614146565b90615dd8565b617b5a565b9661846e61845c61845461841e5f61377e565b61844d61842c8d8c90617b85565b5161844361843b61010961283a565b5f83016171d7565b6020889101617ab7565b8590615dd8565b9a8890616270565b996184688b89906151da565b90615dd8565b60015b15618546575b5f9661848290617ba5565b96898861848e91617b85565b518c6184999061a76a565b8a600101908b5f016184aa90610eb3565b6184b391617bc1565b6184bc9161350d565b6184c59061283a565b9081815f01906184d4916171d7565b8c90602001906184e391617ab7565b6184ee61010b610eb3565b8991908d927f9c62e2cb8cbd10bf3b8a5760977d719fc3475fa67cb7ef9d2e1701f275e885c49161851e90610e1e565b9261852890610e78565b9361853161081b565b91829161853e9183613478565b0390a3618471565b866185596185535f61377e565b9161082d565b1161847757909397929695506020919499985061857f61857a61012c613ab0565b611278565b906185b86387565d149192919261859761010b610eb3565b936185c36185a36177a1565b976185ac61081b565b98899788968795613ac1565b855260048501617cb7565b03925af19081156192a4575f91619276575b50946185ea6185e561012e617cff565b6124af565b9063b6b55f25909190919061860061010b610eb3565b90803b15619271576186255f936186309561861961081b565b96879586948593613ac1565b835260048301610a62565b03925af19081619245575b50155f1461924057600161864d617d0c565b634e487b7114619203575b6191fe575b5f8061866a61013161283a565b8361867361081b565b908161867e81617d92565b03925af161868a617dba565b505f146191ac5761869c61013161283a565b6186db6186c97f1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d92610e78565b926186d261081b565b91829182610a62565b0390a25b6187135f806186ec6177a1565b866186f561081b565b908161870081617d92565b03925af161870c617dba565b5015614567565b61917e5761872261010261283a565b61873c6187366187315f613764565b61088b565b9161088b565b14155f146191695761876b61875a6187546004616a8c565b5b617eff565b618765610120610eb3565b90615dd8565b9261877f61877a61012d617f2f565b611384565b5f63e36aee7891618791610121610eb3565b906187e26187a0848c01617f3c565b6187c97f7c6eeb003d4a6dc5ebf549935c6ffb814ba1f060f1af8a0b11c2aa94a8e716e4617f4c565b18946187ed6187d661081b565b96879586948594613ac1565b845260048401613478565b03915afa801561916457618813915f91619142575b509461880d86618043565b90615dd8565b9661881d88618063565b9561883a6188358a61882f6001614146565b90615dd8565b61811f565b976188746188498a8c9061814a565b5161886061885861012f61283a565b5f83016171d7565b602061886d610130610eb3565b9101617ab7565b61887d87618043565b5b8061889161888b5f61377e565b9161082d565b111561890e576188a090617ba5565b9587876188ac9161816a565b6188b59061818a565b9a6188bf90617ba5565b9a8a8c6188cb9161814a565b5181815f01906188da916171d7565b6188e561011c610eb3565b90602001906188f391617ab7565b898c6188fe9161816a565b90618908916171d7565b9561887e565b50919498909295979693986189ae618927610120610eb3565b5b6189a961896761896161895c61893d8961a76a565b61895660018901916189505f8b01610eb3565b90617bc1565b9061350d565b61283a565b95617ba5565b946189976189768d889061814a565b51618983835f83016171d7565b602061899061011c610eb3565b9101617ab7565b6189a4899187909261816a565b6171d7565b617ba5565b90816189c26189bc5f61377e565b9161082d565b11156189e357906189a961896761896161895c6189ae949350505050618928565b50506189f1618a4f91617ba5565b618a2a6189ff89839061814a565b51618a16618a0e61010961283a565b5f83016171d7565b6020618a2361011c610eb3565b9101617ab7565b618a4a618a3861010961283a565b618a45879184909261816a565b6171d7565b617ba5565b618a88618a5d88839061814a565b51618a74618a6c61010561283a565b5f83016171d7565b6020618a8161011c610eb3565b9101617ab7565b618aa8618a9661010561283a565b618aa3869184909261816a565b6171d7565b618abb618ab56001614146565b9161082d565b115f1461913d57618b04618ad987618ad36001614146565b9061814a565b51618af0618ae861010261283a565b5f83016171d7565b6020618afd61011c610eb3565b9101617ab7565b618b2b618b1261010261283a565b618b2685618b206001614146565b9061816a565b6171d7565b5b618b6a618b4287618b3c5f61377e565b9061814a565b51618b56618b4e6177a1565b5f83016171d7565b6020618b6361011c610eb3565b9101617ab7565b618b8d618b756177a1565b618b8885618b825f61377e565b9061816a565b6171d7565b618ba0618b9b6101296170bf565b6133e0565b63b33266da87823b1561913857618bd692618bcb5f8094618bbf61081b565b96879586948593613ac1565b835260048301618231565b03925af1801561913357619107575b506020618bfb618bf661012b61825d565b61186f565b636578f11390618c5e5f618c1061010b610eb3565b93618c69618c20838b9901617f3c565b618c497f2a8612ecb5cb17da87f8befda0480288e2d053de55d9d7d4dc4899077cf5aeda61826d565b18618c5261081b565b98899788968795613ac1565b8552600485016182ff565b03925af1801561910257618c88915f916190d4575b5097949792618043565b93618c9d618c97848790615dd8565b97618043565b945b85618cb2618cac5f61377e565b9161082d565b1115618d6b5790618cc6618cf09992617ba5565b95618cd288889061814a565b5198618ce8618ce25f8c0161818a565b92617ba5565b9a8b91617ba5565b99618cfc61010b610eb3565b618d0c6020600194959301617f3c565b938c93618d60618d4e618d48618d427f27c21fe4cea1a3367aa491829dd4dd824296c00910626150464cba8ea5ebb3f496610e1e565b96610e78565b96610e1e565b96618d5761081b565b93849384618341565b0390a4909794618c9f565b909195929693979450618d89618d82610120610eb3565b915b617ba5565b618d9484829061814a565b5191618db3618dad618da75f860161818a565b92617ba5565b96617ba5565b92618dbf61010b610eb3565b905f91618dd060208a959301617f3c565b938693618e24618e12618e0c618e067f27c21fe4cea1a3367aa491829dd4dd824296c00910626150464cba8ea5ebb3f496610e1e565b96610e78565b96610e1e565b96618e1b61081b565b93849384618341565b0390a484618e3a618e345f61377e565b9161082d565b1115618e4c57618d8990949194618d84565b619052945090618e5e618f0192617ba5565b90618e74618e6d86849061814a565b5191617ba5565b93618e8061010b610eb3565b90618e8c61011f610eb3565b91618ea56020618e9d5f870161818a565b939501617f3c565b938793618ef9618ee7618ee1618edb7faa858ae20a26d00a9ea528972d537e68a51a0744226d9ea1fc9b22492dc282a596610e1e565b96610e78565b96610e1e565b96618ef061081b565b93849384610efc565b0390a4617ba5565b618f16618f0f84839061814a565b5192617ba5565b91618f2261010b610eb3565b618f396020618f325f850161818a565b9301617f3c565b918491618f8d618f7b618f75618f6f7f838ec9dd2530548892bff113f5ffb0138d2efc63c7f59bb6571e8c923b63126094610e1e565b94610e78565b94610e1e565b94618f8461081b565b91829182610a62565b0390a4618fa3618f9d6001614146565b9161082d565b115f146190ce57618fc8618fc183618fbb6001614146565b9061814a565b5191617ba5565b90618fd461010b610eb3565b90618fec6020618fe55f840161818a565b9201617f3c565b929161903f61902d6190276190217f3901b6430c99dc290ee88ff84c4de6091ad7eac335b58e92ef5cbb0793abf4f694610e1e565b94610e78565b94610e1e565b9461903661081b565b91829182610a62565b0390a45b61904c5f61377e565b9061814a565b519261905f61010b610eb3565b90619075602061906d6177a1565b949601617f3c565b9093946190c96190b76190b16190ab7f8c551ec2b6f186753e27f1cf46f84b57f4f83f721e8c1e6170ae512845ced59196610e1e565b96610e78565b96610e1e565b966190c061081b565b93849384610efc565b0390a4565b50619043565b6190f5915060203d81116190fb575b6190ed8183610c22565b810190617bf2565b5f618c7e565b503d6190e3565b613b2d565b619126905f3d811161912c575b61911e8183610c22565b810190613ac7565b5f618be5565b503d619114565b613b2d565b613abd565b618b2c565b61915e91503d805f833e6191568183610c22565b81019061800e565b5f618802565b613b2d565b61876b61875a6191796003617ee3565b618755565b826191876177a1565b6191a861919261081b565b928392630aa7db6360e11b845260048401617eb0565b0390fd5b6191b761013161283a565b6191f66191e47f154fb6c686c977314af35b730a16b83facda5265b2abec7237c730f63e42215a92610e78565b926191ed61081b565b91829182617e24565b0390a26186df565b613b2d565b61920b617d30565b90619217575b50618658565b90505f908061923561922f61922a617d71565b61082d565b9161082d565b03156192115761a7a6565b61865d565b619264905f3d811161926a575b61925c8183610c22565b810190613ac7565b5f61863b565b503d619252565b613abd565b619297915060203d811161929d575b61928f8183610c22565b810190617bf2565b5f6185d5565b503d619285565b613b2d565b6192b29061082d565b5f1981146192c05760010190565b6145eb565b6192d96192d15f613764565b61010161399c565b6192ed6192e55f613764565b61010261399c565b6193016192f95f613764565b61010561399c565b61931561930d5f61377e565b6101086135f2565b6193296193215f613764565b61010961399c565b61934661933e6193395f1961675b565b6137b6565b61010a6135f2565b61936461935c61935761010b610eb3565b6192a9565b61010b6135f2565b61939f61939a619375610125610eb3565b619394619383610125610eb3565b61938e610126610eb3565b90616270565b90615dd8565b61a165565b6193bc6193b7426193b161010c610eb3565b90615dd8565b617885565b565b6193c6613741565b506193e66193d5610125610eb3565b6193e061010e610eb3565b906137f4565b90565b6193f290610e6c565b90565b6193fe306193e9565b61943061942a7f000000000000000000000000000000000000000000000000000000000000000061088b565b9161088b565b14801561945a575b61943e57565b5f63703e46dd60e11b81528061945660048201610976565b0390fd5b5061946361a7b6565b61949561948f7f000000000000000000000000000000000000000000000000000000000000000061088b565b9161088b565b1415619438565b6194ad906194a86167c1565b6194af565b565b506194b86168a9565b565b6194c39061949c565b565b6194ce90610e50565b90565b6194da906194c5565b90565b6194e690610e6c565b90565b6194f281611a9f565b036194f957565b5f80fd5b9050519061950a826194e9565b565b9060208282031261952557619522915f016194fd565b90565b610825565b9190619558602061954261953d866194d1565b6194dd565b6352d1902d9061955061081b565b938492613ac1565b8252818061956860048201610976565b03915afa80915f926195f8575b50155f146195a957505090600161958a57505b565b6195a5905f918291634c9c8ce360e01b8352600483016108a4565b0390fd5b92836195c46195be6195b9614cdc565b611a9f565b91611a9f565b036195d9576195d492935061a7dc565b619588565b6195f4845f918291632a87526960e21b835260048301611aaf565b0390fd5b61961a91925060203d8111619621575b6196128183610c22565b81019061950c565b905f619575565b503d619608565b619631306193e9565b61966361965d7f000000000000000000000000000000000000000000000000000000000000000061088b565b9161088b565b0361966a57565b5f63703e46dd60e11b81528061968260048201610976565b0390fd5b61969a61969561969f92614143565b610e1b565b614f04565b90565b6196cf6196ad61a869565b6196c86196c26196bd6001619686565b614f04565b91614f04565b1415614567565b6196d557565b5f63f92ee8a960e01b8152806196ed60048201610976565b0390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0090565b61971d61a141565b61973561972b5f830161283a565b915f84910161399c565b906197696197637f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093610e78565b91610e78565b9161977261081b565b8061977c81610976565b0390a3565b619790619796919392936108fe565b926108fe565b91828103925f8285128183121692851391121516176197b157565b6145eb565b60407f727265642e000000000000000000000000000000000000000000000000000000917f5468652063757272656e742045544820626964207072696365206973206772655f8201527f61746572207468616e2074686520616d6f756e7420796f75207472616e73666560208201520152565b61983660456060926125fc565b61983f816197b6565b0190565b91604061987492949361986d619862606083018381035f850152619829565b966020830190610a55565b0190610a55565b565b60207f206265656e207573656420666f722062696464696e672e000000000000000000917f546869732052616e646f6d2057616c6b204e46542068617320616c72656164795f8201520152565b6198d060376040926125fc565b6198d981619876565b0190565b91906199009060206198f8604086018681035f8801526198c3565b940190610a55565b565b1561990a5750565b61992c9061991661081b565b91829163c35947c560e01b8352600483016198dd565b0390fd5b61993c61994191610e9a565b611f9b565b90565b61994e9054619930565b90565b9060208282031261996a57619967915f01617f85565b90565b610825565b60207f6e646f6d2057616c6b204e46542e000000000000000000000000000000000000917f596f7520617265206e6f7420746865206f776e6572206f6620746869732052615f8201520152565b6199c9602e6040926125fc565b6199d28161996f565b0190565b606090619a0f619a169496959396619a056199fa608085018581035f8701526199bc565b986020850190611fe8565b6040830190610a55565b0190610897565b565b9290919215619a2657505050565b619a4890619a3261081b565b938493630b81342760e31b8552600485016199d6565b0390fd5b619a56600261411a565b90565b619a6d619a68619a72926108fe565b610e1b565b6108fe565b90565b9194619ac2619ab7619acc9360a096619aaa619ad39a9c9b999c619aa060c08a01945f8b0190610901565b60208901906172ea565b8682036040880152612610565b986060850190610a55565b6080830190610a55565b0190610a55565b565b5f7f45544820726566756e64207472616e73666572206661696c65642e0000000000910152565b619b09601b6020926125fc565b619b1281619ad5565b0190565b916040619b47929493619b40619b35606083018381035f850152619afc565b966020830190610897565b0190610a55565b565b9190619b5c619b575f61379a565b613816565b91619b7b83619b73619b6d8461082d565b9161082d565b101515614567565b61a1225750619b91619b8c5f61379a565b614a56565b9280619bad619ba7619ba25f61379a565b6108fe565b916108fe565b125f1461a11457835b90619bd2619bc334613553565b619bcc84613553565b90619781565b9485619bee619be8619be35f61379a565b6108fe565b916108fe565b145f1461a079575b81619c11619c0b619c065f61379a565b6108fe565b916108fe565b125f14619f2d57619cbf619caf619cc7925b619c6986619c635f619c53619c45610104619c3f61010b610eb3565b90610e3a565b619c4d6177a1565b90610e84565b0191619c5e83610eb3565b615dd8565b906135f2565b619c7461010161283a565b619c8e619c88619c835f613764565b61088b565b9161088b565b14619f0b575b619ca981619ca3610112610eb3565b90616270565b90615dd8565b619cb96001614146565b90615dd8565b6101116135f2565b619d20619cfc619ceb619cdb610115610eb3565b619ce56001614146565b90615dd8565b619cf6610133610eb3565b906151da565b619d1a619d0a610133610eb3565b619d146001614146565b90615dd8565b90616270565b90619d2d826101156135f2565b84619d40619d3a5f61377e565b9161082d565b11619e6f575b619d4f8461a468565b619d5a61010b610eb3565b91619dd0619d6f619d696177a1565b95613553565b915f1993969790619d81610124610eb3565b91619dbe619db8619db27f1d1f406c89e99504a7222aaba50bf96b9f91f522ebc6ea825255145919e102ec99610e1e565b99610e78565b99619a59565b99619dc761081b565b96879687619a75565b0390a480619dee619de8619de35f61379a565b6108fe565b916108fe565b13619df7575b50565b619e335f80619e046177a1565b619e0d856137b6565b619e1561081b565b9081619e2081617d92565b03925af1619e2c617dba565b5015614567565b15619df457619e49619e436177a1565b916137b6565b90619e6b619e5561081b565b928392630aa7db6360e11b845260048401619b16565b0390fd5b619e82619e7d6101296170bf565b6133e0565b6340c10f19619e8f6177a1565b8792803b15619f0657619eb55f8094619ec0619ea961081b565b97889687958694613ac1565b8452600484016170cc565b03925af18015619f0157619ed5575b50619d46565b619ef4905f3d8111619efa575b619eec8183610c22565b810190613ac7565b5f619ecf565b503d619ee2565b613b2d565b613abd565b619f28619f2082619f1a619a4c565b906151da565b61010f6135f2565b619c94565b619fc390619f73619f51619f4c610119619f46876137b6565b906130bc565b610eb3565b619f63619f5d5f61377e565b9161082d565b14619f6d856137b6565b90619902565b619f7b6177a1565b906020619f91619f8c61012a619944565b611fdc565b636352211e90619fb8619fa3886137b6565b92619fac61081b565b97889485938493613ac1565b835260048301610a62565b03915afa90811561a0745761a01a619cbf93619ff6619ff0619caf95619cc7985f9161a046575b5061088b565b9161088b565b1461a00261012a619944565b61a00b886137b6565b9061a0146177a1565b92619a18565b61a04161a0276001614146565b61a03c61011961a036896137b6565b906130bc565b6135f2565b619c23565b61a067915060203d811161a06d575b61a05f8183610c22565b810190619951565b5f619fea565b503d61a055565b613b2d565b8561a09461a08e61a0895f61379a565b6108fe565b916108fe565b135f1461a0ec5761a0b061a0a9610113610eb3565b3a906151da565b61a0cb61a0c561a0bf896137b6565b9261082d565b9161082d565b111561a0d7575b619bf6565b9150935061a0e45f61379a565b93349161a0d2565b82349061a11061a0fa61081b565b92839263814ac7ff60e01b845260048401619843565b0390fd5b61a11d8461417e565b619bb6565b8261a13d5f9283926335465b3160e01b845260048401613478565b0390fd5b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930090565b61a171816101256135f2565b61a1a77f07417920574ce0bdfe987af0575c8793cc73a29d7830760ad459d0e569b5b79b9161a19e61081b565b91829182610a62565b0390a1565b61a1b461354f565b5061a1d861a1c061a1db565b61a1d361a1ce610115610eb3565b613553565b61356f565b90565b61a1e361354f565b5061a20961a1f042613553565b61a20361a1fe610114610eb3565b613553565b9061356f565b90565b61a218816101106135f2565b61a24e7fb6f6af60099e44041a78b3561ed029b98bf03fdb0efbbb2eb15e1f3d7d9230379161a24561081b565b91829182610a62565b0390a1565b61a25f8161010e6135f2565b61a2957ffdf6043c89a5f304289122dcc9f8bd78bb111b5d4f409e2fc2e6c141a1110b799161a28c61081b565b91829182610a62565b0390a1565b61a2a2613741565b50151590565b5f90565b61a2b461a2a8565b5061a2cd61a2c861a2c3616f5e565b61a2e0565b61a887565b90565b5f90565b61a2dd90611a9f565b90565b61a2f29061a2ec61a2d0565b5061a2d4565b90565b5d565b61a3219161a304613741565b508161a31861a3128361082d565b9161082d565b1191909161a894565b90565b90565b5190565b5f7f4d65737361676520697320746f6f206c6f6e672e000000000000000000000000910152565b61a35f60146020926125fc565b61a3688161a32b565b0190565b919061a38f90602061a387604086018681035f88015261a352565b940190610a55565b565b1561a3995750565b61a3bb9061a3a561081b565b91829163271d43ff60e21b83526004830161a36c565b0390fd5b60207f207368616c6c206265204554482e000000000000000000000000000000000000917f5468652066697273742062696420696e20612062696464696e6720726f756e645f8201520152565b61a419602e6040926125fc565b61a4228161a3bf565b0190565b61a43b9060208101905f81830391015261a40c565b90565b1561a44557565b61a44d61081b565b63b5a45a4960e01b81528061a4646004820161a426565b0390fd5b61a4b29061a4ac61a4a761a48361a47e8461a324565b61a327565b61a49f61a49961a49461011a610eb3565b61082d565b9161082d565b11159261a324565b61a327565b9061a391565b61a4bd61010161283a565b61a4d761a4d161a4cc5f613764565b61088b565b9161088b565b145f1461a61f5761a4e661a984565b61a5023461a4fc61a4f65f61377e565b9161082d565b1161a43e565b61a50e426101146135f2565b61a52b61a5234261a51d6141c0565b90615dd8565b6101246135f2565b61a53661010b610eb3565b429061a57761a5657f028a52641badd593b7f30072734c0b97e449213f55b5c3663756427340accd3c92610e1e565b9261a56e61081b565b91829182610a62565b0390a25b61a58e61a5866177a1565b61010161399c565b61a5e961a5b061a5ab61010361a5a561010b610eb3565b9061144d565b61664c565b5f61a5e261a5bf828401610eb3565b61a5dd61a5ca6177a1565b61a5d860018701849061350d565b61399c565b6192a9565b91016135f2565b61a61d42600261a61761a60961010461a60361010b610eb3565b90610e3a565b61a6116177a1565b90610e84565b016135f2565b565b61a6276178cc565b61a62f61a8b8565b61a57b565b61a64061a64591610e9a565b610e1e565b90565b90565b61a65f61a65a61a6649261a648565b610e1b565b616921565b90565b61a66f613741565b5061a6a061a69061a68a4361a6846001614146565b90615db3565b4061a634565b61a69a6001616afc565b90616943565b61a6b44861a6ae604061696d565b90616989565b1861a6bd61aad7565b9061a71b575b5061a6cc61ae4d565b9061a700575b5061a6db61afe7565b9061a6e5575b5090565b61a6f99061a6f360c061a64b565b90616989565b185f61a6e1565b61a7149061a70e6080616927565b90616989565b185f61a6d2565b61a73b61a7409161a72a614ca5565b5061a7356001614146565b90615db3565b61ac4b565b9061a74b575b61a6c3565b61a7549061a634565b185f61a746565b600161a767910161082d565b90565b61a79e5f61a7a39261a77a613741565b5061a79882820161a79261a78d82617f3c565b61a75b565b90617ab7565b01617f3c565b61b103565b90565b634e487b715f526020526024601cfd5b61a7be613509565b5061a7d95f61a7d361a7ce614cdc565b61b118565b0161283a565b90565b9061a7e68261b11b565b8161a8117fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b91610e78565b9061a81a61081b565b8061a82481610976565b0390a261a8308161a327565b61a84261a83c5f61377e565b9161082d565b115f1461a8565761a8529161b1a6565b505b565b505061a86061b170565b61a854565b5f90565b61a87161a865565b5061a8845f61a87e6196f1565b01614f7b565b90565b61a88f61a2a8565b505c90565b61a8ae61a8b4929361a8a4613741565b508094189161a29a565b906137d2565b1890565b61a8e861a8e061a8c6614c05565b61a8db61a8d4610124610eb3565b429061a2f8565b614170565b6101246135f2565b565b60207f20616374697665207965742e0000000000000000000000000000000000000000917f5468652063757272656e742062696464696e6720726f756e64206973206e6f745f8201520152565b61a944602c6040926125fc565b61a94d8161a8ea565b0190565b91604061a98292949361a97b61a970606083018381035f85015261a937565b966020830190610a55565b0190610a55565b565b61a98f61010d610eb3565b61a9ad4261a9a561a99f8461082d565b9161082d565b101515614567565b61a9b45750565b429061a9d761a9c161081b565b9283926302dbf17b60e31b84526004840161a951565b0390fd5b61a9ef61a9ea61a9f492613a21565b610e1b565b610880565b90565b61aa009061a9db565b90565b61aa0c90610e50565b90565b61aa189061aa03565b90565b61aa2d61aa28606461a9f7565b61aa0f565b90565b61aa3990610e6c565b90565b61aa5061aa4b61aa55926169cf565b610e1b565b61082d565b90565b60207f642e000000000000000000000000000000000000000000000000000000000000917f4172625379732e617262426c6f636b4e756d6265722063616c6c206661696c655f8201520152565b61aab260226040926125fc565b61aabb8161aa58565b0190565b61aad49060208101905f81830391015261aaa5565b90565b61aadf61a2a8565b5061aae8613741565b61aaf0617db5565b505f8061ab0361aafe61aa1b565b61aa30565b600461ab3a63a3b1b31d60e01b61ab2b61ab1b61081b565b9384926020840190815201610976565b60208201810382520382610c22565b82602082019151925af19161ab4d617dba565b8361aba0575b5061ab5e8315614567565b61ab65575b565b7fa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d61ab8e61081b565b8061ab988161aabf565b0390a161ab63565b909161abab8261a327565b61abbe61abb8602061aa3c565b9161082d565b145f1461abe8575061abe090602061abd58261a327565b818301019101617bf2565b905b5f61ab53565b919250505f9161abe2565b5f7f4172625379732e617262426c6f636b486173682063616c6c206661696c65642e910152565b61ac26602080926125fc565b61ac2f8161abf3565b0190565b61ac489060208101905f81830391015261ac1a565b90565b61ac5361a2a8565b505f8061ac5e614ca5565b9261ac67617db5565b50600461acb161ac7d61ac7861aa1b565b61aa30565b9261aca26315a03d4160e11b9161ac9261081b565b9485936020850190815201610a62565b60208201810382520382610c22565b82602082019151925af19161acc4617dba565b8361ad17575b5061acd58315614567565b61acdc575b565b7fa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d61ad0561081b565b8061ad0f8161ac33565b0390a161acda565b909161ad228261a327565b61ad3561ad2f602061aa3c565b9161082d565b145f1461ad5f575061ad5790602061ad4c8261a327565b81830101910161950c565b905b5f61acca565b919250505f9161ad59565b90565b61ad8161ad7c61ad869261ad6a565b610e1b565b610880565b90565b61ad929061ad6d565b90565b61ad9e90610e50565b90565b61adaa9061ad95565b90565b61adbf61adba606c61ad89565b61ada1565b90565b61adcb90610e6c565b90565b60207f696c65642e000000000000000000000000000000000000000000000000000000917f417262476173496e666f2e6765744761734261636b6c6f672063616c6c2066615f8201520152565b61ae2860256040926125fc565b61ae318161adce565b0190565b61ae4a9060208101905f81830391015261ae1b565b90565b61ae5561a2a8565b5061ae5e613741565b61ae66617db5565b505f8061ae7961ae7461adad565b61adc2565b600461aeaf62eadae160e51b61aea061ae9061081b565b9384926020840190815201610976565b60208201810382520382610c22565b82602082019151925af19161aec2617dba565b8361af15575b5061aed38315614567565b61aeda575b565b7fa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d61af0361081b565b8061af0d8161ae35565b0390a161aed8565b909161af208261a327565b61af3361af2d602061aa3c565b9161082d565b145f1461af5d575061af5590602061af4a8261a327565b818301019101617bf2565b905b5f61aec8565b919250505f9161af57565b60207f655570646174652063616c6c206661696c65642e000000000000000000000000917f417262476173496e666f2e6765744c3150726963696e67556e69747353696e635f8201520152565b61afc260346040926125fc565b61afcb8161af68565b0190565b61afe49060208101905f81830391015261afb5565b90565b61afef61a2a8565b5061aff8613741565b61b000617db5565b505f8061b01361b00e61adad565b61adc2565b600461b04a6377f8098360e11b61b03b61b02b61081b565b9384926020840190815201610976565b60208201810382520382610c22565b82602082019151925af19161b05d617dba565b8361b0b0575b5061b06e8315614567565b61b075575b565b7fa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d61b09e61081b565b8061b0a88161afcf565b0390a161b073565b909161b0bb8261a327565b61b0ce61b0c8602061aa3c565b9161082d565b145f1461b0f8575061b0f090602061b0e58261a327565b818301019101617bf2565b905b5f61b063565b919250505f9161b0f2565b61b1159061b10f613741565b5061b1d5565b90565b90565b803b61b12f61b1295f61377e565b9161082d565b1461b1515761b14f905f61b14961b144614cdc565b61b118565b0161399c565b565b61b16c905f918291634c9c8ce360e01b8352600483016108a4565b0390fd5b3461b18361b17d5f61377e565b9161082d565b1161b18a57565b5f63b398979f60e01b81528061b1a260048201610976565b0390fd5b5f8061b1d29361b1b4617db5565b508390602081019051915af49061b1c9617dba565b9091909161b1e7565b90565b61b1dd613741565b505f5260205f2090565b9061b1fb9061b1f4617db5565b5015614567565b5f1461b207575061b26b565b61b2108261a327565b61b22261b21c5f61377e565b9161082d565b148061b250575b61b231575090565b61b24c905f918291639996b31560e01b8352600483016108a4565b0390fd5b50803b61b26561b25f5f61377e565b9161082d565b1461b229565b61b2748161a327565b61b28661b2805f61377e565b9161082d565b115f1461b29557805190602001fd5b5f63d6bda27560e01b81528061b2ad60048201610976565b0390fdfea26469706673582212204809a810ea5458912a11a264c756cdb0d30093b2d37fc4e3efffcb95a9df1e7c64736f6c63430008220033",
}

// CosmicSignatureGameV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use CosmicSignatureGameV2MetaData.ABI instead.
var CosmicSignatureGameV2ABI = CosmicSignatureGameV2MetaData.ABI

// CosmicSignatureGameV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CosmicSignatureGameV2MetaData.Bin instead.
var CosmicSignatureGameV2Bin = CosmicSignatureGameV2MetaData.Bin

// DeployCosmicSignatureGameV2 deploys a new Ethereum contract, binding an instance of CosmicSignatureGameV2 to it.
func DeployCosmicSignatureGameV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CosmicSignatureGameV2, error) {
	parsed, err := CosmicSignatureGameV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CosmicSignatureGameV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CosmicSignatureGameV2{CosmicSignatureGameV2Caller: CosmicSignatureGameV2Caller{contract: contract}, CosmicSignatureGameV2Transactor: CosmicSignatureGameV2Transactor{contract: contract}, CosmicSignatureGameV2Filterer: CosmicSignatureGameV2Filterer{contract: contract}}, nil
}

// CosmicSignatureGameV2 is an auto generated Go binding around an Ethereum contract.
type CosmicSignatureGameV2 struct {
	CosmicSignatureGameV2Caller     // Read-only binding to the contract
	CosmicSignatureGameV2Transactor // Write-only binding to the contract
	CosmicSignatureGameV2Filterer   // Log filterer for contract events
}

// CosmicSignatureGameV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type CosmicSignatureGameV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmicSignatureGameV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CosmicSignatureGameV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmicSignatureGameV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CosmicSignatureGameV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmicSignatureGameV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CosmicSignatureGameV2Session struct {
	Contract     *CosmicSignatureGameV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CosmicSignatureGameV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CosmicSignatureGameV2CallerSession struct {
	Contract *CosmicSignatureGameV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// CosmicSignatureGameV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CosmicSignatureGameV2TransactorSession struct {
	Contract     *CosmicSignatureGameV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// CosmicSignatureGameV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type CosmicSignatureGameV2Raw struct {
	Contract *CosmicSignatureGameV2 // Generic contract binding to access the raw methods on
}

// CosmicSignatureGameV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CosmicSignatureGameV2CallerRaw struct {
	Contract *CosmicSignatureGameV2Caller // Generic read-only contract binding to access the raw methods on
}

// CosmicSignatureGameV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CosmicSignatureGameV2TransactorRaw struct {
	Contract *CosmicSignatureGameV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCosmicSignatureGameV2 creates a new instance of CosmicSignatureGameV2, bound to a specific deployed contract.
func NewCosmicSignatureGameV2(address common.Address, backend bind.ContractBackend) (*CosmicSignatureGameV2, error) {
	contract, err := bindCosmicSignatureGameV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2{CosmicSignatureGameV2Caller: CosmicSignatureGameV2Caller{contract: contract}, CosmicSignatureGameV2Transactor: CosmicSignatureGameV2Transactor{contract: contract}, CosmicSignatureGameV2Filterer: CosmicSignatureGameV2Filterer{contract: contract}}, nil
}

// NewCosmicSignatureGameV2Caller creates a new read-only instance of CosmicSignatureGameV2, bound to a specific deployed contract.
func NewCosmicSignatureGameV2Caller(address common.Address, caller bind.ContractCaller) (*CosmicSignatureGameV2Caller, error) {
	contract, err := bindCosmicSignatureGameV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2Caller{contract: contract}, nil
}

// NewCosmicSignatureGameV2Transactor creates a new write-only instance of CosmicSignatureGameV2, bound to a specific deployed contract.
func NewCosmicSignatureGameV2Transactor(address common.Address, transactor bind.ContractTransactor) (*CosmicSignatureGameV2Transactor, error) {
	contract, err := bindCosmicSignatureGameV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2Transactor{contract: contract}, nil
}

// NewCosmicSignatureGameV2Filterer creates a new log filterer instance of CosmicSignatureGameV2, bound to a specific deployed contract.
func NewCosmicSignatureGameV2Filterer(address common.Address, filterer bind.ContractFilterer) (*CosmicSignatureGameV2Filterer, error) {
	contract, err := bindCosmicSignatureGameV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2Filterer{contract: contract}, nil
}

// bindCosmicSignatureGameV2 binds a generic wrapper to an already deployed contract.
func bindCosmicSignatureGameV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CosmicSignatureGameV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmicSignatureGameV2.Contract.CosmicSignatureGameV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.CosmicSignatureGameV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.CosmicSignatureGameV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmicSignatureGameV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) UPGRADEINTERFACEVERSION() (string, error) {
	return _CosmicSignatureGameV2.Contract.UPGRADEINTERFACEVERSION(&_CosmicSignatureGameV2.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CosmicSignatureGameV2.Contract.UPGRADEINTERFACEVERSION(&_CosmicSignatureGameV2.CallOpts)
}

// BidCstRewardAmountMultiplier is a free data retrieval call binding the contract method 0xb30f5bb1.
//
// Solidity: function bidCstRewardAmountMultiplier() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) BidCstRewardAmountMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "bidCstRewardAmountMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BidCstRewardAmountMultiplier is a free data retrieval call binding the contract method 0xb30f5bb1.
//
// Solidity: function bidCstRewardAmountMultiplier() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) BidCstRewardAmountMultiplier() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.BidCstRewardAmountMultiplier(&_CosmicSignatureGameV2.CallOpts)
}

// BidCstRewardAmountMultiplier is a free data retrieval call binding the contract method 0xb30f5bb1.
//
// Solidity: function bidCstRewardAmountMultiplier() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) BidCstRewardAmountMultiplier() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.BidCstRewardAmountMultiplier(&_CosmicSignatureGameV2.CallOpts)
}

// BidMessageLengthMaxLimit is a free data retrieval call binding the contract method 0x6c0613c0.
//
// Solidity: function bidMessageLengthMaxLimit() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) BidMessageLengthMaxLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "bidMessageLengthMaxLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BidMessageLengthMaxLimit is a free data retrieval call binding the contract method 0x6c0613c0.
//
// Solidity: function bidMessageLengthMaxLimit() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) BidMessageLengthMaxLimit() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.BidMessageLengthMaxLimit(&_CosmicSignatureGameV2.CallOpts)
}

// BidMessageLengthMaxLimit is a free data retrieval call binding the contract method 0x6c0613c0.
//
// Solidity: function bidMessageLengthMaxLimit() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) BidMessageLengthMaxLimit() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.BidMessageLengthMaxLimit(&_CosmicSignatureGameV2.CallOpts)
}

// BidderAddresses is a free data retrieval call binding the contract method 0x2fb3c48f.
//
// Solidity: function bidderAddresses(uint256 roundNum) view returns(uint256 numItems)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) BidderAddresses(opts *bind.CallOpts, roundNum *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "bidderAddresses", roundNum)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BidderAddresses is a free data retrieval call binding the contract method 0x2fb3c48f.
//
// Solidity: function bidderAddresses(uint256 roundNum) view returns(uint256 numItems)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) BidderAddresses(roundNum *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.BidderAddresses(&_CosmicSignatureGameV2.CallOpts, roundNum)
}

// BidderAddresses is a free data retrieval call binding the contract method 0x2fb3c48f.
//
// Solidity: function bidderAddresses(uint256 roundNum) view returns(uint256 numItems)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) BidderAddresses(roundNum *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.BidderAddresses(&_CosmicSignatureGameV2.CallOpts, roundNum)
}

// BiddersInfo is a free data retrieval call binding the contract method 0x17887731.
//
// Solidity: function biddersInfo(uint256 roundNum, address bidderAddress) view returns(uint256 totalSpentEthAmount, uint256 totalSpentCstAmount, uint256 lastBidTimeStamp)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) BiddersInfo(opts *bind.CallOpts, roundNum *big.Int, bidderAddress common.Address) (struct {
	TotalSpentEthAmount *big.Int
	TotalSpentCstAmount *big.Int
	LastBidTimeStamp    *big.Int
}, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "biddersInfo", roundNum, bidderAddress)

	outstruct := new(struct {
		TotalSpentEthAmount *big.Int
		TotalSpentCstAmount *big.Int
		LastBidTimeStamp    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalSpentEthAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalSpentCstAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastBidTimeStamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BiddersInfo is a free data retrieval call binding the contract method 0x17887731.
//
// Solidity: function biddersInfo(uint256 roundNum, address bidderAddress) view returns(uint256 totalSpentEthAmount, uint256 totalSpentCstAmount, uint256 lastBidTimeStamp)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) BiddersInfo(roundNum *big.Int, bidderAddress common.Address) (struct {
	TotalSpentEthAmount *big.Int
	TotalSpentCstAmount *big.Int
	LastBidTimeStamp    *big.Int
}, error) {
	return _CosmicSignatureGameV2.Contract.BiddersInfo(&_CosmicSignatureGameV2.CallOpts, roundNum, bidderAddress)
}

// BiddersInfo is a free data retrieval call binding the contract method 0x17887731.
//
// Solidity: function biddersInfo(uint256 roundNum, address bidderAddress) view returns(uint256 totalSpentEthAmount, uint256 totalSpentCstAmount, uint256 lastBidTimeStamp)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) BiddersInfo(roundNum *big.Int, bidderAddress common.Address) (struct {
	TotalSpentEthAmount *big.Int
	TotalSpentCstAmount *big.Int
	LastBidTimeStamp    *big.Int
}, error) {
	return _CosmicSignatureGameV2.Contract.BiddersInfo(&_CosmicSignatureGameV2.CallOpts, roundNum, bidderAddress)
}

// CharityAddress is a free data retrieval call binding the contract method 0xafcf2fc4.
//
// Solidity: function charityAddress() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) CharityAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "charityAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CharityAddress is a free data retrieval call binding the contract method 0xafcf2fc4.
//
// Solidity: function charityAddress() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) CharityAddress() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.CharityAddress(&_CosmicSignatureGameV2.CallOpts)
}

// CharityAddress is a free data retrieval call binding the contract method 0xafcf2fc4.
//
// Solidity: function charityAddress() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) CharityAddress() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.CharityAddress(&_CosmicSignatureGameV2.CallOpts)
}

// CharityEthDonationAmountPercentage is a free data retrieval call binding the contract method 0xbe720ad5.
//
// Solidity: function charityEthDonationAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) CharityEthDonationAmountPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "charityEthDonationAmountPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CharityEthDonationAmountPercentage is a free data retrieval call binding the contract method 0xbe720ad5.
//
// Solidity: function charityEthDonationAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) CharityEthDonationAmountPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.CharityEthDonationAmountPercentage(&_CosmicSignatureGameV2.CallOpts)
}

// CharityEthDonationAmountPercentage is a free data retrieval call binding the contract method 0xbe720ad5.
//
// Solidity: function charityEthDonationAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) CharityEthDonationAmountPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.CharityEthDonationAmountPercentage(&_CosmicSignatureGameV2.CallOpts)
}

// ChronoWarriorAddress is a free data retrieval call binding the contract method 0x6b7cbe85.
//
// Solidity: function chronoWarriorAddress() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) ChronoWarriorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "chronoWarriorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChronoWarriorAddress is a free data retrieval call binding the contract method 0x6b7cbe85.
//
// Solidity: function chronoWarriorAddress() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) ChronoWarriorAddress() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.ChronoWarriorAddress(&_CosmicSignatureGameV2.CallOpts)
}

// ChronoWarriorAddress is a free data retrieval call binding the contract method 0x6b7cbe85.
//
// Solidity: function chronoWarriorAddress() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) ChronoWarriorAddress() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.ChronoWarriorAddress(&_CosmicSignatureGameV2.CallOpts)
}

// ChronoWarriorDuration is a free data retrieval call binding the contract method 0x1824d5e7.
//
// Solidity: function chronoWarriorDuration() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) ChronoWarriorDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "chronoWarriorDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChronoWarriorDuration is a free data retrieval call binding the contract method 0x1824d5e7.
//
// Solidity: function chronoWarriorDuration() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) ChronoWarriorDuration() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.ChronoWarriorDuration(&_CosmicSignatureGameV2.CallOpts)
}

// ChronoWarriorDuration is a free data retrieval call binding the contract method 0x1824d5e7.
//
// Solidity: function chronoWarriorDuration() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) ChronoWarriorDuration() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.ChronoWarriorDuration(&_CosmicSignatureGameV2.CallOpts)
}

// ChronoWarriorEthPrizeAmountPercentage is a free data retrieval call binding the contract method 0x54ada1d6.
//
// Solidity: function chronoWarriorEthPrizeAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) ChronoWarriorEthPrizeAmountPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "chronoWarriorEthPrizeAmountPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChronoWarriorEthPrizeAmountPercentage is a free data retrieval call binding the contract method 0x54ada1d6.
//
// Solidity: function chronoWarriorEthPrizeAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) ChronoWarriorEthPrizeAmountPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.ChronoWarriorEthPrizeAmountPercentage(&_CosmicSignatureGameV2.CallOpts)
}

// ChronoWarriorEthPrizeAmountPercentage is a free data retrieval call binding the contract method 0x54ada1d6.
//
// Solidity: function chronoWarriorEthPrizeAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) ChronoWarriorEthPrizeAmountPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.ChronoWarriorEthPrizeAmountPercentage(&_CosmicSignatureGameV2.CallOpts)
}

// CosmicSignatureNftStakingTotalEthRewardAmountPercentage is a free data retrieval call binding the contract method 0xf7bea078.
//
// Solidity: function cosmicSignatureNftStakingTotalEthRewardAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) CosmicSignatureNftStakingTotalEthRewardAmountPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "cosmicSignatureNftStakingTotalEthRewardAmountPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CosmicSignatureNftStakingTotalEthRewardAmountPercentage is a free data retrieval call binding the contract method 0xf7bea078.
//
// Solidity: function cosmicSignatureNftStakingTotalEthRewardAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) CosmicSignatureNftStakingTotalEthRewardAmountPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.CosmicSignatureNftStakingTotalEthRewardAmountPercentage(&_CosmicSignatureGameV2.CallOpts)
}

// CosmicSignatureNftStakingTotalEthRewardAmountPercentage is a free data retrieval call binding the contract method 0xf7bea078.
//
// Solidity: function cosmicSignatureNftStakingTotalEthRewardAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) CosmicSignatureNftStakingTotalEthRewardAmountPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.CosmicSignatureNftStakingTotalEthRewardAmountPercentage(&_CosmicSignatureGameV2.CallOpts)
}

// CstDutchAuctionBeginningBidPrice is a free data retrieval call binding the contract method 0x1b410319.
//
// Solidity: function cstDutchAuctionBeginningBidPrice() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) CstDutchAuctionBeginningBidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "cstDutchAuctionBeginningBidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CstDutchAuctionBeginningBidPrice is a free data retrieval call binding the contract method 0x1b410319.
//
// Solidity: function cstDutchAuctionBeginningBidPrice() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) CstDutchAuctionBeginningBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.CstDutchAuctionBeginningBidPrice(&_CosmicSignatureGameV2.CallOpts)
}

// CstDutchAuctionBeginningBidPrice is a free data retrieval call binding the contract method 0x1b410319.
//
// Solidity: function cstDutchAuctionBeginningBidPrice() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) CstDutchAuctionBeginningBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.CstDutchAuctionBeginningBidPrice(&_CosmicSignatureGameV2.CallOpts)
}

// CstDutchAuctionBeginningBidPriceMinLimit is a free data retrieval call binding the contract method 0xbb4b3e6f.
//
// Solidity: function cstDutchAuctionBeginningBidPriceMinLimit() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) CstDutchAuctionBeginningBidPriceMinLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "cstDutchAuctionBeginningBidPriceMinLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CstDutchAuctionBeginningBidPriceMinLimit is a free data retrieval call binding the contract method 0xbb4b3e6f.
//
// Solidity: function cstDutchAuctionBeginningBidPriceMinLimit() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) CstDutchAuctionBeginningBidPriceMinLimit() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.CstDutchAuctionBeginningBidPriceMinLimit(&_CosmicSignatureGameV2.CallOpts)
}

// CstDutchAuctionBeginningBidPriceMinLimit is a free data retrieval call binding the contract method 0xbb4b3e6f.
//
// Solidity: function cstDutchAuctionBeginningBidPriceMinLimit() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) CstDutchAuctionBeginningBidPriceMinLimit() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.CstDutchAuctionBeginningBidPriceMinLimit(&_CosmicSignatureGameV2.CallOpts)
}

// CstDutchAuctionBeginningTimeStamp is a free data retrieval call binding the contract method 0x27700481.
//
// Solidity: function cstDutchAuctionBeginningTimeStamp() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) CstDutchAuctionBeginningTimeStamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "cstDutchAuctionBeginningTimeStamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CstDutchAuctionBeginningTimeStamp is a free data retrieval call binding the contract method 0x27700481.
//
// Solidity: function cstDutchAuctionBeginningTimeStamp() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) CstDutchAuctionBeginningTimeStamp() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.CstDutchAuctionBeginningTimeStamp(&_CosmicSignatureGameV2.CallOpts)
}

// CstDutchAuctionBeginningTimeStamp is a free data retrieval call binding the contract method 0x27700481.
//
// Solidity: function cstDutchAuctionBeginningTimeStamp() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) CstDutchAuctionBeginningTimeStamp() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.CstDutchAuctionBeginningTimeStamp(&_CosmicSignatureGameV2.CallOpts)
}

// CstDutchAuctionDuration is a free data retrieval call binding the contract method 0x9302020f.
//
// Solidity: function cstDutchAuctionDuration() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) CstDutchAuctionDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "cstDutchAuctionDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CstDutchAuctionDuration is a free data retrieval call binding the contract method 0x9302020f.
//
// Solidity: function cstDutchAuctionDuration() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) CstDutchAuctionDuration() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.CstDutchAuctionDuration(&_CosmicSignatureGameV2.CallOpts)
}

// CstDutchAuctionDuration is a free data retrieval call binding the contract method 0x9302020f.
//
// Solidity: function cstDutchAuctionDuration() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) CstDutchAuctionDuration() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.CstDutchAuctionDuration(&_CosmicSignatureGameV2.CallOpts)
}

// CstDutchAuctionDurationChangeDivisor is a free data retrieval call binding the contract method 0xda9931dd.
//
// Solidity: function cstDutchAuctionDurationChangeDivisor() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) CstDutchAuctionDurationChangeDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "cstDutchAuctionDurationChangeDivisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CstDutchAuctionDurationChangeDivisor is a free data retrieval call binding the contract method 0xda9931dd.
//
// Solidity: function cstDutchAuctionDurationChangeDivisor() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) CstDutchAuctionDurationChangeDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.CstDutchAuctionDurationChangeDivisor(&_CosmicSignatureGameV2.CallOpts)
}

// CstDutchAuctionDurationChangeDivisor is a free data retrieval call binding the contract method 0xda9931dd.
//
// Solidity: function cstDutchAuctionDurationChangeDivisor() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) CstDutchAuctionDurationChangeDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.CstDutchAuctionDurationChangeDivisor(&_CosmicSignatureGameV2.CallOpts)
}

// CstPrizeAmount is a free data retrieval call binding the contract method 0x320c435c.
//
// Solidity: function cstPrizeAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) CstPrizeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "cstPrizeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CstPrizeAmount is a free data retrieval call binding the contract method 0x320c435c.
//
// Solidity: function cstPrizeAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) CstPrizeAmount() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.CstPrizeAmount(&_CosmicSignatureGameV2.CallOpts)
}

// CstPrizeAmount is a free data retrieval call binding the contract method 0x320c435c.
//
// Solidity: function cstPrizeAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) CstPrizeAmount() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.CstPrizeAmount(&_CosmicSignatureGameV2.CallOpts)
}

// DelayDurationBeforeRoundActivation is a free data retrieval call binding the contract method 0xb9cf9ba5.
//
// Solidity: function delayDurationBeforeRoundActivation() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) DelayDurationBeforeRoundActivation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "delayDurationBeforeRoundActivation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayDurationBeforeRoundActivation is a free data retrieval call binding the contract method 0xb9cf9ba5.
//
// Solidity: function delayDurationBeforeRoundActivation() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) DelayDurationBeforeRoundActivation() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.DelayDurationBeforeRoundActivation(&_CosmicSignatureGameV2.CallOpts)
}

// DelayDurationBeforeRoundActivation is a free data retrieval call binding the contract method 0xb9cf9ba5.
//
// Solidity: function delayDurationBeforeRoundActivation() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) DelayDurationBeforeRoundActivation() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.DelayDurationBeforeRoundActivation(&_CosmicSignatureGameV2.CallOpts)
}

// EnduranceChampionAddress is a free data retrieval call binding the contract method 0x9e50acc9.
//
// Solidity: function enduranceChampionAddress() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) EnduranceChampionAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "enduranceChampionAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EnduranceChampionAddress is a free data retrieval call binding the contract method 0x9e50acc9.
//
// Solidity: function enduranceChampionAddress() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) EnduranceChampionAddress() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.EnduranceChampionAddress(&_CosmicSignatureGameV2.CallOpts)
}

// EnduranceChampionAddress is a free data retrieval call binding the contract method 0x9e50acc9.
//
// Solidity: function enduranceChampionAddress() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) EnduranceChampionAddress() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.EnduranceChampionAddress(&_CosmicSignatureGameV2.CallOpts)
}

// EnduranceChampionDuration is a free data retrieval call binding the contract method 0x5863a705.
//
// Solidity: function enduranceChampionDuration() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) EnduranceChampionDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "enduranceChampionDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnduranceChampionDuration is a free data retrieval call binding the contract method 0x5863a705.
//
// Solidity: function enduranceChampionDuration() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) EnduranceChampionDuration() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.EnduranceChampionDuration(&_CosmicSignatureGameV2.CallOpts)
}

// EnduranceChampionDuration is a free data retrieval call binding the contract method 0x5863a705.
//
// Solidity: function enduranceChampionDuration() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) EnduranceChampionDuration() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.EnduranceChampionDuration(&_CosmicSignatureGameV2.CallOpts)
}

// EnduranceChampionStartTimeStamp is a free data retrieval call binding the contract method 0x250fadb6.
//
// Solidity: function enduranceChampionStartTimeStamp() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) EnduranceChampionStartTimeStamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "enduranceChampionStartTimeStamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnduranceChampionStartTimeStamp is a free data retrieval call binding the contract method 0x250fadb6.
//
// Solidity: function enduranceChampionStartTimeStamp() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) EnduranceChampionStartTimeStamp() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.EnduranceChampionStartTimeStamp(&_CosmicSignatureGameV2.CallOpts)
}

// EnduranceChampionStartTimeStamp is a free data retrieval call binding the contract method 0x250fadb6.
//
// Solidity: function enduranceChampionStartTimeStamp() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) EnduranceChampionStartTimeStamp() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.EnduranceChampionStartTimeStamp(&_CosmicSignatureGameV2.CallOpts)
}

// EthBidPriceIncreaseDivisor is a free data retrieval call binding the contract method 0xa9742016.
//
// Solidity: function ethBidPriceIncreaseDivisor() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) EthBidPriceIncreaseDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "ethBidPriceIncreaseDivisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthBidPriceIncreaseDivisor is a free data retrieval call binding the contract method 0xa9742016.
//
// Solidity: function ethBidPriceIncreaseDivisor() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) EthBidPriceIncreaseDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.EthBidPriceIncreaseDivisor(&_CosmicSignatureGameV2.CallOpts)
}

// EthBidPriceIncreaseDivisor is a free data retrieval call binding the contract method 0xa9742016.
//
// Solidity: function ethBidPriceIncreaseDivisor() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) EthBidPriceIncreaseDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.EthBidPriceIncreaseDivisor(&_CosmicSignatureGameV2.CallOpts)
}

// EthBidRefundAmountInGasToSwallowMaxLimit is a free data retrieval call binding the contract method 0x9aa1b38d.
//
// Solidity: function ethBidRefundAmountInGasToSwallowMaxLimit() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) EthBidRefundAmountInGasToSwallowMaxLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "ethBidRefundAmountInGasToSwallowMaxLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthBidRefundAmountInGasToSwallowMaxLimit is a free data retrieval call binding the contract method 0x9aa1b38d.
//
// Solidity: function ethBidRefundAmountInGasToSwallowMaxLimit() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) EthBidRefundAmountInGasToSwallowMaxLimit() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.EthBidRefundAmountInGasToSwallowMaxLimit(&_CosmicSignatureGameV2.CallOpts)
}

// EthBidRefundAmountInGasToSwallowMaxLimit is a free data retrieval call binding the contract method 0x9aa1b38d.
//
// Solidity: function ethBidRefundAmountInGasToSwallowMaxLimit() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) EthBidRefundAmountInGasToSwallowMaxLimit() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.EthBidRefundAmountInGasToSwallowMaxLimit(&_CosmicSignatureGameV2.CallOpts)
}

// EthDonationWithInfoRecords is a free data retrieval call binding the contract method 0xb5d1f06f.
//
// Solidity: function ethDonationWithInfoRecords(uint256 ) view returns(uint256 roundNum, address donorAddress, uint256 amount, string data)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) EthDonationWithInfoRecords(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RoundNum     *big.Int
	DonorAddress common.Address
	Amount       *big.Int
	Data         string
}, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "ethDonationWithInfoRecords", arg0)

	outstruct := new(struct {
		RoundNum     *big.Int
		DonorAddress common.Address
		Amount       *big.Int
		Data         string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundNum = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DonorAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// EthDonationWithInfoRecords is a free data retrieval call binding the contract method 0xb5d1f06f.
//
// Solidity: function ethDonationWithInfoRecords(uint256 ) view returns(uint256 roundNum, address donorAddress, uint256 amount, string data)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) EthDonationWithInfoRecords(arg0 *big.Int) (struct {
	RoundNum     *big.Int
	DonorAddress common.Address
	Amount       *big.Int
	Data         string
}, error) {
	return _CosmicSignatureGameV2.Contract.EthDonationWithInfoRecords(&_CosmicSignatureGameV2.CallOpts, arg0)
}

// EthDonationWithInfoRecords is a free data retrieval call binding the contract method 0xb5d1f06f.
//
// Solidity: function ethDonationWithInfoRecords(uint256 ) view returns(uint256 roundNum, address donorAddress, uint256 amount, string data)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) EthDonationWithInfoRecords(arg0 *big.Int) (struct {
	RoundNum     *big.Int
	DonorAddress common.Address
	Amount       *big.Int
	Data         string
}, error) {
	return _CosmicSignatureGameV2.Contract.EthDonationWithInfoRecords(&_CosmicSignatureGameV2.CallOpts, arg0)
}

// EthDutchAuctionBeginningBidPrice is a free data retrieval call binding the contract method 0xc87baab5.
//
// Solidity: function ethDutchAuctionBeginningBidPrice() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) EthDutchAuctionBeginningBidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "ethDutchAuctionBeginningBidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthDutchAuctionBeginningBidPrice is a free data retrieval call binding the contract method 0xc87baab5.
//
// Solidity: function ethDutchAuctionBeginningBidPrice() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) EthDutchAuctionBeginningBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.EthDutchAuctionBeginningBidPrice(&_CosmicSignatureGameV2.CallOpts)
}

// EthDutchAuctionBeginningBidPrice is a free data retrieval call binding the contract method 0xc87baab5.
//
// Solidity: function ethDutchAuctionBeginningBidPrice() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) EthDutchAuctionBeginningBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.EthDutchAuctionBeginningBidPrice(&_CosmicSignatureGameV2.CallOpts)
}

// EthDutchAuctionDurationDivisor is a free data retrieval call binding the contract method 0xd1f8fcf2.
//
// Solidity: function ethDutchAuctionDurationDivisor() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) EthDutchAuctionDurationDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "ethDutchAuctionDurationDivisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthDutchAuctionDurationDivisor is a free data retrieval call binding the contract method 0xd1f8fcf2.
//
// Solidity: function ethDutchAuctionDurationDivisor() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) EthDutchAuctionDurationDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.EthDutchAuctionDurationDivisor(&_CosmicSignatureGameV2.CallOpts)
}

// EthDutchAuctionDurationDivisor is a free data retrieval call binding the contract method 0xd1f8fcf2.
//
// Solidity: function ethDutchAuctionDurationDivisor() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) EthDutchAuctionDurationDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.EthDutchAuctionDurationDivisor(&_CosmicSignatureGameV2.CallOpts)
}

// EthDutchAuctionEndingBidPriceDivisor is a free data retrieval call binding the contract method 0xebaa1ea8.
//
// Solidity: function ethDutchAuctionEndingBidPriceDivisor() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) EthDutchAuctionEndingBidPriceDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "ethDutchAuctionEndingBidPriceDivisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthDutchAuctionEndingBidPriceDivisor is a free data retrieval call binding the contract method 0xebaa1ea8.
//
// Solidity: function ethDutchAuctionEndingBidPriceDivisor() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) EthDutchAuctionEndingBidPriceDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.EthDutchAuctionEndingBidPriceDivisor(&_CosmicSignatureGameV2.CallOpts)
}

// EthDutchAuctionEndingBidPriceDivisor is a free data retrieval call binding the contract method 0xebaa1ea8.
//
// Solidity: function ethDutchAuctionEndingBidPriceDivisor() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) EthDutchAuctionEndingBidPriceDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.EthDutchAuctionEndingBidPriceDivisor(&_CosmicSignatureGameV2.CallOpts)
}

// GetBidCstRewardAmount is a free data retrieval call binding the contract method 0xbaab4430.
//
// Solidity: function getBidCstRewardAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetBidCstRewardAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getBidCstRewardAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBidCstRewardAmount is a free data retrieval call binding the contract method 0xbaab4430.
//
// Solidity: function getBidCstRewardAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetBidCstRewardAmount() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetBidCstRewardAmount(&_CosmicSignatureGameV2.CallOpts)
}

// GetBidCstRewardAmount is a free data retrieval call binding the contract method 0xbaab4430.
//
// Solidity: function getBidCstRewardAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetBidCstRewardAmount() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetBidCstRewardAmount(&_CosmicSignatureGameV2.CallOpts)
}

// GetBidCstRewardAmountAdvanced is a free data retrieval call binding the contract method 0x0a120648.
//
// Solidity: function getBidCstRewardAmountAdvanced(int256 currentTimeOffset_) view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetBidCstRewardAmountAdvanced(opts *bind.CallOpts, currentTimeOffset_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getBidCstRewardAmountAdvanced", currentTimeOffset_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBidCstRewardAmountAdvanced is a free data retrieval call binding the contract method 0x0a120648.
//
// Solidity: function getBidCstRewardAmountAdvanced(int256 currentTimeOffset_) view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetBidCstRewardAmountAdvanced(currentTimeOffset_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetBidCstRewardAmountAdvanced(&_CosmicSignatureGameV2.CallOpts, currentTimeOffset_)
}

// GetBidCstRewardAmountAdvanced is a free data retrieval call binding the contract method 0x0a120648.
//
// Solidity: function getBidCstRewardAmountAdvanced(int256 currentTimeOffset_) view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetBidCstRewardAmountAdvanced(currentTimeOffset_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetBidCstRewardAmountAdvanced(&_CosmicSignatureGameV2.CallOpts, currentTimeOffset_)
}

// GetBidderAddressAt is a free data retrieval call binding the contract method 0x000ac9f1.
//
// Solidity: function getBidderAddressAt(uint256 roundNum_, uint256 bidIndex_) view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetBidderAddressAt(opts *bind.CallOpts, roundNum_ *big.Int, bidIndex_ *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getBidderAddressAt", roundNum_, bidIndex_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBidderAddressAt is a free data retrieval call binding the contract method 0x000ac9f1.
//
// Solidity: function getBidderAddressAt(uint256 roundNum_, uint256 bidIndex_) view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetBidderAddressAt(roundNum_ *big.Int, bidIndex_ *big.Int) (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.GetBidderAddressAt(&_CosmicSignatureGameV2.CallOpts, roundNum_, bidIndex_)
}

// GetBidderAddressAt is a free data retrieval call binding the contract method 0x000ac9f1.
//
// Solidity: function getBidderAddressAt(uint256 roundNum_, uint256 bidIndex_) view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetBidderAddressAt(roundNum_ *big.Int, bidIndex_ *big.Int) (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.GetBidderAddressAt(&_CosmicSignatureGameV2.CallOpts, roundNum_, bidIndex_)
}

// GetBidderTotalSpentAmounts is a free data retrieval call binding the contract method 0xfd9b3747.
//
// Solidity: function getBidderTotalSpentAmounts(uint256 roundNum_, address bidderAddress_) view returns(uint256, uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetBidderTotalSpentAmounts(opts *bind.CallOpts, roundNum_ *big.Int, bidderAddress_ common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getBidderTotalSpentAmounts", roundNum_, bidderAddress_)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetBidderTotalSpentAmounts is a free data retrieval call binding the contract method 0xfd9b3747.
//
// Solidity: function getBidderTotalSpentAmounts(uint256 roundNum_, address bidderAddress_) view returns(uint256, uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetBidderTotalSpentAmounts(roundNum_ *big.Int, bidderAddress_ common.Address) (*big.Int, *big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetBidderTotalSpentAmounts(&_CosmicSignatureGameV2.CallOpts, roundNum_, bidderAddress_)
}

// GetBidderTotalSpentAmounts is a free data retrieval call binding the contract method 0xfd9b3747.
//
// Solidity: function getBidderTotalSpentAmounts(uint256 roundNum_, address bidderAddress_) view returns(uint256, uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetBidderTotalSpentAmounts(roundNum_ *big.Int, bidderAddress_ common.Address) (*big.Int, *big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetBidderTotalSpentAmounts(&_CosmicSignatureGameV2.CallOpts, roundNum_, bidderAddress_)
}

// GetCharityEthDonationAmount is a free data retrieval call binding the contract method 0x0eb16be6.
//
// Solidity: function getCharityEthDonationAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetCharityEthDonationAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getCharityEthDonationAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCharityEthDonationAmount is a free data retrieval call binding the contract method 0x0eb16be6.
//
// Solidity: function getCharityEthDonationAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetCharityEthDonationAmount() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetCharityEthDonationAmount(&_CosmicSignatureGameV2.CallOpts)
}

// GetCharityEthDonationAmount is a free data retrieval call binding the contract method 0x0eb16be6.
//
// Solidity: function getCharityEthDonationAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetCharityEthDonationAmount() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetCharityEthDonationAmount(&_CosmicSignatureGameV2.CallOpts)
}

// GetChronoWarriorEthPrizeAmount is a free data retrieval call binding the contract method 0x2665c882.
//
// Solidity: function getChronoWarriorEthPrizeAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetChronoWarriorEthPrizeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getChronoWarriorEthPrizeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChronoWarriorEthPrizeAmount is a free data retrieval call binding the contract method 0x2665c882.
//
// Solidity: function getChronoWarriorEthPrizeAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetChronoWarriorEthPrizeAmount() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetChronoWarriorEthPrizeAmount(&_CosmicSignatureGameV2.CallOpts)
}

// GetChronoWarriorEthPrizeAmount is a free data retrieval call binding the contract method 0x2665c882.
//
// Solidity: function getChronoWarriorEthPrizeAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetChronoWarriorEthPrizeAmount() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetChronoWarriorEthPrizeAmount(&_CosmicSignatureGameV2.CallOpts)
}

// GetCosmicSignatureNftStakingTotalEthRewardAmount is a free data retrieval call binding the contract method 0x5f0112fe.
//
// Solidity: function getCosmicSignatureNftStakingTotalEthRewardAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetCosmicSignatureNftStakingTotalEthRewardAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getCosmicSignatureNftStakingTotalEthRewardAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCosmicSignatureNftStakingTotalEthRewardAmount is a free data retrieval call binding the contract method 0x5f0112fe.
//
// Solidity: function getCosmicSignatureNftStakingTotalEthRewardAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetCosmicSignatureNftStakingTotalEthRewardAmount() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetCosmicSignatureNftStakingTotalEthRewardAmount(&_CosmicSignatureGameV2.CallOpts)
}

// GetCosmicSignatureNftStakingTotalEthRewardAmount is a free data retrieval call binding the contract method 0x5f0112fe.
//
// Solidity: function getCosmicSignatureNftStakingTotalEthRewardAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetCosmicSignatureNftStakingTotalEthRewardAmount() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetCosmicSignatureNftStakingTotalEthRewardAmount(&_CosmicSignatureGameV2.CallOpts)
}

// GetCstDutchAuctionDurations is a free data retrieval call binding the contract method 0xb700db5f.
//
// Solidity: function getCstDutchAuctionDurations() view returns(uint256, int256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetCstDutchAuctionDurations(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getCstDutchAuctionDurations")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetCstDutchAuctionDurations is a free data retrieval call binding the contract method 0xb700db5f.
//
// Solidity: function getCstDutchAuctionDurations() view returns(uint256, int256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetCstDutchAuctionDurations() (*big.Int, *big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetCstDutchAuctionDurations(&_CosmicSignatureGameV2.CallOpts)
}

// GetCstDutchAuctionDurations is a free data retrieval call binding the contract method 0xb700db5f.
//
// Solidity: function getCstDutchAuctionDurations() view returns(uint256, int256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetCstDutchAuctionDurations() (*big.Int, *big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetCstDutchAuctionDurations(&_CosmicSignatureGameV2.CallOpts)
}

// GetDurationElapsedSinceRoundActivation is a free data retrieval call binding the contract method 0x040d4d31.
//
// Solidity: function getDurationElapsedSinceRoundActivation() view returns(int256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetDurationElapsedSinceRoundActivation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getDurationElapsedSinceRoundActivation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDurationElapsedSinceRoundActivation is a free data retrieval call binding the contract method 0x040d4d31.
//
// Solidity: function getDurationElapsedSinceRoundActivation() view returns(int256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetDurationElapsedSinceRoundActivation() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetDurationElapsedSinceRoundActivation(&_CosmicSignatureGameV2.CallOpts)
}

// GetDurationElapsedSinceRoundActivation is a free data retrieval call binding the contract method 0x040d4d31.
//
// Solidity: function getDurationElapsedSinceRoundActivation() view returns(int256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetDurationElapsedSinceRoundActivation() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetDurationElapsedSinceRoundActivation(&_CosmicSignatureGameV2.CallOpts)
}

// GetDurationUntilMainPrize is a free data retrieval call binding the contract method 0x36750d2c.
//
// Solidity: function getDurationUntilMainPrize() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetDurationUntilMainPrize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getDurationUntilMainPrize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDurationUntilMainPrize is a free data retrieval call binding the contract method 0x36750d2c.
//
// Solidity: function getDurationUntilMainPrize() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetDurationUntilMainPrize() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetDurationUntilMainPrize(&_CosmicSignatureGameV2.CallOpts)
}

// GetDurationUntilMainPrize is a free data retrieval call binding the contract method 0x36750d2c.
//
// Solidity: function getDurationUntilMainPrize() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetDurationUntilMainPrize() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetDurationUntilMainPrize(&_CosmicSignatureGameV2.CallOpts)
}

// GetDurationUntilMainPrizeRaw is a free data retrieval call binding the contract method 0x37b99cc7.
//
// Solidity: function getDurationUntilMainPrizeRaw() view returns(int256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetDurationUntilMainPrizeRaw(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getDurationUntilMainPrizeRaw")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDurationUntilMainPrizeRaw is a free data retrieval call binding the contract method 0x37b99cc7.
//
// Solidity: function getDurationUntilMainPrizeRaw() view returns(int256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetDurationUntilMainPrizeRaw() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetDurationUntilMainPrizeRaw(&_CosmicSignatureGameV2.CallOpts)
}

// GetDurationUntilMainPrizeRaw is a free data retrieval call binding the contract method 0x37b99cc7.
//
// Solidity: function getDurationUntilMainPrizeRaw() view returns(int256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetDurationUntilMainPrizeRaw() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetDurationUntilMainPrizeRaw(&_CosmicSignatureGameV2.CallOpts)
}

// GetDurationUntilRoundActivation is a free data retrieval call binding the contract method 0xef22d15b.
//
// Solidity: function getDurationUntilRoundActivation() view returns(int256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetDurationUntilRoundActivation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getDurationUntilRoundActivation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDurationUntilRoundActivation is a free data retrieval call binding the contract method 0xef22d15b.
//
// Solidity: function getDurationUntilRoundActivation() view returns(int256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetDurationUntilRoundActivation() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetDurationUntilRoundActivation(&_CosmicSignatureGameV2.CallOpts)
}

// GetDurationUntilRoundActivation is a free data retrieval call binding the contract method 0xef22d15b.
//
// Solidity: function getDurationUntilRoundActivation() view returns(int256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetDurationUntilRoundActivation() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetDurationUntilRoundActivation(&_CosmicSignatureGameV2.CallOpts)
}

// GetEthDutchAuctionDurations is a free data retrieval call binding the contract method 0xfbaf5084.
//
// Solidity: function getEthDutchAuctionDurations() view returns(uint256, int256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetEthDutchAuctionDurations(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getEthDutchAuctionDurations")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetEthDutchAuctionDurations is a free data retrieval call binding the contract method 0xfbaf5084.
//
// Solidity: function getEthDutchAuctionDurations() view returns(uint256, int256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetEthDutchAuctionDurations() (*big.Int, *big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetEthDutchAuctionDurations(&_CosmicSignatureGameV2.CallOpts)
}

// GetEthDutchAuctionDurations is a free data retrieval call binding the contract method 0xfbaf5084.
//
// Solidity: function getEthDutchAuctionDurations() view returns(uint256, int256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetEthDutchAuctionDurations() (*big.Int, *big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetEthDutchAuctionDurations(&_CosmicSignatureGameV2.CallOpts)
}

// GetEthPlusRandomWalkNftBidPrice is a free data retrieval call binding the contract method 0x27995f07.
//
// Solidity: function getEthPlusRandomWalkNftBidPrice(uint256 ethBidPrice_) pure returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetEthPlusRandomWalkNftBidPrice(opts *bind.CallOpts, ethBidPrice_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getEthPlusRandomWalkNftBidPrice", ethBidPrice_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthPlusRandomWalkNftBidPrice is a free data retrieval call binding the contract method 0x27995f07.
//
// Solidity: function getEthPlusRandomWalkNftBidPrice(uint256 ethBidPrice_) pure returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetEthPlusRandomWalkNftBidPrice(ethBidPrice_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetEthPlusRandomWalkNftBidPrice(&_CosmicSignatureGameV2.CallOpts, ethBidPrice_)
}

// GetEthPlusRandomWalkNftBidPrice is a free data retrieval call binding the contract method 0x27995f07.
//
// Solidity: function getEthPlusRandomWalkNftBidPrice(uint256 ethBidPrice_) pure returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetEthPlusRandomWalkNftBidPrice(ethBidPrice_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetEthPlusRandomWalkNftBidPrice(&_CosmicSignatureGameV2.CallOpts, ethBidPrice_)
}

// GetInitialDurationUntilMainPrize is a free data retrieval call binding the contract method 0x2b8dcbba.
//
// Solidity: function getInitialDurationUntilMainPrize() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetInitialDurationUntilMainPrize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getInitialDurationUntilMainPrize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitialDurationUntilMainPrize is a free data retrieval call binding the contract method 0x2b8dcbba.
//
// Solidity: function getInitialDurationUntilMainPrize() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetInitialDurationUntilMainPrize() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetInitialDurationUntilMainPrize(&_CosmicSignatureGameV2.CallOpts)
}

// GetInitialDurationUntilMainPrize is a free data retrieval call binding the contract method 0x2b8dcbba.
//
// Solidity: function getInitialDurationUntilMainPrize() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetInitialDurationUntilMainPrize() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetInitialDurationUntilMainPrize(&_CosmicSignatureGameV2.CallOpts)
}

// GetMainEthPrizeAmount is a free data retrieval call binding the contract method 0x5b0a45d9.
//
// Solidity: function getMainEthPrizeAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetMainEthPrizeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getMainEthPrizeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMainEthPrizeAmount is a free data retrieval call binding the contract method 0x5b0a45d9.
//
// Solidity: function getMainEthPrizeAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetMainEthPrizeAmount() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetMainEthPrizeAmount(&_CosmicSignatureGameV2.CallOpts)
}

// GetMainEthPrizeAmount is a free data retrieval call binding the contract method 0x5b0a45d9.
//
// Solidity: function getMainEthPrizeAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetMainEthPrizeAmount() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetMainEthPrizeAmount(&_CosmicSignatureGameV2.CallOpts)
}

// GetMainPrizeTimeIncrement is a free data retrieval call binding the contract method 0x4f734612.
//
// Solidity: function getMainPrizeTimeIncrement() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetMainPrizeTimeIncrement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getMainPrizeTimeIncrement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMainPrizeTimeIncrement is a free data retrieval call binding the contract method 0x4f734612.
//
// Solidity: function getMainPrizeTimeIncrement() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetMainPrizeTimeIncrement() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetMainPrizeTimeIncrement(&_CosmicSignatureGameV2.CallOpts)
}

// GetMainPrizeTimeIncrement is a free data retrieval call binding the contract method 0x4f734612.
//
// Solidity: function getMainPrizeTimeIncrement() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetMainPrizeTimeIncrement() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetMainPrizeTimeIncrement(&_CosmicSignatureGameV2.CallOpts)
}

// GetNextCstBidPrice is a free data retrieval call binding the contract method 0x6e95d286.
//
// Solidity: function getNextCstBidPrice() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetNextCstBidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getNextCstBidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextCstBidPrice is a free data retrieval call binding the contract method 0x6e95d286.
//
// Solidity: function getNextCstBidPrice() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetNextCstBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetNextCstBidPrice(&_CosmicSignatureGameV2.CallOpts)
}

// GetNextCstBidPrice is a free data retrieval call binding the contract method 0x6e95d286.
//
// Solidity: function getNextCstBidPrice() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetNextCstBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetNextCstBidPrice(&_CosmicSignatureGameV2.CallOpts)
}

// GetNextCstBidPriceAdvanced is a free data retrieval call binding the contract method 0xb6a94f42.
//
// Solidity: function getNextCstBidPriceAdvanced(int256 currentTimeOffset_) view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetNextCstBidPriceAdvanced(opts *bind.CallOpts, currentTimeOffset_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getNextCstBidPriceAdvanced", currentTimeOffset_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextCstBidPriceAdvanced is a free data retrieval call binding the contract method 0xb6a94f42.
//
// Solidity: function getNextCstBidPriceAdvanced(int256 currentTimeOffset_) view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetNextCstBidPriceAdvanced(currentTimeOffset_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetNextCstBidPriceAdvanced(&_CosmicSignatureGameV2.CallOpts, currentTimeOffset_)
}

// GetNextCstBidPriceAdvanced is a free data retrieval call binding the contract method 0xb6a94f42.
//
// Solidity: function getNextCstBidPriceAdvanced(int256 currentTimeOffset_) view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetNextCstBidPriceAdvanced(currentTimeOffset_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetNextCstBidPriceAdvanced(&_CosmicSignatureGameV2.CallOpts, currentTimeOffset_)
}

// GetNextEthBidPrice is a free data retrieval call binding the contract method 0x62ed9b53.
//
// Solidity: function getNextEthBidPrice() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetNextEthBidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getNextEthBidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextEthBidPrice is a free data retrieval call binding the contract method 0x62ed9b53.
//
// Solidity: function getNextEthBidPrice() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetNextEthBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetNextEthBidPrice(&_CosmicSignatureGameV2.CallOpts)
}

// GetNextEthBidPrice is a free data retrieval call binding the contract method 0x62ed9b53.
//
// Solidity: function getNextEthBidPrice() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetNextEthBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetNextEthBidPrice(&_CosmicSignatureGameV2.CallOpts)
}

// GetNextEthBidPriceAdvanced is a free data retrieval call binding the contract method 0x4e452010.
//
// Solidity: function getNextEthBidPriceAdvanced(int256 currentTimeOffset_) view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetNextEthBidPriceAdvanced(opts *bind.CallOpts, currentTimeOffset_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getNextEthBidPriceAdvanced", currentTimeOffset_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextEthBidPriceAdvanced is a free data retrieval call binding the contract method 0x4e452010.
//
// Solidity: function getNextEthBidPriceAdvanced(int256 currentTimeOffset_) view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetNextEthBidPriceAdvanced(currentTimeOffset_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetNextEthBidPriceAdvanced(&_CosmicSignatureGameV2.CallOpts, currentTimeOffset_)
}

// GetNextEthBidPriceAdvanced is a free data retrieval call binding the contract method 0x4e452010.
//
// Solidity: function getNextEthBidPriceAdvanced(int256 currentTimeOffset_) view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetNextEthBidPriceAdvanced(currentTimeOffset_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetNextEthBidPriceAdvanced(&_CosmicSignatureGameV2.CallOpts, currentTimeOffset_)
}

// GetRaffleTotalEthPrizeAmountForBidders is a free data retrieval call binding the contract method 0xa35286d1.
//
// Solidity: function getRaffleTotalEthPrizeAmountForBidders() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetRaffleTotalEthPrizeAmountForBidders(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getRaffleTotalEthPrizeAmountForBidders")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRaffleTotalEthPrizeAmountForBidders is a free data retrieval call binding the contract method 0xa35286d1.
//
// Solidity: function getRaffleTotalEthPrizeAmountForBidders() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetRaffleTotalEthPrizeAmountForBidders() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetRaffleTotalEthPrizeAmountForBidders(&_CosmicSignatureGameV2.CallOpts)
}

// GetRaffleTotalEthPrizeAmountForBidders is a free data retrieval call binding the contract method 0xa35286d1.
//
// Solidity: function getRaffleTotalEthPrizeAmountForBidders() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetRaffleTotalEthPrizeAmountForBidders() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetRaffleTotalEthPrizeAmountForBidders(&_CosmicSignatureGameV2.CallOpts)
}

// GetTotalNumBids is a free data retrieval call binding the contract method 0xfd77310f.
//
// Solidity: function getTotalNumBids(uint256 roundNum_) view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) GetTotalNumBids(opts *bind.CallOpts, roundNum_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "getTotalNumBids", roundNum_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalNumBids is a free data retrieval call binding the contract method 0xfd77310f.
//
// Solidity: function getTotalNumBids(uint256 roundNum_) view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) GetTotalNumBids(roundNum_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetTotalNumBids(&_CosmicSignatureGameV2.CallOpts, roundNum_)
}

// GetTotalNumBids is a free data retrieval call binding the contract method 0xfd77310f.
//
// Solidity: function getTotalNumBids(uint256 roundNum_) view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) GetTotalNumBids(roundNum_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.GetTotalNumBids(&_CosmicSignatureGameV2.CallOpts, roundNum_)
}

// InitialDurationUntilMainPrizeDivisor is a free data retrieval call binding the contract method 0x44a4b917.
//
// Solidity: function initialDurationUntilMainPrizeDivisor() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) InitialDurationUntilMainPrizeDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "initialDurationUntilMainPrizeDivisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialDurationUntilMainPrizeDivisor is a free data retrieval call binding the contract method 0x44a4b917.
//
// Solidity: function initialDurationUntilMainPrizeDivisor() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) InitialDurationUntilMainPrizeDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.InitialDurationUntilMainPrizeDivisor(&_CosmicSignatureGameV2.CallOpts)
}

// InitialDurationUntilMainPrizeDivisor is a free data retrieval call binding the contract method 0x44a4b917.
//
// Solidity: function initialDurationUntilMainPrizeDivisor() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) InitialDurationUntilMainPrizeDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.InitialDurationUntilMainPrizeDivisor(&_CosmicSignatureGameV2.CallOpts)
}

// LastBidderAddress is a free data retrieval call binding the contract method 0xe5b3cd14.
//
// Solidity: function lastBidderAddress() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) LastBidderAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "lastBidderAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastBidderAddress is a free data retrieval call binding the contract method 0xe5b3cd14.
//
// Solidity: function lastBidderAddress() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) LastBidderAddress() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.LastBidderAddress(&_CosmicSignatureGameV2.CallOpts)
}

// LastBidderAddress is a free data retrieval call binding the contract method 0xe5b3cd14.
//
// Solidity: function lastBidderAddress() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) LastBidderAddress() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.LastBidderAddress(&_CosmicSignatureGameV2.CallOpts)
}

// LastCstBidderAddress is a free data retrieval call binding the contract method 0xad4b0e8a.
//
// Solidity: function lastCstBidderAddress() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) LastCstBidderAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "lastCstBidderAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastCstBidderAddress is a free data retrieval call binding the contract method 0xad4b0e8a.
//
// Solidity: function lastCstBidderAddress() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) LastCstBidderAddress() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.LastCstBidderAddress(&_CosmicSignatureGameV2.CallOpts)
}

// LastCstBidderAddress is a free data retrieval call binding the contract method 0xad4b0e8a.
//
// Solidity: function lastCstBidderAddress() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) LastCstBidderAddress() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.LastCstBidderAddress(&_CosmicSignatureGameV2.CallOpts)
}

// MainEthPrizeAmountPercentage is a free data retrieval call binding the contract method 0xf444b298.
//
// Solidity: function mainEthPrizeAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) MainEthPrizeAmountPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "mainEthPrizeAmountPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MainEthPrizeAmountPercentage is a free data retrieval call binding the contract method 0xf444b298.
//
// Solidity: function mainEthPrizeAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) MainEthPrizeAmountPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.MainEthPrizeAmountPercentage(&_CosmicSignatureGameV2.CallOpts)
}

// MainEthPrizeAmountPercentage is a free data retrieval call binding the contract method 0xf444b298.
//
// Solidity: function mainEthPrizeAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) MainEthPrizeAmountPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.MainEthPrizeAmountPercentage(&_CosmicSignatureGameV2.CallOpts)
}

// MainPrizeTime is a free data retrieval call binding the contract method 0x18305de2.
//
// Solidity: function mainPrizeTime() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) MainPrizeTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "mainPrizeTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MainPrizeTime is a free data retrieval call binding the contract method 0x18305de2.
//
// Solidity: function mainPrizeTime() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) MainPrizeTime() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.MainPrizeTime(&_CosmicSignatureGameV2.CallOpts)
}

// MainPrizeTime is a free data retrieval call binding the contract method 0x18305de2.
//
// Solidity: function mainPrizeTime() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) MainPrizeTime() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.MainPrizeTime(&_CosmicSignatureGameV2.CallOpts)
}

// MainPrizeTimeIncrementInMicroSeconds is a free data retrieval call binding the contract method 0xeb13430e.
//
// Solidity: function mainPrizeTimeIncrementInMicroSeconds() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) MainPrizeTimeIncrementInMicroSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "mainPrizeTimeIncrementInMicroSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MainPrizeTimeIncrementInMicroSeconds is a free data retrieval call binding the contract method 0xeb13430e.
//
// Solidity: function mainPrizeTimeIncrementInMicroSeconds() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) MainPrizeTimeIncrementInMicroSeconds() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.MainPrizeTimeIncrementInMicroSeconds(&_CosmicSignatureGameV2.CallOpts)
}

// MainPrizeTimeIncrementInMicroSeconds is a free data retrieval call binding the contract method 0xeb13430e.
//
// Solidity: function mainPrizeTimeIncrementInMicroSeconds() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) MainPrizeTimeIncrementInMicroSeconds() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.MainPrizeTimeIncrementInMicroSeconds(&_CosmicSignatureGameV2.CallOpts)
}

// MainPrizeTimeIncrementIncreaseDivisor is a free data retrieval call binding the contract method 0x56732241.
//
// Solidity: function mainPrizeTimeIncrementIncreaseDivisor() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) MainPrizeTimeIncrementIncreaseDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "mainPrizeTimeIncrementIncreaseDivisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MainPrizeTimeIncrementIncreaseDivisor is a free data retrieval call binding the contract method 0x56732241.
//
// Solidity: function mainPrizeTimeIncrementIncreaseDivisor() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) MainPrizeTimeIncrementIncreaseDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.MainPrizeTimeIncrementIncreaseDivisor(&_CosmicSignatureGameV2.CallOpts)
}

// MainPrizeTimeIncrementIncreaseDivisor is a free data retrieval call binding the contract method 0x56732241.
//
// Solidity: function mainPrizeTimeIncrementIncreaseDivisor() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) MainPrizeTimeIncrementIncreaseDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.MainPrizeTimeIncrementIncreaseDivisor(&_CosmicSignatureGameV2.CallOpts)
}

// MarketingWallet is a free data retrieval call binding the contract method 0x75f0a874.
//
// Solidity: function marketingWallet() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) MarketingWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "marketingWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketingWallet is a free data retrieval call binding the contract method 0x75f0a874.
//
// Solidity: function marketingWallet() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) MarketingWallet() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.MarketingWallet(&_CosmicSignatureGameV2.CallOpts)
}

// MarketingWallet is a free data retrieval call binding the contract method 0x75f0a874.
//
// Solidity: function marketingWallet() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) MarketingWallet() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.MarketingWallet(&_CosmicSignatureGameV2.CallOpts)
}

// MarketingWalletCstContributionAmount is a free data retrieval call binding the contract method 0x4164b95b.
//
// Solidity: function marketingWalletCstContributionAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) MarketingWalletCstContributionAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "marketingWalletCstContributionAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketingWalletCstContributionAmount is a free data retrieval call binding the contract method 0x4164b95b.
//
// Solidity: function marketingWalletCstContributionAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) MarketingWalletCstContributionAmount() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.MarketingWalletCstContributionAmount(&_CosmicSignatureGameV2.CallOpts)
}

// MarketingWalletCstContributionAmount is a free data retrieval call binding the contract method 0x4164b95b.
//
// Solidity: function marketingWalletCstContributionAmount() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) MarketingWalletCstContributionAmount() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.MarketingWalletCstContributionAmount(&_CosmicSignatureGameV2.CallOpts)
}

// NextEthBidPrice is a free data retrieval call binding the contract method 0xefeb248a.
//
// Solidity: function nextEthBidPrice() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) NextEthBidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "nextEthBidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextEthBidPrice is a free data retrieval call binding the contract method 0xefeb248a.
//
// Solidity: function nextEthBidPrice() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) NextEthBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.NextEthBidPrice(&_CosmicSignatureGameV2.CallOpts)
}

// NextEthBidPrice is a free data retrieval call binding the contract method 0xefeb248a.
//
// Solidity: function nextEthBidPrice() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) NextEthBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.NextEthBidPrice(&_CosmicSignatureGameV2.CallOpts)
}

// NextRoundFirstCstDutchAuctionBeginningBidPrice is a free data retrieval call binding the contract method 0xc7e7a601.
//
// Solidity: function nextRoundFirstCstDutchAuctionBeginningBidPrice() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) NextRoundFirstCstDutchAuctionBeginningBidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "nextRoundFirstCstDutchAuctionBeginningBidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRoundFirstCstDutchAuctionBeginningBidPrice is a free data retrieval call binding the contract method 0xc7e7a601.
//
// Solidity: function nextRoundFirstCstDutchAuctionBeginningBidPrice() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) NextRoundFirstCstDutchAuctionBeginningBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.NextRoundFirstCstDutchAuctionBeginningBidPrice(&_CosmicSignatureGameV2.CallOpts)
}

// NextRoundFirstCstDutchAuctionBeginningBidPrice is a free data retrieval call binding the contract method 0xc7e7a601.
//
// Solidity: function nextRoundFirstCstDutchAuctionBeginningBidPrice() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) NextRoundFirstCstDutchAuctionBeginningBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.NextRoundFirstCstDutchAuctionBeginningBidPrice(&_CosmicSignatureGameV2.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) Nft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "nft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) Nft() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.Nft(&_CosmicSignatureGameV2.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) Nft() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.Nft(&_CosmicSignatureGameV2.CallOpts)
}

// NumEthDonationWithInfoRecords is a free data retrieval call binding the contract method 0x0b5f95ae.
//
// Solidity: function numEthDonationWithInfoRecords() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) NumEthDonationWithInfoRecords(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "numEthDonationWithInfoRecords")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumEthDonationWithInfoRecords is a free data retrieval call binding the contract method 0x0b5f95ae.
//
// Solidity: function numEthDonationWithInfoRecords() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) NumEthDonationWithInfoRecords() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.NumEthDonationWithInfoRecords(&_CosmicSignatureGameV2.CallOpts)
}

// NumEthDonationWithInfoRecords is a free data retrieval call binding the contract method 0x0b5f95ae.
//
// Solidity: function numEthDonationWithInfoRecords() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) NumEthDonationWithInfoRecords() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.NumEthDonationWithInfoRecords(&_CosmicSignatureGameV2.CallOpts)
}

// NumRaffleCosmicSignatureNftsForBidders is a free data retrieval call binding the contract method 0x4c2a4a33.
//
// Solidity: function numRaffleCosmicSignatureNftsForBidders() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) NumRaffleCosmicSignatureNftsForBidders(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "numRaffleCosmicSignatureNftsForBidders")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRaffleCosmicSignatureNftsForBidders is a free data retrieval call binding the contract method 0x4c2a4a33.
//
// Solidity: function numRaffleCosmicSignatureNftsForBidders() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) NumRaffleCosmicSignatureNftsForBidders() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.NumRaffleCosmicSignatureNftsForBidders(&_CosmicSignatureGameV2.CallOpts)
}

// NumRaffleCosmicSignatureNftsForBidders is a free data retrieval call binding the contract method 0x4c2a4a33.
//
// Solidity: function numRaffleCosmicSignatureNftsForBidders() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) NumRaffleCosmicSignatureNftsForBidders() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.NumRaffleCosmicSignatureNftsForBidders(&_CosmicSignatureGameV2.CallOpts)
}

// NumRaffleCosmicSignatureNftsForRandomWalkNftStakers is a free data retrieval call binding the contract method 0xe2f9185f.
//
// Solidity: function numRaffleCosmicSignatureNftsForRandomWalkNftStakers() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) NumRaffleCosmicSignatureNftsForRandomWalkNftStakers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "numRaffleCosmicSignatureNftsForRandomWalkNftStakers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRaffleCosmicSignatureNftsForRandomWalkNftStakers is a free data retrieval call binding the contract method 0xe2f9185f.
//
// Solidity: function numRaffleCosmicSignatureNftsForRandomWalkNftStakers() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) NumRaffleCosmicSignatureNftsForRandomWalkNftStakers() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.NumRaffleCosmicSignatureNftsForRandomWalkNftStakers(&_CosmicSignatureGameV2.CallOpts)
}

// NumRaffleCosmicSignatureNftsForRandomWalkNftStakers is a free data retrieval call binding the contract method 0xe2f9185f.
//
// Solidity: function numRaffleCosmicSignatureNftsForRandomWalkNftStakers() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) NumRaffleCosmicSignatureNftsForRandomWalkNftStakers() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.NumRaffleCosmicSignatureNftsForRandomWalkNftStakers(&_CosmicSignatureGameV2.CallOpts)
}

// NumRaffleEthPrizesForBidders is a free data retrieval call binding the contract method 0xf11400f0.
//
// Solidity: function numRaffleEthPrizesForBidders() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) NumRaffleEthPrizesForBidders(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "numRaffleEthPrizesForBidders")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRaffleEthPrizesForBidders is a free data retrieval call binding the contract method 0xf11400f0.
//
// Solidity: function numRaffleEthPrizesForBidders() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) NumRaffleEthPrizesForBidders() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.NumRaffleEthPrizesForBidders(&_CosmicSignatureGameV2.CallOpts)
}

// NumRaffleEthPrizesForBidders is a free data retrieval call binding the contract method 0xf11400f0.
//
// Solidity: function numRaffleEthPrizesForBidders() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) NumRaffleEthPrizesForBidders() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.NumRaffleEthPrizesForBidders(&_CosmicSignatureGameV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) Owner() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.Owner(&_CosmicSignatureGameV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) Owner() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.Owner(&_CosmicSignatureGameV2.CallOpts)
}

// PrevEnduranceChampionDuration is a free data retrieval call binding the contract method 0xeaace302.
//
// Solidity: function prevEnduranceChampionDuration() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) PrevEnduranceChampionDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "prevEnduranceChampionDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrevEnduranceChampionDuration is a free data retrieval call binding the contract method 0xeaace302.
//
// Solidity: function prevEnduranceChampionDuration() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) PrevEnduranceChampionDuration() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.PrevEnduranceChampionDuration(&_CosmicSignatureGameV2.CallOpts)
}

// PrevEnduranceChampionDuration is a free data retrieval call binding the contract method 0xeaace302.
//
// Solidity: function prevEnduranceChampionDuration() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) PrevEnduranceChampionDuration() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.PrevEnduranceChampionDuration(&_CosmicSignatureGameV2.CallOpts)
}

// PrizesWallet is a free data retrieval call binding the contract method 0x2afa2580.
//
// Solidity: function prizesWallet() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) PrizesWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "prizesWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrizesWallet is a free data retrieval call binding the contract method 0x2afa2580.
//
// Solidity: function prizesWallet() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) PrizesWallet() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.PrizesWallet(&_CosmicSignatureGameV2.CallOpts)
}

// PrizesWallet is a free data retrieval call binding the contract method 0x2afa2580.
//
// Solidity: function prizesWallet() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) PrizesWallet() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.PrizesWallet(&_CosmicSignatureGameV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) ProxiableUUID() ([32]byte, error) {
	return _CosmicSignatureGameV2.Contract.ProxiableUUID(&_CosmicSignatureGameV2.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) ProxiableUUID() ([32]byte, error) {
	return _CosmicSignatureGameV2.Contract.ProxiableUUID(&_CosmicSignatureGameV2.CallOpts)
}

// RaffleTotalEthPrizeAmountForBiddersPercentage is a free data retrieval call binding the contract method 0x477adf2a.
//
// Solidity: function raffleTotalEthPrizeAmountForBiddersPercentage() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) RaffleTotalEthPrizeAmountForBiddersPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "raffleTotalEthPrizeAmountForBiddersPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaffleTotalEthPrizeAmountForBiddersPercentage is a free data retrieval call binding the contract method 0x477adf2a.
//
// Solidity: function raffleTotalEthPrizeAmountForBiddersPercentage() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) RaffleTotalEthPrizeAmountForBiddersPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.RaffleTotalEthPrizeAmountForBiddersPercentage(&_CosmicSignatureGameV2.CallOpts)
}

// RaffleTotalEthPrizeAmountForBiddersPercentage is a free data retrieval call binding the contract method 0x477adf2a.
//
// Solidity: function raffleTotalEthPrizeAmountForBiddersPercentage() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) RaffleTotalEthPrizeAmountForBiddersPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.RaffleTotalEthPrizeAmountForBiddersPercentage(&_CosmicSignatureGameV2.CallOpts)
}

// RandomWalkNft is a free data retrieval call binding the contract method 0x755b4ef7.
//
// Solidity: function randomWalkNft() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) RandomWalkNft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "randomWalkNft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RandomWalkNft is a free data retrieval call binding the contract method 0x755b4ef7.
//
// Solidity: function randomWalkNft() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) RandomWalkNft() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.RandomWalkNft(&_CosmicSignatureGameV2.CallOpts)
}

// RandomWalkNft is a free data retrieval call binding the contract method 0x755b4ef7.
//
// Solidity: function randomWalkNft() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) RandomWalkNft() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.RandomWalkNft(&_CosmicSignatureGameV2.CallOpts)
}

// RoundActivationTime is a free data retrieval call binding the contract method 0x6e970834.
//
// Solidity: function roundActivationTime() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) RoundActivationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "roundActivationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundActivationTime is a free data retrieval call binding the contract method 0x6e970834.
//
// Solidity: function roundActivationTime() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) RoundActivationTime() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.RoundActivationTime(&_CosmicSignatureGameV2.CallOpts)
}

// RoundActivationTime is a free data retrieval call binding the contract method 0x6e970834.
//
// Solidity: function roundActivationTime() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) RoundActivationTime() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.RoundActivationTime(&_CosmicSignatureGameV2.CallOpts)
}

// RoundNum is a free data retrieval call binding the contract method 0x119b22b3.
//
// Solidity: function roundNum() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) RoundNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "roundNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundNum is a free data retrieval call binding the contract method 0x119b22b3.
//
// Solidity: function roundNum() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) RoundNum() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.RoundNum(&_CosmicSignatureGameV2.CallOpts)
}

// RoundNum is a free data retrieval call binding the contract method 0x119b22b3.
//
// Solidity: function roundNum() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) RoundNum() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.RoundNum(&_CosmicSignatureGameV2.CallOpts)
}

// StakingWalletCosmicSignatureNft is a free data retrieval call binding the contract method 0xa922ab5d.
//
// Solidity: function stakingWalletCosmicSignatureNft() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) StakingWalletCosmicSignatureNft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "stakingWalletCosmicSignatureNft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingWalletCosmicSignatureNft is a free data retrieval call binding the contract method 0xa922ab5d.
//
// Solidity: function stakingWalletCosmicSignatureNft() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) StakingWalletCosmicSignatureNft() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.StakingWalletCosmicSignatureNft(&_CosmicSignatureGameV2.CallOpts)
}

// StakingWalletCosmicSignatureNft is a free data retrieval call binding the contract method 0xa922ab5d.
//
// Solidity: function stakingWalletCosmicSignatureNft() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) StakingWalletCosmicSignatureNft() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.StakingWalletCosmicSignatureNft(&_CosmicSignatureGameV2.CallOpts)
}

// StakingWalletRandomWalkNft is a free data retrieval call binding the contract method 0x2d809e88.
//
// Solidity: function stakingWalletRandomWalkNft() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) StakingWalletRandomWalkNft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "stakingWalletRandomWalkNft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingWalletRandomWalkNft is a free data retrieval call binding the contract method 0x2d809e88.
//
// Solidity: function stakingWalletRandomWalkNft() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) StakingWalletRandomWalkNft() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.StakingWalletRandomWalkNft(&_CosmicSignatureGameV2.CallOpts)
}

// StakingWalletRandomWalkNft is a free data retrieval call binding the contract method 0x2d809e88.
//
// Solidity: function stakingWalletRandomWalkNft() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) StakingWalletRandomWalkNft() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.StakingWalletRandomWalkNft(&_CosmicSignatureGameV2.CallOpts)
}

// TimeoutDurationToClaimMainPrize is a free data retrieval call binding the contract method 0x3b9d292e.
//
// Solidity: function timeoutDurationToClaimMainPrize() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) TimeoutDurationToClaimMainPrize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "timeoutDurationToClaimMainPrize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeoutDurationToClaimMainPrize is a free data retrieval call binding the contract method 0x3b9d292e.
//
// Solidity: function timeoutDurationToClaimMainPrize() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) TimeoutDurationToClaimMainPrize() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.TimeoutDurationToClaimMainPrize(&_CosmicSignatureGameV2.CallOpts)
}

// TimeoutDurationToClaimMainPrize is a free data retrieval call binding the contract method 0x3b9d292e.
//
// Solidity: function timeoutDurationToClaimMainPrize() view returns(uint256)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) TimeoutDurationToClaimMainPrize() (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.TimeoutDurationToClaimMainPrize(&_CosmicSignatureGameV2.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) Token() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.Token(&_CosmicSignatureGameV2.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) Token() (common.Address, error) {
	return _CosmicSignatureGameV2.Contract.Token(&_CosmicSignatureGameV2.CallOpts)
}

// TryGetCurrentChampions is a free data retrieval call binding the contract method 0xcb720d4d.
//
// Solidity: function tryGetCurrentChampions() view returns(address enduranceChampionAddress_, uint256 enduranceChampionDuration_, address chronoWarriorAddress_, uint256 chronoWarriorDuration_)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) TryGetCurrentChampions(opts *bind.CallOpts) (struct {
	EnduranceChampionAddress  common.Address
	EnduranceChampionDuration *big.Int
	ChronoWarriorAddress      common.Address
	ChronoWarriorDuration     *big.Int
}, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "tryGetCurrentChampions")

	outstruct := new(struct {
		EnduranceChampionAddress  common.Address
		EnduranceChampionDuration *big.Int
		ChronoWarriorAddress      common.Address
		ChronoWarriorDuration     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EnduranceChampionAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.EnduranceChampionDuration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ChronoWarriorAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.ChronoWarriorDuration = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TryGetCurrentChampions is a free data retrieval call binding the contract method 0xcb720d4d.
//
// Solidity: function tryGetCurrentChampions() view returns(address enduranceChampionAddress_, uint256 enduranceChampionDuration_, address chronoWarriorAddress_, uint256 chronoWarriorDuration_)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) TryGetCurrentChampions() (struct {
	EnduranceChampionAddress  common.Address
	EnduranceChampionDuration *big.Int
	ChronoWarriorAddress      common.Address
	ChronoWarriorDuration     *big.Int
}, error) {
	return _CosmicSignatureGameV2.Contract.TryGetCurrentChampions(&_CosmicSignatureGameV2.CallOpts)
}

// TryGetCurrentChampions is a free data retrieval call binding the contract method 0xcb720d4d.
//
// Solidity: function tryGetCurrentChampions() view returns(address enduranceChampionAddress_, uint256 enduranceChampionDuration_, address chronoWarriorAddress_, uint256 chronoWarriorDuration_)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) TryGetCurrentChampions() (struct {
	EnduranceChampionAddress  common.Address
	EnduranceChampionDuration *big.Int
	ChronoWarriorAddress      common.Address
	ChronoWarriorDuration     *big.Int
}, error) {
	return _CosmicSignatureGameV2.Contract.TryGetCurrentChampions(&_CosmicSignatureGameV2.CallOpts)
}

// UsedRandomWalkNfts is a free data retrieval call binding the contract method 0xebb9bc5c.
//
// Solidity: function usedRandomWalkNfts(uint256 nftId) view returns(uint256 nftWasUsed)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Caller) UsedRandomWalkNfts(opts *bind.CallOpts, nftId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV2.contract.Call(opts, &out, "usedRandomWalkNfts", nftId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsedRandomWalkNfts is a free data retrieval call binding the contract method 0xebb9bc5c.
//
// Solidity: function usedRandomWalkNfts(uint256 nftId) view returns(uint256 nftWasUsed)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) UsedRandomWalkNfts(nftId *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.UsedRandomWalkNfts(&_CosmicSignatureGameV2.CallOpts, nftId)
}

// UsedRandomWalkNfts is a free data retrieval call binding the contract method 0xebb9bc5c.
//
// Solidity: function usedRandomWalkNfts(uint256 nftId) view returns(uint256 nftWasUsed)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2CallerSession) UsedRandomWalkNfts(nftId *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV2.Contract.UsedRandomWalkNfts(&_CosmicSignatureGameV2.CallOpts, nftId)
}

// BidWithCst is a paid mutator transaction binding the contract method 0xd7559b9c.
//
// Solidity: function bidWithCst(uint256 priceMaxLimit_, string message_, uint256 bidCstRewardAmountMinLimit_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) BidWithCst(opts *bind.TransactOpts, priceMaxLimit_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "bidWithCst", priceMaxLimit_, message_, bidCstRewardAmountMinLimit_)
}

// BidWithCst is a paid mutator transaction binding the contract method 0xd7559b9c.
//
// Solidity: function bidWithCst(uint256 priceMaxLimit_, string message_, uint256 bidCstRewardAmountMinLimit_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) BidWithCst(priceMaxLimit_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.BidWithCst(&_CosmicSignatureGameV2.TransactOpts, priceMaxLimit_, message_, bidCstRewardAmountMinLimit_)
}

// BidWithCst is a paid mutator transaction binding the contract method 0xd7559b9c.
//
// Solidity: function bidWithCst(uint256 priceMaxLimit_, string message_, uint256 bidCstRewardAmountMinLimit_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) BidWithCst(priceMaxLimit_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.BidWithCst(&_CosmicSignatureGameV2.TransactOpts, priceMaxLimit_, message_, bidCstRewardAmountMinLimit_)
}

// BidWithCstAndDonateNft is a paid mutator transaction binding the contract method 0x329b95a5.
//
// Solidity: function bidWithCstAndDonateNft(uint256 priceMaxLimit_, string message_, uint256 bidCstRewardAmountMinLimit_, address nftAddress_, uint256 nftId_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) BidWithCstAndDonateNft(opts *bind.TransactOpts, priceMaxLimit_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "bidWithCstAndDonateNft", priceMaxLimit_, message_, bidCstRewardAmountMinLimit_, nftAddress_, nftId_)
}

// BidWithCstAndDonateNft is a paid mutator transaction binding the contract method 0x329b95a5.
//
// Solidity: function bidWithCstAndDonateNft(uint256 priceMaxLimit_, string message_, uint256 bidCstRewardAmountMinLimit_, address nftAddress_, uint256 nftId_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) BidWithCstAndDonateNft(priceMaxLimit_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.BidWithCstAndDonateNft(&_CosmicSignatureGameV2.TransactOpts, priceMaxLimit_, message_, bidCstRewardAmountMinLimit_, nftAddress_, nftId_)
}

// BidWithCstAndDonateNft is a paid mutator transaction binding the contract method 0x329b95a5.
//
// Solidity: function bidWithCstAndDonateNft(uint256 priceMaxLimit_, string message_, uint256 bidCstRewardAmountMinLimit_, address nftAddress_, uint256 nftId_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) BidWithCstAndDonateNft(priceMaxLimit_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.BidWithCstAndDonateNft(&_CosmicSignatureGameV2.TransactOpts, priceMaxLimit_, message_, bidCstRewardAmountMinLimit_, nftAddress_, nftId_)
}

// BidWithCstAndDonateToken is a paid mutator transaction binding the contract method 0x11b0d1fe.
//
// Solidity: function bidWithCstAndDonateToken(uint256 priceMaxLimit_, string message_, uint256 bidCstRewardAmountMinLimit_, address tokenAddress_, uint256 amount_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) BidWithCstAndDonateToken(opts *bind.TransactOpts, priceMaxLimit_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "bidWithCstAndDonateToken", priceMaxLimit_, message_, bidCstRewardAmountMinLimit_, tokenAddress_, amount_)
}

// BidWithCstAndDonateToken is a paid mutator transaction binding the contract method 0x11b0d1fe.
//
// Solidity: function bidWithCstAndDonateToken(uint256 priceMaxLimit_, string message_, uint256 bidCstRewardAmountMinLimit_, address tokenAddress_, uint256 amount_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) BidWithCstAndDonateToken(priceMaxLimit_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.BidWithCstAndDonateToken(&_CosmicSignatureGameV2.TransactOpts, priceMaxLimit_, message_, bidCstRewardAmountMinLimit_, tokenAddress_, amount_)
}

// BidWithCstAndDonateToken is a paid mutator transaction binding the contract method 0x11b0d1fe.
//
// Solidity: function bidWithCstAndDonateToken(uint256 priceMaxLimit_, string message_, uint256 bidCstRewardAmountMinLimit_, address tokenAddress_, uint256 amount_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) BidWithCstAndDonateToken(priceMaxLimit_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.BidWithCstAndDonateToken(&_CosmicSignatureGameV2.TransactOpts, priceMaxLimit_, message_, bidCstRewardAmountMinLimit_, tokenAddress_, amount_)
}

// BidWithEth is a paid mutator transaction binding the contract method 0x928880fa.
//
// Solidity: function bidWithEth(int256 randomWalkNftId_, string message_, uint256 bidCstRewardAmountMinLimit_) payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) BidWithEth(opts *bind.TransactOpts, randomWalkNftId_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "bidWithEth", randomWalkNftId_, message_, bidCstRewardAmountMinLimit_)
}

// BidWithEth is a paid mutator transaction binding the contract method 0x928880fa.
//
// Solidity: function bidWithEth(int256 randomWalkNftId_, string message_, uint256 bidCstRewardAmountMinLimit_) payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) BidWithEth(randomWalkNftId_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.BidWithEth(&_CosmicSignatureGameV2.TransactOpts, randomWalkNftId_, message_, bidCstRewardAmountMinLimit_)
}

// BidWithEth is a paid mutator transaction binding the contract method 0x928880fa.
//
// Solidity: function bidWithEth(int256 randomWalkNftId_, string message_, uint256 bidCstRewardAmountMinLimit_) payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) BidWithEth(randomWalkNftId_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.BidWithEth(&_CosmicSignatureGameV2.TransactOpts, randomWalkNftId_, message_, bidCstRewardAmountMinLimit_)
}

// BidWithEthAndDonateNft is a paid mutator transaction binding the contract method 0xb78d1e2a.
//
// Solidity: function bidWithEthAndDonateNft(int256 randomWalkNftId_, string message_, uint256 bidCstRewardAmountMinLimit_, address nftAddress_, uint256 nftId_) payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) BidWithEthAndDonateNft(opts *bind.TransactOpts, randomWalkNftId_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "bidWithEthAndDonateNft", randomWalkNftId_, message_, bidCstRewardAmountMinLimit_, nftAddress_, nftId_)
}

// BidWithEthAndDonateNft is a paid mutator transaction binding the contract method 0xb78d1e2a.
//
// Solidity: function bidWithEthAndDonateNft(int256 randomWalkNftId_, string message_, uint256 bidCstRewardAmountMinLimit_, address nftAddress_, uint256 nftId_) payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) BidWithEthAndDonateNft(randomWalkNftId_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.BidWithEthAndDonateNft(&_CosmicSignatureGameV2.TransactOpts, randomWalkNftId_, message_, bidCstRewardAmountMinLimit_, nftAddress_, nftId_)
}

// BidWithEthAndDonateNft is a paid mutator transaction binding the contract method 0xb78d1e2a.
//
// Solidity: function bidWithEthAndDonateNft(int256 randomWalkNftId_, string message_, uint256 bidCstRewardAmountMinLimit_, address nftAddress_, uint256 nftId_) payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) BidWithEthAndDonateNft(randomWalkNftId_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.BidWithEthAndDonateNft(&_CosmicSignatureGameV2.TransactOpts, randomWalkNftId_, message_, bidCstRewardAmountMinLimit_, nftAddress_, nftId_)
}

// BidWithEthAndDonateToken is a paid mutator transaction binding the contract method 0x876d5c36.
//
// Solidity: function bidWithEthAndDonateToken(int256 randomWalkNftId_, string message_, uint256 bidCstRewardAmountMinLimit_, address tokenAddress_, uint256 amount_) payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) BidWithEthAndDonateToken(opts *bind.TransactOpts, randomWalkNftId_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "bidWithEthAndDonateToken", randomWalkNftId_, message_, bidCstRewardAmountMinLimit_, tokenAddress_, amount_)
}

// BidWithEthAndDonateToken is a paid mutator transaction binding the contract method 0x876d5c36.
//
// Solidity: function bidWithEthAndDonateToken(int256 randomWalkNftId_, string message_, uint256 bidCstRewardAmountMinLimit_, address tokenAddress_, uint256 amount_) payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) BidWithEthAndDonateToken(randomWalkNftId_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.BidWithEthAndDonateToken(&_CosmicSignatureGameV2.TransactOpts, randomWalkNftId_, message_, bidCstRewardAmountMinLimit_, tokenAddress_, amount_)
}

// BidWithEthAndDonateToken is a paid mutator transaction binding the contract method 0x876d5c36.
//
// Solidity: function bidWithEthAndDonateToken(int256 randomWalkNftId_, string message_, uint256 bidCstRewardAmountMinLimit_, address tokenAddress_, uint256 amount_) payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) BidWithEthAndDonateToken(randomWalkNftId_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.BidWithEthAndDonateToken(&_CosmicSignatureGameV2.TransactOpts, randomWalkNftId_, message_, bidCstRewardAmountMinLimit_, tokenAddress_, amount_)
}

// ClaimMainPrize is a paid mutator transaction binding the contract method 0x448c6eb1.
//
// Solidity: function claimMainPrize() returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) ClaimMainPrize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "claimMainPrize")
}

// ClaimMainPrize is a paid mutator transaction binding the contract method 0x448c6eb1.
//
// Solidity: function claimMainPrize() returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) ClaimMainPrize() (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.ClaimMainPrize(&_CosmicSignatureGameV2.TransactOpts)
}

// ClaimMainPrize is a paid mutator transaction binding the contract method 0x448c6eb1.
//
// Solidity: function claimMainPrize() returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) ClaimMainPrize() (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.ClaimMainPrize(&_CosmicSignatureGameV2.TransactOpts)
}

// DonateEth is a paid mutator transaction binding the contract method 0xaadd1b03.
//
// Solidity: function donateEth() payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) DonateEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "donateEth")
}

// DonateEth is a paid mutator transaction binding the contract method 0xaadd1b03.
//
// Solidity: function donateEth() payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) DonateEth() (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.DonateEth(&_CosmicSignatureGameV2.TransactOpts)
}

// DonateEth is a paid mutator transaction binding the contract method 0xaadd1b03.
//
// Solidity: function donateEth() payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) DonateEth() (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.DonateEth(&_CosmicSignatureGameV2.TransactOpts)
}

// DonateEthWithInfo is a paid mutator transaction binding the contract method 0x23b31cfc.
//
// Solidity: function donateEthWithInfo(string data_) payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) DonateEthWithInfo(opts *bind.TransactOpts, data_ string) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "donateEthWithInfo", data_)
}

// DonateEthWithInfo is a paid mutator transaction binding the contract method 0x23b31cfc.
//
// Solidity: function donateEthWithInfo(string data_) payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) DonateEthWithInfo(data_ string) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.DonateEthWithInfo(&_CosmicSignatureGameV2.TransactOpts, data_)
}

// DonateEthWithInfo is a paid mutator transaction binding the contract method 0x23b31cfc.
//
// Solidity: function donateEthWithInfo(string data_) payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) DonateEthWithInfo(data_ string) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.DonateEthWithInfo(&_CosmicSignatureGameV2.TransactOpts, data_)
}

// HalveEthDutchAuctionEndingBidPrice is a paid mutator transaction binding the contract method 0xdfcd00d1.
//
// Solidity: function halveEthDutchAuctionEndingBidPrice() returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) HalveEthDutchAuctionEndingBidPrice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "halveEthDutchAuctionEndingBidPrice")
}

// HalveEthDutchAuctionEndingBidPrice is a paid mutator transaction binding the contract method 0xdfcd00d1.
//
// Solidity: function halveEthDutchAuctionEndingBidPrice() returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) HalveEthDutchAuctionEndingBidPrice() (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.HalveEthDutchAuctionEndingBidPrice(&_CosmicSignatureGameV2.TransactOpts)
}

// HalveEthDutchAuctionEndingBidPrice is a paid mutator transaction binding the contract method 0xdfcd00d1.
//
// Solidity: function halveEthDutchAuctionEndingBidPrice() returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) HalveEthDutchAuctionEndingBidPrice() (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.HalveEthDutchAuctionEndingBidPrice(&_CosmicSignatureGameV2.TransactOpts)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) InitializeV2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "initializeV2")
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) InitializeV2() (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.InitializeV2(&_CosmicSignatureGameV2.TransactOpts)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) InitializeV2() (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.InitializeV2(&_CosmicSignatureGameV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.RenounceOwnership(&_CosmicSignatureGameV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.RenounceOwnership(&_CosmicSignatureGameV2.TransactOpts)
}

// SetBidCstRewardAmountMultiplier is a paid mutator transaction binding the contract method 0x09632366.
//
// Solidity: function setBidCstRewardAmountMultiplier(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetBidCstRewardAmountMultiplier(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setBidCstRewardAmountMultiplier", newValue_)
}

// SetBidCstRewardAmountMultiplier is a paid mutator transaction binding the contract method 0x09632366.
//
// Solidity: function setBidCstRewardAmountMultiplier(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetBidCstRewardAmountMultiplier(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetBidCstRewardAmountMultiplier(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetBidCstRewardAmountMultiplier is a paid mutator transaction binding the contract method 0x09632366.
//
// Solidity: function setBidCstRewardAmountMultiplier(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetBidCstRewardAmountMultiplier(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetBidCstRewardAmountMultiplier(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetBidMessageLengthMaxLimit is a paid mutator transaction binding the contract method 0x543f416f.
//
// Solidity: function setBidMessageLengthMaxLimit(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetBidMessageLengthMaxLimit(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setBidMessageLengthMaxLimit", newValue_)
}

// SetBidMessageLengthMaxLimit is a paid mutator transaction binding the contract method 0x543f416f.
//
// Solidity: function setBidMessageLengthMaxLimit(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetBidMessageLengthMaxLimit(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetBidMessageLengthMaxLimit(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetBidMessageLengthMaxLimit is a paid mutator transaction binding the contract method 0x543f416f.
//
// Solidity: function setBidMessageLengthMaxLimit(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetBidMessageLengthMaxLimit(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetBidMessageLengthMaxLimit(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetCharityAddress is a paid mutator transaction binding the contract method 0x0c9be46d.
//
// Solidity: function setCharityAddress(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetCharityAddress(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setCharityAddress", newValue_)
}

// SetCharityAddress is a paid mutator transaction binding the contract method 0x0c9be46d.
//
// Solidity: function setCharityAddress(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetCharityAddress(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetCharityAddress(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetCharityAddress is a paid mutator transaction binding the contract method 0x0c9be46d.
//
// Solidity: function setCharityAddress(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetCharityAddress(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetCharityAddress(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetCharityEthDonationAmountPercentage is a paid mutator transaction binding the contract method 0x2d829a2d.
//
// Solidity: function setCharityEthDonationAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetCharityEthDonationAmountPercentage(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setCharityEthDonationAmountPercentage", newValue_)
}

// SetCharityEthDonationAmountPercentage is a paid mutator transaction binding the contract method 0x2d829a2d.
//
// Solidity: function setCharityEthDonationAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetCharityEthDonationAmountPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetCharityEthDonationAmountPercentage(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetCharityEthDonationAmountPercentage is a paid mutator transaction binding the contract method 0x2d829a2d.
//
// Solidity: function setCharityEthDonationAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetCharityEthDonationAmountPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetCharityEthDonationAmountPercentage(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetChronoWarriorEthPrizeAmountPercentage is a paid mutator transaction binding the contract method 0x1f1b4aa4.
//
// Solidity: function setChronoWarriorEthPrizeAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetChronoWarriorEthPrizeAmountPercentage(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setChronoWarriorEthPrizeAmountPercentage", newValue_)
}

// SetChronoWarriorEthPrizeAmountPercentage is a paid mutator transaction binding the contract method 0x1f1b4aa4.
//
// Solidity: function setChronoWarriorEthPrizeAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetChronoWarriorEthPrizeAmountPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetChronoWarriorEthPrizeAmountPercentage(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetChronoWarriorEthPrizeAmountPercentage is a paid mutator transaction binding the contract method 0x1f1b4aa4.
//
// Solidity: function setChronoWarriorEthPrizeAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetChronoWarriorEthPrizeAmountPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetChronoWarriorEthPrizeAmountPercentage(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetCosmicSignatureNft is a paid mutator transaction binding the contract method 0x44acc12a.
//
// Solidity: function setCosmicSignatureNft(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetCosmicSignatureNft(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setCosmicSignatureNft", newValue_)
}

// SetCosmicSignatureNft is a paid mutator transaction binding the contract method 0x44acc12a.
//
// Solidity: function setCosmicSignatureNft(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetCosmicSignatureNft(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetCosmicSignatureNft(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetCosmicSignatureNft is a paid mutator transaction binding the contract method 0x44acc12a.
//
// Solidity: function setCosmicSignatureNft(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetCosmicSignatureNft(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetCosmicSignatureNft(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetCosmicSignatureNftStakingTotalEthRewardAmountPercentage is a paid mutator transaction binding the contract method 0x4feb78b7.
//
// Solidity: function setCosmicSignatureNftStakingTotalEthRewardAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetCosmicSignatureNftStakingTotalEthRewardAmountPercentage(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setCosmicSignatureNftStakingTotalEthRewardAmountPercentage", newValue_)
}

// SetCosmicSignatureNftStakingTotalEthRewardAmountPercentage is a paid mutator transaction binding the contract method 0x4feb78b7.
//
// Solidity: function setCosmicSignatureNftStakingTotalEthRewardAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetCosmicSignatureNftStakingTotalEthRewardAmountPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetCosmicSignatureNftStakingTotalEthRewardAmountPercentage(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetCosmicSignatureNftStakingTotalEthRewardAmountPercentage is a paid mutator transaction binding the contract method 0x4feb78b7.
//
// Solidity: function setCosmicSignatureNftStakingTotalEthRewardAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetCosmicSignatureNftStakingTotalEthRewardAmountPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetCosmicSignatureNftStakingTotalEthRewardAmountPercentage(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetCosmicSignatureToken is a paid mutator transaction binding the contract method 0x9646d758.
//
// Solidity: function setCosmicSignatureToken(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetCosmicSignatureToken(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setCosmicSignatureToken", newValue_)
}

// SetCosmicSignatureToken is a paid mutator transaction binding the contract method 0x9646d758.
//
// Solidity: function setCosmicSignatureToken(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetCosmicSignatureToken(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetCosmicSignatureToken(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetCosmicSignatureToken is a paid mutator transaction binding the contract method 0x9646d758.
//
// Solidity: function setCosmicSignatureToken(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetCosmicSignatureToken(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetCosmicSignatureToken(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetCstDutchAuctionBeginningBidPriceMinLimit is a paid mutator transaction binding the contract method 0x88ce802c.
//
// Solidity: function setCstDutchAuctionBeginningBidPriceMinLimit(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetCstDutchAuctionBeginningBidPriceMinLimit(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setCstDutchAuctionBeginningBidPriceMinLimit", newValue_)
}

// SetCstDutchAuctionBeginningBidPriceMinLimit is a paid mutator transaction binding the contract method 0x88ce802c.
//
// Solidity: function setCstDutchAuctionBeginningBidPriceMinLimit(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetCstDutchAuctionBeginningBidPriceMinLimit(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetCstDutchAuctionBeginningBidPriceMinLimit(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetCstDutchAuctionBeginningBidPriceMinLimit is a paid mutator transaction binding the contract method 0x88ce802c.
//
// Solidity: function setCstDutchAuctionBeginningBidPriceMinLimit(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetCstDutchAuctionBeginningBidPriceMinLimit(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetCstDutchAuctionBeginningBidPriceMinLimit(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetCstDutchAuctionDuration is a paid mutator transaction binding the contract method 0xde704b41.
//
// Solidity: function setCstDutchAuctionDuration(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetCstDutchAuctionDuration(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setCstDutchAuctionDuration", newValue_)
}

// SetCstDutchAuctionDuration is a paid mutator transaction binding the contract method 0xde704b41.
//
// Solidity: function setCstDutchAuctionDuration(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetCstDutchAuctionDuration(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetCstDutchAuctionDuration(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetCstDutchAuctionDuration is a paid mutator transaction binding the contract method 0xde704b41.
//
// Solidity: function setCstDutchAuctionDuration(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetCstDutchAuctionDuration(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetCstDutchAuctionDuration(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetCstDutchAuctionDurationChangeDivisor is a paid mutator transaction binding the contract method 0x04338479.
//
// Solidity: function setCstDutchAuctionDurationChangeDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetCstDutchAuctionDurationChangeDivisor(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setCstDutchAuctionDurationChangeDivisor", newValue_)
}

// SetCstDutchAuctionDurationChangeDivisor is a paid mutator transaction binding the contract method 0x04338479.
//
// Solidity: function setCstDutchAuctionDurationChangeDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetCstDutchAuctionDurationChangeDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetCstDutchAuctionDurationChangeDivisor(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetCstDutchAuctionDurationChangeDivisor is a paid mutator transaction binding the contract method 0x04338479.
//
// Solidity: function setCstDutchAuctionDurationChangeDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetCstDutchAuctionDurationChangeDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetCstDutchAuctionDurationChangeDivisor(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetCstPrizeAmount is a paid mutator transaction binding the contract method 0xecb5776e.
//
// Solidity: function setCstPrizeAmount(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetCstPrizeAmount(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setCstPrizeAmount", newValue_)
}

// SetCstPrizeAmount is a paid mutator transaction binding the contract method 0xecb5776e.
//
// Solidity: function setCstPrizeAmount(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetCstPrizeAmount(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetCstPrizeAmount(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetCstPrizeAmount is a paid mutator transaction binding the contract method 0xecb5776e.
//
// Solidity: function setCstPrizeAmount(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetCstPrizeAmount(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetCstPrizeAmount(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetDelayDurationBeforeRoundActivation is a paid mutator transaction binding the contract method 0x09794bee.
//
// Solidity: function setDelayDurationBeforeRoundActivation(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetDelayDurationBeforeRoundActivation(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setDelayDurationBeforeRoundActivation", newValue_)
}

// SetDelayDurationBeforeRoundActivation is a paid mutator transaction binding the contract method 0x09794bee.
//
// Solidity: function setDelayDurationBeforeRoundActivation(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetDelayDurationBeforeRoundActivation(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetDelayDurationBeforeRoundActivation(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetDelayDurationBeforeRoundActivation is a paid mutator transaction binding the contract method 0x09794bee.
//
// Solidity: function setDelayDurationBeforeRoundActivation(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetDelayDurationBeforeRoundActivation(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetDelayDurationBeforeRoundActivation(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetEthBidPriceIncreaseDivisor is a paid mutator transaction binding the contract method 0xf49efe9d.
//
// Solidity: function setEthBidPriceIncreaseDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetEthBidPriceIncreaseDivisor(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setEthBidPriceIncreaseDivisor", newValue_)
}

// SetEthBidPriceIncreaseDivisor is a paid mutator transaction binding the contract method 0xf49efe9d.
//
// Solidity: function setEthBidPriceIncreaseDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetEthBidPriceIncreaseDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetEthBidPriceIncreaseDivisor(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetEthBidPriceIncreaseDivisor is a paid mutator transaction binding the contract method 0xf49efe9d.
//
// Solidity: function setEthBidPriceIncreaseDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetEthBidPriceIncreaseDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetEthBidPriceIncreaseDivisor(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetEthBidRefundAmountInGasToSwallowMaxLimit is a paid mutator transaction binding the contract method 0x441b3289.
//
// Solidity: function setEthBidRefundAmountInGasToSwallowMaxLimit(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetEthBidRefundAmountInGasToSwallowMaxLimit(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setEthBidRefundAmountInGasToSwallowMaxLimit", newValue_)
}

// SetEthBidRefundAmountInGasToSwallowMaxLimit is a paid mutator transaction binding the contract method 0x441b3289.
//
// Solidity: function setEthBidRefundAmountInGasToSwallowMaxLimit(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetEthBidRefundAmountInGasToSwallowMaxLimit(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetEthBidRefundAmountInGasToSwallowMaxLimit(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetEthBidRefundAmountInGasToSwallowMaxLimit is a paid mutator transaction binding the contract method 0x441b3289.
//
// Solidity: function setEthBidRefundAmountInGasToSwallowMaxLimit(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetEthBidRefundAmountInGasToSwallowMaxLimit(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetEthBidRefundAmountInGasToSwallowMaxLimit(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetEthDutchAuctionDurationDivisor is a paid mutator transaction binding the contract method 0xf34d411c.
//
// Solidity: function setEthDutchAuctionDurationDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetEthDutchAuctionDurationDivisor(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setEthDutchAuctionDurationDivisor", newValue_)
}

// SetEthDutchAuctionDurationDivisor is a paid mutator transaction binding the contract method 0xf34d411c.
//
// Solidity: function setEthDutchAuctionDurationDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetEthDutchAuctionDurationDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetEthDutchAuctionDurationDivisor(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetEthDutchAuctionDurationDivisor is a paid mutator transaction binding the contract method 0xf34d411c.
//
// Solidity: function setEthDutchAuctionDurationDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetEthDutchAuctionDurationDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetEthDutchAuctionDurationDivisor(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetEthDutchAuctionEndingBidPriceDivisor is a paid mutator transaction binding the contract method 0xddd6df07.
//
// Solidity: function setEthDutchAuctionEndingBidPriceDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetEthDutchAuctionEndingBidPriceDivisor(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setEthDutchAuctionEndingBidPriceDivisor", newValue_)
}

// SetEthDutchAuctionEndingBidPriceDivisor is a paid mutator transaction binding the contract method 0xddd6df07.
//
// Solidity: function setEthDutchAuctionEndingBidPriceDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetEthDutchAuctionEndingBidPriceDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetEthDutchAuctionEndingBidPriceDivisor(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetEthDutchAuctionEndingBidPriceDivisor is a paid mutator transaction binding the contract method 0xddd6df07.
//
// Solidity: function setEthDutchAuctionEndingBidPriceDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetEthDutchAuctionEndingBidPriceDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetEthDutchAuctionEndingBidPriceDivisor(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetInitialDurationUntilMainPrizeDivisor is a paid mutator transaction binding the contract method 0x77fa1027.
//
// Solidity: function setInitialDurationUntilMainPrizeDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetInitialDurationUntilMainPrizeDivisor(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setInitialDurationUntilMainPrizeDivisor", newValue_)
}

// SetInitialDurationUntilMainPrizeDivisor is a paid mutator transaction binding the contract method 0x77fa1027.
//
// Solidity: function setInitialDurationUntilMainPrizeDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetInitialDurationUntilMainPrizeDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetInitialDurationUntilMainPrizeDivisor(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetInitialDurationUntilMainPrizeDivisor is a paid mutator transaction binding the contract method 0x77fa1027.
//
// Solidity: function setInitialDurationUntilMainPrizeDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetInitialDurationUntilMainPrizeDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetInitialDurationUntilMainPrizeDivisor(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetMainEthPrizeAmountPercentage is a paid mutator transaction binding the contract method 0x6b59acb8.
//
// Solidity: function setMainEthPrizeAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetMainEthPrizeAmountPercentage(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setMainEthPrizeAmountPercentage", newValue_)
}

// SetMainEthPrizeAmountPercentage is a paid mutator transaction binding the contract method 0x6b59acb8.
//
// Solidity: function setMainEthPrizeAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetMainEthPrizeAmountPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetMainEthPrizeAmountPercentage(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetMainEthPrizeAmountPercentage is a paid mutator transaction binding the contract method 0x6b59acb8.
//
// Solidity: function setMainEthPrizeAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetMainEthPrizeAmountPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetMainEthPrizeAmountPercentage(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetMainPrizeTimeIncrementInMicroSeconds is a paid mutator transaction binding the contract method 0xa4be0d40.
//
// Solidity: function setMainPrizeTimeIncrementInMicroSeconds(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetMainPrizeTimeIncrementInMicroSeconds(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setMainPrizeTimeIncrementInMicroSeconds", newValue_)
}

// SetMainPrizeTimeIncrementInMicroSeconds is a paid mutator transaction binding the contract method 0xa4be0d40.
//
// Solidity: function setMainPrizeTimeIncrementInMicroSeconds(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetMainPrizeTimeIncrementInMicroSeconds(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetMainPrizeTimeIncrementInMicroSeconds(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetMainPrizeTimeIncrementInMicroSeconds is a paid mutator transaction binding the contract method 0xa4be0d40.
//
// Solidity: function setMainPrizeTimeIncrementInMicroSeconds(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetMainPrizeTimeIncrementInMicroSeconds(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetMainPrizeTimeIncrementInMicroSeconds(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetMainPrizeTimeIncrementIncreaseDivisor is a paid mutator transaction binding the contract method 0xcfb4e599.
//
// Solidity: function setMainPrizeTimeIncrementIncreaseDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetMainPrizeTimeIncrementIncreaseDivisor(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setMainPrizeTimeIncrementIncreaseDivisor", newValue_)
}

// SetMainPrizeTimeIncrementIncreaseDivisor is a paid mutator transaction binding the contract method 0xcfb4e599.
//
// Solidity: function setMainPrizeTimeIncrementIncreaseDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetMainPrizeTimeIncrementIncreaseDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetMainPrizeTimeIncrementIncreaseDivisor(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetMainPrizeTimeIncrementIncreaseDivisor is a paid mutator transaction binding the contract method 0xcfb4e599.
//
// Solidity: function setMainPrizeTimeIncrementIncreaseDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetMainPrizeTimeIncrementIncreaseDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetMainPrizeTimeIncrementIncreaseDivisor(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetMarketingWallet is a paid mutator transaction binding the contract method 0x5d098b38.
//
// Solidity: function setMarketingWallet(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetMarketingWallet(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setMarketingWallet", newValue_)
}

// SetMarketingWallet is a paid mutator transaction binding the contract method 0x5d098b38.
//
// Solidity: function setMarketingWallet(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetMarketingWallet(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetMarketingWallet(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetMarketingWallet is a paid mutator transaction binding the contract method 0x5d098b38.
//
// Solidity: function setMarketingWallet(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetMarketingWallet(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetMarketingWallet(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetMarketingWalletCstContributionAmount is a paid mutator transaction binding the contract method 0xd9ab9eaa.
//
// Solidity: function setMarketingWalletCstContributionAmount(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetMarketingWalletCstContributionAmount(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setMarketingWalletCstContributionAmount", newValue_)
}

// SetMarketingWalletCstContributionAmount is a paid mutator transaction binding the contract method 0xd9ab9eaa.
//
// Solidity: function setMarketingWalletCstContributionAmount(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetMarketingWalletCstContributionAmount(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetMarketingWalletCstContributionAmount(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetMarketingWalletCstContributionAmount is a paid mutator transaction binding the contract method 0xd9ab9eaa.
//
// Solidity: function setMarketingWalletCstContributionAmount(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetMarketingWalletCstContributionAmount(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetMarketingWalletCstContributionAmount(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetNumRaffleCosmicSignatureNftsForBidders is a paid mutator transaction binding the contract method 0x1e9cbb7e.
//
// Solidity: function setNumRaffleCosmicSignatureNftsForBidders(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetNumRaffleCosmicSignatureNftsForBidders(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setNumRaffleCosmicSignatureNftsForBidders", newValue_)
}

// SetNumRaffleCosmicSignatureNftsForBidders is a paid mutator transaction binding the contract method 0x1e9cbb7e.
//
// Solidity: function setNumRaffleCosmicSignatureNftsForBidders(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetNumRaffleCosmicSignatureNftsForBidders(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetNumRaffleCosmicSignatureNftsForBidders(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetNumRaffleCosmicSignatureNftsForBidders is a paid mutator transaction binding the contract method 0x1e9cbb7e.
//
// Solidity: function setNumRaffleCosmicSignatureNftsForBidders(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetNumRaffleCosmicSignatureNftsForBidders(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetNumRaffleCosmicSignatureNftsForBidders(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetNumRaffleCosmicSignatureNftsForRandomWalkNftStakers is a paid mutator transaction binding the contract method 0x135f3d28.
//
// Solidity: function setNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setNumRaffleCosmicSignatureNftsForRandomWalkNftStakers", newValue_)
}

// SetNumRaffleCosmicSignatureNftsForRandomWalkNftStakers is a paid mutator transaction binding the contract method 0x135f3d28.
//
// Solidity: function setNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetNumRaffleCosmicSignatureNftsForRandomWalkNftStakers is a paid mutator transaction binding the contract method 0x135f3d28.
//
// Solidity: function setNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetNumRaffleEthPrizesForBidders is a paid mutator transaction binding the contract method 0x2f894cd7.
//
// Solidity: function setNumRaffleEthPrizesForBidders(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetNumRaffleEthPrizesForBidders(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setNumRaffleEthPrizesForBidders", newValue_)
}

// SetNumRaffleEthPrizesForBidders is a paid mutator transaction binding the contract method 0x2f894cd7.
//
// Solidity: function setNumRaffleEthPrizesForBidders(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetNumRaffleEthPrizesForBidders(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetNumRaffleEthPrizesForBidders(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetNumRaffleEthPrizesForBidders is a paid mutator transaction binding the contract method 0x2f894cd7.
//
// Solidity: function setNumRaffleEthPrizesForBidders(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetNumRaffleEthPrizesForBidders(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetNumRaffleEthPrizesForBidders(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetPrizesWallet is a paid mutator transaction binding the contract method 0x6c17e3cc.
//
// Solidity: function setPrizesWallet(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetPrizesWallet(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setPrizesWallet", newValue_)
}

// SetPrizesWallet is a paid mutator transaction binding the contract method 0x6c17e3cc.
//
// Solidity: function setPrizesWallet(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetPrizesWallet(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetPrizesWallet(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetPrizesWallet is a paid mutator transaction binding the contract method 0x6c17e3cc.
//
// Solidity: function setPrizesWallet(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetPrizesWallet(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetPrizesWallet(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetRaffleTotalEthPrizeAmountForBiddersPercentage is a paid mutator transaction binding the contract method 0xfdfb9ba4.
//
// Solidity: function setRaffleTotalEthPrizeAmountForBiddersPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetRaffleTotalEthPrizeAmountForBiddersPercentage(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setRaffleTotalEthPrizeAmountForBiddersPercentage", newValue_)
}

// SetRaffleTotalEthPrizeAmountForBiddersPercentage is a paid mutator transaction binding the contract method 0xfdfb9ba4.
//
// Solidity: function setRaffleTotalEthPrizeAmountForBiddersPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetRaffleTotalEthPrizeAmountForBiddersPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetRaffleTotalEthPrizeAmountForBiddersPercentage(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetRaffleTotalEthPrizeAmountForBiddersPercentage is a paid mutator transaction binding the contract method 0xfdfb9ba4.
//
// Solidity: function setRaffleTotalEthPrizeAmountForBiddersPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetRaffleTotalEthPrizeAmountForBiddersPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetRaffleTotalEthPrizeAmountForBiddersPercentage(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetRandomWalkNft is a paid mutator transaction binding the contract method 0x9edeaf8e.
//
// Solidity: function setRandomWalkNft(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetRandomWalkNft(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setRandomWalkNft", newValue_)
}

// SetRandomWalkNft is a paid mutator transaction binding the contract method 0x9edeaf8e.
//
// Solidity: function setRandomWalkNft(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetRandomWalkNft(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetRandomWalkNft(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetRandomWalkNft is a paid mutator transaction binding the contract method 0x9edeaf8e.
//
// Solidity: function setRandomWalkNft(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetRandomWalkNft(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetRandomWalkNft(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetRoundActivationTime is a paid mutator transaction binding the contract method 0x2b91c7bb.
//
// Solidity: function setRoundActivationTime(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetRoundActivationTime(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setRoundActivationTime", newValue_)
}

// SetRoundActivationTime is a paid mutator transaction binding the contract method 0x2b91c7bb.
//
// Solidity: function setRoundActivationTime(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetRoundActivationTime(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetRoundActivationTime(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetRoundActivationTime is a paid mutator transaction binding the contract method 0x2b91c7bb.
//
// Solidity: function setRoundActivationTime(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetRoundActivationTime(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetRoundActivationTime(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetStakingWalletCosmicSignatureNft is a paid mutator transaction binding the contract method 0x5a1e5bde.
//
// Solidity: function setStakingWalletCosmicSignatureNft(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetStakingWalletCosmicSignatureNft(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setStakingWalletCosmicSignatureNft", newValue_)
}

// SetStakingWalletCosmicSignatureNft is a paid mutator transaction binding the contract method 0x5a1e5bde.
//
// Solidity: function setStakingWalletCosmicSignatureNft(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetStakingWalletCosmicSignatureNft(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetStakingWalletCosmicSignatureNft(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetStakingWalletCosmicSignatureNft is a paid mutator transaction binding the contract method 0x5a1e5bde.
//
// Solidity: function setStakingWalletCosmicSignatureNft(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetStakingWalletCosmicSignatureNft(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetStakingWalletCosmicSignatureNft(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetStakingWalletRandomWalkNft is a paid mutator transaction binding the contract method 0xb4f1b134.
//
// Solidity: function setStakingWalletRandomWalkNft(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetStakingWalletRandomWalkNft(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setStakingWalletRandomWalkNft", newValue_)
}

// SetStakingWalletRandomWalkNft is a paid mutator transaction binding the contract method 0xb4f1b134.
//
// Solidity: function setStakingWalletRandomWalkNft(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetStakingWalletRandomWalkNft(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetStakingWalletRandomWalkNft(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetStakingWalletRandomWalkNft is a paid mutator transaction binding the contract method 0xb4f1b134.
//
// Solidity: function setStakingWalletRandomWalkNft(address newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetStakingWalletRandomWalkNft(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetStakingWalletRandomWalkNft(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetTimeoutDurationToClaimMainPrize is a paid mutator transaction binding the contract method 0xf0bdab7c.
//
// Solidity: function setTimeoutDurationToClaimMainPrize(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) SetTimeoutDurationToClaimMainPrize(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "setTimeoutDurationToClaimMainPrize", newValue_)
}

// SetTimeoutDurationToClaimMainPrize is a paid mutator transaction binding the contract method 0xf0bdab7c.
//
// Solidity: function setTimeoutDurationToClaimMainPrize(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) SetTimeoutDurationToClaimMainPrize(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetTimeoutDurationToClaimMainPrize(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// SetTimeoutDurationToClaimMainPrize is a paid mutator transaction binding the contract method 0xf0bdab7c.
//
// Solidity: function setTimeoutDurationToClaimMainPrize(uint256 newValue_) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) SetTimeoutDurationToClaimMainPrize(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.SetTimeoutDurationToClaimMainPrize(&_CosmicSignatureGameV2.TransactOpts, newValue_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.TransferOwnership(&_CosmicSignatureGameV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.TransferOwnership(&_CosmicSignatureGameV2.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.UpgradeToAndCall(&_CosmicSignatureGameV2.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.UpgradeToAndCall(&_CosmicSignatureGameV2.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicSignatureGameV2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Session) Receive() (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.Receive(&_CosmicSignatureGameV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2TransactorSession) Receive() (*types.Transaction, error) {
	return _CosmicSignatureGameV2.Contract.Receive(&_CosmicSignatureGameV2.TransactOpts)
}

// CosmicSignatureGameV2ArbitrumErrorIterator is returned from FilterArbitrumError and is used to iterate over the raw logs and unpacked data for ArbitrumError events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2ArbitrumErrorIterator struct {
	Event *CosmicSignatureGameV2ArbitrumError // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2ArbitrumErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2ArbitrumError)
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
		it.Event = new(CosmicSignatureGameV2ArbitrumError)
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
func (it *CosmicSignatureGameV2ArbitrumErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2ArbitrumErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2ArbitrumError represents a ArbitrumError event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2ArbitrumError struct {
	ErrStr string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterArbitrumError is a free log retrieval operation binding the contract event 0xa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d.
//
// Solidity: event ArbitrumError(string errStr)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterArbitrumError(opts *bind.FilterOpts) (*CosmicSignatureGameV2ArbitrumErrorIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "ArbitrumError")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2ArbitrumErrorIterator{contract: _CosmicSignatureGameV2.contract, event: "ArbitrumError", logs: logs, sub: sub}, nil
}

// WatchArbitrumError is a free log subscription operation binding the contract event 0xa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d.
//
// Solidity: event ArbitrumError(string errStr)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchArbitrumError(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2ArbitrumError) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "ArbitrumError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2ArbitrumError)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "ArbitrumError", log); err != nil {
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

// ParseArbitrumError is a log parse operation binding the contract event 0xa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d.
//
// Solidity: event ArbitrumError(string errStr)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseArbitrumError(log types.Log) (*CosmicSignatureGameV2ArbitrumError, error) {
	event := new(CosmicSignatureGameV2ArbitrumError)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "ArbitrumError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2BidCstRewardAmountMultiplierChangedIterator is returned from FilterBidCstRewardAmountMultiplierChanged and is used to iterate over the raw logs and unpacked data for BidCstRewardAmountMultiplierChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2BidCstRewardAmountMultiplierChangedIterator struct {
	Event *CosmicSignatureGameV2BidCstRewardAmountMultiplierChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2BidCstRewardAmountMultiplierChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2BidCstRewardAmountMultiplierChanged)
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
		it.Event = new(CosmicSignatureGameV2BidCstRewardAmountMultiplierChanged)
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
func (it *CosmicSignatureGameV2BidCstRewardAmountMultiplierChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2BidCstRewardAmountMultiplierChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2BidCstRewardAmountMultiplierChanged represents a BidCstRewardAmountMultiplierChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2BidCstRewardAmountMultiplierChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBidCstRewardAmountMultiplierChanged is a free log retrieval operation binding the contract event 0x40b9c59af8c486ccf8c7cc73df5a51e7cc29747ea7d39f99632ecaf9caa2ed1f.
//
// Solidity: event BidCstRewardAmountMultiplierChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterBidCstRewardAmountMultiplierChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2BidCstRewardAmountMultiplierChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "BidCstRewardAmountMultiplierChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2BidCstRewardAmountMultiplierChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "BidCstRewardAmountMultiplierChanged", logs: logs, sub: sub}, nil
}

// WatchBidCstRewardAmountMultiplierChanged is a free log subscription operation binding the contract event 0x40b9c59af8c486ccf8c7cc73df5a51e7cc29747ea7d39f99632ecaf9caa2ed1f.
//
// Solidity: event BidCstRewardAmountMultiplierChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchBidCstRewardAmountMultiplierChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2BidCstRewardAmountMultiplierChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "BidCstRewardAmountMultiplierChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2BidCstRewardAmountMultiplierChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "BidCstRewardAmountMultiplierChanged", log); err != nil {
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

// ParseBidCstRewardAmountMultiplierChanged is a log parse operation binding the contract event 0x40b9c59af8c486ccf8c7cc73df5a51e7cc29747ea7d39f99632ecaf9caa2ed1f.
//
// Solidity: event BidCstRewardAmountMultiplierChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseBidCstRewardAmountMultiplierChanged(log types.Log) (*CosmicSignatureGameV2BidCstRewardAmountMultiplierChanged, error) {
	event := new(CosmicSignatureGameV2BidCstRewardAmountMultiplierChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "BidCstRewardAmountMultiplierChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2BidMessageLengthMaxLimitChangedIterator is returned from FilterBidMessageLengthMaxLimitChanged and is used to iterate over the raw logs and unpacked data for BidMessageLengthMaxLimitChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2BidMessageLengthMaxLimitChangedIterator struct {
	Event *CosmicSignatureGameV2BidMessageLengthMaxLimitChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2BidMessageLengthMaxLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2BidMessageLengthMaxLimitChanged)
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
		it.Event = new(CosmicSignatureGameV2BidMessageLengthMaxLimitChanged)
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
func (it *CosmicSignatureGameV2BidMessageLengthMaxLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2BidMessageLengthMaxLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2BidMessageLengthMaxLimitChanged represents a BidMessageLengthMaxLimitChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2BidMessageLengthMaxLimitChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBidMessageLengthMaxLimitChanged is a free log retrieval operation binding the contract event 0x157c413b0549fd4f45aab72b7828304fb2c45dad53de0f1128c5eabf3aaabaf8.
//
// Solidity: event BidMessageLengthMaxLimitChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterBidMessageLengthMaxLimitChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2BidMessageLengthMaxLimitChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "BidMessageLengthMaxLimitChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2BidMessageLengthMaxLimitChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "BidMessageLengthMaxLimitChanged", logs: logs, sub: sub}, nil
}

// WatchBidMessageLengthMaxLimitChanged is a free log subscription operation binding the contract event 0x157c413b0549fd4f45aab72b7828304fb2c45dad53de0f1128c5eabf3aaabaf8.
//
// Solidity: event BidMessageLengthMaxLimitChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchBidMessageLengthMaxLimitChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2BidMessageLengthMaxLimitChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "BidMessageLengthMaxLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2BidMessageLengthMaxLimitChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "BidMessageLengthMaxLimitChanged", log); err != nil {
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

// ParseBidMessageLengthMaxLimitChanged is a log parse operation binding the contract event 0x157c413b0549fd4f45aab72b7828304fb2c45dad53de0f1128c5eabf3aaabaf8.
//
// Solidity: event BidMessageLengthMaxLimitChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseBidMessageLengthMaxLimitChanged(log types.Log) (*CosmicSignatureGameV2BidMessageLengthMaxLimitChanged, error) {
	event := new(CosmicSignatureGameV2BidMessageLengthMaxLimitChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "BidMessageLengthMaxLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2BidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2BidPlacedIterator struct {
	Event *CosmicSignatureGameV2BidPlaced // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2BidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2BidPlaced)
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
		it.Event = new(CosmicSignatureGameV2BidPlaced)
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
func (it *CosmicSignatureGameV2BidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2BidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2BidPlaced represents a BidPlaced event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2BidPlaced struct {
	RoundNum                *big.Int
	LastBidderAddress       common.Address
	PaidEthPrice            *big.Int
	PaidCstPrice            *big.Int
	RandomWalkNftId         *big.Int
	Message                 string
	BidCstRewardAmount      *big.Int
	CstDutchAuctionDuration *big.Int
	MainPrizeTime           *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x1d1f406c89e99504a7222aaba50bf96b9f91f522ebc6ea825255145919e102ec.
//
// Solidity: event BidPlaced(uint256 indexed roundNum, address indexed lastBidderAddress, int256 paidEthPrice, int256 paidCstPrice, int256 indexed randomWalkNftId, string message, uint256 bidCstRewardAmount, uint256 cstDutchAuctionDuration, uint256 mainPrizeTime)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterBidPlaced(opts *bind.FilterOpts, roundNum []*big.Int, lastBidderAddress []common.Address, randomWalkNftId []*big.Int) (*CosmicSignatureGameV2BidPlacedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var lastBidderAddressRule []interface{}
	for _, lastBidderAddressItem := range lastBidderAddress {
		lastBidderAddressRule = append(lastBidderAddressRule, lastBidderAddressItem)
	}

	var randomWalkNftIdRule []interface{}
	for _, randomWalkNftIdItem := range randomWalkNftId {
		randomWalkNftIdRule = append(randomWalkNftIdRule, randomWalkNftIdItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "BidPlaced", roundNumRule, lastBidderAddressRule, randomWalkNftIdRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2BidPlacedIterator{contract: _CosmicSignatureGameV2.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x1d1f406c89e99504a7222aaba50bf96b9f91f522ebc6ea825255145919e102ec.
//
// Solidity: event BidPlaced(uint256 indexed roundNum, address indexed lastBidderAddress, int256 paidEthPrice, int256 paidCstPrice, int256 indexed randomWalkNftId, string message, uint256 bidCstRewardAmount, uint256 cstDutchAuctionDuration, uint256 mainPrizeTime)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2BidPlaced, roundNum []*big.Int, lastBidderAddress []common.Address, randomWalkNftId []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var lastBidderAddressRule []interface{}
	for _, lastBidderAddressItem := range lastBidderAddress {
		lastBidderAddressRule = append(lastBidderAddressRule, lastBidderAddressItem)
	}

	var randomWalkNftIdRule []interface{}
	for _, randomWalkNftIdItem := range randomWalkNftId {
		randomWalkNftIdRule = append(randomWalkNftIdRule, randomWalkNftIdItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "BidPlaced", roundNumRule, lastBidderAddressRule, randomWalkNftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2BidPlaced)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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

// ParseBidPlaced is a log parse operation binding the contract event 0x1d1f406c89e99504a7222aaba50bf96b9f91f522ebc6ea825255145919e102ec.
//
// Solidity: event BidPlaced(uint256 indexed roundNum, address indexed lastBidderAddress, int256 paidEthPrice, int256 paidCstPrice, int256 indexed randomWalkNftId, string message, uint256 bidCstRewardAmount, uint256 cstDutchAuctionDuration, uint256 mainPrizeTime)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseBidPlaced(log types.Log) (*CosmicSignatureGameV2BidPlaced, error) {
	event := new(CosmicSignatureGameV2BidPlaced)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2CharityAddressChangedIterator is returned from FilterCharityAddressChanged and is used to iterate over the raw logs and unpacked data for CharityAddressChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2CharityAddressChangedIterator struct {
	Event *CosmicSignatureGameV2CharityAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2CharityAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2CharityAddressChanged)
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
		it.Event = new(CosmicSignatureGameV2CharityAddressChanged)
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
func (it *CosmicSignatureGameV2CharityAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2CharityAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2CharityAddressChanged represents a CharityAddressChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2CharityAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCharityAddressChanged is a free log retrieval operation binding the contract event 0x1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c.
//
// Solidity: event CharityAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterCharityAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*CosmicSignatureGameV2CharityAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "CharityAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2CharityAddressChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "CharityAddressChanged", logs: logs, sub: sub}, nil
}

// WatchCharityAddressChanged is a free log subscription operation binding the contract event 0x1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c.
//
// Solidity: event CharityAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchCharityAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2CharityAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "CharityAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2CharityAddressChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "CharityAddressChanged", log); err != nil {
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
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseCharityAddressChanged(log types.Log) (*CosmicSignatureGameV2CharityAddressChanged, error) {
	event := new(CosmicSignatureGameV2CharityAddressChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "CharityAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2CharityEthDonationAmountPercentageChangedIterator is returned from FilterCharityEthDonationAmountPercentageChanged and is used to iterate over the raw logs and unpacked data for CharityEthDonationAmountPercentageChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2CharityEthDonationAmountPercentageChangedIterator struct {
	Event *CosmicSignatureGameV2CharityEthDonationAmountPercentageChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2CharityEthDonationAmountPercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2CharityEthDonationAmountPercentageChanged)
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
		it.Event = new(CosmicSignatureGameV2CharityEthDonationAmountPercentageChanged)
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
func (it *CosmicSignatureGameV2CharityEthDonationAmountPercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2CharityEthDonationAmountPercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2CharityEthDonationAmountPercentageChanged represents a CharityEthDonationAmountPercentageChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2CharityEthDonationAmountPercentageChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCharityEthDonationAmountPercentageChanged is a free log retrieval operation binding the contract event 0xfe65b6d5007c66dadebc5095104ccd672c070f396dfdcfe1ee7f54201b9efa6d.
//
// Solidity: event CharityEthDonationAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterCharityEthDonationAmountPercentageChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2CharityEthDonationAmountPercentageChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "CharityEthDonationAmountPercentageChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2CharityEthDonationAmountPercentageChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "CharityEthDonationAmountPercentageChanged", logs: logs, sub: sub}, nil
}

// WatchCharityEthDonationAmountPercentageChanged is a free log subscription operation binding the contract event 0xfe65b6d5007c66dadebc5095104ccd672c070f396dfdcfe1ee7f54201b9efa6d.
//
// Solidity: event CharityEthDonationAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchCharityEthDonationAmountPercentageChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2CharityEthDonationAmountPercentageChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "CharityEthDonationAmountPercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2CharityEthDonationAmountPercentageChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "CharityEthDonationAmountPercentageChanged", log); err != nil {
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

// ParseCharityEthDonationAmountPercentageChanged is a log parse operation binding the contract event 0xfe65b6d5007c66dadebc5095104ccd672c070f396dfdcfe1ee7f54201b9efa6d.
//
// Solidity: event CharityEthDonationAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseCharityEthDonationAmountPercentageChanged(log types.Log) (*CosmicSignatureGameV2CharityEthDonationAmountPercentageChanged, error) {
	event := new(CosmicSignatureGameV2CharityEthDonationAmountPercentageChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "CharityEthDonationAmountPercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2ChronoWarriorEthPrizeAmountPercentageChangedIterator is returned from FilterChronoWarriorEthPrizeAmountPercentageChanged and is used to iterate over the raw logs and unpacked data for ChronoWarriorEthPrizeAmountPercentageChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2ChronoWarriorEthPrizeAmountPercentageChangedIterator struct {
	Event *CosmicSignatureGameV2ChronoWarriorEthPrizeAmountPercentageChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2ChronoWarriorEthPrizeAmountPercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2ChronoWarriorEthPrizeAmountPercentageChanged)
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
		it.Event = new(CosmicSignatureGameV2ChronoWarriorEthPrizeAmountPercentageChanged)
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
func (it *CosmicSignatureGameV2ChronoWarriorEthPrizeAmountPercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2ChronoWarriorEthPrizeAmountPercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2ChronoWarriorEthPrizeAmountPercentageChanged represents a ChronoWarriorEthPrizeAmountPercentageChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2ChronoWarriorEthPrizeAmountPercentageChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChronoWarriorEthPrizeAmountPercentageChanged is a free log retrieval operation binding the contract event 0x5581e31f5b8d4b3c45b8ab8bf67e3602ce588b423905eb6ad34bd6bc3c848699.
//
// Solidity: event ChronoWarriorEthPrizeAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterChronoWarriorEthPrizeAmountPercentageChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2ChronoWarriorEthPrizeAmountPercentageChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "ChronoWarriorEthPrizeAmountPercentageChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2ChronoWarriorEthPrizeAmountPercentageChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "ChronoWarriorEthPrizeAmountPercentageChanged", logs: logs, sub: sub}, nil
}

// WatchChronoWarriorEthPrizeAmountPercentageChanged is a free log subscription operation binding the contract event 0x5581e31f5b8d4b3c45b8ab8bf67e3602ce588b423905eb6ad34bd6bc3c848699.
//
// Solidity: event ChronoWarriorEthPrizeAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchChronoWarriorEthPrizeAmountPercentageChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2ChronoWarriorEthPrizeAmountPercentageChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "ChronoWarriorEthPrizeAmountPercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2ChronoWarriorEthPrizeAmountPercentageChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "ChronoWarriorEthPrizeAmountPercentageChanged", log); err != nil {
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

// ParseChronoWarriorEthPrizeAmountPercentageChanged is a log parse operation binding the contract event 0x5581e31f5b8d4b3c45b8ab8bf67e3602ce588b423905eb6ad34bd6bc3c848699.
//
// Solidity: event ChronoWarriorEthPrizeAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseChronoWarriorEthPrizeAmountPercentageChanged(log types.Log) (*CosmicSignatureGameV2ChronoWarriorEthPrizeAmountPercentageChanged, error) {
	event := new(CosmicSignatureGameV2ChronoWarriorEthPrizeAmountPercentageChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "ChronoWarriorEthPrizeAmountPercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2ChronoWarriorPrizePaidIterator is returned from FilterChronoWarriorPrizePaid and is used to iterate over the raw logs and unpacked data for ChronoWarriorPrizePaid events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2ChronoWarriorPrizePaidIterator struct {
	Event *CosmicSignatureGameV2ChronoWarriorPrizePaid // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2ChronoWarriorPrizePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2ChronoWarriorPrizePaid)
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
		it.Event = new(CosmicSignatureGameV2ChronoWarriorPrizePaid)
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
func (it *CosmicSignatureGameV2ChronoWarriorPrizePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2ChronoWarriorPrizePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2ChronoWarriorPrizePaid represents a ChronoWarriorPrizePaid event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2ChronoWarriorPrizePaid struct {
	RoundNum                  *big.Int
	WinnerIndex               *big.Int
	ChronoWarriorAddress      common.Address
	EthPrizeAmount            *big.Int
	CstPrizeAmount            *big.Int
	PrizeCosmicSignatureNftId *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterChronoWarriorPrizePaid is a free log retrieval operation binding the contract event 0xaa858ae20a26d00a9ea528972d537e68a51a0744226d9ea1fc9b22492dc282a5.
//
// Solidity: event ChronoWarriorPrizePaid(uint256 indexed roundNum, uint256 winnerIndex, address indexed chronoWarriorAddress, uint256 ethPrizeAmount, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterChronoWarriorPrizePaid(opts *bind.FilterOpts, roundNum []*big.Int, chronoWarriorAddress []common.Address, prizeCosmicSignatureNftId []*big.Int) (*CosmicSignatureGameV2ChronoWarriorPrizePaidIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	var chronoWarriorAddressRule []interface{}
	for _, chronoWarriorAddressItem := range chronoWarriorAddress {
		chronoWarriorAddressRule = append(chronoWarriorAddressRule, chronoWarriorAddressItem)
	}

	var prizeCosmicSignatureNftIdRule []interface{}
	for _, prizeCosmicSignatureNftIdItem := range prizeCosmicSignatureNftId {
		prizeCosmicSignatureNftIdRule = append(prizeCosmicSignatureNftIdRule, prizeCosmicSignatureNftIdItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "ChronoWarriorPrizePaid", roundNumRule, chronoWarriorAddressRule, prizeCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2ChronoWarriorPrizePaidIterator{contract: _CosmicSignatureGameV2.contract, event: "ChronoWarriorPrizePaid", logs: logs, sub: sub}, nil
}

// WatchChronoWarriorPrizePaid is a free log subscription operation binding the contract event 0xaa858ae20a26d00a9ea528972d537e68a51a0744226d9ea1fc9b22492dc282a5.
//
// Solidity: event ChronoWarriorPrizePaid(uint256 indexed roundNum, uint256 winnerIndex, address indexed chronoWarriorAddress, uint256 ethPrizeAmount, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchChronoWarriorPrizePaid(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2ChronoWarriorPrizePaid, roundNum []*big.Int, chronoWarriorAddress []common.Address, prizeCosmicSignatureNftId []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	var chronoWarriorAddressRule []interface{}
	for _, chronoWarriorAddressItem := range chronoWarriorAddress {
		chronoWarriorAddressRule = append(chronoWarriorAddressRule, chronoWarriorAddressItem)
	}

	var prizeCosmicSignatureNftIdRule []interface{}
	for _, prizeCosmicSignatureNftIdItem := range prizeCosmicSignatureNftId {
		prizeCosmicSignatureNftIdRule = append(prizeCosmicSignatureNftIdRule, prizeCosmicSignatureNftIdItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "ChronoWarriorPrizePaid", roundNumRule, chronoWarriorAddressRule, prizeCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2ChronoWarriorPrizePaid)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "ChronoWarriorPrizePaid", log); err != nil {
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

// ParseChronoWarriorPrizePaid is a log parse operation binding the contract event 0xaa858ae20a26d00a9ea528972d537e68a51a0744226d9ea1fc9b22492dc282a5.
//
// Solidity: event ChronoWarriorPrizePaid(uint256 indexed roundNum, uint256 winnerIndex, address indexed chronoWarriorAddress, uint256 ethPrizeAmount, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseChronoWarriorPrizePaid(log types.Log) (*CosmicSignatureGameV2ChronoWarriorPrizePaid, error) {
	event := new(CosmicSignatureGameV2ChronoWarriorPrizePaid)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "ChronoWarriorPrizePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2CosmicSignatureNftAddressChangedIterator is returned from FilterCosmicSignatureNftAddressChanged and is used to iterate over the raw logs and unpacked data for CosmicSignatureNftAddressChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2CosmicSignatureNftAddressChangedIterator struct {
	Event *CosmicSignatureGameV2CosmicSignatureNftAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2CosmicSignatureNftAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2CosmicSignatureNftAddressChanged)
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
		it.Event = new(CosmicSignatureGameV2CosmicSignatureNftAddressChanged)
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
func (it *CosmicSignatureGameV2CosmicSignatureNftAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2CosmicSignatureNftAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2CosmicSignatureNftAddressChanged represents a CosmicSignatureNftAddressChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2CosmicSignatureNftAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCosmicSignatureNftAddressChanged is a free log retrieval operation binding the contract event 0x5bde6238168795ba4e77972a2bdaa5a465f7c9a5d05817f5e8d3fed2e5a4fa60.
//
// Solidity: event CosmicSignatureNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterCosmicSignatureNftAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*CosmicSignatureGameV2CosmicSignatureNftAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "CosmicSignatureNftAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2CosmicSignatureNftAddressChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "CosmicSignatureNftAddressChanged", logs: logs, sub: sub}, nil
}

// WatchCosmicSignatureNftAddressChanged is a free log subscription operation binding the contract event 0x5bde6238168795ba4e77972a2bdaa5a465f7c9a5d05817f5e8d3fed2e5a4fa60.
//
// Solidity: event CosmicSignatureNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchCosmicSignatureNftAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2CosmicSignatureNftAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "CosmicSignatureNftAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2CosmicSignatureNftAddressChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "CosmicSignatureNftAddressChanged", log); err != nil {
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

// ParseCosmicSignatureNftAddressChanged is a log parse operation binding the contract event 0x5bde6238168795ba4e77972a2bdaa5a465f7c9a5d05817f5e8d3fed2e5a4fa60.
//
// Solidity: event CosmicSignatureNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseCosmicSignatureNftAddressChanged(log types.Log) (*CosmicSignatureGameV2CosmicSignatureNftAddressChanged, error) {
	event := new(CosmicSignatureGameV2CosmicSignatureNftAddressChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "CosmicSignatureNftAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2CosmicSignatureNftStakingTotalEthRewardAmountPercentageChangedIterator is returned from FilterCosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged and is used to iterate over the raw logs and unpacked data for CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2CosmicSignatureNftStakingTotalEthRewardAmountPercentageChangedIterator struct {
	Event *CosmicSignatureGameV2CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2CosmicSignatureNftStakingTotalEthRewardAmountPercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged)
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
		it.Event = new(CosmicSignatureGameV2CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged)
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
func (it *CosmicSignatureGameV2CosmicSignatureNftStakingTotalEthRewardAmountPercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2CosmicSignatureNftStakingTotalEthRewardAmountPercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged represents a CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged is a free log retrieval operation binding the contract event 0x9e44c04f534af356419a731b967f3d56bc748b1f3fdbda7a89f4e1a321ade934.
//
// Solidity: event CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterCosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2CosmicSignatureNftStakingTotalEthRewardAmountPercentageChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2CosmicSignatureNftStakingTotalEthRewardAmountPercentageChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged", logs: logs, sub: sub}, nil
}

// WatchCosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged is a free log subscription operation binding the contract event 0x9e44c04f534af356419a731b967f3d56bc748b1f3fdbda7a89f4e1a321ade934.
//
// Solidity: event CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchCosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged", log); err != nil {
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

// ParseCosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged is a log parse operation binding the contract event 0x9e44c04f534af356419a731b967f3d56bc748b1f3fdbda7a89f4e1a321ade934.
//
// Solidity: event CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseCosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged(log types.Log) (*CosmicSignatureGameV2CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged, error) {
	event := new(CosmicSignatureGameV2CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2CosmicSignatureTokenAddressChangedIterator is returned from FilterCosmicSignatureTokenAddressChanged and is used to iterate over the raw logs and unpacked data for CosmicSignatureTokenAddressChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2CosmicSignatureTokenAddressChangedIterator struct {
	Event *CosmicSignatureGameV2CosmicSignatureTokenAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2CosmicSignatureTokenAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2CosmicSignatureTokenAddressChanged)
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
		it.Event = new(CosmicSignatureGameV2CosmicSignatureTokenAddressChanged)
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
func (it *CosmicSignatureGameV2CosmicSignatureTokenAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2CosmicSignatureTokenAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2CosmicSignatureTokenAddressChanged represents a CosmicSignatureTokenAddressChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2CosmicSignatureTokenAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCosmicSignatureTokenAddressChanged is a free log retrieval operation binding the contract event 0x9b3eda10f1724a2cf9f7dae4ac263c77908df4d00e92f1377b66fc8be37cd8c1.
//
// Solidity: event CosmicSignatureTokenAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterCosmicSignatureTokenAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*CosmicSignatureGameV2CosmicSignatureTokenAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "CosmicSignatureTokenAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2CosmicSignatureTokenAddressChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "CosmicSignatureTokenAddressChanged", logs: logs, sub: sub}, nil
}

// WatchCosmicSignatureTokenAddressChanged is a free log subscription operation binding the contract event 0x9b3eda10f1724a2cf9f7dae4ac263c77908df4d00e92f1377b66fc8be37cd8c1.
//
// Solidity: event CosmicSignatureTokenAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchCosmicSignatureTokenAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2CosmicSignatureTokenAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "CosmicSignatureTokenAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2CosmicSignatureTokenAddressChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "CosmicSignatureTokenAddressChanged", log); err != nil {
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

// ParseCosmicSignatureTokenAddressChanged is a log parse operation binding the contract event 0x9b3eda10f1724a2cf9f7dae4ac263c77908df4d00e92f1377b66fc8be37cd8c1.
//
// Solidity: event CosmicSignatureTokenAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseCosmicSignatureTokenAddressChanged(log types.Log) (*CosmicSignatureGameV2CosmicSignatureTokenAddressChanged, error) {
	event := new(CosmicSignatureGameV2CosmicSignatureTokenAddressChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "CosmicSignatureTokenAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2CstDutchAuctionBeginningBidPriceMinLimitChangedIterator is returned from FilterCstDutchAuctionBeginningBidPriceMinLimitChanged and is used to iterate over the raw logs and unpacked data for CstDutchAuctionBeginningBidPriceMinLimitChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2CstDutchAuctionBeginningBidPriceMinLimitChangedIterator struct {
	Event *CosmicSignatureGameV2CstDutchAuctionBeginningBidPriceMinLimitChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2CstDutchAuctionBeginningBidPriceMinLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2CstDutchAuctionBeginningBidPriceMinLimitChanged)
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
		it.Event = new(CosmicSignatureGameV2CstDutchAuctionBeginningBidPriceMinLimitChanged)
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
func (it *CosmicSignatureGameV2CstDutchAuctionBeginningBidPriceMinLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2CstDutchAuctionBeginningBidPriceMinLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2CstDutchAuctionBeginningBidPriceMinLimitChanged represents a CstDutchAuctionBeginningBidPriceMinLimitChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2CstDutchAuctionBeginningBidPriceMinLimitChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCstDutchAuctionBeginningBidPriceMinLimitChanged is a free log retrieval operation binding the contract event 0x4e8c80fe79d13b8663de9f6981925ae24c498cc07b0ebd4f4789fa78ca17caff.
//
// Solidity: event CstDutchAuctionBeginningBidPriceMinLimitChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterCstDutchAuctionBeginningBidPriceMinLimitChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2CstDutchAuctionBeginningBidPriceMinLimitChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "CstDutchAuctionBeginningBidPriceMinLimitChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2CstDutchAuctionBeginningBidPriceMinLimitChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "CstDutchAuctionBeginningBidPriceMinLimitChanged", logs: logs, sub: sub}, nil
}

// WatchCstDutchAuctionBeginningBidPriceMinLimitChanged is a free log subscription operation binding the contract event 0x4e8c80fe79d13b8663de9f6981925ae24c498cc07b0ebd4f4789fa78ca17caff.
//
// Solidity: event CstDutchAuctionBeginningBidPriceMinLimitChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchCstDutchAuctionBeginningBidPriceMinLimitChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2CstDutchAuctionBeginningBidPriceMinLimitChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "CstDutchAuctionBeginningBidPriceMinLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2CstDutchAuctionBeginningBidPriceMinLimitChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "CstDutchAuctionBeginningBidPriceMinLimitChanged", log); err != nil {
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

// ParseCstDutchAuctionBeginningBidPriceMinLimitChanged is a log parse operation binding the contract event 0x4e8c80fe79d13b8663de9f6981925ae24c498cc07b0ebd4f4789fa78ca17caff.
//
// Solidity: event CstDutchAuctionBeginningBidPriceMinLimitChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseCstDutchAuctionBeginningBidPriceMinLimitChanged(log types.Log) (*CosmicSignatureGameV2CstDutchAuctionBeginningBidPriceMinLimitChanged, error) {
	event := new(CosmicSignatureGameV2CstDutchAuctionBeginningBidPriceMinLimitChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "CstDutchAuctionBeginningBidPriceMinLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChangedIterator is returned from FilterCstDutchAuctionDurationChangeDivisorChanged and is used to iterate over the raw logs and unpacked data for CstDutchAuctionDurationChangeDivisorChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChangedIterator struct {
	Event *CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChanged)
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
		it.Event = new(CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChanged)
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
func (it *CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChanged represents a CstDutchAuctionDurationChangeDivisorChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCstDutchAuctionDurationChangeDivisorChanged is a free log retrieval operation binding the contract event 0xacbc6b6929088e4b2d043625fa7248e00fb6658a425eb9d9dc8c37b18c7f3e6f.
//
// Solidity: event CstDutchAuctionDurationChangeDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterCstDutchAuctionDurationChangeDivisorChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "CstDutchAuctionDurationChangeDivisorChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "CstDutchAuctionDurationChangeDivisorChanged", logs: logs, sub: sub}, nil
}

// WatchCstDutchAuctionDurationChangeDivisorChanged is a free log subscription operation binding the contract event 0xacbc6b6929088e4b2d043625fa7248e00fb6658a425eb9d9dc8c37b18c7f3e6f.
//
// Solidity: event CstDutchAuctionDurationChangeDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchCstDutchAuctionDurationChangeDivisorChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "CstDutchAuctionDurationChangeDivisorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "CstDutchAuctionDurationChangeDivisorChanged", log); err != nil {
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

// ParseCstDutchAuctionDurationChangeDivisorChanged is a log parse operation binding the contract event 0xacbc6b6929088e4b2d043625fa7248e00fb6658a425eb9d9dc8c37b18c7f3e6f.
//
// Solidity: event CstDutchAuctionDurationChangeDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseCstDutchAuctionDurationChangeDivisorChanged(log types.Log) (*CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChanged, error) {
	event := new(CosmicSignatureGameV2CstDutchAuctionDurationChangeDivisorChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "CstDutchAuctionDurationChangeDivisorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2CstDutchAuctionDurationChangedIterator is returned from FilterCstDutchAuctionDurationChanged and is used to iterate over the raw logs and unpacked data for CstDutchAuctionDurationChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2CstDutchAuctionDurationChangedIterator struct {
	Event *CosmicSignatureGameV2CstDutchAuctionDurationChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2CstDutchAuctionDurationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2CstDutchAuctionDurationChanged)
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
		it.Event = new(CosmicSignatureGameV2CstDutchAuctionDurationChanged)
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
func (it *CosmicSignatureGameV2CstDutchAuctionDurationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2CstDutchAuctionDurationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2CstDutchAuctionDurationChanged represents a CstDutchAuctionDurationChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2CstDutchAuctionDurationChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCstDutchAuctionDurationChanged is a free log retrieval operation binding the contract event 0x4abea08c196329c357e3175d011af39a8625be99ef0ba5a0f3547a95534fedb7.
//
// Solidity: event CstDutchAuctionDurationChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterCstDutchAuctionDurationChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2CstDutchAuctionDurationChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "CstDutchAuctionDurationChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2CstDutchAuctionDurationChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "CstDutchAuctionDurationChanged", logs: logs, sub: sub}, nil
}

// WatchCstDutchAuctionDurationChanged is a free log subscription operation binding the contract event 0x4abea08c196329c357e3175d011af39a8625be99ef0ba5a0f3547a95534fedb7.
//
// Solidity: event CstDutchAuctionDurationChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchCstDutchAuctionDurationChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2CstDutchAuctionDurationChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "CstDutchAuctionDurationChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2CstDutchAuctionDurationChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "CstDutchAuctionDurationChanged", log); err != nil {
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

// ParseCstDutchAuctionDurationChanged is a log parse operation binding the contract event 0x4abea08c196329c357e3175d011af39a8625be99ef0ba5a0f3547a95534fedb7.
//
// Solidity: event CstDutchAuctionDurationChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseCstDutchAuctionDurationChanged(log types.Log) (*CosmicSignatureGameV2CstDutchAuctionDurationChanged, error) {
	event := new(CosmicSignatureGameV2CstDutchAuctionDurationChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "CstDutchAuctionDurationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2CstPrizeAmountChangedIterator is returned from FilterCstPrizeAmountChanged and is used to iterate over the raw logs and unpacked data for CstPrizeAmountChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2CstPrizeAmountChangedIterator struct {
	Event *CosmicSignatureGameV2CstPrizeAmountChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2CstPrizeAmountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2CstPrizeAmountChanged)
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
		it.Event = new(CosmicSignatureGameV2CstPrizeAmountChanged)
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
func (it *CosmicSignatureGameV2CstPrizeAmountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2CstPrizeAmountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2CstPrizeAmountChanged represents a CstPrizeAmountChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2CstPrizeAmountChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCstPrizeAmountChanged is a free log retrieval operation binding the contract event 0xd95e7f967f9370c11deb15ffbb191b9f2e9795ab0738db5bc72bd2794978f32d.
//
// Solidity: event CstPrizeAmountChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterCstPrizeAmountChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2CstPrizeAmountChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "CstPrizeAmountChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2CstPrizeAmountChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "CstPrizeAmountChanged", logs: logs, sub: sub}, nil
}

// WatchCstPrizeAmountChanged is a free log subscription operation binding the contract event 0xd95e7f967f9370c11deb15ffbb191b9f2e9795ab0738db5bc72bd2794978f32d.
//
// Solidity: event CstPrizeAmountChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchCstPrizeAmountChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2CstPrizeAmountChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "CstPrizeAmountChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2CstPrizeAmountChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "CstPrizeAmountChanged", log); err != nil {
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

// ParseCstPrizeAmountChanged is a log parse operation binding the contract event 0xd95e7f967f9370c11deb15ffbb191b9f2e9795ab0738db5bc72bd2794978f32d.
//
// Solidity: event CstPrizeAmountChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseCstPrizeAmountChanged(log types.Log) (*CosmicSignatureGameV2CstPrizeAmountChanged, error) {
	event := new(CosmicSignatureGameV2CstPrizeAmountChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "CstPrizeAmountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2DelayDurationBeforeRoundActivationChangedIterator is returned from FilterDelayDurationBeforeRoundActivationChanged and is used to iterate over the raw logs and unpacked data for DelayDurationBeforeRoundActivationChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2DelayDurationBeforeRoundActivationChangedIterator struct {
	Event *CosmicSignatureGameV2DelayDurationBeforeRoundActivationChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2DelayDurationBeforeRoundActivationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2DelayDurationBeforeRoundActivationChanged)
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
		it.Event = new(CosmicSignatureGameV2DelayDurationBeforeRoundActivationChanged)
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
func (it *CosmicSignatureGameV2DelayDurationBeforeRoundActivationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2DelayDurationBeforeRoundActivationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2DelayDurationBeforeRoundActivationChanged represents a DelayDurationBeforeRoundActivationChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2DelayDurationBeforeRoundActivationChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelayDurationBeforeRoundActivationChanged is a free log retrieval operation binding the contract event 0xb0868a729f47ae3829aaafe3ca2975d3db2148553c854112f598be6d91ef0d28.
//
// Solidity: event DelayDurationBeforeRoundActivationChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterDelayDurationBeforeRoundActivationChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2DelayDurationBeforeRoundActivationChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "DelayDurationBeforeRoundActivationChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2DelayDurationBeforeRoundActivationChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "DelayDurationBeforeRoundActivationChanged", logs: logs, sub: sub}, nil
}

// WatchDelayDurationBeforeRoundActivationChanged is a free log subscription operation binding the contract event 0xb0868a729f47ae3829aaafe3ca2975d3db2148553c854112f598be6d91ef0d28.
//
// Solidity: event DelayDurationBeforeRoundActivationChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchDelayDurationBeforeRoundActivationChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2DelayDurationBeforeRoundActivationChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "DelayDurationBeforeRoundActivationChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2DelayDurationBeforeRoundActivationChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "DelayDurationBeforeRoundActivationChanged", log); err != nil {
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

// ParseDelayDurationBeforeRoundActivationChanged is a log parse operation binding the contract event 0xb0868a729f47ae3829aaafe3ca2975d3db2148553c854112f598be6d91ef0d28.
//
// Solidity: event DelayDurationBeforeRoundActivationChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseDelayDurationBeforeRoundActivationChanged(log types.Log) (*CosmicSignatureGameV2DelayDurationBeforeRoundActivationChanged, error) {
	event := new(CosmicSignatureGameV2DelayDurationBeforeRoundActivationChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "DelayDurationBeforeRoundActivationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2EnduranceChampionPrizePaidIterator is returned from FilterEnduranceChampionPrizePaid and is used to iterate over the raw logs and unpacked data for EnduranceChampionPrizePaid events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2EnduranceChampionPrizePaidIterator struct {
	Event *CosmicSignatureGameV2EnduranceChampionPrizePaid // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2EnduranceChampionPrizePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2EnduranceChampionPrizePaid)
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
		it.Event = new(CosmicSignatureGameV2EnduranceChampionPrizePaid)
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
func (it *CosmicSignatureGameV2EnduranceChampionPrizePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2EnduranceChampionPrizePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2EnduranceChampionPrizePaid represents a EnduranceChampionPrizePaid event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2EnduranceChampionPrizePaid struct {
	RoundNum                  *big.Int
	EnduranceChampionAddress  common.Address
	CstPrizeAmount            *big.Int
	PrizeCosmicSignatureNftId *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterEnduranceChampionPrizePaid is a free log retrieval operation binding the contract event 0x838ec9dd2530548892bff113f5ffb0138d2efc63c7f59bb6571e8c923b631260.
//
// Solidity: event EnduranceChampionPrizePaid(uint256 indexed roundNum, address indexed enduranceChampionAddress, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterEnduranceChampionPrizePaid(opts *bind.FilterOpts, roundNum []*big.Int, enduranceChampionAddress []common.Address, prizeCosmicSignatureNftId []*big.Int) (*CosmicSignatureGameV2EnduranceChampionPrizePaidIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var enduranceChampionAddressRule []interface{}
	for _, enduranceChampionAddressItem := range enduranceChampionAddress {
		enduranceChampionAddressRule = append(enduranceChampionAddressRule, enduranceChampionAddressItem)
	}

	var prizeCosmicSignatureNftIdRule []interface{}
	for _, prizeCosmicSignatureNftIdItem := range prizeCosmicSignatureNftId {
		prizeCosmicSignatureNftIdRule = append(prizeCosmicSignatureNftIdRule, prizeCosmicSignatureNftIdItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "EnduranceChampionPrizePaid", roundNumRule, enduranceChampionAddressRule, prizeCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2EnduranceChampionPrizePaidIterator{contract: _CosmicSignatureGameV2.contract, event: "EnduranceChampionPrizePaid", logs: logs, sub: sub}, nil
}

// WatchEnduranceChampionPrizePaid is a free log subscription operation binding the contract event 0x838ec9dd2530548892bff113f5ffb0138d2efc63c7f59bb6571e8c923b631260.
//
// Solidity: event EnduranceChampionPrizePaid(uint256 indexed roundNum, address indexed enduranceChampionAddress, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchEnduranceChampionPrizePaid(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2EnduranceChampionPrizePaid, roundNum []*big.Int, enduranceChampionAddress []common.Address, prizeCosmicSignatureNftId []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var enduranceChampionAddressRule []interface{}
	for _, enduranceChampionAddressItem := range enduranceChampionAddress {
		enduranceChampionAddressRule = append(enduranceChampionAddressRule, enduranceChampionAddressItem)
	}

	var prizeCosmicSignatureNftIdRule []interface{}
	for _, prizeCosmicSignatureNftIdItem := range prizeCosmicSignatureNftId {
		prizeCosmicSignatureNftIdRule = append(prizeCosmicSignatureNftIdRule, prizeCosmicSignatureNftIdItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "EnduranceChampionPrizePaid", roundNumRule, enduranceChampionAddressRule, prizeCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2EnduranceChampionPrizePaid)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "EnduranceChampionPrizePaid", log); err != nil {
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

// ParseEnduranceChampionPrizePaid is a log parse operation binding the contract event 0x838ec9dd2530548892bff113f5ffb0138d2efc63c7f59bb6571e8c923b631260.
//
// Solidity: event EnduranceChampionPrizePaid(uint256 indexed roundNum, address indexed enduranceChampionAddress, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseEnduranceChampionPrizePaid(log types.Log) (*CosmicSignatureGameV2EnduranceChampionPrizePaid, error) {
	event := new(CosmicSignatureGameV2EnduranceChampionPrizePaid)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "EnduranceChampionPrizePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2EthBidPriceIncreaseDivisorChangedIterator is returned from FilterEthBidPriceIncreaseDivisorChanged and is used to iterate over the raw logs and unpacked data for EthBidPriceIncreaseDivisorChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2EthBidPriceIncreaseDivisorChangedIterator struct {
	Event *CosmicSignatureGameV2EthBidPriceIncreaseDivisorChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2EthBidPriceIncreaseDivisorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2EthBidPriceIncreaseDivisorChanged)
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
		it.Event = new(CosmicSignatureGameV2EthBidPriceIncreaseDivisorChanged)
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
func (it *CosmicSignatureGameV2EthBidPriceIncreaseDivisorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2EthBidPriceIncreaseDivisorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2EthBidPriceIncreaseDivisorChanged represents a EthBidPriceIncreaseDivisorChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2EthBidPriceIncreaseDivisorChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthBidPriceIncreaseDivisorChanged is a free log retrieval operation binding the contract event 0xdeb71e1d117914859ddde484a9810084d9ee399077d5cd8dcbdf8770d33d8ae4.
//
// Solidity: event EthBidPriceIncreaseDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterEthBidPriceIncreaseDivisorChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2EthBidPriceIncreaseDivisorChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "EthBidPriceIncreaseDivisorChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2EthBidPriceIncreaseDivisorChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "EthBidPriceIncreaseDivisorChanged", logs: logs, sub: sub}, nil
}

// WatchEthBidPriceIncreaseDivisorChanged is a free log subscription operation binding the contract event 0xdeb71e1d117914859ddde484a9810084d9ee399077d5cd8dcbdf8770d33d8ae4.
//
// Solidity: event EthBidPriceIncreaseDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchEthBidPriceIncreaseDivisorChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2EthBidPriceIncreaseDivisorChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "EthBidPriceIncreaseDivisorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2EthBidPriceIncreaseDivisorChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "EthBidPriceIncreaseDivisorChanged", log); err != nil {
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

// ParseEthBidPriceIncreaseDivisorChanged is a log parse operation binding the contract event 0xdeb71e1d117914859ddde484a9810084d9ee399077d5cd8dcbdf8770d33d8ae4.
//
// Solidity: event EthBidPriceIncreaseDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseEthBidPriceIncreaseDivisorChanged(log types.Log) (*CosmicSignatureGameV2EthBidPriceIncreaseDivisorChanged, error) {
	event := new(CosmicSignatureGameV2EthBidPriceIncreaseDivisorChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "EthBidPriceIncreaseDivisorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2EthBidRefundAmountInGasToSwallowMaxLimitChangedIterator is returned from FilterEthBidRefundAmountInGasToSwallowMaxLimitChanged and is used to iterate over the raw logs and unpacked data for EthBidRefundAmountInGasToSwallowMaxLimitChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2EthBidRefundAmountInGasToSwallowMaxLimitChangedIterator struct {
	Event *CosmicSignatureGameV2EthBidRefundAmountInGasToSwallowMaxLimitChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2EthBidRefundAmountInGasToSwallowMaxLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2EthBidRefundAmountInGasToSwallowMaxLimitChanged)
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
		it.Event = new(CosmicSignatureGameV2EthBidRefundAmountInGasToSwallowMaxLimitChanged)
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
func (it *CosmicSignatureGameV2EthBidRefundAmountInGasToSwallowMaxLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2EthBidRefundAmountInGasToSwallowMaxLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2EthBidRefundAmountInGasToSwallowMaxLimitChanged represents a EthBidRefundAmountInGasToSwallowMaxLimitChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2EthBidRefundAmountInGasToSwallowMaxLimitChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthBidRefundAmountInGasToSwallowMaxLimitChanged is a free log retrieval operation binding the contract event 0xa787f26546d7eeea63d70fc31736f27ad28329e95982f3bc5a7e0280f497bbf5.
//
// Solidity: event EthBidRefundAmountInGasToSwallowMaxLimitChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterEthBidRefundAmountInGasToSwallowMaxLimitChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2EthBidRefundAmountInGasToSwallowMaxLimitChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "EthBidRefundAmountInGasToSwallowMaxLimitChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2EthBidRefundAmountInGasToSwallowMaxLimitChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "EthBidRefundAmountInGasToSwallowMaxLimitChanged", logs: logs, sub: sub}, nil
}

// WatchEthBidRefundAmountInGasToSwallowMaxLimitChanged is a free log subscription operation binding the contract event 0xa787f26546d7eeea63d70fc31736f27ad28329e95982f3bc5a7e0280f497bbf5.
//
// Solidity: event EthBidRefundAmountInGasToSwallowMaxLimitChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchEthBidRefundAmountInGasToSwallowMaxLimitChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2EthBidRefundAmountInGasToSwallowMaxLimitChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "EthBidRefundAmountInGasToSwallowMaxLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2EthBidRefundAmountInGasToSwallowMaxLimitChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "EthBidRefundAmountInGasToSwallowMaxLimitChanged", log); err != nil {
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

// ParseEthBidRefundAmountInGasToSwallowMaxLimitChanged is a log parse operation binding the contract event 0xa787f26546d7eeea63d70fc31736f27ad28329e95982f3bc5a7e0280f497bbf5.
//
// Solidity: event EthBidRefundAmountInGasToSwallowMaxLimitChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseEthBidRefundAmountInGasToSwallowMaxLimitChanged(log types.Log) (*CosmicSignatureGameV2EthBidRefundAmountInGasToSwallowMaxLimitChanged, error) {
	event := new(CosmicSignatureGameV2EthBidRefundAmountInGasToSwallowMaxLimitChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "EthBidRefundAmountInGasToSwallowMaxLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2EthDonatedIterator is returned from FilterEthDonated and is used to iterate over the raw logs and unpacked data for EthDonated events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2EthDonatedIterator struct {
	Event *CosmicSignatureGameV2EthDonated // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2EthDonatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2EthDonated)
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
		it.Event = new(CosmicSignatureGameV2EthDonated)
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
func (it *CosmicSignatureGameV2EthDonatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2EthDonatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2EthDonated represents a EthDonated event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2EthDonated struct {
	RoundNum     *big.Int
	DonorAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEthDonated is a free log retrieval operation binding the contract event 0xe32cacf203d00685e2b4d8b0a90e7cd8f3f8a208fdf116f4bb36abe08b7d548e.
//
// Solidity: event EthDonated(uint256 indexed roundNum, address indexed donorAddress, uint256 amount)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterEthDonated(opts *bind.FilterOpts, roundNum []*big.Int, donorAddress []common.Address) (*CosmicSignatureGameV2EthDonatedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "EthDonated", roundNumRule, donorAddressRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2EthDonatedIterator{contract: _CosmicSignatureGameV2.contract, event: "EthDonated", logs: logs, sub: sub}, nil
}

// WatchEthDonated is a free log subscription operation binding the contract event 0xe32cacf203d00685e2b4d8b0a90e7cd8f3f8a208fdf116f4bb36abe08b7d548e.
//
// Solidity: event EthDonated(uint256 indexed roundNum, address indexed donorAddress, uint256 amount)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchEthDonated(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2EthDonated, roundNum []*big.Int, donorAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "EthDonated", roundNumRule, donorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2EthDonated)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "EthDonated", log); err != nil {
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

// ParseEthDonated is a log parse operation binding the contract event 0xe32cacf203d00685e2b4d8b0a90e7cd8f3f8a208fdf116f4bb36abe08b7d548e.
//
// Solidity: event EthDonated(uint256 indexed roundNum, address indexed donorAddress, uint256 amount)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseEthDonated(log types.Log) (*CosmicSignatureGameV2EthDonated, error) {
	event := new(CosmicSignatureGameV2EthDonated)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "EthDonated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2EthDonatedWithInfoIterator is returned from FilterEthDonatedWithInfo and is used to iterate over the raw logs and unpacked data for EthDonatedWithInfo events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2EthDonatedWithInfoIterator struct {
	Event *CosmicSignatureGameV2EthDonatedWithInfo // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2EthDonatedWithInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2EthDonatedWithInfo)
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
		it.Event = new(CosmicSignatureGameV2EthDonatedWithInfo)
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
func (it *CosmicSignatureGameV2EthDonatedWithInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2EthDonatedWithInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2EthDonatedWithInfo represents a EthDonatedWithInfo event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2EthDonatedWithInfo struct {
	RoundNum                       *big.Int
	DonorAddress                   common.Address
	Amount                         *big.Int
	EthDonationWithInfoRecordIndex *big.Int
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterEthDonatedWithInfo is a free log retrieval operation binding the contract event 0xa08049565b10d44a06dca9bf05685b39bc370352043c5a003e8d35d45ebdc53f.
//
// Solidity: event EthDonatedWithInfo(uint256 indexed roundNum, address indexed donorAddress, uint256 amount, uint256 indexed ethDonationWithInfoRecordIndex)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterEthDonatedWithInfo(opts *bind.FilterOpts, roundNum []*big.Int, donorAddress []common.Address, ethDonationWithInfoRecordIndex []*big.Int) (*CosmicSignatureGameV2EthDonatedWithInfoIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}

	var ethDonationWithInfoRecordIndexRule []interface{}
	for _, ethDonationWithInfoRecordIndexItem := range ethDonationWithInfoRecordIndex {
		ethDonationWithInfoRecordIndexRule = append(ethDonationWithInfoRecordIndexRule, ethDonationWithInfoRecordIndexItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "EthDonatedWithInfo", roundNumRule, donorAddressRule, ethDonationWithInfoRecordIndexRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2EthDonatedWithInfoIterator{contract: _CosmicSignatureGameV2.contract, event: "EthDonatedWithInfo", logs: logs, sub: sub}, nil
}

// WatchEthDonatedWithInfo is a free log subscription operation binding the contract event 0xa08049565b10d44a06dca9bf05685b39bc370352043c5a003e8d35d45ebdc53f.
//
// Solidity: event EthDonatedWithInfo(uint256 indexed roundNum, address indexed donorAddress, uint256 amount, uint256 indexed ethDonationWithInfoRecordIndex)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchEthDonatedWithInfo(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2EthDonatedWithInfo, roundNum []*big.Int, donorAddress []common.Address, ethDonationWithInfoRecordIndex []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}

	var ethDonationWithInfoRecordIndexRule []interface{}
	for _, ethDonationWithInfoRecordIndexItem := range ethDonationWithInfoRecordIndex {
		ethDonationWithInfoRecordIndexRule = append(ethDonationWithInfoRecordIndexRule, ethDonationWithInfoRecordIndexItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "EthDonatedWithInfo", roundNumRule, donorAddressRule, ethDonationWithInfoRecordIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2EthDonatedWithInfo)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "EthDonatedWithInfo", log); err != nil {
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

// ParseEthDonatedWithInfo is a log parse operation binding the contract event 0xa08049565b10d44a06dca9bf05685b39bc370352043c5a003e8d35d45ebdc53f.
//
// Solidity: event EthDonatedWithInfo(uint256 indexed roundNum, address indexed donorAddress, uint256 amount, uint256 indexed ethDonationWithInfoRecordIndex)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseEthDonatedWithInfo(log types.Log) (*CosmicSignatureGameV2EthDonatedWithInfo, error) {
	event := new(CosmicSignatureGameV2EthDonatedWithInfo)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "EthDonatedWithInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2EthDutchAuctionDurationDivisorChangedIterator is returned from FilterEthDutchAuctionDurationDivisorChanged and is used to iterate over the raw logs and unpacked data for EthDutchAuctionDurationDivisorChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2EthDutchAuctionDurationDivisorChangedIterator struct {
	Event *CosmicSignatureGameV2EthDutchAuctionDurationDivisorChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2EthDutchAuctionDurationDivisorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2EthDutchAuctionDurationDivisorChanged)
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
		it.Event = new(CosmicSignatureGameV2EthDutchAuctionDurationDivisorChanged)
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
func (it *CosmicSignatureGameV2EthDutchAuctionDurationDivisorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2EthDutchAuctionDurationDivisorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2EthDutchAuctionDurationDivisorChanged represents a EthDutchAuctionDurationDivisorChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2EthDutchAuctionDurationDivisorChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDutchAuctionDurationDivisorChanged is a free log retrieval operation binding the contract event 0xfdf6043c89a5f304289122dcc9f8bd78bb111b5d4f409e2fc2e6c141a1110b79.
//
// Solidity: event EthDutchAuctionDurationDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterEthDutchAuctionDurationDivisorChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2EthDutchAuctionDurationDivisorChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "EthDutchAuctionDurationDivisorChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2EthDutchAuctionDurationDivisorChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "EthDutchAuctionDurationDivisorChanged", logs: logs, sub: sub}, nil
}

// WatchEthDutchAuctionDurationDivisorChanged is a free log subscription operation binding the contract event 0xfdf6043c89a5f304289122dcc9f8bd78bb111b5d4f409e2fc2e6c141a1110b79.
//
// Solidity: event EthDutchAuctionDurationDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchEthDutchAuctionDurationDivisorChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2EthDutchAuctionDurationDivisorChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "EthDutchAuctionDurationDivisorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2EthDutchAuctionDurationDivisorChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "EthDutchAuctionDurationDivisorChanged", log); err != nil {
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

// ParseEthDutchAuctionDurationDivisorChanged is a log parse operation binding the contract event 0xfdf6043c89a5f304289122dcc9f8bd78bb111b5d4f409e2fc2e6c141a1110b79.
//
// Solidity: event EthDutchAuctionDurationDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseEthDutchAuctionDurationDivisorChanged(log types.Log) (*CosmicSignatureGameV2EthDutchAuctionDurationDivisorChanged, error) {
	event := new(CosmicSignatureGameV2EthDutchAuctionDurationDivisorChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "EthDutchAuctionDurationDivisorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2EthDutchAuctionEndingBidPriceDivisorChangedIterator is returned from FilterEthDutchAuctionEndingBidPriceDivisorChanged and is used to iterate over the raw logs and unpacked data for EthDutchAuctionEndingBidPriceDivisorChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2EthDutchAuctionEndingBidPriceDivisorChangedIterator struct {
	Event *CosmicSignatureGameV2EthDutchAuctionEndingBidPriceDivisorChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2EthDutchAuctionEndingBidPriceDivisorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2EthDutchAuctionEndingBidPriceDivisorChanged)
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
		it.Event = new(CosmicSignatureGameV2EthDutchAuctionEndingBidPriceDivisorChanged)
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
func (it *CosmicSignatureGameV2EthDutchAuctionEndingBidPriceDivisorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2EthDutchAuctionEndingBidPriceDivisorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2EthDutchAuctionEndingBidPriceDivisorChanged represents a EthDutchAuctionEndingBidPriceDivisorChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2EthDutchAuctionEndingBidPriceDivisorChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDutchAuctionEndingBidPriceDivisorChanged is a free log retrieval operation binding the contract event 0xb6f6af60099e44041a78b3561ed029b98bf03fdb0efbbb2eb15e1f3d7d923037.
//
// Solidity: event EthDutchAuctionEndingBidPriceDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterEthDutchAuctionEndingBidPriceDivisorChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2EthDutchAuctionEndingBidPriceDivisorChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "EthDutchAuctionEndingBidPriceDivisorChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2EthDutchAuctionEndingBidPriceDivisorChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "EthDutchAuctionEndingBidPriceDivisorChanged", logs: logs, sub: sub}, nil
}

// WatchEthDutchAuctionEndingBidPriceDivisorChanged is a free log subscription operation binding the contract event 0xb6f6af60099e44041a78b3561ed029b98bf03fdb0efbbb2eb15e1f3d7d923037.
//
// Solidity: event EthDutchAuctionEndingBidPriceDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchEthDutchAuctionEndingBidPriceDivisorChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2EthDutchAuctionEndingBidPriceDivisorChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "EthDutchAuctionEndingBidPriceDivisorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2EthDutchAuctionEndingBidPriceDivisorChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "EthDutchAuctionEndingBidPriceDivisorChanged", log); err != nil {
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

// ParseEthDutchAuctionEndingBidPriceDivisorChanged is a log parse operation binding the contract event 0xb6f6af60099e44041a78b3561ed029b98bf03fdb0efbbb2eb15e1f3d7d923037.
//
// Solidity: event EthDutchAuctionEndingBidPriceDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseEthDutchAuctionEndingBidPriceDivisorChanged(log types.Log) (*CosmicSignatureGameV2EthDutchAuctionEndingBidPriceDivisorChanged, error) {
	event := new(CosmicSignatureGameV2EthDutchAuctionEndingBidPriceDivisorChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "EthDutchAuctionEndingBidPriceDivisorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2FirstBidPlacedInRoundIterator is returned from FilterFirstBidPlacedInRound and is used to iterate over the raw logs and unpacked data for FirstBidPlacedInRound events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2FirstBidPlacedInRoundIterator struct {
	Event *CosmicSignatureGameV2FirstBidPlacedInRound // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2FirstBidPlacedInRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2FirstBidPlacedInRound)
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
		it.Event = new(CosmicSignatureGameV2FirstBidPlacedInRound)
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
func (it *CosmicSignatureGameV2FirstBidPlacedInRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2FirstBidPlacedInRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2FirstBidPlacedInRound represents a FirstBidPlacedInRound event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2FirstBidPlacedInRound struct {
	RoundNum       *big.Int
	BlockTimeStamp *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFirstBidPlacedInRound is a free log retrieval operation binding the contract event 0x028a52641badd593b7f30072734c0b97e449213f55b5c3663756427340accd3c.
//
// Solidity: event FirstBidPlacedInRound(uint256 indexed roundNum, uint256 blockTimeStamp)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterFirstBidPlacedInRound(opts *bind.FilterOpts, roundNum []*big.Int) (*CosmicSignatureGameV2FirstBidPlacedInRoundIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "FirstBidPlacedInRound", roundNumRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2FirstBidPlacedInRoundIterator{contract: _CosmicSignatureGameV2.contract, event: "FirstBidPlacedInRound", logs: logs, sub: sub}, nil
}

// WatchFirstBidPlacedInRound is a free log subscription operation binding the contract event 0x028a52641badd593b7f30072734c0b97e449213f55b5c3663756427340accd3c.
//
// Solidity: event FirstBidPlacedInRound(uint256 indexed roundNum, uint256 blockTimeStamp)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchFirstBidPlacedInRound(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2FirstBidPlacedInRound, roundNum []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "FirstBidPlacedInRound", roundNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2FirstBidPlacedInRound)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "FirstBidPlacedInRound", log); err != nil {
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

// ParseFirstBidPlacedInRound is a log parse operation binding the contract event 0x028a52641badd593b7f30072734c0b97e449213f55b5c3663756427340accd3c.
//
// Solidity: event FirstBidPlacedInRound(uint256 indexed roundNum, uint256 blockTimeStamp)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseFirstBidPlacedInRound(log types.Log) (*CosmicSignatureGameV2FirstBidPlacedInRound, error) {
	event := new(CosmicSignatureGameV2FirstBidPlacedInRound)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "FirstBidPlacedInRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2FundTransferFailedIterator is returned from FilterFundTransferFailed and is used to iterate over the raw logs and unpacked data for FundTransferFailed events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2FundTransferFailedIterator struct {
	Event *CosmicSignatureGameV2FundTransferFailed // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2FundTransferFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2FundTransferFailed)
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
		it.Event = new(CosmicSignatureGameV2FundTransferFailed)
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
func (it *CosmicSignatureGameV2FundTransferFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2FundTransferFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2FundTransferFailed represents a FundTransferFailed event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2FundTransferFailed struct {
	ErrStr             string
	DestinationAddress common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFundTransferFailed is a free log retrieval operation binding the contract event 0x154fb6c686c977314af35b730a16b83facda5265b2abec7237c730f63e42215a.
//
// Solidity: event FundTransferFailed(string errStr, address indexed destinationAddress, uint256 amount)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterFundTransferFailed(opts *bind.FilterOpts, destinationAddress []common.Address) (*CosmicSignatureGameV2FundTransferFailedIterator, error) {

	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "FundTransferFailed", destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2FundTransferFailedIterator{contract: _CosmicSignatureGameV2.contract, event: "FundTransferFailed", logs: logs, sub: sub}, nil
}

// WatchFundTransferFailed is a free log subscription operation binding the contract event 0x154fb6c686c977314af35b730a16b83facda5265b2abec7237c730f63e42215a.
//
// Solidity: event FundTransferFailed(string errStr, address indexed destinationAddress, uint256 amount)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchFundTransferFailed(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2FundTransferFailed, destinationAddress []common.Address) (event.Subscription, error) {

	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "FundTransferFailed", destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2FundTransferFailed)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "FundTransferFailed", log); err != nil {
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

// ParseFundTransferFailed is a log parse operation binding the contract event 0x154fb6c686c977314af35b730a16b83facda5265b2abec7237c730f63e42215a.
//
// Solidity: event FundTransferFailed(string errStr, address indexed destinationAddress, uint256 amount)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseFundTransferFailed(log types.Log) (*CosmicSignatureGameV2FundTransferFailed, error) {
	event := new(CosmicSignatureGameV2FundTransferFailed)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "FundTransferFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2FundsTransferredToCharityIterator is returned from FilterFundsTransferredToCharity and is used to iterate over the raw logs and unpacked data for FundsTransferredToCharity events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2FundsTransferredToCharityIterator struct {
	Event *CosmicSignatureGameV2FundsTransferredToCharity // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2FundsTransferredToCharityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2FundsTransferredToCharity)
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
		it.Event = new(CosmicSignatureGameV2FundsTransferredToCharity)
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
func (it *CosmicSignatureGameV2FundsTransferredToCharityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2FundsTransferredToCharityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2FundsTransferredToCharity represents a FundsTransferredToCharity event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2FundsTransferredToCharity struct {
	CharityAddress common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFundsTransferredToCharity is a free log retrieval operation binding the contract event 0x1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d.
//
// Solidity: event FundsTransferredToCharity(address indexed charityAddress, uint256 amount)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterFundsTransferredToCharity(opts *bind.FilterOpts, charityAddress []common.Address) (*CosmicSignatureGameV2FundsTransferredToCharityIterator, error) {

	var charityAddressRule []interface{}
	for _, charityAddressItem := range charityAddress {
		charityAddressRule = append(charityAddressRule, charityAddressItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "FundsTransferredToCharity", charityAddressRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2FundsTransferredToCharityIterator{contract: _CosmicSignatureGameV2.contract, event: "FundsTransferredToCharity", logs: logs, sub: sub}, nil
}

// WatchFundsTransferredToCharity is a free log subscription operation binding the contract event 0x1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d.
//
// Solidity: event FundsTransferredToCharity(address indexed charityAddress, uint256 amount)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchFundsTransferredToCharity(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2FundsTransferredToCharity, charityAddress []common.Address) (event.Subscription, error) {

	var charityAddressRule []interface{}
	for _, charityAddressItem := range charityAddress {
		charityAddressRule = append(charityAddressRule, charityAddressItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "FundsTransferredToCharity", charityAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2FundsTransferredToCharity)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "FundsTransferredToCharity", log); err != nil {
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
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseFundsTransferredToCharity(log types.Log) (*CosmicSignatureGameV2FundsTransferredToCharity, error) {
	event := new(CosmicSignatureGameV2FundsTransferredToCharity)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "FundsTransferredToCharity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2InitialDurationUntilMainPrizeDivisorChangedIterator is returned from FilterInitialDurationUntilMainPrizeDivisorChanged and is used to iterate over the raw logs and unpacked data for InitialDurationUntilMainPrizeDivisorChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2InitialDurationUntilMainPrizeDivisorChangedIterator struct {
	Event *CosmicSignatureGameV2InitialDurationUntilMainPrizeDivisorChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2InitialDurationUntilMainPrizeDivisorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2InitialDurationUntilMainPrizeDivisorChanged)
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
		it.Event = new(CosmicSignatureGameV2InitialDurationUntilMainPrizeDivisorChanged)
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
func (it *CosmicSignatureGameV2InitialDurationUntilMainPrizeDivisorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2InitialDurationUntilMainPrizeDivisorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2InitialDurationUntilMainPrizeDivisorChanged represents a InitialDurationUntilMainPrizeDivisorChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2InitialDurationUntilMainPrizeDivisorChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInitialDurationUntilMainPrizeDivisorChanged is a free log retrieval operation binding the contract event 0xb5edd1f338b34c8f5dd3b1c5cc12f05653c495713c282bf588d34cf14fad0f89.
//
// Solidity: event InitialDurationUntilMainPrizeDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterInitialDurationUntilMainPrizeDivisorChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2InitialDurationUntilMainPrizeDivisorChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "InitialDurationUntilMainPrizeDivisorChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2InitialDurationUntilMainPrizeDivisorChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "InitialDurationUntilMainPrizeDivisorChanged", logs: logs, sub: sub}, nil
}

// WatchInitialDurationUntilMainPrizeDivisorChanged is a free log subscription operation binding the contract event 0xb5edd1f338b34c8f5dd3b1c5cc12f05653c495713c282bf588d34cf14fad0f89.
//
// Solidity: event InitialDurationUntilMainPrizeDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchInitialDurationUntilMainPrizeDivisorChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2InitialDurationUntilMainPrizeDivisorChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "InitialDurationUntilMainPrizeDivisorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2InitialDurationUntilMainPrizeDivisorChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "InitialDurationUntilMainPrizeDivisorChanged", log); err != nil {
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

// ParseInitialDurationUntilMainPrizeDivisorChanged is a log parse operation binding the contract event 0xb5edd1f338b34c8f5dd3b1c5cc12f05653c495713c282bf588d34cf14fad0f89.
//
// Solidity: event InitialDurationUntilMainPrizeDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseInitialDurationUntilMainPrizeDivisorChanged(log types.Log) (*CosmicSignatureGameV2InitialDurationUntilMainPrizeDivisorChanged, error) {
	event := new(CosmicSignatureGameV2InitialDurationUntilMainPrizeDivisorChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "InitialDurationUntilMainPrizeDivisorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2InitializedIterator struct {
	Event *CosmicSignatureGameV2Initialized // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2Initialized)
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
		it.Event = new(CosmicSignatureGameV2Initialized)
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
func (it *CosmicSignatureGameV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2Initialized represents a Initialized event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*CosmicSignatureGameV2InitializedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2InitializedIterator{contract: _CosmicSignatureGameV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2Initialized) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2Initialized)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseInitialized(log types.Log) (*CosmicSignatureGameV2Initialized, error) {
	event := new(CosmicSignatureGameV2Initialized)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2LastCstBidderPrizePaidIterator is returned from FilterLastCstBidderPrizePaid and is used to iterate over the raw logs and unpacked data for LastCstBidderPrizePaid events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2LastCstBidderPrizePaidIterator struct {
	Event *CosmicSignatureGameV2LastCstBidderPrizePaid // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2LastCstBidderPrizePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2LastCstBidderPrizePaid)
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
		it.Event = new(CosmicSignatureGameV2LastCstBidderPrizePaid)
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
func (it *CosmicSignatureGameV2LastCstBidderPrizePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2LastCstBidderPrizePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2LastCstBidderPrizePaid represents a LastCstBidderPrizePaid event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2LastCstBidderPrizePaid struct {
	RoundNum                  *big.Int
	LastCstBidderAddress      common.Address
	CstPrizeAmount            *big.Int
	PrizeCosmicSignatureNftId *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterLastCstBidderPrizePaid is a free log retrieval operation binding the contract event 0x3901b6430c99dc290ee88ff84c4de6091ad7eac335b58e92ef5cbb0793abf4f6.
//
// Solidity: event LastCstBidderPrizePaid(uint256 indexed roundNum, address indexed lastCstBidderAddress, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterLastCstBidderPrizePaid(opts *bind.FilterOpts, roundNum []*big.Int, lastCstBidderAddress []common.Address, prizeCosmicSignatureNftId []*big.Int) (*CosmicSignatureGameV2LastCstBidderPrizePaidIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var lastCstBidderAddressRule []interface{}
	for _, lastCstBidderAddressItem := range lastCstBidderAddress {
		lastCstBidderAddressRule = append(lastCstBidderAddressRule, lastCstBidderAddressItem)
	}

	var prizeCosmicSignatureNftIdRule []interface{}
	for _, prizeCosmicSignatureNftIdItem := range prizeCosmicSignatureNftId {
		prizeCosmicSignatureNftIdRule = append(prizeCosmicSignatureNftIdRule, prizeCosmicSignatureNftIdItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "LastCstBidderPrizePaid", roundNumRule, lastCstBidderAddressRule, prizeCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2LastCstBidderPrizePaidIterator{contract: _CosmicSignatureGameV2.contract, event: "LastCstBidderPrizePaid", logs: logs, sub: sub}, nil
}

// WatchLastCstBidderPrizePaid is a free log subscription operation binding the contract event 0x3901b6430c99dc290ee88ff84c4de6091ad7eac335b58e92ef5cbb0793abf4f6.
//
// Solidity: event LastCstBidderPrizePaid(uint256 indexed roundNum, address indexed lastCstBidderAddress, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchLastCstBidderPrizePaid(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2LastCstBidderPrizePaid, roundNum []*big.Int, lastCstBidderAddress []common.Address, prizeCosmicSignatureNftId []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var lastCstBidderAddressRule []interface{}
	for _, lastCstBidderAddressItem := range lastCstBidderAddress {
		lastCstBidderAddressRule = append(lastCstBidderAddressRule, lastCstBidderAddressItem)
	}

	var prizeCosmicSignatureNftIdRule []interface{}
	for _, prizeCosmicSignatureNftIdItem := range prizeCosmicSignatureNftId {
		prizeCosmicSignatureNftIdRule = append(prizeCosmicSignatureNftIdRule, prizeCosmicSignatureNftIdItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "LastCstBidderPrizePaid", roundNumRule, lastCstBidderAddressRule, prizeCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2LastCstBidderPrizePaid)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "LastCstBidderPrizePaid", log); err != nil {
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

// ParseLastCstBidderPrizePaid is a log parse operation binding the contract event 0x3901b6430c99dc290ee88ff84c4de6091ad7eac335b58e92ef5cbb0793abf4f6.
//
// Solidity: event LastCstBidderPrizePaid(uint256 indexed roundNum, address indexed lastCstBidderAddress, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseLastCstBidderPrizePaid(log types.Log) (*CosmicSignatureGameV2LastCstBidderPrizePaid, error) {
	event := new(CosmicSignatureGameV2LastCstBidderPrizePaid)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "LastCstBidderPrizePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2MainEthPrizeAmountPercentageChangedIterator is returned from FilterMainEthPrizeAmountPercentageChanged and is used to iterate over the raw logs and unpacked data for MainEthPrizeAmountPercentageChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2MainEthPrizeAmountPercentageChangedIterator struct {
	Event *CosmicSignatureGameV2MainEthPrizeAmountPercentageChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2MainEthPrizeAmountPercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2MainEthPrizeAmountPercentageChanged)
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
		it.Event = new(CosmicSignatureGameV2MainEthPrizeAmountPercentageChanged)
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
func (it *CosmicSignatureGameV2MainEthPrizeAmountPercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2MainEthPrizeAmountPercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2MainEthPrizeAmountPercentageChanged represents a MainEthPrizeAmountPercentageChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2MainEthPrizeAmountPercentageChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMainEthPrizeAmountPercentageChanged is a free log retrieval operation binding the contract event 0xb5a05ec7911dd5450a7fa4ae54d50d9d83af6e256db8fc76c82edd7b659cf8bc.
//
// Solidity: event MainEthPrizeAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterMainEthPrizeAmountPercentageChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2MainEthPrizeAmountPercentageChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "MainEthPrizeAmountPercentageChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2MainEthPrizeAmountPercentageChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "MainEthPrizeAmountPercentageChanged", logs: logs, sub: sub}, nil
}

// WatchMainEthPrizeAmountPercentageChanged is a free log subscription operation binding the contract event 0xb5a05ec7911dd5450a7fa4ae54d50d9d83af6e256db8fc76c82edd7b659cf8bc.
//
// Solidity: event MainEthPrizeAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchMainEthPrizeAmountPercentageChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2MainEthPrizeAmountPercentageChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "MainEthPrizeAmountPercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2MainEthPrizeAmountPercentageChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "MainEthPrizeAmountPercentageChanged", log); err != nil {
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

// ParseMainEthPrizeAmountPercentageChanged is a log parse operation binding the contract event 0xb5a05ec7911dd5450a7fa4ae54d50d9d83af6e256db8fc76c82edd7b659cf8bc.
//
// Solidity: event MainEthPrizeAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseMainEthPrizeAmountPercentageChanged(log types.Log) (*CosmicSignatureGameV2MainEthPrizeAmountPercentageChanged, error) {
	event := new(CosmicSignatureGameV2MainEthPrizeAmountPercentageChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "MainEthPrizeAmountPercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2MainPrizeClaimedIterator is returned from FilterMainPrizeClaimed and is used to iterate over the raw logs and unpacked data for MainPrizeClaimed events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2MainPrizeClaimedIterator struct {
	Event *CosmicSignatureGameV2MainPrizeClaimed // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2MainPrizeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2MainPrizeClaimed)
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
		it.Event = new(CosmicSignatureGameV2MainPrizeClaimed)
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
func (it *CosmicSignatureGameV2MainPrizeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2MainPrizeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2MainPrizeClaimed represents a MainPrizeClaimed event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2MainPrizeClaimed struct {
	RoundNum                             *big.Int
	BeneficiaryAddress                   common.Address
	EthPrizeAmount                       *big.Int
	CstPrizeAmount                       *big.Int
	PrizeCosmicSignatureNftId            *big.Int
	TimeoutTimeToWithdrawSecondaryPrizes *big.Int
	Raw                                  types.Log // Blockchain specific contextual infos
}

// FilterMainPrizeClaimed is a free log retrieval operation binding the contract event 0x8c551ec2b6f186753e27f1cf46f84b57f4f83f721e8c1e6170ae512845ced591.
//
// Solidity: event MainPrizeClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, uint256 ethPrizeAmount, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId, uint256 timeoutTimeToWithdrawSecondaryPrizes)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterMainPrizeClaimed(opts *bind.FilterOpts, roundNum []*big.Int, beneficiaryAddress []common.Address, prizeCosmicSignatureNftId []*big.Int) (*CosmicSignatureGameV2MainPrizeClaimedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	var prizeCosmicSignatureNftIdRule []interface{}
	for _, prizeCosmicSignatureNftIdItem := range prizeCosmicSignatureNftId {
		prizeCosmicSignatureNftIdRule = append(prizeCosmicSignatureNftIdRule, prizeCosmicSignatureNftIdItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "MainPrizeClaimed", roundNumRule, beneficiaryAddressRule, prizeCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2MainPrizeClaimedIterator{contract: _CosmicSignatureGameV2.contract, event: "MainPrizeClaimed", logs: logs, sub: sub}, nil
}

// WatchMainPrizeClaimed is a free log subscription operation binding the contract event 0x8c551ec2b6f186753e27f1cf46f84b57f4f83f721e8c1e6170ae512845ced591.
//
// Solidity: event MainPrizeClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, uint256 ethPrizeAmount, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId, uint256 timeoutTimeToWithdrawSecondaryPrizes)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchMainPrizeClaimed(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2MainPrizeClaimed, roundNum []*big.Int, beneficiaryAddress []common.Address, prizeCosmicSignatureNftId []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	var prizeCosmicSignatureNftIdRule []interface{}
	for _, prizeCosmicSignatureNftIdItem := range prizeCosmicSignatureNftId {
		prizeCosmicSignatureNftIdRule = append(prizeCosmicSignatureNftIdRule, prizeCosmicSignatureNftIdItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "MainPrizeClaimed", roundNumRule, beneficiaryAddressRule, prizeCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2MainPrizeClaimed)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "MainPrizeClaimed", log); err != nil {
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

// ParseMainPrizeClaimed is a log parse operation binding the contract event 0x8c551ec2b6f186753e27f1cf46f84b57f4f83f721e8c1e6170ae512845ced591.
//
// Solidity: event MainPrizeClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, uint256 ethPrizeAmount, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId, uint256 timeoutTimeToWithdrawSecondaryPrizes)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseMainPrizeClaimed(log types.Log) (*CosmicSignatureGameV2MainPrizeClaimed, error) {
	event := new(CosmicSignatureGameV2MainPrizeClaimed)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "MainPrizeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2MainPrizeTimeIncrementInMicroSecondsChangedIterator is returned from FilterMainPrizeTimeIncrementInMicroSecondsChanged and is used to iterate over the raw logs and unpacked data for MainPrizeTimeIncrementInMicroSecondsChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2MainPrizeTimeIncrementInMicroSecondsChangedIterator struct {
	Event *CosmicSignatureGameV2MainPrizeTimeIncrementInMicroSecondsChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2MainPrizeTimeIncrementInMicroSecondsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2MainPrizeTimeIncrementInMicroSecondsChanged)
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
		it.Event = new(CosmicSignatureGameV2MainPrizeTimeIncrementInMicroSecondsChanged)
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
func (it *CosmicSignatureGameV2MainPrizeTimeIncrementInMicroSecondsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2MainPrizeTimeIncrementInMicroSecondsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2MainPrizeTimeIncrementInMicroSecondsChanged represents a MainPrizeTimeIncrementInMicroSecondsChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2MainPrizeTimeIncrementInMicroSecondsChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMainPrizeTimeIncrementInMicroSecondsChanged is a free log retrieval operation binding the contract event 0x07417920574ce0bdfe987af0575c8793cc73a29d7830760ad459d0e569b5b79b.
//
// Solidity: event MainPrizeTimeIncrementInMicroSecondsChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterMainPrizeTimeIncrementInMicroSecondsChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2MainPrizeTimeIncrementInMicroSecondsChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "MainPrizeTimeIncrementInMicroSecondsChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2MainPrizeTimeIncrementInMicroSecondsChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "MainPrizeTimeIncrementInMicroSecondsChanged", logs: logs, sub: sub}, nil
}

// WatchMainPrizeTimeIncrementInMicroSecondsChanged is a free log subscription operation binding the contract event 0x07417920574ce0bdfe987af0575c8793cc73a29d7830760ad459d0e569b5b79b.
//
// Solidity: event MainPrizeTimeIncrementInMicroSecondsChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchMainPrizeTimeIncrementInMicroSecondsChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2MainPrizeTimeIncrementInMicroSecondsChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "MainPrizeTimeIncrementInMicroSecondsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2MainPrizeTimeIncrementInMicroSecondsChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "MainPrizeTimeIncrementInMicroSecondsChanged", log); err != nil {
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

// ParseMainPrizeTimeIncrementInMicroSecondsChanged is a log parse operation binding the contract event 0x07417920574ce0bdfe987af0575c8793cc73a29d7830760ad459d0e569b5b79b.
//
// Solidity: event MainPrizeTimeIncrementInMicroSecondsChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseMainPrizeTimeIncrementInMicroSecondsChanged(log types.Log) (*CosmicSignatureGameV2MainPrizeTimeIncrementInMicroSecondsChanged, error) {
	event := new(CosmicSignatureGameV2MainPrizeTimeIncrementInMicroSecondsChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "MainPrizeTimeIncrementInMicroSecondsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2MainPrizeTimeIncrementIncreaseDivisorChangedIterator is returned from FilterMainPrizeTimeIncrementIncreaseDivisorChanged and is used to iterate over the raw logs and unpacked data for MainPrizeTimeIncrementIncreaseDivisorChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2MainPrizeTimeIncrementIncreaseDivisorChangedIterator struct {
	Event *CosmicSignatureGameV2MainPrizeTimeIncrementIncreaseDivisorChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2MainPrizeTimeIncrementIncreaseDivisorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2MainPrizeTimeIncrementIncreaseDivisorChanged)
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
		it.Event = new(CosmicSignatureGameV2MainPrizeTimeIncrementIncreaseDivisorChanged)
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
func (it *CosmicSignatureGameV2MainPrizeTimeIncrementIncreaseDivisorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2MainPrizeTimeIncrementIncreaseDivisorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2MainPrizeTimeIncrementIncreaseDivisorChanged represents a MainPrizeTimeIncrementIncreaseDivisorChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2MainPrizeTimeIncrementIncreaseDivisorChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMainPrizeTimeIncrementIncreaseDivisorChanged is a free log retrieval operation binding the contract event 0x4636d3e567b27988879babd22e50f49104ca65647933fc1623ff3d4d807438d2.
//
// Solidity: event MainPrizeTimeIncrementIncreaseDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterMainPrizeTimeIncrementIncreaseDivisorChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2MainPrizeTimeIncrementIncreaseDivisorChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "MainPrizeTimeIncrementIncreaseDivisorChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2MainPrizeTimeIncrementIncreaseDivisorChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "MainPrizeTimeIncrementIncreaseDivisorChanged", logs: logs, sub: sub}, nil
}

// WatchMainPrizeTimeIncrementIncreaseDivisorChanged is a free log subscription operation binding the contract event 0x4636d3e567b27988879babd22e50f49104ca65647933fc1623ff3d4d807438d2.
//
// Solidity: event MainPrizeTimeIncrementIncreaseDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchMainPrizeTimeIncrementIncreaseDivisorChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2MainPrizeTimeIncrementIncreaseDivisorChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "MainPrizeTimeIncrementIncreaseDivisorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2MainPrizeTimeIncrementIncreaseDivisorChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "MainPrizeTimeIncrementIncreaseDivisorChanged", log); err != nil {
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

// ParseMainPrizeTimeIncrementIncreaseDivisorChanged is a log parse operation binding the contract event 0x4636d3e567b27988879babd22e50f49104ca65647933fc1623ff3d4d807438d2.
//
// Solidity: event MainPrizeTimeIncrementIncreaseDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseMainPrizeTimeIncrementIncreaseDivisorChanged(log types.Log) (*CosmicSignatureGameV2MainPrizeTimeIncrementIncreaseDivisorChanged, error) {
	event := new(CosmicSignatureGameV2MainPrizeTimeIncrementIncreaseDivisorChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "MainPrizeTimeIncrementIncreaseDivisorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2MarketingWalletAddressChangedIterator is returned from FilterMarketingWalletAddressChanged and is used to iterate over the raw logs and unpacked data for MarketingWalletAddressChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2MarketingWalletAddressChangedIterator struct {
	Event *CosmicSignatureGameV2MarketingWalletAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2MarketingWalletAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2MarketingWalletAddressChanged)
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
		it.Event = new(CosmicSignatureGameV2MarketingWalletAddressChanged)
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
func (it *CosmicSignatureGameV2MarketingWalletAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2MarketingWalletAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2MarketingWalletAddressChanged represents a MarketingWalletAddressChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2MarketingWalletAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketingWalletAddressChanged is a free log retrieval operation binding the contract event 0x4d03942c29c20d1bccfe551e9d148c917c5a44fb558a4fc60270d8f76fb75f54.
//
// Solidity: event MarketingWalletAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterMarketingWalletAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*CosmicSignatureGameV2MarketingWalletAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "MarketingWalletAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2MarketingWalletAddressChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "MarketingWalletAddressChanged", logs: logs, sub: sub}, nil
}

// WatchMarketingWalletAddressChanged is a free log subscription operation binding the contract event 0x4d03942c29c20d1bccfe551e9d148c917c5a44fb558a4fc60270d8f76fb75f54.
//
// Solidity: event MarketingWalletAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchMarketingWalletAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2MarketingWalletAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "MarketingWalletAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2MarketingWalletAddressChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "MarketingWalletAddressChanged", log); err != nil {
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
// Solidity: event MarketingWalletAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseMarketingWalletAddressChanged(log types.Log) (*CosmicSignatureGameV2MarketingWalletAddressChanged, error) {
	event := new(CosmicSignatureGameV2MarketingWalletAddressChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "MarketingWalletAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2MarketingWalletCstContributionAmountChangedIterator is returned from FilterMarketingWalletCstContributionAmountChanged and is used to iterate over the raw logs and unpacked data for MarketingWalletCstContributionAmountChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2MarketingWalletCstContributionAmountChangedIterator struct {
	Event *CosmicSignatureGameV2MarketingWalletCstContributionAmountChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2MarketingWalletCstContributionAmountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2MarketingWalletCstContributionAmountChanged)
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
		it.Event = new(CosmicSignatureGameV2MarketingWalletCstContributionAmountChanged)
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
func (it *CosmicSignatureGameV2MarketingWalletCstContributionAmountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2MarketingWalletCstContributionAmountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2MarketingWalletCstContributionAmountChanged represents a MarketingWalletCstContributionAmountChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2MarketingWalletCstContributionAmountChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketingWalletCstContributionAmountChanged is a free log retrieval operation binding the contract event 0x2652e6657dd1ed89d7bdcb70f8827cc8449ac4536ebf604dbb2465cdad264020.
//
// Solidity: event MarketingWalletCstContributionAmountChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterMarketingWalletCstContributionAmountChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2MarketingWalletCstContributionAmountChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "MarketingWalletCstContributionAmountChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2MarketingWalletCstContributionAmountChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "MarketingWalletCstContributionAmountChanged", logs: logs, sub: sub}, nil
}

// WatchMarketingWalletCstContributionAmountChanged is a free log subscription operation binding the contract event 0x2652e6657dd1ed89d7bdcb70f8827cc8449ac4536ebf604dbb2465cdad264020.
//
// Solidity: event MarketingWalletCstContributionAmountChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchMarketingWalletCstContributionAmountChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2MarketingWalletCstContributionAmountChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "MarketingWalletCstContributionAmountChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2MarketingWalletCstContributionAmountChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "MarketingWalletCstContributionAmountChanged", log); err != nil {
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

// ParseMarketingWalletCstContributionAmountChanged is a log parse operation binding the contract event 0x2652e6657dd1ed89d7bdcb70f8827cc8449ac4536ebf604dbb2465cdad264020.
//
// Solidity: event MarketingWalletCstContributionAmountChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseMarketingWalletCstContributionAmountChanged(log types.Log) (*CosmicSignatureGameV2MarketingWalletCstContributionAmountChanged, error) {
	event := new(CosmicSignatureGameV2MarketingWalletCstContributionAmountChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "MarketingWalletCstContributionAmountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForBiddersChangedIterator is returned from FilterNumRaffleCosmicSignatureNftsForBiddersChanged and is used to iterate over the raw logs and unpacked data for NumRaffleCosmicSignatureNftsForBiddersChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForBiddersChangedIterator struct {
	Event *CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForBiddersChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForBiddersChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForBiddersChanged)
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
		it.Event = new(CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForBiddersChanged)
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
func (it *CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForBiddersChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForBiddersChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForBiddersChanged represents a NumRaffleCosmicSignatureNftsForBiddersChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForBiddersChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNumRaffleCosmicSignatureNftsForBiddersChanged is a free log retrieval operation binding the contract event 0x85d8bf21006916836edd67a5afeed2e891cf876a6c9cd9babf7f42f4b007c24f.
//
// Solidity: event NumRaffleCosmicSignatureNftsForBiddersChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterNumRaffleCosmicSignatureNftsForBiddersChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForBiddersChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "NumRaffleCosmicSignatureNftsForBiddersChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForBiddersChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "NumRaffleCosmicSignatureNftsForBiddersChanged", logs: logs, sub: sub}, nil
}

// WatchNumRaffleCosmicSignatureNftsForBiddersChanged is a free log subscription operation binding the contract event 0x85d8bf21006916836edd67a5afeed2e891cf876a6c9cd9babf7f42f4b007c24f.
//
// Solidity: event NumRaffleCosmicSignatureNftsForBiddersChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchNumRaffleCosmicSignatureNftsForBiddersChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForBiddersChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "NumRaffleCosmicSignatureNftsForBiddersChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForBiddersChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "NumRaffleCosmicSignatureNftsForBiddersChanged", log); err != nil {
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

// ParseNumRaffleCosmicSignatureNftsForBiddersChanged is a log parse operation binding the contract event 0x85d8bf21006916836edd67a5afeed2e891cf876a6c9cd9babf7f42f4b007c24f.
//
// Solidity: event NumRaffleCosmicSignatureNftsForBiddersChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseNumRaffleCosmicSignatureNftsForBiddersChanged(log types.Log) (*CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForBiddersChanged, error) {
	event := new(CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForBiddersChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "NumRaffleCosmicSignatureNftsForBiddersChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChangedIterator is returned from FilterNumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged and is used to iterate over the raw logs and unpacked data for NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChangedIterator struct {
	Event *CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged)
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
		it.Event = new(CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged)
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
func (it *CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged represents a NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged is a free log retrieval operation binding the contract event 0x3312247fcf207243294680d9103851c8709e19be3d46ee4b1bff5e12d1c5ca7b.
//
// Solidity: event NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterNumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged", logs: logs, sub: sub}, nil
}

// WatchNumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged is a free log subscription operation binding the contract event 0x3312247fcf207243294680d9103851c8709e19be3d46ee4b1bff5e12d1c5ca7b.
//
// Solidity: event NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchNumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged", log); err != nil {
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

// ParseNumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged is a log parse operation binding the contract event 0x3312247fcf207243294680d9103851c8709e19be3d46ee4b1bff5e12d1c5ca7b.
//
// Solidity: event NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseNumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged(log types.Log) (*CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged, error) {
	event := new(CosmicSignatureGameV2NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2NumRaffleEthPrizesForBiddersChangedIterator is returned from FilterNumRaffleEthPrizesForBiddersChanged and is used to iterate over the raw logs and unpacked data for NumRaffleEthPrizesForBiddersChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2NumRaffleEthPrizesForBiddersChangedIterator struct {
	Event *CosmicSignatureGameV2NumRaffleEthPrizesForBiddersChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2NumRaffleEthPrizesForBiddersChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2NumRaffleEthPrizesForBiddersChanged)
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
		it.Event = new(CosmicSignatureGameV2NumRaffleEthPrizesForBiddersChanged)
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
func (it *CosmicSignatureGameV2NumRaffleEthPrizesForBiddersChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2NumRaffleEthPrizesForBiddersChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2NumRaffleEthPrizesForBiddersChanged represents a NumRaffleEthPrizesForBiddersChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2NumRaffleEthPrizesForBiddersChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNumRaffleEthPrizesForBiddersChanged is a free log retrieval operation binding the contract event 0x4787028773c8f14bc3b4bc41f43a02329ae41105823287201a34782c530d35fd.
//
// Solidity: event NumRaffleEthPrizesForBiddersChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterNumRaffleEthPrizesForBiddersChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2NumRaffleEthPrizesForBiddersChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "NumRaffleEthPrizesForBiddersChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2NumRaffleEthPrizesForBiddersChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "NumRaffleEthPrizesForBiddersChanged", logs: logs, sub: sub}, nil
}

// WatchNumRaffleEthPrizesForBiddersChanged is a free log subscription operation binding the contract event 0x4787028773c8f14bc3b4bc41f43a02329ae41105823287201a34782c530d35fd.
//
// Solidity: event NumRaffleEthPrizesForBiddersChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchNumRaffleEthPrizesForBiddersChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2NumRaffleEthPrizesForBiddersChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "NumRaffleEthPrizesForBiddersChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2NumRaffleEthPrizesForBiddersChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "NumRaffleEthPrizesForBiddersChanged", log); err != nil {
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

// ParseNumRaffleEthPrizesForBiddersChanged is a log parse operation binding the contract event 0x4787028773c8f14bc3b4bc41f43a02329ae41105823287201a34782c530d35fd.
//
// Solidity: event NumRaffleEthPrizesForBiddersChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseNumRaffleEthPrizesForBiddersChanged(log types.Log) (*CosmicSignatureGameV2NumRaffleEthPrizesForBiddersChanged, error) {
	event := new(CosmicSignatureGameV2NumRaffleEthPrizesForBiddersChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "NumRaffleEthPrizesForBiddersChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2OwnershipTransferredIterator struct {
	Event *CosmicSignatureGameV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2OwnershipTransferred)
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
		it.Event = new(CosmicSignatureGameV2OwnershipTransferred)
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
func (it *CosmicSignatureGameV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2OwnershipTransferred represents a OwnershipTransferred event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CosmicSignatureGameV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2OwnershipTransferredIterator{contract: _CosmicSignatureGameV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2OwnershipTransferred)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseOwnershipTransferred(log types.Log) (*CosmicSignatureGameV2OwnershipTransferred, error) {
	event := new(CosmicSignatureGameV2OwnershipTransferred)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2PrizesWalletAddressChangedIterator is returned from FilterPrizesWalletAddressChanged and is used to iterate over the raw logs and unpacked data for PrizesWalletAddressChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2PrizesWalletAddressChangedIterator struct {
	Event *CosmicSignatureGameV2PrizesWalletAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2PrizesWalletAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2PrizesWalletAddressChanged)
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
		it.Event = new(CosmicSignatureGameV2PrizesWalletAddressChanged)
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
func (it *CosmicSignatureGameV2PrizesWalletAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2PrizesWalletAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2PrizesWalletAddressChanged represents a PrizesWalletAddressChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2PrizesWalletAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPrizesWalletAddressChanged is a free log retrieval operation binding the contract event 0xb4cecfe1346c94da27291cf5a02969d5fe0b5c36eca49b04fcd60841d28c5e13.
//
// Solidity: event PrizesWalletAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterPrizesWalletAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*CosmicSignatureGameV2PrizesWalletAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "PrizesWalletAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2PrizesWalletAddressChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "PrizesWalletAddressChanged", logs: logs, sub: sub}, nil
}

// WatchPrizesWalletAddressChanged is a free log subscription operation binding the contract event 0xb4cecfe1346c94da27291cf5a02969d5fe0b5c36eca49b04fcd60841d28c5e13.
//
// Solidity: event PrizesWalletAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchPrizesWalletAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2PrizesWalletAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "PrizesWalletAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2PrizesWalletAddressChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "PrizesWalletAddressChanged", log); err != nil {
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

// ParsePrizesWalletAddressChanged is a log parse operation binding the contract event 0xb4cecfe1346c94da27291cf5a02969d5fe0b5c36eca49b04fcd60841d28c5e13.
//
// Solidity: event PrizesWalletAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParsePrizesWalletAddressChanged(log types.Log) (*CosmicSignatureGameV2PrizesWalletAddressChanged, error) {
	event := new(CosmicSignatureGameV2PrizesWalletAddressChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "PrizesWalletAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2RaffleTotalEthPrizeAmountForBiddersPercentageChangedIterator is returned from FilterRaffleTotalEthPrizeAmountForBiddersPercentageChanged and is used to iterate over the raw logs and unpacked data for RaffleTotalEthPrizeAmountForBiddersPercentageChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2RaffleTotalEthPrizeAmountForBiddersPercentageChangedIterator struct {
	Event *CosmicSignatureGameV2RaffleTotalEthPrizeAmountForBiddersPercentageChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2RaffleTotalEthPrizeAmountForBiddersPercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2RaffleTotalEthPrizeAmountForBiddersPercentageChanged)
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
		it.Event = new(CosmicSignatureGameV2RaffleTotalEthPrizeAmountForBiddersPercentageChanged)
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
func (it *CosmicSignatureGameV2RaffleTotalEthPrizeAmountForBiddersPercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2RaffleTotalEthPrizeAmountForBiddersPercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2RaffleTotalEthPrizeAmountForBiddersPercentageChanged represents a RaffleTotalEthPrizeAmountForBiddersPercentageChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2RaffleTotalEthPrizeAmountForBiddersPercentageChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRaffleTotalEthPrizeAmountForBiddersPercentageChanged is a free log retrieval operation binding the contract event 0xbfcd8fb930a57c1598c9760db19c84ec766546c3c9a8565611df8302482bfb17.
//
// Solidity: event RaffleTotalEthPrizeAmountForBiddersPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterRaffleTotalEthPrizeAmountForBiddersPercentageChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2RaffleTotalEthPrizeAmountForBiddersPercentageChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "RaffleTotalEthPrizeAmountForBiddersPercentageChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2RaffleTotalEthPrizeAmountForBiddersPercentageChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "RaffleTotalEthPrizeAmountForBiddersPercentageChanged", logs: logs, sub: sub}, nil
}

// WatchRaffleTotalEthPrizeAmountForBiddersPercentageChanged is a free log subscription operation binding the contract event 0xbfcd8fb930a57c1598c9760db19c84ec766546c3c9a8565611df8302482bfb17.
//
// Solidity: event RaffleTotalEthPrizeAmountForBiddersPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchRaffleTotalEthPrizeAmountForBiddersPercentageChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2RaffleTotalEthPrizeAmountForBiddersPercentageChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "RaffleTotalEthPrizeAmountForBiddersPercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2RaffleTotalEthPrizeAmountForBiddersPercentageChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "RaffleTotalEthPrizeAmountForBiddersPercentageChanged", log); err != nil {
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

// ParseRaffleTotalEthPrizeAmountForBiddersPercentageChanged is a log parse operation binding the contract event 0xbfcd8fb930a57c1598c9760db19c84ec766546c3c9a8565611df8302482bfb17.
//
// Solidity: event RaffleTotalEthPrizeAmountForBiddersPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseRaffleTotalEthPrizeAmountForBiddersPercentageChanged(log types.Log) (*CosmicSignatureGameV2RaffleTotalEthPrizeAmountForBiddersPercentageChanged, error) {
	event := new(CosmicSignatureGameV2RaffleTotalEthPrizeAmountForBiddersPercentageChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "RaffleTotalEthPrizeAmountForBiddersPercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2RaffleWinnerBidderEthPrizeAllocatedIterator is returned from FilterRaffleWinnerBidderEthPrizeAllocated and is used to iterate over the raw logs and unpacked data for RaffleWinnerBidderEthPrizeAllocated events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2RaffleWinnerBidderEthPrizeAllocatedIterator struct {
	Event *CosmicSignatureGameV2RaffleWinnerBidderEthPrizeAllocated // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2RaffleWinnerBidderEthPrizeAllocatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2RaffleWinnerBidderEthPrizeAllocated)
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
		it.Event = new(CosmicSignatureGameV2RaffleWinnerBidderEthPrizeAllocated)
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
func (it *CosmicSignatureGameV2RaffleWinnerBidderEthPrizeAllocatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2RaffleWinnerBidderEthPrizeAllocatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2RaffleWinnerBidderEthPrizeAllocated represents a RaffleWinnerBidderEthPrizeAllocated event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2RaffleWinnerBidderEthPrizeAllocated struct {
	RoundNum       *big.Int
	WinnerIndex    *big.Int
	WinnerAddress  common.Address
	EthPrizeAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRaffleWinnerBidderEthPrizeAllocated is a free log retrieval operation binding the contract event 0x9c62e2cb8cbd10bf3b8a5760977d719fc3475fa67cb7ef9d2e1701f275e885c4.
//
// Solidity: event RaffleWinnerBidderEthPrizeAllocated(uint256 indexed roundNum, uint256 winnerIndex, address indexed winnerAddress, uint256 ethPrizeAmount)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterRaffleWinnerBidderEthPrizeAllocated(opts *bind.FilterOpts, roundNum []*big.Int, winnerAddress []common.Address) (*CosmicSignatureGameV2RaffleWinnerBidderEthPrizeAllocatedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	var winnerAddressRule []interface{}
	for _, winnerAddressItem := range winnerAddress {
		winnerAddressRule = append(winnerAddressRule, winnerAddressItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "RaffleWinnerBidderEthPrizeAllocated", roundNumRule, winnerAddressRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2RaffleWinnerBidderEthPrizeAllocatedIterator{contract: _CosmicSignatureGameV2.contract, event: "RaffleWinnerBidderEthPrizeAllocated", logs: logs, sub: sub}, nil
}

// WatchRaffleWinnerBidderEthPrizeAllocated is a free log subscription operation binding the contract event 0x9c62e2cb8cbd10bf3b8a5760977d719fc3475fa67cb7ef9d2e1701f275e885c4.
//
// Solidity: event RaffleWinnerBidderEthPrizeAllocated(uint256 indexed roundNum, uint256 winnerIndex, address indexed winnerAddress, uint256 ethPrizeAmount)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchRaffleWinnerBidderEthPrizeAllocated(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2RaffleWinnerBidderEthPrizeAllocated, roundNum []*big.Int, winnerAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	var winnerAddressRule []interface{}
	for _, winnerAddressItem := range winnerAddress {
		winnerAddressRule = append(winnerAddressRule, winnerAddressItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "RaffleWinnerBidderEthPrizeAllocated", roundNumRule, winnerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2RaffleWinnerBidderEthPrizeAllocated)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "RaffleWinnerBidderEthPrizeAllocated", log); err != nil {
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

// ParseRaffleWinnerBidderEthPrizeAllocated is a log parse operation binding the contract event 0x9c62e2cb8cbd10bf3b8a5760977d719fc3475fa67cb7ef9d2e1701f275e885c4.
//
// Solidity: event RaffleWinnerBidderEthPrizeAllocated(uint256 indexed roundNum, uint256 winnerIndex, address indexed winnerAddress, uint256 ethPrizeAmount)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseRaffleWinnerBidderEthPrizeAllocated(log types.Log) (*CosmicSignatureGameV2RaffleWinnerBidderEthPrizeAllocated, error) {
	event := new(CosmicSignatureGameV2RaffleWinnerBidderEthPrizeAllocated)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "RaffleWinnerBidderEthPrizeAllocated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2RaffleWinnerPrizePaidIterator is returned from FilterRaffleWinnerPrizePaid and is used to iterate over the raw logs and unpacked data for RaffleWinnerPrizePaid events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2RaffleWinnerPrizePaidIterator struct {
	Event *CosmicSignatureGameV2RaffleWinnerPrizePaid // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2RaffleWinnerPrizePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2RaffleWinnerPrizePaid)
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
		it.Event = new(CosmicSignatureGameV2RaffleWinnerPrizePaid)
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
func (it *CosmicSignatureGameV2RaffleWinnerPrizePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2RaffleWinnerPrizePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2RaffleWinnerPrizePaid represents a RaffleWinnerPrizePaid event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2RaffleWinnerPrizePaid struct {
	RoundNum                    *big.Int
	WinnerIsRandomWalkNftStaker bool
	WinnerIndex                 *big.Int
	WinnerAddress               common.Address
	CstPrizeAmount              *big.Int
	PrizeCosmicSignatureNftId   *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterRaffleWinnerPrizePaid is a free log retrieval operation binding the contract event 0x27c21fe4cea1a3367aa491829dd4dd824296c00910626150464cba8ea5ebb3f4.
//
// Solidity: event RaffleWinnerPrizePaid(uint256 indexed roundNum, bool winnerIsRandomWalkNftStaker, uint256 winnerIndex, address indexed winnerAddress, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterRaffleWinnerPrizePaid(opts *bind.FilterOpts, roundNum []*big.Int, winnerAddress []common.Address, prizeCosmicSignatureNftId []*big.Int) (*CosmicSignatureGameV2RaffleWinnerPrizePaidIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	var winnerAddressRule []interface{}
	for _, winnerAddressItem := range winnerAddress {
		winnerAddressRule = append(winnerAddressRule, winnerAddressItem)
	}

	var prizeCosmicSignatureNftIdRule []interface{}
	for _, prizeCosmicSignatureNftIdItem := range prizeCosmicSignatureNftId {
		prizeCosmicSignatureNftIdRule = append(prizeCosmicSignatureNftIdRule, prizeCosmicSignatureNftIdItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "RaffleWinnerPrizePaid", roundNumRule, winnerAddressRule, prizeCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2RaffleWinnerPrizePaidIterator{contract: _CosmicSignatureGameV2.contract, event: "RaffleWinnerPrizePaid", logs: logs, sub: sub}, nil
}

// WatchRaffleWinnerPrizePaid is a free log subscription operation binding the contract event 0x27c21fe4cea1a3367aa491829dd4dd824296c00910626150464cba8ea5ebb3f4.
//
// Solidity: event RaffleWinnerPrizePaid(uint256 indexed roundNum, bool winnerIsRandomWalkNftStaker, uint256 winnerIndex, address indexed winnerAddress, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchRaffleWinnerPrizePaid(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2RaffleWinnerPrizePaid, roundNum []*big.Int, winnerAddress []common.Address, prizeCosmicSignatureNftId []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	var winnerAddressRule []interface{}
	for _, winnerAddressItem := range winnerAddress {
		winnerAddressRule = append(winnerAddressRule, winnerAddressItem)
	}

	var prizeCosmicSignatureNftIdRule []interface{}
	for _, prizeCosmicSignatureNftIdItem := range prizeCosmicSignatureNftId {
		prizeCosmicSignatureNftIdRule = append(prizeCosmicSignatureNftIdRule, prizeCosmicSignatureNftIdItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "RaffleWinnerPrizePaid", roundNumRule, winnerAddressRule, prizeCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2RaffleWinnerPrizePaid)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "RaffleWinnerPrizePaid", log); err != nil {
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

// ParseRaffleWinnerPrizePaid is a log parse operation binding the contract event 0x27c21fe4cea1a3367aa491829dd4dd824296c00910626150464cba8ea5ebb3f4.
//
// Solidity: event RaffleWinnerPrizePaid(uint256 indexed roundNum, bool winnerIsRandomWalkNftStaker, uint256 winnerIndex, address indexed winnerAddress, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseRaffleWinnerPrizePaid(log types.Log) (*CosmicSignatureGameV2RaffleWinnerPrizePaid, error) {
	event := new(CosmicSignatureGameV2RaffleWinnerPrizePaid)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "RaffleWinnerPrizePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2RandomWalkNftAddressChangedIterator is returned from FilterRandomWalkNftAddressChanged and is used to iterate over the raw logs and unpacked data for RandomWalkNftAddressChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2RandomWalkNftAddressChangedIterator struct {
	Event *CosmicSignatureGameV2RandomWalkNftAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2RandomWalkNftAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2RandomWalkNftAddressChanged)
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
		it.Event = new(CosmicSignatureGameV2RandomWalkNftAddressChanged)
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
func (it *CosmicSignatureGameV2RandomWalkNftAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2RandomWalkNftAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2RandomWalkNftAddressChanged represents a RandomWalkNftAddressChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2RandomWalkNftAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRandomWalkNftAddressChanged is a free log retrieval operation binding the contract event 0xdab38e33e6e11cbb8b085bba9d7426d5e5af01bcc46d2c5957645e4d09e8c49c.
//
// Solidity: event RandomWalkNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterRandomWalkNftAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*CosmicSignatureGameV2RandomWalkNftAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "RandomWalkNftAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2RandomWalkNftAddressChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "RandomWalkNftAddressChanged", logs: logs, sub: sub}, nil
}

// WatchRandomWalkNftAddressChanged is a free log subscription operation binding the contract event 0xdab38e33e6e11cbb8b085bba9d7426d5e5af01bcc46d2c5957645e4d09e8c49c.
//
// Solidity: event RandomWalkNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchRandomWalkNftAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2RandomWalkNftAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "RandomWalkNftAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2RandomWalkNftAddressChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "RandomWalkNftAddressChanged", log); err != nil {
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

// ParseRandomWalkNftAddressChanged is a log parse operation binding the contract event 0xdab38e33e6e11cbb8b085bba9d7426d5e5af01bcc46d2c5957645e4d09e8c49c.
//
// Solidity: event RandomWalkNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseRandomWalkNftAddressChanged(log types.Log) (*CosmicSignatureGameV2RandomWalkNftAddressChanged, error) {
	event := new(CosmicSignatureGameV2RandomWalkNftAddressChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "RandomWalkNftAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2RoundActivationTimeChangedIterator is returned from FilterRoundActivationTimeChanged and is used to iterate over the raw logs and unpacked data for RoundActivationTimeChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2RoundActivationTimeChangedIterator struct {
	Event *CosmicSignatureGameV2RoundActivationTimeChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2RoundActivationTimeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2RoundActivationTimeChanged)
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
		it.Event = new(CosmicSignatureGameV2RoundActivationTimeChanged)
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
func (it *CosmicSignatureGameV2RoundActivationTimeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2RoundActivationTimeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2RoundActivationTimeChanged represents a RoundActivationTimeChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2RoundActivationTimeChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoundActivationTimeChanged is a free log retrieval operation binding the contract event 0x9a2159c1f277ddd727551baedc6a6c4cba77cc5219c8563ee3b15fb67548d89b.
//
// Solidity: event RoundActivationTimeChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterRoundActivationTimeChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2RoundActivationTimeChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "RoundActivationTimeChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2RoundActivationTimeChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "RoundActivationTimeChanged", logs: logs, sub: sub}, nil
}

// WatchRoundActivationTimeChanged is a free log subscription operation binding the contract event 0x9a2159c1f277ddd727551baedc6a6c4cba77cc5219c8563ee3b15fb67548d89b.
//
// Solidity: event RoundActivationTimeChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchRoundActivationTimeChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2RoundActivationTimeChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "RoundActivationTimeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2RoundActivationTimeChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "RoundActivationTimeChanged", log); err != nil {
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

// ParseRoundActivationTimeChanged is a log parse operation binding the contract event 0x9a2159c1f277ddd727551baedc6a6c4cba77cc5219c8563ee3b15fb67548d89b.
//
// Solidity: event RoundActivationTimeChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseRoundActivationTimeChanged(log types.Log) (*CosmicSignatureGameV2RoundActivationTimeChanged, error) {
	event := new(CosmicSignatureGameV2RoundActivationTimeChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "RoundActivationTimeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2StakingWalletCosmicSignatureNftAddressChangedIterator is returned from FilterStakingWalletCosmicSignatureNftAddressChanged and is used to iterate over the raw logs and unpacked data for StakingWalletCosmicSignatureNftAddressChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2StakingWalletCosmicSignatureNftAddressChangedIterator struct {
	Event *CosmicSignatureGameV2StakingWalletCosmicSignatureNftAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2StakingWalletCosmicSignatureNftAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2StakingWalletCosmicSignatureNftAddressChanged)
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
		it.Event = new(CosmicSignatureGameV2StakingWalletCosmicSignatureNftAddressChanged)
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
func (it *CosmicSignatureGameV2StakingWalletCosmicSignatureNftAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2StakingWalletCosmicSignatureNftAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2StakingWalletCosmicSignatureNftAddressChanged represents a StakingWalletCosmicSignatureNftAddressChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2StakingWalletCosmicSignatureNftAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakingWalletCosmicSignatureNftAddressChanged is a free log retrieval operation binding the contract event 0x4da1815cd654922275d14d2335fd9a0dd0aa6a0d0ff87fb4cc872ebe9704596f.
//
// Solidity: event StakingWalletCosmicSignatureNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterStakingWalletCosmicSignatureNftAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*CosmicSignatureGameV2StakingWalletCosmicSignatureNftAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "StakingWalletCosmicSignatureNftAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2StakingWalletCosmicSignatureNftAddressChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "StakingWalletCosmicSignatureNftAddressChanged", logs: logs, sub: sub}, nil
}

// WatchStakingWalletCosmicSignatureNftAddressChanged is a free log subscription operation binding the contract event 0x4da1815cd654922275d14d2335fd9a0dd0aa6a0d0ff87fb4cc872ebe9704596f.
//
// Solidity: event StakingWalletCosmicSignatureNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchStakingWalletCosmicSignatureNftAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2StakingWalletCosmicSignatureNftAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "StakingWalletCosmicSignatureNftAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2StakingWalletCosmicSignatureNftAddressChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "StakingWalletCosmicSignatureNftAddressChanged", log); err != nil {
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

// ParseStakingWalletCosmicSignatureNftAddressChanged is a log parse operation binding the contract event 0x4da1815cd654922275d14d2335fd9a0dd0aa6a0d0ff87fb4cc872ebe9704596f.
//
// Solidity: event StakingWalletCosmicSignatureNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseStakingWalletCosmicSignatureNftAddressChanged(log types.Log) (*CosmicSignatureGameV2StakingWalletCosmicSignatureNftAddressChanged, error) {
	event := new(CosmicSignatureGameV2StakingWalletCosmicSignatureNftAddressChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "StakingWalletCosmicSignatureNftAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2StakingWalletRandomWalkNftAddressChangedIterator is returned from FilterStakingWalletRandomWalkNftAddressChanged and is used to iterate over the raw logs and unpacked data for StakingWalletRandomWalkNftAddressChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2StakingWalletRandomWalkNftAddressChangedIterator struct {
	Event *CosmicSignatureGameV2StakingWalletRandomWalkNftAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2StakingWalletRandomWalkNftAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2StakingWalletRandomWalkNftAddressChanged)
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
		it.Event = new(CosmicSignatureGameV2StakingWalletRandomWalkNftAddressChanged)
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
func (it *CosmicSignatureGameV2StakingWalletRandomWalkNftAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2StakingWalletRandomWalkNftAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2StakingWalletRandomWalkNftAddressChanged represents a StakingWalletRandomWalkNftAddressChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2StakingWalletRandomWalkNftAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakingWalletRandomWalkNftAddressChanged is a free log retrieval operation binding the contract event 0xbf6e296f85d08cc1ab124aed644bf4b19e4a726a7aea53e3784ab1341738a040.
//
// Solidity: event StakingWalletRandomWalkNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterStakingWalletRandomWalkNftAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*CosmicSignatureGameV2StakingWalletRandomWalkNftAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "StakingWalletRandomWalkNftAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2StakingWalletRandomWalkNftAddressChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "StakingWalletRandomWalkNftAddressChanged", logs: logs, sub: sub}, nil
}

// WatchStakingWalletRandomWalkNftAddressChanged is a free log subscription operation binding the contract event 0xbf6e296f85d08cc1ab124aed644bf4b19e4a726a7aea53e3784ab1341738a040.
//
// Solidity: event StakingWalletRandomWalkNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchStakingWalletRandomWalkNftAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2StakingWalletRandomWalkNftAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "StakingWalletRandomWalkNftAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2StakingWalletRandomWalkNftAddressChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "StakingWalletRandomWalkNftAddressChanged", log); err != nil {
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

// ParseStakingWalletRandomWalkNftAddressChanged is a log parse operation binding the contract event 0xbf6e296f85d08cc1ab124aed644bf4b19e4a726a7aea53e3784ab1341738a040.
//
// Solidity: event StakingWalletRandomWalkNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseStakingWalletRandomWalkNftAddressChanged(log types.Log) (*CosmicSignatureGameV2StakingWalletRandomWalkNftAddressChanged, error) {
	event := new(CosmicSignatureGameV2StakingWalletRandomWalkNftAddressChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "StakingWalletRandomWalkNftAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2TimeoutDurationToClaimMainPrizeChangedIterator is returned from FilterTimeoutDurationToClaimMainPrizeChanged and is used to iterate over the raw logs and unpacked data for TimeoutDurationToClaimMainPrizeChanged events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2TimeoutDurationToClaimMainPrizeChangedIterator struct {
	Event *CosmicSignatureGameV2TimeoutDurationToClaimMainPrizeChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2TimeoutDurationToClaimMainPrizeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2TimeoutDurationToClaimMainPrizeChanged)
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
		it.Event = new(CosmicSignatureGameV2TimeoutDurationToClaimMainPrizeChanged)
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
func (it *CosmicSignatureGameV2TimeoutDurationToClaimMainPrizeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2TimeoutDurationToClaimMainPrizeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2TimeoutDurationToClaimMainPrizeChanged represents a TimeoutDurationToClaimMainPrizeChanged event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2TimeoutDurationToClaimMainPrizeChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTimeoutDurationToClaimMainPrizeChanged is a free log retrieval operation binding the contract event 0x37a332914fac995349420c0419b4423a19dcb762017f691442a0782ce4bf417a.
//
// Solidity: event TimeoutDurationToClaimMainPrizeChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterTimeoutDurationToClaimMainPrizeChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV2TimeoutDurationToClaimMainPrizeChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "TimeoutDurationToClaimMainPrizeChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2TimeoutDurationToClaimMainPrizeChangedIterator{contract: _CosmicSignatureGameV2.contract, event: "TimeoutDurationToClaimMainPrizeChanged", logs: logs, sub: sub}, nil
}

// WatchTimeoutDurationToClaimMainPrizeChanged is a free log subscription operation binding the contract event 0x37a332914fac995349420c0419b4423a19dcb762017f691442a0782ce4bf417a.
//
// Solidity: event TimeoutDurationToClaimMainPrizeChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchTimeoutDurationToClaimMainPrizeChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2TimeoutDurationToClaimMainPrizeChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "TimeoutDurationToClaimMainPrizeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2TimeoutDurationToClaimMainPrizeChanged)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "TimeoutDurationToClaimMainPrizeChanged", log); err != nil {
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

// ParseTimeoutDurationToClaimMainPrizeChanged is a log parse operation binding the contract event 0x37a332914fac995349420c0419b4423a19dcb762017f691442a0782ce4bf417a.
//
// Solidity: event TimeoutDurationToClaimMainPrizeChanged(uint256 newValue)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseTimeoutDurationToClaimMainPrizeChanged(log types.Log) (*CosmicSignatureGameV2TimeoutDurationToClaimMainPrizeChanged, error) {
	event := new(CosmicSignatureGameV2TimeoutDurationToClaimMainPrizeChanged)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "TimeoutDurationToClaimMainPrizeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV2UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2UpgradedIterator struct {
	Event *CosmicSignatureGameV2Upgraded // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV2UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV2Upgraded)
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
		it.Event = new(CosmicSignatureGameV2Upgraded)
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
func (it *CosmicSignatureGameV2UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV2UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV2Upgraded represents a Upgraded event raised by the CosmicSignatureGameV2 contract.
type CosmicSignatureGameV2Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CosmicSignatureGameV2UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV2UpgradedIterator{contract: _CosmicSignatureGameV2.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV2Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CosmicSignatureGameV2.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV2Upgraded)
				if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CosmicSignatureGameV2 *CosmicSignatureGameV2Filterer) ParseUpgraded(log types.Log) (*CosmicSignatureGameV2Upgraded, error) {
	event := new(CosmicSignatureGameV2Upgraded)
	if err := _CosmicSignatureGameV2.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
