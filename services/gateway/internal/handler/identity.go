package handler

import (
	"net/http"

	"github.com/ashishSharma1203/rideflow/services/gateway/internal/service"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegisterUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserHandler struct {
	identityService *service.IdentityService
}

func NewIdentityHandler(identityService *service.IdentityService) *UserHandler {
	return &UserHandler{
		identityService: identityService,
	}
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	var req RegisterUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": "invalid request body",
			},
		)
	}

	res, err := h.identityService.RegisterUser(c.Request().Context(), service.RegisterUserInput{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return c.JSON(
			httpStatusFromError(err),
			map[string]string{
				"error": userFacingError(err),
			},
		)
	}

	return c.JSON(
		http.StatusCreated,
		map[string]string{
			"user_id": res.UserID,
		},
	)
}

func httpStatusFromError(err error) int {
	statusErr, ok := status.FromError(err)
	if !ok {
		return http.StatusInternalServerError
	}

	switch statusErr.Code() {
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.AlreadyExists:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func userFacingError(err error) string {
	statusErr, ok := status.FromError(err)
	if !ok {
		return "internal server error"
	}

	switch statusErr.Code() {
	case codes.InvalidArgument, codes.AlreadyExists:
		return statusErr.Message()
	default:
		return "internal server error"
	}
}
