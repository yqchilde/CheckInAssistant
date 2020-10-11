package model

import "time"

type Class struct {
	ClassID       string `json:"class_id" xorm:"char(36) pk"`
	ClassName     string `json:"class_name" xorm:"varchar(30)"`
	ClassCapacity int    `json:"class_capacity" xorm:"int(3)"`
	Creator       string `json:"creator" xorm:"char(36) index"`

	RoomID              string `json:"room_id" xorm:"-"`
	CurrentPeopleNumber int64  `json:"current_people_number" xorm:"-"`

	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"-" xorm:"updated"`
	DeletedAt time.Time `json:"-" xorm:"deleted"`
}

type ClassRoom struct {
	ClassRoomID string `json:"class_room_id" xorm:"char(36) pk"`
	ClassID     string `json:"class_id" xorm:"char(36) index"`
	RoomID      string `json:"room_id" xorm:"char(36) index"`

	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"-" xorm:"updated"`
	DeletedAt time.Time `json:"-" xorm:"deleted"`
}
