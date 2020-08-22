package handlers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// hashAndSalt hash and salt a given password.
func hashAndSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// comparePasswords compare a hashed password with a plain string password.
func comparePasswords(hashedPassword string, plainPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
