package utils

import (
	"crypto/rand"
	"strconv"
)

func GenerateRandomNumber(length uint) (uint, error) {

	const numbers = "0123456789"

	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return 0, err
	}

	for i := range b {
		b[i] = numbers[int(b[i])%len(numbers)]
	}

	result, err := strconv.Atoi(string(b))
	if err != nil {
		return 0, err
	}
	return uint(result), nil
}
