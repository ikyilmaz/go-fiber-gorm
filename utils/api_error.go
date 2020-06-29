package utils

import (
	"net/http"
	"time"
)

type APIError struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	StatusText string `json:"statusText"`
	Timestamp  int64  `json:"timestamp"`
}

func NewAPIError(statusCode int, msg string) *APIError {
	status := "fail"

	if statusCode >= 500 {
		status = "error"
	}

	return &APIError{status, msg, statusCode, http.StatusText(statusCode), time.Now().Unix()}
}

func (a *APIError) Error() string { return a.Message }

func getDefaultMessage(defaultMsg string, msg ...string) string {
	if len(msg) > 0 {
		return msg[0]
	}

	return defaultMsg
}

func BadRequest(msg ...string) *APIError {
	return NewAPIError(400, getDefaultMessage("Bad Request", msg...))
}

func NotFound(msg ...string) *APIError {
	return NewAPIError(404, getDefaultMessage("Not Found", msg...))
}

func Forbidden(msg ...string) *APIError {
	return NewAPIError(403, getDefaultMessage("Forbidden", msg...))
}
