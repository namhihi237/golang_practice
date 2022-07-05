package models

type Order struct {
	Id            int64   `gorm:"primary_key; auto_increment; not null; index;" json:"id"`
	UserId        int64   `gorm:"not null; index;" json:"user_id"`
	Total         float64 `gorm:"not null;" json:"total"`
	Status_abc    string  `gorm:"not null; size:20;" json:"status"`
	PaymentMethod string  `gorm:"not null; size:20; column:payment_methodssss" json:"payment_method"`
	CreatedAt     int64   `gorm:"not null;" json:"created_at"`
	UpdatedAt     int64   `gorm:"not null;" json:"updated_at"`

	OrderItems []OrderItem `json:"order_items"`
	User       User        `json:"user"`
}
