package store

import (
	"database/sql"
	"fmt"
	"time"
)

// TimeText returns a scan target that stores a timestamp column into a
// string field formatted with time.RFC3339Nano. That is byte-identical to
// what the legacy database/sql layer produced when scanning timestamptz into
// string fields (database/sql formats time.Time with RFC3339Nano during
// convertAssign), which is the format all golden files pin. Scanning SQL
// NULL fails, matching database/sql's "converting NULL to string" error;
// queries use COALESCE where NULL is legitimate.
func TimeText(dst *string) sql.Scanner {
	return timeTextScanner{dst: dst}
}

type timeTextScanner struct{ dst *string }

func (s timeTextScanner) Scan(src any) error {
	t, ok := src.(time.Time)
	if !ok {
		return fmt.Errorf("cannot scan %T into timestamp text", src)
	}
	*s.dst = t.Format(time.RFC3339Nano)
	return nil
}
