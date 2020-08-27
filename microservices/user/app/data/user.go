package data

// User defines the structure for an API user
// swagger:model
type User struct {
	ID string `json:"-" bson:"_id,omitempty"`
	// The username of this user
	// required: true
	// example: john
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	// The email address of this user
	// required: true
	// example: john@provider.net
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"-" bson:"password,omitempty"`
	// The authentication token of this user
	// example: eyJhb[...]FYqf4
	JWT string `json:"jwt" bson:"-"`
}
