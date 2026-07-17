// Package buildjson maintains the combined-JSON build artifacts consumed by
// the abigen generate directives (contracts/*/buildjson/combined.json).
//
// The deployed contracts' Solidity sources live outside this repository, so
// the artifacts are reconstructed from the generated bindings' embedded
// bind.MetaData. The embedded ABI is lossy in exactly one way: abigen
// strips every whitespace rune — including the spaces inside solc's
// "struct Foo.Bar" / "enum Foo.Bar" / "contract Foo" internalType values —
// before embedding, and geth's ABI parser needs those spaces to derive
// tuple struct names. RestoreInternalTypeSpaces reverses that known loss
// when artifacts are written, and StripABIWhitespace replays abigen's
// stripping when artifacts are compared against bindings.
//
// Extract parses the binding sources and recovers, per binding file, every
// contract type with its ABI and bytecode; the drift test in this package
// pins artifacts ⇄ bindings in both directions, and `make generate-check`
// proves a pinned-abigen regeneration reproduces the committed bindings
// byte-for-byte.
package buildjson

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// Meta is one contract type recovered from a binding file: the embedded ABI
// JSON (compact, as generated) and the creation bytecode without the 0x
// prefix ("" for non-deployable types).
type Meta struct {
	ABI string
	Bin string
}

// Extract parses every non-test .go file in dir and returns the contract
// metadata keyed by type name. Files without bind.MetaData declarations
// (doc.go, generate.go) contribute nothing; a type declared twice is an
// error (the package must build, so this cannot happen in practice).
func Extract(dir string) (map[string]Meta, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	fset := token.NewFileSet()
	out := make(map[string]Meta)
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		file, err := parser.ParseFile(fset, filepath.Join(dir, name), nil, 0)
		if err != nil {
			return nil, fmt.Errorf("parse %s: %w", name, err)
		}
		metas, err := extractFile(file)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", name, err)
		}
		for typeName, meta := range metas {
			if _, dup := out[typeName]; dup {
				return nil, fmt.Errorf("%s: type %s declared in more than one file", name, typeName)
			}
			out[typeName] = meta
		}
	}
	return out, nil
}

// extractFile collects every `var XMetaData = &bind.MetaData{...}`
// declaration of one parsed file.
func extractFile(file *ast.File) (map[string]Meta, error) {
	metas := make(map[string]Meta)
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.VAR {
			continue
		}
		for _, spec := range gen.Specs {
			vs, ok := spec.(*ast.ValueSpec)
			if !ok || len(vs.Names) != 1 || len(vs.Values) != 1 {
				continue
			}
			varName := vs.Names[0].Name
			typeName, ok := strings.CutSuffix(varName, "MetaData")
			if !ok || typeName == "" {
				continue
			}
			lit := compositeLit(vs.Values[0])
			if lit == nil || !isBindMetaData(lit.Type) {
				continue
			}
			meta, err := metaFromLit(lit)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", varName, err)
			}
			metas[typeName] = meta
		}
	}
	return metas, nil
}

// compositeLit unwraps `&T{...}` (or a plain `T{...}`) to the composite
// literal.
func compositeLit(expr ast.Expr) *ast.CompositeLit {
	if unary, ok := expr.(*ast.UnaryExpr); ok && unary.Op == token.AND {
		expr = unary.X
	}
	lit, _ := expr.(*ast.CompositeLit)
	return lit
}

// isBindMetaData reports whether the composite literal type is
// bind.MetaData (any import alias whose selector is MetaData qualifies —
// the bindings use the canonical `bind` name).
func isBindMetaData(expr ast.Expr) bool {
	sel, ok := expr.(*ast.SelectorExpr)
	return ok && sel.Sel.Name == "MetaData"
}

// metaFromLit reads the ABI and Bin string fields of a bind.MetaData
// composite literal.
func metaFromLit(lit *ast.CompositeLit) (Meta, error) {
	var meta Meta
	for _, elt := range lit.Elts {
		kv, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		key, ok := kv.Key.(*ast.Ident)
		if !ok {
			continue
		}
		basic, ok := kv.Value.(*ast.BasicLit)
		if !ok || basic.Kind != token.STRING {
			continue
		}
		value, err := strconv.Unquote(basic.Value)
		if err != nil {
			return Meta{}, fmt.Errorf("unquote %s: %w", key.Name, err)
		}
		switch key.Name {
		case "ABI":
			meta.ABI = value
		case "Bin":
			meta.Bin = strings.TrimPrefix(value, "0x")
		}
	}
	if meta.ABI == "" {
		return Meta{}, fmt.Errorf("no ABI literal found")
	}
	return meta, nil
}

