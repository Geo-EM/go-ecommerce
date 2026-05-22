package auth

import "e-commerce/internal/utils"

// TODO: Make it more secure
func GenerateVerificationCode() (uint, error) {
	return utils.GenerateRandomNumber(6)
}
