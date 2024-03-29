package library

import (
	"github.com/alexandr-io/backend/common/typeconv"
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/data/permissions"
	bookProgressServ "github.com/alexandr-io/backend/library/internal/bookprogress"
	groupServ "github.com/alexandr-io/backend/library/internal/group"
	userLibraryServ "github.com/alexandr-io/backend/library/internal/userlibrary"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Serv instance of library service
var Serv *Service

// Service is the struct containing database repository needed for library methods of the interface
type Service struct {
	repo             Repository
	userLibraryRepo  userLibraryServ.Repository
	groupRepo        groupServ.Repository
	bookProgressRepo bookProgressServ.Repository
}

// NewService create and set instance of Service
func NewService(repo Repository, userLibrary userLibraryServ.Repository, group groupServ.Repository, bookProgress bookProgressServ.Repository) *Service {
	Serv = &Service{repo: repo, userLibraryRepo: userLibrary, groupRepo: group, bookProgressRepo: bookProgress}
	return Serv
}

// CreateLibrary create a library
func (s *Service) CreateLibrary(library data.Library, userID primitive.ObjectID) (*data.Library, error) {
	// insert library
	insertedLibrary, err := s.repo.Create(library)
	if err != nil {
		return nil, err
	}

	// insert user library
	libraryID, err := primitive.ObjectIDFromHex(insertedLibrary.ID)
	if err != nil {
		return nil, data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error())
	}
	userLibrary := data.UserLibrary{
		UserID:      userID,
		LibraryID:   libraryID,
		Permissions: permissions.PermissionLibrary{Owner: typeconv.BoolPtr(true)},
	}
	if _, err = s.userLibraryRepo.Create(userLibrary); err != nil {
		return nil, err
	}

	// insert permission group for new library
	if _, err = s.groupRepo.Create(permissions.Group{
		LibraryID:   libraryID,
		Name:        "everyone",
		Description: "Group with every user in it.",
		Priority:    -1,
		Permissions: permissions.PermissionLibrary{
			Owner:                typeconv.BoolPtr(false),
			Admin:                typeconv.BoolPtr(false),
			BookDelete:           typeconv.BoolPtr(false),
			BookUpload:           typeconv.BoolPtr(false),
			BookUpdate:           typeconv.BoolPtr(false),
			BookDisplay:          typeconv.BoolPtr(true),
			BookRead:             typeconv.BoolPtr(true),
			LibraryUpdate:        typeconv.BoolPtr(false),
			LibraryDelete:        typeconv.BoolPtr(false),
			UserInvite:           typeconv.BoolPtr(false),
			UserRemove:           typeconv.BoolPtr(false),
			UserPermissionManage: typeconv.BoolPtr(false),
		},
	}); err != nil {
		return nil, err
	}

	return insertedLibrary, nil
}

// CreateDefaultLibrary create a default library
func (s *Service) CreateDefaultLibrary(userID primitive.ObjectID) error {
	library := data.Library{
		Name:        "Bookshelf",
		Description: "The default library",
	}
	if _, err := s.CreateLibrary(library, userID); err != nil {
		return err
	}
	return nil
}

// ReadLibrary read a library
func (s *Service) ReadLibrary(libraryID primitive.ObjectID) (*data.Library, error) {
	return s.repo.Read(libraryID)
}

// DeleteLibrary delete a library
func (s *Service) DeleteLibrary(libraryID primitive.ObjectID) error {
	// Check if all data exist before deleting
	if _, err := s.repo.Read(libraryID); err != nil {
		return err
	}
	if _, err := s.userLibraryRepo.ReadFromLibraryID(libraryID); err != nil {
		return err
	}

	if err := s.repo.Delete(libraryID); err != nil {
		return err
	}
	if err := s.userLibraryRepo.Delete(libraryID); err != nil {
		return err
	}
	_ = s.bookProgressRepo.Delete(data.BookProgressData{LibraryID: libraryID})
	return nil
}
