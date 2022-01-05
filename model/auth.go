package model

// Authentication is for login data.
type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Token is for storing token information for correct login credentials.
type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}
