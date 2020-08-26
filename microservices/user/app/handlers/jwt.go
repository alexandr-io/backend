package handlers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alexandr-io/backend_errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

const (
	issuer   = "alexandrio_user_service"
	audience = "alexandrio_backend"
)

// newGlobalJWT creat and sign a new global jwt token for connection.
// In case of error, the proper http error is set to the context and an empty string is returned.
func newGlobalJWT(ctx *fiber.Ctx, username string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = issuer                                          // Who created and signed the token
	claims["sub"] = string(ctx.Fasthttp.Request.Header.UserAgent()) // Whom the token refers to
	claims["aud"] = audience                                        // Who or what the token is intended for
	claims["username"] = username                                   // Username of the auth user
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()           // Expire in 72 hours

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		backend_errors.InternalServerError(ctx, err)
		return ""
	}
	log.Println("New JWT for " + username + ": " + signedToken)
	return signedToken
}

// ExtractJWTClaims extract the map of claims contained in the JWT of the given context.
func ExtractJWTClaims(ctx *fiber.Ctx) (jwt.MapClaims, bool) {
	token, ok := ctx.Locals("jwt").(*jwt.Token)
	if !ok {
		log.Println("Error casting locals jwt to *jwt.Token")
		return nil, false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Error casting token.Claims to jwt.MapClaims")
		return nil, false
	}
	return claims, true
}

// ExtractJWTUsername extract the username contained in the claims of the JWT of the given context.
func ExtractJWTUsername(ctx *fiber.Ctx) (string, bool) {
	claims, ok := ExtractJWTClaims(ctx)
	if !ok {
		return "", false
	}
	username, ok := claims["username"].(string)
	if !ok {
		fmt.Println("Error casting claims[\"username\"] to string")
		return "", false
	}
	return username, true
}
