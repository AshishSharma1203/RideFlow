package main

import (
	"log"
	"net"

	identityv1 "github.com/ashishSharma1203/rideflow/api/gen/identity/v1"

	"github.com/ashishSharma1203/rideflow/services/identity/internal/security/bcrypt"
	"github.com/ashishSharma1203/rideflow/services/identity/internal/service"
	transportgrpc "github.com/ashishSharma1203/rideflow/services/identity/internal/transport/grpc"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}

	server := grpc.NewServer()

	passwordHasher := bcrypt.NewHasher()

	identityService :=
		service.NewIdentityService(
			passwordHasher,
		)

	identityv1.RegisterIdentityServiceServer(
		server,
		transportgrpc.NewServer(identityService),
	)

	log.Println("identity service listening on :50051")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("serve error: %v", err)
	}
}
