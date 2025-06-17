package handlers

import (
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"clinic-management/backend/managers"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConsultationMedicationHandler struct {
	groupName                     string
	consultationMedicationManager managers.ConsultationMedicationManager
}

func NewConsultationMedicationHandler(consultationMedicationManager managers.ConsultationMedicationManager) *ConsultationMedicationHandler {
	return &ConsultationMedicationHandler{
		"api/medication",
		consultationMedicationManager,
	}
}

func (handler *ConsultationMedicationHandler) RegisterApis(r *gin.Engine) {
	consultationMedicationGroup := r.Group(handler.groupName)
	consultationMedicationGroup.POST("/create", handler.InsertConsultationMedicationDetails)
}

// func (handler *AppointmentHandler) GetAppointmentsOfDoctorByDate(ctx *gin.Context) {
// 	query := &requestData.AppointmentSearchQuery{}
// 	if err := ctx.ShouldBindQuery(&query); err != nil {
// 		common.SendError(ctx, http.StatusBadRequest, 0, "Invalid query parameters", err)
// 		return
// 	}
// 	appointmentManagerResponse, err := handler.appointmentManager.GetAppointmentsOfDoctorByDate(query)
// 	// This error block will work when the store procedure catch block catches an error
// 	if common.HandleServerError(ctx, appointmentManagerResponse, err) {
// 		return // Exit if an error occurred (response is already sent)
// 	}

// 	fmt.Printf("appointments.Data type: %T\n", appointmentManagerResponse.Data)
// 	fmt.Println("appointments.Data content:", appointmentManagerResponse.Data)

// 	// Use ParseJSONResponse to parse the appointmentManagerResponse data
// 	parsedData := common.ParseJSONResponse(appointmentManagerResponse, ctx)
// 	fmt.Println("parsedData", parsedData)

// 	response := builder.BuildAppointmentDTOs(parsedData)

// 	fmt.Println("response", response)

// 	common.SendSuccess(ctx, http.StatusOK, appointmentManagerResponse.Status, appointmentManagerResponse.StatusMessage, response)

// 	log.Println("Appointments of doctor fetched successfully")
// }

func (handler *ConsultationMedicationHandler) InsertConsultationMedicationDetails(ctx *gin.Context) {
	consultationMedicationData := requestData.NewConsultationMedicationObj()

	// Bind the incoming JSON to the consultation medication Data object
	if err := common.BindJSONAndValidate(ctx, &consultationMedicationData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	jsonData, err := json.MarshalIndent(consultationMedicationData, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	fmt.Println("consultationMedicationData:", string(jsonData))

	consultationMedicationManagerResponse, err := handler.consultationMedicationManager.CreateConsultationMedicationData(consultationMedicationData)

	if common.HandleServerError(ctx, consultationMedicationManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	//Use ParseJSONResponse to parse the consultationMedicationManagerResponse data
	parsedData := common.ParseJSONResponse(consultationMedicationManagerResponse, ctx)
	fmt.Println("parsed data: ", parsedData)

	response := builder.BuildConsultationMedicationDTOs(parsedData)
	fmt.Println("response: ", response)

	common.SendSuccess(ctx, http.StatusCreated, consultationMedicationManagerResponse.Status, consultationMedicationManagerResponse.StatusMessage, response)
	log.Println("consultation medication details created successfully")
}

// func (handler *AppointmentHandler) UpdateAppointment(ctx *gin.Context) {
// 	appointmentData := requestData.NewAppointmentObj()

// 	// Bind the incoming JSON to the appointmentData object
// 	if err := common.BindJSONAndValidate(ctx, &appointmentData); err != nil {
// 		return // Error response is already handled in BindJSONAndValidate
// 	}

// 	appointmentId, err := common.GetParamAsUint(ctx, "appointmentId")
// 	if err != nil {
// 		return // The function already sends an error response, so just return
// 	}

// 	// Assign the department ID to the departmentData object
// 	appointmentData.AppointmentID = uint(appointmentId)

// 	appointmentManagerResponse, err := handler.appointmentManager.UpdateAppointment(appointmentData)

// 	if common.HandleServerError(ctx, appointmentManagerResponse, err) {
// 		return // Exit if an error occurred (response is already sent)
// 	}

// 	//Use ParseJSONResponse to parse the appointmentManagerResponse data
// 	parsedData := common.ParseJSONResponse(appointmentManagerResponse, ctx)

// 	response := builder.BuildAppointmentDTOs(parsedData)

// 	common.SendSuccess(ctx, http.StatusCreated, appointmentManagerResponse.Status, appointmentManagerResponse.StatusMessage, response)
// 	log.Println("appointment updated successfully")
// }

// func (handler *AppointmentHandler) GetPatientByAppointmentActiveStatus(ctx *gin.Context) {
// 	query := &requestData.ActiveAppointmentPatientSearchQuery{}
// 	if err := ctx.ShouldBindQuery(&query); err != nil {
// 		common.SendError(ctx, http.StatusBadRequest, 0, "Invalid query parameters", err)
// 		return
// 	}
// 	appointmentManagerResponse, err := handler.appointmentManager.GetPatientByAppointmentActiveStatus(query)
// 	// This error block will work when the store procedure catch block catches an error
// 	if common.HandleServerError(ctx, appointmentManagerResponse, err) {
// 		return // Exit if an error occurred (response is already sent)
// 	}

// 	fmt.Printf("appointments.Data type: %T\n", appointmentManagerResponse.Data)
// 	fmt.Println("appointments.Data content:", appointmentManagerResponse.Data)

// 	// Use ParseJSONResponse to parse the appointmentManagerResponse data
// 	parsedData := common.ParseJSONResponse(appointmentManagerResponse, ctx)
// 	fmt.Println("parsedData", parsedData)

// 	response := builder.BuildAppointmentDTOs(parsedData)

// 	fmt.Println("response", response)

// 	common.SendSuccess(ctx, http.StatusOK, appointmentManagerResponse.Status, appointmentManagerResponse.StatusMessage, response)

// 	log.Println("Appointments of doctor fetched successfully")
// }
