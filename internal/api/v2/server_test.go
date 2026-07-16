package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame/contractstate"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	rwmodel "github.com/PredictionExplorer/augur-explorer/internal/model/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

type fakeBidReader struct {
	page func(context.Context, int64, cgstore.BidPageCursor, int) ([]cgmodel.CGBidRec, bool, error)
	item func(context.Context, int64, int64) (cgmodel.CGBidRec, error)
}

type fakeRoundReader struct {
	page func(context.Context, *cgstore.RoundPageCursor, int) ([]cgmodel.CGRoundRec, bool, error)
	item func(context.Context, int64) (cgmodel.CGRoundRec, error)
}

type fakeCurrentRoundReader struct {
	statistics func(context.Context, int64) (cgmodel.CGRoundStats, error)
	bidCount   func(context.Context, int64) (int64, error)
}

type fakeRoundPrizeReader struct {
	exists func(context.Context, int64) (bool, error)
	page   func(context.Context, int64, *cgstore.PrizePageCursor, int) ([]cgmodel.CGPrizeHistory, bool, error)
}

type fakeRoundRaffleReader struct {
	exists func(context.Context, int64) (bool, error)
	eth    func(context.Context, int64, *cgstore.RaffleEthDepositPageCursor, int) ([]cgstore.RaffleEthDepositRecord, bool, error)
	nft    func(context.Context, int64, bool, *cgstore.RaffleNFTWinnerPageCursor, int) ([]cgmodel.CGRaffleNFTWinnerRec, bool, error)
}

type fakeRoundDonationReader struct {
	eth   func(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundEthDonationRecord, bool, error)
	erc20 func(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundERC20DonationRecord, bool, error)
	nft   func(context.Context, int64, *cgstore.DonationPageCursor, int) ([]cgstore.RoundNFTDonationRecord, bool, error)
}

type fakeStatisticsReader struct {
	global       func(context.Context) (cgstore.GlobalStatisticsRecord, error)
	counters     func(context.Context) (cgmodel.CGRecordCounters, error)
	roi          func(context.Context, int, cgstore.ROILeaderboardSort, *cgstore.ROILeaderboardPageCursor, int) ([]cgstore.ROILeaderboardRecord, bool, error)
	claims       func(context.Context, *cgstore.ClaimSummaryCursor, int) ([]cgstore.ClaimSummaryRecord, bool, error)
	exists       func(context.Context, int64) (bool, error)
	summary      func(context.Context, int64) (cgstore.ClaimSummaryRecord, error)
	transactions func(context.Context, int64, *cgstore.ClaimEventCursor, int) ([]cgstore.ClaimTransactionRecord, bool, error)
	attached     func(context.Context, int64, *cgstore.ClaimEventCursor, int) ([]cgstore.AttachedTokenRecord, bool, error)
	unclaimed    func(context.Context, int64, *cgstore.UnclaimedItemCursor, int) ([]cgstore.UnclaimedItemRecord, bool, error)
}

type fakeBiddingAnalyticsReader struct {
	frequency func(context.Context, int, int, int) ([]cgmodel.CGBidFrequencyBucket, error)
	ratio     func(context.Context, int, int, int) ([]cgmodel.CGBidTypeRatioBucket, error)
	periods   func(context.Context, int, int, int, int, int) ([]cgmodel.CGTopBidderInfo, []cgmodel.CGBidderActivePeriod, bool, error)
	bounds    func(context.Context) (int64, int64, error)
}

type fakeContractAddressReader struct {
	get func(context.Context) (cgmodel.CosmicGameContractAddrs, error)
}

type fakeParticipantReader struct {
	bidders           func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.BidderParticipantRecord, bool, error)
	winners           func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.WinnerParticipantRecord, bool, error)
	donors            func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.DonorParticipantRecord, bool, error)
	cstStakers        func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.CSTStakerParticipantRecord, bool, error)
	randomWalkStakers func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.RandomWalkStakerParticipantRecord, bool, error)
	dualStakers       func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.DualStakerParticipantRecord, bool, error)
}

type fakeUserReader struct {
	addressID func(context.Context, string) (int64, error)
	profile   func(context.Context, int64) (cgstore.UserProfileRecord, error)
	bids      func(context.Context, int64, *cgstore.UserBidPageCursor, int) ([]cgmodel.CGBidRec, bool, error)
}

type fakeUserHistoryReader struct {
	addressID     func(context.Context, string) (int64, error)
	prizes        func(context.Context, int64, *cgstore.UserPrizePageCursor, int) ([]cgmodel.CGPrizeHistory, bool, error)
	deposits      func(context.Context, int64, *bool, *cgstore.UserEventPageCursor, int) ([]cgstore.UserRaffleEthDepositRecord, bool, error)
	nftWins       func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserRaffleNftWinRecord, bool, error)
	ethDonations  func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.RoundEthDonationRecord, bool, error)
	ercDonations  func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.RoundERC20DonationRecord, bool, error)
	nftDonations  func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.RoundNFTDonationRecord, bool, error)
	donatedNfts   func(context.Context, int64, *bool, *cgstore.UserEventPageCursor, int) ([]cgstore.UserDonatedNftRecord, bool, error)
	donatedTokens func(context.Context, int64, *cgstore.UserDonatedErc20PageCursor, int) ([]cgstore.UserDonatedErc20Record, bool, error)
}

type fakeUserStakingReader struct {
	addressID     func(context.Context, string) (int64, error)
	cstActions    func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserStakingActionRecord, bool, error)
	rwalkActions  func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserStakingActionRecord, bool, error)
	cstStaked     func(context.Context, int64, *cgstore.UserStakingTokenPageCursor, int) ([]cgstore.UserStakedCstTokenRecord, bool, error)
	rwalkStaked   func(context.Context, int64, *cgstore.UserStakingTokenPageCursor, int) ([]cgstore.UserStakedRwalkTokenRecord, bool, error)
	deposits      func(context.Context, int64, *bool, *cgstore.UserStakingDepositPageCursor, int) ([]cgstore.UserStakingDepositRecord, bool, error)
	depositExists func(context.Context, int64) (bool, error)
	depositReward func(context.Context, int64, int64, *cgstore.UserStakingRewardPageCursor, int) ([]cgstore.UserStakingDepositRewardRecord, bool, error)
	tokenRewards  func(context.Context, int64, *cgstore.UserStakingTokenPageCursor, int) ([]cgstore.UserStakingTokenRewardRecord, bool, error)
	tokenExists   func(context.Context, int64) (bool, error)
	tokenDeposits func(context.Context, int64, int64, *cgstore.UserStakingTokenDepositPageCursor, int) ([]cgstore.UserStakingTokenRewardDepositRecord, bool, error)
}

type fakeUserActivityReader struct {
	addressID    func(context.Context, string) (int64, error)
	ownedTokens  func(context.Context, int64, *cgstore.UserTokenPageCursor, int) ([]cgstore.UserOwnedTokenRecord, bool, error)
	csTransfers  func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserCosmicSignatureTransferRecord, bool, error)
	ctTransfers  func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserCosmicTokenTransferRecord, bool, error)
	mktRewards   func(context.Context, int64, *cgstore.UserEventPageCursor, int) ([]cgstore.UserMarketingRewardRecord, bool, error)
	tokenSummary func(context.Context, int64) (cgstore.UserCosmicTokenSummaryRecord, error)
	pending      func(context.Context, int64) (cgstore.UserPendingWinningsRecord, error)
}

type fakeGlobalDirectoryReader struct {
	tokens          func(context.Context, cgstore.GlobalTokenFilter, *cgstore.GlobalTokenPageCursor, int) ([]cgstore.GlobalTokenRecord, bool, error)
	tokenDetail     func(context.Context, int64) (cgstore.GlobalTokenDetailRecord, error)
	tokenExists     func(context.Context, int64) (bool, error)
	nameHistory     func(context.Context, int64, *cgstore.TokenEventPageCursor, int) ([]cgstore.TokenNameChangeRecord, bool, error)
	tokenTransfers  func(context.Context, int64, *cgstore.TokenEventPageCursor, int) ([]cgstore.TokenTransferRecord, bool, error)
	csHolders       func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.CosmicSignatureHolderRecord, bool, error)
	ctHolders       func(context.Context, *cgstore.ParticipantPageCursor, int) ([]cgstore.CosmicTokenHolderRecord, bool, error)
	tokenStatistics func(context.Context) (cgstore.CosmicTokenStatisticsRecord, error)
	supplyByBid     func(context.Context, *cgstore.SupplyChangePageCursor, int) ([]cgstore.SupplyChangeRecord, bool, error)
	supplyDaily     func(context.Context, time.Time, time.Time) ([]cgstore.DailySupplyRecord, error)
	mktRewards      func(context.Context, *cgstore.UserEventPageCursor, int) ([]cgstore.MarketingRewardRecord, bool, error)
}

