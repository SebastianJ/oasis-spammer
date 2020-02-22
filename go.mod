module github.com/SebastianJ/oasis-spammer

go 1.13

replace (
	github.com/tendermint/iavl => github.com/oasislabs/iavl v0.12.0-ekiden3
	github.com/tendermint/tendermint => github.com/oasislabs/tendermint v0.32.8-oasis2
	golang.org/x/crypto/curve25519 => github.com/oasislabs/ed25519/extra/x25519 v0.0.0-20191022155220-a426dcc8ad5f
	golang.org/x/crypto/ed25519 => github.com/oasislabs/ed25519 v0.0.0-20191109133925-b197a691e30d
)

require (
	github.com/fxamacker/cbor v1.3.2
	github.com/oasislabs/oasis-core/go v0.0.0-20200210160800-cef4925f3105
	github.com/prometheus/client_golang v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.6.1 // indirect
	github.com/urfave/cli v1.22.2
	golang.org/x/net v0.0.0-20190923162816-aa69164e4478 // indirect
	golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect
	google.golang.org/grpc v1.25.1
	gopkg.in/yaml.v2 v2.2.7
)
