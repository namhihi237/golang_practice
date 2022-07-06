package models

import (
	"errors"
	"time"
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

func GetCategoryByName(name string) (*Category, error) {
	var category Category
	err := db.Where("name = ?", name).Limit(1).Find(&category).Error
	if err != nil {
		return nil, err
	}

	if category.Id <= 0 {
		return nil, nil
	}

	return &category, nil
}

func GetCategoryById(id int64) (*Category, error) {
	var category Category
	err := db.Where("id = ?", id).Limit(1).Find(&category).Error
	if err != nil {
		return nil, err
	}

	if category.Id <= 0 {
		return nil, nil
	}

	return &category, nil
}

func CreateCategory(name string, image string, description string) error {
	// check name is exist
	c, err := GetCategoryByName(name)
	if err != nil {
		return err
	}

	if c == nil {
		category := &Category{
			Name:        name,
			Image:       image,
			Description: description,
		}
		return db.Create(category).Error
	}

	return errors.New("Category name is exist")
}

func UpdateCategory(id int64, name string, image string, description string) error {
	// check id is exist
	c, err := GetCategoryById(id)
	if err != nil {
		return err
	}

	if c == nil {
		return errors.New("Category not found")
	}

	// check name is exist
	c, err = GetCategoryByName(name)
	if err != nil {
		return err
	}

	if c != nil {
		return errors.New("Category name is exist")
	}

	category := &Category{
		Id:          id,
		Name:        name,
		Image:       image,
		Description: description,
	}
	if err := db.Model(&Category{}).Where("id = ?", id).Updates(category).Error; err != nil {
		return err
	}

	return nil
}
