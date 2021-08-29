package store

import (
	"log"

	"xorm.io/xorm"

	"CheckInAssistant/model"
)

type DBClassRoom struct {
	engine xorm.Interface
}

func NewDBClassRoom(session ...xorm.Interface) *DBClassRoom {
	if len(session) > 0 {
		return &DBClassRoom{engine: session[0]}
	}
	return &DBClassRoom{engine: engine}
}

// Insert insert a class room
func (db *DBClassRoom) Insert(c *model.ClassRoom) error {
	_, e := db.engine.Insert(c)
	if e != nil {
		log.Printf("DBClassRoom Insert failed: %s", e.Error())
	}
	return e
}

// GetByID return a class room by id
func (db *DBClassRoom) GetByID(id string) (bool, *model.ClassRoom, error) {
	classRoom := new(model.ClassRoom)
	has, err := db.engine.ID(id).NoAutoCondition(true).Get(classRoom)
	if err != nil {
		log.Printf("DBClassRoom GetByID failed: %s", err.Error())
	}
	return has, classRoom, err
}
