package permission

import (
	"github.com/alexandr-io/backend/library/data"
	grpcclient "github.com/alexandr-io/backend/library/grpc/client"
)

var Serv *Service

type Service struct{}

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
