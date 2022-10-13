package http

import (
	"github.com/cymon1997/go-template/core/usecase"
	"github.com/cymon1997/go-template/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Register(engine *gin.Engine)
}

type handlerImpl struct {
	mw middleware.Middleware
	uc usecase.Usecase
}

func New(mw middleware.Middleware, uc usecase.Usecase) Handler {
	return &handlerImpl{
		mw: mw,
		uc: uc,
	}
}

func (h *handlerImpl) Register(engine *gin.Engine) {
	// TODO: Register your endpoints here
	root := engine.Group("", h.mw.AddContextID)
	root.GET("/version", h.Version)
}
