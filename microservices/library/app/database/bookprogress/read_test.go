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

func TestRead(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.RunOpts("from id", mtest.NewOptions().ClientType(mtest.Mock), func(mt *mtest.T) {
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

			mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bsonD))
			bookProgress, err := RetrieveFromIDs(expectedBookProgress.UserID, expectedBookProgress.BookID, expectedBookProgress.LibraryID)
			assert.Nil(t, err)
			assert.Equal(t, &expectedBookProgress, bookProgress)
		})
		mt.Run("error", func(t *mtest.T) {
			database.BookProgressCollection = mt.Coll

			user, err := RetrieveFromIDs(primitive.NilObjectID, primitive.NilObjectID, primitive.NilObjectID)
			assert.NotNil(t, err)
			assert.Nil(t, user)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
		mt.Run("no document", func(t *mtest.T) {
			database.BookProgressCollection = mt.Coll

			mt.AddMockResponses(bson.D{
				{"ok", 1},
				{"cursor", bson.D{
					{"id", int64(0)},
					{"ns", "foo.bar"},
					{string(mtest.FirstBatch), bson.A{}},
				}},
			})
			invitation, err := RetrieveFromIDs(primitive.NilObjectID, primitive.NilObjectID, primitive.NilObjectID)
			assert.NotNil(t, err)
			assert.Nil(t, invitation)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusNotFound, e.Code)
		})
	})
}
