package dto

type LoginResponse struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
}

type RegisterResponse struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
}
