package common

import (
	"github.com/gin-gonic/gin"
)

// Standard response format
type APIResponse struct {
	Status  int         `json:"status"`    
	Data    interface{} `json:"data"`      
	Error   string      `json:"error,omitempty"` 
}

// SendSuccess creates a standard success response
func SendSuccess(ctx *gin.Context, status int, data interface{}) {
	
	ctx.JSON(status, APIResponse{
		Status:  status,
		Data:    data,
	})
}

// SendError creates a standard error response
func SendError(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, APIResponse{
		Status:  status,
		Error:   err.Error(),
	})
}
