package models

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserPayload struct {
	User User `json:"user"`
}

type TokenResponse struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			Token string `json:"token"`
		} `json:"attributes"`
	} `json:"data"`
}
