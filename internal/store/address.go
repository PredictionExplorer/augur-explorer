// Address-id resolution on the pgx-native Store, with a bounded per-Store
// LRU cache. This is the Phase 1 replacement for the package-level address
// cache in lookups.go: the mapping address→id is immutable and append-only,
// so cached entries can never go stale — the LRU bound exists purely to cap
// memory on long-running processes.

package store

import (
	"container/list"
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
)

// DefaultAddressCacheSize bounds the per-Store address-id cache. At ~64
// bytes per entry the default costs a few megabytes and covers far more
// distinct addresses than one ETL batch or API burst touches.
const DefaultAddressCacheSize = 65536

// addressCache is a thread-safe bounded LRU map from address string to
// address_id.
type addressCache struct {
	mu    sync.Mutex
	max   int
	order *list.List // front = most recently used
	items map[string]*list.Element
}

type addressCacheEntry struct {
	addr string
	aid  int64
}

func newAddressCache(maxEntries int) *addressCache {
	if maxEntries <= 0 {
		maxEntries = DefaultAddressCacheSize
	}
	return &addressCache{
		max:   maxEntries,
		order: list.New(),
		items: make(map[string]*list.Element, maxEntries),
	}
}

func (c *addressCache) get(addr string) (int64, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	el, ok := c.items[addr]
	if !ok {
		return 0, false
	}
	c.order.MoveToFront(el)
	return el.Value.(*addressCacheEntry).aid, true
}

func (c *addressCache) put(addr string, aid int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.items[addr]; ok {
		// address→id never changes; refresh recency only.
		c.order.MoveToFront(el)
		return
	}
	c.items[addr] = c.order.PushFront(&addressCacheEntry{addr: addr, aid: aid})
	if c.order.Len() > c.max {
		oldest := c.order.Back()
		c.order.Remove(oldest)
		delete(c.items, oldest.Value.(*addressCacheEntry).addr)
	}
}

// putAll flushes a transaction overlay into the cache (InTx, after commit).
func (c *addressCache) putAll(entries map[string]int64) {
	for addr, aid := range entries {
		c.put(addr, aid)
	}
}

func (c *addressCache) len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.order.Len()
}

func (c *addressCache) reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.order.Init()
	c.items = make(map[string]*list.Element, c.max)
}

// cachedAddressID consults the shared LRU and, inside a transaction begun
// by this Store, the transaction's overlay of not-yet-committed creations.
func (s *Store) cachedAddressID(ctx context.Context, addr string) (int64, bool) {
	if aid, ok := s.addrCache.get(addr); ok {
		return aid, true
	}
	if st := s.txState(ctx); st != nil {
		aid, ok := st.overlay[addr]
		return aid, ok
	}
	return 0, false
}

// cacheAddressID records a resolved id. Inside a transaction the entry goes
// to the transaction overlay only — a rollback discards it, a commit
// flushes it into the shared LRU (InTx) — so the shared cache can never
// hold an id whose address row was rolled away.
func (s *Store) cacheAddressID(ctx context.Context, addr string, aid int64) {
	if st := s.txState(ctx); st != nil {
		if st.overlay == nil {
			st.overlay = make(map[string]int64, 8)
		}
		st.overlay[addr] = aid
		return
	}
	s.addrCache.put(addr, aid)
}

// LookupOrCreateAddress returns the address_id for addr, inserting a new
// address row (recorded against blockNum/txID) when it doesn't exist yet.
// Results are memoized in the Store's bounded LRU cache (transaction-safe:
// see cacheAddressID).
func (s *Store) LookupOrCreateAddress(ctx context.Context, addr string, blockNum, txID int64) (int64, error) {
	if aid, ok := s.cachedAddressID(ctx, addr); ok {
		return aid, nil
	}
	if len(addr) == 0 {
		return 0, fmt.Errorf("lookup/create address: empty address (block %d, tx %d)", blockNum, txID)
	}
	q := s.q(ctx)
	var aid int64
	err := q.QueryRow(ctx, "SELECT address_id FROM address WHERE addr=$1", addr).Scan(&aid)
	if errors.Is(err, pgx.ErrNoRows) {
		// ON CONFLICT DO NOTHING keeps the create race-safe in both modes:
		// a plain INSERT losing the race raises a unique violation, which
		// inside a transaction would poison every later statement (25P02);
		// a conflict here simply returns no row and the re-read below picks
		// up the winner. The existing row keeps its first-seen block.
		err = q.QueryRow(ctx,
			"INSERT INTO address(addr,block_num,tx_id) VALUES($1,$2,$3) ON CONFLICT (addr) DO NOTHING RETURNING address_id",
			addr, blockNum, txID).Scan(&aid)
		if errors.Is(err, pgx.ErrNoRows) {
			// A concurrent writer created the row between our SELECT and
			// INSERT; re-read so both callers agree on the id.
			err = q.QueryRow(ctx, "SELECT address_id FROM address WHERE addr=$1", addr).Scan(&aid)
		}
	}
	if err != nil {
		return 0, WrapError("lookup/create address "+addr, err)
	}
	s.cacheAddressID(ctx, addr, aid)
	return aid, nil
}

// LookupAddressID returns the address_id for addr; a missing address yields
// a wrapped ErrNotFound. Hits populate the same cache as
// LookupOrCreateAddress.
func (s *Store) LookupAddressID(ctx context.Context, addr string) (int64, error) {
	if aid, ok := s.cachedAddressID(ctx, addr); ok {
		return aid, nil
	}
	var aid int64
	err := s.q(ctx).QueryRow(ctx, "SELECT address_id FROM address WHERE addr=$1", addr).Scan(&aid)
	if err != nil {
		return 0, WrapError("address id lookup for "+addr, err)
	}
	s.cacheAddressID(ctx, addr, aid)
	return aid, nil
}

// AddressByID returns the address string for an address_id; a missing id
// yields a wrapped ErrNotFound.
func (s *Store) AddressByID(ctx context.Context, aid int64) (string, error) {
	var addr string
	err := s.q(ctx).QueryRow(ctx, "SELECT addr FROM address WHERE address_id=$1", aid).Scan(&addr)
	if err != nil {
		return "", WrapError(fmt.Sprintf("address lookup for id %d", aid), err)
	}
	return addr, nil
}

// ResetAddressCache clears the address-id cache. Test harnesses that
// truncate and re-seed the address table between cases must call it,
// otherwise ids cached from a previous seeding would leak into the next one.
func (s *Store) ResetAddressCache() { s.addrCache.reset() }
