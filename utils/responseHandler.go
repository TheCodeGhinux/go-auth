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

// SendError formats and sends a standardized error response
func SendError(c *gin.Context, statusCode int, message interface{}) {
	c.JSON(statusCode, ErrorResponse{
		Status:     "Error",
		StatusCode: statusCode,
		Message:    message,
	})
}
