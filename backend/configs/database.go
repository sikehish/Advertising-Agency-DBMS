package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() {

	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

	db, err := gorm.Open(mysql.Open(dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local"),&gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	db.LogMode(true)

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
