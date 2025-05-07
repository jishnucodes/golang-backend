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
	EmployeeType  uint      `json:"employeeType"`
	CreatedAt     string    `json:"createdAt"`
	CreatedBy     string    `json:"createdBy"`
	ModifiedAt    string    `json:"modifiedAt"`
	ModifiedBy    string    `json:"modifiedBy"`
}


type SearchQuery struct {
	EmployeeType *int `form:"employeeType" json:"employeeType"`
}

func NewEmployeeObj() *EmployeeObj {
	return &EmployeeObj{}
}

