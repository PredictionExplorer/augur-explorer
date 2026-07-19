package covergate

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"path"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var profileLinePattern = regexp.MustCompile(
	`^(.+):([0-9]+)\.([0-9]+),([0-9]+)\.([0-9]+) ([0-9]+) ([0-9]+)$`,
)

// Block is one unique statement block from a Go coverage profile.
type Block struct {
	File        string
	StartLine   int
	StartColumn int
	EndLine     int
	EndColumn   int
	Statements  int
	Count       uint64
}

// Profile is a normalized, deduplicated Go coverage profile.
type Profile struct {
	Mode   string
	Blocks []Block
}

type blockKey struct {
	file        string
	startLine   int
	startColumn int
	endLine     int
	endColumn   int
}

// ParseProfile parses a Go text coverage profile and unions duplicate blocks
// emitted by separate test binaries.
func ParseProfile(reader io.Reader) (Profile, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return Profile{}, fmt.Errorf("read coverage mode: %w", err)
		}
		return Profile{}, errors.New("coverage profile is empty")
	}
	modeLine := strings.TrimSpace(scanner.Text())
	if !strings.HasPrefix(modeLine, "mode: ") {
		return Profile{}, fmt.Errorf("invalid coverage mode line %q", modeLine)
	}
	mode := strings.TrimPrefix(modeLine, "mode: ")
	switch mode {
	case "set", "count", "atomic":
	default:
		return Profile{}, fmt.Errorf("unsupported coverage mode %q", mode)
	}

	blocks := make(map[blockKey]Block)
	lineNumber := 1
	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		block, err := parseProfileLine(line)
		if err != nil {
			return Profile{}, fmt.Errorf("coverage profile line %d: %w", lineNumber, err)
		}
		key := blockKey{
			file:        block.File,
			startLine:   block.StartLine,
			startColumn: block.StartColumn,
			endLine:     block.EndLine,
			endColumn:   block.EndColumn,
		}
		if previous, ok := blocks[key]; ok {
			if previous.Statements != block.Statements {
				return Profile{}, fmt.Errorf(
					"coverage profile line %d: duplicate block has %d and %d statements",
					lineNumber,
					previous.Statements,
					block.Statements,
				)
			}
			if block.Count > previous.Count {
				previous.Count = block.Count
				blocks[key] = previous
			}
			continue
		}
		blocks[key] = block
	}
	if err := scanner.Err(); err != nil {
		return Profile{}, fmt.Errorf("read coverage profile: %w", err)
	}
	if len(blocks) == 0 {
		return Profile{}, errors.New("coverage profile contains no statement blocks")
	}

	normalized := make([]Block, 0, len(blocks))
	for _, block := range blocks {
		normalized = append(normalized, block)
	}
	slices.SortFunc(normalized, func(left, right Block) int {
		if compared := strings.Compare(left.File, right.File); compared != 0 {
			return compared
		}
		if left.StartLine != right.StartLine {
			return left.StartLine - right.StartLine
		}
		if left.StartColumn != right.StartColumn {
			return left.StartColumn - right.StartColumn
		}
		if left.EndLine != right.EndLine {
			return left.EndLine - right.EndLine
		}
		return left.EndColumn - right.EndColumn
	})
	return Profile{Mode: mode, Blocks: normalized}, nil
}

func parseProfileLine(line string) (Block, error) {
	match := profileLinePattern.FindStringSubmatch(line)
	if match == nil {
		return Block{}, fmt.Errorf("invalid statement block %q", line)
	}
	values := make([]int, 0, 4)
	for _, raw := range match[2:6] {
		value, err := strconv.Atoi(raw)
		if err != nil || value < 1 {
			return Block{}, fmt.Errorf("invalid positive block field %q", raw)
		}
		values = append(values, value)
	}
	statements, err := strconv.Atoi(match[6])
	if err != nil || statements < 0 {
		return Block{}, fmt.Errorf("invalid statement count %q", match[6])
	}
	count, err := strconv.ParseUint(match[7], 10, 64)
	if err != nil {
		return Block{}, fmt.Errorf("invalid execution count %q", match[7])
	}
	file, err := normalizeProfilePath(match[1])
	if err != nil {
		return Block{}, err
	}
	if values[2] < values[0] ||
		(values[2] == values[0] && values[3] < values[1]) {
		return Block{}, errors.New("coverage block ends before it starts")
	}
	return Block{
		File:        file,
		StartLine:   values[0],
		StartColumn: values[1],
		EndLine:     values[2],
		EndColumn:   values[3],
		Statements:  statements,
		Count:       count,
	}, nil
}

func normalizeProfilePath(file string) (string, error) {
	file = strings.ReplaceAll(file, `\`, "/")
	file = strings.TrimPrefix(file, "./")
	best := -1
	for _, marker := range []string{"/cmd/", "/internal/"} {
		if index := strings.Index(file, marker); index >= 0 &&
			(best < 0 || index < best) {
			best = index + 1
		}
	}
	if best >= 0 {
		file = file[best:]
	}
	if !strings.HasPrefix(file, "cmd/") && !strings.HasPrefix(file, "internal/") {
		return "", fmt.Errorf("coverage path is outside cmd/ and internal/: %q", file)
	}
	cleaned := path.Clean(file)
	if cleaned == "." || strings.HasPrefix(cleaned, "../") {
		return "", fmt.Errorf("unsafe coverage path %q", file)
	}
	return cleaned, nil
}

// Metric is one weighted statement-coverage result.
type Metric struct {
	Covered int     `json:"covered"`
	Total   int     `json:"total"`
	Percent float64 `json:"percent"`
}

// Analysis contains repository-wide and per-package coverage metrics.
type Analysis struct {
	LegacyInternal Metric            `json:"legacyInternal"`
	Internal       Metric            `json:"internal"`
	Production     Metric            `json:"production"`
	Packages       map[string]Metric `json:"packages"`
}

// Analyze computes raw legacy and policy-filtered production metrics.
func Analyze(profile Profile, policy Policy) Analysis {
	var legacyInternal metricBuilder
	var internal metricBuilder
	var production metricBuilder
	packages := make(map[string]*metricBuilder)
	for _, block := range profile.Blocks {
		isInternal := strings.HasPrefix(block.File, "internal/")
		if isInternal {
			legacyInternal.add(block)
		}
		if policy.excluded(block.File) {
			continue
		}
		if !isInternal && !strings.HasPrefix(block.File, "cmd/") {
			continue
		}
		production.add(block)
		if isInternal {
			internal.add(block)
		}
		packageName := path.Dir(block.File)
		builder := packages[packageName]
		if builder == nil {
			builder = &metricBuilder{}
			packages[packageName] = builder
		}
		builder.add(block)
	}
	packageMetrics := make(map[string]Metric, len(packages))
	for packageName, builder := range packages {
		packageMetrics[packageName] = builder.metric()
	}
	return Analysis{
		LegacyInternal: legacyInternal.metric(),
		Internal:       internal.metric(),
		Production:     production.metric(),
		Packages:       packageMetrics,
	}
}

type metricBuilder struct {
	covered int
	total   int
}

func (builder *metricBuilder) add(block Block) {
	builder.total += block.Statements
	if block.Count > 0 {
		builder.covered += block.Statements
	}
}

func (builder *metricBuilder) metric() Metric {
	percent := 0.0
	if builder.total > 0 {
		percent = 100 * float64(builder.covered) / float64(builder.total)
	}
	return Metric{
		Covered: builder.covered,
		Total:   builder.total,
		Percent: percent,
	}
}
