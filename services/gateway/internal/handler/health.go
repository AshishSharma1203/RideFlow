package handler

import (
	"context"
	"net/http"

	identityv1 "github.com/ashishSharma1203/rideflow/api/gen/identity/v1"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	identityClient identityv1.IdentityServiceClient
}

func NewHealthHandler(
	identityClient identityv1.IdentityServiceClient,
) *HealthHandler {
	return &HealthHandler{
		identityClient: identityClient,
	}
}

func (h *HealthHandler) Health(c echo.Context) error {

	resp, err := h.identityClient.HealthCheck(
		context.Background(),
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