package managers

import (
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"clinic-management/backend/spResponse"
	"encoding/json"
	"fmt"
)

// DoctorManager interface defines the methods for doctor operations
type DoctorManager interface {
	GetDoctors() (*spResponse.Result, error)
	GetADoctor(doctorData *requestData.DoctorObj) (*spResponse.Result, error)
	CreateDoctor(doctorData *requestData.DoctorObj) (*spResponse.Result, error)
	UpdateDoctor(doctorData *requestData.DoctorObj) (*spResponse.Result, error)
	DeleteADoctor(doctorData *requestData.DoctorObj) (*spResponse.Result, error)
}

type doctorManager struct {
	// Struct that implements the DoctorManager interface
}

func NewDoctorManager() DoctorManager {
	return &doctorManager{}
}

func (dm *doctorManager) GetDoctors() (*spResponse.Result, error) {
	// Create an instance of StoredProcedureExecutor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure and capture the result
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_ListDoctors", nil)
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	// Print the result data and its type for debugging
	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (dm *doctorManager) GetADoctor(doctorData *requestData.DoctorObj) (*spResponse.Result, error) {
	// Convert input to DTO
	doctorDTO := builder.BuildDoctorDTO(doctorData)
	fmt.Println("doctorDTO:", doctorDTO)

	// Convert DTO to JSON
	doctorJSON, err := json.Marshal(doctorDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal doctor data: %w", err)
	}

	fmt.Println("doctorJSON", string(doctorJSON))

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the doctor data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_GetADoctorDetails @DoctorJSON = ?", []interface{}{string(doctorJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (dm *doctorManager) CreateDoctor(doctorData *requestData.DoctorObj) (*spResponse.Result, error) {
	// Convert input to DTO
	doctorDTO := builder.BuildDoctorDTO(doctorData)
	fmt.Println("doctorDTO:", doctorDTO)

	// Convert DTO to JSON
	doctorJSON, err := json.Marshal(doctorDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal doctor data: %w", err)
	}

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the doctor data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_InsertDoctorDetails @DoctorJSON = ?", []interface{}{string(doctorJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (dm *doctorManager) UpdateDoctor(doctorData *requestData.DoctorObj) (*spResponse.Result, error) {
	// Convert input to DTO
	doctorDTO := builder.BuildDoctorDTO(doctorData)
	fmt.Println("doctorDTO:", doctorDTO)

	// Convert DTO to JSON
	doctorJSON, err := json.Marshal(doctorDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal doctor data: %w", err)
	}

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the doctor data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_UpdateDoctorDetails @DoctorJSON = ?", []interface{}{string(doctorJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (dm *doctorManager) DeleteADoctor(doctorData *requestData.DoctorObj) (*spResponse.Result, error) {
	// Convert input to DTO
	doctorDTO := builder.BuildDoctorDTO(doctorData)
	fmt.Println("doctorDTO:", doctorDTO)

	// Convert DTO to JSON
	doctorJSON, err := json.Marshal(doctorDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal doctor data: %w", err)
	}

	fmt.Println("doctorJSON", string(doctorJSON))

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the doctor data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_DeleteDoctor @DoctorJSON = ?", []interface{}{string(doctorJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}
