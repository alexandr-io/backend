package data

import (
	"encoding/json"
	"log"

	"github.com/alexandr-io/berrors"
)

type KafkaResponseMessageUUIDData struct {
	UUID string      `json:"uuid"`
	Data interface{} `json:"data"`
}

type KafkaResponseMessage struct {
	UUID string `json:"uuid"`
	Data struct {
		Code    int         `json:"code"`
		Content interface{} `json:"content"`
	} `json:"data"`
}

type KafkaResponseMessageBadRequest struct {
	UUID string `json:"uuid"`
	Data struct {
		Code    int              `json:"code"`
		Content berrors.BadInput `json:"content"`
	} `json:"data"`
}

func GetBadInputJson(rawMessage []byte) ([]byte, error) {
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
