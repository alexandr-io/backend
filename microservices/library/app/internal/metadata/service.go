package permission

import (
	"github.com/alexandr-io/backend/library/data"
	grpcclient "github.com/alexandr-io/backend/library/grpc/client"
)

// Serv instance of metadata service
var Serv *Service

// Service is the struct containing database repository needed for metadata methods of the interface
type Service struct{}

// NewService create and set instance of Service
func NewService() *Service {
	Serv = &Service{}
	return Serv
}

// RequestMetadata request book metadata to metadata MS
func (s *Service) RequestMetadata(title string, authors string) (*data.Book, error) {
	response, err := grpcclient.Metadata(title, authors)
	if err != nil {
		return nil, err
	}
	return response, nil
}
