package builder

import (
	"clinic-management/backend/common"
	"encoding/base64"
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

// BuildUserDTOs constructs a slice of UserDTO from []map[string]interface{}
func BuildUserDTOs(usersData []map[string]interface{}) []*UserDTO {
	var userDTOs []*UserDTO

	for _, userMap := range usersData {
		userDTO := &UserDTO{
			UserID:        toUint(userMap["UserID"]),
			FirstName:     toString(userMap["FirstName"]),
			LastName:      toString(userMap["LastName"]),
			DOB:           parseTime(userMap["DOB"]),
			Gender:        toString(userMap["Gender"]),
			ContactNumber: toString(userMap["ContactNumber"]),
			Email:         toString(userMap["Email"]),
			Address:       toString(userMap["Address"]),
			Role:          toString(userMap["Role"]),
			BiometricData: decodeBase64(toString(userMap["BiometricData"])),
			// Password:      toString(userMap["Password"]),
			CreatedAt:     parseTime(userMap["CreatedAt"]),
			CreatedBy:     toString(userMap["CreatedBy"]),
			ModifiedAt:    parseTime(userMap["ModifiedAt"]),
			ModifiedBy:    toString(userMap["ModifiedBy"]),
		}

		userDTOs = append(userDTOs, userDTO)
	}

	return userDTOs
}

// Helper functions for safe type conversion

// Safely convert to uint, handling nil and float64 values
func toUint(value interface{}) uint {
	if value == nil {
		return 0
	}
	switch v := value.(type) {
	case float64:
		return uint(v) // JSON unmarshals numbers as float64
	case int:
		return uint(v)
	}
	return 0
}

// Safely convert to string
func toString(value interface{}) string {
	if value == nil {
		return ""
	}
	if v, ok := value.(string); ok {
		return v
	}
	return ""
}

// Safely decode base64-encoded biometric data
func decodeBase64(value string) []byte {
	if value == "" {
		return nil
	}
	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return nil
	}
	return data
}

// Safely parse time.Time from RFC3339 string
func parseTime(value interface{}) time.Time {
	if value == nil {
		return time.Time{}
	}
	if str, ok := value.(string); ok {
		t, err := time.Parse(time.RFC3339, str)
		if err == nil {
			return t
		}
	}
	return time.Time{}
}




