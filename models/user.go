package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        int64      `gorm:"primary_key; auto_increment; not null; index;" json:"id"`
	UserName  string     `gorm:"not null; size:255; unique;" json:"user_name"`
	Password  string     `gorm:"not null; size:255;" json:"password"`
	Email     string     `gorm:"not null; size:255; unique;" json:"email"`
	FullName  string     `gorm:"default: null; size:255;" json:"full_name"`
	Phone     string     `gorm:"size:20; default: null;" json:"phone"`
	Address   string     `gorm:"size:255; default: null;" json:"address"`
	Gender    string     `gorm:"size:10; default: null;" json:"gender"`
	IsActive  bool       `gorm:"default: false;" json:"is_active"`
	IsBlocked bool       `gorm:"default:false;" json:"is_blocked"`
	CreatedAt *time.Time `gorm:"not null;" json:"created_at"`
	UpdatedAt *time.Time `gorm:"not null;" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	Cart   Cart    `json:"cart"`
	Orders []Order `json:"orders"`
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := db.Where("user_name = ?", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

func CheckUserByUsernameOrEmail(username string, email string) (bool, error) {
	var user User
	err := db.Where("user_name = ? or email = ?", username, email).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.Id > 0 {
		return true, nil
	}

	return false, nil
}

func CreateUser(username string, password string, email string) error {
	user := User{
		UserName: username,
		Password: password,
		Email:    email,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func GetUserById(id int64) (*User, error) {
	var user User
	var selectStr = "user_name, email, full_name, phone, address, gender"
	err := db.Select(selectStr).Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

func ActiveAccount(id int64) error {
	user := User{
		Id: id,
	}

	if err := db.Model(&user).Update("is_active", true).Error; err != nil {
		return err
	}

	return nil
}
