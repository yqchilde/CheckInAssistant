package store

import (
	"log"

	"xorm.io/xorm"

	"CheckInAssistant/model"
)

type DBRoom struct {
	engine xorm.Interface
}

func NewDBRoom(session ...xorm.Interface) *DBRoom {
	if len(session) > 0 {
		return &DBRoom{engine: session[0]}
	}
	return &DBRoom{engine: engine}
}

// Insert insert a room
func (db *DBRoom) Insert(u *model.Room) error {
	_, e := db.engine.Insert(u)
	if e != nil {
		log.Printf("DBRoom Insert failed: %s", e.Error())
	}
	return e
}

// GetByID return a room by id
func (db *DBRoom) GetByID(id string) (bool, *model.Room, error) {
	room := new(model.Room)
	has, err := db.engine.ID(id).NoAutoCondition(true).Get(room)
	if err != nil {
		log.Printf("DBRoom GetByID failed: %s", err.Error())
	}
	return has, room, err
}

// GetByUserIDAndType get room list by user id and type
func (db *DBRoom) GetByUserIDAndType(userID, tabIndex string) ([]*model.Room, error) {
	var roomList []*model.Room

	switch tabIndex {
	case model.TabTypeLeft:
		err := db.engine.Where("creator = ?", userID).Find(&roomList)
		if err != nil {
			log.Printf("DBRoom GetByUserIDAndType failed: %s", err.Error())
		}
		return roomList, err
	case model.TabTypeRight:
		query := `SELECT r.*, cr.class_id FROM room r
			      INNER JOIN class_room cr ON r.room_id = cr.room_id
				  INNER JOIN connection c ON c.user_id = ? AND cr.class_id = c.class_id
				  WHERE r.creator <> ? AND (r.deleted_at IS NULL OR r.deleted_at = '0001-01-01 00:00:00')
				  ORDER BY cr.created_at DESC`
		err := db.engine.SQL(query, userID, userID).Find(&roomList)
		if err != nil {
			log.Printf("DBRoom GetByUserIDAndType failed: %s", err.Error())
		}
		return roomList, err
	}
	return nil, nil
}

// DeleteClassRoom delete class room by room id
func (db *DBRoom) DeleteRoom(roomID string) error {
	_, err := db.engine.Where("room_id = ?", roomID).Delete(new(model.Room))
	if err != nil {
		log.Printf("DBRoom DeleteRoom failed: %s", err.Error())
	}
	return err
}
