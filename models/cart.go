package models

import "time"

type Cart struct {
	Id        int64      `gorm:"primary_key; auto_increment; not null; index;" json:"id"`
	UserId    int64      `gorm:"not null; index;" json:"user_id"`
	CreatedAt *time.Time `gorm:"not null;" json:"created_at"`
	UpdatedAt *time.Time `gorm:"not null;" json:"updated_at"`

	CartItems []CartItem `json:"cart_items"`
}
