package http

import (
	"log"
	"net/http"

	"github.com/cymon1997/go-template/internal/errors"
	iHttp "github.com/cymon1997/go-template/internal/http"
	"github.com/gin-gonic/gin"
)

func (h *handlerImpl) Status(ctx *gin.Context) {
	resp, err := h.uc.Status(ctx)
	if err != nil {
		log.Println("error execute Status usecase", err)
		ctx.JSON(errors.GetStatus(err), iHttp.BaseResponse{
			Code:    errors.GetCode(err),
			Message: errors.GetMessage(err),
			Payload: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, iHttp.BaseResponse{
		Code:    iHttp.CodeOK,
		Payload: resp,
	})
}
