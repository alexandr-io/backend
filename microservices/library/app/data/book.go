package data

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//
// Structs
//

// Book defines the structure for an API book
type Book struct {
	ID          string   `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string   `json:"title,omitempty" bson:"title,omitempty"`
	Author      string   `json:"author,omitempty" bson:"author,omitempty"`
	Publisher   string   `json:"publisher,omitempty" bson:"publisher,omitempty"`
	Description string   `json:"description,omitempty" bson:"description,omitempty"`
	CoverID     string   `json:"cover_id,omitempty" bson:"cover_id,omitempty"`
	Tags        []string `json:"tags,omitempty" bson:"tags,omitempty"`
	LibraryID   string   `json:"library_id,omitempty" bson:"library_id,omitempty"`
	UploaderID  string   `json:"-" bson:"uploader_id,omitempty"`
}

// BookData is the structure of a book in the database
type BookData struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	LibraryID   primitive.ObjectID `bson:"library_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Author      string             `bson:"author,omitempty"`
	Publisher   string             `bson:"publisher,omitempty"`
	Description string             `bson:"description,omitempty"`
	CoverID     string             `bson:"cover_id,omitempty"`
	Tags        []string           `bson:"tags,omitempty"`
	UploaderID  string             `bson:"uploader_id,omitempty"`
}

//
// Methods
//

// ToBookData create a BookData object from a Book and return it
func (book *Book) ToBookData() (BookData, error) {
	bookData := BookData{
		Title:       book.Title,
		Author:      book.Author,
		Publisher:   book.Publisher,
		Description: book.Description,
		CoverID:     book.CoverID,
		Tags:        book.Tags,
		UploaderID:  book.UploaderID,
	}
	if book.ID != "" {
		bookID, err := primitive.ObjectIDFromHex(book.ID)
		if err != nil {
			return BookData{}, NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
		}
		bookData.ID = bookID
	}
	if book.LibraryID != "" {
		libraryID, err := primitive.ObjectIDFromHex(book.LibraryID)
		if err != nil {
			return BookData{}, NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
		}
		bookData.LibraryID = libraryID
	}

	return bookData, nil
}

// ToBook create a Book object from a BookData and return it
func (bookData *BookData) ToBook() Book {
	return Book{
		ID:          bookData.ID.Hex(),
		LibraryID:   bookData.LibraryID.Hex(),
		Title:       bookData.Title,
		Author:      bookData.Author,
		Publisher:   bookData.Publisher,
		Description: bookData.Description,
		CoverID:     bookData.CoverID,
		Tags:        bookData.Tags,
		UploaderID:  bookData.UploaderID,
	}
}
