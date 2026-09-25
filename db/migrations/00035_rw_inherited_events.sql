-- +goose Up
-- RandomWalkNFT (ERC721Enumerable + Ownable) emits three OpenZeppelin events
-- the RandomWalk ETL did not store: Approval, ApprovalForAll and
-- OwnershipTransferred. Every event a deployed contract can emit is
-- indexed (TestEveryABIEventHasHandler in internal/indexer/randomwalk).

CREATE TABLE rw_approval ( -- IERC721.sol:Approval (RandomWalkNFT)
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	owner_aid		BIGINT NOT NULL,
	approved_aid	BIGINT NOT NULL,		-- zero address clears the approval
	token_id		BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE INDEX rw_approval_token_idx ON rw_approval(token_id, block_num DESC);
CREATE TABLE rw_approval_for_all ( -- IERC721.sol:ApprovalForAll (RandomWalkNFT)
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	owner_aid		BIGINT NOT NULL,
	operator_aid	BIGINT NOT NULL,
	approved		BOOLEAN NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE INDEX rw_approval_for_all_owner_idx ON rw_approval_for_all(owner_aid, block_num DESC);
CREATE TABLE rw_ownership_transferred ( -- Ownable.sol:OwnershipTransferred (RandomWalkNFT)
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	prev_owner_aid	BIGINT NOT NULL,
	new_owner_aid	BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);

-- +goose Down
DROP TABLE rw_ownership_transferred;
DROP TABLE rw_approval_for_all;
DROP TABLE rw_approval;
