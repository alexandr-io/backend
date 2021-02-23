package internal

import (
	"github.com/alexandr-io/backend/auth/data"
	authJWT "github.com/alexandr-io/backend/auth/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

// Auth is the internal logic function used to check if a JWT is valid and if the user exist.
func Auth(jwt string) (*data.User, error) {
	tokenObject, err := authJWT.Validate(jwt)
	if err != nil {
		return nil, err
	}

	// get user from refresh token
	fiberApp := fiber.New()
	ctx := fiberApp.AcquireCtx(&fasthttp.RequestCtx{})
	ctx.Locals("jwt", tokenObject)
	user, err := authJWT.GetUserFromContextJWT(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
