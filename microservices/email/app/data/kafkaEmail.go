package data

// KafkaEmail is the data sent in kafka to create and send an email
type KafkaEmail struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Type     string `json:"type"`
	Data     string `json:"data"`
}
