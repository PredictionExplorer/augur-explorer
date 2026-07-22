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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit\",\"type\":\"uint256\"}],\"name\":\"BidCstRewardAmountMinLimitNotReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"BidHasBeenPlacedInCurrentRound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"}],\"name\":\"CallerIsNotNftOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientReceivedBidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"InvalidOperationInCurrentState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"lastBidderAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"durationUntilOperationIsPermitted\",\"type\":\"uint256\"}],\"name\":\"MainPrizeClaimDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"mainPrizeTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"MainPrizeEarlyClaim\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"NoBidsPlacedInCurrentRound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"roundActivationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"RoundIsActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"roundActivationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"RoundIsInactive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"messageLength\",\"type\":\"uint256\"}],\"name\":\"TooLongBidMessage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"randomWalkNftId\",\"type\":\"uint256\"}],\"name\":\"UsedRandomWalkNft\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"WrongBidType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ArbitrumError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"BidCstRewardAmountMultiplierChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"BidMessageLengthMaxLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lastBidderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"paidEthPrice\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"paidCstPrice\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"randomWalkNftId\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstDutchAuctionDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mainPrizeTime\",\"type\":\"uint256\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CharityAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CharityEthDonationAmountPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"ChronoWarriorEthPrizeAmountPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chronoWarriorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeCosmicSignatureNftId\",\"type\":\"uint256\"}],\"name\":\"ChronoWarriorPrizePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICosmicSignatureNft\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CosmicSignatureNftAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractICosmicSignatureToken\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CosmicSignatureTokenAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CstDutchAuctionBeginningBidPriceMinLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CstDutchAuctionDurationChangeDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CstDutchAuctionDurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CstPrizeAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"DelayDurationBeforeRoundActivationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"enduranceChampionAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeCosmicSignatureNftId\",\"type\":\"uint256\"}],\"name\":\"EnduranceChampionPrizePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"EthBidPriceIncreaseDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"EthBidRefundAmountInGasToSwallowMaxLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthDonated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ethDonationWithInfoRecordIndex\",\"type\":\"uint256\"}],\"name\":\"EthDonatedWithInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"EthDutchAuctionDurationDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"EthDutchAuctionEndingBidPriceDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockTimeStamp\",\"type\":\"uint256\"}],\"name\":\"FirstBidPlacedInRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"charityAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsTransferredToCharity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"InitialDurationUntilMainPrizeDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"LastBidderBidCstRewardAmountPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lastCstBidderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeCosmicSignatureNftId\",\"type\":\"uint256\"}],\"name\":\"LastCstBidderPrizePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MainEthPrizeAmountPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiaryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeFirstCosmicSignatureNftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prizeNumCosmicSignatureNfts\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeoutTimeToWithdrawSecondaryPrizes\",\"type\":\"uint256\"}],\"name\":\"MainPrizeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MainPrizeNumCosmicSignatureNftsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MainPrizeTimeIncrementInMicroSecondsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MainPrizeTimeIncrementIncreaseDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"MarketingWalletAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MarketingWalletCstContributionAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"NumRaffleCosmicSignatureNftsForBiddersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"NumRaffleEthPrizesForBiddersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIPrizesWallet\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"PrizesWalletAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RaffleTotalEthPrizeAmountForBiddersPercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winnerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethPrizeAmount\",\"type\":\"uint256\"}],\"name\":\"RaffleWinnerBidderEthPrizeAllocated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"winnerIsRandomWalkNftStaker\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winnerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cstPrizeAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prizeCosmicSignatureNftId\",\"type\":\"uint256\"}],\"name\":\"RaffleWinnerPrizePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIRandomWalkNFT\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"RandomWalkNftAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RoundActivationTimeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RoundLateBidDurationDivisorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RoundLateBidPricePremiumAmountBaseMultiplierChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RoundLateBidPricePremiumAmountExponentChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIStakingWalletCosmicSignatureNft\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"StakingWalletCosmicSignatureNftAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIStakingWalletRandomWalkNft\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"StakingWalletRandomWalkNftAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"TimeoutDurationToClaimMainPrizeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bidCstRewardAmountMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bidMessageLengthMaxLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceMaxLimit_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"}],\"name\":\"bidWithCst\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceMaxLimit_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"bidWithCstAndDonateNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceMaxLimit_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"bidWithCstAndDonateToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"randomWalkNftId_\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"}],\"name\":\"bidWithEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"randomWalkNftId_\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC721\",\"name\":\"nftAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"bidWithEthAndDonateNft\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"randomWalkNftId_\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidCstRewardAmountMinLimit_\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"bidWithEthAndDonateToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"}],\"name\":\"bidderAddresses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"numItems\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidderAddress\",\"type\":\"address\"}],\"name\":\"biddersInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalSpentEthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSpentCstAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBidTimeStamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"}],\"name\":\"championDurations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"enduranceChampion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chronoWarrior\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charityAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"charityEthDonationAmountPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chronoWarriorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chronoWarriorDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chronoWarriorEthPrizeAmountPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimMainPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cosmicSignatureNftStakingTotalEthRewardAmountPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionBeginningBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionBeginningBidPriceMinLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionBeginningTimeStamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstDutchAuctionDurationChangeDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cstPrizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayDurationBeforeRoundActivation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donateEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"data_\",\"type\":\"string\"}],\"name\":\"donateEthWithInfo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enduranceChampionAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enduranceChampionDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enduranceChampionStartTimeStamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethBidPriceIncreaseDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethBidRefundAmountInGasToSwallowMaxLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ethDonationWithInfoRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethDutchAuctionBeginningBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethDutchAuctionDurationDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethDutchAuctionEndingBidPriceDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBidCstRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"currentTimeOffset_\",\"type\":\"int256\"}],\"name\":\"getBidCstRewardAmountAdvanced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBidCstRewardAmountPerMainPrizeTimeIncrement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidIndex_\",\"type\":\"uint256\"}],\"name\":\"getBidderAddressAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidderAddress_\",\"type\":\"address\"}],\"name\":\"getBidderTotalSpentAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCharityEthDonationAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChronoWarriorEthPrizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCosmicSignatureNftStakingTotalEthRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCstDutchAuctionBeginningBidPriceMinLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCstDutchAuctionDurations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDurationElapsedSinceRoundActivation\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDurationUntilMainPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDurationUntilMainPrizeRaw\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDurationUntilRoundActivation\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEthDutchAuctionDurations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethBidPrice_\",\"type\":\"uint256\"}],\"name\":\"getEthPlusRandomWalkNftBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitialDurationUntilMainPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMainEthPrizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMainPrizeTimeIncrement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextCstBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"currentTimeOffset_\",\"type\":\"int256\"}],\"name\":\"getNextCstBidPriceAdvanced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextEthBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"currentTimeOffset_\",\"type\":\"int256\"}],\"name\":\"getNextEthBidPriceAdvanced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRaffleTotalEthPrizeAmountForBidders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoundLateBidDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"}],\"name\":\"getTotalNumBids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"halveEthDutchAuctionEndingBidPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialDurationUntilMainPrizeDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBidderAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBidderBidCstRewardAmountPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCstBidderAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainEthPrizeAmountPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainPrizeNumCosmicSignatureNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainPrizeTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainPrizeTimeIncrementInMicroSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainPrizeTimeIncrementIncreaseDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketingWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketingWalletCstContributionAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextEthBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRoundFirstCstDutchAuctionBeginningBidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractCosmicSignatureNft\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numEthDonationWithInfoRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleCosmicSignatureNftsForBidders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleCosmicSignatureNftsForRandomWalkNftStakers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRaffleEthPrizesForBidders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prevEnduranceChampionDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizesWallet\",\"outputs\":[{\"internalType\":\"contractPrizesWallet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raffleTotalEthPrizeAmountForBiddersPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomWalkNft\",\"outputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reinitialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundActivationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundLateBidDurationDivisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundLateBidPricePremiumAmountBaseMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundLateBidPricePremiumAmountExponent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setBidCstRewardAmountMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setBidMessageLengthMaxLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setCharityAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCharityEthDonationAmountPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setChronoWarriorEthPrizeAmountPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICosmicSignatureNft\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setCosmicSignatureNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCosmicSignatureNftStakingTotalEthRewardAmountPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICosmicSignatureToken\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setCosmicSignatureToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCstDutchAuctionBeginningBidPriceMinLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCstDutchAuctionDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCstDutchAuctionDurationChangeDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setCstPrizeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setDelayDurationBeforeRoundActivation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setEthBidPriceIncreaseDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setEthBidRefundAmountInGasToSwallowMaxLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setEthDutchAuctionDurationDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setEthDutchAuctionEndingBidPriceDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setInitialDurationUntilMainPrizeDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setLastBidderBidCstRewardAmountPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setMainEthPrizeAmountPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setMainPrizeNumCosmicSignatureNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setMainPrizeTimeIncrementInMicroSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setMainPrizeTimeIncrementIncreaseDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setMarketingWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setMarketingWalletCstContributionAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleCosmicSignatureNftsForBidders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleCosmicSignatureNftsForRandomWalkNftStakers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setNumRaffleEthPrizesForBidders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPrizesWallet\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setPrizesWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setRaffleTotalEthPrizeAmountForBiddersPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRandomWalkNFT\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setRandomWalkNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setRoundActivationTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setRoundLateBidDurationDivisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setRoundLateBidPricePremiumAmountBaseMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setRoundLateBidPricePremiumAmountExponent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakingWalletCosmicSignatureNft\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setStakingWalletCosmicSignatureNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakingWalletRandomWalkNft\",\"name\":\"newValue_\",\"type\":\"address\"}],\"name\":\"setStakingWalletRandomWalkNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue_\",\"type\":\"uint256\"}],\"name\":\"setTimeoutDurationToClaimMainPrize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingWalletCosmicSignatureNft\",\"outputs\":[{\"internalType\":\"contractStakingWalletCosmicSignatureNft\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingWalletRandomWalkNft\",\"outputs\":[{\"internalType\":\"contractStakingWalletRandomWalkNft\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeoutDurationToClaimMainPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractCosmicSignatureToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tryGetCurrentChampions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"enduranceChampionAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"enduranceChampionDuration_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"chronoWarriorAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chronoWarriorDuration_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"usedRandomWalkNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nftWasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523461003e5761001161004d565b610019610043565b61be8e6103ed82396080518181816192190152818161927e015261944c015261be8e90f35b610049565b60405190565b5f80fd5b610055610057565b565b61005f610061565b565b61006961006b565b565b610073610075565b565b61007d61007f565b565b610087610089565b565b610091610093565b565b61009b61009d565b565b6100a56100a7565b565b6100af6100b1565b565b6100b96100bb565b565b6100c36100c5565b565b6100cd6100cf565b565b6100d76100e1565b6100df610310565b565b6100e96100eb565b565b6100f36100f5565b565b6100fd6100ff565b565b610107610109565b565b610111610113565b565b61011b61011d565b565b610125610127565b565b61012f610131565b565b61013961013b565b565b610143610145565b565b61014d61014f565b565b610157610159565b565b610161610163565b565b61016b61016d565b565b610175610177565b565b61017f610181565b565b61018961018b565b565b610193610195565b565b61019d61019f565b565b6101a76101a9565b565b6101b16101b3565b565b6101bb6101bd565b565b6101c56101c7565b565b6101cf610213565b565b60018060a01b031690565b90565b6101f36101ee6101f8926101d1565b6101dc565b6101d1565b90565b610204906101df565b90565b610210906101fb565b90565b61021c30610207565b608052565b60401c90565b60ff1690565b61023961023e91610221565b610227565b90565b61024b905461022d565b90565b5f0190565b5f1c90565b60018060401b031690565b61026f61027491610253565b610258565b90565b6102819054610263565b90565b60018060401b031690565b5f1b90565b906102a560018060401b039161028f565b9181191691161790565b6102c36102be6102c892610284565b6101dc565b610284565b90565b90565b906102e36102de6102ea926102af565b6102cb565b8254610294565b9055565b6102f790610284565b9052565b919061030e905f602085019401906102ee565b565b6103186103c8565b6103235f8201610241565b6103ac576103325f8201610277565b61034a61034460018060401b03610284565b91610284565b03610353575b50565b610366905f60018060401b0391016102ce565b60018060401b036103a37fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29161039a610043565b918291826102fb565b0390a15f610350565b5f63f92ee8a960e01b8152806103c46004820161024e565b0390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009056fe6080604052600436101561001d575b3661394e5761001b616dce565b005b6100275f356108f4565b80620ac9f1146108ef5780623d9695146108ea578063040d4d31146108e557806304338479146108e057806309632366146108db57806309794bee146108d65780630a120648146108d15780630b5f95ae146108cc5780630c9be46d146108c75780630eb16be6146108c2578063119b22b3146108bd57806311b0d1fe146108b8578063135f3d28146108b357806317887731146108ae5780631824d5e7146108a957806318305de2146108a45780631aaba5a51461089f5780631b4103191461089a5780631e9cbb7e146108955780631f1b4aa41461089057806323b31cfc1461088b578063250fadb6146108865780632665c88214610881578063277004811461087c57806327995f07146108775780632afa2580146108725780632b8dcbba1461086d5780632b91c7bb146108685780632d809e88146108635780632d829a2d1461085e5780632f894cd7146108595780632fb3c48f14610854578063320c435c1461084f578063329b95a51461084a57806336750d2c1461084557806337b99cc7146108405780633b9d292e1461083b5780634164b95b14610836578063441b328914610831578063448c6eb11461082c57806344a4b9171461082757806344acc12a14610822578063477adf2a1461081d57806347ccca02146108185780634c2a4a33146108135780634e4520101461080e5780634f1ef286146108095780634f734612146108045780634feb78b7146107ff57806352d1902d146107fa578063543f416f146107f557806354ada1d6146107f057806356732241146107eb5780635863a705146107e65780635a1e5bde146107e15780635b0a45d9146107dc5780635d098b38146107d75780635f0112fe146107d25780635fdf49cb146107cd57806360ef8841146107c857806362ed9b53146107c35780636b59acb8146107be5780636b7cbe85146107b95780636c0613c0146107b45780636c17e3cc146107af5780636c2eb350146107aa5780636e95d286146107a55780636e970834146107a0578063715018a61461079b57806371b6d01914610796578063755b4ef71461079157806375ef3b9c1461078c57806375f0a8741461078757806377fa10271461078257806387292a851461077d578063876d5c361461077857806388ce802c146107735780638af9f8b61461076e5780638c94e9ba146107695780638da5cb5b14610764578063928880fa1461075f5780639302020f1461075a5780639646d7581461075557806399bf353d146107505780639aa1b38d1461074b5780639e50acc9146107465780639edeaf8e146107415780639f9a7fcf1461073c578063a35286d114610737578063a375f48214610732578063a4be0d401461072d578063a922ab5d14610728578063a974201614610723578063aadd1b031461071e578063ad3cb1cc14610719578063ad4b0e8a14610714578063afcf2fc41461070f578063b30f5bb11461070a578063b4f1b13414610705578063b5d1f06f14610700578063b6a94f42146106fb578063b700db5f146106f6578063b78d1e2a146106f1578063b9cf9ba5146106ec578063baab4430146106e7578063bb4b3e6f146106e2578063be720ad5146106dd578063c52d8549146106d8578063c7e7a601146106d3578063c87baab5146106ce578063cb720d4d146106c9578063cfb4e599146106c4578063d1f8fcf2146106bf578063d7559b9c146106ba578063d9ab9eaa146106b5578063da9931dd146106b0578063dd5f6587146106ab578063ddd6df07146106a6578063de704b41146106a1578063dfcd00d11461069c578063e2f9185f14610697578063e5b3cd1414610692578063eaace3021461068d578063eb13430e14610688578063ebaa1ea814610683578063ebb9bc5c1461067e578063ecb5776e14610679578063ef22d15b14610674578063efeb248a1461066f578063f0bdab7c1461066a578063f11400f014610665578063f2fde38b14610660578063f34d411c1461065b578063f444b29814610656578063f49efe9d14610651578063f7bea0781461064c578063fbaf508414610647578063fc0c546a14610642578063fd77310f1461063d578063fd9b3747146106385763fdfb9ba40361000e5761391b565b6138e4565b6138af565b61387a565b6137d5565b6137a0565b61375d565b613728565b6136e5565b6136b2565b61367d565b61363a565b613605565b6135c0565b61358d565b613558565b6134f3565b6134ae565b613469565b613424565b6133df565b61339c565b613369565b613336565b613301565b6132cc565b613289565b613255565b6131cd565b61318a565b613151565b6130de565b613099565b613054565b61300f565b612fca565b612f85565b612f50565b612f13565b612e6f565b612e17565b612dde565b612b9c565b612b1a565b612ad5565b612a90565b612a4b565b612949565b612914565b6128cf565b61282d565b6127f8565b6127c3565b61278e565b61274b565b6126c9565b612684565b61263f565b6125fc565b61257a565b612540565b6124b8565b612485565b61244f565b6123bc565b61238f565b6122ee565b6122bb565b612286565b612243565b61220e565b61216a565b612127565b6120f2565b6120ad565b61207a565b612047565b611fc5565b611f80565b611f0c565b611ed7565b611ea2565b611e6d565b611e28565b611df5565b611dc0565b611d8d565b611d0b565b611cc6565b611c81565b611c3e565b611c09565b611bb1565b611b7c565b611b52565b611a5c565b611a27565b6119e2565b61193e565b6118fb565b611879565b611836565b611803565b6117ce565b611789565b611744565b61170f565b6116d8565b611606565b6115c1565b61155f565b61152c565b6114f7565b611455565b611420565b6113eb565b611347565b611312565b6112cd565b611298565b61125e565b6111ae565b61117b565b611146565b611103565b6110ce565b611089565b611040565b610ecd565b610e96565b610cd9565b610c67565b610c34565b610bbe565b610b89565b610af3565b610ac0565b610a8d565b610a58565b6109f1565b610998565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b90565b6109188161090c565b0361091f57565b5f80fd5b905035906109308261090f565b565b919060408382031261095a578061094e610957925f8601610923565b93602001610923565b90565b610904565b60018060a01b031690565b6109739061095f565b90565b61097f9061096a565b9052565b9190610996905f60208501940190610976565b565b346109c9576109c56109b46109ae366004610932565b9061396c565b6109bc6108fa565b91829182610983565b0390f35b610900565b906020828203126109e7576109e4915f01610923565b90565b610904565b5f0190565b34610a1f57610a09610a043660046109ce565b613a43565b610a116108fa565b80610a1b816109ec565b0390f35b610900565b5f910312610a2e57565b610904565b90565b610a3f90610a33565b9052565b9190610a56905f60208501940190610a36565b565b34610a8857610a68366004610a24565b610a84610a73613a7c565b610a7b6108fa565b91829182610a43565b0390f35b610900565b34610abb57610aa5610aa03660046109ce565b613b1a565b610aad6108fa565b80610ab7816109ec565b0390f35b610900565b34610aee57610ad8610ad33660046109ce565b613b92565b610ae06108fa565b80610aea816109ec565b0390f35b610900565b34610b2157610b0b610b063660046109ce565b613bf7565b610b136108fa565b80610b1d816109ec565b0390f35b610900565b610b2f81610a33565b03610b3657565b5f80fd5b90503590610b4782610b26565b565b90602082820312610b6257610b5f915f01610b3a565b90565b610904565b610b709061090c565b9052565b9190610b87905f60208501940190610b67565b565b34610bb957610bb5610ba4610b9f366004610b49565b613c06565b610bac6108fa565b91829182610b74565b0390f35b610900565b34610bee57610bce366004610a24565b610bea610bd9613c1b565b610be16108fa565b91829182610b74565b0390f35b610900565b610bfc8161096a565b03610c0357565b5f80fd5b90503590610c1482610bf3565b565b90602082820312610c2f57610c2c915f01610c07565b90565b610904565b34610c6257610c4c610c47366004610c16565b613cf8565b610c546108fa565b80610c5e816109ec565b0390f35b610900565b34610c9757610c77366004610a24565b610c93610c82613d72565b610c8a6108fa565b91829182610b74565b0390f35b610900565b1c90565b90565b610cb3906008610cb89302610c9c565b610ca0565b90565b90610cc69154610ca3565b90565b610cd661010b5f90610cbb565b90565b34610d0957610ce9366004610a24565b610d05610cf4610cc9565b610cfc6108fa565b91829182610b74565b0390f35b610900565b5f80fd5b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b90610d3e90610d16565b810190811067ffffffffffffffff821117610d5857604052565b610d20565b90610d70610d696108fa565b9283610d34565b565b67ffffffffffffffff8111610d9057610d8c602091610d16565b0190565b610d20565b90825f939282370152565b90929192610db5610db082610d72565b610d5d565b93818552602085019082840111610dd157610dcf92610d95565b565b610d12565b9080601f83011215610df457816020610df193359101610da0565b90565b610d0e565b610e029061096a565b90565b610e0e81610df9565b03610e1557565b5f80fd5b90503590610e2682610e05565b565b919060a083820312610e9157610e40815f8501610923565b92602081013567ffffffffffffffff8111610e8c5782610e61918301610dd6565b92610e89610e728460408501610923565b93610e808160608601610e19565b93608001610923565b90565b610908565b610904565b34610ec857610eb2610ea9366004610e28565b93929092613f28565b610eba6108fa565b80610ec4816109ec565b0390f35b610900565b34610efb57610ee5610ee03660046109ce565b613fa4565b610eed6108fa565b80610ef7816109ec565b0390f35b610900565b9190604083820312610f285780610f1c610f25925f8601610923565b93602001610c07565b90565b610904565b90565b610f44610f3f610f499261090c565b610f2d565b61090c565b90565b90610f5690610f30565b5f5260205260405f2090565b610f76610f71610f7b9261095f565b610f2d565b61095f565b90565b610f8790610f62565b90565b610f9390610f7e565b90565b90610fa090610f8a565b5f5260205260405f2090565b5f1c90565b610fbd610fc291610fac565b610ca0565b90565b610fcf9054610fb1565b90565b90610fe2610fe792610104610f4c565b610f96565b610ff25f8201610fc5565b9161100b600261100460018501610fc5565b9301610fc5565b90565b60409061103761103e949695939661102d60608401985f850190610b67565b6020830190610b67565b0190610b67565b565b346110745761107061105c611056366004610f00565b90610fd2565b6110679391936108fa565b9384938461100e565b0390f35b610900565b61108661010a5f90610cbb565b90565b346110b957611099366004610a24565b6110b56110a4611079565b6110ac6108fa565b91829182610b74565b0390f35b610900565b6110cb6101245f90610cbb565b90565b346110fe576110de366004610a24565b6110fa6110e96110be565b6110f16108fa565b91829182610b74565b0390f35b610900565b346111315761111b6111163660046109ce565b61401c565b6111236108fa565b8061112d816109ec565b0390f35b610900565b6111436101165f90610cbb565b90565b3461117657611156366004610a24565b611172611161611136565b6111696108fa565b91829182610b74565b0390f35b610900565b346111a95761119361118e3660046109ce565b614094565b61119b6108fa565b806111a5816109ec565b0390f35b610900565b346111dc576111c66111c13660046109ce565b61410c565b6111ce6108fa565b806111d8816109ec565b0390f35b610900565b5f80fd5b5f80fd5b909182601f830112156112235781359167ffffffffffffffff831161121e57602001926001830284011161121957565b6111e5565b6111e1565b610d0e565b90602082820312611259575f82013567ffffffffffffffff81116112545761125092016111e9565b9091565b610908565b610904565b61127261126c366004611228565b9061446c565b61127a6108fa565b80611284816109ec565b0390f35b6112956101065f90610cbb565b90565b346112c8576112a8366004610a24565b6112c46112b3611288565b6112bb6108fa565b91829182610b74565b0390f35b610900565b346112fd576112dd366004610a24565b6112f96112e8614484565b6112f06108fa565b91829182610b74565b0390f35b610900565b61130f6101145f90610cbb565b90565b3461134257611322366004610a24565b61133e61132d611302565b6113356108fa565b91829182610b74565b0390f35b610900565b346113775761137361136261135d3660046109ce565b614528565b61136a6108fa565b91829182610b74565b0390f35b610900565b60018060a01b031690565b61139790600861139c9302610c9c565b61137c565b90565b906113aa9154611387565b90565b6113ba61012c5f9061139f565b90565b6113c690610f7e565b90565b6113d2906113bd565b9052565b91906113e9905f602085019401906113c9565b565b3461141b576113fb366004610a24565b6114176114066113ad565b61140e6108fa565b918291826113d6565b0390f35b610900565b3461145057611430366004610a24565b61144c61143b61456a565b6114436108fa565b91829182610b74565b0390f35b610900565b346114835761146d6114683660046109ce565b6145c6565b6114756108fa565b8061147f816109ec565b0390f35b610900565b60018060a01b031690565b6114a39060086114a89302610c9c565b611488565b90565b906114b69154611493565b90565b6114c661012d5f906114ab565b90565b6114d290610f7e565b90565b6114de906114c9565b9052565b91906114f5905f602085019401906114d5565b565b3461152757611507366004610a24565b6115236115126114b9565b61151a6108fa565b918291826114e2565b0390f35b610900565b3461155a5761154461153f3660046109ce565b61463e565b61154c6108fa565b80611556816109ec565b0390f35b610900565b3461158d576115776115723660046109ce565b6146b6565b61157f6108fa565b80611589816109ec565b0390f35b610900565b9061159c90610f30565b5f5260205260405f2090565b5f6115b86115be92610103611592565b01610fc5565b90565b346115f1576115ed6115dc6115d73660046109ce565b6115a8565b6115e46108fa565b91829182610b74565b0390f35b610900565b61160361011c5f90610cbb565b90565b3461163657611616366004610a24565b6116326116216115f6565b6116296108fa565b91829182610b74565b0390f35b610900565b6116449061096a565b90565b6116508161163b565b0361165757565b5f80fd5b9050359061166882611647565b565b919060a0838203126116d357611682815f8501610923565b92602081013567ffffffffffffffff81116116ce57826116a3918301610dd6565b926116cb6116b48460408501610923565b936116c2816060860161165b565b93608001610923565b90565b610908565b610904565b3461170a576116f46116eb36600461166a565b939290926147f0565b6116fc6108fa565b80611706816109ec565b0390f35b610900565b3461173f5761171f366004610a24565b61173b61172a614856565b6117326108fa565b91829182610b74565b0390f35b610900565b3461177457611754366004610a24565b61177061175f6148a5565b6117676108fa565b91829182610a43565b0390f35b610900565b6117866101275f90610cbb565b90565b346117b957611799366004610a24565b6117b56117a4611779565b6117ac6108fa565b91829182610b74565b0390f35b610900565b6117cb6101305f90610cbb565b90565b346117fe576117de366004610a24565b6117fa6117e96117be565b6117f16108fa565b91829182610b74565b0390f35b610900565b346118315761181b6118163660046109ce565b614943565b6118236108fa565b8061182d816109ec565b0390f35b610900565b3461186457611846366004610a24565b61184e614d67565b6118566108fa565b80611860816109ec565b0390f35b610900565b6118766101235f90610cbb565b90565b346118a957611889366004610a24565b6118a5611894611869565b61189c6108fa565b91829182610b74565b0390f35b610900565b6118b79061096a565b90565b6118c3816118ae565b036118ca57565b5f80fd5b905035906118db826118ba565b565b906020828203126118f6576118f3915f016118ce565b90565b610904565b346119295761191361190e3660046118dd565b614e7c565b61191b6108fa565b80611925816109ec565b0390f35b610900565b61193b61011e5f90610cbb565b90565b3461196e5761194e366004610a24565b61196a61195961192e565b6119616108fa565b91829182610b74565b0390f35b610900565b60018060a01b031690565b61198e9060086119939302610c9c565b611973565b90565b906119a1915461197e565b90565b6119b161012b5f90611996565b90565b6119bd90610f7e565b90565b6119c9906119b4565b9052565b91906119e0905f602085019401906119c0565b565b34611a12576119f2366004610a24565b611a0e6119fd6119a4565b611a056108fa565b918291826119cd565b0390f35b610900565b611a246101205f90610cbb565b90565b34611a5757611a37366004610a24565b611a53611a42611a17565b611a4a6108fa565b91829182610b74565b0390f35b610900565b34611a8c57611a88611a77611a72366004610b49565b614e87565b611a7f6108fa565b91829182610b74565b0390f35b610900565b67ffffffffffffffff8111611aaf57611aab602091610d16565b0190565b610d20565b90929192611ac9611ac482611a91565b610d5d565b93818552602085019082840111611ae557611ae392610d95565b565b610d12565b9080601f83011215611b0857816020611b0593359101611ab4565b90565b610d0e565b919091604081840312611b4d57611b26835f8301610c07565b92602082013567ffffffffffffffff8111611b4857611b459201611aea565b90565b610908565b610904565b611b66611b60366004611b0d565b90614ec5565b611b6e6108fa565b80611b78816109ec565b0390f35b34611bac57611b8c366004610a24565b611ba8611b97614f25565b611b9f6108fa565b91829182610b74565b0390f35b610900565b34611bdf57611bc9611bc43660046109ce565b614fba565b611bd16108fa565b80611bdb816109ec565b0390f35b610900565b90565b611bf090611be4565b9052565b9190611c07905f60208501940190611be7565b565b34611c3957611c19366004610a24565b611c35611c24615034565b611c2c6108fa565b91829182611bf4565b0390f35b610900565b34611c6c57611c56611c513660046109ce565b6150b4565b611c5e6108fa565b80611c68816109ec565b0390f35b610900565b611c7e61011d5f90610cbb565b90565b34611cb157611c91366004610a24565b611cad611c9c611c71565b611ca46108fa565b91829182610b74565b0390f35b610900565b611cc36101265f90610cbb565b90565b34611cf657611cd6366004610a24565b611cf2611ce1611cb6565b611ce96108fa565b91829182610b74565b0390f35b610900565b611d086101075f90610cbb565b90565b34611d3b57611d1b366004610a24565b611d37611d26611cfb565b611d2e6108fa565b91829182610b74565b0390f35b610900565b611d499061096a565b90565b611d5581611d40565b03611d5c57565b5f80fd5b90503590611d6d82611d4c565b565b90602082820312611d8857611d85915f01611d60565b90565b610904565b34611dbb57611da5611da0366004611d6f565b6151ca565b611dad6108fa565b80611db7816109ec565b0390f35b610900565b34611df057611dd0366004610a24565b611dec611ddb6151d5565b611de36108fa565b91829182610b74565b0390f35b610900565b34611e2357611e0d611e08366004610c16565b61529a565b611e156108fa565b80611e1f816109ec565b0390f35b610900565b34611e5857611e38366004610a24565b611e54611e436152a5565b611e4b6108fa565b91829182610b74565b0390f35b610900565b611e6a6101395f90610cbb565b90565b34611e9d57611e7d366004610a24565b611e99611e88611e5d565b611e906108fa565b91829182610b74565b0390f35b610900565b34611ed257611eb2366004610a24565b611ece611ebd6152e2565b611ec56108fa565b91829182610b74565b0390f35b610900565b34611f0757611ee7366004610a24565b611f03611ef261530d565b611efa6108fa565b91829182610b74565b0390f35b610900565b34611f3a57611f24611f1f3660046109ce565b615397565b611f2c6108fa565b80611f36816109ec565b0390f35b610900565b60018060a01b031690565b611f5a906008611f5f9302610c9c565b611f3f565b90565b90611f6d9154611f4a565b90565b611f7d6101095f90611f62565b90565b34611fb057611f90366004610a24565b611fac611f9b611f70565b611fa36108fa565b91829182610983565b0390f35b610900565b611fc261011a5f90610cbb565b90565b34611ff557611fd5366004610a24565b611ff1611fe0611fb5565b611fe86108fa565b91829182610b74565b0390f35b610900565b6120039061096a565b90565b61200f81611ffa565b0361201657565b5f80fd5b9050359061202782612006565b565b906020828203126120425761203f915f0161201a565b90565b610904565b346120755761205f61205a366004612029565b6154ad565b6120676108fa565b80612071816109ec565b0390f35b610900565b346120a85761208a366004610a24565b612092615a7b565b61209a6108fa565b806120a4816109ec565b0390f35b610900565b346120dd576120bd366004610a24565b6120d96120c8615a85565b6120d06108fa565b91829182610b74565b0390f35b610900565b6120ef61010d5f90610cbb565b90565b3461212257612102366004610a24565b61211e61210d6120e2565b6121156108fa565b91829182610b74565b0390f35b610900565b3461215557612137366004610a24565b61213f615ac7565b6121476108fa565b80612151816109ec565b0390f35b610900565b6121676101355f90610cbb565b90565b3461219a5761217a366004610a24565b61219661218561215a565b61218d6108fa565b91829182610b74565b0390f35b610900565b60018060a01b031690565b6121ba9060086121bf9302610c9c565b61219f565b90565b906121cd91546121aa565b90565b6121dd61012a5f906121c2565b90565b6121e990610f7e565b90565b6121f5906121e0565b9052565b919061220c905f602085019401906121ec565b565b3461223e5761221e366004610a24565b61223a6122296121d0565b6122316108fa565b918291826121f9565b0390f35b610900565b346122715761225b6122563660046109ce565b615b3e565b6122636108fa565b8061226d816109ec565b0390f35b610900565b61228361012f5f90611f62565b90565b346122b657612296366004610a24565b6122b26122a1612276565b6122a96108fa565b91829182610983565b0390f35b610900565b346122e9576122d36122ce3660046109ce565b615bb6565b6122db6108fa565b806122e5816109ec565b0390f35b610900565b3461231c576123066123013660046109ce565b615c2e565b61230e6108fa565b80612318816109ec565b0390f35b610900565b919060a08382031261238a57612339815f8501610b3a565b92602081013567ffffffffffffffff8111612385578261235a918301610dd6565b9261238261236b8460408501610923565b936123798160608601610e19565b93608001610923565b90565b610908565b610904565b6123a661239d366004612321565b93929092615d11565b6123ae6108fa565b806123b8816109ec565b0390f35b346123ea576123d46123cf3660046109ce565b615d8d565b6123dc6108fa565b806123e6816109ec565b0390f35b610900565b906123f990610f30565b5f5260205260405f2090565b612411906101346123ef565b9061242960016124225f8501610fc5565b9301610fc5565b90565b91602061244d92949361244660408201965f830190610b67565b0190610b67565b565b34612480576124676124623660046109ce565b612405565b9061247c6124736108fa565b9283928361242c565b0390f35b610900565b346124b35761249d6124983660046109ce565b615e05565b6124a56108fa565b806124af816109ec565b0390f35b610900565b346124e8576124c8366004610a24565b6124e46124d3615e10565b6124db6108fa565b91829182610983565b0390f35b610900565b909160608284031261253b57612505835f8401610b3a565b9260208301359067ffffffffffffffff82116125365761252a81612533938601610dd6565b93604001610923565b90565b610908565b610904565b61255461254e3660046124ed565b91615e5a565b61255c6108fa565b80612566816109ec565b0390f35b6125776101155f90610cbb565b90565b346125aa5761258a366004610a24565b6125a661259561256a565b61259d6108fa565b91829182610b74565b0390f35b610900565b6125b89061096a565b90565b6125c4816125af565b036125cb57565b5f80fd5b905035906125dc826125bb565b565b906020828203126125f7576125f4915f016125cf565b90565b610904565b3461262a5761261461260f3660046125de565b615f72565b61261c6108fa565b80612626816109ec565b0390f35b610900565b61263c6101365f90610cbb565b90565b3461266f5761264f366004610a24565b61266b61265a61262f565b6126626108fa565b91829182610b74565b0390f35b610900565b6126816101135f90610cbb565b90565b346126b457612694366004610a24565b6126b061269f612674565b6126a76108fa565b91829182610b74565b0390f35b610900565b6126c66101055f90611f62565b90565b346126f9576126d9366004610a24565b6126f56126e46126b9565b6126ec6108fa565b91829182610983565b0390f35b610900565b6127079061096a565b90565b612713816126fe565b0361271a57565b5f80fd5b9050359061272b8261270a565b565b9060208282031261274657612743915f0161271e565b90565b610904565b346127795761276361275e36600461272d565b616088565b61276b6108fa565b80612775816109ec565b0390f35b610900565b61278b6101385f90610cbb565b90565b346127be5761279e366004610a24565b6127ba6127a961277e565b6127b16108fa565b91829182610b74565b0390f35b610900565b346127f3576127d3366004610a24565b6127ef6127de616093565b6127e66108fa565b91829182610b74565b0390f35b610900565b3461282857612808366004610a24565b6128246128136160d0565b61281b6108fa565b91829182610b74565b0390f35b610900565b3461285b576128456128403660046109ce565b616115565b61284d6108fa565b80612857816109ec565b0390f35b610900565b60018060a01b031690565b61287b9060086128809302610c9c565b612860565b90565b9061288e915461286b565b90565b61289e61012e5f90612883565b90565b6128aa90610f7e565b90565b6128b6906128a1565b9052565b91906128cd905f602085019401906128ad565b565b346128ff576128df366004610a24565b6128fb6128ea612891565b6128f26108fa565b918291826128ba565b0390f35b610900565b6129116101125f90610cbb565b90565b3461294457612924366004610a24565b61294061292f612904565b6129376108fa565b91829182610b74565b0390f35b610900565b612954366004610a24565b61295c61619c565b6129646108fa565b8061296e816109ec565b0390f35b9061298461297f83610d72565b610d5d565b918252565b5f7f352e302e30000000000000000000000000000000000000000000000000000000910152565b6129ba6005612972565b906129c760208301612989565b565b6129d16129b0565b90565b6129dc6129c9565b90565b6129e76129d4565b90565b5190565b60209181520190565b90825f9392825e0152565b612a21612a2a602093612a2f93612a18816129ea565b938480936129ee565b958691016129f7565b610d16565b0190565b612a489160208201915f818403910152612a02565b90565b34612a7b57612a5b366004610a24565b612a77612a666129df565b612a6e6108fa565b91829182612a33565b0390f35b610900565b612a8d6101025f90611f62565b90565b34612ac057612aa0366004610a24565b612abc612aab612a80565b612ab36108fa565b91829182610983565b0390f35b610900565b612ad26101315f90611f62565b90565b34612b0557612ae5366004610a24565b612b01612af0612ac5565b612af86108fa565b91829182610983565b0390f35b610900565b612b1761011b5f90610cbb565b90565b34612b4a57612b2a366004610a24565b612b46612b35612b0a565b612b3d6108fa565b91829182610b74565b0390f35b610900565b612b589061096a565b90565b612b6481612b4f565b03612b6b57565b5f80fd5b90503590612b7c82612b5b565b565b90602082820312612b9757612b94915f01612b6f565b90565b610904565b34612bca57612bb4612baf366004612b7e565b6162b1565b612bbc6108fa565b80612bc6816109ec565b0390f35b610900565b634e487b7160e01b5f52603260045260245ffd5b5490565b5f5260205f2090565b612bf981612be3565b821015612c1357612c0b600491612be7565b910201905f90565b612bcf565b612c24612c2991610fac565b611f3f565b90565b612c369054612c18565b90565b634e487b7160e01b5f52602260045260245ffd5b9060016002830492168015612c6d575b6020831014612c6857565b612c39565b91607f1691612c5d565b60209181520190565b5f5260205f2090565b905f9291805490612ca3612c9c83612c4d565b8094612c77565b916001811690815f14612cfa5750600114612cbe575b505050565b612ccb9192939450612c80565b915f925b818410612ce257505001905f8080612cb9565b60018160209295939554848601520191019290612ccf565b92949550505060ff19168252151560200201905f8080612cb9565b90612d1f91612c89565b90565b90612d42612d3b92612d326108fa565b93848092612d15565b0383610d34565b565b61010090612d5182612be3565b811015612d9757612d6191612bf0565b5090612d6e5f8301610fc5565b91612d7b60018201612c2c565b91612d946003612d8d60028501610fc5565b9301612d22565b90565b5f80fd5b9092612dce90612dc4612ddb9694612dba60808601975f870190610b67565b6020850190610976565b6040830190610b67565b6060818403910152612a02565b90565b34612e1257612e0e612df9612df43660046109ce565b612d44565b90612e059492946108fa565b94859485612d9b565b0390f35b610900565b34612e4757612e43612e32612e2d366004610b49565b6162bc565b612e3a6108fa565b91829182610b74565b0390f35b610900565b916020612e6d929493612e6660408201965f830190610b67565b0190610a36565b565b34612ea057612e7f366004610a24565b612e876162d1565b90612e9c612e936108fa565b92839283612e4c565b0390f35b610900565b919060a083820312612f0e57612ebd815f8501610b3a565b92602081013567ffffffffffffffff8111612f095782612ede918301610dd6565b92612f06612eef8460408501610923565b93612efd816060860161165b565b93608001610923565b90565b610908565b610904565b612f2a612f21366004612ea5565b939290926163c9565b612f326108fa565b80612f3c816109ec565b0390f35b612f4d61010c5f90610cbb565b90565b34612f8057612f60366004610a24565b612f7c612f6b612f40565b612f736108fa565b91829182610b74565b0390f35b610900565b34612fb557612f95366004610a24565b612fb1612fa06163d8565b612fa86108fa565b91829182610b74565b0390f35b610900565b612fc76101185f90610cbb565b90565b34612ffa57612fda366004610a24565b612ff6612fe5612fba565b612fed6108fa565b91829182610b74565b0390f35b610900565b61300c6101325f90610cbb565b90565b3461303f5761301f366004610a24565b61303b61302a612fff565b6130326108fa565b91829182610b74565b0390f35b610900565b6130516101375f90610cbb565b90565b3461308457613064366004610a24565b61308061306f613044565b6130776108fa565b91829182610b74565b0390f35b610900565b6130966101175f90610cbb565b90565b346130c9576130a9366004610a24565b6130c56130b4613089565b6130bc6108fa565b91829182610b74565b0390f35b610900565b6130db61010f5f90610cbb565b90565b3461310e576130ee366004610a24565b61310a6130f96130ce565b6131016108fa565b91829182610b74565b0390f35b610900565b61314861314f9461313e606094989795613134608086019a5f870190610976565b6020850190610b67565b6040830190610976565b0190610b67565b565b3461318557613161366004610a24565b61318161316c61641a565b906131789492946108fa565b94859485613113565b0390f35b610900565b346131b8576131a261319d3660046109ce565b616696565b6131aa6108fa565b806131b4816109ec565b0390f35b610900565b6131ca61010e5f90610cbb565b90565b346131fd576131dd366004610a24565b6131f96131e86131bd565b6131f06108fa565b91829182610b74565b0390f35b610900565b90916060828403126132505761321a835f8401610923565b9260208301359067ffffffffffffffff821161324b5761323f81613248938601610dd6565b93604001610923565b90565b610908565b610904565b346132845761326e613268366004613202565b916166cd565b6132766108fa565b80613280816109ec565b0390f35b610900565b346132b7576132a161329c3660046109ce565b616747565b6132a96108fa565b806132b3816109ec565b0390f35b610900565b6132c96101335f90610cbb565b90565b346132fc576132dc366004610a24565b6132f86132e76132bc565b6132ef6108fa565b91829182610b74565b0390f35b610900565b3461333157613311366004610a24565b61332d61331c616752565b6133246108fa565b91829182610b74565b0390f35b610900565b346133645761334e6133493660046109ce565b61679f565b6133566108fa565b80613360816109ec565b0390f35b610900565b346133975761338161337c3660046109ce565b616817565b6133896108fa565b80613393816109ec565b0390f35b610900565b346133ca576133ac366004610a24565b6133b46169f5565b6133bc6108fa565b806133c6816109ec565b0390f35b610900565b6133dc6101215f90610cbb565b90565b3461340f576133ef366004610a24565b61340b6133fa6133cf565b6134026108fa565b91829182610b74565b0390f35b610900565b6134216101015f90611f62565b90565b3461345457613434366004610a24565b61345061343f613414565b6134476108fa565b91829182610983565b0390f35b610900565b6134666101085f90610cbb565b90565b3461349957613479366004610a24565b613495613484613459565b61348c6108fa565b91829182610b74565b0390f35b610900565b6134ab6101255f90610cbb565b90565b346134de576134be366004610a24565b6134da6134c961349e565b6134d16108fa565b91829182610b74565b0390f35b610900565b6134f06101105f90610cbb565b90565b3461352357613503366004610a24565b61351f61350e6134e3565b6135166108fa565b91829182610b74565b0390f35b610900565b9061353290610f30565b5f5260205260405f2090565b61355590613550610119915f92613528565b610cbb565b90565b346135885761358461357361356e3660046109ce565b61353e565b61357b6108fa565b91829182610b74565b0390f35b610900565b346135bb576135a56135a03660046109ce565b616a6c565b6135ad6108fa565b806135b7816109ec565b0390f35b610900565b346135f0576135d0366004610a24565b6135ec6135db616a85565b6135e36108fa565b91829182610a43565b0390f35b610900565b6136026101115f90610cbb565b90565b3461363557613615366004610a24565b6136316136206135f5565b6136286108fa565b91829182610b74565b0390f35b610900565b346136685761365261364d3660046109ce565b616b0e565b61365a6108fa565b80613664816109ec565b0390f35b610900565b61367a61011f5f90610cbb565b90565b346136ad5761368d366004610a24565b6136a961369861366d565b6136a06108fa565b91829182610b74565b0390f35b610900565b346136e0576136ca6136c5366004610c16565b616b7e565b6136d26108fa565b806136dc816109ec565b0390f35b610900565b34613713576136fd6136f83660046109ce565b616bba565b6137056108fa565b8061370f816109ec565b0390f35b610900565b6137256101285f90610cbb565b90565b3461375857613738366004610a24565b613754613743613718565b61374b6108fa565b91829182610b74565b0390f35b610900565b3461378b576137756137703660046109ce565b616c32565b61377d6108fa565b80613787816109ec565b0390f35b610900565b61379d6101225f90610cbb565b90565b346137d0576137b0366004610a24565b6137cc6137bb613790565b6137c36108fa565b91829182610b74565b0390f35b610900565b34613806576137e5366004610a24565b6137ed616c3d565b906138026137f96108fa565b92839283612e4c565b0390f35b610900565b60018060a01b031690565b61382690600861382b9302610c9c565b61380b565b90565b906138399154613816565b90565b6138496101295f9061382e565b90565b61385590610f7e565b90565b6138619061384c565b9052565b9190613878905f60208501940190613858565b565b346138aa5761388a366004610a24565b6138a661389561383c565b61389d6108fa565b91829182613865565b0390f35b610900565b346138df576138db6138ca6138c53660046109ce565b616c66565b6138d26108fa565b91829182610b74565b0390f35b610900565b34613916576138fd6138f7366004610f00565b90616c93565b906139126139096108fa565b9283928361242c565b0390f35b610900565b346139495761393361392e3660046109ce565b616d4a565b61393b6108fa565b80613945816109ec565b0390f35b610900565b5f80fd5b5f90565b9061396090610f30565b5f5260205260405f2090565b61399591600161398a61399093613981613952565b50610103611592565b01613956565b612c2c565b90565b6139a9906139a4616dd8565b6139ab565b565b6139bc906139b7616ec0565b6139fc565b565b5f1b90565b906139cf5f19916139be565b9181191691161790565b90565b906139f16139ec6139f892610f30565b6139d9565b82546139c3565b9055565b613a08816101386139dc565b613a3e7fc63013cf34a6f7b20983b293d1787e833f8de2db868e904525fc2910df652a9791613a356108fa565b91829182610b74565b0390a1565b613a4c90613998565b565b5f90565b613a66613a61613a6b9261090c565b610f2d565b610a33565b90565b90613a799103610a33565b90565b613a84613a4e565b50613aaa613a9142613a52565b613aa4613a9f61010d610fc5565b613a52565b90613a6e565b90565b613abe90613ab9616dd8565b613ac0565b565b613ad190613acc616ec0565b613ad3565b565b613adf816101336139dc565b613b157facbc6b6929088e4b2d043625fa7248e00fb6658a425eb9d9dc8c37b18c7f3e6f91613b0c6108fa565b91829182610b74565b0390a1565b613b2390613aad565b565b613b3690613b31616dd8565b613b38565b565b613b4990613b44616ec0565b613b4b565b565b613b578161011b6139dc565b613b8d7f40b9c59af8c486ccf8c7cc73df5a51e7cc29747ea7d39f99632ecaf9caa2ed1f91613b846108fa565b91829182610b74565b0390a1565b613b9b90613b25565b565b613bae90613ba9616dd8565b613bb0565b565b613bbc8161010c6139dc565b613bf27fb0868a729f47ae3829aaafe3ca2975d3db2148553c854112f598be6d91ef0d2891613be96108fa565b91829182610b74565b0390a1565b613c0090613b9d565b565b5f90565b613c1890613c12613c02565b50616f24565b90565b613c23613c02565b50613c2f610100612be3565b90565b613c4390613c3e616dd8565b613c45565b565b613c5690613c51616ec0565b613c58565b565b613c6a90613c658161706c565b613caa565b565b90613c7d60018060a01b03916139be565b9181191691161790565b90565b90613c9f613c9a613ca692610f8a565b613c87565b8254613c6c565b9055565b613cb681610131613c8a565b613ce07f1c7efd98583e33a9cc6adff48a97abdaaf43e5c9e918d5ec3e75e93e1dafcf6c91610f8a565b90613ce96108fa565b80613cf3816109ec565b0390a2565b613d0190613c32565b565b613d0c90610f7e565b90565b90613d1a910261090c565b90565b90565b613d34613d2f613d3992613d1d565b610f2d565b61090c565b90565b634e487b7160e01b5f52601260045260245ffd5b613d5c613d629161090c565b9161090c565b908115613d6d570490565b613d3c565b613d7a613c02565b50613dac613d9c613d8a30613d03565b31613d96610132610fc5565b90613d0f565b613da66064613d20565b90613d50565b90565b90613dc494939291613dbf6170fb565b613e6f565b613dcc617140565b565b613dda613ddf91610fac565b61137c565b90565b613dec9054613dce565b90565b5f80fd5b60e01b90565b5f910312613e0357565b610904565b613e1190610f7e565b90565b613e1d90613e08565b9052565b613e56613e5d94613e4c606094989795613e42608086019a5f870190610b67565b6020850190610976565b6040830190613e14565b0190610b67565b565b613e676108fa565b3d5f823e3d90fd5b91613e7f92949394919091617296565b613e92613e8d61012c613de2565b6113bd565b9063e2051c7e91613ea461010b610fc5565b91613ead6174bc565b9490823b15613f23575f94613ee08692613ed594613ec96108fa565b998a9889978896613df3565b865260048601613e21565b03925af18015613f1e57613ef2575b50565b613f11905f3d8111613f17575b613f098183610d34565b810190613df9565b5f613eef565b503d613eff565b613e5f565b613def565b90613f3594939291613daf565b565b613f4890613f43616dd8565b613f4a565b565b613f5b90613f56616ec0565b613f5d565b565b613f69816101216139dc565b613f9f7f3312247fcf207243294680d9103851c8709e19be3d46ee4b1bff5e12d1c5ca7b91613f966108fa565b91829182610b74565b0390a1565b613fad90613f37565b565b613fc090613fbb616dd8565b613fc2565b565b613fd390613fce616ec0565b613fd5565b565b613fe1816101376139dc565b6140177fcb78cca7628d232a9c7beef53b62f7204d9eacb44de85a8f593e6b0bb72a16219161400e6108fa565b91829182610b74565b0390a1565b61402590613faf565b565b61403890614033616dd8565b61403a565b565b61404b90614046616ec0565b61404d565b565b614059816101206139dc565b61408f7f85d8bf21006916836edd67a5afeed2e891cf876a6c9cd9babf7f42f4b007c24f916140866108fa565b91829182610b74565b0390a1565b61409d90614027565b565b6140b0906140ab616dd8565b6140b2565b565b6140c3906140be616ec0565b6140c5565b565b6140d18161011d6139dc565b6141077f5581e31f5b8d4b3c45b8ab8bf67e3602ce588b423905eb6ad34bd6bc3c848699916140fe6108fa565b91829182610b74565b0390a1565b6141159061409f565b565b90614129916141246170fb565b614392565b614131617140565b565b90565b5490565b5f5260205f2090565b61414c81614136565b8210156141665761415e60049161413a565b910201905f90565b612bcf565b61417481614136565b680100000000000000008110156141985761419491600182018155614143565b9091565b610d20565b90565b5090565b601f602091010490565b1b90565b919060086141cd9102916141c75f19846141ae565b926141ae565b9181191691161790565b91906141ed6141e86141f593610f30565b6139d9565b9083546141b2565b9055565b61420b91614205613c02565b916141d7565b565b5f5b82811061421b57505050565b8061422a5f60019385016141f9565b0161420f565b9190601f8111614240575b505050565b81811161424d575b61423b565b61426261425c61428094612c80565b916141a4565b91602061426e826141a4565b9110614288575b80910191039061420d565b5f8080614248565b505f614275565b9061429f905f1990600802610c9c565b191690565b816142ae9161428f565b906002021790565b916142c190826141a0565b9067ffffffffffffffff8211614380576142e5826142df8554612c4d565b85614230565b5f90601f831160011461431857918091614307935f9261430c575b50506142a4565b90555b565b90915001355f80614300565b601f1983169161432785612c80565b925f5b8181106143685750916002939185600196941061434e575b5050500201905561430a565b61435e910135601f84169061428f565b90555f8080614342565b9193602060018192878701358155019501920161432a565b610d20565b9061439092916142b6565b565b614400906143a1610100612be3565b9260036143c06143ba6143b5610100614133565b61416b565b5061419d565b926143d76143cf61010b610fc5565b5f86016139dc565b6143eb6143e26174bc565b60018601613c8a565b6143f834600286016139dc565b919201614385565b61440b61010b610fc5565b6144136174bc565b34929161446761445561444f6144497fa08049565b10d44a06dca9bf05685b39bc370352043c5a003e8d35d45ebdc53f94610f30565b94610f8a565b94610f30565b9461445e6108fa565b91829182610b74565b0390a4565b9061447691614117565b565b61448190610f7e565b90565b61448c613c02565b506144be6144ae61449c30614478565b316144a861011d610fc5565b90613d0f565b6144b86064613d20565b90613d50565b90565b90565b6144d86144d36144dd926144c1565b610f2d565b61090c565b90565b6144ea60026144c4565b90565b90565b6145046144ff614509926144ed565b610f2d565b61090c565b90565b90614517910361090c565b90565b90614525910161090c565b90565b61455961456791614537613c02565b506145536145436144e0565b61454d60016144f0565b9061450c565b9061451a565b6145616144e0565b90613d50565b90565b614572613c02565b50614592614581610125610fc5565b61458c610123610fc5565b90613d50565b90565b6145a6906145a1616dd8565b6145a8565b565b6145b9906145b4617548565b6145bb565b565b6145c4906175a0565b565b6145cf90614595565b565b6145e2906145dd616dd8565b6145e4565b565b6145f5906145f0616ec0565b6145f7565b565b614603816101326139dc565b6146397ffe65b6d5007c66dadebc5095104ccd672c070f396dfdcfe1ee7f54201b9efa6d916146306108fa565b91829182610b74565b0390a1565b614647906145d1565b565b61465a90614655616dd8565b61465c565b565b61466d90614668616ec0565b61466f565b565b61467b8161011f6139dc565b6146b17f4787028773c8f14bc3b4bc41f43a02329ae41105823287201a34782c530d35fd916146a86108fa565b91829182610b74565b0390a1565b6146bf90614649565b565b906146d6949392916146d16170fb565b614737565b6146de617140565b565b6146e990610f7e565b90565b6146f5906146e0565b9052565b61472e6147359461472460609498979561471a608086019a5f870190610b67565b6020850190610976565b60408301906146ec565b0190610b67565b565b9161474792949394919091617296565b61475a61475561012c613de2565b6113bd565b9063fe673fd39161476c61010b610fc5565b916147756174bc565b9490823b156147eb575f946147a8869261479d946147916108fa565b998a9889978896613df3565b8652600486016146f9565b03925af180156147e6576147ba575b50565b6147d9905f3d81116147df575b6147d18183610d34565b810190613df9565b5f6147b7565b503d6147c7565b613e5f565b613def565b906147fd949392916146c1565b565b90565b61481661481161481b926147ff565b610f2d565b610a33565b90565b61483261482d614837926147ff565b610f2d565b61090c565b90565b61484e61484961485392610a33565b610f2d565b61090c565b90565b61485e613c02565b506148676148a5565b8061488261487c6148775f614802565b610a33565b91610a33565b135f14614896576148929061483a565b5b90565b506148a05f61481e565b614893565b6148ad613a4e565b506148d36148c46148bf610124610fc5565b613a52565b6148cd42613a52565b90613a6e565b90565b6148e7906148e2616dd8565b6148e9565b565b6148fa906148f5616ec0565b6148fc565b565b614908816101136139dc565b61493e7fa787f26546d7eeea63d70fc31736f27ad28329e95982f3bc5a7e0280f497bbf5916149356108fa565b91829182610b74565b0390a1565b61494c906148d6565b565b6149566170fb565b61495e614bc6565b614966617140565b565b61497c614977614981926147ff565b610f2d565b61095f565b90565b61498d90614968565b90565b151590565b60207f757272656e742062696464696e6720726f756e64207965742e00000000000000917f54686572652068617665206265656e206e6f206269647320696e2074686520635f8201520152565b6149ef60396040926129ee565b6149f881614995565b0190565b614a119060208101905f8183039101526149e2565b90565b634e487b7160e01b5f52601160045260245ffd5b614a37614a3d91939293610a33565b92610a33565b808301925f82851215818312169285129112151617614a5857565b614a14565b60607f2e00000000000000000000000000000000000000000000000000000000000000917f4f6e6c7920746865206c61737420626964646572206973207065726d697474655f8201527f6420746f20636c61696d207468652062696464696e6720726f756e64206d616960208201527f6e207072697a65206265666f726520612074696d656f7574206578706972657360408201520152565b614b0360616080926129ee565b614b0c81614a5d565b0190565b606090614b49614b509496959396614b3f614b34608085018581035f870152614af6565b986020850190610976565b6040830190610976565b0190610b67565b565b5f7f4e6f7420656e6f7567682074696d652068617320656c61707365642e00000000910152565b614b86601c6020926129ee565b614b8f81614b52565b0190565b916040614bc4929493614bbd614bb2606083018381035f850152614b79565b966020830190610b67565b0190610b67565b565b614bce6174bc565b614bea614be4614bdf610101612c2c565b61096a565b9161096a565b145f14614c7b57614c1942614c11614c0b614c06610124610fc5565b61090c565b9161090c565b101515614990565b614c49575b614c266175e7565b614c2f42617715565b614c3761779c565b614c3f61813b565b614c476190e5565b565b614c54610124610fc5565b4290614c77614c616108fa565b928392638d31bb1560e01b845260048401614b93565b0390fd5b614cab614c89610101612c2c565b614ca3614c9d614c985f614984565b61096a565b9161096a565b141515614990565b614d4457614cd3614cba6148a5565b614ccd614cc8610127610fc5565b613a52565b90614a28565b614cf981614cf1614ceb614ce65f614802565b610a33565b91610a33565b131515614990565b614d035750614c1e565b614d0e610101612c2c565b614d40614d22614d1c6174bc565b9361483a565b614d2a6108fa565b93849363336598a360e21b855260048501614b10565b0390fd5b614d4c6108fa565b6318844a7d60e31b815280614d63600482016149fc565b0390fd5b614d6f61494e565b565b614d8290614d7d616dd8565b614d84565b565b614d9590614d90616ec0565b614da3565b565b614da090610f7e565b90565b614dbd90614db8614db382614d97565b61706c565b614e1e565b565b614dc890610f62565b90565b614dd490614dbf565b90565b614de090614dbf565b90565b90565b90614dfb614df6614e0292614dd7565b614de3565b8254613c6c565b9055565b614e0f90610f62565b90565b614e1b90614e06565b90565b614e3a614e32614e2d83614d97565b614dcb565b61012b614de6565b614e647f5bde6238168795ba4e77972a2bdaa5a465f7c9a5d05817f5e8d3fed2e5a4fa6091614e12565b90614e6d6108fa565b80614e77816109ec565b0390a2565b614e8590614d71565b565b614e9990614e93613c02565b506191de565b90565b90614eae91614ea9619208565b614eb0565b565b90614ec391614ebe816192cd565b61933d565b565b90614ecf91614e9c565b565b90565b614ee8614ee3614eed92614ed1565b610f2d565b61090c565b90565b614efb6103e8614ed4565b90565b614f17614f09614ef0565b614f11614ef0565b90613d0f565b90565b614f22614efe565b90565b614f2d613c02565b50614f4a614f3c610125610fc5565b614f44614f1a565b90613d50565b90565b614f5e90614f59616dd8565b614f60565b565b614f7190614f6c616ec0565b614f73565b565b614f7f816101226139dc565b614fb57f9e44c04f534af356419a731b967f3d56bc748b1f3fdbda7a89f4e1a321ade93491614fac6108fa565b91829182610b74565b0390a1565b614fc390614f4d565b565b5f90565b614fda90614fd561943b565b615028565b90565b90565b614ff4614fef614ff992614fdd565b6139be565b611be4565b90565b6150257f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc614fe0565b90565b50615031614ffc565b90565b61504461503f614fc5565b614fc9565b90565b61505890615053616dd8565b61505a565b565b61506b90615066616ec0565b61506d565b565b6150798161011a6139dc565b6150af7f157c413b0549fd4f45aab72b7828304fb2c45dad53de0f1128c5eabf3aaabaf8916150a66108fa565b91829182610b74565b0390a1565b6150bd90615047565b565b6150d0906150cb616dd8565b6150d2565b565b6150e3906150de616ec0565b6150f1565b565b6150ee90610f7e565b90565b61510b90615106615101826150e5565b61706c565b61516c565b565b61511690610f62565b90565b6151229061510d565b90565b61512e9061510d565b90565b90565b9061514961514461515092615125565b615131565b8254613c6c565b9055565b61515d90610f62565b90565b61516990615154565b90565b61518861518061517b836150e5565b615119565b61012e615134565b6151b27f4da1815cd654922275d14d2335fd9a0dd0aa6a0d0ff87fb4cc872ebe9704596f91615160565b906151bb6108fa565b806151c5816109ec565b0390a2565b6151d3906150bf565b565b6151dd613c02565b5061520f6151ff6151ed30613d03565b316151f9610128610fc5565b90613d0f565b6152096064613d20565b90613d50565b90565b6152239061521e616dd8565b615225565b565b61523690615231616ec0565b615238565b565b61524a906152458161706c565b61524c565b565b6152588161012f613c8a565b6152827f4d03942c29c20d1bccfe551e9d148c917c5a44fb558a4fc60270d8f76fb75f5491610f8a565b9061528b6108fa565b80615295816109ec565b0390a2565b6152a390615212565b565b6152ad613c02565b506152df6152cf6152bd30614478565b316152c9610122610fc5565b90613d0f565b6152d96064613d20565b90613d50565b90565b6152ea613c02565b5061530a6152f9610125610fc5565b615304610135610fc5565b90613d50565b90565b615315613c02565b506153276153225f614802565b614e87565b90565b61533b90615336616dd8565b61533d565b565b61534e90615349616ec0565b615350565b565b61535c816101286139dc565b6153927fb5a05ec7911dd5450a7fa4ae54d50d9d83af6e256db8fc76c82edd7b659cf8bc916153896108fa565b91829182610b74565b0390a1565b6153a09061532a565b565b6153b3906153ae616dd8565b6153b5565b565b6153c6906153c1616ec0565b6153d4565b565b6153d190610f7e565b90565b6153ee906153e96153e4826153c8565b61706c565b61544f565b565b6153f990610f62565b90565b615405906153f0565b90565b615411906153f0565b90565b90565b9061542c61542761543392615408565b615414565b8254613c6c565b9055565b61544090610f62565b90565b61544c90615437565b90565b61546b61546361545e836153c8565b6153fc565b61012c615417565b6154957fb4cecfe1346c94da27291cf5a02969d5fe0b5c36eca49b04fcd60841d28c5e1391615443565b9061549e6108fa565b806154a8816109ec565b0390a2565b6154b6906153a2565b565b6154c06154c2565b565b6154ca61564d565b565b90565b6154e36154de6154e8926154cc565b610f2d565b61090c565b90565b6154f560036154cf565b90565b67ffffffffffffffff1690565b61551961551461551e9261090c565b610f2d565b6154f8565b90565b60401c90565b60ff1690565b61553961553e91615521565b615527565b90565b61554b905461552d565b90565b67ffffffffffffffff1690565b61556761556c91610fac565b61554e565b90565b615579905461555b565b90565b9061558f67ffffffffffffffff916139be565b9181191691161790565b6155ad6155a86155b2926154f8565b610f2d565b6154f8565b90565b90565b906155cd6155c86155d492615599565b6155b5565b825461557c565b9055565b60401b90565b906155f268ff0000000000000000916155d8565b9181191691161790565b61560590614990565b90565b90565b9061562061561b615627926155fc565b615608565b82546155de565b9055565b615634906154f8565b9052565b919061564b905f6020850194019061562b565b565b61565d6156586154eb565b615505565b615665619499565b6156705f8201615541565b8015615700575b6156e4576156a99061568b835f83016155b8565b61569860015f830161560b565b6156a0615a07565b5f80910161560b565b6156df7fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2916156d66108fa565b91829182615638565b0390a1565b5f63f92ee8a960e01b8152806156fc600482016109ec565b0390fd5b5061570c5f820161556f565b61571e615718846154f8565b916154f8565b1015615677565b90565b61573c61573761574192615725565b610f2d565b61090c565b90565b61574f610e10615728565b90565b6157616157679193929361090c565b9261090c565b9161577383820261090c565b92818404149015171561578257565b614a14565b90565b61579e6157996157a392615787565b610f2d565b61090c565b90565b6157b16104b061578a565b90565b6157c06157c69161090c565b9161090c565b9081156157d1570490565b613d3c565b6157e56157eb9193929361090c565b9261090c565b82018092116157f657565b614a14565b61584961583b61581a61580c615744565b615814614f1a565b90615752565b6158356158256157a6565b61582f60026144c4565b906157b4565b906157d6565b6158436157a6565b906157b4565b90565b90565b61586361585e6158689261584c565b610f2d565b61090c565b90565b615875600d61584f565b90565b90565b61588f61588a61589492615878565b610f2d565b61090c565b90565b6158b6906158b06158aa6158bb9461090c565b9161090c565b906141ae565b61090c565b90565b6158db623671796158d66158d061586b565b9161587b565b615897565b90565b90565b6158f56158f06158fa926158de565b610f2d565b61090c565b90565b61590760086158e1565b90565b90565b61592161591c6159269261590a565b610f2d565b61090c565b90565b61593a670de0b6b3a764000061590d565b90565b90565b61595461594f6159599261593d565b610f2d565b61090c565b90565b90565b61597361596e6159789261595c565b610f2d565b61090c565b90565b6159cb6159bb6159ab61599d61598f615929565b615997615744565b90615752565b6159a5614f1a565b90615752565b6159b5601e615940565b906157d6565b6159c5603c61595f565b906157b4565b90565b90565b6159e56159e06159ea926159ce565b610f2d565b61090c565b90565b6159f7605a6159d1565b90565b615a0460036154cf565b90565b615a1a615a126157fb565b6101356139dc565b615a2d615a256158be565b6101366139dc565b615a40615a386158fd565b6101376139dc565b615a53615a4b61597b565b61011b6139dc565b615a66615a5e6159ed565b6101386139dc565b615a79615a716159fa565b6101396139dc565b565b615a836154b8565b565b615a8d613c02565b50615a9f615a9a5f614802565b6162bc565b90565b615aaa616dd8565b615ab2615ab4565b565b615ac5615ac05f614984565b6194bd565b565b615acf615aa2565b565b615ae290615add616dd8565b615ae4565b565b615af590615af0616ec0565b615af7565b565b615b03816101366139dc565b615b397f169f25ec19cc5b518dc57adf05bb2d85155b1d6c60767e9f0e3ced18eac0ab7791615b306108fa565b91829182610b74565b0390a1565b615b4790615ad1565b565b615b5a90615b55616dd8565b615b5c565b565b615b6d90615b68616ec0565b615b6f565b565b615b7b816101236139dc565b615bb17fb5edd1f338b34c8f5dd3b1c5cc12f05653c495713c282bf588d34cf14fad0f8991615ba86108fa565b91829182610b74565b0390a1565b615bbf90615b49565b565b615bd290615bcd616dd8565b615bd4565b565b615be590615be0616ec0565b615be7565b565b615bf3816101396139dc565b615c297f616bfcaa6490f55f6e57a4deedac1db04d0d6826deb84fad86cc43439bcf356491615c206108fa565b91829182610b74565b0390a1565b615c3790615bc1565b565b90615c4e94939291615c496170fb565b615c58565b615c56617140565b565b91615c68929493949190916198f1565b615c7b615c7661012c613de2565b6113bd565b9063e2051c7e91615c8d61010b610fc5565b91615c966174bc565b9490823b15615d0c575f94615cc98692615cbe94615cb26108fa565b998a9889978896613df3565b865260048601613e21565b03925af18015615d0757615cdb575b50565b615cfa905f3d8111615d00575b615cf28183610d34565b810190613df9565b5f615cd8565b503d615ce8565b613e5f565b613def565b90615d1e94939291615c39565b565b615d3190615d2c616dd8565b615d33565b565b615d4490615d3f616ec0565b615d46565b565b615d52816101186139dc565b615d887f4e8c80fe79d13b8663de9f6981925ae24c498cc07b0ebd4f4789fa78ca17caff91615d7f6108fa565b91829182610b74565b0390a1565b615d9690615d20565b565b615da990615da4616dd8565b615dab565b565b615dbc90615db7616ec0565b615dbe565b565b615dca816101356139dc565b615e007f7acba37d1b2d934e554139ff1d470d2cce50a7b6e56870577249b87564a95a6991615df76108fa565b91829182610b74565b0390a1565b615e0e90615d98565b565b615e18613952565b50615e2b5f615e25619de0565b01612c2c565b90565b90615e419291615e3c6170fb565b615e4b565b615e49617140565b565b91615e58929190916198f1565b565b90615e659291615e2e565b565b615e7890615e73616dd8565b615e7a565b565b615e8b90615e86616ec0565b615e99565b565b615e9690610f7e565b90565b615eb390615eae615ea982615e8d565b61706c565b615f14565b565b615ebe90610f62565b90565b615eca90615eb5565b90565b615ed690615eb5565b90565b90565b90615ef1615eec615ef892615ecd565b615ed9565b8254613c6c565b9055565b615f0590610f62565b90565b615f1190615efc565b90565b615f30615f28615f2383615e8d565b615ec1565b610129615edc565b615f5a7f9b3eda10f1724a2cf9f7dae4ac263c77908df4d00e92f1377b66fc8be37cd8c191615f08565b90615f636108fa565b80615f6d816109ec565b0390a2565b615f7b90615e67565b565b615f8e90615f89616dd8565b615f90565b565b615fa190615f9c616ec0565b615faf565b565b615fac90610f7e565b90565b615fc990615fc4615fbf82615fa3565b61706c565b61602a565b565b615fd490610f62565b90565b615fe090615fcb565b90565b615fec90615fcb565b90565b90565b9061600761600261600e92615fe3565b615fef565b8254613c6c565b9055565b61601b90610f62565b90565b61602790616012565b90565b61604661603e61603983615fa3565b615fd7565b61012a615ff2565b6160707fdab38e33e6e11cbb8b085bba9d7426d5e5af01bcc46d2c5957645e4d09e8c49c9161601e565b906160796108fa565b80616083816109ec565b0390a2565b61609190615f7d565b565b61609b613c02565b506160cd6160bd6160ab30614478565b316160b761011e610fc5565b90613d0f565b6160c76064613d20565b90613d50565b90565b6160d8613c02565b506160e1619e04565b90565b6160f5906160f0616dd8565b6160f7565b565b61610890616103616ec0565b61610a565b565b61611390619e18565b565b61611e906160e4565b565b6161286170fb565b61613061613a565b616138617140565b565b61614561010b610fc5565b61614d6174bc565b349161619761618561617f7fe32cacf203d00685e2b4d8b0a90e7cd8f3f8a208fdf116f4bb36abe08b7d548e93610f30565b93610f8a565b9361618e6108fa565b91829182610b74565b0390a3565b6161a4616120565b565b6161b7906161b2616dd8565b6161b9565b565b6161ca906161c5616ec0565b6161d8565b565b6161d590610f7e565b90565b6161f2906161ed6161e8826161cc565b61706c565b616253565b565b6161fd90610f62565b90565b616209906161f4565b90565b616215906161f4565b90565b90565b9061623061622b6162379261620c565b616218565b8254613c6c565b9055565b61624490610f62565b90565b6162509061623b565b90565b61626f616267616262836161cc565b616200565b61012d61621b565b6162997fbf6e296f85d08cc1ab124aed644bf4b19e4a726a7aea53e3784ab1341738a04091616247565b906162a26108fa565b806162ac816109ec565b0390a2565b6162ba906161a6565b565b6162ce906162c8613c02565b50619e5f565b90565b6162d9613c02565b506162e2613a4e565b506162eb619fff565b91909190565b90616306949392916163016170fb565b616310565b61630e617140565b565b91616320929493949190916198f1565b61633361632e61012c613de2565b6113bd565b9063fe673fd39161634561010b610fc5565b9161634e6174bc565b9490823b156163c4575f9461638186926163769461636a6108fa565b998a9889978896613df3565b8652600486016146f9565b03925af180156163bf57616393575b50565b6163b2905f3d81116163b8575b6163aa8183610d34565b810190613df9565b5f616390565b503d6163a0565b613e5f565b613def565b906163d6949392916162f1565b565b6163e0613c02565b506163f26163ed5f614802565b613c06565b90565b61640461640a9193929361090c565b9261090c565b820391821161641557565b614a14565b616422613952565b9061642b613c02565b90616434613952565b9061643d613c02565b90616449610101612c2c565b61646361645d6164585f614984565b61096a565b9161096a565b0361646b575b565b935050505061647b610105612c2c565b90616487610106610fc5565b616492610107610fc5565b9261649e610108610fc5565b926164aa610109612c2c565b926164b661010a610fc5565b926164ed60026164e76164d66101046164d061010b610fc5565b90610f4c565b6164e1610101612c2c565b90610f96565b01610fc5565b966164f94289906163f5565b928261651561650f61650a5f614984565b61096a565b9161096a565b145f1461657e575050506165429061653c616531610101612c2c565b9791965b42926157d6565b906163f5565b61654b81613a52565b61656561655f61655a86613a52565b610a33565b91610a33565b13616571575b50616469565b925090508391905f61656b565b928098919792986165976165918a61090c565b9161090c565b116165ac575b50509061653c61654292616535565b96926165c88299936165c26165ce9487906157d6565b926157d6565b906163f5565b906165d882613a52565b6165f26165ec6165e788613a52565b610a33565b91610a33565b13616617575b5050616542909461653c61660d610101612c2c565b979196919261659d565b9094506165429193509392905f6165f8565b61663a90616635616dd8565b61663c565b565b61664d90616648616ec0565b61664f565b565b61665b816101266139dc565b6166917f4636d3e567b27988879babd22e50f49104ca65647933fc1623ff3d4d807438d2916166886108fa565b91829182610b74565b0390a1565b61669f90616629565b565b906166b492916166af6170fb565b6166be565b6166bc617140565b565b916166cb92919091617296565b565b906166d892916166a1565b565b6166eb906166e6616dd8565b6166ed565b565b6166fe906166f9616ec0565b616700565b565b61670c816101306139dc565b6167427f2652e6657dd1ed89d7bdcb70f8827cc8449ac4536ebf604dbb2465cdad264020916167396108fa565b91829182610b74565b0390a1565b616750906166da565b565b61675a613c02565b5061676b616766614f25565b61a026565b90565b61677f9061677a616dd8565b616781565b565b6167929061678d616ec0565b616794565b565b61679d9061a05b565b565b6167a89061676e565b565b6167bb906167b6616dd8565b6167bd565b565b6167ce906167c9616ec0565b6167d0565b565b6167dc816101156139dc565b6168127f4abea08c196329c357e3175d011af39a8625be99ef0ba5a0f3547a95534fedb7916168096108fa565b91829182610b74565b0390a1565b616820906167aa565b565b61682a616dd8565b616832616834565b565b61683c61683e565b565b616846617548565b61684e6168a9565b565b5f7f546f6f206561726c792e00000000000000000000000000000000000000000000910152565b616884600a6020926129ee565b61688d81616850565b0190565b6168a69060208101905f818303910152616877565b90565b6168d86168b4616c3d565b91906168d16168cb6168c68593613a52565b610a33565b91610a33565b1315614990565b6169d2576169d0906169cb6169c66169b66168f4610110610fc5565b6169b061693361692261691261690b61010f610fc5565b85906157b4565b61691c60016144f0565b906157d6565b9261692d60026144c4565b90615752565b956169aa6169a461699361698261696861695861695161010f610fc5565b8d906157b4565b61696260016144f0565b906157d6565b96616971613c02565b5061697d61010f610fc5565b6163f5565b61698d610125610fc5565b90615752565b9461699f61010f610fc5565b6163f5565b9161483a565b90615752565b906157b4565b6169c060016144f0565b906157d6565b61a0a2565b61a05b565b565b6169da6108fa565b63a29f5c4d60e01b8152806169f160048201616891565b0390fd5b6169fd616822565b565b616a1090616a0b616dd8565b616a12565b565b616a2390616a1e616ec0565b616a25565b565b616a318161011c6139dc565b616a677fd95e7f967f9370c11deb15ffbb191b9f2e9795ab0738db5bc72bd2794978f32d91616a5e6108fa565b91829182610b74565b0390a1565b616a75906169ff565b565b616a82905f03610a33565b90565b616a8d613a4e565b50616a9e616a99613a7c565b616a77565b90565b616ab290616aad616dd8565b616ab4565b565b616ac590616ac0616ec0565b616ac7565b565b616ad3816101276139dc565b616b097f37a332914fac995349420c0419b4423a19dcb762017f691442a0782ce4bf417a91616b006108fa565b91829182610b74565b0390a1565b616b1790616aa1565b565b616b2a90616b25616dd8565b616b2c565b565b80616b47616b41616b3c5f614984565b61096a565b9161096a565b14616b5757616b55906194bd565b565b616b7a616b635f614984565b5f918291631e4fbdf760e01b835260048301610983565b0390fd5b616b8790616b19565b565b616b9a90616b95616dd8565b616b9c565b565b616bad90616ba8616ec0565b616baf565b565b616bb89061a0a2565b565b616bc390616b89565b565b616bd690616bd1616dd8565b616bd8565b565b616be990616be4616ec0565b616beb565b565b616bf7816101126139dc565b616c2d7fdeb71e1d117914859ddde484a9810084d9ee399077d5cd8dcbdf8770d33d8ae491616c246108fa565b91829182610b74565b0390a1565b616c3b90616bc5565b565b616c45613c02565b50616c4e613a4e565b50616c5761a0e9565b90616c60613a7c565b90565b90565b5f616c87616c82616c8d93616c79613c02565b50610103611592565b616c63565b01610fc5565b90565b90565b616cc291616cb8616cbd92616ca6613c02565b50616caf613c02565b50610104610f4c565b610f96565b616c90565b90616cda6001616cd35f8501610fc5565b9301610fc5565b90565b616cee90616ce9616dd8565b616cf0565b565b616d0190616cfc616ec0565b616d03565b565b616d0f8161011e6139dc565b616d457fbfcd8fb930a57c1598c9760db19c84ec766546c3c9a8565611df8302482bfb1791616d3c6108fa565b91829182610b74565b0390a1565b616d5390616cdd565b565b616d5d6170fb565b616d65616da5565b616d6d617140565b565b90565b616d86616d81616d8b92616d6f565b610f2d565b610a33565b90565b616d975f612972565b90565b616da2616d8e565b90565b616dcc5f19616db45f91616d72565b90616dc6616dc0616d9a565b9161481e565b916198f1565b565b616dd6616d55565b565b616de0615e10565b616df9616df3616dee6174bc565b61096a565b9161096a565b03616e0057565b616e22616e0b6174bc565b5f91829163118cdaa760e01b835260048301610983565b0390fd5b60207f65616479206163746976652e0000000000000000000000000000000000000000917f5468652063757272656e742062696464696e6720726f756e6420697320616c725f8201520152565b616e80602c6040926129ee565b616e8981616e26565b0190565b916040616ebe929493616eb7616eac606083018381035f850152616e73565b966020830190610b67565b0190610b67565b565b616ecb61010d610fc5565b616ee842616ee1616edb8461090c565b9161090c565b1015614990565b616eef5750565b4290616f12616efc6108fa565b92839263d0fd11df60e01b845260048401616e8d565b0390fd5b90616f219101610a33565b90565b616f8a90616f30613c02565b50616f3c610101612c2c565b616f56616f50616f4b5f614984565b61096a565b9161096a565b145f14616fd257616f84616f7e616f6e61010d610fc5565b5b92616f7942613a52565b616f16565b91613a52565b90613a6e565b616f935f61481e565b9080616faf616fa9616fa45f614802565b610a33565b91610a33565b13616fb9575b5090565b616fcc9150616fc79061483a565b61a026565b5f616fb5565b616f84616f7e61700e6002617008616ff7610104616ff161010b610fc5565b90610f4c565b617002610101612c2c565b90610f96565b01610fc5565b616f6f565b5f7f5468652070726f76696465642061646472657373206973207a65726f2e000000910152565b617047601d6020926129ee565b61705081617013565b0190565b6170699060208101905f81830391015261703a565b90565b61708661708061707b5f614984565b61096a565b9161096a565b1461708d57565b6170956108fa565b63eac0d38960e01b8152806170ac60048201617054565b0390fd5b90565b6170c76170c26170cc926170b0565b6139be565b611be4565b90565b6170f87f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f006170b3565b90565b61710361a118565b6171245761712261711a6171156170cf565b61a14c565b60019061a161565b565b5f633ee5aeb560e01b81528061713c600482016109ec565b0390fd5b61715a61715361714e6170cf565b61a14c565b5f9061a161565b565b60407f642e000000000000000000000000000000000000000000000000000000000000917f5468652063757272656e742043535420626964207072696365206973206772655f8201527f61746572207468616e20746865206d6178696d756d20796f7520616c6c6f776560208201520152565b6171dc60426060926129ee565b6171e58161715c565b0190565b91604061721a929493617213617208606083018381035f8501526171cf565b966020830190610b67565b0190610b67565b565b61722660026144c4565b90565b61723290616d72565b9052565b919461728361727861728d9360a09661726b6172949a9c9b999c61726160c08a01945f8b0190617229565b6020890190610a36565b8682036040880152612a02565b986060850190610b67565b6080830190610b67565b0190610b67565b565b6172a76172a25f614802565b613c06565b926172c6846172be6172b88461090c565b9161090c565b101515614990565b61749d57506172dc6172d75f614802565b6162bc565b906172fb826172f36172ed8461090c565b9161090c565b111515614990565b617477575061730b81849061a164565b6173528161734c600161733c61732e61010461732861010b610fc5565b90610f4c565b6173366174bc565b90610f96565b019161734783610fc5565b6157d6565b906139dc565b61735e426101146139dc565b6173816173738261736d61721c565b90615752565b61737b619e04565b9061a170565b61738d816101166139dc565b617398610102612c2c565b6173b26173ac6173a75f614984565b61096a565b9161096a565b14617465575b506173cc6173c46174bc565b610102613c8a565b6173d461a19c565b6173dd8361a310565b6173e861010b610fc5565b906173f16174bc565b926174606174005f1992613a52565b925f19969790617411610124610fc5565b9161744e6174486174427f1d1f406c89e99504a7222aaba50bf96b9f91f522ebc6ea825255145919e102ec99610f30565b99610f8a565b99616d72565b996174576108fa565b96879687617236565b0390a4565b617471906101176139dc565b5f6173b8565b906174996174836108fa565b92839263814ac7ff60e01b8452600484016171e9565b0390fd5b836174b85f9283926335465b3160e01b84526004840161242c565b0390fd5b6174c4613952565b503390565b60207f207468652063757272656e742062696464696e6720726f756e642e0000000000917f41206269642068617320616c7265616479206265656e20706c6163656420696e5f8201520152565b617523603b6040926129ee565b61752c816174c9565b0190565b6175459060208101905f818303910152617516565b90565b617577617556610101612c2c565b61757061756a6175655f614984565b61096a565b9161096a565b1415614990565b61757d57565b6175856108fa565b634283f4b960e01b81528061759c60048201617530565b0390fd5b6175ac8161010d6139dc565b6175e27f9a2159c1f277ddd727551baedc6a6c4cba77cc5219c8563ee3b15fb67548d89b916175d96108fa565b91829182610b74565b0390a1565b61761d600261761761760661010461760061010b610fc5565b90610f4c565b617611610101612c2c565b90610f96565b01610fc5565b6176284282906163f5565b617633610105612c2c565b61764d6176476176425f614984565b61096a565b9161096a565b145f146176845761767961768192617671617669610101612c2c565b610105613c8a565b6101066139dc565b6101076139dc565b5b565b806176a161769b617696610107610fc5565b61090c565b9161090c565b116176ae575b5050617682565b617706826176d26176cd61770e956176c7610107610fc5565b906157d6565b617715565b6176e86176e0610107610fc5565b6101086139dc565b6176fe6176f6610101612c2c565b610105613c8a565b6101066139dc565b6101076139dc565b5f806176a7565b61773e90617738617727610106610fc5565b617732610108610fc5565b906157d6565b906163f5565b61774781613a52565b61776b61776561776061775b61010a610fc5565b613a52565b610a33565b91610a33565b13617774575b50565b6177969061778e617786610105612c2c565b610109613c8a565b61010a6139dc565b5f617771565b6177a461a4dc565b565b6177b06020610d5d565b90565b5f90565b6177bf6177a6565b906020826177cb6177b3565b81525050565b6177d96177b7565b90565b906177e69061090c565b9052565b67ffffffffffffffff81116178025760208091020190565b610d20565b90617819617814836177ea565b610d5d565b918252565b6178286040610d5d565b90565b5f90565b61783761781e565b906020808361784461782b565b81520161784f6177b3565b81525050565b61785d61782f565b90565b5f5b82811061786e57505050565b602090617879617855565b8184015201617862565b906178a861789083617807565b9260208061789e86936177ea565b9201910390617860565b565b5190565b906178b8826178aa565b8110156178c9576020809102010190565b612bcf565b906178d89061096a565b9052565b6178e59061090c565b5f81146178f3576001900390565b614a14565b61790461790a9161090c565b9161090c565b908115617915570690565b613d3c565b905051906179278261090f565b565b906020828203126179425761793f915f0161791a565b90565b610904565b60209181520190565b60200190565b61795f9061096a565b9052565b61796c9061090c565b9052565b90602080617992936179885f8201515f860190617956565b0151910190617963565b565b906179a181604093617970565b0190565b60200190565b906179c86179c26179bb846178aa565b8093617947565b92617950565b905f5b8181106179d85750505090565b9091926179f16179eb6001928651617994565b946179a5565b91019190916179cb565b617a1f617a2c949293617a1560608401955f850190610b67565b6020830190610976565b60408184039101526179ab565b90565b617a3b617a4091610fac565b612860565b90565b617a4d9054617a2f565b90565b5f9060033d11617a5d575b565b905060045f803e617a6e5f516108f4565b90617a5b565b5f5f9160233d11617a82575b565b915050602060045f3e6001905f5190617a80565b90565b617aad617aa8617ab292617a96565b610f2d565b61090c565b90565b617abf6012617a99565b90565b905090565b617ad25f8092617ac2565b0190565b617adf90617ac7565b90565b90617af4617aef83611a91565b610d5d565b918252565b606090565b3d5f14617b1957617b0e3d617ae2565b903d5f602084013e5b565b617b21617af9565b90617b17565b5f7f455448207472616e7366657220746f2063686172697479206661696c65642e00910152565b617b5b601f6020926129ee565b617b6481617b27565b0190565b9190617b8b906020617b83604086018681035f880152617b4e565b940190610b67565b565b60207f696e207072697a652062656e6566696369617279206661696c65642e00000000917f455448207472616e7366657220746f2062696464696e6720726f756e64206d615f8201520152565b617be7603c6040926129ee565b617bf081617b8d565b0190565b916040617c25929493617c1e617c13606083018381035f850152617bda565b966020830190610976565b0190610b67565b565b60ff1690565b617c41617c3c617c46926154cc565b610f2d565b617c27565b90565b90565b617c60617c5b617c6592617c49565b610f2d565b617c27565b90565b617c7c617c77617c8192617c27565b610f2d565b61090c565b90565b617c90617c9591610fac565b611488565b90565b617ca29054617c84565b90565b617caf905161090c565b90565b90565b617cc9617cc4617cce92617cb2565b610f2d565b61090c565b90565b67ffffffffffffffff8111617ce95760208091020190565b610d20565b90505190617cfb82610bf3565b565b90929192617d12617d0d82617cd1565b610d5d565b9381855260208086019202830192818411617d4f57915b838310617d365750505050565b60208091617d448486617cee565b815201920191617d29565b6111e5565b9080601f83011215617d7257816020617d6f93519101617cfd565b90565b610d0e565b90602082820312617da7575f82015167ffffffffffffffff8111617da257617d9f9201617d54565b90565b610908565b610904565b5190565b90617dc2617dbd83617cd1565b610d5d565b918252565b369037565b90617df1617dd983617db0565b92602080617de78693617cd1565b9201910390617dc7565b565b67ffffffffffffffff8111617e0b5760208091020190565b610d20565b90617e22617e1d83617df3565b610d5d565b918252565b617e316040610d5d565b90565b617e3c617e27565b9060208083617e4961782b565b815201617e546177b3565b81525050565b617e62617e34565b90565b5f5b828110617e7357505050565b602090617e7e617e5a565b8184015201617e67565b90617ead617e9583617e10565b92602080617ea38693617df3565b9201910390617e65565b565b5190565b90617ebd82617eaf565b811015617ece576020809102010190565b612bcf565b90617edd82617dac565b811015617eee576020809102010190565b612bcf565b617efd905161096a565b90565b617f0c617f1191610fac565b61380b565b90565b617f1e9054617f00565b90565b60209181520190565b60200190565b90602080617f5293617f485f8201515f860190617956565b0151910190617963565b565b90617f6181604093617f30565b0190565b60200190565b90617f88617f82617f7b84617eaf565b8093617f21565b92617f2a565b905f5b818110617f985750505090565b909192617fb1617fab6001928651617f54565b94617f65565b9101919091617f8b565b617fd09160208201915f818403910152617f6b565b90565b617fdf617fe491610fac565b611973565b90565b617ff19054617fd3565b90565b90565b61800b61800661801092617ff4565b610f2d565b61090c565b90565b60209181520190565b60200190565b9061802f81602093617956565b0190565b60200190565b9061805661805061804984617dac565b8093618013565b9261801c565b905f5b8181106180665750505090565b90919261807f6180796001928651618022565b94618033565b9101919091618059565b9392906180b46040916180bc946180a760608901925f8a0190610b67565b8782036020890152618039565b940190610b67565b565b6180f36180fa946180e96060949897956180df608086019a5f870190610b67565b6020850190610b67565b6040830190610b67565b0190610b67565b565b61810590614990565b9052565b604090618132618139949695939661812860608401985f8501906180fc565b6020830190610b67565b0190610b67565b565b6181436177d1565b9061815761814f61a5c1565b5f84016177dc565b61817661817161010361816b61010b610fc5565b90611592565b616c63565b9061817f613c02565b506181886151d5565b91618191614484565b618199613d72565b926181a26152a5565b906181ab616093565b956181b761011f610fc5565b936181d46181cf866181c960016144f0565b906157d6565b617883565b9661823761822561821d6181e75f61481e565b6182166181f58d8c906178ae565b5161820c618204610109612c2c565b5f83016178ce565b60208891016177dc565b85906157d6565b9a88906157b4565b996182318b8990615752565b906157d6565b60015b1561830f575b5f9661824b906178dc565b968988618257916178ae565b518c6182629061a6b5565b8a600101908b5f0161827390610fc5565b61827c916178f8565b61828591613956565b61828e90612c2c565b9081815f019061829d916178ce565b8c90602001906182ac916177dc565b6182b761010b610fc5565b8991908d927f9c62e2cb8cbd10bf3b8a5760977d719fc3475fa67cb7ef9d2e1701f275e885c4916182e790610f30565b926182f190610f8a565b936182fa6108fa565b918291618307918361242c565b0390a361823a565b8661832261831c5f61481e565b9161090c565b1161824057909398506020919499979296955061834861834361012c613de2565b6113bd565b906183816387565d149192919261836061010b610fc5565b9361838c61836c6174bc565b976183756108fa565b98899788968795613df3565b8552600485016179fb565b03925af19081156190d1575f916190a3575b50966183b36183ae61012e617a43565b6128a1565b9063b6b55f2590919091906183c961010b610fc5565b90803b1561909e576183ee5f936183f9956183e26108fa565b96879586948593613df3565b835260048301610b74565b03925af19081619072575b50155f1461906d576001618416617a50565b634e487b7114619030575b61902b575b5f80618433610131612c2c565b8361843c6108fa565b908161844781617ad6565b03925af1618453617afe565b505f14618fd957618465610131612c2c565b6184a46184927f1222634ba80f397fa40371bab63974e6f97da8200777ec79e731c69bb6a2351d92610f8a565b9261849b6108fa565b91829182610b74565b0390a25b6184dc5f806184b56174bc565b866184be6108fa565b90816184c981617ad6565b03925af16184d5617afe565b5015614990565b618fab576184eb610102612c2c565b6185056184ff6184fa5f614984565b61096a565b9161096a565b14155f14618f965761853461852361851d6004617c4c565b5b617c68565b61852e610120610fc5565b906157d6565b9261854861854361012d617c98565b6114c9565b5f63e36aee789161855a610121610fc5565b906185ab618569848b01617ca5565b6185927f7c6eeb003d4a6dc5ebf549935c6ffb814ba1f060f1af8a0b11c2aa94a8e716e4617cb5565b18946185b661859f6108fa565b96879586948594613df3565b84526004840161242c565b03915afa8015618f91576185dc915f91618f6f575b50946185d686617dac565b906157d6565b9561860d618608886186026185f2610139610fc5565b6185fc60016144f0565b906163f5565b906157d6565b617dcc565b9761866c61862d6186288a61862260016144f0565b906157d6565b617e88565b9861866761863c8b8390617eb3565b5161865361864b61012f612c2c565b5f83016178ce565b6020618660610130610fc5565b91016177dc565b6178dc565b966186a361867b8a8a90617eb3565b5161868f6186876174bc565b5f83016178ce565b602061869c61011c610fc5565b91016177dc565b6186b56186af8b617dac565b5b6178dc565b6186d26186c06174bc565b6186cd8d91849092617ed3565b6178ce565b806186e56186df8b61090c565b9161090c565b11156186f4576186b5906186b0565b506186fe87617dac565b5b8061871261870c5f61481e565b9161090c565b111561878957618721906178dc565b9561878361874161873b6187368b8b90617ed3565b617ef3565b9a6178dc565b996187716187508d8d90617eb3565b5161875d835f83016178ce565b602061876a61011c610fc5565b91016177dc565b61877e8d918c9092617ed3565b6178ce565b956186ff565b509297969093989591946188286187a1610120610fc5565b5b6188236187e16187db6187d66187b78a61a6b5565b6187d060018901916187ca5f8b01610fc5565b906178f8565b90613956565b612c2c565b956178dc565b946188116187f08d8890617eb3565b516187fd835f83016178ce565b602061880a61011c610fc5565b91016177dc565b61881e8b91879092617ed3565b6178ce565b6178dc565b908161883c6188365f61481e565b9161090c565b111561885d57906188236187e16187db6187d66188289493505050506187a2565b505061886b6188c9916178dc565b6188a4618879898390617eb3565b51618890618888610109612c2c565b5f83016178ce565b602061889d61011c610fc5565b91016177dc565b6188c46188b2610109612c2c565b6188bf8991849092617ed3565b6178ce565b6178dc565b6189026188d7888390617eb3565b516188ee6188e6610105612c2c565b5f83016178ce565b60206188fb61011c610fc5565b91016177dc565b618922618910610105612c2c565b61891d8891849092617ed3565b6178ce565b61893461892e5f61481e565b9161090c565b115f14618f6a5761897c6189518761894b5f61481e565b90617eb3565b51618968618960610102612c2c565b5f83016178ce565b602061897561011c610fc5565b91016177dc565b6189a261898a610102612c2c565b61899d876189975f61481e565b90617ed3565b6178ce565b5b6189b66189b1610129617f14565b61384c565b63b33266da87823b15618f65576189ec926189e15f80946189d56108fa565b96879586948593613df3565b835260048301617fbb565b03925af18015618f6057618f34575b506020618a11618a0c61012b617fe7565b6119b4565b636578f11390618a745f618a2661010b610fc5565b93618a7f618a36838d9a01617ca5565b618a5f7f2a8612ecb5cb17da87f8befda0480288e2d053de55d9d7d4dc4899077cf5aeda617ff7565b18618a686108fa565b998a9788968795613df3565b855260048501618089565b03925af18015618f2f57618b65925f91618f01575b509794969792618abc618ab7618aa989617eaf565b618ab1613c02565b506178dc565b6178dc565b97618ac8888a90617eb3565b519989618ad68782906157d6565b809c618b5d618b0e618ae961010b610fc5565b94618b09618b026020618afa6174bc565b999601617ca5565b9598617dac565b6163f5565b96618b4b618b45618b3f7f9314e78538382a9fc6cd54cee107c8a7721d172c43b2e48c2e7a51f200fc378897610f30565b97610f8a565b97610f30565b97618b546108fa565b948594856180be565b0390a4617dac565b945b85618b7a618b745f61481e565b9161090c565b1115618c3257618b89906178dc565b92618b95858590617eb3565b5195618bb4618bae618ba85f8a01617ef3565b926178dc565b986178dc565b96618bc061010b610fc5565b90600191618bd260208c959301617ca5565b938a93618c26618c14618c0e618c087f27c21fe4cea1a3367aa491829dd4dd824296c00910626150464cba8ea5ebb3f496610f30565b96610f8a565b96610f30565b96618c1d6108fa565b93849384618109565b0390a492959495618b67565b9195909450618c4c618c45610120610fc5565b925b6178dc565b618c57858290617eb3565b5192618c76618c70618c6a5f8701617ef3565b926178dc565b936178dc565b93618c8261010b610fc5565b905f91618c93602087959301617ca5565b938793618ce7618cd5618ccf618cc97f27c21fe4cea1a3367aa491829dd4dd824296c00910626150464cba8ea5ebb3f496610f30565b96610f8a565b96610f30565b96618cde6108fa565b93849384618109565b0390a481618cfd618cf75f61481e565b9161090c565b1115618d0f57618c4c90929192618c47565b618dc1929150618d1e906178dc565b90618d34618d2d868490617eb3565b51916178dc565b93618d4061010b610fc5565b90618d4c61011f610fc5565b91618d656020618d5d5f8701617ef3565b939501617ca5565b938793618db9618da7618da1618d9b7faa858ae20a26d00a9ea528972d537e68a51a0744226d9ea1fc9b22492dc282a596610f30565b96610f8a565b96610f30565b96618db06108fa565b9384938461100e565b0390a46178dc565b90618dd7618dd0848490617eb3565b51916178dc565b90618de361010b610fc5565b90618dfb6020618df45f8401617ef3565b9201617ca5565b9291618e4e618e3c618e36618e307f838ec9dd2530548892bff113f5ffb0138d2efc63c7f59bb6571e8c923b63126094610f30565b94610f8a565b94610f30565b94618e456108fa565b91829182610b74565b0390a4618e63618e5d5f61481e565b9161090c565b115f14618efa57618e7d90618e775f61481e565b90617eb3565b51618e8961010b610fc5565b90618ea16020618e9a5f8401617ef3565b9201617ca5565b9291618ef4618ee2618edc618ed67f3901b6430c99dc290ee88ff84c4de6091ad7eac335b58e92ef5cbb0793abf4f694610f30565b94610f8a565b94610f30565b94618eeb6108fa565b91829182610b74565b0390a45b565b5050618ef8565b618f22915060203d8111618f28575b618f1a8183610d34565b810190617929565b5f618a94565b503d618f10565b613e5f565b618f53905f3d8111618f59575b618f4b8183610d34565b810190613df9565b5f6189fb565b503d618f41565b613e5f565b613def565b6189a3565b618f8b91503d805f833e618f838183610d34565b810190617d77565b5f6185cb565b613e5f565b618534618523618fa66003617c2d565b61851e565b82618fb46174bc565b618fd5618fbf6108fa565b928392630aa7db6360e11b845260048401617bf4565b0390fd5b618fe4610131612c2c565b6190236190117f154fb6c686c977314af35b730a16b83facda5265b2abec7237c730f63e42215a92610f8a565b9261901a6108fa565b91829182617b68565b0390a26184a8565b613e5f565b619038617a74565b90619044575b50618421565b90505f908061906261905c619057617ab5565b61090c565b9161090c565b031561903e5761a6f1565b618426565b619091905f3d8111619097575b6190898183610d34565b810190613df9565b5f618404565b503d61907f565b613def565b6190c4915060203d81116190ca575b6190bc8183610d34565b810190617929565b5f61839e565b503d6190b2565b613e5f565b60016190e2910161090c565b90565b6190f96190f15f614984565b610101613c8a565b61910d6191055f614984565b610102613c8a565b6191216191195f614984565b610105613c8a565b61913561912d5f61481e565b6101086139dc565b6191496191415f614984565b610109613c8a565b61916661915e6191595f19616d72565b61483a565b61010a6139dc565b61918461917c61917761010b610fc5565b6190d6565b61010b6139dc565b6191bf6191ba619195610125610fc5565b6191b46191a3610125610fc5565b6191ae610126610fc5565b90613d50565b9061451a565b619e18565b6191dc6191d7426191d161010c610fc5565b9061451a565b6175a0565b565b6191f9906191ea613c02565b506191f48161a701565b61a86e565b90565b61920590610f7e565b90565b619211306191fc565b61924361923d7f000000000000000000000000000000000000000000000000000000000000000061096a565b9161096a565b14801561926d575b61925157565b5f63703e46dd60e11b815280619269600482016109ec565b0390fd5b5061927661a9bb565b6192a86192a27f000000000000000000000000000000000000000000000000000000000000000061096a565b9161096a565b141561924b565b6192c0906192bb616dd8565b6192c2565b565b506192cb616ec0565b565b6192d6906192af565b565b6192e190610f62565b90565b6192ed906192d8565b90565b6192f990610f7e565b90565b61930581611be4565b0361930c57565b5f80fd5b9050519061931d826192fc565b565b9060208282031261933857619335915f01619310565b90565b610904565b919061936b6020619355619350866192e4565b6192f0565b6352d1902d906193636108fa565b938492613df3565b8252818061937b600482016109ec565b03915afa80915f9261940b575b50155f146193bc57505090600161939d57505b565b6193b8905f918291634c9c8ce360e01b835260048301610983565b0390fd5b92836193d76193d16193cc614ffc565b611be4565b91611be4565b036193ec576193e792935061a9e1565b61939b565b619407845f918291632a87526960e21b835260048301611bf4565b0390fd5b61942d91925060203d8111619434575b6194258183610d34565b81019061931f565b905f619388565b503d61941b565b619444306191fc565b6194766194707f000000000000000000000000000000000000000000000000000000000000000061096a565b9161096a565b0361947d57565b5f63703e46dd60e11b815280619495600482016109ec565b0390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0090565b6194c5619de0565b6194dd6194d35f8301612c2c565b915f849101613c8a565b9061951161950b7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e093610f8a565b91610f8a565b9161951a6108fa565b80619524816109ec565b0390a3565b61953861953e91939293610a33565b92610a33565b91828103925f82851281831216928513911215161761955957565b614a14565b60407f727265642e000000000000000000000000000000000000000000000000000000917f5468652063757272656e742045544820626964207072696365206973206772655f8201527f61746572207468616e2074686520616d6f756e7420796f75207472616e73666560208201520152565b6195de60456060926129ee565b6195e78161955e565b0190565b91604061961c92949361961561960a606083018381035f8501526195d1565b966020830190610b67565b0190610b67565b565b60207f206265656e207573656420666f722062696464696e672e000000000000000000917f546869732052616e646f6d2057616c6b204e46542068617320616c72656164795f8201520152565b61967860376040926129ee565b6196818161961e565b0190565b91906196a89060206196a0604086018681035f88015261966b565b940190610b67565b565b156196b25750565b6196d4906196be6108fa565b91829163c35947c560e01b835260048301619685565b0390fd5b6196e46196e991610fac565b61219f565b90565b6196f690546196d8565b90565b906020828203126197125761970f915f01617cee565b90565b610904565b60207f6e646f6d2057616c6b204e46542e000000000000000000000000000000000000917f596f7520617265206e6f7420746865206f776e6572206f6620746869732052615f8201520152565b619771602e6040926129ee565b61977a81619717565b0190565b6060906197b76197be94969593966197ad6197a2608085018581035f870152619764565b9860208501906121ec565b6040830190610b67565b0190610976565b565b92909192156197ce57505050565b6197f0906197da6108fa565b938493630b81342760e31b85526004850161977e565b0390fd5b6197fe60026144c4565b90565b61981561981061981a92610a33565b610f2d565b610a33565b90565b919461986a61985f6198749360a09661985261987b9a9c9b999c61984860c08a01945f8b0190610a36565b6020890190617229565b8682036040880152612a02565b986060850190610b67565b6080830190610b67565b0190610b67565b565b5f7f45544820726566756e64207472616e73666572206661696c65642e0000000000910152565b6198b1601b6020926129ee565b6198ba8161987d565b0190565b9160406198ef9294936198e86198dd606083018381035f8501526198a4565b966020830190610976565b0190610b67565b565b91906199046198ff5f614802565b613c06565b916199238361991b6199158461090c565b9161090c565b101515614990565b619dc157506199396199345f614802565b614e87565b928061995561994f61994a5f614802565b610a33565b91610a33565b125f14619db357835b9061997a61996b34613a52565b61997484613a52565b90619529565b948561999661999061998b5f614802565b610a33565b91610a33565b145f14619d18575b816199b96199b36199ae5f614802565b610a33565b91610a33565b125f14619bcc57619a67619a57619a6f925b619a1186619a0b5f6199fb6199ed6101046199e761010b610fc5565b90610f4c565b6199f56174bc565b90610f96565b0191619a0683610fc5565b6157d6565b906139dc565b619a1c610101612c2c565b619a36619a30619a2b5f614984565b61096a565b9161096a565b14619baa575b619a5181619a4b610112610fc5565b906157b4565b906157d6565b619a6160016144f0565b906157d6565b6101116139dc565b619a7761aa6a565b90619a818561aa7e565b619a8a8461a310565b619a9561010b610fc5565b91619b0b619aaa619aa46174bc565b95613a52565b915f1993969790619abc610124610fc5565b91619af9619af3619aed7f1d1f406c89e99504a7222aaba50bf96b9f91f522ebc6ea825255145919e102ec99610f30565b99610f8a565b99619801565b99619b026108fa565b9687968761981d565b0390a480619b29619b23619b1e5f614802565b610a33565b91610a33565b13619b32575b50565b619b6e5f80619b3f6174bc565b619b488561483a565b619b506108fa565b9081619b5b81617ad6565b03925af1619b67617afe565b5015614990565b15619b2f57619b84619b7e6174bc565b9161483a565b90619ba6619b906108fa565b928392630aa7db6360e11b8452600484016198be565b0390fd5b619bc7619bbf82619bb96197f4565b90615752565b61010f6139dc565b619a3c565b619c6290619c12619bf0619beb610119619be58761483a565b90613528565b610fc5565b619c02619bfc5f61481e565b9161090c565b14619c0c8561483a565b906196aa565b619c1a6174bc565b906020619c30619c2b61012a6196ec565b6121e0565b636352211e90619c57619c428861483a565b92619c4b6108fa565b97889485938493613df3565b835260048301610b74565b03915afa908115619d1357619cb9619a6793619c95619c8f619a5795619a6f985f91619ce5575b5061096a565b9161096a565b14619ca161012a6196ec565b619caa8861483a565b90619cb36174bc565b926197c0565b619ce0619cc660016144f0565b619cdb610119619cd58961483a565b90613528565b6139dc565b6199cb565b619d06915060203d8111619d0c575b619cfe8183610d34565b8101906196f9565b5f619c89565b503d619cf4565b613e5f565b85619d33619d2d619d285f614802565b610a33565b91610a33565b135f14619d8b57619d4f619d48610113610fc5565b3a90615752565b619d6a619d64619d5e8961483a565b9261090c565b9161090c565b1115619d76575b61999e565b91509350619d835f614802565b933491619d71565b823490619daf619d996108fa565b92839263814ac7ff60e01b8452600484016195eb565b0390fd5b619dbc84614528565b61995e565b82619ddc5f9283926335465b3160e01b84526004840161242c565b0390fd5b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930090565b619e0c613c02565b50619e1561aa96565b90565b619e24816101256139dc565b619e5a7f07417920574ce0bdfe987af0575c8793cc73a29d7830760ad459d0e569b5b79b91619e516108fa565b91829182610b74565b0390a1565b619f4390619e6b613c02565b50619e7e619e7761aacf565b8290616f16565b619e89610102612c2c565b619ea3619e9d619e985f614984565b61096a565b9161096a565b145f14619fef57619eb5610117610fc5565b5b90619ebf61ab2c565b90619ec8613c02565b5082619ee4619ede619ed98561a026565b61090c565b9161090c565b115f14619f5757619efd90619ef883613a52565b613a6e565b9182619f19619f13619f0e5f614802565b610a33565b91610a33565b135f14619f4657619f3c92619f31619f37929161483a565b90613d0f565b613d50565b5b5b61a86e565b90565b505050619f525f61481e565b619f3d565b809150619f74619f6e619f695f614802565b610a33565b91610a33565b13155f14619fa557619f99619f94619f8f619f9f9493616a77565b61483a565b61a026565b9061451a565b5b619f3e565b619fb1619fb69161483a565b61a026565b80619fc9619fc38461090c565b9161090c565b105f14619fdf57619fd99161450c565b5b619fa0565b5050619fea5f61481e565b619fda565b619ffa610116610fc5565b619eb6565b61a007613c02565b5061a010613a4e565b5061a01961aacf565b9061a02261ab65565b9190565b61a04761a0589161a035613c02565b5061a04161011b610fc5565b90613d0f565b61a052610125610fc5565b90613d50565b90565b61a067816101106139dc565b61a09d7fb6f6af60099e44041a78b3561ed029b98bf03fdb0efbbb2eb15e1f3d7d9230379161a0946108fa565b91829182610b74565b0390a1565b61a0ae8161010e6139dc565b61a0e47ffdf6043c89a5f304289122dcc9f8bd78bb111b5d4f409e2fc2e6c141a1110b799161a0db6108fa565b91829182610b74565b0390a1565b61a0f1613c02565b5061a11161a100610125610fc5565b61a10b61010e610fc5565b90613d50565b90565b5f90565b61a12061a114565b5061a13961a13461a12f6170cf565b61a14c565b61ac53565b90565b5f90565b61a14990611be4565b90565b61a15e9061a15861a13c565b5061a140565b90565b5d565b9061a16e9161ae50565b565b61a1999161a17c613c02565b508161a19061a18a8361090c565b9161090c565b1191909161b128565b90565b61a1a4613c02565b5061a1ad61b14c565b90565b90565b5190565b5f7f4d65737361676520697320746f6f206c6f6e672e000000000000000000000000910152565b61a1eb60146020926129ee565b61a1f48161a1b7565b0190565b919061a21b90602061a213604086018681035f88015261a1de565b940190610b67565b565b1561a2255750565b61a2479061a2316108fa565b91829163271d43ff60e21b83526004830161a1f8565b0390fd5b60207f207368616c6c206265204554482e000000000000000000000000000000000000917f5468652066697273742062696420696e20612062696464696e6720726f756e645f8201520152565b61a2a5602e6040926129ee565b61a2ae8161a24b565b0190565b61a2c79060208101905f81830391015261a298565b90565b1561a2d157565b61a2d96108fa565b63b5a45a4960e01b81528061a2f06004820161a2b2565b0390fd5b61a2fd9061090c565b5f19811461a30b5760010190565b614a14565b61a35a9061a35461a34f61a32b61a3268461a1b0565b61a1b3565b61a34761a34161a33c61011a610fc5565b61090c565b9161090c565b11159261a1b0565b61a1b3565b9061a21d565b61a365610101612c2c565b61a37f61a37961a3745f614984565b61096a565b9161096a565b145f1461a4c75761a38e61b222565b61a3aa3461a3a461a39e5f61481e565b9161090c565b1161a2ca565b61a3b6426101146139dc565b61a3d361a3cb4261a3c561456a565b906157d6565b6101246139dc565b61a3de61010b610fc5565b429061a41f61a40d7f028a52641badd593b7f30072734c0b97e449213f55b5c3663756427340accd3c92610f30565b9261a4166108fa565b91829182610b74565b0390a25b61a43661a42e6174bc565b610101613c8a565b61a49161a45861a45361010361a44d61010b610fc5565b90611592565b616c63565b5f61a48a61a467828401610fc5565b61a48561a4726174bc565b61a480600187018490613956565b613c8a565b61a2f4565b91016139dc565b61a4c542600261a4bf61a4b161010461a4ab61010b610fc5565b90610f4c565b61a4b96174bc565b90610f96565b016139dc565b565b61a4cf6175e7565b61a4d761b160565b61a423565b61a4e461b27c565b565b61a4f261a4f791610fac565b610f30565b90565b61a50e61a50961a513926144ed565b610f2d565b617c27565b90565b61a5359061a52f61a52961a53a94617c27565b9161090c565b90610c9c565b61090c565b90565b90565b61a55461a54f61a5599261a53d565b610f2d565b617c27565b90565b61a57b9061a57561a56f61a58094617c27565b9161090c565b906141ae565b61090c565b90565b90565b61a59a61a59561a59f9261a583565b610f2d565b617c27565b90565b90565b61a5b961a5b461a5be9261a5a2565b610f2d565b617c27565b90565b61a5c9613c02565b5061a5fa61a5ea61a5e44361a5de60016144f0565b906163f5565b4061a4e6565b61a5f4600161a4fa565b9061a516565b61a60e4861a608604061a540565b9061a55c565b1861a61761b3c9565b9061a675575b5061a62661b73f565b9061a65a575b5061a63561b8d9565b9061a63f575b5090565b61a6539061a64d60c061a5a5565b9061a55c565b185f61a63b565b61a66e9061a668608061a586565b9061a55c565b185f61a62c565b61a69561a69a9161a684614fc5565b5061a68f60016144f0565b906163f5565b61b53d565b9061a6a5575b61a61d565b61a6ae9061a4e6565b185f61a6a0565b61a6e95f61a6ee9261a6c5613c02565b5061a6e382820161a6dd61a6d882617ca5565b6190d6565b906177dc565b01617ca5565b61b9f5565b90565b634e487b715f526020526024601cfd5b61a709613c02565b5061a712613c02565b5061a71e610101612c2c565b61a73861a73261a72d5f614984565b61096a565b9161096a565b145f1461a8165761a75b61a74d61010f610fc5565b9161a756613a7c565b616f16565b8061a77661a77061a76b5f614802565b610a33565b91610a33565b13155f1461a78457505b5b90565b9061a7ad61a79d8261a797610110610fc5565b90613d50565b61a7a760016144f0565b9061451a565b9161a7b661a0e9565b9061a7c08161483a565b61a7d261a7cc8461090c565b9161090c565b105f1461a80e579061a7fd61a8029261a7f761a7f161a808978761450c565b9161483a565b90613d0f565b613d50565b9061450c565b5b61a780565b50505061a809565b5061a822610111610fc5565b61a781565b9061a83d61a83761a8449361090c565b9161090c565b900a61090c565b90565b61a8669061a86061a85a61a86b9461090c565b9161090c565b90610c9c565b61090c565b90565b91909161a879613c02565b509161a886610101612c2c565b61a8a061a89a61a8955f614984565b61096a565b9161096a565b0361a8a9575b50565b61a8c261a8b46152e2565b9161a8bd6148a5565b613a6e565b9061a8d661a8cf82613a52565b8390613a6e565b918261a8f261a8ec61a8e75f614802565b610a33565b91610a33565b1361a8ff575b505061a8a6565b61a97d61a96c61a95b61a94a61a9ac97989661a9a69661a9849661a93361a92d61a9285f614802565b610a33565b91610a33565b125f1461a9b55761a9449150613a52565b5b61483a565b61a955610136610fc5565b90613d0f565b61a966610125610fc5565b90613d50565b61a977610137610fc5565b9061a827565b8390613d0f565b61a9a061a992610137610fc5565b61a99a61586b565b90613d0f565b9061a847565b9061451a565b905f808061a8f8565b5061a945565b61a9c3613952565b5061a9de5f61a9d861a9d3614ffc565b61ba0a565b01612c2c565b90565b9061a9eb8261ba0d565b8161aa167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b91610f8a565b9061aa1f6108fa565b8061aa29816109ec565b0390a261aa358161a1b3565b61aa4761aa415f61481e565b9161090c565b115f1461aa5b5761aa579161ba98565b505b565b505061aa6561ba62565b61aa59565b61aa72613c02565b5061aa7b61bac7565b90565b61aa879061badb565b565b61aa9360036154cf565b90565b61aa9e613c02565b5061aacc61aabe61aaad61aa89565b61aab861011b610fc5565b90613d0f565b61aac6614f1a565b90613d50565b90565b61aad7613a4e565b5061aafd61aae442613a52565b61aaf761aaf2610114610fc5565b613a52565b90613a6e565b90565b90565b61ab1761ab1261ab1c9261ab00565b610f2d565b61090c565b90565b61ab29600c61ab03565b90565b61ab34613c02565b5061ab6261ab5461ab4361ab1f565b61ab4e610125610fc5565b90613d0f565b61ab5c614f1a565b90613d50565b90565b61ab6d613c02565b5061ab79610102612c2c565b61ab9361ab8d61ab885f614984565b61096a565b9161096a565b145f1461ac435761aba5610117610fc5565b5b8061abb961abb35f61481e565b9161090c565b1461ac365761abc961011b610fc5565b61abd161ab2c565b918161abe561abdf5f61481e565b9161090c565b1461ac315761ac2e929161ac2461ac0a61ac299361ac04610125610fc5565b90613d0f565b61ac1e8361ac1860016144f0565b9061450c565b9061451a565b613d50565b61bd0e565b90565b505090565b5061ac405f61481e565b90565b61ac4e610116610fc5565b61aba6565b61ac5b61a114565b505c90565b91602061ac8192949361ac7a60408201965f830190610976565b0190610b67565b565b61ac9761ac9261ac9c926144c1565b610f2d565b617c27565b90565b67ffffffffffffffff811161acb75760208091020190565b610d20565b9061acce61acc98361ac9f565b610d5d565b918252565b61acdd6040610d5d565b90565b5f90565b61acec61acd3565b906020808361acf961782b565b81520161ad0461ace0565b81525050565b61ad1261ace4565b90565b5f5b82811061ad2357505050565b60209061ad2e61ad0a565b818401520161ad17565b9061ad5d61ad458361acbc565b9260208061ad53869361ac9f565b920191039061ad15565b565b5190565b9061ad6d8261ad5f565b81101561ad7e576020809102010190565b612bcf565b9061ad8d90610a33565b9052565b60209181520190565b60200190565b61ada990610a33565b9052565b9060208061adcf9361adc55f8201515f860190617956565b015191019061ada0565b565b9061adde8160409361adad565b0190565b60200190565b9061ae0561adff61adf88461ad5f565b809361ad91565b9261ad9a565b905f5b81811061ae155750505090565b90919261ae2e61ae28600192865161add1565b9461ade2565b910191909161ae08565b61ae4d9160208201915f81840391015261ade8565b90565b908061ae6461ae5e5f61481e565b9161090c565b115f1461b0895761ae748161bd3a565b9061ae80610101612c2c565b908161ae9c61ae9661ae915f614984565b61096a565b9161096a565b145f1461b06e5761af7061af5361af4e61aebf61aeb9600261ac83565b5b617c68565b9361af2061af0461aeff61aed28861ad38565b9a61aefa8c5f61aef361aee36174bc565b9261aeed8361481e565b9061ad63565b51016178ce565b613a52565b616a77565b602061af198b61af135f61481e565b9061ad63565b510161ad83565b61af4761af2b6174bc565b5f61af408b61af3a60016144f0565b9061ad63565b51016178ce565b869061450c565b613a52565b602061af698761af6360016144f0565b9061ad63565b510161ad83565b61af8361af7d60026144c4565b9161090c565b1161b01e575b505061af9e61af99610129617f14565b61384c565b9063b355121490823b1561b0195761afd59261afca5f809461afbe6108fa565b96879586948593613df3565b83526004830161ae38565b03925af1801561b0145761afe8575b505b565b61b007905f3d811161b00d575b61afff8183610d34565b810190613df9565b5f61afe4565b503d61aff5565b613e5f565b613def565b61b0679161b04561b04a925f61b03e8761b03860026144c4565b9061ad63565b51016178ce565b613a52565b602061b0608461b05a60026144c4565b9061ad63565b510161ad83565b5f8061af89565b61af7061af5361af4e61aebf61b0846003617c2d565b61aeba565b5061b09d61b098610129617f14565b61384c565b90639dc29fac9061b0ac6174bc565b9092803b1561b1235761b0d25f809461b0dd61b0c66108fa565b97889687958694613df3565b84526004840161ac60565b03925af1801561b11e5761b0f2575b5061afe6565b61b111905f3d811161b117575b61b1098183610d34565b810190613df9565b5f61b0ec565b503d61b0ff565b613e5f565b613def565b61b14261b148929361b138613c02565b508094189161bd6e565b90613d0f565b1890565b61b154613c02565b5061b15d61ab65565b90565b61b18661b17e61b16e614f25565b61b179610124610fc5565b61451a565b6101246139dc565b565b60207f20616374697665207965742e0000000000000000000000000000000000000000917f5468652063757272656e742062696464696e6720726f756e64206973206e6f745f8201520152565b61b1e2602c6040926129ee565b61b1eb8161b188565b0190565b91604061b22092949361b21961b20e606083018381035f85015261b1d5565b966020830190610b67565b0190610b67565b565b61b22d61010d610fc5565b61b24b4261b24361b23d8461090c565b9161090c565b101515614990565b61b2525750565b429061b27561b25f6108fa565b9283926302dbf17b60e31b84526004840161b1ef565b0390fd5b90565b61b2c861b29e61b29961013461b29361010b610fc5565b906123ef565b61b279565b61b2b461b2ac610107610fc5565b5f83016139dc565b600161b2c161010a610fc5565b91016139dc565b565b61b2de61b2d961b2e392613d1d565b610f2d565b61095f565b90565b61b2ef9061b2ca565b90565b61b2fb90610f62565b90565b61b3079061b2f2565b90565b61b31c61b317606461b2e6565b61b2fe565b90565b61b32890610f7e565b90565b90565b61b34261b33d61b3479261b32b565b610f2d565b61090c565b90565b60207f642e000000000000000000000000000000000000000000000000000000000000917f4172625379732e617262426c6f636b4e756d6265722063616c6c206661696c655f8201520152565b61b3a460226040926129ee565b61b3ad8161b34a565b0190565b61b3c69060208101905f81830391015261b397565b90565b61b3d161a114565b5061b3da613c02565b61b3e2617af9565b505f8061b3f561b3f061b30a565b61b31f565b600461b42c63a3b1b31d60e01b61b41d61b40d6108fa565b93849260208401908152016109ec565b60208201810382520382610d34565b82602082019151925af19161b43f617afe565b8361b492575b5061b4508315614990565b61b457575b565b7fa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d61b4806108fa565b8061b48a8161b3b1565b0390a161b455565b909161b49d8261a1b3565b61b4b061b4aa602061b32e565b9161090c565b145f1461b4da575061b4d290602061b4c78261a1b3565b818301019101617929565b905b5f61b445565b919250505f9161b4d4565b5f7f4172625379732e617262426c6f636b486173682063616c6c206661696c65642e910152565b61b518602080926129ee565b61b5218161b4e5565b0190565b61b53a9060208101905f81830391015261b50c565b90565b61b54561a114565b505f8061b550614fc5565b9261b559617af9565b50600461b5a361b56f61b56a61b30a565b61b31f565b9261b5946315a03d4160e11b9161b5846108fa565b9485936020850190815201610b74565b60208201810382520382610d34565b82602082019151925af19161b5b6617afe565b8361b609575b5061b5c78315614990565b61b5ce575b565b7fa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d61b5f76108fa565b8061b6018161b525565b0390a161b5cc565b909161b6148261a1b3565b61b62761b621602061b32e565b9161090c565b145f1461b651575061b64990602061b63e8261a1b3565b81830101910161931f565b905b5f61b5bc565b919250505f9161b64b565b90565b61b67361b66e61b6789261b65c565b610f2d565b61095f565b90565b61b6849061b65f565b90565b61b69090610f62565b90565b61b69c9061b687565b90565b61b6b161b6ac606c61b67b565b61b693565b90565b61b6bd90610f7e565b90565b60207f696c65642e000000000000000000000000000000000000000000000000000000917f417262476173496e666f2e6765744761734261636b6c6f672063616c6c2066615f8201520152565b61b71a60256040926129ee565b61b7238161b6c0565b0190565b61b73c9060208101905f81830391015261b70d565b90565b61b74761a114565b5061b750613c02565b61b758617af9565b505f8061b76b61b76661b69f565b61b6b4565b600461b7a162eadae160e51b61b79261b7826108fa565b93849260208401908152016109ec565b60208201810382520382610d34565b82602082019151925af19161b7b4617afe565b8361b807575b5061b7c58315614990565b61b7cc575b565b7fa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d61b7f56108fa565b8061b7ff8161b727565b0390a161b7ca565b909161b8128261a1b3565b61b82561b81f602061b32e565b9161090c565b145f1461b84f575061b84790602061b83c8261a1b3565b818301019101617929565b905b5f61b7ba565b919250505f9161b849565b60207f655570646174652063616c6c206661696c65642e000000000000000000000000917f417262476173496e666f2e6765744c3150726963696e67556e69747353696e635f8201520152565b61b8b460346040926129ee565b61b8bd8161b85a565b0190565b61b8d69060208101905f81830391015261b8a7565b90565b61b8e161a114565b5061b8ea613c02565b61b8f2617af9565b505f8061b90561b90061b69f565b61b6b4565b600461b93c6377f8098360e11b61b92d61b91d6108fa565b93849260208401908152016109ec565b60208201810382520382610d34565b82602082019151925af19161b94f617afe565b8361b9a2575b5061b9608315614990565b61b967575b565b7fa0f59128cf0144d4891de440cb7fffc98340af3ee8b041d4add05fc0248c392d61b9906108fa565b8061b99a8161b8c1565b0390a161b965565b909161b9ad8261a1b3565b61b9c061b9ba602061b32e565b9161090c565b145f1461b9ea575061b9e290602061b9d78261a1b3565b818301019101617929565b905b5f61b955565b919250505f9161b9e4565b61ba079061ba01613c02565b5061bd7c565b90565b90565b803b61ba2161ba1b5f61481e565b9161090c565b1461ba435761ba41905f61ba3b61ba36614ffc565b61ba0a565b01613c8a565b565b61ba5e905f918291634c9c8ce360e01b835260048301610983565b0390fd5b3461ba7561ba6f5f61481e565b9161090c565b1161ba7c57565b5f63b398979f60e01b81528061ba94600482016109ec565b0390fd5b5f8061bac49361baa6617af9565b508390602081019051915af49061babb617afe565b9091909161bd8e565b90565b61bacf613c02565b5061bad861ab65565b90565b8061baee61bae85f61481e565b9161090c565b1161baf7575b50565b61bb008161bd3a565b61bb0b610101612c2c565b8061bb2661bb2061bb1b5f614984565b61096a565b9161096a565b145f1461bbd6575061bb4161bb3c610129617f14565b61384c565b9161bb596340c10f199261bb536174bc565b9261450c565b92803b1561bbd15761bb7e5f809461bb8961bb726108fa565b97889687958694613df3565b84526004840161ac60565b03925af1801561bbcc5761bba0575b505b5f61baf4565b61bbbf905f3d811161bbc5575b61bbb78183610d34565b810190613df9565b5f61bb98565b503d61bbad565b613e5f565b613def565b9061bc5c61bc799261bc4061bc2461bbf661bbf160026144c4565b617e88565b9661bc1d61bc026174bc565b5f61bc168b61bc108361481e565b90617eb3565b51016178ce565b859061450c565b602061bc398861bc335f61481e565b90617eb3565b51016177dc565b5f61bc558661bc4f60016144f0565b90617eb3565b51016178ce565b602061bc728461bc6c60016144f0565b90617eb3565b51016177dc565b61bc8c61bc87610129617f14565b61384c565b9063b33266da90823b1561bd095761bcc39261bcb85f809461bcac6108fa565b96879586948593613df3565b835260048301617fbb565b03925af1801561bd045761bcd8575b5061bb9a565b61bcf7905f3d811161bcfd575b61bcef8183610d34565b810190613df9565b5f61bcd2565b503d61bce5565b613e5f565b613def565b61bd379161bd1a613c02565b508161bd2e61bd288361090c565b9161090c565b1091909161b128565b90565b61bd5b61bd6b9161bd49613c02565b5061bd55610138610fc5565b90613d0f565b61bd656064613d20565b90613d50565b90565b61bd76613c02565b50151590565b61bd84613c02565b505f5260205f2090565b9061bda29061bd9b617af9565b5015614990565b5f1461bdae575061be12565b61bdb78261a1b3565b61bdc961bdc35f61481e565b9161090c565b148061bdf7575b61bdd8575090565b61bdf3905f918291639996b31560e01b835260048301610983565b0390fd5b50803b61be0c61be065f61481e565b9161090c565b1461bdd0565b61be1b8161a1b3565b61be2d61be275f61481e565b9161090c565b115f1461be3c57805190602001fd5b5f63d6bda27560e01b81528061be54600482016109ec565b0390fdfea26469706673582212203b120a08ff23404b5252be9157a28bc2a118a4810cdefb9c8712cfedcca5ad7364736f6c63430008220033",
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

