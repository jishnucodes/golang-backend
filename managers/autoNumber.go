package managers

import (
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"clinic-management/backend/spResponse"
	"encoding/json"
	"fmt"
)

type AutoNumberManager interface {
	GenerateAutoNumber(autoNumberData *requestData.AutoNumberObj) (*spResponse.Result, error)
}

type autoNumberManager struct {
	
}


func NewAutoNumberManager() AutoNumberManager {	
	return &autoNumberManager{}
}

func (manager *autoNumberManager) GenerateAutoNumber(autoNumberData *requestData.AutoNumberObj) (*spResponse.Result, error) {

	// Convert input to DTO
	autoNumberDTO := builder.BuildAutoNumberDTO(autoNumberData)
	fmt.Println("autoNumberDTO:", autoNumberDTO)

	// Convert DTO to JSON
	autoNumberJSON, err := json.Marshal(autoNumberDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal autoNumber data: %w", err)
	}

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()
	
	// Execute the stored procedure with the autoNumber data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_GenerateAutoNumberCode @AutoNumberJSON = ?", []interface{}{string(autoNumberJSON)})
	if err != nil {
		return nil, fmt.Errorf("failed to execute stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
	
}

