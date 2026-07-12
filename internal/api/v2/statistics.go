package v2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// GetCosmicGameStatistics implements GET /api/v2/cosmicgame/statistics.
func (s *Server) GetCosmicGameStatistics(
	ctx context.Context,
	_ GetCosmicGameStatisticsRequestObject,
) (GetCosmicGameStatisticsResponseObject, error) {
	const instance = "/api/v2/cosmicgame/statistics"
	record, err := s.statistics.CosmicGameGlobalStatistics(ctx)
	if err != nil {
		s.logInternal(ctx, "get global statistics", err)
		return getCosmicGameStatisticsInternal(internalProblem(instance)), nil
	}
	result, err := mapGlobalStatistics(record)
	if err != nil {
		s.logInternal(ctx, "map global statistics", err)
		return getCosmicGameStatisticsInternal(internalProblem(instance)), nil
	}
	return GetCosmicGameStatistics200JSONResponse{
		CosmicGameGlobalStatisticsJSONResponse: CosmicGameGlobalStatisticsJSONResponse(result),
	}, nil
}

// GetCosmicGameCounters implements GET /api/v2/cosmicgame/statistics/counters.
func (s *Server) GetCosmicGameCounters(
	ctx context.Context,
	_ GetCosmicGameCountersRequestObject,
) (GetCosmicGameCountersResponseObject, error) {
	const instance = "/api/v2/cosmicgame/statistics/counters"
	record, err := s.statistics.RecordCounters(ctx)
	if err != nil {
		s.logInternal(ctx, "get game counters", err)
		return getCosmicGameCountersInternal(internalProblem(instance)), nil
	}
	result, err := mapCounters(record)
	if err != nil {
		s.logInternal(ctx, "map game counters", err)
		return getCosmicGameCountersInternal(internalProblem(instance)), nil
	}
	return GetCosmicGameCounters200JSONResponse{
		CosmicGameCountersJSONResponse: CosmicGameCountersJSONResponse(result),
	}, nil
}

