package managers

import (
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"clinic-management/backend/spResponse"
	"encoding/json"
	"fmt"
)

// DepartmentManager interface defines the methods for department management
type DepartmentManager interface {
	GetDepartments() (*spResponse.Result, error)
	GetADepartment(departmentData *requestData.DepartmentObj) (*spResponse.Result, error)
	CreateDepartment(departmentData *requestData.DepartmentObj) (*spResponse.Result, error)
	UpdateDepartment(departmentData *requestData.DepartmentObj) (*spResponse.Result, error)
	DeleteDepartment(departmentData *requestData.DepartmentObj) (*spResponse.Result, error)
}

type departmentManager struct {
	// Struct that implements the DepartmentManager interface
}

func NewDepartmentManager() DepartmentManager {
	return &departmentManager{}
}

func (dm *departmentManager) GetDepartments() (*spResponse.Result, error) {
	// Create an instance of StoredProcedureExecutor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure and capture the result
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_ListDepartments", nil)
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	// Print the result data and its type for debugging
	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (dm *departmentManager) GetADepartment(departmentData *requestData.DepartmentObj) (*spResponse.Result, error) {
	// Convert input to DTO
	departmentDTO := builder.BuildDepartmentObj(departmentData)
	fmt.Println("departmentDTO:", departmentDTO)

	// Convert DTO to JSON
	departmentJSON, err := json.Marshal(departmentDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal department data: %w", err)
	}

	fmt.Println("departmentJSON", string(departmentJSON))

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the department data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_GetADepartment @DepartmentJSON = ?", []interface{}{string(departmentJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (dm *departmentManager) CreateDepartment(departmentData *requestData.DepartmentObj) (*spResponse.Result, error) {
	// Convert input to DTO
	departmentDTO := builder.BuildDepartmentObj(departmentData)
	fmt.Println("departmentDTO:", departmentDTO)

	// Convert DTO to JSON
	departmentJSON, err := json.Marshal(departmentDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal department data: %w", err)
	}

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the department data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_InsertDepartment @DepartmentJSON = ?", []interface{}{string(departmentJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (dm *departmentManager) UpdateDepartment(departmentData *requestData.DepartmentObj) (*spResponse.Result, error) {
	// Convert input to DTO
	departmentDTO := builder.BuildDepartmentObj(departmentData)
	fmt.Println("departmentDTO:", departmentDTO)

	// Convert DTO to JSON
	departmentJSON, err := json.Marshal(departmentDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal department data: %w", err)
	}

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the department data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_UpdateDepartment @DepartmentJSON = ?", []interface{}{string(departmentJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (dm *departmentManager) DeleteDepartment(departmentData *requestData.DepartmentObj) (*spResponse.Result, error) {
	// Convert input to DTO
	departmentDTO := builder.BuildDepartmentObj(departmentData)
	fmt.Println("departmentDTO:", departmentDTO)

	// Convert DTO to JSON
	departmentJSON, err := json.Marshal(departmentDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal department data: %w", err)
	}

	fmt.Println("departmentJSON", string(departmentJSON))

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the department data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_DeleteDepartment @DepartmentJSON = ?", []interface{}{string(departmentJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}
