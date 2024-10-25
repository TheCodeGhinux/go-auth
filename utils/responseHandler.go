package utils

import (
	"github.com/gin-gonic/gin"
)

// ErrorResponse defines the structure of the error response
type ErrorResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"message"`
}

type ApiResponse struct {
	Message    string      `json:"message"`
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
}


// SendError formats and sends a standardized error response
func SendError(c *gin.Context, statusCode int, message interface{}) {
	c.JSON(statusCode, ErrorResponse{
		Status:     "Error",
		StatusCode: statusCode,
		Message:    message,
	})
}

func RespondHandler(c *gin.Context, message string, statusCode int, data interface{}) {
	response := ApiResponse{
		Message:    message,
		Status:     "Sucess",
		StatusCode: statusCode,
		Data:       data,
	}
	c.JSON(statusCode, response)
}