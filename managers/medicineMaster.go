package managers

import (
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"clinic-management/backend/spResponse"
	"encoding/json"
	"fmt"
)

// DepartmentManager interface defines the methods for Medicine Master management
type MedicineMasterManager interface {
	GetMedicines() (*spResponse.Result, error)
	GetAMedicine(medicineMasterData *requestData.MedicineMasterObj) (*spResponse.Result, error)
	CreateMedicine(medicineMasterData *requestData.MedicineMasterObj) (*spResponse.Result, error)
	UpdateMedicine(medicineMasterData *requestData.MedicineMasterObj) (*spResponse.Result, error)
	SoftDeleteMedicine(medicineMasterData *requestData.MedicineMasterObj) (*spResponse.Result, error)
	DeleteMedicine(medicineMasterData *requestData.MedicineMasterObj) (*spResponse.Result, error)
}

type medicineMasterManager struct {
	// Struct that implements the medicineMaster interface
}

func NewMedicineMasterManager() MedicineMasterManager {
	return &medicineMasterManager{}
}

func (dm *medicineMasterManager) GetMedicines() (*spResponse.Result, error) {
	// Create an instance of StoredProcedureExecutor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure and capture the result
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_ListMedicines", nil)
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	// Print the result data and its type for debugging
	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (dm *medicineMasterManager) GetAMedicine(medicineMasterData *requestData.MedicineMasterObj) (*spResponse.Result, error) {
	// Convert input to DTO
	medicineMasterDTO := builder.BuildMedicineMasterObj(medicineMasterData)
	fmt.Println("medicineMasterDTO:", medicineMasterDTO)

	// Convert DTO to JSON
	medicineMasterJSON, err := json.Marshal(medicineMasterDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal medicine data: %w", err)
	}

	fmt.Println("medicineMasterJSON", string(medicineMasterJSON))

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the department data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_GetAMedicine @MedicineJSON = ?", []interface{}{string(medicineMasterJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (dm *medicineMasterManager) CreateMedicine(medicineMasterData *requestData.MedicineMasterObj) (*spResponse.Result, error) {
	// Convert input to DTO
	medicineMasterDTO := builder.BuildMedicineMasterObj(medicineMasterData)
	fmt.Println("medicineMasterDTO:", medicineMasterDTO)

	// Convert DTO to JSON
	medicineMasterJSON, err := json.Marshal(medicineMasterDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal medicine data: %w", err)
	}
    fmt.Println("medicineJSON",string(medicineMasterJSON))
	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the medicine data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_InsertMedicine @MedicineJSON = ?", []interface{}{string(medicineMasterJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (dm *medicineMasterManager) UpdateMedicine(medicineMasterData *requestData.MedicineMasterObj) (*spResponse.Result, error) {
	// Convert input to DTO
	medicineMasterDTO := builder.BuildMedicineMasterObj(medicineMasterData)
	fmt.Println("medicineMasterDTO:", medicineMasterDTO)

	// Convert DTO to JSON
	medicineMasterJSON, err := json.Marshal(medicineMasterDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal medicine data: %w", err)
	}

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the medicine data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_UpdateMedicine @MedicineJSON = ?", []interface{}{string(medicineMasterJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (dm *medicineMasterManager) SoftDeleteMedicine(medicineMasterData *requestData.MedicineMasterObj) (*spResponse.Result, error) {
	// Convert input to DTO
	medicineMasterDTO := builder.BuildMedicineMasterObj(medicineMasterData)
	fmt.Println("medicineMasterDTO:", medicineMasterDTO)

	// Convert DTO to JSON
	medicineMasterJSON, err := json.Marshal(medicineMasterDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal medicine data: %w", err)
	}

	fmt.Println("MedicineMasterJSON", string(medicineMasterJSON))

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the medicine data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_SoftDeleteMedicine @MedicineJSON = ?", []interface{}{string(medicineMasterJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (dm *medicineMasterManager) DeleteMedicine(medicineMasterData *requestData.MedicineMasterObj) (*spResponse.Result, error) {
	// Convert input to DTO
	medicineMasterDTO := builder.BuildMedicineMasterObj(medicineMasterData)
	fmt.Println("medicineMasterDTO:", medicineMasterDTO)

	// Convert DTO to JSON
	medicineMasterJSON, err := json.Marshal(medicineMasterDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal medicine data: %w", err)
	}

	fmt.Println("MedicineMasterJSON", string(medicineMasterJSON))

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the medicine data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_DeleteMedicine @MedicineJSON = ?", []interface{}{string(medicineMasterJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}
