package data

// LibraryCreation is the body parameter given to create a new  library
// swagger:model
type LibraryOwner struct {
	// The username of the library owner
	// required: true
	// example: john
	Username string `json:"username,omitempty"`
}
