package permission

import (
	"github.com/alexandr-io/backend/library/data"
)

// Internal metadata service interface
type Internal interface {
	RequestMetadata(title string, authors string) (*data.Book, error)
}
