package database

import (
	"fmt"
	"os"

	"github.com/riyan-eng/go-restfull-api-psql/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	var err error
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_TIMEZONE"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Connection failed to Database")
	}

	fmt.Println("Connection Opened to Database")

	// migrate database
	DB.AutoMigrate(&models.Note{})
	fmt.Println("Database migrated")
}
