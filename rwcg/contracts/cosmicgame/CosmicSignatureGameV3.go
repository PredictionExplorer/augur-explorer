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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit\",\"type\":\"uint256\"}],\"name\":\"BidCstRewardAmountMinLimitNotReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"BidHasBeenPlacedInCurrentRound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"}],\"name\":\"CallerIsNotNftOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientReceivedBidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"InvalidOperationInCurrentState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"lastBidderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"durationUntilOperationIsPermitted\",\"type\":\"uint256\"}],\"name\":\"MainPrizeClaimDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"mainPrizeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"MainPrizeEarlyClaim\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"NoBidsPlacedInCurrentRound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"roundActivationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"RoundIsActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"roundActivationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"RoundIsInactive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"messageLength\",\"type\":\"uint256\"}],\"name\":\"TooLongBidMessage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"randomWalkNftId\",\"type\":\"uint256\"}],\"name\":\"UsedRandomWalkNft\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"WrongBidType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ArbitrumError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"BidCstRewardAmountMultiplierChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"BidMessageLengthMaxLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lastBidderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"paidEthPrice\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"paidCstPrice\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"randomWalkNftId\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstDutchAuctionDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mainPrizeTime\",\"type\":\"uint256\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CharityAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CharityEthDonationAmountPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"ChronoWarriorEthPrizeAmountPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chronoWarriorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeCosmicSignatureNftId\",\"type\":\"uint256\"}],\"name\":\"ChronoWarriorPrizePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICosmicSignatureNft\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CosmicSignatureNftAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICosmicSignatureToken\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CosmicSignatureTokenAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CstDutchAuctionBeginningBidPriceMinLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CstDutchAuctionDurationChangeDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CstDutchAuctionDurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CstPrizeAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"DelayDurationBeforeRoundActivationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"enduranceChampionAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeCosmicSignatureNftId\",\"type\":\"uint256\"}],\"name\":\"EnduranceChampionPrizePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"EthBidPriceIncreaseDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"EthBidRefundAmountInGasToSwallowMaxLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthDonated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ethDonationWithInfoRecordIndex\",\"type\":\"uint256\"}],\"name\":\"EthDonatedWithInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"EthDutchAuctionDurationDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"EthDutchAuctionEndingBidPriceDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"FirstBidPlacedInRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"charityAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsTransferredToCharity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"InitialDurationUntilMainPrizeDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"LastBidderBidCstRewardAmountPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lastCstBidderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeCosmicSignatureNftId\",\"type\":\"uint256\"}],\"name\":\"LastCstBidderPrizePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MainEthPrizeAmountPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeFirstCosmicSignatureNftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prizeNumCosmicSignatureNfts\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeoutTimeToWithdrawSecondaryPrizes\",\"type\":\"uint256\"}],\"name\":\"MainPrizeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MainPrizeNumCosmicSignatureNftsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MainPrizeTimeIncrementInMicroSecondsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MainPrizeTimeIncrementIncreaseDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"MarketingWalletAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MarketingWalletCstContributionAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"NumRaffleCosmicSignatureNftsForBiddersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"NumRaffleEthPrizesForBiddersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIPrizesWallet\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"PrizesWalletAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RaffleTotalEthPrizeAmountForBiddersPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winnerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethPrizeAmount\",\"type\":\"uint256\"}],\"name\":\"RaffleWinnerBidderEthPrizeAllocated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"winnerIsRandomWalkNftStaker\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winnerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeCosmicSignatureNftId\",\"type\":\"uint256\"}],\"name\":\"RaffleWinnerPrizePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIRandomWalkNFT\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"RandomWalkNftAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RoundActivationTimeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RoundLateBidDurationDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RoundLateBidPricePremiumAmountBaseMultiplierChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RoundLateBidPricePremiumAmountExponentChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIStakingWalletCosmicSignatureNft\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"StakingWalletCosmicSignatureNftAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIStakingWalletRandomWalkNft\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"StakingWalletRandomWalkNftAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"TimeoutDurationToClaimMainPrizeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bidCstRewardAmountMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bidMessageLengthMaxLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceMaxLimit_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"}],\"name\":\"bidWithCst\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceMaxLimit_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"bidWithCstAndDonateNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceMaxLimit_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"bidWithCstAndDonateToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"randomWalkNftId_\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"}],\"name\":\"bidWithEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"randomWalkNftId_\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"bidWithEthAndDonateNft\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"randomWalkNftId_\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"bidWithEthAndDonateToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"}],\"name\":\"bidderAddresses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"numItems\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidderAddress\",\"type\":\"address\"}],\"name\":\"biddersInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalSpentEthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSpentCstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBidTimeStamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"}],\"name\":\"championDurations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"enduranceChampion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chronoWarrior\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charityAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charityEthDonationAmountPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chronoWarriorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chronoWarriorDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chronoWarriorEthPrizeAmountPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimMainPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cosmicSignatureNftStakingTotalEthRewardAmountPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionBeginningBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionBeginningBidPriceMinLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionBeginningTimeStamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionDurationChangeDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstPrizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayDurationBeforeRoundActivation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donateEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"data_\",\"type\":\"string\"}],\"name\":\"donateEthWithInfo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enduranceChampionAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enduranceChampionDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enduranceChampionStartTimeStamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethBidPriceIncreaseDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethBidRefundAmountInGasToSwallowMaxLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ethDonationWithInfoRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethDutchAuctionBeginningBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethDutchAuctionDurationDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethDutchAuctionEndingBidPriceDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBidCstRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"currentTimeOffset_\",\"type\":\"int256\"}],\"name\":\"getBidCstRewardAmountAdvanced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidIndex_\",\"type\":\"uint256\"}],\"name\":\"getBidderAddressAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidderAddress_\",\"type\":\"address\"}],\"name\":\"getBidderTotalSpentAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCharityEthDonationAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChronoWarriorEthPrizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCosmicSignatureNftStakingTotalEthRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCstDutchAuctionDurations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDurationElapsedSinceRoundActivation\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDurationUntilMainPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDurationUntilMainPrizeRaw\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDurationUntilRoundActivation\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEthDutchAuctionDurations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethBidPrice_\",\"type\":\"uint256\"}],\"name\":\"getEthPlusRandomWalkNftBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitialDurationUntilMainPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMainEthPrizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMainPrizeTimeIncrement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextCstBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"currentTimeOffset_\",\"type\":\"int256\"}],\"name\":\"getNextCstBidPriceAdvanced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextEthBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"currentTimeOffset_\",\"type\":\"int256\"}],\"name\":\"getNextEthBidPriceAdvanced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRaffleTotalEthPrizeAmountForBidders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoundLateBidDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"}],\"name\":\"getTotalNumBids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"halveEthDutchAuctionEndingBidPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialDurationUntilMainPrizeDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBidderAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBidderBidCstRewardAmountPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCstBidderAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainEthPrizeAmountPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainPrizeNumCosmicSignatureNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainPrizeTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainPrizeTimeIncrementInMicroSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainPrizeTimeIncrementIncreaseDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketingWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketingWalletCstContributionAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextEthBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRoundFirstCstDutchAuctionBeginningBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractCosmicSignatureNft\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numEthDonationWithInfoRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleCosmicSignatureNftsForBidders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleCosmicSignatureNftsForRandomWalkNftStakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleEthPrizesForBidders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prevEnduranceChampionDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizesWallet\",\"outputs\":[{\"internalType\":\"contractPrizesWallet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raffleTotalEthPrizeAmountForBiddersPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomWalkNft\",\"outputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reinitialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundActivationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundLateBidDurationDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundLateBidPricePremiumAmountBaseMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundLateBidPricePremiumAmountExponent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setBidCstRewardAmountMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setBidMessageLengthMaxLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setCharityAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCharityEthDonationAmountPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setChronoWarriorEthPrizeAmountPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICosmicSignatureNft\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setCosmicSignatureNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCosmicSignatureNftStakingTotalEthRewardAmountPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICosmicSignatureToken\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setCosmicSignatureToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCstDutchAuctionBeginningBidPriceMinLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCstDutchAuctionDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCstDutchAuctionDurationChangeDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCstPrizeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setDelayDurationBeforeRoundActivation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setEthBidPriceIncreaseDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setEthBidRefundAmountInGasToSwallowMaxLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setEthDutchAuctionDurationDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setEthDutchAuctionEndingBidPriceDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setInitialDurationUntilMainPrizeDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setLastBidderBidCstRewardAmountPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setMainEthPrizeAmountPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setMainPrizeNumCosmicSignatureNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setMainPrizeTimeIncrementInMicroSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setMainPrizeTimeIncrementIncreaseDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setMarketingWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setMarketingWalletCstContributionAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleCosmicSignatureNftsForBidders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleCosmicSignatureNftsForRandomWalkNftStakers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleEthPrizesForBidders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPrizesWallet\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setPrizesWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setRaffleTotalEthPrizeAmountForBiddersPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRandomWalkNFT\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setRandomWalkNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setRoundActivationTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setRoundLateBidDurationDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setRoundLateBidPricePremiumAmountBaseMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setRoundLateBidPricePremiumAmountExponent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakingWalletCosmicSignatureNft\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setStakingWalletCosmicSignatureNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakingWalletRandomWalkNft\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setStakingWalletRandomWalkNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setTimeoutDurationToClaimMainPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingWalletCosmicSignatureNft\",\"outputs\":[{\"internalType\":\"contractStakingWalletCosmicSignatureNft\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingWalletRandomWalkNft\",\"outputs\":[{\"internalType\":\"contractStakingWalletRandomWalkNft\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeoutDurationToClaimMainPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractCosmicSignatureToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tryGetCurrentChampions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"enduranceChampionAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"enduranceChampionDuration_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"chronoWarriorAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chronoWarriorDuration_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"usedRandomWalkNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nftWasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523461003e5761001161004d565b610019610043565b61bae86103ed82396080518181816191b90152818161921e01526193ec015261bae890f35b610049565b60405190565b5f80fd5b610055610057565b565b61005f610061565b565b61006961006b565b565b610073610075565b565b61007d61007f565b565b610087610089565b565b610091610093565b565b61009b61009d565b565b6100a56100a7565b565b6100af6100b1565b565b6100b96100bb565b565b6100c36100c5565b565b6100cd6100cf565b565b6100d76100e1565b6100df610310565b565b6100e96100eb565b565b6100f36100f5565b565b6100fd6100ff565b565b610107610109565b565b610111610113565b565b61011b61011d565b565b610125610127565b565b61012f610131565b565b61013961013b565b565b610143610145565b565b61014d61014f565b565b610157610159565b565b610161610163565b565b61016b61016d565b565b610175610177565b565b61017f610181565b565b61018961018b565b565b610193610195565b565b61019d61019f565b565b6101a76101a9565b565b6101b16101b3565b565b6101bb6101bd565b565b6101c56101c7565b565b6101cf610213565b565b60018060a01b031690565b90565b6101f36101ee6101f8926101d1565b6101dc565b6101d1565b90565b610204906101df565b90565b610210906101fb565b90565b61021c30610207565b608052565b60401c90565b60ff1690565b61023961023e91610221565b610227565b90565b61024b905461022d565b90565b5f0190565b5f1c90565b60018060401b031690565b61026f61027491610253565b610258565b90565b6102819054610263565b90565b60018060401b031690565b5f1b90565b906102a560018060401b039161028f565b9181191691161790565b6102c36102be6102c892610284565b6101dc565b610284565b90565b90565b906102e36102de6102ea926102af565b6102cb565b8254610294565b9055565b6102f790610284565b9052565b919061030e905f602085019401906102ee565b565b6103186103c8565b6103235f8201610241565b6103ac576103325f8201610277565b61034a61034460018060401b03610284565b91610284565b03610353575b50565b610366905f60018060401b0391016102ce565b60018060401b036103a37fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29161039a610043565b918291826102fb565b0390a15f610350565b5f63f92ee8a960e01b8152806103c46004820161024e565b0390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009056fe6080604052600436101561001d575b366138c45761001b616d1e565b005b6100275f356108d4565b80620ac9f1146108cf5780623d9695146108ca578063040d4d31146108c557806304338479146108c057806309632366146108bb57806309794bee146108b65780630a120648146108b15780630b5f95ae146108ac5780630c9be46d146108a75780630eb16be6146108a2578063119b22b31461089d57806311b0d1fe14610898578063135f3d2814610893578063178877311461088e5780631824d5e71461088957806318305de2146108845780631aaba5a51461087f5780631b4103191461087a5780631e9cbb7e146108755780631f1b4aa41461087057806323b31cfc1461086b578063250fadb6146108665780632665c88214610861578063277004811461085c57806327995f07146108575780632afa2580146108525780632b8dcbba1461084d5780632b91c7bb146108485780632d809e88146108435780632d829a2d1461083e5780632f894cd7146108395780632fb3c48f14610834578063320c435c1461082f578063329b95a51461082a57806336750d2c1461082557806337b99cc7146108205780633b9d292e1461081b5780634164b95b14610816578063441b328914610811578063448c6eb11461080c57806344a4b9171461080757806344acc12a14610802578063477adf2a146107fd57806347ccca02146107f85780634c2a4a33146107f35780634e452010146107ee5780634f1ef286146107e95780634f734612146107e45780634feb78b7146107df57806352d1902d146107da578063543f416f146107d557806354ada1d6146107d057806356732241146107cb5780635863a705146107c65780635a1e5bde146107c15780635b0a45d9146107bc5780635d098b38146107b75780635f0112fe146107b25780635fdf49cb146107ad57806360ef8841146107a857806362ed9b53146107a35780636b59acb81461079e5780636b7cbe85146107995780636c0613c0146107945780636c17e3cc1461078f5780636c2eb3501461078a5780636e95d286146107855780636e97083414610780578063715018a61461077b57806371b6d01914610776578063755b4ef71461077157806375ef3b9c1461076c57806375f0a8741461076757806377fa10271461076257806387292a851461075d578063876d5c361461075857806388ce802c146107535780638af9f8b61461074e5780638c94e9ba146107495780638da5cb5b14610744578063928880fa1461073f5780639302020f1461073a5780639646d7581461073557806399bf353d146107305780639aa1b38d1461072b5780639e50acc9146107265780639edeaf8e146107215780639f9a7fcf1461071c578063a35286d114610717578063a4be0d4014610712578063a922ab5d1461070d578063a974201614610708578063aadd1b0314610703578063ad3cb1cc146106fe578063ad4b0e8a146106f9578063afcf2fc4146106f4578063b30f5bb1146106ef578063b4f1b134146106ea578063b5d1f06f146106e5578063b6a94f42146106e0578063b700db5f146106db578063b78d1e2a146106d6578063b9cf9ba5146106d1578063baab4430146106cc578063bb4b3e6f146106c7578063be720ad5146106c2578063c52d8549146106bd578063c7e7a601146106b8578063c87baab5146106b3578063cb720d4d146106ae578063cfb4e599146106a9578063d1f8fcf2146106a4578063d7559b9c1461069f578063d9ab9eaa1461069a578063da9931dd14610695578063ddd6df0714610690578063de704b411461068b578063dfcd00d114610686578063e2f9185f14610681578063e5b3cd141461067c578063eaace30214610677578063eb13430e14610672578063ebaa1ea81461066d578063ebb9bc5c14610668578063ecb5776e14610663578063ef22d15b1461065e578063efeb248a14610659578063f0bdab7c14610654578063f11400f01461064f578063f2fde38b1461064a578063f34d411c14610645578063f444b29814610640578063f49efe9d1461063b578063f7bea07814610636578063fbaf508414610631578063fc0c546a1461062c578063fd77310f14610627578063fd9b3747146106225763fdfb9ba40361000e57613891565b61385a565b613825565b6137f0565b61374b565b613716565b6136d3565b61369e565b61365b565b613628565b6135f3565b6135b0565b61357b565b613536565b613503565b6134ce565b613469565b613424565b6133df565b61339a565b613355565b613312565b6132df565b6132ac565b613277565b613234565b613200565b613178565b613135565b6130fc565b613089565b613044565b612fff565b612fba565b612f75565b612f30565b612efb565b612ebe565b612e1a565b612dc2565b612d89565b612b47565b612ac5565b612a80565b612a3b565b6129f6565b6128f4565b6128bf565b61287a565b6127d8565b6127a3565b61276e565b61272b565b6126a9565b612664565b61261f565b6125dc565b61255a565b612520565b612498565b612465565b61242f565b61239c565b61236f565b6122ce565b61229b565b612266565b612223565b6121ee565b61214a565b612107565b6120d2565b61208d565b61205a565b612027565b611fa5565b611f60565b611eec565b611eb7565b611e82565b611e4d565b611e08565b611dd5565b611da0565b611d6d565b611ceb565b611ca6565b611c61565b611c1e565b611be9565b611b91565b611b5c565b611b32565b611a3c565b611a07565b6119c2565b61191e565b6118db565b611859565b611816565b6117e3565b6117ae565b611769565b611724565b6116ef565b6116b8565b6115e6565b6115a1565b61153f565b61150c565b6114d7565b611435565b611400565b6113cb565b611327565b6112f2565b6112ad565b611278565b61123e565b61118e565b61115b565b611126565b6110e3565b6110ae565b611069565b611020565b610ead565b610e76565b610cb9565b610c47565b610c14565b610b9e565b610b69565b610ad3565b610aa0565b610a6d565b610a38565b6109d1565b610978565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b90565b6108f8816108ec565b036108ff57565b5f80fd5b90503590610910826108ef565b565b919060408382031261093a578061092e610937925f8601610903565b93602001610903565b90565b6108e4565b60018060a01b031690565b6109539061093f565b90565b61095f9061094a565b9052565b9190610976905f60208501940190610956565b565b346109a9576109a561099461098e366004610912565b906138e2565b61099c6108da565b91829182610963565b0390f35b6108e0565b906020828203126109c7576109c4915f01610903565b90565b6108e4565b5f0190565b346109ff576109e96109e43660046109ae565b6139b9565b6109f16108da565b806109fb816109cc565b0390f35b6108e0565b5f910312610a0e57565b6108e4565b90565b610a1f90610a13565b9052565b9190610a36905f60208501940190610a16565b565b34610a6857610a48366004610a04565b610a64610a536139f2565b610a5b6108da565b91829182610a23565b0390f35b6108e0565b34610a9b57610a85610a803660046109ae565b613a90565b610a8d6108da565b80610a97816109cc565b0390f35b6108e0565b34610ace57610ab8610ab33660046109ae565b613b08565b610ac06108da565b80610aca816109cc565b0390f35b6108e0565b34610b0157610aeb610ae63660046109ae565b613b6d565b610af36108da565b80610afd816109cc565b0390f35b6108e0565b610b0f81610a13565b03610b1657565b5f80fd5b90503590610b2782610b06565b565b90602082820312610b4257610b3f915f01610b1a565b90565b6108e4565b610b50906108ec565b9052565b9190610b67905f60208501940190610b47565b565b34610b9957610b95610b84610b7f366004610b29565b613b7c565b610b8c6108da565b91829182610b54565b0390f35b6108e0565b34610bce57610bae366004610a04565b610bca610bb9613b91565b610bc16108da565b91829182610b54565b0390f35b6108e0565b610bdc8161094a565b03610be357565b5f80fd5b90503590610bf482610bd3565b565b90602082820312610c0f57610c0c915f01610be7565b90565b6108e4565b34610c4257610c2c610c27366004610bf6565b613c6e565b610c346108da565b80610c3e816109cc565b0390f35b6108e0565b34610c7757610c57366004610a04565b610c73610c62613ce8565b610c6a6108da565b91829182610b54565b0390f35b6108e0565b1c90565b90565b610c93906008610c989302610c7c565b610c80565b90565b90610ca69154610c83565b90565b610cb661010b5f90610c9b565b90565b34610ce957610cc9366004610a04565b610ce5610cd4610ca9565b610cdc6108da565b91829182610b54565b0390f35b6108e0565b5f80fd5b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b90610d1e90610cf6565b810190811067ffffffffffffffff821117610d3857604052565b610d00565b90610d50610d496108da565b9283610d14565b565b67ffffffffffffffff8111610d7057610d6c602091610cf6565b0190565b610d00565b90825f939282370152565b90929192610d95610d9082610d52565b610d3d565b93818552602085019082840111610db157610daf92610d75565b565b610cf2565b9080601f83011215610dd457816020610dd193359101610d80565b90565b610cee565b610de29061094a565b90565b610dee81610dd9565b03610df557565b5f80fd5b90503590610e0682610de5565b565b919060a083820312610e7157610e20815f8501610903565b92602081013567ffffffffffffffff8111610e6c5782610e41918301610db6565b92610e69610e528460408501610903565b93610e608160608601610df9565b93608001610903565b90565b6108e8565b6108e4565b34610ea857610e92610e89366004610e08565b93929092613e9e565b610e9a6108da565b80610ea4816109cc565b0390f35b6108e0565b34610edb57610ec5610ec03660046109ae565b613f1a565b610ecd6108da565b80610ed7816109cc565b0390f35b6108e0565b9190604083820312610f085780610efc610f05925f8601610903565b93602001610be7565b90565b6108e4565b90565b610f24610f1f610f29926108ec565b610f0d565b6108ec565b90565b90610f3690610f10565b5f5260205260405f2090565b610f56610f51610f5b9261093f565b610f0d565b61093f565b90565b610f6790610f42565b90565b610f7390610f5e565b90565b90610f8090610f6a565b5f5260205260405f2090565b5f1c90565b610f9d610fa291610f8c565b610c80565b90565b610faf9054610f91565b90565b90610fc2610fc792610104610f2c565b610f76565b610fd25f8201610fa5565b91610feb6002610fe460018501610fa5565b9301610fa5565b90565b60409061101761101e949695939661100d60608401985f850190610b47565b6020830190610b47565b0190610b47565b565b346110545761105061103c611036366004610ee0565b90610fb2565b6110479391936108da565b93849384610fee565b0390f35b6108e0565b61106661010a5f90610c9b565b90565b3461109957611079366004610a04565b611095611084611059565b61108c6108da565b91829182610b54565b0390f35b6108e0565b6110ab6101245f90610c9b565b90565b346110de576110be366004610a04565b6110da6110c961109e565b6110d16108da565b91829182610b54565b0390f35b6108e0565b34611111576110fb6110f63660046109ae565b613f92565b6111036108da565b8061110d816109cc565b0390f35b6108e0565b6111236101165f90610c9b565b90565b3461115657611136366004610a04565b611152611141611116565b6111496108da565b91829182610b54565b0390f35b6108e0565b346111895761117361116e3660046109ae565b61400a565b61117b6108da565b80611185816109cc565b0390f35b6108e0565b346111bc576111a66111a13660046109ae565b614082565b6111ae6108da565b806111b8816109cc565b0390f35b6108e0565b5f80fd5b5f80fd5b909182601f830112156112035781359167ffffffffffffffff83116111fe5760200192600183028401116111f957565b6111c5565b6111c1565b610cee565b90602082820312611239575f82013567ffffffffffffffff81116112345761123092016111c9565b9091565b6108e8565b6108e4565b61125261124c366004611208565b906143e2565b61125a6108da565b80611264816109cc565b0390f35b6112756101065f90610c9b565b90565b346112a857611288366004610a04565b6112a4611293611268565b61129b6108da565b91829182610b54565b0390f35b6108e0565b346112dd576112bd366004610a04565b6112d96112c86143fa565b6112d06108da565b91829182610b54565b0390f35b6108e0565b6112ef6101145f90610c9b565b90565b3461132257611302366004610a04565b61131e61130d6112e2565b6113156108da565b91829182610b54565b0390f35b6108e0565b346113575761135361134261133d3660046109ae565b61449e565b61134a6108da565b91829182610b54565b0390f35b6108e0565b60018060a01b031690565b61137790600861137c9302610c7c565b61135c565b90565b9061138a9154611367565b90565b61139a61012c5f9061137f565b90565b6113a690610f5e565b90565b6113b29061139d565b9052565b91906113c9905f602085019401906113a9565b565b346113fb576113db366004610a04565b6113f76113e661138d565b6113ee6108da565b918291826113b6565b0390f35b6108e0565b3461143057611410366004610a04565b61142c61141b6144e0565b6114236108da565b91829182610b54565b0390f35b6108e0565b346114635761144d6114483660046109ae565b61453c565b6114556108da565b8061145f816109cc565b0390f35b6108e0565b60018060a01b031690565b6114839060086114889302610c7c565b611468565b90565b906114969154611473565b90565b6114a661012d5f9061148b565b90565b6114b290610f5e565b90565b6114be906114a9565b9052565b91906114d5905f602085019401906114b5565b565b34611507576114e7366004610a04565b6115036114f2611499565b6114fa6108da565b918291826114c2565b0390f35b6108e0565b3461153a5761152461151f3660046109ae565b6145b4565b61152c6108da565b80611536816109cc565b0390f35b6108e0565b3461156d576115576115523660046109ae565b61462c565b61155f6108da565b80611569816109cc565b0390f35b6108e0565b9061157c90610f10565b5f5260205260405f2090565b5f61159861159e92610103611572565b01610fa5565b90565b346115d1576115cd6115bc6115b73660046109ae565b611588565b6115c46108da565b91829182610b54565b0390f35b6108e0565b6115e361011c5f90610c9b565b90565b34611616576115f6366004610a04565b6116126116016115d6565b6116096108da565b91829182610b54565b0390f35b6108e0565b6116249061094a565b90565b6116308161161b565b0361163757565b5f80fd5b9050359061164882611627565b565b919060a0838203126116b357611662815f8501610903565b92602081013567ffffffffffffffff81116116ae5782611683918301610db6565b926116ab6116948460408501610903565b936116a2816060860161163b565b93608001610903565b90565b6108e8565b6108e4565b346116ea576116d46116cb36600461164a565b93929092614766565b6116dc6108da565b806116e6816109cc565b0390f35b6108e0565b3461171f576116ff366004610a04565b61171b61170a6147cc565b6117126108da565b91829182610b54565b0390f35b6108e0565b3461175457611734366004610a04565b61175061173f61481b565b6117476108da565b91829182610a23565b0390f35b6108e0565b6117666101275f90610c9b565b90565b3461179957611779366004610a04565b611795611784611759565b61178c6108da565b91829182610b54565b0390f35b6108e0565b6117ab6101305f90610c9b565b90565b346117de576117be366004610a04565b6117da6117c961179e565b6117d16108da565b91829182610b54565b0390f35b6108e0565b34611811576117fb6117f63660046109ae565b6148b9565b6118036108da565b8061180d816109cc565b0390f35b6108e0565b3461184457611826366004610a04565b61182e614cdd565b6118366108da565b80611840816109cc565b0390f35b6108e0565b6118566101235f90610c9b565b90565b3461188957611869366004610a04565b611885611874611849565b61187c6108da565b91829182610b54565b0390f35b6108e0565b6118979061094a565b90565b6118a38161188e565b036118aa57565b5f80fd5b905035906118bb8261189a565b565b906020828203126118d6576118d3915f016118ae565b90565b6108e4565b34611909576118f36118ee3660046118bd565b614df2565b6118fb6108da565b80611905816109cc565b0390f35b6108e0565b61191b61011e5f90610c9b565b90565b3461194e5761192e366004610a04565b61194a61193961190e565b6119416108da565b91829182610b54565b0390f35b6108e0565b60018060a01b031690565b61196e9060086119739302610c7c565b611953565b90565b90611981915461195e565b90565b61199161012b5f90611976565b90565b61199d90610f5e565b90565b6119a990611994565b9052565b91906119c0905f602085019401906119a0565b565b346119f2576119d2366004610a04565b6119ee6119dd611984565b6119e56108da565b918291826119ad565b0390f35b6108e0565b611a046101205f90610c9b565b90565b34611a3757611a17366004610a04565b611a33611a226119f7565b611a2a6108da565b91829182610b54565b0390f35b6108e0565b34611a6c57611a68611a57611a52366004610b29565b614dfd565b611a5f6108da565b91829182610b54565b0390f35b6108e0565b67ffffffffffffffff8111611a8f57611a8b602091610cf6565b0190565b610d00565b90929192611aa9611aa482611a71565b610d3d565b93818552602085019082840111611ac557611ac392610d75565b565b610cf2565b9080601f83011215611ae857816020611ae593359101611a94565b90565b610cee565b919091604081840312611b2d57611b06835f8301610be7565b92602082013567ffffffffffffffff8111611b2857611b259201611aca565b90565b6108e8565b6108e4565b611b46611b40366004611aed565b90614e3b565b611b4e6108da565b80611b58816109cc565b0390f35b34611b8c57611b6c366004610a04565b611b88611b77614e9b565b611b7f6108da565b91829182610b54565b0390f35b6108e0565b34611bbf57611ba9611ba43660046109ae565b614f30565b611bb16108da565b80611bbb816109cc565b0390f35b6108e0565b90565b611bd090611bc4565b9052565b9190611be7905f60208501940190611bc7565b565b34611c1957611bf9366004610a04565b611c15611c04614faa565b611c0c6108da565b91829182611bd4565b0390f35b6108e0565b34611c4c57611c36611c313660046109ae565b61502a565b611c3e6108da565b80611c48816109cc565b0390f35b6108e0565b611c5e61011d5f90610c9b565b90565b34611c9157611c71366004610a04565b611c8d611c7c611c51565b611c846108da565b91829182610b54565b0390f35b6108e0565b611ca36101265f90610c9b565b90565b34611cd657611cb6366004610a04565b611cd2611cc1611c96565b611cc96108da565b91829182610b54565b0390f35b6108e0565b611ce86101075f90610c9b565b90565b34611d1b57611cfb366004610a04565b611d17611d06611cdb565b611d0e6108da565b91829182610b54565b0390f35b6108e0565b611d299061094a565b90565b611d3581611d20565b03611d3c57565b5f80fd5b90503590611d4d82611d2c565b565b90602082820312611d6857611d65915f01611d40565b90565b6108e4565b34611d9b57611d85611d80366004611d4f565b615140565b611d8d6108da565b80611d97816109cc565b0390f35b6108e0565b34611dd057611db0366004610a04565b611dcc611dbb61514b565b611dc36108da565b91829182610b54565b0390f35b6108e0565b34611e0357611ded611de8366004610bf6565b615210565b611df56108da565b80611dff816109cc565b0390f35b6108e0565b34611e3857611e18366004610a04565b611e34611e2361521b565b611e2b6108da565b91829182610b54565b0390f35b6108e0565b611e4a6101395f90610c9b565b90565b34611e7d57611e5d366004610a04565b611e79611e68611e3d565b611e706108da565b91829182610b54565b0390f35b6108e0565b34611eb257611e92366004610a04565b611eae611e9d615258565b611ea56108da565b91829182610b54565b0390f35b6108e0565b34611ee757611ec7366004610a04565b611ee3611ed2615283565b611eda6108da565b91829182610b54565b0390f35b6108e0565b34611f1a57611f04611eff3660046109ae565b61530d565b611f0c6108da565b80611f16816109cc565b0390f35b6108e0565b60018060a01b031690565b611f3a906008611f3f9302610c7c565b611f1f565b90565b90611f4d9154611f2a565b90565b611f5d6101095f90611f42565b90565b34611f9057611f70366004610a04565b611f8c611f7b611f50565b611f836108da565b91829182610963565b0390f35b6108e0565b611fa261011a5f90610c9b565b90565b34611fd557611fb5366004610a04565b611fd1611fc0611f95565b611fc86108da565b91829182610b54565b0390f35b6108e0565b611fe39061094a565b90565b611fef81611fda565b03611ff657565b5f80fd5b9050359061200782611fe6565b565b906020828203126120225761201f915f01611ffa565b90565b6108e4565b346120555761203f61203a366004612009565b615423565b6120476108da565b80612051816109cc565b0390f35b6108e0565b346120885761206a366004610a04565b6120726159f1565b61207a6108da565b80612084816109cc565b0390f35b6108e0565b346120bd5761209d366004610a04565b6120b96120a86159fb565b6120b06108da565b91829182610b54565b0390f35b6108e0565b6120cf61010d5f90610c9b565b90565b34612102576120e2366004610a04565b6120fe6120ed6120c2565b6120f56108da565b91829182610b54565b0390f35b6108e0565b3461213557612117366004610a04565b61211f615a3d565b6121276108da565b80612131816109cc565b0390f35b6108e0565b6121476101355f90610c9b565b90565b3461217a5761215a366004610a04565b61217661216561213a565b61216d6108da565b91829182610b54565b0390f35b6108e0565b60018060a01b031690565b61219a90600861219f9302610c7c565b61217f565b90565b906121ad915461218a565b90565b6121bd61012a5f906121a2565b90565b6121c990610f5e565b90565b6121d5906121c0565b9052565b91906121ec905f602085019401906121cc565b565b3461221e576121fe366004610a04565b61221a6122096121b0565b6122116108da565b918291826121d9565b0390f35b6108e0565b346122515761223b6122363660046109ae565b615ab4565b6122436108da565b8061224d816109cc565b0390f35b6108e0565b61226361012f5f90611f42565b90565b3461229657612276366004610a04565b612292612281612256565b6122896108da565b91829182610963565b0390f35b6108e0565b346122c9576122b36122ae3660046109ae565b615b2c565b6122bb6108da565b806122c5816109cc565b0390f35b6108e0565b346122fc576122e66122e13660046109ae565b615ba4565b6122ee6108da565b806122f8816109cc565b0390f35b6108e0565b919060a08382031261236a57612319815f8501610b1a565b92602081013567ffffffffffffffff8111612365578261233a918301610db6565b9261236261234b8460408501610903565b936123598160608601610df9565b93608001610903565b90565b6108e8565b6108e4565b61238661237d366004612301565b93929092615c87565b61238e6108da565b80612398816109cc565b0390f35b346123ca576123b46123af3660046109ae565b615d03565b6123bc6108da565b806123c6816109cc565b0390f35b6108e0565b906123d990610f10565b5f5260205260405f2090565b6123f1906101346123cf565b9061240960016124025f8501610fa5565b9301610fa5565b90565b91602061242d92949361242660408201965f830190610b47565b0190610b47565b565b34612460576124476124423660046109ae565b6123e5565b9061245c6124536108da565b9283928361240c565b0390f35b6108e0565b346124935761247d6124783660046109ae565b615d7b565b6124856108da565b8061248f816109cc565b0390f35b6108e0565b346124c8576124a8366004610a04565b6124c46124b3615d86565b6124bb6108da565b91829182610963565b0390f35b6108e0565b909160608284031261251b576124e5835f8401610b1a565b9260208301359067ffffffffffffffff82116125165761250a81612513938601610db6565b93604001610903565b90565b6108e8565b6108e4565b61253461252e3660046124cd565b91615dd0565b61253c6108da565b80612546816109cc565b0390f35b6125576101155f90610c9b565b90565b3461258a5761256a366004610a04565b61258661257561254a565b61257d6108da565b91829182610b54565b0390f35b6108e0565b6125989061094a565b90565b6125a48161258f565b036125ab57565b5f80fd5b905035906125bc8261259b565b565b906020828203126125d7576125d4915f016125af565b90565b6108e4565b3461260a576125f46125ef3660046125be565b615ee8565b6125fc6108da565b80612606816109cc565b0390f35b6108e0565b61261c6101365f90610c9b565b90565b3461264f5761262f366004610a04565b61264b61263a61260f565b6126426108da565b91829182610b54565b0390f35b6108e0565b6126616101135f90610c9b565b90565b3461269457612674366004610a04565b61269061267f612654565b6126876108da565b91829182610b54565b0390f35b6108e0565b6126a66101055f90611f42565b90565b346126d9576126b9366004610a04565b6126d56126c4612699565b6126cc6108da565b91829182610963565b0390f35b6108e0565b6126e79061094a565b90565b6126f3816126de565b036126fa57565b5f80fd5b9050359061270b826126ea565b565b9060208282031261272657612723915f016126fe565b90565b6108e4565b346127595761274361273e36600461270d565b615ffe565b61274b6108da565b80612755816109cc565b0390f35b6108e0565b61276b6101385f90610c9b565b90565b3461279e5761277e366004610a04565b61279a61278961275e565b6127916108da565b91829182610b54565b0390f35b6108e0565b346127d3576127b3366004610a04565b6127cf6127be616009565b6127c66108da565b91829182610b54565b0390f35b6108e0565b34612806576127f06127eb3660046109ae565b616077565b6127f86108da565b80612802816109cc565b0390f35b6108e0565b60018060a01b031690565b61282690600861282b9302610c7c565b61280b565b90565b906128399154612816565b90565b61284961012e5f9061282e565b90565b61285590610f5e565b90565b6128619061284c565b9052565b9190612878905f60208501940190612858565b565b346128aa5761288a366004610a04565b6128a661289561283c565b61289d6108da565b91829182612865565b0390f35b6108e0565b6128bc6101125f90610c9b565b90565b346128ef576128cf366004610a04565b6128eb6128da6128af565b6128e26108da565b91829182610b54565b0390f35b6108e0565b6128ff366004610a04565b6129076160fe565b61290f6108da565b80612919816109cc565b0390f35b9061292f61292a83610d52565b610d3d565b918252565b5f7f352e302e30000000000000000000000000000000000000000000000000000000910152565b612965600561291d565b9061297260208301612934565b565b61297c61295b565b90565b612987612974565b90565b61299261297f565b90565b5190565b60209181520190565b90825f9392825e0152565b6129cc6129d56020936129da936129c381612995565b93848093612999565b958691016129a2565b610cf6565b0190565b6129f39160208201915f8184039101526129ad565b90565b34612a2657612a06366004610a04565b612a22612a1161298a565b612a196108da565b918291826129de565b0390f35b6108e0565b612a386101025f90611f42565b90565b34612a6b57612a4b366004610a04565b612a67612a56612a2b565b612a5e6108da565b91829182610963565b0390f35b6108e0565b612a7d6101315f90611f42565b90565b34612ab057612a90366004610a04565b612aac612a9b612a70565b612aa36108da565b91829182610963565b0390f35b6108e0565b612ac261011b5f90610c9b565b90565b34612af557612ad5366004610a04565b612af1612ae0612ab5565b612ae86108da565b91829182610b54565b0390f35b6108e0565b612b039061094a565b90565b612b0f81612afa565b03612b1657565b5f80fd5b90503590612b2782612b06565b565b90602082820312612b4257612b3f915f01612b1a565b90565b6108e4565b34612b7557612b5f612b5a366004612b29565b616213565b612b676108da565b80612b71816109cc565b0390f35b6108e0565b634e487b7160e01b5f52603260045260245ffd5b5490565b5f5260205f2090565b612ba481612b8e565b821015612bbe57612bb6600491612b92565b910201905f90565b612b7a565b612bcf612bd491610f8c565b611f1f565b90565b612be19054612bc3565b90565b634e487b7160e01b5f52602260045260245ffd5b9060016002830492168015612c18575b6020831014612c1357565b612be4565b91607f1691612c08565b60209181520190565b5f5260205f2090565b905f9291805490612c4e612c4783612bf8565b8094612c22565b916001811690815f14612ca55750600114612c69575b505050565b612c769192939450612c2b565b915f925b818410612c8d57505001905f8080612c64565b60018160209295939554848601520191019290612c7a565b92949550505060ff19168252151560200201905f8080612c64565b90612cca91612c34565b90565b90612ced612ce692612cdd6108da565b93848092612cc0565b0383610d14565b565b61010090612cfc82612b8e565b811015612d4257612d0c91612b9b565b5090612d195f8301610fa5565b91612d2660018201612bd7565b91612d3f6003612d3860028501610fa5565b9301612ccd565b90565b5f80fd5b9092612d7990612d6f612d869694612d6560808601975f870190610b47565b6020850190610956565b6040830190610b47565b60608184039101526129ad565b90565b34612dbd57612db9612da4612d9f3660046109ae565b612cef565b90612db09492946108da565b94859485612d46565b0390f35b6108e0565b34612df257612dee612ddd612dd8366004610b29565b61621e565b612de56108da565b91829182610b54565b0390f35b6108e0565b916020612e18929493612e1160408201965f830190610b47565b0190610a16565b565b34612e4b57612e2a366004610a04565b612e32616233565b90612e47612e3e6108da565b92839283612df7565b0390f35b6108e0565b919060a083820312612eb957612e68815f8501610b1a565b92602081013567ffffffffffffffff8111612eb45782612e89918301610db6565b92612eb1612e9a8460408501610903565b93612ea8816060860161163b565b93608001610903565b90565b6108e8565b6108e4565b612ed5612ecc366004612e50565b93929092616335565b612edd6108da565b80612ee7816109cc565b0390f35b612ef861010c5f90610c9b565b90565b34612f2b57612f0b366004610a04565b612f27612f16612eeb565b612f1e6108da565b91829182610b54565b0390f35b6108e0565b34612f6057612f40366004610a04565b612f5c612f4b616344565b612f536108da565b91829182610b54565b0390f35b6108e0565b612f726101185f90610c9b565b90565b34612fa557612f85366004610a04565b612fa1612f90612f65565b612f986108da565b91829182610b54565b0390f35b6108e0565b612fb76101325f90610c9b565b90565b34612fea57612fca366004610a04565b612fe6612fd5612faa565b612fdd6108da565b91829182610b54565b0390f35b6108e0565b612ffc6101375f90610c9b565b90565b3461302f5761300f366004610a04565b61302b61301a612fef565b6130226108da565b91829182610b54565b0390f35b6108e0565b6130416101175f90610c9b565b90565b3461307457613054366004610a04565b61307061305f613034565b6130676108da565b91829182610b54565b0390f35b6108e0565b61308661010f5f90610c9b565b90565b346130b957613099366004610a04565b6130b56130a4613079565b6130ac6108da565b91829182610b54565b0390f35b6108e0565b6130f36130fa946130e96060949897956130df608086019a5f870190610956565b6020850190610b47565b6040830190610956565b0190610b47565b565b346131305761310c366004610a04565b61312c613117616386565b906131239492946108da565b948594856130be565b0390f35b6108e0565b346131635761314d6131483660046109ae565b616602565b6131556108da565b8061315f816109cc565b0390f35b6108e0565b61317561010e5f90610c9b565b90565b346131a857613188366004610a04565b6131a4613193613168565b61319b6108da565b91829182610b54565b0390f35b6108e0565b90916060828403126131fb576131c5835f8401610903565b9260208301359067ffffffffffffffff82116131f6576131ea816131f3938601610db6565b93604001610903565b90565b6108e8565b6108e4565b3461322f576132196132133660046131ad565b91616639565b6132216108da565b8061322b816109cc565b0390f35b6108e0565b346132625761324c6132473660046109ae565b6166b3565b6132546108da565b8061325e816109cc565b0390f35b6108e0565b6132746101335f90610c9b565b90565b346132a757613287366004610a04565b6132a3613292613267565b61329a6108da565b91829182610b54565b0390f35b6108e0565b346132da576132c46132bf3660046109ae565b6166ef565b6132cc6108da565b806132d6816109cc565b0390f35b6108e0565b3461330d576132f76132f23660046109ae565b616767565b6132ff6108da565b80613309816109cc565b0390f35b6108e0565b3461334057613322366004610a04565b61332a616945565b6133326108da565b8061333c816109cc565b0390f35b6108e0565b6133526101215f90610c9b565b90565b3461338557613365366004610a04565b613381613370613345565b6133786108da565b91829182610b54565b0390f35b6108e0565b6133976101015f90611f42565b90565b346133ca576133aa366004610a04565b6133c66133b561338a565b6133bd6108da565b91829182610963565b0390f35b6108e0565b6133dc6101085f90610c9b565b90565b3461340f576133ef366004610a04565b61340b6133fa6133cf565b6134026108da565b91829182610b54565b0390f35b6108e0565b6134216101255f90610c9b565b90565b3461345457613434366004610a04565b61345061343f613414565b6134476108da565b91829182610b54565b0390f35b6108e0565b6134666101105f90610c9b565b90565b3461349957613479366004610a04565b613495613484613459565b61348c6108da565b91829182610b54565b0390f35b6108e0565b906134a890610f10565b5f5260205260405f2090565b6134cb906134c6610119915f9261349e565b610c9b565b90565b346134fe576134fa6134e96134e43660046109ae565b6134b4565b6134f16108da565b91829182610b54565b0390f35b6108e0565b346135315761351b6135163660046109ae565b6169bc565b6135236108da565b8061352d816109cc565b0390f35b6108e0565b3461356657613546366004610a04565b6135626135516169d5565b6135596108da565b91829182610a23565b0390f35b6108e0565b6135786101115f90610c9b565b90565b346135ab5761358b366004610a04565b6135a761359661356b565b61359e6108da565b91829182610b54565b0390f35b6108e0565b346135de576135c86135c33660046109ae565b616a5e565b6135d06108da565b806135da816109cc565b0390f35b6108e0565b6135f061011f5f90610c9b565b90565b3461362357613603366004610a04565b61361f61360e6135e3565b6136166108da565b91829182610b54565b0390f35b6108e0565b346136565761364061363b366004610bf6565b616ace565b6136486108da565b80613652816109cc565b0390f35b6108e0565b346136895761367361366e3660046109ae565b616b0a565b61367b6108da565b80613685816109cc565b0390f35b6108e0565b61369b6101285f90610c9b565b90565b346136ce576136ae366004610a04565b6136ca6136b961368e565b6136c16108da565b91829182610b54565b0390f35b6108e0565b34613701576136eb6136e63660046109ae565b616b82565b6136f36108da565b806136fd816109cc565b0390f35b6108e0565b6137136101225f90610c9b565b90565b3461374657613726366004610a04565b613742613731613706565b6137396108da565b91829182610b54565b0390f35b6108e0565b3461377c5761375b366004610a04565b613763616b8d565b9061377861376f6108da565b92839283612df7565b0390f35b6108e0565b60018060a01b031690565b61379c9060086137a19302610c7c565b613781565b90565b906137af915461378c565b90565b6137bf6101295f906137a4565b90565b6137cb90610f5e565b90565b6137d7906137c2565b9052565b91906137ee905f602085019401906137ce565b565b3461382057613800366004610a04565b61381c61380b6137b2565b6138136108da565b918291826137db565b0390f35b6108e0565b346138555761385161384061383b3660046109ae565b616bb6565b6138486108da565b91829182610b54565b0390f35b6108e0565b3461388c5761387361386d366004610ee0565b90616be3565b9061388861387f6108da565b9283928361240c565b0390f35b6108e0565b346138bf576138a96138a43660046109ae565b616c9a565b6138b16108da565b806138bb816109cc565b0390f35b6108e0565b5f80fd5b5f90565b906138d690610f10565b5f5260205260405f2090565b61390b916001613900613906936138f76138c8565b50610103611572565b016138cc565b612bd7565b90565b61391f9061391a616d28565b613921565b565b6139329061392d616e10565b613972565b565b5f1b90565b906139455f1991613934565b9181191691161790565b90565b9061396761396261396e92610f10565b61394f565b8254613939565b9055565b61397e81610138613952565b6139b47fc63013cf34a6f7b20983b293d1787e833f8de2db868e904525fc2910df652a97916139ab6108da565b91829182610b54565b0390a1565b6139c29061390e565b565b5f90565b6139dc6139d76139e1926108ec565b610f0d565b610a13565b90565b906139ef9103610a13565b90565b6139fa6139c4565b50613a20613a07426139c8565b613a1a613a1561010d610fa5565b6139c8565b906139e4565b90565b613a3490613a2f616d28565b613a36565b565b613a4790613a42616e10565b613a49565b565b613a5581610133613952565b613a8b7facbc6b6929088e4b2d043625fa7248e00fb6658a425eb9d9dc8c37b18c7f3e6f91613a826108da565b91829182610b54565b0390a1565b613a9990613a23565b565b613aac90613aa7616d28565b613aae565b565b613abf90613aba616e10565b613ac1565b565b613acd8161011b613952565b613b037f40b9c59af8c486ccf8c7cc73df5a51e7cc29747ea7d39f99632ecaf9caa2ed1f91613afa6108da565b91829182610b54565b0390a1565b613b1190613a9b565b565b613b2490613b1f616d28565b613b26565b565b613b328161010c613952565b613b687fb0868a729f47ae3829aaafe3ca2975d3db2148553c854112f598be6d91ef0d2891613b5f6108da565b91829182610b54565b0390a1565b613b7690613b13565b565b5f90565b613b8e90613b88613b78565b50616e74565b90565b613b99613b78565b50613ba5610100612b8e565b90565b613bb990613bb4616d28565b613bbb565b565b613bcc90613bc7616e10565b613bce565b565b613be090613bdb81616fdc565b613c20565b565b90613bf360018060a01b0391613934565b9181191691161790565b90565b90613c15613c10613c1c92610f6a565b613bfd565b8254613be2565b9055565b613c2c81610131613c00565b613c567f1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c91610f6a565b90613c5f6108da565b80613c69816109cc565b0390a2565b613c7790613ba8565b565b613c8290610f5e565b90565b90613c9091026108ec565b90565b90565b613caa613ca5613caf92613c93565b610f0d565b6108ec565b90565b634e487b7160e01b5f52601260045260245ffd5b613cd2613cd8916108ec565b916108ec565b908115613ce3570490565b613cb2565b613cf0613b78565b50613d22613d12613d0030613c79565b31613d0c610132610fa5565b90613c85565b613d1c6064613c96565b90613cc6565b90565b90613d3a94939291613d3561706b565b613de5565b613d426170b0565b565b613d50613d5591610f8c565b61135c565b90565b613d629054613d44565b90565b5f80fd5b60e01b90565b5f910312613d7957565b6108e4565b613d8790610f5e565b90565b613d9390613d7e565b9052565b613dcc613dd394613dc2606094989795613db8608086019a5f870190610b47565b6020850190610956565b6040830190613d8a565b0190610b47565b565b613ddd6108da565b3d5f823e3d90fd5b91613df592949394919091617206565b613e08613e0361012c613d58565b61139d565b9063e2051c7e91613e1a61010b610fa5565b91613e2361745c565b9490823b15613e99575f94613e568692613e4b94613e3f6108da565b998a9889978896613d69565b865260048601613d97565b03925af18015613e9457613e68575b50565b613e87905f3d8111613e8d575b613e7f8183610d14565b810190613d6f565b5f613e65565b503d613e75565b613dd5565b613d65565b90613eab94939291613d25565b565b613ebe90613eb9616d28565b613ec0565b565b613ed190613ecc616e10565b613ed3565b565b613edf81610121613952565b613f157f3312247fcf207243294680d9103851c8709e19be3d46ee4b1bff5e12d1c5ca7b91613f0c6108da565b91829182610b54565b0390a1565b613f2390613ead565b565b613f3690613f31616d28565b613f38565b565b613f4990613f44616e10565b613f4b565b565b613f5781610137613952565b613f8d7fcb78cca7628d232a9c7beef53b62f7204d9eacb44de85a8f593e6b0bb72a162191613f846108da565b91829182610b54565b0390a1565b613f9b90613f25565b565b613fae90613fa9616d28565b613fb0565b565b613fc190613fbc616e10565b613fc3565b565b613fcf81610120613952565b6140057f85d8bf21006916836edd67a5afeed2e891cf876a6c9cd9babf7f42f4b007c24f91613ffc6108da565b91829182610b54565b0390a1565b61401390613f9d565b565b61402690614021616d28565b614028565b565b61403990614034616e10565b61403b565b565b6140478161011d613952565b61407d7f5581e31f5b8d4b3c45b8ab8bf67e3602ce588b423905eb6ad34bd6bc3c848699916140746108da565b91829182610b54565b0390a1565b61408b90614015565b565b9061409f9161409a61706b565b614308565b6140a76170b0565b565b90565b5490565b5f5260205f2090565b6140c2816140ac565b8210156140dc576140d46004916140b0565b910201905f90565b612b7a565b6140ea816140ac565b6801000000000000000081101561410e5761410a916001820181556140b9565b9091565b610d00565b90565b5090565b601f602091010490565b1b90565b9190600861414391029161413d5f1984614124565b92614124565b9181191691161790565b919061416361415e61416b93610f10565b61394f565b908354614128565b9055565b6141819161417b613b78565b9161414d565b565b5f5b82811061419157505050565b806141a05f600193850161416f565b01614185565b9190601f81116141b6575b505050565b8181116141c3575b6141b1565b6141d86141d26141f694612c2b565b9161411a565b9160206141e48261411a565b91106141fe575b809101910390614183565b5f80806141be565b505f6141eb565b90614215905f1990600802610c7c565b191690565b8161422491614205565b906002021790565b916142379082614116565b9067ffffffffffffffff82116142f65761425b826142558554612bf8565b856141a6565b5f90601f831160011461428e5791809161427d935f92614282575b505061421a565b90555b565b90915001355f80614276565b601f1983169161429d85612c2b565b925f5b8181106142de575091600293918560019694106142c4575b50505002019055614280565b6142d4910135601f841690614205565b90555f80806142b8565b919360206001819287870135815501950192016142a0565b610d00565b90614306929161422c565b565b61437690614317610100612b8e565b92600361433661433061432b6101006140a9565b6140e1565b50614113565b9261434d61434561010b610fa5565b5f8601613952565b61436161435861745c565b60018601613c00565b61436e3460028601613952565b9192016142fb565b61438161010b610fa5565b61438961745c565b3492916143dd6143cb6143c56143bf7fa08049565b10d44a06dca9bf05685b39bc370352043c5a003e8d35d45ebdc53f94610f10565b94610f6a565b94610f10565b946143d46108da565b91829182610b54565b0390a4565b906143ec9161408d565b565b6143f790610f5e565b90565b614402613b78565b50614434614424614412306143ee565b3161441e61011d610fa5565b90613c85565b61442e6064613c96565b90613cc6565b90565b90565b61444e61444961445392614437565b610f0d565b6108ec565b90565b614460600261443a565b90565b90565b61447a61447561447f92614463565b610f0d565b6108ec565b90565b9061448d91036108ec565b90565b9061449b91016108ec565b90565b6144cf6144dd916144ad613b78565b506144c96144b9614456565b6144c36001614466565b90614482565b90614490565b6144d7614456565b90613cc6565b90565b6144e8613b78565b506145086144f7610125610fa5565b614502610123610fa5565b90613cc6565b90565b61451c90614517616d28565b61451e565b565b61452f9061452a6174e8565b614531565b565b61453a90617540565b565b6145459061450b565b565b61455890614553616d28565b61455a565b565b61456b90614566616e10565b61456d565b565b61457981610132613952565b6145af7ffe65b6d5007c66dadebc5095104ccd672c070f396dfdcfe1ee7f54201b9efa6d916145a66108da565b91829182610b54565b0390a1565b6145bd90614547565b565b6145d0906145cb616d28565b6145d2565b565b6145e3906145de616e10565b6145e5565b565b6145f18161011f613952565b6146277f4787028773c8f14bc3b4bc41f43a02329ae41105823287201a34782c530d35fd9161461e6108da565b91829182610b54565b0390a1565b614635906145bf565b565b9061464c9493929161464761706b565b6146ad565b6146546170b0565b565b61465f90610f5e565b90565b61466b90614656565b9052565b6146a46146ab9461469a606094989795614690608086019a5f870190610b47565b6020850190610956565b6040830190614662565b0190610b47565b565b916146bd92949394919091617206565b6146d06146cb61012c613d58565b61139d565b9063fe673fd3916146e261010b610fa5565b916146eb61745c565b9490823b15614761575f9461471e8692614713946147076108da565b998a9889978896613d69565b86526004860161466f565b03925af1801561475c57614730575b50565b61474f905f3d8111614755575b6147478183610d14565b810190613d6f565b5f61472d565b503d61473d565b613dd5565b613d65565b9061477394939291614637565b565b90565b61478c61478761479192614775565b610f0d565b610a13565b90565b6147a86147a36147ad92614775565b610f0d565b6108ec565b90565b6147c46147bf6147c992610a13565b610f0d565b6108ec565b90565b6147d4613b78565b506147dd61481b565b806147f86147f26147ed5f614778565b610a13565b91610a13565b135f1461480c57614808906147b0565b5b90565b506148165f614794565b614809565b6148236139c4565b5061484961483a614835610124610fa5565b6139c8565b614843426139c8565b906139e4565b90565b61485d90614858616d28565b61485f565b565b6148709061486b616e10565b614872565b565b61487e81610113613952565b6148b47fa787f26546d7eeea63d70fc31736f27ad28329e95982f3bc5a7e0280f497bbf5916148ab6108da565b91829182610b54565b0390a1565b6148c29061484c565b565b6148cc61706b565b6148d4614b3c565b6148dc6170b0565b565b6148f26148ed6148f792614775565b610f0d565b61093f565b90565b614903906148de565b90565b151590565b60207f757272656e742062696464696e6720726f756e64207965742e00000000000000917f54686572652068617665206265656e206e6f206269647320696e2074686520635f8201520152565b6149656039604092612999565b61496e8161490b565b0190565b6149879060208101905f818303910152614958565b90565b634e487b7160e01b5f52601160045260245ffd5b6149ad6149b391939293610a13565b92610a13565b808301925f828512158183121692851291121516176149ce57565b61498a565b60607f2e00000000000000000000000000000000000000000000000000000000000000917f4f6e6c7920746865206c61737420626964646572206973207065726d697474655f8201527f6420746f20636c61696d207468652062696464696e6720726f756e64206d616960208201527f6e207072697a65206265666f726520612074696d656f7574206578706972657360408201520152565b614a796061608092612999565b614a82816149d3565b0190565b606090614abf614ac69496959396614ab5614aaa608085018581035f870152614a6c565b986020850190610956565b6040830190610956565b0190610b47565b565b5f7f4e6f7420656e6f7567682074696d652068617320656c61707365642e00000000910152565b614afc601c602092612999565b614b0581614ac8565b0190565b916040614b3a929493614b33614b28606083018381035f850152614aef565b966020830190610b47565b0190610b47565b565b614b4461745c565b614b60614b5a614b55610101612bd7565b61094a565b9161094a565b145f14614bf157614b8f42614b87614b81614b7c610124610fa5565b6108ec565b916108ec565b101515614906565b614bbf575b614b9c617587565b614ba5426176b5565b614bad61773c565b614bb56180db565b614bbd619085565b565b614bca610124610fa5565b4290614bed614bd76108da565b928392638d31bb1560e01b845260048401614b09565b0390fd5b614c21614bff610101612bd7565b614c19614c13614c0e5f6148fa565b61094a565b9161094a565b141515614906565b614cba57614c49614c3061481b565b614c43614c3e610127610fa5565b6139c8565b9061499e565b614c6f81614c67614c61614c5c5f614778565b610a13565b91610a13565b131515614906565b614c795750614b94565b614c84610101612bd7565b614cb6614c98614c9261745c565b936147b0565b614ca06108da565b93849363336598a360e21b855260048501614a86565b0390fd5b614cc26108da565b6318844a7d60e31b815280614cd960048201614972565b0390fd5b614ce56148c4565b565b614cf890614cf3616d28565b614cfa565b565b614d0b90614d06616e10565b614d19565b565b614d1690610f5e565b90565b614d3390614d2e614d2982614d0d565b616fdc565b614d94565b565b614d3e90610f42565b90565b614d4a90614d35565b90565b614d5690614d35565b90565b90565b90614d71614d6c614d7892614d4d565b614d59565b8254613be2565b9055565b614d8590610f42565b90565b614d9190614d7c565b90565b614db0614da8614da383614d0d565b614d41565b61012b614d5c565b614dda7f5bde6238168795ba4e77972a2bdaa5a465f7c9a5d05817f5e8d3fed2e5a4fa6091614d88565b90614de36108da565b80614ded816109cc565b0390a2565b614dfb90614ce7565b565b614e0f90614e09613b78565b5061917e565b90565b90614e2491614e1f6191a8565b614e26565b565b90614e3991614e348161926d565b6192dd565b565b90614e4591614e12565b565b90565b614e5e614e59614e6392614e47565b610f0d565b6108ec565b90565b614e716103e8614e4a565b90565b614e8d614e7f614e66565b614e87614e66565b90613c85565b90565b614e98614e74565b90565b614ea3613b78565b50614ec0614eb2610125610fa5565b614eba614e90565b90613cc6565b90565b614ed490614ecf616d28565b614ed6565b565b614ee790614ee2616e10565b614ee9565b565b614ef581610122613952565b614f2b7f9e44c04f534af356419a731b967f3d56bc748b1f3fdbda7a89f4e1a321ade93491614f226108da565b91829182610b54565b0390a1565b614f3990614ec3565b565b5f90565b614f5090614f4b6193db565b614f9e565b90565b90565b614f6a614f65614f6f92614f53565b613934565b611bc4565b90565b614f9b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc614f56565b90565b50614fa7614f72565b90565b614fba614fb5614f3b565b614f3f565b90565b614fce90614fc9616d28565b614fd0565b565b614fe190614fdc616e10565b614fe3565b565b614fef8161011a613952565b6150257f157c413b0549fd4f45aab72b7828304fb2c45dad53de0f1128c5eabf3aaabaf89161501c6108da565b91829182610b54565b0390a1565b61503390614fbd565b565b61504690615041616d28565b615048565b565b61505990615054616e10565b615067565b565b61506490610f5e565b90565b6150819061507c6150778261505b565b616fdc565b6150e2565b565b61508c90610f42565b90565b61509890615083565b90565b6150a490615083565b90565b90565b906150bf6150ba6150c69261509b565b6150a7565b8254613be2565b9055565b6150d390610f42565b90565b6150df906150ca565b90565b6150fe6150f66150f18361505b565b61508f565b61012e6150aa565b6151287f4da1815cd654922275d14d2335fd9a0dd0aa6a0d0ff87fb4cc872ebe9704596f916150d6565b906151316108da565b8061513b816109cc565b0390a2565b61514990615035565b565b615153613b78565b5061518561517561516330613c79565b3161516f610128610fa5565b90613c85565b61517f6064613c96565b90613cc6565b90565b61519990615194616d28565b61519b565b565b6151ac906151a7616e10565b6151ae565b565b6151c0906151bb81616fdc565b6151c2565b565b6151ce8161012f613c00565b6151f87f4d03942c29c20d1bccfe551e9d148c917c5a44fb558a4fc60270d8f76fb75f5491610f6a565b906152016108da565b8061520b816109cc565b0390a2565b61521990615188565b565b615223613b78565b50615255615245615233306143ee565b3161523f610122610fa5565b90613c85565b61524f6064613c96565b90613cc6565b90565b615260613b78565b5061528061526f610125610fa5565b61527a610135610fa5565b90613cc6565b90565b61528b613b78565b5061529d6152985f614778565b614dfd565b90565b6152b1906152ac616d28565b6152b3565b565b6152c4906152bf616e10565b6152c6565b565b6152d281610128613952565b6153087fb5a05ec7911dd5450a7fa4ae54d50d9d83af6e256db8fc76c82edd7b659cf8bc916152ff6108da565b91829182610b54565b0390a1565b615316906152a0565b565b61532990615324616d28565b61532b565b565b61533c90615337616e10565b61534a565b565b61534790610f5e565b90565b6153649061535f61535a8261533e565b616fdc565b6153c5565b565b61536f90610f42565b90565b61537b90615366565b90565b61538790615366565b90565b90565b906153a261539d6153a99261537e565b61538a565b8254613be2565b9055565b6153b690610f42565b90565b6153c2906153ad565b90565b6153e16153d96153d48361533e565b615372565b61012c61538d565b61540b7fb4cecfe1346c94da27291cf5a02969d5fe0b5c36eca49b04fcd60841d28c5e13916153b9565b906154146108da565b8061541e816109cc565b0390a2565b61542c90615318565b565b615436615438565b565b6154406155c3565b565b90565b61545961545461545e92615442565b610f0d565b6108ec565b90565b61546b6003615445565b90565b67ffffffffffffffff1690565b61548f61548a615494926108ec565b610f0d565b61546e565b90565b60401c90565b60ff1690565b6154af6154b491615497565b61549d565b90565b6154c190546154a3565b90565b67ffffffffffffffff1690565b6154dd6154e291610f8c565b6154c4565b90565b6154ef90546154d1565b90565b9061550567ffffffffffffffff91613934565b9181191691161790565b61552361551e6155289261546e565b610f0d565b61546e565b90565b90565b9061554361553e61554a9261550f565b61552b565b82546154f2565b9055565b60401b90565b9061556868ff00000000000000009161554e565b9181191691161790565b61557b90614906565b90565b90565b9061559661559161559d92615572565b61557e565b8254615554565b9055565b6155aa9061546e565b9052565b91906155c1905f602085019401906155a1565b565b6155d36155ce615461565b61547b565b6155db619439565b6155e65f82016154b7565b8015615676575b61565a5761561f90615601835f830161552e565b61560e60015f8301615581565b61561661597d565b5f809101615581565b6156557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29161564c6108da565b918291826155ae565b0390a1565b5f63f92ee8a960e01b815280615672600482016109cc565b0390fd5b506156825f82016154e5565b61569461568e8461546e565b9161546e565b10156155ed565b90565b6156b26156ad6156b79261569b565b610f0d565b6108ec565b90565b6156c5610e1061569e565b90565b6156d76156dd919392936108ec565b926108ec565b916156e98382026108ec565b9281840414901517156156f857565b61498a565b90565b61571461570f615719926156fd565b610f0d565b6108ec565b90565b6157276104b0615700565b90565b61573661573c916108ec565b916108ec565b908115615747570490565b613cb2565b61575b615761919392936108ec565b926108ec565b820180921161576c57565b61498a565b6157bf6157b16157906157826156ba565b61578a614e90565b906156c8565b6157ab61579b61571c565b6157a5600261443a565b9061572a565b9061574c565b6157b961571c565b9061572a565b90565b90565b6157d96157d46157de926157c2565b610f0d565b6108ec565b90565b6157eb600d6157c5565b90565b90565b61580561580061580a926157ee565b610f0d565b6108ec565b90565b61582c90615826615820615831946108ec565b916108ec565b90614124565b6108ec565b90565b6158516236717961584c6158466157e1565b916157f1565b61580d565b90565b90565b61586b61586661587092615854565b610f0d565b6108ec565b90565b61587d6008615857565b90565b90565b61589761589261589c92615880565b610f0d565b6108ec565b90565b6158b0670de0b6b3a7640000615883565b90565b90565b6158ca6158c56158cf926158b3565b610f0d565b6108ec565b90565b90565b6158e96158e46158ee926158d2565b610f0d565b6108ec565b90565b61594161593161592161591361590561589f565b61590d6156ba565b906156c8565b61591b614e90565b906156c8565b61592b601e6158b6565b9061574c565b61593b603c6158d5565b9061572a565b90565b90565b61595b61595661596092615944565b610f0d565b6108ec565b90565b61596d605a615947565b90565b61597a6003615445565b90565b615990615988615771565b610135613952565b6159a361599b615834565b610136613952565b6159b66159ae615873565b610137613952565b6159c96159c16158f1565b61011b613952565b6159dc6159d4615963565b610138613952565b6159ef6159e7615970565b610139613952565b565b6159f961542e565b565b615a03613b78565b50615a15615a105f614778565b61621e565b90565b615a20616d28565b615a28615a2a565b565b615a3b615a365f6148fa565b61945d565b565b615a45615a18565b565b615a5890615a53616d28565b615a5a565b565b615a6b90615a66616e10565b615a6d565b565b615a7981610136613952565b615aaf7f169f25ec19cc5b518dc57adf05bb2d85155b1d6c60767e9f0e3ced18eac0ab7791615aa66108da565b91829182610b54565b0390a1565b615abd90615a47565b565b615ad090615acb616d28565b615ad2565b565b615ae390615ade616e10565b615ae5565b565b615af181610123613952565b615b277fb5edd1f338b34c8f5dd3b1c5cc12f05653c495713c282bf588d34cf14fad0f8991615b1e6108da565b91829182610b54565b0390a1565b615b3590615abf565b565b615b4890615b43616d28565b615b4a565b565b615b5b90615b56616e10565b615b5d565b565b615b6981610139613952565b615b9f7f616bfcaa6490f55f6e57a4deedac1db04d0d6826deb84fad86cc43439bcf356491615b966108da565b91829182610b54565b0390a1565b615bad90615b37565b565b90615bc494939291615bbf61706b565b615bce565b615bcc6170b0565b565b91615bde92949394919091619891565b615bf1615bec61012c613d58565b61139d565b9063e2051c7e91615c0361010b610fa5565b91615c0c61745c565b9490823b15615c82575f94615c3f8692615c3494615c286108da565b998a9889978896613d69565b865260048601613d97565b03925af18015615c7d57615c51575b50565b615c70905f3d8111615c76575b615c688183610d14565b810190613d6f565b5f615c4e565b503d615c5e565b613dd5565b613d65565b90615c9494939291615baf565b565b615ca790615ca2616d28565b615ca9565b565b615cba90615cb5616e10565b615cbc565b565b615cc881610118613952565b615cfe7f4e8c80fe79d13b8663de9f6981925ae24c498cc07b0ebd4f4789fa78ca17caff91615cf56108da565b91829182610b54565b0390a1565b615d0c90615c96565b565b615d1f90615d1a616d28565b615d21565b565b615d3290615d2d616e10565b615d34565b565b615d4081610135613952565b615d767f7acba37d1b2d934e554139ff1d470d2cce50a7b6e56870577249b87564a95a6991615d6d6108da565b91829182610b54565b0390a1565b615d8490615d0e565b565b615d8e6138c8565b50615da15f615d9b619ddd565b01612bd7565b90565b90615db79291615db261706b565b615dc1565b615dbf6170b0565b565b91615dce92919091619891565b565b90615ddb9291615da4565b565b615dee90615de9616d28565b615df0565b565b615e0190615dfc616e10565b615e0f565b565b615e0c90610f5e565b90565b615e2990615e24615e1f82615e03565b616fdc565b615e8a565b565b615e3490610f42565b90565b615e4090615e2b565b90565b615e4c90615e2b565b90565b90565b90615e67615e62615e6e92615e43565b615e4f565b8254613be2565b9055565b615e7b90610f42565b90565b615e8790615e72565b90565b615ea6615e9e615e9983615e03565b615e37565b610129615e52565b615ed07f9b3eda10f1724a2cf9f7dae4ac263c77908df4d00e92f1377b66fc8be37cd8c191615e7e565b90615ed96108da565b80615ee3816109cc565b0390a2565b615ef190615ddd565b565b615f0490615eff616d28565b615f06565b565b615f1790615f12616e10565b615f25565b565b615f2290610f5e565b90565b615f3f90615f3a615f3582615f19565b616fdc565b615fa0565b565b615f4a90610f42565b90565b615f5690615f41565b90565b615f6290615f41565b90565b90565b90615f7d615f78615f8492615f59565b615f65565b8254613be2565b9055565b615f9190610f42565b90565b615f9d90615f88565b90565b615fbc615fb4615faf83615f19565b615f4d565b61012a615f68565b615fe67fdab38e33e6e11cbb8b085bba9d7426d5e5af01bcc46d2c5957645e4d09e8c49c91615f94565b90615fef6108da565b80615ff9816109cc565b0390a2565b61600790615ef3565b565b616011613b78565b50616043616033616021306143ee565b3161602d61011e610fa5565b90613c85565b61603d6064613c96565b90613cc6565b90565b61605790616052616d28565b616059565b565b61606a90616065616e10565b61606c565b565b61607590619e01565b565b61608090616046565b565b61608a61706b565b61609261609c565b61609a6170b0565b565b6160a761010b610fa5565b6160af61745c565b34916160f96160e76160e17fe32cacf203d00685e2b4d8b0a90e7cd8f3f8a208fdf116f4bb36abe08b7d548e93610f10565b93610f6a565b936160f06108da565b91829182610b54565b0390a3565b616106616082565b565b61611990616114616d28565b61611b565b565b61612c90616127616e10565b61613a565b565b61613790610f5e565b90565b6161549061614f61614a8261612e565b616fdc565b6161b5565b565b61615f90610f42565b90565b61616b90616156565b90565b61617790616156565b90565b90565b9061619261618d6161999261616e565b61617a565b8254613be2565b9055565b6161a690610f42565b90565b6161b29061619d565b90565b6161d16161c96161c48361612e565b616162565b61012d61617d565b6161fb7fbf6e296f85d08cc1ab124aed644bf4b19e4a726a7aea53e3784ab1341738a040916161a9565b906162046108da565b8061620e816109cc565b0390a2565b61621c90616108565b565b6162309061622a613b78565b50619e48565b90565b61623b613b78565b506162446139c4565b5061624d619e66565b90616259610115610fa5565b9190565b906162729493929161626d61706b565b61627c565b61627a6170b0565b565b9161628c92949394919091619891565b61629f61629a61012c613d58565b61139d565b9063fe673fd3916162b161010b610fa5565b916162ba61745c565b9490823b15616330575f946162ed86926162e2946162d66108da565b998a9889978896613d69565b86526004860161466f565b03925af1801561632b576162ff575b50565b61631e905f3d8111616324575b6163168183610d14565b810190613d6f565b5f6162fc565b503d61630c565b613dd5565b613d65565b906163429493929161625d565b565b61634c613b78565b5061635e6163595f614778565b613b7c565b90565b616370616376919392936108ec565b926108ec565b820391821161638157565b61498a565b61638e6138c8565b90616397613b78565b906163a06138c8565b906163a9613b78565b906163b5610101612bd7565b6163cf6163c96163c45f6148fa565b61094a565b9161094a565b036163d7575b565b93505050506163e7610105612bd7565b906163f3610106610fa5565b6163fe610107610fa5565b9261640a610108610fa5565b92616416610109612bd7565b9261642261010a610fa5565b92616459600261645361644261010461643c61010b610fa5565b90610f2c565b61644d610101612bd7565b90610f76565b01610fa5565b96616465428990616361565b928261648161647b6164765f6148fa565b61094a565b9161094a565b145f146164ea575050506164ae906164a861649d610101612bd7565b9791965b429261574c565b90616361565b6164b7816139c8565b6164d16164cb6164c6866139c8565b610a13565b91610a13565b136164dd575b506163d5565b925090508391905f6164d7565b928098919792986165036164fd8a6108ec565b916108ec565b11616518575b5050906164a86164ae926164a1565b969261653482999361652e61653a94879061574c565b9261574c565b90616361565b90616544826139c8565b61655e616558616553886139c8565b610a13565b91610a13565b13616583575b50506164ae90946164a8616579610101612bd7565b9791969192616509565b9094506164ae9193509392905f616564565b6165a6906165a1616d28565b6165a8565b565b6165b9906165b4616e10565b6165bb565b565b6165c781610126613952565b6165fd7f4636d3e567b27988879babd22e50f49104ca65647933fc1623ff3d4d807438d2916165f46108da565b91829182610b54565b0390a1565b61660b90616595565b565b90616620929161661b61706b565b61662a565b6166286170b0565b565b9161663792919091617206565b565b90616644929161660d565b565b61665790616652616d28565b616659565b565b61666a90616665616e10565b61666c565b565b61667881610130613952565b6166ae7f2652e6657dd1ed89d7bdcb70f8827cc8449ac4536ebf604dbb2465cdad264020916166a56108da565b91829182610b54565b0390a1565b6166bc90616646565b565b6166cf906166ca616d28565b6166d1565b565b6166e2906166dd616e10565b6166e4565b565b6166ed90619e97565b565b6166f8906166be565b565b61670b90616706616d28565b61670d565b565b61671e90616719616e10565b616720565b565b61672c81610115613952565b6167627f4abea08c196329c357e3175d011af39a8625be99ef0ba5a0f3547a95534fedb7916167596108da565b91829182610b54565b0390a1565b616770906166fa565b565b61677a616d28565b616782616784565b565b61678c61678e565b565b6167966174e8565b61679e6167f9565b565b5f7f546f6f206561726c792e00000000000000000000000000000000000000000000910152565b6167d4600a602092612999565b6167dd816167a0565b0190565b6167f69060208101905f8183039101526167c7565b90565b616828616804616b8d565b919061682161681b61681685936139c8565b610a13565b91610a13565b1315614906565b616922576169209061691b616916616906616844610110610fa5565b61690061688361687261686261685b61010f610fa5565b859061572a565b61686c6001614466565b9061574c565b9261687d600261443a565b906156c8565b956168fa6168f46168e36168d26168b86168a86168a161010f610fa5565b8d9061572a565b6168b26001614466565b9061574c565b966168c1613b78565b506168cd61010f610fa5565b616361565b6168dd610125610fa5565b906156c8565b946168ef61010f610fa5565b616361565b916147b0565b906156c8565b9061572a565b6169106001614466565b9061574c565b619ede565b619e97565b565b61692a6108da565b63a29f5c4d60e01b815280616941600482016167e1565b0390fd5b61694d616772565b565b6169609061695b616d28565b616962565b565b6169739061696e616e10565b616975565b565b6169818161011c613952565b6169b77fd95e7f967f9370c11deb15ffbb191b9f2e9795ab0738db5bc72bd2794978f32d916169ae6108da565b91829182610b54565b0390a1565b6169c59061694f565b565b6169d2905f03610a13565b90565b6169dd6139c4565b506169ee6169e96139f2565b6169c7565b90565b616a02906169fd616d28565b616a04565b565b616a1590616a10616e10565b616a17565b565b616a2381610127613952565b616a597f37a332914fac995349420c0419b4423a19dcb762017f691442a0782ce4bf417a91616a506108da565b91829182610b54565b0390a1565b616a67906169f1565b565b616a7a90616a75616d28565b616a7c565b565b80616a97616a91616a8c5f6148fa565b61094a565b9161094a565b14616aa757616aa59061945d565b565b616aca616ab35f6148fa565b5f918291631e4fbdf760e01b835260048301610963565b0390fd5b616ad790616a69565b565b616aea90616ae5616d28565b616aec565b565b616afd90616af8616e10565b616aff565b565b616b0890619ede565b565b616b1390616ad9565b565b616b2690616b21616d28565b616b28565b565b616b3990616b34616e10565b616b3b565b565b616b4781610112613952565b616b7d7fdeb71e1d117914859ddde484a9810084d9ee399077d5cd8dcbdf8770d33d8ae491616b746108da565b91829182610b54565b0390a1565b616b8b90616b15565b565b616b95613b78565b50616b9e6139c4565b50616ba7619f25565b90616bb06139f2565b90565b90565b5f616bd7616bd2616bdd93616bc9613b78565b50610103611572565b616bb3565b01610fa5565b90565b90565b616c1291616c08616c0d92616bf6613b78565b50616bff613b78565b50610104610f2c565b610f76565b616be0565b90616c2a6001616c235f8501610fa5565b9301610fa5565b90565b616c3e90616c39616d28565b616c40565b565b616c5190616c4c616e10565b616c53565b565b616c5f8161011e613952565b616c957fbfcd8fb930a57c1598c9760db19c84ec766546c3c9a8565611df8302482bfb1791616c8c6108da565b91829182610b54565b0390a1565b616ca390616c2d565b565b616cad61706b565b616cb5616cf5565b616cbd6170b0565b565b90565b616cd6616cd1616cdb92616cbf565b610f0d565b610a13565b90565b616ce75f61291d565b90565b616cf2616cde565b90565b616d1c5f19616d045f91616cc2565b90616d16616d10616cea565b91614794565b91619891565b565b616d26616ca5565b565b616d30615d86565b616d49616d43616d3e61745c565b61094a565b9161094a565b03616d5057565b616d72616d5b61745c565b5f91829163118cdaa760e01b835260048301610963565b0390fd5b60207f65616479206163746976652e0000000000000000000000000000000000000000917f5468652063757272656e742062696464696e6720726f756e6420697320616c725f8201520152565b616dd0602c604092612999565b616dd981616d76565b0190565b916040616e0e929493616e07616dfc606083018381035f850152616dc3565b966020830190610b47565b0190610b47565b565b616e1b61010d610fa5565b616e3842616e31616e2b846108ec565b916108ec565b1015614906565b616e3f5750565b4290616e62616e4c6108da565b92839263d0fd11df60e01b845260048401616ddd565b0390fd5b90616e719101610a13565b90565b616eda90616e80613b78565b50616e8c610101612bd7565b616ea6616ea0616e9b5f6148fa565b61094a565b9161094a565b145f14616f4257616ed4616ece616ebe61010d610fa5565b5b92616ec9426139c8565b616e66565b916139c8565b906139e4565b616ee35f614794565b9080616eff616ef9616ef45f614778565b610a13565b91610a13565b13616f09575b5090565b616f3c9150616f1a616f2b916147b0565b616f2561011b610fa5565b90613c85565b616f36610125610fa5565b90613cc6565b5f616f05565b616ed4616ece616f7e6002616f78616f67610104616f6161010b610fa5565b90610f2c565b616f72610101612bd7565b90610f76565b01610fa5565b616ebf565b5f7f5468652070726f76696465642061646472657373206973207a65726f2e000000910152565b616fb7601d602092612999565b616fc081616f83565b0190565b616fd99060208101905f818303910152616faa565b90565b616ff6616ff0616feb5f6148fa565b61094a565b9161094a565b14616ffd57565b6170056108da565b63eac0d38960e01b81528061701c60048201616fc4565b0390fd5b90565b61703761703261703c92617020565b613934565b611bc4565b90565b6170687f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00617023565b90565b617073619f54565b6170945761709261708a61708561703f565b619f88565b600190619f9d565b565b5f633ee5aeb560e01b8152806170ac600482016109cc565b0390fd5b6170ca6170c36170be61703f565b619f88565b5f90619f9d565b565b60407f642e000000000000000000000000000000000000000000000000000000000000917f5468652063757272656e742043535420626964207072696365206973206772655f8201527f61746572207468616e20746865206d6178696d756d20796f7520616c6c6f776560208201520152565b61714c6042606092612999565b617155816170cc565b0190565b91604061718a929493617183617178606083018381035f85015261713f565b966020830190610b47565b0190610b47565b565b617196600261443a565b90565b6171a290616cc2565b9052565b91946171f36171e86171fd9360a0966171db6172049a9c9b999c6171d160c08a01945f8b0190617199565b6020890190610a16565b86820360408801526129ad565b986060850190610b47565b6080830190610b47565b0190610b47565b565b6172176172125f614778565b613b7c565b926172368461722e617228846108ec565b916108ec565b101515614906565b61743d575061724c6172475f614778565b61621e565b9061726b8261726361725d846108ec565b916108ec565b111515614906565b617417575061727b818490619fa0565b6172c2816172bc60016172ac61729e61010461729861010b610fa5565b90610f2c565b6172a661745c565b90610f76565b01916172b783610fa5565b61574c565b90613952565b6172ce42610114613952565b6172f46172e3826172dd61718c565b906156c8565b6172ee610118610fa5565b90619fac565b61730081610116613952565b61730b610102612bd7565b61732561731f61731a5f6148fa565b61094a565b9161094a565b14617405575b5061733f61733761745c565b610102613c00565b61736861734d610115610fa5565b6173628161735c610133610fa5565b9061572a565b9061574c565b61737481610115613952565b61737d8361a138565b61738861010b610fa5565b9061739161745c565b926174006173a05f19926139c8565b925f199697906173b1610124610fa5565b916173ee6173e86173e27f1d1f406c89e99504a7222aaba50bf96b9f91f522ebc6ea825255145919e102ec99610f10565b99610f6a565b99616cc2565b996173f76108da565b968796876171a6565b0390a4565b61741190610117613952565b5f61732b565b906174396174236108da565b92839263814ac7ff60e01b845260048401617159565b0390fd5b836174585f9283926335465b3160e01b84526004840161240c565b0390fd5b6174646138c8565b503390565b60207f207468652063757272656e742062696464696e6720726f756e642e0000000000917f41206269642068617320616c7265616479206265656e20706c6163656420696e5f8201520152565b6174c3603b604092612999565b6174cc81617469565b0190565b6174e59060208101905f8183039101526174b6565b90565b6175176174f6610101612bd7565b61751061750a6175055f6148fa565b61094a565b9161094a565b1415614906565b61751d57565b6175256108da565b634283f4b960e01b81528061753c600482016174d0565b0390fd5b61754c8161010d613952565b6175827f9a2159c1f277ddd727551baedc6a6c4cba77cc5219c8563ee3b15fb67548d89b916175796108da565b91829182610b54565b0390a1565b6175bd60026175b76175a66101046175a061010b610fa5565b90610f2c565b6175b1610101612bd7565b90610f76565b01610fa5565b6175c8428290616361565b6175d3610105612bd7565b6175ed6175e76175e25f6148fa565b61094a565b9161094a565b145f146176245761761961762192617611617609610101612bd7565b610105613c00565b610106613952565b610107613952565b5b565b8061764161763b617636610107610fa5565b6108ec565b916108ec565b1161764e575b5050617622565b6176a68261767261766d6176ae95617667610107610fa5565b9061574c565b6176b5565b617688617680610107610fa5565b610108613952565b61769e617696610101612bd7565b610105613c00565b610106613952565b610107613952565b5f80617647565b6176de906176d86176c7610106610fa5565b6176d2610108610fa5565b9061574c565b90616361565b6176e7816139c8565b61770b6177056177006176fb61010a610fa5565b6139c8565b610a13565b91610a13565b13617714575b50565b6177369061772e617726610105612bd7565b610109613c00565b61010a613952565b5f617711565b61774461a304565b565b6177506020610d3d565b90565b5f90565b61775f617746565b9060208261776b617753565b81525050565b617779617757565b90565b90617786906108ec565b9052565b67ffffffffffffffff81116177a25760208091020190565b610d00565b906177b96177b48361778a565b610d3d565b918252565b6177c86040610d3d565b90565b5f90565b6177d76177be565b90602080836177e46177cb565b8152016177ef617753565b81525050565b6177fd6177cf565b90565b5f5b82811061780e57505050565b6020906178196177f5565b8184015201617802565b90617848617830836177a7565b9260208061783e869361778a565b9201910390617800565b565b5190565b906178588261784a565b811015617869576020809102010190565b612b7a565b906178789061094a565b9052565b617885906108ec565b5f8114617893576001900390565b61498a565b6178a46178aa916108ec565b916108ec565b9081156178b5570690565b613cb2565b905051906178c7826108ef565b565b906020828203126178e2576178df915f016178ba565b90565b6108e4565b60209181520190565b60200190565b6178ff9061094a565b9052565b61790c906108ec565b9052565b90602080617932936179285f8201515f8601906178f6565b0151910190617903565b565b9061794181604093617910565b0190565b60200190565b9061796861796261795b8461784a565b80936178e7565b926178f0565b905f5b8181106179785750505090565b90919261799161798b6001928651617934565b94617945565b910191909161796b565b6179bf6179cc9492936179b560608401955f850190610b47565b6020830190610956565b604081840391015261794b565b90565b6179db6179e091610f8c565b61280b565b90565b6179ed90546179cf565b90565b5f9060033d116179fd575b565b905060045f803e617a0e5f516108d4565b906179fb565b5f5f9160233d11617a22575b565b915050602060045f3e6001905f5190617a20565b90565b617a4d617a48617a5292617a36565b610f0d565b6108ec565b90565b617a5f6012617a39565b90565b905090565b617a725f8092617a62565b0190565b617a7f90617a67565b90565b90617a94617a8f83611a71565b610d3d565b918252565b606090565b3d5f14617ab957617aae3d617a82565b903d5f602084013e5b565b617ac1617a99565b90617ab7565b5f7f455448207472616e7366657220746f2063686172697479206661696c65642e00910152565b617afb601f602092612999565b617b0481617ac7565b0190565b9190617b2b906020617b23604086018681035f880152617aee565b940190610b47565b565b60207f696e207072697a652062656e6566696369617279206661696c65642e00000000917f455448207472616e7366657220746f2062696464696e6720726f756e64206d615f8201520152565b617b87603c604092612999565b617b9081617b2d565b0190565b916040617bc5929493617bbe617bb3606083018381035f850152617b7a565b966020830190610956565b0190610b47565b565b60ff1690565b617be1617bdc617be692615442565b610f0d565b617bc7565b90565b90565b617c00617bfb617c0592617be9565b610f0d565b617bc7565b90565b617c1c617c17617c2192617bc7565b610f0d565b6108ec565b90565b617c30617c3591610f8c565b611468565b90565b617c429054617c24565b90565b617c4f90516108ec565b90565b90565b617c69617c64617c6e92617c52565b610f0d565b6108ec565b90565b67ffffffffffffffff8111617c895760208091020190565b610d00565b90505190617c9b82610bd3565b565b90929192617cb2617cad82617c71565b610d3d565b9381855260208086019202830192818411617cef57915b838310617cd65750505050565b60208091617ce48486617c8e565b815201920191617cc9565b6111c5565b9080601f83011215617d1257816020617d0f93519101617c9d565b90565b610cee565b90602082820312617d47575f82015167ffffffffffffffff8111617d4257617d3f9201617cf4565b90565b6108e8565b6108e4565b5190565b90617d62617d5d83617c71565b610d3d565b918252565b369037565b90617d91617d7983617d50565b92602080617d878693617c71565b9201910390617d67565b565b67ffffffffffffffff8111617dab5760208091020190565b610d00565b90617dc2617dbd83617d93565b610d3d565b918252565b617dd16040610d3d565b90565b617ddc617dc7565b9060208083617de96177cb565b815201617df4617753565b81525050565b617e02617dd4565b90565b5f5b828110617e1357505050565b602090617e1e617dfa565b8184015201617e07565b90617e4d617e3583617db0565b92602080617e438693617d93565b9201910390617e05565b565b5190565b90617e5d82617e4f565b811015617e6e576020809102010190565b612b7a565b90617e7d82617d4c565b811015617e8e576020809102010190565b612b7a565b617e9d905161094a565b90565b617eac617eb191610f8c565b613781565b90565b617ebe9054617ea0565b90565b60209181520190565b60200190565b90602080617ef293617ee85f8201515f8601906178f6565b0151910190617903565b565b90617f0181604093617ed0565b0190565b60200190565b90617f28617f22617f1b84617e4f565b8093617ec1565b92617eca565b905f5b818110617f385750505090565b909192617f51617f4b6001928651617ef4565b94617f05565b9101919091617f2b565b617f709160208201915f818403910152617f0b565b90565b617f7f617f8491610f8c565b611953565b90565b617f919054617f73565b90565b90565b617fab617fa6617fb092617f94565b610f0d565b6108ec565b90565b60209181520190565b60200190565b90617fcf816020936178f6565b0190565b60200190565b90617ff6617ff0617fe984617d4c565b8093617fb3565b92617fbc565b905f5b8181106180065750505090565b90919261801f6180196001928651617fc2565b94617fd3565b9101919091617ff9565b93929061805460409161805c9461804760608901925f8a0190610b47565b8782036020890152617fd9565b940190610b47565b565b61809361809a9461808960609498979561807f608086019a5f870190610b47565b6020850190610b47565b6040830190610b47565b0190610b47565b565b6180a590614906565b9052565b6040906180d26180d994969593966180c860608401985f85019061809c565b6020830190610b47565b0190610b47565b565b6180e3617771565b906180f76180ef61a3e9565b5f840161777c565b61811661811161010361810b61010b610fa5565b90611572565b616bb3565b9061811f613b78565b5061812861514b565b916181316143fa565b618139613ce8565b9261814261521b565b9061814b616009565b9561815761011f610fa5565b9361817461816f866181696001614466565b9061574c565b617823565b966181d76181c56181bd6181875f614794565b6181b66181958d8c9061784e565b516181ac6181a4610109612bd7565b5f830161786e565b602088910161777c565b859061574c565b9a889061572a565b996181d18b89906156c8565b9061574c565b60015b156182af575b5f966181eb9061787c565b9689886181f79161784e565b518c6182029061a4dd565b8a600101908b5f0161821390610fa5565b61821c91617898565b618225916138cc565b61822e90612bd7565b9081815f019061823d9161786e565b8c906020019061824c9161777c565b61825761010b610fa5565b8991908d927f9c62e2cb8cbd10bf3b8a5760977d719fc3475fa67cb7ef9d2e1701f275e885c49161828790610f10565b9261829190610f6a565b9361829a6108da565b9182916182a7918361240c565b0390a36181da565b866182c26182bc5f614794565b916108ec565b116181e05790939850602091949997929695506182e86182e361012c613d58565b61139d565b906183216387565d149192919261830061010b610fa5565b9361832c61830c61745c565b976183156108da565b98899788968795613d69565b85526004850161799b565b03925af1908115619071575f91619043575b509661835361834e61012e6179e3565b61284c565b9063b6b55f25909190919061836961010b610fa5565b90803b1561903e5761838e5f93618399956183826108da565b96879586948593613d69565b835260048301610b54565b03925af19081619012575b50155f1461900d5760016183b66179f0565b634e487b7114618fd0575b618fcb575b5f806183d3610131612bd7565b836183dc6108da565b90816183e781617a76565b03925af16183f3617a9e565b505f14618f7957618405610131612bd7565b6184446184327f1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d92610f6a565b9261843b6108da565b91829182610b54565b0390a25b61847c5f8061845561745c565b8661845e6108da565b908161846981617a76565b03925af1618475617a9e565b5015614906565b618f4b5761848b610102612bd7565b6184a561849f61849a5f6148fa565b61094a565b9161094a565b14155f14618f36576184d46184c36184bd6004617bec565b5b617c08565b6184ce610120610fa5565b9061574c565b926184e86184e361012d617c38565b6114a9565b5f63e36aee78916184fa610121610fa5565b9061854b618509848b01617c45565b6185327f7c6eeb003d4a6dc5ebf549935c6ffb814ba1f060f1af8a0b11c2aa94a8e716e4617c55565b189461855661853f6108da565b96879586948594613d69565b84526004840161240c565b03915afa8015618f315761857c915f91618f0f575b509461857686617d4c565b9061574c565b956185ad6185a8886185a2618592610139610fa5565b61859c6001614466565b90616361565b9061574c565b617d6c565b9761860c6185cd6185c88a6185c26001614466565b9061574c565b617e28565b986186076185dc8b8390617e53565b516185f36185eb61012f612bd7565b5f830161786e565b6020618600610130610fa5565b910161777c565b61787c565b9661864361861b8a8a90617e53565b5161862f61862761745c565b5f830161786e565b602061863c61011c610fa5565b910161777c565b61865561864f8b617d4c565b5b61787c565b61867261866061745c565b61866d8d91849092617e73565b61786e565b8061868561867f8b6108ec565b916108ec565b11156186945761865590618650565b5061869e87617d4c565b5b806186b26186ac5f614794565b916108ec565b1115618729576186c19061787c565b956187236186e16186db6186d68b8b90617e73565b617e93565b9a61787c565b996187116186f08d8d90617e53565b516186fd835f830161786e565b602061870a61011c610fa5565b910161777c565b61871e8d918c9092617e73565b61786e565b9561869f565b509297969093989591946187c8618741610120610fa5565b5b6187c361878161877b6187766187578a61a4dd565b618770600189019161876a5f8b01610fa5565b90617898565b906138cc565b612bd7565b9561787c565b946187b16187908d8890617e53565b5161879d835f830161786e565b60206187aa61011c610fa5565b910161777c565b6187be8b91879092617e73565b61786e565b61787c565b90816187dc6187d65f614794565b916108ec565b11156187fd57906187c361878161877b6187766187c8949350505050618742565b505061880b6188699161787c565b618844618819898390617e53565b51618830618828610109612bd7565b5f830161786e565b602061883d61011c610fa5565b910161777c565b618864618852610109612bd7565b61885f8991849092617e73565b61786e565b61787c565b6188a2618877888390617e53565b5161888e618886610105612bd7565b5f830161786e565b602061889b61011c610fa5565b910161777c565b6188c26188b0610105612bd7565b6188bd8891849092617e73565b61786e565b6188d46188ce5f614794565b916108ec565b115f14618f0a5761891c6188f1876188eb5f614794565b90617e53565b51618908618900610102612bd7565b5f830161786e565b602061891561011c610fa5565b910161777c565b61894261892a610102612bd7565b61893d876189375f614794565b90617e73565b61786e565b5b618956618951610129617eb4565b6137c2565b63b33266da87823b15618f055761898c926189815f80946189756108da565b96879586948593613d69565b835260048301617f5b565b03925af18015618f0057618ed4575b5060206189b16189ac61012b617f87565b611994565b636578f11390618a145f6189c661010b610fa5565b93618a1f6189d6838d9a01617c45565b6189ff7f2a8612ecb5cb17da87f8befda0480288e2d053de55d9d7d4dc4899077cf5aeda617f97565b18618a086108da565b998a9788968795613d69565b855260048501618029565b03925af18015618ecf57618b05925f91618ea1575b509794969792618a5c618a57618a4989617e4f565b618a51613b78565b5061787c565b61787c565b97618a68888a90617e53565b519989618a7687829061574c565b809c618afd618aae618a8961010b610fa5565b94618aa9618aa26020618a9a61745c565b999601617c45565b9598617d4c565b616361565b96618aeb618ae5618adf7f9314e78538382a9fc6cd54cee107c8a7721d172c43b2e48c2e7a51f200fc378897610f10565b97610f6a565b97610f10565b97618af46108da565b9485948561805e565b0390a4617d4c565b945b85618b1a618b145f614794565b916108ec565b1115618bd257618b299061787c565b92618b35858590617e53565b5195618b54618b4e618b485f8a01617e93565b9261787c565b9861787c565b96618b6061010b610fa5565b90600191618b7260208c959301617c45565b938a93618bc6618bb4618bae618ba87f27c21fe4cea1a3367aa491829dd4dd824296c00910626150464cba8ea5ebb3f496610f10565b96610f6a565b96610f10565b96618bbd6108da565b938493846180a9565b0390a492959495618b07565b9195909450618bec618be5610120610fa5565b925b61787c565b618bf7858290617e53565b5192618c16618c10618c0a5f8701617e93565b9261787c565b9361787c565b93618c2261010b610fa5565b905f91618c33602087959301617c45565b938793618c87618c75618c6f618c697f27c21fe4cea1a3367aa491829dd4dd824296c00910626150464cba8ea5ebb3f496610f10565b96610f6a565b96610f10565b96618c7e6108da565b938493846180a9565b0390a481618c9d618c975f614794565b916108ec565b1115618caf57618bec90929192618be7565b618d61929150618cbe9061787c565b90618cd4618ccd868490617e53565b519161787c565b93618ce061010b610fa5565b90618cec61011f610fa5565b91618d056020618cfd5f8701617e93565b939501617c45565b938793618d59618d47618d41618d3b7faa858ae20a26d00a9ea528972d537e68a51a0744226d9ea1fc9b22492dc282a596610f10565b96610f6a565b96610f10565b96618d506108da565b93849384610fee565b0390a461787c565b90618d77618d70848490617e53565b519161787c565b90618d8361010b610fa5565b90618d9b6020618d945f8401617e93565b9201617c45565b9291618dee618ddc618dd6618dd07f838ec9dd2530548892bff113f5ffb0138d2efc63c7f59bb6571e8c923b63126094610f10565b94610f6a565b94610f10565b94618de56108da565b91829182610b54565b0390a4618e03618dfd5f614794565b916108ec565b115f14618e9a57618e1d90618e175f614794565b90617e53565b51618e2961010b610fa5565b90618e416020618e3a5f8401617e93565b9201617c45565b9291618e94618e82618e7c618e767f3901b6430c99dc290ee88ff84c4de6091ad7eac335b58e92ef5cbb0793abf4f694610f10565b94610f6a565b94610f10565b94618e8b6108da565b91829182610b54565b0390a45b565b5050618e98565b618ec2915060203d8111618ec8575b618eba8183610d14565b8101906178c9565b5f618a34565b503d618eb0565b613dd5565b618ef3905f3d8111618ef9575b618eeb8183610d14565b810190613d6f565b5f61899b565b503d618ee1565b613dd5565b613d65565b618943565b618f2b91503d805f833e618f238183610d14565b810190617d17565b5f61856b565b613dd5565b6184d46184c3618f466003617bcd565b6184be565b82618f5461745c565b618f75618f5f6108da565b928392630aa7db6360e11b845260048401617b94565b0390fd5b618f84610131612bd7565b618fc3618fb17f154fb6c686c977314af35b730a16b83facda5265b2abec7237c730f63e42215a92610f6a565b92618fba6108da565b91829182617b08565b0390a2618448565b613dd5565b618fd8617a14565b90618fe4575b506183c1565b90505f9080619002618ffc618ff7617a55565b6108ec565b916108ec565b0315618fde5761a519565b6183c6565b619031905f3d8111619037575b6190298183610d14565b810190613d6f565b5f6183a4565b503d61901f565b613d65565b619064915060203d811161906a575b61905c8183610d14565b8101906178c9565b5f61833e565b503d619052565b613dd5565b600161908291016108ec565b90565b6190996190915f6148fa565b610101613c00565b6190ad6190a55f6148fa565b610102613c00565b6190c16190b95f6148fa565b610105613c00565b6190d56190cd5f614794565b610108613952565b6190e96190e15f6148fa565b610109613c00565b6191066190fe6190f95f19616cc2565b6147b0565b61010a613952565b61912461911c61911761010b610fa5565b619076565b61010b613952565b61915f61915a619135610125610fa5565b619154619143610125610fa5565b61914e610126610fa5565b90613cc6565b90614490565b619e01565b61917c6191774261917161010c610fa5565b90614490565b617540565b565b6191999061918a613b78565b506191948161a529565b61a696565b90565b6191a590610f5e565b90565b6191b13061919c565b6191e36191dd7f000000000000000000000000000000000000000000000000000000000000000061094a565b9161094a565b14801561920d575b6191f157565b5f63703e46dd60e11b815280619209600482016109cc565b0390fd5b5061921661a7e3565b6192486192427f000000000000000000000000000000000000000000000000000000000000000061094a565b9161094a565b14156191eb565b6192609061925b616d28565b619262565b565b5061926b616e10565b565b6192769061924f565b565b61928190610f42565b90565b61928d90619278565b90565b61929990610f5e565b90565b6192a581611bc4565b036192ac57565b5f80fd5b905051906192bd8261929c565b565b906020828203126192d8576192d5915f016192b0565b90565b6108e4565b919061930b60206192f56192f086619284565b619290565b6352d1902d906193036108da565b938492613d69565b8252818061931b600482016109cc565b03915afa80915f926193ab575b50155f1461935c57505090600161933d57505b565b619358905f918291634c9c8ce360e01b835260048301610963565b0390fd5b928361937761937161936c614f72565b611bc4565b91611bc4565b0361938c5761938792935061a809565b61933b565b6193a7845f918291632a87526960e21b835260048301611bd4565b0390fd5b6193cd91925060203d81116193d4575b6193c58183610d14565b8101906192bf565b905f619328565b503d6193bb565b6193e43061919c565b6194166194107f000000000000000000000000000000000000000000000000000000000000000061094a565b9161094a565b0361941d57565b5f63703e46dd60e11b815280619435600482016109cc565b0390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0090565b619465619ddd565b61947d6194735f8301612bd7565b915f849101613c00565b906194b16194ab7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093610f6a565b91610f6a565b916194ba6108da565b806194c4816109cc565b0390a3565b6194d86194de91939293610a13565b92610a13565b91828103925f8285128183121692851391121516176194f957565b61498a565b60407f727265642e000000000000000000000000000000000000000000000000000000917f5468652063757272656e742045544820626964207072696365206973206772655f8201527f61746572207468616e2074686520616d6f756e7420796f75207472616e73666560208201520152565b61957e6045606092612999565b619587816194fe565b0190565b9160406195bc9294936195b56195aa606083018381035f850152619571565b966020830190610b47565b0190610b47565b565b60207f206265656e207573656420666f722062696464696e672e000000000000000000917f546869732052616e646f6d2057616c6b204e46542068617320616c72656164795f8201520152565b6196186037604092612999565b619621816195be565b0190565b9190619648906020619640604086018681035f88015261960b565b940190610b47565b565b156196525750565b6196749061965e6108da565b91829163c35947c560e01b835260048301619625565b0390fd5b61968461968991610f8c565b61217f565b90565b6196969054619678565b90565b906020828203126196b2576196af915f01617c8e565b90565b6108e4565b60207f6e646f6d2057616c6b204e46542e000000000000000000000000000000000000917f596f7520617265206e6f7420746865206f776e6572206f6620746869732052615f8201520152565b619711602e604092612999565b61971a816196b7565b0190565b60609061975761975e949695939661974d619742608085018581035f870152619704565b9860208501906121cc565b6040830190610b47565b0190610956565b565b929091921561976e57505050565b6197909061977a6108da565b938493630b81342760e31b85526004850161971e565b0390fd5b61979e600261443a565b90565b6197b56197b06197ba92610a13565b610f0d565b610a13565b90565b919461980a6197ff6198149360a0966197f261981b9a9c9b999c6197e860c08a01945f8b0190610a16565b6020890190617199565b86820360408801526129ad565b986060850190610b47565b6080830190610b47565b0190610b47565b565b5f7f45544820726566756e64207472616e73666572206661696c65642e0000000000910152565b619851601b602092612999565b61985a8161981d565b0190565b91604061988f92949361988861987d606083018381035f850152619844565b966020830190610956565b0190610b47565b565b91906198a461989f5f614778565b613b7c565b916198c3836198bb6198b5846108ec565b916108ec565b101515614906565b619dbe57506198d96198d45f614778565b614dfd565b92806198f56198ef6198ea5f614778565b610a13565b91610a13565b125f14619db057835b9061991a61990b346139c8565b619914846139c8565b906194c9565b948561993661993061992b5f614778565b610a13565b91610a13565b145f14619d15575b8161995961995361994e5f614778565b610a13565b91610a13565b125f14619bc957619a076199f7619a0f925b6199b1866199ab5f61999b61998d61010461998761010b610fa5565b90610f2c565b61999561745c565b90610f76565b01916199a683610fa5565b61574c565b90613952565b6199bc610101612bd7565b6199d66199d06199cb5f6148fa565b61094a565b9161094a565b14619ba7575b6199f1816199eb610112610fa5565b9061572a565b9061574c565b619a016001614466565b9061574c565b610111613952565b619a68619a44619a33619a23610115610fa5565b619a2d6001614466565b9061574c565b619a3e610133610fa5565b906156c8565b619a62619a52610133610fa5565b619a5c6001614466565b9061574c565b9061572a565b90619a7582610115613952565b619a7e8561a892565b619a878461a138565b619a9261010b610fa5565b91619b08619aa7619aa161745c565b956139c8565b915f1993969790619ab9610124610fa5565b91619af6619af0619aea7f1d1f406c89e99504a7222aaba50bf96b9f91f522ebc6ea825255145919e102ec99610f10565b99610f6a565b996197a1565b99619aff6108da565b968796876197bd565b0390a480619b26619b20619b1b5f614778565b610a13565b91610a13565b13619b2f575b50565b619b6b5f80619b3c61745c565b619b45856147b0565b619b4d6108da565b9081619b5881617a76565b03925af1619b64617a9e565b5015614906565b15619b2c57619b81619b7b61745c565b916147b0565b90619ba3619b8d6108da565b928392630aa7db6360e11b84526004840161985e565b0390fd5b619bc4619bbc82619bb6619794565b906156c8565b61010f613952565b6199dc565b619c5f90619c0f619bed619be8610119619be2876147b0565b9061349e565b610fa5565b619bff619bf95f614794565b916108ec565b14619c09856147b0565b9061964a565b619c1761745c565b906020619c2d619c2861012a61968c565b6121c0565b636352211e90619c54619c3f886147b0565b92619c486108da565b97889485938493613d69565b835260048301610b54565b03915afa908115619d1057619cb6619a0793619c92619c8c6199f795619a0f985f91619ce2575b5061094a565b9161094a565b14619c9e61012a61968c565b619ca7886147b0565b90619cb061745c565b92619760565b619cdd619cc36001614466565b619cd8610119619cd2896147b0565b9061349e565b613952565b61996b565b619d03915060203d8111619d09575b619cfb8183610d14565b810190619699565b5f619c86565b503d619cf1565b613dd5565b85619d30619d2a619d255f614778565b610a13565b91610a13565b135f14619d8857619d4c619d45610113610fa5565b3a906156c8565b619d67619d61619d5b896147b0565b926108ec565b916108ec565b1115619d73575b61993e565b91509350619d805f614778565b933491619d6e565b823490619dac619d966108da565b92839263814ac7ff60e01b84526004840161958b565b0390fd5b619db98461449e565b6198fe565b82619dd95f9283926335465b3160e01b84526004840161240c565b0390fd5b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930090565b619e0d81610125613952565b619e437f07417920574ce0bdfe987af0575c8793cc73a29d7830760ad459d0e569b5b79b91619e3a6108da565b91829182610b54565b0390a1565b619e6390619e54613b78565b50619e5e8161a89d565b61a696565b90565b619e6e6139c4565b50619e94619e7b426139c8565b619e8e619e89610114610fa5565b6139c8565b906139e4565b90565b619ea381610110613952565b619ed97fb6f6af60099e44041a78b3561ed029b98bf03fdb0efbbb2eb15e1f3d7d92303791619ed06108da565b91829182610b54565b0390a1565b619eea8161010e613952565b619f207ffdf6043c89a5f304289122dcc9f8bd78bb111b5d4f409e2fc2e6c141a1110b7991619f176108da565b91829182610b54565b0390a1565b619f2d613b78565b50619f4d619f3c610125610fa5565b619f4761010e610fa5565b90613cc6565b90565b5f90565b619f5c619f50565b50619f75619f70619f6b61703f565b619f88565b61a95a565b90565b5f90565b619f8590611bc4565b90565b619f9a90619f94619f78565b50619f7c565b90565b5d565b90619faa9161ab3b565b565b619fd591619fb8613b78565b5081619fcc619fc6836108ec565b916108ec565b1191909161ada7565b90565b90565b5190565b5f7f4d65737361676520697320746f6f206c6f6e672e000000000000000000000000910152565b61a0136014602092612999565b61a01c81619fdf565b0190565b919061a04390602061a03b604086018681035f88015261a006565b940190610b47565b565b1561a04d5750565b61a06f9061a0596108da565b91829163271d43ff60e21b83526004830161a020565b0390fd5b60207f207368616c6c206265204554482e000000000000000000000000000000000000917f5468652066697273742062696420696e20612062696464696e6720726f756e645f8201520152565b61a0cd602e604092612999565b61a0d68161a073565b0190565b61a0ef9060208101905f81830391015261a0c0565b90565b1561a0f957565b61a1016108da565b63b5a45a4960e01b81528061a1186004820161a0da565b0390fd5b61a125906108ec565b5f19811461a1335760010190565b61498a565b61a1829061a17c61a17761a15361a14e84619fd8565b619fdb565b61a16f61a16961a16461011a610fa5565b6108ec565b916108ec565b111592619fd8565b619fdb565b9061a045565b61a18d610101612bd7565b61a1a761a1a161a19c5f6148fa565b61094a565b9161094a565b145f1461a2ef5761a1b661ae8d565b61a1d23461a1cc61a1c65f614794565b916108ec565b1161a0f2565b61a1de42610114613952565b61a1fb61a1f34261a1ed6144e0565b9061574c565b610124613952565b61a20661010b610fa5565b429061a24761a2357f028a52641badd593b7f30072734c0b97e449213f55b5c3663756427340accd3c92610f10565b9261a23e6108da565b91829182610b54565b0390a25b61a25e61a25661745c565b610101613c00565b61a2b961a28061a27b61010361a27561010b610fa5565b90611572565b616bb3565b5f61a2b261a28f828401610fa5565b61a2ad61a29a61745c565b61a2a86001870184906138cc565b613c00565b61a11c565b9101613952565b61a2ed42600261a2e761a2d961010461a2d361010b610fa5565b90610f2c565b61a2e161745c565b90610f76565b01613952565b565b61a2f7617587565b61a2ff61adcb565b61a24b565b61a30c61aee7565b565b61a31a61a31f91610f8c565b610f10565b90565b61a33661a33161a33b92614463565b610f0d565b617bc7565b90565b61a35d9061a35761a35161a36294617bc7565b916108ec565b90610c7c565b6108ec565b90565b90565b61a37c61a37761a3819261a365565b610f0d565b617bc7565b90565b61a3a39061a39d61a39761a3a894617bc7565b916108ec565b90614124565b6108ec565b90565b90565b61a3c261a3bd61a3c79261a3ab565b610f0d565b617bc7565b90565b90565b61a3e161a3dc61a3e69261a3ca565b610f0d565b617bc7565b90565b61a3f1613b78565b5061a42261a41261a40c4361a4066001614466565b90616361565b4061a30e565b61a41c600161a322565b9061a33e565b61a4364861a430604061a368565b9061a384565b1861a43f61b034565b9061a49d575b5061a44e61b3aa565b9061a482575b5061a45d61b544565b9061a467575b5090565b61a47b9061a47560c061a3cd565b9061a384565b185f61a463565b61a4969061a490608061a3ae565b9061a384565b185f61a454565b61a4bd61a4c29161a4ac614f3b565b5061a4b76001614466565b90616361565b61b1a8565b9061a4cd575b61a445565b61a4d69061a30e565b185f61a4c8565b61a5115f61a5169261a4ed613b78565b5061a50b82820161a50561a50082617c45565b619076565b9061777c565b01617c45565b61b660565b90565b634e487b715f526020526024601cfd5b61a531613b78565b5061a53a613b78565b5061a546610101612bd7565b61a56061a55a61a5555f6148fa565b61094a565b9161094a565b145f1461a63e5761a58361a57561010f610fa5565b9161a57e6139f2565b616e66565b8061a59e61a59861a5935f614778565b610a13565b91610a13565b13155f1461a5ac57505b5b90565b9061a5d561a5c58261a5bf610110610fa5565b90613cc6565b61a5cf6001614466565b90614490565b9161a5de619f25565b9061a5e8816147b0565b61a5fa61a5f4846108ec565b916108ec565b105f1461a636579061a62561a62a9261a61f61a61961a6309787614482565b916147b0565b90613c85565b613cc6565b90614482565b5b61a5a8565b50505061a631565b5061a64a610111610fa5565b61a5a9565b9061a66561a65f61a66c936108ec565b916108ec565b900a6108ec565b90565b61a68e9061a68861a68261a693946108ec565b916108ec565b90610c7c565b6108ec565b90565b91909161a6a1613b78565b509161a6ae610101612bd7565b61a6c861a6c261a6bd5f6148fa565b61094a565b9161094a565b0361a6d1575b50565b61a6ea61a6dc615258565b9161a6e561481b565b6139e4565b9061a6fe61a6f7826139c8565b83906139e4565b918261a71a61a71461a70f5f614778565b610a13565b91610a13565b1361a727575b505061a6ce565b61a7a561a79461a78361a77261a7d497989661a7ce9661a7ac9661a75b61a75561a7505f614778565b610a13565b91610a13565b125f1461a7dd5761a76c91506139c8565b5b6147b0565b61a77d610136610fa5565b90613c85565b61a78e610125610fa5565b90613cc6565b61a79f610137610fa5565b9061a64f565b8390613c85565b61a7c861a7ba610137610fa5565b61a7c26157e1565b90613c85565b9061a66f565b90614490565b905f808061a720565b5061a76d565b61a7eb6138c8565b5061a8065f61a80061a7fb614f72565b61b675565b01612bd7565b90565b9061a8138261b678565b8161a83e7fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b91610f6a565b9061a8476108da565b8061a851816109cc565b0390a261a85d81619fdb565b61a86f61a8695f614794565b916108ec565b115f1461a8835761a87f9161b703565b505b565b505061a88d61b6cd565b61a881565b61a89b9061b732565b565b61a8b79061a8a9613b78565b5061a8b261b965565b6139e4565b8061a8d261a8cc61a8c75f614778565b610a13565b91610a13565b131561a94d5761a92661a9379161a8ea610102612bd7565b61a90461a8fe61a8f95f6148fa565b61094a565b9161094a565b145f1461a93a5761a92061a919610117610fa5565b5b916147b0565b90613c85565b61a931610115610fa5565b90613cc6565b90565b61a92061a948610116610fa5565b61a91a565b5061a9575f614794565b90565b61a962619f50565b505c90565b91602061a98892949361a98160408201965f830190610956565b0190610b47565b565b67ffffffffffffffff811161a9a25760208091020190565b610d00565b9061a9b961a9b48361a98a565b610d3d565b918252565b61a9c86040610d3d565b90565b5f90565b61a9d761a9be565b906020808361a9e46177cb565b81520161a9ef61a9cb565b81525050565b61a9fd61a9cf565b90565b5f5b82811061aa0e57505050565b60209061aa1961a9f5565b818401520161aa02565b9061aa4861aa308361a9a7565b9260208061aa3e869361a98a565b920191039061aa00565b565b5190565b9061aa588261aa4a565b81101561aa69576020809102010190565b612b7a565b9061aa7890610a13565b9052565b60209181520190565b60200190565b61aa9490610a13565b9052565b9060208061aaba9361aab05f8201515f8601906178f6565b015191019061aa8b565b565b9061aac98160409361aa98565b0190565b60200190565b9061aaf061aaea61aae38461aa4a565b809361aa7c565b9261aa85565b905f5b81811061ab005750505090565b90919261ab1961ab13600192865161aabc565b9461aacd565b910191909161aaf3565b61ab389160208201915f81840391015261aad3565b90565b908061ab4f61ab495f614794565b916108ec565b115f1461ad085761ac588161ac2961ac0c61ac0761ab6f61ac759661b994565b9361abd961abbd61abb861ab8b61ab866003615445565b61aa23565b9a61abb38c5f61abac61ab9c61745c565b9261aba683614794565b9061aa4e565b510161786e565b6139c8565b6169c7565b602061abd28b61abcc5f614794565b9061aa4e565b510161aa6e565b61ac0061abe461745c565b5f61abf98b61abf36001614466565b9061aa4e565b510161786e565b8490614482565b6139c8565b602061ac228761ac1c6001614466565b9061aa4e565b510161aa6e565b61ac5361ac37610101612bd7565b5f61ac4c8761ac46600261443a565b9061aa4e565b510161786e565b6139c8565b602061ac6e8461ac68600261443a565b9061aa4e565b510161aa6e565b61ac8861ac83610129617eb4565b6137c2565b9063b355121490823b1561ad035761acbf9261acb45f809461aca86108da565b96879586948593613d69565b83526004830161ab23565b03925af1801561acfe5761acd2575b505b565b61acf1905f3d811161acf7575b61ace98183610d14565b810190613d6f565b5f61acce565b503d61acdf565b613dd5565b613d65565b5061ad1c61ad17610129617eb4565b6137c2565b90639dc29fac9061ad2b61745c565b9092803b1561ada25761ad515f809461ad5c61ad456108da565b97889687958694613d69565b84526004840161a967565b03925af1801561ad9d5761ad71575b5061acd0565b61ad90905f3d811161ad96575b61ad888183610d14565b810190613d6f565b5f61ad6b565b503d61ad7e565b613dd5565b613d65565b61adc161adc7929361adb7613b78565b508094189161b9c8565b90613c85565b1890565b61adf161ade961add9614e9b565b61ade4610124610fa5565b614490565b610124613952565b565b60207f20616374697665207965742e0000000000000000000000000000000000000000917f5468652063757272656e742062696464696e6720726f756e64206973206e6f745f8201520152565b61ae4d602c604092612999565b61ae568161adf3565b0190565b91604061ae8b92949361ae8461ae79606083018381035f85015261ae40565b966020830190610b47565b0190610b47565b565b61ae9861010d610fa5565b61aeb64261aeae61aea8846108ec565b916108ec565b101515614906565b61aebd5750565b429061aee061aeca6108da565b9283926302dbf17b60e31b84526004840161ae5a565b0390fd5b90565b61af3361af0961af0461013461aefe61010b610fa5565b906123cf565b61aee4565b61af1f61af17610107610fa5565b5f8301613952565b600161af2c61010a610fa5565b9101613952565b565b61af4961af4461af4e92613c93565b610f0d565b61093f565b90565b61af5a9061af35565b90565b61af6690610f42565b90565b61af729061af5d565b90565b61af8761af82606461af51565b61af69565b90565b61af9390610f5e565b90565b90565b61afad61afa861afb29261af96565b610f0d565b6108ec565b90565b60207f642e000000000000000000000000000000000000000000000000000000000000917f4172625379732e617262426c6f636b4e756d6265722063616c6c206661696c655f8201520152565b61b00f6022604092612999565b61b0188161afb5565b0190565b61b0319060208101905f81830391015261b002565b90565b61b03c619f50565b5061b045613b78565b61b04d617a99565b505f8061b06061b05b61af75565b61af8a565b600461b09763a3b1b31d60e01b61b08861b0786108da565b93849260208401908152016109cc565b60208201810382520382610d14565b82602082019151925af19161b0aa617a9e565b8361b0fd575b5061b0bb8315614906565b61b0c2575b565b7fa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d61b0eb6108da565b8061b0f58161b01c565b0390a161b0c0565b909161b10882619fdb565b61b11b61b115602061af99565b916108ec565b145f1461b145575061b13d90602061b13282619fdb565b8183010191016178c9565b905b5f61b0b0565b919250505f9161b13f565b5f7f4172625379732e617262426c6f636b486173682063616c6c206661696c65642e910152565b61b18360208092612999565b61b18c8161b150565b0190565b61b1a59060208101905f81830391015261b177565b90565b61b1b0619f50565b505f8061b1bb614f3b565b9261b1c4617a99565b50600461b20e61b1da61b1d561af75565b61af8a565b9261b1ff6315a03d4160e11b9161b1ef6108da565b9485936020850190815201610b54565b60208201810382520382610d14565b82602082019151925af19161b221617a9e565b8361b274575b5061b2328315614906565b61b239575b565b7fa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d61b2626108da565b8061b26c8161b190565b0390a161b237565b909161b27f82619fdb565b61b29261b28c602061af99565b916108ec565b145f1461b2bc575061b2b490602061b2a982619fdb565b8183010191016192bf565b905b5f61b227565b919250505f9161b2b6565b90565b61b2de61b2d961b2e39261b2c7565b610f0d565b61093f565b90565b61b2ef9061b2ca565b90565b61b2fb90610f42565b90565b61b3079061b2f2565b90565b61b31c61b317606c61b2e6565b61b2fe565b90565b61b32890610f5e565b90565b60207f696c65642e000000000000000000000000000000000000000000000000000000917f417262476173496e666f2e6765744761734261636b6c6f672063616c6c2066615f8201520152565b61b3856025604092612999565b61b38e8161b32b565b0190565b61b3a79060208101905f81830391015261b378565b90565b61b3b2619f50565b5061b3bb613b78565b61b3c3617a99565b505f8061b3d661b3d161b30a565b61b31f565b600461b40c62eadae160e51b61b3fd61b3ed6108da565b93849260208401908152016109cc565b60208201810382520382610d14565b82602082019151925af19161b41f617a9e565b8361b472575b5061b4308315614906565b61b437575b565b7fa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d61b4606108da565b8061b46a8161b392565b0390a161b435565b909161b47d82619fdb565b61b49061b48a602061af99565b916108ec565b145f1461b4ba575061b4b290602061b4a782619fdb565b8183010191016178c9565b905b5f61b425565b919250505f9161b4b4565b60207f655570646174652063616c6c206661696c65642e000000000000000000000000917f417262476173496e666f2e6765744c3150726963696e67556e69747353696e635f8201520152565b61b51f6034604092612999565b61b5288161b4c5565b0190565b61b5419060208101905f81830391015261b512565b90565b61b54c619f50565b5061b555613b78565b61b55d617a99565b505f8061b57061b56b61b30a565b61b31f565b600461b5a76377f8098360e11b61b59861b5886108da565b93849260208401908152016109cc565b60208201810382520382610d14565b82602082019151925af19161b5ba617a9e565b8361b60d575b5061b5cb8315614906565b61b5d2575b565b7fa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d61b5fb6108da565b8061b6058161b52c565b0390a161b5d0565b909161b61882619fdb565b61b62b61b625602061af99565b916108ec565b145f1461b655575061b64d90602061b64282619fdb565b8183010191016178c9565b905b5f61b5c0565b919250505f9161b64f565b61b6729061b66c613b78565b5061b9d6565b90565b90565b803b61b68c61b6865f614794565b916108ec565b1461b6ae5761b6ac905f61b6a661b6a1614f72565b61b675565b01613c00565b565b61b6c9905f918291634c9c8ce360e01b835260048301610963565b0390fd5b3461b6e061b6da5f614794565b916108ec565b1161b6e757565b5f63b398979f60e01b81528061b6ff600482016109cc565b0390fd5b5f8061b72f9361b711617a99565b508390602081019051915af49061b726617a9e565b9091909161b9e8565b90565b8061b74561b73f5f614794565b916108ec565b1161b74e575b50565b61b7578161b994565b61b762610101612bd7565b8061b77d61b77761b7725f6148fa565b61094a565b9161094a565b145f1461b82d575061b79861b793610129617eb4565b6137c2565b9161b7b06340c10f199261b7aa61745c565b92614482565b92803b1561b8285761b7d55f809461b7e061b7c96108da565b97889687958694613d69565b84526004840161a967565b03925af1801561b8235761b7f7575b505b5f61b74b565b61b816905f3d811161b81c575b61b80e8183610d14565b810190613d6f565b5f61b7ef565b503d61b804565b613dd5565b613d65565b9061b8b361b8d09261b89761b87b61b84d61b848600261443a565b617e28565b9661b87461b85961745c565b5f61b86d8b61b86783614794565b90617e53565b510161786e565b8590614482565b602061b8908861b88a5f614794565b90617e53565b510161777c565b5f61b8ac8661b8a66001614466565b90617e53565b510161786e565b602061b8c98461b8c36001614466565b90617e53565b510161777c565b61b8e361b8de610129617eb4565b6137c2565b9063b33266da90823b1561b9605761b91a9261b90f5f809461b9036108da565b96879586948593613d69565b835260048301617f5b565b03925af1801561b95b5761b92f575b5061b7f1565b61b94e905f3d811161b954575b61b9468183610d14565b810190613d6f565b5f61b929565b503d61b93c565b613dd5565b613d65565b61b96d6139c4565b5061b99161b979619e66565b61b98c61b987610115610fa5565b6139c8565b6139e4565b90565b61b9b561b9c59161b9a3613b78565b5061b9af610138610fa5565b90613c85565b61b9bf6064613c96565b90613cc6565b90565b61b9d0613b78565b50151590565b61b9de613b78565b505f5260205f2090565b9061b9fc9061b9f5617a99565b5015614906565b5f1461ba08575061ba6c565b61ba1182619fdb565b61ba2361ba1d5f614794565b916108ec565b148061ba51575b61ba32575090565b61ba4d905f918291639996b31560e01b835260048301610963565b0390fd5b50803b61ba6661ba605f614794565b916108ec565b1461ba2a565b61ba7581619fdb565b61ba8761ba815f614794565b916108ec565b115f1461ba9657805190602001fd5b5f63d6bda27560e01b81528061baae600482016109cc565b0390fdfea2646970667358221220454848d27760530cb0e9e3db293e909bbe4fb8a1e14fecca507ea7e42bee8cc964736f6c63430008220033",
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

