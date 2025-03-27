package common

import (
	"clinic-management/backend/database"
	"fmt"
	"log"
)

// StoredProcedureExecutor defines the interface for executing stored procedures
type StoredProcedureExecutor interface {
	ExecuteStoredProcedure(spName string, params []interface{}, result interface{}) (interface{}, error)
}

// storedProcedureExecutor is the concrete implementation of the StoredProcedureExecutor interface
type storedProcedureExecutor struct {
}

// NewStoredProcedureExecutor returns a new instance of the storedProcedureExecutor
func NewStoredProcedureExecutor() *storedProcedureExecutor {
	return &storedProcedureExecutor{}
}



// ExecuteStoredProcedure executes a stored procedure and returns the result or error
// func (sp *storedProcedureExecutor) ExecuteStoredProcedure(spName string, params []interface{}, result interface{}) (string, error) {
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
// 		// if err := database.DB.ScanRows(rows, result); err != nil {
// 		// 	// Try scanning into a string message if not scanning into struct
// 		// 	if scanErr := rows.Scan(&message); scanErr == nil {
// 		// 		isString = true
// 		// 		continue
// 		// 	}
// 		// 	log.Println("error scanning row: %w", err)
// 		// 	return "", fmt.Errorf("error scanning row: %w", err)
// 		// }
// 		if err := query.Find(result).Error; err != nil {
// 			log.Printf("Error executing stored procedure %s: %v\n", spName, err)
// 			return "", fmt.Errorf("error executing stored procedure %s: %w", spName, err)
// 		}
// 	}

// 	// Return the result as message or empty string
// 	if isString {
// 		return message, nil
// 	}
// 	return "", nil
// }


func (sp *storedProcedureExecutor) ExecuteStoredProcedure(spName string, params []interface{}, result interface{}) (interface{}, error) {
    // Use GORM Raw to execute the stored procedure and map the result
    query := database.DB.Raw(spName, params...)
    
    // Execute the query and map the result to the 'result' variable
    if err := query.Find(result).Error; err != nil {
        log.Printf("Error executing stored procedure %s: %v\n", spName, err)
        return nil, fmt.Errorf("error executing stored procedure %s: %w", spName, err)
    }

    // Return the result data
    return result, nil
}



