package models

type OrderItem struct {
	Id        int64   `gorm:"primary_key; auto_increment; not null; index;" json:"id"`
	OrderId   int64   `gorm:"not null; index;" json:"order_id"`
	ProductId int64   `gorm:"not null; index;" json:"product_id"`
	Quantity  int64   `gorm:"not null;" json:"quantity"`
	Price     float64 `gorm:"not null;" json:"price"`
	CreatedAt int64   `gorm:"not null;" json:"created_at"`
	UpdatedAt int64   `gorm:"not null;" json:"updated_at"`

	Product Product `json:"product"`
	Order   Order   `json:"order"`
}
