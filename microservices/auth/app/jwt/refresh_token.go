package jwt

import (
	"time"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/redis"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// newRefreshToken creat and sign a new refresh jwt token.
// In case of error, the proper http error is set to the context and an empty string is returned.
func newRefreshToken(ctx *fiber.Ctx, userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = issuer                                           // Who created and signed the token
	claims["sub"] = string(ctx.Context().Request.Header.UserAgent()) // Whom the token refers to
	claims["aud"] = audience                                         // Who or what the token is intended for
	claims["user_id"] = userID                                       // User ID of the auth user
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()       // Expire in 30 days

	secret := RandomString(16)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if err = redis.RefreshToken.Create(ctx.Context(), signedToken, secret); err != nil {
		return "", err
	}
	return signedToken, nil
}
