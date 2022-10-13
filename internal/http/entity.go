package http

type BaseResponse struct {
	Code    Code        `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

type Code string
