package router

import (
	"github.com/gin-gonic/gin"

	"CheckInAssistant/pkg/app"
	"CheckInAssistant/service"
)

// GetShareUrl return application share url info
func GetShareUrl(ctx *gin.Context)  {
	data, err := service.NewAppService(ctx).GetShareUrl()
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", data)
}