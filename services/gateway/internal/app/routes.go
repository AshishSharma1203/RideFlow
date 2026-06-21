package app

import middleware "github.com/ashishSharma1203/rideflow/pkg/middlerware"

func (a *App) registerRoutes() {

	// Public routes
	//  health check remains unversioned as it's used by load balancers and monitoring tools that expect a simple endpoint

	a.Echo.GET(
		"/health",
		a.HealthHandler.Health,
	)

	v1 := a.Echo.Group("/api/v1")

	v1.POST(
		"/users/register",
		a.IdentityHandler.RegisterUser,
	)
	v1.POST(
		"/users/login",
		a.IdentityHandler.LoginUser,
	)

	// protected routes guarded by authentication middleware
	protected := v1.Group(
		"",
		middleware.Auth(a.TokenValidator),
	)

	protected.GET(
		"/profile",
		a.ProfileHandler.GetProfile,
	)
}