// ChampionDurations is a free data retrieval call binding the contract method 0x8af9f8b6.
//
// Solidity: function championDurations(uint256 roundNum) view returns(uint256 enduranceChampion, uint256 chronoWarrior)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) ChampionDurations(opts *bind.CallOpts, roundNum *big.Int) (struct {
	EnduranceChampion *big.Int
	ChronoWarrior     *big.Int
}, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "championDurations", roundNum)

	outstruct := new(struct {
		EnduranceChampion *big.Int
		ChronoWarrior     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EnduranceChampion = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ChronoWarrior = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ChampionDurations is a free data retrieval call binding the contract method 0x8af9f8b6.
//
// Solidity: function championDurations(uint256 roundNum) view returns(uint256 enduranceChampion, uint256 chronoWarrior)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) ChampionDurations(roundNum *big.Int) (struct {
	EnduranceChampion *big.Int
	ChronoWarrior     *big.Int
}, error) {
	return _CosmicSignatureGameV3.Contract.ChampionDurations(&_CosmicSignatureGameV3.CallOpts, roundNum)
}

// ChampionDurations is a free data retrieval call binding the contract method 0x8af9f8b6.
//
// Solidity: function championDurations(uint256 roundNum) view returns(uint256 enduranceChampion, uint256 chronoWarrior)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) ChampionDurations(roundNum *big.Int) (struct {
	EnduranceChampion *big.Int
	ChronoWarrior     *big.Int
}, error) {
	return _CosmicSignatureGameV3.Contract.ChampionDurations(&_CosmicSignatureGameV3.CallOpts, roundNum)
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

// LastBidderBidCstRewardAmountPercentage is a free data retrieval call binding the contract method 0x9f9a7fcf.
//
// Solidity: function lastBidderBidCstRewardAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) LastBidderBidCstRewardAmountPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "lastBidderBidCstRewardAmountPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBidderBidCstRewardAmountPercentage is a free data retrieval call binding the contract method 0x9f9a7fcf.
//
// Solidity: function lastBidderBidCstRewardAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) LastBidderBidCstRewardAmountPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.LastBidderBidCstRewardAmountPercentage(&_CosmicSignatureGameV3.CallOpts)
}

