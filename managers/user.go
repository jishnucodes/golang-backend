package managers

import (
	// "clinic-management/backend/builder"
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"clinic-management/backend/result"
	"encoding/json"

	// "clinic-management/backend/models"
	// "encoding/json"
	"fmt"
	// "log"
)

// this is the interface which have the methods. Methods are the functions which is
// also a method of userStruct
type UserManager interface {
	// Login(userData *common.UserObj) (*returnData, error)
	GetUsers() (*result.Result, error)
    GetAUser(userData *common.UserObj) (*result.Result, error)
	// CreateUser(userData *common.UserObj) (*returnData, error)
    // UpdateUser(userData *common.UserObj) (*returnData, error)
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

// func (um *userManager) Login(userData *common.UserObj) (*returnData, error) {

// 	userDTO := builder.BuildUserDTO(userData)
// 	fmt.Println("userDTO:", userDTO)

// 	//convert DTO to JSON
// 	userJSON, err := json.Marshal(userDTO)
//     if err != nil {
//         return nil, fmt.Errorf("failed to marshal user data: %w", err)
//     }

// 	// Create an instance of the stored procedure executor
//     spExecutor := common.NewStoredProcedureExecutor()

// 	var user string

//     // Execute the stored procedure with the user data
//     data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_UserLogin @UserJSON = ?", []interface{}{string(userJSON)}, &user)
//     if err != nil {

//         return nil, fmt.Errorf("error executing stored procedure: %w", err)
//     }

//     fmt.Println("data:", data)
//     fmt.Println("data type:", fmt.Sprintf("%T", data))  // Print the type of data

//     // Type assert the data returned by ExecuteStoredProcedure
//     resultUser, ok := data.(*string)
//     if !ok {
//         // Log and handle the error appropriately if the type assertion fails
//         log.Println("unexpected result type, expected *builder.UserDTO")
//         return nil, fmt.Errorf("unexpected result type, expected *builder.UserDTO")
//     }

//     // Prepare the result object
//     fmt.Println("Raw JSON:", *resultUser)

//     parsedData, err := common.ParseJSONResponse(resultUser)
//     if err != nil {
// 		fmt.Println("Error:", err)
// 		return nil, fmt.Errorf("failed to parse outer JSON: %w", err)
// 	}

//     // Prepare the result object
//     result := &returnData{
//         Message: parsedData.OuterData[0]["statusMessage"].(string),  // Use status message from the outer JSON
//         Result:    parsedData.UserData,                               // The actual user data
//     }

//     return result, nil

// }

func (um *userManager) GetUsers() (*result.Result, error) {
    // Create an instance of StoredProcedureExecutor
    spExecutor := common.NewStoredProcedureExecutor()


    

    // Define a variable to hold the users (as a slice, not a pointer)
    // var users result.Result

    // Execute the stored procedure and capture the result
    data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_ListUsers", nil)
    if err != nil {
        return nil, fmt.Errorf("error executing stored procedure: %w", err)
    }

    // Print the result data and its type for debugging
    fmt.Println("data:", data)
    fmt.Println("data type:", fmt.Sprintf("%T", data)) 

    // resultUser, ok := data.(*result.Result)
    // if !ok {
    //     // Log and handle the error appropriately if the type assertion fails
    //     log.Println("unexpected result type, expected *builder.UserDTO")
    //     return nil, fmt.Errorf("unexpected result type, expected *builder.UserDTO")
    // }

    // fmt.Println("Raw JSON:", *resultUser)

    // parsedData, err := common.ParseJSONResponse(resultUser.Data)
    // if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return nil, fmt.Errorf("failed to parse outer JSON: %w", err)
	// }

    // // Prepare the result object
    // result := &Result{
    //     StatusMessage: parsedData.OuterData[0]["statusMessage"].(string),  // Use status message from the outer JSON
    //     Response:    parsedData.UserData,                               // The actual user data
    // }

    return data, nil
}

func (um *userManager) GetAUser(userData *common.UserObj) (*result.Result, error) {
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

    // resultUser, ok := data.(*result.Result)
    // if !ok {
    //     // Log and handle the error appropriately if the type assertion fails
    //     log.Println("unexpected result type, expected *builder.UserDTO")
    //     return nil, fmt.Errorf("unexpected result type, expected *builder.UserDTO")
    // }

    // fmt.Println("Raw JSON:", *resultUser)

    // parsedData, err := common.ParseJSONResponse(resultUser)
    // if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return nil, fmt.Errorf("failed to parse outer JSON: %w", err)
	// }

    // // Prepare the result object
    // result := &returnData{
    //     Message: parsedData.OuterData[0]["statusMessage"].(string),  // Use status message from the outer JSON
    //     Result:    parsedData.UserData,                               // The actual user data
    // }

    return data, nil
}

// func (um *userManager) CreateUser(userData *common.UserObj) (*returnData, error) {
//     // Convert input to DTO
//     userDTO := builder.BuildUserDTO(userData)
//     fmt.Println("userDTO:", userDTO)

//     // Convert DTO to JSON
//     userJSON, err := json.Marshal(userDTO)
//     if err != nil {
//         return nil, fmt.Errorf("failed to marshal user data: %w", err)
//     }

//     // Create an instance of the stored procedure executor
//     spExecutor := common.NewStoredProcedureExecutor()

//     var user string //need to change the struct as like the data

//     // Execute the stored procedure with the user data
//     data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_InsertUser @UserJSON = ?", []interface{}{string(userJSON)}, &user)
//     if err != nil {
//         return nil, fmt.Errorf("error executing stored procedure: %w", err)
//     }

//     fmt.Println("data:", data)
//     fmt.Println("data type:", fmt.Sprintf("%T", data)) // Print the type of data

    

//     resultUser, ok := data.(*string)
//     if !ok {
//         // Log and handle the error appropriately if the type assertion fails
//         log.Println("unexpected result type, expected *builder.UserDTO")
//         return nil, fmt.Errorf("unexpected result type, expected *builder.UserDTO")
//     }

//     fmt.Println("Raw JSON:", *resultUser)

//     parsedData, err := common.ParseJSONResponse(resultUser)
//     if err != nil {
// 		fmt.Println("Error:", err)
// 		return nil, fmt.Errorf("failed to parse outer JSON: %w", err)
// 	}

//     // Prepare the result object
//     result := &returnData{
//         Message: parsedData.OuterData[0]["statusMessage"].(string),  // Use status message from the outer JSON
//         Result:    parsedData.UserData,                               // The actual user data
//     }

//     return result, nil
// }


// func (um *userManager) UpdateUser(userData *common.UserObj) (*returnData, error) {
//     // Convert input to DTO
//     userDTO := builder.BuildUserDTO(userData)
//     fmt.Println("userDTO:", userDTO)

//     // Convert DTO to JSON
//     userJSON, err := json.Marshal(userDTO)
//     if err != nil {
//         return nil, fmt.Errorf("failed to marshal user data: %w", err)
//     }

//     // Create an instance of the stored procedure executor
//     spExecutor := common.NewStoredProcedureExecutor()

// 	var user string

//     // Execute the stored procedure with the user data
//     data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_UpdateUser @UserJSON = ?", []interface{}{string(userJSON)}, &user)
//     if err != nil {

//         return nil, fmt.Errorf("error executing stored procedure: %w", err)
//     }
//     fmt.Println("data:", data)
//     fmt.Println("data type:", fmt.Sprintf("%T", data))  // Print the type of data

//     // Type assert the data returned by ExecuteStoredProcedure
//     resultUser, ok := data.(*string)
//     if !ok {
//         // Log and handle the error appropriately if the type assertion fails
//         log.Println("unexpected result type, expected *builder.UserDTO")
//         return nil, fmt.Errorf("unexpected result type, expected *builder.UserDTO")
//     }

//     parsedData, err := common.ParseJSONResponse(resultUser)
//     if err != nil {
// 		fmt.Println("Error:", err)
// 		return nil, fmt.Errorf("failed to parse outer JSON: %w", err)
// 	}

//     // Prepare the result object
//     result := &returnData{
//         Message: parsedData.OuterData[0]["statusMessage"].(string),  // Use status message from the outer JSON
//         Result:    parsedData.UserData,                               // The actual user data
//     }

//     return result, nil
// }








