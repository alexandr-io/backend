package data

type KafkaMail struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Type     string `json:"type"`
	Data     string `json:"data"`
}
