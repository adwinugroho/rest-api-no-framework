package http

type (
	//Request parameter
	Request struct {
		Protocol string
		Host     string
		Port     int
		Path     string //url
		Body     interface{}
	}

	//Response body
	Response struct {
		Status       bool        `json:"status"`
		Code         int         `json:"code"`
		ErrorMessage string      `json:"errMessage,omitempty"`
		Data         interface{} `json:"data,omitempty"`
		Message      interface{} `json:"message,omitempty"`
	}
)
