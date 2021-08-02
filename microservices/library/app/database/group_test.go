package database

import (
	"testing"

	"github.com/alexandr-io/backend/common/typeconv"
	"github.com/alexandr-io/backend/library/data/permissions"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestGroupCreate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		GroupDB = &GroupCollection{collection: mt.Coll}
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
		groupResponse, err := GroupDB.Create(group)
		assert.Nil(t, err)
		assert.Equal(t, &group, groupResponse)
	})

	mt.Run("error update many", func(mt *mtest.T) {
		GroupDB = &GroupCollection{collection: mt.Coll}
		mt.AddMockResponses(bson.D{{"ok", 0}})
		group, err := GroupDB.Create(permissions.Group{})
		assert.NotNil(t, err)
		assert.Nil(t, group)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})

	mt.Run("error insert", func(mt *mtest.T) {
		GroupDB = &GroupCollection{collection: mt.Coll}
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		mt.AddMockResponses(bson.D{{"ok", 0}})
		group, err := GroupDB.Create(permissions.Group{})
		assert.NotNil(t, err)
		assert.Nil(t, group)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})
}

func TestGroupReadFromIDAndLibraryID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		GroupDB = &GroupCollection{collection: mt.Coll}
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
		group, err := GroupDB.ReadFromIDAndLibraryID(expectedGroup.ID, expectedGroup.LibraryID)
		assert.Nil(t, err)
		assert.Equal(t, &expectedGroup, group)
	})

	mt.Run("error", func(mt *mtest.T) {
		GroupDB = &GroupCollection{collection: mt.Coll}

		group, err := GroupDB.ReadFromIDAndLibraryID(primitive.NewObjectID(), primitive.NewObjectID())
		assert.NotNil(t, err)
		assert.Nil(t, group)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})

	mt.Run("no document", func(mt *mtest.T) {
		GroupDB = &GroupCollection{collection: mt.Coll}
		mt.AddMockResponses(bson.D{
			{"ok", 1},
			{"cursor", bson.D{
				{"id", int64(0)},
				{"ns", "foo.bar"},
				{string(mtest.FirstBatch), bson.A{}},
			}},
		})
		group, err := GroupDB.ReadFromIDAndLibraryID(primitive.NewObjectID(), primitive.NewObjectID())
		assert.NotNil(t, err)
		assert.Nil(t, group)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusNotFound, e.Code)
	})
}

func TestGroupReadFromIDListAndLibraryID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		GroupDB = &GroupCollection{collection: mt.Coll}
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
		groups, err := GroupDB.ReadFromIDListAndLibraryID(ids, libraryID)
		assert.Nil(t, err)
		assert.Equal(t, &expectedGroups, groups)
	})

	mt.Run("success no ID", func(mt *mtest.T) {
		GroupDB = &GroupCollection{collection: mt.Coll}

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
		groups, err := GroupDB.ReadFromIDListAndLibraryID(nil, libraryID)
		assert.Nil(t, err)
		assert.Equal(t, &expectedGroups, groups)
	})

	mt.Run("find error", func(mt *mtest.T) {
		GroupDB = &GroupCollection{collection: mt.Coll}

		groups, err := GroupDB.ReadFromIDListAndLibraryID(nil, primitive.NewObjectID())
		assert.NotNil(t, err)
		assert.Nil(t, groups)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})

	mt.Run("cursor read all error", func(mt *mtest.T) {
		GroupDB = &GroupCollection{collection: mt.Coll}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{}))
		groups, err := GroupDB.ReadFromIDListAndLibraryID(nil, primitive.NewObjectID())
		assert.NotNil(t, err)
		assert.Nil(t, groups)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})
}

func TestGroupUpdate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		GroupDB = &GroupCollection{collection: mt.Coll}
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
		group, err := GroupDB.Update(expectedGroup)
		assert.Nil(t, err)
		assert.Equal(t, &expectedGroup, group)
	})

	mt.Run("error", func(mt *mtest.T) {
		GroupDB = &GroupCollection{collection: mt.Coll}
		group, err := GroupDB.Update(permissions.Group{})
		assert.NotNil(t, err)
		assert.Nil(t, group)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})

	mt.Run("no document", func(mt *mtest.T) {
		GroupDB = &GroupCollection{collection: mt.Coll}
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
		group, err := GroupDB.Update(sentGroup)
		assert.NotNil(t, err)
		assert.Nil(t, group)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusNotFound, e.Code)
	})
}

func TestGroupDelete(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		GroupDB = &GroupCollection{collection: mt.Coll}
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
		err = GroupDB.Delete(group.ID)
		assert.Nil(t, err)
	})

	mt.Run("not found", func(mt *mtest.T) {
		GroupDB = &GroupCollection{collection: mt.Coll}
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		//mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 0}})
		err := GroupDB.Delete(primitive.NewObjectID())
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusNotFound, e.Code)
	})

	mt.Run("error", func(mt *mtest.T) {
		GroupDB = &GroupCollection{collection: mt.Coll}
		err := GroupDB.Delete(primitive.NewObjectID())
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})

	mt.Run("error update many", func(mt *mtest.T) {
		GroupDB = &GroupCollection{collection: mt.Coll}
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
		err = GroupDB.Delete(primitive.NewObjectID())
		assert.NotNil(t, err)

		e, ok := err.(*fiber.Error)
		assert.True(t, ok)
		assert.Equal(t, fiber.StatusInternalServerError, e.Code)
	})
}