// ListCosmicGameRoiLeaderboard implements
// GET /api/v2/cosmicgame/statistics/leaderboard/roi.
func (s *Server) ListCosmicGameRoiLeaderboard(
	ctx context.Context,
	request ListCosmicGameRoiLeaderboardRequestObject,
) (ListCosmicGameRoiLeaderboardResponseObject, error) {
	const instance = "/api/v2/cosmicgame/statistics/leaderboard/roi"
	sortBy := NetProfit
	if request.Params.Sort != nil {
		sortBy = *request.Params.Sort
	}
	if !sortBy.Valid() {
		return listROILeaderboardBadRequest(newProblem(
			http.StatusBadRequest, "invalid-parameter", "Invalid parameter",
			"Sort must be one of netProfit, roi, winRate, spent, nfts, or bids.",
			instance,
		)), nil
	}
	minBids := 5
	if request.Params.MinBids != nil {
		minBids = *request.Params.MinBids
	}
	if minBids < 0 {
		return listROILeaderboardBadRequest(newProblem(
			http.StatusBadRequest, "invalid-parameter", "Invalid parameter",
			"MinBids must be zero or greater.", instance,
		)), nil
	}
	limit, validLimit := resolvePageLimit(request.Params.Limit)
	if !validLimit {
		return listROILeaderboardBadRequest(newProblem(
			http.StatusBadRequest, "invalid-parameter", "Invalid parameter",
			pageLimitProblemDetail(), instance,
		)), nil
	}
	storeSort := roiStoreSort(sortBy)
	var after *cgstore.ROILeaderboardPageCursor
	if request.Params.Cursor != nil {
		cursor, err := decodeROILeaderboardCursor(*request.Params.Cursor, sortBy, minBids)
		if err != nil {
			return listROILeaderboardBadRequest(newProblem(
				http.StatusBadRequest, "invalid-cursor", "Invalid cursor",
				"The cursor is malformed, unsupported, or belongs to another ROI sort or minBids filter.",
				instance,
			)), nil
		}
		after = &cgstore.ROILeaderboardPageCursor{
			Sort:      storeSort,
			MinBids:   minBids,
			SortValue: cursor.SortValue,
			Secondary: cursor.Secondary,
			BidderAid: cursor.BidderAid,
		}
	}
	records, hasMore, err := s.statistics.ROILeaderboardPage(ctx, minBids, storeSort, after, limit)
	if err != nil {
		s.logInternal(ctx, "list ROI leaderboard", err,
			"sort", string(sortBy), "min_bids", minBids)
		return listROILeaderboardInternal(internalProblem(instance)), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate ROI leaderboard page cardinality", err,
			"sort", string(sortBy), "min_bids", minBids)
		return listROILeaderboardInternal(internalProblem(instance)), nil
	}
	data := make([]RoiLeaderboardEntry, 0, len(records))
	previous := after
	for i := range records {
		if records[i].BidderAid < 1 ||
			(previous != nil && !roiRecordFollows(records[i], *previous, storeSort)) {
			err := errors.New("repository returned an unordered ROI leaderboard page")
			s.logInternal(ctx, "validate ROI leaderboard page", err,
				"sort", string(sortBy), "bidder_aid", records[i].BidderAid)
			return listROILeaderboardInternal(internalProblem(instance)), nil
		}
		mapped, err := mapROILeaderboardEntry(records[i])
		if err != nil {
			s.logInternal(ctx, "map ROI leaderboard entry", err,
				"sort", string(sortBy), "bidder_aid", records[i].BidderAid)
			return listROILeaderboardInternal(internalProblem(instance)), nil
		}
		data = append(data, mapped)
		previous = &cgstore.ROILeaderboardPageCursor{
			Sort:      storeSort,
			MinBids:   minBids,
			SortValue: cgstore.ROILeaderboardSortValue(records[i], storeSort),
			Secondary: roiSecondaryValue(records[i], storeSort),
			BidderAid: records[i].BidderAid,
		}
	}
	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 || previous == nil {
			s.logInternal(ctx, "list ROI leaderboard",
				errors.New("repository reported another ROI page without a cursor row"))
			return listROILeaderboardInternal(internalProblem(instance)), nil
		}
		next, err := encodeROILeaderboardCursor(roiLeaderboardCursor{
			Version:   roiLeaderboardCursorVersion,
			Sort:      sortBy,
			MinBids:   minBids,
			SortValue: previous.SortValue,
			Secondary: previous.Secondary,
			BidderAid: previous.BidderAid,
		})
		if err != nil {
			s.logInternal(ctx, "encode ROI leaderboard cursor", err)
			return listROILeaderboardInternal(internalProblem(instance)), nil
		}
		meta.NextCursor = &next
	}
	return ListCosmicGameRoiLeaderboard200JSONResponse{
		RoiLeaderboardPageJSONResponse: RoiLeaderboardPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameClaims implements GET /api/v2/cosmicgame/statistics/claims.
func (s *Server) ListCosmicGameClaims(
	ctx context.Context,
	request ListCosmicGameClaimsRequestObject,
) (ListCosmicGameClaimsResponseObject, error) {
	const instance = "/api/v2/cosmicgame/statistics/claims"
	limit, validLimit := resolvePageLimit(request.Params.Limit)
	if !validLimit {
		return listClaimsBadRequest(newProblem(
			http.StatusBadRequest, "invalid-parameter", "Invalid parameter",
			pageLimitProblemDetail(), instance,
		)), nil
	}
	var after *cgstore.ClaimSummaryCursor
	if request.Params.Cursor != nil {
		cursor, err := decodeClaimSummaryCursor(*request.Params.Cursor)
		if err != nil {
			return listClaimsBadRequest(newProblem(
				http.StatusBadRequest, "invalid-cursor", "Invalid cursor",
				"The cursor is malformed or unsupported.", instance,
			)), nil
		}
		after = &cgstore.ClaimSummaryCursor{
			RoundNum:   cursor.Round,
			EventLogID: cursor.EventLogID,
		}
	}
	records, hasMore, err := s.statistics.ClaimsSummaryPage(ctx, after, limit)
	if err != nil {
		s.logInternal(ctx, "list claim summaries", err)
		return listClaimsInternal(internalProblem(instance)), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate claim summary page cardinality", err)
		return listClaimsInternal(internalProblem(instance)), nil
	}
	data := make([]ClaimSummary, 0, len(records))
	previous := after
	now := time.Now()
	for i := range records {
		current := cgstore.ClaimSummaryCursor{
			RoundNum:   records[i].RoundNum,
			EventLogID: records[i].EventLogID,
		}
		if previous != nil && !claimSummaryFollows(current, *previous) {
			s.logInternal(ctx, "validate claim summary page",
				errors.New("repository returned an unordered claim summary page"),
				"round", current.RoundNum, "event_log_id", current.EventLogID)
			return listClaimsInternal(internalProblem(instance)), nil
		}
		mapped, err := mapClaimSummary(records[i], now)
		if err != nil {
			s.logInternal(ctx, "map claim summary", err, "round", records[i].RoundNum)
			return listClaimsInternal(internalProblem(instance)), nil
		}
		data = append(data, mapped)
		previous = &current
	}
	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 || previous == nil {
			s.logInternal(ctx, "list claim summaries",
				errors.New("repository reported another claim page without a cursor row"))
			return listClaimsInternal(internalProblem(instance)), nil
		}
		next, err := encodeClaimSummaryCursor(claimSummaryCursor{
			Version:    claimSummaryCursorVersion,
			Round:      previous.RoundNum,
			EventLogID: previous.EventLogID,
		})
		if err != nil {
			s.logInternal(ctx, "encode claim summary cursor", err)
			return listClaimsInternal(internalProblem(instance)), nil
		}
		meta.NextCursor = &next
	}
	return ListCosmicGameClaims200JSONResponse{
		ClaimSummaryPageJSONResponse: ClaimSummaryPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// GetRoundClaims implements GET /api/v2/cosmicgame/rounds/{round}/claims.
func (s *Server) GetRoundClaims(
	ctx context.Context,
	request GetRoundClaimsRequestObject,
) (GetRoundClaimsResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/rounds/%d/claims", request.Round)
	if request.Round < 0 {
		return getRoundClaimsBadRequest(newProblem(
			http.StatusBadRequest, "invalid-parameter", "Invalid parameter",
			"Round must be zero or greater.", instance,
		)), nil
	}
	limit, validLimit := resolvePageLimit(request.Params.Limit)
	if !validLimit {
		return getRoundClaimsBadRequest(newProblem(
			http.StatusBadRequest, "invalid-parameter", "Invalid parameter",
			pageLimitProblemDetail(), instance,
		)), nil
	}
	transactionAfter, problem := decodeClaimEventRequestCursor(
		request.Params.ClaimTransactionsCursor, request.Round,
		claimDetailTransactions, instance,
	)
	if problem != nil {
		return getRoundClaimsBadRequest(*problem), nil
	}
	attachedAfter, problem := decodeClaimEventRequestCursor(
		request.Params.AttachedTokensCursor, request.Round,
		claimDetailAttached, instance,
	)
	if problem != nil {
		return getRoundClaimsBadRequest(*problem), nil
	}
	unclaimedAfter, problem := decodeUnclaimedRequestCursor(
		request.Params.UnclaimedItemsCursor, request.Round, instance,
	)
	if problem != nil {
		return getRoundClaimsBadRequest(*problem), nil
	}
	exists, err := s.statistics.CompletedRoundExists(ctx, request.Round)
	if err != nil {
		s.logInternal(ctx, "check completed round for claims", err, "round", request.Round)
		return getRoundClaimsInternal(internalProblem(instance)), nil
	}
	if !exists {
		return getRoundClaimsNotFound(roundNotFoundProblem(instance)), nil
	}
	summaryRecord, err := s.statistics.ClaimSummaryByRound(ctx, request.Round)
	if err != nil {
		s.logInternal(ctx, "get round claim summary", err, "round", request.Round)
		return getRoundClaimsInternal(internalProblem(instance)), nil
	}
	summary, err := mapClaimSummary(summaryRecord, time.Now())
	if err != nil || summary.Round != request.Round {
		if err == nil {
			err = errors.New("repository returned another round's claim summary")
		}
		s.logInternal(ctx, "map round claim summary", err, "round", request.Round)
		return getRoundClaimsInternal(internalProblem(instance)), nil
	}
	transactionRecords, transactionMore, err := s.statistics.ClaimTransactionsPage(
		ctx, request.Round, transactionAfter, limit,
	)
	if err != nil {
		s.logInternal(ctx, "list round claim transactions", err, "round", request.Round)
		return getRoundClaimsInternal(internalProblem(instance)), nil
	}
	transactions, err := buildClaimTransactionPage(
		transactionRecords, transactionMore, request.Round, transactionAfter, limit,
	)
	if err != nil {
		s.logInternal(ctx, "build round claim transaction page", err, "round", request.Round)
		return getRoundClaimsInternal(internalProblem(instance)), nil
	}
	attachedRecords, attachedMore, err := s.statistics.AttachedTokensPage(
		ctx, request.Round, attachedAfter, limit,
	)
	if err != nil {
		s.logInternal(ctx, "list round attached tokens", err, "round", request.Round)
		return getRoundClaimsInternal(internalProblem(instance)), nil
	}
	attached, err := buildAttachedTokenPage(
		attachedRecords, attachedMore, request.Round, attachedAfter, limit,
	)
	if err != nil {
		s.logInternal(ctx, "build round attached-token page", err, "round", request.Round)
		return getRoundClaimsInternal(internalProblem(instance)), nil
	}
	unclaimedRecords, unclaimedMore, err := s.statistics.UnclaimedItemsPage(
		ctx, request.Round, unclaimedAfter, limit,
	)
	if err != nil {
		s.logInternal(ctx, "list round unclaimed items", err, "round", request.Round)
		return getRoundClaimsInternal(internalProblem(instance)), nil
	}
	unclaimed, err := buildUnclaimedItemPage(
		unclaimedRecords, unclaimedMore, request.Round, unclaimedAfter, limit,
	)
	if err != nil {
		s.logInternal(ctx, "build round unclaimed-item page", err, "round", request.Round)
		return getRoundClaimsInternal(internalProblem(instance)), nil
	}
	return GetRoundClaims200JSONResponse{
		RoundClaimsDetailJSONResponse: RoundClaimsDetailJSONResponse{
			AttachedTokens:    attached,
			ClaimTransactions: transactions,
			Summary:           summary,
			UnclaimedItems:    unclaimed,
		},
	}, nil
}

func buildClaimTransactionPage(
	records []cgstore.ClaimTransactionRecord,
	hasMore bool,
	round int64,
	after *cgstore.ClaimEventCursor,
	limit int,
) (ClaimTransactionPage, error) {
	if err := validatePageCardinality(len(records), limit); err != nil {
		return ClaimTransactionPage{}, err
	}
	data := make([]AssetClaimTransaction, 0, len(records))
	previous := int64(0)
	if after != nil {
		previous = after.EventLogID
	}
	for i := range records {
		if records[i].RoundNum != round || records[i].EventLogID <= previous {
			return ClaimTransactionPage{}, errors.New("unordered or out-of-scope claim transaction page")
		}
		mapped, err := mapAssetClaimTransaction(records[i])
		if err != nil {
			return ClaimTransactionPage{}, err
		}
		data = append(data, mapped)
		previous = records[i].EventLogID
	}
	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 {
			return ClaimTransactionPage{}, errors.New("claim transaction page hasMore without row")
		}
		next, err := encodeClaimDetailCursor(claimDetailCursor{
			Version:    claimDetailCursorVersion,
			Round:      round,
			Section:    claimDetailTransactions,
			EventLogID: previous,
		})
		if err != nil {
			return ClaimTransactionPage{}, err
		}
		meta.NextCursor = &next
	}
	return ClaimTransactionPage{Data: data, Meta: meta}, nil
}

func buildAttachedTokenPage(
	records []cgstore.AttachedTokenRecord,
	hasMore bool,
	round int64,
	after *cgstore.ClaimEventCursor,
	limit int,
) (AttachedTokenPage, error) {
	if err := validatePageCardinality(len(records), limit); err != nil {
		return AttachedTokenPage{}, err
	}
	data := make([]AttachedToken, 0, len(records))
	previous := int64(0)
	if after != nil {
		previous = after.EventLogID
	}
	for i := range records {
		if records[i].RoundNum != round || records[i].EventLogID <= previous {
			return AttachedTokenPage{}, errors.New("unordered or out-of-scope attached-token page")
		}
		mapped, err := mapAttachedToken(records[i])
		if err != nil {
			return AttachedTokenPage{}, err
		}
		data = append(data, mapped)
		previous = records[i].EventLogID
	}
	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 {
			return AttachedTokenPage{}, errors.New("attached-token page hasMore without row")
		}
		next, err := encodeClaimDetailCursor(claimDetailCursor{
			Version:    claimDetailCursorVersion,
			Round:      round,
			Section:    claimDetailAttached,
			EventLogID: previous,
		})
		if err != nil {
			return AttachedTokenPage{}, err
		}
		meta.NextCursor = &next
	}
	return AttachedTokenPage{Data: data, Meta: meta}, nil
}

