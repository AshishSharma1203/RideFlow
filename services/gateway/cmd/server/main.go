package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/ashishSharma1203/rideflow/services/gateway/internal/client"
	"github.com/ashishSharma1203/rideflow/services/gateway/internal/handler"
)

func main() {

	e := echo.New()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	defer stop()

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

	go func() {

		if err := e.Start(":8080"); err != nil &&
			err != http.ErrServerClosed {

			log.Fatalf("echo start error: %v", err)
		}

	}()

	<-ctx.Done()

	log.Println("shutdown signal received")
	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Printf("echo shutdown error: %v", err)
	}
	if err := identityClient.Conn.Close(); err != nil {
		log.Printf(
			"identity grpc close error: %v",
			err,
		)
	}

	log.Println("gateway stopped")
}
