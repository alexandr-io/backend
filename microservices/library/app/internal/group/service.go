package group

import (
	"github.com/alexandr-io/backend/library/data/permissions"
	userLibraryServ "github.com/alexandr-io/backend/library/internal/user_library"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Serv *Service

type Service struct {
	repo            Repository
	userLibraryRepo userLibraryServ.Repository
}

func NewService(repo Repository, userLibrary userLibraryServ.Repository) *Service {
	Serv = &Service{repo: repo, userLibraryRepo: userLibrary}
	return Serv
}

func (s *Service) CreateGroup(group permissions.Group) (*permissions.Group, error) {
	return s.repo.Create(group)
}

func (s *Service) ReadGroupFromIDAndLibraryID(groupID primitive.ObjectID, libraryID primitive.ObjectID) (*permissions.Group, error) {
	return s.repo.ReadFromIDAndLibraryID(groupID, libraryID)
}

func (s *Service) ReadGroupFromLibrary(libraryID primitive.ObjectID, userID primitive.ObjectID) (*[]permissions.Group, error) {
	library, err := s.userLibraryRepo.ReadFromUserIDAndLibraryID(userID, libraryID)
	if err != nil {
		return nil, err
	}

	return s.repo.ReadFromIDListAndLibraryID(library.Groups, libraryID)
}

func (s *Service) UpdateGroup(group permissions.Group) (*permissions.Group, error) {
	return s.repo.Update(group)
}

func (s *Service) DeleteGroup(groupID primitive.ObjectID) error {
	return s.repo.Delete(groupID)
}

func (s *Service) AddUserToGroup(userID primitive.ObjectID, groupID primitive.ObjectID, libraryID primitive.ObjectID) error {
	library, err := s.userLibraryRepo.ReadFromUserIDAndLibraryID(userID, libraryID)
	if err != nil {
		return err
	}

	group, err := s.repo.ReadFromIDAndLibraryID(groupID, libraryID)
	if err != nil {
		return err
	}

	library.Groups = append(library.Groups, group.ID)

	_, err = s.userLibraryRepo.Update(*library)
	if err != nil {
		return err
	}
	return nil
}
