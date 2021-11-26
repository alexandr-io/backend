package group

import (
	"testing"

	"github.com/alexandr-io/backend/common/typeconv"
	"github.com/alexandr-io/backend/library/data/permissions"
	mockGroup "github.com/alexandr-io/backend/library/internal/group/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateGroup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGroupRepo := mockGroup.NewMockRepository(ctrl)
	NewService(mockGroupRepo, nil)

	groupSent := permissions.Group{
		LibraryID:   primitive.NewObjectID(),
		Name:        "Foo",
		Description: "bar",
		Priority:    1,
		Permissions: permissions.PermissionLibrary{
			BookUpload:  typeconv.BoolPtr(true),
			BookUpdate:  typeconv.BoolPtr(true),
			BookDisplay: typeconv.BoolPtr(true),
			BookRead:    typeconv.BoolPtr(true),
		},
	}
	groupExpectedResult := groupSent
	groupExpectedResult.ID = primitive.NewObjectID()

	mockGroupRepo.EXPECT().Create(groupSent).Return(&groupExpectedResult, nil)

	groupResult, err := Serv.CreateGroup(groupSent)

	assert.Nil(t, err)
	assert.NotNil(t, groupResult)
	assert.Equal(t, &groupExpectedResult, groupResult)
}
