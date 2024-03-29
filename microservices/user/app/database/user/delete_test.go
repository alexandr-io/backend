package user

import (
	"testing"

	"github.com/alexandr-io/backend/user/database"

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
			database.UserCollection = mt.Coll
			id := primitive.NewObjectID()
			mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 1}})
			err := Delete(id)
			assert.Nil(t, err)
		})
		mt.Run("not found", func(t *mtest.T) {
			database.UserCollection = mt.Coll
			id := primitive.NewObjectID()
			mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 0}})
			err := Delete(id)
			assert.NotNil(t, err)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusUnauthorized, e.Code)
		})
		mt.Run("error", func(t *mtest.T) {
			database.UserCollection = mt.Coll
			id := primitive.NewObjectID()
			mt.AddMockResponses(bson.D{{"ok", 0}})
			err := Delete(id)
			assert.NotNil(t, err)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
	})
}
