package jwt

import (
	"time"

	"github.com/ashishSharma1203/rideflow/pkg/auth"
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

func NewTokenValidator(secretKey string) *JWTMaker{
	return &JWTMaker{
		secretKey: secretKey,
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

func (m *JWTMaker) ValidateAccessToken(token string) (auth.Claims, error) {
	return m.validateToken(token)
}

func (m *JWTMaker) ValidateRefreshToken(token string) (auth.Claims, error) {
	return m.validateToken(token)
}

func (m *JWTMaker) validateToken(token string) (auth.Claims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, auth.ErrInvalidToken
		}
		return []byte(m.secretKey), nil
	})
	if err != nil {
		return auth.Claims{}, err
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		userID, ok := claims["user_id"].(string)
		if !ok {
			return auth.Claims{}, auth.ErrInvalidToken
		}
		expiry, ok := claims["exp"].(float64)
		if !ok {
			return auth.Claims{}, auth.ErrInvalidToken
		}
		return auth.Claims{
			UserID: userID,
			Expiry: int64(expiry),
		}, nil
	}
	return auth.Claims{}, auth.ErrInvalidToken
}
