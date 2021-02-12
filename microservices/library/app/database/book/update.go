package book

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
	"github.com/gofiber/fiber/v2"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Update update a book in a library.
func Update(DBBook data.BookData) (*data.BookData, error) {
	// Update data
	invitationCollection := database.Instance.Db.Collection(database.CollectionBook)
	err := invitationCollection.FindOneAndUpdate(
		context.Background(),
		bson.D{
			{"_id", DBBook.ID},
		},
		bson.D{{"$set", DBBook}},
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&DBBook)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return &DBBook, nil
}