func buildUnclaimedItemPage(
	records []cgstore.UnclaimedItemRecord,
	hasMore bool,
	round int64,
	after *cgstore.UnclaimedItemCursor,
	limit int,
) (UnclaimedItemPage, error) {
	if err := validatePageCardinality(len(records), limit); err != nil {
		return UnclaimedItemPage{}, err
	}
	data := make([]UnclaimedItem, 0, len(records))
	previous := cgstore.UnclaimedItemCursor{Segment: -1}
	if after != nil {
		previous = *after
	}
	for i := range records {
		current := cgstore.UnclaimedItemCursor{
			Segment: records[i].Segment,
			Key:     records[i].Key,
		}
		if records[i].RoundNum != round || !unclaimedCursorFollows(current, previous) {
			return UnclaimedItemPage{}, errors.New("unordered or out-of-scope unclaimed-item page")
		}
		mapped, err := mapUnclaimedItem(records[i])
		if err != nil {
			return UnclaimedItemPage{}, err
		}
		data = append(data, mapped)
		previous = current
	}
	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 {
			return UnclaimedItemPage{}, errors.New("unclaimed-item page hasMore without row")
		}
		next, err := encodeClaimDetailCursor(claimDetailCursor{
			Version: claimDetailCursorVersion,
			Round:   round,
			Section: claimDetailUnclaimed,
			Segment: previous.Segment,
			Key:     previous.Key,
		})
		if err != nil {
			return UnclaimedItemPage{}, err
		}
		meta.NextCursor = &next
	}
	return UnclaimedItemPage{Data: data, Meta: meta}, nil
}

