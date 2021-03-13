package server

import (
	"context"
	"fmt"

	grpcemail "github.com/alexandr-io/backend/grpc/email"
	"github.com/alexandr-io/backend/mail/data"
	"github.com/alexandr-io/backend/mail/internal"

	"github.com/golang/protobuf/ptypes/empty"
)

// SendEmail is a gRPC server method that take an email information to send one
func (s *server) SendEmail(_ context.Context, in *grpcemail.SendEmailRequest) (*empty.Empty, error) {
	fmt.Printf("[gRPC]: Send email received: %+v\n", in.String())
	internal.CreateMailFromMessage(data.Email{
		Email:    in.GetEmail(),
		Username: in.GetUsername(),
		Type:     in.GetType(),
		Data:     in.GetData(),
	})
	return &empty.Empty{}, nil
}
