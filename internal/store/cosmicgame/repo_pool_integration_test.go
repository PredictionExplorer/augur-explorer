//go:build integration

package cosmicgame

import "github.com/jackc/pgx/v5/pgxpool"

// pool exposes the raw connection pool for the integration tests' fixture
// mutation and cleanup SQL, which must run outside any transaction. It lives
// in a test file because no production code needs the raw pool.
func (r *Repo) pool() *pgxpool.Pool { return r.store.Pool() }
