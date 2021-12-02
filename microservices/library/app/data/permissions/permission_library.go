package permissions

// PermissionLibrary is the struct containing all the permissions of a user in a library
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
	return *object.Owner
}

// IsAdmin check if a user is an admin of the library
func (object *PermissionLibrary) IsAdmin() bool {
	return object.IsOwner() || *object.Admin
}

// CanDeleteBook check if the user can delete a book for the requested library
func (object *PermissionLibrary) CanDeleteBook() bool {
	return object.IsAdmin() || *object.BookDelete
}

// CanUploadBook check if the user can upload a book for the requested library
func (object *PermissionLibrary) CanUploadBook() bool {
	return object.IsAdmin() || *object.BookUpload
}

// CanUpdateBook check if the user can update a book for the requested library
func (object *PermissionLibrary) CanUpdateBook() bool {
	return object.IsAdmin() || *object.BookUpdate
}

// CanDisplayBook check if the user can see the books for the requested library
func (object *PermissionLibrary) CanDisplayBook() bool {
	return object.IsAdmin() || *object.BookDisplay
}

// CanReadBook check if the user can read the books for the requested library
func (object *PermissionLibrary) CanReadBook() bool {
	return object.IsAdmin() || *object.BookRead
}

// CanUpdateLibrary check if the user can update the library information
func (object *PermissionLibrary) CanUpdateLibrary() bool {
	return object.IsAdmin() || *object.LibraryUpdate
}

// CanDeleteLibrary check if the user can delete the library
func (object *PermissionLibrary) CanDeleteLibrary() bool {
	return object.IsAdmin() || *object.LibraryDelete
}

// CanInviteUser check if the user can invite users to a library
func (object *PermissionLibrary) CanInviteUser() bool {
	return object.IsAdmin() || *object.UserInvite
}

// CanRemoveUser check if the user can remove users from a library
func (object *PermissionLibrary) CanRemoveUser() bool {
	return object.IsAdmin() || *object.UserRemove
}

// CanManagePermissions check if the user can manage other users permissions for a library
func (object *PermissionLibrary) CanManagePermissions() bool {
	return object.IsAdmin() || *object.UserPermissionManage
}
