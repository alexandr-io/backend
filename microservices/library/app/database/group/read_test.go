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

func TestRead(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.RunOpts("from id", mtest.NewOptions().ClientType(mtest.Mock), func(mt *mtest.T) {
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

			mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bsonD))
			group, err := GetFromIDAndLibraryID(expectedGroup.ID.Hex(), expectedGroup.LibraryID.Hex())
			assert.Nil(t, err)
			assert.Equal(t, &expectedGroup, group)
		})
		mt.Run("error", func(t *mtest.T) {
			database.GroupCollection = mt.Coll

			group, err := GetFromIDAndLibraryID(primitive.NewObjectID().Hex(), primitive.NewObjectID().Hex())
			assert.NotNil(t, err)
			assert.Nil(t, group)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
		mt.Run("no document", func(t *mtest.T) {
			database.GroupCollection = mt.Coll
			mt.AddMockResponses(bson.D{
				{"ok", 1},
				{"cursor", bson.D{
					{"id", int64(0)},
					{"ns", "foo.bar"},
					{string(mtest.FirstBatch), bson.A{}},
				}},
			})
			group, err := GetFromIDAndLibraryID(primitive.NewObjectID().Hex(), primitive.NewObjectID().Hex())
			assert.NotNil(t, err)
			assert.Nil(t, group)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusNotFound, e.Code)
		})
		mt.Run("error ObjectID groupID", func(t *mtest.T) {
			database.GroupCollection = mt.Coll

			group, err := GetFromIDAndLibraryID("42", primitive.NilObjectID.Hex())
			assert.NotNil(t, err)
			assert.Nil(t, group)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusBadRequest, e.Code)
		})
		mt.Run("error ObjectID libraryID", func(t *mtest.T) {
			database.GroupCollection = mt.Coll

			group, err := GetFromIDAndLibraryID(primitive.NewObjectID().Hex(), "42")
			assert.NotNil(t, err)
			assert.Nil(t, group)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusBadRequest, e.Code)
		})
	})

	mt.RunOpts("from id list", mtest.NewOptions().ClientType(mtest.Mock), func(mt *mtest.T) {
		database.Instance.Db = mt.DB
		mt.Run("success", func(mt *mtest.T) {
			database.GroupCollection = mt.Coll
			libraryID := primitive.NewObjectID()

			expectedGroups := []permissions.Group{
				{
					ID:          primitive.NewObjectID(),
					LibraryID:   libraryID,
					Name:        "Teacher",
					Description: "can manage students",
					Priority:    1,
					Permissions: permissions.PermissionLibrary{
						Admin: typeconv.BoolPtr(true),
					},
				},
				{
					ID:          primitive.NewObjectID(),
					LibraryID:   libraryID,
					Name:        "Student",
					Description: "can read books",
					Priority:    1,
					Permissions: permissions.PermissionLibrary{
						BookRead: typeconv.BoolPtr(true),
					},
				},
			}
			ids := []primitive.ObjectID{
				expectedGroups[0].ID,
				expectedGroups[1].ID,
			}
			bsonD1, err := typeconv.ToDoc(expectedGroups[0])
			assert.Nil(t, err)
			bsonD2, err := typeconv.ToDoc(expectedGroups[1])
			assert.Nil(t, err)

			// Create mock responses.
			find := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bsonD1)
			getMore := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, bsonD2)
			killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
			mt.AddMockResponses(find, getMore, killCursors)
			groups, err := GetFromIDListAndLibraryID(ids, libraryID.Hex())
			assert.Nil(t, err)
			assert.Equal(t, &expectedGroups, groups)
		})
		mt.Run("success no ID", func(mt *mtest.T) {
			database.GroupCollection = mt.Coll

			libraryID := primitive.NewObjectID()
			expectedGroups := []permissions.Group{
				{
					ID:          primitive.NewObjectID(),
					LibraryID:   libraryID,
					Name:        "Teacher",
					Description: "can manage students",
					Priority:    1,
					Permissions: permissions.PermissionLibrary{
						Admin: typeconv.BoolPtr(true),
					},
				},
				{
					ID:          primitive.NewObjectID(),
					LibraryID:   libraryID,
					Name:        "Student",
					Description: "can read books",
					Priority:    1,
					Permissions: permissions.PermissionLibrary{
						BookRead: typeconv.BoolPtr(true),
					},
				},
			}
			bsonD1, err := typeconv.ToDoc(expectedGroups[0])
			assert.Nil(t, err)
			bsonD2, err := typeconv.ToDoc(expectedGroups[1])
			assert.Nil(t, err)

			// Create mock responses.
			find := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bsonD1)
			getMore := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, bsonD2)
			killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
			mt.AddMockResponses(find, getMore, killCursors)
			groups, err := GetFromIDListAndLibraryID(nil, libraryID.Hex())
			assert.Nil(t, err)
			assert.Equal(t, &expectedGroups, groups)
		})
		mt.Run("find error", func(t *mtest.T) {
			database.GroupCollection = mt.Coll

			groups, err := GetFromIDListAndLibraryID(nil, primitive.NewObjectID().Hex())
			assert.NotNil(t, err)
			assert.Nil(t, groups)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
		mt.Run("cursor read all error", func(t *mtest.T) {
			database.GroupCollection = mt.Coll

			mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{}))
			groups, err := GetFromIDListAndLibraryID(nil, primitive.NewObjectID().Hex())
			assert.NotNil(t, err)
			assert.Nil(t, groups)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusInternalServerError, e.Code)
		})
		mt.Run("error ObjectID libraryID", func(t *mtest.T) {
			database.GroupCollection = mt.Coll

			groups, err := GetFromIDListAndLibraryID(nil, "42")
			assert.NotNil(t, err)
			assert.Nil(t, groups)

			e, ok := err.(*fiber.Error)
			assert.True(t, ok)
			assert.Equal(t, fiber.StatusBadRequest, e.Code)
		})
	})
}
