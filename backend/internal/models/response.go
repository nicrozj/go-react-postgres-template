package models

type Response struct {
	Success bool `json:"success"`
	Status  int  `json:"status"`
	Data    any  `json:"data,omitempty,omitzero"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewErrorResponse(status int, err error) *ErrorResponse {
	return &ErrorResponse{
		Success: false,
		Status:  status,
		Error:   err.Error(),
	}
}
