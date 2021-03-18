package data

// User defines the structure for an API user
type User struct {
	ID            string `json:"-" bson:"_id,omitempty"`
	Username      string `json:"username" bson:"username,omitempty"`
	Email         string `json:"email" bson:"email,omitempty"`
	EmailVerified bool   `json:"email_verified" bson:"email_verified,omitempty"`
	Password      string `json:"-" bson:"password,omitempty"`
	AuthToken     string `json:"auth_token,omitempty" bson:"-"`
	RefreshToken  string `json:"refresh_token,omitempty" bson:"-"`
}

// UserUpdate defines the structure for an API update user
type UserUpdate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
