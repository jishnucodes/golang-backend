package requestData

// RoleMaster represents the role master table in the database
type AutoNumberObj struct {
	Type      uint   `json:"type"`
	Prefix    string `json:"prefix"`
	Increment uint   `json:"increment"`
	Suffix    string `json:"suffix"`
	GeneratedCode string `json:"generatedCode"`
}

func NewAutoNumberObj() *AutoNumberObj {
	return &AutoNumberObj{}
}
