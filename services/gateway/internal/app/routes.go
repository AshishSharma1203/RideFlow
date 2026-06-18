package app

func (a *App) registerRoutes() {

	a.Echo.GET(
		"/health",
		a.HealthHandler.Health,
	)

	a.Echo.POST(
		"/register",
		a.IdentityHandler.RegisterUser,
	)
}
