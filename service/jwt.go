package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	Secret    string
	Issuer    string
	ExpiresAt string
}

type jwtClaims struct {
	UserID  uint `json:"user_id"`
	IsAdmin bool `json:"is_admin"`
	jwt.RegisteredClaims
}

type JWTTokenResponse struct {
	Token string `json:"token"`
}

func NewJWTService(secret string, issuer string, expiryTime string) *JWTService {
	return &JWTService{Secret: secret, Issuer: issuer, ExpiresAt: expiryTime}
}

// GenerateToken creates a JWT for a given user ID
func (j *JWTService) GenerateToken(userID uint, isAdmin bool) (string, error) {
	claims := &jwtClaims{
		UserID:  userID,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.Secret))
}

// ValidateToken parses & validates a token
func (j *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
