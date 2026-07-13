// Package smoketest contains the context-aware API smoke-test engine used by
// opsctl.
package smoketest

import (
	"context"
	"database/sql"
	"fmt"
)

const zeroAddress = "0x0000000000000000000000000000000000000000"

// Params holds real URL parameter values used to exercise API endpoints.
type Params struct {
	UserAddress        string
	CSTStakerAddress   string
	RoundNumber        string
	BidEventLogID      string
	BidRound           string
	BidPosition        string
	TokenID            string
	TokenName          string
	ETHDonationID      string
	NFTDonationID      string
	ERC20DonationID    string
	CSTActionID        string
	RandomWalkActionID string
	DepositID          string
	NFTTokenAddress    string
	TimestampMin       string
	TimestampMax       string
	FromDate           string
	ToDate             string
}

// ParameterSource loads smoke-test URL parameters.
type ParameterSource interface {
	Parameters(ctx context.Context) (Params, error)
}

// SQLParameterSource loads parameters from the CosmicGame database. Individual
// missing/erroring queries retain the legacy type-valid fallback behavior.
type SQLParameterSource struct {
	DB *sql.DB
}

// DefaultParams returns type-valid fallback values for every endpoint.
func DefaultParams() Params {
	return Params{
		UserAddress:        zeroAddress,
		CSTStakerAddress:   zeroAddress,
		RoundNumber:        "0",
		BidEventLogID:      "1",
		BidRound:           "0",
		BidPosition:        "1",
		TokenID:            "0",
		TokenName:          "a",
		ETHDonationID:      "1",
		NFTDonationID:      "1",
		ERC20DonationID:    "1",
		CSTActionID:        "0",
		RandomWalkActionID: "0",
		DepositID:          "0",
		NFTTokenAddress:    zeroAddress,
		TimestampMin:       "0",
		TimestampMax:       "2000000000",
		FromDate:           "20230101",
		ToDate:             "20300101",
	}
}

// WithDefaults fills missing fields without overwriting values supplied by a
// parameter source.
func WithDefaults(params Params) Params {
	defaults := DefaultParams()
	fill := func(value *string, fallback string) {
		if *value == "" {
			*value = fallback
		}
	}
	fill(&params.UserAddress, defaults.UserAddress)
	fill(&params.CSTStakerAddress, params.UserAddress)
	fill(&params.RoundNumber, defaults.RoundNumber)
	fill(&params.BidEventLogID, defaults.BidEventLogID)
	fill(&params.BidRound, params.RoundNumber)
	fill(&params.BidPosition, defaults.BidPosition)
	fill(&params.TokenID, defaults.TokenID)
	fill(&params.TokenName, defaults.TokenName)
	fill(&params.ETHDonationID, defaults.ETHDonationID)
	fill(&params.NFTDonationID, defaults.NFTDonationID)
	fill(&params.ERC20DonationID, defaults.ERC20DonationID)
	fill(&params.CSTActionID, defaults.CSTActionID)
	fill(&params.RandomWalkActionID, defaults.RandomWalkActionID)
	fill(&params.DepositID, defaults.DepositID)
	fill(&params.NFTTokenAddress, defaults.NFTTokenAddress)
	fill(&params.TimestampMin, defaults.TimestampMin)
	fill(&params.TimestampMax, defaults.TimestampMax)
	fill(&params.FromDate, defaults.FromDate)
	fill(&params.ToDate, defaults.ToDate)
	return params
}

// Parameters implements ParameterSource.
func (s SQLParameterSource) Parameters(ctx context.Context) (Params, error) {
	if s.DB == nil {
		return Params{}, fmt.Errorf("parameter database is nil")
	}
	return loadParameters(func(query, fallback string) (string, error) {
		return queryValue(ctx, s.DB, query, fallback)
	})
}

