package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
)

type MedicineMasterDTO struct {
	MedicineID    uint   `json:"medicineId"`
	MedicineName  string `json:"medicineName"`
	GenericName   string `json:"genericName"`
	BrandName     string `json:"brandName"`
	Composition   string `json:"composition"`
	Form          string `json:"form"`
	Strength      uint `json:"strength"`
	Unit          string    `json:"unit"`
	Manufacturer  string `json:"manufacturer"`
	HsnCode       string `json:"hsnCode"`
	GstPercentage float64 `json:"gstPercentage"`
	ReorderLevel  uint `json:"reorderLevel"`
	Status        uint    `json:"status"`
	CreatedAt     string `json:"createdAt"`
	CreatedBy     uint   `json:"createdBy"`
	ModifiedAt    string `json:"modifiedAt"`
	ModifiedBy    uint   `json:"modifiedBy"`
}

func BuildMedicineMasterObj(medicine *requestData.MedicineMasterObj) *MedicineMasterDTO {
	var medicineMasterObj MedicineMasterDTO
	medicineMasterObj.MedicineID = medicine.MedicineID
	medicineMasterObj.MedicineName = medicine.MedicineName
	medicineMasterObj.GenericName = medicine.GenericName
	medicineMasterObj.BrandName = medicine.BrandName
	medicineMasterObj.Composition = medicine.Composition
	medicineMasterObj.Form = medicine.Form
	medicineMasterObj.Strength = medicine.Strength
	medicineMasterObj.Unit = medicine.Unit
	medicineMasterObj.Manufacturer = medicine.Manufacturer
	medicineMasterObj.HsnCode = medicine.HsnCode
	medicineMasterObj.GstPercentage = medicine.GstPercentage
	medicineMasterObj.ReorderLevel = medicine.ReorderLevel
	medicineMasterObj.Status = uint(medicine.Status)
	medicineMasterObj.CreatedAt = medicine.CreatedAt
	medicineMasterObj.CreatedBy = medicine.CreatedBy
	medicineMasterObj.ModifiedAt = medicine.ModifiedAt
	medicineMasterObj.ModifiedBy = medicine.ModifiedBy

	return &medicineMasterObj
}

func BuildMedicineMasterDTOs(medicineMasterData []map[string]interface{}) []*MedicineMasterDTO {
	var medicineMasterDTOs []*MedicineMasterDTO

	for _, medicineMap := range medicineMasterData {
		medicineDTO := &MedicineMasterDTO{
			MedicineID:    common.ToUint(medicineMap["MedicineID"]),
			MedicineName:  common.ToString(medicineMap["MedicineName"]),
			GenericName:   common.ToString(medicineMap["GenericName"]),
			BrandName:     common.ToString(medicineMap["BrandName"]),
			Composition:   common.ToString(medicineMap["Composition"]),
			Form:          common.ToString(medicineMap["Form"]),
			Strength:      common.ToUint(medicineMap["Strength"]),
			Unit:          common.ToString(medicineMap["Unit"]),
			Manufacturer:  common.ToString(medicineMap["Manufacturer"]),
			HsnCode:       common.ToString(medicineMap["HsnCode"]),
			GstPercentage: common.ToFloat64(medicineMap["GstPercentage"]),
			ReorderLevel:  common.ToUint(medicineMap["ReorderLevel"]),
			Status:        common.ToUint(medicineMap["Status"]),
			CreatedAt:     common.ToString(medicineMap["CreatedAt"]),
			CreatedBy:     common.ToUint(medicineMap["CreatedBy"]),
			ModifiedAt:    common.ToString(medicineMap["ModifiedAt"]),
			ModifiedBy:    common.ToUint(medicineMap["ModifiedBy"]),
		}
		medicineMasterDTOs = append(medicineMasterDTOs, medicineDTO)
	}

	return medicineMasterDTOs
}
