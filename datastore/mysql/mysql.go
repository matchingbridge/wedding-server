package mysql

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"wedding/models"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("RDS_HOSTNAME")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("RDS_PORT")
	if port == "" {
		port = "3306"
	}
	database := os.Getenv("RDS_DB_NAME")
	if database == "" {
		database = "wedding"
	}
	user := os.Getenv("RDS_USERNAME")
	if user == "" {
		user = "root"
	}
	password := os.Getenv("RDS_PASSWORD")
	if password == "" {
		password = "gusqo123"
	}
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=%s", user, password, host, port, database, "Asia%2FSeoul")
	var err error
	DB, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func CreateTable() {
	tables := []interface{}{
		(*models.Auth)(nil),
		(*models.Chat)(nil),
		(*models.Match)(nil),
		(*models.Review)(nil),
		(*models.Suggestion)(nil),
		(*models.User)(nil),
	}

	if err := DB.AutoMigrate(tables...); err != nil {
		panic(err)
	}
}
