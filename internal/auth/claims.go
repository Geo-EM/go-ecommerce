package auth

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	UserID uint   `json:"userId"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// Ensure JwtClaims implements jwt.Claims interface at compile time
var _ jwt.Claims = (*JwtClaims)(nil)
