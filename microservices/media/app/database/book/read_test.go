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

func TestRead(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.RunOpts("from book id", mtest.NewOptions().ClientType(mtest.Mock), func(mt *mtest.T) {
		database.Instance.Db = mt.DB
		mt.Run("success", func(mt *mtest.T) {
			database.BookCollection = mt.Coll
			expectedBook := data.Book{
				ID:        primitive.NewObjectID(),
				LibraryID: primitive.NewObjectID(),
				Path:      "path/to/book",
				CoverPath: "path/to/cover",
			}
			bytes, err := bson.Marshal(expectedBook)
			assert.Nil(t, err)
			var bsonD bson.D
			err = bson.Unmarshal(bytes, &bsonD)
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

			book, err := GetFromID(sentBook.ID)
			assert.NotNil(t, err)
			assert.Nil(t, book)

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
			book, err := GetFromID(primitive.NewObjectID())
			assert.NotNil(t, err)
			assert.Nil(t, book)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusNotFound, e.Code)
		})
	})
}
