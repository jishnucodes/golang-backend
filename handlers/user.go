package handlers

import (
	// "clinic-management/backend/builder"
	// "clinic-management/backend/builder"
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"encoding/json"
	"strconv"

	// "clinic-management/backend/errorHandlers"
	"clinic-management/backend/managers"
	"fmt"
	"log"
	"net/http"


	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	groupName   string
	userManager managers.UserManager
}

func NewUserHandler(userManager managers.UserManager) *UserHandler {
	return &UserHandler{
		"api/user",
		userManager,
	}
}

func (handler *UserHandler) RegisterApis(r *gin.Engine) {
	userGroup := r.Group(handler.groupName)
	userGroup.GET("/list", handler.UserList)
	userGroup.GET("/:userId", handler.GetAUser)
	// userGroup.POST("/create", handler.InsertUser)
	// userGroup.POST("/login", handler.Login)
	// userGroup.PUT("/update/:userId", handler.UpdateUser)
}




// func (handler *UserHandler) Login(ctx *gin.Context) {
//     userData := common.NewUserOb()

//     // Bind the incoming JSON to the userData object
//     err := ctx.BindJSON(&userData)
//     if err != nil {
//         log.Println("binding user details from json is failed:", err)
//         common.SendError(ctx, http.StatusBadRequest, "binding user details from json is failed", err)
//         return
//     }

//     // Call the Login method and get the result
//     result, err := handler.userManager.Login(userData)

// 	fmt.Println("Result", result)

//     // Debug prints to check result types and values
//     // fmt.Printf("Type of result.Result: %T\n", result.Result)
//     // fmt.Printf("Value of result.Result: %#v\n", result.Result)

// 	//this is mainly used to handle the sql custom error messages
// 	errors := errorHandlers.HandleErrorResponse(ctx, result.Result, err)
//     if errors != nil {
//         return 
//     }

//     // If no error message exists, continue with building the response
//     response := builder.BuildUserDTOs(result.Result)
//     fmt.Println("response", result.Result)
//     fmt.Println("error", err)

//     // If there was an error from the login function
//     if err != nil {
//         log.Println("user login failed", err)
//         common.SendError(ctx, http.StatusInternalServerError, result.Message, err)
//         return
//     }

//     // If everything is fine, send success response
//     common.SendSuccess(ctx, http.StatusOK, result.Message, response)
//     log.Println("user logged in successfully")
// }


func (handler *UserHandler) UserList(ctx *gin.Context) {

	users, err := handler.userManager.GetUsers()

	if err != nil {
		log.Println("user fetching failed", err)
		common.SendError(ctx, http.StatusInternalServerError, 0, users.StatusMessage, err)
		return
	}

	fmt.Printf("users.Data type: %T\n", users.Data)
	fmt.Println("users.Data content:", users.Data)

	// users.Data is a string, so directly unmarshal
	var parsedData []map[string]interface{}
	unmarshalError := json.Unmarshal([]byte(users.Data), &parsedData)
	if unmarshalError != nil {
		log.Println("Failed to parse user data:", unmarshalError)
		common.SendError(ctx, http.StatusInternalServerError, 0, "Failed to parse user data", unmarshalError)
		return
	}

	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildUserDTOs(parsedData)

	fmt.Println("response", response)

	common.SendSuccess(ctx, http.StatusOK, users.Status, users.StatusMessage, response)

	log.Println("users fetched successfully")

}

