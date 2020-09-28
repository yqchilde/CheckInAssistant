package model

import "time"

type User struct {
	UserID   string `json:"user_id" xorm:"char(36) pk"`
	Username string `json:"username"`
	Phone    string `json:"phone" xorm:"notnull"`
	Password string `json:"-" xorm:"notnull"`
	Email    string `json:"email"`
	ClassID  string `json:"class_id" xorm:"char(36)" index`
	RealName string `json:"real_name" xorm:"varchar(20)"`

	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
	DeletedAt time.Time `json:"-" xorm:"deleted"`
}
