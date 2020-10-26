package data

// BookInfo defines the structure of the information on a book
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
	// The url of the cover on the media server
	// required: true
	// example: https://media.alexandrio.cloud/ef45f[...]feUEgE7.png
	CoverUrl string `json:"cover_url,omitempty" bson:"cover_url,omitempty"`
}

// Library defines the structure for an API library
// swagger:model
type Library struct {
	// The username of the library owner
	// required: true
	// example: john
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	// The lists containing the information on the books in the library
	// required: true
	Books []BookInfo `json:"books,omitempty" bson:"books"`
}
