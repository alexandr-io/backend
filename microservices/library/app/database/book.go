package database

import (
	"context"

	"github.com/alexandr-io/backend/library/data"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var BookDB *BookCollection

type BookCollection struct {
	collection *mongo.Collection
}

const bookCollectionName = "book"

func NewBookCollection(db *mongo.Database) *BookCollection {
	return &BookCollection{collection: db.Collection(bookCollectionName)}
}

// Create a book
func (c *BookCollection) Create(book data.Book) (*data.Book, error) {
	result, err := c.collection.InsertOne(context.Background(), book)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	book.ID = result.InsertedID.(primitive.ObjectID)
	return &book, nil
}

// Read books infos
func (c *BookCollection) Read(filter bson.D) (*[]data.Book, error) {
	var books []data.Book

	cursor, err := c.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	if err = cursor.All(context.Background(), &books); err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if len(books) == 0 {
		return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "No book found.")
	}

	return &books, nil
}

// Update a book
func (c *BookCollection) Update(book data.Book) (*data.Book, error) {
	if err := c.collection.FindOneAndUpdate(
		context.Background(),
		bson.D{
			{"_id", book.ID},
		},
		bson.D{{"$set", book}},
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&book); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, data.NewHTTPErrorInfo(fiber.StatusNotFound, "Book not found.")
		}
		return nil, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	return &book, nil
}

// Delete a book
func (c *BookCollection) Delete(id primitive.ObjectID) error {
	result, err := c.collection.DeleteOne(
		context.Background(),
		bson.D{
			{Key: "_id", Value: id},
		},
	)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if result.DeletedCount == 0 {
		return data.NewHTTPErrorInfo(fiber.StatusNotFound, "can't find book to delete")
	}
	return nil
}
