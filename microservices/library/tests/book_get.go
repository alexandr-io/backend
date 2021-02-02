package tests

// bookRetrieve defines the structure of an API book retrieval. Copy for test
type bookRetrieve struct {
	ID        *string `json:"book_id,omitempty"`
	LibraryID *string `json:"library_id,omitempty"`
}
