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

// DoctorAvailabilityHandler handles doctor availability-related API requests
type DoctorAvailabilityHandler struct {
	groupName                 string
	doctorAvailabilityManager managers.DoctorAvailabilityManager
}

// NewDoctorAvailabilityHandler creates a new instance of DoctorAvailabilityHandler
func NewDoctorAvailabilityHandler(doctorAvailabilityManager managers.DoctorAvailabilityManager) *DoctorAvailabilityHandler {
	return &DoctorAvailabilityHandler{
		groupName:                 "api/doctorAvailability",
		doctorAvailabilityManager: doctorAvailabilityManager,
	}
}

// RegisterApis registers all doctor availability-related API routes
func (h *DoctorAvailabilityHandler) RegisterApis(router *gin.Engine) {
	doctorAvailabilityGroup := router.Group(h.groupName)

	doctorAvailabilityGroup.GET("/list", h.DoctorAvailabilityList)
	doctorAvailabilityGroup.GET("/:doctorId/:availabilityId", h.GetADoctorAvailability)
	doctorAvailabilityGroup.POST("/insert", h.InsertDoctorAvailability)
	doctorAvailabilityGroup.PUT("/update/:availabilityId", h.UpdateDoctorAvailability)
	doctorAvailabilityGroup.DELETE("/:availabilityId", h.DeleteADoctorAvailability)

}

// DoctorAvailabilityList handles the GET request to list all doctor availabilities
func (h *DoctorAvailabilityHandler) DoctorAvailabilityList(ctx *gin.Context) {
	doctorAvailabilityManagerResponse, err := h.doctorAvailabilityManager.GetDoctorAvailabilities()
	if common.HandleServerError(ctx, doctorAvailabilityManagerResponse, err) {
		return
	}
	fmt.Printf("doctors.Data type: %T\n", doctorAvailabilityManagerResponse.Data)
	fmt.Println("doctors.Data content:", doctorAvailabilityManagerResponse.Data)

	// Use ParseJSONResponse to parse the doctorAvailabilityManagerResponse data
	parsedData := common.ParseJSONResponse(doctorAvailabilityManagerResponse, ctx)

	response := builder.BuildDoctorAvailabilityDTOs(parsedData)

	fmt.Println("response", response)

	common.SendSuccess(ctx, http.StatusOK, doctorAvailabilityManagerResponse.Status, doctorAvailabilityManagerResponse.StatusMessage, response)
	log.Println("doctors availability fetched successfully")
}

