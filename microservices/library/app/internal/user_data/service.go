package user_data

import (
	"github.com/alexandr-io/backend/library/data"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Serv instance of user data service
var Serv *Service

// Service is the struct containing database repository needed for user data methods of the interface
type Service struct {
	repo Repository
}

// NewService create and set instance of Service
func NewService(repo Repository) *Service {
	Serv = &Service{repo: repo}
	return Serv
}

// CreateUserData create a user data
func (s *Service) CreateUserData(userData data.UserData) (*data.UserData, error) {
	return s.repo.Create(userData)
}

// ReadUserData read a user data
func (s *Service) ReadUserData(userID, bookID, dataID primitive.ObjectID) (*data.UserData, error) {
	return s.repo.Read(userID, bookID, dataID)
}

// ReadAllUserData read many user data
func (s *Service) ReadAllUserData(userID, libraryID, bookID primitive.ObjectID) (*[]data.UserData, error) {
	return s.repo.ReadAll(userID, libraryID, bookID)
}

// UpdateUserData update user data
func (s *Service) UpdateUserData(userData data.UserData) (*data.UserData, error) {
	return s.repo.Update(userData)
}

// DeleteUserData delete user data
func (s *Service) DeleteUserData(userID, libraryID, bookID, dataID primitive.ObjectID) error {
	return s.repo.Delete(userID, libraryID, bookID, dataID)
}

// DeleteAllUserDataIn delete all user data
func (s *Service) DeleteAllUserDataIn(userID, libraryID, bookID primitive.ObjectID) error {
	return s.repo.DeleteAllIn(userID, libraryID, bookID)
}
