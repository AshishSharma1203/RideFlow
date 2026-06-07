package handler

import (
	"net/http"

	"github.com/ashishSharma1203/rideflow/services/gateway/internal/service"
	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	healthService *service.HealthService
}

func NewHealthHandler(
	healthService *service.HealthService,
) *HealthHandler {
	return &HealthHandler{
		healthService: healthService,
	}
}

func (h *HealthHandler) Health(c echo.Context) error {
	status, err := h.healthService.Check(c.Request().Context())
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
			"status": status,
		},
	)
}