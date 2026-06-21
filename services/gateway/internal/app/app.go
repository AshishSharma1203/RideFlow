package app

import (
	"context"
	"fmt"

	echo "github.com/labstack/echo/v4"

	authjwt "github.com/ashishSharma1203/rideflow/pkg/auth/jwt"

	"github.com/ashishSharma1203/rideflow/pkg/auth"
	"github.com/ashishSharma1203/rideflow/services/gateway/internal/client"
	"github.com/ashishSharma1203/rideflow/services/gateway/internal/config"
	"github.com/ashishSharma1203/rideflow/services/gateway/internal/handler"
	"github.com/ashishSharma1203/rideflow/services/gateway/internal/service"
)

type App struct {
	Echo   *echo.Echo
	Config *config.Config

	IdentityClient *client.IdentityClient
	TokenValidator auth.TokenValidator

	HealthService   *service.HealthService
	IdentityService *service.IdentityService

	HealthHandler   *handler.HealthHandler
	IdentityHandler *handler.UserHandler
	ProfileHandler  *handler.ProfileHandler
}

func New(cfg *config.Config) (*App, error) {
	e := echo.New()

	identityClient, err := client.NewIdentityClient(cfg.Identity.GRPCAddr)
	if err != nil {
		return nil, err
	}
	tokenValidator := authjwt.NewTokenValidator(cfg.JWT.SecretKey)
	healthService := service.NewHealthService(identityClient)
	identityService := service.NewIdentityService(identityClient)
	healthHandler := handler.NewHealthHandler(healthService)
	identityHandler := handler.NewIdentityHandler(identityService)
	profileHandler := handler.NewProfileHandler()

	app := &App{
		Echo:           e,
		Config:         cfg,
		IdentityClient: identityClient,

		TokenValidator: tokenValidator,

		HealthService:   healthService,
		IdentityService: identityService,

		HealthHandler:   healthHandler,
		IdentityHandler: identityHandler,
		ProfileHandler:  profileHandler,
	}

	app.registerRoutes()

	return app, nil
}

func (a *App) Start() error {
	// return a.Echo.Start(":8080")
	addr := fmt.Sprintf(":%d", a.Config.Server.HTTPPort)
	return a.Echo.Start(addr)
}

func (a *App) Shutdown(
	ctx context.Context,
) error {

	if err := a.Echo.Shutdown(ctx); err != nil {
		return err
	}

	return a.IdentityClient.Conn.Close()
}
