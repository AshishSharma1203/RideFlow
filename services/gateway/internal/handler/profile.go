package handler

import (
	"net/http"

	middleware "github.com/ashishSharma1203/rideflow/pkg/middlerware"
	"github.com/labstack/echo/v4"
)

type ProfileResponse struct {
	UserID string `json:"user_id"`
}

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
		ProfileResponse{
			UserID: claims.UserID,
		},
	)
}
