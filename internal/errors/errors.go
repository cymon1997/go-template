package errors

import (
	"fmt"

	"github.com/cymon1997/go-template/internal/http"
)

type Err struct {
	Status  int
	Code    http.Code
	Message string
}

func (e *Err) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
