package common

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
	Role          string    `json:"role"`
	BiometricData []byte    `json:"biometricData"`
	PasswordHash  string    `json:"password"`
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

