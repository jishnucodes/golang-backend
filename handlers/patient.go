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
	patientGroup.DELETE("/delete/:patientId", handler.DeletePatient)
}

func (handler *PatientHandler) PatientList(ctx *gin.Context) {
	patientManagerResponse, err := handler.patientManager.GetPatients()
	if common.HandleServerError(ctx, patientManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("patients.Data type: %T\n", patientManagerResponse.Data)
	fmt.Println("patients.Data content:", patientManagerResponse.Data)

	//Use ParseJSONResponse to parse the patientManagerResponse data
	parsedData := common.ParseJSONResponse(patientManagerResponse, ctx)

	response := builder.BuildPatientDTOs(parsedData)

	fmt.Println("response", response)

	common.SendSuccess(ctx, http.StatusOK, patientManagerResponse.Status, patientManagerResponse.StatusMessage, response)

	log.Println("patients fetched successfully")
}

func (handler *PatientHandler) GetAPatient(ctx *gin.Context) {
	// Create a new patient object
	patientData := requestData.NewPatientObj()

	patientId, err := common.GetParamAsUint(ctx, "patientId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the patient ID to the patientData object
	patientData.PatientID = uint(patientId)

	// Call the get a patient method in the patient manager
	patientManagerResponse, err := handler.patientManager.GetAPatient(patientData)
	if common.HandleServerError(ctx, patientManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("patient.Data type: %T\n", patientManagerResponse.Data)
	fmt.Println("patient.Data content:", patientManagerResponse.Data)

	//Use ParseJSONResponse to parse the patientManagerResponse data
	parsedData := common.ParseJSONResponse(patientManagerResponse, ctx)

	fmt.Println("parsedData", parsedData)

	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildPatientDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, patientManagerResponse.Status, patientManagerResponse.StatusMessage, response)
	log.Println("patient fetched successfully")
}

func (handler *PatientHandler) InsertPatient(ctx *gin.Context) {
	patientData := requestData.NewPatientObj()

	// Bind the incoming JSON to the patientData object
	if err := common.BindJSONAndValidate(ctx, &patientData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	patientManagerResponse, err := handler.patientManager.CreatePatient(patientData)

	if common.HandleServerError(ctx, patientManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	//Use ParseJSONResponse to parse the patientManagerResponse data
	parsedData := common.ParseJSONResponse(patientManagerResponse, ctx)

	response := builder.BuildPatientDTOs(parsedData)

	common.SendSuccess(ctx, http.StatusCreated, patientManagerResponse.Status, patientManagerResponse.StatusMessage, response)
	log.Println("patient created successfully")
}

func (handler *PatientHandler) UpdatePatient(ctx *gin.Context) {
	// Create a new patient object
	patientData := requestData.NewPatientObj()

	// Bind the incoming JSON to the patientData object
	if err := common.BindJSONAndValidate(ctx, &patientData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	patientId, err := common.GetParamAsUint(ctx, "patientId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the patient ID to the patientData object
	patientData.PatientID = uint(patientId)

	// Call the update method in the patient manager
	patientManagerResponse, err := handler.patientManager.UpdatePatient(patientData)

	if common.HandleServerError(ctx, patientManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	//Use ParseJSONResponse to parse the patientManagerResponse data
	parsedData := common.ParseJSONResponse(patientManagerResponse, ctx)

	response := builder.BuildPatientDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, patientManagerResponse.Status, patientManagerResponse.StatusMessage, response)
	log.Println("patient updated successfully")
}

func (handler *PatientHandler) DeletePatient(ctx *gin.Context) {
	// Create a new patient object
	patientData := requestData.NewPatientObj()

	patientId, err := common.GetParamAsUint(ctx, "patientId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the patient ID to the patientData object
	patientData.PatientID = uint(patientId)

	// Call the delete method in the patient manager
	patientManagerResponse, err := handler.patientManager.DeletePatient(patientData)

	if common.HandleServerError(ctx, patientManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("patient.Data type: %T\n", patientManagerResponse.Data)
	fmt.Println("patient.Data content:", patientManagerResponse.Data)

	//Use ParseJSONResponse to parse the patientManagerResponse data
	parsedData := common.ParseJSONResponse(patientManagerResponse, ctx)
	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildPatientDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, patientManagerResponse.Status, patientManagerResponse.StatusMessage, response)
	log.Println("patient deleted successfully")
}
