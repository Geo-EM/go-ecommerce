package auth

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	Secret   string
	Issuer   string
	TokenTTL time.Duration
}

func NewTokenService(secret, issuer string, ttl time.Duration) (*TokenService, error) {
	if secret == "" {
		return nil, errors.New("jwt secret required")
	}

	if issuer == "" {
		return nil, errors.New("jwt issuer required")
	}

	if ttl <= 0 {
		return nil, errors.New("token ttl must be positive")
	}

	return &TokenService{
		Secret:   secret,
		Issuer:   issuer,
		TokenTTL: ttl,
	}, nil
}

func (t TokenService) GenerateToken(userID uint, email, role string) (string, error) {

	if userID == 0 || email == "" || role == "" {
		return "", errors.New("invalid user data")
	}

	now := time.Now()

	claims := JwtClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.Itoa(int(userID)),
			Issuer:    t.Issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(t.TokenTTL)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &claims)

	return token.SignedString([]byte(t.Secret))
}

func (t TokenService) ValidateToken(tokenString string) (*JwtClaims, error) {

	if tokenString == "" {
		return nil, errors.New("token is empty")
	}

	claims := &JwtClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(t.Secret), nil
		},
	)

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.Issuer != t.Issuer {
		return nil, errors.New("invalid token issuer")
	}

	return claims, nil
}
