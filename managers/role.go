package managers

import (
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"clinic-management/backend/spResponse"
	"encoding/json"
	"fmt"
)

// RoleManager interface defines the methods for role management
type RoleManager interface {
	GetRoles() (*spResponse.Result, error)
	// GetARole(roleData *requestData.RoleObj) (*spResponse.Result, error)
	CreateRole(roleData *requestData.RoleObj) (*spResponse.Result, error)
	UpdateRole(roleData *requestData.RoleObj) (*spResponse.Result, error)
	DeleteARole(roleData *requestData.RoleObj) (*spResponse.Result, error)
}

type roleManager struct {
	// Struct that implements the RoleManager interface
}

func NewRoleManager() RoleManager {
	return &roleManager{}
}

func (rm *roleManager) GetRoles() (*spResponse.Result, error) {
	// Create an instance of StoredProcedureExecutor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure and capture the result
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_ListRoles", nil)
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	// Print the result data and its type for debugging
	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

// func (rm *roleManager) GetARole(roleData *requestData.RoleObj) (*spResponse.Result, error) {
// 	// Convert input to DTO
// 	roleDTO := builder.BuildRoleDTO(roleData)
// 	fmt.Println("roleDTO:", roleDTO)

// 	// Convert DTO to JSON
// 	roleJSON, err := json.Marshal(roleDTO)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to marshal role data: %w", err)
// 	}

// 	fmt.Println("roleJSON", string(roleJSON))

// 	// Create an instance of the stored procedure executor
// 	spExecutor := common.NewStoredProcedureExecutor()

// 	// Execute the stored procedure with the role data
// 	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_GetARole @RoleJSON = ?", []interface{}{string(roleJSON)})
// 	if err != nil {
// 		return nil, fmt.Errorf("error executing stored procedure: %w", err)
// 	}

// 	fmt.Println("data:", data)
// 	fmt.Println("data type:", fmt.Sprintf("%T", data))

// 	return data, nil
// }

func (rm *roleManager) CreateRole(roleData *requestData.RoleObj) (*spResponse.Result, error) {
	// Convert input to DTO
	roleDTO := builder.BuildRoleDTO(roleData)
	fmt.Println("roleDTO:", roleDTO)

	// Convert DTO to JSON
	roleJSON, err := json.Marshal(roleDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal role data: %w", err)
	}

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the role data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_InsertRole @RoleJSON = ?", []interface{}{string(roleJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (rm *roleManager) UpdateRole(roleData *requestData.RoleObj) (*spResponse.Result, error) {
	// Convert input to DTO
	roleDTO := builder.BuildRoleDTO(roleData)
	fmt.Println("roleDTO:", roleDTO)

	// Convert DTO to JSON
	roleJSON, err := json.Marshal(roleDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal role data: %w", err)
	}

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the role data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_UpdateRole @RoleJSON = ?", []interface{}{string(roleJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (rm *roleManager) DeleteARole(roleData *requestData.RoleObj) (*spResponse.Result, error) {
	// Convert input to DTO
	roleDTO := builder.BuildRoleDTO(roleData)
	fmt.Println("roleDTO:", roleDTO)

	// Convert DTO to JSON
	roleJSON, err := json.Marshal(roleDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal role data: %w", err)
	}

	fmt.Println("roleJSON", string(roleJSON))

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the role data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_DeleteRole @RoleJSON = ?", []interface{}{string(roleJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}
