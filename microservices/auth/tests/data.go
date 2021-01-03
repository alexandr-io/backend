package tests

type user struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	AuthToken    string `json:"auth_token"`
	RefreshToken string `json:"refresh_token"`
}

type invitation struct {
	Token string `json:"token"`
}
