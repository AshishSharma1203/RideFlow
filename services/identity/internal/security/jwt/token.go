package jwt

import (
	"time"

	"github.com/ashishSharma1203/rideflow/services/identity/internal/security"
	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secretKey              string
	accessTokenExpiration  time.Duration
	refreshTokenExpiration time.Duration
}

func NewJWTMaker(
	secretKey string,
	accessTokenExpiration time.Duration,
	refreshTokenExpiration time.Duration,
) *JWTMaker {
	return &JWTMaker{
		secretKey:              secretKey,
		accessTokenExpiration:  accessTokenExpiration,
		refreshTokenExpiration: refreshTokenExpiration,
	}
}

func (m *JWTMaker) GenerateAccessToken(userID string) (string, error) {
	return m.generateToken(userID, m.accessTokenExpiration)
}

func (m *JWTMaker) GenerateRefreshToken(userID string) (string, error) {
	return m.generateToken(userID, m.refreshTokenExpiration)
}

func (m *JWTMaker) generateToken(
	userID string,
	duration time.Duration,
) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secretKey))
}

func (m *JWTMaker) ValidateAccessToken(token string) (security.Claims, error) {
	return m.validateToken(token)
}

func (m *JWTMaker) ValidateRefreshToken(token string) (security.Claims, error) {
	return m.validateToken(token)
}

func (m *JWTMaker) validateToken(token string) (security.Claims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, security.ErrInvalidToken
		}
		return []byte(m.secretKey), nil
	})
	if err != nil {
		return security.Claims{}, err
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		userID, ok := claims["user_id"].(string)
		if !ok {
			return security.Claims{}, security.ErrInvalidToken
		}
		expiry, ok := claims["exp"].(float64)
		if !ok {
			return security.Claims{}, security.ErrInvalidToken
		}
		return security.Claims{
			UserID: userID,
			Expiry: int64(expiry),
		}, nil
	}
	return security.Claims{}, security.ErrInvalidToken
}
