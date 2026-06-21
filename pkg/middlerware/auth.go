package middleware

import (
	"net/http"
	"strings"

	"github.com/ashishSharma1203/rideflow/pkg/auth"
	echo "github.com/labstack/echo/v4"
)

// ClaimsContextKey is the single source of truth for where validated claims are
// stored in the Echo context. Echo's Set/Get are string-keyed, so centralizing
// the key in one constant avoids stringly-typed bugs across packages.
const ClaimsContextKey = "claims"

// Auth returns Echo middleware that authenticates a request using a bearer
// access token.
//
// It depends ONLY on auth.TokenValidator: by construction this middleware can
// verify tokens but can never mint them (least privilege, enforced by the type
// system). The validator is captured via closure — idiomatic dependency
// injection for middleware.
func Auth(validator auth.TokenValidator) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header.Get(echo.HeaderAuthorization)
			if header == "" {
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"missing authorization header",
				)
			}

			// Expect exactly: "Bearer <token>". The scheme is case-insensitive
			// per RFC 6750.
			parts := strings.SplitN(header, " ", 2)
			if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"invalid authorization header",
				)
			}

			token := strings.TrimSpace(parts[1])
			if token == "" {
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"missing bearer token",
				)
			}

			claims, err := validator.ValidateAccessToken(token)
			if err != nil {
				// Every validation failure collapses to a generic 401. We
				// deliberately do not leak err.Error() to the caller — that
				// would expose internals (signing method, parsing details).
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"invalid or expired token",
				)
			}

			c.Set(ClaimsContextKey, claims)

			return next(c)
		}
	}
}

// ClaimsFromContext safely retrieves the claims that Auth stored. Any handler
// running behind Auth can rely on ok == true; handlers reachable without Auth
// should check ok before trusting the result.
func ClaimsFromContext(c echo.Context) (auth.Claims, bool) {
	claims, ok := c.Get(ClaimsContextKey).(auth.Claims)
	return claims, ok
}