type fakeGlobalStakingReader struct {
	cstActions   func(context.Context, *cgstore.GlobalStakingActionPageCursor, int) ([]cgstore.GlobalStakingActionRecord, bool, error)
	rwalkActions func(context.Context, *cgstore.GlobalStakingActionPageCursor, int) ([]cgstore.GlobalStakingActionRecord, bool, error)
	cstAction    func(context.Context, int64) (cgmodel.CGStakeUnstakeCombined, error)
	rwalkAction  func(context.Context, int64) (cgmodel.CGStakeUnstakeCombined, error)
	cstStaked    func(context.Context, *cgstore.GlobalStakedTokenPageCursor, int) ([]cgstore.GlobalStakedCstTokenRecord, bool, error)
	rwalkStaked  func(context.Context, *cgstore.GlobalStakedTokenPageCursor, int) ([]cgstore.GlobalStakedRwalkTokenRecord, bool, error)
	deposits     func(context.Context, *cgstore.GlobalStakingDepositPageCursor, int) ([]cgstore.GlobalStakingDepositRecord, bool, error)
	roundExists  func(context.Context, int64) (bool, error)
	roundRewards func(context.Context, int64, *cgstore.RoundStakingRewardPageCursor, int) ([]cgstore.RoundStakingRewardRecord, bool, error)
	raffleWins   func(context.Context, bool, *cgstore.GlobalStakerRafflePageCursor, int) ([]cgmodel.CGRaffleNFTWinnerRec, bool, error)
}

func (f fakeGlobalDirectoryReader) CosmicSignatureTokensGlobalPage(
	ctx context.Context,
	filter cgstore.GlobalTokenFilter,
	after *cgstore.GlobalTokenPageCursor,
	limit int,
) ([]cgstore.GlobalTokenRecord, bool, error) {
	if f.tokens == nil {
		return []cgstore.GlobalTokenRecord{}, false, nil
	}
	return f.tokens(ctx, filter, after, limit)
}

func (f fakeGlobalDirectoryReader) CosmicSignatureTokenDetailV2(
	ctx context.Context,
	tokenID int64,
) (cgstore.GlobalTokenDetailRecord, error) {
	if f.tokenDetail == nil {
		return cgstore.GlobalTokenDetailRecord{}, store.ErrNotFound
	}
	return f.tokenDetail(ctx, tokenID)
}

func (f fakeGlobalDirectoryReader) CosmicSignatureTokenExists(
	ctx context.Context,
	tokenID int64,
) (bool, error) {
	if f.tokenExists == nil {
		return true, nil
	}
	return f.tokenExists(ctx, tokenID)
}

func (f fakeGlobalDirectoryReader) TokenNameHistoryPage(
	ctx context.Context,
	tokenID int64,
	after *cgstore.TokenEventPageCursor,
	limit int,
) ([]cgstore.TokenNameChangeRecord, bool, error) {
	if f.nameHistory == nil {
		return []cgstore.TokenNameChangeRecord{}, false, nil
	}
	return f.nameHistory(ctx, tokenID, after, limit)
}

func (f fakeGlobalDirectoryReader) TokenTransfersPage(
	ctx context.Context,
	tokenID int64,
	after *cgstore.TokenEventPageCursor,
	limit int,
) ([]cgstore.TokenTransferRecord, bool, error) {
	if f.tokenTransfers == nil {
		return []cgstore.TokenTransferRecord{}, false, nil
	}
	return f.tokenTransfers(ctx, tokenID, after, limit)
}

func (f fakeGlobalDirectoryReader) CosmicSignatureHoldersPage(
	ctx context.Context,
	after *cgstore.ParticipantPageCursor,
	limit int,
) ([]cgstore.CosmicSignatureHolderRecord, bool, error) {
	if f.csHolders == nil {
		return []cgstore.CosmicSignatureHolderRecord{}, false, nil
	}
	return f.csHolders(ctx, after, limit)
}

func (f fakeGlobalDirectoryReader) CosmicTokenHoldersPage(
	ctx context.Context,
	after *cgstore.ParticipantPageCursor,
	limit int,
) ([]cgstore.CosmicTokenHolderRecord, bool, error) {
	if f.ctHolders == nil {
		return []cgstore.CosmicTokenHolderRecord{}, false, nil
	}
	return f.ctHolders(ctx, after, limit)
}

func (f fakeGlobalDirectoryReader) CosmicTokenStatisticsV2(
	ctx context.Context,
) (cgstore.CosmicTokenStatisticsRecord, error) {
	if f.tokenStatistics == nil {
		return zeroCosmicTokenStatisticsRecord(), nil
	}
	return f.tokenStatistics(ctx)
}

// zeroCosmicTokenStatisticsRecord is a mapper-valid all-zero snapshot.
func zeroCosmicTokenStatisticsRecord() cgstore.CosmicTokenStatisticsRecord {
	return cgstore.CosmicTokenStatisticsRecord{
		TotalSupplyWei:             "0",
		BiddingRewardsWei:          "0",
		MainPrizesWei:              "0",
		RafflePrizesWei:            "0",
		ChronoWarriorPrizesWei:     "0",
		EnduranceChampionPrizesWei: "0",
		LastCstBidderPrizesWei:     "0",
		MarketingRewardsWei:        "0",
		TotalEarnedWei:             "0",
		ConsumedInBidsWei:          "0",
		NetWei:                     "0",
	}
}

func (f fakeGlobalDirectoryReader) CosmicTokenSupplyByBidPage(
	ctx context.Context,
	after *cgstore.SupplyChangePageCursor,
	limit int,
) ([]cgstore.SupplyChangeRecord, bool, error) {
	if f.supplyByBid == nil {
		return []cgstore.SupplyChangeRecord{}, false, nil
	}
	return f.supplyByBid(ctx, after, limit)
}

func (f fakeGlobalDirectoryReader) CosmicTokenSupplyDaily(
	ctx context.Context,
	from, to time.Time,
) ([]cgstore.DailySupplyRecord, error) {
	if f.supplyDaily == nil {
		return []cgstore.DailySupplyRecord{}, nil
	}
	return f.supplyDaily(ctx, from, to)
}

func (f fakeGlobalDirectoryReader) MarketingRewardsGlobalPage(
	ctx context.Context,
	after *cgstore.UserEventPageCursor,
	limit int,
) ([]cgstore.MarketingRewardRecord, bool, error) {
	if f.mktRewards == nil {
		return []cgstore.MarketingRewardRecord{}, false, nil
	}
	return f.mktRewards(ctx, after, limit)
}

func (f fakeGlobalStakingReader) GlobalCstStakingActionsPage(
	ctx context.Context,
	after *cgstore.GlobalStakingActionPageCursor,
	limit int,
) ([]cgstore.GlobalStakingActionRecord, bool, error) {
	if f.cstActions == nil {
		return []cgstore.GlobalStakingActionRecord{}, false, nil
	}
	return f.cstActions(ctx, after, limit)
}

func (f fakeGlobalStakingReader) GlobalRwalkStakingActionsPage(
	ctx context.Context,
	after *cgstore.GlobalStakingActionPageCursor,
	limit int,
) ([]cgstore.GlobalStakingActionRecord, bool, error) {
	if f.rwalkActions == nil {
		return []cgstore.GlobalStakingActionRecord{}, false, nil
	}
	return f.rwalkActions(ctx, after, limit)
}

func (f fakeGlobalStakingReader) StakeActionCstInfo(
	ctx context.Context,
	actionID int64,
) (cgmodel.CGStakeUnstakeCombined, error) {
	if f.cstAction == nil {
		return cgmodel.CGStakeUnstakeCombined{}, store.ErrNotFound
	}
	return f.cstAction(ctx, actionID)
}

func (f fakeGlobalStakingReader) StakeActionRwalkInfo(
	ctx context.Context,
	actionID int64,
) (cgmodel.CGStakeUnstakeCombined, error) {
	if f.rwalkAction == nil {
		return cgmodel.CGStakeUnstakeCombined{}, store.ErrNotFound
	}
	return f.rwalkAction(ctx, actionID)
}

func (f fakeGlobalStakingReader) GlobalStakedCstTokensPage(
	ctx context.Context,
	after *cgstore.GlobalStakedTokenPageCursor,
	limit int,
) ([]cgstore.GlobalStakedCstTokenRecord, bool, error) {
	if f.cstStaked == nil {
		return []cgstore.GlobalStakedCstTokenRecord{}, false, nil
	}
	return f.cstStaked(ctx, after, limit)
}

