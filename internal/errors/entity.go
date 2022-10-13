package errors

import "net/http"

var (
	ErrUnknown = &Err{
		Status:  http.StatusInternalServerError,
		Code:    "ERROR_UNKNOWN",
		Message: "Unexpected error",
	}
)
