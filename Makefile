# RWCG backend — single entry point for building, testing and linting.
#
# All binaries are placed in ./bin. Run `make help` for a summary.

SHELL := /bin/bash
BIN   := bin

# Every runnable service and CLI in cmd/.
COMMANDS := apiserver cg-etl rw-etl notibot freezer-scan freezer-verify \
            srvmonitor loganomaly imggen-monitor rwalk-alarm \
            cgctl rwctl opsctl covergate

# Build identity, stamped into internal/version (see /version and --version).
# Overridable for reproducible builds: make build VERSION=v1.2.3 ...
VERSION    ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo unknown)
COMMIT     ?= $(shell git rev-parse HEAD 2>/dev/null || echo unknown)
BUILD_DATE ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
VERSION_PKG := github.com/PredictionExplorer/augur-explorer/internal/version
LDFLAGS := -X $(VERSION_PKG).version=$(VERSION) \
           -X $(VERSION_PKG).commit=$(COMMIT) \
           -X $(VERSION_PKG).buildDate=$(BUILD_DATE)

.PHONY: all build $(COMMANDS) generate generate-check test test-integration coverage-check hooks-install fuzz-smoke lint migrate-up fmt vet vuln clean help

all: build

## build: compile every command into ./bin
build:
	@mkdir -p $(BIN)
	go build -ldflags "$(LDFLAGS)" -o $(BIN)/ ./cmd/...

# Convenience per-command targets, e.g. `make apiserver`.
$(COMMANDS):
	@mkdir -p $(BIN)
	go build -ldflags "$(LDFLAGS)" -o $(BIN)/$@ ./cmd/$@

## generate: regenerate committed OpenAPI v2 server/models
generate:
	go generate ./internal/api/v2

## generate-check: fail when committed OpenAPI v2 generated code is stale
generate-check:
	go generate ./internal/api/v2
	git diff --exit-code -- internal/api/v2/api.gen.go

## test: run unit tests with the race detector
test:
	go test -race -shuffle=on ./...

## test-integration: run tests that need Docker (testcontainers)
test-integration:
	go test -race -tags=integration -timeout 20m ./...

## coverage-check: run the staged-diff and repository coverage gates
coverage-check:
	./scripts/coverage-gate.sh

## hooks-install: install the pre-commit hook (coverage blocking activates at 90%)
hooks-install:
	./scripts/install-hooks.sh

## fuzz-smoke: run every fuzz target briefly (10s each; FUZZTIME=30s to change)
fuzz-smoke:
	./scripts/fuzz-all.sh $${FUZZTIME:-10s}

## lint: run golangci-lint (install: brew install golangci-lint)
lint:
	golangci-lint run

## migrate-up: apply database migrations (uses PGSQL_* env or GOOSE_DBSTRING)
migrate-up:
	goose -dir db/migrations postgres "$${GOOSE_DBSTRING}" up

## fmt: format all Go code
fmt:
	gofmt -w .

## vet: run go vet
vet:
	go vet ./...

## vuln: scan dependencies for known vulnerabilities
vuln:
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## clean: remove build artifacts
clean:
	rm -rf $(BIN)

## help: list available targets
help:
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/## /  /'