func (f fakeGlobalStakingReader) GlobalStakedRwalkTokensPage(
	ctx context.Context,
	after *cgstore.GlobalStakedTokenPageCursor,
	limit int,
) ([]cgstore.GlobalStakedRwalkTokenRecord, bool, error) {
	if f.rwalkStaked == nil {
		return []cgstore.GlobalStakedRwalkTokenRecord{}, false, nil
	}
	return f.rwalkStaked(ctx, after, limit)
}

func (f fakeGlobalStakingReader) GlobalStakingDepositsPage(
	ctx context.Context,
	after *cgstore.GlobalStakingDepositPageCursor,
	limit int,
) ([]cgstore.GlobalStakingDepositRecord, bool, error) {
	if f.deposits == nil {
		return []cgstore.GlobalStakingDepositRecord{}, false, nil
	}
	return f.deposits(ctx, after, limit)
}

func (f fakeGlobalStakingReader) CompletedRoundExists(
	ctx context.Context,
	round int64,
) (bool, error) {
	if f.roundExists == nil {
		return true, nil
	}
	return f.roundExists(ctx, round)
}

func (f fakeGlobalStakingReader) RoundStakingRewardsPage(
	ctx context.Context,
	round int64,
	after *cgstore.RoundStakingRewardPageCursor,
	limit int,
) ([]cgstore.RoundStakingRewardRecord, bool, error) {
	if f.roundRewards == nil {
		return []cgstore.RoundStakingRewardRecord{}, false, nil
	}
	return f.roundRewards(ctx, round, after, limit)
}

func (f fakeGlobalStakingReader) GlobalStakerRaffleNftWinsPage(
	ctx context.Context,
	isRwalk bool,
	after *cgstore.GlobalStakerRafflePageCursor,
	limit int,
) ([]cgmodel.CGRaffleNFTWinnerRec, bool, error) {
	if f.raffleWins == nil {
		return []cgmodel.CGRaffleNFTWinnerRec{}, false, nil
	}
	return f.raffleWins(ctx, isRwalk, after, limit)
}

type fakeRandomWalkReader struct {
	tokens       func(context.Context, rwstore.TokenFilter, rwstore.TokenSort, *rwstore.TokenPageCursor, int) ([]rwstore.TokenRecord, bool, error)
	tokenDetail  func(context.Context, int64) (rwstore.TokenDetailRecord, error)
	tokenExists  func(context.Context, int64) (bool, error)
	nameHistory  func(context.Context, int64, *rwstore.EventPageCursor, int) ([]rwstore.TokenNameChangeRecord, bool, error)
	tokenEvents  func(context.Context, int64, *rwstore.EventPageCursor, int) ([]rwstore.TokenEventRecord, bool, error)
	activeOffers func(context.Context, rwstore.OfferSort, *rwstore.OfferPageCursor, int) ([]rwstore.OfferRecord, bool, error)
	offerHistory func(context.Context, *rwstore.EventPageCursor, int) ([]rwstore.OfferHistoryRecord, bool, error)
	userOffers   func(context.Context, int64, *rwstore.EventPageCursor, int) ([]rwstore.OfferHistoryRecord, bool, error)
	trades       func(context.Context, *rwstore.EventPageCursor, int) ([]rwstore.TradeRecord, bool, error)
	floorPrice   func(context.Context) (rwstore.FloorPriceRecord, error)
	addressID    func(context.Context, string) (int64, error)
	profile      func(context.Context, int64) (rwstore.UserProfileRecord, error)
	userTokens   func(context.Context, int64, *rwstore.TokenPageCursor, int) ([]rwstore.OwnedTokenRecord, bool, error)
	statistics   func(context.Context) (rwstore.StatisticsRecord, error)
	volumeSeries func(context.Context, int, int, int) (string, []rwstore.VolumeBucketRecord, error)
	floorSeries  func(context.Context, int, int, int) ([]rwstore.FloorPointRecord, error)
	mintReport   func(context.Context) ([]rwstore.MonthlyMintRecord, error)
	withdrawals  func(context.Context, *rwstore.EventPageCursor, int) ([]rwstore.WithdrawalRecord, bool, error)
	addrs        func(context.Context) (rwmodel.ContractAddresses, error)
}

func (f fakeRandomWalkReader) TokensPage(
	ctx context.Context,
	filter rwstore.TokenFilter,
	sort rwstore.TokenSort,
	after *rwstore.TokenPageCursor,
	limit int,
) ([]rwstore.TokenRecord, bool, error) {
	if f.tokens == nil {
		return []rwstore.TokenRecord{}, false, nil
	}
	return f.tokens(ctx, filter, sort, after, limit)
}

func (f fakeRandomWalkReader) TokenDetailV2(
	ctx context.Context,
	tokenID int64,
) (rwstore.TokenDetailRecord, error) {
	if f.tokenDetail == nil {
		return rwstore.TokenDetailRecord{}, store.ErrNotFound
	}
	return f.tokenDetail(ctx, tokenID)
}

func (f fakeRandomWalkReader) CollectionTokenExists(
	ctx context.Context,
	tokenID int64,
) (bool, error) {
	if f.tokenExists == nil {
		return true, nil
	}
	return f.tokenExists(ctx, tokenID)
}

func (f fakeRandomWalkReader) TokenNameChangesPageV2(
	ctx context.Context,
	tokenID int64,
	after *rwstore.EventPageCursor,
	limit int,
) ([]rwstore.TokenNameChangeRecord, bool, error) {
	if f.nameHistory == nil {
		return []rwstore.TokenNameChangeRecord{}, false, nil
	}
	return f.nameHistory(ctx, tokenID, after, limit)
}

func (f fakeRandomWalkReader) TokenEventsPage(
	ctx context.Context,
	tokenID int64,
	after *rwstore.EventPageCursor,
	limit int,
) ([]rwstore.TokenEventRecord, bool, error) {
	if f.tokenEvents == nil {
		return []rwstore.TokenEventRecord{}, false, nil
	}
	return f.tokenEvents(ctx, tokenID, after, limit)
}

func (f fakeRandomWalkReader) ActiveOffersPage(
	ctx context.Context,
	sort rwstore.OfferSort,
	after *rwstore.OfferPageCursor,
	limit int,
) ([]rwstore.OfferRecord, bool, error) {
	if f.activeOffers == nil {
		return []rwstore.OfferRecord{}, false, nil
	}
	return f.activeOffers(ctx, sort, after, limit)
}

func (f fakeRandomWalkReader) OfferHistoryPage(
	ctx context.Context,
	after *rwstore.EventPageCursor,
	limit int,
) ([]rwstore.OfferHistoryRecord, bool, error) {
	if f.offerHistory == nil {
		return []rwstore.OfferHistoryRecord{}, false, nil
	}
	return f.offerHistory(ctx, after, limit)
}

func (f fakeRandomWalkReader) UserOffersPage(
	ctx context.Context,
	userAid int64,
	after *rwstore.EventPageCursor,
	limit int,
) ([]rwstore.OfferHistoryRecord, bool, error) {
	if f.userOffers == nil {
		return []rwstore.OfferHistoryRecord{}, false, nil
	}
	return f.userOffers(ctx, userAid, after, limit)
}

func (f fakeRandomWalkReader) TradesPage(
	ctx context.Context,
	after *rwstore.EventPageCursor,
	limit int,
) ([]rwstore.TradeRecord, bool, error) {
	if f.trades == nil {
		return []rwstore.TradeRecord{}, false, nil
	}
	return f.trades(ctx, after, limit)
}

func (f fakeRandomWalkReader) FloorPriceV2(ctx context.Context) (rwstore.FloorPriceRecord, error) {
	if f.floorPrice == nil {
		return rwstore.FloorPriceRecord{}, nil
	}
	return f.floorPrice(ctx)
}

func (f fakeRandomWalkReader) UserAddressID(ctx context.Context, address string) (int64, error) {
	if f.addressID == nil {
		return 1, nil
	}
	return f.addressID(ctx, address)
}

func (f fakeRandomWalkReader) UserProfileV2(
	ctx context.Context,
	userAid int64,
) (rwstore.UserProfileRecord, error) {
	if f.profile == nil {
		return rwstore.UserProfileRecord{
			Aid:              userAid,
			Address:          "0x0000000000000000000000000000000000000001",
			TradingVolumeWei: "0",
			ProfitWei:        "0",
		}, nil
	}
	return f.profile(ctx, userAid)
}

func (f fakeRandomWalkReader) UserTokensPage(
	ctx context.Context,
	userAid int64,
	after *rwstore.TokenPageCursor,
	limit int,
) ([]rwstore.OwnedTokenRecord, bool, error) {
	if f.userTokens == nil {
		return []rwstore.OwnedTokenRecord{}, false, nil
	}
	return f.userTokens(ctx, userAid, after, limit)
}

