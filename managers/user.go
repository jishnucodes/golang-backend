package managers

import (
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	// "clinic-management/backend/models"
	"encoding/json"
	"fmt"
	"log"
)

// this is the interface which have the methods. Methods are the functions which is
// also a method of userStruct
type UserManager interface {
	Login(userData *common.UserObj) (*returnData, error)
	GetUsers() (*returnData, error)
    GetAUser(userData *common.UserObj) (*returnData, error)
	CreateUser(userData *common.UserObj) (*returnData, error)
    UpdateUser(userData *common.UserObj) (*returnData, error)
}

type returnData struct {
    Users  []builder.UserDTO  `json:"users,omitempty"`   // Slice for user records
    User    *builder.UserDTO   `json:"user,omitempty"` 
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

	var user string

    // Execute the stored procedure with the user data
    data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_UserLogin @UserJSON = ?", []interface{}{string(userJSON)}, &user)
    if err != nil {

        return nil, fmt.Errorf("error executing stored procedure: %w", err)
    }

    fmt.Println("data:", data)
    fmt.Println("data type:", fmt.Sprintf("%T", data))  // Print the type of data

    // Type assert the data returned by ExecuteStoredProcedure
    resultUser, ok := data.(*string)
    if !ok {
        // Log and handle the error appropriately if the type assertion fails
        log.Println("unexpected result type, expected *builder.UserDTO")
        return nil, fmt.Errorf("unexpected result type, expected *builder.UserDTO")
    }

    // Prepare the result object
    result := &returnData{}
    if resultUser != nil && *resultUser != ""{
        result.Message = *resultUser
    } else {
        // Handle the case where no user data is returned
        result.Message = "user login failed"
    }

    return result, nil

}

func (um *userManager) GetUsers() (*returnData, error) {
    // Create an instance of StoredProcedureExecutor
    spExecutor := common.NewStoredProcedureExecutor()

    // Define a variable to hold the users (as a slice, not a pointer)
    var users []builder.UserDTO

    // Execute the stored procedure and capture the result
    data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_ListUsers", nil, &users)
    if err != nil {
        return nil, fmt.Errorf("error executing stored procedure: %w", err)
    }

    // Print the result data and its type for debugging
    fmt.Println("data:", data)
    fmt.Println("data type:", fmt.Sprintf("%T", data))  // Should print: *([]builder.UserDTO)

    // Type assertion to *[]builder.UserDTO
    resultUser, ok := data.(*[]builder.UserDTO)
    if !ok {
        log.Println("unexpected result type, expected *[]builder.UserDTO")
        return nil, fmt.Errorf("unexpected result type, expected *[]builder.UserDTO")
    }

    // Prepare the result object to return to the caller
    result := &returnData{}

    // Check if resultUser is not nil and contains data
    if resultUser != nil && len(*resultUser) > 0 {
        // Dereference the pointer before assigning
        result.Users = *resultUser
        result.Message = "Users fetched successfully"
        result.User = nil
    } else {
        result.Message = "No users found"
    }

    return result, nil
}

func (um *userManager) GetAUser(userData *common.UserObj) (*returnData, error) {
    // Convert input to DTO
    userDTO := builder.BuildUserDTO(userData)
    fmt.Println("userDTO:", userDTO)

    // Convert DTO to JSON
    userJSON, err := json.Marshal(userDTO)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal user data: %w", err)
    }

    fmt.Println("userJSON", string(userJSON))

    // Create an instance of the stored procedure executor
    spExecutor := common.NewStoredProcedureExecutor()

    var user builder.UserDTO

    // Execute the stored procedure with the user data
    data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_GetAUser @UserJSON = ?", []interface{}{string(userJSON)}, &user)
    if err != nil {
        return nil, fmt.Errorf("error executing stored procedure: %w", err)
    }

    fmt.Println("data:", data)
    fmt.Println("data type:", fmt.Sprintf("%T", data))  // Print the type of data

    // Type assert the data returned by ExecuteStoredProcedure
    resultUser, ok := data.(*builder.UserDTO)
    if !ok {
        // Log and handle the error appropriately if the type assertion fails
        log.Println("unexpected result type, expected *builder.UserDTO")
        return nil, fmt.Errorf("unexpected result type, expected *builder.UserDTO")
    }

    

    // Prepare the result object to return to the caller
    result := &returnData{}
    if resultUser != nil {
        result.User = resultUser
        result.Message = "success"
    } else {
        // Handle the case where no user data is returned
        result.Message = "No user found with the provided details"
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

	var user string

    // Execute the stored procedure with the user data
    data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_InsertUser @UserJSON = ?", []interface{}{string(userJSON)}, &user)
    if err != nil {

        return nil, fmt.Errorf("error executing stored procedure: %w", err)
    }

    fmt.Println("data:", data)
    fmt.Println("data type:", fmt.Sprintf("%T", data))  // Print the type of data

    // Type assert the data returned by ExecuteStoredProcedure
    resultUser, ok := data.(*string)
    if !ok {
        // Log and handle the error appropriately if the type assertion fails
        log.Println("unexpected result type, expected *builder.UserDTO")
        return nil, fmt.Errorf("unexpected result type, expected *builder.UserDTO")
    }

    // Prepare the result object
    result := &returnData{}
    if resultUser != nil && *resultUser != ""{
        result.Message = *resultUser
    } else {
        // Handle the case where no user data is returned
        result.Message = "user creation failed"
    }

    return result, nil
}

func (um *userManager) UpdateUser(userData *common.UserObj) (*returnData, error) {
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

	var user string

    // Execute the stored procedure with the user data
    data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_UpdateUser @UserJSON = ?", []interface{}{string(userJSON)}, &user)
    if err != nil {

        return nil, fmt.Errorf("error executing stored procedure: %w", err)
    }
    fmt.Println("data:", data)
    fmt.Println("data type:", fmt.Sprintf("%T", data))  // Print the type of data

    // Type assert the data returned by ExecuteStoredProcedure
    resultUser, ok := data.(*string)
    if !ok {
        // Log and handle the error appropriately if the type assertion fails
        log.Println("unexpected result type, expected *builder.UserDTO")
        return nil, fmt.Errorf("unexpected result type, expected *builder.UserDTO")
    }

    // Prepare the result object
    result := &returnData{}
    if resultUser != nil && *resultUser != ""{
        result.Message = *resultUser
    } else {
        // Handle the case where no user data is returned
        result.Message = "user creation failed"
    }

    return result, nil
}








