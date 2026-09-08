//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"
)

func chatTestTransaction(t *testing.T, run func(context.Context, *Repo)) {
	t.Helper()
	r := repo(t)
	rollback := errors.New("chat fixture rollback")
	err := r.store.InTx(context.Background(), func(ctx context.Context) error { run(ctx, r); return rollback })
	if !errors.Is(err, rollback) {
		t.Fatalf("chat transaction: %v", err)
	}
}

func TestChatMessagesFilteringAndBoundaries(t *testing.T) {
	chatTestTransaction(t, func(ctx context.Context, r *Repo) {
		if _, err := r.q(ctx).Exec(ctx, `UPDATE cg_bid SET msg='visible',time_stamp=TO_TIMESTAMP(1767226000) WHERE round_num=0`); err != nil {
			t.Fatal(err)
		}
		first, err := r.ChatMessagesPage(ctx, 0, nil, false, 2)
		if err != nil {
			t.Fatal(err)
		}
		if len(first.Records) != 2 || !first.HasMore || first.Records[0].EventLogID != 5010 || first.Records[1].EventLogID != 5008 {
			t.Fatalf("first=%+v", first)
		}
		last := first.Records[1].ChatPosition
		older, err := r.ChatMessagesPage(ctx, 0, &last, false, 2)
		if err != nil {
			t.Fatal(err)
		}
		// Ban identity 2002 is the database row, whose event identity is 5006.
		if len(older.Records) != 1 || older.Records[0].EventLogID != 5004 || older.HasMore {
			t.Fatalf("ban identity or older boundary=%+v", older)
		}
		oldest := older.Records[0].ChatPosition
		newer, err := r.ChatMessagesPage(ctx, 0, &oldest, true, 1)
		if err != nil || len(newer.Records) != 1 || newer.Records[0].EventLogID != 5008 || !newer.HasMore {
			t.Fatalf("newer=%+v,%v", newer, err)
		}
		boundary := newer.Records[0].ChatPosition
		final, err := r.ChatMessagesPage(ctx, 0, &boundary, true, 1)
		if err != nil || len(final.Records) != 1 || final.Records[0].EventLogID != 5010 || final.HasMore {
			t.Fatalf("final=%+v,%v", final, err)
		}
		snapshot, err := r.ChatContext(ctx, 0)
		if err != nil || snapshot.Revision != first.Revision || len(snapshot.Records) != 4 {
			t.Fatalf("context=%+v,%v", snapshot, err)
		}
		if snapshot.Records[0].EventLogID != 5004 || snapshot.Records[1].EventLogID != 5006 {
			t.Fatalf("context must retain all gestures: %+v", snapshot)
		}
		if _, err := r.q(ctx).Exec(ctx, `UPDATE cg_bid SET msg=$1 WHERE evtlog_id=5010`, "\ufeff\n\u3000"); err != nil {
			t.Fatal(err)
		}
		filtered, err := r.ChatMessagesPage(ctx, 0, nil, false, 1)
		if err != nil || len(filtered.Records) != 1 || filtered.Records[0].EventLogID != 5008 || !filtered.HasMore {
			t.Fatalf("blank filtering before limit=%+v,%v", filtered, err)
		}
		if filtered.Revision <= first.Revision {
			t.Fatal("message edit did not invalidate revision")
		}
	})
}

