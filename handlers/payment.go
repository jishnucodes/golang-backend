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

type PaymentHandler struct {
	groupName      string
	paymentManager managers.PaymentManager
}

func NewPaymentHandler(paymentManager managers.PaymentManager) *PaymentHandler {
	return &PaymentHandler{
		"api/payment",
		paymentManager,
	}
}


func (handler *PaymentHandler) RegisterApis(r *gin.Engine) {
	paymentGroup := r.Group(handler.groupName)
	paymentGroup.GET("/confirmedAppointment", handler.GetConfirmedAppointmentById)
}

func (handler *PaymentHandler) GetConfirmedAppointmentById(ctx *gin.Context) {
	query := &requestData.ConfirmedAppointmentSearchQuery{}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		common.SendError(ctx, http.StatusBadRequest, 0, "Invalid query parameters", err)
		return
	}
	paymentManagerResponse, err := handler.paymentManager.GetConfirmedAppointmentById(query)
	// This error block will work when the store procedure catch block catches an error
	if common.HandleServerError(ctx, paymentManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("payments.Data type: %T\n", paymentManagerResponse.Data)
	fmt.Println("payments.Data content:", paymentManagerResponse.Data)

	// Use ParseJSONResponse to parse the appointmentManagerResponse data
	parsedData := common.ParseJSONResponse(paymentManagerResponse, ctx)
	fmt.Println("parsedData", parsedData)

	response := builder.BuildPaymentDTOs(parsedData)

	fmt.Println("response", response)

	common.SendSuccess(ctx, http.StatusOK, paymentManagerResponse.Status, paymentManagerResponse.StatusMessage, response)

	log.Println("Confirmed payment details fetched successfully")
}