package handlers

import (
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
	return signedToken
}
