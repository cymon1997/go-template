package middleware

import (
	"github.com/cymon1997/go-template/internal/context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (m *middlewareImpl) AddContextID(ctx *gin.Context) {
	ctx.Set(string(context.IDKey), uuid.New().String())
}
