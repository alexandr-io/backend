package handlers

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

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
