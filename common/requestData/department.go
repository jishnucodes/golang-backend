package requestData

import "time"

type DepartmentObj struct {
	DepartmentID     uint      `json:"departmentId"`
	DepartmentName   string    `json:"departmentName"`
	HeadOfDepartment uint      `json:"headOfDepartment"`
	ContactNumber    string    `json:"contactNumber"`
	Email            string    `json:"email"`
	Location         string    `json:"location"`
	OperatingHours   string    `json:"operatingHours"`
	NumberOfStaff    int       `json:"numberOfStaff"`
	ServicesOffered  string    `json:"servicesOffered"`
	Status           int       `json:"status"`
	CreatedAt        time.Time `json:"createdAt"`
	CreatedBy        string    `json:"createdBy"`
	ModifiedAt       time.Time `json:"modifiedAt"`
	ModifiedBy       string    `json:"modifiedBy"`
}

func NewDepartmentObj() *DepartmentObj {
	return &DepartmentObj{}
}
