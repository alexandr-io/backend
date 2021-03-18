package data

type UpdatePassword struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type UserUpdatePassword struct {
	UserID          string `json:"user_ID"`
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}
