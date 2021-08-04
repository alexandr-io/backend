package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/alexandr-io/backend/common/typeconv"
	"github.com/alexandr-io/backend/library/data/permissions"
	groupServ "github.com/alexandr-io/backend/library/internal/group"
	mockGroup "github.com/alexandr-io/backend/library/internal/group/mock"
	"github.com/alexandr-io/backend/library/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateGroup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGroupServ := mockGroup.NewMockInternal(ctrl)
	app := fiber.New()
	CreateGroupHandlers(app)
	groupServ.Serv = mockGroupServ

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
	groupSentJSON, err := json.Marshal(groupSent)
	assert.Nil(t, err)

	groupExpectedResult := groupSent
	groupExpectedResult.ID = primitive.NewObjectID()

	mockGroupServ.EXPECT().CreateGroup(groupSent).Return(&groupExpectedResult, nil)

	req, _ := http.NewRequest(
		"POST",
		"/library/"+groupSent.LibraryID.Hex()+"/group",
		bytes.NewBuffer(groupSentJSON),
	)
	req.Header.Set("Content-Type", "application/json")
	middleware.Test = true
	res, err := app.Test(req, -1)
	assert.Nil(t, err)
	defer res.Body.Close()
	assert.Equal(t, 201, res.StatusCode)

	var responseData permissions.Group
	err = json.NewDecoder(res.Body).Decode(&responseData)
	assert.Nil(t, err)
	assert.Equal(t, groupExpectedResult, responseData)
}
