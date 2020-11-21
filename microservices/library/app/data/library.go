package data

// Library defines the structure for an API library
type Library struct {
	ID string `json:"-" bson:"_id,omitempty"`
	// The name of the library
	// required: true
	// example: Cartoons
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	// A brief description of the library
	// required: false
	// example: General library with every unsorted books
	Description string `json:"description" bson:"description"`
	// The lists containing the information on the books of this library  (the list can be empty `[]`)
	// required: true
	Books []Book `json:"books" bson:"books"`
}

// Libraries defines the structure for an API libraries
type Libraries struct {
	UserID    string   `bson:"user_id,omitempty"`
	Libraries []string `bson:"libraries"`
}

// LibraryName defines the structure for an API library name
type LibraryName struct {
	// The name of a library
	// required: true
	// example: general
	Name string `json:"name" bson:"name"`
	// The ID of the library
	// required: false
	// example: "egGefeHUfeg[...]EfYhef"
	ID string `json:"id" bson:"-"`
}

// LibrariesNames is the format of the data return when fetching libraries names.
// swagger:model
type LibrariesNames struct {
	// The names of the libraries of a user
	// required: true
	// example: [{"name": "Action", "id": "egGefeHUfeg[...]EfYhef"}]
	Libraries []LibraryName `json:"libraries" bson:"name"`
}