func decodeClaimEventRequestCursor(
	encoded *string,
	round int64,
	section claimDetailSection,
	instance string,
) (*cgstore.ClaimEventCursor, *Problem) {
	if encoded == nil {
		return nil, nil
	}
	cursor, err := decodeClaimDetailCursor(*encoded, round, section)
	if err != nil {
		problem := newProblem(
			http.StatusBadRequest, "invalid-cursor", "Invalid cursor",
			"The cursor is malformed, unsupported, or belongs to another round or claim section.",
			instance,
		)
		return nil, &problem
	}
	return &cgstore.ClaimEventCursor{EventLogID: cursor.EventLogID}, nil
}

func decodeUnclaimedRequestCursor(
	encoded *string,
	round int64,
	instance string,
) (*cgstore.UnclaimedItemCursor, *Problem) {
	if encoded == nil {
		return nil, nil
	}
	cursor, err := decodeClaimDetailCursor(*encoded, round, claimDetailUnclaimed)
	if err != nil {
		problem := newProblem(
			http.StatusBadRequest, "invalid-cursor", "Invalid cursor",
			"The cursor is malformed, unsupported, or belongs to another round or claim section.",
			instance,
		)
		return nil, &problem
	}
	return &cgstore.UnclaimedItemCursor{Segment: cursor.Segment, Key: cursor.Key}, nil
}

