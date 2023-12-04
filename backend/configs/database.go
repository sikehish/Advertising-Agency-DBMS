package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sikehish/Advertising-Agency-DBMS/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	db, err := gorm.Open(mysql.Open(dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local"),&gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	DB = db

	DB.AutoMigrate(&models.Client{})
}

func GetDB() *gorm.DB {
	return DB
}
