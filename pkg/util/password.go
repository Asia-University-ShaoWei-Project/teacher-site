package util

import (
	"context"
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

func GeneralSalt(ctx context.Context, saltSize int) []byte {
	var salt = make([]byte, saltSize)
	rand.Read(salt[:])
	return salt
}

func GeneralHashPassword(ctx context.Context, password string, salt []byte, hashCost int) string {
	saltPassword := append([]byte(password), salt...)
	hashPassword, _ := bcrypt.GenerateFromPassword(saltPassword, hashCost)
	return string(hashPassword)
}
