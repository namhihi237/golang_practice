package models

import "time"

type User struct {
	Id        int64      `gorm:"primary_key; auto_increment; not null; index;" json:"id"`
	UserName  string     `gorm:"not null; size:255; unique;" json:"user_name"`
	Password  string     `gorm:"not null; size:255;" json:"password"`
	Email     string     `gorm:"not null; size:255; unique;" json:"email"`
	FullName  string     `gorm:"not null; size:255;" json:"full_name"`
	Phone     string     `gorm:"size:20;" json:"phone"`
	Address   string     `gorm:"size:255;" json:"address"`
	Gender    string     `gorm:"size:10;" json:"gender"`
	IsActive  bool       `gorm:"default: true;" json:"is_active"`
	IsBlocked bool       `gorm:"default:false;" json:"is_blocked"`
	CreatedAt *time.Time `gorm:"not null;" json:"created_at"`
	UpdatedAt *time.Time `gorm:"not null;" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	Cart   Cart    `json:"cart"`
	Orders []Order `json:"orders"`
}
