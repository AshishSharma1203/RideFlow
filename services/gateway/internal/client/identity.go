package client

import (
	identityv1 "github.com/ashishSharma1203/rideflow/api/gen/identity/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IdentityClient struct {
	Client identityv1.IdentityServiceClient
	Conn   *grpc.ClientConn
}

func NewIdentityClient(add string) (*IdentityClient, error) {

	conn, err := grpc.NewClient(
		add,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		return nil, err
	}

	return &IdentityClient{
		Client: identityv1.NewIdentityServiceClient(conn),
		Conn:   conn,
	}, nil
}
