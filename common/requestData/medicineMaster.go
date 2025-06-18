package requestData

// import "time"

type MedicineMasterObj struct {
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
	Status        uint   `json:"status"`
	CreatedAt     string `json:"createdAt"`
	CreatedBy     uint   `json:"createdBy"`
	ModifiedAt    string `json:"modifiedAt"`
	ModifiedBy    uint   `json:"modifiedBy"`
}

func NewMedicineMasterObj() *MedicineMasterObj {
	return &MedicineMasterObj{}
}
