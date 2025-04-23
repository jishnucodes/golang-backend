package managers

import (
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"clinic-management/backend/spResponse"
	"encoding/json"
	"fmt"
)

// PatientManager interface defines the methods for patient management
type PatientManager interface {
	GetPatients() (*spResponse.Result, error)
	GetAPatient(patientData *requestData.PatientObj) (*spResponse.Result, error)
	CreatePatient(patientData *requestData.PatientObj) (*spResponse.Result, error)
	UpdatePatient(patientData *requestData.PatientObj) (*spResponse.Result, error)
	DeletePatient(patientData *requestData.PatientObj) (*spResponse.Result, error)
}

type patientManager struct {
	// Struct that implements the PatientManager interface
}

// NewPatientManager creates a new instance of PatientManager
func NewPatientManager() PatientManager {
	return &patientManager{}
}

// GetPatients retrieves all patients
func (pm *patientManager) GetPatients() (*spResponse.Result, error) {
	// Create an instance of StoredProcedureExecutor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure and capture the result
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_ListPatients", nil)
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	// Print the result data and its type for debugging
	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

// GetAPatient retrieves a specific patient
func (pm *patientManager) GetAPatient(patientData *requestData.PatientObj) (*spResponse.Result, error) {
	// Convert input to DTO
	patientDTO := builder.BuildPatientDTO(patientData)
	fmt.Println("patientDTO:", patientDTO)

	// Convert DTO to JSON
	patientJSON, err := json.Marshal(patientDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal patient data: %w", err)
	}

	fmt.Println("patientJSON", string(patientJSON))

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the patient data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_GetAPatient @PatientJSON = ?", []interface{}{string(patientJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

// CreatePatient creates a new patient
func (pm *patientManager) CreatePatient(patientData *requestData.PatientObj) (*spResponse.Result, error) {
	// Convert input to DTO
	patientDTO := builder.BuildPatientDTO(patientData)
	fmt.Println("patientDTO:", patientDTO)

	// Convert DTO to JSON
	patientJSON, err := json.Marshal(patientDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal patient data: %w", err)
	}

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the patient data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_InsertPatient @PatientJSON = ?", []interface{}{string(patientJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

// UpdatePatient updates an existing patient
func (pm *patientManager) UpdatePatient(patientData *requestData.PatientObj) (*spResponse.Result, error) {
	// Convert input to DTO
	patientDTO := builder.BuildPatientDTO(patientData)
	fmt.Println("patientDTO:", patientDTO)

	// Convert DTO to JSON
	patientJSON, err := json.Marshal(patientDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal patient data: %w", err)
	}

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the patient data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_UpdatePatient @PatientJSON = ?", []interface{}{string(patientJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

// DeletePatient deletes a patient
func (pm *patientManager) DeletePatient(patientData *requestData.PatientObj) (*spResponse.Result, error) {
	// Convert input to DTO
	patientDTO := builder.BuildPatientDTO(patientData)
	fmt.Println("patientDTO:", patientDTO)

	// Convert DTO to JSON
	patientJSON, err := json.Marshal(patientDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal patient data: %w", err)
	}

	fmt.Println("patientJSON", string(patientJSON))

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the patient data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_DeletePatient @PatientJSON = ?", []interface{}{string(patientJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}
