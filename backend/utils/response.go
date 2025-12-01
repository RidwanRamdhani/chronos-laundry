package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SuccessResponse sends a successful response
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse sends an error response
func ErrorResponse(c *gin.Context, statusCode int, message string, err string) {
	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Error:   err,
	})
}

// BadRequest sends a 400 error
func BadRequest(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusBadRequest, "Bad Request", message)
}

// Unauthorized sends a 401 error
func Unauthorized(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", message)
}

// Forbidden sends a 403 error
func Forbidden(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusForbidden, "Forbidden", message)
}

// NotFound sends a 404 error
func NotFound(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, "Not Found", message)
}

// InternalServerError sends a 500 error
func InternalServerError(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", message)
}

// Conflict sends a 409 error
func Conflict(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusConflict, "Conflict", message)
}
