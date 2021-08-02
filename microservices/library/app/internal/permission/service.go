package permission

import (
	"github.com/alexandr-io/backend/library/data/permissions"
	groupServ "github.com/alexandr-io/backend/library/internal/group"
	userLibraryServ "github.com/alexandr-io/backend/library/internal/user_library"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Serv *Service

type Service struct {
	userLibraryRepo userLibraryServ.Repository
	groupRepo       groupServ.Repository
}

func NewService(userLibrary userLibraryServ.Repository, group groupServ.Repository) *Service {
	Serv = &Service{userLibraryRepo: userLibrary, groupRepo: group}
	return Serv
}

func getGroupHigherPermission(groups []permissions.Group, userPermissions permissions.PermissionLibrary) *permissions.PermissionLibrary {

	for _, currGroup := range groups {
		current := currGroup.Permissions

		if userPermissions.Owner == nil {
			userPermissions.Owner = current.Owner
		}

		if userPermissions.Admin == nil {
			userPermissions.Admin = current.Admin
		}

		if userPermissions.BookDelete == nil {
			userPermissions.BookDelete = current.BookDelete
		}

		if userPermissions.BookUpload == nil {
			userPermissions.BookUpload = current.BookUpload
		}

		if userPermissions.BookUpdate == nil {
			userPermissions.BookUpdate = current.BookUpdate
		}

		if userPermissions.BookDisplay == nil {
			userPermissions.BookDisplay = current.BookDisplay
		}

		if userPermissions.BookRead == nil {
			userPermissions.BookRead = current.BookRead
		}

		if userPermissions.LibraryUpdate == nil {
			userPermissions.LibraryUpdate = current.LibraryUpdate
		}

		if userPermissions.LibraryDelete == nil {
			userPermissions.LibraryDelete = current.LibraryDelete
		}

		if userPermissions.UserInvite == nil {
			userPermissions.UserInvite = current.UserInvite
		}

		if userPermissions.UserRemove == nil {
			userPermissions.UserRemove = current.UserRemove
		}

		if userPermissions.UserPermissionManage == nil {
			userPermissions.UserPermissionManage = current.UserPermissionManage
		}
	}
	return &userPermissions
}

// GetUserLibraryPermission retrieve the permissions of the user
func (s *Service) GetUserLibraryPermission(userID primitive.ObjectID, libraryID primitive.ObjectID) (*permissions.PermissionLibrary, error) {
	userLibrary, err := s.userLibraryRepo.ReadFromUserIDAndLibraryID(userID, libraryID)
	if err != nil {
		return nil, err
	}

	groups, err := s.groupRepo.ReadFromIDListAndLibraryID(userLibrary.Groups, libraryID)
	if err != nil {
		return nil, err
	}

	permissionsList := getGroupHigherPermission(*groups, userLibrary.Permissions)
	return permissionsList, nil
}
