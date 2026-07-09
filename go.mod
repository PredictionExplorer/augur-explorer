module github.com/PredictionExplorer/augur-explorer

go 1.26.5

// Fix ambiguous btcd import
exclude github.com/btcsuite/btcd/chaincfg/chainhash v1.0.1

require (
	github.com/ethereum/go-ethereum v1.17.4
	github.com/gin-gonic/gin v1.12.0 // indirect
	github.com/lib/pq v1.12.3
	//	github.com/wealdtech/go-ens 2be8e3e5fa10e897e9957584302fae93d43d6cc1
	golang.org/x/crypto v0.53.0 // indirect
	golang.org/x/net v0.56.0 // indirect
)

require (
	github.com/DataDog/zstd v1.5.2 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/ProjectZKM/Ziren/crates/go-runtime/zkvm_runtime v0.0.0-20251001021608-1fe7b43fc4d6 // indirect
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/andersfylling/disgord v0.32.3
	github.com/andersfylling/snowflake/v5 v5.0.1 // indirect
	github.com/aws/aws-sdk-go v1.38.20 // indirect
	github.com/bits-and-blooms/bitset v1.20.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/consensys/gnark-crypto v0.18.1 // indirect
	github.com/crate-crypto/go-eth-kzg v1.5.0 // indirect
	github.com/deckarep/golang-set/v2 v2.6.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/deepmap/oapi-codegen v1.8.2 // indirect
	github.com/ethereum/c-kzg-4844/v2 v2.1.6 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/golang/snappy v1.0.0
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/holiman/uint256 v1.3.2 // indirect
	github.com/influxdata/line-protocol v0.0.0-20210311194329-9aa0e372d097 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/klauspost/compress v1.18.5 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/nsf/termbox-go v1.1.1
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/shirou/gopsutil v3.21.5+incompatible // indirect
	github.com/supranational/blst v0.3.16 // indirect
	github.com/tklauser/go-sysconf v0.3.16 // indirect
	github.com/tklauser/numcpus v0.11.0 // indirect
	github.com/u2takey/ffmpeg-go v0.4.1
	github.com/u2takey/go-utils v0.3.1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/sync v0.21.0 // indirect
	golang.org/x/sys v0.46.0
	golang.org/x/text v0.38.0 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
	nhooyr.io/websocket v1.8.7 // indirect
)

require (
	github.com/goccy/go-yaml v1.19.2
	github.com/jackc/pgx/v5 v5.10.0
	github.com/pressly/goose/v3 v3.27.2
	github.com/prometheus/client_golang v1.23.2
	github.com/spf13/cobra v1.10.2
	github.com/testcontainers/testcontainers-go v0.43.0
	github.com/testcontainers/testcontainers-go/modules/postgres v0.43.0
	golang.org/x/time v0.15.0
)

require (
	dario.cat/mergo v1.0.2 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20250102033503-faa5f7b0171c // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/cp v1.1.1 // indirect
	github.com/containerd/errdefs v1.0.0 // indirect
	github.com/containerd/errdefs/pkg v0.3.0 // indirect
	github.com/containerd/log v0.1.0 // indirect
	github.com/containerd/platforms v0.2.1 // indirect
	github.com/cpuguy83/dockercfg v0.3.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/distribution/reference v0.6.0 // indirect
	github.com/docker/go-connections v0.7.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/ebitengine/purego v0.10.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fjl/jsonw v0.1.0 // indirect
	github.com/gballet/go-libpcsclite v0.0.0-20191108122812-4678299bea08 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/magiconair/properties v1.8.10 // indirect
	github.com/mfridman/interpolate v0.0.2 // indirect
	github.com/moby/docker-image-spec v1.3.1 // indirect
	github.com/moby/go-archive v0.2.0 // indirect
	github.com/moby/moby/api v1.55.0 // indirect
	github.com/moby/moby/client v0.5.0 // indirect
	github.com/moby/patternmatcher v0.6.1 // indirect
	github.com/moby/sys/sequential v0.6.0 // indirect
	github.com/moby/sys/user v0.4.0 // indirect
	github.com/moby/sys/userns v0.1.0 // indirect
	github.com/moby/term v0.5.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.1 // indirect
	github.com/peterh/liner v1.2.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20240221224432-82ca36839d55 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.66.1 // indirect
	github.com/prometheus/procfs v0.20.1 // indirect
	github.com/sethvargo/go-retry v0.3.0 // indirect
	github.com/shirou/gopsutil/v4 v4.26.5 // indirect
	github.com/sirupsen/logrus v1.9.4 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	github.com/stretchr/testify v1.11.1 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.68.0 // indirect
	go.opentelemetry.io/otel v1.43.0 // indirect
	go.opentelemetry.io/otel/metric v1.43.0 // indirect
	go.opentelemetry.io/otel/trace v1.43.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.yaml.in/yaml/v2 v2.4.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
