package data

// UserRegister is the body parameter given to register a new user to the database.
type UserRegister struct {
	Email           string  `json:"email" validate:"required,email"`
	Username        string  `json:"username" validate:"required"`
	Password        string  `json:"password" validate:"required"`
	ConfirmPassword string  `json:"confirm_password" validate:"required"`
	InvitationToken *string `json:"invitation_token,omitempty" validate:"required,len=10"`
}
