package database

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/media/data"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// CollectionBook is the name of the book collection in mongodb
const CollectionBook = "book"

//
// Getters
//

// FindOneWithFilter fill the given object with a mongodb single result filtered by the given filters.
func FindOneWithFilter(object interface{}, filters interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := Instance.Db.Collection(CollectionBook)
	filteredSingleResult := collection.FindOne(ctx, filters)
	return filteredSingleResult.Decode(object)
}

// GetBookByID retrieve a book from the ID given by the library MS.
func GetBookByID(book *data.Book) (*data.Book, error) {
	filter := bson.D{{Key: "book_id", Value: book.ID}}
	err := FindOneWithFilter(book, filter)
	if err != nil {
		return book, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "You do not have access to this book.")
	}

	return book, nil
}

//
// Setters
//

// InsertBook create a book on the database
func InsertBook(book data.Book) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bookCollection := Instance.Db.Collection(CollectionBook)
	insertedResult, err := bookCollection.InsertOne(ctx, book)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return insertedResult, nil
}
