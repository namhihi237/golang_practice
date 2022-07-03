package models

type CartItem struct {
	Id        int64 `gorm:"primary_key; auto_increment; not null; index;" json:"id"`
	CartId    int64 `gorm:"not null; index;" json:"cart_id"`
	ProductId int64 `gorm:"not null; index;" json:"product_id"`
	Quantity  int64 `gorm:"not null;" json:"quantity"`
	CreatedAt int64 `gorm:"not null;" json:"created_at"`
	UpdatedAt int64 `gorm:"not null;" json:"updated_at"`

	Product Product `json:"product"`
	Cart    Cart    `json:"cart"`
}
