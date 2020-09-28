package model

import "time"

type Class struct {
	ClassID       string `json:"class_id" xorm:"char(36) pk"`
	ClassName     string `json:"class_name" xorm:"varchar(30)"`
	ClassCapacity int    `json:"class_capacity" xorm:"int(3)"`
	UserID        string `json:"user_id" xorm:"char(36) index"`

	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
	DeletedAt time.Time `json:"deleted_at" xorm:"deleted"`
}

type ClassRoom struct {
	RoomID   string `json:"room_id" xorm:"char(36) pk"`
	RoomName string `json:"room_name" xorm:"varchar(30)"`
	ClassID  string `json:"class_id" xorm:"char(36) index"`
	UserID   string `json:"user_id" xorm:"char(36) index"`

	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
	DeletedAt time.Time `json:"deleted_at" xorm:"deleted"`
}
