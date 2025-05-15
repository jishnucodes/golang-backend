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

type DoctorHandler struct {
	groupName     string
	doctorManager managers.DoctorManager
}

func NewDoctorHandler(doctorManager managers.DoctorManager) *DoctorHandler {
	return &DoctorHandler{
		"api/doctor",
		doctorManager,
	}
}

func (handler *DoctorHandler) RegisterApis(r *gin.Engine) {
	doctorGroup := r.Group(handler.groupName)
	doctorGroup.GET("/list", handler.DoctorList)
	doctorGroup.GET("/:employeeId", handler.GetADoctor)
	doctorGroup.POST("/create", handler.InsertDoctor)
	doctorGroup.PUT("/update/:doctorId", handler.UpdateDoctor)
	doctorGroup.DELETE("/delete/:doctorId", handler.DeleteADoctor)
}

func (handler *DoctorHandler) DoctorList(ctx *gin.Context) {

	query := &requestData.DoctorSearchQuery{}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		common.SendError(ctx, http.StatusBadRequest, 0, "Invalid query parameters", err)
		return
	}
	doctorManagerResponse, err := handler.doctorManager.GetDoctors(query)
	// This error block will work when the store procedure catch block catches an error
	if common.HandleServerError(ctx, doctorManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("doctors.Data type: %T\n", doctorManagerResponse.Data)
	fmt.Println("doctors.Data content:", doctorManagerResponse.Data)

	// Use ParseJSONResponse to parse the doctorManagerResponse data
	parsedData := common.ParseJSONResponse(doctorManagerResponse, ctx)
	fmt.Println("parsedData", parsedData)

	response := builder.BuildEmployeeDTOs(parsedData)

	fmt.Println("response", response)

	common.SendSuccess(ctx, http.StatusOK, doctorManagerResponse.Status, doctorManagerResponse.StatusMessage, response)

	log.Println("doctors fetched successfully")
}

func (handler *DoctorHandler) GetADoctor(ctx *gin.Context) {
	// Create a new doctor object
	doctorData := requestData.NewDoctorObj()

	employeeId, err := common.GetParamAsUint(ctx, "employeeId")
	if err != nil {
		return // The function already sends an error response, so just return
	}
	// Assign the doctor ID to the doctorData object
	doctorData.EmployeeID = uint(employeeId)

	// Call the get a doctor method in the doctor manager
	doctorManagerResponse, err := handler.doctorManager.GetADoctor(doctorData)
	if common.HandleServerError(ctx, doctorManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("doctor.Data type: %T\n", doctorManagerResponse.Data)
	fmt.Println("doctor.Data content:", doctorManagerResponse.Data)

	// Use ParseJSONResponse to parse the doctorManagerResponse data
	parsedData := common.ParseJSONResponse(doctorManagerResponse, ctx)

	fmt.Println("parsedData", parsedData)

	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildDoctorDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, doctorManagerResponse.Status, doctorManagerResponse.StatusMessage, response)
	log.Println("doctor fetched successfully")
}

func (handler *DoctorHandler) InsertDoctor(ctx *gin.Context) {
	doctorData := requestData.NewDoctorObj()

	// Bind the incoming JSON to the doctorData object
	if err := common.BindJSONAndValidate(ctx, &doctorData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	doctorManagerResponse, err := handler.doctorManager.CreateDoctor(doctorData)

	if common.HandleServerError(ctx, doctorManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	// Use ParseJSONResponse to parse the doctorManagerResponse data
	parsedData := common.ParseJSONResponse(doctorManagerResponse, ctx)

	response := builder.BuildDoctorDTOs(parsedData)

	common.SendSuccess(ctx, http.StatusCreated, doctorManagerResponse.Status, doctorManagerResponse.StatusMessage, response)
	log.Println("doctor created successfully")
}

func (handler *DoctorHandler) UpdateDoctor(ctx *gin.Context) {
	// Create a new doctor object
	doctorData := requestData.NewDoctorObj()

	// Bind the incoming JSON to the doctorData object
	if err := common.BindJSONAndValidate(ctx, &doctorData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	doctorId, err := common.GetParamAsUint(ctx, "doctorId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the doctor ID to the doctorData object
	doctorData.DoctorID = uint(doctorId)

	// Call the update method in the doctor manager
	doctorManagerResponse, err := handler.doctorManager.UpdateDoctor(doctorData)

	if common.HandleServerError(ctx, doctorManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	// Use ParseJSONResponse to parse the doctorManagerResponse data
	parsedData := common.ParseJSONResponse(doctorManagerResponse, ctx)

	response := builder.BuildDoctorDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, doctorManagerResponse.Status, doctorManagerResponse.StatusMessage, response)
	log.Println("doctor updated successfully")
}

func (handler *DoctorHandler) DeleteADoctor(ctx *gin.Context) {
	// Create a new doctor object
	doctorData := requestData.NewDoctorObj()

	doctorId, err := common.GetParamAsUint(ctx, "doctorId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the doctor ID to the doctorData object
	doctorData.DoctorID = uint(doctorId)

	// Call the delete a doctor method in the doctor manager
	doctorManagerResponse, err := handler.doctorManager.DeleteADoctor(doctorData)

	if common.HandleServerError(ctx, doctorManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("doctor.Data type: %T\n", doctorManagerResponse.Data)
	fmt.Println("doctor.Data content:", doctorManagerResponse.Data)

	// Use ParseJSONResponse to parse the doctorManagerResponse data
	parsedData := common.ParseJSONResponse(doctorManagerResponse, ctx)
	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildDoctorDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, doctorManagerResponse.Status, doctorManagerResponse.StatusMessage, response)
	log.Println("doctor deleted successfully")
}
