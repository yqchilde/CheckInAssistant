package store

import (
	"log"

	"xorm.io/xorm"

	"CheckInAssistant/model"
	"CheckInAssistant/viewmodel"
)

type DBClass struct {
	engine xorm.Interface
}

func NewDBClass(session ...xorm.Interface) *DBClass {
	if len(session) > 0 {
		return &DBClass{engine: session[0]}
	}
	return &DBClass{engine: engine}
}

// Insert insert a user
func (db *DBClass) Insert(c *model.Class) error {
	_, e := db.engine.Insert(c)
	if e != nil {
		log.Printf("DBClass Insert failed: %s", e.Error())
	}
	return e
}

// GetByID return a user by id
func (db *DBClass) GetByID(id string) (bool, *model.Class, error) {
	class := new(model.Class)
	has, err := db.engine.ID(id).NoAutoCondition(true).Get(class)
	if err != nil {
		log.Printf("DBClass GetByID failed: %s", err.Error())
	}
	return has, class, err
}

// GetByUserIDAndType get class list by user id and type
func (db *DBClass) GetByUserIDAndType(userID, tabIndex string) ([]*model.Class, error) {
	var classList []*model.Class

	switch tabIndex {
	case model.TabTypeLeft:
		err := db.engine.Where("creator = ?", userID).Desc("created_at").Find(&classList)
		if err != nil {
			log.Printf("DBClass GetByUserIDAndType failed: %s", err.Error())
		}
		return classList, err
	case model.TabTypeRight:
		query := `SELECT c.* FROM class c
				  INNER JOIN connection ct ON c.class_id = ct.class_id AND ct.user_id = ?
					AND (ct.deleted_at IS NULL OR ct.deleted_at = '0001-01-01 00:00:00')
				  WHERE c.creator <> ? AND (c.deleted_at IS NULL OR c.deleted_at = '0001-01-01 00:00:00')
				  ORDER BY c.created_at DESC`
		err := db.engine.SQL(query, userID, userID).Find(&classList)
		if err != nil {
			log.Printf("DBClass GetByUserIDAndType failed: %s", err.Error())
		}
		return classList, err
	}
	return nil, nil
}

// UpdateClass update class info
func (db *DBClass) UpdateClass(class *model.Class) error {
	_, err := db.engine.ID(class.ClassID).Cols("class_name", "class_capacity").Update(class)
	if err != nil {
		log.Printf("DBClass UpdateClass failed: %s", err.Error())
	}
	return err
}

// DeleteClass delete class
func (db *DBClass) DeleteClass(classID string) error {
	_, err := db.engine.ID(classID).Delete(new(model.Class))
	if err != nil {
		log.Printf("DBClass DeleteClass failed: %s", err.Error())
	}
	return err
}

// GetClassRoomList return class room list
func (db *DBClass) GetClassRoomList(roomID string) ([]*model.Class, error) {
	var cmList []*model.Class
	query := `SELECT c.* FROM class c
			  INNER JOIN room r ON r.room_id = ?
			  WHERE (c.deleted_at IS NULL OR c.deleted_at = "0001-01-01 00:00:00")`
	err := db.engine.SQL(query, roomID).Find(&cmList)
	if err != nil {
		log.Printf("DBClassRoom GetClassRoomList failed: %s", err.Error())
	}
	return cmList, err
}

// GetClassUserList get class user list by class id
func (db *DBClass) GetClassUserList(classID string) ([]*viewmodel.User, error) {
	var users []*viewmodel.User
	query := `SELECT c.class_name, u.* FROM user u 
			  INNER JOIN connection ct ON u.user_id = ct.user_id AND (ct.deleted_at IS NULL OR ct.deleted_at = '0001-01-01 00:00:00')
			  INNER JOIN class c ON c.class_id = ct.class_id AND c.class_id = ?
			  WHERE (u.deleted_at IS NULL OR u.deleted_at = '0001-01-01 00:00:00')`
	err := db.engine.SQL(query, classID).Find(&users)
	if err != nil {
		log.Printf("DBClassRoom GetClassUserList failed: %s", err.Error())
	}
	return users, err
}
