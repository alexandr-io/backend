package data

// User defines the structure for an API user
// swagger:model
type User struct {
	ID string `json:"-" bson:"_id,omitempty"`
	// The username of this user
	// required: true
	// example: john
	Username string `json:"username" bson:"username,omitempty"`
	// The email address of this user
	// required: true
	// example: john@provider.net
	Email    string `json:"email" bson:"email,omitempty"`
	Password string `json:"-" bson:"password,omitempty"`
	// The authentication token of this user. Valid for 15 minutes.
	// example: eyJhb[...]FYqf4
	AuthToken string `json:"auth_token,omitempty" bson:"-"`
	// The refresh token of this user. Valid for 30 days. Used to get a new auth token.
	// example: eyJhb[...]FYqf4
	RefreshToken string `json:"refresh_token,omitempty" bson:"-"`
}
