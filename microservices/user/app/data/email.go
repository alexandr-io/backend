package data

// Email is the data needed to send an email
type Email struct {
	Email    string
	Username string
	Type     string
	Data     string
}

type EmailVerification struct {
	OldEmail string `json:"old_email"`
	NewEmail string `json:"new_email"`
}
