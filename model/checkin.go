package model

import "time"

type CheckIn struct {
	CheckInID   string    `json:"check_in_id" xorm:"char(36) pk"`
	Creator     string    `json:"creator" xorm:"char(36) index"`
	RoomID      string    `json:"room_id" xorm:"char(36) index"`
	CheckInWay  int       `json:"check_in_way" xorm:"tinyint(1)"`
	BeginTime   time.Time `json:"begin_time"`
	EndTime     time.Time `json:"end_time"`
	Duration    int       `json:"duration" xorm:"-"`
	CheckInCode string    `json:"check_in_code,omitempty" xorm:"char(4)"`
	Longitude   string    `json:"longitude,omitempty" xorm:"varchar(20)"`
	Latitude    string    `json:"latitude,omitempty" xorm:"varchar(20)"`

	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"-" xorm:"updated"`
	DeletedAt time.Time `json:"-" xorm:"deleted"`
}

type CheckInItem struct {
	ID        int    `json:"id" xorm:"pk autoincr"`
	CheckInID string `json:"check_in_id" xorm:"char(36) index"`
	UserID    string `json:"user_id" xorm:"char(36) index"`

	CreatedAt time.Time `json:"created_at" xorm:"created"`
	DeletedAt time.Time `json:"-" xorm:"deleted"`
}
