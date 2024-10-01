package internal

import (
	"app/app/domain"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

func GenerateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	claims := &domain.JwtCustomClaims{
		Email: user.Email,
		ID:    strconv.FormatInt(user.ID, 10),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "AMONIC Airlines",
			Subject:   "Authentication",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry))),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func GenerateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	claims := &domain.JwtCustomRefreshClaims{
		ID: strconv.FormatInt(user.ID, 10),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "AMONIC Airlines",
			Subject:   "Authentication",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry))),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}
