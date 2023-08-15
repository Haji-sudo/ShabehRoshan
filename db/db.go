package db

import (
	"fmt"
	"log"
	"os"

	"github.com/haji-sudo/ShabehRoshan/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var (
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		dbName   = os.Getenv("DB_NAME")
		uri      = fmt.Sprintf("postgres://%v:%v@%v:%v/%v", user, password, host, port, dbName)
	)

	var err error
	DB, err = gorm.Open(
		postgres.Open(uri),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatalln(err)
	}

	autoMigrateTables()
}

func autoMigrateTables() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Token{},
		&models.Profile{},
		&models.SocialLogin{},
		&models.Post{},
		&models.Comment{},
		&models.Like{},
		&models.Tag{},
		&models.Category{},
		&models.Subscription{},
	)
	if err != nil {
		log.Fatalln(err)
	}
}
