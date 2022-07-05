package models

import (
	"time"

	"gorm.io/gorm"
)

type UserType struct {
	Id        int64      `gorm:"primary_key; auto_increment; not null; index;" json:"id"`
	Name      string     `gorm:"not null; size:255; unique;" json:"name"`
	CreatedAt *time.Time `gorm:"not null;" json:"created_at"`
	UpdatedAt *time.Time `gorm:"not null;" json:"updated_at"`
}

func GetUserTypeByName(name string) (*UserType, error) {
	var userType UserType
	err := db.Where("name = ?", name).First(&userType).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if userType.Id <= 0 {
		return nil, nil
	}

	return &userType, nil
}
