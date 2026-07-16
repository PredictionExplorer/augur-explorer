// The RandomWalk beauty-contest ranking slice: four bounded reads and the
// first three v2 write operations (ADR-0008). Elo math, the canonical vote
// message and signer recovery live in internal/beautyrank, shared with the
// frozen v1 mini-app.

package v2

import (
	"context"
	"errors"
	"math"
	"net/http"
	"strings"

	"github.com/PredictionExplorer/augur-explorer/internal/beautyrank"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

const (
	// #nosec G101 -- route instances, not credentials.
	rankingRandomTokensInstance = "/api/v2/randomwalk/ranking/random-tokens"
	rankingPairInstance         = "/api/v2/randomwalk/ranking/pair"
	rankingRatingsInstance      = "/api/v2/randomwalk/ranking/ratings"
	rankingStatisticsInstance   = "/api/v2/randomwalk/ranking/statistics"
	rankingChallengesInstance   = "/api/v2/randomwalk/ranking/challenges"
	rankingVotesInstance        = "/api/v2/randomwalk/ranking/votes"
	rankingMatchesInstance      = "/api/v2/randomwalk/ranking/matches"
)

const (
	defaultRankingSampleSize = 2
	maxRankingSampleSize     = 100
	// maxRankingPairAttempts bounds the voter-aware re-roll, matching v1.
	maxRankingPairAttempts = 50
)

// GetRandomWalkRankingRandomTokens implements
// GET /api/v2/randomwalk/ranking/random-tokens.
func (s *Server) GetRandomWalkRankingRandomTokens(
	ctx context.Context,
	request GetRandomWalkRankingRandomTokensRequestObject,
) (GetRandomWalkRankingRandomTokensResponseObject, error) {
	internal := func() GetRandomWalkRankingRandomTokensResponseObject {
		return GetRandomWalkRankingRandomTokens500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(rankingRandomTokensInstance),
			),
		}
	}
	sampleSize := defaultRankingSampleSize
	if request.Params.SampleSize != nil {
		sampleSize = *request.Params.SampleSize
	}
	if sampleSize < 1 || sampleSize > maxRankingSampleSize {
		return GetRandomWalkRankingRandomTokens400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(
				invalidSampleSizeProblem(rankingRandomTokensInstance),
			),
		}, nil
	}
	ids, err := s.exploreRankingTokenIDs(ctx, sampleSize)
	if err != nil {
		s.logInternal(ctx, "get ranking random tokens", err)
		return internal(), nil
	}
	tokenIDs, err := mapRankingTokenIDs(ids, sampleSize)
	if err != nil {
		s.logInternal(ctx, "map ranking random tokens", err)
		return internal(), nil
	}
	return GetRandomWalkRankingRandomTokens200JSONResponse{
		RandomWalkRankingRandomTokensJSONResponse: RandomWalkRankingRandomTokensJSONResponse{
			TokenIds: tokenIDs,
		},
	}, nil
}

// GetRandomWalkRankingPair implements GET /api/v2/randomwalk/ranking/pair.
func (s *Server) GetRandomWalkRankingPair(
	ctx context.Context,
	request GetRandomWalkRankingPairRequestObject,
) (GetRandomWalkRankingPairResponseObject, error) {
	internal := func() GetRandomWalkRankingPairResponseObject {
		return GetRandomWalkRankingPair500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(rankingPairInstance),
			),
		}
	}

	voterAid := int64(0)
	if request.Params.Voter != nil {
		address, _, valid := userAddressInput(*request.Params.Voter)
		if !valid {
			return GetRandomWalkRankingPair400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(
					invalidUserAddressProblem(rankingPairInstance),
				),
			}, nil
		}
		aid, err := s.ranking.UserAddressID(ctx, address)
		switch {
		case errors.Is(err, store.ErrNotFound):
			// A wallet the indexer has never seen cannot have voted:
			// every pair is fresh, so no filtering is needed.
		case err != nil:
			s.logInternal(ctx, "resolve ranking voter", err, "voter", address)
			return internal(), nil
		default:
			voterAid = aid
		}
	}

	attempts := 1
	if voterAid > 0 {
		attempts = maxRankingPairAttempts
	}
	pairExhausted := voterAid > 0
	var ids []int64
	for attempt := 0; attempt < attempts; attempt++ {
		var err error
		ids, err = s.exploreRankingTokenIDs(ctx, 2)
		if err != nil {
			s.logInternal(ctx, "get ranking pair", err)
			return internal(), nil
		}
		if len(ids) < 2 {
			return GetRandomWalkRankingPair404ApplicationProblemPlusJSONResponse{
				NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(
					rankingPairUnavailableProblem(),
				),
			}, nil
		}
		if voterAid == 0 {
			break
		}
		voted, err := s.ranking.HasRankingVoteForVoterPair(ctx, voterAid, ids[0], ids[1])
		if err != nil {
			s.logInternal(ctx, "check ranking pair vote", err)
			return internal(), nil
		}
		if !voted {
			pairExhausted = false
			break
		}
	}
	if ids[0] < 0 || ids[1] < 0 || ids[0] == ids[1] {
		s.logInternal(ctx, "map ranking pair", errors.New("repository returned an invalid pair"),
			"first", ids[0], "second", ids[1])
		return internal(), nil
	}
	return GetRandomWalkRankingPair200JSONResponse{
		RandomWalkRankingPairJSONResponse: RandomWalkRankingPairJSONResponse{
			FirstTokenId:  ids[0],
			SecondTokenId: ids[1],
			PairExhausted: pairExhausted,
		},
	}, nil
}

