package v2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

const bannedBidsInstance = "/api/v2/cosmicgame/moderation/banned-bids"

// ListCosmicGameBannedBids implements
// GET /api/v2/cosmicgame/moderation/banned-bids.
func (s *Server) ListCosmicGameBannedBids(
	ctx context.Context,
	request ListCosmicGameBannedBidsRequestObject,
) (ListCosmicGameBannedBidsResponseObject, error) {
	badRequest := func(problem Problem) ListCosmicGameBannedBidsResponseObject {
		return ListCosmicGameBannedBids400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
		}
	}
	internal := func() ListCosmicGameBannedBidsResponseObject {
		return ListCosmicGameBannedBids500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(bannedBidsInstance),
			),
		}
	}

	limit, valid := resolvePageLimit(request.Params.Limit)
	if !valid {
		return badRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			bannedBidsInstance,
		)), nil
	}

	var after *cgstore.BannedBidPageCursor
	if request.Params.Cursor != nil {
		cursor, err := decodeBidBanCursor(*request.Params.Cursor)
		if err != nil {
			return badRequest(newProblem(
				http.StatusBadRequest,
				"invalid-cursor",
				"Invalid cursor",
				"The cursor is malformed or uses an unsupported version.",
				bannedBidsInstance,
			)), nil
		}
		after = &cgstore.BannedBidPageCursor{ID: cursor.ID}
	}

	records, hasMore, err := s.bids.BannedBidsPage(ctx, after, limit)
	if err != nil {
		s.logInternal(ctx, "list banned bids", err)
		return internal(), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate banned-bid page cardinality", err)
		return internal(), nil
	}

	data := make([]BannedBid, 0, len(records))
	var previousID int64
	hasPrevious := false
	if after != nil {
		previousID = after.ID
		hasPrevious = true
	}
	for i := range records {
		record := records[i]
		if record.Id < 1 || (hasPrevious && record.Id >= previousID) {
			s.logInternal(ctx, "validate banned-bid page",
				errors.New("repository returned an unordered banned-bid row"),
				"bid_id", record.BidId,
				"row_id", record.Id)
			return internal(), nil
		}
		ban, err := mapBannedBid(record)
		if err != nil {
			s.logInternal(ctx, "map banned bid", err, "bid_id", record.BidId)
			return internal(), nil
		}
		data = append(data, ban)
		previousID = record.Id
		hasPrevious = true
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 || !hasPrevious {
			s.logInternal(ctx, "list banned bids",
				errors.New("repository reported another page without a cursor row"))
			return internal(), nil
		}
		next, err := encodeBidBanCursor(bidBanCursor{
			Version: bidBanCursorVersion,
			ID:      previousID,
		})
		if err != nil {
			s.logInternal(ctx, "encode banned-bid cursor", err)
			return internal(), nil
		}
		meta.NextCursor = &next
	}

	return ListCosmicGameBannedBids200JSONResponse{
		CosmicGameBannedBidPageJSONResponse: CosmicGameBannedBidPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// CreateCosmicGameBidBan implements
// POST /api/v2/cosmicgame/moderation/banned-bids.
func (s *Server) CreateCosmicGameBidBan(
	ctx context.Context,
	request CreateCosmicGameBidBanRequestObject,
) (CreateCosmicGameBidBanResponseObject, error) {
	if request.Body == nil || request.Body.BidId < 1 {
		return CreateCosmicGameBidBan400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(newProblem(
				http.StatusBadRequest,
				"invalid-bid",
				"Invalid bid",
				"Bid id must be one or greater.",
				bannedBidsInstance,
			)),
		}, nil
	}
	bidID := request.Body.BidId

	bidderAddress, err := s.bids.BidderAddressForBid(ctx, bidID)
	if errors.Is(err, store.ErrNotFound) {
		return CreateCosmicGameBidBan404ApplicationProblemPlusJSONResponse{
			NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(
				bidNotFoundProblem(bannedBidsInstance, bidID),
			),
		}, nil
	}
	if err != nil {
		s.logInternal(ctx, "resolve banned-bid owner", err, "bid_id", bidID)
		return createBidBanInternal(), nil
	}
	if !ethcommon.IsHexAddress(bidderAddress) {
		s.logInternal(ctx, "validate banned-bid owner",
			errors.New("repository returned an invalid bidder address"),
			"bid_id", bidID)
		return createBidBanInternal(), nil
	}
	bidderAddress = ethcommon.HexToAddress(bidderAddress).Hex()

	record, err := s.bids.CreateBannedBid(ctx, bidID, bidderAddress, s.now().UTC())
	if errors.Is(err, store.ErrConflict) {
		return CreateCosmicGameBidBan409ApplicationProblemPlusJSONResponse{
			ConflictApplicationProblemPlusJSONResponse: ConflictApplicationProblemPlusJSONResponse(newProblem(
				http.StatusConflict,
				"bid-already-banned",
				"Bid already banned",
				"An active ban already exists for this bid.",
				bannedBidsInstance,
			)),
		}, nil
	}
	if err != nil {
		s.logInternal(ctx, "create bid ban", err, "bid_id", bidID)
		return createBidBanInternal(), nil
	}
	if record.BidId != bidID || !strings.EqualFold(record.UserAddr, bidderAddress) {
		s.logInternal(ctx, "validate created bid ban",
			errors.New("repository returned a bid ban outside the requested identity"),
			"bid_id", bidID)
		return createBidBanInternal(), nil
	}
	ban, err := mapBannedBid(record)
	if err != nil {
		s.logInternal(ctx, "map created bid ban", err, "bid_id", bidID)
		return createBidBanInternal(), nil
	}

	return CreateCosmicGameBidBan201JSONResponse{
		CosmicGameBannedBidCreatedJSONResponse: CosmicGameBannedBidCreatedJSONResponse(ban),
	}, nil
}