func (handler *UserHandler) GetAUser(ctx *gin.Context) {
	// Create a new user object
	userData := common.NewUserOb()

	// Bind the request body to the userData object
	// err := ctx.BindJSON(&userData)
	// if err != nil {
	//     log.Println("binding user details from JSON failed:", err)
	//     common.SendError(ctx, http.StatusBadRequest, err)
	//     return
	// }

	// Extract the user ID from the URL parameters
	userIdStr, ok := ctx.Params.Get("userId")
	if !ok {
		log.Println("user ID is missing in the request")
		common.SendError(ctx, http.StatusBadRequest, 0,"", fmt.Errorf("user ID is required"))
		return
	}

	// Convert string userId to uint
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		log.Println("invalid user ID format:", err)
		common.SendError(ctx, http.StatusBadRequest,0, "", fmt.Errorf("invalid user ID format"))
		return
	}

	// Assign the user ID to the userData object
	userData.UserID = uint(userId)

	// Call the update method in the user manager
	user, err := handler.userManager.GetAUser(userData)

	if err != nil {
		log.Println("user update failed:", err)
		common.SendError(ctx, http.StatusInternalServerError,0, user.StatusMessage, err)
		return
	}

	fmt.Printf("user.Data type: %T\n", user.Data)
	fmt.Println("user.Data content:", user.Data)

	// users.Data is a string, so directly unmarshal
	var parsedData []map[string]interface{}
	unmarshalError := json.Unmarshal([]byte(user.Data), &parsedData)
	if unmarshalError != nil {
		log.Println("Failed to parse user data:", unmarshalError)
		common.SendError(ctx, http.StatusInternalServerError, 0, "Failed to parse user data", unmarshalError)
		return
	}

	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildUserDTOs(parsedData)

	// errors := errorHandlers.HandleErrorResponse(ctx, user.Result, err)
    // if errors != nil {
    //     return 
    // }

	// response := builder.BuildUserDTOs(user.Result)

	

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, user.Status, user.StatusMessage, response)
	log.Println("user fetched successfully")
}

// func (handler *UserHandler) InsertUser(ctx *gin.Context) {
// 	userData := common.NewUserOb()

// 	err := ctx.BindJSON(&userData)

// 	if err != nil {
// 		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		log.Println("binding user details from json is failed:", err)
// 		common.SendError(nil, http.StatusBadRequest, "", err)
// 		return
// 	}

// 	newUser, err := handler.userManager.CreateUser(userData)

// 	errors := errorHandlers.HandleErrorResponse(ctx, newUser.Result, err)
//     if errors != nil {
//         return 
//     }

// 	response := builder.BuildUserDTOs(newUser.Result)

// 	fmt.Println("error", err)

// 	if err != nil {
// 		log.Println("user insertion failed", err)
// 		common.SendError(ctx, http.StatusInternalServerError, "", err)
// 		return
// 	}

// 	common.SendSuccess(ctx, http.StatusOK, newUser.Message, response)
// 	log.Println("user created successfully")
// }

// func (handler *UserHandler) UpdateUser(ctx *gin.Context) {
// 	// Create a new user object
// 	userData := common.NewUserOb()

// 	// Bind the request body to the userData object
// 	err := ctx.BindJSON(&userData)
// 	if err != nil {
// 		log.Println("binding user details from JSON failed:", err)
// 		common.SendError(ctx, http.StatusBadRequest, "binding user details from JSON failed", err)
// 		return
// 	}

// 	// Extract the user ID from the URL parameters
// 	userIdStr, ok := ctx.Params.Get("userId")
// 	if !ok {
// 		log.Println("user ID is missing in the request")
// 		common.SendError(ctx, http.StatusBadRequest, "user ID is missing in the request", fmt.Errorf("user ID is required"))
// 		return
// 	}

// 	// Convert string userId to uint
// 	userId, err := strconv.ParseUint(userIdStr, 10, 32)
// 	if err != nil {
// 		log.Println("invalid user ID format:", err)
// 		common.SendError(ctx, http.StatusBadRequest, "invalid user ID format", fmt.Errorf("invalid user ID format"))
// 		return
// 	}

// 	// Assign the user ID to the userData object
// 	userData.UserID = uint(userId)

// 	// Call the update method in the user manager
// 	updatedUser, err := handler.userManager.UpdateUser(userData)

// 	errors := errorHandlers.HandleErrorResponse(ctx, updatedUser.Result, err)
//     if errors != nil {
//         return 
//     }

// 	response := builder.BuildUserDTOs(updatedUser.Result)

// 	if err != nil {
// 		log.Println("user update failed:", err)
// 		common.SendError(ctx, http.StatusInternalServerError, "user update failed", err)
// 		return
// 	}

// 	// Send success response
// 	common.SendSuccess(ctx, http.StatusOK, updatedUser.Message, response)
// 	log.Println("user updated successfully")
// }
