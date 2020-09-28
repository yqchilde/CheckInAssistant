package model

import "time"

type Sign struct {
	SignID    string    `json:"sign_id" xorm:"char(36) pk"`
	UserID    string    `json:"user_id" xorm:"char(36) index"`
	RoomID    string    `json:"room_id" xorm:"char(36) index"`
	SignInWay int       `json:"sign_in_way" xorm:"tinyint(1)"`
	BeginTime time.Time `json:"begin_time"`
	SignTime  time.Time `json:"sign_time"`

	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
	DeletedAt time.Time `json:"deleted_at" xorm:"deleted"`
}
