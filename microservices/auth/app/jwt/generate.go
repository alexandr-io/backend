package jwt

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GenerateNewRefreshTokenAndAuthToken generate both the auth and refresh token.
// The first returned string is the refresh token and the second one is the auth token.
func GenerateNewRefreshTokenAndAuthToken(
	ctx *fiber.Ctx, userID string) (string, string, error) {

	refreshToken, err := newRefreshToken(ctx, userID)
	if err != nil {
		return "", "", err
	}
	authToken, err := newGlobalAuthToken(ctx, userID)
	if err != nil {
		return "", "", err
	}
	return refreshToken, authToken, nil
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890123456789!@#$%^&*()_+<>?:\"|{}[]'."

// RandomString generate a random string of the given length with the charters in charset.
func RandomString(length int) string {
	seededRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
