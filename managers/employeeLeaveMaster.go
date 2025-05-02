package managers

import (
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"clinic-management/backend/spResponse"
	"encoding/json"
	"fmt"
)

// EmployeeManager interface defines the methods for employee management
type EmployeeLeaveManager interface {
	GetEmployeeLeaves() (*spResponse.Result, error)
	GetAEmployeeLeave(employeeLeaveData *requestData.EmployeeLeaveObj) (*spResponse.Result, error)
	CreateEmployeeLeave(employeeLeaveData *requestData.EmployeeLeaveObj) (*spResponse.Result, error)
	UpdateEmployeeLeave(employeeLeaveData *requestData.EmployeeLeaveObj) (*spResponse.Result, error)
	DeleteEmployeeLeave(employeeLeaveData *requestData.EmployeeLeaveObj) (*spResponse.Result, error)
}

type employeeLeaveManager struct {
	// Struct that implements the EmployeeManager interface
}

func NewEmployeeLeaveManager() EmployeeLeaveManager {
	return &employeeLeaveManager{}
}

func (em *employeeLeaveManager) GetEmployeeLeaves() (*spResponse.Result, error) {
	// Create an instance of StoredProcedureExecutor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure and capture the result
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_ListEmployeeLeaves", nil)
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	// Print the result data and its type for debugging
	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (em *employeeLeaveManager) GetAEmployeeLeave(employeeLeaveData *requestData.EmployeeLeaveObj) (*spResponse.Result, error) {
	// Convert input to DTO
	employeeLeaveDTO := builder.BuildEmployeeLeaveDTO(employeeLeaveData)
	fmt.Println("employeeLeaveDTO:", employeeLeaveDTO)

	// Convert DTO to JSON
	employeeLeaveJSON, err := json.Marshal(employeeLeaveDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal employee data: %w", err)
	}

	fmt.Println("employeeLeaveJSON", string(employeeLeaveJSON))

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the employee data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_GetAEmployeeLeave @EmployeeLeaveJSON = ?", []interface{}{string(employeeLeaveJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (em *employeeLeaveManager) CreateEmployeeLeave(employeeLeaveData *requestData.EmployeeLeaveObj) (*spResponse.Result, error) {
	// Convert input to DTO
	employeeLeaveDTO := builder.BuildEmployeeLeaveDTO(employeeLeaveData)
	fmt.Println("employeeLeaveDTO:", employeeLeaveDTO)

	// Convert DTO to JSON
	employeeLeaveJSON, err := json.Marshal(employeeLeaveDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal employee data: %w", err)
	}

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the employee data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_InsertEmployeeLeave @EmployeeLeaveJSON = ?", []interface{}{string(employeeLeaveJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (em *employeeLeaveManager) UpdateEmployeeLeave(employeeLeaveData *requestData.EmployeeLeaveObj) (*spResponse.Result, error) {
	// Convert input to DTO
	employeeLeaveDTO := builder.BuildEmployeeLeaveDTO(employeeLeaveData)
	fmt.Println("employeeLeaveDTO:", employeeLeaveDTO)

	// Convert DTO to JSON
	employeeLeaveJSON, err := json.Marshal(employeeLeaveDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal employee data: %w", err)
	}

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the employee data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_UpdateEmployeeLeave @EmployeeLeaveJSON = ?", []interface{}{string(employeeLeaveJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (em *employeeLeaveManager) DeleteEmployeeLeave(employeeLeaveData *requestData.EmployeeLeaveObj) (*spResponse.Result, error) {
	// Convert input to DTO
	employeeLeaveDTO := builder.BuildEmployeeLeaveDTO(employeeLeaveData)
	fmt.Println("employeeLeaveDTO:", employeeLeaveDTO)

	// Convert DTO to JSON
	employeeLeaveJSON, err := json.Marshal(employeeLeaveDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal employee data: %w", err)
	}

	fmt.Println("employeeLeaveJSON", string(employeeLeaveJSON))

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the employee data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_DeleteEmployeeLeave @EmployeeLeaveJSON = ?", []interface{}{string(employeeLeaveJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}
