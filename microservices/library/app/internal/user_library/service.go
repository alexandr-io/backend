package user_library

import (
	"github.com/alexandr-io/backend/library/data"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Serv *Service

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	Serv = &Service{repo: repo}
	return Serv
}

func (s *Service) CreateUserLibrary(userLibrary data.UserLibrary) (*data.UserLibrary, error) {
	return s.repo.Create(userLibrary)
}

func (s *Service) ReadUserLibraryFromUserID(userID string) (*[]data.Library, error) {
	return s.repo.ReadFromUserID(userID)
}

func (s *Service) ReadUserLibraryFromUserIDAndLibraryID(userID primitive.ObjectID, libraryID primitive.ObjectID) (*data.UserLibrary, error) {
	return s.repo.ReadFromUserIDAndLibraryID(userID, libraryID)
}

func (s *Service) UpdateUserLibrary(library data.UserLibrary) (*data.UserLibrary, error) {
	return s.repo.Update(library)
}

func (s *Service) DeleteUserLibrary(id primitive.ObjectID) error {
	return s.repo.Delete(id)
}
