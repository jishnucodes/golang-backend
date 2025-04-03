package common

import (
	"fmt"
	"log"
	"net/http"

	"clinic-management/backend/spResponse" // Import the response struct

	"github.com/gin-gonic/gin"
)

// HandleManagerError checks for errors in userManager responses and sends an appropriate response.
func HandleManagerError(ctx *gin.Context, response *spResponse.Result, err error, action string) bool {
	if err != nil {
		log.Printf("%s failed: %v", action, err) // Log error with context
		SendError(ctx, http.StatusInternalServerError, response.Status, response.StatusMessage, err)
		return true // Return true to indicate an error occurred
	}
	return false // Return false if there's no error
}

// HandleJSONParseError logs and sends a response if JSON parsing fails.
func HandleRequestError(ctx *gin.Context, err error, message string) error {
	if err != nil {
		log.Println(message, err) // Log the error with a custom message
		SendError(ctx, http.StatusBadRequest, 0, message, err) // Send error response
		return err // Return the error so the caller knows parsing failed
	}
	return nil // Return nil if there's no error
}

func HandleHTTPError(ctx *gin.Context, managerResponse *spResponse.Result) error {
	if managerResponse == nil {
		err := fmt.Errorf("managerResponse is nil")
		log.Println("Error:", err)
		SendError(ctx, http.StatusInternalServerError, 0, "Internal Server Error", err)
		return err
	}

	err := fmt.Errorf(managerResponse.StatusMessage)
	log.Println("Error:", managerResponse.StatusMessage, err)
	SendError(ctx, http.StatusInternalServerError, managerResponse.Status, managerResponse.StatusMessage, fmt.Errorf(managerResponse.StatusMessage))
	
	return err
}
