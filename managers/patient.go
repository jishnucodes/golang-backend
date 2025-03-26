package managers

import (
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"clinic-management/backend/models"
	"encoding/json"
	"fmt"
)

// this is the interface which have the methods. Methods are the functions which is
// also a method of userStruct
type PatientManager interface {
	CreatePatient(patientData *common.PatientCreationInput) (*returnPatientData, error) 
	// CreateUser(userData *common.UserCreationInput) (*returnData, error)
}

type returnPatientData struct {
    Patients   []models.CMSPatients `json:"result,omitempty"`   // Slice for user records
    Message string           `json:"message,omitempty"` // Message for string results
}

type patientManager struct {
	//this is a struct which defines the methods
}

func NewPatientManager() PatientManager {
	return &patientManager{}
}



func (pm *patientManager) CreatePatient(patientData *common.PatientCreationInput) (*returnPatientData, error) {
    // Convert input to DTO
    patientDTO := builder.BuildPatientDTO(patientData)
    fmt.Println("userDTO:", patientDTO)

    // Convert DTO to JSON
    patientJSON, err := json.Marshal(patientDTO)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal user data: %w", err)
    }

    // Create an instance of the stored procedure executor
    spExecutor := common.NewStoredProcedureExecutor()

	var patient models.CMSPatients

    // Execute the stored procedure with the user data
    data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_InsertPatient @PatientJSON = ?", []interface{}{string(patientJSON)}, &patient)
    if err != nil {

        return nil, fmt.Errorf("error executing stored procedure: %w", err)
    }

    // Prepare the result object
    result := &returnPatientData{}
    if data != "" {
        result.Message = data
    } else {
        result.Message = "User created successfully"
    }

    return result, nil
}





