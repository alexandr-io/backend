package internal

import (
	"github.com/alexandr-io/backend/library/data/permissions"
	"github.com/alexandr-io/backend/library/database/group"
	"github.com/alexandr-io/backend/library/database/libraries"
)

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
func GetUserLibraryPermission(userID string, libraryID string) (*permissions.PermissionLibrary, error) {
	userLibrary, err := libraries.GetFromUserIDAndLibraryID(userID, libraryID)
	if err != nil {
		return nil, err
	}

	groups, err := group.GetFromIDListAndLibraryID(userLibrary.Groups, libraryID)
	if err != nil {
		return nil, err
	}

	permissionsList := getGroupHigherPermission(*groups, userLibrary.Permissions)
	return permissionsList, nil
}
