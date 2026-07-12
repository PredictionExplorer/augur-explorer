package v2

import (
	"errors"
	"fmt"
)

const (
	defaultPageLimit = 50
	maxPageLimit     = 200
)

func resolvePageLimit(requested *int) (int, bool) {
	limit := defaultPageLimit
	if requested != nil {
		limit = *requested
	}
	return limit, limit >= 1 && limit <= maxPageLimit
}

func pageLimitProblemDetail() string {
	return fmt.Sprintf("Limit must be between 1 and %d.", maxPageLimit)
}

func validatePageCardinality(recordCount, limit int) error {
	if recordCount > limit {
		return errors.New("repository returned more rows than requested")
	}
	return nil
}
