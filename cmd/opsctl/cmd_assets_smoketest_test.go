package main

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestAssetsCommandSurface(t *testing.T) {
	t.Parallel()
	assetsCmd := newAssetsCmd()
	if assetsCmd.Use != "assets" || assetsCmd.Short == "" {
		t.Fatalf("assets command = use %q short %q", assetsCmd.Use, assetsCmd.Short)
	}

	inventory := newAssetsInventoryCmd()
	if inventory.Use != "inventory" || inventory.Args(inventory, []string{"extra"}) == nil {
		t.Fatalf("inventory command surface changed")
	}
	assertFlagDefault(t, inventory, "db", "")
	assertFlagDefault(t, inventory, "base", ".")
	assertFlagDefault(t, inventory, "schema", "public")
	assertFlagDefault(t, inventory, "missing-only", "false")
	assertFlagDefault(t, inventory, "all", "false")

	thumbnails := newAssetsGenThumbnailsCmd()
	if thumbnails.Use != "gen-thumbnails" || thumbnails.Args(thumbnails, []string{"extra"}) == nil {
		t.Fatalf("thumbnail command surface changed")
	}
	assertFlagDefault(t, thumbnails, "db", "")
	assertFlagDefault(t, thumbnails, "base", "")
	assertFlagDefault(t, thumbnails, "schema", "public")
	assertFlagDefault(t, thumbnails, "force", "false")
	assertFlagDefault(t, thumbnails, "magick", "")
	assertFlagDefault(t, thumbnails, "min-age", "10")

	verify := newAssetsVerifyTokenImagesCmd()
	if verify.Use != "verify-token-images" || verify.Args(verify, []string{"extra"}) == nil {
		t.Fatalf("verify-token-images command surface changed")
	}
}

func TestSmoketestCommandSurface(t *testing.T) {
	t.Parallel()
	cmd := newSmoketestCmd()
	if cmd.Use != "smoketest" || cmd.Short == "" || cmd.Long == "" {
		t.Fatalf("smoketest command = use %q short %q", cmd.Use, cmd.Short)
	}
	if err := cmd.Args(cmd, []string{"extra"}); err == nil {
		t.Fatal("smoketest accepted a positional argument")
	}
	assertFlagDefault(t, cmd, "suite", "v2")
}

func TestAssetRunnersRejectInvalidSchemaBeforeExternalSetup(t *testing.T) {
	t.Parallel()
	cmd := &cobra.Command{}
	if err := runAssetInventory(cmd, "", t.TempDir(), "bad.schema", false, false); err == nil ||
		!strings.Contains(err.Error(), "invalid database schema") {
		t.Fatalf("inventory error = %v", err)
	}
	if err := runGenThumbnails(cmd, "", t.TempDir(), "bad.schema", false, "", 10); err == nil ||
		!strings.Contains(err.Error(), "invalid database schema") {
		t.Fatalf("thumbnail error = %v", err)
	}
}

func TestSmoketestAPIBase(t *testing.T) {
	t.Run("explicit base", func(t *testing.T) {
		t.Setenv("API_BASE", " https://api.example.test/root/// ")
		t.Setenv("HTTP_PORT", "1234")
		if got := smoketestAPIBase(); got != "https://api.example.test/root" {
			t.Fatalf("base = %q", got)
		}
	})
	t.Run("configured port", func(t *testing.T) {
		t.Setenv("API_BASE", "")
		t.Setenv("HTTP_PORT", "8123")
		if got := smoketestAPIBase(); got != "http://127.0.0.1:8123" {
			t.Fatalf("base = %q", got)
		}
	})
	t.Run("default port", func(t *testing.T) {
		t.Setenv("API_BASE", "")
		t.Setenv("HTTP_PORT", "")
		if got := smoketestAPIBase(); got != "http://127.0.0.1:9090" {
			t.Fatalf("base = %q", got)
		}
	})
}

