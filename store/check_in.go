package store

import (
	"log"

	"xorm.io/xorm"

	"CheckInAssistant/model"
	"CheckInAssistant/viewmodel"
)

type DBCheckIn struct {
	engine xorm.Interface
}

func NewDBCheckIn(session ...xorm.Interface) *DBCheckIn {
	if len(session) > 0 {
		return &DBCheckIn{engine: session[0]}
	}
	return &DBCheckIn{engine: engine}
}

// Insert insert a checkin
func (db *DBCheckIn) Insert(c *model.CheckIn) error {
	_, e := db.engine.Insert(c)
	if e != nil {
		log.Printf("DBCheckIn Insert failed: %s", e.Error())
	}
	return e
}

// GetByID return a checkin by id
func (db *DBCheckIn) GetByID(id string) (bool, *model.CheckIn, error) {
	checkIn := new(model.CheckIn)
	has, err := db.engine.ID(id).NoAutoCondition(true).Get(checkIn)
	if err != nil {
		log.Printf("DBCheckIn GetByID failed: %s", err.Error())
	}
	return has, checkIn, err
}

// GetCheckInDetail get checkin detail
func (db *DBCheckIn) GetCheckInDetail(checkInID string) (*viewmodel.CheckIn, error) {
	checkIn := new(viewmodel.CheckIn)
	query := `SELECT ci.check_in_id, ci.begin_time, ci.end_time, ci.check_in_way, ci.check_in_code, 
				ci.longitude, ci.latitude, r.room_id, r.room_name, u.user_id AS creator, 
				u.real_name AS creator_name FROM check_in ci 
			  INNER JOIN room r ON ci.room_id = r.room_id AND ci.check_in_id = ?
			  INNER JOIN user u ON ci.creator = u.user_id`
	_, err := db.engine.SQL(query, checkInID).Get(checkIn)
	if err != nil {
		log.Printf("DBCheckIn GetCheckInDetail failed: %s", err.Error())
	}
	return checkIn, err
}

// GetByUserIDAndType get check in list by user id and type
func (db *DBCheckIn) GetByUserIDAndType(userID, tabIndex string) ([]*viewmodel.CheckIn, error) {
	var checkInList []*viewmodel.CheckIn

	switch tabIndex {
	case model.TabTypeLeft:
		query := `SELECT ci.check_in_id, ci.begin_time, ci.end_time, ci.check_in_way, ci.check_in_code, 
				    ci.longitude, ci.latitude, r.room_id, r.room_name, u.user_id AS creator, 
					u.real_name AS creator_name FROM check_in ci 
			      INNER JOIN room r ON ci.room_id = r.room_id
			      INNER JOIN user u ON ci.creator = u.user_id AND ci.creator = ?
				  ORDER BY ci.created_at DESC`
		err := db.engine.SQL(query, userID).Find(&checkInList)
		if err != nil {
			log.Printf("DBCheckIn GetByUserIDAndType failed: %s", err.Error())
		}
		return checkInList, err
	case model.TabTypeRight:
		query := `SELECT ci.*, t.room_name FROM check_in ci 
				  INNER JOIN (
					SELECT r.room_id, r.room_name FROM room r
					INNER JOIN class_room cr ON r.room_id = cr.room_id
					INNER JOIN connection c ON c.user_id = ? AND cr.class_id = c.class_id
					WHERE r.creator <> ? AND (r.deleted_at IS NULL OR r.deleted_at = '0001-01-01 00:00:00')
				  ) t ON ci.room_id = t.room_id
				  WHERE (ci.deleted_at IS NULL OR ci.deleted_at = '0001-01-01 00:00:00')
				  ORDER BY ci.created_at DESC`
		err := db.engine.SQL(query, userID, userID).Find(&checkInList)
		if err != nil {
			log.Printf("DBCheckIn GetByUserIDAndType failed: %s", err.Error())
		}
		return checkInList, err
	}
	return nil, nil
}
