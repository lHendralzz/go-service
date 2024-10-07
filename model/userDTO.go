package model

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	Token string `json:"token"`
}
