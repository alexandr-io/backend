package jwt

import (
	"github.com/alexandr-io/backend/auth/data"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// ParseJWT return a jwt.Token from the given string using the given secret
func ParseJWT(token string, secret string) (*jwt.Token, error) {
	tokenObject, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, err.Error())
	}
	return tokenObject, nil
}
