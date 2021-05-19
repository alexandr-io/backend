package group

import (
	"testing"

	"github.com/alexandr-io/backend/common/typeconv"
	"github.com/alexandr-io/backend/library/data/permissions"
	"github.com/alexandr-io/backend/library/database"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.RunOpts("insert one", mtest.NewOptions().ClientType(mtest.Mock), func(mt *mtest.T) {
		database.Instance.Db = mt.DB
		mt.Run("success", func(mt *mtest.T) {
			database.GroupCollection = mt.Coll
			group := permissions.Group{
				ID:          primitive.NewObjectID(),
				LibraryID:   primitive.NewObjectID(),
				Name:        "Teacher",
				Description: "can manage students",
				Priority:    1,
				Permissions: permissions.PermissionLibrary{
					Admin: typeconv.BoolPtr(true),
				},
			}

			mt.AddMockResponses(mtest.CreateSuccessResponse())
			mt.AddMockResponses(mtest.CreateSuccessResponse(bson.E{Key: "id", Value: group.ID}))
			groupResponse, err := Insert(group)
			assert.Nil(t, err)
			assert.Equal(t, &group, groupResponse)
		})
		mt.Run("error update many", func(t *mtest.T) {
			database.GroupCollection = mt.Coll
			mt.AddMockResponses(bson.D{{"ok", 0}})
			group, err := Insert(permissions.Group{})
			assert.NotNil(t, err)
			assert.Nil(t, group)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
		mt.Run("error insert", func(t *mtest.T) {
			database.GroupCollection = mt.Coll
			mt.AddMockResponses(mtest.CreateSuccessResponse())
			mt.AddMockResponses(bson.D{{"ok", 0}})
			group, err := Insert(permissions.Group{})
			assert.NotNil(t, err)
			assert.Nil(t, group)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
	})
}
