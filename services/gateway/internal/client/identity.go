package client

import (
	identityv1 "github.com/ashishSharma1203/rideflow/api/gen/identity/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewIdentityClient() (identityv1.IdentityServiceClient, error) {

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		return nil, err
	}

	return identityv1.NewIdentityServiceClient(conn), nil
}