package managers

import (
	// "clinic-management/backend/builder"
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"clinic-management/backend/spResponse"
	"encoding/json"

	// "clinic-management/backend/models"
	// "encoding/json"
	"fmt"
	// "log"
)

// this is the interface which have the methods. Methods are the functions which is
// also a method of userStruct
type UserManager interface {
	Login(userData *common.UserObj) (*spResponse.Result, error)
	GetUsers() (*spResponse.Result, error)
    GetAUser(userData *common.UserObj) (*spResponse.Result, error)
	CreateUser(userData *common.UserObj) (*spResponse.Result, error)
    UpdateUser(userData *common.UserObj) (*spResponse.Result, error)
}

// type returnData struct {
//     Users  []builder.UserDTO  `json:"users,omitempty"`   // Slice for user records
//     Result []map[string]interface{}
//     User    *builder.UserDTO   `json:"user,omitempty"` 
//     Message string           `json:"message,omitempty"` // Message for string results
//     Status int               `json:"status,omitempty"`
// }

type Result struct {
	Data string  `json:"data,omitempty"`
	Status int 					    `json:"status,omitempty"`
	StatusCode string				`json:"statusCode,omitempty"`
	StatusMessage string			`json:"statusMessage,omitempty"`

}



type userManager struct {
	//this is a struct which defines the methods
}

func NewUserManager() UserManager {
	return &userManager{}
}

func (um *userManager) Login(userData *common.UserObj) (*spResponse.Result, error) {

	userDTO := builder.BuildUserDTO(userData)
	fmt.Println("userDTO:", userDTO)

	//convert DTO to JSON
	userJSON, err := json.Marshal(userDTO)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal user data: %w", err)
    }

	// Create an instance of the stored procedure executor
    spExecutor := common.NewStoredProcedureExecutor()


    // Execute the stored procedure with the user data
    data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_UserLogin @UserJSON = ?", []interface{}{string(userJSON)})
    if err != nil {

        return nil, fmt.Errorf("error executing stored procedure: %w", err)
    }

    fmt.Println("data:", data)
    fmt.Println("data type:", fmt.Sprintf("%T", data))  // Print the type of data

    return data, nil

}

func (um *userManager) GetUsers() (*spResponse.Result, error) {
    // Create an instance of StoredProcedureExecutor
    spExecutor := common.NewStoredProcedureExecutor()

    // Execute the stored procedure and capture the result
    data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_ListUsers", nil)
    if err != nil {
        return nil, fmt.Errorf("error executing stored procedure: %w", err)
    }

    // Print the result data and its type for debugging
    fmt.Println("data:", data)
    fmt.Println("data type:", fmt.Sprintf("%T", data)) 

    return data, nil
}

func (um *userManager) GetAUser(userData *common.UserObj) (*spResponse.Result, error) {
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

    // var user result.Result

    // Execute the stored procedure with the user data
    data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_GetAUser @UserJSON = ?", []interface{}{string(userJSON)})
    if err != nil {
        return nil, fmt.Errorf("error executing stored procedure: %w", err)
    }

    fmt.Println("data:", data)
    fmt.Println("data type:", fmt.Sprintf("%T", data))  // Print the type of data

    return data, nil
}

func (um *userManager) CreateUser(userData *common.UserObj) (*spResponse.Result, error) {
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


    // Execute the stored procedure with the user data
    data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_InsertUser @UserJSON = ?", []interface{}{string(userJSON)})
    if err != nil {
        return nil, fmt.Errorf("error executing stored procedure: %w", err)
    }

    fmt.Println("data:", data)
    fmt.Println("data type:", fmt.Sprintf("%T", data)) // Print the type of data

    return data, nil
}


func (um *userManager) UpdateUser(userData *common.UserObj) (*spResponse.Result, error) {
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

    // Execute the stored procedure with the user data
    data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_UpdateUser @UserJSON = ?", []interface{}{string(userJSON)})
    if err != nil {

        return nil, fmt.Errorf("error executing stored procedure: %w", err)
    }
    fmt.Println("data:", data)
    fmt.Println("data type:", fmt.Sprintf("%T", data))  // Print the type of data

    return data, nil
}








