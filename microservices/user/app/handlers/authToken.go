package handlers

import (
	"log"
	"os"
	"time"

	"github.com/alexandr-io/berrors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

const (
	issuer   = "alexandrio_user_service"
	audience = "alexandrio_backend"
)

// newGlobalAuthToken creat and sign a new global jwt auth token for connection.
// In case of error, the proper http error is set to the context and an empty string is returned.
func newGlobalAuthToken(ctx *fiber.Ctx, userID string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = issuer                                          // Who created and signed the token
	claims["sub"] = string(ctx.Fasthttp.Request.Header.UserAgent()) // Whom the token refers to
	claims["aud"] = audience                                        // Who or what the token is intended for
	claims["user_id"] = userID                                      // User ID of the auth user
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()         // Expire in 15 minutes

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		berrors.InternalServerError(ctx, err)
		return ""
	}
	log.Println("New JWT for " + userID + ": " + signedToken)
	return signedToken
}
