package common

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"clinic-management/backend/spResponse" // Import the response struct

	"github.com/gin-gonic/gin"
)

// HandleManagerError checks for errors in userManager responses and sends an appropriate response.
// func HandleManagerError(ctx *gin.Context, response *spResponse.Result, err error, action string) bool {
// 	if err != nil {
// 		log.Printf("%s failed: %v", action, err) // Log error with context
// 		SendError(ctx, http.StatusInternalServerError, response.Status, response.StatusMessage, err)
// 		return true // Return true to indicate an error occurred
// 	}
// 	return false // Return false if there's no error
// }

// HandleJSONParseError logs and sends a response if JSON parsing fails.
func HandleRequestError(ctx *gin.Context, err error, message string) error {
	if err != nil {
		log.Println(message, err) // Log the error with a custom message
		SendError(ctx, http.StatusBadRequest, 0, message, err) // Send error response
		return err // Return the error so the caller knows parsing failed
	}
	return nil // Return nil if there's no error
}

func HandleServerError(ctx *gin.Context, managerResponse *spResponse.Result, err error) bool {
	if managerResponse == nil {
		err := fmt.Errorf("managerResponse is nil")
		fmt.Println("working server error -1 on deleting")
		log.Println("Error:", err)
		SendError(ctx, http.StatusInternalServerError, 0, "Internal Server Error", err)
		return true
	}

	// If there's no provided error, use the response message
	if err == nil && managerResponse.Status == 0 {
		err = fmt.Errorf(managerResponse.StatusMessage)
		fmt.Println("working server error -2 on deleting")
		log.Println("Error:", err)
		SendError(ctx, http.StatusInternalServerError, managerResponse.Status, managerResponse.StatusMessage, err)
		return true
	}else if err != nil && managerResponse.Status == 0 {
		fmt.Println("working server error -3 on deleting")
		log.Println("Error:", err) // Log the error with a custom message
		SendError(ctx, http.StatusBadRequest, managerResponse.Status, managerResponse.StatusMessage, err) // Send error response
		return true
	}
	fmt.Println("working server error -4 on deleting")
	return false
	

}


func GetParamAsUint(ctx *gin.Context, paramName string) (uint, error) {
	paramValue, ok := ctx.Params.Get(paramName)
	if !ok {
		err := fmt.Errorf("%s is required", paramName)
		log.Println("Error:", err)
		// SendError(ctx, http.StatusBadRequest, 0, fmt.Sprintf("%s is missing in the request", paramName), err)
		HandleRequestError(ctx, err, fmt.Sprintf("%s is missing in the request", paramName)); 
		return 0, err
	}

	// Convert paramValue to uint
	parsedValue, err := strconv.ParseUint(paramValue, 10, 32)
	if err != nil {
		err = fmt.Errorf("invalid %s format", paramName)
		log.Println("Error:", err)
		// SendError(ctx, http.StatusBadRequest, 0, fmt.Sprintf("invalid %s format", paramName), err)
		HandleRequestError(ctx, err, fmt.Sprintf("invalid %s format", paramName))
		return 0, err
	}

	return uint(parsedValue), nil
}

