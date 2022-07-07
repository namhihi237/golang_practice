package models

import "gorm.io/gorm"

type CategoryProduct struct {
	CategoryId int64 `gorm:"not null; index;" json:"category_id"`
	ProductId  int64 `gorm:"not null; index;" json:"product_id"`
	IsActive   bool  `gorm:"not null; default: true" json:"is_active"`

	Product Product `json:"product"`
}

func CountProductByCategory(categoryId int64) (int64, error) {
	var count int64
	err := db.Model(&CategoryProduct{}).Where("category_id = ? and is_active = true", categoryId).Count(&count).Error
	return count, err
}

func CreateCategoryProducts(tx *gorm.DB, categoryIds []int64, product int64) error {
	var categoryProducts []CategoryProduct
	for _, categoryId := range categoryIds {
		categoryProducts = append(categoryProducts, CategoryProduct{
			CategoryId: categoryId,
			ProductId:  product,
			IsActive:   true,
		})
	}

	return tx.CreateInBatches(&categoryProducts, len(categoryProducts)).Error
}
