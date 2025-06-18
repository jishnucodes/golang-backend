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

type MedicineMasterHandler struct {
	groupName         string
	medicineMasterManager managers.MedicineMasterManager
}
func NewMedicineMasterHandler(medicineMasterManager managers.MedicineMasterManager) *MedicineMasterHandler {
	return &MedicineMasterHandler{
		"api/medicine-master",
		medicineMasterManager,
	}
}

func (handler *MedicineMasterHandler) RegisterApis(r *gin.Engine) {
	medicineMasterGroup := r.Group(handler.groupName)
	medicineMasterGroup.GET("/list", handler.MedicineList)
	medicineMasterGroup.GET("/:medicineId", handler.GetAMedicine)
	medicineMasterGroup.POST("/create", handler.InsertMedicine)
	medicineMasterGroup.PUT("/update/:medicineId", handler.UpdateMedicine)
	medicineMasterGroup.DELETE("/delete/:medicineId", handler.DeleteMedicine)
	medicineMasterGroup.PUT("/soft-delete/:medicineId", handler.SoftDeleteMedicine) 


	//soft delete is not implemented yet
}

func (handler *MedicineMasterHandler) MedicineList(ctx *gin.Context) {
	medicineManagerResponse, err := handler.medicineMasterManager.GetMedicines()
	if common.HandleServerError(ctx, medicineManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("medicines.Data type: %T\n", medicineManagerResponse.Data)
	fmt.Println("medicines.Data content:", medicineManagerResponse.Data)

	// Use ParseJSONResponse to parse the medicineManagerResponse data
	parsedData := common.ParseJSONResponse(medicineManagerResponse, ctx)

	fmt.Println("parsedData", parsedData)

	response := builder.BuildMedicineMasterDTOs(parsedData)

	fmt.Println("response", response)

	common.SendSuccess(ctx, http.StatusOK, medicineManagerResponse.Status, medicineManagerResponse.StatusMessage, response)

	log.Println("medicines fetched successfully")
}

func (handler *MedicineMasterHandler) InsertMedicine(ctx *gin.Context) {
	medicineData := requestData.NewMedicineMasterObj()

	// Bind the incoming JSON to the medicineData object
	if err := common.BindJSONAndValidate(ctx, &medicineData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}

	medicineManagerResponse, err := handler.medicineMasterManager.CreateMedicine(medicineData)

	if common.HandleServerError(ctx, medicineManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	// Use ParseJSONResponse to parse the medicineManagerResponse data
	parsedData := common.ParseJSONResponse(medicineManagerResponse, ctx)

	response := builder.BuildMedicineMasterDTOs(parsedData)

	common.SendSuccess(ctx, http.StatusCreated, medicineManagerResponse.Status, medicineManagerResponse.StatusMessage, response)
	log.Println("medicine created successfully")
}
 func (handler *MedicineMasterHandler) UpdateMedicine(ctx *gin.Context) {
	// Create a new medicine object
	medicineData := requestData.NewMedicineMasterObj()
	// Bind the incoming JSON to the medicineData object
	if err := common.BindJSONAndValidate(ctx, &medicineData); err != nil {
		return // Error response is already handled in BindJSONAndValidate
	}
	medicineId, err := common.GetParamAsUint(ctx, "medicineId")
	if err != nil {
		return // The function already sends an error response, so just return
	}
	// Assign the medicine ID to the medicineData object
	medicineData.MedicineID = uint(medicineId)
	// Call the update method in the medicine manager
	medicineManagerResponse, err := handler.medicineMasterManager.UpdateMedicine(medicineData)
	if common.HandleServerError(ctx, medicineManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}
	// Use ParseJSONResponse to parse the medicineManagerResponse data
	parsedData := common.ParseJSONResponse(medicineManagerResponse, ctx)
	response := builder.BuildMedicineMasterDTOs(parsedData)
	// Send success response
	common.SendSuccess(ctx, http.StatusOK, medicineManagerResponse.Status, medicineManagerResponse.StatusMessage, response)
	log.Println("medicine updated successfully")
}

 func (handler *MedicineMasterHandler) DeleteMedicine(ctx *gin.Context) {
	// Create a new medicine object
	medicineData := requestData.NewMedicineMasterObj()
	medicineId, err := common.GetParamAsUint(ctx, "medicineId")
	if err != nil {
		return // The function already sends an error response, so just return
	}
	// Assign the medicine ID to the medicineData object
	medicineData.MedicineID = uint(medicineId)
	// Call the delete method in the medicine manager
	medicineManagerResponse, err := handler.medicineMasterManager.DeleteMedicine(medicineData)
	if common.HandleServerError(ctx, medicineManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}
	fmt.Printf("medicine.Data type: %T\n", medicineManagerResponse.Data)
	fmt.Println("medicine.Data content:", medicineManagerResponse.Data)
	// Use ParseJSONResponse to parse the medicineManagerResponse data
	parsedData := common.ParseJSONResponse(medicineManagerResponse, ctx)
	fmt.Printf("Parsed data type: %T\n", parsedData)
	response := builder.BuildMedicineMasterDTOs(parsedData)
	// Send success response
	common.SendSuccess(ctx, http.StatusOK, medicineManagerResponse.Status, medicineManagerResponse.StatusMessage, response)
	log.Println("medicine deleted successfully")	
}

func (handler *MedicineMasterHandler) SoftDeleteMedicine(ctx *gin.Context) {
	// Create a new medicine object	
	medicineData := requestData.NewMedicineMasterObj()
	medicineId, err := common.GetParamAsUint(ctx, "medicineId")
	if err != nil {
		return // The function already sends an error response, so just return
	}
	// Assign the medicine ID to the medicineData object
	medicineData.MedicineID = uint(medicineId)
	// Call the soft delete method in the medicine manager
	medicineManagerResponse, err := handler.medicineMasterManager.SoftDeleteMedicine(medicineData)
	if common.HandleServerError(ctx, medicineManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}	
	fmt.Printf("medicine.Data type: %T\n", medicineManagerResponse.Data)
	fmt.Println("medicine.Data content:", medicineManagerResponse.Data)
	// Use ParseJSONResponse to parse the medicineManagerResponse data
	parsedData := common.ParseJSONResponse(medicineManagerResponse, ctx)
	fmt.Printf("Parsed data type: %T\n", parsedData)
	response := builder.BuildMedicineMasterDTOs(parsedData)
	// Send success response
	common.SendSuccess(ctx, http.StatusOK, medicineManagerResponse.Status, medicineManagerResponse.StatusMessage, response)
	log.Println("medicine soft deleted successfully")
}
 
func (handler *MedicineMasterHandler) GetAMedicine(ctx *gin.Context) {
	// Create a new medicine object
	medicineData := requestData.NewMedicineMasterObj()

	medicineId, err := common.GetParamAsUint(ctx, "medicineId")
	if err != nil {
		return // The function already sends an error response, so just return
	}

	// Assign the medicine ID to the medicineData object
	medicineData.MedicineID = uint(medicineId)

	// Call the get a medicine method in the medicine manager
	medicineManagerResponse, err := handler.medicineMasterManager.GetAMedicine(medicineData)
	if common.HandleServerError(ctx, medicineManagerResponse, err) {
		return // Exit if an error occurred (response is already sent)
	}

	fmt.Printf("medicine.Data type: %T\n", medicineManagerResponse.Data)
	fmt.Println("medicine.Data content:", medicineManagerResponse.Data)

	// Use ParseJSONResponse to parse the medicineManagerResponse data
	parsedData := common.ParseJSONResponse(medicineManagerResponse, ctx)

	fmt.Println("parsedData", parsedData)

	fmt.Printf("Parsed data type: %T\n", parsedData)

	response := builder.BuildMedicineMasterDTOs(parsedData)

	// Send success response
	common.SendSuccess(ctx, http.StatusOK, medicineManagerResponse.Status, medicineManagerResponse.StatusMessage, response)
	log.Println("medicine fetched successfully")
}