// GetADoctorAvailability handles the GET request to retrieve a specific doctor availability
func (h *DoctorAvailabilityHandler) GetADoctorAvailability(c *gin.Context) {

	// Create a new doctor availability object
	availabilityData := requestData.NewDoctorAvailabilityObj()

	doctorId, err := common.GetParamAsUint(c, "doctorId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	availabilityId, err := common.GetParamAsUint(c, "availabilityId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the doctor ID and availability ID to the doctor availability object
	availabilityData.DoctorID = uint(doctorId)
	availabilityData.AvailabilityID = uint(availabilityId)

	// Call the get a doctor availability method in the doctor availability manager
	doctorAvailabilityManagerResponse, err := h.doctorAvailabilityManager.GetADoctorAvailability(availabilityData)
	if common.HandleServerError(c, doctorAvailabilityManagerResponse, err) {
		return
	}

	// Use ParseJSONResponse to parse the doctorAvailabilityManagerResponse data
	parsedData := common.ParseJSONResponse(doctorAvailabilityManagerResponse, c)

	response := builder.BuildDoctorAvailabilityDTOs(parsedData)

	fmt.Println("response", response)

	common.SendSuccess(c, http.StatusOK, doctorAvailabilityManagerResponse.Status, doctorAvailabilityManagerResponse.StatusMessage, response)
	log.Println("doctor availability fetched successfully")

}

// InsertDoctorAvailability handles the POST request to create a new doctor availability
func (h *DoctorAvailabilityHandler) InsertDoctorAvailability(ctx *gin.Context) {

	doctorAvailabilityData := requestData.NewDoctorAvailabilityObj()

	// Bind the incoming JSON to the doctorAvailabilityData object
	if err := common.BindJSONAndValidate(ctx, &doctorAvailabilityData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	doctorAvailabilityData.AvailableTimeStart = common.ParseTime(doctorAvailabilityData.AvailableTimeStart)
	doctorAvailabilityData.AvailableTimeEnd = common.ParseTime(doctorAvailabilityData.AvailableTimeEnd)

	doctorAvailabilityManagerResponse, err := h.doctorAvailabilityManager.CreateDoctorAvailability(doctorAvailabilityData)
	if common.HandleServerError(ctx, doctorAvailabilityManagerResponse, err) {
		return
	}

	// Use ParseJSONResponse to parse the doctorAvailabilityManagerResponse data
	parsedData := common.ParseJSONResponse(doctorAvailabilityManagerResponse, ctx)

	response := builder.BuildDoctorAvailabilityDTOs(parsedData)

	fmt.Println("response", response)

	common.SendSuccess(ctx, http.StatusOK, doctorAvailabilityManagerResponse.Status, doctorAvailabilityManagerResponse.StatusMessage, response)
	log.Println("doctor availability created successfully")
}

// UpdateDoctorAvailability handles the PUT request to update an existing doctor availability
func (h *DoctorAvailabilityHandler) UpdateDoctorAvailability(ctx *gin.Context) {

	doctorAvailabilityData := requestData.NewDoctorAvailabilityObj()

	// Bind the incoming JSON to the doctorAvailabilityData object
	if err := common.BindJSONAndValidate(ctx, &doctorAvailabilityData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	availabilityId, err := common.GetParamAsUint(ctx, "availabilityId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the availability ID to the doctorAvailabilityData object
	doctorAvailabilityData.AvailabilityID = uint(availabilityId)

	// Call the update doctor availability method in the doctor availability manager
	doctorAvailabilityManagerResponse, err := h.doctorAvailabilityManager.UpdateDoctorAvailability(doctorAvailabilityData)
	if common.HandleServerError(ctx, doctorAvailabilityManagerResponse, err) {
		return
	}

	// Use ParseJSONResponse to parse the doctorAvailabilityManagerResponse data
	parsedData := common.ParseJSONResponse(doctorAvailabilityManagerResponse, ctx)

	response := builder.BuildDoctorAvailabilityDTOs(parsedData)

	fmt.Println("response", response)

	common.SendSuccess(ctx, http.StatusOK, doctorAvailabilityManagerResponse.Status, doctorAvailabilityManagerResponse.StatusMessage, response)

}

// DeleteADoctorAvailability handles the DELETE request to remove a doctor availability
func (h *DoctorAvailabilityHandler) DeleteADoctorAvailability(ctx *gin.Context) {

	// Create a new doctor availability object
	availabilityData := requestData.NewDoctorAvailabilityObj()

	availabilityId, err := common.GetParamAsUint(ctx, "availabilityId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the availability ID to the doctorAvailabilityData object
	availabilityData.AvailabilityID = uint(availabilityId)

	// Call the delete doctor availability method in the doctor availability manager
	doctorAvailabilityManagerResponse, err := h.doctorAvailabilityManager.DeleteADoctorAvailability(availabilityData)
	if common.HandleServerError(ctx, doctorAvailabilityManagerResponse, err) {
		return
	}

	// Use ParseJSONResponse to parse the doctorAvailabilityManagerResponse data
	parsedData := common.ParseJSONResponse(doctorAvailabilityManagerResponse, ctx)

	response := builder.BuildDoctorAvailabilityDTOs(parsedData)

	fmt.Println("response", response)

	common.SendSuccess(ctx, http.StatusOK, doctorAvailabilityManagerResponse.Status, doctorAvailabilityManagerResponse.StatusMessage, response)
	log.Println("doctor availability deleted successfully")
}
