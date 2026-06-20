package auth

type TokenManager interface {
	GenerateAccessToken(userID string) (string, error)
	GenerateRefreshToken(UserID string) (string, error)
	ValidateAccessToken(token string) (Claims, error)
	ValidateRefreshToken(token string) (Claims, error)
}
