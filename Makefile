# RWCG backend — single entry point for building, testing and linting.
#
# All binaries are placed in ./bin. Run `make help` for a summary.

SHELL := /bin/bash
BIN   := bin

# Every runnable service and CLI in cmd/.
COMMANDS := apiserver cg-etl rw-etl notibot freezer-scan freezer-verify \
            srvmonitor loganomaly imggen-monitor rwalk-alarm \
            cgctl rwctl opsctl

.PHONY: all build $(COMMANDS) test test-integration fuzz-smoke lint migrate-up fmt vet vuln clean help

all: build

## build: compile every command into ./bin
build:
	@mkdir -p $(BIN)
	go build -o $(BIN)/ ./cmd/...

# Convenience per-command targets, e.g. `make apiserver`.
$(COMMANDS):
	@mkdir -p $(BIN)
	go build -o $(BIN)/$@ ./cmd/$@

## test: run unit tests with the race detector
test:
	go test -race -shuffle=on ./...

## test-integration: run tests that need Docker (testcontainers)
test-integration:
	go test -race -tags=integration -timeout 20m ./...

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
