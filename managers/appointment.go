package managers

import (
	"clinic-management/backend/builder"
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"clinic-management/backend/spResponse"
	"encoding/json"
	"fmt"
)

type AppointmentManager interface {
	GetAppointmentsOfDoctorByDate(query *requestData.AppointmentSearchQuery) (*spResponse.Result, error) 
	CreateAppointment(appointmentData *requestData.AppointmentObj) (*spResponse.Result, error)
}

type appointmentManager struct {
}

func NewAppointmentManager() AppointmentManager {
	return &appointmentManager{}
}

func (am *appointmentManager) GetAppointmentsOfDoctorByDate(query *requestData.AppointmentSearchQuery) (*spResponse.Result, error) {
	// Create an instance of StoredProcedureExecutor
	spExecutor := common.NewStoredProcedureExecutor()

	// Convert query to JSON
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal query: %w", err)
	}
	fmt.Println("queryJSON:", string(queryJSON))

	// Execute the stored procedure and capture the result
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_ListAppointsOfDoctorByDate @AppointmentJSON = ?", []interface{}{string(queryJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	// Print the result data and its type for debugging
	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}

func (am *appointmentManager) CreateAppointment(appointmentData *requestData.AppointmentObj) (*spResponse.Result, error) {
	// Convert input to DTO
	appointmentDTO := builder.BuildAppointmentDTO(appointmentData)
	fmt.Println("appointmentDTO:", appointmentDTO)

	// Convert DTO to JSON
	appointmentJSON, err := json.Marshal(appointmentDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal appointment data: %w", err)
	}

	// Create an instance of the stored procedure executor
	spExecutor := common.NewStoredProcedureExecutor()

	// Execute the stored procedure with the appointment data
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_InsertAppointment @AppointmentJSON = ?", []interface{}{string(appointmentJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}