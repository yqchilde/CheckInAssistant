package store

import (
	"log"

	"xorm.io/xorm"

	"CheckInAssistant/model"
)

type DBCheckInItem struct {
	engine xorm.Interface
}

func NewDBCheckInItem(session ...xorm.Interface) *DBCheckInItem {
	if len(session) > 0 {
		return &DBCheckInItem{engine: session[0]}
	}
	return &DBCheckInItem{engine: engine}
}

// Insert insert a check in item
func (db *DBCheckInItem) Insert(c *model.CheckInItem) error {
	_, e := db.engine.Insert(c)
	if e != nil {
		log.Printf("DBCheckInItem Insert failed: %s", e.Error())
	}
	return e
}

// GetByCheckInIDAndUserID ...
func (db *DBCheckInItem) GetByCheckInIDAndUserID(checkInID, userID string) (bool, *model.CheckInItem, error) {
	checkInItem := new(model.CheckInItem)
	has, e := db.engine.Where("check_in_id = ? AND user_id = ?", checkInID, userID).Desc("created_at").Get(checkInItem)
	if e != nil {
		log.Printf("DBCheckInItem GetByCheckInIDAndUserID failed: %s", e.Error())
	}

	return has, checkInItem, e
}

// GetByCheckInID return check in item list by check in id
func (db *DBCheckInItem) GetByCheckInID(checkInID string) ([]*model.CheckInItem, error) {
	var checkInItem []*model.CheckInItem
	e := db.engine.Where("check_in_id = ?", checkInID).Find(&checkInItem)
	if e != nil {
		log.Printf("DBCheckInItem GetByCheckInID failed: %s", e.Error())
	}
	return checkInItem, e
}
