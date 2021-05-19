package book

import (
	"testing"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.RunOpts("insert one", mtest.NewOptions().ClientType(mtest.Mock), func(mt *mtest.T) {
		database.Instance.Db = mt.DB
		mt.Run("success", func(mt *mtest.T) {
			database.BookCollection = mt.Coll
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
			bookResponse, err := Insert(book)
			assert.Nil(t, err)
			assert.Equal(t, &book, bookResponse)
		})
		mt.Run("error", func(t *mtest.T) {
			database.BookCollection = mt.Coll
			mt.AddMockResponses(bson.D{{"ok", 0}})
			book, err := Insert(data.Book{})
			assert.NotNil(t, err)
			assert.Nil(t, book)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
	})
}
