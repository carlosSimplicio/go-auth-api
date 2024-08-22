package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password []byte) (hashedPassword []byte, err error) {
	rounds := 10
	return bcrypt.GenerateFromPassword(password, rounds)
}

func ComparePassword(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}