// ListRandomWalkRankingRatings implements
// GET /api/v2/randomwalk/ranking/ratings.
func (s *Server) ListRandomWalkRankingRatings(
	ctx context.Context,
	request ListRandomWalkRankingRatingsRequestObject,
) (ListRandomWalkRankingRatingsResponseObject, error) {
	badRequest := func(problem Problem) ListRandomWalkRankingRatingsResponseObject {
		return ListRandomWalkRankingRatings400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
		}
	}
	internal := func() ListRandomWalkRankingRatingsResponseObject {
		return ListRandomWalkRankingRatings500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(rankingRatingsInstance),
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
			rankingRatingsInstance,
		)), nil
	}
	var after *rwstore.RankingRatingPageCursor
	if request.Params.Cursor != nil {
		cursor, err := decodeRandomWalkRankingRatingCursor(*request.Params.Cursor)
		if err != nil {
			return badRequest(invalidRandomWalkCursorProblem(rankingRatingsInstance)), nil
		}
		after = &rwstore.RankingRatingPageCursor{Rating: cursor.Rating, TokenID: cursor.TokenID}
	}
	addrs, err := s.ranking.ContractAddrs(ctx)
	if err != nil {
		s.logInternal(ctx, "resolve ranking contract registry", err)
		return internal(), nil
	}
	records, hasMore, err := s.ranking.RankingRatingsPage(ctx, addrs.RandomWalkAid, after, limit)
	if err != nil {
		s.logInternal(ctx, "list ranking ratings", err)
		return internal(), nil
	}
	data, err := mapRankingRatings(records, after, limit)
	if err != nil {
		s.logInternal(ctx, "map ranking ratings", err)
		return internal(), nil
	}
	meta := PageMeta{Limit: limit}
	if hasMore {
		last := records[len(records)-1]
		encoded, err := encodeRandomWalkRankingRatingCursor(randomWalkRankingRatingCursor{
			Version:  randomWalkCursorVersion,
			Resource: randomWalkResourceRankingRatings,
			Rating:   last.Rating,
			TokenID:  last.TokenID,
		})
		if err != nil {
			s.logInternal(ctx, "encode ranking rating cursor", err)
			return internal(), nil
		}
		meta.NextCursor = &encoded
	}
	return ListRandomWalkRankingRatings200JSONResponse{
		RandomWalkRankingRatingPageJSONResponse: RandomWalkRankingRatingPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// GetRandomWalkRankingStatistics implements
// GET /api/v2/randomwalk/ranking/statistics.
func (s *Server) GetRandomWalkRankingStatistics(
	ctx context.Context,
	_ GetRandomWalkRankingStatisticsRequestObject,
) (GetRandomWalkRankingStatisticsResponseObject, error) {
	internal := func() GetRandomWalkRankingStatisticsResponseObject {
		return GetRandomWalkRankingStatistics500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(rankingStatisticsInstance),
			),
		}
	}
	record, err := s.ranking.RankingStatistics(ctx)
	if err != nil {
		s.logInternal(ctx, "get ranking statistics", err)
		return internal(), nil
	}
	statistics, err := mapRankingStatistics(record)
	if err != nil {
		s.logInternal(ctx, "map ranking statistics", err)
		return internal(), nil
	}
	return GetRandomWalkRankingStatistics200JSONResponse{
		RandomWalkRankingStatisticsJSONResponse: RandomWalkRankingStatisticsJSONResponse(statistics),
	}, nil
}

