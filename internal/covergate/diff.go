package covergate

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var hunkPattern = regexp.MustCompile(
	`^@@ -([0-9]+)(?:,([0-9]+))? \+([0-9]+)(?:,([0-9]+))? @@`,
)

// ChangedLines maps repository-relative files to added or modified line
// numbers in a unified Git diff.
type ChangedLines map[string]map[int]struct{}

// ParseUnifiedDiff parses added and modified lines from a Git unified diff.
func ParseUnifiedDiff(reader io.Reader) (ChangedLines, error) {
	changed := make(ChangedLines)
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, 64*1024), 4*1024*1024)

	var currentFile string
	oldLine := 0
	newLine := 0
	inHunk := false
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "+++ "):
			file, err := diffPath(strings.TrimPrefix(line, "+++ "))
			if err != nil {
				return nil, fmt.Errorf("diff line %d: %w", lineNumber, err)
			}
			currentFile = file
			inHunk = false
		case strings.HasPrefix(line, "@@ "):
			if currentFile == "" {
				return nil, fmt.Errorf("diff line %d: hunk has no target file", lineNumber)
			}
			match := hunkPattern.FindStringSubmatch(line)
			if match == nil {
				return nil, fmt.Errorf("diff line %d: malformed hunk header", lineNumber)
			}
			var err error
			oldLine, err = strconv.Atoi(match[1])
			if err != nil {
				return nil, fmt.Errorf("diff line %d: invalid old line", lineNumber)
			}
			newLine, err = strconv.Atoi(match[3])
			if err != nil {
				return nil, fmt.Errorf("diff line %d: invalid new line", lineNumber)
			}
			inHunk = true
		case inHunk && strings.HasPrefix(line, "+"):
			if currentFile != "/dev/null" {
				lines := changed[currentFile]
				if lines == nil {
					lines = make(map[int]struct{})
					changed[currentFile] = lines
				}
				lines[newLine] = struct{}{}
			}
			newLine++
		case inHunk && strings.HasPrefix(line, "-"):
			oldLine++
		case inHunk && strings.HasPrefix(line, " "):
			oldLine++
			newLine++
		case inHunk && strings.HasPrefix(line, `\`):
			// "\ No newline at end of file" does not consume either side.
		case strings.HasPrefix(line, "diff --git ") ||
			strings.HasPrefix(line, "index ") ||
			strings.HasPrefix(line, "--- ") ||
			strings.HasPrefix(line, "new file mode ") ||
			strings.HasPrefix(line, "deleted file mode ") ||
			strings.HasPrefix(line, "similarity index ") ||
			strings.HasPrefix(line, "rename from ") ||
			strings.HasPrefix(line, "rename to "):
			inHunk = false
		default:
			if inHunk {
				return nil, fmt.Errorf("diff line %d: malformed hunk body", lineNumber)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read unified diff: %w", err)
	}
	return changed, nil
}

func diffPath(raw string) (string, error) {
	if raw == "/dev/null" {
		return raw, nil
	}
	if strings.HasPrefix(raw, `"`) {
		unquoted, err := strconv.Unquote(raw)
		if err != nil {
			return "", fmt.Errorf("invalid quoted diff path: %w", err)
		}
		raw = unquoted
	}
	raw = strings.TrimPrefix(raw, "b/")
	return normalizeProfilePath(raw)
}

// PatchAnalysis is the changed-code statement coverage result.
type PatchAnalysis struct {
	Metric       Metric           `json:"metric"`
	Applicable   bool             `json:"applicable"`
	ChangedFiles int              `json:"changedFiles"`
	ChangedLines int              `json:"changedLines"`
	Uncovered    []UncoveredBlock `json:"uncovered"`
}

// UncoveredBlock identifies a changed executable block missed by the suite.
type UncoveredBlock struct {
	File       string `json:"file"`
	StartLine  int    `json:"startLine"`
	EndLine    int    `json:"endLine"`
	Statements int    `json:"statements"`
}

// AnalyzePatch intersects changed lines with policy-covered statement blocks.
func AnalyzePatch(profile Profile, policy Policy, changed ChangedLines) PatchAnalysis {
	var builder metricBuilder
	changedLineCount := 0
	changedFiles := 0
	uncovered := make([]UncoveredBlock, 0)
	for file, lines := range changed {
		if policy.excluded(file) ||
			(!strings.HasPrefix(file, "cmd/") && !strings.HasPrefix(file, "internal/")) {
			continue
		}
		changedFiles++
		changedLineCount += len(lines)
	}
	for _, block := range profile.Blocks {
		lines := changed[block.File]
		if len(lines) == 0 || policy.excluded(block.File) {
			continue
		}
		for line := range lines {
			if line >= block.StartLine && line <= block.EndLine {
				builder.add(block)
				if block.Count == 0 && block.Statements > 0 {
					uncovered = append(uncovered, UncoveredBlock{
						File:       block.File,
						StartLine:  block.StartLine,
						EndLine:    block.EndLine,
						Statements: block.Statements,
					})
				}
				break
			}
		}
	}
	metric := builder.metric()
	return PatchAnalysis{
		Metric:       metric,
		Applicable:   metric.Total > 0,
		ChangedFiles: changedFiles,
		ChangedLines: changedLineCount,
		Uncovered:    uncovered,
	}
}
