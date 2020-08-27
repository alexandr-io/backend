package handlers

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Alexandr-io/Backend/User/database"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

// ExtractJWTUserID extract the user_id contained in the claims of the JWT of the given context.
func ExtractJWTUserID(ctx *fiber.Ctx) (string, bool) {
	claims, ok := ExtractJWTClaims(ctx)
	if !ok {
		return "", false
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		fmt.Println("Error casting claims[\"user_id\"] to string")
		return "", false
	}
	return userID, true
}

// ExtractJWTUsername extract the username from the user_id contained in the claims of the JWT of the given context.
func ExtractJWTUsername(ctx *fiber.Ctx) (string, bool) {
	userID, ok := ExtractJWTUserID(ctx)
	if !ok {
		return "", ok
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return "", false
	}
	user, ok := database.GetUserByID(ctx, userObjectID)
	if !ok {
		return "", ok
	}
	return user.Username, true
}

// generateNewRefreshTokenAndAuthToken generate both the auth and refresh token.
func generateNewRefreshTokenAndAuthToken(
	ctx *fiber.Ctx, userID string) (refreshToken string, authToken string, ok bool) {
	ok = true
	refreshToken = newRefreshToken(ctx, userID)
	authToken = newGlobalAuthToken(ctx, userID)
	if refreshToken == "" || authToken == "" {
		ok = false
	}
	return
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890123456789!@#$%^&*()_+<>?:\"|{}[]'."

// randomString generate a random string of the given length with the charters in charset.
func randomString(length int) string {
	seededRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
