package data

// UserSendResetPasswordEmail is the data sent to the /password/reset POST route
type UserSendResetPasswordEmail struct {
	Email string `json:"email" validate:"required,email"`
}
