package data

import (
	"encoding/json"
	"log"

	"github.com/alexandr-io/berrors"
)

// KafkaResponseMessageUUIDData is used to get the UUID from a kafka message.
type KafkaResponseMessageUUIDData struct {
	UUID string      `json:"uuid"`
	Data interface{} `json:"data"`
}

// KafkaResponseMessage is used to get the Code from a kafka message.
type KafkaResponseMessage struct {
	UUID string `json:"uuid"`
	Data struct {
		Code    int         `json:"code"`
		Content interface{} `json:"content"`
	} `json:"data"`
}

// KafkaResponseMessageBadRequest is used to get the Content of a kafka message.
type KafkaResponseMessageBadRequest struct {
	UUID string `json:"uuid"`
	Data struct {
		Code    int              `json:"code"`
		Content berrors.BadInput `json:"content"`
	} `json:"data"`
}

// GetBadInputJSON return a marshal JSON of berrors.BadInput from a kafka message.
func GetBadInputJSON(rawMessage []byte) ([]byte, error) {
	var badRequest KafkaResponseMessageBadRequest
	if err := json.Unmarshal(rawMessage, &badRequest); err != nil {
		log.Println(err)
		return nil, err
	}
	badRequestJSON, err := json.Marshal(badRequest.Data.Content)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return badRequestJSON, nil
}
