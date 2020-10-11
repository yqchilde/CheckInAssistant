package model

import "time"

type AppUpgrade struct {
	ID        int    `json:"id" xorm:"pk autoincr"`              // App upgrade id
	Version   string `json:"version" xorm:"varchar(10)"`         // App version number
	Url       string `json:"url"`                                // App down url
	Notes     string `json:"notes"`                              // App upgrade notes
	ShareDown string `json:"share_down"`                         // App share down url to friends
	Status    int    `json:"status" xorm:"tinyint(1) default 0"` // status[0:not online; 1: online]

	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"-" xorm:"updated"`
	DeletedAt time.Time `json:"-" xorm:"deleted"`
}
