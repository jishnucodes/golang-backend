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
type EmployeeManager interface {
	GetEmployees() (*spResponse.Result, error)
	GetAEmployee(employeeData *requestData.EmployeeObj) (*spResponse.Result, error)
	CreateEmployee(employeeData *requestData.EmployeeObj) (*spResponse.Result, error)
	UpdateEmployee(employeeData *requestData.EmployeeObj) (*spResponse.Result, error)
	DeleteEmployee(employeeData *requestData.EmployeeObj) (*spResponse.Result, error)
}

type employeeManager struct {
	// Struct that implements the EmployeeManager interface
}

func NewEmployeeManager() EmployeeManager {
	return &employeeManager{}
}

func (em *employeeManager) GetEmployees() (*spResponse.Result, error) {
	// Create an instance of StoredProcedureExecutor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure and capture the result
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_ListEmployees", nil)
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	// Print the result data and its type for debugging
	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (em *employeeManager) GetAEmployee(employeeData *requestData.EmployeeObj) (*spResponse.Result, error) {
	// Convert input to DTO
	employeeDTO := builder.BuildEmployeeDTO(employeeData)
	fmt.Println("employeeDTO:", employeeDTO)

	// Convert DTO to JSON
	employeeJSON, err := json.Marshal(employeeDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal employee data: %w", err)
	}

	fmt.Println("employeeJSON", string(employeeJSON))

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the employee data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_GetAEmployee @EmployeeJSON = ?", []interface{}{string(employeeJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (em *employeeManager) CreateEmployee(employeeData *requestData.EmployeeObj) (*spResponse.Result, error) {
	// Convert input to DTO
	employeeDTO := builder.BuildEmployeeDTO(employeeData)
	fmt.Println("employeeDTO:", employeeDTO)

	// Convert DTO to JSON
	employeeJSON, err := json.Marshal(employeeDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal employee data: %w", err)
	}

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the employee data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_InsertEmployee @EmployeeJSON = ?", []interface{}{string(employeeJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (em *employeeManager) UpdateEmployee(employeeData *requestData.EmployeeObj) (*spResponse.Result, error) {
	// Convert input to DTO
	employeeDTO := builder.BuildEmployeeDTO(employeeData)
	fmt.Println("employeeDTO:", employeeDTO)

	// Convert DTO to JSON
	employeeJSON, err := json.Marshal(employeeDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal employee data: %w", err)
	}

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the employee data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_UpdateEmployee @EmployeeJSON = ?", []interface{}{string(employeeJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (em *employeeManager) DeleteEmployee(employeeData *requestData.EmployeeObj) (*spResponse.Result, error) {
	// Convert input to DTO
	employeeDTO := builder.BuildEmployeeDTO(employeeData)
	fmt.Println("employeeDTO:", employeeDTO)

	// Convert DTO to JSON
	employeeJSON, err := json.Marshal(employeeDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal employee data: %w", err)
	}

	fmt.Println("employeeJSON", string(employeeJSON))

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the employee data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_DeleteEmployee @EmployeeJSON = ?", []interface{}{string(employeeJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}