// LastBidderBidCstRewardAmountPercentage is a free data retrieval call binding the contract method 0x9f9a7fcf.
//
// Solidity: function lastBidderBidCstRewardAmountPercentage() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) LastBidderBidCstRewardAmountPercentage() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.LastBidderBidCstRewardAmountPercentage(&_CosmicSignatureGameV3.CallOpts)
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

// SetLastBidderBidCstRewardAmountPercentage is a paid mutator transaction binding the contract method 0x003d9695.
//
// Solidity: function setLastBidderBidCstRewardAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Transactor) SetLastBidderBidCstRewardAmountPercentage(opts *bind.TransactOpts, newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.contract.Transact(opts, "setLastBidderBidCstRewardAmountPercentage", newValue_)
}

// SetLastBidderBidCstRewardAmountPercentage is a paid mutator transaction binding the contract method 0x003d9695.
//
// Solidity: function setLastBidderBidCstRewardAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) SetLastBidderBidCstRewardAmountPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetLastBidderBidCstRewardAmountPercentage(&_CosmicSignatureGameV3.TransactOpts, newValue_)
}

// SetLastBidderBidCstRewardAmountPercentage is a paid mutator transaction binding the contract method 0x003d9695.
//
// Solidity: function setLastBidderBidCstRewardAmountPercentage(uint256 newValue_) returns()
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3TransactorSession) SetLastBidderBidCstRewardAmountPercentage(newValue_ *big.Int) (*types.Transaction, error) {
	return _CosmicSignatureGameV3.Contract.SetLastBidderBidCstRewardAmountPercentage(&_CosmicSignatureGameV3.TransactOpts, newValue_)
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

// CosmicSignatureGameV3LastBidderBidCstRewardAmountPercentageChangedIterator is returned from FilterLastBidderBidCstRewardAmountPercentageChanged and is used to iterate over the raw logs and unpacked data for LastBidderBidCstRewardAmountPercentageChanged events raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3LastBidderBidCstRewardAmountPercentageChangedIterator struct {
	Event *CosmicSignatureGameV3LastBidderBidCstRewardAmountPercentageChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CosmicSignatureGameV3LastBidderBidCstRewardAmountPercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmicSignatureGameV3LastBidderBidCstRewardAmountPercentageChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CosmicSignatureGameV3LastBidderBidCstRewardAmountPercentageChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CosmicSignatureGameV3LastBidderBidCstRewardAmountPercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmicSignatureGameV3LastBidderBidCstRewardAmountPercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmicSignatureGameV3LastBidderBidCstRewardAmountPercentageChanged represents a LastBidderBidCstRewardAmountPercentageChanged event raised by the CosmicSignatureGameV3 contract.
type CosmicSignatureGameV3LastBidderBidCstRewardAmountPercentageChanged struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLastBidderBidCstRewardAmountPercentageChanged is a free log retrieval operation binding the contract event 0xc63013cf34a6f7b20983b293d1787e833f8de2db868e904525fc2910df652a97.
//
// Solidity: event LastBidderBidCstRewardAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) FilterLastBidderBidCstRewardAmountPercentageChanged(opts *bind.FilterOpts) (*CosmicSignatureGameV3LastBidderBidCstRewardAmountPercentageChangedIterator, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.FilterLogs(opts, "LastBidderBidCstRewardAmountPercentageChanged")
	if err != nil {
		return nil, err
	}
	return &CosmicSignatureGameV3LastBidderBidCstRewardAmountPercentageChangedIterator{contract: _CosmicSignatureGameV3.contract, event: "LastBidderBidCstRewardAmountPercentageChanged", logs: logs, sub: sub}, nil
}

