package data

type UserSendResetPasswordEmail struct {
	Email string `json:"email" validate:"required,email"`
}
