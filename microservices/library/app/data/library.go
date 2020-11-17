package data

// BookInfo defines the structure of the information on a book
// swagger:model
type BookInfo struct {
	// The title of the book
	// required: true
	// example: Pride and Prejudice
	Title string `json:"title,omitempty" bson:"title,omitempty"`
	// The author of the book
	// required: true
	// example: Jane Austen
	Author string `json:"author,omitempty" bson:"author,omitempty"`
	// The publisher of the book
	// required: true
	// example: Public domain
	Publisher string `json:"publisher,omitempty" bson:"publisher,omitempty"`
	// The description of the book
	// required: true
	// example: Pride and Prejudice is set in rural England in the early 19th century [...] and by prejudice against Darcy’s snobbery.
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	// The id of the cover on the media server
	// required: true
	// example: ef45f[...]feUEgE7
	CoverID string `json:"cover_id,omitempty" bson:"cover_id,omitempty"`
}

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
	Books []BookInfo `json:"books" bson:"books"`
}

type Libraries struct {
	UserID    string   `bson:"user_id,omitempty"`
	Libraries []string `bson:"libraries"`
}

// LibrariesNames is the format of the data return when fetching libraries names.
// swagger:model
type LibrariesNames struct {
	// The names of the libraries of a user
	// required: true
	// example: ["general", "action", "adventure"]
	Names []string `json:"names" bson:"name"`
}

type LibraryName struct {
	// The name of a library
	// required: true
	// example: general
	Name string `json:"name" bson:"name"`
}
