package res

import "strings"

// responce struct for the fomation
type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Errors     interface{} `json:"errors,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

// If we Get any error tihis will be show in json format
func ErrorResponse(statusCode int, message string, err string, data interface{}) Response {

	spiltedError := strings.Split(err, "\n")

	return Response{ // calling the responce struct
		StatusCode: statusCode,
		Message:    message,
		Errors:     spiltedError,
		Data:       data,
	}
}

// Function for SuccessResponse in json format
func SuccessResponse(statusCode int, message string, data interface{}) Response {

	return Response{ // calling the responce struct
		StatusCode: statusCode,
		Message:    message,
		Errors:     nil,
		Data:       data,
	}
}
