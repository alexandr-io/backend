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
	_ "go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.RunOpts("find one", mtest.NewOptions().ClientType(mtest.Mock), func(mt *mtest.T) {
		database.Instance.Db = mt.DB
		mt.Run("success", func(mt *mtest.T) {
			var token = "dOG8UVzaLk"
			id := primitive.NewObjectID()

			mt.AddMockResponses(mtest.CreateSuccessResponse(bson.E{Key: "id", Value: id}))
			invitation, err := Insert(mt.Coll, data.Invitation{
				ID:     id,
				Token:  token,
				Used:   nil,
				UserID: nil,
			})
			assert.Nil(t, err)
			assert.Equal(t, &data.Invitation{
				ID:     id,
				Token:  token,
				Used:   nil,
				UserID: nil,
			}, invitation)
		})
		mt.Run("error", func(t *mtest.T) {
			mt.AddMockResponses(bson.D{{"ok", 0}})
			invitation, err := Insert(mt.Coll, data.Invitation{})
			assert.NotNil(t, err)
			assert.Nil(t, invitation)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
	})
}
