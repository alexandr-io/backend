package handlers

import (
	"time"

	"github.com/Alexandr-io/Backend/User/redis"
	"github.com/alexandr-io/backend_errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

// newRefreshToken creat and sign a new refresh jwt token.
// In case of error, the proper http error is set to the context and an empty string is returned.
func newRefreshToken(ctx *fiber.Ctx, userID string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = issuer                                          // Who created and signed the token
	claims["sub"] = string(ctx.Fasthttp.Request.Header.UserAgent()) // Whom the token refers to
	claims["aud"] = audience                                        // Who or what the token is intended for
	claims["user_id"] = userID                                      // Username of the auth user
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()      // Expire in 30 days

	secret := randomString(16)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		backend_errors.InternalServerError(ctx, err)
		return ""
	}
	if err := redis.StoreRefreshToken(ctx, signedToken, secret); err != nil {
		backend_errors.InternalServerError(ctx, err)
		return ""
	}
	return signedToken
}
