package models

import "time"

type Category struct {
	Id          int64      `gorm:"primary_key; auto_increment; not null; index;" json:"id"`
	Name        string     `gorm:"not null; size:255; unique;" json:"name"`
	Image       string     `gorm:"size:255;" json:"image"`
	Description string     `gorm:"size:255;" json:"description"`
	CreatedAt   *time.Time `gorm:"not null;" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"not null;" json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`

	Products []Product `json:"products" gorm:"many2many:category_products;"`
}
