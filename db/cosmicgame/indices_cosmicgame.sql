CREATE INDEX bid_bidder_aid_idx			ON cg_bid			(bidder_aid);
CREATE INDEX prize_winner_aid_idx		ON cg_prize_claim	(winner_aid);
CREATE INDEX prize_num_idx				ON cg_prize_claim	(round_num);
CREATE INDEX mint_tokid_idx				ON cg_mint_event	(token_id);
