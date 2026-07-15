# Multi-stage build for any RWCG service.
#
# Build a specific service with:
#   docker build -f deploy/docker/Dockerfile --build-arg CMD=apiserver \
#     --build-arg VERSION="$(git describe --tags --always --dirty)" \
#     --build-arg COMMIT="$(git rev-parse HEAD)" \
#     --build-arg BUILD_DATE="$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
#     -t rwcg/apiserver .
#
# CMD must be a directory name under cmd/ (apiserver, cg-etl, rw-etl, notibot, ...).
# VERSION/COMMIT/BUILD_DATE stamp internal/version (served at /version and by
# --version); unset they degrade to "unknown" because the build context
# carries no .git directory.

FROM golang:1.26-alpine AS build
ARG CMD=apiserver
ARG VERSION=unknown
ARG COMMIT=unknown
ARG BUILD_DATE=unknown
WORKDIR /src

# Cache module downloads separately from source changes.
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath \
    -ldflags="-s -w \
      -X github.com/PredictionExplorer/augur-explorer/internal/version.version=${VERSION} \
      -X github.com/PredictionExplorer/augur-explorer/internal/version.commit=${COMMIT} \
      -X github.com/PredictionExplorer/augur-explorer/internal/version.buildDate=${BUILD_DATE}" \
    -o /out/service ./cmd/${CMD}

FROM alpine:3.22
ARG CMD=apiserver
RUN adduser -D -H rwcg && apk add --no-cache ca-certificates tzdata
USER rwcg
COPY --from=build /out/service /usr/local/bin/service

WORKDIR /app
ENTRYPOINT ["/usr/local/bin/service"]
