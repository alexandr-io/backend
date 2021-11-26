package tests

// PermissionLibrary is the struct containing all the permissions of an user in a library, Copy for test
type PermissionLibrary struct {
	Owner                bool
	Admin                bool
	BookDelete           bool
	BookUpload           bool
	BookUpdate           bool
	BookDisplay          bool
	BookRead             bool
	LibraryUpdate        bool
	LibraryDelete        bool
	UserInvite           bool
	UserRemove           bool
	UserPermissionManage bool
}
