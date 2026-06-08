package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ashishSharma1203/rideflow/services/gateway/internal/app"
)

func main() {

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := app.Start(); err != nil &&
			err != http.ErrServerClosed {

			log.Fatalf("echo start error: %v", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel :=
		context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
	defer cancel()

	app.Shutdown(shutdownCtx)
}
