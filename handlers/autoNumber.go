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

type AutoNumberHandler struct {
	groupName   string
	autoNumberManager managers.AutoNumberManager
}

func NewAutoNumberHandler(autoNumberManager managers.AutoNumberManager) *AutoNumberHandler {
	return &AutoNumberHandler{
		"api/autoNumber",
		autoNumberManager,
	}
}

func (handler *AutoNumberHandler) RegisterApis(r *gin.Engine) {
	autoNumberGroup := r.Group(handler.groupName)
	// autoNumberGroup.GET("/list", handler.RoleList)
	// autoNumberGroup.GET("/:roleId", handler.GetARole)
	// autoNumberGroup.POST("/create", handler.InsertRole)
	autoNumberGroup.POST("/generate", handler.GenerateAutoNumber)
	
}


func (handler *AutoNumberHandler) GenerateAutoNumber(ctx *gin.Context) {
	autoNumberData := requestData.NewAutoNumberObj()

	// Bind the incoming JSON to the roleData object
	if err := common.BindJSONAndValidate(ctx, &autoNumberData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	autoNumberManagerResponse, err := handler.autoNumberManager.GenerateAutoNumber(autoNumberData)

	if common.HandleServerError(ctx, autoNumberManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	//Use ParseJSONResponse to parse the roleManagerResponse data
	parsedData := common.ParseJSONResponse(autoNumberManagerResponse, ctx)

	fmt.Println("parsedData:", parsedData)

	response := builder.BuildAutoNumberDTOs(parsedData)

	common.SendSuccess(ctx, http.StatusCreated, autoNumberManagerResponse.Status, autoNumberManagerResponse.StatusMessage, response)
	log.Println("autoNumber created successfully")
}

