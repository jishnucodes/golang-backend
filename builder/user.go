package builder

import (
	"clinic-management/backend/common"
	"time"
)

// UserDTO represents the structure of the user data.
type UserDTO struct {
	UserID        uint     
	FirstName     string    
	LastName      string    
	DOB           time.Time                     
	Gender        string    
	ContactNumber string    
	Email         string    
	Address       string    
	Role          string    
	BiometricData []byte    
	Password  string    
	CreatedAt     time.Time 
	CreatedBy     string                   
	ModifiedAt    time.Time 
	ModifiedBy    string    


}

// BuildUserDTO constructs and returns a UserDTO from userData.
func BuildUserDTO(userData *common.UserObj) *UserDTO {
	var userObj UserDTO

	userObj.UserID = userData.UserID
	userObj.FirstName = userData.FirstName
	userObj.LastName = userData.LastName
	userObj.DOB = userData.DOB
	userObj.Gender = userData.Gender
	userObj.ContactNumber = userData.ContactNumber
	userObj.Email = userData.Email
	userObj.Address = userData.Address
	userObj.Role = userData.Role
	userObj.BiometricData = userData.BiometricData
	userObj.Password = userData.PasswordHash
	userObj.CreatedAt = userData.CreatedAt
	userObj.CreatedBy = userData.CreatedBy
	userObj.ModifiedAt = userData.ModifiedAt
	userObj.ModifiedBy = userData.ModifiedBy

	return &userObj;
}
