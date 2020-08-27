package handlers

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/Alexandr-io/Backend/User/database"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"

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
		fmt.Println("Error casting claims[\"user_id\"] to string")
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

func parseJWT(ctx *fiber.Ctx, token string, secret string) (*jwt.Token, bool) {
	tokenObject, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		log.Println(err)
		_ = ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired JWT"})
		return nil, false
	}
	return tokenObject, true
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
