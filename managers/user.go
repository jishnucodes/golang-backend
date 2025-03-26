package managers

import (
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"clinic-management/backend/models"
	"encoding/json"
	"fmt"
)

// this is the interface which have the methods. Methods are the functions which is
// also a method of userStruct
type UserManager interface {
	Login(userData *common.UserObj) (*returnData, error)
	GetUsers() (*returnData, error)
	CreateUser(userData *common.UserObj) (*returnData, error)
}

type returnData struct {
    Users   []models.CMSUser `json:"users,omitempty"`   // Slice for user records
    Message string           `json:"message,omitempty"` // Message for string results
}

type userManager struct {
	//this is a struct which defines the methods
}

func NewUserManager() UserManager {
	return &userManager{}
}

func (um *userManager) Login(userData *common.UserObj) (*returnData, error) {

	userDTO := builder.BuildUserDTO(userData)
	fmt.Println("userDTO:", userDTO)

	//convert DTO to JSON
	userJSON, err := json.Marshal(userDTO)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal user data: %w", err)
    }

	// Create an instance of the stored procedure executor
    spExecutor := common.NewStoredProcedureExecutor()

	var user models.CMSUser

    // Execute the stored procedure with the user data
    data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_UserLogin @UserJSON = ?", []interface{}{string(userJSON)}, &user)
    if err != nil {

        return nil, fmt.Errorf("error executing stored procedure: %w", err)
    }

    // Prepare the result object
    result := &returnData{}
    if data != "" {
        result.Message = data
    } else {
        result.Message = "User created successfully"
    }

    return result, nil

}

func (um *userManager) GetUsers() (*returnData, error) {
	// Create an instance of StoredProcedureExecutor
	spExecutor := common.NewStoredProcedureExecutor()

	// Define a variable to hold the users
	var users []models.CMSUser

	// Execute the stored procedure and capture the result
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_ListUsers", nil, &users)
	if err != nil {
		
		return nil, err
	}

	// Prepare the result
	result := &returnData{}

	if data != "" {
		result.Message = data
	} else if len(users) > 0 {
		result.Users = users
	} else {
		result.Message = "No users found."
	}

	return result, nil
}

func (um *userManager) CreateUser(userData *common.UserObj) (*returnData, error) {
    // Convert input to DTO
    userDTO := builder.BuildUserDTO(userData)
    fmt.Println("userDTO:", userDTO)

    // Convert DTO to JSON
    userJSON, err := json.Marshal(userDTO)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal user data: %w", err)
    }

    // Create an instance of the stored procedure executor
    spExecutor := common.NewStoredProcedureExecutor()

	var user models.CMSUser

    // Execute the stored procedure with the user data
    data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_InsertUser @UserJSON = ?", []interface{}{string(userJSON)}, &user)
    if err != nil {

        return nil, fmt.Errorf("error executing stored procedure: %w", err)
    }

    // Prepare the result object
    result := &returnData{}
    if data != "" {
        result.Message = data
    } else {
        result.Message = "User created successfully"
    }

    return result, nil
}






