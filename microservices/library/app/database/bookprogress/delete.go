package bookprogress

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
)

// Delete deletes the user's book progress data entry in the database
func Delete(bookUserData data.BookProgressData) error {
	if result, err := database.BookProgressCollection.DeleteOne(context.Background(), bookUserData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	} else if result.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusNotFound, "Progress not found.")
	}

	return nil
}
