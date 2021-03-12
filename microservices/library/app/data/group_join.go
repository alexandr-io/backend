package data

// GroupJoinData contain the user's IS of the user to add to the designed group
type GroupJoinData struct {
	UserID string `json:"user_id" validation:"required"`
}
