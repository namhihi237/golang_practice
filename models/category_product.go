package models

type CategoryProduct struct {
	CategoryId int64 `gorm:"not null; index;" json:"category_id"`
	ProductId  int64 `gorm:"not null; index;" json:"product_id"`
}
