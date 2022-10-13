package router

import "github.com/gin-gonic/gin"

type Router interface {
	Instance() *gin.Engine
}

type routerImpl struct {
	engine *gin.Engine
}

func (r *routerImpl) Instance() *gin.Engine {
	return r.engine
}

func New() Router {
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	return &routerImpl{
		engine: engine,
	}
}
