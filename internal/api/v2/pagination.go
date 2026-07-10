package v2

import "fmt"

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
