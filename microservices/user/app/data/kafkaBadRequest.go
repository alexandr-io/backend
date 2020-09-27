package data

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alexandr-io/berrors"
)

// KafkaBadRequest is the JSON struct used in kafka communication in case of a bad request (e.g. username already taken).
type KafkaBadRequest struct {
	UUID string              `json:"uuid"`
	Data KafkaBadRequestData `json:"data"`
}

// KafkaBadRequestData is the data containing the error description of a bad request.
type KafkaBadRequestData struct {
	Code    int              `json:"code"`
	Content berrors.BadInput `json:"content"`
}

// CreateKafkaBadRequestMessage return a JSON of KafkaBadRequest from an id (UUID) and a berrors.BadInput.
func CreateKafkaBadRequestMessage(id string, content berrors.BadInput) ([]byte, error) {
	message := KafkaBadRequest{
		UUID: id,
		Data: KafkaBadRequestData{
			Code:    http.StatusBadRequest,
			Content: content,
		},
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	return messageJSON, err
}