// CreateRandomWalkRankingChallenge implements
// POST /api/v2/randomwalk/ranking/challenges.
func (s *Server) CreateRandomWalkRankingChallenge(
	ctx context.Context,
	_ CreateRandomWalkRankingChallengeRequestObject,
) (CreateRandomWalkRankingChallengeResponseObject, error) {
	internal := func() CreateRandomWalkRankingChallengeResponseObject {
		return CreateRandomWalkRankingChallenge500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(rankingChallengesInstance),
			),
		}
	}
	nonce, err := beautyrank.NewNonce(s.entropy)
	if err != nil {
		s.logInternal(ctx, "generate ranking challenge nonce", err)
		return internal(), nil
	}
	// The expiry comes from the database clock — the same clock the vote's
	// nonce consumption compares against — so validity cannot drift with
	// the process clock.
	expiresAt, err := s.ranking.CreateRankingVoteNonce(ctx, nonce, beautyrank.ChallengeTTL)
	if err != nil {
		s.logInternal(ctx, "store ranking challenge nonce", err)
		return internal(), nil
	}
	return CreateRandomWalkRankingChallenge201JSONResponse{
		RandomWalkRankingChallengeCreatedJSONResponse: RandomWalkRankingChallengeCreatedJSONResponse{
			Nonce:     nonce,
			ExpiresAt: expiresAt,
		},
	}, nil
}

// CreateRandomWalkRankingVote implements
// POST /api/v2/randomwalk/ranking/votes.
func (s *Server) CreateRandomWalkRankingVote(
	ctx context.Context,
	request CreateRandomWalkRankingVoteRequestObject,
) (CreateRandomWalkRankingVoteResponseObject, error) {
	badRequest := func(problem Problem) CreateRandomWalkRankingVoteResponseObject {
		return CreateRandomWalkRankingVote400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
		}
	}
	internal := func() CreateRandomWalkRankingVoteResponseObject {
		return CreateRandomWalkRankingVote500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(rankingVotesInstance),
			),
		}
	}
	body := request.Body
	if detail := rankingPairProblemDetail(body.FirstTokenId, body.SecondTokenId, body.WinnerTokenId); detail != "" {
		return badRequest(invalidRankingPairProblem(rankingVotesInstance, detail)), nil
	}
	nonce := strings.TrimSpace(body.Nonce)
	if nonce == "" {
		return badRequest(invalidNonceProblem(rankingVotesInstance, "A challenge nonce is required.")), nil
	}
	if strings.TrimSpace(body.Signature) == "" {
		return badRequest(invalidSignatureProblem(rankingVotesInstance, "A wallet signature is required.")), nil
	}
	if !s.rankingConfig.voteChainAllowed(body.ChainId) {
		return badRequest(newProblem(
			http.StatusBadRequest,
			"chain-not-allowed",
			"Chain not allowed",
			"Votes signed on this chain id are not accepted.",
			rankingVotesInstance,
		)), nil
	}

	message := beautyrank.VoteMessage(
		body.ChainId, nonce, body.FirstTokenId, body.SecondTokenId, body.WinnerTokenId)
	signer, err := beautyrank.RecoverSigner(message, body.Signature)
	if err != nil {
		return badRequest(invalidSignatureProblem(
			rankingVotesInstance,
			"The signature does not verify against the canonical vote message.",
		)), nil
	}

	firstWon := body.WinnerTokenId == body.FirstTokenId
	newFirst, newSecond, err := s.computeRankingOutcome(
		ctx, body.FirstTokenId, body.SecondTokenId, firstWon)
	if err != nil {
		s.logInternal(ctx, "compute ranking vote outcome", err)
		return internal(), nil
	}

	voterAid, err := s.ranking.EnsureVoterAddress(ctx, signer.Hex())
	if err != nil {
		s.logInternal(ctx, "resolve ranking vote signer", err)
		return internal(), nil
	}

	err = s.ranking.RecordSignedRankingVote(
		ctx, nonce, body.FirstTokenId, body.SecondTokenId, firstWon, newFirst, newSecond, voterAid)
	switch {
	case errors.Is(err, rwstore.ErrRankingNonceInvalid):
		return badRequest(invalidNonceProblem(
			rankingVotesInstance,
			"The challenge nonce is unknown, expired or already used.",
		)), nil
	case errors.Is(err, store.ErrConflict):
		return CreateRandomWalkRankingVote409ApplicationProblemPlusJSONResponse{
			ConflictApplicationProblemPlusJSONResponse: ConflictApplicationProblemPlusJSONResponse(newProblem(
				http.StatusConflict,
				"already-voted",
				"Already voted",
				"This wallet already voted on this pair.",
				rankingVotesInstance,
			)),
		}, nil
	case err != nil:
		s.logInternal(ctx, "record ranking vote", err)
		return internal(), nil
	}

	return CreateRandomWalkRankingVote201JSONResponse{
		RandomWalkRankingVoteRecordedJSONResponse: RandomWalkRankingVoteRecordedJSONResponse{
			FirstToken:    RankingRatingChange{TokenId: body.FirstTokenId, Rating: newFirst},
			SecondToken:   RankingRatingChange{TokenId: body.SecondTokenId, Rating: newSecond},
			WinnerTokenId: body.WinnerTokenId,
			VoterAddress:  signer.Hex(),
		},
	}, nil
}

