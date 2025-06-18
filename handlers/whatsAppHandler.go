package handlers

import (
	"clinic-management/backend/common"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type WhatsAppHandler struct {
	groupName string
	// departmentManager managers.DepartmentManager
}

func NewWhatsAppHandler() *WhatsAppHandler {
	return &WhatsAppHandler{
		"api/whatsAppChat",
	}
}


func (handler *WhatsAppHandler) RegisterApis(r *gin.Engine) {
	chatGroup := r.Group(handler.groupName)
	chatGroup.POST("/addChat", handler.whatsAppChat)
}

type InputType struct {
	Message string   `json:"message"`
}

func (handler *WhatsAppHandler) whatsAppChat(ctx *gin.Context) {
	// departmentManagerResponse, err := handler.departmentManager.GetDepartments()
	// if common.HandleServerError(ctx, departmentManagerResponse, err) {
	// 	return // Exit if an error occurred (response is already sent)
	// }

	// fmt.Printf("departments.Data type: %T\n", departmentManagerResponse.Data)
	// fmt.Println("departments.Data content:", departmentManagerResponse.Data)

	// //Use ParseJSONResponse to parse the departmentManagerResponse data
	// parsedData := common.ParseJSONResponse(departmentManagerResponse, ctx)

	// fmt.Println("parsedData", parsedData)

	// response := builder.BuildDepartmentDTOs(parsedData)

	// fmt.Println("response", response)

	// common.SendSuccess(ctx, http.StatusOK, departmentManagerResponse.Status, departmentManagerResponse.StatusMessage, response)

	var messageData InputType

	// Bind the incoming JSON to the departmentData object
	if err := common.BindJSONAndValidate(ctx, &messageData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}
	fmt.Println("chat", messageData)
	log.Println("message hitted successfully", messageData)
}