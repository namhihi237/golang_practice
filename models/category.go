package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

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

func CreateCategory(name string, image string, description string) error {
	// check name is exist
	var c Category
	err := db.Where("name = ?", name).First(&c).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if c.Id > 0 {
		return errors.New("Category name is exist")
	}

	category := &Category{
		Name:        name,
		Image:       image,
		Description: description,
	}
	return db.Create(category).Error
}
