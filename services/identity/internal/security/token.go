package security

type Claims struct {
	UserID string `json:"user_id"`
	Expiry int64  `json:"expiry"`
}

type TokenManager interface {
	GenerateAccessToken(
		userID string,
	) (string, error)

	GenerateRefreshToken(
		userID string,
	) (string, error)

	ValidateAccessToken(
		token string,
	) (Claims, error)

	ValidateRefreshToken(
		token string,
	) (Claims, error)
}
