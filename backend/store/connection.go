package store

import (
	"log"

	"xorm.io/xorm"

	"CheckInAssistant/model"
)

type DBConnection struct {
	engine xorm.Interface
}

func NewDBConnection(session ...xorm.Interface) *DBConnection {
	if len(session) > 0 {
		return &DBConnection{engine: session[0]}
	}
	return &DBConnection{engine: engine}
}

// Insert insert a connection
func (db *DBConnection) Insert(c *model.Connection) error {
	_, e := db.engine.Insert(c)
	if e != nil {
		log.Printf("DBConnection Insert failed: %s", e.Error())
	}
	return e
}

// DeleteByClassIDAndUserID ...
func (db *DBConnection) DeleteByClassIDAndUserID(classID, userID string) error {
	_, e := db.engine.Where("class_id = ? AND user_id = ?", classID, userID).Delete(new(model.Connection))
	if e != nil {
		log.Printf("DBConnection DeleteByClassIDAndUserID failed: %s", e.Error())
	}
	return e
}

// GetByClassID return class list
func (db *DBConnection) GetByClassID(classID string) (int64, error) {
	count, e := db.engine.Where("class_id = ?", classID).Count(new(model.Connection))
	if e != nil {
		log.Printf("DBConnection GetByClassID failed: %s", e.Error())
	}
	return count, e
}
