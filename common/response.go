package common

import (
	"github.com/gin-gonic/gin"
)

// Standard response format
type APIResponse struct {
	Status  bool         `json:"status"`
	StatusCode int  	`json:"statusCode"`
	Message string      `json:"message"`    
	Data    interface{} `json:"data"`      
	Error   string      `json:"error,omitempty"` 
}

// SendSuccess creates a standard success response
func SendSuccess(ctx *gin.Context, httpStatus int, status int, message string, data interface{}) {
	// Ensure status is either `true` or `false`
	statusBool := status == 1

	ctx.JSON(httpStatus, APIResponse{
		Status:     statusBool,
		StatusCode: httpStatus,
		Message:    message,
		Data:       data,
	})
}


// SendError creates a standard error response
func SendError(ctx *gin.Context, httpStatus int, status int, message string, err error) {
	statusBool := status == 1

	// Ensure err is not nil before calling err.Error()
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}

	ctx.JSON(httpStatus, APIResponse{
		Status:     statusBool,
		StatusCode: httpStatus,
		Message:    message,
		Error:      errorMessage,
	})
}

