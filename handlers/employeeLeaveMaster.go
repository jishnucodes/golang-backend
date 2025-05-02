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

type EmployeeLeaveHandler struct {
	groupName       string
	employeeLeaveManager managers.EmployeeLeaveManager
}

func NewEmployeeLeaveHandler(employeeLeaveManager managers.EmployeeLeaveManager) *EmployeeLeaveHandler {
	return &EmployeeLeaveHandler{
		"api/employeeLeave",
		employeeLeaveManager,
	}
}

func (handler *EmployeeLeaveHandler) RegisterApis(r *gin.Engine) {
	employeeGroup := r.Group(handler.groupName)
	employeeGroup.GET("/list", handler.EmployeeLeaveList)
	employeeGroup.GET("/:employeeId", handler.GetAEmployeeLeave)
	employeeGroup.POST("/create", handler.InsertEmployeeLeave)
	employeeGroup.PUT("/update/:employeeId", handler.UpdateEmployeeLeave)
	employeeGroup.DELETE("/delete/:employeeId", handler.DeleteEmployeeLeave)
}

func (handler *EmployeeLeaveHandler) EmployeeLeaveList(ctx *gin.Context) {
	employeeLeaveManagerResponse, err := handler.employeeLeaveManager.GetEmployeeLeaves()
	if common.HandleServerError(ctx, employeeLeaveManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("employeeLeaves.Data type: %T\n", employeeLeaveManagerResponse.Data)
	fmt.Println("employeeLeaves.Data content:", employeeLeaveManagerResponse.Data)

	//Use ParseJSONResponse to parse the employeeManagerResponse data
	parsedData := common.ParseJSONResponse(employeeLeaveManagerResponse, ctx)

	response := builder.BuildEmployeeLeaveDTOs(parsedData)

	fmt.Println("response", response)

	common.SendSuccess(ctx, http.StatusOK, employeeLeaveManagerResponse.Status, employeeLeaveManagerResponse.StatusMessage, response)

	log.Println("employeeLeaves fetched successfully")
}

func (handler *EmployeeLeaveHandler) GetAEmployeeLeave(ctx *gin.Context) {
	// Create a new employeeLeave object
	employeeLeaveData := requestData.NewEmployeeLeaveObj()

	employeeId, err := common.GetParamAsUint(ctx, "employeeId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the employee ID to the employeeData object
	employeeLeaveData.ID = uint(employeeId)

	// Call the get an employeeLeave method in the employeeLeave manager
	employeeLeaveManagerResponse, err := handler.employeeLeaveManager.GetAEmployeeLeave(employeeLeaveData)
	if common.HandleServerError(ctx, employeeLeaveManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("employeeLeave.Data type: %T\n", employeeLeaveManagerResponse.Data)
	fmt.Println("employeeLeave.Data content:", employeeLeaveManagerResponse.Data)

	//Use ParseJSONResponse to parse the employeeManagerResponse data
	parsedData := common.ParseJSONResponse(employeeLeaveManagerResponse, ctx)

	fmt.Println("parsedData", parsedData)

	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildEmployeeLeaveDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, employeeLeaveManagerResponse.Status, employeeLeaveManagerResponse.StatusMessage, response)
	log.Println("employeeLeave fetched successfully")
}

func (handler *EmployeeLeaveHandler) InsertEmployeeLeave(ctx *gin.Context) {
	employeeLeaveData := requestData.NewEmployeeLeaveObj()

	// Bind the incoming JSON to the employeeLeaveData object
	if err := common.BindJSONAndValidate(ctx, &employeeLeaveData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	employeeLeaveManagerResponse, err := handler.employeeLeaveManager.CreateEmployeeLeave(employeeLeaveData)

	if common.HandleServerError(ctx, employeeLeaveManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	//Use ParseJSONResponse to parse the employeeManagerResponse data
	parsedData := common.ParseJSONResponse(employeeLeaveManagerResponse, ctx)

	response := builder.BuildEmployeeLeaveDTOs(parsedData)

	common.SendSuccess(ctx, http.StatusCreated, employeeLeaveManagerResponse.Status, employeeLeaveManagerResponse.StatusMessage, response)
	log.Println("employeeLeave created successfully")
}

func (handler *EmployeeLeaveHandler) UpdateEmployeeLeave(ctx *gin.Context) {
	// Create a new employeeLeave object
	employeeLeaveData := requestData.NewEmployeeLeaveObj()

	// Bind the incoming JSON to the employeeLeaveData object
	if err := common.BindJSONAndValidate(ctx, &employeeLeaveData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	employeeId, err := common.GetParamAsUint(ctx, "employeeId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the employee ID to the employeeLeaveData object
	employeeLeaveData.ID = uint(employeeId)

	// Call the update method in the employeeLeave manager
	employeeLeaveManagerResponse, err := handler.employeeLeaveManager.UpdateEmployeeLeave(employeeLeaveData)

	if common.HandleServerError(ctx, employeeLeaveManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	//Use ParseJSONResponse to parse the employeeManagerResponse data
	parsedData := common.ParseJSONResponse(employeeLeaveManagerResponse, ctx)

	response := builder.BuildEmployeeLeaveDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, employeeLeaveManagerResponse.Status, employeeLeaveManagerResponse.StatusMessage, response)
	log.Println("employeeLeave updated successfully")
}

func (handler *EmployeeLeaveHandler) DeleteEmployeeLeave(ctx *gin.Context) {
	// Create a new employeeLeave object
	employeeLeaveData := requestData.NewEmployeeLeaveObj()

	employeeId, err := common.GetParamAsUint(ctx, "employeeId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the employee ID to the employeeLeaveData object
	employeeLeaveData.ID = uint(employeeId)

	// Call the delete method in the employeeLeave manager
	employeeLeaveManagerResponse, err := handler.employeeLeaveManager.DeleteEmployeeLeave(employeeLeaveData)

	if common.HandleServerError(ctx, employeeLeaveManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("employeeLeave.Data type: %T\n", employeeLeaveManagerResponse.Data)
	fmt.Println("employeeLeave.Data content:", employeeLeaveManagerResponse.Data)

	//Use ParseJSONResponse to parse the employeeManagerResponse data
	parsedData := common.ParseJSONResponse(employeeLeaveManagerResponse, ctx)
	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildEmployeeLeaveDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, employeeLeaveManagerResponse.Status, employeeLeaveManagerResponse.StatusMessage, response)
	log.Println("employeeLeave deleted successfully")
}
