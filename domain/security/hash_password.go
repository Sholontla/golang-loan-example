package security

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func SetPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Panic("Error ")
	}
	return string(hashedPassword), nil
}

func ComparePassword(s string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(s), []byte(password))
}
