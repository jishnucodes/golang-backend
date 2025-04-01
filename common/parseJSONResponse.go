package common

import (
	"encoding/json"
	"fmt"
)

type ParsedData struct {
	OuterData []map[string]interface{}
	UserData  []map[string]interface{}
}

// Function to parse outer and nested JSON data
func ParseJSONResponse(jsonStr *string) (*ParsedData, error) {
	// Step 1: Unmarshal outer JSON
	var outerData []map[string]interface{}
	err := json.Unmarshal([]byte(*jsonStr), &outerData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse outer JSON: %w", err)
	}

	if len(outerData) == 0 {
		return &ParsedData{OuterData: nil, UserData: nil}, nil
	}

	// Step 2: Extract the nested "data" field
	nestedDataStr, ok := outerData[0]["data"].(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'data' field in JSON")
	}

	// Step 3: Unmarshal the nested JSON string
	var userData []map[string]interface{}
	err = json.Unmarshal([]byte(nestedDataStr), &userData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse nested JSON data: %w", err)
	}

	// Return both parsed structures
	return &ParsedData{OuterData: outerData, UserData: userData}, nil
}