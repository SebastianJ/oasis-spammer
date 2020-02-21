module github.com/SebastianJ/oasis-spammer

go 1.13

replace (
	github.com/tendermint/iavl => github.com/oasislabs/iavl v0.12.0-ekiden3
	github.com/tendermint/tendermint => github.com/oasislabs/tendermint v0.32.8-oasis2
	golang.org/x/crypto/curve25519 => github.com/oasislabs/ed25519/extra/x25519 v0.0.0-20191022155220-a426dcc8ad5f
	golang.org/x/crypto/ed25519 => github.com/oasislabs/ed25519 v0.0.0-20191109133925-b197a691e30d
)

require (
	github.com/oasislabs/oasis-core/go v0.0.0-20200210160800-cef4925f3105
	github.com/urfave/cli v1.22.2
	google.golang.org/grpc v1.25.1
)
