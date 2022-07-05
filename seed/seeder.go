package seed

import (
	"fmt"
	"log"
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
		log.Println("Error getting env: ", err)
		return
	}
	var adminExist *models.Admin
	// check if admin exists and create

	err = db.Where("user_name = ?", env.AdminUsername).First(&adminExist).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {

			hashPassword, err := utils.HashPassword(env.AdminPassword)
			if err != nil {
				log.Println(err)
				return
			}
			admin := map[string]interface{}{
				"user_name":  env.AdminUsername,
				"password":   hashPassword,
				"email":      env.AdminEmail,
				"created_at": time.Now().UTC(),
				"updated_at": time.Now().UTC(),
				"is_active":  true,
			}

			err = db.Model(&models.Admin{}).Create(admin).Error
			fmt.Println("admin created", err)
			if err != nil {
				log.Println("Error creating admin:", err)
			}
		} else {
			log.Println("Error get admin:", err)
		}
	}
}
