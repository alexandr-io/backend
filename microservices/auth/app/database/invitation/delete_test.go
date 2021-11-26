package invitation

import (
	"testing"

	"github.com/alexandr-io/backend/auth/database"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestDelete(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.RunOpts("delete one", mtest.NewOptions().ClientType(mtest.Mock).CreateClient(true), func(mt *mtest.T) {
		database.Instance.Db = mt.DB
		mt.Run("success", func(mt *mtest.T) {
			database.InvitationCollection = mt.Coll
			var token = "dOG8UVzaLk"
			mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 1}})
			err := Delete(token)
			assert.Nil(t, err)
		})
		mt.Run("not found", func(t *mtest.T) {
			database.InvitationCollection = mt.Coll
			var token = "dOG8UVzaLk"
			mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 0}})
			err := Delete(token)
			assert.NotNil(t, err)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusUnauthorized, e.Code)
		})
		mt.Run("error", func(t *mtest.T) {
			database.InvitationCollection = mt.Coll
			var token = "dOG8UVzaLk"
			mt.AddMockResponses(bson.D{{"ok", 0}})
			err := Delete(token)
			assert.NotNil(t, err)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
	})
}
