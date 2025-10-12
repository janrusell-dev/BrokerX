package utils

import "github.com/gin-gonic/gin"

// SuccessResponse sends a standardized success response
func SuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"status": "success",
		"data":   data,
	})
}

// ErrorResponse sends a standardized error response
func ErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	response := gin.H{
		"status":  "error",
		"message": message,
	}

	if err != nil {
		response["error"] = err.Error()
	}

	c.JSON(statusCode, response)
}

// ValidationErrorResponse sends a validation error response
func ValidationErrorResponse(c *gin.Context, errors map[string]string) {
	c.JSON(400, gin.H{
		"status":  "error",
		"message": "Validation failed",
		"errors":  errors,
	})
}

// MessageResponse sends a simple message response
func MessageResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"status":  "success",
		"message": message,
	})
}
