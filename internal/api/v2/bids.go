package v2

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"

	cgprimitives "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// ListRoundBids implements GET /api/v2/cosmicgame/rounds/{round}/bids.
func (s *Server) ListRoundBids(ctx context.Context, request ListRoundBidsRequestObject) (ListRoundBidsResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/rounds/%d/bids", request.Round)
	if request.Round < 0 {
		return listBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			"Round must be zero or greater.",
			instance,
		)), nil
	}

	limit, validLimit := resolvePageLimit(request.Params.Limit)
	if !validLimit {
		return listBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			instance,
		)), nil
	}

	after := cgstore.BidPageCursor{}
	if request.Params.Cursor != nil {
		cursor, err := decodeBidCursor(*request.Params.Cursor, request.Round)
		if err != nil {
			return listBadRequest(newProblem(
				http.StatusBadRequest,
				"invalid-cursor",
				"Invalid cursor",
				"The cursor is malformed, unsupported, or belongs to another round.",
				instance,
			)), nil
		}
		after = cgstore.BidPageCursor{
			BidPosition: cursor.BidPosition,
			EventLogID:  cursor.EventLogID,
		}
	}

	records, hasMore, err := s.bids.BidsByRoundPage(ctx, request.Round, after, limit)
	if err != nil {
		s.logInternal(ctx, "list round bids", err, "round", request.Round)
		return listInternal(internalProblem(instance)), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate round bid page cardinality", err, "round", request.Round)
		return listInternal(internalProblem(instance)), nil
	}

	data := make([]Bid, 0, len(records))
	previous := after
	for i := range records {
		if records[i].RoundNum != request.Round ||
			records[i].BidPosition < previous.BidPosition ||
			(records[i].BidPosition == previous.BidPosition && records[i].Tx.EvtLogId <= previous.EventLogID) {
			err := errors.New("repository returned an out-of-scope or unordered bid page")
			s.logInternal(ctx, "validate round bid page", err,
				"round", request.Round,
				"event_log_id", records[i].Tx.EvtLogId)
			return listInternal(internalProblem(instance)), nil
		}
		bid, err := mapBid(records[i])
		if err != nil {
			s.logInternal(ctx, "map round bid", err,
				"round", request.Round,
				"event_log_id", records[i].Tx.EvtLogId)
			return listInternal(internalProblem(instance)), nil
		}
		data = append(data, bid)
		previous = cgstore.BidPageCursor{
			BidPosition: records[i].BidPosition,
			EventLogID:  records[i].Tx.EvtLogId,
		}
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 {
			err := errors.New("repository reported another page without returning a cursor row")
			s.logInternal(ctx, "list round bids", err, "round", request.Round)
			return listInternal(internalProblem(instance)), nil
		}
		last := records[len(records)-1]
		next, err := encodeBidCursor(bidCursor{
			Version:     bidCursorVersion,
			Round:       request.Round,
			BidPosition: last.BidPosition,
			EventLogID:  last.Tx.EvtLogId,
		})
		if err != nil {
			s.logInternal(ctx, "encode round bid cursor", err, "round", request.Round)
			return listInternal(internalProblem(instance)), nil
		}
		meta.NextCursor = &next
	}

	return ListRoundBids200JSONResponse{
		RoundBidPageJSONResponse: RoundBidPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// GetRoundBid implements
// GET /api/v2/cosmicgame/rounds/{round}/bids/{position}.
func (s *Server) GetRoundBid(ctx context.Context, request GetRoundBidRequestObject) (GetRoundBidResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/rounds/%d/bids/%d", request.Round, request.Position)
	if request.Round < 0 || request.Position < 1 {
		return getBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			"Round must be zero or greater and position must be one or greater.",
			instance,
		)), nil
	}

	record, err := s.bids.BidByRoundAndPosition(ctx, request.Round, request.Position)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return GetRoundBid404ApplicationProblemPlusJSONResponse{
				NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(newProblem(
					http.StatusNotFound,
					"bid-not-found",
					"Bid not found",
					"No bid exists at that round and position.",
					instance,
				)),
			}, nil
		}
		s.logInternal(ctx, "get round bid", err,
			"round", request.Round,
			"position", request.Position)
		return getInternal(internalProblem(instance)), nil
	}
	if record.RoundNum != request.Round || record.BidPosition != request.Position {
		err := errors.New("repository returned a bid outside the requested identity")
		s.logInternal(ctx, "validate round bid", err,
			"round", request.Round,
			"position", request.Position,
			"event_log_id", record.Tx.EvtLogId)
		return getInternal(internalProblem(instance)), nil
	}

	bid, err := mapBid(record)
	if err != nil {
		s.logInternal(ctx, "map round bid", err,
			"round", request.Round,
			"position", request.Position,
			"event_log_id", record.Tx.EvtLogId)
		return getInternal(internalProblem(instance)), nil
	}
	return GetRoundBid200JSONResponse{BidJSONResponse: BidJSONResponse(bid)}, nil
}

