package model

import "time"

type Connection struct {
	ID int `json:"id" xorm:"pk autoincr"`

	ClassID string `json:"class_id" xorm:"char(36) index"`
	UserID string `json:"user_id" xorm:"char(36) index"`

	CreatedAt time.Time `json:"created_at" xorm:"created"`
	DeletedAt time.Time `json:"-" xorm:"deleted"`
}
