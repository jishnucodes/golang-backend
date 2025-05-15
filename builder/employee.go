package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
)

// EmployeeDTO represents the structure of the employee data for API responses
type EmployeeDTO struct {
	EmployeeID    uint   `json:"employeeId"`
	EmployeeCode  string `json:"employeeCode"`
	UserName      string `json:"userName"`
	Password      string `json:"password"`
	ProfileImage  string `json:"profileImage"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phoneNumber"`
	MobileNumber  string `json:"mobileNumber"`
	Address       string `json:"address"`
	BloodGroup    string `json:"bloodGroup"`
	HireDate      string `json:"hireDate"`
	Specialty     string `json:"specialty"`
	ConsultationFee float64 `json:"consultationFee"`
	JobTitle      string `json:"jobTitle"`
	DepartmentID  uint   `json:"departmentId"`
	Department    uint `json:"department"`
	Type          uint   `json:"type"`
	EmployeeType  uint   `json:"employeeType"`
	IsOnLeave    bool      `json:"isOnLeave"`
	LeaveDateFrom string    `json:"leaveDateFrom"`
	LeaveDateTo   string    `json:"leaveDateTo"`
	CreatedAt     string `json:"createdAt"`
	CreatedBy     uint `json:"createdBy"`
	ModifiedAt    string `json:"modifiedAt"`
	ModifiedBy    uint `json:"modifiedBy"`
}

// BuildEmployeeDTO constructs and returns an EmployeeDTO from employeeData
func BuildEmployeeDTO(employeeData *requestData.EmployeeObj) *EmployeeDTO {
	var employeeObj EmployeeDTO

	employeeObj.EmployeeID = employeeData.EmployeeID
	employeeObj.EmployeeCode = employeeData.EmployeeCode
	employeeObj.UserName = employeeData.UserName
	employeeObj.Password = employeeData.Password
	employeeObj.ProfileImage = employeeData.ProfileImage
	employeeObj.FirstName = employeeData.FirstName
	employeeObj.LastName = employeeData.LastName
	employeeObj.Email = employeeData.Email
	employeeObj.PhoneNumber = employeeData.PhoneNumber
	employeeObj.MobileNumber = employeeData.MobileNumber
	employeeObj.Address = employeeData.Address
	employeeObj.BloodGroup = employeeData.BloodGroup
	employeeObj.HireDate = employeeData.HireDate
	employeeObj.JobTitle = employeeData.JobTitle
	if employeeData.DepartmentID != 0 {
		employeeObj.DepartmentID = employeeData.DepartmentID
	} else {
		employeeObj.DepartmentID = employeeData.Department
	}
	
	// employeeObj.Department = employeeData.Department
	employeeObj.Type = employeeData.Type
	employeeObj.EmployeeType = employeeData.EmployeeType
	employeeObj.IsOnLeave = employeeData.IsOnLeave
	employeeObj.LeaveDateFrom = employeeData.LeaveDateFrom
	employeeObj.LeaveDateTo = employeeData.LeaveDateTo
	employeeObj.CreatedAt = employeeData.CreatedAt
	employeeObj.CreatedBy = employeeData.CreatedBy
	employeeObj.ModifiedAt = employeeData.ModifiedAt
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
			UserName:      common.ToString(employeeMap["UserName"]),
			Password:      common.ToString(employeeMap["Password"]),
			ProfileImage:  common.ToString(employeeMap["ProfileImage"]),
			FirstName:     common.ToString(employeeMap["FirstName"]),
			LastName:      common.ToString(employeeMap["LastName"]),
			Email:         common.ToString(employeeMap["Email"]),
			PhoneNumber:   common.ToString(employeeMap["PhoneNumber"]),
			MobileNumber:  common.ToString(employeeMap["MobileNumber"]),
			Address:       common.ToString(employeeMap["Address"]),
			BloodGroup:    common.ToString(employeeMap["BloodGroup"]),
			HireDate:      common.ToString(employeeMap["HireDate"]),
			Specialty:     common.ToString(employeeMap["Specialty"]),
			ConsultationFee: common.ToFloat64(employeeMap["ConsultationFee"]),
			JobTitle:      common.ToString(employeeMap["JobTitle"]),
			DepartmentID:  common.ToUint(employeeMap["Department"]),
			Type:          common.ToUint(employeeMap["Type"]),
			EmployeeType:  common.ToUint(employeeMap["EmployeeType"]),
			IsOnLeave:     common.ToBool(employeeMap["IsOnLeave"]),
			LeaveDateFrom: common.ToString(employeeMap["LeaveDateFrom"]),
			LeaveDateTo:   common.ToString(employeeMap["LeaveDateTo"]),
			CreatedAt:     common.ToString(employeeMap["CreatedAt"]),
			CreatedBy:     common.ToUint(employeeMap["CreatedBy"]),
			ModifiedAt:    common.ToString(employeeMap["ModifiedAt"]),
			ModifiedBy:    common.ToUint(employeeMap["ModifiedBy"]),
		}
		employeeDTOs = append(employeeDTOs, employeeDTO)
	}

	return employeeDTOs
}
