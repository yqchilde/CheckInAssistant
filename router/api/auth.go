package api

import (
	"CheckInAssistant/pkg/app"
	"github.com/gin-gonic/gin"
)

type auth struct {
	UserName string `valid:"Required; MaxSize(50)"`
	PassWord string `valid:"Required; MaxSize(50)"`
}

func GetAuth(ctx *gin.Context) {
	appG := app.Gin{
		C: ctx,
	}

	// valication
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	a := auth{
		UserName: username,
		PassWord: password,
	}

	
}