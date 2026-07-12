package covergate

import (
	"errors"
	"fmt"
	"io"
	"math"
	"slices"
	"strings"
	"testing"
)

const testProfile = `mode: atomic
github.com/example/project/internal/core/core.go:10.1,12.2 2 0
github.com/example/project/internal/core/core.go:10.1,12.2 2 7
github.com/example/project/internal/core/core.go:20.1,20.9 1 0
github.com/example/project/internal/api/v2/api.gen.go:1.1,3.2 3 9
github.com/example/project/internal/testutil/helper.go:1.1,2.2 2 9
github.com/example/project/cmd/tool/main.go:5.1,9.2 4 0
`

func TestParseProfileDeduplicatesAndNormalizes(t *testing.T) {
	t.Parallel()
	profile, err := ParseProfile(strings.NewReader(testProfile))
	if err != nil {
		t.Fatalf("ParseProfile: %v", err)
	}
	if profile.Mode != "atomic" || len(profile.Blocks) != 5 {
		t.Fatalf("profile = %+v", profile)
	}
	firstCore := findBlock(t, profile, "internal/core/core.go", 10)
	if firstCore.Count != 7 || firstCore.Statements != 2 {
		t.Fatalf("deduplicated block = %+v", firstCore)
	}
}

func TestAnalyzeCoverageScopes(t *testing.T) {
	t.Parallel()
	profile, err := ParseProfile(strings.NewReader(testProfile))
	if err != nil {
		t.Fatal(err)
	}
	policy, err := DecodePolicy(strings.NewReader(validPolicyJSON))
	if err != nil {
		t.Fatal(err)
	}
	analysis := Analyze(profile, policy)
	assertMetric(t, analysis.LegacyInternal, 7, 8)
	assertMetric(t, analysis.Internal, 2, 3)
	assertMetric(t, analysis.Production, 2, 7)
	assertMetric(t, analysis.Packages["internal/core"], 2, 3)
	assertMetric(t, analysis.Packages["cmd/tool"], 0, 4)
	if _, ok := analysis.Packages["internal/testutil"]; ok {
		t.Fatal("excluded test utility package was reported")
	}
}

func TestParseProfileRejectsMalformedInput(t *testing.T) {
	t.Parallel()
	cases := map[string]string{
		"empty":               "",
		"missing mode":        "internal/a.go:1.1,1.2 1 1\n",
		"bad mode":            "mode: other\ninternal/a.go:1.1,1.2 1 1\n",
		"no blocks":           "mode: set\n",
		"bad line":            "mode: set\nnot-a-block\n",
		"zero coordinate":     "mode: set\ninternal/a.go:0.1,1.2 1 1\n",
		"outside scope":       "mode: set\ngithub.com/example/other/a.go:1.1,1.2 1 1\n",
		"negative statements": "mode: set\ninternal/a.go:1.1,1.2 -1 1\n",
		"statement overflow":  "mode: set\ninternal/a.go:1.1,1.2 999999999999999999999999999999 1\n",
		"count overflow":      "mode: set\ninternal/a.go:1.1,1.2 1 999999999999999999999999999999\n",
		"negative ordering":   "mode: set\ninternal/a.go:2.1,1.2 1 1\n",
		"unsafe clean path":   "mode: set\ninternal/../../a.go:1.1,1.2 1 1\n",
		"duplicate mismatch": "mode: set\n" +
			"internal/a.go:1.1,1.2 1 0\ninternal/a.go:1.1,1.2 2 1\n",
	}
	for name, profile := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, err := ParseProfile(strings.NewReader(profile)); err == nil {
				t.Fatal("malformed profile was accepted")
			}
		})
	}
}

func TestProfileSortsEverySourceCoordinate(t *testing.T) {
	t.Parallel()
	profile, err := ParseProfile(strings.NewReader(`mode: count

internal/b.go:1.1,1.2 1 0
internal/a.go:1.2,1.3 1 0
internal/a.go:1.1,2.2 1 0
internal/a.go:1.1,2.1 1 0
internal/a.go:1.1,1.3 1 0
`))
	if err != nil {
		t.Fatal(err)
	}
	got := make([]string, len(profile.Blocks))
	for index, block := range profile.Blocks {
		got[index] = strings.Join([]string{
			block.File,
			fmt.Sprintf("%d.%d", block.StartLine, block.StartColumn),
			fmt.Sprintf("%d.%d", block.EndLine, block.EndColumn),
		}, ":")
	}
	want := []string{
		"internal/a.go:1.1:1.3",
		"internal/a.go:1.1:2.1",
		"internal/a.go:1.1:2.2",
		"internal/a.go:1.2:1.3",
		"internal/b.go:1.1:1.2",
	}
	if !slices.Equal(got, want) {
		t.Fatalf("sorted blocks = %v, want %v", got, want)
	}
}

func TestParseProfilePropagatesReaderErrors(t *testing.T) {
	t.Parallel()
	if _, err := ParseProfile(errorReader{}); !errors.Is(err, errTestRead) {
		t.Fatalf("initial reader error = %v", err)
	}
	reader := &dataThenErrorReader{
		data: []byte("mode: atomic\ninternal/a.go:1.1,1.2 1 1\n"),
	}
	if _, err := ParseProfile(reader); !errors.Is(err, errTestRead) {
		t.Fatalf("trailing reader error = %v", err)
	}
}

func TestNormalizeProfilePath(t *testing.T) {
	t.Parallel()
	cases := map[string]string{
		"internal/a.go":                              "internal/a.go",
		"./cmd/tool/main.go":                         "cmd/tool/main.go",
		"github.com/example/project/internal/a.go":   "internal/a.go",
		`C:\repo\project\internal\api\handler.go`:    "internal/api/handler.go",
		"/Users/example/project/cmd/tool/command.go": "cmd/tool/command.go",
	}
	for input, want := range cases {
		got, err := normalizeProfilePath(input)
		if err != nil || got != want {
			t.Errorf("normalizeProfilePath(%q) = %q, %v; want %q", input, got, err, want)
		}
	}
}

func findBlock(t *testing.T, profile Profile, file string, startLine int) Block {
	t.Helper()
	for _, block := range profile.Blocks {
		if block.File == file && block.StartLine == startLine {
			return block
		}
	}
	t.Fatalf("block %s:%d not found", file, startLine)
	return Block{}
}

func assertMetric(t *testing.T, metric Metric, covered, total int) {
	t.Helper()
	if metric.Covered != covered || metric.Total != total {
		t.Fatalf("metric = %+v; want %d/%d", metric, covered, total)
	}
	wantPercent := 100 * float64(covered) / float64(total)
	if math.Abs(metric.Percent-wantPercent) > 1e-9 {
		t.Fatalf("percent = %f, want %f", metric.Percent, wantPercent)
	}
}

var errTestRead = errors.New("test read failure")

type errorReader struct{}

func (errorReader) Read([]byte) (int, error) { return 0, errTestRead }

type dataThenErrorReader struct {
	data []byte
	done bool
}

func (reader *dataThenErrorReader) Read(buffer []byte) (int, error) {
	if reader.done {
		return 0, errTestRead
	}
	reader.done = true
	return copy(buffer, reader.data), nil
}

var _ io.Reader = (*dataThenErrorReader)(nil)
