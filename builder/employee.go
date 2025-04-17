package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
)

// EmployeeDTO represents the structure of the employee data for API responses
type EmployeeDTO struct {
	EmployeeID    uint   `json:"employeeId"`
	EmployeeCode  string `json:"employeeCode"`
	UserID        uint   `json:"userId"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phoneNumber"`
	MobileNumber  string `json:"mobileNumber"`
	Address       string `json:"address"`
	BloodGroup    string `json:"bloodGroup"`
	HireDate      string `json:"hireDate"`
	JobTitle      string `json:"jobTitle"`
	DepartmentID  uint   `json:"departmentId"`
	EmployeeType  string `json:"employeeType"`
	CreatedAt     string `json:"createdAt"`
	CreatedBy     string `json:"createdBy"`
	ModifiedAt    string `json:"modifiedAt"`
	ModifiedBy    string `json:"modifiedBy"`
}

// BuildEmployeeDTO constructs and returns an EmployeeDTO from employeeData
func BuildEmployeeDTO(employeeData *requestData.EmployeeObj) *EmployeeDTO {
	var employeeObj EmployeeDTO

	employeeObj.EmployeeID = employeeData.EmployeeID
	employeeObj.EmployeeCode = employeeData.EmployeeCode
	employeeObj.UserID = employeeData.UserID
	employeeObj.FirstName = employeeData.FirstName
	employeeObj.LastName = employeeData.LastName
	employeeObj.Email = employeeData.Email
	employeeObj.PhoneNumber = employeeData.PhoneNumber
	employeeObj.MobileNumber = employeeData.MobileNumber
	employeeObj.Address = employeeData.Address
	employeeObj.BloodGroup = employeeData.BloodGroup
	employeeObj.HireDate = employeeData.HireDate.Format("2006-01-02 15:04:05")
	employeeObj.JobTitle = employeeData.JobTitle
	employeeObj.DepartmentID = employeeData.DepartmentID
	employeeObj.EmployeeType = employeeData.EmployeeType
	employeeObj.CreatedAt = employeeData.CreatedAt.Format("2006-01-02 15:04:05")
	employeeObj.CreatedBy = employeeData.CreatedBy
	employeeObj.ModifiedAt = employeeData.ModifiedAt.Format("2006-01-02 15:04:05")
	employeeObj.ModifiedBy = employeeData.ModifiedBy

	return &employeeObj
}

// BuildEmployeeDTOs constructs a slice of EmployeeDTO from []map[string]interface{}
func BuildEmployeeDTOs(employeesData []map[string]interface{}) []*EmployeeDTO {
	var employeeDTOs []*EmployeeDTO

	for _, employeeMap := range employeesData {
		employeeDTO := &EmployeeDTO{
			EmployeeID:    common.ToUint(employeeMap["EmployeeID"]),
			EmployeeCode:  common.ToString(employeeMap["EmployeeCode"]),
			UserID:        common.ToUint(employeeMap["UserID"]),
			FirstName:     common.ToString(employeeMap["FirstName"]),
			LastName:      common.ToString(employeeMap["LastName"]),
			Email:         common.ToString(employeeMap["Email"]),
			PhoneNumber:   common.ToString(employeeMap["PhoneNumber"]),
			MobileNumber:  common.ToString(employeeMap["MobileNumber"]),
			Address:       common.ToString(employeeMap["Address"]),
			BloodGroup:    common.ToString(employeeMap["BloodGroup"]),
			HireDate:      common.ToString(employeeMap["HireDate"]),
			JobTitle:      common.ToString(employeeMap["JobTitle"]),
			DepartmentID:  common.ToUint(employeeMap["Department"]),
			EmployeeType:  common.ToString(employeeMap["EmployeeType"]),
			CreatedAt:     common.ToString(employeeMap["CreatedAt"]),
			CreatedBy:     common.ToString(employeeMap["CreatedBy"]),
			ModifiedAt:    common.ToString(employeeMap["ModifiedAt"]),
			ModifiedBy:    common.ToString(employeeMap["ModifiedBy"]),
		}
		employeeDTOs = append(employeeDTOs, employeeDTO)
	}

	return employeeDTOs
}
