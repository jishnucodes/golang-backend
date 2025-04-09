package errorHandlers

import (
	"clinic-management/backend/common"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleErrorResponse checks if the result contains an error message, logs it, and sends an error response
func HandleErrorResponse(ctx *gin.Context, result []map[string]interface{}, err error) error{
    // Check if result contains at least one item
    if len(result) > 0 {
        firstResult := result[0] // Get the first map in the slice

        // Check if "Message" key exists in the map
        if errorRaw, exists := firstResult["Message"]; exists {
            if errorMessage, ok := errorRaw.(string); ok {
                log.Println("Login failed:", errorMessage)

                // If there is an error, set err to the message
                if err == nil {
                    err = fmt.Errorf("%v", errorMessage) // Create a new error with the message
                }

                // Send the error response with the message
                common.SendError(ctx, http.StatusUnauthorized, 0, errorMessage, err)
                return err
            } else {
                // If "Message" is not a string, log the unexpected type
                log.Println("Unexpected type for Message:", errorRaw)
            }
        }
    } 
	return nil 
}