// WatchLastBidderBidCstRewardAmountPercentageChanged is a free log subscription operation binding the contract event 0xc63013cf34a6f7b20983b293d1787e833f8de2db868e904525fc2910df652a97.
//
// Solidity: event LastBidderBidCstRewardAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) WatchLastBidderBidCstRewardAmountPercentageChanged(opts *bind.WatchOpts, sink chan<- *CosmicSignatureGameV3LastBidderBidCstRewardAmountPercentageChanged) (event.Subscription, error) {

	logs, sub, err := _CosmicSignatureGameV3.contract.WatchLogs(opts, "LastBidderBidCstRewardAmountPercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmicSignatureGameV3LastBidderBidCstRewardAmountPercentageChanged)
				if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "LastBidderBidCstRewardAmountPercentageChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLastBidderBidCstRewardAmountPercentageChanged is a log parse operation binding the contract event 0xc63013cf34a6f7b20983b293d1787e833f8de2db868e904525fc2910df652a97.
//
// Solidity: event LastBidderBidCstRewardAmountPercentageChanged(uint256 newValue)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Filterer) ParseLastBidderBidCstRewardAmountPercentageChanged(log types.Log) (*CosmicSignatureGameV3LastBidderBidCstRewardAmountPercentageChanged, error) {
	event := new(CosmicSignatureGameV3LastBidderBidCstRewardAmountPercentageChanged)
	if err := _CosmicSignatureGameV3.contract.UnpackLog(event, "LastBidderBidCstRewardAmountPercentageChanged", log); err != nil {
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
