package store

import (
	"log"
	"strings"

	"xorm.io/xorm"

	"CheckInAssistant/model"
)

type DBUser struct {
	engine xorm.Interface
}

func NewDBUser(session ...xorm.Interface) *DBUser {
	if len(session) > 0 {
		return &DBUser{engine: session[0]}
	}
	return &DBUser{engine: engine}
}

// Insert insert a user
func (db *DBUser) Insert(u *model.User) error {
	_, e := db.engine.Insert(u)
	if e != nil {
		log.Printf("DBUser Insert failed: %s", e.Error())
	}
	return e
}

// GetByID return a user by id
func (db *DBUser) GetByID(id string) (bool, *model.User, error) {
	user := new(model.User)
	has, err := db.engine.ID(id).NoAutoCondition(true).Get(user)
	if err != nil {
		log.Printf("DBUser GetByID failed: %s", err.Error())
	}
	return has, user, err
}

// GetByPhone return a user by id
func (db *DBUser) GetByPhone(phone string) (bool, *model.User, error) {
	user := new(model.User)
	has, err := db.engine.Where("phone = ?", phone).Get(user)
	if err != nil {
		log.Printf("DBUser GetByPhone failed: %s", err.Error())
	}
	return has, user, err
}

// GetByClassIDList return user by class id list
func (db *DBUser) GetByClassIDList(classIDList []string) ([]*model.User, error) {
	var (
		users  []*model.User
		idList = `"` + strings.Join(classIDList, `","`) + `"`
	)

	query := `SELECT u.* FROM user u
			  INNER JOIN connection c ON u.user_id = c.user_id 
				AND c.class_id IN (` + idList + `)
			  WHERE (u.deleted_at IS NULL OR u.deleted_at = "0001-01-01 00:00:00")
			  ORDER BY u.created_at DESC`
	err := db.engine.SQL(query).Find(&users)
	if err != nil {
		log.Printf("DBUser GetByClassID failed: %s", err.Error())
	}
	return users, err
}

// UpdateUserInfo update user info
func (db *DBUser) UpdateUserInfo(user *model.User) error {
	_, err := db.engine.ID(user.UserID).Cols("password", "email", "class_id", "real_name").Update(user)
	if err != nil {
		log.Printf("DBUser UpdateUserInfo failed: %s", err.Error())
	}
	return err
}
