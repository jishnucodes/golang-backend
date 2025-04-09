package managers

import (
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"clinic-management/backend/spResponse"
	"encoding/json"
	"fmt"
)

// PatientManager interface defines the methods for patient management
type PatientManager interface {
	GetPatients() (*spResponse.Result, error)
	GetAPatient(patientData *common.PatientCreationInput) (*spResponse.Result, error)
	CreatePatient(patientData *common.PatientCreationInput) (*spResponse.Result, error)
	UpdatePatient(patientData *common.PatientCreationInput) (*spResponse.Result, error)
	DeleteAPatient(patientData *common.PatientCreationInput) (*spResponse.Result, error)
}

type patientManager struct {
	// struct that implements the PatientManager interface
}

// NewPatientManager creates a new instance of PatientManager
func NewPatientManager() PatientManager {
	return &patientManager{}
}

// GetPatients retrieves all patients
func (pm *patientManager) GetPatients() (*spResponse.Result, error) {
	spExecutor := common.NewStoredProcedureExecutor()

	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_ListPatients", nil)
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

// GetAPatient retrieves a specific patient
func (pm *patientManager) GetAPatient(patientData *common.PatientCreationInput) (*spResponse.Result, error) {
	patientDTO := builder.BuildPatientDTO(patientData)
	fmt.Println("patientDTO:", patientDTO)

	patientJSON, err := json.Marshal(patientDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal patient data: %w", err)
	}

	spExecutor := common.NewStoredProcedureExecutor()

	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_GetAPatient @PatientJSON = ?", []interface{}{string(patientJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

// CreatePatient creates a new patient
func (pm *patientManager) CreatePatient(patientData *common.PatientCreationInput) (*spResponse.Result, error) {
	patientDTO := builder.BuildPatientDTO(patientData)
	fmt.Println("patientDTO:", patientDTO)

	patientJSON, err := json.Marshal(patientDTO)
	fmt.Println("patientJSON:", string(patientJSON))	
	if err != nil {
		return nil, fmt.Errorf("failed to marshal patient data: %w", err)
	}

	spExecutor := common.NewStoredProcedureExecutor()

	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_InsertPatient @PatientJSON = ?", []interface{}{string(patientJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

// UpdatePatient updates an existing patient
func (pm *patientManager) UpdatePatient(patientData *common.PatientCreationInput) (*spResponse.Result, error) {
	patientDTO := builder.BuildPatientDTO(patientData)
	fmt.Println("patientDTO:", patientDTO)

	patientJSON, err := json.Marshal(patientDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal patient data: %w", err)
	}

	spExecutor := common.NewStoredProcedureExecutor()

	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_UpdatePatient @PatientJSON = ?", []interface{}{string(patientJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

// DeleteAPatient deletes a patient
func (pm *patientManager) DeleteAPatient(patientData *common.PatientCreationInput) (*spResponse.Result, error) {
	patientDTO := builder.BuildPatientDTO(patientData)
	fmt.Println("patientDTO:", patientDTO)

	patientJSON, err := json.Marshal(patientDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal patient data: %w", err)
	}

	spExecutor := common.NewStoredProcedureExecutor()

	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_DeletePatient @PatientJSON = ?", []interface{}{string(patientJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}
