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
	autoMigrate()

	// change column name
	renameColumn(&Admin{}, "username", "user_name")
	AddColumn(&CategoryProduct{}, "is_active")

	if db.Migrator().HasIndex(&Image{}, "url") {
		db.Migrator().DropIndex(&Image{}, "url")
	}
	if db.Migrator().HasIndex(&Product{}, "name") {
		db.Migrator().DropIndex(&Product{}, "name")
	}
}

func GetDb() *gorm.DB {
	return db
}

func autoMigrate() {
	checkTableExistAndMigrate(&UserType{})
	checkTableExistAndMigrate(&User{})
	checkTableExistAndMigrate(&Product{})
	checkTableExistAndMigrate(&Category{})
	checkTableExistAndMigrate(&CategoryProduct{})
	checkTableExistAndMigrate(&Image{})
	checkTableExistAndMigrate(&Cart{})
	checkTableExistAndMigrate(&CartItem{})
	checkTableExistAndMigrate(&Order{})
	checkTableExistAndMigrate(&OrderItem{})
	checkTableExistAndMigrate(&Admin{})
}

func checkTableExistAndMigrate(dts interface{}) {
	if !db.Migrator().HasTable(dts) {
		db.AutoMigrate(dts)
	}
}

// Note: when update name: run this func and update models
func renameColumn(dts interface{}, oldName string, newName string) {
	if db.Migrator().HasTable(dts) {
		if db.Migrator().HasColumn(dts, oldName) {
			db.Migrator().RenameColumn(dts, oldName, newName)
		}
	}
}

func AddColumn(dts interface{}, columnName string) {
	if db.Migrator().HasTable(dts) {
		if !db.Migrator().HasColumn(dts, columnName) {
			db.Migrator().AddColumn(dts, columnName)
		}
	}
}
