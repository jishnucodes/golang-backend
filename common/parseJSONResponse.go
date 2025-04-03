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
func ParseJSONResponse(managerResponse *spResponse.Result, ctx *gin.Context) ([]map[string]interface{}, error) {
	var parsedData []map[string]interface{}

	// Check if Data is not empty and is valid JSON
	if managerResponse.Data != "" && json.Valid([]byte(managerResponse.Data)) {
		err := json.Unmarshal([]byte(managerResponse.Data), &parsedData)
		if err != nil {
			log.Println("Failed to parse user data:", err)
			// SendError(ctx, http.StatusBadRequest, 0, "Failed to parse user data", err)
			if errors := HandleRequestError(ctx, err, "Failed to parse user data"); errors != nil {
				return nil, errors // Exit if an error occurred (response is already sent)
			}
			
			return nil, err
		}
		return parsedData, nil
	}

	// Handle invalid response
	// SendError(
	// 	ctx,
	// 	http.StatusInternalServerError,
	// 	managerResponse.Status,
	// 	managerResponse.StatusMessage,
	// 	fmt.Errorf(managerResponse.StatusMessage),
	// )
	if err := HandleHTTPError(ctx, managerResponse); err != nil {
		return nil, err// Exit if an error occurred (response is already sent)
	}
	
	
	return nil, fmt.Errorf(managerResponse.StatusMessage)
}
