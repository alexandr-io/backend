package data

import "encoding/json"

// KafkaAuthRequest is the data sent to the auth MS to check the validity of a jwt.
type KafkaAuthRequest struct {
	JWT string `json:"jwt"`
}

// CreateAuthRequestMessage return a JSON of KafkaAuthRequest containing a jwt.
func CreateAuthRequestMessage(jwt string) ([]byte, error) {
	message := KafkaAuthRequest{
		JWT: jwt,
	}

	return json.Marshal(message)
}
