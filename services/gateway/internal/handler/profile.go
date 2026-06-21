package handler

import (
	"net/http"

	middleware "github.com/ashishSharma1203/rideflow/pkg/middlerware"
	"github.com/labstack/echo/v4"
)

type ProfileHandler struct {
}

func NewProfileHandler() *ProfileHandler {
	return &ProfileHandler{}
}

func (h *ProfileHandler) GetProfile(c echo.Context) error {
	// Implement the logic to get the profile
	claims, ok := middleware.ClaimsFromContext(c)

	if !ok {
		return echo.NewHTTPError(
			http.StatusUnauthorized,
			"missing claims",
		)
	}

	return c.JSON(
		http.StatusOK,
		map[string]string{
			"user_id": claims.UserID,
		},
	)
}
