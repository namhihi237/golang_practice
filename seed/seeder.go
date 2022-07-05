package seed

import (
	"fmt"
	"practice/config"
	"practice/models"
	"practice/pkg/utils"
	"time"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	fmt.Println("Seeding data...")
	env, err := config.GetEnv()
	if err != nil {
		fmt.Println("Error getting env: ", err)
		return
	}
	var adminExist *models.Admin
	// check if admin exists and create

	initUserTypes(db)

	err = db.Where("user_name = ?", env.AdminUsername).First(&adminExist).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {

			hashPassword, err := utils.HashPassword(env.AdminPassword)
			if err != nil {
				fmt.Println(err)
				return
			}

			// select user type admin
			var userType models.UserType
			err = db.Where("name = ?", "admin").First(&userType).Error
			if err != nil {
				fmt.Println("Error getting user type: ", err)
				return
			}

			admin := map[string]interface{}{
				"user_name":    env.AdminUsername,
				"password":     hashPassword,
				"email":        env.AdminEmail,
				"created_at":   time.Now().UTC(),
				"updated_at":   time.Now().UTC(),
				"is_active":    true,
				"user_type_id": userType.Id,
			}

			err = db.Model(&models.Admin{}).Create(admin).Error
			fmt.Println("admin created", err)
			if err != nil {
				fmt.Println("Error creating admin:", err)
			}
		} else {
			fmt.Println("Error get admin:", err)
		}
	}
}

func initUserTypes(db *gorm.DB) {
	fmt.Println("Seeding user types...")
	userTypes := []models.UserType{
		{
			Name: "admin",
		},
		{
			Name: "user",
		},
	}
	for _, userType := range userTypes {
		err := db.Where("name = ?", userType.Name).First(&userType).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				err = db.Create(&userType).Error
				if err != nil {
					fmt.Println("Error creating user type:", err)
				}
			} else {
				fmt.Println("Error get user type:", err)
			}
		}
	}
}
