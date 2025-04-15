package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
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
	UserName      string    `json:"userName"`
	UserType      uint       `json:"userType"`
	ProfilePic    string    `json:"profilePic"`
	BiometricData []byte    `json:"biometricData"`
	Password      string    `json:"password"`
	Active        uint       `json:"active"`
	CreatedAt     string    `json:"createdAt"`
	CreatedBy     string    `json:"createdBy"`
	ModifiedAt    string    `json:"modifiedAt"`
	ModifiedBy    string    `json:"modifiedBy"`
}


// BuildUserDTO constructs and returns a UserDTO from userData.
func BuildUserDTO(userData *requestData.UserObj) *UserDTO {
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
	userObj.UserName = userData.UserName
	userObj.UserType = userData.UserType
	userObj.ProfilePic = userData.ProfilePic
	userObj.BiometricData = userData.BiometricData
	userObj.Password = userData.PasswordHash
	userObj.Active = userData.Active
	userObj.CreatedAt = userData.CreatedAt.Format("2006-01-02 15:04:05")
	userObj.CreatedBy = userData.CreatedBy
	userObj.ModifiedAt = userData.ModifiedAt.Format("2006-01-02 15:04:05")
	userObj.ModifiedBy = userData.ModifiedBy

	return &userObj
}

// BuildUserDTOs constructs a slice of UserDTO from []map[string]interface{}
func BuildUserDTOs(usersData []map[string]interface{}) []*UserDTO {
	var userDTOs []*UserDTO

	for _, userMap := range usersData {
		userDTO := &UserDTO{
			UserID:        common.ToUint(userMap["UserID"]),
			FirstName:     common.ToString(userMap["FirstName"]),
			LastName:      common.ToString(userMap["LastName"]),
			DOB:           common.ParseTime(userMap["DOB"]),
			Gender:        common.ToString(userMap["Gender"]),
			ContactNumber: common.ToString(userMap["ContactNumber"]),
			Email:         common.ToString(userMap["Email"]),
			Address:       common.ToString(userMap["Address"]),
			UserName:      common.ToString(userMap["UserName"]),
			UserType:      common.ToUint(userMap["UserType"]),
			ProfilePic:    common.ToString(userMap["ProfilePic"]),
			BiometricData: common.DecodeBase64(common.ToString(userMap["BiometricData"])),
			Active:        common.ToUint(userMap["Active"]),
			// Password:      toString(userMap["Password"]),
			CreatedAt:     common.ToString(userMap["CreatedAt"]),
			CreatedBy:     common.ToString(userMap["CreatedBy"]),
			ModifiedAt:    common.ToString(userMap["ModifiedAt"]),
			ModifiedBy:    common.ToString(userMap["ModifiedBy"]),
		}

		userDTOs = append(userDTOs, userDTO)
	}

	return userDTOs
}

// Helper functions for safe type conversion






