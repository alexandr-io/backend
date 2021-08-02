package book

import (
	"github.com/alexandr-io/backend/library/data"
	BookServ "github.com/alexandr-io/backend/library/internal/book/interface"
	bookProgressServ "github.com/alexandr-io/backend/library/internal/book_progress"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Serv *Service

type Service struct {
	repo             BookServ.Repository
	bookProgressRepo bookProgressServ.Repository
}

func NewService(repo BookServ.Repository, bookProgress bookProgressServ.Repository) *Service {
	Serv = &Service{repo: repo, bookProgressRepo: bookProgress}
	return Serv
}

func (s *Service) CreateBook(book data.Book) (*data.Book, error) {
	return s.repo.Create(book)
}

func (s *Service) ReadFromID(id primitive.ObjectID) (*data.Book, error) {
	bookList, err := s.repo.Read(bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return nil, err
	}
	return &(*bookList)[0], nil
}

func (s *Service) ReadFromLibraryID(libraryID primitive.ObjectID) (*[]data.Book, error) {
	return s.repo.Read(bson.D{{Key: "library_id", Value: libraryID}})
}

func (s *Service) UpdateBook(book data.Book) (*data.Book, error) {
	return s.repo.Update(book)
}

func (s *Service) DeleteBook(id primitive.ObjectID) error {
	// TODO: create logic so that when book progress delete failed, the book previously deleted in restored
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return s.bookProgressRepo.Delete(data.BookProgressData{BookID: id})
}
