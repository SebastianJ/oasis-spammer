package rpc

import (
	"fmt"

	cmnGrpc "github.com/oasislabs/oasis-core/go/common/grpc"
	consensus "github.com/oasislabs/oasis-core/go/consensus/api"
	"google.golang.org/grpc"
)

// ConsensusClient - initiate a new consensus client
func ConsensusClient(address string) (*grpc.ClientConn, consensus.ClientBackend, error) {
	conn, err := Connect(address)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to establish connection with node %s", address)
	}

	client := consensus.NewConsensusClient(conn)
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
