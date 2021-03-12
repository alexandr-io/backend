package permissions

// PermissionLibrary is the struct containing all the permissions of an user in a library
type PermissionLibrary struct {
	Owner *bool `json:"owner,omitempty" bson:"owner"`
	Admin *bool `json:"admin,omitempty" bson:"admin"`

	BookDelete  *bool `json:"book_delete,omitempty" bson:"book_delete"`
	BookUpload  *bool `json:"book_upload,omitempty" bson:"book_upload"`
	BookUpdate  *bool `json:"book_update,omitempty" bson:"book_update"`
	BookDisplay *bool `json:"book_display,omitempty" bson:"book_display"`
	BookRead    *bool `json:"book_read,omitempty" bson:"book_read"`

	LibraryUpdate *bool `json:"library_update,omitempty" bson:"library_update"`
	LibraryDelete *bool `json:"library_delete,omitempty" bson:"library_delete"`

	UserInvite *bool `json:"user_invite,omitempty" bson:"user_invite"`
	UserRemove *bool `json:"user_remove,omitempty" bson:"user_remove"`

	UserPermissionManage *bool `json:"user_permissions_manage,omitempty" bson:"user_permissions_manage"`
}

// IsOwner check if a user is the owner of the library
func (object *PermissionLibrary) IsOwner() bool {
	if *object.Owner == true {
		return true
	}
	return false
}

// IsAdmin check if a user is an admin of the library
func (object *PermissionLibrary) IsAdmin() bool {
	if object.IsOwner() == true {
		return true
	}

	if *object.Admin == true {
		return true
	}
	return false
}

// CanDeleteBook check if the user can delete a book for the requested library
func (object *PermissionLibrary) CanDeleteBook() bool {
	if object.IsAdmin() == true {
		return true
	}

	if *object.BookDelete == true {
		return true
	}
	return false
}

// CanUploadBook check if the user can upload a book for the requested library
func (object *PermissionLibrary) CanUploadBook() bool {
	if object.IsAdmin() == true {
		return true
	}

	if *object.BookUpload == true {
		return true
	}
	return false
}

// CanUpdateBook check if the user can update a book for the requested library
func (object *PermissionLibrary) CanUpdateBook() bool {
	if object.IsAdmin() == true {
		return true
	}

	if *object.BookUpdate == true {
		return true
	}
	return false
}

// CanDisplayBook check if the user can see the books for the requested library
func (object *PermissionLibrary) CanDisplayBook() bool {
	if object.IsAdmin() == true {
		return true
	}

	if *object.BookDisplay == true {
		return true
	}
	return false
}

// CanReadBook check if the user can read the books for the requested library
func (object *PermissionLibrary) CanReadBook() bool {
	if object.IsAdmin() == true {
		return true
	}

	if *object.BookRead == true {
		return true
	}
	return false
}

// CanUpdateLibrary check if the user can update the library information
func (object *PermissionLibrary) CanUpdateLibrary() bool {
	if object.IsAdmin() == true {
		return true
	}

	if *object.LibraryUpdate == true {
		return true
	}
	return false
}

// CanDeleteLibrary check if the user can delete the library
func (object *PermissionLibrary) CanDeleteLibrary() bool {
	if object.IsAdmin() == true {
		return true
	}

	if *object.LibraryDelete == true {
		return true
	}
	return false
}

// CanInviteUser check if the user can invite users to a library
func (object *PermissionLibrary) CanInviteUser() bool {
	if object.IsAdmin() == true {
		return true
	}

	if *object.UserInvite == true {
		return true
	}
	return false
}

// CanRemoveUser check if the user can remove users from a library
func (object *PermissionLibrary) CanRemoveUser() bool {
	if object.IsAdmin() == true {
		return true
	}

	if *object.UserRemove == true {
		return true
	}
	return false
}

// CanManagePermissions check if the user can manage other users permissions for a library
func (object *PermissionLibrary) CanManagePermissions() bool {
	if object.IsAdmin() == true {
		return true
	}

	if *object.UserPermissionManage == true {
		return true
	}
	return false
}

// BoolPtr create a bointer to a boolean
// TODO: move in common folder
func BoolPtr(value bool) *bool {
	return &value
}
