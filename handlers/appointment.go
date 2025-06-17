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

type AppointmentHandler struct {
	groupName          string
	appointmentManager managers.AppointmentManager
}


func NewAppointmentHandler(appointmentManager managers.AppointmentManager) *AppointmentHandler {
	return &AppointmentHandler{
		"api/appointment",
		appointmentManager,
	}
}

func (handler *AppointmentHandler) RegisterApis(r *gin.Engine) {
	appointmentGroup := r.Group(handler.groupName)
	appointmentGroup.GET("/list", handler.GetAppointmentsOfDoctorByDate)
	appointmentGroup.POST("/create", handler.InsertAppointment)
	appointmentGroup.PUT("/update/:appointmentId", handler.UpdateAppointment)
	appointmentGroup.GET("/active", handler.GetPatientByAppointmentActiveStatus)
}


func (handler *AppointmentHandler) GetAppointmentsOfDoctorByDate(ctx *gin.Context) {
	query := &requestData.AppointmentSearchQuery{}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		common.SendError(ctx, http.StatusBadRequest, 0, "Invalid query parameters", err)
		return
	}
	appointmentManagerResponse, err := handler.appointmentManager.GetAppointmentsOfDoctorByDate(query)
	// This error block will work when the store procedure catch block catches an error
	if common.HandleServerError(ctx, appointmentManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("appointments.Data type: %T\n", appointmentManagerResponse.Data)
	fmt.Println("appointments.Data content:", appointmentManagerResponse.Data)

	// Use ParseJSONResponse to parse the appointmentManagerResponse data
	parsedData := common.ParseJSONResponse(appointmentManagerResponse, ctx)
	fmt.Println("parsedData", parsedData)

	response := builder.BuildAppointmentDTOs(parsedData)

	fmt.Println("response", response)

	common.SendSuccess(ctx, http.StatusOK, appointmentManagerResponse.Status, appointmentManagerResponse.StatusMessage, response)

	log.Println("Appointments of doctor fetched successfully")
}

func (handler *AppointmentHandler) InsertAppointment(ctx *gin.Context) {
	appointmentData := requestData.NewAppointmentObj()

	// Bind the incoming JSON to the appointmentData object
	if err := common.BindJSONAndValidate(ctx, &appointmentData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	appointmentManagerResponse, err := handler.appointmentManager.CreateAppointment(appointmentData)

	if common.HandleServerError(ctx, appointmentManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	//Use ParseJSONResponse to parse the appointmentManagerResponse data
	parsedData := common.ParseJSONResponse(appointmentManagerResponse, ctx)

	response := builder.BuildAppointmentDTOs(parsedData)

	common.SendSuccess(ctx, http.StatusCreated, appointmentManagerResponse.Status, appointmentManagerResponse.StatusMessage, response)
	log.Println("appointment created successfully")
}

func (handler *AppointmentHandler) UpdateAppointment(ctx *gin.Context) {
	appointmentData := requestData.NewAppointmentObj()

	// Bind the incoming JSON to the appointmentData object
	if err := common.BindJSONAndValidate(ctx, &appointmentData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	appointmentId, err := common.GetParamAsUint(ctx, "appointmentId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the department ID to the departmentData object
	appointmentData.AppointmentID = uint(appointmentId)

	appointmentManagerResponse, err := handler.appointmentManager.UpdateAppointment(appointmentData)

	if common.HandleServerError(ctx, appointmentManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	//Use ParseJSONResponse to parse the appointmentManagerResponse data
	parsedData := common.ParseJSONResponse(appointmentManagerResponse, ctx)

	response := builder.BuildAppointmentDTOs(parsedData)

	common.SendSuccess(ctx, http.StatusCreated, appointmentManagerResponse.Status, appointmentManagerResponse.StatusMessage, response)
	log.Println("appointment updated successfully")
}

func (handler *AppointmentHandler) GetPatientByAppointmentActiveStatus(ctx *gin.Context) {
	query := &requestData.ActiveAppointmentPatientSearchQuery{}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		common.SendError(ctx, http.StatusBadRequest, 0, "Invalid query parameters", err)
		return
	}
	appointmentManagerResponse, err := handler.appointmentManager.GetPatientByAppointmentActiveStatus(query)
	// This error block will work when the store procedure catch block catches an error
	if common.HandleServerError(ctx, appointmentManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("appointments.Data type: %T\n", appointmentManagerResponse.Data)
	fmt.Println("appointments.Data content:", appointmentManagerResponse.Data)

	// Use ParseJSONResponse to parse the appointmentManagerResponse data
	parsedData := common.ParseJSONResponse(appointmentManagerResponse, ctx)
	fmt.Println("parsedData", parsedData)

	response := builder.BuildAppointmentDTOs(parsedData)

	fmt.Println("response", response)

	common.SendSuccess(ctx, http.StatusOK, appointmentManagerResponse.Status, appointmentManagerResponse.StatusMessage, response)

	log.Println("Appointments of doctor fetched successfully")
}
