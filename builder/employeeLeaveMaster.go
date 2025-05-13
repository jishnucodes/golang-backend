package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
)

// EmployeeDTO represents the structure of the employee data for API responses
type EmployeeLeaveDTO struct {
	ID            uint   `json:"id"`
	EmployeeID    uint   `json:"employeeId"`
	DateFrom      string `json:"dateFrom"`
	DateTo        string `json:"dateTo"`
	TimeFrom      string `json:"timeFrom"`
	TimeTo        string `json:"timeTo"`
	Approved      uint   `json:"approved"`
	Remarks       string `json:"remarks"`
	LeaveType     uint   `json:"leaveType"`
	CreatedAt     string `json:"createdAt"`
	CreatedBy     uint   `json:"createdBy"`
	ModifiedAt    string `json:"modifiedAt"`
	ModifiedBy    uint   `json:"modifiedBy"`
}

// BuildEmployeeDTO constructs and returns an EmployeeDTO from employeeData
func BuildEmployeeLeaveDTO(employeeLeaveData *requestData.EmployeeLeaveObj) *EmployeeLeaveDTO {
	var employeeLeaveObj EmployeeLeaveDTO

	employeeLeaveObj.ID = employeeLeaveData.ID
	employeeLeaveObj.EmployeeID = employeeLeaveData.EmployeeID
	employeeLeaveObj.DateFrom = employeeLeaveData.DateFrom.Format("2006-01-02")
	employeeLeaveObj.DateTo = employeeLeaveData.DateTo.Format("2006-01-02")
	employeeLeaveObj.TimeFrom = employeeLeaveData.TimeFrom
	employeeLeaveObj.TimeTo = employeeLeaveData.TimeTo
	employeeLeaveObj.Approved = employeeLeaveData.Approved
	employeeLeaveObj.Remarks = employeeLeaveData.Remarks
	employeeLeaveObj.LeaveType = employeeLeaveData.LeaveType
	employeeLeaveObj.CreatedAt = employeeLeaveData.CreatedAt
	employeeLeaveObj.CreatedBy = employeeLeaveData.CreatedBy
	employeeLeaveObj.ModifiedAt = employeeLeaveData.ModifiedAt
	employeeLeaveObj.ModifiedBy = employeeLeaveData.ModifiedBy

	return &employeeLeaveObj
}

// BuildEmployeeDTOs constructs a slice of EmployeeDTO from []map[string]interface{}
func BuildEmployeeLeaveDTOs(employeeLeaveData []map[string]interface{}) []*EmployeeLeaveDTO {
	var employeeLeaveDTOs []*EmployeeLeaveDTO

	for _, employeeLeaveMap := range employeeLeaveData {
		employeeLeaveDTO := &EmployeeLeaveDTO{	
			ID:            common.ToUint(employeeLeaveMap["ID"]),
			EmployeeID:    common.ToUint(employeeLeaveMap["EmployeeID"]),
			DateFrom:      common.ToString(employeeLeaveMap["DateFrom"]),
			DateTo:        common.ToString(employeeLeaveMap["DateTo"]),
			TimeFrom:      common.ToString(employeeLeaveMap["TimeFrom"]),
			TimeTo:        common.ToString(employeeLeaveMap["TimeTo"]),
			Approved:      common.ToUint(employeeLeaveMap["Approved"]),
			Remarks:       common.ToString(employeeLeaveMap["Remarks"]),
			LeaveType:     common.ToUint(employeeLeaveMap["LeaveType"]),
			CreatedAt:     common.ToString(employeeLeaveMap["CreatedAt"]),
			CreatedBy:     common.ToUint(employeeLeaveMap["CreatedBy"]),
			ModifiedAt:    common.ToString(employeeLeaveMap["ModifiedAt"]),
			ModifiedBy:    common.ToUint(employeeLeaveMap["ModifiedBy"]),
		}
		employeeLeaveDTOs = append(employeeLeaveDTOs, employeeLeaveDTO)
	}

	return employeeLeaveDTOs
}
