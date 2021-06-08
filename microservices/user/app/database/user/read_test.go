package user

import (
	"testing"

	"github.com/alexandr-io/backend/user/data"
	"github.com/alexandr-io/backend/user/database"

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
			database.UserCollection = mt.Coll
			expectedUser := data.User{
				ID:            primitive.NewObjectID(),
				Username:      "test",
				Email:         "test@test.test",
				EmailVerified: true,
				Password:      "42",
			}
			bytes, err := bson.Marshal(expectedUser)
			assert.Nil(t, err)
			var bsonD bson.D
			err = bson.Unmarshal(bytes, &bsonD)
			assert.Nil(t, err)

			mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bsonD))
			user, err := FromID(expectedUser.ID)
			assert.Nil(t, err)
			assert.Equal(t, &expectedUser, user)
		})
		mt.Run("error", func(t *mtest.T) {
			database.UserCollection = mt.Coll
			sentUser := data.User{
				ID: primitive.NewObjectID(),
			}

			user, err := FromID(sentUser.ID)
			assert.NotNil(t, err)
			assert.Nil(t, user)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusUnauthorized, e.Code)
		})
	})

	mt.RunOpts("from email", mtest.NewOptions().ClientType(mtest.Mock), func(mt *mtest.T) {
		database.Instance.Db = mt.DB
		mt.Run("success", func(mt *mtest.T) {
			database.UserCollection = mt.Coll
			expectedUser := data.User{
				ID:            primitive.NewObjectID(),
				Username:      "test",
				Email:         "test@test.test",
				EmailVerified: true,
				Password:      "42",
			}
			bytes, err := bson.Marshal(expectedUser)
			assert.Nil(t, err)
			var bsonD bson.D
			err = bson.Unmarshal(bytes, &bsonD)
			assert.Nil(t, err)

			mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bsonD))
			user, err := FromEmail(expectedUser.Email)
			assert.Nil(t, err)
			assert.Equal(t, &expectedUser, user)
		})
		mt.Run("error", func(t *mtest.T) {
			database.UserCollection = mt.Coll
			sentUser := data.User{
				Email: "test@test.com",
			}

			user, err := FromEmail(sentUser.Email)
			assert.NotNil(t, err)
			assert.Nil(t, user)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
		mt.Run("no document", func(t *mtest.T) {
			database.UserCollection = mt.Coll
			var email = "test@test.com"

			mt.AddMockResponses(bson.D{
				{"ok", 1},
				{"cursor", bson.D{
					{"id", int64(0)},
					{"ns", "foo.bar"},
					{string(mtest.FirstBatch), bson.A{}},
				}},
			})
			user, err := FromEmail(email)
			assert.NotNil(t, err)
			assert.Nil(t, user)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusNotFound, e.Code)
		})
	})

	mt.RunOpts("from username", mtest.NewOptions().ClientType(mtest.Mock), func(mt *mtest.T) {
		database.Instance.Db = mt.DB
		mt.Run("success", func(mt *mtest.T) {
			database.UserCollection = mt.Coll
			expectedUser := data.User{
				ID:            primitive.NewObjectID(),
				Username:      "test",
				Email:         "test@test.test",
				EmailVerified: true,
				Password:      "42",
			}
			bytes, err := bson.Marshal(expectedUser)
			assert.Nil(t, err)
			var bsonD bson.D
			err = bson.Unmarshal(bytes, &bsonD)
			assert.Nil(t, err)

			mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bsonD))
			user, err := FromUsername(expectedUser.Email)
			assert.Nil(t, err)
			assert.Equal(t, &expectedUser, user)
		})
		mt.Run("error", func(t *mtest.T) {
			database.UserCollection = mt.Coll
			sentUser := data.User{
				Email: "test@test.com",
			}

			user, err := FromUsername(sentUser.Email)
			assert.NotNil(t, err)
			assert.Nil(t, user)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
		mt.Run("no document", func(t *mtest.T) {
			database.UserCollection = mt.Coll
			var email = "test@test.com"

			mt.AddMockResponses(bson.D{
				{"ok", 1},
				{"cursor", bson.D{
					{"id", int64(0)},
					{"ns", "foo.bar"},
					{string(mtest.FirstBatch), bson.A{}},
				}},
			})
			user, err := FromUsername(email)
			assert.NotNil(t, err)
			assert.Nil(t, user)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusNotFound, e.Code)
		})
	})

	mt.RunOpts("from login", mtest.NewOptions().ClientType(mtest.Mock), func(mt *mtest.T) {
		database.Instance.Db = mt.DB
		mt.Run("success email", func(mt *mtest.T) {
			database.UserCollection = mt.Coll
			expectedUser := data.User{
				ID:            primitive.NewObjectID(),
				Username:      "test",
				Email:         "test@test.test",
				EmailVerified: true,
				Password:      "42",
			}
			bytes, err := bson.Marshal(expectedUser)
			assert.Nil(t, err)
			var bsonD bson.D
			err = bson.Unmarshal(bytes, &bsonD)
			assert.Nil(t, err)

			mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bsonD))
			user, err := FromLogin(expectedUser.Email)
			assert.Nil(t, err)
			assert.Equal(t, &expectedUser, user)
		})
		mt.Run("success username", func(mt *mtest.T) {
			database.UserCollection = mt.Coll
			expectedUser := data.User{
				ID:            primitive.NewObjectID(),
				Username:      "test",
				Email:         "test@test.test",
				EmailVerified: true,
				Password:      "42",
			}
			bytes, err := bson.Marshal(expectedUser)
			assert.Nil(t, err)
			var bsonD bson.D
			err = bson.Unmarshal(bytes, &bsonD)
			assert.Nil(t, err)

			mt.AddMockResponses(bson.D{{"ok", 0}})
			mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bsonD))
			user, err := FromLogin(expectedUser.Email)
			assert.Nil(t, err)
			assert.Equal(t, &expectedUser, user)
		})
		mt.Run("error nothing found", func(mt *mtest.T) {
			database.UserCollection = mt.Coll
			expectedUser := data.User{
				ID:            primitive.NewObjectID(),
				Username:      "test",
				Email:         "test@test.test",
				EmailVerified: true,
				Password:      "42",
			}
			bytes, err := bson.Marshal(expectedUser)
			assert.Nil(t, err)
			var bsonD bson.D
			err = bson.Unmarshal(bytes, &bsonD)
			assert.Nil(t, err)

			mt.AddMockResponses(bson.D{{"ok", 0}})
			mt.AddMockResponses(bson.D{{"ok", 0}})
			user, err := FromLogin(expectedUser.Email)
			assert.NotNil(t, err)
			assert.Nil(t, user)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusUnauthorized, e.Code)
		})
	})

}
