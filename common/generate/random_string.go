package generate

import (
	"math/rand"
	"time"
)

const charset2 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890123456789"

// RandomStringNoSpecialChar generate a random string of the given length with the charters in charset2.
func RandomStringNoSpecialChar(length int) string {
	seededRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset2[seededRand.Intn(len(charset2))]
	}
	return string(b)
}
