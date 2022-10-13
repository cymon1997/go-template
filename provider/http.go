package provider

import (
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/cymon1997/go-template/internal/http/middleware"
	iRouter "github.com/cymon1997/go-template/internal/http/router"
)

var (
	mw     middleware.Middleware
	syncMw sync.Once

	router     iRouter.Router
	syncRouter sync.Once

	httpServer     *gin.Engine
	syncHttpServer sync.Once
)

func GetMiddleware() middleware.Middleware {
	syncMw.Do(func() {
		mw = middleware.New()
	})
	return mw
}

func GetRouter() iRouter.Router {
	syncRouter.Do(func() {
		router = iRouter.New()
	})
	return router
}

func GetHttpServer() *gin.Engine {
	syncHttpServer.Do(func() {
		httpServer = GetRouter().Instance()
		GetInboundHttpHandler().Register(httpServer)
	})
	return httpServer
}
