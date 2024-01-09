// utils/response.go
package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents a standard error response structure
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// RespondJSON sends a JSON response with the specified status code and data
func RespondJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"status": "success", "data": data})
}

// RespondError sends a JSON error response with the specified status code and message
func RespondError(c *gin.Context, statusCode int, message string) {
	errorResponse := ErrorResponse{Status: statusCode, Message: message}
	c.JSON(statusCode, errorResponse)
}

// RespondBadRequest sends a JSON error response for bad request (400)
func RespondBadRequest(c *gin.Context, message string) {
	RespondError(c, http.StatusBadRequest, message)
}

// RespondUnauthorized sends a JSON error response for unauthorized access (401)
func RespondUnauthorized(c *gin.Context, message string) {
	RespondError(c, http.StatusUnauthorized, message)
}

// RespondNotFound sends a JSON error response for resource not found (404)
func RespondNotFound(c *gin.Context, message string) {
	RespondError(c, http.StatusNotFound, message)
}

// RespondInternalServerError sends a JSON error response for internal server error (500)
func RespondInternalServerError(c *gin.Context, message string) {
	RespondError(c, http.StatusInternalServerError, message)
}

// RespondWithSuccess sends a JSON success response with the specified status code and data
func RespondWithSuccess(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"status": "success", "data": data})
}
