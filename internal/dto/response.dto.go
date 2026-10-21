package dto

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}

type ErrorResponse struct {
	Succes  bool   `json:"success"`
	Message string `json:"message"`
}
