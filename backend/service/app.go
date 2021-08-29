package service

import (
	"errors"

	"github.com/gin-gonic/gin"

	"CheckInAssistant/model"
	"CheckInAssistant/store"
)

type AppService struct {
	Ctx *gin.Context
}

func NewAppService(ctx *gin.Context) *AppService {
	return &AppService{Ctx: ctx}
}

// GetShareUrl return application share url info
func (a *AppService) GetShareUrl() (*model.AppUpgrade, error) {
	has, app, err := store.NewDBApp().GetLatestInfo()
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("The app info not exists. ")
	}

	return app, err
}
