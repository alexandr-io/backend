package permissions

type Permission string

const (
	// Owner permission determine if a user is the library owner, can't be assign to more/less than 1 user
	// The Owner has every permissions
	Owner Permission = "owner"
	// Admin permission give a user every permissions
	Admin = "admin"

	// BookDelete permission allow a user to delete books from the library
	BookDelete = "book_delete"
	// BookUpload permission allow a user to upload books to the library
	BookUpload = "book_upload"
	// BookUpdate permission allow a user to modify the metadata of a book
	BookUpdate = "book_update"

	// BookDisplay permission allow a user to retrieve the books of the library and it's metadata
	// Given by default
	BookDisplay = "book_display"
	// BookRead permission allow a user to read books from the library
	// Given by default
	BookRead = "book_read"

	// LibraryUpdate permission allow a user to edit the library title
	LibraryUpdate = "library_update"
	// LibraryDelete permission allow a user to delete the library and every books in it
	LibraryDelete = "library_delete"

	// UserInvite permission allow a user to invite other users to the library
	UserInvite = "user_invite"
	// UserRemove permission allow a user to remove other users from the library
	// Not applicable to an Admin except if the user is the Owner
	UserRemove = "user_remove"

	// The UserPermissionManage permission allow a user to manage other users in the library permissions
	UserPermissionManage = "user_permissions_manage"
)