func roiStoreSort(sortBy RoiLeaderboardSort) cgstore.ROILeaderboardSort {
	return cgstore.ROILeaderboardSort(sortBy)
}

func roiSecondaryValue(record cgstore.ROILeaderboardRecord, sortBy cgstore.ROILeaderboardSort) int64 {
	if sortBy == cgstore.ROILeaderboardWinRate {
		return record.RoundsParticipated
	}
	return 0
}

func roiRecordFollows(
	record cgstore.ROILeaderboardRecord,
	previous cgstore.ROILeaderboardPageCursor,
	sortBy cgstore.ROILeaderboardSort,
) bool {
	currentValue := cgstore.ROILeaderboardSortValue(record, sortBy)
	comparison, err := compareDecimal(currentValue, previous.SortValue)
	if err != nil || comparison > 0 {
		return false
	}
	if comparison < 0 {
		return true
	}
	if sortBy == cgstore.ROILeaderboardWinRate {
		if record.RoundsParticipated > previous.Secondary {
			return false
		}
		if record.RoundsParticipated < previous.Secondary {
			return true
		}
	}
	return record.BidderAid > previous.BidderAid
}

func claimSummaryFollows(current, previous cgstore.ClaimSummaryCursor) bool {
	return current.RoundNum < previous.RoundNum ||
		(current.RoundNum == previous.RoundNum && current.EventLogID < previous.EventLogID)
}

