package data

// UserSendResetPasswordEmail is the data sent to the /password/reset POST route
type UserSendResetPasswordEmail struct {
	Email string `json:"email" validate:"required,email"`
}

// UserResetPasswordToken is the data sent to the /password/reset GET route
type UserResetPasswordToken struct {
	Token string `json:"token" validate:"required,len=6"`
}
