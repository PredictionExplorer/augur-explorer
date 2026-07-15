package version

import (
	"bytes"
	"encoding/json"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"
)

func buildInfo(settings map[string]string, mainVersion string) *debug.BuildInfo {
	bi := &debug.BuildInfo{}
	bi.Main.Version = mainVersion
	for k, v := range settings {
		bi.Settings = append(bi.Settings, debug.BuildSetting{Key: k, Value: v})
	}
	return bi
}

func TestResolvePrecedence(t *testing.T) {
	vcs := map[string]string{
		"vcs.revision": "85941dbaffffffffffffffffffffffffffffffff",
		"vcs.time":     "2026-07-13T18:00:00Z",
		"vcs.modified": "false",
	}

	tests := []struct {
		name                        string
		ldVersion, ldCommit, ldDate string
		bi                          *debug.BuildInfo
		want                        Info
	}{
		{
			name:      "ldflags win over build info",
			ldVersion: "v1.4.0", ldCommit: "deadbeef", ldDate: "2026-07-14T00:00:00Z",
			bi: buildInfo(vcs, "(devel)"),
			want: Info{
				Version: "v1.4.0", Commit: "deadbeef", BuildDate: "2026-07-14T00:00:00Z",
			},
		},
		{
			name: "vcs metadata fills unstamped fields",
			bi:   buildInfo(vcs, "(devel)"),
			want: Info{
				Version:   "devel",
				Commit:    "85941dbaffffffffffffffffffffffffffffffff",
				BuildDate: "2026-07-13T18:00:00Z",
			},
		},
		{
			name: "dirty tree marks the vcs commit",
			bi: buildInfo(map[string]string{
				"vcs.revision": "85941dba",
				"vcs.time":     "2026-07-13T18:00:00Z",
				"vcs.modified": "true",
			}, "(devel)"),
			want: Info{
				Version:   "devel",
				Commit:    "85941dba-dirty",
				BuildDate: "2026-07-13T18:00:00Z",
			},
		},
		{
			name:     "stamped commit is never marked dirty",
			ldCommit: "deadbeef",
			bi: buildInfo(map[string]string{
				"vcs.revision": "85941dba",
				"vcs.modified": "true",
			}, "(devel)"),
			want: Info{
				Version:   "devel",
				Commit:    "deadbeef",
				BuildDate: unknown,
			},
		},
		{
			name: "module version used when tagged",
			bi:   buildInfo(nil, "v2.3.1"),
			want: Info{
				Version:   "v2.3.1",
				Commit:    unknown,
				BuildDate: unknown,
			},
		},
		{
			name: "nil build info degrades to placeholders",
			bi:   nil,
			want: Info{
				Version:   "devel",
				Commit:    unknown,
				BuildDate: unknown,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resolve(tt.ldVersion, tt.ldCommit, tt.ldDate, tt.bi)
			tt.want.GoVersion = runtime.Version()
			if got != tt.want {
				t.Errorf("resolve() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestGetNeverEmpty(t *testing.T) {
	info := Get()
	if info.Version == "" || info.Commit == "" || info.BuildDate == "" || info.GoVersion == "" {
		t.Fatalf("Get() returned an empty field: %+v", info)
	}
}

func TestStringFormat(t *testing.T) {
	i := Info{Version: "v1.0.0", Commit: "abc123", BuildDate: "2026-07-13T18:00:00Z", GoVersion: "go1.26.5"}
	want := "v1.0.0 (commit abc123, built 2026-07-13T18:00:00Z, go1.26.5)"
	if got := i.String(); got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}

func TestJSONShape(t *testing.T) {
	i := Info{Version: "v1.0.0", Commit: "abc", BuildDate: "d", GoVersion: "go1.26.5"}
	data, err := json.Marshal(i)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]string
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"version", "commit", "buildDate", "goVersion"} {
		if m[key] == "" {
			t.Errorf("JSON field %q missing or empty in %s", key, data)
		}
	}
	if len(m) != 4 {
		t.Errorf("JSON has %d fields, want 4: %s", len(m), data)
	}
}

func TestLogAttrs(t *testing.T) {
	attrs := LogAttrs()
	keys := make(map[string]bool)
	for _, a := range attrs {
		if a.Value.String() == "" {
			t.Errorf("attr %q has empty value", a.Key)
		}
		keys[a.Key] = true
	}
	for _, want := range []string{"version", "commit", "build_date", "go_version"} {
		if !keys[want] {
			t.Errorf("LogAttrs() missing %q (got %v)", want, keys)
		}
	}
}

func TestHandleFlag(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want bool
	}{
		{"double dash", []string{"--version"}, true},
		{"single dash", []string{"-version"}, true},
		{"anywhere in args", []string{"--twitter", "--version"}, true},
		{"no flag", []string{"--twitter"}, false},
		{"empty args", nil, false},
		{"not a prefix match", []string{"--versions"}, false},
		{"positional value ignored", []string{"version"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			got := HandleFlag(tt.args, &buf)
			if got != tt.want {
				t.Fatalf("HandleFlag(%v) = %v, want %v", tt.args, got, tt.want)
			}
			if tt.want && !strings.Contains(buf.String(), "commit") {
				t.Errorf("output %q does not look like a version line", buf.String())
			}
			if !tt.want && buf.Len() != 0 {
				t.Errorf("HandleFlag printed %q without the flag", buf.String())
			}
		})
	}
}