func unclaimedCursorFollows(current, previous cgstore.UnclaimedItemCursor) bool {
	return current.Segment > previous.Segment ||
		(current.Segment == previous.Segment && current.Key > previous.Key)
}

func getCosmicGameStatisticsInternal(problem Problem) GetCosmicGameStatisticsResponseObject {
	return GetCosmicGameStatistics500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}

func getCosmicGameCountersInternal(problem Problem) GetCosmicGameCountersResponseObject {
	return GetCosmicGameCounters500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}

func listROILeaderboardBadRequest(problem Problem) ListCosmicGameRoiLeaderboardResponseObject {
	return ListCosmicGameRoiLeaderboard400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func listROILeaderboardInternal(problem Problem) ListCosmicGameRoiLeaderboardResponseObject {
	return ListCosmicGameRoiLeaderboard500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}

func listClaimsBadRequest(problem Problem) ListCosmicGameClaimsResponseObject {
	return ListCosmicGameClaims400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func listClaimsInternal(problem Problem) ListCosmicGameClaimsResponseObject {
	return ListCosmicGameClaims500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}

func getRoundClaimsBadRequest(problem Problem) GetRoundClaimsResponseObject {
	return GetRoundClaims400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func getRoundClaimsNotFound(problem Problem) GetRoundClaimsResponseObject {
	return GetRoundClaims404ApplicationProblemPlusJSONResponse{
		NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(problem),
	}
}

func getRoundClaimsInternal(problem Problem) GetRoundClaimsResponseObject {
	return GetRoundClaims500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}
