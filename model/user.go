package model

import (
	"time"
)

type User struct {
	UserID   string `json:"user_id" xorm:"char(36) pk"`
	Phone    string `json:"phone" xorm:"notnull"`
	Password string `json:"-" xorm:"notnull"`
	Email    string `json:"email"`
	Role     int    `json:"role" xorm:"tinyint(1)"`
	ClassID  string `json:"class_id" xorm:"char(36) index"`
	RealName string `json:"real_name" xorm:"varchar(20)"`

	Pwd           string `json:"password,omitempty" xorm:"-"`
	CheckInStatus int    `json:"check_in_status" xorm:"-"`

	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"-" xorm:"updated"`
	DeletedAt time.Time `json:"-" xorm:"deleted"`
}
