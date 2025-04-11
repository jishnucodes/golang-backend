package managers

import (
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"clinic-management/backend/spResponse"
	"encoding/json"
	"fmt"
)

// DoctorAvailabilityManager interface defines the methods for doctor availability operations
type DoctorAvailabilityManager interface {
	GetDoctorAvailabilities() (*spResponse.Result, error)
	GetADoctorAvailability(availabilityData *requestData.DoctorAvailabilityObj) (*spResponse.Result, error)
	CreateDoctorAvailability(availabilityData *requestData.DoctorAvailabilityObj) (*spResponse.Result, error)
	UpdateDoctorAvailability(availabilityData *requestData.DoctorAvailabilityObj) (*spResponse.Result, error)
	DeleteADoctorAvailability(availabilityData *requestData.DoctorAvailabilityObj) (*spResponse.Result, error)
}

type doctorAvailabilityManager struct {
	// Struct that implements the DoctorAvailabilityManager interface
}

func NewDoctorAvailabilityManager() DoctorAvailabilityManager {
	return &doctorAvailabilityManager{}
}

func (dam *doctorAvailabilityManager) GetDoctorAvailabilities() (*spResponse.Result, error) {
	// Create an instance of StoredProcedureExecutor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure and capture the result
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_ListDoctorsAvailability", nil)
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	// Print the result data and its type for debugging
	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (dam *doctorAvailabilityManager) GetADoctorAvailability(availabilityData *requestData.DoctorAvailabilityObj) (*spResponse.Result, error) {
	// Convert input to DTO
	availabilityDTO := builder.BuildDoctorAvailabilityDTO(availabilityData)
	fmt.Println("availabilityDTO:", availabilityDTO)

	// Convert DTO to JSON
	availabilityJSON, err := json.Marshal(availabilityDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal availability data: %w", err)
	}

	fmt.Println("availabilityJSON", string(availabilityJSON))

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the availability data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_GetADoctorAvailabilityDetails @DoctorAvailabilityJSON = ?", []interface{}{string(availabilityJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (dam *doctorAvailabilityManager) CreateDoctorAvailability(availabilityData *requestData.DoctorAvailabilityObj) (*spResponse.Result, error) {
	// Convert input to DTO
	fmt.Println("availabilityData:", availabilityData)
	availabilityDTO := builder.BuildDoctorAvailabilityDTO(availabilityData)
	fmt.Println("availabilityDTO:", availabilityDTO)

	// Convert DTO to JSON
	availabilityJSON, err := json.Marshal(availabilityDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal availability data: %w", err)
	}

	fmt.Println("availabilityJSON", string(availabilityJSON))

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the availability data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_InsertDoctorAvailabilityDetails @DoctorAvailabilityJSON = ?", []interface{}{string(availabilityJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (dam *doctorAvailabilityManager) UpdateDoctorAvailability(availabilityData *requestData.DoctorAvailabilityObj) (*spResponse.Result, error) {
	// Convert input to DTO
	availabilityDTO := builder.BuildDoctorAvailabilityDTO(availabilityData)
	fmt.Println("availabilityDTO:", availabilityDTO)

	// Convert DTO to JSON
	availabilityJSON, err := json.Marshal(availabilityDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal availability data: %w", err)
	}

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the availability data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_UpdateDoctorAvailabilityDetails @DoctorAvailabilityJSON = ?", []interface{}{string(availabilityJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (dam *doctorAvailabilityManager) DeleteADoctorAvailability(availabilityData *requestData.DoctorAvailabilityObj) (*spResponse.Result, error) {
	// Convert input to DTO
	availabilityDTO := builder.BuildDoctorAvailabilityDTO(availabilityData)
	fmt.Println("availabilityDTO:", availabilityDTO)

	// Convert DTO to JSON
	availabilityJSON, err := json.Marshal(availabilityDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal availability data: %w", err)
	}

	fmt.Println("availabilityJSON", string(availabilityJSON))

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the availability data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_DeleteDoctorAvailabilityDetails @DoctorAvailabilityJSON = ?", []interface{}{string(availabilityJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}
