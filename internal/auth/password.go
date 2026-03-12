package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func validatePasswordLength(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	if len(password) > 71 {
		return errors.New("password must be less than 71 bytes/characters")
	}

	return nil
}

func HashPassword(rawPassword string) (string, error) {

	if err := validatePasswordLength(rawPassword); err != nil {
		return "", err
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(rawPassword),
		bcrypt.DefaultCost,
	)

	return string(hash), err
}

func ValidatePassword(rawPassword, hashedPassword string) error {
	if err := validatePasswordLength(rawPassword); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword)); err != nil {
		return errors.New("invalid credentials")
	}

	return nil
}
