package data

import (
	"encoding/json"
	"log"

	"github.com/alexandr-io/berrors"
)

// KafkaResponseMessage is used to get the Code from a kafka message.
type KafkaResponseMessage struct {
	Code    int         `json:"code"`
	Content interface{} `json:"content"`
}

// KafkaResponseMessageBadRequest is used to get the Content of a kafka message.
type KafkaResponseMessageBadRequest struct {
	Code    int              `json:"code"`
	Content berrors.BadInput `json:"content"`
}

// GetBadInputJSON return a marshal JSON of berrors.BadInput from a kafka message.
func GetBadInputJSON(rawMessage []byte) ([]byte, error) {
	var badRequest KafkaResponseMessageBadRequest
	if err := json.Unmarshal(rawMessage, &badRequest); err != nil {
		log.Println(err)
		return nil, err
	}
	badRequestJSON, err := json.Marshal(badRequest.Content)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return badRequestJSON, nil
}
