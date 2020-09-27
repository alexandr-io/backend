package data

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alexandr-io/berrors"
)

type KafkaBadRequest struct {
	UUID string              `json:"uuid"`
	Data KafkaBadRequestData `json:"data"`
}

type KafkaBadRequestData struct {
	Code    int              `json:"code"`
	Content berrors.BadInput `json:"content"`
}

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