// GetBidCstRewardAmountPerMainPrizeTimeIncrement is a free data retrieval call binding the contract method 0xdd5f6587.
//
// Solidity: function getBidCstRewardAmountPerMainPrizeTimeIncrement() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetBidCstRewardAmountPerMainPrizeTimeIncrement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getBidCstRewardAmountPerMainPrizeTimeIncrement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBidCstRewardAmountPerMainPrizeTimeIncrement is a free data retrieval call binding the contract method 0xdd5f6587.
//
// Solidity: function getBidCstRewardAmountPerMainPrizeTimeIncrement() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetBidCstRewardAmountPerMainPrizeTimeIncrement() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetBidCstRewardAmountPerMainPrizeTimeIncrement(&_CosmicSignatureGameV3.CallOpts)
}

// GetBidCstRewardAmountPerMainPrizeTimeIncrement is a free data retrieval call binding the contract method 0xdd5f6587.
//
// Solidity: function getBidCstRewardAmountPerMainPrizeTimeIncrement() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetBidCstRewardAmountPerMainPrizeTimeIncrement() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetBidCstRewardAmountPerMainPrizeTimeIncrement(&_CosmicSignatureGameV3.CallOpts)
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

// GetCstDutchAuctionBeginningBidPriceMinLimit is a free data retrieval call binding the contract method 0xa375f482.
//
// Solidity: function getCstDutchAuctionBeginningBidPriceMinLimit() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Caller) GetCstDutchAuctionBeginningBidPriceMinLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmicSignatureGameV3.contract.Call(opts, &out, "getCstDutchAuctionBeginningBidPriceMinLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCstDutchAuctionBeginningBidPriceMinLimit is a free data retrieval call binding the contract method 0xa375f482.
//
// Solidity: function getCstDutchAuctionBeginningBidPriceMinLimit() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3Session) GetCstDutchAuctionBeginningBidPriceMinLimit() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetCstDutchAuctionBeginningBidPriceMinLimit(&_CosmicSignatureGameV3.CallOpts)
}

// GetCstDutchAuctionBeginningBidPriceMinLimit is a free data retrieval call binding the contract method 0xa375f482.
//
// Solidity: function getCstDutchAuctionBeginningBidPriceMinLimit() view returns(uint256)
func (_CosmicSignatureGameV3 *CosmicSignatureGameV3CallerSession) GetCstDutchAuctionBeginningBidPriceMinLimit() (*big.Int, error) {
	return _CosmicSignatureGameV3.Contract.GetCstDutchAuctionBeginningBidPriceMinLimit(&_CosmicSignatureGameV3.CallOpts)
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
