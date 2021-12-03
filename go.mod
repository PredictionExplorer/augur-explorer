module github.com/PredictionExplorer/augur-explorer

go 1.14

replace (
	github.com/ethereum/go-ethereum => github.com/ethereum/go-ethereum v1.9.13
	github.com/ethereum/go-ethereum@v0.0.0-00010101000000-000000000000 => github.com/ethereum/go-ethereum v1.9.13
)

require (
	//	github.com/0xProject/0x-mesh v9.4.0-beta+incompatible
	//	github.com/0xProject/0x-mesh v7.1.0-beta+incompatible
	//	github.com/0xProject/0x-mesh v9.2.0-incompatible
	//	github.com/0xProject/0x-mesh v7.1.0-beta+incompatible
	github.com/0xProject/0x-mesh v0.0.0-20200801025701-8123878dc210
	github.com/andersfylling/disgord v0.29.0 // indirect
	github.com/andersfylling/snowflake/v3 v3.0.2 // indirect
	github.com/ethereum/go-ethereum v1.9.23
	github.com/gin-gonic/autotls v0.0.0-20200518075542-45033372a9ad
	github.com/gin-gonic/gin v1.6.3
	github.com/gomodule/oauth1 v0.1.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/lib/pq v1.7.1
	github.com/libp2p/go-libp2p-peer v0.2.0 // indirect
	github.com/libp2p/go-libp2p-peerstore v0.2.6 // indirect
	github.com/multiformats/go-multiaddr v0.2.2 // indirect
	github.com/multiformats/go-multiaddr-dns v0.2.0 // indirect
	github.com/plaid/go-envvar v1.1.0
	github.com/sergi/go-diff v1.0.0 // indirect
	github.com/sirupsen/logrus v1.6.0 // indirect
	github.com/wealdtech/go-ens v1.0.0 // indirect
	github.com/wealdtech/go-ens/v3 v3.4.4-0.20200816153207-2be8e3e5fa10
	//	github.com/wealdtech/go-ens 2be8e3e5fa10e897e9957584302fae93d43d6cc1
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
)
