package randomwalk

import "testing"

func TestActiveOffersOrderClause(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{0, " ORDER BY o.id"},
		{1, " ORDER BY o.price DESC"},
		{2, " ORDER BY o.price ASC"},
		{3, " ORDER BY o.id"},
		{-1, " ORDER BY o.id"},
	}
	for _, tc := range cases {
		if got := activeOffersOrderClause(tc.in); got != tc.want {
			t.Errorf("activeOffersOrderClause(%d) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestMonthName(t *testing.T) {
	cases := []struct {
		in   int64
		want string
	}{
		{1, "January"}, {2, "February"}, {3, "March"}, {4, "April"},
		{5, "May"}, {6, "June"}, {7, "July"}, {8, "August"},
		{9, "September"}, {10, "October"}, {11, "November"}, {12, "December"},
		{0, "???"}, {13, "???"}, {-1, "???"},
	}
	for _, tc := range cases {
		if got := monthName(tc.in); got != tc.want {
			t.Errorf("monthName(%d) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func FuzzActiveOffersOrderClause(f *testing.F) {
	for _, seed := range []int{0, 1, 2, 3, -1, 1 << 30} {
		f.Add(seed)
	}
	whitelist := map[string]bool{
		" ORDER BY o.id":         true,
		" ORDER BY o.price DESC": true,
		" ORDER BY o.price ASC":  true,
	}
	f.Fuzz(func(t *testing.T, orderBy int) {
		got := activeOffersOrderClause(orderBy)
		if !whitelist[got] {
			t.Fatalf("activeOffersOrderClause(%d) = %q: not in the whitelist", orderBy, got)
		}
	})
}