// combinedJSON is the solc-combined-output subset abigen consumes
// (common/compiler solcOutputV8 rules: abi as a JSON document, bin without
// the 0x prefix).
type combinedJSON struct {
	Contracts map[string]combinedContract `json:"contracts"`
	Version   string                      `json:"version"`
}

type combinedContract struct {
	ABI json.RawMessage `json:"abi"`
	Bin string          `json:"bin"`
}

// artifactVersion documents artifact provenance inside each file.
const artifactVersion = "reconstructed from the committed Go bindings; see contracts/README.md"

// Encode renders one package's contract set as a stable, readable
// combined.json document (keys sorted by encoding/json, trailing newline).
// Each type is qualified as "<Type>.sol:<Type>" — abigen only uses the part
// after the colon.
func Encode(metas map[string]Meta) ([]byte, error) {
	doc := combinedJSON{
		Contracts: make(map[string]combinedContract, len(metas)),
		Version:   artifactVersion,
	}
	for typeName, meta := range metas {
		doc.Contracts[typeName+".sol:"+typeName] = combinedContract{
			ABI: json.RawMessage(meta.ABI),
			Bin: meta.Bin,
		}
	}
	data, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(data, '\n'), nil
}

// Decode parses a combined.json artifact back into the per-type metadata
// (ABI compacted to the canonical abigen form, bin without 0x).
func Decode(data []byte) (map[string]Meta, error) {
	var doc combinedJSON
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, err
	}
	metas := make(map[string]Meta, len(doc.Contracts))
	for qualified, contract := range doc.Contracts {
		typeName := qualified[strings.LastIndex(qualified, ":")+1:]
		abi, err := canonicalABI(contract.ABI)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", qualified, err)
		}
		metas[typeName] = Meta{ABI: abi, Bin: contract.Bin}
	}
	return metas, nil
}

// canonicalABI round-trips an ABI document the way abigen does
// (json.Unmarshal into any + json.Marshal), yielding the sorted-key compact
// form abigen embeds (before its whitespace stripping).
func canonicalABI(raw json.RawMessage) (string, error) {
	var doc any
	if err := json.Unmarshal(raw, &doc); err != nil {
		return "", err
	}
	compact, err := json.Marshal(doc)
	if err != nil {
		return "", err
	}
	return string(compact), nil
}

// StripABIWhitespace removes every whitespace rune, replicating what abigen
// does to the ABI before embedding it in a binding (accounts/abi/abigen:
// strings.Map over unicode.IsSpace). Comparing artifacts against bindings
// must go through this.
func StripABIWhitespace(abi string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, abi)
}

// internalTypePrefixes are solc's composite internalType markers. In real
// solc output each is followed by one space and the source-level type name;
// abigen's whitespace stripping glues them together.
var internalTypePrefixes = []string{"struct", "enum", "contract"}

// RestoreInternalTypeSpaces reverses abigen's whitespace stripping inside a
// binding-embedded ABI document: every "internalType" value reading
// "structFoo.Bar", "enumFoo.Bar" or "contractFoo" regains the space solc
// emitted. Elementary internalTypes (uint256, string, bytes32, ...) never
// start with a marker directly followed by a capitalized name, so the
// restoration is unambiguous for solc-produced ABIs; without it a
// regeneration would lose every named tuple struct (geth derives
// TupleRawName from the "struct " prefix).
func RestoreInternalTypeSpaces(abi string) (string, error) {
	var doc any
	if err := json.Unmarshal([]byte(abi), &doc); err != nil {
		return "", err
	}
	restoreInternalTypes(doc)
	out, err := json.Marshal(doc)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// restoreInternalTypes walks the decoded ABI document and fixes every
// internalType string in place.
func restoreInternalTypes(node any) {
	switch v := node.(type) {
	case map[string]any:
		if s, ok := v["internalType"].(string); ok {
			v["internalType"] = restoreOnePrefix(s)
		}
		for _, child := range v {
			restoreInternalTypes(child)
		}
	case []any:
		for _, child := range v {
			restoreInternalTypes(child)
		}
	}
}

// restoreOnePrefix inserts the stripped space after a composite-type marker
// ("structFoo.Bar" → "struct Foo.Bar"). Values that already carry a space,
// equal a bare marker, or continue in lowercase (elementary types like
// "string") pass through unchanged.
func restoreOnePrefix(value string) string {
	for _, prefix := range internalTypePrefixes {
		rest, ok := strings.CutPrefix(value, prefix)
		if !ok || rest == "" || strings.HasPrefix(rest, " ") {
			continue
		}
		if r := rest[0]; r >= 'A' && r <= 'Z' {
			return prefix + " " + rest
		}
	}
	return value
}

// SortedKeys returns the map keys sorted, for deterministic reports.
func SortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
