package handlers

import (
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"clinic-management/backend/managers"
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
	appointmentGroup.POST("/create", handler.InsertAppointment)
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