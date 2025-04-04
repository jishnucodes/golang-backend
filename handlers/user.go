package handlers

import (
	// "clinic-management/backend/builder"
	// "clinic-management/backend/builder"
	"clinic-management/backend/builder"
	"clinic-management/backend/common"

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
	if common.HandleServerError(ctx, userManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Println("Result", userManagerResponse)

	//Use ParseJSONResponse to parse the userManagerResponse data
	parsedData := common.ParseJSONResponse(userManagerResponse, ctx)

	// If no error message exists, continue with building the response
	response := builder.BuildUserDTOs(parsedData)

	// If everything is fine, send success response
	common.SendSuccess(ctx, http.StatusOK, userManagerResponse.Status, userManagerResponse.StatusMessage, response)

	log.Println("user logged in successfully")
}

func (handler *UserHandler) UserList(ctx *gin.Context) {

	userManagerResponse, err := handler.userManager.GetUsers()
	//this error block will work when the store procedure catch block catches an error
	if common.HandleServerError(ctx, userManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("users.Data type: %T\n", userManagerResponse.Data)
	fmt.Println("users.Data content:", userManagerResponse.Data)

	//Use ParseJSONResponse to parse the userManagerResponse data
	parsedData := common.ParseJSONResponse(userManagerResponse, ctx)

	response := builder.BuildUserDTOs(parsedData)

	fmt.Println("response", response)

	common.SendSuccess(ctx, http.StatusOK, userManagerResponse.Status, userManagerResponse.StatusMessage, response)

	log.Println("users fetched successfully")

}

func (handler *UserHandler) GetAUser(ctx *gin.Context) {
	// Create a new user object
	userData := common.NewUserOb()

	userId, err := common.GetParamAsUint(ctx, "userId")
	if err != nil {
		return // The function already sends an error response, so just return
	}
	// Assign the user ID to the userData object
	userData.UserID = uint(userId)

	// Call the get a user method in the user manager
	userManagerResponse, err := handler.userManager.GetAUser(userData)
	if common.HandleServerError(ctx, userManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("user.Data type: %T\n", userManagerResponse.Data)
	fmt.Println("user.Data content:", userManagerResponse.Data)

	//Use ParseJSONResponse to parse the userManagerResponse data
	parsedData := common.ParseJSONResponse(userManagerResponse, ctx)

	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildUserDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, userManagerResponse.Status, userManagerResponse.StatusMessage, response)
	log.Println("user fetched successfully")
}

func (handler *UserHandler) InsertUser(ctx *gin.Context) {
	userData := common.NewUserOb()

	// Bind the incoming JSON to the userData object
	if err := common.BindJSONAndValidate(ctx, &userData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	userManagerResponse, err := handler.userManager.CreateUser(userData)

	if common.HandleServerError(ctx, userManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	//Use ParseJSONResponse to parse the userManagerResponse data
	parsedData := common.ParseJSONResponse(userManagerResponse, ctx)

	response := builder.BuildUserDTOs(parsedData)

	common.SendSuccess(ctx, http.StatusCreated, userManagerResponse.Status, userManagerResponse.StatusMessage, response)
	log.Println("user created successfully")
}

func (handler *UserHandler) UpdateUser(ctx *gin.Context) {
	// Create a new user object
	userData := common.NewUserOb()

	// Bind the incoming JSON to the userData object
	if err := common.BindJSONAndValidate(ctx, &userData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	userId, err := common.GetParamAsUint(ctx, "userId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the user ID to the userData object
	userData.UserID = uint(userId)

	// Call the update method in the user manager
	userManagerResponse, err := handler.userManager.UpdateUser(userData)

	if common.HandleServerError(ctx, userManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	//Use ParseJSONResponse to parse the userManagerResponse data
	parsedData := common.ParseJSONResponse(userManagerResponse, ctx)

	response := builder.BuildUserDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, userManagerResponse.Status, userManagerResponse.StatusMessage, response)
	log.Println("user updated successfully")
}

func (handler *UserHandler) DeleteAUser(ctx *gin.Context) {
	// Create a new user object
	userData := common.NewUserOb()

	userId, err := common.GetParamAsUint(ctx, "userId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the user ID to the userData object
	userData.UserID = uint(userId)

	// Call the get a user method in the user manager
	userManagerResponse, err := handler.userManager.DeleteAUser(userData)

	if common.HandleServerError(ctx, userManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("user.Data type: %T\n", userManagerResponse.Data)
	fmt.Println("user.Data content:", userManagerResponse.Data)

	//Use ParseJSONResponse to parse the userManagerResponse data
	parsedData := common.ParseJSONResponse(userManagerResponse, ctx)
	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildUserDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, userManagerResponse.Status, userManagerResponse.StatusMessage, response)
	log.Println("user deleted successfully")
}
