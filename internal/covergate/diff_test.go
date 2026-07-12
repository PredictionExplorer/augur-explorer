package covergate

import (
	"errors"
	"strings"
	"testing"
)

const testDiff = `diff --git a/internal/core/core.go b/internal/core/core.go
index 1111111..2222222 100644
--- a/internal/core/core.go
+++ b/internal/core/core.go
@@ -9,2 +9,3 @@
 unchanged
-old
+new
+another
diff --git a/cmd/tool/main.go b/cmd/tool/main.go
index 3333333..4444444 100644
--- a/cmd/tool/main.go
+++ b/cmd/tool/main.go
@@ -5 +5 @@
-old command
+new command
`

func TestParseUnifiedDiff(t *testing.T) {
	t.Parallel()
	changed, err := ParseUnifiedDiff(strings.NewReader(testDiff))
	if err != nil {
		t.Fatalf("ParseUnifiedDiff: %v", err)
	}
	for _, line := range []int{10, 11} {
		if _, ok := changed["internal/core/core.go"][line]; !ok {
			t.Errorf("internal line %d not marked changed: %#v", line, changed)
		}
	}
	if _, ok := changed["cmd/tool/main.go"][5]; !ok {
		t.Errorf("command line not marked changed: %#v", changed)
	}
}

func TestParseUnifiedDiffHandlesNewDeletedAndQuotedFiles(t *testing.T) {
	t.Parallel()
	diff := `diff --git "a/internal/core/new file.go" "b/internal/core/new file.go"
new file mode 100644
--- /dev/null
+++ "b/internal/core/new file.go"
@@ -0,0 +1,2 @@
+package core
+var Added = true
diff --git a/internal/core/old.go b/internal/core/old.go
deleted file mode 100644
--- a/internal/core/old.go
+++ /dev/null
@@ -1 +0,0 @@
-package core
`
	changed, err := ParseUnifiedDiff(strings.NewReader(diff))
	if err != nil {
		t.Fatal(err)
	}
	if len(changed["internal/core/new file.go"]) != 2 {
		t.Fatalf("new-file changes = %#v", changed)
	}
	if _, exists := changed["/dev/null"]; exists {
		t.Fatal("deleted file produced changed lines")
	}
}

func TestAnalyzePatch(t *testing.T) {
	t.Parallel()
	profile, err := ParseProfile(strings.NewReader(testProfile))
	if err != nil {
		t.Fatal(err)
	}
	policy, err := DecodePolicy(strings.NewReader(validPolicyJSON))
	if err != nil {
		t.Fatal(err)
	}
	changed, err := ParseUnifiedDiff(strings.NewReader(testDiff))
	if err != nil {
		t.Fatal(err)
	}
	patch := AnalyzePatch(profile, policy, changed)
	assertMetric(t, patch.Metric, 2, 6)
	if !patch.Applicable || patch.ChangedFiles != 2 || patch.ChangedLines != 3 ||
		len(patch.Uncovered) != 1 ||
		patch.Uncovered[0].File != "cmd/tool/main.go" ||
		patch.Uncovered[0].Statements != 4 {
		t.Fatalf("patch = %+v", patch)
	}
}

func TestAnalyzePatchReturnsNotApplicableForNonExecutableChanges(t *testing.T) {
	t.Parallel()
	profile, err := ParseProfile(strings.NewReader(testProfile))
	if err != nil {
		t.Fatal(err)
	}
	policy, err := DecodePolicy(strings.NewReader(validPolicyJSON))
	if err != nil {
		t.Fatal(err)
	}
	patch := AnalyzePatch(profile, policy, ChangedLines{
		"internal/core/core.go": {1000: {}},
	})
	if patch.Applicable || patch.Metric.Total != 0 {
		t.Fatalf("patch = %+v", patch)
	}
}

func TestParseUnifiedDiffRejectsMalformedInput(t *testing.T) {
	t.Parallel()
	cases := []string{
		"@@ -1 +1 @@\n+line\n",
		"+++ b/internal/core/a.go\n@@ malformed @@\n",
		"+++ b/internal/core/a.go\n@@ -1 +1 @@\n?line\n",
		"+++ b/../../outside.go\n",
		"+++ \"unterminated\n",
		"+++ b/internal/core/a.go\n@@ -999999999999999999999999999999 +1 @@\n+line\n",
		"+++ b/internal/core/a.go\n@@ -1 +999999999999999999999999999999 @@\n+line\n",
	}
	for _, diff := range cases {
		if _, err := ParseUnifiedDiff(strings.NewReader(diff)); err == nil {
			t.Errorf("malformed diff was accepted: %q", diff)
		}
	}
	if _, err := ParseUnifiedDiff(errorReader{}); !errors.Is(err, errTestRead) {
		t.Fatalf("reader error = %v", err)
	}
}

func TestAnalyzePatchIgnoresPathsOutsidePolicyScope(t *testing.T) {
	t.Parallel()
	profile, err := ParseProfile(strings.NewReader(testProfile))
	if err != nil {
		t.Fatal(err)
	}
	policy, err := DecodePolicy(strings.NewReader(validPolicyJSON))
	if err != nil {
		t.Fatal(err)
	}
	patch := AnalyzePatch(profile, policy, ChangedLines{
		"README.md":                   {1: {}},
		"internal/testutil/helper.go": {1: {}},
	})
	if patch.Applicable || patch.ChangedFiles != 0 || patch.ChangedLines != 0 {
		t.Fatalf("patch = %+v", patch)
	}
}
