package router

import (
	"github.com/gin-gonic/gin"

	"CheckInAssistant/model"
	"CheckInAssistant/pkg/app"
	"CheckInAssistant/service"
	"CheckInAssistant/viewmodel"
)

// UserLogin ...
func UserLogin(ctx *gin.Context) {
	user := new(viewmodel.User)
	errDecode := ctx.ShouldBindJSON(user)
	if errDecode != nil {
		app.NewErrorBodyJSON(ctx.Writer)
		return
	}

	if user.Phone == "" {
		app.NewErrorMissField(ctx.Writer, "phone")
		return
	}
	if user.Password == "" {
		app.NewErrorMissField(ctx.Writer, "password")
		return
	}

	data, err := service.NewUserService(ctx).UserLogin(user)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", data)
}

// UserRegister ...
func UserRegister(ctx *gin.Context) {
	user := new(viewmodel.User)
	errDecode := ctx.ShouldBindJSON(user)
	if errDecode != nil {
		app.NewErrorBodyJSON(ctx.Writer)
		return
	}

	if user.Phone == "" {
		app.NewErrorMissField(ctx.Writer, "phone")
		return
	}
	if user.Password == "" {
		app.NewErrorMissField(ctx.Writer, "password")
		return
	}

	err := service.NewUserService(ctx).UserRegister(user)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", nil)
}

// UpdateUserInfo ...
func UpdateUserInfo(ctx *gin.Context) {
	user := new(model.User)
	errDecode := ctx.ShouldBindJSON(user)
	if errDecode != nil {
		app.NewErrorBodyJSON(ctx.Writer)
		return
	}

	typ := ctx.Query("type")
	if typ == "" {
		app.NewErrorMissField(ctx.Writer, "password")
		return
	}

	err := service.NewUserService(ctx).UpdateUserInfo(typ, user)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", nil)
}

// GetUserInfo get user info
func GetUserInfo(ctx *gin.Context) {
	userID := ctx.Query("user_id")

	data, err := service.NewUserService(ctx).GetUserInfo(userID)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", data)
}
