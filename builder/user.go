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
	Role          string    `json:"role"`
	BiometricData []byte    `json:"biometricData"`
	Password      string    `json:"password"`
	CreatedAt     time.Time `json:"createdAt"`
	CreatedBy     string    `json:"createdBy"`
	ModifiedAt    time.Time `json:"modifiedAt"`
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
	userObj.Role = userData.Role
	userObj.BiometricData = userData.BiometricData
	userObj.Password = userData.PasswordHash
	userObj.CreatedAt = userData.CreatedAt
	userObj.CreatedBy = userData.CreatedBy
	userObj.ModifiedAt = userData.ModifiedAt
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
			Role:          common.ToString(userMap["Role"]),
			BiometricData: common.DecodeBase64(common.ToString(userMap["BiometricData"])),
			// Password:      toString(userMap["Password"]),
			CreatedAt:     common.ParseTime(userMap["CreatedAt"]),
			CreatedBy:     common.ToString(userMap["CreatedBy"]),
			ModifiedAt:    common.ParseTime(userMap["ModifiedAt"]),
			ModifiedBy:    common.ToString(userMap["ModifiedBy"]),
		}

		userDTOs = append(userDTOs, userDTO)
	}

	return userDTOs
}

// Helper functions for safe type conversion






