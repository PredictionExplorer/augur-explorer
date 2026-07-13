package randomwalk

import (
	"encoding/json"
	"math"
)

// JSONNullFloat64 is an optional float for API JSON: encodes as null when unset or non-finite
// (encoding/json rejects NaN and ±Inf in plain float64).
type JSONNullFloat64 struct {
	Valid bool
	Value float64
}

func (j JSONNullFloat64) MarshalJSON() ([]byte, error) {
	if !j.Valid {
		return []byte("null"), nil
	}
	x := j.Value
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return []byte("null"), nil
	}
	return json.Marshal(x)
}
