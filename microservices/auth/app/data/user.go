package data

// User defines the structure for an API user
type User struct {
	ID           string `json:"-" bson:"_id,omitempty"`
	Username     string `json:"username,omitempty" bson:"username,omitempty"`
	Email        string `json:"email,omitempty" bson:"email,omitempty"`
	AuthToken    string `json:"auth_token,omitempty" bson:"-"`
	RefreshToken string `json:"refresh_token,omitempty" bson:"-"`
}
