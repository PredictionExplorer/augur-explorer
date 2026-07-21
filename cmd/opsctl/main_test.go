package main

import (
	"bytes"
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRootCommandTree(t *testing.T) {
	t.Parallel()
	root := newRootCmd()
	want := map[string][]string{
		"archive":      {"corpus-export", "export", "node-fill", "verify"},
		"assets":       {"gen-thumbnails", "inventory", "verify-token-images"},
		"db":           {"evtlog-diff", "verify"},
		"scan":         {"cst-auction-len"},
		"smoketest":    {},
		"tx-collector": {"run", "verify"},
	}

	if got := childCommandNames(root); !reflect.DeepEqual(got, sortedKeys(want)) {
		t.Fatalf("root commands = %v, want %v", got, sortedKeys(want))
	}
	for name, wantChildren := range want {
		command := childCommand(t, root, name)
		if got := childCommandNames(command); !reflect.DeepEqual(got, wantChildren) {
			t.Errorf("%s subcommands = %v, want %v", name, got, wantChildren)
		}
	}
}

func TestRootCommandsAreFresh(t *testing.T) {
	t.Parallel()
	first := newRootCmd()
	second := newRootCmd()

	firstExport := childCommand(t, childCommand(t, first, "archive"), "export")
	if err := firstExport.Flags().Set("project", "both"); err != nil {
		t.Fatalf("setting first command flag: %v", err)
	}
	secondExport := childCommand(t, childCommand(t, second, "archive"), "export")
	if got, err := secondExport.Flags().GetString("project"); err != nil || got != "" {
		t.Fatalf("second command inherited flag state: value=%q error=%v", got, err)
	}
	if first == second || firstExport == secondExport {
		t.Fatal("command constructors reused mutable cobra commands")
	}
}

func TestRequiredFlagsFailBeforeExternalSetup(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		args []string
		want string
	}{
		{name: "archive corpus export", args: []string{"archive", "corpus-export"}, want: `required flag(s) "db", "project", "tx-hash"`},
		{name: "archive export", args: []string{"archive", "export"}, want: `required flag(s) "dst", "project", "src"`},
		{name: "archive node fill", args: []string{"archive", "node-fill"}, want: `required flag(s) "db", "project"`},
		{name: "archive verify", args: []string{"archive", "verify"}, want: `required flag(s) "db", "project"`},
		{name: "database diff", args: []string{"db", "evtlog-diff"}, want: `required flag(s) "primary", "secondary"`},
		{name: "database verify", args: []string{"db", "verify"}, want: `required flag(s) "primary", "secondary"`},
		{name: "collector run", args: []string{"tx-collector", "run"}, want: `required flag(s) "config"`},
		{name: "collector verify", args: []string{"tx-collector", "verify"}, want: `required flag(s) "config"`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			root := newRootCmd()
			var output bytes.Buffer
			root.SetOut(&output)
			root.SetErr(&output)
			root.SetArgs(test.args)
			err := root.Execute()
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("Execute(%v) error = %v, want containing %q", test.args, err, test.want)
			}
		})
	}
}

func TestRootVersionFlag(t *testing.T) {
	t.Parallel()
	root := newRootCmd()
	var output bytes.Buffer
	root.SetOut(&output)
	root.SetErr(&output)
	root.SetArgs([]string{"--version"})
	if err := root.Execute(); err != nil {
		t.Fatalf("Execute(--version) error = %v", err)
	}
	got := output.String()
	if !strings.HasPrefix(got, "opsctl version ") || !strings.Contains(got, "commit") {
		t.Fatalf("--version output = %q, want the build identity line", got)
	}
}

func childCommand(t *testing.T, parent *cobra.Command, name string) *cobra.Command {
	t.Helper()
	for _, command := range parent.Commands() {
		if command.Name() == name {
			return command
		}
	}
	t.Fatalf("%s has no child %q", parent.CommandPath(), name)
	return nil
}

func childCommandNames(command *cobra.Command) []string {
	names := make([]string, 0, len(command.Commands()))
	for _, child := range command.Commands() {
		names = append(names, child.Name())
	}
	slices.Sort(names)
	return names
}

func sortedKeys(values map[string][]string) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}
