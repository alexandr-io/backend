package bookprogress

import (
	"testing"
	"time"

	"github.com/alexandr-io/backend/common/typeconv"
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"

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
			database.BookProgressCollection = mt.Coll
			expectedBookProgress := data.BookProgressData{
				UserID:       primitive.NewObjectID(),
				BookID:       primitive.NewObjectID(),
				LibraryID:    primitive.NewObjectID(),
				Progress:     "42",
				LastReadDate: time.Time{},
			}
			bsonD, err := typeconv.ToDoc(expectedBookProgress)
			assert.Nil(t, err)

			mt.AddMockResponses(bson.D{
				{"ok", 1},
				{"value", bsonD},
			})
			bookProgress, err := Upsert(expectedBookProgress)
			assert.Nil(t, err)
			assert.Equal(t, &expectedBookProgress, bookProgress)
		})
		mt.Run("error", func(t *mtest.T) {
			database.BookProgressCollection = mt.Coll
			bookProgress, err := Upsert(data.BookProgressData{})
			assert.NotNil(t, err)
			assert.Nil(t, bookProgress)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
		mt.Run("no document", func(t *mtest.T) {
			database.BookCollection = mt.Coll
			expectedBookProgress := data.BookProgressData{
				UserID:       primitive.NewObjectID(),
				BookID:       primitive.NewObjectID(),
				LibraryID:    primitive.NewObjectID(),
				Progress:     "42",
				LastReadDate: time.Time{},
			}
			bsonD, err := typeconv.ToDoc(expectedBookProgress)
			assert.Nil(t, err)

			mt.AddMockResponses(bson.D{
				{"ok", 1},
				{"value", bsonD},
			})
			bookProgress, err := Upsert(expectedBookProgress)
			assert.Nil(t, err)
			assert.Equal(t, &expectedBookProgress, bookProgress)
		})
	})
}
