package auth

import "e-commerce/internal/utils"

func GenerateVerificationCode() (uint, error) {
	return utils.GenerateRandomNumber(6)
}
