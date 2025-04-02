package result

type Result struct {
	Data string  `json:"data,omitempty"`
	Status int 					    `json:"status,omitempty"`
	StatusCode  string				`json:"statusCode,omitempty"`
	StatusMessage string			`json:"statusMessage,omitempty"`

}

func NewResult() *Result {
	return &Result{}
}