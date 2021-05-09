package invitation

import (
	"testing"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/auth/database"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestRead(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.RunOpts("read one", mtest.NewOptions().ClientType(mtest.Mock), func(mt *mtest.T) {
		database.Instance.Db = mt.DB
		mt.Run("success", func(mt *mtest.T) {
			expectedInvitation := data.Invitation{
				ID:    primitive.NewObjectID(),
				Token: "dOG8UVzaLk",
			}
			bytes, err := bson.Marshal(expectedInvitation)
			assert.Nil(t, err)
			var bsonD bson.D
			err = bson.Unmarshal(bytes, &bsonD)
			assert.Nil(t, err)

			mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bsonD))
			invitation, err := GetFromToken(mt.Coll, expectedInvitation.Token)
			assert.Nil(t, err)
			assert.Equal(t, &expectedInvitation, invitation)
		})
		mt.Run("error", func(t *mtest.T) {
			sentInvitation := data.Invitation{
				ID:    primitive.NewObjectID(),
				Token: "dOG8UVzaLk",
			}

			invitation, err := GetFromToken(mt.Coll, sentInvitation.Token)
			assert.NotNil(t, err)
			assert.Nil(t, invitation)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusUnauthorized, e.Code)
		})
		mt.Run("no document", func(t *mtest.T) {
			var token = "dOG8UVzaLk"

			mt.AddMockResponses(bson.D{
				{"ok", 1},
				{"cursor", bson.D{
					{"id", int64(0)},
					{"ns", "foo.bar"},
					{string(mtest.FirstBatch), bson.A{}},
				}},
			})
			invitation, err := GetFromToken(mt.Coll, token)
			assert.NotNil(t, err)
			assert.Nil(t, invitation)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusNotFound, e.Code)
		})
	})
}
