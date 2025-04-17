package requestData

import (
	"time"
)

type UserObj struct {
	UserID        uint      `json:"userid" gorm:"primaryKey"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	DOB           time.Time `json:"dateOfBirth"`
	Gender        string    `json:"gender"`
	ContactNumber string    `json:"contactNumber"`
	Email         string    `json:"email"`
	Address       string    `json:"address"`
	UserName      string    `json:"userName"`
	UserType      uint       `json:"userType"`
	Role          []RoleObj  `json:"role"` // Changed to []RoleObj to match the expected input type
	ProfilePic    string    `json:"profilePic"`	
	BiometricData []byte    `json:"biometricData"`
	PasswordHash  string    `json:"password"`
	Active        uint       `json:"active"`
	CreatedAt     time.Time `json:"createdAt"`
	CreatedBy     string    `json:"createdBy"`
	ModifiedAt    time.Time `json:"modifiedAt"`
	ModifiedBy    string    `json:"modifiedBy"`
}
type UserCreationInput struct {
	FirstName string `json:"firstName"` //this is called field tag (the format of object)
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserCreationInput() *UserCreationInput {
	return &UserCreationInput{}
}

func NewUserLoginInput() *UserLoginInput {
	return &UserLoginInput{}
}

func NewUserOb() *UserObj {
	return &UserObj{}
}

