package cosmicgame

import (
	"context"
	"fmt"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// ChatPosition is a deterministic message boundary; the zero event ID and
// Unix epoch identify an empty feed's synchronization boundary.
type ChatPosition struct {
	OccurredAt time.Time
	EventLogID int64
}

// ChatMessageRecord deliberately excludes rewards and attached assets.
type ChatMessageRecord struct {
	ChatPosition

	Round             int64
	Position          int64
	BidderAddress     string
	Message           string
	BidType           int64
	EthPriceWei       string
	CstPriceWei       string
	RandomWalkTokenID int64
	TransactionHash   string
}

// ChatMessagePage couples bounded messages to their consistent history revision.
type ChatMessagePage struct {
	Records  []ChatMessageRecord
	Revision int64
	Latest   ChatPosition
	HasMore  bool
}

// ChatContextRecord is complete chronological metadata for reconstructing
// cycle moments and participant counts. It never includes message bodies.
type ChatContextRecord struct {
	ChatPosition

	Round                          int64
	Position                       int64
	BidderAddress                  string
	BidType                        int64
	PrizeAt                        time.Time
	CstDutchAuctionDurationSeconds int64
}

// ChatContextSnapshot holds complete minimal metadata and its history revision.
type ChatContextSnapshot struct {
	Records  []ChatContextRecord
	Revision int64
}

const visibleChatMessage = `b.round_num=$1 AND cg_chat_has_text(b.msg)
	AND NOT EXISTS (SELECT 1 FROM cg_banned_bids ban WHERE ban.bid_id=b.id)`

// ChatMessagesPage reads the page, latest visible boundary, and invalidation
// revision in a single MVCC snapshot. after selects ascending catch-up pages;
// otherwise the cursor selects older rows in newest-first order.
func (r *Repo) ChatMessagesPage(ctx context.Context, round int64, cursor *ChatPosition, after bool, limit int) (ChatMessagePage, error) {
	const op = "chat message page"
	out := ChatMessagePage{Records: make([]ChatMessageRecord, 0)}
	if round < 0 || limit < 1 || limit > 200 || (after && cursor == nil) {
		return out, fmt.Errorf("%s: invalid parameters", op)
	}
	if cursor != nil && (cursor.EventLogID < 0 || cursor.OccurredAt.Before(time.Unix(0, 0)) ||
		(cursor.EventLogID == 0 && (!after || !cursor.OccurredAt.Equal(time.Unix(0, 0))))) {
		return out, fmt.Errorf("%s: invalid cursor", op)
	}
	order, comparison := "DESC", "<"
	if after {
		order, comparison = "ASC", ">"
	}
	where, paging := visibleChatMessage, "LIMIT $2"
	args := []any{round, limit + 1}
	if cursor != nil {
		where += " AND (b.time_stamp,b.evtlog_id)" + comparison + "($2,$3)"
		paging = "LIMIT $4"
		args = []any{round, cursor.OccurredAt, cursor.EventLogID, limit + 1}
	}
	// Every interpolated SQL clause above is selected from constant literals.
	query := `WITH latest AS (
		SELECT b.time_stamp,b.evtlog_id FROM cg_bid b WHERE ` + visibleChatMessage + `
		ORDER BY b.time_stamp DESC,b.evtlog_id DESC LIMIT 1
	)
	SELECT s.revision,COALESCE(l.time_stamp,'epoch'::timestamptz),COALESCE(l.evtlog_id,0),
		COALESCE(p.evtlog_id,0),COALESCE(p.round_num,0),COALESCE(p.bid_position,0),
		COALESCE(p.addr,''),COALESCE(p.time_stamp,'epoch'::timestamptz),COALESCE(p.msg,''),
		COALESCE(p.bid_type,0),COALESCE(p.eth_price,''),COALESCE(p.cst_price,''),
		COALESCE(p.rwalk_nft_id,-1),COALESCE(p.tx_hash,'')
	FROM cg_chat_revision s LEFT JOIN latest l ON TRUE
	LEFT JOIN LATERAL (
		SELECT b.evtlog_id,b.round_num,b.bid_position,ba.addr,b.time_stamp,b.msg,b.bid_type,
			b.eth_price::text,b.cst_price::text,b.rwalk_nft_id,t.tx_hash
		FROM cg_bid b LEFT JOIN address ba ON ba.address_id=b.bidder_aid
		LEFT JOIN transaction t ON t.id=b.tx_id
		WHERE ` + where + ` ORDER BY b.time_stamp ` + order + `,b.evtlog_id ` + order + ` ` + paging + `
	) p ON TRUE WHERE s.singleton ORDER BY p.time_stamp ` + order + `,p.evtlog_id ` + order
	rows, err := r.q(ctx).Query(ctx, query, args...)
	if err != nil {
		return out, store.WrapError(op, err)
	}
	defer rows.Close()
	for rows.Next() {
		var record ChatMessageRecord
		if err := rows.Scan(&out.Revision, &out.Latest.OccurredAt, &out.Latest.EventLogID,
			&record.EventLogID, &record.Round, &record.Position, &record.BidderAddress, &record.OccurredAt,
			&record.Message, &record.BidType, &record.EthPriceWei, &record.CstPriceWei, &record.RandomWalkTokenID, &record.TransactionHash); err != nil {
			return out, store.WrapError(op, err)
		}
		if record.EventLogID != 0 {
			out.Records = append(out.Records, record)
		}
	}
	if err := rows.Err(); err != nil {
		return out, store.WrapError(op, err)
	}
	if out.Revision < 1 {
		return out, fmt.Errorf("%s: revision state missing", op)
	}
	if len(out.Records) > limit {
		out.Records = out.Records[:limit]
		out.HasMore = true
	}
	return out, nil
}

// ChatContext returns all minimal cycle metadata. Its O(cycle-size) cost is
// explicit: this compatibility snapshot preserves client-derived milestones
// until an authoritative aggregate/event API replaces that reconstruction.
func (r *Repo) ChatContext(ctx context.Context, round int64) (ChatContextSnapshot, error) {
	const op = "chat context"
	out := ChatContextSnapshot{Records: make([]ChatContextRecord, 0)}
	if round < 0 {
		return out, fmt.Errorf("%s: invalid round", op)
	}
	const query = `SELECT s.revision,COALESCE(p.evtlog_id,0),COALESCE(p.round_num,0),
		COALESCE(p.bid_position,0),COALESCE(p.addr,''),COALESCE(p.time_stamp,'epoch'::timestamptz),
		COALESCE(p.bid_type,0),COALESCE(p.prize_time,'epoch'::timestamptz),COALESCE(p.duration,-1)
	FROM cg_chat_revision s LEFT JOIN LATERAL (
		SELECT b.evtlog_id,b.round_num,b.bid_position,ba.addr,b.time_stamp,b.bid_type,b.prize_time,
			CASE WHEN b.cst_dutch_auction_duration>=0 THEN b.cst_dutch_auction_duration::bigint ELSE -1 END AS duration
		FROM cg_bid b LEFT JOIN address ba ON ba.address_id=b.bidder_aid
		WHERE b.round_num=$1
	) p ON TRUE WHERE s.singleton ORDER BY p.time_stamp ASC,p.evtlog_id ASC`
	rows, err := r.q(ctx).Query(ctx, query, round)
	if err != nil {
		return out, store.WrapError(op, err)
	}
	defer rows.Close()
	for rows.Next() {
		var record ChatContextRecord
		if err := rows.Scan(&out.Revision, &record.EventLogID, &record.Round, &record.Position,
			&record.BidderAddress, &record.OccurredAt, &record.BidType, &record.PrizeAt, &record.CstDutchAuctionDurationSeconds); err != nil {
			return out, store.WrapError(op, err)
		}
		if record.EventLogID != 0 {
			out.Records = append(out.Records, record)
		}
	}
	if err := rows.Err(); err != nil {
		return out, store.WrapError(op, err)
	}
	if out.Revision < 1 {
		return out, fmt.Errorf("%s: revision state missing", op)
	}
	return out, nil
}
