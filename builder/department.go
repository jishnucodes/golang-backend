package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
)

type DepartmentDTO struct {
	DepartmentID     uint      `json:"departmentId"`
	DepartmentName   string    `json:"departmentName"`
	HeadOfDepartment uint      `json:"headOfDepartment"`
	ContactNumber    string    `json:"contactNumber"`
	Email            string    `json:"email"`
	Location         string    `json:"location"`
	OperatingHours   string    `json:"operatingHours"`
	NumberOfStaff    uint       `json:"numberOfStaff"`
	ServicesOffered  string    `json:"servicesOffered"`
	Status           uint      `json:"status"`
	CreatedAt        string    `json:"createdAt"`
	CreatedBy        uint    `json:"createdBy"`
	ModifiedAt       string    `json:"modifiedAt"`
	ModifiedBy       uint    `json:"modifiedBy"`
}


func BuildDepartmentObj(department *requestData.DepartmentObj) *DepartmentDTO {
	var departmentObj DepartmentDTO
	departmentObj.DepartmentID = department.DepartmentID
	departmentObj.DepartmentName = department.DepartmentName
	departmentObj.HeadOfDepartment = department.HeadOfDepartment
	departmentObj.ContactNumber = department.ContactNumber
	departmentObj.Email = department.Email
	departmentObj.Location = department.Location
	departmentObj.OperatingHours = department.OperatingHours
	departmentObj.NumberOfStaff = uint(department.NumberOfStaff)
	departmentObj.ServicesOffered = department.ServicesOffered
	departmentObj.Status = uint(department.Status)
	departmentObj.CreatedAt = department.CreatedAt
	departmentObj.CreatedBy = department.CreatedBy
	departmentObj.ModifiedAt = department.ModifiedAt
	departmentObj.ModifiedBy = department.ModifiedBy

	return &departmentObj
}	


func BuildDepartmentDTOs(departmentData []map[string]interface{}) []*DepartmentDTO {
	var departmentDTOs []*DepartmentDTO

	for _, departmentMap := range departmentData {
		departmentDTO := &DepartmentDTO{
			DepartmentID:     common.ToUint(departmentMap["DepartmentID"]),
			DepartmentName:   common.ToString(departmentMap["DepartmentName"]),
			HeadOfDepartment: common.ToUint(departmentMap["HeadOfDepartment"]),
			ContactNumber:    common.ToString(departmentMap["ContactNumber"]),
			Email:            common.ToString(departmentMap["Email"]),
			Location:         common.ToString(departmentMap["Location"]),
			OperatingHours:   common.ToString(departmentMap["OperatingHours"]),
			NumberOfStaff:    common.ToUint(departmentMap["NumberOfStaff"]),
			ServicesOffered:  common.ToString(departmentMap["ServicesOffered"]),
			Status:           common.ToUint(departmentMap["Status"]),
			CreatedAt:        common.ToString(departmentMap["CreatedAt"]),
			CreatedBy:        common.ToUint(departmentMap["CreatedBy"]),
			ModifiedAt:       common.ToString(departmentMap["ModifiedAt"]),
			ModifiedBy:       common.ToUint(departmentMap["ModifiedBy"]),
		}
		departmentDTOs = append(departmentDTOs, departmentDTO)
	}
	return departmentDTOs
}
