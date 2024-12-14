package initializers

import (
	"fmt"
	"os"

	"github.com/xuhe2/intelligence-software-expr3/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// database
var Db *gorm.DB

// connect to pgsql database
func ConnectToDB() {
	// load db info from .env file
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	// make connection
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", DB_USER, DB_PASSWORD, DB_NAME)
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// auto create table
	Db.AutoMigrate(&models.User{})
}