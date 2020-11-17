package data

// UserLogin is the body parameter given to login a new user.
type UserLogin struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}
