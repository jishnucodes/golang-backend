package managers

import (
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"clinic-management/backend/spResponse"
	"encoding/json"
	"fmt"
)

type ConsultationMedicationManager interface {
	// GetAppointmentsOfDoctorByDate(query *requestData.AppointmentSearchQuery) (*spResponse.Result, error)
	CreateConsultationMedicationData(consultationMedicationData *requestData.ConsultationMedicationObj) (*spResponse.Result, error)
	UpdateConsultationMedicationData(consultationMedicationData *requestData.ConsultationMedicationObj) (*spResponse.Result, error)
	// GetPatientByAppointmentActiveStatus(query *requestData.ActiveAppointmentPatientSearchQuery) (*spResponse.Result, error)
}

type consultationMedicationManager struct {
}

func NewConsultationMedicationManager() ConsultationMedicationManager {
	return &consultationMedicationManager{}
}

// func (am *appointmentManager) GetAppointmentsOfDoctorByDate(query *requestData.AppointmentSearchQuery) (*spResponse.Result, error) {
// 	// Create an instance of StoredProcedureExecutor
// 	spExecutor := common.NewStoredProcedureExecutor()

// 	// Convert query to JSON
// 	queryJSON, err := json.Marshal(query)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to marshal query: %w", err)
// 	}
// 	fmt.Println("queryJSON:", string(queryJSON))

// 	// Execute the stored procedure and capture the result
// 	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_ListAppointsOfDoctorByDate @AppointmentJSON = ?", []interface{}{string(queryJSON)})
// 	if err != nil {
// 		return nil, fmt.Errorf("error executing stored procedure: %w", err)
// 	}

// 	// Print the result data and its type for debugging
// 	fmt.Println("data:", data)
// 	fmt.Println("data type:", fmt.Sprintf("%T", data))

// 	return data, nil
// }

func (cm *consultationMedicationManager) CreateConsultationMedicationData(consultationMedicationData *requestData.ConsultationMedicationObj) (*spResponse.Result, error) {
	// Convert input to DTO
	consultationMedicationDTO := builder.BuildConsultationMedicationDTO(consultationMedicationData)
	fmt.Println("consultationMedicationDTO:", consultationMedicationDTO)

	// Convert DTO to JSON
	consultationMedicationJSON, err := json.Marshal(consultationMedicationDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal medication data: %w", err)
	}

	fmt.Println("consultationMedicationJSON", string(consultationMedicationJSON))

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the consultation data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_InsertConsultationMedicationDetails @ConsultationMedicationJSON = ?", []interface{}{string(consultationMedicationJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (cm *consultationMedicationManager) UpdateConsultationMedicationData(consultationMedicationData *requestData.ConsultationMedicationObj) (*spResponse.Result, error) {
	// Convert input to DTO
	consultationMedicationDTO := builder.BuildConsultationMedicationDTO(consultationMedicationData)
	fmt.Println("consultationMedicationDTO:", consultationMedicationDTO)

	// Convert DTO to JSON
	consultationMedicationJSON, err := json.Marshal(consultationMedicationDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal medication data: %w", err)
	}

	fmt.Println("consultationMedicationJSON", string(consultationMedicationJSON))

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the appointment data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_UpdateConsultationMedicationDetails @ConsultationMedicationJSON = ?", []interface{}{string(consultationMedicationJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

// func (am *appointmentManager) GetPatientByAppointmentActiveStatus(query *requestData.ActiveAppointmentPatientSearchQuery) (*spResponse.Result, error) {
// 	// Create an instance of StoredProcedureExecutor
// 	spExecutor := common.NewStoredProcedureExecutor()

// 	// Convert query to JSON
// 	queryJSON, err := json.Marshal(query)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to marshal query: %w", err)
// 	}
// 	fmt.Println("queryJSON:", string(queryJSON))

// 	// Execute the stored procedure and capture the result
// 	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_getPatientByAppointmentActiveStatus @AppointmentJSON = ?", []interface{}{string(queryJSON)})
// 	if err != nil {
// 		return nil, fmt.Errorf("error executing stored procedure: %w", err)
// 	}

// 	// Print the result data and its type for debugging
// 	fmt.Println("data:", data)
// 	fmt.Println("data type:", fmt.Sprintf("%T", data))

// 	return data, nil
// }
