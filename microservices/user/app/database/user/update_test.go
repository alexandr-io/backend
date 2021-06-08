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

func TestUpdate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.RunOpts("find one and update", mtest.NewOptions().ClientType(mtest.Mock), func(mt *mtest.T) {
		database.Instance.Db = mt.DB
		mt.Run("success", func(mt *mtest.T) {
			database.UserCollection = mt.Coll
			expectedUser := data.User{
				ID:            primitive.NewObjectID(),
				Username:      "test",
				Email:         "test",
				EmailVerified: true,
				Password:      "42",
			}
			bytes, err := bson.Marshal(expectedUser)
			assert.Nil(t, err)
			var bsonD bson.D
			err = bson.Unmarshal(bytes, &bsonD)
			assert.Nil(t, err)

			mt.AddMockResponses(bson.D{
				{"ok", 1},
				{"value", bsonD},
			})
			user, err := Update(expectedUser.ID, expectedUser)
			assert.Nil(t, err)
			assert.Equal(t, &expectedUser, user)
		})
		mt.Run("error", func(t *mtest.T) {
			database.UserCollection = mt.Coll
			user, err := Update(primitive.NewObjectID(), data.User{})
			assert.NotNil(t, err)
			assert.Nil(t, user)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
		mt.Run("no document", func(t *mtest.T) {
			database.UserCollection = mt.Coll
			sentUser := data.User{
				ID:            primitive.NewObjectID(),
				Username:      "test",
				Email:         "test",
				EmailVerified: true,
				Password:      "42",
			}

			mt.AddMockResponses(bson.D{
				{"ok", 1},
			})
			user, err := Update(sentUser.ID, sentUser)
			assert.NotNil(t, err)
			assert.Nil(t, user)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusNotFound, e.Code)
		})
	})
}
