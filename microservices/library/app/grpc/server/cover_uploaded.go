package server

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/grpc"
	grpclibrary "github.com/alexandr-io/backend/grpc/library"
	"github.com/alexandr-io/backend/library/data"
	"github.com/alexandr-io/backend/library/database/book"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CoverUploaded is a gRPC server method that take a book ID and the url to an uploaded cover. Used to store book cover url in metadata
func (s *server) CoverUploaded(_ context.Context, in *grpclibrary.CoverUploadedRequest) (*empty.Empty, error) {
	fmt.Printf("[gRPC]: Cover uploaded received: %+v\n", in.String())

	bookID, err := primitive.ObjectIDFromHex(in.GetBookID())
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(data.NewHTTPErrorInfo(fiber.StatusBadRequest, "Book ID is incorrect"))
	}
	bookData, err := book.GetFromID(bookID)
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(err)
	}

	coverURLS := append(bookData.Thumbnails, in.GetCoverURL())
	fmt.Println(coverURLS)
	if _, err = book.Update(data.Book{ID: bookID, Thumbnails: coverURLS}); err != nil {
		return nil, grpc.FiberErrorToGRPC(err)
	}

	return &empty.Empty{}, nil
}
