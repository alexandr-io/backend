package handlers

import (
	"fmt"
	"time"

	"github.com/alexandr-io/backend/auth/data"
	authJWT "github.com/alexandr-io/backend/auth/jwt"
	"github.com/alexandr-io/backend/auth/redis"

	"github.com/gofiber/fiber/v2"
)

// Logout the connected user
func Logout(ctx *fiber.Ctx) error {
	// get jwt claims to get the expiration date
	claims, err := authJWT.ExtractJWTClaims(ctx)
	if err != nil {
		return err
	}
	expInt, ok := claims["exp"].(float64)
	if !ok {
		fmt.Println(claims["exp"])
		fmt.Printf("Type = %v", claims["exp"])
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, "Error casting claims[\"exp\"] to int64")
	}
	exp := time.Until(time.Unix(int64(expInt), 0))
	fmt.Println(exp.Minutes())
	// get jwt to store in redis
	jwt, err := authJWT.ExtractJWTFromContext(ctx)
	if err != nil {
		return err
	}

	if err := redis.StoreAuthTokenBlackList(ctx, jwt.Raw, exp); err != nil {
		return err
	}

	if err := ctx.SendStatus(fiber.StatusNoContent); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
