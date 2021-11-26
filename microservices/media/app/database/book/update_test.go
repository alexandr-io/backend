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

func TestUpdate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.RunOpts("find one and update", mtest.NewOptions().ClientType(mtest.Mock), func(mt *mtest.T) {
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

			mt.AddMockResponses(bson.D{
				{"ok", 1},
				{"value", bsonD},
			})
			book, err := Update(expectedBook)
			assert.Nil(t, err)
			assert.Equal(t, &expectedBook, book)
		})
		mt.Run("error", func(t *mtest.T) {
			database.BookCollection = mt.Coll
			book, err := Update(data.Book{})
			assert.NotNil(t, err)
			assert.Nil(t, book)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
		mt.Run("no document", func(t *mtest.T) {
			database.BookCollection = mt.Coll
			sentBook := data.Book{
				ID:        primitive.NewObjectID(),
				LibraryID: primitive.NewObjectID(),
				Path:      "path/to/book",
				CoverPath: "path/to/cover",
			}

			mt.AddMockResponses(bson.D{
				{"ok", 1},
			})
			book, err := Update(sentBook)
			assert.NotNil(t, err)
			assert.Nil(t, book)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusNotFound, e.Code)
		})
	})
}
