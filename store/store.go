package store

import (
	"CheckInAssistant/model"
	"CheckInAssistant/pkg/setting"
)

var (
	dbModels = []interface{}{
		new(model.User),
		new(model.AppUpgrade),
		new(model.Class),
		new(model.ClassRoom),
		new(model.Sign),
	}
)

func Setup(mysqlInit, redisInit bool) {
	showSQL = setting.MySQLSetting.ShowSQL
	syncDbModels = setting.MySQLSetting.ShowDB

	if setting.MySQLSetting.LinkUrl != "" {
		dataSourceName = setting.MySQLSetting.LinkUrl
	}

	if mysqlInit {
		initMySQL()
	}
}

func Clear() {
	if engine != nil {
		_ = engine.Close()
	}
}
