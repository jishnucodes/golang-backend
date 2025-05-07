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

type EmployeeHandler struct {
	groupName       string
	employeeManager managers.EmployeeManager
}


func NewEmployeeHandler(employeeManager managers.EmployeeManager) *EmployeeHandler {
	return &EmployeeHandler{
		"api/employee",
		employeeManager,
	}
}

func (handler *EmployeeHandler) RegisterApis(r *gin.Engine) {
	employeeGroup := r.Group(handler.groupName)
	employeeGroup.GET("/list", handler.EmployeeList)
	employeeGroup.GET("/:employeeId", handler.GetAEmployee)
	employeeGroup.POST("/create", handler.InsertEmployee)
	employeeGroup.PUT("/update/:employeeId", handler.UpdateEmployee)
	employeeGroup.DELETE("/delete/:employeeId", handler.DeleteEmployee)
}

func (handler *EmployeeHandler) EmployeeList(ctx *gin.Context) {
	query := &requestData.SearchQuery{}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		common.SendError(ctx, http.StatusBadRequest, 0, "Invalid query parameters", err)
		return
	}

	employeeManagerResponse, err := handler.employeeManager.GetEmployees(query)
	if common.HandleServerError(ctx, employeeManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("employees.Data type: %T\n", employeeManagerResponse.Data)
	fmt.Println("employees.Data content:", employeeManagerResponse.Data)

	//Use ParseJSONResponse to parse the employeeManagerResponse data
	parsedData := common.ParseJSONResponse(employeeManagerResponse, ctx)

	response := builder.BuildEmployeeDTOs(parsedData)

	fmt.Println("response", response)

	common.SendSuccess(ctx, http.StatusOK, employeeManagerResponse.Status, employeeManagerResponse.StatusMessage, response)

	log.Println("employees fetched successfully")
}

func (handler *EmployeeHandler) GetAEmployee(ctx *gin.Context) {
	// Create a new employee object
	employeeData := requestData.NewEmployeeObj()

	employeeId, err := common.GetParamAsUint(ctx, "employeeId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the employee ID to the employeeData object
	employeeData.EmployeeID = uint(employeeId)

	// Call the get an employee method in the employee manager
	employeeManagerResponse, err := handler.employeeManager.GetAEmployee(employeeData)
	if common.HandleServerError(ctx, employeeManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("employee.Data type: %T\n", employeeManagerResponse.Data)
	fmt.Println("employee.Data content:", employeeManagerResponse.Data)

	//Use ParseJSONResponse to parse the employeeManagerResponse data
	parsedData := common.ParseJSONResponse(employeeManagerResponse, ctx)

	fmt.Println("parsedData", parsedData)

	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildEmployeeDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, employeeManagerResponse.Status, employeeManagerResponse.StatusMessage, response)
	log.Println("employee fetched successfully")
}

func (handler *EmployeeHandler) InsertEmployee(ctx *gin.Context) {
	employeeData := requestData.NewEmployeeObj()

	// Bind the incoming JSON to the employeeData object
	if err := common.BindJSONAndValidate(ctx, &employeeData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	employeeManagerResponse, err := handler.employeeManager.CreateEmployee(employeeData)

	if common.HandleServerError(ctx, employeeManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	//Use ParseJSONResponse to parse the employeeManagerResponse data
	parsedData := common.ParseJSONResponse(employeeManagerResponse, ctx)

	response := builder.BuildEmployeeDTOs(parsedData)

	common.SendSuccess(ctx, http.StatusCreated, employeeManagerResponse.Status, employeeManagerResponse.StatusMessage, response)
	log.Println("employee created successfully")
}

func (handler *EmployeeHandler) UpdateEmployee(ctx *gin.Context) {
	// Create a new employee object
	employeeData := requestData.NewEmployeeObj()

	// Bind the incoming JSON to the employeeData object
	if err := common.BindJSONAndValidate(ctx, &employeeData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	employeeId, err := common.GetParamAsUint(ctx, "employeeId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the employee ID to the employeeData object
	employeeData.EmployeeID = uint(employeeId)

	// Call the update method in the employee manager
	employeeManagerResponse, err := handler.employeeManager.UpdateEmployee(employeeData)

	if common.HandleServerError(ctx, employeeManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	//Use ParseJSONResponse to parse the employeeManagerResponse data
	parsedData := common.ParseJSONResponse(employeeManagerResponse, ctx)

	response := builder.BuildEmployeeDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, employeeManagerResponse.Status, employeeManagerResponse.StatusMessage, response)
	log.Println("employee updated successfully")
}

func (handler *EmployeeHandler) DeleteEmployee(ctx *gin.Context) {
	// Create a new employee object
	employeeData := requestData.NewEmployeeObj()

	employeeId, err := common.GetParamAsUint(ctx, "employeeId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the employee ID to the employeeData object
	employeeData.EmployeeID = uint(employeeId)

	// Call the delete method in the employee manager
	employeeManagerResponse, err := handler.employeeManager.DeleteEmployee(employeeData)

	if common.HandleServerError(ctx, employeeManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("employee.Data type: %T\n", employeeManagerResponse.Data)
	fmt.Println("employee.Data content:", employeeManagerResponse.Data)

	//Use ParseJSONResponse to parse the employeeManagerResponse data
	parsedData := common.ParseJSONResponse(employeeManagerResponse, ctx)
	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildEmployeeDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, employeeManagerResponse.Status, employeeManagerResponse.StatusMessage, response)
	log.Println("employee deleted successfully")
}
