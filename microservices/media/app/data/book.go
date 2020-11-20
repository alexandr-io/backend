package data

// Book is the structure of an API book.
type Book struct {
	ID        string `json:"book_id" bson:"book_id,omitempty"`
	LibraryID string `json:"-" bson:"library_id,omitempty"`
	Path      string `json:"-" bson:"path"`
}