func TestChatMessagesCatchUpBurstAndReorg(t *testing.T) {
	chatTestTransaction(t, func(ctx context.Context, r *Repo) {
		before, err := r.ChatMessagesPage(ctx, 0, nil, false, 50)
		if err != nil {
			t.Fatal(err)
		}
		if err := r.DeleteBid(ctx, 99999999); err != nil {
			t.Fatal(err)
		}
		if _, err := r.q(ctx).Exec(ctx, `INSERT INTO evt_log(id,block_num,tx_id,contract_aid,topic0_sig,log_index,log_rlp)
			SELECT 800000+n,101,1002,2,'12345678',900+n,''::bytea FROM generate_series(1,55) n`); err != nil {
			t.Fatal(err)
		}
		if _, err := r.q(ctx).Exec(ctx, `INSERT INTO cg_bid(id,evtlog_id,block_num,tx_id,time_stamp,contract_aid,bidder_aid,rwalk_nft_id,round_num,bid_type,bid_position,prize_time,eth_price,cst_price,cst_reward,msg)
			SELECT 800000+n,800000+n,101,1002,TO_TIMESTAMP(1767227000+n),2,21,-1,0,0,100+n,TO_TIMESTAMP(1767237000+n),1,-1,0,'burst' FROM generate_series(1,55) n`); err != nil {
			t.Fatal(err)
		}
		page, err := r.ChatMessagesPage(ctx, 0, &before.Latest, true, 50)
		if err != nil || len(page.Records) != 50 || !page.HasMore {
			t.Fatalf("burst page=%+v,%v", page, err)
		}
		if page.Revision != before.Revision {
			t.Fatal("append/no-op delete invalidated history")
		}
		for i, record := range page.Records {
			if record.EventLogID != 800001+int64(i) {
				t.Fatalf("burst skipped row: %+v", record)
			}
		}
		last := page.Records[49].ChatPosition
		remaining, err := r.ChatMessagesPage(ctx, 0, &last, true, 50)
		if err != nil || len(remaining.Records) != 5 || remaining.HasMore || remaining.Records[4].EventLogID != 800055 {
			t.Fatalf("remaining=%+v,%v", remaining, err)
		}
		// Cascading event deletion models the indexer's reorganization path.
		if _, err := r.q(ctx).Exec(ctx, `DELETE FROM evt_log WHERE id=800055`); err != nil {
			t.Fatal(err)
		}
		after, err := r.ChatMessagesPage(ctx, 0, nil, false, 50)
		if err != nil || after.Revision <= page.Revision || after.Latest.EventLogID != 800054 {
			t.Fatalf("reorg=%+v,%v", after, err)
		}
	})
}

func TestChatModerationRevisionAndRollback(t *testing.T) {
	r := repo(t)
	base, err := r.ChatMessagesPage(context.Background(), 0, nil, false, 50)
	if err != nil {
		t.Fatal(err)
	}
	chatTestTransaction(t, func(ctx context.Context, r *Repo) {
		if _, err := r.q(ctx).Exec(ctx, `INSERT INTO cg_banned_bids(bid_id,user_addr,created_at) VALUES(2001,'alice',1)`); err != nil {
			t.Fatal(err)
		}
		banned, err := r.ChatMessagesPage(ctx, 0, nil, false, 50)
		if err != nil || banned.Revision <= base.Revision || len(banned.Records) != len(base.Records)-1 {
			t.Fatalf("ban=%+v,%v", banned, err)
		}
		if _, err := r.q(ctx).Exec(ctx, `DELETE FROM cg_banned_bids WHERE bid_id=2001`); err != nil {
			t.Fatal(err)
		}
		unbanned, err := r.ChatMessagesPage(ctx, 0, nil, false, 50)
		if err != nil || unbanned.Revision <= banned.Revision || len(unbanned.Records) != len(base.Records) {
			t.Fatalf("unban=%+v,%v", unbanned, err)
		}
	})
	after, err := r.ChatMessagesPage(context.Background(), 0, nil, false, 50)
	if err != nil || after.Revision != base.Revision {
		t.Fatalf("revision escaped rollback: %+v,%v", after, err)
	}
}

