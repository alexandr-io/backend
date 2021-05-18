package book

import (
	"testing"

	"github.com/alexandr-io/backend/media/data"
	"github.com/alexandr-io/backend/media/database"

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
			id := primitive.NewObjectID()
			libraryID := primitive.NewObjectID()
			book := data.Book{
				ID:        id,
				LibraryID: libraryID,
				Path:      "path/to/book",
				CoverPath: "path/to/cover",
			}

			mt.AddMockResponses(mtest.CreateSuccessResponse(bson.E{Key: "id", Value: id}))
			bookResponse, err := Insert(book)
			assert.Nil(t, err)
			assert.Equal(t, &book, bookResponse)
		})
		mt.Run("error", func(t *mtest.T) {
			database.BookCollection = mt.Coll
			mt.AddMockResponses(bson.D{{"ok", 0}})
			invitation, err := Insert(data.Book{})
			assert.NotNil(t, err)
			assert.Nil(t, invitation)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
		mt.Run("error duplicate book id", func(t *mtest.T) {
			database.BookCollection = mt.Coll
			mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
				Index:   1,
				Code:    11000,
				Message: "duplicate key error",
			}))
			mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{}))
			invitation, err := Insert(data.Book{})
			assert.NotNil(t, err)
			assert.Nil(t, invitation)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
	})
}
