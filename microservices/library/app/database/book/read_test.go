package book

import (
	"github.com/alexandr-io/backend/common/typeconv"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"testing"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestRead(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.RunOpts("from id", mtest.NewOptions().ClientType(mtest.Mock), func(mt *mtest.T) {
		database.Instance.Db = mt.DB
		mt.Run("success", func(mt *mtest.T) {
			database.BookCollection = mt.Coll
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

			mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bsonD))
			book, err := GetFromID(expectedBook.ID)
			assert.Nil(t, err)
			assert.Equal(t, &expectedBook, book)
		})
		mt.Run("error", func(t *mtest.T) {
			database.BookCollection = mt.Coll
			sentBook := data.Book{
				ID: primitive.NewObjectID(),
			}

			user, err := GetFromID(sentBook.ID)
			assert.NotNil(t, err)
			assert.Nil(t, user)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
		mt.Run("no document", func(t *mtest.T) {
			database.BookCollection = mt.Coll
			mt.AddMockResponses(bson.D{
				{"ok", 1},
				{"cursor", bson.D{
					{"id", int64(0)},
					{"ns", "foo.bar"},
					{string(mtest.FirstBatch), bson.A{}},
				}},
			})
			user, err := GetFromID(primitive.NewObjectID())
			assert.NotNil(t, err)
			assert.Nil(t, user)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusNotFound, e.Code)
		})
	})

	mt.RunOpts("books from library id", mtest.NewOptions().ClientType(mtest.Mock), func(mt *mtest.T) {
		database.Instance.Db = mt.DB
		mt.Run("success", func(mt *mtest.T) {
			database.BookCollection = mt.Coll
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
			books, err := GetBooksFromLibraryID(libraryID)
			assert.Nil(t, err)
			assert.Equal(t, &expectedBooks, books)
		})
		mt.Run("find error", func(t *mtest.T) {
			database.BookCollection = mt.Coll

			user, err := GetBooksFromLibraryID(primitive.NewObjectID())
			assert.NotNil(t, err)
			assert.Nil(t, user)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
		mt.Run("cursor read all error", func(t *mtest.T) {
			database.BookCollection = mt.Coll

			mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{}))
			user, err := GetBooksFromLibraryID(primitive.NewObjectID())
			assert.NotNil(t, err)
			assert.Nil(t, user)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
	})
}