func (f fakeRandomWalkReader) StatisticsV2(ctx context.Context) (rwstore.StatisticsRecord, error) {
	if f.statistics == nil {
		return rwstore.StatisticsRecord{
			TokenTradingVolumeWei:  "0",
			MintFundsWei:           "0",
			MarketTradingVolumeWei: "0",
		}, nil
	}
	return f.statistics(ctx)
}

func (f fakeRandomWalkReader) TradingVolumeSeries(
	ctx context.Context,
	from, to, interval int,
) (string, []rwstore.VolumeBucketRecord, error) {
	if f.volumeSeries == nil {
		return "0", []rwstore.VolumeBucketRecord{}, nil
	}
	return f.volumeSeries(ctx, from, to, interval)
}

func (f fakeRandomWalkReader) ListingFloorSeries(
	ctx context.Context,
	from, to, interval int,
) ([]rwstore.FloorPointRecord, error) {
	if f.floorSeries == nil {
		return []rwstore.FloorPointRecord{}, nil
	}
	return f.floorSeries(ctx, from, to, interval)
}

func (f fakeRandomWalkReader) MintReportV2(ctx context.Context) ([]rwstore.MonthlyMintRecord, error) {
	if f.mintReport == nil {
		return []rwstore.MonthlyMintRecord{}, nil
	}
	return f.mintReport(ctx)
}

func (f fakeRandomWalkReader) WithdrawalsPage(
	ctx context.Context,
	after *rwstore.EventPageCursor,
	limit int,
) ([]rwstore.WithdrawalRecord, bool, error) {
	if f.withdrawals == nil {
		return []rwstore.WithdrawalRecord{}, false, nil
	}
	return f.withdrawals(ctx, after, limit)
}

func (f fakeRandomWalkReader) ContractAddrs(ctx context.Context) (rwmodel.ContractAddresses, error) {
	if f.addrs == nil {
		return rwmodel.ContractAddresses{
			MarketPlace:    "0x1200000000000000000000000000000000000012",
			RandomWalk:     "0x8000000000000000000000000000000000000008",
			MarketPlaceAid: 12,
			RandomWalkAid:  8,
		}, nil
	}
	return f.addrs(ctx)
}

type fakeContractState struct {
	snapshot func() contractstate.Snapshot
}

func (f fakeRoundReader) PrizeClaimsPage(ctx context.Context, after *cgstore.RoundPageCursor, limit int) ([]cgmodel.CGRoundRec, bool, error) {
	if f.page == nil {
		return []cgmodel.CGRoundRec{}, false, nil
	}
	return f.page(ctx, after, limit)
}

func (f fakeRoundReader) RoundInfo(ctx context.Context, round int64) (cgmodel.CGRoundRec, error) {
	if f.item == nil {
		return cgmodel.CGRoundRec{}, store.ErrNotFound
	}
	return f.item(ctx, round)
}

func (f fakeBidReader) BidsByRoundPage(ctx context.Context, round int64, after cgstore.BidPageCursor, limit int) ([]cgmodel.CGBidRec, bool, error) {
	if f.page == nil {
		return []cgmodel.CGBidRec{}, false, nil
	}
	return f.page(ctx, round, after, limit)
}

func (f fakeBidReader) BidByRoundAndPosition(ctx context.Context, round, position int64) (cgmodel.CGBidRec, error) {
	if f.item == nil {
		return cgmodel.CGBidRec{}, store.ErrNotFound
	}
	return f.item(ctx, round, position)
}

func (f fakeCurrentRoundReader) CosmicGameRoundStatistics(ctx context.Context, round int64) (cgmodel.CGRoundStats, error) {
	if f.statistics == nil {
		return cgmodel.CGRoundStats{RoundNum: round}, nil
	}
	return f.statistics(ctx, round)
}

func (f fakeCurrentRoundReader) BidCountForRound(ctx context.Context, round int64) (int64, error) {
	if f.bidCount == nil {
		return 0, nil
	}
	return f.bidCount(ctx, round)
}

func (f fakeRoundPrizeReader) CompletedRoundExists(ctx context.Context, round int64) (bool, error) {
	if f.exists == nil {
		return true, nil
	}
	return f.exists(ctx, round)
}

func (f fakeRoundPrizeReader) AllPrizesForRoundPage(
	ctx context.Context,
	round int64,
	after *cgstore.PrizePageCursor,
	limit int,
) ([]cgmodel.CGPrizeHistory, bool, error) {
	if f.page == nil {
		return []cgmodel.CGPrizeHistory{}, false, nil
	}
	return f.page(ctx, round, after, limit)
}

func (f fakeRoundRaffleReader) CompletedRoundExists(ctx context.Context, round int64) (bool, error) {
	if f.exists == nil {
		return true, nil
	}
	return f.exists(ctx, round)
}

func (f fakeRoundRaffleReader) RaffleEthDepositsByRoundPage(
	ctx context.Context,
	round int64,
	after *cgstore.RaffleEthDepositPageCursor,
	limit int,
) ([]cgstore.RaffleEthDepositRecord, bool, error) {
	if f.eth == nil {
		return []cgstore.RaffleEthDepositRecord{}, false, nil
	}
	return f.eth(ctx, round, after, limit)
}

func (f fakeRoundRaffleReader) RaffleNFTWinnersByRoundPage(
	ctx context.Context,
	round int64,
	isStaker bool,
	after *cgstore.RaffleNFTWinnerPageCursor,
	limit int,
) ([]cgmodel.CGRaffleNFTWinnerRec, bool, error) {
	if f.nft == nil {
		return []cgmodel.CGRaffleNFTWinnerRec{}, false, nil
	}
	return f.nft(ctx, round, isStaker, after, limit)
}

func (f fakeRoundDonationReader) EthDonationsByRoundPage(
	ctx context.Context,
	round int64,
	after *cgstore.DonationPageCursor,
	limit int,
) ([]cgstore.RoundEthDonationRecord, bool, error) {
	if f.eth == nil {
		return []cgstore.RoundEthDonationRecord{}, false, nil
	}
	return f.eth(ctx, round, after, limit)
}

func (f fakeRoundDonationReader) ERC20DonationsByRoundPage(
	ctx context.Context,
	round int64,
	after *cgstore.DonationPageCursor,
	limit int,
) ([]cgstore.RoundERC20DonationRecord, bool, error) {
	if f.erc20 == nil {
		return []cgstore.RoundERC20DonationRecord{}, false, nil
	}
	return f.erc20(ctx, round, after, limit)
}

func (f fakeRoundDonationReader) NFTDonationsByRoundPage(
	ctx context.Context,
	round int64,
	after *cgstore.DonationPageCursor,
	limit int,
) ([]cgstore.RoundNFTDonationRecord, bool, error) {
	if f.nft == nil {
		return []cgstore.RoundNFTDonationRecord{}, false, nil
	}
	return f.nft(ctx, round, after, limit)
}

func (f fakeStatisticsReader) CosmicGameGlobalStatistics(ctx context.Context) (cgstore.GlobalStatisticsRecord, error) {
	if f.global == nil {
		return validGlobalStatisticsRecord(), nil
	}
	return f.global(ctx)
}

func (f fakeStatisticsReader) RecordCounters(ctx context.Context) (cgmodel.CGRecordCounters, error) {
	if f.counters == nil {
		return cgmodel.CGRecordCounters{}, nil
	}
	return f.counters(ctx)
}

func (f fakeStatisticsReader) ROILeaderboardPage(
	ctx context.Context,
	minBids int,
	sort cgstore.ROILeaderboardSort,
	after *cgstore.ROILeaderboardPageCursor,
	limit int,
) ([]cgstore.ROILeaderboardRecord, bool, error) {
	if f.roi == nil {
		return []cgstore.ROILeaderboardRecord{}, false, nil
	}
	return f.roi(ctx, minBids, sort, after, limit)
}

func (f fakeStatisticsReader) ClaimsSummaryPage(
	ctx context.Context,
	after *cgstore.ClaimSummaryCursor,
	limit int,
) ([]cgstore.ClaimSummaryRecord, bool, error) {
	if f.claims == nil {
		return []cgstore.ClaimSummaryRecord{}, false, nil
	}
	return f.claims(ctx, after, limit)
}

func (f fakeStatisticsReader) CompletedRoundExists(ctx context.Context, round int64) (bool, error) {
	if f.exists == nil {
		return true, nil
	}
	return f.exists(ctx, round)
}

func (f fakeStatisticsReader) ClaimSummaryByRound(ctx context.Context, round int64) (cgstore.ClaimSummaryRecord, error) {
	if f.summary == nil {
		return validClaimSummaryRecord(round), nil
	}
	return f.summary(ctx, round)
}