// RecordRandomWalkRankingMatch implements
// POST /api/v2/randomwalk/ranking/matches. Authentication is enforced by
// the adminKeyMiddleware before this handler runs.
func (s *Server) RecordRandomWalkRankingMatch(
	ctx context.Context,
	request RecordRandomWalkRankingMatchRequestObject,
) (RecordRandomWalkRankingMatchResponseObject, error) {
	internal := func() RecordRandomWalkRankingMatchResponseObject {
		return RecordRandomWalkRankingMatch500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(rankingMatchesInstance),
			),
		}
	}
	body := request.Body
	if detail := rankingPairProblemDetail(body.FirstTokenId, body.SecondTokenId, body.WinnerTokenId); detail != "" {
		return RecordRandomWalkRankingMatch400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(
				invalidRankingPairProblem(rankingMatchesInstance, detail),
			),
		}, nil
	}
	firstWon := body.WinnerTokenId == body.FirstTokenId
	newFirst, newSecond, err := s.computeRankingOutcome(
		ctx, body.FirstTokenId, body.SecondTokenId, firstWon)
	if err != nil {
		s.logInternal(ctx, "compute ranking match outcome", err)
		return internal(), nil
	}
	if err := s.ranking.RecordRankingMatch(
		ctx, body.FirstTokenId, body.SecondTokenId, firstWon, newFirst, newSecond); err != nil {
		s.logInternal(ctx, "record ranking match", err)
		return internal(), nil
	}
	return RecordRandomWalkRankingMatch201JSONResponse{
		RandomWalkRankingMatchRecordedJSONResponse: RandomWalkRankingMatchRecordedJSONResponse{
			FirstToken:    RankingRatingChange{TokenId: body.FirstTokenId, Rating: newFirst},
			SecondToken:   RankingRatingChange{TokenId: body.SecondTokenId, Rating: newSecond},
			WinnerTokenId: body.WinnerTokenId,
		},
	}, nil
}

// exploreRankingTokenIDs samples limit tokens for the explorer and pair
// selection, scoped to the configured exploration bound.
func (s *Server) exploreRankingTokenIDs(ctx context.Context, limit int) ([]int64, error) {
	addrs, err := s.ranking.ContractAddrs(ctx)
	if err != nil {
		return nil, err
	}
	return s.ranking.ExploreRandomTokenIDs(
		ctx, addrs.RandomWalkAid, s.rankingConfig.exploreMaxTokenID(), limit)
}

// computeRankingOutcome reads the match count and both current ratings,
// then applies the shared Elo update.
func (s *Server) computeRankingOutcome(
	ctx context.Context,
	first, second int64,
	firstWon bool,
) (newFirst, newSecond float64, err error) {
	total, err := s.ranking.CountRankingMatches(ctx)
	if err != nil {
		return 0, 0, err
	}
	ratingFirst, ratingSecond, err := s.ranking.RatingPair(ctx, first, second)
	if err != nil {
		return 0, 0, err
	}
	score := 0.0
	if firstWon {
		score = 1.0
	}
	newFirst, newSecond = beautyrank.EloUpdate(ratingFirst, ratingSecond, score, total)
	return newFirst, newSecond, nil
}

// rankingPairProblemDetail validates a (first, second, winner) triple and
// returns the problem detail for the first violated rule, or "" when the
// triple is valid.
func rankingPairProblemDetail(first, second, winner int64) string {
	switch {
	case first < 0 || second < 0:
		return "Token ids must be zero or greater."
	case first == second:
		return "The two token ids must differ."
	case winner != first && winner != second:
		return "winnerTokenId must equal firstTokenId or secondTokenId."
	default:
		return ""
	}
}

