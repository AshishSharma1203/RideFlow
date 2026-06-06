package main

import (
	"log"

	"github.com/labstack/echo/v4"

	"github.com/ashishSharma1203/rideflow/services/gateway/internal/client"
	"github.com/ashishSharma1203/rideflow/services/gateway/internal/handler"
)

func main() {

	e := echo.New()

	identityClient, err := client.NewIdentityClient()
	if err != nil {
		log.Fatalf("identity client error: %v", err)
	}

	healthHandler := handler.NewHealthHandler(
		identityClient,
	)

	e.GET(
		"/health",
		healthHandler.Health,
	)

	log.Println("gateway listening on :8080")

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}