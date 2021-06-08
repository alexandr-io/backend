package bookprogress

import (
	"testing"
	"time"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestDelete(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.RunOpts("delete one", mtest.NewOptions().ClientType(mtest.Mock).CreateClient(true), func(mt *mtest.T) {
		database.Instance.Db = mt.DB
		mt.Run("success", func(mt *mtest.T) {
			database.BookProgressCollection = mt.Coll

			sentBookProgress := data.BookProgressData{
				UserID:       primitive.NewObjectID(),
				BookID:       primitive.NewObjectID(),
				LibraryID:    primitive.NewObjectID(),
				Progress:     "42",
				LastReadDate: time.Time{},
			}

			mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 1}})
			err := Delete(sentBookProgress)
			assert.Nil(t, err)
		})
		mt.Run("not found", func(t *mtest.T) {
			database.BookProgressCollection = mt.Coll

			mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 0}})
			err := Delete(data.BookProgressData{})
			assert.NotNil(t, err)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusNotFound, e.Code)
		})
		mt.Run("error", func(t *mtest.T) {
			database.BookProgressCollection = mt.Coll

			mt.AddMockResponses(bson.D{{"ok", 0}})
			err := Delete(data.BookProgressData{})
			assert.NotNil(t, err)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
	})
}
