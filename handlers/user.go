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
	userGroup.POST("/create", handler.InsertUser)
	userGroup.POST("/login", handler.Login)
	userGroup.PUT("/update/:userId", handler.UpdateUser)
	userGroup.DELETE("/delete/:userId", handler.DeleteAUser)
}




func (handler *UserHandler) Login(ctx *gin.Context) {
    userData := common.NewUserOb()

    // Bind the incoming JSON to the userData object
    if err := common.BindJSONAndValidate(ctx, &userData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}
	
    // Call the Login method and get the result
    userManagerResponse, err := handler.userManager.Login(userData)

	// If there was an error from the login function, this error block will work when the store procedure catch block catches an error
    if common.HandleManagerError(ctx, userManagerResponse, err, "User login") {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Println("Result", userManagerResponse)

	//Use ParseJSONResponse to parse the userManagerResponse data
    parsedData, parseErr := common.ParseJSONResponse(userManagerResponse, ctx)
    if parseErr != nil {
        return // `ParseJSONResponse` already sends an error response, no need to send another one
    }

    // If no error message exists, continue with building the response
    response := builder.BuildUserDTOs(parsedData)
    
    // If everything is fine, send success response
	common.SendSuccess(ctx, http.StatusOK, userManagerResponse.Status, userManagerResponse.StatusMessage, response)

    log.Println("user logged in successfully")
}

func (handler *UserHandler) UserList(ctx *gin.Context) {

	userManagerResponse, err := handler.userManager.GetUsers()
	//this error block will work when the store procedure catch block catches an error
	if common.HandleManagerError(ctx, userManagerResponse, err, "User fetching") {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("users.Data type: %T\n", userManagerResponse.Data)
	fmt.Println("users.Data content:", userManagerResponse.Data)

	// users.Data is a string, so directly unmarshal
	var parsedData []map[string]interface{}

	if userManagerResponse.Data != "" && json.Valid([]byte(userManagerResponse.Data))  {

		unmarshalError := json.Unmarshal([]byte(userManagerResponse.Data), &parsedData)
		if unmarshalError != nil {
			log.Println("Failed to parse user data:", unmarshalError)
			common.SendError(ctx, http.StatusBadRequest, 0, "Failed to parse user data", unmarshalError)
			return
		}
	}else {
		common.SendError(
			ctx, 
			http.StatusInternalServerError, 
			userManagerResponse.Status, 
			userManagerResponse.StatusMessage, 
			fmt.Errorf(userManagerResponse.StatusMessage),
		)
		log.Println("fetching users list failed:", userManagerResponse.StatusMessage)
		return
	}

	response := builder.BuildUserDTOs(parsedData)

	fmt.Println("response", response)

	common.SendSuccess(ctx, http.StatusOK, userManagerResponse.Status, userManagerResponse.StatusMessage, response)

	log.Println("users fetched successfully")

}

func (handler *UserHandler) GetAUser(ctx *gin.Context) {
	// Create a new user object
	userData := common.NewUserOb()

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
		common.SendError(ctx, http.StatusBadRequest,0, "invalid user ID format", fmt.Errorf("invalid user ID format"))
		return
	}

	// Assign the user ID to the userData object
	userData.UserID = uint(userId)

	// Call the get a user method in the user manager
	userManagerResponse, err := handler.userManager.GetAUser(userData)
	//this error block will work when the store procedure catch block catches an error
	if err != nil {
		log.Println("user update failed:", err)
		common.SendError(ctx, http.StatusInternalServerError, userManagerResponse.Status, userManagerResponse.StatusMessage, err)
		return
	}

	fmt.Printf("user.Data type: %T\n", userManagerResponse.Data)
	fmt.Println("user.Data content:", userManagerResponse.Data)

	// users.Data is a string, so directly unmarshal
	var parsedData []map[string]interface{}
	if userManagerResponse.Data != "" && json.Valid([]byte(userManagerResponse.Data)) {

		unmarshalError := json.Unmarshal([]byte(userManagerResponse.Data), &parsedData)
		if unmarshalError != nil {
			log.Println("Failed to parse user data:", unmarshalError)
			common.SendError(ctx, http.StatusBadRequest, 0, "Failed to parse user data", unmarshalError)
			return
		}
	}else {
		common.SendError(
			ctx, 
			http.StatusInternalServerError, 
			userManagerResponse.Status, 
			userManagerResponse.StatusMessage, 
			fmt.Errorf(userManagerResponse.StatusMessage),
		)
		log.Printf("Fetching user of ID %d failed: %s", userId, userManagerResponse.StatusMessage)
		return
	}

	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildUserDTOs(parsedData)
	
	// Send success response
	common.SendSuccess(ctx, http.StatusOK, userManagerResponse.Status, userManagerResponse.StatusMessage, response)
	log.Println("user fetched successfully")
}

func (handler *UserHandler) InsertUser(ctx *gin.Context) {
	userData := common.NewUserOb()

	err := ctx.BindJSON(&userData)

	if err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("binding user details from json is failed:", err)
		common.SendError(nil, http.StatusBadRequest, 0, "binding user details from json is failed", err)
		return
	}

	userManagerResponse, err := handler.userManager.CreateUser(userData)

	if err != nil {
		log.Println("user update failed:", err)
		common.SendError(ctx, http.StatusInternalServerError, userManagerResponse.Status, userManagerResponse.StatusMessage, err)
		return
	}

	var parsedData []map[string]interface{}
	if userManagerResponse.Data != "" && json.Valid([]byte(userManagerResponse.Data)) {

		unmarshalError := json.Unmarshal([]byte(userManagerResponse.Data), &parsedData)
		if unmarshalError != nil {
			log.Println("Failed to parse user data:", unmarshalError)
			common.SendError(ctx, http.StatusBadRequest, 0, "Failed to parse user data", unmarshalError)
			return
		}
	}else {
		common.SendError(
			ctx, 
			http.StatusInternalServerError, 
			userManagerResponse.Status, 
			userManagerResponse.StatusMessage, 
			fmt.Errorf(userManagerResponse.StatusMessage),
		)
		log.Printf("Creating a user failed: %s", userManagerResponse.StatusMessage)
		return
	}

	response := builder.BuildUserDTOs(parsedData)

	common.SendSuccess(ctx, http.StatusCreated,userManagerResponse.Status, userManagerResponse.StatusMessage, response)
	log.Println("user created successfully")
}

