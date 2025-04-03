package spResponse

type Result struct {
	Data string  `json:"data,omitempty"` //json string
	Status int 					    `json:"status,omitempty"`
	StatusCode  string				`json:"statusCode,omitempty"`
	StatusMessage string			`json:"statusMessage,omitempty"`

}

func NewResult() *Result {
	return &Result{}
}