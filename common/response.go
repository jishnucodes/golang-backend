package common

import (
	"github.com/gin-gonic/gin"
)

// Standard response format
type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`    
	Data    interface{} `json:"data"`      
	Error   string      `json:"error,omitempty"` 
}

// SendSuccess creates a standard success response
func SendSuccess(ctx *gin.Context, status int, message string, data interface{}) {
	
	ctx.JSON(status, APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

// SendError creates a standard error response
func SendError(ctx *gin.Context, status int, message string, err error) {
	ctx.JSON(status, APIResponse{
		Status:  status,
		Message: message,
		Error:   err.Error(),
	})
}
