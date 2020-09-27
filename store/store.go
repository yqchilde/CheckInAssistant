package store

import "CheckInAssistant/pkg/setting"

var (
	dbModels = []interface{}{
		//new(),
	}
)

func InitStore(mysqlInit, redisInit bool) {
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

	}
}
