# Multi-stage build for any RWCG service.
#
# Build a specific service with:
#   docker build -f deploy/docker/Dockerfile --build-arg CMD=apiserver -t rwcg/apiserver .
#
# CMD must be a directory name under cmd/ (apiserver, cg-etl, rw-etl, notibot, ...).

FROM golang:1.26-alpine AS build
ARG CMD=apiserver
WORKDIR /src

# Cache module downloads separately from source changes.
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/service ./cmd/${CMD}

FROM alpine:3.22
ARG CMD=apiserver
RUN adduser -D -H rwcg && apk add --no-cache ca-certificates tzdata
USER rwcg
COPY --from=build /out/service /usr/local/bin/service

WORKDIR /app
ENTRYPOINT ["/usr/local/bin/service"]
