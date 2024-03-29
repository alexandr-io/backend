package server

import (
	"context"
	"fmt"
	"net/url"

	"github.com/alexandr-io/backend/grpc"
	grpclibrary "github.com/alexandr-io/backend/grpc/library"
	"github.com/alexandr-io/backend/library/data"
	bookserv "github.com/alexandr-io/backend/library/internal/book"

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
	bookData, err := bookserv.Serv.ReadFromID(bookID)
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(err)
	}

	coverURLS := bookData.Thumbnails
	urlCover, err := url.Parse(in.GetCoverURL())
	if err != nil {
		return nil, grpc.FiberErrorToGRPC(data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error()))
	}
	for i, link := range bookData.Thumbnails {
		u, err := url.Parse(link)
		if err != nil {
			return nil, grpc.FiberErrorToGRPC(data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error()))
		}
		if u.Host == urlCover.Host {
			coverURLS = append(coverURLS[:i], coverURLS[i+1:]...)
			break
		}
	}
	coverURLS = append(coverURLS, in.GetCoverURL())

	fmt.Println(coverURLS)
	if _, err = bookserv.Serv.UpdateBook(data.Book{ID: bookID, Thumbnails: coverURLS}); err != nil {
		return nil, grpc.FiberErrorToGRPC(err)
	}

	return &empty.Empty{}, nil
}
