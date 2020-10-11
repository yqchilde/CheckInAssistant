package viewmodel

import (
	"time"

	"CheckInAssistant/model"
)

type ClassRoom struct {
	RoomID      string   `json:"room_id"`
	RoomName    string   `json:"room_name"`
	Creator     string   `json:"creator"`
	ClassIDList []string `json:"class_id_list,omitempty"`

	ClassList   []*model.Class `json:"class_list,omitempty"`
	ClassMember []*model.User  `json:"class_member,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
}