func invalidRankingPairProblem(instance, detail string) Problem {
	return newProblem(http.StatusBadRequest, "invalid-pair", "Invalid pair", detail, instance)
}

func invalidNonceProblem(instance, detail string) Problem {
	return newProblem(http.StatusBadRequest, "invalid-nonce", "Invalid nonce", detail, instance)
}

func invalidSignatureProblem(instance, detail string) Problem {
	return newProblem(http.StatusBadRequest, "invalid-signature", "Invalid signature", detail, instance)
}

func invalidSampleSizeProblem(instance string) Problem {
	return newProblem(
		http.StatusBadRequest,
		"invalid-parameter",
		"Invalid parameter",
		"sampleSize must be between 1 and 100.",
		instance,
	)
}

func rankingPairUnavailableProblem() Problem {
	return newProblem(
		http.StatusNotFound,
		"ranking-pair-unavailable",
		"Ranking pair unavailable",
		"The collection does not hold two tokens to compare.",
		rankingPairInstance,
	)
}

// mapRankingTokenIDs enforces the exploration contract: at most sampleSize
// non-negative distinct ids. The empty collection maps to an empty list.
func mapRankingTokenIDs(ids []int64, sampleSize int) ([]int64, error) {
	if len(ids) > sampleSize {
		return nil, errors.New("repository returned more tokens than requested")
	}
	seen := make(map[int64]struct{}, len(ids))
	out := make([]int64, 0, len(ids))
	for _, id := range ids {
		if id < 0 {
			return nil, errors.New("repository returned a negative token id")
		}
		if _, dup := seen[id]; dup {
			return nil, errors.New("repository returned a duplicate token id")
		}
		seen[id] = struct{}{}
		out = append(out, id)
	}
	return out, nil
}

// mapRankingRatings enforces the directory contract: at most limit rows in
// strictly ascending (rating, tokenId) order resuming after the cursor,
// with finite ratings and non-negative counts.
func mapRankingRatings(
	records []rwstore.RankingRatingRecord,
	after *rwstore.RankingRatingPageCursor,
	limit int,
) ([]RankingRating, error) {
	if err := validatePageCardinality(len(records), limit); err != nil {
		return nil, err
	}
	previousRating := math.Inf(-1)
	previousToken := int64(-1)
	if after != nil {
		previousRating, previousToken = after.Rating, after.TokenID
	}
	data := make([]RankingRating, 0, len(records))
	for i := range records {
		record := records[i]
		if record.TokenID < 0 {
			return nil, errors.New("negative token id in rating directory")
		}
		if record.MatchCount < 0 {
			return nil, errors.New("negative match count in rating directory")
		}
		if math.IsNaN(record.Rating) || math.IsInf(record.Rating, 0) {
			return nil, errors.New("non-finite rating in rating directory")
		}
		if record.Rating < previousRating ||
			(record.Rating == previousRating && record.TokenID <= previousToken) {
			return nil, errors.New("rating directory out of order")
		}
		data = append(data, RankingRating{
			TokenId:    record.TokenID,
			Rating:     record.Rating,
			MatchCount: record.MatchCount,
		})
		previousRating, previousToken = record.Rating, record.TokenID
	}
	return data, nil
}

// mapRankingStatistics enforces the snapshot invariants: wallet votes are
// a subset of all votes, and voters exist exactly when wallet votes do.
func mapRankingStatistics(record rwstore.RankingStatisticsRecord) (RankingStatistics, error) {
	if record.TotalVotes < 0 || record.WalletVotes < 0 ||
		record.DistinctVoters < 0 || record.RatedTokens < 0 {
		return RankingStatistics{}, errors.New("negative ranking counter")
	}
	if record.WalletVotes > record.TotalVotes {
		return RankingStatistics{}, errors.New("wallet votes exceed total votes")
	}
	if record.DistinctVoters > record.WalletVotes {
		return RankingStatistics{}, errors.New("distinct voters exceed wallet votes")
	}
	if (record.WalletVotes == 0) != (record.DistinctVoters == 0) {
		return RankingStatistics{}, errors.New("wallet votes and distinct voters disagree")
	}
	return RankingStatistics{
		TotalVotes:     record.TotalVotes,
		WalletVotes:    record.WalletVotes,
		DistinctVoters: record.DistinctVoters,
		RatedTokens:    record.RatedTokens,
	}, nil
}
