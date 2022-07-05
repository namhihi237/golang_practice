package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id         int64      `gorm:"primary_key; auto_increment; not null; index;" json:"id"`
	UserName   string     `gorm:"not null; size:255; unique;" json:"user_name"`
	Password   string     `gorm:"not null; size:255;" json:"password"`
	Email      string     `gorm:"not null; size:255; unique;" json:"email"`
	FullName   string     `gorm:"default: null; size:255;" json:"full_name"`
	Phone      string     `gorm:"size:20; default: null;" json:"phone"`
	Address    string     `gorm:"size:255; default: null;" json:"address"`
	Gender     string     `gorm:"size:10; default: null;" json:"gender"`
	Birthday   *time.Time `gorm:"default: null;" json:"birthday"`
	IsActive   bool       `gorm:"default: false;" json:"is_active"`
	IsBlocked  bool       `gorm:"default:false;" json:"is_blocked"`
	UserTypeId int        `json:"user_type_id"`
	CreatedAt  *time.Time `gorm:"not null;" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"not null;" json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`

	Cart     Cart     `json:"cart"`
	UserType UserType `json:"user_type"`
	Orders   []Order  `json:"orders"`
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := db.Joins("UserType").Where("user_name = ?", username).First(&user).Error
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

	//get user type
	userType, err := GetUserTypeByName("user")
	if err != nil {
		return err
	}

	if userType == nil {
		return errors.New("user type not found")
	}

	user := User{
		UserName:   username,
		Password:   password,
		Email:      email,
		UserTypeId: int(userType.Id),
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func GetUserById(id int64) (*User, error) {
	var user User
	err := db.Joins("UserType").Where("users.id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if user.Id <= 0 {
		return nil, nil
	}

	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if user.Id <= 0 {
		return nil, nil
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

func UpdateUser(id int64, data interface{}) error {
	if err := db.Model(&User{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
