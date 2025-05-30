package managers

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"clinic-management/backend/spResponse"
	"encoding/json"
	"fmt"
)

type PaymentManager interface{
	GetConfirmedAppointmentById(query *requestData.ConfirmedAppointmentSearchQuery) (*spResponse.Result, error)
}

type paymentManager struct {
}


func NewPaymentManager() PaymentManager {
	return &paymentManager{}
}

func (pm *paymentManager) GetConfirmedAppointmentById(query *requestData.ConfirmedAppointmentSearchQuery) (*spResponse.Result, error) {
	// Create an instance of StoredProcedureExecutor
	spExecutor := common.NewStoredProcedureExecutor()

	// Convert query to JSON
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal query: %w", err)
	}
	fmt.Println("queryJSON:", string(queryJSON))

	// Execute the stored procedure and capture the result
	data, err := spExecutor.ExecuteStoredProcedure("EXEC sp_CMS_ConfirmAppointment @AppointmentJSON = ?", []interface{}{string(queryJSON)})
	if err != nil {
		return nil, fmt.Errorf("error executing stored procedure: %w", err)
	}

	// Print the result data and its type for debugging
	fmt.Println("data:", data)
	fmt.Println("data type:", fmt.Sprintf("%T", data))

	return data, nil
}