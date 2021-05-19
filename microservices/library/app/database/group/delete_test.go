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

func TestDelete(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.RunOpts("delete one", mtest.NewOptions().ClientType(mtest.Mock).CreateClient(true), func(mt *mtest.T) {
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
			bsonD, err := typeconv.ToDoc(group)
			assert.Nil(t, err)

			mt.AddMockResponses(bson.D{
				{"ok", 1},
				{"value", bsonD},
			})
			mt.AddMockResponses(mtest.CreateSuccessResponse())
			err = Delete(group.ID.Hex())
			assert.Nil(t, err)
		})
		mt.Run("not found", func(t *mtest.T) {
			database.GroupCollection = mt.Coll
			mt.AddMockResponses(mtest.CreateSuccessResponse())
			//mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 0}})
			err := Delete(primitive.NewObjectID().Hex())
			assert.NotNil(t, err)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusNotFound, e.Code)
		})
		mt.Run("error", func(t *mtest.T) {
			database.BookCollection = mt.Coll
			err := Delete(primitive.NewObjectID().Hex())
			assert.NotNil(t, err)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
		mt.Run("error objectID", func(t *mtest.T) {
			database.BookCollection = mt.Coll
			err := Delete("42")
			assert.NotNil(t, err)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusBadRequest, e.Code)
		})
		mt.Run("error update many", func(t *mtest.T) {
			database.BookCollection = mt.Coll
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
			bsonD, err := typeconv.ToDoc(group)
			assert.Nil(t, err)

			mt.AddMockResponses(bson.D{
				{"ok", 1},
				{"value", bsonD},
			})
			mt.AddMockResponses(bson.D{{"ok", 0}})
			err = Delete(primitive.NewObjectID().Hex())
			assert.NotNil(t, err)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
	})
}