func TestChatEmptyWhitespaceAndErrors(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	empty, err := r.ChatMessagesPage(ctx, 99999, nil, false, 50)
	if err != nil || empty.Records == nil || len(empty.Records) != 0 || empty.Revision < 1 || empty.Latest.EventLogID != 0 || !empty.Latest.OccurredAt.Equal(time.Unix(0, 0)) {
		t.Fatalf("empty=%+v,%v", empty, err)
	}
	snapshot, err := r.ChatContext(ctx, 99999)
	if err != nil || snapshot.Records == nil || len(snapshot.Records) != 0 || snapshot.Revision != empty.Revision {
		t.Fatalf("empty context=%+v,%v", snapshot, err)
	}
	for _, value := range []string{"", "\t\n\r\v\f ", "\u00a0\u1680\u2000\u2001\u2002\u2003\u2004\u2005\u2006\u2007\u2008\u2009\u200a\u2028\u2029\u202f\u205f\u3000\ufeff"} {
		var visible bool
		if err := r.q(ctx).QueryRow(ctx, `SELECT cg_chat_has_text($1)`, value).Scan(&visible); err != nil {
			t.Fatal(err)
		}
		if visible {
			t.Errorf("whitespace treated as text: %q", value)
		}
	}
	for _, value := range []string{"hello", " \nhello\ufeff", "\u0085"} {
		var visible bool
		if err := r.q(ctx).QueryRow(ctx, `SELECT cg_chat_has_text($1)`, value).Scan(&visible); err != nil || !visible {
			t.Fatalf("text=%q,%v,%v", value, visible, err)
		}
	}
	if _, err := r.ChatMessagesPage(ctx, -1, nil, false, 50); err == nil {
		t.Fatal("negative round accepted")
	}
	if _, err := r.ChatMessagesPage(ctx, 0, nil, false, 201); err == nil {
		t.Fatal("oversized page accepted")
	}
	if _, err := r.ChatMessagesPage(ctx, 0, nil, true, 50); err == nil {
		t.Fatal("missing after accepted")
	}
	if _, err := r.ChatMessagesPage(ctx, 0, &ChatPosition{}, false, 50); err == nil {
		t.Fatal("invalid cursor accepted")
	}
	if _, err := r.ChatContext(ctx, -1); err == nil {
		t.Fatal("negative context round accepted")
	}
	cancelled, cancel := context.WithCancel(ctx)
	cancel()
	if _, err := r.ChatMessagesPage(cancelled, 0, nil, false, 50); !errors.Is(err, context.Canceled) {
		t.Fatalf("message cancellation=%v", err)
	}
	if _, err := r.ChatContext(cancelled, 0); !errors.Is(err, context.Canceled) {
		t.Fatalf("context cancellation=%v", err)
	}
}

// Revision state is mandatory even for an empty cycle. A missing migration or
// damaged row must never look like a valid empty history to the client.
func TestChatRevisionStateFailures(t *testing.T) {
	for _, tc := range []struct {
		name  string
		setup []string
		want  string
	}{
		{
			name:  "missing singleton",
			setup: []string{`DELETE FROM cg_chat_revision`},
			want:  "revision state missing",
		},
		{
			name: "unreadable revision",
			// Model schema drift inside a transaction; both the schema and the
			// row are restored by the same rollback as every other chat fixture.
			setup: []string{
				`ALTER TABLE cg_chat_revision ALTER COLUMN revision DROP NOT NULL`,
				`UPDATE cg_chat_revision SET revision=NULL`,
			},
			want: "can't scan",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			chatTestTransaction(t, func(ctx context.Context, r *Repo) {
				for _, statement := range tc.setup {
					if _, err := r.q(ctx).Exec(ctx, statement); err != nil {
						t.Fatal(err)
					}
				}
				if _, err := r.ChatMessagesPage(ctx, 99999, nil, false, 50); err == nil || !strings.Contains(err.Error(), tc.want) || !strings.Contains(err.Error(), "chat message page") {
					t.Fatalf("message revision error = %v, want %q with operation context", err, tc.want)
				}
				if _, err := r.ChatContext(ctx, 99999); err == nil || !strings.Contains(err.Error(), tc.want) || !strings.Contains(err.Error(), "chat context") {
					t.Fatalf("context revision error = %v, want %q with operation context", err, tc.want)
				}
			})
			// Verify that a failure cannot leave shared revision state damaged.
			if page, err := repo(t).ChatMessagesPage(context.Background(), 99999, nil, false, 50); err != nil || page.Revision < 1 {
				t.Fatalf("revision state after rollback = %+v, %v", page, err)
			}
		})
	}
}
