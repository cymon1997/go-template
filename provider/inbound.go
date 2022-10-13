package provider

import (
	"sync"

	"github.com/cymon1997/go-template/core/inbound/http"
)

var (
	httpHandler     http.Handler
	syncHttpHandler sync.Once
)

func GetInboundHttpHandler() http.Handler {
	syncHttpHandler.Do(func() {
		httpHandler = http.New(GetMiddleware(), GetUsecase())
	})
	return httpHandler
}
