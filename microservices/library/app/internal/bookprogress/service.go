package bookprogress

import (
	"time"

	"github.com/alexandr-io/backend/library/data"
	bookServ "github.com/alexandr-io/backend/library/internal/book/interface"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Serv instance of book progress service
var Serv *Service

// Service is the struct containing database repository needed for book progress methods of the interface
type Service struct {
	repo     Repository
	bookRepo bookServ.Repository
}

// NewService create and set instance of Service
func NewService(repo Repository, book bookServ.Repository) *Service {
	Serv = &Service{repo: repo, bookRepo: book}
	return Serv
}

// UpsertProgression upsert a progression
func (s *Service) UpsertProgression(bookProgress data.BookProgressData) (*data.BookProgressData, error) {
	bookProgress.LastReadDate = time.Now()

	// Check book existence
	if _, err := s.bookRepo.Read(bson.D{{"_id", bookProgress.BookID}}); err != nil {
		return nil, err
	}

	// Update / Insert data
	updatedBookProgress, err := s.repo.Upsert(bookProgress)
	if err != nil {
		return nil, err
	}

	return updatedBookProgress, nil
}

// ReadProgressionFromIDs read a progression
func (s *Service) ReadProgressionFromIDs(userID primitive.ObjectID, bookID primitive.ObjectID, libraryID primitive.ObjectID) (*data.BookProgressData, error) {
	return s.repo.ReadFromIDs(userID, bookID, libraryID)
}

// DeleteProgression delete a progression
func (s *Service) DeleteProgression(bookProgress data.BookProgressData) error {
	return s.repo.Delete(bookProgress)
}
