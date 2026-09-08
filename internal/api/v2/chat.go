package v2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// ListRoundMessages exposes bounded older pages and ascending catch-up pages.
func (s *Server) ListRoundMessages(ctx context.Context, request ListRoundMessagesRequestObject) (ListRoundMessagesResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/rounds/%d/messages", request.Round)
	bad := func(code, detail string) ListRoundMessagesResponseObject {
		return ListRoundMessages400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(newProblem(http.StatusBadRequest, code, "Invalid request", detail, instance)),
		}
	}
	fail := func(err error) ListRoundMessagesResponseObject {
		s.logInternal(ctx, "list round messages", err, "round", request.Round)
		return ListRoundMessages500ApplicationProblemPlusJSONResponse{InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(internalProblem(instance))}
	}
	limit, valid := resolvePageLimit(request.Params.Limit)
	if request.Round < 0 {
		return bad("invalid-parameter", "Round must be zero or greater."), nil
	}
	if !valid {
		return bad("invalid-parameter", pageLimitProblemDetail()), nil
	}
	if request.Params.Cursor != nil && request.Params.After != nil {
		return bad("invalid-parameter", "Cursor and after are mutually exclusive."), nil
	}
	direction, encoded := "older", request.Params.Cursor
	if request.Params.After != nil {
		direction, encoded = "after", request.Params.After
	}
	var position *cgstore.ChatPosition
	var revision int64
	if encoded != nil {
		cursor, err := decodeChatCursor(*encoded, request.Round, direction)
		if err != nil {
			return bad("invalid-cursor", "The cursor is malformed, unsupported, or belongs to another cycle or direction."), nil
		}
		position = &cgstore.ChatPosition{OccurredAt: cursor.OccurredAt, EventLogID: cursor.EventLogID}
		revision = cursor.Revision
	}
	after := direction == "after"
	page, err := s.bids.ChatMessagesPage(ctx, request.Round, position, after, limit)
	if err != nil {
		return fail(err), nil
	}
	if page.Revision < 1 {
		return fail(errors.New("invalid chat revision")), nil
	}
	if position != nil && revision != page.Revision {
		problem := newProblem(http.StatusConflict, "feed-reset-required", "Chat history changed", "Reload the message feed and chat context before continuing.", instance)
		return ListRoundMessages409ApplicationProblemPlusJSONResponse{ConflictApplicationProblemPlusJSONResponse: ConflictApplicationProblemPlusJSONResponse(problem)}, nil
	}
	if err := validatePageCardinality(len(page.Records), limit); err != nil {
		return fail(err), nil
	}
	if page.HasMore && len(page.Records) != limit {
		return fail(errors.New("incomplete chat continuation page")), nil
	}
	if !validChatCursor(chatCursor{Version: chatCursorVersion, Round: request.Round, Direction: "after", OccurredAt: page.Latest.OccurredAt, EventLogID: page.Latest.EventLogID, Revision: page.Revision}) {
		return fail(errors.New("invalid latest chat boundary")), nil
	}
	data := make([]ChatMessage, 0, len(page.Records))
	previous := position
	for _, record := range page.Records {
		if record.Round != request.Round || compareChatPosition(record.ChatPosition, page.Latest) > 0 {
			return fail(errors.New("chat row outside snapshot")), nil
		}
		if previous != nil {
			comparison := compareChatPosition(record.ChatPosition, *previous)
			if (after && comparison <= 0) || (!after && comparison >= 0) {
				return fail(errors.New("unordered chat page")), nil
			}
		}
		item, err := mapChatMessage(record)
		if err != nil {
			return fail(err), nil
		}
		data = append(data, item)
		p := record.ChatPosition
		previous = &p
	}
	boundary := page.Latest
	if after {
		boundary = *previous
	}
	syncCursor, err := encodeChatCursor(request.Round, "after", boundary, page.Revision)
	if err != nil {
		return fail(err), nil
	}
	meta := ChatPageMeta{Limit: limit, SyncCursor: syncCursor, HasMore: after && page.HasMore, Revision: strconv.FormatInt(page.Revision, 10)}
	if !after && page.HasMore {
		cursor, err := encodeChatCursor(request.Round, "older", *previous, page.Revision)
		if err != nil {
			return fail(err), nil
		}
		meta.NextCursor = &cursor
	}
	return ListRoundMessages200JSONResponse{Data: data, Meta: meta}, nil
}

