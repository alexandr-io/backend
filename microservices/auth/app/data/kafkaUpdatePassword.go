package data

import "encoding/json"

// KafkaUpdatePassword is the JSON struct sent to the user MS using the kafka topic `user.password.update`.
type KafkaUpdatePassword struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

// Marshall marshall a KafkaUpdatePassword
func (update *KafkaUpdatePassword) Marshall() ([]byte, error) {
	return json.Marshal(update)
}
