package main

import (
	"fmt"
	"log"
	"net"

	identityv1 "github.com/ashishSharma1203/rideflow/api/gen/identity/v1"

	"github.com/ashishSharma1203/rideflow/services/identity/internal/database"
	"github.com/ashishSharma1203/rideflow/services/identity/internal/config"
	"github.com/ashishSharma1203/rideflow/services/identity/internal/security/bcrypt"
	"github.com/ashishSharma1203/rideflow/services/identity/internal/service"
	transportgrpc "github.com/ashishSharma1203/rideflow/services/identity/internal/transport/grpc"

	"google.golang.org/grpc"
)

func main() {

	// 1. Load and validate your configuration layout
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to initialize configs: %v", err)
	}

	db, err := database.New(cfg.Postgres)
	if err != nil {
		log.Fatalf("failed to initialize postgres: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("postgres close error: %v", err)
		}
	}()

	// 2. Format the int port into a valid address string (e.g., ":50051")
	addr := fmt.Sprintf(":%d", cfg.Server.GRPCPort)

	// 3. Pass the string address to net.Listen
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen error on address %s: %v", addr, err)
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
