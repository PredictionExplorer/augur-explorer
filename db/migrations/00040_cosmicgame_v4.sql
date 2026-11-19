-- +goose Up
-- Cosmic Signature V4 (in development): the game proxy gains the art-auction
-- wiring events and the per-cycle ChaosZero market funding record, and a new
-- non-upgradeable CosmicSignatureArtAuction contract runs English auctions
-- over the eight extra Cosmic Signature NFTs each V4 round mints into its
-- custody. One table per event; the canonical column set matches the rest of
-- the cg_* event tables (evtlog_id references evt_log so a reorg cascade
-- removes the row; UNIQUE(evtlog_id) makes the writes idempotent).

-- ISystemEventsV4/ICosmicSignatureGameV4.sol:ArtAuctionConfigured (game proxy).
-- Announces the art-auction contract selected during initializeV4 (or by a
-- later upgrade). The ETL grows its watched-address set from these rows,
-- mirroring cg_adm_prizes_wallet_addr.
CREATE TABLE cg_art_auction_configured (
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	art_auction_aid BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);

-- ICosmicSignatureGameV4.sol:ArtAuctionOpeningBidChanged (game proxy).
CREATE TABLE cg_adm_art_auction_opening_bid (
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	opening_bid     DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);

-- ICosmicSignatureGameV4.sol:ArtAuctionNftsMinted (game proxy; one per
-- claimed round, alongside MainPrizeClaimed).
CREATE TABLE cg_art_auction_nfts_minted (
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	round_num       BIGINT NOT NULL,
	art_auction_aid BIGINT NOT NULL,
	first_nft_id    BIGINT NOT NULL,
	num_nfts        BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);

-- ICosmicSignatureGameV4.sol:ChaosZeroCycleFunded (game proxy; 8000 CST
-- minted and transferred to the ChaosZero market each cycle).
CREATE TABLE cg_chaos_zero_cycle_funded (
	id                BIGSERIAL PRIMARY KEY,
	evtlog_id         BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num         BIGINT NOT NULL,
	tx_id             BIGINT NOT NULL,
	time_stamp        TIMESTAMPTZ NOT NULL,
	contract_aid      BIGINT NOT NULL,
	market_aid        BIGINT NOT NULL,
	cycle             BIGINT NOT NULL,
	amount            DECIMAL NOT NULL,
	cumulative_amount DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);

-- ICosmicSignatureArtAuction.sol:AuctionBatchRegistered.
CREATE TABLE cg_art_auction_batch (
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	round_num       BIGINT NOT NULL,
	first_nft_id    BIGINT NOT NULL,
	num_nfts        BIGINT NOT NULL,
	eligible_at     TIMESTAMPTZ NOT NULL,
	opening_bid     DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);

-- ICosmicSignatureArtAuction.sol:AuctionBidPlaced. amount is the bidder's
-- total standing bid; eth_paid the ETH sent with this transaction (lower on
-- a self-raise); end_at the auction deadline after anti-sniping extension.
CREATE TABLE cg_art_auction_bid (
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	nft_id          BIGINT NOT NULL,
	round_num       BIGINT NOT NULL,
	bidder_aid      BIGINT NOT NULL,
	amount          DECIMAL NOT NULL,
	eth_paid        DECIMAL NOT NULL,
	end_at          TIMESTAMPTZ NOT NULL,
	UNIQUE(evtlog_id)
);

-- ICosmicSignatureArtAuction.sol:AuctionRefundCredited (outbid; pull-based).
CREATE TABLE cg_art_auction_refund_credited (
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	bidder_aid      BIGINT NOT NULL,
	nft_id          BIGINT NOT NULL,
	amount          DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);

-- ICosmicSignatureArtAuction.sol:AuctionRefundWithdrawn.
CREATE TABLE cg_art_auction_refund_withdrawn (
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	bidder_aid      BIGINT NOT NULL,
	recipient_aid   BIGINT NOT NULL,
	amount          DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);

-- ICosmicSignatureArtAuction.sol:AuctionSettled.
CREATE TABLE cg_art_auction_settled (
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	nft_id          BIGINT NOT NULL,
	round_num       BIGINT NOT NULL,
	winner_aid      BIGINT NOT NULL,
	amount          DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);

-- ICosmicSignatureArtAuction.sol:AuctionNftDelivered (winner chose the
-- receiving address).
CREATE TABLE cg_art_auction_nft_delivered (
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	nft_id          BIGINT NOT NULL,
	winner_aid      BIGINT NOT NULL,
	recipient_aid   BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);

-- ICosmicSignatureArtAuction.sol:AuctionProceedsForwarded (winning bid sent
-- back into the game via donateEth; shows up as an EthDonated row from the
-- art-auction contract in the same transaction).
CREATE TABLE cg_art_auction_proceeds (
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	nft_id          BIGINT NOT NULL,
	round_num       BIGINT NOT NULL,
	amount          DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);

-- ICosmicSignatureArtAuction.sol:AuctionSurplusForwarded (forced/unexpected
-- ETH swept into the game as a donation).
CREATE TABLE cg_art_auction_surplus (
	id              BIGSERIAL PRIMARY KEY,
	evtlog_id       BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num       BIGINT NOT NULL,
	tx_id           BIGINT NOT NULL,
	time_stamp      TIMESTAMPTZ NOT NULL,
	contract_aid    BIGINT NOT NULL,
	amount          DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);

-- +goose Down
DROP TABLE IF EXISTS cg_art_auction_surplus;
DROP TABLE IF EXISTS cg_art_auction_proceeds;
DROP TABLE IF EXISTS cg_art_auction_nft_delivered;
DROP TABLE IF EXISTS cg_art_auction_settled;
DROP TABLE IF EXISTS cg_art_auction_refund_withdrawn;
DROP TABLE IF EXISTS cg_art_auction_refund_credited;
DROP TABLE IF EXISTS cg_art_auction_bid;
DROP TABLE IF EXISTS cg_art_auction_batch;
DROP TABLE IF EXISTS cg_chaos_zero_cycle_funded;
DROP TABLE IF EXISTS cg_art_auction_nfts_minted;
DROP TABLE IF EXISTS cg_adm_art_auction_opening_bid;
DROP TABLE IF EXISTS cg_art_auction_configured;
