package auth

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	Secret string
	Issuer string
}

func (tokenService TokenService) GenerateToken(userID uint, email, role string) (string, error) {
	if userID == 0 || email == "" || role == "" {
		return "", errors.New("invalid user data")
	}

	if tokenService.Secret == "" || tokenService.Issuer == "" {
		return "", errors.New("token service not configured")
	}

	now := time.Now()

	claims := JwtClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.Itoa(int(userID)),
			Issuer:    tokenService.Issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(7 * 24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &claims)

	return token.SignedString([]byte(tokenService.Secret))
}

func (tokenService TokenService) ValidateToken(tokenString string) (*JwtClaims, error) {

	if tokenString == "" {
		return nil, errors.New("token is empty")
	}

	tokenStrings := strings.Split(tokenString, " ")
	if len(tokenStrings) != 2 || tokenStrings[0] != "Bearer" {
		return nil, errors.New("invalid token format")
	}

	tokenString = tokenStrings[1]

	claims := &JwtClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(tokenService.Secret), nil
		},
	)

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.Issuer != tokenService.Issuer {
		return nil, errors.New("invalid token issuer")
	}

	return claims, nil
}
