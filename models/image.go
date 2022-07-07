package models

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	Id        int64      `gorm:"primary_key; auto_increment; not null; index;" json:"id"`
	Url       string     `gorm:"not null; size:255;" json:"url"`
	ProductId int64      `gorm:"not null; index;" json:"product_id"`
	IsDefault bool       `gorm:"default: false;" json:"is_default"`
	CreatedAt *time.Time `gorm:"not null;" json:"created_at"`
	UpdatedAt *time.Time `gorm:"not null;" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func CreateBulkImages(tx *gorm.DB, images []Image, productId int64) error {

	err := tx.CreateInBatches(&images, len(images)).Error
	if err != nil {
		return err
	}
	return nil
}