func (f fakeStatisticsReader) ClaimTransactionsPage(
	ctx context.Context,
	round int64,
	after *cgstore.ClaimEventCursor,
	limit int,
) ([]cgstore.ClaimTransactionRecord, bool, error) {
	if f.transactions == nil {
		return []cgstore.ClaimTransactionRecord{}, false, nil
	}
	return f.transactions(ctx, round, after, limit)
}

func (f fakeStatisticsReader) AttachedTokensPage(
	ctx context.Context,
	round int64,
	after *cgstore.ClaimEventCursor,
	limit int,
) ([]cgstore.AttachedTokenRecord, bool, error) {
	if f.attached == nil {
		return []cgstore.AttachedTokenRecord{}, false, nil
	}
	return f.attached(ctx, round, after, limit)
}

func (f fakeStatisticsReader) UnclaimedItemsPage(
	ctx context.Context,
	round int64,
	after *cgstore.UnclaimedItemCursor,
	limit int,
) ([]cgstore.UnclaimedItemRecord, bool, error) {
	if f.unclaimed == nil {
		return []cgstore.UnclaimedItemRecord{}, false, nil
	}
	return f.unclaimed(ctx, round, after, limit)
}

func (f fakeBiddingAnalyticsReader) BidFrequencyByPeriodBounded(
	ctx context.Context,
	from int,
	to int,
	interval int,
) ([]cgmodel.CGBidFrequencyBucket, error) {
	if f.frequency == nil {
		return []cgmodel.CGBidFrequencyBucket{}, nil
	}
	return f.frequency(ctx, from, to, interval)
}

func (f fakeBiddingAnalyticsReader) BidTypeRatioByPeriodBounded(
	ctx context.Context,
	from int,
	to int,
	interval int,
) ([]cgmodel.CGBidTypeRatioBucket, error) {
	if f.ratio == nil {
		return []cgmodel.CGBidTypeRatioBucket{}, nil
	}
	return f.ratio(ctx, from, to, interval)
}

func (f fakeBiddingAnalyticsReader) TopBidderActivePeriodsBounded(
	ctx context.Context,
	top int,
	from int,
	to int,
	gapHours int,
	minBids int,
) ([]cgmodel.CGTopBidderInfo, []cgmodel.CGBidderActivePeriod, bool, error) {
	if f.periods == nil {
		return []cgmodel.CGTopBidderInfo{}, []cgmodel.CGBidderActivePeriod{}, false, nil
	}
	return f.periods(ctx, top, from, to, gapHours, minBids)
}

func (f fakeBiddingAnalyticsReader) BidTimeBounds(ctx context.Context) (int64, int64, error) {
	if f.bounds == nil {
		return 0, 0, nil
	}
	return f.bounds(ctx)
}

func (f fakeContractAddressReader) ContractAddrs(ctx context.Context) (cgmodel.CosmicGameContractAddrs, error) {
	if f.get == nil {
		return cgmodel.CosmicGameContractAddrs{}, nil
	}
	return f.get(ctx)
}

func (f fakeParticipantReader) BidderParticipantsPage(
	ctx context.Context,
	after *cgstore.ParticipantPageCursor,
	limit int,
) ([]cgstore.BidderParticipantRecord, bool, error) {
	if f.bidders == nil {
		return []cgstore.BidderParticipantRecord{}, false, nil
	}
	return f.bidders(ctx, after, limit)
}

func (f fakeParticipantReader) WinnerParticipantsPage(
	ctx context.Context,
	after *cgstore.ParticipantPageCursor,
	limit int,
) ([]cgstore.WinnerParticipantRecord, bool, error) {
	if f.winners == nil {
		return []cgstore.WinnerParticipantRecord{}, false, nil
	}
	return f.winners(ctx, after, limit)
}

func (f fakeParticipantReader) DonorParticipantsPage(
	ctx context.Context,
	after *cgstore.ParticipantPageCursor,
	limit int,
) ([]cgstore.DonorParticipantRecord, bool, error) {
	if f.donors == nil {
		return []cgstore.DonorParticipantRecord{}, false, nil
	}
	return f.donors(ctx, after, limit)
}

func (f fakeParticipantReader) CSTStakerParticipantsPage(
	ctx context.Context,
	after *cgstore.ParticipantPageCursor,
	limit int,
) ([]cgstore.CSTStakerParticipantRecord, bool, error) {
	if f.cstStakers == nil {
		return []cgstore.CSTStakerParticipantRecord{}, false, nil
	}
	return f.cstStakers(ctx, after, limit)
}

func (f fakeParticipantReader) RandomWalkStakerParticipantsPage(
	ctx context.Context,
	after *cgstore.ParticipantPageCursor,
	limit int,
) ([]cgstore.RandomWalkStakerParticipantRecord, bool, error) {
	if f.randomWalkStakers == nil {
		return []cgstore.RandomWalkStakerParticipantRecord{}, false, nil
	}
	return f.randomWalkStakers(ctx, after, limit)
}

func (f fakeParticipantReader) DualStakerParticipantsPage(
	ctx context.Context,
	after *cgstore.ParticipantPageCursor,
	limit int,
) ([]cgstore.DualStakerParticipantRecord, bool, error) {
	if f.dualStakers == nil {
		return []cgstore.DualStakerParticipantRecord{}, false, nil
	}
	return f.dualStakers(ctx, after, limit)
}

func (f fakeUserReader) UserAddressID(ctx context.Context, address string) (int64, error) {
	if f.addressID == nil {
		return 1, nil
	}
	return f.addressID(ctx, address)
}

func (f fakeUserReader) UserProfile(ctx context.Context, userAid int64) (cgstore.UserProfileRecord, error) {
	if f.profile == nil {
		return cgstore.UserProfileRecord{
			Address:               "0x0000000000000000000000000000000000000001",
			TotalETHSpentWei:      "0",
			TotalCSTSpentWei:      "0",
			MaxMainPrizeETHWei:    "0",
			TotalETHWonWei:        "0",
			RaffleETHTotalWei:     "0",
			RaffleCSTTotalWei:     "0",
			ETHDonatedWei:         "0",
			CSTTotalRewardWei:     "0",
			CSTUnclaimedRewardWei: "0",
		}, nil
	}
	return f.profile(ctx, userAid)
}

func (f fakeUserReader) BidsByUserPage(
	ctx context.Context,
	userAid int64,
	after *cgstore.UserBidPageCursor,
	limit int,
) ([]cgmodel.CGBidRec, bool, error) {
	if f.bids == nil {
		return []cgmodel.CGBidRec{}, false, nil
	}
	return f.bids(ctx, userAid, after, limit)
}

func (f fakeUserHistoryReader) UserAddressID(ctx context.Context, address string) (int64, error) {
	if f.addressID == nil {
		return 1, nil
	}
	return f.addressID(ctx, address)
}

func (f fakeUserHistoryReader) UserPrizesPage(
	ctx context.Context,
	userAid int64,
	after *cgstore.UserPrizePageCursor,
	limit int,
) ([]cgmodel.CGPrizeHistory, bool, error) {
	if f.prizes == nil {
		return []cgmodel.CGPrizeHistory{}, false, nil
	}
	return f.prizes(ctx, userAid, after, limit)
}

func (f fakeUserHistoryReader) UserRaffleEthDepositsPage(
	ctx context.Context,
	userAid int64,
	claimed *bool,
	after *cgstore.UserEventPageCursor,
	limit int,
) ([]cgstore.UserRaffleEthDepositRecord, bool, error) {
	if f.deposits == nil {
		return []cgstore.UserRaffleEthDepositRecord{}, false, nil
	}
	return f.deposits(ctx, userAid, claimed, after, limit)
}

func (f fakeUserHistoryReader) UserRaffleNftWinsPage(
	ctx context.Context,
	userAid int64,
	after *cgstore.UserEventPageCursor,
	limit int,
) ([]cgstore.UserRaffleNftWinRecord, bool, error) {
	if f.nftWins == nil {
		return []cgstore.UserRaffleNftWinRecord{}, false, nil
	}
	return f.nftWins(ctx, userAid, after, limit)
}

func (f fakeUserHistoryReader) EthDonationsByUserPage(
	ctx context.Context,
	userAid int64,
	after *cgstore.UserEventPageCursor,
	limit int,
) ([]cgstore.RoundEthDonationRecord, bool, error) {
	if f.ethDonations == nil {
		return []cgstore.RoundEthDonationRecord{}, false, nil
	}
	return f.ethDonations(ctx, userAid, after, limit)
}

