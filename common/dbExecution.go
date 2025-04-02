package common

import (
	"clinic-management/backend/database"
	"clinic-management/backend/result"
	"fmt"
	"log"
)

// StoredProcedureExecutor defines the interface for executing stored procedures
type StoredProcedureExecutor interface {
	ExecuteStoredProcedure(spName string, params []interface{}, result interface{}) (interface{}, error)
    // ExecuteStoredProcedure(spName string, params []interface{}, result interface{}) ([]map[string]interface{}, error)
}

// storedProcedureExecutor is the concrete implementation of the StoredProcedureExecutor interface
type storedProcedureExecutor struct {
}

// NewStoredProcedureExecutor returns a new instance of the storedProcedureExecutor
func NewStoredProcedureExecutor() *storedProcedureExecutor {
	return &storedProcedureExecutor{}
}



// ExecuteStoredProcedure executes a stored procedure and returns the result or error
// func (sp *storedProcedureExecutor) ExecuteStoredProcedure(spName string, params []interface{}, result interface{}) (interface{}, error) {
// 	// Execute the stored procedure
// 	rows, err := database.DB.Raw(spName, params...).Rows()
// 	if err != nil {
// 		log.Printf("Error executing stored procedure %s: %v\n", spName, err)
// 		return "", fmt.Errorf("error executing stored procedure %s: %w", spName, err)
// 	}
// 	defer rows.Close()

// 	fmt.Println("rows: ", rows)

// 	var message string
// 	isString := false

// 	// Process the rows returned by the stored procedure
// 	for rows.Next() {
// 		// Scan into the result depending on the expected type (users or message)
// 		if err := database.DB.ScanRows(rows, result); err != nil {
// 			// Try scanning into a string message if not scanning into struct
// 			if scanErr := rows.Scan(&message); scanErr == nil {
// 				isString = true
// 				continue
// 			}
// 			log.Println("error scanning row: %w", err)
// 			return "", fmt.Errorf("error scanning row: %w", err)
// 		}
// 		// if err := query.Find(result).Error; err != nil {
// 		// 	log.Printf("Error executing stored procedure %s: %v\n", spName, err)
// 		// 	return "", fmt.Errorf("error executing stored procedure %s: %w", spName, err)
// 		// }
// 	}

// 	// Return the result as message or empty string
// 	if isString {
// 		return message, nil
// 	}
// 	return "", nil
// }




// func (sp *storedProcedureExecutor) ExecuteStoredProcedure(spName string, params []interface{}, result *result.Result) (interface{}, error) {
//     // Use GORM Raw to execute the stored procedure and map the result
//     query := database.DB.Raw(spName, params...)

//     fmt.Println("query", query)
    
//     // Execute the query and map the result to the 'result' variable
//     if err := query.Find(&result).Error; err != nil {
//         log.Printf("Error executing stored procedure %s: %v\n", spName, err)
//         return nil, fmt.Errorf("error executing stored procedure %s: %w", spName, err)
//     }

//     // Return the result data
//     return result, nil
// }

func (sp *storedProcedureExecutor) ExecuteStoredProcedure(spName string, params []interface{}) (*result.Result, error) {
    var response result.Result

    rows, err := database.DB.Raw(spName, params...).Rows()
    if err != nil {
        log.Printf("Error executing stored procedure %s: %v\n", spName, err)
        return nil, fmt.Errorf("error executing stored procedure %s: %w", spName, err)
    }
    defer rows.Close()

    // Manually scan values
    for rows.Next() {
        if err := rows.Scan(&response.Data, &response.Status, &response.StatusCode, &response.StatusMessage); err != nil {
            log.Println("Error scanning row:", err)
            return nil, err
        }
    }

    return &response, nil
}




// func (sp *storedProcedureExecutor) ExecuteStoredProcedure(spName string, params []interface{}, result interface{}) ([]map[string]interface{}, error) {
// 	// Step 1: Execute the stored procedure
// 	query := database.DB.Raw(spName, params...)

// 	// Execute the query and map the result to the 'result' variable
// 	if err := query.Find(result).Error; err != nil {
// 		log.Printf("Error executing stored procedure %s: %v\n", spName, err)
// 		return nil, fmt.Errorf("error executing stored procedure %s: %w", spName, err)
// 	}

// 	// Step 2: Convert result to JSON
// 	data, err := json.Marshal(result)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to marshal result to JSON: %w", err)
// 	}

// 	// Step 3: Unmarshal JSON into []map[string]interface{}
// 	var outerData []map[string]interface{}
// 	err = json.Unmarshal(data, &outerData)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to parse JSON: %w", err)
// 	}

// 	// Step 4: Extract nested JSON data if "data" field exists
// 	var usersData []map[string]interface{}

// 	if len(outerData) > 0 {
// 		if nestedDataStr, ok := outerData[0]["data"].(string); ok {
// 			// Unmarshal nested data
// 			err = json.Unmarshal([]byte(nestedDataStr), &usersData)
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to parse nested JSON data: %w", err)
// 			}
// 		} else {
// 			// No nested data, return the outer data
// 			usersData = outerData
// 		}
// 	}

//     fmt.Println(usersData)

// 	// Return the parsed data
// 	return usersData, nil
// }



