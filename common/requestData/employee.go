package requestData

// import "time"

type EmployeeObj struct {
	EmployeeID    uint      `json:"employeeId"`
	EmployeeCode  string    `json:"employeeCode"`
	UserName      string    `json:"userName"`
	Password      string    `json:"password"`
	ProfileImage  string    `json:"profileImage"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"` 
	Email         string    `json:"email"`
	PhoneNumber   string    `json:"phoneNumber"`
	MobileNumber  string    `json:"mobileNumber"`
	Address       string    `json:"address"`
	BloodGroup    string    `json:"bloodGroup"`
	HireDate      string    `json:"hireDate"`
	JobTitle      string    `json:"jobTitle"`
	Type           uint      `json:"type"`
	DepartmentID  uint      `json:"departmentId"`
	Department    uint      `json:"department"`
	EmployeeType  uint      `json:"employeeType"`
	IsOnLeave    bool      `json:"isOnLeave"`
	LeaveDateFrom string    `json:"leaveDateFrom"`
	LeaveDateTo   string    `json:"leaveDateTo"`
	CreatedAt     string    `json:"createdAt"`
	CreatedBy     uint      `json:"createdBy"`
	ModifiedAt    string    `json:"modifiedAt"`
	ModifiedBy    uint      `json:"modifiedBy"`
}


type SearchQuery struct {
	EmployeeType *int `form:"employeeType" json:"employeeType"`
}

type DoctorSearchQuery struct {
	EmployeeType *int `form:"employeeType" json:"employeeType"`
	DepartmentID *int `form:"departmentId" json:"departmentId"`
	InputDate    *string `form:"inputDate" json:"inputDate"`
}

func NewEmployeeObj() *EmployeeObj {
	return &EmployeeObj{}
}