func mapChatMessage(record cgstore.ChatMessageRecord) (ChatMessage, error) {
	if err := validateChatIdentity(record.ChatPosition, record.Round, record.Position, record.BidderAddress); err != nil {
		return ChatMessage{}, err
	}
	if !isTransactionHash(record.TransactionHash) || !chatHasText(record.Message) {
		return ChatMessage{}, errors.New("invalid chat message or transaction")
	}
	item := ChatMessage{
		EventLogId: record.EventLogID, Round: record.Round, Position: record.Position,
		BidderAddress: ethcommon.HexToAddress(record.BidderAddress).Hex(), OccurredAt: record.OccurredAt.UTC(),
		Message: record.Message, BidType: mapBidType(record.BidType), TransactionHash: strings.ToLower(record.TransactionHash),
	}
	var err error
	if item.EthPriceWei, err = optionalAmount(record.EthPriceWei); err != nil {
		return ChatMessage{}, err
	}
	if item.CstPriceWei, err = optionalAmount(record.CstPriceWei); err != nil {
		return ChatMessage{}, err
	}
	if record.RandomWalkTokenID < -1 {
		return ChatMessage{}, errors.New("invalid RandomWalk token ID")
	}
	if record.RandomWalkTokenID >= 0 {
		id := record.RandomWalkTokenID
		item.RandomWalkTokenId = &id
	}
	return item, nil
}

// ECMAScript trim's exact whitespace set, shared with migration 00030.
func chatHasText(value string) bool {
	return strings.Trim(value, "\u0009\u000a\u000b\u000c\u000d\u0020\u00a0\u1680\u2000\u2001\u2002\u2003\u2004\u2005\u2006\u2007\u2008\u2009\u200a\u2028\u2029\u202f\u205f\u3000\ufeff") != ""
}

func validateChatIdentity(position cgstore.ChatPosition, round, ordinal int64, address string) error {
	if position.EventLogID < 1 || position.OccurredAt.Before(time.Unix(0, 0)) || round < 0 || ordinal < 1 || !ethcommon.IsHexAddress(address) {
		return errors.New("invalid chat identity")
	}
	return nil
}

// GetRoundChatContext returns minimal complete metadata for existing client milestones.
func (s *Server) GetRoundChatContext(ctx context.Context, request GetRoundChatContextRequestObject) (GetRoundChatContextResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/rounds/%d/chat-context", request.Round)
	if request.Round < 0 {
		problem := newProblem(http.StatusBadRequest, "invalid-parameter", "Invalid request", "Round must be zero or greater.", instance)
		return GetRoundChatContext400ApplicationProblemPlusJSONResponse{BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem)}, nil
	}
	fail := func(err error) GetRoundChatContextResponseObject {
		s.logInternal(ctx, "get round chat context", err, "round", request.Round)
		return GetRoundChatContext500ApplicationProblemPlusJSONResponse{InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(internalProblem(instance))}
	}
	snapshot, err := s.bids.ChatContext(ctx, request.Round)
	if err != nil {
		return fail(err), nil
	}
	if snapshot.Revision < 1 {
		return fail(errors.New("invalid chat context revision")), nil
	}
	response := GetRoundChatContext200JSONResponse{Data: make([]ChatContextEntry, 0, len(snapshot.Records))}
	response.Meta.Revision = strconv.FormatInt(snapshot.Revision, 10)
	var previous *cgstore.ChatPosition
	for _, record := range snapshot.Records {
		if record.Round != request.Round || (previous != nil && compareChatPosition(record.ChatPosition, *previous) <= 0) || record.PrizeAt.IsZero() || record.CstDutchAuctionDurationSeconds < -1 {
			return fail(errors.New("invalid or unordered chat context")), nil
		}
		if err := validateChatIdentity(record.ChatPosition, record.Round, record.Position, record.BidderAddress); err != nil {
			return fail(err), nil
		}
		item := ChatContextEntry{
			EventLogId: record.EventLogID, Round: record.Round, Position: record.Position,
			BidderAddress: ethcommon.HexToAddress(record.BidderAddress).Hex(), OccurredAt: record.OccurredAt.UTC(),
			BidType: mapBidType(record.BidType), PrizeAt: record.PrizeAt.UTC(),
		}
		if record.CstDutchAuctionDurationSeconds >= 0 {
			duration := record.CstDutchAuctionDurationSeconds
			item.CstDutchAuctionDurationSeconds = &duration
		}
		response.Data = append(response.Data, item)
		p := record.ChatPosition
		previous = &p
	}
	return response, nil
}
