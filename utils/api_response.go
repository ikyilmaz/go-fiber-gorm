package utils

import (
	"net/http"
	"time"
)

type APIResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	StatusText string      `json:"statusText"`
	Data       interface{} `json:"data"`
	Timestamp  int64       `json:"timestamp"`
}

func NewAPIResponse(statusCode int, data interface{}) *APIResponse {
	return &APIResponse{"success", statusCode, http.StatusText(statusCode), data, time.Now().Unix()}
}

func OK(data interface{}) *APIResponse        { return NewAPIResponse(200, data) }
func Created(data interface{}) *APIResponse   { return NewAPIResponse(201, data) }
func NoContent(data interface{}) *APIResponse { return NewAPIResponse(204, data) }
