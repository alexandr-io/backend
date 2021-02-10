package data

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BookCreation defines the structure for an API book for creation
type BookCreation struct {
	ID          string   `json:"book_id,omitempty"`
	Title       string   `json:"title,omitempty" validate:"required"`
	Author      string   `json:"author,omitempty"`
	Publisher   string   `json:"publisher,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	LibraryID   string   `json:"library_id,omitempty"`
	UploaderID  string   `json:"-"`
}

// BookRetrieve defines the structure for an API book for retrieval
type BookRetrieve struct {
	ID         string `json:"book_id,omitempty" validate:"required"`
	LibraryID  string `json:"library_id,omitempty" validate:"required"`
	UploaderID string `json:"-"`
}

// Book defines the structure for an API book
type Book struct {
	ID          string   `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string   `json:"title,omitempty" bson:"title,omitempty"`
	Author      string   `json:"author,omitempty" bson:"author,omitempty"`
	Publisher   string   `json:"publisher,omitempty" bson:"publisher,omitempty"`
	Description string   `json:"description,omitempty" bson:"description,omitempty"`
	CoverID     string   `json:"cover_id,omitempty" bson:"cover_id,omitempty"`
	Tags        []string `json:"tags,omitempty" bson:"tags,omitempty"`
	UploaderID  string   `json:"-" bson:"uploader_id,omitempty"`
}

// ToBookData create a BookData object from a Book and return it
func (book *Book) ToBookData() (BookData, error) {

	id, err := primitive.ObjectIDFromHex(book.ID)
	if err != nil {
		return BookData{}, NewHTTPErrorInfo(fiber.StatusUnauthorized, err.Error())
	}

	return BookData{
		ID:          id,
		Title:       book.Title,
		Author:      book.Author,
		Publisher:   book.Publisher,
		Description: book.Description,
		CoverID:     book.CoverID,
		Tags:        book.Tags,
		UploaderID:  book.UploaderID,
	}, nil
}

// BookData is the structure of a book in the database
type BookData struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Author      string             `bson:"author,omitempty"`
	Publisher   string             `bson:"publisher,omitempty"`
	Description string             `bson:"description,omitempty"`
	CoverID     string             `bson:"cover_id,omitempty"`
	Tags        []string           `bson:"tags,omitempty"`
	UploaderID  string             `bson:"uploader_id,omitempty"`
}

// ToBook create a Book object from a BookData and return it
func (bookData *BookData) ToBook() Book {
	return Book{
		ID:          bookData.ID.Hex(),
		Title:       bookData.Title,
		Author:      bookData.Author,
		Publisher:   bookData.Publisher,
		Description: bookData.Description,
		CoverID:     bookData.CoverID,
		Tags:        bookData.Tags,
		UploaderID:  bookData.UploaderID,
	}
}