// DeleteCosmicGameBidBan implements
// DELETE /api/v2/cosmicgame/moderation/banned-bids/{bidId}.
func (s *Server) DeleteCosmicGameBidBan(
	ctx context.Context,
	request DeleteCosmicGameBidBanRequestObject,
) (DeleteCosmicGameBidBanResponseObject, error) {
	instance := fmt.Sprintf("%s/%d", bannedBidsInstance, request.BidId)
	if request.BidId < 1 {
		return DeleteCosmicGameBidBan400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(newProblem(
				http.StatusBadRequest,
				"invalid-bid",
				"Invalid bid",
				"Bid id must be one or greater.",
				instance,
			)),
		}, nil
	}

	removed, err := s.bids.RemoveBannedBid(ctx, request.BidId)
	if err != nil {
		s.logInternal(ctx, "delete bid ban", err, "bid_id", request.BidId)
		return DeleteCosmicGameBidBan500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(instance),
			),
		}, nil
	}
	if !removed {
		return DeleteCosmicGameBidBan404ApplicationProblemPlusJSONResponse{
			NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(newProblem(
				http.StatusNotFound,
				"bid-ban-not-found",
				"Bid ban not found",
				"No active ban exists for this bid.",
				instance,
			)),
		}, nil
	}
	return DeleteCosmicGameBidBan204Response{}, nil
}

func mapBannedBid(record cgmodel.CGBannedBidRec) (BannedBid, error) {
	if record.Id < 1 || record.BidId < 1 || record.CreatedAt < 0 {
		return BannedBid{}, errors.New("invalid banned-bid identity")
	}
	if !ethcommon.IsHexAddress(record.UserAddr) {
		return BannedBid{}, errors.New("invalid banned-bid address")
	}
	bannedAt := time.Unix(record.CreatedAt, 0).UTC()
	if bannedAt.Year() < 1 || bannedAt.Year() > 9999 {
		return BannedBid{}, errors.New("invalid banned-bid timestamp")
	}
	return BannedBid{
		BannedAt:      bannedAt,
		BidId:         record.BidId,
		BidderAddress: ethcommon.HexToAddress(record.UserAddr).Hex(),
	}, nil
}

func bidNotFoundProblem(instance string, bidID int64) Problem {
	return newProblem(
		http.StatusNotFound,
		"bid-not-found",
		"Bid not found",
		fmt.Sprintf("No indexed bid exists with id %d.", bidID),
		instance,
	)
}

func createBidBanInternal() CreateCosmicGameBidBanResponseObject {
	return CreateCosmicGameBidBan500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
			internalProblem(bannedBidsInstance),
		),
	}
}
