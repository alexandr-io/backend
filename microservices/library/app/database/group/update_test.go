package group

import (
	"github.com/alexandr-io/backend/library/data/permissions"
	"github.com/gofiber/fiber/v2"
	"testing"

	"github.com/alexandr-io/backend/common/typeconv"
	"github.com/alexandr-io/backend/library/database"

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
			database.GroupCollection = mt.Coll
			expectedGroup := permissions.Group{
				ID:          primitive.NewObjectID(),
				LibraryID:   primitive.NewObjectID(),
				Name:        "Teacher",
				Description: "can manage students",
				Priority:    1,
				Permissions: permissions.PermissionLibrary{
					Admin: typeconv.BoolPtr(true),
				},
			}
			bsonD, err := typeconv.ToDoc(expectedGroup)
			assert.Nil(t, err)

			mt.AddMockResponses(bson.D{
				{"ok", 1},
				{"value", bsonD},
			})
			group, err := Update(expectedGroup)
			assert.Nil(t, err)
			assert.Equal(t, &expectedGroup, group)
		})
		mt.Run("error", func(t *mtest.T) {
			database.GroupCollection = mt.Coll
			group, err := Update(permissions.Group{})
			assert.NotNil(t, err)
			assert.Nil(t, group)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
		mt.Run("no document", func(t *mtest.T) {
			database.GroupCollection = mt.Coll
			sentGroup := permissions.Group{
				ID:          primitive.NewObjectID(),
				LibraryID:   primitive.NewObjectID(),
				Name:        "Teacher",
				Description: "can manage students",
				Priority:    1,
				Permissions: permissions.PermissionLibrary{
					Admin: typeconv.BoolPtr(true),
				},
			}

			mt.AddMockResponses(bson.D{
				{"ok", 1},
			})
			group, err := Update(sentGroup)
			assert.NotNil(t, err)
			assert.Nil(t, group)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusNotFound, e.Code)
		})
	})

	//TODO: add test for AddUserToGroup
}
