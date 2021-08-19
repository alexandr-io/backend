package database

import (
	"testing"

	"github.com/alexandr-io/backend/common/typeconv"
	"github.com/alexandr-io/backend/library/data"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestBookCreate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		BookDB = &BookCollection{collection: mt.Coll}
		book := data.Book{
			ID:                  primitive.NewObjectID(),
			LibraryID:           primitive.NewObjectID(),
			UploaderID:          primitive.NewObjectID(),
			Title:               "1984",
			Author:              "George Orwell",
			Description:         "The Party told you to reject the evidence of your eyes and ears. It was their final, most essential command.",
			Categories:          []string{"dystopian fiction", "utopian fiction", "political fiction", "social science fiction"},
			Thumbnails:          nil,
			PublishedDate:       "1949",
			MaturityRating:      "false",
			Language:            "en",
			IndustryIdentifiers: nil,
			PageCount:           328,
		}

		mt.AddMockResponses(mtest.CreateSuccessResponse(bson.E{Key: "id", Value: book.ID}))
		bookResponse, err := BookDB.Create(book)
		assert.Nil(t, err)
		assert.Equal(t, &book, bookResponse)
	})

	mt.Run("error", func(mt *mtest.T) {
		BookDB = &BookCollection{collection: mt.Coll}
		mt.AddMockResponses(bson.D{{"ok", 0}})
		book, err := BookDB.Create(data.Book{})
		assert.NotNil(t, err)
		assert.Nil(t, book)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})
}

func TestBookReadOneBook(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		BookDB = &BookCollection{collection: mt.Coll}
		expectedBook := data.Book{
			ID:                  primitive.NewObjectID(),
			LibraryID:           primitive.NewObjectID(),
			UploaderID:          primitive.NewObjectID(),
			Title:               "1984",
			Author:              "George Orwell",
			Description:         "The Party told you to reject the evidence of your eyes and ears. It was their final, most essential command.",
			Categories:          []string{"dystopian fiction", "utopian fiction", "political fiction", "social science fiction"},
			Thumbnails:          nil,
			PublishedDate:       "1949",
			MaturityRating:      "false",
			Language:            "en",
			IndustryIdentifiers: nil,
			PageCount:           328,
		}
		bsonD, err := typeconv.ToDoc(expectedBook)
		assert.Nil(t, err)

		find := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bsonD)
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(find, killCursors)

		book, err := BookDB.Read(bson.D{{"_id", expectedBook.ID}})
		assert.Nil(t, err)
		assert.Equal(t, &[]data.Book{expectedBook}, book)
	})

	mt.Run("no document", func(mt *mtest.T) {
		BookDB = &BookCollection{collection: mt.Coll}
		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"cursor", bson.D{
				{"id", int64(0)},
				{"ns", "foo.bar"},
				{string(mtest.FirstBatch), bson.A{}},
			}},
		})
		book, err := BookDB.Read(bson.D{{"_id", primitive.NewObjectID()}})
		assert.NotNil(t, err)
		assert.Nil(t, book)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusNotFound, e.Code)
	})
}

func TestBookReadBooks(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		BookDB = &BookCollection{collection: mt.Coll}
		libraryID := primitive.NewObjectID()

		expectedBooks := []data.Book{
			{
				ID:        primitive.NewObjectID(),
				LibraryID: libraryID,
				Title:     "1984",
				Author:    "George Orwell",
			},
			{
				ID:        primitive.NewObjectID(),
				LibraryID: libraryID,
				Title:     "Animal Farm",
				Author:    "George Orwell",
			},
		}
		bsonD1, err := typeconv.ToDoc(expectedBooks[0])
		assert.Nil(t, err)
		bsonD2, err := typeconv.ToDoc(expectedBooks[1])
		assert.Nil(t, err)

		// Create mock responses.
		find := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bsonD1)
		getMore := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, bsonD2)
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(find, getMore, killCursors)
		books, err := BookDB.Read(bson.D{{"library_id", libraryID}})
		assert.Nil(t, err)
		assert.Equal(t, &expectedBooks, books)
	})

	mt.Run("find error", func(mt *mtest.T) {
		BookDB = &BookCollection{collection: mt.Coll}

		books, err := BookDB.Read(bson.D{{"library_id", primitive.NewObjectID()}})
		assert.NotNil(t, err)
		assert.Nil(t, books)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})

	mt.Run("cursor read all error", func(mt *mtest.T) {
		BookDB = &BookCollection{collection: mt.Coll}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{}))
		books, err := BookDB.Read(bson.D{{"library_id", primitive.NewObjectID()}})
		assert.NotNil(t, err)
		assert.Nil(t, books)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})
}

func TestBookUpdate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		BookDB = &BookCollection{collection: mt.Coll}
		expectedBook := data.Book{
			ID:                  primitive.NewObjectID(),
			LibraryID:           primitive.NewObjectID(),
			UploaderID:          primitive.NewObjectID(),
			Title:               "1984",
			Author:              "George Orwell",
			Description:         "The Party told you to reject the evidence of your eyes and ears. It was their final, most essential command.",
			Categories:          []string{"dystopian fiction", "utopian fiction", "political fiction", "social science fiction"},
			Thumbnails:          nil,
			PublishedDate:       "1949",
			MaturityRating:      "false",
			Language:            "en",
			IndustryIdentifiers: nil,
			PageCount:           328,
		}
		bsonD, err := typeconv.ToDoc(expectedBook)
		assert.Nil(t, err)

		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"value", bsonD},
		})
		book, err := BookDB.Update(expectedBook)
		assert.Nil(t, err)
		assert.Equal(t, &expectedBook, book)
	})
	mt.Run("error", func(mt *mtest.T) {
		BookDB = &BookCollection{collection: mt.Coll}
		book, err := BookDB.Update(data.Book{})
		assert.NotNil(t, err)
		assert.Nil(t, book)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})
	mt.Run("no document", func(mt *mtest.T) {
		BookDB = &BookCollection{collection: mt.Coll}
		sentBook := data.Book{
			ID:                  primitive.NewObjectID(),
			LibraryID:           primitive.NewObjectID(),
			UploaderID:          primitive.NewObjectID(),
			Title:               "1984",
			Author:              "George Orwell",
			Description:         "The Party told you to reject the evidence of your eyes and ears. It was their final, most essential command.",
			Categories:          []string{"dystopian fiction", "utopian fiction", "political fiction", "social science fiction"},
			Thumbnails:          nil,
			PublishedDate:       "1949",
			MaturityRating:      "false",
			Language:            "en",
			IndustryIdentifiers: nil,
			PageCount:           328,
		}

		mt.AddMockResponses(bson.D{
			{"ok", 1},
		})
		book, err := BookDB.Update(sentBook)
		assert.NotNil(t, err)
		assert.Nil(t, book)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusNotFound, e.Code)
	})
}

func TestBookDelete(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		BookDB = &BookCollection{collection: mt.Coll}
		//database.BookProgressCollection = mt.Coll
		id := primitive.NewObjectID()
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 1}})
		//mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 1}})
		err := BookDB.Delete(id)
		assert.Nil(t, err)
	})

	mt.Run("not found", func(mt *mtest.T) {
		BookDB = &BookCollection{collection: mt.Coll}
		id := primitive.NewObjectID()
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 0}})
		err := BookDB.Delete(id)
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusNotFound, e.Code)
	})

	mt.Run("error", func(mt *mtest.T) {
		BookDB = &BookCollection{collection: mt.Coll}
		id := primitive.NewObjectID()
		mt.AddMockResponses(bson.D{{"ok", 0}})
		err := BookDB.Delete(id)
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})
}
