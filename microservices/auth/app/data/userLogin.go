package data

// UserLogin is the body parameter given to login a new user.
// swagger:model
type UserLogin struct {
	// The email or the username of the user
	// required: true
	// example: john@provider.net
	Login string `json:"login" validate:"required"`
	// The password of the user
	// required: true
	// example: leHAiOjE1OTgzNz
	Password string `json:"password" validate:"required"`
}
