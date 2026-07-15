// Package version reports what build of the software is running: the
// version/tag, VCS commit and build timestamp. Release builds stamp the
// values through -ldflags (see the Makefile and deploy/docker/Dockerfile);
// plain `go build` binaries fall back to the toolchain's embedded VCS
// metadata (debug.ReadBuildInfo), so the information is best-effort but
// never empty.
package version

import (
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"runtime/debug"
)

// Stamped via -ldflags="-X github.com/PredictionExplorer/augur-explorer/internal/version.version=..."
// (and .commit= / .buildDate=). Empty means "not stamped": Get falls back to
// the build info embedded by the Go toolchain.
var (
	version   string
	commit    string
	buildDate string
)

// unknown is the placeholder for fields that neither ldflags nor the
// toolchain's VCS stamping could fill (e.g. a build outside a git checkout).
const unknown = "unknown"

// Info is the resolved build identity of the running binary.
type Info struct {
	// Version is the release tag or `git describe` output ("devel" when
	// nothing was stamped and the module has no version).
	Version string `json:"version"`
	// Commit is the full VCS revision, with a "-dirty" suffix when the
	// working tree had uncommitted changes.
	Commit string `json:"commit"`
	// BuildDate is the UTC RFC 3339 commit or build timestamp.
	BuildDate string `json:"buildDate"`
	// GoVersion is the toolchain that produced the binary.
	GoVersion string `json:"goVersion"`
}

// Get resolves the build identity: ldflags-stamped values win, the
// toolchain's embedded VCS metadata fills the gaps.
func Get() Info {
	bi, _ := debug.ReadBuildInfo()
	return resolve(version, commit, buildDate, bi)
}

// resolve merges the stamped values with build-info fallbacks. Split from
// Get so tests can exercise every precedence rule with synthetic inputs.
func resolve(ldVersion, ldCommit, ldDate string, bi *debug.BuildInfo) Info {
	info := Info{
		Version:   ldVersion,
		Commit:    ldCommit,
		BuildDate: ldDate,
		GoVersion: runtime.Version(),
	}

	var vcsRevision, vcsTime, vcsModified string
	if bi != nil {
		for _, s := range bi.Settings {
			switch s.Key {
			case "vcs.revision":
				vcsRevision = s.Value
			case "vcs.time":
				vcsTime = s.Value
			case "vcs.modified":
				vcsModified = s.Value
			}
		}
	}

	if info.Version == "" {
		if bi != nil && bi.Main.Version != "" && bi.Main.Version != "(devel)" {
			info.Version = bi.Main.Version
		} else {
			info.Version = "devel"
		}
	}
	if info.Commit == "" {
		info.Commit = vcsRevision
	}
	if info.Commit != "" && vcsModified == "true" && ldCommit == "" {
		info.Commit += "-dirty"
	}
	if info.Commit == "" {
		info.Commit = unknown
	}
	if info.BuildDate == "" {
		info.BuildDate = vcsTime
	}
	if info.BuildDate == "" {
		info.BuildDate = unknown
	}
	return info
}

// String renders the identity on one line, e.g.
// "v1.4.0 (commit 85941dba…, built 2026-07-13T18:00:00Z, go1.26.5)".
func String() string {
	return Get().String()
}

// String renders i on one line.
func (i Info) String() string {
	return fmt.Sprintf("%s (commit %s, built %s, %s)", i.Version, i.Commit, i.BuildDate, i.GoVersion)
}

// LogAttrs returns the identity as slog attributes for the one startup
// "build info" record every service emits.
func LogAttrs() []slog.Attr {
	i := Get()
	return []slog.Attr{
		slog.String("version", i.Version),
		slog.String("commit", i.Commit),
		slog.String("build_date", i.BuildDate),
		slog.String("go_version", i.GoVersion),
	}
}

// HandleFlag implements the --version convention for the plain-main
// binaries: when args contains "--version" or "-version" it prints the
// one-line identity to out and reports true (the caller exits). Cobra CLIs
// use root.Version instead.
func HandleFlag(args []string, out io.Writer) bool {
	for _, a := range args {
		if a == "--version" || a == "-version" {
			fmt.Fprintln(out, String())
			return true
		}
	}
	return false
}
