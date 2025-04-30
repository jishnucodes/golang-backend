package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
)

type AutoNumberDTO struct {
	Type      uint   `json:"type"`
	Prefix    string `json:"prefix"`
	Increment uint   `json:"increment"`
	Suffix    string `json:"suffix"`
	GeneratedCode string `json:"generatedCode"`
}

func BuildAutoNumberDTO(autoNumber *requestData.AutoNumberObj) *AutoNumberDTO {
	var autoNumberDTO AutoNumberDTO
	autoNumberDTO.Type = autoNumber.Type
	autoNumberDTO.Prefix = autoNumber.Prefix
	autoNumberDTO.Increment = autoNumber.Increment
	autoNumberDTO.Suffix = autoNumber.Suffix
	autoNumberDTO.GeneratedCode = autoNumber.GeneratedCode
	return &autoNumberDTO
}


func BuildAutoNumberDTOs(autoNumberData []map[string]interface{}) []*AutoNumberDTO {
	var autoNumberDTOs []*AutoNumberDTO

	for _, autoNumberMap := range autoNumberData {
		autoNumberDTO := &AutoNumberDTO{
			Type:      common.ToUint(autoNumberMap["Type"]),
			Prefix:    common.ToString(autoNumberMap["Prefix"]),
			Increment: common.ToUint(autoNumberMap["Increment"]),
			Suffix:    common.ToString(autoNumberMap["Suffix"]),
			GeneratedCode: common.ToString(autoNumberMap["GeneratedCode"]),
		}
		autoNumberDTOs = append(autoNumberDTOs, autoNumberDTO)
	}
	return autoNumberDTOs
}




