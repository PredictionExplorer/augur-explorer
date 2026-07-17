// Command extract reconstructs the combined-JSON build artifact of one or
// more binding packages from their committed Go bindings (see the buildjson
// package and contracts/README.md). It refreshes
// contracts/<pkg>/buildjson/combined.json after bindings are regenerated
// from real solc output on the contract author's machine:
//
//	go run ./contracts/internal/buildjson/extract contracts/cosmicgame contracts/randomwalk
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/PredictionExplorer/augur-explorer/contracts/internal/buildjson"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "extract: %v\n", err)
		os.Exit(1)
	}
}

func run(dirs []string) error {
	if len(dirs) == 0 {
		return fmt.Errorf("usage: extract <binding package dirs>")
	}
	for _, dir := range dirs {
		metas, err := buildjson.Extract(dir)
		if err != nil {
			return err
		}
		if len(metas) == 0 {
			return fmt.Errorf("%s: no bind.MetaData declarations found", dir)
		}
		// Bindings embed the ABI with all whitespace stripped; restore the
		// internalType spaces so abigen re-derives the named tuple structs
		// (see the buildjson package comment).
		for typeName, meta := range metas {
			restored, err := buildjson.RestoreInternalTypeSpaces(meta.ABI)
			if err != nil {
				return fmt.Errorf("%s: %s: %w", dir, typeName, err)
			}
			meta.ABI = restored
			metas[typeName] = meta
		}
		data, err := buildjson.Encode(metas)
		if err != nil {
			return fmt.Errorf("%s: %w", dir, err)
		}
		outPath := filepath.Join(dir, "buildjson", "combined.json")
		if err := os.MkdirAll(filepath.Dir(outPath), 0o750); err != nil { // #nosec G703 -- operator CLI; the target directory is its command-line argument
			return err
		}
		if err := os.WriteFile(outPath, data, 0o600); err != nil { // #nosec G703 -- operator CLI; the target directory is its command-line argument
			return err
		}
		fmt.Printf("%s: %d contract type(s)\n", outPath, len(metas))
	}
	return nil
}
