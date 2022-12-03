package httpjson

type Payload struct {
	StatusCode int    `json:"status_code"`
	StatusName string `json:"status_name"`
	Message    string `json:"message,omitempty"`
}
