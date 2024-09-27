package handler

// ResponseData is the response structure for the application.
type ResponseData struct {
	// Data is the response data.
	Data interface{} `json:"data,omitempty"`
}

// ResponseError is the response structure for the application.
type ResponseError struct {
	// Errors is the response errors.
	Errors []Error `json:"errors,omitempty"`
}

// Error is the error structure for the application.
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