func (f fakeUserHistoryReader) ERC20DonationsByUserPage(
	ctx context.Context,
	userAid int64,
	after *cgstore.UserEventPageCursor,
	limit int,
) ([]cgstore.RoundERC20DonationRecord, bool, error) {
	if f.ercDonations == nil {
		return []cgstore.RoundERC20DonationRecord{}, false, nil
	}
	return f.ercDonations(ctx, userAid, after, limit)
}

func (f fakeUserHistoryReader) NFTDonationsByUserPage(
	ctx context.Context,
	userAid int64,
	after *cgstore.UserEventPageCursor,
	limit int,
) ([]cgstore.RoundNFTDonationRecord, bool, error) {
	if f.nftDonations == nil {
		return []cgstore.RoundNFTDonationRecord{}, false, nil
	}
	return f.nftDonations(ctx, userAid, after, limit)
}

func (f fakeUserHistoryReader) UserDonatedNftsPage(
	ctx context.Context,
	userAid int64,
	claimed *bool,
	after *cgstore.UserEventPageCursor,
	limit int,
) ([]cgstore.UserDonatedNftRecord, bool, error) {
	if f.donatedNfts == nil {
		return []cgstore.UserDonatedNftRecord{}, false, nil
	}
	return f.donatedNfts(ctx, userAid, claimed, after, limit)
}

func (f fakeUserHistoryReader) UserDonatedErc20Page(
	ctx context.Context,
	userAid int64,
	after *cgstore.UserDonatedErc20PageCursor,
	limit int,
) ([]cgstore.UserDonatedErc20Record, bool, error) {
	if f.donatedTokens == nil {
		return []cgstore.UserDonatedErc20Record{}, false, nil
	}
	return f.donatedTokens(ctx, userAid, after, limit)
}

func (f fakeUserStakingReader) UserAddressID(ctx context.Context, address string) (int64, error) {
	if f.addressID == nil {
		return 1, nil
	}
	return f.addressID(ctx, address)
}

func (f fakeUserStakingReader) UserCstStakingActionsPage(
	ctx context.Context,
	userAid int64,
	after *cgstore.UserEventPageCursor,
	limit int,
) ([]cgstore.UserStakingActionRecord, bool, error) {
	if f.cstActions == nil {
		return []cgstore.UserStakingActionRecord{}, false, nil
	}
	return f.cstActions(ctx, userAid, after, limit)
}

func (f fakeUserStakingReader) UserRwalkStakingActionsPage(
	ctx context.Context,
	userAid int64,
	after *cgstore.UserEventPageCursor,
	limit int,
) ([]cgstore.UserStakingActionRecord, bool, error) {
	if f.rwalkActions == nil {
		return []cgstore.UserStakingActionRecord{}, false, nil
	}
	return f.rwalkActions(ctx, userAid, after, limit)
}

func (f fakeUserStakingReader) UserStakedCstTokensPage(
	ctx context.Context,
	userAid int64,
	after *cgstore.UserStakingTokenPageCursor,
	limit int,
) ([]cgstore.UserStakedCstTokenRecord, bool, error) {
	if f.cstStaked == nil {
		return []cgstore.UserStakedCstTokenRecord{}, false, nil
	}
	return f.cstStaked(ctx, userAid, after, limit)
}

func (f fakeUserStakingReader) UserStakedRwalkTokensPage(
	ctx context.Context,
	userAid int64,
	after *cgstore.UserStakingTokenPageCursor,
	limit int,
) ([]cgstore.UserStakedRwalkTokenRecord, bool, error) {
	if f.rwalkStaked == nil {
		return []cgstore.UserStakedRwalkTokenRecord{}, false, nil
	}
	return f.rwalkStaked(ctx, userAid, after, limit)
}

func (f fakeUserStakingReader) UserStakingDepositsPage(
	ctx context.Context,
	userAid int64,
	claimed *bool,
	after *cgstore.UserStakingDepositPageCursor,
	limit int,
) ([]cgstore.UserStakingDepositRecord, bool, error) {
	if f.deposits == nil {
		return []cgstore.UserStakingDepositRecord{}, false, nil
	}
	return f.deposits(ctx, userAid, claimed, after, limit)
}

func (f fakeUserStakingReader) StakingDepositExists(ctx context.Context, depositID int64) (bool, error) {
	if f.depositExists == nil {
		return true, nil
	}
	return f.depositExists(ctx, depositID)
}

func (f fakeUserStakingReader) UserStakingDepositRewardsPage(
	ctx context.Context,
	userAid int64,
	depositID int64,
	after *cgstore.UserStakingRewardPageCursor,
	limit int,
) ([]cgstore.UserStakingDepositRewardRecord, bool, error) {
	if f.depositReward == nil {
		return []cgstore.UserStakingDepositRewardRecord{}, false, nil
	}
	return f.depositReward(ctx, userAid, depositID, after, limit)
}

func (f fakeUserStakingReader) UserStakingTokenRewardsPage(
	ctx context.Context,
	userAid int64,
	after *cgstore.UserStakingTokenPageCursor,
	limit int,
) ([]cgstore.UserStakingTokenRewardRecord, bool, error) {
	if f.tokenRewards == nil {
		return []cgstore.UserStakingTokenRewardRecord{}, false, nil
	}
	return f.tokenRewards(ctx, userAid, after, limit)
}

func (f fakeUserStakingReader) CosmicSignatureTokenExists(ctx context.Context, tokenID int64) (bool, error) {
	if f.tokenExists == nil {
		return true, nil
	}
	return f.tokenExists(ctx, tokenID)
}

func (f fakeUserStakingReader) UserStakingTokenRewardDepositsPage(
	ctx context.Context,
	userAid int64,
	tokenID int64,
	after *cgstore.UserStakingTokenDepositPageCursor,
	limit int,
) ([]cgstore.UserStakingTokenRewardDepositRecord, bool, error) {
	if f.tokenDeposits == nil {
		return []cgstore.UserStakingTokenRewardDepositRecord{}, false, nil
	}
	return f.tokenDeposits(ctx, userAid, tokenID, after, limit)
}

func (f fakeUserActivityReader) UserAddressID(ctx context.Context, address string) (int64, error) {
	if f.addressID == nil {
		return 1, nil
	}
	return f.addressID(ctx, address)
}

func (f fakeUserActivityReader) UserCosmicSignatureTokensPage(
	ctx context.Context,
	userAid int64,
	after *cgstore.UserTokenPageCursor,
	limit int,
) ([]cgstore.UserOwnedTokenRecord, bool, error) {
	if f.ownedTokens == nil {
		return []cgstore.UserOwnedTokenRecord{}, false, nil
	}
	return f.ownedTokens(ctx, userAid, after, limit)
}

func (f fakeUserActivityReader) UserCosmicSignatureTransfersPage(
	ctx context.Context,
	userAid int64,
	after *cgstore.UserEventPageCursor,
	limit int,
) ([]cgstore.UserCosmicSignatureTransferRecord, bool, error) {
	if f.csTransfers == nil {
		return []cgstore.UserCosmicSignatureTransferRecord{}, false, nil
	}
	return f.csTransfers(ctx, userAid, after, limit)
}

func (f fakeUserActivityReader) UserCosmicTokenTransfersPage(
	ctx context.Context,
	userAid int64,
	after *cgstore.UserEventPageCursor,
	limit int,
) ([]cgstore.UserCosmicTokenTransferRecord, bool, error) {
	if f.ctTransfers == nil {
		return []cgstore.UserCosmicTokenTransferRecord{}, false, nil
	}
	return f.ctTransfers(ctx, userAid, after, limit)
}

func (f fakeUserActivityReader) UserMarketingRewardsPage(
	ctx context.Context,
	userAid int64,
	after *cgstore.UserEventPageCursor,
	limit int,
) ([]cgstore.UserMarketingRewardRecord, bool, error) {
	if f.mktRewards == nil {
		return []cgstore.UserMarketingRewardRecord{}, false, nil
	}
	return f.mktRewards(ctx, userAid, after, limit)
}

func (f fakeUserActivityReader) UserCosmicTokenSummaryV2(
	ctx context.Context,
	userAid int64,
) (cgstore.UserCosmicTokenSummaryRecord, error) {
	if f.tokenSummary == nil {
		return zeroUserCosmicTokenSummaryRecord(), nil
	}
	return f.tokenSummary(ctx, userAid)
}

func (f fakeUserActivityReader) UserPendingWinnings(
	ctx context.Context,
	userAid int64,
) (cgstore.UserPendingWinningsRecord, error) {
	if f.pending == nil {
		return cgstore.UserPendingWinningsRecord{
			RaffleEthWei:        "0",
			ChronoWarriorEthWei: "0",
			StakingRewardWei:    "0",
		}, nil
	}
	return f.pending(ctx, userAid)
}

