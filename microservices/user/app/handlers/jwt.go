package handlers

import (
	"log"

	"github.com/alexandr-io/backend/user/database"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// extractJWTFromContext extract a jwt from the context.
func extractJWTFromContext(ctx *fiber.Ctx) (*jwt.Token, bool) {
	token, ok := ctx.Locals("jwt").(*jwt.Token)
	if !ok {
		log.Println("Error casting locals jwt to *jwt.Token")
		return nil, false
	}
	return token, true
}

// extractJWTClaims extract the map of claims contained in the JWT of the given context.
func extractJWTClaims(ctx *fiber.Ctx) (jwt.MapClaims, bool) {
	token, ok := extractJWTFromContext(ctx)
	if !ok {
		return nil, false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Error casting token.Claims to jwt.MapClaims")
		return nil, false
	}
	return claims, true
}

// extractJWTUserID extract the user_id contained in the claims of the JWT of the given context.
func extractJWTUserID(ctx *fiber.Ctx) (string, bool) {
	claims, ok := extractJWTClaims(ctx)
	if !ok {
		return "", false
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		log.Println("Error casting claims[\"user_id\"] to string")
		return "", false
	}
	return userID, true
}

// ExtractJWTUsername extract the username from the user_id contained in the claims of the JWT of the given context.
func ExtractJWTUsername(ctx *fiber.Ctx) (string, bool) {
	userID, ok := extractJWTUserID(ctx)
	if !ok {
		return "", ok
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return "", false
	}
	user, err := database.GetUserByID(userObjectID)
	if err != nil {
		return "", false
	}
	return user.Username, true
}
