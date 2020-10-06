package data

// UserRegister is the body parameter given to register a new user to the database.
// swagger:model
type UserRegister struct {
	// The email of the user
	// required: true
	// example: john@provider.net
	Email string `json:"email" validate:"required,email"`
	// The username of the user
	// required: true
	// example: john
	Username string `json:"username" validate:"required"`
	// The password of the user
	// required: true
	// example: leHAiOjE1OTgzNz
	Password string `json:"password" validate:"required"`
	// The confirmation password of the user
	// required: true
	// example: leHAiOjE1OTgzNz
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}
