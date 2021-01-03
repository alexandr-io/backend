package internal

import (
	"context"

	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database"
)

func CanUserModifyBook(userID string, libraryID string, bookID string) (bool, error) {
	if ok, err := HasUserAccessToLibraryFromID(userID, libraryID); err != nil {
		return false, err
	} else if !ok {
		return false, nil
	}

	bookRetrieve := data.BookRetrieve{
		ID:         bookID,
		LibraryID:  libraryID,
		UploaderID: userID,
	}
	book, err := database.BookRetrieve(context.Background(), bookRetrieve)
	if err != nil {
		return false, err
	}
	if book.UploaderID == userID {
		return true, nil
	}
	return false, nil
}
