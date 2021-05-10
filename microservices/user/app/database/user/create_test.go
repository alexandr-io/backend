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

func TestCreate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.RunOpts("find one", mtest.NewOptions().ClientType(mtest.Mock), func(mt *mtest.T) {
		database.Instance.Db = mt.DB
		mt.Run("success", func(mt *mtest.T) {
			database.UserCollection = mt.Coll
			id := primitive.NewObjectID()
			user := data.User{
				ID:            id,
				Username:      "test",
				Email:         "test@test.test",
				EmailVerified: false,
				Password:      "test",
			}

			mt.AddMockResponses(mtest.CreateSuccessResponse(bson.E{Key: "id", Value: id}))
			userResponse, err := Insert(user)
			assert.Nil(t, err)
			assert.Equal(t, &user, userResponse)
		})
		mt.Run("error", func(t *mtest.T) {
			database.UserCollection = mt.Coll
			mt.AddMockResponses(bson.D{{"ok", 0}})
			invitation, err := Insert(data.User{})
			assert.NotNil(t, err)
			assert.Nil(t, invitation)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
		mt.Run("error duplicate email", func(t *mtest.T) {
			database.UserCollection = mt.Coll
			mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
				Index:   1,
				Code:    11000,
				Message: "duplicate key error",
			}))
			mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{}))
			invitation, err := Insert(data.User{})
			assert.NotNil(t, err)
			assert.Nil(t, invitation)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusBadRequest, e.Code)
		})
		mt.Run("error duplicate username", func(t *mtest.T) {
			mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
				Index:   1,
				Code:    11000,
				Message: "duplicate key error",
			}))
			mt.AddMockResponses(bson.D{
				{"ok", 1},
				{"cursor", bson.D{
					{"id", int64(0)},
					{"ns", "foo.bar"},
					{string(mtest.FirstBatch), bson.A{}},
				}},
			})
			mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{{"email", "test@test.test"}}))
			invitation, err := Insert(data.User{})
			assert.NotNil(t, err)
			assert.Nil(t, invitation)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusBadRequest, e.Code)
		})
		mt.Run("error duplicate but nothing found", func(t *mtest.T) {
			database.UserCollection = mt.Coll
			mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
				Index:   1,
				Code:    11000,
				Message: "duplicate key error",
			}))
			invitation, err := Insert(data.User{})
			assert.NotNil(t, err)
			assert.Nil(t, invitation)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
	})
}
