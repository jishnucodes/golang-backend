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

type PatientHandler struct {
	groupName      string
	patientManager managers.PatientManager
}

func NewPatientHandler(patientManager managers.PatientManager) *PatientHandler {
	return &PatientHandler{
		"api/patient",
		patientManager,
	}
}

func (handler *PatientHandler) RegisterApis(r *gin.Engine) {
	patientGroup := r.Group(handler.groupName)
	patientGroup.GET("/list", handler.PatientList)
	patientGroup.GET("/:patientId", handler.GetAPatient)
	patientGroup.POST("/create", handler.InsertPatient)
	patientGroup.PUT("/update/:patientId", handler.UpdatePatient)
	patientGroup.DELETE("/delete/:patientId", handler.DeleteAPatient)
}

func (handler *PatientHandler) PatientList(ctx *gin.Context) {
	patientManagerResponse, err := handler.patientManager.GetPatients()
	if common.HandleServerError(ctx, patientManagerResponse, err) {
		return
	}

	fmt.Printf("patients.Data type: %T\n", patientManagerResponse.Data)
	fmt.Println("patients.Data content:", patientManagerResponse.Data)

	parsedData := common.ParseJSONResponse(patientManagerResponse, ctx)
	response := builder.BuildPatientDTOs(parsedData)

	fmt.Println("response", response)

	common.SendSuccess(ctx, http.StatusOK, patientManagerResponse.Status, patientManagerResponse.StatusMessage, response)
	log.Println("patients fetched successfully")
}

func (handler *PatientHandler) GetAPatient(ctx *gin.Context) {
	patientData := requestData.NewPatientCreationInput()

	patientId, err := common.GetParamAsUint(ctx, "patientId")
	if err != nil {
		return
	}
	patientData.PatientID = uint(patientId)

	patientManagerResponse, err := handler.patientManager.GetAPatient(patientData)
	if common.HandleServerError(ctx, patientManagerResponse, err) {
		return
	}

	fmt.Printf("patient.Data type: %T\n", patientManagerResponse.Data)
	fmt.Println("patient.Data content:", patientManagerResponse.Data)

	parsedData := common.ParseJSONResponse(patientManagerResponse, ctx)
	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildPatientDTOs(parsedData)

	common.SendSuccess(ctx, http.StatusOK, patientManagerResponse.Status, patientManagerResponse.StatusMessage, response)
	log.Println("patient fetched successfully")
}

func (handler *PatientHandler) InsertPatient(ctx *gin.Context) {
	patientData := requestData.NewPatientCreationInput()
	// patientData.DOB = common.ParseTime(patientData.DOB)

	if err := common.BindJSONAndValidate(ctx, &patientData); err != nil {
		return
	}

	patientManagerResponse, err := handler.patientManager.CreatePatient(patientData)
	if common.HandleServerError(ctx, patientManagerResponse, err) {
		return
	}

	parsedData := common.ParseJSONResponse(patientManagerResponse, ctx)
	fmt.Println("parsedData:", parsedData)	
	response := builder.BuildPatientDTOs(parsedData)
	

	common.SendSuccess(ctx, http.StatusCreated, patientManagerResponse.Status, patientManagerResponse.StatusMessage, response)
	log.Println("patient created successfully")
}

func (handler *PatientHandler) UpdatePatient(ctx *gin.Context) {
	patientData := requestData.NewPatientCreationInput()

	if err := common.BindJSONAndValidate(ctx, &patientData); err != nil {
		return
	}

	patientId, err := common.GetParamAsUint(ctx, "patientId")
	if err != nil {
		return
	}

	patientData.PatientID = uint(patientId)

	patientManagerResponse, err := handler.patientManager.UpdatePatient(patientData)
	if common.HandleServerError(ctx, patientManagerResponse, err) {
		return
	}

	parsedData := common.ParseJSONResponse(patientManagerResponse, ctx)
	response := builder.BuildPatientDTOs(parsedData)

	common.SendSuccess(ctx, http.StatusOK, patientManagerResponse.Status, patientManagerResponse.StatusMessage, response)
	log.Println("patient updated successfully")
}

func (handler *PatientHandler) DeleteAPatient(ctx *gin.Context) {
	patientData := requestData.NewPatientCreationInput()

	patientId, err := common.GetParamAsUint(ctx, "patientId")
	if err != nil {
		return
	}

	patientData.PatientID = uint(patientId)

	patientManagerResponse, err := handler.patientManager.DeleteAPatient(patientData)
	if common.HandleServerError(ctx, patientManagerResponse, err) {
		return
	}

	fmt.Printf("patient.Data type: %T\n", patientManagerResponse.Data)
	fmt.Println("patient.Data content:", patientManagerResponse.Data)

	parsedData := common.ParseJSONResponse(patientManagerResponse, ctx)
	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildPatientDTOs(parsedData)

	common.SendSuccess(ctx, http.StatusOK, patientManagerResponse.Status, patientManagerResponse.StatusMessage, response)
	log.Println("patient deleted successfully")
}
