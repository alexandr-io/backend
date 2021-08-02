package server

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/grpc"
	grpclibrary "github.com/alexandr-io/backend/grpc/library"
	"github.com/alexandr-io/backend/library/data"
	libraryServ "github.com/alexandr-io/backend/library/internal/library"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateLibrary is a gRPC server method that take a token and return user data. Used to check if the token is valid.
func (s *server) CreateLibrary(_ context.Context, in *grpclibrary.CreateLibraryRequest) (*empty.Empty, error) {
	fmt.Printf("[gRPC]: Create library received: %+v\n", in.String())
	userID, err := primitive.ObjectIDFromHex(in.GetUserID())
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(data.NewHTTPErrorInfo(fiber.StatusBadRequest, err.Error()))
	}

	if err = libraryServ.Serv.CreateDefaultLibrary(userID); err != nil {
		return nil, grpc.FiberErrorToGRPC(err)
	}

	return &empty.Empty{}, nil
}
