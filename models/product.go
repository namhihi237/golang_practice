package models

import "time"

type Product struct {
	Id          int64      `gorm:"primary_key; auto_increment; not null; index;" json:"id"`
	Name        string     `gorm:"not null; size:255;" json:"name"`
	Price       float64    `gorm:"not null;" json:"price"`
	Amount      int64      `gorm:"not null;" json:"amount"`
	Description string     `gorm:"size:255;" json:"description"`
	CreatedAt   *time.Time `gorm:"not null;" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"not null;" json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`

	Categories []Category `json:"categories" gorm:"many2many:category_products;"`
	Images     []Image    `json:"images"`
}

func CreateProduct(CategoryIds []int64, ImageUrls []string, product *Product) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&product).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := CreateCategoryProducts(tx, CategoryIds, product.Id); err != nil {
		tx.Rollback()
		return err
	}

	images := make([]Image, len(ImageUrls))
	for i, url := range ImageUrls {
		images[i] = Image{
			Url:       url,
			ProductId: product.Id,
			IsDefault: i == 0,
		}
	}

	if err := CreateBulkImages(tx, images, product.Id); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
