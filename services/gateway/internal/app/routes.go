package app

func (a *App) registerRoutes() {

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
}
