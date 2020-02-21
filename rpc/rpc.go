package rpc

import (
	"context"
	"fmt"
	"os"

	"github.com/oasislabs/oasis-core/go/common/crypto/signature"
	cmnGrpc "github.com/oasislabs/oasis-core/go/common/grpc"
	consensus "github.com/oasislabs/oasis-core/go/consensus/api"
	"github.com/oasislabs/oasis-core/go/staking/api"
	"google.golang.org/grpc"
)

// CurrentNonce - get the current nonce for a given signer
func CurrentNonce(signer signature.Signer, address string) (uint64, error) {
	conn, client, err := StakingClient(address)
	defer conn.Close()

	if err != nil {
		return 0, err
	}

	ctx := context.Background()

	var acct *api.Account
	acct, err = client.AccountInfo(ctx, &api.OwnerQuery{Owner: signer.Public(), Height: consensus.HeightLatest})
	if err != nil {
		return 0, err
	}

	fmt.Printf("Acct: %+v\n", acct)

	os.Exit(-1)

	return 0, nil
}

// ConsensusClient - initiate a new consensus client
func ConsensusClient(address string) (*grpc.ClientConn, consensus.ClientBackend, error) {
	conn, err := Connect(address)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to establish connection with node %s", address)
	}

	client := consensus.NewConsensusClient(conn)
	return conn, client, nil
}

// StakingClient - initiate a new staking client
func StakingClient(address string) (*grpc.ClientConn, api.Backend, error) {
	conn, err := Connect(address)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to establish connection with node %s", address)
	}

	client := api.NewStakingClient(conn)
	return conn, client, nil
}

// Connect - connect to grpc
func Connect(address string) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.WaitForReady(false)))

	conn, err := cmnGrpc.Dial(
		address,
		opts...,
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