// zeroUserCosmicTokenSummaryRecord is the store-level zero shape the SQL
// summary produces for an indexed wallet without activity.
func zeroUserCosmicTokenSummaryRecord() cgstore.UserCosmicTokenSummaryRecord {
	return cgstore.UserCosmicTokenSummaryRecord{
		BalanceWei:                 "0",
		BiddingRewardsWei:          "0",
		MainPrizesWei:              "0",
		RafflePrizesWei:            "0",
		ChronoWarriorPrizesWei:     "0",
		EnduranceChampionPrizesWei: "0",
		LastCstBidderPrizesWei:     "0",
		MarketingRewardsWei:        "0",
		TotalEarnedWei:             "0",
		ConsumedInBidsWei:          "0",
		NetWei:                     "0",
	}
}

func (f fakeContractState) Snapshot() contractstate.Snapshot {
	if f.snapshot == nil {
		return contractstate.Snapshot{}
	}
	return f.snapshot()
}

func TestListRoundBidsPaginatesWithOpaqueCursor(t *testing.T) {
	t.Parallel()

	var gotRound int64
	var gotAfter cgstore.BidPageCursor
	var gotLimit int
	first := validBidRecord()
	first.RoundNum, first.BidPosition, first.Tx.EvtLogId = 9, 1, 100
	second := validBidRecord()
	second.RoundNum, second.BidPosition, second.Tx.EvtLogId = 9, 2, 101

	server := newTestServer(t, fakeBidReader{
		page: func(_ context.Context, round int64, after cgstore.BidPageCursor, limit int) ([]cgmodel.CGBidRec, bool, error) {
			gotRound, gotAfter, gotLimit = round, after, limit
			return []cgmodel.CGBidRec{first, second}, true, nil
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/rounds/9/bids?limit=2")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if gotRound != 9 || gotAfter != (cgstore.BidPageCursor{}) || gotLimit != 2 {
		t.Fatalf("repository args = (%d,%+v,%d)", gotRound, gotAfter, gotLimit)
	}

	var page RoundBidPage
	decodeResponse(t, response, &page)
	if len(page.Data) != 2 || page.Meta.Limit != 2 || page.Meta.NextCursor == nil {
		t.Fatalf("page = %+v", page)
	}
	cursor, err := decodeBidCursor(*page.Meta.NextCursor, 9)
	if err != nil {
		t.Fatalf("decode next cursor: %v", err)
	}
	if cursor.BidPosition != 2 || cursor.EventLogID != 101 {
		t.Fatalf("next cursor = %+v", cursor)
	}
}

func TestListRoundBidsDecodesContinuationCursor(t *testing.T) {
	t.Parallel()

	encoded, err := encodeBidCursor(bidCursor{
		Version:     bidCursorVersion,
		Round:       3,
		BidPosition: 7,
		EventLogID:  88,
	})
	if err != nil {
		t.Fatal(err)
	}

	var gotAfter cgstore.BidPageCursor
	server := newTestServer(t, fakeBidReader{
		page: func(_ context.Context, _ int64, after cgstore.BidPageCursor, limit int) ([]cgmodel.CGBidRec, bool, error) {
			gotAfter = after
			if limit != defaultPageLimit {
				t.Errorf("default limit = %d, want %d", limit, defaultPageLimit)
			}
			return []cgmodel.CGBidRec{}, false, nil
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/rounds/3/bids?cursor="+encoded)
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	if gotAfter != (cgstore.BidPageCursor{BidPosition: 7, EventLogID: 88}) {
		t.Fatalf("decoded repository cursor = %+v", gotAfter)
	}
	if !bytes.Contains(response.Body.Bytes(), []byte(`"data":[]`)) {
		t.Fatalf("empty data was not encoded as []: %s", response.Body.String())
	}
	var page RoundBidPage
	decodeResponse(t, response, &page)
	if page.Meta.NextCursor != nil {
		t.Fatalf("exhausted page has next cursor %q", *page.Meta.NextCursor)
	}
}

func TestListRoundBidsRejectsInvalidInput(t *testing.T) {
	t.Parallel()

	crossRound, err := encodeBidCursor(bidCursor{
		Version:     bidCursorVersion,
		Round:       2,
		BidPosition: 1,
		EventLogID:  1,
	})
	if err != nil {
		t.Fatal(err)
	}

	tests := map[string]string{
		"negative round":   "/api/v2/cosmicgame/rounds/-1/bids",
		"zero limit":       "/api/v2/cosmicgame/rounds/1/bids?limit=0",
		"excessive limit":  "/api/v2/cosmicgame/rounds/1/bids?limit=201",
		"duplicate limit":  "/api/v2/cosmicgame/rounds/1/bids?limit=1&limit=2",
		"malformed cursor": "/api/v2/cosmicgame/rounds/1/bids?cursor=not-a-cursor",
		"oversized cursor": "/api/v2/cosmicgame/rounds/1/bids?cursor=" + strings.Repeat("a", maxCursorLength+1),
		"cross-round":      "/api/v2/cosmicgame/rounds/1/bids?cursor=" + crossRound,
		"bind round":       "/api/v2/cosmicgame/rounds/not-a-number/bids",
		"bind limit":       "/api/v2/cosmicgame/rounds/1/bids?limit=wat",
	}
	for name, path := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newTestServer(t, fakeBidReader{}), path)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}
}

func TestListRoundBidsHidesRepositoryErrors(t *testing.T) {
	t.Parallel()

	server := newTestServer(t, fakeBidReader{
		page: func(context.Context, int64, cgstore.BidPageCursor, int) ([]cgmodel.CGBidRec, bool, error) {
			return nil, false, errors.New("password=super-secret")
		},
	})
	response := serve(t, server, "/api/v2/cosmicgame/rounds/1/bids")
	assertProblem(t, response, http.StatusInternalServerError)
	if strings.Contains(response.Body.String(), "super-secret") {
		t.Fatalf("internal error leaked: %s", response.Body.String())
	}
}

func TestListRoundBidsRejectsInconsistentRepositoryPage(t *testing.T) {
	t.Parallel()

	t.Run("has more without row", func(t *testing.T) {
		t.Parallel()
		server := newTestServer(t, fakeBidReader{
			page: func(context.Context, int64, cgstore.BidPageCursor, int) ([]cgmodel.CGBidRec, bool, error) {
				return []cgmodel.CGBidRec{}, true, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/rounds/1/bids")
		assertProblem(t, response, http.StatusInternalServerError)
	})

	t.Run("out of order", func(t *testing.T) {
		t.Parallel()
		first := validBidRecord()
		first.RoundNum, first.BidPosition, first.Tx.EvtLogId = 1, 2, 20
		second := validBidRecord()
		second.RoundNum, second.BidPosition, second.Tx.EvtLogId = 1, 1, 10
		server := newTestServer(t, fakeBidReader{
			page: func(context.Context, int64, cgstore.BidPageCursor, int) ([]cgmodel.CGBidRec, bool, error) {
				return []cgmodel.CGBidRec{first, second}, false, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/rounds/1/bids")
		assertProblem(t, response, http.StatusInternalServerError)
	})
}

func TestGetRoundBidResponses(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		record := validBidRecord()
		record.RoundNum, record.BidPosition = 4, 2
		server := newTestServer(t, fakeBidReader{
			item: func(_ context.Context, round, position int64) (cgmodel.CGBidRec, error) {
				if round != 4 || position != 2 {
					t.Fatalf("repository args = (%d,%d)", round, position)
				}
				return record, nil
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/rounds/4/bids/2")
		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
		}
		var bid Bid
		decodeResponse(t, response, &bid)
		if bid.Round != 4 || bid.Position != 2 {
			t.Fatalf("bid = %+v", bid)
		}
	})

	t.Run("not found", func(t *testing.T) {
		t.Parallel()
		response := serve(t, newTestServer(t, fakeBidReader{}), "/api/v2/cosmicgame/rounds/4/bids/99")
		assertProblem(t, response, http.StatusNotFound)
	})

	t.Run("invalid position", func(t *testing.T) {
		t.Parallel()
		response := serve(t, newTestServer(t, fakeBidReader{}), "/api/v2/cosmicgame/rounds/4/bids/0")
		assertProblem(t, response, http.StatusBadRequest)
	})

	t.Run("repository failure", func(t *testing.T) {
		t.Parallel()
		server := newTestServer(t, fakeBidReader{
			item: func(context.Context, int64, int64) (cgmodel.CGBidRec, error) {
				return cgmodel.CGBidRec{}, errors.New("private database detail")
			},
		})
		response := serve(t, server, "/api/v2/cosmicgame/rounds/4/bids/2")
		assertProblem(t, response, http.StatusInternalServerError)
		if strings.Contains(response.Body.String(), "private database detail") {
			t.Fatalf("internal error leaked: %s", response.Body.String())
		}
	})
}

func TestNewServerValidatesDependencies(t *testing.T) {
	t.Parallel()

	if _, err := NewServer(nil, nil, nil); err == nil {
		t.Fatal("NewServer accepted nil dependencies")
	}
	if _, err := NewServer(&store.Store{}, nil, nil); err == nil {
		t.Fatal("NewServer accepted a nil contract state")
	}
	if _, err := NewServer(
		&store.Store{},
		&contractstate.State{},
		nil,
		WithClock(nil),
	); err == nil {
		t.Fatal("NewServer accepted a nil clock")
	}
	fixedNow := time.Unix(123, 0)
	configured, err := NewServer(
		&store.Store{},
		&contractstate.State{},
		nil,
		WithClock(func() time.Time { return fixedNow }),
	)
	if err != nil || !configured.now().Equal(fixedNow) {
		t.Fatalf("NewServer clock option: server=%v err=%v", configured, err)
	}
	if _, err := newServer(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil bid repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil round repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil current-round repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, fakeCurrentRoundReader{}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil round-prize repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, fakeCurrentRoundReader{}, fakeRoundPrizeReader{}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil round-raffle repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, fakeCurrentRoundReader{}, fakeRoundPrizeReader{}, fakeRoundRaffleReader{}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil round-donation repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, fakeCurrentRoundReader{}, fakeRoundPrizeReader{}, fakeRoundRaffleReader{}, fakeRoundDonationReader{}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil statistics repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, fakeCurrentRoundReader{}, fakeRoundPrizeReader{}, fakeRoundRaffleReader{}, fakeRoundDonationReader{}, fakeStatisticsReader{}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil bidding analytics repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, fakeCurrentRoundReader{}, fakeRoundPrizeReader{}, fakeRoundRaffleReader{}, fakeRoundDonationReader{}, fakeStatisticsReader{}, fakeBiddingAnalyticsReader{}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil contract-address repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, fakeCurrentRoundReader{}, fakeRoundPrizeReader{}, fakeRoundRaffleReader{}, fakeRoundDonationReader{}, fakeStatisticsReader{}, fakeBiddingAnalyticsReader{}, fakeContractAddressReader{}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil participant repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, fakeCurrentRoundReader{}, fakeRoundPrizeReader{}, fakeRoundRaffleReader{}, fakeRoundDonationReader{}, fakeStatisticsReader{}, fakeBiddingAnalyticsReader{}, fakeContractAddressReader{}, fakeParticipantReader{}, nil, nil, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil user repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, fakeCurrentRoundReader{}, fakeRoundPrizeReader{}, fakeRoundRaffleReader{}, fakeRoundDonationReader{}, fakeStatisticsReader{}, fakeBiddingAnalyticsReader{}, fakeContractAddressReader{}, fakeParticipantReader{}, fakeUserReader{}, nil, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil user-history repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, fakeCurrentRoundReader{}, fakeRoundPrizeReader{}, fakeRoundRaffleReader{}, fakeRoundDonationReader{}, fakeStatisticsReader{}, fakeBiddingAnalyticsReader{}, fakeContractAddressReader{}, fakeParticipantReader{}, fakeUserReader{}, fakeUserHistoryReader{}, nil, nil, nil, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil user-staking repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, fakeCurrentRoundReader{}, fakeRoundPrizeReader{}, fakeRoundRaffleReader{}, fakeRoundDonationReader{}, fakeStatisticsReader{}, fakeBiddingAnalyticsReader{}, fakeContractAddressReader{}, fakeParticipantReader{}, fakeUserReader{}, fakeUserHistoryReader{}, fakeUserStakingReader{}, nil, nil, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil user-activity repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, fakeCurrentRoundReader{}, fakeRoundPrizeReader{}, fakeRoundRaffleReader{}, fakeRoundDonationReader{}, fakeStatisticsReader{}, fakeBiddingAnalyticsReader{}, fakeContractAddressReader{}, fakeParticipantReader{}, fakeUserReader{}, fakeUserHistoryReader{}, fakeUserStakingReader{}, fakeUserActivityReader{}, nil, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil global-directory repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, fakeCurrentRoundReader{}, fakeRoundPrizeReader{}, fakeRoundRaffleReader{}, fakeRoundDonationReader{}, fakeStatisticsReader{}, fakeBiddingAnalyticsReader{}, fakeContractAddressReader{}, fakeParticipantReader{}, fakeUserReader{}, fakeUserHistoryReader{}, fakeUserStakingReader{}, fakeUserActivityReader{}, fakeGlobalDirectoryReader{}, nil, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil global-staking repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, fakeCurrentRoundReader{}, fakeRoundPrizeReader{}, fakeRoundRaffleReader{}, fakeRoundDonationReader{}, fakeStatisticsReader{}, fakeBiddingAnalyticsReader{}, fakeContractAddressReader{}, fakeParticipantReader{}, fakeUserReader{}, fakeUserHistoryReader{}, fakeUserStakingReader{}, fakeUserActivityReader{}, fakeGlobalDirectoryReader{}, fakeGlobalStakingReader{}, nil, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil random-walk repository")
	}
	if _, err := newServer(nil, fakeBidReader{}, fakeRoundReader{}, fakeCurrentRoundReader{}, fakeRoundPrizeReader{}, fakeRoundRaffleReader{}, fakeRoundDonationReader{}, fakeStatisticsReader{}, fakeBiddingAnalyticsReader{}, fakeContractAddressReader{}, fakeParticipantReader{}, fakeUserReader{}, fakeUserHistoryReader{}, fakeUserStakingReader{}, fakeUserActivityReader{}, fakeGlobalDirectoryReader{}, fakeGlobalStakingReader{}, fakeRandomWalkReader{}, nil, nil); err == nil {
		t.Fatal("newServer accepted a nil contract state")
	}
	server, err := newServer(
		nil,
		fakeBidReader{},
		fakeRoundReader{},
		fakeCurrentRoundReader{},
		fakeRoundPrizeReader{},
		fakeRoundRaffleReader{},
		fakeRoundDonationReader{},
		fakeStatisticsReader{},
		fakeBiddingAnalyticsReader{},
		fakeContractAddressReader{},
		fakeParticipantReader{},
		fakeUserReader{},
		fakeUserHistoryReader{},
		fakeUserStakingReader{},
		fakeUserActivityReader{},
		fakeGlobalDirectoryReader{},
		fakeGlobalStakingReader{},
		fakeRandomWalkReader{},
		fakeContractState{},
		nil,
	)
	if err != nil {
		t.Fatalf("newServer rejected test dependencies: %v", err)
	}
	if server.logger == nil {
		t.Fatal("newServer did not install a default logger")
	}
	if server.now == nil {
		t.Fatal("newServer did not install a clock")
	}
}

func newTestServer(t *testing.T, bids bidReader) *Server {
	t.Helper()
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	server, err := newServer(
		nil,
		bids,
		fakeRoundReader{},
		fakeCurrentRoundReader{},
		fakeRoundPrizeReader{},
		fakeRoundRaffleReader{},
		fakeRoundDonationReader{},
		fakeStatisticsReader{},
		fakeBiddingAnalyticsReader{},
		fakeContractAddressReader{},
		fakeParticipantReader{},
		fakeUserReader{},
		fakeUserHistoryReader{},
		fakeUserStakingReader{},
		fakeUserActivityReader{},
		fakeGlobalDirectoryReader{},
		fakeGlobalStakingReader{},
		fakeRandomWalkReader{},
		fakeContractState{},
		logger,
	)
	if err != nil {
		t.Fatalf("newServer: %v", err)
	}
	return server
}

func serve(t *testing.T, server *Server, target string) *httptest.ResponseRecorder {
	t.Helper()
	router := httpx.NewRouter()
	server.RegisterRoutes(router)
	request := httptest.NewRequest(http.MethodGet, target, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	return response
}

func decodeResponse(t *testing.T, response *httptest.ResponseRecorder, target any) {
	t.Helper()
	if err := json.Unmarshal(response.Body.Bytes(), target); err != nil {
		t.Fatalf("decode response: %v\n%s", err, response.Body.String())
	}
}

func assertProblem(t *testing.T, response *httptest.ResponseRecorder, status int) {
	t.Helper()
	if response.Code != status {
		t.Fatalf("status = %d, want %d; body=%s", response.Code, status, response.Body.String())
	}
	if got := response.Header().Get("Content-Type"); got != "application/problem+json" {
		t.Fatalf("Content-Type = %q", got)
	}
	var problem Problem
	decodeResponse(t, response, &problem)
	if problem.Status != status || problem.Type == "" || problem.Title == "" {
		t.Fatalf("problem = %+v", problem)
	}
}
