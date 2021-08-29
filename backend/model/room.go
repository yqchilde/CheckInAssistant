package model

import "time"

type Room struct {
	RoomID   string `json:"room_id" xorm:"char(36) pk"`
	RoomName string `json:"room_name" xorm:"varchar(30)"`
	Creator  string `json:"creator" xorm:"char(36) index"`

	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"-" xorm:"updated"`
	DeletedAt time.Time `json:"-" xorm:"deleted"`
}
