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

type RegisterUserResponse struct {
	UserId string `json:"user_id"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
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
			ErrorResponse{
				Error: "invalid request body",
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
			ErrorResponse{
				Error: userFacingError(err),
			},
		)
	}

	return c.JSON(
		http.StatusCreated,
		RegisterUserResponse{
			UserId: res.UserID,
		},
	)
}
func (h *UserHandler) LoginUser(c echo.Context) error {
	var req LoginUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			ErrorResponse{
				Error: "invalid request body",
			},
		)
	}

	res, err := h.identityService.LoginUser(c.Request().Context(), service.LoginUserInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return c.JSON(
			httpStatusFromError(err),
			ErrorResponse{
				Error: userFacingError(err),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		LoginUserResponse{
			UserID:       res.UserID,
			Username:     res.Username,
			Email:        res.Email,
			AccessToken:  res.AccessToken,
			RefreshToken: res.RefreshToken,
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
