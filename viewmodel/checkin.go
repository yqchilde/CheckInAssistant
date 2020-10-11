package viewmodel

import "time"

type CheckIn struct {
	CheckInID    string    `json:"check_in_id"`
	RoomID       string    `json:"room_id"`
	RoomName     string    `json:"room_name"`
	Creator      string    `json:"creator"`
	CreatorName  string    `json:"creator_name"`
	CheckInWay   int       `json:"check_in_way"`
	BeginTime    time.Time `json:"begin_time"`
	EndTime      time.Time `json:"end_time"`
	CheckInCode  string    `json:"check_in_code,omitempty"`
	Longitude    string    `json:"longitude,omitempty"`
	Latitude     string    `json:"latitude,omitempty"`
	PeopleTotal  int       `json:"people_total"`
	CheckInTotal int       `json:"check_in_total"`
	Status       int       `json:"status"`
}
