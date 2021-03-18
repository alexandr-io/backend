package data

// UpdatePassword is the data used for the PUT route to update a password as a logged user
type UpdatePassword struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

// UserUpdatePassword is the data used to send to the User MS via gRPC to update a password for a logged user
type UserUpdatePassword struct {
	UserID          string
	CurrentPassword string
	NewPassword     string
}
