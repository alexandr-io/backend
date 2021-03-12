package data

// GroupJoin contain the user's ID of the user to add to the designed group
type GroupJoin struct {
	UserID string `json:"user_id" validation:"required"`
}
