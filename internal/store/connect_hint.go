package store

import (
	"fmt"
	"os"
	"strings"
)

// ConnectHint returns operator guidance for a failed New call: how to fix a
// password mismatch and the status of the DATABASE_URL / PGSQL_* environment
// variables. The binaries print it (to stderr or their log) next to the
// connect error so a misconfigured host is diagnosable without reading
// source. Returns "" when err is nil. Secrets never appear: the password
// reports only its length and DATABASE_URL (which embeds the password) only
// whether it is set.
func ConnectHint(err error) string {
	if err == nil {
		return ""
	}
	var b strings.Builder
	if strings.Contains(err.Error(), "password authentication failed") {
		b.WriteString("Hint: PGSQL_PASSWORD does not match the password set in PostgreSQL for this user.\n")
		b.WriteString("  Option A – Make PostgreSQL use your env password:\n")
		b.WriteString("    psql -U postgres -c \"ALTER USER cosmicgame PASSWORD 'YOUR_ENV_PASSWORD';\"\n")
		b.WriteString("  Option B – Discover the correct password and set it in your env:\n")
		b.WriteString("    psql -h 127.0.0.1 -U cosmicgame -d cosmicgame -W   # type the working password, then set that in PGSQL_PASSWORD\n")
		b.WriteString("  Option C – Use Unix socket (no password) if pg_hba allows trust/peer for local:\n")
		b.WriteString("    unset PGSQL_HOST\n")
	}
	b.WriteString("Environment variable status:\n")
	if v := os.Getenv("DATABASE_URL"); v == "" {
		b.WriteString("  DATABASE_URL: not set (or empty)\n")
	} else {
		b.WriteString("  DATABASE_URL: set (wins over PGSQL_*; never echoed — it embeds the password)\n")
	}
	for _, name := range []string{"PGSQL_USERNAME", "PGSQL_PASSWORD", "PGSQL_DATABASE", "PGSQL_HOST"} {
		v := os.Getenv(name)
		switch {
		case v == "":
			fmt.Fprintf(&b, "  %s: not set (or empty)\n", name)
		case name == "PGSQL_PASSWORD":
			fmt.Fprintf(&b, "  %s: set (length %d)\n", name, len(v))
		default:
			fmt.Fprintf(&b, "  %s: set = %q\n", name, v)
		}
	}
	return b.String()
}
