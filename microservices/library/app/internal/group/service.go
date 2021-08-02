package group

import (
	"github.com/alexandr-io/backend/library/data/permissions"
	userLibraryServ "github.com/alexandr-io/backend/library/internal/userlibrary"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Serv instance of group service
var Serv *Service

// Service is the struct containing database repository needed for group methods of the interface
type Service struct {
	repo            Repository
	userLibraryRepo userLibraryServ.Repository
}

// NewService create and set instance of Service
func NewService(repo Repository, userLibrary userLibraryServ.Repository) *Service {
	Serv = &Service{repo: repo, userLibraryRepo: userLibrary}
	return Serv
}

// CreateGroup create a group
func (s *Service) CreateGroup(group permissions.Group) (*permissions.Group, error) {
	return s.repo.Create(group)
}

// ReadGroupFromIDAndLibraryID read a group
func (s *Service) ReadGroupFromIDAndLibraryID(groupID primitive.ObjectID, libraryID primitive.ObjectID) (*permissions.Group, error) {
	return s.repo.ReadFromIDAndLibraryID(groupID, libraryID)
}

// ReadGroupFromLibrary read groups
func (s *Service) ReadGroupFromLibrary(libraryID primitive.ObjectID, userID primitive.ObjectID) (*[]permissions.Group, error) {
	library, err := s.userLibraryRepo.ReadFromUserIDAndLibraryID(userID, libraryID)
	if err != nil {
		return nil, err
	}

	return s.repo.ReadFromIDListAndLibraryID(library.Groups, libraryID)
}

// UpdateGroup update a group
func (s *Service) UpdateGroup(group permissions.Group) (*permissions.Group, error) {
	return s.repo.Update(group)
}

// DeleteGroup delete a group
func (s *Service) DeleteGroup(groupID primitive.ObjectID) error {
	return s.repo.Delete(groupID)
}

// AddUserToGroup add a user to a group
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