func TestRunnerEnvironmentErrors(t *testing.T) {
	for _, name := range []string{
		"PGSQL_HOST", "PGSQL_USERNAME", "PGSQL_DATABASE", "PGSQL_PASSWORD", "NFT_ASSETS_ROOT",
	} {
		t.Setenv(name, "")
	}
	cmd := &cobra.Command{}
	if err := runAssetInventory(cmd, "", t.TempDir(), "public", false, false); err == nil ||
		!strings.Contains(err.Error(), "no --db flag") {
		t.Fatalf("inventory error = %v", err)
	}
	if err := runGenThumbnails(cmd, "", "", "public", false, "", 10); err == nil ||
		!strings.Contains(err.Error(), "NFT_ASSETS_ROOT") {
		t.Fatalf("thumbnail error = %v", err)
	}
	if err := runSmoketest(cmd); err == nil ||
		!strings.Contains(err.Error(), "PGSQL_HOST / PGSQL_USERNAME / PGSQL_DATABASE") {
		t.Fatalf("smoketest error = %v", err)
	}
}

func TestCommandsRejectUnsafeValuesBeforeExternalSetup(t *testing.T) {
	t.Run("negative thumbnail age", func(t *testing.T) {
		cmd := &cobra.Command{}
		err := runGenThumbnails(cmd, "", t.TempDir(), "public", false, "", -1)
		if err == nil || !strings.Contains(err.Error(), "--min-age must be non-negative") {
			t.Fatalf("negative min-age error = %v", err)
		}
	})
	t.Run("malformed scan contract", func(t *testing.T) {
		cmd := newScanCstAuctionLenCmd()
		cmd.SilenceErrors = true
		cmd.SilenceUsage = true
		cmd.SetArgs([]string{"--contract", "not-an-address"})
		err := cmd.Execute()
		if err == nil || !strings.Contains(err.Error(), "invalid --contract address") {
			t.Fatalf("malformed contract error = %v", err)
		}
	})
}

func TestRedactConn(t *testing.T) {
	t.Parallel()
	// #nosec G101 -- deliberately fake secrets exercise redaction.
	tests := map[string]string{
		"postgres://user:secret@db.example/app?sslmode=disable":       "postgres://user:%2A%2A%2A@db.example/app?sslmode=disable",
		"https://user:secret@rpc.example/path?api_key=hidden&chain=1": "https://user:%2A%2A%2A@rpc.example/path?api_key=%2A%2A%2A&chain=1",
		"https://rpc.example/path?token=hidden":                       "https://rpc.example/path?token=%2A%2A%2A",
		"host=db user=app password=secret dbname=app":                 "host=db user=app password=*** dbname=app",
		`host=db password="two words" user=app`:                       "host=db password=*** user=app",
		"host=db user=app":                                            "host=db user=app",
	}
	for input, want := range tests {
		if got := redactConn(input); got != want {
			t.Errorf("redactConn(%q) = %q, want %q", input, got, want)
		}
	}
	rpc := redactRPCURL("https://user:secret@rpc.example/v3/project-secret?token=query-secret")
	for _, secret := range []string{"secret", "project-secret", "query-secret"} {
		if strings.Contains(rpc, secret) {
			t.Errorf("redactRPCURL leaked %q in %q", secret, rpc)
		}
	}
	if !strings.Contains(rpc, "rpc.example") {
		t.Errorf("redactRPCURL removed endpoint identity: %q", rpc)
	}
}

func TestResolveMagickExplicitMissing(t *testing.T) {
	t.Parallel()
	_, err := resolveMagick("/definitely/not/a/real/imagemagick-binary")
	if err == nil || !strings.Contains(err.Error(), "not found on PATH") {
		t.Fatalf("error = %v", err)
	}
}

func assertFlagDefault(t *testing.T, cmd *cobra.Command, name, want string) {
	t.Helper()
	if got := cmd.Flags().Lookup(name); got == nil || got.DefValue != want {
		t.Fatalf("--%s default = %#v, want %q", name, got, want)
	}
}
