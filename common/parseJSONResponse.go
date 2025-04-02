package common

import (
	"encoding/json"
	"fmt"
)

// ParsedData holds both the outer and nested JSON data
type ParsedData struct {
	OuterData []map[string]interface{} `json:"outerData"`
	UserData  []map[string]interface{} `json:"userData"`
}

// ParseJSONResponse parses a JSON response and extracts data from the "data" field
func ParseJSONResponse(jsonStr string) (*ParsedData, error) {
	// Step 1: Unmarshal the outer JSON
	var outerData []map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &outerData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse outer JSON: %w", err)
	}

	// Return early if there's no data
	if len(outerData) == 0 {
		return &ParsedData{OuterData: nil, UserData: nil}, nil
	}

	// Step 2: Extract the "data" field and check its type
	rawData, exists := outerData[0]["data"]
	if !exists {
		return nil, fmt.Errorf("missing 'data' field in JSON")
	}

	// Step 3: Ensure "data" is an array of objects
	userData, ok := rawData.([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid format for 'data' field")
	}

	// Convert userData to []map[string]interface{}
	var userDataParsed []map[string]interface{}
	for _, item := range userData {
		if userMap, valid := item.(map[string]interface{}); valid {
			userDataParsed = append(userDataParsed, userMap)
		} else {
			return nil, fmt.Errorf("invalid user data format")
		}
	}

	// Return both parsed structures
	return &ParsedData{OuterData: outerData, UserData: userDataParsed}, nil
}
