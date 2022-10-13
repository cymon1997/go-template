package errors

import (
	"net/http"

	iHttp "github.com/cymon1997/go-template/internal/http"
)

func GetStatus(err error) int {
	errs, ok := err.(*Err)
	if ok {
		return errs.Status
	}
	return http.StatusInternalServerError
}

func GetCode(err error) iHttp.Code {
	errs, ok := err.(*Err)
	if ok {
		return errs.Code
	}
	return iHttp.CodeUnknown
}

func GetMessage(err error) string {
	errs, ok := err.(*Err)
	if ok {
		return errs.Message
	}
	return errs.Error()
}
