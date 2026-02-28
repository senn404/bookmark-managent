package service

import (
	"bytes"
	"crypto/rand"
	"math/big"
)

const (
	charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"
	passLength = 16
)

type passwordService struct {
}

type Password interface {
	GeneratePassword() (string, error)
}

func NewPassword() Password {
	return &passwordService{}
}

// GeneratePassword generates a random password of length 16 using a cryptographically secure random number generator.
func (s *passwordService) GeneratePassword() (string, error) {
	var strBuilder bytes.Buffer

	for i := 0; i < passLength; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}

		strBuilder.WriteByte(charset[randomIndex.Int64()])
	}

	return strBuilder.String(), nil
}
