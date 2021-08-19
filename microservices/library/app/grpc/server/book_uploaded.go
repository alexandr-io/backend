package server

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/grpc"
	grpclibrary "github.com/alexandr-io/backend/grpc/library"
	"github.com/alexandr-io/backend/library/data"
	bookserv "github.com/alexandr-io/backend/library/internal/book"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BookUploaded is a gRPC server method that take a book ID and the book type of an uploaded book. Used to store book file type url in metadata
func (s *server) BookUploaded(_ context.Context, in *grpclibrary.BookUploadedRequest) (*empty.Empty, error) {
	fmt.Printf("[gRPC]: Book uploaded received: %+v\n", in.String())

	bookID, err := primitive.ObjectIDFromHex(in.GetBookID())
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(data.NewHTTPErrorInfo(fiber.StatusBadRequest, "Book ID is incorrect"))
	}

	if _, err = bookserv.Serv.UpdateBook(data.Book{ID: bookID, FileType: in.Type}); err != nil {
		return nil, grpc.FiberErrorToGRPC(err)
	}

	return &empty.Empty{}, nil
}
