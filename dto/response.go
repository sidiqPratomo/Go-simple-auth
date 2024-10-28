package dto

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string                   `json:"message"`
	Details []ValidationErrorDetails `json:"details,omitempty"`
}

type ValidationErrorDetails struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}