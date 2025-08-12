// # Common JSON response formatting
package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// APIResponse - common response structure
type APIResponse struct {
	Status  string      `json:"status"`          // success or error
	Message string      `json:"message"`         // message for the client
	Data    interface{} `json:"data,omitempty"`  // actual data (optional)
	Error   string      `json:"error,omitempty"` // error details (optional)
}

// SuccessResponse - send success response
func SuccessResponse(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// ErrorResponse - send error response
func ErrorResponse(c echo.Context, message string, err error, statusCode int) error {
	return c.JSON(statusCode, APIResponse{
		Status:  "error",
		Message: message,
		Error:   err.Error(),
	})
}
