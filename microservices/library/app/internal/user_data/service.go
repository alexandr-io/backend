package user_data

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

func (s *Service) CreateUserData(userData data.UserData) (*data.UserData, error) {
	return s.repo.Create(userData)
}

func (s *Service) ReadUserData(userID, bookID, dataID primitive.ObjectID) (*data.UserData, error) {
	return s.repo.Read(userID, bookID, dataID)
}

func (s *Service) ReadAllUserData(userID, libraryID, bookID primitive.ObjectID) (*[]data.UserData, error) {
	return s.repo.ReadAll(userID, libraryID, bookID)
}

func (s *Service) UpdateUserData(userData data.UserData) (*data.UserData, error) {
	return s.repo.Update(userData)
}

func (s *Service) DeleteUserData(userID, libraryID, bookID, dataID primitive.ObjectID) error {
	return s.repo.Delete(userID, libraryID, bookID, dataID)
}

func (s *Service) DeleteAllUserDataIn(userID, libraryID, bookID primitive.ObjectID) error {
	return s.repo.DeleteAllIn(userID, libraryID, bookID)
}
