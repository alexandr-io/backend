package middleware

import (
	"fmt"
	"os"

	"github.com/alexandr-io/backend/auth/data"
	authJWT "github.com/alexandr-io/backend/auth/jwt"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protected is the middleware to protect routes with jwt
func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: jwtError,
		Filter: func(ctx *fiber.Ctx) bool {
			jwt, err := authJWT.ExtractJWTFromHeader(ctx)
			if err != nil {
				errorInfo, _ := data.ErrorInfoUnmarshal(err.Error())
				fmt.Println("Middleware:", errorInfo.Message)
				return true
			}
			if _, err := authJWT.Validate(jwt); err != nil {
				errorInfo, _ := data.ErrorInfoUnmarshal(err.Error())
				fmt.Println("Middleware:", errorInfo.Message)
				return true
			}
			return false
		},
		ContextKey: "jwt",
	})
}

// jwtError manage errors for the jwt middleware
func jwtError(_ *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}
	return data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Invalid or expired JWT")
}
