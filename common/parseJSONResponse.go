package common

import (
	"clinic-management/backend/spResponse"
	"encoding/json"
	"fmt"
	"log"
	// "net/http"

	"github.com/gin-gonic/gin"
)

// ParsedData struct (Define it properly)


// ParseJSONResponse parses a JSON response and extracts data from the "data" field
func ParseJSONResponse(managerResponse *spResponse.Result, ctx *gin.Context) ([]map[string]interface{}) {
	var parsedData []map[string]interface{}

	// Check if Data is not empty and is valid JSON
	if managerResponse.Data != "nil" && json.Valid([]byte(managerResponse.Data)) {
		err := json.Unmarshal([]byte(managerResponse.Data), &parsedData)
		if err != nil {
			log.Println("Failed to parse user data:", err)
			// SendError(ctx, http.StatusBadRequest, 0, "Failed to parse user data", err)
			if errors := HandleRequestError(ctx, err, "Failed to parse user data"); errors != nil {
				return nil // Exit if an error occurred (response is already sent)
			}
			
			return nil
		}
		return parsedData
	}

	fmt.Println("working this when deleting")
	return nil
}
