package database

import (
	"context"
	"time"

	"github.com/alexandr-io/backend/library/data"
	bson2 "github.com/globalsign/mgo/bson"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//
// Getters
//

// BookRetrieve search and return the metadata of a book on the mongo database
func BookRetrieve(ctx context.Context, bookRetrieve data.BookRetrieve) (data.Book, error) {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	collection := Instance.Db.Collection(CollectionLibrary)

	id, err := primitive.ObjectIDFromHex(bookRetrieve.LibraryID)
	if err != nil {
		return data.Book{}, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	library := &data.Library{}
	libraryFilter := bson.D{{Key: "_id", Value: id}}

	filterOptions := options.FindOne().SetProjection(bson.D{{"books", true}})
	filteredByLibraryIDSingleResult := collection.FindOne(ctx, libraryFilter, filterOptions)
	if err := filteredByLibraryIDSingleResult.Decode(library); err != nil {
		return data.Book{}, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	for _, book := range library.Books {
		if book.ID.Hex() == bookRetrieve.ID {
			return book, nil
		}
	}

	return data.Book{}, data.NewHTTPErrorInfo(fiber.StatusUnauthorized, "Book matching query does not exist")
}

//
// Setters
//

// BookCreate create the metadata of a book on the mongo database
func BookCreate(ctx context.Context, bookCreation data.BookCreation) (data.Book, error) {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	collection := Instance.Db.Collection(CollectionLibrary)

	book := data.Book{
		Title:       bookCreation.Title,
		Author:      bookCreation.Author,
		Publisher:   bookCreation.Publisher,
		Description: bookCreation.Description,
		Tags:        bookCreation.Tags,
		UploaderID:  bookCreation.UploaderID,
	}

	generatedID, err := primitive.ObjectIDFromHex(bson2.NewObjectId().Hex())
	if err != nil {
		return book, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	book.ID = generatedID

	id, err := primitive.ObjectIDFromHex(bookCreation.LibraryID)
	if err != nil {
		return data.Book{}, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	library := &data.Library{}
	libraryFilter := bson.D{{Key: "_id", Value: id}}

	filterOptions := options.FindOne().SetProjection(bson.D{{"books", true}})
	filteredByLibraryIDSingleResult := collection.FindOne(ctx, libraryFilter, filterOptions)
	if err := filteredByLibraryIDSingleResult.Decode(library); err != nil {
		return data.Book{}, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	library.Books = append(library.Books, book)

	updateValues := bson.D{{Key: "$set", Value: bson.D{{Key: "books", Value: library.Books}}}}
	updateResult := collection.FindOneAndUpdate(ctx, libraryFilter, updateValues)
	if updateResult.Err() != nil {
		return data.Book{}, data.NewHTTPErrorInfo(fiber.StatusInternalServerError, updateResult.Err().Error())
	}

	_, err = UserDataCreate(ctx, data.UserData{UserID: bookCreation.UploaderID})
	if err != nil {
		return book, data.NewHTTPErrorInfo(fiber.StatusPartialContent, "Failed to create user data")
	}

	return book, nil
}

// BookUpdate update the metadata of a book
func BookUpdate(ctx context.Context, libraryIDStr string, book data.Book) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	libraryID, err := primitive.ObjectIDFromHex(libraryIDStr)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	bookFilter := bson.D{
		{"_id", libraryID},
		{"books._id", book.ID},
	}

	collection := Instance.Db.Collection(CollectionLibrary)
	_, err = collection.UpdateOne(ctx, bookFilter, bson.D{{"$set", bson.D{{"books.$", book}}}})
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	return nil

	/*
		cursor, err := collection.Aggregate(ctx, mongo.Pipeline{
			bson.D{{"$match", bookFilter}},
			bson.D{{"$replaceRoot", bson.D{{"newRoot", bson.D{{"$first", "$books"}}}}}},
		})
		if err != nil {
			return err
		}
		var showsLoaded data.Book

		if !cursor.Next(ctx) {
			return data.NewHTTPErrorInfo(fiber.StatusNotFound, "The resource you requested does not exist.")
		}
		if err = cursor.Decode(&showsLoaded); err != nil {
			return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
		}
	*/
}

// BookDelete delete the metadata of a book on the mongo database
func BookDelete(ctx context.Context, bookRetrieve data.BookRetrieve) error {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	collection := Instance.Db.Collection(CollectionLibrary)

	id, err := primitive.ObjectIDFromHex(bookRetrieve.LibraryID)
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	library := &data.Library{}
	libraryFilter := bson.D{{Key: "_id", Value: id}}

	filterOptions := options.FindOne().SetProjection(bson.D{{"books", true}})
	filteredByLibraryIDSingleResult := collection.FindOne(ctx, libraryFilter, filterOptions)
	if err := filteredByLibraryIDSingleResult.Decode(library); err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}

	var bookList []data.Book
	for _, book := range library.Books {
		if book.ID.Hex() != bookRetrieve.ID {
			bookList = append(bookList, book)
		}
	}

	if bookList == nil {
		bookList = []data.Book{}
	}
	updateValues := bson.D{{Key: "$set", Value: bson.D{{Key: "books", Value: bookList}}}}
	updateResult := collection.FindOneAndUpdate(ctx, libraryFilter, updateValues)
	if updateResult.Err() != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, updateResult.Err().Error())
	}

	return nil
}
