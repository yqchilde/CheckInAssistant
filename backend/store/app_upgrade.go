package store

import (
	"log"

	"xorm.io/xorm"

	"CheckInAssistant/model"
)

type DBApp struct {
	engine xorm.Interface
}

func NewDBApp(session ...xorm.Interface) *DBApp {
	if len(session) > 0 {
		return &DBApp{engine: session[0]}
	}
	return &DBApp{engine: engine}
}

// GetByID return a user by id
func (db *DBApp) GetLatestInfo() (bool, *model.AppUpgrade, error) {
	app := new(model.AppUpgrade)
	has, err := db.engine.Desc("created_at").NoAutoCondition(true).Get(app)
	if err != nil {
		log.Printf("DBApp GetLatestInfo failed: %s", err.Error())
	}
	return has, app, err
}
