package jwt

import (
	"github.com/alexandr-io/backend/auth/data"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// ExtractJWTFromHeader extract a jwt from the header.
func ExtractJWTFromHeader(ctx *fiber.Ctx) (string, error) {
	auth := string(ctx.Request().Header.Peek("Authorization"))

	l := len("Bearer")
	if len(auth) > l+1 && auth[:l] == "Bearer" {
		return auth[l+1:], nil
	}
	return "", data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Missing or malformed JWT")
}

// ExtractJWTFromContext extract a jwt from the context.
func ExtractJWTFromContext(ctx *fiber.Ctx) (*jwt.Token, error) {
	token, ok := ctx.Locals("jwt").(*jwt.Token)
	if !ok {
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Error casting locals jwt to *jwt.Token")
	}
	return token, nil
}

// ExtractJWTClaims extract the map of claims contained in the JWT of the given context.
func ExtractJWTClaims(ctx *fiber.Ctx) (jwt.MapClaims, error) {
	token, err := ExtractJWTFromContext(ctx)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Error casting token.Claims to jwt.MapClaims")
	}
	return claims, nil
}

// ExtractJWTUserID extract the user_id contained in the claims of the JWT of the given context.
func ExtractJWTUserID(ctx *fiber.Ctx) (string, error) {
	claims, err := ExtractJWTClaims(ctx)
	if err != nil {
		return "", err
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Error casting claims[\"user_id\"] to string")
	}
	return userID, nil
}
