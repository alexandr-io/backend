package grpcclient

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/grpc"
	grpclibrary "github.com/alexandr-io/backend/grpc/library"

	"github.com/gofiber/fiber/v2"
)

// CreateLibrary get a data.User containing an ID or an email and return the complete user data
func CreateLibrary(userID string) error {
	if libraryClient == nil {
		go InitClients()
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, "gRPC library client not initialized")
	}

	createLibraryRequest := grpclibrary.CreateLibraryRequest{
		UserID: userID,
	}
	fmt.Printf("[gRPC]: Login sent: %+v\n", createLibraryRequest.String())
	_, err := libraryClient.CreateLibrary(context.Background(), &createLibraryRequest)
	if err != nil {
		return grpc.ErrorToFiber(err)
	}
	return nil
}
