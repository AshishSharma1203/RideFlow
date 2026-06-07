package app

import (
	"context"

	echo "github.com/labstack/echo/v4"

	"github.com/ashishSharma1203/rideflow/services/gateway/internal/client"
	"github.com/ashishSharma1203/rideflow/services/gateway/internal/handler"
	"github.com/ashishSharma1203/rideflow/services/gateway/internal/service"
)

type App struct {
	Echo *echo.Echo

	IdentityClient *client.IdentityClient

	HealthService *service.HealthService

	HealthHandler *handler.HealthHandler
}

func New() (*App, error) {
	e := echo.New()
	identityClient,err:=client.NewIdentityClient()
	if(err!=nil){
		return nil,err;
	}
	healthService:=service.NewHealthService(identityClient)
	healthHandler:=handler.NewHealthHandler(healthService)
	  app := &App{
        Echo: e,

        IdentityClient: identityClient,

        HealthService: healthService,

        HealthHandler: healthHandler,
    }

    app.registerRoutes()

    return app, nil
}

func (a *App) Start() error {
    return a.Echo.Start(":8080")
}

func (a *App) Shutdown(
    ctx context.Context,
) error {

    if err := a.Echo.Shutdown(ctx); err != nil {
        return err
    }

    return a.IdentityClient.Conn.Close()
}