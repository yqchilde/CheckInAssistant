package viewmodel

import "time"

type User struct {
	UserID    string `json:"user_id"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Role      int    `json:"role"`
	ClassID   string `json:"class_id"`
	RealName  string `json:"real_name"`
	ClassName string `json:"class_name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
