package data

// Library defines the structure for an API library
type Library struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Description string `json:"description" bson:"description"`
	// The lists containing the information on the books of this library (the list can be empty `[]`)
	Books []BookData `json:"books" bson:"books"`
}

// Libraries defines the structure for an API libraries
type Libraries struct {
	UserID    string   `bson:"user_id,omitempty"`
	Libraries []string `bson:"libraries"`
}

// LibraryName defines the structure for an API library name
type LibraryName struct {
	Name string `json:"name" bson:"name" validate:"required"`
	ID   string `json:"id" bson:"-"`
}

// LibrariesNames is the format of the data return when fetching libraries names.
// swagger:model
type LibrariesNames struct {
	Libraries []LibraryName `json:"libraries" bson:"name"`
}
