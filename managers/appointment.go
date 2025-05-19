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
	CreateAppointment(appointmentData *requestData.AppointmentObj) (*spResponse.Result, error)
}

type appointmentManager struct {
}

func NewAppointmentManager() AppointmentManager {
	return &appointmentManager{}
}



func (dm *appointmentManager) CreateAppointment(appointmentData *requestData.AppointmentObj) (*spResponse.Result, error) {
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