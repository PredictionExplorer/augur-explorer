package store

import (
	"fmt"
	"sync"
	"testing"
)

func TestAddressCachePutGet(t *testing.T) {
	c := newAddressCache(4)
	if _, ok := c.get("0xa"); ok {
		t.Fatal("empty cache reported a hit")
	}
	c.put("0xa", 1)
	c.put("0xb", 2)
	if aid, ok := c.get("0xa"); !ok || aid != 1 {
		t.Fatalf("get(0xa) = %d, %v; want 1, true", aid, ok)
	}
	if aid, ok := c.get("0xb"); !ok || aid != 2 {
		t.Fatalf("get(0xb) = %d, %v; want 2, true", aid, ok)
	}
	if got := c.len(); got != 2 {
		t.Fatalf("len = %d, want 2", got)
	}
}

func TestAddressCachePutExistingKeepsID(t *testing.T) {
	// address→id is immutable; a second put must not overwrite (it can only
	// refresh recency).
	c := newAddressCache(4)
	c.put("0xa", 1)
	c.put("0xa", 99)
	if aid, _ := c.get("0xa"); aid != 1 {
		t.Fatalf("get after duplicate put = %d, want original 1", aid)
	}
	if got := c.len(); got != 1 {
		t.Fatalf("len after duplicate put = %d, want 1", got)
	}
}

func TestAddressCacheEvictsLeastRecentlyUsed(t *testing.T) {
	c := newAddressCache(3)
	c.put("0xa", 1)
	c.put("0xb", 2)
	c.put("0xc", 3)

	// Touch 0xa so 0xb becomes the least recently used entry.
	if _, ok := c.get("0xa"); !ok {
		t.Fatal("0xa missing before eviction")
	}
	c.put("0xd", 4)

	if _, ok := c.get("0xb"); ok {
		t.Error("0xb should have been evicted (least recently used)")
	}
	for addr, want := range map[string]int64{"0xa": 1, "0xc": 3, "0xd": 4} {
		if aid, ok := c.get(addr); !ok || aid != want {
			t.Errorf("get(%s) = %d, %v; want %d, true", addr, aid, ok, want)
		}
	}
	if got := c.len(); got != 3 {
		t.Fatalf("len = %d, want bound 3", got)
	}
}

func TestAddressCacheBoundHolds(t *testing.T) {
	const bound = 16
	c := newAddressCache(bound)
	for i := 0; i < 10*bound; i++ {
		c.put(fmt.Sprintf("0x%040x", i), int64(i))
	}
	if got := c.len(); got != bound {
		t.Fatalf("len after overfill = %d, want %d", got, bound)
	}
	// The most recent `bound` entries survive.
	for i := 9*bound + 1; i < 10*bound; i++ {
		if _, ok := c.get(fmt.Sprintf("0x%040x", i)); !ok {
			t.Errorf("recent entry %d evicted", i)
		}
	}
}

func TestAddressCacheReset(t *testing.T) {
	c := newAddressCache(4)
	c.put("0xa", 1)
	c.reset()
	if got := c.len(); got != 0 {
		t.Fatalf("len after reset = %d, want 0", got)
	}
	if _, ok := c.get("0xa"); ok {
		t.Fatal("entry survived reset")
	}
	// The cache stays usable after reset.
	c.put("0xb", 2)
	if aid, ok := c.get("0xb"); !ok || aid != 2 {
		t.Fatalf("get after reset+put = %d, %v; want 2, true", aid, ok)
	}
}

func TestAddressCacheZeroSizeUsesDefault(t *testing.T) {
	c := newAddressCache(0)
	if c.max != DefaultAddressCacheSize {
		t.Fatalf("max = %d, want default %d", c.max, DefaultAddressCacheSize)
	}
}

// TestAddressCacheConcurrent exercises the cache from many goroutines; run
// with -race it proves the locking is sound.
func TestAddressCacheConcurrent(t *testing.T) {
	c := newAddressCache(64)
	var wg sync.WaitGroup
	for g := 0; g < 8; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			for i := 0; i < 500; i++ {
				addr := fmt.Sprintf("0x%040x", i%100)
				if aid, ok := c.get(addr); ok && aid != int64(i%100) {
					t.Errorf("corrupted entry: get(%s) = %d", addr, aid)
					return
				}
				c.put(addr, int64(i%100))
				if g == 0 && i%97 == 0 {
					c.reset()
				}
			}
		}(g)
	}
	wg.Wait()
	if c.len() > 64 {
		t.Fatalf("len %d exceeds bound", c.len())
	}
}
