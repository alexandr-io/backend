package data

import "github.com/alexandr-io/backend/library/data/permissions"

// User is the struct to retrieve the users' permissions
type User struct {
	ID          string
	Username    string
	Email       string
	Permissions []permissions.Permission
}

func (user *User) isPermissionInPermissions(permission permissions.Permission) bool {
	for _, currentPermission := range user.Permissions {
		if currentPermission == permission {
			return true
		}
	}
	return false
}

// IsOwner check if a user is the owner of the library
func (user *User) IsOwner() bool {
	return user.isPermissionInPermissions(permissions.Owner)
}

// IsAdmin check if a user is an admin of the library
func (user *User) IsAdmin() bool {
	if user.IsOwner() == true {
		return true
	}
	return user.isPermissionInPermissions(permissions.Admin)
}

// CanDeleteBook check if the user can delete a book for the requested library
func (user *User) CanDeleteBook() bool {
	if user.IsAdmin() == true {
		return true
	}
	return user.isPermissionInPermissions(permissions.BookDelete)
}

// CanUploadBook check if the user can upload a book for the requested library
func (user *User) CanUploadBook() bool {
	if user.IsAdmin() == true {
		return true
	}
	return user.isPermissionInPermissions(permissions.BookUpload)
}

// CanUpdateBook check if the user can update a book for the requested library
func (user *User) CanUpdateBook() bool {
	if user.IsAdmin() == true {
		return true
	}
	return user.isPermissionInPermissions(permissions.BookUpdate)
}

// CanSeeBooks check if the user can see the books for the requested library
func (user *User) CanSeeBooks() bool {
	if user.IsAdmin() == true {
		return true
	}
	return user.isPermissionInPermissions(permissions.BookDisplay)
}

// CanReadBooks check if the user can read the books for the requested library
func (user *User) CanReadBooks() bool {
	if user.IsAdmin() == true {
		return true
	}
	return user.isPermissionInPermissions(permissions.BookRead)
}

// CanUpdateLibrary check if the user can update the library information
func (user *User) CanUpdateLibrary() bool {
	if user.IsAdmin() == true {
		return true
	}
	return user.isPermissionInPermissions(permissions.LibraryUpdate)
}

// CanDeleteLibrary check if the user can delete the library
func (user *User) CanDeleteLibrary() bool {
	if user.IsAdmin() == true {
		return true
	}
	return user.isPermissionInPermissions(permissions.LibraryDelete)
}

// CanInviteUser check if the user can invite users to a library
func (user *User) CanInviteUser() bool {
	if user.IsAdmin() == true {
		return true
	}
	return user.isPermissionInPermissions(permissions.UserInvite)
}

// CanRemoveUser check if the user can remove users from a library
func (user *User) CanRemoveUser() bool {
	if user.IsAdmin() == true {
		return true
	}
	return user.isPermissionInPermissions(permissions.UserRemove)
}

// CanManagePermissions check if the user can manage other users permissions for a library
func (user *User) CanManagePermissions(userTarget *User) bool {
	if user.IsOwner() {
		return true
	}
	if user.IsAdmin() {
		if userTarget.IsAdmin() {
			return false
		}
		return true
	}
	return user.isPermissionInPermissions(permissions.UserPermissionManage)
}
