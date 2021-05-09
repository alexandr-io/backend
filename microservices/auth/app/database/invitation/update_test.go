package invitation

import (
	"testing"
	"time"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/database"

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
			expectedInvitation := data.Invitation{
				ID:     primitive.NewObjectID(),
				Token:  "dOG8UVzaLk",
				Used:   func() *time.Time { now := time.Now(); return &now }(),
				UserID: func() *primitive.ObjectID { id := primitive.NewObjectID(); return &id }(),
			}
			bytes, err := bson.Marshal(expectedInvitation)
			assert.Nil(t, err)
			var bsonD bson.D
			err = bson.Unmarshal(bytes, &bsonD)
			assert.Nil(t, err)

			mt.AddMockResponses(bson.D{
				{"ok", 1},
				{"value", bsonD},
			})
			invitation, err := Update(mt.Coll, expectedInvitation)
			assert.Nil(t, err)
			assert.Equal(t, &expectedInvitation, invitation)
		})
		mt.Run("error", func(t *mtest.T) {
			invitation, err := Update(mt.Coll, data.Invitation{})
			assert.NotNil(t, err)
			assert.Nil(t, invitation)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
		mt.Run("no document", func(t *mtest.T) {
			sentInvitation := data.Invitation{
				ID:     primitive.NewObjectID(),
				Token:  "dOG8UVzaLk",
				Used:   func() *time.Time { now := time.Now(); return &now }(),
				UserID: func() *primitive.ObjectID { id := primitive.NewObjectID(); return &id }(),
			}

			mt.AddMockResponses(bson.D{
				{"ok", 1},
			})
			invitation, err := Update(mt.Coll, sentInvitation)
			assert.NotNil(t, err)
			assert.Nil(t, invitation)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusNotFound, e.Code)
		})
	})
}
