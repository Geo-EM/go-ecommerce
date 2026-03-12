package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(rawPassword string) (string, error) {
	if len(rawPassword) < 8 {
		return "", errors.New("password must be at least 8 characters")
	}

	if len(rawPassword) > 72 {
		return "", errors.New("password must be less than 72 bytes")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(rawPassword),
		bcrypt.DefaultCost,
	)

	return string(hash), err
}

func ValidatePassword(rawPassword, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(rawPassword),
	) == nil
}
