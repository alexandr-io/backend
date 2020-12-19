package data

import (
	"encoding/json"
)

// KafkaAuthResponse is the data sent by the auth MS to inform of the validity of a jwt.
type KafkaAuthResponse struct {
	Code    int       `json:"code"`
	Content KafkaUser `json:"content"`
}

// CreateAuthResponseMessage return a JSON of KafkaAuthResponse containing a http code and a KafkaUser.
func CreateAuthResponseMessage(code int, content KafkaUser) ([]byte, error) {
	message := KafkaAuthResponse{
		Code:    code,
		Content: content,
	}

	return json.Marshal(message)
}
