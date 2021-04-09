package book

import (
	"context"

	"github.com/alexandr-io/backend/media/data"
	"github.com/alexandr-io/backend/media/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Update take a data.Book to update a book in database
func Update(ctx context.Context, book data.Book) (*data.Book, error) {
	bookCollection := database.Instance.Db.Collection(database.CollectionBook)

	if err := bookCollection.FindOneAndUpdate(
		ctx,
		bson.D{
			{"book_id", book.ID},
			{"library_id", book.LibraryID},
		},
		bson.D{
			{"$set", book},
		},
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&book); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Invitation not found.")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &book, nil
}