func loadParameters(queryValue func(query, fallback string) (string, error)) (Params, error) {
	params := DefaultParams()
	load := func(dst *string, query, fallback string) error {
		value, err := queryValue(query, fallback)
		if err != nil {
			return err
		}
		*dst = value
		return nil
	}

	if err := load(&params.UserAddress,
		"SELECT a.addr FROM cg_bid b JOIN address a ON a.address_id=b.bidder_aid LIMIT 1",
		params.UserAddress); err != nil {
		return Params{}, err
	}
	if err := load(&params.CSTStakerAddress,
		"SELECT a.addr FROM cg_nft_staked_cst s JOIN address a ON a.address_id=s.staker_aid LIMIT 1",
		params.UserAddress); err != nil {
		return Params{}, err
	}

	latestBidRound, err := queryValue(
		"SELECT round_num FROM cg_bid ORDER BY round_num DESC LIMIT 1", params.RoundNumber)
	if err != nil {
		return Params{}, err
	}
	if err := load(&params.RoundNumber,
		"SELECT round_num FROM cg_prize_claim ORDER BY round_num DESC LIMIT 1",
		latestBidRound); err != nil {
		return Params{}, err
	}
	if err := load(&params.BidEventLogID,
		"SELECT evtlog_id FROM cg_bid ORDER BY id DESC LIMIT 1", params.BidEventLogID); err != nil {
		return Params{}, err
	}
	if err := load(&params.BidRound,
		"SELECT round_num FROM cg_bid ORDER BY id DESC LIMIT 1", params.RoundNumber); err != nil {
		return Params{}, err
	}
	if err := load(&params.BidPosition,
		"SELECT bid_position FROM cg_bid ORDER BY id DESC LIMIT 1", params.BidPosition); err != nil {
		return Params{}, err
	}
	if err := load(&params.TokenID,
		"SELECT token_id FROM cg_mint_event ORDER BY token_id DESC LIMIT 1", params.TokenID); err != nil {
		return Params{}, err
	}
	if err := load(&params.TokenName,
		"SELECT token_name FROM cg_token_name WHERE COALESCE(token_name,'')<>'' LIMIT 1", params.TokenName); err != nil {
		return Params{}, err
	}
	if err := load(&params.ETHDonationID,
		"SELECT id FROM cg_eth_donated_wi ORDER BY id DESC LIMIT 1", params.ETHDonationID); err != nil {
		return Params{}, err
	}
	if err := load(&params.NFTDonationID,
		"SELECT id FROM cg_nft_donation ORDER BY id DESC LIMIT 1", params.NFTDonationID); err != nil {
		return Params{}, err
	}
	if err := load(&params.ERC20DonationID,
		"SELECT id FROM cg_erc20_donation ORDER BY id DESC LIMIT 1", params.ERC20DonationID); err != nil {
		return Params{}, err
	}
	if err := load(&params.CSTActionID,
		"SELECT action_id FROM cg_nft_staked_cst ORDER BY action_id DESC LIMIT 1", params.CSTActionID); err != nil {
		return Params{}, err
	}
	if err := load(&params.RandomWalkActionID,
		"SELECT action_id FROM cg_nft_staked_rwalk ORDER BY action_id DESC LIMIT 1", params.RandomWalkActionID); err != nil {
		return Params{}, err
	}
	if err := load(&params.DepositID,
		"SELECT deposit_id FROM cg_staking_eth_deposit ORDER BY deposit_id DESC LIMIT 1", params.DepositID); err != nil {
		return Params{}, err
	}
	if err := load(&params.NFTTokenAddress,
		"SELECT a.addr FROM cg_nft_donation d JOIN address a ON a.address_id=d.token_aid LIMIT 1",
		params.NFTTokenAddress); err != nil {
		return Params{}, err
	}
	if err := load(&params.TimestampMin,
		"SELECT EXTRACT(EPOCH FROM MIN(time_stamp))::bigint::text FROM cg_bid", params.TimestampMin); err != nil {
		return Params{}, err
	}
	if err := load(&params.TimestampMax,
		"SELECT EXTRACT(EPOCH FROM MAX(time_stamp))::bigint::text FROM cg_bid", params.TimestampMax); err != nil {
		return Params{}, err
	}
	if err := load(&params.FromDate,
		"SELECT to_char(MIN(time_stamp),'YYYYMMDD') FROM cg_bid", params.FromDate); err != nil {
		return Params{}, err
	}
	if err := load(&params.ToDate,
		"SELECT to_char(MAX(time_stamp) + interval '1 day','YYYYMMDD') FROM cg_bid", params.ToDate); err != nil {
		return Params{}, err
	}
	return WithDefaults(params), nil
}

func queryValue(ctx context.Context, db *sql.DB, query, fallback string) (string, error) {
	var value sql.NullString
	if err := db.QueryRowContext(ctx, query).Scan(&value); err != nil {
		if ctxErr := ctx.Err(); ctxErr != nil {
			return "", ctxErr
		}
		return fallback, nil
	}
	if !value.Valid || value.String == "" {
		return fallback, nil
	}
	return value.String, nil
}
