package common

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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
		log.Println(message, err)                              // Log the error with a custom message
		SendError(ctx, http.StatusBadRequest, 0, message, err) // Send error response
		return err                                             // Return the error so the caller knows parsing failed
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
	} else if err != nil && managerResponse.Status == 0 {
		fmt.Println("working server error -3 on deleting")
		log.Println("Error:", err)                                                                        // Log the error with a custom message
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
		HandleRequestError(ctx, err, fmt.Sprintf("%s is missing in the request", paramName))
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

// Safely convert to uint, handling nil and float64 values
// func ToUint(value interface{}) uint {
// 	if value == nil {
// 		return 0
// 	}
// 	switch v := value.(type) {
// 	case float64:
// 		return uint(v) // JSON unmarshals numbers as float64
// 	case int:
// 		return uint(v)
// 	}
// 	return 0
// }

func ToUint(value interface{}) uint {
	if value == nil {
		return 0
	}

	switch v := value.(type) {
	case uint:
		return v
	case int:
		return uint(v)
	case int64:
		return uint(v)
	case float64:
		return uint(v)
	case float32:
		return uint(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		if i, err := strconv.ParseUint(v, 10, 64); err == nil {
			return uint(i)
		}
	case json.Number:
		if i, err := v.Int64(); err == nil {
			return uint(i)
		}
	}

	return 0
}


// Safely convert to string
func ToString(value interface{}) string {
	if value == nil {
		return ""
	}
	if v, ok := value.(string); ok {
		return v
	}
	return ""
}

// Safely decode base64-encoded biometric data
func DecodeBase64(value string) []byte {
	if value == "" {
		return nil
	}
	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return nil
	}
	return data
}

// Safely parse time.Time from RFC3339 string
func ParseTime(value interface{}) time.Time {
	if value == nil {
		fmt.Println("value is nil")
		return time.Time{}
	}
	if str, ok := value.(string); ok {
		fmt.Println("str is string")
		t, err := time.Parse(time.RFC3339, str)
		if err == nil {
			return t
		}
	}
	if str, ok := value.(string); ok {
		layout := "15:04:05"
		fmt.Println("str is string")
		t, err := time.Parse(layout, str)
		if err == nil {
			return t
		}
	}
	fmt.Println("time is nil")
	return time.Time{}
}

// Safely convert to float64
func ToFloat64(value interface{}) float64 {
	if value == nil {
		return 0
	}
	switch v := value.(type) {
	case float64:
		return v
	case int:
		return float64(v)
	case string:
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
	}
	return 0
}

func ToInt(value interface{}) int {
	if value == nil {
		return 0
	}
	switch v := value.(type) {
	case int:
		return v
	case string:
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return 0
}


func ToBool(value interface{}) bool {
	if value == nil {
		return false
	}
	switch v := value.(type) {
	case bool:
		return v
	case string:
		if b, err := strconv.ParseBool(v); err == nil {
			return b
		}
	}
	return false
}
