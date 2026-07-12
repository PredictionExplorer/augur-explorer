package covergate

import (
	"strings"
	"testing"
)

const validPolicyJSON = `{
	"version": 1,
	"legacyInternalFloor": 75.4,
	"internalFloor": 70,
	"productionFloor": 49.4,
	"patchFloor": 95,
	"internalTarget": 90,
	"productionTarget": 90,
	"criticalPackageTarget": 95,
	"commitGateEnabled": false,
	"commitFloor": 90,
	"excludeFiles": ["internal/api/v2/api.gen.go"],
	"excludePrefixes": ["internal/testutil/"]
}`

func TestDecodePolicy(t *testing.T) {
	t.Parallel()
	policy, err := DecodePolicy(strings.NewReader(validPolicyJSON))
	if err != nil {
		t.Fatalf("DecodePolicy: %v", err)
	}
	if policy.Version != 1 || policy.PatchFloor != 95 ||
		policy.CommitGateEnabled || policy.CommitFloor != 90 ||
		!policy.excluded("internal/api/v2/api.gen.go") ||
		!policy.excluded("internal/testutil/golden.go") ||
		policy.excluded("internal/api/v2/server.go") {
		t.Fatalf("policy = %+v", policy)
	}
}

func TestDecodePolicyRejectsInvalidDocuments(t *testing.T) {
	t.Parallel()
	cases := map[string]string{
		"empty":         "",
		"unknown field": strings.Replace(validPolicyJSON, `"version": 1,`, `"version": 1, "other": true,`, 1),
		"trailing":      validPolicyJSON + `{}`,
		"bad trailing":  validPolicyJSON + `x`,
		"version":       strings.Replace(validPolicyJSON, `"version": 1`, `"version": 2`, 1),
		"negative":      strings.Replace(validPolicyJSON, `"patchFloor": 95`, `"patchFloor": -1`, 1),
		"over 100":      strings.Replace(validPolicyJSON, `"patchFloor": 95`, `"patchFloor": 101`, 1),
		"target below":  strings.Replace(validPolicyJSON, `"internalTarget": 90`, `"internalTarget": 60`, 1),
		"production target below": strings.Replace(
			validPolicyJSON, `"productionTarget": 90`, `"productionTarget": 40`, 1,
		),
		"commit floor over 100": strings.Replace(validPolicyJSON, `"commitFloor": 90`, `"commitFloor": 101`, 1),
		"commit gate too early": strings.NewReplacer(
			`"commitGateEnabled": false`, `"commitGateEnabled": true`,
			`"internalFloor": 70`, `"internalFloor": 89`,
		).Replace(validPolicyJSON),
		"absolute path": strings.Replace(validPolicyJSON, `"internal/api/v2/api.gen.go"`, `"/tmp/file.go"`, 1),
		"parent path":   strings.Replace(validPolicyJSON, `"internal/api/v2/api.gen.go"`, `"../file.go"`, 1),
		"empty path":    strings.Replace(validPolicyJSON, `"internal/api/v2/api.gen.go"`, `""`, 1),
		"windows path":  strings.Replace(validPolicyJSON, `"internal/api/v2/api.gen.go"`, `"internal\\file.go"`, 1),
	}
	for name, document := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, err := DecodePolicy(strings.NewReader(document)); err == nil {
				t.Fatal("invalid policy was accepted")
			}
		})
	}
}
