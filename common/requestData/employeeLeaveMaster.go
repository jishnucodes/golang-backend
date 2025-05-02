package requestData

import "time"

type EmployeeLeaveObj struct {
	ID            uint      `json:"id"`
	EmployeeID    uint      `json:"employeeId"`
	DateFrom      time.Time `json:"dateFrom"`
	DateTo        time.Time `json:"dateTo"`
	TimeFrom      string    `json:"timeFrom"`
	TimeTo        string     `json:"timeTo"`
	Approved      uint      `json:"approved"`
	Remarks       string    `json:"remarks"`
	LeaveType     uint      `json:"leaveType"`
	CreatedAt     string    `json:"createdAt"`
	CreatedBy     string    `json:"createdBy"`
	ModifiedAt    string    `json:"modifiedAt"`
	ModifiedBy    string    `json:"modifiedBy"`
}

func NewEmployeeLeaveObj() *EmployeeLeaveObj {
	return &EmployeeLeaveObj{}
}

