package internal

import (
	"github.com/alexandr-io/backend/auth/data"
	authJWT "github.com/alexandr-io/backend/auth/jwt"
	"github.com/alexandr-io/backend/auth/kafka/producers"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

// Auth is the internal logic function used to check if a JWT is valid and if the user exist.
func Auth(key string, message data.KafkaAuthRequest) error {
	tokenObject, err := authJWT.Validate(message.JWT)
	if err != nil {
		return producers.SendErrorAuthMessage(key, err)
	}

	// get user from refresh token
	fiberApp := fiber.New()
	ctx := fiberApp.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Locals("jwt", tokenObject)
	user, err := authJWT.GetUserFromContextJWT(ctx)
	if err != nil {
		return producers.SendErrorAuthMessage(key, err)
	}

	return producers.SendAuthResponseMessage(key, fiber.StatusOK, user)
}
