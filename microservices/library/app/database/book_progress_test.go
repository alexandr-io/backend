package database

import (
	"testing"
	"time"

	"github.com/alexandr-io/backend/common/typeconv"
	"github.com/alexandr-io/backend/library/data"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestProgressUpdate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		BookProgressDB = &BookProgressCollection{collection: mt.Coll}
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
		bookProgress, err := BookProgressDB.Upsert(expectedBookProgress)
		assert.Nil(t, err)
		assert.Equal(t, &expectedBookProgress, bookProgress)
	})

	mt.Run("error", func(mt *mtest.T) {
		BookProgressDB = &BookProgressCollection{collection: mt.Coll}
		bookProgress, err := BookProgressDB.Upsert(data.BookProgressData{})
		assert.NotNil(t, err)
		assert.Nil(t, bookProgress)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})

	mt.Run("no document", func(mt *mtest.T) {
		BookProgressDB = &BookProgressCollection{collection: mt.Coll}
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
		bookProgress, err := BookProgressDB.Upsert(expectedBookProgress)
		assert.Nil(t, err)
		assert.Equal(t, &expectedBookProgress, bookProgress)
	})
}

func TestProgressRead(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		BookProgressDB = &BookProgressCollection{collection: mt.Coll}
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
		bookProgress, err := BookProgressDB.ReadFromIDs(expectedBookProgress.UserID, expectedBookProgress.BookID, expectedBookProgress.LibraryID)
		assert.Nil(t, err)
		assert.Equal(t, &expectedBookProgress, bookProgress)
	})

	mt.Run("error", func(mt *mtest.T) {
		BookProgressDB = &BookProgressCollection{collection: mt.Coll}

		book, err := BookProgressDB.ReadFromIDs(primitive.NilObjectID, primitive.NilObjectID, primitive.NilObjectID)
		assert.NotNil(t, err)
		assert.Nil(t, book)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})

	mt.Run("no document", func(mt *mtest.T) {
		BookProgressDB = &BookProgressCollection{collection: mt.Coll}

		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"cursor", bson.D{
				{"id", int64(0)},
				{"ns", "foo.bar"},
				{string(mtest.FirstBatch), bson.A{}},
			}},
		})
		invitation, err := BookProgressDB.ReadFromIDs(primitive.NilObjectID, primitive.NilObjectID, primitive.NilObjectID)
		assert.NotNil(t, err)
		assert.Nil(t, invitation)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusNotFound, e.Code)
	})
}

func TestProgressDelete(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		BookProgressDB = &BookProgressCollection{collection: mt.Coll}

		sentBookProgress := data.BookProgressData{
			UserID:       primitive.NewObjectID(),
			BookID:       primitive.NewObjectID(),
			LibraryID:    primitive.NewObjectID(),
			Progress:     "42",
			LastReadDate: time.Time{},
		}

		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 1}})
		err := BookProgressDB.Delete(sentBookProgress)
		assert.Nil(t, err)
	})

	mt.Run("not found", func(mt *mtest.T) {
		BookProgressDB = &BookProgressCollection{collection: mt.Coll}

		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 0}})
		err := BookProgressDB.Delete(data.BookProgressData{})
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusNotFound, e.Code)
	})

	mt.Run("error", func(mt *mtest.T) {
		BookProgressDB = &BookProgressCollection{collection: mt.Coll}

		mt.AddMockResponses(bson.D{{"ok", 0}})
		err := BookProgressDB.Delete(data.BookProgressData{})
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})
}
