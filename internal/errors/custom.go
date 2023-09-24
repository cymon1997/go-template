package errors

import "net/http"

var (
	ErrUnknown = &Err{
		Status:  http.StatusInternalServerError,
		Code:    "ERROR_UNKNOWN",
		Message: "Unexpected error",
	}

	ErrDatabaseOp = &Err{
		Status:  http.StatusInternalServerError,
		Code:    "ERROR_DATABASE_OPERATION",
		Message: "Error while query to database",
	}
)