func (handler *UserHandler) UpdateUser(ctx *gin.Context) {
	// Create a new user object
	userData := common.NewUserOb()

	// Bind the request body to the userData object
	err := ctx.BindJSON(&userData)
	if err != nil {
		log.Println("binding user details from JSON failed:", err)
		common.SendError(nil, http.StatusBadRequest, 0, "binding user details from json is failed", err)
		return
	}

	// Extract the user ID from the URL parameters
	userIdStr, ok := ctx.Params.Get("userId")
	if !ok {
		log.Println("user ID is missing in the request")
		common.SendError(ctx, http.StatusBadRequest, 0, "user ID is missing in the request", fmt.Errorf("user ID is required"))
		return
	}

	// Convert string userId to uint
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		log.Println("invalid user ID format:", err)
		common.SendError(ctx, http.StatusBadRequest, 0, "invalid user ID format", fmt.Errorf("invalid user ID format"))
		return
	}

	// Assign the user ID to the userData object
	userData.UserID = uint(userId)

	// Call the update method in the user manager
	userManagerResponse, err := handler.userManager.UpdateUser(userData)

	if err != nil {
		log.Println("user update failed:", err)
		common.SendError(ctx, http.StatusInternalServerError, userManagerResponse.Status, userManagerResponse.StatusMessage, err)
		return
	}

	var parsedData []map[string]interface{}
	if userManagerResponse.Data != "" && json.Valid([]byte(userManagerResponse.Data)) {

		unmarshalError := json.Unmarshal([]byte(userManagerResponse.Data), &parsedData)
		if unmarshalError != nil {
			log.Println("Failed to parse user data:", unmarshalError)
			common.SendError(ctx, http.StatusBadRequest, 0, "Failed to parse user data", unmarshalError)
			return
		}
	}else {
		common.SendError(
			ctx, 
			http.StatusInternalServerError, 
			userManagerResponse.Status, 
			userManagerResponse.StatusMessage, 
			fmt.Errorf(userManagerResponse.StatusMessage),
		)
		log.Printf("Updating a user of ID %d failed: %s", userId, userManagerResponse.StatusMessage)
		return
	}

	response := builder.BuildUserDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, userManagerResponse.Status, userManagerResponse.StatusMessage, response)
	log.Println("user updated successfully")
}

func (handler *UserHandler) DeleteAUser(ctx *gin.Context) {
	// Create a new user object
	userData := common.NewUserOb()

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
		common.SendError(ctx, http.StatusBadRequest,0, "invalid user ID format", fmt.Errorf("invalid user ID format"))
		return
	}

	// Assign the user ID to the userData object
	userData.UserID = uint(userId)

	// Call the get a user method in the user manager
	userManagerResponse, err := handler.userManager.DeleteAUser(userData)
	//this error block will work when the store procedure catch block catches an error
	if err != nil {
		log.Println("user delete failed:", err)
		common.SendError(ctx, http.StatusInternalServerError, userManagerResponse.Status, userManagerResponse.StatusMessage, err)
		return
	}

	fmt.Printf("user.Data type: %T\n", userManagerResponse.Data)
	fmt.Println("user.Data content:", userManagerResponse.Data)

	// users.Data is a string, so directly unmarshal
	var parsedData []map[string]interface{}
	if userManagerResponse.Data != "" && json.Valid([]byte(userManagerResponse.Data)) {

		unmarshalError := json.Unmarshal([]byte(userManagerResponse.Data), &parsedData)
		if unmarshalError != nil {
			log.Println("Failed to parse user data:", unmarshalError)
			common.SendError(ctx, http.StatusBadRequest, 0, "Failed to parse user data", unmarshalError)
			return
		}
	}else if userManagerResponse.Status == 0  {
		common.SendError(
			ctx, 
			http.StatusInternalServerError, 
			userManagerResponse.Status, 
			userManagerResponse.StatusMessage, 
			fmt.Errorf(userManagerResponse.StatusMessage),
		)
		log.Printf("Deleting user of ID %d failed: %s", userId, userManagerResponse.StatusMessage)
		return
	}

	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildUserDTOs(parsedData)
	
	// Send success response
	common.SendSuccess(ctx, http.StatusOK, userManagerResponse.Status, userManagerResponse.StatusMessage, response)
	log.Println("user deleted successfully")
}
