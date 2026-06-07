package handler

import (
	"net/http"

	identityv1 "github.com/ashishSharma1203/rideflow/api/gen/identity/v1"
	"github.com/ashishSharma1203/rideflow/services/gateway/internal/client"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	identityClient *client.IdentityClient
}

func NewHealthHandler(
	identityClient *client.IdentityClient,
) *HealthHandler {
	return &HealthHandler{
		identityClient: identityClient,
	}
}

func (h *HealthHandler) Health(c echo.Context) error {

	resp, err := h.identityClient.Client.HealthCheck(
		c.Request().Context(),
		&identityv1.HealthCheckRequest{},
	)

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		map[string]string{
			"status": resp.Status,
		},
	)
}