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

type DepartmentHandler struct {
	groupName         string
	departmentManager managers.DepartmentManager
}

func NewDepartmentHandler(departmentManager managers.DepartmentManager) *DepartmentHandler {
	return &DepartmentHandler{
		"api/department",
		departmentManager,
	}
}

func (handler *DepartmentHandler) RegisterApis(r *gin.Engine) {
	departmentGroup := r.Group(handler.groupName)
	departmentGroup.GET("/list", handler.DepartmentList)
	departmentGroup.GET("/:departmentId", handler.GetADepartment)
	departmentGroup.POST("/create", handler.InsertDepartment)
	departmentGroup.PUT("/update/:departmentId", handler.UpdateDepartment)
	departmentGroup.DELETE("/delete/:departmentId", handler.DeleteDepartment)
}

func (handler *DepartmentHandler) DepartmentList(ctx *gin.Context) {
	departmentManagerResponse, err := handler.departmentManager.GetDepartments()
	if common.HandleServerError(ctx, departmentManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("departments.Data type: %T\n", departmentManagerResponse.Data)
	fmt.Println("departments.Data content:", departmentManagerResponse.Data)

	//Use ParseJSONResponse to parse the departmentManagerResponse data
	parsedData := common.ParseJSONResponse(departmentManagerResponse, ctx)

	fmt.Println("parsedData", parsedData)

	response := builder.BuildDepartmentDTOs(parsedData)

	fmt.Println("response", response)

	common.SendSuccess(ctx, http.StatusOK, departmentManagerResponse.Status, departmentManagerResponse.StatusMessage, response)

	log.Println("departments fetched successfully")
}

func (handler *DepartmentHandler) InsertDepartment(ctx *gin.Context) {
	departmentData := requestData.NewDepartmentObj()

	// Bind the incoming JSON to the departmentData object
	if err := common.BindJSONAndValidate(ctx, &departmentData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	departmentManagerResponse, err := handler.departmentManager.CreateDepartment(departmentData)

	if common.HandleServerError(ctx, departmentManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	//Use ParseJSONResponse to parse the departmentManagerResponse data
	parsedData := common.ParseJSONResponse(departmentManagerResponse, ctx)

	response := builder.BuildDepartmentDTOs(parsedData)

	common.SendSuccess(ctx, http.StatusCreated, departmentManagerResponse.Status, departmentManagerResponse.StatusMessage, response)
	log.Println("department created successfully")
}

func (handler *DepartmentHandler) UpdateDepartment(ctx *gin.Context) {
	// Create a new department object
	departmentData := requestData.NewDepartmentObj()

	// Bind the incoming JSON to the departmentData object
	if err := common.BindJSONAndValidate(ctx, &departmentData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	departmentId, err := common.GetParamAsUint(ctx, "departmentId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the department ID to the departmentData object
	departmentData.DepartmentID = uint(departmentId)

	// Call the update method in the department manager
	departmentManagerResponse, err := handler.departmentManager.UpdateDepartment(departmentData)

	if common.HandleServerError(ctx, departmentManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	//Use ParseJSONResponse to parse the departmentManagerResponse data
	parsedData := common.ParseJSONResponse(departmentManagerResponse, ctx)

	response := builder.BuildDepartmentDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, departmentManagerResponse.Status, departmentManagerResponse.StatusMessage, response)
	log.Println("department updated successfully")
}

func (handler *DepartmentHandler) DeleteDepartment(ctx *gin.Context) {
	// Create a new department object
	departmentData := requestData.NewDepartmentObj()

	departmentId, err := common.GetParamAsUint(ctx, "departmentId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the department ID to the departmentData object
	departmentData.DepartmentID = uint(departmentId)

	// Call the delete method in the department manager
	departmentManagerResponse, err := handler.departmentManager.DeleteDepartment(departmentData)

	if common.HandleServerError(ctx, departmentManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("department.Data type: %T\n", departmentManagerResponse.Data)
	fmt.Println("department.Data content:", departmentManagerResponse.Data)

	//Use ParseJSONResponse to parse the departmentManagerResponse data
	parsedData := common.ParseJSONResponse(departmentManagerResponse, ctx)
	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildDepartmentDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, departmentManagerResponse.Status, departmentManagerResponse.StatusMessage, response)
	log.Println("department deleted successfully")
}

func (handler *DepartmentHandler) GetADepartment(ctx *gin.Context) {
	// Create a new department object
	departmentData := requestData.NewDepartmentObj()

	departmentId, err := common.GetParamAsUint(ctx, "departmentId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the department ID to the departmentData object
	departmentData.DepartmentID = uint(departmentId)

	// Call the get a department method in the department manager
	departmentManagerResponse, err := handler.departmentManager.GetADepartment(departmentData)
	if common.HandleServerError(ctx, departmentManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("department.Data type: %T\n", departmentManagerResponse.Data)
	fmt.Println("department.Data content:", departmentManagerResponse.Data)

	//Use ParseJSONResponse to parse the departmentManagerResponse data
	parsedData := common.ParseJSONResponse(departmentManagerResponse, ctx)

	fmt.Println("parsedData", parsedData)

	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildDepartmentDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, departmentManagerResponse.Status, departmentManagerResponse.StatusMessage, response)
	log.Println("department fetched successfully")
}
