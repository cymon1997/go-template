package middleware

import (
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	AddContextID(ctx *gin.Context)
}

type middlewareImpl struct {
}

func New() Middleware {
	return &middlewareImpl{}
}
