package app

import middleware "github.com/ashishSharma1203/rideflow/pkg/middlerware"

func (a *App) registerRoutes() {

	// Public routes
	a.Echo.GET(
		"/health",
		a.HealthHandler.Health,
	)

	a.Echo.POST(
		"/users/register",
		a.IdentityHandler.RegisterUser,
	)
	a.Echo.POST(
		"/users/login",
		a.IdentityHandler.LoginUser,
	)

	// protected routes
	protected := a.Echo.Group(
		"/api/v1",
		middleware.Auth(a.TokenValidator),
	)

	protected.GET(
		"/profile",
		a.ProfileHandler.GetProfile,
	)
}
