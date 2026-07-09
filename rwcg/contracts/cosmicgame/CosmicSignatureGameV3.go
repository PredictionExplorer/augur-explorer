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

// CosmicSignatureGameV3MetaData contains all meta data concerning the CosmicSignatureGameV3 contract.
var CosmicSignatureGameV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit\",\"type\":\"uint256\"}],\"name\":\"BidCstRewardAmountMinLimitNotReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"BidHasBeenPlacedInCurrentRound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"}],\"name\":\"CallerIsNotNftOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientReceivedBidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"InvalidOperationInCurrentState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"lastBidderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"durationUntilOperationIsPermitted\",\"type\":\"uint256\"}],\"name\":\"MainPrizeClaimDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"mainPrizeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"MainPrizeEarlyClaim\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"NoBidsPlacedInCurrentRound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"roundActivationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"RoundIsActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"roundActivationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"RoundIsInactive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"messageLength\",\"type\":\"uint256\"}],\"name\":\"TooLongBidMessage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"randomWalkNftId\",\"type\":\"uint256\"}],\"name\":\"UsedRandomWalkNft\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"WrongBidType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ArbitrumError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"BidCstRewardAmountMultiplierChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"BidCstRewardAmountPerMinuteChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"BidMessageLengthMaxLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lastBidderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"paidEthPrice\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"paidCstPrice\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"randomWalkNftId\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstDutchAuctionDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mainPrizeTime\",\"type\":\"uint256\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CharityAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CharityEthDonationAmountPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"ChronoWarriorEthPrizeAmountPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chronoWarriorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeCosmicSignatureNftId\",\"type\":\"uint256\"}],\"name\":\"ChronoWarriorPrizePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICosmicSignatureNft\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CosmicSignatureNftAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICosmicSignatureToken\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CosmicSignatureTokenAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CstDutchAuctionBeginningBidPriceMinLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CstDutchAuctionDurationChangeDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CstDutchAuctionDurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CstPrizeAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"DelayDurationBeforeRoundActivationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"enduranceChampionAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeCosmicSignatureNftId\",\"type\":\"uint256\"}],\"name\":\"EnduranceChampionPrizePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"EthBidPriceIncreaseDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"EthBidRefundAmountInGasToSwallowMaxLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthDonated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ethDonationWithInfoRecordIndex\",\"type\":\"uint256\"}],\"name\":\"EthDonatedWithInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"EthDutchAuctionDurationDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"EthDutchAuctionEndingBidPriceDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"FirstBidPlacedInRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"charityAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsTransferredToCharity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"InitialDurationUntilMainPrizeDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lastCstBidderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeCosmicSignatureNftId\",\"type\":\"uint256\"}],\"name\":\"LastCstBidderPrizePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MainEthPrizeAmountPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeFirstCosmicSignatureNftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prizeNumCosmicSignatureNfts\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeoutTimeToWithdrawSecondaryPrizes\",\"type\":\"uint256\"}],\"name\":\"MainPrizeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MainPrizeNumCosmicSignatureNftsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MainPrizeTimeIncrementInMicroSecondsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MainPrizeTimeIncrementIncreaseDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"MarketingWalletAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MarketingWalletCstContributionAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"NumRaffleCosmicSignatureNftsForBiddersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"NumRaffleEthPrizesForBiddersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIPrizesWallet\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"PrizesWalletAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RaffleTotalEthPrizeAmountForBiddersPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winnerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethPrizeAmount\",\"type\":\"uint256\"}],\"name\":\"RaffleWinnerBidderEthPrizeAllocated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"winnerIsRandomWalkNftStaker\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winnerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeCosmicSignatureNftId\",\"type\":\"uint256\"}],\"name\":\"RaffleWinnerPrizePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIRandomWalkNFT\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"RandomWalkNftAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RoundActivationTimeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RoundLateBidDurationDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RoundLateBidPricePremiumAmountBaseMultiplierChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RoundLateBidPricePremiumAmountExponentChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIStakingWalletCosmicSignatureNft\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"StakingWalletCosmicSignatureNftAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIStakingWalletRandomWalkNft\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"StakingWalletRandomWalkNftAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"TimeoutDurationToClaimMainPrizeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bidCstRewardAmountMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bidCstRewardAmountPerMinute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bidMessageLengthMaxLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceMaxLimit_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"}],\"name\":\"bidWithCst\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceMaxLimit_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"bidWithCstAndDonateNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceMaxLimit_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"bidWithCstAndDonateToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"randomWalkNftId_\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"}],\"name\":\"bidWithEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"randomWalkNftId_\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"bidWithEthAndDonateNft\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"randomWalkNftId_\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"bidWithEthAndDonateToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"}],\"name\":\"bidderAddresses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"numItems\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidderAddress\",\"type\":\"address\"}],\"name\":\"biddersInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalSpentEthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSpentCstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBidTimeStamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charityAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charityEthDonationAmountPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chronoWarriorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chronoWarriorDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chronoWarriorEthPrizeAmountPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimMainPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cosmicSignatureNftStakingTotalEthRewardAmountPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionBeginningBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionBeginningBidPriceMinLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionBeginningTimeStamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionDurationChangeDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstPrizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayDurationBeforeRoundActivation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donateEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"data_\",\"type\":\"string\"}],\"name\":\"donateEthWithInfo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enduranceChampionAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enduranceChampionDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enduranceChampionStartTimeStamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethBidPriceIncreaseDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethBidRefundAmountInGasToSwallowMaxLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ethDonationWithInfoRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethDutchAuctionBeginningBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethDutchAuctionDurationDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethDutchAuctionEndingBidPriceDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBidCstRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"currentTimeOffset_\",\"type\":\"int256\"}],\"name\":\"getBidCstRewardAmountAdvanced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidIndex_\",\"type\":\"uint256\"}],\"name\":\"getBidderAddressAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidderAddress_\",\"type\":\"address\"}],\"name\":\"getBidderTotalSpentAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCharityEthDonationAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChronoWarriorEthPrizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCosmicSignatureNftStakingTotalEthRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCstDutchAuctionDurations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDurationElapsedSinceRoundActivation\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDurationUntilMainPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDurationUntilMainPrizeRaw\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDurationUntilRoundActivation\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEthDutchAuctionDurations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethBidPrice_\",\"type\":\"uint256\"}],\"name\":\"getEthPlusRandomWalkNftBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitialDurationUntilMainPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMainEthPrizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMainPrizeTimeIncrement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextCstBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"currentTimeOffset_\",\"type\":\"int256\"}],\"name\":\"getNextCstBidPriceAdvanced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextEthBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"currentTimeOffset_\",\"type\":\"int256\"}],\"name\":\"getNextEthBidPriceAdvanced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRaffleTotalEthPrizeAmountForBidders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoundLateBidDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"}],\"name\":\"getTotalNumBids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"halveEthDutchAuctionEndingBidPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialDurationUntilMainPrizeDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBidderAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCstBidderAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainEthPrizeAmountPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainPrizeNumCosmicSignatureNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainPrizeTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainPrizeTimeIncrementInMicroSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainPrizeTimeIncrementIncreaseDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketingWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketingWalletCstContributionAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextEthBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRoundFirstCstDutchAuctionBeginningBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractCosmicSignatureNft\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numEthDonationWithInfoRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleCosmicSignatureNftsForBidders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleCosmicSignatureNftsForRandomWalkNftStakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleEthPrizesForBidders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prevEnduranceChampionDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizesWallet\",\"outputs\":[{\"internalType\":\"contractPrizesWallet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raffleTotalEthPrizeAmountForBiddersPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomWalkNft\",\"outputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reinitialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundActivationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundLateBidDurationDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundLateBidPricePremiumAmountBaseMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundLateBidPricePremiumAmountExponent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setBidCstRewardAmountMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setBidCstRewardAmountPerMinute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setBidMessageLengthMaxLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setCharityAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCharityEthDonationAmountPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setChronoWarriorEthPrizeAmountPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICosmicSignatureNft\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setCosmicSignatureNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCosmicSignatureNftStakingTotalEthRewardAmountPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICosmicSignatureToken\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setCosmicSignatureToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCstDutchAuctionBeginningBidPriceMinLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCstDutchAuctionDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCstDutchAuctionDurationChangeDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCstPrizeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setDelayDurationBeforeRoundActivation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setEthBidPriceIncreaseDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setEthBidRefundAmountInGasToSwallowMaxLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setEthDutchAuctionDurationDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setEthDutchAuctionEndingBidPriceDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setInitialDurationUntilMainPrizeDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setMainEthPrizeAmountPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setMainPrizeNumCosmicSignatureNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setMainPrizeTimeIncrementInMicroSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setMainPrizeTimeIncrementIncreaseDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setMarketingWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setMarketingWalletCstContributionAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleCosmicSignatureNftsForBidders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleCosmicSignatureNftsForRandomWalkNftStakers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleEthPrizesForBidders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPrizesWallet\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setPrizesWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setRaffleTotalEthPrizeAmountForBiddersPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRandomWalkNFT\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setRandomWalkNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setRoundActivationTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setRoundLateBidDurationDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setRoundLateBidPricePremiumAmountBaseMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setRoundLateBidPricePremiumAmountExponent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakingWalletCosmicSignatureNft\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setStakingWalletCosmicSignatureNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakingWalletRandomWalkNft\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setStakingWalletRandomWalkNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setTimeoutDurationToClaimMainPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingWalletCosmicSignatureNft\",\"outputs\":[{\"internalType\":\"contractStakingWalletCosmicSignatureNft\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingWalletRandomWalkNft\",\"outputs\":[{\"internalType\":\"contractStakingWalletRandomWalkNft\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeoutDurationToClaimMainPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractCosmicSignatureToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tryGetCurrentChampions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"enduranceChampionAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"enduranceChampionDuration_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"chronoWarriorAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chronoWarriorDuration_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"usedRandomWalkNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nftWasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523461003e5761001161004d565b610019610043565b61b9a16103d9823960805181818161930c01528181619371015261953f015261b9a190f35b610049565b60405190565b5f80fd5b610055610057565b565b61005f610061565b565b61006961006b565b565b610073610075565b565b61007d61007f565b565b610087610089565b565b610091610093565b565b61009b61009d565b565b6100a56100a7565b565b6100af6100b1565b565b6100b96100bb565b565b6100c36100c5565b565b6100cd6100cf565b565b6100d76100d9565b565b6100e16100e3565b565b6100eb6100ed565b565b6100f56100f7565b565b6100ff610101565b565b61010961010b565b565b610113610115565b565b61011d61011f565b565b610127610129565b565b610131610133565b565b61013b61013d565b565b610145610147565b565b61014f610159565b6101576102fc565b565b610161610163565b565b61016b61016d565b565b610175610177565b565b61017f610181565b565b61018961018b565b565b610193610195565b565b61019d61019f565b565b6101a76101a9565b565b6101b16101b3565b565b6101bb6101ff565b565b60018060a01b031690565b90565b6101df6101da6101e4926101bd565b6101c8565b6101bd565b90565b6101f0906101cb565b90565b6101fc906101e7565b90565b610208306101f3565b608052565b60401c90565b60ff1690565b61022561022a9161020d565b610213565b90565b6102379054610219565b90565b5f0190565b5f1c90565b60018060401b031690565b61025b6102609161023f565b610244565b90565b61026d905461024f565b90565b60018060401b031690565b5f1b90565b9061029160018060401b039161027b565b9181191691161790565b6102af6102aa6102b492610270565b6101c8565b610270565b90565b90565b906102cf6102ca6102d69261029b565b6102b7565b8254610280565b9055565b6102e390610270565b9052565b91906102fa905f602085019401906102da565b565b6103046103b4565b61030f5f820161022d565b6103985761031e5f8201610263565b61033661033060018060401b03610270565b91610270565b0361033f575b50565b610352905f60018060401b0391016102ba565b60018060401b0361038f7fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d291610386610043565b918291826102e7565b0390a15f61033c565b5f63f92ee8a960e01b8152806103b06004820161023a565b0390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009056fe6080604052600436101561001d575b366138425761001b616cfc565b005b6100275f356108c5565b80620ac9f1146108c0578063040d4d31146108bb57806304338479146108b657806309632366146108b157806309794bee146108ac5780630a120648146108a75780630b5f95ae146108a25780630c9be46d1461089d5780630eb16be614610898578063119b22b31461089357806311b0d1fe1461088e578063135f3d281461088957806317887731146108845780631824d5e71461087f57806318305de21461087a5780631aaba5a5146108755780631b410319146108705780631e9cbb7e1461086b5780631f1b4aa41461086657806323b31cfc14610861578063250fadb61461085c57806325ed499c146108575780632665c88214610852578063277004811461084d57806327995f07146108485780632afa2580146108435780632b8dcbba1461083e5780632b91c7bb146108395780632d809e88146108345780632d829a2d1461082f5780632f894cd71461082a5780632fb3c48f14610825578063320c435c14610820578063329b95a51461081b57806336750d2c1461081657806337b99cc7146108115780633b9d292e1461080c5780634164b95b14610807578063441b328914610802578063448c6eb1146107fd57806344a4b917146107f857806344acc12a146107f3578063477adf2a146107ee57806347ccca02146107e95780634ad8a90e146107e45780634c2a4a33146107df5780634e452010146107da5780634f1ef286146107d55780634f734612146107d05780634feb78b7146107cb57806352d1902d146107c6578063543f416f146107c157806354ada1d6146107bc57806356732241146107b75780635863a705146107b25780635a1e5bde146107ad5780635b0a45d9146107a85780635d098b38146107a35780635f0112fe1461079e5780635fdf49cb1461079957806360ef88411461079457806362ed9b531461078f5780636b59acb81461078a5780636b7cbe85146107855780636c0613c0146107805780636c17e3cc1461077b5780636c2eb350146107765780636e95d286146107715780636e9708341461076c578063715018a61461076757806371b6d01914610762578063755b4ef71461075d57806375ef3b9c1461075857806375f0a8741461075357806377fa10271461074e57806387292a8514610749578063876d5c361461074457806388ce802c1461073f5780638c94e9ba1461073a5780638da5cb5b14610735578063928880fa146107305780639302020f1461072b5780639646d7581461072657806399bf353d146107215780639aa1b38d1461071c5780639e50acc9146107175780639edeaf8e14610712578063a35286d11461070d578063a4be0d4014610708578063a922ab5d14610703578063a9742016146106fe578063aadd1b03146106f9578063ad3cb1cc146106f4578063ad4b0e8a146106ef578063afcf2fc4146106ea578063b30f5bb1146106e5578063b4f1b134146106e0578063b5d1f06f146106db578063b6a94f42146106d6578063b700db5f146106d1578063b78d1e2a146106cc578063b9cf9ba5146106c7578063baab4430146106c2578063bb4b3e6f146106bd578063be720ad5146106b8578063c52d8549146106b3578063c7e7a601146106ae578063c87baab5146106a9578063cb720d4d146106a4578063cfb4e5991461069f578063d1f8fcf21461069a578063d7559b9c14610695578063d9ab9eaa14610690578063da9931dd1461068b578063ddd6df0714610686578063de704b4114610681578063dfcd00d11461067c578063e2f9185f14610677578063e5b3cd1414610672578063eaace3021461066d578063eb13430e14610668578063ebaa1ea814610663578063ebb9bc5c1461065e578063ecb5776e14610659578063ef22d15b14610654578063efeb248a1461064f578063f0bdab7c1461064a578063f11400f014610645578063f2fde38b14610640578063f34d411c1461063b578063f444b29814610636578063f49efe9d14610631578063f7bea0781461062c578063fbaf508414610627578063fc0c546a14610622578063fd77310f1461061d578063fd9b3747146106185763fdfb9ba40361000e5761380f565b6137d8565b613780565b61374b565b6136a6565b613671565b61362e565b6135f9565b6135b6565b613583565b61354e565b61350b565b6134d6565b613491565b61345e565b613429565b6133c4565b61337f565b61333a565b6132f5565b6132b0565b61326d565b61323a565b613207565b6131d2565b61318f565b61315b565b6130d3565b613090565b613057565b612fe4565b612f9f565b612f5a565b612f15565b612ed0565b612e8b565b612e56565b612e19565b612d75565b612d1d565b612ce4565b612aa2565b612a20565b6129db565b612996565b612951565b61284f565b61281a565b6127d5565b612733565b6126fe565b6126cb565b612649565b612604565b6125bf565b61257c565b6124fa565b6124c0565b612438565b612405565b6123d2565b6123a5565b612304565b6122d1565b61229c565b612259565b612224565b612180565b61213d565b612108565b6120c3565b612090565b61205d565b611fdb565b611f96565b611f22565b611eed565b611eb8565b611e83565b611e3e565b611e0b565b611dd6565b611da3565b611d21565b611cdc565b611c97565b611c54565b611c1f565b611bc7565b611b92565b611b68565b611a72565b611a3d565b6119f8565b6119b3565b61190f565b6118cc565b61184a565b611807565b6117d4565b61179f565b61175a565b611715565b6116e0565b6116a9565b6115d7565b611592565b611530565b6114fd565b6114c8565b611426565b6113f1565b6113bc565b611318565b6112e3565b61129e565b61126b565b611236565b6111fc565b61114c565b611119565b6110e4565b6110a1565b61106c565b611027565b610fde565b610e6b565b610e34565b610c77565b610c05565b610bd2565b610b5c565b610b27565b610a91565b610a5e565b610a2b565b6109d3565b610969565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b90565b6108e9816108dd565b036108f057565b5f80fd5b90503590610901826108e0565b565b919060408382031261092b578061091f610928925f86016108f4565b936020016108f4565b90565b6108d5565b60018060a01b031690565b61094490610930565b90565b6109509061093b565b9052565b9190610967905f60208501940190610947565b565b3461099a5761099661098561097f366004610903565b90613860565b61098d6108cb565b91829182610954565b0390f35b6108d1565b5f9103126109a957565b6108d5565b90565b6109ba906109ae565b9052565b91906109d1905f602085019401906109b1565b565b34610a03576109e336600461099f565b6109ff6109ee6138ba565b6109f66108cb565b918291826109be565b0390f35b6108d1565b90602082820312610a2157610a1e915f016108f4565b90565b6108d5565b5f0190565b34610a5957610a43610a3e366004610a08565b613996565b610a4b6108cb565b80610a5581610a26565b0390f35b6108d1565b34610a8c57610a76610a71366004610a08565b613a0e565b610a7e6108cb565b80610a8881610a26565b0390f35b6108d1565b34610abf57610aa9610aa4366004610a08565b613a73565b610ab16108cb565b80610abb81610a26565b0390f35b6108d1565b610acd816109ae565b03610ad457565b5f80fd5b90503590610ae582610ac4565b565b90602082820312610b0057610afd915f01610ad8565b90565b6108d5565b610b0e906108dd565b9052565b9190610b25905f60208501940190610b05565b565b34610b5757610b53610b42610b3d366004610ae7565b613b72565b610b4a6108cb565b91829182610b12565b0390f35b6108d1565b34610b8c57610b6c36600461099f565b610b88610b77613c80565b610b7f6108cb565b91829182610b12565b0390f35b6108d1565b610b9a8161093b565b03610ba157565b5f80fd5b90503590610bb282610b91565b565b90602082820312610bcd57610bca915f01610ba5565b90565b6108d5565b34610c0057610bea610be5366004610bb4565b613d5d565b610bf26108cb565b80610bfc81610a26565b0390f35b6108d1565b34610c3557610c1536600461099f565b610c31610c20613d93565b610c286108cb565b91829182610b12565b0390f35b6108d1565b1c90565b90565b610c51906008610c569302610c3a565b610c3e565b90565b90610c649154610c41565b90565b610c7461010b5f90610c59565b90565b34610ca757610c8736600461099f565b610ca3610c92610c67565b610c9a6108cb565b91829182610b12565b0390f35b6108d1565b5f80fd5b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b90610cdc90610cb4565b810190811067ffffffffffffffff821117610cf657604052565b610cbe565b90610d0e610d076108cb565b9283610cd2565b565b67ffffffffffffffff8111610d2e57610d2a602091610cb4565b0190565b610cbe565b90825f939282370152565b90929192610d53610d4e82610d10565b610cfb565b93818552602085019082840111610d6f57610d6d92610d33565b565b610cb0565b9080601f83011215610d9257816020610d8f93359101610d3e565b90565b610cac565b610da09061093b565b90565b610dac81610d97565b03610db357565b5f80fd5b90503590610dc482610da3565b565b919060a083820312610e2f57610dde815f85016108f4565b92602081013567ffffffffffffffff8111610e2a5782610dff918301610d74565b92610e27610e1084604085016108f4565b93610e1e8160608601610db7565b936080016108f4565b90565b6108d9565b6108d5565b34610e6657610e50610e47366004610dc6565b93929092613f49565b610e586108cb565b80610e6281610a26565b0390f35b6108d1565b34610e9957610e83610e7e366004610a08565b613fc5565b610e8b6108cb565b80610e9581610a26565b0390f35b6108d1565b9190604083820312610ec65780610eba610ec3925f86016108f4565b93602001610ba5565b90565b6108d5565b90565b610ee2610edd610ee7926108dd565b610ecb565b6108dd565b90565b90610ef490610ece565b5f5260205260405f2090565b610f14610f0f610f1992610930565b610ecb565b610930565b90565b610f2590610f00565b90565b610f3190610f1c565b90565b90610f3e90610f28565b5f5260205260405f2090565b5f1c90565b610f5b610f6091610f4a565b610c3e565b90565b610f6d9054610f4f565b90565b90610f80610f8592610104610eea565b610f34565b610f905f8201610f63565b91610fa96002610fa260018501610f63565b9301610f63565b90565b604090610fd5610fdc9496959396610fcb60608401985f850190610b05565b6020830190610b05565b0190610b05565b565b346110125761100e610ffa610ff4366004610e9e565b90610f70565b6110059391936108cb565b93849384610fac565b0390f35b6108d1565b61102461010a5f90610c59565b90565b346110575761103736600461099f565b611053611042611017565b61104a6108cb565b91829182610b12565b0390f35b6108d1565b6110696101245f90610c59565b90565b3461109c5761107c36600461099f565b61109861108761105c565b61108f6108cb565b91829182610b12565b0390f35b6108d1565b346110cf576110b96110b4366004610a08565b61403d565b6110c16108cb565b806110cb81610a26565b0390f35b6108d1565b6110e16101165f90610c59565b90565b34611114576110f436600461099f565b6111106110ff6110d4565b6111076108cb565b91829182610b12565b0390f35b6108d1565b346111475761113161112c366004610a08565b6140b5565b6111396108cb565b8061114381610a26565b0390f35b6108d1565b3461117a5761116461115f366004610a08565b61412d565b61116c6108cb565b8061117681610a26565b0390f35b6108d1565b5f80fd5b5f80fd5b909182601f830112156111c15781359167ffffffffffffffff83116111bc5760200192600183028401116111b757565b611183565b61117f565b610cac565b906020828203126111f7575f82013567ffffffffffffffff81116111f2576111ee9201611187565b9091565b6108d9565b6108d5565b61121061120a3660046111c6565b9061448d565b6112186108cb565b8061122281610a26565b0390f35b6112336101065f90610c59565b90565b346112665761124636600461099f565b611262611251611226565b6112596108cb565b91829182610b12565b0390f35b6108d1565b346112995761128361127e366004610a08565b614506565b61128b6108cb565b8061129581610a26565b0390f35b6108d1565b346112ce576112ae36600461099f565b6112ca6112b961451d565b6112c16108cb565b91829182610b12565b0390f35b6108d1565b6112e06101145f90610c59565b90565b34611313576112f336600461099f565b61130f6112fe6112d3565b6113066108cb565b91829182610b12565b0390f35b6108d1565b346113485761134461133361132e366004610a08565b6145c1565b61133b6108cb565b91829182610b12565b0390f35b6108d1565b60018060a01b031690565b61136890600861136d9302610c3a565b61134d565b90565b9061137b9154611358565b90565b61138b61012c5f90611370565b90565b61139790610f1c565b90565b6113a39061138e565b9052565b91906113ba905f6020850194019061139a565b565b346113ec576113cc36600461099f565b6113e86113d761137e565b6113df6108cb565b918291826113a7565b0390f35b6108d1565b346114215761140136600461099f565b61141d61140c614603565b6114146108cb565b91829182610b12565b0390f35b6108d1565b346114545761143e611439366004610a08565b61465f565b6114466108cb565b8061145081610a26565b0390f35b6108d1565b60018060a01b031690565b6114749060086114799302610c3a565b611459565b90565b906114879154611464565b90565b61149761012d5f9061147c565b90565b6114a390610f1c565b90565b6114af9061149a565b9052565b91906114c6905f602085019401906114a6565b565b346114f8576114d836600461099f565b6114f46114e361148a565b6114eb6108cb565b918291826114b3565b0390f35b6108d1565b3461152b57611515611510366004610a08565b6146d7565b61151d6108cb565b8061152781610a26565b0390f35b6108d1565b3461155e57611548611543366004610a08565b61474f565b6115506108cb565b8061155a81610a26565b0390f35b6108d1565b9061156d90610ece565b5f5260205260405f2090565b5f61158961158f92610103611563565b01610f63565b90565b346115c2576115be6115ad6115a8366004610a08565b611579565b6115b56108cb565b91829182610b12565b0390f35b6108d1565b6115d461011c5f90610c59565b90565b34611607576115e736600461099f565b6116036115f26115c7565b6115fa6108cb565b91829182610b12565b0390f35b6108d1565b6116159061093b565b90565b6116218161160c565b0361162857565b5f80fd5b9050359061163982611618565b565b919060a0838203126116a457611653815f85016108f4565b92602081013567ffffffffffffffff811161169f5782611674918301610d74565b9261169c61168584604085016108f4565b93611693816060860161162c565b936080016108f4565b90565b6108d9565b6108d5565b346116db576116c56116bc36600461163b565b93929092614889565b6116cd6108cb565b806116d781610a26565b0390f35b6108d1565b34611710576116f036600461099f565b61170c6116fb614898565b6117036108cb565b91829182610b12565b0390f35b6108d1565b346117455761172536600461099f565b6117416117306148e7565b6117386108cb565b918291826109be565b0390f35b6108d1565b6117576101275f90610c59565b90565b3461178a5761176a36600461099f565b61178661177561174a565b61177d6108cb565b91829182610b12565b0390f35b6108d1565b61179c6101305f90610c59565b90565b346117cf576117af36600461099f565b6117cb6117ba61178f565b6117c26108cb565b91829182610b12565b0390f35b6108d1565b34611802576117ec6117e7366004610a08565b614985565b6117f46108cb565b806117fe81610a26565b0390f35b6108d1565b346118355761181736600461099f565b61181f614d79565b6118276108cb565b8061183181610a26565b0390f35b6108d1565b6118476101235f90610c59565b90565b3461187a5761185a36600461099f565b61187661186561183a565b61186d6108cb565b91829182610b12565b0390f35b6108d1565b6118889061093b565b90565b6118948161187f565b0361189b57565b5f80fd5b905035906118ac8261188b565b565b906020828203126118c7576118c4915f0161189f565b90565b6108d5565b346118fa576118e46118df3660046118ae565b614e8e565b6118ec6108cb565b806118f681610a26565b0390f35b6108d1565b61190c61011e5f90610c59565b90565b3461193f5761191f36600461099f565b61193b61192a6118ff565b6119326108cb565b91829182610b12565b0390f35b6108d1565b60018060a01b031690565b61195f9060086119649302610c3a565b611944565b90565b90611972915461194f565b90565b61198261012b5f90611967565b90565b61198e90610f1c565b90565b61199a90611985565b9052565b91906119b1905f60208501940190611991565b565b346119e3576119c336600461099f565b6119df6119ce611975565b6119d66108cb565b9182918261199e565b0390f35b6108d1565b6119f56101375f90610c59565b90565b34611a2857611a0836600461099f565b611a24611a136119e8565b611a1b6108cb565b91829182610b12565b0390f35b6108d1565b611a3a6101205f90610c59565b90565b34611a6d57611a4d36600461099f565b611a69611a58611a2d565b611a606108cb565b91829182610b12565b0390f35b6108d1565b34611aa257611a9e611a8d611a88366004610ae7565b614e99565b611a956108cb565b91829182610b12565b0390f35b6108d1565b67ffffffffffffffff8111611ac557611ac1602091610cb4565b0190565b610cbe565b90929192611adf611ada82611aa7565b610cfb565b93818552602085019082840111611afb57611af992610d33565b565b610cb0565b9080601f83011215611b1e57816020611b1b93359101611aca565b90565b610cac565b919091604081840312611b6357611b3c835f8301610ba5565b92602082013567ffffffffffffffff8111611b5e57611b5b9201611b00565b90565b6108d9565b6108d5565b611b7c611b76366004611b23565b90614ee0565b611b846108cb565b80611b8e81610a26565b0390f35b34611bc257611ba236600461099f565b611bbe611bad614f40565b611bb56108cb565b91829182610b12565b0390f35b6108d1565b34611bf557611bdf611bda366004610a08565b614fd5565b611be76108cb565b80611bf181610a26565b0390f35b6108d1565b90565b611c0690611bfa565b9052565b9190611c1d905f60208501940190611bfd565b565b34611c4f57611c2f36600461099f565b611c4b611c3a61504f565b611c426108cb565b91829182611c0a565b0390f35b6108d1565b34611c8257611c6c611c67366004610a08565b6150cf565b611c746108cb565b80611c7e81610a26565b0390f35b6108d1565b611c9461011d5f90610c59565b90565b34611cc757611ca736600461099f565b611cc3611cb2611c87565b611cba6108cb565b91829182610b12565b0390f35b6108d1565b611cd96101265f90610c59565b90565b34611d0c57611cec36600461099f565b611d08611cf7611ccc565b611cff6108cb565b91829182610b12565b0390f35b6108d1565b611d1e6101075f90610c59565b90565b34611d5157611d3136600461099f565b611d4d611d3c611d11565b611d446108cb565b91829182610b12565b0390f35b6108d1565b611d5f9061093b565b90565b611d6b81611d56565b03611d7257565b5f80fd5b90503590611d8382611d62565b565b90602082820312611d9e57611d9b915f01611d76565b90565b6108d5565b34611dd157611dbb611db6366004611d85565b6151e5565b611dc36108cb565b80611dcd81610a26565b0390f35b6108d1565b34611e0657611de636600461099f565b611e02611df16151f0565b611df96108cb565b91829182610b12565b0390f35b6108d1565b34611e3957611e23611e1e366004610bb4565b6152b5565b611e2b6108cb565b80611e3581610a26565b0390f35b6108d1565b34611e6e57611e4e36600461099f565b611e6a611e596152c0565b611e616108cb565b91829182610b12565b0390f35b6108d1565b611e806101385f90610c59565b90565b34611eb357611e9336600461099f565b611eaf611e9e611e73565b611ea66108cb565b91829182610b12565b0390f35b6108d1565b34611ee857611ec836600461099f565b611ee4611ed36152fd565b611edb6108cb565b91829182610b12565b0390f35b6108d1565b34611f1d57611efd36600461099f565b611f19611f08615328565b611f106108cb565b91829182610b12565b0390f35b6108d1565b34611f5057611f3a611f35366004610a08565b6153b2565b611f426108cb565b80611f4c81610a26565b0390f35b6108d1565b60018060a01b031690565b611f70906008611f759302610c3a565b611f55565b90565b90611f839154611f60565b90565b611f936101095f90611f78565b90565b34611fc657611fa636600461099f565b611fc2611fb1611f86565b611fb96108cb565b91829182610954565b0390f35b6108d1565b611fd861011a5f90610c59565b90565b3461200b57611feb36600461099f565b612007611ff6611fcb565b611ffe6108cb565b91829182610b12565b0390f35b6108d1565b6120199061093b565b90565b61202581612010565b0361202c57565b5f80fd5b9050359061203d8261201c565b565b9060208282031261205857612055915f01612030565b90565b6108d5565b3461208b5761207561207036600461203f565b6154c8565b61207d6108cb565b8061208781610a26565b0390f35b6108d1565b346120be576120a036600461099f565b6120a86159c6565b6120b06108cb565b806120ba81610a26565b0390f35b6108d1565b346120f3576120d336600461099f565b6120ef6120de6159d0565b6120e66108cb565b91829182610b12565b0390f35b6108d1565b61210561010d5f90610c59565b90565b346121385761211836600461099f565b6121346121236120f8565b61212b6108cb565b91829182610b12565b0390f35b6108d1565b3461216b5761214d36600461099f565b612155615a12565b61215d6108cb565b8061216781610a26565b0390f35b6108d1565b61217d6101345f90610c59565b90565b346121b05761219036600461099f565b6121ac61219b612170565b6121a36108cb565b91829182610b12565b0390f35b6108d1565b60018060a01b031690565b6121d09060086121d59302610c3a565b6121b5565b90565b906121e391546121c0565b90565b6121f361012a5f906121d8565b90565b6121ff90610f1c565b90565b61220b906121f6565b9052565b9190612222905f60208501940190612202565b565b346122545761223436600461099f565b61225061223f6121e6565b6122476108cb565b9182918261220f565b0390f35b6108d1565b346122875761227161226c366004610a08565b615a89565b6122796108cb565b8061228381610a26565b0390f35b6108d1565b61229961012f5f90611f78565b90565b346122cc576122ac36600461099f565b6122c86122b761228c565b6122bf6108cb565b91829182610954565b0390f35b6108d1565b346122ff576122e96122e4366004610a08565b615b01565b6122f16108cb565b806122fb81610a26565b0390f35b6108d1565b346123325761231c612317366004610a08565b615b79565b6123246108cb565b8061232e81610a26565b0390f35b6108d1565b919060a0838203126123a05761234f815f8501610ad8565b92602081013567ffffffffffffffff811161239b5782612370918301610d74565b9261239861238184604085016108f4565b9361238f8160608601610db7565b936080016108f4565b90565b6108d9565b6108d5565b6123bc6123b3366004612337565b93929092615c5c565b6123c46108cb565b806123ce81610a26565b0390f35b34612400576123ea6123e5366004610a08565b615cd8565b6123f26108cb565b806123fc81610a26565b0390f35b6108d1565b346124335761241d612418366004610a08565b615d50565b6124256108cb565b8061242f81610a26565b0390f35b6108d1565b346124685761244836600461099f565b612464612453615d5b565b61245b6108cb565b91829182610954565b0390f35b6108d1565b90916060828403126124bb57612485835f8401610ad8565b9260208301359067ffffffffffffffff82116124b6576124aa816124b3938601610d74565b936040016108f4565b90565b6108d9565b6108d5565b6124d46124ce36600461246d565b91615da5565b6124dc6108cb565b806124e681610a26565b0390f35b6124f76101155f90610c59565b90565b3461252a5761250a36600461099f565b6125266125156124ea565b61251d6108cb565b91829182610b12565b0390f35b6108d1565b6125389061093b565b90565b6125448161252f565b0361254b57565b5f80fd5b9050359061255c8261253b565b565b9060208282031261257757612574915f0161254f565b90565b6108d5565b346125aa5761259461258f36600461255e565b615ebd565b61259c6108cb565b806125a681610a26565b0390f35b6108d1565b6125bc6101355f90610c59565b90565b346125ef576125cf36600461099f565b6125eb6125da6125af565b6125e26108cb565b91829182610b12565b0390f35b6108d1565b6126016101135f90610c59565b90565b346126345761261436600461099f565b61263061261f6125f4565b6126276108cb565b91829182610b12565b0390f35b6108d1565b6126466101055f90611f78565b90565b346126795761265936600461099f565b612675612664612639565b61266c6108cb565b91829182610954565b0390f35b6108d1565b6126879061093b565b90565b6126938161267e565b0361269a57565b5f80fd5b905035906126ab8261268a565b565b906020828203126126c6576126c3915f0161269e565b90565b6108d5565b346126f9576126e36126de3660046126ad565b615fd3565b6126eb6108cb565b806126f581610a26565b0390f35b6108d1565b3461272e5761270e36600461099f565b61272a612719615fde565b6127216108cb565b91829182610b12565b0390f35b6108d1565b346127615761274b612746366004610a08565b61604c565b6127536108cb565b8061275d81610a26565b0390f35b6108d1565b60018060a01b031690565b6127819060086127869302610c3a565b612766565b90565b906127949154612771565b90565b6127a461012e5f90612789565b90565b6127b090610f1c565b90565b6127bc906127a7565b9052565b91906127d3905f602085019401906127b3565b565b34612805576127e536600461099f565b6128016127f0612797565b6127f86108cb565b918291826127c0565b0390f35b6108d1565b6128176101125f90610c59565b90565b3461284a5761282a36600461099f565b61284661283561280a565b61283d6108cb565b91829182610b12565b0390f35b6108d1565b61285a36600461099f565b6128626160d3565b61286a6108cb565b8061287481610a26565b0390f35b9061288a61288583610d10565b610cfb565b918252565b5f7f352e302e30000000000000000000000000000000000000000000000000000000910152565b6128c06005612878565b906128cd6020830161288f565b565b6128d76128b6565b90565b6128e26128cf565b90565b6128ed6128da565b90565b5190565b60209181520190565b90825f9392825e0152565b6129276129306020936129359361291e816128f0565b938480936128f4565b958691016128fd565b610cb4565b0190565b61294e9160208201915f818403910152612908565b90565b346129815761296136600461099f565b61297d61296c6128e5565b6129746108cb565b91829182612939565b0390f35b6108d1565b6129936101025f90611f78565b90565b346129c6576129a636600461099f565b6129c26129b1612986565b6129b96108cb565b91829182610954565b0390f35b6108d1565b6129d86101315f90611f78565b90565b34612a0b576129eb36600461099f565b612a076129f66129cb565b6129fe6108cb565b91829182610954565b0390f35b6108d1565b612a1d61011b5f90610c59565b90565b34612a5057612a3036600461099f565b612a4c612a3b612a10565b612a436108cb565b91829182610b12565b0390f35b6108d1565b612a5e9061093b565b90565b612a6a81612a55565b03612a7157565b5f80fd5b90503590612a8282612a61565b565b90602082820312612a9d57612a9a915f01612a75565b90565b6108d5565b34612ad057612aba612ab5366004612a84565b6161e8565b612ac26108cb565b80612acc81610a26565b0390f35b6108d1565b634e487b7160e01b5f52603260045260245ffd5b5490565b5f5260205f2090565b612aff81612ae9565b821015612b1957612b11600491612aed565b910201905f90565b612ad5565b612b2a612b2f91610f4a565b611f55565b90565b612b3c9054612b1e565b90565b634e487b7160e01b5f52602260045260245ffd5b9060016002830492168015612b73575b6020831014612b6e57565b612b3f565b91607f1691612b63565b60209181520190565b5f5260205f2090565b905f9291805490612ba9612ba283612b53565b8094612b7d565b916001811690815f14612c005750600114612bc4575b505050565b612bd19192939450612b86565b915f925b818410612be857505001905f8080612bbf565b60018160209295939554848601520191019290612bd5565b92949550505060ff19168252151560200201905f8080612bbf565b90612c2591612b8f565b90565b90612c48612c4192612c386108cb565b93848092612c1b565b0383610cd2565b565b61010090612c5782612ae9565b811015612c9d57612c6791612af6565b5090612c745f8301610f63565b91612c8160018201612b32565b91612c9a6003612c9360028501610f63565b9301612c28565b90565b5f80fd5b9092612cd490612cca612ce19694612cc060808601975f870190610b05565b6020850190610947565b6040830190610b05565b6060818403910152612908565b90565b34612d1857612d14612cff612cfa366004610a08565b612c4a565b90612d0b9492946108cb565b94859485612ca1565b0390f35b6108d1565b34612d4d57612d49612d38612d33366004610ae7565b6161f3565b612d406108cb565b91829182610b12565b0390f35b6108d1565b916020612d73929493612d6c60408201965f830190610b05565b01906109b1565b565b34612da657612d8536600461099f565b612d8d616211565b90612da2612d996108cb565b92839283612d52565b0390f35b6108d1565b919060a083820312612e1457612dc3815f8501610ad8565b92602081013567ffffffffffffffff8111612e0f5782612de4918301610d74565b92612e0c612df584604085016108f4565b93612e03816060860161162c565b936080016108f4565b90565b6108d9565b6108d5565b612e30612e27366004612dab565b93929092616313565b612e386108cb565b80612e4281610a26565b0390f35b612e5361010c5f90610c59565b90565b34612e8657612e6636600461099f565b612e82612e71612e46565b612e796108cb565b91829182610b12565b0390f35b6108d1565b34612ebb57612e9b36600461099f565b612eb7612ea6616322565b612eae6108cb565b91829182610b12565b0390f35b6108d1565b612ecd6101185f90610c59565b90565b34612f0057612ee036600461099f565b612efc612eeb612ec0565b612ef36108cb565b91829182610b12565b0390f35b6108d1565b612f126101325f90610c59565b90565b34612f4557612f2536600461099f565b612f41612f30612f05565b612f386108cb565b91829182610b12565b0390f35b6108d1565b612f576101365f90610c59565b90565b34612f8a57612f6a36600461099f565b612f86612f75612f4a565b612f7d6108cb565b91829182610b12565b0390f35b6108d1565b612f9c6101175f90610c59565b90565b34612fcf57612faf36600461099f565b612fcb612fba612f8f565b612fc26108cb565b91829182610b12565b0390f35b6108d1565b612fe161010f5f90610c59565b90565b3461301457612ff436600461099f565b613010612fff612fd4565b6130076108cb565b91829182610b12565b0390f35b6108d1565b61304e6130559461304460609498979561303a608086019a5f870190610947565b6020850190610b05565b6040830190610947565b0190610b05565b565b3461308b5761306736600461099f565b613087613072616364565b9061307e9492946108cb565b94859485613019565b0390f35b6108d1565b346130be576130a86130a3366004610a08565b6165e0565b6130b06108cb565b806130ba81610a26565b0390f35b6108d1565b6130d061010e5f90610c59565b90565b34613103576130e336600461099f565b6130ff6130ee6130c3565b6130f66108cb565b91829182610b12565b0390f35b6108d1565b909160608284031261315657613120835f84016108f4565b9260208301359067ffffffffffffffff8211613151576131458161314e938601610d74565b936040016108f4565b90565b6108d9565b6108d5565b3461318a5761317461316e366004613108565b91616617565b61317c6108cb565b8061318681610a26565b0390f35b6108d1565b346131bd576131a76131a2366004610a08565b616691565b6131af6108cb565b806131b981610a26565b0390f35b6108d1565b6131cf6101335f90610c59565b90565b34613202576131e236600461099f565b6131fe6131ed6131c2565b6131f56108cb565b91829182610b12565b0390f35b6108d1565b346132355761321f61321a366004610a08565b6166cd565b6132276108cb565b8061323181610a26565b0390f35b6108d1565b346132685761325261324d366004610a08565b616745565b61325a6108cb565b8061326481610a26565b0390f35b6108d1565b3461329b5761327d36600461099f565b613285616923565b61328d6108cb565b8061329781610a26565b0390f35b6108d1565b6132ad6101215f90610c59565b90565b346132e0576132c036600461099f565b6132dc6132cb6132a0565b6132d36108cb565b91829182610b12565b0390f35b6108d1565b6132f26101015f90611f78565b90565b346133255761330536600461099f565b6133216133106132e5565b6133186108cb565b91829182610954565b0390f35b6108d1565b6133376101085f90610c59565b90565b3461336a5761334a36600461099f565b61336661335561332a565b61335d6108cb565b91829182610b12565b0390f35b6108d1565b61337c6101255f90610c59565b90565b346133af5761338f36600461099f565b6133ab61339a61336f565b6133a26108cb565b91829182610b12565b0390f35b6108d1565b6133c16101105f90610c59565b90565b346133f4576133d436600461099f565b6133f06133df6133b4565b6133e76108cb565b91829182610b12565b0390f35b6108d1565b9061340390610ece565b5f5260205260405f2090565b61342690613421610119915f926133f9565b610c59565b90565b346134595761345561344461343f366004610a08565b61340f565b61344c6108cb565b91829182610b12565b0390f35b6108d1565b3461348c57613476613471366004610a08565b61699a565b61347e6108cb565b8061348881610a26565b0390f35b6108d1565b346134c1576134a136600461099f565b6134bd6134ac6169b3565b6134b46108cb565b918291826109be565b0390f35b6108d1565b6134d36101115f90610c59565b90565b34613506576134e636600461099f565b6135026134f16134c6565b6134f96108cb565b91829182610b12565b0390f35b6108d1565b346135395761352361351e366004610a08565b616a3c565b61352b6108cb565b8061353581610a26565b0390f35b6108d1565b61354b61011f5f90610c59565b90565b3461357e5761355e36600461099f565b61357a61356961353e565b6135716108cb565b91829182610b12565b0390f35b6108d1565b346135b15761359b613596366004610bb4565b616aac565b6135a36108cb565b806135ad81610a26565b0390f35b6108d1565b346135e4576135ce6135c9366004610a08565b616ae8565b6135d66108cb565b806135e081610a26565b0390f35b6108d1565b6135f66101285f90610c59565b90565b346136295761360936600461099f565b6136256136146135e9565b61361c6108cb565b91829182610b12565b0390f35b6108d1565b3461365c57613646613641366004610a08565b616b60565b61364e6108cb565b8061365881610a26565b0390f35b6108d1565b61366e6101225f90610c59565b90565b346136a15761368136600461099f565b61369d61368c613661565b6136946108cb565b91829182610b12565b0390f35b6108d1565b346136d7576136b636600461099f565b6136be616b6b565b906136d36136ca6108cb565b92839283612d52565b0390f35b6108d1565b60018060a01b031690565b6136f79060086136fc9302610c3a565b6136dc565b90565b9061370a91546136e7565b90565b61371a6101295f906136ff565b90565b61372690610f1c565b90565b6137329061371d565b9052565b9190613749905f60208501940190613729565b565b3461377b5761375b36600461099f565b61377761376661370d565b61376e6108cb565b91829182613736565b0390f35b6108d1565b346137b0576137ac61379b613796366004610a08565b616b94565b6137a36108cb565b91829182610b12565b0390f35b6108d1565b9160206137d69294936137cf60408201965f830190610b05565b0190610b05565b565b3461380a576137f16137eb366004610e9e565b90616bc1565b906138066137fd6108cb565b928392836137b5565b0390f35b6108d1565b3461383d57613827613822366004610a08565b616c78565b61382f6108cb565b8061383981610a26565b0390f35b6108d1565b5f80fd5b5f90565b9061385490610ece565b5f5260205260405f2090565b61388991600161387e61388493613875613846565b50610103611563565b0161384a565b612b32565b90565b5f90565b6138a461389f6138a9926108dd565b610ecb565b6109ae565b90565b906138b791036109ae565b90565b6138c261388c565b506138e86138cf42613890565b6138e26138dd61010d610f63565b613890565b906138ac565b90565b6138fc906138f7616d06565b6138fe565b565b61390f9061390a616dee565b61394f565b565b5f1b90565b906139225f1991613911565b9181191691161790565b90565b9061394461393f61394b92610ece565b61392c565b8254613916565b9055565b61395b8161013361392f565b6139917facbc6b6929088e4b2d043625fa7248e00fb6658a425eb9d9dc8c37b18c7f3e6f916139886108cb565b91829182610b12565b0390a1565b61399f906138eb565b565b6139b2906139ad616d06565b6139b4565b565b6139c5906139c0616dee565b6139c7565b565b6139d38161011b61392f565b613a097f40b9c59af8c486ccf8c7cc73df5a51e7cc29747ea7d39f99632ecaf9caa2ed1f91613a006108cb565b91829182610b12565b0390a1565b613a17906139a1565b565b613a2a90613a25616d06565b613a2c565b565b613a388161010c61392f565b613a6e7fb0868a729f47ae3829aaafe3ca2975d3db2148553c854112f598be6d91ef0d2891613a656108cb565b91829182610b12565b0390a1565b613a7c90613a19565b565b5f90565b90565b613a99613a94613a9e92613a82565b610ecb565b610930565b90565b613aaa90613a85565b90565b90613ab891016109ae565b90565b613acf613aca613ad492613a82565b610ecb565b6108dd565b90565b613aeb613ae6613af092613a82565b610ecb565b6109ae565b90565b613b07613b02613b0c926109ae565b610ecb565b6108dd565b90565b90613b1a91026108dd565b90565b90565b613b34613b2f613b3992613b1d565b610ecb565b6108dd565b90565b634e487b7160e01b5f52601260045260245ffd5b613b5c613b62916108dd565b916108dd565b908115613b6d570490565b613b3c565b613bd890613b7e613a7e565b50613b8a610101612b32565b613ba4613b9e613b995f613aa1565b61093b565b9161093b565b145f14613c3f57613bd2613bcc613bbc61010d610f63565b5b92613bc742613890565b613aad565b91613890565b906138ac565b613be15f613abb565b9080613bfd613bf7613bf25f613ad7565b6109ae565b916109ae565b13613c07575b5090565b613c399150613c18613c2991613af3565b613c23610137610f63565b90613b0f565b613c33603c613b20565b90613b50565b5f613c03565b613bd2613bcc613c7b6002613c75613c64610104613c5e61010b610f63565b90610eea565b613c6f610101612b32565b90610f34565b01610f63565b613bbd565b613c88613a7e565b50613c94610100612ae9565b90565b613ca890613ca3616d06565b613caa565b565b613cbb90613cb6616dee565b613cbd565b565b613ccf90613cca81616e9d565b613d0f565b565b90613ce260018060a01b0391613911565b9181191691161790565b90565b90613d04613cff613d0b92610f28565b613cec565b8254613cd1565b9055565b613d1b81610131613cef565b613d457f1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c91610f28565b90613d4e6108cb565b80613d5881610a26565b0390a2565b613d6690613c97565b565b613d7190610f1c565b90565b90565b613d8b613d86613d9092613d74565b610ecb565b6108dd565b90565b613d9b613a7e565b50613dcd613dbd613dab30613d68565b31613db7610132610f63565b90613b0f565b613dc76064613d77565b90613b50565b90565b90613de594939291613de0616f2c565b613e90565b613ded616f71565b565b613dfb613e0091610f4a565b61134d565b90565b613e0d9054613def565b90565b5f80fd5b60e01b90565b5f910312613e2457565b6108d5565b613e3290610f1c565b90565b613e3e90613e29565b9052565b613e77613e7e94613e6d606094989795613e63608086019a5f870190610b05565b6020850190610947565b6040830190613e35565b0190610b05565b565b613e886108cb565b3d5f823e3d90fd5b91613ea0929493949190916170c7565b613eb3613eae61012c613e03565b61138e565b9063e2051c7e91613ec561010b610f63565b91613ece61731d565b9490823b15613f44575f94613f018692613ef694613eea6108cb565b998a9889978896613e14565b865260048601613e42565b03925af18015613f3f57613f13575b50565b613f32905f3d8111613f38575b613f2a8183610cd2565b810190613e1a565b5f613f10565b503d613f20565b613e80565b613e10565b90613f5694939291613dd0565b565b613f6990613f64616d06565b613f6b565b565b613f7c90613f77616dee565b613f7e565b565b613f8a8161012161392f565b613fc07f3312247fcf207243294680d9103851c8709e19be3d46ee4b1bff5e12d1c5ca7b91613fb76108cb565b91829182610b12565b0390a1565b613fce90613f58565b565b613fe190613fdc616d06565b613fe3565b565b613ff490613fef616dee565b613ff6565b565b6140028161013661392f565b6140387fcb78cca7628d232a9c7beef53b62f7204d9eacb44de85a8f593e6b0bb72a16219161402f6108cb565b91829182610b12565b0390a1565b61404690613fd0565b565b61405990614054616d06565b61405b565b565b61406c90614067616dee565b61406e565b565b61407a8161012061392f565b6140b07f85d8bf21006916836edd67a5afeed2e891cf876a6c9cd9babf7f42f4b007c24f916140a76108cb565b91829182610b12565b0390a1565b6140be90614048565b565b6140d1906140cc616d06565b6140d3565b565b6140e4906140df616dee565b6140e6565b565b6140f28161011d61392f565b6141287f5581e31f5b8d4b3c45b8ab8bf67e3602ce588b423905eb6ad34bd6bc3c8486999161411f6108cb565b91829182610b12565b0390a1565b614136906140c0565b565b9061414a91614145616f2c565b6143b3565b614152616f71565b565b90565b5490565b5f5260205f2090565b61416d81614157565b8210156141875761417f60049161415b565b910201905f90565b612ad5565b61419581614157565b680100000000000000008110156141b9576141b591600182018155614164565b9091565b610cbe565b90565b5090565b601f602091010490565b1b90565b919060086141ee9102916141e85f19846141cf565b926141cf565b9181191691161790565b919061420e61420961421693610ece565b61392c565b9083546141d3565b9055565b61422c91614226613a7e565b916141f8565b565b5f5b82811061423c57505050565b8061424b5f600193850161421a565b01614230565b9190601f8111614261575b505050565b81811161426e575b61425c565b61428361427d6142a194612b86565b916141c5565b91602061428f826141c5565b91106142a9575b80910191039061422e565b5f8080614269565b505f614296565b906142c0905f1990600802610c3a565b191690565b816142cf916142b0565b906002021790565b916142e290826141c1565b9067ffffffffffffffff82116143a157614306826143008554612b53565b85614251565b5f90601f831160011461433957918091614328935f9261432d575b50506142c5565b90555b565b90915001355f80614321565b601f1983169161434885612b86565b925f5b8181106143895750916002939185600196941061436f575b5050500201905561432b565b61437f910135601f8416906142b0565b90555f8080614363565b9193602060018192878701358155019501920161434b565b610cbe565b906143b192916142d7565b565b614421906143c2610100612ae9565b9260036143e16143db6143d6610100614154565b61418c565b506141be565b926143f86143f061010b610f63565b5f860161392f565b61440c61440361731d565b60018601613cef565b614419346002860161392f565b9192016143a6565b61442c61010b610f63565b61443461731d565b34929161448861447661447061446a7fa08049565b10d44a06dca9bf05685b39bc370352043c5a003e8d35d45ebdc53f94610ece565b94610f28565b94610ece565b9461447f6108cb565b91829182610b12565b0390a4565b9061449791614138565b565b6144aa906144a5616d06565b6144ac565b565b6144bd906144b8616dee565b6144bf565b565b6144cb8161013761392f565b6145017f961b02838cd95976d0cac2e65ed131f45e19a84369d91d59d613dec94a0638c6916144f86108cb565b91829182610b12565b0390a1565b61450f90614499565b565b61451a90610f1c565b90565b614525613a7e565b5061455761454761453530614511565b3161454161011d610f63565b90613b0f565b6145516064613d77565b90613b50565b90565b90565b61457161456c6145769261455a565b610ecb565b6108dd565b90565b614583600261455d565b90565b90565b61459d6145986145a292614586565b610ecb565b6108dd565b90565b906145b091036108dd565b90565b906145be91016108dd565b90565b6145f2614600916145d0613a7e565b506145ec6145dc614579565b6145e66001614589565b906145a5565b906145b3565b6145fa614579565b90613b50565b90565b61460b613a7e565b5061462b61461a610125610f63565b614625610123610f63565b90613b50565b90565b61463f9061463a616d06565b614641565b565b6146529061464d6173a9565b614654565b565b61465d90617401565b565b6146689061462e565b565b61467b90614676616d06565b61467d565b565b61468e90614689616dee565b614690565b565b61469c8161013261392f565b6146d27ffe65b6d5007c66dadebc5095104ccd672c070f396dfdcfe1ee7f54201b9efa6d916146c96108cb565b91829182610b12565b0390a1565b6146e09061466a565b565b6146f3906146ee616d06565b6146f5565b565b61470690614701616dee565b614708565b565b6147148161011f61392f565b61474a7f4787028773c8f14bc3b4bc41f43a02329ae41105823287201a34782c530d35fd916147416108cb565b91829182610b12565b0390a1565b614758906146e2565b565b9061476f9493929161476a616f2c565b6147d0565b614777616f71565b565b61478290610f1c565b90565b61478e90614779565b9052565b6147c76147ce946147bd6060949897956147b3608086019a5f870190610b05565b6020850190610947565b6040830190614785565b0190610b05565b565b916147e0929493949190916170c7565b6147f36147ee61012c613e03565b61138e565b9063fe673fd39161480561010b610f63565b9161480e61731d565b9490823b15614884575f9461484186926148369461482a6108cb565b998a9889978896613e14565b865260048601614792565b03925af1801561487f57614853575b50565b614872905f3d8111614878575b61486a8183610cd2565b810190613e1a565b5f614850565b503d614860565b613e80565b613e10565b906148969493929161475a565b565b6148a0613a7e565b506148a96148e7565b806148c46148be6148b95f613ad7565b6109ae565b916109ae565b135f146148d8576148d490613af3565b5b90565b506148e25f613abb565b6148d5565b6148ef61388c565b50614915614906614901610124610f63565b613890565b61490f42613890565b906138ac565b90565b61492990614924616d06565b61492b565b565b61493c90614937616dee565b61493e565b565b61494a8161011361392f565b6149807fa787f26546d7eeea63d70fc31736f27ad28329e95982f3bc5a7e0280f497bbf5916149776108cb565b91829182610b12565b0390a1565b61498e90614918565b565b614998616f2c565b6149a0614be0565b6149a8616f71565b565b151590565b60207f757272656e742062696464696e6720726f756e64207965742e00000000000000917f54686572652068617665206265656e206e6f206269647320696e2074686520635f8201520152565b614a0960396040926128f4565b614a12816149af565b0190565b614a2b9060208101905f8183039101526149fc565b90565b634e487b7160e01b5f52601160045260245ffd5b614a51614a57919392936109ae565b926109ae565b808301925f82851215818312169285129112151617614a7257565b614a2e565b60607f2e00000000000000000000000000000000000000000000000000000000000000917f4f6e6c7920746865206c61737420626964646572206973207065726d697474655f8201527f6420746f20636c61696d207468652062696464696e6720726f756e64206d616960208201527f6e207072697a65206265666f726520612074696d656f7574206578706972657360408201520152565b614b1d60616080926128f4565b614b2681614a77565b0190565b606090614b63614b6a9496959396614b59614b4e608085018581035f870152614b10565b986020850190610947565b6040830190610947565b0190610b05565b565b5f7f4e6f7420656e6f7567682074696d652068617320656c61707365642e00000000910152565b614ba0601c6020926128f4565b614ba981614b6c565b0190565b916040614bde929493614bd7614bcc606083018381035f850152614b93565b966020830190610b05565b0190610b05565b565b614be861731d565b614c04614bfe614bf9610101612b32565b61093b565b9161093b565b145f14614c8d57614c3342614c2b614c25614c20610124610f63565b6108dd565b916108dd565b1015156149aa565b614c5b575b614c40617448565b614c4942617576565b614c51617f92565b614c59618f3c565b565b614c66610124610f63565b4290614c89614c736108cb565b928392638d31bb1560e01b845260048401614bad565b0390fd5b614cbd614c9b610101612b32565b614cb5614caf614caa5f613aa1565b61093b565b9161093b565b1415156149aa565b614d5657614ce5614ccc6148e7565b614cdf614cda610127610f63565b613890565b90614a42565b614d0b81614d03614cfd614cf85f613ad7565b6109ae565b916109ae565b1315156149aa565b614d155750614c38565b614d20610101612b32565b614d52614d34614d2e61731d565b93613af3565b614d3c6108cb565b93849363336598a360e21b855260048501614b2a565b0390fd5b614d5e6108cb565b6318844a7d60e31b815280614d7560048201614a16565b0390fd5b614d81614990565b565b614d9490614d8f616d06565b614d96565b565b614da790614da2616dee565b614db5565b565b614db290610f1c565b90565b614dcf90614dca614dc582614da9565b616e9d565b614e30565b565b614dda90610f00565b90565b614de690614dd1565b90565b614df290614dd1565b90565b90565b90614e0d614e08614e1492614de9565b614df5565b8254613cd1565b9055565b614e2190610f00565b90565b614e2d90614e18565b90565b614e4c614e44614e3f83614da9565b614ddd565b61012b614df8565b614e767f5bde6238168795ba4e77972a2bdaa5a465f7c9a5d05817f5e8d3fed2e5a4fa6091614e24565b90614e7f6108cb565b80614e8981610a26565b0390a2565b614e9790614d83565b565b614eb490614ea5613a7e565b50614eaf81619035565b6191a2565b90565b90614ec991614ec46192fb565b614ecb565b565b90614ede91614ed9816193c0565b619430565b565b90614eea91614eb7565b565b90565b614f03614efe614f0892614eec565b610ecb565b6108dd565b90565b614f166103e8614eef565b90565b614f32614f24614f0b565b614f2c614f0b565b90613b0f565b90565b614f3d614f19565b90565b614f48613a7e565b50614f65614f57610125610f63565b614f5f614f35565b90613b50565b90565b614f7990614f74616d06565b614f7b565b565b614f8c90614f87616dee565b614f8e565b565b614f9a8161012261392f565b614fd07f9e44c04f534af356419a731b967f3d56bc748b1f3fdbda7a89f4e1a321ade93491614fc76108cb565b91829182610b12565b0390a1565b614fde90614f68565b565b5f90565b614ff590614ff061952e565b615043565b90565b90565b61500f61500a61501492614ff8565b613911565b611bfa565b90565b6150407f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc614ffb565b90565b5061504c615017565b90565b61505f61505a614fe0565b614fe4565b90565b6150739061506e616d06565b615075565b565b61508690615081616dee565b615088565b565b6150948161011a61392f565b6150ca7f157c413b0549fd4f45aab72b7828304fb2c45dad53de0f1128c5eabf3aaabaf8916150c16108cb565b91829182610b12565b0390a1565b6150d890615062565b565b6150eb906150e6616d06565b6150ed565b565b6150fe906150f9616dee565b61510c565b565b61510990610f1c565b90565b6151269061512161511c82615100565b616e9d565b615187565b565b61513190610f00565b90565b61513d90615128565b90565b61514990615128565b90565b90565b9061516461515f61516b92615140565b61514c565b8254613cd1565b9055565b61517890610f00565b90565b6151849061516f565b90565b6151a361519b61519683615100565b615134565b61012e61514f565b6151cd7f4da1815cd654922275d14d2335fd9a0dd0aa6a0d0ff87fb4cc872ebe9704596f9161517b565b906151d66108cb565b806151e081610a26565b0390a2565b6151ee906150da565b565b6151f8613a7e565b5061522a61521a61520830613d68565b31615214610128610f63565b90613b0f565b6152246064613d77565b90613b50565b90565b61523e90615239616d06565b615240565b565b6152519061524c616dee565b615253565b565b6152659061526081616e9d565b615267565b565b6152738161012f613cef565b61529d7f4d03942c29c20d1bccfe551e9d148c917c5a44fb558a4fc60270d8f76fb75f5491610f28565b906152a66108cb565b806152b081610a26565b0390a2565b6152be9061522d565b565b6152c8613a7e565b506152fa6152ea6152d830614511565b316152e4610122610f63565b90613b0f565b6152f46064613d77565b90613b50565b90565b615305613a7e565b50615325615314610125610f63565b61531f610134610f63565b90613b50565b90565b615330613a7e565b5061534261533d5f613ad7565b614e99565b90565b61535690615351616d06565b615358565b565b61536990615364616dee565b61536b565b565b6153778161012861392f565b6153ad7fb5a05ec7911dd5450a7fa4ae54d50d9d83af6e256db8fc76c82edd7b659cf8bc916153a46108cb565b91829182610b12565b0390a1565b6153bb90615345565b565b6153ce906153c9616d06565b6153d0565b565b6153e1906153dc616dee565b6153ef565b565b6153ec90610f1c565b90565b615409906154046153ff826153e3565b616e9d565b61546a565b565b61541490610f00565b90565b6154209061540b565b90565b61542c9061540b565b90565b90565b9061544761544261544e92615423565b61542f565b8254613cd1565b9055565b61545b90610f00565b90565b61546790615452565b90565b61548661547e615479836153e3565b615417565b61012c615432565b6154b07fb4cecfe1346c94da27291cf5a02969d5fe0b5c36eca49b04fcd60841d28c5e139161545e565b906154b96108cb565b806154c381610a26565b0390a2565b6154d1906153bd565b565b6154db6154dd565b565b6154e5615668565b565b90565b6154fe6154f9615503926154e7565b610ecb565b6108dd565b90565b61551060036154ea565b90565b67ffffffffffffffff1690565b61553461552f615539926108dd565b610ecb565b615513565b90565b60401c90565b60ff1690565b6155546155599161553c565b615542565b90565b6155669054615548565b90565b67ffffffffffffffff1690565b61558261558791610f4a565b615569565b90565b6155949054615576565b90565b906155aa67ffffffffffffffff91613911565b9181191691161790565b6155c86155c36155cd92615513565b610ecb565b615513565b90565b90565b906155e86155e36155ef926155b4565b6155d0565b8254615597565b9055565b60401b90565b9061560d68ff0000000000000000916155f3565b9181191691161790565b615620906149aa565b90565b90565b9061563b61563661564292615617565b615623565b82546155f9565b9055565b61564f90615513565b9052565b9190615666905f60208501940190615646565b565b615678615673615506565b615520565b61568061958c565b61568b5f820161555c565b801561571b575b6156ff576156c4906156a6835f83016155d3565b6156b360015f8301615626565b6156bb615965565b5f809101615626565b6156fa7fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2916156f16108cb565b91829182615653565b0390a1565b5f63f92ee8a960e01b81528061571760048201610a26565b0390fd5b506157275f820161558a565b61573961573384615513565b91615513565b1015615692565b90565b61575761575261575c92615740565b610ecb565b6108dd565b90565b61576a610e10615743565b90565b61577c615782919392936108dd565b926108dd565b9161578e8382026108dd565b92818404149015171561579d57565b614a2e565b90565b6157b96157b46157be926157a2565b610ecb565b6108dd565b90565b6157cc6104b06157a5565b90565b6157db6157e1916108dd565b916108dd565b9081156157ec570490565b613b3c565b615800615806919392936108dd565b926108dd565b820180921161581157565b614a2e565b61586461585661583561582761575f565b61582f614f35565b9061576d565b6158506158406157c1565b61584a600261455d565b906157cf565b906157f1565b61585e6157c1565b906157cf565b90565b90565b61587e61587961588392615867565b610ecb565b6108dd565b90565b615890600d61586a565b90565b90565b6158aa6158a56158af92615893565b610ecb565b6108dd565b90565b6158d1906158cb6158c56158d6946108dd565b916108dd565b906141cf565b6108dd565b90565b6158f6623671796158f16158eb615886565b91615896565b6158b2565b90565b90565b61591061590b615915926158f9565b610ecb565b6108dd565b90565b61592260086158fc565b90565b90565b61593c61593761594192615925565b610ecb565b6108dd565b90565b615955670de0b6b3a7640000615928565b90565b61596260036154ea565b90565b615978615970615816565b61013461392f565b61598b6159836158d9565b61013561392f565b61599e615996615918565b61013661392f565b6159b16159a9615944565b61013761392f565b6159c46159bc615958565b61013861392f565b565b6159ce6154d3565b565b6159d8613a7e565b506159ea6159e55f613ad7565b6161f3565b90565b6159f5616d06565b6159fd6159ff565b565b615a10615a0b5f613aa1565b6195b0565b565b615a1a6159ed565b565b615a2d90615a28616d06565b615a2f565b565b615a4090615a3b616dee565b615a42565b565b615a4e8161013561392f565b615a847f169f25ec19cc5b518dc57adf05bb2d85155b1d6c60767e9f0e3ced18eac0ab7791615a7b6108cb565b91829182610b12565b0390a1565b615a9290615a1c565b565b615aa590615aa0616d06565b615aa7565b565b615ab890615ab3616dee565b615aba565b565b615ac68161012361392f565b615afc7fb5edd1f338b34c8f5dd3b1c5cc12f05653c495713c282bf588d34cf14fad0f8991615af36108cb565b91829182610b12565b0390a1565b615b0a90615a94565b565b615b1d90615b18616d06565b615b1f565b565b615b3090615b2b616dee565b615b32565b565b615b3e8161013861392f565b615b747f616bfcaa6490f55f6e57a4deedac1db04d0d6826deb84fad86cc43439bcf356491615b6b6108cb565b91829182610b12565b0390a1565b615b8290615b0c565b565b90615b9994939291615b94616f2c565b615ba3565b615ba1616f71565b565b91615bb3929493949190916199e4565b615bc6615bc161012c613e03565b61138e565b9063e2051c7e91615bd861010b610f63565b91615be161731d565b9490823b15615c57575f94615c148692615c0994615bfd6108cb565b998a9889978896613e14565b865260048601613e42565b03925af18015615c5257615c26575b50565b615c45905f3d8111615c4b575b615c3d8183610cd2565b810190613e1a565b5f615c23565b503d615c33565b613e80565b613e10565b90615c6994939291615b84565b565b615c7c90615c77616d06565b615c7e565b565b615c8f90615c8a616dee565b615c91565b565b615c9d8161011861392f565b615cd37f4e8c80fe79d13b8663de9f6981925ae24c498cc07b0ebd4f4789fa78ca17caff91615cca6108cb565b91829182610b12565b0390a1565b615ce190615c6b565b565b615cf490615cef616d06565b615cf6565b565b615d0790615d02616dee565b615d09565b565b615d158161013461392f565b615d4b7f7acba37d1b2d934e554139ff1d470d2cce50a7b6e56870577249b87564a95a6991615d426108cb565b91829182610b12565b0390a1565b615d5990615ce3565b565b615d63613846565b50615d765f615d70619f30565b01612b32565b90565b90615d8c9291615d87616f2c565b615d96565b615d94616f71565b565b91615da3929190916199e4565b565b90615db09291615d79565b565b615dc390615dbe616d06565b615dc5565b565b615dd690615dd1616dee565b615de4565b565b615de190610f1c565b90565b615dfe90615df9615df482615dd8565b616e9d565b615e5f565b565b615e0990610f00565b90565b615e1590615e00565b90565b615e2190615e00565b90565b90565b90615e3c615e37615e4392615e18565b615e24565b8254613cd1565b9055565b615e5090610f00565b90565b615e5c90615e47565b90565b615e7b615e73615e6e83615dd8565b615e0c565b610129615e27565b615ea57f9b3eda10f1724a2cf9f7dae4ac263c77908df4d00e92f1377b66fc8be37cd8c191615e53565b90615eae6108cb565b80615eb881610a26565b0390a2565b615ec690615db2565b565b615ed990615ed4616d06565b615edb565b565b615eec90615ee7616dee565b615efa565b565b615ef790610f1c565b90565b615f1490615f0f615f0a82615eee565b616e9d565b615f75565b565b615f1f90610f00565b90565b615f2b90615f16565b90565b615f3790615f16565b90565b90565b90615f52615f4d615f5992615f2e565b615f3a565b8254613cd1565b9055565b615f6690610f00565b90565b615f7290615f5d565b90565b615f91615f89615f8483615eee565b615f22565b61012a615f3d565b615fbb7fdab38e33e6e11cbb8b085bba9d7426d5e5af01bcc46d2c5957645e4d09e8c49c91615f69565b90615fc46108cb565b80615fce81610a26565b0390a2565b615fdc90615ec8565b565b615fe6613a7e565b50616018616008615ff630614511565b3161600261011e610f63565b90613b0f565b6160126064613d77565b90613b50565b90565b61602c90616027616d06565b61602e565b565b61603f9061603a616dee565b616041565b565b61604a90619f54565b565b6160559061601b565b565b61605f616f2c565b616067616071565b61606f616f71565b565b61607c61010b610f63565b61608461731d565b34916160ce6160bc6160b67fe32cacf203d00685e2b4d8b0a90e7cd8f3f8a208fdf116f4bb36abe08b7d548e93610ece565b93610f28565b936160c56108cb565b91829182610b12565b0390a3565b6160db616057565b565b6160ee906160e9616d06565b6160f0565b565b616101906160fc616dee565b61610f565b565b61610c90610f1c565b90565b6161299061612461611f82616103565b616e9d565b61618a565b565b61613490610f00565b90565b6161409061612b565b90565b61614c9061612b565b90565b90565b9061616761616261616e92616143565b61614f565b8254613cd1565b9055565b61617b90610f00565b90565b61618790616172565b90565b6161a661619e61619983616103565b616137565b61012d616152565b6161d07fbf6e296f85d08cc1ab124aed644bf4b19e4a726a7aea53e3784ab1341738a0409161617e565b906161d96108cb565b806161e381610a26565b0390a2565b6161f1906160dd565b565b61620e906161ff613a7e565b5061620981619f9b565b6191a2565b90565b616219613a7e565b5061622261388c565b5061622b61a058565b90616237610115610f63565b9190565b906162509493929161624b616f2c565b61625a565b616258616f71565b565b9161626a929493949190916199e4565b61627d61627861012c613e03565b61138e565b9063fe673fd39161628f61010b610f63565b9161629861731d565b9490823b1561630e575f946162cb86926162c0946162b46108cb565b998a9889978896613e14565b865260048601614792565b03925af18015616309576162dd575b50565b6162fc905f3d8111616302575b6162f48183610cd2565b810190613e1a565b5f6162da565b503d6162ea565b613e80565b613e10565b906163209493929161623b565b565b61632a613a7e565b5061633c6163375f613ad7565b613b72565b90565b61634e616354919392936108dd565b926108dd565b820391821161635f57565b614a2e565b61636c613846565b90616375613a7e565b9061637e613846565b90616387613a7e565b90616393610101612b32565b6163ad6163a76163a25f613aa1565b61093b565b9161093b565b036163b5575b565b93505050506163c5610105612b32565b906163d1610106610f63565b6163dc610107610f63565b926163e8610108610f63565b926163f4610109612b32565b9261640061010a610f63565b92616437600261643161642061010461641a61010b610f63565b90610eea565b61642b610101612b32565b90610f34565b01610f63565b9661644342899061633f565b928261645f6164596164545f613aa1565b61093b565b9161093b565b145f146164c85750505061648c9061648661647b610101612b32565b9791965b42926157f1565b9061633f565b61649581613890565b6164af6164a96164a486613890565b6109ae565b916109ae565b136164bb575b506163b3565b925090508391905f6164b5565b928098919792986164e16164db8a6108dd565b916108dd565b116164f6575b50509061648661648c9261647f565b969261651282999361650c6165189487906157f1565b926157f1565b9061633f565b9061652282613890565b61653c61653661653188613890565b6109ae565b916109ae565b13616561575b505061648c9094616486616557610101612b32565b97919691926164e7565b90945061648c9193509392905f616542565b6165849061657f616d06565b616586565b565b61659790616592616dee565b616599565b565b6165a58161012661392f565b6165db7f4636d3e567b27988879babd22e50f49104ca65647933fc1623ff3d4d807438d2916165d26108cb565b91829182610b12565b0390a1565b6165e990616573565b565b906165fe92916165f9616f2c565b616608565b616606616f71565b565b91616615929190916170c7565b565b9061662292916165eb565b565b61663590616630616d06565b616637565b565b61664890616643616dee565b61664a565b565b6166568161013061392f565b61668c7f2652e6657dd1ed89d7bdcb70f8827cc8449ac4536ebf604dbb2465cdad264020916166836108cb565b91829182610b12565b0390a1565b61669a90616624565b565b6166ad906166a8616d06565b6166af565b565b6166c0906166bb616dee565b6166c2565b565b6166cb9061a089565b565b6166d69061669c565b565b6166e9906166e4616d06565b6166eb565b565b6166fc906166f7616dee565b6166fe565b565b61670a8161011561392f565b6167407f4abea08c196329c357e3175d011af39a8625be99ef0ba5a0f3547a95534fedb7916167376108cb565b91829182610b12565b0390a1565b61674e906166d8565b565b616758616d06565b616760616762565b565b61676a61676c565b565b6167746173a9565b61677c6167d7565b565b5f7f546f6f206561726c792e00000000000000000000000000000000000000000000910152565b6167b2600a6020926128f4565b6167bb8161677e565b0190565b6167d49060208101905f8183039101526167a5565b90565b6168066167e2616b6b565b91906167ff6167f96167f48593613890565b6109ae565b916109ae565b13156149aa565b616900576168fe906168f96168f46168e4616822610110610f63565b6168de61686161685061684061683961010f610f63565b85906157cf565b61684a6001614589565b906157f1565b9261685b600261455d565b9061576d565b956168d86168d26168c16168b061689661688661687f61010f610f63565b8d906157cf565b6168906001614589565b906157f1565b9661689f613a7e565b506168ab61010f610f63565b61633f565b6168bb610125610f63565b9061576d565b946168cd61010f610f63565b61633f565b91613af3565b9061576d565b906157cf565b6168ee6001614589565b906157f1565b61a0d0565b61a089565b565b6169086108cb565b63a29f5c4d60e01b81528061691f600482016167bf565b0390fd5b61692b616750565b565b61693e90616939616d06565b616940565b565b6169519061694c616dee565b616953565b565b61695f8161011c61392f565b6169957fd95e7f967f9370c11deb15ffbb191b9f2e9795ab0738db5bc72bd2794978f32d9161698c6108cb565b91829182610b12565b0390a1565b6169a39061692d565b565b6169b0905f036109ae565b90565b6169bb61388c565b506169cc6169c76138ba565b6169a5565b90565b6169e0906169db616d06565b6169e2565b565b6169f3906169ee616dee565b6169f5565b565b616a018161012761392f565b616a377f37a332914fac995349420c0419b4423a19dcb762017f691442a0782ce4bf417a91616a2e6108cb565b91829182610b12565b0390a1565b616a45906169cf565b565b616a5890616a53616d06565b616a5a565b565b80616a75616a6f616a6a5f613aa1565b61093b565b9161093b565b14616a8557616a83906195b0565b565b616aa8616a915f613aa1565b5f918291631e4fbdf760e01b835260048301610954565b0390fd5b616ab590616a47565b565b616ac890616ac3616d06565b616aca565b565b616adb90616ad6616dee565b616add565b565b616ae69061a0d0565b565b616af190616ab7565b565b616b0490616aff616d06565b616b06565b565b616b1790616b12616dee565b616b19565b565b616b258161011261392f565b616b5b7fdeb71e1d117914859ddde484a9810084d9ee399077d5cd8dcbdf8770d33d8ae491616b526108cb565b91829182610b12565b0390a1565b616b6990616af3565b565b616b73613a7e565b50616b7c61388c565b50616b8561a117565b90616b8e6138ba565b90565b90565b5f616bb5616bb0616bbb93616ba7613a7e565b50610103611563565b616b91565b01610f63565b90565b90565b616bf091616be6616beb92616bd4613a7e565b50616bdd613a7e565b50610104610eea565b610f34565b616bbe565b90616c086001616c015f8501610f63565b9301610f63565b90565b616c1c90616c17616d06565b616c1e565b565b616c2f90616c2a616dee565b616c31565b565b616c3d8161011e61392f565b616c737fbfcd8fb930a57c1598c9760db19c84ec766546c3c9a8565611df8302482bfb1791616c6a6108cb565b91829182610b12565b0390a1565b616c8190616c0b565b565b616c8b616f2c565b616c93616cd3565b616c9b616f71565b565b90565b616cb4616caf616cb992616c9d565b610ecb565b6109ae565b90565b616cc55f612878565b90565b616cd0616cbc565b90565b616cfa5f19616ce25f91616ca0565b90616cf4616cee616cc8565b91613abb565b916199e4565b565b616d04616c83565b565b616d0e615d5b565b616d27616d21616d1c61731d565b61093b565b9161093b565b03616d2e57565b616d50616d3961731d565b5f91829163118cdaa760e01b835260048301610954565b0390fd5b60207f65616479206163746976652e0000000000000000000000000000000000000000917f5468652063757272656e742062696464696e6720726f756e6420697320616c725f8201520152565b616dae602c6040926128f4565b616db781616d54565b0190565b916040616dec929493616de5616dda606083018381035f850152616da1565b966020830190610b05565b0190610b05565b565b616df961010d610f63565b616e1642616e0f616e09846108dd565b916108dd565b10156149aa565b616e1d5750565b4290616e40616e2a6108cb565b92839263d0fd11df60e01b845260048401616dbb565b0390fd5b5f7f5468652070726f76696465642061646472657373206973207a65726f2e000000910152565b616e78601d6020926128f4565b616e8181616e44565b0190565b616e9a9060208101905f818303910152616e6b565b90565b616eb7616eb1616eac5f613aa1565b61093b565b9161093b565b14616ebe57565b616ec66108cb565b63eac0d38960e01b815280616edd60048201616e85565b0390fd5b90565b616ef8616ef3616efd92616ee1565b613911565b611bfa565b90565b616f297f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00616ee4565b90565b616f3461a146565b616f5557616f53616f4b616f46616f00565b61a17a565b60019061a18f565b565b5f633ee5aeb560e01b815280616f6d60048201610a26565b0390fd5b616f8b616f84616f7f616f00565b61a17a565b5f9061a18f565b565b60407f642e000000000000000000000000000000000000000000000000000000000000917f5468652063757272656e742043535420626964207072696365206973206772655f8201527f61746572207468616e20746865206d6178696d756d20796f7520616c6c6f776560208201520152565b61700d60426060926128f4565b61701681616f8d565b0190565b91604061704b929493617044617039606083018381035f850152617000565b966020830190610b05565b0190610b05565b565b617057600261455d565b90565b61706390616ca0565b9052565b91946170b46170a96170be9360a09661709c6170c59a9c9b999c61709260c08a01945f8b019061705a565b60208901906109b1565b8682036040880152612908565b986060850190610b05565b6080830190610b05565b0190610b05565b565b6170d86170d35f613ad7565b613b72565b926170f7846170ef6170e9846108dd565b916108dd565b1015156149aa565b6172fe575061710d6171085f613ad7565b6161f3565b9061712c8261712461711e846108dd565b916108dd565b1115156149aa565b6172d8575061713c83829061a382565b6171838161717d600161716d61715f61010461715961010b610f63565b90610eea565b61716761731d565b90610f34565b019161717883610f63565b6157f1565b9061392f565b61718f4261011461392f565b6171b56171a48261719e61704d565b9061576d565b6171af610118610f63565b9061a658565b6171c18161011661392f565b6171cc610102612b32565b6171e66171e06171db5f613aa1565b61093b565b9161093b565b146172c6575b506172006171f861731d565b610102613cef565b61722961720e610115610f63565b6172238161721d610133610f63565b906157cf565b906157f1565b6172358161011561392f565b61723e8361a7e4565b61724961010b610f63565b9061725261731d565b926172c16172615f1992613890565b925f19969790617272610124610f63565b916172af6172a96172a37f1d1f406c89e99504a7222aaba50bf96b9f91f522ebc6ea825255145919e102ec99610ece565b99610f28565b99616ca0565b996172b86108cb565b96879687617067565b0390a4565b6172d29061011761392f565b5f6171ec565b906172fa6172e46108cb565b92839263814ac7ff60e01b84526004840161701a565b0390fd5b836173195f9283926335465b3160e01b8452600484016137b5565b0390fd5b617325613846565b503390565b60207f207468652063757272656e742062696464696e6720726f756e642e0000000000917f41206269642068617320616c7265616479206265656e20706c6163656420696e5f8201520152565b617384603b6040926128f4565b61738d8161732a565b0190565b6173a69060208101905f818303910152617377565b90565b6173d86173b7610101612b32565b6173d16173cb6173c65f613aa1565b61093b565b9161093b565b14156149aa565b6173de57565b6173e66108cb565b634283f4b960e01b8152806173fd60048201617391565b0390fd5b61740d8161010d61392f565b6174437f9a2159c1f277ddd727551baedc6a6c4cba77cc5219c8563ee3b15fb67548d89b9161743a6108cb565b91829182610b12565b0390a1565b61747e600261747861746761010461746161010b610f63565b90610eea565b617472610101612b32565b90610f34565b01610f63565b61748942829061633f565b617494610105612b32565b6174ae6174a86174a35f613aa1565b61093b565b9161093b565b145f146174e5576174da6174e2926174d26174ca610101612b32565b610105613cef565b61010661392f565b61010761392f565b5b565b806175026174fc6174f7610107610f63565b6108dd565b916108dd565b1161750f575b50506174e3565b6175678261753361752e61756f95617528610107610f63565b906157f1565b617576565b617549617541610107610f63565b61010861392f565b61755f617557610101612b32565b610105613cef565b61010661392f565b61010761392f565b5f80617508565b61759f90617599617588610106610f63565b617593610108610f63565b906157f1565b9061633f565b6175a881613890565b6175cc6175c66175c16175bc61010a610f63565b613890565b6109ae565b916109ae565b136175d5575b50565b6175f7906175ef6175e7610105612b32565b610109613cef565b61010a61392f565b5f6175d2565b6176076020610cfb565b90565b5f90565b6176166175fd565b9060208261762261760a565b81525050565b61763061760e565b90565b9061763d906108dd565b9052565b67ffffffffffffffff81116176595760208091020190565b610cbe565b9061767061766b83617641565b610cfb565b918252565b61767f6040610cfb565b90565b5f90565b61768e617675565b906020808361769b617682565b8152016176a661760a565b81525050565b6176b4617686565b90565b5f5b8281106176c557505050565b6020906176d06176ac565b81840152016176b9565b906176ff6176e78361765e565b926020806176f58693617641565b92019103906176b7565b565b5190565b9061770f82617701565b811015617720576020809102010190565b612ad5565b9061772f9061093b565b9052565b61773c906108dd565b5f811461774a576001900390565b614a2e565b61775b617761916108dd565b916108dd565b90811561776c570690565b613b3c565b9050519061777e826108e0565b565b9060208282031261779957617796915f01617771565b90565b6108d5565b60209181520190565b60200190565b6177b69061093b565b9052565b6177c3906108dd565b9052565b906020806177e9936177df5f8201515f8601906177ad565b01519101906177ba565b565b906177f8816040936177c7565b0190565b60200190565b9061781f61781961781284617701565b809361779e565b926177a7565b905f5b81811061782f5750505090565b90919261784861784260019286516177eb565b946177fc565b9101919091617822565b61787661788394929361786c60608401955f850190610b05565b6020830190610947565b6040818403910152617802565b90565b61789261789791610f4a565b612766565b90565b6178a49054617886565b90565b5f9060033d116178b4575b565b905060045f803e6178c55f516108c5565b906178b2565b5f5f9160233d116178d9575b565b915050602060045f3e6001905f51906178d7565b90565b6179046178ff617909926178ed565b610ecb565b6108dd565b90565b61791660126178f0565b90565b905090565b6179295f8092617919565b0190565b6179369061791e565b90565b9061794b61794683611aa7565b610cfb565b918252565b606090565b3d5f14617970576179653d617939565b903d5f602084013e5b565b617978617950565b9061796e565b5f7f455448207472616e7366657220746f2063686172697479206661696c65642e00910152565b6179b2601f6020926128f4565b6179bb8161797e565b0190565b91906179e29060206179da604086018681035f8801526179a5565b940190610b05565b565b60207f696e207072697a652062656e6566696369617279206661696c65642e00000000917f455448207472616e7366657220746f2062696464696e6720726f756e64206d615f8201520152565b617a3e603c6040926128f4565b617a47816179e4565b0190565b916040617a7c929493617a75617a6a606083018381035f850152617a31565b966020830190610947565b0190610b05565b565b60ff1690565b617a98617a93617a9d926154e7565b610ecb565b617a7e565b90565b90565b617ab7617ab2617abc92617aa0565b610ecb565b617a7e565b90565b617ad3617ace617ad892617a7e565b610ecb565b6108dd565b90565b617ae7617aec91610f4a565b611459565b90565b617af99054617adb565b90565b617b0690516108dd565b90565b90565b617b20617b1b617b2592617b09565b610ecb565b6108dd565b90565b67ffffffffffffffff8111617b405760208091020190565b610cbe565b90505190617b5282610b91565b565b90929192617b69617b6482617b28565b610cfb565b9381855260208086019202830192818411617ba657915b838310617b8d5750505050565b60208091617b9b8486617b45565b815201920191617b80565b611183565b9080601f83011215617bc957816020617bc693519101617b54565b90565b610cac565b90602082820312617bfe575f82015167ffffffffffffffff8111617bf957617bf69201617bab565b90565b6108d9565b6108d5565b5190565b90617c19617c1483617b28565b610cfb565b918252565b369037565b90617c48617c3083617c07565b92602080617c3e8693617b28565b9201910390617c1e565b565b67ffffffffffffffff8111617c625760208091020190565b610cbe565b90617c79617c7483617c4a565b610cfb565b918252565b617c886040610cfb565b90565b617c93617c7e565b9060208083617ca0617682565b815201617cab61760a565b81525050565b617cb9617c8b565b90565b5f5b828110617cca57505050565b602090617cd5617cb1565b8184015201617cbe565b90617d04617cec83617c67565b92602080617cfa8693617c4a565b9201910390617cbc565b565b5190565b90617d1482617d06565b811015617d25576020809102010190565b612ad5565b90617d3482617c03565b811015617d45576020809102010190565b612ad5565b617d54905161093b565b90565b617d63617d6891610f4a565b6136dc565b90565b617d759054617d57565b90565b60209181520190565b60200190565b90602080617da993617d9f5f8201515f8601906177ad565b01519101906177ba565b565b90617db881604093617d87565b0190565b60200190565b90617ddf617dd9617dd284617d06565b8093617d78565b92617d81565b905f5b818110617def5750505090565b909192617e08617e026001928651617dab565b94617dbc565b9101919091617de2565b617e279160208201915f818403910152617dc2565b90565b617e36617e3b91610f4a565b611944565b90565b617e489054617e2a565b90565b90565b617e62617e5d617e6792617e4b565b610ecb565b6108dd565b90565b60209181520190565b60200190565b90617e86816020936177ad565b0190565b60200190565b90617ead617ea7617ea084617c03565b8093617e6a565b92617e73565b905f5b818110617ebd5750505090565b909192617ed6617ed06001928651617e79565b94617e8a565b9101919091617eb0565b939290617f0b604091617f1394617efe60608901925f8a0190610b05565b8782036020890152617e90565b940190610b05565b565b617f4a617f5194617f40606094989795617f36608086019a5f870190610b05565b6020850190610b05565b6040830190610b05565b0190610b05565b565b617f5c906149aa565b9052565b604090617f89617f909496959396617f7f60608401985f850190617f53565b6020830190610b05565b0190610b05565b565b617f9a617628565b90617fae617fa661aa8b565b5f8401617633565b617fcd617fc8610103617fc261010b610f63565b90611563565b616b91565b90617fd6613a7e565b50617fdf6151f0565b91617fe861451d565b617ff0613d93565b92617ff96152c0565b90618002615fde565b9561800e61011f610f63565b9361802b618026866180206001614589565b906157f1565b6176da565b9661808e61807c61807461803e5f613abb565b61806d61804c8d8c90617705565b5161806361805b610109612b32565b5f8301617725565b6020889101617633565b85906157f1565b9a88906157cf565b996180888b899061576d565b906157f1565b60015b15618166575b5f966180a290617733565b9689886180ae91617705565b518c6180b99061ab7f565b8a600101908b5f016180ca90610f63565b6180d39161774f565b6180dc9161384a565b6180e590612b32565b9081815f01906180f491617725565b8c906020019061810391617633565b61810e61010b610f63565b8991908d927f9c62e2cb8cbd10bf3b8a5760977d719fc3475fa67cb7ef9d2e1701f275e885c49161813e90610ece565b9261814890610f28565b936181516108cb565b91829161815e91836137b5565b0390a3618091565b866181796181735f613abb565b916108dd565b1161809757909398506020919499979296955061819f61819a61012c613e03565b61138e565b906181d86387565d14919291926181b761010b610f63565b936181e36181c361731d565b976181cc6108cb565b98899788968795613e14565b855260048501617852565b03925af1908115618f28575f91618efa575b509661820a61820561012e61789a565b6127a7565b9063b6b55f25909190919061822061010b610f63565b90803b15618ef5576182455f93618250956182396108cb565b96879586948593613e14565b835260048301610b12565b03925af19081618ec9575b50155f14618ec457600161826d6178a7565b634e487b7114618e87575b618e82575b5f8061828a610131612b32565b836182936108cb565b908161829e8161792d565b03925af16182aa617955565b505f14618e30576182bc610131612b32565b6182fb6182e97f1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d92610f28565b926182f26108cb565b91829182610b12565b0390a25b6183335f8061830c61731d565b866183156108cb565b90816183208161792d565b03925af161832c617955565b50156149aa565b618e0257618342610102612b32565b61835c6183566183515f613aa1565b61093b565b9161093b565b14155f14618ded5761838b61837a6183746004617aa3565b5b617abf565b618385610120610f63565b906157f1565b9261839f61839a61012d617aef565b61149a565b5f63e36aee78916183b1610121610f63565b906184026183c0848b01617afc565b6183e97f7c6eeb003d4a6dc5ebf549935c6ffb814ba1f060f1af8a0b11c2aa94a8e716e4617b0c565b189461840d6183f66108cb565b96879586948594613e14565b8452600484016137b5565b03915afa8015618de857618433915f91618dc6575b509461842d86617c03565b906157f1565b9561846461845f88618459618449610138610f63565b6184536001614589565b9061633f565b906157f1565b617c23565b976184c361848461847f8a6184796001614589565b906157f1565b617cdf565b986184be6184938b8390617d0a565b516184aa6184a261012f612b32565b5f8301617725565b60206184b7610130610f63565b9101617633565b617733565b966184fa6184d28a8a90617d0a565b516184e66184de61731d565b5f8301617725565b60206184f361011c610f63565b9101617633565b61850c6185068b617c03565b5b617733565b61852961851761731d565b6185248d91849092617d2a565b617725565b8061853c6185368b6108dd565b916108dd565b111561854b5761850c90618507565b5061855587617c03565b5b806185696185635f613abb565b916108dd565b11156185e05761857890617733565b956185da61859861859261858d8b8b90617d2a565b617d4a565b9a617733565b996185c86185a78d8d90617d0a565b516185b4835f8301617725565b60206185c161011c610f63565b9101617633565b6185d58d918c9092617d2a565b617725565b95618556565b5092979690939895919461867f6185f8610120610f63565b5b61867a61863861863261862d61860e8a61ab7f565b61862760018901916186215f8b01610f63565b9061774f565b9061384a565b612b32565b95617733565b946186686186478d8890617d0a565b51618654835f8301617725565b602061866161011c610f63565b9101617633565b6186758b91879092617d2a565b617725565b617733565b908161869361868d5f613abb565b916108dd565b11156186b4579061867a61863861863261862d61867f9493505050506185f9565b50506186c261872091617733565b6186fb6186d0898390617d0a565b516186e76186df610109612b32565b5f8301617725565b60206186f461011c610f63565b9101617633565b61871b618709610109612b32565b6187168991849092617d2a565b617725565b617733565b61875961872e888390617d0a565b5161874561873d610105612b32565b5f8301617725565b602061875261011c610f63565b9101617633565b618779618767610105612b32565b6187748891849092617d2a565b617725565b61878b6187855f613abb565b916108dd565b115f14618dc1576187d36187a8876187a25f613abb565b90617d0a565b516187bf6187b7610102612b32565b5f8301617725565b60206187cc61011c610f63565b9101617633565b6187f96187e1610102612b32565b6187f4876187ee5f613abb565b90617d2a565b617725565b5b61880d618808610129617d6b565b61371d565b63b33266da87823b15618dbc57618843926188385f809461882c6108cb565b96879586948593613e14565b835260048301617e12565b03925af18015618db757618d8b575b50602061886861886361012b617e3e565b611985565b636578f113906188cb5f61887d61010b610f63565b936188d661888d838d9a01617afc565b6188b67f2a8612ecb5cb17da87f8befda0480288e2d053de55d9d7d4dc4899077cf5aeda617e4e565b186188bf6108cb565b998a9788968795613e14565b855260048501617ee0565b03925af18015618d86576189bc925f91618d58575b50979496979261891361890e61890089617d06565b618908613a7e565b50617733565b617733565b9761891f888a90617d0a565b51998961892d8782906157f1565b809c6189b461896561894061010b610f63565b94618960618959602061895161731d565b999601617afc565b9598617c03565b61633f565b966189a261899c6189967f9314e78538382a9fc6cd54cee107c8a7721d172c43b2e48c2e7a51f200fc378897610ece565b97610f28565b97610ece565b976189ab6108cb565b94859485617f15565b0390a4617c03565b945b856189d16189cb5f613abb565b916108dd565b1115618a89576189e090617733565b926189ec858590617d0a565b5195618a0b618a056189ff5f8a01617d4a565b92617733565b98617733565b96618a1761010b610f63565b90600191618a2960208c959301617afc565b938a93618a7d618a6b618a65618a5f7f27c21fe4cea1a3367aa491829dd4dd824296c00910626150464cba8ea5ebb3f496610ece565b96610f28565b96610ece565b96618a746108cb565b93849384617f60565b0390a4929594956189be565b9195909450618aa3618a9c610120610f63565b925b617733565b618aae858290617d0a565b5192618acd618ac7618ac15f8701617d4a565b92617733565b93617733565b93618ad961010b610f63565b905f91618aea602087959301617afc565b938793618b3e618b2c618b26618b207f27c21fe4cea1a3367aa491829dd4dd824296c00910626150464cba8ea5ebb3f496610ece565b96610f28565b96610ece565b96618b356108cb565b93849384617f60565b0390a481618b54618b4e5f613abb565b916108dd565b1115618b6657618aa390929192618a9e565b618c18929150618b7590617733565b90618b8b618b84868490617d0a565b5191617733565b93618b9761010b610f63565b90618ba361011f610f63565b91618bbc6020618bb45f8701617d4a565b939501617afc565b938793618c10618bfe618bf8618bf27faa858ae20a26d00a9ea528972d537e68a51a0744226d9ea1fc9b22492dc282a596610ece565b96610f28565b96610ece565b96618c076108cb565b93849384610fac565b0390a4617733565b90618c2e618c27848490617d0a565b5191617733565b90618c3a61010b610f63565b90618c526020618c4b5f8401617d4a565b9201617afc565b9291618ca5618c93618c8d618c877f838ec9dd2530548892bff113f5ffb0138d2efc63c7f59bb6571e8c923b63126094610ece565b94610f28565b94610ece565b94618c9c6108cb565b91829182610b12565b0390a4618cba618cb45f613abb565b916108dd565b115f14618d5157618cd490618cce5f613abb565b90617d0a565b51618ce061010b610f63565b90618cf86020618cf15f8401617d4a565b9201617afc565b9291618d4b618d39618d33618d2d7f3901b6430c99dc290ee88ff84c4de6091ad7eac335b58e92ef5cbb0793abf4f694610ece565b94610f28565b94610ece565b94618d426108cb565b91829182610b12565b0390a45b565b5050618d4f565b618d79915060203d8111618d7f575b618d718183610cd2565b810190617780565b5f6188eb565b503d618d67565b613e80565b618daa905f3d8111618db0575b618da28183610cd2565b810190613e1a565b5f618852565b503d618d98565b613e80565b613e10565b6187fa565b618de291503d805f833e618dda8183610cd2565b810190617bce565b5f618422565b613e80565b61838b61837a618dfd6003617a84565b618375565b82618e0b61731d565b618e2c618e166108cb565b928392630aa7db6360e11b845260048401617a4b565b0390fd5b618e3b610131612b32565b618e7a618e687f154fb6c686c977314af35b730a16b83facda5265b2abec7237c730f63e42215a92610f28565b92618e716108cb565b918291826179bf565b0390a26182ff565b613e80565b618e8f6178cb565b90618e9b575b50618278565b90505f9080618eb9618eb3618eae61790c565b6108dd565b916108dd565b0315618e955761abbb565b61827d565b618ee8905f3d8111618eee575b618ee08183610cd2565b810190613e1a565b5f61825b565b503d618ed6565b613e10565b618f1b915060203d8111618f21575b618f138183610cd2565b810190617780565b5f6181f5565b503d618f09565b613e80565b6001618f3991016108dd565b90565b618f50618f485f613aa1565b610101613cef565b618f64618f5c5f613aa1565b610102613cef565b618f78618f705f613aa1565b610105613cef565b618f8c618f845f613abb565b61010861392f565b618fa0618f985f613aa1565b610109613cef565b618fbd618fb5618fb05f19616ca0565b613af3565b61010a61392f565b618fdb618fd3618fce61010b610f63565b618f2d565b61010b61392f565b619016619011618fec610125610f63565b61900b618ffa610125610f63565b619005610126610f63565b90613b50565b906145b3565b619f54565b61903361902e4261902861010c610f63565b906145b3565b617401565b565b61903d613a7e565b50619046613a7e565b50619052610101612b32565b61906c6190666190615f613aa1565b61093b565b9161093b565b145f1461914a5761908f61908161010f610f63565b9161908a6138ba565b613aad565b806190aa6190a461909f5f613ad7565b6109ae565b916109ae565b13155f146190b857505b5b90565b906190e16190d1826190cb610110610f63565b90613b50565b6190db6001614589565b906145b3565b916190ea61a117565b906190f481613af3565b619106619100846108dd565b916108dd565b105f1461914257906191316191369261912b61912561913c97876145a5565b91613af3565b90613b0f565b613b50565b906145a5565b5b6190b4565b50505061913d565b50619156610111610f63565b6190b5565b9061917161916b619178936108dd565b916108dd565b900a6108dd565b90565b61919a9061919461918e61919f946108dd565b916108dd565b90610c3a565b6108dd565b90565b9190916191ad613a7e565b50916191ba610101612b32565b6191d46191ce6191c95f613aa1565b61093b565b9161093b565b036191dd575b50565b6191f66191e86152fd565b916191f16148e7565b6138ac565b9061920a61920382613890565b83906138ac565b918261922661922061921b5f613ad7565b6109ae565b916109ae565b13619233575b50506191da565b6192b16192a061928f61927e6192e09798966192da966192b89661926761926161925c5f613ad7565b6109ae565b916109ae565b125f146192e9576192789150613890565b5b613af3565b619289610135610f63565b90613b0f565b61929a610125610f63565b90613b50565b6192ab610136610f63565b9061915b565b8390613b0f565b6192d46192c6610136610f63565b6192ce615886565b90613b0f565b9061917b565b906145b3565b905f808061922c565b50619279565b6192f890610f1c565b90565b619304306192ef565b6193366193307f000000000000000000000000000000000000000000000000000000000000000061093b565b9161093b565b148015619360575b61934457565b5f63703e46dd60e11b81528061935c60048201610a26565b0390fd5b5061936961abcb565b61939b6193957f000000000000000000000000000000000000000000000000000000000000000061093b565b9161093b565b141561933e565b6193b3906193ae616d06565b6193b5565b565b506193be616dee565b565b6193c9906193a2565b565b6193d490610f00565b90565b6193e0906193cb565b90565b6193ec90610f1c565b90565b6193f881611bfa565b036193ff57565b5f80fd5b90505190619410826193ef565b565b9060208282031261942b57619428915f01619403565b90565b6108d5565b919061945e6020619448619443866193d7565b6193e3565b6352d1902d906194566108cb565b938492613e14565b8252818061946e60048201610a26565b03915afa80915f926194fe575b50155f146194af57505090600161949057505b565b6194ab905f918291634c9c8ce360e01b835260048301610954565b0390fd5b92836194ca6194c46194bf615017565b611bfa565b91611bfa565b036194df576194da92935061abf1565b61948e565b6194fa845f918291632a87526960e21b835260048301611c0a565b0390fd5b61952091925060203d8111619527575b6195188183610cd2565b810190619412565b905f61947b565b503d61950e565b619537306192ef565b6195696195637f000000000000000000000000000000000000000000000000000000000000000061093b565b9161093b565b0361957057565b5f63703e46dd60e11b81528061958860048201610a26565b0390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0090565b6195b8619f30565b6195d06195c65f8301612b32565b915f849101613cef565b906196046195fe7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093610f28565b91610f28565b9161960d6108cb565b8061961781610a26565b0390a3565b61962b619631919392936109ae565b926109ae565b91828103925f82851281831216928513911215161761964c57565b614a2e565b60407f727265642e000000000000000000000000000000000000000000000000000000917f5468652063757272656e742045544820626964207072696365206973206772655f8201527f61746572207468616e2074686520616d6f756e7420796f75207472616e73666560208201520152565b6196d160456060926128f4565b6196da81619651565b0190565b91604061970f9294936197086196fd606083018381035f8501526196c4565b966020830190610b05565b0190610b05565b565b60207f206265656e207573656420666f722062696464696e672e000000000000000000917f546869732052616e646f6d2057616c6b204e46542068617320616c72656164795f8201520152565b61976b60376040926128f4565b61977481619711565b0190565b919061979b906020619793604086018681035f88015261975e565b940190610b05565b565b156197a55750565b6197c7906197b16108cb565b91829163c35947c560e01b835260048301619778565b0390fd5b6197d76197dc91610f4a565b6121b5565b90565b6197e990546197cb565b90565b9060208282031261980557619802915f01617b45565b90565b6108d5565b60207f6e646f6d2057616c6b204e46542e000000000000000000000000000000000000917f596f7520617265206e6f7420746865206f776e6572206f6620746869732052615f8201520152565b619864602e6040926128f4565b61986d8161980a565b0190565b6060906198aa6198b194969593966198a0619895608085018581035f870152619857565b986020850190612202565b6040830190610b05565b0190610947565b565b92909192156198c157505050565b6198e3906198cd6108cb565b938493630b81342760e31b855260048501619871565b0390fd5b6198f1600261455d565b90565b61990861990361990d926109ae565b610ecb565b6109ae565b90565b919461995d6199526199679360a09661994561996e9a9c9b999c61993b60c08a01945f8b01906109b1565b602089019061705a565b8682036040880152612908565b986060850190610b05565b6080830190610b05565b0190610b05565b565b5f7f45544820726566756e64207472616e73666572206661696c65642e0000000000910152565b6199a4601b6020926128f4565b6199ad81619970565b0190565b9160406199e29294936199db6199d0606083018381035f850152619997565b966020830190610947565b0190610b05565b565b91906199f76199f25f613ad7565b613b72565b91619a1683619a0e619a08846108dd565b916108dd565b1015156149aa565b619f115750619a2c619a275f613ad7565b614e99565b9280619a48619a42619a3d5f613ad7565b6109ae565b916109ae565b125f14619f0357835b90619a6d619a5e34613890565b619a6784613890565b9061961c565b9485619a89619a83619a7e5f613ad7565b6109ae565b916109ae565b145f14619e68575b81619aac619aa6619aa15f613ad7565b6109ae565b916109ae565b125f14619d1c57619b5a619b4a619b62925b619b0486619afe5f619aee619ae0610104619ada61010b610f63565b90610eea565b619ae861731d565b90610f34565b0191619af983610f63565b6157f1565b9061392f565b619b0f610101612b32565b619b29619b23619b1e5f613aa1565b61093b565b9161093b565b14619cfa575b619b4481619b3e610112610f63565b906157cf565b906157f1565b619b546001614589565b906157f1565b61011161392f565b619bbb619b97619b86619b76610115610f63565b619b806001614589565b906157f1565b619b91610133610f63565b9061576d565b619bb5619ba5610133610f63565b619baf6001614589565b906157f1565b906157cf565b90619bc88261011561392f565b619bd18561ac7a565b619bda8461a7e4565b619be561010b610f63565b91619c5b619bfa619bf461731d565b95613890565b915f1993969790619c0c610124610f63565b91619c49619c43619c3d7f1d1f406c89e99504a7222aaba50bf96b9f91f522ebc6ea825255145919e102ec99610ece565b99610f28565b996198f4565b99619c526108cb565b96879687619910565b0390a480619c79619c73619c6e5f613ad7565b6109ae565b916109ae565b13619c82575b50565b619cbe5f80619c8f61731d565b619c9885613af3565b619ca06108cb565b9081619cab8161792d565b03925af1619cb7617955565b50156149aa565b15619c7f57619cd4619cce61731d565b91613af3565b90619cf6619ce06108cb565b928392630aa7db6360e11b8452600484016199b1565b0390fd5b619d17619d0f82619d096198e7565b9061576d565b61010f61392f565b619b2f565b619db290619d62619d40619d3b610119619d3587613af3565b906133f9565b610f63565b619d52619d4c5f613abb565b916108dd565b14619d5c85613af3565b9061979d565b619d6a61731d565b906020619d80619d7b61012a6197df565b6121f6565b636352211e90619da7619d9288613af3565b92619d9b6108cb565b97889485938493613e14565b835260048301610b12565b03915afa908115619e6357619e09619b5a93619de5619ddf619b4a95619b62985f91619e35575b5061093b565b9161093b565b14619df161012a6197df565b619dfa88613af3565b90619e0361731d565b926198b3565b619e30619e166001614589565b619e2b610119619e2589613af3565b906133f9565b61392f565b619abe565b619e56915060203d8111619e5c575b619e4e8183610cd2565b8101906197ec565b5f619dd9565b503d619e44565b613e80565b85619e83619e7d619e785f613ad7565b6109ae565b916109ae565b135f14619edb57619e9f619e98610113610f63565b3a9061576d565b619eba619eb4619eae89613af3565b926108dd565b916108dd565b1115619ec6575b619a91565b91509350619ed35f613ad7565b933491619ec1565b823490619eff619ee96108cb565b92839263814ac7ff60e01b8452600484016196de565b0390fd5b619f0c846145c1565b619a51565b82619f2c5f9283926335465b3160e01b8452600484016137b5565b0390fd5b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930090565b619f608161012561392f565b619f967f07417920574ce0bdfe987af0575c8793cc73a29d7830760ad459d0e569b5b79b91619f8d6108cb565b91829182610b12565b0390a1565b619fb590619fa7613a7e565b50619fb061aeae565b6138ac565b80619fd0619fca619fc55f613ad7565b6109ae565b916109ae565b131561a04b5761a02461a03591619fe8610102612b32565b61a002619ffc619ff75f613aa1565b61093b565b9161093b565b145f1461a0385761a01e61a017610117610f63565b5b91613af3565b90613b0f565b61a02f610115610f63565b90613b50565b90565b61a01e61a046610116610f63565b61a018565b5061a0555f613abb565b90565b61a06061388c565b5061a08661a06d42613890565b61a08061a07b610114610f63565b613890565b906138ac565b90565b61a0958161011061392f565b61a0cb7fb6f6af60099e44041a78b3561ed029b98bf03fdb0efbbb2eb15e1f3d7d9230379161a0c26108cb565b91829182610b12565b0390a1565b61a0dc8161010e61392f565b61a1127ffdf6043c89a5f304289122dcc9f8bd78bb111b5d4f409e2fc2e6c141a1110b799161a1096108cb565b91829182610b12565b0390a1565b61a11f613a7e565b5061a13f61a12e610125610f63565b61a13961010e610f63565b90613b50565b90565b5f90565b61a14e61a142565b5061a16761a16261a15d616f00565b61a17a565b61aedd565b90565b5f90565b61a17790611bfa565b90565b61a18c9061a18661a16a565b5061a16e565b90565b5d565b91602061a1b392949361a1ac60408201965f830190610947565b0190610b05565b565b61a1c961a1c461a1ce9261455a565b610ecb565b617a7e565b90565b67ffffffffffffffff811161a1e95760208091020190565b610cbe565b9061a20061a1fb8361a1d1565b610cfb565b918252565b61a20f6040610cfb565b90565b5f90565b61a21e61a205565b906020808361a22b617682565b81520161a23661a212565b81525050565b61a24461a216565b90565b5f5b82811061a25557505050565b60209061a26061a23c565b818401520161a249565b9061a28f61a2778361a1ee565b9260208061a285869361a1d1565b920191039061a247565b565b5190565b9061a29f8261a291565b81101561a2b0576020809102010190565b612ad5565b9061a2bf906109ae565b9052565b60209181520190565b60200190565b61a2db906109ae565b9052565b9060208061a3019361a2f75f8201515f8601906177ad565b015191019061a2d2565b565b9061a3108160409361a2df565b0190565b60200190565b9061a33761a33161a32a8461a291565b809361a2c3565b9261a2cc565b905f5b81811061a3475750505090565b90919261a36061a35a600192865161a303565b9461a314565b910191909161a33a565b61a37f9160208201915f81840391015261a31a565b90565b8061a39561a38f5f613abb565b916108dd565b115f1461a5b95761a3a7610101612b32565b61a3b08261af16565b918161a3cc61a3c661a3c15f613aa1565b61093b565b9161093b565b145f1461a59e5761a4a061a48361a47e61a3ef61a3e9600261a1b5565b5b617abf565b9361a45061a43461a42f61a4028861a26a565b9a61a42a8c5f61a42361a41361731d565b9261a41d83613abb565b9061a295565b5101617725565b613890565b6169a5565b602061a4498b61a4435f613abb565b9061a295565b510161a2b5565b61a47761a45b61731d565b5f61a4708b61a46a6001614589565b9061a295565b5101617725565b86906145a5565b613890565b602061a4998761a4936001614589565b9061a295565b510161a2b5565b61a4b361a4ad600261455d565b916108dd565b1161a54e575b505061a4ce61a4c9610129617d6b565b61371d565b9063b355121490823b1561a5495761a5059261a4fa5f809461a4ee6108cb565b96879586948593613e14565b83526004830161a36a565b03925af1801561a5445761a518575b505b565b61a537905f3d811161a53d575b61a52f8183610cd2565b810190613e1a565b5f61a514565b503d61a525565b613e80565b613e10565b61a5979161a57561a57a925f61a56e8761a568600261455d565b9061a295565b5101617725565b613890565b602061a5908461a58a600261455d565b9061a295565b510161a2b5565b5f8061a4b9565b61a4a061a48361a47e61a3ef61a5b46003617a84565b61a3ea565b5061a5cd61a5c8610129617d6b565b61371d565b90639dc29fac9061a5dc61731d565b9092803b1561a6535761a6025f809461a60d61a5f66108cb565b97889687958694613e14565b84526004840161a192565b03925af1801561a64e5761a622575b5061a516565b61a641905f3d811161a647575b61a6398183610cd2565b810190613e1a565b5f61a61c565b503d61a62f565b613e80565b613e10565b61a6819161a664613a7e565b508161a67861a672836108dd565b916108dd565b1191909161af47565b90565b90565b5190565b5f7f4d65737361676520697320746f6f206c6f6e672e000000000000000000000000910152565b61a6bf60146020926128f4565b61a6c88161a68b565b0190565b919061a6ef90602061a6e7604086018681035f88015261a6b2565b940190610b05565b565b1561a6f95750565b61a71b9061a7056108cb565b91829163271d43ff60e21b83526004830161a6cc565b0390fd5b60207f207368616c6c206265204554482e000000000000000000000000000000000000917f5468652066697273742062696420696e20612062696464696e6720726f756e645f8201520152565b61a779602e6040926128f4565b61a7828161a71f565b0190565b61a79b9060208101905f81830391015261a76c565b90565b1561a7a557565b61a7ad6108cb565b63b5a45a4960e01b81528061a7c46004820161a786565b0390fd5b61a7d1906108dd565b5f19811461a7df5760010190565b614a2e565b61a82e9061a82861a82361a7ff61a7fa8461a684565b61a687565b61a81b61a81561a81061011a610f63565b6108dd565b916108dd565b11159261a684565b61a687565b9061a6f1565b61a839610101612b32565b61a85361a84d61a8485f613aa1565b61093b565b9161093b565b145f1461a99b5761a86261b02d565b61a87e3461a87861a8725f613abb565b916108dd565b1161a79e565b61a88a4261011461392f565b61a8a761a89f4261a899614603565b906157f1565b61012461392f565b61a8b261010b610f63565b429061a8f361a8e17f028a52641badd593b7f30072734c0b97e449213f55b5c3663756427340accd3c92610ece565b9261a8ea6108cb565b91829182610b12565b0390a25b61a90a61a90261731d565b610101613cef565b61a96561a92c61a92761010361a92161010b610f63565b90611563565b616b91565b5f61a95e61a93b828401610f63565b61a95961a94661731d565b61a95460018701849061384a565b613cef565b61a7c8565b910161392f565b61a99942600261a99361a98561010461a97f61010b610f63565b90610eea565b61a98d61731d565b90610f34565b0161392f565b565b61a9a3617448565b61a9ab61af6b565b61a8f7565b61a9bc61a9c191610f4a565b610ece565b90565b61a9d861a9d361a9dd92614586565b610ecb565b617a7e565b90565b61a9ff9061a9f961a9f361aa0494617a7e565b916108dd565b90610c3a565b6108dd565b90565b90565b61aa1e61aa1961aa239261aa07565b610ecb565b617a7e565b90565b61aa459061aa3f61aa3961aa4a94617a7e565b916108dd565b906141cf565b6108dd565b90565b90565b61aa6461aa5f61aa699261aa4d565b610ecb565b617a7e565b90565b90565b61aa8361aa7e61aa889261aa6c565b610ecb565b617a7e565b90565b61aa93613a7e565b5061aac461aab461aaae4361aaa86001614589565b9061633f565b4061a9b0565b61aabe600161a9c4565b9061a9e0565b61aad84861aad2604061aa0a565b9061aa26565b1861aae161b183565b9061ab3f575b5061aaf061b4f9565b9061ab24575b5061aaff61b693565b9061ab09575b5090565b61ab1d9061ab1760c061aa6f565b9061aa26565b185f61ab05565b61ab389061ab32608061aa50565b9061aa26565b185f61aaf6565b61ab5f61ab649161ab4e614fe0565b5061ab596001614589565b9061633f565b61b2f7565b9061ab6f575b61aae7565b61ab789061a9b0565b185f61ab6a565b61abb35f61abb89261ab8f613a7e565b5061abad82820161aba761aba282617afc565b618f2d565b90617633565b01617afc565b61b7af565b90565b634e487b715f526020526024601cfd5b61abd3613846565b5061abee5f61abe861abe3615017565b61b7c4565b01612b32565b90565b9061abfb8261b7c7565b8161ac267fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b91610f28565b9061ac2f6108cb565b8061ac3981610a26565b0390a261ac458161a687565b61ac5761ac515f613abb565b916108dd565b115f1461ac6b5761ac679161b852565b505b565b505061ac7561b81c565b61ac69565b8061ac8d61ac875f613abb565b916108dd565b1161ac96575b50565b61aca1610101612b32565b61acaa8261af16565b908061acc661acc061acbb5f613aa1565b61093b565b9161093b565b145f1461ad76575061ace161acdc610129617d6b565b61371d565b9161acf96340c10f199261acf361731d565b926145a5565b92803b1561ad715761ad1e5f809461ad2961ad126108cb565b97889687958694613e14565b84526004840161a192565b03925af1801561ad6c5761ad40575b505b5f61ac93565b61ad5f905f3d811161ad65575b61ad578183610cd2565b810190613e1a565b5f61ad38565b503d61ad4d565b613e80565b613e10565b9161ae199161adfc9161adb061ad9461ad8f600261455d565b617cdf565b955f61ada98861ada383613abb565b90617d0a565b5101617725565b61add082602061adc98861adc35f613abb565b90617d0a565b5101617633565b61adf761addb61731d565b5f61adf08861adea6001614589565b90617d0a565b5101617725565b6145a5565b602061ae128461ae0c6001614589565b90617d0a565b5101617633565b61ae2c61ae27610129617d6b565b61371d565b9063b33266da90823b1561aea95761ae639261ae585f809461ae4c6108cb565b96879586948593613e14565b835260048301617e12565b03925af1801561aea45761ae78575b5061ad3a565b61ae97905f3d811161ae9d575b61ae8f8183610cd2565b810190613e1a565b5f61ae72565b503d61ae85565b613e80565b613e10565b61aeb661388c565b5061aeda61aec261a058565b61aed561aed0610115610f63565b613890565b6138ac565b90565b61aee561a142565b505c90565b90565b61af0161aefc61af069261aeea565b610ecb565b6108dd565b90565b61af13605a61aeed565b90565b61af3461af449161af25613a7e565b5061af2e61af09565b90613b0f565b61af3e6064613d77565b90613b50565b90565b61af6161af67929361af57613a7e565b508094189161b881565b90613b0f565b1890565b61af9161af8961af79614f40565b61af84610124610f63565b6145b3565b61012461392f565b565b60207f20616374697665207965742e0000000000000000000000000000000000000000917f5468652063757272656e742062696464696e6720726f756e64206973206e6f745f8201520152565b61afed602c6040926128f4565b61aff68161af93565b0190565b91604061b02b92949361b02461b019606083018381035f85015261afe0565b966020830190610b05565b0190610b05565b565b61b03861010d610f63565b61b0564261b04e61b048846108dd565b916108dd565b1015156149aa565b61b05d5750565b429061b08061b06a6108cb565b9283926302dbf17b60e31b84526004840161affa565b0390fd5b61b09861b09361b09d92613d74565b610ecb565b610930565b90565b61b0a99061b084565b90565b61b0b590610f00565b90565b61b0c19061b0ac565b90565b61b0d661b0d1606461b0a0565b61b0b8565b90565b61b0e290610f1c565b90565b90565b61b0fc61b0f761b1019261b0e5565b610ecb565b6108dd565b90565b60207f642e000000000000000000000000000000000000000000000000000000000000917f4172625379732e617262426c6f636b4e756d6265722063616c6c206661696c655f8201520152565b61b15e60226040926128f4565b61b1678161b104565b0190565b61b1809060208101905f81830391015261b151565b90565b61b18b61a142565b5061b194613a7e565b61b19c617950565b505f8061b1af61b1aa61b0c4565b61b0d9565b600461b1e663a3b1b31d60e01b61b1d761b1c76108cb565b9384926020840190815201610a26565b60208201810382520382610cd2565b82602082019151925af19161b1f9617955565b8361b24c575b5061b20a83156149aa565b61b211575b565b7fa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d61b23a6108cb565b8061b2448161b16b565b0390a161b20f565b909161b2578261a687565b61b26a61b264602061b0e8565b916108dd565b145f1461b294575061b28c90602061b2818261a687565b818301019101617780565b905b5f61b1ff565b919250505f9161b28e565b5f7f4172625379732e617262426c6f636b486173682063616c6c206661696c65642e910152565b61b2d2602080926128f4565b61b2db8161b29f565b0190565b61b2f49060208101905f81830391015261b2c6565b90565b61b2ff61a142565b505f8061b30a614fe0565b9261b313617950565b50600461b35d61b32961b32461b0c4565b61b0d9565b9261b34e6315a03d4160e11b9161b33e6108cb565b9485936020850190815201610b12565b60208201810382520382610cd2565b82602082019151925af19161b370617955565b8361b3c3575b5061b38183156149aa565b61b388575b565b7fa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d61b3b16108cb565b8061b3bb8161b2df565b0390a161b386565b909161b3ce8261a687565b61b3e161b3db602061b0e8565b916108dd565b145f1461b40b575061b40390602061b3f88261a687565b818301019101619412565b905b5f61b376565b919250505f9161b405565b90565b61b42d61b42861b4329261b416565b610ecb565b610930565b90565b61b43e9061b419565b90565b61b44a90610f00565b90565b61b4569061b441565b90565b61b46b61b466606c61b435565b61b44d565b90565b61b47790610f1c565b90565b60207f696c65642e000000000000000000000000000000000000000000000000000000917f417262476173496e666f2e6765744761734261636b6c6f672063616c6c2066615f8201520152565b61b4d460256040926128f4565b61b4dd8161b47a565b0190565b61b4f69060208101905f81830391015261b4c7565b90565b61b50161a142565b5061b50a613a7e565b61b512617950565b505f8061b52561b52061b459565b61b46e565b600461b55b62eadae160e51b61b54c61b53c6108cb565b9384926020840190815201610a26565b60208201810382520382610cd2565b82602082019151925af19161b56e617955565b8361b5c1575b5061b57f83156149aa565b61b586575b565b7fa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d61b5af6108cb565b8061b5b98161b4e1565b0390a161b584565b909161b5cc8261a687565b61b5df61b5d9602061b0e8565b916108dd565b145f1461b609575061b60190602061b5f68261a687565b818301019101617780565b905b5f61b574565b919250505f9161b603565b60207f655570646174652063616c6c206661696c65642e000000000000000000000000917f417262476173496e666f2e6765744c3150726963696e67556e69747353696e635f8201520152565b61b66e60346040926128f4565b61b6778161b614565b0190565b61b6909060208101905f81830391015261b661565b90565b61b69b61a142565b5061b6a4613a7e565b61b6ac617950565b505f8061b6bf61b6ba61b459565b61b46e565b600461b6f66377f8098360e11b61b6e761b6d76108cb565b9384926020840190815201610a26565b60208201810382520382610cd2565b82602082019151925af19161b709617955565b8361b75c575b5061b71a83156149aa565b61b721575b565b7fa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d61b74a6108cb565b8061b7548161b67b565b0390a161b71f565b909161b7678261a687565b61b77a61b774602061b0e8565b916108dd565b145f1461b7a4575061b79c90602061b7918261a687565b818301019101617780565b905b5f61b70f565b919250505f9161b79e565b61b7c19061b7bb613a7e565b5061b88f565b90565b90565b803b61b7db61b7d55f613abb565b916108dd565b1461b7fd5761b7fb905f61b7f561b7f0615017565b61b7c4565b01613cef565b565b61b818905f918291634c9c8ce360e01b835260048301610954565b0390fd5b3461b82f61b8295f613abb565b916108dd565b1161b83657565b5f63b398979f60e01b81528061b84e60048201610a26565b0390fd5b5f8061b87e9361b860617950565b508390602081019051915af49061b875617955565b9091909161b8a1565b90565b61b889613a7e565b50151590565b61b897613a7e565b505f5260205f2090565b9061b8b59061b8ae617950565b50156149aa565b5f1461b8c1575061b925565b61b8ca8261a687565b61b8dc61b8d65f613abb565b916108dd565b148061b90a575b61b8eb575090565b61b906905f918291639996b31560e01b835260048301610954565b0390fd5b50803b61b91f61b9195f613abb565b916108dd565b1461b8e3565b61b92e8161a687565b61b94061b93a5f613abb565b916108dd565b115f1461b94f57805190602001fd5b5f63d6bda27560e01b81528061b96760048201610a26565b0390fdfea264697066735822122060225b922a8b312d5c13be59a1efc86d52dd5cb3879c2f2a277b3ee4b36bd52a64736f6c63430008220033",
}

// CosmicSignatureGameV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use CosmicSignatureGameV3MetaData.ABI instead.
var CosmicSignatureGameV3ABI = CosmicSignatureGameV3MetaData.ABI

// CosmicSignatureGameV3Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CosmicSignatureGameV3MetaData.Bin instead.
var CosmicSignatureGameV3Bin = CosmicSignatureGameV3MetaData.Bin

// DeployCosmicSignatureGameV3 deploys a new Ethereum contract, binding an instance of CosmicSignatureGameV3 to it.
func DeployCosmicSignatureGameV3(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CosmicSignatureGameV3, error) {
	parsed, err := CosmicSignatureGameV3MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CosmicSignatureGameV3Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CosmicSignatureGameV3{CosmicSignatureGameV3Caller: CosmicSignatureGameV3Caller{contract: contract}, CosmicSignatureGameV3Transactor: CosmicSignatureGameV3Transactor{contract: contract}, CosmicSignatureGameV3Filterer: CosmicSignatureGameV3Filterer{contract: contract}}, nil
}

// CosmicSignatureGameV3 is an auto generated Go binding around an Ethereum contract.
type CosmicSignatureGameV3 struct {
	CosmicSignatureGameV3Caller     // Read-only binding to the contract
	CosmicSignatureGameV3Transactor // Write-only binding to the contract
	CosmicSignatureGameV3Filterer   // Log filterer for contract events
}

// CosmicSignatureGameV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type CosmicSignatureGameV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmicSignatureGameV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CosmicSignatureGameV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmicSignatureGameV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CosmicSignatureGameV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmicSignatureGameV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CosmicSignatureGameV3Session struct {
	Contract     *CosmicSignatureGameV3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CosmicSignatureGameV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CosmicSignatureGameV3CallerSession struct {
	Contract *CosmicSignatureGameV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// CosmicSignatureGameV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CosmicSignatureGameV3TransactorSession struct {
	Contract     *CosmicSignatureGameV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// CosmicSignatureGameV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type CosmicSignatureGameV3Raw struct {
	Contract *CosmicSignatureGameV3 // Generic contract binding to access the raw methods on
}

// CosmicSignatureGameV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CosmicSignatureGameV3CallerRaw struct {
	Contract *CosmicSignatureGameV3Caller // Generic read-only contract binding to access the raw methods on
}

// CosmicSignatureGameV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CosmicSignatureGameV3TransactorRaw struct {
	Contract *CosmicSignatureGameV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCosmicSignatureGameV3 creates a new instance of CosmicSignatureGameV3, bound to a specific deployed contract.
func NewCosmicSignatureGameV3(address common.Address, backend bind.ContractBackend) (*CosmicSignatureGameV3, error) {
	contract, err := bindCosmicSignatureGameV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3{CosmicSignatureGameV3Caller: CosmicSignatureGameV3Caller{contract: contract}, CosmicSignatureGameV3Transactor: CosmicSignatureGameV3Transactor{contract: contract}, CosmicSignatureGameV3Filterer: CosmicSignatureGameV3Filterer{contract: contract}}, nil
}

// NewCosmicSignatureGameV3Caller creates a new read-only instance of CosmicSignatureGameV3, bound to a specific deployed contract.
func NewCosmicSignatureGameV3Caller(address common.Address, caller bind.ContractCaller) (*CosmicSignatureGameV3Caller, error) {
	contract, err := bindCosmicSignatureGameV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3Caller{contract: contract}, nil
}

// NewCosmicSignatureGameV3Transactor creates a new write-only instance of CosmicSignatureGameV3, bound to a specific deployed contract.
func NewCosmicSignatureGameV3Transactor(address common.Address, transactor bind.ContractTransactor) (*CosmicSignatureGameV3Transactor, error) {
	contract, err := bindCosmicSignatureGameV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3Transactor{contract: contract}, nil
}

// NewCosmicSignatureGameV3Filterer creates a new log filterer instance of CosmicSignatureGameV3, bound to a specific deployed contract.
func NewCosmicSignatureGameV3Filterer(address common.Address, filterer bind.ContractFilterer) (*CosmicSignatureGameV3Filterer, error) {
	contract, err := bindCosmicSignatureGameV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3Filterer{contract: contract}, nil
}

// bindCosmicSignatureGameV3 binds a generic wrapper to an already deployed contract.
func bindCosmicSignatureGameV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CosmicSignatureGameV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmicSignatureGameV3.Contract.CosmicSignatureGameV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.CosmicSignatureGameV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.CosmicSignatureGameV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmicSignatureGameV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) UPGRADEINTERFACEVERSION() (string, error) {
	return _CosmicSignatureGameV3.Contract.UPGRADEINTERFACEVERSION(&_CosmicSignatureGameV3.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _CosmicSignatureGameV3.Contract.UPGRADEINTERFACEVERSION(&_CosmicSignatureGameV3.CallOpts)
}

// BidCstRewardAmountMultiplier is a free data retrieval call binding the contract method 0xb30f5bb1.
//
// Solidity: function bidCstRewardAmountMultiplier() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) BidCstRewardAmountMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "bidCstRewardAmountMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BidCstRewardAmountMultiplier is a free data retrieval call binding the contract method 0xb30f5bb1.
//
// Solidity: function bidCstRewardAmountMultiplier() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) BidCstRewardAmountMultiplier() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.BidCstRewardAmountMultiplier(&_CosmicSignatureGameV3.CallOpts)
}

// BidCstRewardAmountMultiplier is a free data retrieval call binding the contract method 0xb30f5bb1.
//
// Solidity: function bidCstRewardAmountMultiplier() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) BidCstRewardAmountMultiplier() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.BidCstRewardAmountMultiplier(&_CosmicSignatureGameV3.CallOpts)
}

// BidCstRewardAmountPerMinute is a free data retrieval call binding the contract method 0x4ad8a90e.
//
// Solidity: function bidCstRewardAmountPerMinute() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) BidCstRewardAmountPerMinute(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "bidCstRewardAmountPerMinute")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BidCstRewardAmountPerMinute is a free data retrieval call binding the contract method 0x4ad8a90e.
//
// Solidity: function bidCstRewardAmountPerMinute() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) BidCstRewardAmountPerMinute() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.BidCstRewardAmountPerMinute(&_CosmicSignatureGameV3.CallOpts)
}

// BidCstRewardAmountPerMinute is a free data retrieval call binding the contract method 0x4ad8a90e.
//
// Solidity: function bidCstRewardAmountPerMinute() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) BidCstRewardAmountPerMinute() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.BidCstRewardAmountPerMinute(&_CosmicSignatureGameV3.CallOpts)
}

// BidMessageLengthMaxLimit is a free data retrieval call binding the contract method 0x6c0613c0.
//
// Solidity: function bidMessageLengthMaxLimit() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) BidMessageLengthMaxLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "bidMessageLengthMaxLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BidMessageLengthMaxLimit is a free data retrieval call binding the contract method 0x6c0613c0.
//
// Solidity: function bidMessageLengthMaxLimit() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) BidMessageLengthMaxLimit() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.BidMessageLengthMaxLimit(&_CosmicSignatureGameV3.CallOpts)
}

// BidMessageLengthMaxLimit is a free data retrieval call binding the contract method 0x6c0613c0.
//
// Solidity: function bidMessageLengthMaxLimit() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) BidMessageLengthMaxLimit() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.BidMessageLengthMaxLimit(&_CosmicSignatureGameV3.CallOpts)
}

// BidderAddresses is a free data retrieval call binding the contract method 0x2fb3c48f.
//
// Solidity: function bidderAddresses(uint256 roundNum) view returns(uint256 numItems)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) BidderAddresses(opts *bind.CallOpts, roundNum *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "bidderAddresses", roundNum)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BidderAddresses is a free data retrieval call binding the contract method 0x2fb3c48f.
//
// Solidity: function bidderAddresses(uint256 roundNum) view returns(uint256 numItems)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) BidderAddresses(roundNum *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.BidderAddresses(&_CosmicSignatureGameV3.CallOpts, roundNum)
}

// BidderAddresses is a free data retrieval call binding the contract method 0x2fb3c48f.
//
// Solidity: function bidderAddresses(uint256 roundNum) view returns(uint256 numItems)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) BidderAddresses(roundNum *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.BidderAddresses(&_CosmicSignatureGameV3.CallOpts, roundNum)
}

// BiddersInfo is a free data retrieval call binding the contract method 0x17887731.
//
// Solidity: function biddersInfo(uint256 roundNum, address bidderAddress) view returns(uint256 totalSpentEthAmount, uint256 totalSpentCstAmount, uint256 lastBidTimeStamp)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) BiddersInfo(opts *bind.CallOpts, roundNum *big.Int, bidderAddress common.Address) (struct {
	TotalSpentEthAmount *big.Int
	TotalSpentCstAmount *big.Int
	LastBidTimeStamp    *big.Int
}, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "biddersInfo", roundNum, bidderAddress)

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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) BiddersInfo(roundNum *big.Int, bidderAddress common.Address) (struct {
	TotalSpentEthAmount *big.Int
	TotalSpentCstAmount *big.Int
	LastBidTimeStamp    *big.Int
}, error) {
	return _CosmicSignatureGameV3.Contract.BiddersInfo(&_CosmicSignatureGameV3.CallOpts, roundNum, bidderAddress)
}

// BiddersInfo is a free data retrieval call binding the contract method 0x17887731.
//
// Solidity: function biddersInfo(uint256 roundNum, address bidderAddress) view returns(uint256 totalSpentEthAmount, uint256 totalSpentCstAmount, uint256 lastBidTimeStamp)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) BiddersInfo(roundNum *big.Int, bidderAddress common.Address) (struct {
	TotalSpentEthAmount *big.Int
	TotalSpentCstAmount *big.Int
	LastBidTimeStamp    *big.Int
}, error) {
	return _CosmicSignatureGameV3.Contract.BiddersInfo(&_CosmicSignatureGameV3.CallOpts, roundNum, bidderAddress)
}

// CharityAddress is a free data retrieval call binding the contract method 0xafcf2fc4.
//
// Solidity: function charityAddress() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) CharityAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "charityAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CharityAddress is a free data retrieval call binding the contract method 0xafcf2fc4.
//
// Solidity: function charityAddress() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) CharityAddress() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.CharityAddress(&_CosmicSignatureGameV3.CallOpts)
}

// CharityAddress is a free data retrieval call binding the contract method 0xafcf2fc4.
//
// Solidity: function charityAddress() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) CharityAddress() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.CharityAddress(&_CosmicSignatureGameV3.CallOpts)
}

// CharityEthDonationAmountPercentage is a free data retrieval call binding the contract method 0xbe720ad5.
//
// Solidity: function charityEthDonationAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) CharityEthDonationAmountPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "charityEthDonationAmountPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CharityEthDonationAmountPercentage is a free data retrieval call binding the contract method 0xbe720ad5.
//
// Solidity: function charityEthDonationAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) CharityEthDonationAmountPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.CharityEthDonationAmountPercentage(&_CosmicSignatureGameV3.CallOpts)
}

// CharityEthDonationAmountPercentage is a free data retrieval call binding the contract method 0xbe720ad5.
//
// Solidity: function charityEthDonationAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) CharityEthDonationAmountPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.CharityEthDonationAmountPercentage(&_CosmicSignatureGameV3.CallOpts)
}

// ChronoWarriorAddress is a free data retrieval call binding the contract method 0x6b7cbe85.
//
// Solidity: function chronoWarriorAddress() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) ChronoWarriorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "chronoWarriorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChronoWarriorAddress is a free data retrieval call binding the contract method 0x6b7cbe85.
//
// Solidity: function chronoWarriorAddress() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) ChronoWarriorAddress() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.ChronoWarriorAddress(&_CosmicSignatureGameV3.CallOpts)
}

// ChronoWarriorAddress is a free data retrieval call binding the contract method 0x6b7cbe85.
//
// Solidity: function chronoWarriorAddress() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) ChronoWarriorAddress() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.ChronoWarriorAddress(&_CosmicSignatureGameV3.CallOpts)
}

// ChronoWarriorDuration is a free data retrieval call binding the contract method 0x1824d5e7.
//
// Solidity: function chronoWarriorDuration() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) ChronoWarriorDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "chronoWarriorDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChronoWarriorDuration is a free data retrieval call binding the contract method 0x1824d5e7.
//
// Solidity: function chronoWarriorDuration() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) ChronoWarriorDuration() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.ChronoWarriorDuration(&_CosmicSignatureGameV3.CallOpts)
}

// ChronoWarriorDuration is a free data retrieval call binding the contract method 0x1824d5e7.
//
// Solidity: function chronoWarriorDuration() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) ChronoWarriorDuration() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.ChronoWarriorDuration(&_CosmicSignatureGameV3.CallOpts)
}

// ChronoWarriorEthPrizeAmountPercentage is a free data retrieval call binding the contract method 0x54ada1d6.
//
// Solidity: function chronoWarriorEthPrizeAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) ChronoWarriorEthPrizeAmountPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "chronoWarriorEthPrizeAmountPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChronoWarriorEthPrizeAmountPercentage is a free data retrieval call binding the contract method 0x54ada1d6.
//
// Solidity: function chronoWarriorEthPrizeAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) ChronoWarriorEthPrizeAmountPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.ChronoWarriorEthPrizeAmountPercentage(&_CosmicSignatureGameV3.CallOpts)
}

// ChronoWarriorEthPrizeAmountPercentage is a free data retrieval call binding the contract method 0x54ada1d6.
//
// Solidity: function chronoWarriorEthPrizeAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) ChronoWarriorEthPrizeAmountPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.ChronoWarriorEthPrizeAmountPercentage(&_CosmicSignatureGameV3.CallOpts)
}

// CosmicSignatureNftStakingTotalEthRewardAmountPercentage is a free data retrieval call binding the contract method 0xf7bea078.
//
// Solidity: function cosmicSignatureNftStakingTotalEthRewardAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) CosmicSignatureNftStakingTotalEthRewardAmountPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "cosmicSignatureNftStakingTotalEthRewardAmountPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CosmicSignatureNftStakingTotalEthRewardAmountPercentage is a free data retrieval call binding the contract method 0xf7bea078.
//
// Solidity: function cosmicSignatureNftStakingTotalEthRewardAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) CosmicSignatureNftStakingTotalEthRewardAmountPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.CosmicSignatureNftStakingTotalEthRewardAmountPercentage(&_CosmicSignatureGameV3.CallOpts)
}

// CosmicSignatureNftStakingTotalEthRewardAmountPercentage is a free data retrieval call binding the contract method 0xf7bea078.
//
// Solidity: function cosmicSignatureNftStakingTotalEthRewardAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) CosmicSignatureNftStakingTotalEthRewardAmountPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.CosmicSignatureNftStakingTotalEthRewardAmountPercentage(&_CosmicSignatureGameV3.CallOpts)
}

// CstDutchAuctionBeginningBidPrice is a free data retrieval call binding the contract method 0x1b410319.
//
// Solidity: function cstDutchAuctionBeginningBidPrice() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) CstDutchAuctionBeginningBidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "cstDutchAuctionBeginningBidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CstDutchAuctionBeginningBidPrice is a free data retrieval call binding the contract method 0x1b410319.
//
// Solidity: function cstDutchAuctionBeginningBidPrice() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) CstDutchAuctionBeginningBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.CstDutchAuctionBeginningBidPrice(&_CosmicSignatureGameV3.CallOpts)
}

// CstDutchAuctionBeginningBidPrice is a free data retrieval call binding the contract method 0x1b410319.
//
// Solidity: function cstDutchAuctionBeginningBidPrice() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) CstDutchAuctionBeginningBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.CstDutchAuctionBeginningBidPrice(&_CosmicSignatureGameV3.CallOpts)
}

// CstDutchAuctionBeginningBidPriceMinLimit is a free data retrieval call binding the contract method 0xbb4b3e6f.
//
// Solidity: function cstDutchAuctionBeginningBidPriceMinLimit() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) CstDutchAuctionBeginningBidPriceMinLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "cstDutchAuctionBeginningBidPriceMinLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CstDutchAuctionBeginningBidPriceMinLimit is a free data retrieval call binding the contract method 0xbb4b3e6f.
//
// Solidity: function cstDutchAuctionBeginningBidPriceMinLimit() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) CstDutchAuctionBeginningBidPriceMinLimit() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.CstDutchAuctionBeginningBidPriceMinLimit(&_CosmicSignatureGameV3.CallOpts)
}

// CstDutchAuctionBeginningBidPriceMinLimit is a free data retrieval call binding the contract method 0xbb4b3e6f.
//
// Solidity: function cstDutchAuctionBeginningBidPriceMinLimit() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) CstDutchAuctionBeginningBidPriceMinLimit() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.CstDutchAuctionBeginningBidPriceMinLimit(&_CosmicSignatureGameV3.CallOpts)
}

// CstDutchAuctionBeginningTimeStamp is a free data retrieval call binding the contract method 0x27700481.
//
// Solidity: function cstDutchAuctionBeginningTimeStamp() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) CstDutchAuctionBeginningTimeStamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "cstDutchAuctionBeginningTimeStamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CstDutchAuctionBeginningTimeStamp is a free data retrieval call binding the contract method 0x27700481.
//
// Solidity: function cstDutchAuctionBeginningTimeStamp() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) CstDutchAuctionBeginningTimeStamp() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.CstDutchAuctionBeginningTimeStamp(&_CosmicSignatureGameV3.CallOpts)
}

// CstDutchAuctionBeginningTimeStamp is a free data retrieval call binding the contract method 0x27700481.
//
// Solidity: function cstDutchAuctionBeginningTimeStamp() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) CstDutchAuctionBeginningTimeStamp() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.CstDutchAuctionBeginningTimeStamp(&_CosmicSignatureGameV3.CallOpts)
}

// CstDutchAuctionDuration is a free data retrieval call binding the contract method 0x9302020f.
//
// Solidity: function cstDutchAuctionDuration() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) CstDutchAuctionDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "cstDutchAuctionDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CstDutchAuctionDuration is a free data retrieval call binding the contract method 0x9302020f.
//
// Solidity: function cstDutchAuctionDuration() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) CstDutchAuctionDuration() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.CstDutchAuctionDuration(&_CosmicSignatureGameV3.CallOpts)
}

// CstDutchAuctionDuration is a free data retrieval call binding the contract method 0x9302020f.
//
// Solidity: function cstDutchAuctionDuration() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) CstDutchAuctionDuration() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.CstDutchAuctionDuration(&_CosmicSignatureGameV3.CallOpts)
}

// CstDutchAuctionDurationChangeDivisor is a free data retrieval call binding the contract method 0xda9931dd.
//
// Solidity: function cstDutchAuctionDurationChangeDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) CstDutchAuctionDurationChangeDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "cstDutchAuctionDurationChangeDivisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CstDutchAuctionDurationChangeDivisor is a free data retrieval call binding the contract method 0xda9931dd.
//
// Solidity: function cstDutchAuctionDurationChangeDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) CstDutchAuctionDurationChangeDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.CstDutchAuctionDurationChangeDivisor(&_CosmicSignatureGameV3.CallOpts)
}

// CstDutchAuctionDurationChangeDivisor is a free data retrieval call binding the contract method 0xda9931dd.
//
// Solidity: function cstDutchAuctionDurationChangeDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) CstDutchAuctionDurationChangeDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.CstDutchAuctionDurationChangeDivisor(&_CosmicSignatureGameV3.CallOpts)
}

// CstPrizeAmount is a free data retrieval call binding the contract method 0x320c435c.
//
// Solidity: function cstPrizeAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) CstPrizeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "cstPrizeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CstPrizeAmount is a free data retrieval call binding the contract method 0x320c435c.
//
// Solidity: function cstPrizeAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) CstPrizeAmount() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.CstPrizeAmount(&_CosmicSignatureGameV3.CallOpts)
}

// CstPrizeAmount is a free data retrieval call binding the contract method 0x320c435c.
//
// Solidity: function cstPrizeAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) CstPrizeAmount() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.CstPrizeAmount(&_CosmicSignatureGameV3.CallOpts)
}

// DelayDurationBeforeRoundActivation is a free data retrieval call binding the contract method 0xb9cf9ba5.
//
// Solidity: function delayDurationBeforeRoundActivation() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) DelayDurationBeforeRoundActivation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "delayDurationBeforeRoundActivation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayDurationBeforeRoundActivation is a free data retrieval call binding the contract method 0xb9cf9ba5.
//
// Solidity: function delayDurationBeforeRoundActivation() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) DelayDurationBeforeRoundActivation() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.DelayDurationBeforeRoundActivation(&_CosmicSignatureGameV3.CallOpts)
}

// DelayDurationBeforeRoundActivation is a free data retrieval call binding the contract method 0xb9cf9ba5.
//
// Solidity: function delayDurationBeforeRoundActivation() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) DelayDurationBeforeRoundActivation() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.DelayDurationBeforeRoundActivation(&_CosmicSignatureGameV3.CallOpts)
}

// EnduranceChampionAddress is a free data retrieval call binding the contract method 0x9e50acc9.
//
// Solidity: function enduranceChampionAddress() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) EnduranceChampionAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "enduranceChampionAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EnduranceChampionAddress is a free data retrieval call binding the contract method 0x9e50acc9.
//
// Solidity: function enduranceChampionAddress() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) EnduranceChampionAddress() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.EnduranceChampionAddress(&_CosmicSignatureGameV3.CallOpts)
}

// EnduranceChampionAddress is a free data retrieval call binding the contract method 0x9e50acc9.
//
// Solidity: function enduranceChampionAddress() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) EnduranceChampionAddress() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.EnduranceChampionAddress(&_CosmicSignatureGameV3.CallOpts)
}

// EnduranceChampionDuration is a free data retrieval call binding the contract method 0x5863a705.
//
// Solidity: function enduranceChampionDuration() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) EnduranceChampionDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "enduranceChampionDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnduranceChampionDuration is a free data retrieval call binding the contract method 0x5863a705.
//
// Solidity: function enduranceChampionDuration() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) EnduranceChampionDuration() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.EnduranceChampionDuration(&_CosmicSignatureGameV3.CallOpts)
}

// EnduranceChampionDuration is a free data retrieval call binding the contract method 0x5863a705.
//
// Solidity: function enduranceChampionDuration() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) EnduranceChampionDuration() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.EnduranceChampionDuration(&_CosmicSignatureGameV3.CallOpts)
}

// EnduranceChampionStartTimeStamp is a free data retrieval call binding the contract method 0x250fadb6.
//
// Solidity: function enduranceChampionStartTimeStamp() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) EnduranceChampionStartTimeStamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "enduranceChampionStartTimeStamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnduranceChampionStartTimeStamp is a free data retrieval call binding the contract method 0x250fadb6.
//
// Solidity: function enduranceChampionStartTimeStamp() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) EnduranceChampionStartTimeStamp() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.EnduranceChampionStartTimeStamp(&_CosmicSignatureGameV3.CallOpts)
}

// EnduranceChampionStartTimeStamp is a free data retrieval call binding the contract method 0x250fadb6.
//
// Solidity: function enduranceChampionStartTimeStamp() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) EnduranceChampionStartTimeStamp() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.EnduranceChampionStartTimeStamp(&_CosmicSignatureGameV3.CallOpts)
}

// EthBidPriceIncreaseDivisor is a free data retrieval call binding the contract method 0xa9742016.
//
// Solidity: function ethBidPriceIncreaseDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) EthBidPriceIncreaseDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "ethBidPriceIncreaseDivisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthBidPriceIncreaseDivisor is a free data retrieval call binding the contract method 0xa9742016.
//
// Solidity: function ethBidPriceIncreaseDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) EthBidPriceIncreaseDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.EthBidPriceIncreaseDivisor(&_CosmicSignatureGameV3.CallOpts)
}

// EthBidPriceIncreaseDivisor is a free data retrieval call binding the contract method 0xa9742016.
//
// Solidity: function ethBidPriceIncreaseDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) EthBidPriceIncreaseDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.EthBidPriceIncreaseDivisor(&_CosmicSignatureGameV3.CallOpts)
}

// EthBidRefundAmountInGasToSwallowMaxLimit is a free data retrieval call binding the contract method 0x9aa1b38d.
//
// Solidity: function ethBidRefundAmountInGasToSwallowMaxLimit() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) EthBidRefundAmountInGasToSwallowMaxLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "ethBidRefundAmountInGasToSwallowMaxLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthBidRefundAmountInGasToSwallowMaxLimit is a free data retrieval call binding the contract method 0x9aa1b38d.
//
// Solidity: function ethBidRefundAmountInGasToSwallowMaxLimit() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) EthBidRefundAmountInGasToSwallowMaxLimit() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.EthBidRefundAmountInGasToSwallowMaxLimit(&_CosmicSignatureGameV3.CallOpts)
}

// EthBidRefundAmountInGasToSwallowMaxLimit is a free data retrieval call binding the contract method 0x9aa1b38d.
//
// Solidity: function ethBidRefundAmountInGasToSwallowMaxLimit() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) EthBidRefundAmountInGasToSwallowMaxLimit() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.EthBidRefundAmountInGasToSwallowMaxLimit(&_CosmicSignatureGameV3.CallOpts)
}

// EthDonationWithInfoRecords is a free data retrieval call binding the contract method 0xb5d1f06f.
//
// Solidity: function ethDonationWithInfoRecords(uint256 ) view returns(uint256 roundNum, address donorAddress, uint256 amount, string data)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) EthDonationWithInfoRecords(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RoundNum     *big.Int
	DonorAddress common.Address
	Amount       *big.Int
	Data         string
}, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "ethDonationWithInfoRecords", arg0)

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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) EthDonationWithInfoRecords(arg0 *big.Int) (struct {
	RoundNum     *big.Int
	DonorAddress common.Address
	Amount       *big.Int
	Data         string
}, error) {
	return _CosmicSignatureGameV3.Contract.EthDonationWithInfoRecords(&_CosmicSignatureGameV3.CallOpts, arg0)
}

// EthDonationWithInfoRecords is a free data retrieval call binding the contract method 0xb5d1f06f.
//
// Solidity: function ethDonationWithInfoRecords(uint256 ) view returns(uint256 roundNum, address donorAddress, uint256 amount, string data)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) EthDonationWithInfoRecords(arg0 *big.Int) (struct {
	RoundNum     *big.Int
	DonorAddress common.Address
	Amount       *big.Int
	Data         string
}, error) {
	return _CosmicSignatureGameV3.Contract.EthDonationWithInfoRecords(&_CosmicSignatureGameV3.CallOpts, arg0)
}

// EthDutchAuctionBeginningBidPrice is a free data retrieval call binding the contract method 0xc87baab5.
//
// Solidity: function ethDutchAuctionBeginningBidPrice() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) EthDutchAuctionBeginningBidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "ethDutchAuctionBeginningBidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthDutchAuctionBeginningBidPrice is a free data retrieval call binding the contract method 0xc87baab5.
//
// Solidity: function ethDutchAuctionBeginningBidPrice() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) EthDutchAuctionBeginningBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.EthDutchAuctionBeginningBidPrice(&_CosmicSignatureGameV3.CallOpts)
}

// EthDutchAuctionBeginningBidPrice is a free data retrieval call binding the contract method 0xc87baab5.
//
// Solidity: function ethDutchAuctionBeginningBidPrice() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) EthDutchAuctionBeginningBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.EthDutchAuctionBeginningBidPrice(&_CosmicSignatureGameV3.CallOpts)
}

// EthDutchAuctionDurationDivisor is a free data retrieval call binding the contract method 0xd1f8fcf2.
//
// Solidity: function ethDutchAuctionDurationDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) EthDutchAuctionDurationDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "ethDutchAuctionDurationDivisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthDutchAuctionDurationDivisor is a free data retrieval call binding the contract method 0xd1f8fcf2.
//
// Solidity: function ethDutchAuctionDurationDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) EthDutchAuctionDurationDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.EthDutchAuctionDurationDivisor(&_CosmicSignatureGameV3.CallOpts)
}

// EthDutchAuctionDurationDivisor is a free data retrieval call binding the contract method 0xd1f8fcf2.
//
// Solidity: function ethDutchAuctionDurationDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) EthDutchAuctionDurationDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.EthDutchAuctionDurationDivisor(&_CosmicSignatureGameV3.CallOpts)
}

// EthDutchAuctionEndingBidPriceDivisor is a free data retrieval call binding the contract method 0xebaa1ea8.
//
// Solidity: function ethDutchAuctionEndingBidPriceDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) EthDutchAuctionEndingBidPriceDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "ethDutchAuctionEndingBidPriceDivisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthDutchAuctionEndingBidPriceDivisor is a free data retrieval call binding the contract method 0xebaa1ea8.
//
// Solidity: function ethDutchAuctionEndingBidPriceDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) EthDutchAuctionEndingBidPriceDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.EthDutchAuctionEndingBidPriceDivisor(&_CosmicSignatureGameV3.CallOpts)
}

// EthDutchAuctionEndingBidPriceDivisor is a free data retrieval call binding the contract method 0xebaa1ea8.
//
// Solidity: function ethDutchAuctionEndingBidPriceDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) EthDutchAuctionEndingBidPriceDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.EthDutchAuctionEndingBidPriceDivisor(&_CosmicSignatureGameV3.CallOpts)
}

// GetBidCstRewardAmount is a free data retrieval call binding the contract method 0xbaab4430.
//
// Solidity: function getBidCstRewardAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetBidCstRewardAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getBidCstRewardAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBidCstRewardAmount is a free data retrieval call binding the contract method 0xbaab4430.
//
// Solidity: function getBidCstRewardAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetBidCstRewardAmount() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetBidCstRewardAmount(&_CosmicSignatureGameV3.CallOpts)
}

// GetBidCstRewardAmount is a free data retrieval call binding the contract method 0xbaab4430.
//
// Solidity: function getBidCstRewardAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetBidCstRewardAmount() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetBidCstRewardAmount(&_CosmicSignatureGameV3.CallOpts)
}

// GetBidCstRewardAmountAdvanced is a free data retrieval call binding the contract method 0x0a120648.
//
// Solidity: function getBidCstRewardAmountAdvanced(int256 currentTimeOffset_) view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetBidCstRewardAmountAdvanced(opts *bind.CallOpts, currentTimeOffset_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getBidCstRewardAmountAdvanced", currentTimeOffset_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBidCstRewardAmountAdvanced is a free data retrieval call binding the contract method 0x0a120648.
//
// Solidity: function getBidCstRewardAmountAdvanced(int256 currentTimeOffset_) view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetBidCstRewardAmountAdvanced(currentTimeOffset_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetBidCstRewardAmountAdvanced(&_CosmicSignatureGameV3.CallOpts, currentTimeOffset_)
}

// GetBidCstRewardAmountAdvanced is a free data retrieval call binding the contract method 0x0a120648.
//
// Solidity: function getBidCstRewardAmountAdvanced(int256 currentTimeOffset_) view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetBidCstRewardAmountAdvanced(currentTimeOffset_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetBidCstRewardAmountAdvanced(&_CosmicSignatureGameV3.CallOpts, currentTimeOffset_)
}

// GetBidderAddressAt is a free data retrieval call binding the contract method 0x000ac9f1.
//
// Solidity: function getBidderAddressAt(uint256 roundNum_, uint256 bidIndex_) view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetBidderAddressAt(opts *bind.CallOpts, roundNum_ *big.Int, bidIndex_ *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getBidderAddressAt", roundNum_, bidIndex_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBidderAddressAt is a free data retrieval call binding the contract method 0x000ac9f1.
//
// Solidity: function getBidderAddressAt(uint256 roundNum_, uint256 bidIndex_) view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetBidderAddressAt(roundNum_ *big.Int, bidIndex_ *big.Int) (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.GetBidderAddressAt(&_CosmicSignatureGameV3.CallOpts, roundNum_, bidIndex_)
}

// GetBidderAddressAt is a free data retrieval call binding the contract method 0x000ac9f1.
//
// Solidity: function getBidderAddressAt(uint256 roundNum_, uint256 bidIndex_) view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetBidderAddressAt(roundNum_ *big.Int, bidIndex_ *big.Int) (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.GetBidderAddressAt(&_CosmicSignatureGameV3.CallOpts, roundNum_, bidIndex_)
}

// GetBidderTotalSpentAmounts is a free data retrieval call binding the contract method 0xfd9b3747.
//
// Solidity: function getBidderTotalSpentAmounts(uint256 roundNum_, address bidderAddress_) view returns(uint256, uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetBidderTotalSpentAmounts(opts *bind.CallOpts, roundNum_ *big.Int, bidderAddress_ common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getBidderTotalSpentAmounts", roundNum_, bidderAddress_)

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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetBidderTotalSpentAmounts(roundNum_ *big.Int, bidderAddress_ common.Address) (*big.Int, *big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetBidderTotalSpentAmounts(&_CosmicSignatureGameV3.CallOpts, roundNum_, bidderAddress_)
}

// GetBidderTotalSpentAmounts is a free data retrieval call binding the contract method 0xfd9b3747.
//
// Solidity: function getBidderTotalSpentAmounts(uint256 roundNum_, address bidderAddress_) view returns(uint256, uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetBidderTotalSpentAmounts(roundNum_ *big.Int, bidderAddress_ common.Address) (*big.Int, *big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetBidderTotalSpentAmounts(&_CosmicSignatureGameV3.CallOpts, roundNum_, bidderAddress_)
}

// GetCharityEthDonationAmount is a free data retrieval call binding the contract method 0x0eb16be6.
//
// Solidity: function getCharityEthDonationAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetCharityEthDonationAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getCharityEthDonationAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCharityEthDonationAmount is a free data retrieval call binding the contract method 0x0eb16be6.
//
// Solidity: function getCharityEthDonationAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetCharityEthDonationAmount() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetCharityEthDonationAmount(&_CosmicSignatureGameV3.CallOpts)
}

// GetCharityEthDonationAmount is a free data retrieval call binding the contract method 0x0eb16be6.
//
// Solidity: function getCharityEthDonationAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetCharityEthDonationAmount() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetCharityEthDonationAmount(&_CosmicSignatureGameV3.CallOpts)
}

// GetChronoWarriorEthPrizeAmount is a free data retrieval call binding the contract method 0x2665c882.
//
// Solidity: function getChronoWarriorEthPrizeAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetChronoWarriorEthPrizeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getChronoWarriorEthPrizeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChronoWarriorEthPrizeAmount is a free data retrieval call binding the contract method 0x2665c882.
//
// Solidity: function getChronoWarriorEthPrizeAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetChronoWarriorEthPrizeAmount() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetChronoWarriorEthPrizeAmount(&_CosmicSignatureGameV3.CallOpts)
}

// GetChronoWarriorEthPrizeAmount is a free data retrieval call binding the contract method 0x2665c882.
//
// Solidity: function getChronoWarriorEthPrizeAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetChronoWarriorEthPrizeAmount() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetChronoWarriorEthPrizeAmount(&_CosmicSignatureGameV3.CallOpts)
}

// GetCosmicSignatureNftStakingTotalEthRewardAmount is a free data retrieval call binding the contract method 0x5f0112fe.
//
// Solidity: function getCosmicSignatureNftStakingTotalEthRewardAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetCosmicSignatureNftStakingTotalEthRewardAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getCosmicSignatureNftStakingTotalEthRewardAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCosmicSignatureNftStakingTotalEthRewardAmount is a free data retrieval call binding the contract method 0x5f0112fe.
//
// Solidity: function getCosmicSignatureNftStakingTotalEthRewardAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetCosmicSignatureNftStakingTotalEthRewardAmount() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetCosmicSignatureNftStakingTotalEthRewardAmount(&_CosmicSignatureGameV3.CallOpts)
}

// GetCosmicSignatureNftStakingTotalEthRewardAmount is a free data retrieval call binding the contract method 0x5f0112fe.
//
// Solidity: function getCosmicSignatureNftStakingTotalEthRewardAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetCosmicSignatureNftStakingTotalEthRewardAmount() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetCosmicSignatureNftStakingTotalEthRewardAmount(&_CosmicSignatureGameV3.CallOpts)
}

// GetCstDutchAuctionDurations is a free data retrieval call binding the contract method 0xb700db5f.
//
// Solidity: function getCstDutchAuctionDurations() view returns(uint256, int256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetCstDutchAuctionDurations(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getCstDutchAuctionDurations")

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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetCstDutchAuctionDurations() (*big.Int, *big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetCstDutchAuctionDurations(&_CosmicSignatureGameV3.CallOpts)
}

// GetCstDutchAuctionDurations is a free data retrieval call binding the contract method 0xb700db5f.
//
// Solidity: function getCstDutchAuctionDurations() view returns(uint256, int256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetCstDutchAuctionDurations() (*big.Int, *big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetCstDutchAuctionDurations(&_CosmicSignatureGameV3.CallOpts)
}

// GetDurationElapsedSinceRoundActivation is a free data retrieval call binding the contract method 0x040d4d31.
//
// Solidity: function getDurationElapsedSinceRoundActivation() view returns(int256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetDurationElapsedSinceRoundActivation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getDurationElapsedSinceRoundActivation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDurationElapsedSinceRoundActivation is a free data retrieval call binding the contract method 0x040d4d31.
//
// Solidity: function getDurationElapsedSinceRoundActivation() view returns(int256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetDurationElapsedSinceRoundActivation() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetDurationElapsedSinceRoundActivation(&_CosmicSignatureGameV3.CallOpts)
}

// GetDurationElapsedSinceRoundActivation is a free data retrieval call binding the contract method 0x040d4d31.
//
// Solidity: function getDurationElapsedSinceRoundActivation() view returns(int256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetDurationElapsedSinceRoundActivation() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetDurationElapsedSinceRoundActivation(&_CosmicSignatureGameV3.CallOpts)
}

// GetDurationUntilMainPrize is a free data retrieval call binding the contract method 0x36750d2c.
//
// Solidity: function getDurationUntilMainPrize() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetDurationUntilMainPrize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getDurationUntilMainPrize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDurationUntilMainPrize is a free data retrieval call binding the contract method 0x36750d2c.
//
// Solidity: function getDurationUntilMainPrize() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetDurationUntilMainPrize() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetDurationUntilMainPrize(&_CosmicSignatureGameV3.CallOpts)
}

// GetDurationUntilMainPrize is a free data retrieval call binding the contract method 0x36750d2c.
//
// Solidity: function getDurationUntilMainPrize() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetDurationUntilMainPrize() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetDurationUntilMainPrize(&_CosmicSignatureGameV3.CallOpts)
}

// GetDurationUntilMainPrizeRaw is a free data retrieval call binding the contract method 0x37b99cc7.
//
// Solidity: function getDurationUntilMainPrizeRaw() view returns(int256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetDurationUntilMainPrizeRaw(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getDurationUntilMainPrizeRaw")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDurationUntilMainPrizeRaw is a free data retrieval call binding the contract method 0x37b99cc7.
//
// Solidity: function getDurationUntilMainPrizeRaw() view returns(int256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetDurationUntilMainPrizeRaw() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetDurationUntilMainPrizeRaw(&_CosmicSignatureGameV3.CallOpts)
}

// GetDurationUntilMainPrizeRaw is a free data retrieval call binding the contract method 0x37b99cc7.
//
// Solidity: function getDurationUntilMainPrizeRaw() view returns(int256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetDurationUntilMainPrizeRaw() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetDurationUntilMainPrizeRaw(&_CosmicSignatureGameV3.CallOpts)
}

// GetDurationUntilRoundActivation is a free data retrieval call binding the contract method 0xef22d15b.
//
// Solidity: function getDurationUntilRoundActivation() view returns(int256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetDurationUntilRoundActivation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getDurationUntilRoundActivation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDurationUntilRoundActivation is a free data retrieval call binding the contract method 0xef22d15b.
//
// Solidity: function getDurationUntilRoundActivation() view returns(int256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetDurationUntilRoundActivation() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetDurationUntilRoundActivation(&_CosmicSignatureGameV3.CallOpts)
}

// GetDurationUntilRoundActivation is a free data retrieval call binding the contract method 0xef22d15b.
//
// Solidity: function getDurationUntilRoundActivation() view returns(int256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetDurationUntilRoundActivation() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetDurationUntilRoundActivation(&_CosmicSignatureGameV3.CallOpts)
}

// GetEthDutchAuctionDurations is a free data retrieval call binding the contract method 0xfbaf5084.
//
// Solidity: function getEthDutchAuctionDurations() view returns(uint256, int256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetEthDutchAuctionDurations(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getEthDutchAuctionDurations")

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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetEthDutchAuctionDurations() (*big.Int, *big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetEthDutchAuctionDurations(&_CosmicSignatureGameV3.CallOpts)
}

// GetEthDutchAuctionDurations is a free data retrieval call binding the contract method 0xfbaf5084.
//
// Solidity: function getEthDutchAuctionDurations() view returns(uint256, int256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetEthDutchAuctionDurations() (*big.Int, *big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetEthDutchAuctionDurations(&_CosmicSignatureGameV3.CallOpts)
}

// GetEthPlusRandomWalkNftBidPrice is a free data retrieval call binding the contract method 0x27995f07.
//
// Solidity: function getEthPlusRandomWalkNftBidPrice(uint256 ethBidPrice_) pure returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetEthPlusRandomWalkNftBidPrice(opts *bind.CallOpts, ethBidPrice_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getEthPlusRandomWalkNftBidPrice", ethBidPrice_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthPlusRandomWalkNftBidPrice is a free data retrieval call binding the contract method 0x27995f07.
//
// Solidity: function getEthPlusRandomWalkNftBidPrice(uint256 ethBidPrice_) pure returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetEthPlusRandomWalkNftBidPrice(ethBidPrice_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetEthPlusRandomWalkNftBidPrice(&_CosmicSignatureGameV3.CallOpts, ethBidPrice_)
}

// GetEthPlusRandomWalkNftBidPrice is a free data retrieval call binding the contract method 0x27995f07.
//
// Solidity: function getEthPlusRandomWalkNftBidPrice(uint256 ethBidPrice_) pure returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetEthPlusRandomWalkNftBidPrice(ethBidPrice_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetEthPlusRandomWalkNftBidPrice(&_CosmicSignatureGameV3.CallOpts, ethBidPrice_)
}

// GetInitialDurationUntilMainPrize is a free data retrieval call binding the contract method 0x2b8dcbba.
//
// Solidity: function getInitialDurationUntilMainPrize() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetInitialDurationUntilMainPrize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getInitialDurationUntilMainPrize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitialDurationUntilMainPrize is a free data retrieval call binding the contract method 0x2b8dcbba.
//
// Solidity: function getInitialDurationUntilMainPrize() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetInitialDurationUntilMainPrize() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetInitialDurationUntilMainPrize(&_CosmicSignatureGameV3.CallOpts)
}

// GetInitialDurationUntilMainPrize is a free data retrieval call binding the contract method 0x2b8dcbba.
//
// Solidity: function getInitialDurationUntilMainPrize() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetInitialDurationUntilMainPrize() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetInitialDurationUntilMainPrize(&_CosmicSignatureGameV3.CallOpts)
}

// GetMainEthPrizeAmount is a free data retrieval call binding the contract method 0x5b0a45d9.
//
// Solidity: function getMainEthPrizeAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetMainEthPrizeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getMainEthPrizeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMainEthPrizeAmount is a free data retrieval call binding the contract method 0x5b0a45d9.
//
// Solidity: function getMainEthPrizeAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetMainEthPrizeAmount() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetMainEthPrizeAmount(&_CosmicSignatureGameV3.CallOpts)
}

// GetMainEthPrizeAmount is a free data retrieval call binding the contract method 0x5b0a45d9.
//
// Solidity: function getMainEthPrizeAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetMainEthPrizeAmount() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetMainEthPrizeAmount(&_CosmicSignatureGameV3.CallOpts)
}

// GetMainPrizeTimeIncrement is a free data retrieval call binding the contract method 0x4f734612.
//
// Solidity: function getMainPrizeTimeIncrement() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetMainPrizeTimeIncrement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getMainPrizeTimeIncrement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMainPrizeTimeIncrement is a free data retrieval call binding the contract method 0x4f734612.
//
// Solidity: function getMainPrizeTimeIncrement() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetMainPrizeTimeIncrement() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetMainPrizeTimeIncrement(&_CosmicSignatureGameV3.CallOpts)
}

// GetMainPrizeTimeIncrement is a free data retrieval call binding the contract method 0x4f734612.
//
// Solidity: function getMainPrizeTimeIncrement() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetMainPrizeTimeIncrement() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetMainPrizeTimeIncrement(&_CosmicSignatureGameV3.CallOpts)
}

// GetNextCstBidPrice is a free data retrieval call binding the contract method 0x6e95d286.
//
// Solidity: function getNextCstBidPrice() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetNextCstBidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getNextCstBidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextCstBidPrice is a free data retrieval call binding the contract method 0x6e95d286.
//
// Solidity: function getNextCstBidPrice() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetNextCstBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetNextCstBidPrice(&_CosmicSignatureGameV3.CallOpts)
}

// GetNextCstBidPrice is a free data retrieval call binding the contract method 0x6e95d286.
//
// Solidity: function getNextCstBidPrice() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetNextCstBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetNextCstBidPrice(&_CosmicSignatureGameV3.CallOpts)
}

// GetNextCstBidPriceAdvanced is a free data retrieval call binding the contract method 0xb6a94f42.
//
// Solidity: function getNextCstBidPriceAdvanced(int256 currentTimeOffset_) view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetNextCstBidPriceAdvanced(opts *bind.CallOpts, currentTimeOffset_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getNextCstBidPriceAdvanced", currentTimeOffset_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextCstBidPriceAdvanced is a free data retrieval call binding the contract method 0xb6a94f42.
//
// Solidity: function getNextCstBidPriceAdvanced(int256 currentTimeOffset_) view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetNextCstBidPriceAdvanced(currentTimeOffset_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetNextCstBidPriceAdvanced(&_CosmicSignatureGameV3.CallOpts, currentTimeOffset_)
}

// GetNextCstBidPriceAdvanced is a free data retrieval call binding the contract method 0xb6a94f42.
//
// Solidity: function getNextCstBidPriceAdvanced(int256 currentTimeOffset_) view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetNextCstBidPriceAdvanced(currentTimeOffset_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetNextCstBidPriceAdvanced(&_CosmicSignatureGameV3.CallOpts, currentTimeOffset_)
}

// GetNextEthBidPrice is a free data retrieval call binding the contract method 0x62ed9b53.
//
// Solidity: function getNextEthBidPrice() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetNextEthBidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getNextEthBidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextEthBidPrice is a free data retrieval call binding the contract method 0x62ed9b53.
//
// Solidity: function getNextEthBidPrice() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetNextEthBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetNextEthBidPrice(&_CosmicSignatureGameV3.CallOpts)
}

// GetNextEthBidPrice is a free data retrieval call binding the contract method 0x62ed9b53.
//
// Solidity: function getNextEthBidPrice() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetNextEthBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetNextEthBidPrice(&_CosmicSignatureGameV3.CallOpts)
}

// GetNextEthBidPriceAdvanced is a free data retrieval call binding the contract method 0x4e452010.
//
// Solidity: function getNextEthBidPriceAdvanced(int256 currentTimeOffset_) view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetNextEthBidPriceAdvanced(opts *bind.CallOpts, currentTimeOffset_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getNextEthBidPriceAdvanced", currentTimeOffset_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextEthBidPriceAdvanced is a free data retrieval call binding the contract method 0x4e452010.
//
// Solidity: function getNextEthBidPriceAdvanced(int256 currentTimeOffset_) view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetNextEthBidPriceAdvanced(currentTimeOffset_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetNextEthBidPriceAdvanced(&_CosmicSignatureGameV3.CallOpts, currentTimeOffset_)
}

// GetNextEthBidPriceAdvanced is a free data retrieval call binding the contract method 0x4e452010.
//
// Solidity: function getNextEthBidPriceAdvanced(int256 currentTimeOffset_) view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetNextEthBidPriceAdvanced(currentTimeOffset_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetNextEthBidPriceAdvanced(&_CosmicSignatureGameV3.CallOpts, currentTimeOffset_)
}

// GetRaffleTotalEthPrizeAmountForBidders is a free data retrieval call binding the contract method 0xa35286d1.
//
// Solidity: function getRaffleTotalEthPrizeAmountForBidders() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetRaffleTotalEthPrizeAmountForBidders(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getRaffleTotalEthPrizeAmountForBidders")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRaffleTotalEthPrizeAmountForBidders is a free data retrieval call binding the contract method 0xa35286d1.
//
// Solidity: function getRaffleTotalEthPrizeAmountForBidders() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetRaffleTotalEthPrizeAmountForBidders() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetRaffleTotalEthPrizeAmountForBidders(&_CosmicSignatureGameV3.CallOpts)
}

// GetRaffleTotalEthPrizeAmountForBidders is a free data retrieval call binding the contract method 0xa35286d1.
//
// Solidity: function getRaffleTotalEthPrizeAmountForBidders() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetRaffleTotalEthPrizeAmountForBidders() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetRaffleTotalEthPrizeAmountForBidders(&_CosmicSignatureGameV3.CallOpts)
}

// GetRoundLateBidDuration is a free data retrieval call binding the contract method 0x60ef8841.
//
// Solidity: function getRoundLateBidDuration() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetRoundLateBidDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getRoundLateBidDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoundLateBidDuration is a free data retrieval call binding the contract method 0x60ef8841.
//
// Solidity: function getRoundLateBidDuration() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetRoundLateBidDuration() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetRoundLateBidDuration(&_CosmicSignatureGameV3.CallOpts)
}

// GetRoundLateBidDuration is a free data retrieval call binding the contract method 0x60ef8841.
//
// Solidity: function getRoundLateBidDuration() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetRoundLateBidDuration() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetRoundLateBidDuration(&_CosmicSignatureGameV3.CallOpts)
}

// GetTotalNumBids is a free data retrieval call binding the contract method 0xfd77310f.
//
// Solidity: function getTotalNumBids(uint256 roundNum_) view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetTotalNumBids(opts *bind.CallOpts, roundNum_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getTotalNumBids", roundNum_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalNumBids is a free data retrieval call binding the contract method 0xfd77310f.
//
// Solidity: function getTotalNumBids(uint256 roundNum_) view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetTotalNumBids(roundNum_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetTotalNumBids(&_CosmicSignatureGameV3.CallOpts, roundNum_)
}

// GetTotalNumBids is a free data retrieval call binding the contract method 0xfd77310f.
//
// Solidity: function getTotalNumBids(uint256 roundNum_) view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetTotalNumBids(roundNum_ *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetTotalNumBids(&_CosmicSignatureGameV3.CallOpts, roundNum_)
}

// InitialDurationUntilMainPrizeDivisor is a free data retrieval call binding the contract method 0x44a4b917.
//
// Solidity: function initialDurationUntilMainPrizeDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) InitialDurationUntilMainPrizeDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "initialDurationUntilMainPrizeDivisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialDurationUntilMainPrizeDivisor is a free data retrieval call binding the contract method 0x44a4b917.
//
// Solidity: function initialDurationUntilMainPrizeDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) InitialDurationUntilMainPrizeDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.InitialDurationUntilMainPrizeDivisor(&_CosmicSignatureGameV3.CallOpts)
}

// InitialDurationUntilMainPrizeDivisor is a free data retrieval call binding the contract method 0x44a4b917.
//
// Solidity: function initialDurationUntilMainPrizeDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) InitialDurationUntilMainPrizeDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.InitialDurationUntilMainPrizeDivisor(&_CosmicSignatureGameV3.CallOpts)
}

// LastBidderAddress is a free data retrieval call binding the contract method 0xe5b3cd14.
//
// Solidity: function lastBidderAddress() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) LastBidderAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "lastBidderAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastBidderAddress is a free data retrieval call binding the contract method 0xe5b3cd14.
//
// Solidity: function lastBidderAddress() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) LastBidderAddress() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.LastBidderAddress(&_CosmicSignatureGameV3.CallOpts)
}

// LastBidderAddress is a free data retrieval call binding the contract method 0xe5b3cd14.
//
// Solidity: function lastBidderAddress() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) LastBidderAddress() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.LastBidderAddress(&_CosmicSignatureGameV3.CallOpts)
}

// LastCstBidderAddress is a free data retrieval call binding the contract method 0xad4b0e8a.
//
// Solidity: function lastCstBidderAddress() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) LastCstBidderAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "lastCstBidderAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastCstBidderAddress is a free data retrieval call binding the contract method 0xad4b0e8a.
//
// Solidity: function lastCstBidderAddress() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) LastCstBidderAddress() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.LastCstBidderAddress(&_CosmicSignatureGameV3.CallOpts)
}

// LastCstBidderAddress is a free data retrieval call binding the contract method 0xad4b0e8a.
//
// Solidity: function lastCstBidderAddress() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) LastCstBidderAddress() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.LastCstBidderAddress(&_CosmicSignatureGameV3.CallOpts)
}

// MainEthPrizeAmountPercentage is a free data retrieval call binding the contract method 0xf444b298.
//
// Solidity: function mainEthPrizeAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) MainEthPrizeAmountPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "mainEthPrizeAmountPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MainEthPrizeAmountPercentage is a free data retrieval call binding the contract method 0xf444b298.
//
// Solidity: function mainEthPrizeAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) MainEthPrizeAmountPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.MainEthPrizeAmountPercentage(&_CosmicSignatureGameV3.CallOpts)
}

// MainEthPrizeAmountPercentage is a free data retrieval call binding the contract method 0xf444b298.
//
// Solidity: function mainEthPrizeAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) MainEthPrizeAmountPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.MainEthPrizeAmountPercentage(&_CosmicSignatureGameV3.CallOpts)
}

// MainPrizeNumCosmicSignatureNfts is a free data retrieval call binding the contract method 0x5fdf49cb.
//
// Solidity: function mainPrizeNumCosmicSignatureNfts() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) MainPrizeNumCosmicSignatureNfts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "mainPrizeNumCosmicSignatureNfts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MainPrizeNumCosmicSignatureNfts is a free data retrieval call binding the contract method 0x5fdf49cb.
//
// Solidity: function mainPrizeNumCosmicSignatureNfts() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) MainPrizeNumCosmicSignatureNfts() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.MainPrizeNumCosmicSignatureNfts(&_CosmicSignatureGameV3.CallOpts)
}

// MainPrizeNumCosmicSignatureNfts is a free data retrieval call binding the contract method 0x5fdf49cb.
//
// Solidity: function mainPrizeNumCosmicSignatureNfts() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) MainPrizeNumCosmicSignatureNfts() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.MainPrizeNumCosmicSignatureNfts(&_CosmicSignatureGameV3.CallOpts)
}

// MainPrizeTime is a free data retrieval call binding the contract method 0x18305de2.
//
// Solidity: function mainPrizeTime() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) MainPrizeTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "mainPrizeTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MainPrizeTime is a free data retrieval call binding the contract method 0x18305de2.
//
// Solidity: function mainPrizeTime() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) MainPrizeTime() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.MainPrizeTime(&_CosmicSignatureGameV3.CallOpts)
}

// MainPrizeTime is a free data retrieval call binding the contract method 0x18305de2.
//
// Solidity: function mainPrizeTime() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) MainPrizeTime() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.MainPrizeTime(&_CosmicSignatureGameV3.CallOpts)
}

// MainPrizeTimeIncrementInMicroSeconds is a free data retrieval call binding the contract method 0xeb13430e.
//
// Solidity: function mainPrizeTimeIncrementInMicroSeconds() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) MainPrizeTimeIncrementInMicroSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "mainPrizeTimeIncrementInMicroSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MainPrizeTimeIncrementInMicroSeconds is a free data retrieval call binding the contract method 0xeb13430e.
//
// Solidity: function mainPrizeTimeIncrementInMicroSeconds() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) MainPrizeTimeIncrementInMicroSeconds() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.MainPrizeTimeIncrementInMicroSeconds(&_CosmicSignatureGameV3.CallOpts)
}

// MainPrizeTimeIncrementInMicroSeconds is a free data retrieval call binding the contract method 0xeb13430e.
//
// Solidity: function mainPrizeTimeIncrementInMicroSeconds() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) MainPrizeTimeIncrementInMicroSeconds() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.MainPrizeTimeIncrementInMicroSeconds(&_CosmicSignatureGameV3.CallOpts)
}

// MainPrizeTimeIncrementIncreaseDivisor is a free data retrieval call binding the contract method 0x56732241.
//
// Solidity: function mainPrizeTimeIncrementIncreaseDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) MainPrizeTimeIncrementIncreaseDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "mainPrizeTimeIncrementIncreaseDivisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MainPrizeTimeIncrementIncreaseDivisor is a free data retrieval call binding the contract method 0x56732241.
//
// Solidity: function mainPrizeTimeIncrementIncreaseDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) MainPrizeTimeIncrementIncreaseDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.MainPrizeTimeIncrementIncreaseDivisor(&_CosmicSignatureGameV3.CallOpts)
}

// MainPrizeTimeIncrementIncreaseDivisor is a free data retrieval call binding the contract method 0x56732241.
//
// Solidity: function mainPrizeTimeIncrementIncreaseDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) MainPrizeTimeIncrementIncreaseDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.MainPrizeTimeIncrementIncreaseDivisor(&_CosmicSignatureGameV3.CallOpts)
}

// MarketingWallet is a free data retrieval call binding the contract method 0x75f0a874.
//
// Solidity: function marketingWallet() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) MarketingWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "marketingWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketingWallet is a free data retrieval call binding the contract method 0x75f0a874.
//
// Solidity: function marketingWallet() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) MarketingWallet() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.MarketingWallet(&_CosmicSignatureGameV3.CallOpts)
}

// MarketingWallet is a free data retrieval call binding the contract method 0x75f0a874.
//
// Solidity: function marketingWallet() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) MarketingWallet() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.MarketingWallet(&_CosmicSignatureGameV3.CallOpts)
}

// MarketingWalletCstContributionAmount is a free data retrieval call binding the contract method 0x4164b95b.
//
// Solidity: function marketingWalletCstContributionAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) MarketingWalletCstContributionAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "marketingWalletCstContributionAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketingWalletCstContributionAmount is a free data retrieval call binding the contract method 0x4164b95b.
//
// Solidity: function marketingWalletCstContributionAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) MarketingWalletCstContributionAmount() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.MarketingWalletCstContributionAmount(&_CosmicSignatureGameV3.CallOpts)
}

// MarketingWalletCstContributionAmount is a free data retrieval call binding the contract method 0x4164b95b.
//
// Solidity: function marketingWalletCstContributionAmount() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) MarketingWalletCstContributionAmount() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.MarketingWalletCstContributionAmount(&_CosmicSignatureGameV3.CallOpts)
}

// NextEthBidPrice is a free data retrieval call binding the contract method 0xefeb248a.
//
// Solidity: function nextEthBidPrice() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) NextEthBidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "nextEthBidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextEthBidPrice is a free data retrieval call binding the contract method 0xefeb248a.
//
// Solidity: function nextEthBidPrice() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) NextEthBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.NextEthBidPrice(&_CosmicSignatureGameV3.CallOpts)
}

// NextEthBidPrice is a free data retrieval call binding the contract method 0xefeb248a.
//
// Solidity: function nextEthBidPrice() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) NextEthBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.NextEthBidPrice(&_CosmicSignatureGameV3.CallOpts)
}

// NextRoundFirstCstDutchAuctionBeginningBidPrice is a free data retrieval call binding the contract method 0xc7e7a601.
//
// Solidity: function nextRoundFirstCstDutchAuctionBeginningBidPrice() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) NextRoundFirstCstDutchAuctionBeginningBidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "nextRoundFirstCstDutchAuctionBeginningBidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRoundFirstCstDutchAuctionBeginningBidPrice is a free data retrieval call binding the contract method 0xc7e7a601.
//
// Solidity: function nextRoundFirstCstDutchAuctionBeginningBidPrice() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) NextRoundFirstCstDutchAuctionBeginningBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.NextRoundFirstCstDutchAuctionBeginningBidPrice(&_CosmicSignatureGameV3.CallOpts)
}

// NextRoundFirstCstDutchAuctionBeginningBidPrice is a free data retrieval call binding the contract method 0xc7e7a601.
//
// Solidity: function nextRoundFirstCstDutchAuctionBeginningBidPrice() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) NextRoundFirstCstDutchAuctionBeginningBidPrice() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.NextRoundFirstCstDutchAuctionBeginningBidPrice(&_CosmicSignatureGameV3.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) Nft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "nft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) Nft() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.Nft(&_CosmicSignatureGameV3.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) Nft() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.Nft(&_CosmicSignatureGameV3.CallOpts)
}

// NumEthDonationWithInfoRecords is a free data retrieval call binding the contract method 0x0b5f95ae.
//
// Solidity: function numEthDonationWithInfoRecords() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) NumEthDonationWithInfoRecords(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "numEthDonationWithInfoRecords")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumEthDonationWithInfoRecords is a free data retrieval call binding the contract method 0x0b5f95ae.
//
// Solidity: function numEthDonationWithInfoRecords() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) NumEthDonationWithInfoRecords() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.NumEthDonationWithInfoRecords(&_CosmicSignatureGameV3.CallOpts)
}

// NumEthDonationWithInfoRecords is a free data retrieval call binding the contract method 0x0b5f95ae.
//
// Solidity: function numEthDonationWithInfoRecords() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) NumEthDonationWithInfoRecords() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.NumEthDonationWithInfoRecords(&_CosmicSignatureGameV3.CallOpts)
}

// NumRaffleCosmicSignatureNftsForBidders is a free data retrieval call binding the contract method 0x4c2a4a33.
//
// Solidity: function numRaffleCosmicSignatureNftsForBidders() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) NumRaffleCosmicSignatureNftsForBidders(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "numRaffleCosmicSignatureNftsForBidders")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRaffleCosmicSignatureNftsForBidders is a free data retrieval call binding the contract method 0x4c2a4a33.
//
// Solidity: function numRaffleCosmicSignatureNftsForBidders() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) NumRaffleCosmicSignatureNftsForBidders() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.NumRaffleCosmicSignatureNftsForBidders(&_CosmicSignatureGameV3.CallOpts)
}

// NumRaffleCosmicSignatureNftsForBidders is a free data retrieval call binding the contract method 0x4c2a4a33.
//
// Solidity: function numRaffleCosmicSignatureNftsForBidders() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) NumRaffleCosmicSignatureNftsForBidders() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.NumRaffleCosmicSignatureNftsForBidders(&_CosmicSignatureGameV3.CallOpts)
}

// NumRaffleCosmicSignatureNftsForRandomWalkNftStakers is a free data retrieval call binding the contract method 0xe2f9185f.
//
// Solidity: function numRaffleCosmicSignatureNftsForRandomWalkNftStakers() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) NumRaffleCosmicSignatureNftsForRandomWalkNftStakers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "numRaffleCosmicSignatureNftsForRandomWalkNftStakers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRaffleCosmicSignatureNftsForRandomWalkNftStakers is a free data retrieval call binding the contract method 0xe2f9185f.
//
// Solidity: function numRaffleCosmicSignatureNftsForRandomWalkNftStakers() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) NumRaffleCosmicSignatureNftsForRandomWalkNftStakers() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.NumRaffleCosmicSignatureNftsForRandomWalkNftStakers(&_CosmicSignatureGameV3.CallOpts)
}

// NumRaffleCosmicSignatureNftsForRandomWalkNftStakers is a free data retrieval call binding the contract method 0xe2f9185f.
//
// Solidity: function numRaffleCosmicSignatureNftsForRandomWalkNftStakers() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) NumRaffleCosmicSignatureNftsForRandomWalkNftStakers() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.NumRaffleCosmicSignatureNftsForRandomWalkNftStakers(&_CosmicSignatureGameV3.CallOpts)
}

// NumRaffleEthPrizesForBidders is a free data retrieval call binding the contract method 0xf11400f0.
//
// Solidity: function numRaffleEthPrizesForBidders() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) NumRaffleEthPrizesForBidders(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "numRaffleEthPrizesForBidders")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRaffleEthPrizesForBidders is a free data retrieval call binding the contract method 0xf11400f0.
//
// Solidity: function numRaffleEthPrizesForBidders() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) NumRaffleEthPrizesForBidders() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.NumRaffleEthPrizesForBidders(&_CosmicSignatureGameV3.CallOpts)
}

// NumRaffleEthPrizesForBidders is a free data retrieval call binding the contract method 0xf11400f0.
//
// Solidity: function numRaffleEthPrizesForBidders() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) NumRaffleEthPrizesForBidders() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.NumRaffleEthPrizesForBidders(&_CosmicSignatureGameV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) Owner() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.Owner(&_CosmicSignatureGameV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) Owner() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.Owner(&_CosmicSignatureGameV3.CallOpts)
}

// PrevEnduranceChampionDuration is a free data retrieval call binding the contract method 0xeaace302.
//
// Solidity: function prevEnduranceChampionDuration() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) PrevEnduranceChampionDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "prevEnduranceChampionDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrevEnduranceChampionDuration is a free data retrieval call binding the contract method 0xeaace302.
//
// Solidity: function prevEnduranceChampionDuration() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) PrevEnduranceChampionDuration() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.PrevEnduranceChampionDuration(&_CosmicSignatureGameV3.CallOpts)
}

// PrevEnduranceChampionDuration is a free data retrieval call binding the contract method 0xeaace302.
//
// Solidity: function prevEnduranceChampionDuration() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) PrevEnduranceChampionDuration() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.PrevEnduranceChampionDuration(&_CosmicSignatureGameV3.CallOpts)
}

// PrizesWallet is a free data retrieval call binding the contract method 0x2afa2580.
//
// Solidity: function prizesWallet() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) PrizesWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "prizesWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrizesWallet is a free data retrieval call binding the contract method 0x2afa2580.
//
// Solidity: function prizesWallet() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) PrizesWallet() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.PrizesWallet(&_CosmicSignatureGameV3.CallOpts)
}

// PrizesWallet is a free data retrieval call binding the contract method 0x2afa2580.
//
// Solidity: function prizesWallet() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) PrizesWallet() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.PrizesWallet(&_CosmicSignatureGameV3.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) ProxiableUUID() ([32]byte, error) {
	return _CosmicSignatureGameV3.Contract.ProxiableUUID(&_CosmicSignatureGameV3.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) ProxiableUUID() ([32]byte, error) {
	return _CosmicSignatureGameV3.Contract.ProxiableUUID(&_CosmicSignatureGameV3.CallOpts)
}

// RaffleTotalEthPrizeAmountForBiddersPercentage is a free data retrieval call binding the contract method 0x477adf2a.
//
// Solidity: function raffleTotalEthPrizeAmountForBiddersPercentage() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) RaffleTotalEthPrizeAmountForBiddersPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "raffleTotalEthPrizeAmountForBiddersPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaffleTotalEthPrizeAmountForBiddersPercentage is a free data retrieval call binding the contract method 0x477adf2a.
//
// Solidity: function raffleTotalEthPrizeAmountForBiddersPercentage() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) RaffleTotalEthPrizeAmountForBiddersPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.RaffleTotalEthPrizeAmountForBiddersPercentage(&_CosmicSignatureGameV3.CallOpts)
}

// RaffleTotalEthPrizeAmountForBiddersPercentage is a free data retrieval call binding the contract method 0x477adf2a.
//
// Solidity: function raffleTotalEthPrizeAmountForBiddersPercentage() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) RaffleTotalEthPrizeAmountForBiddersPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.RaffleTotalEthPrizeAmountForBiddersPercentage(&_CosmicSignatureGameV3.CallOpts)
}

// RandomWalkNft is a free data retrieval call binding the contract method 0x755b4ef7.
//
// Solidity: function randomWalkNft() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) RandomWalkNft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "randomWalkNft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RandomWalkNft is a free data retrieval call binding the contract method 0x755b4ef7.
//
// Solidity: function randomWalkNft() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) RandomWalkNft() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.RandomWalkNft(&_CosmicSignatureGameV3.CallOpts)
}

// RandomWalkNft is a free data retrieval call binding the contract method 0x755b4ef7.
//
// Solidity: function randomWalkNft() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) RandomWalkNft() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.RandomWalkNft(&_CosmicSignatureGameV3.CallOpts)
}

// RoundActivationTime is a free data retrieval call binding the contract method 0x6e970834.
//
// Solidity: function roundActivationTime() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) RoundActivationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "roundActivationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundActivationTime is a free data retrieval call binding the contract method 0x6e970834.
//
// Solidity: function roundActivationTime() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) RoundActivationTime() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.RoundActivationTime(&_CosmicSignatureGameV3.CallOpts)
}

// RoundActivationTime is a free data retrieval call binding the contract method 0x6e970834.
//
// Solidity: function roundActivationTime() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) RoundActivationTime() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.RoundActivationTime(&_CosmicSignatureGameV3.CallOpts)
}

// RoundLateBidDurationDivisor is a free data retrieval call binding the contract method 0x71b6d019.
//
// Solidity: function roundLateBidDurationDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) RoundLateBidDurationDivisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "roundLateBidDurationDivisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundLateBidDurationDivisor is a free data retrieval call binding the contract method 0x71b6d019.
//
// Solidity: function roundLateBidDurationDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) RoundLateBidDurationDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.RoundLateBidDurationDivisor(&_CosmicSignatureGameV3.CallOpts)
}

// RoundLateBidDurationDivisor is a free data retrieval call binding the contract method 0x71b6d019.
//
// Solidity: function roundLateBidDurationDivisor() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) RoundLateBidDurationDivisor() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.RoundLateBidDurationDivisor(&_CosmicSignatureGameV3.CallOpts)
}

// RoundLateBidPricePremiumAmountBaseMultiplier is a free data retrieval call binding the contract method 0x99bf353d.
//
// Solidity: function roundLateBidPricePremiumAmountBaseMultiplier() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) RoundLateBidPricePremiumAmountBaseMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "roundLateBidPricePremiumAmountBaseMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundLateBidPricePremiumAmountBaseMultiplier is a free data retrieval call binding the contract method 0x99bf353d.
//
// Solidity: function roundLateBidPricePremiumAmountBaseMultiplier() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) RoundLateBidPricePremiumAmountBaseMultiplier() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.RoundLateBidPricePremiumAmountBaseMultiplier(&_CosmicSignatureGameV3.CallOpts)
}

// RoundLateBidPricePremiumAmountBaseMultiplier is a free data retrieval call binding the contract method 0x99bf353d.
//
// Solidity: function roundLateBidPricePremiumAmountBaseMultiplier() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) RoundLateBidPricePremiumAmountBaseMultiplier() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.RoundLateBidPricePremiumAmountBaseMultiplier(&_CosmicSignatureGameV3.CallOpts)
}

// RoundLateBidPricePremiumAmountExponent is a free data retrieval call binding the contract method 0xc52d8549.
//
// Solidity: function roundLateBidPricePremiumAmountExponent() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) RoundLateBidPricePremiumAmountExponent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "roundLateBidPricePremiumAmountExponent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundLateBidPricePremiumAmountExponent is a free data retrieval call binding the contract method 0xc52d8549.
//
// Solidity: function roundLateBidPricePremiumAmountExponent() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) RoundLateBidPricePremiumAmountExponent() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.RoundLateBidPricePremiumAmountExponent(&_CosmicSignatureGameV3.CallOpts)
}

// RoundLateBidPricePremiumAmountExponent is a free data retrieval call binding the contract method 0xc52d8549.
//
// Solidity: function roundLateBidPricePremiumAmountExponent() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) RoundLateBidPricePremiumAmountExponent() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.RoundLateBidPricePremiumAmountExponent(&_CosmicSignatureGameV3.CallOpts)
}

// RoundNum is a free data retrieval call binding the contract method 0x119b22b3.
//
// Solidity: function roundNum() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) RoundNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "roundNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundNum is a free data retrieval call binding the contract method 0x119b22b3.
//
// Solidity: function roundNum() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) RoundNum() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.RoundNum(&_CosmicSignatureGameV3.CallOpts)
}

// RoundNum is a free data retrieval call binding the contract method 0x119b22b3.
//
// Solidity: function roundNum() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) RoundNum() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.RoundNum(&_CosmicSignatureGameV3.CallOpts)
}

// StakingWalletCosmicSignatureNft is a free data retrieval call binding the contract method 0xa922ab5d.
//
// Solidity: function stakingWalletCosmicSignatureNft() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) StakingWalletCosmicSignatureNft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "stakingWalletCosmicSignatureNft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingWalletCosmicSignatureNft is a free data retrieval call binding the contract method 0xa922ab5d.
//
// Solidity: function stakingWalletCosmicSignatureNft() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) StakingWalletCosmicSignatureNft() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.StakingWalletCosmicSignatureNft(&_CosmicSignatureGameV3.CallOpts)
}

// StakingWalletCosmicSignatureNft is a free data retrieval call binding the contract method 0xa922ab5d.
//
// Solidity: function stakingWalletCosmicSignatureNft() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) StakingWalletCosmicSignatureNft() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.StakingWalletCosmicSignatureNft(&_CosmicSignatureGameV3.CallOpts)
}

// StakingWalletRandomWalkNft is a free data retrieval call binding the contract method 0x2d809e88.
//
// Solidity: function stakingWalletRandomWalkNft() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) StakingWalletRandomWalkNft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "stakingWalletRandomWalkNft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingWalletRandomWalkNft is a free data retrieval call binding the contract method 0x2d809e88.
//
// Solidity: function stakingWalletRandomWalkNft() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) StakingWalletRandomWalkNft() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.StakingWalletRandomWalkNft(&_CosmicSignatureGameV3.CallOpts)
}

// StakingWalletRandomWalkNft is a free data retrieval call binding the contract method 0x2d809e88.
//
// Solidity: function stakingWalletRandomWalkNft() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) StakingWalletRandomWalkNft() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.StakingWalletRandomWalkNft(&_CosmicSignatureGameV3.CallOpts)
}

// TimeoutDurationToClaimMainPrize is a free data retrieval call binding the contract method 0x3b9d292e.
//
// Solidity: function timeoutDurationToClaimMainPrize() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) TimeoutDurationToClaimMainPrize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "timeoutDurationToClaimMainPrize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeoutDurationToClaimMainPrize is a free data retrieval call binding the contract method 0x3b9d292e.
//
// Solidity: function timeoutDurationToClaimMainPrize() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) TimeoutDurationToClaimMainPrize() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.TimeoutDurationToClaimMainPrize(&_CosmicSignatureGameV3.CallOpts)
}

// TimeoutDurationToClaimMainPrize is a free data retrieval call binding the contract method 0x3b9d292e.
//
// Solidity: function timeoutDurationToClaimMainPrize() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) TimeoutDurationToClaimMainPrize() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.TimeoutDurationToClaimMainPrize(&_CosmicSignatureGameV3.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) Token() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.Token(&_CosmicSignatureGameV3.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) Token() (common.Address, error) {
	return _CosmicSignatureGameV3.Contract.Token(&_CosmicSignatureGameV3.CallOpts)
}

// TryGetCurrentChampions is a free data retrieval call binding the contract method 0xcb720d4d.
//
// Solidity: function tryGetCurrentChampions() view returns(address enduranceChampionAddress_, uint256 enduranceChampionDuration_, address chronoWarriorAddress_, uint256 chronoWarriorDuration_)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) TryGetCurrentChampions(opts *bind.CallOpts) (struct {
	EnduranceChampionAddress  common.Address
	EnduranceChampionDuration *big.Int
	ChronoWarriorAddress      common.Address
	ChronoWarriorDuration     *big.Int
}, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "tryGetCurrentChampions")

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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) TryGetCurrentChampions() (struct {
	EnduranceChampionAddress  common.Address
	EnduranceChampionDuration *big.Int
	ChronoWarriorAddress      common.Address
	ChronoWarriorDuration     *big.Int
}, error) {
	return _CosmicSignatureGameV3.Contract.TryGetCurrentChampions(&_CosmicSignatureGameV3.CallOpts)
}

// TryGetCurrentChampions is a free data retrieval call binding the contract method 0xcb720d4d.
//
// Solidity: function tryGetCurrentChampions() view returns(address enduranceChampionAddress_, uint256 enduranceChampionDuration_, address chronoWarriorAddress_, uint256 chronoWarriorDuration_)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) TryGetCurrentChampions() (struct {
	EnduranceChampionAddress  common.Address
	EnduranceChampionDuration *big.Int
	ChronoWarriorAddress      common.Address
	ChronoWarriorDuration     *big.Int
}, error) {
	return _CosmicSignatureGameV3.Contract.TryGetCurrentChampions(&_CosmicSignatureGameV3.CallOpts)
}

// UsedRandomWalkNfts is a free data retrieval call binding the contract method 0xebb9bc5c.
//
// Solidity: function usedRandomWalkNfts(uint256 nftId) view returns(uint256 nftWasUsed)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) UsedRandomWalkNfts(opts *bind.CallOpts, nftId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "usedRandomWalkNfts", nftId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsedRandomWalkNfts is a free data retrieval call binding the contract method 0xebb9bc5c.
//
// Solidity: function usedRandomWalkNfts(uint256 nftId) view returns(uint256 nftWasUsed)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) UsedRandomWalkNfts(nftId *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.UsedRandomWalkNfts(&_CosmicSignatureGameV3.CallOpts, nftId)
}

// UsedRandomWalkNfts is a free data retrieval call binding the contract method 0xebb9bc5c.
//
// Solidity: function usedRandomWalkNfts(uint256 nftId) view returns(uint256 nftWasUsed)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) UsedRandomWalkNfts(nftId *big.Int) (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.UsedRandomWalkNfts(&_CosmicSignatureGameV3.CallOpts, nftId)
}

// BidWithCst is a paid mutator transaction binding the contract method 0xd7559b9c.
//
// Solidity: function bidWithCst(uint256 priceMaxLimit_, string message_, uint256 bidCstRewardAmountMinLimit_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) BidWithCst(opts *bind.TransactOpts, priceMaxLimit_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "bidWithCst", priceMaxLimit_, message_, bidCstRewardAmountMinLimit_)
}

// BidWithCst is a paid mutator transaction binding the contract method 0xd7559b9c.
//
// Solidity: function bidWithCst(uint256 priceMaxLimit_, string message_, uint256 bidCstRewardAmountMinLimit_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) BidWithCst(priceMaxLimit_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.BidWithCst(&_CosmicSignatureGameV3.TransactOpts, priceMaxLimit_, message_, bidCstRewardAmountMinLimit_)
}

// BidWithCst is a paid mutator transaction binding the contract method 0xd7559b9c.
//
// Solidity: function bidWithCst(uint256 priceMaxLimit_, string message_, uint256 bidCstRewardAmountMinLimit_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) BidWithCst(priceMaxLimit_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.BidWithCst(&_CosmicSignatureGameV3.TransactOpts, priceMaxLimit_, message_, bidCstRewardAmountMinLimit_)
}

// BidWithCstAndDonateNft is a paid mutator transaction binding the contract method 0x329b95a5.
//
// Solidity: function bidWithCstAndDonateNft(uint256 priceMaxLimit_, string message_, uint256 bidCstRewardAmountMinLimit_, address nftAddress_, uint256 nftId_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) BidWithCstAndDonateNft(opts *bind.TransactOpts, priceMaxLimit_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "bidWithCstAndDonateNft", priceMaxLimit_, message_, bidCstRewardAmountMinLimit_, nftAddress_, nftId_)
}

// BidWithCstAndDonateNft is a paid mutator transaction binding the contract method 0x329b95a5.
//
// Solidity: function bidWithCstAndDonateNft(uint256 priceMaxLimit_, string message_, uint256 bidCstRewardAmountMinLimit_, address nftAddress_, uint256 nftId_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) BidWithCstAndDonateNft(priceMaxLimit_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.BidWithCstAndDonateNft(&_CosmicSignatureGameV3.TransactOpts, priceMaxLimit_, message_, bidCstRewardAmountMinLimit_, nftAddress_, nftId_)
}

// BidWithCstAndDonateNft is a paid mutator transaction binding the contract method 0x329b95a5.
//
// Solidity: function bidWithCstAndDonateNft(uint256 priceMaxLimit_, string message_, uint256 bidCstRewardAmountMinLimit_, address nftAddress_, uint256 nftId_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) BidWithCstAndDonateNft(priceMaxLimit_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.BidWithCstAndDonateNft(&_CosmicSignatureGameV3.TransactOpts, priceMaxLimit_, message_, bidCstRewardAmountMinLimit_, nftAddress_, nftId_)
}

// BidWithCstAndDonateToken is a paid mutator transaction binding the contract method 0x11b0d1fe.
//
// Solidity: function bidWithCstAndDonateToken(uint256 priceMaxLimit_, string message_, uint256 bidCstRewardAmountMinLimit_, address tokenAddress_, uint256 amount_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) BidWithCstAndDonateToken(opts *bind.TransactOpts, priceMaxLimit_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "bidWithCstAndDonateToken", priceMaxLimit_, message_, bidCstRewardAmountMinLimit_, tokenAddress_, amount_)
}

// BidWithCstAndDonateToken is a paid mutator transaction binding the contract method 0x11b0d1fe.
//
// Solidity: function bidWithCstAndDonateToken(uint256 priceMaxLimit_, string message_, uint256 bidCstRewardAmountMinLimit_, address tokenAddress_, uint256 amount_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) BidWithCstAndDonateToken(priceMaxLimit_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.BidWithCstAndDonateToken(&_CosmicSignatureGameV3.TransactOpts, priceMaxLimit_, message_, bidCstRewardAmountMinLimit_, tokenAddress_, amount_)
}

// BidWithCstAndDonateToken is a paid mutator transaction binding the contract method 0x11b0d1fe.
//
// Solidity: function bidWithCstAndDonateToken(uint256 priceMaxLimit_, string message_, uint256 bidCstRewardAmountMinLimit_, address tokenAddress_, uint256 amount_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) BidWithCstAndDonateToken(priceMaxLimit_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.BidWithCstAndDonateToken(&_CosmicSignatureGameV3.TransactOpts, priceMaxLimit_, message_, bidCstRewardAmountMinLimit_, tokenAddress_, amount_)
}

// BidWithEth is a paid mutator transaction binding the contract method 0x928880fa.
//
// Solidity: function bidWithEth(int256 randomWalkNftId_, string message_, uint256 bidCstRewardAmountMinLimit_) payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) BidWithEth(opts *bind.TransactOpts, randomWalkNftId_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "bidWithEth", randomWalkNftId_, message_, bidCstRewardAmountMinLimit_)
}

// BidWithEth is a paid mutator transaction binding the contract method 0x928880fa.
//
// Solidity: function bidWithEth(int256 randomWalkNftId_, string message_, uint256 bidCstRewardAmountMinLimit_) payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) BidWithEth(randomWalkNftId_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.BidWithEth(&_CosmicSignatureGameV3.TransactOpts, randomWalkNftId_, message_, bidCstRewardAmountMinLimit_)
}

// BidWithEth is a paid mutator transaction binding the contract method 0x928880fa.
//
// Solidity: function bidWithEth(int256 randomWalkNftId_, string message_, uint256 bidCstRewardAmountMinLimit_) payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) BidWithEth(randomWalkNftId_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.BidWithEth(&_CosmicSignatureGameV3.TransactOpts, randomWalkNftId_, message_, bidCstRewardAmountMinLimit_)
}

// BidWithEthAndDonateNft is a paid mutator transaction binding the contract method 0xb78d1e2a.
//
// Solidity: function bidWithEthAndDonateNft(int256 randomWalkNftId_, string message_, uint256 bidCstRewardAmountMinLimit_, address nftAddress_, uint256 nftId_) payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) BidWithEthAndDonateNft(opts *bind.TransactOpts, randomWalkNftId_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "bidWithEthAndDonateNft", randomWalkNftId_, message_, bidCstRewardAmountMinLimit_, nftAddress_, nftId_)
}

// BidWithEthAndDonateNft is a paid mutator transaction binding the contract method 0xb78d1e2a.
//
// Solidity: function bidWithEthAndDonateNft(int256 randomWalkNftId_, string message_, uint256 bidCstRewardAmountMinLimit_, address nftAddress_, uint256 nftId_) payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) BidWithEthAndDonateNft(randomWalkNftId_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.BidWithEthAndDonateNft(&_CosmicSignatureGameV3.TransactOpts, randomWalkNftId_, message_, bidCstRewardAmountMinLimit_, nftAddress_, nftId_)
}

// BidWithEthAndDonateNft is a paid mutator transaction binding the contract method 0xb78d1e2a.
//
// Solidity: function bidWithEthAndDonateNft(int256 randomWalkNftId_, string message_, uint256 bidCstRewardAmountMinLimit_, address nftAddress_, uint256 nftId_) payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) BidWithEthAndDonateNft(randomWalkNftId_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, nftAddress_ common.Address, nftId_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.BidWithEthAndDonateNft(&_CosmicSignatureGameV3.TransactOpts, randomWalkNftId_, message_, bidCstRewardAmountMinLimit_, nftAddress_, nftId_)
}

// BidWithEthAndDonateToken is a paid mutator transaction binding the contract method 0x876d5c36.
//
// Solidity: function bidWithEthAndDonateToken(int256 randomWalkNftId_, string message_, uint256 bidCstRewardAmountMinLimit_, address tokenAddress_, uint256 amount_) payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) BidWithEthAndDonateToken(opts *bind.TransactOpts, randomWalkNftId_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "bidWithEthAndDonateToken", randomWalkNftId_, message_, bidCstRewardAmountMinLimit_, tokenAddress_, amount_)
}

// BidWithEthAndDonateToken is a paid mutator transaction binding the contract method 0x876d5c36.
//
// Solidity: function bidWithEthAndDonateToken(int256 randomWalkNftId_, string message_, uint256 bidCstRewardAmountMinLimit_, address tokenAddress_, uint256 amount_) payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) BidWithEthAndDonateToken(randomWalkNftId_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.BidWithEthAndDonateToken(&_CosmicSignatureGameV3.TransactOpts, randomWalkNftId_, message_, bidCstRewardAmountMinLimit_, tokenAddress_, amount_)
}

// BidWithEthAndDonateToken is a paid mutator transaction binding the contract method 0x876d5c36.
//
// Solidity: function bidWithEthAndDonateToken(int256 randomWalkNftId_, string message_, uint256 bidCstRewardAmountMinLimit_, address tokenAddress_, uint256 amount_) payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) BidWithEthAndDonateToken(randomWalkNftId_ *big.Int, message_ string, bidCstRewardAmountMinLimit_ *big.Int, tokenAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.BidWithEthAndDonateToken(&_CosmicSignatureGameV3.TransactOpts, randomWalkNftId_, message_, bidCstRewardAmountMinLimit_, tokenAddress_, amount_)
}

// ClaimMainPrize is a paid mutator transaction binding the contract method 0x448c6eb1.
//
// Solidity: function claimMainPrize() returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) ClaimMainPrize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "claimMainPrize")
}

// ClaimMainPrize is a paid mutator transaction binding the contract method 0x448c6eb1.
//
// Solidity: function claimMainPrize() returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) ClaimMainPrize() (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.ClaimMainPrize(&_CosmicSignatureGameV3.TransactOpts)
}

// ClaimMainPrize is a paid mutator transaction binding the contract method 0x448c6eb1.
//
// Solidity: function claimMainPrize() returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) ClaimMainPrize() (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.ClaimMainPrize(&_CosmicSignatureGameV3.TransactOpts)
}

// DonateEth is a paid mutator transaction binding the contract method 0xaadd1b03.
//
// Solidity: function donateEth() payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) DonateEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "donateEth")
}

// DonateEth is a paid mutator transaction binding the contract method 0xaadd1b03.
//
// Solidity: function donateEth() payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) DonateEth() (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.DonateEth(&_CosmicSignatureGameV3.TransactOpts)
}

// DonateEth is a paid mutator transaction binding the contract method 0xaadd1b03.
//
// Solidity: function donateEth() payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) DonateEth() (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.DonateEth(&_CosmicSignatureGameV3.TransactOpts)
}

// DonateEthWithInfo is a paid mutator transaction binding the contract method 0x23b31cfc.
//
// Solidity: function donateEthWithInfo(string data_) payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) DonateEthWithInfo(opts *bind.TransactOpts, data_ string) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "donateEthWithInfo", data_)
}

// DonateEthWithInfo is a paid mutator transaction binding the contract method 0x23b31cfc.
//
// Solidity: function donateEthWithInfo(string data_) payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) DonateEthWithInfo(data_ string) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.DonateEthWithInfo(&_CosmicSignatureGameV3.TransactOpts, data_)
}

// DonateEthWithInfo is a paid mutator transaction binding the contract method 0x23b31cfc.
//
// Solidity: function donateEthWithInfo(string data_) payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) DonateEthWithInfo(data_ string) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.DonateEthWithInfo(&_CosmicSignatureGameV3.TransactOpts, data_)
}

// HalveEthDutchAuctionEndingBidPrice is a paid mutator transaction binding the contract method 0xdfcd00d1.
//
// Solidity: function halveEthDutchAuctionEndingBidPrice() returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) HalveEthDutchAuctionEndingBidPrice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "halveEthDutchAuctionEndingBidPrice")
}

// HalveEthDutchAuctionEndingBidPrice is a paid mutator transaction binding the contract method 0xdfcd00d1.
//
// Solidity: function halveEthDutchAuctionEndingBidPrice() returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) HalveEthDutchAuctionEndingBidPrice() (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.HalveEthDutchAuctionEndingBidPrice(&_CosmicSignatureGameV3.TransactOpts)
}

// HalveEthDutchAuctionEndingBidPrice is a paid mutator transaction binding the contract method 0xdfcd00d1.
//
// Solidity: function halveEthDutchAuctionEndingBidPrice() returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) HalveEthDutchAuctionEndingBidPrice() (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.HalveEthDutchAuctionEndingBidPrice(&_CosmicSignatureGameV3.TransactOpts)
}

// Reinitialize is a paid mutator transaction binding the contract method 0x6c2eb350.
//
// Solidity: function reinitialize() returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) Reinitialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "reinitialize")
}

// Reinitialize is a paid mutator transaction binding the contract method 0x6c2eb350.
//
// Solidity: function reinitialize() returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) Reinitialize() (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.Reinitialize(&_CosmicSignatureGameV3.TransactOpts)
}

// Reinitialize is a paid mutator transaction binding the contract method 0x6c2eb350.
//
// Solidity: function reinitialize() returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) Reinitialize() (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.Reinitialize(&_CosmicSignatureGameV3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) RenounceOwnership() (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.RenounceOwnership(&_CosmicSignatureGameV3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.RenounceOwnership(&_CosmicSignatureGameV3.TransactOpts)
}

// SetBidCstRewardAmountMultiplier is a paid mutator transaction binding the contract method 0x09632366.
//
// Solidity: function setBidCstRewardAmountMultiplier(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetBidCstRewardAmountMultiplier(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setBidCstRewardAmountMultiplier", newValue_)
}

// SetBidCstRewardAmountMultiplier is a paid mutator transaction binding the contract method 0x09632366.
//
// Solidity: function setBidCstRewardAmountMultiplier(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetBidCstRewardAmountMultiplier(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetBidCstRewardAmountMultiplier(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetBidCstRewardAmountMultiplier is a paid mutator transaction binding the contract method 0x09632366.
//
// Solidity: function setBidCstRewardAmountMultiplier(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetBidCstRewardAmountMultiplier(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetBidCstRewardAmountMultiplier(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetBidCstRewardAmountPerMinute is a paid mutator transaction binding the contract method 0x25ed499c.
//
// Solidity: function setBidCstRewardAmountPerMinute(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetBidCstRewardAmountPerMinute(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setBidCstRewardAmountPerMinute", newValue_)
}

// SetBidCstRewardAmountPerMinute is a paid mutator transaction binding the contract method 0x25ed499c.
//
// Solidity: function setBidCstRewardAmountPerMinute(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetBidCstRewardAmountPerMinute(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetBidCstRewardAmountPerMinute(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetBidCstRewardAmountPerMinute is a paid mutator transaction binding the contract method 0x25ed499c.
//
// Solidity: function setBidCstRewardAmountPerMinute(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetBidCstRewardAmountPerMinute(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetBidCstRewardAmountPerMinute(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetBidMessageLengthMaxLimit is a paid mutator transaction binding the contract method 0x543f416f.
//
// Solidity: function setBidMessageLengthMaxLimit(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetBidMessageLengthMaxLimit(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setBidMessageLengthMaxLimit", newValue_)
}

// SetBidMessageLengthMaxLimit is a paid mutator transaction binding the contract method 0x543f416f.
//
// Solidity: function setBidMessageLengthMaxLimit(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetBidMessageLengthMaxLimit(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetBidMessageLengthMaxLimit(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetBidMessageLengthMaxLimit is a paid mutator transaction binding the contract method 0x543f416f.
//
// Solidity: function setBidMessageLengthMaxLimit(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetBidMessageLengthMaxLimit(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetBidMessageLengthMaxLimit(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetCharityAddress is a paid mutator transaction binding the contract method 0x0c9be46d.
//
// Solidity: function setCharityAddress(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetCharityAddress(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setCharityAddress", newValue_)
}

// SetCharityAddress is a paid mutator transaction binding the contract method 0x0c9be46d.
//
// Solidity: function setCharityAddress(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetCharityAddress(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetCharityAddress(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetCharityAddress is a paid mutator transaction binding the contract method 0x0c9be46d.
//
// Solidity: function setCharityAddress(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetCharityAddress(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetCharityAddress(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetCharityEthDonationAmountPercentage is a paid mutator transaction binding the contract method 0x2d829a2d.
//
// Solidity: function setCharityEthDonationAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetCharityEthDonationAmountPercentage(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setCharityEthDonationAmountPercentage", newValue_)
}

// SetCharityEthDonationAmountPercentage is a paid mutator transaction binding the contract method 0x2d829a2d.
//
// Solidity: function setCharityEthDonationAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetCharityEthDonationAmountPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetCharityEthDonationAmountPercentage(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetCharityEthDonationAmountPercentage is a paid mutator transaction binding the contract method 0x2d829a2d.
//
// Solidity: function setCharityEthDonationAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetCharityEthDonationAmountPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetCharityEthDonationAmountPercentage(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetChronoWarriorEthPrizeAmountPercentage is a paid mutator transaction binding the contract method 0x1f1b4aa4.
//
// Solidity: function setChronoWarriorEthPrizeAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetChronoWarriorEthPrizeAmountPercentage(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setChronoWarriorEthPrizeAmountPercentage", newValue_)
}

// SetChronoWarriorEthPrizeAmountPercentage is a paid mutator transaction binding the contract method 0x1f1b4aa4.
//
// Solidity: function setChronoWarriorEthPrizeAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetChronoWarriorEthPrizeAmountPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetChronoWarriorEthPrizeAmountPercentage(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetChronoWarriorEthPrizeAmountPercentage is a paid mutator transaction binding the contract method 0x1f1b4aa4.
//
// Solidity: function setChronoWarriorEthPrizeAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetChronoWarriorEthPrizeAmountPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetChronoWarriorEthPrizeAmountPercentage(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetCosmicSignatureNft is a paid mutator transaction binding the contract method 0x44acc12a.
//
// Solidity: function setCosmicSignatureNft(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetCosmicSignatureNft(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setCosmicSignatureNft", newValue_)
}

// SetCosmicSignatureNft is a paid mutator transaction binding the contract method 0x44acc12a.
//
// Solidity: function setCosmicSignatureNft(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetCosmicSignatureNft(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetCosmicSignatureNft(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetCosmicSignatureNft is a paid mutator transaction binding the contract method 0x44acc12a.
//
// Solidity: function setCosmicSignatureNft(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetCosmicSignatureNft(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetCosmicSignatureNft(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetCosmicSignatureNftStakingTotalEthRewardAmountPercentage is a paid mutator transaction binding the contract method 0x4feb78b7.
//
// Solidity: function setCosmicSignatureNftStakingTotalEthRewardAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetCosmicSignatureNftStakingTotalEthRewardAmountPercentage(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setCosmicSignatureNftStakingTotalEthRewardAmountPercentage", newValue_)
}

// SetCosmicSignatureNftStakingTotalEthRewardAmountPercentage is a paid mutator transaction binding the contract method 0x4feb78b7.
//
// Solidity: function setCosmicSignatureNftStakingTotalEthRewardAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetCosmicSignatureNftStakingTotalEthRewardAmountPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetCosmicSignatureNftStakingTotalEthRewardAmountPercentage(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetCosmicSignatureNftStakingTotalEthRewardAmountPercentage is a paid mutator transaction binding the contract method 0x4feb78b7.
//
// Solidity: function setCosmicSignatureNftStakingTotalEthRewardAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetCosmicSignatureNftStakingTotalEthRewardAmountPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetCosmicSignatureNftStakingTotalEthRewardAmountPercentage(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetCosmicSignatureToken is a paid mutator transaction binding the contract method 0x9646d758.
//
// Solidity: function setCosmicSignatureToken(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetCosmicSignatureToken(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setCosmicSignatureToken", newValue_)
}

// SetCosmicSignatureToken is a paid mutator transaction binding the contract method 0x9646d758.
//
// Solidity: function setCosmicSignatureToken(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetCosmicSignatureToken(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetCosmicSignatureToken(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetCosmicSignatureToken is a paid mutator transaction binding the contract method 0x9646d758.
//
// Solidity: function setCosmicSignatureToken(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetCosmicSignatureToken(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetCosmicSignatureToken(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetCstDutchAuctionBeginningBidPriceMinLimit is a paid mutator transaction binding the contract method 0x88ce802c.
//
// Solidity: function setCstDutchAuctionBeginningBidPriceMinLimit(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetCstDutchAuctionBeginningBidPriceMinLimit(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setCstDutchAuctionBeginningBidPriceMinLimit", newValue_)
}

// SetCstDutchAuctionBeginningBidPriceMinLimit is a paid mutator transaction binding the contract method 0x88ce802c.
//
// Solidity: function setCstDutchAuctionBeginningBidPriceMinLimit(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetCstDutchAuctionBeginningBidPriceMinLimit(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetCstDutchAuctionBeginningBidPriceMinLimit(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetCstDutchAuctionBeginningBidPriceMinLimit is a paid mutator transaction binding the contract method 0x88ce802c.
//
// Solidity: function setCstDutchAuctionBeginningBidPriceMinLimit(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetCstDutchAuctionBeginningBidPriceMinLimit(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetCstDutchAuctionBeginningBidPriceMinLimit(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetCstDutchAuctionDuration is a paid mutator transaction binding the contract method 0xde704b41.
//
// Solidity: function setCstDutchAuctionDuration(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetCstDutchAuctionDuration(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setCstDutchAuctionDuration", newValue_)
}

// SetCstDutchAuctionDuration is a paid mutator transaction binding the contract method 0xde704b41.
//
// Solidity: function setCstDutchAuctionDuration(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetCstDutchAuctionDuration(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetCstDutchAuctionDuration(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetCstDutchAuctionDuration is a paid mutator transaction binding the contract method 0xde704b41.
//
// Solidity: function setCstDutchAuctionDuration(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetCstDutchAuctionDuration(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetCstDutchAuctionDuration(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetCstDutchAuctionDurationChangeDivisor is a paid mutator transaction binding the contract method 0x04338479.
//
// Solidity: function setCstDutchAuctionDurationChangeDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetCstDutchAuctionDurationChangeDivisor(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setCstDutchAuctionDurationChangeDivisor", newValue_)
}

// SetCstDutchAuctionDurationChangeDivisor is a paid mutator transaction binding the contract method 0x04338479.
//
// Solidity: function setCstDutchAuctionDurationChangeDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetCstDutchAuctionDurationChangeDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetCstDutchAuctionDurationChangeDivisor(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetCstDutchAuctionDurationChangeDivisor is a paid mutator transaction binding the contract method 0x04338479.
//
// Solidity: function setCstDutchAuctionDurationChangeDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetCstDutchAuctionDurationChangeDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetCstDutchAuctionDurationChangeDivisor(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetCstPrizeAmount is a paid mutator transaction binding the contract method 0xecb5776e.
//
// Solidity: function setCstPrizeAmount(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetCstPrizeAmount(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setCstPrizeAmount", newValue_)
}

// SetCstPrizeAmount is a paid mutator transaction binding the contract method 0xecb5776e.
//
// Solidity: function setCstPrizeAmount(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetCstPrizeAmount(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetCstPrizeAmount(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetCstPrizeAmount is a paid mutator transaction binding the contract method 0xecb5776e.
//
// Solidity: function setCstPrizeAmount(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetCstPrizeAmount(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetCstPrizeAmount(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetDelayDurationBeforeRoundActivation is a paid mutator transaction binding the contract method 0x09794bee.
//
// Solidity: function setDelayDurationBeforeRoundActivation(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetDelayDurationBeforeRoundActivation(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setDelayDurationBeforeRoundActivation", newValue_)
}

// SetDelayDurationBeforeRoundActivation is a paid mutator transaction binding the contract method 0x09794bee.
//
// Solidity: function setDelayDurationBeforeRoundActivation(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetDelayDurationBeforeRoundActivation(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetDelayDurationBeforeRoundActivation(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetDelayDurationBeforeRoundActivation is a paid mutator transaction binding the contract method 0x09794bee.
//
// Solidity: function setDelayDurationBeforeRoundActivation(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetDelayDurationBeforeRoundActivation(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetDelayDurationBeforeRoundActivation(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetEthBidPriceIncreaseDivisor is a paid mutator transaction binding the contract method 0xf49efe9d.
//
// Solidity: function setEthBidPriceIncreaseDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetEthBidPriceIncreaseDivisor(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setEthBidPriceIncreaseDivisor", newValue_)
}

// SetEthBidPriceIncreaseDivisor is a paid mutator transaction binding the contract method 0xf49efe9d.
//
// Solidity: function setEthBidPriceIncreaseDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetEthBidPriceIncreaseDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetEthBidPriceIncreaseDivisor(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetEthBidPriceIncreaseDivisor is a paid mutator transaction binding the contract method 0xf49efe9d.
//
// Solidity: function setEthBidPriceIncreaseDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetEthBidPriceIncreaseDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetEthBidPriceIncreaseDivisor(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetEthBidRefundAmountInGasToSwallowMaxLimit is a paid mutator transaction binding the contract method 0x441b3289.
//
// Solidity: function setEthBidRefundAmountInGasToSwallowMaxLimit(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetEthBidRefundAmountInGasToSwallowMaxLimit(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setEthBidRefundAmountInGasToSwallowMaxLimit", newValue_)
}

// SetEthBidRefundAmountInGasToSwallowMaxLimit is a paid mutator transaction binding the contract method 0x441b3289.
//
// Solidity: function setEthBidRefundAmountInGasToSwallowMaxLimit(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetEthBidRefundAmountInGasToSwallowMaxLimit(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetEthBidRefundAmountInGasToSwallowMaxLimit(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetEthBidRefundAmountInGasToSwallowMaxLimit is a paid mutator transaction binding the contract method 0x441b3289.
//
// Solidity: function setEthBidRefundAmountInGasToSwallowMaxLimit(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetEthBidRefundAmountInGasToSwallowMaxLimit(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetEthBidRefundAmountInGasToSwallowMaxLimit(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetEthDutchAuctionDurationDivisor is a paid mutator transaction binding the contract method 0xf34d411c.
//
// Solidity: function setEthDutchAuctionDurationDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetEthDutchAuctionDurationDivisor(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setEthDutchAuctionDurationDivisor", newValue_)
}

// SetEthDutchAuctionDurationDivisor is a paid mutator transaction binding the contract method 0xf34d411c.
//
// Solidity: function setEthDutchAuctionDurationDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetEthDutchAuctionDurationDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetEthDutchAuctionDurationDivisor(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetEthDutchAuctionDurationDivisor is a paid mutator transaction binding the contract method 0xf34d411c.
//
// Solidity: function setEthDutchAuctionDurationDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetEthDutchAuctionDurationDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetEthDutchAuctionDurationDivisor(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetEthDutchAuctionEndingBidPriceDivisor is a paid mutator transaction binding the contract method 0xddd6df07.
//
// Solidity: function setEthDutchAuctionEndingBidPriceDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetEthDutchAuctionEndingBidPriceDivisor(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setEthDutchAuctionEndingBidPriceDivisor", newValue_)
}

// SetEthDutchAuctionEndingBidPriceDivisor is a paid mutator transaction binding the contract method 0xddd6df07.
//
// Solidity: function setEthDutchAuctionEndingBidPriceDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetEthDutchAuctionEndingBidPriceDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetEthDutchAuctionEndingBidPriceDivisor(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetEthDutchAuctionEndingBidPriceDivisor is a paid mutator transaction binding the contract method 0xddd6df07.
//
// Solidity: function setEthDutchAuctionEndingBidPriceDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetEthDutchAuctionEndingBidPriceDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetEthDutchAuctionEndingBidPriceDivisor(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetInitialDurationUntilMainPrizeDivisor is a paid mutator transaction binding the contract method 0x77fa1027.
//
// Solidity: function setInitialDurationUntilMainPrizeDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetInitialDurationUntilMainPrizeDivisor(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setInitialDurationUntilMainPrizeDivisor", newValue_)
}

// SetInitialDurationUntilMainPrizeDivisor is a paid mutator transaction binding the contract method 0x77fa1027.
//
// Solidity: function setInitialDurationUntilMainPrizeDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetInitialDurationUntilMainPrizeDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetInitialDurationUntilMainPrizeDivisor(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetInitialDurationUntilMainPrizeDivisor is a paid mutator transaction binding the contract method 0x77fa1027.
//
// Solidity: function setInitialDurationUntilMainPrizeDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetInitialDurationUntilMainPrizeDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetInitialDurationUntilMainPrizeDivisor(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetMainEthPrizeAmountPercentage is a paid mutator transaction binding the contract method 0x6b59acb8.
//
// Solidity: function setMainEthPrizeAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetMainEthPrizeAmountPercentage(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setMainEthPrizeAmountPercentage", newValue_)
}

// SetMainEthPrizeAmountPercentage is a paid mutator transaction binding the contract method 0x6b59acb8.
//
// Solidity: function setMainEthPrizeAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetMainEthPrizeAmountPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetMainEthPrizeAmountPercentage(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetMainEthPrizeAmountPercentage is a paid mutator transaction binding the contract method 0x6b59acb8.
//
// Solidity: function setMainEthPrizeAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetMainEthPrizeAmountPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetMainEthPrizeAmountPercentage(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetMainPrizeNumCosmicSignatureNfts is a paid mutator transaction binding the contract method 0x87292a85.
//
// Solidity: function setMainPrizeNumCosmicSignatureNfts(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetMainPrizeNumCosmicSignatureNfts(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setMainPrizeNumCosmicSignatureNfts", newValue_)
}

// SetMainPrizeNumCosmicSignatureNfts is a paid mutator transaction binding the contract method 0x87292a85.
//
// Solidity: function setMainPrizeNumCosmicSignatureNfts(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetMainPrizeNumCosmicSignatureNfts(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetMainPrizeNumCosmicSignatureNfts(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetMainPrizeNumCosmicSignatureNfts is a paid mutator transaction binding the contract method 0x87292a85.
//
// Solidity: function setMainPrizeNumCosmicSignatureNfts(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetMainPrizeNumCosmicSignatureNfts(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetMainPrizeNumCosmicSignatureNfts(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetMainPrizeTimeIncrementInMicroSeconds is a paid mutator transaction binding the contract method 0xa4be0d40.
//
// Solidity: function setMainPrizeTimeIncrementInMicroSeconds(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetMainPrizeTimeIncrementInMicroSeconds(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setMainPrizeTimeIncrementInMicroSeconds", newValue_)
}

// SetMainPrizeTimeIncrementInMicroSeconds is a paid mutator transaction binding the contract method 0xa4be0d40.
//
// Solidity: function setMainPrizeTimeIncrementInMicroSeconds(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetMainPrizeTimeIncrementInMicroSeconds(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetMainPrizeTimeIncrementInMicroSeconds(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetMainPrizeTimeIncrementInMicroSeconds is a paid mutator transaction binding the contract method 0xa4be0d40.
//
// Solidity: function setMainPrizeTimeIncrementInMicroSeconds(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetMainPrizeTimeIncrementInMicroSeconds(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetMainPrizeTimeIncrementInMicroSeconds(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetMainPrizeTimeIncrementIncreaseDivisor is a paid mutator transaction binding the contract method 0xcfb4e599.
//
// Solidity: function setMainPrizeTimeIncrementIncreaseDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetMainPrizeTimeIncrementIncreaseDivisor(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setMainPrizeTimeIncrementIncreaseDivisor", newValue_)
}

// SetMainPrizeTimeIncrementIncreaseDivisor is a paid mutator transaction binding the contract method 0xcfb4e599.
//
// Solidity: function setMainPrizeTimeIncrementIncreaseDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetMainPrizeTimeIncrementIncreaseDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetMainPrizeTimeIncrementIncreaseDivisor(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetMainPrizeTimeIncrementIncreaseDivisor is a paid mutator transaction binding the contract method 0xcfb4e599.
//
// Solidity: function setMainPrizeTimeIncrementIncreaseDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetMainPrizeTimeIncrementIncreaseDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetMainPrizeTimeIncrementIncreaseDivisor(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetMarketingWallet is a paid mutator transaction binding the contract method 0x5d098b38.
//
// Solidity: function setMarketingWallet(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetMarketingWallet(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setMarketingWallet", newValue_)
}

// SetMarketingWallet is a paid mutator transaction binding the contract method 0x5d098b38.
//
// Solidity: function setMarketingWallet(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetMarketingWallet(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetMarketingWallet(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetMarketingWallet is a paid mutator transaction binding the contract method 0x5d098b38.
//
// Solidity: function setMarketingWallet(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetMarketingWallet(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetMarketingWallet(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetMarketingWalletCstContributionAmount is a paid mutator transaction binding the contract method 0xd9ab9eaa.
//
// Solidity: function setMarketingWalletCstContributionAmount(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetMarketingWalletCstContributionAmount(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setMarketingWalletCstContributionAmount", newValue_)
}

// SetMarketingWalletCstContributionAmount is a paid mutator transaction binding the contract method 0xd9ab9eaa.
//
// Solidity: function setMarketingWalletCstContributionAmount(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetMarketingWalletCstContributionAmount(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetMarketingWalletCstContributionAmount(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetMarketingWalletCstContributionAmount is a paid mutator transaction binding the contract method 0xd9ab9eaa.
//
// Solidity: function setMarketingWalletCstContributionAmount(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetMarketingWalletCstContributionAmount(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetMarketingWalletCstContributionAmount(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetNumRaffleCosmicSignatureNftsForBidders is a paid mutator transaction binding the contract method 0x1e9cbb7e.
//
// Solidity: function setNumRaffleCosmicSignatureNftsForBidders(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetNumRaffleCosmicSignatureNftsForBidders(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setNumRaffleCosmicSignatureNftsForBidders", newValue_)
}

// SetNumRaffleCosmicSignatureNftsForBidders is a paid mutator transaction binding the contract method 0x1e9cbb7e.
//
// Solidity: function setNumRaffleCosmicSignatureNftsForBidders(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetNumRaffleCosmicSignatureNftsForBidders(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetNumRaffleCosmicSignatureNftsForBidders(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetNumRaffleCosmicSignatureNftsForBidders is a paid mutator transaction binding the contract method 0x1e9cbb7e.
//
// Solidity: function setNumRaffleCosmicSignatureNftsForBidders(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetNumRaffleCosmicSignatureNftsForBidders(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetNumRaffleCosmicSignatureNftsForBidders(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetNumRaffleCosmicSignatureNftsForRandomWalkNftStakers is a paid mutator transaction binding the contract method 0x135f3d28.
//
// Solidity: function setNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setNumRaffleCosmicSignatureNftsForRandomWalkNftStakers", newValue_)
}

// SetNumRaffleCosmicSignatureNftsForRandomWalkNftStakers is a paid mutator transaction binding the contract method 0x135f3d28.
//
// Solidity: function setNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetNumRaffleCosmicSignatureNftsForRandomWalkNftStakers is a paid mutator transaction binding the contract method 0x135f3d28.
//
// Solidity: function setNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetNumRaffleEthPrizesForBidders is a paid mutator transaction binding the contract method 0x2f894cd7.
//
// Solidity: function setNumRaffleEthPrizesForBidders(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetNumRaffleEthPrizesForBidders(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setNumRaffleEthPrizesForBidders", newValue_)
}

// SetNumRaffleEthPrizesForBidders is a paid mutator transaction binding the contract method 0x2f894cd7.
//
// Solidity: function setNumRaffleEthPrizesForBidders(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetNumRaffleEthPrizesForBidders(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetNumRaffleEthPrizesForBidders(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetNumRaffleEthPrizesForBidders is a paid mutator transaction binding the contract method 0x2f894cd7.
//
// Solidity: function setNumRaffleEthPrizesForBidders(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetNumRaffleEthPrizesForBidders(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetNumRaffleEthPrizesForBidders(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetPrizesWallet is a paid mutator transaction binding the contract method 0x6c17e3cc.
//
// Solidity: function setPrizesWallet(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetPrizesWallet(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setPrizesWallet", newValue_)
}

// SetPrizesWallet is a paid mutator transaction binding the contract method 0x6c17e3cc.
//
// Solidity: function setPrizesWallet(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetPrizesWallet(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetPrizesWallet(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetPrizesWallet is a paid mutator transaction binding the contract method 0x6c17e3cc.
//
// Solidity: function setPrizesWallet(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetPrizesWallet(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetPrizesWallet(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetRaffleTotalEthPrizeAmountForBiddersPercentage is a paid mutator transaction binding the contract method 0xfdfb9ba4.
//
// Solidity: function setRaffleTotalEthPrizeAmountForBiddersPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetRaffleTotalEthPrizeAmountForBiddersPercentage(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setRaffleTotalEthPrizeAmountForBiddersPercentage", newValue_)
}

// SetRaffleTotalEthPrizeAmountForBiddersPercentage is a paid mutator transaction binding the contract method 0xfdfb9ba4.
//
// Solidity: function setRaffleTotalEthPrizeAmountForBiddersPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetRaffleTotalEthPrizeAmountForBiddersPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetRaffleTotalEthPrizeAmountForBiddersPercentage(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetRaffleTotalEthPrizeAmountForBiddersPercentage is a paid mutator transaction binding the contract method 0xfdfb9ba4.
//
// Solidity: function setRaffleTotalEthPrizeAmountForBiddersPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetRaffleTotalEthPrizeAmountForBiddersPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetRaffleTotalEthPrizeAmountForBiddersPercentage(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetRandomWalkNft is a paid mutator transaction binding the contract method 0x9edeaf8e.
//
// Solidity: function setRandomWalkNft(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetRandomWalkNft(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setRandomWalkNft", newValue_)
}

// SetRandomWalkNft is a paid mutator transaction binding the contract method 0x9edeaf8e.
//
// Solidity: function setRandomWalkNft(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetRandomWalkNft(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetRandomWalkNft(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetRandomWalkNft is a paid mutator transaction binding the contract method 0x9edeaf8e.
//
// Solidity: function setRandomWalkNft(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetRandomWalkNft(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetRandomWalkNft(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetRoundActivationTime is a paid mutator transaction binding the contract method 0x2b91c7bb.
//
// Solidity: function setRoundActivationTime(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetRoundActivationTime(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setRoundActivationTime", newValue_)
}

// SetRoundActivationTime is a paid mutator transaction binding the contract method 0x2b91c7bb.
//
// Solidity: function setRoundActivationTime(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetRoundActivationTime(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetRoundActivationTime(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetRoundActivationTime is a paid mutator transaction binding the contract method 0x2b91c7bb.
//
// Solidity: function setRoundActivationTime(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetRoundActivationTime(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetRoundActivationTime(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetRoundLateBidDurationDivisor is a paid mutator transaction binding the contract method 0x8c94e9ba.
//
// Solidity: function setRoundLateBidDurationDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetRoundLateBidDurationDivisor(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setRoundLateBidDurationDivisor", newValue_)
}

// SetRoundLateBidDurationDivisor is a paid mutator transaction binding the contract method 0x8c94e9ba.
//
// Solidity: function setRoundLateBidDurationDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetRoundLateBidDurationDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetRoundLateBidDurationDivisor(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetRoundLateBidDurationDivisor is a paid mutator transaction binding the contract method 0x8c94e9ba.
//
// Solidity: function setRoundLateBidDurationDivisor(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetRoundLateBidDurationDivisor(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetRoundLateBidDurationDivisor(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetRoundLateBidPricePremiumAmountBaseMultiplier is a paid mutator transaction binding the contract method 0x75ef3b9c.
//
// Solidity: function setRoundLateBidPricePremiumAmountBaseMultiplier(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetRoundLateBidPricePremiumAmountBaseMultiplier(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setRoundLateBidPricePremiumAmountBaseMultiplier", newValue_)
}

// SetRoundLateBidPricePremiumAmountBaseMultiplier is a paid mutator transaction binding the contract method 0x75ef3b9c.
//
// Solidity: function setRoundLateBidPricePremiumAmountBaseMultiplier(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetRoundLateBidPricePremiumAmountBaseMultiplier(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetRoundLateBidPricePremiumAmountBaseMultiplier(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetRoundLateBidPricePremiumAmountBaseMultiplier is a paid mutator transaction binding the contract method 0x75ef3b9c.
//
// Solidity: function setRoundLateBidPricePremiumAmountBaseMultiplier(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetRoundLateBidPricePremiumAmountBaseMultiplier(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetRoundLateBidPricePremiumAmountBaseMultiplier(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetRoundLateBidPricePremiumAmountExponent is a paid mutator transaction binding the contract method 0x1aaba5a5.
//
// Solidity: function setRoundLateBidPricePremiumAmountExponent(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetRoundLateBidPricePremiumAmountExponent(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setRoundLateBidPricePremiumAmountExponent", newValue_)
}

// SetRoundLateBidPricePremiumAmountExponent is a paid mutator transaction binding the contract method 0x1aaba5a5.
//
// Solidity: function setRoundLateBidPricePremiumAmountExponent(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetRoundLateBidPricePremiumAmountExponent(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetRoundLateBidPricePremiumAmountExponent(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetRoundLateBidPricePremiumAmountExponent is a paid mutator transaction binding the contract method 0x1aaba5a5.
//
// Solidity: function setRoundLateBidPricePremiumAmountExponent(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetRoundLateBidPricePremiumAmountExponent(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetRoundLateBidPricePremiumAmountExponent(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetStakingWalletCosmicSignatureNft is a paid mutator transaction binding the contract method 0x5a1e5bde.
//
// Solidity: function setStakingWalletCosmicSignatureNft(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetStakingWalletCosmicSignatureNft(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setStakingWalletCosmicSignatureNft", newValue_)
}

// SetStakingWalletCosmicSignatureNft is a paid mutator transaction binding the contract method 0x5a1e5bde.
//
// Solidity: function setStakingWalletCosmicSignatureNft(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetStakingWalletCosmicSignatureNft(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetStakingWalletCosmicSignatureNft(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetStakingWalletCosmicSignatureNft is a paid mutator transaction binding the contract method 0x5a1e5bde.
//
// Solidity: function setStakingWalletCosmicSignatureNft(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetStakingWalletCosmicSignatureNft(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetStakingWalletCosmicSignatureNft(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetStakingWalletRandomWalkNft is a paid mutator transaction binding the contract method 0xb4f1b134.
//
// Solidity: function setStakingWalletRandomWalkNft(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetStakingWalletRandomWalkNft(opts *bind.TransactOpts, newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setStakingWalletRandomWalkNft", newValue_)
}

// SetStakingWalletRandomWalkNft is a paid mutator transaction binding the contract method 0xb4f1b134.
//
// Solidity: function setStakingWalletRandomWalkNft(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetStakingWalletRandomWalkNft(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetStakingWalletRandomWalkNft(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetStakingWalletRandomWalkNft is a paid mutator transaction binding the contract method 0xb4f1b134.
//
// Solidity: function setStakingWalletRandomWalkNft(address newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetStakingWalletRandomWalkNft(newValue_ common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetStakingWalletRandomWalkNft(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetTimeoutDurationToClaimMainPrize is a paid mutator transaction binding the contract method 0xf0bdab7c.
//
// Solidity: function setTimeoutDurationToClaimMainPrize(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetTimeoutDurationToClaimMainPrize(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setTimeoutDurationToClaimMainPrize", newValue_)
}

// SetTimeoutDurationToClaimMainPrize is a paid mutator transaction binding the contract method 0xf0bdab7c.
//
// Solidity: function setTimeoutDurationToClaimMainPrize(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetTimeoutDurationToClaimMainPrize(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetTimeoutDurationToClaimMainPrize(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetTimeoutDurationToClaimMainPrize is a paid mutator transaction binding the contract method 0xf0bdab7c.
//
// Solidity: function setTimeoutDurationToClaimMainPrize(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetTimeoutDurationToClaimMainPrize(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetTimeoutDurationToClaimMainPrize(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.TransferOwnership(&_CosmicSignatureGameV3.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.TransferOwnership(&_CosmicSignatureGameV3.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.UpgradeToAndCall(&_CosmicSignatureGameV3.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.UpgradeToAndCall(&_CosmicSignatureGameV3.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) Receive() (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.Receive(&_CosmicSignatureGameV3.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) Receive() (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.Receive(&_CosmicSignatureGameV3.TransactOpts)
}

// CosmicSignatureGameV3ArbitrumErrorIterator is returned from FilterArbitrumError and is used to iterate over the raw logs and unpacked data for ArbitrumError events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3ArbitrumErrorIterator struct {
	Event *CosmicSignatureGameV3ArbitrumError // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3ArbitrumErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3ArbitrumError)
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
		it.Event = new(CosmicSignatureGameV3ArbitrumError)
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
func (it *CosmicSignatureGameV3ArbitrumErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3ArbitrumErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3ArbitrumError represents a ArbitrumError event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3ArbitrumError struct {
	ErrStr string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterArbitrumError is a free log retrieval operation binding the contract event 0xa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d.
//
// Solidity: event ArbitrumError(string errStr)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterArbitrumError(opts *bind.FilterOpts) (*CosmicSignatureGameV3ArbitrumErrorIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "ArbitrumError")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3ArbitrumErrorIterator{contract: _CosmicSignatureGameV3.contract, event: "ArbitrumError", logs: logs, sub: sub}, nil
}

// WatchArbitrumError is a free log subscription operation binding the contract event 0xa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d.
//
// Solidity: event ArbitrumError(string errStr)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchArbitrumError(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3ArbitrumError) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "ArbitrumError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3ArbitrumError)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "ArbitrumError", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseArbitrumError(log types.Log) (*CosmicSignatureGameV3ArbitrumError, error) {
	event := new(CosmicSignatureGameV3ArbitrumError)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "ArbitrumError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3BidCstRewardAmountMultiplierChangedIterator is returned from FilterBidCstRewardAmountMultiplierChanged and is used to iterate over the raw logs and unpacked data for BidCstRewardAmountMultiplierChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3BidCstRewardAmountMultiplierChangedIterator struct {
	Event *CosmicSignatureGameV3BidCstRewardAmountMultiplierChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3BidCstRewardAmountMultiplierChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3BidCstRewardAmountMultiplierChanged)
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
		it.Event = new(CosmicSignatureGameV3BidCstRewardAmountMultiplierChanged)
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
func (it *CosmicSignatureGameV3BidCstRewardAmountMultiplierChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3BidCstRewardAmountMultiplierChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3BidCstRewardAmountMultiplierChanged represents a BidCstRewardAmountMultiplierChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3BidCstRewardAmountMultiplierChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBidCstRewardAmountMultiplierChanged is a free log retrieval operation binding the contract event 0x40b9c59af8c486ccf8c7cc73df5a51e7cc29747ea7d39f99632ecaf9caa2ed1f.
//
// Solidity: event BidCstRewardAmountMultiplierChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterBidCstRewardAmountMultiplierChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3BidCstRewardAmountMultiplierChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "BidCstRewardAmountMultiplierChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3BidCstRewardAmountMultiplierChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "BidCstRewardAmountMultiplierChanged", logs: logs, sub: sub}, nil
}

// WatchBidCstRewardAmountMultiplierChanged is a free log subscription operation binding the contract event 0x40b9c59af8c486ccf8c7cc73df5a51e7cc29747ea7d39f99632ecaf9caa2ed1f.
//
// Solidity: event BidCstRewardAmountMultiplierChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchBidCstRewardAmountMultiplierChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3BidCstRewardAmountMultiplierChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "BidCstRewardAmountMultiplierChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3BidCstRewardAmountMultiplierChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "BidCstRewardAmountMultiplierChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseBidCstRewardAmountMultiplierChanged(log types.Log) (*CosmicSignatureGameV3BidCstRewardAmountMultiplierChanged, error) {
	event := new(CosmicSignatureGameV3BidCstRewardAmountMultiplierChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "BidCstRewardAmountMultiplierChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3BidCstRewardAmountPerMinuteChangedIterator is returned from FilterBidCstRewardAmountPerMinuteChanged and is used to iterate over the raw logs and unpacked data for BidCstRewardAmountPerMinuteChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3BidCstRewardAmountPerMinuteChangedIterator struct {
	Event *CosmicSignatureGameV3BidCstRewardAmountPerMinuteChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3BidCstRewardAmountPerMinuteChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3BidCstRewardAmountPerMinuteChanged)
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
		it.Event = new(CosmicSignatureGameV3BidCstRewardAmountPerMinuteChanged)
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
func (it *CosmicSignatureGameV3BidCstRewardAmountPerMinuteChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3BidCstRewardAmountPerMinuteChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3BidCstRewardAmountPerMinuteChanged represents a BidCstRewardAmountPerMinuteChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3BidCstRewardAmountPerMinuteChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBidCstRewardAmountPerMinuteChanged is a free log retrieval operation binding the contract event 0x961b02838cd95976d0cac2e65ed131f45e19a84369d91d59d613dec94a0638c6.
//
// Solidity: event BidCstRewardAmountPerMinuteChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterBidCstRewardAmountPerMinuteChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3BidCstRewardAmountPerMinuteChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "BidCstRewardAmountPerMinuteChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3BidCstRewardAmountPerMinuteChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "BidCstRewardAmountPerMinuteChanged", logs: logs, sub: sub}, nil
}

// WatchBidCstRewardAmountPerMinuteChanged is a free log subscription operation binding the contract event 0x961b02838cd95976d0cac2e65ed131f45e19a84369d91d59d613dec94a0638c6.
//
// Solidity: event BidCstRewardAmountPerMinuteChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchBidCstRewardAmountPerMinuteChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3BidCstRewardAmountPerMinuteChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "BidCstRewardAmountPerMinuteChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3BidCstRewardAmountPerMinuteChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "BidCstRewardAmountPerMinuteChanged", log); err != nil {
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

// ParseBidCstRewardAmountPerMinuteChanged is a log parse operation binding the contract event 0x961b02838cd95976d0cac2e65ed131f45e19a84369d91d59d613dec94a0638c6.
//
// Solidity: event BidCstRewardAmountPerMinuteChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseBidCstRewardAmountPerMinuteChanged(log types.Log) (*CosmicSignatureGameV3BidCstRewardAmountPerMinuteChanged, error) {
	event := new(CosmicSignatureGameV3BidCstRewardAmountPerMinuteChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "BidCstRewardAmountPerMinuteChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3BidMessageLengthMaxLimitChangedIterator is returned from FilterBidMessageLengthMaxLimitChanged and is used to iterate over the raw logs and unpacked data for BidMessageLengthMaxLimitChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3BidMessageLengthMaxLimitChangedIterator struct {
	Event *CosmicSignatureGameV3BidMessageLengthMaxLimitChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3BidMessageLengthMaxLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3BidMessageLengthMaxLimitChanged)
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
		it.Event = new(CosmicSignatureGameV3BidMessageLengthMaxLimitChanged)
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
func (it *CosmicSignatureGameV3BidMessageLengthMaxLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3BidMessageLengthMaxLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3BidMessageLengthMaxLimitChanged represents a BidMessageLengthMaxLimitChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3BidMessageLengthMaxLimitChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBidMessageLengthMaxLimitChanged is a free log retrieval operation binding the contract event 0x157c413b0549fd4f45aab72b7828304fb2c45dad53de0f1128c5eabf3aaabaf8.
//
// Solidity: event BidMessageLengthMaxLimitChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterBidMessageLengthMaxLimitChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3BidMessageLengthMaxLimitChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "BidMessageLengthMaxLimitChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3BidMessageLengthMaxLimitChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "BidMessageLengthMaxLimitChanged", logs: logs, sub: sub}, nil
}

// WatchBidMessageLengthMaxLimitChanged is a free log subscription operation binding the contract event 0x157c413b0549fd4f45aab72b7828304fb2c45dad53de0f1128c5eabf3aaabaf8.
//
// Solidity: event BidMessageLengthMaxLimitChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchBidMessageLengthMaxLimitChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3BidMessageLengthMaxLimitChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "BidMessageLengthMaxLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3BidMessageLengthMaxLimitChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "BidMessageLengthMaxLimitChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseBidMessageLengthMaxLimitChanged(log types.Log) (*CosmicSignatureGameV3BidMessageLengthMaxLimitChanged, error) {
	event := new(CosmicSignatureGameV3BidMessageLengthMaxLimitChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "BidMessageLengthMaxLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3BidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3BidPlacedIterator struct {
	Event *CosmicSignatureGameV3BidPlaced // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3BidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3BidPlaced)
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
		it.Event = new(CosmicSignatureGameV3BidPlaced)
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
func (it *CosmicSignatureGameV3BidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3BidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3BidPlaced represents a BidPlaced event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3BidPlaced struct {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterBidPlaced(opts *bind.FilterOpts, roundNum []*big.Int, lastBidderAddress []common.Address, randomWalkNftId []*big.Int) (*CosmicSignatureGameV3BidPlacedIterator, error) {

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

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "BidPlaced", roundNumRule, lastBidderAddressRule, randomWalkNftIdRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3BidPlacedIterator{contract: _CosmicSignatureGameV3.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x1d1f406c89e99504a7222aaba50bf96b9f91f522ebc6ea825255145919e102ec.
//
// Solidity: event BidPlaced(uint256 indexed roundNum, address indexed lastBidderAddress, int256 paidEthPrice, int256 paidCstPrice, int256 indexed randomWalkNftId, string message, uint256 bidCstRewardAmount, uint256 cstDutchAuctionDuration, uint256 mainPrizeTime)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3BidPlaced, roundNum []*big.Int, lastBidderAddress []common.Address, randomWalkNftId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "BidPlaced", roundNumRule, lastBidderAddressRule, randomWalkNftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3BidPlaced)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseBidPlaced(log types.Log) (*CosmicSignatureGameV3BidPlaced, error) {
	event := new(CosmicSignatureGameV3BidPlaced)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3CharityAddressChangedIterator is returned from FilterCharityAddressChanged and is used to iterate over the raw logs and unpacked data for CharityAddressChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3CharityAddressChangedIterator struct {
	Event *CosmicSignatureGameV3CharityAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3CharityAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3CharityAddressChanged)
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
		it.Event = new(CosmicSignatureGameV3CharityAddressChanged)
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
func (it *CosmicSignatureGameV3CharityAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3CharityAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3CharityAddressChanged represents a CharityAddressChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3CharityAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCharityAddressChanged is a free log retrieval operation binding the contract event 0x1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c.
//
// Solidity: event CharityAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterCharityAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*CosmicSignatureGameV3CharityAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "CharityAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3CharityAddressChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "CharityAddressChanged", logs: logs, sub: sub}, nil
}

// WatchCharityAddressChanged is a free log subscription operation binding the contract event 0x1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c.
//
// Solidity: event CharityAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchCharityAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3CharityAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "CharityAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3CharityAddressChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "CharityAddressChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseCharityAddressChanged(log types.Log) (*CosmicSignatureGameV3CharityAddressChanged, error) {
	event := new(CosmicSignatureGameV3CharityAddressChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "CharityAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3CharityEthDonationAmountPercentageChangedIterator is returned from FilterCharityEthDonationAmountPercentageChanged and is used to iterate over the raw logs and unpacked data for CharityEthDonationAmountPercentageChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3CharityEthDonationAmountPercentageChangedIterator struct {
	Event *CosmicSignatureGameV3CharityEthDonationAmountPercentageChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3CharityEthDonationAmountPercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3CharityEthDonationAmountPercentageChanged)
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
		it.Event = new(CosmicSignatureGameV3CharityEthDonationAmountPercentageChanged)
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
func (it *CosmicSignatureGameV3CharityEthDonationAmountPercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3CharityEthDonationAmountPercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3CharityEthDonationAmountPercentageChanged represents a CharityEthDonationAmountPercentageChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3CharityEthDonationAmountPercentageChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCharityEthDonationAmountPercentageChanged is a free log retrieval operation binding the contract event 0xfe65b6d5007c66dadebc5095104ccd672c070f396dfdcfe1ee7f54201b9efa6d.
//
// Solidity: event CharityEthDonationAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterCharityEthDonationAmountPercentageChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3CharityEthDonationAmountPercentageChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "CharityEthDonationAmountPercentageChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3CharityEthDonationAmountPercentageChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "CharityEthDonationAmountPercentageChanged", logs: logs, sub: sub}, nil
}

// WatchCharityEthDonationAmountPercentageChanged is a free log subscription operation binding the contract event 0xfe65b6d5007c66dadebc5095104ccd672c070f396dfdcfe1ee7f54201b9efa6d.
//
// Solidity: event CharityEthDonationAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchCharityEthDonationAmountPercentageChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3CharityEthDonationAmountPercentageChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "CharityEthDonationAmountPercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3CharityEthDonationAmountPercentageChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "CharityEthDonationAmountPercentageChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseCharityEthDonationAmountPercentageChanged(log types.Log) (*CosmicSignatureGameV3CharityEthDonationAmountPercentageChanged, error) {
	event := new(CosmicSignatureGameV3CharityEthDonationAmountPercentageChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "CharityEthDonationAmountPercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3ChronoWarriorEthPrizeAmountPercentageChangedIterator is returned from FilterChronoWarriorEthPrizeAmountPercentageChanged and is used to iterate over the raw logs and unpacked data for ChronoWarriorEthPrizeAmountPercentageChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3ChronoWarriorEthPrizeAmountPercentageChangedIterator struct {
	Event *CosmicSignatureGameV3ChronoWarriorEthPrizeAmountPercentageChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3ChronoWarriorEthPrizeAmountPercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3ChronoWarriorEthPrizeAmountPercentageChanged)
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
		it.Event = new(CosmicSignatureGameV3ChronoWarriorEthPrizeAmountPercentageChanged)
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
func (it *CosmicSignatureGameV3ChronoWarriorEthPrizeAmountPercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3ChronoWarriorEthPrizeAmountPercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3ChronoWarriorEthPrizeAmountPercentageChanged represents a ChronoWarriorEthPrizeAmountPercentageChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3ChronoWarriorEthPrizeAmountPercentageChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChronoWarriorEthPrizeAmountPercentageChanged is a free log retrieval operation binding the contract event 0x5581e31f5b8d4b3c45b8ab8bf67e3602ce588b423905eb6ad34bd6bc3c848699.
//
// Solidity: event ChronoWarriorEthPrizeAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterChronoWarriorEthPrizeAmountPercentageChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3ChronoWarriorEthPrizeAmountPercentageChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "ChronoWarriorEthPrizeAmountPercentageChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3ChronoWarriorEthPrizeAmountPercentageChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "ChronoWarriorEthPrizeAmountPercentageChanged", logs: logs, sub: sub}, nil
}

// WatchChronoWarriorEthPrizeAmountPercentageChanged is a free log subscription operation binding the contract event 0x5581e31f5b8d4b3c45b8ab8bf67e3602ce588b423905eb6ad34bd6bc3c848699.
//
// Solidity: event ChronoWarriorEthPrizeAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchChronoWarriorEthPrizeAmountPercentageChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3ChronoWarriorEthPrizeAmountPercentageChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "ChronoWarriorEthPrizeAmountPercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3ChronoWarriorEthPrizeAmountPercentageChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "ChronoWarriorEthPrizeAmountPercentageChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseChronoWarriorEthPrizeAmountPercentageChanged(log types.Log) (*CosmicSignatureGameV3ChronoWarriorEthPrizeAmountPercentageChanged, error) {
	event := new(CosmicSignatureGameV3ChronoWarriorEthPrizeAmountPercentageChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "ChronoWarriorEthPrizeAmountPercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3ChronoWarriorPrizePaidIterator is returned from FilterChronoWarriorPrizePaid and is used to iterate over the raw logs and unpacked data for ChronoWarriorPrizePaid events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3ChronoWarriorPrizePaidIterator struct {
	Event *CosmicSignatureGameV3ChronoWarriorPrizePaid // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3ChronoWarriorPrizePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3ChronoWarriorPrizePaid)
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
		it.Event = new(CosmicSignatureGameV3ChronoWarriorPrizePaid)
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
func (it *CosmicSignatureGameV3ChronoWarriorPrizePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3ChronoWarriorPrizePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3ChronoWarriorPrizePaid represents a ChronoWarriorPrizePaid event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3ChronoWarriorPrizePaid struct {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterChronoWarriorPrizePaid(opts *bind.FilterOpts, roundNum []*big.Int, chronoWarriorAddress []common.Address, prizeCosmicSignatureNftId []*big.Int) (*CosmicSignatureGameV3ChronoWarriorPrizePaidIterator, error) {

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

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "ChronoWarriorPrizePaid", roundNumRule, chronoWarriorAddressRule, prizeCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3ChronoWarriorPrizePaidIterator{contract: _CosmicSignatureGameV3.contract, event: "ChronoWarriorPrizePaid", logs: logs, sub: sub}, nil
}

// WatchChronoWarriorPrizePaid is a free log subscription operation binding the contract event 0xaa858ae20a26d00a9ea528972d537e68a51a0744226d9ea1fc9b22492dc282a5.
//
// Solidity: event ChronoWarriorPrizePaid(uint256 indexed roundNum, uint256 winnerIndex, address indexed chronoWarriorAddress, uint256 ethPrizeAmount, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchChronoWarriorPrizePaid(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3ChronoWarriorPrizePaid, roundNum []*big.Int, chronoWarriorAddress []common.Address, prizeCosmicSignatureNftId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "ChronoWarriorPrizePaid", roundNumRule, chronoWarriorAddressRule, prizeCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3ChronoWarriorPrizePaid)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "ChronoWarriorPrizePaid", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseChronoWarriorPrizePaid(log types.Log) (*CosmicSignatureGameV3ChronoWarriorPrizePaid, error) {
	event := new(CosmicSignatureGameV3ChronoWarriorPrizePaid)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "ChronoWarriorPrizePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3CosmicSignatureNftAddressChangedIterator is returned from FilterCosmicSignatureNftAddressChanged and is used to iterate over the raw logs and unpacked data for CosmicSignatureNftAddressChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3CosmicSignatureNftAddressChangedIterator struct {
	Event *CosmicSignatureGameV3CosmicSignatureNftAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3CosmicSignatureNftAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3CosmicSignatureNftAddressChanged)
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
		it.Event = new(CosmicSignatureGameV3CosmicSignatureNftAddressChanged)
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
func (it *CosmicSignatureGameV3CosmicSignatureNftAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3CosmicSignatureNftAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3CosmicSignatureNftAddressChanged represents a CosmicSignatureNftAddressChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3CosmicSignatureNftAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCosmicSignatureNftAddressChanged is a free log retrieval operation binding the contract event 0x5bde6238168795ba4e77972a2bdaa5a465f7c9a5d05817f5e8d3fed2e5a4fa60.
//
// Solidity: event CosmicSignatureNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterCosmicSignatureNftAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*CosmicSignatureGameV3CosmicSignatureNftAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "CosmicSignatureNftAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3CosmicSignatureNftAddressChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "CosmicSignatureNftAddressChanged", logs: logs, sub: sub}, nil
}

// WatchCosmicSignatureNftAddressChanged is a free log subscription operation binding the contract event 0x5bde6238168795ba4e77972a2bdaa5a465f7c9a5d05817f5e8d3fed2e5a4fa60.
//
// Solidity: event CosmicSignatureNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchCosmicSignatureNftAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3CosmicSignatureNftAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "CosmicSignatureNftAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3CosmicSignatureNftAddressChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "CosmicSignatureNftAddressChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseCosmicSignatureNftAddressChanged(log types.Log) (*CosmicSignatureGameV3CosmicSignatureNftAddressChanged, error) {
	event := new(CosmicSignatureGameV3CosmicSignatureNftAddressChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "CosmicSignatureNftAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3CosmicSignatureNftStakingTotalEthRewardAmountPercentageChangedIterator is returned from FilterCosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged and is used to iterate over the raw logs and unpacked data for CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3CosmicSignatureNftStakingTotalEthRewardAmountPercentageChangedIterator struct {
	Event *CosmicSignatureGameV3CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3CosmicSignatureNftStakingTotalEthRewardAmountPercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged)
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
		it.Event = new(CosmicSignatureGameV3CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged)
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
func (it *CosmicSignatureGameV3CosmicSignatureNftStakingTotalEthRewardAmountPercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3CosmicSignatureNftStakingTotalEthRewardAmountPercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged represents a CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged is a free log retrieval operation binding the contract event 0x9e44c04f534af356419a731b967f3d56bc748b1f3fdbda7a89f4e1a321ade934.
//
// Solidity: event CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterCosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3CosmicSignatureNftStakingTotalEthRewardAmountPercentageChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3CosmicSignatureNftStakingTotalEthRewardAmountPercentageChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged", logs: logs, sub: sub}, nil
}

// WatchCosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged is a free log subscription operation binding the contract event 0x9e44c04f534af356419a731b967f3d56bc748b1f3fdbda7a89f4e1a321ade934.
//
// Solidity: event CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchCosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseCosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged(log types.Log) (*CosmicSignatureGameV3CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged, error) {
	event := new(CosmicSignatureGameV3CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3CosmicSignatureTokenAddressChangedIterator is returned from FilterCosmicSignatureTokenAddressChanged and is used to iterate over the raw logs and unpacked data for CosmicSignatureTokenAddressChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3CosmicSignatureTokenAddressChangedIterator struct {
	Event *CosmicSignatureGameV3CosmicSignatureTokenAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3CosmicSignatureTokenAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3CosmicSignatureTokenAddressChanged)
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
		it.Event = new(CosmicSignatureGameV3CosmicSignatureTokenAddressChanged)
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
func (it *CosmicSignatureGameV3CosmicSignatureTokenAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3CosmicSignatureTokenAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3CosmicSignatureTokenAddressChanged represents a CosmicSignatureTokenAddressChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3CosmicSignatureTokenAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCosmicSignatureTokenAddressChanged is a free log retrieval operation binding the contract event 0x9b3eda10f1724a2cf9f7dae4ac263c77908df4d00e92f1377b66fc8be37cd8c1.
//
// Solidity: event CosmicSignatureTokenAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterCosmicSignatureTokenAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*CosmicSignatureGameV3CosmicSignatureTokenAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "CosmicSignatureTokenAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3CosmicSignatureTokenAddressChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "CosmicSignatureTokenAddressChanged", logs: logs, sub: sub}, nil
}

// WatchCosmicSignatureTokenAddressChanged is a free log subscription operation binding the contract event 0x9b3eda10f1724a2cf9f7dae4ac263c77908df4d00e92f1377b66fc8be37cd8c1.
//
// Solidity: event CosmicSignatureTokenAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchCosmicSignatureTokenAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3CosmicSignatureTokenAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "CosmicSignatureTokenAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3CosmicSignatureTokenAddressChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "CosmicSignatureTokenAddressChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseCosmicSignatureTokenAddressChanged(log types.Log) (*CosmicSignatureGameV3CosmicSignatureTokenAddressChanged, error) {
	event := new(CosmicSignatureGameV3CosmicSignatureTokenAddressChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "CosmicSignatureTokenAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3CstDutchAuctionBeginningBidPriceMinLimitChangedIterator is returned from FilterCstDutchAuctionBeginningBidPriceMinLimitChanged and is used to iterate over the raw logs and unpacked data for CstDutchAuctionBeginningBidPriceMinLimitChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3CstDutchAuctionBeginningBidPriceMinLimitChangedIterator struct {
	Event *CosmicSignatureGameV3CstDutchAuctionBeginningBidPriceMinLimitChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3CstDutchAuctionBeginningBidPriceMinLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3CstDutchAuctionBeginningBidPriceMinLimitChanged)
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
		it.Event = new(CosmicSignatureGameV3CstDutchAuctionBeginningBidPriceMinLimitChanged)
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
func (it *CosmicSignatureGameV3CstDutchAuctionBeginningBidPriceMinLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3CstDutchAuctionBeginningBidPriceMinLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3CstDutchAuctionBeginningBidPriceMinLimitChanged represents a CstDutchAuctionBeginningBidPriceMinLimitChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3CstDutchAuctionBeginningBidPriceMinLimitChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCstDutchAuctionBeginningBidPriceMinLimitChanged is a free log retrieval operation binding the contract event 0x4e8c80fe79d13b8663de9f6981925ae24c498cc07b0ebd4f4789fa78ca17caff.
//
// Solidity: event CstDutchAuctionBeginningBidPriceMinLimitChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterCstDutchAuctionBeginningBidPriceMinLimitChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3CstDutchAuctionBeginningBidPriceMinLimitChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "CstDutchAuctionBeginningBidPriceMinLimitChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3CstDutchAuctionBeginningBidPriceMinLimitChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "CstDutchAuctionBeginningBidPriceMinLimitChanged", logs: logs, sub: sub}, nil
}

// WatchCstDutchAuctionBeginningBidPriceMinLimitChanged is a free log subscription operation binding the contract event 0x4e8c80fe79d13b8663de9f6981925ae24c498cc07b0ebd4f4789fa78ca17caff.
//
// Solidity: event CstDutchAuctionBeginningBidPriceMinLimitChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchCstDutchAuctionBeginningBidPriceMinLimitChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3CstDutchAuctionBeginningBidPriceMinLimitChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "CstDutchAuctionBeginningBidPriceMinLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3CstDutchAuctionBeginningBidPriceMinLimitChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "CstDutchAuctionBeginningBidPriceMinLimitChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseCstDutchAuctionBeginningBidPriceMinLimitChanged(log types.Log) (*CosmicSignatureGameV3CstDutchAuctionBeginningBidPriceMinLimitChanged, error) {
	event := new(CosmicSignatureGameV3CstDutchAuctionBeginningBidPriceMinLimitChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "CstDutchAuctionBeginningBidPriceMinLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3CstDutchAuctionDurationChangeDivisorChangedIterator is returned from FilterCstDutchAuctionDurationChangeDivisorChanged and is used to iterate over the raw logs and unpacked data for CstDutchAuctionDurationChangeDivisorChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3CstDutchAuctionDurationChangeDivisorChangedIterator struct {
	Event *CosmicSignatureGameV3CstDutchAuctionDurationChangeDivisorChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3CstDutchAuctionDurationChangeDivisorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3CstDutchAuctionDurationChangeDivisorChanged)
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
		it.Event = new(CosmicSignatureGameV3CstDutchAuctionDurationChangeDivisorChanged)
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
func (it *CosmicSignatureGameV3CstDutchAuctionDurationChangeDivisorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3CstDutchAuctionDurationChangeDivisorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3CstDutchAuctionDurationChangeDivisorChanged represents a CstDutchAuctionDurationChangeDivisorChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3CstDutchAuctionDurationChangeDivisorChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCstDutchAuctionDurationChangeDivisorChanged is a free log retrieval operation binding the contract event 0xacbc6b6929088e4b2d043625fa7248e00fb6658a425eb9d9dc8c37b18c7f3e6f.
//
// Solidity: event CstDutchAuctionDurationChangeDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterCstDutchAuctionDurationChangeDivisorChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3CstDutchAuctionDurationChangeDivisorChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "CstDutchAuctionDurationChangeDivisorChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3CstDutchAuctionDurationChangeDivisorChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "CstDutchAuctionDurationChangeDivisorChanged", logs: logs, sub: sub}, nil
}

// WatchCstDutchAuctionDurationChangeDivisorChanged is a free log subscription operation binding the contract event 0xacbc6b6929088e4b2d043625fa7248e00fb6658a425eb9d9dc8c37b18c7f3e6f.
//
// Solidity: event CstDutchAuctionDurationChangeDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchCstDutchAuctionDurationChangeDivisorChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3CstDutchAuctionDurationChangeDivisorChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "CstDutchAuctionDurationChangeDivisorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3CstDutchAuctionDurationChangeDivisorChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "CstDutchAuctionDurationChangeDivisorChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseCstDutchAuctionDurationChangeDivisorChanged(log types.Log) (*CosmicSignatureGameV3CstDutchAuctionDurationChangeDivisorChanged, error) {
	event := new(CosmicSignatureGameV3CstDutchAuctionDurationChangeDivisorChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "CstDutchAuctionDurationChangeDivisorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3CstDutchAuctionDurationChangedIterator is returned from FilterCstDutchAuctionDurationChanged and is used to iterate over the raw logs and unpacked data for CstDutchAuctionDurationChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3CstDutchAuctionDurationChangedIterator struct {
	Event *CosmicSignatureGameV3CstDutchAuctionDurationChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3CstDutchAuctionDurationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3CstDutchAuctionDurationChanged)
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
		it.Event = new(CosmicSignatureGameV3CstDutchAuctionDurationChanged)
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
func (it *CosmicSignatureGameV3CstDutchAuctionDurationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3CstDutchAuctionDurationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3CstDutchAuctionDurationChanged represents a CstDutchAuctionDurationChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3CstDutchAuctionDurationChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCstDutchAuctionDurationChanged is a free log retrieval operation binding the contract event 0x4abea08c196329c357e3175d011af39a8625be99ef0ba5a0f3547a95534fedb7.
//
// Solidity: event CstDutchAuctionDurationChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterCstDutchAuctionDurationChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3CstDutchAuctionDurationChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "CstDutchAuctionDurationChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3CstDutchAuctionDurationChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "CstDutchAuctionDurationChanged", logs: logs, sub: sub}, nil
}

// WatchCstDutchAuctionDurationChanged is a free log subscription operation binding the contract event 0x4abea08c196329c357e3175d011af39a8625be99ef0ba5a0f3547a95534fedb7.
//
// Solidity: event CstDutchAuctionDurationChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchCstDutchAuctionDurationChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3CstDutchAuctionDurationChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "CstDutchAuctionDurationChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3CstDutchAuctionDurationChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "CstDutchAuctionDurationChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseCstDutchAuctionDurationChanged(log types.Log) (*CosmicSignatureGameV3CstDutchAuctionDurationChanged, error) {
	event := new(CosmicSignatureGameV3CstDutchAuctionDurationChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "CstDutchAuctionDurationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3CstPrizeAmountChangedIterator is returned from FilterCstPrizeAmountChanged and is used to iterate over the raw logs and unpacked data for CstPrizeAmountChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3CstPrizeAmountChangedIterator struct {
	Event *CosmicSignatureGameV3CstPrizeAmountChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3CstPrizeAmountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3CstPrizeAmountChanged)
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
		it.Event = new(CosmicSignatureGameV3CstPrizeAmountChanged)
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
func (it *CosmicSignatureGameV3CstPrizeAmountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3CstPrizeAmountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3CstPrizeAmountChanged represents a CstPrizeAmountChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3CstPrizeAmountChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCstPrizeAmountChanged is a free log retrieval operation binding the contract event 0xd95e7f967f9370c11deb15ffbb191b9f2e9795ab0738db5bc72bd2794978f32d.
//
// Solidity: event CstPrizeAmountChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterCstPrizeAmountChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3CstPrizeAmountChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "CstPrizeAmountChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3CstPrizeAmountChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "CstPrizeAmountChanged", logs: logs, sub: sub}, nil
}

// WatchCstPrizeAmountChanged is a free log subscription operation binding the contract event 0xd95e7f967f9370c11deb15ffbb191b9f2e9795ab0738db5bc72bd2794978f32d.
//
// Solidity: event CstPrizeAmountChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchCstPrizeAmountChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3CstPrizeAmountChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "CstPrizeAmountChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3CstPrizeAmountChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "CstPrizeAmountChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseCstPrizeAmountChanged(log types.Log) (*CosmicSignatureGameV3CstPrizeAmountChanged, error) {
	event := new(CosmicSignatureGameV3CstPrizeAmountChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "CstPrizeAmountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3DelayDurationBeforeRoundActivationChangedIterator is returned from FilterDelayDurationBeforeRoundActivationChanged and is used to iterate over the raw logs and unpacked data for DelayDurationBeforeRoundActivationChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3DelayDurationBeforeRoundActivationChangedIterator struct {
	Event *CosmicSignatureGameV3DelayDurationBeforeRoundActivationChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3DelayDurationBeforeRoundActivationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3DelayDurationBeforeRoundActivationChanged)
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
		it.Event = new(CosmicSignatureGameV3DelayDurationBeforeRoundActivationChanged)
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
func (it *CosmicSignatureGameV3DelayDurationBeforeRoundActivationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3DelayDurationBeforeRoundActivationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3DelayDurationBeforeRoundActivationChanged represents a DelayDurationBeforeRoundActivationChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3DelayDurationBeforeRoundActivationChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelayDurationBeforeRoundActivationChanged is a free log retrieval operation binding the contract event 0xb0868a729f47ae3829aaafe3ca2975d3db2148553c854112f598be6d91ef0d28.
//
// Solidity: event DelayDurationBeforeRoundActivationChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterDelayDurationBeforeRoundActivationChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3DelayDurationBeforeRoundActivationChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "DelayDurationBeforeRoundActivationChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3DelayDurationBeforeRoundActivationChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "DelayDurationBeforeRoundActivationChanged", logs: logs, sub: sub}, nil
}

// WatchDelayDurationBeforeRoundActivationChanged is a free log subscription operation binding the contract event 0xb0868a729f47ae3829aaafe3ca2975d3db2148553c854112f598be6d91ef0d28.
//
// Solidity: event DelayDurationBeforeRoundActivationChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchDelayDurationBeforeRoundActivationChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3DelayDurationBeforeRoundActivationChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "DelayDurationBeforeRoundActivationChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3DelayDurationBeforeRoundActivationChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "DelayDurationBeforeRoundActivationChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseDelayDurationBeforeRoundActivationChanged(log types.Log) (*CosmicSignatureGameV3DelayDurationBeforeRoundActivationChanged, error) {
	event := new(CosmicSignatureGameV3DelayDurationBeforeRoundActivationChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "DelayDurationBeforeRoundActivationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3EnduranceChampionPrizePaidIterator is returned from FilterEnduranceChampionPrizePaid and is used to iterate over the raw logs and unpacked data for EnduranceChampionPrizePaid events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3EnduranceChampionPrizePaidIterator struct {
	Event *CosmicSignatureGameV3EnduranceChampionPrizePaid // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3EnduranceChampionPrizePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3EnduranceChampionPrizePaid)
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
		it.Event = new(CosmicSignatureGameV3EnduranceChampionPrizePaid)
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
func (it *CosmicSignatureGameV3EnduranceChampionPrizePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3EnduranceChampionPrizePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3EnduranceChampionPrizePaid represents a EnduranceChampionPrizePaid event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3EnduranceChampionPrizePaid struct {
	RoundNum                  *big.Int
	EnduranceChampionAddress  common.Address
	CstPrizeAmount            *big.Int
	PrizeCosmicSignatureNftId *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterEnduranceChampionPrizePaid is a free log retrieval operation binding the contract event 0x838ec9dd2530548892bff113f5ffb0138d2efc63c7f59bb6571e8c923b631260.
//
// Solidity: event EnduranceChampionPrizePaid(uint256 indexed roundNum, address indexed enduranceChampionAddress, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterEnduranceChampionPrizePaid(opts *bind.FilterOpts, roundNum []*big.Int, enduranceChampionAddress []common.Address, prizeCosmicSignatureNftId []*big.Int) (*CosmicSignatureGameV3EnduranceChampionPrizePaidIterator, error) {

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

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "EnduranceChampionPrizePaid", roundNumRule, enduranceChampionAddressRule, prizeCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3EnduranceChampionPrizePaidIterator{contract: _CosmicSignatureGameV3.contract, event: "EnduranceChampionPrizePaid", logs: logs, sub: sub}, nil
}

// WatchEnduranceChampionPrizePaid is a free log subscription operation binding the contract event 0x838ec9dd2530548892bff113f5ffb0138d2efc63c7f59bb6571e8c923b631260.
//
// Solidity: event EnduranceChampionPrizePaid(uint256 indexed roundNum, address indexed enduranceChampionAddress, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchEnduranceChampionPrizePaid(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3EnduranceChampionPrizePaid, roundNum []*big.Int, enduranceChampionAddress []common.Address, prizeCosmicSignatureNftId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "EnduranceChampionPrizePaid", roundNumRule, enduranceChampionAddressRule, prizeCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3EnduranceChampionPrizePaid)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "EnduranceChampionPrizePaid", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseEnduranceChampionPrizePaid(log types.Log) (*CosmicSignatureGameV3EnduranceChampionPrizePaid, error) {
	event := new(CosmicSignatureGameV3EnduranceChampionPrizePaid)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "EnduranceChampionPrizePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3EthBidPriceIncreaseDivisorChangedIterator is returned from FilterEthBidPriceIncreaseDivisorChanged and is used to iterate over the raw logs and unpacked data for EthBidPriceIncreaseDivisorChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3EthBidPriceIncreaseDivisorChangedIterator struct {
	Event *CosmicSignatureGameV3EthBidPriceIncreaseDivisorChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3EthBidPriceIncreaseDivisorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3EthBidPriceIncreaseDivisorChanged)
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
		it.Event = new(CosmicSignatureGameV3EthBidPriceIncreaseDivisorChanged)
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
func (it *CosmicSignatureGameV3EthBidPriceIncreaseDivisorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3EthBidPriceIncreaseDivisorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3EthBidPriceIncreaseDivisorChanged represents a EthBidPriceIncreaseDivisorChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3EthBidPriceIncreaseDivisorChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthBidPriceIncreaseDivisorChanged is a free log retrieval operation binding the contract event 0xdeb71e1d117914859ddde484a9810084d9ee399077d5cd8dcbdf8770d33d8ae4.
//
// Solidity: event EthBidPriceIncreaseDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterEthBidPriceIncreaseDivisorChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3EthBidPriceIncreaseDivisorChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "EthBidPriceIncreaseDivisorChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3EthBidPriceIncreaseDivisorChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "EthBidPriceIncreaseDivisorChanged", logs: logs, sub: sub}, nil
}

// WatchEthBidPriceIncreaseDivisorChanged is a free log subscription operation binding the contract event 0xdeb71e1d117914859ddde484a9810084d9ee399077d5cd8dcbdf8770d33d8ae4.
//
// Solidity: event EthBidPriceIncreaseDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchEthBidPriceIncreaseDivisorChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3EthBidPriceIncreaseDivisorChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "EthBidPriceIncreaseDivisorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3EthBidPriceIncreaseDivisorChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "EthBidPriceIncreaseDivisorChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseEthBidPriceIncreaseDivisorChanged(log types.Log) (*CosmicSignatureGameV3EthBidPriceIncreaseDivisorChanged, error) {
	event := new(CosmicSignatureGameV3EthBidPriceIncreaseDivisorChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "EthBidPriceIncreaseDivisorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3EthBidRefundAmountInGasToSwallowMaxLimitChangedIterator is returned from FilterEthBidRefundAmountInGasToSwallowMaxLimitChanged and is used to iterate over the raw logs and unpacked data for EthBidRefundAmountInGasToSwallowMaxLimitChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3EthBidRefundAmountInGasToSwallowMaxLimitChangedIterator struct {
	Event *CosmicSignatureGameV3EthBidRefundAmountInGasToSwallowMaxLimitChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3EthBidRefundAmountInGasToSwallowMaxLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3EthBidRefundAmountInGasToSwallowMaxLimitChanged)
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
		it.Event = new(CosmicSignatureGameV3EthBidRefundAmountInGasToSwallowMaxLimitChanged)
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
func (it *CosmicSignatureGameV3EthBidRefundAmountInGasToSwallowMaxLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3EthBidRefundAmountInGasToSwallowMaxLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3EthBidRefundAmountInGasToSwallowMaxLimitChanged represents a EthBidRefundAmountInGasToSwallowMaxLimitChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3EthBidRefundAmountInGasToSwallowMaxLimitChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthBidRefundAmountInGasToSwallowMaxLimitChanged is a free log retrieval operation binding the contract event 0xa787f26546d7eeea63d70fc31736f27ad28329e95982f3bc5a7e0280f497bbf5.
//
// Solidity: event EthBidRefundAmountInGasToSwallowMaxLimitChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterEthBidRefundAmountInGasToSwallowMaxLimitChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3EthBidRefundAmountInGasToSwallowMaxLimitChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "EthBidRefundAmountInGasToSwallowMaxLimitChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3EthBidRefundAmountInGasToSwallowMaxLimitChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "EthBidRefundAmountInGasToSwallowMaxLimitChanged", logs: logs, sub: sub}, nil
}

// WatchEthBidRefundAmountInGasToSwallowMaxLimitChanged is a free log subscription operation binding the contract event 0xa787f26546d7eeea63d70fc31736f27ad28329e95982f3bc5a7e0280f497bbf5.
//
// Solidity: event EthBidRefundAmountInGasToSwallowMaxLimitChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchEthBidRefundAmountInGasToSwallowMaxLimitChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3EthBidRefundAmountInGasToSwallowMaxLimitChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "EthBidRefundAmountInGasToSwallowMaxLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3EthBidRefundAmountInGasToSwallowMaxLimitChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "EthBidRefundAmountInGasToSwallowMaxLimitChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseEthBidRefundAmountInGasToSwallowMaxLimitChanged(log types.Log) (*CosmicSignatureGameV3EthBidRefundAmountInGasToSwallowMaxLimitChanged, error) {
	event := new(CosmicSignatureGameV3EthBidRefundAmountInGasToSwallowMaxLimitChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "EthBidRefundAmountInGasToSwallowMaxLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3EthDonatedIterator is returned from FilterEthDonated and is used to iterate over the raw logs and unpacked data for EthDonated events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3EthDonatedIterator struct {
	Event *CosmicSignatureGameV3EthDonated // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3EthDonatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3EthDonated)
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
		it.Event = new(CosmicSignatureGameV3EthDonated)
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
func (it *CosmicSignatureGameV3EthDonatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3EthDonatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3EthDonated represents a EthDonated event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3EthDonated struct {
	RoundNum     *big.Int
	DonorAddress common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEthDonated is a free log retrieval operation binding the contract event 0xe32cacf203d00685e2b4d8b0a90e7cd8f3f8a208fdf116f4bb36abe08b7d548e.
//
// Solidity: event EthDonated(uint256 indexed roundNum, address indexed donorAddress, uint256 amount)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterEthDonated(opts *bind.FilterOpts, roundNum []*big.Int, donorAddress []common.Address) (*CosmicSignatureGameV3EthDonatedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "EthDonated", roundNumRule, donorAddressRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3EthDonatedIterator{contract: _CosmicSignatureGameV3.contract, event: "EthDonated", logs: logs, sub: sub}, nil
}

// WatchEthDonated is a free log subscription operation binding the contract event 0xe32cacf203d00685e2b4d8b0a90e7cd8f3f8a208fdf116f4bb36abe08b7d548e.
//
// Solidity: event EthDonated(uint256 indexed roundNum, address indexed donorAddress, uint256 amount)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchEthDonated(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3EthDonated, roundNum []*big.Int, donorAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var donorAddressRule []interface{}
	for _, donorAddressItem := range donorAddress {
		donorAddressRule = append(donorAddressRule, donorAddressItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "EthDonated", roundNumRule, donorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3EthDonated)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "EthDonated", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseEthDonated(log types.Log) (*CosmicSignatureGameV3EthDonated, error) {
	event := new(CosmicSignatureGameV3EthDonated)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "EthDonated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3EthDonatedWithInfoIterator is returned from FilterEthDonatedWithInfo and is used to iterate over the raw logs and unpacked data for EthDonatedWithInfo events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3EthDonatedWithInfoIterator struct {
	Event *CosmicSignatureGameV3EthDonatedWithInfo // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3EthDonatedWithInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3EthDonatedWithInfo)
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
		it.Event = new(CosmicSignatureGameV3EthDonatedWithInfo)
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
func (it *CosmicSignatureGameV3EthDonatedWithInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3EthDonatedWithInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3EthDonatedWithInfo represents a EthDonatedWithInfo event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3EthDonatedWithInfo struct {
	RoundNum                       *big.Int
	DonorAddress                   common.Address
	Amount                         *big.Int
	EthDonationWithInfoRecordIndex *big.Int
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterEthDonatedWithInfo is a free log retrieval operation binding the contract event 0xa08049565b10d44a06dca9bf05685b39bc370352043c5a003e8d35d45ebdc53f.
//
// Solidity: event EthDonatedWithInfo(uint256 indexed roundNum, address indexed donorAddress, uint256 amount, uint256 indexed ethDonationWithInfoRecordIndex)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterEthDonatedWithInfo(opts *bind.FilterOpts, roundNum []*big.Int, donorAddress []common.Address, ethDonationWithInfoRecordIndex []*big.Int) (*CosmicSignatureGameV3EthDonatedWithInfoIterator, error) {

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

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "EthDonatedWithInfo", roundNumRule, donorAddressRule, ethDonationWithInfoRecordIndexRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3EthDonatedWithInfoIterator{contract: _CosmicSignatureGameV3.contract, event: "EthDonatedWithInfo", logs: logs, sub: sub}, nil
}

// WatchEthDonatedWithInfo is a free log subscription operation binding the contract event 0xa08049565b10d44a06dca9bf05685b39bc370352043c5a003e8d35d45ebdc53f.
//
// Solidity: event EthDonatedWithInfo(uint256 indexed roundNum, address indexed donorAddress, uint256 amount, uint256 indexed ethDonationWithInfoRecordIndex)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchEthDonatedWithInfo(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3EthDonatedWithInfo, roundNum []*big.Int, donorAddress []common.Address, ethDonationWithInfoRecordIndex []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "EthDonatedWithInfo", roundNumRule, donorAddressRule, ethDonationWithInfoRecordIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3EthDonatedWithInfo)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "EthDonatedWithInfo", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseEthDonatedWithInfo(log types.Log) (*CosmicSignatureGameV3EthDonatedWithInfo, error) {
	event := new(CosmicSignatureGameV3EthDonatedWithInfo)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "EthDonatedWithInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3EthDutchAuctionDurationDivisorChangedIterator is returned from FilterEthDutchAuctionDurationDivisorChanged and is used to iterate over the raw logs and unpacked data for EthDutchAuctionDurationDivisorChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3EthDutchAuctionDurationDivisorChangedIterator struct {
	Event *CosmicSignatureGameV3EthDutchAuctionDurationDivisorChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3EthDutchAuctionDurationDivisorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3EthDutchAuctionDurationDivisorChanged)
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
		it.Event = new(CosmicSignatureGameV3EthDutchAuctionDurationDivisorChanged)
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
func (it *CosmicSignatureGameV3EthDutchAuctionDurationDivisorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3EthDutchAuctionDurationDivisorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3EthDutchAuctionDurationDivisorChanged represents a EthDutchAuctionDurationDivisorChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3EthDutchAuctionDurationDivisorChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDutchAuctionDurationDivisorChanged is a free log retrieval operation binding the contract event 0xfdf6043c89a5f304289122dcc9f8bd78bb111b5d4f409e2fc2e6c141a1110b79.
//
// Solidity: event EthDutchAuctionDurationDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterEthDutchAuctionDurationDivisorChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3EthDutchAuctionDurationDivisorChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "EthDutchAuctionDurationDivisorChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3EthDutchAuctionDurationDivisorChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "EthDutchAuctionDurationDivisorChanged", logs: logs, sub: sub}, nil
}

// WatchEthDutchAuctionDurationDivisorChanged is a free log subscription operation binding the contract event 0xfdf6043c89a5f304289122dcc9f8bd78bb111b5d4f409e2fc2e6c141a1110b79.
//
// Solidity: event EthDutchAuctionDurationDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchEthDutchAuctionDurationDivisorChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3EthDutchAuctionDurationDivisorChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "EthDutchAuctionDurationDivisorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3EthDutchAuctionDurationDivisorChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "EthDutchAuctionDurationDivisorChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseEthDutchAuctionDurationDivisorChanged(log types.Log) (*CosmicSignatureGameV3EthDutchAuctionDurationDivisorChanged, error) {
	event := new(CosmicSignatureGameV3EthDutchAuctionDurationDivisorChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "EthDutchAuctionDurationDivisorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3EthDutchAuctionEndingBidPriceDivisorChangedIterator is returned from FilterEthDutchAuctionEndingBidPriceDivisorChanged and is used to iterate over the raw logs and unpacked data for EthDutchAuctionEndingBidPriceDivisorChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3EthDutchAuctionEndingBidPriceDivisorChangedIterator struct {
	Event *CosmicSignatureGameV3EthDutchAuctionEndingBidPriceDivisorChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3EthDutchAuctionEndingBidPriceDivisorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3EthDutchAuctionEndingBidPriceDivisorChanged)
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
		it.Event = new(CosmicSignatureGameV3EthDutchAuctionEndingBidPriceDivisorChanged)
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
func (it *CosmicSignatureGameV3EthDutchAuctionEndingBidPriceDivisorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3EthDutchAuctionEndingBidPriceDivisorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3EthDutchAuctionEndingBidPriceDivisorChanged represents a EthDutchAuctionEndingBidPriceDivisorChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3EthDutchAuctionEndingBidPriceDivisorChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEthDutchAuctionEndingBidPriceDivisorChanged is a free log retrieval operation binding the contract event 0xb6f6af60099e44041a78b3561ed029b98bf03fdb0efbbb2eb15e1f3d7d923037.
//
// Solidity: event EthDutchAuctionEndingBidPriceDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterEthDutchAuctionEndingBidPriceDivisorChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3EthDutchAuctionEndingBidPriceDivisorChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "EthDutchAuctionEndingBidPriceDivisorChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3EthDutchAuctionEndingBidPriceDivisorChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "EthDutchAuctionEndingBidPriceDivisorChanged", logs: logs, sub: sub}, nil
}

// WatchEthDutchAuctionEndingBidPriceDivisorChanged is a free log subscription operation binding the contract event 0xb6f6af60099e44041a78b3561ed029b98bf03fdb0efbbb2eb15e1f3d7d923037.
//
// Solidity: event EthDutchAuctionEndingBidPriceDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchEthDutchAuctionEndingBidPriceDivisorChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3EthDutchAuctionEndingBidPriceDivisorChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "EthDutchAuctionEndingBidPriceDivisorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3EthDutchAuctionEndingBidPriceDivisorChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "EthDutchAuctionEndingBidPriceDivisorChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseEthDutchAuctionEndingBidPriceDivisorChanged(log types.Log) (*CosmicSignatureGameV3EthDutchAuctionEndingBidPriceDivisorChanged, error) {
	event := new(CosmicSignatureGameV3EthDutchAuctionEndingBidPriceDivisorChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "EthDutchAuctionEndingBidPriceDivisorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3FirstBidPlacedInRoundIterator is returned from FilterFirstBidPlacedInRound and is used to iterate over the raw logs and unpacked data for FirstBidPlacedInRound events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3FirstBidPlacedInRoundIterator struct {
	Event *CosmicSignatureGameV3FirstBidPlacedInRound // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3FirstBidPlacedInRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3FirstBidPlacedInRound)
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
		it.Event = new(CosmicSignatureGameV3FirstBidPlacedInRound)
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
func (it *CosmicSignatureGameV3FirstBidPlacedInRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3FirstBidPlacedInRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3FirstBidPlacedInRound represents a FirstBidPlacedInRound event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3FirstBidPlacedInRound struct {
	RoundNum       *big.Int
	BlockTimeStamp *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFirstBidPlacedInRound is a free log retrieval operation binding the contract event 0x028a52641badd593b7f30072734c0b97e449213f55b5c3663756427340accd3c.
//
// Solidity: event FirstBidPlacedInRound(uint256 indexed roundNum, uint256 blockTimeStamp)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterFirstBidPlacedInRound(opts *bind.FilterOpts, roundNum []*big.Int) (*CosmicSignatureGameV3FirstBidPlacedInRoundIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "FirstBidPlacedInRound", roundNumRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3FirstBidPlacedInRoundIterator{contract: _CosmicSignatureGameV3.contract, event: "FirstBidPlacedInRound", logs: logs, sub: sub}, nil
}

// WatchFirstBidPlacedInRound is a free log subscription operation binding the contract event 0x028a52641badd593b7f30072734c0b97e449213f55b5c3663756427340accd3c.
//
// Solidity: event FirstBidPlacedInRound(uint256 indexed roundNum, uint256 blockTimeStamp)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchFirstBidPlacedInRound(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3FirstBidPlacedInRound, roundNum []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "FirstBidPlacedInRound", roundNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3FirstBidPlacedInRound)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "FirstBidPlacedInRound", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseFirstBidPlacedInRound(log types.Log) (*CosmicSignatureGameV3FirstBidPlacedInRound, error) {
	event := new(CosmicSignatureGameV3FirstBidPlacedInRound)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "FirstBidPlacedInRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3FundTransferFailedIterator is returned from FilterFundTransferFailed and is used to iterate over the raw logs and unpacked data for FundTransferFailed events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3FundTransferFailedIterator struct {
	Event *CosmicSignatureGameV3FundTransferFailed // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3FundTransferFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3FundTransferFailed)
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
		it.Event = new(CosmicSignatureGameV3FundTransferFailed)
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
func (it *CosmicSignatureGameV3FundTransferFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3FundTransferFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3FundTransferFailed represents a FundTransferFailed event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3FundTransferFailed struct {
	ErrStr             string
	DestinationAddress common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFundTransferFailed is a free log retrieval operation binding the contract event 0x154fb6c686c977314af35b730a16b83facda5265b2abec7237c730f63e42215a.
//
// Solidity: event FundTransferFailed(string errStr, address indexed destinationAddress, uint256 amount)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterFundTransferFailed(opts *bind.FilterOpts, destinationAddress []common.Address) (*CosmicSignatureGameV3FundTransferFailedIterator, error) {

	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "FundTransferFailed", destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3FundTransferFailedIterator{contract: _CosmicSignatureGameV3.contract, event: "FundTransferFailed", logs: logs, sub: sub}, nil
}

// WatchFundTransferFailed is a free log subscription operation binding the contract event 0x154fb6c686c977314af35b730a16b83facda5265b2abec7237c730f63e42215a.
//
// Solidity: event FundTransferFailed(string errStr, address indexed destinationAddress, uint256 amount)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchFundTransferFailed(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3FundTransferFailed, destinationAddress []common.Address) (event.Subscription, error) {

	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "FundTransferFailed", destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3FundTransferFailed)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "FundTransferFailed", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseFundTransferFailed(log types.Log) (*CosmicSignatureGameV3FundTransferFailed, error) {
	event := new(CosmicSignatureGameV3FundTransferFailed)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "FundTransferFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3FundsTransferredToCharityIterator is returned from FilterFundsTransferredToCharity and is used to iterate over the raw logs and unpacked data for FundsTransferredToCharity events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3FundsTransferredToCharityIterator struct {
	Event *CosmicSignatureGameV3FundsTransferredToCharity // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3FundsTransferredToCharityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3FundsTransferredToCharity)
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
		it.Event = new(CosmicSignatureGameV3FundsTransferredToCharity)
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
func (it *CosmicSignatureGameV3FundsTransferredToCharityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3FundsTransferredToCharityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3FundsTransferredToCharity represents a FundsTransferredToCharity event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3FundsTransferredToCharity struct {
	CharityAddress common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFundsTransferredToCharity is a free log retrieval operation binding the contract event 0x1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d.
//
// Solidity: event FundsTransferredToCharity(address indexed charityAddress, uint256 amount)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterFundsTransferredToCharity(opts *bind.FilterOpts, charityAddress []common.Address) (*CosmicSignatureGameV3FundsTransferredToCharityIterator, error) {

	var charityAddressRule []interface{}
	for _, charityAddressItem := range charityAddress {
		charityAddressRule = append(charityAddressRule, charityAddressItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "FundsTransferredToCharity", charityAddressRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3FundsTransferredToCharityIterator{contract: _CosmicSignatureGameV3.contract, event: "FundsTransferredToCharity", logs: logs, sub: sub}, nil
}

// WatchFundsTransferredToCharity is a free log subscription operation binding the contract event 0x1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d.
//
// Solidity: event FundsTransferredToCharity(address indexed charityAddress, uint256 amount)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchFundsTransferredToCharity(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3FundsTransferredToCharity, charityAddress []common.Address) (event.Subscription, error) {

	var charityAddressRule []interface{}
	for _, charityAddressItem := range charityAddress {
		charityAddressRule = append(charityAddressRule, charityAddressItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "FundsTransferredToCharity", charityAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3FundsTransferredToCharity)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "FundsTransferredToCharity", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseFundsTransferredToCharity(log types.Log) (*CosmicSignatureGameV3FundsTransferredToCharity, error) {
	event := new(CosmicSignatureGameV3FundsTransferredToCharity)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "FundsTransferredToCharity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3InitialDurationUntilMainPrizeDivisorChangedIterator is returned from FilterInitialDurationUntilMainPrizeDivisorChanged and is used to iterate over the raw logs and unpacked data for InitialDurationUntilMainPrizeDivisorChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3InitialDurationUntilMainPrizeDivisorChangedIterator struct {
	Event *CosmicSignatureGameV3InitialDurationUntilMainPrizeDivisorChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3InitialDurationUntilMainPrizeDivisorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3InitialDurationUntilMainPrizeDivisorChanged)
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
		it.Event = new(CosmicSignatureGameV3InitialDurationUntilMainPrizeDivisorChanged)
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
func (it *CosmicSignatureGameV3InitialDurationUntilMainPrizeDivisorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3InitialDurationUntilMainPrizeDivisorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3InitialDurationUntilMainPrizeDivisorChanged represents a InitialDurationUntilMainPrizeDivisorChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3InitialDurationUntilMainPrizeDivisorChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInitialDurationUntilMainPrizeDivisorChanged is a free log retrieval operation binding the contract event 0xb5edd1f338b34c8f5dd3b1c5cc12f05653c495713c282bf588d34cf14fad0f89.
//
// Solidity: event InitialDurationUntilMainPrizeDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterInitialDurationUntilMainPrizeDivisorChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3InitialDurationUntilMainPrizeDivisorChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "InitialDurationUntilMainPrizeDivisorChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3InitialDurationUntilMainPrizeDivisorChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "InitialDurationUntilMainPrizeDivisorChanged", logs: logs, sub: sub}, nil
}

// WatchInitialDurationUntilMainPrizeDivisorChanged is a free log subscription operation binding the contract event 0xb5edd1f338b34c8f5dd3b1c5cc12f05653c495713c282bf588d34cf14fad0f89.
//
// Solidity: event InitialDurationUntilMainPrizeDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchInitialDurationUntilMainPrizeDivisorChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3InitialDurationUntilMainPrizeDivisorChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "InitialDurationUntilMainPrizeDivisorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3InitialDurationUntilMainPrizeDivisorChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "InitialDurationUntilMainPrizeDivisorChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseInitialDurationUntilMainPrizeDivisorChanged(log types.Log) (*CosmicSignatureGameV3InitialDurationUntilMainPrizeDivisorChanged, error) {
	event := new(CosmicSignatureGameV3InitialDurationUntilMainPrizeDivisorChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "InitialDurationUntilMainPrizeDivisorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3InitializedIterator struct {
	Event *CosmicSignatureGameV3Initialized // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3Initialized)
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
		it.Event = new(CosmicSignatureGameV3Initialized)
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
func (it *CosmicSignatureGameV3InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3Initialized represents a Initialized event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterInitialized(opts *bind.FilterOpts) (*CosmicSignatureGameV3InitializedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3InitializedIterator{contract: _CosmicSignatureGameV3.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3Initialized) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3Initialized)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseInitialized(log types.Log) (*CosmicSignatureGameV3Initialized, error) {
	event := new(CosmicSignatureGameV3Initialized)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3LastCstBidderPrizePaidIterator is returned from FilterLastCstBidderPrizePaid and is used to iterate over the raw logs and unpacked data for LastCstBidderPrizePaid events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3LastCstBidderPrizePaidIterator struct {
	Event *CosmicSignatureGameV3LastCstBidderPrizePaid // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3LastCstBidderPrizePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3LastCstBidderPrizePaid)
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
		it.Event = new(CosmicSignatureGameV3LastCstBidderPrizePaid)
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
func (it *CosmicSignatureGameV3LastCstBidderPrizePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3LastCstBidderPrizePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3LastCstBidderPrizePaid represents a LastCstBidderPrizePaid event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3LastCstBidderPrizePaid struct {
	RoundNum                  *big.Int
	LastCstBidderAddress      common.Address
	CstPrizeAmount            *big.Int
	PrizeCosmicSignatureNftId *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterLastCstBidderPrizePaid is a free log retrieval operation binding the contract event 0x3901b6430c99dc290ee88ff84c4de6091ad7eac335b58e92ef5cbb0793abf4f6.
//
// Solidity: event LastCstBidderPrizePaid(uint256 indexed roundNum, address indexed lastCstBidderAddress, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterLastCstBidderPrizePaid(opts *bind.FilterOpts, roundNum []*big.Int, lastCstBidderAddress []common.Address, prizeCosmicSignatureNftId []*big.Int) (*CosmicSignatureGameV3LastCstBidderPrizePaidIterator, error) {

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

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "LastCstBidderPrizePaid", roundNumRule, lastCstBidderAddressRule, prizeCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3LastCstBidderPrizePaidIterator{contract: _CosmicSignatureGameV3.contract, event: "LastCstBidderPrizePaid", logs: logs, sub: sub}, nil
}

// WatchLastCstBidderPrizePaid is a free log subscription operation binding the contract event 0x3901b6430c99dc290ee88ff84c4de6091ad7eac335b58e92ef5cbb0793abf4f6.
//
// Solidity: event LastCstBidderPrizePaid(uint256 indexed roundNum, address indexed lastCstBidderAddress, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchLastCstBidderPrizePaid(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3LastCstBidderPrizePaid, roundNum []*big.Int, lastCstBidderAddress []common.Address, prizeCosmicSignatureNftId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "LastCstBidderPrizePaid", roundNumRule, lastCstBidderAddressRule, prizeCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3LastCstBidderPrizePaid)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "LastCstBidderPrizePaid", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseLastCstBidderPrizePaid(log types.Log) (*CosmicSignatureGameV3LastCstBidderPrizePaid, error) {
	event := new(CosmicSignatureGameV3LastCstBidderPrizePaid)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "LastCstBidderPrizePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3MainEthPrizeAmountPercentageChangedIterator is returned from FilterMainEthPrizeAmountPercentageChanged and is used to iterate over the raw logs and unpacked data for MainEthPrizeAmountPercentageChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3MainEthPrizeAmountPercentageChangedIterator struct {
	Event *CosmicSignatureGameV3MainEthPrizeAmountPercentageChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3MainEthPrizeAmountPercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3MainEthPrizeAmountPercentageChanged)
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
		it.Event = new(CosmicSignatureGameV3MainEthPrizeAmountPercentageChanged)
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
func (it *CosmicSignatureGameV3MainEthPrizeAmountPercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3MainEthPrizeAmountPercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3MainEthPrizeAmountPercentageChanged represents a MainEthPrizeAmountPercentageChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3MainEthPrizeAmountPercentageChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMainEthPrizeAmountPercentageChanged is a free log retrieval operation binding the contract event 0xb5a05ec7911dd5450a7fa4ae54d50d9d83af6e256db8fc76c82edd7b659cf8bc.
//
// Solidity: event MainEthPrizeAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterMainEthPrizeAmountPercentageChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3MainEthPrizeAmountPercentageChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "MainEthPrizeAmountPercentageChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3MainEthPrizeAmountPercentageChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "MainEthPrizeAmountPercentageChanged", logs: logs, sub: sub}, nil
}

// WatchMainEthPrizeAmountPercentageChanged is a free log subscription operation binding the contract event 0xb5a05ec7911dd5450a7fa4ae54d50d9d83af6e256db8fc76c82edd7b659cf8bc.
//
// Solidity: event MainEthPrizeAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchMainEthPrizeAmountPercentageChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3MainEthPrizeAmountPercentageChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "MainEthPrizeAmountPercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3MainEthPrizeAmountPercentageChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "MainEthPrizeAmountPercentageChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseMainEthPrizeAmountPercentageChanged(log types.Log) (*CosmicSignatureGameV3MainEthPrizeAmountPercentageChanged, error) {
	event := new(CosmicSignatureGameV3MainEthPrizeAmountPercentageChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "MainEthPrizeAmountPercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3MainPrizeClaimedIterator is returned from FilterMainPrizeClaimed and is used to iterate over the raw logs and unpacked data for MainPrizeClaimed events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3MainPrizeClaimedIterator struct {
	Event *CosmicSignatureGameV3MainPrizeClaimed // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3MainPrizeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3MainPrizeClaimed)
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
		it.Event = new(CosmicSignatureGameV3MainPrizeClaimed)
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
func (it *CosmicSignatureGameV3MainPrizeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3MainPrizeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3MainPrizeClaimed represents a MainPrizeClaimed event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3MainPrizeClaimed struct {
	RoundNum                             *big.Int
	BeneficiaryAddress                   common.Address
	EthPrizeAmount                       *big.Int
	CstPrizeAmount                       *big.Int
	PrizeFirstCosmicSignatureNftId       *big.Int
	PrizeNumCosmicSignatureNfts          *big.Int
	TimeoutTimeToWithdrawSecondaryPrizes *big.Int
	Raw                                  types.Log // Blockchain specific contextual infos
}

// FilterMainPrizeClaimed is a free log retrieval operation binding the contract event 0x9314e78538382a9fc6cd54cee107c8a7721d172c43b2e48c2e7a51f200fc3788.
//
// Solidity: event MainPrizeClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, uint256 ethPrizeAmount, uint256 cstPrizeAmount, uint256 indexed prizeFirstCosmicSignatureNftId, uint256 prizeNumCosmicSignatureNfts, uint256 timeoutTimeToWithdrawSecondaryPrizes)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterMainPrizeClaimed(opts *bind.FilterOpts, roundNum []*big.Int, beneficiaryAddress []common.Address, prizeFirstCosmicSignatureNftId []*big.Int) (*CosmicSignatureGameV3MainPrizeClaimedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	var prizeFirstCosmicSignatureNftIdRule []interface{}
	for _, prizeFirstCosmicSignatureNftIdItem := range prizeFirstCosmicSignatureNftId {
		prizeFirstCosmicSignatureNftIdRule = append(prizeFirstCosmicSignatureNftIdRule, prizeFirstCosmicSignatureNftIdItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "MainPrizeClaimed", roundNumRule, beneficiaryAddressRule, prizeFirstCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3MainPrizeClaimedIterator{contract: _CosmicSignatureGameV3.contract, event: "MainPrizeClaimed", logs: logs, sub: sub}, nil
}

// WatchMainPrizeClaimed is a free log subscription operation binding the contract event 0x9314e78538382a9fc6cd54cee107c8a7721d172c43b2e48c2e7a51f200fc3788.
//
// Solidity: event MainPrizeClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, uint256 ethPrizeAmount, uint256 cstPrizeAmount, uint256 indexed prizeFirstCosmicSignatureNftId, uint256 prizeNumCosmicSignatureNfts, uint256 timeoutTimeToWithdrawSecondaryPrizes)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchMainPrizeClaimed(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3MainPrizeClaimed, roundNum []*big.Int, beneficiaryAddress []common.Address, prizeFirstCosmicSignatureNftId []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}
	var beneficiaryAddressRule []interface{}
	for _, beneficiaryAddressItem := range beneficiaryAddress {
		beneficiaryAddressRule = append(beneficiaryAddressRule, beneficiaryAddressItem)
	}

	var prizeFirstCosmicSignatureNftIdRule []interface{}
	for _, prizeFirstCosmicSignatureNftIdItem := range prizeFirstCosmicSignatureNftId {
		prizeFirstCosmicSignatureNftIdRule = append(prizeFirstCosmicSignatureNftIdRule, prizeFirstCosmicSignatureNftIdItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "MainPrizeClaimed", roundNumRule, beneficiaryAddressRule, prizeFirstCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3MainPrizeClaimed)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "MainPrizeClaimed", log); err != nil {
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

// ParseMainPrizeClaimed is a log parse operation binding the contract event 0x9314e78538382a9fc6cd54cee107c8a7721d172c43b2e48c2e7a51f200fc3788.
//
// Solidity: event MainPrizeClaimed(uint256 indexed roundNum, address indexed beneficiaryAddress, uint256 ethPrizeAmount, uint256 cstPrizeAmount, uint256 indexed prizeFirstCosmicSignatureNftId, uint256 prizeNumCosmicSignatureNfts, uint256 timeoutTimeToWithdrawSecondaryPrizes)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseMainPrizeClaimed(log types.Log) (*CosmicSignatureGameV3MainPrizeClaimed, error) {
	event := new(CosmicSignatureGameV3MainPrizeClaimed)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "MainPrizeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3MainPrizeNumCosmicSignatureNftsChangedIterator is returned from FilterMainPrizeNumCosmicSignatureNftsChanged and is used to iterate over the raw logs and unpacked data for MainPrizeNumCosmicSignatureNftsChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3MainPrizeNumCosmicSignatureNftsChangedIterator struct {
	Event *CosmicSignatureGameV3MainPrizeNumCosmicSignatureNftsChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3MainPrizeNumCosmicSignatureNftsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3MainPrizeNumCosmicSignatureNftsChanged)
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
		it.Event = new(CosmicSignatureGameV3MainPrizeNumCosmicSignatureNftsChanged)
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
func (it *CosmicSignatureGameV3MainPrizeNumCosmicSignatureNftsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3MainPrizeNumCosmicSignatureNftsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3MainPrizeNumCosmicSignatureNftsChanged represents a MainPrizeNumCosmicSignatureNftsChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3MainPrizeNumCosmicSignatureNftsChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMainPrizeNumCosmicSignatureNftsChanged is a free log retrieval operation binding the contract event 0x616bfcaa6490f55f6e57a4deedac1db04d0d6826deb84fad86cc43439bcf3564.
//
// Solidity: event MainPrizeNumCosmicSignatureNftsChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterMainPrizeNumCosmicSignatureNftsChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3MainPrizeNumCosmicSignatureNftsChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "MainPrizeNumCosmicSignatureNftsChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3MainPrizeNumCosmicSignatureNftsChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "MainPrizeNumCosmicSignatureNftsChanged", logs: logs, sub: sub}, nil
}

// WatchMainPrizeNumCosmicSignatureNftsChanged is a free log subscription operation binding the contract event 0x616bfcaa6490f55f6e57a4deedac1db04d0d6826deb84fad86cc43439bcf3564.
//
// Solidity: event MainPrizeNumCosmicSignatureNftsChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchMainPrizeNumCosmicSignatureNftsChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3MainPrizeNumCosmicSignatureNftsChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "MainPrizeNumCosmicSignatureNftsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3MainPrizeNumCosmicSignatureNftsChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "MainPrizeNumCosmicSignatureNftsChanged", log); err != nil {
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

// ParseMainPrizeNumCosmicSignatureNftsChanged is a log parse operation binding the contract event 0x616bfcaa6490f55f6e57a4deedac1db04d0d6826deb84fad86cc43439bcf3564.
//
// Solidity: event MainPrizeNumCosmicSignatureNftsChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseMainPrizeNumCosmicSignatureNftsChanged(log types.Log) (*CosmicSignatureGameV3MainPrizeNumCosmicSignatureNftsChanged, error) {
	event := new(CosmicSignatureGameV3MainPrizeNumCosmicSignatureNftsChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "MainPrizeNumCosmicSignatureNftsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3MainPrizeTimeIncrementInMicroSecondsChangedIterator is returned from FilterMainPrizeTimeIncrementInMicroSecondsChanged and is used to iterate over the raw logs and unpacked data for MainPrizeTimeIncrementInMicroSecondsChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3MainPrizeTimeIncrementInMicroSecondsChangedIterator struct {
	Event *CosmicSignatureGameV3MainPrizeTimeIncrementInMicroSecondsChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3MainPrizeTimeIncrementInMicroSecondsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3MainPrizeTimeIncrementInMicroSecondsChanged)
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
		it.Event = new(CosmicSignatureGameV3MainPrizeTimeIncrementInMicroSecondsChanged)
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
func (it *CosmicSignatureGameV3MainPrizeTimeIncrementInMicroSecondsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3MainPrizeTimeIncrementInMicroSecondsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3MainPrizeTimeIncrementInMicroSecondsChanged represents a MainPrizeTimeIncrementInMicroSecondsChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3MainPrizeTimeIncrementInMicroSecondsChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMainPrizeTimeIncrementInMicroSecondsChanged is a free log retrieval operation binding the contract event 0x07417920574ce0bdfe987af0575c8793cc73a29d7830760ad459d0e569b5b79b.
//
// Solidity: event MainPrizeTimeIncrementInMicroSecondsChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterMainPrizeTimeIncrementInMicroSecondsChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3MainPrizeTimeIncrementInMicroSecondsChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "MainPrizeTimeIncrementInMicroSecondsChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3MainPrizeTimeIncrementInMicroSecondsChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "MainPrizeTimeIncrementInMicroSecondsChanged", logs: logs, sub: sub}, nil
}

// WatchMainPrizeTimeIncrementInMicroSecondsChanged is a free log subscription operation binding the contract event 0x07417920574ce0bdfe987af0575c8793cc73a29d7830760ad459d0e569b5b79b.
//
// Solidity: event MainPrizeTimeIncrementInMicroSecondsChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchMainPrizeTimeIncrementInMicroSecondsChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3MainPrizeTimeIncrementInMicroSecondsChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "MainPrizeTimeIncrementInMicroSecondsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3MainPrizeTimeIncrementInMicroSecondsChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "MainPrizeTimeIncrementInMicroSecondsChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseMainPrizeTimeIncrementInMicroSecondsChanged(log types.Log) (*CosmicSignatureGameV3MainPrizeTimeIncrementInMicroSecondsChanged, error) {
	event := new(CosmicSignatureGameV3MainPrizeTimeIncrementInMicroSecondsChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "MainPrizeTimeIncrementInMicroSecondsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3MainPrizeTimeIncrementIncreaseDivisorChangedIterator is returned from FilterMainPrizeTimeIncrementIncreaseDivisorChanged and is used to iterate over the raw logs and unpacked data for MainPrizeTimeIncrementIncreaseDivisorChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3MainPrizeTimeIncrementIncreaseDivisorChangedIterator struct {
	Event *CosmicSignatureGameV3MainPrizeTimeIncrementIncreaseDivisorChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3MainPrizeTimeIncrementIncreaseDivisorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3MainPrizeTimeIncrementIncreaseDivisorChanged)
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
		it.Event = new(CosmicSignatureGameV3MainPrizeTimeIncrementIncreaseDivisorChanged)
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
func (it *CosmicSignatureGameV3MainPrizeTimeIncrementIncreaseDivisorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3MainPrizeTimeIncrementIncreaseDivisorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3MainPrizeTimeIncrementIncreaseDivisorChanged represents a MainPrizeTimeIncrementIncreaseDivisorChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3MainPrizeTimeIncrementIncreaseDivisorChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMainPrizeTimeIncrementIncreaseDivisorChanged is a free log retrieval operation binding the contract event 0x4636d3e567b27988879babd22e50f49104ca65647933fc1623ff3d4d807438d2.
//
// Solidity: event MainPrizeTimeIncrementIncreaseDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterMainPrizeTimeIncrementIncreaseDivisorChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3MainPrizeTimeIncrementIncreaseDivisorChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "MainPrizeTimeIncrementIncreaseDivisorChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3MainPrizeTimeIncrementIncreaseDivisorChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "MainPrizeTimeIncrementIncreaseDivisorChanged", logs: logs, sub: sub}, nil
}

// WatchMainPrizeTimeIncrementIncreaseDivisorChanged is a free log subscription operation binding the contract event 0x4636d3e567b27988879babd22e50f49104ca65647933fc1623ff3d4d807438d2.
//
// Solidity: event MainPrizeTimeIncrementIncreaseDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchMainPrizeTimeIncrementIncreaseDivisorChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3MainPrizeTimeIncrementIncreaseDivisorChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "MainPrizeTimeIncrementIncreaseDivisorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3MainPrizeTimeIncrementIncreaseDivisorChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "MainPrizeTimeIncrementIncreaseDivisorChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseMainPrizeTimeIncrementIncreaseDivisorChanged(log types.Log) (*CosmicSignatureGameV3MainPrizeTimeIncrementIncreaseDivisorChanged, error) {
	event := new(CosmicSignatureGameV3MainPrizeTimeIncrementIncreaseDivisorChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "MainPrizeTimeIncrementIncreaseDivisorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3MarketingWalletAddressChangedIterator is returned from FilterMarketingWalletAddressChanged and is used to iterate over the raw logs and unpacked data for MarketingWalletAddressChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3MarketingWalletAddressChangedIterator struct {
	Event *CosmicSignatureGameV3MarketingWalletAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3MarketingWalletAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3MarketingWalletAddressChanged)
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
		it.Event = new(CosmicSignatureGameV3MarketingWalletAddressChanged)
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
func (it *CosmicSignatureGameV3MarketingWalletAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3MarketingWalletAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3MarketingWalletAddressChanged represents a MarketingWalletAddressChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3MarketingWalletAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketingWalletAddressChanged is a free log retrieval operation binding the contract event 0x4d03942c29c20d1bccfe551e9d148c917c5a44fb558a4fc60270d8f76fb75f54.
//
// Solidity: event MarketingWalletAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterMarketingWalletAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*CosmicSignatureGameV3MarketingWalletAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "MarketingWalletAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3MarketingWalletAddressChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "MarketingWalletAddressChanged", logs: logs, sub: sub}, nil
}

// WatchMarketingWalletAddressChanged is a free log subscription operation binding the contract event 0x4d03942c29c20d1bccfe551e9d148c917c5a44fb558a4fc60270d8f76fb75f54.
//
// Solidity: event MarketingWalletAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchMarketingWalletAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3MarketingWalletAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "MarketingWalletAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3MarketingWalletAddressChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "MarketingWalletAddressChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseMarketingWalletAddressChanged(log types.Log) (*CosmicSignatureGameV3MarketingWalletAddressChanged, error) {
	event := new(CosmicSignatureGameV3MarketingWalletAddressChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "MarketingWalletAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3MarketingWalletCstContributionAmountChangedIterator is returned from FilterMarketingWalletCstContributionAmountChanged and is used to iterate over the raw logs and unpacked data for MarketingWalletCstContributionAmountChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3MarketingWalletCstContributionAmountChangedIterator struct {
	Event *CosmicSignatureGameV3MarketingWalletCstContributionAmountChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3MarketingWalletCstContributionAmountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3MarketingWalletCstContributionAmountChanged)
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
		it.Event = new(CosmicSignatureGameV3MarketingWalletCstContributionAmountChanged)
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
func (it *CosmicSignatureGameV3MarketingWalletCstContributionAmountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3MarketingWalletCstContributionAmountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3MarketingWalletCstContributionAmountChanged represents a MarketingWalletCstContributionAmountChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3MarketingWalletCstContributionAmountChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketingWalletCstContributionAmountChanged is a free log retrieval operation binding the contract event 0x2652e6657dd1ed89d7bdcb70f8827cc8449ac4536ebf604dbb2465cdad264020.
//
// Solidity: event MarketingWalletCstContributionAmountChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterMarketingWalletCstContributionAmountChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3MarketingWalletCstContributionAmountChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "MarketingWalletCstContributionAmountChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3MarketingWalletCstContributionAmountChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "MarketingWalletCstContributionAmountChanged", logs: logs, sub: sub}, nil
}

// WatchMarketingWalletCstContributionAmountChanged is a free log subscription operation binding the contract event 0x2652e6657dd1ed89d7bdcb70f8827cc8449ac4536ebf604dbb2465cdad264020.
//
// Solidity: event MarketingWalletCstContributionAmountChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchMarketingWalletCstContributionAmountChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3MarketingWalletCstContributionAmountChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "MarketingWalletCstContributionAmountChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3MarketingWalletCstContributionAmountChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "MarketingWalletCstContributionAmountChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseMarketingWalletCstContributionAmountChanged(log types.Log) (*CosmicSignatureGameV3MarketingWalletCstContributionAmountChanged, error) {
	event := new(CosmicSignatureGameV3MarketingWalletCstContributionAmountChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "MarketingWalletCstContributionAmountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForBiddersChangedIterator is returned from FilterNumRaffleCosmicSignatureNftsForBiddersChanged and is used to iterate over the raw logs and unpacked data for NumRaffleCosmicSignatureNftsForBiddersChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForBiddersChangedIterator struct {
	Event *CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForBiddersChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForBiddersChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForBiddersChanged)
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
		it.Event = new(CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForBiddersChanged)
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
func (it *CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForBiddersChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForBiddersChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForBiddersChanged represents a NumRaffleCosmicSignatureNftsForBiddersChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForBiddersChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNumRaffleCosmicSignatureNftsForBiddersChanged is a free log retrieval operation binding the contract event 0x85d8bf21006916836edd67a5afeed2e891cf876a6c9cd9babf7f42f4b007c24f.
//
// Solidity: event NumRaffleCosmicSignatureNftsForBiddersChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterNumRaffleCosmicSignatureNftsForBiddersChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForBiddersChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "NumRaffleCosmicSignatureNftsForBiddersChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForBiddersChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "NumRaffleCosmicSignatureNftsForBiddersChanged", logs: logs, sub: sub}, nil
}

// WatchNumRaffleCosmicSignatureNftsForBiddersChanged is a free log subscription operation binding the contract event 0x85d8bf21006916836edd67a5afeed2e891cf876a6c9cd9babf7f42f4b007c24f.
//
// Solidity: event NumRaffleCosmicSignatureNftsForBiddersChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchNumRaffleCosmicSignatureNftsForBiddersChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForBiddersChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "NumRaffleCosmicSignatureNftsForBiddersChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForBiddersChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "NumRaffleCosmicSignatureNftsForBiddersChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseNumRaffleCosmicSignatureNftsForBiddersChanged(log types.Log) (*CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForBiddersChanged, error) {
	event := new(CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForBiddersChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "NumRaffleCosmicSignatureNftsForBiddersChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChangedIterator is returned from FilterNumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged and is used to iterate over the raw logs and unpacked data for NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChangedIterator struct {
	Event *CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged)
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
		it.Event = new(CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged)
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
func (it *CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged represents a NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged is a free log retrieval operation binding the contract event 0x3312247fcf207243294680d9103851c8709e19be3d46ee4b1bff5e12d1c5ca7b.
//
// Solidity: event NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterNumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged", logs: logs, sub: sub}, nil
}

// WatchNumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged is a free log subscription operation binding the contract event 0x3312247fcf207243294680d9103851c8709e19be3d46ee4b1bff5e12d1c5ca7b.
//
// Solidity: event NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchNumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseNumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged(log types.Log) (*CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged, error) {
	event := new(CosmicSignatureGameV3NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3NumRaffleEthPrizesForBiddersChangedIterator is returned from FilterNumRaffleEthPrizesForBiddersChanged and is used to iterate over the raw logs and unpacked data for NumRaffleEthPrizesForBiddersChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3NumRaffleEthPrizesForBiddersChangedIterator struct {
	Event *CosmicSignatureGameV3NumRaffleEthPrizesForBiddersChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3NumRaffleEthPrizesForBiddersChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3NumRaffleEthPrizesForBiddersChanged)
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
		it.Event = new(CosmicSignatureGameV3NumRaffleEthPrizesForBiddersChanged)
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
func (it *CosmicSignatureGameV3NumRaffleEthPrizesForBiddersChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3NumRaffleEthPrizesForBiddersChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3NumRaffleEthPrizesForBiddersChanged represents a NumRaffleEthPrizesForBiddersChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3NumRaffleEthPrizesForBiddersChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNumRaffleEthPrizesForBiddersChanged is a free log retrieval operation binding the contract event 0x4787028773c8f14bc3b4bc41f43a02329ae41105823287201a34782c530d35fd.
//
// Solidity: event NumRaffleEthPrizesForBiddersChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterNumRaffleEthPrizesForBiddersChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3NumRaffleEthPrizesForBiddersChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "NumRaffleEthPrizesForBiddersChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3NumRaffleEthPrizesForBiddersChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "NumRaffleEthPrizesForBiddersChanged", logs: logs, sub: sub}, nil
}

// WatchNumRaffleEthPrizesForBiddersChanged is a free log subscription operation binding the contract event 0x4787028773c8f14bc3b4bc41f43a02329ae41105823287201a34782c530d35fd.
//
// Solidity: event NumRaffleEthPrizesForBiddersChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchNumRaffleEthPrizesForBiddersChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3NumRaffleEthPrizesForBiddersChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "NumRaffleEthPrizesForBiddersChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3NumRaffleEthPrizesForBiddersChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "NumRaffleEthPrizesForBiddersChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseNumRaffleEthPrizesForBiddersChanged(log types.Log) (*CosmicSignatureGameV3NumRaffleEthPrizesForBiddersChanged, error) {
	event := new(CosmicSignatureGameV3NumRaffleEthPrizesForBiddersChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "NumRaffleEthPrizesForBiddersChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3OwnershipTransferredIterator struct {
	Event *CosmicSignatureGameV3OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3OwnershipTransferred)
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
		it.Event = new(CosmicSignatureGameV3OwnershipTransferred)
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
func (it *CosmicSignatureGameV3OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3OwnershipTransferred represents a OwnershipTransferred event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CosmicSignatureGameV3OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3OwnershipTransferredIterator{contract: _CosmicSignatureGameV3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3OwnershipTransferred)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseOwnershipTransferred(log types.Log) (*CosmicSignatureGameV3OwnershipTransferred, error) {
	event := new(CosmicSignatureGameV3OwnershipTransferred)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3PrizesWalletAddressChangedIterator is returned from FilterPrizesWalletAddressChanged and is used to iterate over the raw logs and unpacked data for PrizesWalletAddressChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3PrizesWalletAddressChangedIterator struct {
	Event *CosmicSignatureGameV3PrizesWalletAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3PrizesWalletAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3PrizesWalletAddressChanged)
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
		it.Event = new(CosmicSignatureGameV3PrizesWalletAddressChanged)
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
func (it *CosmicSignatureGameV3PrizesWalletAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3PrizesWalletAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3PrizesWalletAddressChanged represents a PrizesWalletAddressChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3PrizesWalletAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPrizesWalletAddressChanged is a free log retrieval operation binding the contract event 0xb4cecfe1346c94da27291cf5a02969d5fe0b5c36eca49b04fcd60841d28c5e13.
//
// Solidity: event PrizesWalletAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterPrizesWalletAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*CosmicSignatureGameV3PrizesWalletAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "PrizesWalletAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3PrizesWalletAddressChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "PrizesWalletAddressChanged", logs: logs, sub: sub}, nil
}

// WatchPrizesWalletAddressChanged is a free log subscription operation binding the contract event 0xb4cecfe1346c94da27291cf5a02969d5fe0b5c36eca49b04fcd60841d28c5e13.
//
// Solidity: event PrizesWalletAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchPrizesWalletAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3PrizesWalletAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "PrizesWalletAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3PrizesWalletAddressChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "PrizesWalletAddressChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParsePrizesWalletAddressChanged(log types.Log) (*CosmicSignatureGameV3PrizesWalletAddressChanged, error) {
	event := new(CosmicSignatureGameV3PrizesWalletAddressChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "PrizesWalletAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3RaffleTotalEthPrizeAmountForBiddersPercentageChangedIterator is returned from FilterRaffleTotalEthPrizeAmountForBiddersPercentageChanged and is used to iterate over the raw logs and unpacked data for RaffleTotalEthPrizeAmountForBiddersPercentageChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3RaffleTotalEthPrizeAmountForBiddersPercentageChangedIterator struct {
	Event *CosmicSignatureGameV3RaffleTotalEthPrizeAmountForBiddersPercentageChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3RaffleTotalEthPrizeAmountForBiddersPercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3RaffleTotalEthPrizeAmountForBiddersPercentageChanged)
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
		it.Event = new(CosmicSignatureGameV3RaffleTotalEthPrizeAmountForBiddersPercentageChanged)
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
func (it *CosmicSignatureGameV3RaffleTotalEthPrizeAmountForBiddersPercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3RaffleTotalEthPrizeAmountForBiddersPercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3RaffleTotalEthPrizeAmountForBiddersPercentageChanged represents a RaffleTotalEthPrizeAmountForBiddersPercentageChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3RaffleTotalEthPrizeAmountForBiddersPercentageChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRaffleTotalEthPrizeAmountForBiddersPercentageChanged is a free log retrieval operation binding the contract event 0xbfcd8fb930a57c1598c9760db19c84ec766546c3c9a8565611df8302482bfb17.
//
// Solidity: event RaffleTotalEthPrizeAmountForBiddersPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterRaffleTotalEthPrizeAmountForBiddersPercentageChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3RaffleTotalEthPrizeAmountForBiddersPercentageChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "RaffleTotalEthPrizeAmountForBiddersPercentageChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3RaffleTotalEthPrizeAmountForBiddersPercentageChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "RaffleTotalEthPrizeAmountForBiddersPercentageChanged", logs: logs, sub: sub}, nil
}

// WatchRaffleTotalEthPrizeAmountForBiddersPercentageChanged is a free log subscription operation binding the contract event 0xbfcd8fb930a57c1598c9760db19c84ec766546c3c9a8565611df8302482bfb17.
//
// Solidity: event RaffleTotalEthPrizeAmountForBiddersPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchRaffleTotalEthPrizeAmountForBiddersPercentageChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3RaffleTotalEthPrizeAmountForBiddersPercentageChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "RaffleTotalEthPrizeAmountForBiddersPercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3RaffleTotalEthPrizeAmountForBiddersPercentageChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "RaffleTotalEthPrizeAmountForBiddersPercentageChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseRaffleTotalEthPrizeAmountForBiddersPercentageChanged(log types.Log) (*CosmicSignatureGameV3RaffleTotalEthPrizeAmountForBiddersPercentageChanged, error) {
	event := new(CosmicSignatureGameV3RaffleTotalEthPrizeAmountForBiddersPercentageChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "RaffleTotalEthPrizeAmountForBiddersPercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3RaffleWinnerBidderEthPrizeAllocatedIterator is returned from FilterRaffleWinnerBidderEthPrizeAllocated and is used to iterate over the raw logs and unpacked data for RaffleWinnerBidderEthPrizeAllocated events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3RaffleWinnerBidderEthPrizeAllocatedIterator struct {
	Event *CosmicSignatureGameV3RaffleWinnerBidderEthPrizeAllocated // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3RaffleWinnerBidderEthPrizeAllocatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3RaffleWinnerBidderEthPrizeAllocated)
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
		it.Event = new(CosmicSignatureGameV3RaffleWinnerBidderEthPrizeAllocated)
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
func (it *CosmicSignatureGameV3RaffleWinnerBidderEthPrizeAllocatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3RaffleWinnerBidderEthPrizeAllocatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3RaffleWinnerBidderEthPrizeAllocated represents a RaffleWinnerBidderEthPrizeAllocated event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3RaffleWinnerBidderEthPrizeAllocated struct {
	RoundNum       *big.Int
	WinnerIndex    *big.Int
	WinnerAddress  common.Address
	EthPrizeAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRaffleWinnerBidderEthPrizeAllocated is a free log retrieval operation binding the contract event 0x9c62e2cb8cbd10bf3b8a5760977d719fc3475fa67cb7ef9d2e1701f275e885c4.
//
// Solidity: event RaffleWinnerBidderEthPrizeAllocated(uint256 indexed roundNum, uint256 winnerIndex, address indexed winnerAddress, uint256 ethPrizeAmount)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterRaffleWinnerBidderEthPrizeAllocated(opts *bind.FilterOpts, roundNum []*big.Int, winnerAddress []common.Address) (*CosmicSignatureGameV3RaffleWinnerBidderEthPrizeAllocatedIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	var winnerAddressRule []interface{}
	for _, winnerAddressItem := range winnerAddress {
		winnerAddressRule = append(winnerAddressRule, winnerAddressItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "RaffleWinnerBidderEthPrizeAllocated", roundNumRule, winnerAddressRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3RaffleWinnerBidderEthPrizeAllocatedIterator{contract: _CosmicSignatureGameV3.contract, event: "RaffleWinnerBidderEthPrizeAllocated", logs: logs, sub: sub}, nil
}

// WatchRaffleWinnerBidderEthPrizeAllocated is a free log subscription operation binding the contract event 0x9c62e2cb8cbd10bf3b8a5760977d719fc3475fa67cb7ef9d2e1701f275e885c4.
//
// Solidity: event RaffleWinnerBidderEthPrizeAllocated(uint256 indexed roundNum, uint256 winnerIndex, address indexed winnerAddress, uint256 ethPrizeAmount)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchRaffleWinnerBidderEthPrizeAllocated(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3RaffleWinnerBidderEthPrizeAllocated, roundNum []*big.Int, winnerAddress []common.Address) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	var winnerAddressRule []interface{}
	for _, winnerAddressItem := range winnerAddress {
		winnerAddressRule = append(winnerAddressRule, winnerAddressItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "RaffleWinnerBidderEthPrizeAllocated", roundNumRule, winnerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3RaffleWinnerBidderEthPrizeAllocated)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "RaffleWinnerBidderEthPrizeAllocated", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseRaffleWinnerBidderEthPrizeAllocated(log types.Log) (*CosmicSignatureGameV3RaffleWinnerBidderEthPrizeAllocated, error) {
	event := new(CosmicSignatureGameV3RaffleWinnerBidderEthPrizeAllocated)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "RaffleWinnerBidderEthPrizeAllocated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3RaffleWinnerPrizePaidIterator is returned from FilterRaffleWinnerPrizePaid and is used to iterate over the raw logs and unpacked data for RaffleWinnerPrizePaid events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3RaffleWinnerPrizePaidIterator struct {
	Event *CosmicSignatureGameV3RaffleWinnerPrizePaid // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3RaffleWinnerPrizePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3RaffleWinnerPrizePaid)
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
		it.Event = new(CosmicSignatureGameV3RaffleWinnerPrizePaid)
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
func (it *CosmicSignatureGameV3RaffleWinnerPrizePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3RaffleWinnerPrizePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3RaffleWinnerPrizePaid represents a RaffleWinnerPrizePaid event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3RaffleWinnerPrizePaid struct {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterRaffleWinnerPrizePaid(opts *bind.FilterOpts, roundNum []*big.Int, winnerAddress []common.Address, prizeCosmicSignatureNftId []*big.Int) (*CosmicSignatureGameV3RaffleWinnerPrizePaidIterator, error) {

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

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "RaffleWinnerPrizePaid", roundNumRule, winnerAddressRule, prizeCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3RaffleWinnerPrizePaidIterator{contract: _CosmicSignatureGameV3.contract, event: "RaffleWinnerPrizePaid", logs: logs, sub: sub}, nil
}

// WatchRaffleWinnerPrizePaid is a free log subscription operation binding the contract event 0x27c21fe4cea1a3367aa491829dd4dd824296c00910626150464cba8ea5ebb3f4.
//
// Solidity: event RaffleWinnerPrizePaid(uint256 indexed roundNum, bool winnerIsRandomWalkNftStaker, uint256 winnerIndex, address indexed winnerAddress, uint256 cstPrizeAmount, uint256 indexed prizeCosmicSignatureNftId)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchRaffleWinnerPrizePaid(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3RaffleWinnerPrizePaid, roundNum []*big.Int, winnerAddress []common.Address, prizeCosmicSignatureNftId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "RaffleWinnerPrizePaid", roundNumRule, winnerAddressRule, prizeCosmicSignatureNftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3RaffleWinnerPrizePaid)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "RaffleWinnerPrizePaid", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseRaffleWinnerPrizePaid(log types.Log) (*CosmicSignatureGameV3RaffleWinnerPrizePaid, error) {
	event := new(CosmicSignatureGameV3RaffleWinnerPrizePaid)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "RaffleWinnerPrizePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3RandomWalkNftAddressChangedIterator is returned from FilterRandomWalkNftAddressChanged and is used to iterate over the raw logs and unpacked data for RandomWalkNftAddressChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3RandomWalkNftAddressChangedIterator struct {
	Event *CosmicSignatureGameV3RandomWalkNftAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3RandomWalkNftAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3RandomWalkNftAddressChanged)
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
		it.Event = new(CosmicSignatureGameV3RandomWalkNftAddressChanged)
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
func (it *CosmicSignatureGameV3RandomWalkNftAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3RandomWalkNftAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3RandomWalkNftAddressChanged represents a RandomWalkNftAddressChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3RandomWalkNftAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRandomWalkNftAddressChanged is a free log retrieval operation binding the contract event 0xdab38e33e6e11cbb8b085bba9d7426d5e5af01bcc46d2c5957645e4d09e8c49c.
//
// Solidity: event RandomWalkNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterRandomWalkNftAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*CosmicSignatureGameV3RandomWalkNftAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "RandomWalkNftAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3RandomWalkNftAddressChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "RandomWalkNftAddressChanged", logs: logs, sub: sub}, nil
}

// WatchRandomWalkNftAddressChanged is a free log subscription operation binding the contract event 0xdab38e33e6e11cbb8b085bba9d7426d5e5af01bcc46d2c5957645e4d09e8c49c.
//
// Solidity: event RandomWalkNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchRandomWalkNftAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3RandomWalkNftAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "RandomWalkNftAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3RandomWalkNftAddressChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "RandomWalkNftAddressChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseRandomWalkNftAddressChanged(log types.Log) (*CosmicSignatureGameV3RandomWalkNftAddressChanged, error) {
	event := new(CosmicSignatureGameV3RandomWalkNftAddressChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "RandomWalkNftAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3RoundActivationTimeChangedIterator is returned from FilterRoundActivationTimeChanged and is used to iterate over the raw logs and unpacked data for RoundActivationTimeChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3RoundActivationTimeChangedIterator struct {
	Event *CosmicSignatureGameV3RoundActivationTimeChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3RoundActivationTimeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3RoundActivationTimeChanged)
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
		it.Event = new(CosmicSignatureGameV3RoundActivationTimeChanged)
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
func (it *CosmicSignatureGameV3RoundActivationTimeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3RoundActivationTimeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3RoundActivationTimeChanged represents a RoundActivationTimeChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3RoundActivationTimeChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoundActivationTimeChanged is a free log retrieval operation binding the contract event 0x9a2159c1f277ddd727551baedc6a6c4cba77cc5219c8563ee3b15fb67548d89b.
//
// Solidity: event RoundActivationTimeChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterRoundActivationTimeChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3RoundActivationTimeChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "RoundActivationTimeChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3RoundActivationTimeChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "RoundActivationTimeChanged", logs: logs, sub: sub}, nil
}

// WatchRoundActivationTimeChanged is a free log subscription operation binding the contract event 0x9a2159c1f277ddd727551baedc6a6c4cba77cc5219c8563ee3b15fb67548d89b.
//
// Solidity: event RoundActivationTimeChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchRoundActivationTimeChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3RoundActivationTimeChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "RoundActivationTimeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3RoundActivationTimeChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "RoundActivationTimeChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseRoundActivationTimeChanged(log types.Log) (*CosmicSignatureGameV3RoundActivationTimeChanged, error) {
	event := new(CosmicSignatureGameV3RoundActivationTimeChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "RoundActivationTimeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3RoundLateBidDurationDivisorChangedIterator is returned from FilterRoundLateBidDurationDivisorChanged and is used to iterate over the raw logs and unpacked data for RoundLateBidDurationDivisorChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3RoundLateBidDurationDivisorChangedIterator struct {
	Event *CosmicSignatureGameV3RoundLateBidDurationDivisorChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3RoundLateBidDurationDivisorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3RoundLateBidDurationDivisorChanged)
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
		it.Event = new(CosmicSignatureGameV3RoundLateBidDurationDivisorChanged)
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
func (it *CosmicSignatureGameV3RoundLateBidDurationDivisorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3RoundLateBidDurationDivisorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3RoundLateBidDurationDivisorChanged represents a RoundLateBidDurationDivisorChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3RoundLateBidDurationDivisorChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoundLateBidDurationDivisorChanged is a free log retrieval operation binding the contract event 0x7acba37d1b2d934e554139ff1d470d2cce50a7b6e56870577249b87564a95a69.
//
// Solidity: event RoundLateBidDurationDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterRoundLateBidDurationDivisorChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3RoundLateBidDurationDivisorChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "RoundLateBidDurationDivisorChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3RoundLateBidDurationDivisorChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "RoundLateBidDurationDivisorChanged", logs: logs, sub: sub}, nil
}

// WatchRoundLateBidDurationDivisorChanged is a free log subscription operation binding the contract event 0x7acba37d1b2d934e554139ff1d470d2cce50a7b6e56870577249b87564a95a69.
//
// Solidity: event RoundLateBidDurationDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchRoundLateBidDurationDivisorChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3RoundLateBidDurationDivisorChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "RoundLateBidDurationDivisorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3RoundLateBidDurationDivisorChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "RoundLateBidDurationDivisorChanged", log); err != nil {
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

// ParseRoundLateBidDurationDivisorChanged is a log parse operation binding the contract event 0x7acba37d1b2d934e554139ff1d470d2cce50a7b6e56870577249b87564a95a69.
//
// Solidity: event RoundLateBidDurationDivisorChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseRoundLateBidDurationDivisorChanged(log types.Log) (*CosmicSignatureGameV3RoundLateBidDurationDivisorChanged, error) {
	event := new(CosmicSignatureGameV3RoundLateBidDurationDivisorChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "RoundLateBidDurationDivisorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3RoundLateBidPricePremiumAmountBaseMultiplierChangedIterator is returned from FilterRoundLateBidPricePremiumAmountBaseMultiplierChanged and is used to iterate over the raw logs and unpacked data for RoundLateBidPricePremiumAmountBaseMultiplierChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3RoundLateBidPricePremiumAmountBaseMultiplierChangedIterator struct {
	Event *CosmicSignatureGameV3RoundLateBidPricePremiumAmountBaseMultiplierChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3RoundLateBidPricePremiumAmountBaseMultiplierChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3RoundLateBidPricePremiumAmountBaseMultiplierChanged)
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
		it.Event = new(CosmicSignatureGameV3RoundLateBidPricePremiumAmountBaseMultiplierChanged)
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
func (it *CosmicSignatureGameV3RoundLateBidPricePremiumAmountBaseMultiplierChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3RoundLateBidPricePremiumAmountBaseMultiplierChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3RoundLateBidPricePremiumAmountBaseMultiplierChanged represents a RoundLateBidPricePremiumAmountBaseMultiplierChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3RoundLateBidPricePremiumAmountBaseMultiplierChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoundLateBidPricePremiumAmountBaseMultiplierChanged is a free log retrieval operation binding the contract event 0x169f25ec19cc5b518dc57adf05bb2d85155b1d6c60767e9f0e3ced18eac0ab77.
//
// Solidity: event RoundLateBidPricePremiumAmountBaseMultiplierChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterRoundLateBidPricePremiumAmountBaseMultiplierChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3RoundLateBidPricePremiumAmountBaseMultiplierChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "RoundLateBidPricePremiumAmountBaseMultiplierChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3RoundLateBidPricePremiumAmountBaseMultiplierChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "RoundLateBidPricePremiumAmountBaseMultiplierChanged", logs: logs, sub: sub}, nil
}

// WatchRoundLateBidPricePremiumAmountBaseMultiplierChanged is a free log subscription operation binding the contract event 0x169f25ec19cc5b518dc57adf05bb2d85155b1d6c60767e9f0e3ced18eac0ab77.
//
// Solidity: event RoundLateBidPricePremiumAmountBaseMultiplierChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchRoundLateBidPricePremiumAmountBaseMultiplierChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3RoundLateBidPricePremiumAmountBaseMultiplierChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "RoundLateBidPricePremiumAmountBaseMultiplierChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3RoundLateBidPricePremiumAmountBaseMultiplierChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "RoundLateBidPricePremiumAmountBaseMultiplierChanged", log); err != nil {
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

// ParseRoundLateBidPricePremiumAmountBaseMultiplierChanged is a log parse operation binding the contract event 0x169f25ec19cc5b518dc57adf05bb2d85155b1d6c60767e9f0e3ced18eac0ab77.
//
// Solidity: event RoundLateBidPricePremiumAmountBaseMultiplierChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseRoundLateBidPricePremiumAmountBaseMultiplierChanged(log types.Log) (*CosmicSignatureGameV3RoundLateBidPricePremiumAmountBaseMultiplierChanged, error) {
	event := new(CosmicSignatureGameV3RoundLateBidPricePremiumAmountBaseMultiplierChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "RoundLateBidPricePremiumAmountBaseMultiplierChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3RoundLateBidPricePremiumAmountExponentChangedIterator is returned from FilterRoundLateBidPricePremiumAmountExponentChanged and is used to iterate over the raw logs and unpacked data for RoundLateBidPricePremiumAmountExponentChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3RoundLateBidPricePremiumAmountExponentChangedIterator struct {
	Event *CosmicSignatureGameV3RoundLateBidPricePremiumAmountExponentChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3RoundLateBidPricePremiumAmountExponentChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3RoundLateBidPricePremiumAmountExponentChanged)
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
		it.Event = new(CosmicSignatureGameV3RoundLateBidPricePremiumAmountExponentChanged)
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
func (it *CosmicSignatureGameV3RoundLateBidPricePremiumAmountExponentChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3RoundLateBidPricePremiumAmountExponentChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3RoundLateBidPricePremiumAmountExponentChanged represents a RoundLateBidPricePremiumAmountExponentChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3RoundLateBidPricePremiumAmountExponentChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoundLateBidPricePremiumAmountExponentChanged is a free log retrieval operation binding the contract event 0xcb78cca7628d232a9c7beef53b62f7204d9eacb44de85a8f593e6b0bb72a1621.
//
// Solidity: event RoundLateBidPricePremiumAmountExponentChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterRoundLateBidPricePremiumAmountExponentChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3RoundLateBidPricePremiumAmountExponentChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "RoundLateBidPricePremiumAmountExponentChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3RoundLateBidPricePremiumAmountExponentChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "RoundLateBidPricePremiumAmountExponentChanged", logs: logs, sub: sub}, nil
}

// WatchRoundLateBidPricePremiumAmountExponentChanged is a free log subscription operation binding the contract event 0xcb78cca7628d232a9c7beef53b62f7204d9eacb44de85a8f593e6b0bb72a1621.
//
// Solidity: event RoundLateBidPricePremiumAmountExponentChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchRoundLateBidPricePremiumAmountExponentChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3RoundLateBidPricePremiumAmountExponentChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "RoundLateBidPricePremiumAmountExponentChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3RoundLateBidPricePremiumAmountExponentChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "RoundLateBidPricePremiumAmountExponentChanged", log); err != nil {
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

// ParseRoundLateBidPricePremiumAmountExponentChanged is a log parse operation binding the contract event 0xcb78cca7628d232a9c7beef53b62f7204d9eacb44de85a8f593e6b0bb72a1621.
//
// Solidity: event RoundLateBidPricePremiumAmountExponentChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseRoundLateBidPricePremiumAmountExponentChanged(log types.Log) (*CosmicSignatureGameV3RoundLateBidPricePremiumAmountExponentChanged, error) {
	event := new(CosmicSignatureGameV3RoundLateBidPricePremiumAmountExponentChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "RoundLateBidPricePremiumAmountExponentChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3StakingWalletCosmicSignatureNftAddressChangedIterator is returned from FilterStakingWalletCosmicSignatureNftAddressChanged and is used to iterate over the raw logs and unpacked data for StakingWalletCosmicSignatureNftAddressChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3StakingWalletCosmicSignatureNftAddressChangedIterator struct {
	Event *CosmicSignatureGameV3StakingWalletCosmicSignatureNftAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3StakingWalletCosmicSignatureNftAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3StakingWalletCosmicSignatureNftAddressChanged)
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
		it.Event = new(CosmicSignatureGameV3StakingWalletCosmicSignatureNftAddressChanged)
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
func (it *CosmicSignatureGameV3StakingWalletCosmicSignatureNftAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3StakingWalletCosmicSignatureNftAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3StakingWalletCosmicSignatureNftAddressChanged represents a StakingWalletCosmicSignatureNftAddressChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3StakingWalletCosmicSignatureNftAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakingWalletCosmicSignatureNftAddressChanged is a free log retrieval operation binding the contract event 0x4da1815cd654922275d14d2335fd9a0dd0aa6a0d0ff87fb4cc872ebe9704596f.
//
// Solidity: event StakingWalletCosmicSignatureNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterStakingWalletCosmicSignatureNftAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*CosmicSignatureGameV3StakingWalletCosmicSignatureNftAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "StakingWalletCosmicSignatureNftAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3StakingWalletCosmicSignatureNftAddressChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "StakingWalletCosmicSignatureNftAddressChanged", logs: logs, sub: sub}, nil
}

// WatchStakingWalletCosmicSignatureNftAddressChanged is a free log subscription operation binding the contract event 0x4da1815cd654922275d14d2335fd9a0dd0aa6a0d0ff87fb4cc872ebe9704596f.
//
// Solidity: event StakingWalletCosmicSignatureNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchStakingWalletCosmicSignatureNftAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3StakingWalletCosmicSignatureNftAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "StakingWalletCosmicSignatureNftAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3StakingWalletCosmicSignatureNftAddressChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "StakingWalletCosmicSignatureNftAddressChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseStakingWalletCosmicSignatureNftAddressChanged(log types.Log) (*CosmicSignatureGameV3StakingWalletCosmicSignatureNftAddressChanged, error) {
	event := new(CosmicSignatureGameV3StakingWalletCosmicSignatureNftAddressChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "StakingWalletCosmicSignatureNftAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3StakingWalletRandomWalkNftAddressChangedIterator is returned from FilterStakingWalletRandomWalkNftAddressChanged and is used to iterate over the raw logs and unpacked data for StakingWalletRandomWalkNftAddressChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3StakingWalletRandomWalkNftAddressChangedIterator struct {
	Event *CosmicSignatureGameV3StakingWalletRandomWalkNftAddressChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3StakingWalletRandomWalkNftAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3StakingWalletRandomWalkNftAddressChanged)
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
		it.Event = new(CosmicSignatureGameV3StakingWalletRandomWalkNftAddressChanged)
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
func (it *CosmicSignatureGameV3StakingWalletRandomWalkNftAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3StakingWalletRandomWalkNftAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3StakingWalletRandomWalkNftAddressChanged represents a StakingWalletRandomWalkNftAddressChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3StakingWalletRandomWalkNftAddressChanged struct {
	NewValue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakingWalletRandomWalkNftAddressChanged is a free log retrieval operation binding the contract event 0xbf6e296f85d08cc1ab124aed644bf4b19e4a726a7aea53e3784ab1341738a040.
//
// Solidity: event StakingWalletRandomWalkNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterStakingWalletRandomWalkNftAddressChanged(opts *bind.FilterOpts, newValue []common.Address) (*CosmicSignatureGameV3StakingWalletRandomWalkNftAddressChangedIterator, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "StakingWalletRandomWalkNftAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3StakingWalletRandomWalkNftAddressChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "StakingWalletRandomWalkNftAddressChanged", logs: logs, sub: sub}, nil
}

// WatchStakingWalletRandomWalkNftAddressChanged is a free log subscription operation binding the contract event 0xbf6e296f85d08cc1ab124aed644bf4b19e4a726a7aea53e3784ab1341738a040.
//
// Solidity: event StakingWalletRandomWalkNftAddressChanged(address indexed newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchStakingWalletRandomWalkNftAddressChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3StakingWalletRandomWalkNftAddressChanged, newValue []common.Address) (event.Subscription, error) {

	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "StakingWalletRandomWalkNftAddressChanged", newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3StakingWalletRandomWalkNftAddressChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "StakingWalletRandomWalkNftAddressChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseStakingWalletRandomWalkNftAddressChanged(log types.Log) (*CosmicSignatureGameV3StakingWalletRandomWalkNftAddressChanged, error) {
	event := new(CosmicSignatureGameV3StakingWalletRandomWalkNftAddressChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "StakingWalletRandomWalkNftAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3TimeoutDurationToClaimMainPrizeChangedIterator is returned from FilterTimeoutDurationToClaimMainPrizeChanged and is used to iterate over the raw logs and unpacked data for TimeoutDurationToClaimMainPrizeChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3TimeoutDurationToClaimMainPrizeChangedIterator struct {
	Event *CosmicSignatureGameV3TimeoutDurationToClaimMainPrizeChanged // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3TimeoutDurationToClaimMainPrizeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3TimeoutDurationToClaimMainPrizeChanged)
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
		it.Event = new(CosmicSignatureGameV3TimeoutDurationToClaimMainPrizeChanged)
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
func (it *CosmicSignatureGameV3TimeoutDurationToClaimMainPrizeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3TimeoutDurationToClaimMainPrizeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3TimeoutDurationToClaimMainPrizeChanged represents a TimeoutDurationToClaimMainPrizeChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3TimeoutDurationToClaimMainPrizeChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTimeoutDurationToClaimMainPrizeChanged is a free log retrieval operation binding the contract event 0x37a332914fac995349420c0419b4423a19dcb762017f691442a0782ce4bf417a.
//
// Solidity: event TimeoutDurationToClaimMainPrizeChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterTimeoutDurationToClaimMainPrizeChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3TimeoutDurationToClaimMainPrizeChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "TimeoutDurationToClaimMainPrizeChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3TimeoutDurationToClaimMainPrizeChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "TimeoutDurationToClaimMainPrizeChanged", logs: logs, sub: sub}, nil
}

// WatchTimeoutDurationToClaimMainPrizeChanged is a free log subscription operation binding the contract event 0x37a332914fac995349420c0419b4423a19dcb762017f691442a0782ce4bf417a.
//
// Solidity: event TimeoutDurationToClaimMainPrizeChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchTimeoutDurationToClaimMainPrizeChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3TimeoutDurationToClaimMainPrizeChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "TimeoutDurationToClaimMainPrizeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3TimeoutDurationToClaimMainPrizeChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "TimeoutDurationToClaimMainPrizeChanged", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseTimeoutDurationToClaimMainPrizeChanged(log types.Log) (*CosmicSignatureGameV3TimeoutDurationToClaimMainPrizeChanged, error) {
	event := new(CosmicSignatureGameV3TimeoutDurationToClaimMainPrizeChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "TimeoutDurationToClaimMainPrizeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmicSignatureGameV3UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3UpgradedIterator struct {
	Event *CosmicSignatureGameV3Upgraded // Event containing the contract specifics and raw log

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
func (it *CosmicSignatureGameV3UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3Upgraded)
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
		it.Event = new(CosmicSignatureGameV3Upgraded)
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
func (it *CosmicSignatureGameV3UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3Upgraded represents a Upgraded event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CosmicSignatureGameV3UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3UpgradedIterator{contract: _CosmicSignatureGameV3.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3Upgraded)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseUpgraded(log types.Log) (*CosmicSignatureGameV3Upgraded, error) {
	event := new(CosmicSignatureGameV3Upgraded)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
