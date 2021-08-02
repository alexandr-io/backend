package book_progress

import (
	"time"

	"github.com/alexandr-io/backend/library/data"
	bookServ "github.com/alexandr-io/backend/library/internal/book/interface"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Serv *Service

type Service struct {
	repo     Repository
	bookRepo bookServ.Repository
}

func NewService(repo Repository, book bookServ.Repository) *Service {
	Serv = &Service{repo: repo, bookRepo: book}
	return Serv
}

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

func (s *Service) ReadProgressionFromIDs(userID primitive.ObjectID, bookID primitive.ObjectID, libraryID primitive.ObjectID) (*data.BookProgressData, error) {
	return s.repo.ReadFromIDs(userID, bookID, libraryID)
}

func (s *Service) DeleteProgression(bookProgress data.BookProgressData) error {
	return s.repo.Delete(bookProgress)
}
