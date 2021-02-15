package bookprogress

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/gofiber/fiber/v2"
)

// Delete deletes the user's book progress data entry in the database
func Delete(ctx context.Context, bookUserData data.BookProgressData) error {
	collection := database.Instance.Db.Collection(database.CollectionBookProgress)

	if result, err := collection.DeleteOne(ctx, bookUserData); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	} else if result.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusNotFound, "Progress not found.")
	}

	return nil
}
