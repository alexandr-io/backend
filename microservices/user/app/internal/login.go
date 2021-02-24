package internal

import (
	"log"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database/user"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Login is the internal logic function used to login a user.
func Login(login string, password string) (*data.User, error) {
	// Get the user from it's login
	userData, err := user.FromLogin(login)
	if err != nil {
		return nil, err
	}

	// Check the user's password
	if !comparePasswords(userData.Password, []byte(password)) {
		return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "User not found")
	}

	return userData, nil
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
