package errors

import (
	"errors"
	"net/http"

	iHttp "github.com/cymon1997/go-template/internal/http"
)

func GetStatus(err error) int {
	var errs *Err
	ok := errors.As(err, &errs)
	if ok {
		return errs.Status
	}
	return http.StatusInternalServerError
}

func GetCode(err error) iHttp.Code {
	var errs *Err
	ok := errors.As(err, &errs)
	if ok {
		return errs.Code
	}
	return "UNKNOWN"
}

func GetMessage(err error) string {
	var errs *Err
	ok := errors.As(err, &errs)
	if ok {
		return errs.Message
	}
	return errs.Error()
}
