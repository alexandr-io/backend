package grpcclient

import (
	"context"
	"fmt"

	"github.com/alexandr-io/backend/auth/data"
	"github.com/alexandr-io/backend/grpc"
	grpcemail "github.com/alexandr-io/backend/grpc/email"

	"github.com/gofiber/fiber/v2"
)

// SendEmail create and send an email from the given information
func SendEmail(email data.Email) error {
	if emailClient == nil {
		go InitClients()
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, "gRPC email client not initialized")
	}

	sendEmailRequest := grpcemail.SendEmailRequest{
		Email:    email.Email,
		Username: email.Username,
		Type:     email.Type,
		Data:     email.Data,
	}
	fmt.Printf("[gRPC]: Send email sent: %+v\n", sendEmailRequest.String())
	_, err := emailClient.SendEmail(context.Background(), &sendEmailRequest)
	if err != nil {
		return grpc.ErrorToFiber(err)
	}
	return nil
}