func mapBid(record cgprimitives.CGBidRec) (Bid, error) {
	if record.Tx.EvtLogId < 1 || record.Tx.BlockNum < 0 || record.RoundNum < 0 || record.BidPosition < 1 {
		return Bid{}, errors.New("invalid bid identity fields")
	}
	occurredAt, err := time.Parse(time.RFC3339Nano, record.Tx.DateTime)
	if err != nil {
		return Bid{}, fmt.Errorf("parse bid timestamp: %w", err)
	}
	prizeAt, err := time.Parse(time.RFC3339Nano, record.PrizeTimeDate)
	if err != nil {
		return Bid{}, fmt.Errorf("parse prize timestamp: %w", err)
	}
	if !ethcommon.IsHexAddress(record.BidderAddr) {
		return Bid{}, errors.New("invalid bidder address")
	}
	if !isTransactionHash(record.Tx.TxHash) {
		return Bid{}, errors.New("invalid transaction hash")
	}
	cstReward, err := requiredAmount(record.CSTReward)
	if err != nil {
		return Bid{}, fmt.Errorf("cst reward: %w", err)
	}

	bid := Bid{
		BidType:         mapBidType(record.BidType),
		BidderAddress:   ethcommon.HexToAddress(record.BidderAddr).Hex(),
		BlockNumber:     record.Tx.BlockNum,
		CstRewardWei:    cstReward,
		EventLogId:      record.Tx.EvtLogId,
		OccurredAt:      occurredAt.UTC(),
		Position:        record.BidPosition,
		PrizeAt:         prizeAt.UTC(),
		Round:           record.RoundNum,
		TransactionHash: strings.ToLower(record.Tx.TxHash),
	}

	if bid.EthPriceWei, err = optionalAmount(record.EthPrice); err != nil {
		return Bid{}, fmt.Errorf("eth price: %w", err)
	}
	if bid.CstPriceWei, err = optionalAmount(record.CstPrice); err != nil {
		return Bid{}, fmt.Errorf("cst price: %w", err)
	}
	if bid.BidCstRewardAmountWei, err = optionalAmount(record.BidCstRewardAmount); err != nil {
		return Bid{}, fmt.Errorf("bid cst reward: %w", err)
	}
	if record.CstDutchAuctionDurationInt >= 0 {
		value := record.CstDutchAuctionDurationInt
		bid.CstDutchAuctionDurationSeconds = &value
	}
	if record.RWalkNFTId >= 0 {
		tokenID := record.RWalkNFTId
		bid.RandomWalkTokenId = &tokenID
	}
	if record.Message != "" {
		message := record.Message
		bid.Message = &message
	}
	if record.NFTDonationTokenId >= 0 {
		if !ethcommon.IsHexAddress(record.NFTDonationTokenAddr) {
			return Bid{}, errors.New("invalid NFT donation token address")
		}
		bid.NftDonation = &NftDonation{
			TokenAddress: ethcommon.HexToAddress(record.NFTDonationTokenAddr).Hex(),
			TokenId:      record.NFTDonationTokenId,
			TokenUri:     record.NFTTokenURI,
		}
	}
	if record.DonatedERC20TokenAddr != "" || record.DonatedERC20TokenAmount != "" {
		if !ethcommon.IsHexAddress(record.DonatedERC20TokenAddr) {
			return Bid{}, errors.New("invalid ERC-20 donation token address")
		}
		amount, err := requiredAmount(record.DonatedERC20TokenAmount)
		if err != nil {
			return Bid{}, fmt.Errorf("ERC-20 donation amount: %w", err)
		}
		bid.Erc20Donation = &Erc20Donation{
			AmountWei:    amount,
			TokenAddress: ethcommon.HexToAddress(record.DonatedERC20TokenAddr).Hex(),
		}
	}
	return bid, nil
}

func mapBidType(value int64) BidType {
	switch value {
	case 0:
		return Eth
	case 1:
		return RandomWalk
	case 2:
		return Cst
	default:
		return Unknown
	}
}

func optionalAmount(value string) (*string, error) {
	if value == "-1" {
		return nil, nil
	}
	amount, err := requiredAmount(value)
	if err != nil {
		return nil, err
	}
	return &amount, nil
}

func requiredAmount(value string) (string, error) {
	if value == "" {
		return "", errors.New("amount is empty")
	}
	amount, ok := new(big.Int).SetString(value, 10)
	if !ok || amount.Sign() < 0 {
		return "", fmt.Errorf("invalid non-negative decimal %q", value)
	}
	return amount.String(), nil
}

func isTransactionHash(value string) bool {
	if len(value) != 66 || !strings.HasPrefix(value, "0x") {
		return false
	}
	decoded, err := hex.DecodeString(value[2:])
	return err == nil && len(decoded) == 32
}

func (s *Server) logInternal(ctx context.Context, message string, err error, attrs ...any) {
	if errors.Is(err, context.Canceled) {
		return
	}
	attrs = append(attrs, "error", err)
	s.logger.ErrorContext(ctx, message, attrs...)
}

func listBadRequest(problem Problem) ListRoundBidsResponseObject {
	return ListRoundBids400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func listInternal(problem Problem) ListRoundBidsResponseObject {
	return ListRoundBids500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}

func getBadRequest(problem Problem) GetRoundBidResponseObject {
	return GetRoundBid400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func getInternal(problem Problem) GetRoundBidResponseObject {
	return GetRoundBid500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}
