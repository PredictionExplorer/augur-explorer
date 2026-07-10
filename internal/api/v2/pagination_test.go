package v2

import "testing"

func TestResolvePageLimit(t *testing.T) {
	t.Parallel()

	value := func(v int) *int { return &v }
	tests := map[string]struct {
		requested *int
		want      int
		valid     bool
	}{
		"default": {want: defaultPageLimit, valid: true},
		"minimum": {requested: value(1), want: 1, valid: true},
		"maximum": {requested: value(maxPageLimit), want: maxPageLimit, valid: true},
		"zero":    {requested: value(0), want: 0},
		"too high": {
			requested: value(maxPageLimit + 1),
			want:      maxPageLimit + 1,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, valid := resolvePageLimit(tc.requested)
			if got != tc.want || valid != tc.valid {
				t.Fatalf("resolvePageLimit = (%d,%v), want (%d,%v)", got, valid, tc.want, tc.valid)
			}
		})
	}
	if got := pageLimitProblemDetail(); got != "Limit must be between 1 and 200." {
		t.Fatalf("problem detail = %q", got)
	}
}
