package models

type Admin struct {
	Id        int64  `gorm:"primary_key; auto_increment; not null; index;" json:"id"`
	Username  string `gorm:"not null; size:255; unique;" json:"username"`
	Password  string `gorm:"not null; size:255;" json:"password"`
	Email     string `gorm:"not null; size:255; unique;" json:"email"`
	IsActive  bool   `gorm:"not null; default: true;" json:"is_active"`
	CreatedAt int64  `gorm:"not null;" json:"created_at"`
	UpdatedAt int64  `gorm:"not null;" json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
}
