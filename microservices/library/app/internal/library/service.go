package library

import (
	"github.com/alexandr-io/backend/common/typeconv"
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/data/permissions"
	bookProgressServ "github.com/alexandr-io/backend/library/internal/book_progress"
	groupServ "github.com/alexandr-io/backend/library/internal/group"
	userLibraryServ "github.com/alexandr-io/backend/library/internal/user_library"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Serv *Service

type Service struct {
	repo             Repository
	userLibraryRepo  userLibraryServ.Repository
	groupRepo        groupServ.Repository
	bookProgressRepo bookProgressServ.Repository
}

func NewService(repo Repository, userLibrary userLibraryServ.Repository, group groupServ.Repository, bookProgress bookProgressServ.Repository) *Service {
	Serv = &Service{repo: repo, userLibraryRepo: userLibrary, groupRepo: group, bookProgressRepo: bookProgress}
	return Serv
}

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

func (s *Service) CreateDefaultLibrary(userID primitive.ObjectID) error {
	library := data.Library{
		Name:        "Bookshelf",
		Description: "The default library",
	}
	if _, err := s.repo.Create(library); err != nil {
		return err
	}
	return nil
}

func (s *Service) ReadLibrary(libraryID primitive.ObjectID) (*data.Library, error) {
	return s.repo.Read(libraryID)
}

func (s *Service) DeleteLibrary(libraryID primitive.ObjectID) error {
	// TODO: create logic so that when library delete failed, the book previously deleted in restored
	if err := s.repo.Delete(libraryID); err != nil {
		return err
	}
	s.userLibraryRepo.Delete(libraryID)
	s.bookProgressRepo.Delete(data.BookProgressData{LibraryID: libraryID})
	return nil
}
