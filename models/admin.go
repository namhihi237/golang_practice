package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	Id        int64      `gorm:"primary_key; auto_increment; not null; index;" json:"id"`
	UserName  string     `gorm:"not null; size:255; unique;" json:"user_name"`
	Password  string     `gorm:"not null; size:255;" json:"password"`
	Email     string     `gorm:"not null; size:255; unique;" json:"email"`
	IsActive  bool       `gorm:"not null; default: true;" json:"is_active"`
	CreatedAt *time.Time `gorm:"not null; default current_timestamp;" json:"created_at"`
	UpdatedAt *time.Time `gorm:"not null; default current_timestamp;" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func GetAdminByUserName(userName string) (*Admin, error) {
	var admin Admin
	err := db.Where("user_name = ?", userName).First(&admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &admin, nil
}
