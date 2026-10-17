// Package httpx holds HTTP helpers shared by the API handlers.
package httpx

import "github.com/gin-gonic/gin"

// ErrorResponse is the body of every non-2xx API response.
type ErrorResponse struct {
	Error string `json:"error"`
}

// Error aborts the request with the given status and message.
func Error(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, ErrorResponse{Error: message})
}
