package store

import "xorm.io/xorm"

var (
	dataSourceName = ""
	engine         *xorm.Engine
	showSQL        = false
	syncDbModels   = false
)

func initMySQL()  {

}