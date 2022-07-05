package models

import (
	"os"
	"practice/config"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"log"

	"gorm.io/driver/mysql"
)

var db *gorm.DB

func SetUp() {
	var err error
	var env config.Env

	env, err = config.GetEnv()
	if err != nil {
		log.Fatal(err)
	}

	var loggerConfig logger.Interface
	if env.Go_env != "production" {
		loggerConfig = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				Colorful: false,
				LogLevel: logger.Info | logger.Warn | logger.Error,
			},
		)
	}

	db, err = gorm.Open(mysql.Open(env.DbUrl), &gorm.Config{
		Logger: loggerConfig,
	})

	if err != nil {
		log.Fatal("Can't connect to database: ", err)
	}

	log.Printf("Connected to database: %s", env.DbUrl)

	// migrate the schema
	db.AutoMigrate(&UserType{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&CategoryProduct{})
	db.AutoMigrate(&Image{})
	db.AutoMigrate(&Cart{})
	db.AutoMigrate(&CartItem{})
	db.AutoMigrate(&Order{})
	db.AutoMigrate(&OrderItem{})
	db.AutoMigrate(&Admin{})

	// migrate the schema
	// db.Migrator().RenameColumn(&Admin{}, "username", "user_name")

}

func GetDb() *gorm.DB {
	return db
}
