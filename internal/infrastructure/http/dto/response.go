package dto

type Response struct {
	Data any `json:"data"`
}

type FailedResponse struct {
	Message string `json:"message"`
}
