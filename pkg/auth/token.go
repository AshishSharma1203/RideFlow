package auth

type TokenManager interface {
	TokenGenerator
	TokenValidator
}
type TokenGenerator interface {
	GenerateAccessToken(userID string) (string, error)
	GenerateRefreshToken(UserID string) (string, error)
}

type TokenValidator interface {
	ValidateAccessToken(token string) (Claims, error)
	ValidateRefreshToken(token string) (Claims, error)
}
