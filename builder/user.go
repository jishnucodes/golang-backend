package builder

import (
	"clinic-management/backend/common"
	"time"
)

// UserDTO represents the structure of the user data.
type UserDTO struct {
	UserID        uint      `json:"userId"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	DOB           time.Time `json:"dateOfBirth"`
	Gender        string    `json:"gender"`
	ContactNumber string    `json:"contactNumber"`
	Email         string    `json:"email"`
	Address       string    `json:"address"`
	Role          string    `json:"role"`
	BiometricData []byte    `json:"biometricData"`
	Password      string    `json:"password"`
	CreatedAt     time.Time `json:"createdAt"`
	CreatedBy     string    `json:"createdBy"`
	ModifiedAt    time.Time `json:"modifiedAt"`
	ModifiedBy    string    `json:"modifiedBy"`
}

// BuildUserDTO constructs and returns a UserDTO from userData.
func BuildUserDTO(userData *common.UserObj) *UserDTO {
	var userObj UserDTO

	// Mapping fields from UserObj to UserDTO
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

	return &userObj
}


// BuildUserDTOs accepts a slice of UserDTO and returns it (or processes it further).
func BuildUserDTOs(userDTOList []UserDTO) []*UserDTO {
	var userDTOs []*UserDTO 
	// so we can just append them directly to the return list.
	for _, userDTO := range userDTOList {
		userDTOs = append(userDTOs, &userDTO)
	}
	return userDTOs
}