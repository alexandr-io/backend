package userlibrary

import (
	"github.com/alexandr-io/backend/library/data"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Serv instance of user library service
var Serv *Service

// Service is the struct containing database repository needed for user library methods of the interface
type Service struct {
	repo Repository
}

// NewService create and set instance of Service
func NewService(repo Repository) *Service {
	Serv = &Service{repo: repo}
	return Serv
}

// CreateUserLibrary create user library
func (s *Service) CreateUserLibrary(userLibrary data.UserLibrary) (*data.UserLibrary, error) {
	return s.repo.Create(userLibrary)
}

// ReadUserLibraryFromUserID read many user library
func (s *Service) ReadUserLibraryFromUserID(userID string) (*[]data.Library, error) {
	return s.repo.ReadFromUserID(userID)
}

// ReadUserLibraryFromUserIDAndLibraryID read a user library
func (s *Service) ReadUserLibraryFromUserIDAndLibraryID(userID primitive.ObjectID, libraryID primitive.ObjectID) (*data.UserLibrary, error) {
	return s.repo.ReadFromUserIDAndLibraryID(userID, libraryID)
}

// UpdateUserLibrary update user library
func (s *Service) UpdateUserLibrary(library data.UserLibrary) (*data.UserLibrary, error) {
	return s.repo.Update(library)
}

// DeleteUserLibrary delete user library
func (s *Service) DeleteUserLibrary(id primitive.ObjectID) error {
	return s.repo.Delete(id)
}
