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

type CategoryResponse struct {
	Id          int64      `json:"id"`
	Name        string     `json:"name"`
	Image       string     `json:"image"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func GetCategoryByName(name string) (*CategoryResponse, error) {
	var category CategoryResponse
	err := db.Where("name = ?", name).Limit(1).Find(&category).Error
	if err != nil {
		return nil, err
	}

	if category.Id <= 0 {
		return nil, nil
	}

	return &category, nil
}

func GetCategoryById(id int64) (*CategoryResponse, error) {
	var category CategoryResponse
	err := db.Model(&Category{}).Where("id = ?", id).Limit(1).Find(&category).Error
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

func GetCategories(page int, limit int) ([]CategoryResponse, error) {
	var categories []CategoryResponse
	if page > 0 && limit > 0 {
		db.Model(&Category{}).
			Offset((page - 1) * limit).
			Limit(limit).
			Find(&categories)
	} else {
		db.Model(&Category{}).Find(&categories)
	}

	return categories, nil
}

func CountCategory() int64 {
	var count int64
	db.Model(&Category{}).Count(&count)
	return count
}

func DeleteCategory(id int64) error {
	category := &Category{
		Id: id,
	}
	return db.Delete(category).Error
}
