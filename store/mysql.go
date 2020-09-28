package store

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/core"
	"xorm.io/xorm"
)

var (
	dataSourceName = ""
	engine         *xorm.Engine
	showSQL        = false
	syncDbModels   = false
)

func initMySQL() {
	var err error
	engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Init mysql connection failed: %s \n", err.Error())
	}
	log.Println("Init mysql connection...")
	engine.SetMapper(core.GonicMapper{})

	engine.ShowSQL(showSQL)
	if syncDbModels {
		for _, dbm := range dbModels {
			exists, errExists := engine.IsTableExist(dbm)
			if errExists != nil {
				log.Printf("Sysnc table select error: %s", errExists.Error())
				continue
			}

			if !exists {
				errCreate := engine.CreateTables(dbm)
				if errCreate != nil {
					log.Printf("Sysnc table create error: %s", errCreate.Error())
					continue
				}
			}

			errSync := engine.Sync2(dbm)
			if errSync != nil {
				log.Printf("Sysnc table sync error: %+v, %s \n", dbm, errSync.Error())
			}
		}
	}
}
