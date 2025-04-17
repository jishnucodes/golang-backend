package handlers

import (
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"clinic-management/backend/managers"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	groupName   string
	roleManager managers.RoleManager
}

func NewRoleHandler(roleManager managers.RoleManager) *RoleHandler {
	return &RoleHandler{
		"api/role",
		roleManager,
	}
}

func (handler *RoleHandler) RegisterApis(r *gin.Engine) {
	roleGroup := r.Group(handler.groupName)
	roleGroup.GET("/list", handler.RoleList)
	// roleGroup.GET("/:roleId", handler.GetARole)
	roleGroup.POST("/create", handler.InsertRole)
	roleGroup.PUT("/update/:roleId", handler.UpdateRole)
	roleGroup.DELETE("/delete/:roleId", handler.DeleteARole)
}

func (handler *RoleHandler) RoleList(ctx *gin.Context) {
	roleManagerResponse, err := handler.roleManager.GetRoles()
	if common.HandleServerError(ctx, roleManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("roles.Data type: %T\n", roleManagerResponse.Data)
	fmt.Println("roles.Data content:", roleManagerResponse.Data)

	//Use ParseJSONResponse to parse the roleManagerResponse data
	parsedData := common.ParseJSONResponse(roleManagerResponse, ctx)

	response := builder.BuildRoleDTOs(parsedData)

	fmt.Println("response", response)

	common.SendSuccess(ctx, http.StatusOK, roleManagerResponse.Status, roleManagerResponse.StatusMessage, response)

	log.Println("roles fetched successfully")
}

// func (handler *RoleHandler) GetARole(ctx *gin.Context) {
// 	// Create a new role object
// 	roleData := requestData.NewRoleObj()

// 	roleId, err := common.GetParamAsUint(ctx, "roleId")
// 	if err != nil {
// 		return // The function already sends an error response, so just return
// 	}
// 	// Assign the role ID to the roleData object
// 	roleData.Id = uint(roleId)

// 	// Call the get a role method in the role manager
// 	roleManagerResponse, err := handler.roleManager.GetARole(roleData)
// 	if common.HandleServerError(ctx, roleManagerResponse, err) {
// 		return // Exit if an error occurred (response is already sent)
// 	}

// 	fmt.Printf("role.Data type: %T\n", roleManagerResponse.Data)
// 	fmt.Println("role.Data content:", roleManagerResponse.Data)

// 	//Use ParseJSONResponse to parse the roleManagerResponse data
// 	parsedData := common.ParseJSONResponse(roleManagerResponse, ctx)

// 	fmt.Println("parsedData", parsedData)

// 	fmt.Printf("Parsed data type: %T\n", parsedData)

// 	response := builder.BuildRoleDTOs(parsedData)

// 	// Send success response
// 	common.SendSuccess(ctx, http.StatusOK, roleManagerResponse.Status, roleManagerResponse.StatusMessage, response)
// 	log.Println("role fetched successfully")
// }

func (handler *RoleHandler) InsertRole(ctx *gin.Context) {
	roleData := requestData.NewRoleObj()

	// Bind the incoming JSON to the roleData object
	if err := common.BindJSONAndValidate(ctx, &roleData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	roleManagerResponse, err := handler.roleManager.CreateRole(roleData)

	if common.HandleServerError(ctx, roleManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	//Use ParseJSONResponse to parse the roleManagerResponse data
	parsedData := common.ParseJSONResponse(roleManagerResponse, ctx)

	response := builder.BuildRoleDTOs(parsedData)

	common.SendSuccess(ctx, http.StatusCreated, roleManagerResponse.Status, roleManagerResponse.StatusMessage, response)
	log.Println("role created successfully")
}

func (handler *RoleHandler) UpdateRole(ctx *gin.Context) {
	// Create a new role object
	roleData := requestData.NewRoleObj()

	// Bind the incoming JSON to the roleData object
	if err := common.BindJSONAndValidate(ctx, &roleData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	roleId, err := common.GetParamAsUint(ctx, "roleId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the role ID to the roleData object
	roleData.Id = uint(roleId)

	// Call the update method in the role manager
	roleManagerResponse, err := handler.roleManager.UpdateRole(roleData)

	if common.HandleServerError(ctx, roleManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	//Use ParseJSONResponse to parse the roleManagerResponse data
	parsedData := common.ParseJSONResponse(roleManagerResponse, ctx)

	response := builder.BuildRoleDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, roleManagerResponse.Status, roleManagerResponse.StatusMessage, response)
	log.Println("role updated successfully")
}

func (handler *RoleHandler) DeleteARole(ctx *gin.Context) {
	// Create a new role object
	roleData := requestData.NewRoleObj()

	roleId, err := common.GetParamAsUint(ctx, "roleId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the role ID to the roleData object
	roleData.Id = uint(roleId)

	// Call the delete method in the role manager
	roleManagerResponse, err := handler.roleManager.DeleteARole(roleData)

	if common.HandleServerError(ctx, roleManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("role.Data type: %T\n", roleManagerResponse.Data)
	fmt.Println("role.Data content:", roleManagerResponse.Data)

	//Use ParseJSONResponse to parse the roleManagerResponse data
	parsedData := common.ParseJSONResponse(roleManagerResponse, ctx)
	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildRoleDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, roleManagerResponse.Status, roleManagerResponse.StatusMessage, response)
	log.Println("role deleted successfully")
}
