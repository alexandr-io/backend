package data

// KafkaUserResponse is the data used for a success response in kafka for a registration.
type KafkaUserResponse struct {
	Code    int                      `json:"code"`
	Content KafkaUserResponseContent `json:"content"`
}

// KafkaUserResponseContent contain the user fields of a success response of a registration.
type KafkaUserResponseContent struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
