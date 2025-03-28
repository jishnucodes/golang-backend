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
    Result []map[string]interface{}
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
    var users string

    // Execute the stored procedure and capture the result
    data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_ListUsers", nil, &users)
    if err != nil {
        return nil, fmt.Errorf("error executing stored procedure: %w", err)
    }

    // Print the result data and its type for debugging
    fmt.Println("data:", data)
    fmt.Println("data type:", fmt.Sprintf("%T", data)) 

    resultUser, ok := data.(*string)
    if !ok {
        // Log and handle the error appropriately if the type assertion fails
        log.Println("unexpected result type, expected *builder.UserDTO")
        return nil, fmt.Errorf("unexpected result type, expected *builder.UserDTO")
    }

    fmt.Println("Raw JSON:", *resultUser)

     // Parse the JSON string into a map or struct
     // Step 1: Unmarshal the outer JSON array
    var outerData []map[string]interface{}
    err = json.Unmarshal([]byte(*resultUser), &outerData)
    if err != nil {
        return nil, fmt.Errorf("failed to parse outer JSON: %w", err)
    }

    if len(outerData) == 0 {
        return &returnData{Message: "No data returned"}, nil
    }

    // Extract the nested JSON string from the "data" field
    nestedDataStr, ok := outerData[0]["data"].(string)
    if !ok {
        return nil, fmt.Errorf("missing or invalid 'data' field in JSON")
    }

    // Step 2: Unmarshal the nested JSON string into a slice of users
    var usersData []map[string]interface{}
    err = json.Unmarshal([]byte(nestedDataStr), &usersData)
    if err != nil {
        return nil, fmt.Errorf("failed to parse nested JSON data: %w", err)
    }

    // Prepare the result object
    result := &returnData{
        Message: outerData[0]["statusMessage"].(string),  // Use status message from the outer JSON
        Result:    usersData,                               // The actual user data
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

    var user string

    // Execute the stored procedure with the user data
    data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_GetAUser @UserJSON = ?", []interface{}{string(userJSON)}, &user)
    if err != nil {
        return nil, fmt.Errorf("error executing stored procedure: %w", err)
    }

    fmt.Println("data:", data)
    fmt.Println("data type:", fmt.Sprintf("%T", data))  // Print the type of data

    resultUser, ok := data.(*string)
    if !ok {
        // Log and handle the error appropriately if the type assertion fails
        log.Println("unexpected result type, expected *builder.UserDTO")
        return nil, fmt.Errorf("unexpected result type, expected *builder.UserDTO")
    }

    fmt.Println("Raw JSON:", *resultUser)

     // Parse the JSON string into a map or struct
     // Step 1: Unmarshal the outer JSON array
    var outerData []map[string]interface{}
    err = json.Unmarshal([]byte(*resultUser), &outerData)
    if err != nil {
        return nil, fmt.Errorf("failed to parse outer JSON: %w", err)
    }

    if len(outerData) == 0 {
        return &returnData{Message: "No data returned"}, nil
    }

    // Extract the nested JSON string from the "data" field
    nestedDataStr, ok := outerData[0]["data"].(string)
    if !ok {
        return nil, fmt.Errorf("missing or invalid 'data' field in JSON")
    }

    // Step 2: Unmarshal the nested JSON string into a slice of users
    var usersData []map[string]interface{}
    err = json.Unmarshal([]byte(nestedDataStr), &usersData)
    if err != nil {
        return nil, fmt.Errorf("failed to parse nested JSON data: %w", err)
    }

    // Prepare the result object
    result := &returnData{
        Message: outerData[0]["statusMessage"].(string),  // Use status message from the outer JSON
        Result:    usersData,                               // The actual user data
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
    fmt.Println("data type:", fmt.Sprintf("%T", data)) // Print the type of data

    // Handle the case when the result is a string (as in some simple responses)
    if resultUser, ok := data.(*string); ok {
        fmt.Println("Raw JSON:", *resultUser)

        // Step 1: Unmarshal the outer JSON array
        var outerData []map[string]interface{}
        err = json.Unmarshal([]byte(*resultUser), &outerData)
        if err != nil {
            return nil, fmt.Errorf("failed to parse outer JSON: %w", err)
        }

        if len(outerData) == 0 {
            return &returnData{Message: "No data returned"}, nil
        }

        // Extract the relevant fields from the response (data, status, etc.)
        dataField, ok := outerData[0]["data"].(interface{})
        if !ok {
            return nil, fmt.Errorf("missing or invalid 'data' field in JSON")
        }

        // Check if data is a string or a number and handle accordingly
        var dataValue string
        switch v := dataField.(type) {
        case string:
            dataValue = v
        case float64:
            dataValue = fmt.Sprintf("%f", v) // Convert number to string
        default:
            return nil, fmt.Errorf("unexpected data type: %T", v)
        }

        statusMessage, _ := outerData[0]["statusMessage"].(string)
        status, _ := outerData[0]["status"].(float64) // float64 because JSON numbers are unmarshalled as float64
        statusCode, _ := outerData[0]["statusCode"].(string)

        // Prepare the result object
        result := &returnData{
            Message: statusMessage, // Use status message from the outer JSON
            Result: []map[string]interface{}{  // <-- This is a slice containing a map
                {
                    "data":       dataValue,
                    "status":     status,
                    "statusCode": statusCode,
                },
            },
        }

        // Return result
        return result, nil
    }

    // Handle the case where the result is a simple string (e.g., user ID or status message)
    if resultStr, ok := data.(*string); ok {
        fmt.Println("Simple result:", *resultStr)

        // In case the result is just a status message or ID, we can return it as such
        return &returnData{
            Message: *resultStr,
        }, nil
    }

    // Handle the case when the result type is neither a string nor a structured object
    return nil, fmt.Errorf("unexpected result type, expected *string or a structured response")
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








